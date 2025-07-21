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

type DmLTPAUserAttributeNameAndValue struct {
	LtpaUserAttributeName        types.String `tfsdk:"ltpa_user_attribute_name"`
	LtpaUserAttributeType        types.String `tfsdk:"ltpa_user_attribute_type"`
	LtpaUserAttributeStaticValue types.String `tfsdk:"ltpa_user_attribute_static_value"`
	LtpaUserAttributeXPathValue  types.String `tfsdk:"ltpa_user_attribute_x_path_value"`
}

var DmLTPAUserAttributeNameAndValueObjectType = map[string]attr.Type{
	"ltpa_user_attribute_name":         types.StringType,
	"ltpa_user_attribute_type":         types.StringType,
	"ltpa_user_attribute_static_value": types.StringType,
	"ltpa_user_attribute_x_path_value": types.StringType,
}
var DmLTPAUserAttributeNameAndValueObjectDefault = map[string]attr.Value{
	"ltpa_user_attribute_name":         types.StringNull(),
	"ltpa_user_attribute_type":         types.StringNull(),
	"ltpa_user_attribute_static_value": types.StringNull(),
	"ltpa_user_attribute_x_path_value": types.StringNull(),
}
var DmLTPAUserAttributeNameAndValueDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"ltpa_user_attribute_name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Attribute name", "", "").String,
			Computed:            true,
		},
		"ltpa_user_attribute_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Attribute type", "", "").AddStringEnum("static", "xpath").String,
			Computed:            true,
		},
		"ltpa_user_attribute_static_value": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Static attribute value", "", "").String,
			Computed:            true,
		},
		"ltpa_user_attribute_x_path_value": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("XPath attribute value", "", "").String,
			Computed:            true,
		},
	},
}
var DmLTPAUserAttributeNameAndValueResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"ltpa_user_attribute_name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Attribute name", "", "").String,
			Required:            true,
		},
		"ltpa_user_attribute_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Attribute type", "", "").AddStringEnum("static", "xpath").String,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("static", "xpath"),
			},
		},
		"ltpa_user_attribute_static_value": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Static attribute value", "", "").String,
			Optional:            true,
		},
		"ltpa_user_attribute_x_path_value": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("XPath attribute value", "", "").String,
			Optional:            true,
		},
	},
}

func (data DmLTPAUserAttributeNameAndValue) IsNull() bool {
	if !data.LtpaUserAttributeName.IsNull() {
		return false
	}
	if !data.LtpaUserAttributeType.IsNull() {
		return false
	}
	if !data.LtpaUserAttributeStaticValue.IsNull() {
		return false
	}
	if !data.LtpaUserAttributeXPathValue.IsNull() {
		return false
	}
	return true
}

func (data DmLTPAUserAttributeNameAndValue) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.LtpaUserAttributeName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LTPAUserAttributeName`, data.LtpaUserAttributeName.ValueString())
	}
	if !data.LtpaUserAttributeType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LTPAUserAttributeType`, data.LtpaUserAttributeType.ValueString())
	}
	if !data.LtpaUserAttributeStaticValue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LTPAUserAttributeStaticValue`, data.LtpaUserAttributeStaticValue.ValueString())
	}
	if !data.LtpaUserAttributeXPathValue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LTPAUserAttributeXPathValue`, data.LtpaUserAttributeXPathValue.ValueString())
	}
	return body
}

func (data *DmLTPAUserAttributeNameAndValue) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `LTPAUserAttributeName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LtpaUserAttributeName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LtpaUserAttributeName = types.StringNull()
	}
	if value := res.Get(pathRoot + `LTPAUserAttributeType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LtpaUserAttributeType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LtpaUserAttributeType = types.StringNull()
	}
	if value := res.Get(pathRoot + `LTPAUserAttributeStaticValue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LtpaUserAttributeStaticValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LtpaUserAttributeStaticValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `LTPAUserAttributeXPathValue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LtpaUserAttributeXPathValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LtpaUserAttributeXPathValue = types.StringNull()
	}
}

func (data *DmLTPAUserAttributeNameAndValue) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `LTPAUserAttributeName`); value.Exists() && !data.LtpaUserAttributeName.IsNull() {
		data.LtpaUserAttributeName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LtpaUserAttributeName = types.StringNull()
	}
	if value := res.Get(pathRoot + `LTPAUserAttributeType`); value.Exists() && !data.LtpaUserAttributeType.IsNull() {
		data.LtpaUserAttributeType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LtpaUserAttributeType = types.StringNull()
	}
	if value := res.Get(pathRoot + `LTPAUserAttributeStaticValue`); value.Exists() && !data.LtpaUserAttributeStaticValue.IsNull() {
		data.LtpaUserAttributeStaticValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LtpaUserAttributeStaticValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `LTPAUserAttributeXPathValue`); value.Exists() && !data.LtpaUserAttributeXPathValue.IsNull() {
		data.LtpaUserAttributeXPathValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LtpaUserAttributeXPathValue = types.StringNull()
	}
}
