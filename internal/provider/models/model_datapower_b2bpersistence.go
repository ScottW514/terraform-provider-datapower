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

type B2BPersistence struct {
	Enabled           types.Bool        `tfsdk:"enabled"`
	UserSummary       types.String      `tfsdk:"user_summary"`
	RaidVolume        types.String      `tfsdk:"raid_volume"`
	StorageSize       types.Int64       `tfsdk:"storage_size"`
	HaEnabled         types.Bool        `tfsdk:"ha_enabled"`
	HaOtherHosts      *DmB2BHAHost      `tfsdk:"ha_other_hosts"`
	HaLocalIp         types.String      `tfsdk:"ha_local_ip"`
	HaLocalPort       types.Int64       `tfsdk:"ha_local_port"`
	HaVirtualIp       types.String      `tfsdk:"ha_virtual_ip"`
	DependencyActions []*actions.Action `tfsdk:"dependency_actions"`
}

var B2BPersistenceObjectType = map[string]attr.Type{
	"enabled":            types.BoolType,
	"user_summary":       types.StringType,
	"raid_volume":        types.StringType,
	"storage_size":       types.Int64Type,
	"ha_enabled":         types.BoolType,
	"ha_other_hosts":     types.ObjectType{AttrTypes: DmB2BHAHostObjectType},
	"ha_local_ip":        types.StringType,
	"ha_local_port":      types.Int64Type,
	"ha_virtual_ip":      types.StringType,
	"dependency_actions": actions.ActionsListType,
}

func (data B2BPersistence) GetPath() string {
	rest_path := "/mgmt/config/default/B2BPersistence/B2BPersistence"
	return rest_path
}

func (data B2BPersistence) IsNull() bool {
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.RaidVolume.IsNull() {
		return false
	}
	if !data.StorageSize.IsNull() {
		return false
	}
	if !data.HaEnabled.IsNull() {
		return false
	}
	if data.HaOtherHosts != nil {
		if !data.HaOtherHosts.IsNull() {
			return false
		}
	}
	if !data.HaLocalIp.IsNull() {
		return false
	}
	if !data.HaLocalPort.IsNull() {
		return false
	}
	if !data.HaVirtualIp.IsNull() {
		return false
	}
	return true
}

func (data B2BPersistence) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "B2BPersistence.name", path.Base("/mgmt/config/default/B2BPersistence/B2BPersistence"))
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.RaidVolume.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RaidVolume`, data.RaidVolume.ValueString())
	}
	if !data.StorageSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StorageSize`, data.StorageSize.ValueInt64())
	}
	if !data.HaEnabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HAEnabled`, tfutils.StringFromBool(data.HaEnabled, ""))
	}
	if data.HaOtherHosts != nil {
		if !data.HaOtherHosts.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`HAOtherHosts`, data.HaOtherHosts.ToBody(ctx, ""))
		}
	}
	if !data.HaLocalIp.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HALocalIP`, data.HaLocalIp.ValueString())
	}
	if !data.HaLocalPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HALocalPort`, data.HaLocalPort.ValueInt64())
	}
	if !data.HaVirtualIp.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HAVirtualIP`, data.HaVirtualIp.ValueString())
	}
	return body
}

func (data *B2BPersistence) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `RaidVolume`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RaidVolume = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RaidVolume = types.StringNull()
	}
	if value := res.Get(pathRoot + `StorageSize`); value.Exists() {
		data.StorageSize = types.Int64Value(value.Int())
	} else {
		data.StorageSize = types.Int64Value(1024)
	}
	if value := res.Get(pathRoot + `HAEnabled`); value.Exists() {
		data.HaEnabled = tfutils.BoolFromString(value.String())
	} else {
		data.HaEnabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HAOtherHosts`); value.Exists() {
		data.HaOtherHosts = &DmB2BHAHost{}
		data.HaOtherHosts.FromBody(ctx, "", value)
	} else {
		data.HaOtherHosts = nil
	}
	if value := res.Get(pathRoot + `HALocalIP`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HaLocalIp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HaLocalIp = types.StringNull()
	}
	if value := res.Get(pathRoot + `HALocalPort`); value.Exists() {
		data.HaLocalPort = types.Int64Value(value.Int())
	} else {
		data.HaLocalPort = types.Int64Value(1320)
	}
	if value := res.Get(pathRoot + `HAVirtualIP`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HaVirtualIp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HaVirtualIp = types.StringNull()
	}
}

func (data *B2BPersistence) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `RaidVolume`); value.Exists() && !data.RaidVolume.IsNull() {
		data.RaidVolume = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RaidVolume = types.StringNull()
	}
	if value := res.Get(pathRoot + `StorageSize`); value.Exists() && !data.StorageSize.IsNull() {
		data.StorageSize = types.Int64Value(value.Int())
	} else if data.StorageSize.ValueInt64() != 1024 {
		data.StorageSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `HAEnabled`); value.Exists() && !data.HaEnabled.IsNull() {
		data.HaEnabled = tfutils.BoolFromString(value.String())
	} else if data.HaEnabled.ValueBool() {
		data.HaEnabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HAOtherHosts`); value.Exists() {
		data.HaOtherHosts.UpdateFromBody(ctx, "", value)
	} else {
		data.HaOtherHosts = nil
	}
	if value := res.Get(pathRoot + `HALocalIP`); value.Exists() && !data.HaLocalIp.IsNull() {
		data.HaLocalIp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HaLocalIp = types.StringNull()
	}
	if value := res.Get(pathRoot + `HALocalPort`); value.Exists() && !data.HaLocalPort.IsNull() {
		data.HaLocalPort = types.Int64Value(value.Int())
	} else if data.HaLocalPort.ValueInt64() != 1320 {
		data.HaLocalPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `HAVirtualIP`); value.Exists() && !data.HaVirtualIp.IsNull() {
		data.HaVirtualIp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HaVirtualIp = types.StringNull()
	}
}
