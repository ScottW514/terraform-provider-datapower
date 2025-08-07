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

type GatewayPeeringManager struct {
	AppDomain                types.String      `tfsdk:"app_domain"`
	Enabled                  types.Bool        `tfsdk:"enabled"`
	UserSummary              types.String      `tfsdk:"user_summary"`
	ApiConnectGatewayService types.String      `tfsdk:"api_connect_gateway_service"`
	RateLimit                types.String      `tfsdk:"rate_limit"`
	Subscription             types.String      `tfsdk:"subscription"`
	RatelimitModule          types.String      `tfsdk:"ratelimit_module"`
	DependencyActions        []*actions.Action `tfsdk:"dependency_actions"`
}

var GatewayPeeringManagerObjectType = map[string]attr.Type{
	"app_domain":                  types.StringType,
	"enabled":                     types.BoolType,
	"user_summary":                types.StringType,
	"api_connect_gateway_service": types.StringType,
	"rate_limit":                  types.StringType,
	"subscription":                types.StringType,
	"ratelimit_module":            types.StringType,
	"dependency_actions":          actions.ActionsListType,
}

func (data GatewayPeeringManager) GetPath() string {
	rest_path := "/mgmt/config/{domain}/GatewayPeeringManager/default"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data GatewayPeeringManager) IsNull() bool {
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.ApiConnectGatewayService.IsNull() {
		return false
	}
	if !data.RateLimit.IsNull() {
		return false
	}
	if !data.Subscription.IsNull() {
		return false
	}
	if !data.RatelimitModule.IsNull() {
		return false
	}
	return true
}

func (data GatewayPeeringManager) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "GatewayPeeringManager.name", path.Base("/mgmt/config/{domain}/GatewayPeeringManager/default"))
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.ApiConnectGatewayService.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APIConnectGatewayService`, data.ApiConnectGatewayService.ValueString())
	}
	if !data.RateLimit.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RateLimit`, data.RateLimit.ValueString())
	}
	if !data.Subscription.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Subscription`, data.Subscription.ValueString())
	}
	if !data.RatelimitModule.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RatelimitModule`, data.RatelimitModule.ValueString())
	}
	return body
}

func (data *GatewayPeeringManager) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `APIConnectGatewayService`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ApiConnectGatewayService = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ApiConnectGatewayService = types.StringNull()
	}
	if value := res.Get(pathRoot + `RateLimit`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RateLimit = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RateLimit = types.StringNull()
	}
	if value := res.Get(pathRoot + `Subscription`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Subscription = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Subscription = types.StringNull()
	}
	if value := res.Get(pathRoot + `RatelimitModule`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RatelimitModule = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RatelimitModule = types.StringNull()
	}
}

func (data *GatewayPeeringManager) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `APIConnectGatewayService`); value.Exists() && !data.ApiConnectGatewayService.IsNull() {
		data.ApiConnectGatewayService = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ApiConnectGatewayService = types.StringNull()
	}
	if value := res.Get(pathRoot + `RateLimit`); value.Exists() && !data.RateLimit.IsNull() {
		data.RateLimit = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RateLimit = types.StringNull()
	}
	if value := res.Get(pathRoot + `Subscription`); value.Exists() && !data.Subscription.IsNull() {
		data.Subscription = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Subscription = types.StringNull()
	}
	if value := res.Get(pathRoot + `RatelimitModule`); value.Exists() && !data.RatelimitModule.IsNull() {
		data.RatelimitModule = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RatelimitModule = types.StringNull()
	}
}
