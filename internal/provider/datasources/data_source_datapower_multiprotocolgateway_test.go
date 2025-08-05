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

func TestAccDataSourceMultiProtocolGateway(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_MultiProtocolGateway") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_MultiProtocolGateway")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.MultiProtocolGatewayTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.id", "MultiProtocolGateway_name"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.back_http_version", "HTTP/1.1"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.http2_required", "false"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.request_type", "soap"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.response_type", "soap"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.follow_redirects", "true"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.rewrite_location_header", "false"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.style_policy", "default"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.type", "static-backend"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.allow_compression", "false"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.enable_compressed_payload_passthrough", "false"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.allow_cache_control_header", "false"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsm_agent_monitor", "false"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsm_agent_monitor_pcm", "all-messages"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.proxy_http_response", "false"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.priority", "normal"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.xml_manager", "default"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.ssl_client_config_type", "client"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.default_param_namespace", "http://www.datapower.com/param/config"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.query_param_namespace", "http://www.datapower.com/param/query"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.propagate_uri", "true"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.monitor_processing_policy", "terminate-at-first-throttle"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.request_attachments", "strip"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.response_attachments", "strip"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.request_attachments_flow_control", "false"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.response_attachments_flow_control", "false"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.root_part_not_first_action", "process-in-order"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.front_attachment_format", "dynamic"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.back_attachment_format", "dynamic"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.mime_front_headers", "true"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.mime_back_headers", "true"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.stream_output_to_back", "buffer-until-verification"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.stream_output_to_front", "buffer-until-verification"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.gateway_parser_limits", "false"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.parser_limits_element_depth", "512"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.parser_limits_attribute_count", "128"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.parser_limits_max_node_size", "33554432"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.parser_limits_external_references", "forbid"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.parser_limits_max_prefixes", "1024"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.parser_limits_max_namespaces", "1024"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.parser_limits_max_local_names", "60000"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.parser_limits_attachment_byte_count", "2000000000"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.debug_mode", "off"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.debug_history", "25"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.flow_control", "false"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.soap_schema_url", "store:///schemas/soap-envelope.xsd"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.front_timeout", "120"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.back_timeout", "120"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.front_persistent_timeout", "180"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.back_persistent_timeout", "180"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.include_response_type_encoding", "false"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.persistent_connections", "true"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.loop_detection", "false"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.do_host_rewriting", "true"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.do_chunked_upload", "false"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.process_http_errors", "true"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.http_client_ip_label", "X-Client-IP"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.http_log_cor_id_label", "X-Global-Transaction-ID"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsa_mode", "sync2sync"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsa_require_aaa", "true"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsa_strip", "true"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsa_default_reply_to", "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsa_default_fault_to", "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsa_force", "false"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsa_gen_style", "sync"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsahttp_async_response_code", "204"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsa_timeout", "120"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsrm_enabled", "false"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsrm_sequence_expiration", "3600"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsrm_destination_accept_create_sequence", "true"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsrm_destination_maximum_sequences", "400"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsrm_destination_in_order", "false"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsrm_destination_maximum_in_order_queue_length", "10"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsrm_destination_accept_offers", "false"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsrm_front_force", "false"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsrm_back_force", "false"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsrm_back_create_sequence", "false"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsrm_front_create_sequence", "false"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsrm_source_make_offer", "false"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsrm_uses_sequence_ssl", "false"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsrm_source_maximum_sequences", "400"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsrm_source_retransmission_interval", "10"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsrm_source_exponential_backoff", "true"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsrm_source_maximum_retransmissions", "4"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsrm_source_maximum_queue_length", "30"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsrm_source_request_ack_count", "1"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsrm_source_inactivity_close", "360"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.force_policy_exec", "false"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.rewrite_errors", "true"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.delay_errors", "true"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.delay_errors_duration", "1000"),
				}...),
			},
		},
	})
}
