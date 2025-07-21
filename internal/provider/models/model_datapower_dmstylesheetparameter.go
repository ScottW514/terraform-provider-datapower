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

type DmStylesheetParameter struct {
	ParameterName  types.String `tfsdk:"parameter_name"`
	ParameterValue types.String `tfsdk:"parameter_value"`
}

var DmStylesheetParameterObjectType = map[string]attr.Type{
	"parameter_name":  types.StringType,
	"parameter_value": types.StringType,
}
var DmStylesheetParameterObjectDefault = map[string]attr.Value{
	"parameter_name":  types.StringNull(),
	"parameter_value": types.StringNull(),
}
var DmStylesheetParameterDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"parameter_name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Parameter Name", "", "").String,
			Computed:            true,
		},
		"parameter_value": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Parameter Value", "", "").String,
			Computed:            true,
		},
	},
}
var DmStylesheetParameterResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"parameter_name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Parameter Name", "", "").String,
			Optional:            true,
		},
		"parameter_value": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Parameter Value", "", "").String,
			Required:            true,
		},
	},
}

func (data DmStylesheetParameter) IsNull() bool {
	if !data.ParameterName.IsNull() {
		return false
	}
	if !data.ParameterValue.IsNull() {
		return false
	}
	return true
}

func (data DmStylesheetParameter) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.ParameterName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ParameterName`, data.ParameterName.ValueString())
	}
	if !data.ParameterValue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ParameterValue`, data.ParameterValue.ValueString())
	}
	return body
}

func (data *DmStylesheetParameter) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `ParameterName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ParameterName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ParameterName = types.StringNull()
	}
	if value := res.Get(pathRoot + `ParameterValue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ParameterValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ParameterValue = types.StringNull()
	}
}

func (data *DmStylesheetParameter) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `ParameterName`); value.Exists() && !data.ParameterName.IsNull() {
		data.ParameterName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ParameterName = types.StringNull()
	}
	if value := res.Get(pathRoot + `ParameterValue`); value.Exists() && !data.ParameterValue.IsNull() {
		data.ParameterValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ParameterValue = types.StringNull()
	}
}
