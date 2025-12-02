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

type MgmtInterface struct {
	Enabled           types.Bool                  `tfsdk:"enabled"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	LocalPort         types.Int64                 `tfsdk:"local_port"`
	UserAgent         types.String                `tfsdk:"user_agent"`
	Acl               types.String                `tfsdk:"acl"`
	SlmPeering        types.Int64                 `tfsdk:"slm_peering"`
	Mode              *DmXMLMgmtModes             `tfsdk:"mode"`
	SslConfigType     types.String                `tfsdk:"ssl_config_type"`
	SslServer         types.String                `tfsdk:"ssl_server"`
	SslSniServer      types.String                `tfsdk:"ssl_sni_server"`
	LocalAddress      types.String                `tfsdk:"local_address"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
	ProviderTarget    types.String                `tfsdk:"provider_target"`
}

var MgmtInterfaceSLMPeeringCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "mode",
	AttrType:    "DmXMLMgmtModes",
	AttrDefault: "",
	Value:       []string{"slm"},
}

var MgmtInterfaceSLMPeeringIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var MgmtInterfaceSSLServerIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "ssl_config_type",
	AttrType:    "String",
	AttrDefault: "server",
	Value:       []string{"server"},
}

var MgmtInterfaceSSLSNIServerCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "ssl_config_type",
	AttrType:    "String",
	AttrDefault: "server",
	Value:       []string{"sni"},
}

var MgmtInterfaceSSLSNIServerIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var MgmtInterfaceObjectType = map[string]attr.Type{
	"provider_target":    types.StringType,
	"enabled":            types.BoolType,
	"user_summary":       types.StringType,
	"local_port":         types.Int64Type,
	"user_agent":         types.StringType,
	"acl":                types.StringType,
	"slm_peering":        types.Int64Type,
	"mode":               types.ObjectType{AttrTypes: DmXMLMgmtModesObjectType},
	"ssl_config_type":    types.StringType,
	"ssl_server":         types.StringType,
	"ssl_sni_server":     types.StringType,
	"local_address":      types.StringType,
	"dependency_actions": actions.ActionsListType,
}

func (data MgmtInterface) GetPath() string {
	rest_path := "/mgmt/config/default/MgmtInterface/XMLMgmt-Settings"
	return rest_path
}

func (data MgmtInterface) IsNull() bool {
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
	if !data.Acl.IsNull() {
		return false
	}
	if !data.SlmPeering.IsNull() {
		return false
	}
	if data.Mode != nil {
		if !data.Mode.IsNull() {
			return false
		}
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
	if !data.LocalAddress.IsNull() {
		return false
	}
	return true
}
func (data *MgmtInterface) ToDefault() {
	data.Enabled = types.BoolValue(false)
	data.UserSummary = types.StringNull()
	data.LocalPort = types.Int64Value(5550)
	data.UserAgent = types.StringNull()
	data.Acl = types.StringValue("xml-mgmt")
	data.SlmPeering = types.Int64Value(10)
	data.Mode = &DmXMLMgmtModes{}
	data.Mode.ToDefault()
	data.SslConfigType = types.StringValue("server")
	data.SslServer = types.StringNull()
	data.SslSniServer = types.StringNull()
	data.LocalAddress = types.StringValue("0.0.0.0")
}

func (data MgmtInterface) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "MgmtInterface.name", path.Base("/mgmt/config/default/MgmtInterface/XMLMgmt-Settings"))

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
	if !data.Acl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ACL`, data.Acl.ValueString())
	}
	if !data.SlmPeering.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SLMPeering`, data.SlmPeering.ValueInt64())
	}
	if data.Mode != nil {
		if !data.Mode.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`Mode`, data.Mode.ToBody(ctx, ""))
		}
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
	if !data.LocalAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalAddress`, data.LocalAddress.ValueString())
	}
	return body
}

func (data *MgmtInterface) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
		data.LocalPort = types.Int64Value(5550)
	}
	if value := res.Get(pathRoot + `UserAgent`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserAgent = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserAgent = types.StringNull()
	}
	if value := res.Get(pathRoot + `ACL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Acl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Acl = types.StringValue("xml-mgmt")
	}
	if value := res.Get(pathRoot + `SLMPeering`); value.Exists() {
		data.SlmPeering = types.Int64Value(value.Int())
	} else {
		data.SlmPeering = types.Int64Value(10)
	}
	if value := res.Get(pathRoot + `Mode`); value.Exists() {
		data.Mode = &DmXMLMgmtModes{}
		data.Mode.FromBody(ctx, "", value)
	} else {
		data.Mode = nil
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
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalAddress = types.StringValue("0.0.0.0")
	}
}

func (data *MgmtInterface) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else if data.Enabled.ValueBool() {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalPort`); value.Exists() && !data.LocalPort.IsNull() {
		data.LocalPort = types.Int64Value(value.Int())
	} else if data.LocalPort.ValueInt64() != 5550 {
		data.LocalPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `UserAgent`); value.Exists() && !data.UserAgent.IsNull() {
		data.UserAgent = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserAgent = types.StringNull()
	}
	if value := res.Get(pathRoot + `ACL`); value.Exists() && !data.Acl.IsNull() {
		data.Acl = tfutils.ParseStringFromGJSON(value)
	} else if data.Acl.ValueString() != "xml-mgmt" {
		data.Acl = types.StringNull()
	}
	if value := res.Get(pathRoot + `SLMPeering`); value.Exists() && !data.SlmPeering.IsNull() {
		data.SlmPeering = types.Int64Value(value.Int())
	} else if data.SlmPeering.ValueInt64() != 10 {
		data.SlmPeering = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Mode`); value.Exists() {
		data.Mode.UpdateFromBody(ctx, "", value)
	} else {
		data.Mode = nil
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
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && !data.LocalAddress.IsNull() {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else if data.LocalAddress.ValueString() != "0.0.0.0" {
		data.LocalAddress = types.StringNull()
	}
}
