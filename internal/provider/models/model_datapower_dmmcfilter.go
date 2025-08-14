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

type DmMCFilter struct {
	FilterName      types.String `tfsdk:"filter_name"`
	Type            types.String `tfsdk:"type"`
	HttpName        types.String `tfsdk:"http_name"`
	HttpValue       types.String `tfsdk:"http_value"`
	XPathExpression types.String `tfsdk:"x_path_expression"`
	XPathValue      types.String `tfsdk:"x_path_value"`
}

var DmMCFilterObjectType = map[string]attr.Type{
	"filter_name":       types.StringType,
	"type":              types.StringType,
	"http_name":         types.StringType,
	"http_value":        types.StringType,
	"x_path_expression": types.StringType,
	"x_path_value":      types.StringType,
}
var DmMCFilterObjectDefault = map[string]attr.Value{
	"filter_name":       types.StringNull(),
	"type":              types.StringNull(),
	"http_name":         types.StringNull(),
	"http_value":        types.StringNull(),
	"x_path_expression": types.StringNull(),
	"x_path_value":      types.StringNull(),
}
var DmMCFilterDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"filter_name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the name for the message content filter.", "", "").String,
			Computed:            true,
		},
		"type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Select the type of the message content filter.", "", "").AddStringEnum("http", "xpath").String,
			Computed:            true,
		},
		"http_name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of the HTTP header field. Available for HTTP header-based filters.", "", "").String,
			Computed:            true,
		},
		"http_value": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enter the explicit, literal string value for the HTTP header field. Wildcards are not supported. Available for HTTP header-based filters.", "", "").String,
			Computed:            true,
		},
		"x_path_expression": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enter the XPath expression or use the builder to define the XPath expression that is used to evaluate the messages to obtain the XPath value. Available for XPath-based filters.", "", "").String,
			Computed:            true,
		},
		"x_path_value": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enter the contents of the element for the XPath expression. Available for XPath-based filters.", "", "").String,
			Computed:            true,
		},
	},
}
var DmMCFilterResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"filter_name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the name for the message content filter.", "", "").String,
			Required:            true,
		},
		"type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Select the type of the message content filter.", "", "").AddStringEnum("http", "xpath").String,
			Required:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("http", "xpath"),
			},
		},
		"http_name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of the HTTP header field. Available for HTTP header-based filters.", "", "").String,
			Optional:            true,
		},
		"http_value": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enter the explicit, literal string value for the HTTP header field. Wildcards are not supported. Available for HTTP header-based filters.", "", "").String,
			Optional:            true,
		},
		"x_path_expression": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enter the XPath expression or use the builder to define the XPath expression that is used to evaluate the messages to obtain the XPath value. Available for XPath-based filters.", "", "").String,
			Optional:            true,
		},
		"x_path_value": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enter the contents of the element for the XPath expression. Available for XPath-based filters.", "", "").String,
			Optional:            true,
		},
	},
}

func (data DmMCFilter) IsNull() bool {
	if !data.FilterName.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
		return false
	}
	if !data.HttpName.IsNull() {
		return false
	}
	if !data.HttpValue.IsNull() {
		return false
	}
	if !data.XPathExpression.IsNull() {
		return false
	}
	if !data.XPathValue.IsNull() {
		return false
	}
	return true
}

func (data DmMCFilter) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.FilterName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FilterName`, data.FilterName.ValueString())
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Type`, data.Type.ValueString())
	}
	if !data.HttpName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HttpName`, data.HttpName.ValueString())
	}
	if !data.HttpValue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HttpValue`, data.HttpValue.ValueString())
	}
	if !data.XPathExpression.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XPathExpression`, data.XPathExpression.ValueString())
	}
	if !data.XPathValue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XPathValue`, data.XPathValue.ValueString())
	}
	return body
}

func (data *DmMCFilter) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `FilterName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.FilterName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FilterName = types.StringNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `HttpName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HttpName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HttpName = types.StringNull()
	}
	if value := res.Get(pathRoot + `HttpValue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HttpValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HttpValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `XPathExpression`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.XPathExpression = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XPathExpression = types.StringNull()
	}
	if value := res.Get(pathRoot + `XPathValue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.XPathValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XPathValue = types.StringNull()
	}
}

func (data *DmMCFilter) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `FilterName`); value.Exists() && !data.FilterName.IsNull() {
		data.FilterName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FilterName = types.StringNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && !data.Type.IsNull() {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `HttpName`); value.Exists() && !data.HttpName.IsNull() {
		data.HttpName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HttpName = types.StringNull()
	}
	if value := res.Get(pathRoot + `HttpValue`); value.Exists() && !data.HttpValue.IsNull() {
		data.HttpValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HttpValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `XPathExpression`); value.Exists() && !data.XPathExpression.IsNull() {
		data.XPathExpression = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XPathExpression = types.StringNull()
	}
	if value := res.Get(pathRoot + `XPathValue`); value.Exists() && !data.XPathValue.IsNull() {
		data.XPathValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XPathValue = types.StringNull()
	}
}
