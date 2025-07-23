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
	"path"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type WebGUI struct {
	Enabled              types.Bool   `tfsdk:"enabled"`
	UserSummary          types.String `tfsdk:"user_summary"`
	LocalPort            types.Int64  `tfsdk:"local_port"`
	UserAgent            types.String `tfsdk:"user_agent"`
	SaveConfigOverwrites types.Bool   `tfsdk:"save_config_overwrites"`
	IdleTimeout          types.Int64  `tfsdk:"idle_timeout"`
	Acl                  types.String `tfsdk:"acl"`
	SslServerConfigType  types.String `tfsdk:"ssl_server_config_type"`
	SslServer            types.String `tfsdk:"ssl_server"`
	SslsniServer         types.String `tfsdk:"sslsni_server"`
	EnableSts            types.Bool   `tfsdk:"enable_sts"`
	LocalAddress         types.String `tfsdk:"local_address"`
}

var WebGUIObjectType = map[string]attr.Type{
	"enabled":                types.BoolType,
	"user_summary":           types.StringType,
	"local_port":             types.Int64Type,
	"user_agent":             types.StringType,
	"save_config_overwrites": types.BoolType,
	"idle_timeout":           types.Int64Type,
	"acl":                    types.StringType,
	"ssl_server_config_type": types.StringType,
	"ssl_server":             types.StringType,
	"sslsni_server":          types.StringType,
	"enable_sts":             types.BoolType,
	"local_address":          types.StringType,
}

func (data WebGUI) GetPath() string {
	rest_path := "/mgmt/config/default/WebGUI/WebGUI-Settings"
	return rest_path
}

func (data WebGUI) IsNull() bool {
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.LocalPort.IsNull() {
		return false
	}
	if !data.UserAgent.IsNull() {
		return false
	}
	if !data.SaveConfigOverwrites.IsNull() {
		return false
	}
	if !data.IdleTimeout.IsNull() {
		return false
	}
	if !data.Acl.IsNull() {
		return false
	}
	if !data.SslServerConfigType.IsNull() {
		return false
	}
	if !data.SslServer.IsNull() {
		return false
	}
	if !data.SslsniServer.IsNull() {
		return false
	}
	if !data.EnableSts.IsNull() {
		return false
	}
	if !data.LocalAddress.IsNull() {
		return false
	}
	return true
}

func (data WebGUI) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "WebGUI.name", path.Base("/mgmt/config/default/WebGUI/WebGUI-Settings"))
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.LocalPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalPort`, data.LocalPort.ValueInt64())
	}
	if !data.UserAgent.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserAgent`, data.UserAgent.ValueString())
	}
	if !data.SaveConfigOverwrites.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SaveConfigOverwrites`, tfutils.StringFromBool(data.SaveConfigOverwrites, ""))
	}
	if !data.IdleTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IdleTimeout`, data.IdleTimeout.ValueInt64())
	}
	if !data.Acl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ACL`, data.Acl.ValueString())
	}
	if !data.SslServerConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLServerConfigType`, data.SslServerConfigType.ValueString())
	}
	if !data.SslServer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLServer`, data.SslServer.ValueString())
	}
	if !data.SslsniServer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLSNIServer`, data.SslsniServer.ValueString())
	}
	if !data.EnableSts.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnableSTS`, tfutils.StringFromBool(data.EnableSts, ""))
	}
	if !data.LocalAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalAddress`, data.LocalAddress.ValueString())
	}
	return body
}

func (data *WebGUI) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalPort`); value.Exists() {
		data.LocalPort = types.Int64Value(value.Int())
	} else {
		data.LocalPort = types.Int64Value(9090)
	}
	if value := res.Get(pathRoot + `UserAgent`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserAgent = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserAgent = types.StringNull()
	}
	if value := res.Get(pathRoot + `SaveConfigOverwrites`); value.Exists() {
		data.SaveConfigOverwrites = tfutils.BoolFromString(value.String())
	} else {
		data.SaveConfigOverwrites = types.BoolNull()
	}
	if value := res.Get(pathRoot + `IdleTimeout`); value.Exists() {
		data.IdleTimeout = types.Int64Value(value.Int())
	} else {
		data.IdleTimeout = types.Int64Value(600)
	}
	if value := res.Get(pathRoot + `ACL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Acl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Acl = types.StringValue("web-mgmt")
	}
	if value := res.Get(pathRoot + `SSLServerConfigType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslServerConfigType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslServerConfigType = types.StringValue("server")
	}
	if value := res.Get(pathRoot + `SSLServer`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLSNIServer`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslsniServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslsniServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `EnableSTS`); value.Exists() {
		data.EnableSts = tfutils.BoolFromString(value.String())
	} else {
		data.EnableSts = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalAddress = types.StringValue("0.0.0.0")
	}
}

func (data *WebGUI) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else if !data.Enabled.ValueBool() {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalPort`); value.Exists() && !data.LocalPort.IsNull() {
		data.LocalPort = types.Int64Value(value.Int())
	} else if data.LocalPort.ValueInt64() != 9090 {
		data.LocalPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `UserAgent`); value.Exists() && !data.UserAgent.IsNull() {
		data.UserAgent = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserAgent = types.StringNull()
	}
	if value := res.Get(pathRoot + `SaveConfigOverwrites`); value.Exists() && !data.SaveConfigOverwrites.IsNull() {
		data.SaveConfigOverwrites = tfutils.BoolFromString(value.String())
	} else if !data.SaveConfigOverwrites.ValueBool() {
		data.SaveConfigOverwrites = types.BoolNull()
	}
	if value := res.Get(pathRoot + `IdleTimeout`); value.Exists() && !data.IdleTimeout.IsNull() {
		data.IdleTimeout = types.Int64Value(value.Int())
	} else if data.IdleTimeout.ValueInt64() != 600 {
		data.IdleTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ACL`); value.Exists() && !data.Acl.IsNull() {
		data.Acl = tfutils.ParseStringFromGJSON(value)
	} else if data.Acl.ValueString() != "web-mgmt" {
		data.Acl = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLServerConfigType`); value.Exists() && !data.SslServerConfigType.IsNull() {
		data.SslServerConfigType = tfutils.ParseStringFromGJSON(value)
	} else if data.SslServerConfigType.ValueString() != "server" {
		data.SslServerConfigType = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLServer`); value.Exists() && !data.SslServer.IsNull() {
		data.SslServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLSNIServer`); value.Exists() && !data.SslsniServer.IsNull() {
		data.SslsniServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslsniServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `EnableSTS`); value.Exists() && !data.EnableSts.IsNull() {
		data.EnableSts = tfutils.BoolFromString(value.String())
	} else if !data.EnableSts.ValueBool() {
		data.EnableSts = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && !data.LocalAddress.IsNull() {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else if data.LocalAddress.ValueString() != "0.0.0.0" {
		data.LocalAddress = types.StringNull()
	}
}
