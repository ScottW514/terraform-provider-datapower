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

type Probe struct {
	AppDomain      types.String      `tfsdk:"app_domain"`
	Enabled        types.Bool        `tfsdk:"enabled"`
	UserSummary    types.String      `tfsdk:"user_summary"`
	MaxRecords     types.Int64       `tfsdk:"max_records"`
	Expiration     types.Int64       `tfsdk:"expiration"`
	GatewayPeering types.String      `tfsdk:"gateway_peering"`
	ObjectActions  []*actions.Action `tfsdk:"object_actions"`
}

var ProbeObjectType = map[string]attr.Type{
	"app_domain":      types.StringType,
	"enabled":         types.BoolType,
	"user_summary":    types.StringType,
	"max_records":     types.Int64Type,
	"expiration":      types.Int64Type,
	"gateway_peering": types.StringType,
	"object_actions":  actions.ActionsListType,
}

func (data Probe) GetPath() string {
	rest_path := "/mgmt/config/{domain}/Probe/default"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data Probe) IsNull() bool {
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
	if !data.Expiration.IsNull() {
		return false
	}
	if !data.GatewayPeering.IsNull() {
		return false
	}
	return true
}

func (data Probe) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "Probe.name", path.Base("/mgmt/config/{domain}/Probe/default"))
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.MaxRecords.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxRecords`, data.MaxRecords.ValueInt64())
	}
	if !data.Expiration.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Expiration`, data.Expiration.ValueInt64())
	}
	if !data.GatewayPeering.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GatewayPeering`, data.GatewayPeering.ValueString())
	}
	return body
}

func (data *Probe) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
		data.MaxRecords = types.Int64Value(1000)
	}
	if value := res.Get(pathRoot + `Expiration`); value.Exists() {
		data.Expiration = types.Int64Value(value.Int())
	} else {
		data.Expiration = types.Int64Value(3600)
	}
	if value := res.Get(pathRoot + `GatewayPeering`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.GatewayPeering = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GatewayPeering = types.StringNull()
	}
}

func (data *Probe) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	} else if data.MaxRecords.ValueInt64() != 1000 {
		data.MaxRecords = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Expiration`); value.Exists() && !data.Expiration.IsNull() {
		data.Expiration = types.Int64Value(value.Int())
	} else if data.Expiration.ValueInt64() != 3600 {
		data.Expiration = types.Int64Null()
	}
	if value := res.Get(pathRoot + `GatewayPeering`); value.Exists() && !data.GatewayPeering.IsNull() {
		data.GatewayPeering = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GatewayPeering = types.StringNull()
	}
}
