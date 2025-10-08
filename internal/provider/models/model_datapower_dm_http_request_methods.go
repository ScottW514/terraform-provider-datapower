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

type DmHTTPRequestMethods struct {
	Post    types.Bool `tfsdk:"post"`
	Get     types.Bool `tfsdk:"get"`
	Put     types.Bool `tfsdk:"put"`
	Patch   types.Bool `tfsdk:"patch"`
	Head    types.Bool `tfsdk:"head"`
	Options types.Bool `tfsdk:"options"`
	Trace   types.Bool `tfsdk:"trace"`
	Delete  types.Bool `tfsdk:"delete"`
	Connect types.Bool `tfsdk:"connect"`
}

var DmHTTPRequestMethodsObjectType = map[string]attr.Type{
	"post":    types.BoolType,
	"get":     types.BoolType,
	"put":     types.BoolType,
	"patch":   types.BoolType,
	"head":    types.BoolType,
	"options": types.BoolType,
	"trace":   types.BoolType,
	"delete":  types.BoolType,
	"connect": types.BoolType,
}
var DmHTTPRequestMethodsObjectDefault = map[string]attr.Value{
	"post":    types.BoolValue(true),
	"get":     types.BoolValue(true),
	"put":     types.BoolValue(false),
	"patch":   types.BoolValue(false),
	"head":    types.BoolValue(true),
	"options": types.BoolValue(false),
	"trace":   types.BoolValue(false),
	"delete":  types.BoolValue(false),
	"connect": types.BoolValue(false),
}

func GetDmHTTPRequestMethodsDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmHTTPRequestMethodsDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"post": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "POST method",
				Computed:            true,
			},
			"get": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "GET method",
				Computed:            true,
			},
			"put": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "PUT method",
				Computed:            true,
			},
			"patch": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "PATCH method",
				Computed:            true,
			},
			"head": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "HEAD method",
				Computed:            true,
			},
			"options": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "OPTIONS method",
				Computed:            true,
			},
			"trace": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "TRACE method",
				Computed:            true,
			},
			"delete": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "DELETE method",
				Computed:            true,
			},
			"connect": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "CONNECT method",
				Computed:            true,
			},
		},
	}
	DmHTTPRequestMethodsDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmHTTPRequestMethodsDataSourceSchema
}
func GetDmHTTPRequestMethodsResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmHTTPRequestMethodsResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmHTTPRequestMethodsObjectType,
				DmHTTPRequestMethodsObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"post": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("POST method", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"get": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("GET method", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"put": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("PUT method", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"patch": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("PATCH method", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"head": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HEAD method", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"options": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("OPTIONS method", "", "").AddDefaultValue("false").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("CONNECT method", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
	DmHTTPRequestMethodsResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmHTTPRequestMethodsResourceSchema.Required = true
	} else {
		DmHTTPRequestMethodsResourceSchema.Optional = true
		DmHTTPRequestMethodsResourceSchema.Computed = true
	}
	return DmHTTPRequestMethodsResourceSchema
}

func (data DmHTTPRequestMethods) IsNull() bool {
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
	return true
}

func (data DmHTTPRequestMethods) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Post.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`POST`, tfutils.StringFromBool(data.Post, ""))
	}
	if !data.Get.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GET`, tfutils.StringFromBool(data.Get, ""))
	}
	if !data.Put.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PUT`, tfutils.StringFromBool(data.Put, ""))
	}
	if !data.Patch.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PATCH`, tfutils.StringFromBool(data.Patch, ""))
	}
	if !data.Head.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HEAD`, tfutils.StringFromBool(data.Head, ""))
	}
	if !data.Options.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OPTIONS`, tfutils.StringFromBool(data.Options, ""))
	}
	if !data.Trace.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TRACE`, tfutils.StringFromBool(data.Trace, ""))
	}
	if !data.Delete.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DELETE`, tfutils.StringFromBool(data.Delete, ""))
	}
	if !data.Connect.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CONNECT`, tfutils.StringFromBool(data.Connect, ""))
	}
	return body
}

func (data *DmHTTPRequestMethods) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
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
}

func (data *DmHTTPRequestMethods) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `POST`); value.Exists() && !data.Post.IsNull() {
		data.Post = tfutils.BoolFromString(value.String())
	} else if !data.Post.ValueBool() {
		data.Post = types.BoolNull()
	}
	if value := res.Get(pathRoot + `GET`); value.Exists() && !data.Get.IsNull() {
		data.Get = tfutils.BoolFromString(value.String())
	} else if !data.Get.ValueBool() {
		data.Get = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PUT`); value.Exists() && !data.Put.IsNull() {
		data.Put = tfutils.BoolFromString(value.String())
	} else if data.Put.ValueBool() {
		data.Put = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PATCH`); value.Exists() && !data.Patch.IsNull() {
		data.Patch = tfutils.BoolFromString(value.String())
	} else if data.Patch.ValueBool() {
		data.Patch = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HEAD`); value.Exists() && !data.Head.IsNull() {
		data.Head = tfutils.BoolFromString(value.String())
	} else if !data.Head.ValueBool() {
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
}
