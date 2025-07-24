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

var MQManagerTestConfig = ModelTestConfig{
	Name: "MQManager",
	Resource: `
resource "datapower_mqmanager" "test" {
  id = "MQManager_name"
  app_domain = "acc_test_domain"
  host_name = "localhost"
  cache_timeout = 60
  xml_manager = "default"
}`,
	Data: `
data "datapower_mqmanager" "test" {
  depends_on = [ datapower_mqmanager.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"XMLManager": &XMLManagerTestConfig,
	},
	ReferencesTo: map[string]*ModelTestConfig{
		"XMLManager": &XMLManagerTestConfig,
	},
	TestPre: `
`,
}
