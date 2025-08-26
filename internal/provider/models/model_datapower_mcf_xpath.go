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

type MCFXPath struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	XpathExpression   types.String                `tfsdk:"xpath_expression"`
	XpathValue        types.String                `tfsdk:"xpath_value"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var MCFXPathObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"xpath_expression":   types.StringType,
	"xpath_value":        types.StringType,
	"user_summary":       types.StringType,
	"dependency_actions": actions.ActionsListType,
}

func (data MCFXPath) GetPath() string {
	rest_path := "/mgmt/config/{domain}/MCFXPath"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data MCFXPath) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.XpathExpression.IsNull() {
		return false
	}
	if !data.XpathValue.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	return true
}

func (data MCFXPath) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.XpathExpression.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XPathExpression`, data.XpathExpression.ValueString())
	}
	if !data.XpathValue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XPathValue`, data.XpathValue.ValueString())
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	return body
}

func (data *MCFXPath) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `XPathExpression`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.XpathExpression = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XpathExpression = types.StringNull()
	}
	if value := res.Get(pathRoot + `XPathValue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.XpathValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XpathValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
}

func (data *MCFXPath) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `XPathExpression`); value.Exists() && !data.XpathExpression.IsNull() {
		data.XpathExpression = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XpathExpression = types.StringNull()
	}
	if value := res.Get(pathRoot + `XPathValue`); value.Exists() && !data.XpathValue.IsNull() {
		data.XpathValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XpathValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
}
