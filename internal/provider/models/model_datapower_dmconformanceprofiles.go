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

type DmConformanceProfiles struct {
	Bp10  types.Bool `tfsdk:"bp10"`
	Bp11  types.Bool `tfsdk:"bp11"`
	Ap10  types.Bool `tfsdk:"ap10"`
	Bsp10 types.Bool `tfsdk:"bsp10"`
}

var DmConformanceProfilesObjectType = map[string]attr.Type{
	"bp10":  types.BoolType,
	"bp11":  types.BoolType,
	"ap10":  types.BoolType,
	"bsp10": types.BoolType,
}
var DmConformanceProfilesObjectDefault = map[string]attr.Value{
	"bp10":  types.BoolValue(false),
	"bp11":  types.BoolValue(true),
	"ap10":  types.BoolValue(true),
	"bsp10": types.BoolValue(true),
}
var DmConformanceProfilesDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"bp10": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WS-I BP 1.0", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"bp11": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WS-I BP 1.1", "", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"ap10": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WS-I AP 1.0", "", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"bsp10": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WS-I BSP 1.0", "", "").AddDefaultValue("true").String,
			Computed:            true,
		},
	},
}
var DmConformanceProfilesResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmConformanceProfilesObjectType,
			DmConformanceProfilesObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"bp10": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WS-I BP 1.0", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"bp11": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WS-I BP 1.1", "", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"ap10": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WS-I AP 1.0", "", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"bsp10": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WS-I BSP 1.0", "", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
	},
}

func (data DmConformanceProfiles) IsNull() bool {
	if !data.Bp10.IsNull() {
		return false
	}
	if !data.Bp11.IsNull() {
		return false
	}
	if !data.Ap10.IsNull() {
		return false
	}
	if !data.Bsp10.IsNull() {
		return false
	}
	return true
}
func GetDmConformanceProfilesDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmConformanceProfilesDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmConformanceProfilesDataSourceSchema
}

func GetDmConformanceProfilesResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmConformanceProfilesResourceSchema.Required = true
	} else {
		DmConformanceProfilesResourceSchema.Optional = true
		DmConformanceProfilesResourceSchema.Computed = true
	}
	DmConformanceProfilesResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmConformanceProfilesResourceSchema
}

func (data DmConformanceProfiles) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Bp10.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BP10`, tfutils.StringFromBool(data.Bp10, ""))
	}
	if !data.Bp11.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BP11`, tfutils.StringFromBool(data.Bp11, ""))
	}
	if !data.Ap10.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AP10`, tfutils.StringFromBool(data.Ap10, ""))
	}
	if !data.Bsp10.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BSP10`, tfutils.StringFromBool(data.Bsp10, ""))
	}
	return body
}

func (data *DmConformanceProfiles) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `BP10`); value.Exists() {
		data.Bp10 = tfutils.BoolFromString(value.String())
	} else {
		data.Bp10 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `BP11`); value.Exists() {
		data.Bp11 = tfutils.BoolFromString(value.String())
	} else {
		data.Bp11 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AP10`); value.Exists() {
		data.Ap10 = tfutils.BoolFromString(value.String())
	} else {
		data.Ap10 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `BSP10`); value.Exists() {
		data.Bsp10 = tfutils.BoolFromString(value.String())
	} else {
		data.Bsp10 = types.BoolNull()
	}
}

func (data *DmConformanceProfiles) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `BP10`); value.Exists() && !data.Bp10.IsNull() {
		data.Bp10 = tfutils.BoolFromString(value.String())
	} else if data.Bp10.ValueBool() {
		data.Bp10 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `BP11`); value.Exists() && !data.Bp11.IsNull() {
		data.Bp11 = tfutils.BoolFromString(value.String())
	} else if !data.Bp11.ValueBool() {
		data.Bp11 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AP10`); value.Exists() && !data.Ap10.IsNull() {
		data.Ap10 = tfutils.BoolFromString(value.String())
	} else if !data.Ap10.ValueBool() {
		data.Ap10 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `BSP10`); value.Exists() && !data.Bsp10.IsNull() {
		data.Bsp10 = tfutils.BoolFromString(value.String())
	} else if !data.Bsp10.ValueBool() {
		data.Bsp10 = types.BoolNull()
	}
}
