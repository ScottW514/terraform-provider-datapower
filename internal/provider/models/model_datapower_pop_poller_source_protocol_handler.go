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

type POPPollerSourceProtocolHandler struct {
	Id                  types.String                `tfsdk:"id"`
	AppDomain           types.String                `tfsdk:"app_domain"`
	UserSummary         types.String                `tfsdk:"user_summary"`
	MailServer          types.String                `tfsdk:"mail_server"`
	Port                types.Int64                 `tfsdk:"port"`
	ConnSecurity        types.String                `tfsdk:"conn_security"`
	AuthMethod          types.String                `tfsdk:"auth_method"`
	Account             types.String                `tfsdk:"account"`
	PasswordAlias       types.String                `tfsdk:"password_alias"`
	DelayBetweenPolls   types.Int64                 `tfsdk:"delay_between_polls"`
	MaxMessagesPerPoll  types.Int64                 `tfsdk:"max_messages_per_poll"`
	SslClientConfigType types.String                `tfsdk:"ssl_client_config_type"`
	SslClient           types.String                `tfsdk:"ssl_client"`
	DependencyActions   []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var POPPollerSourceProtocolHandlerObjectType = map[string]attr.Type{
	"id":                     types.StringType,
	"app_domain":             types.StringType,
	"user_summary":           types.StringType,
	"mail_server":            types.StringType,
	"port":                   types.Int64Type,
	"conn_security":          types.StringType,
	"auth_method":            types.StringType,
	"account":                types.StringType,
	"password_alias":         types.StringType,
	"delay_between_polls":    types.Int64Type,
	"max_messages_per_poll":  types.Int64Type,
	"ssl_client_config_type": types.StringType,
	"ssl_client":             types.StringType,
	"dependency_actions":     actions.ActionsListType,
}

func (data POPPollerSourceProtocolHandler) GetPath() string {
	rest_path := "/mgmt/config/{domain}/POPPollerSourceProtocolHandler"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data POPPollerSourceProtocolHandler) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.MailServer.IsNull() {
		return false
	}
	if !data.Port.IsNull() {
		return false
	}
	if !data.ConnSecurity.IsNull() {
		return false
	}
	if !data.AuthMethod.IsNull() {
		return false
	}
	if !data.Account.IsNull() {
		return false
	}
	if !data.PasswordAlias.IsNull() {
		return false
	}
	if !data.DelayBetweenPolls.IsNull() {
		return false
	}
	if !data.MaxMessagesPerPoll.IsNull() {
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

func (data POPPollerSourceProtocolHandler) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.MailServer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MailServer`, data.MailServer.ValueString())
	}
	if !data.Port.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Port`, data.Port.ValueInt64())
	}
	if !data.ConnSecurity.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ConnSecurity`, data.ConnSecurity.ValueString())
	}
	if !data.AuthMethod.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AuthMethod`, data.AuthMethod.ValueString())
	}
	if !data.Account.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Account`, data.Account.ValueString())
	}
	if !data.PasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PasswordAlias`, data.PasswordAlias.ValueString())
	}
	if !data.DelayBetweenPolls.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DelayBetweenPolls`, data.DelayBetweenPolls.ValueInt64())
	}
	if !data.MaxMessagesPerPoll.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxMessagesPerPoll`, data.MaxMessagesPerPoll.ValueInt64())
	}
	if !data.SslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClientConfigType`, data.SslClientConfigType.ValueString())
	}
	if !data.SslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClient`, data.SslClient.ValueString())
	}
	return body
}

func (data *POPPollerSourceProtocolHandler) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `MailServer`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MailServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MailServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `Port`); value.Exists() {
		data.Port = types.Int64Value(value.Int())
	} else {
		data.Port = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ConnSecurity`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ConnSecurity = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ConnSecurity = types.StringValue("none")
	}
	if value := res.Get(pathRoot + `AuthMethod`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuthMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuthMethod = types.StringValue("basic")
	}
	if value := res.Get(pathRoot + `Account`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Account = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Account = types.StringNull()
	}
	if value := res.Get(pathRoot + `PasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `DelayBetweenPolls`); value.Exists() {
		data.DelayBetweenPolls = types.Int64Value(value.Int())
	} else {
		data.DelayBetweenPolls = types.Int64Value(300)
	}
	if value := res.Get(pathRoot + `MaxMessagesPerPoll`); value.Exists() {
		data.MaxMessagesPerPoll = types.Int64Value(value.Int())
	} else {
		data.MaxMessagesPerPoll = types.Int64Value(5)
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

func (data *POPPollerSourceProtocolHandler) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `MailServer`); value.Exists() && !data.MailServer.IsNull() {
		data.MailServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MailServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `Port`); value.Exists() && !data.Port.IsNull() {
		data.Port = types.Int64Value(value.Int())
	} else {
		data.Port = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ConnSecurity`); value.Exists() && !data.ConnSecurity.IsNull() {
		data.ConnSecurity = tfutils.ParseStringFromGJSON(value)
	} else if data.ConnSecurity.ValueString() != "none" {
		data.ConnSecurity = types.StringNull()
	}
	if value := res.Get(pathRoot + `AuthMethod`); value.Exists() && !data.AuthMethod.IsNull() {
		data.AuthMethod = tfutils.ParseStringFromGJSON(value)
	} else if data.AuthMethod.ValueString() != "basic" {
		data.AuthMethod = types.StringNull()
	}
	if value := res.Get(pathRoot + `Account`); value.Exists() && !data.Account.IsNull() {
		data.Account = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Account = types.StringNull()
	}
	if value := res.Get(pathRoot + `PasswordAlias`); value.Exists() && !data.PasswordAlias.IsNull() {
		data.PasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `DelayBetweenPolls`); value.Exists() && !data.DelayBetweenPolls.IsNull() {
		data.DelayBetweenPolls = types.Int64Value(value.Int())
	} else if data.DelayBetweenPolls.ValueInt64() != 300 {
		data.DelayBetweenPolls = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxMessagesPerPoll`); value.Exists() && !data.MaxMessagesPerPoll.IsNull() {
		data.MaxMessagesPerPoll = types.Int64Value(value.Int())
	} else if data.MaxMessagesPerPoll.ValueInt64() != 5 {
		data.MaxMessagesPerPoll = types.Int64Null()
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
