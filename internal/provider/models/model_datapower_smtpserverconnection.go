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

type SMTPServerConnection struct {
	Id                   types.String      `tfsdk:"id"`
	AppDomain            types.String      `tfsdk:"app_domain"`
	UserSummary          types.String      `tfsdk:"user_summary"`
	MailServerHost       types.String      `tfsdk:"mail_server_host"`
	MailServerPort       types.Int64       `tfsdk:"mail_server_port"`
	Options              *DmSMTPOptions    `tfsdk:"options"`
	Auth                 types.String      `tfsdk:"auth"`
	AccountName          types.String      `tfsdk:"account_name"`
	AccountPasswordAlias types.String      `tfsdk:"account_password_alias"`
	SslClientConfigType  types.String      `tfsdk:"ssl_client_config_type"`
	SslClient            types.String      `tfsdk:"ssl_client"`
	ObjectActions        []*actions.Action `tfsdk:"object_actions"`
}

var SMTPServerConnectionObjectType = map[string]attr.Type{
	"id":                     types.StringType,
	"app_domain":             types.StringType,
	"user_summary":           types.StringType,
	"mail_server_host":       types.StringType,
	"mail_server_port":       types.Int64Type,
	"options":                types.ObjectType{AttrTypes: DmSMTPOptionsObjectType},
	"auth":                   types.StringType,
	"account_name":           types.StringType,
	"account_password_alias": types.StringType,
	"ssl_client_config_type": types.StringType,
	"ssl_client":             types.StringType,
	"object_actions":         actions.ActionsListType,
}

func (data SMTPServerConnection) GetPath() string {
	rest_path := "/mgmt/config/{domain}/SMTPServerConnection"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data SMTPServerConnection) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.MailServerHost.IsNull() {
		return false
	}
	if !data.MailServerPort.IsNull() {
		return false
	}
	if data.Options != nil {
		if !data.Options.IsNull() {
			return false
		}
	}
	if !data.Auth.IsNull() {
		return false
	}
	if !data.AccountName.IsNull() {
		return false
	}
	if !data.AccountPasswordAlias.IsNull() {
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

func (data SMTPServerConnection) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.MailServerHost.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MailServerHost`, data.MailServerHost.ValueString())
	}
	if !data.MailServerPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MailServerPort`, data.MailServerPort.ValueInt64())
	}
	if data.Options != nil {
		if !data.Options.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`Options`, data.Options.ToBody(ctx, ""))
		}
	}
	if !data.Auth.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Auth`, data.Auth.ValueString())
	}
	if !data.AccountName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AccountName`, data.AccountName.ValueString())
	}
	if !data.AccountPasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AccountPasswordAlias`, data.AccountPasswordAlias.ValueString())
	}
	if !data.SslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClientConfigType`, data.SslClientConfigType.ValueString())
	}
	if !data.SslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClient`, data.SslClient.ValueString())
	}
	return body
}

func (data *SMTPServerConnection) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `MailServerHost`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MailServerHost = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MailServerHost = types.StringNull()
	}
	if value := res.Get(pathRoot + `MailServerPort`); value.Exists() {
		data.MailServerPort = types.Int64Value(value.Int())
	} else {
		data.MailServerPort = types.Int64Value(25)
	}
	if value := res.Get(pathRoot + `Options`); value.Exists() {
		data.Options = &DmSMTPOptions{}
		data.Options.FromBody(ctx, "", value)
	} else {
		data.Options = nil
	}
	if value := res.Get(pathRoot + `Auth`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Auth = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Auth = types.StringValue("plain")
	}
	if value := res.Get(pathRoot + `AccountName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AccountName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AccountName = types.StringNull()
	}
	if value := res.Get(pathRoot + `AccountPasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AccountPasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AccountPasswordAlias = types.StringNull()
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

func (data *SMTPServerConnection) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `MailServerHost`); value.Exists() && !data.MailServerHost.IsNull() {
		data.MailServerHost = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MailServerHost = types.StringNull()
	}
	if value := res.Get(pathRoot + `MailServerPort`); value.Exists() && !data.MailServerPort.IsNull() {
		data.MailServerPort = types.Int64Value(value.Int())
	} else if data.MailServerPort.ValueInt64() != 25 {
		data.MailServerPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Options`); value.Exists() {
		data.Options.UpdateFromBody(ctx, "", value)
	} else {
		data.Options = nil
	}
	if value := res.Get(pathRoot + `Auth`); value.Exists() && !data.Auth.IsNull() {
		data.Auth = tfutils.ParseStringFromGJSON(value)
	} else if data.Auth.ValueString() != "plain" {
		data.Auth = types.StringNull()
	}
	if value := res.Get(pathRoot + `AccountName`); value.Exists() && !data.AccountName.IsNull() {
		data.AccountName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AccountName = types.StringNull()
	}
	if value := res.Get(pathRoot + `AccountPasswordAlias`); value.Exists() && !data.AccountPasswordAlias.IsNull() {
		data.AccountPasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AccountPasswordAlias = types.StringNull()
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
