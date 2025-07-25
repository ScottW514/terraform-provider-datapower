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
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	DataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	ResourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmSLMStatement struct {
	SlmId                                      types.Int64  `tfsdk:"slm_id"`
	UserString                                 types.String `tfsdk:"user_string"`
	CredClass                                  types.String `tfsdk:"cred_class"`
	RsrcClass                                  types.String `tfsdk:"rsrc_class"`
	Schedule                                   types.String `tfsdk:"schedule"`
	Action                                     types.String `tfsdk:"action"`
	ThreshIntervalLength                       types.Int64  `tfsdk:"thresh_interval_length"`
	ThreshIntervalType                         types.String `tfsdk:"thresh_interval_type"`
	ThreshAlgorithm                            types.String `tfsdk:"thresh_algorithm"`
	ThresholdType                              types.String `tfsdk:"threshold_type"`
	ThresholdLevel                             types.Int64  `tfsdk:"threshold_level"`
	ReleaseThresholdLevel                      types.Int64  `tfsdk:"release_threshold_level"`
	BurstLimit                                 types.Int64  `tfsdk:"burst_limit"`
	ReportingAggregationInterval               types.Int64  `tfsdk:"reporting_aggregation_interval"`
	MaximumTotalReportingRecords               types.Int64  `tfsdk:"maximum_total_reporting_records"`
	AutoGeneratedByWebGui                      types.Bool   `tfsdk:"auto_generated_by_web_gui"`
	MaximumResourcesAndCredentialsForThreshold types.Int64  `tfsdk:"maximum_resources_and_credentials_for_threshold"`
}

var DmSLMStatementObjectType = map[string]attr.Type{
	"slm_id":                          types.Int64Type,
	"user_string":                     types.StringType,
	"cred_class":                      types.StringType,
	"rsrc_class":                      types.StringType,
	"schedule":                        types.StringType,
	"action":                          types.StringType,
	"thresh_interval_length":          types.Int64Type,
	"thresh_interval_type":            types.StringType,
	"thresh_algorithm":                types.StringType,
	"threshold_type":                  types.StringType,
	"threshold_level":                 types.Int64Type,
	"release_threshold_level":         types.Int64Type,
	"burst_limit":                     types.Int64Type,
	"reporting_aggregation_interval":  types.Int64Type,
	"maximum_total_reporting_records": types.Int64Type,
	"auto_generated_by_web_gui":       types.BoolType,
	"maximum_resources_and_credentials_for_threshold": types.Int64Type,
}
var DmSLMStatementObjectDefault = map[string]attr.Value{
	"slm_id":                          types.Int64Null(),
	"user_string":                     types.StringNull(),
	"cred_class":                      types.StringNull(),
	"rsrc_class":                      types.StringNull(),
	"schedule":                        types.StringNull(),
	"action":                          types.StringNull(),
	"thresh_interval_length":          types.Int64Null(),
	"thresh_interval_type":            types.StringValue("fixed"),
	"thresh_algorithm":                types.StringValue("greater-than"),
	"threshold_type":                  types.StringValue("count-all"),
	"threshold_level":                 types.Int64Value(0),
	"release_threshold_level":         types.Int64Value(0),
	"burst_limit":                     types.Int64Null(),
	"reporting_aggregation_interval":  types.Int64Null(),
	"maximum_total_reporting_records": types.Int64Value(5000),
	"auto_generated_by_web_gui":       types.BoolValue(false),
	"maximum_resources_and_credentials_for_threshold": types.Int64Value(5000),
}
var DmSLMStatementDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"slm_id": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Identifier", "", "").String,
			Computed:            true,
		},
		"user_string": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("User annotation", "", "").String,
			Computed:            true,
		},
		"cred_class": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Credential class", "", "slmcredclass").String,
			Computed:            true,
		},
		"rsrc_class": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Resource class", "", "slmrsrcclass").String,
			Computed:            true,
		},
		"schedule": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Schedule", "", "slmschedule").String,
			Computed:            true,
		},
		"action": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Action", "", "slmaction").String,
			Computed:            true,
		},
		"thresh_interval_length": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Threshold interval length", "", "").String,
			Computed:            true,
		},
		"thresh_interval_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Threshold interval type", "", "").AddStringEnum("fixed", "moving", "concurrent").AddDefaultValue("fixed").String,
			Computed:            true,
		},
		"thresh_algorithm": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Threshold algorithm", "", "").AddStringEnum("greater-than", "less-than", "token-bucket", "high-low-thresholds").AddDefaultValue("greater-than").String,
			Computed:            true,
		},
		"threshold_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Threshold type", "", "").AddStringEnum("count-all", "count-errors", "latency-internal", "latency-backend", "latency-total", "payload-request", "payload-response", "payload-total").AddDefaultValue("count-all").String,
			Computed:            true,
		},
		"threshold_level": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Threshold level", "", "").AddIntegerRange(0, 9007199254740991).AddDefaultValue("0").String,
			Computed:            true,
		},
		"release_threshold_level": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("High-low release level", "", "").AddIntegerRange(0, 9007199254740991).AddDefaultValue("0").String,
			Computed:            true,
		},
		"burst_limit": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Burst limit", "", "").String,
			Computed:            true,
		},
		"reporting_aggregation_interval": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Reporting aggregation interval", "", "").String,
			Computed:            true,
		},
		"maximum_total_reporting_records": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Max records across intervals", "", "").AddIntegerRange(0, 4294967295).AddDefaultValue("5000").String,
			Computed:            true,
		},
		"auto_generated_by_web_gui": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Auto-generated (read-only)", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"maximum_resources_and_credentials_for_threshold": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Max credentials-resource combinations", "", "").AddIntegerRange(0, 4294967295).AddDefaultValue("5000").String,
			Computed:            true,
		},
	},
}
var DmSLMStatementResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"slm_id": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Identifier", "", "").String,
			Required:            true,
		},
		"user_string": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("User annotation", "", "").String,
			Optional:            true,
		},
		"cred_class": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Credential class", "", "slmcredclass").String,
			Optional:            true,
		},
		"rsrc_class": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Resource class", "", "slmrsrcclass").String,
			Optional:            true,
		},
		"schedule": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Schedule", "", "slmschedule").String,
			Optional:            true,
		},
		"action": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Action", "", "slmaction").String,
			Required:            true,
		},
		"thresh_interval_length": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Threshold interval length", "", "").String,
			Optional:            true,
		},
		"thresh_interval_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Threshold interval type", "", "").AddStringEnum("fixed", "moving", "concurrent").AddDefaultValue("fixed").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("fixed", "moving", "concurrent"),
			},
			Default: stringdefault.StaticString("fixed"),
		},
		"thresh_algorithm": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Threshold algorithm", "", "").AddStringEnum("greater-than", "less-than", "token-bucket", "high-low-thresholds").AddDefaultValue("greater-than").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("greater-than", "less-than", "token-bucket", "high-low-thresholds"),
			},
			Default: stringdefault.StaticString("greater-than"),
		},
		"threshold_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Threshold type", "", "").AddStringEnum("count-all", "count-errors", "latency-internal", "latency-backend", "latency-total", "payload-request", "payload-response", "payload-total").AddDefaultValue("count-all").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("count-all", "count-errors", "latency-internal", "latency-backend", "latency-total", "payload-request", "payload-response", "payload-total"),
			},
			Default: stringdefault.StaticString("count-all"),
		},
		"threshold_level": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Threshold level", "", "").AddIntegerRange(0, 9007199254740991).AddDefaultValue("0").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(0, 9007199254740991),
			},
			Default: int64default.StaticInt64(0),
		},
		"release_threshold_level": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("High-low release level", "", "").AddIntegerRange(0, 9007199254740991).AddDefaultValue("0").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(0, 9007199254740991),
			},
			Default: int64default.StaticInt64(0),
		},
		"burst_limit": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Burst limit", "", "").String,
			Optional:            true,
		},
		"reporting_aggregation_interval": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Reporting aggregation interval", "", "").String,
			Optional:            true,
		},
		"maximum_total_reporting_records": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Max records across intervals", "", "").AddIntegerRange(0, 4294967295).AddDefaultValue("5000").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(0, 4294967295),
			},
			Default: int64default.StaticInt64(5000),
		},
		"auto_generated_by_web_gui": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Auto-generated (read-only)", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"maximum_resources_and_credentials_for_threshold": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Max credentials-resource combinations", "", "").AddIntegerRange(0, 4294967295).AddDefaultValue("5000").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(0, 4294967295),
			},
			Default: int64default.StaticInt64(5000),
		},
	},
}

func (data DmSLMStatement) IsNull() bool {
	if !data.SlmId.IsNull() {
		return false
	}
	if !data.UserString.IsNull() {
		return false
	}
	if !data.CredClass.IsNull() {
		return false
	}
	if !data.RsrcClass.IsNull() {
		return false
	}
	if !data.Schedule.IsNull() {
		return false
	}
	if !data.Action.IsNull() {
		return false
	}
	if !data.ThreshIntervalLength.IsNull() {
		return false
	}
	if !data.ThreshIntervalType.IsNull() {
		return false
	}
	if !data.ThreshAlgorithm.IsNull() {
		return false
	}
	if !data.ThresholdType.IsNull() {
		return false
	}
	if !data.ThresholdLevel.IsNull() {
		return false
	}
	if !data.ReleaseThresholdLevel.IsNull() {
		return false
	}
	if !data.BurstLimit.IsNull() {
		return false
	}
	if !data.ReportingAggregationInterval.IsNull() {
		return false
	}
	if !data.MaximumTotalReportingRecords.IsNull() {
		return false
	}
	if !data.AutoGeneratedByWebGui.IsNull() {
		return false
	}
	if !data.MaximumResourcesAndCredentialsForThreshold.IsNull() {
		return false
	}
	return true
}

func (data DmSLMStatement) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.SlmId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SLMId`, data.SlmId.ValueInt64())
	}
	if !data.UserString.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserString`, data.UserString.ValueString())
	}
	if !data.CredClass.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CredClass`, data.CredClass.ValueString())
	}
	if !data.RsrcClass.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RsrcClass`, data.RsrcClass.ValueString())
	}
	if !data.Schedule.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Schedule`, data.Schedule.ValueString())
	}
	if !data.Action.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Action`, data.Action.ValueString())
	}
	if !data.ThreshIntervalLength.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ThreshIntervalLength`, data.ThreshIntervalLength.ValueInt64())
	}
	if !data.ThreshIntervalType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ThreshIntervalType`, data.ThreshIntervalType.ValueString())
	}
	if !data.ThreshAlgorithm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ThreshAlgorithm`, data.ThreshAlgorithm.ValueString())
	}
	if !data.ThresholdType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ThresholdType`, data.ThresholdType.ValueString())
	}
	if !data.ThresholdLevel.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ThresholdLevel`, data.ThresholdLevel.ValueInt64())
	}
	if !data.ReleaseThresholdLevel.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ReleaseThresholdLevel`, data.ReleaseThresholdLevel.ValueInt64())
	}
	if !data.BurstLimit.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BurstLimit`, data.BurstLimit.ValueInt64())
	}
	if !data.ReportingAggregationInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ReportingAggregationInterval`, data.ReportingAggregationInterval.ValueInt64())
	}
	if !data.MaximumTotalReportingRecords.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaximumTotalReportingRecords`, data.MaximumTotalReportingRecords.ValueInt64())
	}
	if !data.AutoGeneratedByWebGui.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AutoGeneratedByWebGUI`, tfutils.StringFromBool(data.AutoGeneratedByWebGui, ""))
	}
	if !data.MaximumResourcesAndCredentialsForThreshold.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaximumResourcesAndCredentialsForThreshold`, data.MaximumResourcesAndCredentialsForThreshold.ValueInt64())
	}
	return body
}

func (data *DmSLMStatement) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `SLMId`); value.Exists() {
		data.SlmId = types.Int64Value(value.Int())
	} else {
		data.SlmId = types.Int64Null()
	}
	if value := res.Get(pathRoot + `UserString`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserString = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserString = types.StringNull()
	}
	if value := res.Get(pathRoot + `CredClass`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CredClass = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CredClass = types.StringNull()
	}
	if value := res.Get(pathRoot + `RsrcClass`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RsrcClass = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RsrcClass = types.StringNull()
	}
	if value := res.Get(pathRoot + `Schedule`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Schedule = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Schedule = types.StringNull()
	}
	if value := res.Get(pathRoot + `Action`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Action = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Action = types.StringNull()
	}
	if value := res.Get(pathRoot + `ThreshIntervalLength`); value.Exists() {
		data.ThreshIntervalLength = types.Int64Value(value.Int())
	} else {
		data.ThreshIntervalLength = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ThreshIntervalType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ThreshIntervalType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ThreshIntervalType = types.StringValue("fixed")
	}
	if value := res.Get(pathRoot + `ThreshAlgorithm`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ThreshAlgorithm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ThreshAlgorithm = types.StringValue("greater-than")
	}
	if value := res.Get(pathRoot + `ThresholdType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ThresholdType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ThresholdType = types.StringValue("count-all")
	}
	if value := res.Get(pathRoot + `ThresholdLevel`); value.Exists() {
		data.ThresholdLevel = types.Int64Value(value.Int())
	} else {
		data.ThresholdLevel = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `ReleaseThresholdLevel`); value.Exists() {
		data.ReleaseThresholdLevel = types.Int64Value(value.Int())
	} else {
		data.ReleaseThresholdLevel = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `BurstLimit`); value.Exists() {
		data.BurstLimit = types.Int64Value(value.Int())
	} else {
		data.BurstLimit = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ReportingAggregationInterval`); value.Exists() {
		data.ReportingAggregationInterval = types.Int64Value(value.Int())
	} else {
		data.ReportingAggregationInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaximumTotalReportingRecords`); value.Exists() {
		data.MaximumTotalReportingRecords = types.Int64Value(value.Int())
	} else {
		data.MaximumTotalReportingRecords = types.Int64Value(5000)
	}
	if value := res.Get(pathRoot + `AutoGeneratedByWebGUI`); value.Exists() {
		data.AutoGeneratedByWebGui = tfutils.BoolFromString(value.String())
	} else {
		data.AutoGeneratedByWebGui = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MaximumResourcesAndCredentialsForThreshold`); value.Exists() {
		data.MaximumResourcesAndCredentialsForThreshold = types.Int64Value(value.Int())
	} else {
		data.MaximumResourcesAndCredentialsForThreshold = types.Int64Value(5000)
	}
}

func (data *DmSLMStatement) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `SLMId`); value.Exists() && !data.SlmId.IsNull() {
		data.SlmId = types.Int64Value(value.Int())
	} else {
		data.SlmId = types.Int64Null()
	}
	if value := res.Get(pathRoot + `UserString`); value.Exists() && !data.UserString.IsNull() {
		data.UserString = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserString = types.StringNull()
	}
	if value := res.Get(pathRoot + `CredClass`); value.Exists() && !data.CredClass.IsNull() {
		data.CredClass = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CredClass = types.StringNull()
	}
	if value := res.Get(pathRoot + `RsrcClass`); value.Exists() && !data.RsrcClass.IsNull() {
		data.RsrcClass = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RsrcClass = types.StringNull()
	}
	if value := res.Get(pathRoot + `Schedule`); value.Exists() && !data.Schedule.IsNull() {
		data.Schedule = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Schedule = types.StringNull()
	}
	if value := res.Get(pathRoot + `Action`); value.Exists() && !data.Action.IsNull() {
		data.Action = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Action = types.StringNull()
	}
	if value := res.Get(pathRoot + `ThreshIntervalLength`); value.Exists() && !data.ThreshIntervalLength.IsNull() {
		data.ThreshIntervalLength = types.Int64Value(value.Int())
	} else {
		data.ThreshIntervalLength = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ThreshIntervalType`); value.Exists() && !data.ThreshIntervalType.IsNull() {
		data.ThreshIntervalType = tfutils.ParseStringFromGJSON(value)
	} else if data.ThreshIntervalType.ValueString() != "fixed" {
		data.ThreshIntervalType = types.StringNull()
	}
	if value := res.Get(pathRoot + `ThreshAlgorithm`); value.Exists() && !data.ThreshAlgorithm.IsNull() {
		data.ThreshAlgorithm = tfutils.ParseStringFromGJSON(value)
	} else if data.ThreshAlgorithm.ValueString() != "greater-than" {
		data.ThreshAlgorithm = types.StringNull()
	}
	if value := res.Get(pathRoot + `ThresholdType`); value.Exists() && !data.ThresholdType.IsNull() {
		data.ThresholdType = tfutils.ParseStringFromGJSON(value)
	} else if data.ThresholdType.ValueString() != "count-all" {
		data.ThresholdType = types.StringNull()
	}
	if value := res.Get(pathRoot + `ThresholdLevel`); value.Exists() && !data.ThresholdLevel.IsNull() {
		data.ThresholdLevel = types.Int64Value(value.Int())
	} else if data.ThresholdLevel.ValueInt64() != 0 {
		data.ThresholdLevel = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ReleaseThresholdLevel`); value.Exists() && !data.ReleaseThresholdLevel.IsNull() {
		data.ReleaseThresholdLevel = types.Int64Value(value.Int())
	} else if data.ReleaseThresholdLevel.ValueInt64() != 0 {
		data.ReleaseThresholdLevel = types.Int64Null()
	}
	if value := res.Get(pathRoot + `BurstLimit`); value.Exists() && !data.BurstLimit.IsNull() {
		data.BurstLimit = types.Int64Value(value.Int())
	} else {
		data.BurstLimit = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ReportingAggregationInterval`); value.Exists() && !data.ReportingAggregationInterval.IsNull() {
		data.ReportingAggregationInterval = types.Int64Value(value.Int())
	} else {
		data.ReportingAggregationInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaximumTotalReportingRecords`); value.Exists() && !data.MaximumTotalReportingRecords.IsNull() {
		data.MaximumTotalReportingRecords = types.Int64Value(value.Int())
	} else if data.MaximumTotalReportingRecords.ValueInt64() != 5000 {
		data.MaximumTotalReportingRecords = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AutoGeneratedByWebGUI`); value.Exists() && !data.AutoGeneratedByWebGui.IsNull() {
		data.AutoGeneratedByWebGui = tfutils.BoolFromString(value.String())
	} else if data.AutoGeneratedByWebGui.ValueBool() {
		data.AutoGeneratedByWebGui = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MaximumResourcesAndCredentialsForThreshold`); value.Exists() && !data.MaximumResourcesAndCredentialsForThreshold.IsNull() {
		data.MaximumResourcesAndCredentialsForThreshold = types.Int64Value(value.Int())
	} else if data.MaximumResourcesAndCredentialsForThreshold.ValueInt64() != 5000 {
		data.MaximumResourcesAndCredentialsForThreshold = types.Int64Null()
	}
}
