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

func TestAccResourceThrottler(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Throttler") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Throttler")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testutils.ThrottlerTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_throttler.test", "enabled", "true"),
			resource.TestCheckResourceAttr("datapower_throttler.test", "throttle_at", "0"),
			resource.TestCheckResourceAttr("datapower_throttler.test", "terminate_at", "0"),
			resource.TestCheckResourceAttr("datapower_throttler.test", "temp_fs_throttle_at", "0"),
			resource.TestCheckResourceAttr("datapower_throttler.test", "temp_fs_terminate_at", "0"),
			resource.TestCheckResourceAttr("datapower_throttler.test", "qname_warn_at", "10"),
			resource.TestCheckResourceAttr("datapower_throttler.test", "timeout", "30"),
			resource.TestCheckResourceAttr("datapower_throttler.test", "statistics", "false"),
			resource.TestCheckResourceAttr("datapower_throttler.test", "log_level", "debug"),
			resource.TestCheckResourceAttr("datapower_throttler.test", "environmental_log", "true"),
			resource.TestCheckResourceAttr("datapower_throttler.test", "backlog_size", "0"),
			resource.TestCheckResourceAttr("datapower_throttler.test", "backlog_timeout", "30"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
