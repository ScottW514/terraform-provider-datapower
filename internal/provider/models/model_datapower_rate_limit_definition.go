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
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type RateLimitDefinition struct {
	Id                  types.String                `tfsdk:"id"`
	AppDomain           types.String                `tfsdk:"app_domain"`
	ShortName           types.String                `tfsdk:"short_name"`
	UserSummary         types.String                `tfsdk:"user_summary"`
	Type                types.String                `tfsdk:"type"`
	Rate                types.Int64                 `tfsdk:"rate"`
	Interval            types.Int64                 `tfsdk:"interval"`
	Unit                types.String                `tfsdk:"unit"`
	HardLimit           types.Bool                  `tfsdk:"hard_limit"`
	IsClient            types.Bool                  `tfsdk:"is_client"`
	UseApiName          types.Bool                  `tfsdk:"use_api_name"`
	UseAppId            types.Bool                  `tfsdk:"use_app_id"`
	UseClientId         types.Bool                  `tfsdk:"use_client_id"`
	AutoReplenish       types.Bool                  `tfsdk:"auto_replenish"`
	DynamicValue        types.String                `tfsdk:"dynamic_value"`
	Weight              types.String                `tfsdk:"weight"`
	ResponseHeaders     types.Bool                  `tfsdk:"response_headers"`
	EmulateBurstHeaders types.Bool                  `tfsdk:"emulate_burst_headers"`
	UseIntervalOffset   types.Bool                  `tfsdk:"use_interval_offset"`
	AllowCacheFallback  types.Bool                  `tfsdk:"allow_cache_fallback"`
	UseCache            types.Bool                  `tfsdk:"use_cache"`
	Parameters          types.List                  `tfsdk:"parameters"`
	DependencyActions   []*actions.DependencyAction `tfsdk:"dependency_actions"`
	ProviderTarget      types.String                `tfsdk:"provider_target"`
}

var RateLimitDefinitionIntervalIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "rate",
			AttrType:    "Int64",
			AttrDefault: "",
			Value:       []string{"0"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "rate",
			Value:       []string{"count"},
		},
	},
}

var RateLimitDefinitionUnitIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "rate",
			AttrType:    "Int64",
			AttrDefault: "",
			Value:       []string{"0"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "rate",
			Value:       []string{"count"},
		},
	},
}

var RateLimitDefinitionHardLimitIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "rate",
			AttrType:    "Int64",
			AttrDefault: "",
			Value:       []string{"0"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "rate",
			Value:       []string{"burst"},
		},
	},
}

var RateLimitDefinitionIsClientIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "rate",
	AttrType:    "Int64",
	AttrDefault: "",
	Value:       []string{"0"},
}

var RateLimitDefinitionUseApiNameIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "rate",
	AttrType:    "Int64",
	AttrDefault: "",
	Value:       []string{"0"},
}

var RateLimitDefinitionUseAppIdIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "rate",
	AttrType:    "Int64",
	AttrDefault: "",
	Value:       []string{"0"},
}

var RateLimitDefinitionUseClientIdIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "rate",
	AttrType:    "Int64",
	AttrDefault: "",
	Value:       []string{"0"},
}

var RateLimitDefinitionAutoReplenishIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "rate",
	Value:       []string{"count"},
}

var RateLimitDefinitionDynamicValueIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "rate",
	AttrType:    "Int64",
	AttrDefault: "",
	Value:       []string{"0"},
}

var RateLimitDefinitionWeightIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "rate",
	AttrType:    "Int64",
	AttrDefault: "",
	Value:       []string{"0"},
}

var RateLimitDefinitionEmulateBurstHeadersIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "rate",
	Value:       []string{"rate"},
}

var RateLimitDefinitionUseIntervalOffsetIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "rate",
	Value:       []string{"count"},
}

var RateLimitDefinitionObjectType = map[string]attr.Type{
	"provider_target":       types.StringType,
	"id":                    types.StringType,
	"app_domain":            types.StringType,
	"short_name":            types.StringType,
	"user_summary":          types.StringType,
	"type":                  types.StringType,
	"rate":                  types.Int64Type,
	"interval":              types.Int64Type,
	"unit":                  types.StringType,
	"hard_limit":            types.BoolType,
	"is_client":             types.BoolType,
	"use_api_name":          types.BoolType,
	"use_app_id":            types.BoolType,
	"use_client_id":         types.BoolType,
	"auto_replenish":        types.BoolType,
	"dynamic_value":         types.StringType,
	"weight":                types.StringType,
	"response_headers":      types.BoolType,
	"emulate_burst_headers": types.BoolType,
	"use_interval_offset":   types.BoolType,
	"allow_cache_fallback":  types.BoolType,
	"use_cache":             types.BoolType,
	"parameters":            types.ListType{ElemType: types.ObjectType{AttrTypes: DmRateLimitDefinitionNameValuePairObjectType}},
	"dependency_actions":    actions.ActionsListType,
}

func (data RateLimitDefinition) GetPath() string {
	rest_path := "/mgmt/config/{domain}/RateLimitDefinition"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data RateLimitDefinition) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.ShortName.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
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
	if !data.AutoReplenish.IsNull() {
		return false
	}
	if !data.DynamicValue.IsNull() {
		return false
	}
	if !data.Weight.IsNull() {
		return false
	}
	if !data.ResponseHeaders.IsNull() {
		return false
	}
	if !data.EmulateBurstHeaders.IsNull() {
		return false
	}
	if !data.UseIntervalOffset.IsNull() {
		return false
	}
	if !data.AllowCacheFallback.IsNull() {
		return false
	}
	if !data.UseCache.IsNull() {
		return false
	}
	if !data.Parameters.IsNull() {
		return false
	}
	return true
}

func (data RateLimitDefinition) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.ShortName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ShortName`, data.ShortName.ValueString())
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Type`, data.Type.ValueString())
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
	if !data.AutoReplenish.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AutoReplenish`, tfutils.StringFromBool(data.AutoReplenish, ""))
	}
	if !data.DynamicValue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DynamicValue`, data.DynamicValue.ValueString())
	}
	if !data.Weight.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Weight`, data.Weight.ValueString())
	}
	if !data.ResponseHeaders.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ResponseHeaders`, tfutils.StringFromBool(data.ResponseHeaders, ""))
	}
	if !data.EmulateBurstHeaders.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EmulateBurstHeaders`, tfutils.StringFromBool(data.EmulateBurstHeaders, ""))
	}
	if !data.UseIntervalOffset.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseIntervalOffset`, tfutils.StringFromBool(data.UseIntervalOffset, ""))
	}
	if !data.AllowCacheFallback.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllowCacheFallback`, tfutils.StringFromBool(data.AllowCacheFallback, ""))
	}
	if !data.UseCache.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseCache`, tfutils.StringFromBool(data.UseCache, ""))
	}
	if !data.Parameters.IsNull() {
		var dataValues []DmRateLimitDefinitionNameValuePair
		data.Parameters.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.SetRaw(body, pathRoot+`Parameters`+".-1", val.ToBody(ctx, ""))
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`Parameters`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`Parameters`, "[]")
	}
	return body
}

func (data *RateLimitDefinition) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `ShortName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ShortName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ShortName = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringValue("rate")
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
		data.Unit = types.StringValue("minute")
	}
	if value := res.Get(pathRoot + `HardLimit`); value.Exists() {
		data.HardLimit = tfutils.BoolFromString(value.String())
	} else {
		data.HardLimit = types.BoolNull()
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
	if value := res.Get(pathRoot + `AutoReplenish`); value.Exists() {
		data.AutoReplenish = tfutils.BoolFromString(value.String())
	} else {
		data.AutoReplenish = types.BoolNull()
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
	if value := res.Get(pathRoot + `ResponseHeaders`); value.Exists() {
		data.ResponseHeaders = tfutils.BoolFromString(value.String())
	} else {
		data.ResponseHeaders = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EmulateBurstHeaders`); value.Exists() {
		data.EmulateBurstHeaders = tfutils.BoolFromString(value.String())
	} else {
		data.EmulateBurstHeaders = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UseIntervalOffset`); value.Exists() {
		data.UseIntervalOffset = tfutils.BoolFromString(value.String())
	} else {
		data.UseIntervalOffset = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowCacheFallback`); value.Exists() {
		data.AllowCacheFallback = tfutils.BoolFromString(value.String())
	} else {
		data.AllowCacheFallback = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UseCache`); value.Exists() {
		data.UseCache = tfutils.BoolFromString(value.String())
	} else {
		data.UseCache = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Parameters`); value.Exists() {
		l := []DmRateLimitDefinitionNameValuePair{}
		if value := res.Get(`Parameters`); value.Exists() {
			for _, v := range value.Array() {
				item := DmRateLimitDefinitionNameValuePair{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Parameters, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmRateLimitDefinitionNameValuePairObjectType}, l)
		} else {
			data.Parameters = types.ListNull(types.ObjectType{AttrTypes: DmRateLimitDefinitionNameValuePairObjectType})
		}
	} else {
		data.Parameters = types.ListNull(types.ObjectType{AttrTypes: DmRateLimitDefinitionNameValuePairObjectType})
	}
}

func (data *RateLimitDefinition) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `ShortName`); value.Exists() && !data.ShortName.IsNull() {
		data.ShortName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ShortName = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && !data.Type.IsNull() {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else if data.Type.ValueString() != "rate" {
		data.Type = types.StringNull()
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
	} else if data.Unit.ValueString() != "minute" {
		data.Unit = types.StringNull()
	}
	if value := res.Get(pathRoot + `HardLimit`); value.Exists() && !data.HardLimit.IsNull() {
		data.HardLimit = tfutils.BoolFromString(value.String())
	} else if !data.HardLimit.ValueBool() {
		data.HardLimit = types.BoolNull()
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
	if value := res.Get(pathRoot + `AutoReplenish`); value.Exists() && !data.AutoReplenish.IsNull() {
		data.AutoReplenish = tfutils.BoolFromString(value.String())
	} else if !data.AutoReplenish.ValueBool() {
		data.AutoReplenish = types.BoolNull()
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
	if value := res.Get(pathRoot + `ResponseHeaders`); value.Exists() && !data.ResponseHeaders.IsNull() {
		data.ResponseHeaders = tfutils.BoolFromString(value.String())
	} else if !data.ResponseHeaders.ValueBool() {
		data.ResponseHeaders = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EmulateBurstHeaders`); value.Exists() && !data.EmulateBurstHeaders.IsNull() {
		data.EmulateBurstHeaders = tfutils.BoolFromString(value.String())
	} else if data.EmulateBurstHeaders.ValueBool() {
		data.EmulateBurstHeaders = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UseIntervalOffset`); value.Exists() && !data.UseIntervalOffset.IsNull() {
		data.UseIntervalOffset = tfutils.BoolFromString(value.String())
	} else if !data.UseIntervalOffset.ValueBool() {
		data.UseIntervalOffset = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowCacheFallback`); value.Exists() && !data.AllowCacheFallback.IsNull() {
		data.AllowCacheFallback = tfutils.BoolFromString(value.String())
	} else if !data.AllowCacheFallback.ValueBool() {
		data.AllowCacheFallback = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UseCache`); value.Exists() && !data.UseCache.IsNull() {
		data.UseCache = tfutils.BoolFromString(value.String())
	} else if data.UseCache.ValueBool() {
		data.UseCache = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Parameters`); value.Exists() && !data.Parameters.IsNull() {
		l := []DmRateLimitDefinitionNameValuePair{}
		e := []DmRateLimitDefinitionNameValuePair{}
		data.Parameters.ElementsAs(ctx, &e, false)
		if len(value.Array()) == len(e) {
			for i, v := range value.Array() {
				item := e[i]
				item.UpdateFromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		} else {
			for _, v := range value.Array() {
				item := DmRateLimitDefinitionNameValuePair{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Parameters, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmRateLimitDefinitionNameValuePairObjectType}, l)
		} else {
			data.Parameters = types.ListNull(types.ObjectType{AttrTypes: DmRateLimitDefinitionNameValuePairObjectType})
		}
	} else {
		data.Parameters = types.ListNull(types.ObjectType{AttrTypes: DmRateLimitDefinitionNameValuePairObjectType})
	}
}
