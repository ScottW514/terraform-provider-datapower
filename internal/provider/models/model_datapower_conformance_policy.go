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

type ConformancePolicy struct {
	Id                           types.String                `tfsdk:"id"`
	AppDomain                    types.String                `tfsdk:"app_domain"`
	UserSummary                  types.String                `tfsdk:"user_summary"`
	Profiles                     *DmConformanceProfiles      `tfsdk:"profiles"`
	IgnoredRequirements          types.List                  `tfsdk:"ignored_requirements"`
	FixupStylesheets             types.List                  `tfsdk:"fixup_stylesheets"`
	AssertBp10Conformance        types.Bool                  `tfsdk:"assert_bp10_conformance"`
	ReportLevel                  types.String                `tfsdk:"report_level"`
	LogTarget                    types.String                `tfsdk:"log_target"`
	RejectLevel                  types.String                `tfsdk:"reject_level"`
	RejectIncludeSummary         types.Bool                  `tfsdk:"reject_include_summary"`
	ResultIsConformanceReport    types.Bool                  `tfsdk:"result_is_conformance_report"`
	ResponsePropertiesEnabled    types.Bool                  `tfsdk:"response_properties_enabled"`
	ResponseReportLevel          types.String                `tfsdk:"response_report_level"`
	ResponseLogTarget            types.String                `tfsdk:"response_log_target"`
	ResponseRejectLevel          types.String                `tfsdk:"response_reject_level"`
	ResponseRejectIncludeSummary types.Bool                  `tfsdk:"response_reject_include_summary"`
	DependencyActions            []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var ConformancePolicyLogTargetCondVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "report_level",
	AttrType:    "String",
	AttrDefault: "never",
	Value:       []string{"never"},
}
var ConformancePolicyRejectIncludeSummaryCondVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "reject_level",
	AttrType:    "String",
	AttrDefault: "never",
	Value:       []string{"never"},
}
var ConformancePolicyResponseReportLevelCondVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "response_properties_enabled",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}
var ConformancePolicyResponseLogTargetCondVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "response_report_level",
	AttrType:    "String",
	AttrDefault: "never",
	Value:       []string{"never"},
}
var ConformancePolicyResponseRejectLevelCondVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "response_properties_enabled",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}
var ConformancePolicyResponseRejectIncludeSummaryCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "response_reject_level",
			AttrType:    "String",
			AttrDefault: "never",
			Value:       []string{"never"},
		},
	},
}
var ConformancePolicyAssertBP10ConformanceIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "profiles",
	AttrType:    "DmConformanceProfiles",
	AttrDefault: "",
	Value:       []string{"bp10"},
}
var ConformancePolicyLogTargetIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var ConformancePolicyRejectIncludeSummaryIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var ConformancePolicyResponseReportLevelIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "response_properties_enabled",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}
var ConformancePolicyResponseLogTargetIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "response_properties_enabled",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}
var ConformancePolicyResponseRejectLevelIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "response_properties_enabled",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}
var ConformancePolicyResponseRejectIncludeSummaryIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "response_properties_enabled",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var ConformancePolicyObjectType = map[string]attr.Type{
	"id":                              types.StringType,
	"app_domain":                      types.StringType,
	"user_summary":                    types.StringType,
	"profiles":                        types.ObjectType{AttrTypes: DmConformanceProfilesObjectType},
	"ignored_requirements":            types.ListType{ElemType: types.StringType},
	"fixup_stylesheets":               types.ListType{ElemType: types.StringType},
	"assert_bp10_conformance":         types.BoolType,
	"report_level":                    types.StringType,
	"log_target":                      types.StringType,
	"reject_level":                    types.StringType,
	"reject_include_summary":          types.BoolType,
	"result_is_conformance_report":    types.BoolType,
	"response_properties_enabled":     types.BoolType,
	"response_report_level":           types.StringType,
	"response_log_target":             types.StringType,
	"response_reject_level":           types.StringType,
	"response_reject_include_summary": types.BoolType,
	"dependency_actions":              actions.ActionsListType,
}

func (data ConformancePolicy) GetPath() string {
	rest_path := "/mgmt/config/{domain}/ConformancePolicy"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data ConformancePolicy) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if data.Profiles != nil {
		if !data.Profiles.IsNull() {
			return false
		}
	}
	if !data.IgnoredRequirements.IsNull() {
		return false
	}
	if !data.FixupStylesheets.IsNull() {
		return false
	}
	if !data.AssertBp10Conformance.IsNull() {
		return false
	}
	if !data.ReportLevel.IsNull() {
		return false
	}
	if !data.LogTarget.IsNull() {
		return false
	}
	if !data.RejectLevel.IsNull() {
		return false
	}
	if !data.RejectIncludeSummary.IsNull() {
		return false
	}
	if !data.ResultIsConformanceReport.IsNull() {
		return false
	}
	if !data.ResponsePropertiesEnabled.IsNull() {
		return false
	}
	if !data.ResponseReportLevel.IsNull() {
		return false
	}
	if !data.ResponseLogTarget.IsNull() {
		return false
	}
	if !data.ResponseRejectLevel.IsNull() {
		return false
	}
	if !data.ResponseRejectIncludeSummary.IsNull() {
		return false
	}
	return true
}

func (data ConformancePolicy) ToBody(ctx context.Context, pathRoot string) string {
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
	if data.Profiles != nil {
		if !data.Profiles.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`Profiles`, data.Profiles.ToBody(ctx, ""))
		}
	}
	if !data.IgnoredRequirements.IsNull() {
		var dataValues []string
		data.IgnoredRequirements.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`IgnoredRequirements`+".-1", map[string]string{"value": val})
		}
	}
	if !data.FixupStylesheets.IsNull() {
		var dataValues []string
		data.FixupStylesheets.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`FixupStylesheets`+".-1", map[string]string{"value": val})
		}
	}
	if !data.AssertBp10Conformance.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AssertBP10Conformance`, tfutils.StringFromBool(data.AssertBp10Conformance, ""))
	}
	if !data.ReportLevel.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ReportLevel`, data.ReportLevel.ValueString())
	}
	if !data.LogTarget.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LogTarget`, data.LogTarget.ValueString())
	}
	if !data.RejectLevel.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RejectLevel`, data.RejectLevel.ValueString())
	}
	if !data.RejectIncludeSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RejectIncludeSummary`, tfutils.StringFromBool(data.RejectIncludeSummary, ""))
	}
	if !data.ResultIsConformanceReport.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ResultIsConformanceReport`, tfutils.StringFromBool(data.ResultIsConformanceReport, ""))
	}
	if !data.ResponsePropertiesEnabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ResponsePropertiesEnabled`, tfutils.StringFromBool(data.ResponsePropertiesEnabled, ""))
	}
	if !data.ResponseReportLevel.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ResponseReportLevel`, data.ResponseReportLevel.ValueString())
	}
	if !data.ResponseLogTarget.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ResponseLogTarget`, data.ResponseLogTarget.ValueString())
	}
	if !data.ResponseRejectLevel.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ResponseRejectLevel`, data.ResponseRejectLevel.ValueString())
	}
	if !data.ResponseRejectIncludeSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ResponseRejectIncludeSummary`, tfutils.StringFromBool(data.ResponseRejectIncludeSummary, ""))
	}
	return body
}

func (data *ConformancePolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Profiles`); value.Exists() {
		data.Profiles = &DmConformanceProfiles{}
		data.Profiles.FromBody(ctx, "", value)
	} else {
		data.Profiles = nil
	}
	if value := res.Get(pathRoot + `IgnoredRequirements`); value.Exists() {
		data.IgnoredRequirements = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.IgnoredRequirements = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `FixupStylesheets`); value.Exists() {
		data.FixupStylesheets = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.FixupStylesheets = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `AssertBP10Conformance`); value.Exists() {
		data.AssertBp10Conformance = tfutils.BoolFromString(value.String())
	} else {
		data.AssertBp10Conformance = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ReportLevel`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ReportLevel = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ReportLevel = types.StringValue("never")
	}
	if value := res.Get(pathRoot + `LogTarget`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LogTarget = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LogTarget = types.StringNull()
	}
	if value := res.Get(pathRoot + `RejectLevel`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RejectLevel = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RejectLevel = types.StringValue("never")
	}
	if value := res.Get(pathRoot + `RejectIncludeSummary`); value.Exists() {
		data.RejectIncludeSummary = tfutils.BoolFromString(value.String())
	} else {
		data.RejectIncludeSummary = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ResultIsConformanceReport`); value.Exists() {
		data.ResultIsConformanceReport = tfutils.BoolFromString(value.String())
	} else {
		data.ResultIsConformanceReport = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ResponsePropertiesEnabled`); value.Exists() {
		data.ResponsePropertiesEnabled = tfutils.BoolFromString(value.String())
	} else {
		data.ResponsePropertiesEnabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ResponseReportLevel`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ResponseReportLevel = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ResponseReportLevel = types.StringValue("never")
	}
	if value := res.Get(pathRoot + `ResponseLogTarget`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ResponseLogTarget = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ResponseLogTarget = types.StringNull()
	}
	if value := res.Get(pathRoot + `ResponseRejectLevel`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ResponseRejectLevel = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ResponseRejectLevel = types.StringValue("never")
	}
	if value := res.Get(pathRoot + `ResponseRejectIncludeSummary`); value.Exists() {
		data.ResponseRejectIncludeSummary = tfutils.BoolFromString(value.String())
	} else {
		data.ResponseRejectIncludeSummary = types.BoolNull()
	}
}

func (data *ConformancePolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Profiles`); value.Exists() {
		data.Profiles.UpdateFromBody(ctx, "", value)
	} else {
		data.Profiles = nil
	}
	if value := res.Get(pathRoot + `IgnoredRequirements`); value.Exists() && !data.IgnoredRequirements.IsNull() {
		data.IgnoredRequirements = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.IgnoredRequirements = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `FixupStylesheets`); value.Exists() && !data.FixupStylesheets.IsNull() {
		data.FixupStylesheets = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.FixupStylesheets = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `AssertBP10Conformance`); value.Exists() && !data.AssertBp10Conformance.IsNull() {
		data.AssertBp10Conformance = tfutils.BoolFromString(value.String())
	} else if data.AssertBp10Conformance.ValueBool() {
		data.AssertBp10Conformance = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ReportLevel`); value.Exists() && !data.ReportLevel.IsNull() {
		data.ReportLevel = tfutils.ParseStringFromGJSON(value)
	} else if data.ReportLevel.ValueString() != "never" {
		data.ReportLevel = types.StringNull()
	}
	if value := res.Get(pathRoot + `LogTarget`); value.Exists() && !data.LogTarget.IsNull() {
		data.LogTarget = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LogTarget = types.StringNull()
	}
	if value := res.Get(pathRoot + `RejectLevel`); value.Exists() && !data.RejectLevel.IsNull() {
		data.RejectLevel = tfutils.ParseStringFromGJSON(value)
	} else if data.RejectLevel.ValueString() != "never" {
		data.RejectLevel = types.StringNull()
	}
	if value := res.Get(pathRoot + `RejectIncludeSummary`); value.Exists() && !data.RejectIncludeSummary.IsNull() {
		data.RejectIncludeSummary = tfutils.BoolFromString(value.String())
	} else if data.RejectIncludeSummary.ValueBool() {
		data.RejectIncludeSummary = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ResultIsConformanceReport`); value.Exists() && !data.ResultIsConformanceReport.IsNull() {
		data.ResultIsConformanceReport = tfutils.BoolFromString(value.String())
	} else if data.ResultIsConformanceReport.ValueBool() {
		data.ResultIsConformanceReport = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ResponsePropertiesEnabled`); value.Exists() && !data.ResponsePropertiesEnabled.IsNull() {
		data.ResponsePropertiesEnabled = tfutils.BoolFromString(value.String())
	} else if data.ResponsePropertiesEnabled.ValueBool() {
		data.ResponsePropertiesEnabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ResponseReportLevel`); value.Exists() && !data.ResponseReportLevel.IsNull() {
		data.ResponseReportLevel = tfutils.ParseStringFromGJSON(value)
	} else if data.ResponseReportLevel.ValueString() != "never" {
		data.ResponseReportLevel = types.StringNull()
	}
	if value := res.Get(pathRoot + `ResponseLogTarget`); value.Exists() && !data.ResponseLogTarget.IsNull() {
		data.ResponseLogTarget = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ResponseLogTarget = types.StringNull()
	}
	if value := res.Get(pathRoot + `ResponseRejectLevel`); value.Exists() && !data.ResponseRejectLevel.IsNull() {
		data.ResponseRejectLevel = tfutils.ParseStringFromGJSON(value)
	} else if data.ResponseRejectLevel.ValueString() != "never" {
		data.ResponseRejectLevel = types.StringNull()
	}
	if value := res.Get(pathRoot + `ResponseRejectIncludeSummary`); value.Exists() && !data.ResponseRejectIncludeSummary.IsNull() {
		data.ResponseRejectIncludeSummary = tfutils.BoolFromString(value.String())
	} else if data.ResponseRejectIncludeSummary.ValueBool() {
		data.ResponseRejectIncludeSummary = types.BoolNull()
	}
}
