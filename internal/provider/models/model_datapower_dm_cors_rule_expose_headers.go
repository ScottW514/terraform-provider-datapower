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

type DmCORSRuleExposeHeaders struct {
	Predefined types.Bool   `tfsdk:"predefined"`
	Backend    types.Bool   `tfsdk:"backend"`
	Custom     types.String `tfsdk:"custom"`
}

var DmCORSRuleExposeHeadersObjectType = map[string]attr.Type{
	"predefined": types.BoolType,
	"backend":    types.BoolType,
	"custom":     types.StringType,
}
var DmCORSRuleExposeHeadersObjectDefault = map[string]attr.Value{
	"predefined": types.BoolValue(true),
	"backend":    types.BoolValue(false),
	"custom":     types.StringNull(),
}

func GetDmCORSRuleExposeHeadersDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmCORSRuleExposeHeadersDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"predefined": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Append the gateway-predefined value.", "predefined", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"backend": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Append the value in the response from the target endpoint.", "backend", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"custom": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Append a custom string as the value.", "custom", "").String,
				Computed:            true,
			},
		},
	}
	DmCORSRuleExposeHeadersDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmCORSRuleExposeHeadersDataSourceSchema
}
func GetDmCORSRuleExposeHeadersResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmCORSRuleExposeHeadersResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmCORSRuleExposeHeadersObjectType,
				DmCORSRuleExposeHeadersObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"predefined": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Append the gateway-predefined value.", "predefined", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"backend": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Append the value in the response from the target endpoint.", "backend", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"custom": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Append a custom string as the value.", "custom", "").String,
				Optional:            true,
			},
		},
	}
	DmCORSRuleExposeHeadersResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmCORSRuleExposeHeadersResourceSchema.Required = true
	} else {
		DmCORSRuleExposeHeadersResourceSchema.Optional = true
		DmCORSRuleExposeHeadersResourceSchema.Computed = true
	}
	return DmCORSRuleExposeHeadersResourceSchema
}

func (data DmCORSRuleExposeHeaders) IsNull() bool {
	if !data.Predefined.IsNull() {
		return false
	}
	if !data.Backend.IsNull() {
		return false
	}
	if !data.Custom.IsNull() {
		return false
	}
	return true
}

func (data DmCORSRuleExposeHeaders) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Predefined.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Predefined`, tfutils.StringFromBool(data.Predefined, ""))
	}
	if !data.Backend.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Backend`, tfutils.StringFromBool(data.Backend, ""))
	}
	if !data.Custom.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Custom`, data.Custom.ValueString())
	}
	return body
}

func (data *DmCORSRuleExposeHeaders) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Predefined`); value.Exists() {
		data.Predefined = tfutils.BoolFromString(value.String())
	} else {
		data.Predefined = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Backend`); value.Exists() {
		data.Backend = tfutils.BoolFromString(value.String())
	} else {
		data.Backend = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Custom`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Custom = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Custom = types.StringNull()
	}
}

func (data *DmCORSRuleExposeHeaders) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Predefined`); value.Exists() && !data.Predefined.IsNull() {
		data.Predefined = tfutils.BoolFromString(value.String())
	} else if !data.Predefined.ValueBool() {
		data.Predefined = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Backend`); value.Exists() && !data.Backend.IsNull() {
		data.Backend = tfutils.BoolFromString(value.String())
	} else if data.Backend.ValueBool() {
		data.Backend = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Custom`); value.Exists() && !data.Custom.IsNull() {
		data.Custom = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Custom = types.StringNull()
	}
}
