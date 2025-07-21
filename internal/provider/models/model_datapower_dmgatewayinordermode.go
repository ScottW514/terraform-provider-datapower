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

type DmGatewayInOrderMode struct {
	Request  types.Bool `tfsdk:"request"`
	Backend  types.Bool `tfsdk:"backend"`
	Response types.Bool `tfsdk:"response"`
}

var DmGatewayInOrderModeObjectType = map[string]attr.Type{
	"request":  types.BoolType,
	"backend":  types.BoolType,
	"response": types.BoolType,
}
var DmGatewayInOrderModeObjectDefault = map[string]attr.Value{
	"request":  types.BoolValue(false),
	"backend":  types.BoolValue(false),
	"response": types.BoolValue(false),
}
var DmGatewayInOrderModeDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"request": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Request rule in order", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"backend": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Backend in order", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"response": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Response rule in order", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
	},
}
var DmGatewayInOrderModeResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmGatewayInOrderModeObjectType,
			DmGatewayInOrderModeObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"request": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Request rule in order", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"backend": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Backend in order", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"response": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Response rule in order", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
	},
}

func (data DmGatewayInOrderMode) IsNull() bool {
	if !data.Request.IsNull() {
		return false
	}
	if !data.Backend.IsNull() {
		return false
	}
	if !data.Response.IsNull() {
		return false
	}
	return true
}
func GetDmGatewayInOrderModeDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmGatewayInOrderModeDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmGatewayInOrderModeDataSourceSchema
}

func GetDmGatewayInOrderModeResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmGatewayInOrderModeResourceSchema.Required = true
	} else {
		DmGatewayInOrderModeResourceSchema.Optional = true
		DmGatewayInOrderModeResourceSchema.Computed = true
	}
	DmGatewayInOrderModeResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmGatewayInOrderModeResourceSchema
}

func (data DmGatewayInOrderMode) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Request.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Request`, tfutils.StringFromBool(data.Request))
	}
	if !data.Backend.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Backend`, tfutils.StringFromBool(data.Backend))
	}
	if !data.Response.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Response`, tfutils.StringFromBool(data.Response))
	}
	return body
}

func (data *DmGatewayInOrderMode) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Request`); value.Exists() {
		data.Request = tfutils.BoolFromString(value.String())
	} else {
		data.Request = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Backend`); value.Exists() {
		data.Backend = tfutils.BoolFromString(value.String())
	} else {
		data.Backend = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Response`); value.Exists() {
		data.Response = tfutils.BoolFromString(value.String())
	} else {
		data.Response = types.BoolNull()
	}
}

func (data *DmGatewayInOrderMode) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Request`); value.Exists() && !data.Request.IsNull() {
		data.Request = tfutils.BoolFromString(value.String())
	} else if data.Request.ValueBool() {
		data.Request = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Backend`); value.Exists() && !data.Backend.IsNull() {
		data.Backend = tfutils.BoolFromString(value.String())
	} else if data.Backend.ValueBool() {
		data.Backend = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Response`); value.Exists() && !data.Response.IsNull() {
		data.Response = tfutils.BoolFromString(value.String())
	} else if data.Response.ValueBool() {
		data.Response = types.BoolNull()
	}
}
