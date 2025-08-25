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

type DmOAuthAZGrantType struct {
	Code         types.Bool `tfsdk:"code"`
	Implicit     types.Bool `tfsdk:"implicit"`
	Password     types.Bool `tfsdk:"password"`
	Client       types.Bool `tfsdk:"client"`
	Jwt          types.Bool `tfsdk:"jwt"`
	Novalidate   types.Bool `tfsdk:"novalidate"`
	Oidc         types.Bool `tfsdk:"oidc"`
	Saml20bearer types.Bool `tfsdk:"saml20bearer"`
}

var DmOAuthAZGrantTypeObjectType = map[string]attr.Type{
	"code":         types.BoolType,
	"implicit":     types.BoolType,
	"password":     types.BoolType,
	"client":       types.BoolType,
	"jwt":          types.BoolType,
	"novalidate":   types.BoolType,
	"oidc":         types.BoolType,
	"saml20bearer": types.BoolType,
}
var DmOAuthAZGrantTypeObjectDefault = map[string]attr.Value{
	"code":         types.BoolValue(false),
	"implicit":     types.BoolValue(false),
	"password":     types.BoolValue(false),
	"client":       types.BoolValue(false),
	"jwt":          types.BoolValue(false),
	"novalidate":   types.BoolValue(false),
	"oidc":         types.BoolValue(false),
	"saml20bearer": types.BoolValue(false),
}

func GetDmOAuthAZGrantTypeDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmOAuthAZGrantTypeDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"code": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Authorization Code Grant", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"implicit": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Implicit Grant", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"password": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Resource Owner Password Credential Grant", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"client": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Client Credentials Grant", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"jwt": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("JWT Grant", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"novalidate": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Disable Validation Grant", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"oidc": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("OpenID Connect", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"saml20bearer": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
		},
	}
	DmOAuthAZGrantTypeDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmOAuthAZGrantTypeDataSourceSchema
}
func GetDmOAuthAZGrantTypeResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmOAuthAZGrantTypeResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmOAuthAZGrantTypeObjectType,
				DmOAuthAZGrantTypeObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"code": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Authorization Code Grant", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"implicit": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Implicit Grant", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"password": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Resource Owner Password Credential Grant", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"client": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Client Credentials Grant", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"jwt": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("JWT Grant", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"novalidate": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Disable Validation Grant", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"oidc": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("OpenID Connect", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"saml20bearer": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
	DmOAuthAZGrantTypeResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmOAuthAZGrantTypeResourceSchema.Required = true
	} else {
		DmOAuthAZGrantTypeResourceSchema.Optional = true
		DmOAuthAZGrantTypeResourceSchema.Computed = true
	}
	return DmOAuthAZGrantTypeResourceSchema
}

func (data DmOAuthAZGrantType) IsNull() bool {
	if !data.Code.IsNull() {
		return false
	}
	if !data.Implicit.IsNull() {
		return false
	}
	if !data.Password.IsNull() {
		return false
	}
	if !data.Client.IsNull() {
		return false
	}
	if !data.Jwt.IsNull() {
		return false
	}
	if !data.Novalidate.IsNull() {
		return false
	}
	if !data.Oidc.IsNull() {
		return false
	}
	if !data.Saml20bearer.IsNull() {
		return false
	}
	return true
}

func (data DmOAuthAZGrantType) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Code.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`code`, tfutils.StringFromBool(data.Code, ""))
	}
	if !data.Implicit.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`implicit`, tfutils.StringFromBool(data.Implicit, ""))
	}
	if !data.Password.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`password`, tfutils.StringFromBool(data.Password, ""))
	}
	if !data.Client.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`client`, tfutils.StringFromBool(data.Client, ""))
	}
	if !data.Jwt.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`jwt`, tfutils.StringFromBool(data.Jwt, ""))
	}
	if !data.Novalidate.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`novalidate`, tfutils.StringFromBool(data.Novalidate, ""))
	}
	if !data.Oidc.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`oidc`, tfutils.StringFromBool(data.Oidc, ""))
	}
	if !data.Saml20bearer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`saml20bearer`, tfutils.StringFromBool(data.Saml20bearer, ""))
	}
	return body
}

func (data *DmOAuthAZGrantType) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `code`); value.Exists() {
		data.Code = tfutils.BoolFromString(value.String())
	} else {
		data.Code = types.BoolNull()
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
	if value := res.Get(pathRoot + `client`); value.Exists() {
		data.Client = tfutils.BoolFromString(value.String())
	} else {
		data.Client = types.BoolNull()
	}
	if value := res.Get(pathRoot + `jwt`); value.Exists() {
		data.Jwt = tfutils.BoolFromString(value.String())
	} else {
		data.Jwt = types.BoolNull()
	}
	if value := res.Get(pathRoot + `novalidate`); value.Exists() {
		data.Novalidate = tfutils.BoolFromString(value.String())
	} else {
		data.Novalidate = types.BoolNull()
	}
	if value := res.Get(pathRoot + `oidc`); value.Exists() {
		data.Oidc = tfutils.BoolFromString(value.String())
	} else {
		data.Oidc = types.BoolNull()
	}
	if value := res.Get(pathRoot + `saml20bearer`); value.Exists() {
		data.Saml20bearer = tfutils.BoolFromString(value.String())
	} else {
		data.Saml20bearer = types.BoolNull()
	}
}

func (data *DmOAuthAZGrantType) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `code`); value.Exists() && !data.Code.IsNull() {
		data.Code = tfutils.BoolFromString(value.String())
	} else if data.Code.ValueBool() {
		data.Code = types.BoolNull()
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
	if value := res.Get(pathRoot + `client`); value.Exists() && !data.Client.IsNull() {
		data.Client = tfutils.BoolFromString(value.String())
	} else if data.Client.ValueBool() {
		data.Client = types.BoolNull()
	}
	if value := res.Get(pathRoot + `jwt`); value.Exists() && !data.Jwt.IsNull() {
		data.Jwt = tfutils.BoolFromString(value.String())
	} else if data.Jwt.ValueBool() {
		data.Jwt = types.BoolNull()
	}
	if value := res.Get(pathRoot + `novalidate`); value.Exists() && !data.Novalidate.IsNull() {
		data.Novalidate = tfutils.BoolFromString(value.String())
	} else if data.Novalidate.ValueBool() {
		data.Novalidate = types.BoolNull()
	}
	if value := res.Get(pathRoot + `oidc`); value.Exists() && !data.Oidc.IsNull() {
		data.Oidc = tfutils.BoolFromString(value.String())
	} else if data.Oidc.ValueBool() {
		data.Oidc = types.BoolNull()
	}
	if value := res.Get(pathRoot + `saml20bearer`); value.Exists() && !data.Saml20bearer.IsNull() {
		data.Saml20bearer = tfutils.BoolFromString(value.String())
	} else if data.Saml20bearer.ValueBool() {
		data.Saml20bearer = types.BoolNull()
	}
}
