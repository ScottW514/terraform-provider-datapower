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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	DataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	ResourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmConfigModifyType struct {
	Match    types.String `tfsdk:"match"`
	Type     types.String `tfsdk:"type"`
	Property types.String `tfsdk:"property"`
	Value    types.String `tfsdk:"value"`
}

var DmConfigModifyTypeObjectType = map[string]attr.Type{
	"match":    types.StringType,
	"type":     types.StringType,
	"property": types.StringType,
	"value":    types.StringType,
}
var DmConfigModifyTypeObjectDefault = map[string]attr.Value{
	"match":    types.StringNull(),
	"type":     types.StringNull(),
	"property": types.StringNull(),
	"value":    types.StringNull(),
}
var DmConfigModifyTypeDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"match": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Configuration Match", "match", "").String,
			Computed:            true,
		},
		"type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Modification Type", "type", "").AddStringEnum("add", "delete", "change").String,
			Computed:            true,
		},
		"property": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Configuration Property", "name", "").String,
			Computed:            true,
		},
		"value": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Configuration Value", "value", "").String,
			Computed:            true,
		},
	},
}
var DmConfigModifyTypeResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"match": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Configuration Match", "match", "").String,
			Required:            true,
		},
		"type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Modification Type", "type", "").AddStringEnum("add", "delete", "change").String,
			Required:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("add", "delete", "change"),
			},
		},
		"property": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Configuration Property", "name", "").String,
			Optional:            true,
		},
		"value": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Configuration Value", "value", "").String,
			Optional:            true,
		},
	},
}

func (data DmConfigModifyType) IsNull() bool {
	if !data.Match.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
		return false
	}
	if !data.Property.IsNull() {
		return false
	}
	if !data.Value.IsNull() {
		return false
	}
	return true
}

func (data DmConfigModifyType) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Match.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Match`, data.Match.ValueString())
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Type`, data.Type.ValueString())
	}
	if !data.Property.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Property`, data.Property.ValueString())
	}
	if !data.Value.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Value`, data.Value.ValueString())
	}
	return body
}

func (data *DmConfigModifyType) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Match`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Match = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Match = types.StringNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `Property`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Property = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Property = types.StringNull()
	}
	if value := res.Get(pathRoot + `Value`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Value = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Value = types.StringNull()
	}
}

func (data *DmConfigModifyType) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Match`); value.Exists() && !data.Match.IsNull() {
		data.Match = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Match = types.StringNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && !data.Type.IsNull() {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `Property`); value.Exists() && !data.Property.IsNull() {
		data.Property = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Property = types.StringNull()
	}
	if value := res.Get(pathRoot + `Value`); value.Exists() && !data.Value.IsNull() {
		data.Value = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Value = types.StringNull()
	}
}
