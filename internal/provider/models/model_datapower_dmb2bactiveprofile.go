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

type DmB2BActiveProfile struct {
	PartnerProfile types.String `tfsdk:"partner_profile"`
	ProfileEnabled types.Bool   `tfsdk:"profile_enabled"`
	ProfileDest    types.String `tfsdk:"profile_dest"`
}

var DmB2BActiveProfileObjectType = map[string]attr.Type{
	"partner_profile": types.StringType,
	"profile_enabled": types.BoolType,
	"profile_dest":    types.StringType,
}
var DmB2BActiveProfileObjectDefault = map[string]attr.Value{
	"partner_profile": types.StringNull(),
	"profile_enabled": types.BoolValue(true),
	"profile_dest":    types.StringNull(),
}
var DmB2BActiveProfileDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"partner_profile": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Partner profile", "profile", "b2bprofile").String,
			Computed:            true,
		},
		"profile_enabled": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enable profile", "enabled", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"profile_dest": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Profile destination", "destination", "").String,
			Computed:            true,
		},
	},
}
var DmB2BActiveProfileResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"partner_profile": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Partner profile", "profile", "b2bprofile").String,
			Required:            true,
		},
		"profile_enabled": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enable profile", "enabled", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"profile_dest": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Profile destination", "destination", "").String,
			Optional:            true,
		},
	},
}

func (data DmB2BActiveProfile) IsNull() bool {
	if !data.PartnerProfile.IsNull() {
		return false
	}
	if !data.ProfileEnabled.IsNull() {
		return false
	}
	if !data.ProfileDest.IsNull() {
		return false
	}
	return true
}

func (data DmB2BActiveProfile) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.PartnerProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PartnerProfile`, data.PartnerProfile.ValueString())
	}
	if !data.ProfileEnabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ProfileEnabled`, tfutils.StringFromBool(data.ProfileEnabled, ""))
	}
	if !data.ProfileDest.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ProfileDest`, data.ProfileDest.ValueString())
	}
	return body
}

func (data *DmB2BActiveProfile) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `PartnerProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PartnerProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PartnerProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProfileEnabled`); value.Exists() {
		data.ProfileEnabled = tfutils.BoolFromString(value.String())
	} else {
		data.ProfileEnabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ProfileDest`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ProfileDest = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProfileDest = types.StringNull()
	}
}

func (data *DmB2BActiveProfile) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `PartnerProfile`); value.Exists() && !data.PartnerProfile.IsNull() {
		data.PartnerProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PartnerProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProfileEnabled`); value.Exists() && !data.ProfileEnabled.IsNull() {
		data.ProfileEnabled = tfutils.BoolFromString(value.String())
	} else if !data.ProfileEnabled.ValueBool() {
		data.ProfileEnabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ProfileDest`); value.Exists() && !data.ProfileDest.IsNull() {
		data.ProfileDest = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProfileDest = types.StringNull()
	}
}
