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

type XACMLPDP struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	EqualPolicies     types.Bool                  `tfsdk:"equal_policies"`
	GeneralPolicy     types.String                `tfsdk:"general_policy"`
	CombiningAlg      types.String                `tfsdk:"combining_alg"`
	DependentPolicy   types.List                  `tfsdk:"dependent_policy"`
	Directory         types.List                  `tfsdk:"directory"`
	CacheTtl          types.Int64                 `tfsdk:"cache_ttl"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var XACMLPDPGeneralPolicyCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "equal_policies",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var XACMLPDPGeneralPolicyIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var XACMLPDPCombiningAlgCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "equal_policies",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var XACMLPDPCombiningAlgIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var XACMLPDPObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"user_summary":       types.StringType,
	"equal_policies":     types.BoolType,
	"general_policy":     types.StringType,
	"combining_alg":      types.StringType,
	"dependent_policy":   types.ListType{ElemType: types.StringType},
	"directory":          types.ListType{ElemType: types.StringType},
	"cache_ttl":          types.Int64Type,
	"dependency_actions": actions.ActionsListType,
}

func (data XACMLPDP) GetPath() string {
	rest_path := "/mgmt/config/{domain}/XACMLPDP"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data XACMLPDP) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.EqualPolicies.IsNull() {
		return false
	}
	if !data.GeneralPolicy.IsNull() {
		return false
	}
	if !data.CombiningAlg.IsNull() {
		return false
	}
	if !data.DependentPolicy.IsNull() {
		return false
	}
	if !data.Directory.IsNull() {
		return false
	}
	if !data.CacheTtl.IsNull() {
		return false
	}
	return true
}

func (data XACMLPDP) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.EqualPolicies.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EqualPolicies`, tfutils.StringFromBool(data.EqualPolicies, ""))
	}
	if !data.GeneralPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GeneralPolicy`, data.GeneralPolicy.ValueString())
	}
	if !data.CombiningAlg.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CombiningAlg`, data.CombiningAlg.ValueString())
	}
	if !data.DependentPolicy.IsNull() {
		var dataValues []string
		data.DependentPolicy.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.Set(body, pathRoot+`DependentPolicy`+".-1", map[string]string{"value": val})
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`DependentPolicy`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`DependentPolicy`, "[]")
	}
	if !data.Directory.IsNull() {
		var dataValues []string
		data.Directory.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.Set(body, pathRoot+`Directory`+".-1", map[string]string{"value": val})
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`Directory`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`Directory`, "[]")
	}
	if !data.CacheTtl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CacheTTL`, data.CacheTtl.ValueInt64())
	}
	return body
}

func (data *XACMLPDP) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `EqualPolicies`); value.Exists() {
		data.EqualPolicies = tfutils.BoolFromString(value.String())
	} else {
		data.EqualPolicies = types.BoolNull()
	}
	if value := res.Get(pathRoot + `GeneralPolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.GeneralPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GeneralPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `CombiningAlg`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CombiningAlg = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CombiningAlg = types.StringNull()
	}
	if value := res.Get(pathRoot + `DependentPolicy`); value.Exists() {
		data.DependentPolicy = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.DependentPolicy = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `Directory`); value.Exists() {
		data.Directory = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Directory = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `CacheTTL`); value.Exists() {
		data.CacheTtl = types.Int64Value(value.Int())
	} else {
		data.CacheTtl = types.Int64Null()
	}
}

func (data *XACMLPDP) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `EqualPolicies`); value.Exists() && !data.EqualPolicies.IsNull() {
		data.EqualPolicies = tfutils.BoolFromString(value.String())
	} else if data.EqualPolicies.ValueBool() {
		data.EqualPolicies = types.BoolNull()
	}
	if value := res.Get(pathRoot + `GeneralPolicy`); value.Exists() && !data.GeneralPolicy.IsNull() {
		data.GeneralPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GeneralPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `CombiningAlg`); value.Exists() && !data.CombiningAlg.IsNull() {
		data.CombiningAlg = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CombiningAlg = types.StringNull()
	}
	if value := res.Get(pathRoot + `DependentPolicy`); value.Exists() && !data.DependentPolicy.IsNull() {
		data.DependentPolicy = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.DependentPolicy = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `Directory`); value.Exists() && !data.Directory.IsNull() {
		data.Directory = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Directory = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `CacheTTL`); value.Exists() && !data.CacheTtl.IsNull() {
		data.CacheTtl = types.Int64Value(value.Int())
	} else {
		data.CacheTtl = types.Int64Null()
	}
}
