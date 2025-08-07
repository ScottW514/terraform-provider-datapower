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

type processData struct {
	// [domain][objectType][objectId]{action,...}
	saveDomains []string
	client      *client.DatapowerClient
	mu          sync.Mutex
}

var process = processData{}

// Called by resources prior to create/update runs.
// Triggers quiesce calls, and queues the rest for post-process execution. Duplicate actions on the same domain/resource are ignored.
func PreProcess(ctx context.Context, diag *diag.Diagnostics, client *client.DatapowerClient, srcDomain string, actions []*DependencyAction, operation Operation, saveDefault bool) {
	process.mu.Lock()
	defer process.mu.Unlock()
	process.client = client
	if !slices.Contains(process.saveDomains, srcDomain) {
		process.saveDomains = append(process.saveDomains, srcDomain)
	}
	if saveDefault && !slices.Contains(process.saveDomains, "default") {
		process.saveDomains = append(process.saveDomains, "default")
	}
	for _, item := range actions {
		var name string
		if (operation == Create && !item.OnCreate.ValueBool()) ||
			(operation == Update && !item.OnUpdate.ValueBool()) ||
			(operation == Delete && !item.OnDelete.ValueBool()) ||
			item.Action.ValueString() != "quiesce" {
			continue
		}
		if item.TargetType.ValueString() == "resource_datapower_domain" {
			name = "default"
		} else {
			name = item.TargetId.ValueString()
		}

		quiesceAction(diag, item.TargetDomain.ValueString(), item.TargetType.ValueString(), name, true)

		if diag.HasError() {
			return
		}

	}
}

// Called prior to exit of resource Create, Update and Delete operations
func PostProcess(ctx context.Context, diag *diag.Diagnostics, client *client.DatapowerClient, actions []*DependencyAction, operation Operation) {
	process.mu.Lock()
	defer process.mu.Unlock()
	process.client = client
	for _, item := range actions {
		var name string
		if (operation == Create && !item.OnCreate.ValueBool()) ||
			(operation == Update && !item.OnUpdate.ValueBool()) ||
			(operation == Delete && !item.OnDelete.ValueBool()) {
			continue
		}
		if item.TargetType.ValueString() == "resource_datapower_domain" {
			name = "default"
		} else {
			name = item.TargetId.ValueString()
		}

		domain := item.TargetDomain.ValueString()
		targetType := item.TargetType.ValueString()
		action := item.Action.ValueString()

		if action == "quiesce" {
			quiesceAction(diag, domain, targetType, name, false)
		} else {
			res, err := process.client.Post(fmt.Sprintf("/mgmt/actionqueue/%s", domain),
				getBody(domainUnquiesceBody, targetType, name, action, false))
			if err == nil {
				if res.StatusCode() != 202 {
					diag.AddError("Action Error", fmt.Sprintf("Failed to `%s` object (%s/%s/%s), expected 202 got %d",
						action, domain, targetType,
						name, res.StatusCode()))
					return
				}
			} else {
				diag.AddError("Action Error", fmt.Sprintf("Failed to `%s` object (%s/%s/%s), got error: %s",
					action, domain, targetType, name, err))
				return
			}

		}

		if diag.HasError() {
			return
		}

	}
}

// Called upon exit of provider Serve method
func SaveDomains() error {
	for _, domain := range process.saveDomains {
		res, err := process.client.Post(fmt.Sprintf("/mgmt/actionqueue/%s", domain), `{"SaveConfig": 0}`)
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

func quiesceAction(diag *diag.Diagnostics, domain, objectType, objectId string, quiesce bool) {
	quiesceAction := "quiesce"
	if !quiesce {
		quiesceAction = "unquiesce"
	}

	res, err := process.client.Post(fmt.Sprintf("/mgmt/actionqueue/%s", domain), getBody(domain, objectType, objectId, "quiesce", quiesce))
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
			sRes, err := process.client.Get(fmt.Sprintf("/mgmt/status/%s/%sSummary", domain, actionMap[objectType].ObjectName))
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
			body = domainQuiesceBody
		} else if objectType == "resource_datapower_domain" && !pre {
			body = domainUnquiesceBody
		} else if objectType != "resource_datapower_domain" && pre {
			body = serviceQuiesceBody
		} else {
			body = serviceUnquiesceBody
		}
	} else {
		body = actionMap[objectType].ValidActions[action].Body
	}
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(body, "{type}", actionMap[objectType].ObjectName), "{name}", objectId), "{domain}", domain)
}
