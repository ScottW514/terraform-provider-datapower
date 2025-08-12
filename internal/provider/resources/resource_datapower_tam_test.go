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

package resources_test

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccResourceTAM(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_TAM") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_TAM")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testutils.TAMTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_tam.test", "ad_use_ad", "false"),
			resource.TestCheckResourceAttr("datapower_tam.test", "use_local_mode", "false"),
			resource.TestCheckResourceAttr("datapower_tam.test", "poll_interval", "default"),
			resource.TestCheckResourceAttr("datapower_tam.test", "listen_mode", "false"),
			resource.TestCheckResourceAttr("datapower_tam.test", "returning_user_attributes", "false"),
			resource.TestCheckResourceAttr("datapower_tam.test", "ldap_use_ssl", "false"),
			resource.TestCheckResourceAttr("datapower_tam.test", "ldapssl_port", "636"),
			resource.TestCheckResourceAttr("datapower_tam.test", "tam_use_fips", "false"),
			resource.TestCheckResourceAttr("datapower_tam.test", "tam_use_basic_user", "false"),
			resource.TestCheckResourceAttr("datapower_tam.test", "user_principal_attribute", "uid"),
			resource.TestCheckResourceAttr("datapower_tam.test", "user_no_duplicates", "true"),
			resource.TestCheckResourceAttr("datapower_tam.test", "user_suffix_optimiser", "false"),
			resource.TestCheckResourceAttr("datapower_tam.test", "auto_retry", "false"),
			resource.TestCheckResourceAttr("datapower_tam.test", "retry_interval", "180"),
			resource.TestCheckResourceAttr("datapower_tam.test", "retry_attempts", "3"),
			resource.TestCheckResourceAttr("datapower_tam.test", "long_retry_interval", "900"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
