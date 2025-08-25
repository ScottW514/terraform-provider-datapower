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

type DmJWTGenMethod struct {
	Sign    types.Bool `tfsdk:"sign"`
	Encrypt types.Bool `tfsdk:"encrypt"`
}

var DmJWTGenMethodObjectType = map[string]attr.Type{
	"sign":    types.BoolType,
	"encrypt": types.BoolType,
}
var DmJWTGenMethodObjectDefault = map[string]attr.Value{
	"sign":    types.BoolValue(false),
	"encrypt": types.BoolValue(false),
}

func GetDmJWTGenMethodDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmJWTGenMethodDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"sign": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Sign the JWT", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"encrypt": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Encrypt the JWT", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
		},
	}
	DmJWTGenMethodDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmJWTGenMethodDataSourceSchema
}
func GetDmJWTGenMethodResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmJWTGenMethodResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmJWTGenMethodObjectType,
				DmJWTGenMethodObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"sign": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Sign the JWT", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"encrypt": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Encrypt the JWT", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
	DmJWTGenMethodResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmJWTGenMethodResourceSchema.Required = true
	} else {
		DmJWTGenMethodResourceSchema.Optional = true
		DmJWTGenMethodResourceSchema.Computed = true
	}
	return DmJWTGenMethodResourceSchema
}

func (data DmJWTGenMethod) IsNull() bool {
	if !data.Sign.IsNull() {
		return false
	}
	if !data.Encrypt.IsNull() {
		return false
	}
	return true
}

func (data DmJWTGenMethod) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Sign.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`sign`, tfutils.StringFromBool(data.Sign, ""))
	}
	if !data.Encrypt.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`encrypt`, tfutils.StringFromBool(data.Encrypt, ""))
	}
	return body
}

func (data *DmJWTGenMethod) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `sign`); value.Exists() {
		data.Sign = tfutils.BoolFromString(value.String())
	} else {
		data.Sign = types.BoolNull()
	}
	if value := res.Get(pathRoot + `encrypt`); value.Exists() {
		data.Encrypt = tfutils.BoolFromString(value.String())
	} else {
		data.Encrypt = types.BoolNull()
	}
}

func (data *DmJWTGenMethod) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `sign`); value.Exists() && !data.Sign.IsNull() {
		data.Sign = tfutils.BoolFromString(value.String())
	} else if data.Sign.ValueBool() {
		data.Sign = types.BoolNull()
	}
	if value := res.Get(pathRoot + `encrypt`); value.Exists() && !data.Encrypt.IsNull() {
		data.Encrypt = tfutils.BoolFromString(value.String())
	} else if data.Encrypt.ValueBool() {
		data.Encrypt = types.BoolNull()
	}
}
