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

type AuditLog struct {
	Enabled    types.Bool   `tfsdk:"enabled"`
	Size       types.Int64  `tfsdk:"size"`
	Rotate     types.Int64  `tfsdk:"rotate"`
	AuditLevel types.String `tfsdk:"audit_level"`
}

var AuditLogObjectType = map[string]attr.Type{
	"enabled":     types.BoolType,
	"size":        types.Int64Type,
	"rotate":      types.Int64Type,
	"audit_level": types.StringType,
}

func (data AuditLog) GetPath() string {
	rest_path := "/mgmt/config/default/AuditLog/AuditLog-Settings"
	return rest_path
}

func (data AuditLog) IsNull() bool {
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.Size.IsNull() {
		return false
	}
	if !data.Rotate.IsNull() {
		return false
	}
	if !data.AuditLevel.IsNull() {
		return false
	}
	return true
}

func (data AuditLog) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "AuditLog.name", path.Base("/mgmt/config/default/AuditLog/AuditLog-Settings"))
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, true))
	}
	if !data.Size.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Size`, data.Size.ValueInt64())
	}
	if !data.Rotate.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Rotate`, data.Rotate.ValueInt64())
	}
	if !data.AuditLevel.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AuditLevel`, data.AuditLevel.ValueString())
	}
	return body
}

func (data *AuditLog) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Size`); value.Exists() {
		data.Size = types.Int64Value(value.Int())
	} else {
		data.Size = types.Int64Value(1000)
	}
	if value := res.Get(pathRoot + `Rotate`); value.Exists() {
		data.Rotate = types.Int64Value(value.Int())
	} else {
		data.Rotate = types.Int64Value(3)
	}
	if value := res.Get(pathRoot + `AuditLevel`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuditLevel = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuditLevel = types.StringValue("standard")
	}
}

func (data *AuditLog) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else if !data.Enabled.ValueBool() {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Size`); value.Exists() && !data.Size.IsNull() {
		data.Size = types.Int64Value(value.Int())
	} else if data.Size.ValueInt64() != 1000 {
		data.Size = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Rotate`); value.Exists() && !data.Rotate.IsNull() {
		data.Rotate = types.Int64Value(value.Int())
	} else if data.Rotate.ValueInt64() != 3 {
		data.Rotate = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AuditLevel`); value.Exists() && !data.AuditLevel.IsNull() {
		data.AuditLevel = tfutils.ParseStringFromGJSON(value)
	} else if data.AuditLevel.ValueString() != "standard" {
		data.AuditLevel = types.StringNull()
	}
}
