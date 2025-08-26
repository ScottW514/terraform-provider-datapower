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

type DmOAuthRSSetHeader struct {
	Owner    types.Bool `tfsdk:"owner"`
	ClientId types.Bool `tfsdk:"client_id"`
	Scope    types.Bool `tfsdk:"scope"`
	Miscinfo types.Bool `tfsdk:"miscinfo"`
}

var DmOAuthRSSetHeaderObjectType = map[string]attr.Type{
	"owner":     types.BoolType,
	"client_id": types.BoolType,
	"scope":     types.BoolType,
	"miscinfo":  types.BoolType,
}
var DmOAuthRSSetHeaderObjectDefault = map[string]attr.Value{
	"owner":     types.BoolValue(false),
	"client_id": types.BoolValue(false),
	"scope":     types.BoolValue(false),
	"miscinfo":  types.BoolValue(false),
}

func GetDmOAuthRSSetHeaderDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmOAuthRSSetHeaderDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"owner": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Resource Owner", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"client_id": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Client ID", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"scope": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Scope", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"miscinfo": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Customized Info", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
		},
	}
	DmOAuthRSSetHeaderDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmOAuthRSSetHeaderDataSourceSchema
}
func GetDmOAuthRSSetHeaderResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmOAuthRSSetHeaderResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmOAuthRSSetHeaderObjectType,
				DmOAuthRSSetHeaderObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"owner": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Resource Owner", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"client_id": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Client ID", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"scope": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Scope", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"miscinfo": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Customized Info", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
	DmOAuthRSSetHeaderResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmOAuthRSSetHeaderResourceSchema.Required = true
	} else {
		DmOAuthRSSetHeaderResourceSchema.Optional = true
		DmOAuthRSSetHeaderResourceSchema.Computed = true
	}
	return DmOAuthRSSetHeaderResourceSchema
}

func (data DmOAuthRSSetHeader) IsNull() bool {
	if !data.Owner.IsNull() {
		return false
	}
	if !data.ClientId.IsNull() {
		return false
	}
	if !data.Scope.IsNull() {
		return false
	}
	if !data.Miscinfo.IsNull() {
		return false
	}
	return true
}

func (data DmOAuthRSSetHeader) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Owner.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`owner`, tfutils.StringFromBool(data.Owner, ""))
	}
	if !data.ClientId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`clientid`, tfutils.StringFromBool(data.ClientId, ""))
	}
	if !data.Scope.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`scope`, tfutils.StringFromBool(data.Scope, ""))
	}
	if !data.Miscinfo.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`miscinfo`, tfutils.StringFromBool(data.Miscinfo, ""))
	}
	return body
}

func (data *DmOAuthRSSetHeader) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `owner`); value.Exists() {
		data.Owner = tfutils.BoolFromString(value.String())
	} else {
		data.Owner = types.BoolNull()
	}
	if value := res.Get(pathRoot + `clientid`); value.Exists() {
		data.ClientId = tfutils.BoolFromString(value.String())
	} else {
		data.ClientId = types.BoolNull()
	}
	if value := res.Get(pathRoot + `scope`); value.Exists() {
		data.Scope = tfutils.BoolFromString(value.String())
	} else {
		data.Scope = types.BoolNull()
	}
	if value := res.Get(pathRoot + `miscinfo`); value.Exists() {
		data.Miscinfo = tfutils.BoolFromString(value.String())
	} else {
		data.Miscinfo = types.BoolNull()
	}
}

func (data *DmOAuthRSSetHeader) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `owner`); value.Exists() && !data.Owner.IsNull() {
		data.Owner = tfutils.BoolFromString(value.String())
	} else if data.Owner.ValueBool() {
		data.Owner = types.BoolNull()
	}
	if value := res.Get(pathRoot + `clientid`); value.Exists() && !data.ClientId.IsNull() {
		data.ClientId = tfutils.BoolFromString(value.String())
	} else if data.ClientId.ValueBool() {
		data.ClientId = types.BoolNull()
	}
	if value := res.Get(pathRoot + `scope`); value.Exists() && !data.Scope.IsNull() {
		data.Scope = tfutils.BoolFromString(value.String())
	} else if data.Scope.ValueBool() {
		data.Scope = types.BoolNull()
	}
	if value := res.Get(pathRoot + `miscinfo`); value.Exists() && !data.Miscinfo.IsNull() {
		data.Miscinfo = tfutils.BoolFromString(value.String())
	} else if data.Miscinfo.ValueBool() {
		data.Miscinfo = types.BoolNull()
	}
}
