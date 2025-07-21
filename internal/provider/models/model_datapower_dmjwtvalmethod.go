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

type DmJWTValMethod struct {
	Decrypt    types.Bool `tfsdk:"decrypt"`
	Verify     types.Bool `tfsdk:"verify"`
	Customized types.Bool `tfsdk:"customized"`
}

var DmJWTValMethodObjectType = map[string]attr.Type{
	"decrypt":    types.BoolType,
	"verify":     types.BoolType,
	"customized": types.BoolType,
}
var DmJWTValMethodObjectDefault = map[string]attr.Value{
	"decrypt":    types.BoolValue(false),
	"verify":     types.BoolValue(false),
	"customized": types.BoolValue(false),
}
var DmJWTValMethodDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"decrypt": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Decrypt", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"verify": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Verify", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"customized": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Custom processing", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
	},
}
var DmJWTValMethodResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmJWTValMethodObjectType,
			DmJWTValMethodObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"decrypt": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Decrypt", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"verify": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Verify", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"customized": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Custom processing", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
	},
}

func (data DmJWTValMethod) IsNull() bool {
	if !data.Decrypt.IsNull() {
		return false
	}
	if !data.Verify.IsNull() {
		return false
	}
	if !data.Customized.IsNull() {
		return false
	}
	return true
}
func GetDmJWTValMethodDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmJWTValMethodDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmJWTValMethodDataSourceSchema
}

func GetDmJWTValMethodResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmJWTValMethodResourceSchema.Required = true
	} else {
		DmJWTValMethodResourceSchema.Optional = true
		DmJWTValMethodResourceSchema.Computed = true
	}
	DmJWTValMethodResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmJWTValMethodResourceSchema
}

func (data DmJWTValMethod) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Decrypt.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`decrypt`, tfutils.StringFromBool(data.Decrypt))
	}
	if !data.Verify.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`verify`, tfutils.StringFromBool(data.Verify))
	}
	if !data.Customized.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`customized`, tfutils.StringFromBool(data.Customized))
	}
	return body
}

func (data *DmJWTValMethod) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `decrypt`); value.Exists() {
		data.Decrypt = tfutils.BoolFromString(value.String())
	} else {
		data.Decrypt = types.BoolNull()
	}
	if value := res.Get(pathRoot + `verify`); value.Exists() {
		data.Verify = tfutils.BoolFromString(value.String())
	} else {
		data.Verify = types.BoolNull()
	}
	if value := res.Get(pathRoot + `customized`); value.Exists() {
		data.Customized = tfutils.BoolFromString(value.String())
	} else {
		data.Customized = types.BoolNull()
	}
}

func (data *DmJWTValMethod) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `decrypt`); value.Exists() && !data.Decrypt.IsNull() {
		data.Decrypt = tfutils.BoolFromString(value.String())
	} else if data.Decrypt.ValueBool() {
		data.Decrypt = types.BoolNull()
	}
	if value := res.Get(pathRoot + `verify`); value.Exists() && !data.Verify.IsNull() {
		data.Verify = tfutils.BoolFromString(value.String())
	} else if data.Verify.ValueBool() {
		data.Verify = types.BoolNull()
	}
	if value := res.Get(pathRoot + `customized`); value.Exists() && !data.Customized.IsNull() {
		data.Customized = tfutils.BoolFromString(value.String())
	} else if data.Customized.ValueBool() {
		data.Customized = types.BoolNull()
	}
}
