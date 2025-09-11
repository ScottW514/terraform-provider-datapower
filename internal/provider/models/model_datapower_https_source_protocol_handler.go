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

type HTTPSSourceProtocolHandler struct {
	Id                            types.String                `tfsdk:"id"`
	AppDomain                     types.String                `tfsdk:"app_domain"`
	UserSummary                   types.String                `tfsdk:"user_summary"`
	LocalAddress                  types.String                `tfsdk:"local_address"`
	LocalPort                     types.Int64                 `tfsdk:"local_port"`
	HttpVersion                   types.String                `tfsdk:"http_version"`
	AllowedFeatures               *DmSourceHTTPFeatureType    `tfsdk:"allowed_features"`
	PersistentConnections         types.Bool                  `tfsdk:"persistent_connections"`
	MaxPersistentConnectionsReuse types.Int64                 `tfsdk:"max_persistent_connections_reuse"`
	AllowCompression              types.Bool                  `tfsdk:"allow_compression"`
	AllowWebSocketUpgrade         types.Bool                  `tfsdk:"allow_web_socket_upgrade"`
	WebSocketIdleTimeout          types.Int64                 `tfsdk:"web_socket_idle_timeout"`
	MaxUrlLen                     types.Int64                 `tfsdk:"max_url_len"`
	MaxTotalHdrLen                types.Int64                 `tfsdk:"max_total_hdr_len"`
	MaxHdrCount                   types.Int64                 `tfsdk:"max_hdr_count"`
	MaxNameHdrLen                 types.Int64                 `tfsdk:"max_name_hdr_len"`
	MaxValueHdrLen                types.Int64                 `tfsdk:"max_value_hdr_len"`
	MaxQueryStringLen             types.Int64                 `tfsdk:"max_query_string_len"`
	Acl                           types.String                `tfsdk:"acl"`
	CredentialCharset             types.String                `tfsdk:"credential_charset"`
	SslServerConfigType           types.String                `tfsdk:"ssl_server_config_type"`
	SslServer                     types.String                `tfsdk:"ssl_server"`
	SslSniServer                  types.String                `tfsdk:"ssl_sni_server"`
	Http2MaxStreams               types.Int64                 `tfsdk:"http2_max_streams"`
	Http2MaxFrameSize             types.Int64                 `tfsdk:"http2_max_frame_size"`
	Http2StreamHeader             types.Bool                  `tfsdk:"http2_stream_header"`
	ChunkedEncoding               types.Bool                  `tfsdk:"chunked_encoding"`
	HeaderTimeout                 types.Int64                 `tfsdk:"header_timeout"`
	Http2IdleTimeout              types.Int64                 `tfsdk:"http2_idle_timeout"`
	DependencyActions             []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var HTTPSSourceProtocolHandlerWebSocketIdleTimeoutCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "allow_web_socket_upgrade",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "http_version",
			AttrType:    "String",
			AttrDefault: "HTTP/1.1",
			Value:       []string{"HTTP/1.0"},
		},
	},
}
var HTTPSSourceProtocolHandlerSSLServerCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "ssl_server_config_type",
	AttrType:    "String",
	AttrDefault: "server",
	Value:       []string{"server"},
}
var HTTPSSourceProtocolHandlerSSLSNIServerCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "ssl_server_config_type",
	AttrType:    "String",
	AttrDefault: "server",
	Value:       []string{"sni"},
}
var HTTPSSourceProtocolHandlerMaxPersistentConnectionsReuseIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "persistent_connections",
	AttrType:    "Bool",
	AttrDefault: "true",
	Value:       []string{"false"},
}
var HTTPSSourceProtocolHandlerAllowWebSocketUpgradeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "http_version",
	AttrType:    "String",
	AttrDefault: "HTTP/1.1",
	Value:       []string{"HTTP/1.0"},
}
var HTTPSSourceProtocolHandlerWebSocketIdleTimeoutIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var HTTPSSourceProtocolHandlerSSLServerIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var HTTPSSourceProtocolHandlerSSLSNIServerIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var HTTPSSourceProtocolHandlerObjectType = map[string]attr.Type{
	"id":                               types.StringType,
	"app_domain":                       types.StringType,
	"user_summary":                     types.StringType,
	"local_address":                    types.StringType,
	"local_port":                       types.Int64Type,
	"http_version":                     types.StringType,
	"allowed_features":                 types.ObjectType{AttrTypes: DmSourceHTTPFeatureTypeObjectType},
	"persistent_connections":           types.BoolType,
	"max_persistent_connections_reuse": types.Int64Type,
	"allow_compression":                types.BoolType,
	"allow_web_socket_upgrade":         types.BoolType,
	"web_socket_idle_timeout":          types.Int64Type,
	"max_url_len":                      types.Int64Type,
	"max_total_hdr_len":                types.Int64Type,
	"max_hdr_count":                    types.Int64Type,
	"max_name_hdr_len":                 types.Int64Type,
	"max_value_hdr_len":                types.Int64Type,
	"max_query_string_len":             types.Int64Type,
	"acl":                              types.StringType,
	"credential_charset":               types.StringType,
	"ssl_server_config_type":           types.StringType,
	"ssl_server":                       types.StringType,
	"ssl_sni_server":                   types.StringType,
	"http2_max_streams":                types.Int64Type,
	"http2_max_frame_size":             types.Int64Type,
	"http2_stream_header":              types.BoolType,
	"chunked_encoding":                 types.BoolType,
	"header_timeout":                   types.Int64Type,
	"http2_idle_timeout":               types.Int64Type,
	"dependency_actions":               actions.ActionsListType,
}

func (data HTTPSSourceProtocolHandler) GetPath() string {
	rest_path := "/mgmt/config/{domain}/HTTPSSourceProtocolHandler"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data HTTPSSourceProtocolHandler) IsNull() bool {
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
	if data.AllowedFeatures != nil {
		if !data.AllowedFeatures.IsNull() {
			return false
		}
	}
	if !data.PersistentConnections.IsNull() {
		return false
	}
	if !data.MaxPersistentConnectionsReuse.IsNull() {
		return false
	}
	if !data.AllowCompression.IsNull() {
		return false
	}
	if !data.AllowWebSocketUpgrade.IsNull() {
		return false
	}
	if !data.WebSocketIdleTimeout.IsNull() {
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
	if !data.Http2MaxStreams.IsNull() {
		return false
	}
	if !data.Http2MaxFrameSize.IsNull() {
		return false
	}
	if !data.Http2StreamHeader.IsNull() {
		return false
	}
	if !data.ChunkedEncoding.IsNull() {
		return false
	}
	if !data.HeaderTimeout.IsNull() {
		return false
	}
	if !data.Http2IdleTimeout.IsNull() {
		return false
	}
	return true
}

func (data HTTPSSourceProtocolHandler) ToBody(ctx context.Context, pathRoot string) string {
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
	if data.AllowedFeatures != nil {
		if !data.AllowedFeatures.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`AllowedFeatures`, data.AllowedFeatures.ToBody(ctx, ""))
		}
	}
	if !data.PersistentConnections.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PersistentConnections`, tfutils.StringFromBool(data.PersistentConnections, ""))
	}
	if !data.MaxPersistentConnectionsReuse.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxPersistentConnectionsReuse`, data.MaxPersistentConnectionsReuse.ValueInt64())
	}
	if !data.AllowCompression.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllowCompression`, tfutils.StringFromBool(data.AllowCompression, ""))
	}
	if !data.AllowWebSocketUpgrade.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllowWebSocketUpgrade`, tfutils.StringFromBool(data.AllowWebSocketUpgrade, ""))
	}
	if !data.WebSocketIdleTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WebSocketIdleTimeout`, data.WebSocketIdleTimeout.ValueInt64())
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
	if !data.Http2MaxStreams.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP2MaxStreams`, data.Http2MaxStreams.ValueInt64())
	}
	if !data.Http2MaxFrameSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP2MaxFrameSize`, data.Http2MaxFrameSize.ValueInt64())
	}
	if !data.Http2StreamHeader.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP2StreamHeader`, tfutils.StringFromBool(data.Http2StreamHeader, ""))
	}
	if !data.ChunkedEncoding.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ChunkedEncoding`, tfutils.StringFromBool(data.ChunkedEncoding, ""))
	}
	if !data.HeaderTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HeaderTimeout`, data.HeaderTimeout.ValueInt64())
	}
	if !data.Http2IdleTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP2IdleTimeout`, data.Http2IdleTimeout.ValueInt64())
	}
	return body
}

func (data *HTTPSSourceProtocolHandler) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
		data.LocalPort = types.Int64Value(443)
	}
	if value := res.Get(pathRoot + `HTTPVersion`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HttpVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HttpVersion = types.StringValue("HTTP/1.1")
	}
	if value := res.Get(pathRoot + `AllowedFeatures`); value.Exists() {
		data.AllowedFeatures = &DmSourceHTTPFeatureType{}
		data.AllowedFeatures.FromBody(ctx, "", value)
	} else {
		data.AllowedFeatures = nil
	}
	if value := res.Get(pathRoot + `PersistentConnections`); value.Exists() {
		data.PersistentConnections = tfutils.BoolFromString(value.String())
	} else {
		data.PersistentConnections = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MaxPersistentConnectionsReuse`); value.Exists() {
		data.MaxPersistentConnectionsReuse = types.Int64Value(value.Int())
	} else {
		data.MaxPersistentConnectionsReuse = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AllowCompression`); value.Exists() {
		data.AllowCompression = tfutils.BoolFromString(value.String())
	} else {
		data.AllowCompression = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowWebSocketUpgrade`); value.Exists() {
		data.AllowWebSocketUpgrade = tfutils.BoolFromString(value.String())
	} else {
		data.AllowWebSocketUpgrade = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WebSocketIdleTimeout`); value.Exists() {
		data.WebSocketIdleTimeout = types.Int64Value(value.Int())
	} else {
		data.WebSocketIdleTimeout = types.Int64Null()
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
		data.MaxHdrCount = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxNameHdrLen`); value.Exists() {
		data.MaxNameHdrLen = types.Int64Value(value.Int())
	} else {
		data.MaxNameHdrLen = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxValueHdrLen`); value.Exists() {
		data.MaxValueHdrLen = types.Int64Value(value.Int())
	} else {
		data.MaxValueHdrLen = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxQueryStringLen`); value.Exists() {
		data.MaxQueryStringLen = types.Int64Value(value.Int())
	} else {
		data.MaxQueryStringLen = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ACL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Acl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Acl = types.StringNull()
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
	if value := res.Get(pathRoot + `HTTP2MaxStreams`); value.Exists() {
		data.Http2MaxStreams = types.Int64Value(value.Int())
	} else {
		data.Http2MaxStreams = types.Int64Value(100)
	}
	if value := res.Get(pathRoot + `HTTP2MaxFrameSize`); value.Exists() {
		data.Http2MaxFrameSize = types.Int64Value(value.Int())
	} else {
		data.Http2MaxFrameSize = types.Int64Value(16384)
	}
	if value := res.Get(pathRoot + `HTTP2StreamHeader`); value.Exists() {
		data.Http2StreamHeader = tfutils.BoolFromString(value.String())
	} else {
		data.Http2StreamHeader = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ChunkedEncoding`); value.Exists() {
		data.ChunkedEncoding = tfutils.BoolFromString(value.String())
	} else {
		data.ChunkedEncoding = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HeaderTimeout`); value.Exists() {
		data.HeaderTimeout = types.Int64Value(value.Int())
	} else {
		data.HeaderTimeout = types.Int64Value(30000)
	}
	if value := res.Get(pathRoot + `HTTP2IdleTimeout`); value.Exists() {
		data.Http2IdleTimeout = types.Int64Value(value.Int())
	} else {
		data.Http2IdleTimeout = types.Int64Null()
	}
}

func (data *HTTPSSourceProtocolHandler) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	} else if data.LocalPort.ValueInt64() != 443 {
		data.LocalPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `HTTPVersion`); value.Exists() && !data.HttpVersion.IsNull() {
		data.HttpVersion = tfutils.ParseStringFromGJSON(value)
	} else if data.HttpVersion.ValueString() != "HTTP/1.1" {
		data.HttpVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `AllowedFeatures`); value.Exists() {
		data.AllowedFeatures.UpdateFromBody(ctx, "", value)
	} else {
		data.AllowedFeatures = nil
	}
	if value := res.Get(pathRoot + `PersistentConnections`); value.Exists() && !data.PersistentConnections.IsNull() {
		data.PersistentConnections = tfutils.BoolFromString(value.String())
	} else if !data.PersistentConnections.ValueBool() {
		data.PersistentConnections = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MaxPersistentConnectionsReuse`); value.Exists() && !data.MaxPersistentConnectionsReuse.IsNull() {
		data.MaxPersistentConnectionsReuse = types.Int64Value(value.Int())
	} else {
		data.MaxPersistentConnectionsReuse = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AllowCompression`); value.Exists() && !data.AllowCompression.IsNull() {
		data.AllowCompression = tfutils.BoolFromString(value.String())
	} else if data.AllowCompression.ValueBool() {
		data.AllowCompression = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowWebSocketUpgrade`); value.Exists() && !data.AllowWebSocketUpgrade.IsNull() {
		data.AllowWebSocketUpgrade = tfutils.BoolFromString(value.String())
	} else if data.AllowWebSocketUpgrade.ValueBool() {
		data.AllowWebSocketUpgrade = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WebSocketIdleTimeout`); value.Exists() && !data.WebSocketIdleTimeout.IsNull() {
		data.WebSocketIdleTimeout = types.Int64Value(value.Int())
	} else {
		data.WebSocketIdleTimeout = types.Int64Null()
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
	} else {
		data.MaxHdrCount = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxNameHdrLen`); value.Exists() && !data.MaxNameHdrLen.IsNull() {
		data.MaxNameHdrLen = types.Int64Value(value.Int())
	} else {
		data.MaxNameHdrLen = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxValueHdrLen`); value.Exists() && !data.MaxValueHdrLen.IsNull() {
		data.MaxValueHdrLen = types.Int64Value(value.Int())
	} else {
		data.MaxValueHdrLen = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxQueryStringLen`); value.Exists() && !data.MaxQueryStringLen.IsNull() {
		data.MaxQueryStringLen = types.Int64Value(value.Int())
	} else {
		data.MaxQueryStringLen = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ACL`); value.Exists() && !data.Acl.IsNull() {
		data.Acl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Acl = types.StringNull()
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
	if value := res.Get(pathRoot + `HTTP2MaxStreams`); value.Exists() && !data.Http2MaxStreams.IsNull() {
		data.Http2MaxStreams = types.Int64Value(value.Int())
	} else if data.Http2MaxStreams.ValueInt64() != 100 {
		data.Http2MaxStreams = types.Int64Null()
	}
	if value := res.Get(pathRoot + `HTTP2MaxFrameSize`); value.Exists() && !data.Http2MaxFrameSize.IsNull() {
		data.Http2MaxFrameSize = types.Int64Value(value.Int())
	} else if data.Http2MaxFrameSize.ValueInt64() != 16384 {
		data.Http2MaxFrameSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `HTTP2StreamHeader`); value.Exists() && !data.Http2StreamHeader.IsNull() {
		data.Http2StreamHeader = tfutils.BoolFromString(value.String())
	} else if data.Http2StreamHeader.ValueBool() {
		data.Http2StreamHeader = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ChunkedEncoding`); value.Exists() && !data.ChunkedEncoding.IsNull() {
		data.ChunkedEncoding = tfutils.BoolFromString(value.String())
	} else if !data.ChunkedEncoding.ValueBool() {
		data.ChunkedEncoding = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HeaderTimeout`); value.Exists() && !data.HeaderTimeout.IsNull() {
		data.HeaderTimeout = types.Int64Value(value.Int())
	} else if data.HeaderTimeout.ValueInt64() != 30000 {
		data.HeaderTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `HTTP2IdleTimeout`); value.Exists() && !data.Http2IdleTimeout.IsNull() {
		data.Http2IdleTimeout = types.Int64Value(value.Int())
	} else {
		data.Http2IdleTimeout = types.Int64Null()
	}
}
