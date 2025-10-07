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

type DmHeaderInjection struct {
	Direction      types.String `tfsdk:"direction"`
	HeaderTag      types.String `tfsdk:"header_tag"`
	HeaderTagValue types.String `tfsdk:"header_tag_value"`
}

var DmHeaderInjectionObjectType = map[string]attr.Type{
	"direction":        types.StringType,
	"header_tag":       types.StringType,
	"header_tag_value": types.StringType,
}
var DmHeaderInjectionObjectDefault = map[string]attr.Value{
	"direction":        types.StringValue("front"),
	"header_tag":       types.StringNull(),
	"header_tag_value": types.StringNull(),
}

func GetDmHeaderInjectionDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmHeaderInjectionDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"direction": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the direction of the message.", "", "").AddStringEnum("front", "back").AddDefaultValue("front").String,
				Computed:            true,
			},
			"header_tag": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of the header to inject. Even though the headers are not defined in the original request, the device provides the specified headers to the backend server.", "", "").String,
				Computed:            true,
			},
			"header_tag_value": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the value for the header tag.", "", "").String,
				Computed:            true,
			},
		},
	}
	return DmHeaderInjectionDataSourceSchema
}
func GetDmHeaderInjectionResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmHeaderInjectionResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"direction": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the direction of the message.", "", "").AddStringEnum("front", "back").AddDefaultValue("front").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("front", "back"),
				},
				Default: stringdefault.StaticString("front"),
			},
			"header_tag": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of the header to inject. Even though the headers are not defined in the original request, the device provides the specified headers to the backend server.", "", "").String,
				Optional:            true,
			},
			"header_tag_value": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the value for the header tag.", "", "").String,
				Required:            true,
			},
		},
	}
	return DmHeaderInjectionResourceSchema
}

func (data DmHeaderInjection) IsNull() bool {
	if !data.Direction.IsNull() {
		return false
	}
	if !data.HeaderTag.IsNull() {
		return false
	}
	if !data.HeaderTagValue.IsNull() {
		return false
	}
	return true
}

func (data DmHeaderInjection) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Direction.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Direction`, data.Direction.ValueString())
	}
	if !data.HeaderTag.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HeaderTag`, data.HeaderTag.ValueString())
	}
	if !data.HeaderTagValue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HeaderTagValue`, data.HeaderTagValue.ValueString())
	}
	return body
}

func (data *DmHeaderInjection) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Direction`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Direction = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Direction = types.StringValue("front")
	}
	if value := res.Get(pathRoot + `HeaderTag`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HeaderTag = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HeaderTag = types.StringNull()
	}
	if value := res.Get(pathRoot + `HeaderTagValue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HeaderTagValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HeaderTagValue = types.StringNull()
	}
}

func (data *DmHeaderInjection) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Direction`); value.Exists() && !data.Direction.IsNull() {
		data.Direction = tfutils.ParseStringFromGJSON(value)
	} else if data.Direction.ValueString() != "front" {
		data.Direction = types.StringNull()
	}
	if value := res.Get(pathRoot + `HeaderTag`); value.Exists() && !data.HeaderTag.IsNull() {
		data.HeaderTag = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HeaderTag = types.StringNull()
	}
	if value := res.Get(pathRoot + `HeaderTagValue`); value.Exists() && !data.HeaderTagValue.IsNull() {
		data.HeaderTagValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HeaderTagValue = types.StringNull()
	}
}
