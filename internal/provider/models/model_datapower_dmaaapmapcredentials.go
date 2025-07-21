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

type DmAAAPMapCredentials struct {
	McMethod       types.String `tfsdk:"mc_method"`
	McCustomUrl    types.String `tfsdk:"mc_custom_url"`
	McMapUrl       types.String `tfsdk:"mc_map_url"`
	McMapXPath     types.String `tfsdk:"mc_map_x_path"`
	MctfimEndpoint types.String `tfsdk:"mctfim_endpoint"`
}

var DmAAAPMapCredentialsObjectType = map[string]attr.Type{
	"mc_method":       types.StringType,
	"mc_custom_url":   types.StringType,
	"mc_map_url":      types.StringType,
	"mc_map_x_path":   types.StringType,
	"mctfim_endpoint": types.StringType,
}
var DmAAAPMapCredentialsObjectDefault = map[string]attr.Value{
	"mc_method":       types.StringValue("none"),
	"mc_custom_url":   types.StringNull(),
	"mc_map_url":      types.StringNull(),
	"mc_map_x_path":   types.StringNull(),
	"mctfim_endpoint": types.StringNull(),
}
var DmAAAPMapCredentialsDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"mc_method": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Method", "method", "").AddStringEnum("none", "custom", "xmlfile", "xpath", "ws-secureconversation", "TFIM").AddDefaultValue("none").String,
			Computed:            true,
		},
		"mc_custom_url": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Custom URL", "custom-url", "").String,
			Computed:            true,
		},
		"mc_map_url": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("AAA information file URL", "xmlfile-url", "").String,
			Computed:            true,
		},
		"mc_map_x_path": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("XPath expression", "xpath", "").String,
			Computed:            true,
		},
		"mctfim_endpoint": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Federated Identity Manager endpoint (deprecated)", "tfim", "tfimendpoint").String,
			Computed:            true,
		},
	},
}
var DmAAAPMapCredentialsResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmAAAPMapCredentialsObjectType,
			DmAAAPMapCredentialsObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"mc_method": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Method", "method", "").AddStringEnum("none", "custom", "xmlfile", "xpath", "ws-secureconversation", "TFIM").AddDefaultValue("none").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("none", "custom", "xmlfile", "xpath", "ws-secureconversation", "TFIM"),
			},
			Default: stringdefault.StaticString("none"),
		},
		"mc_custom_url": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Custom URL", "custom-url", "").String,
			Optional:            true,
		},
		"mc_map_url": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("AAA information file URL", "xmlfile-url", "").String,
			Optional:            true,
		},
		"mc_map_x_path": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("XPath expression", "xpath", "").String,
			Optional:            true,
		},
		"mctfim_endpoint": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Federated Identity Manager endpoint (deprecated)", "tfim", "tfimendpoint").String,
			Optional:            true,
		},
	},
}

func (data DmAAAPMapCredentials) IsNull() bool {
	if !data.McMethod.IsNull() {
		return false
	}
	if !data.McCustomUrl.IsNull() {
		return false
	}
	if !data.McMapUrl.IsNull() {
		return false
	}
	if !data.McMapXPath.IsNull() {
		return false
	}
	if !data.MctfimEndpoint.IsNull() {
		return false
	}
	return true
}
func GetDmAAAPMapCredentialsDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmAAAPMapCredentialsDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmAAAPMapCredentialsDataSourceSchema
}

func GetDmAAAPMapCredentialsResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmAAAPMapCredentialsResourceSchema.Required = true
	} else {
		DmAAAPMapCredentialsResourceSchema.Optional = true
		DmAAAPMapCredentialsResourceSchema.Computed = true
	}
	DmAAAPMapCredentialsResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmAAAPMapCredentialsResourceSchema
}

func (data DmAAAPMapCredentials) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.McMethod.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCMethod`, data.McMethod.ValueString())
	}
	if !data.McCustomUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCCustomURL`, data.McCustomUrl.ValueString())
	}
	if !data.McMapUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCMapURL`, data.McMapUrl.ValueString())
	}
	if !data.McMapXPath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCMapXPath`, data.McMapXPath.ValueString())
	}
	if !data.MctfimEndpoint.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCTFIMEndpoint`, data.MctfimEndpoint.ValueString())
	}
	return body
}

func (data *DmAAAPMapCredentials) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `MCMethod`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.McMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McMethod = types.StringValue("none")
	}
	if value := res.Get(pathRoot + `MCCustomURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.McCustomUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McCustomUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCMapURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.McMapUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McMapUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCMapXPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.McMapXPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McMapXPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCTFIMEndpoint`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MctfimEndpoint = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MctfimEndpoint = types.StringNull()
	}
}

func (data *DmAAAPMapCredentials) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `MCMethod`); value.Exists() && !data.McMethod.IsNull() {
		data.McMethod = tfutils.ParseStringFromGJSON(value)
	} else if data.McMethod.ValueString() != "none" {
		data.McMethod = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCCustomURL`); value.Exists() && !data.McCustomUrl.IsNull() {
		data.McCustomUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McCustomUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCMapURL`); value.Exists() && !data.McMapUrl.IsNull() {
		data.McMapUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McMapUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCMapXPath`); value.Exists() && !data.McMapXPath.IsNull() {
		data.McMapXPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McMapXPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCTFIMEndpoint`); value.Exists() && !data.MctfimEndpoint.IsNull() {
		data.MctfimEndpoint = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MctfimEndpoint = types.StringNull()
	}
}
