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

type AS2ProxySourceProtocolHandler struct {
	Id                            types.String                `tfsdk:"id"`
	AppDomain                     types.String                `tfsdk:"app_domain"`
	UserSummary                   types.String                `tfsdk:"user_summary"`
	LocalAddress                  types.String                `tfsdk:"local_address"`
	LocalPort                     types.Int64                 `tfsdk:"local_port"`
	HttpVersion                   types.String                `tfsdk:"http_version"`
	AllowedFeatures               *DmSourceAS2FeatureType     `tfsdk:"allowed_features"`
	PersistentConnections         types.Bool                  `tfsdk:"persistent_connections"`
	MaxPersistentConnectionsReuse types.Int64                 `tfsdk:"max_persistent_connections_reuse"`
	AllowCompression              types.Bool                  `tfsdk:"allow_compression"`
	MaxUrlLen                     types.Int64                 `tfsdk:"max_url_len"`
	MaxTotalHdrLen                types.Int64                 `tfsdk:"max_total_hdr_len"`
	MaxHdrCount                   types.Int64                 `tfsdk:"max_hdr_count"`
	MaxNameHdrLen                 types.Int64                 `tfsdk:"max_name_hdr_len"`
	MaxValueHdrLen                types.Int64                 `tfsdk:"max_value_hdr_len"`
	Acl                           types.String                `tfsdk:"acl"`
	CredentialCharset             types.String                `tfsdk:"credential_charset"`
	RemoteAddress                 types.String                `tfsdk:"remote_address"`
	RemotePort                    types.Int64                 `tfsdk:"remote_port"`
	RemoteConnectionTimeout       types.Int64                 `tfsdk:"remote_connection_timeout"`
	XmlManager                    types.String                `tfsdk:"xml_manager"`
	EnablePassthrough             types.Bool                  `tfsdk:"enable_passthrough"`
	EnableVisibilityEvent         types.Bool                  `tfsdk:"enable_visibility_event"`
	VisibilityEventEndpoint       types.String                `tfsdk:"visibility_event_endpoint"`
	EnableHmacAuthentication      types.Bool                  `tfsdk:"enable_hmac_authentication"`
	HmacPassphraseAlias           types.String                `tfsdk:"hmac_passphrase_alias"`
	SslServerConfigType           types.String                `tfsdk:"ssl_server_config_type"`
	SslServer                     types.String                `tfsdk:"ssl_server"`
	SslSniServer                  types.String                `tfsdk:"ssl_sni_server"`
	SslClientConfigType           types.String                `tfsdk:"ssl_client_config_type"`
	SslClient                     types.String                `tfsdk:"ssl_client"`
	DependencyActions             []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var AS2ProxySourceProtocolHandlerVisibilityEventEndpointCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "enable_visibility_event",
	AttrType:    "Bool",
	AttrDefault: "true",
	Value:       []string{"true"},
}
var AS2ProxySourceProtocolHandlerHmacPassphraseAliasCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "enable_visibility_event",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "enable_hmac_authentication",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
	},
}

var AS2ProxySourceProtocolHandlerObjectType = map[string]attr.Type{
	"id":                               types.StringType,
	"app_domain":                       types.StringType,
	"user_summary":                     types.StringType,
	"local_address":                    types.StringType,
	"local_port":                       types.Int64Type,
	"http_version":                     types.StringType,
	"allowed_features":                 types.ObjectType{AttrTypes: DmSourceAS2FeatureTypeObjectType},
	"persistent_connections":           types.BoolType,
	"max_persistent_connections_reuse": types.Int64Type,
	"allow_compression":                types.BoolType,
	"max_url_len":                      types.Int64Type,
	"max_total_hdr_len":                types.Int64Type,
	"max_hdr_count":                    types.Int64Type,
	"max_name_hdr_len":                 types.Int64Type,
	"max_value_hdr_len":                types.Int64Type,
	"acl":                              types.StringType,
	"credential_charset":               types.StringType,
	"remote_address":                   types.StringType,
	"remote_port":                      types.Int64Type,
	"remote_connection_timeout":        types.Int64Type,
	"xml_manager":                      types.StringType,
	"enable_passthrough":               types.BoolType,
	"enable_visibility_event":          types.BoolType,
	"visibility_event_endpoint":        types.StringType,
	"enable_hmac_authentication":       types.BoolType,
	"hmac_passphrase_alias":            types.StringType,
	"ssl_server_config_type":           types.StringType,
	"ssl_server":                       types.StringType,
	"ssl_sni_server":                   types.StringType,
	"ssl_client_config_type":           types.StringType,
	"ssl_client":                       types.StringType,
	"dependency_actions":               actions.ActionsListType,
}

func (data AS2ProxySourceProtocolHandler) GetPath() string {
	rest_path := "/mgmt/config/{domain}/AS2ProxySourceProtocolHandler"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data AS2ProxySourceProtocolHandler) IsNull() bool {
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
	if !data.Acl.IsNull() {
		return false
	}
	if !data.CredentialCharset.IsNull() {
		return false
	}
	if !data.RemoteAddress.IsNull() {
		return false
	}
	if !data.RemotePort.IsNull() {
		return false
	}
	if !data.RemoteConnectionTimeout.IsNull() {
		return false
	}
	if !data.XmlManager.IsNull() {
		return false
	}
	if !data.EnablePassthrough.IsNull() {
		return false
	}
	if !data.EnableVisibilityEvent.IsNull() {
		return false
	}
	if !data.VisibilityEventEndpoint.IsNull() {
		return false
	}
	if !data.EnableHmacAuthentication.IsNull() {
		return false
	}
	if !data.HmacPassphraseAlias.IsNull() {
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
	if !data.SslClientConfigType.IsNull() {
		return false
	}
	if !data.SslClient.IsNull() {
		return false
	}
	return true
}

func (data AS2ProxySourceProtocolHandler) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Acl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ACL`, data.Acl.ValueString())
	}
	if !data.CredentialCharset.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CredentialCharset`, data.CredentialCharset.ValueString())
	}
	if !data.RemoteAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemoteAddress`, data.RemoteAddress.ValueString())
	}
	if !data.RemotePort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemotePort`, data.RemotePort.ValueInt64())
	}
	if !data.RemoteConnectionTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemoteConnectionTimeout`, data.RemoteConnectionTimeout.ValueInt64())
	}
	if !data.XmlManager.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XMLManager`, data.XmlManager.ValueString())
	}
	if !data.EnablePassthrough.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnablePassthrough`, tfutils.StringFromBool(data.EnablePassthrough, ""))
	}
	if !data.EnableVisibilityEvent.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnableVisibilityEvent`, tfutils.StringFromBool(data.EnableVisibilityEvent, ""))
	}
	if !data.VisibilityEventEndpoint.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`VisibilityEventEndpoint`, data.VisibilityEventEndpoint.ValueString())
	}
	if !data.EnableHmacAuthentication.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnableHmacAuthentication`, tfutils.StringFromBool(data.EnableHmacAuthentication, ""))
	}
	if !data.HmacPassphraseAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HmacPassphraseAlias`, data.HmacPassphraseAlias.ValueString())
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
	if !data.SslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClientConfigType`, data.SslClientConfigType.ValueString())
	}
	if !data.SslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClient`, data.SslClient.ValueString())
	}
	return body
}

func (data *AS2ProxySourceProtocolHandler) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `AllowedFeatures`); value.Exists() {
		data.AllowedFeatures = &DmSourceAS2FeatureType{}
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
	if value := res.Get(pathRoot + `RemoteConnectionTimeout`); value.Exists() {
		data.RemoteConnectionTimeout = types.Int64Value(value.Int())
	} else {
		data.RemoteConnectionTimeout = types.Int64Value(60)
	}
	if value := res.Get(pathRoot + `XMLManager`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.XmlManager = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XmlManager = types.StringValue("default")
	}
	if value := res.Get(pathRoot + `EnablePassthrough`); value.Exists() {
		data.EnablePassthrough = tfutils.BoolFromString(value.String())
	} else {
		data.EnablePassthrough = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EnableVisibilityEvent`); value.Exists() {
		data.EnableVisibilityEvent = tfutils.BoolFromString(value.String())
	} else {
		data.EnableVisibilityEvent = types.BoolNull()
	}
	if value := res.Get(pathRoot + `VisibilityEventEndpoint`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.VisibilityEventEndpoint = tfutils.ParseStringFromGJSON(value)
	} else {
		data.VisibilityEventEndpoint = types.StringNull()
	}
	if value := res.Get(pathRoot + `EnableHmacAuthentication`); value.Exists() {
		data.EnableHmacAuthentication = tfutils.BoolFromString(value.String())
	} else {
		data.EnableHmacAuthentication = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HmacPassphraseAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HmacPassphraseAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HmacPassphraseAlias = types.StringNull()
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
}

func (data *AS2ProxySourceProtocolHandler) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `RemoteConnectionTimeout`); value.Exists() && !data.RemoteConnectionTimeout.IsNull() {
		data.RemoteConnectionTimeout = types.Int64Value(value.Int())
	} else if data.RemoteConnectionTimeout.ValueInt64() != 60 {
		data.RemoteConnectionTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `XMLManager`); value.Exists() && !data.XmlManager.IsNull() {
		data.XmlManager = tfutils.ParseStringFromGJSON(value)
	} else if data.XmlManager.ValueString() != "default" {
		data.XmlManager = types.StringNull()
	}
	if value := res.Get(pathRoot + `EnablePassthrough`); value.Exists() && !data.EnablePassthrough.IsNull() {
		data.EnablePassthrough = tfutils.BoolFromString(value.String())
	} else if !data.EnablePassthrough.ValueBool() {
		data.EnablePassthrough = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EnableVisibilityEvent`); value.Exists() && !data.EnableVisibilityEvent.IsNull() {
		data.EnableVisibilityEvent = tfutils.BoolFromString(value.String())
	} else if !data.EnableVisibilityEvent.ValueBool() {
		data.EnableVisibilityEvent = types.BoolNull()
	}
	if value := res.Get(pathRoot + `VisibilityEventEndpoint`); value.Exists() && !data.VisibilityEventEndpoint.IsNull() {
		data.VisibilityEventEndpoint = tfutils.ParseStringFromGJSON(value)
	} else {
		data.VisibilityEventEndpoint = types.StringNull()
	}
	if value := res.Get(pathRoot + `EnableHmacAuthentication`); value.Exists() && !data.EnableHmacAuthentication.IsNull() {
		data.EnableHmacAuthentication = tfutils.BoolFromString(value.String())
	} else if !data.EnableHmacAuthentication.ValueBool() {
		data.EnableHmacAuthentication = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HmacPassphraseAlias`); value.Exists() && !data.HmacPassphraseAlias.IsNull() {
		data.HmacPassphraseAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HmacPassphraseAlias = types.StringNull()
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
}
