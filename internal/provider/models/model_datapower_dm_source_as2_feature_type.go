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

type DmSourceAS2FeatureType struct {
	Http1D0             types.Bool `tfsdk:"http_1d0"`
	Http1D1             types.Bool `tfsdk:"http_1d1"`
	QueryString         types.Bool `tfsdk:"query_string"`
	FragmentIdentifiers types.Bool `tfsdk:"fragment_identifiers"`
	DotDot              types.Bool `tfsdk:"dot_dot"`
	CmdExe              types.Bool `tfsdk:"cmd_exe"`
}

var DmSourceAS2FeatureTypeObjectType = map[string]attr.Type{
	"http_1d0":             types.BoolType,
	"http_1d1":             types.BoolType,
	"query_string":         types.BoolType,
	"fragment_identifiers": types.BoolType,
	"dot_dot":              types.BoolType,
	"cmd_exe":              types.BoolType,
}
var DmSourceAS2FeatureTypeObjectDefault = map[string]attr.Value{
	"http_1d0":             types.BoolValue(true),
	"http_1d1":             types.BoolValue(true),
	"query_string":         types.BoolValue(true),
	"fragment_identifiers": types.BoolValue(true),
	"dot_dot":              types.BoolValue(false),
	"cmd_exe":              types.BoolValue(false),
}

func GetDmSourceAS2FeatureTypeDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmSourceAS2FeatureTypeDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"http_1d0": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Supports HTTP 1.0 requests", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"http_1d1": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Supports HTTP 1.1 requests", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"query_string": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allows requests when the URL contains a ? (query string)", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"fragment_identifiers": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("UAllows requests when the URL contains a # (fragment identifier)", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"dot_dot": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allows requests when the URL contains a .. (dotdot)", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"cmd_exe": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allows requests when the URL contains cmd.exe", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
		},
	}
	DmSourceAS2FeatureTypeDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmSourceAS2FeatureTypeDataSourceSchema
}
func GetDmSourceAS2FeatureTypeResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmSourceAS2FeatureTypeResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmSourceAS2FeatureTypeObjectType,
				DmSourceAS2FeatureTypeObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"http_1d0": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Supports HTTP 1.0 requests", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"http_1d1": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Supports HTTP 1.1 requests", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"query_string": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allows requests when the URL contains a ? (query string)", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"fragment_identifiers": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("UAllows requests when the URL contains a # (fragment identifier)", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"dot_dot": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allows requests when the URL contains a .. (dotdot)", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"cmd_exe": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allows requests when the URL contains cmd.exe", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
	DmSourceAS2FeatureTypeResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmSourceAS2FeatureTypeResourceSchema.Required = true
	} else {
		DmSourceAS2FeatureTypeResourceSchema.Optional = true
		DmSourceAS2FeatureTypeResourceSchema.Computed = true
	}
	return DmSourceAS2FeatureTypeResourceSchema
}

func (data DmSourceAS2FeatureType) IsNull() bool {
	if !data.Http1D0.IsNull() {
		return false
	}
	if !data.Http1D1.IsNull() {
		return false
	}
	if !data.QueryString.IsNull() {
		return false
	}
	if !data.FragmentIdentifiers.IsNull() {
		return false
	}
	if !data.DotDot.IsNull() {
		return false
	}
	if !data.CmdExe.IsNull() {
		return false
	}
	return true
}

func (data DmSourceAS2FeatureType) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Http1D0.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-1\.0`, tfutils.StringFromBool(data.Http1D0, ""))
	}
	if !data.Http1D1.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-1\.1`, tfutils.StringFromBool(data.Http1D1, ""))
	}
	if !data.QueryString.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`QueryString`, tfutils.StringFromBool(data.QueryString, ""))
	}
	if !data.FragmentIdentifiers.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FragmentIdentifiers`, tfutils.StringFromBool(data.FragmentIdentifiers, ""))
	}
	if !data.DotDot.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DotDot`, tfutils.StringFromBool(data.DotDot, ""))
	}
	if !data.CmdExe.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CmdExe`, tfutils.StringFromBool(data.CmdExe, ""))
	}
	return body
}

func (data *DmSourceAS2FeatureType) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `HTTP-1\.0`); value.Exists() {
		data.Http1D0 = tfutils.BoolFromString(value.String())
	} else {
		data.Http1D0 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-1\.1`); value.Exists() {
		data.Http1D1 = tfutils.BoolFromString(value.String())
	} else {
		data.Http1D1 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `QueryString`); value.Exists() {
		data.QueryString = tfutils.BoolFromString(value.String())
	} else {
		data.QueryString = types.BoolNull()
	}
	if value := res.Get(pathRoot + `FragmentIdentifiers`); value.Exists() {
		data.FragmentIdentifiers = tfutils.BoolFromString(value.String())
	} else {
		data.FragmentIdentifiers = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DotDot`); value.Exists() {
		data.DotDot = tfutils.BoolFromString(value.String())
	} else {
		data.DotDot = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CmdExe`); value.Exists() {
		data.CmdExe = tfutils.BoolFromString(value.String())
	} else {
		data.CmdExe = types.BoolNull()
	}
}

func (data *DmSourceAS2FeatureType) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `HTTP-1\.0`); value.Exists() && !data.Http1D0.IsNull() {
		data.Http1D0 = tfutils.BoolFromString(value.String())
	} else if !data.Http1D0.ValueBool() {
		data.Http1D0 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-1\.1`); value.Exists() && !data.Http1D1.IsNull() {
		data.Http1D1 = tfutils.BoolFromString(value.String())
	} else if !data.Http1D1.ValueBool() {
		data.Http1D1 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `QueryString`); value.Exists() && !data.QueryString.IsNull() {
		data.QueryString = tfutils.BoolFromString(value.String())
	} else if !data.QueryString.ValueBool() {
		data.QueryString = types.BoolNull()
	}
	if value := res.Get(pathRoot + `FragmentIdentifiers`); value.Exists() && !data.FragmentIdentifiers.IsNull() {
		data.FragmentIdentifiers = tfutils.BoolFromString(value.String())
	} else if !data.FragmentIdentifiers.ValueBool() {
		data.FragmentIdentifiers = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DotDot`); value.Exists() && !data.DotDot.IsNull() {
		data.DotDot = tfutils.BoolFromString(value.String())
	} else if data.DotDot.ValueBool() {
		data.DotDot = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CmdExe`); value.Exists() && !data.CmdExe.IsNull() {
		data.CmdExe = tfutils.BoolFromString(value.String())
	} else if data.CmdExe.ValueBool() {
		data.CmdExe = types.BoolNull()
	}
}
