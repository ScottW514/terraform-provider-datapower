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

type AssemblyActionParse struct {
	Id                     types.String                     `tfsdk:"id"`
	AppDomain              types.String                     `tfsdk:"app_domain"`
	ParseSettingsReference *DmDynamicParseSettingsReference `tfsdk:"parse_settings_reference"`
	Input                  types.String                     `tfsdk:"input"`
	WarnOnEmptyInput       types.Bool                       `tfsdk:"warn_on_empty_input"`
	UseContentType         types.Bool                       `tfsdk:"use_content_type"`
	Output                 types.String                     `tfsdk:"output"`
	UserSummary            types.String                     `tfsdk:"user_summary"`
	Title                  types.String                     `tfsdk:"title"`
	CorrelationPath        types.String                     `tfsdk:"correlation_path"`
	ActionDebug            types.Bool                       `tfsdk:"action_debug"`
	DependencyActions      []*actions.DependencyAction      `tfsdk:"dependency_actions"`
}

var AssemblyActionParseObjectType = map[string]attr.Type{
	"id":                       types.StringType,
	"app_domain":               types.StringType,
	"parse_settings_reference": types.ObjectType{AttrTypes: DmDynamicParseSettingsReferenceObjectType},
	"input":                    types.StringType,
	"warn_on_empty_input":      types.BoolType,
	"use_content_type":         types.BoolType,
	"output":                   types.StringType,
	"user_summary":             types.StringType,
	"title":                    types.StringType,
	"correlation_path":         types.StringType,
	"action_debug":             types.BoolType,
	"dependency_actions":       actions.ActionsListType,
}

func (data AssemblyActionParse) GetPath() string {
	rest_path := "/mgmt/config/{domain}/AssemblyActionParse"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data AssemblyActionParse) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if data.ParseSettingsReference != nil {
		if !data.ParseSettingsReference.IsNull() {
			return false
		}
	}
	if !data.Input.IsNull() {
		return false
	}
	if !data.WarnOnEmptyInput.IsNull() {
		return false
	}
	if !data.UseContentType.IsNull() {
		return false
	}
	if !data.Output.IsNull() {
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

func (data AssemblyActionParse) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if data.ParseSettingsReference != nil {
		if !data.ParseSettingsReference.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`ParseSettingsReference`, data.ParseSettingsReference.ToBody(ctx, ""))
		}
	}
	if !data.Input.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Input`, data.Input.ValueString())
	}
	if !data.WarnOnEmptyInput.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WarnOnEmptyInput`, tfutils.StringFromBool(data.WarnOnEmptyInput, ""))
	}
	if !data.UseContentType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseContentType`, tfutils.StringFromBool(data.UseContentType, ""))
	}
	if !data.Output.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Output`, data.Output.ValueString())
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

func (data *AssemblyActionParse) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `ParseSettingsReference`); value.Exists() {
		data.ParseSettingsReference = &DmDynamicParseSettingsReference{}
		data.ParseSettingsReference.FromBody(ctx, "", value)
	} else {
		data.ParseSettingsReference = nil
	}
	if value := res.Get(pathRoot + `Input`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Input = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Input = types.StringValue("message")
	}
	if value := res.Get(pathRoot + `WarnOnEmptyInput`); value.Exists() {
		data.WarnOnEmptyInput = tfutils.BoolFromString(value.String())
	} else {
		data.WarnOnEmptyInput = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UseContentType`); value.Exists() {
		data.UseContentType = tfutils.BoolFromString(value.String())
	} else {
		data.UseContentType = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Output`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Output = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Output = types.StringNull()
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

func (data *AssemblyActionParse) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `ParseSettingsReference`); value.Exists() {
		data.ParseSettingsReference.UpdateFromBody(ctx, "", value)
	} else {
		data.ParseSettingsReference = nil
	}
	if value := res.Get(pathRoot + `Input`); value.Exists() && !data.Input.IsNull() {
		data.Input = tfutils.ParseStringFromGJSON(value)
	} else if data.Input.ValueString() != "message" {
		data.Input = types.StringNull()
	}
	if value := res.Get(pathRoot + `WarnOnEmptyInput`); value.Exists() && !data.WarnOnEmptyInput.IsNull() {
		data.WarnOnEmptyInput = tfutils.BoolFromString(value.String())
	} else if data.WarnOnEmptyInput.ValueBool() {
		data.WarnOnEmptyInput = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UseContentType`); value.Exists() && !data.UseContentType.IsNull() {
		data.UseContentType = tfutils.BoolFromString(value.String())
	} else if data.UseContentType.ValueBool() {
		data.UseContentType = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Output`); value.Exists() && !data.Output.IsNull() {
		data.Output = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Output = types.StringNull()
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
