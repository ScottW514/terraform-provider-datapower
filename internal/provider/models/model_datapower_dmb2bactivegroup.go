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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmB2BActiveGroup struct {
	ProfileGroup types.String `tfsdk:"profile_group"`
	GroupEnabled types.Bool   `tfsdk:"group_enabled"`
}

var DmB2BActiveGroupObjectType = map[string]attr.Type{
	"profile_group": types.StringType,
	"group_enabled": types.BoolType,
}
var DmB2BActiveGroupObjectDefault = map[string]attr.Value{
	"profile_group": types.StringNull(),
	"group_enabled": types.BoolValue(true),
}
var DmB2BActiveGroupDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"profile_group": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Profile group", "group", "b2bprofilegroup").String,
			Computed:            true,
		},
		"group_enabled": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enable group", "enabled", "").AddDefaultValue("true").String,
			Computed:            true,
		},
	},
}
var DmB2BActiveGroupResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"profile_group": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Profile group", "group", "b2bprofilegroup").String,
			Optional:            true,
		},
		"group_enabled": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enable group", "enabled", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
	},
}

func (data DmB2BActiveGroup) IsNull() bool {
	if !data.ProfileGroup.IsNull() {
		return false
	}
	if !data.GroupEnabled.IsNull() {
		return false
	}
	return true
}

func (data DmB2BActiveGroup) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.ProfileGroup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ProfileGroup`, data.ProfileGroup.ValueString())
	}
	if !data.GroupEnabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GroupEnabled`, tfutils.StringFromBool(data.GroupEnabled, ""))
	}
	return body
}

func (data *DmB2BActiveGroup) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `ProfileGroup`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ProfileGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProfileGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `GroupEnabled`); value.Exists() {
		data.GroupEnabled = tfutils.BoolFromString(value.String())
	} else {
		data.GroupEnabled = types.BoolNull()
	}
}

func (data *DmB2BActiveGroup) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `ProfileGroup`); value.Exists() && !data.ProfileGroup.IsNull() {
		data.ProfileGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProfileGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `GroupEnabled`); value.Exists() && !data.GroupEnabled.IsNull() {
		data.GroupEnabled = tfutils.BoolFromString(value.String())
	} else if !data.GroupEnabled.ValueBool() {
		data.GroupEnabled = types.BoolNull()
	}
}
