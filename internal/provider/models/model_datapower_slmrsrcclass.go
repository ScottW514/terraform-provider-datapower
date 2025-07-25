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
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type SLMRsrcClass struct {
	Id                          types.String `tfsdk:"id"`
	AppDomain                   types.String `tfsdk:"app_domain"`
	UserSummary                 types.String `tfsdk:"user_summary"`
	RsrcType                    types.String `tfsdk:"rsrc_type"`
	RsrcMatchType               types.String `tfsdk:"rsrc_match_type"`
	RsrcValue                   types.List   `tfsdk:"rsrc_value"`
	Stylesheet                  types.String `tfsdk:"stylesheet"`
	XPathFilter                 types.String `tfsdk:"x_path_filter"`
	Subscription                types.String `tfsdk:"subscription"`
	WsrrSubscription            types.String `tfsdk:"wsrr_subscription"`
	WsrrSavedSearchSubscription types.String `tfsdk:"wsrr_saved_search_subscription"`
}

var SLMRsrcClassObjectType = map[string]attr.Type{
	"id":                             types.StringType,
	"app_domain":                     types.StringType,
	"user_summary":                   types.StringType,
	"rsrc_type":                      types.StringType,
	"rsrc_match_type":                types.StringType,
	"rsrc_value":                     types.ListType{ElemType: types.StringType},
	"stylesheet":                     types.StringType,
	"x_path_filter":                  types.StringType,
	"subscription":                   types.StringType,
	"wsrr_subscription":              types.StringType,
	"wsrr_saved_search_subscription": types.StringType,
}

func (data SLMRsrcClass) GetPath() string {
	rest_path := "/mgmt/config/{domain}/SLMRsrcClass"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data SLMRsrcClass) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.RsrcType.IsNull() {
		return false
	}
	if !data.RsrcMatchType.IsNull() {
		return false
	}
	if !data.RsrcValue.IsNull() {
		return false
	}
	if !data.Stylesheet.IsNull() {
		return false
	}
	if !data.XPathFilter.IsNull() {
		return false
	}
	if !data.Subscription.IsNull() {
		return false
	}
	if !data.WsrrSubscription.IsNull() {
		return false
	}
	if !data.WsrrSavedSearchSubscription.IsNull() {
		return false
	}
	return true
}

func (data SLMRsrcClass) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.RsrcType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RsrcType`, data.RsrcType.ValueString())
	}
	if !data.RsrcMatchType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RsrcMatchType`, data.RsrcMatchType.ValueString())
	}
	if !data.RsrcValue.IsNull() {
		var values []string
		data.RsrcValue.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`RsrcValue`+".-1", map[string]string{"value": val})
		}
	}
	if !data.Stylesheet.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Stylesheet`, data.Stylesheet.ValueString())
	}
	if !data.XPathFilter.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XPathFilter`, data.XPathFilter.ValueString())
	}
	if !data.Subscription.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Subscription`, data.Subscription.ValueString())
	}
	if !data.WsrrSubscription.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSRRSubscription`, data.WsrrSubscription.ValueString())
	}
	if !data.WsrrSavedSearchSubscription.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSRRSavedSearchSubscription`, data.WsrrSavedSearchSubscription.ValueString())
	}
	return body
}

func (data *SLMRsrcClass) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `RsrcType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RsrcType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RsrcType = types.StringValue("aaa-mapped-resource")
	}
	if value := res.Get(pathRoot + `RsrcMatchType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RsrcMatchType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RsrcMatchType = types.StringValue("per-extracted-value")
	}
	if value := res.Get(pathRoot + `RsrcValue`); value.Exists() {
		data.RsrcValue = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.RsrcValue = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `Stylesheet`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Stylesheet = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Stylesheet = types.StringNull()
	}
	if value := res.Get(pathRoot + `XPathFilter`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.XPathFilter = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XPathFilter = types.StringNull()
	}
	if value := res.Get(pathRoot + `Subscription`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Subscription = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Subscription = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSRRSubscription`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsrrSubscription = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsrrSubscription = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSRRSavedSearchSubscription`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsrrSavedSearchSubscription = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsrrSavedSearchSubscription = types.StringNull()
	}
}

func (data *SLMRsrcClass) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `RsrcType`); value.Exists() && !data.RsrcType.IsNull() {
		data.RsrcType = tfutils.ParseStringFromGJSON(value)
	} else if data.RsrcType.ValueString() != "aaa-mapped-resource" {
		data.RsrcType = types.StringNull()
	}
	if value := res.Get(pathRoot + `RsrcMatchType`); value.Exists() && !data.RsrcMatchType.IsNull() {
		data.RsrcMatchType = tfutils.ParseStringFromGJSON(value)
	} else if data.RsrcMatchType.ValueString() != "per-extracted-value" {
		data.RsrcMatchType = types.StringNull()
	}
	if value := res.Get(pathRoot + `RsrcValue`); value.Exists() && !data.RsrcValue.IsNull() {
		data.RsrcValue = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.RsrcValue = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `Stylesheet`); value.Exists() && !data.Stylesheet.IsNull() {
		data.Stylesheet = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Stylesheet = types.StringNull()
	}
	if value := res.Get(pathRoot + `XPathFilter`); value.Exists() && !data.XPathFilter.IsNull() {
		data.XPathFilter = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XPathFilter = types.StringNull()
	}
	if value := res.Get(pathRoot + `Subscription`); value.Exists() && !data.Subscription.IsNull() {
		data.Subscription = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Subscription = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSRRSubscription`); value.Exists() && !data.WsrrSubscription.IsNull() {
		data.WsrrSubscription = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsrrSubscription = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSRRSavedSearchSubscription`); value.Exists() && !data.WsrrSavedSearchSubscription.IsNull() {
		data.WsrrSavedSearchSubscription = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsrrSavedSearchSubscription = types.StringNull()
	}
}
