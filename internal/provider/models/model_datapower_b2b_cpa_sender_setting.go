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

type B2BCPASenderSetting struct {
	Id                     types.String                `tfsdk:"id"`
	AppDomain              types.String                `tfsdk:"app_domain"`
	UserSummary            types.String                `tfsdk:"user_summary"`
	EnabledDocType         *DmB2BEnabledDocType        `tfsdk:"enabled_doc_type"`
	DestEndpointUrl        types.String                `tfsdk:"dest_endpoint_url"`
	UserName               types.String                `tfsdk:"user_name"`
	PasswordAlias          types.String                `tfsdk:"password_alias"`
	ConnectionTimeout      types.Int64                 `tfsdk:"connection_timeout"`
	SyncReplyMode          types.String                `tfsdk:"sync_reply_mode"`
	DuplicateElimination   types.String                `tfsdk:"duplicate_elimination"`
	AckRequested           types.String                `tfsdk:"ack_requested"`
	AckSignatureRequested  types.String                `tfsdk:"ack_signature_requested"`
	Retry                  types.Bool                  `tfsdk:"retry"`
	MaxRetries             types.Int64                 `tfsdk:"max_retries"`
	RetryInterval          types.Int64                 `tfsdk:"retry_interval"`
	PersistDuration        types.Int64                 `tfsdk:"persist_duration"`
	IncludeTimeToLive      types.Bool                  `tfsdk:"include_time_to_live"`
	EncryptionRequired     types.Bool                  `tfsdk:"encryption_required"`
	EncryptCert            types.String                `tfsdk:"encrypt_cert"`
	EncryptAlgorithm       types.String                `tfsdk:"encrypt_algorithm"`
	SignatureRequired      types.Bool                  `tfsdk:"signature_required"`
	SignIdCred             types.String                `tfsdk:"sign_id_cred"`
	SignatureAlgorithm     types.String                `tfsdk:"signature_algorithm"`
	SignDigestAlgorithm    types.String                `tfsdk:"sign_digest_algorithm"`
	SignatureC14nAlgorithm types.String                `tfsdk:"signature_c14n_algorithm"`
	SslClientConfigType    types.String                `tfsdk:"ssl_client_config_type"`
	SslClient              types.String                `tfsdk:"ssl_client"`
	DependencyActions      []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var B2BCPASenderSettingMaxRetriesCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "retry",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var B2BCPASenderSettingMaxRetriesIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "retry",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var B2BCPASenderSettingRetryIntervalCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "retry",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var B2BCPASenderSettingRetryIntervalIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "retry",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var B2BCPASenderSettingIncludeTimeToLiveIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "retry",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var B2BCPASenderSettingEncryptCertCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "encryption_required",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var B2BCPASenderSettingEncryptAlgorithmCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "encryption_required",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var B2BCPASenderSettingSignIdCredCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "signature_required",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var B2BCPASenderSettingSignatureAlgorithmCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "signature_required",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var B2BCPASenderSettingSignDigestAlgorithmCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "signature_required",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var B2BCPASenderSettingSignatureC14NAlgorithmCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "signature_required",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var B2BCPASenderSettingSSLClientConfigTypeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-url-protocol-not-in-list",
	Attribute:   "dest_endpoint_url",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"ebms2s"},
}

var B2BCPASenderSettingSSLClientCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ssl_client_config_type",
			AttrType:    "String",
			AttrDefault: "client",
			Value:       []string{"client"},
		},
		{
			Evaluation:  "property-url-protocol-in-list",
			Attribute:   "dest_endpoint_url",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"ebms2s"},
		},
	},
}

var B2BCPASenderSettingSSLClientIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var B2BCPASenderSettingObjectType = map[string]attr.Type{
	"id":                       types.StringType,
	"app_domain":               types.StringType,
	"user_summary":             types.StringType,
	"enabled_doc_type":         types.ObjectType{AttrTypes: DmB2BEnabledDocTypeObjectType},
	"dest_endpoint_url":        types.StringType,
	"user_name":                types.StringType,
	"password_alias":           types.StringType,
	"connection_timeout":       types.Int64Type,
	"sync_reply_mode":          types.StringType,
	"duplicate_elimination":    types.StringType,
	"ack_requested":            types.StringType,
	"ack_signature_requested":  types.StringType,
	"retry":                    types.BoolType,
	"max_retries":              types.Int64Type,
	"retry_interval":           types.Int64Type,
	"persist_duration":         types.Int64Type,
	"include_time_to_live":     types.BoolType,
	"encryption_required":      types.BoolType,
	"encrypt_cert":             types.StringType,
	"encrypt_algorithm":        types.StringType,
	"signature_required":       types.BoolType,
	"sign_id_cred":             types.StringType,
	"signature_algorithm":      types.StringType,
	"sign_digest_algorithm":    types.StringType,
	"signature_c14n_algorithm": types.StringType,
	"ssl_client_config_type":   types.StringType,
	"ssl_client":               types.StringType,
	"dependency_actions":       actions.ActionsListType,
}

func (data B2BCPASenderSetting) GetPath() string {
	rest_path := "/mgmt/config/{domain}/B2BCPASenderSetting"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data B2BCPASenderSetting) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if data.EnabledDocType != nil {
		if !data.EnabledDocType.IsNull() {
			return false
		}
	}
	if !data.DestEndpointUrl.IsNull() {
		return false
	}
	if !data.UserName.IsNull() {
		return false
	}
	if !data.PasswordAlias.IsNull() {
		return false
	}
	if !data.ConnectionTimeout.IsNull() {
		return false
	}
	if !data.SyncReplyMode.IsNull() {
		return false
	}
	if !data.DuplicateElimination.IsNull() {
		return false
	}
	if !data.AckRequested.IsNull() {
		return false
	}
	if !data.AckSignatureRequested.IsNull() {
		return false
	}
	if !data.Retry.IsNull() {
		return false
	}
	if !data.MaxRetries.IsNull() {
		return false
	}
	if !data.RetryInterval.IsNull() {
		return false
	}
	if !data.PersistDuration.IsNull() {
		return false
	}
	if !data.IncludeTimeToLive.IsNull() {
		return false
	}
	if !data.EncryptionRequired.IsNull() {
		return false
	}
	if !data.EncryptCert.IsNull() {
		return false
	}
	if !data.EncryptAlgorithm.IsNull() {
		return false
	}
	if !data.SignatureRequired.IsNull() {
		return false
	}
	if !data.SignIdCred.IsNull() {
		return false
	}
	if !data.SignatureAlgorithm.IsNull() {
		return false
	}
	if !data.SignDigestAlgorithm.IsNull() {
		return false
	}
	if !data.SignatureC14nAlgorithm.IsNull() {
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

func (data B2BCPASenderSetting) ToBody(ctx context.Context, pathRoot string) string {
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
	if data.EnabledDocType != nil {
		if !data.EnabledDocType.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`EnabledDocType`, data.EnabledDocType.ToBody(ctx, ""))
		}
	}
	if !data.DestEndpointUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DestEndpointURL`, data.DestEndpointUrl.ValueString())
	}
	if !data.UserName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserName`, data.UserName.ValueString())
	}
	if !data.PasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PasswordAlias`, data.PasswordAlias.ValueString())
	}
	if !data.ConnectionTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ConnectionTimeout`, data.ConnectionTimeout.ValueInt64())
	}
	if !data.SyncReplyMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SyncReplyMode`, data.SyncReplyMode.ValueString())
	}
	if !data.DuplicateElimination.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DuplicateElimination`, data.DuplicateElimination.ValueString())
	}
	if !data.AckRequested.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AckRequested`, data.AckRequested.ValueString())
	}
	if !data.AckSignatureRequested.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AckSignatureRequested`, data.AckSignatureRequested.ValueString())
	}
	if !data.Retry.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Retry`, tfutils.StringFromBool(data.Retry, ""))
	}
	if !data.MaxRetries.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxRetries`, data.MaxRetries.ValueInt64())
	}
	if !data.RetryInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RetryInterval`, data.RetryInterval.ValueInt64())
	}
	if !data.PersistDuration.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PersistDuration`, data.PersistDuration.ValueInt64())
	}
	if !data.IncludeTimeToLive.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IncludeTimeToLive`, tfutils.StringFromBool(data.IncludeTimeToLive, ""))
	}
	if !data.EncryptionRequired.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EncryptionRequired`, tfutils.StringFromBool(data.EncryptionRequired, ""))
	}
	if !data.EncryptCert.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EncryptCert`, data.EncryptCert.ValueString())
	}
	if !data.EncryptAlgorithm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EncryptAlgorithm`, data.EncryptAlgorithm.ValueString())
	}
	if !data.SignatureRequired.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SignatureRequired`, tfutils.StringFromBool(data.SignatureRequired, ""))
	}
	if !data.SignIdCred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SignIdCred`, data.SignIdCred.ValueString())
	}
	if !data.SignatureAlgorithm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SignatureAlgorithm`, data.SignatureAlgorithm.ValueString())
	}
	if !data.SignDigestAlgorithm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SignDigestAlgorithm`, data.SignDigestAlgorithm.ValueString())
	}
	if !data.SignatureC14nAlgorithm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SignatureC14NAlgorithm`, data.SignatureC14nAlgorithm.ValueString())
	}
	if !data.SslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClientConfigType`, data.SslClientConfigType.ValueString())
	}
	if !data.SslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClient`, data.SslClient.ValueString())
	}
	return body
}

func (data *B2BCPASenderSetting) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `EnabledDocType`); value.Exists() {
		data.EnabledDocType = &DmB2BEnabledDocType{}
		data.EnabledDocType.FromBody(ctx, "", value)
	} else {
		data.EnabledDocType = nil
	}
	if value := res.Get(pathRoot + `DestEndpointURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DestEndpointUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DestEndpointUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserName = types.StringNull()
	}
	if value := res.Get(pathRoot + `PasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `ConnectionTimeout`); value.Exists() {
		data.ConnectionTimeout = types.Int64Value(value.Int())
	} else {
		data.ConnectionTimeout = types.Int64Value(300)
	}
	if value := res.Get(pathRoot + `SyncReplyMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SyncReplyMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SyncReplyMode = types.StringValue("none")
	}
	if value := res.Get(pathRoot + `DuplicateElimination`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DuplicateElimination = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DuplicateElimination = types.StringValue("always")
	}
	if value := res.Get(pathRoot + `AckRequested`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AckRequested = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AckRequested = types.StringValue("never")
	}
	if value := res.Get(pathRoot + `AckSignatureRequested`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AckSignatureRequested = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AckSignatureRequested = types.StringValue("never")
	}
	if value := res.Get(pathRoot + `Retry`); value.Exists() {
		data.Retry = tfutils.BoolFromString(value.String())
	} else {
		data.Retry = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MaxRetries`); value.Exists() {
		data.MaxRetries = types.Int64Value(value.Int())
	} else {
		data.MaxRetries = types.Int64Value(3)
	}
	if value := res.Get(pathRoot + `RetryInterval`); value.Exists() {
		data.RetryInterval = types.Int64Value(value.Int())
	} else {
		data.RetryInterval = types.Int64Value(1800)
	}
	if value := res.Get(pathRoot + `PersistDuration`); value.Exists() {
		data.PersistDuration = types.Int64Value(value.Int())
	} else {
		data.PersistDuration = types.Int64Null()
	}
	if value := res.Get(pathRoot + `IncludeTimeToLive`); value.Exists() {
		data.IncludeTimeToLive = tfutils.BoolFromString(value.String())
	} else {
		data.IncludeTimeToLive = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EncryptionRequired`); value.Exists() {
		data.EncryptionRequired = tfutils.BoolFromString(value.String())
	} else {
		data.EncryptionRequired = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EncryptCert`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EncryptCert = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EncryptCert = types.StringNull()
	}
	if value := res.Get(pathRoot + `EncryptAlgorithm`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EncryptAlgorithm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EncryptAlgorithm = types.StringValue("http://www.w3.org/2001/04/xmlenc#tripledes-cbc")
	}
	if value := res.Get(pathRoot + `SignatureRequired`); value.Exists() {
		data.SignatureRequired = tfutils.BoolFromString(value.String())
	} else {
		data.SignatureRequired = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SignIdCred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SignIdCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SignIdCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `SignatureAlgorithm`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SignatureAlgorithm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SignatureAlgorithm = types.StringValue("rsa-sha1")
	}
	if value := res.Get(pathRoot + `SignDigestAlgorithm`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SignDigestAlgorithm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SignDigestAlgorithm = types.StringValue("sha1")
	}
	if value := res.Get(pathRoot + `SignatureC14NAlgorithm`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SignatureC14nAlgorithm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SignatureC14nAlgorithm = types.StringValue("exc-c14n")
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

func (data *B2BCPASenderSetting) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `EnabledDocType`); value.Exists() {
		data.EnabledDocType.UpdateFromBody(ctx, "", value)
	} else {
		data.EnabledDocType = nil
	}
	if value := res.Get(pathRoot + `DestEndpointURL`); value.Exists() && !data.DestEndpointUrl.IsNull() {
		data.DestEndpointUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DestEndpointUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserName`); value.Exists() && !data.UserName.IsNull() {
		data.UserName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserName = types.StringNull()
	}
	if value := res.Get(pathRoot + `PasswordAlias`); value.Exists() && !data.PasswordAlias.IsNull() {
		data.PasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `ConnectionTimeout`); value.Exists() && !data.ConnectionTimeout.IsNull() {
		data.ConnectionTimeout = types.Int64Value(value.Int())
	} else if data.ConnectionTimeout.ValueInt64() != 300 {
		data.ConnectionTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `SyncReplyMode`); value.Exists() && !data.SyncReplyMode.IsNull() {
		data.SyncReplyMode = tfutils.ParseStringFromGJSON(value)
	} else if data.SyncReplyMode.ValueString() != "none" {
		data.SyncReplyMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `DuplicateElimination`); value.Exists() && !data.DuplicateElimination.IsNull() {
		data.DuplicateElimination = tfutils.ParseStringFromGJSON(value)
	} else if data.DuplicateElimination.ValueString() != "always" {
		data.DuplicateElimination = types.StringNull()
	}
	if value := res.Get(pathRoot + `AckRequested`); value.Exists() && !data.AckRequested.IsNull() {
		data.AckRequested = tfutils.ParseStringFromGJSON(value)
	} else if data.AckRequested.ValueString() != "never" {
		data.AckRequested = types.StringNull()
	}
	if value := res.Get(pathRoot + `AckSignatureRequested`); value.Exists() && !data.AckSignatureRequested.IsNull() {
		data.AckSignatureRequested = tfutils.ParseStringFromGJSON(value)
	} else if data.AckSignatureRequested.ValueString() != "never" {
		data.AckSignatureRequested = types.StringNull()
	}
	if value := res.Get(pathRoot + `Retry`); value.Exists() && !data.Retry.IsNull() {
		data.Retry = tfutils.BoolFromString(value.String())
	} else if data.Retry.ValueBool() {
		data.Retry = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MaxRetries`); value.Exists() && !data.MaxRetries.IsNull() {
		data.MaxRetries = types.Int64Value(value.Int())
	} else if data.MaxRetries.ValueInt64() != 3 {
		data.MaxRetries = types.Int64Null()
	}
	if value := res.Get(pathRoot + `RetryInterval`); value.Exists() && !data.RetryInterval.IsNull() {
		data.RetryInterval = types.Int64Value(value.Int())
	} else if data.RetryInterval.ValueInt64() != 1800 {
		data.RetryInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `PersistDuration`); value.Exists() && !data.PersistDuration.IsNull() {
		data.PersistDuration = types.Int64Value(value.Int())
	} else {
		data.PersistDuration = types.Int64Null()
	}
	if value := res.Get(pathRoot + `IncludeTimeToLive`); value.Exists() && !data.IncludeTimeToLive.IsNull() {
		data.IncludeTimeToLive = tfutils.BoolFromString(value.String())
	} else if !data.IncludeTimeToLive.ValueBool() {
		data.IncludeTimeToLive = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EncryptionRequired`); value.Exists() && !data.EncryptionRequired.IsNull() {
		data.EncryptionRequired = tfutils.BoolFromString(value.String())
	} else if data.EncryptionRequired.ValueBool() {
		data.EncryptionRequired = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EncryptCert`); value.Exists() && !data.EncryptCert.IsNull() {
		data.EncryptCert = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EncryptCert = types.StringNull()
	}
	if value := res.Get(pathRoot + `EncryptAlgorithm`); value.Exists() && !data.EncryptAlgorithm.IsNull() {
		data.EncryptAlgorithm = tfutils.ParseStringFromGJSON(value)
	} else if data.EncryptAlgorithm.ValueString() != "http://www.w3.org/2001/04/xmlenc#tripledes-cbc" {
		data.EncryptAlgorithm = types.StringNull()
	}
	if value := res.Get(pathRoot + `SignatureRequired`); value.Exists() && !data.SignatureRequired.IsNull() {
		data.SignatureRequired = tfutils.BoolFromString(value.String())
	} else if data.SignatureRequired.ValueBool() {
		data.SignatureRequired = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SignIdCred`); value.Exists() && !data.SignIdCred.IsNull() {
		data.SignIdCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SignIdCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `SignatureAlgorithm`); value.Exists() && !data.SignatureAlgorithm.IsNull() {
		data.SignatureAlgorithm = tfutils.ParseStringFromGJSON(value)
	} else if data.SignatureAlgorithm.ValueString() != "rsa-sha1" {
		data.SignatureAlgorithm = types.StringNull()
	}
	if value := res.Get(pathRoot + `SignDigestAlgorithm`); value.Exists() && !data.SignDigestAlgorithm.IsNull() {
		data.SignDigestAlgorithm = tfutils.ParseStringFromGJSON(value)
	} else if data.SignDigestAlgorithm.ValueString() != "sha1" {
		data.SignDigestAlgorithm = types.StringNull()
	}
	if value := res.Get(pathRoot + `SignatureC14NAlgorithm`); value.Exists() && !data.SignatureC14nAlgorithm.IsNull() {
		data.SignatureC14nAlgorithm = tfutils.ParseStringFromGJSON(value)
	} else if data.SignatureC14nAlgorithm.ValueString() != "exc-c14n" {
		data.SignatureC14nAlgorithm = types.StringNull()
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
