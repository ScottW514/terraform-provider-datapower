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

type DmJWTClaims struct {
	Aud    types.Bool `tfsdk:"aud"`
	Nbf    types.Bool `tfsdk:"nbf"`
	Iat    types.Bool `tfsdk:"iat"`
	Jti    types.Bool `tfsdk:"jti"`
	Nonce  types.Bool `tfsdk:"nonce"`
	Custom types.Bool `tfsdk:"custom"`
}

var DmJWTClaimsObjectType = map[string]attr.Type{
	"aud":    types.BoolType,
	"nbf":    types.BoolType,
	"iat":    types.BoolType,
	"jti":    types.BoolType,
	"nonce":  types.BoolType,
	"custom": types.BoolType,
}
var DmJWTClaimsObjectDefault = map[string]attr.Value{
	"aud":    types.BoolValue(false),
	"nbf":    types.BoolValue(false),
	"iat":    types.BoolValue(false),
	"jti":    types.BoolValue(false),
	"nonce":  types.BoolValue(false),
	"custom": types.BoolValue(false),
}
var DmJWTClaimsDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"aud": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Audience", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"nbf": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Not before", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"iat": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Issued at", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"jti": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("JWT ID", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"nonce": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Nonce", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"custom": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Custom", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
	},
}
var DmJWTClaimsResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmJWTClaimsObjectType,
			DmJWTClaimsObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"aud": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Audience", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"nbf": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Not before", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"iat": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Issued at", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"jti": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("JWT ID", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"nonce": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Nonce", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"custom": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Custom", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
	},
}

func (data DmJWTClaims) IsNull() bool {
	if !data.Aud.IsNull() {
		return false
	}
	if !data.Nbf.IsNull() {
		return false
	}
	if !data.Iat.IsNull() {
		return false
	}
	if !data.Jti.IsNull() {
		return false
	}
	if !data.Nonce.IsNull() {
		return false
	}
	if !data.Custom.IsNull() {
		return false
	}
	return true
}
func GetDmJWTClaimsDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmJWTClaimsDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmJWTClaimsDataSourceSchema
}

func GetDmJWTClaimsResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmJWTClaimsResourceSchema.Required = true
	} else {
		DmJWTClaimsResourceSchema.Optional = true
		DmJWTClaimsResourceSchema.Computed = true
	}
	DmJWTClaimsResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmJWTClaimsResourceSchema
}

func (data DmJWTClaims) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Aud.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`aud`, tfutils.StringFromBool(data.Aud))
	}
	if !data.Nbf.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`nbf`, tfutils.StringFromBool(data.Nbf))
	}
	if !data.Iat.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`iat`, tfutils.StringFromBool(data.Iat))
	}
	if !data.Jti.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`jti`, tfutils.StringFromBool(data.Jti))
	}
	if !data.Nonce.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`nonce`, tfutils.StringFromBool(data.Nonce))
	}
	if !data.Custom.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`custom`, tfutils.StringFromBool(data.Custom))
	}
	return body
}

func (data *DmJWTClaims) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `aud`); value.Exists() {
		data.Aud = tfutils.BoolFromString(value.String())
	} else {
		data.Aud = types.BoolNull()
	}
	if value := res.Get(pathRoot + `nbf`); value.Exists() {
		data.Nbf = tfutils.BoolFromString(value.String())
	} else {
		data.Nbf = types.BoolNull()
	}
	if value := res.Get(pathRoot + `iat`); value.Exists() {
		data.Iat = tfutils.BoolFromString(value.String())
	} else {
		data.Iat = types.BoolNull()
	}
	if value := res.Get(pathRoot + `jti`); value.Exists() {
		data.Jti = tfutils.BoolFromString(value.String())
	} else {
		data.Jti = types.BoolNull()
	}
	if value := res.Get(pathRoot + `nonce`); value.Exists() {
		data.Nonce = tfutils.BoolFromString(value.String())
	} else {
		data.Nonce = types.BoolNull()
	}
	if value := res.Get(pathRoot + `custom`); value.Exists() {
		data.Custom = tfutils.BoolFromString(value.String())
	} else {
		data.Custom = types.BoolNull()
	}
}

func (data *DmJWTClaims) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `aud`); value.Exists() && !data.Aud.IsNull() {
		data.Aud = tfutils.BoolFromString(value.String())
	} else if data.Aud.ValueBool() {
		data.Aud = types.BoolNull()
	}
	if value := res.Get(pathRoot + `nbf`); value.Exists() && !data.Nbf.IsNull() {
		data.Nbf = tfutils.BoolFromString(value.String())
	} else if data.Nbf.ValueBool() {
		data.Nbf = types.BoolNull()
	}
	if value := res.Get(pathRoot + `iat`); value.Exists() && !data.Iat.IsNull() {
		data.Iat = tfutils.BoolFromString(value.String())
	} else if data.Iat.ValueBool() {
		data.Iat = types.BoolNull()
	}
	if value := res.Get(pathRoot + `jti`); value.Exists() && !data.Jti.IsNull() {
		data.Jti = tfutils.BoolFromString(value.String())
	} else if data.Jti.ValueBool() {
		data.Jti = types.BoolNull()
	}
	if value := res.Get(pathRoot + `nonce`); value.Exists() && !data.Nonce.IsNull() {
		data.Nonce = tfutils.BoolFromString(value.String())
	} else if data.Nonce.ValueBool() {
		data.Nonce = types.BoolNull()
	}
	if value := res.Get(pathRoot + `custom`); value.Exists() && !data.Custom.IsNull() {
		data.Custom = tfutils.BoolFromString(value.String())
	} else if data.Custom.ValueBool() {
		data.Custom = types.BoolNull()
	}
}
