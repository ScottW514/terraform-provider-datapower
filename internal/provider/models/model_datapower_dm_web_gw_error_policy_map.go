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

type DmWebGWErrorPolicyMap struct {
	Match  types.String `tfsdk:"match"`
	Action types.String `tfsdk:"action"`
}

var DmWebGWErrorPolicyMapObjectType = map[string]attr.Type{
	"match":  types.StringType,
	"action": types.StringType,
}
var DmWebGWErrorPolicyMapObjectDefault = map[string]attr.Value{
	"match":  types.StringNull(),
	"action": types.StringNull(),
}

func GetDmWebGWErrorPolicyMapDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmWebGWErrorPolicyMapDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"match": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Select the matching rule that defines the matching criteria.",
				Computed:            true,
			},
			"action": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Select the error action that handles the errors and generates responses.",
				Computed:            true,
			},
		},
	}
	return DmWebGWErrorPolicyMapDataSourceSchema
}
func GetDmWebGWErrorPolicyMapResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmWebGWErrorPolicyMapResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"match": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the matching rule that defines the matching criteria.", "", "matching").String,
				Required:            true,
			},
			"action": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the error action that handles the errors and generates responses.", "", "mpgw_error_action").String,
				Required:            true,
			},
		},
	}
	return DmWebGWErrorPolicyMapResourceSchema
}

func (data DmWebGWErrorPolicyMap) IsNull() bool {
	if !data.Match.IsNull() {
		return false
	}
	if !data.Action.IsNull() {
		return false
	}
	return true
}

func (data DmWebGWErrorPolicyMap) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Match.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Match`, data.Match.ValueString())
	}
	if !data.Action.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Action`, data.Action.ValueString())
	}
	return body
}

func (data *DmWebGWErrorPolicyMap) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Match`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Match = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Match = types.StringNull()
	}
	if value := res.Get(pathRoot + `Action`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Action = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Action = types.StringNull()
	}
}

func (data *DmWebGWErrorPolicyMap) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Match`); value.Exists() && !data.Match.IsNull() {
		data.Match = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Match = types.StringNull()
	}
	if value := res.Get(pathRoot + `Action`); value.Exists() && !data.Action.IsNull() {
		data.Action = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Action = types.StringNull()
	}
}
