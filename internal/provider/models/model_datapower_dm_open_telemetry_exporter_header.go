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

type DmOpenTelemetryExporterHeader struct {
	HeaderName  types.String `tfsdk:"header_name"`
	HeaderValue types.String `tfsdk:"header_value"`
}

var DmOpenTelemetryExporterHeaderObjectType = map[string]attr.Type{
	"header_name":  types.StringType,
	"header_value": types.StringType,
}
var DmOpenTelemetryExporterHeaderObjectDefault = map[string]attr.Value{
	"header_name":  types.StringNull(),
	"header_value": types.StringNull(),
}
var DmOpenTelemetryExporterHeaderDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"header_name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Name", "", "").String,
			Computed:            true,
		},
		"header_value": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Value", "", "").String,
			Computed:            true,
		},
	},
}
var DmOpenTelemetryExporterHeaderResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"header_name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Name", "", "").String,
			Optional:            true,
		},
		"header_value": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Value", "", "").String,
			Required:            true,
		},
	},
}

func (data DmOpenTelemetryExporterHeader) IsNull() bool {
	if !data.HeaderName.IsNull() {
		return false
	}
	if !data.HeaderValue.IsNull() {
		return false
	}
	return true
}

func (data DmOpenTelemetryExporterHeader) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.HeaderName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`headerName`, data.HeaderName.ValueString())
	}
	if !data.HeaderValue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`headerValue`, data.HeaderValue.ValueString())
	}
	return body
}

func (data *DmOpenTelemetryExporterHeader) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `headerName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HeaderName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HeaderName = types.StringNull()
	}
	if value := res.Get(pathRoot + `headerValue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HeaderValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HeaderValue = types.StringNull()
	}
}

func (data *DmOpenTelemetryExporterHeader) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `headerName`); value.Exists() && !data.HeaderName.IsNull() {
		data.HeaderName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HeaderName = types.StringNull()
	}
	if value := res.Get(pathRoot + `headerValue`); value.Exists() && !data.HeaderValue.IsNull() {
		data.HeaderValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HeaderValue = types.StringNull()
	}
}
