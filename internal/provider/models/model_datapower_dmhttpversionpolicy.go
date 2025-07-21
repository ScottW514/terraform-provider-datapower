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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmHTTPVersionPolicy struct {
	RegExp        types.String `tfsdk:"reg_exp"`
	Version       types.String `tfsdk:"version"`
	Http2Required types.Bool   `tfsdk:"http2_required"`
}

var DmHTTPVersionPolicyObjectType = map[string]attr.Type{
	"reg_exp":        types.StringType,
	"version":        types.StringType,
	"http2_required": types.BoolType,
}
var DmHTTPVersionPolicyObjectDefault = map[string]attr.Value{
	"reg_exp":        types.StringNull(),
	"version":        types.StringValue("HTTP/1.1"),
	"http2_required": types.BoolValue(false),
}
var DmHTTPVersionPolicyDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"reg_exp": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL matching expression", "", "").String,
			Computed:            true,
		},
		"version": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("HTTP version", "", "").AddStringEnum("HTTP/1.0", "HTTP/1.1", "HTTP/2").AddDefaultValue("HTTP/1.1").String,
			Computed:            true,
		},
		"http2_required": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Require HTTP/2", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
	},
}
var DmHTTPVersionPolicyResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"reg_exp": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL matching expression", "", "").String,
			Required:            true,
		},
		"version": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("HTTP version", "", "").AddStringEnum("HTTP/1.0", "HTTP/1.1", "HTTP/2").AddDefaultValue("HTTP/1.1").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("HTTP/1.0", "HTTP/1.1", "HTTP/2"),
			},
			Default: stringdefault.StaticString("HTTP/1.1"),
		},
		"http2_required": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Require HTTP/2", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
	},
}

func (data DmHTTPVersionPolicy) IsNull() bool {
	if !data.RegExp.IsNull() {
		return false
	}
	if !data.Version.IsNull() {
		return false
	}
	if !data.Http2Required.IsNull() {
		return false
	}
	return true
}

func (data DmHTTPVersionPolicy) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.RegExp.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RegExp`, data.RegExp.ValueString())
	}
	if !data.Version.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Version`, data.Version.ValueString())
	}
	if !data.Http2Required.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTP2Required`, tfutils.StringFromBool(data.Http2Required))
	}
	return body
}

func (data *DmHTTPVersionPolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `RegExp`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RegExp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RegExp = types.StringNull()
	}
	if value := res.Get(pathRoot + `Version`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Version = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Version = types.StringValue("HTTP/1.1")
	}
	if value := res.Get(pathRoot + `HTTP2Required`); value.Exists() {
		data.Http2Required = tfutils.BoolFromString(value.String())
	} else {
		data.Http2Required = types.BoolNull()
	}
}

func (data *DmHTTPVersionPolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `RegExp`); value.Exists() && !data.RegExp.IsNull() {
		data.RegExp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RegExp = types.StringNull()
	}
	if value := res.Get(pathRoot + `Version`); value.Exists() && !data.Version.IsNull() {
		data.Version = tfutils.ParseStringFromGJSON(value)
	} else if data.Version.ValueString() != "HTTP/1.1" {
		data.Version = types.StringNull()
	}
	if value := res.Get(pathRoot + `HTTP2Required`); value.Exists() && !data.Http2Required.IsNull() {
		data.Http2Required = tfutils.BoolFromString(value.String())
	} else if data.Http2Required.ValueBool() {
		data.Http2Required = types.BoolNull()
	}
}
