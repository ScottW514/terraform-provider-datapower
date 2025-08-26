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

type EBMS2SourceProtocolHandler struct {
	Id                    types.String                `tfsdk:"id"`
	AppDomain             types.String                `tfsdk:"app_domain"`
	UserSummary           types.String                `tfsdk:"user_summary"`
	LocalAddress          types.String                `tfsdk:"local_address"`
	LocalPort             types.Int64                 `tfsdk:"local_port"`
	HttpVersion           types.String                `tfsdk:"http_version"`
	PersistentConnections types.Bool                  `tfsdk:"persistent_connections"`
	AllowCompression      types.Bool                  `tfsdk:"allow_compression"`
	MaxUrlLen             types.Int64                 `tfsdk:"max_url_len"`
	MaxTotalHdrLen        types.Int64                 `tfsdk:"max_total_hdr_len"`
	MaxHdrCount           types.Int64                 `tfsdk:"max_hdr_count"`
	MaxNameHdrLen         types.Int64                 `tfsdk:"max_name_hdr_len"`
	MaxValueHdrLen        types.Int64                 `tfsdk:"max_value_hdr_len"`
	MaxQueryStringLen     types.Int64                 `tfsdk:"max_query_string_len"`
	Acl                   types.String                `tfsdk:"acl"`
	AaaPolicy             types.String                `tfsdk:"aaa_policy"`
	CredentialCharset     types.String                `tfsdk:"credential_charset"`
	SslServerConfigType   types.String                `tfsdk:"ssl_server_config_type"`
	SslServer             types.String                `tfsdk:"ssl_server"`
	SslSniServer          types.String                `tfsdk:"ssl_sni_server"`
	DependencyActions     []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var EBMS2SourceProtocolHandlerObjectType = map[string]attr.Type{
	"id":                     types.StringType,
	"app_domain":             types.StringType,
	"user_summary":           types.StringType,
	"local_address":          types.StringType,
	"local_port":             types.Int64Type,
	"http_version":           types.StringType,
	"persistent_connections": types.BoolType,
	"allow_compression":      types.BoolType,
	"max_url_len":            types.Int64Type,
	"max_total_hdr_len":      types.Int64Type,
	"max_hdr_count":          types.Int64Type,
	"max_name_hdr_len":       types.Int64Type,
	"max_value_hdr_len":      types.Int64Type,
	"max_query_string_len":   types.Int64Type,
	"acl":                    types.StringType,
	"aaa_policy":             types.StringType,
	"credential_charset":     types.StringType,
	"ssl_server_config_type": types.StringType,
	"ssl_server":             types.StringType,
	"ssl_sni_server":         types.StringType,
	"dependency_actions":     actions.ActionsListType,
}

func (data EBMS2SourceProtocolHandler) GetPath() string {
	rest_path := "/mgmt/config/{domain}/EBMS2SourceProtocolHandler"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data EBMS2SourceProtocolHandler) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.LocalAddress.IsNull() {
		return false
	}
	if !data.LocalPort.IsNull() {
		return false
	}
	if !data.HttpVersion.IsNull() {
		return false
	}
	if !data.PersistentConnections.IsNull() {
		return false
	}
	if !data.AllowCompression.IsNull() {
		return false
	}
	if !data.MaxUrlLen.IsNull() {
		return false
	}
	if !data.MaxTotalHdrLen.IsNull() {
		return false
	}
	if !data.MaxHdrCount.IsNull() {
		return false
	}
	if !data.MaxNameHdrLen.IsNull() {
		return false
	}
	if !data.MaxValueHdrLen.IsNull() {
		return false
	}
	if !data.MaxQueryStringLen.IsNull() {
		return false
	}
	if !data.Acl.IsNull() {
		return false
	}
	if !data.AaaPolicy.IsNull() {
		return false
	}
	if !data.CredentialCharset.IsNull() {
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
	return true
}

func (data EBMS2SourceProtocolHandler) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.LocalAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalAddress`, data.LocalAddress.ValueString())
	}
	if !data.LocalPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalPort`, data.LocalPort.ValueInt64())
	}
	if !data.HttpVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTPVersion`, data.HttpVersion.ValueString())
	}
	if !data.PersistentConnections.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PersistentConnections`, tfutils.StringFromBool(data.PersistentConnections, ""))
	}
	if !data.AllowCompression.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllowCompression`, tfutils.StringFromBool(data.AllowCompression, ""))
	}
	if !data.MaxUrlLen.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxURLLen`, data.MaxUrlLen.ValueInt64())
	}
	if !data.MaxTotalHdrLen.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxTotalHdrLen`, data.MaxTotalHdrLen.ValueInt64())
	}
	if !data.MaxHdrCount.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxHdrCount`, data.MaxHdrCount.ValueInt64())
	}
	if !data.MaxNameHdrLen.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxNameHdrLen`, data.MaxNameHdrLen.ValueInt64())
	}
	if !data.MaxValueHdrLen.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxValueHdrLen`, data.MaxValueHdrLen.ValueInt64())
	}
	if !data.MaxQueryStringLen.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxQueryStringLen`, data.MaxQueryStringLen.ValueInt64())
	}
	if !data.Acl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ACL`, data.Acl.ValueString())
	}
	if !data.AaaPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AAAPolicy`, data.AaaPolicy.ValueString())
	}
	if !data.CredentialCharset.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CredentialCharset`, data.CredentialCharset.ValueString())
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
	return body
}

func (data *EBMS2SourceProtocolHandler) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalAddress = types.StringValue("0.0.0.0")
	}
	if value := res.Get(pathRoot + `LocalPort`); value.Exists() {
		data.LocalPort = types.Int64Value(value.Int())
	} else {
		data.LocalPort = types.Int64Value(80)
	}
	if value := res.Get(pathRoot + `HTTPVersion`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HttpVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HttpVersion = types.StringValue("HTTP/1.1")
	}
	if value := res.Get(pathRoot + `PersistentConnections`); value.Exists() {
		data.PersistentConnections = tfutils.BoolFromString(value.String())
	} else {
		data.PersistentConnections = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowCompression`); value.Exists() {
		data.AllowCompression = tfutils.BoolFromString(value.String())
	} else {
		data.AllowCompression = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MaxURLLen`); value.Exists() {
		data.MaxUrlLen = types.Int64Value(value.Int())
	} else {
		data.MaxUrlLen = types.Int64Value(16384)
	}
	if value := res.Get(pathRoot + `MaxTotalHdrLen`); value.Exists() {
		data.MaxTotalHdrLen = types.Int64Value(value.Int())
	} else {
		data.MaxTotalHdrLen = types.Int64Value(128000)
	}
	if value := res.Get(pathRoot + `MaxHdrCount`); value.Exists() {
		data.MaxHdrCount = types.Int64Value(value.Int())
	} else {
		data.MaxHdrCount = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `MaxNameHdrLen`); value.Exists() {
		data.MaxNameHdrLen = types.Int64Value(value.Int())
	} else {
		data.MaxNameHdrLen = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `MaxValueHdrLen`); value.Exists() {
		data.MaxValueHdrLen = types.Int64Value(value.Int())
	} else {
		data.MaxValueHdrLen = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `MaxQueryStringLen`); value.Exists() {
		data.MaxQueryStringLen = types.Int64Value(value.Int())
	} else {
		data.MaxQueryStringLen = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `ACL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Acl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Acl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AAAPolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AaaPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AaaPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `CredentialCharset`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CredentialCharset = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CredentialCharset = types.StringValue("protocol")
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
}

func (data *EBMS2SourceProtocolHandler) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && !data.LocalAddress.IsNull() {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else if data.LocalAddress.ValueString() != "0.0.0.0" {
		data.LocalAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalPort`); value.Exists() && !data.LocalPort.IsNull() {
		data.LocalPort = types.Int64Value(value.Int())
	} else if data.LocalPort.ValueInt64() != 80 {
		data.LocalPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `HTTPVersion`); value.Exists() && !data.HttpVersion.IsNull() {
		data.HttpVersion = tfutils.ParseStringFromGJSON(value)
	} else if data.HttpVersion.ValueString() != "HTTP/1.1" {
		data.HttpVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `PersistentConnections`); value.Exists() && !data.PersistentConnections.IsNull() {
		data.PersistentConnections = tfutils.BoolFromString(value.String())
	} else if !data.PersistentConnections.ValueBool() {
		data.PersistentConnections = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowCompression`); value.Exists() && !data.AllowCompression.IsNull() {
		data.AllowCompression = tfutils.BoolFromString(value.String())
	} else if data.AllowCompression.ValueBool() {
		data.AllowCompression = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MaxURLLen`); value.Exists() && !data.MaxUrlLen.IsNull() {
		data.MaxUrlLen = types.Int64Value(value.Int())
	} else if data.MaxUrlLen.ValueInt64() != 16384 {
		data.MaxUrlLen = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxTotalHdrLen`); value.Exists() && !data.MaxTotalHdrLen.IsNull() {
		data.MaxTotalHdrLen = types.Int64Value(value.Int())
	} else if data.MaxTotalHdrLen.ValueInt64() != 128000 {
		data.MaxTotalHdrLen = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxHdrCount`); value.Exists() && !data.MaxHdrCount.IsNull() {
		data.MaxHdrCount = types.Int64Value(value.Int())
	} else if data.MaxHdrCount.ValueInt64() != 0 {
		data.MaxHdrCount = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxNameHdrLen`); value.Exists() && !data.MaxNameHdrLen.IsNull() {
		data.MaxNameHdrLen = types.Int64Value(value.Int())
	} else if data.MaxNameHdrLen.ValueInt64() != 0 {
		data.MaxNameHdrLen = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxValueHdrLen`); value.Exists() && !data.MaxValueHdrLen.IsNull() {
		data.MaxValueHdrLen = types.Int64Value(value.Int())
	} else if data.MaxValueHdrLen.ValueInt64() != 0 {
		data.MaxValueHdrLen = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxQueryStringLen`); value.Exists() && !data.MaxQueryStringLen.IsNull() {
		data.MaxQueryStringLen = types.Int64Value(value.Int())
	} else if data.MaxQueryStringLen.ValueInt64() != 0 {
		data.MaxQueryStringLen = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ACL`); value.Exists() && !data.Acl.IsNull() {
		data.Acl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Acl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AAAPolicy`); value.Exists() && !data.AaaPolicy.IsNull() {
		data.AaaPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AaaPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `CredentialCharset`); value.Exists() && !data.CredentialCharset.IsNull() {
		data.CredentialCharset = tfutils.ParseStringFromGJSON(value)
	} else if data.CredentialCharset.ValueString() != "protocol" {
		data.CredentialCharset = types.StringNull()
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
}
