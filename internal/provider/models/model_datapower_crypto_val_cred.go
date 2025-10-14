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

type CryptoValCred struct {
	Id                 types.String                `tfsdk:"id"`
	AppDomain          types.String                `tfsdk:"app_domain"`
	Certificate        types.List                  `tfsdk:"certificate"`
	CertValidationMode types.String                `tfsdk:"cert_validation_mode"`
	UseCrl             types.Bool                  `tfsdk:"use_crl"`
	RequireCrl         types.Bool                  `tfsdk:"require_crl"`
	CrlDpHandling      types.String                `tfsdk:"crl_dp_handling"`
	InitialPolicySet   types.List                  `tfsdk:"initial_policy_set"`
	ExplicitPolicy     types.Bool                  `tfsdk:"explicit_policy"`
	CheckDates         types.Bool                  `tfsdk:"check_dates"`
	DependencyActions  []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var CryptoValCredRequireCRLIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "use_crl",
	AttrType:    "Bool",
	AttrDefault: "true",
	Value:       []string{"false"},
}

var CryptoValCredCRLDPHandlingIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "use_crl",
	AttrType:    "Bool",
	AttrDefault: "true",
	Value:       []string{"false"},
}

var CryptoValCredInitialPolicySetIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "cert_validation_mode",
	AttrType:    "String",
	AttrDefault: "legacy",
	Value:       []string{"pkix"},
}

var CryptoValCredExplicitPolicyIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "cert_validation_mode",
	AttrType:    "String",
	AttrDefault: "legacy",
	Value:       []string{"pkix"},
}

var CryptoValCredObjectType = map[string]attr.Type{
	"id":                   types.StringType,
	"app_domain":           types.StringType,
	"certificate":          types.ListType{ElemType: types.StringType},
	"cert_validation_mode": types.StringType,
	"use_crl":              types.BoolType,
	"require_crl":          types.BoolType,
	"crl_dp_handling":      types.StringType,
	"initial_policy_set":   types.ListType{ElemType: types.StringType},
	"explicit_policy":      types.BoolType,
	"check_dates":          types.BoolType,
	"dependency_actions":   actions.ActionsListType,
}

func (data CryptoValCred) GetPath() string {
	rest_path := "/mgmt/config/{domain}/CryptoValCred"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data CryptoValCred) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Certificate.IsNull() {
		return false
	}
	if !data.CertValidationMode.IsNull() {
		return false
	}
	if !data.UseCrl.IsNull() {
		return false
	}
	if !data.RequireCrl.IsNull() {
		return false
	}
	if !data.CrlDpHandling.IsNull() {
		return false
	}
	if !data.InitialPolicySet.IsNull() {
		return false
	}
	if !data.ExplicitPolicy.IsNull() {
		return false
	}
	if !data.CheckDates.IsNull() {
		return false
	}
	return true
}

func (data CryptoValCred) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.Certificate.IsNull() {
		var dataValues []string
		data.Certificate.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.Set(body, pathRoot+`Certificate`+".-1", map[string]string{"value": val})
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`Certificate`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`Certificate`, "[]")
	}
	if !data.CertValidationMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CertValidationMode`, data.CertValidationMode.ValueString())
	}
	if !data.UseCrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseCRL`, tfutils.StringFromBool(data.UseCrl, ""))
	}
	if !data.RequireCrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RequireCRL`, tfutils.StringFromBool(data.RequireCrl, ""))
	}
	if !data.CrlDpHandling.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CRLDPHandling`, data.CrlDpHandling.ValueString())
	}
	if !data.InitialPolicySet.IsNull() {
		var dataValues []string
		data.InitialPolicySet.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.Set(body, pathRoot+`InitialPolicySet`+".-1", map[string]string{"value": val})
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`InitialPolicySet`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`InitialPolicySet`, "[]")
	}
	if !data.ExplicitPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ExplicitPolicy`, tfutils.StringFromBool(data.ExplicitPolicy, ""))
	}
	if !data.CheckDates.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CheckDates`, tfutils.StringFromBool(data.CheckDates, ""))
	}
	return body
}

func (data *CryptoValCred) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Certificate`); value.Exists() {
		data.Certificate = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Certificate = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `CertValidationMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CertValidationMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CertValidationMode = types.StringValue("legacy")
	}
	if value := res.Get(pathRoot + `UseCRL`); value.Exists() {
		data.UseCrl = tfutils.BoolFromString(value.String())
	} else {
		data.UseCrl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RequireCRL`); value.Exists() {
		data.RequireCrl = tfutils.BoolFromString(value.String())
	} else {
		data.RequireCrl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CRLDPHandling`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CrlDpHandling = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CrlDpHandling = types.StringValue("ignore")
	}
	if value := res.Get(pathRoot + `InitialPolicySet`); value.Exists() {
		data.InitialPolicySet = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.InitialPolicySet = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `ExplicitPolicy`); value.Exists() {
		data.ExplicitPolicy = tfutils.BoolFromString(value.String())
	} else {
		data.ExplicitPolicy = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CheckDates`); value.Exists() {
		data.CheckDates = tfutils.BoolFromString(value.String())
	} else {
		data.CheckDates = types.BoolNull()
	}
}

func (data *CryptoValCred) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Certificate`); value.Exists() && !data.Certificate.IsNull() {
		data.Certificate = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Certificate = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `CertValidationMode`); value.Exists() && !data.CertValidationMode.IsNull() {
		data.CertValidationMode = tfutils.ParseStringFromGJSON(value)
	} else if data.CertValidationMode.ValueString() != "legacy" {
		data.CertValidationMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `UseCRL`); value.Exists() && !data.UseCrl.IsNull() {
		data.UseCrl = tfutils.BoolFromString(value.String())
	} else if !data.UseCrl.ValueBool() {
		data.UseCrl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RequireCRL`); value.Exists() && !data.RequireCrl.IsNull() {
		data.RequireCrl = tfutils.BoolFromString(value.String())
	} else if data.RequireCrl.ValueBool() {
		data.RequireCrl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CRLDPHandling`); value.Exists() && !data.CrlDpHandling.IsNull() {
		data.CrlDpHandling = tfutils.ParseStringFromGJSON(value)
	} else if data.CrlDpHandling.ValueString() != "ignore" {
		data.CrlDpHandling = types.StringNull()
	}
	if value := res.Get(pathRoot + `InitialPolicySet`); value.Exists() && !data.InitialPolicySet.IsNull() {
		data.InitialPolicySet = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.InitialPolicySet = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `ExplicitPolicy`); value.Exists() && !data.ExplicitPolicy.IsNull() {
		data.ExplicitPolicy = tfutils.BoolFromString(value.String())
	} else if data.ExplicitPolicy.ValueBool() {
		data.ExplicitPolicy = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CheckDates`); value.Exists() && !data.CheckDates.IsNull() {
		data.CheckDates = tfutils.BoolFromString(value.String())
	} else if !data.CheckDates.ValueBool() {
		data.CheckDates = types.BoolNull()
	}
}
