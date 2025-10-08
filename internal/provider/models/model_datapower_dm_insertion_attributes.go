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

type DmInsertionAttributes struct {
	Secure   types.Bool `tfsdk:"secure"`
	Httponly types.Bool `tfsdk:"httponly"`
}

var DmInsertionAttributesObjectType = map[string]attr.Type{
	"secure":   types.BoolType,
	"httponly": types.BoolType,
}
var DmInsertionAttributesObjectDefault = map[string]attr.Value{
	"secure":   types.BoolValue(false),
	"httponly": types.BoolValue(false),
}

func GetDmInsertionAttributesDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmInsertionAttributesDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"secure": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Secure",
				Computed:            true,
			},
			"httponly": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "HttpOnly",
				Computed:            true,
			},
		},
	}
	DmInsertionAttributesDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmInsertionAttributesDataSourceSchema
}
func GetDmInsertionAttributesResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmInsertionAttributesResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmInsertionAttributesObjectType,
				DmInsertionAttributesObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"secure": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Secure", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"httponly": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HttpOnly", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
	DmInsertionAttributesResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmInsertionAttributesResourceSchema.Required = true
	} else {
		DmInsertionAttributesResourceSchema.Optional = true
		DmInsertionAttributesResourceSchema.Computed = true
	}
	return DmInsertionAttributesResourceSchema
}

func (data DmInsertionAttributes) IsNull() bool {
	if !data.Secure.IsNull() {
		return false
	}
	if !data.Httponly.IsNull() {
		return false
	}
	return true
}

func (data DmInsertionAttributes) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Secure.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`secure`, tfutils.StringFromBool(data.Secure, ""))
	}
	if !data.Httponly.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`httponly`, tfutils.StringFromBool(data.Httponly, ""))
	}
	return body
}

func (data *DmInsertionAttributes) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `secure`); value.Exists() {
		data.Secure = tfutils.BoolFromString(value.String())
	} else {
		data.Secure = types.BoolNull()
	}
	if value := res.Get(pathRoot + `httponly`); value.Exists() {
		data.Httponly = tfutils.BoolFromString(value.String())
	} else {
		data.Httponly = types.BoolNull()
	}
}

func (data *DmInsertionAttributes) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `secure`); value.Exists() && !data.Secure.IsNull() {
		data.Secure = tfutils.BoolFromString(value.String())
	} else if data.Secure.ValueBool() {
		data.Secure = types.BoolNull()
	}
	if value := res.Get(pathRoot + `httponly`); value.Exists() && !data.Httponly.IsNull() {
		data.Httponly = tfutils.BoolFromString(value.String())
	} else if data.Httponly.ValueBool() {
		data.Httponly = types.BoolNull()
	}
}
