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

type CookieAttributePolicy struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	CookieAttribute   *DmCookieAttribute          `tfsdk:"cookie_attribute"`
	Domain            types.String                `tfsdk:"domain"`
	Path              types.String                `tfsdk:"path"`
	Interval          types.Int64                 `tfsdk:"interval"`
	CustomAttribute   types.String                `tfsdk:"custom_attribute"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var CookieAttributePolicyDomainCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "cookie_attribute",
	AttrType:    "DmCookieAttribute",
	AttrDefault: "",
	Value:       []string{"domain"},
}
var CookieAttributePolicyPathCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "cookie_attribute",
	AttrType:    "DmCookieAttribute",
	AttrDefault: "",
	Value:       []string{"path"},
}
var CookieAttributePolicyIntervalCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "cookie_attribute",
	AttrType:    "DmCookieAttribute",
	AttrDefault: "",
	Value:       []string{"max-age", "expires"},
}
var CookieAttributePolicyCustomAttributeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "cookie_attribute",
	AttrType:    "DmCookieAttribute",
	AttrDefault: "",
	Value:       []string{"custom"},
}

var CookieAttributePolicyObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"user_summary":       types.StringType,
	"cookie_attribute":   types.ObjectType{AttrTypes: DmCookieAttributeObjectType},
	"domain":             types.StringType,
	"path":               types.StringType,
	"interval":           types.Int64Type,
	"custom_attribute":   types.StringType,
	"dependency_actions": actions.ActionsListType,
}

func (data CookieAttributePolicy) GetPath() string {
	rest_path := "/mgmt/config/{domain}/CookieAttributePolicy"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data CookieAttributePolicy) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if data.CookieAttribute != nil {
		if !data.CookieAttribute.IsNull() {
			return false
		}
	}
	if !data.Domain.IsNull() {
		return false
	}
	if !data.Path.IsNull() {
		return false
	}
	if !data.Interval.IsNull() {
		return false
	}
	if !data.CustomAttribute.IsNull() {
		return false
	}
	return true
}

func (data CookieAttributePolicy) ToBody(ctx context.Context, pathRoot string) string {
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
	if data.CookieAttribute != nil {
		if !data.CookieAttribute.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`CookieAttribute`, data.CookieAttribute.ToBody(ctx, ""))
		}
	}
	if !data.Domain.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Domain`, data.Domain.ValueString())
	}
	if !data.Path.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Path`, data.Path.ValueString())
	}
	if !data.Interval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Interval`, data.Interval.ValueInt64())
	}
	if !data.CustomAttribute.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CustomAttribute`, data.CustomAttribute.ValueString())
	}
	return body
}

func (data *CookieAttributePolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `CookieAttribute`); value.Exists() {
		data.CookieAttribute = &DmCookieAttribute{}
		data.CookieAttribute.FromBody(ctx, "", value)
	} else {
		data.CookieAttribute = nil
	}
	if value := res.Get(pathRoot + `Domain`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Domain = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Domain = types.StringNull()
	}
	if value := res.Get(pathRoot + `Path`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Path = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Path = types.StringValue("/")
	}
	if value := res.Get(pathRoot + `Interval`); value.Exists() {
		data.Interval = types.Int64Value(value.Int())
	} else {
		data.Interval = types.Int64Value(3600)
	}
	if value := res.Get(pathRoot + `CustomAttribute`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CustomAttribute = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CustomAttribute = types.StringNull()
	}
}

func (data *CookieAttributePolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `CookieAttribute`); value.Exists() {
		data.CookieAttribute.UpdateFromBody(ctx, "", value)
	} else {
		data.CookieAttribute = nil
	}
	if value := res.Get(pathRoot + `Domain`); value.Exists() && !data.Domain.IsNull() {
		data.Domain = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Domain = types.StringNull()
	}
	if value := res.Get(pathRoot + `Path`); value.Exists() && !data.Path.IsNull() {
		data.Path = tfutils.ParseStringFromGJSON(value)
	} else if data.Path.ValueString() != "/" {
		data.Path = types.StringNull()
	}
	if value := res.Get(pathRoot + `Interval`); value.Exists() && !data.Interval.IsNull() {
		data.Interval = types.Int64Value(value.Int())
	} else if data.Interval.ValueInt64() != 3600 {
		data.Interval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `CustomAttribute`); value.Exists() && !data.CustomAttribute.IsNull() {
		data.CustomAttribute = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CustomAttribute = types.StringNull()
	}
}
