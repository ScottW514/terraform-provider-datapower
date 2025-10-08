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

type DmCookieAttribute struct {
	Domain   types.Bool `tfsdk:"domain"`
	Path     types.Bool `tfsdk:"path"`
	Secure   types.Bool `tfsdk:"secure"`
	Httponly types.Bool `tfsdk:"httponly"`
	MaxAge   types.Bool `tfsdk:"max_age"`
	Expires  types.Bool `tfsdk:"expires"`
	Custom   types.Bool `tfsdk:"custom"`
}

var DmCookieAttributeObjectType = map[string]attr.Type{
	"domain":   types.BoolType,
	"path":     types.BoolType,
	"secure":   types.BoolType,
	"httponly": types.BoolType,
	"max_age":  types.BoolType,
	"expires":  types.BoolType,
	"custom":   types.BoolType,
}
var DmCookieAttributeObjectDefault = map[string]attr.Value{
	"domain":   types.BoolValue(false),
	"path":     types.BoolValue(false),
	"secure":   types.BoolValue(false),
	"httponly": types.BoolValue(false),
	"max_age":  types.BoolValue(false),
	"expires":  types.BoolValue(false),
	"custom":   types.BoolValue(false),
}

func GetDmCookieAttributeDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmCookieAttributeDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"domain": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Domain",
				Computed:            true,
			},
			"path": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Path",
				Computed:            true,
			},
			"secure": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Secure",
				Computed:            true,
			},
			"httponly": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "HTTPOnly",
				Computed:            true,
			},
			"max_age": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Max-Age",
				Computed:            true,
			},
			"expires": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Expires",
				Computed:            true,
			},
			"custom": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Custom attributes",
				Computed:            true,
			},
		},
	}
	DmCookieAttributeDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmCookieAttributeDataSourceSchema
}
func GetDmCookieAttributeResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmCookieAttributeResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmCookieAttributeObjectType,
				DmCookieAttributeObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"domain": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Domain", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"path": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Path", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"secure": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Secure", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"httponly": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTPOnly", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"max_age": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Max-Age", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"expires": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Expires", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"custom": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Custom attributes", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
	DmCookieAttributeResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmCookieAttributeResourceSchema.Required = true
	} else {
		DmCookieAttributeResourceSchema.Optional = true
		DmCookieAttributeResourceSchema.Computed = true
	}
	return DmCookieAttributeResourceSchema
}

func (data DmCookieAttribute) IsNull() bool {
	if !data.Domain.IsNull() {
		return false
	}
	if !data.Path.IsNull() {
		return false
	}
	if !data.Secure.IsNull() {
		return false
	}
	if !data.Httponly.IsNull() {
		return false
	}
	if !data.MaxAge.IsNull() {
		return false
	}
	if !data.Expires.IsNull() {
		return false
	}
	if !data.Custom.IsNull() {
		return false
	}
	return true
}

func (data DmCookieAttribute) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Domain.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`domain`, tfutils.StringFromBool(data.Domain, ""))
	}
	if !data.Path.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`path`, tfutils.StringFromBool(data.Path, ""))
	}
	if !data.Secure.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`secure`, tfutils.StringFromBool(data.Secure, ""))
	}
	if !data.Httponly.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`httponly`, tfutils.StringFromBool(data.Httponly, ""))
	}
	if !data.MaxAge.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`max-age`, tfutils.StringFromBool(data.MaxAge, ""))
	}
	if !data.Expires.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`expires`, tfutils.StringFromBool(data.Expires, ""))
	}
	if !data.Custom.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`custom`, tfutils.StringFromBool(data.Custom, ""))
	}
	return body
}

func (data *DmCookieAttribute) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `domain`); value.Exists() {
		data.Domain = tfutils.BoolFromString(value.String())
	} else {
		data.Domain = types.BoolNull()
	}
	if value := res.Get(pathRoot + `path`); value.Exists() {
		data.Path = tfutils.BoolFromString(value.String())
	} else {
		data.Path = types.BoolNull()
	}
	if value := res.Get(pathRoot + `secure`); value.Exists() {
		data.Secure = tfutils.BoolFromString(value.String())
	} else {
		data.Secure = types.BoolNull()
	}
	if value := res.Get(pathRoot + `httponly`); value.Exists() {
		data.Httponly = tfutils.BoolFromString(value.String())
	} else {
		data.Httponly = types.BoolNull()
	}
	if value := res.Get(pathRoot + `max-age`); value.Exists() {
		data.MaxAge = tfutils.BoolFromString(value.String())
	} else {
		data.MaxAge = types.BoolNull()
	}
	if value := res.Get(pathRoot + `expires`); value.Exists() {
		data.Expires = tfutils.BoolFromString(value.String())
	} else {
		data.Expires = types.BoolNull()
	}
	if value := res.Get(pathRoot + `custom`); value.Exists() {
		data.Custom = tfutils.BoolFromString(value.String())
	} else {
		data.Custom = types.BoolNull()
	}
}

func (data *DmCookieAttribute) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `domain`); value.Exists() && !data.Domain.IsNull() {
		data.Domain = tfutils.BoolFromString(value.String())
	} else if data.Domain.ValueBool() {
		data.Domain = types.BoolNull()
	}
	if value := res.Get(pathRoot + `path`); value.Exists() && !data.Path.IsNull() {
		data.Path = tfutils.BoolFromString(value.String())
	} else if data.Path.ValueBool() {
		data.Path = types.BoolNull()
	}
	if value := res.Get(pathRoot + `secure`); value.Exists() && !data.Secure.IsNull() {
		data.Secure = tfutils.BoolFromString(value.String())
	} else if data.Secure.ValueBool() {
		data.Secure = types.BoolNull()
	}
	if value := res.Get(pathRoot + `httponly`); value.Exists() && !data.Httponly.IsNull() {
		data.Httponly = tfutils.BoolFromString(value.String())
	} else if data.Httponly.ValueBool() {
		data.Httponly = types.BoolNull()
	}
	if value := res.Get(pathRoot + `max-age`); value.Exists() && !data.MaxAge.IsNull() {
		data.MaxAge = tfutils.BoolFromString(value.String())
	} else if data.MaxAge.ValueBool() {
		data.MaxAge = types.BoolNull()
	}
	if value := res.Get(pathRoot + `expires`); value.Exists() && !data.Expires.IsNull() {
		data.Expires = tfutils.BoolFromString(value.String())
	} else if data.Expires.ValueBool() {
		data.Expires = types.BoolNull()
	}
	if value := res.Get(pathRoot + `custom`); value.Exists() && !data.Custom.IsNull() {
		data.Custom = tfutils.BoolFromString(value.String())
	} else if data.Custom.ValueBool() {
		data.Custom = types.BoolNull()
	}
}
