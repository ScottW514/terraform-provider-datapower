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

var B2BPersistenceTestConfig = ModelTestConfig{
	Name: "B2BPersistence",
	Resource: `
resource "datapower_b2bpersistence" "test" {
  raid_volume = "raid0"
  storage_size = 1024
  ha_enabled = false
}`,
	Data: `
data "datapower_b2bpersistence" "test" {
  depends_on = [ datapower_b2bpersistence.test ]
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"RaidVolume": &RaidVolumeTestConfig,
	},
	ReferencesTo: map[string]*ModelTestConfig{
		"RaidVolume": &RaidVolumeTestConfig,
	},
	TestPre: `
`,
}
