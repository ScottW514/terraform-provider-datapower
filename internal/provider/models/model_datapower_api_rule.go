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

type APIRule struct {
	Id                 types.String                `tfsdk:"id"`
	AppDomain          types.String                `tfsdk:"app_domain"`
	DynamicActionsMode types.Bool                  `tfsdk:"dynamic_actions_mode"`
	Actions            types.List                  `tfsdk:"actions"`
	DynamicActions     types.List                  `tfsdk:"dynamic_actions"`
	UserSummary        types.String                `tfsdk:"user_summary"`
	DependencyActions  []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var APIRuleActionsIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "dynamic_actions_mode",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var APIRuleDynamicActionsCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "dynamic_actions_mode",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var APIRuleDynamicActionsIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "dynamic_actions_mode",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var APIRuleObjectType = map[string]attr.Type{
	"id":                   types.StringType,
	"app_domain":           types.StringType,
	"dynamic_actions_mode": types.BoolType,
	"actions":              types.ListType{ElemType: types.StringType},
	"dynamic_actions":      types.ListType{ElemType: types.ObjectType{AttrTypes: DmDynamicStylePolicyActionBaseReferenceObjectType}},
	"user_summary":         types.StringType,
	"dependency_actions":   actions.ActionsListType,
}

func (data APIRule) GetPath() string {
	rest_path := "/mgmt/config/{domain}/APIRule"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data APIRule) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.DynamicActionsMode.IsNull() {
		return false
	}
	if !data.Actions.IsNull() {
		return false
	}
	if !data.DynamicActions.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	return true
}

func (data APIRule) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.DynamicActionsMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DynamicActionsMode`, tfutils.StringFromBool(data.DynamicActionsMode, ""))
	}
	if !data.Actions.IsNull() {
		var dataValues []string
		data.Actions.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.Set(body, pathRoot+`Actions`+".-1", map[string]string{"value": val})
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`Actions`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`Actions`, "[]")
	}
	if !data.DynamicActions.IsNull() {
		var dataValues []DmDynamicStylePolicyActionBaseReference
		data.DynamicActions.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.SetRaw(body, pathRoot+`DynamicActions`+".-1", val.ToBody(ctx, ""))
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`DynamicActions`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`DynamicActions`, "[]")
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	return body
}

func (data *APIRule) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `DynamicActionsMode`); value.Exists() {
		data.DynamicActionsMode = tfutils.BoolFromString(value.String())
	} else {
		data.DynamicActionsMode = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Actions`); value.Exists() {
		data.Actions = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Actions = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `DynamicActions`); value.Exists() {
		l := []DmDynamicStylePolicyActionBaseReference{}
		if value := res.Get(`DynamicActions`); value.Exists() {
			for _, v := range value.Array() {
				item := DmDynamicStylePolicyActionBaseReference{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.DynamicActions, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmDynamicStylePolicyActionBaseReferenceObjectType}, l)
		} else {
			data.DynamicActions = types.ListNull(types.ObjectType{AttrTypes: DmDynamicStylePolicyActionBaseReferenceObjectType})
		}
	} else {
		data.DynamicActions = types.ListNull(types.ObjectType{AttrTypes: DmDynamicStylePolicyActionBaseReferenceObjectType})
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
}

func (data *APIRule) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `DynamicActionsMode`); value.Exists() && !data.DynamicActionsMode.IsNull() {
		data.DynamicActionsMode = tfutils.BoolFromString(value.String())
	} else if data.DynamicActionsMode.ValueBool() {
		data.DynamicActionsMode = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Actions`); value.Exists() && !data.Actions.IsNull() {
		data.Actions = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Actions = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `DynamicActions`); value.Exists() && !data.DynamicActions.IsNull() {
		l := []DmDynamicStylePolicyActionBaseReference{}
		e := []DmDynamicStylePolicyActionBaseReference{}
		data.DynamicActions.ElementsAs(ctx, &e, false)
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
				item := DmDynamicStylePolicyActionBaseReference{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.DynamicActions, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmDynamicStylePolicyActionBaseReferenceObjectType}, l)
		} else {
			data.DynamicActions = types.ListNull(types.ObjectType{AttrTypes: DmDynamicStylePolicyActionBaseReferenceObjectType})
		}
	} else {
		data.DynamicActions = types.ListNull(types.ObjectType{AttrTypes: DmDynamicStylePolicyActionBaseReferenceObjectType})
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
}
