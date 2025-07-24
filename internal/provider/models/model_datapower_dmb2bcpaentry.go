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

type DmB2BCPAEntry struct {
	Cpa             types.String `tfsdk:"cpa"`
	Collaboration   types.String `tfsdk:"collaboration"`
	InternalPartner types.String `tfsdk:"internal_partner"`
	ExternalPartner types.String `tfsdk:"external_partner"`
}

var DmB2BCPAEntryObjectType = map[string]attr.Type{
	"cpa":              types.StringType,
	"collaboration":    types.StringType,
	"internal_partner": types.StringType,
	"external_partner": types.StringType,
}
var DmB2BCPAEntryObjectDefault = map[string]attr.Value{
	"cpa":              types.StringNull(),
	"collaboration":    types.StringNull(),
	"internal_partner": types.StringNull(),
	"external_partner": types.StringNull(),
}
var DmB2BCPAEntryDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"cpa": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("CPA", "cpa", "b2bcpa").String,
			Computed:            true,
		},
		"collaboration": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Service", "collaboration", "b2bcpacollaboration").String,
			Computed:            true,
		},
		"internal_partner": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Internal partner", "internal-partner", "b2bprofile").String,
			Computed:            true,
		},
		"external_partner": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("External partner", "external-partner", "b2bprofile").String,
			Computed:            true,
		},
	},
}
var DmB2BCPAEntryResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"cpa": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("CPA", "cpa", "b2bcpa").String,
			Required:            true,
		},
		"collaboration": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Service", "collaboration", "b2bcpacollaboration").String,
			Required:            true,
		},
		"internal_partner": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Internal partner", "internal-partner", "b2bprofile").String,
			Required:            true,
		},
		"external_partner": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("External partner", "external-partner", "b2bprofile").String,
			Required:            true,
		},
	},
}

func (data DmB2BCPAEntry) IsNull() bool {
	if !data.Cpa.IsNull() {
		return false
	}
	if !data.Collaboration.IsNull() {
		return false
	}
	if !data.InternalPartner.IsNull() {
		return false
	}
	if !data.ExternalPartner.IsNull() {
		return false
	}
	return true
}

func (data DmB2BCPAEntry) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Cpa.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CPA`, data.Cpa.ValueString())
	}
	if !data.Collaboration.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Collaboration`, data.Collaboration.ValueString())
	}
	if !data.InternalPartner.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`InternalPartner`, data.InternalPartner.ValueString())
	}
	if !data.ExternalPartner.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ExternalPartner`, data.ExternalPartner.ValueString())
	}
	return body
}

func (data *DmB2BCPAEntry) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `CPA`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Cpa = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Cpa = types.StringNull()
	}
	if value := res.Get(pathRoot + `Collaboration`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Collaboration = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Collaboration = types.StringNull()
	}
	if value := res.Get(pathRoot + `InternalPartner`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.InternalPartner = tfutils.ParseStringFromGJSON(value)
	} else {
		data.InternalPartner = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExternalPartner`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ExternalPartner = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExternalPartner = types.StringNull()
	}
}

func (data *DmB2BCPAEntry) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `CPA`); value.Exists() && !data.Cpa.IsNull() {
		data.Cpa = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Cpa = types.StringNull()
	}
	if value := res.Get(pathRoot + `Collaboration`); value.Exists() && !data.Collaboration.IsNull() {
		data.Collaboration = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Collaboration = types.StringNull()
	}
	if value := res.Get(pathRoot + `InternalPartner`); value.Exists() && !data.InternalPartner.IsNull() {
		data.InternalPartner = tfutils.ParseStringFromGJSON(value)
	} else {
		data.InternalPartner = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExternalPartner`); value.Exists() && !data.ExternalPartner.IsNull() {
		data.ExternalPartner = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExternalPartner = types.StringNull()
	}
}
