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

var HTTPSSourceProtocolHandlerTestConfig = ModelTestConfig{
	Name: "HTTPSSourceProtocolHandler",
	Resource: `
resource "datapower_httpssourceprotocolhandler" "test" {
  id = "HTTPSSourceProtocolHandler_name"
  app_domain = "acc_test_domain"
  ssl_server_config_type = "server"
  ssl_server = datapower_sslserverprofile.test.id
}`,
	Data: `
data "datapower_httpssourceprotocolhandler" "test" {
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"SSLServerProfile": &SSLServerProfileTestConfig,
	},
	ReferencesTo: map[string]*ModelTestConfig{
		"SSLServerProfile": &SSLServerProfileTestConfig,
	},
	TestPre: `
`,
}
