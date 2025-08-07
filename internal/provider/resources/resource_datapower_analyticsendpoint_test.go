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

func TestAccResourceAnalyticsEndpoint(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_AnalyticsEndpoint") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_AnalyticsEndpoint")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testconfig.AnalyticsEndpointTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_analyticsendpoint.test", "max_records", "1024"),
			resource.TestCheckResourceAttr("datapower_analyticsendpoint.test", "max_records_memory_kb", "512"),
			resource.TestCheckResourceAttr("datapower_analyticsendpoint.test", "max_delivery_memory_mb", "512"),
			resource.TestCheckResourceAttr("datapower_analyticsendpoint.test", "interval", "600"),
			resource.TestCheckResourceAttr("datapower_analyticsendpoint.test", "delivery_connections", "1"),
			resource.TestCheckResourceAttr("datapower_analyticsendpoint.test", "enable_jwt", "false"),
			resource.TestCheckResourceAttr("datapower_analyticsendpoint.test", "persistent_connection", "true"),
			resource.TestCheckResourceAttr("datapower_analyticsendpoint.test", "timeout", "90"),
			resource.TestCheckResourceAttr("datapower_analyticsendpoint.test", "persistent_timeout", "60"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
	actions.PostProcess()
}
