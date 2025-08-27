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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type SSLProxyService struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	Priority          types.String                `tfsdk:"priority"`
	LocalPort         types.Int64                 `tfsdk:"local_port"`
	RemoteAddress     types.String                `tfsdk:"remote_address"`
	RemotePort        types.Int64                 `tfsdk:"remote_port"`
	FrontTimeout      types.Int64                 `tfsdk:"front_timeout"`
	BackTimeout       types.Int64                 `tfsdk:"back_timeout"`
	ConnTimeout       types.Int64                 `tfsdk:"conn_timeout"`
	ConnLimit         types.Int64                 `tfsdk:"conn_limit"`
	SslConfigType     types.String                `tfsdk:"ssl_config_type"`
	SslServer         types.String                `tfsdk:"ssl_server"`
	SslSniServer      types.String                `tfsdk:"ssl_sni_server"`
	SslClient         types.String                `tfsdk:"ssl_client"`
	LocalAddress      types.String                `tfsdk:"local_address"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var SSLProxyServiceSSLServerIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "ssl_config_type",
	AttrType:    "String",
	AttrDefault: "server",
	Value:       []string{"server"},
}
var SSLProxyServiceSSLSNIServerIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "ssl_config_type",
	AttrType:    "String",
	AttrDefault: "server",
	Value:       []string{"sni"},
}
var SSLProxyServiceSSLClientIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "ssl_config_type",
	AttrType:    "String",
	AttrDefault: "server",
	Value:       []string{"proxy"},
}

var SSLProxyServiceObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"user_summary":       types.StringType,
	"priority":           types.StringType,
	"local_port":         types.Int64Type,
	"remote_address":     types.StringType,
	"remote_port":        types.Int64Type,
	"front_timeout":      types.Int64Type,
	"back_timeout":       types.Int64Type,
	"conn_timeout":       types.Int64Type,
	"conn_limit":         types.Int64Type,
	"ssl_config_type":    types.StringType,
	"ssl_server":         types.StringType,
	"ssl_sni_server":     types.StringType,
	"ssl_client":         types.StringType,
	"local_address":      types.StringType,
	"dependency_actions": actions.ActionsListType,
}

func (data SSLProxyService) GetPath() string {
	rest_path := "/mgmt/config/{domain}/SSLProxyService"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data SSLProxyService) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Priority.IsNull() {
		return false
	}
	if !data.LocalPort.IsNull() {
		return false
	}
	if !data.RemoteAddress.IsNull() {
		return false
	}
	if !data.RemotePort.IsNull() {
		return false
	}
	if !data.FrontTimeout.IsNull() {
		return false
	}
	if !data.BackTimeout.IsNull() {
		return false
	}
	if !data.ConnTimeout.IsNull() {
		return false
	}
	if !data.ConnLimit.IsNull() {
		return false
	}
	if !data.SslConfigType.IsNull() {
		return false
	}
	if !data.SslServer.IsNull() {
		return false
	}
	if !data.SslSniServer.IsNull() {
		return false
	}
	if !data.SslClient.IsNull() {
		return false
	}
	if !data.LocalAddress.IsNull() {
		return false
	}
	return true
}

func (data SSLProxyService) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Priority.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Priority`, data.Priority.ValueString())
	}
	if !data.LocalPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalPort`, data.LocalPort.ValueInt64())
	}
	if !data.RemoteAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemoteAddress`, data.RemoteAddress.ValueString())
	}
	if !data.RemotePort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemotePort`, data.RemotePort.ValueInt64())
	}
	if !data.FrontTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FrontTimeout`, data.FrontTimeout.ValueInt64())
	}
	if !data.BackTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BackTimeout`, data.BackTimeout.ValueInt64())
	}
	if !data.ConnTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ConnTimeout`, data.ConnTimeout.ValueInt64())
	}
	if !data.ConnLimit.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ConnLimit`, data.ConnLimit.ValueInt64())
	}
	if !data.SslConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLConfigType`, data.SslConfigType.ValueString())
	}
	if !data.SslServer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLServer`, data.SslServer.ValueString())
	}
	if !data.SslSniServer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLSNIServer`, data.SslSniServer.ValueString())
	}
	if !data.SslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClient`, data.SslClient.ValueString())
	}
	if !data.LocalAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalAddress`, data.LocalAddress.ValueString())
	}
	return body
}

func (data *SSLProxyService) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Priority`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Priority = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Priority = types.StringValue("normal")
	}
	if value := res.Get(pathRoot + `LocalPort`); value.Exists() {
		data.LocalPort = types.Int64Value(value.Int())
	} else {
		data.LocalPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `RemoteAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RemoteAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemotePort`); value.Exists() {
		data.RemotePort = types.Int64Value(value.Int())
	} else {
		data.RemotePort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `FrontTimeout`); value.Exists() {
		data.FrontTimeout = types.Int64Value(value.Int())
	} else {
		data.FrontTimeout = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `BackTimeout`); value.Exists() {
		data.BackTimeout = types.Int64Value(value.Int())
	} else {
		data.BackTimeout = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `ConnTimeout`); value.Exists() {
		data.ConnTimeout = types.Int64Value(value.Int())
	} else {
		data.ConnTimeout = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `ConnLimit`); value.Exists() {
		data.ConnLimit = types.Int64Value(value.Int())
	} else {
		data.ConnLimit = types.Int64Value(100)
	}
	if value := res.Get(pathRoot + `SSLConfigType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslConfigType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslConfigType = types.StringValue("server")
	}
	if value := res.Get(pathRoot + `SSLServer`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLSNIServer`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslSniServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslSniServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLClient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalAddress = types.StringValue("0.0.0.0")
	}
}

func (data *SSLProxyService) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Priority`); value.Exists() && !data.Priority.IsNull() {
		data.Priority = tfutils.ParseStringFromGJSON(value)
	} else if data.Priority.ValueString() != "normal" {
		data.Priority = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalPort`); value.Exists() && !data.LocalPort.IsNull() {
		data.LocalPort = types.Int64Value(value.Int())
	} else {
		data.LocalPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `RemoteAddress`); value.Exists() && !data.RemoteAddress.IsNull() {
		data.RemoteAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemotePort`); value.Exists() && !data.RemotePort.IsNull() {
		data.RemotePort = types.Int64Value(value.Int())
	} else {
		data.RemotePort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `FrontTimeout`); value.Exists() && !data.FrontTimeout.IsNull() {
		data.FrontTimeout = types.Int64Value(value.Int())
	} else if data.FrontTimeout.ValueInt64() != 0 {
		data.FrontTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `BackTimeout`); value.Exists() && !data.BackTimeout.IsNull() {
		data.BackTimeout = types.Int64Value(value.Int())
	} else if data.BackTimeout.ValueInt64() != 0 {
		data.BackTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ConnTimeout`); value.Exists() && !data.ConnTimeout.IsNull() {
		data.ConnTimeout = types.Int64Value(value.Int())
	} else if data.ConnTimeout.ValueInt64() != 0 {
		data.ConnTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ConnLimit`); value.Exists() && !data.ConnLimit.IsNull() {
		data.ConnLimit = types.Int64Value(value.Int())
	} else if data.ConnLimit.ValueInt64() != 100 {
		data.ConnLimit = types.Int64Null()
	}
	if value := res.Get(pathRoot + `SSLConfigType`); value.Exists() && !data.SslConfigType.IsNull() {
		data.SslConfigType = tfutils.ParseStringFromGJSON(value)
	} else if data.SslConfigType.ValueString() != "server" {
		data.SslConfigType = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLServer`); value.Exists() && !data.SslServer.IsNull() {
		data.SslServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLSNIServer`); value.Exists() && !data.SslSniServer.IsNull() {
		data.SslSniServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslSniServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLClient`); value.Exists() && !data.SslClient.IsNull() {
		data.SslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && !data.LocalAddress.IsNull() {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else if data.LocalAddress.ValueString() != "0.0.0.0" {
		data.LocalAddress = types.StringNull()
	}
}
