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

type Matching struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	MatchRules        types.List                  `tfsdk:"match_rules"`
	MatchWithPcre     types.Bool                  `tfsdk:"match_with_pcre"`
	CombineWithOr     types.Bool                  `tfsdk:"combine_with_or"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var MatchingObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"user_summary":       types.StringType,
	"match_rules":        types.ListType{ElemType: types.ObjectType{AttrTypes: DmMatchRuleObjectType}},
	"match_with_pcre":    types.BoolType,
	"combine_with_or":    types.BoolType,
	"dependency_actions": actions.ActionsListType,
}

func (data Matching) GetPath() string {
	rest_path := "/mgmt/config/{domain}/Matching"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data Matching) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.MatchRules.IsNull() {
		return false
	}
	if !data.MatchWithPcre.IsNull() {
		return false
	}
	if !data.CombineWithOr.IsNull() {
		return false
	}
	return true
}

func (data Matching) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.MatchRules.IsNull() {
		var values []DmMatchRule
		data.MatchRules.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`MatchRules`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.MatchWithPcre.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MatchWithPCRE`, tfutils.StringFromBool(data.MatchWithPcre, ""))
	}
	if !data.CombineWithOr.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CombineWithOr`, tfutils.StringFromBool(data.CombineWithOr, ""))
	}
	return body
}

func (data *Matching) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `MatchRules`); value.Exists() {
		l := []DmMatchRule{}
		if value := res.Get(`MatchRules`); value.Exists() {
			for _, v := range value.Array() {
				item := DmMatchRule{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.MatchRules, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmMatchRuleObjectType}, l)
		} else {
			data.MatchRules = types.ListNull(types.ObjectType{AttrTypes: DmMatchRuleObjectType})
		}
	} else {
		data.MatchRules = types.ListNull(types.ObjectType{AttrTypes: DmMatchRuleObjectType})
	}
	if value := res.Get(pathRoot + `MatchWithPCRE`); value.Exists() {
		data.MatchWithPcre = tfutils.BoolFromString(value.String())
	} else {
		data.MatchWithPcre = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CombineWithOr`); value.Exists() {
		data.CombineWithOr = tfutils.BoolFromString(value.String())
	} else {
		data.CombineWithOr = types.BoolNull()
	}
}

func (data *Matching) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `MatchRules`); value.Exists() && !data.MatchRules.IsNull() {
		l := []DmMatchRule{}
		for _, v := range value.Array() {
			item := DmMatchRule{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.MatchRules, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmMatchRuleObjectType}, l)
		} else {
			data.MatchRules = types.ListNull(types.ObjectType{AttrTypes: DmMatchRuleObjectType})
		}
	} else {
		data.MatchRules = types.ListNull(types.ObjectType{AttrTypes: DmMatchRuleObjectType})
	}
	if value := res.Get(pathRoot + `MatchWithPCRE`); value.Exists() && !data.MatchWithPcre.IsNull() {
		data.MatchWithPcre = tfutils.BoolFromString(value.String())
	} else if data.MatchWithPcre.ValueBool() {
		data.MatchWithPcre = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CombineWithOr`); value.Exists() && !data.CombineWithOr.IsNull() {
		data.CombineWithOr = tfutils.BoolFromString(value.String())
	} else if data.CombineWithOr.ValueBool() {
		data.CombineWithOr = types.BoolNull()
	}
}
