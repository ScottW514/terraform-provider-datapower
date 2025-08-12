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
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccDataSourceXSLProxyService(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_XSLProxyService") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_XSLProxyService")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testutils.XSLProxyServiceTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_xslproxyservice.test", "result.0.id", "AccTest_XSLProxyService"),
					resource.TestCheckResourceAttr("data.datapower_xslproxyservice.test", "result.0.type", "static-backend"),
					resource.TestCheckResourceAttr("data.datapower_xslproxyservice.test", "result.0.xml_manager", "default"),
					resource.TestCheckResourceAttr("data.datapower_xslproxyservice.test", "result.0.style_policy", "default"),
					resource.TestCheckResourceAttr("data.datapower_xslproxyservice.test", "result.0.credential_charset", "protocol"),
					resource.TestCheckResourceAttr("data.datapower_xslproxyservice.test", "result.0.ssl_config_type", "server"),
					resource.TestCheckResourceAttr("data.datapower_xslproxyservice.test", "result.0.priority", "normal"),
					resource.TestCheckResourceAttr("data.datapower_xslproxyservice.test", "result.0.http_timeout", "120"),
					resource.TestCheckResourceAttr("data.datapower_xslproxyservice.test", "result.0.http_persist_timeout", "180"),
					resource.TestCheckResourceAttr("data.datapower_xslproxyservice.test", "result.0.do_host_rewrite", "true"),
					resource.TestCheckResourceAttr("data.datapower_xslproxyservice.test", "result.0.suppress_http_warnings", "false"),
					resource.TestCheckResourceAttr("data.datapower_xslproxyservice.test", "result.0.http_compression", "false"),
					resource.TestCheckResourceAttr("data.datapower_xslproxyservice.test", "result.0.http_include_response_type_encoding", "false"),
					resource.TestCheckResourceAttr("data.datapower_xslproxyservice.test", "result.0.always_show_errors", "false"),
					resource.TestCheckResourceAttr("data.datapower_xslproxyservice.test", "result.0.disallow_get", "false"),
					resource.TestCheckResourceAttr("data.datapower_xslproxyservice.test", "result.0.disallow_empty_response", "false"),
					resource.TestCheckResourceAttr("data.datapower_xslproxyservice.test", "result.0.http_persistent_connections", "true"),
					resource.TestCheckResourceAttr("data.datapower_xslproxyservice.test", "result.0.http_client_ip_label", "X-Client-IP"),
					resource.TestCheckResourceAttr("data.datapower_xslproxyservice.test", "result.0.http_log_cor_id_label", "X-Global-Transaction-ID"),
					resource.TestCheckResourceAttr("data.datapower_xslproxyservice.test", "result.0.http_proxy_port", "800"),
					resource.TestCheckResourceAttr("data.datapower_xslproxyservice.test", "result.0.do_chunked_upload", "false"),
					resource.TestCheckResourceAttr("data.datapower_xslproxyservice.test", "result.0.default_param_namespace", "http://www.datapower.com/param/config"),
					resource.TestCheckResourceAttr("data.datapower_xslproxyservice.test", "result.0.query_param_namespace", "http://www.datapower.com/param/query"),
					resource.TestCheckResourceAttr("data.datapower_xslproxyservice.test", "result.0.force_policy_exec", "false"),
					resource.TestCheckResourceAttr("data.datapower_xslproxyservice.test", "result.0.monitor_processing_policy", "terminate-at-first-throttle"),
					resource.TestCheckResourceAttr("data.datapower_xslproxyservice.test", "result.0.debug_mode", "off"),
					resource.TestCheckResourceAttr("data.datapower_xslproxyservice.test", "result.0.debug_history", "25"),
					resource.TestCheckResourceAttr("data.datapower_xslproxyservice.test", "result.0.local_address", "0.0.0.0"),
				}...),
			},
		},
	})
}
