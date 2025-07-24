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

func TestAccDataSourceRateLimitDefinition(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_RateLimitDefinition") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_RateLimitDefinition")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.RateLimitDefinitionTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_ratelimitdefinition.test", "result.0.id", "RateLimitDefinition_name"),
					resource.TestCheckResourceAttr("data.datapower_ratelimitdefinition.test", "result.0.type", "rate"),
					resource.TestCheckResourceAttr("data.datapower_ratelimitdefinition.test", "result.0.interval", "1"),
					resource.TestCheckResourceAttr("data.datapower_ratelimitdefinition.test", "result.0.unit", "minute"),
					resource.TestCheckResourceAttr("data.datapower_ratelimitdefinition.test", "result.0.hard_limit", "true"),
					resource.TestCheckResourceAttr("data.datapower_ratelimitdefinition.test", "result.0.is_client", "true"),
					resource.TestCheckResourceAttr("data.datapower_ratelimitdefinition.test", "result.0.use_api_name", "false"),
					resource.TestCheckResourceAttr("data.datapower_ratelimitdefinition.test", "result.0.use_app_id", "false"),
					resource.TestCheckResourceAttr("data.datapower_ratelimitdefinition.test", "result.0.use_client_id", "false"),
					resource.TestCheckResourceAttr("data.datapower_ratelimitdefinition.test", "result.0.auto_replenish", "true"),
					resource.TestCheckResourceAttr("data.datapower_ratelimitdefinition.test", "result.0.weight", "1"),
					resource.TestCheckResourceAttr("data.datapower_ratelimitdefinition.test", "result.0.response_headers", "true"),
					resource.TestCheckResourceAttr("data.datapower_ratelimitdefinition.test", "result.0.emulate_burst_headers", "false"),
					resource.TestCheckResourceAttr("data.datapower_ratelimitdefinition.test", "result.0.use_interval_offset", "true"),
					resource.TestCheckResourceAttr("data.datapower_ratelimitdefinition.test", "result.0.allow_cache_fallback", "true"),
					resource.TestCheckResourceAttr("data.datapower_ratelimitdefinition.test", "result.0.use_cache", "false"),
				}...),
			},
		},
	})
}
