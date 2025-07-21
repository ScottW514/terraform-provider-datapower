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

type DmSourceHTTPFeatureType struct {
	Http1D0             types.Bool `tfsdk:"http_1d0"`
	Http1D1             types.Bool `tfsdk:"http_1d1"`
	Http2D0             types.Bool `tfsdk:"http_2d0"`
	Post                types.Bool `tfsdk:"post"`
	Get                 types.Bool `tfsdk:"get"`
	Put                 types.Bool `tfsdk:"put"`
	Patch               types.Bool `tfsdk:"patch"`
	Head                types.Bool `tfsdk:"head"`
	Options             types.Bool `tfsdk:"options"`
	Trace               types.Bool `tfsdk:"trace"`
	Delete              types.Bool `tfsdk:"delete"`
	Connect             types.Bool `tfsdk:"connect"`
	CustomMethods       types.Bool `tfsdk:"custom_methods"`
	QueryString         types.Bool `tfsdk:"query_string"`
	FragmentIdentifiers types.Bool `tfsdk:"fragment_identifiers"`
	DotDot              types.Bool `tfsdk:"dot_dot"`
	DotDotInPath        types.Bool `tfsdk:"dot_dot_in_path"`
	DotDotInQueryString types.Bool `tfsdk:"dot_dot_in_query_string"`
	CmdExe              types.Bool `tfsdk:"cmd_exe"`
}

var DmSourceHTTPFeatureTypeObjectType = map[string]attr.Type{
	"http_1d0":                types.BoolType,
	"http_1d1":                types.BoolType,
	"http_2d0":                types.BoolType,
	"post":                    types.BoolType,
	"get":                     types.BoolType,
	"put":                     types.BoolType,
	"patch":                   types.BoolType,
	"head":                    types.BoolType,
	"options":                 types.BoolType,
	"trace":                   types.BoolType,
	"delete":                  types.BoolType,
	"connect":                 types.BoolType,
	"custom_methods":          types.BoolType,
	"query_string":            types.BoolType,
	"fragment_identifiers":    types.BoolType,
	"dot_dot":                 types.BoolType,
	"dot_dot_in_path":         types.BoolType,
	"dot_dot_in_query_string": types.BoolType,
	"cmd_exe":                 types.BoolType,
}
var DmSourceHTTPFeatureTypeObjectDefault = map[string]attr.Value{
	"http_1d0":                types.BoolValue(true),
	"http_1d1":                types.BoolValue(true),
	"http_2d0":                types.BoolValue(false),
	"post":                    types.BoolValue(true),
	"get":                     types.BoolValue(false),
	"put":                     types.BoolValue(true),
	"patch":                   types.BoolValue(false),
	"head":                    types.BoolValue(false),
	"options":                 types.BoolValue(false),
	"trace":                   types.BoolValue(false),
	"delete":                  types.BoolValue(false),
	"connect":                 types.BoolValue(false),
	"custom_methods":          types.BoolValue(false),
	"query_string":            types.BoolValue(true),
	"fragment_identifiers":    types.BoolValue(true),
	"dot_dot":                 types.BoolValue(false),
	"dot_dot_in_path":         types.BoolValue(false),
	"dot_dot_in_query_string": types.BoolValue(false),
	"cmd_exe":                 types.BoolValue(false),
}
var DmSourceHTTPFeatureTypeDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"http_1d0": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("HTTP/1.0", "", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"http_1d1": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("HTTP/1.1", "", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"http_2d0": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("HTTP/2", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"post": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("POST method", "", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"get": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("GET method", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"put": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("PUT method", "", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"patch": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("PATCH method", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"head": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("HEAD method", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"options": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("OPTIONS", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"trace": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("TRACE method", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"delete": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("DELETE method", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"connect": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"custom_methods": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Custom methods", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"query_string": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL with ?", "", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"fragment_identifiers": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL with #", "", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"dot_dot": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL with ..", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"dot_dot_in_path": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL with .. in path", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"dot_dot_in_query_string": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL with .. after ?", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"cmd_exe": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL with cmd.exe", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
	},
}
var DmSourceHTTPFeatureTypeResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmSourceHTTPFeatureTypeObjectType,
			DmSourceHTTPFeatureTypeObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"http_1d0": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("HTTP/1.0", "", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"http_1d1": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("HTTP/1.1", "", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"http_2d0": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("HTTP/2", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"post": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("POST method", "", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"get": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("GET method", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"put": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("PUT method", "", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"patch": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("PATCH method", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"head": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("HEAD method", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"options": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("OPTIONS", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"trace": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("TRACE method", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"delete": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("DELETE method", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"connect": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"custom_methods": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Custom methods", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"query_string": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL with ?", "", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"fragment_identifiers": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL with #", "", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"dot_dot": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL with ..", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"dot_dot_in_path": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL with .. in path", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"dot_dot_in_query_string": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL with .. after ?", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"cmd_exe": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL with cmd.exe", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
	},
}

func (data DmSourceHTTPFeatureType) IsNull() bool {
	if !data.Http1D0.IsNull() {
		return false
	}
	if !data.Http1D1.IsNull() {
		return false
	}
	if !data.Http2D0.IsNull() {
		return false
	}
	if !data.Post.IsNull() {
		return false
	}
	if !data.Get.IsNull() {
		return false
	}
	if !data.Put.IsNull() {
		return false
	}
	if !data.Patch.IsNull() {
		return false
	}
	if !data.Head.IsNull() {
		return false
	}
	if !data.Options.IsNull() {
		return false
	}
	if !data.Trace.IsNull() {
		return false
	}
	if !data.Delete.IsNull() {
		return false
	}
	if !data.Connect.IsNull() {
		return false
	}
	if !data.CustomMethods.IsNull() {
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
	if !data.DotDotInPath.IsNull() {
		return false
	}
	if !data.DotDotInQueryString.IsNull() {
		return false
	}
	if !data.CmdExe.IsNull() {
		return false
	}
	return true
}
func GetDmSourceHTTPFeatureTypeDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmSourceHTTPFeatureTypeDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmSourceHTTPFeatureTypeDataSourceSchema
}

func GetDmSourceHTTPFeatureTypeResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmSourceHTTPFeatureTypeResourceSchema.Required = true
	} else {
		DmSourceHTTPFeatureTypeResourceSchema.Optional = true
		DmSourceHTTPFeatureTypeResourceSchema.Computed = true
	}
	DmSourceHTTPFeatureTypeResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmSourceHTTPFeatureTypeResourceSchema
}

func (data DmSourceHTTPFeatureType) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Http1D0.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-1\.0`, tfutils.StringFromBool(data.Http1D0))
	}
	if !data.Http1D1.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-1\.1`, tfutils.StringFromBool(data.Http1D1))
	}
	if !data.Http2D0.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-2\.0`, tfutils.StringFromBool(data.Http2D0))
	}
	if !data.Post.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`POST`, tfutils.StringFromBool(data.Post))
	}
	if !data.Get.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GET`, tfutils.StringFromBool(data.Get))
	}
	if !data.Put.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PUT`, tfutils.StringFromBool(data.Put))
	}
	if !data.Patch.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PATCH`, tfutils.StringFromBool(data.Patch))
	}
	if !data.Head.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HEAD`, tfutils.StringFromBool(data.Head))
	}
	if !data.Options.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OPTIONS`, tfutils.StringFromBool(data.Options))
	}
	if !data.Trace.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TRACE`, tfutils.StringFromBool(data.Trace))
	}
	if !data.Delete.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DELETE`, tfutils.StringFromBool(data.Delete))
	}
	if !data.Connect.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CONNECT`, tfutils.StringFromBool(data.Connect))
	}
	if !data.CustomMethods.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CustomMethods`, tfutils.StringFromBool(data.CustomMethods))
	}
	if !data.QueryString.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`QueryString`, tfutils.StringFromBool(data.QueryString))
	}
	if !data.FragmentIdentifiers.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FragmentIdentifiers`, tfutils.StringFromBool(data.FragmentIdentifiers))
	}
	if !data.DotDot.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DotDot`, tfutils.StringFromBool(data.DotDot))
	}
	if !data.DotDotInPath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DotDotInPath`, tfutils.StringFromBool(data.DotDotInPath))
	}
	if !data.DotDotInQueryString.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DotDotInQueryString`, tfutils.StringFromBool(data.DotDotInQueryString))
	}
	if !data.CmdExe.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CmdExe`, tfutils.StringFromBool(data.CmdExe))
	}
	return body
}

func (data *DmSourceHTTPFeatureType) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `HTTP-2\.0`); value.Exists() {
		data.Http2D0 = tfutils.BoolFromString(value.String())
	} else {
		data.Http2D0 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `POST`); value.Exists() {
		data.Post = tfutils.BoolFromString(value.String())
	} else {
		data.Post = types.BoolNull()
	}
	if value := res.Get(pathRoot + `GET`); value.Exists() {
		data.Get = tfutils.BoolFromString(value.String())
	} else {
		data.Get = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PUT`); value.Exists() {
		data.Put = tfutils.BoolFromString(value.String())
	} else {
		data.Put = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PATCH`); value.Exists() {
		data.Patch = tfutils.BoolFromString(value.String())
	} else {
		data.Patch = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HEAD`); value.Exists() {
		data.Head = tfutils.BoolFromString(value.String())
	} else {
		data.Head = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OPTIONS`); value.Exists() {
		data.Options = tfutils.BoolFromString(value.String())
	} else {
		data.Options = types.BoolNull()
	}
	if value := res.Get(pathRoot + `TRACE`); value.Exists() {
		data.Trace = tfutils.BoolFromString(value.String())
	} else {
		data.Trace = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DELETE`); value.Exists() {
		data.Delete = tfutils.BoolFromString(value.String())
	} else {
		data.Delete = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CONNECT`); value.Exists() {
		data.Connect = tfutils.BoolFromString(value.String())
	} else {
		data.Connect = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CustomMethods`); value.Exists() {
		data.CustomMethods = tfutils.BoolFromString(value.String())
	} else {
		data.CustomMethods = types.BoolNull()
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
	if value := res.Get(pathRoot + `DotDotInPath`); value.Exists() {
		data.DotDotInPath = tfutils.BoolFromString(value.String())
	} else {
		data.DotDotInPath = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DotDotInQueryString`); value.Exists() {
		data.DotDotInQueryString = tfutils.BoolFromString(value.String())
	} else {
		data.DotDotInQueryString = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CmdExe`); value.Exists() {
		data.CmdExe = tfutils.BoolFromString(value.String())
	} else {
		data.CmdExe = types.BoolNull()
	}
}

func (data *DmSourceHTTPFeatureType) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `HTTP-2\.0`); value.Exists() && !data.Http2D0.IsNull() {
		data.Http2D0 = tfutils.BoolFromString(value.String())
	} else if data.Http2D0.ValueBool() {
		data.Http2D0 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `POST`); value.Exists() && !data.Post.IsNull() {
		data.Post = tfutils.BoolFromString(value.String())
	} else if !data.Post.ValueBool() {
		data.Post = types.BoolNull()
	}
	if value := res.Get(pathRoot + `GET`); value.Exists() && !data.Get.IsNull() {
		data.Get = tfutils.BoolFromString(value.String())
	} else if data.Get.ValueBool() {
		data.Get = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PUT`); value.Exists() && !data.Put.IsNull() {
		data.Put = tfutils.BoolFromString(value.String())
	} else if !data.Put.ValueBool() {
		data.Put = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PATCH`); value.Exists() && !data.Patch.IsNull() {
		data.Patch = tfutils.BoolFromString(value.String())
	} else if data.Patch.ValueBool() {
		data.Patch = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HEAD`); value.Exists() && !data.Head.IsNull() {
		data.Head = tfutils.BoolFromString(value.String())
	} else if data.Head.ValueBool() {
		data.Head = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OPTIONS`); value.Exists() && !data.Options.IsNull() {
		data.Options = tfutils.BoolFromString(value.String())
	} else if data.Options.ValueBool() {
		data.Options = types.BoolNull()
	}
	if value := res.Get(pathRoot + `TRACE`); value.Exists() && !data.Trace.IsNull() {
		data.Trace = tfutils.BoolFromString(value.String())
	} else if data.Trace.ValueBool() {
		data.Trace = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DELETE`); value.Exists() && !data.Delete.IsNull() {
		data.Delete = tfutils.BoolFromString(value.String())
	} else if data.Delete.ValueBool() {
		data.Delete = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CONNECT`); value.Exists() && !data.Connect.IsNull() {
		data.Connect = tfutils.BoolFromString(value.String())
	} else if data.Connect.ValueBool() {
		data.Connect = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CustomMethods`); value.Exists() && !data.CustomMethods.IsNull() {
		data.CustomMethods = tfutils.BoolFromString(value.String())
	} else if data.CustomMethods.ValueBool() {
		data.CustomMethods = types.BoolNull()
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
	if value := res.Get(pathRoot + `DotDotInPath`); value.Exists() && !data.DotDotInPath.IsNull() {
		data.DotDotInPath = tfutils.BoolFromString(value.String())
	} else if data.DotDotInPath.ValueBool() {
		data.DotDotInPath = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DotDotInQueryString`); value.Exists() && !data.DotDotInQueryString.IsNull() {
		data.DotDotInQueryString = tfutils.BoolFromString(value.String())
	} else if data.DotDotInQueryString.ValueBool() {
		data.DotDotInQueryString = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CmdExe`); value.Exists() && !data.CmdExe.IsNull() {
		data.CmdExe = tfutils.BoolFromString(value.String())
	} else if data.CmdExe.ValueBool() {
		data.CmdExe = types.BoolNull()
	}
}
