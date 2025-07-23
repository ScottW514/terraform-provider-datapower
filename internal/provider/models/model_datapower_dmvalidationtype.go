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

type DmValidationType struct {
	Name            types.String `tfsdk:"name"`
	Value           types.String `tfsdk:"value"`
	Fixup           types.String `tfsdk:"fixup"`
	MapValue        types.String `tfsdk:"map_value"`
	Xss             types.Bool   `tfsdk:"xss"`
	XssPatternsFile types.String `tfsdk:"xss_patterns_file"`
}

var DmValidationTypeObjectType = map[string]attr.Type{
	"name":              types.StringType,
	"value":             types.StringType,
	"fixup":             types.StringType,
	"map_value":         types.StringType,
	"xss":               types.BoolType,
	"xss_patterns_file": types.StringType,
}
var DmValidationTypeObjectDefault = map[string]attr.Value{
	"name":              types.StringNull(),
	"value":             types.StringNull(),
	"fixup":             types.StringValue("error"),
	"map_value":         types.StringNull(),
	"xss":               types.BoolValue(false),
	"xss_patterns_file": types.StringValue("store:///XSS-Patterns.xml"),
}
var DmValidationTypeDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Name Expression", "", "").String,
			Computed:            true,
		},
		"value": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Value Constraint", "", "").String,
			Computed:            true,
		},
		"fixup": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Failure Policy", "", "").AddStringEnum("passthrough", "strip", "error", "set").AddDefaultValue("error").String,
			Computed:            true,
		},
		"map_value": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Map Value", "", "").String,
			Computed:            true,
		},
		"xss": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Check XSS", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"xss_patterns_file": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("XSS (Cross Site Scripting) Protection Patterns File", "", "").AddDefaultValue("store:///XSS-Patterns.xml").String,
			Computed:            true,
		},
	},
}
var DmValidationTypeResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Name Expression", "", "").String,
			Required:            true,
		},
		"value": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Value Constraint", "", "").String,
			Required:            true,
		},
		"fixup": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Failure Policy", "", "").AddStringEnum("passthrough", "strip", "error", "set").AddDefaultValue("error").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("passthrough", "strip", "error", "set"),
			},
			Default: stringdefault.StaticString("error"),
		},
		"map_value": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Map Value", "", "").String,
			Optional:            true,
		},
		"xss": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Check XSS", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"xss_patterns_file": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("XSS (Cross Site Scripting) Protection Patterns File", "", "").AddDefaultValue("store:///XSS-Patterns.xml").String,
			Computed:            true,
			Optional:            true,
			Default:             stringdefault.StaticString("store:///XSS-Patterns.xml"),
		},
	},
}

func (data DmValidationType) IsNull() bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.Value.IsNull() {
		return false
	}
	if !data.Fixup.IsNull() {
		return false
	}
	if !data.MapValue.IsNull() {
		return false
	}
	if !data.Xss.IsNull() {
		return false
	}
	if !data.XssPatternsFile.IsNull() {
		return false
	}
	return true
}

func (data DmValidationType) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Name`, data.Name.ValueString())
	}
	if !data.Value.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Value`, data.Value.ValueString())
	}
	if !data.Fixup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Fixup`, data.Fixup.ValueString())
	}
	if !data.MapValue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MapValue`, data.MapValue.ValueString())
	}
	if !data.Xss.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XSS`, tfutils.StringFromBool(data.Xss, ""))
	}
	if !data.XssPatternsFile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XSSPatternsFile`, data.XssPatternsFile.ValueString())
	}
	return body
}

func (data *DmValidationType) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Value`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Value = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Value = types.StringNull()
	}
	if value := res.Get(pathRoot + `Fixup`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Fixup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Fixup = types.StringValue("error")
	}
	if value := res.Get(pathRoot + `MapValue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MapValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MapValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `XSS`); value.Exists() {
		data.Xss = tfutils.BoolFromString(value.String())
	} else {
		data.Xss = types.BoolNull()
	}
	if value := res.Get(pathRoot + `XSSPatternsFile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.XssPatternsFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XssPatternsFile = types.StringValue("store:///XSS-Patterns.xml")
	}
}

func (data *DmValidationType) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && !data.Name.IsNull() {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Value`); value.Exists() && !data.Value.IsNull() {
		data.Value = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Value = types.StringNull()
	}
	if value := res.Get(pathRoot + `Fixup`); value.Exists() && !data.Fixup.IsNull() {
		data.Fixup = tfutils.ParseStringFromGJSON(value)
	} else if data.Fixup.ValueString() != "error" {
		data.Fixup = types.StringNull()
	}
	if value := res.Get(pathRoot + `MapValue`); value.Exists() && !data.MapValue.IsNull() {
		data.MapValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MapValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `XSS`); value.Exists() && !data.Xss.IsNull() {
		data.Xss = tfutils.BoolFromString(value.String())
	} else if data.Xss.ValueBool() {
		data.Xss = types.BoolNull()
	}
	if value := res.Get(pathRoot + `XSSPatternsFile`); value.Exists() && !data.XssPatternsFile.IsNull() {
		data.XssPatternsFile = tfutils.ParseStringFromGJSON(value)
	} else if data.XssPatternsFile.ValueString() != "store:///XSS-Patterns.xml" {
		data.XssPatternsFile = types.StringNull()
	}
}
