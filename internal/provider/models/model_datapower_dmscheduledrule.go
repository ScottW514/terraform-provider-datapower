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

type DmScheduledRule struct {
	Rule     types.String `tfsdk:"rule"`
	Interval types.Int64  `tfsdk:"interval"`
}

var DmScheduledRuleObjectType = map[string]attr.Type{
	"rule":     types.StringType,
	"interval": types.Int64Type,
}
var DmScheduledRuleObjectDefault = map[string]attr.Value{
	"rule":     types.StringNull(),
	"interval": types.Int64Null(),
}
var DmScheduledRuleDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"rule": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Rule", "", "stylepolicyrule").String,
			Computed:            true,
		},
		"interval": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Interval", "", "").String,
			Computed:            true,
		},
	},
}
var DmScheduledRuleResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"rule": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Rule", "", "stylepolicyrule").String,
			Required:            true,
		},
		"interval": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Interval", "", "").String,
			Optional:            true,
		},
	},
}

func (data DmScheduledRule) IsNull() bool {
	if !data.Rule.IsNull() {
		return false
	}
	if !data.Interval.IsNull() {
		return false
	}
	return true
}

func (data DmScheduledRule) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Rule.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Rule`, data.Rule.ValueString())
	}
	if !data.Interval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Interval`, data.Interval.ValueInt64())
	}
	return body
}

func (data *DmScheduledRule) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Rule`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Rule = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Rule = types.StringNull()
	}
	if value := res.Get(pathRoot + `Interval`); value.Exists() {
		data.Interval = types.Int64Value(value.Int())
	} else {
		data.Interval = types.Int64Null()
	}
}

func (data *DmScheduledRule) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Rule`); value.Exists() && !data.Rule.IsNull() {
		data.Rule = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Rule = types.StringNull()
	}
	if value := res.Get(pathRoot + `Interval`); value.Exists() && !data.Interval.IsNull() {
		data.Interval = types.Int64Value(value.Int())
	} else {
		data.Interval = types.Int64Null()
	}
}
