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

type DmMtomRule struct {
	XPath       types.String `tfsdk:"x_path"`
	ContentType types.String `tfsdk:"content_type"`
	ContentId   types.String `tfsdk:"content_id"`
}

var DmMtomRuleObjectType = map[string]attr.Type{
	"x_path":       types.StringType,
	"content_type": types.StringType,
	"content_id":   types.StringType,
}
var DmMtomRuleObjectDefault = map[string]attr.Value{
	"x_path":       types.StringNull(),
	"content_type": types.StringNull(),
	"content_id":   types.StringNull(),
}

func GetDmMtomRuleDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmMtomRuleDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"x_path": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("An XPath expression that defines which schema element or schema elements are subject to this rule. Click XPath Tool to launch the XPath expression builder. To use this tool, upload an example document. After the example document is loaded, the tool allows you to select the elements to construct the corresponding XPath expression. After selecting the elements, click Done to return to the MTOM Rules Properties window.", "select", "").String,
				Computed:            true,
			},
			"content_type": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The Content Type for the extracted element data. This option overrides the value that is attained from the xmlmime:contentType attribute. If the provided XPath matches more than one element, each corresponding MIME attachment part will contain a content-type header with this value. If different content-type values are required, selective XPath expressions are required.", "content-type", "").String,
				Computed:            true,
			},
			"content_id": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("If not explicitly configured, content IDs are automatically generated. Using this property allows for the explicit configuration of content-id headers and their associated href values. Rules that match multiple data elements result in one attachment part for all matched elements. The resulting attachment part contains data from the last match only.", "content-id", "").String,
				Computed:            true,
			},
		},
	}
	return DmMtomRuleDataSourceSchema
}
func GetDmMtomRuleResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmMtomRuleResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"x_path": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("An XPath expression that defines which schema element or schema elements are subject to this rule. Click XPath Tool to launch the XPath expression builder. To use this tool, upload an example document. After the example document is loaded, the tool allows you to select the elements to construct the corresponding XPath expression. After selecting the elements, click Done to return to the MTOM Rules Properties window.", "select", "").String,
				Required:            true,
			},
			"content_type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The Content Type for the extracted element data. This option overrides the value that is attained from the xmlmime:contentType attribute. If the provided XPath matches more than one element, each corresponding MIME attachment part will contain a content-type header with this value. If different content-type values are required, selective XPath expressions are required.", "content-type", "").String,
				Optional:            true,
			},
			"content_id": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("If not explicitly configured, content IDs are automatically generated. Using this property allows for the explicit configuration of content-id headers and their associated href values. Rules that match multiple data elements result in one attachment part for all matched elements. The resulting attachment part contains data from the last match only.", "content-id", "").String,
				Optional:            true,
			},
		},
	}
	return DmMtomRuleResourceSchema
}

func (data DmMtomRule) IsNull() bool {
	if !data.XPath.IsNull() {
		return false
	}
	if !data.ContentType.IsNull() {
		return false
	}
	if !data.ContentId.IsNull() {
		return false
	}
	return true
}

func (data DmMtomRule) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.XPath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XPath`, data.XPath.ValueString())
	}
	if !data.ContentType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ContentType`, data.ContentType.ValueString())
	}
	if !data.ContentId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ContentID`, data.ContentId.ValueString())
	}
	return body
}

func (data *DmMtomRule) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `XPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.XPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `ContentType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ContentType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ContentType = types.StringNull()
	}
	if value := res.Get(pathRoot + `ContentID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ContentId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ContentId = types.StringNull()
	}
}

func (data *DmMtomRule) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `XPath`); value.Exists() && !data.XPath.IsNull() {
		data.XPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `ContentType`); value.Exists() && !data.ContentType.IsNull() {
		data.ContentType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ContentType = types.StringNull()
	}
	if value := res.Get(pathRoot + `ContentID`); value.Exists() && !data.ContentId.IsNull() {
		data.ContentId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ContentId = types.StringNull()
	}
}
