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

type DmAPIResponseSchema struct {
	StatusCode     types.String `tfsdk:"status_code"`
	ResponseSchema types.String `tfsdk:"response_schema"`
}

var DmAPIResponseSchemaObjectType = map[string]attr.Type{
	"status_code":     types.StringType,
	"response_schema": types.StringType,
}
var DmAPIResponseSchemaObjectDefault = map[string]attr.Value{
	"status_code":     types.StringNull(),
	"response_schema": types.StringNull(),
}
var DmAPIResponseSchemaDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"status_code": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Status code", "", "").String,
			Computed:            true,
		},
		"response_schema": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Schema", "", "apischema").String,
			Computed:            true,
		},
	},
}
var DmAPIResponseSchemaResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"status_code": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Status code", "", "").String,
			Required:            true,
		},
		"response_schema": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Schema", "", "apischema").String,
			Optional:            true,
		},
	},
}

func (data DmAPIResponseSchema) IsNull() bool {
	if !data.StatusCode.IsNull() {
		return false
	}
	if !data.ResponseSchema.IsNull() {
		return false
	}
	return true
}

func (data DmAPIResponseSchema) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.StatusCode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StatusCode`, data.StatusCode.ValueString())
	}
	if !data.ResponseSchema.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ResponseSchema`, data.ResponseSchema.ValueString())
	}
	return body
}

func (data *DmAPIResponseSchema) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `StatusCode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.StatusCode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.StatusCode = types.StringNull()
	}
	if value := res.Get(pathRoot + `ResponseSchema`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ResponseSchema = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ResponseSchema = types.StringNull()
	}
}

func (data *DmAPIResponseSchema) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `StatusCode`); value.Exists() && !data.StatusCode.IsNull() {
		data.StatusCode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.StatusCode = types.StringNull()
	}
	if value := res.Get(pathRoot + `ResponseSchema`); value.Exists() && !data.ResponseSchema.IsNull() {
		data.ResponseSchema = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ResponseSchema = types.StringNull()
	}
}
