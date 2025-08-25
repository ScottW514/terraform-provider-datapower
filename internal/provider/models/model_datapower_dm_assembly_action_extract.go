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

type DmAssemblyActionExtract struct {
	Capture   types.String `tfsdk:"capture"`
	Transform types.String `tfsdk:"transform"`
}

var DmAssemblyActionExtractObjectType = map[string]attr.Type{
	"capture":   types.StringType,
	"transform": types.StringType,
}
var DmAssemblyActionExtractObjectDefault = map[string]attr.Value{
	"capture":   types.StringNull(),
	"transform": types.StringNull(),
}

func GetDmAssemblyActionExtractDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmAssemblyActionExtractDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"capture": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the path expression that identifies the field.", "capture", "").String,
				Computed:            true,
			},
			"transform": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the expression that defines how to transform the content.", "transform", "").String,
				Computed:            true,
			},
		},
	}
	return DmAssemblyActionExtractDataSourceSchema
}
func GetDmAssemblyActionExtractResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmAssemblyActionExtractResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"capture": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the path expression that identifies the field.", "capture", "").String,
				Required:            true,
			},
			"transform": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the expression that defines how to transform the content.", "transform", "").String,
				Optional:            true,
			},
		},
	}
	return DmAssemblyActionExtractResourceSchema
}

func (data DmAssemblyActionExtract) IsNull() bool {
	if !data.Capture.IsNull() {
		return false
	}
	if !data.Transform.IsNull() {
		return false
	}
	return true
}

func (data DmAssemblyActionExtract) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Capture.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Capture`, data.Capture.ValueString())
	}
	if !data.Transform.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Transform`, data.Transform.ValueString())
	}
	return body
}

func (data *DmAssemblyActionExtract) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Capture`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Capture = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Capture = types.StringNull()
	}
	if value := res.Get(pathRoot + `Transform`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Transform = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Transform = types.StringNull()
	}
}

func (data *DmAssemblyActionExtract) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Capture`); value.Exists() && !data.Capture.IsNull() {
		data.Capture = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Capture = types.StringNull()
	}
	if value := res.Get(pathRoot + `Transform`); value.Exists() && !data.Transform.IsNull() {
		data.Transform = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Transform = types.StringNull()
	}
}
