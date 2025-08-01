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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	DataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	ResourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmAAAPMapResource struct {
	MrMethod               types.String `tfsdk:"mr_method"`
	MrCustomUrl            types.String `tfsdk:"mr_custom_url"`
	MrMapUrl               types.String `tfsdk:"mr_map_url"`
	MrMapXPath             types.String `tfsdk:"mr_map_x_path"`
	MrtamMap               types.String `tfsdk:"mrtam_map"`
	MrtamInstancePrefix    types.String `tfsdk:"mrtam_instance_prefix"`
	MrtamWebSealDynUrlFile types.String `tfsdk:"mrtam_web_seal_dyn_url_file"`
}

var DmAAAPMapResourceObjectType = map[string]attr.Type{
	"mr_method":                   types.StringType,
	"mr_custom_url":               types.StringType,
	"mr_map_url":                  types.StringType,
	"mr_map_x_path":               types.StringType,
	"mrtam_map":                   types.StringType,
	"mrtam_instance_prefix":       types.StringType,
	"mrtam_web_seal_dyn_url_file": types.StringType,
}
var DmAAAPMapResourceObjectDefault = map[string]attr.Value{
	"mr_method":                   types.StringValue("none"),
	"mr_custom_url":               types.StringNull(),
	"mr_map_url":                  types.StringNull(),
	"mr_map_x_path":               types.StringNull(),
	"mrtam_map":                   types.StringValue("WebSEAL"),
	"mrtam_instance_prefix":       types.StringNull(),
	"mrtam_web_seal_dyn_url_file": types.StringNull(),
}
var DmAAAPMapResourceDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"mr_method": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Method", "method", "").AddStringEnum("none", "custom", "xmlfile", "xpath", "tivoli").AddDefaultValue("none").String,
			Computed:            true,
		},
		"mr_custom_url": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Custom URL", "custom-url", "").String,
			Computed:            true,
		},
		"mr_map_url": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("AAA information file URL", "xmlfile-url", "").String,
			Computed:            true,
		},
		"mr_map_x_path": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("XPath expression", "xpath", "").String,
			Computed:            true,
		},
		"mrtam_map": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Object space mapping", "tam-mapping", "").AddStringEnum("TFIM", "TAMBI", "WebSEAL", "Custom").AddDefaultValue("WebSEAL").String,
			Computed:            true,
		},
		"mrtam_instance_prefix": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Object space prefix", "tam-prefix", "").String,
			Computed:            true,
		},
		"mrtam_web_seal_dyn_url_file": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WebSEAL DynURL mapping file", "webseal-dynurl-file", "").String,
			Computed:            true,
		},
	},
}
var DmAAAPMapResourceResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmAAAPMapResourceObjectType,
			DmAAAPMapResourceObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"mr_method": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Method", "method", "").AddStringEnum("none", "custom", "xmlfile", "xpath", "tivoli").AddDefaultValue("none").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("none", "custom", "xmlfile", "xpath", "tivoli"),
			},
			Default: stringdefault.StaticString("none"),
		},
		"mr_custom_url": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Custom URL", "custom-url", "").String,
			Optional:            true,
		},
		"mr_map_url": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("AAA information file URL", "xmlfile-url", "").String,
			Optional:            true,
		},
		"mr_map_x_path": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("XPath expression", "xpath", "").String,
			Optional:            true,
		},
		"mrtam_map": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Object space mapping", "tam-mapping", "").AddStringEnum("TFIM", "TAMBI", "WebSEAL", "Custom").AddDefaultValue("WebSEAL").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("TFIM", "TAMBI", "WebSEAL", "Custom"),
			},
			Default: stringdefault.StaticString("WebSEAL"),
		},
		"mrtam_instance_prefix": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Object space prefix", "tam-prefix", "").String,
			Optional:            true,
		},
		"mrtam_web_seal_dyn_url_file": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WebSEAL DynURL mapping file", "webseal-dynurl-file", "").String,
			Optional:            true,
		},
	},
}

func (data DmAAAPMapResource) IsNull() bool {
	if !data.MrMethod.IsNull() {
		return false
	}
	if !data.MrCustomUrl.IsNull() {
		return false
	}
	if !data.MrMapUrl.IsNull() {
		return false
	}
	if !data.MrMapXPath.IsNull() {
		return false
	}
	if !data.MrtamMap.IsNull() {
		return false
	}
	if !data.MrtamInstancePrefix.IsNull() {
		return false
	}
	if !data.MrtamWebSealDynUrlFile.IsNull() {
		return false
	}
	return true
}
func GetDmAAAPMapResourceDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmAAAPMapResourceDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmAAAPMapResourceDataSourceSchema
}

func GetDmAAAPMapResourceResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmAAAPMapResourceResourceSchema.Required = true
	} else {
		DmAAAPMapResourceResourceSchema.Optional = true
		DmAAAPMapResourceResourceSchema.Computed = true
	}
	DmAAAPMapResourceResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmAAAPMapResourceResourceSchema
}

func (data DmAAAPMapResource) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.MrMethod.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MRMethod`, data.MrMethod.ValueString())
	}
	if !data.MrCustomUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MRCustomURL`, data.MrCustomUrl.ValueString())
	}
	if !data.MrMapUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MRMapURL`, data.MrMapUrl.ValueString())
	}
	if !data.MrMapXPath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MRMapXPath`, data.MrMapXPath.ValueString())
	}
	if !data.MrtamMap.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MRTAMMap`, data.MrtamMap.ValueString())
	}
	if !data.MrtamInstancePrefix.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MRTAMInstancePrefix`, data.MrtamInstancePrefix.ValueString())
	}
	if !data.MrtamWebSealDynUrlFile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MRTAMWebSEALDynURLFile`, data.MrtamWebSealDynUrlFile.ValueString())
	}
	return body
}

func (data *DmAAAPMapResource) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `MRMethod`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MrMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MrMethod = types.StringValue("none")
	}
	if value := res.Get(pathRoot + `MRCustomURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MrCustomUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MrCustomUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `MRMapURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MrMapUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MrMapUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `MRMapXPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MrMapXPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MrMapXPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `MRTAMMap`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MrtamMap = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MrtamMap = types.StringValue("WebSEAL")
	}
	if value := res.Get(pathRoot + `MRTAMInstancePrefix`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MrtamInstancePrefix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MrtamInstancePrefix = types.StringNull()
	}
	if value := res.Get(pathRoot + `MRTAMWebSEALDynURLFile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MrtamWebSealDynUrlFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MrtamWebSealDynUrlFile = types.StringNull()
	}
}

func (data *DmAAAPMapResource) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `MRMethod`); value.Exists() && !data.MrMethod.IsNull() {
		data.MrMethod = tfutils.ParseStringFromGJSON(value)
	} else if data.MrMethod.ValueString() != "none" {
		data.MrMethod = types.StringNull()
	}
	if value := res.Get(pathRoot + `MRCustomURL`); value.Exists() && !data.MrCustomUrl.IsNull() {
		data.MrCustomUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MrCustomUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `MRMapURL`); value.Exists() && !data.MrMapUrl.IsNull() {
		data.MrMapUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MrMapUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `MRMapXPath`); value.Exists() && !data.MrMapXPath.IsNull() {
		data.MrMapXPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MrMapXPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `MRTAMMap`); value.Exists() && !data.MrtamMap.IsNull() {
		data.MrtamMap = tfutils.ParseStringFromGJSON(value)
	} else if data.MrtamMap.ValueString() != "WebSEAL" {
		data.MrtamMap = types.StringNull()
	}
	if value := res.Get(pathRoot + `MRTAMInstancePrefix`); value.Exists() && !data.MrtamInstancePrefix.IsNull() {
		data.MrtamInstancePrefix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MrtamInstancePrefix = types.StringNull()
	}
	if value := res.Get(pathRoot + `MRTAMWebSEALDynURLFile`); value.Exists() && !data.MrtamWebSealDynUrlFile.IsNull() {
		data.MrtamWebSealDynUrlFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MrtamWebSealDynUrlFile = types.StringNull()
	}
}
