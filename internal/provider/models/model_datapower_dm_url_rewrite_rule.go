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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
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

var DmURLRewriteRuleInputReplaceRegexpCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"header-rewrite"},
}
var DmURLRewriteRuleHeaderCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"header-rewrite"},
}
var DmURLRewriteRuleStyleReplaceRegexpIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"header-rewrite", "content-type"},
}
var DmURLRewriteRuleInputUnescapeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"header-rewrite", "content-type"},
}
var DmURLRewriteRuleStylesheetUnescapeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"header-rewrite", "content-type"},
}
var DmURLRewriteRuleHeaderIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"post-body", "rewrite", "absolute-rewrite", "content-type"},
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

func GetDmURLRewriteRuleDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmURLRewriteRuleDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"type": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the type of rule for the URL Rewrite Policy.", "type", "").AddStringEnum("rewrite", "absolute-rewrite", "post-body", "header-rewrite", "content-type").String,
				Computed:            true,
			},
			"match_regexp": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify a PCRE that defines the match condition that triggers the rewrite rule. Depending on the rule type, a candidate URL or header field is matched against the expression.</p><ul><li>For <b>absolute-rewrite</b> , <b>content-type</b> , and <b>post-body</b> , defines the expression to be matched against the URL. For example, .* or * matches any string, while (.*)xsl=(.*)\\?(.*) matches a text subpattern followed by xsl= followed by a text subpattern followed by a ? followed by a text subpattern.</li><li>For <b>header-rewrite</b> , defines the expression to be matched against the contents of a specific header field. For example, *.* matches any value.</li></ul><p>PCRE documentation is available at http://www.pcre.org</p>", "match", "").String,
				Computed:            true,
			},
			"input_replace_regexp": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify a Perl-style replacement that defines the rewritten URL, header field, or HTTP POST body.</p><ul><li>For <b>absolute-rewrite</b> , defines the rewritten URL. If the match pattern is *, specify the complete replacement. If the match pattern is (.*)xsl=(.*)\\?(.*), specify the evaluation replacement for any text subpattern or retain the original subpattern. To retain the first subpattern, specify $1; to retain the second text subpattern, specify $2; and so forth. To replace the second text subpattern only, specify $1xsl=ident.xsl$3. <p>If a rewritten URL begins with a host name or port that is different from the configured remote address, the host name or port portion of the rewritten URL is ignored.</p></li><li>For <b>content-type</b> , define the replace value for the Content-Type header.</li><li>For <b>header-rewrite</b> , define the replacement value for the specified header.</li><li>For <b>post-body</b> , define the rewritten body of the HTTP POST. If the match pattern is .*, specify the complete replacement. If the match pattern is (.*)xsl=(.*)\\?(.*), specify the evaluation replacement for any text subpattern or retain the original subpattern. To retain the first subpattern, specify $1; to retain the second text subpattern, specify $2; and so forth. To omit the second text subpattern only, specify $1$3.</li></ul>", "input-expression", "").String,
				Computed:            true,
			},
			"style_replace_regexp": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify a Perl-style replacement that identifies the replacement stylesheet. This option is available for <b>absolute-rewrite</b> and <b>post-body</b> only.</p><p>If the match pattern is .*, specify the complete replacement. If the match pattern is (.*)xsl=(.*)\\?(.*), specify the evaluation replacement for any text subpattern or retain the original subpattern. To retain the first subpattern, specify $1; to retain the second text subpattern, specify $2; and so forth. To retain the second text subpattern only and not use the third text subpattern, specify http://mantis:8000$2.</p>", "stylesheet-expression", "").String,
				Computed:            true,
			},
			"input_unescape": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Replace URL-encoded characters (for example, \"%2F\") with the equivalent literal character. Select on to replace escape sequences, or off to retain them.", "input-unescape", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"stylesheet_unescape": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Replace URL-encoded characters (for example, \"%2F\") with the equivalent literal character. Select on to replace escape sequences, or off to retain them.", "stylesheet-unescape", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"header": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name HTTP Header to Rewrite", "", "").AddDefaultValue("none").String,
				Computed:            true,
			},
			"normalize_url": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Normalize URL by converting '\\' to '/' and compressing '.' and '..'", "normalize-url", "").AddDefaultValue("false").String,
				Computed:            true,
			},
		},
	}
	return DmURLRewriteRuleDataSourceSchema
}
func GetDmURLRewriteRuleResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmURLRewriteRuleResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the type of rule for the URL Rewrite Policy.", "type", "").AddStringEnum("rewrite", "absolute-rewrite", "post-body", "header-rewrite", "content-type").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("rewrite", "absolute-rewrite", "post-body", "header-rewrite", "content-type"),
				},
			},
			"match_regexp": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify a PCRE that defines the match condition that triggers the rewrite rule. Depending on the rule type, a candidate URL or header field is matched against the expression.</p><ul><li>For <b>absolute-rewrite</b> , <b>content-type</b> , and <b>post-body</b> , defines the expression to be matched against the URL. For example, .* or * matches any string, while (.*)xsl=(.*)\\?(.*) matches a text subpattern followed by xsl= followed by a text subpattern followed by a ? followed by a text subpattern.</li><li>For <b>header-rewrite</b> , defines the expression to be matched against the contents of a specific header field. For example, *.* matches any value.</li></ul><p>PCRE documentation is available at http://www.pcre.org</p>", "match", "").String,
				Required:            true,
			},
			"input_replace_regexp": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify a Perl-style replacement that defines the rewritten URL, header field, or HTTP POST body.</p><ul><li>For <b>absolute-rewrite</b> , defines the rewritten URL. If the match pattern is *, specify the complete replacement. If the match pattern is (.*)xsl=(.*)\\?(.*), specify the evaluation replacement for any text subpattern or retain the original subpattern. To retain the first subpattern, specify $1; to retain the second text subpattern, specify $2; and so forth. To replace the second text subpattern only, specify $1xsl=ident.xsl$3. <p>If a rewritten URL begins with a host name or port that is different from the configured remote address, the host name or port portion of the rewritten URL is ignored.</p></li><li>For <b>content-type</b> , define the replace value for the Content-Type header.</li><li>For <b>header-rewrite</b> , define the replacement value for the specified header.</li><li>For <b>post-body</b> , define the rewritten body of the HTTP POST. If the match pattern is .*, specify the complete replacement. If the match pattern is (.*)xsl=(.*)\\?(.*), specify the evaluation replacement for any text subpattern or retain the original subpattern. To retain the first subpattern, specify $1; to retain the second text subpattern, specify $2; and so forth. To omit the second text subpattern only, specify $1$3.</li></ul>", "input-expression", "").String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmURLRewriteRuleInputReplaceRegexpCondVal, validators.Evaluation{}, false),
				},
			},
			"style_replace_regexp": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify a Perl-style replacement that identifies the replacement stylesheet. This option is available for <b>absolute-rewrite</b> and <b>post-body</b> only.</p><p>If the match pattern is .*, specify the complete replacement. If the match pattern is (.*)xsl=(.*)\\?(.*), specify the evaluation replacement for any text subpattern or retain the original subpattern. To retain the first subpattern, specify $1; to retain the second text subpattern, specify $2; and so forth. To retain the second text subpattern only and not use the third text subpattern, specify http://mantis:8000$2.</p>", "stylesheet-expression", "").String,
				Optional:            true,
			},
			"input_unescape": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Replace URL-encoded characters (for example, \"%2F\") with the equivalent literal character. Select on to replace escape sequences, or off to retain them.", "input-unescape", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"stylesheet_unescape": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Replace URL-encoded characters (for example, \"%2F\") with the equivalent literal character. Select on to replace escape sequences, or off to retain them.", "stylesheet-unescape", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"header": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name HTTP Header to Rewrite", "", "").AddDefaultValue("none").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmURLRewriteRuleHeaderCondVal, validators.Evaluation{}, true),
				},
				Default: stringdefault.StaticString("none"),
			},
			"normalize_url": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Normalize URL by converting '\\' to '/' and compressing '.' and '..'", "normalize-url", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
	return DmURLRewriteRuleResourceSchema
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
		body, _ = sjson.Set(body, pathRoot+`InputUnescape`, tfutils.StringFromBool(data.InputUnescape, ""))
	}
	if !data.StylesheetUnescape.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StylesheetUnescape`, tfutils.StringFromBool(data.StylesheetUnescape, ""))
	}
	if !data.Header.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Header`, data.Header.ValueString())
	}
	if !data.NormalizeUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`NormalizeURL`, tfutils.StringFromBool(data.NormalizeUrl, ""))
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
