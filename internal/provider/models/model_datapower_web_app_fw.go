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

type WebAppFW struct {
	Id                      types.String                `tfsdk:"id"`
	AppDomain               types.String                `tfsdk:"app_domain"`
	UserSummary             types.String                `tfsdk:"user_summary"`
	Priority                types.String                `tfsdk:"priority"`
	FrontSide               types.List                  `tfsdk:"front_side"`
	RemoteAddress           types.String                `tfsdk:"remote_address"`
	RemotePort              types.Int64                 `tfsdk:"remote_port"`
	StylePolicy             types.String                `tfsdk:"style_policy"`
	XmlManager              types.String                `tfsdk:"xml_manager"`
	ErrorPolicy             types.String                `tfsdk:"error_policy"`
	UriNormalization        types.Bool                  `tfsdk:"uri_normalization"`
	RewriteErrors           types.Bool                  `tfsdk:"rewrite_errors"`
	DelayErrors             types.Bool                  `tfsdk:"delay_errors"`
	DelayErrorsDuration     types.Int64                 `tfsdk:"delay_errors_duration"`
	StreamOutputToBack      types.String                `tfsdk:"stream_output_to_back"`
	StreamOutputToFront     types.String                `tfsdk:"stream_output_to_front"`
	FrontTimeout            types.Int64                 `tfsdk:"front_timeout"`
	BackTimeout             types.Int64                 `tfsdk:"back_timeout"`
	FrontPersistentTimeout  types.Int64                 `tfsdk:"front_persistent_timeout"`
	AllowCacheControlHeader types.Bool                  `tfsdk:"allow_cache_control_header"`
	BackPersistentTimeout   types.Int64                 `tfsdk:"back_persistent_timeout"`
	FrontHttpVersion        types.String                `tfsdk:"front_http_version"`
	BackHttpVersion         types.String                `tfsdk:"back_http_version"`
	RequestSideSecurity     types.Bool                  `tfsdk:"request_side_security"`
	ResponseSideSecurity    types.Bool                  `tfsdk:"response_side_security"`
	DoChunkedUpload         types.Bool                  `tfsdk:"do_chunked_upload"`
	FollowRedirects         types.Bool                  `tfsdk:"follow_redirects"`
	HttpClientIpLabel       types.String                `tfsdk:"http_client_ip_label"`
	HttpLogCorIdLabel       types.String                `tfsdk:"http_log_cor_id_label"`
	DebugMode               types.String                `tfsdk:"debug_mode"`
	DebugHistory            types.Int64                 `tfsdk:"debug_history"`
	DebugTrigger            types.List                  `tfsdk:"debug_trigger"`
	UrlRewritePolicy        types.String                `tfsdk:"url_rewrite_policy"`
	DoHostRewriting         types.Bool                  `tfsdk:"do_host_rewriting"`
	SslConfigType           types.String                `tfsdk:"ssl_config_type"`
	SslServer               types.String                `tfsdk:"ssl_server"`
	SslSniServer            types.String                `tfsdk:"ssl_sni_server"`
	SslClient               types.String                `tfsdk:"ssl_client"`
	DependencyActions       []*actions.DependencyAction `tfsdk:"dependency_actions"`
	ProviderTarget          types.String                `tfsdk:"provider_target"`
}

var WebAppFWDelayErrorsIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "rewrite_errors",
	AttrType:    "Bool",
	AttrDefault: "true",
	Value:       []string{"true"},
}

var WebAppFWDelayErrorsDurationCondVal = validators.Evaluation{
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

var WebAppFWDelayErrorsDurationIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var WebAppFWDebugHistoryCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "debug_mode",
	AttrType:    "String",
	AttrDefault: "off",
	Value:       []string{"true"},
}

var WebAppFWDebugHistoryIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var WebAppFWDebugTriggerIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "debug_mode",
	AttrType:    "String",
	AttrDefault: "off",
	Value:       []string{"true"},
}

var WebAppFWSSLServerIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "ssl_config_type",
	AttrType:    "String",
	AttrDefault: "server",
	Value:       []string{"server"},
}

var WebAppFWSSLSNIServerIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "ssl_config_type",
	AttrType:    "String",
	AttrDefault: "server",
	Value:       []string{"sni"},
}

var WebAppFWSSLClientIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "ssl_config_type",
	AttrType:    "String",
	AttrDefault: "server",
	Value:       []string{"proxy"},
}

var WebAppFWObjectType = map[string]attr.Type{
	"provider_target":            types.StringType,
	"id":                         types.StringType,
	"app_domain":                 types.StringType,
	"user_summary":               types.StringType,
	"priority":                   types.StringType,
	"front_side":                 types.ListType{ElemType: types.ObjectType{AttrTypes: DmFrontSideObjectType}},
	"remote_address":             types.StringType,
	"remote_port":                types.Int64Type,
	"style_policy":               types.StringType,
	"xml_manager":                types.StringType,
	"error_policy":               types.StringType,
	"uri_normalization":          types.BoolType,
	"rewrite_errors":             types.BoolType,
	"delay_errors":               types.BoolType,
	"delay_errors_duration":      types.Int64Type,
	"stream_output_to_back":      types.StringType,
	"stream_output_to_front":     types.StringType,
	"front_timeout":              types.Int64Type,
	"back_timeout":               types.Int64Type,
	"front_persistent_timeout":   types.Int64Type,
	"allow_cache_control_header": types.BoolType,
	"back_persistent_timeout":    types.Int64Type,
	"front_http_version":         types.StringType,
	"back_http_version":          types.StringType,
	"request_side_security":      types.BoolType,
	"response_side_security":     types.BoolType,
	"do_chunked_upload":          types.BoolType,
	"follow_redirects":           types.BoolType,
	"http_client_ip_label":       types.StringType,
	"http_log_cor_id_label":      types.StringType,
	"debug_mode":                 types.StringType,
	"debug_history":              types.Int64Type,
	"debug_trigger":              types.ListType{ElemType: types.ObjectType{AttrTypes: DmMSDebugTriggerTypeObjectType}},
	"url_rewrite_policy":         types.StringType,
	"do_host_rewriting":          types.BoolType,
	"ssl_config_type":            types.StringType,
	"ssl_server":                 types.StringType,
	"ssl_sni_server":             types.StringType,
	"ssl_client":                 types.StringType,
	"dependency_actions":         actions.ActionsListType,
}

func (data WebAppFW) GetPath() string {
	rest_path := "/mgmt/config/{domain}/WebAppFW"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data WebAppFW) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Priority.IsNull() {
		return false
	}
	if !data.FrontSide.IsNull() {
		return false
	}
	if !data.RemoteAddress.IsNull() {
		return false
	}
	if !data.RemotePort.IsNull() {
		return false
	}
	if !data.StylePolicy.IsNull() {
		return false
	}
	if !data.XmlManager.IsNull() {
		return false
	}
	if !data.ErrorPolicy.IsNull() {
		return false
	}
	if !data.UriNormalization.IsNull() {
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
	if !data.StreamOutputToBack.IsNull() {
		return false
	}
	if !data.StreamOutputToFront.IsNull() {
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
	if !data.AllowCacheControlHeader.IsNull() {
		return false
	}
	if !data.BackPersistentTimeout.IsNull() {
		return false
	}
	if !data.FrontHttpVersion.IsNull() {
		return false
	}
	if !data.BackHttpVersion.IsNull() {
		return false
	}
	if !data.RequestSideSecurity.IsNull() {
		return false
	}
	if !data.ResponseSideSecurity.IsNull() {
		return false
	}
	if !data.DoChunkedUpload.IsNull() {
		return false
	}
	if !data.FollowRedirects.IsNull() {
		return false
	}
	if !data.HttpClientIpLabel.IsNull() {
		return false
	}
	if !data.HttpLogCorIdLabel.IsNull() {
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
	if !data.UrlRewritePolicy.IsNull() {
		return false
	}
	if !data.DoHostRewriting.IsNull() {
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
	return true
}

func (data WebAppFW) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.Priority.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Priority`, data.Priority.ValueString())
	}
	if !data.FrontSide.IsNull() {
		var dataValues []DmFrontSide
		data.FrontSide.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.SetRaw(body, pathRoot+`FrontSide`+".-1", val.ToBody(ctx, ""))
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`FrontSide`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`FrontSide`, "[]")
	}
	if !data.RemoteAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemoteAddress`, data.RemoteAddress.ValueString())
	}
	if !data.RemotePort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemotePort`, data.RemotePort.ValueInt64())
	}
	if !data.StylePolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StylePolicy`, data.StylePolicy.ValueString())
	}
	if !data.XmlManager.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XMLManager`, data.XmlManager.ValueString())
	}
	if !data.ErrorPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ErrorPolicy`, data.ErrorPolicy.ValueString())
	}
	if !data.UriNormalization.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`URINormalization`, tfutils.StringFromBool(data.UriNormalization, ""))
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
	if !data.StreamOutputToBack.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StreamOutputToBack`, data.StreamOutputToBack.ValueString())
	}
	if !data.StreamOutputToFront.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StreamOutputToFront`, data.StreamOutputToFront.ValueString())
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
	if !data.AllowCacheControlHeader.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllowCacheControlHeader`, tfutils.StringFromBool(data.AllowCacheControlHeader, ""))
	}
	if !data.BackPersistentTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BackPersistentTimeout`, data.BackPersistentTimeout.ValueInt64())
	}
	if !data.FrontHttpVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FrontHTTPVersion`, data.FrontHttpVersion.ValueString())
	}
	if !data.BackHttpVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BackHTTPVersion`, data.BackHttpVersion.ValueString())
	}
	if !data.RequestSideSecurity.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RequestSideSecurity`, tfutils.StringFromBool(data.RequestSideSecurity, ""))
	}
	if !data.ResponseSideSecurity.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ResponseSideSecurity`, tfutils.StringFromBool(data.ResponseSideSecurity, ""))
	}
	if !data.DoChunkedUpload.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DoChunkedUpload`, tfutils.StringFromBool(data.DoChunkedUpload, ""))
	}
	if !data.FollowRedirects.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FollowRedirects`, tfutils.StringFromBool(data.FollowRedirects, ""))
	}
	if !data.HttpClientIpLabel.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTPClientIPLabel`, data.HttpClientIpLabel.ValueString())
	}
	if !data.HttpLogCorIdLabel.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTPLogCorIDLabel`, data.HttpLogCorIdLabel.ValueString())
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
	if !data.UrlRewritePolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`URLRewritePolicy`, data.UrlRewritePolicy.ValueString())
	}
	if !data.DoHostRewriting.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DoHostRewriting`, tfutils.StringFromBool(data.DoHostRewriting, ""))
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
	return body
}

func (data *WebAppFW) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
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
	if value := res.Get(pathRoot + `FrontSide`); value.Exists() {
		l := []DmFrontSide{}
		if value := res.Get(`FrontSide`); value.Exists() {
			for _, v := range value.Array() {
				item := DmFrontSide{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.FrontSide, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmFrontSideObjectType}, l)
		} else {
			data.FrontSide = types.ListNull(types.ObjectType{AttrTypes: DmFrontSideObjectType})
		}
	} else {
		data.FrontSide = types.ListNull(types.ObjectType{AttrTypes: DmFrontSideObjectType})
	}
	if value := res.Get(pathRoot + `RemoteAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RemoteAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemotePort`); value.Exists() {
		data.RemotePort = types.Int64Value(value.Int())
	} else {
		data.RemotePort = types.Int64Value(80)
	}
	if value := res.Get(pathRoot + `StylePolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.StylePolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.StylePolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `XMLManager`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.XmlManager = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XmlManager = types.StringValue("default")
	}
	if value := res.Get(pathRoot + `ErrorPolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ErrorPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErrorPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `URINormalization`); value.Exists() {
		data.UriNormalization = tfutils.BoolFromString(value.String())
	} else {
		data.UriNormalization = types.BoolNull()
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
	if value := res.Get(pathRoot + `AllowCacheControlHeader`); value.Exists() {
		data.AllowCacheControlHeader = tfutils.BoolFromString(value.String())
	} else {
		data.AllowCacheControlHeader = types.BoolNull()
	}
	if value := res.Get(pathRoot + `BackPersistentTimeout`); value.Exists() {
		data.BackPersistentTimeout = types.Int64Value(value.Int())
	} else {
		data.BackPersistentTimeout = types.Int64Value(180)
	}
	if value := res.Get(pathRoot + `FrontHTTPVersion`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.FrontHttpVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FrontHttpVersion = types.StringValue("HTTP/1.1")
	}
	if value := res.Get(pathRoot + `BackHTTPVersion`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.BackHttpVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BackHttpVersion = types.StringValue("HTTP/1.1")
	}
	if value := res.Get(pathRoot + `RequestSideSecurity`); value.Exists() {
		data.RequestSideSecurity = tfutils.BoolFromString(value.String())
	} else {
		data.RequestSideSecurity = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ResponseSideSecurity`); value.Exists() {
		data.ResponseSideSecurity = tfutils.BoolFromString(value.String())
	} else {
		data.ResponseSideSecurity = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DoChunkedUpload`); value.Exists() {
		data.DoChunkedUpload = tfutils.BoolFromString(value.String())
	} else {
		data.DoChunkedUpload = types.BoolNull()
	}
	if value := res.Get(pathRoot + `FollowRedirects`); value.Exists() {
		data.FollowRedirects = tfutils.BoolFromString(value.String())
	} else {
		data.FollowRedirects = types.BoolNull()
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
	if value := res.Get(pathRoot + `URLRewritePolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UrlRewritePolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UrlRewritePolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `DoHostRewriting`); value.Exists() {
		data.DoHostRewriting = tfutils.BoolFromString(value.String())
	} else {
		data.DoHostRewriting = types.BoolNull()
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
}

func (data *WebAppFW) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
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
	if value := res.Get(pathRoot + `FrontSide`); value.Exists() && !data.FrontSide.IsNull() {
		l := []DmFrontSide{}
		e := []DmFrontSide{}
		data.FrontSide.ElementsAs(ctx, &e, false)
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
				item := DmFrontSide{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.FrontSide, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmFrontSideObjectType}, l)
		} else {
			data.FrontSide = types.ListNull(types.ObjectType{AttrTypes: DmFrontSideObjectType})
		}
	} else {
		data.FrontSide = types.ListNull(types.ObjectType{AttrTypes: DmFrontSideObjectType})
	}
	if value := res.Get(pathRoot + `RemoteAddress`); value.Exists() && !data.RemoteAddress.IsNull() {
		data.RemoteAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemotePort`); value.Exists() && !data.RemotePort.IsNull() {
		data.RemotePort = types.Int64Value(value.Int())
	} else if data.RemotePort.ValueInt64() != 80 {
		data.RemotePort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `StylePolicy`); value.Exists() && !data.StylePolicy.IsNull() {
		data.StylePolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.StylePolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `XMLManager`); value.Exists() && !data.XmlManager.IsNull() {
		data.XmlManager = tfutils.ParseStringFromGJSON(value)
	} else if data.XmlManager.ValueString() != "default" {
		data.XmlManager = types.StringNull()
	}
	if value := res.Get(pathRoot + `ErrorPolicy`); value.Exists() && !data.ErrorPolicy.IsNull() {
		data.ErrorPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErrorPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `URINormalization`); value.Exists() && !data.UriNormalization.IsNull() {
		data.UriNormalization = tfutils.BoolFromString(value.String())
	} else if !data.UriNormalization.ValueBool() {
		data.UriNormalization = types.BoolNull()
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
	if value := res.Get(pathRoot + `AllowCacheControlHeader`); value.Exists() && !data.AllowCacheControlHeader.IsNull() {
		data.AllowCacheControlHeader = tfutils.BoolFromString(value.String())
	} else if data.AllowCacheControlHeader.ValueBool() {
		data.AllowCacheControlHeader = types.BoolNull()
	}
	if value := res.Get(pathRoot + `BackPersistentTimeout`); value.Exists() && !data.BackPersistentTimeout.IsNull() {
		data.BackPersistentTimeout = types.Int64Value(value.Int())
	} else if data.BackPersistentTimeout.ValueInt64() != 180 {
		data.BackPersistentTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `FrontHTTPVersion`); value.Exists() && !data.FrontHttpVersion.IsNull() {
		data.FrontHttpVersion = tfutils.ParseStringFromGJSON(value)
	} else if data.FrontHttpVersion.ValueString() != "HTTP/1.1" {
		data.FrontHttpVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `BackHTTPVersion`); value.Exists() && !data.BackHttpVersion.IsNull() {
		data.BackHttpVersion = tfutils.ParseStringFromGJSON(value)
	} else if data.BackHttpVersion.ValueString() != "HTTP/1.1" {
		data.BackHttpVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `RequestSideSecurity`); value.Exists() && !data.RequestSideSecurity.IsNull() {
		data.RequestSideSecurity = tfutils.BoolFromString(value.String())
	} else if !data.RequestSideSecurity.ValueBool() {
		data.RequestSideSecurity = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ResponseSideSecurity`); value.Exists() && !data.ResponseSideSecurity.IsNull() {
		data.ResponseSideSecurity = tfutils.BoolFromString(value.String())
	} else if !data.ResponseSideSecurity.ValueBool() {
		data.ResponseSideSecurity = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DoChunkedUpload`); value.Exists() && !data.DoChunkedUpload.IsNull() {
		data.DoChunkedUpload = tfutils.BoolFromString(value.String())
	} else if data.DoChunkedUpload.ValueBool() {
		data.DoChunkedUpload = types.BoolNull()
	}
	if value := res.Get(pathRoot + `FollowRedirects`); value.Exists() && !data.FollowRedirects.IsNull() {
		data.FollowRedirects = tfutils.BoolFromString(value.String())
	} else if !data.FollowRedirects.ValueBool() {
		data.FollowRedirects = types.BoolNull()
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
	if value := res.Get(pathRoot + `URLRewritePolicy`); value.Exists() && !data.UrlRewritePolicy.IsNull() {
		data.UrlRewritePolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UrlRewritePolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `DoHostRewriting`); value.Exists() && !data.DoHostRewriting.IsNull() {
		data.DoHostRewriting = tfutils.BoolFromString(value.String())
	} else if !data.DoHostRewriting.ValueBool() {
		data.DoHostRewriting = types.BoolNull()
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
}
