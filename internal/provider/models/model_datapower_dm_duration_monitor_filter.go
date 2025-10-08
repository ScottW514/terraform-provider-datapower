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

type DmDurationMonitorFilter struct {
	Name   types.String `tfsdk:"name"`
	Value  types.Int64  `tfsdk:"value"`
	Action types.String `tfsdk:"action"`
}

var DmDurationMonitorFilterObjectType = map[string]attr.Type{
	"name":   types.StringType,
	"value":  types.Int64Type,
	"action": types.StringType,
}
var DmDurationMonitorFilterObjectDefault = map[string]attr.Value{
	"name":   types.StringNull(),
	"value":  types.Int64Null(),
	"action": types.StringNull(),
}

func GetDmDurationMonitorFilterDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmDurationMonitorFilterDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"name": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Enter the name of the threshold. This name appears in the logs when the threshold action is taken.",
				Computed:            true,
			},
			"value": DataSourceSchema.Int64Attribute{
				MarkdownDescription: "Enter the threshold value in milliseconds.",
				Computed:            true,
			},
			"action": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Select the monitor policy (Message Filter Action) implemented by the monitor when the target message type exceeds the threshold value. Click the + button to create a new action.",
				Computed:            true,
			},
		},
	}
	return DmDurationMonitorFilterDataSourceSchema
}
func GetDmDurationMonitorFilterResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmDurationMonitorFilterResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"name": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of the threshold. This name appears in the logs when the threshold action is taken.", "", "").String,
				Required:            true,
			},
			"value": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the threshold value in milliseconds.", "", "").String,
				Required:            true,
			},
			"action": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the monitor policy (Message Filter Action) implemented by the monitor when the target message type exceeds the threshold value. Click the + button to create a new action.", "", "filter_action").String,
				Required:            true,
			},
		},
	}
	return DmDurationMonitorFilterResourceSchema
}

func (data DmDurationMonitorFilter) IsNull() bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.Value.IsNull() {
		return false
	}
	if !data.Action.IsNull() {
		return false
	}
	return true
}

func (data DmDurationMonitorFilter) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Name`, data.Name.ValueString())
	}
	if !data.Value.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Value`, data.Value.ValueInt64())
	}
	if !data.Action.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Action`, data.Action.ValueString())
	}
	return body
}

func (data *DmDurationMonitorFilter) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Value`); value.Exists() {
		data.Value = types.Int64Value(value.Int())
	} else {
		data.Value = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Action`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Action = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Action = types.StringNull()
	}
}

func (data *DmDurationMonitorFilter) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && !data.Name.IsNull() {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Value`); value.Exists() && !data.Value.IsNull() {
		data.Value = types.Int64Value(value.Int())
	} else {
		data.Value = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Action`); value.Exists() && !data.Action.IsNull() {
		data.Action = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Action = types.StringNull()
	}
}
