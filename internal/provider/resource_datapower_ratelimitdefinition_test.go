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

func TestAccResourceRateLimitDefinition(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_RateLimitDefinition") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_RateLimitDefinition")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testconfig.RateLimitDefinitionTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_ratelimitdefinition.test", "type", "rate"),
			resource.TestCheckResourceAttr("datapower_ratelimitdefinition.test", "interval", "1"),
			resource.TestCheckResourceAttr("datapower_ratelimitdefinition.test", "unit", "minute"),
			resource.TestCheckResourceAttr("datapower_ratelimitdefinition.test", "hard_limit", "true"),
			resource.TestCheckResourceAttr("datapower_ratelimitdefinition.test", "is_client", "true"),
			resource.TestCheckResourceAttr("datapower_ratelimitdefinition.test", "use_api_name", "false"),
			resource.TestCheckResourceAttr("datapower_ratelimitdefinition.test", "use_app_id", "false"),
			resource.TestCheckResourceAttr("datapower_ratelimitdefinition.test", "use_client_id", "false"),
			resource.TestCheckResourceAttr("datapower_ratelimitdefinition.test", "auto_replenish", "true"),
			resource.TestCheckResourceAttr("datapower_ratelimitdefinition.test", "weight", "1"),
			resource.TestCheckResourceAttr("datapower_ratelimitdefinition.test", "response_headers", "true"),
			resource.TestCheckResourceAttr("datapower_ratelimitdefinition.test", "emulate_burst_headers", "false"),
			resource.TestCheckResourceAttr("datapower_ratelimitdefinition.test", "use_interval_offset", "true"),
			resource.TestCheckResourceAttr("datapower_ratelimitdefinition.test", "allow_cache_fallback", "true"),
			resource.TestCheckResourceAttr("datapower_ratelimitdefinition.test", "use_cache", "false"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
