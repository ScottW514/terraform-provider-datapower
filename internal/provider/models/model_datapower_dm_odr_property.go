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

type DmODRProperty struct {
	PropName  types.String `tfsdk:"prop_name"`
	PropValue types.String `tfsdk:"prop_value"`
}

var DmODRPropertyObjectType = map[string]attr.Type{
	"prop_name":  types.StringType,
	"prop_value": types.StringType,
}
var DmODRPropertyObjectDefault = map[string]attr.Value{
	"prop_name":  types.StringNull(),
	"prop_value": types.StringNull(),
}

func GetDmODRPropertyDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmODRPropertyDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"prop_name": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of a property.", "odr-prop-name", "").String,
				Computed:            true,
			},
			"prop_value": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the value of a property.", "odr-prop-value", "").String,
				Computed:            true,
			},
		},
	}
	return DmODRPropertyDataSourceSchema
}
func GetDmODRPropertyResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmODRPropertyResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"prop_name": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of a property.", "odr-prop-name", "").String,
				Required:            true,
			},
			"prop_value": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the value of a property.", "odr-prop-value", "").String,
				Required:            true,
			},
		},
	}
	return DmODRPropertyResourceSchema
}

func (data DmODRProperty) IsNull() bool {
	if !data.PropName.IsNull() {
		return false
	}
	if !data.PropValue.IsNull() {
		return false
	}
	return true
}

func (data DmODRProperty) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.PropName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PropName`, data.PropName.ValueString())
	}
	if !data.PropValue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PropValue`, data.PropValue.ValueString())
	}
	return body
}

func (data *DmODRProperty) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `PropName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PropName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PropName = types.StringNull()
	}
	if value := res.Get(pathRoot + `PropValue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PropValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PropValue = types.StringNull()
	}
}

func (data *DmODRProperty) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `PropName`); value.Exists() && !data.PropName.IsNull() {
		data.PropName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PropName = types.StringNull()
	}
	if value := res.Get(pathRoot + `PropValue`); value.Exists() && !data.PropValue.IsNull() {
		data.PropValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PropValue = types.StringNull()
	}
}
