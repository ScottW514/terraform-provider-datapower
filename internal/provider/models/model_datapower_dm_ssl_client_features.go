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

type DmSSLClientFeatures struct {
	UseSni                types.Bool `tfsdk:"use_sni"`
	PermitInsecureServers types.Bool `tfsdk:"permit_insecure_servers"`
	Compression           types.Bool `tfsdk:"compression"`
}

var DmSSLClientFeaturesObjectType = map[string]attr.Type{
	"use_sni":                 types.BoolType,
	"permit_insecure_servers": types.BoolType,
	"compression":             types.BoolType,
}
var DmSSLClientFeaturesObjectDefault = map[string]attr.Value{
	"use_sni":                 types.BoolValue(false),
	"permit_insecure_servers": types.BoolValue(false),
	"compression":             types.BoolValue(false),
}

func GetDmSSLClientFeaturesDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmSSLClientFeaturesDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"use_sni": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use SNI", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"permit_insecure_servers": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Permit connections without renegotiation", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"compression": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable compression", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
		},
	}
	DmSSLClientFeaturesDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmSSLClientFeaturesDataSourceSchema
}
func GetDmSSLClientFeaturesResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmSSLClientFeaturesResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmSSLClientFeaturesObjectType,
				DmSSLClientFeaturesObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"use_sni": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use SNI", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"permit_insecure_servers": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Permit connections without renegotiation", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"compression": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable compression", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
	DmSSLClientFeaturesResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmSSLClientFeaturesResourceSchema.Required = true
	} else {
		DmSSLClientFeaturesResourceSchema.Optional = true
		DmSSLClientFeaturesResourceSchema.Computed = true
	}
	return DmSSLClientFeaturesResourceSchema
}

func (data DmSSLClientFeatures) IsNull() bool {
	if !data.UseSni.IsNull() {
		return false
	}
	if !data.PermitInsecureServers.IsNull() {
		return false
	}
	if !data.Compression.IsNull() {
		return false
	}
	return true
}

func (data DmSSLClientFeatures) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.UseSni.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`use-sni`, tfutils.StringFromBool(data.UseSni, ""))
	}
	if !data.PermitInsecureServers.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`permit-insecure-servers`, tfutils.StringFromBool(data.PermitInsecureServers, ""))
	}
	if !data.Compression.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`compression`, tfutils.StringFromBool(data.Compression, ""))
	}
	return body
}

func (data *DmSSLClientFeatures) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `use-sni`); value.Exists() {
		data.UseSni = tfutils.BoolFromString(value.String())
	} else {
		data.UseSni = types.BoolNull()
	}
	if value := res.Get(pathRoot + `permit-insecure-servers`); value.Exists() {
		data.PermitInsecureServers = tfutils.BoolFromString(value.String())
	} else {
		data.PermitInsecureServers = types.BoolNull()
	}
	if value := res.Get(pathRoot + `compression`); value.Exists() {
		data.Compression = tfutils.BoolFromString(value.String())
	} else {
		data.Compression = types.BoolNull()
	}
}

func (data *DmSSLClientFeatures) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `use-sni`); value.Exists() && !data.UseSni.IsNull() {
		data.UseSni = tfutils.BoolFromString(value.String())
	} else if data.UseSni.ValueBool() {
		data.UseSni = types.BoolNull()
	}
	if value := res.Get(pathRoot + `permit-insecure-servers`); value.Exists() && !data.PermitInsecureServers.IsNull() {
		data.PermitInsecureServers = tfutils.BoolFromString(value.String())
	} else if data.PermitInsecureServers.ValueBool() {
		data.PermitInsecureServers = types.BoolNull()
	}
	if value := res.Get(pathRoot + `compression`); value.Exists() && !data.Compression.IsNull() {
		data.Compression = tfutils.BoolFromString(value.String())
	} else if data.Compression.ValueBool() {
		data.Compression = types.BoolNull()
	}
}
