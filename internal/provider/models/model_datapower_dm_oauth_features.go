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

type DmOAuthFeatures struct {
	Verboseerror             types.Bool `tfsdk:"verboseerror"`
	Onetimeuse               types.Bool `tfsdk:"onetimeuse"`
	Pkce                     types.Bool `tfsdk:"pkce"`
	Multipleusesrefreshtoken types.Bool `tfsdk:"multipleusesrefreshtoken"`
}

var DmOAuthFeaturesObjectType = map[string]attr.Type{
	"verboseerror":             types.BoolType,
	"onetimeuse":               types.BoolType,
	"pkce":                     types.BoolType,
	"multipleusesrefreshtoken": types.BoolType,
}
var DmOAuthFeaturesObjectDefault = map[string]attr.Value{
	"verboseerror":             types.BoolValue(false),
	"onetimeuse":               types.BoolValue(false),
	"pkce":                     types.BoolValue(false),
	"multipleusesrefreshtoken": types.BoolValue(false),
}
var DmOAuthFeaturesDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"verboseerror": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Verbose", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"onetimeuse": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("One-time use access token", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"pkce": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Require PKCE for Authorization Code", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"multipleusesrefreshtoken": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Allow reuse of refresh token", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
	},
}
var DmOAuthFeaturesResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmOAuthFeaturesObjectType,
			DmOAuthFeaturesObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"verboseerror": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Verbose", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"onetimeuse": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("One-time use access token", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"pkce": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Require PKCE for Authorization Code", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"multipleusesrefreshtoken": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Allow reuse of refresh token", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
	},
}

func (data DmOAuthFeatures) IsNull() bool {
	if !data.Verboseerror.IsNull() {
		return false
	}
	if !data.Onetimeuse.IsNull() {
		return false
	}
	if !data.Pkce.IsNull() {
		return false
	}
	if !data.Multipleusesrefreshtoken.IsNull() {
		return false
	}
	return true
}
func GetDmOAuthFeaturesDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmOAuthFeaturesDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmOAuthFeaturesDataSourceSchema
}

func GetDmOAuthFeaturesResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmOAuthFeaturesResourceSchema.Required = true
	} else {
		DmOAuthFeaturesResourceSchema.Optional = true
		DmOAuthFeaturesResourceSchema.Computed = true
	}
	DmOAuthFeaturesResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmOAuthFeaturesResourceSchema
}

func (data DmOAuthFeatures) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Verboseerror.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`verboseerror`, tfutils.StringFromBool(data.Verboseerror, ""))
	}
	if !data.Onetimeuse.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`onetimeuse`, tfutils.StringFromBool(data.Onetimeuse, ""))
	}
	if !data.Pkce.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`pkce`, tfutils.StringFromBool(data.Pkce, ""))
	}
	if !data.Multipleusesrefreshtoken.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`multipleusesrefreshtoken`, tfutils.StringFromBool(data.Multipleusesrefreshtoken, ""))
	}
	return body
}

func (data *DmOAuthFeatures) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `verboseerror`); value.Exists() {
		data.Verboseerror = tfutils.BoolFromString(value.String())
	} else {
		data.Verboseerror = types.BoolNull()
	}
	if value := res.Get(pathRoot + `onetimeuse`); value.Exists() {
		data.Onetimeuse = tfutils.BoolFromString(value.String())
	} else {
		data.Onetimeuse = types.BoolNull()
	}
	if value := res.Get(pathRoot + `pkce`); value.Exists() {
		data.Pkce = tfutils.BoolFromString(value.String())
	} else {
		data.Pkce = types.BoolNull()
	}
	if value := res.Get(pathRoot + `multipleusesrefreshtoken`); value.Exists() {
		data.Multipleusesrefreshtoken = tfutils.BoolFromString(value.String())
	} else {
		data.Multipleusesrefreshtoken = types.BoolNull()
	}
}

func (data *DmOAuthFeatures) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `verboseerror`); value.Exists() && !data.Verboseerror.IsNull() {
		data.Verboseerror = tfutils.BoolFromString(value.String())
	} else if data.Verboseerror.ValueBool() {
		data.Verboseerror = types.BoolNull()
	}
	if value := res.Get(pathRoot + `onetimeuse`); value.Exists() && !data.Onetimeuse.IsNull() {
		data.Onetimeuse = tfutils.BoolFromString(value.String())
	} else if data.Onetimeuse.ValueBool() {
		data.Onetimeuse = types.BoolNull()
	}
	if value := res.Get(pathRoot + `pkce`); value.Exists() && !data.Pkce.IsNull() {
		data.Pkce = tfutils.BoolFromString(value.String())
	} else if data.Pkce.ValueBool() {
		data.Pkce = types.BoolNull()
	}
	if value := res.Get(pathRoot + `multipleusesrefreshtoken`); value.Exists() && !data.Multipleusesrefreshtoken.IsNull() {
		data.Multipleusesrefreshtoken = tfutils.BoolFromString(value.String())
	} else if data.Multipleusesrefreshtoken.ValueBool() {
		data.Multipleusesrefreshtoken = types.BoolNull()
	}
}
