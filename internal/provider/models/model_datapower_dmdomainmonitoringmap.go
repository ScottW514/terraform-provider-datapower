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

type DmDomainMonitoringMap struct {
	Audit types.Bool `tfsdk:"audit"`
	Log   types.Bool `tfsdk:"log"`
}

var DmDomainMonitoringMapObjectType = map[string]attr.Type{
	"audit": types.BoolType,
	"log":   types.BoolType,
}
var DmDomainMonitoringMapObjectDefault = map[string]attr.Value{
	"audit": types.BoolValue(false),
	"log":   types.BoolValue(false),
}
var DmDomainMonitoringMapDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"audit": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enable auditing", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"log": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enable logging", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
	},
}
var DmDomainMonitoringMapResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmDomainMonitoringMapObjectType,
			DmDomainMonitoringMapObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"audit": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enable auditing", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"log": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enable logging", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
	},
}

func (data DmDomainMonitoringMap) IsNull() bool {
	if !data.Audit.IsNull() {
		return false
	}
	if !data.Log.IsNull() {
		return false
	}
	return true
}
func GetDmDomainMonitoringMapDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmDomainMonitoringMapDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmDomainMonitoringMapDataSourceSchema
}

func GetDmDomainMonitoringMapResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmDomainMonitoringMapResourceSchema.Required = true
	} else {
		DmDomainMonitoringMapResourceSchema.Optional = true
		DmDomainMonitoringMapResourceSchema.Computed = true
	}
	DmDomainMonitoringMapResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmDomainMonitoringMapResourceSchema
}

func (data DmDomainMonitoringMap) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Audit.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Audit`, tfutils.StringFromBool(data.Audit))
	}
	if !data.Log.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Log`, tfutils.StringFromBool(data.Log))
	}
	return body
}

func (data *DmDomainMonitoringMap) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Audit`); value.Exists() {
		data.Audit = tfutils.BoolFromString(value.String())
	} else {
		data.Audit = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Log`); value.Exists() {
		data.Log = tfutils.BoolFromString(value.String())
	} else {
		data.Log = types.BoolNull()
	}
}

func (data *DmDomainMonitoringMap) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Audit`); value.Exists() && !data.Audit.IsNull() {
		data.Audit = tfutils.BoolFromString(value.String())
	} else if data.Audit.ValueBool() {
		data.Audit = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Log`); value.Exists() && !data.Log.IsNull() {
		data.Log = tfutils.BoolFromString(value.String())
	} else if data.Log.ValueBool() {
		data.Log = types.BoolNull()
	}
}
