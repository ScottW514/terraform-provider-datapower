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

type DmSAMLAttribute struct {
	SourceType   types.String `tfsdk:"source_type"`
	Name         types.String `tfsdk:"name"`
	Format       types.String `tfsdk:"format"`
	XPath        types.String `tfsdk:"x_path"`
	ValueData    types.String `tfsdk:"value_data"`
	SubValueData types.String `tfsdk:"sub_value_data"`
	FriendlyName types.String `tfsdk:"friendly_name"`
}

var DmSAMLAttributeObjectType = map[string]attr.Type{
	"source_type":    types.StringType,
	"name":           types.StringType,
	"format":         types.StringType,
	"x_path":         types.StringType,
	"value_data":     types.StringType,
	"sub_value_data": types.StringType,
	"friendly_name":  types.StringType,
}
var DmSAMLAttributeObjectDefault = map[string]attr.Value{
	"source_type":    types.StringNull(),
	"name":           types.StringNull(),
	"format":         types.StringNull(),
	"x_path":         types.StringNull(),
	"value_data":     types.StringNull(),
	"sub_value_data": types.StringNull(),
	"friendly_name":  types.StringNull(),
}
var DmSAMLAttributeDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"source_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Data Source Type", "", "").AddStringEnum("var", "xpath", "static").String,
			Computed:            true,
		},
		"name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Name", "", "").String,
			Computed:            true,
		},
		"format": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Attribute NameSpace/Format", "", "").String,
			Computed:            true,
		},
		"x_path": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("XPath to Input Message", "", "").String,
			Computed:            true,
		},
		"value_data": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Data for Attribute Value", "", "").String,
			Computed:            true,
		},
		"sub_value_data": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Supplementary Data", "", "").String,
			Computed:            true,
		},
		"friendly_name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Friendly Name", "", "").String,
			Computed:            true,
		},
	},
}
var DmSAMLAttributeResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"source_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Data Source Type", "", "").AddStringEnum("var", "xpath", "static").String,
			Required:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("var", "xpath", "static"),
			},
		},
		"name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Name", "", "").String,
			Optional:            true,
		},
		"format": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Attribute NameSpace/Format", "", "").String,
			Optional:            true,
		},
		"x_path": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("XPath to Input Message", "", "").String,
			Optional:            true,
		},
		"value_data": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Data for Attribute Value", "", "").String,
			Optional:            true,
		},
		"sub_value_data": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Supplementary Data", "", "").String,
			Optional:            true,
		},
		"friendly_name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Friendly Name", "", "").String,
			Optional:            true,
		},
	},
}

func (data DmSAMLAttribute) IsNull() bool {
	if !data.SourceType.IsNull() {
		return false
	}
	if !data.Name.IsNull() {
		return false
	}
	if !data.Format.IsNull() {
		return false
	}
	if !data.XPath.IsNull() {
		return false
	}
	if !data.ValueData.IsNull() {
		return false
	}
	if !data.SubValueData.IsNull() {
		return false
	}
	if !data.FriendlyName.IsNull() {
		return false
	}
	return true
}

func (data DmSAMLAttribute) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.SourceType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SourceType`, data.SourceType.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Name`, data.Name.ValueString())
	}
	if !data.Format.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Format`, data.Format.ValueString())
	}
	if !data.XPath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XPath`, data.XPath.ValueString())
	}
	if !data.ValueData.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ValueData`, data.ValueData.ValueString())
	}
	if !data.SubValueData.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SubValueData`, data.SubValueData.ValueString())
	}
	if !data.FriendlyName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FriendlyName`, data.FriendlyName.ValueString())
	}
	return body
}

func (data *DmSAMLAttribute) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `SourceType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SourceType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SourceType = types.StringNull()
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Format`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Format = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Format = types.StringNull()
	}
	if value := res.Get(pathRoot + `XPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.XPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValueData`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ValueData = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ValueData = types.StringNull()
	}
	if value := res.Get(pathRoot + `SubValueData`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SubValueData = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SubValueData = types.StringNull()
	}
	if value := res.Get(pathRoot + `FriendlyName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.FriendlyName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FriendlyName = types.StringNull()
	}
}

func (data *DmSAMLAttribute) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `SourceType`); value.Exists() && !data.SourceType.IsNull() {
		data.SourceType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SourceType = types.StringNull()
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && !data.Name.IsNull() {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Format`); value.Exists() && !data.Format.IsNull() {
		data.Format = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Format = types.StringNull()
	}
	if value := res.Get(pathRoot + `XPath`); value.Exists() && !data.XPath.IsNull() {
		data.XPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValueData`); value.Exists() && !data.ValueData.IsNull() {
		data.ValueData = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ValueData = types.StringNull()
	}
	if value := res.Get(pathRoot + `SubValueData`); value.Exists() && !data.SubValueData.IsNull() {
		data.SubValueData = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SubValueData = types.StringNull()
	}
	if value := res.Get(pathRoot + `FriendlyName`); value.Exists() && !data.FriendlyName.IsNull() {
		data.FriendlyName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FriendlyName = types.StringNull()
	}
}
