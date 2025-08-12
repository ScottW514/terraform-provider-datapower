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

package datasources_test

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccDataSourceTAM(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_TAM") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_TAM")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testutils.TAMTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_tam.test", "result.0.id", "AccTest_TAM"),
					resource.TestCheckResourceAttr("data.datapower_tam.test", "result.0.ad_use_ad", "false"),
					resource.TestCheckResourceAttr("data.datapower_tam.test", "result.0.use_local_mode", "false"),
					resource.TestCheckResourceAttr("data.datapower_tam.test", "result.0.poll_interval", "default"),
					resource.TestCheckResourceAttr("data.datapower_tam.test", "result.0.listen_mode", "false"),
					resource.TestCheckResourceAttr("data.datapower_tam.test", "result.0.returning_user_attributes", "false"),
					resource.TestCheckResourceAttr("data.datapower_tam.test", "result.0.ldap_use_ssl", "false"),
					resource.TestCheckResourceAttr("data.datapower_tam.test", "result.0.ldapssl_port", "636"),
					resource.TestCheckResourceAttr("data.datapower_tam.test", "result.0.tam_use_fips", "false"),
					resource.TestCheckResourceAttr("data.datapower_tam.test", "result.0.tam_use_basic_user", "false"),
					resource.TestCheckResourceAttr("data.datapower_tam.test", "result.0.user_principal_attribute", "uid"),
					resource.TestCheckResourceAttr("data.datapower_tam.test", "result.0.user_no_duplicates", "true"),
					resource.TestCheckResourceAttr("data.datapower_tam.test", "result.0.user_suffix_optimiser", "false"),
					resource.TestCheckResourceAttr("data.datapower_tam.test", "result.0.auto_retry", "false"),
					resource.TestCheckResourceAttr("data.datapower_tam.test", "result.0.retry_interval", "180"),
					resource.TestCheckResourceAttr("data.datapower_tam.test", "result.0.retry_attempts", "3"),
					resource.TestCheckResourceAttr("data.datapower_tam.test", "result.0.long_retry_interval", "900"),
				}...),
			},
		},
	})
}
