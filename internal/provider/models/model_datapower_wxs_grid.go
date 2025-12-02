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

type WXSGrid struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	Collective        types.String                `tfsdk:"collective"`
	Grid              types.String                `tfsdk:"grid"`
	UserName          types.String                `tfsdk:"user_name"`
	PasswordAlias     types.String                `tfsdk:"password_alias"`
	Timeout           types.Int64                 `tfsdk:"timeout"`
	SslClient         types.String                `tfsdk:"ssl_client"`
	Encrypt           types.Bool                  `tfsdk:"encrypt"`
	EncryptSskey      types.String                `tfsdk:"encrypt_sskey"`
	EncryptAlg        types.String                `tfsdk:"encrypt_alg"`
	KeyObfuscation    types.Bool                  `tfsdk:"key_obfuscation"`
	KeyObfuscationAlg types.String                `tfsdk:"key_obfuscation_alg"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
	ProviderTarget    types.String                `tfsdk:"provider_target"`
}

var WXSGridEncryptSSKeyCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "encrypt",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var WXSGridEncryptSSKeyIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var WXSGridEncryptAlgCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "encrypt",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var WXSGridEncryptAlgIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var WXSGridKeyObfuscationAlgCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "key_obfuscation",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var WXSGridKeyObfuscationAlgIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var WXSGridObjectType = map[string]attr.Type{
	"provider_target":     types.StringType,
	"id":                  types.StringType,
	"app_domain":          types.StringType,
	"user_summary":        types.StringType,
	"collective":          types.StringType,
	"grid":                types.StringType,
	"user_name":           types.StringType,
	"password_alias":      types.StringType,
	"timeout":             types.Int64Type,
	"ssl_client":          types.StringType,
	"encrypt":             types.BoolType,
	"encrypt_sskey":       types.StringType,
	"encrypt_alg":         types.StringType,
	"key_obfuscation":     types.BoolType,
	"key_obfuscation_alg": types.StringType,
	"dependency_actions":  actions.ActionsListType,
}

func (data WXSGrid) GetPath() string {
	rest_path := "/mgmt/config/{domain}/WXSGrid"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data WXSGrid) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Collective.IsNull() {
		return false
	}
	if !data.Grid.IsNull() {
		return false
	}
	if !data.UserName.IsNull() {
		return false
	}
	if !data.PasswordAlias.IsNull() {
		return false
	}
	if !data.Timeout.IsNull() {
		return false
	}
	if !data.SslClient.IsNull() {
		return false
	}
	if !data.Encrypt.IsNull() {
		return false
	}
	if !data.EncryptSskey.IsNull() {
		return false
	}
	if !data.EncryptAlg.IsNull() {
		return false
	}
	if !data.KeyObfuscation.IsNull() {
		return false
	}
	if !data.KeyObfuscationAlg.IsNull() {
		return false
	}
	return true
}

func (data WXSGrid) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Collective.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Collective`, data.Collective.ValueString())
	}
	if !data.Grid.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Grid`, data.Grid.ValueString())
	}
	if !data.UserName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserName`, data.UserName.ValueString())
	}
	if !data.PasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PasswordAlias`, data.PasswordAlias.ValueString())
	}
	if !data.Timeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Timeout`, data.Timeout.ValueInt64())
	}
	if !data.SslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClient`, data.SslClient.ValueString())
	}
	if !data.Encrypt.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Encrypt`, tfutils.StringFromBool(data.Encrypt, ""))
	}
	if !data.EncryptSskey.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EncryptSSKey`, data.EncryptSskey.ValueString())
	}
	if !data.EncryptAlg.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EncryptAlg`, data.EncryptAlg.ValueString())
	}
	if !data.KeyObfuscation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`KeyObfuscation`, tfutils.StringFromBool(data.KeyObfuscation, ""))
	}
	if !data.KeyObfuscationAlg.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`KeyObfuscationAlg`, data.KeyObfuscationAlg.ValueString())
	}
	return body
}

func (data *WXSGrid) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Collective`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Collective = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Collective = types.StringNull()
	}
	if value := res.Get(pathRoot + `Grid`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Grid = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Grid = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserName = types.StringNull()
	}
	if value := res.Get(pathRoot + `PasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `Timeout`); value.Exists() {
		data.Timeout = types.Int64Value(value.Int())
	} else {
		data.Timeout = types.Int64Value(1000)
	}
	if value := res.Get(pathRoot + `SSLClient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `Encrypt`); value.Exists() {
		data.Encrypt = tfutils.BoolFromString(value.String())
	} else {
		data.Encrypt = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EncryptSSKey`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EncryptSskey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EncryptSskey = types.StringNull()
	}
	if value := res.Get(pathRoot + `EncryptAlg`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EncryptAlg = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EncryptAlg = types.StringValue("tripledes-cbc")
	}
	if value := res.Get(pathRoot + `KeyObfuscation`); value.Exists() {
		data.KeyObfuscation = tfutils.BoolFromString(value.String())
	} else {
		data.KeyObfuscation = types.BoolNull()
	}
	if value := res.Get(pathRoot + `KeyObfuscationAlg`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.KeyObfuscationAlg = tfutils.ParseStringFromGJSON(value)
	} else {
		data.KeyObfuscationAlg = types.StringValue("sha256")
	}
}

func (data *WXSGrid) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Collective`); value.Exists() && !data.Collective.IsNull() {
		data.Collective = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Collective = types.StringNull()
	}
	if value := res.Get(pathRoot + `Grid`); value.Exists() && !data.Grid.IsNull() {
		data.Grid = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Grid = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserName`); value.Exists() && !data.UserName.IsNull() {
		data.UserName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserName = types.StringNull()
	}
	if value := res.Get(pathRoot + `PasswordAlias`); value.Exists() && !data.PasswordAlias.IsNull() {
		data.PasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `Timeout`); value.Exists() && !data.Timeout.IsNull() {
		data.Timeout = types.Int64Value(value.Int())
	} else if data.Timeout.ValueInt64() != 1000 {
		data.Timeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `SSLClient`); value.Exists() && !data.SslClient.IsNull() {
		data.SslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `Encrypt`); value.Exists() && !data.Encrypt.IsNull() {
		data.Encrypt = tfutils.BoolFromString(value.String())
	} else if data.Encrypt.ValueBool() {
		data.Encrypt = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EncryptSSKey`); value.Exists() && !data.EncryptSskey.IsNull() {
		data.EncryptSskey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EncryptSskey = types.StringNull()
	}
	if value := res.Get(pathRoot + `EncryptAlg`); value.Exists() && !data.EncryptAlg.IsNull() {
		data.EncryptAlg = tfutils.ParseStringFromGJSON(value)
	} else if data.EncryptAlg.ValueString() != "tripledes-cbc" {
		data.EncryptAlg = types.StringNull()
	}
	if value := res.Get(pathRoot + `KeyObfuscation`); value.Exists() && !data.KeyObfuscation.IsNull() {
		data.KeyObfuscation = tfutils.BoolFromString(value.String())
	} else if data.KeyObfuscation.ValueBool() {
		data.KeyObfuscation = types.BoolNull()
	}
	if value := res.Get(pathRoot + `KeyObfuscationAlg`); value.Exists() && !data.KeyObfuscationAlg.IsNull() {
		data.KeyObfuscationAlg = tfutils.ParseStringFromGJSON(value)
	} else if data.KeyObfuscationAlg.ValueString() != "sha256" {
		data.KeyObfuscationAlg = types.StringNull()
	}
}
