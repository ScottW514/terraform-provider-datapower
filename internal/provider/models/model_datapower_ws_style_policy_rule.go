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

type WSStylePolicyRule struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	Actions           types.List                  `tfsdk:"actions"`
	Direction         types.String                `tfsdk:"direction"`
	InputFormat       types.String                `tfsdk:"input_format"`
	OutputFormat      types.String                `tfsdk:"output_format"`
	NonXmlProcessing  types.Bool                  `tfsdk:"non_xml_processing"`
	Unprocessed       types.Bool                  `tfsdk:"unprocessed"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var WSStylePolicyRuleNonXMLProcessingIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "unprocessed",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var WSStylePolicyRuleUnprocessedIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "non_xml_processing",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var WSStylePolicyRuleObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"actions":            types.ListType{ElemType: types.StringType},
	"direction":          types.StringType,
	"input_format":       types.StringType,
	"output_format":      types.StringType,
	"non_xml_processing": types.BoolType,
	"unprocessed":        types.BoolType,
	"user_summary":       types.StringType,
	"dependency_actions": actions.ActionsListType,
}

func (data WSStylePolicyRule) GetPath() string {
	rest_path := "/mgmt/config/{domain}/WSStylePolicyRule"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data WSStylePolicyRule) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Actions.IsNull() {
		return false
	}
	if !data.Direction.IsNull() {
		return false
	}
	if !data.InputFormat.IsNull() {
		return false
	}
	if !data.OutputFormat.IsNull() {
		return false
	}
	if !data.NonXmlProcessing.IsNull() {
		return false
	}
	if !data.Unprocessed.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	return true
}

func (data WSStylePolicyRule) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.Actions.IsNull() {
		var dataValues []string
		data.Actions.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`Actions`+".-1", map[string]string{"value": val})
		}
	}
	if !data.Direction.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Direction`, data.Direction.ValueString())
	}
	if !data.InputFormat.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`InputFormat`, data.InputFormat.ValueString())
	}
	if !data.OutputFormat.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OutputFormat`, data.OutputFormat.ValueString())
	}
	if !data.NonXmlProcessing.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`NonXMLProcessing`, tfutils.StringFromBool(data.NonXmlProcessing, ""))
	}
	if !data.Unprocessed.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Unprocessed`, tfutils.StringFromBool(data.Unprocessed, ""))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	return body
}

func (data *WSStylePolicyRule) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Actions`); value.Exists() {
		data.Actions = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Actions = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `Direction`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Direction = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Direction = types.StringValue("rule")
	}
	if value := res.Get(pathRoot + `InputFormat`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.InputFormat = tfutils.ParseStringFromGJSON(value)
	} else {
		data.InputFormat = types.StringValue("none")
	}
	if value := res.Get(pathRoot + `OutputFormat`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OutputFormat = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OutputFormat = types.StringValue("none")
	}
	if value := res.Get(pathRoot + `NonXMLProcessing`); value.Exists() {
		data.NonXmlProcessing = tfutils.BoolFromString(value.String())
	} else {
		data.NonXmlProcessing = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Unprocessed`); value.Exists() {
		data.Unprocessed = tfutils.BoolFromString(value.String())
	} else {
		data.Unprocessed = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
}

func (data *WSStylePolicyRule) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Actions`); value.Exists() && !data.Actions.IsNull() {
		data.Actions = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Actions = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `Direction`); value.Exists() && !data.Direction.IsNull() {
		data.Direction = tfutils.ParseStringFromGJSON(value)
	} else if data.Direction.ValueString() != "rule" {
		data.Direction = types.StringNull()
	}
	if value := res.Get(pathRoot + `InputFormat`); value.Exists() && !data.InputFormat.IsNull() {
		data.InputFormat = tfutils.ParseStringFromGJSON(value)
	} else if data.InputFormat.ValueString() != "none" {
		data.InputFormat = types.StringNull()
	}
	if value := res.Get(pathRoot + `OutputFormat`); value.Exists() && !data.OutputFormat.IsNull() {
		data.OutputFormat = tfutils.ParseStringFromGJSON(value)
	} else if data.OutputFormat.ValueString() != "none" {
		data.OutputFormat = types.StringNull()
	}
	if value := res.Get(pathRoot + `NonXMLProcessing`); value.Exists() && !data.NonXmlProcessing.IsNull() {
		data.NonXmlProcessing = tfutils.BoolFromString(value.String())
	} else if data.NonXmlProcessing.ValueBool() {
		data.NonXmlProcessing = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Unprocessed`); value.Exists() && !data.Unprocessed.IsNull() {
		data.Unprocessed = tfutils.BoolFromString(value.String())
	} else if data.Unprocessed.ValueBool() {
		data.Unprocessed = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
}
