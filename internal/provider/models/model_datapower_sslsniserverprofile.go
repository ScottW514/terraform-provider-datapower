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

type SSLSNIServerProfile struct {
	Id                           types.String              `tfsdk:"id"`
	AppDomain                    types.String              `tfsdk:"app_domain"`
	UserSummary                  types.String              `tfsdk:"user_summary"`
	Protocols                    *DmSSLProtoVersionsBitmap `tfsdk:"protocols"`
	SniServerMapping             types.String              `tfsdk:"sni_server_mapping"`
	SniServerDefault             types.String              `tfsdk:"sni_server_default"`
	SslOptions                   *DmSSLOptions             `tfsdk:"ssl_options"`
	MaxSslDuration               types.Int64               `tfsdk:"max_ssl_duration"`
	NumberOfRenegotiationAllowed types.Int64               `tfsdk:"number_of_renegotiation_allowed"`
}

var SSLSNIServerProfileObjectType = map[string]attr.Type{
	"id":                              types.StringType,
	"app_domain":                      types.StringType,
	"user_summary":                    types.StringType,
	"protocols":                       types.ObjectType{AttrTypes: DmSSLProtoVersionsBitmapObjectType},
	"sni_server_mapping":              types.StringType,
	"sni_server_default":              types.StringType,
	"ssl_options":                     types.ObjectType{AttrTypes: DmSSLOptionsObjectType},
	"max_ssl_duration":                types.Int64Type,
	"number_of_renegotiation_allowed": types.Int64Type,
}

func (data SSLSNIServerProfile) GetPath() string {
	rest_path := "/mgmt/config/{domain}/SSLSNIServerProfile"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return strings.TrimSuffix(rest_path, "/")
}

func (data SSLSNIServerProfile) IsNull() bool {
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
	if !data.SniServerMapping.IsNull() {
		return false
	}
	if !data.SniServerDefault.IsNull() {
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
	if !data.NumberOfRenegotiationAllowed.IsNull() {
		return false
	}
	return true
}

func (data SSLSNIServerProfile) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.SniServerMapping.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SNIServerMapping`, data.SniServerMapping.ValueString())
	}
	if !data.SniServerDefault.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SNIServerDefault`, data.SniServerDefault.ValueString())
	}
	if data.SslOptions != nil {
		if !data.SslOptions.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`SSLOptions`, data.SslOptions.ToBody(ctx, ""))
		}
	}
	if !data.MaxSslDuration.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxSSLDuration`, data.MaxSslDuration.ValueInt64())
	}
	if !data.NumberOfRenegotiationAllowed.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`NumberOfRenegotiationAllowed`, data.NumberOfRenegotiationAllowed.ValueInt64())
	}
	return body
}

func (data *SSLSNIServerProfile) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `SNIServerMapping`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SniServerMapping = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SniServerMapping = types.StringNull()
	}
	if value := res.Get(pathRoot + `SNIServerDefault`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SniServerDefault = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SniServerDefault = types.StringNull()
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
		data.MaxSslDuration = types.Int64Value(3600)
	}
	if value := res.Get(pathRoot + `NumberOfRenegotiationAllowed`); value.Exists() {
		data.NumberOfRenegotiationAllowed = types.Int64Value(value.Int())
	} else {
		data.NumberOfRenegotiationAllowed = types.Int64Null()
	}
}

func (data *SSLSNIServerProfile) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `SNIServerMapping`); value.Exists() && !data.SniServerMapping.IsNull() {
		data.SniServerMapping = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SniServerMapping = types.StringNull()
	}
	if value := res.Get(pathRoot + `SNIServerDefault`); value.Exists() && !data.SniServerDefault.IsNull() {
		data.SniServerDefault = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SniServerDefault = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLOptions`); value.Exists() {
		data.SslOptions.UpdateFromBody(ctx, "", value)
	} else {
		data.SslOptions = nil
	}
	if value := res.Get(pathRoot + `MaxSSLDuration`); value.Exists() && !data.MaxSslDuration.IsNull() {
		data.MaxSslDuration = types.Int64Value(value.Int())
	} else if data.MaxSslDuration.ValueInt64() != 3600 {
		data.MaxSslDuration = types.Int64Null()
	}
	if value := res.Get(pathRoot + `NumberOfRenegotiationAllowed`); value.Exists() && !data.NumberOfRenegotiationAllowed.IsNull() {
		data.NumberOfRenegotiationAllowed = types.Int64Value(value.Int())
	} else {
		data.NumberOfRenegotiationAllowed = types.Int64Null()
	}
}
