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

type DmWebAppRequestPolicyMap struct {
	Match types.String `tfsdk:"match"`
	Rule  types.String `tfsdk:"rule"`
}

var DmWebAppRequestPolicyMapObjectType = map[string]attr.Type{
	"match": types.StringType,
	"rule":  types.StringType,
}
var DmWebAppRequestPolicyMapObjectDefault = map[string]attr.Value{
	"match": types.StringNull(),
	"rule":  types.StringNull(),
}
var DmWebAppRequestPolicyMapDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"match": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Matching Rule", "", "matching").String,
			Computed:            true,
		},
		"rule": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Request Profile", "", "webapprequest").String,
			Computed:            true,
		},
	},
}
var DmWebAppRequestPolicyMapResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"match": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Matching Rule", "", "matching").String,
			Required:            true,
		},
		"rule": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Request Profile", "", "webapprequest").String,
			Required:            true,
		},
	},
}

func (data DmWebAppRequestPolicyMap) IsNull() bool {
	if !data.Match.IsNull() {
		return false
	}
	if !data.Rule.IsNull() {
		return false
	}
	return true
}

func (data DmWebAppRequestPolicyMap) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Match.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Match`, data.Match.ValueString())
	}
	if !data.Rule.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Rule`, data.Rule.ValueString())
	}
	return body
}

func (data *DmWebAppRequestPolicyMap) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Match`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Match = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Match = types.StringNull()
	}
	if value := res.Get(pathRoot + `Rule`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Rule = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Rule = types.StringNull()
	}
}

func (data *DmWebAppRequestPolicyMap) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Match`); value.Exists() && !data.Match.IsNull() {
		data.Match = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Match = types.StringNull()
	}
	if value := res.Get(pathRoot + `Rule`); value.Exists() && !data.Rule.IsNull() {
		data.Rule = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Rule = types.StringNull()
	}
}
