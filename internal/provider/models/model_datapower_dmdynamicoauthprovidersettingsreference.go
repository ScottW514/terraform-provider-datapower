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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmDynamicOAuthProviderSettingsReference struct {
	Url     types.String `tfsdk:"url"`
	Literal types.String `tfsdk:"literal"`
	Default types.String `tfsdk:"default"`
}

var DmDynamicOAuthProviderSettingsReferenceObjectType = map[string]attr.Type{
	"url":     types.StringType,
	"literal": types.StringType,
	"default": types.StringType,
}
var DmDynamicOAuthProviderSettingsReferenceObjectDefault = map[string]attr.Value{
	"url":     types.StringNull(),
	"literal": types.StringNull(),
	"default": types.StringNull(),
}
var DmDynamicOAuthProviderSettingsReferenceDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"url": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL reference", "url", "").String,
			Computed:            true,
		},
		"literal": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Literal configuration", "literal", "").String,
			Computed:            true,
		},
		"default": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Object reference", "default", "oauthprovidersettings").String,
			Computed:            true,
		},
	},
}
var DmDynamicOAuthProviderSettingsReferenceResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmDynamicOAuthProviderSettingsReferenceObjectType,
			DmDynamicOAuthProviderSettingsReferenceObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"url": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL reference", "url", "").String,
			Optional:            true,
		},
		"literal": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Literal configuration", "literal", "").String,
			Optional:            true,
		},
		"default": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Object reference", "default", "oauthprovidersettings").String,
			Optional:            true,
		},
	},
}

func (data DmDynamicOAuthProviderSettingsReference) IsNull() bool {
	if !data.Url.IsNull() {
		return false
	}
	if !data.Literal.IsNull() {
		return false
	}
	if !data.Default.IsNull() {
		return false
	}
	return true
}
func GetDmDynamicOAuthProviderSettingsReferenceDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmDynamicOAuthProviderSettingsReferenceDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmDynamicOAuthProviderSettingsReferenceDataSourceSchema
}

func GetDmDynamicOAuthProviderSettingsReferenceResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmDynamicOAuthProviderSettingsReferenceResourceSchema.Required = true
	} else {
		DmDynamicOAuthProviderSettingsReferenceResourceSchema.Optional = true
		DmDynamicOAuthProviderSettingsReferenceResourceSchema.Computed = true
	}
	DmDynamicOAuthProviderSettingsReferenceResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmDynamicOAuthProviderSettingsReferenceResourceSchema
}

func (data DmDynamicOAuthProviderSettingsReference) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Url.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`URL`, data.Url.ValueString())
	}
	if !data.Literal.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Literal`, data.Literal.ValueString())
	}
	if !data.Default.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Default`, data.Default.ValueString())
	}
	return body
}

func (data *DmDynamicOAuthProviderSettingsReference) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `URL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Url = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Url = types.StringNull()
	}
	if value := res.Get(pathRoot + `Literal`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Literal = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Literal = types.StringNull()
	}
	if value := res.Get(pathRoot + `Default`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Default = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Default = types.StringNull()
	}
}

func (data *DmDynamicOAuthProviderSettingsReference) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `URL`); value.Exists() && !data.Url.IsNull() {
		data.Url = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Url = types.StringNull()
	}
	if value := res.Get(pathRoot + `Literal`); value.Exists() && !data.Literal.IsNull() {
		data.Literal = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Literal = types.StringNull()
	}
	if value := res.Get(pathRoot + `Default`); value.Exists() && !data.Default.IsNull() {
		data.Default = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Default = types.StringNull()
	}
}
