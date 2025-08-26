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

package models

import (
	"context"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type WSGateway struct {
	Id                                       types.String                `tfsdk:"id"`
	AppDomain                                types.String                `tfsdk:"app_domain"`
	BackHttpVersion                          types.String                `tfsdk:"back_http_version"`
	RequestType                              types.String                `tfsdk:"request_type"`
	ResponseType                             types.String                `tfsdk:"response_type"`
	FollowRedirects                          types.Bool                  `tfsdk:"follow_redirects"`
	AllowCompression                         types.Bool                  `tfsdk:"allow_compression"`
	EnableCompressedPayloadPassthrough       types.Bool                  `tfsdk:"enable_compressed_payload_passthrough"`
	Type                                     types.String                `tfsdk:"type"`
	AutoCreateSourceS                        types.Bool                  `tfsdk:"auto_create_source_s"`
	SslServerConfigType                      types.String                `tfsdk:"ssl_server_config_type"`
	SslServer                                types.String                `tfsdk:"ssl_server"`
	SslSniServer                             types.String                `tfsdk:"ssl_sni_server"`
	EndpointRewritePolicy                    types.String                `tfsdk:"endpoint_rewrite_policy"`
	LocalEndpointRewrite                     types.String                `tfsdk:"local_endpoint_rewrite"`
	RemoteEndpointRewrite                    types.String                `tfsdk:"remote_endpoint_rewrite"`
	AaaPolicy                                types.String                `tfsdk:"aaa_policy"`
	StylePolicy                              types.String                `tfsdk:"style_policy"`
	RemoteFetchRetry                         *DmNetworkRetry             `tfsdk:"remote_fetch_retry"`
	WsdlCachePolicy                          types.List                  `tfsdk:"wsdl_cache_policy"`
	BaseWsdl                                 types.List                  `tfsdk:"base_wsdl"`
	UserToggles                              types.List                  `tfsdk:"user_toggles"`
	ClientPrincipal                          types.String                `tfsdk:"client_principal"`
	ServerPrincipal                          types.String                `tfsdk:"server_principal"`
	KerberosKeytab                           types.String                `tfsdk:"kerberos_keytab"`
	DecryptKey                               types.String                `tfsdk:"decrypt_key"`
	EncryptedKeySha1CacheLifeTime            types.Int64                 `tfsdk:"encrypted_key_sha1_cache_life_time"`
	PreserveKeyChain                         types.Bool                  `tfsdk:"preserve_key_chain"`
	DecryptWithKeyFromEd                     types.Bool                  `tfsdk:"decrypt_with_key_from_ed"`
	SoapActionPolicy                         types.String                `tfsdk:"soap_action_policy"`
	WsrrSubscriptions                        types.List                  `tfsdk:"wsrr_subscriptions"`
	WsrrSavedSearchSubscriptions             types.List                  `tfsdk:"wsrr_saved_search_subscriptions"`
	OperationPriority                        types.List                  `tfsdk:"operation_priority"`
	OperationConformancePolicy               types.List                  `tfsdk:"operation_conformance_policy"`
	OperationPolicySubjectOptOut             types.List                  `tfsdk:"operation_policy_subject_opt_out"`
	PolicyParameter                          types.List                  `tfsdk:"policy_parameter"`
	ReliableMessaging                        types.List                  `tfsdk:"reliable_messaging"`
	WSmAgentMonitor                          types.Bool                  `tfsdk:"w_sm_agent_monitor"`
	WSmAgentMonitorPcm                       types.String                `tfsdk:"w_sm_agent_monitor_pcm"`
	ProcessRespRulesOnOneWayMep              types.Bool                  `tfsdk:"process_resp_rules_on_one_way_mep"`
	UserSummary                              types.String                `tfsdk:"user_summary"`
	Priority                                 types.String                `tfsdk:"priority"`
	FrontProtocol                            types.List                  `tfsdk:"front_protocol"`
	XmlManager                               types.String                `tfsdk:"xml_manager"`
	UrlRewritePolicy                         types.String                `tfsdk:"url_rewrite_policy"`
	SslClientConfigType                      types.String                `tfsdk:"ssl_client_config_type"`
	SslClient                                types.String                `tfsdk:"ssl_client"`
	FwCred                                   types.String                `tfsdk:"fw_cred"`
	HeaderInjection                          types.List                  `tfsdk:"header_injection"`
	HeaderSuppression                        types.List                  `tfsdk:"header_suppression"`
	StylesheetParameters                     types.List                  `tfsdk:"stylesheet_parameters"`
	DefaultParamNamespace                    types.String                `tfsdk:"default_param_namespace"`
	QueryParamNamespace                      types.String                `tfsdk:"query_param_namespace"`
	BackendUrl                               types.String                `tfsdk:"backend_url"`
	PropagateUri                             types.Bool                  `tfsdk:"propagate_uri"`
	ServiceMonitors                          types.List                  `tfsdk:"service_monitors"`
	CountMonitors                            types.List                  `tfsdk:"count_monitors"`
	DurationMonitors                         types.List                  `tfsdk:"duration_monitors"`
	MonitorProcessingPolicy                  types.String                `tfsdk:"monitor_processing_policy"`
	RequestAttachments                       types.String                `tfsdk:"request_attachments"`
	ResponseAttachments                      types.String                `tfsdk:"response_attachments"`
	RequestAttachmentsFlowControl            types.Bool                  `tfsdk:"request_attachments_flow_control"`
	ResponseAttachmentsFlowControl           types.Bool                  `tfsdk:"response_attachments_flow_control"`
	RootPartNotFirstAction                   types.String                `tfsdk:"root_part_not_first_action"`
	FrontAttachmentFormat                    types.String                `tfsdk:"front_attachment_format"`
	BackAttachmentFormat                     types.String                `tfsdk:"back_attachment_format"`
	MimeFrontHeaders                         types.Bool                  `tfsdk:"mime_front_headers"`
	MimeBackHeaders                          types.Bool                  `tfsdk:"mime_back_headers"`
	StreamOutputToBack                       types.String                `tfsdk:"stream_output_to_back"`
	StreamOutputToFront                      types.String                `tfsdk:"stream_output_to_front"`
	MaxMessageSize                           types.Int64                 `tfsdk:"max_message_size"`
	GatewayParserLimits                      types.Bool                  `tfsdk:"gateway_parser_limits"`
	ParserLimitsElementDepth                 types.Int64                 `tfsdk:"parser_limits_element_depth"`
	ParserLimitsAttributeCount               types.Int64                 `tfsdk:"parser_limits_attribute_count"`
	ParserLimitsMaxNodeSize                  types.Int64                 `tfsdk:"parser_limits_max_node_size"`
	ParserLimitsExternalReferences           types.String                `tfsdk:"parser_limits_external_references"`
	ParserLimitsMaxPrefixes                  types.Int64                 `tfsdk:"parser_limits_max_prefixes"`
	ParserLimitsMaxNamespaces                types.Int64                 `tfsdk:"parser_limits_max_namespaces"`
	ParserLimitsMaxLocalNames                types.Int64                 `tfsdk:"parser_limits_max_local_names"`
	ParserLimitsAttachmentByteCount          types.Int64                 `tfsdk:"parser_limits_attachment_byte_count"`
	ParserLimitsAttachmentPackageByteCount   types.Int64                 `tfsdk:"parser_limits_attachment_package_byte_count"`
	DebugMode                                types.String                `tfsdk:"debug_mode"`
	DebugHistory                             types.Int64                 `tfsdk:"debug_history"`
	DebugTrigger                             types.List                  `tfsdk:"debug_trigger"`
	FlowControl                              types.Bool                  `tfsdk:"flow_control"`
	SoapSchemaUrl                            types.String                `tfsdk:"soap_schema_url"`
	FrontTimeout                             types.Int64                 `tfsdk:"front_timeout"`
	BackTimeout                              types.Int64                 `tfsdk:"back_timeout"`
	FrontPersistentTimeout                   types.Int64                 `tfsdk:"front_persistent_timeout"`
	BackPersistentTimeout                    types.Int64                 `tfsdk:"back_persistent_timeout"`
	IncludeResponseTypeEncoding              types.Bool                  `tfsdk:"include_response_type_encoding"`
	PersistentConnections                    types.Bool                  `tfsdk:"persistent_connections"`
	LoopDetection                            types.Bool                  `tfsdk:"loop_detection"`
	DoHostRewriting                          types.Bool                  `tfsdk:"do_host_rewriting"`
	DoChunkedUpload                          types.Bool                  `tfsdk:"do_chunked_upload"`
	ProcessHttpErrors                        types.Bool                  `tfsdk:"process_http_errors"`
	HttpClientIpLabel                        types.String                `tfsdk:"http_client_ip_label"`
	HttpLogCorIdLabel                        types.String                `tfsdk:"http_log_cor_id_label"`
	LoadBalancerHashHeader                   types.String                `tfsdk:"load_balancer_hash_header"`
	InOrderMode                              *DmGatewayInOrderMode       `tfsdk:"in_order_mode"`
	WsAMode                                  types.String                `tfsdk:"ws_a_mode"`
	WsARequireAaa                            types.Bool                  `tfsdk:"ws_a_require_aaa"`
	WsARewriteReplyTo                        types.String                `tfsdk:"ws_a_rewrite_reply_to"`
	WsARewriteFaultTo                        types.String                `tfsdk:"ws_a_rewrite_fault_to"`
	WsARewriteTo                             types.String                `tfsdk:"ws_a_rewrite_to"`
	WsAStrip                                 types.Bool                  `tfsdk:"ws_a_strip"`
	WsADefaultReplyTo                        types.String                `tfsdk:"ws_a_default_reply_to"`
	WsADefaultFaultTo                        types.String                `tfsdk:"ws_a_default_fault_to"`
	WsAForce                                 types.Bool                  `tfsdk:"ws_a_force"`
	WsAGenStyle                              types.String                `tfsdk:"ws_a_gen_style"`
	WsAHttpAsyncResponseCode                 types.Int64                 `tfsdk:"ws_a_http_async_response_code"`
	WsABackProtocol                          types.String                `tfsdk:"ws_a_back_protocol"`
	WsATimeout                               types.Int64                 `tfsdk:"ws_a_timeout"`
	WsRmEnabled                              types.Bool                  `tfsdk:"ws_rm_enabled"`
	WsRmSequenceExpiration                   types.Int64                 `tfsdk:"ws_rm_sequence_expiration"`
	WsRmAaaPolicy                            types.String                `tfsdk:"ws_rm_aaa_policy"`
	WsRmDestinationAcceptCreateSequence      types.Bool                  `tfsdk:"ws_rm_destination_accept_create_sequence"`
	WsRmDestinationMaximumSequences          types.Int64                 `tfsdk:"ws_rm_destination_maximum_sequences"`
	WsRmDestinationInOrder                   types.Bool                  `tfsdk:"ws_rm_destination_in_order"`
	WsRmDestinationMaximumInOrderQueueLength types.Int64                 `tfsdk:"ws_rm_destination_maximum_in_order_queue_length"`
	WsRmDestinationAcceptOffers              types.Bool                  `tfsdk:"ws_rm_destination_accept_offers"`
	WsRmFrontForce                           types.Bool                  `tfsdk:"ws_rm_front_force"`
	WsRmBackForce                            types.Bool                  `tfsdk:"ws_rm_back_force"`
	WsRmBackCreateSequence                   types.Bool                  `tfsdk:"ws_rm_back_create_sequence"`
	WsRmFrontCreateSequence                  types.Bool                  `tfsdk:"ws_rm_front_create_sequence"`
	WsRmSourceMakeOffer                      types.Bool                  `tfsdk:"ws_rm_source_make_offer"`
	WsRmUsesSequenceSsl                      types.Bool                  `tfsdk:"ws_rm_uses_sequence_ssl"`
	WsRmFrontAcksTo                          types.String                `tfsdk:"ws_rm_front_acks_to"`
	WsRmBackAcksTo                           types.String                `tfsdk:"ws_rm_back_acks_to"`
	WsRmSourceMaximumSequences               types.Int64                 `tfsdk:"ws_rm_source_maximum_sequences"`
	WsRmSourceRetransmissionInterval         types.Int64                 `tfsdk:"ws_rm_source_retransmission_interval"`
	WsRmSourceExponentialBackoff             types.Bool                  `tfsdk:"ws_rm_source_exponential_backoff"`
	WsRmSourceMaximumRetransmissions         types.Int64                 `tfsdk:"ws_rm_source_maximum_retransmissions"`
	WsRmSourceMaximumQueueLength             types.Int64                 `tfsdk:"ws_rm_source_maximum_queue_length"`
	WsRmSourceRequestAckCount                types.Int64                 `tfsdk:"ws_rm_source_request_ack_count"`
	WsRmSourceInactivityClose                types.Int64                 `tfsdk:"ws_rm_source_inactivity_close"`
	ForcePolicyExec                          types.Bool                  `tfsdk:"force_policy_exec"`
	RewriteErrors                            types.Bool                  `tfsdk:"rewrite_errors"`
	DelayErrors                              types.Bool                  `tfsdk:"delay_errors"`
	DelayErrorsDuration                      types.Int64                 `tfsdk:"delay_errors_duration"`
	DependencyActions                        []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var WSGatewayRequestAttachmentsFlowControlCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "request_attachments",
			AttrType:    "String",
			AttrDefault: "strip",
			Value:       []string{"unprocessed"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "request_type",
			AttrType:    "String",
			AttrDefault: "soap",
			Value:       []string{"unprocessed"},
		},
	},
}
var WSGatewayResponseAttachmentsFlowControlCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "response_attachments",
			AttrType:    "String",
			AttrDefault: "strip",
			Value:       []string{"unprocessed"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "response_type",
			AttrType:    "String",
			AttrDefault: "soap",
			Value:       []string{"unprocessed"},
		},
	},
}
var WSGatewayParserLimitsElementDepthCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "gateway_parser_limits",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var WSGatewayParserLimitsAttributeCountCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "gateway_parser_limits",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var WSGatewayParserLimitsMaxNodeSizeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "gateway_parser_limits",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var WSGatewayParserLimitsExternalReferencesCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "gateway_parser_limits",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var WSGatewayParserLimitsMaxPrefixesCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "gateway_parser_limits",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var WSGatewayParserLimitsMaxNamespacesCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "gateway_parser_limits",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var WSGatewayParserLimitsMaxLocalNamesCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "gateway_parser_limits",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var WSGatewayParserLimitsAttachmentByteCountCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "gateway_parser_limits",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var WSGatewayParserLimitsAttachmentPackageByteCountCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "gateway_parser_limits",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var WSGatewayDebugHistoryCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "debug_mode",
	AttrType:    "String",
	AttrDefault: "off",
	Value:       []string{"true"},
}
var WSGatewayWSAHTTPAsyncResponseCodeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "ws_a_mode",
	AttrType:    "String",
	AttrDefault: "sync2sync",
	Value:       []string{"wsa2sync", "wsa2wsa", "sync2wsa"},
}
var WSGatewayWSABackProtocolCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ws_a_gen_style",
			AttrType:    "String",
			AttrDefault: "sync",
			Value:       []string{"async"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ws_a_mode",
			AttrType:    "String",
			AttrDefault: "sync2sync",
			Value:       []string{"wsa2sync", "wsa2wsa", "sync2wsa"},
		},
	},
}
var WSGatewayDelayErrorsDurationCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "delay_errors",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "rewrite_errors",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
	},
}

var WSGatewayObjectType = map[string]attr.Type{
	"id":                                              types.StringType,
	"app_domain":                                      types.StringType,
	"back_http_version":                               types.StringType,
	"request_type":                                    types.StringType,
	"response_type":                                   types.StringType,
	"follow_redirects":                                types.BoolType,
	"allow_compression":                               types.BoolType,
	"enable_compressed_payload_passthrough":           types.BoolType,
	"type":                                            types.StringType,
	"auto_create_source_s":                            types.BoolType,
	"ssl_server_config_type":                          types.StringType,
	"ssl_server":                                      types.StringType,
	"ssl_sni_server":                                  types.StringType,
	"endpoint_rewrite_policy":                         types.StringType,
	"local_endpoint_rewrite":                          types.StringType,
	"remote_endpoint_rewrite":                         types.StringType,
	"aaa_policy":                                      types.StringType,
	"style_policy":                                    types.StringType,
	"remote_fetch_retry":                              types.ObjectType{AttrTypes: DmNetworkRetryObjectType},
	"wsdl_cache_policy":                               types.ListType{ElemType: types.ObjectType{AttrTypes: DmWSDLCachePolicyObjectType}},
	"base_wsdl":                                       types.ListType{ElemType: types.ObjectType{AttrTypes: DmWSBaseWSDLObjectType}},
	"user_toggles":                                    types.ListType{ElemType: types.ObjectType{AttrTypes: DmWSUserTogglesObjectType}},
	"client_principal":                                types.StringType,
	"server_principal":                                types.StringType,
	"kerberos_keytab":                                 types.StringType,
	"decrypt_key":                                     types.StringType,
	"encrypted_key_sha1_cache_life_time":              types.Int64Type,
	"preserve_key_chain":                              types.BoolType,
	"decrypt_with_key_from_ed":                        types.BoolType,
	"soap_action_policy":                              types.StringType,
	"wsrr_subscriptions":                              types.ListType{ElemType: types.ObjectType{AttrTypes: DmWSRRWSDLSourceObjectType}},
	"wsrr_saved_search_subscriptions":                 types.ListType{ElemType: types.ObjectType{AttrTypes: DmWSRRSavedSearchWSDLSourceObjectType}},
	"operation_priority":                              types.ListType{ElemType: types.ObjectType{AttrTypes: DmWSOperationSchedulerPriorityObjectType}},
	"operation_conformance_policy":                    types.ListType{ElemType: types.ObjectType{AttrTypes: DmWSOperationConformancePolicyObjectType}},
	"operation_policy_subject_opt_out":                types.ListType{ElemType: types.ObjectType{AttrTypes: DmWSOperationPolicySubjectOptOutObjectType}},
	"policy_parameter":                                types.ListType{ElemType: types.ObjectType{AttrTypes: DmWSPolicyParametersObjectType}},
	"reliable_messaging":                              types.ListType{ElemType: types.ObjectType{AttrTypes: DmWSOperationReliableMessagingObjectType}},
	"w_sm_agent_monitor":                              types.BoolType,
	"w_sm_agent_monitor_pcm":                          types.StringType,
	"process_resp_rules_on_one_way_mep":               types.BoolType,
	"user_summary":                                    types.StringType,
	"priority":                                        types.StringType,
	"front_protocol":                                  types.ListType{ElemType: types.StringType},
	"xml_manager":                                     types.StringType,
	"url_rewrite_policy":                              types.StringType,
	"ssl_client_config_type":                          types.StringType,
	"ssl_client":                                      types.StringType,
	"fw_cred":                                         types.StringType,
	"header_injection":                                types.ListType{ElemType: types.ObjectType{AttrTypes: DmHeaderInjectionObjectType}},
	"header_suppression":                              types.ListType{ElemType: types.ObjectType{AttrTypes: DmHeaderSuppressionObjectType}},
	"stylesheet_parameters":                           types.ListType{ElemType: types.ObjectType{AttrTypes: DmStylesheetParameterObjectType}},
	"default_param_namespace":                         types.StringType,
	"query_param_namespace":                           types.StringType,
	"backend_url":                                     types.StringType,
	"propagate_uri":                                   types.BoolType,
	"service_monitors":                                types.ListType{ElemType: types.StringType},
	"count_monitors":                                  types.ListType{ElemType: types.StringType},
	"duration_monitors":                               types.ListType{ElemType: types.StringType},
	"monitor_processing_policy":                       types.StringType,
	"request_attachments":                             types.StringType,
	"response_attachments":                            types.StringType,
	"request_attachments_flow_control":                types.BoolType,
	"response_attachments_flow_control":               types.BoolType,
	"root_part_not_first_action":                      types.StringType,
	"front_attachment_format":                         types.StringType,
	"back_attachment_format":                          types.StringType,
	"mime_front_headers":                              types.BoolType,
	"mime_back_headers":                               types.BoolType,
	"stream_output_to_back":                           types.StringType,
	"stream_output_to_front":                          types.StringType,
	"max_message_size":                                types.Int64Type,
	"gateway_parser_limits":                           types.BoolType,
	"parser_limits_element_depth":                     types.Int64Type,
	"parser_limits_attribute_count":                   types.Int64Type,
	"parser_limits_max_node_size":                     types.Int64Type,
	"parser_limits_external_references":               types.StringType,
	"parser_limits_max_prefixes":                      types.Int64Type,
	"parser_limits_max_namespaces":                    types.Int64Type,
	"parser_limits_max_local_names":                   types.Int64Type,
	"parser_limits_attachment_byte_count":             types.Int64Type,
	"parser_limits_attachment_package_byte_count":     types.Int64Type,
	"debug_mode":                                      types.StringType,
	"debug_history":                                   types.Int64Type,
	"debug_trigger":                                   types.ListType{ElemType: types.ObjectType{AttrTypes: DmMSDebugTriggerTypeObjectType}},
	"flow_control":                                    types.BoolType,
	"soap_schema_url":                                 types.StringType,
	"front_timeout":                                   types.Int64Type,
	"back_timeout":                                    types.Int64Type,
	"front_persistent_timeout":                        types.Int64Type,
	"back_persistent_timeout":                         types.Int64Type,
	"include_response_type_encoding":                  types.BoolType,
	"persistent_connections":                          types.BoolType,
	"loop_detection":                                  types.BoolType,
	"do_host_rewriting":                               types.BoolType,
	"do_chunked_upload":                               types.BoolType,
	"process_http_errors":                             types.BoolType,
	"http_client_ip_label":                            types.StringType,
	"http_log_cor_id_label":                           types.StringType,
	"load_balancer_hash_header":                       types.StringType,
	"in_order_mode":                                   types.ObjectType{AttrTypes: DmGatewayInOrderModeObjectType},
	"ws_a_mode":                                       types.StringType,
	"ws_a_require_aaa":                                types.BoolType,
	"ws_a_rewrite_reply_to":                           types.StringType,
	"ws_a_rewrite_fault_to":                           types.StringType,
	"ws_a_rewrite_to":                                 types.StringType,
	"ws_a_strip":                                      types.BoolType,
	"ws_a_default_reply_to":                           types.StringType,
	"ws_a_default_fault_to":                           types.StringType,
	"ws_a_force":                                      types.BoolType,
	"ws_a_gen_style":                                  types.StringType,
	"ws_a_http_async_response_code":                   types.Int64Type,
	"ws_a_back_protocol":                              types.StringType,
	"ws_a_timeout":                                    types.Int64Type,
	"ws_rm_enabled":                                   types.BoolType,
	"ws_rm_sequence_expiration":                       types.Int64Type,
	"ws_rm_aaa_policy":                                types.StringType,
	"ws_rm_destination_accept_create_sequence":        types.BoolType,
	"ws_rm_destination_maximum_sequences":             types.Int64Type,
	"ws_rm_destination_in_order":                      types.BoolType,
	"ws_rm_destination_maximum_in_order_queue_length": types.Int64Type,
	"ws_rm_destination_accept_offers":                 types.BoolType,
	"ws_rm_front_force":                               types.BoolType,
	"ws_rm_back_force":                                types.BoolType,
	"ws_rm_back_create_sequence":                      types.BoolType,
	"ws_rm_front_create_sequence":                     types.BoolType,
	"ws_rm_source_make_offer":                         types.BoolType,
	"ws_rm_uses_sequence_ssl":                         types.BoolType,
	"ws_rm_front_acks_to":                             types.StringType,
	"ws_rm_back_acks_to":                              types.StringType,
	"ws_rm_source_maximum_sequences":                  types.Int64Type,
	"ws_rm_source_retransmission_interval":            types.Int64Type,
	"ws_rm_source_exponential_backoff":                types.BoolType,
	"ws_rm_source_maximum_retransmissions":            types.Int64Type,
	"ws_rm_source_maximum_queue_length":               types.Int64Type,
	"ws_rm_source_request_ack_count":                  types.Int64Type,
	"ws_rm_source_inactivity_close":                   types.Int64Type,
	"force_policy_exec":                               types.BoolType,
	"rewrite_errors":                                  types.BoolType,
	"delay_errors":                                    types.BoolType,
	"delay_errors_duration":                           types.Int64Type,
	"dependency_actions":                              actions.ActionsListType,
}

func (data WSGateway) GetPath() string {
	rest_path := "/mgmt/config/{domain}/WSGateway"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data WSGateway) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.BackHttpVersion.IsNull() {
		return false
	}
	if !data.RequestType.IsNull() {
		return false
	}
	if !data.ResponseType.IsNull() {
		return false
	}
	if !data.FollowRedirects.IsNull() {
		return false
	}
	if !data.AllowCompression.IsNull() {
		return false
	}
	if !data.EnableCompressedPayloadPassthrough.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
		return false
	}
	if !data.AutoCreateSourceS.IsNull() {
		return false
	}
	if !data.SslServerConfigType.IsNull() {
		return false
	}
	if !data.SslServer.IsNull() {
		return false
	}
	if !data.SslSniServer.IsNull() {
		return false
	}
	if !data.EndpointRewritePolicy.IsNull() {
		return false
	}
	if !data.LocalEndpointRewrite.IsNull() {
		return false
	}
	if !data.RemoteEndpointRewrite.IsNull() {
		return false
	}
	if !data.AaaPolicy.IsNull() {
		return false
	}
	if !data.StylePolicy.IsNull() {
		return false
	}
	if data.RemoteFetchRetry != nil {
		if !data.RemoteFetchRetry.IsNull() {
			return false
		}
	}
	if !data.WsdlCachePolicy.IsNull() {
		return false
	}
	if !data.BaseWsdl.IsNull() {
		return false
	}
	if !data.UserToggles.IsNull() {
		return false
	}
	if !data.ClientPrincipal.IsNull() {
		return false
	}
	if !data.ServerPrincipal.IsNull() {
		return false
	}
	if !data.KerberosKeytab.IsNull() {
		return false
	}
	if !data.DecryptKey.IsNull() {
		return false
	}
	if !data.EncryptedKeySha1CacheLifeTime.IsNull() {
		return false
	}
	if !data.PreserveKeyChain.IsNull() {
		return false
	}
	if !data.DecryptWithKeyFromEd.IsNull() {
		return false
	}
	if !data.SoapActionPolicy.IsNull() {
		return false
	}
	if !data.WsrrSubscriptions.IsNull() {
		return false
	}
	if !data.WsrrSavedSearchSubscriptions.IsNull() {
		return false
	}
	if !data.OperationPriority.IsNull() {
		return false
	}
	if !data.OperationConformancePolicy.IsNull() {
		return false
	}
	if !data.OperationPolicySubjectOptOut.IsNull() {
		return false
	}
	if !data.PolicyParameter.IsNull() {
		return false
	}
	if !data.ReliableMessaging.IsNull() {
		return false
	}
	if !data.WSmAgentMonitor.IsNull() {
		return false
	}
	if !data.WSmAgentMonitorPcm.IsNull() {
		return false
	}
	if !data.ProcessRespRulesOnOneWayMep.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Priority.IsNull() {
		return false
	}
	if !data.FrontProtocol.IsNull() {
		return false
	}
	if !data.XmlManager.IsNull() {
		return false
	}
	if !data.UrlRewritePolicy.IsNull() {
		return false
	}
	if !data.SslClientConfigType.IsNull() {
		return false
	}
	if !data.SslClient.IsNull() {
		return false
	}
	if !data.FwCred.IsNull() {
		return false
	}
	if !data.HeaderInjection.IsNull() {
		return false
	}
	if !data.HeaderSuppression.IsNull() {
		return false
	}
	if !data.StylesheetParameters.IsNull() {
		return false
	}
	if !data.DefaultParamNamespace.IsNull() {
		return false
	}
	if !data.QueryParamNamespace.IsNull() {
		return false
	}
	if !data.BackendUrl.IsNull() {
		return false
	}
	if !data.PropagateUri.IsNull() {
		return false
	}
	if !data.ServiceMonitors.IsNull() {
		return false
	}
	if !data.CountMonitors.IsNull() {
		return false
	}
	if !data.DurationMonitors.IsNull() {
		return false
	}
	if !data.MonitorProcessingPolicy.IsNull() {
		return false
	}
	if !data.RequestAttachments.IsNull() {
		return false
	}
	if !data.ResponseAttachments.IsNull() {
		return false
	}
	if !data.RequestAttachmentsFlowControl.IsNull() {
		return false
	}
	if !data.ResponseAttachmentsFlowControl.IsNull() {
		return false
	}
	if !data.RootPartNotFirstAction.IsNull() {
		return false
	}
	if !data.FrontAttachmentFormat.IsNull() {
		return false
	}
	if !data.BackAttachmentFormat.IsNull() {
		return false
	}
	if !data.MimeFrontHeaders.IsNull() {
		return false
	}
	if !data.MimeBackHeaders.IsNull() {
		return false
	}
	if !data.StreamOutputToBack.IsNull() {
		return false
	}
	if !data.StreamOutputToFront.IsNull() {
		return false
	}
	if !data.MaxMessageSize.IsNull() {
		return false
	}
	if !data.GatewayParserLimits.IsNull() {
		return false
	}
	if !data.ParserLimitsElementDepth.IsNull() {
		return false
	}
	if !data.ParserLimitsAttributeCount.IsNull() {
		return false
	}
	if !data.ParserLimitsMaxNodeSize.IsNull() {
		return false
	}
	if !data.ParserLimitsExternalReferences.IsNull() {
		return false
	}
	if !data.ParserLimitsMaxPrefixes.IsNull() {
		return false
	}
	if !data.ParserLimitsMaxNamespaces.IsNull() {
		return false
	}
	if !data.ParserLimitsMaxLocalNames.IsNull() {
		return false
	}
	if !data.ParserLimitsAttachmentByteCount.IsNull() {
		return false
	}
	if !data.ParserLimitsAttachmentPackageByteCount.IsNull() {
		return false
	}
	if !data.DebugMode.IsNull() {
		return false
	}
	if !data.DebugHistory.IsNull() {
		return false
	}
	if !data.DebugTrigger.IsNull() {
		return false
	}
	if !data.FlowControl.IsNull() {
		return false
	}
	if !data.SoapSchemaUrl.IsNull() {
		return false
	}
	if !data.FrontTimeout.IsNull() {
		return false
	}
	if !data.BackTimeout.IsNull() {
		return false
	}
	if !data.FrontPersistentTimeout.IsNull() {
		return false
	}
	if !data.BackPersistentTimeout.IsNull() {
		return false
	}
	if !data.IncludeResponseTypeEncoding.IsNull() {
		return false
	}
	if !data.PersistentConnections.IsNull() {
		return false
	}
	if !data.LoopDetection.IsNull() {
		return false
	}
	if !data.DoHostRewriting.IsNull() {
		return false
	}
	if !data.DoChunkedUpload.IsNull() {
		return false
	}
	if !data.ProcessHttpErrors.IsNull() {
		return false
	}
	if !data.HttpClientIpLabel.IsNull() {
		return false
	}
	if !data.HttpLogCorIdLabel.IsNull() {
		return false
	}
	if !data.LoadBalancerHashHeader.IsNull() {
		return false
	}
	if data.InOrderMode != nil {
		if !data.InOrderMode.IsNull() {
			return false
		}
	}
	if !data.WsAMode.IsNull() {
		return false
	}
	if !data.WsARequireAaa.IsNull() {
		return false
	}
	if !data.WsARewriteReplyTo.IsNull() {
		return false
	}
	if !data.WsARewriteFaultTo.IsNull() {
		return false
	}
	if !data.WsARewriteTo.IsNull() {
		return false
	}
	if !data.WsAStrip.IsNull() {
		return false
	}
	if !data.WsADefaultReplyTo.IsNull() {
		return false
	}
	if !data.WsADefaultFaultTo.IsNull() {
		return false
	}
	if !data.WsAForce.IsNull() {
		return false
	}
	if !data.WsAGenStyle.IsNull() {
		return false
	}
	if !data.WsAHttpAsyncResponseCode.IsNull() {
		return false
	}
	if !data.WsABackProtocol.IsNull() {
		return false
	}
	if !data.WsATimeout.IsNull() {
		return false
	}
	if !data.WsRmEnabled.IsNull() {
		return false
	}
	if !data.WsRmSequenceExpiration.IsNull() {
		return false
	}
	if !data.WsRmAaaPolicy.IsNull() {
		return false
	}
	if !data.WsRmDestinationAcceptCreateSequence.IsNull() {
		return false
	}
	if !data.WsRmDestinationMaximumSequences.IsNull() {
		return false
	}
	if !data.WsRmDestinationInOrder.IsNull() {
		return false
	}
	if !data.WsRmDestinationMaximumInOrderQueueLength.IsNull() {
		return false
	}
	if !data.WsRmDestinationAcceptOffers.IsNull() {
		return false
	}
	if !data.WsRmFrontForce.IsNull() {
		return false
	}
	if !data.WsRmBackForce.IsNull() {
		return false
	}
	if !data.WsRmBackCreateSequence.IsNull() {
		return false
	}
	if !data.WsRmFrontCreateSequence.IsNull() {
		return false
	}
	if !data.WsRmSourceMakeOffer.IsNull() {
		return false
	}
	if !data.WsRmUsesSequenceSsl.IsNull() {
		return false
	}
	if !data.WsRmFrontAcksTo.IsNull() {
		return false
	}
	if !data.WsRmBackAcksTo.IsNull() {
		return false
	}
	if !data.WsRmSourceMaximumSequences.IsNull() {
		return false
	}
	if !data.WsRmSourceRetransmissionInterval.IsNull() {
		return false
	}
	if !data.WsRmSourceExponentialBackoff.IsNull() {
		return false
	}
	if !data.WsRmSourceMaximumRetransmissions.IsNull() {
		return false
	}
	if !data.WsRmSourceMaximumQueueLength.IsNull() {
		return false
	}
	if !data.WsRmSourceRequestAckCount.IsNull() {
		return false
	}
	if !data.WsRmSourceInactivityClose.IsNull() {
		return false
	}
	if !data.ForcePolicyExec.IsNull() {
		return false
	}
	if !data.RewriteErrors.IsNull() {
		return false
	}
	if !data.DelayErrors.IsNull() {
		return false
	}
	if !data.DelayErrorsDuration.IsNull() {
		return false
	}
	return true
}

func (data WSGateway) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.BackHttpVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BackHTTPVersion`, data.BackHttpVersion.ValueString())
	}
	if !data.RequestType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RequestType`, data.RequestType.ValueString())
	}
	if !data.ResponseType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ResponseType`, data.ResponseType.ValueString())
	}
	if !data.FollowRedirects.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FollowRedirects`, tfutils.StringFromBool(data.FollowRedirects, ""))
	}
	if !data.AllowCompression.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllowCompression`, tfutils.StringFromBool(data.AllowCompression, ""))
	}
	if !data.EnableCompressedPayloadPassthrough.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnableCompressedPayloadPassthrough`, tfutils.StringFromBool(data.EnableCompressedPayloadPassthrough, ""))
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Type`, data.Type.ValueString())
	}
	if !data.AutoCreateSourceS.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AutoCreateSources`, tfutils.StringFromBool(data.AutoCreateSourceS, ""))
	}
	if !data.SslServerConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLServerConfigType`, data.SslServerConfigType.ValueString())
	}
	if !data.SslServer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLServer`, data.SslServer.ValueString())
	}
	if !data.SslSniServer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLSNIServer`, data.SslSniServer.ValueString())
	}
	if !data.EndpointRewritePolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EndpointRewritePolicy`, data.EndpointRewritePolicy.ValueString())
	}
	if !data.LocalEndpointRewrite.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalEndpointRewrite`, data.LocalEndpointRewrite.ValueString())
	}
	if !data.RemoteEndpointRewrite.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemoteEndpointRewrite`, data.RemoteEndpointRewrite.ValueString())
	}
	if !data.AaaPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AAAPolicy`, data.AaaPolicy.ValueString())
	}
	if !data.StylePolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StylePolicy`, data.StylePolicy.ValueString())
	}
	if data.RemoteFetchRetry != nil {
		if !data.RemoteFetchRetry.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`RemoteFetchRetry`, data.RemoteFetchRetry.ToBody(ctx, ""))
		}
	}
	if !data.WsdlCachePolicy.IsNull() {
		var dataValues []DmWSDLCachePolicy
		data.WsdlCachePolicy.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`WSDLCachePolicy`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.BaseWsdl.IsNull() {
		var dataValues []DmWSBaseWSDL
		data.BaseWsdl.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`BaseWSDL`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.UserToggles.IsNull() {
		var dataValues []DmWSUserToggles
		data.UserToggles.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`UserToggles`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.ClientPrincipal.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ClientPrincipal`, data.ClientPrincipal.ValueString())
	}
	if !data.ServerPrincipal.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ServerPrincipal`, data.ServerPrincipal.ValueString())
	}
	if !data.KerberosKeytab.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`KerberosKeytab`, data.KerberosKeytab.ValueString())
	}
	if !data.DecryptKey.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DecryptKey`, data.DecryptKey.ValueString())
	}
	if !data.EncryptedKeySha1CacheLifeTime.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EncryptedKeySHA1CacheLifeTime`, data.EncryptedKeySha1CacheLifeTime.ValueInt64())
	}
	if !data.PreserveKeyChain.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PreserveKeyChain`, tfutils.StringFromBool(data.PreserveKeyChain, ""))
	}
	if !data.DecryptWithKeyFromEd.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DecryptWithKeyFromED`, tfutils.StringFromBool(data.DecryptWithKeyFromEd, ""))
	}
	if !data.SoapActionPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SOAPActionPolicy`, data.SoapActionPolicy.ValueString())
	}
	if !data.WsrrSubscriptions.IsNull() {
		var dataValues []DmWSRRWSDLSource
		data.WsrrSubscriptions.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`WSRRSubscriptions`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.WsrrSavedSearchSubscriptions.IsNull() {
		var dataValues []DmWSRRSavedSearchWSDLSource
		data.WsrrSavedSearchSubscriptions.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`WSRRSavedSearchSubscriptions`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.OperationPriority.IsNull() {
		var dataValues []DmWSOperationSchedulerPriority
		data.OperationPriority.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`OperationPriority`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.OperationConformancePolicy.IsNull() {
		var dataValues []DmWSOperationConformancePolicy
		data.OperationConformancePolicy.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`OperationConformancePolicy`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.OperationPolicySubjectOptOut.IsNull() {
		var dataValues []DmWSOperationPolicySubjectOptOut
		data.OperationPolicySubjectOptOut.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`OperationPolicySubjectOptOut`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.PolicyParameter.IsNull() {
		var dataValues []DmWSPolicyParameters
		data.PolicyParameter.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`PolicyParameter`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.ReliableMessaging.IsNull() {
		var dataValues []DmWSOperationReliableMessaging
		data.ReliableMessaging.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`ReliableMessaging`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.WSmAgentMonitor.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSMAgentMonitor`, tfutils.StringFromBool(data.WSmAgentMonitor, ""))
	}
	if !data.WSmAgentMonitorPcm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSMAgentMonitorPCM`, data.WSmAgentMonitorPcm.ValueString())
	}
	if !data.ProcessRespRulesOnOneWayMep.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ProcessRespRulesOnOneWayMEP`, tfutils.StringFromBool(data.ProcessRespRulesOnOneWayMep, ""))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.Priority.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Priority`, data.Priority.ValueString())
	}
	if !data.FrontProtocol.IsNull() {
		var dataValues []string
		data.FrontProtocol.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`FrontProtocol`+".-1", map[string]string{"value": val})
		}
	}
	if !data.XmlManager.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XMLManager`, data.XmlManager.ValueString())
	}
	if !data.UrlRewritePolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`URLRewritePolicy`, data.UrlRewritePolicy.ValueString())
	}
	if !data.SslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClientConfigType`, data.SslClientConfigType.ValueString())
	}
	if !data.SslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClient`, data.SslClient.ValueString())
	}
	if !data.FwCred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FWCred`, data.FwCred.ValueString())
	}
	if !data.HeaderInjection.IsNull() {
		var dataValues []DmHeaderInjection
		data.HeaderInjection.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`HeaderInjection`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.HeaderSuppression.IsNull() {
		var dataValues []DmHeaderSuppression
		data.HeaderSuppression.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`HeaderSuppression`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.StylesheetParameters.IsNull() {
		var dataValues []DmStylesheetParameter
		data.StylesheetParameters.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`StylesheetParameters`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.DefaultParamNamespace.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DefaultParamNamespace`, data.DefaultParamNamespace.ValueString())
	}
	if !data.QueryParamNamespace.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`QueryParamNamespace`, data.QueryParamNamespace.ValueString())
	}
	if !data.BackendUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BackendUrl`, data.BackendUrl.ValueString())
	}
	if !data.PropagateUri.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PropagateURI`, tfutils.StringFromBool(data.PropagateUri, ""))
	}
	if !data.ServiceMonitors.IsNull() {
		var dataValues []string
		data.ServiceMonitors.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`ServiceMonitors`+".-1", map[string]string{"value": val})
		}
	}
	if !data.CountMonitors.IsNull() {
		var dataValues []string
		data.CountMonitors.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`CountMonitors`+".-1", map[string]string{"value": val})
		}
	}
	if !data.DurationMonitors.IsNull() {
		var dataValues []string
		data.DurationMonitors.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`DurationMonitors`+".-1", map[string]string{"value": val})
		}
	}
	if !data.MonitorProcessingPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MonitorProcessingPolicy`, data.MonitorProcessingPolicy.ValueString())
	}
	if !data.RequestAttachments.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RequestAttachments`, data.RequestAttachments.ValueString())
	}
	if !data.ResponseAttachments.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ResponseAttachments`, data.ResponseAttachments.ValueString())
	}
	if !data.RequestAttachmentsFlowControl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RequestAttachmentsFlowControl`, tfutils.StringFromBool(data.RequestAttachmentsFlowControl, ""))
	}
	if !data.ResponseAttachmentsFlowControl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ResponseAttachmentsFlowControl`, tfutils.StringFromBool(data.ResponseAttachmentsFlowControl, ""))
	}
	if !data.RootPartNotFirstAction.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RootPartNotFirstAction`, data.RootPartNotFirstAction.ValueString())
	}
	if !data.FrontAttachmentFormat.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FrontAttachmentFormat`, data.FrontAttachmentFormat.ValueString())
	}
	if !data.BackAttachmentFormat.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BackAttachmentFormat`, data.BackAttachmentFormat.ValueString())
	}
	if !data.MimeFrontHeaders.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MIMEFrontHeaders`, tfutils.StringFromBool(data.MimeFrontHeaders, ""))
	}
	if !data.MimeBackHeaders.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MIMEBackHeaders`, tfutils.StringFromBool(data.MimeBackHeaders, ""))
	}
	if !data.StreamOutputToBack.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StreamOutputToBack`, data.StreamOutputToBack.ValueString())
	}
	if !data.StreamOutputToFront.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StreamOutputToFront`, data.StreamOutputToFront.ValueString())
	}
	if !data.MaxMessageSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxMessageSize`, data.MaxMessageSize.ValueInt64())
	}
	if !data.GatewayParserLimits.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GatewayParserLimits`, tfutils.StringFromBool(data.GatewayParserLimits, ""))
	}
	if !data.ParserLimitsElementDepth.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ParserLimitsElementDepth`, data.ParserLimitsElementDepth.ValueInt64())
	}
	if !data.ParserLimitsAttributeCount.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ParserLimitsAttributeCount`, data.ParserLimitsAttributeCount.ValueInt64())
	}
	if !data.ParserLimitsMaxNodeSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ParserLimitsMaxNodeSize`, data.ParserLimitsMaxNodeSize.ValueInt64())
	}
	if !data.ParserLimitsExternalReferences.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ParserLimitsExternalReferences`, data.ParserLimitsExternalReferences.ValueString())
	}
	if !data.ParserLimitsMaxPrefixes.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ParserLimitsMaxPrefixes`, data.ParserLimitsMaxPrefixes.ValueInt64())
	}
	if !data.ParserLimitsMaxNamespaces.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ParserLimitsMaxNamespaces`, data.ParserLimitsMaxNamespaces.ValueInt64())
	}
	if !data.ParserLimitsMaxLocalNames.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ParserLimitsMaxLocalNames`, data.ParserLimitsMaxLocalNames.ValueInt64())
	}
	if !data.ParserLimitsAttachmentByteCount.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ParserLimitsAttachmentByteCount`, data.ParserLimitsAttachmentByteCount.ValueInt64())
	}
	if !data.ParserLimitsAttachmentPackageByteCount.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ParserLimitsAttachmentPackageByteCount`, data.ParserLimitsAttachmentPackageByteCount.ValueInt64())
	}
	if !data.DebugMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DebugMode`, data.DebugMode.ValueString())
	}
	if !data.DebugHistory.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DebugHistory`, data.DebugHistory.ValueInt64())
	}
	if !data.DebugTrigger.IsNull() {
		var dataValues []DmMSDebugTriggerType
		data.DebugTrigger.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`DebugTrigger`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.FlowControl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FlowControl`, tfutils.StringFromBool(data.FlowControl, ""))
	}
	if !data.SoapSchemaUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SOAPSchemaURL`, data.SoapSchemaUrl.ValueString())
	}
	if !data.FrontTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FrontTimeout`, data.FrontTimeout.ValueInt64())
	}
	if !data.BackTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BackTimeout`, data.BackTimeout.ValueInt64())
	}
	if !data.FrontPersistentTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FrontPersistentTimeout`, data.FrontPersistentTimeout.ValueInt64())
	}
	if !data.BackPersistentTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BackPersistentTimeout`, data.BackPersistentTimeout.ValueInt64())
	}
	if !data.IncludeResponseTypeEncoding.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IncludeResponseTypeEncoding`, tfutils.StringFromBool(data.IncludeResponseTypeEncoding, ""))
	}
	if !data.PersistentConnections.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PersistentConnections`, tfutils.StringFromBool(data.PersistentConnections, ""))
	}
	if !data.LoopDetection.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LoopDetection`, tfutils.StringFromBool(data.LoopDetection, ""))
	}
	if !data.DoHostRewriting.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DoHostRewriting`, tfutils.StringFromBool(data.DoHostRewriting, ""))
	}
	if !data.DoChunkedUpload.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DoChunkedUpload`, tfutils.StringFromBool(data.DoChunkedUpload, ""))
	}
	if !data.ProcessHttpErrors.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ProcessHTTPErrors`, tfutils.StringFromBool(data.ProcessHttpErrors, ""))
	}
	if !data.HttpClientIpLabel.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTPClientIPLabel`, data.HttpClientIpLabel.ValueString())
	}
	if !data.HttpLogCorIdLabel.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTPLogCorIDLabel`, data.HttpLogCorIdLabel.ValueString())
	}
	if !data.LoadBalancerHashHeader.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LoadBalancerHashHeader`, data.LoadBalancerHashHeader.ValueString())
	}
	if data.InOrderMode != nil {
		if !data.InOrderMode.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`InOrderMode`, data.InOrderMode.ToBody(ctx, ""))
		}
	}
	if !data.WsAMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSAMode`, data.WsAMode.ValueString())
	}
	if !data.WsARequireAaa.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSARequireAAA`, tfutils.StringFromBool(data.WsARequireAaa, ""))
	}
	if !data.WsARewriteReplyTo.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSARewriteReplyTo`, data.WsARewriteReplyTo.ValueString())
	}
	if !data.WsARewriteFaultTo.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSARewriteFaultTo`, data.WsARewriteFaultTo.ValueString())
	}
	if !data.WsARewriteTo.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSARewriteTo`, data.WsARewriteTo.ValueString())
	}
	if !data.WsAStrip.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSAStrip`, tfutils.StringFromBool(data.WsAStrip, ""))
	}
	if !data.WsADefaultReplyTo.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSADefaultReplyTo`, data.WsADefaultReplyTo.ValueString())
	}
	if !data.WsADefaultFaultTo.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSADefaultFaultTo`, data.WsADefaultFaultTo.ValueString())
	}
	if !data.WsAForce.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSAForce`, tfutils.StringFromBool(data.WsAForce, ""))
	}
	if !data.WsAGenStyle.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSAGenStyle`, data.WsAGenStyle.ValueString())
	}
	if !data.WsAHttpAsyncResponseCode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSAHTTPAsyncResponseCode`, data.WsAHttpAsyncResponseCode.ValueInt64())
	}
	if !data.WsABackProtocol.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSABackProtocol`, data.WsABackProtocol.ValueString())
	}
	if !data.WsATimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSATimeout`, data.WsATimeout.ValueInt64())
	}
	if !data.WsRmEnabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSRMEnabled`, tfutils.StringFromBool(data.WsRmEnabled, ""))
	}
	if !data.WsRmSequenceExpiration.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSRMSequenceExpiration`, data.WsRmSequenceExpiration.ValueInt64())
	}
	if !data.WsRmAaaPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSRMAAAPolicy`, data.WsRmAaaPolicy.ValueString())
	}
	if !data.WsRmDestinationAcceptCreateSequence.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSRMDestinationAcceptCreateSequence`, tfutils.StringFromBool(data.WsRmDestinationAcceptCreateSequence, ""))
	}
	if !data.WsRmDestinationMaximumSequences.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSRMDestinationMaximumSequences`, data.WsRmDestinationMaximumSequences.ValueInt64())
	}
	if !data.WsRmDestinationInOrder.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSRMDestinationInOrder`, tfutils.StringFromBool(data.WsRmDestinationInOrder, ""))
	}
	if !data.WsRmDestinationMaximumInOrderQueueLength.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSRMDestinationMaximumInOrderQueueLength`, data.WsRmDestinationMaximumInOrderQueueLength.ValueInt64())
	}
	if !data.WsRmDestinationAcceptOffers.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSRMDestinationAcceptOffers`, tfutils.StringFromBool(data.WsRmDestinationAcceptOffers, ""))
	}
	if !data.WsRmFrontForce.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSRMFrontForce`, tfutils.StringFromBool(data.WsRmFrontForce, ""))
	}
	if !data.WsRmBackForce.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSRMBackForce`, tfutils.StringFromBool(data.WsRmBackForce, ""))
	}
	if !data.WsRmBackCreateSequence.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSRMBackCreateSequence`, tfutils.StringFromBool(data.WsRmBackCreateSequence, ""))
	}
	if !data.WsRmFrontCreateSequence.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSRMFrontCreateSequence`, tfutils.StringFromBool(data.WsRmFrontCreateSequence, ""))
	}
	if !data.WsRmSourceMakeOffer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSRMSourceMakeOffer`, tfutils.StringFromBool(data.WsRmSourceMakeOffer, ""))
	}
	if !data.WsRmUsesSequenceSsl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSRMUsesSequenceSSL`, tfutils.StringFromBool(data.WsRmUsesSequenceSsl, ""))
	}
	if !data.WsRmFrontAcksTo.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSRMFrontAcksTo`, data.WsRmFrontAcksTo.ValueString())
	}
	if !data.WsRmBackAcksTo.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSRMBackAcksTo`, data.WsRmBackAcksTo.ValueString())
	}
	if !data.WsRmSourceMaximumSequences.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSRMSourceMaximumSequences`, data.WsRmSourceMaximumSequences.ValueInt64())
	}
	if !data.WsRmSourceRetransmissionInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSRMSourceRetransmissionInterval`, data.WsRmSourceRetransmissionInterval.ValueInt64())
	}
	if !data.WsRmSourceExponentialBackoff.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSRMSourceExponentialBackoff`, tfutils.StringFromBool(data.WsRmSourceExponentialBackoff, ""))
	}
	if !data.WsRmSourceMaximumRetransmissions.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSRMSourceMaximumRetransmissions`, data.WsRmSourceMaximumRetransmissions.ValueInt64())
	}
	if !data.WsRmSourceMaximumQueueLength.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSRMSourceMaximumQueueLength`, data.WsRmSourceMaximumQueueLength.ValueInt64())
	}
	if !data.WsRmSourceRequestAckCount.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSRMSourceRequestAckCount`, data.WsRmSourceRequestAckCount.ValueInt64())
	}
	if !data.WsRmSourceInactivityClose.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSRMSourceInactivityClose`, data.WsRmSourceInactivityClose.ValueInt64())
	}
	if !data.ForcePolicyExec.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ForcePolicyExec`, tfutils.StringFromBool(data.ForcePolicyExec, ""))
	}
	if !data.RewriteErrors.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RewriteErrors`, tfutils.StringFromBool(data.RewriteErrors, ""))
	}
	if !data.DelayErrors.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DelayErrors`, tfutils.StringFromBool(data.DelayErrors, ""))
	}
	if !data.DelayErrorsDuration.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DelayErrorsDuration`, data.DelayErrorsDuration.ValueInt64())
	}
	return body
}

func (data *WSGateway) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `BackHTTPVersion`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.BackHttpVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BackHttpVersion = types.StringValue("HTTP/1.1")
	}
	if value := res.Get(pathRoot + `RequestType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RequestType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RequestType = types.StringValue("soap")
	}
	if value := res.Get(pathRoot + `ResponseType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ResponseType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ResponseType = types.StringValue("soap")
	}
	if value := res.Get(pathRoot + `FollowRedirects`); value.Exists() {
		data.FollowRedirects = tfutils.BoolFromString(value.String())
	} else {
		data.FollowRedirects = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowCompression`); value.Exists() {
		data.AllowCompression = tfutils.BoolFromString(value.String())
	} else {
		data.AllowCompression = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EnableCompressedPayloadPassthrough`); value.Exists() {
		data.EnableCompressedPayloadPassthrough = tfutils.BoolFromString(value.String())
	} else {
		data.EnableCompressedPayloadPassthrough = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringValue("static-from-wsdl")
	}
	if value := res.Get(pathRoot + `AutoCreateSources`); value.Exists() {
		data.AutoCreateSourceS = tfutils.BoolFromString(value.String())
	} else {
		data.AutoCreateSourceS = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SSLServerConfigType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslServerConfigType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslServerConfigType = types.StringValue("server")
	}
	if value := res.Get(pathRoot + `SSLServer`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLSNIServer`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslSniServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslSniServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `EndpointRewritePolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EndpointRewritePolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EndpointRewritePolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalEndpointRewrite`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalEndpointRewrite = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalEndpointRewrite = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemoteEndpointRewrite`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RemoteEndpointRewrite = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteEndpointRewrite = types.StringNull()
	}
	if value := res.Get(pathRoot + `AAAPolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AaaPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AaaPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `StylePolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.StylePolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.StylePolicy = types.StringValue("default")
	}
	if value := res.Get(pathRoot + `RemoteFetchRetry`); value.Exists() {
		data.RemoteFetchRetry = &DmNetworkRetry{}
		data.RemoteFetchRetry.FromBody(ctx, "", value)
	} else {
		data.RemoteFetchRetry = nil
	}
	if value := res.Get(pathRoot + `WSDLCachePolicy`); value.Exists() {
		l := []DmWSDLCachePolicy{}
		if value := res.Get(`WSDLCachePolicy`); value.Exists() {
			for _, v := range value.Array() {
				item := DmWSDLCachePolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.WsdlCachePolicy, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSDLCachePolicyObjectType}, l)
		} else {
			data.WsdlCachePolicy = types.ListNull(types.ObjectType{AttrTypes: DmWSDLCachePolicyObjectType})
		}
	} else {
		data.WsdlCachePolicy = types.ListNull(types.ObjectType{AttrTypes: DmWSDLCachePolicyObjectType})
	}
	if value := res.Get(pathRoot + `BaseWSDL`); value.Exists() {
		l := []DmWSBaseWSDL{}
		if value := res.Get(`BaseWSDL`); value.Exists() {
			for _, v := range value.Array() {
				item := DmWSBaseWSDL{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.BaseWsdl, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSBaseWSDLObjectType}, l)
		} else {
			data.BaseWsdl = types.ListNull(types.ObjectType{AttrTypes: DmWSBaseWSDLObjectType})
		}
	} else {
		data.BaseWsdl = types.ListNull(types.ObjectType{AttrTypes: DmWSBaseWSDLObjectType})
	}
	if value := res.Get(pathRoot + `UserToggles`); value.Exists() {
		l := []DmWSUserToggles{}
		if value := res.Get(`UserToggles`); value.Exists() {
			for _, v := range value.Array() {
				item := DmWSUserToggles{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.UserToggles, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSUserTogglesObjectType}, l)
		} else {
			data.UserToggles = types.ListNull(types.ObjectType{AttrTypes: DmWSUserTogglesObjectType})
		}
	} else {
		data.UserToggles = types.ListNull(types.ObjectType{AttrTypes: DmWSUserTogglesObjectType})
	}
	if value := res.Get(pathRoot + `ClientPrincipal`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ClientPrincipal = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientPrincipal = types.StringNull()
	}
	if value := res.Get(pathRoot + `ServerPrincipal`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ServerPrincipal = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ServerPrincipal = types.StringNull()
	}
	if value := res.Get(pathRoot + `KerberosKeytab`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.KerberosKeytab = tfutils.ParseStringFromGJSON(value)
	} else {
		data.KerberosKeytab = types.StringNull()
	}
	if value := res.Get(pathRoot + `DecryptKey`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DecryptKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DecryptKey = types.StringNull()
	}
	if value := res.Get(pathRoot + `EncryptedKeySHA1CacheLifeTime`); value.Exists() {
		data.EncryptedKeySha1CacheLifeTime = types.Int64Value(value.Int())
	} else {
		data.EncryptedKeySha1CacheLifeTime = types.Int64Null()
	}
	if value := res.Get(pathRoot + `PreserveKeyChain`); value.Exists() {
		data.PreserveKeyChain = tfutils.BoolFromString(value.String())
	} else {
		data.PreserveKeyChain = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DecryptWithKeyFromED`); value.Exists() {
		data.DecryptWithKeyFromEd = tfutils.BoolFromString(value.String())
	} else {
		data.DecryptWithKeyFromEd = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SOAPActionPolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SoapActionPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SoapActionPolicy = types.StringValue("lax")
	}
	if value := res.Get(pathRoot + `WSRRSubscriptions`); value.Exists() {
		l := []DmWSRRWSDLSource{}
		if value := res.Get(`WSRRSubscriptions`); value.Exists() {
			for _, v := range value.Array() {
				item := DmWSRRWSDLSource{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.WsrrSubscriptions, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSRRWSDLSourceObjectType}, l)
		} else {
			data.WsrrSubscriptions = types.ListNull(types.ObjectType{AttrTypes: DmWSRRWSDLSourceObjectType})
		}
	} else {
		data.WsrrSubscriptions = types.ListNull(types.ObjectType{AttrTypes: DmWSRRWSDLSourceObjectType})
	}
	if value := res.Get(pathRoot + `WSRRSavedSearchSubscriptions`); value.Exists() {
		l := []DmWSRRSavedSearchWSDLSource{}
		if value := res.Get(`WSRRSavedSearchSubscriptions`); value.Exists() {
			for _, v := range value.Array() {
				item := DmWSRRSavedSearchWSDLSource{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.WsrrSavedSearchSubscriptions, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSRRSavedSearchWSDLSourceObjectType}, l)
		} else {
			data.WsrrSavedSearchSubscriptions = types.ListNull(types.ObjectType{AttrTypes: DmWSRRSavedSearchWSDLSourceObjectType})
		}
	} else {
		data.WsrrSavedSearchSubscriptions = types.ListNull(types.ObjectType{AttrTypes: DmWSRRSavedSearchWSDLSourceObjectType})
	}
	if value := res.Get(pathRoot + `OperationPriority`); value.Exists() {
		l := []DmWSOperationSchedulerPriority{}
		if value := res.Get(`OperationPriority`); value.Exists() {
			for _, v := range value.Array() {
				item := DmWSOperationSchedulerPriority{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.OperationPriority, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSOperationSchedulerPriorityObjectType}, l)
		} else {
			data.OperationPriority = types.ListNull(types.ObjectType{AttrTypes: DmWSOperationSchedulerPriorityObjectType})
		}
	} else {
		data.OperationPriority = types.ListNull(types.ObjectType{AttrTypes: DmWSOperationSchedulerPriorityObjectType})
	}
	if value := res.Get(pathRoot + `OperationConformancePolicy`); value.Exists() {
		l := []DmWSOperationConformancePolicy{}
		if value := res.Get(`OperationConformancePolicy`); value.Exists() {
			for _, v := range value.Array() {
				item := DmWSOperationConformancePolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.OperationConformancePolicy, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSOperationConformancePolicyObjectType}, l)
		} else {
			data.OperationConformancePolicy = types.ListNull(types.ObjectType{AttrTypes: DmWSOperationConformancePolicyObjectType})
		}
	} else {
		data.OperationConformancePolicy = types.ListNull(types.ObjectType{AttrTypes: DmWSOperationConformancePolicyObjectType})
	}
	if value := res.Get(pathRoot + `OperationPolicySubjectOptOut`); value.Exists() {
		l := []DmWSOperationPolicySubjectOptOut{}
		if value := res.Get(`OperationPolicySubjectOptOut`); value.Exists() {
			for _, v := range value.Array() {
				item := DmWSOperationPolicySubjectOptOut{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.OperationPolicySubjectOptOut, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSOperationPolicySubjectOptOutObjectType}, l)
		} else {
			data.OperationPolicySubjectOptOut = types.ListNull(types.ObjectType{AttrTypes: DmWSOperationPolicySubjectOptOutObjectType})
		}
	} else {
		data.OperationPolicySubjectOptOut = types.ListNull(types.ObjectType{AttrTypes: DmWSOperationPolicySubjectOptOutObjectType})
	}
	if value := res.Get(pathRoot + `PolicyParameter`); value.Exists() {
		l := []DmWSPolicyParameters{}
		if value := res.Get(`PolicyParameter`); value.Exists() {
			for _, v := range value.Array() {
				item := DmWSPolicyParameters{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.PolicyParameter, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSPolicyParametersObjectType}, l)
		} else {
			data.PolicyParameter = types.ListNull(types.ObjectType{AttrTypes: DmWSPolicyParametersObjectType})
		}
	} else {
		data.PolicyParameter = types.ListNull(types.ObjectType{AttrTypes: DmWSPolicyParametersObjectType})
	}
	if value := res.Get(pathRoot + `ReliableMessaging`); value.Exists() {
		l := []DmWSOperationReliableMessaging{}
		if value := res.Get(`ReliableMessaging`); value.Exists() {
			for _, v := range value.Array() {
				item := DmWSOperationReliableMessaging{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.ReliableMessaging, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSOperationReliableMessagingObjectType}, l)
		} else {
			data.ReliableMessaging = types.ListNull(types.ObjectType{AttrTypes: DmWSOperationReliableMessagingObjectType})
		}
	} else {
		data.ReliableMessaging = types.ListNull(types.ObjectType{AttrTypes: DmWSOperationReliableMessagingObjectType})
	}
	if value := res.Get(pathRoot + `WSMAgentMonitor`); value.Exists() {
		data.WSmAgentMonitor = tfutils.BoolFromString(value.String())
	} else {
		data.WSmAgentMonitor = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSMAgentMonitorPCM`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WSmAgentMonitorPcm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WSmAgentMonitorPcm = types.StringValue("all-messages")
	}
	if value := res.Get(pathRoot + `ProcessRespRulesOnOneWayMEP`); value.Exists() {
		data.ProcessRespRulesOnOneWayMep = tfutils.BoolFromString(value.String())
	} else {
		data.ProcessRespRulesOnOneWayMep = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `Priority`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Priority = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Priority = types.StringValue("normal")
	}
	if value := res.Get(pathRoot + `FrontProtocol`); value.Exists() {
		data.FrontProtocol = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.FrontProtocol = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `XMLManager`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.XmlManager = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XmlManager = types.StringValue("default")
	}
	if value := res.Get(pathRoot + `URLRewritePolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UrlRewritePolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UrlRewritePolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLClientConfigType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClientConfigType = types.StringValue("client")
	}
	if value := res.Get(pathRoot + `SSLClient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `FWCred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.FwCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FwCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `HeaderInjection`); value.Exists() {
		l := []DmHeaderInjection{}
		if value := res.Get(`HeaderInjection`); value.Exists() {
			for _, v := range value.Array() {
				item := DmHeaderInjection{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.HeaderInjection, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmHeaderInjectionObjectType}, l)
		} else {
			data.HeaderInjection = types.ListNull(types.ObjectType{AttrTypes: DmHeaderInjectionObjectType})
		}
	} else {
		data.HeaderInjection = types.ListNull(types.ObjectType{AttrTypes: DmHeaderInjectionObjectType})
	}
	if value := res.Get(pathRoot + `HeaderSuppression`); value.Exists() {
		l := []DmHeaderSuppression{}
		if value := res.Get(`HeaderSuppression`); value.Exists() {
			for _, v := range value.Array() {
				item := DmHeaderSuppression{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.HeaderSuppression, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmHeaderSuppressionObjectType}, l)
		} else {
			data.HeaderSuppression = types.ListNull(types.ObjectType{AttrTypes: DmHeaderSuppressionObjectType})
		}
	} else {
		data.HeaderSuppression = types.ListNull(types.ObjectType{AttrTypes: DmHeaderSuppressionObjectType})
	}
	if value := res.Get(pathRoot + `StylesheetParameters`); value.Exists() {
		l := []DmStylesheetParameter{}
		if value := res.Get(`StylesheetParameters`); value.Exists() {
			for _, v := range value.Array() {
				item := DmStylesheetParameter{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.StylesheetParameters, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmStylesheetParameterObjectType}, l)
		} else {
			data.StylesheetParameters = types.ListNull(types.ObjectType{AttrTypes: DmStylesheetParameterObjectType})
		}
	} else {
		data.StylesheetParameters = types.ListNull(types.ObjectType{AttrTypes: DmStylesheetParameterObjectType})
	}
	if value := res.Get(pathRoot + `DefaultParamNamespace`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DefaultParamNamespace = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DefaultParamNamespace = types.StringValue("http://www.datapower.com/param/config")
	}
	if value := res.Get(pathRoot + `QueryParamNamespace`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.QueryParamNamespace = tfutils.ParseStringFromGJSON(value)
	} else {
		data.QueryParamNamespace = types.StringValue("http://www.datapower.com/param/query")
	}
	if value := res.Get(pathRoot + `BackendUrl`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.BackendUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BackendUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `PropagateURI`); value.Exists() {
		data.PropagateUri = tfutils.BoolFromString(value.String())
	} else {
		data.PropagateUri = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ServiceMonitors`); value.Exists() {
		data.ServiceMonitors = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.ServiceMonitors = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `CountMonitors`); value.Exists() {
		data.CountMonitors = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.CountMonitors = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `DurationMonitors`); value.Exists() {
		data.DurationMonitors = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.DurationMonitors = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `MonitorProcessingPolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MonitorProcessingPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MonitorProcessingPolicy = types.StringValue("terminate-at-first-throttle")
	}
	if value := res.Get(pathRoot + `RequestAttachments`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RequestAttachments = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RequestAttachments = types.StringValue("strip")
	}
	if value := res.Get(pathRoot + `ResponseAttachments`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ResponseAttachments = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ResponseAttachments = types.StringValue("strip")
	}
	if value := res.Get(pathRoot + `RequestAttachmentsFlowControl`); value.Exists() {
		data.RequestAttachmentsFlowControl = tfutils.BoolFromString(value.String())
	} else {
		data.RequestAttachmentsFlowControl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ResponseAttachmentsFlowControl`); value.Exists() {
		data.ResponseAttachmentsFlowControl = tfutils.BoolFromString(value.String())
	} else {
		data.ResponseAttachmentsFlowControl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RootPartNotFirstAction`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RootPartNotFirstAction = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RootPartNotFirstAction = types.StringValue("process-in-order")
	}
	if value := res.Get(pathRoot + `FrontAttachmentFormat`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.FrontAttachmentFormat = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FrontAttachmentFormat = types.StringValue("dynamic")
	}
	if value := res.Get(pathRoot + `BackAttachmentFormat`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.BackAttachmentFormat = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BackAttachmentFormat = types.StringValue("dynamic")
	}
	if value := res.Get(pathRoot + `MIMEFrontHeaders`); value.Exists() {
		data.MimeFrontHeaders = tfutils.BoolFromString(value.String())
	} else {
		data.MimeFrontHeaders = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MIMEBackHeaders`); value.Exists() {
		data.MimeBackHeaders = tfutils.BoolFromString(value.String())
	} else {
		data.MimeBackHeaders = types.BoolNull()
	}
	if value := res.Get(pathRoot + `StreamOutputToBack`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.StreamOutputToBack = tfutils.ParseStringFromGJSON(value)
	} else {
		data.StreamOutputToBack = types.StringValue("buffer-until-verification")
	}
	if value := res.Get(pathRoot + `StreamOutputToFront`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.StreamOutputToFront = tfutils.ParseStringFromGJSON(value)
	} else {
		data.StreamOutputToFront = types.StringValue("buffer-until-verification")
	}
	if value := res.Get(pathRoot + `MaxMessageSize`); value.Exists() {
		data.MaxMessageSize = types.Int64Value(value.Int())
	} else {
		data.MaxMessageSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `GatewayParserLimits`); value.Exists() {
		data.GatewayParserLimits = tfutils.BoolFromString(value.String())
	} else {
		data.GatewayParserLimits = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ParserLimitsElementDepth`); value.Exists() {
		data.ParserLimitsElementDepth = types.Int64Value(value.Int())
	} else {
		data.ParserLimitsElementDepth = types.Int64Value(512)
	}
	if value := res.Get(pathRoot + `ParserLimitsAttributeCount`); value.Exists() {
		data.ParserLimitsAttributeCount = types.Int64Value(value.Int())
	} else {
		data.ParserLimitsAttributeCount = types.Int64Value(128)
	}
	if value := res.Get(pathRoot + `ParserLimitsMaxNodeSize`); value.Exists() {
		data.ParserLimitsMaxNodeSize = types.Int64Value(value.Int())
	} else {
		data.ParserLimitsMaxNodeSize = types.Int64Value(33554432)
	}
	if value := res.Get(pathRoot + `ParserLimitsExternalReferences`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ParserLimitsExternalReferences = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ParserLimitsExternalReferences = types.StringValue("forbid")
	}
	if value := res.Get(pathRoot + `ParserLimitsMaxPrefixes`); value.Exists() {
		data.ParserLimitsMaxPrefixes = types.Int64Value(value.Int())
	} else {
		data.ParserLimitsMaxPrefixes = types.Int64Value(1024)
	}
	if value := res.Get(pathRoot + `ParserLimitsMaxNamespaces`); value.Exists() {
		data.ParserLimitsMaxNamespaces = types.Int64Value(value.Int())
	} else {
		data.ParserLimitsMaxNamespaces = types.Int64Value(1024)
	}
	if value := res.Get(pathRoot + `ParserLimitsMaxLocalNames`); value.Exists() {
		data.ParserLimitsMaxLocalNames = types.Int64Value(value.Int())
	} else {
		data.ParserLimitsMaxLocalNames = types.Int64Value(60000)
	}
	if value := res.Get(pathRoot + `ParserLimitsAttachmentByteCount`); value.Exists() {
		data.ParserLimitsAttachmentByteCount = types.Int64Value(value.Int())
	} else {
		data.ParserLimitsAttachmentByteCount = types.Int64Value(2000000000)
	}
	if value := res.Get(pathRoot + `ParserLimitsAttachmentPackageByteCount`); value.Exists() {
		data.ParserLimitsAttachmentPackageByteCount = types.Int64Value(value.Int())
	} else {
		data.ParserLimitsAttachmentPackageByteCount = types.Int64Null()
	}
	if value := res.Get(pathRoot + `DebugMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DebugMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DebugMode = types.StringValue("off")
	}
	if value := res.Get(pathRoot + `DebugHistory`); value.Exists() {
		data.DebugHistory = types.Int64Value(value.Int())
	} else {
		data.DebugHistory = types.Int64Value(25)
	}
	if value := res.Get(pathRoot + `DebugTrigger`); value.Exists() {
		l := []DmMSDebugTriggerType{}
		if value := res.Get(`DebugTrigger`); value.Exists() {
			for _, v := range value.Array() {
				item := DmMSDebugTriggerType{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.DebugTrigger, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmMSDebugTriggerTypeObjectType}, l)
		} else {
			data.DebugTrigger = types.ListNull(types.ObjectType{AttrTypes: DmMSDebugTriggerTypeObjectType})
		}
	} else {
		data.DebugTrigger = types.ListNull(types.ObjectType{AttrTypes: DmMSDebugTriggerTypeObjectType})
	}
	if value := res.Get(pathRoot + `FlowControl`); value.Exists() {
		data.FlowControl = tfutils.BoolFromString(value.String())
	} else {
		data.FlowControl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SOAPSchemaURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SoapSchemaUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SoapSchemaUrl = types.StringValue("store:///schemas/soap-envelope.xsd")
	}
	if value := res.Get(pathRoot + `FrontTimeout`); value.Exists() {
		data.FrontTimeout = types.Int64Value(value.Int())
	} else {
		data.FrontTimeout = types.Int64Value(120)
	}
	if value := res.Get(pathRoot + `BackTimeout`); value.Exists() {
		data.BackTimeout = types.Int64Value(value.Int())
	} else {
		data.BackTimeout = types.Int64Value(120)
	}
	if value := res.Get(pathRoot + `FrontPersistentTimeout`); value.Exists() {
		data.FrontPersistentTimeout = types.Int64Value(value.Int())
	} else {
		data.FrontPersistentTimeout = types.Int64Value(180)
	}
	if value := res.Get(pathRoot + `BackPersistentTimeout`); value.Exists() {
		data.BackPersistentTimeout = types.Int64Value(value.Int())
	} else {
		data.BackPersistentTimeout = types.Int64Value(180)
	}
	if value := res.Get(pathRoot + `IncludeResponseTypeEncoding`); value.Exists() {
		data.IncludeResponseTypeEncoding = tfutils.BoolFromString(value.String())
	} else {
		data.IncludeResponseTypeEncoding = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PersistentConnections`); value.Exists() {
		data.PersistentConnections = tfutils.BoolFromString(value.String())
	} else {
		data.PersistentConnections = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LoopDetection`); value.Exists() {
		data.LoopDetection = tfutils.BoolFromString(value.String())
	} else {
		data.LoopDetection = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DoHostRewriting`); value.Exists() {
		data.DoHostRewriting = tfutils.BoolFromString(value.String())
	} else {
		data.DoHostRewriting = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DoChunkedUpload`); value.Exists() {
		data.DoChunkedUpload = tfutils.BoolFromString(value.String())
	} else {
		data.DoChunkedUpload = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ProcessHTTPErrors`); value.Exists() {
		data.ProcessHttpErrors = tfutils.BoolFromString(value.String())
	} else {
		data.ProcessHttpErrors = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTPClientIPLabel`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HttpClientIpLabel = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HttpClientIpLabel = types.StringValue("X-Client-IP")
	}
	if value := res.Get(pathRoot + `HTTPLogCorIDLabel`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HttpLogCorIdLabel = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HttpLogCorIdLabel = types.StringValue("X-Global-Transaction-ID")
	}
	if value := res.Get(pathRoot + `LoadBalancerHashHeader`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LoadBalancerHashHeader = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LoadBalancerHashHeader = types.StringNull()
	}
	if value := res.Get(pathRoot + `InOrderMode`); value.Exists() {
		data.InOrderMode = &DmGatewayInOrderMode{}
		data.InOrderMode.FromBody(ctx, "", value)
	} else {
		data.InOrderMode = nil
	}
	if value := res.Get(pathRoot + `WSAMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsAMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsAMode = types.StringValue("sync2sync")
	}
	if value := res.Get(pathRoot + `WSARequireAAA`); value.Exists() {
		data.WsARequireAaa = tfutils.BoolFromString(value.String())
	} else {
		data.WsARequireAaa = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSARewriteReplyTo`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsARewriteReplyTo = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsARewriteReplyTo = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSARewriteFaultTo`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsARewriteFaultTo = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsARewriteFaultTo = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSARewriteTo`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsARewriteTo = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsARewriteTo = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSAStrip`); value.Exists() {
		data.WsAStrip = tfutils.BoolFromString(value.String())
	} else {
		data.WsAStrip = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSADefaultReplyTo`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsADefaultReplyTo = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsADefaultReplyTo = types.StringValue("http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous")
	}
	if value := res.Get(pathRoot + `WSADefaultFaultTo`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsADefaultFaultTo = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsADefaultFaultTo = types.StringValue("http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous")
	}
	if value := res.Get(pathRoot + `WSAForce`); value.Exists() {
		data.WsAForce = tfutils.BoolFromString(value.String())
	} else {
		data.WsAForce = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSAGenStyle`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsAGenStyle = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsAGenStyle = types.StringValue("sync")
	}
	if value := res.Get(pathRoot + `WSAHTTPAsyncResponseCode`); value.Exists() {
		data.WsAHttpAsyncResponseCode = types.Int64Value(value.Int())
	} else {
		data.WsAHttpAsyncResponseCode = types.Int64Value(204)
	}
	if value := res.Get(pathRoot + `WSABackProtocol`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsABackProtocol = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsABackProtocol = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSATimeout`); value.Exists() {
		data.WsATimeout = types.Int64Value(value.Int())
	} else {
		data.WsATimeout = types.Int64Value(120)
	}
	if value := res.Get(pathRoot + `WSRMEnabled`); value.Exists() {
		data.WsRmEnabled = tfutils.BoolFromString(value.String())
	} else {
		data.WsRmEnabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSRMSequenceExpiration`); value.Exists() {
		data.WsRmSequenceExpiration = types.Int64Value(value.Int())
	} else {
		data.WsRmSequenceExpiration = types.Int64Value(3600)
	}
	if value := res.Get(pathRoot + `WSRMAAAPolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsRmAaaPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsRmAaaPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSRMDestinationAcceptCreateSequence`); value.Exists() {
		data.WsRmDestinationAcceptCreateSequence = tfutils.BoolFromString(value.String())
	} else {
		data.WsRmDestinationAcceptCreateSequence = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSRMDestinationMaximumSequences`); value.Exists() {
		data.WsRmDestinationMaximumSequences = types.Int64Value(value.Int())
	} else {
		data.WsRmDestinationMaximumSequences = types.Int64Value(400)
	}
	if value := res.Get(pathRoot + `WSRMDestinationInOrder`); value.Exists() {
		data.WsRmDestinationInOrder = tfutils.BoolFromString(value.String())
	} else {
		data.WsRmDestinationInOrder = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSRMDestinationMaximumInOrderQueueLength`); value.Exists() {
		data.WsRmDestinationMaximumInOrderQueueLength = types.Int64Value(value.Int())
	} else {
		data.WsRmDestinationMaximumInOrderQueueLength = types.Int64Value(10)
	}
	if value := res.Get(pathRoot + `WSRMDestinationAcceptOffers`); value.Exists() {
		data.WsRmDestinationAcceptOffers = tfutils.BoolFromString(value.String())
	} else {
		data.WsRmDestinationAcceptOffers = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSRMFrontForce`); value.Exists() {
		data.WsRmFrontForce = tfutils.BoolFromString(value.String())
	} else {
		data.WsRmFrontForce = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSRMBackForce`); value.Exists() {
		data.WsRmBackForce = tfutils.BoolFromString(value.String())
	} else {
		data.WsRmBackForce = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSRMBackCreateSequence`); value.Exists() {
		data.WsRmBackCreateSequence = tfutils.BoolFromString(value.String())
	} else {
		data.WsRmBackCreateSequence = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSRMFrontCreateSequence`); value.Exists() {
		data.WsRmFrontCreateSequence = tfutils.BoolFromString(value.String())
	} else {
		data.WsRmFrontCreateSequence = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSRMSourceMakeOffer`); value.Exists() {
		data.WsRmSourceMakeOffer = tfutils.BoolFromString(value.String())
	} else {
		data.WsRmSourceMakeOffer = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSRMUsesSequenceSSL`); value.Exists() {
		data.WsRmUsesSequenceSsl = tfutils.BoolFromString(value.String())
	} else {
		data.WsRmUsesSequenceSsl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSRMFrontAcksTo`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsRmFrontAcksTo = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsRmFrontAcksTo = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSRMBackAcksTo`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsRmBackAcksTo = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsRmBackAcksTo = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSRMSourceMaximumSequences`); value.Exists() {
		data.WsRmSourceMaximumSequences = types.Int64Value(value.Int())
	} else {
		data.WsRmSourceMaximumSequences = types.Int64Value(400)
	}
	if value := res.Get(pathRoot + `WSRMSourceRetransmissionInterval`); value.Exists() {
		data.WsRmSourceRetransmissionInterval = types.Int64Value(value.Int())
	} else {
		data.WsRmSourceRetransmissionInterval = types.Int64Value(10)
	}
	if value := res.Get(pathRoot + `WSRMSourceExponentialBackoff`); value.Exists() {
		data.WsRmSourceExponentialBackoff = tfutils.BoolFromString(value.String())
	} else {
		data.WsRmSourceExponentialBackoff = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSRMSourceMaximumRetransmissions`); value.Exists() {
		data.WsRmSourceMaximumRetransmissions = types.Int64Value(value.Int())
	} else {
		data.WsRmSourceMaximumRetransmissions = types.Int64Value(4)
	}
	if value := res.Get(pathRoot + `WSRMSourceMaximumQueueLength`); value.Exists() {
		data.WsRmSourceMaximumQueueLength = types.Int64Value(value.Int())
	} else {
		data.WsRmSourceMaximumQueueLength = types.Int64Value(30)
	}
	if value := res.Get(pathRoot + `WSRMSourceRequestAckCount`); value.Exists() {
		data.WsRmSourceRequestAckCount = types.Int64Value(value.Int())
	} else {
		data.WsRmSourceRequestAckCount = types.Int64Value(1)
	}
	if value := res.Get(pathRoot + `WSRMSourceInactivityClose`); value.Exists() {
		data.WsRmSourceInactivityClose = types.Int64Value(value.Int())
	} else {
		data.WsRmSourceInactivityClose = types.Int64Value(360)
	}
	if value := res.Get(pathRoot + `ForcePolicyExec`); value.Exists() {
		data.ForcePolicyExec = tfutils.BoolFromString(value.String())
	} else {
		data.ForcePolicyExec = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RewriteErrors`); value.Exists() {
		data.RewriteErrors = tfutils.BoolFromString(value.String())
	} else {
		data.RewriteErrors = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DelayErrors`); value.Exists() {
		data.DelayErrors = tfutils.BoolFromString(value.String())
	} else {
		data.DelayErrors = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DelayErrorsDuration`); value.Exists() {
		data.DelayErrorsDuration = types.Int64Value(value.Int())
	} else {
		data.DelayErrorsDuration = types.Int64Value(1000)
	}
}

func (data *WSGateway) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `BackHTTPVersion`); value.Exists() && !data.BackHttpVersion.IsNull() {
		data.BackHttpVersion = tfutils.ParseStringFromGJSON(value)
	} else if data.BackHttpVersion.ValueString() != "HTTP/1.1" {
		data.BackHttpVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `RequestType`); value.Exists() && !data.RequestType.IsNull() {
		data.RequestType = tfutils.ParseStringFromGJSON(value)
	} else if data.RequestType.ValueString() != "soap" {
		data.RequestType = types.StringNull()
	}
	if value := res.Get(pathRoot + `ResponseType`); value.Exists() && !data.ResponseType.IsNull() {
		data.ResponseType = tfutils.ParseStringFromGJSON(value)
	} else if data.ResponseType.ValueString() != "soap" {
		data.ResponseType = types.StringNull()
	}
	if value := res.Get(pathRoot + `FollowRedirects`); value.Exists() && !data.FollowRedirects.IsNull() {
		data.FollowRedirects = tfutils.BoolFromString(value.String())
	} else if !data.FollowRedirects.ValueBool() {
		data.FollowRedirects = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowCompression`); value.Exists() && !data.AllowCompression.IsNull() {
		data.AllowCompression = tfutils.BoolFromString(value.String())
	} else if data.AllowCompression.ValueBool() {
		data.AllowCompression = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EnableCompressedPayloadPassthrough`); value.Exists() && !data.EnableCompressedPayloadPassthrough.IsNull() {
		data.EnableCompressedPayloadPassthrough = tfutils.BoolFromString(value.String())
	} else if data.EnableCompressedPayloadPassthrough.ValueBool() {
		data.EnableCompressedPayloadPassthrough = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && !data.Type.IsNull() {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else if data.Type.ValueString() != "static-from-wsdl" {
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `AutoCreateSources`); value.Exists() && !data.AutoCreateSourceS.IsNull() {
		data.AutoCreateSourceS = tfutils.BoolFromString(value.String())
	} else if data.AutoCreateSourceS.ValueBool() {
		data.AutoCreateSourceS = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SSLServerConfigType`); value.Exists() && !data.SslServerConfigType.IsNull() {
		data.SslServerConfigType = tfutils.ParseStringFromGJSON(value)
	} else if data.SslServerConfigType.ValueString() != "server" {
		data.SslServerConfigType = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLServer`); value.Exists() && !data.SslServer.IsNull() {
		data.SslServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLSNIServer`); value.Exists() && !data.SslSniServer.IsNull() {
		data.SslSniServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslSniServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `EndpointRewritePolicy`); value.Exists() && !data.EndpointRewritePolicy.IsNull() {
		data.EndpointRewritePolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EndpointRewritePolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalEndpointRewrite`); value.Exists() && !data.LocalEndpointRewrite.IsNull() {
		data.LocalEndpointRewrite = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalEndpointRewrite = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemoteEndpointRewrite`); value.Exists() && !data.RemoteEndpointRewrite.IsNull() {
		data.RemoteEndpointRewrite = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteEndpointRewrite = types.StringNull()
	}
	if value := res.Get(pathRoot + `AAAPolicy`); value.Exists() && !data.AaaPolicy.IsNull() {
		data.AaaPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AaaPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `StylePolicy`); value.Exists() && !data.StylePolicy.IsNull() {
		data.StylePolicy = tfutils.ParseStringFromGJSON(value)
	} else if data.StylePolicy.ValueString() != "default" {
		data.StylePolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemoteFetchRetry`); value.Exists() {
		data.RemoteFetchRetry.UpdateFromBody(ctx, "", value)
	} else {
		data.RemoteFetchRetry = nil
	}
	if value := res.Get(pathRoot + `WSDLCachePolicy`); value.Exists() && !data.WsdlCachePolicy.IsNull() {
		l := []DmWSDLCachePolicy{}
		for _, v := range value.Array() {
			item := DmWSDLCachePolicy{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.WsdlCachePolicy, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSDLCachePolicyObjectType}, l)
		} else {
			data.WsdlCachePolicy = types.ListNull(types.ObjectType{AttrTypes: DmWSDLCachePolicyObjectType})
		}
	} else {
		data.WsdlCachePolicy = types.ListNull(types.ObjectType{AttrTypes: DmWSDLCachePolicyObjectType})
	}
	if value := res.Get(pathRoot + `BaseWSDL`); value.Exists() && !data.BaseWsdl.IsNull() {
		l := []DmWSBaseWSDL{}
		for _, v := range value.Array() {
			item := DmWSBaseWSDL{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.BaseWsdl, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSBaseWSDLObjectType}, l)
		} else {
			data.BaseWsdl = types.ListNull(types.ObjectType{AttrTypes: DmWSBaseWSDLObjectType})
		}
	} else {
		data.BaseWsdl = types.ListNull(types.ObjectType{AttrTypes: DmWSBaseWSDLObjectType})
	}
	if value := res.Get(pathRoot + `UserToggles`); value.Exists() && !data.UserToggles.IsNull() {
		l := []DmWSUserToggles{}
		for _, v := range value.Array() {
			item := DmWSUserToggles{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.UserToggles, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSUserTogglesObjectType}, l)
		} else {
			data.UserToggles = types.ListNull(types.ObjectType{AttrTypes: DmWSUserTogglesObjectType})
		}
	} else {
		data.UserToggles = types.ListNull(types.ObjectType{AttrTypes: DmWSUserTogglesObjectType})
	}
	if value := res.Get(pathRoot + `ClientPrincipal`); value.Exists() && !data.ClientPrincipal.IsNull() {
		data.ClientPrincipal = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientPrincipal = types.StringNull()
	}
	if value := res.Get(pathRoot + `ServerPrincipal`); value.Exists() && !data.ServerPrincipal.IsNull() {
		data.ServerPrincipal = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ServerPrincipal = types.StringNull()
	}
	if value := res.Get(pathRoot + `KerberosKeytab`); value.Exists() && !data.KerberosKeytab.IsNull() {
		data.KerberosKeytab = tfutils.ParseStringFromGJSON(value)
	} else {
		data.KerberosKeytab = types.StringNull()
	}
	if value := res.Get(pathRoot + `DecryptKey`); value.Exists() && !data.DecryptKey.IsNull() {
		data.DecryptKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DecryptKey = types.StringNull()
	}
	if value := res.Get(pathRoot + `EncryptedKeySHA1CacheLifeTime`); value.Exists() && !data.EncryptedKeySha1CacheLifeTime.IsNull() {
		data.EncryptedKeySha1CacheLifeTime = types.Int64Value(value.Int())
	} else {
		data.EncryptedKeySha1CacheLifeTime = types.Int64Null()
	}
	if value := res.Get(pathRoot + `PreserveKeyChain`); value.Exists() && !data.PreserveKeyChain.IsNull() {
		data.PreserveKeyChain = tfutils.BoolFromString(value.String())
	} else if data.PreserveKeyChain.ValueBool() {
		data.PreserveKeyChain = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DecryptWithKeyFromED`); value.Exists() && !data.DecryptWithKeyFromEd.IsNull() {
		data.DecryptWithKeyFromEd = tfutils.BoolFromString(value.String())
	} else if data.DecryptWithKeyFromEd.ValueBool() {
		data.DecryptWithKeyFromEd = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SOAPActionPolicy`); value.Exists() && !data.SoapActionPolicy.IsNull() {
		data.SoapActionPolicy = tfutils.ParseStringFromGJSON(value)
	} else if data.SoapActionPolicy.ValueString() != "lax" {
		data.SoapActionPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSRRSubscriptions`); value.Exists() && !data.WsrrSubscriptions.IsNull() {
		l := []DmWSRRWSDLSource{}
		for _, v := range value.Array() {
			item := DmWSRRWSDLSource{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.WsrrSubscriptions, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSRRWSDLSourceObjectType}, l)
		} else {
			data.WsrrSubscriptions = types.ListNull(types.ObjectType{AttrTypes: DmWSRRWSDLSourceObjectType})
		}
	} else {
		data.WsrrSubscriptions = types.ListNull(types.ObjectType{AttrTypes: DmWSRRWSDLSourceObjectType})
	}
	if value := res.Get(pathRoot + `WSRRSavedSearchSubscriptions`); value.Exists() && !data.WsrrSavedSearchSubscriptions.IsNull() {
		l := []DmWSRRSavedSearchWSDLSource{}
		for _, v := range value.Array() {
			item := DmWSRRSavedSearchWSDLSource{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.WsrrSavedSearchSubscriptions, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSRRSavedSearchWSDLSourceObjectType}, l)
		} else {
			data.WsrrSavedSearchSubscriptions = types.ListNull(types.ObjectType{AttrTypes: DmWSRRSavedSearchWSDLSourceObjectType})
		}
	} else {
		data.WsrrSavedSearchSubscriptions = types.ListNull(types.ObjectType{AttrTypes: DmWSRRSavedSearchWSDLSourceObjectType})
	}
	if value := res.Get(pathRoot + `OperationPriority`); value.Exists() && !data.OperationPriority.IsNull() {
		l := []DmWSOperationSchedulerPriority{}
		for _, v := range value.Array() {
			item := DmWSOperationSchedulerPriority{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.OperationPriority, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSOperationSchedulerPriorityObjectType}, l)
		} else {
			data.OperationPriority = types.ListNull(types.ObjectType{AttrTypes: DmWSOperationSchedulerPriorityObjectType})
		}
	} else {
		data.OperationPriority = types.ListNull(types.ObjectType{AttrTypes: DmWSOperationSchedulerPriorityObjectType})
	}
	if value := res.Get(pathRoot + `OperationConformancePolicy`); value.Exists() && !data.OperationConformancePolicy.IsNull() {
		l := []DmWSOperationConformancePolicy{}
		for _, v := range value.Array() {
			item := DmWSOperationConformancePolicy{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.OperationConformancePolicy, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSOperationConformancePolicyObjectType}, l)
		} else {
			data.OperationConformancePolicy = types.ListNull(types.ObjectType{AttrTypes: DmWSOperationConformancePolicyObjectType})
		}
	} else {
		data.OperationConformancePolicy = types.ListNull(types.ObjectType{AttrTypes: DmWSOperationConformancePolicyObjectType})
	}
	if value := res.Get(pathRoot + `OperationPolicySubjectOptOut`); value.Exists() && !data.OperationPolicySubjectOptOut.IsNull() {
		l := []DmWSOperationPolicySubjectOptOut{}
		for _, v := range value.Array() {
			item := DmWSOperationPolicySubjectOptOut{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.OperationPolicySubjectOptOut, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSOperationPolicySubjectOptOutObjectType}, l)
		} else {
			data.OperationPolicySubjectOptOut = types.ListNull(types.ObjectType{AttrTypes: DmWSOperationPolicySubjectOptOutObjectType})
		}
	} else {
		data.OperationPolicySubjectOptOut = types.ListNull(types.ObjectType{AttrTypes: DmWSOperationPolicySubjectOptOutObjectType})
	}
	if value := res.Get(pathRoot + `PolicyParameter`); value.Exists() && !data.PolicyParameter.IsNull() {
		l := []DmWSPolicyParameters{}
		for _, v := range value.Array() {
			item := DmWSPolicyParameters{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.PolicyParameter, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSPolicyParametersObjectType}, l)
		} else {
			data.PolicyParameter = types.ListNull(types.ObjectType{AttrTypes: DmWSPolicyParametersObjectType})
		}
	} else {
		data.PolicyParameter = types.ListNull(types.ObjectType{AttrTypes: DmWSPolicyParametersObjectType})
	}
	if value := res.Get(pathRoot + `ReliableMessaging`); value.Exists() && !data.ReliableMessaging.IsNull() {
		l := []DmWSOperationReliableMessaging{}
		for _, v := range value.Array() {
			item := DmWSOperationReliableMessaging{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.ReliableMessaging, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSOperationReliableMessagingObjectType}, l)
		} else {
			data.ReliableMessaging = types.ListNull(types.ObjectType{AttrTypes: DmWSOperationReliableMessagingObjectType})
		}
	} else {
		data.ReliableMessaging = types.ListNull(types.ObjectType{AttrTypes: DmWSOperationReliableMessagingObjectType})
	}
	if value := res.Get(pathRoot + `WSMAgentMonitor`); value.Exists() && !data.WSmAgentMonitor.IsNull() {
		data.WSmAgentMonitor = tfutils.BoolFromString(value.String())
	} else if !data.WSmAgentMonitor.ValueBool() {
		data.WSmAgentMonitor = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSMAgentMonitorPCM`); value.Exists() && !data.WSmAgentMonitorPcm.IsNull() {
		data.WSmAgentMonitorPcm = tfutils.ParseStringFromGJSON(value)
	} else if data.WSmAgentMonitorPcm.ValueString() != "all-messages" {
		data.WSmAgentMonitorPcm = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProcessRespRulesOnOneWayMEP`); value.Exists() && !data.ProcessRespRulesOnOneWayMep.IsNull() {
		data.ProcessRespRulesOnOneWayMep = tfutils.BoolFromString(value.String())
	} else if data.ProcessRespRulesOnOneWayMep.ValueBool() {
		data.ProcessRespRulesOnOneWayMep = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `Priority`); value.Exists() && !data.Priority.IsNull() {
		data.Priority = tfutils.ParseStringFromGJSON(value)
	} else if data.Priority.ValueString() != "normal" {
		data.Priority = types.StringNull()
	}
	if value := res.Get(pathRoot + `FrontProtocol`); value.Exists() && !data.FrontProtocol.IsNull() {
		data.FrontProtocol = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.FrontProtocol = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `XMLManager`); value.Exists() && !data.XmlManager.IsNull() {
		data.XmlManager = tfutils.ParseStringFromGJSON(value)
	} else if data.XmlManager.ValueString() != "default" {
		data.XmlManager = types.StringNull()
	}
	if value := res.Get(pathRoot + `URLRewritePolicy`); value.Exists() && !data.UrlRewritePolicy.IsNull() {
		data.UrlRewritePolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UrlRewritePolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLClientConfigType`); value.Exists() && !data.SslClientConfigType.IsNull() {
		data.SslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else if data.SslClientConfigType.ValueString() != "client" {
		data.SslClientConfigType = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLClient`); value.Exists() && !data.SslClient.IsNull() {
		data.SslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `FWCred`); value.Exists() && !data.FwCred.IsNull() {
		data.FwCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FwCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `HeaderInjection`); value.Exists() && !data.HeaderInjection.IsNull() {
		l := []DmHeaderInjection{}
		for _, v := range value.Array() {
			item := DmHeaderInjection{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.HeaderInjection, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmHeaderInjectionObjectType}, l)
		} else {
			data.HeaderInjection = types.ListNull(types.ObjectType{AttrTypes: DmHeaderInjectionObjectType})
		}
	} else {
		data.HeaderInjection = types.ListNull(types.ObjectType{AttrTypes: DmHeaderInjectionObjectType})
	}
	if value := res.Get(pathRoot + `HeaderSuppression`); value.Exists() && !data.HeaderSuppression.IsNull() {
		l := []DmHeaderSuppression{}
		for _, v := range value.Array() {
			item := DmHeaderSuppression{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.HeaderSuppression, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmHeaderSuppressionObjectType}, l)
		} else {
			data.HeaderSuppression = types.ListNull(types.ObjectType{AttrTypes: DmHeaderSuppressionObjectType})
		}
	} else {
		data.HeaderSuppression = types.ListNull(types.ObjectType{AttrTypes: DmHeaderSuppressionObjectType})
	}
	if value := res.Get(pathRoot + `StylesheetParameters`); value.Exists() && !data.StylesheetParameters.IsNull() {
		l := []DmStylesheetParameter{}
		for _, v := range value.Array() {
			item := DmStylesheetParameter{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.StylesheetParameters, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmStylesheetParameterObjectType}, l)
		} else {
			data.StylesheetParameters = types.ListNull(types.ObjectType{AttrTypes: DmStylesheetParameterObjectType})
		}
	} else {
		data.StylesheetParameters = types.ListNull(types.ObjectType{AttrTypes: DmStylesheetParameterObjectType})
	}
	if value := res.Get(pathRoot + `DefaultParamNamespace`); value.Exists() && !data.DefaultParamNamespace.IsNull() {
		data.DefaultParamNamespace = tfutils.ParseStringFromGJSON(value)
	} else if data.DefaultParamNamespace.ValueString() != "http://www.datapower.com/param/config" {
		data.DefaultParamNamespace = types.StringNull()
	}
	if value := res.Get(pathRoot + `QueryParamNamespace`); value.Exists() && !data.QueryParamNamespace.IsNull() {
		data.QueryParamNamespace = tfutils.ParseStringFromGJSON(value)
	} else if data.QueryParamNamespace.ValueString() != "http://www.datapower.com/param/query" {
		data.QueryParamNamespace = types.StringNull()
	}
	if value := res.Get(pathRoot + `BackendUrl`); value.Exists() && !data.BackendUrl.IsNull() {
		data.BackendUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BackendUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `PropagateURI`); value.Exists() && !data.PropagateUri.IsNull() {
		data.PropagateUri = tfutils.BoolFromString(value.String())
	} else if !data.PropagateUri.ValueBool() {
		data.PropagateUri = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ServiceMonitors`); value.Exists() && !data.ServiceMonitors.IsNull() {
		data.ServiceMonitors = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.ServiceMonitors = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `CountMonitors`); value.Exists() && !data.CountMonitors.IsNull() {
		data.CountMonitors = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.CountMonitors = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `DurationMonitors`); value.Exists() && !data.DurationMonitors.IsNull() {
		data.DurationMonitors = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.DurationMonitors = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `MonitorProcessingPolicy`); value.Exists() && !data.MonitorProcessingPolicy.IsNull() {
		data.MonitorProcessingPolicy = tfutils.ParseStringFromGJSON(value)
	} else if data.MonitorProcessingPolicy.ValueString() != "terminate-at-first-throttle" {
		data.MonitorProcessingPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `RequestAttachments`); value.Exists() && !data.RequestAttachments.IsNull() {
		data.RequestAttachments = tfutils.ParseStringFromGJSON(value)
	} else if data.RequestAttachments.ValueString() != "strip" {
		data.RequestAttachments = types.StringNull()
	}
	if value := res.Get(pathRoot + `ResponseAttachments`); value.Exists() && !data.ResponseAttachments.IsNull() {
		data.ResponseAttachments = tfutils.ParseStringFromGJSON(value)
	} else if data.ResponseAttachments.ValueString() != "strip" {
		data.ResponseAttachments = types.StringNull()
	}
	if value := res.Get(pathRoot + `RequestAttachmentsFlowControl`); value.Exists() && !data.RequestAttachmentsFlowControl.IsNull() {
		data.RequestAttachmentsFlowControl = tfutils.BoolFromString(value.String())
	} else if data.RequestAttachmentsFlowControl.ValueBool() {
		data.RequestAttachmentsFlowControl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ResponseAttachmentsFlowControl`); value.Exists() && !data.ResponseAttachmentsFlowControl.IsNull() {
		data.ResponseAttachmentsFlowControl = tfutils.BoolFromString(value.String())
	} else if data.ResponseAttachmentsFlowControl.ValueBool() {
		data.ResponseAttachmentsFlowControl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RootPartNotFirstAction`); value.Exists() && !data.RootPartNotFirstAction.IsNull() {
		data.RootPartNotFirstAction = tfutils.ParseStringFromGJSON(value)
	} else if data.RootPartNotFirstAction.ValueString() != "process-in-order" {
		data.RootPartNotFirstAction = types.StringNull()
	}
	if value := res.Get(pathRoot + `FrontAttachmentFormat`); value.Exists() && !data.FrontAttachmentFormat.IsNull() {
		data.FrontAttachmentFormat = tfutils.ParseStringFromGJSON(value)
	} else if data.FrontAttachmentFormat.ValueString() != "dynamic" {
		data.FrontAttachmentFormat = types.StringNull()
	}
	if value := res.Get(pathRoot + `BackAttachmentFormat`); value.Exists() && !data.BackAttachmentFormat.IsNull() {
		data.BackAttachmentFormat = tfutils.ParseStringFromGJSON(value)
	} else if data.BackAttachmentFormat.ValueString() != "dynamic" {
		data.BackAttachmentFormat = types.StringNull()
	}
	if value := res.Get(pathRoot + `MIMEFrontHeaders`); value.Exists() && !data.MimeFrontHeaders.IsNull() {
		data.MimeFrontHeaders = tfutils.BoolFromString(value.String())
	} else if !data.MimeFrontHeaders.ValueBool() {
		data.MimeFrontHeaders = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MIMEBackHeaders`); value.Exists() && !data.MimeBackHeaders.IsNull() {
		data.MimeBackHeaders = tfutils.BoolFromString(value.String())
	} else if !data.MimeBackHeaders.ValueBool() {
		data.MimeBackHeaders = types.BoolNull()
	}
	if value := res.Get(pathRoot + `StreamOutputToBack`); value.Exists() && !data.StreamOutputToBack.IsNull() {
		data.StreamOutputToBack = tfutils.ParseStringFromGJSON(value)
	} else if data.StreamOutputToBack.ValueString() != "buffer-until-verification" {
		data.StreamOutputToBack = types.StringNull()
	}
	if value := res.Get(pathRoot + `StreamOutputToFront`); value.Exists() && !data.StreamOutputToFront.IsNull() {
		data.StreamOutputToFront = tfutils.ParseStringFromGJSON(value)
	} else if data.StreamOutputToFront.ValueString() != "buffer-until-verification" {
		data.StreamOutputToFront = types.StringNull()
	}
	if value := res.Get(pathRoot + `MaxMessageSize`); value.Exists() && !data.MaxMessageSize.IsNull() {
		data.MaxMessageSize = types.Int64Value(value.Int())
	} else {
		data.MaxMessageSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `GatewayParserLimits`); value.Exists() && !data.GatewayParserLimits.IsNull() {
		data.GatewayParserLimits = tfutils.BoolFromString(value.String())
	} else if data.GatewayParserLimits.ValueBool() {
		data.GatewayParserLimits = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ParserLimitsElementDepth`); value.Exists() && !data.ParserLimitsElementDepth.IsNull() {
		data.ParserLimitsElementDepth = types.Int64Value(value.Int())
	} else if data.ParserLimitsElementDepth.ValueInt64() != 512 {
		data.ParserLimitsElementDepth = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ParserLimitsAttributeCount`); value.Exists() && !data.ParserLimitsAttributeCount.IsNull() {
		data.ParserLimitsAttributeCount = types.Int64Value(value.Int())
	} else if data.ParserLimitsAttributeCount.ValueInt64() != 128 {
		data.ParserLimitsAttributeCount = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ParserLimitsMaxNodeSize`); value.Exists() && !data.ParserLimitsMaxNodeSize.IsNull() {
		data.ParserLimitsMaxNodeSize = types.Int64Value(value.Int())
	} else if data.ParserLimitsMaxNodeSize.ValueInt64() != 33554432 {
		data.ParserLimitsMaxNodeSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ParserLimitsExternalReferences`); value.Exists() && !data.ParserLimitsExternalReferences.IsNull() {
		data.ParserLimitsExternalReferences = tfutils.ParseStringFromGJSON(value)
	} else if data.ParserLimitsExternalReferences.ValueString() != "forbid" {
		data.ParserLimitsExternalReferences = types.StringNull()
	}
	if value := res.Get(pathRoot + `ParserLimitsMaxPrefixes`); value.Exists() && !data.ParserLimitsMaxPrefixes.IsNull() {
		data.ParserLimitsMaxPrefixes = types.Int64Value(value.Int())
	} else if data.ParserLimitsMaxPrefixes.ValueInt64() != 1024 {
		data.ParserLimitsMaxPrefixes = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ParserLimitsMaxNamespaces`); value.Exists() && !data.ParserLimitsMaxNamespaces.IsNull() {
		data.ParserLimitsMaxNamespaces = types.Int64Value(value.Int())
	} else if data.ParserLimitsMaxNamespaces.ValueInt64() != 1024 {
		data.ParserLimitsMaxNamespaces = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ParserLimitsMaxLocalNames`); value.Exists() && !data.ParserLimitsMaxLocalNames.IsNull() {
		data.ParserLimitsMaxLocalNames = types.Int64Value(value.Int())
	} else if data.ParserLimitsMaxLocalNames.ValueInt64() != 60000 {
		data.ParserLimitsMaxLocalNames = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ParserLimitsAttachmentByteCount`); value.Exists() && !data.ParserLimitsAttachmentByteCount.IsNull() {
		data.ParserLimitsAttachmentByteCount = types.Int64Value(value.Int())
	} else if data.ParserLimitsAttachmentByteCount.ValueInt64() != 2000000000 {
		data.ParserLimitsAttachmentByteCount = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ParserLimitsAttachmentPackageByteCount`); value.Exists() && !data.ParserLimitsAttachmentPackageByteCount.IsNull() {
		data.ParserLimitsAttachmentPackageByteCount = types.Int64Value(value.Int())
	} else {
		data.ParserLimitsAttachmentPackageByteCount = types.Int64Null()
	}
	if value := res.Get(pathRoot + `DebugMode`); value.Exists() && !data.DebugMode.IsNull() {
		data.DebugMode = tfutils.ParseStringFromGJSON(value)
	} else if data.DebugMode.ValueString() != "off" {
		data.DebugMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `DebugHistory`); value.Exists() && !data.DebugHistory.IsNull() {
		data.DebugHistory = types.Int64Value(value.Int())
	} else if data.DebugHistory.ValueInt64() != 25 {
		data.DebugHistory = types.Int64Null()
	}
	if value := res.Get(pathRoot + `DebugTrigger`); value.Exists() && !data.DebugTrigger.IsNull() {
		l := []DmMSDebugTriggerType{}
		for _, v := range value.Array() {
			item := DmMSDebugTriggerType{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.DebugTrigger, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmMSDebugTriggerTypeObjectType}, l)
		} else {
			data.DebugTrigger = types.ListNull(types.ObjectType{AttrTypes: DmMSDebugTriggerTypeObjectType})
		}
	} else {
		data.DebugTrigger = types.ListNull(types.ObjectType{AttrTypes: DmMSDebugTriggerTypeObjectType})
	}
	if value := res.Get(pathRoot + `FlowControl`); value.Exists() && !data.FlowControl.IsNull() {
		data.FlowControl = tfutils.BoolFromString(value.String())
	} else if data.FlowControl.ValueBool() {
		data.FlowControl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SOAPSchemaURL`); value.Exists() && !data.SoapSchemaUrl.IsNull() {
		data.SoapSchemaUrl = tfutils.ParseStringFromGJSON(value)
	} else if data.SoapSchemaUrl.ValueString() != "store:///schemas/soap-envelope.xsd" {
		data.SoapSchemaUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `FrontTimeout`); value.Exists() && !data.FrontTimeout.IsNull() {
		data.FrontTimeout = types.Int64Value(value.Int())
	} else if data.FrontTimeout.ValueInt64() != 120 {
		data.FrontTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `BackTimeout`); value.Exists() && !data.BackTimeout.IsNull() {
		data.BackTimeout = types.Int64Value(value.Int())
	} else if data.BackTimeout.ValueInt64() != 120 {
		data.BackTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `FrontPersistentTimeout`); value.Exists() && !data.FrontPersistentTimeout.IsNull() {
		data.FrontPersistentTimeout = types.Int64Value(value.Int())
	} else if data.FrontPersistentTimeout.ValueInt64() != 180 {
		data.FrontPersistentTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `BackPersistentTimeout`); value.Exists() && !data.BackPersistentTimeout.IsNull() {
		data.BackPersistentTimeout = types.Int64Value(value.Int())
	} else if data.BackPersistentTimeout.ValueInt64() != 180 {
		data.BackPersistentTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `IncludeResponseTypeEncoding`); value.Exists() && !data.IncludeResponseTypeEncoding.IsNull() {
		data.IncludeResponseTypeEncoding = tfutils.BoolFromString(value.String())
	} else if data.IncludeResponseTypeEncoding.ValueBool() {
		data.IncludeResponseTypeEncoding = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PersistentConnections`); value.Exists() && !data.PersistentConnections.IsNull() {
		data.PersistentConnections = tfutils.BoolFromString(value.String())
	} else if !data.PersistentConnections.ValueBool() {
		data.PersistentConnections = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LoopDetection`); value.Exists() && !data.LoopDetection.IsNull() {
		data.LoopDetection = tfutils.BoolFromString(value.String())
	} else if data.LoopDetection.ValueBool() {
		data.LoopDetection = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DoHostRewriting`); value.Exists() && !data.DoHostRewriting.IsNull() {
		data.DoHostRewriting = tfutils.BoolFromString(value.String())
	} else if !data.DoHostRewriting.ValueBool() {
		data.DoHostRewriting = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DoChunkedUpload`); value.Exists() && !data.DoChunkedUpload.IsNull() {
		data.DoChunkedUpload = tfutils.BoolFromString(value.String())
	} else if data.DoChunkedUpload.ValueBool() {
		data.DoChunkedUpload = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ProcessHTTPErrors`); value.Exists() && !data.ProcessHttpErrors.IsNull() {
		data.ProcessHttpErrors = tfutils.BoolFromString(value.String())
	} else if !data.ProcessHttpErrors.ValueBool() {
		data.ProcessHttpErrors = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTPClientIPLabel`); value.Exists() && !data.HttpClientIpLabel.IsNull() {
		data.HttpClientIpLabel = tfutils.ParseStringFromGJSON(value)
	} else if data.HttpClientIpLabel.ValueString() != "X-Client-IP" {
		data.HttpClientIpLabel = types.StringNull()
	}
	if value := res.Get(pathRoot + `HTTPLogCorIDLabel`); value.Exists() && !data.HttpLogCorIdLabel.IsNull() {
		data.HttpLogCorIdLabel = tfutils.ParseStringFromGJSON(value)
	} else if data.HttpLogCorIdLabel.ValueString() != "X-Global-Transaction-ID" {
		data.HttpLogCorIdLabel = types.StringNull()
	}
	if value := res.Get(pathRoot + `LoadBalancerHashHeader`); value.Exists() && !data.LoadBalancerHashHeader.IsNull() {
		data.LoadBalancerHashHeader = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LoadBalancerHashHeader = types.StringNull()
	}
	if value := res.Get(pathRoot + `InOrderMode`); value.Exists() {
		data.InOrderMode.UpdateFromBody(ctx, "", value)
	} else {
		data.InOrderMode = nil
	}
	if value := res.Get(pathRoot + `WSAMode`); value.Exists() && !data.WsAMode.IsNull() {
		data.WsAMode = tfutils.ParseStringFromGJSON(value)
	} else if data.WsAMode.ValueString() != "sync2sync" {
		data.WsAMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSARequireAAA`); value.Exists() && !data.WsARequireAaa.IsNull() {
		data.WsARequireAaa = tfutils.BoolFromString(value.String())
	} else if !data.WsARequireAaa.ValueBool() {
		data.WsARequireAaa = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSARewriteReplyTo`); value.Exists() && !data.WsARewriteReplyTo.IsNull() {
		data.WsARewriteReplyTo = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsARewriteReplyTo = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSARewriteFaultTo`); value.Exists() && !data.WsARewriteFaultTo.IsNull() {
		data.WsARewriteFaultTo = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsARewriteFaultTo = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSARewriteTo`); value.Exists() && !data.WsARewriteTo.IsNull() {
		data.WsARewriteTo = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsARewriteTo = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSAStrip`); value.Exists() && !data.WsAStrip.IsNull() {
		data.WsAStrip = tfutils.BoolFromString(value.String())
	} else if !data.WsAStrip.ValueBool() {
		data.WsAStrip = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSADefaultReplyTo`); value.Exists() && !data.WsADefaultReplyTo.IsNull() {
		data.WsADefaultReplyTo = tfutils.ParseStringFromGJSON(value)
	} else if data.WsADefaultReplyTo.ValueString() != "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous" {
		data.WsADefaultReplyTo = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSADefaultFaultTo`); value.Exists() && !data.WsADefaultFaultTo.IsNull() {
		data.WsADefaultFaultTo = tfutils.ParseStringFromGJSON(value)
	} else if data.WsADefaultFaultTo.ValueString() != "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous" {
		data.WsADefaultFaultTo = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSAForce`); value.Exists() && !data.WsAForce.IsNull() {
		data.WsAForce = tfutils.BoolFromString(value.String())
	} else if data.WsAForce.ValueBool() {
		data.WsAForce = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSAGenStyle`); value.Exists() && !data.WsAGenStyle.IsNull() {
		data.WsAGenStyle = tfutils.ParseStringFromGJSON(value)
	} else if data.WsAGenStyle.ValueString() != "sync" {
		data.WsAGenStyle = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSAHTTPAsyncResponseCode`); value.Exists() && !data.WsAHttpAsyncResponseCode.IsNull() {
		data.WsAHttpAsyncResponseCode = types.Int64Value(value.Int())
	} else if data.WsAHttpAsyncResponseCode.ValueInt64() != 204 {
		data.WsAHttpAsyncResponseCode = types.Int64Null()
	}
	if value := res.Get(pathRoot + `WSABackProtocol`); value.Exists() && !data.WsABackProtocol.IsNull() {
		data.WsABackProtocol = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsABackProtocol = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSATimeout`); value.Exists() && !data.WsATimeout.IsNull() {
		data.WsATimeout = types.Int64Value(value.Int())
	} else if data.WsATimeout.ValueInt64() != 120 {
		data.WsATimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `WSRMEnabled`); value.Exists() && !data.WsRmEnabled.IsNull() {
		data.WsRmEnabled = tfutils.BoolFromString(value.String())
	} else if data.WsRmEnabled.ValueBool() {
		data.WsRmEnabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSRMSequenceExpiration`); value.Exists() && !data.WsRmSequenceExpiration.IsNull() {
		data.WsRmSequenceExpiration = types.Int64Value(value.Int())
	} else if data.WsRmSequenceExpiration.ValueInt64() != 3600 {
		data.WsRmSequenceExpiration = types.Int64Null()
	}
	if value := res.Get(pathRoot + `WSRMAAAPolicy`); value.Exists() && !data.WsRmAaaPolicy.IsNull() {
		data.WsRmAaaPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsRmAaaPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSRMDestinationAcceptCreateSequence`); value.Exists() && !data.WsRmDestinationAcceptCreateSequence.IsNull() {
		data.WsRmDestinationAcceptCreateSequence = tfutils.BoolFromString(value.String())
	} else if !data.WsRmDestinationAcceptCreateSequence.ValueBool() {
		data.WsRmDestinationAcceptCreateSequence = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSRMDestinationMaximumSequences`); value.Exists() && !data.WsRmDestinationMaximumSequences.IsNull() {
		data.WsRmDestinationMaximumSequences = types.Int64Value(value.Int())
	} else if data.WsRmDestinationMaximumSequences.ValueInt64() != 400 {
		data.WsRmDestinationMaximumSequences = types.Int64Null()
	}
	if value := res.Get(pathRoot + `WSRMDestinationInOrder`); value.Exists() && !data.WsRmDestinationInOrder.IsNull() {
		data.WsRmDestinationInOrder = tfutils.BoolFromString(value.String())
	} else if data.WsRmDestinationInOrder.ValueBool() {
		data.WsRmDestinationInOrder = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSRMDestinationMaximumInOrderQueueLength`); value.Exists() && !data.WsRmDestinationMaximumInOrderQueueLength.IsNull() {
		data.WsRmDestinationMaximumInOrderQueueLength = types.Int64Value(value.Int())
	} else if data.WsRmDestinationMaximumInOrderQueueLength.ValueInt64() != 10 {
		data.WsRmDestinationMaximumInOrderQueueLength = types.Int64Null()
	}
	if value := res.Get(pathRoot + `WSRMDestinationAcceptOffers`); value.Exists() && !data.WsRmDestinationAcceptOffers.IsNull() {
		data.WsRmDestinationAcceptOffers = tfutils.BoolFromString(value.String())
	} else if data.WsRmDestinationAcceptOffers.ValueBool() {
		data.WsRmDestinationAcceptOffers = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSRMFrontForce`); value.Exists() && !data.WsRmFrontForce.IsNull() {
		data.WsRmFrontForce = tfutils.BoolFromString(value.String())
	} else if data.WsRmFrontForce.ValueBool() {
		data.WsRmFrontForce = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSRMBackForce`); value.Exists() && !data.WsRmBackForce.IsNull() {
		data.WsRmBackForce = tfutils.BoolFromString(value.String())
	} else if data.WsRmBackForce.ValueBool() {
		data.WsRmBackForce = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSRMBackCreateSequence`); value.Exists() && !data.WsRmBackCreateSequence.IsNull() {
		data.WsRmBackCreateSequence = tfutils.BoolFromString(value.String())
	} else if data.WsRmBackCreateSequence.ValueBool() {
		data.WsRmBackCreateSequence = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSRMFrontCreateSequence`); value.Exists() && !data.WsRmFrontCreateSequence.IsNull() {
		data.WsRmFrontCreateSequence = tfutils.BoolFromString(value.String())
	} else if data.WsRmFrontCreateSequence.ValueBool() {
		data.WsRmFrontCreateSequence = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSRMSourceMakeOffer`); value.Exists() && !data.WsRmSourceMakeOffer.IsNull() {
		data.WsRmSourceMakeOffer = tfutils.BoolFromString(value.String())
	} else if data.WsRmSourceMakeOffer.ValueBool() {
		data.WsRmSourceMakeOffer = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSRMUsesSequenceSSL`); value.Exists() && !data.WsRmUsesSequenceSsl.IsNull() {
		data.WsRmUsesSequenceSsl = tfutils.BoolFromString(value.String())
	} else if data.WsRmUsesSequenceSsl.ValueBool() {
		data.WsRmUsesSequenceSsl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSRMFrontAcksTo`); value.Exists() && !data.WsRmFrontAcksTo.IsNull() {
		data.WsRmFrontAcksTo = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsRmFrontAcksTo = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSRMBackAcksTo`); value.Exists() && !data.WsRmBackAcksTo.IsNull() {
		data.WsRmBackAcksTo = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsRmBackAcksTo = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSRMSourceMaximumSequences`); value.Exists() && !data.WsRmSourceMaximumSequences.IsNull() {
		data.WsRmSourceMaximumSequences = types.Int64Value(value.Int())
	} else if data.WsRmSourceMaximumSequences.ValueInt64() != 400 {
		data.WsRmSourceMaximumSequences = types.Int64Null()
	}
	if value := res.Get(pathRoot + `WSRMSourceRetransmissionInterval`); value.Exists() && !data.WsRmSourceRetransmissionInterval.IsNull() {
		data.WsRmSourceRetransmissionInterval = types.Int64Value(value.Int())
	} else if data.WsRmSourceRetransmissionInterval.ValueInt64() != 10 {
		data.WsRmSourceRetransmissionInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `WSRMSourceExponentialBackoff`); value.Exists() && !data.WsRmSourceExponentialBackoff.IsNull() {
		data.WsRmSourceExponentialBackoff = tfutils.BoolFromString(value.String())
	} else if !data.WsRmSourceExponentialBackoff.ValueBool() {
		data.WsRmSourceExponentialBackoff = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSRMSourceMaximumRetransmissions`); value.Exists() && !data.WsRmSourceMaximumRetransmissions.IsNull() {
		data.WsRmSourceMaximumRetransmissions = types.Int64Value(value.Int())
	} else if data.WsRmSourceMaximumRetransmissions.ValueInt64() != 4 {
		data.WsRmSourceMaximumRetransmissions = types.Int64Null()
	}
	if value := res.Get(pathRoot + `WSRMSourceMaximumQueueLength`); value.Exists() && !data.WsRmSourceMaximumQueueLength.IsNull() {
		data.WsRmSourceMaximumQueueLength = types.Int64Value(value.Int())
	} else if data.WsRmSourceMaximumQueueLength.ValueInt64() != 30 {
		data.WsRmSourceMaximumQueueLength = types.Int64Null()
	}
	if value := res.Get(pathRoot + `WSRMSourceRequestAckCount`); value.Exists() && !data.WsRmSourceRequestAckCount.IsNull() {
		data.WsRmSourceRequestAckCount = types.Int64Value(value.Int())
	} else if data.WsRmSourceRequestAckCount.ValueInt64() != 1 {
		data.WsRmSourceRequestAckCount = types.Int64Null()
	}
	if value := res.Get(pathRoot + `WSRMSourceInactivityClose`); value.Exists() && !data.WsRmSourceInactivityClose.IsNull() {
		data.WsRmSourceInactivityClose = types.Int64Value(value.Int())
	} else if data.WsRmSourceInactivityClose.ValueInt64() != 360 {
		data.WsRmSourceInactivityClose = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ForcePolicyExec`); value.Exists() && !data.ForcePolicyExec.IsNull() {
		data.ForcePolicyExec = tfutils.BoolFromString(value.String())
	} else if data.ForcePolicyExec.ValueBool() {
		data.ForcePolicyExec = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RewriteErrors`); value.Exists() && !data.RewriteErrors.IsNull() {
		data.RewriteErrors = tfutils.BoolFromString(value.String())
	} else if !data.RewriteErrors.ValueBool() {
		data.RewriteErrors = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DelayErrors`); value.Exists() && !data.DelayErrors.IsNull() {
		data.DelayErrors = tfutils.BoolFromString(value.String())
	} else if !data.DelayErrors.ValueBool() {
		data.DelayErrors = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DelayErrorsDuration`); value.Exists() && !data.DelayErrorsDuration.IsNull() {
		data.DelayErrorsDuration = types.Int64Value(value.Int())
	} else if data.DelayErrorsDuration.ValueInt64() != 1000 {
		data.DelayErrorsDuration = types.Int64Null()
	}
}
