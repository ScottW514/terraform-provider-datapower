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

type DmAssemblyLogicExecute struct {
	Condition types.String `tfsdk:"condition"`
	Execute   types.String `tfsdk:"execute"`
}

var DmAssemblyLogicExecuteObjectType = map[string]attr.Type{
	"condition": types.StringType,
	"execute":   types.StringType,
}
var DmAssemblyLogicExecuteObjectDefault = map[string]attr.Value{
	"condition": types.StringNull(),
	"execute":   types.StringNull(),
}

func GetDmAssemblyLogicExecuteDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmAssemblyLogicExecuteDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"condition": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the JSONata expression to match against the input.",
				Computed:            true,
			},
			"execute": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the name of the API rule to run when the JSONata expression evaluates to true. The specified API rule can define further switch assembly actions.",
				Computed:            true,
			},
		},
	}
	return DmAssemblyLogicExecuteDataSourceSchema
}
func GetDmAssemblyLogicExecuteResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmAssemblyLogicExecuteResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"condition": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the JSONata expression to match against the input.", "", "").String,
				Required:            true,
			},
			"execute": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the API rule to run when the JSONata expression evaluates to true. The specified API rule can define further switch assembly actions.", "", "").String,
				Required:            true,
			},
		},
	}
	return DmAssemblyLogicExecuteResourceSchema
}

func (data DmAssemblyLogicExecute) IsNull() bool {
	if !data.Condition.IsNull() {
		return false
	}
	if !data.Execute.IsNull() {
		return false
	}
	return true
}

func (data DmAssemblyLogicExecute) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Condition.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Condition`, data.Condition.ValueString())
	}
	if !data.Execute.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Execute`, data.Execute.ValueString())
	}
	return body
}

func (data *DmAssemblyLogicExecute) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Condition`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Condition = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Condition = types.StringNull()
	}
	if value := res.Get(pathRoot + `Execute`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Execute = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Execute = types.StringNull()
	}
}

func (data *DmAssemblyLogicExecute) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Condition`); value.Exists() && !data.Condition.IsNull() {
		data.Condition = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Condition = types.StringNull()
	}
	if value := res.Get(pathRoot + `Execute`); value.Exists() && !data.Execute.IsNull() {
		data.Execute = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Execute = types.StringNull()
	}
}
