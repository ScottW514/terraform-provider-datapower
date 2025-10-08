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

type DmHTTPHeader struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

var DmHTTPHeaderObjectType = map[string]attr.Type{
	"name":  types.StringType,
	"value": types.StringType,
}
var DmHTTPHeaderObjectDefault = map[string]attr.Value{
	"name":  types.StringNull(),
	"value": types.StringNull(),
}

func GetDmHTTPHeaderDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmHTTPHeaderDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"name": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Enter the name of the HTTP header field to be examined.",
				Computed:            true,
			},
			"value": DataSourceSchema.StringAttribute{
				MarkdownDescription: "<p>Provide a literal or wildcard expression to define a set of values to include in or exclude from the traffic definition. If the contents of the specified header field fulfill the defined match criterion, the match succeeds.</p><p>The following wildcard characters are available:</p><table><tr><td valign=\"top\">asterisk (*)</td><td>Matches 0 or more occurrences of any character</td></tr><tr><td valign=\"top\">question mark (?)</td><td>Matches one occurrence of any single character</td></tr><tr><td valign=\"top\">square brackets ([ ])</td><td>Defines a character or numeric range. <p>For example, [1-5] matches 1, 2, 3, 4, or 5, while xs[dl] matches xsd or xsl.</p></td></tr></table>",
				Computed:            true,
			},
		},
	}
	return DmHTTPHeaderDataSourceSchema
}
func GetDmHTTPHeaderResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmHTTPHeaderResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"name": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of the HTTP header field to be examined.", "", "").String,
				Required:            true,
			},
			"value": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Provide a literal or wildcard expression to define a set of values to include in or exclude from the traffic definition. If the contents of the specified header field fulfill the defined match criterion, the match succeeds.</p><p>The following wildcard characters are available:</p><table><tr><td valign=\"top\">asterisk (*)</td><td>Matches 0 or more occurrences of any character</td></tr><tr><td valign=\"top\">question mark (?)</td><td>Matches one occurrence of any single character</td></tr><tr><td valign=\"top\">square brackets ([ ])</td><td>Defines a character or numeric range. <p>For example, [1-5] matches 1, 2, 3, 4, or 5, while xs[dl] matches xsd or xsl.</p></td></tr></table>", "", "").String,
				Required:            true,
			},
		},
	}
	return DmHTTPHeaderResourceSchema
}

func (data DmHTTPHeader) IsNull() bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.Value.IsNull() {
		return false
	}
	return true
}

func (data DmHTTPHeader) ToBody(ctx context.Context, pathRoot string) string {
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
	return body
}

func (data *DmHTTPHeader) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
}

func (data *DmHTTPHeader) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
}
