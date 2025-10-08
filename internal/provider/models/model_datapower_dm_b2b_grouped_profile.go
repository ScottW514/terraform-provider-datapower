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

type DmB2BGroupedProfile struct {
	PartnerProfile types.String `tfsdk:"partner_profile"`
	ProfileDest    types.String `tfsdk:"profile_dest"`
}

var DmB2BGroupedProfileObjectType = map[string]attr.Type{
	"partner_profile": types.StringType,
	"profile_dest":    types.StringType,
}
var DmB2BGroupedProfileObjectDefault = map[string]attr.Value{
	"partner_profile": types.StringNull(),
	"profile_dest":    types.StringNull(),
}

func GetDmB2BGroupedProfileDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmB2BGroupedProfileDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"partner_profile": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the partner profile.",
				Computed:            true,
			},
			"profile_dest": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the destination for this partner. Without this property, the first destination is used.",
				Computed:            true,
			},
		},
	}
	return DmB2BGroupedProfileDataSourceSchema
}
func GetDmB2BGroupedProfileResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmB2BGroupedProfileResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"partner_profile": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the partner profile.", "profile", "b2b_profile").String,
				Required:            true,
			},
			"profile_dest": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the destination for this partner. Without this property, the first destination is used.", "destination", "").String,
				Optional:            true,
			},
		},
	}
	return DmB2BGroupedProfileResourceSchema
}

func (data DmB2BGroupedProfile) IsNull() bool {
	if !data.PartnerProfile.IsNull() {
		return false
	}
	if !data.ProfileDest.IsNull() {
		return false
	}
	return true
}

func (data DmB2BGroupedProfile) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.PartnerProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PartnerProfile`, data.PartnerProfile.ValueString())
	}
	if !data.ProfileDest.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ProfileDest`, data.ProfileDest.ValueString())
	}
	return body
}

func (data *DmB2BGroupedProfile) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `PartnerProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PartnerProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PartnerProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProfileDest`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ProfileDest = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProfileDest = types.StringNull()
	}
}

func (data *DmB2BGroupedProfile) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `PartnerProfile`); value.Exists() && !data.PartnerProfile.IsNull() {
		data.PartnerProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PartnerProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProfileDest`); value.Exists() && !data.ProfileDest.IsNull() {
		data.ProfileDest = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProfileDest = types.StringNull()
	}
}
