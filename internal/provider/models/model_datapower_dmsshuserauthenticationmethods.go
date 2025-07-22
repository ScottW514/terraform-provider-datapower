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

type DmSSHUserAuthenticationMethods struct {
	Publickey types.Bool `tfsdk:"publickey"`
	Password  types.Bool `tfsdk:"password"`
}

var DmSSHUserAuthenticationMethodsObjectType = map[string]attr.Type{
	"publickey": types.BoolType,
	"password":  types.BoolType,
}
var DmSSHUserAuthenticationMethodsObjectDefault = map[string]attr.Value{
	"publickey": types.BoolValue(true),
	"password":  types.BoolValue(true),
}
var DmSSHUserAuthenticationMethodsDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"publickey": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Public key", "", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"password": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Password", "", "").AddDefaultValue("true").String,
			Computed:            true,
		},
	},
}
var DmSSHUserAuthenticationMethodsResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmSSHUserAuthenticationMethodsObjectType,
			DmSSHUserAuthenticationMethodsObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"publickey": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Public key", "", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"password": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Password", "", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
	},
}

func (data DmSSHUserAuthenticationMethods) IsNull() bool {
	if !data.Publickey.IsNull() {
		return false
	}
	if !data.Password.IsNull() {
		return false
	}
	return true
}
func GetDmSSHUserAuthenticationMethodsDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmSSHUserAuthenticationMethodsDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmSSHUserAuthenticationMethodsDataSourceSchema
}

func GetDmSSHUserAuthenticationMethodsResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmSSHUserAuthenticationMethodsResourceSchema.Required = true
	} else {
		DmSSHUserAuthenticationMethodsResourceSchema.Optional = true
		DmSSHUserAuthenticationMethodsResourceSchema.Computed = true
	}
	DmSSHUserAuthenticationMethodsResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmSSHUserAuthenticationMethodsResourceSchema
}

func (data DmSSHUserAuthenticationMethods) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Publickey.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`publickey`, tfutils.StringFromBool(data.Publickey, false))
	}
	if !data.Password.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`password`, tfutils.StringFromBool(data.Password, false))
	}
	return body
}

func (data *DmSSHUserAuthenticationMethods) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `publickey`); value.Exists() {
		data.Publickey = tfutils.BoolFromString(value.String())
	} else {
		data.Publickey = types.BoolNull()
	}
	if value := res.Get(pathRoot + `password`); value.Exists() {
		data.Password = tfutils.BoolFromString(value.String())
	} else {
		data.Password = types.BoolNull()
	}
}

func (data *DmSSHUserAuthenticationMethods) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `publickey`); value.Exists() && !data.Publickey.IsNull() {
		data.Publickey = tfutils.BoolFromString(value.String())
	} else if !data.Publickey.ValueBool() {
		data.Publickey = types.BoolNull()
	}
	if value := res.Get(pathRoot + `password`); value.Exists() && !data.Password.IsNull() {
		data.Password = tfutils.BoolFromString(value.String())
	} else if !data.Password.ValueBool() {
		data.Password = types.BoolNull()
	}
}
