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

type DmB2BMessageProperties struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
	Type  types.String `tfsdk:"type"`
}

var DmB2BMessagePropertiesObjectType = map[string]attr.Type{
	"name":  types.StringType,
	"value": types.StringType,
	"type":  types.StringType,
}
var DmB2BMessagePropertiesObjectDefault = map[string]attr.Value{
	"name":  types.StringNull(),
	"value": types.StringNull(),
	"type":  types.StringNull(),
}
var DmB2BMessagePropertiesDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Name", "", "").String,
			Computed:            true,
		},
		"value": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Value", "", "").String,
			Computed:            true,
		},
		"type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Type", "", "").String,
			Computed:            true,
		},
	},
}
var DmB2BMessagePropertiesResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Name", "", "").String,
			Required:            true,
		},
		"value": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Value", "", "").String,
			Required:            true,
		},
		"type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Type", "", "").String,
			Optional:            true,
		},
	},
}

func (data DmB2BMessageProperties) IsNull() bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.Value.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
		return false
	}
	return true
}

func (data DmB2BMessageProperties) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Name`, data.Name.ValueString())
	}
	if !data.Value.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Value`, data.Value.ValueString())
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Type`, data.Type.ValueString())
	}
	return body
}

func (data *DmB2BMessageProperties) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Value`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Value = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Value = types.StringNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringNull()
	}
}

func (data *DmB2BMessageProperties) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && !data.Name.IsNull() {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Value`); value.Exists() && !data.Value.IsNull() {
		data.Value = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Value = types.StringNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && !data.Type.IsNull() {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringNull()
	}
}
