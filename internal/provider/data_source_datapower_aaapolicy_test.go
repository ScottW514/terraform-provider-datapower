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

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models/testconfig"
)

func TestAccDataSourceAAAPolicy(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_AAAPolicy") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_AAAPolicy")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.AAAPolicyTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_aaapolicy.test", "result.0.id", "AAAPolicy_name"),
					resource.TestCheckResourceAttr("data.datapower_aaapolicy.test", "result.0.log_allowed", "false"),
					resource.TestCheckResourceAttr("data.datapower_aaapolicy.test", "result.0.log_allowed_level", "info"),
					resource.TestCheckResourceAttr("data.datapower_aaapolicy.test", "result.0.log_rejected", "true"),
					resource.TestCheckResourceAttr("data.datapower_aaapolicy.test", "result.0.log_rejected_level", "warn"),
					resource.TestCheckResourceAttr("data.datapower_aaapolicy.test", "result.0.ping_identity_compatibility", "false"),
					resource.TestCheckResourceAttr("data.datapower_aaapolicy.test", "result.0.do_s_valve", "3"),
					resource.TestCheckResourceAttr("data.datapower_aaapolicy.test", "result.0.ldap_version", "v2"),
					resource.TestCheckResourceAttr("data.datapower_aaapolicy.test", "result.0.enforce_soap_actor", "true"),
					resource.TestCheckResourceAttr("data.datapower_aaapolicy.test", "result.0.dyn_config", "none"),
				}...),
			},
		},
	})
}
