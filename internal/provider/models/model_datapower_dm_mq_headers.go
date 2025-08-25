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

type DmMQHeaders struct {
	Mqcih  types.Bool `tfsdk:"mqcih"`
	Mqdlh  types.Bool `tfsdk:"mqdlh"`
	Mqiih  types.Bool `tfsdk:"mqiih"`
	Mqrfh  types.Bool `tfsdk:"mqrfh"`
	Mqrfh2 types.Bool `tfsdk:"mqrfh2"`
	Mqwih  types.Bool `tfsdk:"mqwih"`
}

var DmMQHeadersObjectType = map[string]attr.Type{
	"mqcih":  types.BoolType,
	"mqdlh":  types.BoolType,
	"mqiih":  types.BoolType,
	"mqrfh":  types.BoolType,
	"mqrfh2": types.BoolType,
	"mqwih":  types.BoolType,
}
var DmMQHeadersObjectDefault = map[string]attr.Value{
	"mqcih":  types.BoolValue(false),
	"mqdlh":  types.BoolValue(false),
	"mqiih":  types.BoolValue(false),
	"mqrfh":  types.BoolValue(false),
	"mqrfh2": types.BoolValue(false),
	"mqwih":  types.BoolValue(false),
}

func GetDmMQHeadersDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmMQHeadersDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"mqcih": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("CICS Bridge Header (MQCIH)", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"mqdlh": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Dead Letter Header (MQDLH)", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"mqiih": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("IMS Information Header (MQIIH)", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"mqrfh": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Rules and Formatting Header (MQRFH)", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"mqrfh2": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Rules and Formatting Header (MQRFH2)", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"mqwih": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Work Information Header (MQWIH)", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
		},
	}
	DmMQHeadersDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmMQHeadersDataSourceSchema
}
func GetDmMQHeadersResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmMQHeadersResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmMQHeadersObjectType,
				DmMQHeadersObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"mqcih": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("CICS Bridge Header (MQCIH)", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"mqdlh": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Dead Letter Header (MQDLH)", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"mqiih": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("IMS Information Header (MQIIH)", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"mqrfh": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Rules and Formatting Header (MQRFH)", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"mqrfh2": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Rules and Formatting Header (MQRFH2)", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"mqwih": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Work Information Header (MQWIH)", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
	DmMQHeadersResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmMQHeadersResourceSchema.Required = true
	} else {
		DmMQHeadersResourceSchema.Optional = true
		DmMQHeadersResourceSchema.Computed = true
	}
	return DmMQHeadersResourceSchema
}

func (data DmMQHeaders) IsNull() bool {
	if !data.Mqcih.IsNull() {
		return false
	}
	if !data.Mqdlh.IsNull() {
		return false
	}
	if !data.Mqiih.IsNull() {
		return false
	}
	if !data.Mqrfh.IsNull() {
		return false
	}
	if !data.Mqrfh2.IsNull() {
		return false
	}
	if !data.Mqwih.IsNull() {
		return false
	}
	return true
}

func (data DmMQHeaders) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Mqcih.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MQCIH`, tfutils.StringFromBool(data.Mqcih, ""))
	}
	if !data.Mqdlh.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MQDLH`, tfutils.StringFromBool(data.Mqdlh, ""))
	}
	if !data.Mqiih.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MQIIH`, tfutils.StringFromBool(data.Mqiih, ""))
	}
	if !data.Mqrfh.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MQRFH`, tfutils.StringFromBool(data.Mqrfh, ""))
	}
	if !data.Mqrfh2.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MQRFH2`, tfutils.StringFromBool(data.Mqrfh2, ""))
	}
	if !data.Mqwih.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MQWIH`, tfutils.StringFromBool(data.Mqwih, ""))
	}
	return body
}

func (data *DmMQHeaders) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `MQCIH`); value.Exists() {
		data.Mqcih = tfutils.BoolFromString(value.String())
	} else {
		data.Mqcih = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MQDLH`); value.Exists() {
		data.Mqdlh = tfutils.BoolFromString(value.String())
	} else {
		data.Mqdlh = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MQIIH`); value.Exists() {
		data.Mqiih = tfutils.BoolFromString(value.String())
	} else {
		data.Mqiih = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MQRFH`); value.Exists() {
		data.Mqrfh = tfutils.BoolFromString(value.String())
	} else {
		data.Mqrfh = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MQRFH2`); value.Exists() {
		data.Mqrfh2 = tfutils.BoolFromString(value.String())
	} else {
		data.Mqrfh2 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MQWIH`); value.Exists() {
		data.Mqwih = tfutils.BoolFromString(value.String())
	} else {
		data.Mqwih = types.BoolNull()
	}
}

func (data *DmMQHeaders) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `MQCIH`); value.Exists() && !data.Mqcih.IsNull() {
		data.Mqcih = tfutils.BoolFromString(value.String())
	} else if data.Mqcih.ValueBool() {
		data.Mqcih = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MQDLH`); value.Exists() && !data.Mqdlh.IsNull() {
		data.Mqdlh = tfutils.BoolFromString(value.String())
	} else if data.Mqdlh.ValueBool() {
		data.Mqdlh = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MQIIH`); value.Exists() && !data.Mqiih.IsNull() {
		data.Mqiih = tfutils.BoolFromString(value.String())
	} else if data.Mqiih.ValueBool() {
		data.Mqiih = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MQRFH`); value.Exists() && !data.Mqrfh.IsNull() {
		data.Mqrfh = tfutils.BoolFromString(value.String())
	} else if data.Mqrfh.ValueBool() {
		data.Mqrfh = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MQRFH2`); value.Exists() && !data.Mqrfh2.IsNull() {
		data.Mqrfh2 = tfutils.BoolFromString(value.String())
	} else if data.Mqrfh2.ValueBool() {
		data.Mqrfh2 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MQWIH`); value.Exists() && !data.Mqwih.IsNull() {
		data.Mqwih = tfutils.BoolFromString(value.String())
	} else if data.Mqwih.ValueBool() {
		data.Mqwih = types.BoolNull()
	}
}
