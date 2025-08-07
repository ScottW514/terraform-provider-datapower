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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/testconfig"
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccResourceAPICollection(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_APICollection") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_APICollection")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testconfig.APICollectionTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_apicollection.test", "sandbox", "false"),
			resource.TestCheckResourceAttr("datapower_apicollection.test", "catalog_id", "default-catalog-id"),
			resource.TestCheckResourceAttr("datapower_apicollection.test", "catalog_name", "default"),
			resource.TestCheckResourceAttr("datapower_apicollection.test", "cache_capacity", "128"),
			resource.TestCheckResourceAttr("datapower_apicollection.test", "use_rate_limit_group", "false"),
			resource.TestCheckResourceAttr("datapower_apicollection.test", "enforce_pre_assembly_rate_limits", "true"),
			resource.TestCheckResourceAttr("datapower_apicollection.test", "api_processing_rule", "default-api-rule"),
			resource.TestCheckResourceAttr("datapower_apicollection.test", "api_error_rule", "default-api-error-rule"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
	actions.PostProcess()
}
