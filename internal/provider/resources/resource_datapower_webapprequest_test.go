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

func TestAccResourceWebAppRequest(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_WebAppRequest") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_WebAppRequest")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testutils.WebAppRequestTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_webapprequest.test", "policy_type", "admission"),
			resource.TestCheckResourceAttr("datapower_webapprequest.test", "ssl_policy", "allow"),
			resource.TestCheckResourceAttr("datapower_webapprequest.test", "max_body_size", "128000000"),
			resource.TestCheckResourceAttr("datapower_webapprequest.test", "xml_policy", "nothing"),
			resource.TestCheckResourceAttr("datapower_webapprequest.test", "non_xml_policy", "nothing"),
			resource.TestCheckResourceAttr("datapower_webapprequest.test", "query_string_policy", "allow"),
			resource.TestCheckResourceAttr("datapower_webapprequest.test", "sql_injection", "false"),
			resource.TestCheckResourceAttr("datapower_webapprequest.test", "max_uri_size", "1024"),
			resource.TestCheckResourceAttr("datapower_webapprequest.test", "uri_filter_unicode", "true"),
			resource.TestCheckResourceAttr("datapower_webapprequest.test", "uri_filter_dot_dot", "true"),
			resource.TestCheckResourceAttr("datapower_webapprequest.test", "uri_filter_exe", "true"),
			resource.TestCheckResourceAttr("datapower_webapprequest.test", "uri_filter_fragment", "truncate"),
			resource.TestCheckResourceAttr("datapower_webapprequest.test", "process_all_cookie", "true"),
			resource.TestCheckResourceAttr("datapower_webapprequest.test", "sql_injection_patterns_file", "store:///SQL-Injection-Patterns.xml"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
