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

type WebSphereJMSServer struct {
	Id                    types.String                `tfsdk:"id"`
	AppDomain             types.String                `tfsdk:"app_domain"`
	Endpoint              types.List                  `tfsdk:"endpoint"`
	TargetTransportChain  types.String                `tfsdk:"target_transport_chain"`
	MessagingBus          types.String                `tfsdk:"messaging_bus"`
	SslCipher             types.String                `tfsdk:"ssl_cipher"`
	Fips                  types.Bool                  `tfsdk:"fips"`
	UserSummary           types.String                `tfsdk:"user_summary"`
	UserName              types.String                `tfsdk:"user_name"`
	PasswordAlias         types.String                `tfsdk:"password_alias"`
	Transactional         types.Bool                  `tfsdk:"transactional"`
	MemoryThreshold       types.Int64                 `tfsdk:"memory_threshold"`
	MaximumMessageSize    types.Int64                 `tfsdk:"maximum_message_size"`
	DefaultMessageType    types.String                `tfsdk:"default_message_type"`
	TotalConnectionLimit  types.Int64                 `tfsdk:"total_connection_limit"`
	SessionsPerConnection types.Int64                 `tfsdk:"sessions_per_connection"`
	AutoRetry             types.Bool                  `tfsdk:"auto_retry"`
	RetryInterval         types.Int64                 `tfsdk:"retry_interval"`
	EnableLogging         types.Bool                  `tfsdk:"enable_logging"`
	SslClientConfigType   types.String                `tfsdk:"ssl_client_config_type"`
	SslClient             types.String                `tfsdk:"ssl_client"`
	DependencyActions     []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var WebSphereJMSServerObjectType = map[string]attr.Type{
	"id":                      types.StringType,
	"app_domain":              types.StringType,
	"endpoint":                types.ListType{ElemType: types.ObjectType{AttrTypes: DmWebSphereJMSEndpointObjectType}},
	"target_transport_chain":  types.StringType,
	"messaging_bus":           types.StringType,
	"ssl_cipher":              types.StringType,
	"fips":                    types.BoolType,
	"user_summary":            types.StringType,
	"user_name":               types.StringType,
	"password_alias":          types.StringType,
	"transactional":           types.BoolType,
	"memory_threshold":        types.Int64Type,
	"maximum_message_size":    types.Int64Type,
	"default_message_type":    types.StringType,
	"total_connection_limit":  types.Int64Type,
	"sessions_per_connection": types.Int64Type,
	"auto_retry":              types.BoolType,
	"retry_interval":          types.Int64Type,
	"enable_logging":          types.BoolType,
	"ssl_client_config_type":  types.StringType,
	"ssl_client":              types.StringType,
	"dependency_actions":      actions.ActionsListType,
}

func (data WebSphereJMSServer) GetPath() string {
	rest_path := "/mgmt/config/{domain}/WebSphereJMSServer"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data WebSphereJMSServer) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Endpoint.IsNull() {
		return false
	}
	if !data.TargetTransportChain.IsNull() {
		return false
	}
	if !data.MessagingBus.IsNull() {
		return false
	}
	if !data.SslCipher.IsNull() {
		return false
	}
	if !data.Fips.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.UserName.IsNull() {
		return false
	}
	if !data.PasswordAlias.IsNull() {
		return false
	}
	if !data.Transactional.IsNull() {
		return false
	}
	if !data.MemoryThreshold.IsNull() {
		return false
	}
	if !data.MaximumMessageSize.IsNull() {
		return false
	}
	if !data.DefaultMessageType.IsNull() {
		return false
	}
	if !data.TotalConnectionLimit.IsNull() {
		return false
	}
	if !data.SessionsPerConnection.IsNull() {
		return false
	}
	if !data.AutoRetry.IsNull() {
		return false
	}
	if !data.RetryInterval.IsNull() {
		return false
	}
	if !data.EnableLogging.IsNull() {
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

func (data WebSphereJMSServer) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.Endpoint.IsNull() {
		var dataValues []DmWebSphereJMSEndpoint
		data.Endpoint.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`Endpoint`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.TargetTransportChain.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TargetTransportChain`, data.TargetTransportChain.ValueString())
	}
	if !data.MessagingBus.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MessagingBus`, data.MessagingBus.ValueString())
	}
	if !data.SslCipher.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLCipher`, data.SslCipher.ValueString())
	}
	if !data.Fips.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FIPS`, tfutils.StringFromBool(data.Fips, ""))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.UserName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserName`, data.UserName.ValueString())
	}
	if !data.PasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PasswordAlias`, data.PasswordAlias.ValueString())
	}
	if !data.Transactional.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Transactional`, tfutils.StringFromBool(data.Transactional, ""))
	}
	if !data.MemoryThreshold.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MemoryThreshold`, data.MemoryThreshold.ValueInt64())
	}
	if !data.MaximumMessageSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaximumMessageSize`, data.MaximumMessageSize.ValueInt64())
	}
	if !data.DefaultMessageType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DefaultMessageType`, data.DefaultMessageType.ValueString())
	}
	if !data.TotalConnectionLimit.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TotalConnectionLimit`, data.TotalConnectionLimit.ValueInt64())
	}
	if !data.SessionsPerConnection.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SessionsPerConnection`, data.SessionsPerConnection.ValueInt64())
	}
	if !data.AutoRetry.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AutoRetry`, tfutils.StringFromBool(data.AutoRetry, ""))
	}
	if !data.RetryInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RetryInterval`, data.RetryInterval.ValueInt64())
	}
	if !data.EnableLogging.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnableLogging`, tfutils.StringFromBool(data.EnableLogging, ""))
	}
	if !data.SslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClientConfigType`, data.SslClientConfigType.ValueString())
	}
	if !data.SslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClient`, data.SslClient.ValueString())
	}
	return body
}

func (data *WebSphereJMSServer) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Endpoint`); value.Exists() {
		l := []DmWebSphereJMSEndpoint{}
		if value := res.Get(`Endpoint`); value.Exists() {
			for _, v := range value.Array() {
				item := DmWebSphereJMSEndpoint{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Endpoint, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWebSphereJMSEndpointObjectType}, l)
		} else {
			data.Endpoint = types.ListNull(types.ObjectType{AttrTypes: DmWebSphereJMSEndpointObjectType})
		}
	} else {
		data.Endpoint = types.ListNull(types.ObjectType{AttrTypes: DmWebSphereJMSEndpointObjectType})
	}
	if value := res.Get(pathRoot + `TargetTransportChain`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TargetTransportChain = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TargetTransportChain = types.StringValue("InboundBasicMessaging")
	}
	if value := res.Get(pathRoot + `MessagingBus`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MessagingBus = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MessagingBus = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLCipher`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslCipher = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslCipher = types.StringNull()
	}
	if value := res.Get(pathRoot + `FIPS`); value.Exists() {
		data.Fips = tfutils.BoolFromString(value.String())
	} else {
		data.Fips = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
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
	if value := res.Get(pathRoot + `Transactional`); value.Exists() {
		data.Transactional = tfutils.BoolFromString(value.String())
	} else {
		data.Transactional = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MemoryThreshold`); value.Exists() {
		data.MemoryThreshold = types.Int64Value(value.Int())
	} else {
		data.MemoryThreshold = types.Int64Value(268435456)
	}
	if value := res.Get(pathRoot + `MaximumMessageSize`); value.Exists() {
		data.MaximumMessageSize = types.Int64Value(value.Int())
	} else {
		data.MaximumMessageSize = types.Int64Value(1048576)
	}
	if value := res.Get(pathRoot + `DefaultMessageType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DefaultMessageType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DefaultMessageType = types.StringValue("byte")
	}
	if value := res.Get(pathRoot + `TotalConnectionLimit`); value.Exists() {
		data.TotalConnectionLimit = types.Int64Value(value.Int())
	} else {
		data.TotalConnectionLimit = types.Int64Value(64)
	}
	if value := res.Get(pathRoot + `SessionsPerConnection`); value.Exists() {
		data.SessionsPerConnection = types.Int64Value(value.Int())
	} else {
		data.SessionsPerConnection = types.Int64Value(100)
	}
	if value := res.Get(pathRoot + `AutoRetry`); value.Exists() {
		data.AutoRetry = tfutils.BoolFromString(value.String())
	} else {
		data.AutoRetry = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RetryInterval`); value.Exists() {
		data.RetryInterval = types.Int64Value(value.Int())
	} else {
		data.RetryInterval = types.Int64Value(1)
	}
	if value := res.Get(pathRoot + `EnableLogging`); value.Exists() {
		data.EnableLogging = tfutils.BoolFromString(value.String())
	} else {
		data.EnableLogging = types.BoolNull()
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

func (data *WebSphereJMSServer) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Endpoint`); value.Exists() && !data.Endpoint.IsNull() {
		l := []DmWebSphereJMSEndpoint{}
		for _, v := range value.Array() {
			item := DmWebSphereJMSEndpoint{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.Endpoint, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWebSphereJMSEndpointObjectType}, l)
		} else {
			data.Endpoint = types.ListNull(types.ObjectType{AttrTypes: DmWebSphereJMSEndpointObjectType})
		}
	} else {
		data.Endpoint = types.ListNull(types.ObjectType{AttrTypes: DmWebSphereJMSEndpointObjectType})
	}
	if value := res.Get(pathRoot + `TargetTransportChain`); value.Exists() && !data.TargetTransportChain.IsNull() {
		data.TargetTransportChain = tfutils.ParseStringFromGJSON(value)
	} else if data.TargetTransportChain.ValueString() != "InboundBasicMessaging" {
		data.TargetTransportChain = types.StringNull()
	}
	if value := res.Get(pathRoot + `MessagingBus`); value.Exists() && !data.MessagingBus.IsNull() {
		data.MessagingBus = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MessagingBus = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLCipher`); value.Exists() && !data.SslCipher.IsNull() {
		data.SslCipher = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslCipher = types.StringNull()
	}
	if value := res.Get(pathRoot + `FIPS`); value.Exists() && !data.Fips.IsNull() {
		data.Fips = tfutils.BoolFromString(value.String())
	} else if data.Fips.ValueBool() {
		data.Fips = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
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
	if value := res.Get(pathRoot + `Transactional`); value.Exists() && !data.Transactional.IsNull() {
		data.Transactional = tfutils.BoolFromString(value.String())
	} else if data.Transactional.ValueBool() {
		data.Transactional = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MemoryThreshold`); value.Exists() && !data.MemoryThreshold.IsNull() {
		data.MemoryThreshold = types.Int64Value(value.Int())
	} else if data.MemoryThreshold.ValueInt64() != 268435456 {
		data.MemoryThreshold = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaximumMessageSize`); value.Exists() && !data.MaximumMessageSize.IsNull() {
		data.MaximumMessageSize = types.Int64Value(value.Int())
	} else if data.MaximumMessageSize.ValueInt64() != 1048576 {
		data.MaximumMessageSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `DefaultMessageType`); value.Exists() && !data.DefaultMessageType.IsNull() {
		data.DefaultMessageType = tfutils.ParseStringFromGJSON(value)
	} else if data.DefaultMessageType.ValueString() != "byte" {
		data.DefaultMessageType = types.StringNull()
	}
	if value := res.Get(pathRoot + `TotalConnectionLimit`); value.Exists() && !data.TotalConnectionLimit.IsNull() {
		data.TotalConnectionLimit = types.Int64Value(value.Int())
	} else if data.TotalConnectionLimit.ValueInt64() != 64 {
		data.TotalConnectionLimit = types.Int64Null()
	}
	if value := res.Get(pathRoot + `SessionsPerConnection`); value.Exists() && !data.SessionsPerConnection.IsNull() {
		data.SessionsPerConnection = types.Int64Value(value.Int())
	} else if data.SessionsPerConnection.ValueInt64() != 100 {
		data.SessionsPerConnection = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AutoRetry`); value.Exists() && !data.AutoRetry.IsNull() {
		data.AutoRetry = tfutils.BoolFromString(value.String())
	} else if !data.AutoRetry.ValueBool() {
		data.AutoRetry = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RetryInterval`); value.Exists() && !data.RetryInterval.IsNull() {
		data.RetryInterval = types.Int64Value(value.Int())
	} else if data.RetryInterval.ValueInt64() != 1 {
		data.RetryInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `EnableLogging`); value.Exists() && !data.EnableLogging.IsNull() {
		data.EnableLogging = tfutils.BoolFromString(value.String())
	} else if data.EnableLogging.ValueBool() {
		data.EnableLogging = types.BoolNull()
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
