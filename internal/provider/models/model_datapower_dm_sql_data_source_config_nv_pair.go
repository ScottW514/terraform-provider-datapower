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

type DmSQLDataSourceConfigNVPair struct {
	ParamName  types.String `tfsdk:"param_name"`
	ParamValue types.String `tfsdk:"param_value"`
}

var DmSQLDataSourceConfigNVPairObjectType = map[string]attr.Type{
	"param_name":  types.StringType,
	"param_value": types.StringType,
}
var DmSQLDataSourceConfigNVPairObjectDefault = map[string]attr.Value{
	"param_name":  types.StringNull(),
	"param_value": types.StringNull(),
}

func GetDmSQLDataSourceConfigNVPairDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmSQLDataSourceConfigNVPairDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"param_name": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Parameter name",
				Computed:            true,
			},
			"param_value": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Parameter value",
				Computed:            true,
			},
		},
	}
	return DmSQLDataSourceConfigNVPairDataSourceSchema
}
func GetDmSQLDataSourceConfigNVPairResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmSQLDataSourceConfigNVPairResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"param_name": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Parameter name", "", "").String,
				Required:            true,
			},
			"param_value": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Parameter value", "", "").String,
				Required:            true,
			},
		},
	}
	return DmSQLDataSourceConfigNVPairResourceSchema
}

func (data DmSQLDataSourceConfigNVPair) IsNull() bool {
	if !data.ParamName.IsNull() {
		return false
	}
	if !data.ParamValue.IsNull() {
		return false
	}
	return true
}

func (data DmSQLDataSourceConfigNVPair) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.ParamName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ParamName`, data.ParamName.ValueString())
	}
	if !data.ParamValue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ParamValue`, data.ParamValue.ValueString())
	}
	return body
}

func (data *DmSQLDataSourceConfigNVPair) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `ParamName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ParamName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ParamName = types.StringNull()
	}
	if value := res.Get(pathRoot + `ParamValue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ParamValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ParamValue = types.StringNull()
	}
}

func (data *DmSQLDataSourceConfigNVPair) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `ParamName`); value.Exists() && !data.ParamName.IsNull() {
		data.ParamName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ParamName = types.StringNull()
	}
	if value := res.Get(pathRoot + `ParamValue`); value.Exists() && !data.ParamValue.IsNull() {
		data.ParamValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ParamValue = types.StringNull()
	}
}
