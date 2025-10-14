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

type AssemblyActionGraphQLExecute struct {
	Id                     types.String                `tfsdk:"id"`
	AppDomain              types.String                `tfsdk:"app_domain"`
	Description            types.String                `tfsdk:"description"`
	Input                  types.String                `tfsdk:"input"`
	Output                 types.String                `tfsdk:"output"`
	TargetMapRule          types.List                  `tfsdk:"target_map_rule"`
	AllowCostIntrospection types.Bool                  `tfsdk:"allow_cost_introspection"`
	UserSummary            types.String                `tfsdk:"user_summary"`
	Title                  types.String                `tfsdk:"title"`
	CorrelationPath        types.String                `tfsdk:"correlation_path"`
	ActionDebug            types.Bool                  `tfsdk:"action_debug"`
	DependencyActions      []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var AssemblyActionGraphQLExecuteObjectType = map[string]attr.Type{
	"id":                       types.StringType,
	"app_domain":               types.StringType,
	"description":              types.StringType,
	"input":                    types.StringType,
	"output":                   types.StringType,
	"target_map_rule":          types.ListType{ElemType: types.ObjectType{AttrTypes: DmTargetMapRuleObjectType}},
	"allow_cost_introspection": types.BoolType,
	"user_summary":             types.StringType,
	"title":                    types.StringType,
	"correlation_path":         types.StringType,
	"action_debug":             types.BoolType,
	"dependency_actions":       actions.ActionsListType,
}

func (data AssemblyActionGraphQLExecute) GetPath() string {
	rest_path := "/mgmt/config/{domain}/AssemblyActionGraphQLExecute"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data AssemblyActionGraphQLExecute) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Description.IsNull() {
		return false
	}
	if !data.Input.IsNull() {
		return false
	}
	if !data.Output.IsNull() {
		return false
	}
	if !data.TargetMapRule.IsNull() {
		return false
	}
	if !data.AllowCostIntrospection.IsNull() {
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

func (data AssemblyActionGraphQLExecute) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Description`, data.Description.ValueString())
	}
	if !data.Input.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Input`, data.Input.ValueString())
	}
	if !data.Output.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Output`, data.Output.ValueString())
	}
	if !data.TargetMapRule.IsNull() {
		var dataValues []DmTargetMapRule
		data.TargetMapRule.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.SetRaw(body, pathRoot+`TargetMapRule`+".-1", val.ToBody(ctx, ""))
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`TargetMapRule`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`TargetMapRule`, "[]")
	}
	if !data.AllowCostIntrospection.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllowCostIntrospection`, tfutils.StringFromBool(data.AllowCostIntrospection, ""))
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

func (data *AssemblyActionGraphQLExecute) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Description`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Description = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get(pathRoot + `Input`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Input = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Input = types.StringValue("message")
	}
	if value := res.Get(pathRoot + `Output`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Output = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Output = types.StringNull()
	}
	if value := res.Get(pathRoot + `TargetMapRule`); value.Exists() {
		l := []DmTargetMapRule{}
		if value := res.Get(`TargetMapRule`); value.Exists() {
			for _, v := range value.Array() {
				item := DmTargetMapRule{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.TargetMapRule, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmTargetMapRuleObjectType}, l)
		} else {
			data.TargetMapRule = types.ListNull(types.ObjectType{AttrTypes: DmTargetMapRuleObjectType})
		}
	} else {
		data.TargetMapRule = types.ListNull(types.ObjectType{AttrTypes: DmTargetMapRuleObjectType})
	}
	if value := res.Get(pathRoot + `AllowCostIntrospection`); value.Exists() {
		data.AllowCostIntrospection = tfutils.BoolFromString(value.String())
	} else {
		data.AllowCostIntrospection = types.BoolNull()
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

func (data *AssemblyActionGraphQLExecute) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Description`); value.Exists() && !data.Description.IsNull() {
		data.Description = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get(pathRoot + `Input`); value.Exists() && !data.Input.IsNull() {
		data.Input = tfutils.ParseStringFromGJSON(value)
	} else if data.Input.ValueString() != "message" {
		data.Input = types.StringNull()
	}
	if value := res.Get(pathRoot + `Output`); value.Exists() && !data.Output.IsNull() {
		data.Output = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Output = types.StringNull()
	}
	if value := res.Get(pathRoot + `TargetMapRule`); value.Exists() && !data.TargetMapRule.IsNull() {
		l := []DmTargetMapRule{}
		e := []DmTargetMapRule{}
		data.TargetMapRule.ElementsAs(ctx, &e, false)
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
				item := DmTargetMapRule{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.TargetMapRule, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmTargetMapRuleObjectType}, l)
		} else {
			data.TargetMapRule = types.ListNull(types.ObjectType{AttrTypes: DmTargetMapRuleObjectType})
		}
	} else {
		data.TargetMapRule = types.ListNull(types.ObjectType{AttrTypes: DmTargetMapRuleObjectType})
	}
	if value := res.Get(pathRoot + `AllowCostIntrospection`); value.Exists() && !data.AllowCostIntrospection.IsNull() {
		data.AllowCostIntrospection = tfutils.BoolFromString(value.String())
	} else if data.AllowCostIntrospection.ValueBool() {
		data.AllowCostIntrospection = types.BoolNull()
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
