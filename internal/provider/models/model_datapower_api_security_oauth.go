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
	OAuthProvider           types.String                `tfsdk:"o_auth_provider"`
	OAuthFlow               types.String                `tfsdk:"o_auth_flow"`
	OAuthScope              types.String                `tfsdk:"o_auth_scope"`
	OAuthAdvScopeUrl        types.String                `tfsdk:"o_auth_adv_scope_url"`
	OAuthAdvScopeTlsProfile types.String                `tfsdk:"o_auth_adv_scope_tls_profile"`
	DependencyActions       []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var APISecurityOAuthObjectType = map[string]attr.Type{
	"id":                           types.StringType,
	"app_domain":                   types.StringType,
	"user_summary":                 types.StringType,
	"o_auth_provider":              types.StringType,
	"o_auth_flow":                  types.StringType,
	"o_auth_scope":                 types.StringType,
	"o_auth_adv_scope_url":         types.StringType,
	"o_auth_adv_scope_tls_profile": types.StringType,
	"dependency_actions":           actions.ActionsListType,
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
	if !data.OAuthProvider.IsNull() {
		return false
	}
	if !data.OAuthFlow.IsNull() {
		return false
	}
	if !data.OAuthScope.IsNull() {
		return false
	}
	if !data.OAuthAdvScopeUrl.IsNull() {
		return false
	}
	if !data.OAuthAdvScopeTlsProfile.IsNull() {
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
	if !data.OAuthProvider.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OAuthProvider`, data.OAuthProvider.ValueString())
	}
	if !data.OAuthFlow.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OAuthFlow`, data.OAuthFlow.ValueString())
	}
	if !data.OAuthScope.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OAuthScope`, data.OAuthScope.ValueString())
	}
	if !data.OAuthAdvScopeUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OAuthAdvScopeURL`, data.OAuthAdvScopeUrl.ValueString())
	}
	if !data.OAuthAdvScopeTlsProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OAuthAdvScopeTLSProfile`, data.OAuthAdvScopeTlsProfile.ValueString())
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
		data.OAuthProvider = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OAuthProvider = types.StringNull()
	}
	if value := res.Get(pathRoot + `OAuthFlow`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OAuthFlow = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OAuthFlow = types.StringValue("implicit")
	}
	if value := res.Get(pathRoot + `OAuthScope`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OAuthScope = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OAuthScope = types.StringNull()
	}
	if value := res.Get(pathRoot + `OAuthAdvScopeURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OAuthAdvScopeUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OAuthAdvScopeUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `OAuthAdvScopeTLSProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OAuthAdvScopeTlsProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OAuthAdvScopeTlsProfile = types.StringNull()
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
	if value := res.Get(pathRoot + `OAuthProvider`); value.Exists() && !data.OAuthProvider.IsNull() {
		data.OAuthProvider = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OAuthProvider = types.StringNull()
	}
	if value := res.Get(pathRoot + `OAuthFlow`); value.Exists() && !data.OAuthFlow.IsNull() {
		data.OAuthFlow = tfutils.ParseStringFromGJSON(value)
	} else if data.OAuthFlow.ValueString() != "implicit" {
		data.OAuthFlow = types.StringNull()
	}
	if value := res.Get(pathRoot + `OAuthScope`); value.Exists() && !data.OAuthScope.IsNull() {
		data.OAuthScope = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OAuthScope = types.StringNull()
	}
	if value := res.Get(pathRoot + `OAuthAdvScopeURL`); value.Exists() && !data.OAuthAdvScopeUrl.IsNull() {
		data.OAuthAdvScopeUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OAuthAdvScopeUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `OAuthAdvScopeTLSProfile`); value.Exists() && !data.OAuthAdvScopeTlsProfile.IsNull() {
		data.OAuthAdvScopeTlsProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OAuthAdvScopeTlsProfile = types.StringNull()
	}
}
