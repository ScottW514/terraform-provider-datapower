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
	"path"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type WebServicesAgent struct {
	AppDomain         types.String      `tfsdk:"app_domain"`
	Enabled           types.Bool        `tfsdk:"enabled"`
	UserSummary       types.String      `tfsdk:"user_summary"`
	MaxRecords        types.Int64       `tfsdk:"max_records"`
	MaxMemoryKb       types.Int64       `tfsdk:"max_memory_kb"`
	CaptureMode       types.String      `tfsdk:"capture_mode"`
	MediationMetrics  types.Bool        `tfsdk:"mediation_metrics"`
	MaxPayloadSizeKb  types.Int64       `tfsdk:"max_payload_size_kb"`
	DependencyActions []*actions.Action `tfsdk:"dependency_actions"`
}

var WebServicesAgentObjectType = map[string]attr.Type{
	"app_domain":          types.StringType,
	"enabled":             types.BoolType,
	"user_summary":        types.StringType,
	"max_records":         types.Int64Type,
	"max_memory_kb":       types.Int64Type,
	"capture_mode":        types.StringType,
	"mediation_metrics":   types.BoolType,
	"max_payload_size_kb": types.Int64Type,
	"dependency_actions":  actions.ActionsListType,
}

func (data WebServicesAgent) GetPath() string {
	rest_path := "/mgmt/config/{domain}/WebServicesAgent/default"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data WebServicesAgent) IsNull() bool {
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.MaxRecords.IsNull() {
		return false
	}
	if !data.MaxMemoryKb.IsNull() {
		return false
	}
	if !data.CaptureMode.IsNull() {
		return false
	}
	if !data.MediationMetrics.IsNull() {
		return false
	}
	if !data.MaxPayloadSizeKb.IsNull() {
		return false
	}
	return true
}

func (data WebServicesAgent) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "WebServicesAgent.name", path.Base("/mgmt/config/{domain}/WebServicesAgent/default"))
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.MaxRecords.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxRecords`, data.MaxRecords.ValueInt64())
	}
	if !data.MaxMemoryKb.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxMemoryKB`, data.MaxMemoryKb.ValueInt64())
	}
	if !data.CaptureMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CaptureMode`, data.CaptureMode.ValueString())
	}
	if !data.MediationMetrics.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MediationMetrics`, tfutils.StringFromBool(data.MediationMetrics, ""))
	}
	if !data.MaxPayloadSizeKb.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxPayloadSizeKB`, data.MaxPayloadSizeKb.ValueInt64())
	}
	return body
}

func (data *WebServicesAgent) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `MaxRecords`); value.Exists() {
		data.MaxRecords = types.Int64Value(value.Int())
	} else {
		data.MaxRecords = types.Int64Value(3000)
	}
	if value := res.Get(pathRoot + `MaxMemoryKB`); value.Exists() {
		data.MaxMemoryKb = types.Int64Value(value.Int())
	} else {
		data.MaxMemoryKb = types.Int64Value(64000)
	}
	if value := res.Get(pathRoot + `CaptureMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CaptureMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CaptureMode = types.StringValue("faults")
	}
	if value := res.Get(pathRoot + `MediationMetrics`); value.Exists() {
		data.MediationMetrics = tfutils.BoolFromString(value.String())
	} else {
		data.MediationMetrics = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MaxPayloadSizeKB`); value.Exists() {
		data.MaxPayloadSizeKb = types.Int64Value(value.Int())
	} else {
		data.MaxPayloadSizeKb = types.Int64Value(0)
	}
}

func (data *WebServicesAgent) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else if data.Enabled.ValueBool() {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `MaxRecords`); value.Exists() && !data.MaxRecords.IsNull() {
		data.MaxRecords = types.Int64Value(value.Int())
	} else if data.MaxRecords.ValueInt64() != 3000 {
		data.MaxRecords = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxMemoryKB`); value.Exists() && !data.MaxMemoryKb.IsNull() {
		data.MaxMemoryKb = types.Int64Value(value.Int())
	} else if data.MaxMemoryKb.ValueInt64() != 64000 {
		data.MaxMemoryKb = types.Int64Null()
	}
	if value := res.Get(pathRoot + `CaptureMode`); value.Exists() && !data.CaptureMode.IsNull() {
		data.CaptureMode = tfutils.ParseStringFromGJSON(value)
	} else if data.CaptureMode.ValueString() != "faults" {
		data.CaptureMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `MediationMetrics`); value.Exists() && !data.MediationMetrics.IsNull() {
		data.MediationMetrics = tfutils.BoolFromString(value.String())
	} else if data.MediationMetrics.ValueBool() {
		data.MediationMetrics = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MaxPayloadSizeKB`); value.Exists() && !data.MaxPayloadSizeKb.IsNull() {
		data.MaxPayloadSizeKb = types.Int64Value(value.Int())
	} else if data.MaxPayloadSizeKb.ValueInt64() != 0 {
		data.MaxPayloadSizeKb = types.Int64Null()
	}
}
