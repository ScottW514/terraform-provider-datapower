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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmMatchRule struct {
	Type            types.String `tfsdk:"type"`
	HttpTag         types.String `tfsdk:"http_tag"`
	HttpValue       types.String `tfsdk:"http_value"`
	Url             types.String `tfsdk:"url"`
	ErrorCode       types.String `tfsdk:"error_code"`
	XpathExpression types.String `tfsdk:"xpath_expression"`
	Method          types.String `tfsdk:"method"`
	CustomMethod    types.String `tfsdk:"custom_method"`
}

var DmMatchRuleObjectType = map[string]attr.Type{
	"type":             types.StringType,
	"http_tag":         types.StringType,
	"http_value":       types.StringType,
	"url":              types.StringType,
	"error_code":       types.StringType,
	"xpath_expression": types.StringType,
	"method":           types.StringType,
	"custom_method":    types.StringType,
}
var DmMatchRuleObjectDefault = map[string]attr.Value{
	"type":             types.StringNull(),
	"http_tag":         types.StringNull(),
	"http_value":       types.StringNull(),
	"url":              types.StringNull(),
	"error_code":       types.StringNull(),
	"xpath_expression": types.StringNull(),
	"method":           types.StringValue("default"),
	"custom_method":    types.StringNull(),
}
var DmMatchRuleDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Matching type", "", "").AddStringEnum("url", "http", "errorcode", "xpath", "fullyqualifiedurl", "host", "method").String,
			Computed:            true,
		},
		"http_tag": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("HTTP header", "", "").String,
			Computed:            true,
		},
		"http_value": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("HTTP value match", "", "").String,
			Computed:            true,
		},
		"url": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL match", "", "").String,
			Computed:            true,
		},
		"error_code": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Error code", "", "").String,
			Computed:            true,
		},
		"xpath_expression": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("XPath expression", "", "").String,
			Computed:            true,
		},
		"method": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("HTTP method", "", "").AddStringEnum("GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS", "TRACE", "custom", "default").AddDefaultValue("default").String,
			Computed:            true,
		},
		"custom_method": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Custom method", "", "").String,
			Computed:            true,
		},
	},
}
var DmMatchRuleResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Matching type", "", "").AddStringEnum("url", "http", "errorcode", "xpath", "fullyqualifiedurl", "host", "method").String,
			Required:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("url", "http", "errorcode", "xpath", "fullyqualifiedurl", "host", "method"),
			},
		},
		"http_tag": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("HTTP header", "", "").String,
			Optional:            true,
		},
		"http_value": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("HTTP value match", "", "").String,
			Optional:            true,
		},
		"url": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL match", "", "").String,
			Optional:            true,
		},
		"error_code": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Error code", "", "").String,
			Optional:            true,
		},
		"xpath_expression": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("XPath expression", "", "").String,
			Optional:            true,
		},
		"method": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("HTTP method", "", "").AddStringEnum("GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS", "TRACE", "custom", "default").AddDefaultValue("default").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS", "TRACE", "custom", "default"),
			},
			Default: stringdefault.StaticString("default"),
		},
		"custom_method": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Custom method", "", "").String,
			Optional:            true,
		},
	},
}

func (data DmMatchRule) IsNull() bool {
	if !data.Type.IsNull() {
		return false
	}
	if !data.HttpTag.IsNull() {
		return false
	}
	if !data.HttpValue.IsNull() {
		return false
	}
	if !data.Url.IsNull() {
		return false
	}
	if !data.ErrorCode.IsNull() {
		return false
	}
	if !data.XpathExpression.IsNull() {
		return false
	}
	if !data.Method.IsNull() {
		return false
	}
	if !data.CustomMethod.IsNull() {
		return false
	}
	return true
}

func (data DmMatchRule) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Type`, data.Type.ValueString())
	}
	if !data.HttpTag.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HttpTag`, data.HttpTag.ValueString())
	}
	if !data.HttpValue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HttpValue`, data.HttpValue.ValueString())
	}
	if !data.Url.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Url`, data.Url.ValueString())
	}
	if !data.ErrorCode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ErrorCode`, data.ErrorCode.ValueString())
	}
	if !data.XpathExpression.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XPATHExpression`, data.XpathExpression.ValueString())
	}
	if !data.Method.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Method`, data.Method.ValueString())
	}
	if !data.CustomMethod.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CustomMethod`, data.CustomMethod.ValueString())
	}
	return body
}

func (data *DmMatchRule) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `HttpTag`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HttpTag = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HttpTag = types.StringNull()
	}
	if value := res.Get(pathRoot + `HttpValue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HttpValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HttpValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `Url`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Url = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Url = types.StringNull()
	}
	if value := res.Get(pathRoot + `ErrorCode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ErrorCode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErrorCode = types.StringNull()
	}
	if value := res.Get(pathRoot + `XPATHExpression`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.XpathExpression = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XpathExpression = types.StringNull()
	}
	if value := res.Get(pathRoot + `Method`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Method = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Method = types.StringValue("default")
	}
	if value := res.Get(pathRoot + `CustomMethod`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CustomMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CustomMethod = types.StringNull()
	}
}

func (data *DmMatchRule) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && !data.Type.IsNull() {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `HttpTag`); value.Exists() && !data.HttpTag.IsNull() {
		data.HttpTag = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HttpTag = types.StringNull()
	}
	if value := res.Get(pathRoot + `HttpValue`); value.Exists() && !data.HttpValue.IsNull() {
		data.HttpValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HttpValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `Url`); value.Exists() && !data.Url.IsNull() {
		data.Url = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Url = types.StringNull()
	}
	if value := res.Get(pathRoot + `ErrorCode`); value.Exists() && !data.ErrorCode.IsNull() {
		data.ErrorCode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErrorCode = types.StringNull()
	}
	if value := res.Get(pathRoot + `XPATHExpression`); value.Exists() && !data.XpathExpression.IsNull() {
		data.XpathExpression = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XpathExpression = types.StringNull()
	}
	if value := res.Get(pathRoot + `Method`); value.Exists() && !data.Method.IsNull() {
		data.Method = tfutils.ParseStringFromGJSON(value)
	} else if data.Method.ValueString() != "default" {
		data.Method = types.StringNull()
	}
	if value := res.Get(pathRoot + `CustomMethod`); value.Exists() && !data.CustomMethod.IsNull() {
		data.CustomMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CustomMethod = types.StringNull()
	}
}
