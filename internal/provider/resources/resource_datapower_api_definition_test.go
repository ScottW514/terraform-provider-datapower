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

func TestAccResourceAPIDefinition(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_APIDefinition") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_APIDefinition")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testutils.APIDefinitionTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_api_definition.test", "version", "1.0.0"),
			resource.TestCheckResourceAttr("datapower_api_definition.test", "base_path", "/"),
			resource.TestCheckResourceAttr("datapower_api_definition.test", "type", "standard"),
			resource.TestCheckResourceAttr("datapower_api_definition.test", "require_api_mutual_tls", "false"),
			resource.TestCheckResourceAttr("datapower_api_definition.test", "api_mutual_tls_header_name", "X-Client-Certificate"),
			resource.TestCheckResourceAttr("datapower_api_definition.test", "cors_toggle", "false"),
			resource.TestCheckResourceAttr("datapower_api_definition.test", "activity_log_toggle", "false"),
			resource.TestCheckResourceAttr("datapower_api_definition.test", "content", "activity"),
			resource.TestCheckResourceAttr("datapower_api_definition.test", "error_content", "payload"),
			resource.TestCheckResourceAttr("datapower_api_definition.test", "message_buffering", "false"),
			resource.TestCheckResourceAttr("datapower_api_definition.test", "deployment_state", "running"),
			resource.TestCheckResourceAttr("datapower_api_definition.test", "share_rate_limit_count", "unset"),
			resource.TestCheckResourceAttr("datapower_api_definition.test", "return_v5responses", "false"),
			resource.TestCheckResourceAttr("datapower_api_definition.test", "copy_id_headers_to_message", "false"),
			resource.TestCheckResourceAttr("datapower_api_definition.test", "enforce_required_params", "true"),
			resource.TestCheckResourceAttr("datapower_api_definition.test", "allow_chunked_uploads", "true"),
			resource.TestCheckResourceAttr("datapower_api_definition.test", "set_v5request_headers", "false"),
			resource.TestCheckResourceAttr("datapower_api_definition.test", "get_raw_body_value", "false"),
			resource.TestCheckResourceAttr("datapower_api_definition.test", "allow_trailing_slash", "false"),
			resource.TestCheckResourceAttr("datapower_api_definition.test", "enforce_all_headers_case_insensitive", "false"),
			resource.TestCheckResourceAttr("datapower_api_definition.test", "enforce_form_data_parameter", "true"),
			resource.TestCheckResourceAttr("datapower_api_definition.test", "force_http500_for_soap11", "false"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
