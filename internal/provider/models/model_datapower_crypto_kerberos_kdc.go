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

type CryptoKerberosKDC struct {
	Id                      types.String                `tfsdk:"id"`
	AppDomain               types.String                `tfsdk:"app_domain"`
	UserSummary             types.String                `tfsdk:"user_summary"`
	Realm                   types.String                `tfsdk:"realm"`
	Server                  types.String                `tfsdk:"server"`
	UseTcp                  types.Bool                  `tfsdk:"use_tcp"`
	ServerPort              types.Int64                 `tfsdk:"server_port"`
	UdpTimeout              types.Int64                 `tfsdk:"udp_timeout"`
	CacheTickets            types.Bool                  `tfsdk:"cache_tickets"`
	MaxCachedTickets        types.Int64                 `tfsdk:"max_cached_tickets"`
	MinCachedTicketValidity types.Int64                 `tfsdk:"min_cached_ticket_validity"`
	DependencyActions       []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var CryptoKerberosKDCObjectType = map[string]attr.Type{
	"id":                         types.StringType,
	"app_domain":                 types.StringType,
	"user_summary":               types.StringType,
	"realm":                      types.StringType,
	"server":                     types.StringType,
	"use_tcp":                    types.BoolType,
	"server_port":                types.Int64Type,
	"udp_timeout":                types.Int64Type,
	"cache_tickets":              types.BoolType,
	"max_cached_tickets":         types.Int64Type,
	"min_cached_ticket_validity": types.Int64Type,
	"dependency_actions":         actions.ActionsListType,
}

func (data CryptoKerberosKDC) GetPath() string {
	rest_path := "/mgmt/config/{domain}/CryptoKerberosKDC"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data CryptoKerberosKDC) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Realm.IsNull() {
		return false
	}
	if !data.Server.IsNull() {
		return false
	}
	if !data.UseTcp.IsNull() {
		return false
	}
	if !data.ServerPort.IsNull() {
		return false
	}
	if !data.UdpTimeout.IsNull() {
		return false
	}
	if !data.CacheTickets.IsNull() {
		return false
	}
	if !data.MaxCachedTickets.IsNull() {
		return false
	}
	if !data.MinCachedTicketValidity.IsNull() {
		return false
	}
	return true
}

func (data CryptoKerberosKDC) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Realm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Realm`, data.Realm.ValueString())
	}
	if !data.Server.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Server`, data.Server.ValueString())
	}
	if !data.UseTcp.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseTCP`, tfutils.StringFromBool(data.UseTcp, ""))
	}
	if !data.ServerPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ServerPort`, data.ServerPort.ValueInt64())
	}
	if !data.UdpTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UDPTimeout`, data.UdpTimeout.ValueInt64())
	}
	if !data.CacheTickets.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CacheTickets`, tfutils.StringFromBool(data.CacheTickets, ""))
	}
	if !data.MaxCachedTickets.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxCachedTickets`, data.MaxCachedTickets.ValueInt64())
	}
	if !data.MinCachedTicketValidity.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MinCachedTicketValidity`, data.MinCachedTicketValidity.ValueInt64())
	}
	return body
}

func (data *CryptoKerberosKDC) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Realm`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Realm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Realm = types.StringNull()
	}
	if value := res.Get(pathRoot + `Server`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Server = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Server = types.StringNull()
	}
	if value := res.Get(pathRoot + `UseTCP`); value.Exists() {
		data.UseTcp = tfutils.BoolFromString(value.String())
	} else {
		data.UseTcp = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ServerPort`); value.Exists() {
		data.ServerPort = types.Int64Value(value.Int())
	} else {
		data.ServerPort = types.Int64Value(88)
	}
	if value := res.Get(pathRoot + `UDPTimeout`); value.Exists() {
		data.UdpTimeout = types.Int64Value(value.Int())
	} else {
		data.UdpTimeout = types.Int64Value(5)
	}
	if value := res.Get(pathRoot + `CacheTickets`); value.Exists() {
		data.CacheTickets = tfutils.BoolFromString(value.String())
	} else {
		data.CacheTickets = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MaxCachedTickets`); value.Exists() {
		data.MaxCachedTickets = types.Int64Value(value.Int())
	} else {
		data.MaxCachedTickets = types.Int64Value(32)
	}
	if value := res.Get(pathRoot + `MinCachedTicketValidity`); value.Exists() {
		data.MinCachedTicketValidity = types.Int64Value(value.Int())
	} else {
		data.MinCachedTicketValidity = types.Int64Value(60)
	}
}

func (data *CryptoKerberosKDC) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Realm`); value.Exists() && !data.Realm.IsNull() {
		data.Realm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Realm = types.StringNull()
	}
	if value := res.Get(pathRoot + `Server`); value.Exists() && !data.Server.IsNull() {
		data.Server = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Server = types.StringNull()
	}
	if value := res.Get(pathRoot + `UseTCP`); value.Exists() && !data.UseTcp.IsNull() {
		data.UseTcp = tfutils.BoolFromString(value.String())
	} else if data.UseTcp.ValueBool() {
		data.UseTcp = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ServerPort`); value.Exists() && !data.ServerPort.IsNull() {
		data.ServerPort = types.Int64Value(value.Int())
	} else if data.ServerPort.ValueInt64() != 88 {
		data.ServerPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `UDPTimeout`); value.Exists() && !data.UdpTimeout.IsNull() {
		data.UdpTimeout = types.Int64Value(value.Int())
	} else if data.UdpTimeout.ValueInt64() != 5 {
		data.UdpTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `CacheTickets`); value.Exists() && !data.CacheTickets.IsNull() {
		data.CacheTickets = tfutils.BoolFromString(value.String())
	} else if !data.CacheTickets.ValueBool() {
		data.CacheTickets = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MaxCachedTickets`); value.Exists() && !data.MaxCachedTickets.IsNull() {
		data.MaxCachedTickets = types.Int64Value(value.Int())
	} else if data.MaxCachedTickets.ValueInt64() != 32 {
		data.MaxCachedTickets = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MinCachedTicketValidity`); value.Exists() && !data.MinCachedTicketValidity.IsNull() {
		data.MinCachedTicketValidity = types.Int64Value(value.Int())
	} else if data.MinCachedTicketValidity.ValueInt64() != 60 {
		data.MinCachedTicketValidity = types.Int64Null()
	}
}
