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

type DmLogIPFilter struct {
	IpAddress types.String `tfsdk:"ip_address"`
}

var DmLogIPFilterObjectType = map[string]attr.Type{
	"ip_address": types.StringType,
}
var DmLogIPFilterObjectDefault = map[string]attr.Value{
	"ip_address": types.StringNull(),
}

func GetDmLogIPFilterDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmLogIPFilterDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"ip_address": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the explicit IP address.",
				Computed:            true,
			},
		},
	}
	return DmLogIPFilterDataSourceSchema
}
func GetDmLogIPFilterResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmLogIPFilterResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"ip_address": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the explicit IP address.", "", "").String,
				Optional:            true,
			},
		},
	}
	return DmLogIPFilterResourceSchema
}

func (data DmLogIPFilter) IsNull() bool {
	if !data.IpAddress.IsNull() {
		return false
	}
	return true
}

func (data DmLogIPFilter) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.IpAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IPAddress`, data.IpAddress.ValueString())
	}
	return body
}

func (data *DmLogIPFilter) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `IPAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.IpAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.IpAddress = types.StringNull()
	}
}

func (data *DmLogIPFilter) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `IPAddress`); value.Exists() && !data.IpAddress.IsNull() {
		data.IpAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.IpAddress = types.StringNull()
	}
}
