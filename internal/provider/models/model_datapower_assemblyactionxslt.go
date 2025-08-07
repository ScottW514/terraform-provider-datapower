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

type AssemblyActionXSLT struct {
	Id                types.String      `tfsdk:"id"`
	AppDomain         types.String      `tfsdk:"app_domain"`
	UsePayload        types.Bool        `tfsdk:"use_payload"`
	Stylesheet        types.String      `tfsdk:"stylesheet"`
	SerializeOutput   types.Bool        `tfsdk:"serialize_output"`
	CompileSettings   types.String      `tfsdk:"compile_settings"`
	UserSummary       types.String      `tfsdk:"user_summary"`
	Title             types.String      `tfsdk:"title"`
	CorrelationPath   types.String      `tfsdk:"correlation_path"`
	ActionDebug       types.Bool        `tfsdk:"action_debug"`
	DependencyActions []*actions.Action `tfsdk:"dependency_actions"`
}

var AssemblyActionXSLTObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"use_payload":        types.BoolType,
	"stylesheet":         types.StringType,
	"serialize_output":   types.BoolType,
	"compile_settings":   types.StringType,
	"user_summary":       types.StringType,
	"title":              types.StringType,
	"correlation_path":   types.StringType,
	"action_debug":       types.BoolType,
	"dependency_actions": actions.ActionsListType,
}

func (data AssemblyActionXSLT) GetPath() string {
	rest_path := "/mgmt/config/{domain}/AssemblyActionXSLT"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data AssemblyActionXSLT) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UsePayload.IsNull() {
		return false
	}
	if !data.Stylesheet.IsNull() {
		return false
	}
	if !data.SerializeOutput.IsNull() {
		return false
	}
	if !data.CompileSettings.IsNull() {
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

func (data AssemblyActionXSLT) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.UsePayload.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UsePayload`, tfutils.StringFromBool(data.UsePayload, ""))
	}
	if !data.Stylesheet.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Stylesheet`, data.Stylesheet.ValueString())
	}
	if !data.SerializeOutput.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SerializeOutput`, tfutils.StringFromBool(data.SerializeOutput, ""))
	}
	if !data.CompileSettings.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CompileSettings`, data.CompileSettings.ValueString())
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

func (data *AssemblyActionXSLT) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `UsePayload`); value.Exists() {
		data.UsePayload = tfutils.BoolFromString(value.String())
	} else {
		data.UsePayload = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Stylesheet`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Stylesheet = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Stylesheet = types.StringNull()
	}
	if value := res.Get(pathRoot + `SerializeOutput`); value.Exists() {
		data.SerializeOutput = tfutils.BoolFromString(value.String())
	} else {
		data.SerializeOutput = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CompileSettings`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CompileSettings = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CompileSettings = types.StringNull()
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

func (data *AssemblyActionXSLT) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `UsePayload`); value.Exists() && !data.UsePayload.IsNull() {
		data.UsePayload = tfutils.BoolFromString(value.String())
	} else if data.UsePayload.ValueBool() {
		data.UsePayload = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Stylesheet`); value.Exists() && !data.Stylesheet.IsNull() {
		data.Stylesheet = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Stylesheet = types.StringNull()
	}
	if value := res.Get(pathRoot + `SerializeOutput`); value.Exists() && !data.SerializeOutput.IsNull() {
		data.SerializeOutput = tfutils.BoolFromString(value.String())
	} else if data.SerializeOutput.ValueBool() {
		data.SerializeOutput = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CompileSettings`); value.Exists() && !data.CompileSettings.IsNull() {
		data.CompileSettings = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CompileSettings = types.StringNull()
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
