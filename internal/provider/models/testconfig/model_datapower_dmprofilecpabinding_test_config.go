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

var DmProfileCPABindingTestConfig = ModelTestConfig{
	Name: "DmProfileCPABinding",
	Model: `{
  internal_partner = datapower_b2bprofile.test.id
  cpa = datapower_b2bcpa.test.id
  collaboration = datapower_b2bcpacollaboration.test.id
}`,
	ModelOnly: true,
	Dependencies: map[string]*ModelTestConfig{
		"B2BProfile":          &B2BProfileTestConfig,
		"B2BCPA":              &B2BCPATestConfig,
		"B2BCPACollaboration": &B2BCPACollaborationTestConfig,
	},
	ReferencesTo: map[string]*ModelTestConfig{
		"B2BProfile":          &B2BProfileTestConfig,
		"B2BCPA":              &B2BCPATestConfig,
		"B2BCPACollaboration": &B2BCPACollaborationTestConfig,
	},
	TestPre: `
`,
}
