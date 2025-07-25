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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmNetworkRetry struct {
	AutomaticRetry    types.Bool  `tfsdk:"automatic_retry"`
	RetryInterval     types.Int64 `tfsdk:"retry_interval"`
	ReportingInterval types.Int64 `tfsdk:"reporting_interval"`
	TotalRetries      types.Int64 `tfsdk:"total_retries"`
}

var DmNetworkRetryObjectType = map[string]attr.Type{
	"automatic_retry":    types.BoolType,
	"retry_interval":     types.Int64Type,
	"reporting_interval": types.Int64Type,
	"total_retries":      types.Int64Type,
}
var DmNetworkRetryObjectDefault = map[string]attr.Value{
	"automatic_retry":    types.BoolValue(false),
	"retry_interval":     types.Int64Value(1),
	"reporting_interval": types.Int64Value(1),
	"total_retries":      types.Int64Value(1),
}
var DmNetworkRetryDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"automatic_retry": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Automatic Retry", "auto-retry", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"retry_interval": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Retry Interval", "retry-interval", "").AddIntegerRange(1, 4294967295).AddDefaultValue("1").String,
			Computed:            true,
		},
		"reporting_interval": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Reporting Interval", "reporting-interval", "").AddIntegerRange(1, 4294967295).AddDefaultValue("1").String,
			Computed:            true,
		},
		"total_retries": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Total Retries", "total-retries", "").AddIntegerRange(1, 4294967295).AddDefaultValue("1").String,
			Computed:            true,
		},
	},
}
var DmNetworkRetryResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmNetworkRetryObjectType,
			DmNetworkRetryObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"automatic_retry": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Automatic Retry", "auto-retry", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"retry_interval": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Retry Interval", "retry-interval", "").AddIntegerRange(1, 4294967295).AddDefaultValue("1").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(1, 4294967295),
			},
			Default: int64default.StaticInt64(1),
		},
		"reporting_interval": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Reporting Interval", "reporting-interval", "").AddIntegerRange(1, 4294967295).AddDefaultValue("1").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(1, 4294967295),
			},
			Default: int64default.StaticInt64(1),
		},
		"total_retries": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Total Retries", "total-retries", "").AddIntegerRange(1, 4294967295).AddDefaultValue("1").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(1, 4294967295),
			},
			Default: int64default.StaticInt64(1),
		},
	},
}

func (data DmNetworkRetry) IsNull() bool {
	if !data.AutomaticRetry.IsNull() {
		return false
	}
	if !data.RetryInterval.IsNull() {
		return false
	}
	if !data.ReportingInterval.IsNull() {
		return false
	}
	if !data.TotalRetries.IsNull() {
		return false
	}
	return true
}
func GetDmNetworkRetryDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmNetworkRetryDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmNetworkRetryDataSourceSchema
}

func GetDmNetworkRetryResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmNetworkRetryResourceSchema.Required = true
	} else {
		DmNetworkRetryResourceSchema.Optional = true
		DmNetworkRetryResourceSchema.Computed = true
	}
	DmNetworkRetryResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmNetworkRetryResourceSchema
}

func (data DmNetworkRetry) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.AutomaticRetry.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AutomaticRetry`, tfutils.StringFromBool(data.AutomaticRetry, ""))
	}
	if !data.RetryInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RetryInterval`, data.RetryInterval.ValueInt64())
	}
	if !data.ReportingInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ReportingInterval`, data.ReportingInterval.ValueInt64())
	}
	if !data.TotalRetries.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TotalRetries`, data.TotalRetries.ValueInt64())
	}
	return body
}

func (data *DmNetworkRetry) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `AutomaticRetry`); value.Exists() {
		data.AutomaticRetry = tfutils.BoolFromString(value.String())
	} else {
		data.AutomaticRetry = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RetryInterval`); value.Exists() {
		data.RetryInterval = types.Int64Value(value.Int())
	} else {
		data.RetryInterval = types.Int64Value(1)
	}
	if value := res.Get(pathRoot + `ReportingInterval`); value.Exists() {
		data.ReportingInterval = types.Int64Value(value.Int())
	} else {
		data.ReportingInterval = types.Int64Value(1)
	}
	if value := res.Get(pathRoot + `TotalRetries`); value.Exists() {
		data.TotalRetries = types.Int64Value(value.Int())
	} else {
		data.TotalRetries = types.Int64Value(1)
	}
}

func (data *DmNetworkRetry) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `AutomaticRetry`); value.Exists() && !data.AutomaticRetry.IsNull() {
		data.AutomaticRetry = tfutils.BoolFromString(value.String())
	} else if data.AutomaticRetry.ValueBool() {
		data.AutomaticRetry = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RetryInterval`); value.Exists() && !data.RetryInterval.IsNull() {
		data.RetryInterval = types.Int64Value(value.Int())
	} else if data.RetryInterval.ValueInt64() != 1 {
		data.RetryInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ReportingInterval`); value.Exists() && !data.ReportingInterval.IsNull() {
		data.ReportingInterval = types.Int64Value(value.Int())
	} else if data.ReportingInterval.ValueInt64() != 1 {
		data.ReportingInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `TotalRetries`); value.Exists() && !data.TotalRetries.IsNull() {
		data.TotalRetries = types.Int64Value(value.Int())
	} else if data.TotalRetries.ValueInt64() != 1 {
		data.TotalRetries = types.Int64Null()
	}
}
