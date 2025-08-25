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

type DmHTTPVersionMask struct {
	Http10 types.Bool `tfsdk:"http_1_0"`
	Http11 types.Bool `tfsdk:"http_1_1"`
}

var DmHTTPVersionMaskObjectType = map[string]attr.Type{
	"http_1_0": types.BoolType,
	"http_1_1": types.BoolType,
}
var DmHTTPVersionMaskObjectDefault = map[string]attr.Value{
	"http_1_0": types.BoolValue(true),
	"http_1_1": types.BoolValue(true),
}

func GetDmHTTPVersionMaskDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmHTTPVersionMaskDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"http_1_0": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP 1.0", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"http_1_1": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP 1.1", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
		},
	}
	DmHTTPVersionMaskDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmHTTPVersionMaskDataSourceSchema
}
func GetDmHTTPVersionMaskResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmHTTPVersionMaskResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmHTTPVersionMaskObjectType,
				DmHTTPVersionMaskObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"http_1_0": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP 1.0", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"http_1_1": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP 1.1", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
		},
	}
	DmHTTPVersionMaskResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmHTTPVersionMaskResourceSchema.Required = true
	} else {
		DmHTTPVersionMaskResourceSchema.Optional = true
		DmHTTPVersionMaskResourceSchema.Computed = true
	}
	return DmHTTPVersionMaskResourceSchema
}

func (data DmHTTPVersionMask) IsNull() bool {
	if !data.Http10.IsNull() {
		return false
	}
	if !data.Http11.IsNull() {
		return false
	}
	return true
}

func (data DmHTTPVersionMask) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Http10.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-1\.0`, tfutils.StringFromBool(data.Http10, ""))
	}
	if !data.Http11.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-1\.1`, tfutils.StringFromBool(data.Http11, ""))
	}
	return body
}

func (data *DmHTTPVersionMask) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `HTTP-1\.0`); value.Exists() {
		data.Http10 = tfutils.BoolFromString(value.String())
	} else {
		data.Http10 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-1\.1`); value.Exists() {
		data.Http11 = tfutils.BoolFromString(value.String())
	} else {
		data.Http11 = types.BoolNull()
	}
}

func (data *DmHTTPVersionMask) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `HTTP-1\.0`); value.Exists() && !data.Http10.IsNull() {
		data.Http10 = tfutils.BoolFromString(value.String())
	} else if !data.Http10.ValueBool() {
		data.Http10 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-1\.1`); value.Exists() && !data.Http11.IsNull() {
		data.Http11 = tfutils.BoolFromString(value.String())
	} else if !data.Http11.ValueBool() {
		data.Http11 = types.BoolNull()
	}
}
