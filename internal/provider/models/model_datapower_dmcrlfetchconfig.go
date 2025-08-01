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

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	DataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	ResourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmCRLFetchConfig struct {
	Name                types.String `tfsdk:"name"`
	FetchType           types.String `tfsdk:"fetch_type"`
	IssuerValcred       types.String `tfsdk:"issuer_valcred"`
	RefreshInterval     types.Int64  `tfsdk:"refresh_interval"`
	Url                 types.String `tfsdk:"url"`
	RemoteAddress       types.String `tfsdk:"remote_address"`
	RemotePort          types.Int64  `tfsdk:"remote_port"`
	Dn                  types.String `tfsdk:"dn"`
	BindDn              types.String `tfsdk:"bind_dn"`
	BindPassAlias       types.String `tfsdk:"bind_pass_alias"`
	LdapVersion         types.String `tfsdk:"ldap_version"`
	LdapReadTimeout     types.Int64  `tfsdk:"ldap_read_timeout"`
	SslClientConfigType types.String `tfsdk:"ssl_client_config_type"`
	SslClient           types.String `tfsdk:"ssl_client"`
}

var DmCRLFetchConfigObjectType = map[string]attr.Type{
	"name":                   types.StringType,
	"fetch_type":             types.StringType,
	"issuer_valcred":         types.StringType,
	"refresh_interval":       types.Int64Type,
	"url":                    types.StringType,
	"remote_address":         types.StringType,
	"remote_port":            types.Int64Type,
	"dn":                     types.StringType,
	"bind_dn":                types.StringType,
	"bind_pass_alias":        types.StringType,
	"ldap_version":           types.StringType,
	"ldap_read_timeout":      types.Int64Type,
	"ssl_client_config_type": types.StringType,
	"ssl_client":             types.StringType,
}
var DmCRLFetchConfigObjectDefault = map[string]attr.Value{
	"name":                   types.StringNull(),
	"fetch_type":             types.StringNull(),
	"issuer_valcred":         types.StringNull(),
	"refresh_interval":       types.Int64Value(240),
	"url":                    types.StringNull(),
	"remote_address":         types.StringNull(),
	"remote_port":            types.Int64Value(389),
	"dn":                     types.StringNull(),
	"bind_dn":                types.StringNull(),
	"bind_pass_alias":        types.StringNull(),
	"ldap_version":           types.StringValue("v2"),
	"ldap_read_timeout":      types.Int64Value(60),
	"ssl_client_config_type": types.StringValue("client"),
	"ssl_client":             types.StringNull(),
}
var DmCRLFetchConfigDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Policy name", "", "").String,
			Computed:            true,
		},
		"fetch_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Protocol", "", "").AddStringEnum("http", "ldap").String,
			Computed:            true,
		},
		"issuer_valcred": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Issuer validation credentials", "issuer", "cryptovalcred").String,
			Computed:            true,
		},
		"refresh_interval": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Refresh interval", "refresh", "").AddIntegerRange(1, 1440).AddDefaultValue("240").String,
			Computed:            true,
		},
		"url": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Fetch URL", "fetch-url", "").String,
			Computed:            true,
		},
		"remote_address": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP server", "remote-address", "").String,
			Computed:            true,
		},
		"remote_port": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP port", "", "").AddDefaultValue("389").String,
			Computed:            true,
		},
		"dn": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP read DN", "read-dn", "").String,
			Computed:            true,
		},
		"bind_dn": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP bind DN", "bind-dn", "").String,
			Computed:            true,
		},
		"bind_pass_alias": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP bind password alias", "bind-pass-alias", "passwordalias").String,
			Computed:            true,
		},
		"ldap_version": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP version", "ldap-version", "").AddStringEnum("v2", "v3").AddDefaultValue("v2").String,
			Computed:            true,
		},
		"ldap_read_timeout": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP Read Timeout", "ldap-readtimeout", "").AddIntegerRange(0, 86400).AddDefaultValue("60").String,
			Computed:            true,
		},
		"ssl_client_config_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("TLS client type", "ssl-client-type", "").AddStringEnum("client").AddDefaultValue("client").String,
			Computed:            true,
		},
		"ssl_client": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "sslclientprofile").String,
			Computed:            true,
		},
	},
}
var DmCRLFetchConfigResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Policy name", "", "").String,
			Required:            true,
		},
		"fetch_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Protocol", "", "").AddStringEnum("http", "ldap").String,
			Required:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("http", "ldap"),
			},
		},
		"issuer_valcred": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Issuer validation credentials", "issuer", "cryptovalcred").String,
			Required:            true,
		},
		"refresh_interval": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Refresh interval", "refresh", "").AddIntegerRange(1, 1440).AddDefaultValue("240").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(1, 1440),
			},
			Default: int64default.StaticInt64(240),
		},
		"url": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Fetch URL", "fetch-url", "").String,
			Optional:            true,
		},
		"remote_address": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP server", "remote-address", "").String,
			Optional:            true,
		},
		"remote_port": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP port", "", "").AddDefaultValue("389").String,
			Computed:            true,
			Optional:            true,
			Default:             int64default.StaticInt64(389),
		},
		"dn": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP read DN", "read-dn", "").String,
			Optional:            true,
		},
		"bind_dn": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP bind DN", "bind-dn", "").String,
			Optional:            true,
		},
		"bind_pass_alias": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP bind password alias", "bind-pass-alias", "passwordalias").String,
			Optional:            true,
		},
		"ldap_version": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP version", "ldap-version", "").AddStringEnum("v2", "v3").AddDefaultValue("v2").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("v2", "v3"),
			},
			Default: stringdefault.StaticString("v2"),
		},
		"ldap_read_timeout": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP Read Timeout", "ldap-readtimeout", "").AddIntegerRange(0, 86400).AddDefaultValue("60").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(0, 86400),
			},
			Default: int64default.StaticInt64(60),
		},
		"ssl_client_config_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("TLS client type", "ssl-client-type", "").AddStringEnum("client").AddDefaultValue("client").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("client"),
			},
			Default: stringdefault.StaticString("client"),
		},
		"ssl_client": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "sslclientprofile").String,
			Optional:            true,
		},
	},
}

func (data DmCRLFetchConfig) IsNull() bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.FetchType.IsNull() {
		return false
	}
	if !data.IssuerValcred.IsNull() {
		return false
	}
	if !data.RefreshInterval.IsNull() {
		return false
	}
	if !data.Url.IsNull() {
		return false
	}
	if !data.RemoteAddress.IsNull() {
		return false
	}
	if !data.RemotePort.IsNull() {
		return false
	}
	if !data.Dn.IsNull() {
		return false
	}
	if !data.BindDn.IsNull() {
		return false
	}
	if !data.BindPassAlias.IsNull() {
		return false
	}
	if !data.LdapVersion.IsNull() {
		return false
	}
	if !data.LdapReadTimeout.IsNull() {
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

func (data DmCRLFetchConfig) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Name`, data.Name.ValueString())
	}
	if !data.FetchType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FetchType`, data.FetchType.ValueString())
	}
	if !data.IssuerValcred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IssuerValcred`, data.IssuerValcred.ValueString())
	}
	if !data.RefreshInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RefreshInterval`, data.RefreshInterval.ValueInt64())
	}
	if !data.Url.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`URL`, data.Url.ValueString())
	}
	if !data.RemoteAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemoteAddress`, data.RemoteAddress.ValueString())
	}
	if !data.RemotePort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemotePort`, data.RemotePort.ValueInt64())
	}
	if !data.Dn.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DN`, data.Dn.ValueString())
	}
	if !data.BindDn.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BindDN`, data.BindDn.ValueString())
	}
	if !data.BindPassAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BindPassAlias`, data.BindPassAlias.ValueString())
	}
	if !data.LdapVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPVersion`, data.LdapVersion.ValueString())
	}
	if !data.LdapReadTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPReadTimeout`, data.LdapReadTimeout.ValueInt64())
	}
	if !data.SslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClientConfigType`, data.SslClientConfigType.ValueString())
	}
	if !data.SslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClient`, data.SslClient.ValueString())
	}
	return body
}

func (data *DmCRLFetchConfig) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `FetchType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.FetchType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FetchType = types.StringNull()
	}
	if value := res.Get(pathRoot + `IssuerValcred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.IssuerValcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.IssuerValcred = types.StringNull()
	}
	if value := res.Get(pathRoot + `RefreshInterval`); value.Exists() {
		data.RefreshInterval = types.Int64Value(value.Int())
	} else {
		data.RefreshInterval = types.Int64Value(240)
	}
	if value := res.Get(pathRoot + `URL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Url = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Url = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemoteAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RemoteAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemotePort`); value.Exists() {
		data.RemotePort = types.Int64Value(value.Int())
	} else {
		data.RemotePort = types.Int64Value(389)
	}
	if value := res.Get(pathRoot + `DN`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Dn = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Dn = types.StringNull()
	}
	if value := res.Get(pathRoot + `BindDN`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.BindDn = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BindDn = types.StringNull()
	}
	if value := res.Get(pathRoot + `BindPassAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.BindPassAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BindPassAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPVersion`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapVersion = types.StringValue("v2")
	}
	if value := res.Get(pathRoot + `LDAPReadTimeout`); value.Exists() {
		data.LdapReadTimeout = types.Int64Value(value.Int())
	} else {
		data.LdapReadTimeout = types.Int64Value(60)
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

func (data *DmCRLFetchConfig) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && !data.Name.IsNull() {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `FetchType`); value.Exists() && !data.FetchType.IsNull() {
		data.FetchType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FetchType = types.StringNull()
	}
	if value := res.Get(pathRoot + `IssuerValcred`); value.Exists() && !data.IssuerValcred.IsNull() {
		data.IssuerValcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.IssuerValcred = types.StringNull()
	}
	if value := res.Get(pathRoot + `RefreshInterval`); value.Exists() && !data.RefreshInterval.IsNull() {
		data.RefreshInterval = types.Int64Value(value.Int())
	} else if data.RefreshInterval.ValueInt64() != 240 {
		data.RefreshInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `URL`); value.Exists() && !data.Url.IsNull() {
		data.Url = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Url = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemoteAddress`); value.Exists() && !data.RemoteAddress.IsNull() {
		data.RemoteAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemotePort`); value.Exists() && !data.RemotePort.IsNull() {
		data.RemotePort = types.Int64Value(value.Int())
	} else if data.RemotePort.ValueInt64() != 389 {
		data.RemotePort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `DN`); value.Exists() && !data.Dn.IsNull() {
		data.Dn = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Dn = types.StringNull()
	}
	if value := res.Get(pathRoot + `BindDN`); value.Exists() && !data.BindDn.IsNull() {
		data.BindDn = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BindDn = types.StringNull()
	}
	if value := res.Get(pathRoot + `BindPassAlias`); value.Exists() && !data.BindPassAlias.IsNull() {
		data.BindPassAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BindPassAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPVersion`); value.Exists() && !data.LdapVersion.IsNull() {
		data.LdapVersion = tfutils.ParseStringFromGJSON(value)
	} else if data.LdapVersion.ValueString() != "v2" {
		data.LdapVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPReadTimeout`); value.Exists() && !data.LdapReadTimeout.IsNull() {
		data.LdapReadTimeout = types.Int64Value(value.Int())
	} else if data.LdapReadTimeout.ValueInt64() != 60 {
		data.LdapReadTimeout = types.Int64Null()
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
