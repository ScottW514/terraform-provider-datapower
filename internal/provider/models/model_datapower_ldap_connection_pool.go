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

type LDAPConnectionPool struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	IdleTimeout       types.Int64                 `tfsdk:"idle_timeout"`
	MaxPoolSize       types.Int64                 `tfsdk:"max_pool_size"`
	RejectOnPoolLimit types.Bool                  `tfsdk:"reject_on_pool_limit"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var LDAPConnectionPoolObjectType = map[string]attr.Type{
	"id":                   types.StringType,
	"app_domain":           types.StringType,
	"user_summary":         types.StringType,
	"idle_timeout":         types.Int64Type,
	"max_pool_size":        types.Int64Type,
	"reject_on_pool_limit": types.BoolType,
	"dependency_actions":   actions.ActionsListType,
}

func (data LDAPConnectionPool) GetPath() string {
	rest_path := "/mgmt/config/{domain}/LDAPConnectionPool"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data LDAPConnectionPool) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.IdleTimeout.IsNull() {
		return false
	}
	if !data.MaxPoolSize.IsNull() {
		return false
	}
	if !data.RejectOnPoolLimit.IsNull() {
		return false
	}
	return true
}

func (data LDAPConnectionPool) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.IdleTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IdleTimeout`, data.IdleTimeout.ValueInt64())
	}
	if !data.MaxPoolSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxPoolSize`, data.MaxPoolSize.ValueInt64())
	}
	if !data.RejectOnPoolLimit.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RejectOnPoolLimit`, tfutils.StringFromBool(data.RejectOnPoolLimit, ""))
	}
	return body
}

func (data *LDAPConnectionPool) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `IdleTimeout`); value.Exists() {
		data.IdleTimeout = types.Int64Value(value.Int())
	} else {
		data.IdleTimeout = types.Int64Value(120)
	}
	if value := res.Get(pathRoot + `MaxPoolSize`); value.Exists() {
		data.MaxPoolSize = types.Int64Value(value.Int())
	} else {
		data.MaxPoolSize = types.Int64Value(35)
	}
	if value := res.Get(pathRoot + `RejectOnPoolLimit`); value.Exists() {
		data.RejectOnPoolLimit = tfutils.BoolFromString(value.String())
	} else {
		data.RejectOnPoolLimit = types.BoolNull()
	}
}

func (data *LDAPConnectionPool) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `IdleTimeout`); value.Exists() && !data.IdleTimeout.IsNull() {
		data.IdleTimeout = types.Int64Value(value.Int())
	} else if data.IdleTimeout.ValueInt64() != 120 {
		data.IdleTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxPoolSize`); value.Exists() && !data.MaxPoolSize.IsNull() {
		data.MaxPoolSize = types.Int64Value(value.Int())
	} else if data.MaxPoolSize.ValueInt64() != 35 {
		data.MaxPoolSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `RejectOnPoolLimit`); value.Exists() && !data.RejectOnPoolLimit.IsNull() {
		data.RejectOnPoolLimit = tfutils.BoolFromString(value.String())
	} else if data.RejectOnPoolLimit.ValueBool() {
		data.RejectOnPoolLimit = types.BoolNull()
	}
}
