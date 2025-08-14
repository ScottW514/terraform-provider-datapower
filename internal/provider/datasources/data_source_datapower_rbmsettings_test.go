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

func TestAccDataSourceRBMSettings(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_RBMSettings") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_RBMSettings")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testutils.RBMSettingsTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "enabled", "true"),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "au_method", "local"),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "auoidc_scope", "openid"),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "auoidc_key_fetch_interval", "30"),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "auldap_search_for_dn", "false"),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "auldap_prefix", "cn="),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "au_force_dnldap_order", "false"),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "au_cache_allow", "absolute"),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "au_cache_ttl", "600"),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "auldap_read_timeout", "60"),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "mc_method", "local"),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "mcldap_search_for_group", "false"),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "mcldap_read_timeout", "60"),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "fallback_login", "disabled"),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "apply_to_cli", "false"),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "restrict_admin_to_serial_port", "false"),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "min_password_length", "6"),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "require_mixed_case", "false"),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "require_digit", "false"),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "require_non_alpha_numeric", "false"),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "disallow_username_substring", "false"),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "do_password_aging", "false"),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "max_password_age", "30"),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "do_password_history", "false"),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "num_old_passwords", "5"),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "cli_timeout", "0"),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "max_failed_login", "0"),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "lockout_period", "1"),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "mc_force_dnldap_order", "false"),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "password_hash_algorithm", "md5crypt"),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "ldapssl_client_config_type", "proxy"),
					resource.TestCheckResourceAttr("data.datapower_rbmsettings.test", "mcldapssl_client_config_type", "proxy"),
				}...),
			},
		},
	})
}
