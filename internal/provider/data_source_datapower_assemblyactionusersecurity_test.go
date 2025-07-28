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

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models/testconfig"
)

func TestAccDataSourceAssemblyActionUserSecurity(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_AssemblyActionUserSecurity") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_AssemblyActionUserSecurity")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.AssemblyActionUserSecurityTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_assemblyactionusersecurity.test", "result.0.id", "AssemblyActionUserSecurity_name"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactionusersecurity.test", "result.0.factor_id", "default"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactionusersecurity.test", "result.0.extract_identity_method", "basic"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactionusersecurity.test", "result.0.ei_stop_on_error", "true"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactionusersecurity.test", "result.0.redirect_time_limit", "300"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactionusersecurity.test", "result.0.ei_default_form", "true"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactionusersecurity.test", "result.0.ei_custom_form_content_security_policy", "default-src 'self'"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactionusersecurity.test", "result.0.ei_form_time_limit", "300"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactionusersecurity.test", "result.0.user_auth_method", "user-registry"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactionusersecurity.test", "result.0.au_stop_on_error", "true"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactionusersecurity.test", "result.0.auth_response_headers_pattern", "(?i)x-api*"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactionusersecurity.test", "result.0.auth_response_credential_header", "X-API-Authenticated-Credential"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactionusersecurity.test", "result.0.user_az_method", "authenticated"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactionusersecurity.test", "result.0.az_stop_on_error", "true"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactionusersecurity.test", "result.0.az_default_form", "true"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactionusersecurity.test", "result.0.az_custom_form_content_security_policy", "default-src 'self'"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactionusersecurity.test", "result.0.az_form_time_limit", "300"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactionusersecurity.test", "result.0.az_table_display_checkboxes", "false"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactionusersecurity.test", "result.0.az_table_dynamic_entries", "user.default.az.dynamic_entries"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactionusersecurity.test", "result.0.action_debug", "false"),
				}...),
			},
		},
	})
}
