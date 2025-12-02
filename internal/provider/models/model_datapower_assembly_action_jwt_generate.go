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

type AssemblyActionJWTGenerate struct {
	Id                  types.String                `tfsdk:"id"`
	AppDomain           types.String                `tfsdk:"app_domain"`
	Jwt                 types.String                `tfsdk:"jwt"`
	JwtIdClaims         types.Bool                  `tfsdk:"jwt_id_claims"`
	IssuerClaim         types.String                `tfsdk:"issuer_claim"`
	SubjectClaim        types.String                `tfsdk:"subject_claim"`
	AudienceClaim       types.String                `tfsdk:"audience_claim"`
	ValidityPeriod      types.Int64                 `tfsdk:"validity_period"`
	PrivateClaim        types.String                `tfsdk:"private_claim"`
	SignJwk             types.String                `tfsdk:"sign_jwk"`
	CryptoAlgorithm     types.String                `tfsdk:"crypto_algorithm"`
	SignCrypto          types.String                `tfsdk:"sign_crypto"`
	CustomKidValueJws   types.String                `tfsdk:"custom_kid_value_jws"`
	EncryptAlgorithm    types.String                `tfsdk:"encrypt_algorithm"`
	EncryptJwk          types.String                `tfsdk:"encrypt_jwk"`
	KeyEncryptAlgorithm types.String                `tfsdk:"key_encrypt_algorithm"`
	EncryptCrypto       types.String                `tfsdk:"encrypt_crypto"`
	CustomKidValueJwe   types.String                `tfsdk:"custom_kid_value_jwe"`
	UserSummary         types.String                `tfsdk:"user_summary"`
	Title               types.String                `tfsdk:"title"`
	CorrelationPath     types.String                `tfsdk:"correlation_path"`
	ActionDebug         types.Bool                  `tfsdk:"action_debug"`
	DependencyActions   []*actions.DependencyAction `tfsdk:"dependency_actions"`
	ProviderTarget      types.String                `tfsdk:"provider_target"`
}

var AssemblyActionJWTGenerateObjectType = map[string]attr.Type{
	"provider_target":       types.StringType,
	"id":                    types.StringType,
	"app_domain":            types.StringType,
	"jwt":                   types.StringType,
	"jwt_id_claims":         types.BoolType,
	"issuer_claim":          types.StringType,
	"subject_claim":         types.StringType,
	"audience_claim":        types.StringType,
	"validity_period":       types.Int64Type,
	"private_claim":         types.StringType,
	"sign_jwk":              types.StringType,
	"crypto_algorithm":      types.StringType,
	"sign_crypto":           types.StringType,
	"custom_kid_value_jws":  types.StringType,
	"encrypt_algorithm":     types.StringType,
	"encrypt_jwk":           types.StringType,
	"key_encrypt_algorithm": types.StringType,
	"encrypt_crypto":        types.StringType,
	"custom_kid_value_jwe":  types.StringType,
	"user_summary":          types.StringType,
	"title":                 types.StringType,
	"correlation_path":      types.StringType,
	"action_debug":          types.BoolType,
	"dependency_actions":    actions.ActionsListType,
}

func (data AssemblyActionJWTGenerate) GetPath() string {
	rest_path := "/mgmt/config/{domain}/AssemblyActionJWTGenerate"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data AssemblyActionJWTGenerate) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Jwt.IsNull() {
		return false
	}
	if !data.JwtIdClaims.IsNull() {
		return false
	}
	if !data.IssuerClaim.IsNull() {
		return false
	}
	if !data.SubjectClaim.IsNull() {
		return false
	}
	if !data.AudienceClaim.IsNull() {
		return false
	}
	if !data.ValidityPeriod.IsNull() {
		return false
	}
	if !data.PrivateClaim.IsNull() {
		return false
	}
	if !data.SignJwk.IsNull() {
		return false
	}
	if !data.CryptoAlgorithm.IsNull() {
		return false
	}
	if !data.SignCrypto.IsNull() {
		return false
	}
	if !data.CustomKidValueJws.IsNull() {
		return false
	}
	if !data.EncryptAlgorithm.IsNull() {
		return false
	}
	if !data.EncryptJwk.IsNull() {
		return false
	}
	if !data.KeyEncryptAlgorithm.IsNull() {
		return false
	}
	if !data.EncryptCrypto.IsNull() {
		return false
	}
	if !data.CustomKidValueJwe.IsNull() {
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

func (data AssemblyActionJWTGenerate) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.JwtIdClaims.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`JWTIDClaims`, tfutils.StringFromBool(data.JwtIdClaims, ""))
	}
	if !data.IssuerClaim.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IssuerClaim`, data.IssuerClaim.ValueString())
	}
	if !data.SubjectClaim.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SubjectClaim`, data.SubjectClaim.ValueString())
	}
	if !data.AudienceClaim.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AudienceClaim`, data.AudienceClaim.ValueString())
	}
	if !data.ValidityPeriod.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ValidityPeriod`, data.ValidityPeriod.ValueInt64())
	}
	if !data.PrivateClaim.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PrivateClaim`, data.PrivateClaim.ValueString())
	}
	if !data.SignJwk.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SignJWK`, data.SignJwk.ValueString())
	}
	if !data.CryptoAlgorithm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CryptoAlgorithm`, data.CryptoAlgorithm.ValueString())
	}
	if !data.SignCrypto.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SignCrypto`, data.SignCrypto.ValueString())
	}
	if !data.CustomKidValueJws.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CustomKIDValueJWS`, data.CustomKidValueJws.ValueString())
	}
	if !data.EncryptAlgorithm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EncryptAlgorithm`, data.EncryptAlgorithm.ValueString())
	}
	if !data.EncryptJwk.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EncryptJWK`, data.EncryptJwk.ValueString())
	}
	if !data.KeyEncryptAlgorithm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`KeyEncryptAlgorithm`, data.KeyEncryptAlgorithm.ValueString())
	}
	if !data.EncryptCrypto.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EncryptCrypto`, data.EncryptCrypto.ValueString())
	}
	if !data.CustomKidValueJwe.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CustomKIDValueJWE`, data.CustomKidValueJwe.ValueString())
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

func (data *AssemblyActionJWTGenerate) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
		data.Jwt = types.StringValue("generated.jwt")
	}
	if value := res.Get(pathRoot + `JWTIDClaims`); value.Exists() {
		data.JwtIdClaims = tfutils.BoolFromString(value.String())
	} else {
		data.JwtIdClaims = types.BoolNull()
	}
	if value := res.Get(pathRoot + `IssuerClaim`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.IssuerClaim = tfutils.ParseStringFromGJSON(value)
	} else {
		data.IssuerClaim = types.StringValue("iss.claim")
	}
	if value := res.Get(pathRoot + `SubjectClaim`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SubjectClaim = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SubjectClaim = types.StringNull()
	}
	if value := res.Get(pathRoot + `AudienceClaim`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AudienceClaim = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AudienceClaim = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidityPeriod`); value.Exists() {
		data.ValidityPeriod = types.Int64Value(value.Int())
	} else {
		data.ValidityPeriod = types.Int64Value(3600)
	}
	if value := res.Get(pathRoot + `PrivateClaim`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PrivateClaim = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PrivateClaim = types.StringNull()
	}
	if value := res.Get(pathRoot + `SignJWK`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SignJwk = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SignJwk = types.StringNull()
	}
	if value := res.Get(pathRoot + `CryptoAlgorithm`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CryptoAlgorithm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CryptoAlgorithm = types.StringNull()
	}
	if value := res.Get(pathRoot + `SignCrypto`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SignCrypto = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SignCrypto = types.StringNull()
	}
	if value := res.Get(pathRoot + `CustomKIDValueJWS`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CustomKidValueJws = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CustomKidValueJws = types.StringNull()
	}
	if value := res.Get(pathRoot + `EncryptAlgorithm`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EncryptAlgorithm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EncryptAlgorithm = types.StringNull()
	}
	if value := res.Get(pathRoot + `EncryptJWK`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EncryptJwk = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EncryptJwk = types.StringNull()
	}
	if value := res.Get(pathRoot + `KeyEncryptAlgorithm`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.KeyEncryptAlgorithm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.KeyEncryptAlgorithm = types.StringNull()
	}
	if value := res.Get(pathRoot + `EncryptCrypto`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EncryptCrypto = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EncryptCrypto = types.StringNull()
	}
	if value := res.Get(pathRoot + `CustomKIDValueJWE`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CustomKidValueJwe = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CustomKidValueJwe = types.StringNull()
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

func (data *AssemblyActionJWTGenerate) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	} else if data.Jwt.ValueString() != "generated.jwt" {
		data.Jwt = types.StringNull()
	}
	if value := res.Get(pathRoot + `JWTIDClaims`); value.Exists() && !data.JwtIdClaims.IsNull() {
		data.JwtIdClaims = tfutils.BoolFromString(value.String())
	} else if data.JwtIdClaims.ValueBool() {
		data.JwtIdClaims = types.BoolNull()
	}
	if value := res.Get(pathRoot + `IssuerClaim`); value.Exists() && !data.IssuerClaim.IsNull() {
		data.IssuerClaim = tfutils.ParseStringFromGJSON(value)
	} else if data.IssuerClaim.ValueString() != "iss.claim" {
		data.IssuerClaim = types.StringNull()
	}
	if value := res.Get(pathRoot + `SubjectClaim`); value.Exists() && !data.SubjectClaim.IsNull() {
		data.SubjectClaim = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SubjectClaim = types.StringNull()
	}
	if value := res.Get(pathRoot + `AudienceClaim`); value.Exists() && !data.AudienceClaim.IsNull() {
		data.AudienceClaim = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AudienceClaim = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidityPeriod`); value.Exists() && !data.ValidityPeriod.IsNull() {
		data.ValidityPeriod = types.Int64Value(value.Int())
	} else if data.ValidityPeriod.ValueInt64() != 3600 {
		data.ValidityPeriod = types.Int64Null()
	}
	if value := res.Get(pathRoot + `PrivateClaim`); value.Exists() && !data.PrivateClaim.IsNull() {
		data.PrivateClaim = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PrivateClaim = types.StringNull()
	}
	if value := res.Get(pathRoot + `SignJWK`); value.Exists() && !data.SignJwk.IsNull() {
		data.SignJwk = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SignJwk = types.StringNull()
	}
	if value := res.Get(pathRoot + `CryptoAlgorithm`); value.Exists() && !data.CryptoAlgorithm.IsNull() {
		data.CryptoAlgorithm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CryptoAlgorithm = types.StringNull()
	}
	if value := res.Get(pathRoot + `SignCrypto`); value.Exists() && !data.SignCrypto.IsNull() {
		data.SignCrypto = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SignCrypto = types.StringNull()
	}
	if value := res.Get(pathRoot + `CustomKIDValueJWS`); value.Exists() && !data.CustomKidValueJws.IsNull() {
		data.CustomKidValueJws = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CustomKidValueJws = types.StringNull()
	}
	if value := res.Get(pathRoot + `EncryptAlgorithm`); value.Exists() && !data.EncryptAlgorithm.IsNull() {
		data.EncryptAlgorithm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EncryptAlgorithm = types.StringNull()
	}
	if value := res.Get(pathRoot + `EncryptJWK`); value.Exists() && !data.EncryptJwk.IsNull() {
		data.EncryptJwk = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EncryptJwk = types.StringNull()
	}
	if value := res.Get(pathRoot + `KeyEncryptAlgorithm`); value.Exists() && !data.KeyEncryptAlgorithm.IsNull() {
		data.KeyEncryptAlgorithm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.KeyEncryptAlgorithm = types.StringNull()
	}
	if value := res.Get(pathRoot + `EncryptCrypto`); value.Exists() && !data.EncryptCrypto.IsNull() {
		data.EncryptCrypto = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EncryptCrypto = types.StringNull()
	}
	if value := res.Get(pathRoot + `CustomKIDValueJWE`); value.Exists() && !data.CustomKidValueJwe.IsNull() {
		data.CustomKidValueJwe = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CustomKidValueJwe = types.StringNull()
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
