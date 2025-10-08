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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmAAATransactionPriority struct {
	Credential    types.String `tfsdk:"credential"`
	Priority      types.String `tfsdk:"priority"`
	Authorization types.Bool   `tfsdk:"authorization"`
}

var DmAAATransactionPriorityObjectType = map[string]attr.Type{
	"credential":    types.StringType,
	"priority":      types.StringType,
	"authorization": types.BoolType,
}
var DmAAATransactionPriorityObjectDefault = map[string]attr.Value{
	"credential":    types.StringNull(),
	"priority":      types.StringNull(),
	"authorization": types.BoolValue(false),
}

func GetDmAAATransactionPriorityDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmAAATransactionPriorityDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"credential": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the name of the output credentials.",
				Computed:            true,
			},
			"priority": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the priority for scheduling or resource allocation.",
				Computed:            true,
			},
			"authorization": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Specify whether to require authorization. By default, authorization is not required.",
				Computed:            true,
			},
		},
	}
	return DmAAATransactionPriorityDataSourceSchema
}
func GetDmAAATransactionPriorityResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmAAATransactionPriorityResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"credential": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the output credentials.", "", "").String,
				Required:            true,
			},
			"priority": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the priority for scheduling or resource allocation.", "", "").AddStringEnum("unknown", "high-min", "high", "high-max", "normal-min", "normal", "normal-max", "low-min", "low", "low-max").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("unknown", "high-min", "high", "high-max", "normal-min", "normal", "normal-max", "low-min", "low", "low-max"),
				},
			},
			"authorization": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to require authorization. By default, authorization is not required.", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
	return DmAAATransactionPriorityResourceSchema
}

func (data DmAAATransactionPriority) IsNull() bool {
	if !data.Credential.IsNull() {
		return false
	}
	if !data.Priority.IsNull() {
		return false
	}
	if !data.Authorization.IsNull() {
		return false
	}
	return true
}

func (data DmAAATransactionPriority) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Credential.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Credential`, data.Credential.ValueString())
	}
	if !data.Priority.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Priority`, data.Priority.ValueString())
	}
	if !data.Authorization.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Authorization`, tfutils.StringFromBool(data.Authorization, ""))
	}
	return body
}

func (data *DmAAATransactionPriority) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Credential`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Credential = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Credential = types.StringNull()
	}
	if value := res.Get(pathRoot + `Priority`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Priority = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Priority = types.StringNull()
	}
	if value := res.Get(pathRoot + `Authorization`); value.Exists() {
		data.Authorization = tfutils.BoolFromString(value.String())
	} else {
		data.Authorization = types.BoolNull()
	}
}

func (data *DmAAATransactionPriority) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Credential`); value.Exists() && !data.Credential.IsNull() {
		data.Credential = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Credential = types.StringNull()
	}
	if value := res.Get(pathRoot + `Priority`); value.Exists() && !data.Priority.IsNull() {
		data.Priority = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Priority = types.StringNull()
	}
	if value := res.Get(pathRoot + `Authorization`); value.Exists() && !data.Authorization.IsNull() {
		data.Authorization = tfutils.BoolFromString(value.String())
	} else if data.Authorization.ValueBool() {
		data.Authorization = types.BoolNull()
	}
}
