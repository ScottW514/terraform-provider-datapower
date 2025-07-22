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

func TestAccDataSourceStylePolicyAction(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_StylePolicyAction") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_StylePolicyAction")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.StylePolicyActionTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_stylepolicyaction.test", "result.0.id", "___StylePolicyAction_test"),
					resource.TestCheckResourceAttr("data.datapower_stylepolicyaction.test", "result.0.type", "xform"),
					resource.TestCheckResourceAttr("data.datapower_stylepolicyaction.test", "result.0.parse_metrics_result_type", "none"),
					resource.TestCheckResourceAttr("data.datapower_stylepolicyaction.test", "result.0.input_language", "xml"),
					resource.TestCheckResourceAttr("data.datapower_stylepolicyaction.test", "result.0.transform_language", "none"),
					resource.TestCheckResourceAttr("data.datapower_stylepolicyaction.test", "result.0.output_language", "default"),
					resource.TestCheckResourceAttr("data.datapower_stylepolicyaction.test", "result.0.named_in_out_location_type", "default"),
					resource.TestCheckResourceAttr("data.datapower_stylepolicyaction.test", "result.0.ssl_client_config_type", "client"),
					resource.TestCheckResourceAttr("data.datapower_stylepolicyaction.test", "result.0.soap_validation", "body"),
					resource.TestCheckResourceAttr("data.datapower_stylepolicyaction.test", "result.0.sql_source_type", "static"),
					resource.TestCheckResourceAttr("data.datapower_stylepolicyaction.test", "result.0.results_mode", "first-available"),
					resource.TestCheckResourceAttr("data.datapower_stylepolicyaction.test", "result.0.iterator_type", "XPATH"),
					resource.TestCheckResourceAttr("data.datapower_stylepolicyaction.test", "result.0.method_rewrite_type", "GET"),
					resource.TestCheckResourceAttr("data.datapower_stylepolicyaction.test", "result.0.method_type", "POST"),
					resource.TestCheckResourceAttr("data.datapower_stylepolicyaction.test", "result.0.method_type2", "POST"),
				}...),
			},
		},
	})
}
