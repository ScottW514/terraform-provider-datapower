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

func TestAccDataSourceB2BGateway(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_B2BGateway") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_B2BGateway")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testutils.B2BGatewayTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_b2b_gateway.test", "result.0.id", "AccTest_B2BGateway"),
					resource.TestCheckResourceAttr("data.datapower_b2b_gateway.test", "result.0.priority", "normal"),
					resource.TestCheckResourceAttr("data.datapower_b2b_gateway.test", "result.0.doc_store_location", "(default)"),
					resource.TestCheckResourceAttr("data.datapower_b2b_gateway.test", "result.0.document_routing_preprocessor_type", "stylesheet"),
					resource.TestCheckResourceAttr("data.datapower_b2b_gateway.test", "result.0.document_routing_preprocessor", "store:///b2b-routing.xsl"),
					resource.TestCheckResourceAttr("data.datapower_b2b_gateway.test", "result.0.document_routing_preprocessor_debug", "false"),
					resource.TestCheckResourceAttr("data.datapower_b2b_gateway.test", "result.0.archive_mode", "purgeonly"),
					resource.TestCheckResourceAttr("data.datapower_b2b_gateway.test", "result.0.archive_minimum_size", "1024"),
					resource.TestCheckResourceAttr("data.datapower_b2b_gateway.test", "result.0.archive_document_age", "90"),
					resource.TestCheckResourceAttr("data.datapower_b2b_gateway.test", "result.0.archive_minimum_documents", "100"),
					resource.TestCheckResourceAttr("data.datapower_b2b_gateway.test", "result.0.disk_use_check_interval", "60"),
					resource.TestCheckResourceAttr("data.datapower_b2b_gateway.test", "result.0.max_document_disk_use", "25165824"),
					resource.TestCheckResourceAttr("data.datapower_b2b_gateway.test", "result.0.archive_monitor", "true"),
					resource.TestCheckResourceAttr("data.datapower_b2b_gateway.test", "result.0.shaping_threshold", "200"),
					resource.TestCheckResourceAttr("data.datapower_b2b_gateway.test", "result.0.xml_manager", "default"),
					resource.TestCheckResourceAttr("data.datapower_b2b_gateway.test", "result.0.debug_mode", "off"),
					resource.TestCheckResourceAttr("data.datapower_b2b_gateway.test", "result.0.debug_history", "25"),
					resource.TestCheckResourceAttr("data.datapower_b2b_gateway.test", "result.0.front_side_timeout", "120"),
				}...),
			},
		},
	})
}
