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

type DmCondition struct {
	Expression      types.String `tfsdk:"expression"`
	ConditionAction types.String `tfsdk:"condition_action"`
}

var DmConditionObjectType = map[string]attr.Type{
	"expression":       types.StringType,
	"condition_action": types.StringType,
}
var DmConditionObjectDefault = map[string]attr.Value{
	"expression":       types.StringNull(),
	"condition_action": types.StringNull(),
}
var DmConditionDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"expression": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("XPath expression", "", "").String,
			Computed:            true,
		},
		"condition_action": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Action", "", "").String,
			Computed:            true,
		},
	},
}
var DmConditionResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"expression": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("XPath expression", "", "").String,
			Required:            true,
		},
		"condition_action": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Action", "", "").String,
			Required:            true,
		},
	},
}

func (data DmCondition) IsNull() bool {
	if !data.Expression.IsNull() {
		return false
	}
	if !data.ConditionAction.IsNull() {
		return false
	}
	return true
}

func (data DmCondition) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Expression.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Expression`, data.Expression.ValueString())
	}
	if !data.ConditionAction.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ConditionAction`, data.ConditionAction.ValueString())
	}
	return body
}

func (data *DmCondition) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Expression`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Expression = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Expression = types.StringNull()
	}
	if value := res.Get(pathRoot + `ConditionAction`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ConditionAction = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ConditionAction = types.StringNull()
	}
}

func (data *DmCondition) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Expression`); value.Exists() && !data.Expression.IsNull() {
		data.Expression = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Expression = types.StringNull()
	}
	if value := res.Get(pathRoot + `ConditionAction`); value.Exists() && !data.ConditionAction.IsNull() {
		data.ConditionAction = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ConditionAction = types.StringNull()
	}
}
