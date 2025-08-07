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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/testconfig"
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccDataSourceB2BCPASenderSetting(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_B2BCPASenderSetting") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_B2BCPASenderSetting")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.B2BCPASenderSettingTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_b2bcpasendersetting.test", "result.0.id", "AccTest_B2BCPASenderSetting"),
					resource.TestCheckResourceAttr("data.datapower_b2bcpasendersetting.test", "result.0.connection_timeout", "300"),
					resource.TestCheckResourceAttr("data.datapower_b2bcpasendersetting.test", "result.0.sync_reply_mode", "none"),
					resource.TestCheckResourceAttr("data.datapower_b2bcpasendersetting.test", "result.0.duplicate_elimination", "always"),
					resource.TestCheckResourceAttr("data.datapower_b2bcpasendersetting.test", "result.0.ack_requested", "never"),
					resource.TestCheckResourceAttr("data.datapower_b2bcpasendersetting.test", "result.0.ack_signature_requested", "never"),
					resource.TestCheckResourceAttr("data.datapower_b2bcpasendersetting.test", "result.0.retry", "false"),
					resource.TestCheckResourceAttr("data.datapower_b2bcpasendersetting.test", "result.0.max_retries", "3"),
					resource.TestCheckResourceAttr("data.datapower_b2bcpasendersetting.test", "result.0.retry_interval", "1800"),
					resource.TestCheckResourceAttr("data.datapower_b2bcpasendersetting.test", "result.0.include_time_to_live", "true"),
					resource.TestCheckResourceAttr("data.datapower_b2bcpasendersetting.test", "result.0.encryption_required", "false"),
					resource.TestCheckResourceAttr("data.datapower_b2bcpasendersetting.test", "result.0.encrypt_algorithm", "http://www.w3.org/2001/04/xmlenc#tripledes-cbc"),
					resource.TestCheckResourceAttr("data.datapower_b2bcpasendersetting.test", "result.0.signature_required", "false"),
					resource.TestCheckResourceAttr("data.datapower_b2bcpasendersetting.test", "result.0.signature_algorithm", "rsa-sha1"),
					resource.TestCheckResourceAttr("data.datapower_b2bcpasendersetting.test", "result.0.sign_digest_algorithm", "sha1"),
					resource.TestCheckResourceAttr("data.datapower_b2bcpasendersetting.test", "result.0.signature_c14n_algorithm", "exc-c14n"),
					resource.TestCheckResourceAttr("data.datapower_b2bcpasendersetting.test", "result.0.ssl_client_config_type", "client"),
				}...),
			},
		},
	})
}
