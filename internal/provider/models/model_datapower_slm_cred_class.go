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

type SLMCredClass struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	CredType          types.String                `tfsdk:"cred_type"`
	CredMatchType     types.String                `tfsdk:"cred_match_type"`
	CredValue         types.List                  `tfsdk:"cred_value"`
	Stylesheet        types.String                `tfsdk:"stylesheet"`
	Header            types.String                `tfsdk:"header"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var SLMCredClassCredMatchTypeCondVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "cred_type",
	AttrType:    "String",
	AttrDefault: "aaa-mapped-credential",
	Value:       []string{"custom-stylesheet"},
}
var SLMCredClassStylesheetCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "cred_type",
	AttrType:    "String",
	AttrDefault: "aaa-mapped-credential",
	Value:       []string{"custom-stylesheet"},
}
var SLMCredClassHeaderCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "cred_type",
	AttrType:    "String",
	AttrDefault: "aaa-mapped-credential",
	Value:       []string{"ip-from-header", "request-header"},
}

var SLMCredClassObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"user_summary":       types.StringType,
	"cred_type":          types.StringType,
	"cred_match_type":    types.StringType,
	"cred_value":         types.ListType{ElemType: types.StringType},
	"stylesheet":         types.StringType,
	"header":             types.StringType,
	"dependency_actions": actions.ActionsListType,
}

func (data SLMCredClass) GetPath() string {
	rest_path := "/mgmt/config/{domain}/SLMCredClass"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data SLMCredClass) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.CredType.IsNull() {
		return false
	}
	if !data.CredMatchType.IsNull() {
		return false
	}
	if !data.CredValue.IsNull() {
		return false
	}
	if !data.Stylesheet.IsNull() {
		return false
	}
	if !data.Header.IsNull() {
		return false
	}
	return true
}

func (data SLMCredClass) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.CredType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CredType`, data.CredType.ValueString())
	}
	if !data.CredMatchType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CredMatchType`, data.CredMatchType.ValueString())
	}
	if !data.CredValue.IsNull() {
		var dataValues []string
		data.CredValue.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`CredValue`+".-1", map[string]string{"value": val})
		}
	}
	if !data.Stylesheet.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Stylesheet`, data.Stylesheet.ValueString())
	}
	if !data.Header.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Header`, data.Header.ValueString())
	}
	return body
}

func (data *SLMCredClass) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `CredType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CredType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CredType = types.StringValue("aaa-mapped-credential")
	}
	if value := res.Get(pathRoot + `CredMatchType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CredMatchType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CredMatchType = types.StringValue("per-extracted-value")
	}
	if value := res.Get(pathRoot + `CredValue`); value.Exists() {
		data.CredValue = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.CredValue = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `Stylesheet`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Stylesheet = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Stylesheet = types.StringNull()
	}
	if value := res.Get(pathRoot + `Header`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Header = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Header = types.StringNull()
	}
}

func (data *SLMCredClass) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `CredType`); value.Exists() && !data.CredType.IsNull() {
		data.CredType = tfutils.ParseStringFromGJSON(value)
	} else if data.CredType.ValueString() != "aaa-mapped-credential" {
		data.CredType = types.StringNull()
	}
	if value := res.Get(pathRoot + `CredMatchType`); value.Exists() && !data.CredMatchType.IsNull() {
		data.CredMatchType = tfutils.ParseStringFromGJSON(value)
	} else if data.CredMatchType.ValueString() != "per-extracted-value" {
		data.CredMatchType = types.StringNull()
	}
	if value := res.Get(pathRoot + `CredValue`); value.Exists() && !data.CredValue.IsNull() {
		data.CredValue = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.CredValue = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `Stylesheet`); value.Exists() && !data.Stylesheet.IsNull() {
		data.Stylesheet = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Stylesheet = types.StringNull()
	}
	if value := res.Get(pathRoot + `Header`); value.Exists() && !data.Header.IsNull() {
		data.Header = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Header = types.StringNull()
	}
}
