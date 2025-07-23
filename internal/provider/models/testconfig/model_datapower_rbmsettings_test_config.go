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

var RBMSettingsTestConfig = ModelTestConfig{
	Name: "RBMSettings",
	Resource: `
resource "datapower_rbmsettings" "test" {
  au_method = "local"
  au_cache_allow = "absolute"
  mc_method = "local"
  min_password_length = 6
  require_mixed_case = false
  require_digit = false
  require_non_alpha_numeric = false
  disallow_username_substring = false
  do_password_aging = false
  do_password_history = false
  cli_timeout = 0
  max_failed_login = 0
}`,
	Data: `
data "datapower_rbmsettings" "test" {
  depends_on = [ datapower_rbmsettings.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	ReferencesTo: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
