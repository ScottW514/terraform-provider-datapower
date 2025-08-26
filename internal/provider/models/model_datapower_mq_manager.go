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

type MQManager struct {
	Id                     types.String                `tfsdk:"id"`
	AppDomain              types.String                `tfsdk:"app_domain"`
	UserSummary            types.String                `tfsdk:"user_summary"`
	HostName               types.String                `tfsdk:"host_name"`
	QMname                 types.String                `tfsdk:"q_mname"`
	Ccsid                  types.Int64                 `tfsdk:"ccsid"`
	ChannelName            types.String                `tfsdk:"channel_name"`
	CspUserId              types.String                `tfsdk:"csp_user_id"`
	CspPasswordAlias       types.String                `tfsdk:"csp_password_alias"`
	Heartbeat              types.Int64                 `tfsdk:"heartbeat"`
	MaximumMessageSize     types.Int64                 `tfsdk:"maximum_message_size"`
	CacheTimeout           types.Int64                 `tfsdk:"cache_timeout"`
	FfstSize               types.Int64                 `tfsdk:"ffst_size"`
	FfstRotate             types.Int64                 `tfsdk:"ffst_rotate"`
	UnitsOfWork            types.Int64                 `tfsdk:"units_of_work"`
	AutomaticBackout       types.Bool                  `tfsdk:"automatic_backout"`
	BackoutThreshold       types.Int64                 `tfsdk:"backout_threshold"`
	BackoutQueueName       types.String                `tfsdk:"backout_queue_name"`
	TotalConnectionLimit   types.Int64                 `tfsdk:"total_connection_limit"`
	InitialConnections     types.Int64                 `tfsdk:"initial_connections"`
	SharingConversations   types.Int64                 `tfsdk:"sharing_conversations"`
	SslKey                 types.String                `tfsdk:"ssl_key"`
	PermitInsecureServers  types.Bool                  `tfsdk:"permit_insecure_servers"`
	SslCipher              types.String                `tfsdk:"ssl_cipher"`
	SslCertLabel           types.String                `tfsdk:"ssl_cert_label"`
	ConvertInput           types.Bool                  `tfsdk:"convert_input"`
	AutoRetry              types.Bool                  `tfsdk:"auto_retry"`
	RetryInterval          types.Int64                 `tfsdk:"retry_interval"`
	RetryAttempts          types.Int64                 `tfsdk:"retry_attempts"`
	LongRetryInterval      types.Int64                 `tfsdk:"long_retry_interval"`
	ReportingInterval      types.Int64                 `tfsdk:"reporting_interval"`
	AlternateUser          types.Bool                  `tfsdk:"alternate_user"`
	LocalAddress           types.String                `tfsdk:"local_address"`
	XmlManager             types.String                `tfsdk:"xml_manager"`
	SslClient              types.String                `tfsdk:"ssl_client"`
	OutboundSni            types.String                `tfsdk:"outbound_sni"`
	OcspCheckExtensions    types.Bool                  `tfsdk:"ocsp_check_extensions"`
	OcspAuthentication     types.String                `tfsdk:"ocsp_authentication"`
	CdpCheckExtensions     types.Bool                  `tfsdk:"cdp_check_extensions"`
	ClientRevocationChecks types.String                `tfsdk:"client_revocation_checks"`
	DependencyActions      []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var MQManagerCSPUserIdCondVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "csp_password_alias",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{""},
}
var MQManagerCSPPasswordAliasCondVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "csp_user_id",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{""},
}

var MQManagerObjectType = map[string]attr.Type{
	"id":                       types.StringType,
	"app_domain":               types.StringType,
	"user_summary":             types.StringType,
	"host_name":                types.StringType,
	"q_mname":                  types.StringType,
	"ccsid":                    types.Int64Type,
	"channel_name":             types.StringType,
	"csp_user_id":              types.StringType,
	"csp_password_alias":       types.StringType,
	"heartbeat":                types.Int64Type,
	"maximum_message_size":     types.Int64Type,
	"cache_timeout":            types.Int64Type,
	"ffst_size":                types.Int64Type,
	"ffst_rotate":              types.Int64Type,
	"units_of_work":            types.Int64Type,
	"automatic_backout":        types.BoolType,
	"backout_threshold":        types.Int64Type,
	"backout_queue_name":       types.StringType,
	"total_connection_limit":   types.Int64Type,
	"initial_connections":      types.Int64Type,
	"sharing_conversations":    types.Int64Type,
	"ssl_key":                  types.StringType,
	"permit_insecure_servers":  types.BoolType,
	"ssl_cipher":               types.StringType,
	"ssl_cert_label":           types.StringType,
	"convert_input":            types.BoolType,
	"auto_retry":               types.BoolType,
	"retry_interval":           types.Int64Type,
	"retry_attempts":           types.Int64Type,
	"long_retry_interval":      types.Int64Type,
	"reporting_interval":       types.Int64Type,
	"alternate_user":           types.BoolType,
	"local_address":            types.StringType,
	"xml_manager":              types.StringType,
	"ssl_client":               types.StringType,
	"outbound_sni":             types.StringType,
	"ocsp_check_extensions":    types.BoolType,
	"ocsp_authentication":      types.StringType,
	"cdp_check_extensions":     types.BoolType,
	"client_revocation_checks": types.StringType,
	"dependency_actions":       actions.ActionsListType,
}

func (data MQManager) GetPath() string {
	rest_path := "/mgmt/config/{domain}/MQManager"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data MQManager) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.HostName.IsNull() {
		return false
	}
	if !data.QMname.IsNull() {
		return false
	}
	if !data.Ccsid.IsNull() {
		return false
	}
	if !data.ChannelName.IsNull() {
		return false
	}
	if !data.CspUserId.IsNull() {
		return false
	}
	if !data.CspPasswordAlias.IsNull() {
		return false
	}
	if !data.Heartbeat.IsNull() {
		return false
	}
	if !data.MaximumMessageSize.IsNull() {
		return false
	}
	if !data.CacheTimeout.IsNull() {
		return false
	}
	if !data.FfstSize.IsNull() {
		return false
	}
	if !data.FfstRotate.IsNull() {
		return false
	}
	if !data.UnitsOfWork.IsNull() {
		return false
	}
	if !data.AutomaticBackout.IsNull() {
		return false
	}
	if !data.BackoutThreshold.IsNull() {
		return false
	}
	if !data.BackoutQueueName.IsNull() {
		return false
	}
	if !data.TotalConnectionLimit.IsNull() {
		return false
	}
	if !data.InitialConnections.IsNull() {
		return false
	}
	if !data.SharingConversations.IsNull() {
		return false
	}
	if !data.SslKey.IsNull() {
		return false
	}
	if !data.PermitInsecureServers.IsNull() {
		return false
	}
	if !data.SslCipher.IsNull() {
		return false
	}
	if !data.SslCertLabel.IsNull() {
		return false
	}
	if !data.ConvertInput.IsNull() {
		return false
	}
	if !data.AutoRetry.IsNull() {
		return false
	}
	if !data.RetryInterval.IsNull() {
		return false
	}
	if !data.RetryAttempts.IsNull() {
		return false
	}
	if !data.LongRetryInterval.IsNull() {
		return false
	}
	if !data.ReportingInterval.IsNull() {
		return false
	}
	if !data.AlternateUser.IsNull() {
		return false
	}
	if !data.LocalAddress.IsNull() {
		return false
	}
	if !data.XmlManager.IsNull() {
		return false
	}
	if !data.SslClient.IsNull() {
		return false
	}
	if !data.OutboundSni.IsNull() {
		return false
	}
	if !data.OcspCheckExtensions.IsNull() {
		return false
	}
	if !data.OcspAuthentication.IsNull() {
		return false
	}
	if !data.CdpCheckExtensions.IsNull() {
		return false
	}
	if !data.ClientRevocationChecks.IsNull() {
		return false
	}
	return true
}

func (data MQManager) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.HostName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HostName`, data.HostName.ValueString())
	}
	if !data.QMname.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`QMname`, data.QMname.ValueString())
	}
	if !data.Ccsid.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CCSID`, data.Ccsid.ValueInt64())
	}
	if !data.ChannelName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ChannelName`, data.ChannelName.ValueString())
	}
	if !data.CspUserId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CSPUserId`, data.CspUserId.ValueString())
	}
	if !data.CspPasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CSPPasswordAlias`, data.CspPasswordAlias.ValueString())
	}
	if !data.Heartbeat.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Heartbeat`, data.Heartbeat.ValueInt64())
	}
	if !data.MaximumMessageSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaximumMessageSize`, data.MaximumMessageSize.ValueInt64())
	}
	if !data.CacheTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CacheTimeout`, data.CacheTimeout.ValueInt64())
	}
	if !data.FfstSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FFSTSize`, data.FfstSize.ValueInt64())
	}
	if !data.FfstRotate.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FFSTRotate`, data.FfstRotate.ValueInt64())
	}
	if !data.UnitsOfWork.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UnitsOfWork`, data.UnitsOfWork.ValueInt64())
	}
	if !data.AutomaticBackout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AutomaticBackout`, tfutils.StringFromBool(data.AutomaticBackout, ""))
	}
	if !data.BackoutThreshold.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BackoutThreshold`, data.BackoutThreshold.ValueInt64())
	}
	if !data.BackoutQueueName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BackoutQueueName`, data.BackoutQueueName.ValueString())
	}
	if !data.TotalConnectionLimit.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TotalConnectionLimit`, data.TotalConnectionLimit.ValueInt64())
	}
	if !data.InitialConnections.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`InitialConnections`, data.InitialConnections.ValueInt64())
	}
	if !data.SharingConversations.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SharingConversations`, data.SharingConversations.ValueInt64())
	}
	if !data.SslKey.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLkey`, data.SslKey.ValueString())
	}
	if !data.PermitInsecureServers.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PermitInsecureServers`, tfutils.StringFromBool(data.PermitInsecureServers, ""))
	}
	if !data.SslCipher.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLcipher`, data.SslCipher.ValueString())
	}
	if !data.SslCertLabel.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLCertLabel`, data.SslCertLabel.ValueString())
	}
	if !data.ConvertInput.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ConvertInput`, tfutils.StringFromBool(data.ConvertInput, ""))
	}
	if !data.AutoRetry.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AutoRetry`, tfutils.StringFromBool(data.AutoRetry, ""))
	}
	if !data.RetryInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RetryInterval`, data.RetryInterval.ValueInt64())
	}
	if !data.RetryAttempts.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RetryAttempts`, data.RetryAttempts.ValueInt64())
	}
	if !data.LongRetryInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LongRetryInterval`, data.LongRetryInterval.ValueInt64())
	}
	if !data.ReportingInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ReportingInterval`, data.ReportingInterval.ValueInt64())
	}
	if !data.AlternateUser.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AlternateUser`, tfutils.StringFromBool(data.AlternateUser, ""))
	}
	if !data.LocalAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalAddress`, data.LocalAddress.ValueString())
	}
	if !data.XmlManager.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XMLManager`, data.XmlManager.ValueString())
	}
	if !data.SslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClient`, data.SslClient.ValueString())
	}
	if !data.OutboundSni.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OutboundSNI`, data.OutboundSni.ValueString())
	}
	if !data.OcspCheckExtensions.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OCSPCheckExtensions`, tfutils.StringFromBool(data.OcspCheckExtensions, ""))
	}
	if !data.OcspAuthentication.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OCSPAuthentication`, data.OcspAuthentication.ValueString())
	}
	if !data.CdpCheckExtensions.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CDPCheckExtensions`, tfutils.StringFromBool(data.CdpCheckExtensions, ""))
	}
	if !data.ClientRevocationChecks.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ClientRevocationChecks`, data.ClientRevocationChecks.ValueString())
	}
	return body
}

func (data *MQManager) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `HostName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HostName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HostName = types.StringNull()
	}
	if value := res.Get(pathRoot + `QMname`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.QMname = tfutils.ParseStringFromGJSON(value)
	} else {
		data.QMname = types.StringNull()
	}
	if value := res.Get(pathRoot + `CCSID`); value.Exists() {
		data.Ccsid = types.Int64Value(value.Int())
	} else {
		data.Ccsid = types.Int64Value(819)
	}
	if value := res.Get(pathRoot + `ChannelName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ChannelName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ChannelName = types.StringValue("SYSTEM.DEF.SVRCONN")
	}
	if value := res.Get(pathRoot + `CSPUserId`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CspUserId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CspUserId = types.StringNull()
	}
	if value := res.Get(pathRoot + `CSPPasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CspPasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CspPasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `Heartbeat`); value.Exists() {
		data.Heartbeat = types.Int64Value(value.Int())
	} else {
		data.Heartbeat = types.Int64Value(300)
	}
	if value := res.Get(pathRoot + `MaximumMessageSize`); value.Exists() {
		data.MaximumMessageSize = types.Int64Value(value.Int())
	} else {
		data.MaximumMessageSize = types.Int64Value(1048576)
	}
	if value := res.Get(pathRoot + `CacheTimeout`); value.Exists() {
		data.CacheTimeout = types.Int64Value(value.Int())
	} else {
		data.CacheTimeout = types.Int64Value(60)
	}
	if value := res.Get(pathRoot + `FFSTSize`); value.Exists() {
		data.FfstSize = types.Int64Value(value.Int())
	} else {
		data.FfstSize = types.Int64Value(500)
	}
	if value := res.Get(pathRoot + `FFSTRotate`); value.Exists() {
		data.FfstRotate = types.Int64Value(value.Int())
	} else {
		data.FfstRotate = types.Int64Value(3)
	}
	if value := res.Get(pathRoot + `UnitsOfWork`); value.Exists() {
		data.UnitsOfWork = types.Int64Value(value.Int())
	} else {
		data.UnitsOfWork = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `AutomaticBackout`); value.Exists() {
		data.AutomaticBackout = tfutils.BoolFromString(value.String())
	} else {
		data.AutomaticBackout = types.BoolNull()
	}
	if value := res.Get(pathRoot + `BackoutThreshold`); value.Exists() {
		data.BackoutThreshold = types.Int64Value(value.Int())
	} else {
		data.BackoutThreshold = types.Int64Null()
	}
	if value := res.Get(pathRoot + `BackoutQueueName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.BackoutQueueName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BackoutQueueName = types.StringNull()
	}
	if value := res.Get(pathRoot + `TotalConnectionLimit`); value.Exists() {
		data.TotalConnectionLimit = types.Int64Value(value.Int())
	} else {
		data.TotalConnectionLimit = types.Int64Value(250)
	}
	if value := res.Get(pathRoot + `InitialConnections`); value.Exists() {
		data.InitialConnections = types.Int64Value(value.Int())
	} else {
		data.InitialConnections = types.Int64Value(1)
	}
	if value := res.Get(pathRoot + `SharingConversations`); value.Exists() {
		data.SharingConversations = types.Int64Value(value.Int())
	} else {
		data.SharingConversations = types.Int64Value(1)
	}
	if value := res.Get(pathRoot + `SSLkey`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslKey = types.StringNull()
	}
	if value := res.Get(pathRoot + `PermitInsecureServers`); value.Exists() {
		data.PermitInsecureServers = tfutils.BoolFromString(value.String())
	} else {
		data.PermitInsecureServers = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SSLcipher`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslCipher = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslCipher = types.StringValue("none")
	}
	if value := res.Get(pathRoot + `SSLCertLabel`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslCertLabel = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslCertLabel = types.StringNull()
	}
	if value := res.Get(pathRoot + `ConvertInput`); value.Exists() {
		data.ConvertInput = tfutils.BoolFromString(value.String())
	} else {
		data.ConvertInput = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AutoRetry`); value.Exists() {
		data.AutoRetry = tfutils.BoolFromString(value.String())
	} else {
		data.AutoRetry = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RetryInterval`); value.Exists() {
		data.RetryInterval = types.Int64Value(value.Int())
	} else {
		data.RetryInterval = types.Int64Value(10)
	}
	if value := res.Get(pathRoot + `RetryAttempts`); value.Exists() {
		data.RetryAttempts = types.Int64Value(value.Int())
	} else {
		data.RetryAttempts = types.Int64Value(6)
	}
	if value := res.Get(pathRoot + `LongRetryInterval`); value.Exists() {
		data.LongRetryInterval = types.Int64Value(value.Int())
	} else {
		data.LongRetryInterval = types.Int64Value(600)
	}
	if value := res.Get(pathRoot + `ReportingInterval`); value.Exists() {
		data.ReportingInterval = types.Int64Value(value.Int())
	} else {
		data.ReportingInterval = types.Int64Value(10)
	}
	if value := res.Get(pathRoot + `AlternateUser`); value.Exists() {
		data.AlternateUser = tfutils.BoolFromString(value.String())
	} else {
		data.AlternateUser = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `XMLManager`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.XmlManager = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XmlManager = types.StringValue("default")
	}
	if value := res.Get(pathRoot + `SSLClient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `OutboundSNI`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OutboundSni = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OutboundSni = types.StringValue("Channel")
	}
	if value := res.Get(pathRoot + `OCSPCheckExtensions`); value.Exists() {
		data.OcspCheckExtensions = tfutils.BoolFromString(value.String())
	} else {
		data.OcspCheckExtensions = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OCSPAuthentication`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OcspAuthentication = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OcspAuthentication = types.StringValue("REQUIRED")
	}
	if value := res.Get(pathRoot + `CDPCheckExtensions`); value.Exists() {
		data.CdpCheckExtensions = tfutils.BoolFromString(value.String())
	} else {
		data.CdpCheckExtensions = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ClientRevocationChecks`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ClientRevocationChecks = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientRevocationChecks = types.StringValue("REQUIRED")
	}
}

func (data *MQManager) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `HostName`); value.Exists() && !data.HostName.IsNull() {
		data.HostName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HostName = types.StringNull()
	}
	if value := res.Get(pathRoot + `QMname`); value.Exists() && !data.QMname.IsNull() {
		data.QMname = tfutils.ParseStringFromGJSON(value)
	} else {
		data.QMname = types.StringNull()
	}
	if value := res.Get(pathRoot + `CCSID`); value.Exists() && !data.Ccsid.IsNull() {
		data.Ccsid = types.Int64Value(value.Int())
	} else if data.Ccsid.ValueInt64() != 819 {
		data.Ccsid = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ChannelName`); value.Exists() && !data.ChannelName.IsNull() {
		data.ChannelName = tfutils.ParseStringFromGJSON(value)
	} else if data.ChannelName.ValueString() != "SYSTEM.DEF.SVRCONN" {
		data.ChannelName = types.StringNull()
	}
	if value := res.Get(pathRoot + `CSPUserId`); value.Exists() && !data.CspUserId.IsNull() {
		data.CspUserId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CspUserId = types.StringNull()
	}
	if value := res.Get(pathRoot + `CSPPasswordAlias`); value.Exists() && !data.CspPasswordAlias.IsNull() {
		data.CspPasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CspPasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `Heartbeat`); value.Exists() && !data.Heartbeat.IsNull() {
		data.Heartbeat = types.Int64Value(value.Int())
	} else if data.Heartbeat.ValueInt64() != 300 {
		data.Heartbeat = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaximumMessageSize`); value.Exists() && !data.MaximumMessageSize.IsNull() {
		data.MaximumMessageSize = types.Int64Value(value.Int())
	} else if data.MaximumMessageSize.ValueInt64() != 1048576 {
		data.MaximumMessageSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `CacheTimeout`); value.Exists() && !data.CacheTimeout.IsNull() {
		data.CacheTimeout = types.Int64Value(value.Int())
	} else if data.CacheTimeout.ValueInt64() != 60 {
		data.CacheTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `FFSTSize`); value.Exists() && !data.FfstSize.IsNull() {
		data.FfstSize = types.Int64Value(value.Int())
	} else if data.FfstSize.ValueInt64() != 500 {
		data.FfstSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `FFSTRotate`); value.Exists() && !data.FfstRotate.IsNull() {
		data.FfstRotate = types.Int64Value(value.Int())
	} else if data.FfstRotate.ValueInt64() != 3 {
		data.FfstRotate = types.Int64Null()
	}
	if value := res.Get(pathRoot + `UnitsOfWork`); value.Exists() && !data.UnitsOfWork.IsNull() {
		data.UnitsOfWork = types.Int64Value(value.Int())
	} else if data.UnitsOfWork.ValueInt64() != 0 {
		data.UnitsOfWork = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AutomaticBackout`); value.Exists() && !data.AutomaticBackout.IsNull() {
		data.AutomaticBackout = tfutils.BoolFromString(value.String())
	} else if data.AutomaticBackout.ValueBool() {
		data.AutomaticBackout = types.BoolNull()
	}
	if value := res.Get(pathRoot + `BackoutThreshold`); value.Exists() && !data.BackoutThreshold.IsNull() {
		data.BackoutThreshold = types.Int64Value(value.Int())
	} else {
		data.BackoutThreshold = types.Int64Null()
	}
	if value := res.Get(pathRoot + `BackoutQueueName`); value.Exists() && !data.BackoutQueueName.IsNull() {
		data.BackoutQueueName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BackoutQueueName = types.StringNull()
	}
	if value := res.Get(pathRoot + `TotalConnectionLimit`); value.Exists() && !data.TotalConnectionLimit.IsNull() {
		data.TotalConnectionLimit = types.Int64Value(value.Int())
	} else if data.TotalConnectionLimit.ValueInt64() != 250 {
		data.TotalConnectionLimit = types.Int64Null()
	}
	if value := res.Get(pathRoot + `InitialConnections`); value.Exists() && !data.InitialConnections.IsNull() {
		data.InitialConnections = types.Int64Value(value.Int())
	} else if data.InitialConnections.ValueInt64() != 1 {
		data.InitialConnections = types.Int64Null()
	}
	if value := res.Get(pathRoot + `SharingConversations`); value.Exists() && !data.SharingConversations.IsNull() {
		data.SharingConversations = types.Int64Value(value.Int())
	} else if data.SharingConversations.ValueInt64() != 1 {
		data.SharingConversations = types.Int64Null()
	}
	if value := res.Get(pathRoot + `SSLkey`); value.Exists() && !data.SslKey.IsNull() {
		data.SslKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslKey = types.StringNull()
	}
	if value := res.Get(pathRoot + `PermitInsecureServers`); value.Exists() && !data.PermitInsecureServers.IsNull() {
		data.PermitInsecureServers = tfutils.BoolFromString(value.String())
	} else if data.PermitInsecureServers.ValueBool() {
		data.PermitInsecureServers = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SSLcipher`); value.Exists() && !data.SslCipher.IsNull() {
		data.SslCipher = tfutils.ParseStringFromGJSON(value)
	} else if data.SslCipher.ValueString() != "none" {
		data.SslCipher = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLCertLabel`); value.Exists() && !data.SslCertLabel.IsNull() {
		data.SslCertLabel = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslCertLabel = types.StringNull()
	}
	if value := res.Get(pathRoot + `ConvertInput`); value.Exists() && !data.ConvertInput.IsNull() {
		data.ConvertInput = tfutils.BoolFromString(value.String())
	} else if !data.ConvertInput.ValueBool() {
		data.ConvertInput = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AutoRetry`); value.Exists() && !data.AutoRetry.IsNull() {
		data.AutoRetry = tfutils.BoolFromString(value.String())
	} else if !data.AutoRetry.ValueBool() {
		data.AutoRetry = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RetryInterval`); value.Exists() && !data.RetryInterval.IsNull() {
		data.RetryInterval = types.Int64Value(value.Int())
	} else if data.RetryInterval.ValueInt64() != 10 {
		data.RetryInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `RetryAttempts`); value.Exists() && !data.RetryAttempts.IsNull() {
		data.RetryAttempts = types.Int64Value(value.Int())
	} else if data.RetryAttempts.ValueInt64() != 6 {
		data.RetryAttempts = types.Int64Null()
	}
	if value := res.Get(pathRoot + `LongRetryInterval`); value.Exists() && !data.LongRetryInterval.IsNull() {
		data.LongRetryInterval = types.Int64Value(value.Int())
	} else if data.LongRetryInterval.ValueInt64() != 600 {
		data.LongRetryInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ReportingInterval`); value.Exists() && !data.ReportingInterval.IsNull() {
		data.ReportingInterval = types.Int64Value(value.Int())
	} else if data.ReportingInterval.ValueInt64() != 10 {
		data.ReportingInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AlternateUser`); value.Exists() && !data.AlternateUser.IsNull() {
		data.AlternateUser = tfutils.BoolFromString(value.String())
	} else if !data.AlternateUser.ValueBool() {
		data.AlternateUser = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && !data.LocalAddress.IsNull() {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `XMLManager`); value.Exists() && !data.XmlManager.IsNull() {
		data.XmlManager = tfutils.ParseStringFromGJSON(value)
	} else if data.XmlManager.ValueString() != "default" {
		data.XmlManager = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLClient`); value.Exists() && !data.SslClient.IsNull() {
		data.SslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `OutboundSNI`); value.Exists() && !data.OutboundSni.IsNull() {
		data.OutboundSni = tfutils.ParseStringFromGJSON(value)
	} else if data.OutboundSni.ValueString() != "Channel" {
		data.OutboundSni = types.StringNull()
	}
	if value := res.Get(pathRoot + `OCSPCheckExtensions`); value.Exists() && !data.OcspCheckExtensions.IsNull() {
		data.OcspCheckExtensions = tfutils.BoolFromString(value.String())
	} else if !data.OcspCheckExtensions.ValueBool() {
		data.OcspCheckExtensions = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OCSPAuthentication`); value.Exists() && !data.OcspAuthentication.IsNull() {
		data.OcspAuthentication = tfutils.ParseStringFromGJSON(value)
	} else if data.OcspAuthentication.ValueString() != "REQUIRED" {
		data.OcspAuthentication = types.StringNull()
	}
	if value := res.Get(pathRoot + `CDPCheckExtensions`); value.Exists() && !data.CdpCheckExtensions.IsNull() {
		data.CdpCheckExtensions = tfutils.BoolFromString(value.String())
	} else if data.CdpCheckExtensions.ValueBool() {
		data.CdpCheckExtensions = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ClientRevocationChecks`); value.Exists() && !data.ClientRevocationChecks.IsNull() {
		data.ClientRevocationChecks = tfutils.ParseStringFromGJSON(value)
	} else if data.ClientRevocationChecks.ValueString() != "REQUIRED" {
		data.ClientRevocationChecks = types.StringNull()
	}
}
