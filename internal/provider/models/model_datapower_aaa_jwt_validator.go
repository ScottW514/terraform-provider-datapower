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

type AAAJWTValidator struct {
	Id                              types.String                `tfsdk:"id"`
	AppDomain                       types.String                `tfsdk:"app_domain"`
	UserSummary                     types.String                `tfsdk:"user_summary"`
	Issuer                          types.String                `tfsdk:"issuer"`
	Aud                             types.String                `tfsdk:"aud"`
	ValMethod                       *DmJWTValMethod             `tfsdk:"val_method"`
	CustomizedScript                types.String                `tfsdk:"customized_script"`
	DecryptCredentialType           types.String                `tfsdk:"decrypt_credential_type"`
	DecryptKey                      types.String                `tfsdk:"decrypt_key"`
	DecryptSSecret                  types.String                `tfsdk:"decrypt_s_secret"`
	DecryptJwk                      types.String                `tfsdk:"decrypt_jwk"`
	DecryptFetchCredUrl             types.String                `tfsdk:"decrypt_fetch_cred_url"`
	DecryptFetchCredSslProfile      types.String                `tfsdk:"decrypt_fetch_cred_ssl_profile"`
	ValidateCustom                  types.String                `tfsdk:"validate_custom"`
	VerifyCredentialType            types.String                `tfsdk:"verify_credential_type"`
	VerifyCertificate               types.String                `tfsdk:"verify_certificate"`
	VerifyCertificateAgainstValCred types.Bool                  `tfsdk:"verify_certificate_against_val_cred"`
	VerifyValCred                   types.String                `tfsdk:"verify_val_cred"`
	VerifySSecret                   types.String                `tfsdk:"verify_s_secret"`
	VerifyJwk                       types.String                `tfsdk:"verify_jwk"`
	VerifyFetchCredUrl              types.String                `tfsdk:"verify_fetch_cred_url"`
	VerifyFetchCredSslProfile       types.String                `tfsdk:"verify_fetch_cred_ssl_profile"`
	Claims                          types.List                  `tfsdk:"claims"`
	UsernameClaim                   types.String                `tfsdk:"username_claim"`
	DependencyActions               []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var AAAJWTValidatorCustomizedScriptCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "val_method",
	AttrType:    "DmJWTValMethod",
	AttrDefault: "",
	Value:       []string{"customized"},
}
var AAAJWTValidatorDecryptCredentialTypeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "val_method",
	AttrType:    "DmJWTValMethod",
	AttrDefault: "",
	Value:       []string{"decrypt"},
}
var AAAJWTValidatorDecryptKeyCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "val_method",
			AttrType:    "DmJWTValMethod",
			AttrDefault: "",
			Value:       []string{"decrypt"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "decrypt_credential_type",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"pkix"},
		},
	},
}
var AAAJWTValidatorDecryptSSecretCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "val_method",
			AttrType:    "DmJWTValMethod",
			AttrDefault: "",
			Value:       []string{"decrypt"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "decrypt_credential_type",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"ssecret"},
		},
	},
}
var AAAJWTValidatorDecryptJWKCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "val_method",
			AttrType:    "DmJWTValMethod",
			AttrDefault: "",
			Value:       []string{"decrypt"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "decrypt_credential_type",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"jwk"},
		},
	},
}
var AAAJWTValidatorDecryptFetchCredURLCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "val_method",
			AttrType:    "DmJWTValMethod",
			AttrDefault: "",
			Value:       []string{"decrypt"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "decrypt_credential_type",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"jwk-remote"},
		},
	},
}
var AAAJWTValidatorValidateCustomCondVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "val_method",
					AttrType:    "DmJWTValMethod",
					AttrDefault: "",
					Value:       []string{"verify"},
				},
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "verify_credential_type",
					AttrType:    "String",
					AttrDefault: "",
					Value:       []string{"custom"},
				},
			},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "val_method",
					AttrType:    "DmJWTValMethod",
					AttrDefault: "",
					Value:       []string{"decrypt"},
				},
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "decrypt_credential_type",
					AttrType:    "String",
					AttrDefault: "",
					Value:       []string{"custom"},
				},
			},
		},
	},
}
var AAAJWTValidatorVerifyCredentialTypeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "val_method",
	AttrType:    "DmJWTValMethod",
	AttrDefault: "",
	Value:       []string{"verify"},
}
var AAAJWTValidatorVerifyCertificateCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "val_method",
			AttrType:    "DmJWTValMethod",
			AttrDefault: "",
			Value:       []string{"verify"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "verify_credential_type",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"pkix"},
		},
	},
}
var AAAJWTValidatorVerifyCertificateAgainstValCredCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "val_method",
			AttrType:    "DmJWTValMethod",
			AttrDefault: "",
			Value:       []string{"verify"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "verify_credential_type",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"pkix"},
		},
	},
}
var AAAJWTValidatorVerifyValCredCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "val_method",
			AttrType:    "DmJWTValMethod",
			AttrDefault: "",
			Value:       []string{"verify"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "verify_credential_type",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"pkix"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "verify_certificate_against_val_cred",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
	},
}
var AAAJWTValidatorVerifySSecretCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "val_method",
			AttrType:    "DmJWTValMethod",
			AttrDefault: "",
			Value:       []string{"verify"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "verify_credential_type",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"ssecret"},
		},
	},
}
var AAAJWTValidatorVerifyJWKCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "val_method",
			AttrType:    "DmJWTValMethod",
			AttrDefault: "",
			Value:       []string{"verify"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "verify_credential_type",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"jwk"},
		},
	},
}
var AAAJWTValidatorVerifyFetchCredURLCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "val_method",
			AttrType:    "DmJWTValMethod",
			AttrDefault: "",
			Value:       []string{"verify"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "verify_credential_type",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"jwk-remote"},
		},
	},
}
var AAAJWTValidatorCustomizedScriptIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var AAAJWTValidatorDecryptCredentialTypeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var AAAJWTValidatorDecryptKeyIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var AAAJWTValidatorDecryptSSecretIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var AAAJWTValidatorDecryptJWKIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var AAAJWTValidatorDecryptFetchCredURLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var AAAJWTValidatorDecryptFetchCredSSLProfileIgnoreVal = validators.Evaluation{
	Evaluation: "logical-not",
	Conditions: []validators.Evaluation{
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "val_method",
					AttrType:    "DmJWTValMethod",
					AttrDefault: "",
					Value:       []string{"decrypt"},
				},
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "decrypt_credential_type",
					AttrType:    "String",
					AttrDefault: "",
					Value:       []string{"jwk-remote"},
				},
			},
		},
	},
}
var AAAJWTValidatorValidateCustomIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var AAAJWTValidatorVerifyCredentialTypeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var AAAJWTValidatorVerifyCertificateIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var AAAJWTValidatorVerifyCertificateAgainstValCredIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var AAAJWTValidatorVerifyValCredIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var AAAJWTValidatorVerifySSecretIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var AAAJWTValidatorVerifyJWKIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var AAAJWTValidatorVerifyFetchCredURLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var AAAJWTValidatorVerifyFetchCredSSLProfileIgnoreVal = validators.Evaluation{
	Evaluation: "logical-not",
	Conditions: []validators.Evaluation{
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "val_method",
					AttrType:    "DmJWTValMethod",
					AttrDefault: "",
					Value:       []string{"verify"},
				},
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "verify_credential_type",
					AttrType:    "String",
					AttrDefault: "",
					Value:       []string{"jwk-remote"},
				},
			},
		},
	},
}

var AAAJWTValidatorObjectType = map[string]attr.Type{
	"id":                                  types.StringType,
	"app_domain":                          types.StringType,
	"user_summary":                        types.StringType,
	"issuer":                              types.StringType,
	"aud":                                 types.StringType,
	"val_method":                          types.ObjectType{AttrTypes: DmJWTValMethodObjectType},
	"customized_script":                   types.StringType,
	"decrypt_credential_type":             types.StringType,
	"decrypt_key":                         types.StringType,
	"decrypt_s_secret":                    types.StringType,
	"decrypt_jwk":                         types.StringType,
	"decrypt_fetch_cred_url":              types.StringType,
	"decrypt_fetch_cred_ssl_profile":      types.StringType,
	"validate_custom":                     types.StringType,
	"verify_credential_type":              types.StringType,
	"verify_certificate":                  types.StringType,
	"verify_certificate_against_val_cred": types.BoolType,
	"verify_val_cred":                     types.StringType,
	"verify_s_secret":                     types.StringType,
	"verify_jwk":                          types.StringType,
	"verify_fetch_cred_url":               types.StringType,
	"verify_fetch_cred_ssl_profile":       types.StringType,
	"claims":                              types.ListType{ElemType: types.ObjectType{AttrTypes: DmClaimObjectType}},
	"username_claim":                      types.StringType,
	"dependency_actions":                  actions.ActionsListType,
}

func (data AAAJWTValidator) GetPath() string {
	rest_path := "/mgmt/config/{domain}/AAAJWTValidator"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data AAAJWTValidator) IsNull() bool {
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
	if !data.Aud.IsNull() {
		return false
	}
	if data.ValMethod != nil {
		if !data.ValMethod.IsNull() {
			return false
		}
	}
	if !data.CustomizedScript.IsNull() {
		return false
	}
	if !data.DecryptCredentialType.IsNull() {
		return false
	}
	if !data.DecryptKey.IsNull() {
		return false
	}
	if !data.DecryptSSecret.IsNull() {
		return false
	}
	if !data.DecryptJwk.IsNull() {
		return false
	}
	if !data.DecryptFetchCredUrl.IsNull() {
		return false
	}
	if !data.DecryptFetchCredSslProfile.IsNull() {
		return false
	}
	if !data.ValidateCustom.IsNull() {
		return false
	}
	if !data.VerifyCredentialType.IsNull() {
		return false
	}
	if !data.VerifyCertificate.IsNull() {
		return false
	}
	if !data.VerifyCertificateAgainstValCred.IsNull() {
		return false
	}
	if !data.VerifyValCred.IsNull() {
		return false
	}
	if !data.VerifySSecret.IsNull() {
		return false
	}
	if !data.VerifyJwk.IsNull() {
		return false
	}
	if !data.VerifyFetchCredUrl.IsNull() {
		return false
	}
	if !data.VerifyFetchCredSslProfile.IsNull() {
		return false
	}
	if !data.Claims.IsNull() {
		return false
	}
	if !data.UsernameClaim.IsNull() {
		return false
	}
	return true
}

func (data AAAJWTValidator) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Aud.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Aud`, data.Aud.ValueString())
	}
	if data.ValMethod != nil {
		if !data.ValMethod.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`ValMethod`, data.ValMethod.ToBody(ctx, ""))
		}
	}
	if !data.CustomizedScript.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CustomizedScript`, data.CustomizedScript.ValueString())
	}
	if !data.DecryptCredentialType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DecryptCredentialType`, data.DecryptCredentialType.ValueString())
	}
	if !data.DecryptKey.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DecryptKey`, data.DecryptKey.ValueString())
	}
	if !data.DecryptSSecret.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DecryptSSecret`, data.DecryptSSecret.ValueString())
	}
	if !data.DecryptJwk.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DecryptJWK`, data.DecryptJwk.ValueString())
	}
	if !data.DecryptFetchCredUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DecryptFetchCredURL`, data.DecryptFetchCredUrl.ValueString())
	}
	if !data.DecryptFetchCredSslProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DecryptFetchCredSSLProfile`, data.DecryptFetchCredSslProfile.ValueString())
	}
	if !data.ValidateCustom.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ValidateCustom`, data.ValidateCustom.ValueString())
	}
	if !data.VerifyCredentialType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`VerifyCredentialType`, data.VerifyCredentialType.ValueString())
	}
	if !data.VerifyCertificate.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`VerifyCertificate`, data.VerifyCertificate.ValueString())
	}
	if !data.VerifyCertificateAgainstValCred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`VerifyCertificateAgainstValCred`, tfutils.StringFromBool(data.VerifyCertificateAgainstValCred, ""))
	}
	if !data.VerifyValCred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`VerifyValCred`, data.VerifyValCred.ValueString())
	}
	if !data.VerifySSecret.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`VerifySSecret`, data.VerifySSecret.ValueString())
	}
	if !data.VerifyJwk.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`VerifyJWK`, data.VerifyJwk.ValueString())
	}
	if !data.VerifyFetchCredUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`VerifyFetchCredURL`, data.VerifyFetchCredUrl.ValueString())
	}
	if !data.VerifyFetchCredSslProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`VerifyFetchCredSSLProfile`, data.VerifyFetchCredSslProfile.ValueString())
	}
	if !data.Claims.IsNull() {
		var dataValues []DmClaim
		data.Claims.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`Claims`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.UsernameClaim.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UsernameClaim`, data.UsernameClaim.ValueString())
	}
	return body
}

func (data *AAAJWTValidator) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
		data.Issuer = types.StringNull()
	}
	if value := res.Get(pathRoot + `Aud`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Aud = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Aud = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValMethod`); value.Exists() {
		data.ValMethod = &DmJWTValMethod{}
		data.ValMethod.FromBody(ctx, "", value)
	} else {
		data.ValMethod = nil
	}
	if value := res.Get(pathRoot + `CustomizedScript`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CustomizedScript = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CustomizedScript = types.StringNull()
	}
	if value := res.Get(pathRoot + `DecryptCredentialType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DecryptCredentialType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DecryptCredentialType = types.StringNull()
	}
	if value := res.Get(pathRoot + `DecryptKey`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DecryptKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DecryptKey = types.StringNull()
	}
	if value := res.Get(pathRoot + `DecryptSSecret`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DecryptSSecret = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DecryptSSecret = types.StringNull()
	}
	if value := res.Get(pathRoot + `DecryptJWK`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DecryptJwk = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DecryptJwk = types.StringNull()
	}
	if value := res.Get(pathRoot + `DecryptFetchCredURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DecryptFetchCredUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DecryptFetchCredUrl = types.StringValue("http://example.com/v3/key")
	}
	if value := res.Get(pathRoot + `DecryptFetchCredSSLProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DecryptFetchCredSslProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DecryptFetchCredSslProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidateCustom`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ValidateCustom = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ValidateCustom = types.StringNull()
	}
	if value := res.Get(pathRoot + `VerifyCredentialType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.VerifyCredentialType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.VerifyCredentialType = types.StringNull()
	}
	if value := res.Get(pathRoot + `VerifyCertificate`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.VerifyCertificate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.VerifyCertificate = types.StringNull()
	}
	if value := res.Get(pathRoot + `VerifyCertificateAgainstValCred`); value.Exists() {
		data.VerifyCertificateAgainstValCred = tfutils.BoolFromString(value.String())
	} else {
		data.VerifyCertificateAgainstValCred = types.BoolNull()
	}
	if value := res.Get(pathRoot + `VerifyValCred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.VerifyValCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.VerifyValCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `VerifySSecret`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.VerifySSecret = tfutils.ParseStringFromGJSON(value)
	} else {
		data.VerifySSecret = types.StringNull()
	}
	if value := res.Get(pathRoot + `VerifyJWK`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.VerifyJwk = tfutils.ParseStringFromGJSON(value)
	} else {
		data.VerifyJwk = types.StringNull()
	}
	if value := res.Get(pathRoot + `VerifyFetchCredURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.VerifyFetchCredUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.VerifyFetchCredUrl = types.StringValue("http://example.com/v3/certs")
	}
	if value := res.Get(pathRoot + `VerifyFetchCredSSLProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.VerifyFetchCredSslProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.VerifyFetchCredSslProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `Claims`); value.Exists() {
		l := []DmClaim{}
		if value := res.Get(`Claims`); value.Exists() {
			for _, v := range value.Array() {
				item := DmClaim{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Claims, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmClaimObjectType}, l)
		} else {
			data.Claims = types.ListNull(types.ObjectType{AttrTypes: DmClaimObjectType})
		}
	} else {
		data.Claims = types.ListNull(types.ObjectType{AttrTypes: DmClaimObjectType})
	}
	if value := res.Get(pathRoot + `UsernameClaim`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UsernameClaim = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UsernameClaim = types.StringValue("sub")
	}
}

func (data *AAAJWTValidator) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	} else {
		data.Issuer = types.StringNull()
	}
	if value := res.Get(pathRoot + `Aud`); value.Exists() && !data.Aud.IsNull() {
		data.Aud = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Aud = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValMethod`); value.Exists() {
		data.ValMethod.UpdateFromBody(ctx, "", value)
	} else {
		data.ValMethod = nil
	}
	if value := res.Get(pathRoot + `CustomizedScript`); value.Exists() && !data.CustomizedScript.IsNull() {
		data.CustomizedScript = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CustomizedScript = types.StringNull()
	}
	if value := res.Get(pathRoot + `DecryptCredentialType`); value.Exists() && !data.DecryptCredentialType.IsNull() {
		data.DecryptCredentialType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DecryptCredentialType = types.StringNull()
	}
	if value := res.Get(pathRoot + `DecryptKey`); value.Exists() && !data.DecryptKey.IsNull() {
		data.DecryptKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DecryptKey = types.StringNull()
	}
	if value := res.Get(pathRoot + `DecryptSSecret`); value.Exists() && !data.DecryptSSecret.IsNull() {
		data.DecryptSSecret = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DecryptSSecret = types.StringNull()
	}
	if value := res.Get(pathRoot + `DecryptJWK`); value.Exists() && !data.DecryptJwk.IsNull() {
		data.DecryptJwk = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DecryptJwk = types.StringNull()
	}
	if value := res.Get(pathRoot + `DecryptFetchCredURL`); value.Exists() && !data.DecryptFetchCredUrl.IsNull() {
		data.DecryptFetchCredUrl = tfutils.ParseStringFromGJSON(value)
	} else if data.DecryptFetchCredUrl.ValueString() != "http://example.com/v3/key" {
		data.DecryptFetchCredUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `DecryptFetchCredSSLProfile`); value.Exists() && !data.DecryptFetchCredSslProfile.IsNull() {
		data.DecryptFetchCredSslProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DecryptFetchCredSslProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidateCustom`); value.Exists() && !data.ValidateCustom.IsNull() {
		data.ValidateCustom = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ValidateCustom = types.StringNull()
	}
	if value := res.Get(pathRoot + `VerifyCredentialType`); value.Exists() && !data.VerifyCredentialType.IsNull() {
		data.VerifyCredentialType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.VerifyCredentialType = types.StringNull()
	}
	if value := res.Get(pathRoot + `VerifyCertificate`); value.Exists() && !data.VerifyCertificate.IsNull() {
		data.VerifyCertificate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.VerifyCertificate = types.StringNull()
	}
	if value := res.Get(pathRoot + `VerifyCertificateAgainstValCred`); value.Exists() && !data.VerifyCertificateAgainstValCred.IsNull() {
		data.VerifyCertificateAgainstValCred = tfutils.BoolFromString(value.String())
	} else if data.VerifyCertificateAgainstValCred.ValueBool() {
		data.VerifyCertificateAgainstValCred = types.BoolNull()
	}
	if value := res.Get(pathRoot + `VerifyValCred`); value.Exists() && !data.VerifyValCred.IsNull() {
		data.VerifyValCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.VerifyValCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `VerifySSecret`); value.Exists() && !data.VerifySSecret.IsNull() {
		data.VerifySSecret = tfutils.ParseStringFromGJSON(value)
	} else {
		data.VerifySSecret = types.StringNull()
	}
	if value := res.Get(pathRoot + `VerifyJWK`); value.Exists() && !data.VerifyJwk.IsNull() {
		data.VerifyJwk = tfutils.ParseStringFromGJSON(value)
	} else {
		data.VerifyJwk = types.StringNull()
	}
	if value := res.Get(pathRoot + `VerifyFetchCredURL`); value.Exists() && !data.VerifyFetchCredUrl.IsNull() {
		data.VerifyFetchCredUrl = tfutils.ParseStringFromGJSON(value)
	} else if data.VerifyFetchCredUrl.ValueString() != "http://example.com/v3/certs" {
		data.VerifyFetchCredUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `VerifyFetchCredSSLProfile`); value.Exists() && !data.VerifyFetchCredSslProfile.IsNull() {
		data.VerifyFetchCredSslProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.VerifyFetchCredSslProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `Claims`); value.Exists() && !data.Claims.IsNull() {
		l := []DmClaim{}
		e := []DmClaim{}
		data.Claims.ElementsAs(ctx, &e, false)
		if len(value.Array()) == len(e) {
			for i, v := range value.Array() {
				item := e[i]
				item.UpdateFromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		} else {
			for _, v := range value.Array() {
				item := DmClaim{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Claims, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmClaimObjectType}, l)
		} else {
			data.Claims = types.ListNull(types.ObjectType{AttrTypes: DmClaimObjectType})
		}
	} else {
		data.Claims = types.ListNull(types.ObjectType{AttrTypes: DmClaimObjectType})
	}
	if value := res.Get(pathRoot + `UsernameClaim`); value.Exists() && !data.UsernameClaim.IsNull() {
		data.UsernameClaim = tfutils.ParseStringFromGJSON(value)
	} else if data.UsernameClaim.ValueString() != "sub" {
		data.UsernameClaim = types.StringNull()
	}
}
