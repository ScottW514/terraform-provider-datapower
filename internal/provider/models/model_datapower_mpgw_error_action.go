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

type MPGWErrorAction struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	Type              types.String                `tfsdk:"type"`
	RemoteUrl         types.String                `tfsdk:"remote_url"`
	LocalUrl          types.String                `tfsdk:"local_url"`
	ErrorRule         types.String                `tfsdk:"error_rule"`
	StatusCode        types.Int64                 `tfsdk:"status_code"`
	ReasonPhrase      types.String                `tfsdk:"reason_phrase"`
	HeaderInjection   types.List                  `tfsdk:"header_injection"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var MPGWErrorActionRemoteURLCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "static",
	Value:       []string{"redirect", "proxy"},
}
var MPGWErrorActionLocalURLCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "static",
	Value:       []string{"static"},
}
var MPGWErrorActionErrorRuleCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "static",
	Value:       []string{"error-rule"},
}

var MPGWErrorActionObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"user_summary":       types.StringType,
	"type":               types.StringType,
	"remote_url":         types.StringType,
	"local_url":          types.StringType,
	"error_rule":         types.StringType,
	"status_code":        types.Int64Type,
	"reason_phrase":      types.StringType,
	"header_injection":   types.ListType{ElemType: types.ObjectType{AttrTypes: DmWebGWErrorRespHeaderInjectionObjectType}},
	"dependency_actions": actions.ActionsListType,
}

func (data MPGWErrorAction) GetPath() string {
	rest_path := "/mgmt/config/{domain}/MPGWErrorAction"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data MPGWErrorAction) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
		return false
	}
	if !data.RemoteUrl.IsNull() {
		return false
	}
	if !data.LocalUrl.IsNull() {
		return false
	}
	if !data.ErrorRule.IsNull() {
		return false
	}
	if !data.StatusCode.IsNull() {
		return false
	}
	if !data.ReasonPhrase.IsNull() {
		return false
	}
	if !data.HeaderInjection.IsNull() {
		return false
	}
	return true
}

func (data MPGWErrorAction) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Type`, data.Type.ValueString())
	}
	if !data.RemoteUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemoteURL`, data.RemoteUrl.ValueString())
	}
	if !data.LocalUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalURL`, data.LocalUrl.ValueString())
	}
	if !data.ErrorRule.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ErrorRule`, data.ErrorRule.ValueString())
	}
	if !data.StatusCode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StatusCode`, data.StatusCode.ValueInt64())
	}
	if !data.ReasonPhrase.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ReasonPhrase`, data.ReasonPhrase.ValueString())
	}
	if !data.HeaderInjection.IsNull() {
		var dataValues []DmWebGWErrorRespHeaderInjection
		data.HeaderInjection.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`HeaderInjection`+".-1", val.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *MPGWErrorAction) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Type`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringValue("static")
	}
	if value := res.Get(pathRoot + `RemoteURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RemoteUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `ErrorRule`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ErrorRule = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErrorRule = types.StringNull()
	}
	if value := res.Get(pathRoot + `StatusCode`); value.Exists() {
		data.StatusCode = types.Int64Value(value.Int())
	} else {
		data.StatusCode = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ReasonPhrase`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ReasonPhrase = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ReasonPhrase = types.StringNull()
	}
	if value := res.Get(pathRoot + `HeaderInjection`); value.Exists() {
		l := []DmWebGWErrorRespHeaderInjection{}
		if value := res.Get(`HeaderInjection`); value.Exists() {
			for _, v := range value.Array() {
				item := DmWebGWErrorRespHeaderInjection{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.HeaderInjection, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWebGWErrorRespHeaderInjectionObjectType}, l)
		} else {
			data.HeaderInjection = types.ListNull(types.ObjectType{AttrTypes: DmWebGWErrorRespHeaderInjectionObjectType})
		}
	} else {
		data.HeaderInjection = types.ListNull(types.ObjectType{AttrTypes: DmWebGWErrorRespHeaderInjectionObjectType})
	}
}

func (data *MPGWErrorAction) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Type`); value.Exists() && !data.Type.IsNull() {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else if data.Type.ValueString() != "static" {
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemoteURL`); value.Exists() && !data.RemoteUrl.IsNull() {
		data.RemoteUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalURL`); value.Exists() && !data.LocalUrl.IsNull() {
		data.LocalUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `ErrorRule`); value.Exists() && !data.ErrorRule.IsNull() {
		data.ErrorRule = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErrorRule = types.StringNull()
	}
	if value := res.Get(pathRoot + `StatusCode`); value.Exists() && !data.StatusCode.IsNull() {
		data.StatusCode = types.Int64Value(value.Int())
	} else {
		data.StatusCode = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ReasonPhrase`); value.Exists() && !data.ReasonPhrase.IsNull() {
		data.ReasonPhrase = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ReasonPhrase = types.StringNull()
	}
	if value := res.Get(pathRoot + `HeaderInjection`); value.Exists() && !data.HeaderInjection.IsNull() {
		l := []DmWebGWErrorRespHeaderInjection{}
		for _, v := range value.Array() {
			item := DmWebGWErrorRespHeaderInjection{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.HeaderInjection, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWebGWErrorRespHeaderInjectionObjectType}, l)
		} else {
			data.HeaderInjection = types.ListNull(types.ObjectType{AttrTypes: DmWebGWErrorRespHeaderInjectionObjectType})
		}
	} else {
		data.HeaderInjection = types.ListNull(types.ObjectType{AttrTypes: DmWebGWErrorRespHeaderInjectionObjectType})
	}
}
