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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmOIDCHybridResponseType struct {
	CodeIdToken            types.Bool `tfsdk:"code_id_token"`
	CodeAccessToken        types.Bool `tfsdk:"code_access_token"`
	CodeIdTokenAccessToken types.Bool `tfsdk:"code_id_token_access_token"`
}

var DmOIDCHybridResponseTypeObjectType = map[string]attr.Type{
	"code_id_token":              types.BoolType,
	"code_access_token":          types.BoolType,
	"code_id_token_access_token": types.BoolType,
}
var DmOIDCHybridResponseTypeObjectDefault = map[string]attr.Value{
	"code_id_token":              types.BoolValue(true),
	"code_access_token":          types.BoolValue(false),
	"code_id_token_access_token": types.BoolValue(false),
}

func GetDmOIDCHybridResponseTypeDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmOIDCHybridResponseTypeDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"code_id_token": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("code id_token", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"code_access_token": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("code token", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"code_id_token_access_token": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("code id_token token", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
		},
	}
	DmOIDCHybridResponseTypeDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmOIDCHybridResponseTypeDataSourceSchema
}
func GetDmOIDCHybridResponseTypeResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmOIDCHybridResponseTypeResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmOIDCHybridResponseTypeObjectType,
				DmOIDCHybridResponseTypeObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"code_id_token": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("code id_token", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"code_access_token": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("code token", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"code_id_token_access_token": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("code id_token token", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
	DmOIDCHybridResponseTypeResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmOIDCHybridResponseTypeResourceSchema.Required = true
	} else {
		DmOIDCHybridResponseTypeResourceSchema.Optional = true
		DmOIDCHybridResponseTypeResourceSchema.Computed = true
	}
	return DmOIDCHybridResponseTypeResourceSchema
}

func (data DmOIDCHybridResponseType) IsNull() bool {
	if !data.CodeIdToken.IsNull() {
		return false
	}
	if !data.CodeAccessToken.IsNull() {
		return false
	}
	if !data.CodeIdTokenAccessToken.IsNull() {
		return false
	}
	return true
}

func (data DmOIDCHybridResponseType) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.CodeIdToken.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`code_id_token`, tfutils.StringFromBool(data.CodeIdToken, ""))
	}
	if !data.CodeAccessToken.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`code_access_token`, tfutils.StringFromBool(data.CodeAccessToken, ""))
	}
	if !data.CodeIdTokenAccessToken.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`code_id_token_access_token`, tfutils.StringFromBool(data.CodeIdTokenAccessToken, ""))
	}
	return body
}

func (data *DmOIDCHybridResponseType) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `code_id_token`); value.Exists() {
		data.CodeIdToken = tfutils.BoolFromString(value.String())
	} else {
		data.CodeIdToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `code_access_token`); value.Exists() {
		data.CodeAccessToken = tfutils.BoolFromString(value.String())
	} else {
		data.CodeAccessToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `code_id_token_access_token`); value.Exists() {
		data.CodeIdTokenAccessToken = tfutils.BoolFromString(value.String())
	} else {
		data.CodeIdTokenAccessToken = types.BoolNull()
	}
}

func (data *DmOIDCHybridResponseType) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `code_id_token`); value.Exists() && !data.CodeIdToken.IsNull() {
		data.CodeIdToken = tfutils.BoolFromString(value.String())
	} else if !data.CodeIdToken.ValueBool() {
		data.CodeIdToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `code_access_token`); value.Exists() && !data.CodeAccessToken.IsNull() {
		data.CodeAccessToken = tfutils.BoolFromString(value.String())
	} else if data.CodeAccessToken.ValueBool() {
		data.CodeAccessToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `code_id_token_access_token`); value.Exists() && !data.CodeIdTokenAccessToken.IsNull() {
		data.CodeIdTokenAccessToken = tfutils.BoolFromString(value.String())
	} else if data.CodeIdTokenAccessToken.ValueBool() {
		data.CodeIdTokenAccessToken = types.BoolNull()
	}
}
