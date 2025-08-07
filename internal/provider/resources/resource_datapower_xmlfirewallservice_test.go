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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/testconfig"
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccResourceXMLFirewallService(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_XMLFirewallService") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_XMLFirewallService")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testconfig.XMLFirewallServiceTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "type", "dynamic-backend"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "xml_manager", "default"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "style_policy", "default"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "max_message_size", "0"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "request_type", "soap"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "response_type", "soap"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "request_attachments", "strip"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "response_attachments", "strip"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "root_part_not_first_action", "process-in-order"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "front_attachment_format", "dynamic"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "back_attachment_format", "dynamic"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "mime_headers", "true"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "rewrite_errors", "true"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "delay_errors", "true"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "delay_errors_duration", "1000"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "soap_schema_url", "store:///schemas/soap-envelope.xsd"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "wsdl_response_policy", "off"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "firewall_parser_limits", "false"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "parser_limits_bytes_scanned", "4194304"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "parser_limits_element_depth", "512"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "parser_limits_attribute_count", "128"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "parser_limits_max_node_size", "33554432"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "parser_limits_max_prefixes", "1024"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "parser_limits_max_namespaces", "1024"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "parser_limits_max_local_names", "60000"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "parser_limits_attachment_byte_count", "2000000000"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "parser_limits_external_references", "forbid"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "credential_charset", "protocol"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "ssl_config_type", "server"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "priority", "normal"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "http_timeout", "120"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "http_persist_timeout", "180"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "do_host_rewrite", "true"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "suppress_http_warnings", "false"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "http_compression", "false"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "http_include_response_type_encoding", "false"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "always_show_errors", "false"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "disallow_get", "false"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "disallow_empty_response", "false"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "http_persistent_connections", "true"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "http_client_ip_label", "X-Client-IP"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "http_log_cor_id_label", "X-Global-Transaction-ID"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "http_proxy_port", "800"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "do_chunked_upload", "false"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "default_param_namespace", "http://www.datapower.com/param/config"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "query_param_namespace", "http://www.datapower.com/param/query"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "force_policy_exec", "false"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "monitor_processing_policy", "terminate-at-first-throttle"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "debug_mode", "off"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "debug_history", "25"),
			resource.TestCheckResourceAttr("datapower_xmlfirewallservice.test", "local_address", "0.0.0.0"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
	actions.PostProcess()
}
