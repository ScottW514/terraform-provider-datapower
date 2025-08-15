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

func TestAccDataSourceOAuthSupportedClient(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_OAuthSupportedClient") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_OAuthSupportedClient")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testutils.OAuthSupportedClientTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_oauth_supported_client.test", "result.0.id", "AccTest_OAuthSupportedClient"),
					resource.TestCheckResourceAttr("data.datapower_oauth_supported_client.test", "result.0.customized", "false"),
					resource.TestCheckResourceAttr("data.datapower_oauth_supported_client.test", "result.0.client_type", "confidential"),
					resource.TestCheckResourceAttr("data.datapower_oauth_supported_client.test", "result.0.check_client_credential", "false"),
					resource.TestCheckResourceAttr("data.datapower_oauth_supported_client.test", "result.0.use_validation_url", "false"),
					resource.TestCheckResourceAttr("data.datapower_oauth_supported_client.test", "result.0.client_authen_method", "secret"),
					resource.TestCheckResourceAttr("data.datapower_oauth_supported_client.test", "result.0.generate_client_secret", "false"),
					resource.TestCheckResourceAttr("data.datapower_oauth_supported_client.test", "result.0.caching", "replay"),
					resource.TestCheckResourceAttr("data.datapower_oauth_supported_client.test", "result.0.custom_scope_check", "false"),
					resource.TestCheckResourceAttr("data.datapower_oauth_supported_client.test", "result.0.local_az_page_url", "store:///OAuth-Generate-HTML.xsl"),
					resource.TestCheckResourceAttr("data.datapower_oauth_supported_client.test", "result.0.dp_state_life_time", "300"),
					resource.TestCheckResourceAttr("data.datapower_oauth_supported_client.test", "result.0.au_code_life_time", "300"),
					resource.TestCheckResourceAttr("data.datapower_oauth_supported_client.test", "result.0.access_token_life_time", "3600"),
					resource.TestCheckResourceAttr("data.datapower_oauth_supported_client.test", "result.0.refresh_token_life_time", "5400"),
					resource.TestCheckResourceAttr("data.datapower_oauth_supported_client.test", "result.0.custom_resource_owner", "false"),
					resource.TestCheckResourceAttr("data.datapower_oauth_supported_client.test", "result.0.validation_urlssl_client_type", "client"),
				}...),
			},
		},
	})
}
