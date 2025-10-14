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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type UserGroup struct {
	Id                types.String                `tfsdk:"id"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	AccessPolicies    types.List                  `tfsdk:"access_policies"`
	CommandGroup      types.List                  `tfsdk:"command_group"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var UserGroupObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"user_summary":       types.StringType,
	"access_policies":    types.ListType{ElemType: types.StringType},
	"command_group":      types.ListType{ElemType: types.StringType},
	"dependency_actions": actions.ActionsListType,
}

func (data UserGroup) GetPath() string {
	rest_path := "/mgmt/config/default/UserGroup"
	return rest_path
}

func (data UserGroup) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.AccessPolicies.IsNull() {
		return false
	}
	if !data.CommandGroup.IsNull() {
		return false
	}
	return true
}

func (data UserGroup) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.AccessPolicies.IsNull() {
		var dataValues []string
		data.AccessPolicies.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.Set(body, pathRoot+`AccessPolicies`+".-1", map[string]string{"value": val})
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`AccessPolicies`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`AccessPolicies`, "[]")
	}
	if !data.CommandGroup.IsNull() {
		var dataValues []string
		data.CommandGroup.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.Set(body, pathRoot+`CommandGroup`+".-1", map[string]string{"value": val})
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`CommandGroup`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`CommandGroup`, "[]")
	}
	return body
}

func (data *UserGroup) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `AccessPolicies`); value.Exists() {
		data.AccessPolicies = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.AccessPolicies = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `CommandGroup`); value.Exists() {
		data.CommandGroup = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.CommandGroup = types.ListNull(types.StringType)
	}
}

func (data *UserGroup) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `AccessPolicies`); value.Exists() && !data.AccessPolicies.IsNull() {
		data.AccessPolicies = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.AccessPolicies = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `CommandGroup`); value.Exists() && !data.CommandGroup.IsNull() {
		data.CommandGroup = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.CommandGroup = types.ListNull(types.StringType)
	}
}
