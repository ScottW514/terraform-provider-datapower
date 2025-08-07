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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/testconfig"
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccDataSourceXMLManager(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_XMLManager") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_XMLManager")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.XMLManagerTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_xmlmanager.test", "result.0.id", "AccTest_XMLManager"),
					resource.TestCheckResourceAttr("data.datapower_xmlmanager.test", "result.0.cache_memory_size", "2147483647"),
					resource.TestCheckResourceAttr("data.datapower_xmlmanager.test", "result.0.cache_size", "256"),
					resource.TestCheckResourceAttr("data.datapower_xmlmanager.test", "result.0.sha1_caching", "true"),
					resource.TestCheckResourceAttr("data.datapower_xmlmanager.test", "result.0.static_document_calls", "true"),
					resource.TestCheckResourceAttr("data.datapower_xmlmanager.test", "result.0.search_results", "true"),
					resource.TestCheckResourceAttr("data.datapower_xmlmanager.test", "result.0.parser_limits_bytes_scanned", "4194304"),
					resource.TestCheckResourceAttr("data.datapower_xmlmanager.test", "result.0.parser_limits_element_depth", "512"),
					resource.TestCheckResourceAttr("data.datapower_xmlmanager.test", "result.0.parser_limits_attribute_count", "128"),
					resource.TestCheckResourceAttr("data.datapower_xmlmanager.test", "result.0.parser_limits_max_node_size", "33554432"),
					resource.TestCheckResourceAttr("data.datapower_xmlmanager.test", "result.0.parser_limits_external_references", "forbid"),
					resource.TestCheckResourceAttr("data.datapower_xmlmanager.test", "result.0.parser_limits_max_prefixes", "1024"),
					resource.TestCheckResourceAttr("data.datapower_xmlmanager.test", "result.0.parser_limits_max_namespaces", "1024"),
					resource.TestCheckResourceAttr("data.datapower_xmlmanager.test", "result.0.parser_limits_max_local_names", "60000"),
					resource.TestCheckResourceAttr("data.datapower_xmlmanager.test", "result.0.doc_cache_max_docs", "5000"),
					resource.TestCheckResourceAttr("data.datapower_xmlmanager.test", "result.0.doc_max_writes", "32768"),
					resource.TestCheckResourceAttr("data.datapower_xmlmanager.test", "result.0.user_agent", "default"),
				}...),
			},
		},
	})
}
