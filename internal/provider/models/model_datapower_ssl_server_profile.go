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

type SSLServerProfile struct {
	Id                           types.String                `tfsdk:"id"`
	AppDomain                    types.String                `tfsdk:"app_domain"`
	UserSummary                  types.String                `tfsdk:"user_summary"`
	Protocols                    *DmSSLProtoVersionsBitmap   `tfsdk:"protocols"`
	Ciphers                      types.List                  `tfsdk:"ciphers"`
	Idcred                       types.String                `tfsdk:"idcred"`
	RequestClientAuth            types.Bool                  `tfsdk:"request_client_auth"`
	RequireClientAuth            types.Bool                  `tfsdk:"require_client_auth"`
	ValidateClientCert           types.Bool                  `tfsdk:"validate_client_cert"`
	SendClientAuthCaList         types.Bool                  `tfsdk:"send_client_auth_ca_list"`
	Valcred                      types.String                `tfsdk:"valcred"`
	Caching                      types.Bool                  `tfsdk:"caching"`
	CacheTimeout                 types.Int64                 `tfsdk:"cache_timeout"`
	CacheSize                    types.Int64                 `tfsdk:"cache_size"`
	SslOptions                   *DmSSLOptions               `tfsdk:"ssl_options"`
	MaxSslDuration               types.Int64                 `tfsdk:"max_ssl_duration"`
	DisableRenegotiation         types.Bool                  `tfsdk:"disable_renegotiation"`
	NumberOfRenegotiationAllowed types.Int64                 `tfsdk:"number_of_renegotiation_allowed"`
	ProhibitResumeOnReneg        types.Bool                  `tfsdk:"prohibit_resume_on_reneg"`
	Compression                  types.Bool                  `tfsdk:"compression"`
	AllowLegacyRenegotiation     types.Bool                  `tfsdk:"allow_legacy_renegotiation"`
	PreferServerCiphers          types.Bool                  `tfsdk:"prefer_server_ciphers"`
	EllipticCurves               types.List                  `tfsdk:"elliptic_curves"`
	PrioritizeChaCha             types.Bool                  `tfsdk:"prioritize_cha_cha"`
	SigAlgs                      types.List                  `tfsdk:"sig_algs"`
	RequireClosureNotification   types.Bool                  `tfsdk:"require_closure_notification"`
	DependencyActions            []*actions.DependencyAction `tfsdk:"dependency_actions"`
	ProviderTarget               types.String                `tfsdk:"provider_target"`
}

var SSLServerProfileRequireClientAuthIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "request_client_auth",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var SSLServerProfileValidateClientCertIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "request_client_auth",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var SSLServerProfileSendClientAuthCAListIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "request_client_auth",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "validate_client_cert",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"false"},
		},
	},
}

var SSLServerProfileValcredCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "request_client_auth",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "validate_client_cert",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
	},
}

var SSLServerProfileValcredIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var SSLServerProfileCacheTimeoutIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "caching",
	AttrType:    "Bool",
	AttrDefault: "true",
	Value:       []string{"false"},
}

var SSLServerProfileCacheSizeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "caching",
	AttrType:    "Bool",
	AttrDefault: "true",
	Value:       []string{"false"},
}

var SSLServerProfileMaxSSLDurationIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "ssl_options",
	AttrType:    "DmSSLOptions",
	AttrDefault: "",
	Value:       []string{"max-duration"},
}

var SSLServerProfileNumberOfRenegotiationAllowedIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "ssl_options",
	AttrType:    "DmSSLOptions",
	AttrDefault: "",
	Value:       []string{"max-renegotiation"},
}

var SSLServerProfileProhibitResumeOnRenegIgnoreVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ssl_options",
			AttrType:    "DmSSLOptions",
			AttrDefault: "",
			Value:       []string{"max-renegotiation"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "number_of_renegotiation_allowed",
			AttrType:    "Int64",
			AttrDefault: "0",
			Value:       []string{"0"},
		},
	},
}

var SSLServerProfileAllowLegacyRenegotiationIgnoreVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ssl_options",
			AttrType:    "DmSSLOptions",
			AttrDefault: "",
			Value:       []string{"max-renegotiation"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "number_of_renegotiation_allowed",
			AttrType:    "Int64",
			AttrDefault: "0",
			Value:       []string{"0"},
		},
	},
}

var SSLServerProfilePrioritizeChaChaIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "prefer_server_ciphers",
	AttrType:    "Bool",
	AttrDefault: "true",
	Value:       []string{"false"},
}

var SSLServerProfileObjectType = map[string]attr.Type{
	"provider_target":                 types.StringType,
	"id":                              types.StringType,
	"app_domain":                      types.StringType,
	"user_summary":                    types.StringType,
	"protocols":                       types.ObjectType{AttrTypes: DmSSLProtoVersionsBitmapObjectType},
	"ciphers":                         types.ListType{ElemType: types.StringType},
	"idcred":                          types.StringType,
	"request_client_auth":             types.BoolType,
	"require_client_auth":             types.BoolType,
	"validate_client_cert":            types.BoolType,
	"send_client_auth_ca_list":        types.BoolType,
	"valcred":                         types.StringType,
	"caching":                         types.BoolType,
	"cache_timeout":                   types.Int64Type,
	"cache_size":                      types.Int64Type,
	"ssl_options":                     types.ObjectType{AttrTypes: DmSSLOptionsObjectType},
	"max_ssl_duration":                types.Int64Type,
	"disable_renegotiation":           types.BoolType,
	"number_of_renegotiation_allowed": types.Int64Type,
	"prohibit_resume_on_reneg":        types.BoolType,
	"compression":                     types.BoolType,
	"allow_legacy_renegotiation":      types.BoolType,
	"prefer_server_ciphers":           types.BoolType,
	"elliptic_curves":                 types.ListType{ElemType: types.StringType},
	"prioritize_cha_cha":              types.BoolType,
	"sig_algs":                        types.ListType{ElemType: types.StringType},
	"require_closure_notification":    types.BoolType,
	"dependency_actions":              actions.ActionsListType,
}

func (data SSLServerProfile) GetPath() string {
	rest_path := "/mgmt/config/{domain}/SSLServerProfile"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data SSLServerProfile) IsNull() bool {
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
	if !data.RequestClientAuth.IsNull() {
		return false
	}
	if !data.RequireClientAuth.IsNull() {
		return false
	}
	if !data.ValidateClientCert.IsNull() {
		return false
	}
	if !data.SendClientAuthCaList.IsNull() {
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
	if data.SslOptions != nil {
		if !data.SslOptions.IsNull() {
			return false
		}
	}
	if !data.MaxSslDuration.IsNull() {
		return false
	}
	if !data.DisableRenegotiation.IsNull() {
		return false
	}
	if !data.NumberOfRenegotiationAllowed.IsNull() {
		return false
	}
	if !data.ProhibitResumeOnReneg.IsNull() {
		return false
	}
	if !data.Compression.IsNull() {
		return false
	}
	if !data.AllowLegacyRenegotiation.IsNull() {
		return false
	}
	if !data.PreferServerCiphers.IsNull() {
		return false
	}
	if !data.EllipticCurves.IsNull() {
		return false
	}
	if !data.PrioritizeChaCha.IsNull() {
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

func (data SSLServerProfile) ToBody(ctx context.Context, pathRoot string) string {
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
		var dataValues []string
		data.Ciphers.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.Set(body, pathRoot+`Ciphers`+".-1", map[string]string{"value": val})
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`Ciphers`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`Ciphers`, "[]")
	}
	if !data.Idcred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Idcred`, data.Idcred.ValueString())
	}
	if !data.RequestClientAuth.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RequestClientAuth`, tfutils.StringFromBool(data.RequestClientAuth, ""))
	}
	if !data.RequireClientAuth.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RequireClientAuth`, tfutils.StringFromBool(data.RequireClientAuth, ""))
	}
	if !data.ValidateClientCert.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ValidateClientCert`, tfutils.StringFromBool(data.ValidateClientCert, ""))
	}
	if !data.SendClientAuthCaList.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SendClientAuthCAList`, tfutils.StringFromBool(data.SendClientAuthCaList, ""))
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
	if data.SslOptions != nil {
		if !data.SslOptions.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`SSLOptions`, data.SslOptions.ToBody(ctx, ""))
		}
	}
	if !data.MaxSslDuration.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxSSLDuration`, data.MaxSslDuration.ValueInt64())
	}
	if !data.DisableRenegotiation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DisableRenegotiation`, tfutils.StringFromBool(data.DisableRenegotiation, ""))
	}
	if !data.NumberOfRenegotiationAllowed.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`NumberOfRenegotiationAllowed`, data.NumberOfRenegotiationAllowed.ValueInt64())
	}
	if !data.ProhibitResumeOnReneg.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ProhibitResumeOnReneg`, tfutils.StringFromBool(data.ProhibitResumeOnReneg, ""))
	}
	if !data.Compression.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Compression`, tfutils.StringFromBool(data.Compression, ""))
	}
	if !data.AllowLegacyRenegotiation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllowLegacyRenegotiation`, tfutils.StringFromBool(data.AllowLegacyRenegotiation, ""))
	}
	if !data.PreferServerCiphers.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PreferServerCiphers`, tfutils.StringFromBool(data.PreferServerCiphers, ""))
	}
	if !data.EllipticCurves.IsNull() {
		var dataValues []string
		data.EllipticCurves.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.Set(body, pathRoot+`EllipticCurves`+".-1", map[string]string{"value": val})
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`EllipticCurves`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`EllipticCurves`, "[]")
	}
	if !data.PrioritizeChaCha.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PrioritizeChaCha`, tfutils.StringFromBool(data.PrioritizeChaCha, ""))
	}
	if !data.SigAlgs.IsNull() {
		var dataValues []string
		data.SigAlgs.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.Set(body, pathRoot+`SigAlgs`+".-1", map[string]string{"value": val})
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`SigAlgs`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`SigAlgs`, "[]")
	}
	if !data.RequireClosureNotification.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RequireClosureNotification`, tfutils.StringFromBool(data.RequireClosureNotification, ""))
	}
	return body
}

func (data *SSLServerProfile) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `RequestClientAuth`); value.Exists() {
		data.RequestClientAuth = tfutils.BoolFromString(value.String())
	} else {
		data.RequestClientAuth = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RequireClientAuth`); value.Exists() {
		data.RequireClientAuth = tfutils.BoolFromString(value.String())
	} else {
		data.RequireClientAuth = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ValidateClientCert`); value.Exists() {
		data.ValidateClientCert = tfutils.BoolFromString(value.String())
	} else {
		data.ValidateClientCert = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SendClientAuthCAList`); value.Exists() {
		data.SendClientAuthCaList = tfutils.BoolFromString(value.String())
	} else {
		data.SendClientAuthCaList = types.BoolNull()
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
		data.CacheSize = types.Int64Value(20)
	}
	if value := res.Get(pathRoot + `SSLOptions`); value.Exists() {
		data.SslOptions = &DmSSLOptions{}
		data.SslOptions.FromBody(ctx, "", value)
	} else {
		data.SslOptions = nil
	}
	if value := res.Get(pathRoot + `MaxSSLDuration`); value.Exists() {
		data.MaxSslDuration = types.Int64Value(value.Int())
	} else {
		data.MaxSslDuration = types.Int64Value(60)
	}
	if value := res.Get(pathRoot + `DisableRenegotiation`); value.Exists() {
		data.DisableRenegotiation = tfutils.BoolFromString(value.String())
	} else {
		data.DisableRenegotiation = types.BoolNull()
	}
	if value := res.Get(pathRoot + `NumberOfRenegotiationAllowed`); value.Exists() {
		data.NumberOfRenegotiationAllowed = types.Int64Value(value.Int())
	} else {
		data.NumberOfRenegotiationAllowed = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `ProhibitResumeOnReneg`); value.Exists() {
		data.ProhibitResumeOnReneg = tfutils.BoolFromString(value.String())
	} else {
		data.ProhibitResumeOnReneg = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Compression`); value.Exists() {
		data.Compression = tfutils.BoolFromString(value.String())
	} else {
		data.Compression = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowLegacyRenegotiation`); value.Exists() {
		data.AllowLegacyRenegotiation = tfutils.BoolFromString(value.String())
	} else {
		data.AllowLegacyRenegotiation = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PreferServerCiphers`); value.Exists() {
		data.PreferServerCiphers = tfutils.BoolFromString(value.String())
	} else {
		data.PreferServerCiphers = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EllipticCurves`); value.Exists() {
		data.EllipticCurves = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.EllipticCurves = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `PrioritizeChaCha`); value.Exists() {
		data.PrioritizeChaCha = tfutils.BoolFromString(value.String())
	} else {
		data.PrioritizeChaCha = types.BoolNull()
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

func (data *SSLServerProfile) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `RequestClientAuth`); value.Exists() && !data.RequestClientAuth.IsNull() {
		data.RequestClientAuth = tfutils.BoolFromString(value.String())
	} else if data.RequestClientAuth.ValueBool() {
		data.RequestClientAuth = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RequireClientAuth`); value.Exists() && !data.RequireClientAuth.IsNull() {
		data.RequireClientAuth = tfutils.BoolFromString(value.String())
	} else if !data.RequireClientAuth.ValueBool() {
		data.RequireClientAuth = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ValidateClientCert`); value.Exists() && !data.ValidateClientCert.IsNull() {
		data.ValidateClientCert = tfutils.BoolFromString(value.String())
	} else if !data.ValidateClientCert.ValueBool() {
		data.ValidateClientCert = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SendClientAuthCAList`); value.Exists() && !data.SendClientAuthCaList.IsNull() {
		data.SendClientAuthCaList = tfutils.BoolFromString(value.String())
	} else if !data.SendClientAuthCaList.ValueBool() {
		data.SendClientAuthCaList = types.BoolNull()
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
	} else if data.CacheSize.ValueInt64() != 20 {
		data.CacheSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `SSLOptions`); value.Exists() {
		data.SslOptions.UpdateFromBody(ctx, "", value)
	} else {
		data.SslOptions = nil
	}
	if value := res.Get(pathRoot + `MaxSSLDuration`); value.Exists() && !data.MaxSslDuration.IsNull() {
		data.MaxSslDuration = types.Int64Value(value.Int())
	} else if data.MaxSslDuration.ValueInt64() != 60 {
		data.MaxSslDuration = types.Int64Null()
	}
	if value := res.Get(pathRoot + `DisableRenegotiation`); value.Exists() && !data.DisableRenegotiation.IsNull() {
		data.DisableRenegotiation = tfutils.BoolFromString(value.String())
	} else if data.DisableRenegotiation.ValueBool() {
		data.DisableRenegotiation = types.BoolNull()
	}
	if value := res.Get(pathRoot + `NumberOfRenegotiationAllowed`); value.Exists() && !data.NumberOfRenegotiationAllowed.IsNull() {
		data.NumberOfRenegotiationAllowed = types.Int64Value(value.Int())
	} else if data.NumberOfRenegotiationAllowed.ValueInt64() != 0 {
		data.NumberOfRenegotiationAllowed = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ProhibitResumeOnReneg`); value.Exists() && !data.ProhibitResumeOnReneg.IsNull() {
		data.ProhibitResumeOnReneg = tfutils.BoolFromString(value.String())
	} else if data.ProhibitResumeOnReneg.ValueBool() {
		data.ProhibitResumeOnReneg = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Compression`); value.Exists() && !data.Compression.IsNull() {
		data.Compression = tfutils.BoolFromString(value.String())
	} else if data.Compression.ValueBool() {
		data.Compression = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowLegacyRenegotiation`); value.Exists() && !data.AllowLegacyRenegotiation.IsNull() {
		data.AllowLegacyRenegotiation = tfutils.BoolFromString(value.String())
	} else if data.AllowLegacyRenegotiation.ValueBool() {
		data.AllowLegacyRenegotiation = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PreferServerCiphers`); value.Exists() && !data.PreferServerCiphers.IsNull() {
		data.PreferServerCiphers = tfutils.BoolFromString(value.String())
	} else if !data.PreferServerCiphers.ValueBool() {
		data.PreferServerCiphers = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EllipticCurves`); value.Exists() && !data.EllipticCurves.IsNull() {
		data.EllipticCurves = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.EllipticCurves = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `PrioritizeChaCha`); value.Exists() && !data.PrioritizeChaCha.IsNull() {
		data.PrioritizeChaCha = tfutils.BoolFromString(value.String())
	} else if data.PrioritizeChaCha.ValueBool() {
		data.PrioritizeChaCha = types.BoolNull()
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
