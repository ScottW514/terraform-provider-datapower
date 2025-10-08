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

type DmSecurityType struct {
	BasicAuth types.Bool `tfsdk:"basic_auth"`
}

var DmSecurityTypeObjectType = map[string]attr.Type{
	"basic_auth": types.BoolType,
}
var DmSecurityTypeObjectDefault = map[string]attr.Value{
	"basic_auth": types.BoolValue(true),
}

func GetDmSecurityTypeDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmSecurityTypeDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"basic_auth": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Basic authentication header",
				Computed:            true,
			},
		},
	}
	DmSecurityTypeDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmSecurityTypeDataSourceSchema
}
func GetDmSecurityTypeResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmSecurityTypeResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmSecurityTypeObjectType,
				DmSecurityTypeObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"basic_auth": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Basic authentication header", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
		},
	}
	DmSecurityTypeResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmSecurityTypeResourceSchema.Required = true
	} else {
		DmSecurityTypeResourceSchema.Optional = true
		DmSecurityTypeResourceSchema.Computed = true
	}
	return DmSecurityTypeResourceSchema
}

func (data DmSecurityType) IsNull() bool {
	if !data.BasicAuth.IsNull() {
		return false
	}
	return true
}

func (data DmSecurityType) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.BasicAuth.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`basic-auth`, tfutils.StringFromBool(data.BasicAuth, ""))
	}
	return body
}

func (data *DmSecurityType) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `basic-auth`); value.Exists() {
		data.BasicAuth = tfutils.BoolFromString(value.String())
	} else {
		data.BasicAuth = types.BoolNull()
	}
}

func (data *DmSecurityType) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `basic-auth`); value.Exists() && !data.BasicAuth.IsNull() {
		data.BasicAuth = tfutils.BoolFromString(value.String())
	} else if !data.BasicAuth.ValueBool() {
		data.BasicAuth = types.BoolNull()
	}
}
