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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models/testconfig"
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccResourceXSLProxyService(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_XSLProxyService") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_XSLProxyService")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testconfig.XSLProxyServiceTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_xslproxyservice.test", "type", "static-backend"),
			resource.TestCheckResourceAttr("datapower_xslproxyservice.test", "xml_manager", "default"),
			resource.TestCheckResourceAttr("datapower_xslproxyservice.test", "style_policy", "default"),
			resource.TestCheckResourceAttr("datapower_xslproxyservice.test", "credential_charset", "protocol"),
			resource.TestCheckResourceAttr("datapower_xslproxyservice.test", "ssl_config_type", "server"),
			resource.TestCheckResourceAttr("datapower_xslproxyservice.test", "priority", "normal"),
			resource.TestCheckResourceAttr("datapower_xslproxyservice.test", "http_timeout", "120"),
			resource.TestCheckResourceAttr("datapower_xslproxyservice.test", "http_persist_timeout", "180"),
			resource.TestCheckResourceAttr("datapower_xslproxyservice.test", "do_host_rewrite", "true"),
			resource.TestCheckResourceAttr("datapower_xslproxyservice.test", "suppress_http_warnings", "false"),
			resource.TestCheckResourceAttr("datapower_xslproxyservice.test", "http_compression", "false"),
			resource.TestCheckResourceAttr("datapower_xslproxyservice.test", "http_include_response_type_encoding", "false"),
			resource.TestCheckResourceAttr("datapower_xslproxyservice.test", "always_show_errors", "false"),
			resource.TestCheckResourceAttr("datapower_xslproxyservice.test", "disallow_get", "false"),
			resource.TestCheckResourceAttr("datapower_xslproxyservice.test", "disallow_empty_response", "false"),
			resource.TestCheckResourceAttr("datapower_xslproxyservice.test", "http_persistent_connections", "true"),
			resource.TestCheckResourceAttr("datapower_xslproxyservice.test", "http_client_ip_label", "X-Client-IP"),
			resource.TestCheckResourceAttr("datapower_xslproxyservice.test", "http_log_cor_id_label", "X-Global-Transaction-ID"),
			resource.TestCheckResourceAttr("datapower_xslproxyservice.test", "http_proxy_port", "800"),
			resource.TestCheckResourceAttr("datapower_xslproxyservice.test", "do_chunked_upload", "false"),
			resource.TestCheckResourceAttr("datapower_xslproxyservice.test", "default_param_namespace", "http://www.datapower.com/param/config"),
			resource.TestCheckResourceAttr("datapower_xslproxyservice.test", "query_param_namespace", "http://www.datapower.com/param/query"),
			resource.TestCheckResourceAttr("datapower_xslproxyservice.test", "force_policy_exec", "false"),
			resource.TestCheckResourceAttr("datapower_xslproxyservice.test", "monitor_processing_policy", "terminate-at-first-throttle"),
			resource.TestCheckResourceAttr("datapower_xslproxyservice.test", "debug_mode", "off"),
			resource.TestCheckResourceAttr("datapower_xslproxyservice.test", "debug_history", "25"),
			resource.TestCheckResourceAttr("datapower_xslproxyservice.test", "local_address", "0.0.0.0"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
	actions.PostProcess()
}
