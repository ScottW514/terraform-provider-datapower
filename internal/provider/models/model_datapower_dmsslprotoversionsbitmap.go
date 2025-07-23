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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	DataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	ResourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmSSLProtoVersionsBitmap struct {
	SslV3   types.Bool `tfsdk:"ssl_v3"`
	TlsV1d0 types.Bool `tfsdk:"tls_v1d0"`
	TlsV1d1 types.Bool `tfsdk:"tls_v1d1"`
	TlsV1d2 types.Bool `tfsdk:"tls_v1d2"`
	TlsV1d3 types.Bool `tfsdk:"tls_v1d3"`
}

var DmSSLProtoVersionsBitmapObjectType = map[string]attr.Type{
	"ssl_v3":   types.BoolType,
	"tls_v1d0": types.BoolType,
	"tls_v1d1": types.BoolType,
	"tls_v1d2": types.BoolType,
	"tls_v1d3": types.BoolType,
}
var DmSSLProtoVersionsBitmapObjectDefault = map[string]attr.Value{
	"ssl_v3":   types.BoolValue(false),
	"tls_v1d0": types.BoolValue(false),
	"tls_v1d1": types.BoolValue(true),
	"tls_v1d2": types.BoolValue(true),
	"tls_v1d3": types.BoolValue(true),
}
var DmSSLProtoVersionsBitmapDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"ssl_v3": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enable SSL version 3", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"tls_v1d0": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enable TLS version 1.0", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"tls_v1d1": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enable TLS version 1.1", "", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"tls_v1d2": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enable TLS version 1.2", "", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"tls_v1d3": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enable TLS version 1.3", "", "").AddDefaultValue("true").String,
			Computed:            true,
		},
	},
}
var DmSSLProtoVersionsBitmapResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmSSLProtoVersionsBitmapObjectType,
			DmSSLProtoVersionsBitmapObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"ssl_v3": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enable SSL version 3", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"tls_v1d0": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enable TLS version 1.0", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"tls_v1d1": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enable TLS version 1.1", "", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"tls_v1d2": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enable TLS version 1.2", "", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"tls_v1d3": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enable TLS version 1.3", "", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
	},
}

func (data DmSSLProtoVersionsBitmap) IsNull() bool {
	if !data.SslV3.IsNull() {
		return false
	}
	if !data.TlsV1d0.IsNull() {
		return false
	}
	if !data.TlsV1d1.IsNull() {
		return false
	}
	if !data.TlsV1d2.IsNull() {
		return false
	}
	if !data.TlsV1d3.IsNull() {
		return false
	}
	return true
}
func GetDmSSLProtoVersionsBitmapDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmSSLProtoVersionsBitmapDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmSSLProtoVersionsBitmapDataSourceSchema
}

func GetDmSSLProtoVersionsBitmapResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmSSLProtoVersionsBitmapResourceSchema.Required = true
	} else {
		DmSSLProtoVersionsBitmapResourceSchema.Optional = true
		DmSSLProtoVersionsBitmapResourceSchema.Computed = true
	}
	DmSSLProtoVersionsBitmapResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmSSLProtoVersionsBitmapResourceSchema
}

func (data DmSSLProtoVersionsBitmap) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.SslV3.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLv3`, tfutils.StringFromBool(data.SslV3, ""))
	}
	if !data.TlsV1d0.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TLSv1d0`, tfutils.StringFromBool(data.TlsV1d0, ""))
	}
	if !data.TlsV1d1.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TLSv1d1`, tfutils.StringFromBool(data.TlsV1d1, ""))
	}
	if !data.TlsV1d2.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TLSv1d2`, tfutils.StringFromBool(data.TlsV1d2, ""))
	}
	if !data.TlsV1d3.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TLSv1d3`, tfutils.StringFromBool(data.TlsV1d3, ""))
	}
	return body
}

func (data *DmSSLProtoVersionsBitmap) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `SSLv3`); value.Exists() {
		data.SslV3 = tfutils.BoolFromString(value.String())
	} else {
		data.SslV3 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `TLSv1d0`); value.Exists() {
		data.TlsV1d0 = tfutils.BoolFromString(value.String())
	} else {
		data.TlsV1d0 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `TLSv1d1`); value.Exists() {
		data.TlsV1d1 = tfutils.BoolFromString(value.String())
	} else {
		data.TlsV1d1 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `TLSv1d2`); value.Exists() {
		data.TlsV1d2 = tfutils.BoolFromString(value.String())
	} else {
		data.TlsV1d2 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `TLSv1d3`); value.Exists() {
		data.TlsV1d3 = tfutils.BoolFromString(value.String())
	} else {
		data.TlsV1d3 = types.BoolNull()
	}
}

func (data *DmSSLProtoVersionsBitmap) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `SSLv3`); value.Exists() && !data.SslV3.IsNull() {
		data.SslV3 = tfutils.BoolFromString(value.String())
	} else if data.SslV3.ValueBool() {
		data.SslV3 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `TLSv1d0`); value.Exists() && !data.TlsV1d0.IsNull() {
		data.TlsV1d0 = tfutils.BoolFromString(value.String())
	} else if data.TlsV1d0.ValueBool() {
		data.TlsV1d0 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `TLSv1d1`); value.Exists() && !data.TlsV1d1.IsNull() {
		data.TlsV1d1 = tfutils.BoolFromString(value.String())
	} else if !data.TlsV1d1.ValueBool() {
		data.TlsV1d1 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `TLSv1d2`); value.Exists() && !data.TlsV1d2.IsNull() {
		data.TlsV1d2 = tfutils.BoolFromString(value.String())
	} else if !data.TlsV1d2.ValueBool() {
		data.TlsV1d2 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `TLSv1d3`); value.Exists() && !data.TlsV1d3.IsNull() {
		data.TlsV1d3 = tfutils.BoolFromString(value.String())
	} else if !data.TlsV1d3.ValueBool() {
		data.TlsV1d3 = types.BoolNull()
	}
}
