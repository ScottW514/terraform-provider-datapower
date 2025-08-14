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

type DmURLMapRule struct {
	Pattern types.String `tfsdk:"pattern"`
}

var DmURLMapRuleObjectType = map[string]attr.Type{
	"pattern": types.StringType,
}
var DmURLMapRuleObjectDefault = map[string]attr.Value{
	"pattern": types.StringNull(),
}
var DmURLMapRuleDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"pattern": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("A shell-style match pattern that defines the URL set contained within the URL map. The following wildcard characters are available when constructing a match pattern. <table><tr><td valign=\"top\">asterisk (*)</td><td valign=\"top\">Matches 0 or more occurrences of any character</td></tr><tr><td valign=\"top\">question mark (?)</td><td valign=\"top\">Matches one occurrence of any single character</td></tr><tr><td valign=\"top\">brackets ( [ ] )</td><td valign=\"top\">Defines a character or numeric range. For example, [1-5] matches 1, 2, 3, 4, or 5, while xs[dl] matches xsd or xsl.</td></tr></table>", "", "").String,
			Computed:            true,
		},
	},
}
var DmURLMapRuleResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"pattern": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("A shell-style match pattern that defines the URL set contained within the URL map. The following wildcard characters are available when constructing a match pattern. <table><tr><td valign=\"top\">asterisk (*)</td><td valign=\"top\">Matches 0 or more occurrences of any character</td></tr><tr><td valign=\"top\">question mark (?)</td><td valign=\"top\">Matches one occurrence of any single character</td></tr><tr><td valign=\"top\">brackets ( [ ] )</td><td valign=\"top\">Defines a character or numeric range. For example, [1-5] matches 1, 2, 3, 4, or 5, while xs[dl] matches xsd or xsl.</td></tr></table>", "", "").String,
			Required:            true,
		},
	},
}

func (data DmURLMapRule) IsNull() bool {
	if !data.Pattern.IsNull() {
		return false
	}
	return true
}

func (data DmURLMapRule) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Pattern.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Pattern`, data.Pattern.ValueString())
	}
	return body
}

func (data *DmURLMapRule) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Pattern`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Pattern = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Pattern = types.StringNull()
	}
}

func (data *DmURLMapRule) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Pattern`); value.Exists() && !data.Pattern.IsNull() {
		data.Pattern = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Pattern = types.StringNull()
	}
}
