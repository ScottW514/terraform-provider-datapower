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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmURLRewriteRule struct {
	Type               types.String `tfsdk:"type"`
	MatchRegexp        types.String `tfsdk:"match_regexp"`
	InputReplaceRegexp types.String `tfsdk:"input_replace_regexp"`
	StyleReplaceRegexp types.String `tfsdk:"style_replace_regexp"`
	InputUnescape      types.Bool   `tfsdk:"input_unescape"`
	StylesheetUnescape types.Bool   `tfsdk:"stylesheet_unescape"`
	Header             types.String `tfsdk:"header"`
	NormalizeUrl       types.Bool   `tfsdk:"normalize_url"`
}

var DmURLRewriteRuleObjectType = map[string]attr.Type{
	"type":                 types.StringType,
	"match_regexp":         types.StringType,
	"input_replace_regexp": types.StringType,
	"style_replace_regexp": types.StringType,
	"input_unescape":       types.BoolType,
	"stylesheet_unescape":  types.BoolType,
	"header":               types.StringType,
	"normalize_url":        types.BoolType,
}
var DmURLRewriteRuleObjectDefault = map[string]attr.Value{
	"type":                 types.StringNull(),
	"match_regexp":         types.StringNull(),
	"input_replace_regexp": types.StringNull(),
	"style_replace_regexp": types.StringNull(),
	"input_unescape":       types.BoolValue(false),
	"stylesheet_unescape":  types.BoolValue(true),
	"header":               types.StringValue("none"),
	"normalize_url":        types.BoolValue(false),
}
var DmURLRewriteRuleDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL Rewrite Type", "type", "").AddStringEnum("rewrite", "absolute-rewrite", "post-body", "header-rewrite", "content-type").String,
			Computed:            true,
		},
		"match_regexp": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Match Expression", "match", "").String,
			Computed:            true,
		},
		"input_replace_regexp": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Input Replace Expression", "input-expression", "").String,
			Computed:            true,
		},
		"style_replace_regexp": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Stylesheet Replace Expression", "stylesheet-expression", "").String,
			Computed:            true,
		},
		"input_unescape": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Input URL Unescape", "input-unescape", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"stylesheet_unescape": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Stylesheet URL Unescape", "stylesheet-unescape", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"header": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Header Name", "", "").AddDefaultValue("none").String,
			Computed:            true,
		},
		"normalize_url": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL Normalization", "normalize-url", "").AddDefaultValue("false").String,
			Computed:            true,
		},
	},
}
var DmURLRewriteRuleResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL Rewrite Type", "type", "").AddStringEnum("rewrite", "absolute-rewrite", "post-body", "header-rewrite", "content-type").String,
			Required:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("rewrite", "absolute-rewrite", "post-body", "header-rewrite", "content-type"),
			},
		},
		"match_regexp": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Match Expression", "match", "").String,
			Required:            true,
		},
		"input_replace_regexp": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Input Replace Expression", "input-expression", "").String,
			Optional:            true,
		},
		"style_replace_regexp": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Stylesheet Replace Expression", "stylesheet-expression", "").String,
			Optional:            true,
		},
		"input_unescape": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Input URL Unescape", "input-unescape", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"stylesheet_unescape": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Stylesheet URL Unescape", "stylesheet-unescape", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"header": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Header Name", "", "").AddDefaultValue("none").String,
			Computed:            true,
			Optional:            true,
			Default:             stringdefault.StaticString("none"),
		},
		"normalize_url": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL Normalization", "normalize-url", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
	},
}

func (data DmURLRewriteRule) IsNull() bool {
	if !data.Type.IsNull() {
		return false
	}
	if !data.MatchRegexp.IsNull() {
		return false
	}
	if !data.InputReplaceRegexp.IsNull() {
		return false
	}
	if !data.StyleReplaceRegexp.IsNull() {
		return false
	}
	if !data.InputUnescape.IsNull() {
		return false
	}
	if !data.StylesheetUnescape.IsNull() {
		return false
	}
	if !data.Header.IsNull() {
		return false
	}
	if !data.NormalizeUrl.IsNull() {
		return false
	}
	return true
}

func (data DmURLRewriteRule) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Type`, data.Type.ValueString())
	}
	if !data.MatchRegexp.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MatchRegexp`, data.MatchRegexp.ValueString())
	}
	if !data.InputReplaceRegexp.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`InputReplaceRegexp`, data.InputReplaceRegexp.ValueString())
	}
	if !data.StyleReplaceRegexp.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StyleReplaceRegexp`, data.StyleReplaceRegexp.ValueString())
	}
	if !data.InputUnescape.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`InputUnescape`, tfutils.StringFromBool(data.InputUnescape))
	}
	if !data.StylesheetUnescape.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StylesheetUnescape`, tfutils.StringFromBool(data.StylesheetUnescape))
	}
	if !data.Header.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Header`, data.Header.ValueString())
	}
	if !data.NormalizeUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`NormalizeURL`, tfutils.StringFromBool(data.NormalizeUrl))
	}
	return body
}

func (data *DmURLRewriteRule) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `MatchRegexp`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MatchRegexp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MatchRegexp = types.StringNull()
	}
	if value := res.Get(pathRoot + `InputReplaceRegexp`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.InputReplaceRegexp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.InputReplaceRegexp = types.StringNull()
	}
	if value := res.Get(pathRoot + `StyleReplaceRegexp`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.StyleReplaceRegexp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.StyleReplaceRegexp = types.StringNull()
	}
	if value := res.Get(pathRoot + `InputUnescape`); value.Exists() {
		data.InputUnescape = tfutils.BoolFromString(value.String())
	} else {
		data.InputUnescape = types.BoolNull()
	}
	if value := res.Get(pathRoot + `StylesheetUnescape`); value.Exists() {
		data.StylesheetUnescape = tfutils.BoolFromString(value.String())
	} else {
		data.StylesheetUnescape = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Header`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Header = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Header = types.StringValue("none")
	}
	if value := res.Get(pathRoot + `NormalizeURL`); value.Exists() {
		data.NormalizeUrl = tfutils.BoolFromString(value.String())
	} else {
		data.NormalizeUrl = types.BoolNull()
	}
}

func (data *DmURLRewriteRule) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && !data.Type.IsNull() {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `MatchRegexp`); value.Exists() && !data.MatchRegexp.IsNull() {
		data.MatchRegexp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MatchRegexp = types.StringNull()
	}
	if value := res.Get(pathRoot + `InputReplaceRegexp`); value.Exists() && !data.InputReplaceRegexp.IsNull() {
		data.InputReplaceRegexp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.InputReplaceRegexp = types.StringNull()
	}
	if value := res.Get(pathRoot + `StyleReplaceRegexp`); value.Exists() && !data.StyleReplaceRegexp.IsNull() {
		data.StyleReplaceRegexp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.StyleReplaceRegexp = types.StringNull()
	}
	if value := res.Get(pathRoot + `InputUnescape`); value.Exists() && !data.InputUnescape.IsNull() {
		data.InputUnescape = tfutils.BoolFromString(value.String())
	} else if data.InputUnescape.ValueBool() {
		data.InputUnescape = types.BoolNull()
	}
	if value := res.Get(pathRoot + `StylesheetUnescape`); value.Exists() && !data.StylesheetUnescape.IsNull() {
		data.StylesheetUnescape = tfutils.BoolFromString(value.String())
	} else if !data.StylesheetUnescape.ValueBool() {
		data.StylesheetUnescape = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Header`); value.Exists() && !data.Header.IsNull() {
		data.Header = tfutils.ParseStringFromGJSON(value)
	} else if data.Header.ValueString() != "none" {
		data.Header = types.StringNull()
	}
	if value := res.Get(pathRoot + `NormalizeURL`); value.Exists() && !data.NormalizeUrl.IsNull() {
		data.NormalizeUrl = tfutils.BoolFromString(value.String())
	} else if data.NormalizeUrl.ValueBool() {
		data.NormalizeUrl = types.BoolNull()
	}
}
