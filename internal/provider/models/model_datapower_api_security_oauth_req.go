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

type APISecurityOAuthReq struct {
	Id                  types.String                `tfsdk:"id"`
	AppDomain           types.String                `tfsdk:"app_domain"`
	UserSummary         types.String                `tfsdk:"user_summary"`
	ApiSecurityOauthDef types.String                `tfsdk:"api_security_oauth_def"`
	OauthAllowedScope   types.String                `tfsdk:"oauth_allowed_scope"`
	DependencyActions   []*actions.DependencyAction `tfsdk:"dependency_actions"`
	ProviderTarget      types.String                `tfsdk:"provider_target"`
}

var APISecurityOAuthReqObjectType = map[string]attr.Type{
	"provider_target":        types.StringType,
	"id":                     types.StringType,
	"app_domain":             types.StringType,
	"user_summary":           types.StringType,
	"api_security_oauth_def": types.StringType,
	"oauth_allowed_scope":    types.StringType,
	"dependency_actions":     actions.ActionsListType,
}

func (data APISecurityOAuthReq) GetPath() string {
	rest_path := "/mgmt/config/{domain}/APISecurityOAuthReq"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data APISecurityOAuthReq) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.ApiSecurityOauthDef.IsNull() {
		return false
	}
	if !data.OauthAllowedScope.IsNull() {
		return false
	}
	return true
}

func (data APISecurityOAuthReq) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.ApiSecurityOauthDef.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APISecurityOAuthDef`, data.ApiSecurityOauthDef.ValueString())
	}
	if !data.OauthAllowedScope.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OAuthAllowedScope`, data.OauthAllowedScope.ValueString())
	}
	return body
}

func (data *APISecurityOAuthReq) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `APISecurityOAuthDef`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ApiSecurityOauthDef = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ApiSecurityOauthDef = types.StringNull()
	}
	if value := res.Get(pathRoot + `OAuthAllowedScope`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OauthAllowedScope = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OauthAllowedScope = types.StringNull()
	}
}

func (data *APISecurityOAuthReq) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `APISecurityOAuthDef`); value.Exists() && !data.ApiSecurityOauthDef.IsNull() {
		data.ApiSecurityOauthDef = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ApiSecurityOauthDef = types.StringNull()
	}
	if value := res.Get(pathRoot + `OAuthAllowedScope`); value.Exists() && !data.OauthAllowedScope.IsNull() {
		data.OauthAllowedScope = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OauthAllowedScope = types.StringNull()
	}
}
