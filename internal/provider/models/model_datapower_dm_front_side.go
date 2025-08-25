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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmFrontSide struct {
	LocalAddress      types.String `tfsdk:"local_address"`
	LocalPort         types.Int64  `tfsdk:"local_port"`
	UseSsl            types.Bool   `tfsdk:"use_ssl"`
	CredentialCharset types.String `tfsdk:"credential_charset"`
}

var DmFrontSideObjectType = map[string]attr.Type{
	"local_address":      types.StringType,
	"local_port":         types.Int64Type,
	"use_ssl":            types.BoolType,
	"credential_charset": types.StringType,
}
var DmFrontSideObjectDefault = map[string]attr.Value{
	"local_address":      types.StringValue("0.0.0.0"),
	"local_port":         types.Int64Value(3000),
	"use_ssl":            types.BoolValue(false),
	"credential_charset": types.StringValue("protocol"),
}

func GetDmFrontSideDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmFrontSideDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"local_address": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The address on which the service listens. The default of 0 indicates that the service is active on all addresses. Click Select Alias to use an alias for this value. Local host aliases help to ease migration tasks between machines.", "", "").AddDefaultValue("0.0.0.0").String,
				Computed:            true,
			},
			"local_port": DataSourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the port monitored by the stateful raw XML over TCP service. Enter a value in the range 1 - 65535. The default value is 3000.", "", "").AddIntegerRange(1, 65535).AddDefaultValue("3000").String,
				Computed:            true,
			},
			"use_ssl": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use TLS for this server. Ensure that the TLS server profile or TLS SNI server profile is configured.", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"credential_charset": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the character encoding of the original basic authentication values. Basic authentication credentials are combined and base64 encoded in the authorization header of the request. The DataPower Gateway transcodes the contents of the authorization header to UTF-8. Defaults to Protocol that is ISO-8859-1, Latin 1.", "credential-charset", "").AddStringEnum("protocol", "ascii", "utf8", "big5", "cp1250", "cp1251", "cp1252", "cp1253", "cp1254", "cp1255", "cp1256", "cp1257", "cp1258", "euc_jp", "euc_kr", "gb18030", "gb2312", "iso2022_jp", "iso2022_kr", "iso8859_1", "iso8859_2", "iso8859_4", "iso8859_5", "iso8859_6", "iso8859_7", "iso8859_8", "iso8859_9", "iso8859_15", "sjis", "tis620", "unicode_le").AddDefaultValue("protocol").String,
				Computed:            true,
			},
		},
	}
	return DmFrontSideDataSourceSchema
}
func GetDmFrontSideResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmFrontSideResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"local_address": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The address on which the service listens. The default of 0 indicates that the service is active on all addresses. Click Select Alias to use an alias for this value. Local host aliases help to ease migration tasks between machines.", "", "").AddDefaultValue("0.0.0.0").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("0.0.0.0"),
			},
			"local_port": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the port monitored by the stateful raw XML over TCP service. Enter a value in the range 1 - 65535. The default value is 3000.", "", "").AddIntegerRange(1, 65535).AddDefaultValue("3000").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(3000),
			},
			"use_ssl": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use TLS for this server. Ensure that the TLS server profile or TLS SNI server profile is configured.", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"credential_charset": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the character encoding of the original basic authentication values. Basic authentication credentials are combined and base64 encoded in the authorization header of the request. The DataPower Gateway transcodes the contents of the authorization header to UTF-8. Defaults to Protocol that is ISO-8859-1, Latin 1.", "credential-charset", "").AddStringEnum("protocol", "ascii", "utf8", "big5", "cp1250", "cp1251", "cp1252", "cp1253", "cp1254", "cp1255", "cp1256", "cp1257", "cp1258", "euc_jp", "euc_kr", "gb18030", "gb2312", "iso2022_jp", "iso2022_kr", "iso8859_1", "iso8859_2", "iso8859_4", "iso8859_5", "iso8859_6", "iso8859_7", "iso8859_8", "iso8859_9", "iso8859_15", "sjis", "tis620", "unicode_le").AddDefaultValue("protocol").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("protocol", "ascii", "utf8", "big5", "cp1250", "cp1251", "cp1252", "cp1253", "cp1254", "cp1255", "cp1256", "cp1257", "cp1258", "euc_jp", "euc_kr", "gb18030", "gb2312", "iso2022_jp", "iso2022_kr", "iso8859_1", "iso8859_2", "iso8859_4", "iso8859_5", "iso8859_6", "iso8859_7", "iso8859_8", "iso8859_9", "iso8859_15", "sjis", "tis620", "unicode_le"),
				},
				Default: stringdefault.StaticString("protocol"),
			},
		},
	}
	return DmFrontSideResourceSchema
}

func (data DmFrontSide) IsNull() bool {
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
	return true
}

func (data DmFrontSide) ToBody(ctx context.Context, pathRoot string) string {
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
	return body
}

func (data *DmFrontSide) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
		data.LocalPort = types.Int64Value(3000)
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
}

func (data *DmFrontSide) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	} else if data.LocalPort.ValueInt64() != 3000 {
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
}
