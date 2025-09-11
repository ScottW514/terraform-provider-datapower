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

type MCFHttpHeader struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	HttpName          types.String                `tfsdk:"http_name"`
	HttpValue         types.String                `tfsdk:"http_value"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var MCFHttpHeaderObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"http_name":          types.StringType,
	"http_value":         types.StringType,
	"user_summary":       types.StringType,
	"dependency_actions": actions.ActionsListType,
}

func (data MCFHttpHeader) GetPath() string {
	rest_path := "/mgmt/config/{domain}/MCFHttpHeader"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data MCFHttpHeader) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.HttpName.IsNull() {
		return false
	}
	if !data.HttpValue.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	return true
}

func (data MCFHttpHeader) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.HttpName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HttpName`, data.HttpName.ValueString())
	}
	if !data.HttpValue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HttpValue`, data.HttpValue.ValueString())
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	return body
}

func (data *MCFHttpHeader) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `HttpName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HttpName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HttpName = types.StringNull()
	}
	if value := res.Get(pathRoot + `HttpValue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HttpValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HttpValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
}

func (data *MCFHttpHeader) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `HttpName`); value.Exists() && !data.HttpName.IsNull() {
		data.HttpName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HttpName = types.StringNull()
	}
	if value := res.Get(pathRoot + `HttpValue`); value.Exists() && !data.HttpValue.IsNull() {
		data.HttpValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HttpValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
}
