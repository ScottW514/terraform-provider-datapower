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

type NFSClientSettings struct {
	Enabled          types.Bool   `tfsdk:"enabled"`
	UserSummary      types.String `tfsdk:"user_summary"`
	MountRefreshTime types.Int64  `tfsdk:"mount_refresh_time"`
}

var NFSClientSettingsObjectType = map[string]attr.Type{
	"enabled":            types.BoolType,
	"user_summary":       types.StringType,
	"mount_refresh_time": types.Int64Type,
}

func (data NFSClientSettings) GetPath() string {
	rest_path := "/mgmt/config/default/NFSClientSettings/NFS-Client-Settings"
	return rest_path
}

func (data NFSClientSettings) IsNull() bool {
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.MountRefreshTime.IsNull() {
		return false
	}
	return true
}

func (data NFSClientSettings) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "NFSClientSettings.name", path.Base("/mgmt/config/default/NFSClientSettings/NFS-Client-Settings"))
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.MountRefreshTime.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MountRefreshTime`, data.MountRefreshTime.ValueInt64())
	}
	return body
}

func (data *NFSClientSettings) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `MountRefreshTime`); value.Exists() {
		data.MountRefreshTime = types.Int64Value(value.Int())
	} else {
		data.MountRefreshTime = types.Int64Value(10)
	}
}

func (data *NFSClientSettings) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `MountRefreshTime`); value.Exists() && !data.MountRefreshTime.IsNull() {
		data.MountRefreshTime = types.Int64Value(value.Int())
	} else if data.MountRefreshTime.ValueInt64() != 10 {
		data.MountRefreshTime = types.Int64Null()
	}
}
