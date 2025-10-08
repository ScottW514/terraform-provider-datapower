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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmExtensionFunction struct {
	ExtensionFunctionNamespace types.String `tfsdk:"extension_function_namespace"`
	ExtensionFunction          types.String `tfsdk:"extension_function"`
	LocalFunctionNamespace     types.String `tfsdk:"local_function_namespace"`
	LocalFunction              types.String `tfsdk:"local_function"`
}

var DmExtensionFunctionObjectType = map[string]attr.Type{
	"extension_function_namespace": types.StringType,
	"extension_function":           types.StringType,
	"local_function_namespace":     types.StringType,
	"local_function":               types.StringType,
}
var DmExtensionFunctionObjectDefault = map[string]attr.Value{
	"extension_function_namespace": types.StringNull(),
	"extension_function":           types.StringNull(),
	"local_function_namespace":     types.StringValue("http://www.datapower.com/extensions"),
	"local_function":               types.StringNull(),
}

func GetDmExtensionFunctionDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmExtensionFunctionDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"extension_function_namespace": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Enter namespace of the extension functions used in the custom stylesheets. For example, \"http://www.fredspace/extensions\".",
				Computed:            true,
			},
			"extension_function": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Enter the extension function within the namespace to map. For example, \"nodeset()\".",
				Computed:            true,
			},
			"local_function_namespace": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Enter the local namespace if it is other than the DataPower default (shown in the box).",
				Computed:            true,
			},
			"local_function": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Enter the local function to use in place of the extension function identified above. For example, \"node-set()\".",
				Computed:            true,
			},
		},
	}
	return DmExtensionFunctionDataSourceSchema
}
func GetDmExtensionFunctionResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmExtensionFunctionResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"extension_function_namespace": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter namespace of the extension functions used in the custom stylesheets. For example, \"http://www.fredspace/extensions\".", "", "").String,
				Required:            true,
			},
			"extension_function": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the extension function within the namespace to map. For example, \"nodeset()\".", "", "").String,
				Required:            true,
			},
			"local_function_namespace": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the local namespace if it is other than the DataPower default (shown in the box).", "", "").AddDefaultValue("http://www.datapower.com/extensions").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("http://www.datapower.com/extensions"),
			},
			"local_function": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the local function to use in place of the extension function identified above. For example, \"node-set()\".", "", "").String,
				Required:            true,
			},
		},
	}
	return DmExtensionFunctionResourceSchema
}

func (data DmExtensionFunction) IsNull() bool {
	if !data.ExtensionFunctionNamespace.IsNull() {
		return false
	}
	if !data.ExtensionFunction.IsNull() {
		return false
	}
	if !data.LocalFunctionNamespace.IsNull() {
		return false
	}
	if !data.LocalFunction.IsNull() {
		return false
	}
	return true
}

func (data DmExtensionFunction) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.ExtensionFunctionNamespace.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ExtensionFunctionNamespace`, data.ExtensionFunctionNamespace.ValueString())
	}
	if !data.ExtensionFunction.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ExtensionFunction`, data.ExtensionFunction.ValueString())
	}
	if !data.LocalFunctionNamespace.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalFunctionNamespace`, data.LocalFunctionNamespace.ValueString())
	}
	if !data.LocalFunction.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalFunction`, data.LocalFunction.ValueString())
	}
	return body
}

func (data *DmExtensionFunction) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `ExtensionFunctionNamespace`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ExtensionFunctionNamespace = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExtensionFunctionNamespace = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExtensionFunction`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ExtensionFunction = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExtensionFunction = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalFunctionNamespace`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalFunctionNamespace = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalFunctionNamespace = types.StringValue("http://www.datapower.com/extensions")
	}
	if value := res.Get(pathRoot + `LocalFunction`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalFunction = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalFunction = types.StringNull()
	}
}

func (data *DmExtensionFunction) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `ExtensionFunctionNamespace`); value.Exists() && !data.ExtensionFunctionNamespace.IsNull() {
		data.ExtensionFunctionNamespace = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExtensionFunctionNamespace = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExtensionFunction`); value.Exists() && !data.ExtensionFunction.IsNull() {
		data.ExtensionFunction = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExtensionFunction = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalFunctionNamespace`); value.Exists() && !data.LocalFunctionNamespace.IsNull() {
		data.LocalFunctionNamespace = tfutils.ParseStringFromGJSON(value)
	} else if data.LocalFunctionNamespace.ValueString() != "http://www.datapower.com/extensions" {
		data.LocalFunctionNamespace = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalFunction`); value.Exists() && !data.LocalFunction.IsNull() {
		data.LocalFunction = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalFunction = types.StringNull()
	}
}
