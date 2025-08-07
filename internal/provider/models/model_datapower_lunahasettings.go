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

type LunaHASettings struct {
	Enabled           types.Bool        `tfsdk:"enabled"`
	Mode              types.String      `tfsdk:"mode"`
	RecoveryCount     types.Int64       `tfsdk:"recovery_count"`
	Interval          types.Int64       `tfsdk:"interval"`
	DependencyActions []*actions.Action `tfsdk:"dependency_actions"`
}

var LunaHASettingsObjectType = map[string]attr.Type{
	"enabled":            types.BoolType,
	"mode":               types.StringType,
	"recovery_count":     types.Int64Type,
	"interval":           types.Int64Type,
	"dependency_actions": actions.ActionsListType,
}

func (data LunaHASettings) GetPath() string {
	rest_path := "/mgmt/config/default/LunaHASettings/LunaHASettings"
	return rest_path
}

func (data LunaHASettings) IsNull() bool {
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.Mode.IsNull() {
		return false
	}
	if !data.RecoveryCount.IsNull() {
		return false
	}
	if !data.Interval.IsNull() {
		return false
	}
	return true
}

func (data LunaHASettings) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "LunaHASettings.name", path.Base("/mgmt/config/default/LunaHASettings/LunaHASettings"))
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.Mode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Mode`, data.Mode.ValueString())
	}
	if !data.RecoveryCount.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Count`, data.RecoveryCount.ValueInt64())
	}
	if !data.Interval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Interval`, data.Interval.ValueInt64())
	}
	return body
}

func (data *LunaHASettings) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Mode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Mode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Mode = types.StringValue("activeBasic")
	}
	if value := res.Get(pathRoot + `Count`); value.Exists() {
		data.RecoveryCount = types.Int64Value(value.Int())
	} else {
		data.RecoveryCount = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Interval`); value.Exists() {
		data.Interval = types.Int64Value(value.Int())
	} else {
		data.Interval = types.Int64Value(60)
	}
}

func (data *LunaHASettings) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else if !data.Enabled.ValueBool() {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Mode`); value.Exists() && !data.Mode.IsNull() {
		data.Mode = tfutils.ParseStringFromGJSON(value)
	} else if data.Mode.ValueString() != "activeBasic" {
		data.Mode = types.StringNull()
	}
	if value := res.Get(pathRoot + `Count`); value.Exists() && !data.RecoveryCount.IsNull() {
		data.RecoveryCount = types.Int64Value(value.Int())
	} else {
		data.RecoveryCount = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Interval`); value.Exists() && !data.Interval.IsNull() {
		data.Interval = types.Int64Value(value.Int())
	} else if data.Interval.ValueInt64() != 60 {
		data.Interval = types.Int64Null()
	}
}
