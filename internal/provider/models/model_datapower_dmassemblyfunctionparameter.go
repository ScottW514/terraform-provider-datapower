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

type DmAssemblyFunctionParameter struct {
	Name        types.String `tfsdk:"name"`
	Label       types.String `tfsdk:"label"`
	Description types.String `tfsdk:"description"`
	Schema      types.String `tfsdk:"schema"`
	Value       types.String `tfsdk:"value"`
	ValueType   types.String `tfsdk:"value_type"`
	Required    types.Bool   `tfsdk:"required"`
}

var DmAssemblyFunctionParameterObjectType = map[string]attr.Type{
	"name":        types.StringType,
	"label":       types.StringType,
	"description": types.StringType,
	"schema":      types.StringType,
	"value":       types.StringType,
	"value_type":  types.StringType,
	"required":    types.BoolType,
}
var DmAssemblyFunctionParameterObjectDefault = map[string]attr.Value{
	"name":        types.StringNull(),
	"label":       types.StringNull(),
	"description": types.StringNull(),
	"schema":      types.StringNull(),
	"value":       types.StringNull(),
	"value_type":  types.StringValue("string"),
	"required":    types.BoolValue(true),
}
var DmAssemblyFunctionParameterDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the parameter.", "name", "").String,
			Computed:            true,
		},
		"label": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the label to explain the parameter to API developers.", "label", "").String,
			Computed:            true,
		},
		"description": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify a description of the parameter to advertise the parameter to API developers.", "description", "").String,
			Computed:            true,
		},
		"schema": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the API schema to verify the parameter type.", "schema", "apischema").String,
			Computed:            true,
		},
		"value": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the default parameter value to pass to the assembly function as a string. The default value is used when no value is passed in the request. The default parameter value is required if the assembly function is called in an assembly function call action and no parameter value is defined in the action. <ul><li>If the default value is a JSON payload, enter the value as a JSON string.</li><li>If the default value is an empty string, it is treated as not specified.</li><li>If the assembly function parameter specifies an API schema, the default parameter value overrides the default value defined in the schema.</li></ul>", "value", "").String,
			Computed:            true,
		},
		"value_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the type of the parameter value.", "value-type", "").AddStringEnum("string", "payload", "message").AddDefaultValue("string").String,
			Computed:            true,
		},
		"required": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether this parameter requires a value in an assembly function call.", "required", "").AddDefaultValue("true").String,
			Computed:            true,
		},
	},
}
var DmAssemblyFunctionParameterResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the parameter.", "name", "").String,
			Required:            true,
		},
		"label": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the label to explain the parameter to API developers.", "label", "").String,
			Optional:            true,
		},
		"description": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify a description of the parameter to advertise the parameter to API developers.", "description", "").String,
			Optional:            true,
		},
		"schema": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the API schema to verify the parameter type.", "schema", "apischema").String,
			Optional:            true,
		},
		"value": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the default parameter value to pass to the assembly function as a string. The default value is used when no value is passed in the request. The default parameter value is required if the assembly function is called in an assembly function call action and no parameter value is defined in the action. <ul><li>If the default value is a JSON payload, enter the value as a JSON string.</li><li>If the default value is an empty string, it is treated as not specified.</li><li>If the assembly function parameter specifies an API schema, the default parameter value overrides the default value defined in the schema.</li></ul>", "value", "").String,
			Optional:            true,
		},
		"value_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the type of the parameter value.", "value-type", "").AddStringEnum("string", "payload", "message").AddDefaultValue("string").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("string", "payload", "message"),
			},
			Default: stringdefault.StaticString("string"),
		},
		"required": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether this parameter requires a value in an assembly function call.", "required", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
	},
}

func (data DmAssemblyFunctionParameter) IsNull() bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.Label.IsNull() {
		return false
	}
	if !data.Description.IsNull() {
		return false
	}
	if !data.Schema.IsNull() {
		return false
	}
	if !data.Value.IsNull() {
		return false
	}
	if !data.ValueType.IsNull() {
		return false
	}
	if !data.Required.IsNull() {
		return false
	}
	return true
}

func (data DmAssemblyFunctionParameter) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Name`, data.Name.ValueString())
	}
	if !data.Label.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Label`, data.Label.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Description`, data.Description.ValueString())
	}
	if !data.Schema.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Schema`, data.Schema.ValueString())
	}
	if !data.Value.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Value`, data.Value.ValueString())
	}
	if !data.ValueType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ValueType`, data.ValueType.ValueString())
	}
	if !data.Required.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Required`, tfutils.StringFromBool(data.Required, ""))
	}
	return body
}

func (data *DmAssemblyFunctionParameter) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Label`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Label = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Label = types.StringNull()
	}
	if value := res.Get(pathRoot + `Description`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Description = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get(pathRoot + `Schema`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Schema = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Schema = types.StringNull()
	}
	if value := res.Get(pathRoot + `Value`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Value = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Value = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValueType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ValueType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ValueType = types.StringValue("string")
	}
	if value := res.Get(pathRoot + `Required`); value.Exists() {
		data.Required = tfutils.BoolFromString(value.String())
	} else {
		data.Required = types.BoolNull()
	}
}

func (data *DmAssemblyFunctionParameter) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && !data.Name.IsNull() {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Label`); value.Exists() && !data.Label.IsNull() {
		data.Label = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Label = types.StringNull()
	}
	if value := res.Get(pathRoot + `Description`); value.Exists() && !data.Description.IsNull() {
		data.Description = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get(pathRoot + `Schema`); value.Exists() && !data.Schema.IsNull() {
		data.Schema = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Schema = types.StringNull()
	}
	if value := res.Get(pathRoot + `Value`); value.Exists() && !data.Value.IsNull() {
		data.Value = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Value = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValueType`); value.Exists() && !data.ValueType.IsNull() {
		data.ValueType = tfutils.ParseStringFromGJSON(value)
	} else if data.ValueType.ValueString() != "string" {
		data.ValueType = types.StringNull()
	}
	if value := res.Get(pathRoot + `Required`); value.Exists() && !data.Required.IsNull() {
		data.Required = tfutils.BoolFromString(value.String())
	} else if !data.Required.ValueBool() {
		data.Required = types.BoolNull()
	}
}
