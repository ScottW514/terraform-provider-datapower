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

type DmAAAPERBitmap struct {
	TargetUrl     types.Bool `tfsdk:"target_url"`
	OriginalUrl   types.Bool `tfsdk:"original_url"`
	RequestUri    types.Bool `tfsdk:"request_uri"`
	RequestOpname types.Bool `tfsdk:"request_opname"`
	HttpMethod    types.Bool `tfsdk:"http_method"`
	XPath         types.Bool `tfsdk:"x_path"`
	Metadata      types.Bool `tfsdk:"metadata"`
}

var DmAAAPERBitmapObjectType = map[string]attr.Type{
	"target_url":     types.BoolType,
	"original_url":   types.BoolType,
	"request_uri":    types.BoolType,
	"request_opname": types.BoolType,
	"http_method":    types.BoolType,
	"x_path":         types.BoolType,
	"metadata":       types.BoolType,
}
var DmAAAPERBitmapObjectDefault = map[string]attr.Value{
	"target_url":     types.BoolValue(false),
	"original_url":   types.BoolValue(false),
	"request_uri":    types.BoolValue(false),
	"request_opname": types.BoolValue(false),
	"http_method":    types.BoolValue(false),
	"x_path":         types.BoolValue(false),
	"metadata":       types.BoolValue(false),
}
var DmAAAPERBitmapDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"target_url": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL sent to back end", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"original_url": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL sent by client", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"request_uri": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URI of top level element in message", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"request_opname": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Local name of request element", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"http_method": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("HTTP operation (GET or POST)", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"x_path": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("XPath expression", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"metadata": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Processing metadata", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
	},
}
var DmAAAPERBitmapResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmAAAPERBitmapObjectType,
			DmAAAPERBitmapObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"target_url": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL sent to back end", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"original_url": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL sent by client", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"request_uri": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URI of top level element in message", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"request_opname": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Local name of request element", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"http_method": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("HTTP operation (GET or POST)", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"x_path": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("XPath expression", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"metadata": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Processing metadata", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
	},
}

func (data DmAAAPERBitmap) IsNull() bool {
	if !data.TargetUrl.IsNull() {
		return false
	}
	if !data.OriginalUrl.IsNull() {
		return false
	}
	if !data.RequestUri.IsNull() {
		return false
	}
	if !data.RequestOpname.IsNull() {
		return false
	}
	if !data.HttpMethod.IsNull() {
		return false
	}
	if !data.XPath.IsNull() {
		return false
	}
	if !data.Metadata.IsNull() {
		return false
	}
	return true
}
func GetDmAAAPERBitmapDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmAAAPERBitmapDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmAAAPERBitmapDataSourceSchema
}

func GetDmAAAPERBitmapResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmAAAPERBitmapResourceSchema.Required = true
	} else {
		DmAAAPERBitmapResourceSchema.Optional = true
		DmAAAPERBitmapResourceSchema.Computed = true
	}
	DmAAAPERBitmapResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmAAAPERBitmapResourceSchema
}

func (data DmAAAPERBitmap) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.TargetUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`target-url`, tfutils.StringFromBool(data.TargetUrl, false))
	}
	if !data.OriginalUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`original-url`, tfutils.StringFromBool(data.OriginalUrl, false))
	}
	if !data.RequestUri.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`request-uri`, tfutils.StringFromBool(data.RequestUri, false))
	}
	if !data.RequestOpname.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`request-opname`, tfutils.StringFromBool(data.RequestOpname, false))
	}
	if !data.HttpMethod.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`http-method`, tfutils.StringFromBool(data.HttpMethod, false))
	}
	if !data.XPath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XPath`, tfutils.StringFromBool(data.XPath, false))
	}
	if !data.Metadata.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`metadata`, tfutils.StringFromBool(data.Metadata, false))
	}
	return body
}

func (data *DmAAAPERBitmap) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `target-url`); value.Exists() {
		data.TargetUrl = tfutils.BoolFromString(value.String())
	} else {
		data.TargetUrl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `original-url`); value.Exists() {
		data.OriginalUrl = tfutils.BoolFromString(value.String())
	} else {
		data.OriginalUrl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `request-uri`); value.Exists() {
		data.RequestUri = tfutils.BoolFromString(value.String())
	} else {
		data.RequestUri = types.BoolNull()
	}
	if value := res.Get(pathRoot + `request-opname`); value.Exists() {
		data.RequestOpname = tfutils.BoolFromString(value.String())
	} else {
		data.RequestOpname = types.BoolNull()
	}
	if value := res.Get(pathRoot + `http-method`); value.Exists() {
		data.HttpMethod = tfutils.BoolFromString(value.String())
	} else {
		data.HttpMethod = types.BoolNull()
	}
	if value := res.Get(pathRoot + `XPath`); value.Exists() {
		data.XPath = tfutils.BoolFromString(value.String())
	} else {
		data.XPath = types.BoolNull()
	}
	if value := res.Get(pathRoot + `metadata`); value.Exists() {
		data.Metadata = tfutils.BoolFromString(value.String())
	} else {
		data.Metadata = types.BoolNull()
	}
}

func (data *DmAAAPERBitmap) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `target-url`); value.Exists() && !data.TargetUrl.IsNull() {
		data.TargetUrl = tfutils.BoolFromString(value.String())
	} else if data.TargetUrl.ValueBool() {
		data.TargetUrl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `original-url`); value.Exists() && !data.OriginalUrl.IsNull() {
		data.OriginalUrl = tfutils.BoolFromString(value.String())
	} else if data.OriginalUrl.ValueBool() {
		data.OriginalUrl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `request-uri`); value.Exists() && !data.RequestUri.IsNull() {
		data.RequestUri = tfutils.BoolFromString(value.String())
	} else if data.RequestUri.ValueBool() {
		data.RequestUri = types.BoolNull()
	}
	if value := res.Get(pathRoot + `request-opname`); value.Exists() && !data.RequestOpname.IsNull() {
		data.RequestOpname = tfutils.BoolFromString(value.String())
	} else if data.RequestOpname.ValueBool() {
		data.RequestOpname = types.BoolNull()
	}
	if value := res.Get(pathRoot + `http-method`); value.Exists() && !data.HttpMethod.IsNull() {
		data.HttpMethod = tfutils.BoolFromString(value.String())
	} else if data.HttpMethod.ValueBool() {
		data.HttpMethod = types.BoolNull()
	}
	if value := res.Get(pathRoot + `XPath`); value.Exists() && !data.XPath.IsNull() {
		data.XPath = tfutils.BoolFromString(value.String())
	} else if data.XPath.ValueBool() {
		data.XPath = types.BoolNull()
	}
	if value := res.Get(pathRoot + `metadata`); value.Exists() && !data.Metadata.IsNull() {
		data.Metadata = tfutils.BoolFromString(value.String())
	} else if data.Metadata.ValueBool() {
		data.Metadata = types.BoolNull()
	}
}
