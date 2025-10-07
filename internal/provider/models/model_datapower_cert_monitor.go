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

type CertMonitor struct {
	Enabled             types.Bool                  `tfsdk:"enabled"`
	UserSummary         types.String                `tfsdk:"user_summary"`
	PollingInterval     types.Int64                 `tfsdk:"polling_interval"`
	ReminderTime        types.Int64                 `tfsdk:"reminder_time"`
	LogLevel            types.String                `tfsdk:"log_level"`
	DisableExpiredCerts types.Bool                  `tfsdk:"disable_expired_certs"`
	DependencyActions   []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var CertMonitorObjectType = map[string]attr.Type{
	"enabled":               types.BoolType,
	"user_summary":          types.StringType,
	"polling_interval":      types.Int64Type,
	"reminder_time":         types.Int64Type,
	"log_level":             types.StringType,
	"disable_expired_certs": types.BoolType,
	"dependency_actions":    actions.ActionsListType,
}

func (data CertMonitor) GetPath() string {
	rest_path := "/mgmt/config/default/CertMonitor/Certificate Monitor"
	return rest_path
}

func (data CertMonitor) IsNull() bool {
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.PollingInterval.IsNull() {
		return false
	}
	if !data.ReminderTime.IsNull() {
		return false
	}
	if !data.LogLevel.IsNull() {
		return false
	}
	if !data.DisableExpiredCerts.IsNull() {
		return false
	}
	return true
}
func (data *CertMonitor) ToDefault() {
	data.Enabled = types.BoolValue(true)
	data.UserSummary = types.StringNull()
	data.PollingInterval = types.Int64Value(1)
	data.ReminderTime = types.Int64Value(30)
	data.LogLevel = types.StringValue("warn")
	data.DisableExpiredCerts = types.BoolValue(false)
}

func (data CertMonitor) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "CertMonitor.name", path.Base("/mgmt/config/default/CertMonitor/Certificate Monitor"))

	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.PollingInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PollingInterval`, data.PollingInterval.ValueInt64())
	}
	if !data.ReminderTime.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ReminderTime`, data.ReminderTime.ValueInt64())
	}
	if !data.LogLevel.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LogLevel`, data.LogLevel.ValueString())
	}
	if !data.DisableExpiredCerts.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DisableExpiredCerts`, tfutils.StringFromBool(data.DisableExpiredCerts, ""))
	}
	return body
}

func (data *CertMonitor) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `PollingInterval`); value.Exists() {
		data.PollingInterval = types.Int64Value(value.Int())
	} else {
		data.PollingInterval = types.Int64Value(1)
	}
	if value := res.Get(pathRoot + `ReminderTime`); value.Exists() {
		data.ReminderTime = types.Int64Value(value.Int())
	} else {
		data.ReminderTime = types.Int64Value(30)
	}
	if value := res.Get(pathRoot + `LogLevel`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LogLevel = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LogLevel = types.StringValue("warn")
	}
	if value := res.Get(pathRoot + `DisableExpiredCerts`); value.Exists() {
		data.DisableExpiredCerts = tfutils.BoolFromString(value.String())
	} else {
		data.DisableExpiredCerts = types.BoolNull()
	}
}

func (data *CertMonitor) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `PollingInterval`); value.Exists() && !data.PollingInterval.IsNull() {
		data.PollingInterval = types.Int64Value(value.Int())
	} else if data.PollingInterval.ValueInt64() != 1 {
		data.PollingInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ReminderTime`); value.Exists() && !data.ReminderTime.IsNull() {
		data.ReminderTime = types.Int64Value(value.Int())
	} else if data.ReminderTime.ValueInt64() != 30 {
		data.ReminderTime = types.Int64Null()
	}
	if value := res.Get(pathRoot + `LogLevel`); value.Exists() && !data.LogLevel.IsNull() {
		data.LogLevel = tfutils.ParseStringFromGJSON(value)
	} else if data.LogLevel.ValueString() != "warn" {
		data.LogLevel = types.StringNull()
	}
	if value := res.Get(pathRoot + `DisableExpiredCerts`); value.Exists() && !data.DisableExpiredCerts.IsNull() {
		data.DisableExpiredCerts = tfutils.BoolFromString(value.String())
	} else if data.DisableExpiredCerts.ValueBool() {
		data.DisableExpiredCerts = types.BoolNull()
	}
}
