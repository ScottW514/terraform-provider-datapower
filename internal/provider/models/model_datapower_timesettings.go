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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type TimeSettings struct {
	Enabled                  types.Bool   `tfsdk:"enabled"`
	LocalTimeZone            types.String `tfsdk:"local_time_zone"`
	CustomTzName             types.String `tfsdk:"custom_tz_name"`
	UtcDirection             types.String `tfsdk:"utc_direction"`
	OffsetHours              types.Int64  `tfsdk:"offset_hours"`
	OffsetMinutes            types.Int64  `tfsdk:"offset_minutes"`
	DaylightOffsetHours      types.Int64  `tfsdk:"daylight_offset_hours"`
	TzNameDst                types.String `tfsdk:"tz_name_dst"`
	DaylightStartMonth       types.String `tfsdk:"daylight_start_month"`
	DaylightStartWeek        types.Int64  `tfsdk:"daylight_start_week"`
	DaylightStartDay         types.String `tfsdk:"daylight_start_day"`
	DaylightStartTimeHours   types.Int64  `tfsdk:"daylight_start_time_hours"`
	DaylightStartTimeMinutes types.Int64  `tfsdk:"daylight_start_time_minutes"`
	DaylightStopMonth        types.String `tfsdk:"daylight_stop_month"`
	DaylightStopWeek         types.Int64  `tfsdk:"daylight_stop_week"`
	DaylightStopDay          types.String `tfsdk:"daylight_stop_day"`
	DaylightStopTimeHours    types.Int64  `tfsdk:"daylight_stop_time_hours"`
	DaylightStopTimeMinutes  types.Int64  `tfsdk:"daylight_stop_time_minutes"`
}

var TimeSettingsObjectType = map[string]attr.Type{
	"enabled":                     types.BoolType,
	"local_time_zone":             types.StringType,
	"custom_tz_name":              types.StringType,
	"utc_direction":               types.StringType,
	"offset_hours":                types.Int64Type,
	"offset_minutes":              types.Int64Type,
	"daylight_offset_hours":       types.Int64Type,
	"tz_name_dst":                 types.StringType,
	"daylight_start_month":        types.StringType,
	"daylight_start_week":         types.Int64Type,
	"daylight_start_day":          types.StringType,
	"daylight_start_time_hours":   types.Int64Type,
	"daylight_start_time_minutes": types.Int64Type,
	"daylight_stop_month":         types.StringType,
	"daylight_stop_week":          types.Int64Type,
	"daylight_stop_day":           types.StringType,
	"daylight_stop_time_hours":    types.Int64Type,
	"daylight_stop_time_minutes":  types.Int64Type,
}

func (data TimeSettings) GetPath() string {
	rest_path := "/mgmt/config/default/TimeSettings/Time"
	return rest_path
}

func (data TimeSettings) IsNull() bool {
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.LocalTimeZone.IsNull() {
		return false
	}
	if !data.CustomTzName.IsNull() {
		return false
	}
	if !data.UtcDirection.IsNull() {
		return false
	}
	if !data.OffsetHours.IsNull() {
		return false
	}
	if !data.OffsetMinutes.IsNull() {
		return false
	}
	if !data.DaylightOffsetHours.IsNull() {
		return false
	}
	if !data.TzNameDst.IsNull() {
		return false
	}
	if !data.DaylightStartMonth.IsNull() {
		return false
	}
	if !data.DaylightStartWeek.IsNull() {
		return false
	}
	if !data.DaylightStartDay.IsNull() {
		return false
	}
	if !data.DaylightStartTimeHours.IsNull() {
		return false
	}
	if !data.DaylightStartTimeMinutes.IsNull() {
		return false
	}
	if !data.DaylightStopMonth.IsNull() {
		return false
	}
	if !data.DaylightStopWeek.IsNull() {
		return false
	}
	if !data.DaylightStopDay.IsNull() {
		return false
	}
	if !data.DaylightStopTimeHours.IsNull() {
		return false
	}
	if !data.DaylightStopTimeMinutes.IsNull() {
		return false
	}
	return true
}

func (data TimeSettings) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "TimeSettings.name", path.Base("/mgmt/config/default/TimeSettings/Time"))
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.LocalTimeZone.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalTimeZone`, data.LocalTimeZone.ValueString())
	}
	if !data.CustomTzName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CustomTZName`, data.CustomTzName.ValueString())
	}
	if !data.UtcDirection.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UTCDirection`, data.UtcDirection.ValueString())
	}
	if !data.OffsetHours.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OffsetHours`, data.OffsetHours.ValueInt64())
	}
	if !data.OffsetMinutes.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OffsetMinutes`, data.OffsetMinutes.ValueInt64())
	}
	if !data.DaylightOffsetHours.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DaylightOffsetHours`, data.DaylightOffsetHours.ValueInt64())
	}
	if !data.TzNameDst.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TZNameDST`, data.TzNameDst.ValueString())
	}
	if !data.DaylightStartMonth.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DaylightStartMonth`, data.DaylightStartMonth.ValueString())
	}
	if !data.DaylightStartWeek.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DaylightStartWeek`, data.DaylightStartWeek.ValueInt64())
	}
	if !data.DaylightStartDay.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DaylightStartDay`, data.DaylightStartDay.ValueString())
	}
	if !data.DaylightStartTimeHours.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DaylightStartTimeHours`, data.DaylightStartTimeHours.ValueInt64())
	}
	if !data.DaylightStartTimeMinutes.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DaylightStartTimeMinutes`, data.DaylightStartTimeMinutes.ValueInt64())
	}
	if !data.DaylightStopMonth.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DaylightStopMonth`, data.DaylightStopMonth.ValueString())
	}
	if !data.DaylightStopWeek.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DaylightStopWeek`, data.DaylightStopWeek.ValueInt64())
	}
	if !data.DaylightStopDay.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DaylightStopDay`, data.DaylightStopDay.ValueString())
	}
	if !data.DaylightStopTimeHours.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DaylightStopTimeHours`, data.DaylightStopTimeHours.ValueInt64())
	}
	if !data.DaylightStopTimeMinutes.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DaylightStopTimeMinutes`, data.DaylightStopTimeMinutes.ValueInt64())
	}
	return body
}

func (data *TimeSettings) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LocalTimeZone`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalTimeZone = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalTimeZone = types.StringValue("EST5EDT")
	}
	if value := res.Get(pathRoot + `CustomTZName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CustomTzName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CustomTzName = types.StringValue("STD")
	}
	if value := res.Get(pathRoot + `UTCDirection`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UtcDirection = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UtcDirection = types.StringValue("West")
	}
	if value := res.Get(pathRoot + `OffsetHours`); value.Exists() {
		data.OffsetHours = types.Int64Value(value.Int())
	} else {
		data.OffsetHours = types.Int64Null()
	}
	if value := res.Get(pathRoot + `OffsetMinutes`); value.Exists() {
		data.OffsetMinutes = types.Int64Value(value.Int())
	} else {
		data.OffsetMinutes = types.Int64Null()
	}
	if value := res.Get(pathRoot + `DaylightOffsetHours`); value.Exists() {
		data.DaylightOffsetHours = types.Int64Value(value.Int())
	} else {
		data.DaylightOffsetHours = types.Int64Value(1)
	}
	if value := res.Get(pathRoot + `TZNameDST`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TzNameDst = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TzNameDst = types.StringValue("DST")
	}
	if value := res.Get(pathRoot + `DaylightStartMonth`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DaylightStartMonth = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DaylightStartMonth = types.StringValue("March")
	}
	if value := res.Get(pathRoot + `DaylightStartWeek`); value.Exists() {
		data.DaylightStartWeek = types.Int64Value(value.Int())
	} else {
		data.DaylightStartWeek = types.Int64Value(2)
	}
	if value := res.Get(pathRoot + `DaylightStartDay`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DaylightStartDay = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DaylightStartDay = types.StringValue("Sunday")
	}
	if value := res.Get(pathRoot + `DaylightStartTimeHours`); value.Exists() {
		data.DaylightStartTimeHours = types.Int64Value(value.Int())
	} else {
		data.DaylightStartTimeHours = types.Int64Value(2)
	}
	if value := res.Get(pathRoot + `DaylightStartTimeMinutes`); value.Exists() {
		data.DaylightStartTimeMinutes = types.Int64Value(value.Int())
	} else {
		data.DaylightStartTimeMinutes = types.Int64Null()
	}
	if value := res.Get(pathRoot + `DaylightStopMonth`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DaylightStopMonth = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DaylightStopMonth = types.StringValue("November")
	}
	if value := res.Get(pathRoot + `DaylightStopWeek`); value.Exists() {
		data.DaylightStopWeek = types.Int64Value(value.Int())
	} else {
		data.DaylightStopWeek = types.Int64Value(1)
	}
	if value := res.Get(pathRoot + `DaylightStopDay`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DaylightStopDay = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DaylightStopDay = types.StringValue("Sunday")
	}
	if value := res.Get(pathRoot + `DaylightStopTimeHours`); value.Exists() {
		data.DaylightStopTimeHours = types.Int64Value(value.Int())
	} else {
		data.DaylightStopTimeHours = types.Int64Value(2)
	}
	if value := res.Get(pathRoot + `DaylightStopTimeMinutes`); value.Exists() {
		data.DaylightStopTimeMinutes = types.Int64Value(value.Int())
	} else {
		data.DaylightStopTimeMinutes = types.Int64Null()
	}
}

func (data *TimeSettings) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else if !data.Enabled.ValueBool() {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LocalTimeZone`); value.Exists() && !data.LocalTimeZone.IsNull() {
		data.LocalTimeZone = tfutils.ParseStringFromGJSON(value)
	} else if data.LocalTimeZone.ValueString() != "EST5EDT" {
		data.LocalTimeZone = types.StringNull()
	}
	if value := res.Get(pathRoot + `CustomTZName`); value.Exists() && !data.CustomTzName.IsNull() {
		data.CustomTzName = tfutils.ParseStringFromGJSON(value)
	} else if data.CustomTzName.ValueString() != "STD" {
		data.CustomTzName = types.StringNull()
	}
	if value := res.Get(pathRoot + `UTCDirection`); value.Exists() && !data.UtcDirection.IsNull() {
		data.UtcDirection = tfutils.ParseStringFromGJSON(value)
	} else if data.UtcDirection.ValueString() != "West" {
		data.UtcDirection = types.StringNull()
	}
	if value := res.Get(pathRoot + `OffsetHours`); value.Exists() && !data.OffsetHours.IsNull() {
		data.OffsetHours = types.Int64Value(value.Int())
	} else {
		data.OffsetHours = types.Int64Null()
	}
	if value := res.Get(pathRoot + `OffsetMinutes`); value.Exists() && !data.OffsetMinutes.IsNull() {
		data.OffsetMinutes = types.Int64Value(value.Int())
	} else {
		data.OffsetMinutes = types.Int64Null()
	}
	if value := res.Get(pathRoot + `DaylightOffsetHours`); value.Exists() && !data.DaylightOffsetHours.IsNull() {
		data.DaylightOffsetHours = types.Int64Value(value.Int())
	} else if data.DaylightOffsetHours.ValueInt64() != 1 {
		data.DaylightOffsetHours = types.Int64Null()
	}
	if value := res.Get(pathRoot + `TZNameDST`); value.Exists() && !data.TzNameDst.IsNull() {
		data.TzNameDst = tfutils.ParseStringFromGJSON(value)
	} else if data.TzNameDst.ValueString() != "DST" {
		data.TzNameDst = types.StringNull()
	}
	if value := res.Get(pathRoot + `DaylightStartMonth`); value.Exists() && !data.DaylightStartMonth.IsNull() {
		data.DaylightStartMonth = tfutils.ParseStringFromGJSON(value)
	} else if data.DaylightStartMonth.ValueString() != "March" {
		data.DaylightStartMonth = types.StringNull()
	}
	if value := res.Get(pathRoot + `DaylightStartWeek`); value.Exists() && !data.DaylightStartWeek.IsNull() {
		data.DaylightStartWeek = types.Int64Value(value.Int())
	} else if data.DaylightStartWeek.ValueInt64() != 2 {
		data.DaylightStartWeek = types.Int64Null()
	}
	if value := res.Get(pathRoot + `DaylightStartDay`); value.Exists() && !data.DaylightStartDay.IsNull() {
		data.DaylightStartDay = tfutils.ParseStringFromGJSON(value)
	} else if data.DaylightStartDay.ValueString() != "Sunday" {
		data.DaylightStartDay = types.StringNull()
	}
	if value := res.Get(pathRoot + `DaylightStartTimeHours`); value.Exists() && !data.DaylightStartTimeHours.IsNull() {
		data.DaylightStartTimeHours = types.Int64Value(value.Int())
	} else if data.DaylightStartTimeHours.ValueInt64() != 2 {
		data.DaylightStartTimeHours = types.Int64Null()
	}
	if value := res.Get(pathRoot + `DaylightStartTimeMinutes`); value.Exists() && !data.DaylightStartTimeMinutes.IsNull() {
		data.DaylightStartTimeMinutes = types.Int64Value(value.Int())
	} else {
		data.DaylightStartTimeMinutes = types.Int64Null()
	}
	if value := res.Get(pathRoot + `DaylightStopMonth`); value.Exists() && !data.DaylightStopMonth.IsNull() {
		data.DaylightStopMonth = tfutils.ParseStringFromGJSON(value)
	} else if data.DaylightStopMonth.ValueString() != "November" {
		data.DaylightStopMonth = types.StringNull()
	}
	if value := res.Get(pathRoot + `DaylightStopWeek`); value.Exists() && !data.DaylightStopWeek.IsNull() {
		data.DaylightStopWeek = types.Int64Value(value.Int())
	} else if data.DaylightStopWeek.ValueInt64() != 1 {
		data.DaylightStopWeek = types.Int64Null()
	}
	if value := res.Get(pathRoot + `DaylightStopDay`); value.Exists() && !data.DaylightStopDay.IsNull() {
		data.DaylightStopDay = tfutils.ParseStringFromGJSON(value)
	} else if data.DaylightStopDay.ValueString() != "Sunday" {
		data.DaylightStopDay = types.StringNull()
	}
	if value := res.Get(pathRoot + `DaylightStopTimeHours`); value.Exists() && !data.DaylightStopTimeHours.IsNull() {
		data.DaylightStopTimeHours = types.Int64Value(value.Int())
	} else if data.DaylightStopTimeHours.ValueInt64() != 2 {
		data.DaylightStopTimeHours = types.Int64Null()
	}
	if value := res.Get(pathRoot + `DaylightStopTimeMinutes`); value.Exists() && !data.DaylightStopTimeMinutes.IsNull() {
		data.DaylightStopTimeMinutes = types.Int64Value(value.Int())
	} else {
		data.DaylightStopTimeMinutes = types.Int64Null()
	}
}
