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

type NFSFilePollerSourceProtocolHandler struct {
	Id                      types.String                `tfsdk:"id"`
	AppDomain               types.String                `tfsdk:"app_domain"`
	UserSummary             types.String                `tfsdk:"user_summary"`
	TargetDirectory         types.String                `tfsdk:"target_directory"`
	DelayBetweenPolls       types.Int64                 `tfsdk:"delay_between_polls"`
	InputFileMatchPattern   types.String                `tfsdk:"input_file_match_pattern"`
	ProcessingRenamePattern types.String                `tfsdk:"processing_rename_pattern"`
	DeleteOnSuccess         types.Bool                  `tfsdk:"delete_on_success"`
	SuccessRenamePattern    types.String                `tfsdk:"success_rename_pattern"`
	DeleteOnError           types.Bool                  `tfsdk:"delete_on_error"`
	ErrorRenamePattern      types.String                `tfsdk:"error_rename_pattern"`
	GenerateResultFile      types.Bool                  `tfsdk:"generate_result_file"`
	ResultNamePattern       types.String                `tfsdk:"result_name_pattern"`
	ProcessingSeizeTimeout  types.Int64                 `tfsdk:"processing_seize_timeout"`
	ProcessingSeizePattern  types.String                `tfsdk:"processing_seize_pattern"`
	XmlManager              types.String                `tfsdk:"xml_manager"`
	MaxTransfersPerPoll     types.Int64                 `tfsdk:"max_transfers_per_poll"`
	DependencyActions       []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var NFSFilePollerSourceProtocolHandlerSuccessRenamePatternIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "delete_on_success",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var NFSFilePollerSourceProtocolHandlerErrorRenamePatternIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "delete_on_error",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var NFSFilePollerSourceProtocolHandlerResultNamePatternCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "generate_result_file",
	AttrType:    "Bool",
	AttrDefault: "true",
	Value:       []string{"true"},
}

var NFSFilePollerSourceProtocolHandlerResultNamePatternIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var NFSFilePollerSourceProtocolHandlerProcessingSeizePatternCondVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "processing_seize_timeout",
	AttrType:    "Int64",
	AttrDefault: "0",
	Value:       []string{"0"},
}

var NFSFilePollerSourceProtocolHandlerProcessingSeizePatternIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "processing_seize_timeout",
	AttrType:    "Int64",
	AttrDefault: "0",
	Value:       []string{"0"},
}

var NFSFilePollerSourceProtocolHandlerObjectType = map[string]attr.Type{
	"id":                        types.StringType,
	"app_domain":                types.StringType,
	"user_summary":              types.StringType,
	"target_directory":          types.StringType,
	"delay_between_polls":       types.Int64Type,
	"input_file_match_pattern":  types.StringType,
	"processing_rename_pattern": types.StringType,
	"delete_on_success":         types.BoolType,
	"success_rename_pattern":    types.StringType,
	"delete_on_error":           types.BoolType,
	"error_rename_pattern":      types.StringType,
	"generate_result_file":      types.BoolType,
	"result_name_pattern":       types.StringType,
	"processing_seize_timeout":  types.Int64Type,
	"processing_seize_pattern":  types.StringType,
	"xml_manager":               types.StringType,
	"max_transfers_per_poll":    types.Int64Type,
	"dependency_actions":        actions.ActionsListType,
}

func (data NFSFilePollerSourceProtocolHandler) GetPath() string {
	rest_path := "/mgmt/config/{domain}/NFSFilePollerSourceProtocolHandler"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data NFSFilePollerSourceProtocolHandler) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.TargetDirectory.IsNull() {
		return false
	}
	if !data.DelayBetweenPolls.IsNull() {
		return false
	}
	if !data.InputFileMatchPattern.IsNull() {
		return false
	}
	if !data.ProcessingRenamePattern.IsNull() {
		return false
	}
	if !data.DeleteOnSuccess.IsNull() {
		return false
	}
	if !data.SuccessRenamePattern.IsNull() {
		return false
	}
	if !data.DeleteOnError.IsNull() {
		return false
	}
	if !data.ErrorRenamePattern.IsNull() {
		return false
	}
	if !data.GenerateResultFile.IsNull() {
		return false
	}
	if !data.ResultNamePattern.IsNull() {
		return false
	}
	if !data.ProcessingSeizeTimeout.IsNull() {
		return false
	}
	if !data.ProcessingSeizePattern.IsNull() {
		return false
	}
	if !data.XmlManager.IsNull() {
		return false
	}
	if !data.MaxTransfersPerPoll.IsNull() {
		return false
	}
	return true
}

func (data NFSFilePollerSourceProtocolHandler) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.TargetDirectory.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TargetDirectory`, data.TargetDirectory.ValueString())
	}
	if !data.DelayBetweenPolls.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DelayBetweenPolls`, data.DelayBetweenPolls.ValueInt64())
	}
	if !data.InputFileMatchPattern.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`InputFileMatchPattern`, data.InputFileMatchPattern.ValueString())
	}
	if !data.ProcessingRenamePattern.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ProcessingRenamePattern`, data.ProcessingRenamePattern.ValueString())
	}
	if !data.DeleteOnSuccess.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DeleteOnSuccess`, tfutils.StringFromBool(data.DeleteOnSuccess, ""))
	}
	if !data.SuccessRenamePattern.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SuccessRenamePattern`, data.SuccessRenamePattern.ValueString())
	}
	if !data.DeleteOnError.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DeleteOnError`, tfutils.StringFromBool(data.DeleteOnError, ""))
	}
	if !data.ErrorRenamePattern.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ErrorRenamePattern`, data.ErrorRenamePattern.ValueString())
	}
	if !data.GenerateResultFile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GenerateResultFile`, tfutils.StringFromBool(data.GenerateResultFile, ""))
	}
	if !data.ResultNamePattern.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ResultNamePattern`, data.ResultNamePattern.ValueString())
	}
	if !data.ProcessingSeizeTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ProcessingSeizeTimeout`, data.ProcessingSeizeTimeout.ValueInt64())
	}
	if !data.ProcessingSeizePattern.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ProcessingSeizePattern`, data.ProcessingSeizePattern.ValueString())
	}
	if !data.XmlManager.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XMLManager`, data.XmlManager.ValueString())
	}
	if !data.MaxTransfersPerPoll.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxTransfersPerPoll`, data.MaxTransfersPerPoll.ValueInt64())
	}
	return body
}

func (data *NFSFilePollerSourceProtocolHandler) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `TargetDirectory`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TargetDirectory = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TargetDirectory = types.StringNull()
	}
	if value := res.Get(pathRoot + `DelayBetweenPolls`); value.Exists() {
		data.DelayBetweenPolls = types.Int64Value(value.Int())
	} else {
		data.DelayBetweenPolls = types.Int64Value(60000)
	}
	if value := res.Get(pathRoot + `InputFileMatchPattern`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.InputFileMatchPattern = tfutils.ParseStringFromGJSON(value)
	} else {
		data.InputFileMatchPattern = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProcessingRenamePattern`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ProcessingRenamePattern = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProcessingRenamePattern = types.StringNull()
	}
	if value := res.Get(pathRoot + `DeleteOnSuccess`); value.Exists() {
		data.DeleteOnSuccess = tfutils.BoolFromString(value.String())
	} else {
		data.DeleteOnSuccess = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SuccessRenamePattern`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SuccessRenamePattern = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SuccessRenamePattern = types.StringValue("$1.processed.ok")
	}
	if value := res.Get(pathRoot + `DeleteOnError`); value.Exists() {
		data.DeleteOnError = tfutils.BoolFromString(value.String())
	} else {
		data.DeleteOnError = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ErrorRenamePattern`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ErrorRenamePattern = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErrorRenamePattern = types.StringValue("$0.processed.error")
	}
	if value := res.Get(pathRoot + `GenerateResultFile`); value.Exists() {
		data.GenerateResultFile = tfutils.BoolFromString(value.String())
	} else {
		data.GenerateResultFile = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ResultNamePattern`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ResultNamePattern = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ResultNamePattern = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProcessingSeizeTimeout`); value.Exists() {
		data.ProcessingSeizeTimeout = types.Int64Value(value.Int())
	} else {
		data.ProcessingSeizeTimeout = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `ProcessingSeizePattern`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ProcessingSeizePattern = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProcessingSeizePattern = types.StringNull()
	}
	if value := res.Get(pathRoot + `XMLManager`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.XmlManager = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XmlManager = types.StringValue("default")
	}
	if value := res.Get(pathRoot + `MaxTransfersPerPoll`); value.Exists() {
		data.MaxTransfersPerPoll = types.Int64Value(value.Int())
	} else {
		data.MaxTransfersPerPoll = types.Int64Value(0)
	}
}

func (data *NFSFilePollerSourceProtocolHandler) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `TargetDirectory`); value.Exists() && !data.TargetDirectory.IsNull() {
		data.TargetDirectory = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TargetDirectory = types.StringNull()
	}
	if value := res.Get(pathRoot + `DelayBetweenPolls`); value.Exists() && !data.DelayBetweenPolls.IsNull() {
		data.DelayBetweenPolls = types.Int64Value(value.Int())
	} else if data.DelayBetweenPolls.ValueInt64() != 60000 {
		data.DelayBetweenPolls = types.Int64Null()
	}
	if value := res.Get(pathRoot + `InputFileMatchPattern`); value.Exists() && !data.InputFileMatchPattern.IsNull() {
		data.InputFileMatchPattern = tfutils.ParseStringFromGJSON(value)
	} else {
		data.InputFileMatchPattern = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProcessingRenamePattern`); value.Exists() && !data.ProcessingRenamePattern.IsNull() {
		data.ProcessingRenamePattern = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProcessingRenamePattern = types.StringNull()
	}
	if value := res.Get(pathRoot + `DeleteOnSuccess`); value.Exists() && !data.DeleteOnSuccess.IsNull() {
		data.DeleteOnSuccess = tfutils.BoolFromString(value.String())
	} else if data.DeleteOnSuccess.ValueBool() {
		data.DeleteOnSuccess = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SuccessRenamePattern`); value.Exists() && !data.SuccessRenamePattern.IsNull() {
		data.SuccessRenamePattern = tfutils.ParseStringFromGJSON(value)
	} else if data.SuccessRenamePattern.ValueString() != "$1.processed.ok" {
		data.SuccessRenamePattern = types.StringNull()
	}
	if value := res.Get(pathRoot + `DeleteOnError`); value.Exists() && !data.DeleteOnError.IsNull() {
		data.DeleteOnError = tfutils.BoolFromString(value.String())
	} else if data.DeleteOnError.ValueBool() {
		data.DeleteOnError = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ErrorRenamePattern`); value.Exists() && !data.ErrorRenamePattern.IsNull() {
		data.ErrorRenamePattern = tfutils.ParseStringFromGJSON(value)
	} else if data.ErrorRenamePattern.ValueString() != "$0.processed.error" {
		data.ErrorRenamePattern = types.StringNull()
	}
	if value := res.Get(pathRoot + `GenerateResultFile`); value.Exists() && !data.GenerateResultFile.IsNull() {
		data.GenerateResultFile = tfutils.BoolFromString(value.String())
	} else if !data.GenerateResultFile.ValueBool() {
		data.GenerateResultFile = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ResultNamePattern`); value.Exists() && !data.ResultNamePattern.IsNull() {
		data.ResultNamePattern = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ResultNamePattern = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProcessingSeizeTimeout`); value.Exists() && !data.ProcessingSeizeTimeout.IsNull() {
		data.ProcessingSeizeTimeout = types.Int64Value(value.Int())
	} else if data.ProcessingSeizeTimeout.ValueInt64() != 0 {
		data.ProcessingSeizeTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ProcessingSeizePattern`); value.Exists() && !data.ProcessingSeizePattern.IsNull() {
		data.ProcessingSeizePattern = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProcessingSeizePattern = types.StringNull()
	}
	if value := res.Get(pathRoot + `XMLManager`); value.Exists() && !data.XmlManager.IsNull() {
		data.XmlManager = tfutils.ParseStringFromGJSON(value)
	} else if data.XmlManager.ValueString() != "default" {
		data.XmlManager = types.StringNull()
	}
	if value := res.Get(pathRoot + `MaxTransfersPerPoll`); value.Exists() && !data.MaxTransfersPerPoll.IsNull() {
		data.MaxTransfersPerPoll = types.Int64Value(value.Int())
	} else if data.MaxTransfersPerPoll.ValueInt64() != 0 {
		data.MaxTransfersPerPoll = types.Int64Null()
	}
}
