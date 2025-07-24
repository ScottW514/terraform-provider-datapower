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

// This file is generated "gen/generator.go"
// !!CHANGES TO THIS FILE WILL BE OVERWRITTEN!!

package testconfig

var APICollectionTestConfig = ModelTestConfig{
	Name: "APICollection",
	Resource: `
resource "datapower_apicollection" "test" {
  id = "APICollection_name"
  app_domain = "acc_test_domain"
  org_id = "orgid"
  org_name = "orgname"
  routing_prefix = ` + DmRoutingPrefixTestConfig.GetModelListConfig() + `
  api_processing_rule = "default-api-rule"
  api_error_rule = "default-api-error-rule"
  plan = [datapower_apiplan.test.id]
}`,
	Data: `
data "datapower_apicollection" "test" {
  depends_on = [ datapower_apicollection.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"DmRoutingPrefix": &DmRoutingPrefixTestConfig,
		"APIRule":         &APIRuleTestConfig,
		"APIPlan":         &APIPlanTestConfig,
	},
	ReferencesTo: map[string]*ModelTestConfig{
		"APIRule": &APIRuleTestConfig,
		"APIPlan": &APIPlanTestConfig,
	},
	TestPre: `
`,
}
