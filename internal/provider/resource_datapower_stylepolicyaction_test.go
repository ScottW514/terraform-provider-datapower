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

func TestAccResourceStylePolicyAction(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_StylePolicyAction") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_StylePolicyAction")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testconfig.StylePolicyActionTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_stylepolicyaction.test", "type", "xform"),
			resource.TestCheckResourceAttr("datapower_stylepolicyaction.test", "parse_metrics_result_type", "none"),
			resource.TestCheckResourceAttr("datapower_stylepolicyaction.test", "input_language", "xml"),
			resource.TestCheckResourceAttr("datapower_stylepolicyaction.test", "transform_language", "none"),
			resource.TestCheckResourceAttr("datapower_stylepolicyaction.test", "output_language", "default"),
			resource.TestCheckResourceAttr("datapower_stylepolicyaction.test", "action_debug", "false"),
			resource.TestCheckResourceAttr("datapower_stylepolicyaction.test", "no_transcode_utf8", "false"),
			resource.TestCheckResourceAttr("datapower_stylepolicyaction.test", "named_in_out_location_type", "default"),
			resource.TestCheckResourceAttr("datapower_stylepolicyaction.test", "ssl_client_config_type", "client"),
			resource.TestCheckResourceAttr("datapower_stylepolicyaction.test", "transactional", "false"),
			resource.TestCheckResourceAttr("datapower_stylepolicyaction.test", "soap_validation", "body"),
			resource.TestCheckResourceAttr("datapower_stylepolicyaction.test", "sql_source_type", "static"),
			resource.TestCheckResourceAttr("datapower_stylepolicyaction.test", "jws_verify_strip_signature", "true"),
			resource.TestCheckResourceAttr("datapower_stylepolicyaction.test", "asynchronous", "false"),
			resource.TestCheckResourceAttr("datapower_stylepolicyaction.test", "results_mode", "first-available"),
			resource.TestCheckResourceAttr("datapower_stylepolicyaction.test", "retry_interval", "1000"),
			resource.TestCheckResourceAttr("datapower_stylepolicyaction.test", "multiple_outputs", "false"),
			resource.TestCheckResourceAttr("datapower_stylepolicyaction.test", "iterator_type", "XPATH"),
			resource.TestCheckResourceAttr("datapower_stylepolicyaction.test", "method_rewrite_type", "GET"),
			resource.TestCheckResourceAttr("datapower_stylepolicyaction.test", "method_type", "POST"),
			resource.TestCheckResourceAttr("datapower_stylepolicyaction.test", "method_type2", "POST"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
