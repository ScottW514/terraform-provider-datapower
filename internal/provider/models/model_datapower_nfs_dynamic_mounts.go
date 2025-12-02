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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type NFSDynamicMounts struct {
	AppDomain           types.String                `tfsdk:"app_domain"`
	Enabled             types.Bool                  `tfsdk:"enabled"`
	UserSummary         types.String                `tfsdk:"user_summary"`
	Version             types.Int64                 `tfsdk:"version"`
	Transport           types.String                `tfsdk:"transport"`
	MountType           types.String                `tfsdk:"mount_type"`
	ReadOnly            types.Bool                  `tfsdk:"read_only"`
	ReadSize            types.Int64                 `tfsdk:"read_size"`
	WriteSize           types.Int64                 `tfsdk:"write_size"`
	Timeout             types.Int64                 `tfsdk:"timeout"`
	Retransmissions     types.Int64                 `tfsdk:"retransmissions"`
	IdleUnmountSeconds  types.Int64                 `tfsdk:"idle_unmount_seconds"`
	MountTimeoutSeconds types.Int64                 `tfsdk:"mount_timeout_seconds"`
	DependencyActions   []*actions.DependencyAction `tfsdk:"dependency_actions"`
	ProviderTarget      types.String                `tfsdk:"provider_target"`
}

var NFSDynamicMountsTransportIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "version",
	AttrType:    "Int64",
	AttrDefault: "3",
	Value:       []string{"4"},
}

var NFSDynamicMountsObjectType = map[string]attr.Type{
	"provider_target":       types.StringType,
	"app_domain":            types.StringType,
	"enabled":               types.BoolType,
	"user_summary":          types.StringType,
	"version":               types.Int64Type,
	"transport":             types.StringType,
	"mount_type":            types.StringType,
	"read_only":             types.BoolType,
	"read_size":             types.Int64Type,
	"write_size":            types.Int64Type,
	"timeout":               types.Int64Type,
	"retransmissions":       types.Int64Type,
	"idle_unmount_seconds":  types.Int64Type,
	"mount_timeout_seconds": types.Int64Type,
	"dependency_actions":    actions.ActionsListType,
}

func (data NFSDynamicMounts) GetPath() string {
	rest_path := "/mgmt/config/{domain}/NFSDynamicMounts/default"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data NFSDynamicMounts) IsNull() bool {
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Version.IsNull() {
		return false
	}
	if !data.Transport.IsNull() {
		return false
	}
	if !data.MountType.IsNull() {
		return false
	}
	if !data.ReadOnly.IsNull() {
		return false
	}
	if !data.ReadSize.IsNull() {
		return false
	}
	if !data.WriteSize.IsNull() {
		return false
	}
	if !data.Timeout.IsNull() {
		return false
	}
	if !data.Retransmissions.IsNull() {
		return false
	}
	if !data.IdleUnmountSeconds.IsNull() {
		return false
	}
	if !data.MountTimeoutSeconds.IsNull() {
		return false
	}
	return true
}
func (data *NFSDynamicMounts) ToDefault() {
	data.Enabled = types.BoolValue(false)
	data.UserSummary = types.StringNull()
	data.Version = types.Int64Value(3)
	data.Transport = types.StringValue("tcp")
	data.MountType = types.StringValue("hard")
	data.ReadOnly = types.BoolValue(false)
	data.ReadSize = types.Int64Value(4096)
	data.WriteSize = types.Int64Value(4096)
	data.Timeout = types.Int64Value(7)
	data.Retransmissions = types.Int64Value(3)
	data.IdleUnmountSeconds = types.Int64Value(900)
	data.MountTimeoutSeconds = types.Int64Value(30)
}

func (data NFSDynamicMounts) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "NFSDynamicMounts.name", path.Base("/mgmt/config/{domain}/NFSDynamicMounts/default"))

	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.Version.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Version`, data.Version.ValueInt64())
	}
	if !data.Transport.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Transport`, data.Transport.ValueString())
	}
	if !data.MountType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MountType`, data.MountType.ValueString())
	}
	if !data.ReadOnly.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ReadOnly`, tfutils.StringFromBool(data.ReadOnly, ""))
	}
	if !data.ReadSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ReadSize`, data.ReadSize.ValueInt64())
	}
	if !data.WriteSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WriteSize`, data.WriteSize.ValueInt64())
	}
	if !data.Timeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Timeout`, data.Timeout.ValueInt64())
	}
	if !data.Retransmissions.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Retransmissions`, data.Retransmissions.ValueInt64())
	}
	if !data.IdleUnmountSeconds.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IdleUnmountSeconds`, data.IdleUnmountSeconds.ValueInt64())
	}
	if !data.MountTimeoutSeconds.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MountTimeoutSeconds`, data.MountTimeoutSeconds.ValueInt64())
	}
	return body
}

func (data *NFSDynamicMounts) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Version`); value.Exists() {
		data.Version = types.Int64Value(value.Int())
	} else {
		data.Version = types.Int64Value(3)
	}
	if value := res.Get(pathRoot + `Transport`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Transport = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Transport = types.StringValue("tcp")
	}
	if value := res.Get(pathRoot + `MountType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MountType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MountType = types.StringValue("hard")
	}
	if value := res.Get(pathRoot + `ReadOnly`); value.Exists() {
		data.ReadOnly = tfutils.BoolFromString(value.String())
	} else {
		data.ReadOnly = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ReadSize`); value.Exists() {
		data.ReadSize = types.Int64Value(value.Int())
	} else {
		data.ReadSize = types.Int64Value(4096)
	}
	if value := res.Get(pathRoot + `WriteSize`); value.Exists() {
		data.WriteSize = types.Int64Value(value.Int())
	} else {
		data.WriteSize = types.Int64Value(4096)
	}
	if value := res.Get(pathRoot + `Timeout`); value.Exists() {
		data.Timeout = types.Int64Value(value.Int())
	} else {
		data.Timeout = types.Int64Value(7)
	}
	if value := res.Get(pathRoot + `Retransmissions`); value.Exists() {
		data.Retransmissions = types.Int64Value(value.Int())
	} else {
		data.Retransmissions = types.Int64Value(3)
	}
	if value := res.Get(pathRoot + `IdleUnmountSeconds`); value.Exists() {
		data.IdleUnmountSeconds = types.Int64Value(value.Int())
	} else {
		data.IdleUnmountSeconds = types.Int64Value(900)
	}
	if value := res.Get(pathRoot + `MountTimeoutSeconds`); value.Exists() {
		data.MountTimeoutSeconds = types.Int64Value(value.Int())
	} else {
		data.MountTimeoutSeconds = types.Int64Value(30)
	}
}

func (data *NFSDynamicMounts) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Version`); value.Exists() && !data.Version.IsNull() {
		data.Version = types.Int64Value(value.Int())
	} else if data.Version.ValueInt64() != 3 {
		data.Version = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Transport`); value.Exists() && !data.Transport.IsNull() {
		data.Transport = tfutils.ParseStringFromGJSON(value)
	} else if data.Transport.ValueString() != "tcp" {
		data.Transport = types.StringNull()
	}
	if value := res.Get(pathRoot + `MountType`); value.Exists() && !data.MountType.IsNull() {
		data.MountType = tfutils.ParseStringFromGJSON(value)
	} else if data.MountType.ValueString() != "hard" {
		data.MountType = types.StringNull()
	}
	if value := res.Get(pathRoot + `ReadOnly`); value.Exists() && !data.ReadOnly.IsNull() {
		data.ReadOnly = tfutils.BoolFromString(value.String())
	} else if data.ReadOnly.ValueBool() {
		data.ReadOnly = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ReadSize`); value.Exists() && !data.ReadSize.IsNull() {
		data.ReadSize = types.Int64Value(value.Int())
	} else if data.ReadSize.ValueInt64() != 4096 {
		data.ReadSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `WriteSize`); value.Exists() && !data.WriteSize.IsNull() {
		data.WriteSize = types.Int64Value(value.Int())
	} else if data.WriteSize.ValueInt64() != 4096 {
		data.WriteSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Timeout`); value.Exists() && !data.Timeout.IsNull() {
		data.Timeout = types.Int64Value(value.Int())
	} else if data.Timeout.ValueInt64() != 7 {
		data.Timeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Retransmissions`); value.Exists() && !data.Retransmissions.IsNull() {
		data.Retransmissions = types.Int64Value(value.Int())
	} else if data.Retransmissions.ValueInt64() != 3 {
		data.Retransmissions = types.Int64Null()
	}
	if value := res.Get(pathRoot + `IdleUnmountSeconds`); value.Exists() && !data.IdleUnmountSeconds.IsNull() {
		data.IdleUnmountSeconds = types.Int64Value(value.Int())
	} else if data.IdleUnmountSeconds.ValueInt64() != 900 {
		data.IdleUnmountSeconds = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MountTimeoutSeconds`); value.Exists() && !data.MountTimeoutSeconds.IsNull() {
		data.MountTimeoutSeconds = types.Int64Value(value.Int())
	} else if data.MountTimeoutSeconds.ValueInt64() != 30 {
		data.MountTimeoutSeconds = types.Int64Null()
	}
}
