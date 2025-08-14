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

type DmAPIDataTypeDefinition struct {
	Name   types.String `tfsdk:"name"`
	Schema types.String `tfsdk:"schema"`
}

var DmAPIDataTypeDefinitionObjectType = map[string]attr.Type{
	"name":   types.StringType,
	"schema": types.StringType,
}
var DmAPIDataTypeDefinitionObjectDefault = map[string]attr.Value{
	"name":   types.StringNull(),
	"schema": types.StringNull(),
}
var DmAPIDataTypeDefinitionDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the data type.", "", "").String,
			Computed:            true,
		},
		"schema": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the API schema object.", "", "apischema").String,
			Computed:            true,
		},
	},
}
var DmAPIDataTypeDefinitionResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the data type.", "", "").String,
			Required:            true,
		},
		"schema": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the API schema object.", "", "apischema").String,
			Required:            true,
		},
	},
}

func (data DmAPIDataTypeDefinition) IsNull() bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.Schema.IsNull() {
		return false
	}
	return true
}

func (data DmAPIDataTypeDefinition) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Name`, data.Name.ValueString())
	}
	if !data.Schema.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Schema`, data.Schema.ValueString())
	}
	return body
}

func (data *DmAPIDataTypeDefinition) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Schema`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Schema = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Schema = types.StringNull()
	}
}

func (data *DmAPIDataTypeDefinition) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && !data.Name.IsNull() {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Schema`); value.Exists() && !data.Schema.IsNull() {
		data.Schema = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Schema = types.StringNull()
	}
}
