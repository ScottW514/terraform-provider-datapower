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

type URLRewritePolicy struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	Direction         types.String                `tfsdk:"direction"`
	UrlRewriteRule    types.List                  `tfsdk:"url_rewrite_rule"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var URLRewritePolicyObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"direction":          types.StringType,
	"url_rewrite_rule":   types.ListType{ElemType: types.ObjectType{AttrTypes: DmURLRewriteRuleObjectType}},
	"dependency_actions": actions.ActionsListType,
}

func (data URLRewritePolicy) GetPath() string {
	rest_path := "/mgmt/config/{domain}/URLRewritePolicy"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data URLRewritePolicy) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Direction.IsNull() {
		return false
	}
	if !data.UrlRewriteRule.IsNull() {
		return false
	}
	return true
}

func (data URLRewritePolicy) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.Direction.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Direction`, data.Direction.ValueString())
	}
	if !data.UrlRewriteRule.IsNull() {
		var dataValues []DmURLRewriteRule
		data.UrlRewriteRule.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`URLRewriteRule`+".-1", val.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *URLRewritePolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Direction`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Direction = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Direction = types.StringValue("all")
	}
	if value := res.Get(pathRoot + `URLRewriteRule`); value.Exists() {
		l := []DmURLRewriteRule{}
		if value := res.Get(`URLRewriteRule`); value.Exists() {
			for _, v := range value.Array() {
				item := DmURLRewriteRule{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.UrlRewriteRule, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmURLRewriteRuleObjectType}, l)
		} else {
			data.UrlRewriteRule = types.ListNull(types.ObjectType{AttrTypes: DmURLRewriteRuleObjectType})
		}
	} else {
		data.UrlRewriteRule = types.ListNull(types.ObjectType{AttrTypes: DmURLRewriteRuleObjectType})
	}
}

func (data *URLRewritePolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Direction`); value.Exists() && !data.Direction.IsNull() {
		data.Direction = tfutils.ParseStringFromGJSON(value)
	} else if data.Direction.ValueString() != "all" {
		data.Direction = types.StringNull()
	}
	if value := res.Get(pathRoot + `URLRewriteRule`); value.Exists() && !data.UrlRewriteRule.IsNull() {
		l := []DmURLRewriteRule{}
		e := []DmURLRewriteRule{}
		data.UrlRewriteRule.ElementsAs(ctx, &e, false)
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
				item := DmURLRewriteRule{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.UrlRewriteRule, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmURLRewriteRuleObjectType}, l)
		} else {
			data.UrlRewriteRule = types.ListNull(types.ObjectType{AttrTypes: DmURLRewriteRuleObjectType})
		}
	} else {
		data.UrlRewriteRule = types.ListNull(types.ObjectType{AttrTypes: DmURLRewriteRuleObjectType})
	}
}
