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

type APISecurityOAuth struct {
	Id                      types.String                `tfsdk:"id"`
	AppDomain               types.String                `tfsdk:"app_domain"`
	UserSummary             types.String                `tfsdk:"user_summary"`
	OauthProvider           types.String                `tfsdk:"oauth_provider"`
	OauthFlow               types.String                `tfsdk:"oauth_flow"`
	OauthScope              types.String                `tfsdk:"oauth_scope"`
	OauthAdvScopeUrl        types.String                `tfsdk:"oauth_adv_scope_url"`
	OauthAdvScopeTlsProfile types.String                `tfsdk:"oauth_adv_scope_tls_profile"`
	DependencyActions       []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var APISecurityOAuthObjectType = map[string]attr.Type{
	"id":                          types.StringType,
	"app_domain":                  types.StringType,
	"user_summary":                types.StringType,
	"oauth_provider":              types.StringType,
	"oauth_flow":                  types.StringType,
	"oauth_scope":                 types.StringType,
	"oauth_adv_scope_url":         types.StringType,
	"oauth_adv_scope_tls_profile": types.StringType,
	"dependency_actions":          actions.ActionsListType,
}

func (data APISecurityOAuth) GetPath() string {
	rest_path := "/mgmt/config/{domain}/APISecurityOAuth"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data APISecurityOAuth) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.OauthProvider.IsNull() {
		return false
	}
	if !data.OauthFlow.IsNull() {
		return false
	}
	if !data.OauthScope.IsNull() {
		return false
	}
	if !data.OauthAdvScopeUrl.IsNull() {
		return false
	}
	if !data.OauthAdvScopeTlsProfile.IsNull() {
		return false
	}
	return true
}

func (data APISecurityOAuth) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.OauthProvider.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OAuthProvider`, data.OauthProvider.ValueString())
	}
	if !data.OauthFlow.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OAuthFlow`, data.OauthFlow.ValueString())
	}
	if !data.OauthScope.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OAuthScope`, data.OauthScope.ValueString())
	}
	if !data.OauthAdvScopeUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OAuthAdvScopeURL`, data.OauthAdvScopeUrl.ValueString())
	}
	if !data.OauthAdvScopeTlsProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OAuthAdvScopeTLSProfile`, data.OauthAdvScopeTlsProfile.ValueString())
	}
	return body
}

func (data *APISecurityOAuth) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `OAuthProvider`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OauthProvider = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OauthProvider = types.StringNull()
	}
	if value := res.Get(pathRoot + `OAuthFlow`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OauthFlow = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OauthFlow = types.StringValue("implicit")
	}
	if value := res.Get(pathRoot + `OAuthScope`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OauthScope = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OauthScope = types.StringNull()
	}
	if value := res.Get(pathRoot + `OAuthAdvScopeURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OauthAdvScopeUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OauthAdvScopeUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `OAuthAdvScopeTLSProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OauthAdvScopeTlsProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OauthAdvScopeTlsProfile = types.StringNull()
	}
}

func (data *APISecurityOAuth) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `OAuthProvider`); value.Exists() && !data.OauthProvider.IsNull() {
		data.OauthProvider = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OauthProvider = types.StringNull()
	}
	if value := res.Get(pathRoot + `OAuthFlow`); value.Exists() && !data.OauthFlow.IsNull() {
		data.OauthFlow = tfutils.ParseStringFromGJSON(value)
	} else if data.OauthFlow.ValueString() != "implicit" {
		data.OauthFlow = types.StringNull()
	}
	if value := res.Get(pathRoot + `OAuthScope`); value.Exists() && !data.OauthScope.IsNull() {
		data.OauthScope = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OauthScope = types.StringNull()
	}
	if value := res.Get(pathRoot + `OAuthAdvScopeURL`); value.Exists() && !data.OauthAdvScopeUrl.IsNull() {
		data.OauthAdvScopeUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OauthAdvScopeUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `OAuthAdvScopeTLSProfile`); value.Exists() && !data.OauthAdvScopeTlsProfile.IsNull() {
		data.OauthAdvScopeTlsProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OauthAdvScopeTlsProfile = types.StringNull()
	}
}
