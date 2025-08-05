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

func TestAccDataSourceXMLFirewallService(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_XMLFirewallService") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_XMLFirewallService")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.XMLFirewallServiceTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.id", "XMLFirewallService_name"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.type", "dynamic-backend"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.xml_manager", "default"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.style_policy", "default"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.max_message_size", "0"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.request_type", "soap"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.response_type", "soap"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.request_attachments", "strip"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.response_attachments", "strip"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.root_part_not_first_action", "process-in-order"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.front_attachment_format", "dynamic"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.back_attachment_format", "dynamic"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.mime_headers", "true"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.rewrite_errors", "true"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.delay_errors", "true"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.delay_errors_duration", "1000"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.soap_schema_url", "store:///schemas/soap-envelope.xsd"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.wsdl_response_policy", "off"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.firewall_parser_limits", "false"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.parser_limits_bytes_scanned", "4194304"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.parser_limits_element_depth", "512"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.parser_limits_attribute_count", "128"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.parser_limits_max_node_size", "33554432"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.parser_limits_max_prefixes", "1024"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.parser_limits_max_namespaces", "1024"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.parser_limits_max_local_names", "60000"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.parser_limits_attachment_byte_count", "2000000000"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.parser_limits_external_references", "forbid"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.credential_charset", "protocol"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.ssl_config_type", "server"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.priority", "normal"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.http_timeout", "120"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.http_persist_timeout", "180"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.do_host_rewrite", "true"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.suppress_http_warnings", "false"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.http_compression", "false"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.http_include_response_type_encoding", "false"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.always_show_errors", "false"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.disallow_get", "false"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.disallow_empty_response", "false"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.http_persistent_connections", "true"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.http_client_ip_label", "X-Client-IP"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.http_log_cor_id_label", "X-Global-Transaction-ID"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.http_proxy_port", "800"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.do_chunked_upload", "false"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.default_param_namespace", "http://www.datapower.com/param/config"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.query_param_namespace", "http://www.datapower.com/param/query"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.force_policy_exec", "false"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.monitor_processing_policy", "terminate-at-first-throttle"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.debug_mode", "off"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.debug_history", "25"),
					resource.TestCheckResourceAttr("data.datapower_xmlfirewallservice.test", "result.0.local_address", "0.0.0.0"),
				}...),
			},
		},
	})
}
