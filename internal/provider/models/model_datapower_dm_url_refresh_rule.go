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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmURLRefreshRule struct {
	UrlMap             types.String `tfsdk:"url_map"`
	UrlRefreshPolicy   types.String `tfsdk:"url_refresh_policy"`
	UrlRefreshInterval types.Int64  `tfsdk:"url_refresh_interval"`
}

var DmURLRefreshRuleURLRefreshIntervalCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "url_refresh_policy",
	AttrType:    "String",
	AttrDefault: "default",
	Value:       []string{"default", "no-flush", "protocol-specified"},
}

var DmURLRefreshRuleObjectType = map[string]attr.Type{
	"url_map":              types.StringType,
	"url_refresh_policy":   types.StringType,
	"url_refresh_interval": types.Int64Type,
}
var DmURLRefreshRuleObjectDefault = map[string]attr.Value{
	"url_map":              types.StringNull(),
	"url_refresh_policy":   types.StringValue("default"),
	"url_refresh_interval": types.Int64Value(0),
}

func GetDmURLRefreshRuleDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmURLRefreshRuleDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"url_map": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("URL maps contain one or more shell-style (wildcard) match patterns. Use the values list to select the URL map that supplies the match criteria for the URL Refresh Policy.", "urlmap", "url_map").String,
				Computed:            true,
			},
			"url_refresh_policy": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how to cache a stylesheet that is obtained with a URL refresh operation. The default is default.", "type", "").AddStringEnum("default", "no-cache", "no-flush", "protocol-specified").AddDefaultValue("default").String,
				Computed:            true,
			},
			"url_refresh_interval": DataSourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Not used when the refresh rule type is no-cache, specifies the update frequency for stylesheets that fulfill the match criteria.", "interval", "").AddDefaultValue("0").String,
				Computed:            true,
			},
		},
	}
	return DmURLRefreshRuleDataSourceSchema
}
func GetDmURLRefreshRuleResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmURLRefreshRuleResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"url_map": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("URL maps contain one or more shell-style (wildcard) match patterns. Use the values list to select the URL map that supplies the match criteria for the URL Refresh Policy.", "urlmap", "url_map").String,
				Required:            true,
			},
			"url_refresh_policy": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how to cache a stylesheet that is obtained with a URL refresh operation. The default is default.", "type", "").AddStringEnum("default", "no-cache", "no-flush", "protocol-specified").AddDefaultValue("default").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default", "no-cache", "no-flush", "protocol-specified"),
				},
				Default: stringdefault.StaticString("default"),
			},
			"url_refresh_interval": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Not used when the refresh rule type is no-cache, specifies the update frequency for stylesheets that fulfill the match criteria.", "interval", "").AddDefaultValue("0").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					validators.ConditionalRequiredInt64(DmURLRefreshRuleURLRefreshIntervalCondVal, validators.Evaluation{}, true),
				},
				Default: int64default.StaticInt64(0),
			},
		},
	}
	return DmURLRefreshRuleResourceSchema
}

func (data DmURLRefreshRule) IsNull() bool {
	if !data.UrlMap.IsNull() {
		return false
	}
	if !data.UrlRefreshPolicy.IsNull() {
		return false
	}
	if !data.UrlRefreshInterval.IsNull() {
		return false
	}
	return true
}

func (data DmURLRefreshRule) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.UrlMap.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`URLMap`, data.UrlMap.ValueString())
	}
	if !data.UrlRefreshPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`URLRefreshPolicy`, data.UrlRefreshPolicy.ValueString())
	}
	if !data.UrlRefreshInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`URLRefreshInterval`, data.UrlRefreshInterval.ValueInt64())
	}
	return body
}

func (data *DmURLRefreshRule) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `URLMap`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UrlMap = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UrlMap = types.StringNull()
	}
	if value := res.Get(pathRoot + `URLRefreshPolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UrlRefreshPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UrlRefreshPolicy = types.StringValue("default")
	}
	if value := res.Get(pathRoot + `URLRefreshInterval`); value.Exists() {
		data.UrlRefreshInterval = types.Int64Value(value.Int())
	} else {
		data.UrlRefreshInterval = types.Int64Value(0)
	}
}

func (data *DmURLRefreshRule) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `URLMap`); value.Exists() && !data.UrlMap.IsNull() {
		data.UrlMap = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UrlMap = types.StringNull()
	}
	if value := res.Get(pathRoot + `URLRefreshPolicy`); value.Exists() && !data.UrlRefreshPolicy.IsNull() {
		data.UrlRefreshPolicy = tfutils.ParseStringFromGJSON(value)
	} else if data.UrlRefreshPolicy.ValueString() != "default" {
		data.UrlRefreshPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `URLRefreshInterval`); value.Exists() && !data.UrlRefreshInterval.IsNull() {
		data.UrlRefreshInterval = types.Int64Value(value.Int())
	} else if data.UrlRefreshInterval.ValueInt64() != 0 {
		data.UrlRefreshInterval = types.Int64Null()
	}
}
