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

func TestAccDataSourceErrorReportSettings(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_ErrorReportSettings") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_ErrorReportSettings")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.ErrorReportSettingsTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_errorreportsettings.test", "enabled", "true"),
					resource.TestCheckResourceAttr("data.datapower_errorreportsettings.test", "upload_report", "false"),
					resource.TestCheckResourceAttr("data.datapower_errorreportsettings.test", "use_smtp", "false"),
					resource.TestCheckResourceAttr("data.datapower_errorreportsettings.test", "internal_state", "false"),
					resource.TestCheckResourceAttr("data.datapower_errorreportsettings.test", "ffdc_packet_capture", "false"),
					resource.TestCheckResourceAttr("data.datapower_errorreportsettings.test", "ffdc_event_log_capture", "false"),
					resource.TestCheckResourceAttr("data.datapower_errorreportsettings.test", "ffdc_memory_leak_capture", "false"),
					resource.TestCheckResourceAttr("data.datapower_errorreportsettings.test", "always_on_startup", "true"),
					resource.TestCheckResourceAttr("data.datapower_errorreportsettings.test", "always_on_shutdown", "true"),
					resource.TestCheckResourceAttr("data.datapower_errorreportsettings.test", "report_history_kept", "5"),
				}...),
			},
		},
	})
	actions.PostProcess()
}
