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

type WebAppErrorHandlingPolicy struct {
	Id                   types.String `tfsdk:"id"`
	AppDomain            types.String `tfsdk:"app_domain"`
	UserSummary          types.String `tfsdk:"user_summary"`
	Type                 types.String `tfsdk:"type"`
	Url                  types.String `tfsdk:"url"`
	ErrorStylePolicyRule types.String `tfsdk:"error_style_policy_rule"`
	ErrorMonitor         types.String `tfsdk:"error_monitor"`
}

var WebAppErrorHandlingPolicyObjectType = map[string]attr.Type{
	"id":                      types.StringType,
	"app_domain":              types.StringType,
	"user_summary":            types.StringType,
	"type":                    types.StringType,
	"url":                     types.StringType,
	"error_style_policy_rule": types.StringType,
	"error_monitor":           types.StringType,
}

func (data WebAppErrorHandlingPolicy) GetPath() string {
	rest_path := "/mgmt/config/{domain}/WebAppErrorHandlingPolicy"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data WebAppErrorHandlingPolicy) IsNull() bool {
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
	if !data.Url.IsNull() {
		return false
	}
	if !data.ErrorStylePolicyRule.IsNull() {
		return false
	}
	if !data.ErrorMonitor.IsNull() {
		return false
	}
	return true
}

func (data WebAppErrorHandlingPolicy) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Url.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`URL`, data.Url.ValueString())
	}
	if !data.ErrorStylePolicyRule.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ErrorStylePolicyRule`, data.ErrorStylePolicyRule.ValueString())
	}
	if !data.ErrorMonitor.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ErrorMonitor`, data.ErrorMonitor.ValueString())
	}
	return body
}

func (data *WebAppErrorHandlingPolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
		data.Type = types.StringValue("standard")
	}
	if value := res.Get(pathRoot + `URL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Url = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Url = types.StringNull()
	}
	if value := res.Get(pathRoot + `ErrorStylePolicyRule`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ErrorStylePolicyRule = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErrorStylePolicyRule = types.StringNull()
	}
	if value := res.Get(pathRoot + `ErrorMonitor`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ErrorMonitor = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErrorMonitor = types.StringNull()
	}
}

func (data *WebAppErrorHandlingPolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	} else if data.Type.ValueString() != "standard" {
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `URL`); value.Exists() && !data.Url.IsNull() {
		data.Url = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Url = types.StringNull()
	}
	if value := res.Get(pathRoot + `ErrorStylePolicyRule`); value.Exists() && !data.ErrorStylePolicyRule.IsNull() {
		data.ErrorStylePolicyRule = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErrorStylePolicyRule = types.StringNull()
	}
	if value := res.Get(pathRoot + `ErrorMonitor`); value.Exists() && !data.ErrorMonitor.IsNull() {
		data.ErrorMonitor = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErrorMonitor = types.StringNull()
	}
}
