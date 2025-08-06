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

type APISecurityTokenManager struct {
	AppDomain                   types.String      `tfsdk:"app_domain"`
	Enabled                     types.Bool        `tfsdk:"enabled"`
	UserSummary                 types.String      `tfsdk:"user_summary"`
	GatewayPeering              types.String      `tfsdk:"gateway_peering"`
	GatewayPeeringExternal      types.String      `tfsdk:"gateway_peering_external"`
	ExpiredTokenCleanupInterval types.Int64       `tfsdk:"expired_token_cleanup_interval"`
	ObjectActions               []*actions.Action `tfsdk:"object_actions"`
}

var APISecurityTokenManagerObjectType = map[string]attr.Type{
	"app_domain":                     types.StringType,
	"enabled":                        types.BoolType,
	"user_summary":                   types.StringType,
	"gateway_peering":                types.StringType,
	"gateway_peering_external":       types.StringType,
	"expired_token_cleanup_interval": types.Int64Type,
	"object_actions":                 actions.ActionsListType,
}

func (data APISecurityTokenManager) GetPath() string {
	rest_path := "/mgmt/config/{domain}/APISecurityTokenManager/default"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data APISecurityTokenManager) IsNull() bool {
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.GatewayPeering.IsNull() {
		return false
	}
	if !data.GatewayPeeringExternal.IsNull() {
		return false
	}
	if !data.ExpiredTokenCleanupInterval.IsNull() {
		return false
	}
	return true
}

func (data APISecurityTokenManager) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "APISecurityTokenManager.name", path.Base("/mgmt/config/{domain}/APISecurityTokenManager/default"))
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.GatewayPeering.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GatewayPeering`, data.GatewayPeering.ValueString())
	}
	if !data.GatewayPeeringExternal.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GatewayPeeringExternal`, data.GatewayPeeringExternal.ValueString())
	}
	if !data.ExpiredTokenCleanupInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ExpiredTokenCleanupInterval`, data.ExpiredTokenCleanupInterval.ValueInt64())
	}
	return body
}

func (data *APISecurityTokenManager) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `GatewayPeering`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.GatewayPeering = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GatewayPeering = types.StringNull()
	}
	if value := res.Get(pathRoot + `GatewayPeeringExternal`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.GatewayPeeringExternal = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GatewayPeeringExternal = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExpiredTokenCleanupInterval`); value.Exists() {
		data.ExpiredTokenCleanupInterval = types.Int64Value(value.Int())
	} else {
		data.ExpiredTokenCleanupInterval = types.Int64Value(180)
	}
}

func (data *APISecurityTokenManager) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `GatewayPeering`); value.Exists() && !data.GatewayPeering.IsNull() {
		data.GatewayPeering = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GatewayPeering = types.StringNull()
	}
	if value := res.Get(pathRoot + `GatewayPeeringExternal`); value.Exists() && !data.GatewayPeeringExternal.IsNull() {
		data.GatewayPeeringExternal = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GatewayPeeringExternal = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExpiredTokenCleanupInterval`); value.Exists() && !data.ExpiredTokenCleanupInterval.IsNull() {
		data.ExpiredTokenCleanupInterval = types.Int64Value(value.Int())
	} else if data.ExpiredTokenCleanupInterval.ValueInt64() != 180 {
		data.ExpiredTokenCleanupInterval = types.Int64Null()
	}
}
