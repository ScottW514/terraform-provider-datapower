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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models/testconfig"
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccDataSourceLogTarget(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_LogTarget") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_LogTarget")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.LogTargetTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_logtarget.test", "result.0.id", "___LogTarget_name"),
					resource.TestCheckResourceAttr("data.datapower_logtarget.test", "result.0.type", "file"),
					resource.TestCheckResourceAttr("data.datapower_logtarget.test", "result.0.priority", "normal"),
					resource.TestCheckResourceAttr("data.datapower_logtarget.test", "result.0.soap_version", "soap11"),
					resource.TestCheckResourceAttr("data.datapower_logtarget.test", "result.0.format", "xml"),
					resource.TestCheckResourceAttr("data.datapower_logtarget.test", "result.0.timestamp_format", "zulu"),
					resource.TestCheckResourceAttr("data.datapower_logtarget.test", "result.0.fixed_format", "false"),
					resource.TestCheckResourceAttr("data.datapower_logtarget.test", "result.0.size", "500"),
					resource.TestCheckResourceAttr("data.datapower_logtarget.test", "result.0.archive_mode", "rotate"),
					resource.TestCheckResourceAttr("data.datapower_logtarget.test", "result.0.upload_method", "ftp"),
					resource.TestCheckResourceAttr("data.datapower_logtarget.test", "result.0.rotate", "3"),
					resource.TestCheckResourceAttr("data.datapower_logtarget.test", "result.0.use_ansi_color", "false"),
					resource.TestCheckResourceAttr("data.datapower_logtarget.test", "result.0.syslog_facility", "user"),
					resource.TestCheckResourceAttr("data.datapower_logtarget.test", "result.0.rate_limit", "100"),
					resource.TestCheckResourceAttr("data.datapower_logtarget.test", "result.0.max_connections", "1"),
					resource.TestCheckResourceAttr("data.datapower_logtarget.test", "result.0.connect_timeout", "60"),
					resource.TestCheckResourceAttr("data.datapower_logtarget.test", "result.0.idle_timeout", "15"),
					resource.TestCheckResourceAttr("data.datapower_logtarget.test", "result.0.active_timeout", "0"),
					resource.TestCheckResourceAttr("data.datapower_logtarget.test", "result.0.feedback_detection", "false"),
					resource.TestCheckResourceAttr("data.datapower_logtarget.test", "result.0.ssl_client_config_type", "client"),
					resource.TestCheckResourceAttr("data.datapower_logtarget.test", "result.0.retry_interval", "1"),
					resource.TestCheckResourceAttr("data.datapower_logtarget.test", "result.0.retry_attempts", "1"),
					resource.TestCheckResourceAttr("data.datapower_logtarget.test", "result.0.long_retry_interval", "20"),
					resource.TestCheckResourceAttr("data.datapower_logtarget.test", "result.0.log_precision", "second"),
					resource.TestCheckResourceAttr("data.datapower_logtarget.test", "result.0.event_buffer_size", "2048"),
				}...),
			},
		},
	})
}
