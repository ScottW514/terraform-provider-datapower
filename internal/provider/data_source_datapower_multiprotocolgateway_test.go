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

func TestAccDataSourceMultiProtocolGateway(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_MultiProtocolGateway") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_MultiProtocolGateway")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.MultiProtocolGatewayTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.id", "MultiProtocolGateway_name"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.back_http_version", "HTTP/1.1"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.request_type", "soap"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.response_type", "soap"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.style_policy", "default"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.type", "static-backend"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsm_agent_monitor_pcm", "all-messages"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.priority", "normal"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.xml_manager", "default"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.ssl_client_config_type", "client"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.default_param_namespace", "http://www.datapower.com/param/config"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.query_param_namespace", "http://www.datapower.com/param/query"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.monitor_processing_policy", "terminate-at-first-throttle"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.request_attachments", "strip"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.response_attachments", "strip"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.root_part_not_first_action", "process-in-order"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.front_attachment_format", "dynamic"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.back_attachment_format", "dynamic"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.stream_output_to_back", "buffer-until-verification"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.stream_output_to_front", "buffer-until-verification"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.parser_limits_external_references", "forbid"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.debug_mode", "off"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.soap_schema_url", "store:///schemas/soap-envelope.xsd"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.http_client_ip_label", "X-Client-IP"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.http_log_cor_id_label", "X-Global-Transaction-ID"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsa_mode", "sync2sync"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsa_default_reply_to", "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsa_default_fault_to", "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous"),
					resource.TestCheckResourceAttr("data.datapower_multiprotocolgateway.test", "result.0.wsa_gen_style", "sync"),
				}...),
			},
		},
	})
}
