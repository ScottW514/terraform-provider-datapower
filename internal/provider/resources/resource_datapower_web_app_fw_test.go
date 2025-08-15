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

func TestAccResourceWebAppFW(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_WebAppFW") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_WebAppFW")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testutils.WebAppFWTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_web_app_fw.test", "priority", "normal"),
			resource.TestCheckResourceAttr("datapower_web_app_fw.test", "remote_port", "80"),
			resource.TestCheckResourceAttr("datapower_web_app_fw.test", "xml_manager", "default"),
			resource.TestCheckResourceAttr("datapower_web_app_fw.test", "uri_normalization", "true"),
			resource.TestCheckResourceAttr("datapower_web_app_fw.test", "rewrite_errors", "true"),
			resource.TestCheckResourceAttr("datapower_web_app_fw.test", "delay_errors", "true"),
			resource.TestCheckResourceAttr("datapower_web_app_fw.test", "delay_errors_duration", "1000"),
			resource.TestCheckResourceAttr("datapower_web_app_fw.test", "stream_output_to_back", "buffer-until-verification"),
			resource.TestCheckResourceAttr("datapower_web_app_fw.test", "stream_output_to_front", "buffer-until-verification"),
			resource.TestCheckResourceAttr("datapower_web_app_fw.test", "front_timeout", "120"),
			resource.TestCheckResourceAttr("datapower_web_app_fw.test", "back_timeout", "120"),
			resource.TestCheckResourceAttr("datapower_web_app_fw.test", "front_persistent_timeout", "180"),
			resource.TestCheckResourceAttr("datapower_web_app_fw.test", "allow_cache_control_header", "false"),
			resource.TestCheckResourceAttr("datapower_web_app_fw.test", "back_persistent_timeout", "180"),
			resource.TestCheckResourceAttr("datapower_web_app_fw.test", "front_http_version", "HTTP/1.1"),
			resource.TestCheckResourceAttr("datapower_web_app_fw.test", "back_http_version", "HTTP/1.1"),
			resource.TestCheckResourceAttr("datapower_web_app_fw.test", "request_side_security", "true"),
			resource.TestCheckResourceAttr("datapower_web_app_fw.test", "response_side_security", "true"),
			resource.TestCheckResourceAttr("datapower_web_app_fw.test", "do_chunked_upload", "false"),
			resource.TestCheckResourceAttr("datapower_web_app_fw.test", "follow_redirects", "true"),
			resource.TestCheckResourceAttr("datapower_web_app_fw.test", "http_client_ip_label", "X-Client-IP"),
			resource.TestCheckResourceAttr("datapower_web_app_fw.test", "http_log_cor_id_label", "X-Global-Transaction-ID"),
			resource.TestCheckResourceAttr("datapower_web_app_fw.test", "debug_mode", "off"),
			resource.TestCheckResourceAttr("datapower_web_app_fw.test", "debug_history", "25"),
			resource.TestCheckResourceAttr("datapower_web_app_fw.test", "do_host_rewriting", "true"),
			resource.TestCheckResourceAttr("datapower_web_app_fw.test", "ssl_config_type", "server"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
