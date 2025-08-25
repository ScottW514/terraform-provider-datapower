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

type DmAPIProtocols struct {
	Http  types.Bool `tfsdk:"http"`
	Https types.Bool `tfsdk:"https"`
	Ws    types.Bool `tfsdk:"ws"`
	Wss   types.Bool `tfsdk:"wss"`
}

var DmAPIProtocolsObjectType = map[string]attr.Type{
	"http":  types.BoolType,
	"https": types.BoolType,
	"ws":    types.BoolType,
	"wss":   types.BoolType,
}
var DmAPIProtocolsObjectDefault = map[string]attr.Value{
	"http":  types.BoolValue(true),
	"https": types.BoolValue(true),
	"ws":    types.BoolValue(true),
	"wss":   types.BoolValue(true),
}

func GetDmAPIProtocolsDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmAPIProtocolsDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"http": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Protocol identifier for HTTP.", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"https": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Protocol identifier for HTTPS.", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"ws": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Protocol identifier for WebSocket. When allowed, the HTTP handler must be configured to allow WebSocket upgrade requests.", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"wss": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Protocol identifier for WebSocket Secure. When allowed, the HTTPS handler must be configured to allow WebSocket upgrade requests.", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
		},
	}
	DmAPIProtocolsDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmAPIProtocolsDataSourceSchema
}
func GetDmAPIProtocolsResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmAPIProtocolsResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmAPIProtocolsObjectType,
				DmAPIProtocolsObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"http": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Protocol identifier for HTTP.", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"https": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Protocol identifier for HTTPS.", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ws": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Protocol identifier for WebSocket. When allowed, the HTTP handler must be configured to allow WebSocket upgrade requests.", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"wss": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Protocol identifier for WebSocket Secure. When allowed, the HTTPS handler must be configured to allow WebSocket upgrade requests.", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
		},
	}
	DmAPIProtocolsResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmAPIProtocolsResourceSchema.Required = true
	} else {
		DmAPIProtocolsResourceSchema.Optional = true
		DmAPIProtocolsResourceSchema.Computed = true
	}
	return DmAPIProtocolsResourceSchema
}

func (data DmAPIProtocols) IsNull() bool {
	if !data.Http.IsNull() {
		return false
	}
	if !data.Https.IsNull() {
		return false
	}
	if !data.Ws.IsNull() {
		return false
	}
	if !data.Wss.IsNull() {
		return false
	}
	return true
}

func (data DmAPIProtocols) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Http.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`http`, tfutils.StringFromBool(data.Http, ""))
	}
	if !data.Https.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`https`, tfutils.StringFromBool(data.Https, ""))
	}
	if !data.Ws.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ws`, tfutils.StringFromBool(data.Ws, ""))
	}
	if !data.Wss.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`wss`, tfutils.StringFromBool(data.Wss, ""))
	}
	return body
}

func (data *DmAPIProtocols) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `http`); value.Exists() {
		data.Http = tfutils.BoolFromString(value.String())
	} else {
		data.Http = types.BoolNull()
	}
	if value := res.Get(pathRoot + `https`); value.Exists() {
		data.Https = tfutils.BoolFromString(value.String())
	} else {
		data.Https = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ws`); value.Exists() {
		data.Ws = tfutils.BoolFromString(value.String())
	} else {
		data.Ws = types.BoolNull()
	}
	if value := res.Get(pathRoot + `wss`); value.Exists() {
		data.Wss = tfutils.BoolFromString(value.String())
	} else {
		data.Wss = types.BoolNull()
	}
}

func (data *DmAPIProtocols) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `http`); value.Exists() && !data.Http.IsNull() {
		data.Http = tfutils.BoolFromString(value.String())
	} else if !data.Http.ValueBool() {
		data.Http = types.BoolNull()
	}
	if value := res.Get(pathRoot + `https`); value.Exists() && !data.Https.IsNull() {
		data.Https = tfutils.BoolFromString(value.String())
	} else if !data.Https.ValueBool() {
		data.Https = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ws`); value.Exists() && !data.Ws.IsNull() {
		data.Ws = tfutils.BoolFromString(value.String())
	} else if !data.Ws.ValueBool() {
		data.Ws = types.BoolNull()
	}
	if value := res.Get(pathRoot + `wss`); value.Exists() && !data.Wss.IsNull() {
		data.Wss = tfutils.BoolFromString(value.String())
	} else if !data.Wss.ValueBool() {
		data.Wss = types.BoolNull()
	}
}
