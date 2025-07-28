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

var AnalyticsEndpointTestConfig = ModelTestConfig{
	Name: "AnalyticsEndpoint",
	Resource: `
resource "datapower_analyticsendpoint" "test" {
  id = "AnalyticsEndpoint_name"
  app_domain = "acc_test_domain"
  analytics_server_url = "https://localhost"
  max_records = 1024
  max_records_memory_kb = 512
  max_delivery_memory_mb = 512
}`,
	Data: `
data "datapower_analyticsendpoint" "test" {
  depends_on = [ datapower_analyticsendpoint.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
