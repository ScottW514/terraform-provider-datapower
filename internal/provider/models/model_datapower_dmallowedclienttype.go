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

type DmAllowedClientType struct {
	Confidential types.Bool `tfsdk:"confidential"`
	Public       types.Bool `tfsdk:"public"`
}

var DmAllowedClientTypeObjectType = map[string]attr.Type{
	"confidential": types.BoolType,
	"public":       types.BoolType,
}
var DmAllowedClientTypeObjectDefault = map[string]attr.Value{
	"confidential": types.BoolValue(true),
	"public":       types.BoolValue(false),
}
var DmAllowedClientTypeDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"confidential": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Confidential", "", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"public": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Public", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
	},
}
var DmAllowedClientTypeResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmAllowedClientTypeObjectType,
			DmAllowedClientTypeObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"confidential": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Confidential", "", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"public": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Public", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
	},
}

func (data DmAllowedClientType) IsNull() bool {
	if !data.Confidential.IsNull() {
		return false
	}
	if !data.Public.IsNull() {
		return false
	}
	return true
}
func GetDmAllowedClientTypeDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmAllowedClientTypeDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmAllowedClientTypeDataSourceSchema
}

func GetDmAllowedClientTypeResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmAllowedClientTypeResourceSchema.Required = true
	} else {
		DmAllowedClientTypeResourceSchema.Optional = true
		DmAllowedClientTypeResourceSchema.Computed = true
	}
	DmAllowedClientTypeResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmAllowedClientTypeResourceSchema
}

func (data DmAllowedClientType) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Confidential.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`confidential`, tfutils.StringFromBool(data.Confidential, ""))
	}
	if !data.Public.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`public`, tfutils.StringFromBool(data.Public, ""))
	}
	return body
}

func (data *DmAllowedClientType) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `confidential`); value.Exists() {
		data.Confidential = tfutils.BoolFromString(value.String())
	} else {
		data.Confidential = types.BoolNull()
	}
	if value := res.Get(pathRoot + `public`); value.Exists() {
		data.Public = tfutils.BoolFromString(value.String())
	} else {
		data.Public = types.BoolNull()
	}
}

func (data *DmAllowedClientType) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `confidential`); value.Exists() && !data.Confidential.IsNull() {
		data.Confidential = tfutils.BoolFromString(value.String())
	} else if !data.Confidential.ValueBool() {
		data.Confidential = types.BoolNull()
	}
	if value := res.Get(pathRoot + `public`); value.Exists() && !data.Public.IsNull() {
		data.Public = tfutils.BoolFromString(value.String())
	} else if data.Public.ValueBool() {
		data.Public = types.BoolNull()
	}
}
