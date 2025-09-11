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

type APISecurityAPIKey struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	Where             types.String                `tfsdk:"where"`
	Type              types.String                `tfsdk:"type"`
	KeyName           types.String                `tfsdk:"key_name"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var APISecurityAPIKeyObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"user_summary":       types.StringType,
	"where":              types.StringType,
	"type":               types.StringType,
	"key_name":           types.StringType,
	"dependency_actions": actions.ActionsListType,
}

func (data APISecurityAPIKey) GetPath() string {
	rest_path := "/mgmt/config/{domain}/APISecurityAPIKey"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data APISecurityAPIKey) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Where.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
		return false
	}
	if !data.KeyName.IsNull() {
		return false
	}
	return true
}

func (data APISecurityAPIKey) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Where.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Where`, data.Where.ValueString())
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Type`, data.Type.ValueString())
	}
	if !data.KeyName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`KeyName`, data.KeyName.ValueString())
	}
	return body
}

func (data *APISecurityAPIKey) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Where`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Where = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Where = types.StringValue("header")
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringValue("id")
	}
	if value := res.Get(pathRoot + `KeyName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.KeyName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.KeyName = types.StringNull()
	}
}

func (data *APISecurityAPIKey) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Where`); value.Exists() && !data.Where.IsNull() {
		data.Where = tfutils.ParseStringFromGJSON(value)
	} else if data.Where.ValueString() != "header" {
		data.Where = types.StringNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && !data.Type.IsNull() {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else if data.Type.ValueString() != "id" {
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `KeyName`); value.Exists() && !data.KeyName.IsNull() {
		data.KeyName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.KeyName = types.StringNull()
	}
}
