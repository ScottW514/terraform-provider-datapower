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

type AssemblyActionJWTValidate struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	Jwt               types.String                `tfsdk:"jwt"`
	OutputClaims      types.String                `tfsdk:"output_claims"`
	IssuerClaim       types.String                `tfsdk:"issuer_claim"`
	AudienceClaim     types.String                `tfsdk:"audience_claim"`
	DecryptCrypto     types.String                `tfsdk:"decrypt_crypto"`
	DecryptJwk        types.String                `tfsdk:"decrypt_jwk"`
	VerifyCrypto      types.String                `tfsdk:"verify_crypto"`
	VerifyJwk         types.String                `tfsdk:"verify_jwk"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	Title             types.String                `tfsdk:"title"`
	CorrelationPath   types.String                `tfsdk:"correlation_path"`
	ActionDebug       types.Bool                  `tfsdk:"action_debug"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
	ProviderTarget    types.String                `tfsdk:"provider_target"`
}

var AssemblyActionJWTValidateObjectType = map[string]attr.Type{
	"provider_target":    types.StringType,
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"jwt":                types.StringType,
	"output_claims":      types.StringType,
	"issuer_claim":       types.StringType,
	"audience_claim":     types.StringType,
	"decrypt_crypto":     types.StringType,
	"decrypt_jwk":        types.StringType,
	"verify_crypto":      types.StringType,
	"verify_jwk":         types.StringType,
	"user_summary":       types.StringType,
	"title":              types.StringType,
	"correlation_path":   types.StringType,
	"action_debug":       types.BoolType,
	"dependency_actions": actions.ActionsListType,
}

func (data AssemblyActionJWTValidate) GetPath() string {
	rest_path := "/mgmt/config/{domain}/AssemblyActionJWTValidate"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data AssemblyActionJWTValidate) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Jwt.IsNull() {
		return false
	}
	if !data.OutputClaims.IsNull() {
		return false
	}
	if !data.IssuerClaim.IsNull() {
		return false
	}
	if !data.AudienceClaim.IsNull() {
		return false
	}
	if !data.DecryptCrypto.IsNull() {
		return false
	}
	if !data.DecryptJwk.IsNull() {
		return false
	}
	if !data.VerifyCrypto.IsNull() {
		return false
	}
	if !data.VerifyJwk.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Title.IsNull() {
		return false
	}
	if !data.CorrelationPath.IsNull() {
		return false
	}
	if !data.ActionDebug.IsNull() {
		return false
	}
	return true
}

func (data AssemblyActionJWTValidate) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.Jwt.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`JWT`, data.Jwt.ValueString())
	}
	if !data.OutputClaims.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OutputClaims`, data.OutputClaims.ValueString())
	}
	if !data.IssuerClaim.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IssuerClaim`, data.IssuerClaim.ValueString())
	}
	if !data.AudienceClaim.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AudienceClaim`, data.AudienceClaim.ValueString())
	}
	if !data.DecryptCrypto.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DecryptCrypto`, data.DecryptCrypto.ValueString())
	}
	if !data.DecryptJwk.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DecryptJWK`, data.DecryptJwk.ValueString())
	}
	if !data.VerifyCrypto.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`VerifyCrypto`, data.VerifyCrypto.ValueString())
	}
	if !data.VerifyJwk.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`VerifyJWK`, data.VerifyJwk.ValueString())
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.Title.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Title`, data.Title.ValueString())
	}
	if !data.CorrelationPath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CorrelationPath`, data.CorrelationPath.ValueString())
	}
	if !data.ActionDebug.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ActionDebug`, tfutils.StringFromBool(data.ActionDebug, ""))
	}
	return body
}

func (data *AssemblyActionJWTValidate) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `JWT`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Jwt = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Jwt = types.StringValue("request.headers.authorization")
	}
	if value := res.Get(pathRoot + `OutputClaims`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OutputClaims = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OutputClaims = types.StringValue("decoded.claims")
	}
	if value := res.Get(pathRoot + `IssuerClaim`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.IssuerClaim = tfutils.ParseStringFromGJSON(value)
	} else {
		data.IssuerClaim = types.StringNull()
	}
	if value := res.Get(pathRoot + `AudienceClaim`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AudienceClaim = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AudienceClaim = types.StringNull()
	}
	if value := res.Get(pathRoot + `DecryptCrypto`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DecryptCrypto = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DecryptCrypto = types.StringNull()
	}
	if value := res.Get(pathRoot + `DecryptJWK`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DecryptJwk = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DecryptJwk = types.StringNull()
	}
	if value := res.Get(pathRoot + `VerifyCrypto`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.VerifyCrypto = tfutils.ParseStringFromGJSON(value)
	} else {
		data.VerifyCrypto = types.StringNull()
	}
	if value := res.Get(pathRoot + `VerifyJWK`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.VerifyJwk = tfutils.ParseStringFromGJSON(value)
	} else {
		data.VerifyJwk = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `Title`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Title = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Title = types.StringNull()
	}
	if value := res.Get(pathRoot + `CorrelationPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CorrelationPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CorrelationPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `ActionDebug`); value.Exists() {
		data.ActionDebug = tfutils.BoolFromString(value.String())
	} else {
		data.ActionDebug = types.BoolNull()
	}
}

func (data *AssemblyActionJWTValidate) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `JWT`); value.Exists() && !data.Jwt.IsNull() {
		data.Jwt = tfutils.ParseStringFromGJSON(value)
	} else if data.Jwt.ValueString() != "request.headers.authorization" {
		data.Jwt = types.StringNull()
	}
	if value := res.Get(pathRoot + `OutputClaims`); value.Exists() && !data.OutputClaims.IsNull() {
		data.OutputClaims = tfutils.ParseStringFromGJSON(value)
	} else if data.OutputClaims.ValueString() != "decoded.claims" {
		data.OutputClaims = types.StringNull()
	}
	if value := res.Get(pathRoot + `IssuerClaim`); value.Exists() && !data.IssuerClaim.IsNull() {
		data.IssuerClaim = tfutils.ParseStringFromGJSON(value)
	} else {
		data.IssuerClaim = types.StringNull()
	}
	if value := res.Get(pathRoot + `AudienceClaim`); value.Exists() && !data.AudienceClaim.IsNull() {
		data.AudienceClaim = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AudienceClaim = types.StringNull()
	}
	if value := res.Get(pathRoot + `DecryptCrypto`); value.Exists() && !data.DecryptCrypto.IsNull() {
		data.DecryptCrypto = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DecryptCrypto = types.StringNull()
	}
	if value := res.Get(pathRoot + `DecryptJWK`); value.Exists() && !data.DecryptJwk.IsNull() {
		data.DecryptJwk = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DecryptJwk = types.StringNull()
	}
	if value := res.Get(pathRoot + `VerifyCrypto`); value.Exists() && !data.VerifyCrypto.IsNull() {
		data.VerifyCrypto = tfutils.ParseStringFromGJSON(value)
	} else {
		data.VerifyCrypto = types.StringNull()
	}
	if value := res.Get(pathRoot + `VerifyJWK`); value.Exists() && !data.VerifyJwk.IsNull() {
		data.VerifyJwk = tfutils.ParseStringFromGJSON(value)
	} else {
		data.VerifyJwk = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `Title`); value.Exists() && !data.Title.IsNull() {
		data.Title = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Title = types.StringNull()
	}
	if value := res.Get(pathRoot + `CorrelationPath`); value.Exists() && !data.CorrelationPath.IsNull() {
		data.CorrelationPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CorrelationPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `ActionDebug`); value.Exists() && !data.ActionDebug.IsNull() {
		data.ActionDebug = tfutils.BoolFromString(value.String())
	} else if data.ActionDebug.ValueBool() {
		data.ActionDebug = types.BoolNull()
	}
}
