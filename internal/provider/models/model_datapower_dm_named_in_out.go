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

type DmNamedInOut struct {
	Name    types.String `tfsdk:"name"`
	Context types.String `tfsdk:"context"`
}

var DmNamedInOutObjectType = map[string]attr.Type{
	"name":    types.StringType,
	"context": types.StringType,
}
var DmNamedInOutObjectDefault = map[string]attr.Value{
	"name":    types.StringNull(),
	"context": types.StringNull(),
}

func GetDmNamedInOutDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmNamedInOutDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"name": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the input or output that the map expects. The name must be the same as a cardname that is identified in the map file.", "", "").String,
				Computed:            true,
			},
			"context": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the DataPower context. This context contains the input data or will contain the output data that corresponds to the input or output that the maps expects. Use <tt>INPUT</tt> to designate the context that contains the original request. Use <tt>OUTPUT</tt> to designate the output context.", "", "").String,
				Computed:            true,
			},
		},
	}
	return DmNamedInOutDataSourceSchema
}
func GetDmNamedInOutResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmNamedInOutResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"name": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the input or output that the map expects. The name must be the same as a cardname that is identified in the map file.", "", "").String,
				Optional:            true,
			},
			"context": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the DataPower context. This context contains the input data or will contain the output data that corresponds to the input or output that the maps expects. Use <tt>INPUT</tt> to designate the context that contains the original request. Use <tt>OUTPUT</tt> to designate the output context.", "", "").String,
				Required:            true,
			},
		},
	}
	return DmNamedInOutResourceSchema
}

func (data DmNamedInOut) IsNull() bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.Context.IsNull() {
		return false
	}
	return true
}

func (data DmNamedInOut) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Name`, data.Name.ValueString())
	}
	if !data.Context.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Context`, data.Context.ValueString())
	}
	return body
}

func (data *DmNamedInOut) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Context`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Context = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Context = types.StringNull()
	}
}

func (data *DmNamedInOut) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && !data.Name.IsNull() {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Context`); value.Exists() && !data.Context.IsNull() {
		data.Context = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Context = types.StringNull()
	}
}
