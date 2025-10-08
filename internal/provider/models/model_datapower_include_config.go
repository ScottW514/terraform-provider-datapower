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

type IncludeConfig struct {
	Id                 types.String                `tfsdk:"id"`
	AppDomain          types.String                `tfsdk:"app_domain"`
	UserSummary        types.String                `tfsdk:"user_summary"`
	Url                types.String                `tfsdk:"url"`
	OnStartup          types.Bool                  `tfsdk:"on_startup"`
	InterfaceDetection types.Bool                  `tfsdk:"interface_detection"`
	DependencyActions  []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var IncludeConfigInterfaceDetectionCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "on_startup",
	AttrType:    "Bool",
	AttrDefault: "true",
	Value:       []string{"true"},
}

var IncludeConfigInterfaceDetectionIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var IncludeConfigObjectType = map[string]attr.Type{
	"id":                  types.StringType,
	"app_domain":          types.StringType,
	"user_summary":        types.StringType,
	"url":                 types.StringType,
	"on_startup":          types.BoolType,
	"interface_detection": types.BoolType,
	"dependency_actions":  actions.ActionsListType,
}

func (data IncludeConfig) GetPath() string {
	rest_path := "/mgmt/config/{domain}/IncludeConfig"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data IncludeConfig) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Url.IsNull() {
		return false
	}
	if !data.OnStartup.IsNull() {
		return false
	}
	if !data.InterfaceDetection.IsNull() {
		return false
	}
	return true
}

func (data IncludeConfig) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Url.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`URL`, data.Url.ValueString())
	}
	if !data.OnStartup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OnStartup`, tfutils.StringFromBool(data.OnStartup, ""))
	}
	if !data.InterfaceDetection.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`InterfaceDetection`, tfutils.StringFromBool(data.InterfaceDetection, ""))
	}
	return body
}

func (data *IncludeConfig) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `URL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Url = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Url = types.StringNull()
	}
	if value := res.Get(pathRoot + `OnStartup`); value.Exists() {
		data.OnStartup = tfutils.BoolFromString(value.String())
	} else {
		data.OnStartup = types.BoolNull()
	}
	if value := res.Get(pathRoot + `InterfaceDetection`); value.Exists() {
		data.InterfaceDetection = tfutils.BoolFromString(value.String())
	} else {
		data.InterfaceDetection = types.BoolNull()
	}
}

func (data *IncludeConfig) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `URL`); value.Exists() && !data.Url.IsNull() {
		data.Url = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Url = types.StringNull()
	}
	if value := res.Get(pathRoot + `OnStartup`); value.Exists() && !data.OnStartup.IsNull() {
		data.OnStartup = tfutils.BoolFromString(value.String())
	} else if !data.OnStartup.ValueBool() {
		data.OnStartup = types.BoolNull()
	}
	if value := res.Get(pathRoot + `InterfaceDetection`); value.Exists() && !data.InterfaceDetection.IsNull() {
		data.InterfaceDetection = tfutils.BoolFromString(value.String())
	} else if data.InterfaceDetection.ValueBool() {
		data.InterfaceDetection = types.BoolNull()
	}
}
