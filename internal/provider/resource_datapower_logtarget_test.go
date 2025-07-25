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

func TestAccResourceLogTarget(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_LogTarget") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_LogTarget")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testconfig.LogTargetTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_logtarget.test", "type", "file"),
			resource.TestCheckResourceAttr("datapower_logtarget.test", "priority", "normal"),
			resource.TestCheckResourceAttr("datapower_logtarget.test", "soap_version", "soap11"),
			resource.TestCheckResourceAttr("datapower_logtarget.test", "format", "xml"),
			resource.TestCheckResourceAttr("datapower_logtarget.test", "timestamp_format", "zulu"),
			resource.TestCheckResourceAttr("datapower_logtarget.test", "fixed_format", "false"),
			resource.TestCheckResourceAttr("datapower_logtarget.test", "size", "500"),
			resource.TestCheckResourceAttr("datapower_logtarget.test", "archive_mode", "rotate"),
			resource.TestCheckResourceAttr("datapower_logtarget.test", "upload_method", "ftp"),
			resource.TestCheckResourceAttr("datapower_logtarget.test", "rotate", "3"),
			resource.TestCheckResourceAttr("datapower_logtarget.test", "use_ansi_color", "false"),
			resource.TestCheckResourceAttr("datapower_logtarget.test", "syslog_facility", "user"),
			resource.TestCheckResourceAttr("datapower_logtarget.test", "rate_limit", "100"),
			resource.TestCheckResourceAttr("datapower_logtarget.test", "max_connections", "1"),
			resource.TestCheckResourceAttr("datapower_logtarget.test", "connect_timeout", "60"),
			resource.TestCheckResourceAttr("datapower_logtarget.test", "idle_timeout", "15"),
			resource.TestCheckResourceAttr("datapower_logtarget.test", "active_timeout", "0"),
			resource.TestCheckResourceAttr("datapower_logtarget.test", "feedback_detection", "false"),
			resource.TestCheckResourceAttr("datapower_logtarget.test", "ssl_client_config_type", "client"),
			resource.TestCheckResourceAttr("datapower_logtarget.test", "retry_interval", "1"),
			resource.TestCheckResourceAttr("datapower_logtarget.test", "retry_attempts", "1"),
			resource.TestCheckResourceAttr("datapower_logtarget.test", "long_retry_interval", "20"),
			resource.TestCheckResourceAttr("datapower_logtarget.test", "log_precision", "second"),
			resource.TestCheckResourceAttr("datapower_logtarget.test", "event_buffer_size", "2048"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
