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

type XSLCoprocService struct {
	Id                        types.String                `tfsdk:"id"`
	AppDomain                 types.String                `tfsdk:"app_domain"`
	UserSummary               types.String                `tfsdk:"user_summary"`
	Priority                  types.String                `tfsdk:"priority"`
	LocalPort                 types.Int64                 `tfsdk:"local_port"`
	XmlManager                types.String                `tfsdk:"xml_manager"`
	UrlRewritePolicy          types.String                `tfsdk:"url_rewrite_policy"`
	StylePolicyRule           types.String                `tfsdk:"style_policy_rule"`
	ConnectionTimeout         types.Int64                 `tfsdk:"connection_timeout"`
	IntermediateResultTimeout types.Int64                 `tfsdk:"intermediate_result_timeout"`
	CacheRelativeUrl          types.Bool                  `tfsdk:"cache_relative_url"`
	UseClientUriResolver      types.Bool                  `tfsdk:"use_client_uri_resolver"`
	CryptoExtensions          types.Bool                  `tfsdk:"crypto_extensions"`
	DefaultParamNamespace     types.String                `tfsdk:"default_param_namespace"`
	DebugMode                 types.String                `tfsdk:"debug_mode"`
	DebugHistory              types.Int64                 `tfsdk:"debug_history"`
	DebugTrigger              types.List                  `tfsdk:"debug_trigger"`
	SslServerConfigType       types.String                `tfsdk:"ssl_server_config_type"`
	SslServer                 types.String                `tfsdk:"ssl_server"`
	SslsniServer              types.String                `tfsdk:"sslsni_server"`
	LocalAddress              types.String                `tfsdk:"local_address"`
	DependencyActions         []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var XSLCoprocServiceDebugHistoryCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "debug_mode",
	AttrType:    "String",
	AttrDefault: "off",
	Value:       []string{"true"},
}

var XSLCoprocServiceObjectType = map[string]attr.Type{
	"id":                          types.StringType,
	"app_domain":                  types.StringType,
	"user_summary":                types.StringType,
	"priority":                    types.StringType,
	"local_port":                  types.Int64Type,
	"xml_manager":                 types.StringType,
	"url_rewrite_policy":          types.StringType,
	"style_policy_rule":           types.StringType,
	"connection_timeout":          types.Int64Type,
	"intermediate_result_timeout": types.Int64Type,
	"cache_relative_url":          types.BoolType,
	"use_client_uri_resolver":     types.BoolType,
	"crypto_extensions":           types.BoolType,
	"default_param_namespace":     types.StringType,
	"debug_mode":                  types.StringType,
	"debug_history":               types.Int64Type,
	"debug_trigger":               types.ListType{ElemType: types.ObjectType{AttrTypes: DmMSDebugTriggerTypeObjectType}},
	"ssl_server_config_type":      types.StringType,
	"ssl_server":                  types.StringType,
	"sslsni_server":               types.StringType,
	"local_address":               types.StringType,
	"dependency_actions":          actions.ActionsListType,
}

func (data XSLCoprocService) GetPath() string {
	rest_path := "/mgmt/config/{domain}/XSLCoprocService"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data XSLCoprocService) IsNull() bool {
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
	if !data.LocalPort.IsNull() {
		return false
	}
	if !data.XmlManager.IsNull() {
		return false
	}
	if !data.UrlRewritePolicy.IsNull() {
		return false
	}
	if !data.StylePolicyRule.IsNull() {
		return false
	}
	if !data.ConnectionTimeout.IsNull() {
		return false
	}
	if !data.IntermediateResultTimeout.IsNull() {
		return false
	}
	if !data.CacheRelativeUrl.IsNull() {
		return false
	}
	if !data.UseClientUriResolver.IsNull() {
		return false
	}
	if !data.CryptoExtensions.IsNull() {
		return false
	}
	if !data.DefaultParamNamespace.IsNull() {
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
	if !data.SslServerConfigType.IsNull() {
		return false
	}
	if !data.SslServer.IsNull() {
		return false
	}
	if !data.SslsniServer.IsNull() {
		return false
	}
	if !data.LocalAddress.IsNull() {
		return false
	}
	return true
}

func (data XSLCoprocService) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.LocalPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalPort`, data.LocalPort.ValueInt64())
	}
	if !data.XmlManager.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XMLManager`, data.XmlManager.ValueString())
	}
	if !data.UrlRewritePolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`URLRewritePolicy`, data.UrlRewritePolicy.ValueString())
	}
	if !data.StylePolicyRule.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StylePolicyRule`, data.StylePolicyRule.ValueString())
	}
	if !data.ConnectionTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ConnectionTimeout`, data.ConnectionTimeout.ValueInt64())
	}
	if !data.IntermediateResultTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IntermediateResultTimeout`, data.IntermediateResultTimeout.ValueInt64())
	}
	if !data.CacheRelativeUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CacheRelativeURL`, tfutils.StringFromBool(data.CacheRelativeUrl, ""))
	}
	if !data.UseClientUriResolver.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseClientURIResolver`, tfutils.StringFromBool(data.UseClientUriResolver, ""))
	}
	if !data.CryptoExtensions.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CryptoExtensions`, tfutils.StringFromBool(data.CryptoExtensions, ""))
	}
	if !data.DefaultParamNamespace.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DefaultParamNamespace`, data.DefaultParamNamespace.ValueString())
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
	if !data.SslServerConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLServerConfigType`, data.SslServerConfigType.ValueString())
	}
	if !data.SslServer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLServer`, data.SslServer.ValueString())
	}
	if !data.SslsniServer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLSNIServer`, data.SslsniServer.ValueString())
	}
	if !data.LocalAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalAddress`, data.LocalAddress.ValueString())
	}
	return body
}

func (data *XSLCoprocService) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `LocalPort`); value.Exists() {
		data.LocalPort = types.Int64Value(value.Int())
	} else {
		data.LocalPort = types.Int64Null()
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
	if value := res.Get(pathRoot + `StylePolicyRule`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.StylePolicyRule = tfutils.ParseStringFromGJSON(value)
	} else {
		data.StylePolicyRule = types.StringNull()
	}
	if value := res.Get(pathRoot + `ConnectionTimeout`); value.Exists() {
		data.ConnectionTimeout = types.Int64Value(value.Int())
	} else {
		data.ConnectionTimeout = types.Int64Value(60)
	}
	if value := res.Get(pathRoot + `IntermediateResultTimeout`); value.Exists() {
		data.IntermediateResultTimeout = types.Int64Value(value.Int())
	} else {
		data.IntermediateResultTimeout = types.Int64Value(20)
	}
	if value := res.Get(pathRoot + `CacheRelativeURL`); value.Exists() {
		data.CacheRelativeUrl = tfutils.BoolFromString(value.String())
	} else {
		data.CacheRelativeUrl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UseClientURIResolver`); value.Exists() {
		data.UseClientUriResolver = tfutils.BoolFromString(value.String())
	} else {
		data.UseClientUriResolver = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CryptoExtensions`); value.Exists() {
		data.CryptoExtensions = tfutils.BoolFromString(value.String())
	} else {
		data.CryptoExtensions = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DefaultParamNamespace`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DefaultParamNamespace = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DefaultParamNamespace = types.StringValue("http://www.datapower.com/param/config")
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
		data.SslsniServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslsniServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalAddress = types.StringValue("0.0.0.0")
	}
}

func (data *XSLCoprocService) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `LocalPort`); value.Exists() && !data.LocalPort.IsNull() {
		data.LocalPort = types.Int64Value(value.Int())
	} else {
		data.LocalPort = types.Int64Null()
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
	if value := res.Get(pathRoot + `StylePolicyRule`); value.Exists() && !data.StylePolicyRule.IsNull() {
		data.StylePolicyRule = tfutils.ParseStringFromGJSON(value)
	} else {
		data.StylePolicyRule = types.StringNull()
	}
	if value := res.Get(pathRoot + `ConnectionTimeout`); value.Exists() && !data.ConnectionTimeout.IsNull() {
		data.ConnectionTimeout = types.Int64Value(value.Int())
	} else if data.ConnectionTimeout.ValueInt64() != 60 {
		data.ConnectionTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `IntermediateResultTimeout`); value.Exists() && !data.IntermediateResultTimeout.IsNull() {
		data.IntermediateResultTimeout = types.Int64Value(value.Int())
	} else if data.IntermediateResultTimeout.ValueInt64() != 20 {
		data.IntermediateResultTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `CacheRelativeURL`); value.Exists() && !data.CacheRelativeUrl.IsNull() {
		data.CacheRelativeUrl = tfutils.BoolFromString(value.String())
	} else if data.CacheRelativeUrl.ValueBool() {
		data.CacheRelativeUrl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UseClientURIResolver`); value.Exists() && !data.UseClientUriResolver.IsNull() {
		data.UseClientUriResolver = tfutils.BoolFromString(value.String())
	} else if !data.UseClientUriResolver.ValueBool() {
		data.UseClientUriResolver = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CryptoExtensions`); value.Exists() && !data.CryptoExtensions.IsNull() {
		data.CryptoExtensions = tfutils.BoolFromString(value.String())
	} else if data.CryptoExtensions.ValueBool() {
		data.CryptoExtensions = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DefaultParamNamespace`); value.Exists() && !data.DefaultParamNamespace.IsNull() {
		data.DefaultParamNamespace = tfutils.ParseStringFromGJSON(value)
	} else if data.DefaultParamNamespace.ValueString() != "http://www.datapower.com/param/config" {
		data.DefaultParamNamespace = types.StringNull()
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
	if value := res.Get(pathRoot + `SSLSNIServer`); value.Exists() && !data.SslsniServer.IsNull() {
		data.SslsniServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslsniServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && !data.LocalAddress.IsNull() {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else if data.LocalAddress.ValueString() != "0.0.0.0" {
		data.LocalAddress = types.StringNull()
	}
}
