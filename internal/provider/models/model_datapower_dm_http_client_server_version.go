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

type DmHTTPClientServerVersion struct {
	Front types.String `tfsdk:"front"`
	Back  types.String `tfsdk:"back"`
}

var DmHTTPClientServerVersionObjectType = map[string]attr.Type{
	"front": types.StringType,
	"back":  types.StringType,
}
var DmHTTPClientServerVersionObjectDefault = map[string]attr.Value{
	"front": types.StringValue("HTTP/1.1"),
	"back":  types.StringValue("HTTP/1.1"),
}

func GetDmHTTPClientServerVersionDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmHTTPClientServerVersionDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"front": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Select the HTTP version to use on the client-side connection. The default is HTTP 1.1.",
				Computed:            true,
			},
			"back": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Select the HTTP version to use on the server-side connection. The default is HTTP 1.1.",
				Computed:            true,
			},
		},
	}
	DmHTTPClientServerVersionDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmHTTPClientServerVersionDataSourceSchema
}
func GetDmHTTPClientServerVersionResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmHTTPClientServerVersionResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmHTTPClientServerVersionObjectType,
				DmHTTPClientServerVersionObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"front": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the HTTP version to use on the client-side connection. The default is HTTP 1.1.", "", "").AddStringEnum("HTTP/1.0", "HTTP/1.1").AddDefaultValue("HTTP/1.1").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("HTTP/1.0", "HTTP/1.1"),
				},
				Default: stringdefault.StaticString("HTTP/1.1"),
			},
			"back": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the HTTP version to use on the server-side connection. The default is HTTP 1.1.", "", "").AddStringEnum("HTTP/1.0", "HTTP/1.1").AddDefaultValue("HTTP/1.1").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("HTTP/1.0", "HTTP/1.1"),
				},
				Default: stringdefault.StaticString("HTTP/1.1"),
			},
		},
	}
	DmHTTPClientServerVersionResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmHTTPClientServerVersionResourceSchema.Required = true
	} else {
		DmHTTPClientServerVersionResourceSchema.Optional = true
		DmHTTPClientServerVersionResourceSchema.Computed = true
	}
	return DmHTTPClientServerVersionResourceSchema
}

func (data DmHTTPClientServerVersion) IsNull() bool {
	if !data.Front.IsNull() {
		return false
	}
	if !data.Back.IsNull() {
		return false
	}
	return true
}

func (data DmHTTPClientServerVersion) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Front.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Front`, data.Front.ValueString())
	}
	if !data.Back.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Back`, data.Back.ValueString())
	}
	return body
}

func (data *DmHTTPClientServerVersion) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Front`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Front = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Front = types.StringValue("HTTP/1.1")
	}
	if value := res.Get(pathRoot + `Back`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Back = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Back = types.StringValue("HTTP/1.1")
	}
}

func (data *DmHTTPClientServerVersion) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Front`); value.Exists() && !data.Front.IsNull() {
		data.Front = tfutils.ParseStringFromGJSON(value)
	} else if data.Front.ValueString() != "HTTP/1.1" {
		data.Front = types.StringNull()
	}
	if value := res.Get(pathRoot + `Back`); value.Exists() && !data.Back.IsNull() {
		data.Back = tfutils.ParseStringFromGJSON(value)
	} else if data.Back.ValueString() != "HTTP/1.1" {
		data.Back = types.StringNull()
	}
}
