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

type MQv9PlusMFTSourceProtocolHandler struct {
	Id                      types.String                `tfsdk:"id"`
	AppDomain               types.String                `tfsdk:"app_domain"`
	UserSummary             types.String                `tfsdk:"user_summary"`
	QueueManager            types.String                `tfsdk:"queue_manager"`
	GetQueue                types.String                `tfsdk:"get_queue"`
	GetMessageOptions       types.Int64                 `tfsdk:"get_message_options"`
	ConcurrentConnections   types.Int64                 `tfsdk:"concurrent_connections"`
	PollingInterval         types.Int64                 `tfsdk:"polling_interval"`
	RetrieveBackoutSettings types.Bool                  `tfsdk:"retrieve_backout_settings"`
	IgnoreBackoutErrors     types.Bool                  `tfsdk:"ignore_backout_errors"`
	UseQmNameInUrl          types.Bool                  `tfsdk:"use_qm_name_in_url"`
	DependencyActions       []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var MQv9PlusMFTSourceProtocolHandlerObjectType = map[string]attr.Type{
	"id":                        types.StringType,
	"app_domain":                types.StringType,
	"user_summary":              types.StringType,
	"queue_manager":             types.StringType,
	"get_queue":                 types.StringType,
	"get_message_options":       types.Int64Type,
	"concurrent_connections":    types.Int64Type,
	"polling_interval":          types.Int64Type,
	"retrieve_backout_settings": types.BoolType,
	"ignore_backout_errors":     types.BoolType,
	"use_qm_name_in_url":        types.BoolType,
	"dependency_actions":        actions.ActionsListType,
}

func (data MQv9PlusMFTSourceProtocolHandler) GetPath() string {
	rest_path := "/mgmt/config/{domain}/MQv9PlusMFTSourceProtocolHandler"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data MQv9PlusMFTSourceProtocolHandler) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.QueueManager.IsNull() {
		return false
	}
	if !data.GetQueue.IsNull() {
		return false
	}
	if !data.GetMessageOptions.IsNull() {
		return false
	}
	if !data.ConcurrentConnections.IsNull() {
		return false
	}
	if !data.PollingInterval.IsNull() {
		return false
	}
	if !data.RetrieveBackoutSettings.IsNull() {
		return false
	}
	if !data.IgnoreBackoutErrors.IsNull() {
		return false
	}
	if !data.UseQmNameInUrl.IsNull() {
		return false
	}
	return true
}

func (data MQv9PlusMFTSourceProtocolHandler) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.QueueManager.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`QueueManager`, data.QueueManager.ValueString())
	}
	if !data.GetQueue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GetQueue`, data.GetQueue.ValueString())
	}
	if !data.GetMessageOptions.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GetMessageOptions`, data.GetMessageOptions.ValueInt64())
	}
	if !data.ConcurrentConnections.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ConcurrentConnections`, data.ConcurrentConnections.ValueInt64())
	}
	if !data.PollingInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PollingInterval`, data.PollingInterval.ValueInt64())
	}
	if !data.RetrieveBackoutSettings.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RetrieveBackoutSettings`, tfutils.StringFromBool(data.RetrieveBackoutSettings, ""))
	}
	if !data.IgnoreBackoutErrors.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IgnoreBackoutErrors`, tfutils.StringFromBool(data.IgnoreBackoutErrors, ""))
	}
	if !data.UseQmNameInUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseQMNameInURL`, tfutils.StringFromBool(data.UseQmNameInUrl, ""))
	}
	return body
}

func (data *MQv9PlusMFTSourceProtocolHandler) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `QueueManager`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.QueueManager = tfutils.ParseStringFromGJSON(value)
	} else {
		data.QueueManager = types.StringNull()
	}
	if value := res.Get(pathRoot + `GetQueue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.GetQueue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GetQueue = types.StringNull()
	}
	if value := res.Get(pathRoot + `GetMessageOptions`); value.Exists() {
		data.GetMessageOptions = types.Int64Value(value.Int())
	} else {
		data.GetMessageOptions = types.Int64Value(32769)
	}
	if value := res.Get(pathRoot + `ConcurrentConnections`); value.Exists() {
		data.ConcurrentConnections = types.Int64Value(value.Int())
	} else {
		data.ConcurrentConnections = types.Int64Value(1)
	}
	if value := res.Get(pathRoot + `PollingInterval`); value.Exists() {
		data.PollingInterval = types.Int64Value(value.Int())
	} else {
		data.PollingInterval = types.Int64Value(30)
	}
	if value := res.Get(pathRoot + `RetrieveBackoutSettings`); value.Exists() {
		data.RetrieveBackoutSettings = tfutils.BoolFromString(value.String())
	} else {
		data.RetrieveBackoutSettings = types.BoolNull()
	}
	if value := res.Get(pathRoot + `IgnoreBackoutErrors`); value.Exists() {
		data.IgnoreBackoutErrors = tfutils.BoolFromString(value.String())
	} else {
		data.IgnoreBackoutErrors = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UseQMNameInURL`); value.Exists() {
		data.UseQmNameInUrl = tfutils.BoolFromString(value.String())
	} else {
		data.UseQmNameInUrl = types.BoolNull()
	}
}

func (data *MQv9PlusMFTSourceProtocolHandler) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `QueueManager`); value.Exists() && !data.QueueManager.IsNull() {
		data.QueueManager = tfutils.ParseStringFromGJSON(value)
	} else {
		data.QueueManager = types.StringNull()
	}
	if value := res.Get(pathRoot + `GetQueue`); value.Exists() && !data.GetQueue.IsNull() {
		data.GetQueue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GetQueue = types.StringNull()
	}
	if value := res.Get(pathRoot + `GetMessageOptions`); value.Exists() && !data.GetMessageOptions.IsNull() {
		data.GetMessageOptions = types.Int64Value(value.Int())
	} else if data.GetMessageOptions.ValueInt64() != 32769 {
		data.GetMessageOptions = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ConcurrentConnections`); value.Exists() && !data.ConcurrentConnections.IsNull() {
		data.ConcurrentConnections = types.Int64Value(value.Int())
	} else if data.ConcurrentConnections.ValueInt64() != 1 {
		data.ConcurrentConnections = types.Int64Null()
	}
	if value := res.Get(pathRoot + `PollingInterval`); value.Exists() && !data.PollingInterval.IsNull() {
		data.PollingInterval = types.Int64Value(value.Int())
	} else if data.PollingInterval.ValueInt64() != 30 {
		data.PollingInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `RetrieveBackoutSettings`); value.Exists() && !data.RetrieveBackoutSettings.IsNull() {
		data.RetrieveBackoutSettings = tfutils.BoolFromString(value.String())
	} else if data.RetrieveBackoutSettings.ValueBool() {
		data.RetrieveBackoutSettings = types.BoolNull()
	}
	if value := res.Get(pathRoot + `IgnoreBackoutErrors`); value.Exists() && !data.IgnoreBackoutErrors.IsNull() {
		data.IgnoreBackoutErrors = tfutils.BoolFromString(value.String())
	} else if data.IgnoreBackoutErrors.ValueBool() {
		data.IgnoreBackoutErrors = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UseQMNameInURL`); value.Exists() && !data.UseQmNameInUrl.IsNull() {
		data.UseQmNameInUrl = tfutils.BoolFromString(value.String())
	} else if data.UseQmNameInUrl.ValueBool() {
		data.UseQmNameInUrl = types.BoolNull()
	}
}
