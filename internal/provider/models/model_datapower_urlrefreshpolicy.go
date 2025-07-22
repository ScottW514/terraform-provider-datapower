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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type URLRefreshPolicy struct {
	Id             types.String `tfsdk:"id"`
	AppDomain      types.String `tfsdk:"app_domain"`
	UserSummary    types.String `tfsdk:"user_summary"`
	UrlRefreshRule types.List   `tfsdk:"url_refresh_rule"`
}

var URLRefreshPolicyObjectType = map[string]attr.Type{
	"id":               types.StringType,
	"app_domain":       types.StringType,
	"user_summary":     types.StringType,
	"url_refresh_rule": types.ListType{ElemType: types.ObjectType{AttrTypes: DmURLRefreshRuleObjectType}},
}

func (data URLRefreshPolicy) GetPath() string {
	rest_path := "/mgmt/config/{domain}/URLRefreshPolicy"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data URLRefreshPolicy) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.UrlRefreshRule.IsNull() {
		return false
	}
	return true
}

func (data URLRefreshPolicy) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.UrlRefreshRule.IsNull() {
		var values []DmURLRefreshRule
		data.UrlRefreshRule.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`URLRefreshRule`+".-1", val.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *URLRefreshPolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `URLRefreshRule`); value.Exists() {
		l := []DmURLRefreshRule{}
		if value := res.Get(`URLRefreshRule`); value.Exists() {
			for _, v := range value.Array() {
				item := DmURLRefreshRule{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.UrlRefreshRule, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmURLRefreshRuleObjectType}, l)
		} else {
			data.UrlRefreshRule = types.ListNull(types.ObjectType{AttrTypes: DmURLRefreshRuleObjectType})
		}
	} else {
		data.UrlRefreshRule = types.ListNull(types.ObjectType{AttrTypes: DmURLRefreshRuleObjectType})
	}
}

func (data *URLRefreshPolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `URLRefreshRule`); value.Exists() && !data.UrlRefreshRule.IsNull() {
		l := []DmURLRefreshRule{}
		for _, v := range value.Array() {
			item := DmURLRefreshRule{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.UrlRefreshRule, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmURLRefreshRuleObjectType}, l)
		} else {
			data.UrlRefreshRule = types.ListNull(types.ObjectType{AttrTypes: DmURLRefreshRuleObjectType})
		}
	} else {
		data.UrlRefreshRule = types.ListNull(types.ObjectType{AttrTypes: DmURLRefreshRuleObjectType})
	}
}
