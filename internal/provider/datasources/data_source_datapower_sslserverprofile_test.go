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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/testconfig"
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccDataSourceSSLServerProfile(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_SSLServerProfile") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_SSLServerProfile")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.SSLServerProfileTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_sslserverprofile.test", "result.0.id", "SSLServerProfile_name"),
					resource.TestCheckResourceAttr("data.datapower_sslserverprofile.test", "result.0.request_client_auth", "false"),
					resource.TestCheckResourceAttr("data.datapower_sslserverprofile.test", "result.0.require_client_auth", "true"),
					resource.TestCheckResourceAttr("data.datapower_sslserverprofile.test", "result.0.validate_client_cert", "true"),
					resource.TestCheckResourceAttr("data.datapower_sslserverprofile.test", "result.0.send_client_auth_ca_list", "true"),
					resource.TestCheckResourceAttr("data.datapower_sslserverprofile.test", "result.0.caching", "true"),
					resource.TestCheckResourceAttr("data.datapower_sslserverprofile.test", "result.0.cache_timeout", "300"),
					resource.TestCheckResourceAttr("data.datapower_sslserverprofile.test", "result.0.cache_size", "20"),
					resource.TestCheckResourceAttr("data.datapower_sslserverprofile.test", "result.0.max_ssl_duration", "60"),
					resource.TestCheckResourceAttr("data.datapower_sslserverprofile.test", "result.0.disable_renegotiation", "false"),
					resource.TestCheckResourceAttr("data.datapower_sslserverprofile.test", "result.0.number_of_renegotiation_allowed", "0"),
					resource.TestCheckResourceAttr("data.datapower_sslserverprofile.test", "result.0.prohibit_resume_on_reneg", "false"),
					resource.TestCheckResourceAttr("data.datapower_sslserverprofile.test", "result.0.compression", "false"),
					resource.TestCheckResourceAttr("data.datapower_sslserverprofile.test", "result.0.allow_legacy_renegotiation", "false"),
					resource.TestCheckResourceAttr("data.datapower_sslserverprofile.test", "result.0.prefer_server_ciphers", "true"),
					resource.TestCheckResourceAttr("data.datapower_sslserverprofile.test", "result.0.prioritize_cha_cha", "false"),
					resource.TestCheckResourceAttr("data.datapower_sslserverprofile.test", "result.0.require_closure_notification", "true"),
				}...),
			},
		},
	})
	actions.PostProcess()
}
