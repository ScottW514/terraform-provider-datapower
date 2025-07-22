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

type DmSAMLStatementType struct {
	Authentication types.Bool `tfsdk:"authentication"`
	Attribute      types.Bool `tfsdk:"attribute"`
	Authorization  types.Bool `tfsdk:"authorization"`
}

var DmSAMLStatementTypeObjectType = map[string]attr.Type{
	"authentication": types.BoolType,
	"attribute":      types.BoolType,
	"authorization":  types.BoolType,
}
var DmSAMLStatementTypeObjectDefault = map[string]attr.Value{
	"authentication": types.BoolValue(true),
	"attribute":      types.BoolValue(true),
	"authorization":  types.BoolValue(false),
}
var DmSAMLStatementTypeDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"authentication": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Authentication statement", "", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"attribute": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Attribute statement", "", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"authorization": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Authorization decision statement", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
	},
}
var DmSAMLStatementTypeResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmSAMLStatementTypeObjectType,
			DmSAMLStatementTypeObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"authentication": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Authentication statement", "", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"attribute": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Attribute statement", "", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"authorization": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Authorization decision statement", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
	},
}

func (data DmSAMLStatementType) IsNull() bool {
	if !data.Authentication.IsNull() {
		return false
	}
	if !data.Attribute.IsNull() {
		return false
	}
	if !data.Authorization.IsNull() {
		return false
	}
	return true
}
func GetDmSAMLStatementTypeDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmSAMLStatementTypeDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmSAMLStatementTypeDataSourceSchema
}

func GetDmSAMLStatementTypeResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmSAMLStatementTypeResourceSchema.Required = true
	} else {
		DmSAMLStatementTypeResourceSchema.Optional = true
		DmSAMLStatementTypeResourceSchema.Computed = true
	}
	DmSAMLStatementTypeResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmSAMLStatementTypeResourceSchema
}

func (data DmSAMLStatementType) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Authentication.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`authentication`, tfutils.StringFromBool(data.Authentication, false))
	}
	if !data.Attribute.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`attribute`, tfutils.StringFromBool(data.Attribute, false))
	}
	if !data.Authorization.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`authorization`, tfutils.StringFromBool(data.Authorization, false))
	}
	return body
}

func (data *DmSAMLStatementType) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `authentication`); value.Exists() {
		data.Authentication = tfutils.BoolFromString(value.String())
	} else {
		data.Authentication = types.BoolNull()
	}
	if value := res.Get(pathRoot + `attribute`); value.Exists() {
		data.Attribute = tfutils.BoolFromString(value.String())
	} else {
		data.Attribute = types.BoolNull()
	}
	if value := res.Get(pathRoot + `authorization`); value.Exists() {
		data.Authorization = tfutils.BoolFromString(value.String())
	} else {
		data.Authorization = types.BoolNull()
	}
}

func (data *DmSAMLStatementType) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `authentication`); value.Exists() && !data.Authentication.IsNull() {
		data.Authentication = tfutils.BoolFromString(value.String())
	} else if !data.Authentication.ValueBool() {
		data.Authentication = types.BoolNull()
	}
	if value := res.Get(pathRoot + `attribute`); value.Exists() && !data.Attribute.IsNull() {
		data.Attribute = tfutils.BoolFromString(value.String())
	} else if !data.Attribute.ValueBool() {
		data.Attribute = types.BoolNull()
	}
	if value := res.Get(pathRoot + `authorization`); value.Exists() && !data.Authorization.IsNull() {
		data.Authorization = tfutils.BoolFromString(value.String())
	} else if data.Authorization.ValueBool() {
		data.Authorization = types.BoolNull()
	}
}
