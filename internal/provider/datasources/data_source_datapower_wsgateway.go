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

package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type WSGatewayList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &WSGatewayDataSource{}
	_ datasource.DataSourceWithConfigure = &WSGatewayDataSource{}
)

func NewWSGatewayDataSource() datasource.DataSource {
	return &WSGatewayDataSource{}
}

type WSGatewayDataSource struct {
	pData *tfutils.ProviderData
}

func (d *WSGatewayDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wsgateway"
}

func (d *WSGatewayDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Web Service Proxy",
		Attributes: map[string]schema.Attribute{
			"app_domain": schema.StringAttribute{
				MarkdownDescription: "The name of the application domain the object belongs to",
				Required:            true,
			},
			"result": schema.ListNestedAttribute{
				MarkdownDescription: "List of objects",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Name of the object. Must be unique among object types in application domain.",
							Computed:            true,
						},
						"app_domain": schema.StringAttribute{
							MarkdownDescription: "The name of the application domain the object belongs to",
							Computed:            true,
						},
						"back_http_version": schema.StringAttribute{
							MarkdownDescription: "HTTP Version to Server",
							Computed:            true,
						},
						"request_type": schema.StringAttribute{
							MarkdownDescription: "Request Type",
							Computed:            true,
						},
						"response_type": schema.StringAttribute{
							MarkdownDescription: "Response Type",
							Computed:            true,
						},
						"follow_redirects": schema.BoolAttribute{
							MarkdownDescription: "Follow Redirects",
							Computed:            true,
						},
						"allow_compression": schema.BoolAttribute{
							MarkdownDescription: "Compression",
							Computed:            true,
						},
						"enable_compressed_payload_passthrough": schema.BoolAttribute{
							MarkdownDescription: "Enable passthrough of compressed payload",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Type",
							Computed:            true,
						},
						"auto_create_sources": schema.BoolAttribute{
							MarkdownDescription: "Auto-create Source Protocols",
							Computed:            true,
						},
						"ssl_server_config_type": schema.StringAttribute{
							MarkdownDescription: "TLS server type",
							Computed:            true,
						},
						"ssl_server": schema.StringAttribute{
							MarkdownDescription: "TLS server profile",
							Computed:            true,
						},
						"sslsni_server": schema.StringAttribute{
							MarkdownDescription: "TLS SNI server profile",
							Computed:            true,
						},
						"endpoint_rewrite_policy": schema.StringAttribute{
							MarkdownDescription: "Endpoint Rewrite Policy",
							Computed:            true,
						},
						"local_endpoint_rewrite": schema.StringAttribute{
							MarkdownDescription: "Local endpoint rewrite",
							Computed:            true,
						},
						"remote_endpoint_rewrite": schema.StringAttribute{
							MarkdownDescription: "Remote endpoint rewrite",
							Computed:            true,
						},
						"aaa_policy": schema.StringAttribute{
							MarkdownDescription: "Authorization AAA Policy",
							Computed:            true,
						},
						"style_policy": schema.StringAttribute{
							MarkdownDescription: "Processing Policy",
							Computed:            true,
						},
						"remote_fetch_retry": models.GetDmNetworkRetryDataSourceSchema("Remote Fetch Retry", "remote-retry", ""),
						"wsdl_cache_policy": schema.ListNestedAttribute{
							MarkdownDescription: "WSDL Cache Policy",
							NestedObject:        models.DmWSDLCachePolicyDataSourceSchema,
							Computed:            true,
						},
						"base_wsdl": schema.ListNestedAttribute{
							MarkdownDescription: "WSDL File",
							NestedObject:        models.DmWSBaseWSDLDataSourceSchema,
							Computed:            true,
						},
						"user_toggles": schema.ListNestedAttribute{
							MarkdownDescription: "User policy",
							NestedObject:        models.DmWSUserTogglesDataSourceSchema,
							Computed:            true,
						},
						"client_principal": schema.StringAttribute{
							MarkdownDescription: "Kerberos Encryptor Principal",
							Computed:            true,
						},
						"server_principal": schema.StringAttribute{
							MarkdownDescription: "Kerberos Decryptor Principal",
							Computed:            true,
						},
						"kerberos_keytab": schema.StringAttribute{
							MarkdownDescription: "Kerberos Keytab",
							Computed:            true,
						},
						"decrypt_key": schema.StringAttribute{
							MarkdownDescription: "Decrypt Key",
							Computed:            true,
						},
						"encrypted_key_sha1_cache_life_time": schema.Int64Attribute{
							MarkdownDescription: "EncryptedKeySHA1 Cache Lifetime",
							Computed:            true,
						},
						"preserve_key_chain": schema.BoolAttribute{
							MarkdownDescription: "Preserve EncryptedKey Chain",
							Computed:            true,
						},
						"decrypt_with_key_from_ed": schema.BoolAttribute{
							MarkdownDescription: "Decrypt with Key from EncryptedData",
							Computed:            true,
						},
						"soap_action_policy": schema.StringAttribute{
							MarkdownDescription: "SOAPAction Policy",
							Computed:            true,
						},
						"wsrr_subscriptions": schema.ListNestedAttribute{
							MarkdownDescription: "WSRR Subscriptions",
							NestedObject:        models.DmWSRRWSDLSourceDataSourceSchema,
							Computed:            true,
						},
						"wsrr_saved_search_subscriptions": schema.ListNestedAttribute{
							MarkdownDescription: "WSRR Saved Search Subscriptions",
							NestedObject:        models.DmWSRRSavedSearchWSDLSourceDataSourceSchema,
							Computed:            true,
						},
						"operation_priority": schema.ListNestedAttribute{
							MarkdownDescription: "Operation Priority",
							NestedObject:        models.DmWSOperationSchedulerPriorityDataSourceSchema,
							Computed:            true,
						},
						"operation_conformance_policy": schema.ListNestedAttribute{
							MarkdownDescription: "Operation Conformance Policy",
							NestedObject:        models.DmWSOperationConformancePolicyDataSourceSchema,
							Computed:            true,
						},
						"operation_policy_subject_opt_out": schema.ListNestedAttribute{
							MarkdownDescription: "Operation Policy Subject Opt Out",
							NestedObject:        models.DmWSOperationPolicySubjectOptOutDataSourceSchema,
							Computed:            true,
						},
						"policy_parameter": schema.ListNestedAttribute{
							MarkdownDescription: "Policy Parameters",
							NestedObject:        models.DmWSPolicyParametersDataSourceSchema,
							Computed:            true,
						},
						"reliable_messaging": schema.ListNestedAttribute{
							MarkdownDescription: "Reliable Messaging",
							NestedObject:        models.DmWSOperationReliableMessagingDataSourceSchema,
							Computed:            true,
						},
						"wsm_agent_monitor": schema.BoolAttribute{
							MarkdownDescription: "Monitor with Web Services Management Agent",
							Computed:            true,
						},
						"wsm_agent_monitor_pcm": schema.StringAttribute{
							MarkdownDescription: "Message capture with Web Services Management Agent",
							Computed:            true,
						},
						"process_resp_rules_on_one_way_mep": schema.BoolAttribute{
							MarkdownDescription: "Enable response rule processing for one-way WSDLs.",
							Computed:            true,
						},
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"priority": schema.StringAttribute{
							MarkdownDescription: "Service Priority",
							Computed:            true,
						},
						"front_protocol": schema.ListAttribute{
							MarkdownDescription: "Front Side Protocol",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"xml_manager": schema.StringAttribute{
							MarkdownDescription: "XML Manager",
							Computed:            true,
						},
						"url_rewrite_policy": schema.StringAttribute{
							MarkdownDescription: "URL Rewrite Policy",
							Computed:            true,
						},
						"ssl_client_config_type": schema.StringAttribute{
							MarkdownDescription: "TLS client type",
							Computed:            true,
						},
						"ssl_client": schema.StringAttribute{
							MarkdownDescription: "TLS client profile",
							Computed:            true,
						},
						"fw_cred": schema.StringAttribute{
							MarkdownDescription: "Crypto Credentials",
							Computed:            true,
						},
						"header_injection": schema.ListNestedAttribute{
							MarkdownDescription: "HTTP Header Injection",
							NestedObject:        models.DmHeaderInjectionDataSourceSchema,
							Computed:            true,
						},
						"header_suppression": schema.ListNestedAttribute{
							MarkdownDescription: "HTTP Header Suppression",
							NestedObject:        models.DmHeaderSuppressionDataSourceSchema,
							Computed:            true,
						},
						"stylesheet_parameters": schema.ListNestedAttribute{
							MarkdownDescription: "Stylesheet Parameter",
							NestedObject:        models.DmStylesheetParameterDataSourceSchema,
							Computed:            true,
						},
						"default_param_namespace": schema.StringAttribute{
							MarkdownDescription: "Default parameter namespace",
							Computed:            true,
						},
						"query_param_namespace": schema.StringAttribute{
							MarkdownDescription: "Query parameter namespace",
							Computed:            true,
						},
						"backend_url": schema.StringAttribute{
							MarkdownDescription: "Backend URL",
							Computed:            true,
						},
						"propagate_uri": schema.BoolAttribute{
							MarkdownDescription: "Propagate URI",
							Computed:            true,
						},
						"service_monitors": schema.ListAttribute{
							MarkdownDescription: "Service Monitors",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"count_monitors": schema.ListAttribute{
							MarkdownDescription: "Count Monitors",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"duration_monitors": schema.ListAttribute{
							MarkdownDescription: "Duration Monitors",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"monitor_processing_policy": schema.StringAttribute{
							MarkdownDescription: "Monitors evaluation method",
							Computed:            true,
						},
						"request_attachments": schema.StringAttribute{
							MarkdownDescription: "Request attachment processing mode",
							Computed:            true,
						},
						"response_attachments": schema.StringAttribute{
							MarkdownDescription: "Response attachment processing mode",
							Computed:            true,
						},
						"request_attachments_flow_control": schema.BoolAttribute{
							MarkdownDescription: "Request attachment flow control mode",
							Computed:            true,
						},
						"response_attachments_flow_control": schema.BoolAttribute{
							MarkdownDescription: "Response attachment flow control mode",
							Computed:            true,
						},
						"root_part_not_first_action": schema.StringAttribute{
							MarkdownDescription: "Action when required root part is not first",
							Computed:            true,
						},
						"front_attachment_format": schema.StringAttribute{
							MarkdownDescription: "Front attachment processing format",
							Computed:            true,
						},
						"back_attachment_format": schema.StringAttribute{
							MarkdownDescription: "Back attachment processing format",
							Computed:            true,
						},
						"mime_front_headers": schema.BoolAttribute{
							MarkdownDescription: "Front Side MIME Header Processing",
							Computed:            true,
						},
						"mime_back_headers": schema.BoolAttribute{
							MarkdownDescription: "Back Side MIME Header Processing",
							Computed:            true,
						},
						"stream_output_to_back": schema.StringAttribute{
							MarkdownDescription: "Stream Output to Back",
							Computed:            true,
						},
						"stream_output_to_front": schema.StringAttribute{
							MarkdownDescription: "Stream Output to Front",
							Computed:            true,
						},
						"max_message_size": schema.Int64Attribute{
							MarkdownDescription: "Maximum Message Size",
							Computed:            true,
						},
						"gateway_parser_limits": schema.BoolAttribute{
							MarkdownDescription: "Parser limits",
							Computed:            true,
						},
						"parser_limits_element_depth": schema.Int64Attribute{
							MarkdownDescription: "XML Element Depth",
							Computed:            true,
						},
						"parser_limits_attribute_count": schema.Int64Attribute{
							MarkdownDescription: "XML Attribute Count",
							Computed:            true,
						},
						"parser_limits_max_node_size": schema.Int64Attribute{
							MarkdownDescription: "XML Maximum Node Size",
							Computed:            true,
						},
						"parser_limits_external_references": schema.StringAttribute{
							MarkdownDescription: "XML External Reference Handling",
							Computed:            true,
						},
						"parser_limits_max_prefixes": schema.Int64Attribute{
							MarkdownDescription: "XML Maximum Distinct Prefixes",
							Computed:            true,
						},
						"parser_limits_max_namespaces": schema.Int64Attribute{
							MarkdownDescription: "XML Maximum Distinct Namespaces",
							Computed:            true,
						},
						"parser_limits_max_local_names": schema.Int64Attribute{
							MarkdownDescription: "XML Maximum Distinct Local Names",
							Computed:            true,
						},
						"parser_limits_attachment_byte_count": schema.Int64Attribute{
							MarkdownDescription: "Attachment Byte Count Limit",
							Computed:            true,
						},
						"parser_limits_attachment_package_byte_count": schema.Int64Attribute{
							MarkdownDescription: "Attachment Package Byte Count Limit",
							Computed:            true,
						},
						"debug_mode": schema.StringAttribute{
							MarkdownDescription: "Probe setting",
							Computed:            true,
						},
						"debug_history": schema.Int64Attribute{
							MarkdownDescription: "Transaction History",
							Computed:            true,
						},
						"debug_trigger": schema.ListNestedAttribute{
							MarkdownDescription: "Probe Triggers",
							NestedObject:        models.DmMSDebugTriggerTypeDataSourceSchema,
							Computed:            true,
						},
						"flow_control": schema.BoolAttribute{
							MarkdownDescription: "Flow Control",
							Computed:            true,
						},
						"soap_schema_url": schema.StringAttribute{
							MarkdownDescription: "SOAP Schema URL",
							Computed:            true,
						},
						"front_timeout": schema.Int64Attribute{
							MarkdownDescription: "Front Side Timeout",
							Computed:            true,
						},
						"back_timeout": schema.Int64Attribute{
							MarkdownDescription: "Back Side Timeout",
							Computed:            true,
						},
						"front_persistent_timeout": schema.Int64Attribute{
							MarkdownDescription: "Front Persistent Timeout",
							Computed:            true,
						},
						"back_persistent_timeout": schema.Int64Attribute{
							MarkdownDescription: "Back Persistent Timeout",
							Computed:            true,
						},
						"include_response_type_encoding": schema.BoolAttribute{
							MarkdownDescription: "Include charset in response-type",
							Computed:            true,
						},
						"persistent_connections": schema.BoolAttribute{
							MarkdownDescription: "Persistent Connections",
							Computed:            true,
						},
						"loop_detection": schema.BoolAttribute{
							MarkdownDescription: "Loop Detection",
							Computed:            true,
						},
						"do_host_rewriting": schema.BoolAttribute{
							MarkdownDescription: "Rewrite Host Names when Redirecting",
							Computed:            true,
						},
						"do_chunked_upload": schema.BoolAttribute{
							MarkdownDescription: "Allow Chunked Uploads",
							Computed:            true,
						},
						"process_http_errors": schema.BoolAttribute{
							MarkdownDescription: "Process Backend Errors",
							Computed:            true,
						},
						"http_client_ip_label": schema.StringAttribute{
							MarkdownDescription: "HTTP Client IP Label",
							Computed:            true,
						},
						"http_log_cor_id_label": schema.StringAttribute{
							MarkdownDescription: "HTTP Global Transaction ID Label",
							Computed:            true,
						},
						"load_balancer_hash_header": schema.StringAttribute{
							MarkdownDescription: "Load Balancer Hash Header",
							Computed:            true,
						},
						"in_order_mode": models.GetDmGatewayInOrderModeDataSourceSchema("Message Processing Modes", "inorder-mode", ""),
						"wsa_mode": schema.StringAttribute{
							MarkdownDescription: "WS-Addressing Mode",
							Computed:            true,
						},
						"wsa_require_aaa": schema.BoolAttribute{
							MarkdownDescription: "Require AAA for WS-Addressing trust",
							Computed:            true,
						},
						"wsa_rewrite_reply_to": schema.StringAttribute{
							MarkdownDescription: "Rewrite WS Addressing Reply To Header",
							Computed:            true,
						},
						"wsa_rewrite_fault_to": schema.StringAttribute{
							MarkdownDescription: "Rewrite WS Addressing Fault To Header",
							Computed:            true,
						},
						"wsa_rewrite_to": schema.StringAttribute{
							MarkdownDescription: "Rewrite WS Addressing To Header",
							Computed:            true,
						},
						"wsa_strip": schema.BoolAttribute{
							MarkdownDescription: "Strip WS-Addressing Headers",
							Computed:            true,
						},
						"wsa_default_reply_to": schema.StringAttribute{
							MarkdownDescription: "Default WS-Addressing Reply-To",
							Computed:            true,
						},
						"wsa_default_fault_to": schema.StringAttribute{
							MarkdownDescription: "Default WS-Addressing Fault-To",
							Computed:            true,
						},
						"wsa_force": schema.BoolAttribute{
							MarkdownDescription: "Force Insertion of WS-Addressing Headers",
							Computed:            true,
						},
						"wsa_gen_style": schema.StringAttribute{
							MarkdownDescription: "WS-Addressing Message Generation Pattern",
							Computed:            true,
						},
						"wsahttp_async_response_code": schema.Int64Attribute{
							MarkdownDescription: "Asynchronous HTTP ResponseCode",
							Computed:            true,
						},
						"wsa_back_protocol": schema.StringAttribute{
							MarkdownDescription: "WS-Addressing Reply Point",
							Computed:            true,
						},
						"wsa_timeout": schema.Int64Attribute{
							MarkdownDescription: "WS-Addressing Asynchronous Reply Timeout",
							Computed:            true,
						},
						"wsrm_enabled": schema.BoolAttribute{
							MarkdownDescription: "WS-ReliableMessaging",
							Computed:            true,
						},
						"wsrm_sequence_expiration": schema.Int64Attribute{
							MarkdownDescription: "Target Sequence Expiration Interval",
							Computed:            true,
						},
						"wsrmaaa_policy": schema.StringAttribute{
							MarkdownDescription: "WS-RM AAA Policy",
							Computed:            true,
						},
						"wsrm_destination_accept_create_sequence": schema.BoolAttribute{
							MarkdownDescription: "Destination Accept Incoming CreateSequence",
							Computed:            true,
						},
						"wsrm_destination_maximum_sequences": schema.Int64Attribute{
							MarkdownDescription: "Destination Maximum Simultaneous Sequences",
							Computed:            true,
						},
						"wsrm_destination_in_order": schema.BoolAttribute{
							MarkdownDescription: "Destination InOrder Delivery Assurance",
							Computed:            true,
						},
						"wsrm_destination_maximum_in_order_queue_length": schema.Int64Attribute{
							MarkdownDescription: "Destination Maximum InOrder Queue Length",
							Computed:            true,
						},
						"wsrm_destination_accept_offers": schema.BoolAttribute{
							MarkdownDescription: "Destination Accept Two-Way Offers",
							Computed:            true,
						},
						"wsrm_front_force": schema.BoolAttribute{
							MarkdownDescription: "WS-RM Required on Request",
							Computed:            true,
						},
						"wsrm_back_force": schema.BoolAttribute{
							MarkdownDescription: "WS-RM Required on Response",
							Computed:            true,
						},
						"wsrm_back_create_sequence": schema.BoolAttribute{
							MarkdownDescription: "Source Create Sequence on Request",
							Computed:            true,
						},
						"wsrm_front_create_sequence": schema.BoolAttribute{
							MarkdownDescription: "Source Create Sequence on Response",
							Computed:            true,
						},
						"wsrm_source_make_offer": schema.BoolAttribute{
							MarkdownDescription: "Source Make Two-Way Offer",
							Computed:            true,
						},
						"wsrm_uses_sequence_ssl": schema.BoolAttribute{
							MarkdownDescription: "Source use SSL/TLS Session Binding",
							Computed:            true,
						},
						"wsrm_front_acks_to": schema.StringAttribute{
							MarkdownDescription: "Source Front Reply Point",
							Computed:            true,
						},
						"wsrm_back_acks_to": schema.StringAttribute{
							MarkdownDescription: "Source Back AcksTo Reply Point",
							Computed:            true,
						},
						"wsrm_source_maximum_sequences": schema.Int64Attribute{
							MarkdownDescription: "Source Maximum Simultaneous Sequences",
							Computed:            true,
						},
						"wsrm_source_retransmission_interval": schema.Int64Attribute{
							MarkdownDescription: "Source Retransmission Interval",
							Computed:            true,
						},
						"wsrm_source_exponential_backoff": schema.BoolAttribute{
							MarkdownDescription: "Source Exponential Backoff",
							Computed:            true,
						},
						"wsrm_source_maximum_retransmissions": schema.Int64Attribute{
							MarkdownDescription: "Source Maximum Retransmissions",
							Computed:            true,
						},
						"wsrm_source_maximum_queue_length": schema.Int64Attribute{
							MarkdownDescription: "Source Maximum Queue Length",
							Computed:            true,
						},
						"wsrm_source_request_ack_count": schema.Int64Attribute{
							MarkdownDescription: "Source Request Ack Count",
							Computed:            true,
						},
						"wsrm_source_inactivity_close": schema.Int64Attribute{
							MarkdownDescription: "Source Inactivity Close Interval",
							Computed:            true,
						},
						"force_policy_exec": schema.BoolAttribute{
							MarkdownDescription: "Process Messages Whose Body Is Empty",
							Computed:            true,
						},
						"rewrite_errors": schema.BoolAttribute{
							MarkdownDescription: "Rewrite Error Messages",
							Computed:            true,
						},
						"delay_errors": schema.BoolAttribute{
							MarkdownDescription: "Delay Error Messages",
							Computed:            true,
						},
						"delay_errors_duration": schema.Int64Attribute{
							MarkdownDescription: "Duration to Delay Error Messages",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *WSGatewayDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *WSGatewayDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data WSGatewayList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.WSGateway{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.WSGateway{}
	if value := res.Get(`WSGateway`); value.Exists() {
		for _, v := range value.Array() {
			item := models.WSGateway{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.WSGatewayObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.WSGatewayObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
