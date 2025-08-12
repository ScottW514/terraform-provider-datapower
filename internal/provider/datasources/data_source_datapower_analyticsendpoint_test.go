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

func TestAccDataSourceAnalyticsEndpoint(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_AnalyticsEndpoint") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_AnalyticsEndpoint")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testutils.AnalyticsEndpointTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_analyticsendpoint.test", "result.0.id", "AccTest_AnalyticsEndpoint"),
					resource.TestCheckResourceAttr("data.datapower_analyticsendpoint.test", "result.0.max_records", "1024"),
					resource.TestCheckResourceAttr("data.datapower_analyticsendpoint.test", "result.0.max_records_memory_kb", "512"),
					resource.TestCheckResourceAttr("data.datapower_analyticsendpoint.test", "result.0.max_delivery_memory_mb", "512"),
					resource.TestCheckResourceAttr("data.datapower_analyticsendpoint.test", "result.0.interval", "600"),
					resource.TestCheckResourceAttr("data.datapower_analyticsendpoint.test", "result.0.delivery_connections", "1"),
					resource.TestCheckResourceAttr("data.datapower_analyticsendpoint.test", "result.0.enable_jwt", "false"),
					resource.TestCheckResourceAttr("data.datapower_analyticsendpoint.test", "result.0.persistent_connection", "true"),
					resource.TestCheckResourceAttr("data.datapower_analyticsendpoint.test", "result.0.timeout", "90"),
					resource.TestCheckResourceAttr("data.datapower_analyticsendpoint.test", "result.0.persistent_timeout", "60"),
				}...),
			},
		},
	})
}
