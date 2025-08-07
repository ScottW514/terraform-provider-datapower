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
	"github.com/scottw514/terraform-provider-datapower/client"
)

type queue struct {
	// [domain][objectType][objectId]{action,...}
	DomainActions map[string]map[string]map[string][]string
	ObjectActions map[string]map[string]map[string][]string
	SaveDomains   []string
	Client        *client.DatapowerClient
	mu            sync.Mutex
}

var actionQueue = queue{}

// Called by resources prior to create/update runs.
// Triggers quiesce calls, and queues the rest for post-process execution. Duplicate actions on the same domain/resource are ignored.
func PreProcess(ctx context.Context, diag *diag.Diagnostics, client *client.DatapowerClient, srcDomain string, actions []*Action, operation Operation) {
	actionQueue.mu.Lock()
	defer actionQueue.mu.Unlock()
	if !slices.Contains(actionQueue.SaveDomains, srcDomain) {
		actionQueue.SaveDomains = append(actionQueue.SaveDomains, srcDomain)
	}
	actionQueue.Client = client
	for _, item := range actions {
		var name string
		if (operation == Create && !item.OnCreate.ValueBool()) ||
			(operation == Update && !item.OnUpdate.ValueBool()) ||
			(operation == Delete && !item.OnDelete.ValueBool()) ||
			alreadyQueued(item.TargetDomain.ValueString(), item.TargetType.ValueString(), item.TargetId.ValueString(), item.Action.ValueString()) {
			continue
		}
		if item.TargetType.ValueString() == "resource_datapower_domain" {
			name = "default"
		} else {
			name = item.TargetId.ValueString()
		}

		if item.Action.ValueString() == "quiesce" {
			quiesceAction(diag, item.TargetDomain.ValueString(), item.TargetType.ValueString(), name, true)
		}
		if diag.HasError() {
			return
		}

		if item.TargetType.ValueString() == "resource_datapower_domain" {
			addToQueue(&actionQueue.DomainActions, item.TargetDomain.ValueString(), item.TargetType.ValueString(), name, item.Action.ValueString())
		} else {
			addToQueue(&actionQueue.ObjectActions, item.TargetDomain.ValueString(), item.TargetType.ValueString(), name, item.Action.ValueString())
		}
	}
}

// Called upon exit of provider Serve method, and at end of tests
func PostProcess() error {
	diag := diag.Diagnostics{}
	processQueuedActions(&diag, &actionQueue.ObjectActions)
	if !diag.HasError() {
		processQueuedActions(&diag, &actionQueue.DomainActions)
	}
	if diag.HasError() {
		for _, e := range diag.Errors() {
			return fmt.Errorf("%s: %s", e.Summary(), e.Detail())
		}
	}

	for _, domain := range actionQueue.SaveDomains {
		res, err := actionQueue.Client.Post(fmt.Sprintf("/mgmt/actionqueue/%s", domain), `{"SaveConfig": 0}`)
		if err == nil {
			if res.StatusCode() != 404 && res.StatusCode() != 401 {
				return fmt.Errorf("failed to save domain `%s`. status: %d body: %s", domain, res.StatusCode(), res.Body())
			}
		} else {
			return fmt.Errorf("failed to save domain `%s`. error: %s", domain, err.Error())
		}
	}
	return nil
}

func quiesceAction(diag *diag.Diagnostics, domain, objectType, objectId string, quiesce bool) {
	quiesceAction := "quiesce"
	if !quiesce {
		quiesceAction = "unquiesce"
	}

	res, err := actionQueue.Client.Post(fmt.Sprintf("/mgmt/actionqueue/%s", domain), getBody(domain, objectType, objectId, "quiesce", quiesce))
	if err == nil {
		if res.StatusCode() != 202 {
			diag.AddError("Action Error", fmt.Sprintf("Failed to %s object (%s/%s/%s), expected 202 got %d",
				quiesceAction, domain, objectType, objectId, res.StatusCode()))
		}
	} else {
		diag.AddError("Action Error", fmt.Sprintf("Failed to %s object (%s/%s/%s), got error: %s",
			quiesceAction, domain, objectType, objectId, err))
	}

	if diag.HasError() {
		return
	}

	if quiesce {
		// We wait for the object to reach quiesce state for up to 60 seconds
		for attempt := 1; attempt <= 120; attempt++ {
			sRes, err := actionQueue.Client.Get(fmt.Sprintf("/mgmt/status/%s/%sSummary", domain, actionMap[objectType].ObjectName))
			if err != nil {
				diag.AddError("Action Error", fmt.Sprintf("Failed to quiesce object (%s/%s/%s), got error: %s", domain, objectType, objectId, err))
				break
			}
			if summary := sRes.Get(fmt.Sprintf("%sSummary", actionMap[objectType].ObjectName)); summary.Exists() {
				for _, obj := range summary.Array() {
					if name := obj.Get(fmt.Sprintf("%s.value", actionMap[objectType].ObjectName)); name.Exists() {
						if status := obj.Get("Quiesce"); status.Exists() && status.String() == "quiesced" {
							return
						}
					}
				}
			} else {
				diag.AddError("Action Error", fmt.Sprintf("Failed while waiting to quiesce object (%s/%s/%s), no object status returned: %s",
					domain, objectType, objectId, sRes.Raw))
				break
			}
			time.Sleep(time.Millisecond * 500)
		}
		diag.AddError("Action Error", fmt.Sprintf("Timeout reached while waiting to quiesce object (%s/%s/%s)", domain, objectType, objectId))
	}
}

func getBody(domain, objectType, objectId, action string, pre bool) string {
	var body string
	if action == "quiesce" {
		if objectType == "resource_datapower_domain" && pre {
			body = DomainQuiesceBody
		} else if objectType == "resource_datapower_domain" && !pre {
			body = DomainUnquiesceBody
		} else if objectType != "resource_datapower_domain" && pre {
			body = ServiceQuiesceBody
		} else {
			body = ServiceUnquiesceBody
		}
	} else {
		body = actionMap[objectType].ValidActions[action].Body
	}
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(body, "{type}", actionMap[objectType].ObjectName), "{name}", objectId), "{domain}", domain)
}

func alreadyQueued(domain, objectType, objectId, action string) bool {
	if searchQueue(actionQueue.DomainActions, domain, objectType, objectId, action) {
		return true
	}
	if searchQueue(actionQueue.ObjectActions, domain, objectType, objectId, action) {
		return true
	}
	return false
}

func searchQueue(queue map[string]map[string]map[string][]string, domain, objectType, objectId, action string) bool {
	if a, ok := queue[domain]; ok {
		if b, ok := a[objectType]; ok {
			if c, ok := b[objectId]; ok {
				if slices.Contains(c, action) {
					return true
				}
			}
		}
	}
	return false
}

func addToQueue(queue *map[string]map[string]map[string][]string, domain, objectType, objectId, action string) {
	if *queue == nil {
		*queue = make(map[string]map[string]map[string][]string)
	}
	if (*queue)[domain] == nil {
		(*queue)[domain] = make(map[string]map[string][]string)
	}
	if (*queue)[domain][objectType] == nil {
		(*queue)[domain][objectType] = make(map[string][]string)
	}
	(*queue)[domain][objectType][objectId] = append((*queue)[domain][objectType][objectId], action)
}

func processQueuedActions(diag *diag.Diagnostics, queue *map[string]map[string]map[string][]string) {
	for domain, typeMap := range *queue {
		for objectType, objectIdMap := range typeMap {
			for objectId, actionList := range objectIdMap {
				for _, action := range actionList {
					if action == "quiesce" {
						quiesceAction(diag, domain, objectType, objectId, false)
						if diag.HasError() {
							return
						}
					} else {
						res, err := actionQueue.Client.Post(fmt.Sprintf("/mgmt/actionqueue/%s", domain), getBody(domain, objectType, objectId, action, false))
						if err == nil {
							if res.StatusCode() != 202 {
								diag.AddError("Action Error", fmt.Sprintf("Failed to `%s` object (%s/%s/%s), expected 202 got %d",
									action, domain, objectType, objectId, res.StatusCode()))
								return
							}
						} else {
							diag.AddError("Action Error", fmt.Sprintf("Failed to `%s` object (%s/%s/%s), got error: %s",
								action, domain, objectType, objectId, err))
							return
						}

					}
				}
			}
		}
	}
}
