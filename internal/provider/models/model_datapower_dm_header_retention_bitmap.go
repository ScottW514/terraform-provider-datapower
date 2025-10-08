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

type DmHeaderRetentionBitmap struct {
	Te             types.Bool `tfsdk:"te"`
	AcceptEncoding types.Bool `tfsdk:"accept_encoding"`
	Range          types.Bool `tfsdk:"range"`
	Mqmd           types.Bool `tfsdk:"mqmd"`
}

var DmHeaderRetentionBitmapObjectType = map[string]attr.Type{
	"te":              types.BoolType,
	"accept_encoding": types.BoolType,
	"range":           types.BoolType,
	"mqmd":            types.BoolType,
}
var DmHeaderRetentionBitmapObjectDefault = map[string]attr.Value{
	"te":              types.BoolValue(false),
	"accept_encoding": types.BoolValue(false),
	"range":           types.BoolValue(false),
	"mqmd":            types.BoolValue(false),
}

func GetDmHeaderRetentionBitmapDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmHeaderRetentionBitmapDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"te": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "TE",
				Computed:            true,
			},
			"accept_encoding": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Accept-Encoding",
				Computed:            true,
			},
			"range": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Range",
				Computed:            true,
			},
			"mqmd": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "MQMD",
				Computed:            true,
			},
		},
	}
	DmHeaderRetentionBitmapDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmHeaderRetentionBitmapDataSourceSchema
}
func GetDmHeaderRetentionBitmapResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmHeaderRetentionBitmapResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmHeaderRetentionBitmapObjectType,
				DmHeaderRetentionBitmapObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"te": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TE", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"accept_encoding": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Accept-Encoding", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"range": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Range", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"mqmd": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("MQMD", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
	DmHeaderRetentionBitmapResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmHeaderRetentionBitmapResourceSchema.Required = true
	} else {
		DmHeaderRetentionBitmapResourceSchema.Optional = true
		DmHeaderRetentionBitmapResourceSchema.Computed = true
	}
	return DmHeaderRetentionBitmapResourceSchema
}

func (data DmHeaderRetentionBitmap) IsNull() bool {
	if !data.Te.IsNull() {
		return false
	}
	if !data.AcceptEncoding.IsNull() {
		return false
	}
	if !data.Range.IsNull() {
		return false
	}
	if !data.Mqmd.IsNull() {
		return false
	}
	return true
}

func (data DmHeaderRetentionBitmap) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Te.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TE`, tfutils.StringFromBool(data.Te, ""))
	}
	if !data.AcceptEncoding.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Accept-Encoding`, tfutils.StringFromBool(data.AcceptEncoding, ""))
	}
	if !data.Range.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Range`, tfutils.StringFromBool(data.Range, ""))
	}
	if !data.Mqmd.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MQMD`, tfutils.StringFromBool(data.Mqmd, ""))
	}
	return body
}

func (data *DmHeaderRetentionBitmap) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `TE`); value.Exists() {
		data.Te = tfutils.BoolFromString(value.String())
	} else {
		data.Te = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Accept-Encoding`); value.Exists() {
		data.AcceptEncoding = tfutils.BoolFromString(value.String())
	} else {
		data.AcceptEncoding = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Range`); value.Exists() {
		data.Range = tfutils.BoolFromString(value.String())
	} else {
		data.Range = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MQMD`); value.Exists() {
		data.Mqmd = tfutils.BoolFromString(value.String())
	} else {
		data.Mqmd = types.BoolNull()
	}
}

func (data *DmHeaderRetentionBitmap) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `TE`); value.Exists() && !data.Te.IsNull() {
		data.Te = tfutils.BoolFromString(value.String())
	} else if data.Te.ValueBool() {
		data.Te = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Accept-Encoding`); value.Exists() && !data.AcceptEncoding.IsNull() {
		data.AcceptEncoding = tfutils.BoolFromString(value.String())
	} else if data.AcceptEncoding.ValueBool() {
		data.AcceptEncoding = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Range`); value.Exists() && !data.Range.IsNull() {
		data.Range = tfutils.BoolFromString(value.String())
	} else if data.Range.ValueBool() {
		data.Range = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MQMD`); value.Exists() && !data.Mqmd.IsNull() {
		data.Mqmd = tfutils.BoolFromString(value.String())
	} else if data.Mqmd.ValueBool() {
		data.Mqmd = types.BoolNull()
	}
}
