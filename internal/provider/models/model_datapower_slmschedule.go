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

type SLMSchedule struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	DaysOfWeek        *DmWeekdayBitmap            `tfsdk:"days_of_week"`
	StartTime         types.String                `tfsdk:"start_time"`
	Duration          types.Int64                 `tfsdk:"duration"`
	StartDate         types.String                `tfsdk:"start_date"`
	StopDate          types.String                `tfsdk:"stop_date"`
	TimeZone          types.String                `tfsdk:"time_zone"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var SLMScheduleObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"user_summary":       types.StringType,
	"days_of_week":       types.ObjectType{AttrTypes: DmWeekdayBitmapObjectType},
	"start_time":         types.StringType,
	"duration":           types.Int64Type,
	"start_date":         types.StringType,
	"stop_date":          types.StringType,
	"time_zone":          types.StringType,
	"dependency_actions": actions.ActionsListType,
}

func (data SLMSchedule) GetPath() string {
	rest_path := "/mgmt/config/{domain}/SLMSchedule"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data SLMSchedule) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if data.DaysOfWeek != nil {
		if !data.DaysOfWeek.IsNull() {
			return false
		}
	}
	if !data.StartTime.IsNull() {
		return false
	}
	if !data.Duration.IsNull() {
		return false
	}
	if !data.StartDate.IsNull() {
		return false
	}
	if !data.StopDate.IsNull() {
		return false
	}
	if !data.TimeZone.IsNull() {
		return false
	}
	return true
}

func (data SLMSchedule) ToBody(ctx context.Context, pathRoot string) string {
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
	if data.DaysOfWeek != nil {
		if !data.DaysOfWeek.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`DaysOfWeek`, data.DaysOfWeek.ToBody(ctx, ""))
		}
	}
	if !data.StartTime.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StartTime`, data.StartTime.ValueString())
	}
	if !data.Duration.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Duration`, data.Duration.ValueInt64())
	}
	if !data.StartDate.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StartDate`, data.StartDate.ValueString())
	}
	if !data.StopDate.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StopDate`, data.StopDate.ValueString())
	}
	if !data.TimeZone.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TimeZone`, data.TimeZone.ValueString())
	}
	return body
}

func (data *SLMSchedule) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `DaysOfWeek`); value.Exists() {
		data.DaysOfWeek = &DmWeekdayBitmap{}
		data.DaysOfWeek.FromBody(ctx, "", value)
	} else {
		data.DaysOfWeek = nil
	}
	if value := res.Get(pathRoot + `StartTime`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.StartTime = tfutils.ParseStringFromGJSON(value)
	} else {
		data.StartTime = types.StringNull()
	}
	if value := res.Get(pathRoot + `Duration`); value.Exists() {
		data.Duration = types.Int64Value(value.Int())
	} else {
		data.Duration = types.Int64Value(1440)
	}
	if value := res.Get(pathRoot + `StartDate`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.StartDate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.StartDate = types.StringNull()
	}
	if value := res.Get(pathRoot + `StopDate`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.StopDate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.StopDate = types.StringNull()
	}
	if value := res.Get(pathRoot + `TimeZone`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TimeZone = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TimeZone = types.StringValue("LOCAL")
	}
}

func (data *SLMSchedule) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `DaysOfWeek`); value.Exists() {
		data.DaysOfWeek.UpdateFromBody(ctx, "", value)
	} else {
		data.DaysOfWeek = nil
	}
	if value := res.Get(pathRoot + `StartTime`); value.Exists() && !data.StartTime.IsNull() {
		data.StartTime = tfutils.ParseStringFromGJSON(value)
	} else {
		data.StartTime = types.StringNull()
	}
	if value := res.Get(pathRoot + `Duration`); value.Exists() && !data.Duration.IsNull() {
		data.Duration = types.Int64Value(value.Int())
	} else if data.Duration.ValueInt64() != 1440 {
		data.Duration = types.Int64Null()
	}
	if value := res.Get(pathRoot + `StartDate`); value.Exists() && !data.StartDate.IsNull() {
		data.StartDate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.StartDate = types.StringNull()
	}
	if value := res.Get(pathRoot + `StopDate`); value.Exists() && !data.StopDate.IsNull() {
		data.StopDate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.StopDate = types.StringNull()
	}
	if value := res.Get(pathRoot + `TimeZone`); value.Exists() && !data.TimeZone.IsNull() {
		data.TimeZone = tfutils.ParseStringFromGJSON(value)
	} else if data.TimeZone.ValueString() != "LOCAL" {
		data.TimeZone = types.StringNull()
	}
}
