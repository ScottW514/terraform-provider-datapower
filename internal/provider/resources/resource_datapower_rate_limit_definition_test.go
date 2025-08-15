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

func TestAccResourceRateLimitDefinition(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_RateLimitDefinition") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_RateLimitDefinition")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testutils.RateLimitDefinitionTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_rate_limit_definition.test", "type", "rate"),
			resource.TestCheckResourceAttr("datapower_rate_limit_definition.test", "interval", "1"),
			resource.TestCheckResourceAttr("datapower_rate_limit_definition.test", "unit", "minute"),
			resource.TestCheckResourceAttr("datapower_rate_limit_definition.test", "hard_limit", "true"),
			resource.TestCheckResourceAttr("datapower_rate_limit_definition.test", "is_client", "true"),
			resource.TestCheckResourceAttr("datapower_rate_limit_definition.test", "use_api_name", "false"),
			resource.TestCheckResourceAttr("datapower_rate_limit_definition.test", "use_app_id", "false"),
			resource.TestCheckResourceAttr("datapower_rate_limit_definition.test", "use_client_id", "false"),
			resource.TestCheckResourceAttr("datapower_rate_limit_definition.test", "auto_replenish", "true"),
			resource.TestCheckResourceAttr("datapower_rate_limit_definition.test", "weight", "1"),
			resource.TestCheckResourceAttr("datapower_rate_limit_definition.test", "response_headers", "true"),
			resource.TestCheckResourceAttr("datapower_rate_limit_definition.test", "emulate_burst_headers", "false"),
			resource.TestCheckResourceAttr("datapower_rate_limit_definition.test", "use_interval_offset", "true"),
			resource.TestCheckResourceAttr("datapower_rate_limit_definition.test", "allow_cache_fallback", "true"),
			resource.TestCheckResourceAttr("datapower_rate_limit_definition.test", "use_cache", "false"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
