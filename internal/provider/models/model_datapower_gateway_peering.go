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

type GatewayPeering struct {
	Id                  types.String                `tfsdk:"id"`
	AppDomain           types.String                `tfsdk:"app_domain"`
	UserSummary         types.String                `tfsdk:"user_summary"`
	PasswordAlias       types.String                `tfsdk:"password_alias"`
	LocalAddress        types.String                `tfsdk:"local_address"`
	LocalPort           types.Int64                 `tfsdk:"local_port"`
	PeerGroup           types.String                `tfsdk:"peer_group"`
	MonitorPort         types.Int64                 `tfsdk:"monitor_port"`
	EnablePeerGroup     types.Bool                  `tfsdk:"enable_peer_group"`
	EnableSsl           types.Bool                  `tfsdk:"enable_ssl"`
	PersistenceLocation types.String                `tfsdk:"persistence_location"`
	LocalDirectory      types.String                `tfsdk:"local_directory"`
	MaxMemory           types.Int64                 `tfsdk:"max_memory"`
	DependencyActions   []*actions.DependencyAction `tfsdk:"dependency_actions"`
	ProviderTarget      types.String                `tfsdk:"provider_target"`
}

var GatewayPeeringLocalAddressCondVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "enable_peer_group",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "enable_peer_group",
					AttrType:    "Bool",
					AttrDefault: "true",
					Value:       []string{"false"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "peer_group",
					AttrType:    "String",
					AttrDefault: "",
					Value:       []string{""},
				},
			},
		},
	},
}

var GatewayPeeringPeerGroupIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "enable_peer_group",
	AttrType:    "Bool",
	AttrDefault: "true",
	Value:       []string{"true"},
}

var GatewayPeeringEnableSSLIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "peer_group",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{""},
}

var GatewayPeeringLocalDirectoryCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "persistence_location",
	AttrType:    "String",
	AttrDefault: "memory",
	Value:       []string{"local"},
}

var GatewayPeeringLocalDirectoryIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var GatewayPeeringMaxMemoryIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "persistence_location",
	AttrType:    "String",
	AttrDefault: "memory",
	Value:       []string{"memory"},
}

var GatewayPeeringObjectType = map[string]attr.Type{
	"provider_target":      types.StringType,
	"id":                   types.StringType,
	"app_domain":           types.StringType,
	"user_summary":         types.StringType,
	"password_alias":       types.StringType,
	"local_address":        types.StringType,
	"local_port":           types.Int64Type,
	"peer_group":           types.StringType,
	"monitor_port":         types.Int64Type,
	"enable_peer_group":    types.BoolType,
	"enable_ssl":           types.BoolType,
	"persistence_location": types.StringType,
	"local_directory":      types.StringType,
	"max_memory":           types.Int64Type,
	"dependency_actions":   actions.ActionsListType,
}

func (data GatewayPeering) GetPath() string {
	rest_path := "/mgmt/config/{domain}/GatewayPeering"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data GatewayPeering) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.PasswordAlias.IsNull() {
		return false
	}
	if !data.LocalAddress.IsNull() {
		return false
	}
	if !data.LocalPort.IsNull() {
		return false
	}
	if !data.PeerGroup.IsNull() {
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
	if !data.PersistenceLocation.IsNull() {
		return false
	}
	if !data.LocalDirectory.IsNull() {
		return false
	}
	if !data.MaxMemory.IsNull() {
		return false
	}
	return true
}

func (data GatewayPeering) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.PasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PasswordAlias`, data.PasswordAlias.ValueString())
	}
	if !data.LocalAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalAddress`, data.LocalAddress.ValueString())
	}
	if !data.LocalPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalPort`, data.LocalPort.ValueInt64())
	}
	if !data.PeerGroup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PeerGroup`, data.PeerGroup.ValueString())
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
	if !data.PersistenceLocation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PersistenceLocation`, data.PersistenceLocation.ValueString())
	}
	if !data.LocalDirectory.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalDirectory`, data.LocalDirectory.ValueString())
	}
	if !data.MaxMemory.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxMemory`, data.MaxMemory.ValueInt64())
	}
	return body
}

func (data *GatewayPeering) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `PasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalPort`); value.Exists() {
		data.LocalPort = types.Int64Value(value.Int())
	} else {
		data.LocalPort = types.Int64Value(16380)
	}
	if value := res.Get(pathRoot + `PeerGroup`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PeerGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PeerGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `MonitorPort`); value.Exists() {
		data.MonitorPort = types.Int64Value(value.Int())
	} else {
		data.MonitorPort = types.Int64Value(26380)
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
	if value := res.Get(pathRoot + `PersistenceLocation`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PersistenceLocation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PersistenceLocation = types.StringValue("memory")
	}
	if value := res.Get(pathRoot + `LocalDirectory`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalDirectory = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalDirectory = types.StringValue("local:///")
	}
	if value := res.Get(pathRoot + `MaxMemory`); value.Exists() {
		data.MaxMemory = types.Int64Value(value.Int())
	} else {
		data.MaxMemory = types.Int64Null()
	}
}

func (data *GatewayPeering) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `PasswordAlias`); value.Exists() && !data.PasswordAlias.IsNull() {
		data.PasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && !data.LocalAddress.IsNull() {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalPort`); value.Exists() && !data.LocalPort.IsNull() {
		data.LocalPort = types.Int64Value(value.Int())
	} else if data.LocalPort.ValueInt64() != 16380 {
		data.LocalPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `PeerGroup`); value.Exists() && !data.PeerGroup.IsNull() {
		data.PeerGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PeerGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `MonitorPort`); value.Exists() && !data.MonitorPort.IsNull() {
		data.MonitorPort = types.Int64Value(value.Int())
	} else if data.MonitorPort.ValueInt64() != 26380 {
		data.MonitorPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `EnablePeerGroup`); value.Exists() && !data.EnablePeerGroup.IsNull() {
		data.EnablePeerGroup = tfutils.BoolFromString(value.String())
	} else if !data.EnablePeerGroup.ValueBool() {
		data.EnablePeerGroup = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EnableSSL`); value.Exists() && !data.EnableSsl.IsNull() {
		data.EnableSsl = tfutils.BoolFromString(value.String())
	} else if data.EnableSsl.ValueBool() {
		data.EnableSsl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PersistenceLocation`); value.Exists() && !data.PersistenceLocation.IsNull() {
		data.PersistenceLocation = tfutils.ParseStringFromGJSON(value)
	} else if data.PersistenceLocation.ValueString() != "memory" {
		data.PersistenceLocation = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalDirectory`); value.Exists() && !data.LocalDirectory.IsNull() {
		data.LocalDirectory = tfutils.ParseStringFromGJSON(value)
	} else if data.LocalDirectory.ValueString() != "local:///" {
		data.LocalDirectory = types.StringNull()
	}
	if value := res.Get(pathRoot + `MaxMemory`); value.Exists() && !data.MaxMemory.IsNull() {
		data.MaxMemory = types.Int64Value(value.Int())
	} else {
		data.MaxMemory = types.Int64Null()
	}
}
