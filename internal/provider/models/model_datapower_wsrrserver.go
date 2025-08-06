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

type WSRRServer struct {
	Id                  types.String      `tfsdk:"id"`
	AppDomain           types.String      `tfsdk:"app_domain"`
	UserSummary         types.String      `tfsdk:"user_summary"`
	SoapUrl             types.String      `tfsdk:"soap_url"`
	ServerPrefix        types.String      `tfsdk:"server_prefix"`
	Username            types.String      `tfsdk:"username"`
	PasswordAlias       types.String      `tfsdk:"password_alias"`
	SslClientConfigType types.String      `tfsdk:"ssl_client_config_type"`
	SslClient           types.String      `tfsdk:"ssl_client"`
	ObjectActions       []*actions.Action `tfsdk:"object_actions"`
}

var WSRRServerObjectType = map[string]attr.Type{
	"id":                     types.StringType,
	"app_domain":             types.StringType,
	"user_summary":           types.StringType,
	"soap_url":               types.StringType,
	"server_prefix":          types.StringType,
	"username":               types.StringType,
	"password_alias":         types.StringType,
	"ssl_client_config_type": types.StringType,
	"ssl_client":             types.StringType,
	"object_actions":         actions.ActionsListType,
}

func (data WSRRServer) GetPath() string {
	rest_path := "/mgmt/config/{domain}/WSRRServer"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data WSRRServer) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.SoapUrl.IsNull() {
		return false
	}
	if !data.ServerPrefix.IsNull() {
		return false
	}
	if !data.Username.IsNull() {
		return false
	}
	if !data.PasswordAlias.IsNull() {
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

func (data WSRRServer) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.SoapUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SoapURL`, data.SoapUrl.ValueString())
	}
	if !data.ServerPrefix.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ServerPrefix`, data.ServerPrefix.ValueString())
	}
	if !data.Username.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Username`, data.Username.ValueString())
	}
	if !data.PasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PasswordAlias`, data.PasswordAlias.ValueString())
	}
	if !data.SslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClientConfigType`, data.SslClientConfigType.ValueString())
	}
	if !data.SslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClient`, data.SslClient.ValueString())
	}
	return body
}

func (data *WSRRServer) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `SoapURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SoapUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SoapUrl = types.StringValue("https://host:9443/WSRRCoreSDO/services/WSRRCoreSDOPort")
	}
	if value := res.Get(pathRoot + `ServerPrefix`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ServerPrefix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ServerPrefix = types.StringNull()
	}
	if value := res.Get(pathRoot + `Username`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Username = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Username = types.StringNull()
	}
	if value := res.Get(pathRoot + `PasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordAlias = types.StringNull()
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

func (data *WSRRServer) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `SoapURL`); value.Exists() && !data.SoapUrl.IsNull() {
		data.SoapUrl = tfutils.ParseStringFromGJSON(value)
	} else if data.SoapUrl.ValueString() != "https://host:9443/WSRRCoreSDO/services/WSRRCoreSDOPort" {
		data.SoapUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `ServerPrefix`); value.Exists() && !data.ServerPrefix.IsNull() {
		data.ServerPrefix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ServerPrefix = types.StringNull()
	}
	if value := res.Get(pathRoot + `Username`); value.Exists() && !data.Username.IsNull() {
		data.Username = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Username = types.StringNull()
	}
	if value := res.Get(pathRoot + `PasswordAlias`); value.Exists() && !data.PasswordAlias.IsNull() {
		data.PasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordAlias = types.StringNull()
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
