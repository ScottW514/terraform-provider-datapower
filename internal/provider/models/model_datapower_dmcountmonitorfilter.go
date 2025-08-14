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

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	DataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	ResourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmCountMonitorFilter struct {
	Name       types.String `tfsdk:"name"`
	Interval   types.Int64  `tfsdk:"interval"`
	RateLimit  types.Int64  `tfsdk:"rate_limit"`
	BurstLimit types.Int64  `tfsdk:"burst_limit"`
	Action     types.String `tfsdk:"action"`
}

var DmCountMonitorFilterObjectType = map[string]attr.Type{
	"name":        types.StringType,
	"interval":    types.Int64Type,
	"rate_limit":  types.Int64Type,
	"burst_limit": types.Int64Type,
	"action":      types.StringType,
}
var DmCountMonitorFilterObjectDefault = map[string]attr.Value{
	"name":        types.StringNull(),
	"interval":    types.Int64Null(),
	"rate_limit":  types.Int64Null(),
	"burst_limit": types.Int64Null(),
	"action":      types.StringNull(),
}
var DmCountMonitorFilterDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of the threshold. This name appears in the log files when the threshold is reached.", "", "").String,
			Computed:            true,
		},
		"interval": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enter the measurement interval in milliseconds. One second is expressed as 1000 here.", "", "").String,
			Computed:            true,
		},
		"rate_limit": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the threshold value as a number of messages. A monitored message type that exceeds this value over the measurement interval specified by the Interval property is subject to the policy identified by the Action values list.", "", "").AddIntegerRange(1, 4294967295).String,
			Computed:            true,
		},
		"burst_limit": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum allowed message burst. Use a value approximately twice the value of Rate Limit.", "", "").String,
			Computed:            true,
		},
		"action": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Select the monitor policy implemented by the monitor when the target message type exceeds the threshold value. This is a list of Message Filter Actions. Click the + button to create a new action.", "", "filteraction").String,
			Computed:            true,
		},
	},
}
var DmCountMonitorFilterResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of the threshold. This name appears in the log files when the threshold is reached.", "", "").String,
			Required:            true,
		},
		"interval": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enter the measurement interval in milliseconds. One second is expressed as 1000 here.", "", "").String,
			Required:            true,
		},
		"rate_limit": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the threshold value as a number of messages. A monitored message type that exceeds this value over the measurement interval specified by the Interval property is subject to the policy identified by the Action values list.", "", "").AddIntegerRange(1, 4294967295).String,
			Required:            true,
			Validators: []validator.Int64{
				int64validator.Between(1, 4294967295),
			},
		},
		"burst_limit": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum allowed message burst. Use a value approximately twice the value of Rate Limit.", "", "").String,
			Required:            true,
		},
		"action": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Select the monitor policy implemented by the monitor when the target message type exceeds the threshold value. This is a list of Message Filter Actions. Click the + button to create a new action.", "", "filteraction").String,
			Required:            true,
		},
	},
}

func (data DmCountMonitorFilter) IsNull() bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.Interval.IsNull() {
		return false
	}
	if !data.RateLimit.IsNull() {
		return false
	}
	if !data.BurstLimit.IsNull() {
		return false
	}
	if !data.Action.IsNull() {
		return false
	}
	return true
}

func (data DmCountMonitorFilter) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Name`, data.Name.ValueString())
	}
	if !data.Interval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Interval`, data.Interval.ValueInt64())
	}
	if !data.RateLimit.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RateLimit`, data.RateLimit.ValueInt64())
	}
	if !data.BurstLimit.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BurstLimit`, data.BurstLimit.ValueInt64())
	}
	if !data.Action.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Action`, data.Action.ValueString())
	}
	return body
}

func (data *DmCountMonitorFilter) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Interval`); value.Exists() {
		data.Interval = types.Int64Value(value.Int())
	} else {
		data.Interval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `RateLimit`); value.Exists() {
		data.RateLimit = types.Int64Value(value.Int())
	} else {
		data.RateLimit = types.Int64Null()
	}
	if value := res.Get(pathRoot + `BurstLimit`); value.Exists() {
		data.BurstLimit = types.Int64Value(value.Int())
	} else {
		data.BurstLimit = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Action`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Action = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Action = types.StringNull()
	}
}

func (data *DmCountMonitorFilter) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && !data.Name.IsNull() {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Interval`); value.Exists() && !data.Interval.IsNull() {
		data.Interval = types.Int64Value(value.Int())
	} else {
		data.Interval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `RateLimit`); value.Exists() && !data.RateLimit.IsNull() {
		data.RateLimit = types.Int64Value(value.Int())
	} else {
		data.RateLimit = types.Int64Null()
	}
	if value := res.Get(pathRoot + `BurstLimit`); value.Exists() && !data.BurstLimit.IsNull() {
		data.BurstLimit = types.Int64Value(value.Int())
	} else {
		data.BurstLimit = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Action`); value.Exists() && !data.Action.IsNull() {
		data.Action = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Action = types.StringNull()
	}
}
