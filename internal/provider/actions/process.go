// Copyright © 2025 Scott Wiederhold <s.e.wiederhold@gmail.com>
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package actions

import (
	"context"
	"fmt"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
)

type processData struct {
	// [domain][objectType][objectId]{action,...}
	saveDomains []string
	Provider    *tfutils.ProviderData
	stateMu     sync.Mutex
}

var Process = processData{
	stateMu: sync.Mutex{},
}

type qstate int

const (
	unquiesced qstate = iota
	quiescing
	quiesced
	unquiescing
	errorState
)

// Called by resources prior to create/update runs.
// Triggers quiesce calls, and queues the rest for post-process execution. Duplicate actions on the same domain/resource are ignored.
func PreProcess(ctx context.Context, diag *diag.Diagnostics, srcDomain string, actions []*DependencyAction, operation Operation, saveDefault bool) {
	Process.stateMu.Lock()
	if !slices.Contains(Process.saveDomains, srcDomain) {
		Process.saveDomains = append(Process.saveDomains, srcDomain)
	}
	if saveDefault && !slices.Contains(Process.saveDomains, "default") {
		Process.saveDomains = append(Process.saveDomains, "default")
	}
	Process.stateMu.Unlock()

	for _, item := range actions {
		var name string
		if (operation == Create && !item.OnCreate.ValueBool()) ||
			(operation == Update && !item.OnUpdate.ValueBool()) ||
			(operation == Delete && !item.OnDelete.ValueBool()) ||
			item.Action.ValueString() != "quiesce" {
			continue
		}

		if item.TargetType.ValueString() == "datapower_domain" {
			name = item.TargetDomain.ValueString()
		} else {
			name = item.TargetId.ValueString()
		}

		submitAction(ctx, diag, item.TargetDomain.ValueString(), item.TargetType.ValueString(), name, item.Action.ValueString(), false)

		if diag.HasError() {
			return
		}

	}
}

// Called prior to exit of resource Create, Update and Delete operations
func PostProcess(ctx context.Context, diag *diag.Diagnostics, actions []*DependencyAction, operation Operation) {
	for _, item := range actions {
		if (operation == Create && !item.OnCreate.ValueBool()) ||
			(operation == Update && !item.OnUpdate.ValueBool()) ||
			(operation == Delete && !item.OnDelete.ValueBool()) {
			continue
		}

		domain := item.TargetDomain.ValueString()
		targetType := item.TargetType.ValueString()
		action := item.Action.ValueString()
		name := item.TargetId.ValueString()

		if targetType == "datapower_domain" {
			name = domain
		}

		submitAction(ctx, diag, domain, targetType, name, action, true)

		if diag.HasError() {
			return
		}
	}
}

// Called upon exit of provider Serve method
func SaveDomains() error {
	for _, domain := range Process.saveDomains {
		res, err := Process.Provider.Client.Post(fmt.Sprintf("/mgmt/actionqueue/%s", domain), `{"SaveConfig": 0}`)
		if err == nil {
			if res.StatusCode() != 200 {
				return fmt.Errorf("failed to save domain `%s`. status: %d body: %s", domain, res.StatusCode(), res.Body())
			}
		} else if !strings.Contains(err.Error(), "status 401") {
			return fmt.Errorf("failed to save domain `%s`. error: %s", domain, err.Error())
		}
	}
	return nil
}

func submitAction(ctx context.Context, diag *diag.Diagnostics, domain, objectType, objectId, action string, postAction bool) {
	postStr := "pre"
	if postAction {
		postStr = "post"
	}

	// If object already in proper quiesce state, skip
	if action == "quiesce" {
		qState, err := getQuiesceState(domain, objectType, objectId)
		if err != nil {
			diag.AddError("Action Error", fmt.Sprintf("Failed to %s-%s object (%s/%s/%s), got error: %s",
				postStr, action, domain, objectType, objectId, err))
			return
		}
		if (!postAction && qState == quiesced) || (postAction && qState == unquiesced) {
			return
		}
	}

	var cbUrl = ""
	res, err := Process.Provider.Client.Post(fmt.Sprintf("/mgmt/actionqueue/%s", domain), getBody(domain, objectType, objectId, action, postAction))
	if err == nil {
		if res.StatusCode() == 200 {
			return
		} else if res.StatusCode() != 202 {
			diag.AddError("Action Error", fmt.Sprintf("Failed to %s-%s object (%s/%s/%s), during operation request expected 202 got %d",
				postStr, action, domain, objectType, objectId, res.StatusCode()))
			return
		}
		if ack := gjson.ParseBytes(res.Body()).Get("_links.location.href"); ack.Exists() {
			cbUrl = ack.String()
		} else {
			diag.AddError("Action Error", fmt.Sprintf("Failed to %s-%s object (%s/%s/%s), invalid response during operation request: %s",
				postStr, action, domain, objectType, objectId, res.Body()))
			return
		}
	} else {
		diag.AddError("Action Error", fmt.Sprintf("Failed to %s-%s object (%s/%s/%s), got error during operation request: %s",
			postStr, action, domain, objectType, objectId, err))
		return
	}

	var exitLoop = false
	timer := time.NewTimer(time.Second * 90)
	ticker := time.NewTicker(time.Millisecond * 500)
	for {
		if exitLoop {
			break
		}
		select {
		case <-timer.C:
			diag.AddError("Action Error", fmt.Sprintf("Failed to %s-%s object (%s/%s/%s), timeout while waiting for action submission status",
				postStr, action, domain, objectType, objectId))
			return
		case <-ctx.Done():
			diag.AddError("Action Error", fmt.Sprintf("Failed to %s-%s object (%s/%s/%s), context expired before completion",
				postStr, action, domain, objectType, objectId))
			return
		case <-ticker.C:
			status, resBody, err := getOperationState(diag, cbUrl)
			if err != nil && !strings.Contains(err.Error(), "status 404") {
				diag.AddError("Action Error", fmt.Sprintf("Failed to %s-%s object (%s/%s/%s), error while reading action submission status: %s",
					postStr, action, domain, objectType, objectId, err))
				return
			} else if err != nil && strings.Contains(err.Error(), "status 404") {
				exitLoop = true
			} else if status == "completed" {
				exitLoop = true
			} else if status != "started" {
				diag.AddError("Action Error", fmt.Sprintf("Failed to %s-%s object (%s/%s/%s), unexpected status (%s) while reading action submission status: %s",
					postStr, action, domain, objectType, objectId, status, resBody))
				return
			}
		}
	}

	if action == "quiesce" {
		timer := time.NewTimer(time.Second * 90)
		ticker := time.NewTicker(time.Millisecond * 500)
		for {
			select {
			case <-timer.C:
				diag.AddError("Action Error", fmt.Sprintf("Failed to %s-%s object (%s/%s/%s), timeout while waiting for object to reach target state",
					postStr, action, domain, objectType, objectId))
				return
			case <-ctx.Done():
				diag.AddError("Action Error", fmt.Sprintf("Failed to %s-%s object (%s/%s/%s), context expired before completion",
					postStr, action, domain, objectType, objectId))
				return
			case <-ticker.C:
				qState, err := getQuiesceState(domain, objectType, objectId)
				if err != nil {
					diag.AddError("Action Error", fmt.Sprintf("Failed to %s-%s object (%s/%s/%s), got error: %s",
						postStr, action, domain, objectType, objectId, err))
					return
				}
				if (!postAction && qState == quiesced) || (postAction && qState == unquiesced) {
					return
				}
			}
		}
	} else if action == "restart" && objectType == "datapower_domain" {
		timer := time.NewTimer(time.Second * 90)
		ticker := time.NewTicker(time.Millisecond * 500)
		for {
			select {
			case <-timer.C:
				diag.AddError("Action Error", fmt.Sprintf("Failed to %s-%s object (%s/%s/%s), timeout while waiting for domain to restart",
					postStr, action, domain, objectType, objectId))
				return
			case <-ctx.Done():
				diag.AddError("Action Error", fmt.Sprintf("Failed to %s-%s object (%s/%s/%s), context expired before completion",
					postStr, action, domain, objectType, objectId))
				return
			case <-ticker.C:
				_, cErr := Process.Provider.Client.Get(fmt.Sprintf("/mgmt/actionqueue/%s/operations", domain))
				if cErr != nil && !strings.Contains(cErr.Error(), "status 404") {
					diag.AddError("Action Error", fmt.Sprintf("Failed to %s-%s object (%s/%s/%s), while waiting for domain to restart got error: %s",
						postStr, action, domain, objectType, objectId, cErr))
					return
				} else if cErr == nil {
					return
				}
			}
		}
	}
}

func getBody(domain, objectType, objectId, action string, postAction bool) string {
	var body string
	if action == "quiesce" {
		if objectType == "datapower_domain" && !postAction {
			body = domainQuiesceBody
		} else if objectType == "datapower_domain" && postAction {
			body = domainUnquiesceBody
		} else if objectType != "datapower_domain" && !postAction {
			body = serviceQuiesceBody
		} else {
			body = serviceUnquiesceBody
		}
	} else {
		body = actionMap[objectType].ValidActions[action].Body
	}
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(body, "{type}", actionMap[objectType].ObjectName), "{name}", objectId), "{domain}", domain)
}

func getQuiesceState(domain, objectType, objectId string) (qstate, error) {
	sRes, err := Process.Provider.Client.Get(fmt.Sprintf("/mgmt/status/%s/%sSummary", domain, actionMap[objectType].ObjectName))
	if err != nil {
		return errorState, err
	}
	if summary := sRes.Get(fmt.Sprintf("%sSummary", actionMap[objectType].ObjectName)); summary.Exists() {
		// DataPower returns single objects as an object in the "{ObjectName}Summary" key, and multiple objects as an array
		if summary.IsArray() {
			// Mulitple objects received
			for _, obj := range summary.Array() {
				namePath := fmt.Sprintf("%s.value", actionMap[objectType].ObjectName)
				if name := obj.Get(namePath); name.Exists() {
					if name.String() != objectId {
						continue
					} else if status := obj.Get("Quiesce"); status.Exists() {
						return stringToQstate(status.String()), nil
					}
				}
			}
			return errorState, fmt.Errorf("object not found in Summary")
		} else {
			// Single object recieved
			namePath := fmt.Sprintf("%s.value", actionMap[objectType].ObjectName)
			if name := summary.Get(namePath); name.Exists() {
				if status := summary.Get("Quiesce"); status.Exists() {
					return stringToQstate(status.String()), nil
				}
			}
		}
	}
	return errorState, fmt.Errorf("invalid Summary response from host: %s", sRes.Raw)
}

func getOperationState(diag *diag.Diagnostics, cbUrl string) (string, string, error) {
	cbRes, err := Process.Provider.Client.Get(cbUrl)
	if err == nil {
		if status := cbRes.Get("status"); status.Exists() {
			if status.String() == "error" {
				return "", "", fmt.Errorf("%s", cbRes.Raw)
			} else {
				return status.String(), cbRes.Raw, nil
			}
		} else {
			return "", "", fmt.Errorf("%s", cbRes.Raw)
		}
	}
	return "", "", err
}

func stringToQstate(qstring string) qstate {
	switch qstring {
	case "quiesced":
		return quiesced
	case "quiescing":
		return quiescing
	case "unquiescing":
		return unquiescing
	default:
		return unquiesced
	}
}
