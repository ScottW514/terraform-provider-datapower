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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models/testconfig"
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccResourceB2BCPASenderSetting(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_B2BCPASenderSetting") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_B2BCPASenderSetting")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testconfig.B2BCPASenderSettingTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_b2bcpasendersetting.test", "connection_timeout", "300"),
			resource.TestCheckResourceAttr("datapower_b2bcpasendersetting.test", "sync_reply_mode", "none"),
			resource.TestCheckResourceAttr("datapower_b2bcpasendersetting.test", "duplicate_elimination", "always"),
			resource.TestCheckResourceAttr("datapower_b2bcpasendersetting.test", "ack_requested", "never"),
			resource.TestCheckResourceAttr("datapower_b2bcpasendersetting.test", "ack_signature_requested", "never"),
			resource.TestCheckResourceAttr("datapower_b2bcpasendersetting.test", "retry", "false"),
			resource.TestCheckResourceAttr("datapower_b2bcpasendersetting.test", "max_retries", "3"),
			resource.TestCheckResourceAttr("datapower_b2bcpasendersetting.test", "retry_interval", "1800"),
			resource.TestCheckResourceAttr("datapower_b2bcpasendersetting.test", "include_time_to_live", "true"),
			resource.TestCheckResourceAttr("datapower_b2bcpasendersetting.test", "encryption_required", "false"),
			resource.TestCheckResourceAttr("datapower_b2bcpasendersetting.test", "encrypt_algorithm", "http://www.w3.org/2001/04/xmlenc#tripledes-cbc"),
			resource.TestCheckResourceAttr("datapower_b2bcpasendersetting.test", "signature_required", "false"),
			resource.TestCheckResourceAttr("datapower_b2bcpasendersetting.test", "signature_algorithm", "rsa-sha1"),
			resource.TestCheckResourceAttr("datapower_b2bcpasendersetting.test", "sign_digest_algorithm", "sha1"),
			resource.TestCheckResourceAttr("datapower_b2bcpasendersetting.test", "signature_c14n_algorithm", "exc-c14n"),
			resource.TestCheckResourceAttr("datapower_b2bcpasendersetting.test", "ssl_client_config_type", "client"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
	actions.PostProcess()
}
