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

type DmHeaderSuppression struct {
	Direction types.String `tfsdk:"direction"`
	HeaderTag types.String `tfsdk:"header_tag"`
}

var DmHeaderSuppressionObjectType = map[string]attr.Type{
	"direction":  types.StringType,
	"header_tag": types.StringType,
}
var DmHeaderSuppressionObjectDefault = map[string]attr.Value{
	"direction":  types.StringNull(),
	"header_tag": types.StringNull(),
}

func GetDmHeaderSuppressionDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmHeaderSuppressionDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"direction": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the direction of the message.", "", "").AddStringEnum("front", "back").String,
				Computed:            true,
			},
			"header_tag": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of the header to suppress. When these headers are defined in the original request, the device removes the specified headers before forwarding the request to the backend server.", "", "").String,
				Computed:            true,
			},
		},
	}
	return DmHeaderSuppressionDataSourceSchema
}
func GetDmHeaderSuppressionResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmHeaderSuppressionResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"direction": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the direction of the message.", "", "").AddStringEnum("front", "back").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("front", "back"),
				},
			},
			"header_tag": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of the header to suppress. When these headers are defined in the original request, the device removes the specified headers before forwarding the request to the backend server.", "", "").String,
				Optional:            true,
			},
		},
	}
	return DmHeaderSuppressionResourceSchema
}

func (data DmHeaderSuppression) IsNull() bool {
	if !data.Direction.IsNull() {
		return false
	}
	if !data.HeaderTag.IsNull() {
		return false
	}
	return true
}

func (data DmHeaderSuppression) ToBody(ctx context.Context, pathRoot string) string {
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
	return body
}

func (data *DmHeaderSuppression) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Direction`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Direction = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Direction = types.StringNull()
	}
	if value := res.Get(pathRoot + `HeaderTag`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HeaderTag = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HeaderTag = types.StringNull()
	}
}

func (data *DmHeaderSuppression) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Direction`); value.Exists() && !data.Direction.IsNull() {
		data.Direction = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Direction = types.StringNull()
	}
	if value := res.Get(pathRoot + `HeaderTag`); value.Exists() && !data.HeaderTag.IsNull() {
		data.HeaderTag = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HeaderTag = types.StringNull()
	}
}
