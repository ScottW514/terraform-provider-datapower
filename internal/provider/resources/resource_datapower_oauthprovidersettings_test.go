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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models/testconfig"
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccResourceOAuthProviderSettings(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_OAuthProviderSettings") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_OAuthProviderSettings")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testconfig.OAuthProviderSettingsTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "enable_debug_mode", "false"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "provider_type", "native"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "apic_provider_base_path", "/"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "apic_authorize_endpoint", "/oauth2/authorize"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "apic_token_endpoint", "/oauth2/token"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "apic_enable_introspection", "false"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "apic_introspect_endpoint", "/oauth2/introspect"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "apic_one_time_use_accesstoken", "false"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "apic_access_token_ttl", "3600"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "apic_auth_code_ttl", "300"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "apic_enable_refresh_token", "false"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "apic_one_time_use_refreshtoken", "true"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "apic_refresh_token_limit", "10"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "apic_refresh_token_ttl", "5400"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "apic_maximum_consent_ttl", "0"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "adv_scope_validation_enabled", "false"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "advanced_scope_url_override", "false"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "adv_scope_url_security_enabled", "false"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "apic_enable_oidc", "true"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "apic_support_pkce", "true"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "apic_require_pkce", "false"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "apic_support_pkce_plain", "false"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "apic_token_type_to_generate", "Bearer"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "metadata_header_for_access_token", "X-API-OAuth-Metadata-For-AccessToken"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "metadata_header_for_payload", "X-API-OAuth-Metadata-For-Payload"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "external_revocation_cache_type", "NoCache"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "external_revocation_cache_time_to_live", "900"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "external_revocation_fail_on_error", "true"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "enable_token_management", "true"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "token_manager_type", "native"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "api_security_token_manager", "default"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "enable_application_revocation", "false"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "application_revocation_endpoint", "/oauth2/revoke"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "enable_owner_revocation", "false"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "owner_revocation_endpoint", "/oauth2/issued"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "third_party_introspect_cache_type", "NoCache"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "third_party_introspect_cache_time_to_live", "900"),
			resource.TestCheckResourceAttr("datapower_oauthprovidersettings.test", "third_party_authorization_header_pass_thru", "false"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
