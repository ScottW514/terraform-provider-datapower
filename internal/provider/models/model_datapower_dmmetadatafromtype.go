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

type DmMetadataFromType struct {
	None              types.Bool `tfsdk:"none"`
	AuthenticationUrl types.Bool `tfsdk:"authentication_url"`
	ExternalUrl       types.Bool `tfsdk:"external_url"`
}

var DmMetadataFromTypeObjectType = map[string]attr.Type{
	"none":               types.BoolType,
	"authentication_url": types.BoolType,
	"external_url":       types.BoolType,
}
var DmMetadataFromTypeObjectDefault = map[string]attr.Value{
	"none":               types.BoolValue(false),
	"authentication_url": types.BoolValue(true),
	"external_url":       types.BoolValue(false),
}
var DmMetadataFromTypeDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"none": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"authentication_url": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Authentication URL", "", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"external_url": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("External metadata URL", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
	},
}
var DmMetadataFromTypeResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmMetadataFromTypeObjectType,
			DmMetadataFromTypeObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"none": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"authentication_url": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Authentication URL", "", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"external_url": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("External metadata URL", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
	},
}

func (data DmMetadataFromType) IsNull() bool {
	if !data.None.IsNull() {
		return false
	}
	if !data.AuthenticationUrl.IsNull() {
		return false
	}
	if !data.ExternalUrl.IsNull() {
		return false
	}
	return true
}
func GetDmMetadataFromTypeDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmMetadataFromTypeDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmMetadataFromTypeDataSourceSchema
}

func GetDmMetadataFromTypeResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmMetadataFromTypeResourceSchema.Required = true
	} else {
		DmMetadataFromTypeResourceSchema.Optional = true
		DmMetadataFromTypeResourceSchema.Computed = true
	}
	DmMetadataFromTypeResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmMetadataFromTypeResourceSchema
}

func (data DmMetadataFromType) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.None.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`none`, tfutils.StringFromBool(data.None, ""))
	}
	if !data.AuthenticationUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`authentication_url`, tfutils.StringFromBool(data.AuthenticationUrl, ""))
	}
	if !data.ExternalUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`external_url`, tfutils.StringFromBool(data.ExternalUrl, ""))
	}
	return body
}

func (data *DmMetadataFromType) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `none`); value.Exists() {
		data.None = tfutils.BoolFromString(value.String())
	} else {
		data.None = types.BoolNull()
	}
	if value := res.Get(pathRoot + `authentication_url`); value.Exists() {
		data.AuthenticationUrl = tfutils.BoolFromString(value.String())
	} else {
		data.AuthenticationUrl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `external_url`); value.Exists() {
		data.ExternalUrl = tfutils.BoolFromString(value.String())
	} else {
		data.ExternalUrl = types.BoolNull()
	}
}

func (data *DmMetadataFromType) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `none`); value.Exists() && !data.None.IsNull() {
		data.None = tfutils.BoolFromString(value.String())
	} else if data.None.ValueBool() {
		data.None = types.BoolNull()
	}
	if value := res.Get(pathRoot + `authentication_url`); value.Exists() && !data.AuthenticationUrl.IsNull() {
		data.AuthenticationUrl = tfutils.BoolFromString(value.String())
	} else if !data.AuthenticationUrl.ValueBool() {
		data.AuthenticationUrl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `external_url`); value.Exists() && !data.ExternalUrl.IsNull() {
		data.ExternalUrl = tfutils.BoolFromString(value.String())
	} else if data.ExternalUrl.ValueBool() {
		data.ExternalUrl = types.BoolNull()
	}
}
