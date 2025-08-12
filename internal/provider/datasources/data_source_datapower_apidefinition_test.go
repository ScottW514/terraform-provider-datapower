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

func TestAccDataSourceAPIDefinition(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_APIDefinition") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_APIDefinition")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testutils.APIDefinitionTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_apidefinition.test", "result.0.id", "AccTest_APIDefinition"),
					resource.TestCheckResourceAttr("data.datapower_apidefinition.test", "result.0.version", "1.0.0"),
					resource.TestCheckResourceAttr("data.datapower_apidefinition.test", "result.0.base_path", "/"),
					resource.TestCheckResourceAttr("data.datapower_apidefinition.test", "result.0.type", "standard"),
					resource.TestCheckResourceAttr("data.datapower_apidefinition.test", "result.0.require_api_mutual_tls", "false"),
					resource.TestCheckResourceAttr("data.datapower_apidefinition.test", "result.0.api_mutual_tls_header_name", "X-Client-Certificate"),
					resource.TestCheckResourceAttr("data.datapower_apidefinition.test", "result.0.cors_toggle", "false"),
					resource.TestCheckResourceAttr("data.datapower_apidefinition.test", "result.0.activity_log_toggle", "false"),
					resource.TestCheckResourceAttr("data.datapower_apidefinition.test", "result.0.content", "activity"),
					resource.TestCheckResourceAttr("data.datapower_apidefinition.test", "result.0.error_content", "payload"),
					resource.TestCheckResourceAttr("data.datapower_apidefinition.test", "result.0.message_buffering", "false"),
					resource.TestCheckResourceAttr("data.datapower_apidefinition.test", "result.0.deployment_state", "running"),
					resource.TestCheckResourceAttr("data.datapower_apidefinition.test", "result.0.share_rate_limit_count", "unset"),
					resource.TestCheckResourceAttr("data.datapower_apidefinition.test", "result.0.return_v5_responses", "false"),
					resource.TestCheckResourceAttr("data.datapower_apidefinition.test", "result.0.copy_id_headers_to_message", "false"),
					resource.TestCheckResourceAttr("data.datapower_apidefinition.test", "result.0.enforce_required_params", "true"),
					resource.TestCheckResourceAttr("data.datapower_apidefinition.test", "result.0.allow_chunked_uploads", "true"),
					resource.TestCheckResourceAttr("data.datapower_apidefinition.test", "result.0.set_v5_request_headers", "false"),
					resource.TestCheckResourceAttr("data.datapower_apidefinition.test", "result.0.get_raw_body_value", "false"),
					resource.TestCheckResourceAttr("data.datapower_apidefinition.test", "result.0.allow_trailing_slash", "false"),
					resource.TestCheckResourceAttr("data.datapower_apidefinition.test", "result.0.enforce_all_headers_case_insensitive", "false"),
					resource.TestCheckResourceAttr("data.datapower_apidefinition.test", "result.0.enforce_form_data_parameter", "true"),
					resource.TestCheckResourceAttr("data.datapower_apidefinition.test", "result.0.force_http500_for_soap11", "false"),
				}...),
			},
		},
	})
}
