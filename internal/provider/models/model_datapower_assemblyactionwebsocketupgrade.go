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
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type AssemblyActionWebSocketUpgrade struct {
	Id                            types.String      `tfsdk:"id"`
	AppDomain                     types.String      `tfsdk:"app_domain"`
	Url                           types.String      `tfsdk:"url"`
	SslClient                     types.String      `tfsdk:"ssl_client"`
	Timeout                       types.Int64       `tfsdk:"timeout"`
	UserName                      types.String      `tfsdk:"user_name"`
	Password                      types.String      `tfsdk:"password"`
	FollowRedirects               types.Bool        `tfsdk:"follow_redirects"`
	DecodeRequestParams           types.Bool        `tfsdk:"decode_request_params"`
	EncodePlusChar                types.Bool        `tfsdk:"encode_plus_char"`
	InjectUserAgentHeader         types.Bool        `tfsdk:"inject_user_agent_header"`
	InjectProxyHeaders            types.Bool        `tfsdk:"inject_proxy_headers"`
	HeaderControlList             types.String      `tfsdk:"header_control_list"`
	ParameterControlList          types.String      `tfsdk:"parameter_control_list"`
	ApiRequestProcessingAssembly  types.String      `tfsdk:"api_request_processing_assembly"`
	ApiResponseProcessingAssembly types.String      `tfsdk:"api_response_processing_assembly"`
	UserSummary                   types.String      `tfsdk:"user_summary"`
	Title                         types.String      `tfsdk:"title"`
	CorrelationPath               types.String      `tfsdk:"correlation_path"`
	ActionDebug                   types.Bool        `tfsdk:"action_debug"`
	ObjectActions                 []*actions.Action `tfsdk:"object_actions"`
}

var AssemblyActionWebSocketUpgradeObjectType = map[string]attr.Type{
	"id":                               types.StringType,
	"app_domain":                       types.StringType,
	"url":                              types.StringType,
	"ssl_client":                       types.StringType,
	"timeout":                          types.Int64Type,
	"user_name":                        types.StringType,
	"password":                         types.StringType,
	"follow_redirects":                 types.BoolType,
	"decode_request_params":            types.BoolType,
	"encode_plus_char":                 types.BoolType,
	"inject_user_agent_header":         types.BoolType,
	"inject_proxy_headers":             types.BoolType,
	"header_control_list":              types.StringType,
	"parameter_control_list":           types.StringType,
	"api_request_processing_assembly":  types.StringType,
	"api_response_processing_assembly": types.StringType,
	"user_summary":                     types.StringType,
	"title":                            types.StringType,
	"correlation_path":                 types.StringType,
	"action_debug":                     types.BoolType,
	"object_actions":                   actions.ActionsListType,
}

func (data AssemblyActionWebSocketUpgrade) GetPath() string {
	rest_path := "/mgmt/config/{domain}/AssemblyActionWebSocketUpgrade"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data AssemblyActionWebSocketUpgrade) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Url.IsNull() {
		return false
	}
	if !data.SslClient.IsNull() {
		return false
	}
	if !data.Timeout.IsNull() {
		return false
	}
	if !data.UserName.IsNull() {
		return false
	}
	if !data.Password.IsNull() {
		return false
	}
	if !data.FollowRedirects.IsNull() {
		return false
	}
	if !data.DecodeRequestParams.IsNull() {
		return false
	}
	if !data.EncodePlusChar.IsNull() {
		return false
	}
	if !data.InjectUserAgentHeader.IsNull() {
		return false
	}
	if !data.InjectProxyHeaders.IsNull() {
		return false
	}
	if !data.HeaderControlList.IsNull() {
		return false
	}
	if !data.ParameterControlList.IsNull() {
		return false
	}
	if !data.ApiRequestProcessingAssembly.IsNull() {
		return false
	}
	if !data.ApiResponseProcessingAssembly.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Title.IsNull() {
		return false
	}
	if !data.CorrelationPath.IsNull() {
		return false
	}
	if !data.ActionDebug.IsNull() {
		return false
	}
	return true
}

func (data AssemblyActionWebSocketUpgrade) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.Url.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Url`, data.Url.ValueString())
	}
	if !data.SslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClient`, data.SslClient.ValueString())
	}
	if !data.Timeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Timeout`, data.Timeout.ValueInt64())
	}
	if !data.UserName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserName`, data.UserName.ValueString())
	}
	if !data.Password.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Password`, data.Password.ValueString())
	}
	if !data.FollowRedirects.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FollowRedirects`, tfutils.StringFromBool(data.FollowRedirects, ""))
	}
	if !data.DecodeRequestParams.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DecodeRequestParams`, tfutils.StringFromBool(data.DecodeRequestParams, ""))
	}
	if !data.EncodePlusChar.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EncodePlusChar`, tfutils.StringFromBool(data.EncodePlusChar, ""))
	}
	if !data.InjectUserAgentHeader.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`InjectUserAgentHeader`, tfutils.StringFromBool(data.InjectUserAgentHeader, ""))
	}
	if !data.InjectProxyHeaders.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`InjectProxyHeaders`, tfutils.StringFromBool(data.InjectProxyHeaders, ""))
	}
	if !data.HeaderControlList.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HeaderControlList`, data.HeaderControlList.ValueString())
	}
	if !data.ParameterControlList.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ParameterControlList`, data.ParameterControlList.ValueString())
	}
	if !data.ApiRequestProcessingAssembly.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APIRequestProcessingAssembly`, data.ApiRequestProcessingAssembly.ValueString())
	}
	if !data.ApiResponseProcessingAssembly.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APIResponseProcessingAssembly`, data.ApiResponseProcessingAssembly.ValueString())
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.Title.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Title`, data.Title.ValueString())
	}
	if !data.CorrelationPath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CorrelationPath`, data.CorrelationPath.ValueString())
	}
	if !data.ActionDebug.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ActionDebug`, tfutils.StringFromBool(data.ActionDebug, ""))
	}
	return body
}

func (data *AssemblyActionWebSocketUpgrade) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Url`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Url = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Url = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLClient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `Timeout`); value.Exists() {
		data.Timeout = types.Int64Value(value.Int())
	} else {
		data.Timeout = types.Int64Value(60)
	}
	if value := res.Get(pathRoot + `UserName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserName = types.StringNull()
	}
	if value := res.Get(pathRoot + `Password`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Password = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Password = types.StringNull()
	}
	if value := res.Get(pathRoot + `FollowRedirects`); value.Exists() {
		data.FollowRedirects = tfutils.BoolFromString(value.String())
	} else {
		data.FollowRedirects = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DecodeRequestParams`); value.Exists() {
		data.DecodeRequestParams = tfutils.BoolFromString(value.String())
	} else {
		data.DecodeRequestParams = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EncodePlusChar`); value.Exists() {
		data.EncodePlusChar = tfutils.BoolFromString(value.String())
	} else {
		data.EncodePlusChar = types.BoolNull()
	}
	if value := res.Get(pathRoot + `InjectUserAgentHeader`); value.Exists() {
		data.InjectUserAgentHeader = tfutils.BoolFromString(value.String())
	} else {
		data.InjectUserAgentHeader = types.BoolNull()
	}
	if value := res.Get(pathRoot + `InjectProxyHeaders`); value.Exists() {
		data.InjectProxyHeaders = tfutils.BoolFromString(value.String())
	} else {
		data.InjectProxyHeaders = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HeaderControlList`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HeaderControlList = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HeaderControlList = types.StringValue("default-accept-all")
	}
	if value := res.Get(pathRoot + `ParameterControlList`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ParameterControlList = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ParameterControlList = types.StringValue("default-reject-all")
	}
	if value := res.Get(pathRoot + `APIRequestProcessingAssembly`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ApiRequestProcessingAssembly = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ApiRequestProcessingAssembly = types.StringNull()
	}
	if value := res.Get(pathRoot + `APIResponseProcessingAssembly`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ApiResponseProcessingAssembly = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ApiResponseProcessingAssembly = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `Title`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Title = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Title = types.StringNull()
	}
	if value := res.Get(pathRoot + `CorrelationPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CorrelationPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CorrelationPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `ActionDebug`); value.Exists() {
		data.ActionDebug = tfutils.BoolFromString(value.String())
	} else {
		data.ActionDebug = types.BoolNull()
	}
}

func (data *AssemblyActionWebSocketUpgrade) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Url`); value.Exists() && !data.Url.IsNull() {
		data.Url = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Url = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLClient`); value.Exists() && !data.SslClient.IsNull() {
		data.SslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `Timeout`); value.Exists() && !data.Timeout.IsNull() {
		data.Timeout = types.Int64Value(value.Int())
	} else if data.Timeout.ValueInt64() != 60 {
		data.Timeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `UserName`); value.Exists() && !data.UserName.IsNull() {
		data.UserName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserName = types.StringNull()
	}
	if value := res.Get(pathRoot + `Password`); value.Exists() && !data.Password.IsNull() {
		data.Password = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Password = types.StringNull()
	}
	if value := res.Get(pathRoot + `FollowRedirects`); value.Exists() && !data.FollowRedirects.IsNull() {
		data.FollowRedirects = tfutils.BoolFromString(value.String())
	} else if data.FollowRedirects.ValueBool() {
		data.FollowRedirects = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DecodeRequestParams`); value.Exists() && !data.DecodeRequestParams.IsNull() {
		data.DecodeRequestParams = tfutils.BoolFromString(value.String())
	} else if data.DecodeRequestParams.ValueBool() {
		data.DecodeRequestParams = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EncodePlusChar`); value.Exists() && !data.EncodePlusChar.IsNull() {
		data.EncodePlusChar = tfutils.BoolFromString(value.String())
	} else if data.EncodePlusChar.ValueBool() {
		data.EncodePlusChar = types.BoolNull()
	}
	if value := res.Get(pathRoot + `InjectUserAgentHeader`); value.Exists() && !data.InjectUserAgentHeader.IsNull() {
		data.InjectUserAgentHeader = tfutils.BoolFromString(value.String())
	} else if !data.InjectUserAgentHeader.ValueBool() {
		data.InjectUserAgentHeader = types.BoolNull()
	}
	if value := res.Get(pathRoot + `InjectProxyHeaders`); value.Exists() && !data.InjectProxyHeaders.IsNull() {
		data.InjectProxyHeaders = tfutils.BoolFromString(value.String())
	} else if data.InjectProxyHeaders.ValueBool() {
		data.InjectProxyHeaders = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HeaderControlList`); value.Exists() && !data.HeaderControlList.IsNull() {
		data.HeaderControlList = tfutils.ParseStringFromGJSON(value)
	} else if data.HeaderControlList.ValueString() != "default-accept-all" {
		data.HeaderControlList = types.StringNull()
	}
	if value := res.Get(pathRoot + `ParameterControlList`); value.Exists() && !data.ParameterControlList.IsNull() {
		data.ParameterControlList = tfutils.ParseStringFromGJSON(value)
	} else if data.ParameterControlList.ValueString() != "default-reject-all" {
		data.ParameterControlList = types.StringNull()
	}
	if value := res.Get(pathRoot + `APIRequestProcessingAssembly`); value.Exists() && !data.ApiRequestProcessingAssembly.IsNull() {
		data.ApiRequestProcessingAssembly = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ApiRequestProcessingAssembly = types.StringNull()
	}
	if value := res.Get(pathRoot + `APIResponseProcessingAssembly`); value.Exists() && !data.ApiResponseProcessingAssembly.IsNull() {
		data.ApiResponseProcessingAssembly = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ApiResponseProcessingAssembly = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `Title`); value.Exists() && !data.Title.IsNull() {
		data.Title = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Title = types.StringNull()
	}
	if value := res.Get(pathRoot + `CorrelationPath`); value.Exists() && !data.CorrelationPath.IsNull() {
		data.CorrelationPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CorrelationPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `ActionDebug`); value.Exists() && !data.ActionDebug.IsNull() {
		data.ActionDebug = tfutils.BoolFromString(value.String())
	} else if data.ActionDebug.ValueBool() {
		data.ActionDebug = types.BoolNull()
	}
}
