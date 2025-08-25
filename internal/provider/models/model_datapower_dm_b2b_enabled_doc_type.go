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

type DmB2BEnabledDocType struct {
	EnableXml     types.Bool `tfsdk:"enable_xml"`
	EnableX12     types.Bool `tfsdk:"enable_x12"`
	EnableEdifact types.Bool `tfsdk:"enable_edifact"`
	EnableBinary  types.Bool `tfsdk:"enable_binary"`
}

var DmB2BEnabledDocTypeObjectType = map[string]attr.Type{
	"enable_xml":     types.BoolType,
	"enable_x12":     types.BoolType,
	"enable_edifact": types.BoolType,
	"enable_binary":  types.BoolType,
}
var DmB2BEnabledDocTypeObjectDefault = map[string]attr.Value{
	"enable_xml":     types.BoolValue(true),
	"enable_x12":     types.BoolValue(true),
	"enable_edifact": types.BoolValue(true),
	"enable_binary":  types.BoolValue(true),
}

func GetDmB2BEnabledDocTypeDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmB2BEnabledDocTypeDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"enable_xml": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"enable_x12": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("X12", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"enable_edifact": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("EDIFACT", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"enable_binary": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Binary", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
		},
	}
	DmB2BEnabledDocTypeDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmB2BEnabledDocTypeDataSourceSchema
}
func GetDmB2BEnabledDocTypeResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmB2BEnabledDocTypeResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmB2BEnabledDocTypeObjectType,
				DmB2BEnabledDocTypeObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"enable_xml": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"enable_x12": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("X12", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"enable_edifact": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("EDIFACT", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"enable_binary": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Binary", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
		},
	}
	DmB2BEnabledDocTypeResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmB2BEnabledDocTypeResourceSchema.Required = true
	} else {
		DmB2BEnabledDocTypeResourceSchema.Optional = true
		DmB2BEnabledDocTypeResourceSchema.Computed = true
	}
	return DmB2BEnabledDocTypeResourceSchema
}

func (data DmB2BEnabledDocType) IsNull() bool {
	if !data.EnableXml.IsNull() {
		return false
	}
	if !data.EnableX12.IsNull() {
		return false
	}
	if !data.EnableEdifact.IsNull() {
		return false
	}
	if !data.EnableBinary.IsNull() {
		return false
	}
	return true
}

func (data DmB2BEnabledDocType) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.EnableXml.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnableXML`, tfutils.StringFromBool(data.EnableXml, ""))
	}
	if !data.EnableX12.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnableX12`, tfutils.StringFromBool(data.EnableX12, ""))
	}
	if !data.EnableEdifact.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnableEDIFACT`, tfutils.StringFromBool(data.EnableEdifact, ""))
	}
	if !data.EnableBinary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnableBinary`, tfutils.StringFromBool(data.EnableBinary, ""))
	}
	return body
}

func (data *DmB2BEnabledDocType) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `EnableXML`); value.Exists() {
		data.EnableXml = tfutils.BoolFromString(value.String())
	} else {
		data.EnableXml = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EnableX12`); value.Exists() {
		data.EnableX12 = tfutils.BoolFromString(value.String())
	} else {
		data.EnableX12 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EnableEDIFACT`); value.Exists() {
		data.EnableEdifact = tfutils.BoolFromString(value.String())
	} else {
		data.EnableEdifact = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EnableBinary`); value.Exists() {
		data.EnableBinary = tfutils.BoolFromString(value.String())
	} else {
		data.EnableBinary = types.BoolNull()
	}
}

func (data *DmB2BEnabledDocType) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `EnableXML`); value.Exists() && !data.EnableXml.IsNull() {
		data.EnableXml = tfutils.BoolFromString(value.String())
	} else if !data.EnableXml.ValueBool() {
		data.EnableXml = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EnableX12`); value.Exists() && !data.EnableX12.IsNull() {
		data.EnableX12 = tfutils.BoolFromString(value.String())
	} else if !data.EnableX12.ValueBool() {
		data.EnableX12 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EnableEDIFACT`); value.Exists() && !data.EnableEdifact.IsNull() {
		data.EnableEdifact = tfutils.BoolFromString(value.String())
	} else if !data.EnableEdifact.ValueBool() {
		data.EnableEdifact = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EnableBinary`); value.Exists() && !data.EnableBinary.IsNull() {
		data.EnableBinary = tfutils.BoolFromString(value.String())
	} else if !data.EnableBinary.ValueBool() {
		data.EnableBinary = types.BoolNull()
	}
}
