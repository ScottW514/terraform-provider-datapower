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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type AMQPBroker struct {
	Id                types.String `tfsdk:"id"`
	AppDomain         types.String `tfsdk:"app_domain"`
	UserSummary       types.String `tfsdk:"user_summary"`
	HostName          types.String `tfsdk:"host_name"`
	Port              types.Int64  `tfsdk:"port"`
	XmlManager        types.String `tfsdk:"xml_manager"`
	ContainerId       types.String `tfsdk:"container_id"`
	Authorization     types.String `tfsdk:"authorization"`
	UserName          types.String `tfsdk:"user_name"`
	PasswordAlias     types.String `tfsdk:"password_alias"`
	MaximumFrameSize  types.Int64  `tfsdk:"maximum_frame_size"`
	AutoRetry         types.Bool   `tfsdk:"auto_retry"`
	RetryInterval     types.Int64  `tfsdk:"retry_interval"`
	RetryAttempts     types.Int64  `tfsdk:"retry_attempts"`
	LongRetryInterval types.Int64  `tfsdk:"long_retry_interval"`
	ReportingInterval types.Int64  `tfsdk:"reporting_interval"`
	SslClient         types.String `tfsdk:"ssl_client"`
}

var AMQPBrokerObjectType = map[string]attr.Type{
	"id":                  types.StringType,
	"app_domain":          types.StringType,
	"user_summary":        types.StringType,
	"host_name":           types.StringType,
	"port":                types.Int64Type,
	"xml_manager":         types.StringType,
	"container_id":        types.StringType,
	"authorization":       types.StringType,
	"user_name":           types.StringType,
	"password_alias":      types.StringType,
	"maximum_frame_size":  types.Int64Type,
	"auto_retry":          types.BoolType,
	"retry_interval":      types.Int64Type,
	"retry_attempts":      types.Int64Type,
	"long_retry_interval": types.Int64Type,
	"reporting_interval":  types.Int64Type,
	"ssl_client":          types.StringType,
}

func (data AMQPBroker) GetPath() string {
	rest_path := "/mgmt/config/{domain}/AMQPBroker"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data AMQPBroker) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.HostName.IsNull() {
		return false
	}
	if !data.Port.IsNull() {
		return false
	}
	if !data.XmlManager.IsNull() {
		return false
	}
	if !data.ContainerId.IsNull() {
		return false
	}
	if !data.Authorization.IsNull() {
		return false
	}
	if !data.UserName.IsNull() {
		return false
	}
	if !data.PasswordAlias.IsNull() {
		return false
	}
	if !data.MaximumFrameSize.IsNull() {
		return false
	}
	if !data.AutoRetry.IsNull() {
		return false
	}
	if !data.RetryInterval.IsNull() {
		return false
	}
	if !data.RetryAttempts.IsNull() {
		return false
	}
	if !data.LongRetryInterval.IsNull() {
		return false
	}
	if !data.ReportingInterval.IsNull() {
		return false
	}
	if !data.SslClient.IsNull() {
		return false
	}
	return true
}

func (data AMQPBroker) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.HostName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HostName`, data.HostName.ValueString())
	}
	if !data.Port.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Port`, data.Port.ValueInt64())
	}
	if !data.XmlManager.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XMLManager`, data.XmlManager.ValueString())
	}
	if !data.ContainerId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ContainerID`, data.ContainerId.ValueString())
	}
	if !data.Authorization.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Authorization`, data.Authorization.ValueString())
	}
	if !data.UserName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserName`, data.UserName.ValueString())
	}
	if !data.PasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PasswordAlias`, data.PasswordAlias.ValueString())
	}
	if !data.MaximumFrameSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaximumFrameSize`, data.MaximumFrameSize.ValueInt64())
	}
	if !data.AutoRetry.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AutoRetry`, tfutils.StringFromBool(data.AutoRetry, ""))
	}
	if !data.RetryInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RetryInterval`, data.RetryInterval.ValueInt64())
	}
	if !data.RetryAttempts.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RetryAttempts`, data.RetryAttempts.ValueInt64())
	}
	if !data.LongRetryInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LongRetryInterval`, data.LongRetryInterval.ValueInt64())
	}
	if !data.ReportingInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ReportingInterval`, data.ReportingInterval.ValueInt64())
	}
	if !data.SslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClient`, data.SslClient.ValueString())
	}
	return body
}

func (data *AMQPBroker) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `HostName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HostName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HostName = types.StringNull()
	}
	if value := res.Get(pathRoot + `Port`); value.Exists() {
		data.Port = types.Int64Value(value.Int())
	} else {
		data.Port = types.Int64Value(5672)
	}
	if value := res.Get(pathRoot + `XMLManager`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.XmlManager = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XmlManager = types.StringValue("default")
	}
	if value := res.Get(pathRoot + `ContainerID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ContainerId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ContainerId = types.StringNull()
	}
	if value := res.Get(pathRoot + `Authorization`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Authorization = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Authorization = types.StringValue("none")
	}
	if value := res.Get(pathRoot + `UserName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserName = types.StringNull()
	}
	if value := res.Get(pathRoot + `PasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `MaximumFrameSize`); value.Exists() {
		data.MaximumFrameSize = types.Int64Value(value.Int())
	} else {
		data.MaximumFrameSize = types.Int64Value(104857600)
	}
	if value := res.Get(pathRoot + `AutoRetry`); value.Exists() {
		data.AutoRetry = tfutils.BoolFromString(value.String())
	} else {
		data.AutoRetry = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RetryInterval`); value.Exists() {
		data.RetryInterval = types.Int64Value(value.Int())
	} else {
		data.RetryInterval = types.Int64Value(10)
	}
	if value := res.Get(pathRoot + `RetryAttempts`); value.Exists() {
		data.RetryAttempts = types.Int64Value(value.Int())
	} else {
		data.RetryAttempts = types.Int64Value(6)
	}
	if value := res.Get(pathRoot + `LongRetryInterval`); value.Exists() {
		data.LongRetryInterval = types.Int64Value(value.Int())
	} else {
		data.LongRetryInterval = types.Int64Value(600)
	}
	if value := res.Get(pathRoot + `ReportingInterval`); value.Exists() {
		data.ReportingInterval = types.Int64Value(value.Int())
	} else {
		data.ReportingInterval = types.Int64Value(10)
	}
	if value := res.Get(pathRoot + `SSLClient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClient = types.StringNull()
	}
}

func (data *AMQPBroker) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `HostName`); value.Exists() && !data.HostName.IsNull() {
		data.HostName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HostName = types.StringNull()
	}
	if value := res.Get(pathRoot + `Port`); value.Exists() && !data.Port.IsNull() {
		data.Port = types.Int64Value(value.Int())
	} else if data.Port.ValueInt64() != 5672 {
		data.Port = types.Int64Null()
	}
	if value := res.Get(pathRoot + `XMLManager`); value.Exists() && !data.XmlManager.IsNull() {
		data.XmlManager = tfutils.ParseStringFromGJSON(value)
	} else if data.XmlManager.ValueString() != "default" {
		data.XmlManager = types.StringNull()
	}
	if value := res.Get(pathRoot + `ContainerID`); value.Exists() && !data.ContainerId.IsNull() {
		data.ContainerId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ContainerId = types.StringNull()
	}
	if value := res.Get(pathRoot + `Authorization`); value.Exists() && !data.Authorization.IsNull() {
		data.Authorization = tfutils.ParseStringFromGJSON(value)
	} else if data.Authorization.ValueString() != "none" {
		data.Authorization = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserName`); value.Exists() && !data.UserName.IsNull() {
		data.UserName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserName = types.StringNull()
	}
	if value := res.Get(pathRoot + `PasswordAlias`); value.Exists() && !data.PasswordAlias.IsNull() {
		data.PasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `MaximumFrameSize`); value.Exists() && !data.MaximumFrameSize.IsNull() {
		data.MaximumFrameSize = types.Int64Value(value.Int())
	} else if data.MaximumFrameSize.ValueInt64() != 104857600 {
		data.MaximumFrameSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AutoRetry`); value.Exists() && !data.AutoRetry.IsNull() {
		data.AutoRetry = tfutils.BoolFromString(value.String())
	} else if !data.AutoRetry.ValueBool() {
		data.AutoRetry = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RetryInterval`); value.Exists() && !data.RetryInterval.IsNull() {
		data.RetryInterval = types.Int64Value(value.Int())
	} else if data.RetryInterval.ValueInt64() != 10 {
		data.RetryInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `RetryAttempts`); value.Exists() && !data.RetryAttempts.IsNull() {
		data.RetryAttempts = types.Int64Value(value.Int())
	} else if data.RetryAttempts.ValueInt64() != 6 {
		data.RetryAttempts = types.Int64Null()
	}
	if value := res.Get(pathRoot + `LongRetryInterval`); value.Exists() && !data.LongRetryInterval.IsNull() {
		data.LongRetryInterval = types.Int64Value(value.Int())
	} else if data.LongRetryInterval.ValueInt64() != 600 {
		data.LongRetryInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ReportingInterval`); value.Exists() && !data.ReportingInterval.IsNull() {
		data.ReportingInterval = types.Int64Value(value.Int())
	} else if data.ReportingInterval.ValueInt64() != 10 {
		data.ReportingInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `SSLClient`); value.Exists() && !data.SslClient.IsNull() {
		data.SslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClient = types.StringNull()
	}
}
