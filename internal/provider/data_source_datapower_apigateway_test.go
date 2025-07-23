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

func TestAccDataSourceAPIGateway(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_APIGateway") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_APIGateway")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.APIGatewayTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_apigateway.test", "result.0.id", "APIGateway_test"),
					resource.TestCheckResourceAttr("data.datapower_apigateway.test", "result.0.cache_memory_size", "2147483647"),
					resource.TestCheckResourceAttr("data.datapower_apigateway.test", "result.0.cache_size", "256"),
					resource.TestCheckResourceAttr("data.datapower_apigateway.test", "result.0.sha1_caching", "true"),
					resource.TestCheckResourceAttr("data.datapower_apigateway.test", "result.0.static_document_calls", "true"),
					resource.TestCheckResourceAttr("data.datapower_apigateway.test", "result.0.doc_cache_max_docs", "5000"),
					resource.TestCheckResourceAttr("data.datapower_apigateway.test", "result.0.doc_max_writes", "32768"),
					resource.TestCheckResourceAttr("data.datapower_apigateway.test", "result.0.share_rate_limit_count", "yes"),
					resource.TestCheckResourceAttr("data.datapower_apigateway.test", "result.0.front_timeout", "120"),
					resource.TestCheckResourceAttr("data.datapower_apigateway.test", "result.0.front_persistent_timeout", "180"),
				}...),
			},
		},
	})
}
