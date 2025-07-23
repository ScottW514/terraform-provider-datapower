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

type DmOAuthRole struct {
	Azsvr types.Bool `tfsdk:"azsvr"`
	Rssvr types.Bool `tfsdk:"rssvr"`
}

var DmOAuthRoleObjectType = map[string]attr.Type{
	"azsvr": types.BoolType,
	"rssvr": types.BoolType,
}
var DmOAuthRoleObjectDefault = map[string]attr.Value{
	"azsvr": types.BoolValue(false),
	"rssvr": types.BoolValue(false),
}
var DmOAuthRoleDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"azsvr": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Authorization and Token Endpoints", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"rssvr": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enforcement Point for Resource Server", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
	},
}
var DmOAuthRoleResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmOAuthRoleObjectType,
			DmOAuthRoleObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"azsvr": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Authorization and Token Endpoints", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"rssvr": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enforcement Point for Resource Server", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
	},
}

func (data DmOAuthRole) IsNull() bool {
	if !data.Azsvr.IsNull() {
		return false
	}
	if !data.Rssvr.IsNull() {
		return false
	}
	return true
}
func GetDmOAuthRoleDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmOAuthRoleDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmOAuthRoleDataSourceSchema
}

func GetDmOAuthRoleResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmOAuthRoleResourceSchema.Required = true
	} else {
		DmOAuthRoleResourceSchema.Optional = true
		DmOAuthRoleResourceSchema.Computed = true
	}
	DmOAuthRoleResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmOAuthRoleResourceSchema
}

func (data DmOAuthRole) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Azsvr.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`azsvr`, tfutils.StringFromBool(data.Azsvr, ""))
	}
	if !data.Rssvr.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`rssvr`, tfutils.StringFromBool(data.Rssvr, ""))
	}
	return body
}

func (data *DmOAuthRole) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `azsvr`); value.Exists() {
		data.Azsvr = tfutils.BoolFromString(value.String())
	} else {
		data.Azsvr = types.BoolNull()
	}
	if value := res.Get(pathRoot + `rssvr`); value.Exists() {
		data.Rssvr = tfutils.BoolFromString(value.String())
	} else {
		data.Rssvr = types.BoolNull()
	}
}

func (data *DmOAuthRole) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `azsvr`); value.Exists() && !data.Azsvr.IsNull() {
		data.Azsvr = tfutils.BoolFromString(value.String())
	} else if data.Azsvr.ValueBool() {
		data.Azsvr = types.BoolNull()
	}
	if value := res.Get(pathRoot + `rssvr`); value.Exists() && !data.Rssvr.IsNull() {
		data.Rssvr = tfutils.BoolFromString(value.String())
	} else if data.Rssvr.ValueBool() {
		data.Rssvr = types.BoolNull()
	}
}
