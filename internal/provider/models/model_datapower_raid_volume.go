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

type RaidVolume struct {
	UserSummary       types.String                `tfsdk:"user_summary"`
	ReadOnly          types.Bool                  `tfsdk:"read_only"`
	Directory         types.String                `tfsdk:"directory"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var RaidVolumeObjectType = map[string]attr.Type{
	"user_summary":       types.StringType,
	"read_only":          types.BoolType,
	"directory":          types.StringType,
	"dependency_actions": actions.ActionsListType,
}

func (data RaidVolume) GetPath() string {
	rest_path := "/mgmt/config/default/RaidVolume/raid0"
	return rest_path
}

func (data RaidVolume) IsNull() bool {
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.ReadOnly.IsNull() {
		return false
	}
	if !data.Directory.IsNull() {
		return false
	}
	return true
}
func (data *RaidVolume) ToDefault() {
	data.UserSummary = types.StringNull()
	data.ReadOnly = types.BoolValue(false)
	data.Directory = types.StringNull()
}

func (data RaidVolume) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "RaidVolume.name", path.Base("/mgmt/config/default/RaidVolume/raid0"))

	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.ReadOnly.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ReadOnly`, tfutils.StringFromBool(data.ReadOnly, ""))
	}
	if !data.Directory.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Directory`, data.Directory.ValueString())
	}
	return body
}

func (data *RaidVolume) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `ReadOnly`); value.Exists() {
		data.ReadOnly = tfutils.BoolFromString(value.String())
	} else {
		data.ReadOnly = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Directory`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Directory = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Directory = types.StringNull()
	}
}

func (data *RaidVolume) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `ReadOnly`); value.Exists() && !data.ReadOnly.IsNull() {
		data.ReadOnly = tfutils.BoolFromString(value.String())
	} else if data.ReadOnly.ValueBool() {
		data.ReadOnly = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Directory`); value.Exists() && !data.Directory.IsNull() {
		data.Directory = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Directory = types.StringNull()
	}
}
