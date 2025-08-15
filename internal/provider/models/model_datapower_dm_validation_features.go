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

type DmValidationFeatures struct {
	Noauthen   types.Bool `tfsdk:"noauthen"`
	Introspect types.Bool `tfsdk:"introspect"`
}

var DmValidationFeaturesObjectType = map[string]attr.Type{
	"noauthen":   types.BoolType,
	"introspect": types.BoolType,
}
var DmValidationFeaturesObjectDefault = map[string]attr.Value{
	"noauthen":   types.BoolValue(false),
	"introspect": types.BoolValue(false),
}
var DmValidationFeaturesDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"noauthen": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Unprotected", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"introspect": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Introspection Format", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
	},
}
var DmValidationFeaturesResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmValidationFeaturesObjectType,
			DmValidationFeaturesObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"noauthen": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Unprotected", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"introspect": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Introspection Format", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
	},
}

func (data DmValidationFeatures) IsNull() bool {
	if !data.Noauthen.IsNull() {
		return false
	}
	if !data.Introspect.IsNull() {
		return false
	}
	return true
}
func GetDmValidationFeaturesDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmValidationFeaturesDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmValidationFeaturesDataSourceSchema
}

func GetDmValidationFeaturesResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmValidationFeaturesResourceSchema.Required = true
	} else {
		DmValidationFeaturesResourceSchema.Optional = true
		DmValidationFeaturesResourceSchema.Computed = true
	}
	DmValidationFeaturesResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmValidationFeaturesResourceSchema
}

func (data DmValidationFeatures) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Noauthen.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`noauthen`, tfutils.StringFromBool(data.Noauthen, ""))
	}
	if !data.Introspect.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`introspect`, tfutils.StringFromBool(data.Introspect, ""))
	}
	return body
}

func (data *DmValidationFeatures) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `noauthen`); value.Exists() {
		data.Noauthen = tfutils.BoolFromString(value.String())
	} else {
		data.Noauthen = types.BoolNull()
	}
	if value := res.Get(pathRoot + `introspect`); value.Exists() {
		data.Introspect = tfutils.BoolFromString(value.String())
	} else {
		data.Introspect = types.BoolNull()
	}
}

func (data *DmValidationFeatures) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `noauthen`); value.Exists() && !data.Noauthen.IsNull() {
		data.Noauthen = tfutils.BoolFromString(value.String())
	} else if data.Noauthen.ValueBool() {
		data.Noauthen = types.BoolNull()
	}
	if value := res.Get(pathRoot + `introspect`); value.Exists() && !data.Introspect.IsNull() {
		data.Introspect = tfutils.BoolFromString(value.String())
	} else if data.Introspect.ValueBool() {
		data.Introspect = types.BoolNull()
	}
}
