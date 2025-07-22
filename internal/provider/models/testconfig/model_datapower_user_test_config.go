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

var UserTestConfig = ModelTestConfig{
	Name: "User",
	Resource: `
resource "datapower_user" "test" {
  id = "0user"
  password = "Password$123"
  access_level = "group-defined"
  group_name = datapower_usergroup.test.id
  snmp_creds = null
  hashed_snmp_creds = null
}`,
	Data: `
data "datapower_user" "test" {
  depends_on = [ datapower_user.test ]
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"UserGroup":        &UserGroupTestConfig,
		"DmSnmpCred":       &DmSnmpCredTestConfig,
		"DmSnmpCredMasked": &DmSnmpCredMaskedTestConfig,
	},
	ReferencesTo: map[string]*ModelTestConfig{
		"UserGroup": &UserGroupTestConfig,
	},
	TestPre: `
`,
}
