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
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type WebAppSessionPolicy struct {
	Id                    types.String                `tfsdk:"id"`
	AppDomain             types.String                `tfsdk:"app_domain"`
	UserSummary           types.String                `tfsdk:"user_summary"`
	AutoRenew             types.Bool                  `tfsdk:"auto_renew"`
	Timeout               types.Int64                 `tfsdk:"timeout"`
	AddressAgnosticCookie types.Bool                  `tfsdk:"address_agnostic_cookie"`
	StartMatches          types.String                `tfsdk:"start_matches"`
	DependencyActions     []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var WebAppSessionPolicyObjectType = map[string]attr.Type{
	"id":                      types.StringType,
	"app_domain":              types.StringType,
	"user_summary":            types.StringType,
	"auto_renew":              types.BoolType,
	"timeout":                 types.Int64Type,
	"address_agnostic_cookie": types.BoolType,
	"start_matches":           types.StringType,
	"dependency_actions":      actions.ActionsListType,
}

func (data WebAppSessionPolicy) GetPath() string {
	rest_path := "/mgmt/config/{domain}/WebAppSessionPolicy"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data WebAppSessionPolicy) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.AutoRenew.IsNull() {
		return false
	}
	if !data.Timeout.IsNull() {
		return false
	}
	if !data.AddressAgnosticCookie.IsNull() {
		return false
	}
	if !data.StartMatches.IsNull() {
		return false
	}
	return true
}

func (data WebAppSessionPolicy) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.AutoRenew.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AutoRenew`, tfutils.StringFromBool(data.AutoRenew, ""))
	}
	if !data.Timeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Timeout`, data.Timeout.ValueInt64())
	}
	if !data.AddressAgnosticCookie.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AddressAgnosticCookie`, tfutils.StringFromBool(data.AddressAgnosticCookie, ""))
	}
	if !data.StartMatches.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StartMatches`, data.StartMatches.ValueString())
	}
	return body
}

func (data *WebAppSessionPolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `AutoRenew`); value.Exists() {
		data.AutoRenew = tfutils.BoolFromString(value.String())
	} else {
		data.AutoRenew = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Timeout`); value.Exists() {
		data.Timeout = types.Int64Value(value.Int())
	} else {
		data.Timeout = types.Int64Value(3600)
	}
	if value := res.Get(pathRoot + `AddressAgnosticCookie`); value.Exists() {
		data.AddressAgnosticCookie = tfutils.BoolFromString(value.String())
	} else {
		data.AddressAgnosticCookie = types.BoolNull()
	}
	if value := res.Get(pathRoot + `StartMatches`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.StartMatches = tfutils.ParseStringFromGJSON(value)
	} else {
		data.StartMatches = types.StringNull()
	}
}

func (data *WebAppSessionPolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `AutoRenew`); value.Exists() && !data.AutoRenew.IsNull() {
		data.AutoRenew = tfutils.BoolFromString(value.String())
	} else if !data.AutoRenew.ValueBool() {
		data.AutoRenew = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Timeout`); value.Exists() && !data.Timeout.IsNull() {
		data.Timeout = types.Int64Value(value.Int())
	} else if data.Timeout.ValueInt64() != 3600 {
		data.Timeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AddressAgnosticCookie`); value.Exists() && !data.AddressAgnosticCookie.IsNull() {
		data.AddressAgnosticCookie = tfutils.BoolFromString(value.String())
	} else if data.AddressAgnosticCookie.ValueBool() {
		data.AddressAgnosticCookie = types.BoolNull()
	}
	if value := res.Get(pathRoot + `StartMatches`); value.Exists() && !data.StartMatches.IsNull() {
		data.StartMatches = tfutils.ParseStringFromGJSON(value)
	} else {
		data.StartMatches = types.StringNull()
	}
}
