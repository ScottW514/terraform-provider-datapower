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

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models/testconfig"
)

func TestAccDataSourceTimeSettings(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_TimeSettings") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_TimeSettings")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.TimeSettingsTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_timesettings.test", "enabled", "true"),
					resource.TestCheckResourceAttr("data.datapower_timesettings.test", "local_time_zone", "EST5EDT"),
					resource.TestCheckResourceAttr("data.datapower_timesettings.test", "custom_tz_name", "STD"),
					resource.TestCheckResourceAttr("data.datapower_timesettings.test", "utc_direction", "West"),
					resource.TestCheckResourceAttr("data.datapower_timesettings.test", "daylight_offset_hours", "1"),
					resource.TestCheckResourceAttr("data.datapower_timesettings.test", "tz_name_dst", "DST"),
					resource.TestCheckResourceAttr("data.datapower_timesettings.test", "daylight_start_month", "March"),
					resource.TestCheckResourceAttr("data.datapower_timesettings.test", "daylight_start_week", "2"),
					resource.TestCheckResourceAttr("data.datapower_timesettings.test", "daylight_start_day", "Sunday"),
					resource.TestCheckResourceAttr("data.datapower_timesettings.test", "daylight_start_time_hours", "2"),
					resource.TestCheckResourceAttr("data.datapower_timesettings.test", "daylight_stop_month", "November"),
					resource.TestCheckResourceAttr("data.datapower_timesettings.test", "daylight_stop_week", "1"),
					resource.TestCheckResourceAttr("data.datapower_timesettings.test", "daylight_stop_day", "Sunday"),
					resource.TestCheckResourceAttr("data.datapower_timesettings.test", "daylight_stop_time_hours", "2"),
				}...),
			},
		},
	})
}
