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
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type AssemblyActionThrow struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	ErrorId           types.String                `tfsdk:"error_id"`
	ErrorText         types.String                `tfsdk:"error_text"`
	ErrorStatusCode   types.String                `tfsdk:"error_status_code"`
	ErrorStatusReason types.String                `tfsdk:"error_status_reason"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	Title             types.String                `tfsdk:"title"`
	CorrelationPath   types.String                `tfsdk:"correlation_path"`
	ActionDebug       types.Bool                  `tfsdk:"action_debug"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var AssemblyActionThrowObjectType = map[string]attr.Type{
	"id":                  types.StringType,
	"app_domain":          types.StringType,
	"error_id":            types.StringType,
	"error_text":          types.StringType,
	"error_status_code":   types.StringType,
	"error_status_reason": types.StringType,
	"user_summary":        types.StringType,
	"title":               types.StringType,
	"correlation_path":    types.StringType,
	"action_debug":        types.BoolType,
	"dependency_actions":  actions.ActionsListType,
}

func (data AssemblyActionThrow) GetPath() string {
	rest_path := "/mgmt/config/{domain}/AssemblyActionThrow"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data AssemblyActionThrow) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.ErrorId.IsNull() {
		return false
	}
	if !data.ErrorText.IsNull() {
		return false
	}
	if !data.ErrorStatusCode.IsNull() {
		return false
	}
	if !data.ErrorStatusReason.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Title.IsNull() {
		return false
	}
	if !data.CorrelationPath.IsNull() {
		return false
	}
	if !data.ActionDebug.IsNull() {
		return false
	}
	return true
}

func (data AssemblyActionThrow) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.ErrorId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ErrorID`, data.ErrorId.ValueString())
	}
	if !data.ErrorText.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ErrorText`, data.ErrorText.ValueString())
	}
	if !data.ErrorStatusCode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ErrorStatusCode`, data.ErrorStatusCode.ValueString())
	}
	if !data.ErrorStatusReason.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ErrorStatusReason`, data.ErrorStatusReason.ValueString())
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.Title.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Title`, data.Title.ValueString())
	}
	if !data.CorrelationPath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CorrelationPath`, data.CorrelationPath.ValueString())
	}
	if !data.ActionDebug.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ActionDebug`, tfutils.StringFromBool(data.ActionDebug, ""))
	}
	return body
}

func (data *AssemblyActionThrow) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `ErrorID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ErrorId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErrorId = types.StringNull()
	}
	if value := res.Get(pathRoot + `ErrorText`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ErrorText = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErrorText = types.StringNull()
	}
	if value := res.Get(pathRoot + `ErrorStatusCode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ErrorStatusCode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErrorStatusCode = types.StringNull()
	}
	if value := res.Get(pathRoot + `ErrorStatusReason`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ErrorStatusReason = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErrorStatusReason = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `Title`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Title = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Title = types.StringNull()
	}
	if value := res.Get(pathRoot + `CorrelationPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CorrelationPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CorrelationPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `ActionDebug`); value.Exists() {
		data.ActionDebug = tfutils.BoolFromString(value.String())
	} else {
		data.ActionDebug = types.BoolNull()
	}
}

func (data *AssemblyActionThrow) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `ErrorID`); value.Exists() && !data.ErrorId.IsNull() {
		data.ErrorId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErrorId = types.StringNull()
	}
	if value := res.Get(pathRoot + `ErrorText`); value.Exists() && !data.ErrorText.IsNull() {
		data.ErrorText = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErrorText = types.StringNull()
	}
	if value := res.Get(pathRoot + `ErrorStatusCode`); value.Exists() && !data.ErrorStatusCode.IsNull() {
		data.ErrorStatusCode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErrorStatusCode = types.StringNull()
	}
	if value := res.Get(pathRoot + `ErrorStatusReason`); value.Exists() && !data.ErrorStatusReason.IsNull() {
		data.ErrorStatusReason = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErrorStatusReason = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `Title`); value.Exists() && !data.Title.IsNull() {
		data.Title = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Title = types.StringNull()
	}
	if value := res.Get(pathRoot + `CorrelationPath`); value.Exists() && !data.CorrelationPath.IsNull() {
		data.CorrelationPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CorrelationPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `ActionDebug`); value.Exists() && !data.ActionDebug.IsNull() {
		data.ActionDebug = tfutils.BoolFromString(value.String())
	} else if data.ActionDebug.ValueBool() {
		data.ActionDebug = types.BoolNull()
	}
}
