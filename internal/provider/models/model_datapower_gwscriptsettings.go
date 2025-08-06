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

type GWScriptSettings struct {
	Enabled                types.Bool        `tfsdk:"enabled"`
	UserSummary            types.String      `tfsdk:"user_summary"`
	FrozenEnabled          types.Bool        `tfsdk:"frozen_enabled"`
	UntrustedCodeMitigated types.Bool        `tfsdk:"untrusted_code_mitigated"`
	ReloadNeeded           types.Bool        `tfsdk:"reload_needed"`
	TerminateTime          types.Int64       `tfsdk:"terminate_time"`
	ObjectActions          []*actions.Action `tfsdk:"object_actions"`
}

var GWScriptSettingsObjectType = map[string]attr.Type{
	"enabled":                  types.BoolType,
	"user_summary":             types.StringType,
	"frozen_enabled":           types.BoolType,
	"untrusted_code_mitigated": types.BoolType,
	"reload_needed":            types.BoolType,
	"terminate_time":           types.Int64Type,
	"object_actions":           actions.ActionsListType,
}

func (data GWScriptSettings) GetPath() string {
	rest_path := "/mgmt/config/default/GWScriptSettings/GatewayScript-Settings"
	return rest_path
}

func (data GWScriptSettings) IsNull() bool {
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.FrozenEnabled.IsNull() {
		return false
	}
	if !data.UntrustedCodeMitigated.IsNull() {
		return false
	}
	if !data.ReloadNeeded.IsNull() {
		return false
	}
	if !data.TerminateTime.IsNull() {
		return false
	}
	return true
}

func (data GWScriptSettings) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "GWScriptSettings.name", path.Base("/mgmt/config/default/GWScriptSettings/GatewayScript-Settings"))
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.FrozenEnabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FrozenEnabled`, tfutils.StringFromBool(data.FrozenEnabled, ""))
	}
	if !data.UntrustedCodeMitigated.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UntrustedCodeMitigated`, tfutils.StringFromBool(data.UntrustedCodeMitigated, ""))
	}
	if !data.ReloadNeeded.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ReloadNeeded`, tfutils.StringFromBool(data.ReloadNeeded, ""))
	}
	if !data.TerminateTime.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TerminateTime`, data.TerminateTime.ValueInt64())
	}
	return body
}

func (data *GWScriptSettings) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `FrozenEnabled`); value.Exists() {
		data.FrozenEnabled = tfutils.BoolFromString(value.String())
	} else {
		data.FrozenEnabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UntrustedCodeMitigated`); value.Exists() {
		data.UntrustedCodeMitigated = tfutils.BoolFromString(value.String())
	} else {
		data.UntrustedCodeMitigated = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ReloadNeeded`); value.Exists() {
		data.ReloadNeeded = tfutils.BoolFromString(value.String())
	} else {
		data.ReloadNeeded = types.BoolNull()
	}
	if value := res.Get(pathRoot + `TerminateTime`); value.Exists() {
		data.TerminateTime = types.Int64Value(value.Int())
	} else {
		data.TerminateTime = types.Int64Null()
	}
}

func (data *GWScriptSettings) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `FrozenEnabled`); value.Exists() && !data.FrozenEnabled.IsNull() {
		data.FrozenEnabled = tfutils.BoolFromString(value.String())
	} else if !data.FrozenEnabled.ValueBool() {
		data.FrozenEnabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UntrustedCodeMitigated`); value.Exists() && !data.UntrustedCodeMitigated.IsNull() {
		data.UntrustedCodeMitigated = tfutils.BoolFromString(value.String())
	} else if !data.UntrustedCodeMitigated.ValueBool() {
		data.UntrustedCodeMitigated = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ReloadNeeded`); value.Exists() && !data.ReloadNeeded.IsNull() {
		data.ReloadNeeded = tfutils.BoolFromString(value.String())
	} else if data.ReloadNeeded.ValueBool() {
		data.ReloadNeeded = types.BoolNull()
	}
	if value := res.Get(pathRoot + `TerminateTime`); value.Exists() && !data.TerminateTime.IsNull() {
		data.TerminateTime = types.Int64Value(value.Int())
	} else {
		data.TerminateTime = types.Int64Null()
	}
}
func (data *GWScriptSettings) UpdateUnknownFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if data.Enabled.IsUnknown() {
		if value := res.Get(pathRoot + `mAdminState`); value.Exists() && !data.Enabled.IsNull() {
			data.Enabled = tfutils.BoolFromString(value.String())
		} else {
			data.Enabled = types.BoolNull()
		}
	}
	if data.UserSummary.IsUnknown() {
		if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
			data.UserSummary = tfutils.ParseStringFromGJSON(value)
		} else {
			data.UserSummary = types.StringNull()
		}
	}
	if data.FrozenEnabled.IsUnknown() {
		if value := res.Get(pathRoot + `FrozenEnabled`); value.Exists() && !data.FrozenEnabled.IsNull() {
			data.FrozenEnabled = tfutils.BoolFromString(value.String())
		} else {
			data.FrozenEnabled = types.BoolNull()
		}
	}
	if data.UntrustedCodeMitigated.IsUnknown() {
		if value := res.Get(pathRoot + `UntrustedCodeMitigated`); value.Exists() && !data.UntrustedCodeMitigated.IsNull() {
			data.UntrustedCodeMitigated = tfutils.BoolFromString(value.String())
		} else {
			data.UntrustedCodeMitigated = types.BoolNull()
		}
	}
	if data.ReloadNeeded.IsUnknown() {
		if value := res.Get(pathRoot + `ReloadNeeded`); value.Exists() && !data.ReloadNeeded.IsNull() {
			data.ReloadNeeded = tfutils.BoolFromString(value.String())
		} else {
			data.ReloadNeeded = types.BoolNull()
		}
	}
	if data.TerminateTime.IsUnknown() {
		if value := res.Get(pathRoot + `TerminateTime`); value.Exists() && !data.TerminateTime.IsNull() {
			data.TerminateTime = types.Int64Value(value.Int())
		} else {
			data.TerminateTime = types.Int64Null()
		}
	}
}
