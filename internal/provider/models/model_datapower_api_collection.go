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

type APICollection struct {
	Id                           types.String                     `tfsdk:"id"`
	AppDomain                    types.String                     `tfsdk:"app_domain"`
	UserSummary                  types.String                     `tfsdk:"user_summary"`
	Sandbox                      types.Bool                       `tfsdk:"sandbox"`
	OrgId                        types.String                     `tfsdk:"org_id"`
	OrgName                      types.String                     `tfsdk:"org_name"`
	CatalogId                    types.String                     `tfsdk:"catalog_id"`
	CatalogName                  types.String                     `tfsdk:"catalog_name"`
	DevPortalEndpoint            types.String                     `tfsdk:"dev_portal_endpoint"`
	CacheCapacity                types.Int64                      `tfsdk:"cache_capacity"`
	RoutingPrefix                types.List                       `tfsdk:"routing_prefix"`
	UseRateLimitGroup            types.Bool                       `tfsdk:"use_rate_limit_group"`
	DefaultRateLimit             types.List                       `tfsdk:"default_rate_limit"`
	RateLimitGroup               types.String                     `tfsdk:"rate_limit_group"`
	AssemblyBurstLimit           types.List                       `tfsdk:"assembly_burst_limit"`
	AssemblyRateLimit            types.List                       `tfsdk:"assembly_rate_limit"`
	AssemblyCountLimit           types.List                       `tfsdk:"assembly_count_limit"`
	EnforcePreAssemblyRateLimits types.Bool                       `tfsdk:"enforce_pre_assembly_rate_limits"`
	ApiProcessingRule            types.String                     `tfsdk:"api_processing_rule"`
	ApiErrorRule                 types.String                     `tfsdk:"api_error_rule"`
	AssemblyPreflow              types.String                     `tfsdk:"assembly_preflow"`
	AssemblyPostflow             types.String                     `tfsdk:"assembly_postflow"`
	Plan                         types.List                       `tfsdk:"plan"`
	AnalyticsEndpoint            types.String                     `tfsdk:"analytics_endpoint"`
	ApplicationType              types.List                       `tfsdk:"application_type"`
	ParseSettingsReference       *DmDynamicParseSettingsReference `tfsdk:"parse_settings_reference"`
	DependencyActions            []*actions.DependencyAction      `tfsdk:"dependency_actions"`
}

var APICollectionDefaultRateLimitIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "use_rate_limit_group",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var APICollectionRateLimitGroupIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "use_rate_limit_group",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var APICollectionObjectType = map[string]attr.Type{
	"id":                               types.StringType,
	"app_domain":                       types.StringType,
	"user_summary":                     types.StringType,
	"sandbox":                          types.BoolType,
	"org_id":                           types.StringType,
	"org_name":                         types.StringType,
	"catalog_id":                       types.StringType,
	"catalog_name":                     types.StringType,
	"dev_portal_endpoint":              types.StringType,
	"cache_capacity":                   types.Int64Type,
	"routing_prefix":                   types.ListType{ElemType: types.ObjectType{AttrTypes: DmRoutingPrefixObjectType}},
	"use_rate_limit_group":             types.BoolType,
	"default_rate_limit":               types.ListType{ElemType: types.ObjectType{AttrTypes: DmAPIRateLimitObjectType}},
	"rate_limit_group":                 types.StringType,
	"assembly_burst_limit":             types.ListType{ElemType: types.ObjectType{AttrTypes: DmAPIBurstLimitObjectType}},
	"assembly_rate_limit":              types.ListType{ElemType: types.ObjectType{AttrTypes: DmAPIRateLimitObjectType}},
	"assembly_count_limit":             types.ListType{ElemType: types.ObjectType{AttrTypes: DmAPICountLimitObjectType}},
	"enforce_pre_assembly_rate_limits": types.BoolType,
	"api_processing_rule":              types.StringType,
	"api_error_rule":                   types.StringType,
	"assembly_preflow":                 types.StringType,
	"assembly_postflow":                types.StringType,
	"plan":                             types.ListType{ElemType: types.StringType},
	"analytics_endpoint":               types.StringType,
	"application_type":                 types.ListType{ElemType: types.StringType},
	"parse_settings_reference":         types.ObjectType{AttrTypes: DmDynamicParseSettingsReferenceObjectType},
	"dependency_actions":               actions.ActionsListType,
}

func (data APICollection) GetPath() string {
	rest_path := "/mgmt/config/{domain}/APICollection"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data APICollection) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Sandbox.IsNull() {
		return false
	}
	if !data.OrgId.IsNull() {
		return false
	}
	if !data.OrgName.IsNull() {
		return false
	}
	if !data.CatalogId.IsNull() {
		return false
	}
	if !data.CatalogName.IsNull() {
		return false
	}
	if !data.DevPortalEndpoint.IsNull() {
		return false
	}
	if !data.CacheCapacity.IsNull() {
		return false
	}
	if !data.RoutingPrefix.IsNull() {
		return false
	}
	if !data.UseRateLimitGroup.IsNull() {
		return false
	}
	if !data.DefaultRateLimit.IsNull() {
		return false
	}
	if !data.RateLimitGroup.IsNull() {
		return false
	}
	if !data.AssemblyBurstLimit.IsNull() {
		return false
	}
	if !data.AssemblyRateLimit.IsNull() {
		return false
	}
	if !data.AssemblyCountLimit.IsNull() {
		return false
	}
	if !data.EnforcePreAssemblyRateLimits.IsNull() {
		return false
	}
	if !data.ApiProcessingRule.IsNull() {
		return false
	}
	if !data.ApiErrorRule.IsNull() {
		return false
	}
	if !data.AssemblyPreflow.IsNull() {
		return false
	}
	if !data.AssemblyPostflow.IsNull() {
		return false
	}
	if !data.Plan.IsNull() {
		return false
	}
	if !data.AnalyticsEndpoint.IsNull() {
		return false
	}
	if !data.ApplicationType.IsNull() {
		return false
	}
	if data.ParseSettingsReference != nil {
		if !data.ParseSettingsReference.IsNull() {
			return false
		}
	}
	return true
}

func (data APICollection) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.Sandbox.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Sandbox`, tfutils.StringFromBool(data.Sandbox, ""))
	}
	if !data.OrgId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OrgID`, data.OrgId.ValueString())
	}
	if !data.OrgName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OrgName`, data.OrgName.ValueString())
	}
	if !data.CatalogId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CatalogID`, data.CatalogId.ValueString())
	}
	if !data.CatalogName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CatalogName`, data.CatalogName.ValueString())
	}
	if !data.DevPortalEndpoint.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DevPortalEndpoint`, data.DevPortalEndpoint.ValueString())
	}
	if !data.CacheCapacity.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CacheCapacity`, data.CacheCapacity.ValueInt64())
	}
	if !data.RoutingPrefix.IsNull() {
		var dataValues []DmRoutingPrefix
		data.RoutingPrefix.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`RoutingPrefix`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.UseRateLimitGroup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseRateLimitGroup`, tfutils.StringFromBool(data.UseRateLimitGroup, ""))
	}
	if !data.DefaultRateLimit.IsNull() {
		var dataValues []DmAPIRateLimit
		data.DefaultRateLimit.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`DefaultRateLimit`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.RateLimitGroup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RateLimitGroup`, data.RateLimitGroup.ValueString())
	}
	if !data.AssemblyBurstLimit.IsNull() {
		var dataValues []DmAPIBurstLimit
		data.AssemblyBurstLimit.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`AssemblyBurstLimit`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.AssemblyRateLimit.IsNull() {
		var dataValues []DmAPIRateLimit
		data.AssemblyRateLimit.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`AssemblyRateLimit`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.AssemblyCountLimit.IsNull() {
		var dataValues []DmAPICountLimit
		data.AssemblyCountLimit.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`AssemblyCountLimit`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.EnforcePreAssemblyRateLimits.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnforcePreAssemblyRateLimits`, tfutils.StringFromBool(data.EnforcePreAssemblyRateLimits, ""))
	}
	if !data.ApiProcessingRule.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APIProcessingRule`, data.ApiProcessingRule.ValueString())
	}
	if !data.ApiErrorRule.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APIErrorRule`, data.ApiErrorRule.ValueString())
	}
	if !data.AssemblyPreflow.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AssemblyPreflow`, data.AssemblyPreflow.ValueString())
	}
	if !data.AssemblyPostflow.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AssemblyPostflow`, data.AssemblyPostflow.ValueString())
	}
	if !data.Plan.IsNull() {
		var dataValues []string
		data.Plan.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`Plan`+".-1", map[string]string{"value": val})
		}
	}
	if !data.AnalyticsEndpoint.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AnalyticsEndpoint`, data.AnalyticsEndpoint.ValueString())
	}
	if !data.ApplicationType.IsNull() {
		var dataValues []string
		data.ApplicationType.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`ApplicationType`+".-1", map[string]string{"value": val})
		}
	}
	if data.ParseSettingsReference != nil {
		if !data.ParseSettingsReference.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`ParseSettingsReference`, data.ParseSettingsReference.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *APICollection) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `Sandbox`); value.Exists() {
		data.Sandbox = tfutils.BoolFromString(value.String())
	} else {
		data.Sandbox = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OrgID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OrgId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OrgId = types.StringNull()
	}
	if value := res.Get(pathRoot + `OrgName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OrgName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OrgName = types.StringNull()
	}
	if value := res.Get(pathRoot + `CatalogID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CatalogId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CatalogId = types.StringValue("default-catalog-id")
	}
	if value := res.Get(pathRoot + `CatalogName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CatalogName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CatalogName = types.StringValue("default")
	}
	if value := res.Get(pathRoot + `DevPortalEndpoint`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DevPortalEndpoint = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DevPortalEndpoint = types.StringNull()
	}
	if value := res.Get(pathRoot + `CacheCapacity`); value.Exists() {
		data.CacheCapacity = types.Int64Value(value.Int())
	} else {
		data.CacheCapacity = types.Int64Value(128)
	}
	if value := res.Get(pathRoot + `RoutingPrefix`); value.Exists() {
		l := []DmRoutingPrefix{}
		if value := res.Get(`RoutingPrefix`); value.Exists() {
			for _, v := range value.Array() {
				item := DmRoutingPrefix{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.RoutingPrefix, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmRoutingPrefixObjectType}, l)
		} else {
			data.RoutingPrefix = types.ListNull(types.ObjectType{AttrTypes: DmRoutingPrefixObjectType})
		}
	} else {
		data.RoutingPrefix = types.ListNull(types.ObjectType{AttrTypes: DmRoutingPrefixObjectType})
	}
	if value := res.Get(pathRoot + `UseRateLimitGroup`); value.Exists() {
		data.UseRateLimitGroup = tfutils.BoolFromString(value.String())
	} else {
		data.UseRateLimitGroup = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DefaultRateLimit`); value.Exists() {
		l := []DmAPIRateLimit{}
		if value := res.Get(`DefaultRateLimit`); value.Exists() {
			for _, v := range value.Array() {
				item := DmAPIRateLimit{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.DefaultRateLimit, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPIRateLimitObjectType}, l)
		} else {
			data.DefaultRateLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIRateLimitObjectType})
		}
	} else {
		data.DefaultRateLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIRateLimitObjectType})
	}
	if value := res.Get(pathRoot + `RateLimitGroup`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RateLimitGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RateLimitGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `AssemblyBurstLimit`); value.Exists() {
		l := []DmAPIBurstLimit{}
		if value := res.Get(`AssemblyBurstLimit`); value.Exists() {
			for _, v := range value.Array() {
				item := DmAPIBurstLimit{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.AssemblyBurstLimit, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPIBurstLimitObjectType}, l)
		} else {
			data.AssemblyBurstLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIBurstLimitObjectType})
		}
	} else {
		data.AssemblyBurstLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIBurstLimitObjectType})
	}
	if value := res.Get(pathRoot + `AssemblyRateLimit`); value.Exists() {
		l := []DmAPIRateLimit{}
		if value := res.Get(`AssemblyRateLimit`); value.Exists() {
			for _, v := range value.Array() {
				item := DmAPIRateLimit{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.AssemblyRateLimit, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPIRateLimitObjectType}, l)
		} else {
			data.AssemblyRateLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIRateLimitObjectType})
		}
	} else {
		data.AssemblyRateLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIRateLimitObjectType})
	}
	if value := res.Get(pathRoot + `AssemblyCountLimit`); value.Exists() {
		l := []DmAPICountLimit{}
		if value := res.Get(`AssemblyCountLimit`); value.Exists() {
			for _, v := range value.Array() {
				item := DmAPICountLimit{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.AssemblyCountLimit, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPICountLimitObjectType}, l)
		} else {
			data.AssemblyCountLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPICountLimitObjectType})
		}
	} else {
		data.AssemblyCountLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPICountLimitObjectType})
	}
	if value := res.Get(pathRoot + `EnforcePreAssemblyRateLimits`); value.Exists() {
		data.EnforcePreAssemblyRateLimits = tfutils.BoolFromString(value.String())
	} else {
		data.EnforcePreAssemblyRateLimits = types.BoolNull()
	}
	if value := res.Get(pathRoot + `APIProcessingRule`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ApiProcessingRule = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ApiProcessingRule = types.StringValue("default-api-rule")
	}
	if value := res.Get(pathRoot + `APIErrorRule`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ApiErrorRule = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ApiErrorRule = types.StringValue("default-api-error-rule")
	}
	if value := res.Get(pathRoot + `AssemblyPreflow`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AssemblyPreflow = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AssemblyPreflow = types.StringNull()
	}
	if value := res.Get(pathRoot + `AssemblyPostflow`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AssemblyPostflow = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AssemblyPostflow = types.StringNull()
	}
	if value := res.Get(pathRoot + `Plan`); value.Exists() {
		data.Plan = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Plan = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `AnalyticsEndpoint`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AnalyticsEndpoint = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AnalyticsEndpoint = types.StringNull()
	}
	if value := res.Get(pathRoot + `ApplicationType`); value.Exists() {
		data.ApplicationType = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.ApplicationType = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `ParseSettingsReference`); value.Exists() {
		data.ParseSettingsReference = &DmDynamicParseSettingsReference{}
		data.ParseSettingsReference.FromBody(ctx, "", value)
	} else {
		data.ParseSettingsReference = nil
	}
}

func (data *APICollection) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `Sandbox`); value.Exists() && !data.Sandbox.IsNull() {
		data.Sandbox = tfutils.BoolFromString(value.String())
	} else if data.Sandbox.ValueBool() {
		data.Sandbox = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OrgID`); value.Exists() && !data.OrgId.IsNull() {
		data.OrgId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OrgId = types.StringNull()
	}
	if value := res.Get(pathRoot + `OrgName`); value.Exists() && !data.OrgName.IsNull() {
		data.OrgName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OrgName = types.StringNull()
	}
	if value := res.Get(pathRoot + `CatalogID`); value.Exists() && !data.CatalogId.IsNull() {
		data.CatalogId = tfutils.ParseStringFromGJSON(value)
	} else if data.CatalogId.ValueString() != "default-catalog-id" {
		data.CatalogId = types.StringNull()
	}
	if value := res.Get(pathRoot + `CatalogName`); value.Exists() && !data.CatalogName.IsNull() {
		data.CatalogName = tfutils.ParseStringFromGJSON(value)
	} else if data.CatalogName.ValueString() != "default" {
		data.CatalogName = types.StringNull()
	}
	if value := res.Get(pathRoot + `DevPortalEndpoint`); value.Exists() && !data.DevPortalEndpoint.IsNull() {
		data.DevPortalEndpoint = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DevPortalEndpoint = types.StringNull()
	}
	if value := res.Get(pathRoot + `CacheCapacity`); value.Exists() && !data.CacheCapacity.IsNull() {
		data.CacheCapacity = types.Int64Value(value.Int())
	} else if data.CacheCapacity.ValueInt64() != 128 {
		data.CacheCapacity = types.Int64Null()
	}
	if value := res.Get(pathRoot + `RoutingPrefix`); value.Exists() && !data.RoutingPrefix.IsNull() {
		l := []DmRoutingPrefix{}
		e := []DmRoutingPrefix{}
		data.RoutingPrefix.ElementsAs(ctx, &e, false)
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
				item := DmRoutingPrefix{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.RoutingPrefix, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmRoutingPrefixObjectType}, l)
		} else {
			data.RoutingPrefix = types.ListNull(types.ObjectType{AttrTypes: DmRoutingPrefixObjectType})
		}
	} else {
		data.RoutingPrefix = types.ListNull(types.ObjectType{AttrTypes: DmRoutingPrefixObjectType})
	}
	if value := res.Get(pathRoot + `UseRateLimitGroup`); value.Exists() && !data.UseRateLimitGroup.IsNull() {
		data.UseRateLimitGroup = tfutils.BoolFromString(value.String())
	} else if data.UseRateLimitGroup.ValueBool() {
		data.UseRateLimitGroup = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DefaultRateLimit`); value.Exists() && !data.DefaultRateLimit.IsNull() {
		l := []DmAPIRateLimit{}
		e := []DmAPIRateLimit{}
		data.DefaultRateLimit.ElementsAs(ctx, &e, false)
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
				item := DmAPIRateLimit{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.DefaultRateLimit, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPIRateLimitObjectType}, l)
		} else {
			data.DefaultRateLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIRateLimitObjectType})
		}
	} else {
		data.DefaultRateLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIRateLimitObjectType})
	}
	if value := res.Get(pathRoot + `RateLimitGroup`); value.Exists() && !data.RateLimitGroup.IsNull() {
		data.RateLimitGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RateLimitGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `AssemblyBurstLimit`); value.Exists() && !data.AssemblyBurstLimit.IsNull() {
		l := []DmAPIBurstLimit{}
		e := []DmAPIBurstLimit{}
		data.AssemblyBurstLimit.ElementsAs(ctx, &e, false)
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
				item := DmAPIBurstLimit{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.AssemblyBurstLimit, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPIBurstLimitObjectType}, l)
		} else {
			data.AssemblyBurstLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIBurstLimitObjectType})
		}
	} else {
		data.AssemblyBurstLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIBurstLimitObjectType})
	}
	if value := res.Get(pathRoot + `AssemblyRateLimit`); value.Exists() && !data.AssemblyRateLimit.IsNull() {
		l := []DmAPIRateLimit{}
		e := []DmAPIRateLimit{}
		data.AssemblyRateLimit.ElementsAs(ctx, &e, false)
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
				item := DmAPIRateLimit{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.AssemblyRateLimit, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPIRateLimitObjectType}, l)
		} else {
			data.AssemblyRateLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIRateLimitObjectType})
		}
	} else {
		data.AssemblyRateLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIRateLimitObjectType})
	}
	if value := res.Get(pathRoot + `AssemblyCountLimit`); value.Exists() && !data.AssemblyCountLimit.IsNull() {
		l := []DmAPICountLimit{}
		e := []DmAPICountLimit{}
		data.AssemblyCountLimit.ElementsAs(ctx, &e, false)
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
				item := DmAPICountLimit{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.AssemblyCountLimit, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPICountLimitObjectType}, l)
		} else {
			data.AssemblyCountLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPICountLimitObjectType})
		}
	} else {
		data.AssemblyCountLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPICountLimitObjectType})
	}
	if value := res.Get(pathRoot + `EnforcePreAssemblyRateLimits`); value.Exists() && !data.EnforcePreAssemblyRateLimits.IsNull() {
		data.EnforcePreAssemblyRateLimits = tfutils.BoolFromString(value.String())
	} else if !data.EnforcePreAssemblyRateLimits.ValueBool() {
		data.EnforcePreAssemblyRateLimits = types.BoolNull()
	}
	if value := res.Get(pathRoot + `APIProcessingRule`); value.Exists() && !data.ApiProcessingRule.IsNull() {
		data.ApiProcessingRule = tfutils.ParseStringFromGJSON(value)
	} else if data.ApiProcessingRule.ValueString() != "default-api-rule" {
		data.ApiProcessingRule = types.StringNull()
	}
	if value := res.Get(pathRoot + `APIErrorRule`); value.Exists() && !data.ApiErrorRule.IsNull() {
		data.ApiErrorRule = tfutils.ParseStringFromGJSON(value)
	} else if data.ApiErrorRule.ValueString() != "default-api-error-rule" {
		data.ApiErrorRule = types.StringNull()
	}
	if value := res.Get(pathRoot + `AssemblyPreflow`); value.Exists() && !data.AssemblyPreflow.IsNull() {
		data.AssemblyPreflow = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AssemblyPreflow = types.StringNull()
	}
	if value := res.Get(pathRoot + `AssemblyPostflow`); value.Exists() && !data.AssemblyPostflow.IsNull() {
		data.AssemblyPostflow = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AssemblyPostflow = types.StringNull()
	}
	if value := res.Get(pathRoot + `Plan`); value.Exists() && !data.Plan.IsNull() {
		data.Plan = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Plan = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `AnalyticsEndpoint`); value.Exists() && !data.AnalyticsEndpoint.IsNull() {
		data.AnalyticsEndpoint = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AnalyticsEndpoint = types.StringNull()
	}
	if value := res.Get(pathRoot + `ApplicationType`); value.Exists() && !data.ApplicationType.IsNull() {
		data.ApplicationType = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.ApplicationType = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `ParseSettingsReference`); value.Exists() {
		data.ParseSettingsReference.UpdateFromBody(ctx, "", value)
	} else {
		data.ParseSettingsReference = nil
	}
}
