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

type XMLFirewallService struct {
	Id                                     types.String                `tfsdk:"id"`
	AppDomain                              types.String                `tfsdk:"app_domain"`
	Type                                   types.String                `tfsdk:"type"`
	XmlManager                             types.String                `tfsdk:"xml_manager"`
	UrlRewritePolicy                       types.String                `tfsdk:"url_rewrite_policy"`
	StylePolicy                            types.String                `tfsdk:"style_policy"`
	MaxMessageSize                         types.Int64                 `tfsdk:"max_message_size"`
	RequestType                            types.String                `tfsdk:"request_type"`
	ResponseType                           types.String                `tfsdk:"response_type"`
	FwCred                                 types.String                `tfsdk:"fw_cred"`
	ServiceMonitors                        types.List                  `tfsdk:"service_monitors"`
	RequestAttachments                     types.String                `tfsdk:"request_attachments"`
	ResponseAttachments                    types.String                `tfsdk:"response_attachments"`
	RootPartNotFirstAction                 types.String                `tfsdk:"root_part_not_first_action"`
	FrontAttachmentFormat                  types.String                `tfsdk:"front_attachment_format"`
	BackAttachmentFormat                   types.String                `tfsdk:"back_attachment_format"`
	MimeHeaders                            types.Bool                  `tfsdk:"mime_headers"`
	RewriteErrors                          types.Bool                  `tfsdk:"rewrite_errors"`
	DelayErrors                            types.Bool                  `tfsdk:"delay_errors"`
	DelayErrorsDuration                    types.Int64                 `tfsdk:"delay_errors_duration"`
	SoapSchemaUrl                          types.String                `tfsdk:"soap_schema_url"`
	WsdlResponsePolicy                     types.String                `tfsdk:"wsdl_response_policy"`
	WsdlFileLocation                       types.String                `tfsdk:"wsdl_file_location"`
	FirewallParserLimits                   types.Bool                  `tfsdk:"firewall_parser_limits"`
	ParserLimitsBytesScanned               types.Int64                 `tfsdk:"parser_limits_bytes_scanned"`
	ParserLimitsElementDepth               types.Int64                 `tfsdk:"parser_limits_element_depth"`
	ParserLimitsAttributeCount             types.Int64                 `tfsdk:"parser_limits_attribute_count"`
	ParserLimitsMaxNodeSize                types.Int64                 `tfsdk:"parser_limits_max_node_size"`
	ParserLimitsMaxPrefixes                types.Int64                 `tfsdk:"parser_limits_max_prefixes"`
	ParserLimitsMaxNamespaces              types.Int64                 `tfsdk:"parser_limits_max_namespaces"`
	ParserLimitsMaxLocalNames              types.Int64                 `tfsdk:"parser_limits_max_local_names"`
	ParserLimitsAttachmentByteCount        types.Int64                 `tfsdk:"parser_limits_attachment_byte_count"`
	ParserLimitsAttachmentPackageByteCount types.Int64                 `tfsdk:"parser_limits_attachment_package_byte_count"`
	ParserLimitsExternalReferences         types.String                `tfsdk:"parser_limits_external_references"`
	CredentialCharset                      types.String                `tfsdk:"credential_charset"`
	SslConfigType                          types.String                `tfsdk:"ssl_config_type"`
	SslServer                              types.String                `tfsdk:"ssl_server"`
	SslSniServer                           types.String                `tfsdk:"ssl_sni_server"`
	SslClient                              types.String                `tfsdk:"ssl_client"`
	UserSummary                            types.String                `tfsdk:"user_summary"`
	Priority                               types.String                `tfsdk:"priority"`
	LocalPort                              types.Int64                 `tfsdk:"local_port"`
	RemoteAddress                          types.String                `tfsdk:"remote_address"`
	RemotePort                             types.Int64                 `tfsdk:"remote_port"`
	Acl                                    types.String                `tfsdk:"acl"`
	HttpTimeout                            types.Int64                 `tfsdk:"http_timeout"`
	HttpPersistTimeout                     types.Int64                 `tfsdk:"http_persist_timeout"`
	DoHostRewrite                          types.Bool                  `tfsdk:"do_host_rewrite"`
	SuppressHttpWarnings                   types.Bool                  `tfsdk:"suppress_http_warnings"`
	HttpCompression                        types.Bool                  `tfsdk:"http_compression"`
	HttpIncludeResponseTypeEncoding        types.Bool                  `tfsdk:"http_include_response_type_encoding"`
	AlwaysShowErrors                       types.Bool                  `tfsdk:"always_show_errors"`
	DisallowGet                            types.Bool                  `tfsdk:"disallow_get"`
	DisallowEmptyResponse                  types.Bool                  `tfsdk:"disallow_empty_response"`
	HttpPersistentConnections              types.Bool                  `tfsdk:"http_persistent_connections"`
	HttpClientIpLabel                      types.String                `tfsdk:"http_client_ip_label"`
	HttpLogCorIdLabel                      types.String                `tfsdk:"http_log_cor_id_label"`
	HttpProxyHost                          types.String                `tfsdk:"http_proxy_host"`
	HttpProxyPort                          types.Int64                 `tfsdk:"http_proxy_port"`
	HttpVersion                            *DmHTTPClientServerVersion  `tfsdk:"http_version"`
	DoChunkedUpload                        types.Bool                  `tfsdk:"do_chunked_upload"`
	HeaderInjection                        types.List                  `tfsdk:"header_injection"`
	HeaderSuppression                      types.List                  `tfsdk:"header_suppression"`
	StylesheetParameters                   types.List                  `tfsdk:"stylesheet_parameters"`
	DefaultParamNamespace                  types.String                `tfsdk:"default_param_namespace"`
	QueryParamNamespace                    types.String                `tfsdk:"query_param_namespace"`
	ForcePolicyExec                        types.Bool                  `tfsdk:"force_policy_exec"`
	CountMonitors                          types.List                  `tfsdk:"count_monitors"`
	DurationMonitors                       types.List                  `tfsdk:"duration_monitors"`
	MonitorProcessingPolicy                types.String                `tfsdk:"monitor_processing_policy"`
	DebugMode                              types.String                `tfsdk:"debug_mode"`
	DebugHistory                           types.Int64                 `tfsdk:"debug_history"`
	DebugTrigger                           types.List                  `tfsdk:"debug_trigger"`
	LocalAddress                           types.String                `tfsdk:"local_address"`
	DependencyActions                      []*actions.DependencyAction `tfsdk:"dependency_actions"`
	ProviderTarget                         types.String                `tfsdk:"provider_target"`
}

var XMLFirewallServiceRequestAttachmentsIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "request_type",
	AttrType:    "String",
	AttrDefault: "soap",
	Value:       []string{"unprocessed"},
}

var XMLFirewallServiceResponseAttachmentsIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "response_type",
	AttrType:    "String",
	AttrDefault: "soap",
	Value:       []string{"unprocessed"},
}

var XMLFirewallServiceRootPartNotFirstActionIgnoreVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "request_attachments",
			AttrType:    "String",
			AttrDefault: "strip",
			Value:       []string{"streaming"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "response_attachments",
			AttrType:    "String",
			AttrDefault: "strip",
			Value:       []string{"streaming"},
		},
	},
}

var XMLFirewallServiceDelayErrorsIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "rewrite_errors",
	AttrType:    "Bool",
	AttrDefault: "true",
	Value:       []string{"true"},
}

var XMLFirewallServiceDelayErrorsDurationCondVal = validators.Evaluation{
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

var XMLFirewallServiceDelayErrorsDurationIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var XMLFirewallServiceSOAPSchemaURLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "request_type",
			AttrType:    "String",
			AttrDefault: "soap",
			Value:       []string{"soap"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "response_type",
			AttrType:    "String",
			AttrDefault: "soap",
			Value:       []string{"soap"},
		},
	},
}

var XMLFirewallServiceWSDLFileLocationIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "wsdl_response_policy",
	AttrType:    "String",
	AttrDefault: "off",
	Value:       []string{"serve"},
}

var XMLFirewallServiceParserLimitsBytesScannedCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "firewall_parser_limits",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var XMLFirewallServiceParserLimitsBytesScannedIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var XMLFirewallServiceParserLimitsElementDepthCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "firewall_parser_limits",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var XMLFirewallServiceParserLimitsElementDepthIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var XMLFirewallServiceParserLimitsAttributeCountCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "firewall_parser_limits",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var XMLFirewallServiceParserLimitsAttributeCountIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var XMLFirewallServiceParserLimitsMaxNodeSizeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "firewall_parser_limits",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var XMLFirewallServiceParserLimitsMaxNodeSizeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var XMLFirewallServiceParserLimitsMaxPrefixesCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "firewall_parser_limits",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var XMLFirewallServiceParserLimitsMaxPrefixesIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var XMLFirewallServiceParserLimitsMaxNamespacesCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "firewall_parser_limits",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var XMLFirewallServiceParserLimitsMaxNamespacesIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var XMLFirewallServiceParserLimitsMaxLocalNamesCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "firewall_parser_limits",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var XMLFirewallServiceParserLimitsMaxLocalNamesIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var XMLFirewallServiceParserLimitsAttachmentByteCountCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "firewall_parser_limits",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var XMLFirewallServiceParserLimitsAttachmentByteCountIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var XMLFirewallServiceParserLimitsAttachmentPackageByteCountCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "firewall_parser_limits",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var XMLFirewallServiceParserLimitsAttachmentPackageByteCountIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var XMLFirewallServiceParserLimitsExternalReferencesCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "firewall_parser_limits",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var XMLFirewallServiceParserLimitsExternalReferencesIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var XMLFirewallServiceSSLServerIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "ssl_config_type",
	AttrType:    "String",
	AttrDefault: "server",
	Value:       []string{"server"},
}

var XMLFirewallServiceSSLSNIServerIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "ssl_config_type",
	AttrType:    "String",
	AttrDefault: "server",
	Value:       []string{"sni"},
}

var XMLFirewallServiceSSLClientIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "ssl_config_type",
	AttrType:    "String",
	AttrDefault: "server",
	Value:       []string{"proxy"},
}

var XMLFirewallServiceRemoteAddressCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "dynamic-backend",
	Value:       []string{"static-backend"},
}

var XMLFirewallServiceRemoteAddressIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var XMLFirewallServiceRemotePortCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "dynamic-backend",
	Value:       []string{"static-backend"},
}

var XMLFirewallServiceRemotePortIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var XMLFirewallServiceDebugHistoryCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "debug_mode",
	AttrType:    "String",
	AttrDefault: "off",
	Value:       []string{"true"},
}

var XMLFirewallServiceDebugHistoryIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var XMLFirewallServiceDebugTriggerIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "debug_mode",
	AttrType:    "String",
	AttrDefault: "off",
	Value:       []string{"true"},
}

var XMLFirewallServiceObjectType = map[string]attr.Type{
	"provider_target":                             types.StringType,
	"id":                                          types.StringType,
	"app_domain":                                  types.StringType,
	"type":                                        types.StringType,
	"xml_manager":                                 types.StringType,
	"url_rewrite_policy":                          types.StringType,
	"style_policy":                                types.StringType,
	"max_message_size":                            types.Int64Type,
	"request_type":                                types.StringType,
	"response_type":                               types.StringType,
	"fw_cred":                                     types.StringType,
	"service_monitors":                            types.ListType{ElemType: types.StringType},
	"request_attachments":                         types.StringType,
	"response_attachments":                        types.StringType,
	"root_part_not_first_action":                  types.StringType,
	"front_attachment_format":                     types.StringType,
	"back_attachment_format":                      types.StringType,
	"mime_headers":                                types.BoolType,
	"rewrite_errors":                              types.BoolType,
	"delay_errors":                                types.BoolType,
	"delay_errors_duration":                       types.Int64Type,
	"soap_schema_url":                             types.StringType,
	"wsdl_response_policy":                        types.StringType,
	"wsdl_file_location":                          types.StringType,
	"firewall_parser_limits":                      types.BoolType,
	"parser_limits_bytes_scanned":                 types.Int64Type,
	"parser_limits_element_depth":                 types.Int64Type,
	"parser_limits_attribute_count":               types.Int64Type,
	"parser_limits_max_node_size":                 types.Int64Type,
	"parser_limits_max_prefixes":                  types.Int64Type,
	"parser_limits_max_namespaces":                types.Int64Type,
	"parser_limits_max_local_names":               types.Int64Type,
	"parser_limits_attachment_byte_count":         types.Int64Type,
	"parser_limits_attachment_package_byte_count": types.Int64Type,
	"parser_limits_external_references":           types.StringType,
	"credential_charset":                          types.StringType,
	"ssl_config_type":                             types.StringType,
	"ssl_server":                                  types.StringType,
	"ssl_sni_server":                              types.StringType,
	"ssl_client":                                  types.StringType,
	"user_summary":                                types.StringType,
	"priority":                                    types.StringType,
	"local_port":                                  types.Int64Type,
	"remote_address":                              types.StringType,
	"remote_port":                                 types.Int64Type,
	"acl":                                         types.StringType,
	"http_timeout":                                types.Int64Type,
	"http_persist_timeout":                        types.Int64Type,
	"do_host_rewrite":                             types.BoolType,
	"suppress_http_warnings":                      types.BoolType,
	"http_compression":                            types.BoolType,
	"http_include_response_type_encoding":         types.BoolType,
	"always_show_errors":                          types.BoolType,
	"disallow_get":                                types.BoolType,
	"disallow_empty_response":                     types.BoolType,
	"http_persistent_connections":                 types.BoolType,
	"http_client_ip_label":                        types.StringType,
	"http_log_cor_id_label":                       types.StringType,
	"http_proxy_host":                             types.StringType,
	"http_proxy_port":                             types.Int64Type,
	"http_version":                                types.ObjectType{AttrTypes: DmHTTPClientServerVersionObjectType},
	"do_chunked_upload":                           types.BoolType,
	"header_injection":                            types.ListType{ElemType: types.ObjectType{AttrTypes: DmHeaderInjectionObjectType}},
	"header_suppression":                          types.ListType{ElemType: types.ObjectType{AttrTypes: DmHeaderSuppressionObjectType}},
	"stylesheet_parameters":                       types.ListType{ElemType: types.ObjectType{AttrTypes: DmStylesheetParameterObjectType}},
	"default_param_namespace":                     types.StringType,
	"query_param_namespace":                       types.StringType,
	"force_policy_exec":                           types.BoolType,
	"count_monitors":                              types.ListType{ElemType: types.StringType},
	"duration_monitors":                           types.ListType{ElemType: types.StringType},
	"monitor_processing_policy":                   types.StringType,
	"debug_mode":                                  types.StringType,
	"debug_history":                               types.Int64Type,
	"debug_trigger":                               types.ListType{ElemType: types.ObjectType{AttrTypes: DmMSDebugTriggerTypeObjectType}},
	"local_address":                               types.StringType,
	"dependency_actions":                          actions.ActionsListType,
}

func (data XMLFirewallService) GetPath() string {
	rest_path := "/mgmt/config/{domain}/XMLFirewallService"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data XMLFirewallService) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
		return false
	}
	if !data.XmlManager.IsNull() {
		return false
	}
	if !data.UrlRewritePolicy.IsNull() {
		return false
	}
	if !data.StylePolicy.IsNull() {
		return false
	}
	if !data.MaxMessageSize.IsNull() {
		return false
	}
	if !data.RequestType.IsNull() {
		return false
	}
	if !data.ResponseType.IsNull() {
		return false
	}
	if !data.FwCred.IsNull() {
		return false
	}
	if !data.ServiceMonitors.IsNull() {
		return false
	}
	if !data.RequestAttachments.IsNull() {
		return false
	}
	if !data.ResponseAttachments.IsNull() {
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
	if !data.MimeHeaders.IsNull() {
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
	if !data.SoapSchemaUrl.IsNull() {
		return false
	}
	if !data.WsdlResponsePolicy.IsNull() {
		return false
	}
	if !data.WsdlFileLocation.IsNull() {
		return false
	}
	if !data.FirewallParserLimits.IsNull() {
		return false
	}
	if !data.ParserLimitsBytesScanned.IsNull() {
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
	if !data.ParserLimitsExternalReferences.IsNull() {
		return false
	}
	if !data.CredentialCharset.IsNull() {
		return false
	}
	if !data.SslConfigType.IsNull() {
		return false
	}
	if !data.SslServer.IsNull() {
		return false
	}
	if !data.SslSniServer.IsNull() {
		return false
	}
	if !data.SslClient.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Priority.IsNull() {
		return false
	}
	if !data.LocalPort.IsNull() {
		return false
	}
	if !data.RemoteAddress.IsNull() {
		return false
	}
	if !data.RemotePort.IsNull() {
		return false
	}
	if !data.Acl.IsNull() {
		return false
	}
	if !data.HttpTimeout.IsNull() {
		return false
	}
	if !data.HttpPersistTimeout.IsNull() {
		return false
	}
	if !data.DoHostRewrite.IsNull() {
		return false
	}
	if !data.SuppressHttpWarnings.IsNull() {
		return false
	}
	if !data.HttpCompression.IsNull() {
		return false
	}
	if !data.HttpIncludeResponseTypeEncoding.IsNull() {
		return false
	}
	if !data.AlwaysShowErrors.IsNull() {
		return false
	}
	if !data.DisallowGet.IsNull() {
		return false
	}
	if !data.DisallowEmptyResponse.IsNull() {
		return false
	}
	if !data.HttpPersistentConnections.IsNull() {
		return false
	}
	if !data.HttpClientIpLabel.IsNull() {
		return false
	}
	if !data.HttpLogCorIdLabel.IsNull() {
		return false
	}
	if !data.HttpProxyHost.IsNull() {
		return false
	}
	if !data.HttpProxyPort.IsNull() {
		return false
	}
	if data.HttpVersion != nil {
		if !data.HttpVersion.IsNull() {
			return false
		}
	}
	if !data.DoChunkedUpload.IsNull() {
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
	if !data.ForcePolicyExec.IsNull() {
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
	if !data.DebugMode.IsNull() {
		return false
	}
	if !data.DebugHistory.IsNull() {
		return false
	}
	if !data.DebugTrigger.IsNull() {
		return false
	}
	if !data.LocalAddress.IsNull() {
		return false
	}
	return true
}

func (data XMLFirewallService) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Type`, data.Type.ValueString())
	}
	if !data.XmlManager.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XMLManager`, data.XmlManager.ValueString())
	}
	if !data.UrlRewritePolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`URLRewritePolicy`, data.UrlRewritePolicy.ValueString())
	}
	if !data.StylePolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StylePolicy`, data.StylePolicy.ValueString())
	}
	if !data.MaxMessageSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxMessageSize`, data.MaxMessageSize.ValueInt64())
	}
	if !data.RequestType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RequestType`, data.RequestType.ValueString())
	}
	if !data.ResponseType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ResponseType`, data.ResponseType.ValueString())
	}
	if !data.FwCred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FWCred`, data.FwCred.ValueString())
	}
	if !data.ServiceMonitors.IsNull() {
		var dataValues []string
		data.ServiceMonitors.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.Set(body, pathRoot+`ServiceMonitors`+".-1", map[string]string{"value": val})
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`ServiceMonitors`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`ServiceMonitors`, "[]")
	}
	if !data.RequestAttachments.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RequestAttachments`, data.RequestAttachments.ValueString())
	}
	if !data.ResponseAttachments.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ResponseAttachments`, data.ResponseAttachments.ValueString())
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
	if !data.MimeHeaders.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MIMEHeaders`, tfutils.StringFromBool(data.MimeHeaders, ""))
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
	if !data.SoapSchemaUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SOAPSchemaURL`, data.SoapSchemaUrl.ValueString())
	}
	if !data.WsdlResponsePolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSDLResponsePolicy`, data.WsdlResponsePolicy.ValueString())
	}
	if !data.WsdlFileLocation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSDLFileLocation`, data.WsdlFileLocation.ValueString())
	}
	if !data.FirewallParserLimits.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FirewallParserLimits`, tfutils.StringFromBool(data.FirewallParserLimits, ""))
	}
	if !data.ParserLimitsBytesScanned.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ParserLimitsBytesScanned`, data.ParserLimitsBytesScanned.ValueInt64())
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
	if !data.ParserLimitsExternalReferences.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ParserLimitsExternalReferences`, data.ParserLimitsExternalReferences.ValueString())
	}
	if !data.CredentialCharset.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CredentialCharset`, data.CredentialCharset.ValueString())
	}
	if !data.SslConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLConfigType`, data.SslConfigType.ValueString())
	}
	if !data.SslServer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLServer`, data.SslServer.ValueString())
	}
	if !data.SslSniServer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLSNIServer`, data.SslSniServer.ValueString())
	}
	if !data.SslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClient`, data.SslClient.ValueString())
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.Priority.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Priority`, data.Priority.ValueString())
	}
	if !data.LocalPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalPort`, data.LocalPort.ValueInt64())
	}
	if !data.RemoteAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemoteAddress`, data.RemoteAddress.ValueString())
	}
	if !data.RemotePort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemotePort`, data.RemotePort.ValueInt64())
	}
	if !data.Acl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ACL`, data.Acl.ValueString())
	}
	if !data.HttpTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTPTimeout`, data.HttpTimeout.ValueInt64())
	}
	if !data.HttpPersistTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTPPersistTimeout`, data.HttpPersistTimeout.ValueInt64())
	}
	if !data.DoHostRewrite.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DoHostRewrite`, tfutils.StringFromBool(data.DoHostRewrite, ""))
	}
	if !data.SuppressHttpWarnings.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SuppressHTTPWarnings`, tfutils.StringFromBool(data.SuppressHttpWarnings, ""))
	}
	if !data.HttpCompression.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTPCompression`, tfutils.StringFromBool(data.HttpCompression, ""))
	}
	if !data.HttpIncludeResponseTypeEncoding.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTPIncludeResponseTypeEncoding`, tfutils.StringFromBool(data.HttpIncludeResponseTypeEncoding, ""))
	}
	if !data.AlwaysShowErrors.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AlwaysShowErrors`, tfutils.StringFromBool(data.AlwaysShowErrors, ""))
	}
	if !data.DisallowGet.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DisallowGet`, tfutils.StringFromBool(data.DisallowGet, ""))
	}
	if !data.DisallowEmptyResponse.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DisallowEmptyResponse`, tfutils.StringFromBool(data.DisallowEmptyResponse, ""))
	}
	if !data.HttpPersistentConnections.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTPPersistentConnections`, tfutils.StringFromBool(data.HttpPersistentConnections, ""))
	}
	if !data.HttpClientIpLabel.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTPClientIPLabel`, data.HttpClientIpLabel.ValueString())
	}
	if !data.HttpLogCorIdLabel.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTPLogCorIDLabel`, data.HttpLogCorIdLabel.ValueString())
	}
	if !data.HttpProxyHost.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTPProxyHost`, data.HttpProxyHost.ValueString())
	}
	if !data.HttpProxyPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTPProxyPort`, data.HttpProxyPort.ValueInt64())
	}
	if data.HttpVersion != nil {
		if !data.HttpVersion.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`HTTPVersion`, data.HttpVersion.ToBody(ctx, ""))
		}
	}
	if !data.DoChunkedUpload.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DoChunkedUpload`, tfutils.StringFromBool(data.DoChunkedUpload, ""))
	}
	if !data.HeaderInjection.IsNull() {
		var dataValues []DmHeaderInjection
		data.HeaderInjection.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.SetRaw(body, pathRoot+`HeaderInjection`+".-1", val.ToBody(ctx, ""))
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`HeaderInjection`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`HeaderInjection`, "[]")
	}
	if !data.HeaderSuppression.IsNull() {
		var dataValues []DmHeaderSuppression
		data.HeaderSuppression.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.SetRaw(body, pathRoot+`HeaderSuppression`+".-1", val.ToBody(ctx, ""))
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`HeaderSuppression`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`HeaderSuppression`, "[]")
	}
	if !data.StylesheetParameters.IsNull() {
		var dataValues []DmStylesheetParameter
		data.StylesheetParameters.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.SetRaw(body, pathRoot+`StylesheetParameters`+".-1", val.ToBody(ctx, ""))
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`StylesheetParameters`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`StylesheetParameters`, "[]")
	}
	if !data.DefaultParamNamespace.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DefaultParamNamespace`, data.DefaultParamNamespace.ValueString())
	}
	if !data.QueryParamNamespace.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`QueryParamNamespace`, data.QueryParamNamespace.ValueString())
	}
	if !data.ForcePolicyExec.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ForcePolicyExec`, tfutils.StringFromBool(data.ForcePolicyExec, ""))
	}
	if !data.CountMonitors.IsNull() {
		var dataValues []string
		data.CountMonitors.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.Set(body, pathRoot+`CountMonitors`+".-1", map[string]string{"value": val})
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`CountMonitors`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`CountMonitors`, "[]")
	}
	if !data.DurationMonitors.IsNull() {
		var dataValues []string
		data.DurationMonitors.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.Set(body, pathRoot+`DurationMonitors`+".-1", map[string]string{"value": val})
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`DurationMonitors`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`DurationMonitors`, "[]")
	}
	if !data.MonitorProcessingPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MonitorProcessingPolicy`, data.MonitorProcessingPolicy.ValueString())
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
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.SetRaw(body, pathRoot+`DebugTrigger`+".-1", val.ToBody(ctx, ""))
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`DebugTrigger`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`DebugTrigger`, "[]")
	}
	if !data.LocalAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalAddress`, data.LocalAddress.ValueString())
	}
	return body
}

func (data *XMLFirewallService) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringValue("dynamic-backend")
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
	if value := res.Get(pathRoot + `StylePolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.StylePolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.StylePolicy = types.StringValue("default")
	}
	if value := res.Get(pathRoot + `MaxMessageSize`); value.Exists() {
		data.MaxMessageSize = types.Int64Value(value.Int())
	} else {
		data.MaxMessageSize = types.Int64Value(0)
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
	if value := res.Get(pathRoot + `FWCred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.FwCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FwCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `ServiceMonitors`); value.Exists() {
		data.ServiceMonitors = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.ServiceMonitors = types.ListNull(types.StringType)
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
	if value := res.Get(pathRoot + `MIMEHeaders`); value.Exists() {
		data.MimeHeaders = tfutils.BoolFromString(value.String())
	} else {
		data.MimeHeaders = types.BoolNull()
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
	if value := res.Get(pathRoot + `SOAPSchemaURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SoapSchemaUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SoapSchemaUrl = types.StringValue("store:///schemas/soap-envelope.xsd")
	}
	if value := res.Get(pathRoot + `WSDLResponsePolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsdlResponsePolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlResponsePolicy = types.StringValue("off")
	}
	if value := res.Get(pathRoot + `WSDLFileLocation`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsdlFileLocation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlFileLocation = types.StringNull()
	}
	if value := res.Get(pathRoot + `FirewallParserLimits`); value.Exists() {
		data.FirewallParserLimits = tfutils.BoolFromString(value.String())
	} else {
		data.FirewallParserLimits = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ParserLimitsBytesScanned`); value.Exists() {
		data.ParserLimitsBytesScanned = types.Int64Value(value.Int())
	} else {
		data.ParserLimitsBytesScanned = types.Int64Value(4194304)
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
	if value := res.Get(pathRoot + `ParserLimitsExternalReferences`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ParserLimitsExternalReferences = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ParserLimitsExternalReferences = types.StringValue("forbid")
	}
	if value := res.Get(pathRoot + `CredentialCharset`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CredentialCharset = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CredentialCharset = types.StringValue("protocol")
	}
	if value := res.Get(pathRoot + `SSLConfigType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslConfigType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslConfigType = types.StringValue("server")
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
	if value := res.Get(pathRoot + `SSLClient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClient = types.StringNull()
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
	if value := res.Get(pathRoot + `LocalPort`); value.Exists() {
		data.LocalPort = types.Int64Value(value.Int())
	} else {
		data.LocalPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `RemoteAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RemoteAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemotePort`); value.Exists() {
		data.RemotePort = types.Int64Value(value.Int())
	} else {
		data.RemotePort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ACL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Acl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Acl = types.StringNull()
	}
	if value := res.Get(pathRoot + `HTTPTimeout`); value.Exists() {
		data.HttpTimeout = types.Int64Value(value.Int())
	} else {
		data.HttpTimeout = types.Int64Value(120)
	}
	if value := res.Get(pathRoot + `HTTPPersistTimeout`); value.Exists() {
		data.HttpPersistTimeout = types.Int64Value(value.Int())
	} else {
		data.HttpPersistTimeout = types.Int64Value(180)
	}
	if value := res.Get(pathRoot + `DoHostRewrite`); value.Exists() {
		data.DoHostRewrite = tfutils.BoolFromString(value.String())
	} else {
		data.DoHostRewrite = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SuppressHTTPWarnings`); value.Exists() {
		data.SuppressHttpWarnings = tfutils.BoolFromString(value.String())
	} else {
		data.SuppressHttpWarnings = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTPCompression`); value.Exists() {
		data.HttpCompression = tfutils.BoolFromString(value.String())
	} else {
		data.HttpCompression = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTPIncludeResponseTypeEncoding`); value.Exists() {
		data.HttpIncludeResponseTypeEncoding = tfutils.BoolFromString(value.String())
	} else {
		data.HttpIncludeResponseTypeEncoding = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AlwaysShowErrors`); value.Exists() {
		data.AlwaysShowErrors = tfutils.BoolFromString(value.String())
	} else {
		data.AlwaysShowErrors = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DisallowGet`); value.Exists() {
		data.DisallowGet = tfutils.BoolFromString(value.String())
	} else {
		data.DisallowGet = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DisallowEmptyResponse`); value.Exists() {
		data.DisallowEmptyResponse = tfutils.BoolFromString(value.String())
	} else {
		data.DisallowEmptyResponse = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTPPersistentConnections`); value.Exists() {
		data.HttpPersistentConnections = tfutils.BoolFromString(value.String())
	} else {
		data.HttpPersistentConnections = types.BoolNull()
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
	if value := res.Get(pathRoot + `HTTPProxyHost`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HttpProxyHost = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HttpProxyHost = types.StringNull()
	}
	if value := res.Get(pathRoot + `HTTPProxyPort`); value.Exists() {
		data.HttpProxyPort = types.Int64Value(value.Int())
	} else {
		data.HttpProxyPort = types.Int64Value(800)
	}
	if value := res.Get(pathRoot + `HTTPVersion`); value.Exists() {
		data.HttpVersion = &DmHTTPClientServerVersion{}
		data.HttpVersion.FromBody(ctx, "", value)
	} else {
		data.HttpVersion = nil
	}
	if value := res.Get(pathRoot + `DoChunkedUpload`); value.Exists() {
		data.DoChunkedUpload = tfutils.BoolFromString(value.String())
	} else {
		data.DoChunkedUpload = types.BoolNull()
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
	if value := res.Get(pathRoot + `ForcePolicyExec`); value.Exists() {
		data.ForcePolicyExec = tfutils.BoolFromString(value.String())
	} else {
		data.ForcePolicyExec = types.BoolNull()
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
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalAddress = types.StringValue("0.0.0.0")
	}
}

func (data *XMLFirewallService) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && !data.Type.IsNull() {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else if data.Type.ValueString() != "dynamic-backend" {
		data.Type = types.StringNull()
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
	if value := res.Get(pathRoot + `StylePolicy`); value.Exists() && !data.StylePolicy.IsNull() {
		data.StylePolicy = tfutils.ParseStringFromGJSON(value)
	} else if data.StylePolicy.ValueString() != "default" {
		data.StylePolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `MaxMessageSize`); value.Exists() && !data.MaxMessageSize.IsNull() {
		data.MaxMessageSize = types.Int64Value(value.Int())
	} else if data.MaxMessageSize.ValueInt64() != 0 {
		data.MaxMessageSize = types.Int64Null()
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
	if value := res.Get(pathRoot + `FWCred`); value.Exists() && !data.FwCred.IsNull() {
		data.FwCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FwCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `ServiceMonitors`); value.Exists() && !data.ServiceMonitors.IsNull() {
		data.ServiceMonitors = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.ServiceMonitors = types.ListNull(types.StringType)
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
	if value := res.Get(pathRoot + `MIMEHeaders`); value.Exists() && !data.MimeHeaders.IsNull() {
		data.MimeHeaders = tfutils.BoolFromString(value.String())
	} else if !data.MimeHeaders.ValueBool() {
		data.MimeHeaders = types.BoolNull()
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
	if value := res.Get(pathRoot + `SOAPSchemaURL`); value.Exists() && !data.SoapSchemaUrl.IsNull() {
		data.SoapSchemaUrl = tfutils.ParseStringFromGJSON(value)
	} else if data.SoapSchemaUrl.ValueString() != "store:///schemas/soap-envelope.xsd" {
		data.SoapSchemaUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLResponsePolicy`); value.Exists() && !data.WsdlResponsePolicy.IsNull() {
		data.WsdlResponsePolicy = tfutils.ParseStringFromGJSON(value)
	} else if data.WsdlResponsePolicy.ValueString() != "off" {
		data.WsdlResponsePolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLFileLocation`); value.Exists() && !data.WsdlFileLocation.IsNull() {
		data.WsdlFileLocation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlFileLocation = types.StringNull()
	}
	if value := res.Get(pathRoot + `FirewallParserLimits`); value.Exists() && !data.FirewallParserLimits.IsNull() {
		data.FirewallParserLimits = tfutils.BoolFromString(value.String())
	} else if data.FirewallParserLimits.ValueBool() {
		data.FirewallParserLimits = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ParserLimitsBytesScanned`); value.Exists() && !data.ParserLimitsBytesScanned.IsNull() {
		data.ParserLimitsBytesScanned = types.Int64Value(value.Int())
	} else if data.ParserLimitsBytesScanned.ValueInt64() != 4194304 {
		data.ParserLimitsBytesScanned = types.Int64Null()
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
	if value := res.Get(pathRoot + `ParserLimitsExternalReferences`); value.Exists() && !data.ParserLimitsExternalReferences.IsNull() {
		data.ParserLimitsExternalReferences = tfutils.ParseStringFromGJSON(value)
	} else if data.ParserLimitsExternalReferences.ValueString() != "forbid" {
		data.ParserLimitsExternalReferences = types.StringNull()
	}
	if value := res.Get(pathRoot + `CredentialCharset`); value.Exists() && !data.CredentialCharset.IsNull() {
		data.CredentialCharset = tfutils.ParseStringFromGJSON(value)
	} else if data.CredentialCharset.ValueString() != "protocol" {
		data.CredentialCharset = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLConfigType`); value.Exists() && !data.SslConfigType.IsNull() {
		data.SslConfigType = tfutils.ParseStringFromGJSON(value)
	} else if data.SslConfigType.ValueString() != "server" {
		data.SslConfigType = types.StringNull()
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
	if value := res.Get(pathRoot + `SSLClient`); value.Exists() && !data.SslClient.IsNull() {
		data.SslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClient = types.StringNull()
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
	if value := res.Get(pathRoot + `LocalPort`); value.Exists() && !data.LocalPort.IsNull() {
		data.LocalPort = types.Int64Value(value.Int())
	} else {
		data.LocalPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `RemoteAddress`); value.Exists() && !data.RemoteAddress.IsNull() {
		data.RemoteAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemotePort`); value.Exists() && !data.RemotePort.IsNull() {
		data.RemotePort = types.Int64Value(value.Int())
	} else {
		data.RemotePort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ACL`); value.Exists() && !data.Acl.IsNull() {
		data.Acl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Acl = types.StringNull()
	}
	if value := res.Get(pathRoot + `HTTPTimeout`); value.Exists() && !data.HttpTimeout.IsNull() {
		data.HttpTimeout = types.Int64Value(value.Int())
	} else if data.HttpTimeout.ValueInt64() != 120 {
		data.HttpTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `HTTPPersistTimeout`); value.Exists() && !data.HttpPersistTimeout.IsNull() {
		data.HttpPersistTimeout = types.Int64Value(value.Int())
	} else if data.HttpPersistTimeout.ValueInt64() != 180 {
		data.HttpPersistTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `DoHostRewrite`); value.Exists() && !data.DoHostRewrite.IsNull() {
		data.DoHostRewrite = tfutils.BoolFromString(value.String())
	} else if !data.DoHostRewrite.ValueBool() {
		data.DoHostRewrite = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SuppressHTTPWarnings`); value.Exists() && !data.SuppressHttpWarnings.IsNull() {
		data.SuppressHttpWarnings = tfutils.BoolFromString(value.String())
	} else if data.SuppressHttpWarnings.ValueBool() {
		data.SuppressHttpWarnings = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTPCompression`); value.Exists() && !data.HttpCompression.IsNull() {
		data.HttpCompression = tfutils.BoolFromString(value.String())
	} else if data.HttpCompression.ValueBool() {
		data.HttpCompression = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTPIncludeResponseTypeEncoding`); value.Exists() && !data.HttpIncludeResponseTypeEncoding.IsNull() {
		data.HttpIncludeResponseTypeEncoding = tfutils.BoolFromString(value.String())
	} else if data.HttpIncludeResponseTypeEncoding.ValueBool() {
		data.HttpIncludeResponseTypeEncoding = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AlwaysShowErrors`); value.Exists() && !data.AlwaysShowErrors.IsNull() {
		data.AlwaysShowErrors = tfutils.BoolFromString(value.String())
	} else if data.AlwaysShowErrors.ValueBool() {
		data.AlwaysShowErrors = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DisallowGet`); value.Exists() && !data.DisallowGet.IsNull() {
		data.DisallowGet = tfutils.BoolFromString(value.String())
	} else if data.DisallowGet.ValueBool() {
		data.DisallowGet = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DisallowEmptyResponse`); value.Exists() && !data.DisallowEmptyResponse.IsNull() {
		data.DisallowEmptyResponse = tfutils.BoolFromString(value.String())
	} else if data.DisallowEmptyResponse.ValueBool() {
		data.DisallowEmptyResponse = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTPPersistentConnections`); value.Exists() && !data.HttpPersistentConnections.IsNull() {
		data.HttpPersistentConnections = tfutils.BoolFromString(value.String())
	} else if !data.HttpPersistentConnections.ValueBool() {
		data.HttpPersistentConnections = types.BoolNull()
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
	if value := res.Get(pathRoot + `HTTPProxyHost`); value.Exists() && !data.HttpProxyHost.IsNull() {
		data.HttpProxyHost = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HttpProxyHost = types.StringNull()
	}
	if value := res.Get(pathRoot + `HTTPProxyPort`); value.Exists() && !data.HttpProxyPort.IsNull() {
		data.HttpProxyPort = types.Int64Value(value.Int())
	} else if data.HttpProxyPort.ValueInt64() != 800 {
		data.HttpProxyPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `HTTPVersion`); value.Exists() {
		data.HttpVersion.UpdateFromBody(ctx, "", value)
	} else {
		data.HttpVersion = nil
	}
	if value := res.Get(pathRoot + `DoChunkedUpload`); value.Exists() && !data.DoChunkedUpload.IsNull() {
		data.DoChunkedUpload = tfutils.BoolFromString(value.String())
	} else if data.DoChunkedUpload.ValueBool() {
		data.DoChunkedUpload = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HeaderInjection`); value.Exists() && !data.HeaderInjection.IsNull() {
		l := []DmHeaderInjection{}
		e := []DmHeaderInjection{}
		data.HeaderInjection.ElementsAs(ctx, &e, false)
		if len(value.Array()) == len(e) {
			for i, v := range value.Array() {
				item := e[i]
				item.UpdateFromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		} else {
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
	if value := res.Get(pathRoot + `HeaderSuppression`); value.Exists() && !data.HeaderSuppression.IsNull() {
		l := []DmHeaderSuppression{}
		e := []DmHeaderSuppression{}
		data.HeaderSuppression.ElementsAs(ctx, &e, false)
		if len(value.Array()) == len(e) {
			for i, v := range value.Array() {
				item := e[i]
				item.UpdateFromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		} else {
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
	if value := res.Get(pathRoot + `StylesheetParameters`); value.Exists() && !data.StylesheetParameters.IsNull() {
		l := []DmStylesheetParameter{}
		e := []DmStylesheetParameter{}
		data.StylesheetParameters.ElementsAs(ctx, &e, false)
		if len(value.Array()) == len(e) {
			for i, v := range value.Array() {
				item := e[i]
				item.UpdateFromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		} else {
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
	if value := res.Get(pathRoot + `ForcePolicyExec`); value.Exists() && !data.ForcePolicyExec.IsNull() {
		data.ForcePolicyExec = tfutils.BoolFromString(value.String())
	} else if data.ForcePolicyExec.ValueBool() {
		data.ForcePolicyExec = types.BoolNull()
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
		e := []DmMSDebugTriggerType{}
		data.DebugTrigger.ElementsAs(ctx, &e, false)
		if len(value.Array()) == len(e) {
			for i, v := range value.Array() {
				item := e[i]
				item.UpdateFromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		} else {
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
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && !data.LocalAddress.IsNull() {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else if data.LocalAddress.ValueString() != "0.0.0.0" {
		data.LocalAddress = types.StringNull()
	}
}
