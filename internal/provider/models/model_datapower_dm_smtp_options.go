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

type DmSMTPOptions struct {
	StartTls types.Bool `tfsdk:"start_tls"`
	Auth     types.Bool `tfsdk:"auth"`
}

var DmSMTPOptionsObjectType = map[string]attr.Type{
	"start_tls": types.BoolType,
	"auth":      types.BoolType,
}
var DmSMTPOptionsObjectDefault = map[string]attr.Value{
	"start_tls": types.BoolValue(false),
	"auth":      types.BoolValue(false),
}
var DmSMTPOptionsDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"start_tls": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("STARTTLS", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"auth": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Authentication", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
	},
}
var DmSMTPOptionsResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmSMTPOptionsObjectType,
			DmSMTPOptionsObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"start_tls": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("STARTTLS", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"auth": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Authentication", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
	},
}

func (data DmSMTPOptions) IsNull() bool {
	if !data.StartTls.IsNull() {
		return false
	}
	if !data.Auth.IsNull() {
		return false
	}
	return true
}
func GetDmSMTPOptionsDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmSMTPOptionsDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmSMTPOptionsDataSourceSchema
}

func GetDmSMTPOptionsResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmSMTPOptionsResourceSchema.Required = true
	} else {
		DmSMTPOptionsResourceSchema.Optional = true
		DmSMTPOptionsResourceSchema.Computed = true
	}
	DmSMTPOptionsResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmSMTPOptionsResourceSchema
}

func (data DmSMTPOptions) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.StartTls.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StartTLS`, tfutils.StringFromBool(data.StartTls, ""))
	}
	if !data.Auth.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Auth`, tfutils.StringFromBool(data.Auth, ""))
	}
	return body
}

func (data *DmSMTPOptions) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `StartTLS`); value.Exists() {
		data.StartTls = tfutils.BoolFromString(value.String())
	} else {
		data.StartTls = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Auth`); value.Exists() {
		data.Auth = tfutils.BoolFromString(value.String())
	} else {
		data.Auth = types.BoolNull()
	}
}

func (data *DmSMTPOptions) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `StartTLS`); value.Exists() && !data.StartTls.IsNull() {
		data.StartTls = tfutils.BoolFromString(value.String())
	} else if data.StartTls.ValueBool() {
		data.StartTls = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Auth`); value.Exists() && !data.Auth.IsNull() {
		data.Auth = tfutils.BoolFromString(value.String())
	} else if data.Auth.ValueBool() {
		data.Auth = types.BoolNull()
	}
}
