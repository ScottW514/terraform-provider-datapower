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
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type AssemblyActionRateLimit struct {
	Id                  types.String                `tfsdk:"id"`
	AppDomain           types.String                `tfsdk:"app_domain"`
	Source              types.String                `tfsdk:"source"`
	BurstLimit          types.List                  `tfsdk:"burst_limit"`
	RateLimit           types.List                  `tfsdk:"rate_limit"`
	CountLimit          types.List                  `tfsdk:"count_limit"`
	RateLimitDefinition types.List                  `tfsdk:"rate_limit_definition"`
	RateLimitGroup      types.String                `tfsdk:"rate_limit_group"`
	GroupAction         types.String                `tfsdk:"group_action"`
	UserSummary         types.String                `tfsdk:"user_summary"`
	Title               types.String                `tfsdk:"title"`
	CorrelationPath     types.String                `tfsdk:"correlation_path"`
	ActionDebug         types.Bool                  `tfsdk:"action_debug"`
	DependencyActions   []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var AssemblyActionRateLimitObjectType = map[string]attr.Type{
	"id":                    types.StringType,
	"app_domain":            types.StringType,
	"source":                types.StringType,
	"burst_limit":           types.ListType{ElemType: types.StringType},
	"rate_limit":            types.ListType{ElemType: types.ObjectType{AttrTypes: DmRateLimitInfoObjectType}},
	"count_limit":           types.ListType{ElemType: types.ObjectType{AttrTypes: DmCountLimitInfoObjectType}},
	"rate_limit_definition": types.ListType{ElemType: types.ObjectType{AttrTypes: DmRateLimitInfoDomainNamedObjectType}},
	"rate_limit_group":      types.StringType,
	"group_action":          types.StringType,
	"user_summary":          types.StringType,
	"title":                 types.StringType,
	"correlation_path":      types.StringType,
	"action_debug":          types.BoolType,
	"dependency_actions":    actions.ActionsListType,
}

func (data AssemblyActionRateLimit) GetPath() string {
	rest_path := "/mgmt/config/{domain}/AssemblyActionRateLimit"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data AssemblyActionRateLimit) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Source.IsNull() {
		return false
	}
	if !data.BurstLimit.IsNull() {
		return false
	}
	if !data.RateLimit.IsNull() {
		return false
	}
	if !data.CountLimit.IsNull() {
		return false
	}
	if !data.RateLimitDefinition.IsNull() {
		return false
	}
	if !data.RateLimitGroup.IsNull() {
		return false
	}
	if !data.GroupAction.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Title.IsNull() {
		return false
	}
	if !data.CorrelationPath.IsNull() {
		return false
	}
	if !data.ActionDebug.IsNull() {
		return false
	}
	return true
}

func (data AssemblyActionRateLimit) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.Source.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Source`, data.Source.ValueString())
	}
	if !data.BurstLimit.IsNull() {
		var dataValues []string
		data.BurstLimit.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`BurstLimit`+".-1", map[string]string{"value": val})
		}
	}
	if !data.RateLimit.IsNull() {
		var dataValues []DmRateLimitInfo
		data.RateLimit.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`RateLimit`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.CountLimit.IsNull() {
		var dataValues []DmCountLimitInfo
		data.CountLimit.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`CountLimit`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.RateLimitDefinition.IsNull() {
		var dataValues []DmRateLimitInfoDomainNamed
		data.RateLimitDefinition.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`RateLimitDefinition`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.RateLimitGroup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RateLimitGroup`, data.RateLimitGroup.ValueString())
	}
	if !data.GroupAction.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GroupAction`, data.GroupAction.ValueString())
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.Title.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Title`, data.Title.ValueString())
	}
	if !data.CorrelationPath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CorrelationPath`, data.CorrelationPath.ValueString())
	}
	if !data.ActionDebug.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ActionDebug`, tfutils.StringFromBool(data.ActionDebug, ""))
	}
	return body
}

func (data *AssemblyActionRateLimit) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Source`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Source = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Source = types.StringValue("plan-default")
	}
	if value := res.Get(pathRoot + `BurstLimit`); value.Exists() {
		data.BurstLimit = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.BurstLimit = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `RateLimit`); value.Exists() {
		l := []DmRateLimitInfo{}
		if value := res.Get(`RateLimit`); value.Exists() {
			for _, v := range value.Array() {
				item := DmRateLimitInfo{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.RateLimit, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmRateLimitInfoObjectType}, l)
		} else {
			data.RateLimit = types.ListNull(types.ObjectType{AttrTypes: DmRateLimitInfoObjectType})
		}
	} else {
		data.RateLimit = types.ListNull(types.ObjectType{AttrTypes: DmRateLimitInfoObjectType})
	}
	if value := res.Get(pathRoot + `CountLimit`); value.Exists() {
		l := []DmCountLimitInfo{}
		if value := res.Get(`CountLimit`); value.Exists() {
			for _, v := range value.Array() {
				item := DmCountLimitInfo{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.CountLimit, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmCountLimitInfoObjectType}, l)
		} else {
			data.CountLimit = types.ListNull(types.ObjectType{AttrTypes: DmCountLimitInfoObjectType})
		}
	} else {
		data.CountLimit = types.ListNull(types.ObjectType{AttrTypes: DmCountLimitInfoObjectType})
	}
	if value := res.Get(pathRoot + `RateLimitDefinition`); value.Exists() {
		l := []DmRateLimitInfoDomainNamed{}
		if value := res.Get(`RateLimitDefinition`); value.Exists() {
			for _, v := range value.Array() {
				item := DmRateLimitInfoDomainNamed{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.RateLimitDefinition, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmRateLimitInfoDomainNamedObjectType}, l)
		} else {
			data.RateLimitDefinition = types.ListNull(types.ObjectType{AttrTypes: DmRateLimitInfoDomainNamedObjectType})
		}
	} else {
		data.RateLimitDefinition = types.ListNull(types.ObjectType{AttrTypes: DmRateLimitInfoDomainNamedObjectType})
	}
	if value := res.Get(pathRoot + `RateLimitGroup`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RateLimitGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RateLimitGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `GroupAction`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.GroupAction = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GroupAction = types.StringValue("consume")
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `Title`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Title = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Title = types.StringNull()
	}
	if value := res.Get(pathRoot + `CorrelationPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CorrelationPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CorrelationPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `ActionDebug`); value.Exists() {
		data.ActionDebug = tfutils.BoolFromString(value.String())
	} else {
		data.ActionDebug = types.BoolNull()
	}
}

func (data *AssemblyActionRateLimit) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Source`); value.Exists() && !data.Source.IsNull() {
		data.Source = tfutils.ParseStringFromGJSON(value)
	} else if data.Source.ValueString() != "plan-default" {
		data.Source = types.StringNull()
	}
	if value := res.Get(pathRoot + `BurstLimit`); value.Exists() && !data.BurstLimit.IsNull() {
		data.BurstLimit = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.BurstLimit = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `RateLimit`); value.Exists() && !data.RateLimit.IsNull() {
		l := []DmRateLimitInfo{}
		for _, v := range value.Array() {
			item := DmRateLimitInfo{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.RateLimit, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmRateLimitInfoObjectType}, l)
		} else {
			data.RateLimit = types.ListNull(types.ObjectType{AttrTypes: DmRateLimitInfoObjectType})
		}
	} else {
		data.RateLimit = types.ListNull(types.ObjectType{AttrTypes: DmRateLimitInfoObjectType})
	}
	if value := res.Get(pathRoot + `CountLimit`); value.Exists() && !data.CountLimit.IsNull() {
		l := []DmCountLimitInfo{}
		for _, v := range value.Array() {
			item := DmCountLimitInfo{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.CountLimit, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmCountLimitInfoObjectType}, l)
		} else {
			data.CountLimit = types.ListNull(types.ObjectType{AttrTypes: DmCountLimitInfoObjectType})
		}
	} else {
		data.CountLimit = types.ListNull(types.ObjectType{AttrTypes: DmCountLimitInfoObjectType})
	}
	if value := res.Get(pathRoot + `RateLimitDefinition`); value.Exists() && !data.RateLimitDefinition.IsNull() {
		l := []DmRateLimitInfoDomainNamed{}
		for _, v := range value.Array() {
			item := DmRateLimitInfoDomainNamed{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.RateLimitDefinition, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmRateLimitInfoDomainNamedObjectType}, l)
		} else {
			data.RateLimitDefinition = types.ListNull(types.ObjectType{AttrTypes: DmRateLimitInfoDomainNamedObjectType})
		}
	} else {
		data.RateLimitDefinition = types.ListNull(types.ObjectType{AttrTypes: DmRateLimitInfoDomainNamedObjectType})
	}
	if value := res.Get(pathRoot + `RateLimitGroup`); value.Exists() && !data.RateLimitGroup.IsNull() {
		data.RateLimitGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RateLimitGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `GroupAction`); value.Exists() && !data.GroupAction.IsNull() {
		data.GroupAction = tfutils.ParseStringFromGJSON(value)
	} else if data.GroupAction.ValueString() != "consume" {
		data.GroupAction = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `Title`); value.Exists() && !data.Title.IsNull() {
		data.Title = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Title = types.StringNull()
	}
	if value := res.Get(pathRoot + `CorrelationPath`); value.Exists() && !data.CorrelationPath.IsNull() {
		data.CorrelationPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CorrelationPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `ActionDebug`); value.Exists() && !data.ActionDebug.IsNull() {
		data.ActionDebug = tfutils.BoolFromString(value.String())
	} else if data.ActionDebug.ValueBool() {
		data.ActionDebug = types.BoolNull()
	}
}
