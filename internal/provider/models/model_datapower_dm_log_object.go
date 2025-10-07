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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmLogObject struct {
	Class            types.String `tfsdk:"class"`
	Object           types.String `tfsdk:"object"`
	FollowReferences types.Bool   `tfsdk:"follow_references"`
}

var DmLogObjectObjectType = map[string]attr.Type{
	"class":             types.StringType,
	"object":            types.StringType,
	"follow_references": types.BoolType,
}
var DmLogObjectObjectDefault = map[string]attr.Value{
	"class":             types.StringValue("AAAPolicy"),
	"object":            types.StringNull(),
	"follow_references": types.BoolValue(false),
}

func GetDmLogObjectDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmLogObjectDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"class": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the object type, which is the object class. With this filter, the log target collects log messages for only the specified object classes or for only particular instances of the specified object class.", "", "").AddDefaultValue("AAAPolicy").String,
				Computed:            true,
			},
			"object": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the instance name of the specified object type. <ul><li>For all instances of an object class, do not specify an object name.</li><li>For a specific instance of an object class, specify its object name.</li></ul>", "", "").String,
				Computed:            true,
			},
			"follow_references": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to include log messages for objects that the specified object instance references. <ul><li>When enabled, include referenced objects.</li><li>When disabled, exclude referenced objects.</li></ul><p><b>Note:</b> Included objects are a static snapshot when you apply the object filter. If referenced objects are added after you apply the object filter, messages for these referenced objects are not logged.</p>", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
		},
	}
	return DmLogObjectDataSourceSchema
}
func GetDmLogObjectResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmLogObjectResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"class": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the object type, which is the object class. With this filter, the log target collects log messages for only the specified object classes or for only particular instances of the specified object class.", "", "").AddDefaultValue("AAAPolicy").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("AAAPolicy"),
			},
			"object": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the instance name of the specified object type. <ul><li>For all instances of an object class, do not specify an object name.</li><li>For a specific instance of an object class, specify its object name.</li></ul>", "", "").String,
				Optional:            true,
			},
			"follow_references": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to include log messages for objects that the specified object instance references. <ul><li>When enabled, include referenced objects.</li><li>When disabled, exclude referenced objects.</li></ul><p><b>Note:</b> Included objects are a static snapshot when you apply the object filter. If referenced objects are added after you apply the object filter, messages for these referenced objects are not logged.</p>", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
	return DmLogObjectResourceSchema
}

func (data DmLogObject) IsNull() bool {
	if !data.Class.IsNull() {
		return false
	}
	if !data.Object.IsNull() {
		return false
	}
	if !data.FollowReferences.IsNull() {
		return false
	}
	return true
}

func (data DmLogObject) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Class.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Class`, data.Class.ValueString())
	}
	if !data.Object.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Object`, data.Object.ValueString())
	}
	if !data.FollowReferences.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FollowReferences`, tfutils.StringFromBool(data.FollowReferences, ""))
	}
	return body
}

func (data *DmLogObject) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Class`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Class = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Class = types.StringValue("AAAPolicy")
	}
	if value := res.Get(pathRoot + `Object`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Object = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Object = types.StringNull()
	}
	if value := res.Get(pathRoot + `FollowReferences`); value.Exists() {
		data.FollowReferences = tfutils.BoolFromString(value.String())
	} else {
		data.FollowReferences = types.BoolNull()
	}
}

func (data *DmLogObject) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Class`); value.Exists() && !data.Class.IsNull() {
		data.Class = tfutils.ParseStringFromGJSON(value)
	} else if data.Class.ValueString() != "AAAPolicy" {
		data.Class = types.StringNull()
	}
	if value := res.Get(pathRoot + `Object`); value.Exists() && !data.Object.IsNull() {
		data.Object = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Object = types.StringNull()
	}
	if value := res.Get(pathRoot + `FollowReferences`); value.Exists() && !data.FollowReferences.IsNull() {
		data.FollowReferences = tfutils.BoolFromString(value.String())
	} else if data.FollowReferences.ValueBool() {
		data.FollowReferences = types.BoolNull()
	}
}
