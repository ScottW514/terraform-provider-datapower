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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type TFIMEndpoint struct {
	Id                   types.String `tfsdk:"id"`
	AppDomain            types.String `tfsdk:"app_domain"`
	UserSummary          types.String `tfsdk:"user_summary"`
	MEndpointType        types.String `tfsdk:"m_endpoint_type"`
	MServerUrl           types.String `tfsdk:"m_server_url"`
	MServerPort          types.Int64  `tfsdk:"m_server_port"`
	MCompatibleMode      types.String `tfsdk:"m_compatible_mode"`
	MReqToken60Format    types.String `tfsdk:"m_req_token60_format"`
	MReqToken61Format    types.String `tfsdk:"m_req_token61_format"`
	MReqToken62Format    types.String `tfsdk:"m_req_token62_format"`
	MReqCustomUrl        types.String `tfsdk:"m_req_custom_url"`
	MAppliesToAddress    types.String `tfsdk:"m_applies_to_address"`
	MIssuer              types.String `tfsdk:"m_issuer"`
	MPortType            types.String `tfsdk:"m_port_type"`
	MOperation           types.String `tfsdk:"m_operation"`
	MSchemaValidate      types.Bool   `tfsdk:"m_schema_validate"`
	MLtpaValueTypeMode   types.String `tfsdk:"m_ltpa_value_type_mode"`
	MStsUsername         types.String `tfsdk:"m_sts_username"`
	MStsPassword         types.String `tfsdk:"m_sts_password"`
	MStsPasswordAlias    types.String `tfsdk:"m_sts_password_alias"`
	MSslClientConfigType types.String `tfsdk:"m_ssl_client_config_type"`
	MSslClient           types.String `tfsdk:"m_ssl_client"`
}

var TFIMEndpointObjectType = map[string]attr.Type{
	"id":                       types.StringType,
	"app_domain":               types.StringType,
	"user_summary":             types.StringType,
	"m_endpoint_type":          types.StringType,
	"m_server_url":             types.StringType,
	"m_server_port":            types.Int64Type,
	"m_compatible_mode":        types.StringType,
	"m_req_token60_format":     types.StringType,
	"m_req_token61_format":     types.StringType,
	"m_req_token62_format":     types.StringType,
	"m_req_custom_url":         types.StringType,
	"m_applies_to_address":     types.StringType,
	"m_issuer":                 types.StringType,
	"m_port_type":              types.StringType,
	"m_operation":              types.StringType,
	"m_schema_validate":        types.BoolType,
	"m_ltpa_value_type_mode":   types.StringType,
	"m_sts_username":           types.StringType,
	"m_sts_password":           types.StringType,
	"m_sts_password_alias":     types.StringType,
	"m_ssl_client_config_type": types.StringType,
	"m_ssl_client":             types.StringType,
}

func (data TFIMEndpoint) GetPath() string {
	rest_path := "/mgmt/config/{domain}/TFIMEndpoint"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data TFIMEndpoint) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.MEndpointType.IsNull() {
		return false
	}
	if !data.MServerUrl.IsNull() {
		return false
	}
	if !data.MServerPort.IsNull() {
		return false
	}
	if !data.MCompatibleMode.IsNull() {
		return false
	}
	if !data.MReqToken60Format.IsNull() {
		return false
	}
	if !data.MReqToken61Format.IsNull() {
		return false
	}
	if !data.MReqToken62Format.IsNull() {
		return false
	}
	if !data.MReqCustomUrl.IsNull() {
		return false
	}
	if !data.MAppliesToAddress.IsNull() {
		return false
	}
	if !data.MIssuer.IsNull() {
		return false
	}
	if !data.MPortType.IsNull() {
		return false
	}
	if !data.MOperation.IsNull() {
		return false
	}
	if !data.MSchemaValidate.IsNull() {
		return false
	}
	if !data.MLtpaValueTypeMode.IsNull() {
		return false
	}
	if !data.MStsUsername.IsNull() {
		return false
	}
	if !data.MStsPassword.IsNull() {
		return false
	}
	if !data.MStsPasswordAlias.IsNull() {
		return false
	}
	if !data.MSslClientConfigType.IsNull() {
		return false
	}
	if !data.MSslClient.IsNull() {
		return false
	}
	return true
}

func (data TFIMEndpoint) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.MEndpointType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mEndpointType`, data.MEndpointType.ValueString())
	}
	if !data.MServerUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mServerURL`, data.MServerUrl.ValueString())
	}
	if !data.MServerPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mServerPort`, data.MServerPort.ValueInt64())
	}
	if !data.MCompatibleMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mCompatibleMode`, data.MCompatibleMode.ValueString())
	}
	if !data.MReqToken60Format.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mReqToken60Format`, data.MReqToken60Format.ValueString())
	}
	if !data.MReqToken61Format.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mReqToken61Format`, data.MReqToken61Format.ValueString())
	}
	if !data.MReqToken62Format.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mReqToken62Format`, data.MReqToken62Format.ValueString())
	}
	if !data.MReqCustomUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mReqCustomURL`, data.MReqCustomUrl.ValueString())
	}
	if !data.MAppliesToAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAppliesToAddress`, data.MAppliesToAddress.ValueString())
	}
	if !data.MIssuer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mIssuer`, data.MIssuer.ValueString())
	}
	if !data.MPortType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mPortType`, data.MPortType.ValueString())
	}
	if !data.MOperation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mOperation`, data.MOperation.ValueString())
	}
	if !data.MSchemaValidate.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mSchemaValidate`, tfutils.StringFromBool(data.MSchemaValidate))
	}
	if !data.MLtpaValueTypeMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mLTPAValueTypeMode`, data.MLtpaValueTypeMode.ValueString())
	}
	if !data.MStsUsername.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mSTSUsername`, data.MStsUsername.ValueString())
	}
	if !data.MStsPassword.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mSTSPassword`, data.MStsPassword.ValueString())
	}
	if !data.MStsPasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mSTSPasswordAlias`, data.MStsPasswordAlias.ValueString())
	}
	if !data.MSslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mSSLClientConfigType`, data.MSslClientConfigType.ValueString())
	}
	if !data.MSslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mSSLClient`, data.MSslClient.ValueString())
	}
	return body
}

func (data *TFIMEndpoint) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `mEndpointType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MEndpointType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MEndpointType = types.StringValue("tokenmapping")
	}
	if value := res.Get(pathRoot + `mServerURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MServerUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MServerUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `mServerPort`); value.Exists() {
		data.MServerPort = types.Int64Value(value.Int())
	} else {
		data.MServerPort = types.Int64Value(9080)
	}
	if value := res.Get(pathRoot + `mCompatibleMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MCompatibleMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MCompatibleMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `mReqToken60Format`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MReqToken60Format = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MReqToken60Format = types.StringNull()
	}
	if value := res.Get(pathRoot + `mReqToken61Format`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MReqToken61Format = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MReqToken61Format = types.StringNull()
	}
	if value := res.Get(pathRoot + `mReqToken62Format`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MReqToken62Format = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MReqToken62Format = types.StringNull()
	}
	if value := res.Get(pathRoot + `mReqCustomURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MReqCustomUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MReqCustomUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `mAppliesToAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MAppliesToAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MAppliesToAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `mIssuer`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MIssuer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MIssuer = types.StringNull()
	}
	if value := res.Get(pathRoot + `mPortType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MPortType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MPortType = types.StringNull()
	}
	if value := res.Get(pathRoot + `mOperation`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MOperation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MOperation = types.StringNull()
	}
	if value := res.Get(pathRoot + `mSchemaValidate`); value.Exists() {
		data.MSchemaValidate = tfutils.BoolFromString(value.String())
	} else {
		data.MSchemaValidate = types.BoolNull()
	}
	if value := res.Get(pathRoot + `mLTPAValueTypeMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MLtpaValueTypeMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MLtpaValueTypeMode = types.StringValue("static")
	}
	if value := res.Get(pathRoot + `mSTSUsername`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MStsUsername = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MStsUsername = types.StringNull()
	}
	if value := res.Get(pathRoot + `mSTSPassword`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MStsPassword = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MStsPassword = types.StringNull()
	}
	if value := res.Get(pathRoot + `mSTSPasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MStsPasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MStsPasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `mSSLClientConfigType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MSslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MSslClientConfigType = types.StringValue("client")
	}
	if value := res.Get(pathRoot + `mSSLClient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MSslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MSslClient = types.StringNull()
	}
}

func (data *TFIMEndpoint) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `mEndpointType`); value.Exists() && !data.MEndpointType.IsNull() {
		data.MEndpointType = tfutils.ParseStringFromGJSON(value)
	} else if data.MEndpointType.ValueString() != "tokenmapping" {
		data.MEndpointType = types.StringNull()
	}
	if value := res.Get(pathRoot + `mServerURL`); value.Exists() && !data.MServerUrl.IsNull() {
		data.MServerUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MServerUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `mServerPort`); value.Exists() && !data.MServerPort.IsNull() {
		data.MServerPort = types.Int64Value(value.Int())
	} else if data.MServerPort.ValueInt64() != 9080 {
		data.MServerPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `mCompatibleMode`); value.Exists() && !data.MCompatibleMode.IsNull() {
		data.MCompatibleMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MCompatibleMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `mReqToken60Format`); value.Exists() && !data.MReqToken60Format.IsNull() {
		data.MReqToken60Format = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MReqToken60Format = types.StringNull()
	}
	if value := res.Get(pathRoot + `mReqToken61Format`); value.Exists() && !data.MReqToken61Format.IsNull() {
		data.MReqToken61Format = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MReqToken61Format = types.StringNull()
	}
	if value := res.Get(pathRoot + `mReqToken62Format`); value.Exists() && !data.MReqToken62Format.IsNull() {
		data.MReqToken62Format = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MReqToken62Format = types.StringNull()
	}
	if value := res.Get(pathRoot + `mReqCustomURL`); value.Exists() && !data.MReqCustomUrl.IsNull() {
		data.MReqCustomUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MReqCustomUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `mAppliesToAddress`); value.Exists() && !data.MAppliesToAddress.IsNull() {
		data.MAppliesToAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MAppliesToAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `mIssuer`); value.Exists() && !data.MIssuer.IsNull() {
		data.MIssuer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MIssuer = types.StringNull()
	}
	if value := res.Get(pathRoot + `mPortType`); value.Exists() && !data.MPortType.IsNull() {
		data.MPortType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MPortType = types.StringNull()
	}
	if value := res.Get(pathRoot + `mOperation`); value.Exists() && !data.MOperation.IsNull() {
		data.MOperation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MOperation = types.StringNull()
	}
	if value := res.Get(pathRoot + `mSchemaValidate`); value.Exists() && !data.MSchemaValidate.IsNull() {
		data.MSchemaValidate = tfutils.BoolFromString(value.String())
	} else if data.MSchemaValidate.ValueBool() {
		data.MSchemaValidate = types.BoolNull()
	}
	if value := res.Get(pathRoot + `mLTPAValueTypeMode`); value.Exists() && !data.MLtpaValueTypeMode.IsNull() {
		data.MLtpaValueTypeMode = tfutils.ParseStringFromGJSON(value)
	} else if data.MLtpaValueTypeMode.ValueString() != "static" {
		data.MLtpaValueTypeMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `mSTSUsername`); value.Exists() && !data.MStsUsername.IsNull() {
		data.MStsUsername = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MStsUsername = types.StringNull()
	}
	if value := res.Get(pathRoot + `mSTSPassword`); value.Exists() && !data.MStsPassword.IsNull() {
		data.MStsPassword = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MStsPassword = types.StringNull()
	}
	if value := res.Get(pathRoot + `mSTSPasswordAlias`); value.Exists() && !data.MStsPasswordAlias.IsNull() {
		data.MStsPasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MStsPasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `mSSLClientConfigType`); value.Exists() && !data.MSslClientConfigType.IsNull() {
		data.MSslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else if data.MSslClientConfigType.ValueString() != "client" {
		data.MSslClientConfigType = types.StringNull()
	}
	if value := res.Get(pathRoot + `mSSLClient`); value.Exists() && !data.MSslClient.IsNull() {
		data.MSslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MSslClient = types.StringNull()
	}
}
