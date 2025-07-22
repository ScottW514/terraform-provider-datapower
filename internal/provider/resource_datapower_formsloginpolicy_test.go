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

func TestAccResourceFormsLoginPolicy(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_FormsLoginPolicy") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_FormsLoginPolicy")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testconfig.FormsLoginPolicyTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_formsloginpolicy.test", "login_form", "/LoginPage.htm"),
			resource.TestCheckResourceAttr("datapower_formsloginpolicy.test", "use_cookie_attributes", "no"),
			resource.TestCheckResourceAttr("datapower_formsloginpolicy.test", "use_ssl_for_login", "true"),
			resource.TestCheckResourceAttr("datapower_formsloginpolicy.test", "enable_migration", "false"),
			resource.TestCheckResourceAttr("datapower_formsloginpolicy.test", "ssl_port", "8080"),
			resource.TestCheckResourceAttr("datapower_formsloginpolicy.test", "error_page", "/ErrorPage.htm"),
			resource.TestCheckResourceAttr("datapower_formsloginpolicy.test", "logout_page", "/LogoutPage.htm"),
			resource.TestCheckResourceAttr("datapower_formsloginpolicy.test", "default_url", "/"),
			resource.TestCheckResourceAttr("datapower_formsloginpolicy.test", "forms_location", "backend"),
			resource.TestCheckResourceAttr("datapower_formsloginpolicy.test", "local_login_form", "store:///LoginPage.htm"),
			resource.TestCheckResourceAttr("datapower_formsloginpolicy.test", "local_error_page", "store:///ErrorPage.htm"),
			resource.TestCheckResourceAttr("datapower_formsloginpolicy.test", "local_logout_page", "store:///LogoutPage.htm"),
			resource.TestCheckResourceAttr("datapower_formsloginpolicy.test", "username_field", "j_username"),
			resource.TestCheckResourceAttr("datapower_formsloginpolicy.test", "password_field", "j_password"),
			resource.TestCheckResourceAttr("datapower_formsloginpolicy.test", "redirect_field", "originalUrl"),
			resource.TestCheckResourceAttr("datapower_formsloginpolicy.test", "form_processing_url", "/j_security_check"),
			resource.TestCheckResourceAttr("datapower_formsloginpolicy.test", "inactivity_timeout", "600"),
			resource.TestCheckResourceAttr("datapower_formsloginpolicy.test", "session_lifetime", "10800"),
			resource.TestCheckResourceAttr("datapower_formsloginpolicy.test", "redirect_url_type", "urlin"),
			resource.TestCheckResourceAttr("datapower_formsloginpolicy.test", "form_support_type", "static"),
			resource.TestCheckResourceAttr("datapower_formsloginpolicy.test", "form_support_script", "store:///Form-Generate-HTML.xsl"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
