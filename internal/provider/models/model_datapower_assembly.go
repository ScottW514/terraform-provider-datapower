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

type Assembly struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	Rule              types.String                `tfsdk:"rule"`
	Catch             types.List                  `tfsdk:"catch"`
	DefaultCatch      types.String                `tfsdk:"default_catch"`
	Finally           types.String                `tfsdk:"finally"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var AssemblyObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"user_summary":       types.StringType,
	"rule":               types.StringType,
	"catch":              types.ListType{ElemType: types.ObjectType{AttrTypes: DmAssemblyCatchObjectType}},
	"default_catch":      types.StringType,
	"finally":            types.StringType,
	"dependency_actions": actions.ActionsListType,
}

func (data Assembly) GetPath() string {
	rest_path := "/mgmt/config/{domain}/Assembly"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data Assembly) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Rule.IsNull() {
		return false
	}
	if !data.Catch.IsNull() {
		return false
	}
	if !data.DefaultCatch.IsNull() {
		return false
	}
	if !data.Finally.IsNull() {
		return false
	}
	return true
}

func (data Assembly) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Rule.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Rule`, data.Rule.ValueString())
	}
	if !data.Catch.IsNull() {
		var dataValues []DmAssemblyCatch
		data.Catch.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`Catch`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.DefaultCatch.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DefaultCatch`, data.DefaultCatch.ValueString())
	}
	if !data.Finally.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Finally`, data.Finally.ValueString())
	}
	return body
}

func (data *Assembly) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Rule`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Rule = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Rule = types.StringNull()
	}
	if value := res.Get(pathRoot + `Catch`); value.Exists() {
		l := []DmAssemblyCatch{}
		if value := res.Get(`Catch`); value.Exists() {
			for _, v := range value.Array() {
				item := DmAssemblyCatch{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Catch, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAssemblyCatchObjectType}, l)
		} else {
			data.Catch = types.ListNull(types.ObjectType{AttrTypes: DmAssemblyCatchObjectType})
		}
	} else {
		data.Catch = types.ListNull(types.ObjectType{AttrTypes: DmAssemblyCatchObjectType})
	}
	if value := res.Get(pathRoot + `DefaultCatch`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DefaultCatch = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DefaultCatch = types.StringNull()
	}
	if value := res.Get(pathRoot + `Finally`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Finally = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Finally = types.StringNull()
	}
}

func (data *Assembly) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Rule`); value.Exists() && !data.Rule.IsNull() {
		data.Rule = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Rule = types.StringNull()
	}
	if value := res.Get(pathRoot + `Catch`); value.Exists() && !data.Catch.IsNull() {
		l := []DmAssemblyCatch{}
		for _, v := range value.Array() {
			item := DmAssemblyCatch{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.Catch, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAssemblyCatchObjectType}, l)
		} else {
			data.Catch = types.ListNull(types.ObjectType{AttrTypes: DmAssemblyCatchObjectType})
		}
	} else {
		data.Catch = types.ListNull(types.ObjectType{AttrTypes: DmAssemblyCatchObjectType})
	}
	if value := res.Get(pathRoot + `DefaultCatch`); value.Exists() && !data.DefaultCatch.IsNull() {
		data.DefaultCatch = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DefaultCatch = types.StringNull()
	}
	if value := res.Get(pathRoot + `Finally`); value.Exists() && !data.Finally.IsNull() {
		data.Finally = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Finally = types.StringNull()
	}
}
