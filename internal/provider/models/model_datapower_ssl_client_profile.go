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

type SSLClientProfile struct {
	Id                            types.String                  `tfsdk:"id"`
	AppDomain                     types.String                  `tfsdk:"app_domain"`
	UserSummary                   types.String                  `tfsdk:"user_summary"`
	Protocols                     *DmSSLProtoVersionsBitmap     `tfsdk:"protocols"`
	Ciphers                       types.List                    `tfsdk:"ciphers"`
	Idcred                        types.String                  `tfsdk:"idcred"`
	ValidateServerCert            types.Bool                    `tfsdk:"validate_server_cert"`
	Valcred                       types.String                  `tfsdk:"valcred"`
	Caching                       types.Bool                    `tfsdk:"caching"`
	CacheTimeout                  types.Int64                   `tfsdk:"cache_timeout"`
	CacheSize                     types.Int64                   `tfsdk:"cache_size"`
	SslClientFeatures             *DmSSLClientFeatures          `tfsdk:"ssl_client_features"`
	EllipticCurves                types.List                    `tfsdk:"elliptic_curves"`
	UseCustomSniHostname          types.Bool                    `tfsdk:"use_custom_sni_hostname"`
	CustomSniHostname             types.String                  `tfsdk:"custom_sni_hostname"`
	ValidateHostname              types.Bool                    `tfsdk:"validate_hostname"`
	HostnameValidationFlags       *DmSSLHostnameValidationFlags `tfsdk:"hostname_validation_flags"`
	HostnameValidationFailOnError types.Bool                    `tfsdk:"hostname_validation_fail_on_error"`
	EnableTls13Compat             types.Bool                    `tfsdk:"enable_tls13_compat"`
	DisableRenegotiation          types.Bool                    `tfsdk:"disable_renegotiation"`
	SigAlgs                       types.List                    `tfsdk:"sig_algs"`
	RequireClosureNotification    types.Bool                    `tfsdk:"require_closure_notification"`
	DependencyActions             []*actions.DependencyAction   `tfsdk:"dependency_actions"`
}

var SSLClientProfileObjectType = map[string]attr.Type{
	"id":                                types.StringType,
	"app_domain":                        types.StringType,
	"user_summary":                      types.StringType,
	"protocols":                         types.ObjectType{AttrTypes: DmSSLProtoVersionsBitmapObjectType},
	"ciphers":                           types.ListType{ElemType: types.StringType},
	"idcred":                            types.StringType,
	"validate_server_cert":              types.BoolType,
	"valcred":                           types.StringType,
	"caching":                           types.BoolType,
	"cache_timeout":                     types.Int64Type,
	"cache_size":                        types.Int64Type,
	"ssl_client_features":               types.ObjectType{AttrTypes: DmSSLClientFeaturesObjectType},
	"elliptic_curves":                   types.ListType{ElemType: types.StringType},
	"use_custom_sni_hostname":           types.BoolType,
	"custom_sni_hostname":               types.StringType,
	"validate_hostname":                 types.BoolType,
	"hostname_validation_flags":         types.ObjectType{AttrTypes: DmSSLHostnameValidationFlagsObjectType},
	"hostname_validation_fail_on_error": types.BoolType,
	"enable_tls13_compat":               types.BoolType,
	"disable_renegotiation":             types.BoolType,
	"sig_algs":                          types.ListType{ElemType: types.StringType},
	"require_closure_notification":      types.BoolType,
	"dependency_actions":                actions.ActionsListType,
}

func (data SSLClientProfile) GetPath() string {
	rest_path := "/mgmt/config/{domain}/SSLClientProfile"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data SSLClientProfile) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if data.Protocols != nil {
		if !data.Protocols.IsNull() {
			return false
		}
	}
	if !data.Ciphers.IsNull() {
		return false
	}
	if !data.Idcred.IsNull() {
		return false
	}
	if !data.ValidateServerCert.IsNull() {
		return false
	}
	if !data.Valcred.IsNull() {
		return false
	}
	if !data.Caching.IsNull() {
		return false
	}
	if !data.CacheTimeout.IsNull() {
		return false
	}
	if !data.CacheSize.IsNull() {
		return false
	}
	if data.SslClientFeatures != nil {
		if !data.SslClientFeatures.IsNull() {
			return false
		}
	}
	if !data.EllipticCurves.IsNull() {
		return false
	}
	if !data.UseCustomSniHostname.IsNull() {
		return false
	}
	if !data.CustomSniHostname.IsNull() {
		return false
	}
	if !data.ValidateHostname.IsNull() {
		return false
	}
	if data.HostnameValidationFlags != nil {
		if !data.HostnameValidationFlags.IsNull() {
			return false
		}
	}
	if !data.HostnameValidationFailOnError.IsNull() {
		return false
	}
	if !data.EnableTls13Compat.IsNull() {
		return false
	}
	if !data.DisableRenegotiation.IsNull() {
		return false
	}
	if !data.SigAlgs.IsNull() {
		return false
	}
	if !data.RequireClosureNotification.IsNull() {
		return false
	}
	return true
}

func (data SSLClientProfile) ToBody(ctx context.Context, pathRoot string) string {
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
	if data.Protocols != nil {
		if !data.Protocols.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`Protocols`, data.Protocols.ToBody(ctx, ""))
		}
	}
	if !data.Ciphers.IsNull() {
		var values []string
		data.Ciphers.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`Ciphers`+".-1", map[string]string{"value": val})
		}
	}
	if !data.Idcred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Idcred`, data.Idcred.ValueString())
	}
	if !data.ValidateServerCert.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ValidateServerCert`, tfutils.StringFromBool(data.ValidateServerCert, ""))
	}
	if !data.Valcred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Valcred`, data.Valcred.ValueString())
	}
	if !data.Caching.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Caching`, tfutils.StringFromBool(data.Caching, ""))
	}
	if !data.CacheTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CacheTimeout`, data.CacheTimeout.ValueInt64())
	}
	if !data.CacheSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CacheSize`, data.CacheSize.ValueInt64())
	}
	if data.SslClientFeatures != nil {
		if !data.SslClientFeatures.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`SSLClientFeatures`, data.SslClientFeatures.ToBody(ctx, ""))
		}
	}
	if !data.EllipticCurves.IsNull() {
		var values []string
		data.EllipticCurves.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`EllipticCurves`+".-1", map[string]string{"value": val})
		}
	}
	if !data.UseCustomSniHostname.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseCustomSNIHostname`, tfutils.StringFromBool(data.UseCustomSniHostname, "flag"))
	}
	if !data.CustomSniHostname.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CustomSNIHostname`, data.CustomSniHostname.ValueString())
	}
	if !data.ValidateHostname.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ValidateHostname`, tfutils.StringFromBool(data.ValidateHostname, ""))
	}
	if data.HostnameValidationFlags != nil {
		if !data.HostnameValidationFlags.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`HostnameValidationFlags`, data.HostnameValidationFlags.ToBody(ctx, ""))
		}
	}
	if !data.HostnameValidationFailOnError.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HostnameValidationFailOnError`, tfutils.StringFromBool(data.HostnameValidationFailOnError, ""))
	}
	if !data.EnableTls13Compat.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnableTLS13Compat`, tfutils.StringFromBool(data.EnableTls13Compat, ""))
	}
	if !data.DisableRenegotiation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DisableRenegotiation`, tfutils.StringFromBool(data.DisableRenegotiation, ""))
	}
	if !data.SigAlgs.IsNull() {
		var values []string
		data.SigAlgs.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`SigAlgs`+".-1", map[string]string{"value": val})
		}
	}
	if !data.RequireClosureNotification.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RequireClosureNotification`, tfutils.StringFromBool(data.RequireClosureNotification, ""))
	}
	return body
}

func (data *SSLClientProfile) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Protocols`); value.Exists() {
		data.Protocols = &DmSSLProtoVersionsBitmap{}
		data.Protocols.FromBody(ctx, "", value)
	} else {
		data.Protocols = nil
	}
	if value := res.Get(pathRoot + `Ciphers`); value.Exists() {
		data.Ciphers = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Ciphers = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `Idcred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Idcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Idcred = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidateServerCert`); value.Exists() {
		data.ValidateServerCert = tfutils.BoolFromString(value.String())
	} else {
		data.ValidateServerCert = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Valcred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Valcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Valcred = types.StringNull()
	}
	if value := res.Get(pathRoot + `Caching`); value.Exists() {
		data.Caching = tfutils.BoolFromString(value.String())
	} else {
		data.Caching = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CacheTimeout`); value.Exists() {
		data.CacheTimeout = types.Int64Value(value.Int())
	} else {
		data.CacheTimeout = types.Int64Value(300)
	}
	if value := res.Get(pathRoot + `CacheSize`); value.Exists() {
		data.CacheSize = types.Int64Value(value.Int())
	} else {
		data.CacheSize = types.Int64Value(100)
	}
	if value := res.Get(pathRoot + `SSLClientFeatures`); value.Exists() {
		data.SslClientFeatures = &DmSSLClientFeatures{}
		data.SslClientFeatures.FromBody(ctx, "", value)
	} else {
		data.SslClientFeatures = nil
	}
	if value := res.Get(pathRoot + `EllipticCurves`); value.Exists() {
		data.EllipticCurves = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.EllipticCurves = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `UseCustomSNIHostname`); value.Exists() {
		data.UseCustomSniHostname = tfutils.BoolFromString(value.String())
	} else {
		data.UseCustomSniHostname = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CustomSNIHostname`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CustomSniHostname = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CustomSniHostname = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidateHostname`); value.Exists() {
		data.ValidateHostname = tfutils.BoolFromString(value.String())
	} else {
		data.ValidateHostname = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HostnameValidationFlags`); value.Exists() {
		data.HostnameValidationFlags = &DmSSLHostnameValidationFlags{}
		data.HostnameValidationFlags.FromBody(ctx, "", value)
	} else {
		data.HostnameValidationFlags = nil
	}
	if value := res.Get(pathRoot + `HostnameValidationFailOnError`); value.Exists() {
		data.HostnameValidationFailOnError = tfutils.BoolFromString(value.String())
	} else {
		data.HostnameValidationFailOnError = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EnableTLS13Compat`); value.Exists() {
		data.EnableTls13Compat = tfutils.BoolFromString(value.String())
	} else {
		data.EnableTls13Compat = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DisableRenegotiation`); value.Exists() {
		data.DisableRenegotiation = tfutils.BoolFromString(value.String())
	} else {
		data.DisableRenegotiation = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SigAlgs`); value.Exists() {
		data.SigAlgs = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.SigAlgs = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `RequireClosureNotification`); value.Exists() {
		data.RequireClosureNotification = tfutils.BoolFromString(value.String())
	} else {
		data.RequireClosureNotification = types.BoolNull()
	}
}

func (data *SSLClientProfile) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Protocols`); value.Exists() {
		data.Protocols.UpdateFromBody(ctx, "", value)
	} else {
		data.Protocols = nil
	}
	if value := res.Get(pathRoot + `Ciphers`); value.Exists() && !data.Ciphers.IsNull() {
		data.Ciphers = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Ciphers = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `Idcred`); value.Exists() && !data.Idcred.IsNull() {
		data.Idcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Idcred = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidateServerCert`); value.Exists() && !data.ValidateServerCert.IsNull() {
		data.ValidateServerCert = tfutils.BoolFromString(value.String())
	} else if !data.ValidateServerCert.ValueBool() {
		data.ValidateServerCert = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Valcred`); value.Exists() && !data.Valcred.IsNull() {
		data.Valcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Valcred = types.StringNull()
	}
	if value := res.Get(pathRoot + `Caching`); value.Exists() && !data.Caching.IsNull() {
		data.Caching = tfutils.BoolFromString(value.String())
	} else if !data.Caching.ValueBool() {
		data.Caching = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CacheTimeout`); value.Exists() && !data.CacheTimeout.IsNull() {
		data.CacheTimeout = types.Int64Value(value.Int())
	} else if data.CacheTimeout.ValueInt64() != 300 {
		data.CacheTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `CacheSize`); value.Exists() && !data.CacheSize.IsNull() {
		data.CacheSize = types.Int64Value(value.Int())
	} else if data.CacheSize.ValueInt64() != 100 {
		data.CacheSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `SSLClientFeatures`); value.Exists() {
		data.SslClientFeatures.UpdateFromBody(ctx, "", value)
	} else {
		data.SslClientFeatures = nil
	}
	if value := res.Get(pathRoot + `EllipticCurves`); value.Exists() && !data.EllipticCurves.IsNull() {
		data.EllipticCurves = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.EllipticCurves = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `UseCustomSNIHostname`); value.Exists() && !data.UseCustomSniHostname.IsNull() {
		data.UseCustomSniHostname = tfutils.BoolFromString(value.String())
	} else if data.UseCustomSniHostname.ValueBool() {
		data.UseCustomSniHostname = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CustomSNIHostname`); value.Exists() && !data.CustomSniHostname.IsNull() {
		data.CustomSniHostname = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CustomSniHostname = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidateHostname`); value.Exists() && !data.ValidateHostname.IsNull() {
		data.ValidateHostname = tfutils.BoolFromString(value.String())
	} else if data.ValidateHostname.ValueBool() {
		data.ValidateHostname = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HostnameValidationFlags`); value.Exists() {
		data.HostnameValidationFlags.UpdateFromBody(ctx, "", value)
	} else {
		data.HostnameValidationFlags = nil
	}
	if value := res.Get(pathRoot + `HostnameValidationFailOnError`); value.Exists() && !data.HostnameValidationFailOnError.IsNull() {
		data.HostnameValidationFailOnError = tfutils.BoolFromString(value.String())
	} else if data.HostnameValidationFailOnError.ValueBool() {
		data.HostnameValidationFailOnError = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EnableTLS13Compat`); value.Exists() && !data.EnableTls13Compat.IsNull() {
		data.EnableTls13Compat = tfutils.BoolFromString(value.String())
	} else if !data.EnableTls13Compat.ValueBool() {
		data.EnableTls13Compat = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DisableRenegotiation`); value.Exists() && !data.DisableRenegotiation.IsNull() {
		data.DisableRenegotiation = tfutils.BoolFromString(value.String())
	} else if !data.DisableRenegotiation.ValueBool() {
		data.DisableRenegotiation = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SigAlgs`); value.Exists() && !data.SigAlgs.IsNull() {
		data.SigAlgs = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.SigAlgs = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `RequireClosureNotification`); value.Exists() && !data.RequireClosureNotification.IsNull() {
		data.RequireClosureNotification = tfutils.BoolFromString(value.String())
	} else if !data.RequireClosureNotification.ValueBool() {
		data.RequireClosureNotification = types.BoolNull()
	}
}
