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

func TestAccDataSourceOAuthProviderSettings(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_OAuthProviderSettings") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_OAuthProviderSettings")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.OAuthProviderSettingsTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.id", "OAuthProviderSettings_name"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.enable_debug_mode", "false"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.provider_type", "native"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.apic_provider_base_path", "/"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.apic_authorize_endpoint", "/oauth2/authorize"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.apic_token_endpoint", "/oauth2/token"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.apic_enable_introspection", "false"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.apic_introspect_endpoint", "/oauth2/introspect"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.apic_one_time_use_accesstoken", "false"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.apic_access_token_ttl", "3600"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.apic_auth_code_ttl", "300"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.apic_enable_refresh_token", "false"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.apic_one_time_use_refreshtoken", "true"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.apic_refresh_token_limit", "10"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.apic_refresh_token_ttl", "5400"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.apic_maximum_consent_ttl", "0"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.adv_scope_validation_enabled", "false"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.advanced_scope_url_override", "false"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.adv_scope_url_security_enabled", "false"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.apic_enable_oidc", "true"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.apic_support_pkce", "true"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.apic_require_pkce", "false"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.apic_support_pkce_plain", "false"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.apic_token_type_to_generate", "Bearer"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.metadata_header_for_access_token", "X-API-OAuth-Metadata-For-AccessToken"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.metadata_header_for_payload", "X-API-OAuth-Metadata-For-Payload"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.external_revocation_cache_type", "NoCache"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.external_revocation_cache_time_to_live", "900"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.external_revocation_fail_on_error", "true"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.enable_token_management", "true"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.token_manager_type", "native"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.api_security_token_manager", "default"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.enable_application_revocation", "false"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.application_revocation_endpoint", "/oauth2/revoke"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.enable_owner_revocation", "false"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.owner_revocation_endpoint", "/oauth2/issued"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.third_party_introspect_cache_type", "NoCache"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.third_party_introspect_cache_time_to_live", "900"),
					resource.TestCheckResourceAttr("data.datapower_oauthprovidersettings.test", "result.0.third_party_authorization_header_pass_thru", "false"),
				}...),
			},
		},
	})
	actions.PostProcess()
}
