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

type AssemblyActionExtract struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	Root              types.String                `tfsdk:"root"`
	Extract           types.List                  `tfsdk:"extract"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	Title             types.String                `tfsdk:"title"`
	CorrelationPath   types.String                `tfsdk:"correlation_path"`
	ActionDebug       types.Bool                  `tfsdk:"action_debug"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var AssemblyActionExtractObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"root":               types.StringType,
	"extract":            types.ListType{ElemType: types.ObjectType{AttrTypes: DmAssemblyActionExtractObjectType}},
	"user_summary":       types.StringType,
	"title":              types.StringType,
	"correlation_path":   types.StringType,
	"action_debug":       types.BoolType,
	"dependency_actions": actions.ActionsListType,
}

func (data AssemblyActionExtract) GetPath() string {
	rest_path := "/mgmt/config/{domain}/AssemblyActionExtract"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data AssemblyActionExtract) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Root.IsNull() {
		return false
	}
	if !data.Extract.IsNull() {
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

func (data AssemblyActionExtract) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.Root.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Root`, data.Root.ValueString())
	}
	if !data.Extract.IsNull() {
		var values []DmAssemblyActionExtract
		data.Extract.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`Extract`+".-1", val.ToBody(ctx, ""))
		}
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

func (data *AssemblyActionExtract) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Root`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Root = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Root = types.StringValue("message.body")
	}
	if value := res.Get(pathRoot + `Extract`); value.Exists() {
		l := []DmAssemblyActionExtract{}
		if value := res.Get(`Extract`); value.Exists() {
			for _, v := range value.Array() {
				item := DmAssemblyActionExtract{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Extract, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAssemblyActionExtractObjectType}, l)
		} else {
			data.Extract = types.ListNull(types.ObjectType{AttrTypes: DmAssemblyActionExtractObjectType})
		}
	} else {
		data.Extract = types.ListNull(types.ObjectType{AttrTypes: DmAssemblyActionExtractObjectType})
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

func (data *AssemblyActionExtract) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Root`); value.Exists() && !data.Root.IsNull() {
		data.Root = tfutils.ParseStringFromGJSON(value)
	} else if data.Root.ValueString() != "message.body" {
		data.Root = types.StringNull()
	}
	if value := res.Get(pathRoot + `Extract`); value.Exists() && !data.Extract.IsNull() {
		l := []DmAssemblyActionExtract{}
		for _, v := range value.Array() {
			item := DmAssemblyActionExtract{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.Extract, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAssemblyActionExtractObjectType}, l)
		} else {
			data.Extract = types.ListNull(types.ObjectType{AttrTypes: DmAssemblyActionExtractObjectType})
		}
	} else {
		data.Extract = types.ListNull(types.ObjectType{AttrTypes: DmAssemblyActionExtractObjectType})
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
