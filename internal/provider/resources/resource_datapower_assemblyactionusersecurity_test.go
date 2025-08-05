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

func TestAccResourceAssemblyActionUserSecurity(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_AssemblyActionUserSecurity") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_AssemblyActionUserSecurity")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testconfig.AssemblyActionUserSecurityTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_assemblyactionusersecurity.test", "factor_id", "default"),
			resource.TestCheckResourceAttr("datapower_assemblyactionusersecurity.test", "extract_identity_method", "basic"),
			resource.TestCheckResourceAttr("datapower_assemblyactionusersecurity.test", "ei_stop_on_error", "true"),
			resource.TestCheckResourceAttr("datapower_assemblyactionusersecurity.test", "redirect_time_limit", "300"),
			resource.TestCheckResourceAttr("datapower_assemblyactionusersecurity.test", "ei_default_form", "true"),
			resource.TestCheckResourceAttr("datapower_assemblyactionusersecurity.test", "ei_custom_form_content_security_policy", "default-src 'self'"),
			resource.TestCheckResourceAttr("datapower_assemblyactionusersecurity.test", "ei_form_time_limit", "300"),
			resource.TestCheckResourceAttr("datapower_assemblyactionusersecurity.test", "user_auth_method", "user-registry"),
			resource.TestCheckResourceAttr("datapower_assemblyactionusersecurity.test", "au_stop_on_error", "true"),
			resource.TestCheckResourceAttr("datapower_assemblyactionusersecurity.test", "auth_response_headers_pattern", "(?i)x-api*"),
			resource.TestCheckResourceAttr("datapower_assemblyactionusersecurity.test", "auth_response_credential_header", "X-API-Authenticated-Credential"),
			resource.TestCheckResourceAttr("datapower_assemblyactionusersecurity.test", "user_az_method", "authenticated"),
			resource.TestCheckResourceAttr("datapower_assemblyactionusersecurity.test", "az_stop_on_error", "true"),
			resource.TestCheckResourceAttr("datapower_assemblyactionusersecurity.test", "az_default_form", "true"),
			resource.TestCheckResourceAttr("datapower_assemblyactionusersecurity.test", "az_custom_form_content_security_policy", "default-src 'self'"),
			resource.TestCheckResourceAttr("datapower_assemblyactionusersecurity.test", "az_form_time_limit", "300"),
			resource.TestCheckResourceAttr("datapower_assemblyactionusersecurity.test", "az_table_display_checkboxes", "false"),
			resource.TestCheckResourceAttr("datapower_assemblyactionusersecurity.test", "az_table_dynamic_entries", "user.default.az.dynamic_entries"),
			resource.TestCheckResourceAttr("datapower_assemblyactionusersecurity.test", "action_debug", "false"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
