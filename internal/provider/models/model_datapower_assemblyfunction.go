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

type AssemblyFunction struct {
	Id          types.String `tfsdk:"id"`
	AppDomain   types.String `tfsdk:"app_domain"`
	UserSummary types.String `tfsdk:"user_summary"`
	Title       types.String `tfsdk:"title"`
	Description types.String `tfsdk:"description"`
	Scope       types.String `tfsdk:"scope"`
	Parameter   types.List   `tfsdk:"parameter"`
	Assembly    types.String `tfsdk:"assembly"`
}

var AssemblyFunctionObjectType = map[string]attr.Type{
	"id":           types.StringType,
	"app_domain":   types.StringType,
	"user_summary": types.StringType,
	"title":        types.StringType,
	"description":  types.StringType,
	"scope":        types.StringType,
	"parameter":    types.ListType{ElemType: types.ObjectType{AttrTypes: DmAssemblyFunctionParameterObjectType}},
	"assembly":     types.StringType,
}

func (data AssemblyFunction) GetPath() string {
	rest_path := "/mgmt/config/{domain}/AssemblyFunction"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data AssemblyFunction) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Title.IsNull() {
		return false
	}
	if !data.Description.IsNull() {
		return false
	}
	if !data.Scope.IsNull() {
		return false
	}
	if !data.Parameter.IsNull() {
		return false
	}
	if !data.Assembly.IsNull() {
		return false
	}
	return true
}

func (data AssemblyFunction) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Title.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Title`, data.Title.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Description`, data.Description.ValueString())
	}
	if !data.Scope.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Scope`, data.Scope.ValueString())
	}
	if !data.Parameter.IsNull() {
		var values []DmAssemblyFunctionParameter
		data.Parameter.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`Parameter`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.Assembly.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Assembly`, data.Assembly.ValueString())
	}
	return body
}

func (data *AssemblyFunction) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Title`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Title = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Title = types.StringNull()
	}
	if value := res.Get(pathRoot + `Description`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Description = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get(pathRoot + `Scope`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Scope = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Scope = types.StringNull()
	}
	if value := res.Get(pathRoot + `Parameter`); value.Exists() {
		l := []DmAssemblyFunctionParameter{}
		if value := res.Get(`Parameter`); value.Exists() {
			for _, v := range value.Array() {
				item := DmAssemblyFunctionParameter{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Parameter, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAssemblyFunctionParameterObjectType}, l)
		} else {
			data.Parameter = types.ListNull(types.ObjectType{AttrTypes: DmAssemblyFunctionParameterObjectType})
		}
	} else {
		data.Parameter = types.ListNull(types.ObjectType{AttrTypes: DmAssemblyFunctionParameterObjectType})
	}
	if value := res.Get(pathRoot + `Assembly`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Assembly = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Assembly = types.StringNull()
	}
}

func (data *AssemblyFunction) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Title`); value.Exists() && !data.Title.IsNull() {
		data.Title = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Title = types.StringNull()
	}
	if value := res.Get(pathRoot + `Description`); value.Exists() && !data.Description.IsNull() {
		data.Description = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get(pathRoot + `Scope`); value.Exists() && !data.Scope.IsNull() {
		data.Scope = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Scope = types.StringNull()
	}
	if value := res.Get(pathRoot + `Parameter`); value.Exists() && !data.Parameter.IsNull() {
		l := []DmAssemblyFunctionParameter{}
		for _, v := range value.Array() {
			item := DmAssemblyFunctionParameter{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.Parameter, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAssemblyFunctionParameterObjectType}, l)
		} else {
			data.Parameter = types.ListNull(types.ObjectType{AttrTypes: DmAssemblyFunctionParameterObjectType})
		}
	} else {
		data.Parameter = types.ListNull(types.ObjectType{AttrTypes: DmAssemblyFunctionParameterObjectType})
	}
	if value := res.Get(pathRoot + `Assembly`); value.Exists() && !data.Assembly.IsNull() {
		data.Assembly = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Assembly = types.StringNull()
	}
}
