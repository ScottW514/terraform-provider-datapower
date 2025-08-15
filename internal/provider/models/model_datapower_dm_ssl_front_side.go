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

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	DataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	ResourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmSSLFrontSide struct {
	LocalAddress        types.String `tfsdk:"local_address"`
	LocalPort           types.Int64  `tfsdk:"local_port"`
	UseSsl              types.Bool   `tfsdk:"use_ssl"`
	CredentialCharset   types.String `tfsdk:"credential_charset"`
	SslServerConfigType types.String `tfsdk:"ssl_server_config_type"`
	SslServer           types.String `tfsdk:"ssl_server"`
	SslsniServer        types.String `tfsdk:"sslsni_server"`
}

var DmSSLFrontSideObjectType = map[string]attr.Type{
	"local_address":          types.StringType,
	"local_port":             types.Int64Type,
	"use_ssl":                types.BoolType,
	"credential_charset":     types.StringType,
	"ssl_server_config_type": types.StringType,
	"ssl_server":             types.StringType,
	"sslsni_server":          types.StringType,
}
var DmSSLFrontSideObjectDefault = map[string]attr.Value{
	"local_address":          types.StringValue("0.0.0.0"),
	"local_port":             types.Int64Null(),
	"use_ssl":                types.BoolValue(false),
	"credential_charset":     types.StringValue("protocol"),
	"ssl_server_config_type": types.StringValue("server"),
	"ssl_server":             types.StringNull(),
	"sslsni_server":          types.StringNull(),
}
var DmSSLFrontSideDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"local_address": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the host alias or local IP address on the appliance that the service listens on. The default value is 0.0.0.0, which denotes all local addresses.", "address", "").AddDefaultValue("0.0.0.0").String,
			Computed:            true,
		},
		"local_port": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the unique, local TCP port that the service listens on. No other service on the DataPower Gateway can use this port. Enter any value in the range 1 - 65535.", "port", "").AddIntegerRange(1, 65535).String,
			Computed:            true,
		},
		"use_ssl": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Indicate whether to use the assigned TLS profile to control connections to this TCP port. When enabled, the service expects HTTPS requests on this port.", "use-ssl", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"credential_charset": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specifies the character encoding of the original basic authentication values. Basic authentication credentials are combined and base64 encoded in the authorization header of the request. The appliance transcodes the contents of the authorization header to UTF-8. Defaults to Protocol which is ISO-8859-1, Latin 1.", "credential-charset", "").AddStringEnum("protocol", "ascii", "utf8", "big5", "cp1250", "cp1251", "cp1252", "cp1253", "cp1254", "cp1255", "cp1256", "cp1257", "cp1258", "euc_jp", "euc_kr", "gb18030", "gb2312", "iso2022_jp", "iso2022_kr", "iso8859_1", "iso8859_2", "iso8859_4", "iso8859_5", "iso8859_6", "iso8859_7", "iso8859_8", "iso8859_9", "iso8859_15", "sjis", "tis620", "unicode_le").AddDefaultValue("protocol").String,
			Computed:            true,
		},
		"ssl_server_config_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS profile type to secure connections between clients and the DataPower Gateway", "ssl-config-type", "").AddStringEnum("server", "sni").AddDefaultValue("server").String,
			Computed:            true,
		},
		"ssl_server": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS server profile to secure connections between clients and the DataPower Gateway", "ssl-server", "ssl_server_profile").String,
			Computed:            true,
		},
		"sslsni_server": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS SNI server profile to secure connections between clients and the DataPower Gateway", "ssl-sni-server", "ssl_sni_server_profile").String,
			Computed:            true,
		},
	},
}
var DmSSLFrontSideResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"local_address": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the host alias or local IP address on the appliance that the service listens on. The default value is 0.0.0.0, which denotes all local addresses.", "address", "").AddDefaultValue("0.0.0.0").String,
			Computed:            true,
			Optional:            true,
			Default:             stringdefault.StaticString("0.0.0.0"),
		},
		"local_port": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the unique, local TCP port that the service listens on. No other service on the DataPower Gateway can use this port. Enter any value in the range 1 - 65535.", "port", "").AddIntegerRange(1, 65535).String,
			Required:            true,
			Validators: []validator.Int64{
				int64validator.Between(1, 65535),
			},
		},
		"use_ssl": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Indicate whether to use the assigned TLS profile to control connections to this TCP port. When enabled, the service expects HTTPS requests on this port.", "use-ssl", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"credential_charset": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specifies the character encoding of the original basic authentication values. Basic authentication credentials are combined and base64 encoded in the authorization header of the request. The appliance transcodes the contents of the authorization header to UTF-8. Defaults to Protocol which is ISO-8859-1, Latin 1.", "credential-charset", "").AddStringEnum("protocol", "ascii", "utf8", "big5", "cp1250", "cp1251", "cp1252", "cp1253", "cp1254", "cp1255", "cp1256", "cp1257", "cp1258", "euc_jp", "euc_kr", "gb18030", "gb2312", "iso2022_jp", "iso2022_kr", "iso8859_1", "iso8859_2", "iso8859_4", "iso8859_5", "iso8859_6", "iso8859_7", "iso8859_8", "iso8859_9", "iso8859_15", "sjis", "tis620", "unicode_le").AddDefaultValue("protocol").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("protocol", "ascii", "utf8", "big5", "cp1250", "cp1251", "cp1252", "cp1253", "cp1254", "cp1255", "cp1256", "cp1257", "cp1258", "euc_jp", "euc_kr", "gb18030", "gb2312", "iso2022_jp", "iso2022_kr", "iso8859_1", "iso8859_2", "iso8859_4", "iso8859_5", "iso8859_6", "iso8859_7", "iso8859_8", "iso8859_9", "iso8859_15", "sjis", "tis620", "unicode_le"),
			},
			Default: stringdefault.StaticString("protocol"),
		},
		"ssl_server_config_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS profile type to secure connections between clients and the DataPower Gateway", "ssl-config-type", "").AddStringEnum("server", "sni").AddDefaultValue("server").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("server", "sni"),
			},
			Default: stringdefault.StaticString("server"),
		},
		"ssl_server": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS server profile to secure connections between clients and the DataPower Gateway", "ssl-server", "ssl_server_profile").String,
			Optional:            true,
		},
		"sslsni_server": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS SNI server profile to secure connections between clients and the DataPower Gateway", "ssl-sni-server", "ssl_sni_server_profile").String,
			Optional:            true,
		},
	},
}

func (data DmSSLFrontSide) IsNull() bool {
	if !data.LocalAddress.IsNull() {
		return false
	}
	if !data.LocalPort.IsNull() {
		return false
	}
	if !data.UseSsl.IsNull() {
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
	if !data.SslsniServer.IsNull() {
		return false
	}
	return true
}

func (data DmSSLFrontSide) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.LocalAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalAddress`, data.LocalAddress.ValueString())
	}
	if !data.LocalPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalPort`, data.LocalPort.ValueInt64())
	}
	if !data.UseSsl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseSSL`, tfutils.StringFromBool(data.UseSsl, ""))
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
	if !data.SslsniServer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLSNIServer`, data.SslsniServer.ValueString())
	}
	return body
}

func (data *DmSSLFrontSide) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalAddress = types.StringValue("0.0.0.0")
	}
	if value := res.Get(pathRoot + `LocalPort`); value.Exists() {
		data.LocalPort = types.Int64Value(value.Int())
	} else {
		data.LocalPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `UseSSL`); value.Exists() {
		data.UseSsl = tfutils.BoolFromString(value.String())
	} else {
		data.UseSsl = types.BoolNull()
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
		data.SslsniServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslsniServer = types.StringNull()
	}
}

func (data *DmSSLFrontSide) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && !data.LocalAddress.IsNull() {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else if data.LocalAddress.ValueString() != "0.0.0.0" {
		data.LocalAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalPort`); value.Exists() && !data.LocalPort.IsNull() {
		data.LocalPort = types.Int64Value(value.Int())
	} else {
		data.LocalPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `UseSSL`); value.Exists() && !data.UseSsl.IsNull() {
		data.UseSsl = tfutils.BoolFromString(value.String())
	} else if data.UseSsl.ValueBool() {
		data.UseSsl = types.BoolNull()
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
	if value := res.Get(pathRoot + `SSLSNIServer`); value.Exists() && !data.SslsniServer.IsNull() {
		data.SslsniServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslsniServer = types.StringNull()
	}
}
