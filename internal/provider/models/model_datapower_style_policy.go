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

type StylePolicy struct {
	Id                   types.String                `tfsdk:"id"`
	AppDomain            types.String                `tfsdk:"app_domain"`
	UserSummary          types.String                `tfsdk:"user_summary"`
	DefStylesheetForSoap types.String                `tfsdk:"def_stylesheet_for_soap"`
	DefStylesheetForXsl  types.String                `tfsdk:"def_stylesheet_for_xsl"`
	DefXQueryForJson     types.String                `tfsdk:"def_x_query_for_json"`
	PolicyMaps           types.List                  `tfsdk:"policy_maps"`
	DependencyActions    []*actions.DependencyAction `tfsdk:"dependency_actions"`
	ProviderTarget       types.String                `tfsdk:"provider_target"`
}

var StylePolicyObjectType = map[string]attr.Type{
	"provider_target":         types.StringType,
	"id":                      types.StringType,
	"app_domain":              types.StringType,
	"user_summary":            types.StringType,
	"def_stylesheet_for_soap": types.StringType,
	"def_stylesheet_for_xsl":  types.StringType,
	"def_x_query_for_json":    types.StringType,
	"policy_maps":             types.ListType{ElemType: types.ObjectType{AttrTypes: DmPolicyMapObjectType}},
	"dependency_actions":      actions.ActionsListType,
}

func (data StylePolicy) GetPath() string {
	rest_path := "/mgmt/config/{domain}/StylePolicy"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data StylePolicy) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.DefStylesheetForSoap.IsNull() {
		return false
	}
	if !data.DefStylesheetForXsl.IsNull() {
		return false
	}
	if !data.DefXQueryForJson.IsNull() {
		return false
	}
	if !data.PolicyMaps.IsNull() {
		return false
	}
	return true
}

func (data StylePolicy) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.DefStylesheetForSoap.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DefStylesheetForSoap`, data.DefStylesheetForSoap.ValueString())
	}
	if !data.DefStylesheetForXsl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DefStylesheetForXsl`, data.DefStylesheetForXsl.ValueString())
	}
	if !data.DefXQueryForJson.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DefXQueryForJSON`, data.DefXQueryForJson.ValueString())
	}
	if !data.PolicyMaps.IsNull() {
		var dataValues []DmPolicyMap
		data.PolicyMaps.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.SetRaw(body, pathRoot+`PolicyMaps`+".-1", val.ToBody(ctx, ""))
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`PolicyMaps`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`PolicyMaps`, "[]")
	}
	return body
}

func (data *StylePolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `DefStylesheetForSoap`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DefStylesheetForSoap = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DefStylesheetForSoap = types.StringValue("store:///filter-reject-all.xsl")
	}
	if value := res.Get(pathRoot + `DefStylesheetForXsl`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DefStylesheetForXsl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DefStylesheetForXsl = types.StringValue("store:///identity.xsl")
	}
	if value := res.Get(pathRoot + `DefXQueryForJSON`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DefXQueryForJson = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DefXQueryForJson = types.StringValue("store:///reject-all-json.xq")
	}
	if value := res.Get(pathRoot + `PolicyMaps`); value.Exists() {
		l := []DmPolicyMap{}
		if value := res.Get(`PolicyMaps`); value.Exists() {
			for _, v := range value.Array() {
				item := DmPolicyMap{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.PolicyMaps, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmPolicyMapObjectType}, l)
		} else {
			data.PolicyMaps = types.ListNull(types.ObjectType{AttrTypes: DmPolicyMapObjectType})
		}
	} else {
		data.PolicyMaps = types.ListNull(types.ObjectType{AttrTypes: DmPolicyMapObjectType})
	}
}

func (data *StylePolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `DefStylesheetForSoap`); value.Exists() && !data.DefStylesheetForSoap.IsNull() {
		data.DefStylesheetForSoap = tfutils.ParseStringFromGJSON(value)
	} else if data.DefStylesheetForSoap.ValueString() != "store:///filter-reject-all.xsl" {
		data.DefStylesheetForSoap = types.StringNull()
	}
	if value := res.Get(pathRoot + `DefStylesheetForXsl`); value.Exists() && !data.DefStylesheetForXsl.IsNull() {
		data.DefStylesheetForXsl = tfutils.ParseStringFromGJSON(value)
	} else if data.DefStylesheetForXsl.ValueString() != "store:///identity.xsl" {
		data.DefStylesheetForXsl = types.StringNull()
	}
	if value := res.Get(pathRoot + `DefXQueryForJSON`); value.Exists() && !data.DefXQueryForJson.IsNull() {
		data.DefXQueryForJson = tfutils.ParseStringFromGJSON(value)
	} else if data.DefXQueryForJson.ValueString() != "store:///reject-all-json.xq" {
		data.DefXQueryForJson = types.StringNull()
	}
	if value := res.Get(pathRoot + `PolicyMaps`); value.Exists() && !data.PolicyMaps.IsNull() {
		l := []DmPolicyMap{}
		e := []DmPolicyMap{}
		data.PolicyMaps.ElementsAs(ctx, &e, false)
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
				item := DmPolicyMap{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.PolicyMaps, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmPolicyMapObjectType}, l)
		} else {
			data.PolicyMaps = types.ListNull(types.ObjectType{AttrTypes: DmPolicyMapObjectType})
		}
	} else {
		data.PolicyMaps = types.ListNull(types.ObjectType{AttrTypes: DmPolicyMapObjectType})
	}
}
