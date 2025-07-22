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

type DmSMFlow struct {
	Frontend types.Bool `tfsdk:"frontend"`
	Backend  types.Bool `tfsdk:"backend"`
}

var DmSMFlowObjectType = map[string]attr.Type{
	"frontend": types.BoolType,
	"backend":  types.BoolType,
}
var DmSMFlowObjectDefault = map[string]attr.Value{
	"frontend": types.BoolValue(false),
	"backend":  types.BoolValue(false),
}
var DmSMFlowDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"frontend": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Responses", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"backend": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Requests", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
	},
}
var DmSMFlowResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmSMFlowObjectType,
			DmSMFlowObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"frontend": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Responses", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"backend": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Requests", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
	},
}

func (data DmSMFlow) IsNull() bool {
	if !data.Frontend.IsNull() {
		return false
	}
	if !data.Backend.IsNull() {
		return false
	}
	return true
}
func GetDmSMFlowDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmSMFlowDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmSMFlowDataSourceSchema
}

func GetDmSMFlowResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmSMFlowResourceSchema.Required = true
	} else {
		DmSMFlowResourceSchema.Optional = true
		DmSMFlowResourceSchema.Computed = true
	}
	DmSMFlowResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmSMFlowResourceSchema
}

func (data DmSMFlow) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Frontend.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`frontend`, tfutils.StringFromBool(data.Frontend, false))
	}
	if !data.Backend.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`backend`, tfutils.StringFromBool(data.Backend, false))
	}
	return body
}

func (data *DmSMFlow) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `frontend`); value.Exists() {
		data.Frontend = tfutils.BoolFromString(value.String())
	} else {
		data.Frontend = types.BoolNull()
	}
	if value := res.Get(pathRoot + `backend`); value.Exists() {
		data.Backend = tfutils.BoolFromString(value.String())
	} else {
		data.Backend = types.BoolNull()
	}
}

func (data *DmSMFlow) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `frontend`); value.Exists() && !data.Frontend.IsNull() {
		data.Frontend = tfutils.BoolFromString(value.String())
	} else if data.Frontend.ValueBool() {
		data.Frontend = types.BoolNull()
	}
	if value := res.Get(pathRoot + `backend`); value.Exists() && !data.Backend.IsNull() {
		data.Backend = tfutils.BoolFromString(value.String())
	} else if data.Backend.ValueBool() {
		data.Backend = types.BoolNull()
	}
}
