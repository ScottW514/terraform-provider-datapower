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

type DmTargetMapRule struct {
	Target  types.String `tfsdk:"target"`
	Execute types.String `tfsdk:"execute"`
}

var DmTargetMapRuleObjectType = map[string]attr.Type{
	"target":  types.StringType,
	"execute": types.StringType,
}
var DmTargetMapRuleObjectDefault = map[string]attr.Value{
	"target":  types.StringNull(),
	"execute": types.StringNull(),
}

func GetDmTargetMapRuleDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmTargetMapRuleDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"target": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the target name that is configured in the schema.",
				Computed:            true,
			},
			"execute": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the function call assembly action to call.",
				Computed:            true,
			},
		},
	}
	return DmTargetMapRuleDataSourceSchema
}
func GetDmTargetMapRuleResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmTargetMapRuleResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"target": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the target name that is configured in the schema.", "target", "").String,
				Required:            true,
			},
			"execute": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the function call assembly action to call.", "execute", "assembly_action_function_call").String,
				Required:            true,
			},
		},
	}
	return DmTargetMapRuleResourceSchema
}

func (data DmTargetMapRule) IsNull() bool {
	if !data.Target.IsNull() {
		return false
	}
	if !data.Execute.IsNull() {
		return false
	}
	return true
}

func (data DmTargetMapRule) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Target.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Target`, data.Target.ValueString())
	}
	if !data.Execute.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Execute`, data.Execute.ValueString())
	}
	return body
}

func (data *DmTargetMapRule) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Target`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Target = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Target = types.StringNull()
	}
	if value := res.Get(pathRoot + `Execute`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Execute = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Execute = types.StringNull()
	}
}

func (data *DmTargetMapRule) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Target`); value.Exists() && !data.Target.IsNull() {
		data.Target = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Target = types.StringNull()
	}
	if value := res.Get(pathRoot + `Execute`); value.Exists() && !data.Execute.IsNull() {
		data.Execute = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Execute = types.StringNull()
	}
}
