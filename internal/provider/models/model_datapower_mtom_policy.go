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

type MTOMPolicy struct {
	Id                 types.String                `tfsdk:"id"`
	AppDomain          types.String                `tfsdk:"app_domain"`
	UserSummary        types.String                `tfsdk:"user_summary"`
	Mode               types.String                `tfsdk:"mode"`
	IncludeContentType types.Bool                  `tfsdk:"include_content_type"`
	Rule               types.List                  `tfsdk:"rule"`
	DependencyActions  []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var MTOMPolicyObjectType = map[string]attr.Type{
	"id":                   types.StringType,
	"app_domain":           types.StringType,
	"user_summary":         types.StringType,
	"mode":                 types.StringType,
	"include_content_type": types.BoolType,
	"rule":                 types.ListType{ElemType: types.ObjectType{AttrTypes: DmMtomRuleObjectType}},
	"dependency_actions":   actions.ActionsListType,
}

func (data MTOMPolicy) GetPath() string {
	rest_path := "/mgmt/config/{domain}/MTOMPolicy"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data MTOMPolicy) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Mode.IsNull() {
		return false
	}
	if !data.IncludeContentType.IsNull() {
		return false
	}
	if !data.Rule.IsNull() {
		return false
	}
	return true
}

func (data MTOMPolicy) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Mode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Mode`, data.Mode.ValueString())
	}
	if !data.IncludeContentType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IncludeContentType`, tfutils.StringFromBool(data.IncludeContentType, ""))
	}
	if !data.Rule.IsNull() {
		var dataValues []DmMtomRule
		data.Rule.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`Rule`+".-1", val.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *MTOMPolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Mode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Mode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Mode = types.StringValue("encode")
	}
	if value := res.Get(pathRoot + `IncludeContentType`); value.Exists() {
		data.IncludeContentType = tfutils.BoolFromString(value.String())
	} else {
		data.IncludeContentType = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Rule`); value.Exists() {
		l := []DmMtomRule{}
		if value := res.Get(`Rule`); value.Exists() {
			for _, v := range value.Array() {
				item := DmMtomRule{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Rule, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmMtomRuleObjectType}, l)
		} else {
			data.Rule = types.ListNull(types.ObjectType{AttrTypes: DmMtomRuleObjectType})
		}
	} else {
		data.Rule = types.ListNull(types.ObjectType{AttrTypes: DmMtomRuleObjectType})
	}
}

func (data *MTOMPolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Mode`); value.Exists() && !data.Mode.IsNull() {
		data.Mode = tfutils.ParseStringFromGJSON(value)
	} else if data.Mode.ValueString() != "encode" {
		data.Mode = types.StringNull()
	}
	if value := res.Get(pathRoot + `IncludeContentType`); value.Exists() && !data.IncludeContentType.IsNull() {
		data.IncludeContentType = tfutils.BoolFromString(value.String())
	} else if !data.IncludeContentType.ValueBool() {
		data.IncludeContentType = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Rule`); value.Exists() && !data.Rule.IsNull() {
		l := []DmMtomRule{}
		for _, v := range value.Array() {
			item := DmMtomRule{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.Rule, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmMtomRuleObjectType}, l)
		} else {
			data.Rule = types.ListNull(types.ObjectType{AttrTypes: DmMtomRuleObjectType})
		}
	} else {
		data.Rule = types.ListNull(types.ObjectType{AttrTypes: DmMtomRuleObjectType})
	}
}
