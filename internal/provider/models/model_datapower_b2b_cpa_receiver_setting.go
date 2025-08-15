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

type B2BCPAReceiverSetting struct {
	Id                    types.String                `tfsdk:"id"`
	AppDomain             types.String                `tfsdk:"app_domain"`
	UserSummary           types.String                `tfsdk:"user_summary"`
	LocalEndpointUri      types.String                `tfsdk:"local_endpoint_uri"`
	SyncReplyMode         types.String                `tfsdk:"sync_reply_mode"`
	AckRequested          types.String                `tfsdk:"ack_requested"`
	AckSignatureRequested types.String                `tfsdk:"ack_signature_requested"`
	AllowDuplicateMessage types.String                `tfsdk:"allow_duplicate_message"`
	PersistDuration       types.Int64                 `tfsdk:"persist_duration"`
	EncryptionRequired    types.Bool                  `tfsdk:"encryption_required"`
	DecryptIdCred         types.String                `tfsdk:"decrypt_id_cred"`
	SignatureRequired     types.Bool                  `tfsdk:"signature_required"`
	VerifyValCred         types.String                `tfsdk:"verify_val_cred"`
	DefaultSignerCert     types.String                `tfsdk:"default_signer_cert"`
	DependencyActions     []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var B2BCPAReceiverSettingObjectType = map[string]attr.Type{
	"id":                      types.StringType,
	"app_domain":              types.StringType,
	"user_summary":            types.StringType,
	"local_endpoint_uri":      types.StringType,
	"sync_reply_mode":         types.StringType,
	"ack_requested":           types.StringType,
	"ack_signature_requested": types.StringType,
	"allow_duplicate_message": types.StringType,
	"persist_duration":        types.Int64Type,
	"encryption_required":     types.BoolType,
	"decrypt_id_cred":         types.StringType,
	"signature_required":      types.BoolType,
	"verify_val_cred":         types.StringType,
	"default_signer_cert":     types.StringType,
	"dependency_actions":      actions.ActionsListType,
}

func (data B2BCPAReceiverSetting) GetPath() string {
	rest_path := "/mgmt/config/{domain}/B2BCPAReceiverSetting"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data B2BCPAReceiverSetting) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.LocalEndpointUri.IsNull() {
		return false
	}
	if !data.SyncReplyMode.IsNull() {
		return false
	}
	if !data.AckRequested.IsNull() {
		return false
	}
	if !data.AckSignatureRequested.IsNull() {
		return false
	}
	if !data.AllowDuplicateMessage.IsNull() {
		return false
	}
	if !data.PersistDuration.IsNull() {
		return false
	}
	if !data.EncryptionRequired.IsNull() {
		return false
	}
	if !data.DecryptIdCred.IsNull() {
		return false
	}
	if !data.SignatureRequired.IsNull() {
		return false
	}
	if !data.VerifyValCred.IsNull() {
		return false
	}
	if !data.DefaultSignerCert.IsNull() {
		return false
	}
	return true
}

func (data B2BCPAReceiverSetting) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.LocalEndpointUri.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalEndpointURI`, data.LocalEndpointUri.ValueString())
	}
	if !data.SyncReplyMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SyncReplyMode`, data.SyncReplyMode.ValueString())
	}
	if !data.AckRequested.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AckRequested`, data.AckRequested.ValueString())
	}
	if !data.AckSignatureRequested.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AckSignatureRequested`, data.AckSignatureRequested.ValueString())
	}
	if !data.AllowDuplicateMessage.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllowDuplicateMessage`, data.AllowDuplicateMessage.ValueString())
	}
	if !data.PersistDuration.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PersistDuration`, data.PersistDuration.ValueInt64())
	}
	if !data.EncryptionRequired.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EncryptionRequired`, tfutils.StringFromBool(data.EncryptionRequired, ""))
	}
	if !data.DecryptIdCred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DecryptIdCred`, data.DecryptIdCred.ValueString())
	}
	if !data.SignatureRequired.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SignatureRequired`, tfutils.StringFromBool(data.SignatureRequired, ""))
	}
	if !data.VerifyValCred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`VerifyValCred`, data.VerifyValCred.ValueString())
	}
	if !data.DefaultSignerCert.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DefaultSignerCert`, data.DefaultSignerCert.ValueString())
	}
	return body
}

func (data *B2BCPAReceiverSetting) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `LocalEndpointURI`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalEndpointUri = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalEndpointUri = types.StringNull()
	}
	if value := res.Get(pathRoot + `SyncReplyMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SyncReplyMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SyncReplyMode = types.StringValue("none")
	}
	if value := res.Get(pathRoot + `AckRequested`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AckRequested = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AckRequested = types.StringValue("perMessage")
	}
	if value := res.Get(pathRoot + `AckSignatureRequested`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AckSignatureRequested = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AckSignatureRequested = types.StringValue("perMessage")
	}
	if value := res.Get(pathRoot + `AllowDuplicateMessage`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AllowDuplicateMessage = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AllowDuplicateMessage = types.StringValue("never")
	}
	if value := res.Get(pathRoot + `PersistDuration`); value.Exists() {
		data.PersistDuration = types.Int64Value(value.Int())
	} else {
		data.PersistDuration = types.Int64Null()
	}
	if value := res.Get(pathRoot + `EncryptionRequired`); value.Exists() {
		data.EncryptionRequired = tfutils.BoolFromString(value.String())
	} else {
		data.EncryptionRequired = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DecryptIdCred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DecryptIdCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DecryptIdCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `SignatureRequired`); value.Exists() {
		data.SignatureRequired = tfutils.BoolFromString(value.String())
	} else {
		data.SignatureRequired = types.BoolNull()
	}
	if value := res.Get(pathRoot + `VerifyValCred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.VerifyValCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.VerifyValCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `DefaultSignerCert`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DefaultSignerCert = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DefaultSignerCert = types.StringNull()
	}
}

func (data *B2BCPAReceiverSetting) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `LocalEndpointURI`); value.Exists() && !data.LocalEndpointUri.IsNull() {
		data.LocalEndpointUri = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalEndpointUri = types.StringNull()
	}
	if value := res.Get(pathRoot + `SyncReplyMode`); value.Exists() && !data.SyncReplyMode.IsNull() {
		data.SyncReplyMode = tfutils.ParseStringFromGJSON(value)
	} else if data.SyncReplyMode.ValueString() != "none" {
		data.SyncReplyMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `AckRequested`); value.Exists() && !data.AckRequested.IsNull() {
		data.AckRequested = tfutils.ParseStringFromGJSON(value)
	} else if data.AckRequested.ValueString() != "perMessage" {
		data.AckRequested = types.StringNull()
	}
	if value := res.Get(pathRoot + `AckSignatureRequested`); value.Exists() && !data.AckSignatureRequested.IsNull() {
		data.AckSignatureRequested = tfutils.ParseStringFromGJSON(value)
	} else if data.AckSignatureRequested.ValueString() != "perMessage" {
		data.AckSignatureRequested = types.StringNull()
	}
	if value := res.Get(pathRoot + `AllowDuplicateMessage`); value.Exists() && !data.AllowDuplicateMessage.IsNull() {
		data.AllowDuplicateMessage = tfutils.ParseStringFromGJSON(value)
	} else if data.AllowDuplicateMessage.ValueString() != "never" {
		data.AllowDuplicateMessage = types.StringNull()
	}
	if value := res.Get(pathRoot + `PersistDuration`); value.Exists() && !data.PersistDuration.IsNull() {
		data.PersistDuration = types.Int64Value(value.Int())
	} else {
		data.PersistDuration = types.Int64Null()
	}
	if value := res.Get(pathRoot + `EncryptionRequired`); value.Exists() && !data.EncryptionRequired.IsNull() {
		data.EncryptionRequired = tfutils.BoolFromString(value.String())
	} else if data.EncryptionRequired.ValueBool() {
		data.EncryptionRequired = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DecryptIdCred`); value.Exists() && !data.DecryptIdCred.IsNull() {
		data.DecryptIdCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DecryptIdCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `SignatureRequired`); value.Exists() && !data.SignatureRequired.IsNull() {
		data.SignatureRequired = tfutils.BoolFromString(value.String())
	} else if data.SignatureRequired.ValueBool() {
		data.SignatureRequired = types.BoolNull()
	}
	if value := res.Get(pathRoot + `VerifyValCred`); value.Exists() && !data.VerifyValCred.IsNull() {
		data.VerifyValCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.VerifyValCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `DefaultSignerCert`); value.Exists() && !data.DefaultSignerCert.IsNull() {
		data.DefaultSignerCert = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DefaultSignerCert = types.StringNull()
	}
}
