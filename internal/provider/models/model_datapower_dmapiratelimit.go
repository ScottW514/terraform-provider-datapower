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

type DmAPIRateLimit struct {
	Name         types.String `tfsdk:"name"`
	Rate         types.Int64  `tfsdk:"rate"`
	Interval     types.Int64  `tfsdk:"interval"`
	Unit         types.String `tfsdk:"unit"`
	HardLimit    types.Bool   `tfsdk:"hard_limit"`
	CacheOnly    types.Bool   `tfsdk:"cache_only"`
	IsClient     types.Bool   `tfsdk:"is_client"`
	UseApiName   types.Bool   `tfsdk:"use_api_name"`
	UseAppId     types.Bool   `tfsdk:"use_app_id"`
	UseClientId  types.Bool   `tfsdk:"use_client_id"`
	DynamicValue types.String `tfsdk:"dynamic_value"`
	Weight       types.String `tfsdk:"weight"`
}

var DmAPIRateLimitObjectType = map[string]attr.Type{
	"name":          types.StringType,
	"rate":          types.Int64Type,
	"interval":      types.Int64Type,
	"unit":          types.StringType,
	"hard_limit":    types.BoolType,
	"cache_only":    types.BoolType,
	"is_client":     types.BoolType,
	"use_api_name":  types.BoolType,
	"use_app_id":    types.BoolType,
	"use_client_id": types.BoolType,
	"dynamic_value": types.StringType,
	"weight":        types.StringType,
}
var DmAPIRateLimitObjectDefault = map[string]attr.Value{
	"name":          types.StringNull(),
	"rate":          types.Int64Null(),
	"interval":      types.Int64Value(1),
	"unit":          types.StringValue("second"),
	"hard_limit":    types.BoolValue(false),
	"cache_only":    types.BoolValue(true),
	"is_client":     types.BoolValue(true),
	"use_api_name":  types.BoolValue(false),
	"use_app_id":    types.BoolValue(false),
	"use_client_id": types.BoolValue(false),
	"dynamic_value": types.StringNull(),
	"weight":        types.StringValue("1"),
}
var DmAPIRateLimitDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Name", "", "").String,
			Computed:            true,
		},
		"rate": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Rate", "", "").AddIntegerRange(0, 4294967295).String,
			Computed:            true,
		},
		"interval": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Interval", "", "").AddIntegerRange(1, 65535).AddDefaultValue("1").String,
			Computed:            true,
		},
		"unit": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Unit", "", "").AddStringEnum("second", "minute", "hour", "day", "week").AddDefaultValue("second").String,
			Computed:            true,
		},
		"hard_limit": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enable hard limit", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"cache_only": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Cache only", "", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"is_client": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Is Client", "", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"use_api_name": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Use API Name", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"use_app_id": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Use Application ID", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"use_client_id": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Use Client ID", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"dynamic_value": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Dynamic Value", "", "").String,
			Computed:            true,
		},
		"weight": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Weight expression", "", "").AddDefaultValue("1").String,
			Computed:            true,
		},
	},
}
var DmAPIRateLimitResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Name", "", "").String,
			Required:            true,
		},
		"rate": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Rate", "", "").AddIntegerRange(0, 4294967295).String,
			Required:            true,
			Validators: []validator.Int64{
				int64validator.Between(0, 4294967295),
			},
		},
		"interval": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Interval", "", "").AddIntegerRange(1, 65535).AddDefaultValue("1").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(1, 65535),
			},
			Default: int64default.StaticInt64(1),
		},
		"unit": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Unit", "", "").AddStringEnum("second", "minute", "hour", "day", "week").AddDefaultValue("second").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("second", "minute", "hour", "day", "week"),
			},
			Default: stringdefault.StaticString("second"),
		},
		"hard_limit": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enable hard limit", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"cache_only": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Cache only", "", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"is_client": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Is Client", "", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"use_api_name": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Use API Name", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"use_app_id": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Use Application ID", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"use_client_id": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Use Client ID", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"dynamic_value": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Dynamic Value", "", "").String,
			Optional:            true,
		},
		"weight": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Weight expression", "", "").AddDefaultValue("1").String,
			Computed:            true,
			Optional:            true,
			Default:             stringdefault.StaticString("1"),
		},
	},
}

func (data DmAPIRateLimit) IsNull() bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.Rate.IsNull() {
		return false
	}
	if !data.Interval.IsNull() {
		return false
	}
	if !data.Unit.IsNull() {
		return false
	}
	if !data.HardLimit.IsNull() {
		return false
	}
	if !data.CacheOnly.IsNull() {
		return false
	}
	if !data.IsClient.IsNull() {
		return false
	}
	if !data.UseApiName.IsNull() {
		return false
	}
	if !data.UseAppId.IsNull() {
		return false
	}
	if !data.UseClientId.IsNull() {
		return false
	}
	if !data.DynamicValue.IsNull() {
		return false
	}
	if !data.Weight.IsNull() {
		return false
	}
	return true
}

func (data DmAPIRateLimit) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Name`, data.Name.ValueString())
	}
	if !data.Rate.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Rate`, data.Rate.ValueInt64())
	}
	if !data.Interval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Interval`, data.Interval.ValueInt64())
	}
	if !data.Unit.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Unit`, data.Unit.ValueString())
	}
	if !data.HardLimit.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HardLimit`, tfutils.StringFromBool(data.HardLimit, ""))
	}
	if !data.CacheOnly.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CacheOnly`, tfutils.StringFromBool(data.CacheOnly, ""))
	}
	if !data.IsClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IsClient`, tfutils.StringFromBool(data.IsClient, ""))
	}
	if !data.UseApiName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseApiName`, tfutils.StringFromBool(data.UseApiName, ""))
	}
	if !data.UseAppId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseAppId`, tfutils.StringFromBool(data.UseAppId, ""))
	}
	if !data.UseClientId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseClientId`, tfutils.StringFromBool(data.UseClientId, ""))
	}
	if !data.DynamicValue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DynamicValue`, data.DynamicValue.ValueString())
	}
	if !data.Weight.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Weight`, data.Weight.ValueString())
	}
	return body
}

func (data *DmAPIRateLimit) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Rate`); value.Exists() {
		data.Rate = types.Int64Value(value.Int())
	} else {
		data.Rate = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Interval`); value.Exists() {
		data.Interval = types.Int64Value(value.Int())
	} else {
		data.Interval = types.Int64Value(1)
	}
	if value := res.Get(pathRoot + `Unit`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Unit = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Unit = types.StringValue("second")
	}
	if value := res.Get(pathRoot + `HardLimit`); value.Exists() {
		data.HardLimit = tfutils.BoolFromString(value.String())
	} else {
		data.HardLimit = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CacheOnly`); value.Exists() {
		data.CacheOnly = tfutils.BoolFromString(value.String())
	} else {
		data.CacheOnly = types.BoolNull()
	}
	if value := res.Get(pathRoot + `IsClient`); value.Exists() {
		data.IsClient = tfutils.BoolFromString(value.String())
	} else {
		data.IsClient = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UseApiName`); value.Exists() {
		data.UseApiName = tfutils.BoolFromString(value.String())
	} else {
		data.UseApiName = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UseAppId`); value.Exists() {
		data.UseAppId = tfutils.BoolFromString(value.String())
	} else {
		data.UseAppId = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UseClientId`); value.Exists() {
		data.UseClientId = tfutils.BoolFromString(value.String())
	} else {
		data.UseClientId = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DynamicValue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DynamicValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DynamicValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `Weight`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Weight = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Weight = types.StringValue("1")
	}
}

func (data *DmAPIRateLimit) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && !data.Name.IsNull() {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Rate`); value.Exists() && !data.Rate.IsNull() {
		data.Rate = types.Int64Value(value.Int())
	} else {
		data.Rate = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Interval`); value.Exists() && !data.Interval.IsNull() {
		data.Interval = types.Int64Value(value.Int())
	} else if data.Interval.ValueInt64() != 1 {
		data.Interval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Unit`); value.Exists() && !data.Unit.IsNull() {
		data.Unit = tfutils.ParseStringFromGJSON(value)
	} else if data.Unit.ValueString() != "second" {
		data.Unit = types.StringNull()
	}
	if value := res.Get(pathRoot + `HardLimit`); value.Exists() && !data.HardLimit.IsNull() {
		data.HardLimit = tfutils.BoolFromString(value.String())
	} else if data.HardLimit.ValueBool() {
		data.HardLimit = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CacheOnly`); value.Exists() && !data.CacheOnly.IsNull() {
		data.CacheOnly = tfutils.BoolFromString(value.String())
	} else if !data.CacheOnly.ValueBool() {
		data.CacheOnly = types.BoolNull()
	}
	if value := res.Get(pathRoot + `IsClient`); value.Exists() && !data.IsClient.IsNull() {
		data.IsClient = tfutils.BoolFromString(value.String())
	} else if !data.IsClient.ValueBool() {
		data.IsClient = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UseApiName`); value.Exists() && !data.UseApiName.IsNull() {
		data.UseApiName = tfutils.BoolFromString(value.String())
	} else if data.UseApiName.ValueBool() {
		data.UseApiName = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UseAppId`); value.Exists() && !data.UseAppId.IsNull() {
		data.UseAppId = tfutils.BoolFromString(value.String())
	} else if data.UseAppId.ValueBool() {
		data.UseAppId = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UseClientId`); value.Exists() && !data.UseClientId.IsNull() {
		data.UseClientId = tfutils.BoolFromString(value.String())
	} else if data.UseClientId.ValueBool() {
		data.UseClientId = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DynamicValue`); value.Exists() && !data.DynamicValue.IsNull() {
		data.DynamicValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DynamicValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `Weight`); value.Exists() && !data.Weight.IsNull() {
		data.Weight = tfutils.ParseStringFromGJSON(value)
	} else if data.Weight.ValueString() != "1" {
		data.Weight = types.StringNull()
	}
}
