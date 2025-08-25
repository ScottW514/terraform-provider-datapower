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

type DmHTTPResponseCodes struct {
	Http100 types.Bool `tfsdk:"http_100"`
	Http101 types.Bool `tfsdk:"http_101"`
	Http200 types.Bool `tfsdk:"http_200"`
	Http201 types.Bool `tfsdk:"http_201"`
	Http202 types.Bool `tfsdk:"http_202"`
	Http203 types.Bool `tfsdk:"http_203"`
	Http204 types.Bool `tfsdk:"http_204"`
	Http205 types.Bool `tfsdk:"http_205"`
	Http206 types.Bool `tfsdk:"http_206"`
	Http300 types.Bool `tfsdk:"http_300"`
	Http301 types.Bool `tfsdk:"http_301"`
	Http302 types.Bool `tfsdk:"http_302"`
	Http303 types.Bool `tfsdk:"http_303"`
	Http304 types.Bool `tfsdk:"http_304"`
	Http305 types.Bool `tfsdk:"http_305"`
	Http307 types.Bool `tfsdk:"http_307"`
	Http400 types.Bool `tfsdk:"http_400"`
	Http401 types.Bool `tfsdk:"http_401"`
	Http402 types.Bool `tfsdk:"http_402"`
	Http403 types.Bool `tfsdk:"http_403"`
	Http404 types.Bool `tfsdk:"http_404"`
	Http405 types.Bool `tfsdk:"http_405"`
	Http406 types.Bool `tfsdk:"http_406"`
	Http407 types.Bool `tfsdk:"http_407"`
	Http408 types.Bool `tfsdk:"http_408"`
	Http409 types.Bool `tfsdk:"http_409"`
	Http410 types.Bool `tfsdk:"http_410"`
	Http411 types.Bool `tfsdk:"http_411"`
	Http412 types.Bool `tfsdk:"http_412"`
	Http413 types.Bool `tfsdk:"http_413"`
	Http500 types.Bool `tfsdk:"http_500"`
	Http503 types.Bool `tfsdk:"http_503"`
}

var DmHTTPResponseCodesObjectType = map[string]attr.Type{
	"http_100": types.BoolType,
	"http_101": types.BoolType,
	"http_200": types.BoolType,
	"http_201": types.BoolType,
	"http_202": types.BoolType,
	"http_203": types.BoolType,
	"http_204": types.BoolType,
	"http_205": types.BoolType,
	"http_206": types.BoolType,
	"http_300": types.BoolType,
	"http_301": types.BoolType,
	"http_302": types.BoolType,
	"http_303": types.BoolType,
	"http_304": types.BoolType,
	"http_305": types.BoolType,
	"http_307": types.BoolType,
	"http_400": types.BoolType,
	"http_401": types.BoolType,
	"http_402": types.BoolType,
	"http_403": types.BoolType,
	"http_404": types.BoolType,
	"http_405": types.BoolType,
	"http_406": types.BoolType,
	"http_407": types.BoolType,
	"http_408": types.BoolType,
	"http_409": types.BoolType,
	"http_410": types.BoolType,
	"http_411": types.BoolType,
	"http_412": types.BoolType,
	"http_413": types.BoolType,
	"http_500": types.BoolType,
	"http_503": types.BoolType,
}
var DmHTTPResponseCodesObjectDefault = map[string]attr.Value{
	"http_100": types.BoolValue(true),
	"http_101": types.BoolValue(false),
	"http_200": types.BoolValue(true),
	"http_201": types.BoolValue(true),
	"http_202": types.BoolValue(true),
	"http_203": types.BoolValue(false),
	"http_204": types.BoolValue(true),
	"http_205": types.BoolValue(false),
	"http_206": types.BoolValue(true),
	"http_300": types.BoolValue(false),
	"http_301": types.BoolValue(true),
	"http_302": types.BoolValue(true),
	"http_303": types.BoolValue(false),
	"http_304": types.BoolValue(true),
	"http_305": types.BoolValue(false),
	"http_307": types.BoolValue(true),
	"http_400": types.BoolValue(false),
	"http_401": types.BoolValue(false),
	"http_402": types.BoolValue(false),
	"http_403": types.BoolValue(false),
	"http_404": types.BoolValue(false),
	"http_405": types.BoolValue(false),
	"http_406": types.BoolValue(false),
	"http_407": types.BoolValue(false),
	"http_408": types.BoolValue(false),
	"http_409": types.BoolValue(false),
	"http_410": types.BoolValue(false),
	"http_411": types.BoolValue(false),
	"http_412": types.BoolValue(false),
	"http_413": types.BoolValue(false),
	"http_500": types.BoolValue(false),
	"http_503": types.BoolValue(false),
}

func GetDmHTTPResponseCodesDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmHTTPResponseCodesDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"http_100": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("100 Continue", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"http_101": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("101 Switching Protocols", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"http_200": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("200 OK", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"http_201": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("201 Created", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"http_202": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("202 Accepted", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"http_203": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("203 Non-Authoritative Information", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"http_204": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("204 No Content", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"http_205": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("205 Reset Content", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"http_206": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("206 Partial Content", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"http_300": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("300 Multiple Choices", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"http_301": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("301 Moved", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"http_302": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("302 Found", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"http_303": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("303 See Other", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"http_304": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("304 Not Modified", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"http_305": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("305 Use Proxy", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"http_307": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("307 Temporary Redirect", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"http_400": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("400 Bad Request", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"http_401": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("401 Unauthorized", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"http_402": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("402 Payment Required", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"http_403": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("403 Forbidden", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"http_404": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("404 Not Found", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"http_405": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("405 Method Not Allowed", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"http_406": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("406 Not Acceptable", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"http_407": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("407 Proxy Authentication Required", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"http_408": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("408 Request Timeout", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"http_409": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("409 Conflict", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"http_410": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("410 Gone", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"http_411": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("411 Length required", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"http_412": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("412 Precondition Failed", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"http_413": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("413 Request Entity Too Large", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"http_500": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("500 Server Error", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"http_503": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("503 Service Unavailable", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
		},
	}
	DmHTTPResponseCodesDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmHTTPResponseCodesDataSourceSchema
}
func GetDmHTTPResponseCodesResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmHTTPResponseCodesResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmHTTPResponseCodesObjectType,
				DmHTTPResponseCodesObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"http_100": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("100 Continue", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"http_101": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("101 Switching Protocols", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_200": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("200 OK", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"http_201": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("201 Created", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"http_202": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("202 Accepted", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"http_203": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("203 Non-Authoritative Information", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_204": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("204 No Content", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"http_205": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("205 Reset Content", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_206": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("206 Partial Content", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"http_300": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("300 Multiple Choices", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_301": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("301 Moved", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"http_302": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("302 Found", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"http_303": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("303 See Other", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_304": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("304 Not Modified", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"http_305": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("305 Use Proxy", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_307": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("307 Temporary Redirect", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"http_400": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("400 Bad Request", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_401": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("401 Unauthorized", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_402": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("402 Payment Required", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_403": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("403 Forbidden", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_404": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("404 Not Found", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_405": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("405 Method Not Allowed", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_406": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("406 Not Acceptable", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_407": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("407 Proxy Authentication Required", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_408": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("408 Request Timeout", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_409": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("409 Conflict", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_410": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("410 Gone", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_411": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("411 Length required", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_412": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("412 Precondition Failed", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_413": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("413 Request Entity Too Large", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_500": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("500 Server Error", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_503": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("503 Service Unavailable", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
	DmHTTPResponseCodesResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmHTTPResponseCodesResourceSchema.Required = true
	} else {
		DmHTTPResponseCodesResourceSchema.Optional = true
		DmHTTPResponseCodesResourceSchema.Computed = true
	}
	return DmHTTPResponseCodesResourceSchema
}

func (data DmHTTPResponseCodes) IsNull() bool {
	if !data.Http100.IsNull() {
		return false
	}
	if !data.Http101.IsNull() {
		return false
	}
	if !data.Http200.IsNull() {
		return false
	}
	if !data.Http201.IsNull() {
		return false
	}
	if !data.Http202.IsNull() {
		return false
	}
	if !data.Http203.IsNull() {
		return false
	}
	if !data.Http204.IsNull() {
		return false
	}
	if !data.Http205.IsNull() {
		return false
	}
	if !data.Http206.IsNull() {
		return false
	}
	if !data.Http300.IsNull() {
		return false
	}
	if !data.Http301.IsNull() {
		return false
	}
	if !data.Http302.IsNull() {
		return false
	}
	if !data.Http303.IsNull() {
		return false
	}
	if !data.Http304.IsNull() {
		return false
	}
	if !data.Http305.IsNull() {
		return false
	}
	if !data.Http307.IsNull() {
		return false
	}
	if !data.Http400.IsNull() {
		return false
	}
	if !data.Http401.IsNull() {
		return false
	}
	if !data.Http402.IsNull() {
		return false
	}
	if !data.Http403.IsNull() {
		return false
	}
	if !data.Http404.IsNull() {
		return false
	}
	if !data.Http405.IsNull() {
		return false
	}
	if !data.Http406.IsNull() {
		return false
	}
	if !data.Http407.IsNull() {
		return false
	}
	if !data.Http408.IsNull() {
		return false
	}
	if !data.Http409.IsNull() {
		return false
	}
	if !data.Http410.IsNull() {
		return false
	}
	if !data.Http411.IsNull() {
		return false
	}
	if !data.Http412.IsNull() {
		return false
	}
	if !data.Http413.IsNull() {
		return false
	}
	if !data.Http500.IsNull() {
		return false
	}
	if !data.Http503.IsNull() {
		return false
	}
	return true
}

func (data DmHTTPResponseCodes) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Http100.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-100`, tfutils.StringFromBool(data.Http100, ""))
	}
	if !data.Http101.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-101`, tfutils.StringFromBool(data.Http101, ""))
	}
	if !data.Http200.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-200`, tfutils.StringFromBool(data.Http200, ""))
	}
	if !data.Http201.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-201`, tfutils.StringFromBool(data.Http201, ""))
	}
	if !data.Http202.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-202`, tfutils.StringFromBool(data.Http202, ""))
	}
	if !data.Http203.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-203`, tfutils.StringFromBool(data.Http203, ""))
	}
	if !data.Http204.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-204`, tfutils.StringFromBool(data.Http204, ""))
	}
	if !data.Http205.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-205`, tfutils.StringFromBool(data.Http205, ""))
	}
	if !data.Http206.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-206`, tfutils.StringFromBool(data.Http206, ""))
	}
	if !data.Http300.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-300`, tfutils.StringFromBool(data.Http300, ""))
	}
	if !data.Http301.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-301`, tfutils.StringFromBool(data.Http301, ""))
	}
	if !data.Http302.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-302`, tfutils.StringFromBool(data.Http302, ""))
	}
	if !data.Http303.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-303`, tfutils.StringFromBool(data.Http303, ""))
	}
	if !data.Http304.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-304`, tfutils.StringFromBool(data.Http304, ""))
	}
	if !data.Http305.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-305`, tfutils.StringFromBool(data.Http305, ""))
	}
	if !data.Http307.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-307`, tfutils.StringFromBool(data.Http307, ""))
	}
	if !data.Http400.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-400`, tfutils.StringFromBool(data.Http400, ""))
	}
	if !data.Http401.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-401`, tfutils.StringFromBool(data.Http401, ""))
	}
	if !data.Http402.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-402`, tfutils.StringFromBool(data.Http402, ""))
	}
	if !data.Http403.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-403`, tfutils.StringFromBool(data.Http403, ""))
	}
	if !data.Http404.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-404`, tfutils.StringFromBool(data.Http404, ""))
	}
	if !data.Http405.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-405`, tfutils.StringFromBool(data.Http405, ""))
	}
	if !data.Http406.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-406`, tfutils.StringFromBool(data.Http406, ""))
	}
	if !data.Http407.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-407`, tfutils.StringFromBool(data.Http407, ""))
	}
	if !data.Http408.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-408`, tfutils.StringFromBool(data.Http408, ""))
	}
	if !data.Http409.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-409`, tfutils.StringFromBool(data.Http409, ""))
	}
	if !data.Http410.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-410`, tfutils.StringFromBool(data.Http410, ""))
	}
	if !data.Http411.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-411`, tfutils.StringFromBool(data.Http411, ""))
	}
	if !data.Http412.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-412`, tfutils.StringFromBool(data.Http412, ""))
	}
	if !data.Http413.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-413`, tfutils.StringFromBool(data.Http413, ""))
	}
	if !data.Http500.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-500`, tfutils.StringFromBool(data.Http500, ""))
	}
	if !data.Http503.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP-503`, tfutils.StringFromBool(data.Http503, ""))
	}
	return body
}

func (data *DmHTTPResponseCodes) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `HTTP-100`); value.Exists() {
		data.Http100 = tfutils.BoolFromString(value.String())
	} else {
		data.Http100 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-101`); value.Exists() {
		data.Http101 = tfutils.BoolFromString(value.String())
	} else {
		data.Http101 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-200`); value.Exists() {
		data.Http200 = tfutils.BoolFromString(value.String())
	} else {
		data.Http200 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-201`); value.Exists() {
		data.Http201 = tfutils.BoolFromString(value.String())
	} else {
		data.Http201 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-202`); value.Exists() {
		data.Http202 = tfutils.BoolFromString(value.String())
	} else {
		data.Http202 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-203`); value.Exists() {
		data.Http203 = tfutils.BoolFromString(value.String())
	} else {
		data.Http203 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-204`); value.Exists() {
		data.Http204 = tfutils.BoolFromString(value.String())
	} else {
		data.Http204 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-205`); value.Exists() {
		data.Http205 = tfutils.BoolFromString(value.String())
	} else {
		data.Http205 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-206`); value.Exists() {
		data.Http206 = tfutils.BoolFromString(value.String())
	} else {
		data.Http206 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-300`); value.Exists() {
		data.Http300 = tfutils.BoolFromString(value.String())
	} else {
		data.Http300 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-301`); value.Exists() {
		data.Http301 = tfutils.BoolFromString(value.String())
	} else {
		data.Http301 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-302`); value.Exists() {
		data.Http302 = tfutils.BoolFromString(value.String())
	} else {
		data.Http302 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-303`); value.Exists() {
		data.Http303 = tfutils.BoolFromString(value.String())
	} else {
		data.Http303 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-304`); value.Exists() {
		data.Http304 = tfutils.BoolFromString(value.String())
	} else {
		data.Http304 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-305`); value.Exists() {
		data.Http305 = tfutils.BoolFromString(value.String())
	} else {
		data.Http305 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-307`); value.Exists() {
		data.Http307 = tfutils.BoolFromString(value.String())
	} else {
		data.Http307 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-400`); value.Exists() {
		data.Http400 = tfutils.BoolFromString(value.String())
	} else {
		data.Http400 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-401`); value.Exists() {
		data.Http401 = tfutils.BoolFromString(value.String())
	} else {
		data.Http401 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-402`); value.Exists() {
		data.Http402 = tfutils.BoolFromString(value.String())
	} else {
		data.Http402 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-403`); value.Exists() {
		data.Http403 = tfutils.BoolFromString(value.String())
	} else {
		data.Http403 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-404`); value.Exists() {
		data.Http404 = tfutils.BoolFromString(value.String())
	} else {
		data.Http404 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-405`); value.Exists() {
		data.Http405 = tfutils.BoolFromString(value.String())
	} else {
		data.Http405 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-406`); value.Exists() {
		data.Http406 = tfutils.BoolFromString(value.String())
	} else {
		data.Http406 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-407`); value.Exists() {
		data.Http407 = tfutils.BoolFromString(value.String())
	} else {
		data.Http407 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-408`); value.Exists() {
		data.Http408 = tfutils.BoolFromString(value.String())
	} else {
		data.Http408 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-409`); value.Exists() {
		data.Http409 = tfutils.BoolFromString(value.String())
	} else {
		data.Http409 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-410`); value.Exists() {
		data.Http410 = tfutils.BoolFromString(value.String())
	} else {
		data.Http410 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-411`); value.Exists() {
		data.Http411 = tfutils.BoolFromString(value.String())
	} else {
		data.Http411 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-412`); value.Exists() {
		data.Http412 = tfutils.BoolFromString(value.String())
	} else {
		data.Http412 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-413`); value.Exists() {
		data.Http413 = tfutils.BoolFromString(value.String())
	} else {
		data.Http413 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-500`); value.Exists() {
		data.Http500 = tfutils.BoolFromString(value.String())
	} else {
		data.Http500 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-503`); value.Exists() {
		data.Http503 = tfutils.BoolFromString(value.String())
	} else {
		data.Http503 = types.BoolNull()
	}
}

func (data *DmHTTPResponseCodes) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `HTTP-100`); value.Exists() && !data.Http100.IsNull() {
		data.Http100 = tfutils.BoolFromString(value.String())
	} else if !data.Http100.ValueBool() {
		data.Http100 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-101`); value.Exists() && !data.Http101.IsNull() {
		data.Http101 = tfutils.BoolFromString(value.String())
	} else if data.Http101.ValueBool() {
		data.Http101 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-200`); value.Exists() && !data.Http200.IsNull() {
		data.Http200 = tfutils.BoolFromString(value.String())
	} else if !data.Http200.ValueBool() {
		data.Http200 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-201`); value.Exists() && !data.Http201.IsNull() {
		data.Http201 = tfutils.BoolFromString(value.String())
	} else if !data.Http201.ValueBool() {
		data.Http201 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-202`); value.Exists() && !data.Http202.IsNull() {
		data.Http202 = tfutils.BoolFromString(value.String())
	} else if !data.Http202.ValueBool() {
		data.Http202 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-203`); value.Exists() && !data.Http203.IsNull() {
		data.Http203 = tfutils.BoolFromString(value.String())
	} else if data.Http203.ValueBool() {
		data.Http203 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-204`); value.Exists() && !data.Http204.IsNull() {
		data.Http204 = tfutils.BoolFromString(value.String())
	} else if !data.Http204.ValueBool() {
		data.Http204 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-205`); value.Exists() && !data.Http205.IsNull() {
		data.Http205 = tfutils.BoolFromString(value.String())
	} else if data.Http205.ValueBool() {
		data.Http205 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-206`); value.Exists() && !data.Http206.IsNull() {
		data.Http206 = tfutils.BoolFromString(value.String())
	} else if !data.Http206.ValueBool() {
		data.Http206 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-300`); value.Exists() && !data.Http300.IsNull() {
		data.Http300 = tfutils.BoolFromString(value.String())
	} else if data.Http300.ValueBool() {
		data.Http300 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-301`); value.Exists() && !data.Http301.IsNull() {
		data.Http301 = tfutils.BoolFromString(value.String())
	} else if !data.Http301.ValueBool() {
		data.Http301 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-302`); value.Exists() && !data.Http302.IsNull() {
		data.Http302 = tfutils.BoolFromString(value.String())
	} else if !data.Http302.ValueBool() {
		data.Http302 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-303`); value.Exists() && !data.Http303.IsNull() {
		data.Http303 = tfutils.BoolFromString(value.String())
	} else if data.Http303.ValueBool() {
		data.Http303 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-304`); value.Exists() && !data.Http304.IsNull() {
		data.Http304 = tfutils.BoolFromString(value.String())
	} else if !data.Http304.ValueBool() {
		data.Http304 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-305`); value.Exists() && !data.Http305.IsNull() {
		data.Http305 = tfutils.BoolFromString(value.String())
	} else if data.Http305.ValueBool() {
		data.Http305 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-307`); value.Exists() && !data.Http307.IsNull() {
		data.Http307 = tfutils.BoolFromString(value.String())
	} else if !data.Http307.ValueBool() {
		data.Http307 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-400`); value.Exists() && !data.Http400.IsNull() {
		data.Http400 = tfutils.BoolFromString(value.String())
	} else if data.Http400.ValueBool() {
		data.Http400 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-401`); value.Exists() && !data.Http401.IsNull() {
		data.Http401 = tfutils.BoolFromString(value.String())
	} else if data.Http401.ValueBool() {
		data.Http401 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-402`); value.Exists() && !data.Http402.IsNull() {
		data.Http402 = tfutils.BoolFromString(value.String())
	} else if data.Http402.ValueBool() {
		data.Http402 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-403`); value.Exists() && !data.Http403.IsNull() {
		data.Http403 = tfutils.BoolFromString(value.String())
	} else if data.Http403.ValueBool() {
		data.Http403 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-404`); value.Exists() && !data.Http404.IsNull() {
		data.Http404 = tfutils.BoolFromString(value.String())
	} else if data.Http404.ValueBool() {
		data.Http404 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-405`); value.Exists() && !data.Http405.IsNull() {
		data.Http405 = tfutils.BoolFromString(value.String())
	} else if data.Http405.ValueBool() {
		data.Http405 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-406`); value.Exists() && !data.Http406.IsNull() {
		data.Http406 = tfutils.BoolFromString(value.String())
	} else if data.Http406.ValueBool() {
		data.Http406 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-407`); value.Exists() && !data.Http407.IsNull() {
		data.Http407 = tfutils.BoolFromString(value.String())
	} else if data.Http407.ValueBool() {
		data.Http407 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-408`); value.Exists() && !data.Http408.IsNull() {
		data.Http408 = tfutils.BoolFromString(value.String())
	} else if data.Http408.ValueBool() {
		data.Http408 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-409`); value.Exists() && !data.Http409.IsNull() {
		data.Http409 = tfutils.BoolFromString(value.String())
	} else if data.Http409.ValueBool() {
		data.Http409 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-410`); value.Exists() && !data.Http410.IsNull() {
		data.Http410 = tfutils.BoolFromString(value.String())
	} else if data.Http410.ValueBool() {
		data.Http410 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-411`); value.Exists() && !data.Http411.IsNull() {
		data.Http411 = tfutils.BoolFromString(value.String())
	} else if data.Http411.ValueBool() {
		data.Http411 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-412`); value.Exists() && !data.Http412.IsNull() {
		data.Http412 = tfutils.BoolFromString(value.String())
	} else if data.Http412.ValueBool() {
		data.Http412 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-413`); value.Exists() && !data.Http413.IsNull() {
		data.Http413 = tfutils.BoolFromString(value.String())
	} else if data.Http413.ValueBool() {
		data.Http413 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-500`); value.Exists() && !data.Http500.IsNull() {
		data.Http500 = tfutils.BoolFromString(value.String())
	} else if data.Http500.ValueBool() {
		data.Http500 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTP-503`); value.Exists() && !data.Http503.IsNull() {
		data.Http503 = tfutils.BoolFromString(value.String())
	} else if data.Http503.ValueBool() {
		data.Http503 = types.BoolNull()
	}
}
