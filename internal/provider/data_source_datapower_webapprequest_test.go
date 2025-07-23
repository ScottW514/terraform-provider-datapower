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

func TestAccDataSourceWebAppRequest(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_WebAppRequest") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_WebAppRequest")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.WebAppRequestTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_webapprequest.test", "result.0.id", "WebAppRequest_test"),
					resource.TestCheckResourceAttr("data.datapower_webapprequest.test", "result.0.policy_type", "admission"),
					resource.TestCheckResourceAttr("data.datapower_webapprequest.test", "result.0.ssl_policy", "allow"),
					resource.TestCheckResourceAttr("data.datapower_webapprequest.test", "result.0.max_body_size", "128000000"),
					resource.TestCheckResourceAttr("data.datapower_webapprequest.test", "result.0.xml_policy", "nothing"),
					resource.TestCheckResourceAttr("data.datapower_webapprequest.test", "result.0.non_xml_policy", "nothing"),
					resource.TestCheckResourceAttr("data.datapower_webapprequest.test", "result.0.query_string_policy", "allow"),
					resource.TestCheckResourceAttr("data.datapower_webapprequest.test", "result.0.sql_injection", "false"),
					resource.TestCheckResourceAttr("data.datapower_webapprequest.test", "result.0.max_uri_size", "1024"),
					resource.TestCheckResourceAttr("data.datapower_webapprequest.test", "result.0.uri_filter_unicode", "true"),
					resource.TestCheckResourceAttr("data.datapower_webapprequest.test", "result.0.uri_filter_dot_dot", "true"),
					resource.TestCheckResourceAttr("data.datapower_webapprequest.test", "result.0.uri_filter_exe", "true"),
					resource.TestCheckResourceAttr("data.datapower_webapprequest.test", "result.0.uri_filter_fragment", "truncate"),
					resource.TestCheckResourceAttr("data.datapower_webapprequest.test", "result.0.process_all_cookie", "true"),
					resource.TestCheckResourceAttr("data.datapower_webapprequest.test", "result.0.sql_injection_patterns_file", "store:///SQL-Injection-Patterns.xml"),
				}...),
			},
		},
	})
}
