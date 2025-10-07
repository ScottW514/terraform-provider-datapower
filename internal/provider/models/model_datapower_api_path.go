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

type APIPath struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	Path              types.String                `tfsdk:"path"`
	Operation         types.List                  `tfsdk:"operation"`
	RequestSchema     types.String                `tfsdk:"request_schema"`
	Parameter         types.List                  `tfsdk:"parameter"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var APIPathObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"user_summary":       types.StringType,
	"path":               types.StringType,
	"operation":          types.ListType{ElemType: types.StringType},
	"request_schema":     types.StringType,
	"parameter":          types.ListType{ElemType: types.ObjectType{AttrTypes: DmAPIParameterObjectType}},
	"dependency_actions": actions.ActionsListType,
}

func (data APIPath) GetPath() string {
	rest_path := "/mgmt/config/{domain}/APIPath"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data APIPath) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Path.IsNull() {
		return false
	}
	if !data.Operation.IsNull() {
		return false
	}
	if !data.RequestSchema.IsNull() {
		return false
	}
	if !data.Parameter.IsNull() {
		return false
	}
	return true
}

func (data APIPath) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Path.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Path`, data.Path.ValueString())
	}
	if !data.Operation.IsNull() {
		var dataValues []string
		data.Operation.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`Operation`+".-1", map[string]string{"value": val})
		}
	}
	if !data.RequestSchema.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RequestSchema`, data.RequestSchema.ValueString())
	}
	if !data.Parameter.IsNull() {
		var dataValues []DmAPIParameter
		data.Parameter.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`Parameter`+".-1", val.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *APIPath) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Path`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Path = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Path = types.StringValue("/")
	}
	if value := res.Get(pathRoot + `Operation`); value.Exists() {
		data.Operation = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Operation = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `RequestSchema`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RequestSchema = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RequestSchema = types.StringNull()
	}
	if value := res.Get(pathRoot + `Parameter`); value.Exists() {
		l := []DmAPIParameter{}
		if value := res.Get(`Parameter`); value.Exists() {
			for _, v := range value.Array() {
				item := DmAPIParameter{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Parameter, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPIParameterObjectType}, l)
		} else {
			data.Parameter = types.ListNull(types.ObjectType{AttrTypes: DmAPIParameterObjectType})
		}
	} else {
		data.Parameter = types.ListNull(types.ObjectType{AttrTypes: DmAPIParameterObjectType})
	}
}

func (data *APIPath) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Path`); value.Exists() && !data.Path.IsNull() {
		data.Path = tfutils.ParseStringFromGJSON(value)
	} else if data.Path.ValueString() != "/" {
		data.Path = types.StringNull()
	}
	if value := res.Get(pathRoot + `Operation`); value.Exists() && !data.Operation.IsNull() {
		data.Operation = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Operation = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `RequestSchema`); value.Exists() && !data.RequestSchema.IsNull() {
		data.RequestSchema = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RequestSchema = types.StringNull()
	}
	if value := res.Get(pathRoot + `Parameter`); value.Exists() && !data.Parameter.IsNull() {
		l := []DmAPIParameter{}
		e := []DmAPIParameter{}
		data.Parameter.ElementsAs(ctx, &e, false)
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
				item := DmAPIParameter{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Parameter, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPIParameterObjectType}, l)
		} else {
			data.Parameter = types.ListNull(types.ObjectType{AttrTypes: DmAPIParameterObjectType})
		}
	} else {
		data.Parameter = types.ListNull(types.ObjectType{AttrTypes: DmAPIParameterObjectType})
	}
}
