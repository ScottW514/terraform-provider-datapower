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

type OAuthSupportedClient struct {
	Id                         types.String          `tfsdk:"id"`
	AppDomain                  types.String          `tfsdk:"app_domain"`
	UserSummary                types.String          `tfsdk:"user_summary"`
	Customized                 types.Bool            `tfsdk:"customized"`
	CustomizedProcessUrl       types.String          `tfsdk:"customized_process_url"`
	OAuthRole                  *DmOAuthRole          `tfsdk:"o_auth_role"`
	AzGrant                    *DmOAuthAZGrantType   `tfsdk:"az_grant"`
	ClientType                 types.String          `tfsdk:"client_type"`
	CheckClientCredential      types.Bool            `tfsdk:"check_client_credential"`
	UseValidationUrl           types.Bool            `tfsdk:"use_validation_url"`
	ClientAuthenMethod         types.String          `tfsdk:"client_authen_method"`
	ClientValCred              types.String          `tfsdk:"client_val_cred"`
	GenerateClientSecret       types.Bool            `tfsdk:"generate_client_secret"`
	ClientSecret               types.String          `tfsdk:"client_secret"`
	Caching                    types.String          `tfsdk:"caching"`
	ValidationUrl              types.String          `tfsdk:"validation_url"`
	ValidationFeatures         *DmValidationFeatures `tfsdk:"validation_features"`
	RedirectUri                types.List            `tfsdk:"redirect_uri"`
	CustomScopeCheck           types.Bool            `tfsdk:"custom_scope_check"`
	Scope                      types.String          `tfsdk:"scope"`
	ScopeUrl                   types.String          `tfsdk:"scope_url"`
	DefaultScope               types.String          `tfsdk:"default_scope"`
	TokenSecret                types.String          `tfsdk:"token_secret"`
	LocalAzPageUrl             types.String          `tfsdk:"local_az_page_url"`
	DpStateLifeTime            types.Int64           `tfsdk:"dp_state_life_time"`
	AuCodeLifeTime             types.Int64           `tfsdk:"au_code_life_time"`
	AccessTokenLifeTime        types.Int64           `tfsdk:"access_token_life_time"`
	RefreshTokenAllowed        types.Int64           `tfsdk:"refresh_token_allowed"`
	RefreshTokenLifeTime       types.Int64           `tfsdk:"refresh_token_life_time"`
	MaxConsentLifeTime         types.Int64           `tfsdk:"max_consent_life_time"`
	CustomResourceOwner        types.Bool            `tfsdk:"custom_resource_owner"`
	ResourceOwnerUrl           types.String          `tfsdk:"resource_owner_url"`
	AdditionalOAuthProcessUrl  types.String          `tfsdk:"additional_o_auth_process_url"`
	RsSetHeader                *DmOAuthRSSetHeader   `tfsdk:"rs_set_header"`
	ValidationUrlsslClientType types.String          `tfsdk:"validation_urlssl_client_type"`
	ValidationUrlsslClient     types.String          `tfsdk:"validation_urlssl_client"`
	JwtGrantValidator          types.String          `tfsdk:"jwt_grant_validator"`
	ClientJwtValidator         types.String          `tfsdk:"client_jwt_validator"`
	OidcidTokenGenerator       types.String          `tfsdk:"oidcid_token_generator"`
	OAuthFeatures              *DmOAuthFeatures      `tfsdk:"o_auth_features"`
}

var OAuthSupportedClientObjectType = map[string]attr.Type{
	"id":                            types.StringType,
	"app_domain":                    types.StringType,
	"user_summary":                  types.StringType,
	"customized":                    types.BoolType,
	"customized_process_url":        types.StringType,
	"o_auth_role":                   types.ObjectType{AttrTypes: DmOAuthRoleObjectType},
	"az_grant":                      types.ObjectType{AttrTypes: DmOAuthAZGrantTypeObjectType},
	"client_type":                   types.StringType,
	"check_client_credential":       types.BoolType,
	"use_validation_url":            types.BoolType,
	"client_authen_method":          types.StringType,
	"client_val_cred":               types.StringType,
	"generate_client_secret":        types.BoolType,
	"client_secret":                 types.StringType,
	"caching":                       types.StringType,
	"validation_url":                types.StringType,
	"validation_features":           types.ObjectType{AttrTypes: DmValidationFeaturesObjectType},
	"redirect_uri":                  types.ListType{ElemType: types.StringType},
	"custom_scope_check":            types.BoolType,
	"scope":                         types.StringType,
	"scope_url":                     types.StringType,
	"default_scope":                 types.StringType,
	"token_secret":                  types.StringType,
	"local_az_page_url":             types.StringType,
	"dp_state_life_time":            types.Int64Type,
	"au_code_life_time":             types.Int64Type,
	"access_token_life_time":        types.Int64Type,
	"refresh_token_allowed":         types.Int64Type,
	"refresh_token_life_time":       types.Int64Type,
	"max_consent_life_time":         types.Int64Type,
	"custom_resource_owner":         types.BoolType,
	"resource_owner_url":            types.StringType,
	"additional_o_auth_process_url": types.StringType,
	"rs_set_header":                 types.ObjectType{AttrTypes: DmOAuthRSSetHeaderObjectType},
	"validation_urlssl_client_type": types.StringType,
	"validation_urlssl_client":      types.StringType,
	"jwt_grant_validator":           types.StringType,
	"client_jwt_validator":          types.StringType,
	"oidcid_token_generator":        types.StringType,
	"o_auth_features":               types.ObjectType{AttrTypes: DmOAuthFeaturesObjectType},
}

func (data OAuthSupportedClient) GetPath() string {
	rest_path := "/mgmt/config/{domain}/OAuthSupportedClient"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data OAuthSupportedClient) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Customized.IsNull() {
		return false
	}
	if !data.CustomizedProcessUrl.IsNull() {
		return false
	}
	if data.OAuthRole != nil {
		if !data.OAuthRole.IsNull() {
			return false
		}
	}
	if data.AzGrant != nil {
		if !data.AzGrant.IsNull() {
			return false
		}
	}
	if !data.ClientType.IsNull() {
		return false
	}
	if !data.CheckClientCredential.IsNull() {
		return false
	}
	if !data.UseValidationUrl.IsNull() {
		return false
	}
	if !data.ClientAuthenMethod.IsNull() {
		return false
	}
	if !data.ClientValCred.IsNull() {
		return false
	}
	if !data.GenerateClientSecret.IsNull() {
		return false
	}
	if !data.ClientSecret.IsNull() {
		return false
	}
	if !data.Caching.IsNull() {
		return false
	}
	if !data.ValidationUrl.IsNull() {
		return false
	}
	if data.ValidationFeatures != nil {
		if !data.ValidationFeatures.IsNull() {
			return false
		}
	}
	if !data.RedirectUri.IsNull() {
		return false
	}
	if !data.CustomScopeCheck.IsNull() {
		return false
	}
	if !data.Scope.IsNull() {
		return false
	}
	if !data.ScopeUrl.IsNull() {
		return false
	}
	if !data.DefaultScope.IsNull() {
		return false
	}
	if !data.TokenSecret.IsNull() {
		return false
	}
	if !data.LocalAzPageUrl.IsNull() {
		return false
	}
	if !data.DpStateLifeTime.IsNull() {
		return false
	}
	if !data.AuCodeLifeTime.IsNull() {
		return false
	}
	if !data.AccessTokenLifeTime.IsNull() {
		return false
	}
	if !data.RefreshTokenAllowed.IsNull() {
		return false
	}
	if !data.RefreshTokenLifeTime.IsNull() {
		return false
	}
	if !data.MaxConsentLifeTime.IsNull() {
		return false
	}
	if !data.CustomResourceOwner.IsNull() {
		return false
	}
	if !data.ResourceOwnerUrl.IsNull() {
		return false
	}
	if !data.AdditionalOAuthProcessUrl.IsNull() {
		return false
	}
	if data.RsSetHeader != nil {
		if !data.RsSetHeader.IsNull() {
			return false
		}
	}
	if !data.ValidationUrlsslClientType.IsNull() {
		return false
	}
	if !data.ValidationUrlsslClient.IsNull() {
		return false
	}
	if !data.JwtGrantValidator.IsNull() {
		return false
	}
	if !data.ClientJwtValidator.IsNull() {
		return false
	}
	if !data.OidcidTokenGenerator.IsNull() {
		return false
	}
	if data.OAuthFeatures != nil {
		if !data.OAuthFeatures.IsNull() {
			return false
		}
	}
	return true
}

func (data OAuthSupportedClient) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Customized.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Customized`, tfutils.StringFromBool(data.Customized, false))
	}
	if !data.CustomizedProcessUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CustomizedProcessUrl`, data.CustomizedProcessUrl.ValueString())
	}
	if data.OAuthRole != nil {
		if !data.OAuthRole.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`OAuthRole`, data.OAuthRole.ToBody(ctx, ""))
		}
	}
	if data.AzGrant != nil {
		if !data.AzGrant.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`AZGrant`, data.AzGrant.ToBody(ctx, ""))
		}
	}
	if !data.ClientType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ClientType`, data.ClientType.ValueString())
	}
	if !data.CheckClientCredential.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CheckClientCredential`, tfutils.StringFromBool(data.CheckClientCredential, false))
	}
	if !data.UseValidationUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseValidationUrl`, tfutils.StringFromBool(data.UseValidationUrl, false))
	}
	if !data.ClientAuthenMethod.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ClientAuthenMethod`, data.ClientAuthenMethod.ValueString())
	}
	if !data.ClientValCred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ClientValCred`, data.ClientValCred.ValueString())
	}
	if !data.GenerateClientSecret.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GenerateClientSecret`, tfutils.StringFromBool(data.GenerateClientSecret, false))
	}
	if !data.ClientSecret.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ClientSecret`, data.ClientSecret.ValueString())
	}
	if !data.Caching.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Caching`, data.Caching.ValueString())
	}
	if !data.ValidationUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ValidationURL`, data.ValidationUrl.ValueString())
	}
	if data.ValidationFeatures != nil {
		if !data.ValidationFeatures.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`ValidationFeatures`, data.ValidationFeatures.ToBody(ctx, ""))
		}
	}
	if !data.RedirectUri.IsNull() {
		var values []string
		data.RedirectUri.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`RedirectURI`+".-1", map[string]string{"value": val})
		}
	}
	if !data.CustomScopeCheck.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CustomScopeCheck`, tfutils.StringFromBool(data.CustomScopeCheck, false))
	}
	if !data.Scope.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Scope`, data.Scope.ValueString())
	}
	if !data.ScopeUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ScopeUrl`, data.ScopeUrl.ValueString())
	}
	if !data.DefaultScope.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DefaultScope`, data.DefaultScope.ValueString())
	}
	if !data.TokenSecret.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TokenSecret`, data.TokenSecret.ValueString())
	}
	if !data.LocalAzPageUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalAZPageUrl`, data.LocalAzPageUrl.ValueString())
	}
	if !data.DpStateLifeTime.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DPStateLifeTime`, data.DpStateLifeTime.ValueInt64())
	}
	if !data.AuCodeLifeTime.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUCodeLifeTime`, data.AuCodeLifeTime.ValueInt64())
	}
	if !data.AccessTokenLifeTime.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AccessTokenLifeTime`, data.AccessTokenLifeTime.ValueInt64())
	}
	if !data.RefreshTokenAllowed.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RefreshTokenAllowed`, data.RefreshTokenAllowed.ValueInt64())
	}
	if !data.RefreshTokenLifeTime.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RefreshTokenLifeTime`, data.RefreshTokenLifeTime.ValueInt64())
	}
	if !data.MaxConsentLifeTime.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxConsentLifeTime`, data.MaxConsentLifeTime.ValueInt64())
	}
	if !data.CustomResourceOwner.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CustomResourceOwner`, tfutils.StringFromBool(data.CustomResourceOwner, false))
	}
	if !data.ResourceOwnerUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ResourceOwnerUrl`, data.ResourceOwnerUrl.ValueString())
	}
	if !data.AdditionalOAuthProcessUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AdditionalOAuthProcessUrl`, data.AdditionalOAuthProcessUrl.ValueString())
	}
	if data.RsSetHeader != nil {
		if !data.RsSetHeader.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`RSSetHeader`, data.RsSetHeader.ToBody(ctx, ""))
		}
	}
	if !data.ValidationUrlsslClientType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ValidationURLSSLClientType`, data.ValidationUrlsslClientType.ValueString())
	}
	if !data.ValidationUrlsslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ValidationURLSSLClient`, data.ValidationUrlsslClient.ValueString())
	}
	if !data.JwtGrantValidator.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`JWTGrantValidator`, data.JwtGrantValidator.ValueString())
	}
	if !data.ClientJwtValidator.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ClientJWTValidator`, data.ClientJwtValidator.ValueString())
	}
	if !data.OidcidTokenGenerator.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OIDCIDTokenGenerator`, data.OidcidTokenGenerator.ValueString())
	}
	if data.OAuthFeatures != nil {
		if !data.OAuthFeatures.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`OAuthFeatures`, data.OAuthFeatures.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *OAuthSupportedClient) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Customized`); value.Exists() {
		data.Customized = tfutils.BoolFromString(value.String())
	} else {
		data.Customized = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CustomizedProcessUrl`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CustomizedProcessUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CustomizedProcessUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `OAuthRole`); value.Exists() {
		data.OAuthRole = &DmOAuthRole{}
		data.OAuthRole.FromBody(ctx, "", value)
	} else {
		data.OAuthRole = nil
	}
	if value := res.Get(pathRoot + `AZGrant`); value.Exists() {
		data.AzGrant = &DmOAuthAZGrantType{}
		data.AzGrant.FromBody(ctx, "", value)
	} else {
		data.AzGrant = nil
	}
	if value := res.Get(pathRoot + `ClientType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ClientType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientType = types.StringValue("confidential")
	}
	if value := res.Get(pathRoot + `CheckClientCredential`); value.Exists() {
		data.CheckClientCredential = tfutils.BoolFromString(value.String())
	} else {
		data.CheckClientCredential = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UseValidationUrl`); value.Exists() {
		data.UseValidationUrl = tfutils.BoolFromString(value.String())
	} else {
		data.UseValidationUrl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ClientAuthenMethod`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ClientAuthenMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientAuthenMethod = types.StringValue("secret")
	}
	if value := res.Get(pathRoot + `ClientValCred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ClientValCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientValCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `GenerateClientSecret`); value.Exists() {
		data.GenerateClientSecret = tfutils.BoolFromString(value.String())
	} else {
		data.GenerateClientSecret = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ClientSecret`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ClientSecret = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientSecret = types.StringNull()
	}
	if value := res.Get(pathRoot + `Caching`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Caching = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Caching = types.StringValue("replay")
	}
	if value := res.Get(pathRoot + `ValidationURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ValidationUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ValidationUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidationFeatures`); value.Exists() {
		data.ValidationFeatures = &DmValidationFeatures{}
		data.ValidationFeatures.FromBody(ctx, "", value)
	} else {
		data.ValidationFeatures = nil
	}
	if value := res.Get(pathRoot + `RedirectURI`); value.Exists() {
		data.RedirectUri = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.RedirectUri = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `CustomScopeCheck`); value.Exists() {
		data.CustomScopeCheck = tfutils.BoolFromString(value.String())
	} else {
		data.CustomScopeCheck = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Scope`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Scope = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Scope = types.StringNull()
	}
	if value := res.Get(pathRoot + `ScopeUrl`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ScopeUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ScopeUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `DefaultScope`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DefaultScope = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DefaultScope = types.StringNull()
	}
	if value := res.Get(pathRoot + `TokenSecret`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TokenSecret = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TokenSecret = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalAZPageUrl`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalAzPageUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalAzPageUrl = types.StringValue("store:///OAuth-Generate-HTML.xsl")
	}
	if value := res.Get(pathRoot + `DPStateLifeTime`); value.Exists() {
		data.DpStateLifeTime = types.Int64Value(value.Int())
	} else {
		data.DpStateLifeTime = types.Int64Value(300)
	}
	if value := res.Get(pathRoot + `AUCodeLifeTime`); value.Exists() {
		data.AuCodeLifeTime = types.Int64Value(value.Int())
	} else {
		data.AuCodeLifeTime = types.Int64Value(300)
	}
	if value := res.Get(pathRoot + `AccessTokenLifeTime`); value.Exists() {
		data.AccessTokenLifeTime = types.Int64Value(value.Int())
	} else {
		data.AccessTokenLifeTime = types.Int64Value(3600)
	}
	if value := res.Get(pathRoot + `RefreshTokenAllowed`); value.Exists() {
		data.RefreshTokenAllowed = types.Int64Value(value.Int())
	} else {
		data.RefreshTokenAllowed = types.Int64Null()
	}
	if value := res.Get(pathRoot + `RefreshTokenLifeTime`); value.Exists() {
		data.RefreshTokenLifeTime = types.Int64Value(value.Int())
	} else {
		data.RefreshTokenLifeTime = types.Int64Value(5400)
	}
	if value := res.Get(pathRoot + `MaxConsentLifeTime`); value.Exists() {
		data.MaxConsentLifeTime = types.Int64Value(value.Int())
	} else {
		data.MaxConsentLifeTime = types.Int64Null()
	}
	if value := res.Get(pathRoot + `CustomResourceOwner`); value.Exists() {
		data.CustomResourceOwner = tfutils.BoolFromString(value.String())
	} else {
		data.CustomResourceOwner = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ResourceOwnerUrl`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ResourceOwnerUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ResourceOwnerUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AdditionalOAuthProcessUrl`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AdditionalOAuthProcessUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AdditionalOAuthProcessUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `RSSetHeader`); value.Exists() {
		data.RsSetHeader = &DmOAuthRSSetHeader{}
		data.RsSetHeader.FromBody(ctx, "", value)
	} else {
		data.RsSetHeader = nil
	}
	if value := res.Get(pathRoot + `ValidationURLSSLClientType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ValidationUrlsslClientType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ValidationUrlsslClientType = types.StringValue("client")
	}
	if value := res.Get(pathRoot + `ValidationURLSSLClient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ValidationUrlsslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ValidationUrlsslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `JWTGrantValidator`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.JwtGrantValidator = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JwtGrantValidator = types.StringNull()
	}
	if value := res.Get(pathRoot + `ClientJWTValidator`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ClientJwtValidator = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientJwtValidator = types.StringNull()
	}
	if value := res.Get(pathRoot + `OIDCIDTokenGenerator`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OidcidTokenGenerator = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OidcidTokenGenerator = types.StringNull()
	}
	if value := res.Get(pathRoot + `OAuthFeatures`); value.Exists() {
		data.OAuthFeatures = &DmOAuthFeatures{}
		data.OAuthFeatures.FromBody(ctx, "", value)
	} else {
		data.OAuthFeatures = nil
	}
}

func (data *OAuthSupportedClient) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Customized`); value.Exists() && !data.Customized.IsNull() {
		data.Customized = tfutils.BoolFromString(value.String())
	} else if data.Customized.ValueBool() {
		data.Customized = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CustomizedProcessUrl`); value.Exists() && !data.CustomizedProcessUrl.IsNull() {
		data.CustomizedProcessUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CustomizedProcessUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `OAuthRole`); value.Exists() {
		data.OAuthRole.UpdateFromBody(ctx, "", value)
	} else {
		data.OAuthRole = nil
	}
	if value := res.Get(pathRoot + `AZGrant`); value.Exists() {
		data.AzGrant.UpdateFromBody(ctx, "", value)
	} else {
		data.AzGrant = nil
	}
	if value := res.Get(pathRoot + `ClientType`); value.Exists() && !data.ClientType.IsNull() {
		data.ClientType = tfutils.ParseStringFromGJSON(value)
	} else if data.ClientType.ValueString() != "confidential" {
		data.ClientType = types.StringNull()
	}
	if value := res.Get(pathRoot + `CheckClientCredential`); value.Exists() && !data.CheckClientCredential.IsNull() {
		data.CheckClientCredential = tfutils.BoolFromString(value.String())
	} else if data.CheckClientCredential.ValueBool() {
		data.CheckClientCredential = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UseValidationUrl`); value.Exists() && !data.UseValidationUrl.IsNull() {
		data.UseValidationUrl = tfutils.BoolFromString(value.String())
	} else if data.UseValidationUrl.ValueBool() {
		data.UseValidationUrl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ClientAuthenMethod`); value.Exists() && !data.ClientAuthenMethod.IsNull() {
		data.ClientAuthenMethod = tfutils.ParseStringFromGJSON(value)
	} else if data.ClientAuthenMethod.ValueString() != "secret" {
		data.ClientAuthenMethod = types.StringNull()
	}
	if value := res.Get(pathRoot + `ClientValCred`); value.Exists() && !data.ClientValCred.IsNull() {
		data.ClientValCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientValCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `GenerateClientSecret`); value.Exists() && !data.GenerateClientSecret.IsNull() {
		data.GenerateClientSecret = tfutils.BoolFromString(value.String())
	} else if !data.GenerateClientSecret.ValueBool() {
		data.GenerateClientSecret = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ClientSecret`); value.Exists() && !data.ClientSecret.IsNull() {
		data.ClientSecret = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientSecret = types.StringNull()
	}
	if value := res.Get(pathRoot + `Caching`); value.Exists() && !data.Caching.IsNull() {
		data.Caching = tfutils.ParseStringFromGJSON(value)
	} else if data.Caching.ValueString() != "replay" {
		data.Caching = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidationURL`); value.Exists() && !data.ValidationUrl.IsNull() {
		data.ValidationUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ValidationUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidationFeatures`); value.Exists() {
		data.ValidationFeatures.UpdateFromBody(ctx, "", value)
	} else {
		data.ValidationFeatures = nil
	}
	if value := res.Get(pathRoot + `RedirectURI`); value.Exists() && !data.RedirectUri.IsNull() {
		data.RedirectUri = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.RedirectUri = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `CustomScopeCheck`); value.Exists() && !data.CustomScopeCheck.IsNull() {
		data.CustomScopeCheck = tfutils.BoolFromString(value.String())
	} else if data.CustomScopeCheck.ValueBool() {
		data.CustomScopeCheck = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Scope`); value.Exists() && !data.Scope.IsNull() {
		data.Scope = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Scope = types.StringNull()
	}
	if value := res.Get(pathRoot + `ScopeUrl`); value.Exists() && !data.ScopeUrl.IsNull() {
		data.ScopeUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ScopeUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `DefaultScope`); value.Exists() && !data.DefaultScope.IsNull() {
		data.DefaultScope = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DefaultScope = types.StringNull()
	}
	if value := res.Get(pathRoot + `TokenSecret`); value.Exists() && !data.TokenSecret.IsNull() {
		data.TokenSecret = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TokenSecret = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalAZPageUrl`); value.Exists() && !data.LocalAzPageUrl.IsNull() {
		data.LocalAzPageUrl = tfutils.ParseStringFromGJSON(value)
	} else if data.LocalAzPageUrl.ValueString() != "store:///OAuth-Generate-HTML.xsl" {
		data.LocalAzPageUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `DPStateLifeTime`); value.Exists() && !data.DpStateLifeTime.IsNull() {
		data.DpStateLifeTime = types.Int64Value(value.Int())
	} else if data.DpStateLifeTime.ValueInt64() != 300 {
		data.DpStateLifeTime = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AUCodeLifeTime`); value.Exists() && !data.AuCodeLifeTime.IsNull() {
		data.AuCodeLifeTime = types.Int64Value(value.Int())
	} else if data.AuCodeLifeTime.ValueInt64() != 300 {
		data.AuCodeLifeTime = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AccessTokenLifeTime`); value.Exists() && !data.AccessTokenLifeTime.IsNull() {
		data.AccessTokenLifeTime = types.Int64Value(value.Int())
	} else if data.AccessTokenLifeTime.ValueInt64() != 3600 {
		data.AccessTokenLifeTime = types.Int64Null()
	}
	if value := res.Get(pathRoot + `RefreshTokenAllowed`); value.Exists() && !data.RefreshTokenAllowed.IsNull() {
		data.RefreshTokenAllowed = types.Int64Value(value.Int())
	} else {
		data.RefreshTokenAllowed = types.Int64Null()
	}
	if value := res.Get(pathRoot + `RefreshTokenLifeTime`); value.Exists() && !data.RefreshTokenLifeTime.IsNull() {
		data.RefreshTokenLifeTime = types.Int64Value(value.Int())
	} else if data.RefreshTokenLifeTime.ValueInt64() != 5400 {
		data.RefreshTokenLifeTime = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxConsentLifeTime`); value.Exists() && !data.MaxConsentLifeTime.IsNull() {
		data.MaxConsentLifeTime = types.Int64Value(value.Int())
	} else {
		data.MaxConsentLifeTime = types.Int64Null()
	}
	if value := res.Get(pathRoot + `CustomResourceOwner`); value.Exists() && !data.CustomResourceOwner.IsNull() {
		data.CustomResourceOwner = tfutils.BoolFromString(value.String())
	} else if data.CustomResourceOwner.ValueBool() {
		data.CustomResourceOwner = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ResourceOwnerUrl`); value.Exists() && !data.ResourceOwnerUrl.IsNull() {
		data.ResourceOwnerUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ResourceOwnerUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AdditionalOAuthProcessUrl`); value.Exists() && !data.AdditionalOAuthProcessUrl.IsNull() {
		data.AdditionalOAuthProcessUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AdditionalOAuthProcessUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `RSSetHeader`); value.Exists() {
		data.RsSetHeader.UpdateFromBody(ctx, "", value)
	} else {
		data.RsSetHeader = nil
	}
	if value := res.Get(pathRoot + `ValidationURLSSLClientType`); value.Exists() && !data.ValidationUrlsslClientType.IsNull() {
		data.ValidationUrlsslClientType = tfutils.ParseStringFromGJSON(value)
	} else if data.ValidationUrlsslClientType.ValueString() != "client" {
		data.ValidationUrlsslClientType = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidationURLSSLClient`); value.Exists() && !data.ValidationUrlsslClient.IsNull() {
		data.ValidationUrlsslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ValidationUrlsslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `JWTGrantValidator`); value.Exists() && !data.JwtGrantValidator.IsNull() {
		data.JwtGrantValidator = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JwtGrantValidator = types.StringNull()
	}
	if value := res.Get(pathRoot + `ClientJWTValidator`); value.Exists() && !data.ClientJwtValidator.IsNull() {
		data.ClientJwtValidator = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientJwtValidator = types.StringNull()
	}
	if value := res.Get(pathRoot + `OIDCIDTokenGenerator`); value.Exists() && !data.OidcidTokenGenerator.IsNull() {
		data.OidcidTokenGenerator = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OidcidTokenGenerator = types.StringNull()
	}
	if value := res.Get(pathRoot + `OAuthFeatures`); value.Exists() {
		data.OAuthFeatures.UpdateFromBody(ctx, "", value)
	} else {
		data.OAuthFeatures = nil
	}
}
