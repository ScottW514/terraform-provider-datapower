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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type AAAJWTGenerator struct {
	Id                 types.String                `tfsdk:"id"`
	AppDomain          types.String                `tfsdk:"app_domain"`
	UserSummary        types.String                `tfsdk:"user_summary"`
	Issuer             types.String                `tfsdk:"issuer"`
	Duration           types.Int64                 `tfsdk:"duration"`
	AdditionalClaims   *DmJWTClaims                `tfsdk:"additional_claims"`
	Audience           types.List                  `tfsdk:"audience"`
	NotBefore          types.Int64                 `tfsdk:"not_before"`
	CustomClaims       types.String                `tfsdk:"custom_claims"`
	GenMethod          *DmJWTGenMethod             `tfsdk:"gen_method"`
	SignAlgorithm      types.String                `tfsdk:"sign_algorithm"`
	SignKey            types.String                `tfsdk:"sign_key"`
	SignSskey          types.String                `tfsdk:"sign_sskey"`
	EncAlgorithm       types.String                `tfsdk:"enc_algorithm"`
	EncryptAlgorithm   types.String                `tfsdk:"encrypt_algorithm"`
	EncryptCertificate types.String                `tfsdk:"encrypt_certificate"`
	EncryptSskey       types.String                `tfsdk:"encrypt_sskey"`
	DependencyActions  []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var AAAJWTGeneratorAudienceCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "additional_claims",
	AttrType:    "DmJWTClaims",
	AttrDefault: "",
	Value:       []string{"aud"},
}

var AAAJWTGeneratorAudienceIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var AAAJWTGeneratorNotBeforeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "additional_claims",
	AttrType:    "DmJWTClaims",
	AttrDefault: "",
	Value:       []string{"nbf"},
}

var AAAJWTGeneratorNotBeforeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var AAAJWTGeneratorCustomClaimsCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "additional_claims",
	AttrType:    "DmJWTClaims",
	AttrDefault: "",
	Value:       []string{"custom"},
}

var AAAJWTGeneratorCustomClaimsIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var AAAJWTGeneratorSignAlgorithmCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "gen_method",
	AttrType:    "DmJWTGenMethod",
	AttrDefault: "",
	Value:       []string{"sign"},
}

var AAAJWTGeneratorSignAlgorithmIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var AAAJWTGeneratorSignKeyCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "gen_method",
			AttrType:    "DmJWTGenMethod",
			AttrDefault: "",
			Value:       []string{"sign"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "sign_algorithm",
			AttrType:    "String",
			AttrDefault: "RS256",
			Value:       []string{"RS256", "RS384", "RS512", "PS256", "PS384", "PS512", "ES256", "ES384", "ES512"},
		},
	},
}

var AAAJWTGeneratorSignKeyIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var AAAJWTGeneratorSignSSKeyCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "gen_method",
			AttrType:    "DmJWTGenMethod",
			AttrDefault: "",
			Value:       []string{"sign"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "sign_algorithm",
			AttrType:    "String",
			AttrDefault: "RS256",
			Value:       []string{"HS256", "HS384", "HS512"},
		},
	},
}

var AAAJWTGeneratorSignSSKeyIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var AAAJWTGeneratorEncAlgorithmCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "gen_method",
	AttrType:    "DmJWTGenMethod",
	AttrDefault: "",
	Value:       []string{"encrypt"},
}

var AAAJWTGeneratorEncAlgorithmIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var AAAJWTGeneratorEncryptAlgorithmCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "gen_method",
	AttrType:    "DmJWTGenMethod",
	AttrDefault: "",
	Value:       []string{"encrypt"},
}

var AAAJWTGeneratorEncryptAlgorithmIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var AAAJWTGeneratorEncryptCertificateCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "gen_method",
			AttrType:    "DmJWTGenMethod",
			AttrDefault: "",
			Value:       []string{"encrypt"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "encrypt_algorithm",
			AttrType:    "String",
			AttrDefault: "RSA1_5",
			Value:       []string{"RSA1_5", "RSA-OAEP", "RSA-OAEP-256"},
		},
	},
}

var AAAJWTGeneratorEncryptCertificateIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var AAAJWTGeneratorEncryptSSKeyCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "gen_method",
			AttrType:    "DmJWTGenMethod",
			AttrDefault: "",
			Value:       []string{"encrypt"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "encrypt_algorithm",
			AttrType:    "String",
			AttrDefault: "RSA1_5",
			Value:       []string{"A128KW", "A192KW", "A256KW", "dir"},
		},
	},
}

var AAAJWTGeneratorEncryptSSKeyIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var AAAJWTGeneratorObjectType = map[string]attr.Type{
	"id":                  types.StringType,
	"app_domain":          types.StringType,
	"user_summary":        types.StringType,
	"issuer":              types.StringType,
	"duration":            types.Int64Type,
	"additional_claims":   types.ObjectType{AttrTypes: DmJWTClaimsObjectType},
	"audience":            types.ListType{ElemType: types.StringType},
	"not_before":          types.Int64Type,
	"custom_claims":       types.StringType,
	"gen_method":          types.ObjectType{AttrTypes: DmJWTGenMethodObjectType},
	"sign_algorithm":      types.StringType,
	"sign_key":            types.StringType,
	"sign_sskey":          types.StringType,
	"enc_algorithm":       types.StringType,
	"encrypt_algorithm":   types.StringType,
	"encrypt_certificate": types.StringType,
	"encrypt_sskey":       types.StringType,
	"dependency_actions":  actions.ActionsListType,
}

func (data AAAJWTGenerator) GetPath() string {
	rest_path := "/mgmt/config/{domain}/AAAJWTGenerator"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data AAAJWTGenerator) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Issuer.IsNull() {
		return false
	}
	if !data.Duration.IsNull() {
		return false
	}
	if data.AdditionalClaims != nil {
		if !data.AdditionalClaims.IsNull() {
			return false
		}
	}
	if !data.Audience.IsNull() {
		return false
	}
	if !data.NotBefore.IsNull() {
		return false
	}
	if !data.CustomClaims.IsNull() {
		return false
	}
	if data.GenMethod != nil {
		if !data.GenMethod.IsNull() {
			return false
		}
	}
	if !data.SignAlgorithm.IsNull() {
		return false
	}
	if !data.SignKey.IsNull() {
		return false
	}
	if !data.SignSskey.IsNull() {
		return false
	}
	if !data.EncAlgorithm.IsNull() {
		return false
	}
	if !data.EncryptAlgorithm.IsNull() {
		return false
	}
	if !data.EncryptCertificate.IsNull() {
		return false
	}
	if !data.EncryptSskey.IsNull() {
		return false
	}
	return true
}

func (data AAAJWTGenerator) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Issuer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Issuer`, data.Issuer.ValueString())
	}
	if !data.Duration.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Duration`, data.Duration.ValueInt64())
	}
	if data.AdditionalClaims != nil {
		if !data.AdditionalClaims.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`AdditionalClaims`, data.AdditionalClaims.ToBody(ctx, ""))
		}
	}
	if !data.Audience.IsNull() {
		var dataValues []string
		data.Audience.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`Audience`+".-1", map[string]string{"value": val})
		}
	}
	if !data.NotBefore.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`NotBefore`, data.NotBefore.ValueInt64())
	}
	if !data.CustomClaims.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CustomClaims`, data.CustomClaims.ValueString())
	}
	if data.GenMethod != nil {
		if !data.GenMethod.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`GenMethod`, data.GenMethod.ToBody(ctx, ""))
		}
	}
	if !data.SignAlgorithm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SignAlgorithm`, data.SignAlgorithm.ValueString())
	}
	if !data.SignKey.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SignKey`, data.SignKey.ValueString())
	}
	if !data.SignSskey.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SignSSKey`, data.SignSskey.ValueString())
	}
	if !data.EncAlgorithm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EncAlgorithm`, data.EncAlgorithm.ValueString())
	}
	if !data.EncryptAlgorithm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EncryptAlgorithm`, data.EncryptAlgorithm.ValueString())
	}
	if !data.EncryptCertificate.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EncryptCertificate`, data.EncryptCertificate.ValueString())
	}
	if !data.EncryptSskey.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EncryptSSKey`, data.EncryptSskey.ValueString())
	}
	return body
}

func (data *AAAJWTGenerator) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Issuer`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Issuer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Issuer = types.StringValue("idg")
	}
	if value := res.Get(pathRoot + `Duration`); value.Exists() {
		data.Duration = types.Int64Value(value.Int())
	} else {
		data.Duration = types.Int64Value(3600)
	}
	if value := res.Get(pathRoot + `AdditionalClaims`); value.Exists() {
		data.AdditionalClaims = &DmJWTClaims{}
		data.AdditionalClaims.FromBody(ctx, "", value)
	} else {
		data.AdditionalClaims = nil
	}
	if value := res.Get(pathRoot + `Audience`); value.Exists() {
		data.Audience = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Audience = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `NotBefore`); value.Exists() {
		data.NotBefore = types.Int64Value(value.Int())
	} else {
		data.NotBefore = types.Int64Null()
	}
	if value := res.Get(pathRoot + `CustomClaims`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CustomClaims = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CustomClaims = types.StringNull()
	}
	if value := res.Get(pathRoot + `GenMethod`); value.Exists() {
		data.GenMethod = &DmJWTGenMethod{}
		data.GenMethod.FromBody(ctx, "", value)
	} else {
		data.GenMethod = nil
	}
	if value := res.Get(pathRoot + `SignAlgorithm`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SignAlgorithm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SignAlgorithm = types.StringValue("RS256")
	}
	if value := res.Get(pathRoot + `SignKey`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SignKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SignKey = types.StringNull()
	}
	if value := res.Get(pathRoot + `SignSSKey`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SignSskey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SignSskey = types.StringNull()
	}
	if value := res.Get(pathRoot + `EncAlgorithm`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EncAlgorithm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EncAlgorithm = types.StringValue("A128CBC-HS256")
	}
	if value := res.Get(pathRoot + `EncryptAlgorithm`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EncryptAlgorithm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EncryptAlgorithm = types.StringValue("RSA1_5")
	}
	if value := res.Get(pathRoot + `EncryptCertificate`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EncryptCertificate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EncryptCertificate = types.StringNull()
	}
	if value := res.Get(pathRoot + `EncryptSSKey`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EncryptSskey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EncryptSskey = types.StringNull()
	}
}

func (data *AAAJWTGenerator) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Issuer`); value.Exists() && !data.Issuer.IsNull() {
		data.Issuer = tfutils.ParseStringFromGJSON(value)
	} else if data.Issuer.ValueString() != "idg" {
		data.Issuer = types.StringNull()
	}
	if value := res.Get(pathRoot + `Duration`); value.Exists() && !data.Duration.IsNull() {
		data.Duration = types.Int64Value(value.Int())
	} else if data.Duration.ValueInt64() != 3600 {
		data.Duration = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AdditionalClaims`); value.Exists() {
		data.AdditionalClaims.UpdateFromBody(ctx, "", value)
	} else {
		data.AdditionalClaims = nil
	}
	if value := res.Get(pathRoot + `Audience`); value.Exists() && !data.Audience.IsNull() {
		data.Audience = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Audience = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `NotBefore`); value.Exists() && !data.NotBefore.IsNull() {
		data.NotBefore = types.Int64Value(value.Int())
	} else {
		data.NotBefore = types.Int64Null()
	}
	if value := res.Get(pathRoot + `CustomClaims`); value.Exists() && !data.CustomClaims.IsNull() {
		data.CustomClaims = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CustomClaims = types.StringNull()
	}
	if value := res.Get(pathRoot + `GenMethod`); value.Exists() {
		data.GenMethod.UpdateFromBody(ctx, "", value)
	} else {
		data.GenMethod = nil
	}
	if value := res.Get(pathRoot + `SignAlgorithm`); value.Exists() && !data.SignAlgorithm.IsNull() {
		data.SignAlgorithm = tfutils.ParseStringFromGJSON(value)
	} else if data.SignAlgorithm.ValueString() != "RS256" {
		data.SignAlgorithm = types.StringNull()
	}
	if value := res.Get(pathRoot + `SignKey`); value.Exists() && !data.SignKey.IsNull() {
		data.SignKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SignKey = types.StringNull()
	}
	if value := res.Get(pathRoot + `SignSSKey`); value.Exists() && !data.SignSskey.IsNull() {
		data.SignSskey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SignSskey = types.StringNull()
	}
	if value := res.Get(pathRoot + `EncAlgorithm`); value.Exists() && !data.EncAlgorithm.IsNull() {
		data.EncAlgorithm = tfutils.ParseStringFromGJSON(value)
	} else if data.EncAlgorithm.ValueString() != "A128CBC-HS256" {
		data.EncAlgorithm = types.StringNull()
	}
	if value := res.Get(pathRoot + `EncryptAlgorithm`); value.Exists() && !data.EncryptAlgorithm.IsNull() {
		data.EncryptAlgorithm = tfutils.ParseStringFromGJSON(value)
	} else if data.EncryptAlgorithm.ValueString() != "RSA1_5" {
		data.EncryptAlgorithm = types.StringNull()
	}
	if value := res.Get(pathRoot + `EncryptCertificate`); value.Exists() && !data.EncryptCertificate.IsNull() {
		data.EncryptCertificate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EncryptCertificate = types.StringNull()
	}
	if value := res.Get(pathRoot + `EncryptSSKey`); value.Exists() && !data.EncryptSskey.IsNull() {
		data.EncryptSskey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EncryptSskey = types.StringNull()
	}
}
