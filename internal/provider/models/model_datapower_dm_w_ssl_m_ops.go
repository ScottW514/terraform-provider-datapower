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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmWSSLMOps struct {
	Operation types.String `tfsdk:"operation"`
	Target    types.String `tfsdk:"target"`
	Severity  types.String `tfsdk:"severity"`
	Threshold types.Int64  `tfsdk:"threshold"`
	Action    types.String `tfsdk:"action"`
}

var DmWSSLMOpsObjectType = map[string]attr.Type{
	"operation": types.StringType,
	"target":    types.StringType,
	"severity":  types.StringType,
	"threshold": types.Int64Type,
	"action":    types.StringType,
}
var DmWSSLMOpsObjectDefault = map[string]attr.Value{
	"operation": types.StringValue("all"),
	"target":    types.StringNull(),
	"severity":  types.StringNull(),
	"threshold": types.Int64Null(),
	"action":    types.StringNull(),
}
var DmWSSLMOpsDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"operation": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the operation to monitor. The operation is defined in the WSDL file. The current implementation is to monitor all operations in the WSDL file.", "operation", "").AddStringEnum("all").AddDefaultValue("all").String,
			Computed:            true,
		},
		"target": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the target activity to monitor. Define the operation for each monitored activity.", "target", "").AddStringEnum("front", "rate").String,
			Computed:            true,
		},
		"severity": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the action threshold with its value. For example, you can define low and high thresholds as transactions rates increase. If the low threshold is 100 transactions/second and that limit is reached, some action is taken. Then, if the high threshold is 300 transactions/second and that limit is reached, another action is taken.", "severity", "").AddStringEnum("low", "high").String,
			Computed:            true,
		},
		"threshold": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the threshold value in TPS to trigger the action.", "threshold", "").String,
			Computed:            true,
		},
		"action": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the action to take when the threshold is reached.", "action", "").AddStringEnum("log", "throttle").String,
			Computed:            true,
		},
	},
}
var DmWSSLMOpsResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"operation": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the operation to monitor. The operation is defined in the WSDL file. The current implementation is to monitor all operations in the WSDL file.", "operation", "").AddStringEnum("all").AddDefaultValue("all").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("all"),
			},
			Default: stringdefault.StaticString("all"),
		},
		"target": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the target activity to monitor. Define the operation for each monitored activity.", "target", "").AddStringEnum("front", "rate").String,
			Required:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("front", "rate"),
			},
		},
		"severity": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the action threshold with its value. For example, you can define low and high thresholds as transactions rates increase. If the low threshold is 100 transactions/second and that limit is reached, some action is taken. Then, if the high threshold is 300 transactions/second and that limit is reached, another action is taken.", "severity", "").AddStringEnum("low", "high").String,
			Required:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("low", "high"),
			},
		},
		"threshold": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the threshold value in TPS to trigger the action.", "threshold", "").String,
			Optional:            true,
		},
		"action": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the action to take when the threshold is reached.", "action", "").AddStringEnum("log", "throttle").String,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("log", "throttle"),
			},
		},
	},
}

func (data DmWSSLMOps) IsNull() bool {
	if !data.Operation.IsNull() {
		return false
	}
	if !data.Target.IsNull() {
		return false
	}
	if !data.Severity.IsNull() {
		return false
	}
	if !data.Threshold.IsNull() {
		return false
	}
	if !data.Action.IsNull() {
		return false
	}
	return true
}

func (data DmWSSLMOps) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Operation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Operation`, data.Operation.ValueString())
	}
	if !data.Target.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Target`, data.Target.ValueString())
	}
	if !data.Severity.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Severity`, data.Severity.ValueString())
	}
	if !data.Threshold.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Threshold`, data.Threshold.ValueInt64())
	}
	if !data.Action.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Action`, data.Action.ValueString())
	}
	return body
}

func (data *DmWSSLMOps) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Operation`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Operation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Operation = types.StringValue("all")
	}
	if value := res.Get(pathRoot + `Target`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Target = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Target = types.StringNull()
	}
	if value := res.Get(pathRoot + `Severity`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Severity = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Severity = types.StringNull()
	}
	if value := res.Get(pathRoot + `Threshold`); value.Exists() {
		data.Threshold = types.Int64Value(value.Int())
	} else {
		data.Threshold = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Action`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Action = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Action = types.StringNull()
	}
}

func (data *DmWSSLMOps) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Operation`); value.Exists() && !data.Operation.IsNull() {
		data.Operation = tfutils.ParseStringFromGJSON(value)
	} else if data.Operation.ValueString() != "all" {
		data.Operation = types.StringNull()
	}
	if value := res.Get(pathRoot + `Target`); value.Exists() && !data.Target.IsNull() {
		data.Target = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Target = types.StringNull()
	}
	if value := res.Get(pathRoot + `Severity`); value.Exists() && !data.Severity.IsNull() {
		data.Severity = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Severity = types.StringNull()
	}
	if value := res.Get(pathRoot + `Threshold`); value.Exists() && !data.Threshold.IsNull() {
		data.Threshold = types.Int64Value(value.Int())
	} else {
		data.Threshold = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Action`); value.Exists() && !data.Action.IsNull() {
		data.Action = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Action = types.StringNull()
	}
}
