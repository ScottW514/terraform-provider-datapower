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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmMetaItem struct {
	MetaCategory types.String `tfsdk:"meta_category"`
	MetaName     types.String `tfsdk:"meta_name"`
	DataSource   types.String `tfsdk:"data_source"`
}

var DmMetaItemDataSourceCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "meta_category",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"header", "variable"},
}

var DmMetaItemDataSourceIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
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

func GetDmMetaItemDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmMetaItemDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"meta_category": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Select the category for the Metadata Item. The Metadata Item selections change according to the selected category. To create a custom Metadata Item, select either Custom Header or Custom Variable. For a custom Metadata Item, specify the name of the metadata item and its data source.",
				Computed:            true,
			},
			"meta_name": DataSourceSchema.StringAttribute{
				MarkdownDescription: "<p>For all except custom items, select a Metadata Item. The list provides an alias for the actual name of a protocol header or system variable. The elements contained in the XML nodeset that is returned by the Processing Metadata object have names that correspond to the actual data source</p><p>For custom items, enter an alphanumeric string for this custom alias. The string cannot contain white space.</p>",
				Computed:            true,
			},
			"data_source": DataSourceSchema.StringAttribute{
				MarkdownDescription: "For custom items only, enter the name of the protocol header or the name of the variable (service, context, or system) that contains the data to be returned in the metadata XML nodeset. The provided value is the name of the element in the returned nodeset that contains the data.",
				Computed:            true,
			},
		},
	}
	return DmMetaItemDataSourceSchema
}
func GetDmMetaItemResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmMetaItemResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"meta_category": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the category for the Metadata Item. The Metadata Item selections change according to the selected category. To create a custom Metadata Item, select either Custom Header or Custom Variable. For a custom Metadata Item, specify the name of the metadata item and its data source.", "", "").AddStringEnum("all", "mq", "tibco", "wasjms", "http", "CUSTOMIZABLE", "header", "variable").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("all", "mq", "tibco", "wasjms", "http", "CUSTOMIZABLE", "header", "variable"),
				},
			},
			"meta_name": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>For all except custom items, select a Metadata Item. The list provides an alias for the actual name of a protocol header or system variable. The elements contained in the XML nodeset that is returned by the Processing Metadata object have names that correspond to the actual data source</p><p>For custom items, enter an alphanumeric string for this custom alias. The string cannot contain white space.</p>", "", "").String,
				Required:            true,
			},
			"data_source": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("For custom items only, enter the name of the protocol header or the name of the variable (service, context, or system) that contains the data to be returned in the metadata XML nodeset. The provided value is the name of the element in the returned nodeset that contains the data.", "", "").AddRequiredWhen(DmMetaItemDataSourceCondVal.String()).AddNotValidWhen(DmMetaItemDataSourceIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmMetaItemDataSourceCondVal, DmMetaItemDataSourceIgnoreVal, false),
				},
			},
		},
	}
	return DmMetaItemResourceSchema
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
