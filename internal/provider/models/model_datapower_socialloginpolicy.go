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

type SocialLoginPolicy struct {
	Id                        types.String                `tfsdk:"id"`
	AppDomain                 types.String                `tfsdk:"app_domain"`
	UserSummary               types.String                `tfsdk:"user_summary"`
	ClientId                  types.String                `tfsdk:"client_id"`
	ClientSecret              types.String                `tfsdk:"client_secret"`
	ClientGrant               types.String                `tfsdk:"client_grant"`
	ClientScope               types.String                `tfsdk:"client_scope"`
	ClientRedirectUri         types.String                `tfsdk:"client_redirect_uri"`
	ClientOptionalQueryParams types.String                `tfsdk:"client_optional_query_params"`
	ClientSslProfile          types.String                `tfsdk:"client_ssl_profile"`
	SocialProvider            types.String                `tfsdk:"social_provider"`
	ProviderAzEndpoint        types.String                `tfsdk:"provider_az_endpoint"`
	ProviderTokenEndpoint     types.String                `tfsdk:"provider_token_endpoint"`
	ValidateJwtToken          types.Bool                  `tfsdk:"validate_jwt_token"`
	JwtValidator              types.String                `tfsdk:"jwt_validator"`
	DependencyActions         []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var SocialLoginPolicyObjectType = map[string]attr.Type{
	"id":                           types.StringType,
	"app_domain":                   types.StringType,
	"user_summary":                 types.StringType,
	"client_id":                    types.StringType,
	"client_secret":                types.StringType,
	"client_grant":                 types.StringType,
	"client_scope":                 types.StringType,
	"client_redirect_uri":          types.StringType,
	"client_optional_query_params": types.StringType,
	"client_ssl_profile":           types.StringType,
	"social_provider":              types.StringType,
	"provider_az_endpoint":         types.StringType,
	"provider_token_endpoint":      types.StringType,
	"validate_jwt_token":           types.BoolType,
	"jwt_validator":                types.StringType,
	"dependency_actions":           actions.ActionsListType,
}

func (data SocialLoginPolicy) GetPath() string {
	rest_path := "/mgmt/config/{domain}/SocialLoginPolicy"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data SocialLoginPolicy) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.ClientId.IsNull() {
		return false
	}
	if !data.ClientSecret.IsNull() {
		return false
	}
	if !data.ClientGrant.IsNull() {
		return false
	}
	if !data.ClientScope.IsNull() {
		return false
	}
	if !data.ClientRedirectUri.IsNull() {
		return false
	}
	if !data.ClientOptionalQueryParams.IsNull() {
		return false
	}
	if !data.ClientSslProfile.IsNull() {
		return false
	}
	if !data.SocialProvider.IsNull() {
		return false
	}
	if !data.ProviderAzEndpoint.IsNull() {
		return false
	}
	if !data.ProviderTokenEndpoint.IsNull() {
		return false
	}
	if !data.ValidateJwtToken.IsNull() {
		return false
	}
	if !data.JwtValidator.IsNull() {
		return false
	}
	return true
}

func (data SocialLoginPolicy) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.ClientId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ClientID`, data.ClientId.ValueString())
	}
	if !data.ClientSecret.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ClientSecret`, data.ClientSecret.ValueString())
	}
	if !data.ClientGrant.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ClientGrant`, data.ClientGrant.ValueString())
	}
	if !data.ClientScope.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ClientScope`, data.ClientScope.ValueString())
	}
	if !data.ClientRedirectUri.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ClientRedirectURI`, data.ClientRedirectUri.ValueString())
	}
	if !data.ClientOptionalQueryParams.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ClientOptionalQueryParams`, data.ClientOptionalQueryParams.ValueString())
	}
	if !data.ClientSslProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ClientSSLProfile`, data.ClientSslProfile.ValueString())
	}
	if !data.SocialProvider.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Provider`, data.SocialProvider.ValueString())
	}
	if !data.ProviderAzEndpoint.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ProviderAZEndpoint`, data.ProviderAzEndpoint.ValueString())
	}
	if !data.ProviderTokenEndpoint.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ProviderTokenEndpoint`, data.ProviderTokenEndpoint.ValueString())
	}
	if !data.ValidateJwtToken.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ValidateJWTToken`, tfutils.StringFromBool(data.ValidateJwtToken, ""))
	}
	if !data.JwtValidator.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`JWTValidator`, data.JwtValidator.ValueString())
	}
	return body
}

func (data *SocialLoginPolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `ClientID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ClientId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientId = types.StringNull()
	}
	if value := res.Get(pathRoot + `ClientSecret`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ClientSecret = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientSecret = types.StringNull()
	}
	if value := res.Get(pathRoot + `ClientGrant`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ClientGrant = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientGrant = types.StringNull()
	}
	if value := res.Get(pathRoot + `ClientScope`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ClientScope = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientScope = types.StringNull()
	}
	if value := res.Get(pathRoot + `ClientRedirectURI`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ClientRedirectUri = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientRedirectUri = types.StringValue("URL-in/social-login-callback")
	}
	if value := res.Get(pathRoot + `ClientOptionalQueryParams`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ClientOptionalQueryParams = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientOptionalQueryParams = types.StringNull()
	}
	if value := res.Get(pathRoot + `ClientSSLProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ClientSslProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientSslProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `Provider`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SocialProvider = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SocialProvider = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProviderAZEndpoint`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ProviderAzEndpoint = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProviderAzEndpoint = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProviderTokenEndpoint`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ProviderTokenEndpoint = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProviderTokenEndpoint = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidateJWTToken`); value.Exists() {
		data.ValidateJwtToken = tfutils.BoolFromString(value.String())
	} else {
		data.ValidateJwtToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `JWTValidator`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.JwtValidator = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JwtValidator = types.StringNull()
	}
}

func (data *SocialLoginPolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `ClientID`); value.Exists() && !data.ClientId.IsNull() {
		data.ClientId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientId = types.StringNull()
	}
	if value := res.Get(pathRoot + `ClientSecret`); value.Exists() && !data.ClientSecret.IsNull() {
		data.ClientSecret = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientSecret = types.StringNull()
	}
	if value := res.Get(pathRoot + `ClientGrant`); value.Exists() && !data.ClientGrant.IsNull() {
		data.ClientGrant = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientGrant = types.StringNull()
	}
	if value := res.Get(pathRoot + `ClientScope`); value.Exists() && !data.ClientScope.IsNull() {
		data.ClientScope = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientScope = types.StringNull()
	}
	if value := res.Get(pathRoot + `ClientRedirectURI`); value.Exists() && !data.ClientRedirectUri.IsNull() {
		data.ClientRedirectUri = tfutils.ParseStringFromGJSON(value)
	} else if data.ClientRedirectUri.ValueString() != "URL-in/social-login-callback" {
		data.ClientRedirectUri = types.StringNull()
	}
	if value := res.Get(pathRoot + `ClientOptionalQueryParams`); value.Exists() && !data.ClientOptionalQueryParams.IsNull() {
		data.ClientOptionalQueryParams = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientOptionalQueryParams = types.StringNull()
	}
	if value := res.Get(pathRoot + `ClientSSLProfile`); value.Exists() && !data.ClientSslProfile.IsNull() {
		data.ClientSslProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientSslProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `Provider`); value.Exists() && !data.SocialProvider.IsNull() {
		data.SocialProvider = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SocialProvider = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProviderAZEndpoint`); value.Exists() && !data.ProviderAzEndpoint.IsNull() {
		data.ProviderAzEndpoint = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProviderAzEndpoint = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProviderTokenEndpoint`); value.Exists() && !data.ProviderTokenEndpoint.IsNull() {
		data.ProviderTokenEndpoint = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProviderTokenEndpoint = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidateJWTToken`); value.Exists() && !data.ValidateJwtToken.IsNull() {
		data.ValidateJwtToken = tfutils.BoolFromString(value.String())
	} else if !data.ValidateJwtToken.ValueBool() {
		data.ValidateJwtToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `JWTValidator`); value.Exists() && !data.JwtValidator.IsNull() {
		data.JwtValidator = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JwtValidator = types.StringNull()
	}
}
