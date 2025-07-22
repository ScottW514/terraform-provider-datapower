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

type WCCService struct {
	Id                  types.String `tfsdk:"id"`
	AppDomain           types.String `tfsdk:"app_domain"`
	UserSummary         types.String `tfsdk:"user_summary"`
	OdcInfoHostname     types.String `tfsdk:"odc_info_hostname"`
	OdcInfoPort         types.Int64  `tfsdk:"odc_info_port"`
	UpdateType          types.String `tfsdk:"update_type"`
	TimeInterval        types.Int64  `tfsdk:"time_interval"`
	SslClientConfigType types.String `tfsdk:"ssl_client_config_type"`
	SslClient           types.String `tfsdk:"ssl_client"`
}

var WCCServiceObjectType = map[string]attr.Type{
	"id":                     types.StringType,
	"app_domain":             types.StringType,
	"user_summary":           types.StringType,
	"odc_info_hostname":      types.StringType,
	"odc_info_port":          types.Int64Type,
	"update_type":            types.StringType,
	"time_interval":          types.Int64Type,
	"ssl_client_config_type": types.StringType,
	"ssl_client":             types.StringType,
}

func (data WCCService) GetPath() string {
	rest_path := "/mgmt/config/{domain}/WCCService"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data WCCService) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.OdcInfoHostname.IsNull() {
		return false
	}
	if !data.OdcInfoPort.IsNull() {
		return false
	}
	if !data.UpdateType.IsNull() {
		return false
	}
	if !data.TimeInterval.IsNull() {
		return false
	}
	if !data.SslClientConfigType.IsNull() {
		return false
	}
	if !data.SslClient.IsNull() {
		return false
	}
	return true
}

func (data WCCService) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.OdcInfoHostname.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ODCInfoHostname`, data.OdcInfoHostname.ValueString())
	}
	if !data.OdcInfoPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ODCInfoPort`, data.OdcInfoPort.ValueInt64())
	}
	if !data.UpdateType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UpdateType`, data.UpdateType.ValueString())
	}
	if !data.TimeInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TimeInterval`, data.TimeInterval.ValueInt64())
	}
	if !data.SslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClientConfigType`, data.SslClientConfigType.ValueString())
	}
	if !data.SslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClient`, data.SslClient.ValueString())
	}
	return body
}

func (data *WCCService) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `ODCInfoHostname`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OdcInfoHostname = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OdcInfoHostname = types.StringNull()
	}
	if value := res.Get(pathRoot + `ODCInfoPort`); value.Exists() {
		data.OdcInfoPort = types.Int64Value(value.Int())
	} else {
		data.OdcInfoPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `UpdateType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UpdateType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UpdateType = types.StringValue("poll")
	}
	if value := res.Get(pathRoot + `TimeInterval`); value.Exists() {
		data.TimeInterval = types.Int64Value(value.Int())
	} else {
		data.TimeInterval = types.Int64Value(10)
	}
	if value := res.Get(pathRoot + `SSLClientConfigType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClientConfigType = types.StringValue("client")
	}
	if value := res.Get(pathRoot + `SSLClient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClient = types.StringNull()
	}
}

func (data *WCCService) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `ODCInfoHostname`); value.Exists() && !data.OdcInfoHostname.IsNull() {
		data.OdcInfoHostname = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OdcInfoHostname = types.StringNull()
	}
	if value := res.Get(pathRoot + `ODCInfoPort`); value.Exists() && !data.OdcInfoPort.IsNull() {
		data.OdcInfoPort = types.Int64Value(value.Int())
	} else {
		data.OdcInfoPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `UpdateType`); value.Exists() && !data.UpdateType.IsNull() {
		data.UpdateType = tfutils.ParseStringFromGJSON(value)
	} else if data.UpdateType.ValueString() != "poll" {
		data.UpdateType = types.StringNull()
	}
	if value := res.Get(pathRoot + `TimeInterval`); value.Exists() && !data.TimeInterval.IsNull() {
		data.TimeInterval = types.Int64Value(value.Int())
	} else if data.TimeInterval.ValueInt64() != 10 {
		data.TimeInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `SSLClientConfigType`); value.Exists() && !data.SslClientConfigType.IsNull() {
		data.SslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else if data.SslClientConfigType.ValueString() != "client" {
		data.SslClientConfigType = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLClient`); value.Exists() && !data.SslClient.IsNull() {
		data.SslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClient = types.StringNull()
	}
}
