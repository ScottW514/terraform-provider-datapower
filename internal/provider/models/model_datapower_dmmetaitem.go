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

type DmMetaItem struct {
	MetaCategory types.String `tfsdk:"meta_category"`
	MetaName     types.String `tfsdk:"meta_name"`
	DataSource   types.String `tfsdk:"data_source"`
}

var DmMetaItemObjectType = map[string]attr.Type{
	"meta_category": types.StringType,
	"meta_name":     types.StringType,
	"data_source":   types.StringType,
}
var DmMetaItemObjectDefault = map[string]attr.Value{
	"meta_category": types.StringNull(),
	"meta_name":     types.StringNull(),
	"data_source":   types.StringNull(),
}
var DmMetaItemDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"meta_category": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Metadata Category", "", "").AddStringEnum("all", "mq", "tibco", "wasjms", "http", "CUSTOMIZABLE", "header", "variable").String,
			Computed:            true,
		},
		"meta_name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Metadata Item", "", "").String,
			Computed:            true,
		},
		"data_source": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Custom Data Source", "", "").String,
			Computed:            true,
		},
	},
}
var DmMetaItemResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"meta_category": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Metadata Category", "", "").AddStringEnum("all", "mq", "tibco", "wasjms", "http", "CUSTOMIZABLE", "header", "variable").String,
			Required:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("all", "mq", "tibco", "wasjms", "http", "CUSTOMIZABLE", "header", "variable"),
			},
		},
		"meta_name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Metadata Item", "", "").String,
			Required:            true,
		},
		"data_source": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Custom Data Source", "", "").String,
			Optional:            true,
		},
	},
}

func (data DmMetaItem) IsNull() bool {
	if !data.MetaCategory.IsNull() {
		return false
	}
	if !data.MetaName.IsNull() {
		return false
	}
	if !data.DataSource.IsNull() {
		return false
	}
	return true
}

func (data DmMetaItem) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.MetaCategory.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MetaCategory`, data.MetaCategory.ValueString())
	}
	if !data.MetaName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MetaName`, data.MetaName.ValueString())
	}
	if !data.DataSource.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DataSource`, data.DataSource.ValueString())
	}
	return body
}

func (data *DmMetaItem) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `MetaCategory`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MetaCategory = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MetaCategory = types.StringNull()
	}
	if value := res.Get(pathRoot + `MetaName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MetaName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MetaName = types.StringNull()
	}
	if value := res.Get(pathRoot + `DataSource`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DataSource = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DataSource = types.StringNull()
	}
}

func (data *DmMetaItem) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `MetaCategory`); value.Exists() && !data.MetaCategory.IsNull() {
		data.MetaCategory = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MetaCategory = types.StringNull()
	}
	if value := res.Get(pathRoot + `MetaName`); value.Exists() && !data.MetaName.IsNull() {
		data.MetaName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MetaName = types.StringNull()
	}
	if value := res.Get(pathRoot + `DataSource`); value.Exists() && !data.DataSource.IsNull() {
		data.DataSource = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DataSource = types.StringNull()
	}
}
