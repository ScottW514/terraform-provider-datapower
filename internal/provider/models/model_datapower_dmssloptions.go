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

type DmSSLOptions struct {
	MaxDuration      types.Bool `tfsdk:"max_duration"`
	MaxRenegotiation types.Bool `tfsdk:"max_renegotiation"`
}

var DmSSLOptionsObjectType = map[string]attr.Type{
	"max_duration":      types.BoolType,
	"max_renegotiation": types.BoolType,
}
var DmSSLOptionsObjectDefault = map[string]attr.Value{
	"max_duration":      types.BoolValue(false),
	"max_renegotiation": types.BoolValue(false),
}
var DmSSLOptionsDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"max_duration": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Set maximum TLS session duration", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"max_renegotiation": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Set maximum number client initiated renegotiation allow", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
	},
}
var DmSSLOptionsResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmSSLOptionsObjectType,
			DmSSLOptionsObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"max_duration": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Set maximum TLS session duration", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"max_renegotiation": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Set maximum number client initiated renegotiation allow", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
	},
}

func (data DmSSLOptions) IsNull() bool {
	if !data.MaxDuration.IsNull() {
		return false
	}
	if !data.MaxRenegotiation.IsNull() {
		return false
	}
	return true
}
func GetDmSSLOptionsDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmSSLOptionsDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmSSLOptionsDataSourceSchema
}

func GetDmSSLOptionsResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmSSLOptionsResourceSchema.Required = true
	} else {
		DmSSLOptionsResourceSchema.Optional = true
		DmSSLOptionsResourceSchema.Computed = true
	}
	DmSSLOptionsResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmSSLOptionsResourceSchema
}

func (data DmSSLOptions) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.MaxDuration.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`max-duration`, tfutils.StringFromBool(data.MaxDuration, false))
	}
	if !data.MaxRenegotiation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`max-renegotiation`, tfutils.StringFromBool(data.MaxRenegotiation, false))
	}
	return body
}

func (data *DmSSLOptions) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `max-duration`); value.Exists() {
		data.MaxDuration = tfutils.BoolFromString(value.String())
	} else {
		data.MaxDuration = types.BoolNull()
	}
	if value := res.Get(pathRoot + `max-renegotiation`); value.Exists() {
		data.MaxRenegotiation = tfutils.BoolFromString(value.String())
	} else {
		data.MaxRenegotiation = types.BoolNull()
	}
}

func (data *DmSSLOptions) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `max-duration`); value.Exists() && !data.MaxDuration.IsNull() {
		data.MaxDuration = tfutils.BoolFromString(value.String())
	} else if data.MaxDuration.ValueBool() {
		data.MaxDuration = types.BoolNull()
	}
	if value := res.Get(pathRoot + `max-renegotiation`); value.Exists() && !data.MaxRenegotiation.IsNull() {
		data.MaxRenegotiation = tfutils.BoolFromString(value.String())
	} else if data.MaxRenegotiation.ValueBool() {
		data.MaxRenegotiation = types.BoolNull()
	}
}
