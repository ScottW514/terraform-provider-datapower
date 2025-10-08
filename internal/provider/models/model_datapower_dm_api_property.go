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
	"regexp"

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

type DmAPIProperty struct {
	PropertyName types.String `tfsdk:"property_name"`
	Catalog      types.String `tfsdk:"catalog"`
	Value        types.String `tfsdk:"value"`
}

var DmAPIPropertyObjectType = map[string]attr.Type{
	"property_name": types.StringType,
	"catalog":       types.StringType,
	"value":         types.StringType,
}
var DmAPIPropertyObjectDefault = map[string]attr.Value{
	"property_name": types.StringNull(),
	"catalog":       types.StringValue("*"),
	"value":         types.StringNull(),
}

func GetDmAPIPropertyDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmAPIPropertyDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"property_name": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the property name.",
				Computed:            true,
			},
			"catalog": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the catalog name. The name must match the name of an API catalog in the API collection. The default value is <tt>*</tt> , which indicates that the value applies to all catalogs.",
				Computed:            true,
			},
			"value": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the property value.",
				Computed:            true,
			},
		},
	}
	return DmAPIPropertyDataSourceSchema
}
func GetDmAPIPropertyResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmAPIPropertyResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"property_name": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the property name.", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_$]( *[a-zA-Z0-9_$-])*$"), "Must match :"+"^[a-zA-Z0-9_$]( *[a-zA-Z0-9_$-])*$"),
				},
			},
			"catalog": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the catalog name. The name must match the name of an API catalog in the API collection. The default value is <tt>*</tt> , which indicates that the value applies to all catalogs.", "", "").AddDefaultValue("*").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("*"),
			},
			"value": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the property value.", "", "").String,
				Optional:            true,
			},
		},
	}
	return DmAPIPropertyResourceSchema
}

func (data DmAPIProperty) IsNull() bool {
	if !data.PropertyName.IsNull() {
		return false
	}
	if !data.Catalog.IsNull() {
		return false
	}
	if !data.Value.IsNull() {
		return false
	}
	return true
}

func (data DmAPIProperty) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.PropertyName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PropertyName`, data.PropertyName.ValueString())
	}
	if !data.Catalog.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Catalog`, data.Catalog.ValueString())
	}
	if !data.Value.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Value`, data.Value.ValueString())
	}
	return body
}

func (data *DmAPIProperty) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `PropertyName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PropertyName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PropertyName = types.StringNull()
	}
	if value := res.Get(pathRoot + `Catalog`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Catalog = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Catalog = types.StringValue("*")
	}
	if value := res.Get(pathRoot + `Value`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Value = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Value = types.StringNull()
	}
}

func (data *DmAPIProperty) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `PropertyName`); value.Exists() && !data.PropertyName.IsNull() {
		data.PropertyName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PropertyName = types.StringNull()
	}
	if value := res.Get(pathRoot + `Catalog`); value.Exists() && !data.Catalog.IsNull() {
		data.Catalog = tfutils.ParseStringFromGJSON(value)
	} else if data.Catalog.ValueString() != "*" {
		data.Catalog = types.StringNull()
	}
	if value := res.Get(pathRoot + `Value`); value.Exists() && !data.Value.IsNull() {
		data.Value = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Value = types.StringNull()
	}
}
