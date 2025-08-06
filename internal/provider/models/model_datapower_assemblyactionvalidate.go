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

type AssemblyActionValidate struct {
	Id              types.String      `tfsdk:"id"`
	AppDomain       types.String      `tfsdk:"app_domain"`
	ValidateAgainst types.String      `tfsdk:"validate_against"`
	ErrorPolicy     types.String      `tfsdk:"error_policy"`
	Schema          types.String      `tfsdk:"schema"`
	Input           types.String      `tfsdk:"input"`
	Output          types.String      `tfsdk:"output"`
	Definition      types.String      `tfsdk:"definition"`
	CompileSettings types.String      `tfsdk:"compile_settings"`
	UserSummary     types.String      `tfsdk:"user_summary"`
	Title           types.String      `tfsdk:"title"`
	CorrelationPath types.String      `tfsdk:"correlation_path"`
	ActionDebug     types.Bool        `tfsdk:"action_debug"`
	ObjectActions   []*actions.Action `tfsdk:"object_actions"`
}

var AssemblyActionValidateObjectType = map[string]attr.Type{
	"id":               types.StringType,
	"app_domain":       types.StringType,
	"validate_against": types.StringType,
	"error_policy":     types.StringType,
	"schema":           types.StringType,
	"input":            types.StringType,
	"output":           types.StringType,
	"definition":       types.StringType,
	"compile_settings": types.StringType,
	"user_summary":     types.StringType,
	"title":            types.StringType,
	"correlation_path": types.StringType,
	"action_debug":     types.BoolType,
	"object_actions":   actions.ActionsListType,
}

func (data AssemblyActionValidate) GetPath() string {
	rest_path := "/mgmt/config/{domain}/AssemblyActionValidate"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data AssemblyActionValidate) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.ValidateAgainst.IsNull() {
		return false
	}
	if !data.ErrorPolicy.IsNull() {
		return false
	}
	if !data.Schema.IsNull() {
		return false
	}
	if !data.Input.IsNull() {
		return false
	}
	if !data.Output.IsNull() {
		return false
	}
	if !data.Definition.IsNull() {
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

func (data AssemblyActionValidate) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.ValidateAgainst.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ValidateAgainst`, data.ValidateAgainst.ValueString())
	}
	if !data.ErrorPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ErrorPolicy`, data.ErrorPolicy.ValueString())
	}
	if !data.Schema.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Schema`, data.Schema.ValueString())
	}
	if !data.Input.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Input`, data.Input.ValueString())
	}
	if !data.Output.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Output`, data.Output.ValueString())
	}
	if !data.Definition.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Definition`, data.Definition.ValueString())
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

func (data *AssemblyActionValidate) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidateAgainst`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ValidateAgainst = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ValidateAgainst = types.StringValue("url")
	}
	if value := res.Get(pathRoot + `ErrorPolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ErrorPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErrorPolicy = types.StringValue("all")
	}
	if value := res.Get(pathRoot + `Schema`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Schema = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Schema = types.StringNull()
	}
	if value := res.Get(pathRoot + `Input`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Input = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Input = types.StringValue("message")
	}
	if value := res.Get(pathRoot + `Output`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Output = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Output = types.StringNull()
	}
	if value := res.Get(pathRoot + `Definition`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Definition = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Definition = types.StringNull()
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

func (data *AssemblyActionValidate) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidateAgainst`); value.Exists() && !data.ValidateAgainst.IsNull() {
		data.ValidateAgainst = tfutils.ParseStringFromGJSON(value)
	} else if data.ValidateAgainst.ValueString() != "url" {
		data.ValidateAgainst = types.StringNull()
	}
	if value := res.Get(pathRoot + `ErrorPolicy`); value.Exists() && !data.ErrorPolicy.IsNull() {
		data.ErrorPolicy = tfutils.ParseStringFromGJSON(value)
	} else if data.ErrorPolicy.ValueString() != "all" {
		data.ErrorPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `Schema`); value.Exists() && !data.Schema.IsNull() {
		data.Schema = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Schema = types.StringNull()
	}
	if value := res.Get(pathRoot + `Input`); value.Exists() && !data.Input.IsNull() {
		data.Input = tfutils.ParseStringFromGJSON(value)
	} else if data.Input.ValueString() != "message" {
		data.Input = types.StringNull()
	}
	if value := res.Get(pathRoot + `Output`); value.Exists() && !data.Output.IsNull() {
		data.Output = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Output = types.StringNull()
	}
	if value := res.Get(pathRoot + `Definition`); value.Exists() && !data.Definition.IsNull() {
		data.Definition = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Definition = types.StringNull()
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
