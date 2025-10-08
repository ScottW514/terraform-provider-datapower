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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmRoutingPrefix struct {
	Type types.String `tfsdk:"type"`
	Name types.String `tfsdk:"name"`
}

var DmRoutingPrefixNameCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "uri",
	Value:       []string{"uri"},
}

var DmRoutingPrefixObjectType = map[string]attr.Type{
	"type": types.StringType,
	"name": types.StringType,
}
var DmRoutingPrefixObjectDefault = map[string]attr.Value{
	"type": types.StringValue("uri"),
	"name": types.StringNull(),
}

func GetDmRoutingPrefixDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmRoutingPrefixDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"type": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the type for the routing prefix.",
				Computed:            true,
			},
			"name": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the routing prefix for API collection routing. <ul><li>When URI, the routing prefix is case sensitive and must begin but not end with a slash (/).</li><li>When hostname, the prefix must not start or end with period (.). Although the request uses the domain qualified hostname, specify only the hostname.</li></ul>",
				Computed:            true,
			},
		},
	}
	return DmRoutingPrefixDataSourceSchema
}
func GetDmRoutingPrefixResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmRoutingPrefixResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the type for the routing prefix.", "", "").AddStringEnum("uri", "host").AddDefaultValue("uri").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("uri", "host"),
				},
				Default: stringdefault.StaticString("uri"),
			},
			"name": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the routing prefix for API collection routing. <ul><li>When URI, the routing prefix is case sensitive and must begin but not end with a slash (/).</li><li>When hostname, the prefix must not start or end with period (.). Although the request uses the domain qualified hostname, specify only the hostname.</li></ul>", "", "").AddRequiredWhen(DmRoutingPrefixNameCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmRoutingPrefixNameCondVal, validators.Evaluation{}, false),
				},
			},
		},
	}
	return DmRoutingPrefixResourceSchema
}

func (data DmRoutingPrefix) IsNull() bool {
	if !data.Type.IsNull() {
		return false
	}
	if !data.Name.IsNull() {
		return false
	}
	return true
}

func (data DmRoutingPrefix) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Type`, data.Type.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Name`, data.Name.ValueString())
	}
	return body
}

func (data *DmRoutingPrefix) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringValue("uri")
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
}

func (data *DmRoutingPrefix) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && !data.Type.IsNull() {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else if data.Type.ValueString() != "uri" {
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && !data.Name.IsNull() {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
}
