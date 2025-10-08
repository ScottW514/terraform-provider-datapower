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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	DataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	ResourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmAAAPEIBitmap struct {
	HttpBasicAuth        types.Bool `tfsdk:"http_basic_auth"`
	WssecUsername        types.Bool `tfsdk:"wssec_username"`
	WssecDerivedKey      types.Bool `tfsdk:"wssec_derived_key"`
	WssecBinaryToken     types.Bool `tfsdk:"wssec_binary_token"`
	WsSecureConversation types.Bool `tfsdk:"ws_secure_conversation"`
	WsTrust              types.Bool `tfsdk:"ws_trust"`
	Kerberos             types.Bool `tfsdk:"kerberos"`
	KerberosSpnego       types.Bool `tfsdk:"kerberos_spnego"`
	ClientSsl            types.Bool `tfsdk:"client_ssl"`
	SamlAttrName         types.Bool `tfsdk:"saml_attr_name"`
	SamlAuthenName       types.Bool `tfsdk:"saml_authen_name"`
	SamlArtifact         types.Bool `tfsdk:"saml_artifact"`
	ClientIpAddress      types.Bool `tfsdk:"client_ip_address"`
	SignerDn             types.Bool `tfsdk:"signer_dn"`
	Token                types.Bool `tfsdk:"token"`
	CookieToken          types.Bool `tfsdk:"cookie_token"`
	Ltpa                 types.Bool `tfsdk:"ltpa"`
	Metadata             types.Bool `tfsdk:"metadata"`
	Jwt                  types.Bool `tfsdk:"jwt"`
	Custom               types.Bool `tfsdk:"custom"`
	HtmlFormsAuth        types.Bool `tfsdk:"html_forms_auth"`
	SocialLogin          types.Bool `tfsdk:"social_login"`
	Oauth                types.Bool `tfsdk:"oauth"`
}

var DmAAAPEIBitmapObjectType = map[string]attr.Type{
	"http_basic_auth":        types.BoolType,
	"wssec_username":         types.BoolType,
	"wssec_derived_key":      types.BoolType,
	"wssec_binary_token":     types.BoolType,
	"ws_secure_conversation": types.BoolType,
	"ws_trust":               types.BoolType,
	"kerberos":               types.BoolType,
	"kerberos_spnego":        types.BoolType,
	"client_ssl":             types.BoolType,
	"saml_attr_name":         types.BoolType,
	"saml_authen_name":       types.BoolType,
	"saml_artifact":          types.BoolType,
	"client_ip_address":      types.BoolType,
	"signer_dn":              types.BoolType,
	"token":                  types.BoolType,
	"cookie_token":           types.BoolType,
	"ltpa":                   types.BoolType,
	"metadata":               types.BoolType,
	"jwt":                    types.BoolType,
	"custom":                 types.BoolType,
	"html_forms_auth":        types.BoolType,
	"social_login":           types.BoolType,
	"oauth":                  types.BoolType,
}
var DmAAAPEIBitmapObjectDefault = map[string]attr.Value{
	"http_basic_auth":        types.BoolValue(false),
	"wssec_username":         types.BoolValue(false),
	"wssec_derived_key":      types.BoolValue(false),
	"wssec_binary_token":     types.BoolValue(false),
	"ws_secure_conversation": types.BoolValue(false),
	"ws_trust":               types.BoolValue(false),
	"kerberos":               types.BoolValue(false),
	"kerberos_spnego":        types.BoolValue(false),
	"client_ssl":             types.BoolValue(false),
	"saml_attr_name":         types.BoolValue(false),
	"saml_authen_name":       types.BoolValue(false),
	"saml_artifact":          types.BoolValue(false),
	"client_ip_address":      types.BoolValue(false),
	"signer_dn":              types.BoolValue(false),
	"token":                  types.BoolValue(false),
	"cookie_token":           types.BoolValue(false),
	"ltpa":                   types.BoolValue(false),
	"metadata":               types.BoolValue(false),
	"jwt":                    types.BoolValue(false),
	"custom":                 types.BoolValue(false),
	"html_forms_auth":        types.BoolValue(false),
	"social_login":           types.BoolValue(false),
	"oauth":                  types.BoolValue(false),
}

func GetDmAAAPEIBitmapDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmAAAPEIBitmapDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"http_basic_auth": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "HTTP Authentication header",
				Computed:            true,
			},
			"wssec_username": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Password-carrying UsernameToken element from WS-Security header",
				Computed:            true,
			},
			"wssec_derived_key": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Derived-key UsernameToken element from WS-Security header",
				Computed:            true,
			},
			"wssec_binary_token": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "BinarySecurityToken element from WS-Security header",
				Computed:            true,
			},
			"ws_secure_conversation": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "WS-SecureConversation identifier",
				Computed:            true,
			},
			"ws_trust": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "WS-Trust Base or Supporting token",
				Computed:            true,
			},
			"kerberos": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Kerberos AP-REQ from WS-Security header",
				Computed:            true,
			},
			"kerberos_spnego": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Kerberos AP-REQ from SPNEGO token",
				Computed:            true,
			},
			"client_ssl": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Subject DN of TLS certificate from connection peer",
				Computed:            true,
			},
			"saml_attr_name": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Name from SAML attribute assertion",
				Computed:            true,
			},
			"saml_authen_name": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Name from SAML authentication assertion",
				Computed:            true,
			},
			"saml_artifact": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "SAML artifact",
				Computed:            true,
			},
			"client_ip_address": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Client IP address",
				Computed:            true,
			},
			"signer_dn": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Subject DN from certificate in message signature",
				Computed:            true,
			},
			"token": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Token extracted from message",
				Computed:            true,
			},
			"cookie_token": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Token extracted as cookie value",
				Computed:            true,
			},
			"ltpa": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "LTPA token",
				Computed:            true,
			},
			"metadata": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Processing metadata",
				Computed:            true,
			},
			"jwt": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "JWT",
				Computed:            true,
			},
			"custom": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Custom processing",
				Computed:            true,
			},
			"html_forms_auth": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "HTML forms-based authentication",
				Computed:            true,
			},
			"social_login": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Redirect to a social login provider",
				Computed:            true,
			},
			"oauth": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "OAuth",
				Computed:            true,
			},
		},
	}
	DmAAAPEIBitmapDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmAAAPEIBitmapDataSourceSchema
}
func GetDmAAAPEIBitmapResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmAAAPEIBitmapResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmAAAPEIBitmapObjectType,
				DmAAAPEIBitmapObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"http_basic_auth": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP Authentication header", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wssec_username": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Password-carrying UsernameToken element from WS-Security header", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wssec_derived_key": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Derived-key UsernameToken element from WS-Security header", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wssec_binary_token": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("BinarySecurityToken element from WS-Security header", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ws_secure_conversation": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WS-SecureConversation identifier", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ws_trust": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WS-Trust Base or Supporting token", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"kerberos": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Kerberos AP-REQ from WS-Security header", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"kerberos_spnego": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Kerberos AP-REQ from SPNEGO token", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"client_ssl": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Subject DN of TLS certificate from connection peer", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"saml_attr_name": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name from SAML attribute assertion", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"saml_authen_name": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name from SAML authentication assertion", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"saml_artifact": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SAML artifact", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"client_ip_address": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Client IP address", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"signer_dn": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Subject DN from certificate in message signature", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"token": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Token extracted from message", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"cookie_token": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Token extracted as cookie value", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ltpa": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LTPA token", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"metadata": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Processing metadata", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"jwt": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("JWT", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"custom": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Custom processing", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"html_forms_auth": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTML forms-based authentication", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"social_login": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Redirect to a social login provider", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"oauth": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("OAuth", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
	DmAAAPEIBitmapResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmAAAPEIBitmapResourceSchema.Required = true
	} else {
		DmAAAPEIBitmapResourceSchema.Optional = true
		DmAAAPEIBitmapResourceSchema.Computed = true
	}
	return DmAAAPEIBitmapResourceSchema
}

func (data DmAAAPEIBitmap) IsNull() bool {
	if !data.HttpBasicAuth.IsNull() {
		return false
	}
	if !data.WssecUsername.IsNull() {
		return false
	}
	if !data.WssecDerivedKey.IsNull() {
		return false
	}
	if !data.WssecBinaryToken.IsNull() {
		return false
	}
	if !data.WsSecureConversation.IsNull() {
		return false
	}
	if !data.WsTrust.IsNull() {
		return false
	}
	if !data.Kerberos.IsNull() {
		return false
	}
	if !data.KerberosSpnego.IsNull() {
		return false
	}
	if !data.ClientSsl.IsNull() {
		return false
	}
	if !data.SamlAttrName.IsNull() {
		return false
	}
	if !data.SamlAuthenName.IsNull() {
		return false
	}
	if !data.SamlArtifact.IsNull() {
		return false
	}
	if !data.ClientIpAddress.IsNull() {
		return false
	}
	if !data.SignerDn.IsNull() {
		return false
	}
	if !data.Token.IsNull() {
		return false
	}
	if !data.CookieToken.IsNull() {
		return false
	}
	if !data.Ltpa.IsNull() {
		return false
	}
	if !data.Metadata.IsNull() {
		return false
	}
	if !data.Jwt.IsNull() {
		return false
	}
	if !data.Custom.IsNull() {
		return false
	}
	if !data.HtmlFormsAuth.IsNull() {
		return false
	}
	if !data.SocialLogin.IsNull() {
		return false
	}
	if !data.Oauth.IsNull() {
		return false
	}
	return true
}

func (data DmAAAPEIBitmap) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.HttpBasicAuth.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`http-basic-auth`, tfutils.StringFromBool(data.HttpBasicAuth, ""))
	}
	if !data.WssecUsername.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`wssec-username`, tfutils.StringFromBool(data.WssecUsername, ""))
	}
	if !data.WssecDerivedKey.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`wssec-derived-key`, tfutils.StringFromBool(data.WssecDerivedKey, ""))
	}
	if !data.WssecBinaryToken.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`wssec-binary-token`, tfutils.StringFromBool(data.WssecBinaryToken, ""))
	}
	if !data.WsSecureConversation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ws-secure-conversation`, tfutils.StringFromBool(data.WsSecureConversation, ""))
	}
	if !data.WsTrust.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ws-trust`, tfutils.StringFromBool(data.WsTrust, ""))
	}
	if !data.Kerberos.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`kerberos`, tfutils.StringFromBool(data.Kerberos, ""))
	}
	if !data.KerberosSpnego.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`kerberos-spnego`, tfutils.StringFromBool(data.KerberosSpnego, ""))
	}
	if !data.ClientSsl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`client-ssl`, tfutils.StringFromBool(data.ClientSsl, ""))
	}
	if !data.SamlAttrName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`saml-attr-name`, tfutils.StringFromBool(data.SamlAttrName, ""))
	}
	if !data.SamlAuthenName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`saml-authen-name`, tfutils.StringFromBool(data.SamlAuthenName, ""))
	}
	if !data.SamlArtifact.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`saml-artifact`, tfutils.StringFromBool(data.SamlArtifact, ""))
	}
	if !data.ClientIpAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`client-ip-address`, tfutils.StringFromBool(data.ClientIpAddress, ""))
	}
	if !data.SignerDn.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`signer-dn`, tfutils.StringFromBool(data.SignerDn, ""))
	}
	if !data.Token.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`token`, tfutils.StringFromBool(data.Token, ""))
	}
	if !data.CookieToken.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`cookie-token`, tfutils.StringFromBool(data.CookieToken, ""))
	}
	if !data.Ltpa.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ltpa`, tfutils.StringFromBool(data.Ltpa, ""))
	}
	if !data.Metadata.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`metadata`, tfutils.StringFromBool(data.Metadata, ""))
	}
	if !data.Jwt.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`jwt`, tfutils.StringFromBool(data.Jwt, ""))
	}
	if !data.Custom.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`custom`, tfutils.StringFromBool(data.Custom, ""))
	}
	if !data.HtmlFormsAuth.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`html-forms-auth`, tfutils.StringFromBool(data.HtmlFormsAuth, ""))
	}
	if !data.SocialLogin.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`social-login`, tfutils.StringFromBool(data.SocialLogin, ""))
	}
	if !data.Oauth.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`oauth`, tfutils.StringFromBool(data.Oauth, ""))
	}
	return body
}

func (data *DmAAAPEIBitmap) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `http-basic-auth`); value.Exists() {
		data.HttpBasicAuth = tfutils.BoolFromString(value.String())
	} else {
		data.HttpBasicAuth = types.BoolNull()
	}
	if value := res.Get(pathRoot + `wssec-username`); value.Exists() {
		data.WssecUsername = tfutils.BoolFromString(value.String())
	} else {
		data.WssecUsername = types.BoolNull()
	}
	if value := res.Get(pathRoot + `wssec-derived-key`); value.Exists() {
		data.WssecDerivedKey = tfutils.BoolFromString(value.String())
	} else {
		data.WssecDerivedKey = types.BoolNull()
	}
	if value := res.Get(pathRoot + `wssec-binary-token`); value.Exists() {
		data.WssecBinaryToken = tfutils.BoolFromString(value.String())
	} else {
		data.WssecBinaryToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ws-secure-conversation`); value.Exists() {
		data.WsSecureConversation = tfutils.BoolFromString(value.String())
	} else {
		data.WsSecureConversation = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ws-trust`); value.Exists() {
		data.WsTrust = tfutils.BoolFromString(value.String())
	} else {
		data.WsTrust = types.BoolNull()
	}
	if value := res.Get(pathRoot + `kerberos`); value.Exists() {
		data.Kerberos = tfutils.BoolFromString(value.String())
	} else {
		data.Kerberos = types.BoolNull()
	}
	if value := res.Get(pathRoot + `kerberos-spnego`); value.Exists() {
		data.KerberosSpnego = tfutils.BoolFromString(value.String())
	} else {
		data.KerberosSpnego = types.BoolNull()
	}
	if value := res.Get(pathRoot + `client-ssl`); value.Exists() {
		data.ClientSsl = tfutils.BoolFromString(value.String())
	} else {
		data.ClientSsl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `saml-attr-name`); value.Exists() {
		data.SamlAttrName = tfutils.BoolFromString(value.String())
	} else {
		data.SamlAttrName = types.BoolNull()
	}
	if value := res.Get(pathRoot + `saml-authen-name`); value.Exists() {
		data.SamlAuthenName = tfutils.BoolFromString(value.String())
	} else {
		data.SamlAuthenName = types.BoolNull()
	}
	if value := res.Get(pathRoot + `saml-artifact`); value.Exists() {
		data.SamlArtifact = tfutils.BoolFromString(value.String())
	} else {
		data.SamlArtifact = types.BoolNull()
	}
	if value := res.Get(pathRoot + `client-ip-address`); value.Exists() {
		data.ClientIpAddress = tfutils.BoolFromString(value.String())
	} else {
		data.ClientIpAddress = types.BoolNull()
	}
	if value := res.Get(pathRoot + `signer-dn`); value.Exists() {
		data.SignerDn = tfutils.BoolFromString(value.String())
	} else {
		data.SignerDn = types.BoolNull()
	}
	if value := res.Get(pathRoot + `token`); value.Exists() {
		data.Token = tfutils.BoolFromString(value.String())
	} else {
		data.Token = types.BoolNull()
	}
	if value := res.Get(pathRoot + `cookie-token`); value.Exists() {
		data.CookieToken = tfutils.BoolFromString(value.String())
	} else {
		data.CookieToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ltpa`); value.Exists() {
		data.Ltpa = tfutils.BoolFromString(value.String())
	} else {
		data.Ltpa = types.BoolNull()
	}
	if value := res.Get(pathRoot + `metadata`); value.Exists() {
		data.Metadata = tfutils.BoolFromString(value.String())
	} else {
		data.Metadata = types.BoolNull()
	}
	if value := res.Get(pathRoot + `jwt`); value.Exists() {
		data.Jwt = tfutils.BoolFromString(value.String())
	} else {
		data.Jwt = types.BoolNull()
	}
	if value := res.Get(pathRoot + `custom`); value.Exists() {
		data.Custom = tfutils.BoolFromString(value.String())
	} else {
		data.Custom = types.BoolNull()
	}
	if value := res.Get(pathRoot + `html-forms-auth`); value.Exists() {
		data.HtmlFormsAuth = tfutils.BoolFromString(value.String())
	} else {
		data.HtmlFormsAuth = types.BoolNull()
	}
	if value := res.Get(pathRoot + `social-login`); value.Exists() {
		data.SocialLogin = tfutils.BoolFromString(value.String())
	} else {
		data.SocialLogin = types.BoolNull()
	}
	if value := res.Get(pathRoot + `oauth`); value.Exists() {
		data.Oauth = tfutils.BoolFromString(value.String())
	} else {
		data.Oauth = types.BoolNull()
	}
}

func (data *DmAAAPEIBitmap) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `http-basic-auth`); value.Exists() && !data.HttpBasicAuth.IsNull() {
		data.HttpBasicAuth = tfutils.BoolFromString(value.String())
	} else if data.HttpBasicAuth.ValueBool() {
		data.HttpBasicAuth = types.BoolNull()
	}
	if value := res.Get(pathRoot + `wssec-username`); value.Exists() && !data.WssecUsername.IsNull() {
		data.WssecUsername = tfutils.BoolFromString(value.String())
	} else if data.WssecUsername.ValueBool() {
		data.WssecUsername = types.BoolNull()
	}
	if value := res.Get(pathRoot + `wssec-derived-key`); value.Exists() && !data.WssecDerivedKey.IsNull() {
		data.WssecDerivedKey = tfutils.BoolFromString(value.String())
	} else if data.WssecDerivedKey.ValueBool() {
		data.WssecDerivedKey = types.BoolNull()
	}
	if value := res.Get(pathRoot + `wssec-binary-token`); value.Exists() && !data.WssecBinaryToken.IsNull() {
		data.WssecBinaryToken = tfutils.BoolFromString(value.String())
	} else if data.WssecBinaryToken.ValueBool() {
		data.WssecBinaryToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ws-secure-conversation`); value.Exists() && !data.WsSecureConversation.IsNull() {
		data.WsSecureConversation = tfutils.BoolFromString(value.String())
	} else if data.WsSecureConversation.ValueBool() {
		data.WsSecureConversation = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ws-trust`); value.Exists() && !data.WsTrust.IsNull() {
		data.WsTrust = tfutils.BoolFromString(value.String())
	} else if data.WsTrust.ValueBool() {
		data.WsTrust = types.BoolNull()
	}
	if value := res.Get(pathRoot + `kerberos`); value.Exists() && !data.Kerberos.IsNull() {
		data.Kerberos = tfutils.BoolFromString(value.String())
	} else if data.Kerberos.ValueBool() {
		data.Kerberos = types.BoolNull()
	}
	if value := res.Get(pathRoot + `kerberos-spnego`); value.Exists() && !data.KerberosSpnego.IsNull() {
		data.KerberosSpnego = tfutils.BoolFromString(value.String())
	} else if data.KerberosSpnego.ValueBool() {
		data.KerberosSpnego = types.BoolNull()
	}
	if value := res.Get(pathRoot + `client-ssl`); value.Exists() && !data.ClientSsl.IsNull() {
		data.ClientSsl = tfutils.BoolFromString(value.String())
	} else if data.ClientSsl.ValueBool() {
		data.ClientSsl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `saml-attr-name`); value.Exists() && !data.SamlAttrName.IsNull() {
		data.SamlAttrName = tfutils.BoolFromString(value.String())
	} else if data.SamlAttrName.ValueBool() {
		data.SamlAttrName = types.BoolNull()
	}
	if value := res.Get(pathRoot + `saml-authen-name`); value.Exists() && !data.SamlAuthenName.IsNull() {
		data.SamlAuthenName = tfutils.BoolFromString(value.String())
	} else if data.SamlAuthenName.ValueBool() {
		data.SamlAuthenName = types.BoolNull()
	}
	if value := res.Get(pathRoot + `saml-artifact`); value.Exists() && !data.SamlArtifact.IsNull() {
		data.SamlArtifact = tfutils.BoolFromString(value.String())
	} else if data.SamlArtifact.ValueBool() {
		data.SamlArtifact = types.BoolNull()
	}
	if value := res.Get(pathRoot + `client-ip-address`); value.Exists() && !data.ClientIpAddress.IsNull() {
		data.ClientIpAddress = tfutils.BoolFromString(value.String())
	} else if data.ClientIpAddress.ValueBool() {
		data.ClientIpAddress = types.BoolNull()
	}
	if value := res.Get(pathRoot + `signer-dn`); value.Exists() && !data.SignerDn.IsNull() {
		data.SignerDn = tfutils.BoolFromString(value.String())
	} else if data.SignerDn.ValueBool() {
		data.SignerDn = types.BoolNull()
	}
	if value := res.Get(pathRoot + `token`); value.Exists() && !data.Token.IsNull() {
		data.Token = tfutils.BoolFromString(value.String())
	} else if data.Token.ValueBool() {
		data.Token = types.BoolNull()
	}
	if value := res.Get(pathRoot + `cookie-token`); value.Exists() && !data.CookieToken.IsNull() {
		data.CookieToken = tfutils.BoolFromString(value.String())
	} else if data.CookieToken.ValueBool() {
		data.CookieToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ltpa`); value.Exists() && !data.Ltpa.IsNull() {
		data.Ltpa = tfutils.BoolFromString(value.String())
	} else if data.Ltpa.ValueBool() {
		data.Ltpa = types.BoolNull()
	}
	if value := res.Get(pathRoot + `metadata`); value.Exists() && !data.Metadata.IsNull() {
		data.Metadata = tfutils.BoolFromString(value.String())
	} else if data.Metadata.ValueBool() {
		data.Metadata = types.BoolNull()
	}
	if value := res.Get(pathRoot + `jwt`); value.Exists() && !data.Jwt.IsNull() {
		data.Jwt = tfutils.BoolFromString(value.String())
	} else if data.Jwt.ValueBool() {
		data.Jwt = types.BoolNull()
	}
	if value := res.Get(pathRoot + `custom`); value.Exists() && !data.Custom.IsNull() {
		data.Custom = tfutils.BoolFromString(value.String())
	} else if data.Custom.ValueBool() {
		data.Custom = types.BoolNull()
	}
	if value := res.Get(pathRoot + `html-forms-auth`); value.Exists() && !data.HtmlFormsAuth.IsNull() {
		data.HtmlFormsAuth = tfutils.BoolFromString(value.String())
	} else if data.HtmlFormsAuth.ValueBool() {
		data.HtmlFormsAuth = types.BoolNull()
	}
	if value := res.Get(pathRoot + `social-login`); value.Exists() && !data.SocialLogin.IsNull() {
		data.SocialLogin = tfutils.BoolFromString(value.String())
	} else if data.SocialLogin.ValueBool() {
		data.SocialLogin = types.BoolNull()
	}
	if value := res.Get(pathRoot + `oauth`); value.Exists() && !data.Oauth.IsNull() {
		data.Oauth = tfutils.BoolFromString(value.String())
	} else if data.Oauth.ValueBool() {
		data.Oauth = types.BoolNull()
	}
}
