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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/testconfig"
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccDataSourceFormsLoginPolicy(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_FormsLoginPolicy") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_FormsLoginPolicy")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.FormsLoginPolicyTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_formsloginpolicy.test", "result.0.id", "AccTest_FormsLoginPolicy"),
					resource.TestCheckResourceAttr("data.datapower_formsloginpolicy.test", "result.0.login_form", "/LoginPage.htm"),
					resource.TestCheckResourceAttr("data.datapower_formsloginpolicy.test", "result.0.use_cookie_attributes", "false"),
					resource.TestCheckResourceAttr("data.datapower_formsloginpolicy.test", "result.0.use_ssl_for_login", "true"),
					resource.TestCheckResourceAttr("data.datapower_formsloginpolicy.test", "result.0.enable_migration", "false"),
					resource.TestCheckResourceAttr("data.datapower_formsloginpolicy.test", "result.0.ssl_port", "8080"),
					resource.TestCheckResourceAttr("data.datapower_formsloginpolicy.test", "result.0.error_page", "/ErrorPage.htm"),
					resource.TestCheckResourceAttr("data.datapower_formsloginpolicy.test", "result.0.logout_page", "/LogoutPage.htm"),
					resource.TestCheckResourceAttr("data.datapower_formsloginpolicy.test", "result.0.default_url", "/"),
					resource.TestCheckResourceAttr("data.datapower_formsloginpolicy.test", "result.0.forms_location", "backend"),
					resource.TestCheckResourceAttr("data.datapower_formsloginpolicy.test", "result.0.local_login_form", "store:///LoginPage.htm"),
					resource.TestCheckResourceAttr("data.datapower_formsloginpolicy.test", "result.0.local_error_page", "store:///ErrorPage.htm"),
					resource.TestCheckResourceAttr("data.datapower_formsloginpolicy.test", "result.0.local_logout_page", "store:///LogoutPage.htm"),
					resource.TestCheckResourceAttr("data.datapower_formsloginpolicy.test", "result.0.username_field", "j_username"),
					resource.TestCheckResourceAttr("data.datapower_formsloginpolicy.test", "result.0.password_field", "j_password"),
					resource.TestCheckResourceAttr("data.datapower_formsloginpolicy.test", "result.0.redirect_field", "originalUrl"),
					resource.TestCheckResourceAttr("data.datapower_formsloginpolicy.test", "result.0.form_processing_url", "/j_security_check"),
					resource.TestCheckResourceAttr("data.datapower_formsloginpolicy.test", "result.0.inactivity_timeout", "600"),
					resource.TestCheckResourceAttr("data.datapower_formsloginpolicy.test", "result.0.session_lifetime", "10800"),
					resource.TestCheckResourceAttr("data.datapower_formsloginpolicy.test", "result.0.redirect_url_type", "urlin"),
					resource.TestCheckResourceAttr("data.datapower_formsloginpolicy.test", "result.0.form_support_type", "static"),
					resource.TestCheckResourceAttr("data.datapower_formsloginpolicy.test", "result.0.form_support_script", "store:///Form-Generate-HTML.xsl"),
				}...),
			},
		},
	})
}
