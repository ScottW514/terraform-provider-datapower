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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmAAAPAuthenticate struct {
	AuMethod                        types.String        `tfsdk:"au_method"`
	AuCustomUrl                     types.String        `tfsdk:"au_custom_url"`
	AuMapUrl                        types.String        `tfsdk:"au_map_url"`
	AuHost                          types.String        `tfsdk:"au_host"`
	AuPort                          types.Int64         `tfsdk:"au_port"`
	AusslValcred                    types.String        `tfsdk:"aussl_valcred"`
	AuCacheAllow                    types.String        `tfsdk:"au_cache_allow"`
	AuCacheTtl                      types.Int64         `tfsdk:"au_cache_ttl"`
	AuKerberosPrincipal             types.String        `tfsdk:"au_kerberos_principal"`
	AuKerberosPassword              types.String        `tfsdk:"au_kerberos_password"`
	AuClearTrustServerUrl           types.String        `tfsdk:"au_clear_trust_server_url"`
	AuClearTrustApplication         types.String        `tfsdk:"au_clear_trust_application"`
	AusamlArtifactResponder         types.String        `tfsdk:"ausaml_artifact_responder"`
	AuKerberosVerifySignature       types.Bool          `tfsdk:"au_kerberos_verify_signature"`
	AuNetegrityBaseUri              types.String        `tfsdk:"au_netegrity_base_uri"`
	AusamlAuthQueryServer           types.String        `tfsdk:"ausaml_auth_query_server"`
	AusamlVersion                   types.String        `tfsdk:"ausaml_version"`
	AuldapPrefix                    types.String        `tfsdk:"auldap_prefix"`
	AuldapSuffix                    types.String        `tfsdk:"auldap_suffix"`
	AuldapLoadBalanceGroup          types.String        `tfsdk:"auldap_load_balance_group"`
	AuKerberosKeytab                types.String        `tfsdk:"au_kerberos_keytab"`
	AuwsTrustUrl                    types.String        `tfsdk:"auws_trust_url"`
	Ausaml2Issuer                   types.String        `tfsdk:"ausaml2_issuer"`
	AuSignerValcred                 types.String        `tfsdk:"au_signer_valcred"`
	AuSignedXPath                   types.String        `tfsdk:"au_signed_x_path"`
	AuldapBindDn                    types.String        `tfsdk:"auldap_bind_dn"`
	AuldapBindPassword              types.String        `tfsdk:"auldap_bind_password"`
	AuldapSearchAttribute           types.String        `tfsdk:"auldap_search_attribute"`
	AultpaTokenVersionsBitmap       *DmLTPATokenVersion `tfsdk:"aultpa_token_versions_bitmap"`
	AultpaKeyFile                   types.String        `tfsdk:"aultpa_key_file"`
	AultpaKeyFilePassword           types.String        `tfsdk:"aultpa_key_file_password"`
	AultpaStashFile                 types.String        `tfsdk:"aultpa_stash_file"`
	AuBinaryTokenX509Valcred        types.String        `tfsdk:"au_binary_token_x509_valcred"`
	AutamServer                     types.String        `tfsdk:"autam_server"`
	AuAllowRemoteTokenReference     types.Bool          `tfsdk:"au_allow_remote_token_reference"`
	AuRemoteTokenProcessService     types.String        `tfsdk:"au_remote_token_process_service"`
	AuwsTrustVersion                types.String        `tfsdk:"auws_trust_version"`
	AuldapSearchForDn               types.Bool          `tfsdk:"auldap_search_for_dn"`
	AuldapSearchParameters          types.String        `tfsdk:"auldap_search_parameters"`
	AuwsTrustRequireClientEntropy   types.Bool          `tfsdk:"auws_trust_require_client_entropy"`
	AuwsTrustClientEntropySize      types.Int64         `tfsdk:"auws_trust_client_entropy_size"`
	AuwsTrustRequireServerEntropy   types.Bool          `tfsdk:"auws_trust_require_server_entropy"`
	AuwsTrustServerEntropySize      types.Int64         `tfsdk:"auws_trust_server_entropy_size"`
	AuwsTrustRequireRstc            types.Bool          `tfsdk:"auws_trust_require_rstc"`
	AuwsTrustRequireAppliesToHeader types.Bool          `tfsdk:"auws_trust_require_applies_to_header"`
	AuwsTrustAppliesToHeader        types.String        `tfsdk:"auws_trust_applies_to_header"`
	AuwsTrustEncryptionCertificate  types.String        `tfsdk:"auws_trust_encryption_certificate"`
	AuzosnssConfig                  types.String        `tfsdk:"auzosnss_config"`
	AuldapAttributes                types.String        `tfsdk:"auldap_attributes"`
	AuSkewTime                      types.Int64         `tfsdk:"au_skew_time"`
	AutampacReturn                  types.Bool          `tfsdk:"autampac_return"`
	AuldapReadTimeout               types.Int64         `tfsdk:"auldap_read_timeout"`
	AusslClientConfigType           types.String        `tfsdk:"aussl_client_config_type"`
	AusslClientProfile              types.String        `tfsdk:"aussl_client_profile"`
	AuldapBindPasswordAlias         types.String        `tfsdk:"auldap_bind_password_alias"`
	AultpaKeyFilePasswordAlias      types.String        `tfsdk:"aultpa_key_file_password_alias"`
	AusmRequestType                 types.String        `tfsdk:"ausm_request_type"`
	AusmCookieFlow                  *DmSMFlow           `tfsdk:"ausm_cookie_flow"`
	AusmHeaderFlow                  *DmSMFlow           `tfsdk:"ausm_header_flow"`
	AusmCookieAttributes            types.String        `tfsdk:"ausm_cookie_attributes"`
	AuCacheControl                  types.String        `tfsdk:"au_cache_control"`
}

var DmAAAPAuthenticateObjectType = map[string]attr.Type{
	"au_method":                            types.StringType,
	"au_custom_url":                        types.StringType,
	"au_map_url":                           types.StringType,
	"au_host":                              types.StringType,
	"au_port":                              types.Int64Type,
	"aussl_valcred":                        types.StringType,
	"au_cache_allow":                       types.StringType,
	"au_cache_ttl":                         types.Int64Type,
	"au_kerberos_principal":                types.StringType,
	"au_kerberos_password":                 types.StringType,
	"au_clear_trust_server_url":            types.StringType,
	"au_clear_trust_application":           types.StringType,
	"ausaml_artifact_responder":            types.StringType,
	"au_kerberos_verify_signature":         types.BoolType,
	"au_netegrity_base_uri":                types.StringType,
	"ausaml_auth_query_server":             types.StringType,
	"ausaml_version":                       types.StringType,
	"auldap_prefix":                        types.StringType,
	"auldap_suffix":                        types.StringType,
	"auldap_load_balance_group":            types.StringType,
	"au_kerberos_keytab":                   types.StringType,
	"auws_trust_url":                       types.StringType,
	"ausaml2_issuer":                       types.StringType,
	"au_signer_valcred":                    types.StringType,
	"au_signed_x_path":                     types.StringType,
	"auldap_bind_dn":                       types.StringType,
	"auldap_bind_password":                 types.StringType,
	"auldap_search_attribute":              types.StringType,
	"aultpa_token_versions_bitmap":         types.ObjectType{AttrTypes: DmLTPATokenVersionObjectType},
	"aultpa_key_file":                      types.StringType,
	"aultpa_key_file_password":             types.StringType,
	"aultpa_stash_file":                    types.StringType,
	"au_binary_token_x509_valcred":         types.StringType,
	"autam_server":                         types.StringType,
	"au_allow_remote_token_reference":      types.BoolType,
	"au_remote_token_process_service":      types.StringType,
	"auws_trust_version":                   types.StringType,
	"auldap_search_for_dn":                 types.BoolType,
	"auldap_search_parameters":             types.StringType,
	"auws_trust_require_client_entropy":    types.BoolType,
	"auws_trust_client_entropy_size":       types.Int64Type,
	"auws_trust_require_server_entropy":    types.BoolType,
	"auws_trust_server_entropy_size":       types.Int64Type,
	"auws_trust_require_rstc":              types.BoolType,
	"auws_trust_require_applies_to_header": types.BoolType,
	"auws_trust_applies_to_header":         types.StringType,
	"auws_trust_encryption_certificate":    types.StringType,
	"auzosnss_config":                      types.StringType,
	"auldap_attributes":                    types.StringType,
	"au_skew_time":                         types.Int64Type,
	"autampac_return":                      types.BoolType,
	"auldap_read_timeout":                  types.Int64Type,
	"aussl_client_config_type":             types.StringType,
	"aussl_client_profile":                 types.StringType,
	"auldap_bind_password_alias":           types.StringType,
	"aultpa_key_file_password_alias":       types.StringType,
	"ausm_request_type":                    types.StringType,
	"ausm_cookie_flow":                     types.ObjectType{AttrTypes: DmSMFlowObjectType},
	"ausm_header_flow":                     types.ObjectType{AttrTypes: DmSMFlowObjectType},
	"ausm_cookie_attributes":               types.StringType,
	"au_cache_control":                     types.StringType,
}
var DmAAAPAuthenticateObjectDefault = map[string]attr.Value{
	"au_method":                            types.StringValue("ldap"),
	"au_custom_url":                        types.StringNull(),
	"au_map_url":                           types.StringNull(),
	"au_host":                              types.StringNull(),
	"au_port":                              types.Int64Value(389),
	"aussl_valcred":                        types.StringNull(),
	"au_cache_allow":                       types.StringValue("absolute"),
	"au_cache_ttl":                         types.Int64Value(3),
	"au_kerberos_principal":                types.StringNull(),
	"au_kerberos_password":                 types.StringNull(),
	"au_clear_trust_server_url":            types.StringNull(),
	"au_clear_trust_application":           types.StringNull(),
	"ausaml_artifact_responder":            types.StringNull(),
	"au_kerberos_verify_signature":         types.BoolValue(true),
	"au_netegrity_base_uri":                types.StringNull(),
	"ausaml_auth_query_server":             types.StringNull(),
	"ausaml_version":                       types.StringValue("1.1"),
	"auldap_prefix":                        types.StringValue("cn="),
	"auldap_suffix":                        types.StringNull(),
	"auldap_load_balance_group":            types.StringNull(),
	"au_kerberos_keytab":                   types.StringNull(),
	"auws_trust_url":                       types.StringNull(),
	"ausaml2_issuer":                       types.StringNull(),
	"au_signer_valcred":                    types.StringNull(),
	"au_signed_x_path":                     types.StringNull(),
	"auldap_bind_dn":                       types.StringNull(),
	"auldap_bind_password":                 types.StringNull(),
	"auldap_search_attribute":              types.StringValue("userPassword"),
	"aultpa_token_versions_bitmap":         types.ObjectValueMust(DmLTPATokenVersionObjectType, DmLTPATokenVersionObjectDefault),
	"aultpa_key_file":                      types.StringNull(),
	"aultpa_key_file_password":             types.StringNull(),
	"aultpa_stash_file":                    types.StringNull(),
	"au_binary_token_x509_valcred":         types.StringNull(),
	"autam_server":                         types.StringNull(),
	"au_allow_remote_token_reference":      types.BoolValue(false),
	"au_remote_token_process_service":      types.StringNull(),
	"auws_trust_version":                   types.StringValue("1.2"),
	"auldap_search_for_dn":                 types.BoolValue(false),
	"auldap_search_parameters":             types.StringNull(),
	"auws_trust_require_client_entropy":    types.BoolValue(false),
	"auws_trust_client_entropy_size":       types.Int64Value(32),
	"auws_trust_require_server_entropy":    types.BoolValue(false),
	"auws_trust_server_entropy_size":       types.Int64Value(32),
	"auws_trust_require_rstc":              types.BoolValue(false),
	"auws_trust_require_applies_to_header": types.BoolValue(false),
	"auws_trust_applies_to_header":         types.StringNull(),
	"auws_trust_encryption_certificate":    types.StringNull(),
	"auzosnss_config":                      types.StringNull(),
	"auldap_attributes":                    types.StringNull(),
	"au_skew_time":                         types.Int64Value(0),
	"autampac_return":                      types.BoolValue(false),
	"auldap_read_timeout":                  types.Int64Value(60),
	"aussl_client_config_type":             types.StringValue("client"),
	"aussl_client_profile":                 types.StringNull(),
	"auldap_bind_password_alias":           types.StringNull(),
	"aultpa_key_file_password_alias":       types.StringNull(),
	"ausm_request_type":                    types.StringValue("webagent"),
	"ausm_cookie_flow":                     types.ObjectValueMust(DmSMFlowObjectType, DmSMFlowObjectDefault),
	"ausm_header_flow":                     types.ObjectValueMust(DmSMFlowObjectType, DmSMFlowObjectDefault),
	"ausm_cookie_attributes":               types.StringNull(),
	"au_cache_control":                     types.StringValue("default"),
}
var DmAAAPAuthenticateDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"au_method": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Method", "method", "").AddStringEnum("xmlfile", "ldap", "tivoli", "netegrity", "oblix", "cleartrust", "radius", "client-ssl", "validate-signer", "saml-signature", "saml-artifact", "saml-authen-query", "ws-trust", "ws-secureconversation", "token", "kerberos", "ltpa", "binarytokenx509", "zosnss", "verified-oauth", "custom").AddDefaultValue("ldap").String,
			Computed:            true,
		},
		"au_custom_url": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Custom URL", "custom-url", "").String,
			Computed:            true,
		},
		"au_map_url": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("AAA information file URL", "xmlfile-url", "").String,
			Computed:            true,
		},
		"au_host": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Host", "remote-host", "").String,
			Computed:            true,
		},
		"au_port": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Port", "remote-port", "").AddDefaultValue("389").String,
			Computed:            true,
		},
		"aussl_valcred": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("TLS client validation credentials", "valcred", "cryptovalcred").String,
			Computed:            true,
		},
		"au_cache_allow": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Cache authentication results", "cache-type", "").AddStringEnum("absolute", "disabled", "maximum", "minimum").AddDefaultValue("absolute").String,
			Computed:            true,
		},
		"au_cache_ttl": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Cache lifetime", "cache-ttl", "").AddIntegerRange(1, 86400).AddDefaultValue("3").String,
			Computed:            true,
		},
		"au_kerberos_principal": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Kerberos principal name", "kerberos-principal", "").String,
			Computed:            true,
		},
		"au_kerberos_password": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "").String,
			Computed:            true,
		},
		"au_clear_trust_server_url": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("ClearTrust server URL", "cleartrust-url", "").String,
			Computed:            true,
		},
		"au_clear_trust_application": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "").String,
			Computed:            true,
		},
		"ausaml_artifact_responder": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML Artifact responder", "saml-artifact-responder", "").String,
			Computed:            true,
		},
		"au_kerberos_verify_signature": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"au_netegrity_base_uri": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("CA Single Sign-On base URI", "netegrity-base-uri", "").String,
			Computed:            true,
		},
		"ausaml_auth_query_server": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML Authentication query server", "saml-authen-query-url", "").String,
			Computed:            true,
		},
		"ausaml_version": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML version", "saml-version", "").AddStringEnum("2.0", "1.1", "1.0").AddDefaultValue("1.1").String,
			Computed:            true,
		},
		"auldap_prefix": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP DN prefix", "ldap-prefix", "").AddDefaultValue("cn=").String,
			Computed:            true,
		},
		"auldap_suffix": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP DN suffix", "ldap-suffix", "").String,
			Computed:            true,
		},
		"auldap_load_balance_group": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP load balancer group", "ldap-lbgroup", "loadbalancergroup").String,
			Computed:            true,
		},
		"au_kerberos_keytab": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Kerberos keytab", "kerberos-keytab", "cryptokerberoskeytab").String,
			Computed:            true,
		},
		"auws_trust_url": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WS-Trust token server", "ws-trust-url", "").String,
			Computed:            true,
		},
		"ausaml2_issuer": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "").String,
			Computed:            true,
		},
		"au_signer_valcred": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Signature validation credentials", "valcred", "cryptovalcred").String,
			Computed:            true,
		},
		"au_signed_x_path": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("XPath expression", "signed-xpath", "").String,
			Computed:            true,
		},
		"auldap_bind_dn": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP bind DN", "ldap-bind-dn", "").String,
			Computed:            true,
		},
		"auldap_bind_password": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP bind password (deprecated)", "ldap-bind-password", "").String,
			Computed:            true,
		},
		"auldap_search_attribute": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP search attribute", "ldap-search-attr", "").AddDefaultValue("userPassword").String,
			Computed:            true,
		},
		"aultpa_token_versions_bitmap": GetDmLTPATokenVersionDataSourceSchema("Acceptable LTPA versions", "lpta-version", ""),
		"aultpa_key_file": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LTPA key file", "lpta-key-file", "").String,
			Computed:            true,
		},
		"aultpa_key_file_password": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LTPA key file password (deprecated)", "lpta-key-password", "").String,
			Computed:            true,
		},
		"aultpa_stash_file": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LTPA stash file", "lpta-stash-file", "").String,
			Computed:            true,
		},
		"au_binary_token_x509_valcred": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("X.509 BinarySecurityToken validation credentials", "x509-bin-token-valcred", "cryptovalcred").String,
			Computed:            true,
		},
		"autam_server": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("IBM Security Access Manager client", "tam", "tam").String,
			Computed:            true,
		},
		"au_allow_remote_token_reference": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Retrieve remote WS-Security token", "remote-token-allowed", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"au_remote_token_process_service": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL to process remote token", "remote-token-url", "").String,
			Computed:            true,
		},
		"auws_trust_version": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WS-Trust compatibility version", "ws-trust-version", "").AddStringEnum("1.3", "1.2", "1.1").AddDefaultValue("1.2").String,
			Computed:            true,
		},
		"auldap_search_for_dn": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Search for DN", "ldap-search-for-dn", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"auldap_search_parameters": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP search parameters", "ldap-search-param", "ldapsearchparameters").String,
			Computed:            true,
		},
		"auws_trust_require_client_entropy": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Require client entropy", "trust-require-client-entropy", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"auws_trust_client_entropy_size": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Client entropy size", "trust-client-entropy-size", "").AddIntegerRange(8, 128).AddDefaultValue("32").String,
			Computed:            true,
		},
		"auws_trust_require_server_entropy": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Require server entropy", "trust-require-server-entropy", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"auws_trust_server_entropy_size": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Server entropy size", "trust-server-entropy-size", "").AddIntegerRange(8, 128).AddDefaultValue("32").String,
			Computed:            true,
		},
		"auws_trust_require_rstc": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Require RequestSecurityTokenCollection", "trust-require-rstc", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"auws_trust_require_applies_to_header": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Require AppliesTo SOAP header", "trust-require-applies-to-header", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"auws_trust_applies_to_header": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("AppliesTo SOAP header", "trust-applies-to-header", "").String,
			Computed:            true,
		},
		"auws_trust_encryption_certificate": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WS-Trust encryption certificate", "trust-encryption-certificate", "cryptocertificate").String,
			Computed:            true,
		},
		"auzosnss_config": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("z/OS NSS client configuration", "zos-nss-au", "zosnssclient").String,
			Computed:            true,
		},
		"auldap_attributes": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("User auxiliary LDAP attributes", "au-ldap-attributes", "").String,
			Computed:            true,
		},
		"au_skew_time": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Skew time", "au-skew-time", "").AddDefaultValue("0").String,
			Computed:            true,
		},
		"autampac_return": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Return Privilege Attribute Certificate", "tam-pac-return", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"auldap_read_timeout": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP read timeout", "ldap-readtimeout", "").AddIntegerRange(0, 86400).AddDefaultValue("60").String,
			Computed:            true,
		},
		"aussl_client_config_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("TLS client type", "ssl-client-type", "").AddStringEnum("proxy", "client").AddDefaultValue("client").String,
			Computed:            true,
		},
		"aussl_client_profile": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "sslclientprofile").String,
			Computed:            true,
		},
		"auldap_bind_password_alias": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP bind password alias", "ldap-bind-password-alias", "passwordalias").String,
			Computed:            true,
		},
		"aultpa_key_file_password_alias": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LTPA key file password alias", "ltpa-key-password-alias", "passwordalias").String,
			Computed:            true,
		},
		"ausm_request_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Request type", "sm-request-type", "").AddStringEnum("webagent", "webservice").AddDefaultValue("webagent").String,
			Computed:            true,
		},
		"ausm_cookie_flow": GetDmSMFlowDataSourceSchema("Session cookie flow", "sm-cookie-flow", ""),
		"ausm_header_flow": GetDmSMFlowDataSourceSchema("CA Single Sign-On header flow", "sm-header-flow", ""),
		"ausm_cookie_attributes": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Cookie attribute policy", "cookie-attributes", "cookieattributepolicy").String,
			Computed:            true,
		},
		"au_cache_control": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Authentication caching", "cache-control", "").AddStringEnum("default", "disable-all", "disable-ldap-failures").AddDefaultValue("default").String,
			Computed:            true,
		},
	},
}
var DmAAAPAuthenticateResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmAAAPAuthenticateObjectType,
			DmAAAPAuthenticateObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"au_method": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Method", "method", "").AddStringEnum("xmlfile", "ldap", "tivoli", "netegrity", "oblix", "cleartrust", "radius", "client-ssl", "validate-signer", "saml-signature", "saml-artifact", "saml-authen-query", "ws-trust", "ws-secureconversation", "token", "kerberos", "ltpa", "binarytokenx509", "zosnss", "verified-oauth", "custom").AddDefaultValue("ldap").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("xmlfile", "ldap", "tivoli", "netegrity", "oblix", "cleartrust", "radius", "client-ssl", "validate-signer", "saml-signature", "saml-artifact", "saml-authen-query", "ws-trust", "ws-secureconversation", "token", "kerberos", "ltpa", "binarytokenx509", "zosnss", "verified-oauth", "custom"),
			},
			Default: stringdefault.StaticString("ldap"),
		},
		"au_custom_url": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Custom URL", "custom-url", "").String,
			Optional:            true,
		},
		"au_map_url": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("AAA information file URL", "xmlfile-url", "").String,
			Optional:            true,
		},
		"au_host": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Host", "remote-host", "").String,
			Optional:            true,
		},
		"au_port": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Port", "remote-port", "").AddDefaultValue("389").String,
			Computed:            true,
			Optional:            true,
			Default:             int64default.StaticInt64(389),
		},
		"aussl_valcred": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("TLS client validation credentials", "valcred", "cryptovalcred").String,
			Optional:            true,
		},
		"au_cache_allow": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Cache authentication results", "cache-type", "").AddStringEnum("absolute", "disabled", "maximum", "minimum").AddDefaultValue("absolute").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("absolute", "disabled", "maximum", "minimum"),
			},
			Default: stringdefault.StaticString("absolute"),
		},
		"au_cache_ttl": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Cache lifetime", "cache-ttl", "").AddIntegerRange(1, 86400).AddDefaultValue("3").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(1, 86400),
			},
			Default: int64default.StaticInt64(3),
		},
		"au_kerberos_principal": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Kerberos principal name", "kerberos-principal", "").String,
			Optional:            true,
		},
		"au_kerberos_password": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "").String,
			Optional:            true,
		},
		"au_clear_trust_server_url": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("ClearTrust server URL", "cleartrust-url", "").String,
			Optional:            true,
		},
		"au_clear_trust_application": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "").String,
			Optional:            true,
		},
		"ausaml_artifact_responder": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML Artifact responder", "saml-artifact-responder", "").String,
			Optional:            true,
		},
		"au_kerberos_verify_signature": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"au_netegrity_base_uri": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("CA Single Sign-On base URI", "netegrity-base-uri", "").String,
			Optional:            true,
		},
		"ausaml_auth_query_server": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML Authentication query server", "saml-authen-query-url", "").String,
			Optional:            true,
		},
		"ausaml_version": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML version", "saml-version", "").AddStringEnum("2.0", "1.1", "1.0").AddDefaultValue("1.1").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("2.0", "1.1", "1.0"),
			},
			Default: stringdefault.StaticString("1.1"),
		},
		"auldap_prefix": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP DN prefix", "ldap-prefix", "").AddDefaultValue("cn=").String,
			Computed:            true,
			Optional:            true,
			Default:             stringdefault.StaticString("cn="),
		},
		"auldap_suffix": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP DN suffix", "ldap-suffix", "").String,
			Optional:            true,
		},
		"auldap_load_balance_group": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP load balancer group", "ldap-lbgroup", "loadbalancergroup").String,
			Optional:            true,
		},
		"au_kerberos_keytab": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Kerberos keytab", "kerberos-keytab", "cryptokerberoskeytab").String,
			Optional:            true,
		},
		"auws_trust_url": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WS-Trust token server", "ws-trust-url", "").String,
			Optional:            true,
		},
		"ausaml2_issuer": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "").String,
			Optional:            true,
		},
		"au_signer_valcred": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Signature validation credentials", "valcred", "cryptovalcred").String,
			Optional:            true,
		},
		"au_signed_x_path": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("XPath expression", "signed-xpath", "").String,
			Optional:            true,
		},
		"auldap_bind_dn": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP bind DN", "ldap-bind-dn", "").String,
			Optional:            true,
		},
		"auldap_bind_password": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP bind password (deprecated)", "ldap-bind-password", "").String,
			Optional:            true,
		},
		"auldap_search_attribute": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP search attribute", "ldap-search-attr", "").AddDefaultValue("userPassword").String,
			Computed:            true,
			Optional:            true,
			Default:             stringdefault.StaticString("userPassword"),
		},
		"aultpa_token_versions_bitmap": GetDmLTPATokenVersionResourceSchema("Acceptable LTPA versions", "lpta-version", "", false),
		"aultpa_key_file": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LTPA key file", "lpta-key-file", "").String,
			Optional:            true,
		},
		"aultpa_key_file_password": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LTPA key file password (deprecated)", "lpta-key-password", "").String,
			Optional:            true,
		},
		"aultpa_stash_file": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LTPA stash file", "lpta-stash-file", "").String,
			Optional:            true,
		},
		"au_binary_token_x509_valcred": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("X.509 BinarySecurityToken validation credentials", "x509-bin-token-valcred", "cryptovalcred").String,
			Optional:            true,
		},
		"autam_server": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("IBM Security Access Manager client", "tam", "tam").String,
			Optional:            true,
		},
		"au_allow_remote_token_reference": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Retrieve remote WS-Security token", "remote-token-allowed", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"au_remote_token_process_service": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL to process remote token", "remote-token-url", "").String,
			Optional:            true,
		},
		"auws_trust_version": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WS-Trust compatibility version", "ws-trust-version", "").AddStringEnum("1.3", "1.2", "1.1").AddDefaultValue("1.2").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("1.3", "1.2", "1.1"),
			},
			Default: stringdefault.StaticString("1.2"),
		},
		"auldap_search_for_dn": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Search for DN", "ldap-search-for-dn", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"auldap_search_parameters": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP search parameters", "ldap-search-param", "ldapsearchparameters").String,
			Optional:            true,
		},
		"auws_trust_require_client_entropy": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Require client entropy", "trust-require-client-entropy", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"auws_trust_client_entropy_size": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Client entropy size", "trust-client-entropy-size", "").AddIntegerRange(8, 128).AddDefaultValue("32").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(8, 128),
			},
			Default: int64default.StaticInt64(32),
		},
		"auws_trust_require_server_entropy": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Require server entropy", "trust-require-server-entropy", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"auws_trust_server_entropy_size": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Server entropy size", "trust-server-entropy-size", "").AddIntegerRange(8, 128).AddDefaultValue("32").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(8, 128),
			},
			Default: int64default.StaticInt64(32),
		},
		"auws_trust_require_rstc": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Require RequestSecurityTokenCollection", "trust-require-rstc", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"auws_trust_require_applies_to_header": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Require AppliesTo SOAP header", "trust-require-applies-to-header", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"auws_trust_applies_to_header": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("AppliesTo SOAP header", "trust-applies-to-header", "").String,
			Optional:            true,
		},
		"auws_trust_encryption_certificate": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WS-Trust encryption certificate", "trust-encryption-certificate", "cryptocertificate").String,
			Optional:            true,
		},
		"auzosnss_config": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("z/OS NSS client configuration", "zos-nss-au", "zosnssclient").String,
			Optional:            true,
		},
		"auldap_attributes": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("User auxiliary LDAP attributes", "au-ldap-attributes", "").String,
			Optional:            true,
		},
		"au_skew_time": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Skew time", "au-skew-time", "").AddDefaultValue("0").String,
			Computed:            true,
			Optional:            true,
			Default:             int64default.StaticInt64(0),
		},
		"autampac_return": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Return Privilege Attribute Certificate", "tam-pac-return", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"auldap_read_timeout": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP read timeout", "ldap-readtimeout", "").AddIntegerRange(0, 86400).AddDefaultValue("60").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(0, 86400),
			},
			Default: int64default.StaticInt64(60),
		},
		"aussl_client_config_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("TLS client type", "ssl-client-type", "").AddStringEnum("proxy", "client").AddDefaultValue("client").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("proxy", "client"),
			},
			Default: stringdefault.StaticString("client"),
		},
		"aussl_client_profile": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "sslclientprofile").String,
			Optional:            true,
		},
		"auldap_bind_password_alias": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP bind password alias", "ldap-bind-password-alias", "passwordalias").String,
			Optional:            true,
		},
		"aultpa_key_file_password_alias": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LTPA key file password alias", "ltpa-key-password-alias", "passwordalias").String,
			Optional:            true,
		},
		"ausm_request_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Request type", "sm-request-type", "").AddStringEnum("webagent", "webservice").AddDefaultValue("webagent").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("webagent", "webservice"),
			},
			Default: stringdefault.StaticString("webagent"),
		},
		"ausm_cookie_flow": GetDmSMFlowResourceSchema("Session cookie flow", "sm-cookie-flow", "", false),
		"ausm_header_flow": GetDmSMFlowResourceSchema("CA Single Sign-On header flow", "sm-header-flow", "", false),
		"ausm_cookie_attributes": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Cookie attribute policy", "cookie-attributes", "cookieattributepolicy").String,
			Optional:            true,
		},
		"au_cache_control": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Authentication caching", "cache-control", "").AddStringEnum("default", "disable-all", "disable-ldap-failures").AddDefaultValue("default").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("default", "disable-all", "disable-ldap-failures"),
			},
			Default: stringdefault.StaticString("default"),
		},
	},
}

func (data DmAAAPAuthenticate) IsNull() bool {
	if !data.AuMethod.IsNull() {
		return false
	}
	if !data.AuCustomUrl.IsNull() {
		return false
	}
	if !data.AuMapUrl.IsNull() {
		return false
	}
	if !data.AuHost.IsNull() {
		return false
	}
	if !data.AuPort.IsNull() {
		return false
	}
	if !data.AusslValcred.IsNull() {
		return false
	}
	if !data.AuCacheAllow.IsNull() {
		return false
	}
	if !data.AuCacheTtl.IsNull() {
		return false
	}
	if !data.AuKerberosPrincipal.IsNull() {
		return false
	}
	if !data.AuKerberosPassword.IsNull() {
		return false
	}
	if !data.AuClearTrustServerUrl.IsNull() {
		return false
	}
	if !data.AuClearTrustApplication.IsNull() {
		return false
	}
	if !data.AusamlArtifactResponder.IsNull() {
		return false
	}
	if !data.AuKerberosVerifySignature.IsNull() {
		return false
	}
	if !data.AuNetegrityBaseUri.IsNull() {
		return false
	}
	if !data.AusamlAuthQueryServer.IsNull() {
		return false
	}
	if !data.AusamlVersion.IsNull() {
		return false
	}
	if !data.AuldapPrefix.IsNull() {
		return false
	}
	if !data.AuldapSuffix.IsNull() {
		return false
	}
	if !data.AuldapLoadBalanceGroup.IsNull() {
		return false
	}
	if !data.AuKerberosKeytab.IsNull() {
		return false
	}
	if !data.AuwsTrustUrl.IsNull() {
		return false
	}
	if !data.Ausaml2Issuer.IsNull() {
		return false
	}
	if !data.AuSignerValcred.IsNull() {
		return false
	}
	if !data.AuSignedXPath.IsNull() {
		return false
	}
	if !data.AuldapBindDn.IsNull() {
		return false
	}
	if !data.AuldapBindPassword.IsNull() {
		return false
	}
	if !data.AuldapSearchAttribute.IsNull() {
		return false
	}
	if data.AultpaTokenVersionsBitmap != nil {
		if !data.AultpaTokenVersionsBitmap.IsNull() {
			return false
		}
	}
	if !data.AultpaKeyFile.IsNull() {
		return false
	}
	if !data.AultpaKeyFilePassword.IsNull() {
		return false
	}
	if !data.AultpaStashFile.IsNull() {
		return false
	}
	if !data.AuBinaryTokenX509Valcred.IsNull() {
		return false
	}
	if !data.AutamServer.IsNull() {
		return false
	}
	if !data.AuAllowRemoteTokenReference.IsNull() {
		return false
	}
	if !data.AuRemoteTokenProcessService.IsNull() {
		return false
	}
	if !data.AuwsTrustVersion.IsNull() {
		return false
	}
	if !data.AuldapSearchForDn.IsNull() {
		return false
	}
	if !data.AuldapSearchParameters.IsNull() {
		return false
	}
	if !data.AuwsTrustRequireClientEntropy.IsNull() {
		return false
	}
	if !data.AuwsTrustClientEntropySize.IsNull() {
		return false
	}
	if !data.AuwsTrustRequireServerEntropy.IsNull() {
		return false
	}
	if !data.AuwsTrustServerEntropySize.IsNull() {
		return false
	}
	if !data.AuwsTrustRequireRstc.IsNull() {
		return false
	}
	if !data.AuwsTrustRequireAppliesToHeader.IsNull() {
		return false
	}
	if !data.AuwsTrustAppliesToHeader.IsNull() {
		return false
	}
	if !data.AuwsTrustEncryptionCertificate.IsNull() {
		return false
	}
	if !data.AuzosnssConfig.IsNull() {
		return false
	}
	if !data.AuldapAttributes.IsNull() {
		return false
	}
	if !data.AuSkewTime.IsNull() {
		return false
	}
	if !data.AutampacReturn.IsNull() {
		return false
	}
	if !data.AuldapReadTimeout.IsNull() {
		return false
	}
	if !data.AusslClientConfigType.IsNull() {
		return false
	}
	if !data.AusslClientProfile.IsNull() {
		return false
	}
	if !data.AuldapBindPasswordAlias.IsNull() {
		return false
	}
	if !data.AultpaKeyFilePasswordAlias.IsNull() {
		return false
	}
	if !data.AusmRequestType.IsNull() {
		return false
	}
	if data.AusmCookieFlow != nil {
		if !data.AusmCookieFlow.IsNull() {
			return false
		}
	}
	if data.AusmHeaderFlow != nil {
		if !data.AusmHeaderFlow.IsNull() {
			return false
		}
	}
	if !data.AusmCookieAttributes.IsNull() {
		return false
	}
	if !data.AuCacheControl.IsNull() {
		return false
	}
	return true
}
func GetDmAAAPAuthenticateDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmAAAPAuthenticateDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmAAAPAuthenticateDataSourceSchema
}

func GetDmAAAPAuthenticateResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmAAAPAuthenticateResourceSchema.Required = true
	} else {
		DmAAAPAuthenticateResourceSchema.Optional = true
		DmAAAPAuthenticateResourceSchema.Computed = true
	}
	DmAAAPAuthenticateResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmAAAPAuthenticateResourceSchema
}

func (data DmAAAPAuthenticate) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.AuMethod.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUMethod`, data.AuMethod.ValueString())
	}
	if !data.AuCustomUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUCustomURL`, data.AuCustomUrl.ValueString())
	}
	if !data.AuMapUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUMapURL`, data.AuMapUrl.ValueString())
	}
	if !data.AuHost.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUHost`, data.AuHost.ValueString())
	}
	if !data.AuPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUPort`, data.AuPort.ValueInt64())
	}
	if !data.AusslValcred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUSSLValcred`, data.AusslValcred.ValueString())
	}
	if !data.AuCacheAllow.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUCacheAllow`, data.AuCacheAllow.ValueString())
	}
	if !data.AuCacheTtl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUCacheTTL`, data.AuCacheTtl.ValueInt64())
	}
	if !data.AuKerberosPrincipal.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUKerberosPrincipal`, data.AuKerberosPrincipal.ValueString())
	}
	if !data.AuKerberosPassword.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUKerberosPassword`, data.AuKerberosPassword.ValueString())
	}
	if !data.AuClearTrustServerUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUClearTrustServerURL`, data.AuClearTrustServerUrl.ValueString())
	}
	if !data.AuClearTrustApplication.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUClearTrustApplication`, data.AuClearTrustApplication.ValueString())
	}
	if !data.AusamlArtifactResponder.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUSAMLArtifactResponder`, data.AusamlArtifactResponder.ValueString())
	}
	if !data.AuKerberosVerifySignature.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUKerberosVerifySignature`, tfutils.StringFromBool(data.AuKerberosVerifySignature, false))
	}
	if !data.AuNetegrityBaseUri.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUNetegrityBaseURI`, data.AuNetegrityBaseUri.ValueString())
	}
	if !data.AusamlAuthQueryServer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUSAMLAuthQueryServer`, data.AusamlAuthQueryServer.ValueString())
	}
	if !data.AusamlVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUSAMLVersion`, data.AusamlVersion.ValueString())
	}
	if !data.AuldapPrefix.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPPrefix`, data.AuldapPrefix.ValueString())
	}
	if !data.AuldapSuffix.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPSuffix`, data.AuldapSuffix.ValueString())
	}
	if !data.AuldapLoadBalanceGroup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPLoadBalanceGroup`, data.AuldapLoadBalanceGroup.ValueString())
	}
	if !data.AuKerberosKeytab.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUKerberosKeytab`, data.AuKerberosKeytab.ValueString())
	}
	if !data.AuwsTrustUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUWSTrustURL`, data.AuwsTrustUrl.ValueString())
	}
	if !data.Ausaml2Issuer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUSAML2Issuer`, data.Ausaml2Issuer.ValueString())
	}
	if !data.AuSignerValcred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUSignerValcred`, data.AuSignerValcred.ValueString())
	}
	if !data.AuSignedXPath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUSignedXPath`, data.AuSignedXPath.ValueString())
	}
	if !data.AuldapBindDn.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPBindDN`, data.AuldapBindDn.ValueString())
	}
	if !data.AuldapBindPassword.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPBindPassword`, data.AuldapBindPassword.ValueString())
	}
	if !data.AuldapSearchAttribute.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPSearchAttribute`, data.AuldapSearchAttribute.ValueString())
	}
	if data.AultpaTokenVersionsBitmap != nil {
		if !data.AultpaTokenVersionsBitmap.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`AULTPATokenVersionsBitmap`, data.AultpaTokenVersionsBitmap.ToBody(ctx, ""))
		}
	}
	if !data.AultpaKeyFile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULTPAKeyFile`, data.AultpaKeyFile.ValueString())
	}
	if !data.AultpaKeyFilePassword.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULTPAKeyFilePassword`, data.AultpaKeyFilePassword.ValueString())
	}
	if !data.AultpaStashFile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULTPAStashFile`, data.AultpaStashFile.ValueString())
	}
	if !data.AuBinaryTokenX509Valcred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUBinaryTokenX509Valcred`, data.AuBinaryTokenX509Valcred.ValueString())
	}
	if !data.AutamServer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUTAMServer`, data.AutamServer.ValueString())
	}
	if !data.AuAllowRemoteTokenReference.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUAllowRemoteTokenReference`, tfutils.StringFromBool(data.AuAllowRemoteTokenReference, false))
	}
	if !data.AuRemoteTokenProcessService.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AURemoteTokenProcessService`, data.AuRemoteTokenProcessService.ValueString())
	}
	if !data.AuwsTrustVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUWSTrustVersion`, data.AuwsTrustVersion.ValueString())
	}
	if !data.AuldapSearchForDn.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPSearchForDN`, tfutils.StringFromBool(data.AuldapSearchForDn, false))
	}
	if !data.AuldapSearchParameters.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPSearchParameters`, data.AuldapSearchParameters.ValueString())
	}
	if !data.AuwsTrustRequireClientEntropy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUWSTrustRequireClientEntropy`, tfutils.StringFromBool(data.AuwsTrustRequireClientEntropy, false))
	}
	if !data.AuwsTrustClientEntropySize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUWSTrustClientEntropySize`, data.AuwsTrustClientEntropySize.ValueInt64())
	}
	if !data.AuwsTrustRequireServerEntropy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUWSTrustRequireServerEntropy`, tfutils.StringFromBool(data.AuwsTrustRequireServerEntropy, false))
	}
	if !data.AuwsTrustServerEntropySize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUWSTrustServerEntropySize`, data.AuwsTrustServerEntropySize.ValueInt64())
	}
	if !data.AuwsTrustRequireRstc.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUWSTrustRequireRSTC`, tfutils.StringFromBool(data.AuwsTrustRequireRstc, false))
	}
	if !data.AuwsTrustRequireAppliesToHeader.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUWSTrustRequireAppliesToHeader`, tfutils.StringFromBool(data.AuwsTrustRequireAppliesToHeader, false))
	}
	if !data.AuwsTrustAppliesToHeader.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUWSTrustAppliesToHeader`, data.AuwsTrustAppliesToHeader.ValueString())
	}
	if !data.AuwsTrustEncryptionCertificate.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUWSTrustEncryptionCertificate`, data.AuwsTrustEncryptionCertificate.ValueString())
	}
	if !data.AuzosnssConfig.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUZOSNSSConfig`, data.AuzosnssConfig.ValueString())
	}
	if !data.AuldapAttributes.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPAttributes`, data.AuldapAttributes.ValueString())
	}
	if !data.AuSkewTime.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUSkewTime`, data.AuSkewTime.ValueInt64())
	}
	if !data.AutampacReturn.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUTAMPACReturn`, tfutils.StringFromBool(data.AutampacReturn, false))
	}
	if !data.AuldapReadTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPReadTimeout`, data.AuldapReadTimeout.ValueInt64())
	}
	if !data.AusslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUSSLClientConfigType`, data.AusslClientConfigType.ValueString())
	}
	if !data.AusslClientProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUSSLClientProfile`, data.AusslClientProfile.ValueString())
	}
	if !data.AuldapBindPasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPBindPasswordAlias`, data.AuldapBindPasswordAlias.ValueString())
	}
	if !data.AultpaKeyFilePasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULTPAKeyFilePasswordAlias`, data.AultpaKeyFilePasswordAlias.ValueString())
	}
	if !data.AusmRequestType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUSMRequestType`, data.AusmRequestType.ValueString())
	}
	if data.AusmCookieFlow != nil {
		if !data.AusmCookieFlow.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`AUSMCookieFlow`, data.AusmCookieFlow.ToBody(ctx, ""))
		}
	}
	if data.AusmHeaderFlow != nil {
		if !data.AusmHeaderFlow.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`AUSMHeaderFlow`, data.AusmHeaderFlow.ToBody(ctx, ""))
		}
	}
	if !data.AusmCookieAttributes.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUSMCookieAttributes`, data.AusmCookieAttributes.ValueString())
	}
	if !data.AuCacheControl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUCacheControl`, data.AuCacheControl.ValueString())
	}
	return body
}

func (data *DmAAAPAuthenticate) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `AUMethod`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuMethod = types.StringValue("ldap")
	}
	if value := res.Get(pathRoot + `AUCustomURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuCustomUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuCustomUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUMapURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuMapUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuMapUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUHost`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuHost = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuHost = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUPort`); value.Exists() {
		data.AuPort = types.Int64Value(value.Int())
	} else {
		data.AuPort = types.Int64Value(389)
	}
	if value := res.Get(pathRoot + `AUSSLValcred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AusslValcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AusslValcred = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUCacheAllow`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuCacheAllow = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuCacheAllow = types.StringValue("absolute")
	}
	if value := res.Get(pathRoot + `AUCacheTTL`); value.Exists() {
		data.AuCacheTtl = types.Int64Value(value.Int())
	} else {
		data.AuCacheTtl = types.Int64Value(3)
	}
	if value := res.Get(pathRoot + `AUKerberosPrincipal`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuKerberosPrincipal = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuKerberosPrincipal = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUKerberosPassword`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuKerberosPassword = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuKerberosPassword = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUClearTrustServerURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuClearTrustServerUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuClearTrustServerUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUClearTrustApplication`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuClearTrustApplication = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuClearTrustApplication = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSAMLArtifactResponder`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AusamlArtifactResponder = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AusamlArtifactResponder = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUKerberosVerifySignature`); value.Exists() {
		data.AuKerberosVerifySignature = tfutils.BoolFromString(value.String())
	} else {
		data.AuKerberosVerifySignature = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AUNetegrityBaseURI`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuNetegrityBaseUri = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuNetegrityBaseUri = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSAMLAuthQueryServer`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AusamlAuthQueryServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AusamlAuthQueryServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSAMLVersion`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AusamlVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AusamlVersion = types.StringValue("1.1")
	}
	if value := res.Get(pathRoot + `AULDAPPrefix`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuldapPrefix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuldapPrefix = types.StringValue("cn=")
	}
	if value := res.Get(pathRoot + `AULDAPSuffix`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuldapSuffix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuldapSuffix = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPLoadBalanceGroup`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuldapLoadBalanceGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuldapLoadBalanceGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUKerberosKeytab`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuKerberosKeytab = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuKerberosKeytab = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUWSTrustURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuwsTrustUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuwsTrustUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSAML2Issuer`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Ausaml2Issuer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Ausaml2Issuer = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSignerValcred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuSignerValcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuSignerValcred = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSignedXPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuSignedXPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuSignedXPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPBindDN`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuldapBindDn = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuldapBindDn = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPBindPassword`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuldapBindPassword = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuldapBindPassword = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPSearchAttribute`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuldapSearchAttribute = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuldapSearchAttribute = types.StringValue("userPassword")
	}
	if value := res.Get(pathRoot + `AULTPATokenVersionsBitmap`); value.Exists() {
		data.AultpaTokenVersionsBitmap = &DmLTPATokenVersion{}
		data.AultpaTokenVersionsBitmap.FromBody(ctx, "", value)
	} else {
		data.AultpaTokenVersionsBitmap = nil
	}
	if value := res.Get(pathRoot + `AULTPAKeyFile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AultpaKeyFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AultpaKeyFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULTPAKeyFilePassword`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AultpaKeyFilePassword = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AultpaKeyFilePassword = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULTPAStashFile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AultpaStashFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AultpaStashFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUBinaryTokenX509Valcred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuBinaryTokenX509Valcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuBinaryTokenX509Valcred = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUTAMServer`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AutamServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AutamServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUAllowRemoteTokenReference`); value.Exists() {
		data.AuAllowRemoteTokenReference = tfutils.BoolFromString(value.String())
	} else {
		data.AuAllowRemoteTokenReference = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AURemoteTokenProcessService`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuRemoteTokenProcessService = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuRemoteTokenProcessService = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUWSTrustVersion`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuwsTrustVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuwsTrustVersion = types.StringValue("1.2")
	}
	if value := res.Get(pathRoot + `AULDAPSearchForDN`); value.Exists() {
		data.AuldapSearchForDn = tfutils.BoolFromString(value.String())
	} else {
		data.AuldapSearchForDn = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AULDAPSearchParameters`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuldapSearchParameters = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuldapSearchParameters = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUWSTrustRequireClientEntropy`); value.Exists() {
		data.AuwsTrustRequireClientEntropy = tfutils.BoolFromString(value.String())
	} else {
		data.AuwsTrustRequireClientEntropy = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AUWSTrustClientEntropySize`); value.Exists() {
		data.AuwsTrustClientEntropySize = types.Int64Value(value.Int())
	} else {
		data.AuwsTrustClientEntropySize = types.Int64Value(32)
	}
	if value := res.Get(pathRoot + `AUWSTrustRequireServerEntropy`); value.Exists() {
		data.AuwsTrustRequireServerEntropy = tfutils.BoolFromString(value.String())
	} else {
		data.AuwsTrustRequireServerEntropy = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AUWSTrustServerEntropySize`); value.Exists() {
		data.AuwsTrustServerEntropySize = types.Int64Value(value.Int())
	} else {
		data.AuwsTrustServerEntropySize = types.Int64Value(32)
	}
	if value := res.Get(pathRoot + `AUWSTrustRequireRSTC`); value.Exists() {
		data.AuwsTrustRequireRstc = tfutils.BoolFromString(value.String())
	} else {
		data.AuwsTrustRequireRstc = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AUWSTrustRequireAppliesToHeader`); value.Exists() {
		data.AuwsTrustRequireAppliesToHeader = tfutils.BoolFromString(value.String())
	} else {
		data.AuwsTrustRequireAppliesToHeader = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AUWSTrustAppliesToHeader`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuwsTrustAppliesToHeader = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuwsTrustAppliesToHeader = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUWSTrustEncryptionCertificate`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuwsTrustEncryptionCertificate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuwsTrustEncryptionCertificate = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUZOSNSSConfig`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuzosnssConfig = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuzosnssConfig = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPAttributes`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuldapAttributes = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuldapAttributes = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSkewTime`); value.Exists() {
		data.AuSkewTime = types.Int64Value(value.Int())
	} else {
		data.AuSkewTime = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `AUTAMPACReturn`); value.Exists() {
		data.AutampacReturn = tfutils.BoolFromString(value.String())
	} else {
		data.AutampacReturn = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AULDAPReadTimeout`); value.Exists() {
		data.AuldapReadTimeout = types.Int64Value(value.Int())
	} else {
		data.AuldapReadTimeout = types.Int64Value(60)
	}
	if value := res.Get(pathRoot + `AUSSLClientConfigType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AusslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AusslClientConfigType = types.StringValue("client")
	}
	if value := res.Get(pathRoot + `AUSSLClientProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AusslClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AusslClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPBindPasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuldapBindPasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuldapBindPasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULTPAKeyFilePasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AultpaKeyFilePasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AultpaKeyFilePasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSMRequestType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AusmRequestType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AusmRequestType = types.StringValue("webagent")
	}
	if value := res.Get(pathRoot + `AUSMCookieFlow`); value.Exists() {
		data.AusmCookieFlow = &DmSMFlow{}
		data.AusmCookieFlow.FromBody(ctx, "", value)
	} else {
		data.AusmCookieFlow = nil
	}
	if value := res.Get(pathRoot + `AUSMHeaderFlow`); value.Exists() {
		data.AusmHeaderFlow = &DmSMFlow{}
		data.AusmHeaderFlow.FromBody(ctx, "", value)
	} else {
		data.AusmHeaderFlow = nil
	}
	if value := res.Get(pathRoot + `AUSMCookieAttributes`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AusmCookieAttributes = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AusmCookieAttributes = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUCacheControl`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuCacheControl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuCacheControl = types.StringValue("default")
	}
}

func (data *DmAAAPAuthenticate) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `AUMethod`); value.Exists() && !data.AuMethod.IsNull() {
		data.AuMethod = tfutils.ParseStringFromGJSON(value)
	} else if data.AuMethod.ValueString() != "ldap" {
		data.AuMethod = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUCustomURL`); value.Exists() && !data.AuCustomUrl.IsNull() {
		data.AuCustomUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuCustomUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUMapURL`); value.Exists() && !data.AuMapUrl.IsNull() {
		data.AuMapUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuMapUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUHost`); value.Exists() && !data.AuHost.IsNull() {
		data.AuHost = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuHost = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUPort`); value.Exists() && !data.AuPort.IsNull() {
		data.AuPort = types.Int64Value(value.Int())
	} else if data.AuPort.ValueInt64() != 389 {
		data.AuPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AUSSLValcred`); value.Exists() && !data.AusslValcred.IsNull() {
		data.AusslValcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AusslValcred = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUCacheAllow`); value.Exists() && !data.AuCacheAllow.IsNull() {
		data.AuCacheAllow = tfutils.ParseStringFromGJSON(value)
	} else if data.AuCacheAllow.ValueString() != "absolute" {
		data.AuCacheAllow = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUCacheTTL`); value.Exists() && !data.AuCacheTtl.IsNull() {
		data.AuCacheTtl = types.Int64Value(value.Int())
	} else if data.AuCacheTtl.ValueInt64() != 3 {
		data.AuCacheTtl = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AUKerberosPrincipal`); value.Exists() && !data.AuKerberosPrincipal.IsNull() {
		data.AuKerberosPrincipal = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuKerberosPrincipal = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUKerberosPassword`); value.Exists() && !data.AuKerberosPassword.IsNull() {
		data.AuKerberosPassword = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuKerberosPassword = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUClearTrustServerURL`); value.Exists() && !data.AuClearTrustServerUrl.IsNull() {
		data.AuClearTrustServerUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuClearTrustServerUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUClearTrustApplication`); value.Exists() && !data.AuClearTrustApplication.IsNull() {
		data.AuClearTrustApplication = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuClearTrustApplication = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSAMLArtifactResponder`); value.Exists() && !data.AusamlArtifactResponder.IsNull() {
		data.AusamlArtifactResponder = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AusamlArtifactResponder = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUKerberosVerifySignature`); value.Exists() && !data.AuKerberosVerifySignature.IsNull() {
		data.AuKerberosVerifySignature = tfutils.BoolFromString(value.String())
	} else if !data.AuKerberosVerifySignature.ValueBool() {
		data.AuKerberosVerifySignature = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AUNetegrityBaseURI`); value.Exists() && !data.AuNetegrityBaseUri.IsNull() {
		data.AuNetegrityBaseUri = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuNetegrityBaseUri = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSAMLAuthQueryServer`); value.Exists() && !data.AusamlAuthQueryServer.IsNull() {
		data.AusamlAuthQueryServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AusamlAuthQueryServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSAMLVersion`); value.Exists() && !data.AusamlVersion.IsNull() {
		data.AusamlVersion = tfutils.ParseStringFromGJSON(value)
	} else if data.AusamlVersion.ValueString() != "1.1" {
		data.AusamlVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPPrefix`); value.Exists() && !data.AuldapPrefix.IsNull() {
		data.AuldapPrefix = tfutils.ParseStringFromGJSON(value)
	} else if data.AuldapPrefix.ValueString() != "cn=" {
		data.AuldapPrefix = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPSuffix`); value.Exists() && !data.AuldapSuffix.IsNull() {
		data.AuldapSuffix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuldapSuffix = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPLoadBalanceGroup`); value.Exists() && !data.AuldapLoadBalanceGroup.IsNull() {
		data.AuldapLoadBalanceGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuldapLoadBalanceGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUKerberosKeytab`); value.Exists() && !data.AuKerberosKeytab.IsNull() {
		data.AuKerberosKeytab = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuKerberosKeytab = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUWSTrustURL`); value.Exists() && !data.AuwsTrustUrl.IsNull() {
		data.AuwsTrustUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuwsTrustUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSAML2Issuer`); value.Exists() && !data.Ausaml2Issuer.IsNull() {
		data.Ausaml2Issuer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Ausaml2Issuer = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSignerValcred`); value.Exists() && !data.AuSignerValcred.IsNull() {
		data.AuSignerValcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuSignerValcred = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSignedXPath`); value.Exists() && !data.AuSignedXPath.IsNull() {
		data.AuSignedXPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuSignedXPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPBindDN`); value.Exists() && !data.AuldapBindDn.IsNull() {
		data.AuldapBindDn = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuldapBindDn = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPBindPassword`); value.Exists() && !data.AuldapBindPassword.IsNull() {
		data.AuldapBindPassword = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuldapBindPassword = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPSearchAttribute`); value.Exists() && !data.AuldapSearchAttribute.IsNull() {
		data.AuldapSearchAttribute = tfutils.ParseStringFromGJSON(value)
	} else if data.AuldapSearchAttribute.ValueString() != "userPassword" {
		data.AuldapSearchAttribute = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULTPATokenVersionsBitmap`); value.Exists() {
		data.AultpaTokenVersionsBitmap.UpdateFromBody(ctx, "", value)
	} else {
		data.AultpaTokenVersionsBitmap = nil
	}
	if value := res.Get(pathRoot + `AULTPAKeyFile`); value.Exists() && !data.AultpaKeyFile.IsNull() {
		data.AultpaKeyFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AultpaKeyFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULTPAKeyFilePassword`); value.Exists() && !data.AultpaKeyFilePassword.IsNull() {
		data.AultpaKeyFilePassword = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AultpaKeyFilePassword = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULTPAStashFile`); value.Exists() && !data.AultpaStashFile.IsNull() {
		data.AultpaStashFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AultpaStashFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUBinaryTokenX509Valcred`); value.Exists() && !data.AuBinaryTokenX509Valcred.IsNull() {
		data.AuBinaryTokenX509Valcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuBinaryTokenX509Valcred = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUTAMServer`); value.Exists() && !data.AutamServer.IsNull() {
		data.AutamServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AutamServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUAllowRemoteTokenReference`); value.Exists() && !data.AuAllowRemoteTokenReference.IsNull() {
		data.AuAllowRemoteTokenReference = tfutils.BoolFromString(value.String())
	} else if data.AuAllowRemoteTokenReference.ValueBool() {
		data.AuAllowRemoteTokenReference = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AURemoteTokenProcessService`); value.Exists() && !data.AuRemoteTokenProcessService.IsNull() {
		data.AuRemoteTokenProcessService = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuRemoteTokenProcessService = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUWSTrustVersion`); value.Exists() && !data.AuwsTrustVersion.IsNull() {
		data.AuwsTrustVersion = tfutils.ParseStringFromGJSON(value)
	} else if data.AuwsTrustVersion.ValueString() != "1.2" {
		data.AuwsTrustVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPSearchForDN`); value.Exists() && !data.AuldapSearchForDn.IsNull() {
		data.AuldapSearchForDn = tfutils.BoolFromString(value.String())
	} else if data.AuldapSearchForDn.ValueBool() {
		data.AuldapSearchForDn = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AULDAPSearchParameters`); value.Exists() && !data.AuldapSearchParameters.IsNull() {
		data.AuldapSearchParameters = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuldapSearchParameters = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUWSTrustRequireClientEntropy`); value.Exists() && !data.AuwsTrustRequireClientEntropy.IsNull() {
		data.AuwsTrustRequireClientEntropy = tfutils.BoolFromString(value.String())
	} else if data.AuwsTrustRequireClientEntropy.ValueBool() {
		data.AuwsTrustRequireClientEntropy = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AUWSTrustClientEntropySize`); value.Exists() && !data.AuwsTrustClientEntropySize.IsNull() {
		data.AuwsTrustClientEntropySize = types.Int64Value(value.Int())
	} else if data.AuwsTrustClientEntropySize.ValueInt64() != 32 {
		data.AuwsTrustClientEntropySize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AUWSTrustRequireServerEntropy`); value.Exists() && !data.AuwsTrustRequireServerEntropy.IsNull() {
		data.AuwsTrustRequireServerEntropy = tfutils.BoolFromString(value.String())
	} else if data.AuwsTrustRequireServerEntropy.ValueBool() {
		data.AuwsTrustRequireServerEntropy = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AUWSTrustServerEntropySize`); value.Exists() && !data.AuwsTrustServerEntropySize.IsNull() {
		data.AuwsTrustServerEntropySize = types.Int64Value(value.Int())
	} else if data.AuwsTrustServerEntropySize.ValueInt64() != 32 {
		data.AuwsTrustServerEntropySize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AUWSTrustRequireRSTC`); value.Exists() && !data.AuwsTrustRequireRstc.IsNull() {
		data.AuwsTrustRequireRstc = tfutils.BoolFromString(value.String())
	} else if data.AuwsTrustRequireRstc.ValueBool() {
		data.AuwsTrustRequireRstc = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AUWSTrustRequireAppliesToHeader`); value.Exists() && !data.AuwsTrustRequireAppliesToHeader.IsNull() {
		data.AuwsTrustRequireAppliesToHeader = tfutils.BoolFromString(value.String())
	} else if data.AuwsTrustRequireAppliesToHeader.ValueBool() {
		data.AuwsTrustRequireAppliesToHeader = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AUWSTrustAppliesToHeader`); value.Exists() && !data.AuwsTrustAppliesToHeader.IsNull() {
		data.AuwsTrustAppliesToHeader = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuwsTrustAppliesToHeader = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUWSTrustEncryptionCertificate`); value.Exists() && !data.AuwsTrustEncryptionCertificate.IsNull() {
		data.AuwsTrustEncryptionCertificate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuwsTrustEncryptionCertificate = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUZOSNSSConfig`); value.Exists() && !data.AuzosnssConfig.IsNull() {
		data.AuzosnssConfig = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuzosnssConfig = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPAttributes`); value.Exists() && !data.AuldapAttributes.IsNull() {
		data.AuldapAttributes = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuldapAttributes = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSkewTime`); value.Exists() && !data.AuSkewTime.IsNull() {
		data.AuSkewTime = types.Int64Value(value.Int())
	} else if data.AuSkewTime.ValueInt64() != 0 {
		data.AuSkewTime = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AUTAMPACReturn`); value.Exists() && !data.AutampacReturn.IsNull() {
		data.AutampacReturn = tfutils.BoolFromString(value.String())
	} else if data.AutampacReturn.ValueBool() {
		data.AutampacReturn = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AULDAPReadTimeout`); value.Exists() && !data.AuldapReadTimeout.IsNull() {
		data.AuldapReadTimeout = types.Int64Value(value.Int())
	} else if data.AuldapReadTimeout.ValueInt64() != 60 {
		data.AuldapReadTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AUSSLClientConfigType`); value.Exists() && !data.AusslClientConfigType.IsNull() {
		data.AusslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else if data.AusslClientConfigType.ValueString() != "client" {
		data.AusslClientConfigType = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSSLClientProfile`); value.Exists() && !data.AusslClientProfile.IsNull() {
		data.AusslClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AusslClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPBindPasswordAlias`); value.Exists() && !data.AuldapBindPasswordAlias.IsNull() {
		data.AuldapBindPasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuldapBindPasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULTPAKeyFilePasswordAlias`); value.Exists() && !data.AultpaKeyFilePasswordAlias.IsNull() {
		data.AultpaKeyFilePasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AultpaKeyFilePasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSMRequestType`); value.Exists() && !data.AusmRequestType.IsNull() {
		data.AusmRequestType = tfutils.ParseStringFromGJSON(value)
	} else if data.AusmRequestType.ValueString() != "webagent" {
		data.AusmRequestType = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSMCookieFlow`); value.Exists() {
		data.AusmCookieFlow.UpdateFromBody(ctx, "", value)
	} else {
		data.AusmCookieFlow = nil
	}
	if value := res.Get(pathRoot + `AUSMHeaderFlow`); value.Exists() {
		data.AusmHeaderFlow.UpdateFromBody(ctx, "", value)
	} else {
		data.AusmHeaderFlow = nil
	}
	if value := res.Get(pathRoot + `AUSMCookieAttributes`); value.Exists() && !data.AusmCookieAttributes.IsNull() {
		data.AusmCookieAttributes = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AusmCookieAttributes = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUCacheControl`); value.Exists() && !data.AuCacheControl.IsNull() {
		data.AuCacheControl = tfutils.ParseStringFromGJSON(value)
	} else if data.AuCacheControl.ValueString() != "default" {
		data.AuCacheControl = types.StringNull()
	}
}
