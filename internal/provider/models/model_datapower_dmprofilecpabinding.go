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

type DmProfileCPABinding struct {
	InternalPartner types.String `tfsdk:"internal_partner"`
	Cpa             types.String `tfsdk:"cpa"`
	Collaboration   types.String `tfsdk:"collaboration"`
	Action          types.String `tfsdk:"action"`
}

var DmProfileCPABindingObjectType = map[string]attr.Type{
	"internal_partner": types.StringType,
	"cpa":              types.StringType,
	"collaboration":    types.StringType,
	"action":           types.StringType,
}
var DmProfileCPABindingObjectDefault = map[string]attr.Value{
	"internal_partner": types.StringNull(),
	"cpa":              types.StringNull(),
	"collaboration":    types.StringNull(),
	"action":           types.StringNull(),
}
var DmProfileCPABindingDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"internal_partner": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("CPA sender (internal partner profile)", "internal-partner", "b2bprofile").String,
			Computed:            true,
		},
		"cpa": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("CPA", "cpa", "b2bcpa").String,
			Computed:            true,
		},
		"collaboration": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Service", "collaboration", "b2bcpacollaboration").String,
			Computed:            true,
		},
		"action": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Action", "action", "").String,
			Computed:            true,
		},
	},
}
var DmProfileCPABindingResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"internal_partner": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("CPA sender (internal partner profile)", "internal-partner", "b2bprofile").String,
			Required:            true,
		},
		"cpa": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("CPA", "cpa", "b2bcpa").String,
			Required:            true,
		},
		"collaboration": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Service", "collaboration", "b2bcpacollaboration").String,
			Required:            true,
		},
		"action": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Action", "action", "").String,
			Optional:            true,
		},
	},
}

func (data DmProfileCPABinding) IsNull() bool {
	if !data.InternalPartner.IsNull() {
		return false
	}
	if !data.Cpa.IsNull() {
		return false
	}
	if !data.Collaboration.IsNull() {
		return false
	}
	if !data.Action.IsNull() {
		return false
	}
	return true
}

func (data DmProfileCPABinding) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.InternalPartner.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`InternalPartner`, data.InternalPartner.ValueString())
	}
	if !data.Cpa.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CPA`, data.Cpa.ValueString())
	}
	if !data.Collaboration.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Collaboration`, data.Collaboration.ValueString())
	}
	if !data.Action.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Action`, data.Action.ValueString())
	}
	return body
}

func (data *DmProfileCPABinding) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `InternalPartner`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.InternalPartner = tfutils.ParseStringFromGJSON(value)
	} else {
		data.InternalPartner = types.StringNull()
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
	if value := res.Get(pathRoot + `Action`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Action = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Action = types.StringNull()
	}
}

func (data *DmProfileCPABinding) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `InternalPartner`); value.Exists() && !data.InternalPartner.IsNull() {
		data.InternalPartner = tfutils.ParseStringFromGJSON(value)
	} else {
		data.InternalPartner = types.StringNull()
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
	if value := res.Get(pathRoot + `Action`); value.Exists() && !data.Action.IsNull() {
		data.Action = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Action = types.StringNull()
	}
}
