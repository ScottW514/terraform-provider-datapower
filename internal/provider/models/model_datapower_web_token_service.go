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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type WebTokenService struct {
	Id                     types.String                `tfsdk:"id"`
	AppDomain              types.String                `tfsdk:"app_domain"`
	UserSummary            types.String                `tfsdk:"user_summary"`
	Priority               types.String                `tfsdk:"priority"`
	XmlManager             types.String                `tfsdk:"xml_manager"`
	RequestType            types.String                `tfsdk:"request_type"`
	FrontSide              types.List                  `tfsdk:"front_side"`
	StylePolicy            types.String                `tfsdk:"style_policy"`
	RewriteErrors          types.Bool                  `tfsdk:"rewrite_errors"`
	DelayErrors            types.Bool                  `tfsdk:"delay_errors"`
	DelayErrorsDuration    types.Int64                 `tfsdk:"delay_errors_duration"`
	FrontTimeout           types.Int64                 `tfsdk:"front_timeout"`
	FrontPersistentTimeout types.Int64                 `tfsdk:"front_persistent_timeout"`
	FrontHttpVersion       types.String                `tfsdk:"front_http_version"`
	HttpClientIpLabel      types.String                `tfsdk:"http_client_ip_label"`
	HttpLogCorIdLabel      types.String                `tfsdk:"http_log_cor_id_label"`
	DebugMode              types.String                `tfsdk:"debug_mode"`
	DebugHistory           types.Int64                 `tfsdk:"debug_history"`
	DependencyActions      []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var WebTokenServiceDelayErrorsDurationCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "delay_errors",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "rewrite_errors",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
	},
}
var WebTokenServiceDebugHistoryCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "debug_mode",
	AttrType:    "String",
	AttrDefault: "off",
	Value:       []string{"true"},
}
var WebTokenServiceDelayErrorsIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "rewrite_errors",
	AttrType:    "Bool",
	AttrDefault: "true",
	Value:       []string{"true"},
}
var WebTokenServiceDelayErrorsDurationIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var WebTokenServiceDebugHistoryIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var WebTokenServiceObjectType = map[string]attr.Type{
	"id":                       types.StringType,
	"app_domain":               types.StringType,
	"user_summary":             types.StringType,
	"priority":                 types.StringType,
	"xml_manager":              types.StringType,
	"request_type":             types.StringType,
	"front_side":               types.ListType{ElemType: types.ObjectType{AttrTypes: DmSSLFrontSideObjectType}},
	"style_policy":             types.StringType,
	"rewrite_errors":           types.BoolType,
	"delay_errors":             types.BoolType,
	"delay_errors_duration":    types.Int64Type,
	"front_timeout":            types.Int64Type,
	"front_persistent_timeout": types.Int64Type,
	"front_http_version":       types.StringType,
	"http_client_ip_label":     types.StringType,
	"http_log_cor_id_label":    types.StringType,
	"debug_mode":               types.StringType,
	"debug_history":            types.Int64Type,
	"dependency_actions":       actions.ActionsListType,
}

func (data WebTokenService) GetPath() string {
	rest_path := "/mgmt/config/{domain}/WebTokenService"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data WebTokenService) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Priority.IsNull() {
		return false
	}
	if !data.XmlManager.IsNull() {
		return false
	}
	if !data.RequestType.IsNull() {
		return false
	}
	if !data.FrontSide.IsNull() {
		return false
	}
	if !data.StylePolicy.IsNull() {
		return false
	}
	if !data.RewriteErrors.IsNull() {
		return false
	}
	if !data.DelayErrors.IsNull() {
		return false
	}
	if !data.DelayErrorsDuration.IsNull() {
		return false
	}
	if !data.FrontTimeout.IsNull() {
		return false
	}
	if !data.FrontPersistentTimeout.IsNull() {
		return false
	}
	if !data.FrontHttpVersion.IsNull() {
		return false
	}
	if !data.HttpClientIpLabel.IsNull() {
		return false
	}
	if !data.HttpLogCorIdLabel.IsNull() {
		return false
	}
	if !data.DebugMode.IsNull() {
		return false
	}
	if !data.DebugHistory.IsNull() {
		return false
	}
	return true
}

func (data WebTokenService) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Priority.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Priority`, data.Priority.ValueString())
	}
	if !data.XmlManager.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XMLManager`, data.XmlManager.ValueString())
	}
	if !data.RequestType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RequestType`, data.RequestType.ValueString())
	}
	if !data.FrontSide.IsNull() {
		var dataValues []DmSSLFrontSide
		data.FrontSide.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`FrontSide`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.StylePolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StylePolicy`, data.StylePolicy.ValueString())
	}
	if !data.RewriteErrors.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RewriteErrors`, tfutils.StringFromBool(data.RewriteErrors, ""))
	}
	if !data.DelayErrors.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DelayErrors`, tfutils.StringFromBool(data.DelayErrors, ""))
	}
	if !data.DelayErrorsDuration.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DelayErrorsDuration`, data.DelayErrorsDuration.ValueInt64())
	}
	if !data.FrontTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FrontTimeout`, data.FrontTimeout.ValueInt64())
	}
	if !data.FrontPersistentTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FrontPersistentTimeout`, data.FrontPersistentTimeout.ValueInt64())
	}
	if !data.FrontHttpVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FrontHTTPVersion`, data.FrontHttpVersion.ValueString())
	}
	if !data.HttpClientIpLabel.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTPClientIPLabel`, data.HttpClientIpLabel.ValueString())
	}
	if !data.HttpLogCorIdLabel.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTPLogCorIDLabel`, data.HttpLogCorIdLabel.ValueString())
	}
	if !data.DebugMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DebugMode`, data.DebugMode.ValueString())
	}
	if !data.DebugHistory.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DebugHistory`, data.DebugHistory.ValueInt64())
	}
	return body
}

func (data *WebTokenService) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Priority`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Priority = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Priority = types.StringValue("normal")
	}
	if value := res.Get(pathRoot + `XMLManager`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.XmlManager = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XmlManager = types.StringValue("default")
	}
	if value := res.Get(pathRoot + `RequestType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RequestType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RequestType = types.StringValue("preprocessed")
	}
	if value := res.Get(pathRoot + `FrontSide`); value.Exists() {
		l := []DmSSLFrontSide{}
		if value := res.Get(`FrontSide`); value.Exists() {
			for _, v := range value.Array() {
				item := DmSSLFrontSide{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.FrontSide, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSSLFrontSideObjectType}, l)
		} else {
			data.FrontSide = types.ListNull(types.ObjectType{AttrTypes: DmSSLFrontSideObjectType})
		}
	} else {
		data.FrontSide = types.ListNull(types.ObjectType{AttrTypes: DmSSLFrontSideObjectType})
	}
	if value := res.Get(pathRoot + `StylePolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.StylePolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.StylePolicy = types.StringValue("default")
	}
	if value := res.Get(pathRoot + `RewriteErrors`); value.Exists() {
		data.RewriteErrors = tfutils.BoolFromString(value.String())
	} else {
		data.RewriteErrors = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DelayErrors`); value.Exists() {
		data.DelayErrors = tfutils.BoolFromString(value.String())
	} else {
		data.DelayErrors = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DelayErrorsDuration`); value.Exists() {
		data.DelayErrorsDuration = types.Int64Value(value.Int())
	} else {
		data.DelayErrorsDuration = types.Int64Value(1000)
	}
	if value := res.Get(pathRoot + `FrontTimeout`); value.Exists() {
		data.FrontTimeout = types.Int64Value(value.Int())
	} else {
		data.FrontTimeout = types.Int64Value(120)
	}
	if value := res.Get(pathRoot + `FrontPersistentTimeout`); value.Exists() {
		data.FrontPersistentTimeout = types.Int64Value(value.Int())
	} else {
		data.FrontPersistentTimeout = types.Int64Value(180)
	}
	if value := res.Get(pathRoot + `FrontHTTPVersion`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.FrontHttpVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FrontHttpVersion = types.StringValue("HTTP/1.1")
	}
	if value := res.Get(pathRoot + `HTTPClientIPLabel`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HttpClientIpLabel = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HttpClientIpLabel = types.StringValue("X-Client-IP")
	}
	if value := res.Get(pathRoot + `HTTPLogCorIDLabel`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HttpLogCorIdLabel = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HttpLogCorIdLabel = types.StringValue("X-Global-Transaction-ID")
	}
	if value := res.Get(pathRoot + `DebugMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DebugMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DebugMode = types.StringValue("off")
	}
	if value := res.Get(pathRoot + `DebugHistory`); value.Exists() {
		data.DebugHistory = types.Int64Value(value.Int())
	} else {
		data.DebugHistory = types.Int64Value(25)
	}
}

func (data *WebTokenService) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Priority`); value.Exists() && !data.Priority.IsNull() {
		data.Priority = tfutils.ParseStringFromGJSON(value)
	} else if data.Priority.ValueString() != "normal" {
		data.Priority = types.StringNull()
	}
	if value := res.Get(pathRoot + `XMLManager`); value.Exists() && !data.XmlManager.IsNull() {
		data.XmlManager = tfutils.ParseStringFromGJSON(value)
	} else if data.XmlManager.ValueString() != "default" {
		data.XmlManager = types.StringNull()
	}
	if value := res.Get(pathRoot + `RequestType`); value.Exists() && !data.RequestType.IsNull() {
		data.RequestType = tfutils.ParseStringFromGJSON(value)
	} else if data.RequestType.ValueString() != "preprocessed" {
		data.RequestType = types.StringNull()
	}
	if value := res.Get(pathRoot + `FrontSide`); value.Exists() && !data.FrontSide.IsNull() {
		l := []DmSSLFrontSide{}
		e := []DmSSLFrontSide{}
		data.FrontSide.ElementsAs(ctx, &e, false)
		if len(value.Array()) == len(e) {
			for i, v := range value.Array() {
				item := e[i]
				item.UpdateFromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		} else {
			for _, v := range value.Array() {
				item := DmSSLFrontSide{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.FrontSide, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSSLFrontSideObjectType}, l)
		} else {
			data.FrontSide = types.ListNull(types.ObjectType{AttrTypes: DmSSLFrontSideObjectType})
		}
	} else {
		data.FrontSide = types.ListNull(types.ObjectType{AttrTypes: DmSSLFrontSideObjectType})
	}
	if value := res.Get(pathRoot + `StylePolicy`); value.Exists() && !data.StylePolicy.IsNull() {
		data.StylePolicy = tfutils.ParseStringFromGJSON(value)
	} else if data.StylePolicy.ValueString() != "default" {
		data.StylePolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `RewriteErrors`); value.Exists() && !data.RewriteErrors.IsNull() {
		data.RewriteErrors = tfutils.BoolFromString(value.String())
	} else if !data.RewriteErrors.ValueBool() {
		data.RewriteErrors = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DelayErrors`); value.Exists() && !data.DelayErrors.IsNull() {
		data.DelayErrors = tfutils.BoolFromString(value.String())
	} else if !data.DelayErrors.ValueBool() {
		data.DelayErrors = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DelayErrorsDuration`); value.Exists() && !data.DelayErrorsDuration.IsNull() {
		data.DelayErrorsDuration = types.Int64Value(value.Int())
	} else if data.DelayErrorsDuration.ValueInt64() != 1000 {
		data.DelayErrorsDuration = types.Int64Null()
	}
	if value := res.Get(pathRoot + `FrontTimeout`); value.Exists() && !data.FrontTimeout.IsNull() {
		data.FrontTimeout = types.Int64Value(value.Int())
	} else if data.FrontTimeout.ValueInt64() != 120 {
		data.FrontTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `FrontPersistentTimeout`); value.Exists() && !data.FrontPersistentTimeout.IsNull() {
		data.FrontPersistentTimeout = types.Int64Value(value.Int())
	} else if data.FrontPersistentTimeout.ValueInt64() != 180 {
		data.FrontPersistentTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `FrontHTTPVersion`); value.Exists() && !data.FrontHttpVersion.IsNull() {
		data.FrontHttpVersion = tfutils.ParseStringFromGJSON(value)
	} else if data.FrontHttpVersion.ValueString() != "HTTP/1.1" {
		data.FrontHttpVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `HTTPClientIPLabel`); value.Exists() && !data.HttpClientIpLabel.IsNull() {
		data.HttpClientIpLabel = tfutils.ParseStringFromGJSON(value)
	} else if data.HttpClientIpLabel.ValueString() != "X-Client-IP" {
		data.HttpClientIpLabel = types.StringNull()
	}
	if value := res.Get(pathRoot + `HTTPLogCorIDLabel`); value.Exists() && !data.HttpLogCorIdLabel.IsNull() {
		data.HttpLogCorIdLabel = tfutils.ParseStringFromGJSON(value)
	} else if data.HttpLogCorIdLabel.ValueString() != "X-Global-Transaction-ID" {
		data.HttpLogCorIdLabel = types.StringNull()
	}
	if value := res.Get(pathRoot + `DebugMode`); value.Exists() && !data.DebugMode.IsNull() {
		data.DebugMode = tfutils.ParseStringFromGJSON(value)
	} else if data.DebugMode.ValueString() != "off" {
		data.DebugMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `DebugHistory`); value.Exists() && !data.DebugHistory.IsNull() {
		data.DebugHistory = types.Int64Value(value.Int())
	} else if data.DebugHistory.ValueInt64() != 25 {
		data.DebugHistory = types.Int64Null()
	}
}
