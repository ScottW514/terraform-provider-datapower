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

type DmTableEntry struct {
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
}

var DmTableEntryObjectType = map[string]attr.Type{
	"name":        types.StringType,
	"description": types.StringType,
}
var DmTableEntryObjectDefault = map[string]attr.Value{
	"name":        types.StringNull(),
	"description": types.StringNull(),
}
var DmTableEntryDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Name", "", "").String,
			Computed:            true,
		},
		"description": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Description", "", "").String,
			Computed:            true,
		},
	},
}
var DmTableEntryResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Name", "", "").String,
			Required:            true,
		},
		"description": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Description", "", "").String,
			Optional:            true,
		},
	},
}

func (data DmTableEntry) IsNull() bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.Description.IsNull() {
		return false
	}
	return true
}

func (data DmTableEntry) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Name`, data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Description`, data.Description.ValueString())
	}
	return body
}

func (data *DmTableEntry) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Description`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Description = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Description = types.StringNull()
	}
}

func (data *DmTableEntry) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && !data.Name.IsNull() {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Description`); value.Exists() && !data.Description.IsNull() {
		data.Description = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Description = types.StringNull()
	}
}
