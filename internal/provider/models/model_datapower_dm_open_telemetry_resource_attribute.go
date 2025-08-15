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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmOpenTelemetryResourceAttribute struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

var DmOpenTelemetryResourceAttributeObjectType = map[string]attr.Type{
	"key":   types.StringType,
	"value": types.StringType,
}
var DmOpenTelemetryResourceAttributeObjectDefault = map[string]attr.Value{
	"key":   types.StringNull(),
	"value": types.StringNull(),
}
var DmOpenTelemetryResourceAttributeDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"key": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Key", "", "").String,
			Computed:            true,
		},
		"value": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Value", "", "").String,
			Computed:            true,
		},
	},
}
var DmOpenTelemetryResourceAttributeResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"key": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Key", "", "").String,
			Optional:            true,
		},
		"value": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Value", "", "").String,
			Required:            true,
		},
	},
}

func (data DmOpenTelemetryResourceAttribute) IsNull() bool {
	if !data.Key.IsNull() {
		return false
	}
	if !data.Value.IsNull() {
		return false
	}
	return true
}

func (data DmOpenTelemetryResourceAttribute) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Key.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Key`, data.Key.ValueString())
	}
	if !data.Value.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Value`, data.Value.ValueString())
	}
	return body
}

func (data *DmOpenTelemetryResourceAttribute) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Key`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Key = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Key = types.StringNull()
	}
	if value := res.Get(pathRoot + `Value`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Value = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Value = types.StringNull()
	}
}

func (data *DmOpenTelemetryResourceAttribute) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Key`); value.Exists() && !data.Key.IsNull() {
		data.Key = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Key = types.StringNull()
	}
	if value := res.Get(pathRoot + `Value`); value.Exists() && !data.Value.IsNull() {
		data.Value = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Value = types.StringNull()
	}
}
