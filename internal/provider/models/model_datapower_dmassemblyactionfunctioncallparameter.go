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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmAssemblyActionFunctionCallParameter struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

var DmAssemblyActionFunctionCallParameterObjectType = map[string]attr.Type{
	"name":  types.StringType,
	"value": types.StringType,
}
var DmAssemblyActionFunctionCallParameterObjectDefault = map[string]attr.Value{
	"name":  types.StringNull(),
	"value": types.StringNull(),
}
var DmAssemblyActionFunctionCallParameterDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the parameter.", "name", "").String,
			Computed:            true,
		},
		"value": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the default parameter value as a string. The default value is used when no value is passed in the request. When you specify a JSON value, enter the value as a JSON string. This value is required in this assembly action if no default value is specified in the assembly function.", "value", "").String,
			Computed:            true,
		},
	},
}
var DmAssemblyActionFunctionCallParameterResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the parameter.", "name", "").String,
			Required:            true,
		},
		"value": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the default parameter value as a string. The default value is used when no value is passed in the request. When you specify a JSON value, enter the value as a JSON string. This value is required in this assembly action if no default value is specified in the assembly function.", "value", "").String,
			Optional:            true,
		},
	},
}

func (data DmAssemblyActionFunctionCallParameter) IsNull() bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.Value.IsNull() {
		return false
	}
	return true
}

func (data DmAssemblyActionFunctionCallParameter) ToBody(ctx context.Context, pathRoot string) string {
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
	return body
}

func (data *DmAssemblyActionFunctionCallParameter) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
}

func (data *DmAssemblyActionFunctionCallParameter) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
}
