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

type DmOAuthProviderGrantType struct {
	None        types.Bool `tfsdk:"none"`
	Implicit    types.Bool `tfsdk:"implicit"`
	Password    types.Bool `tfsdk:"password"`
	Jwt         types.Bool `tfsdk:"jwt"`
	Application types.Bool `tfsdk:"application"`
	AccessCode  types.Bool `tfsdk:"access_code"`
}

var DmOAuthProviderGrantTypeObjectType = map[string]attr.Type{
	"none":        types.BoolType,
	"implicit":    types.BoolType,
	"password":    types.BoolType,
	"jwt":         types.BoolType,
	"application": types.BoolType,
	"access_code": types.BoolType,
}
var DmOAuthProviderGrantTypeObjectDefault = map[string]attr.Value{
	"none":        types.BoolValue(false),
	"implicit":    types.BoolValue(false),
	"password":    types.BoolValue(false),
	"jwt":         types.BoolValue(false),
	"application": types.BoolValue(false),
	"access_code": types.BoolValue(true),
}

func GetDmOAuthProviderGrantTypeDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmOAuthProviderGrantTypeDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"none": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"implicit": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Implicit", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"password": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Resource owner - Password", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"jwt": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Resource owner - JSON Web Token", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"application": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Client credentials", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"access_code": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Authorization code", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
		},
	}
	DmOAuthProviderGrantTypeDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmOAuthProviderGrantTypeDataSourceSchema
}
func GetDmOAuthProviderGrantTypeResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmOAuthProviderGrantTypeResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmOAuthProviderGrantTypeObjectType,
				DmOAuthProviderGrantTypeObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"none": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"implicit": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Implicit", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"password": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Resource owner - Password", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"jwt": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Resource owner - JSON Web Token", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"application": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Client credentials", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"access_code": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Authorization code", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
		},
	}
	DmOAuthProviderGrantTypeResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmOAuthProviderGrantTypeResourceSchema.Required = true
	} else {
		DmOAuthProviderGrantTypeResourceSchema.Optional = true
		DmOAuthProviderGrantTypeResourceSchema.Computed = true
	}
	return DmOAuthProviderGrantTypeResourceSchema
}

func (data DmOAuthProviderGrantType) IsNull() bool {
	if !data.None.IsNull() {
		return false
	}
	if !data.Implicit.IsNull() {
		return false
	}
	if !data.Password.IsNull() {
		return false
	}
	if !data.Jwt.IsNull() {
		return false
	}
	if !data.Application.IsNull() {
		return false
	}
	if !data.AccessCode.IsNull() {
		return false
	}
	return true
}

func (data DmOAuthProviderGrantType) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.None.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`none`, tfutils.StringFromBool(data.None, ""))
	}
	if !data.Implicit.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`implicit`, tfutils.StringFromBool(data.Implicit, ""))
	}
	if !data.Password.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`password`, tfutils.StringFromBool(data.Password, ""))
	}
	if !data.Jwt.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`jwt`, tfutils.StringFromBool(data.Jwt, ""))
	}
	if !data.Application.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`application`, tfutils.StringFromBool(data.Application, ""))
	}
	if !data.AccessCode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`access_code`, tfutils.StringFromBool(data.AccessCode, ""))
	}
	return body
}

func (data *DmOAuthProviderGrantType) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `none`); value.Exists() {
		data.None = tfutils.BoolFromString(value.String())
	} else {
		data.None = types.BoolNull()
	}
	if value := res.Get(pathRoot + `implicit`); value.Exists() {
		data.Implicit = tfutils.BoolFromString(value.String())
	} else {
		data.Implicit = types.BoolNull()
	}
	if value := res.Get(pathRoot + `password`); value.Exists() {
		data.Password = tfutils.BoolFromString(value.String())
	} else {
		data.Password = types.BoolNull()
	}
	if value := res.Get(pathRoot + `jwt`); value.Exists() {
		data.Jwt = tfutils.BoolFromString(value.String())
	} else {
		data.Jwt = types.BoolNull()
	}
	if value := res.Get(pathRoot + `application`); value.Exists() {
		data.Application = tfutils.BoolFromString(value.String())
	} else {
		data.Application = types.BoolNull()
	}
	if value := res.Get(pathRoot + `access_code`); value.Exists() {
		data.AccessCode = tfutils.BoolFromString(value.String())
	} else {
		data.AccessCode = types.BoolNull()
	}
}

func (data *DmOAuthProviderGrantType) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `none`); value.Exists() && !data.None.IsNull() {
		data.None = tfutils.BoolFromString(value.String())
	} else if data.None.ValueBool() {
		data.None = types.BoolNull()
	}
	if value := res.Get(pathRoot + `implicit`); value.Exists() && !data.Implicit.IsNull() {
		data.Implicit = tfutils.BoolFromString(value.String())
	} else if data.Implicit.ValueBool() {
		data.Implicit = types.BoolNull()
	}
	if value := res.Get(pathRoot + `password`); value.Exists() && !data.Password.IsNull() {
		data.Password = tfutils.BoolFromString(value.String())
	} else if data.Password.ValueBool() {
		data.Password = types.BoolNull()
	}
	if value := res.Get(pathRoot + `jwt`); value.Exists() && !data.Jwt.IsNull() {
		data.Jwt = tfutils.BoolFromString(value.String())
	} else if data.Jwt.ValueBool() {
		data.Jwt = types.BoolNull()
	}
	if value := res.Get(pathRoot + `application`); value.Exists() && !data.Application.IsNull() {
		data.Application = tfutils.BoolFromString(value.String())
	} else if data.Application.ValueBool() {
		data.Application = types.BoolNull()
	}
	if value := res.Get(pathRoot + `access_code`); value.Exists() && !data.AccessCode.IsNull() {
		data.AccessCode = tfutils.BoolFromString(value.String())
	} else if !data.AccessCode.ValueBool() {
		data.AccessCode = types.BoolNull()
	}
}
