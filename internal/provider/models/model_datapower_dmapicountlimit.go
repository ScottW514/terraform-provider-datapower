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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmAPICountLimit struct {
	Name         types.String `tfsdk:"name"`
	Count        types.Int64  `tfsdk:"count"`
	HardLimit    types.Bool   `tfsdk:"hard_limit"`
	CacheOnly    types.Bool   `tfsdk:"cache_only"`
	IsClient     types.Bool   `tfsdk:"is_client"`
	UseApiName   types.Bool   `tfsdk:"use_api_name"`
	UseAppId     types.Bool   `tfsdk:"use_app_id"`
	UseClientId  types.Bool   `tfsdk:"use_client_id"`
	DynamicValue types.String `tfsdk:"dynamic_value"`
	Weight       types.String `tfsdk:"weight"`
	AutoDec      types.Bool   `tfsdk:"auto_dec"`
}

var DmAPICountLimitObjectType = map[string]attr.Type{
	"name":          types.StringType,
	"count":         types.Int64Type,
	"hard_limit":    types.BoolType,
	"cache_only":    types.BoolType,
	"is_client":     types.BoolType,
	"use_api_name":  types.BoolType,
	"use_app_id":    types.BoolType,
	"use_client_id": types.BoolType,
	"dynamic_value": types.StringType,
	"weight":        types.StringType,
	"auto_dec":      types.BoolType,
}
var DmAPICountLimitObjectDefault = map[string]attr.Value{
	"name":          types.StringNull(),
	"count":         types.Int64Null(),
	"hard_limit":    types.BoolValue(true),
	"cache_only":    types.BoolValue(true),
	"is_client":     types.BoolValue(false),
	"use_api_name":  types.BoolValue(false),
	"use_app_id":    types.BoolValue(false),
	"use_client_id": types.BoolValue(false),
	"dynamic_value": types.StringNull(),
	"weight":        types.StringValue("1"),
	"auto_dec":      types.BoolValue(true),
}
var DmAPICountLimitDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Name", "", "").String,
			Computed:            true,
		},
		"count": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Count", "", "").String,
			Computed:            true,
		},
		"hard_limit": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enable hard limit", "", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"cache_only": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Cache only", "", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"is_client": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Is Client", "", "").AddDefaultValue("false").String,
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
		"auto_dec": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Auto decrement", "", "").AddDefaultValue("true").String,
			Computed:            true,
		},
	},
}
var DmAPICountLimitResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Name", "", "").String,
			Required:            true,
		},
		"count": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Count", "", "").String,
			Required:            true,
		},
		"hard_limit": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enable hard limit", "", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"cache_only": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Cache only", "", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"is_client": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Is Client", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
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
		"auto_dec": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Auto decrement", "", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
	},
}

func (data DmAPICountLimit) IsNull() bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.Count.IsNull() {
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
	if !data.AutoDec.IsNull() {
		return false
	}
	return true
}

func (data DmAPICountLimit) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Name`, data.Name.ValueString())
	}
	if !data.Count.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Count`, data.Count.ValueInt64())
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
	if !data.AutoDec.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AutoDec`, tfutils.StringFromBool(data.AutoDec, ""))
	}
	return body
}

func (data *DmAPICountLimit) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Count`); value.Exists() {
		data.Count = types.Int64Value(value.Int())
	} else {
		data.Count = types.Int64Null()
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
	if value := res.Get(pathRoot + `AutoDec`); value.Exists() {
		data.AutoDec = tfutils.BoolFromString(value.String())
	} else {
		data.AutoDec = types.BoolNull()
	}
}

func (data *DmAPICountLimit) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && !data.Name.IsNull() {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Count`); value.Exists() && !data.Count.IsNull() {
		data.Count = types.Int64Value(value.Int())
	} else {
		data.Count = types.Int64Null()
	}
	if value := res.Get(pathRoot + `HardLimit`); value.Exists() && !data.HardLimit.IsNull() {
		data.HardLimit = tfutils.BoolFromString(value.String())
	} else if !data.HardLimit.ValueBool() {
		data.HardLimit = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CacheOnly`); value.Exists() && !data.CacheOnly.IsNull() {
		data.CacheOnly = tfutils.BoolFromString(value.String())
	} else if !data.CacheOnly.ValueBool() {
		data.CacheOnly = types.BoolNull()
	}
	if value := res.Get(pathRoot + `IsClient`); value.Exists() && !data.IsClient.IsNull() {
		data.IsClient = tfutils.BoolFromString(value.String())
	} else if data.IsClient.ValueBool() {
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
	if value := res.Get(pathRoot + `AutoDec`); value.Exists() && !data.AutoDec.IsNull() {
		data.AutoDec = tfutils.BoolFromString(value.String())
	} else if !data.AutoDec.ValueBool() {
		data.AutoDec = types.BoolNull()
	}
}
