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

type DmAAAPPostProcess struct {
	PpEnabled                             types.Bool           `tfsdk:"pp_enabled"`
	PpCustomUrl                           types.String         `tfsdk:"pp_custom_url"`
	PpsamlAuthAssertion                   types.Bool           `tfsdk:"ppsaml_auth_assertion"`
	PpsamlServerName                      types.String         `tfsdk:"ppsaml_server_name"`
	PpsamlNameQualifier                   types.String         `tfsdk:"ppsaml_name_qualifier"`
	PpKerberosTicket                      types.Bool           `tfsdk:"pp_kerberos_ticket"`
	PpKerberosClient                      types.String         `tfsdk:"pp_kerberos_client"`
	PpKerberosClientPassword              types.String         `tfsdk:"pp_kerberos_client_password"`
	PpKerberosServer                      types.String         `tfsdk:"pp_kerberos_server"`
	PpwsTrust                             types.Bool           `tfsdk:"ppws_trust"`
	PpTimestamp                           types.Bool           `tfsdk:"pp_timestamp"`
	PpTimestampExpiry                     types.Int64          `tfsdk:"pp_timestamp_expiry"`
	PpAllowRenewal                        types.Bool           `tfsdk:"pp_allow_renewal"`
	PpsamlVersion                         types.String         `tfsdk:"ppsaml_version"`
	PpsamlSendSlo                         types.Bool           `tfsdk:"ppsaml_send_slo"`
	PpsamlsloEndpoint                     types.String         `tfsdk:"ppsamlslo_endpoint"`
	PpwsUsernameToken                     types.Bool           `tfsdk:"ppws_username_token"`
	PpwsUsernameTokenPasswordType         types.String         `tfsdk:"ppws_username_token_password_type"`
	PpsamlValidity                        types.Int64          `tfsdk:"ppsaml_validity"`
	PpsamlSkew                            types.Int64          `tfsdk:"ppsaml_skew"`
	PpwsUsernameTokenIncludePwd           types.Bool           `tfsdk:"ppws_username_token_include_pwd"`
	Ppltpa                                types.Bool           `tfsdk:"ppltpa"`
	PpltpaVersion                         types.String         `tfsdk:"ppltpa_version"`
	PpltpaExpiry                          types.Int64          `tfsdk:"ppltpa_expiry"`
	PpltpaKeyFile                         types.String         `tfsdk:"ppltpa_key_file"`
	PpltpaKeyFilePassword                 types.String         `tfsdk:"ppltpa_key_file_password"`
	PpltpaStashFile                       types.String         `tfsdk:"ppltpa_stash_file"`
	PpKerberosSpnegoToken                 types.Bool           `tfsdk:"pp_kerberos_spnego_token"`
	PpKerberosBstValueType                types.String         `tfsdk:"pp_kerberos_bst_value_type"`
	PpsamlUseWsSec                        types.Bool           `tfsdk:"ppsaml_use_ws_sec"`
	PpKerberosClientKeytab                types.String         `tfsdk:"pp_kerberos_client_keytab"`
	PpUseWsSec                            types.Bool           `tfsdk:"pp_use_ws_sec"`
	PpActorRoleId                         types.String         `tfsdk:"pp_actor_role_id"`
	PptfimTokenMapping                    types.Bool           `tfsdk:"pptfim_token_mapping"`
	PptfimEndpoint                        types.String         `tfsdk:"pptfim_endpoint"`
	PpwsDerivedKeyUsernameToken           types.Bool           `tfsdk:"ppws_derived_key_username_token"`
	PpwsDerivedKeyUsernameTokenIterations types.Int64          `tfsdk:"ppws_derived_key_username_token_iterations"`
	PpwsUsernameTokenAllowReplacement     types.Bool           `tfsdk:"ppws_username_token_allow_replacement"`
	PptfimReplaceMethod                   types.String         `tfsdk:"pptfim_replace_method"`
	PptfimRetrieveMode                    types.String         `tfsdk:"pptfim_retrieve_mode"`
	PphmacSigningAlg                      types.String         `tfsdk:"pphmac_signing_alg"`
	PpSigningHashAlg                      types.String         `tfsdk:"pp_signing_hash_alg"`
	PpwsTrustHeader                       types.Bool           `tfsdk:"ppws_trust_header"`
	PpwsscKeySource                       types.String         `tfsdk:"ppwssc_key_source"`
	PpSharedSecretKey                     types.String         `tfsdk:"pp_shared_secret_key"`
	PpwsTrustRenewalWait                  types.Int64          `tfsdk:"ppws_trust_renewal_wait"`
	PpwsTrustNewInstance                  types.Bool           `tfsdk:"ppws_trust_new_instance"`
	PpwsTrustNewKey                       types.Bool           `tfsdk:"ppws_trust_new_key"`
	PpwsTrustNeverExpire                  types.Bool           `tfsdk:"ppws_trust_never_expire"`
	PpicrxToken                           types.Bool           `tfsdk:"ppicrx_token"`
	PpicrxUserRealm                       types.String         `tfsdk:"ppicrx_user_realm"`
	PpsamlIdentityProvider                types.Bool           `tfsdk:"ppsaml_identity_provider"`
	PpsamlProtocol                        types.String         `tfsdk:"ppsaml_protocol"`
	PpsamlResponseDestination             types.String         `tfsdk:"ppsaml_response_destination"`
	PpResultWrapup                        types.String         `tfsdk:"pp_result_wrapup"`
	PpsamlAssertionType                   *DmSAMLStatementType `tfsdk:"ppsaml_assertion_type"`
	PpsamlSubjectConfirm                  types.String         `tfsdk:"ppsaml_subject_confirm"`
	PpsamlNameId                          types.Bool           `tfsdk:"ppsaml_name_id"`
	PpsamlNameIdFormat                    types.String         `tfsdk:"ppsaml_name_id_format"`
	PpsamlRecipient                       types.String         `tfsdk:"ppsaml_recipient"`
	PpsamlAudience                        types.String         `tfsdk:"ppsaml_audience"`
	PpsamlOmitNotBefore                   types.Bool           `tfsdk:"ppsaml_omit_not_before"`
	PpOneTimeUse                          types.Bool           `tfsdk:"pp_one_time_use"`
	PpsamlProxy                           types.Bool           `tfsdk:"ppsaml_proxy"`
	PpsamlProxyAudience                   types.String         `tfsdk:"ppsaml_proxy_audience"`
	PpsamlProxyCount                      types.Int64          `tfsdk:"ppsaml_proxy_count"`
	PpsamlAuthzAction                     types.String         `tfsdk:"ppsaml_authz_action"`
	PpsamlAttributes                      types.String         `tfsdk:"ppsaml_attributes"`
	PpltpaInsertCookie                    types.Bool           `tfsdk:"ppltpa_insert_cookie"`
	PptampacPropagate                     types.Bool           `tfsdk:"pptampac_propagate"`
	PptamHeader                           types.String         `tfsdk:"pptam_header"`
	PptamHeaderSize                       types.Int64          `tfsdk:"pptam_header_size"`
	PpKerberosUseS4u2Proxy                types.Bool           `tfsdk:"pp_kerberos_use_s4u2_proxy"`
	PpCookieAttributes                    types.String         `tfsdk:"pp_cookie_attributes"`
	PpKerberosUseS4u2SelfAndS4u2Proxy     types.Bool           `tfsdk:"pp_kerberos_use_s4u2_self_and_s4u2_proxy"`
	PpKerberosClientSource                types.String         `tfsdk:"pp_kerberos_client_source"`
	PpKerberosSelf                        types.String         `tfsdk:"pp_kerberos_self"`
	PpKerberosSelfKeytab                  types.String         `tfsdk:"pp_kerberos_self_keytab"`
	PpKerberosClientCustomUrl             types.String         `tfsdk:"pp_kerberos_client_custom_url"`
	PpKerberosClientCtxVar                types.String         `tfsdk:"pp_kerberos_client_ctx_var"`
	PpKerberosServerSource                types.String         `tfsdk:"pp_kerberos_server_source"`
	PpKerberosServerCustomUrl             types.String         `tfsdk:"pp_kerberos_server_custom_url"`
	PpKerberosServerCtxVar                types.String         `tfsdk:"pp_kerberos_server_ctx_var"`
	PpsslClientConfigType                 types.String         `tfsdk:"ppssl_client_config_type"`
	PpsslClientProfile                    types.String         `tfsdk:"ppssl_client_profile"`
	PpltpaKeyFilePasswordAlias            types.String         `tfsdk:"ppltpa_key_file_password_alias"`
	Ppjwt                                 types.Bool           `tfsdk:"ppjwt"`
	PpjwtGenerator                        types.String         `tfsdk:"ppjwt_generator"`
}

var DmAAAPPostProcessObjectType = map[string]attr.Type{
	"pp_enabled":                                 types.BoolType,
	"pp_custom_url":                              types.StringType,
	"ppsaml_auth_assertion":                      types.BoolType,
	"ppsaml_server_name":                         types.StringType,
	"ppsaml_name_qualifier":                      types.StringType,
	"pp_kerberos_ticket":                         types.BoolType,
	"pp_kerberos_client":                         types.StringType,
	"pp_kerberos_client_password":                types.StringType,
	"pp_kerberos_server":                         types.StringType,
	"ppws_trust":                                 types.BoolType,
	"pp_timestamp":                               types.BoolType,
	"pp_timestamp_expiry":                        types.Int64Type,
	"pp_allow_renewal":                           types.BoolType,
	"ppsaml_version":                             types.StringType,
	"ppsaml_send_slo":                            types.BoolType,
	"ppsamlslo_endpoint":                         types.StringType,
	"ppws_username_token":                        types.BoolType,
	"ppws_username_token_password_type":          types.StringType,
	"ppsaml_validity":                            types.Int64Type,
	"ppsaml_skew":                                types.Int64Type,
	"ppws_username_token_include_pwd":            types.BoolType,
	"ppltpa":                                     types.BoolType,
	"ppltpa_version":                             types.StringType,
	"ppltpa_expiry":                              types.Int64Type,
	"ppltpa_key_file":                            types.StringType,
	"ppltpa_key_file_password":                   types.StringType,
	"ppltpa_stash_file":                          types.StringType,
	"pp_kerberos_spnego_token":                   types.BoolType,
	"pp_kerberos_bst_value_type":                 types.StringType,
	"ppsaml_use_ws_sec":                          types.BoolType,
	"pp_kerberos_client_keytab":                  types.StringType,
	"pp_use_ws_sec":                              types.BoolType,
	"pp_actor_role_id":                           types.StringType,
	"pptfim_token_mapping":                       types.BoolType,
	"pptfim_endpoint":                            types.StringType,
	"ppws_derived_key_username_token":            types.BoolType,
	"ppws_derived_key_username_token_iterations": types.Int64Type,
	"ppws_username_token_allow_replacement":      types.BoolType,
	"pptfim_replace_method":                      types.StringType,
	"pptfim_retrieve_mode":                       types.StringType,
	"pphmac_signing_alg":                         types.StringType,
	"pp_signing_hash_alg":                        types.StringType,
	"ppws_trust_header":                          types.BoolType,
	"ppwssc_key_source":                          types.StringType,
	"pp_shared_secret_key":                       types.StringType,
	"ppws_trust_renewal_wait":                    types.Int64Type,
	"ppws_trust_new_instance":                    types.BoolType,
	"ppws_trust_new_key":                         types.BoolType,
	"ppws_trust_never_expire":                    types.BoolType,
	"ppicrx_token":                               types.BoolType,
	"ppicrx_user_realm":                          types.StringType,
	"ppsaml_identity_provider":                   types.BoolType,
	"ppsaml_protocol":                            types.StringType,
	"ppsaml_response_destination":                types.StringType,
	"pp_result_wrapup":                           types.StringType,
	"ppsaml_assertion_type":                      types.ObjectType{AttrTypes: DmSAMLStatementTypeObjectType},
	"ppsaml_subject_confirm":                     types.StringType,
	"ppsaml_name_id":                             types.BoolType,
	"ppsaml_name_id_format":                      types.StringType,
	"ppsaml_recipient":                           types.StringType,
	"ppsaml_audience":                            types.StringType,
	"ppsaml_omit_not_before":                     types.BoolType,
	"pp_one_time_use":                            types.BoolType,
	"ppsaml_proxy":                               types.BoolType,
	"ppsaml_proxy_audience":                      types.StringType,
	"ppsaml_proxy_count":                         types.Int64Type,
	"ppsaml_authz_action":                        types.StringType,
	"ppsaml_attributes":                          types.StringType,
	"ppltpa_insert_cookie":                       types.BoolType,
	"pptampac_propagate":                         types.BoolType,
	"pptam_header":                               types.StringType,
	"pptam_header_size":                          types.Int64Type,
	"pp_kerberos_use_s4u2_proxy":                 types.BoolType,
	"pp_cookie_attributes":                       types.StringType,
	"pp_kerberos_use_s4u2_self_and_s4u2_proxy":   types.BoolType,
	"pp_kerberos_client_source":                  types.StringType,
	"pp_kerberos_self":                           types.StringType,
	"pp_kerberos_self_keytab":                    types.StringType,
	"pp_kerberos_client_custom_url":              types.StringType,
	"pp_kerberos_client_ctx_var":                 types.StringType,
	"pp_kerberos_server_source":                  types.StringType,
	"pp_kerberos_server_custom_url":              types.StringType,
	"pp_kerberos_server_ctx_var":                 types.StringType,
	"ppssl_client_config_type":                   types.StringType,
	"ppssl_client_profile":                       types.StringType,
	"ppltpa_key_file_password_alias":             types.StringType,
	"ppjwt":                                      types.BoolType,
	"ppjwt_generator":                            types.StringType,
}
var DmAAAPPostProcessObjectDefault = map[string]attr.Value{
	"pp_enabled":                                 types.BoolValue(false),
	"pp_custom_url":                              types.StringNull(),
	"ppsaml_auth_assertion":                      types.BoolValue(false),
	"ppsaml_server_name":                         types.StringValue("XS"),
	"ppsaml_name_qualifier":                      types.StringNull(),
	"pp_kerberos_ticket":                         types.BoolValue(false),
	"pp_kerberos_client":                         types.StringNull(),
	"pp_kerberos_client_password":                types.StringNull(),
	"pp_kerberos_server":                         types.StringNull(),
	"ppws_trust":                                 types.BoolValue(false),
	"pp_timestamp":                               types.BoolValue(true),
	"pp_timestamp_expiry":                        types.Int64Value(0),
	"pp_allow_renewal":                           types.BoolValue(false),
	"ppsaml_version":                             types.StringValue("2"),
	"ppsaml_send_slo":                            types.BoolValue(false),
	"ppsamlslo_endpoint":                         types.StringNull(),
	"ppws_username_token":                        types.BoolValue(false),
	"ppws_username_token_password_type":          types.StringValue("Digest"),
	"ppsaml_validity":                            types.Int64Value(0),
	"ppsaml_skew":                                types.Int64Value(0),
	"ppws_username_token_include_pwd":            types.BoolValue(true),
	"ppltpa":                                     types.BoolValue(false),
	"ppltpa_version":                             types.StringValue("LTPA2"),
	"ppltpa_expiry":                              types.Int64Value(600),
	"ppltpa_key_file":                            types.StringNull(),
	"ppltpa_key_file_password":                   types.StringNull(),
	"ppltpa_stash_file":                          types.StringNull(),
	"pp_kerberos_spnego_token":                   types.BoolValue(false),
	"pp_kerberos_bst_value_type":                 types.StringValue("http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ"),
	"ppsaml_use_ws_sec":                          types.BoolValue(false),
	"pp_kerberos_client_keytab":                  types.StringNull(),
	"pp_use_ws_sec":                              types.BoolValue(false),
	"pp_actor_role_id":                           types.StringNull(),
	"pptfim_token_mapping":                       types.BoolValue(false),
	"pptfim_endpoint":                            types.StringNull(),
	"ppws_derived_key_username_token":            types.BoolValue(false),
	"ppws_derived_key_username_token_iterations": types.Int64Value(1000),
	"ppws_username_token_allow_replacement":      types.BoolValue(false),
	"pptfim_replace_method":                      types.StringValue("all"),
	"pptfim_retrieve_mode":                       types.StringValue("CallTFIM"),
	"pphmac_signing_alg":                         types.StringValue("hmac-sha1"),
	"pp_signing_hash_alg":                        types.StringValue("sha1"),
	"ppws_trust_header":                          types.BoolValue(false),
	"ppwssc_key_source":                          types.StringValue("random"),
	"pp_shared_secret_key":                       types.StringNull(),
	"ppws_trust_renewal_wait":                    types.Int64Value(0),
	"ppws_trust_new_instance":                    types.BoolValue(false),
	"ppws_trust_new_key":                         types.BoolValue(false),
	"ppws_trust_never_expire":                    types.BoolValue(false),
	"ppicrx_token":                               types.BoolValue(false),
	"ppicrx_user_realm":                          types.StringNull(),
	"ppsaml_identity_provider":                   types.BoolValue(false),
	"ppsaml_protocol":                            types.StringValue("assertion"),
	"ppsaml_response_destination":                types.StringNull(),
	"pp_result_wrapup":                           types.StringValue("wssec-replace"),
	"ppsaml_assertion_type":                      types.ObjectValueMust(DmSAMLStatementTypeObjectType, DmSAMLStatementTypeObjectDefault),
	"ppsaml_subject_confirm":                     types.StringValue("bearer"),
	"ppsaml_name_id":                             types.BoolValue(true),
	"ppsaml_name_id_format":                      types.StringNull(),
	"ppsaml_recipient":                           types.StringNull(),
	"ppsaml_audience":                            types.StringNull(),
	"ppsaml_omit_not_before":                     types.BoolValue(false),
	"pp_one_time_use":                            types.BoolValue(false),
	"ppsaml_proxy":                               types.BoolValue(false),
	"ppsaml_proxy_audience":                      types.StringNull(),
	"ppsaml_proxy_count":                         types.Int64Value(0),
	"ppsaml_authz_action":                        types.StringValue("AllHTTP"),
	"ppsaml_attributes":                          types.StringNull(),
	"ppltpa_insert_cookie":                       types.BoolValue(true),
	"pptampac_propagate":                         types.BoolValue(false),
	"pptam_header":                               types.StringValue("iv-creds"),
	"pptam_header_size":                          types.Int64Value(0),
	"pp_kerberos_use_s4u2_proxy":                 types.BoolValue(false),
	"pp_cookie_attributes":                       types.StringNull(),
	"pp_kerberos_use_s4u2_self_and_s4u2_proxy":   types.BoolValue(false),
	"pp_kerberos_client_source":                  types.StringValue("mc-output"),
	"pp_kerberos_self":                           types.StringNull(),
	"pp_kerberos_self_keytab":                    types.StringNull(),
	"pp_kerberos_client_custom_url":              types.StringNull(),
	"pp_kerberos_client_ctx_var":                 types.StringNull(),
	"pp_kerberos_server_source":                  types.StringValue("as-is-string"),
	"pp_kerberos_server_custom_url":              types.StringNull(),
	"pp_kerberos_server_ctx_var":                 types.StringNull(),
	"ppssl_client_config_type":                   types.StringValue("proxy"),
	"ppssl_client_profile":                       types.StringNull(),
	"ppltpa_key_file_password_alias":             types.StringNull(),
	"ppjwt":                                      types.BoolValue(false),
	"ppjwt_generator":                            types.StringNull(),
}
var DmAAAPPostProcessDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"pp_enabled": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Run postprocessing custom processing", "custom-processing", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"pp_custom_url": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Custom processing", "custom-url", "").String,
			Computed:            true,
		},
		"ppsaml_auth_assertion": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Generate SAML assertion with SAML authentication statement", "saml-generate-assertion", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ppsaml_server_name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML Issuer identity", "saml-server-name", "").AddDefaultValue("XS").String,
			Computed:            true,
		},
		"ppsaml_name_qualifier": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML name qualifier", "saml-name-qualifier", "").String,
			Computed:            true,
		},
		"pp_kerberos_ticket": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Include a WS-Security Kerberos AP-REQ token", "kerberos-include-token", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"pp_kerberos_client": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Kerberos client principal", "kerberos-client-principal", "").String,
			Computed:            true,
		},
		"pp_kerberos_client_password": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "").String,
			Computed:            true,
		},
		"pp_kerberos_server": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Kerberos server principal", "kerberos-server", "").String,
			Computed:            true,
		},
		"ppws_trust": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Process WS-Trust SCT STS request", "ws-trust-generate-resp", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"pp_timestamp": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Output WS-Trust token time stamp", "ws-trust-add-timestamp", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"pp_timestamp_expiry": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Security context validity", "ws-trust-timestamp-expiry", "").AddIntegerRange(0, 31622400).AddDefaultValue("0").String,
			Computed:            true,
		},
		"pp_allow_renewal": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Allow WS-Trust token renewal", "ws-trust-allow-renewal", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ppsaml_version": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML version", "saml-version", "").AddStringEnum("2", "1.1", "1").AddDefaultValue("2").String,
			Computed:            true,
		},
		"ppsaml_send_slo": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Send SAML Single Logout request (SAML 2.0 only)", "saml-send-slo", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ppsamlslo_endpoint": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML SLO service URL", "saml-slo-endpoint", "").String,
			Computed:            true,
		},
		"ppws_username_token": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Add WS-Security UsernameToken", "wssec-add-user-name-token", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ppws_username_token_password_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WS-Security UsernameToken password type", "wssec-user-name-token-type", "").AddStringEnum("Text", "Digest").AddDefaultValue("Digest").String,
			Computed:            true,
		},
		"ppsaml_validity": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Assertion validity", "saml-validity", "").AddDefaultValue("0").String,
			Computed:            true,
		},
		"ppsaml_skew": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Skew time", "saml-skew", "").AddDefaultValue("0").String,
			Computed:            true,
		},
		"ppws_username_token_include_pwd": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Include password in UsernameToken", "wssec-user-name-token-contains-pwd", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"ppltpa": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Generate LTPA token", "lpta-generate-token", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ppltpa_version": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LTPA token version", "lpta-version", "").AddStringEnum("LTPA", "LTPA1FIPS", "LTPA2", "LTPA2WAS7", "LTPADomino").AddDefaultValue("LTPA2").String,
			Computed:            true,
		},
		"ppltpa_expiry": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LTPA token expiry", "lpta-expiry", "").AddIntegerRange(1, 628992000).AddDefaultValue("600").String,
			Computed:            true,
		},
		"ppltpa_key_file": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LTPA key file", "lpta-key-file", "").String,
			Computed:            true,
		},
		"ppltpa_key_file_password": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LTPA key file password", "lpta-key-file-password", "").String,
			Computed:            true,
		},
		"ppltpa_stash_file": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LTPA stash file", "lpta-stash-file", "").String,
			Computed:            true,
		},
		"pp_kerberos_spnego_token": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Generate Kerberos SPNEGO token", "kerberos-generate-spnego", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"pp_kerberos_bst_value_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("ValueType for generated Kerberos BinarySecurityToken", "kerberos-value-type", "").AddStringEnum("http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ1510", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ1510", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ4120", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ4120", "http://docs.oasis-open.org/wss/2005/xx/oasis-2005xx-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ", "http://docs.oasis-open.org/wss/2005/xx/oasis-2005xx-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ", "http://www.docs.oasis-open.org/wss/2004/07/oasis-000000-wss-kerberos-token-profile-1.0#Kerberosv5_AP_REQ").AddDefaultValue("http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ").String,
			Computed:            true,
		},
		"ppsaml_use_ws_sec": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Wrap SAML assertion in WS-Security Security header", "saml-in-wssec", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"pp_kerberos_client_keytab": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Kerberos client keytab", "kerberos-client-keytab", "cryptokerberoskeytab").String,
			Computed:            true,
		},
		"pp_use_ws_sec": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Wrap token in WS-Security Security header", "wssec-header-wrap-token", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"pp_actor_role_id": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Actor or role identifier", "wssec-actor-role-id", "").String,
			Computed:            true,
		},
		"pptfim_token_mapping": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Request Federated Identity Manager token mapping (deprecated)", "tfim-token-mapping", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"pptfim_endpoint": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Federated Identity Manager endpoint", "tfim-endpoint", "tfimendpoint").String,
			Computed:            true,
		},
		"ppws_derived_key_username_token": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Use Derived-Key variant of WS-Security UsernameToken", "wssec-use-derived-key", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ppws_derived_key_username_token_iterations": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Hashing iteration count", "wssec-derived-key-hash-iter", "").AddIntegerRange(2, 65535).AddDefaultValue("1000").String,
			Computed:            true,
		},
		"ppws_username_token_allow_replacement": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Replace existing UsernameToken", "wssec-replace-existing", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"pptfim_replace_method": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Replacement method (deprecated)", "tfim-replace-method", "").AddStringEnum("all", "replace", "preserve").AddDefaultValue("all").String,
			Computed:            true,
		},
		"pptfim_retrieve_mode": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Retrieval method (deprecated)", "tfim-retrieval-method", "").AddStringEnum("CallTFIM", "FromMC").AddDefaultValue("CallTFIM").String,
			Computed:            true,
		},
		"pphmac_signing_alg": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("HMAC signing algorithm", "hmac-signing-algorithm", "").AddStringEnum("hmac-sha1", "hmac-sha224", "hmac-sha256", "hmac-sha384", "hmac-sha512", "hmac-ripemd160", "hmac-md5").AddDefaultValue("hmac-sha1").String,
			Computed:            true,
		},
		"pp_signing_hash_alg": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Signing message digest algorithm", "message-digest-algorithm", "").AddStringEnum("sha1", "sha256", "sha512", "ripemd160", "sha224", "sha384", "md5").AddDefaultValue("sha1").String,
			Computed:            true,
		},
		"ppws_trust_header": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Return the WS-Trust token as SOAP header", "ws-trust-in-header", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ppwssc_key_source": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Source of shared secret to initialize SecurityContext", "ws-trust-key-source", "").AddStringEnum("client-entropy", "in-kerberos", "in-encryptedkey", "static", "random").AddDefaultValue("random").String,
			Computed:            true,
		},
		"pp_shared_secret_key": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Shared secret key", "ws-trust-shared-key", "cryptosskey").String,
			Computed:            true,
		},
		"ppws_trust_renewal_wait": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Wait time for renewal", "ws-trust-renewal-wait", "").AddIntegerRange(0, 2678400).AddDefaultValue("0").String,
			Computed:            true,
		},
		"ppws_trust_new_instance": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Issue new Instance for WS-Trust renewal", "ws-trust-new-instance", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ppws_trust_new_key": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Update context key for WS-Trust renewal", "ws-trust-new-key", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ppws_trust_never_expire": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WS-Trust SecurityContext never expires", "ws-trust-never-expire", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ppicrx_token": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Generate ICRX token for z/OS identity propagation", "generate-icrx", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ppicrx_user_realm": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("ICRX realm", "icrx-user-realm", "").String,
			Computed:            true,
		},
		"ppsaml_identity_provider": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Generate SAML assertion or response", "generate-saml-assertion", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ppsaml_protocol": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML protocol or profile", "saml-protocol", "").AddStringEnum("assertion", "response").AddDefaultValue("assertion").String,
			Computed:            true,
		},
		"ppsaml_response_destination": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Response destination", "saml-response-destination", "").String,
			Computed:            true,
		},
		"pp_result_wrapup": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Wrap up result", "result-wrapup", "").AddStringEnum("wssec-replace", "wssec-new", "wssec-inject", "soap-body", "none").AddDefaultValue("wssec-replace").String,
			Computed:            true,
		},
		"ppsaml_assertion_type": GetDmSAMLStatementTypeDataSourceSchema("SAML assertion type", "saml-assertion-type", ""),
		"ppsaml_subject_confirm": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML subject confirmation method", "saml-subject-confirm", "").AddStringEnum("bearer", "hok", "sv").AddDefaultValue("bearer").String,
			Computed:            true,
		},
		"ppsaml_name_id": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML subject contains name identifier", "saml-nid", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"ppsaml_name_id_format": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML name identifier format", "saml-nid-format", "").String,
			Computed:            true,
		},
		"ppsaml_recipient": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML recipient", "saml-recipient", "").String,
			Computed:            true,
		},
		"ppsaml_audience": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML audience", "saml-audience", "").String,
			Computed:            true,
		},
		"ppsaml_omit_not_before": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Omit NotBefore attribute", "saml-omit-notbefore", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"pp_one_time_use": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("One time use only", "one-time-use", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ppsaml_proxy": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Allow SAML ProxyRestriction", "saml-proxy", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ppsaml_proxy_audience": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML proxy audience", "saml-proxy-audience", "").String,
			Computed:            true,
		},
		"ppsaml_proxy_count": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML proxy count", "saml-proxy-count", "").AddIntegerRange(0, 65535).AddDefaultValue("0").String,
			Computed:            true,
		},
		"ppsaml_authz_action": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Action for SAML Authorization decision", "saml-authz-action", "").AddStringEnum("AllHTTP", "POST", "GET", "PUT", "HEAD", "General", "Read", "Write", "Execute", "Delete", "Control", "NegatedRead", "NegatedWrite", "NegatedExecute", "NegatedDelete", "NegatedControl").AddDefaultValue("AllHTTP").String,
			Computed:            true,
		},
		"ppsaml_attributes": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML attribute definition", "saml-attributes", "samlattributes").String,
			Computed:            true,
		},
		"ppltpa_insert_cookie": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Insert LTPA Set-Cookie", "ltpa-insert-cookie", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"pptampac_propagate": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Generate an Access Manager PAC token", "propagate-tam-pac", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"pptam_header": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Privilege Attribute Certificate header name", "tam-header", "").AddDefaultValue("iv-creds").String,
			Computed:            true,
		},
		"pptam_header_size": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("PAC header value size", "tam-header-size", "").AddDefaultValue("0").String,
			Computed:            true,
		},
		"pp_kerberos_use_s4u2_proxy": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Use constrained delegation", "kerberos-use-s4u2proxy", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"pp_cookie_attributes": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Cookie attribute policy", "cookie-attributes", "cookieattributepolicy").String,
			Computed:            true,
		},
		"pp_kerberos_use_s4u2_self_and_s4u2_proxy": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Use protocol transition and constrained delegation", "kerberos-use-s4u2self", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"pp_kerberos_client_source": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Kerberos client principal source", "kerberos-client-source", "").AddStringEnum("mc-output", "custom-url", "ctx-var").AddDefaultValue("mc-output").String,
			Computed:            true,
		},
		"pp_kerberos_self": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Kerberos self principal", "kerberos-self-principal", "").String,
			Computed:            true,
		},
		"pp_kerberos_self_keytab": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Kerberos self keytab", "kerberos-self-keytab", "cryptokerberoskeytab").String,
			Computed:            true,
		},
		"pp_kerberos_client_custom_url": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Kerberos client principal - custom processing", "kerberos-client-custom-url", "").String,
			Computed:            true,
		},
		"pp_kerberos_client_ctx_var": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Kerberos client principal - context variable", "kerberos-client-ctx-var", "").String,
			Computed:            true,
		},
		"pp_kerberos_server_source": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Kerberos server principal source", "kerberos-server-source", "").AddStringEnum("as-is-string", "custom-url", "ctx-var").AddDefaultValue("as-is-string").String,
			Computed:            true,
		},
		"pp_kerberos_server_custom_url": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Kerberos server principal - custom processing", "kerberos-server-custom-url", "").String,
			Computed:            true,
		},
		"pp_kerberos_server_ctx_var": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Kerberos server principal - context variable", "kerberos-server-ctx-var", "").String,
			Computed:            true,
		},
		"ppssl_client_config_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("TLS client type", "ssl-client-type", "").AddStringEnum("proxy", "client").AddDefaultValue("proxy").String,
			Computed:            true,
		},
		"ppssl_client_profile": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "sslclientprofile").String,
			Computed:            true,
		},
		"ppltpa_key_file_password_alias": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LTPA key file password alias", "ltpa-key-file-password-alias", "passwordalias").String,
			Computed:            true,
		},
		"ppjwt": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Generate a JWT token", "jwt", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ppjwt_generator": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("JWT Generator settings", "generate-jwt", "aaajwtgenerator").String,
			Computed:            true,
		},
	},
}
var DmAAAPPostProcessResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmAAAPPostProcessObjectType,
			DmAAAPPostProcessObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"pp_enabled": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Run postprocessing custom processing", "custom-processing", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"pp_custom_url": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Custom processing", "custom-url", "").String,
			Optional:            true,
		},
		"ppsaml_auth_assertion": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Generate SAML assertion with SAML authentication statement", "saml-generate-assertion", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ppsaml_server_name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML Issuer identity", "saml-server-name", "").AddDefaultValue("XS").String,
			Computed:            true,
			Optional:            true,
			Default:             stringdefault.StaticString("XS"),
		},
		"ppsaml_name_qualifier": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML name qualifier", "saml-name-qualifier", "").String,
			Optional:            true,
		},
		"pp_kerberos_ticket": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Include a WS-Security Kerberos AP-REQ token", "kerberos-include-token", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"pp_kerberos_client": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Kerberos client principal", "kerberos-client-principal", "").String,
			Optional:            true,
		},
		"pp_kerberos_client_password": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "").String,
			Optional:            true,
		},
		"pp_kerberos_server": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Kerberos server principal", "kerberos-server", "").String,
			Optional:            true,
		},
		"ppws_trust": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Process WS-Trust SCT STS request", "ws-trust-generate-resp", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"pp_timestamp": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Output WS-Trust token time stamp", "ws-trust-add-timestamp", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"pp_timestamp_expiry": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Security context validity", "ws-trust-timestamp-expiry", "").AddIntegerRange(0, 31622400).AddDefaultValue("0").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(0, 31622400),
			},
			Default: int64default.StaticInt64(0),
		},
		"pp_allow_renewal": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Allow WS-Trust token renewal", "ws-trust-allow-renewal", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ppsaml_version": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML version", "saml-version", "").AddStringEnum("2", "1.1", "1").AddDefaultValue("2").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("2", "1.1", "1"),
			},
			Default: stringdefault.StaticString("2"),
		},
		"ppsaml_send_slo": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Send SAML Single Logout request (SAML 2.0 only)", "saml-send-slo", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ppsamlslo_endpoint": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML SLO service URL", "saml-slo-endpoint", "").String,
			Optional:            true,
		},
		"ppws_username_token": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Add WS-Security UsernameToken", "wssec-add-user-name-token", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ppws_username_token_password_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WS-Security UsernameToken password type", "wssec-user-name-token-type", "").AddStringEnum("Text", "Digest").AddDefaultValue("Digest").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("Text", "Digest"),
			},
			Default: stringdefault.StaticString("Digest"),
		},
		"ppsaml_validity": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Assertion validity", "saml-validity", "").AddDefaultValue("0").String,
			Computed:            true,
			Optional:            true,
			Default:             int64default.StaticInt64(0),
		},
		"ppsaml_skew": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Skew time", "saml-skew", "").AddDefaultValue("0").String,
			Computed:            true,
			Optional:            true,
			Default:             int64default.StaticInt64(0),
		},
		"ppws_username_token_include_pwd": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Include password in UsernameToken", "wssec-user-name-token-contains-pwd", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"ppltpa": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Generate LTPA token", "lpta-generate-token", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ppltpa_version": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LTPA token version", "lpta-version", "").AddStringEnum("LTPA", "LTPA1FIPS", "LTPA2", "LTPA2WAS7", "LTPADomino").AddDefaultValue("LTPA2").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("LTPA", "LTPA1FIPS", "LTPA2", "LTPA2WAS7", "LTPADomino"),
			},
			Default: stringdefault.StaticString("LTPA2"),
		},
		"ppltpa_expiry": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LTPA token expiry", "lpta-expiry", "").AddIntegerRange(1, 628992000).AddDefaultValue("600").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(1, 628992000),
			},
			Default: int64default.StaticInt64(600),
		},
		"ppltpa_key_file": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LTPA key file", "lpta-key-file", "").String,
			Optional:            true,
		},
		"ppltpa_key_file_password": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LTPA key file password", "lpta-key-file-password", "").String,
			Optional:            true,
		},
		"ppltpa_stash_file": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LTPA stash file", "lpta-stash-file", "").String,
			Optional:            true,
		},
		"pp_kerberos_spnego_token": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Generate Kerberos SPNEGO token", "kerberos-generate-spnego", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"pp_kerberos_bst_value_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("ValueType for generated Kerberos BinarySecurityToken", "kerberos-value-type", "").AddStringEnum("http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ1510", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ1510", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ4120", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ4120", "http://docs.oasis-open.org/wss/2005/xx/oasis-2005xx-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ", "http://docs.oasis-open.org/wss/2005/xx/oasis-2005xx-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ", "http://www.docs.oasis-open.org/wss/2004/07/oasis-000000-wss-kerberos-token-profile-1.0#Kerberosv5_AP_REQ").AddDefaultValue("http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ1510", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ1510", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ4120", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ4120", "http://docs.oasis-open.org/wss/2005/xx/oasis-2005xx-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ", "http://docs.oasis-open.org/wss/2005/xx/oasis-2005xx-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ", "http://www.docs.oasis-open.org/wss/2004/07/oasis-000000-wss-kerberos-token-profile-1.0#Kerberosv5_AP_REQ"),
			},
			Default: stringdefault.StaticString("http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ"),
		},
		"ppsaml_use_ws_sec": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Wrap SAML assertion in WS-Security Security header", "saml-in-wssec", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"pp_kerberos_client_keytab": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Kerberos client keytab", "kerberos-client-keytab", "cryptokerberoskeytab").String,
			Optional:            true,
		},
		"pp_use_ws_sec": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Wrap token in WS-Security Security header", "wssec-header-wrap-token", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"pp_actor_role_id": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Actor or role identifier", "wssec-actor-role-id", "").String,
			Optional:            true,
		},
		"pptfim_token_mapping": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Request Federated Identity Manager token mapping (deprecated)", "tfim-token-mapping", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"pptfim_endpoint": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Federated Identity Manager endpoint", "tfim-endpoint", "tfimendpoint").String,
			Optional:            true,
		},
		"ppws_derived_key_username_token": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Use Derived-Key variant of WS-Security UsernameToken", "wssec-use-derived-key", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ppws_derived_key_username_token_iterations": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Hashing iteration count", "wssec-derived-key-hash-iter", "").AddIntegerRange(2, 65535).AddDefaultValue("1000").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(2, 65535),
			},
			Default: int64default.StaticInt64(1000),
		},
		"ppws_username_token_allow_replacement": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Replace existing UsernameToken", "wssec-replace-existing", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"pptfim_replace_method": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Replacement method (deprecated)", "tfim-replace-method", "").AddStringEnum("all", "replace", "preserve").AddDefaultValue("all").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("all", "replace", "preserve"),
			},
			Default: stringdefault.StaticString("all"),
		},
		"pptfim_retrieve_mode": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Retrieval method (deprecated)", "tfim-retrieval-method", "").AddStringEnum("CallTFIM", "FromMC").AddDefaultValue("CallTFIM").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("CallTFIM", "FromMC"),
			},
			Default: stringdefault.StaticString("CallTFIM"),
		},
		"pphmac_signing_alg": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("HMAC signing algorithm", "hmac-signing-algorithm", "").AddStringEnum("hmac-sha1", "hmac-sha224", "hmac-sha256", "hmac-sha384", "hmac-sha512", "hmac-ripemd160", "hmac-md5").AddDefaultValue("hmac-sha1").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("hmac-sha1", "hmac-sha224", "hmac-sha256", "hmac-sha384", "hmac-sha512", "hmac-ripemd160", "hmac-md5"),
			},
			Default: stringdefault.StaticString("hmac-sha1"),
		},
		"pp_signing_hash_alg": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Signing message digest algorithm", "message-digest-algorithm", "").AddStringEnum("sha1", "sha256", "sha512", "ripemd160", "sha224", "sha384", "md5").AddDefaultValue("sha1").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("sha1", "sha256", "sha512", "ripemd160", "sha224", "sha384", "md5"),
			},
			Default: stringdefault.StaticString("sha1"),
		},
		"ppws_trust_header": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Return the WS-Trust token as SOAP header", "ws-trust-in-header", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ppwssc_key_source": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Source of shared secret to initialize SecurityContext", "ws-trust-key-source", "").AddStringEnum("client-entropy", "in-kerberos", "in-encryptedkey", "static", "random").AddDefaultValue("random").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("client-entropy", "in-kerberos", "in-encryptedkey", "static", "random"),
			},
			Default: stringdefault.StaticString("random"),
		},
		"pp_shared_secret_key": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Shared secret key", "ws-trust-shared-key", "cryptosskey").String,
			Optional:            true,
		},
		"ppws_trust_renewal_wait": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Wait time for renewal", "ws-trust-renewal-wait", "").AddIntegerRange(0, 2678400).AddDefaultValue("0").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(0, 2678400),
			},
			Default: int64default.StaticInt64(0),
		},
		"ppws_trust_new_instance": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Issue new Instance for WS-Trust renewal", "ws-trust-new-instance", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ppws_trust_new_key": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Update context key for WS-Trust renewal", "ws-trust-new-key", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ppws_trust_never_expire": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WS-Trust SecurityContext never expires", "ws-trust-never-expire", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ppicrx_token": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Generate ICRX token for z/OS identity propagation", "generate-icrx", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ppicrx_user_realm": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("ICRX realm", "icrx-user-realm", "").String,
			Optional:            true,
		},
		"ppsaml_identity_provider": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Generate SAML assertion or response", "generate-saml-assertion", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ppsaml_protocol": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML protocol or profile", "saml-protocol", "").AddStringEnum("assertion", "response").AddDefaultValue("assertion").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("assertion", "response"),
			},
			Default: stringdefault.StaticString("assertion"),
		},
		"ppsaml_response_destination": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Response destination", "saml-response-destination", "").String,
			Optional:            true,
		},
		"pp_result_wrapup": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Wrap up result", "result-wrapup", "").AddStringEnum("wssec-replace", "wssec-new", "wssec-inject", "soap-body", "none").AddDefaultValue("wssec-replace").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("wssec-replace", "wssec-new", "wssec-inject", "soap-body", "none"),
			},
			Default: stringdefault.StaticString("wssec-replace"),
		},
		"ppsaml_assertion_type": GetDmSAMLStatementTypeResourceSchema("SAML assertion type", "saml-assertion-type", "", false),
		"ppsaml_subject_confirm": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML subject confirmation method", "saml-subject-confirm", "").AddStringEnum("bearer", "hok", "sv").AddDefaultValue("bearer").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("bearer", "hok", "sv"),
			},
			Default: stringdefault.StaticString("bearer"),
		},
		"ppsaml_name_id": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML subject contains name identifier", "saml-nid", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"ppsaml_name_id_format": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML name identifier format", "saml-nid-format", "").String,
			Optional:            true,
		},
		"ppsaml_recipient": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML recipient", "saml-recipient", "").String,
			Optional:            true,
		},
		"ppsaml_audience": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML audience", "saml-audience", "").String,
			Optional:            true,
		},
		"ppsaml_omit_not_before": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Omit NotBefore attribute", "saml-omit-notbefore", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"pp_one_time_use": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("One time use only", "one-time-use", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ppsaml_proxy": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Allow SAML ProxyRestriction", "saml-proxy", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ppsaml_proxy_audience": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML proxy audience", "saml-proxy-audience", "").String,
			Optional:            true,
		},
		"ppsaml_proxy_count": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML proxy count", "saml-proxy-count", "").AddIntegerRange(0, 65535).AddDefaultValue("0").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(0, 65535),
			},
			Default: int64default.StaticInt64(0),
		},
		"ppsaml_authz_action": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Action for SAML Authorization decision", "saml-authz-action", "").AddStringEnum("AllHTTP", "POST", "GET", "PUT", "HEAD", "General", "Read", "Write", "Execute", "Delete", "Control", "NegatedRead", "NegatedWrite", "NegatedExecute", "NegatedDelete", "NegatedControl").AddDefaultValue("AllHTTP").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("AllHTTP", "POST", "GET", "PUT", "HEAD", "General", "Read", "Write", "Execute", "Delete", "Control", "NegatedRead", "NegatedWrite", "NegatedExecute", "NegatedDelete", "NegatedControl"),
			},
			Default: stringdefault.StaticString("AllHTTP"),
		},
		"ppsaml_attributes": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML attribute definition", "saml-attributes", "samlattributes").String,
			Optional:            true,
		},
		"ppltpa_insert_cookie": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Insert LTPA Set-Cookie", "ltpa-insert-cookie", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"pptampac_propagate": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Generate an Access Manager PAC token", "propagate-tam-pac", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"pptam_header": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Privilege Attribute Certificate header name", "tam-header", "").AddDefaultValue("iv-creds").String,
			Computed:            true,
			Optional:            true,
			Default:             stringdefault.StaticString("iv-creds"),
		},
		"pptam_header_size": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("PAC header value size", "tam-header-size", "").AddDefaultValue("0").String,
			Computed:            true,
			Optional:            true,
			Default:             int64default.StaticInt64(0),
		},
		"pp_kerberos_use_s4u2_proxy": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Use constrained delegation", "kerberos-use-s4u2proxy", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"pp_cookie_attributes": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Cookie attribute policy", "cookie-attributes", "cookieattributepolicy").String,
			Optional:            true,
		},
		"pp_kerberos_use_s4u2_self_and_s4u2_proxy": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Use protocol transition and constrained delegation", "kerberos-use-s4u2self", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"pp_kerberos_client_source": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Kerberos client principal source", "kerberos-client-source", "").AddStringEnum("mc-output", "custom-url", "ctx-var").AddDefaultValue("mc-output").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("mc-output", "custom-url", "ctx-var"),
			},
			Default: stringdefault.StaticString("mc-output"),
		},
		"pp_kerberos_self": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Kerberos self principal", "kerberos-self-principal", "").String,
			Optional:            true,
		},
		"pp_kerberos_self_keytab": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Kerberos self keytab", "kerberos-self-keytab", "cryptokerberoskeytab").String,
			Optional:            true,
		},
		"pp_kerberos_client_custom_url": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Kerberos client principal - custom processing", "kerberos-client-custom-url", "").String,
			Optional:            true,
		},
		"pp_kerberos_client_ctx_var": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Kerberos client principal - context variable", "kerberos-client-ctx-var", "").String,
			Optional:            true,
		},
		"pp_kerberos_server_source": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Kerberos server principal source", "kerberos-server-source", "").AddStringEnum("as-is-string", "custom-url", "ctx-var").AddDefaultValue("as-is-string").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("as-is-string", "custom-url", "ctx-var"),
			},
			Default: stringdefault.StaticString("as-is-string"),
		},
		"pp_kerberos_server_custom_url": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Kerberos server principal - custom processing", "kerberos-server-custom-url", "").String,
			Optional:            true,
		},
		"pp_kerberos_server_ctx_var": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Kerberos server principal - context variable", "kerberos-server-ctx-var", "").String,
			Optional:            true,
		},
		"ppssl_client_config_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("TLS client type", "ssl-client-type", "").AddStringEnum("proxy", "client").AddDefaultValue("proxy").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("proxy", "client"),
			},
			Default: stringdefault.StaticString("proxy"),
		},
		"ppssl_client_profile": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "sslclientprofile").String,
			Optional:            true,
		},
		"ppltpa_key_file_password_alias": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LTPA key file password alias", "ltpa-key-file-password-alias", "passwordalias").String,
			Optional:            true,
		},
		"ppjwt": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Generate a JWT token", "jwt", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ppjwt_generator": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("JWT Generator settings", "generate-jwt", "aaajwtgenerator").String,
			Optional:            true,
		},
	},
}

func (data DmAAAPPostProcess) IsNull() bool {
	if !data.PpEnabled.IsNull() {
		return false
	}
	if !data.PpCustomUrl.IsNull() {
		return false
	}
	if !data.PpsamlAuthAssertion.IsNull() {
		return false
	}
	if !data.PpsamlServerName.IsNull() {
		return false
	}
	if !data.PpsamlNameQualifier.IsNull() {
		return false
	}
	if !data.PpKerberosTicket.IsNull() {
		return false
	}
	if !data.PpKerberosClient.IsNull() {
		return false
	}
	if !data.PpKerberosClientPassword.IsNull() {
		return false
	}
	if !data.PpKerberosServer.IsNull() {
		return false
	}
	if !data.PpwsTrust.IsNull() {
		return false
	}
	if !data.PpTimestamp.IsNull() {
		return false
	}
	if !data.PpTimestampExpiry.IsNull() {
		return false
	}
	if !data.PpAllowRenewal.IsNull() {
		return false
	}
	if !data.PpsamlVersion.IsNull() {
		return false
	}
	if !data.PpsamlSendSlo.IsNull() {
		return false
	}
	if !data.PpsamlsloEndpoint.IsNull() {
		return false
	}
	if !data.PpwsUsernameToken.IsNull() {
		return false
	}
	if !data.PpwsUsernameTokenPasswordType.IsNull() {
		return false
	}
	if !data.PpsamlValidity.IsNull() {
		return false
	}
	if !data.PpsamlSkew.IsNull() {
		return false
	}
	if !data.PpwsUsernameTokenIncludePwd.IsNull() {
		return false
	}
	if !data.Ppltpa.IsNull() {
		return false
	}
	if !data.PpltpaVersion.IsNull() {
		return false
	}
	if !data.PpltpaExpiry.IsNull() {
		return false
	}
	if !data.PpltpaKeyFile.IsNull() {
		return false
	}
	if !data.PpltpaKeyFilePassword.IsNull() {
		return false
	}
	if !data.PpltpaStashFile.IsNull() {
		return false
	}
	if !data.PpKerberosSpnegoToken.IsNull() {
		return false
	}
	if !data.PpKerberosBstValueType.IsNull() {
		return false
	}
	if !data.PpsamlUseWsSec.IsNull() {
		return false
	}
	if !data.PpKerberosClientKeytab.IsNull() {
		return false
	}
	if !data.PpUseWsSec.IsNull() {
		return false
	}
	if !data.PpActorRoleId.IsNull() {
		return false
	}
	if !data.PptfimTokenMapping.IsNull() {
		return false
	}
	if !data.PptfimEndpoint.IsNull() {
		return false
	}
	if !data.PpwsDerivedKeyUsernameToken.IsNull() {
		return false
	}
	if !data.PpwsDerivedKeyUsernameTokenIterations.IsNull() {
		return false
	}
	if !data.PpwsUsernameTokenAllowReplacement.IsNull() {
		return false
	}
	if !data.PptfimReplaceMethod.IsNull() {
		return false
	}
	if !data.PptfimRetrieveMode.IsNull() {
		return false
	}
	if !data.PphmacSigningAlg.IsNull() {
		return false
	}
	if !data.PpSigningHashAlg.IsNull() {
		return false
	}
	if !data.PpwsTrustHeader.IsNull() {
		return false
	}
	if !data.PpwsscKeySource.IsNull() {
		return false
	}
	if !data.PpSharedSecretKey.IsNull() {
		return false
	}
	if !data.PpwsTrustRenewalWait.IsNull() {
		return false
	}
	if !data.PpwsTrustNewInstance.IsNull() {
		return false
	}
	if !data.PpwsTrustNewKey.IsNull() {
		return false
	}
	if !data.PpwsTrustNeverExpire.IsNull() {
		return false
	}
	if !data.PpicrxToken.IsNull() {
		return false
	}
	if !data.PpicrxUserRealm.IsNull() {
		return false
	}
	if !data.PpsamlIdentityProvider.IsNull() {
		return false
	}
	if !data.PpsamlProtocol.IsNull() {
		return false
	}
	if !data.PpsamlResponseDestination.IsNull() {
		return false
	}
	if !data.PpResultWrapup.IsNull() {
		return false
	}
	if data.PpsamlAssertionType != nil {
		if !data.PpsamlAssertionType.IsNull() {
			return false
		}
	}
	if !data.PpsamlSubjectConfirm.IsNull() {
		return false
	}
	if !data.PpsamlNameId.IsNull() {
		return false
	}
	if !data.PpsamlNameIdFormat.IsNull() {
		return false
	}
	if !data.PpsamlRecipient.IsNull() {
		return false
	}
	if !data.PpsamlAudience.IsNull() {
		return false
	}
	if !data.PpsamlOmitNotBefore.IsNull() {
		return false
	}
	if !data.PpOneTimeUse.IsNull() {
		return false
	}
	if !data.PpsamlProxy.IsNull() {
		return false
	}
	if !data.PpsamlProxyAudience.IsNull() {
		return false
	}
	if !data.PpsamlProxyCount.IsNull() {
		return false
	}
	if !data.PpsamlAuthzAction.IsNull() {
		return false
	}
	if !data.PpsamlAttributes.IsNull() {
		return false
	}
	if !data.PpltpaInsertCookie.IsNull() {
		return false
	}
	if !data.PptampacPropagate.IsNull() {
		return false
	}
	if !data.PptamHeader.IsNull() {
		return false
	}
	if !data.PptamHeaderSize.IsNull() {
		return false
	}
	if !data.PpKerberosUseS4u2Proxy.IsNull() {
		return false
	}
	if !data.PpCookieAttributes.IsNull() {
		return false
	}
	if !data.PpKerberosUseS4u2SelfAndS4u2Proxy.IsNull() {
		return false
	}
	if !data.PpKerberosClientSource.IsNull() {
		return false
	}
	if !data.PpKerberosSelf.IsNull() {
		return false
	}
	if !data.PpKerberosSelfKeytab.IsNull() {
		return false
	}
	if !data.PpKerberosClientCustomUrl.IsNull() {
		return false
	}
	if !data.PpKerberosClientCtxVar.IsNull() {
		return false
	}
	if !data.PpKerberosServerSource.IsNull() {
		return false
	}
	if !data.PpKerberosServerCustomUrl.IsNull() {
		return false
	}
	if !data.PpKerberosServerCtxVar.IsNull() {
		return false
	}
	if !data.PpsslClientConfigType.IsNull() {
		return false
	}
	if !data.PpsslClientProfile.IsNull() {
		return false
	}
	if !data.PpltpaKeyFilePasswordAlias.IsNull() {
		return false
	}
	if !data.Ppjwt.IsNull() {
		return false
	}
	if !data.PpjwtGenerator.IsNull() {
		return false
	}
	return true
}
func GetDmAAAPPostProcessDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmAAAPPostProcessDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmAAAPPostProcessDataSourceSchema
}

func GetDmAAAPPostProcessResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmAAAPPostProcessResourceSchema.Required = true
	} else {
		DmAAAPPostProcessResourceSchema.Optional = true
		DmAAAPPostProcessResourceSchema.Computed = true
	}
	DmAAAPPostProcessResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmAAAPPostProcessResourceSchema
}

func (data DmAAAPPostProcess) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.PpEnabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPEnabled`, tfutils.StringFromBool(data.PpEnabled, ""))
	}
	if !data.PpCustomUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPCustomURL`, data.PpCustomUrl.ValueString())
	}
	if !data.PpsamlAuthAssertion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLAuthAssertion`, tfutils.StringFromBool(data.PpsamlAuthAssertion, ""))
	}
	if !data.PpsamlServerName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLServerName`, data.PpsamlServerName.ValueString())
	}
	if !data.PpsamlNameQualifier.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLNameQualifier`, data.PpsamlNameQualifier.ValueString())
	}
	if !data.PpKerberosTicket.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPKerberosTicket`, tfutils.StringFromBool(data.PpKerberosTicket, ""))
	}
	if !data.PpKerberosClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPKerberosClient`, data.PpKerberosClient.ValueString())
	}
	if !data.PpKerberosClientPassword.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPKerberosClientPassword`, data.PpKerberosClientPassword.ValueString())
	}
	if !data.PpKerberosServer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPKerberosServer`, data.PpKerberosServer.ValueString())
	}
	if !data.PpwsTrust.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPWSTrust`, tfutils.StringFromBool(data.PpwsTrust, ""))
	}
	if !data.PpTimestamp.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPTimestamp`, tfutils.StringFromBool(data.PpTimestamp, ""))
	}
	if !data.PpTimestampExpiry.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPTimestampExpiry`, data.PpTimestampExpiry.ValueInt64())
	}
	if !data.PpAllowRenewal.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPAllowRenewal`, tfutils.StringFromBool(data.PpAllowRenewal, ""))
	}
	if !data.PpsamlVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLVersion`, data.PpsamlVersion.ValueString())
	}
	if !data.PpsamlSendSlo.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLSendSLO`, tfutils.StringFromBool(data.PpsamlSendSlo, ""))
	}
	if !data.PpsamlsloEndpoint.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLSLOEndpoint`, data.PpsamlsloEndpoint.ValueString())
	}
	if !data.PpwsUsernameToken.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPWSUsernameToken`, tfutils.StringFromBool(data.PpwsUsernameToken, ""))
	}
	if !data.PpwsUsernameTokenPasswordType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPWSUsernameTokenPasswordType`, data.PpwsUsernameTokenPasswordType.ValueString())
	}
	if !data.PpsamlValidity.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLValidity`, data.PpsamlValidity.ValueInt64())
	}
	if !data.PpsamlSkew.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLSkew`, data.PpsamlSkew.ValueInt64())
	}
	if !data.PpwsUsernameTokenIncludePwd.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPWSUsernameTokenIncludePwd`, tfutils.StringFromBool(data.PpwsUsernameTokenIncludePwd, ""))
	}
	if !data.Ppltpa.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPLTPA`, tfutils.StringFromBool(data.Ppltpa, ""))
	}
	if !data.PpltpaVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPLTPAVersion`, data.PpltpaVersion.ValueString())
	}
	if !data.PpltpaExpiry.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPLTPAExpiry`, data.PpltpaExpiry.ValueInt64())
	}
	if !data.PpltpaKeyFile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPLTPAKeyFile`, data.PpltpaKeyFile.ValueString())
	}
	if !data.PpltpaKeyFilePassword.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPLTPAKeyFilePassword`, data.PpltpaKeyFilePassword.ValueString())
	}
	if !data.PpltpaStashFile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPLTPAStashFile`, data.PpltpaStashFile.ValueString())
	}
	if !data.PpKerberosSpnegoToken.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPKerberosSPNEGOToken`, tfutils.StringFromBool(data.PpKerberosSpnegoToken, ""))
	}
	if !data.PpKerberosBstValueType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPKerberosBstValueType`, data.PpKerberosBstValueType.ValueString())
	}
	if !data.PpsamlUseWsSec.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLUseWSSec`, tfutils.StringFromBool(data.PpsamlUseWsSec, ""))
	}
	if !data.PpKerberosClientKeytab.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPKerberosClientKeytab`, data.PpKerberosClientKeytab.ValueString())
	}
	if !data.PpUseWsSec.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPUseWSSec`, tfutils.StringFromBool(data.PpUseWsSec, ""))
	}
	if !data.PpActorRoleId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPActorRoleID`, data.PpActorRoleId.ValueString())
	}
	if !data.PptfimTokenMapping.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPTFIMTokenMapping`, tfutils.StringFromBool(data.PptfimTokenMapping, ""))
	}
	if !data.PptfimEndpoint.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPTFIMEndpoint`, data.PptfimEndpoint.ValueString())
	}
	if !data.PpwsDerivedKeyUsernameToken.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPWSDerivedKeyUsernameToken`, tfutils.StringFromBool(data.PpwsDerivedKeyUsernameToken, ""))
	}
	if !data.PpwsDerivedKeyUsernameTokenIterations.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPWSDerivedKeyUsernameTokenIterations`, data.PpwsDerivedKeyUsernameTokenIterations.ValueInt64())
	}
	if !data.PpwsUsernameTokenAllowReplacement.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPWSUsernameTokenAllowReplacement`, tfutils.StringFromBool(data.PpwsUsernameTokenAllowReplacement, ""))
	}
	if !data.PptfimReplaceMethod.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPTFIMReplaceMethod`, data.PptfimReplaceMethod.ValueString())
	}
	if !data.PptfimRetrieveMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPTFIMRetrieveMode`, data.PptfimRetrieveMode.ValueString())
	}
	if !data.PphmacSigningAlg.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPHMACSigningAlg`, data.PphmacSigningAlg.ValueString())
	}
	if !data.PpSigningHashAlg.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSigningHashAlg`, data.PpSigningHashAlg.ValueString())
	}
	if !data.PpwsTrustHeader.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPWSTrustHeader`, tfutils.StringFromBool(data.PpwsTrustHeader, ""))
	}
	if !data.PpwsscKeySource.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPWSSCKeySource`, data.PpwsscKeySource.ValueString())
	}
	if !data.PpSharedSecretKey.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSharedSecretKey`, data.PpSharedSecretKey.ValueString())
	}
	if !data.PpwsTrustRenewalWait.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPWSTrustRenewalWait`, data.PpwsTrustRenewalWait.ValueInt64())
	}
	if !data.PpwsTrustNewInstance.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPWSTrustNewInstance`, tfutils.StringFromBool(data.PpwsTrustNewInstance, ""))
	}
	if !data.PpwsTrustNewKey.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPWSTrustNewKey`, tfutils.StringFromBool(data.PpwsTrustNewKey, ""))
	}
	if !data.PpwsTrustNeverExpire.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPWSTrustNeverExpire`, tfutils.StringFromBool(data.PpwsTrustNeverExpire, ""))
	}
	if !data.PpicrxToken.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPICRXToken`, tfutils.StringFromBool(data.PpicrxToken, ""))
	}
	if !data.PpicrxUserRealm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPICRXUserRealm`, data.PpicrxUserRealm.ValueString())
	}
	if !data.PpsamlIdentityProvider.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLIdentityProvider`, tfutils.StringFromBool(data.PpsamlIdentityProvider, ""))
	}
	if !data.PpsamlProtocol.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLProtocol`, data.PpsamlProtocol.ValueString())
	}
	if !data.PpsamlResponseDestination.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLResponseDestination`, data.PpsamlResponseDestination.ValueString())
	}
	if !data.PpResultWrapup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPResultWrapup`, data.PpResultWrapup.ValueString())
	}
	if data.PpsamlAssertionType != nil {
		if !data.PpsamlAssertionType.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`PPSAMLAssertionType`, data.PpsamlAssertionType.ToBody(ctx, ""))
		}
	}
	if !data.PpsamlSubjectConfirm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLSubjectConfirm`, data.PpsamlSubjectConfirm.ValueString())
	}
	if !data.PpsamlNameId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLNameID`, tfutils.StringFromBool(data.PpsamlNameId, ""))
	}
	if !data.PpsamlNameIdFormat.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLNameIDFormat`, data.PpsamlNameIdFormat.ValueString())
	}
	if !data.PpsamlRecipient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLRecipient`, data.PpsamlRecipient.ValueString())
	}
	if !data.PpsamlAudience.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLAudience`, data.PpsamlAudience.ValueString())
	}
	if !data.PpsamlOmitNotBefore.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLOmitNotBefore`, tfutils.StringFromBool(data.PpsamlOmitNotBefore, ""))
	}
	if !data.PpOneTimeUse.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPOneTimeUse`, tfutils.StringFromBool(data.PpOneTimeUse, ""))
	}
	if !data.PpsamlProxy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLProxy`, tfutils.StringFromBool(data.PpsamlProxy, ""))
	}
	if !data.PpsamlProxyAudience.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLProxyAudience`, data.PpsamlProxyAudience.ValueString())
	}
	if !data.PpsamlProxyCount.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLProxyCount`, data.PpsamlProxyCount.ValueInt64())
	}
	if !data.PpsamlAuthzAction.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLAuthzAction`, data.PpsamlAuthzAction.ValueString())
	}
	if !data.PpsamlAttributes.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLAttributes`, data.PpsamlAttributes.ValueString())
	}
	if !data.PpltpaInsertCookie.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPLTPAInsertCookie`, tfutils.StringFromBool(data.PpltpaInsertCookie, ""))
	}
	if !data.PptampacPropagate.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPTAMPACPropagate`, tfutils.StringFromBool(data.PptampacPropagate, ""))
	}
	if !data.PptamHeader.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPTAMHeader`, data.PptamHeader.ValueString())
	}
	if !data.PptamHeaderSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPTAMHeaderSize`, data.PptamHeaderSize.ValueInt64())
	}
	if !data.PpKerberosUseS4u2Proxy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPKerberosUseS4U2Proxy`, tfutils.StringFromBool(data.PpKerberosUseS4u2Proxy, ""))
	}
	if !data.PpCookieAttributes.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPCookieAttributes`, data.PpCookieAttributes.ValueString())
	}
	if !data.PpKerberosUseS4u2SelfAndS4u2Proxy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPKerberosUseS4U2SelfAndS4U2Proxy`, tfutils.StringFromBool(data.PpKerberosUseS4u2SelfAndS4u2Proxy, ""))
	}
	if !data.PpKerberosClientSource.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPKerberosClientSource`, data.PpKerberosClientSource.ValueString())
	}
	if !data.PpKerberosSelf.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPKerberosSelf`, data.PpKerberosSelf.ValueString())
	}
	if !data.PpKerberosSelfKeytab.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPKerberosSelfKeytab`, data.PpKerberosSelfKeytab.ValueString())
	}
	if !data.PpKerberosClientCustomUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPKerberosClientCustomURL`, data.PpKerberosClientCustomUrl.ValueString())
	}
	if !data.PpKerberosClientCtxVar.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPKerberosClientCtxVar`, data.PpKerberosClientCtxVar.ValueString())
	}
	if !data.PpKerberosServerSource.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPKerberosServerSource`, data.PpKerberosServerSource.ValueString())
	}
	if !data.PpKerberosServerCustomUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPKerberosServerCustomURL`, data.PpKerberosServerCustomUrl.ValueString())
	}
	if !data.PpKerberosServerCtxVar.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPKerberosServerCtxVar`, data.PpKerberosServerCtxVar.ValueString())
	}
	if !data.PpsslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSSLClientConfigType`, data.PpsslClientConfigType.ValueString())
	}
	if !data.PpsslClientProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSSLClientProfile`, data.PpsslClientProfile.ValueString())
	}
	if !data.PpltpaKeyFilePasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPLTPAKeyFilePasswordAlias`, data.PpltpaKeyFilePasswordAlias.ValueString())
	}
	if !data.Ppjwt.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPJWT`, tfutils.StringFromBool(data.Ppjwt, ""))
	}
	if !data.PpjwtGenerator.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPJWTGenerator`, data.PpjwtGenerator.ValueString())
	}
	return body
}

func (data *DmAAAPPostProcess) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `PPEnabled`); value.Exists() {
		data.PpEnabled = tfutils.BoolFromString(value.String())
	} else {
		data.PpEnabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPCustomURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpCustomUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpCustomUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLAuthAssertion`); value.Exists() {
		data.PpsamlAuthAssertion = tfutils.BoolFromString(value.String())
	} else {
		data.PpsamlAuthAssertion = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPSAMLServerName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpsamlServerName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpsamlServerName = types.StringValue("XS")
	}
	if value := res.Get(pathRoot + `PPSAMLNameQualifier`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpsamlNameQualifier = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpsamlNameQualifier = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPKerberosTicket`); value.Exists() {
		data.PpKerberosTicket = tfutils.BoolFromString(value.String())
	} else {
		data.PpKerberosTicket = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPKerberosClient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpKerberosClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpKerberosClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPKerberosClientPassword`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpKerberosClientPassword = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpKerberosClientPassword = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPKerberosServer`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpKerberosServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpKerberosServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPWSTrust`); value.Exists() {
		data.PpwsTrust = tfutils.BoolFromString(value.String())
	} else {
		data.PpwsTrust = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPTimestamp`); value.Exists() {
		data.PpTimestamp = tfutils.BoolFromString(value.String())
	} else {
		data.PpTimestamp = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPTimestampExpiry`); value.Exists() {
		data.PpTimestampExpiry = types.Int64Value(value.Int())
	} else {
		data.PpTimestampExpiry = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `PPAllowRenewal`); value.Exists() {
		data.PpAllowRenewal = tfutils.BoolFromString(value.String())
	} else {
		data.PpAllowRenewal = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPSAMLVersion`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpsamlVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpsamlVersion = types.StringValue("2")
	}
	if value := res.Get(pathRoot + `PPSAMLSendSLO`); value.Exists() {
		data.PpsamlSendSlo = tfutils.BoolFromString(value.String())
	} else {
		data.PpsamlSendSlo = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPSAMLSLOEndpoint`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpsamlsloEndpoint = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpsamlsloEndpoint = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPWSUsernameToken`); value.Exists() {
		data.PpwsUsernameToken = tfutils.BoolFromString(value.String())
	} else {
		data.PpwsUsernameToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPWSUsernameTokenPasswordType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpwsUsernameTokenPasswordType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpwsUsernameTokenPasswordType = types.StringValue("Digest")
	}
	if value := res.Get(pathRoot + `PPSAMLValidity`); value.Exists() {
		data.PpsamlValidity = types.Int64Value(value.Int())
	} else {
		data.PpsamlValidity = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `PPSAMLSkew`); value.Exists() {
		data.PpsamlSkew = types.Int64Value(value.Int())
	} else {
		data.PpsamlSkew = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `PPWSUsernameTokenIncludePwd`); value.Exists() {
		data.PpwsUsernameTokenIncludePwd = tfutils.BoolFromString(value.String())
	} else {
		data.PpwsUsernameTokenIncludePwd = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPLTPA`); value.Exists() {
		data.Ppltpa = tfutils.BoolFromString(value.String())
	} else {
		data.Ppltpa = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPLTPAVersion`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpltpaVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpltpaVersion = types.StringValue("LTPA2")
	}
	if value := res.Get(pathRoot + `PPLTPAExpiry`); value.Exists() {
		data.PpltpaExpiry = types.Int64Value(value.Int())
	} else {
		data.PpltpaExpiry = types.Int64Value(600)
	}
	if value := res.Get(pathRoot + `PPLTPAKeyFile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpltpaKeyFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpltpaKeyFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPLTPAKeyFilePassword`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpltpaKeyFilePassword = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpltpaKeyFilePassword = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPLTPAStashFile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpltpaStashFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpltpaStashFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPKerberosSPNEGOToken`); value.Exists() {
		data.PpKerberosSpnegoToken = tfutils.BoolFromString(value.String())
	} else {
		data.PpKerberosSpnegoToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPKerberosBstValueType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpKerberosBstValueType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpKerberosBstValueType = types.StringValue("http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ")
	}
	if value := res.Get(pathRoot + `PPSAMLUseWSSec`); value.Exists() {
		data.PpsamlUseWsSec = tfutils.BoolFromString(value.String())
	} else {
		data.PpsamlUseWsSec = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPKerberosClientKeytab`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpKerberosClientKeytab = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpKerberosClientKeytab = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPUseWSSec`); value.Exists() {
		data.PpUseWsSec = tfutils.BoolFromString(value.String())
	} else {
		data.PpUseWsSec = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPActorRoleID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpActorRoleId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpActorRoleId = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPTFIMTokenMapping`); value.Exists() {
		data.PptfimTokenMapping = tfutils.BoolFromString(value.String())
	} else {
		data.PptfimTokenMapping = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPTFIMEndpoint`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PptfimEndpoint = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PptfimEndpoint = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPWSDerivedKeyUsernameToken`); value.Exists() {
		data.PpwsDerivedKeyUsernameToken = tfutils.BoolFromString(value.String())
	} else {
		data.PpwsDerivedKeyUsernameToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPWSDerivedKeyUsernameTokenIterations`); value.Exists() {
		data.PpwsDerivedKeyUsernameTokenIterations = types.Int64Value(value.Int())
	} else {
		data.PpwsDerivedKeyUsernameTokenIterations = types.Int64Value(1000)
	}
	if value := res.Get(pathRoot + `PPWSUsernameTokenAllowReplacement`); value.Exists() {
		data.PpwsUsernameTokenAllowReplacement = tfutils.BoolFromString(value.String())
	} else {
		data.PpwsUsernameTokenAllowReplacement = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPTFIMReplaceMethod`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PptfimReplaceMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PptfimReplaceMethod = types.StringValue("all")
	}
	if value := res.Get(pathRoot + `PPTFIMRetrieveMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PptfimRetrieveMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PptfimRetrieveMode = types.StringValue("CallTFIM")
	}
	if value := res.Get(pathRoot + `PPHMACSigningAlg`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PphmacSigningAlg = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PphmacSigningAlg = types.StringValue("hmac-sha1")
	}
	if value := res.Get(pathRoot + `PPSigningHashAlg`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpSigningHashAlg = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpSigningHashAlg = types.StringValue("sha1")
	}
	if value := res.Get(pathRoot + `PPWSTrustHeader`); value.Exists() {
		data.PpwsTrustHeader = tfutils.BoolFromString(value.String())
	} else {
		data.PpwsTrustHeader = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPWSSCKeySource`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpwsscKeySource = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpwsscKeySource = types.StringValue("random")
	}
	if value := res.Get(pathRoot + `PPSharedSecretKey`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpSharedSecretKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpSharedSecretKey = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPWSTrustRenewalWait`); value.Exists() {
		data.PpwsTrustRenewalWait = types.Int64Value(value.Int())
	} else {
		data.PpwsTrustRenewalWait = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `PPWSTrustNewInstance`); value.Exists() {
		data.PpwsTrustNewInstance = tfutils.BoolFromString(value.String())
	} else {
		data.PpwsTrustNewInstance = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPWSTrustNewKey`); value.Exists() {
		data.PpwsTrustNewKey = tfutils.BoolFromString(value.String())
	} else {
		data.PpwsTrustNewKey = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPWSTrustNeverExpire`); value.Exists() {
		data.PpwsTrustNeverExpire = tfutils.BoolFromString(value.String())
	} else {
		data.PpwsTrustNeverExpire = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPICRXToken`); value.Exists() {
		data.PpicrxToken = tfutils.BoolFromString(value.String())
	} else {
		data.PpicrxToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPICRXUserRealm`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpicrxUserRealm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpicrxUserRealm = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLIdentityProvider`); value.Exists() {
		data.PpsamlIdentityProvider = tfutils.BoolFromString(value.String())
	} else {
		data.PpsamlIdentityProvider = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPSAMLProtocol`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpsamlProtocol = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpsamlProtocol = types.StringValue("assertion")
	}
	if value := res.Get(pathRoot + `PPSAMLResponseDestination`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpsamlResponseDestination = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpsamlResponseDestination = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPResultWrapup`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpResultWrapup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpResultWrapup = types.StringValue("wssec-replace")
	}
	if value := res.Get(pathRoot + `PPSAMLAssertionType`); value.Exists() {
		data.PpsamlAssertionType = &DmSAMLStatementType{}
		data.PpsamlAssertionType.FromBody(ctx, "", value)
	} else {
		data.PpsamlAssertionType = nil
	}
	if value := res.Get(pathRoot + `PPSAMLSubjectConfirm`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpsamlSubjectConfirm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpsamlSubjectConfirm = types.StringValue("bearer")
	}
	if value := res.Get(pathRoot + `PPSAMLNameID`); value.Exists() {
		data.PpsamlNameId = tfutils.BoolFromString(value.String())
	} else {
		data.PpsamlNameId = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPSAMLNameIDFormat`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpsamlNameIdFormat = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpsamlNameIdFormat = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLRecipient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpsamlRecipient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpsamlRecipient = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLAudience`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpsamlAudience = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpsamlAudience = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLOmitNotBefore`); value.Exists() {
		data.PpsamlOmitNotBefore = tfutils.BoolFromString(value.String())
	} else {
		data.PpsamlOmitNotBefore = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPOneTimeUse`); value.Exists() {
		data.PpOneTimeUse = tfutils.BoolFromString(value.String())
	} else {
		data.PpOneTimeUse = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPSAMLProxy`); value.Exists() {
		data.PpsamlProxy = tfutils.BoolFromString(value.String())
	} else {
		data.PpsamlProxy = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPSAMLProxyAudience`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpsamlProxyAudience = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpsamlProxyAudience = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLProxyCount`); value.Exists() {
		data.PpsamlProxyCount = types.Int64Value(value.Int())
	} else {
		data.PpsamlProxyCount = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `PPSAMLAuthzAction`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpsamlAuthzAction = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpsamlAuthzAction = types.StringValue("AllHTTP")
	}
	if value := res.Get(pathRoot + `PPSAMLAttributes`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpsamlAttributes = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpsamlAttributes = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPLTPAInsertCookie`); value.Exists() {
		data.PpltpaInsertCookie = tfutils.BoolFromString(value.String())
	} else {
		data.PpltpaInsertCookie = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPTAMPACPropagate`); value.Exists() {
		data.PptampacPropagate = tfutils.BoolFromString(value.String())
	} else {
		data.PptampacPropagate = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPTAMHeader`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PptamHeader = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PptamHeader = types.StringValue("iv-creds")
	}
	if value := res.Get(pathRoot + `PPTAMHeaderSize`); value.Exists() {
		data.PptamHeaderSize = types.Int64Value(value.Int())
	} else {
		data.PptamHeaderSize = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `PPKerberosUseS4U2Proxy`); value.Exists() {
		data.PpKerberosUseS4u2Proxy = tfutils.BoolFromString(value.String())
	} else {
		data.PpKerberosUseS4u2Proxy = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPCookieAttributes`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpCookieAttributes = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpCookieAttributes = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPKerberosUseS4U2SelfAndS4U2Proxy`); value.Exists() {
		data.PpKerberosUseS4u2SelfAndS4u2Proxy = tfutils.BoolFromString(value.String())
	} else {
		data.PpKerberosUseS4u2SelfAndS4u2Proxy = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPKerberosClientSource`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpKerberosClientSource = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpKerberosClientSource = types.StringValue("mc-output")
	}
	if value := res.Get(pathRoot + `PPKerberosSelf`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpKerberosSelf = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpKerberosSelf = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPKerberosSelfKeytab`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpKerberosSelfKeytab = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpKerberosSelfKeytab = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPKerberosClientCustomURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpKerberosClientCustomUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpKerberosClientCustomUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPKerberosClientCtxVar`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpKerberosClientCtxVar = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpKerberosClientCtxVar = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPKerberosServerSource`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpKerberosServerSource = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpKerberosServerSource = types.StringValue("as-is-string")
	}
	if value := res.Get(pathRoot + `PPKerberosServerCustomURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpKerberosServerCustomUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpKerberosServerCustomUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPKerberosServerCtxVar`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpKerberosServerCtxVar = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpKerberosServerCtxVar = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSSLClientConfigType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpsslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpsslClientConfigType = types.StringValue("proxy")
	}
	if value := res.Get(pathRoot + `PPSSLClientProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpsslClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpsslClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPLTPAKeyFilePasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpltpaKeyFilePasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpltpaKeyFilePasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPJWT`); value.Exists() {
		data.Ppjwt = tfutils.BoolFromString(value.String())
	} else {
		data.Ppjwt = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPJWTGenerator`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpjwtGenerator = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpjwtGenerator = types.StringNull()
	}
}

func (data *DmAAAPPostProcess) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `PPEnabled`); value.Exists() && !data.PpEnabled.IsNull() {
		data.PpEnabled = tfutils.BoolFromString(value.String())
	} else if data.PpEnabled.ValueBool() {
		data.PpEnabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPCustomURL`); value.Exists() && !data.PpCustomUrl.IsNull() {
		data.PpCustomUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpCustomUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLAuthAssertion`); value.Exists() && !data.PpsamlAuthAssertion.IsNull() {
		data.PpsamlAuthAssertion = tfutils.BoolFromString(value.String())
	} else if data.PpsamlAuthAssertion.ValueBool() {
		data.PpsamlAuthAssertion = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPSAMLServerName`); value.Exists() && !data.PpsamlServerName.IsNull() {
		data.PpsamlServerName = tfutils.ParseStringFromGJSON(value)
	} else if data.PpsamlServerName.ValueString() != "XS" {
		data.PpsamlServerName = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLNameQualifier`); value.Exists() && !data.PpsamlNameQualifier.IsNull() {
		data.PpsamlNameQualifier = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpsamlNameQualifier = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPKerberosTicket`); value.Exists() && !data.PpKerberosTicket.IsNull() {
		data.PpKerberosTicket = tfutils.BoolFromString(value.String())
	} else if data.PpKerberosTicket.ValueBool() {
		data.PpKerberosTicket = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPKerberosClient`); value.Exists() && !data.PpKerberosClient.IsNull() {
		data.PpKerberosClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpKerberosClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPKerberosClientPassword`); value.Exists() && !data.PpKerberosClientPassword.IsNull() {
		data.PpKerberosClientPassword = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpKerberosClientPassword = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPKerberosServer`); value.Exists() && !data.PpKerberosServer.IsNull() {
		data.PpKerberosServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpKerberosServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPWSTrust`); value.Exists() && !data.PpwsTrust.IsNull() {
		data.PpwsTrust = tfutils.BoolFromString(value.String())
	} else if data.PpwsTrust.ValueBool() {
		data.PpwsTrust = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPTimestamp`); value.Exists() && !data.PpTimestamp.IsNull() {
		data.PpTimestamp = tfutils.BoolFromString(value.String())
	} else if !data.PpTimestamp.ValueBool() {
		data.PpTimestamp = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPTimestampExpiry`); value.Exists() && !data.PpTimestampExpiry.IsNull() {
		data.PpTimestampExpiry = types.Int64Value(value.Int())
	} else if data.PpTimestampExpiry.ValueInt64() != 0 {
		data.PpTimestampExpiry = types.Int64Null()
	}
	if value := res.Get(pathRoot + `PPAllowRenewal`); value.Exists() && !data.PpAllowRenewal.IsNull() {
		data.PpAllowRenewal = tfutils.BoolFromString(value.String())
	} else if data.PpAllowRenewal.ValueBool() {
		data.PpAllowRenewal = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPSAMLVersion`); value.Exists() && !data.PpsamlVersion.IsNull() {
		data.PpsamlVersion = tfutils.ParseStringFromGJSON(value)
	} else if data.PpsamlVersion.ValueString() != "2" {
		data.PpsamlVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLSendSLO`); value.Exists() && !data.PpsamlSendSlo.IsNull() {
		data.PpsamlSendSlo = tfutils.BoolFromString(value.String())
	} else if data.PpsamlSendSlo.ValueBool() {
		data.PpsamlSendSlo = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPSAMLSLOEndpoint`); value.Exists() && !data.PpsamlsloEndpoint.IsNull() {
		data.PpsamlsloEndpoint = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpsamlsloEndpoint = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPWSUsernameToken`); value.Exists() && !data.PpwsUsernameToken.IsNull() {
		data.PpwsUsernameToken = tfutils.BoolFromString(value.String())
	} else if data.PpwsUsernameToken.ValueBool() {
		data.PpwsUsernameToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPWSUsernameTokenPasswordType`); value.Exists() && !data.PpwsUsernameTokenPasswordType.IsNull() {
		data.PpwsUsernameTokenPasswordType = tfutils.ParseStringFromGJSON(value)
	} else if data.PpwsUsernameTokenPasswordType.ValueString() != "Digest" {
		data.PpwsUsernameTokenPasswordType = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLValidity`); value.Exists() && !data.PpsamlValidity.IsNull() {
		data.PpsamlValidity = types.Int64Value(value.Int())
	} else if data.PpsamlValidity.ValueInt64() != 0 {
		data.PpsamlValidity = types.Int64Null()
	}
	if value := res.Get(pathRoot + `PPSAMLSkew`); value.Exists() && !data.PpsamlSkew.IsNull() {
		data.PpsamlSkew = types.Int64Value(value.Int())
	} else if data.PpsamlSkew.ValueInt64() != 0 {
		data.PpsamlSkew = types.Int64Null()
	}
	if value := res.Get(pathRoot + `PPWSUsernameTokenIncludePwd`); value.Exists() && !data.PpwsUsernameTokenIncludePwd.IsNull() {
		data.PpwsUsernameTokenIncludePwd = tfutils.BoolFromString(value.String())
	} else if !data.PpwsUsernameTokenIncludePwd.ValueBool() {
		data.PpwsUsernameTokenIncludePwd = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPLTPA`); value.Exists() && !data.Ppltpa.IsNull() {
		data.Ppltpa = tfutils.BoolFromString(value.String())
	} else if data.Ppltpa.ValueBool() {
		data.Ppltpa = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPLTPAVersion`); value.Exists() && !data.PpltpaVersion.IsNull() {
		data.PpltpaVersion = tfutils.ParseStringFromGJSON(value)
	} else if data.PpltpaVersion.ValueString() != "LTPA2" {
		data.PpltpaVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPLTPAExpiry`); value.Exists() && !data.PpltpaExpiry.IsNull() {
		data.PpltpaExpiry = types.Int64Value(value.Int())
	} else if data.PpltpaExpiry.ValueInt64() != 600 {
		data.PpltpaExpiry = types.Int64Null()
	}
	if value := res.Get(pathRoot + `PPLTPAKeyFile`); value.Exists() && !data.PpltpaKeyFile.IsNull() {
		data.PpltpaKeyFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpltpaKeyFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPLTPAKeyFilePassword`); value.Exists() && !data.PpltpaKeyFilePassword.IsNull() {
		data.PpltpaKeyFilePassword = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpltpaKeyFilePassword = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPLTPAStashFile`); value.Exists() && !data.PpltpaStashFile.IsNull() {
		data.PpltpaStashFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpltpaStashFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPKerberosSPNEGOToken`); value.Exists() && !data.PpKerberosSpnegoToken.IsNull() {
		data.PpKerberosSpnegoToken = tfutils.BoolFromString(value.String())
	} else if data.PpKerberosSpnegoToken.ValueBool() {
		data.PpKerberosSpnegoToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPKerberosBstValueType`); value.Exists() && !data.PpKerberosBstValueType.IsNull() {
		data.PpKerberosBstValueType = tfutils.ParseStringFromGJSON(value)
	} else if data.PpKerberosBstValueType.ValueString() != "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ" {
		data.PpKerberosBstValueType = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLUseWSSec`); value.Exists() && !data.PpsamlUseWsSec.IsNull() {
		data.PpsamlUseWsSec = tfutils.BoolFromString(value.String())
	} else if data.PpsamlUseWsSec.ValueBool() {
		data.PpsamlUseWsSec = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPKerberosClientKeytab`); value.Exists() && !data.PpKerberosClientKeytab.IsNull() {
		data.PpKerberosClientKeytab = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpKerberosClientKeytab = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPUseWSSec`); value.Exists() && !data.PpUseWsSec.IsNull() {
		data.PpUseWsSec = tfutils.BoolFromString(value.String())
	} else if data.PpUseWsSec.ValueBool() {
		data.PpUseWsSec = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPActorRoleID`); value.Exists() && !data.PpActorRoleId.IsNull() {
		data.PpActorRoleId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpActorRoleId = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPTFIMTokenMapping`); value.Exists() && !data.PptfimTokenMapping.IsNull() {
		data.PptfimTokenMapping = tfutils.BoolFromString(value.String())
	} else if data.PptfimTokenMapping.ValueBool() {
		data.PptfimTokenMapping = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPTFIMEndpoint`); value.Exists() && !data.PptfimEndpoint.IsNull() {
		data.PptfimEndpoint = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PptfimEndpoint = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPWSDerivedKeyUsernameToken`); value.Exists() && !data.PpwsDerivedKeyUsernameToken.IsNull() {
		data.PpwsDerivedKeyUsernameToken = tfutils.BoolFromString(value.String())
	} else if data.PpwsDerivedKeyUsernameToken.ValueBool() {
		data.PpwsDerivedKeyUsernameToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPWSDerivedKeyUsernameTokenIterations`); value.Exists() && !data.PpwsDerivedKeyUsernameTokenIterations.IsNull() {
		data.PpwsDerivedKeyUsernameTokenIterations = types.Int64Value(value.Int())
	} else if data.PpwsDerivedKeyUsernameTokenIterations.ValueInt64() != 1000 {
		data.PpwsDerivedKeyUsernameTokenIterations = types.Int64Null()
	}
	if value := res.Get(pathRoot + `PPWSUsernameTokenAllowReplacement`); value.Exists() && !data.PpwsUsernameTokenAllowReplacement.IsNull() {
		data.PpwsUsernameTokenAllowReplacement = tfutils.BoolFromString(value.String())
	} else if data.PpwsUsernameTokenAllowReplacement.ValueBool() {
		data.PpwsUsernameTokenAllowReplacement = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPTFIMReplaceMethod`); value.Exists() && !data.PptfimReplaceMethod.IsNull() {
		data.PptfimReplaceMethod = tfutils.ParseStringFromGJSON(value)
	} else if data.PptfimReplaceMethod.ValueString() != "all" {
		data.PptfimReplaceMethod = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPTFIMRetrieveMode`); value.Exists() && !data.PptfimRetrieveMode.IsNull() {
		data.PptfimRetrieveMode = tfutils.ParseStringFromGJSON(value)
	} else if data.PptfimRetrieveMode.ValueString() != "CallTFIM" {
		data.PptfimRetrieveMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPHMACSigningAlg`); value.Exists() && !data.PphmacSigningAlg.IsNull() {
		data.PphmacSigningAlg = tfutils.ParseStringFromGJSON(value)
	} else if data.PphmacSigningAlg.ValueString() != "hmac-sha1" {
		data.PphmacSigningAlg = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSigningHashAlg`); value.Exists() && !data.PpSigningHashAlg.IsNull() {
		data.PpSigningHashAlg = tfutils.ParseStringFromGJSON(value)
	} else if data.PpSigningHashAlg.ValueString() != "sha1" {
		data.PpSigningHashAlg = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPWSTrustHeader`); value.Exists() && !data.PpwsTrustHeader.IsNull() {
		data.PpwsTrustHeader = tfutils.BoolFromString(value.String())
	} else if data.PpwsTrustHeader.ValueBool() {
		data.PpwsTrustHeader = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPWSSCKeySource`); value.Exists() && !data.PpwsscKeySource.IsNull() {
		data.PpwsscKeySource = tfutils.ParseStringFromGJSON(value)
	} else if data.PpwsscKeySource.ValueString() != "random" {
		data.PpwsscKeySource = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSharedSecretKey`); value.Exists() && !data.PpSharedSecretKey.IsNull() {
		data.PpSharedSecretKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpSharedSecretKey = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPWSTrustRenewalWait`); value.Exists() && !data.PpwsTrustRenewalWait.IsNull() {
		data.PpwsTrustRenewalWait = types.Int64Value(value.Int())
	} else if data.PpwsTrustRenewalWait.ValueInt64() != 0 {
		data.PpwsTrustRenewalWait = types.Int64Null()
	}
	if value := res.Get(pathRoot + `PPWSTrustNewInstance`); value.Exists() && !data.PpwsTrustNewInstance.IsNull() {
		data.PpwsTrustNewInstance = tfutils.BoolFromString(value.String())
	} else if data.PpwsTrustNewInstance.ValueBool() {
		data.PpwsTrustNewInstance = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPWSTrustNewKey`); value.Exists() && !data.PpwsTrustNewKey.IsNull() {
		data.PpwsTrustNewKey = tfutils.BoolFromString(value.String())
	} else if data.PpwsTrustNewKey.ValueBool() {
		data.PpwsTrustNewKey = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPWSTrustNeverExpire`); value.Exists() && !data.PpwsTrustNeverExpire.IsNull() {
		data.PpwsTrustNeverExpire = tfutils.BoolFromString(value.String())
	} else if data.PpwsTrustNeverExpire.ValueBool() {
		data.PpwsTrustNeverExpire = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPICRXToken`); value.Exists() && !data.PpicrxToken.IsNull() {
		data.PpicrxToken = tfutils.BoolFromString(value.String())
	} else if data.PpicrxToken.ValueBool() {
		data.PpicrxToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPICRXUserRealm`); value.Exists() && !data.PpicrxUserRealm.IsNull() {
		data.PpicrxUserRealm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpicrxUserRealm = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLIdentityProvider`); value.Exists() && !data.PpsamlIdentityProvider.IsNull() {
		data.PpsamlIdentityProvider = tfutils.BoolFromString(value.String())
	} else if data.PpsamlIdentityProvider.ValueBool() {
		data.PpsamlIdentityProvider = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPSAMLProtocol`); value.Exists() && !data.PpsamlProtocol.IsNull() {
		data.PpsamlProtocol = tfutils.ParseStringFromGJSON(value)
	} else if data.PpsamlProtocol.ValueString() != "assertion" {
		data.PpsamlProtocol = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLResponseDestination`); value.Exists() && !data.PpsamlResponseDestination.IsNull() {
		data.PpsamlResponseDestination = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpsamlResponseDestination = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPResultWrapup`); value.Exists() && !data.PpResultWrapup.IsNull() {
		data.PpResultWrapup = tfutils.ParseStringFromGJSON(value)
	} else if data.PpResultWrapup.ValueString() != "wssec-replace" {
		data.PpResultWrapup = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLAssertionType`); value.Exists() {
		data.PpsamlAssertionType.UpdateFromBody(ctx, "", value)
	} else {
		data.PpsamlAssertionType = nil
	}
	if value := res.Get(pathRoot + `PPSAMLSubjectConfirm`); value.Exists() && !data.PpsamlSubjectConfirm.IsNull() {
		data.PpsamlSubjectConfirm = tfutils.ParseStringFromGJSON(value)
	} else if data.PpsamlSubjectConfirm.ValueString() != "bearer" {
		data.PpsamlSubjectConfirm = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLNameID`); value.Exists() && !data.PpsamlNameId.IsNull() {
		data.PpsamlNameId = tfutils.BoolFromString(value.String())
	} else if !data.PpsamlNameId.ValueBool() {
		data.PpsamlNameId = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPSAMLNameIDFormat`); value.Exists() && !data.PpsamlNameIdFormat.IsNull() {
		data.PpsamlNameIdFormat = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpsamlNameIdFormat = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLRecipient`); value.Exists() && !data.PpsamlRecipient.IsNull() {
		data.PpsamlRecipient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpsamlRecipient = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLAudience`); value.Exists() && !data.PpsamlAudience.IsNull() {
		data.PpsamlAudience = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpsamlAudience = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLOmitNotBefore`); value.Exists() && !data.PpsamlOmitNotBefore.IsNull() {
		data.PpsamlOmitNotBefore = tfutils.BoolFromString(value.String())
	} else if data.PpsamlOmitNotBefore.ValueBool() {
		data.PpsamlOmitNotBefore = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPOneTimeUse`); value.Exists() && !data.PpOneTimeUse.IsNull() {
		data.PpOneTimeUse = tfutils.BoolFromString(value.String())
	} else if data.PpOneTimeUse.ValueBool() {
		data.PpOneTimeUse = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPSAMLProxy`); value.Exists() && !data.PpsamlProxy.IsNull() {
		data.PpsamlProxy = tfutils.BoolFromString(value.String())
	} else if data.PpsamlProxy.ValueBool() {
		data.PpsamlProxy = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPSAMLProxyAudience`); value.Exists() && !data.PpsamlProxyAudience.IsNull() {
		data.PpsamlProxyAudience = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpsamlProxyAudience = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLProxyCount`); value.Exists() && !data.PpsamlProxyCount.IsNull() {
		data.PpsamlProxyCount = types.Int64Value(value.Int())
	} else if data.PpsamlProxyCount.ValueInt64() != 0 {
		data.PpsamlProxyCount = types.Int64Null()
	}
	if value := res.Get(pathRoot + `PPSAMLAuthzAction`); value.Exists() && !data.PpsamlAuthzAction.IsNull() {
		data.PpsamlAuthzAction = tfutils.ParseStringFromGJSON(value)
	} else if data.PpsamlAuthzAction.ValueString() != "AllHTTP" {
		data.PpsamlAuthzAction = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLAttributes`); value.Exists() && !data.PpsamlAttributes.IsNull() {
		data.PpsamlAttributes = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpsamlAttributes = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPLTPAInsertCookie`); value.Exists() && !data.PpltpaInsertCookie.IsNull() {
		data.PpltpaInsertCookie = tfutils.BoolFromString(value.String())
	} else if !data.PpltpaInsertCookie.ValueBool() {
		data.PpltpaInsertCookie = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPTAMPACPropagate`); value.Exists() && !data.PptampacPropagate.IsNull() {
		data.PptampacPropagate = tfutils.BoolFromString(value.String())
	} else if data.PptampacPropagate.ValueBool() {
		data.PptampacPropagate = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPTAMHeader`); value.Exists() && !data.PptamHeader.IsNull() {
		data.PptamHeader = tfutils.ParseStringFromGJSON(value)
	} else if data.PptamHeader.ValueString() != "iv-creds" {
		data.PptamHeader = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPTAMHeaderSize`); value.Exists() && !data.PptamHeaderSize.IsNull() {
		data.PptamHeaderSize = types.Int64Value(value.Int())
	} else if data.PptamHeaderSize.ValueInt64() != 0 {
		data.PptamHeaderSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `PPKerberosUseS4U2Proxy`); value.Exists() && !data.PpKerberosUseS4u2Proxy.IsNull() {
		data.PpKerberosUseS4u2Proxy = tfutils.BoolFromString(value.String())
	} else if data.PpKerberosUseS4u2Proxy.ValueBool() {
		data.PpKerberosUseS4u2Proxy = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPCookieAttributes`); value.Exists() && !data.PpCookieAttributes.IsNull() {
		data.PpCookieAttributes = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpCookieAttributes = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPKerberosUseS4U2SelfAndS4U2Proxy`); value.Exists() && !data.PpKerberosUseS4u2SelfAndS4u2Proxy.IsNull() {
		data.PpKerberosUseS4u2SelfAndS4u2Proxy = tfutils.BoolFromString(value.String())
	} else if data.PpKerberosUseS4u2SelfAndS4u2Proxy.ValueBool() {
		data.PpKerberosUseS4u2SelfAndS4u2Proxy = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPKerberosClientSource`); value.Exists() && !data.PpKerberosClientSource.IsNull() {
		data.PpKerberosClientSource = tfutils.ParseStringFromGJSON(value)
	} else if data.PpKerberosClientSource.ValueString() != "mc-output" {
		data.PpKerberosClientSource = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPKerberosSelf`); value.Exists() && !data.PpKerberosSelf.IsNull() {
		data.PpKerberosSelf = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpKerberosSelf = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPKerberosSelfKeytab`); value.Exists() && !data.PpKerberosSelfKeytab.IsNull() {
		data.PpKerberosSelfKeytab = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpKerberosSelfKeytab = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPKerberosClientCustomURL`); value.Exists() && !data.PpKerberosClientCustomUrl.IsNull() {
		data.PpKerberosClientCustomUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpKerberosClientCustomUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPKerberosClientCtxVar`); value.Exists() && !data.PpKerberosClientCtxVar.IsNull() {
		data.PpKerberosClientCtxVar = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpKerberosClientCtxVar = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPKerberosServerSource`); value.Exists() && !data.PpKerberosServerSource.IsNull() {
		data.PpKerberosServerSource = tfutils.ParseStringFromGJSON(value)
	} else if data.PpKerberosServerSource.ValueString() != "as-is-string" {
		data.PpKerberosServerSource = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPKerberosServerCustomURL`); value.Exists() && !data.PpKerberosServerCustomUrl.IsNull() {
		data.PpKerberosServerCustomUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpKerberosServerCustomUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPKerberosServerCtxVar`); value.Exists() && !data.PpKerberosServerCtxVar.IsNull() {
		data.PpKerberosServerCtxVar = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpKerberosServerCtxVar = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSSLClientConfigType`); value.Exists() && !data.PpsslClientConfigType.IsNull() {
		data.PpsslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else if data.PpsslClientConfigType.ValueString() != "proxy" {
		data.PpsslClientConfigType = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSSLClientProfile`); value.Exists() && !data.PpsslClientProfile.IsNull() {
		data.PpsslClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpsslClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPLTPAKeyFilePasswordAlias`); value.Exists() && !data.PpltpaKeyFilePasswordAlias.IsNull() {
		data.PpltpaKeyFilePasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpltpaKeyFilePasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPJWT`); value.Exists() && !data.Ppjwt.IsNull() {
		data.Ppjwt = tfutils.BoolFromString(value.String())
	} else if data.Ppjwt.ValueBool() {
		data.Ppjwt = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPJWTGenerator`); value.Exists() && !data.PpjwtGenerator.IsNull() {
		data.PpjwtGenerator = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpjwtGenerator = types.StringNull()
	}
}
