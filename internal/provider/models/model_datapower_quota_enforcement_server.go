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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type QuotaEnforcementServer struct {
	Enabled              types.Bool                  `tfsdk:"enabled"`
	UserSummary          types.String                `tfsdk:"user_summary"`
	PasswordAlias        types.String                `tfsdk:"password_alias"`
	RaidVolume           types.String                `tfsdk:"raid_volume"`
	ServerPort           types.Int64                 `tfsdk:"server_port"`
	MonitorPort          types.Int64                 `tfsdk:"monitor_port"`
	EnablePeerGroup      types.Bool                  `tfsdk:"enable_peer_group"`
	EnableSsl            types.Bool                  `tfsdk:"enable_ssl"`
	SslCryptoKey         types.String                `tfsdk:"ssl_crypto_key"`
	SslCryptoCertificate types.String                `tfsdk:"ssl_crypto_certificate"`
	IpAddress            types.String                `tfsdk:"ip_address"`
	Peers                types.List                  `tfsdk:"peers"`
	Priority             types.Int64                 `tfsdk:"priority"`
	StrictMode           types.Bool                  `tfsdk:"strict_mode"`
	DependencyActions    []*actions.DependencyAction `tfsdk:"dependency_actions"`
	ProviderTarget       types.String                `tfsdk:"provider_target"`
}

var QuotaEnforcementServerEnableSSLIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "enable_peer_group",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var QuotaEnforcementServerSSLCryptoKeyCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "enable_peer_group",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "enable_ssl",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
	},
}

var QuotaEnforcementServerSSLCryptoKeyIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "enable_peer_group",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "enable_ssl",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"false"},
		},
	},
}

var QuotaEnforcementServerSSLCryptoCertificateCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "enable_peer_group",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "enable_ssl",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
	},
}

var QuotaEnforcementServerSSLCryptoCertificateIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "enable_peer_group",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "enable_ssl",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"false"},
		},
	},
}

var QuotaEnforcementServerIPAddressCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "enable_peer_group",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var QuotaEnforcementServerIPAddressIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "enable_peer_group",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var QuotaEnforcementServerPeersIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "enable_peer_group",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var QuotaEnforcementServerPriorityCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "enable_peer_group",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var QuotaEnforcementServerPriorityIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "enable_peer_group",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var QuotaEnforcementServerStrictModeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "enable_peer_group",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var QuotaEnforcementServerObjectType = map[string]attr.Type{
	"provider_target":        types.StringType,
	"enabled":                types.BoolType,
	"user_summary":           types.StringType,
	"password_alias":         types.StringType,
	"raid_volume":            types.StringType,
	"server_port":            types.Int64Type,
	"monitor_port":           types.Int64Type,
	"enable_peer_group":      types.BoolType,
	"enable_ssl":             types.BoolType,
	"ssl_crypto_key":         types.StringType,
	"ssl_crypto_certificate": types.StringType,
	"ip_address":             types.StringType,
	"peers":                  types.ListType{ElemType: types.StringType},
	"priority":               types.Int64Type,
	"strict_mode":            types.BoolType,
	"dependency_actions":     actions.ActionsListType,
}

func (data QuotaEnforcementServer) GetPath() string {
	rest_path := "/mgmt/config/default/QuotaEnforcementServer/QuotaEnforcementServer"
	return rest_path
}

func (data QuotaEnforcementServer) IsNull() bool {
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.PasswordAlias.IsNull() {
		return false
	}
	if !data.RaidVolume.IsNull() {
		return false
	}
	if !data.ServerPort.IsNull() {
		return false
	}
	if !data.MonitorPort.IsNull() {
		return false
	}
	if !data.EnablePeerGroup.IsNull() {
		return false
	}
	if !data.EnableSsl.IsNull() {
		return false
	}
	if !data.SslCryptoKey.IsNull() {
		return false
	}
	if !data.SslCryptoCertificate.IsNull() {
		return false
	}
	if !data.IpAddress.IsNull() {
		return false
	}
	if !data.Peers.IsNull() {
		return false
	}
	if !data.Priority.IsNull() {
		return false
	}
	if !data.StrictMode.IsNull() {
		return false
	}
	return true
}
func (data *QuotaEnforcementServer) ToDefault() {
	data.Enabled = types.BoolValue(true)
	data.UserSummary = types.StringNull()
	data.PasswordAlias = types.StringNull()
	data.RaidVolume = types.StringNull()
	data.ServerPort = types.Int64Value(16379)
	data.MonitorPort = types.Int64Value(26379)
	data.EnablePeerGroup = types.BoolValue(false)
	data.EnableSsl = types.BoolValue(true)
	data.SslCryptoKey = types.StringNull()
	data.SslCryptoCertificate = types.StringNull()
	data.IpAddress = types.StringNull()
	data.Peers = types.ListNull(types.StringType)
	data.Priority = types.Int64Value(100)
	data.StrictMode = types.BoolValue(true)
}

func (data QuotaEnforcementServer) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "QuotaEnforcementServer.name", path.Base("/mgmt/config/default/QuotaEnforcementServer/QuotaEnforcementServer"))

	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.PasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PasswordAlias`, data.PasswordAlias.ValueString())
	}
	if !data.RaidVolume.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RaidVolume`, data.RaidVolume.ValueString())
	}
	if !data.ServerPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ServerPort`, data.ServerPort.ValueInt64())
	}
	if !data.MonitorPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MonitorPort`, data.MonitorPort.ValueInt64())
	}
	if !data.EnablePeerGroup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnablePeerGroup`, tfutils.StringFromBool(data.EnablePeerGroup, ""))
	}
	if !data.EnableSsl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnableSSL`, tfutils.StringFromBool(data.EnableSsl, ""))
	}
	if !data.SslCryptoKey.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLCryptoKey`, data.SslCryptoKey.ValueString())
	}
	if !data.SslCryptoCertificate.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLCryptoCertificate`, data.SslCryptoCertificate.ValueString())
	}
	if !data.IpAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IPAddress`, data.IpAddress.ValueString())
	}
	if !data.Peers.IsNull() {
		var dataValues []string
		data.Peers.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.Set(body, pathRoot+`Peers`+".-1", map[string]string{"value": val})
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`Peers`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`Peers`, "[]")
	}
	if !data.Priority.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Priority`, data.Priority.ValueInt64())
	}
	if !data.StrictMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StrictMode`, tfutils.StringFromBool(data.StrictMode, ""))
	}
	return body
}

func (data *QuotaEnforcementServer) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `PasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `RaidVolume`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RaidVolume = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RaidVolume = types.StringNull()
	}
	if value := res.Get(pathRoot + `ServerPort`); value.Exists() {
		data.ServerPort = types.Int64Value(value.Int())
	} else {
		data.ServerPort = types.Int64Value(16379)
	}
	if value := res.Get(pathRoot + `MonitorPort`); value.Exists() {
		data.MonitorPort = types.Int64Value(value.Int())
	} else {
		data.MonitorPort = types.Int64Value(26379)
	}
	if value := res.Get(pathRoot + `EnablePeerGroup`); value.Exists() {
		data.EnablePeerGroup = tfutils.BoolFromString(value.String())
	} else {
		data.EnablePeerGroup = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EnableSSL`); value.Exists() {
		data.EnableSsl = tfutils.BoolFromString(value.String())
	} else {
		data.EnableSsl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SSLCryptoKey`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslCryptoKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslCryptoKey = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLCryptoCertificate`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslCryptoCertificate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslCryptoCertificate = types.StringNull()
	}
	if value := res.Get(pathRoot + `IPAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.IpAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.IpAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `Peers`); value.Exists() {
		data.Peers = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Peers = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `Priority`); value.Exists() {
		data.Priority = types.Int64Value(value.Int())
	} else {
		data.Priority = types.Int64Value(100)
	}
	if value := res.Get(pathRoot + `StrictMode`); value.Exists() {
		data.StrictMode = tfutils.BoolFromString(value.String())
	} else {
		data.StrictMode = types.BoolNull()
	}
}

func (data *QuotaEnforcementServer) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `PasswordAlias`); value.Exists() && !data.PasswordAlias.IsNull() {
		data.PasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `RaidVolume`); value.Exists() && !data.RaidVolume.IsNull() {
		data.RaidVolume = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RaidVolume = types.StringNull()
	}
	if value := res.Get(pathRoot + `ServerPort`); value.Exists() && !data.ServerPort.IsNull() {
		data.ServerPort = types.Int64Value(value.Int())
	} else if data.ServerPort.ValueInt64() != 16379 {
		data.ServerPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MonitorPort`); value.Exists() && !data.MonitorPort.IsNull() {
		data.MonitorPort = types.Int64Value(value.Int())
	} else if data.MonitorPort.ValueInt64() != 26379 {
		data.MonitorPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `EnablePeerGroup`); value.Exists() && !data.EnablePeerGroup.IsNull() {
		data.EnablePeerGroup = tfutils.BoolFromString(value.String())
	} else if data.EnablePeerGroup.ValueBool() {
		data.EnablePeerGroup = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EnableSSL`); value.Exists() && !data.EnableSsl.IsNull() {
		data.EnableSsl = tfutils.BoolFromString(value.String())
	} else if !data.EnableSsl.ValueBool() {
		data.EnableSsl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SSLCryptoKey`); value.Exists() && !data.SslCryptoKey.IsNull() {
		data.SslCryptoKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslCryptoKey = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLCryptoCertificate`); value.Exists() && !data.SslCryptoCertificate.IsNull() {
		data.SslCryptoCertificate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslCryptoCertificate = types.StringNull()
	}
	if value := res.Get(pathRoot + `IPAddress`); value.Exists() && !data.IpAddress.IsNull() {
		data.IpAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.IpAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `Peers`); value.Exists() && !data.Peers.IsNull() {
		data.Peers = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Peers = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `Priority`); value.Exists() && !data.Priority.IsNull() {
		data.Priority = types.Int64Value(value.Int())
	} else if data.Priority.ValueInt64() != 100 {
		data.Priority = types.Int64Null()
	}
	if value := res.Get(pathRoot + `StrictMode`); value.Exists() && !data.StrictMode.IsNull() {
		data.StrictMode = tfutils.BoolFromString(value.String())
	} else if !data.StrictMode.ValueBool() {
		data.StrictMode = types.BoolNull()
	}
}
