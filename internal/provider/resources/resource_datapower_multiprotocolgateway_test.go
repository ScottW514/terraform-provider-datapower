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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models/testconfig"
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccResourceMultiProtocolGateway(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_MultiProtocolGateway") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_MultiProtocolGateway")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testconfig.MultiProtocolGatewayTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "back_http_version", "HTTP/1.1"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "http2_required", "false"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "request_type", "soap"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "response_type", "soap"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "follow_redirects", "true"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "rewrite_location_header", "false"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "style_policy", "default"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "type", "static-backend"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "allow_compression", "false"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "enable_compressed_payload_passthrough", "false"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "allow_cache_control_header", "false"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsm_agent_monitor", "false"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsm_agent_monitor_pcm", "all-messages"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "proxy_http_response", "false"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "priority", "normal"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "xml_manager", "default"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "ssl_client_config_type", "client"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "default_param_namespace", "http://www.datapower.com/param/config"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "query_param_namespace", "http://www.datapower.com/param/query"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "propagate_uri", "true"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "monitor_processing_policy", "terminate-at-first-throttle"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "request_attachments", "strip"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "response_attachments", "strip"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "request_attachments_flow_control", "false"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "response_attachments_flow_control", "false"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "root_part_not_first_action", "process-in-order"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "front_attachment_format", "dynamic"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "back_attachment_format", "dynamic"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "mime_front_headers", "true"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "mime_back_headers", "true"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "stream_output_to_back", "buffer-until-verification"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "stream_output_to_front", "buffer-until-verification"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "gateway_parser_limits", "false"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "parser_limits_element_depth", "512"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "parser_limits_attribute_count", "128"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "parser_limits_max_node_size", "33554432"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "parser_limits_external_references", "forbid"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "parser_limits_max_prefixes", "1024"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "parser_limits_max_namespaces", "1024"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "parser_limits_max_local_names", "60000"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "parser_limits_attachment_byte_count", "2000000000"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "debug_mode", "off"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "debug_history", "25"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "flow_control", "false"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "soap_schema_url", "store:///schemas/soap-envelope.xsd"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "front_timeout", "120"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "back_timeout", "120"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "front_persistent_timeout", "180"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "back_persistent_timeout", "180"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "include_response_type_encoding", "false"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "persistent_connections", "true"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "loop_detection", "false"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "do_host_rewriting", "true"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "do_chunked_upload", "false"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "process_http_errors", "true"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "http_client_ip_label", "X-Client-IP"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "http_log_cor_id_label", "X-Global-Transaction-ID"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsa_mode", "sync2sync"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsa_require_aaa", "true"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsa_strip", "true"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsa_default_reply_to", "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsa_default_fault_to", "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsa_force", "false"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsa_gen_style", "sync"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsahttp_async_response_code", "204"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsa_timeout", "120"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsrm_enabled", "false"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsrm_sequence_expiration", "3600"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsrm_destination_accept_create_sequence", "true"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsrm_destination_maximum_sequences", "400"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsrm_destination_in_order", "false"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsrm_destination_maximum_in_order_queue_length", "10"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsrm_destination_accept_offers", "false"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsrm_front_force", "false"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsrm_back_force", "false"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsrm_back_create_sequence", "false"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsrm_front_create_sequence", "false"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsrm_source_make_offer", "false"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsrm_uses_sequence_ssl", "false"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsrm_source_maximum_sequences", "400"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsrm_source_retransmission_interval", "10"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsrm_source_exponential_backoff", "true"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsrm_source_maximum_retransmissions", "4"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsrm_source_maximum_queue_length", "30"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsrm_source_request_ack_count", "1"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "wsrm_source_inactivity_close", "360"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "force_policy_exec", "false"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "rewrite_errors", "true"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "delay_errors", "true"),
			resource.TestCheckResourceAttr("datapower_multiprotocolgateway.test", "delay_errors_duration", "1000"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
