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

func TestAccResourceXMLManager(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_XMLManager") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_XMLManager")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testutils.XMLManagerTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_xml_manager.test", "cache_memory_size", "2147483647"),
			resource.TestCheckResourceAttr("datapower_xml_manager.test", "cache_size", "256"),
			resource.TestCheckResourceAttr("datapower_xml_manager.test", "sha1_caching", "true"),
			resource.TestCheckResourceAttr("datapower_xml_manager.test", "static_document_calls", "true"),
			resource.TestCheckResourceAttr("datapower_xml_manager.test", "search_results", "true"),
			resource.TestCheckResourceAttr("datapower_xml_manager.test", "parser_limits_bytes_scanned", "4194304"),
			resource.TestCheckResourceAttr("datapower_xml_manager.test", "parser_limits_element_depth", "512"),
			resource.TestCheckResourceAttr("datapower_xml_manager.test", "parser_limits_attribute_count", "128"),
			resource.TestCheckResourceAttr("datapower_xml_manager.test", "parser_limits_max_node_size", "33554432"),
			resource.TestCheckResourceAttr("datapower_xml_manager.test", "parser_limits_external_references", "forbid"),
			resource.TestCheckResourceAttr("datapower_xml_manager.test", "parser_limits_max_prefixes", "1024"),
			resource.TestCheckResourceAttr("datapower_xml_manager.test", "parser_limits_max_namespaces", "1024"),
			resource.TestCheckResourceAttr("datapower_xml_manager.test", "parser_limits_max_local_names", "60000"),
			resource.TestCheckResourceAttr("datapower_xml_manager.test", "doc_cache_max_docs", "5000"),
			resource.TestCheckResourceAttr("datapower_xml_manager.test", "doc_max_writes", "32768"),
			resource.TestCheckResourceAttr("datapower_xml_manager.test", "user_agent", "default"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
