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
	"path"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type Throttler struct {
	Enabled           types.Bool                  `tfsdk:"enabled"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	ThrottleAt        types.Int64                 `tfsdk:"throttle_at"`
	TerminateAt       types.Int64                 `tfsdk:"terminate_at"`
	TempFsThrottleAt  types.Int64                 `tfsdk:"temp_fs_throttle_at"`
	TempFsTerminateAt types.Int64                 `tfsdk:"temp_fs_terminate_at"`
	QnameWarnAt       types.Int64                 `tfsdk:"qname_warn_at"`
	Timeout           types.Int64                 `tfsdk:"timeout"`
	Statistics        types.Bool                  `tfsdk:"statistics"`
	LogLevel          types.String                `tfsdk:"log_level"`
	EnvironmentalLog  types.Bool                  `tfsdk:"environmental_log"`
	BacklogSize       types.Int64                 `tfsdk:"backlog_size"`
	BacklogTimeout    types.Int64                 `tfsdk:"backlog_timeout"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var ThrottlerObjectType = map[string]attr.Type{
	"enabled":              types.BoolType,
	"user_summary":         types.StringType,
	"throttle_at":          types.Int64Type,
	"terminate_at":         types.Int64Type,
	"temp_fs_throttle_at":  types.Int64Type,
	"temp_fs_terminate_at": types.Int64Type,
	"qname_warn_at":        types.Int64Type,
	"timeout":              types.Int64Type,
	"statistics":           types.BoolType,
	"log_level":            types.StringType,
	"environmental_log":    types.BoolType,
	"backlog_size":         types.Int64Type,
	"backlog_timeout":      types.Int64Type,
	"dependency_actions":   actions.ActionsListType,
}

func (data Throttler) GetPath() string {
	rest_path := "/mgmt/config/default/Throttler/Throttler"
	return rest_path
}

func (data Throttler) IsNull() bool {
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.ThrottleAt.IsNull() {
		return false
	}
	if !data.TerminateAt.IsNull() {
		return false
	}
	if !data.TempFsThrottleAt.IsNull() {
		return false
	}
	if !data.TempFsTerminateAt.IsNull() {
		return false
	}
	if !data.QnameWarnAt.IsNull() {
		return false
	}
	if !data.Timeout.IsNull() {
		return false
	}
	if !data.Statistics.IsNull() {
		return false
	}
	if !data.LogLevel.IsNull() {
		return false
	}
	if !data.EnvironmentalLog.IsNull() {
		return false
	}
	if !data.BacklogSize.IsNull() {
		return false
	}
	if !data.BacklogTimeout.IsNull() {
		return false
	}
	return true
}

func (data Throttler) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "Throttler.name", path.Base("/mgmt/config/default/Throttler/Throttler"))
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.ThrottleAt.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ThrottleAt`, data.ThrottleAt.ValueInt64())
	}
	if !data.TerminateAt.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TerminateAt`, data.TerminateAt.ValueInt64())
	}
	if !data.TempFsThrottleAt.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TempFSThrottleAt`, data.TempFsThrottleAt.ValueInt64())
	}
	if !data.TempFsTerminateAt.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TempFSTerminateAt`, data.TempFsTerminateAt.ValueInt64())
	}
	if !data.QnameWarnAt.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`QnameWarnAt`, data.QnameWarnAt.ValueInt64())
	}
	if !data.Timeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Timeout`, data.Timeout.ValueInt64())
	}
	if !data.Statistics.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Statistics`, tfutils.StringFromBool(data.Statistics, ""))
	}
	if !data.LogLevel.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LogLevel`, data.LogLevel.ValueString())
	}
	if !data.EnvironmentalLog.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnvironmentalLog`, tfutils.StringFromBool(data.EnvironmentalLog, ""))
	}
	if !data.BacklogSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BacklogSize`, data.BacklogSize.ValueInt64())
	}
	if !data.BacklogTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BacklogTimeout`, data.BacklogTimeout.ValueInt64())
	}
	return body
}

func (data *Throttler) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `ThrottleAt`); value.Exists() {
		data.ThrottleAt = types.Int64Value(value.Int())
	} else {
		data.ThrottleAt = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `TerminateAt`); value.Exists() {
		data.TerminateAt = types.Int64Value(value.Int())
	} else {
		data.TerminateAt = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `TempFSThrottleAt`); value.Exists() {
		data.TempFsThrottleAt = types.Int64Value(value.Int())
	} else {
		data.TempFsThrottleAt = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `TempFSTerminateAt`); value.Exists() {
		data.TempFsTerminateAt = types.Int64Value(value.Int())
	} else {
		data.TempFsTerminateAt = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `QnameWarnAt`); value.Exists() {
		data.QnameWarnAt = types.Int64Value(value.Int())
	} else {
		data.QnameWarnAt = types.Int64Value(10)
	}
	if value := res.Get(pathRoot + `Timeout`); value.Exists() {
		data.Timeout = types.Int64Value(value.Int())
	} else {
		data.Timeout = types.Int64Value(30)
	}
	if value := res.Get(pathRoot + `Statistics`); value.Exists() {
		data.Statistics = tfutils.BoolFromString(value.String())
	} else {
		data.Statistics = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LogLevel`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LogLevel = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LogLevel = types.StringValue("debug")
	}
	if value := res.Get(pathRoot + `EnvironmentalLog`); value.Exists() {
		data.EnvironmentalLog = tfutils.BoolFromString(value.String())
	} else {
		data.EnvironmentalLog = types.BoolNull()
	}
	if value := res.Get(pathRoot + `BacklogSize`); value.Exists() {
		data.BacklogSize = types.Int64Value(value.Int())
	} else {
		data.BacklogSize = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `BacklogTimeout`); value.Exists() {
		data.BacklogTimeout = types.Int64Value(value.Int())
	} else {
		data.BacklogTimeout = types.Int64Value(30)
	}
}

func (data *Throttler) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else if !data.Enabled.ValueBool() {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `ThrottleAt`); value.Exists() && !data.ThrottleAt.IsNull() {
		data.ThrottleAt = types.Int64Value(value.Int())
	} else if data.ThrottleAt.ValueInt64() != 0 {
		data.ThrottleAt = types.Int64Null()
	}
	if value := res.Get(pathRoot + `TerminateAt`); value.Exists() && !data.TerminateAt.IsNull() {
		data.TerminateAt = types.Int64Value(value.Int())
	} else if data.TerminateAt.ValueInt64() != 0 {
		data.TerminateAt = types.Int64Null()
	}
	if value := res.Get(pathRoot + `TempFSThrottleAt`); value.Exists() && !data.TempFsThrottleAt.IsNull() {
		data.TempFsThrottleAt = types.Int64Value(value.Int())
	} else if data.TempFsThrottleAt.ValueInt64() != 0 {
		data.TempFsThrottleAt = types.Int64Null()
	}
	if value := res.Get(pathRoot + `TempFSTerminateAt`); value.Exists() && !data.TempFsTerminateAt.IsNull() {
		data.TempFsTerminateAt = types.Int64Value(value.Int())
	} else if data.TempFsTerminateAt.ValueInt64() != 0 {
		data.TempFsTerminateAt = types.Int64Null()
	}
	if value := res.Get(pathRoot + `QnameWarnAt`); value.Exists() && !data.QnameWarnAt.IsNull() {
		data.QnameWarnAt = types.Int64Value(value.Int())
	} else if data.QnameWarnAt.ValueInt64() != 10 {
		data.QnameWarnAt = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Timeout`); value.Exists() && !data.Timeout.IsNull() {
		data.Timeout = types.Int64Value(value.Int())
	} else if data.Timeout.ValueInt64() != 30 {
		data.Timeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Statistics`); value.Exists() && !data.Statistics.IsNull() {
		data.Statistics = tfutils.BoolFromString(value.String())
	} else if data.Statistics.ValueBool() {
		data.Statistics = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LogLevel`); value.Exists() && !data.LogLevel.IsNull() {
		data.LogLevel = tfutils.ParseStringFromGJSON(value)
	} else if data.LogLevel.ValueString() != "debug" {
		data.LogLevel = types.StringNull()
	}
	if value := res.Get(pathRoot + `EnvironmentalLog`); value.Exists() && !data.EnvironmentalLog.IsNull() {
		data.EnvironmentalLog = tfutils.BoolFromString(value.String())
	} else if !data.EnvironmentalLog.ValueBool() {
		data.EnvironmentalLog = types.BoolNull()
	}
	if value := res.Get(pathRoot + `BacklogSize`); value.Exists() && !data.BacklogSize.IsNull() {
		data.BacklogSize = types.Int64Value(value.Int())
	} else if data.BacklogSize.ValueInt64() != 0 {
		data.BacklogSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `BacklogTimeout`); value.Exists() && !data.BacklogTimeout.IsNull() {
		data.BacklogTimeout = types.Int64Value(value.Int())
	} else if data.BacklogTimeout.ValueInt64() != 30 {
		data.BacklogTimeout = types.Int64Null()
	}
}
