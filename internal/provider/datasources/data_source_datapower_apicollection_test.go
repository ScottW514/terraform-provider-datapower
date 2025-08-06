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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models/testconfig"
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccDataSourceAPICollection(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_APICollection") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_APICollection")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.APICollectionTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_apicollection.test", "result.0.id", "APICollection_name"),
					resource.TestCheckResourceAttr("data.datapower_apicollection.test", "result.0.sandbox", "false"),
					resource.TestCheckResourceAttr("data.datapower_apicollection.test", "result.0.catalog_id", "default-catalog-id"),
					resource.TestCheckResourceAttr("data.datapower_apicollection.test", "result.0.catalog_name", "default"),
					resource.TestCheckResourceAttr("data.datapower_apicollection.test", "result.0.cache_capacity", "128"),
					resource.TestCheckResourceAttr("data.datapower_apicollection.test", "result.0.use_rate_limit_group", "false"),
					resource.TestCheckResourceAttr("data.datapower_apicollection.test", "result.0.enforce_pre_assembly_rate_limits", "true"),
					resource.TestCheckResourceAttr("data.datapower_apicollection.test", "result.0.api_processing_rule", "default-api-rule"),
					resource.TestCheckResourceAttr("data.datapower_apicollection.test", "result.0.api_error_rule", "default-api-error-rule"),
				}...),
			},
		},
	})
	actions.PostProcess()
}
