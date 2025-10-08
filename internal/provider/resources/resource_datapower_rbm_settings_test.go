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

func TestAccResourceRBMSettings(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_RBMSettings") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_RBMSettings")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testutils.RBMSettingsTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_rbm_settings.test", "enabled", "true"),
			resource.TestCheckResourceAttr("datapower_rbm_settings.test", "au_method", "local"),
			resource.TestCheckResourceAttr("datapower_rbm_settings.test", "au_oidc_scope", "openid"),
			resource.TestCheckResourceAttr("datapower_rbm_settings.test", "au_oidc_key_fetch_interval", "30"),
			resource.TestCheckResourceAttr("datapower_rbm_settings.test", "au_ldap_search_for_dn", "false"),
			resource.TestCheckResourceAttr("datapower_rbm_settings.test", "au_ldap_prefix", "cn="),
			resource.TestCheckResourceAttr("datapower_rbm_settings.test", "au_force_dn_ldap_order", "false"),
			resource.TestCheckResourceAttr("datapower_rbm_settings.test", "au_cache_allow", "absolute"),
			resource.TestCheckResourceAttr("datapower_rbm_settings.test", "au_cache_ttl", "600"),
			resource.TestCheckResourceAttr("datapower_rbm_settings.test", "au_ldap_read_timeout", "60"),
			resource.TestCheckResourceAttr("datapower_rbm_settings.test", "mc_method", "local"),
			resource.TestCheckResourceAttr("datapower_rbm_settings.test", "mc_ldap_search_for_group", "false"),
			resource.TestCheckResourceAttr("datapower_rbm_settings.test", "mc_ldap_read_timeout", "60"),
			resource.TestCheckResourceAttr("datapower_rbm_settings.test", "fallback_login", "disabled"),
			resource.TestCheckResourceAttr("datapower_rbm_settings.test", "apply_to_cli", "false"),
			resource.TestCheckResourceAttr("datapower_rbm_settings.test", "restrict_admin_to_serial_port", "false"),
			resource.TestCheckResourceAttr("datapower_rbm_settings.test", "min_password_length", "6"),
			resource.TestCheckResourceAttr("datapower_rbm_settings.test", "require_mixed_case", "false"),
			resource.TestCheckResourceAttr("datapower_rbm_settings.test", "require_digit", "false"),
			resource.TestCheckResourceAttr("datapower_rbm_settings.test", "require_non_alpha_numeric", "false"),
			resource.TestCheckResourceAttr("datapower_rbm_settings.test", "disallow_username_substring", "false"),
			resource.TestCheckResourceAttr("datapower_rbm_settings.test", "do_password_aging", "false"),
			resource.TestCheckResourceAttr("datapower_rbm_settings.test", "max_password_age", "30"),
			resource.TestCheckResourceAttr("datapower_rbm_settings.test", "do_password_history", "false"),
			resource.TestCheckResourceAttr("datapower_rbm_settings.test", "num_old_passwords", "5"),
			resource.TestCheckResourceAttr("datapower_rbm_settings.test", "cli_timeout", "0"),
			resource.TestCheckResourceAttr("datapower_rbm_settings.test", "max_failed_login", "0"),
			resource.TestCheckResourceAttr("datapower_rbm_settings.test", "lockout_period", "1"),
			resource.TestCheckResourceAttr("datapower_rbm_settings.test", "mc_force_dn_ldap_order", "false"),
			resource.TestCheckResourceAttr("datapower_rbm_settings.test", "password_hash_algorithm", "md5crypt"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
