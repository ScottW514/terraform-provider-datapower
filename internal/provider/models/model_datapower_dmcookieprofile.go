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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmCookieProfile struct {
	CookiePolicy          types.String `tfsdk:"cookie_policy"`
	Type                  types.String `tfsdk:"type"`
	Key                   types.String `tfsdk:"key"`
	IPinWatermark         types.Bool   `tfsdk:"i_pin_watermark"`
	CookieGnvc            types.String `tfsdk:"cookie_gnvc"`
	UseSharedSecretObject types.Bool   `tfsdk:"use_shared_secret_object"`
}

var DmCookieProfileObjectType = map[string]attr.Type{
	"cookie_policy":            types.StringType,
	"type":                     types.StringType,
	"key":                      types.StringType,
	"i_pin_watermark":          types.BoolType,
	"cookie_gnvc":              types.StringType,
	"use_shared_secret_object": types.BoolType,
}
var DmCookieProfileObjectDefault = map[string]attr.Value{
	"cookie_policy":            types.StringValue("allow"),
	"type":                     types.StringValue("none"),
	"key":                      types.StringValue("secret"),
	"i_pin_watermark":          types.BoolValue(true),
	"cookie_gnvc":              types.StringNull(),
	"use_shared_secret_object": types.BoolValue(false),
}
var DmCookieProfileDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"cookie_policy": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Allow Cookies", "", "").AddStringEnum("allow", "require", "deny").AddDefaultValue("allow").String,
			Computed:            true,
		},
		"type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Sign or Encrypt Cookies", "", "").AddStringEnum("none", "sign", "encrypt").AddDefaultValue("none").String,
			Computed:            true,
		},
		"key": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Secret Key Passphrase", "", "").AddDefaultValue("secret").String,
			Computed:            true,
		},
		"i_pin_watermark": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("IP Address specific Cookies", "", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"cookie_gnvc": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Cookie Content Name-Value Profile", "request-cookie-gnvc", "namevalueprofile").String,
			Computed:            true,
		},
		"use_shared_secret_object": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Use Shared Secret Key", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
	},
}
var DmCookieProfileResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmCookieProfileObjectType,
			DmCookieProfileObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"cookie_policy": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Allow Cookies", "", "").AddStringEnum("allow", "require", "deny").AddDefaultValue("allow").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("allow", "require", "deny"),
			},
			Default: stringdefault.StaticString("allow"),
		},
		"type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Sign or Encrypt Cookies", "", "").AddStringEnum("none", "sign", "encrypt").AddDefaultValue("none").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("none", "sign", "encrypt"),
			},
			Default: stringdefault.StaticString("none"),
		},
		"key": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Secret Key Passphrase", "", "").AddDefaultValue("secret").String,
			Computed:            true,
			Optional:            true,
			Default:             stringdefault.StaticString("secret"),
		},
		"i_pin_watermark": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("IP Address specific Cookies", "", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"cookie_gnvc": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Cookie Content Name-Value Profile", "request-cookie-gnvc", "namevalueprofile").String,
			Optional:            true,
		},
		"use_shared_secret_object": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Use Shared Secret Key", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
	},
}

func (data DmCookieProfile) IsNull() bool {
	if !data.CookiePolicy.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
		return false
	}
	if !data.Key.IsNull() {
		return false
	}
	if !data.IPinWatermark.IsNull() {
		return false
	}
	if !data.CookieGnvc.IsNull() {
		return false
	}
	if !data.UseSharedSecretObject.IsNull() {
		return false
	}
	return true
}
func GetDmCookieProfileDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmCookieProfileDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmCookieProfileDataSourceSchema
}

func GetDmCookieProfileResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmCookieProfileResourceSchema.Required = true
	} else {
		DmCookieProfileResourceSchema.Optional = true
		DmCookieProfileResourceSchema.Computed = true
	}
	DmCookieProfileResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmCookieProfileResourceSchema
}

func (data DmCookieProfile) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.CookiePolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CookiePolicy`, data.CookiePolicy.ValueString())
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Type`, data.Type.ValueString())
	}
	if !data.Key.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Key`, data.Key.ValueString())
	}
	if !data.IPinWatermark.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IPinWatermark`, tfutils.StringFromBool(data.IPinWatermark, ""))
	}
	if !data.CookieGnvc.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CookieGNVC`, data.CookieGnvc.ValueString())
	}
	if !data.UseSharedSecretObject.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseSharedSecretObject`, tfutils.StringFromBool(data.UseSharedSecretObject, ""))
	}
	return body
}

func (data *DmCookieProfile) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `CookiePolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CookiePolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CookiePolicy = types.StringValue("allow")
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringValue("none")
	}
	if value := res.Get(pathRoot + `Key`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Key = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Key = types.StringValue("secret")
	}
	if value := res.Get(pathRoot + `IPinWatermark`); value.Exists() {
		data.IPinWatermark = tfutils.BoolFromString(value.String())
	} else {
		data.IPinWatermark = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CookieGNVC`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CookieGnvc = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CookieGnvc = types.StringNull()
	}
	if value := res.Get(pathRoot + `UseSharedSecretObject`); value.Exists() {
		data.UseSharedSecretObject = tfutils.BoolFromString(value.String())
	} else {
		data.UseSharedSecretObject = types.BoolNull()
	}
}

func (data *DmCookieProfile) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `CookiePolicy`); value.Exists() && !data.CookiePolicy.IsNull() {
		data.CookiePolicy = tfutils.ParseStringFromGJSON(value)
	} else if data.CookiePolicy.ValueString() != "allow" {
		data.CookiePolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && !data.Type.IsNull() {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else if data.Type.ValueString() != "none" {
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `Key`); value.Exists() && !data.Key.IsNull() {
		data.Key = tfutils.ParseStringFromGJSON(value)
	} else if data.Key.ValueString() != "secret" {
		data.Key = types.StringNull()
	}
	if value := res.Get(pathRoot + `IPinWatermark`); value.Exists() && !data.IPinWatermark.IsNull() {
		data.IPinWatermark = tfutils.BoolFromString(value.String())
	} else if !data.IPinWatermark.ValueBool() {
		data.IPinWatermark = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CookieGNVC`); value.Exists() && !data.CookieGnvc.IsNull() {
		data.CookieGnvc = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CookieGnvc = types.StringNull()
	}
	if value := res.Get(pathRoot + `UseSharedSecretObject`); value.Exists() && !data.UseSharedSecretObject.IsNull() {
		data.UseSharedSecretObject = tfutils.BoolFromString(value.String())
	} else if data.UseSharedSecretObject.ValueBool() {
		data.UseSharedSecretObject = types.BoolNull()
	}
}
