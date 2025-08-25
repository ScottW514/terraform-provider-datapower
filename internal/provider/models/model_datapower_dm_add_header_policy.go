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

type DmAddHeaderPolicy struct {
	RegExp    types.String `tfsdk:"reg_exp"`
	AddHeader types.String `tfsdk:"add_header"`
	AddValue  types.String `tfsdk:"add_value"`
}

var DmAddHeaderPolicyObjectType = map[string]attr.Type{
	"reg_exp":    types.StringType,
	"add_header": types.StringType,
	"add_value":  types.StringType,
}
var DmAddHeaderPolicyObjectDefault = map[string]attr.Value{
	"reg_exp":    types.StringNull(),
	"add_header": types.StringNull(),
	"add_value":  types.StringNull(),
}

func GetDmAddHeaderPolicyDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmAddHeaderPolicyDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"reg_exp": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the shell-style expression to define the URL set.", "", "").String,
				Computed:            true,
			},
			"add_header": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the header.", "", "").String,
				Computed:            true,
			},
			"add_value": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the value for the header as a string.", "", "").String,
				Computed:            true,
			},
		},
	}
	return DmAddHeaderPolicyDataSourceSchema
}
func GetDmAddHeaderPolicyResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmAddHeaderPolicyResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"reg_exp": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the shell-style expression to define the URL set.", "", "").String,
				Required:            true,
			},
			"add_header": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the header.", "", "").String,
				Required:            true,
			},
			"add_value": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the value for the header as a string.", "", "").String,
				Required:            true,
			},
		},
	}
	return DmAddHeaderPolicyResourceSchema
}

func (data DmAddHeaderPolicy) IsNull() bool {
	if !data.RegExp.IsNull() {
		return false
	}
	if !data.AddHeader.IsNull() {
		return false
	}
	if !data.AddValue.IsNull() {
		return false
	}
	return true
}

func (data DmAddHeaderPolicy) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.RegExp.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RegExp`, data.RegExp.ValueString())
	}
	if !data.AddHeader.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AddHeader`, data.AddHeader.ValueString())
	}
	if !data.AddValue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AddValue`, data.AddValue.ValueString())
	}
	return body
}

func (data *DmAddHeaderPolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `RegExp`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RegExp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RegExp = types.StringNull()
	}
	if value := res.Get(pathRoot + `AddHeader`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AddHeader = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AddHeader = types.StringNull()
	}
	if value := res.Get(pathRoot + `AddValue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AddValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AddValue = types.StringNull()
	}
}

func (data *DmAddHeaderPolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `RegExp`); value.Exists() && !data.RegExp.IsNull() {
		data.RegExp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RegExp = types.StringNull()
	}
	if value := res.Get(pathRoot + `AddHeader`); value.Exists() && !data.AddHeader.IsNull() {
		data.AddHeader = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AddHeader = types.StringNull()
	}
	if value := res.Get(pathRoot + `AddValue`); value.Exists() && !data.AddValue.IsNull() {
		data.AddValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AddValue = types.StringNull()
	}
}
