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

func TestAccResourceB2BGateway(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_B2BGateway") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_B2BGateway")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testutils.B2BGatewayTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_b2bgateway.test", "priority", "normal"),
			resource.TestCheckResourceAttr("datapower_b2bgateway.test", "doc_store_location", "(default)"),
			resource.TestCheckResourceAttr("datapower_b2bgateway.test", "document_routing_preprocessor_type", "stylesheet"),
			resource.TestCheckResourceAttr("datapower_b2bgateway.test", "document_routing_preprocessor", "store:///b2b-routing.xsl"),
			resource.TestCheckResourceAttr("datapower_b2bgateway.test", "document_routing_preprocessor_debug", "false"),
			resource.TestCheckResourceAttr("datapower_b2bgateway.test", "archive_mode", "archpurge"),
			resource.TestCheckResourceAttr("datapower_b2bgateway.test", "archive_minimum_size", "1024"),
			resource.TestCheckResourceAttr("datapower_b2bgateway.test", "archive_document_age", "90"),
			resource.TestCheckResourceAttr("datapower_b2bgateway.test", "archive_minimum_documents", "100"),
			resource.TestCheckResourceAttr("datapower_b2bgateway.test", "disk_use_check_interval", "60"),
			resource.TestCheckResourceAttr("datapower_b2bgateway.test", "max_document_disk_use", "25165824"),
			resource.TestCheckResourceAttr("datapower_b2bgateway.test", "archive_monitor", "true"),
			resource.TestCheckResourceAttr("datapower_b2bgateway.test", "shaping_threshold", "200"),
			resource.TestCheckResourceAttr("datapower_b2bgateway.test", "xml_manager", "default"),
			resource.TestCheckResourceAttr("datapower_b2bgateway.test", "debug_mode", "off"),
			resource.TestCheckResourceAttr("datapower_b2bgateway.test", "debug_history", "25"),
			resource.TestCheckResourceAttr("datapower_b2bgateway.test", "front_side_timeout", "120"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
