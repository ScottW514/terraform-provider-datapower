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

type DmLTPATokenVersion struct {
	Ltpa       types.Bool `tfsdk:"ltpa"`
	Ltpa2      types.Bool `tfsdk:"ltpa2"`
	LtpaDomino types.Bool `tfsdk:"ltpa_domino"`
}

var DmLTPATokenVersionObjectType = map[string]attr.Type{
	"ltpa":        types.BoolType,
	"ltpa2":       types.BoolType,
	"ltpa_domino": types.BoolType,
}
var DmLTPATokenVersionObjectDefault = map[string]attr.Value{
	"ltpa":        types.BoolValue(false),
	"ltpa2":       types.BoolValue(true),
	"ltpa_domino": types.BoolValue(false),
}
var DmLTPATokenVersionDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"ltpa": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WebSphere version 1", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ltpa2": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WebSphere version 2", "", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"ltpa_domino": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Domino", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
	},
}
var DmLTPATokenVersionResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmLTPATokenVersionObjectType,
			DmLTPATokenVersionObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"ltpa": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WebSphere version 1", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ltpa2": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WebSphere version 2", "", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"ltpa_domino": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Domino", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
	},
}

func (data DmLTPATokenVersion) IsNull() bool {
	if !data.Ltpa.IsNull() {
		return false
	}
	if !data.Ltpa2.IsNull() {
		return false
	}
	if !data.LtpaDomino.IsNull() {
		return false
	}
	return true
}
func GetDmLTPATokenVersionDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmLTPATokenVersionDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmLTPATokenVersionDataSourceSchema
}

func GetDmLTPATokenVersionResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmLTPATokenVersionResourceSchema.Required = true
	} else {
		DmLTPATokenVersionResourceSchema.Optional = true
		DmLTPATokenVersionResourceSchema.Computed = true
	}
	DmLTPATokenVersionResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmLTPATokenVersionResourceSchema
}

func (data DmLTPATokenVersion) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Ltpa.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LTPA`, tfutils.StringFromBool(data.Ltpa, false))
	}
	if !data.Ltpa2.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LTPA2`, tfutils.StringFromBool(data.Ltpa2, false))
	}
	if !data.LtpaDomino.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LTPADomino`, tfutils.StringFromBool(data.LtpaDomino, false))
	}
	return body
}

func (data *DmLTPATokenVersion) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `LTPA`); value.Exists() {
		data.Ltpa = tfutils.BoolFromString(value.String())
	} else {
		data.Ltpa = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LTPA2`); value.Exists() {
		data.Ltpa2 = tfutils.BoolFromString(value.String())
	} else {
		data.Ltpa2 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LTPADomino`); value.Exists() {
		data.LtpaDomino = tfutils.BoolFromString(value.String())
	} else {
		data.LtpaDomino = types.BoolNull()
	}
}

func (data *DmLTPATokenVersion) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `LTPA`); value.Exists() && !data.Ltpa.IsNull() {
		data.Ltpa = tfutils.BoolFromString(value.String())
	} else if data.Ltpa.ValueBool() {
		data.Ltpa = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LTPA2`); value.Exists() && !data.Ltpa2.IsNull() {
		data.Ltpa2 = tfutils.BoolFromString(value.String())
	} else if !data.Ltpa2.ValueBool() {
		data.Ltpa2 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LTPADomino`); value.Exists() && !data.LtpaDomino.IsNull() {
		data.LtpaDomino = tfutils.BoolFromString(value.String())
	} else if data.LtpaDomino.ValueBool() {
		data.LtpaDomino = types.BoolNull()
	}
}
