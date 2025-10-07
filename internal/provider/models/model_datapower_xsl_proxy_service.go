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

type XSLProxyService struct {
	Id                              types.String                `tfsdk:"id"`
	AppDomain                       types.String                `tfsdk:"app_domain"`
	Type                            types.String                `tfsdk:"type"`
	XmlManager                      types.String                `tfsdk:"xml_manager"`
	UrlRewritePolicy                types.String                `tfsdk:"url_rewrite_policy"`
	StylePolicy                     types.String                `tfsdk:"style_policy"`
	CredentialCharset               types.String                `tfsdk:"credential_charset"`
	SslConfigType                   types.String                `tfsdk:"ssl_config_type"`
	SslServer                       types.String                `tfsdk:"ssl_server"`
	SslSniServer                    types.String                `tfsdk:"ssl_sni_server"`
	SslClient                       types.String                `tfsdk:"ssl_client"`
	UserSummary                     types.String                `tfsdk:"user_summary"`
	Priority                        types.String                `tfsdk:"priority"`
	LocalPort                       types.Int64                 `tfsdk:"local_port"`
	RemoteAddress                   types.String                `tfsdk:"remote_address"`
	RemotePort                      types.Int64                 `tfsdk:"remote_port"`
	Acl                             types.String                `tfsdk:"acl"`
	HttpTimeout                     types.Int64                 `tfsdk:"http_timeout"`
	HttpPersistTimeout              types.Int64                 `tfsdk:"http_persist_timeout"`
	DoHostRewrite                   types.Bool                  `tfsdk:"do_host_rewrite"`
	SuppressHttpWarnings            types.Bool                  `tfsdk:"suppress_http_warnings"`
	HttpCompression                 types.Bool                  `tfsdk:"http_compression"`
	HttpIncludeResponseTypeEncoding types.Bool                  `tfsdk:"http_include_response_type_encoding"`
	AlwaysShowErrors                types.Bool                  `tfsdk:"always_show_errors"`
	DisallowGet                     types.Bool                  `tfsdk:"disallow_get"`
	DisallowEmptyResponse           types.Bool                  `tfsdk:"disallow_empty_response"`
	HttpPersistentConnections       types.Bool                  `tfsdk:"http_persistent_connections"`
	HttpClientIpLabel               types.String                `tfsdk:"http_client_ip_label"`
	HttpLogCorIdLabel               types.String                `tfsdk:"http_log_cor_id_label"`
	HttpProxyHost                   types.String                `tfsdk:"http_proxy_host"`
	HttpProxyPort                   types.Int64                 `tfsdk:"http_proxy_port"`
	HttpVersion                     *DmHTTPClientServerVersion  `tfsdk:"http_version"`
	DoChunkedUpload                 types.Bool                  `tfsdk:"do_chunked_upload"`
	HeaderInjection                 types.List                  `tfsdk:"header_injection"`
	HeaderSuppression               types.List                  `tfsdk:"header_suppression"`
	StylesheetParameters            types.List                  `tfsdk:"stylesheet_parameters"`
	DefaultParamNamespace           types.String                `tfsdk:"default_param_namespace"`
	QueryParamNamespace             types.String                `tfsdk:"query_param_namespace"`
	ForcePolicyExec                 types.Bool                  `tfsdk:"force_policy_exec"`
	CountMonitors                   types.List                  `tfsdk:"count_monitors"`
	DurationMonitors                types.List                  `tfsdk:"duration_monitors"`
	MonitorProcessingPolicy         types.String                `tfsdk:"monitor_processing_policy"`
	DebugMode                       types.String                `tfsdk:"debug_mode"`
	DebugHistory                    types.Int64                 `tfsdk:"debug_history"`
	DebugTrigger                    types.List                  `tfsdk:"debug_trigger"`
	LocalAddress                    types.String                `tfsdk:"local_address"`
	DependencyActions               []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var XSLProxyServiceRemoteAddressCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "static-backend",
	Value:       []string{"static-backend"},
}
var XSLProxyServiceRemotePortCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "static-backend",
	Value:       []string{"static-backend"},
}
var XSLProxyServiceDebugHistoryCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "debug_mode",
	AttrType:    "String",
	AttrDefault: "off",
	Value:       []string{"true"},
}
var XSLProxyServiceSSLServerIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "ssl_config_type",
	AttrType:    "String",
	AttrDefault: "server",
	Value:       []string{"server"},
}
var XSLProxyServiceSSLSNIServerIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "ssl_config_type",
	AttrType:    "String",
	AttrDefault: "server",
	Value:       []string{"sni"},
}
var XSLProxyServiceSSLClientIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "ssl_config_type",
	AttrType:    "String",
	AttrDefault: "server",
	Value:       []string{"proxy"},
}
var XSLProxyServiceRemoteAddressIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var XSLProxyServiceRemotePortIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var XSLProxyServiceDebugHistoryIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var XSLProxyServiceDebugTriggerIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "debug_mode",
	AttrType:    "String",
	AttrDefault: "off",
	Value:       []string{"true"},
}

var XSLProxyServiceObjectType = map[string]attr.Type{
	"id":                                  types.StringType,
	"app_domain":                          types.StringType,
	"type":                                types.StringType,
	"xml_manager":                         types.StringType,
	"url_rewrite_policy":                  types.StringType,
	"style_policy":                        types.StringType,
	"credential_charset":                  types.StringType,
	"ssl_config_type":                     types.StringType,
	"ssl_server":                          types.StringType,
	"ssl_sni_server":                      types.StringType,
	"ssl_client":                          types.StringType,
	"user_summary":                        types.StringType,
	"priority":                            types.StringType,
	"local_port":                          types.Int64Type,
	"remote_address":                      types.StringType,
	"remote_port":                         types.Int64Type,
	"acl":                                 types.StringType,
	"http_timeout":                        types.Int64Type,
	"http_persist_timeout":                types.Int64Type,
	"do_host_rewrite":                     types.BoolType,
	"suppress_http_warnings":              types.BoolType,
	"http_compression":                    types.BoolType,
	"http_include_response_type_encoding": types.BoolType,
	"always_show_errors":                  types.BoolType,
	"disallow_get":                        types.BoolType,
	"disallow_empty_response":             types.BoolType,
	"http_persistent_connections":         types.BoolType,
	"http_client_ip_label":                types.StringType,
	"http_log_cor_id_label":               types.StringType,
	"http_proxy_host":                     types.StringType,
	"http_proxy_port":                     types.Int64Type,
	"http_version":                        types.ObjectType{AttrTypes: DmHTTPClientServerVersionObjectType},
	"do_chunked_upload":                   types.BoolType,
	"header_injection":                    types.ListType{ElemType: types.ObjectType{AttrTypes: DmHeaderInjectionObjectType}},
	"header_suppression":                  types.ListType{ElemType: types.ObjectType{AttrTypes: DmHeaderSuppressionObjectType}},
	"stylesheet_parameters":               types.ListType{ElemType: types.ObjectType{AttrTypes: DmStylesheetParameterObjectType}},
	"default_param_namespace":             types.StringType,
	"query_param_namespace":               types.StringType,
	"force_policy_exec":                   types.BoolType,
	"count_monitors":                      types.ListType{ElemType: types.StringType},
	"duration_monitors":                   types.ListType{ElemType: types.StringType},
	"monitor_processing_policy":           types.StringType,
	"debug_mode":                          types.StringType,
	"debug_history":                       types.Int64Type,
	"debug_trigger":                       types.ListType{ElemType: types.ObjectType{AttrTypes: DmMSDebugTriggerTypeObjectType}},
	"local_address":                       types.StringType,
	"dependency_actions":                  actions.ActionsListType,
}

func (data XSLProxyService) GetPath() string {
	rest_path := "/mgmt/config/{domain}/XSLProxyService"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data XSLProxyService) IsNull() bool {
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

func (data XSLProxyService) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.ForcePolicyExec.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ForcePolicyExec`, tfutils.StringFromBool(data.ForcePolicyExec, ""))
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
	if !data.LocalAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalAddress`, data.LocalAddress.ValueString())
	}
	return body
}

func (data *XSLProxyService) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
		data.Type = types.StringValue("static-backend")
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

func (data *XSLProxyService) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	} else if data.Type.ValueString() != "static-backend" {
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
