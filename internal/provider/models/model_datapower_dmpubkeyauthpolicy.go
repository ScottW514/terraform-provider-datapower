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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	DataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	ResourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmPubkeyAuthPolicy struct {
	RegExp    types.String `tfsdk:"reg_exp"`
	CryptoKey types.String `tfsdk:"crypto_key"`
}

var DmPubkeyAuthPolicyObjectType = map[string]attr.Type{
	"reg_exp":    types.StringType,
	"crypto_key": types.StringType,
}
var DmPubkeyAuthPolicyObjectDefault = map[string]attr.Value{
	"reg_exp":    types.StringNull(),
	"crypto_key": types.StringNull(),
}
var DmPubkeyAuthPolicyDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"reg_exp": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL matching expression", "", "").String,
			Computed:            true,
		},
		"crypto_key": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Private Key", "", "cryptokey").String,
			Computed:            true,
		},
	},
}
var DmPubkeyAuthPolicyResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"reg_exp": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL matching expression", "", "").String,
			Required:            true,
		},
		"crypto_key": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Private Key", "", "cryptokey").String,
			Required:            true,
		},
	},
}

func (data DmPubkeyAuthPolicy) IsNull() bool {
	if !data.RegExp.IsNull() {
		return false
	}
	if !data.CryptoKey.IsNull() {
		return false
	}
	return true
}

func (data DmPubkeyAuthPolicy) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.RegExp.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RegExp`, data.RegExp.ValueString())
	}
	if !data.CryptoKey.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CryptoKey`, data.CryptoKey.ValueString())
	}
	return body
}

func (data *DmPubkeyAuthPolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `RegExp`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RegExp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RegExp = types.StringNull()
	}
	if value := res.Get(pathRoot + `CryptoKey`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CryptoKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CryptoKey = types.StringNull()
	}
}

func (data *DmPubkeyAuthPolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `RegExp`); value.Exists() && !data.RegExp.IsNull() {
		data.RegExp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RegExp = types.StringNull()
	}
	if value := res.Get(pathRoot + `CryptoKey`); value.Exists() && !data.CryptoKey.IsNull() {
		data.CryptoKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CryptoKey = types.StringNull()
	}
}
