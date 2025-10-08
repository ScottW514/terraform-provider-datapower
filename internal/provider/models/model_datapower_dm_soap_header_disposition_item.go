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

type DmSOAPHeaderDispositionItem struct {
	Namespace      types.String `tfsdk:"namespace"`
	LocalName      types.String `tfsdk:"local_name"`
	ChildLocalName types.String `tfsdk:"child_local_name"`
	Action         types.String `tfsdk:"action"`
}

var DmSOAPHeaderDispositionItemObjectType = map[string]attr.Type{
	"namespace":        types.StringType,
	"local_name":       types.StringType,
	"child_local_name": types.StringType,
	"action":           types.StringType,
}
var DmSOAPHeaderDispositionItemObjectDefault = map[string]attr.Value{
	"namespace":        types.StringNull(),
	"local_name":       types.StringNull(),
	"child_local_name": types.StringNull(),
	"action":           types.StringValue("processed"),
}

func GetDmSOAPHeaderDispositionItemDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmSOAPHeaderDispositionItemDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"namespace": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the namespace URI of the SOAP header element, the default value is a blank string, which indicates no restriction.",
				Computed:            true,
			},
			"local_name": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the local name of the SOAP header element, the default value is a blank string, which indicates no restriction.",
				Computed:            true,
			},
			"child_local_name": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the local name of the SOAP header's child element, the default value is a blank string, which indicates no restriction.",
				Computed:            true,
			},
			"action": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify what action to take for this SOAP header and/or child element.",
				Computed:            true,
			},
		},
	}
	return DmSOAPHeaderDispositionItemDataSourceSchema
}
func GetDmSOAPHeaderDispositionItemResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmSOAPHeaderDispositionItemResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"namespace": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the namespace URI of the SOAP header element, the default value is a blank string, which indicates no restriction.", "", "").String,
				Optional:            true,
			},
			"local_name": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the local name of the SOAP header element, the default value is a blank string, which indicates no restriction.", "", "").String,
				Optional:            true,
			},
			"child_local_name": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the local name of the SOAP header's child element, the default value is a blank string, which indicates no restriction.", "", "").String,
				Optional:            true,
			},
			"action": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify what action to take for this SOAP header and/or child element.", "", "").AddStringEnum("processed", "unprocessed", "keep", "remove", "fault").AddDefaultValue("processed").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("processed", "unprocessed", "keep", "remove", "fault"),
				},
				Default: stringdefault.StaticString("processed"),
			},
		},
	}
	return DmSOAPHeaderDispositionItemResourceSchema
}

func (data DmSOAPHeaderDispositionItem) IsNull() bool {
	if !data.Namespace.IsNull() {
		return false
	}
	if !data.LocalName.IsNull() {
		return false
	}
	if !data.ChildLocalName.IsNull() {
		return false
	}
	if !data.Action.IsNull() {
		return false
	}
	return true
}

func (data DmSOAPHeaderDispositionItem) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Namespace.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Namespace`, data.Namespace.ValueString())
	}
	if !data.LocalName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalName`, data.LocalName.ValueString())
	}
	if !data.ChildLocalName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ChildLocalName`, data.ChildLocalName.ValueString())
	}
	if !data.Action.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Action`, data.Action.ValueString())
	}
	return body
}

func (data *DmSOAPHeaderDispositionItem) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Namespace`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Namespace = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Namespace = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalName = types.StringNull()
	}
	if value := res.Get(pathRoot + `ChildLocalName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ChildLocalName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ChildLocalName = types.StringNull()
	}
	if value := res.Get(pathRoot + `Action`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Action = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Action = types.StringValue("processed")
	}
}

func (data *DmSOAPHeaderDispositionItem) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Namespace`); value.Exists() && !data.Namespace.IsNull() {
		data.Namespace = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Namespace = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalName`); value.Exists() && !data.LocalName.IsNull() {
		data.LocalName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalName = types.StringNull()
	}
	if value := res.Get(pathRoot + `ChildLocalName`); value.Exists() && !data.ChildLocalName.IsNull() {
		data.ChildLocalName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ChildLocalName = types.StringNull()
	}
	if value := res.Get(pathRoot + `Action`); value.Exists() && !data.Action.IsNull() {
		data.Action = tfutils.ParseStringFromGJSON(value)
	} else if data.Action.ValueString() != "processed" {
		data.Action = types.StringNull()
	}
}
