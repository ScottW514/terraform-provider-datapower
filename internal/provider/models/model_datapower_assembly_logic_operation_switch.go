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

type AssemblyLogicOperationSwitch struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	Case              types.List                  `tfsdk:"case"`
	Otherwise         types.String                `tfsdk:"otherwise"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	Title             types.String                `tfsdk:"title"`
	CorrelationPath   types.String                `tfsdk:"correlation_path"`
	ActionDebug       types.Bool                  `tfsdk:"action_debug"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var AssemblyLogicOperationSwitchObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"case":               types.ListType{ElemType: types.ObjectType{AttrTypes: DmAssemblyLogicOperationSwitchCaseObjectType}},
	"otherwise":          types.StringType,
	"user_summary":       types.StringType,
	"title":              types.StringType,
	"correlation_path":   types.StringType,
	"action_debug":       types.BoolType,
	"dependency_actions": actions.ActionsListType,
}

func (data AssemblyLogicOperationSwitch) GetPath() string {
	rest_path := "/mgmt/config/{domain}/AssemblyLogicOperationSwitch"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data AssemblyLogicOperationSwitch) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Case.IsNull() {
		return false
	}
	if !data.Otherwise.IsNull() {
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

func (data AssemblyLogicOperationSwitch) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.Case.IsNull() {
		var dataValues []DmAssemblyLogicOperationSwitchCase
		data.Case.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.SetRaw(body, pathRoot+`Case`+".-1", val.ToBody(ctx, ""))
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`Case`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`Case`, "[]")
	}
	if !data.Otherwise.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Otherwise`, data.Otherwise.ValueString())
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

func (data *AssemblyLogicOperationSwitch) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Case`); value.Exists() {
		l := []DmAssemblyLogicOperationSwitchCase{}
		if value := res.Get(`Case`); value.Exists() {
			for _, v := range value.Array() {
				item := DmAssemblyLogicOperationSwitchCase{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Case, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAssemblyLogicOperationSwitchCaseObjectType}, l)
		} else {
			data.Case = types.ListNull(types.ObjectType{AttrTypes: DmAssemblyLogicOperationSwitchCaseObjectType})
		}
	} else {
		data.Case = types.ListNull(types.ObjectType{AttrTypes: DmAssemblyLogicOperationSwitchCaseObjectType})
	}
	if value := res.Get(pathRoot + `Otherwise`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Otherwise = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Otherwise = types.StringNull()
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

func (data *AssemblyLogicOperationSwitch) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Case`); value.Exists() && !data.Case.IsNull() {
		l := []DmAssemblyLogicOperationSwitchCase{}
		e := []DmAssemblyLogicOperationSwitchCase{}
		data.Case.ElementsAs(ctx, &e, false)
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
				item := DmAssemblyLogicOperationSwitchCase{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Case, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAssemblyLogicOperationSwitchCaseObjectType}, l)
		} else {
			data.Case = types.ListNull(types.ObjectType{AttrTypes: DmAssemblyLogicOperationSwitchCaseObjectType})
		}
	} else {
		data.Case = types.ListNull(types.ObjectType{AttrTypes: DmAssemblyLogicOperationSwitchCaseObjectType})
	}
	if value := res.Get(pathRoot + `Otherwise`); value.Exists() && !data.Otherwise.IsNull() {
		data.Otherwise = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Otherwise = types.StringNull()
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
