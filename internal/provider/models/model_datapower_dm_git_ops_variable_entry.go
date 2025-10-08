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

type DmGitOpsVariableEntry struct {
	VariableName  types.String `tfsdk:"variable_name"`
	VariableValue types.String `tfsdk:"variable_value"`
}

var DmGitOpsVariableEntryObjectType = map[string]attr.Type{
	"variable_name":  types.StringType,
	"variable_value": types.StringType,
}
var DmGitOpsVariableEntryObjectDefault = map[string]attr.Value{
	"variable_name":  types.StringNull(),
	"variable_value": types.StringNull(),
}

func GetDmGitOpsVariableEntryDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmGitOpsVariableEntryDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"variable_name": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Variable name",
				Computed:            true,
			},
			"variable_value": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Variable value",
				Computed:            true,
			},
		},
	}
	return DmGitOpsVariableEntryDataSourceSchema
}
func GetDmGitOpsVariableEntryResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmGitOpsVariableEntryResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"variable_name": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Variable name", "name", "").String,
				Required:            true,
			},
			"variable_value": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Variable value", "value", "").String,
				Required:            true,
			},
		},
	}
	return DmGitOpsVariableEntryResourceSchema
}

func (data DmGitOpsVariableEntry) IsNull() bool {
	if !data.VariableName.IsNull() {
		return false
	}
	if !data.VariableValue.IsNull() {
		return false
	}
	return true
}

func (data DmGitOpsVariableEntry) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.VariableName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`VariableName`, data.VariableName.ValueString())
	}
	if !data.VariableValue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`VariableValue`, data.VariableValue.ValueString())
	}
	return body
}

func (data *DmGitOpsVariableEntry) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `VariableName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.VariableName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.VariableName = types.StringNull()
	}
	if value := res.Get(pathRoot + `VariableValue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.VariableValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.VariableValue = types.StringNull()
	}
}

func (data *DmGitOpsVariableEntry) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `VariableName`); value.Exists() && !data.VariableName.IsNull() {
		data.VariableName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.VariableName = types.StringNull()
	}
	if value := res.Get(pathRoot + `VariableValue`); value.Exists() && !data.VariableValue.IsNull() {
		data.VariableValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.VariableValue = types.StringNull()
	}
}
