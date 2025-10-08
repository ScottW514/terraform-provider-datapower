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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmMCFilter struct {
	FilterName      types.String `tfsdk:"filter_name"`
	Type            types.String `tfsdk:"type"`
	HttpName        types.String `tfsdk:"http_name"`
	HttpValue       types.String `tfsdk:"http_value"`
	XpathExpression types.String `tfsdk:"xpath_expression"`
	XpathValue      types.String `tfsdk:"xpath_value"`
}

var DmMCFilterHttpNameCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"http"},
}

var DmMCFilterHttpNameIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmMCFilterHttpValueCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"http"},
}

var DmMCFilterHttpValueIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmMCFilterXPathExpressionCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"xpath"},
}

var DmMCFilterXPathExpressionIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmMCFilterXPathValueCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"xpath"},
}

var DmMCFilterXPathValueIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmMCFilterObjectType = map[string]attr.Type{
	"filter_name":      types.StringType,
	"type":             types.StringType,
	"http_name":        types.StringType,
	"http_value":       types.StringType,
	"xpath_expression": types.StringType,
	"xpath_value":      types.StringType,
}
var DmMCFilterObjectDefault = map[string]attr.Value{
	"filter_name":      types.StringNull(),
	"type":             types.StringNull(),
	"http_name":        types.StringNull(),
	"http_value":       types.StringNull(),
	"xpath_expression": types.StringNull(),
	"xpath_value":      types.StringNull(),
}

func GetDmMCFilterDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmMCFilterDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"filter_name": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the name for the message content filter.",
				Computed:            true,
			},
			"type": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Select the type of the message content filter.",
				Computed:            true,
			},
			"http_name": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Enter the name of the HTTP header field. Available for HTTP header-based filters.",
				Computed:            true,
			},
			"http_value": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Enter the explicit, literal string value for the HTTP header field. Wildcards are not supported. Available for HTTP header-based filters.",
				Computed:            true,
			},
			"xpath_expression": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Enter the XPath expression or use the builder to define the XPath expression that is used to evaluate the messages to obtain the XPath value. Available for XPath-based filters.",
				Computed:            true,
			},
			"xpath_value": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Enter the contents of the element for the XPath expression. Available for XPath-based filters.",
				Computed:            true,
			},
		},
	}
	return DmMCFilterDataSourceSchema
}
func GetDmMCFilterResourceSchema() ResourceSchema.NestedAttributeObject {
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
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of the HTTP header field. Available for HTTP header-based filters.", "", "").AddRequiredWhen(DmMCFilterHttpNameCondVal.String()).AddNotValidWhen(DmMCFilterHttpNameIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmMCFilterHttpNameCondVal, DmMCFilterHttpNameIgnoreVal, false),
				},
			},
			"http_value": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the explicit, literal string value for the HTTP header field. Wildcards are not supported. Available for HTTP header-based filters.", "", "").AddRequiredWhen(DmMCFilterHttpValueCondVal.String()).AddNotValidWhen(DmMCFilterHttpValueIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmMCFilterHttpValueCondVal, DmMCFilterHttpValueIgnoreVal, false),
				},
			},
			"xpath_expression": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the XPath expression or use the builder to define the XPath expression that is used to evaluate the messages to obtain the XPath value. Available for XPath-based filters.", "", "").AddRequiredWhen(DmMCFilterXPathExpressionCondVal.String()).AddNotValidWhen(DmMCFilterXPathExpressionIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmMCFilterXPathExpressionCondVal, DmMCFilterXPathExpressionIgnoreVal, false),
				},
			},
			"xpath_value": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the contents of the element for the XPath expression. Available for XPath-based filters.", "", "").AddRequiredWhen(DmMCFilterXPathValueCondVal.String()).AddNotValidWhen(DmMCFilterXPathValueIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmMCFilterXPathValueCondVal, DmMCFilterXPathValueIgnoreVal, false),
				},
			},
		},
	}
	return DmMCFilterResourceSchema
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
	if !data.XpathExpression.IsNull() {
		return false
	}
	if !data.XpathValue.IsNull() {
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
	if !data.XpathExpression.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XPathExpression`, data.XpathExpression.ValueString())
	}
	if !data.XpathValue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XPathValue`, data.XpathValue.ValueString())
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
		data.XpathExpression = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XpathExpression = types.StringNull()
	}
	if value := res.Get(pathRoot + `XPathValue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.XpathValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XpathValue = types.StringNull()
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
	if value := res.Get(pathRoot + `XPathExpression`); value.Exists() && !data.XpathExpression.IsNull() {
		data.XpathExpression = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XpathExpression = types.StringNull()
	}
	if value := res.Get(pathRoot + `XPathValue`); value.Exists() && !data.XpathValue.IsNull() {
		data.XpathValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XpathValue = types.StringNull()
	}
}
