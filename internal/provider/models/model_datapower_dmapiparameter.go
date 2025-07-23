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

type DmAPIParameter struct {
	Name           types.String `tfsdk:"name"`
	Required       types.Bool   `tfsdk:"required"`
	Type           types.String `tfsdk:"type"`
	Where          types.String `tfsdk:"where"`
	SchemaOrFormat types.String `tfsdk:"schema_or_format"`
	Maximum        types.String `tfsdk:"maximum"`
	Minimum        types.String `tfsdk:"minimum"`
	Pattern        types.String `tfsdk:"pattern"`
	Enum           types.String `tfsdk:"enum"`
}

var DmAPIParameterObjectType = map[string]attr.Type{
	"name":             types.StringType,
	"required":         types.BoolType,
	"type":             types.StringType,
	"where":            types.StringType,
	"schema_or_format": types.StringType,
	"maximum":          types.StringType,
	"minimum":          types.StringType,
	"pattern":          types.StringType,
	"enum":             types.StringType,
}
var DmAPIParameterObjectDefault = map[string]attr.Value{
	"name":             types.StringNull(),
	"required":         types.BoolValue(false),
	"type":             types.StringValue("string"),
	"where":            types.StringValue("path"),
	"schema_or_format": types.StringNull(),
	"maximum":          types.StringNull(),
	"minimum":          types.StringNull(),
	"pattern":          types.StringNull(),
	"enum":             types.StringNull(),
}
var DmAPIParameterDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Name", "", "").String,
			Computed:            true,
		},
		"required": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Required", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Type", "", "").AddStringEnum("integer", "long", "float", "double", "string", "byte", "binary", "boolean", "date", "dateTime", "password", "array", "object", "file", "number").AddDefaultValue("string").String,
			Computed:            true,
		},
		"where": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Where", "", "").AddStringEnum("path", "query", "body", "formdata", "header").AddDefaultValue("path").String,
			Computed:            true,
		},
		"schema_or_format": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Format", "", "").String,
			Computed:            true,
		},
		"maximum": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Maximum", "", "").String,
			Computed:            true,
		},
		"minimum": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Minimum", "", "").String,
			Computed:            true,
		},
		"pattern": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Pattern", "", "").String,
			Computed:            true,
		},
		"enum": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enum", "", "").String,
			Computed:            true,
		},
	},
}
var DmAPIParameterResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Name", "", "").String,
			Required:            true,
		},
		"required": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Required", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Type", "", "").AddStringEnum("integer", "long", "float", "double", "string", "byte", "binary", "boolean", "date", "dateTime", "password", "array", "object", "file", "number").AddDefaultValue("string").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("integer", "long", "float", "double", "string", "byte", "binary", "boolean", "date", "dateTime", "password", "array", "object", "file", "number"),
			},
			Default: stringdefault.StaticString("string"),
		},
		"where": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Where", "", "").AddStringEnum("path", "query", "body", "formdata", "header").AddDefaultValue("path").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("path", "query", "body", "formdata", "header"),
			},
			Default: stringdefault.StaticString("path"),
		},
		"schema_or_format": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Format", "", "").String,
			Optional:            true,
		},
		"maximum": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Maximum", "", "").String,
			Optional:            true,
		},
		"minimum": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Minimum", "", "").String,
			Optional:            true,
		},
		"pattern": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Pattern", "", "").String,
			Optional:            true,
		},
		"enum": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enum", "", "").String,
			Optional:            true,
		},
	},
}

func (data DmAPIParameter) IsNull() bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.Required.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
		return false
	}
	if !data.Where.IsNull() {
		return false
	}
	if !data.SchemaOrFormat.IsNull() {
		return false
	}
	if !data.Maximum.IsNull() {
		return false
	}
	if !data.Minimum.IsNull() {
		return false
	}
	if !data.Pattern.IsNull() {
		return false
	}
	if !data.Enum.IsNull() {
		return false
	}
	return true
}

func (data DmAPIParameter) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Name`, data.Name.ValueString())
	}
	if !data.Required.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Required`, tfutils.StringFromBool(data.Required, ""))
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Type`, data.Type.ValueString())
	}
	if !data.Where.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Where`, data.Where.ValueString())
	}
	if !data.SchemaOrFormat.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SchemaOrFormat`, data.SchemaOrFormat.ValueString())
	}
	if !data.Maximum.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Maximum`, data.Maximum.ValueString())
	}
	if !data.Minimum.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Minimum`, data.Minimum.ValueString())
	}
	if !data.Pattern.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Pattern`, data.Pattern.ValueString())
	}
	if !data.Enum.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Enum`, data.Enum.ValueString())
	}
	return body
}

func (data *DmAPIParameter) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Required`); value.Exists() {
		data.Required = tfutils.BoolFromString(value.String())
	} else {
		data.Required = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringValue("string")
	}
	if value := res.Get(pathRoot + `Where`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Where = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Where = types.StringValue("path")
	}
	if value := res.Get(pathRoot + `SchemaOrFormat`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SchemaOrFormat = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SchemaOrFormat = types.StringNull()
	}
	if value := res.Get(pathRoot + `Maximum`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Maximum = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Maximum = types.StringNull()
	}
	if value := res.Get(pathRoot + `Minimum`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Minimum = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Minimum = types.StringNull()
	}
	if value := res.Get(pathRoot + `Pattern`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Pattern = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Pattern = types.StringNull()
	}
	if value := res.Get(pathRoot + `Enum`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Enum = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Enum = types.StringNull()
	}
}

func (data *DmAPIParameter) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && !data.Name.IsNull() {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Required`); value.Exists() && !data.Required.IsNull() {
		data.Required = tfutils.BoolFromString(value.String())
	} else if data.Required.ValueBool() {
		data.Required = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && !data.Type.IsNull() {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else if data.Type.ValueString() != "string" {
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `Where`); value.Exists() && !data.Where.IsNull() {
		data.Where = tfutils.ParseStringFromGJSON(value)
	} else if data.Where.ValueString() != "path" {
		data.Where = types.StringNull()
	}
	if value := res.Get(pathRoot + `SchemaOrFormat`); value.Exists() && !data.SchemaOrFormat.IsNull() {
		data.SchemaOrFormat = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SchemaOrFormat = types.StringNull()
	}
	if value := res.Get(pathRoot + `Maximum`); value.Exists() && !data.Maximum.IsNull() {
		data.Maximum = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Maximum = types.StringNull()
	}
	if value := res.Get(pathRoot + `Minimum`); value.Exists() && !data.Minimum.IsNull() {
		data.Minimum = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Minimum = types.StringNull()
	}
	if value := res.Get(pathRoot + `Pattern`); value.Exists() && !data.Pattern.IsNull() {
		data.Pattern = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Pattern = types.StringNull()
	}
	if value := res.Get(pathRoot + `Enum`); value.Exists() && !data.Enum.IsNull() {
		data.Enum = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Enum = types.StringNull()
	}
}
