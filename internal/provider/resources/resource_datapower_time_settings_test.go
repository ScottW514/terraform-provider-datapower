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

func TestAccResourceTimeSettings(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_TimeSettings") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_TimeSettings")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testutils.TimeSettingsTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_time_settings.test", "enabled", "true"),
			resource.TestCheckResourceAttr("datapower_time_settings.test", "local_time_zone", "EST5EDT"),
			resource.TestCheckResourceAttr("datapower_time_settings.test", "custom_tz_name", "STD"),
			resource.TestCheckResourceAttr("datapower_time_settings.test", "utc_direction", "West"),
			resource.TestCheckResourceAttr("datapower_time_settings.test", "daylight_offset_hours", "1"),
			resource.TestCheckResourceAttr("datapower_time_settings.test", "tz_name_dst", "DST"),
			resource.TestCheckResourceAttr("datapower_time_settings.test", "daylight_start_month", "March"),
			resource.TestCheckResourceAttr("datapower_time_settings.test", "daylight_start_week", "2"),
			resource.TestCheckResourceAttr("datapower_time_settings.test", "daylight_start_day", "Sunday"),
			resource.TestCheckResourceAttr("datapower_time_settings.test", "daylight_start_time_hours", "2"),
			resource.TestCheckResourceAttr("datapower_time_settings.test", "daylight_stop_month", "November"),
			resource.TestCheckResourceAttr("datapower_time_settings.test", "daylight_stop_week", "1"),
			resource.TestCheckResourceAttr("datapower_time_settings.test", "daylight_stop_day", "Sunday"),
			resource.TestCheckResourceAttr("datapower_time_settings.test", "daylight_stop_time_hours", "2"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
