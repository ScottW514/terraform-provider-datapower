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

type DmInputEncoding struct {
	InputMatch types.String `tfsdk:"input_match"`
	Encoding   types.String `tfsdk:"encoding"`
}

var DmInputEncodingObjectType = map[string]attr.Type{
	"input_match": types.StringType,
	"encoding":    types.StringType,
}
var DmInputEncodingObjectDefault = map[string]attr.Value{
	"input_match": types.StringNull(),
	"encoding":    types.StringNull(),
}

func GetDmInputEncodingDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmInputEncodingDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"input_match": DataSourceSchema.StringAttribute{
				MarkdownDescription: "The PCRE pattern that will be matched against the name of an HTTP form field. If the name matches this PCRE, the associated value will be processed by the rules of the Encoding. If a form field has no name, an empty string is used for the pattern matching.",
				Computed:            true,
			},
			"encoding": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Select how to translate the value of a form field to the contents of the &lt;arg> element in the generated XML.",
				Computed:            true,
			},
		},
	}
	return DmInputEncodingDataSourceSchema
}
func GetDmInputEncodingResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmInputEncodingResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"input_match": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The PCRE pattern that will be matched against the name of an HTTP form field. If the name matches this PCRE, the associated value will be processed by the rules of the Encoding. If a form field has no name, an empty string is used for the pattern matching.", "input-name", "").String,
				Optional:            true,
			},
			"encoding": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select how to translate the value of a form field to the contents of the &lt;arg> element in the generated XML.", "encoding", "").AddStringEnum("plain", "urlencoded", "xml", "urlencoded-xml", "base64", "base64-text", "base64-xml", "json").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("plain", "urlencoded", "xml", "urlencoded-xml", "base64", "base64-text", "base64-xml", "json"),
				},
			},
		},
	}
	return DmInputEncodingResourceSchema
}

func (data DmInputEncoding) IsNull() bool {
	if !data.InputMatch.IsNull() {
		return false
	}
	if !data.Encoding.IsNull() {
		return false
	}
	return true
}

func (data DmInputEncoding) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.InputMatch.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`InputMatch`, data.InputMatch.ValueString())
	}
	if !data.Encoding.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Encoding`, data.Encoding.ValueString())
	}
	return body
}

func (data *DmInputEncoding) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `InputMatch`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.InputMatch = tfutils.ParseStringFromGJSON(value)
	} else {
		data.InputMatch = types.StringNull()
	}
	if value := res.Get(pathRoot + `Encoding`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Encoding = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Encoding = types.StringNull()
	}
}

func (data *DmInputEncoding) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `InputMatch`); value.Exists() && !data.InputMatch.IsNull() {
		data.InputMatch = tfutils.ParseStringFromGJSON(value)
	} else {
		data.InputMatch = types.StringNull()
	}
	if value := res.Get(pathRoot + `Encoding`); value.Exists() && !data.Encoding.IsNull() {
		data.Encoding = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Encoding = types.StringNull()
	}
}
