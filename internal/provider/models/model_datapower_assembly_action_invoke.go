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

type AssemblyActionInvoke struct {
	Id                    types.String                `tfsdk:"id"`
	AppDomain             types.String                `tfsdk:"app_domain"`
	Url                   types.String                `tfsdk:"url"`
	SslClient             types.String                `tfsdk:"ssl_client"`
	Timeout               types.Int64                 `tfsdk:"timeout"`
	UserName              types.String                `tfsdk:"user_name"`
	Password              types.String                `tfsdk:"password"`
	Method                types.String                `tfsdk:"method"`
	BackendType           types.String                `tfsdk:"backend_type"`
	GraphqlSendType       types.String                `tfsdk:"graphql_send_type"`
	Compression           types.Bool                  `tfsdk:"compression"`
	CacheType             types.String                `tfsdk:"cache_type"`
	TimeToLive            types.Int64                 `tfsdk:"time_to_live"`
	CacheUnsafeResponse   types.Bool                  `tfsdk:"cache_unsafe_response"`
	CacheKey              types.String                `tfsdk:"cache_key"`
	FollowRedirects       types.Bool                  `tfsdk:"follow_redirects"`
	HttpVersion           types.String                `tfsdk:"http_version"`
	Http2Required         types.Bool                  `tfsdk:"http2_required"`
	DoChunkedUpload       types.Bool                  `tfsdk:"do_chunked_upload"`
	PersistentConnection  types.Bool                  `tfsdk:"persistent_connection"`
	StopOnError           types.Bool                  `tfsdk:"stop_on_error"`
	ErrorTypes            *DmInvokeErrorType          `tfsdk:"error_types"`
	Output                types.String                `tfsdk:"output"`
	DecodeRequestParams   types.Bool                  `tfsdk:"decode_request_params"`
	EncodePlusChar        types.Bool                  `tfsdk:"encode_plus_char"`
	KeepPayload           types.Bool                  `tfsdk:"keep_payload"`
	InjectUserAgentHeader types.Bool                  `tfsdk:"inject_user_agent_header"`
	InjectProxyHeaders    types.Bool                  `tfsdk:"inject_proxy_headers"`
	HeaderControlList     types.String                `tfsdk:"header_control_list"`
	ParameterControlList  types.String                `tfsdk:"parameter_control_list"`
	UserSummary           types.String                `tfsdk:"user_summary"`
	Title                 types.String                `tfsdk:"title"`
	CorrelationPath       types.String                `tfsdk:"correlation_path"`
	ActionDebug           types.Bool                  `tfsdk:"action_debug"`
	DependencyActions     []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var AssemblyActionInvokeGraphQLSendTypeCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "method",
			AttrType:    "String",
			AttrDefault: "Keep",
			Value:       []string{"Keep", "POST"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "backend_type",
			AttrType:    "String",
			AttrDefault: "detect",
			Value:       []string{"detect", "graphql"},
		},
	},
}
var AssemblyActionInvokeTimeToLiveCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "cache_type",
	AttrType:    "String",
	AttrDefault: "Protocol",
	Value:       []string{"TimeToLive"},
}
var AssemblyActionInvokeErrorTypesCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "stop_on_error",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var AssemblyActionInvokeGraphQLSendTypeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "method",
			AttrType:    "String",
			AttrDefault: "Keep",
			Value:       []string{"Keep", "POST"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "backend_type",
			AttrType:    "String",
			AttrDefault: "detect",
			Value:       []string{"detect", "graphql"},
		},
	},
}
var AssemblyActionInvokeTimeToLiveIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "cache_type",
	AttrType:    "String",
	AttrDefault: "Protocol",
	Value:       []string{"TimeToLive"},
}
var AssemblyActionInvokeCacheUnsafeResponseIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "cache_type",
	AttrType:    "String",
	AttrDefault: "Protocol",
	Value:       []string{"TimeToLive"},
}
var AssemblyActionInvokeHTTP2RequiredIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "http_version",
	AttrType:    "String",
	AttrDefault: "HTTP/1.1",
	Value:       []string{"HTTP/2"},
}
var AssemblyActionInvokeErrorTypesIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "stop_on_error",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var AssemblyActionInvokeObjectType = map[string]attr.Type{
	"id":                       types.StringType,
	"app_domain":               types.StringType,
	"url":                      types.StringType,
	"ssl_client":               types.StringType,
	"timeout":                  types.Int64Type,
	"user_name":                types.StringType,
	"password":                 types.StringType,
	"method":                   types.StringType,
	"backend_type":             types.StringType,
	"graphql_send_type":        types.StringType,
	"compression":              types.BoolType,
	"cache_type":               types.StringType,
	"time_to_live":             types.Int64Type,
	"cache_unsafe_response":    types.BoolType,
	"cache_key":                types.StringType,
	"follow_redirects":         types.BoolType,
	"http_version":             types.StringType,
	"http2_required":           types.BoolType,
	"do_chunked_upload":        types.BoolType,
	"persistent_connection":    types.BoolType,
	"stop_on_error":            types.BoolType,
	"error_types":              types.ObjectType{AttrTypes: DmInvokeErrorTypeObjectType},
	"output":                   types.StringType,
	"decode_request_params":    types.BoolType,
	"encode_plus_char":         types.BoolType,
	"keep_payload":             types.BoolType,
	"inject_user_agent_header": types.BoolType,
	"inject_proxy_headers":     types.BoolType,
	"header_control_list":      types.StringType,
	"parameter_control_list":   types.StringType,
	"user_summary":             types.StringType,
	"title":                    types.StringType,
	"correlation_path":         types.StringType,
	"action_debug":             types.BoolType,
	"dependency_actions":       actions.ActionsListType,
}

func (data AssemblyActionInvoke) GetPath() string {
	rest_path := "/mgmt/config/{domain}/AssemblyActionInvoke"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data AssemblyActionInvoke) IsNull() bool {
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
	if !data.Method.IsNull() {
		return false
	}
	if !data.BackendType.IsNull() {
		return false
	}
	if !data.GraphqlSendType.IsNull() {
		return false
	}
	if !data.Compression.IsNull() {
		return false
	}
	if !data.CacheType.IsNull() {
		return false
	}
	if !data.TimeToLive.IsNull() {
		return false
	}
	if !data.CacheUnsafeResponse.IsNull() {
		return false
	}
	if !data.CacheKey.IsNull() {
		return false
	}
	if !data.FollowRedirects.IsNull() {
		return false
	}
	if !data.HttpVersion.IsNull() {
		return false
	}
	if !data.Http2Required.IsNull() {
		return false
	}
	if !data.DoChunkedUpload.IsNull() {
		return false
	}
	if !data.PersistentConnection.IsNull() {
		return false
	}
	if !data.StopOnError.IsNull() {
		return false
	}
	if data.ErrorTypes != nil {
		if !data.ErrorTypes.IsNull() {
			return false
		}
	}
	if !data.Output.IsNull() {
		return false
	}
	if !data.DecodeRequestParams.IsNull() {
		return false
	}
	if !data.EncodePlusChar.IsNull() {
		return false
	}
	if !data.KeepPayload.IsNull() {
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

func (data AssemblyActionInvoke) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Method.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Method`, data.Method.ValueString())
	}
	if !data.BackendType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BackendType`, data.BackendType.ValueString())
	}
	if !data.GraphqlSendType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GraphQLSendType`, data.GraphqlSendType.ValueString())
	}
	if !data.Compression.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Compression`, tfutils.StringFromBool(data.Compression, ""))
	}
	if !data.CacheType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CacheType`, data.CacheType.ValueString())
	}
	if !data.TimeToLive.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TimeToLive`, data.TimeToLive.ValueInt64())
	}
	if !data.CacheUnsafeResponse.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CacheUnsafeResponse`, tfutils.StringFromBool(data.CacheUnsafeResponse, ""))
	}
	if !data.CacheKey.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CacheKey`, data.CacheKey.ValueString())
	}
	if !data.FollowRedirects.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FollowRedirects`, tfutils.StringFromBool(data.FollowRedirects, ""))
	}
	if !data.HttpVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTPVersion`, data.HttpVersion.ValueString())
	}
	if !data.Http2Required.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP2Required`, tfutils.StringFromBool(data.Http2Required, ""))
	}
	if !data.DoChunkedUpload.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DoChunkedUpload`, tfutils.StringFromBool(data.DoChunkedUpload, ""))
	}
	if !data.PersistentConnection.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PersistentConnection`, tfutils.StringFromBool(data.PersistentConnection, ""))
	}
	if !data.StopOnError.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StopOnError`, tfutils.StringFromBool(data.StopOnError, ""))
	}
	if data.ErrorTypes != nil {
		if !data.ErrorTypes.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`ErrorTypes`, data.ErrorTypes.ToBody(ctx, ""))
		}
	}
	if !data.Output.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Output`, data.Output.ValueString())
	}
	if !data.DecodeRequestParams.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DecodeRequestParams`, tfutils.StringFromBool(data.DecodeRequestParams, ""))
	}
	if !data.EncodePlusChar.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EncodePlusChar`, tfutils.StringFromBool(data.EncodePlusChar, ""))
	}
	if !data.KeepPayload.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`KeepPayload`, tfutils.StringFromBool(data.KeepPayload, ""))
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

func (data *AssemblyActionInvoke) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Method`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Method = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Method = types.StringValue("Keep")
	}
	if value := res.Get(pathRoot + `BackendType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.BackendType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BackendType = types.StringValue("detect")
	}
	if value := res.Get(pathRoot + `GraphQLSendType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.GraphqlSendType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GraphqlSendType = types.StringValue("detect")
	}
	if value := res.Get(pathRoot + `Compression`); value.Exists() {
		data.Compression = tfutils.BoolFromString(value.String())
	} else {
		data.Compression = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CacheType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CacheType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CacheType = types.StringValue("Protocol")
	}
	if value := res.Get(pathRoot + `TimeToLive`); value.Exists() {
		data.TimeToLive = types.Int64Value(value.Int())
	} else {
		data.TimeToLive = types.Int64Value(900)
	}
	if value := res.Get(pathRoot + `CacheUnsafeResponse`); value.Exists() {
		data.CacheUnsafeResponse = tfutils.BoolFromString(value.String())
	} else {
		data.CacheUnsafeResponse = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CacheKey`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CacheKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CacheKey = types.StringNull()
	}
	if value := res.Get(pathRoot + `FollowRedirects`); value.Exists() {
		data.FollowRedirects = tfutils.BoolFromString(value.String())
	} else {
		data.FollowRedirects = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTPVersion`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HttpVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HttpVersion = types.StringValue("HTTP/1.1")
	}
	if value := res.Get(pathRoot + `HTTP2Required`); value.Exists() {
		data.Http2Required = tfutils.BoolFromString(value.String())
	} else {
		data.Http2Required = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DoChunkedUpload`); value.Exists() {
		data.DoChunkedUpload = tfutils.BoolFromString(value.String())
	} else {
		data.DoChunkedUpload = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PersistentConnection`); value.Exists() {
		data.PersistentConnection = tfutils.BoolFromString(value.String())
	} else {
		data.PersistentConnection = types.BoolNull()
	}
	if value := res.Get(pathRoot + `StopOnError`); value.Exists() {
		data.StopOnError = tfutils.BoolFromString(value.String())
	} else {
		data.StopOnError = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ErrorTypes`); value.Exists() {
		data.ErrorTypes = &DmInvokeErrorType{}
		data.ErrorTypes.FromBody(ctx, "", value)
	} else {
		data.ErrorTypes = nil
	}
	if value := res.Get(pathRoot + `Output`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Output = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Output = types.StringValue("message")
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
	if value := res.Get(pathRoot + `KeepPayload`); value.Exists() {
		data.KeepPayload = tfutils.BoolFromString(value.String())
	} else {
		data.KeepPayload = types.BoolNull()
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

func (data *AssemblyActionInvoke) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Method`); value.Exists() && !data.Method.IsNull() {
		data.Method = tfutils.ParseStringFromGJSON(value)
	} else if data.Method.ValueString() != "Keep" {
		data.Method = types.StringNull()
	}
	if value := res.Get(pathRoot + `BackendType`); value.Exists() && !data.BackendType.IsNull() {
		data.BackendType = tfutils.ParseStringFromGJSON(value)
	} else if data.BackendType.ValueString() != "detect" {
		data.BackendType = types.StringNull()
	}
	if value := res.Get(pathRoot + `GraphQLSendType`); value.Exists() && !data.GraphqlSendType.IsNull() {
		data.GraphqlSendType = tfutils.ParseStringFromGJSON(value)
	} else if data.GraphqlSendType.ValueString() != "detect" {
		data.GraphqlSendType = types.StringNull()
	}
	if value := res.Get(pathRoot + `Compression`); value.Exists() && !data.Compression.IsNull() {
		data.Compression = tfutils.BoolFromString(value.String())
	} else if data.Compression.ValueBool() {
		data.Compression = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CacheType`); value.Exists() && !data.CacheType.IsNull() {
		data.CacheType = tfutils.ParseStringFromGJSON(value)
	} else if data.CacheType.ValueString() != "Protocol" {
		data.CacheType = types.StringNull()
	}
	if value := res.Get(pathRoot + `TimeToLive`); value.Exists() && !data.TimeToLive.IsNull() {
		data.TimeToLive = types.Int64Value(value.Int())
	} else if data.TimeToLive.ValueInt64() != 900 {
		data.TimeToLive = types.Int64Null()
	}
	if value := res.Get(pathRoot + `CacheUnsafeResponse`); value.Exists() && !data.CacheUnsafeResponse.IsNull() {
		data.CacheUnsafeResponse = tfutils.BoolFromString(value.String())
	} else if data.CacheUnsafeResponse.ValueBool() {
		data.CacheUnsafeResponse = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CacheKey`); value.Exists() && !data.CacheKey.IsNull() {
		data.CacheKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CacheKey = types.StringNull()
	}
	if value := res.Get(pathRoot + `FollowRedirects`); value.Exists() && !data.FollowRedirects.IsNull() {
		data.FollowRedirects = tfutils.BoolFromString(value.String())
	} else if data.FollowRedirects.ValueBool() {
		data.FollowRedirects = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTPVersion`); value.Exists() && !data.HttpVersion.IsNull() {
		data.HttpVersion = tfutils.ParseStringFromGJSON(value)
	} else if data.HttpVersion.ValueString() != "HTTP/1.1" {
		data.HttpVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `HTTP2Required`); value.Exists() && !data.Http2Required.IsNull() {
		data.Http2Required = tfutils.BoolFromString(value.String())
	} else if data.Http2Required.ValueBool() {
		data.Http2Required = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DoChunkedUpload`); value.Exists() && !data.DoChunkedUpload.IsNull() {
		data.DoChunkedUpload = tfutils.BoolFromString(value.String())
	} else if !data.DoChunkedUpload.ValueBool() {
		data.DoChunkedUpload = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PersistentConnection`); value.Exists() && !data.PersistentConnection.IsNull() {
		data.PersistentConnection = tfutils.BoolFromString(value.String())
	} else if !data.PersistentConnection.ValueBool() {
		data.PersistentConnection = types.BoolNull()
	}
	if value := res.Get(pathRoot + `StopOnError`); value.Exists() && !data.StopOnError.IsNull() {
		data.StopOnError = tfutils.BoolFromString(value.String())
	} else if data.StopOnError.ValueBool() {
		data.StopOnError = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ErrorTypes`); value.Exists() {
		data.ErrorTypes.UpdateFromBody(ctx, "", value)
	} else {
		data.ErrorTypes = nil
	}
	if value := res.Get(pathRoot + `Output`); value.Exists() && !data.Output.IsNull() {
		data.Output = tfutils.ParseStringFromGJSON(value)
	} else if data.Output.ValueString() != "message" {
		data.Output = types.StringNull()
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
	if value := res.Get(pathRoot + `KeepPayload`); value.Exists() && !data.KeepPayload.IsNull() {
		data.KeepPayload = tfutils.BoolFromString(value.String())
	} else if data.KeepPayload.ValueBool() {
		data.KeepPayload = types.BoolNull()
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
