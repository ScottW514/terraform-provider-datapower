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

type RateLimitDefinitionGroup struct {
	Id                   types.String                `tfsdk:"id"`
	AppDomain            types.String                `tfsdk:"app_domain"`
	UserSummary          types.String                `tfsdk:"user_summary"`
	UpdateOnExceed       types.String                `tfsdk:"update_on_exceed"`
	RateLimitDefinitions types.List                  `tfsdk:"rate_limit_definitions"`
	DependencyActions    []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var RateLimitDefinitionGroupObjectType = map[string]attr.Type{
	"id":                     types.StringType,
	"app_domain":             types.StringType,
	"user_summary":           types.StringType,
	"update_on_exceed":       types.StringType,
	"rate_limit_definitions": types.ListType{ElemType: types.StringType},
	"dependency_actions":     actions.ActionsListType,
}

func (data RateLimitDefinitionGroup) GetPath() string {
	rest_path := "/mgmt/config/{domain}/RateLimitDefinitionGroup"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data RateLimitDefinitionGroup) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.UpdateOnExceed.IsNull() {
		return false
	}
	if !data.RateLimitDefinitions.IsNull() {
		return false
	}
	return true
}

func (data RateLimitDefinitionGroup) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.UpdateOnExceed.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UpdateOnExceed`, data.UpdateOnExceed.ValueString())
	}
	if !data.RateLimitDefinitions.IsNull() {
		var dataValues []string
		data.RateLimitDefinitions.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`RateLimitDefinitions`+".-1", map[string]string{"value": val})
		}
	}
	return body
}

func (data *RateLimitDefinitionGroup) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `UpdateOnExceed`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UpdateOnExceed = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UpdateOnExceed = types.StringValue("all")
	}
	if value := res.Get(pathRoot + `RateLimitDefinitions`); value.Exists() {
		data.RateLimitDefinitions = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.RateLimitDefinitions = types.ListNull(types.StringType)
	}
}

func (data *RateLimitDefinitionGroup) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `UpdateOnExceed`); value.Exists() && !data.UpdateOnExceed.IsNull() {
		data.UpdateOnExceed = tfutils.ParseStringFromGJSON(value)
	} else if data.UpdateOnExceed.ValueString() != "all" {
		data.UpdateOnExceed = types.StringNull()
	}
	if value := res.Get(pathRoot + `RateLimitDefinitions`); value.Exists() && !data.RateLimitDefinitions.IsNull() {
		data.RateLimitDefinitions = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.RateLimitDefinitions = types.ListNull(types.StringType)
	}
}
