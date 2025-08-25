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

type DmWeekdayBitmap struct {
	Sunday    types.Bool `tfsdk:"sunday"`
	Monday    types.Bool `tfsdk:"monday"`
	Tuesday   types.Bool `tfsdk:"tuesday"`
	Wednesday types.Bool `tfsdk:"wednesday"`
	Thursday  types.Bool `tfsdk:"thursday"`
	Friday    types.Bool `tfsdk:"friday"`
	Saturday  types.Bool `tfsdk:"saturday"`
}

var DmWeekdayBitmapObjectType = map[string]attr.Type{
	"sunday":    types.BoolType,
	"monday":    types.BoolType,
	"tuesday":   types.BoolType,
	"wednesday": types.BoolType,
	"thursday":  types.BoolType,
	"friday":    types.BoolType,
	"saturday":  types.BoolType,
}
var DmWeekdayBitmapObjectDefault = map[string]attr.Value{
	"sunday":    types.BoolValue(false),
	"monday":    types.BoolValue(false),
	"tuesday":   types.BoolValue(false),
	"wednesday": types.BoolValue(false),
	"thursday":  types.BoolValue(false),
	"friday":    types.BoolValue(false),
	"saturday":  types.BoolValue(false),
}

func GetDmWeekdayBitmapDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmWeekdayBitmapDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"sunday": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Sunday", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"monday": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Monday", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"tuesday": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Tuesday", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"wednesday": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Wednesday", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"thursday": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Thursday", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"friday": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Friday", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"saturday": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Saturday", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
		},
	}
	DmWeekdayBitmapDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmWeekdayBitmapDataSourceSchema
}
func GetDmWeekdayBitmapResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmWeekdayBitmapResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmWeekdayBitmapObjectType,
				DmWeekdayBitmapObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"sunday": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Sunday", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"monday": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Monday", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"tuesday": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Tuesday", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wednesday": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Wednesday", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"thursday": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Thursday", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"friday": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Friday", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"saturday": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Saturday", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
	DmWeekdayBitmapResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmWeekdayBitmapResourceSchema.Required = true
	} else {
		DmWeekdayBitmapResourceSchema.Optional = true
		DmWeekdayBitmapResourceSchema.Computed = true
	}
	return DmWeekdayBitmapResourceSchema
}

func (data DmWeekdayBitmap) IsNull() bool {
	if !data.Sunday.IsNull() {
		return false
	}
	if !data.Monday.IsNull() {
		return false
	}
	if !data.Tuesday.IsNull() {
		return false
	}
	if !data.Wednesday.IsNull() {
		return false
	}
	if !data.Thursday.IsNull() {
		return false
	}
	if !data.Friday.IsNull() {
		return false
	}
	if !data.Saturday.IsNull() {
		return false
	}
	return true
}

func (data DmWeekdayBitmap) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Sunday.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Sunday`, tfutils.StringFromBool(data.Sunday, ""))
	}
	if !data.Monday.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Monday`, tfutils.StringFromBool(data.Monday, ""))
	}
	if !data.Tuesday.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Tuesday`, tfutils.StringFromBool(data.Tuesday, ""))
	}
	if !data.Wednesday.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Wednesday`, tfutils.StringFromBool(data.Wednesday, ""))
	}
	if !data.Thursday.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Thursday`, tfutils.StringFromBool(data.Thursday, ""))
	}
	if !data.Friday.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Friday`, tfutils.StringFromBool(data.Friday, ""))
	}
	if !data.Saturday.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Saturday`, tfutils.StringFromBool(data.Saturday, ""))
	}
	return body
}

func (data *DmWeekdayBitmap) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Sunday`); value.Exists() {
		data.Sunday = tfutils.BoolFromString(value.String())
	} else {
		data.Sunday = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Monday`); value.Exists() {
		data.Monday = tfutils.BoolFromString(value.String())
	} else {
		data.Monday = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Tuesday`); value.Exists() {
		data.Tuesday = tfutils.BoolFromString(value.String())
	} else {
		data.Tuesday = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Wednesday`); value.Exists() {
		data.Wednesday = tfutils.BoolFromString(value.String())
	} else {
		data.Wednesday = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Thursday`); value.Exists() {
		data.Thursday = tfutils.BoolFromString(value.String())
	} else {
		data.Thursday = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Friday`); value.Exists() {
		data.Friday = tfutils.BoolFromString(value.String())
	} else {
		data.Friday = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Saturday`); value.Exists() {
		data.Saturday = tfutils.BoolFromString(value.String())
	} else {
		data.Saturday = types.BoolNull()
	}
}

func (data *DmWeekdayBitmap) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Sunday`); value.Exists() && !data.Sunday.IsNull() {
		data.Sunday = tfutils.BoolFromString(value.String())
	} else if data.Sunday.ValueBool() {
		data.Sunday = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Monday`); value.Exists() && !data.Monday.IsNull() {
		data.Monday = tfutils.BoolFromString(value.String())
	} else if data.Monday.ValueBool() {
		data.Monday = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Tuesday`); value.Exists() && !data.Tuesday.IsNull() {
		data.Tuesday = tfutils.BoolFromString(value.String())
	} else if data.Tuesday.ValueBool() {
		data.Tuesday = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Wednesday`); value.Exists() && !data.Wednesday.IsNull() {
		data.Wednesday = tfutils.BoolFromString(value.String())
	} else if data.Wednesday.ValueBool() {
		data.Wednesday = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Thursday`); value.Exists() && !data.Thursday.IsNull() {
		data.Thursday = tfutils.BoolFromString(value.String())
	} else if data.Thursday.ValueBool() {
		data.Thursday = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Friday`); value.Exists() && !data.Friday.IsNull() {
		data.Friday = tfutils.BoolFromString(value.String())
	} else if data.Friday.ValueBool() {
		data.Friday = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Saturday`); value.Exists() && !data.Saturday.IsNull() {
		data.Saturday = tfutils.BoolFromString(value.String())
	} else if data.Saturday.ValueBool() {
		data.Saturday = types.BoolNull()
	}
}
