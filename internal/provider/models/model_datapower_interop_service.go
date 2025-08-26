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

type InteropService struct {
	Enabled             types.Bool                  `tfsdk:"enabled"`
	UserSummary         types.String                `tfsdk:"user_summary"`
	XmlManager          types.String                `tfsdk:"xml_manager"`
	AaaPolicy           types.String                `tfsdk:"aaa_policy"`
	HttpService         types.Bool                  `tfsdk:"http_service"`
	LocalAddress        types.String                `tfsdk:"local_address"`
	LocalPort           types.Int64                 `tfsdk:"local_port"`
	Acl                 types.String                `tfsdk:"acl"`
	HttpsService        types.Bool                  `tfsdk:"https_service"`
	HttpsLocalAddress   types.String                `tfsdk:"https_local_address"`
	HttpsLocalPort      types.Int64                 `tfsdk:"https_local_port"`
	HttpsAcl            types.String                `tfsdk:"https_acl"`
	SslServerConfigType types.String                `tfsdk:"ssl_server_config_type"`
	SslServer           types.String                `tfsdk:"ssl_server"`
	SslSniServer        types.String                `tfsdk:"ssl_sni_server"`
	DependencyActions   []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var InteropServiceLocalAddressCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "http_service",
	AttrType:    "Bool",
	AttrDefault: "true",
	Value:       []string{"true"},
}
var InteropServiceLocalPortCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "http_service",
	AttrType:    "Bool",
	AttrDefault: "true",
	Value:       []string{"true"},
}
var InteropServiceHttpsLocalAddressCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "https_service",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var InteropServiceHttpsLocalPortCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "https_service",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var InteropServiceSSLServerConfigTypeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "https_service",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var InteropServiceSSLServerCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "https_service",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ssl_server_config_type",
			AttrType:    "String",
			AttrDefault: "server",
			Value:       []string{"server"},
		},
	},
}
var InteropServiceSSLSNIServerCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "https_service",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ssl_server_config_type",
			AttrType:    "String",
			AttrDefault: "server",
			Value:       []string{"sni"},
		},
	},
}

var InteropServiceObjectType = map[string]attr.Type{
	"enabled":                types.BoolType,
	"user_summary":           types.StringType,
	"xml_manager":            types.StringType,
	"aaa_policy":             types.StringType,
	"http_service":           types.BoolType,
	"local_address":          types.StringType,
	"local_port":             types.Int64Type,
	"acl":                    types.StringType,
	"https_service":          types.BoolType,
	"https_local_address":    types.StringType,
	"https_local_port":       types.Int64Type,
	"https_acl":              types.StringType,
	"ssl_server_config_type": types.StringType,
	"ssl_server":             types.StringType,
	"ssl_sni_server":         types.StringType,
	"dependency_actions":     actions.ActionsListType,
}

func (data InteropService) GetPath() string {
	rest_path := "/mgmt/config/default/InteropService/IOP-Settings"
	return rest_path
}

func (data InteropService) IsNull() bool {
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.XmlManager.IsNull() {
		return false
	}
	if !data.AaaPolicy.IsNull() {
		return false
	}
	if !data.HttpService.IsNull() {
		return false
	}
	if !data.LocalAddress.IsNull() {
		return false
	}
	if !data.LocalPort.IsNull() {
		return false
	}
	if !data.Acl.IsNull() {
		return false
	}
	if !data.HttpsService.IsNull() {
		return false
	}
	if !data.HttpsLocalAddress.IsNull() {
		return false
	}
	if !data.HttpsLocalPort.IsNull() {
		return false
	}
	if !data.HttpsAcl.IsNull() {
		return false
	}
	if !data.SslServerConfigType.IsNull() {
		return false
	}
	if !data.SslServer.IsNull() {
		return false
	}
	if !data.SslSniServer.IsNull() {
		return false
	}
	return true
}

func (data InteropService) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "InteropService.name", path.Base("/mgmt/config/default/InteropService/IOP-Settings"))

	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.XmlManager.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XMLManager`, data.XmlManager.ValueString())
	}
	if !data.AaaPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AAAPolicy`, data.AaaPolicy.ValueString())
	}
	if !data.HttpService.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HttpService`, tfutils.StringFromBool(data.HttpService, ""))
	}
	if !data.LocalAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalAddress`, data.LocalAddress.ValueString())
	}
	if !data.LocalPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalPort`, data.LocalPort.ValueInt64())
	}
	if !data.Acl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ACL`, data.Acl.ValueString())
	}
	if !data.HttpsService.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HttpsService`, tfutils.StringFromBool(data.HttpsService, ""))
	}
	if !data.HttpsLocalAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HttpsLocalAddress`, data.HttpsLocalAddress.ValueString())
	}
	if !data.HttpsLocalPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HttpsLocalPort`, data.HttpsLocalPort.ValueInt64())
	}
	if !data.HttpsAcl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HttpsACL`, data.HttpsAcl.ValueString())
	}
	if !data.SslServerConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLServerConfigType`, data.SslServerConfigType.ValueString())
	}
	if !data.SslServer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLServer`, data.SslServer.ValueString())
	}
	if !data.SslSniServer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLSNIServer`, data.SslSniServer.ValueString())
	}
	return body
}

func (data *InteropService) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `XMLManager`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.XmlManager = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XmlManager = types.StringNull()
	}
	if value := res.Get(pathRoot + `AAAPolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AaaPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AaaPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `HttpService`); value.Exists() {
		data.HttpService = tfutils.BoolFromString(value.String())
	} else {
		data.HttpService = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalAddress = types.StringValue("0.0.0.0")
	}
	if value := res.Get(pathRoot + `LocalPort`); value.Exists() {
		data.LocalPort = types.Int64Value(value.Int())
	} else {
		data.LocalPort = types.Int64Value(9990)
	}
	if value := res.Get(pathRoot + `ACL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Acl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Acl = types.StringNull()
	}
	if value := res.Get(pathRoot + `HttpsService`); value.Exists() {
		data.HttpsService = tfutils.BoolFromString(value.String())
	} else {
		data.HttpsService = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HttpsLocalAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HttpsLocalAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HttpsLocalAddress = types.StringValue("0.0.0.0")
	}
	if value := res.Get(pathRoot + `HttpsLocalPort`); value.Exists() {
		data.HttpsLocalPort = types.Int64Value(value.Int())
	} else {
		data.HttpsLocalPort = types.Int64Value(9991)
	}
	if value := res.Get(pathRoot + `HttpsACL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HttpsAcl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HttpsAcl = types.StringNull()
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
		data.SslSniServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslSniServer = types.StringNull()
	}
}

func (data *InteropService) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `XMLManager`); value.Exists() && !data.XmlManager.IsNull() {
		data.XmlManager = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XmlManager = types.StringNull()
	}
	if value := res.Get(pathRoot + `AAAPolicy`); value.Exists() && !data.AaaPolicy.IsNull() {
		data.AaaPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AaaPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `HttpService`); value.Exists() && !data.HttpService.IsNull() {
		data.HttpService = tfutils.BoolFromString(value.String())
	} else if !data.HttpService.ValueBool() {
		data.HttpService = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && !data.LocalAddress.IsNull() {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else if data.LocalAddress.ValueString() != "0.0.0.0" {
		data.LocalAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalPort`); value.Exists() && !data.LocalPort.IsNull() {
		data.LocalPort = types.Int64Value(value.Int())
	} else if data.LocalPort.ValueInt64() != 9990 {
		data.LocalPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ACL`); value.Exists() && !data.Acl.IsNull() {
		data.Acl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Acl = types.StringNull()
	}
	if value := res.Get(pathRoot + `HttpsService`); value.Exists() && !data.HttpsService.IsNull() {
		data.HttpsService = tfutils.BoolFromString(value.String())
	} else if data.HttpsService.ValueBool() {
		data.HttpsService = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HttpsLocalAddress`); value.Exists() && !data.HttpsLocalAddress.IsNull() {
		data.HttpsLocalAddress = tfutils.ParseStringFromGJSON(value)
	} else if data.HttpsLocalAddress.ValueString() != "0.0.0.0" {
		data.HttpsLocalAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `HttpsLocalPort`); value.Exists() && !data.HttpsLocalPort.IsNull() {
		data.HttpsLocalPort = types.Int64Value(value.Int())
	} else if data.HttpsLocalPort.ValueInt64() != 9991 {
		data.HttpsLocalPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `HttpsACL`); value.Exists() && !data.HttpsAcl.IsNull() {
		data.HttpsAcl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HttpsAcl = types.StringNull()
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
	if value := res.Get(pathRoot + `SSLSNIServer`); value.Exists() && !data.SslSniServer.IsNull() {
		data.SslSniServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslSniServer = types.StringNull()
	}
}
