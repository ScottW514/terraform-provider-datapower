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
	PpwsDerivedKeyUsernameToken           types.Bool           `tfsdk:"ppws_derived_key_username_token"`
	PpwsDerivedKeyUsernameTokenIterations types.Int64          `tfsdk:"ppws_derived_key_username_token_iterations"`
	PpwsUsernameTokenAllowReplacement     types.Bool           `tfsdk:"ppws_username_token_allow_replacement"`
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
	"ppws_derived_key_username_token":            types.BoolType,
	"ppws_derived_key_username_token_iterations": types.Int64Type,
	"ppws_username_token_allow_replacement":      types.BoolType,
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
	"ppws_derived_key_username_token":            types.BoolValue(false),
	"ppws_derived_key_username_token_iterations": types.Int64Value(1000),
	"ppws_username_token_allow_replacement":      types.BoolValue(false),
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
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to run a custom stylesheet or GatewayScript file.", "custom-processing", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"pp_custom_url": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the custom file for the postprocessing activity.", "custom-url", "").String,
			Computed:            true,
		},
		"ppsaml_auth_assertion": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate a SAML assertion that contains a SAML authentication statement for the authenticated user identity.", "saml-generate-assertion", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ppsaml_server_name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the value of the <tt>saml:Issuer</tt> of the generated SAML assertion or SAML SLO request. The default value is XS.</p><ul><li>If generating an SAML assertion, identifies the server that makes the assertion.</li><li>If sending an SLO request, identifies the issuer that sends the request.</li></ul>", "saml-server-name", "").AddDefaultValue("XS").String,
			Computed:            true,
		},
		"ppsaml_name_qualifier": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the value of the NameQualifier attribute of the NameIdentifier in the generated SAML assertion. Although the attribute is an optional attribute, some SAML implementations require that this attribute must be present.", "saml-name-qualifier", "").String,
			Computed:            true,
		},
		"pp_kerberos_ticket": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to include a WS-Security Kerberos AP-REQ BinarySecurityToken for the specified client and server principals in the WS-Security header. By default, token are not included.", "kerberos-include-token", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"pp_kerberos_client": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the client identity (cname of the Kerberos ticket) for the Kerberos client principal.", "kerberos-client-principal", "").String,
			Computed:            true,
		},
		"pp_kerberos_client_password": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "").String,
			Computed:            true,
		},
		"pp_kerberos_server": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the server identity (sname of the Kerberos ticket) for the Kerberos server principal.", "kerberos-server", "").String,
			Computed:            true,
		},
		"ppws_trust": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate the appropriate security token response for a valid WS-Trust SecurityContextToken request.", "ws-trust-generate-resp", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"pp_timestamp": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate a WS-Trust token time stamp for the security token response.", "ws-trust-add-timestamp", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"pp_timestamp_expiry": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the validity duration for the WS-Trust SCT in seconds to issue a new security context or to renew a context instance with new instance. Enter a value in the range 0 - 31622400. The default value is 0, which uses the value of the <tt>var://system/AAA/defaultexpiry</tt> variable if defined. If you did not define this variable, the value is 14400. If this setting is to renew a security context or instance, the value 0 means to use the old duration for the renewed cycle.", "ws-trust-timestamp-expiry", "").AddIntegerRange(0, 31622400).AddDefaultValue("0").String,
			Computed:            true,
		},
		"pp_allow_renewal": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether WS-Trust tokens can have their lifetime period reset without a new bootstrapping authentication event. If the WS-Trust request asks to renew the issued token, this setting is ignored.", "ws-trust-allow-renewal", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ppsaml_version": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the protocol level of SAML messages. The version affects the identity extraction from the original message and the format of messages. The default value is 1.1.", "saml-version", "").AddStringEnum("2", "1.1", "1").AddDefaultValue("2").String,
			Computed:            true,
		},
		"ppsaml_send_slo": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to send a SAML Logout (SLO) request to revoke the SAML Assertion token that is used for single-sign-on (SSO). The SLO is a request-response that the DataPower&#174; Gateway handles differently when it is working as a service provider (SP) or identity provider (IdP).</p><ul><li>When an SP, the DataPower Gateway sends an SLO request to the SAML SLO endpoint (IdP). On response, the DataPower Gateway processes the SLO response for its status.</li><li>When an IdP, the request to the DataPower Gateway contains the SLO request. Postprocessing validates against the SAML metadata file and sends the corresponding endpoint the SLO response.</li></ul>", "saml-send-slo", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ppsamlslo_endpoint": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the endpoint URL for SAML 2.0 Single Logout (SLO) messages. This endpoint is the authority that authenticated the assertion subject.", "saml-slo-endpoint", "").String,
			Computed:            true,
		},
		"ppws_username_token": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to add a WS-Security UsernameToken. The username and password are taken from the output of the map credentials phase.", "wssec-add-user-name-token", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ppws_username_token_password_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the type of password that the UsernameToken provides. By default, use the digest of the password as defined in the \"Web Services Security UsernameToken Profile 1.0\" specification.", "wssec-user-name-token-type", "").AddStringEnum("Text", "Digest").AddDefaultValue("Digest").String,
			Computed:            true,
		},
		"ppsaml_validity": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the validity duration of the SAML assertion in seconds. This value and the skew time are for fine control of the validity duration. The default value is 0.", "saml-validity", "").AddDefaultValue("0").String,
			Computed:            true,
		},
		"ppsaml_skew": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the acceptable skew interval in seconds. The IdP and SP system clocks can have a skew time. When the SAML assertion is generated, the expiration takes the skew time setting into account. <ul><li>When <tt>NotBefore</tt> has the value of <tt>(CurrentTime - SkewTime)</tt> .</li><li>When <tt>NotOnOrAfter</tt> has the value of <tt>(CurrentTime + Validity + SkewTime)</tt> .</li></ul>", "saml-skew", "").AddDefaultValue("0").String,
			Computed:            true,
		},
		"ppws_username_token_include_pwd": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the WS-Security UsernameToken must include the password. By default, the token must contain the password.", "wssec-user-name-token-contains-pwd", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"ppltpa": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate an LTPA token.", "lpta-generate-token", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ppltpa_version": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the LTPA token version to generate. By default, generates a WebSphere version 2 token.", "lpta-version", "").AddStringEnum("LTPA", "LTPA1FIPS", "LTPA2", "LTPA2WAS7", "LTPADomino").AddDefaultValue("LTPA2").String,
			Computed:            true,
		},
		"ppltpa_expiry": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the lifetime of LTPA token in seconds. Enter a value in the range 1 - 628992000. The default value is 600.", "lpta-expiry", "").AddIntegerRange(1, 628992000).AddDefaultValue("600").String,
			Computed:            true,
		},
		"ppltpa_key_file": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the LTPA key file that secures the LTPA token. The LTPA key file contains the crypto material to create an LTPA token that can be consumed by WebSphere or Domino. <ul><li>For WebSphere tokens, you must export the LTPA key file from WebSphere. This file has portions encrypted by a password.</li><li>For Domino tokens, the key file should contain only the base 64-encoded Domino shared secret.</li></ul>", "lpta-key-file", "").String,
			Computed:            true,
		},
		"ppltpa_key_file_password": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Use the LTPA key file password alias.", "lpta-key-file-password", "").String,
			Computed:            true,
		},
		"ppltpa_stash_file": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the file that contains the LTPA key file password.", "lpta-stash-file", "").String,
			Computed:            true,
		},
		"pp_kerberos_spnego_token": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate an SPNEGO token to insert into the HTTP WWW-Authenticate header.", "kerberos-generate-spnego", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"pp_kerberos_bst_value_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the value for the <tt>ValueType</tt> attribute of the WS-Security BinarySecurityToken. The Kerberos AP-REQ message contains the <tt>ValueType</tt> attribute. The default value is for WSS Kerberos Token Profile 1.1 (GSS).", "kerberos-value-type", "").AddStringEnum("http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ1510", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ1510", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ4120", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ4120", "http://docs.oasis-open.org/wss/2005/xx/oasis-2005xx-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ", "http://docs.oasis-open.org/wss/2005/xx/oasis-2005xx-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ", "http://www.docs.oasis-open.org/wss/2004/07/oasis-000000-wss-kerberos-token-profile-1.0#Kerberosv5_AP_REQ").AddDefaultValue("http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ").String,
			Computed:            true,
		},
		"ppsaml_use_ws_sec": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify where to insert the SAML assertion. By default, the assertion is inserted as a child element of the SOAP header. When enabled, the assertion is inserted in a WS-Security-compliant header as defined by the WS-Security SAML token profile.", "saml-in-wssec", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"pp_kerberos_client_keytab": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the Kerberos keytab that defines the keytab for the client. This keytab is required to authenticate the client to the KDC.", "kerberos-client-keytab", "cryptokerberoskeytab").String,
			Computed:            true,
		},
		"pp_use_ws_sec": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the token can be wrapped by the WS-Security <tt>wsse:Security</tt> header. This setting for the LTPA token. By default, the token cannot be wrapped by this header. When enabled, generate a WS-Security header that contains the token.", "wssec-header-wrap-token", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"pp_actor_role_id": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the identifier for the SOAP 1.1 actor or SOAP 1.2 role for processing a WS-Security Security header. The DataPower Gateway works as that actor or role in consuming the input and generating the output for the next SOAP endpoint. This setting is meaningful when a SOAP message is being used for WS-Security 1.0 or 1.1. <table border=\"1\"><tr><td valign=\"left\">http://schemas.xmlsoap.org/soap/actor/next</td><td>Each receiver, including the intermediary and ultimate receiver, can process the Security header.</td></tr><tr><td valign=\"left\">http://www.w3.org/2003/05/soap-envelope/role/none</td><td>No one can process the Security header.</td></tr><tr><td valign=\"left\">http://www.w3.org/2003/05/soap-envelope/role/next</td><td>Each receiver, including the intermediary and ultimate receiver, can process the Security header.</td></tr><tr><td valign=\"left\">http://www.w3.org/2003/05/soap-envelope/role/ultimateReceiver</td><td>The ultimate receiver of the message can process the Security header. This value is the default value if such setting is not configured.</td></tr><tr><td valign=\"left\">&lt;blank or empty string></td><td>The empty string \"\" (without quotation marks) indicates that no actor or role identifier is configured. If no actor or role setting is configured, the ultimate receiver is assumed during message processing, and no actor or role attribute is added during the generation of the Security header. <p>This value does not generate an attribute with an empty value, which is the behavior as defined by the USE_MESSAGE_BASE_URI constant string. There cannot be more than one Security header that omits the actor or role identifier.</p></td></tr><tr><td valign=\"left\">USE_MESSAGE_BASE_URI</td><td>The constant value indicates that the actor or role identifier is the base URL of the message. If the SOAP message is transported over HTTP, the base URI is the Request-URI of the HTTP request.</td></tr><tr><td valign=\"left\">any other customized string</td><td>You can input any string to identify the actor or role of the Security header.</td></tr></table>", "wssec-actor-role-id", "").String,
			Computed:            true,
		},
		"ppws_derived_key_username_token": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate a derived key from a password. By default, a derived key is not generated. When enabled, the process adds a WS-Security derived-key UsernameToken to the message and adds an HMAC signature with the derived-key. The username and password are taken from the output of the map credentials phase.", "wssec-use-derived-key", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ppws_derived_key_username_token_iterations": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the number of hashing cycles during the generation of a derived key from a password. The minimum value is 2. The default value is 1000.", "wssec-derived-key-hash-iter", "").AddIntegerRange(2, 65535).AddDefaultValue("1000").String,
			Computed:            true,
		},
		"ppws_username_token_allow_replacement": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to retain the original token, not generate a new one, if the message already contains a UsernameToken. By default, the original otken is retained. When enabled, the generated token replaces any existing ones.", "wssec-replace-existing", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"pphmac_signing_alg": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the HMAC algorithm to sign the token. This property is available to request a WS-Security UsernameToken in postprocessing and WS-Security Derived-Key UsernameToken is added to the message with an HMAC signature. The default value is hmac-sha1.", "hmac-signing-algorithm", "").AddStringEnum("hmac-sha1", "hmac-sha224", "hmac-sha256", "hmac-sha384", "hmac-sha512", "hmac-ripemd160", "hmac-md5").AddDefaultValue("hmac-sha1").String,
			Computed:            true,
		},
		"pp_signing_hash_alg": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the algorithm for the message digest for the generation of a digital signature. This algorithm is for only the UsernameToken postprocessing method. The default value is sha1.", "message-digest-algorithm", "").AddStringEnum("sha1", "sha256", "sha512", "ripemd160", "sha224", "sha384", "md5").AddDefaultValue("sha1").String,
			Computed:            true,
		},
		"ppws_trust_header": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to return the WS-Trust token as a SOAP header. By default, the token is put in the SOAP body. When enabled, return the token as a SOAP header by wrapping the <tt>wst:RequestedSecurityToken</tt> by a <tt>wst:IssuedToken</tt> .", "ws-trust-in-header", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ppwssc_key_source": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the source of the key. For WS-Trust postprocessing, the DataPower Gateway works as an on-box WS-Trust security token service that is backed by WS-SecureConversation. A symmetric shared secret key is needed to initialize the WS-SecureConversation SecurityContext. By default, a random key is generated.", "ws-trust-key-source", "").AddStringEnum("client-entropy", "in-kerberos", "in-encryptedkey", "static", "random").AddDefaultValue("random").String,
			Computed:            true,
		},
		"pp_shared_secret_key": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the shared secret key as the WS-Trust key source.", "ws-trust-shared-key", "cryptosskey").String,
			Computed:            true,
		},
		"ppws_trust_renewal_wait": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration to allow the STS to keep an expired SecurityContext token in seconds. After a WS-Trust token expires, it can be removed from the STS and cannot be renewed. Therefore, the token must be renewed before expiry. Enter a value in the range of 0 - 2678400. The default value is 0. <p>The token is issued or renewed with a 1-hour wait time in the following situation.</p><ul><li>The WS-Trust request asks that the issued token can be renewed after expiration.</li><li>This setting has a value of 0.</li></ul>", "ws-trust-renewal-wait", "").AddIntegerRange(0, 2678400).AddDefaultValue("0").String,
			Computed:            true,
		},
		"ppws_trust_new_instance": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the STS renewal request issues a new instance for WS-Trust renewal. By default, the STS renewal request renews the existing instance. When enabled, the STS renewal request creates a new instance.", "ws-trust-new-instance", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ppws_trust_new_key": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to update the context key for WS-Trust renewal.By default, the SCT renewal request uses the existing shared secret key. When enabled, the SCT renewal request does not use the existing shared secret key.", "ws-trust-new-key", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ppws_trust_never_expire": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the WS-Trust security context expires. By default, the security context expires. When enabled, the security context never expires.However, you can change the duration afterward with an explicit duration in seconds before expiry.", "ws-trust-never-expire", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ppicrx_token": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate an Extended Identity Context Reference (ICRX) for z/OS identity propagation from the authenticated credentials. When generated, the WS-Security binary token with an ICRX token is inserted into the WS-Security header. You can use this token interoperability with the CICS Transaction Server for z/OS identity propagation support.", "generate-icrx", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ppicrx_user_realm": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the realm of a user for ICRX identity propagation. The ICRX realm is defined in the SAF configuration. Generally, this value is the equivalent of the prefix for a DN in a user registry.", "icrx-user-realm", "").String,
			Computed:            true,
		},
		"ppsaml_identity_provider": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to generate a SAML assertion. The SAML assertion can contain an authentication statement, an authorization statement, an attribute statement, or any combination of these statements. The SAML attribute value can be a user LDAP Attribute value that can be retrieved in the following ways.</p><ul><li>Directly by the LDAP authentication or authorization method with the list of LDAP attribute names that are defined by user auxiliary LDAP attributes.</li><li>Indirectly with the <tt>var://context/ldap/auxiliary-attributes</tt> variable in a stylesheet or GatewayScript file. A call with <tt>dp:ldap-search</tt> to the user registry, and put the <tt>attribute-value</tt> elements of search result to the variable.</li></ul><p>To sign the SAML assertion, configure a WS-Security sign action or SAML enveloped sign action after the AAA action in the processing rule.</p>", "generate-saml-assertion", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ppsaml_protocol": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the SAML protocol to wrap up the SAML assertion. By default, the SAML assertion can be put to WS-Security wrap-up later.", "saml-protocol", "").AddStringEnum("assertion", "response").AddDefaultValue("assertion").String,
			Computed:            true,
		},
		"ppsaml_response_destination": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the destination for a SAML response. This information can prevent malicious forwarding of requests to unintended recipients, which is a required protection by some protocol bindings. If it is present, the actual recipient must check that the URI reference identifies the location at which the message was received. If it does not check that the URI reference identifies the location, the request must be discarded. Some protocol bindings might require the use of this attribute.", "saml-response-destination", "").String,
			Computed:            true,
		},
		"pp_result_wrapup": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the method to generate the result. When the DataPower Gateway is configured for SOAP or WS-Security processing, different output methods can be used. By default, generates the results to an existing WS-Security message and replaces the same token in the requesting message.", "result-wrapup", "").AddStringEnum("wssec-replace", "wssec-new", "wssec-inject", "soap-body", "none").AddDefaultValue("wssec-replace").String,
			Computed:            true,
		},
		"ppsaml_assertion_type": GetDmSAMLStatementTypeDataSourceSchema("Specify the supported SAML statement types. By default, supports both attributes and authentication statements.", "saml-assertion-type", ""),
		"ppsaml_subject_confirm": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the method that allows the destination system to confirm the subject of the SAML assertion. By default, the subject is bearer.", "saml-subject-confirm", "").AddStringEnum("bearer", "hok", "sv").AddDefaultValue("bearer").String,
			Computed:            true,
		},
		"ppsaml_name_id": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the SAML Subject element contains the name identifier. By default, the SAML subject contains the name identifier. When disabled, the SAML subject does not contain the name identifier. Use this value if the subject confirmation method is holder-of-key because the key represent the same entity as the subject.", "saml-nid", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"ppsaml_name_id_format": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the URI reference that represents the classification of string-based identifier information. Any standard or arbitrary URI is allowed. If the value is an empty string, the DataPower Gateway attempts to determine the value from the AAA context. Some SAML protocols require a specified value, such as <tt>urn:oasis:names:tc:SAML:2.0:nameid-format:entity</tt> or <tt>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</tt> .", "saml-nid-format", "").String,
			Computed:            true,
		},
		"ppsaml_recipient": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify a URI that identifies the entity or location that an attesting entity can present the assertion to. Any standard or arbitrary URI is allowed. If the value is an empty string, the optional attribute is not generated. This setting is applicable for only SAML 2.0.", "saml-recipient", "").String,
			Computed:            true,
		},
		"ppsaml_audience": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify URI references that identify an intended audience. Enter any number of the audience URIs to process the generated SAML assertion. If the value is an empty string, the SAML audience is not restricted. If there is more than one audience URI, use a + delimiter between URIs. In this case, you must convert any URI that contains the + characters to \\+.", "saml-audience", "").String,
			Computed:            true,
		},
		"ppsaml_omit_not_before": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("whether to omit the <tt>NotBefore</tt> attribute in the SAML assertion. When omitted, the assertion is considered valid even before the time it was issued. By default, the <tt>NotBefore</tt> attribute is not omitted. When enabled, the <tt>NotBefore</tt> attribute in the SAML assertion is omitted. This behavior might be required to respond to an <tt>AuthnRequest</tt> .", "saml-omit-notbefore", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"pp_one_time_use": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the destination system or relying party should cache the generated token. The generated token might contain the property for this characteristic, which is especially practical for SAML assertions. By default, the destination system can cache the generated token. When enabled, he destination system should not cache the generated token.", "one-time-use", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ppsaml_proxy": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to allow SAML proxy restriction. The generated SAML assertion provides limitations that the asserting party imposes on relying parties that want to act as asserting parties.</p><ul><li>A relying party that acts as an asserting party can issue subsequent assertions that are based on the information in the original assertion.</li><li>The relying party cannot issue an assertion that violates these restrictions.</li></ul><p>By default, proxy restrictions are not allowd. When enabled, proxy restrictions are allows.</p>", "saml-proxy", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ppsaml_proxy_audience": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the set of audiences (proxy) to whom the asserting party permits new assertions to be issued based on this assertion. If the value is an empty string, the audience for the <tt>ProxyRestriction</tt> is not issued with this SAML assertion. If there is more than one audience URI, use a + delimiter between URIs. In this case, you must convert any URI that contains the + characters to \\+.", "saml-proxy-audience", "").String,
			Computed:            true,
		},
		"ppsaml_proxy_count": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum number of indirections that the asserting party permits between this assertion and an assertion that was issued. Enter a value in the range 0 - 65535. The default value is 0. A value of 0 indicates that a relying party must not issue an assertion to another relying party based on this assertion. If greater than zero, any assertion that is issued must itself contain a <tt>ProxyRestriction</tt> element with a <tt>Count</tt> value of at most one less than this value.", "saml-proxy-count", "").AddIntegerRange(0, 65535).AddDefaultValue("0").String,
			Computed:            true,
		},
		"ppsaml_authz_action": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the standard action that the subject can take on the resource. The SAML specification defines the list of action identifiers with corresponding namespace URIs. By default, all HTTP operations, where <tt>urn:oasis:names:tc:SAML:1.0:action:ghpp</tt> is the namespace URI.", "saml-authz-action", "").AddStringEnum("AllHTTP", "POST", "GET", "PUT", "HEAD", "General", "Read", "Write", "Execute", "Delete", "Control", "NegatedRead", "NegatedWrite", "NegatedExecute", "NegatedDelete", "NegatedControl").AddDefaultValue("AllHTTP").String,
			Computed:            true,
		},
		"ppsaml_attributes": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of an existing SAML attributes. The SAML attributes define the information to put in the SAML assertion to generate the attribute statement. Each SAML attribute requires the name, format or namespace, and value. The value can be from a DataPower variable.", "saml-attributes", "samlattributes").String,
			Computed:            true,
		},
		"ppltpa_insert_cookie": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to insert a <tt>Set-Cookie</tt> header in the response that contains the LTPA token. This setting is for generating LTPA tokens that are not wrapped in the WS-Security <tt>wsse:Security</tt> header. By default, inserts a Set-Cookie header in the response. When disabled, does not insert a Set-Cookie header in the response.", "ltpa-insert-cookie", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"pptampac_propagate": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to add the Access Manager privilege attribute certificate (PAC) token to an HTTP header. The PAC token was returned from the previous authentication or authorization phase. By default, does not add the PAC token. When enabled, adds the PAC token.", "propagate-tam-pac", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"pptam_header": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the HTTP header to store the token in. The default value is iv_creds, which is HTTP header that WebSEAL uses to write headers.", "tam-header", "").AddDefaultValue("iv-creds").String,
			Computed:            true,
		},
		"pptam_header_size": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum size in bytes of HTTP headers. A value of 0 disables this function. If the value is nonzero, the PAC token is split across multiple headers of the specified length. The default value is 0.", "tam-header-size", "").AddDefaultValue("0").String,
			Computed:            true,
		},
		"pp_kerberos_use_s4u2_proxy": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to use constrained delegation, namely S4U2Proxy, when a WS-Security Kerberos AP-REQ token or a Kerberos SPNEGO token is generated. By default, does not use constrained delegation. When enabled, uses constrained delegation.", "kerberos-use-s4u2proxy", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"pp_cookie_attributes": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the cookie attribute policy to include standard or custom attributes in the cookie. The response message that contains a <tt>Set-Cookie</tt> header is updated with the attributes defined in this policy.", "cookie-attributes", "cookieattributepolicy").String,
			Computed:            true,
		},
		"pp_kerberos_use_s4u2_self_and_s4u2_proxy": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to use protocol transition, namely S4U2Self, and then use constrained delegation, namely S4U2Proxy.</p><ul><li>Use S4U2Self to convert a non-Kerberos token to a Kerberos token to the DataPower Gateway itself.</li><li>Use S4U2Proxy to generate a WS-Security Kerberos AP-REQ token or a Kerberos SPNEGO token.</li></ul><p>By default, does not use protocol transition and constrained delegation. When enabled, uses protocol transition and constrained delegation.</p>", "kerberos-use-s4u2self", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"pp_kerberos_client_source": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify where to get the principal name of the Kerberos client. By default, uses the output of credential mapping. The client principal is based on the authenticated identity, which is followed by the corresponding realm name. For example, if the authenticated user is <tt>alice</tt> , the client principal name can be <tt>HTTP/alice.datapower.com@DATAPOWER.COM</tt> . The client principal must be present in the KDC for S4U2Self to work.", "kerberos-client-source", "").AddStringEnum("mc-output", "custom-url", "ctx-var").AddDefaultValue("mc-output").String,
			Computed:            true,
		},
		"pp_kerberos_self": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the principal name of the DataPower Gateway.", "kerberos-self-principal", "").String,
			Computed:            true,
		},
		"pp_kerberos_self_keytab": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the Kerberos keytab that defines the keytab for the DataPower Gateway. This keytab is required to authenticate the DataPower Gateway to the KDC.", "kerberos-self-keytab", "cryptokerberoskeytab").String,
			Computed:            true,
		},
		"pp_kerberos_client_custom_url": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the location of the stylesheet or GatewayScript file. This file returns the client principal name within the <tt>kerberos-client-principal</tt> element. This file gets the following input.</p><ul><li>The output of all the steps that are executed in this AAA action.</li><li>The incoming request message.</li></ul>", "kerberos-client-custom-url", "").String,
			Computed:            true,
		},
		"pp_kerberos_client_ctx_var": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the context variable. The value of this context variable is used as the Kerberos client principal. This context variable must be specified in the <tt>var://context/name</tt> format. For example, <tt>var://context/AAA/krb-client-princ</tt> . You can use the set variable action to set this variable in the processing rule before the AAA action.", "kerberos-client-ctx-var", "").String,
			Computed:            true,
		},
		"pp_kerberos_server_source": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify where to get the principal name of the Kerberos server. By default, the server principal name is the value that is specified by the Kerberos server principal property. Ensure that the server principal is in the correct format. For example, <tt>HTTP/was-backend.datapower.com@DATAPOWER.COM</tt> .", "kerberos-server-source", "").AddStringEnum("as-is-string", "custom-url", "ctx-var").AddDefaultValue("as-is-string").String,
			Computed:            true,
		},
		"pp_kerberos_server_custom_url": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the location of the stylesheet or GatewayScript file. This file returns the client principal name within the <tt>kerberos-server-principal</tt> element.</p><p>When constrained delegation is not used, this file gets the following input.</p><ul><li>The output of all phases that this AAA action processes.</li><li>The incoming request message.</li></ul><p>When constrained delegation is used, this file gets the following input.</p><ul><li>The output of only the identity extraction phase.</li><li>The incoming request message.</li></ul>", "kerberos-server-custom-url", "").String,
			Computed:            true,
		},
		"pp_kerberos_server_ctx_var": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the context variable. The value of this context variable is used as the Kerberos server principal. This context variable must be specified in the <tt>var://context/name format</tt> . For example, <tt>var:///context/AAA/krb-server-princ</tt> . You can use the set variable action to set this variable in the processing rule before the AAA action.", "kerberos-server-ctx-var", "").String,
			Computed:            true,
		},
		"ppssl_client_config_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS client type to secure connections.", "ssl-client-type", "").AddStringEnum("proxy", "client").AddDefaultValue("proxy").String,
			Computed:            true,
		},
		"ppssl_client_profile": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS client profile to secure connections.", "ssl-client", "sslclientprofile").String,
			Computed:            true,
		},
		"ppltpa_key_file_password_alias": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the the alias for password of the LTPA key file.", "ltpa-key-file-password-alias", "passwordalias").String,
			Computed:            true,
		},
		"ppjwt": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate a JWT token.", "jwt", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ppjwt_generator": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the JWT generator.", "generate-jwt", "aaajwtgenerator").String,
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
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to run a custom stylesheet or GatewayScript file.", "custom-processing", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"pp_custom_url": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the custom file for the postprocessing activity.", "custom-url", "").String,
			Optional:            true,
		},
		"ppsaml_auth_assertion": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate a SAML assertion that contains a SAML authentication statement for the authenticated user identity.", "saml-generate-assertion", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ppsaml_server_name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the value of the <tt>saml:Issuer</tt> of the generated SAML assertion or SAML SLO request. The default value is XS.</p><ul><li>If generating an SAML assertion, identifies the server that makes the assertion.</li><li>If sending an SLO request, identifies the issuer that sends the request.</li></ul>", "saml-server-name", "").AddDefaultValue("XS").String,
			Computed:            true,
			Optional:            true,
			Default:             stringdefault.StaticString("XS"),
		},
		"ppsaml_name_qualifier": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the value of the NameQualifier attribute of the NameIdentifier in the generated SAML assertion. Although the attribute is an optional attribute, some SAML implementations require that this attribute must be present.", "saml-name-qualifier", "").String,
			Optional:            true,
		},
		"pp_kerberos_ticket": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to include a WS-Security Kerberos AP-REQ BinarySecurityToken for the specified client and server principals in the WS-Security header. By default, token are not included.", "kerberos-include-token", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"pp_kerberos_client": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the client identity (cname of the Kerberos ticket) for the Kerberos client principal.", "kerberos-client-principal", "").String,
			Optional:            true,
		},
		"pp_kerberos_client_password": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "").String,
			Optional:            true,
		},
		"pp_kerberos_server": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the server identity (sname of the Kerberos ticket) for the Kerberos server principal.", "kerberos-server", "").String,
			Optional:            true,
		},
		"ppws_trust": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate the appropriate security token response for a valid WS-Trust SecurityContextToken request.", "ws-trust-generate-resp", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"pp_timestamp": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate a WS-Trust token time stamp for the security token response.", "ws-trust-add-timestamp", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"pp_timestamp_expiry": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the validity duration for the WS-Trust SCT in seconds to issue a new security context or to renew a context instance with new instance. Enter a value in the range 0 - 31622400. The default value is 0, which uses the value of the <tt>var://system/AAA/defaultexpiry</tt> variable if defined. If you did not define this variable, the value is 14400. If this setting is to renew a security context or instance, the value 0 means to use the old duration for the renewed cycle.", "ws-trust-timestamp-expiry", "").AddIntegerRange(0, 31622400).AddDefaultValue("0").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(0, 31622400),
			},
			Default: int64default.StaticInt64(0),
		},
		"pp_allow_renewal": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether WS-Trust tokens can have their lifetime period reset without a new bootstrapping authentication event. If the WS-Trust request asks to renew the issued token, this setting is ignored.", "ws-trust-allow-renewal", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ppsaml_version": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the protocol level of SAML messages. The version affects the identity extraction from the original message and the format of messages. The default value is 1.1.", "saml-version", "").AddStringEnum("2", "1.1", "1").AddDefaultValue("2").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("2", "1.1", "1"),
			},
			Default: stringdefault.StaticString("2"),
		},
		"ppsaml_send_slo": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to send a SAML Logout (SLO) request to revoke the SAML Assertion token that is used for single-sign-on (SSO). The SLO is a request-response that the DataPower&#174; Gateway handles differently when it is working as a service provider (SP) or identity provider (IdP).</p><ul><li>When an SP, the DataPower Gateway sends an SLO request to the SAML SLO endpoint (IdP). On response, the DataPower Gateway processes the SLO response for its status.</li><li>When an IdP, the request to the DataPower Gateway contains the SLO request. Postprocessing validates against the SAML metadata file and sends the corresponding endpoint the SLO response.</li></ul>", "saml-send-slo", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ppsamlslo_endpoint": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the endpoint URL for SAML 2.0 Single Logout (SLO) messages. This endpoint is the authority that authenticated the assertion subject.", "saml-slo-endpoint", "").String,
			Optional:            true,
		},
		"ppws_username_token": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to add a WS-Security UsernameToken. The username and password are taken from the output of the map credentials phase.", "wssec-add-user-name-token", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ppws_username_token_password_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the type of password that the UsernameToken provides. By default, use the digest of the password as defined in the \"Web Services Security UsernameToken Profile 1.0\" specification.", "wssec-user-name-token-type", "").AddStringEnum("Text", "Digest").AddDefaultValue("Digest").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("Text", "Digest"),
			},
			Default: stringdefault.StaticString("Digest"),
		},
		"ppsaml_validity": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the validity duration of the SAML assertion in seconds. This value and the skew time are for fine control of the validity duration. The default value is 0.", "saml-validity", "").AddDefaultValue("0").String,
			Computed:            true,
			Optional:            true,
			Default:             int64default.StaticInt64(0),
		},
		"ppsaml_skew": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the acceptable skew interval in seconds. The IdP and SP system clocks can have a skew time. When the SAML assertion is generated, the expiration takes the skew time setting into account. <ul><li>When <tt>NotBefore</tt> has the value of <tt>(CurrentTime - SkewTime)</tt> .</li><li>When <tt>NotOnOrAfter</tt> has the value of <tt>(CurrentTime + Validity + SkewTime)</tt> .</li></ul>", "saml-skew", "").AddDefaultValue("0").String,
			Computed:            true,
			Optional:            true,
			Default:             int64default.StaticInt64(0),
		},
		"ppws_username_token_include_pwd": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the WS-Security UsernameToken must include the password. By default, the token must contain the password.", "wssec-user-name-token-contains-pwd", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"ppltpa": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate an LTPA token.", "lpta-generate-token", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ppltpa_version": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the LTPA token version to generate. By default, generates a WebSphere version 2 token.", "lpta-version", "").AddStringEnum("LTPA", "LTPA1FIPS", "LTPA2", "LTPA2WAS7", "LTPADomino").AddDefaultValue("LTPA2").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("LTPA", "LTPA1FIPS", "LTPA2", "LTPA2WAS7", "LTPADomino"),
			},
			Default: stringdefault.StaticString("LTPA2"),
		},
		"ppltpa_expiry": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the lifetime of LTPA token in seconds. Enter a value in the range 1 - 628992000. The default value is 600.", "lpta-expiry", "").AddIntegerRange(1, 628992000).AddDefaultValue("600").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(1, 628992000),
			},
			Default: int64default.StaticInt64(600),
		},
		"ppltpa_key_file": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the LTPA key file that secures the LTPA token. The LTPA key file contains the crypto material to create an LTPA token that can be consumed by WebSphere or Domino. <ul><li>For WebSphere tokens, you must export the LTPA key file from WebSphere. This file has portions encrypted by a password.</li><li>For Domino tokens, the key file should contain only the base 64-encoded Domino shared secret.</li></ul>", "lpta-key-file", "").String,
			Optional:            true,
		},
		"ppltpa_key_file_password": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Use the LTPA key file password alias.", "lpta-key-file-password", "").String,
			Optional:            true,
		},
		"ppltpa_stash_file": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the file that contains the LTPA key file password.", "lpta-stash-file", "").String,
			Optional:            true,
		},
		"pp_kerberos_spnego_token": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate an SPNEGO token to insert into the HTTP WWW-Authenticate header.", "kerberos-generate-spnego", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"pp_kerberos_bst_value_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the value for the <tt>ValueType</tt> attribute of the WS-Security BinarySecurityToken. The Kerberos AP-REQ message contains the <tt>ValueType</tt> attribute. The default value is for WSS Kerberos Token Profile 1.1 (GSS).", "kerberos-value-type", "").AddStringEnum("http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ1510", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ1510", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ4120", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ4120", "http://docs.oasis-open.org/wss/2005/xx/oasis-2005xx-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ", "http://docs.oasis-open.org/wss/2005/xx/oasis-2005xx-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ", "http://www.docs.oasis-open.org/wss/2004/07/oasis-000000-wss-kerberos-token-profile-1.0#Kerberosv5_AP_REQ").AddDefaultValue("http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ1510", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ1510", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ4120", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ4120", "http://docs.oasis-open.org/wss/2005/xx/oasis-2005xx-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ", "http://docs.oasis-open.org/wss/2005/xx/oasis-2005xx-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ", "http://www.docs.oasis-open.org/wss/2004/07/oasis-000000-wss-kerberos-token-profile-1.0#Kerberosv5_AP_REQ"),
			},
			Default: stringdefault.StaticString("http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ"),
		},
		"ppsaml_use_ws_sec": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify where to insert the SAML assertion. By default, the assertion is inserted as a child element of the SOAP header. When enabled, the assertion is inserted in a WS-Security-compliant header as defined by the WS-Security SAML token profile.", "saml-in-wssec", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"pp_kerberos_client_keytab": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the Kerberos keytab that defines the keytab for the client. This keytab is required to authenticate the client to the KDC.", "kerberos-client-keytab", "cryptokerberoskeytab").String,
			Optional:            true,
		},
		"pp_use_ws_sec": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the token can be wrapped by the WS-Security <tt>wsse:Security</tt> header. This setting for the LTPA token. By default, the token cannot be wrapped by this header. When enabled, generate a WS-Security header that contains the token.", "wssec-header-wrap-token", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"pp_actor_role_id": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the identifier for the SOAP 1.1 actor or SOAP 1.2 role for processing a WS-Security Security header. The DataPower Gateway works as that actor or role in consuming the input and generating the output for the next SOAP endpoint. This setting is meaningful when a SOAP message is being used for WS-Security 1.0 or 1.1. <table border=\"1\"><tr><td valign=\"left\">http://schemas.xmlsoap.org/soap/actor/next</td><td>Each receiver, including the intermediary and ultimate receiver, can process the Security header.</td></tr><tr><td valign=\"left\">http://www.w3.org/2003/05/soap-envelope/role/none</td><td>No one can process the Security header.</td></tr><tr><td valign=\"left\">http://www.w3.org/2003/05/soap-envelope/role/next</td><td>Each receiver, including the intermediary and ultimate receiver, can process the Security header.</td></tr><tr><td valign=\"left\">http://www.w3.org/2003/05/soap-envelope/role/ultimateReceiver</td><td>The ultimate receiver of the message can process the Security header. This value is the default value if such setting is not configured.</td></tr><tr><td valign=\"left\">&lt;blank or empty string></td><td>The empty string \"\" (without quotation marks) indicates that no actor or role identifier is configured. If no actor or role setting is configured, the ultimate receiver is assumed during message processing, and no actor or role attribute is added during the generation of the Security header. <p>This value does not generate an attribute with an empty value, which is the behavior as defined by the USE_MESSAGE_BASE_URI constant string. There cannot be more than one Security header that omits the actor or role identifier.</p></td></tr><tr><td valign=\"left\">USE_MESSAGE_BASE_URI</td><td>The constant value indicates that the actor or role identifier is the base URL of the message. If the SOAP message is transported over HTTP, the base URI is the Request-URI of the HTTP request.</td></tr><tr><td valign=\"left\">any other customized string</td><td>You can input any string to identify the actor or role of the Security header.</td></tr></table>", "wssec-actor-role-id", "").String,
			Optional:            true,
		},
		"ppws_derived_key_username_token": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate a derived key from a password. By default, a derived key is not generated. When enabled, the process adds a WS-Security derived-key UsernameToken to the message and adds an HMAC signature with the derived-key. The username and password are taken from the output of the map credentials phase.", "wssec-use-derived-key", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ppws_derived_key_username_token_iterations": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the number of hashing cycles during the generation of a derived key from a password. The minimum value is 2. The default value is 1000.", "wssec-derived-key-hash-iter", "").AddIntegerRange(2, 65535).AddDefaultValue("1000").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(2, 65535),
			},
			Default: int64default.StaticInt64(1000),
		},
		"ppws_username_token_allow_replacement": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to retain the original token, not generate a new one, if the message already contains a UsernameToken. By default, the original otken is retained. When enabled, the generated token replaces any existing ones.", "wssec-replace-existing", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"pphmac_signing_alg": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the HMAC algorithm to sign the token. This property is available to request a WS-Security UsernameToken in postprocessing and WS-Security Derived-Key UsernameToken is added to the message with an HMAC signature. The default value is hmac-sha1.", "hmac-signing-algorithm", "").AddStringEnum("hmac-sha1", "hmac-sha224", "hmac-sha256", "hmac-sha384", "hmac-sha512", "hmac-ripemd160", "hmac-md5").AddDefaultValue("hmac-sha1").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("hmac-sha1", "hmac-sha224", "hmac-sha256", "hmac-sha384", "hmac-sha512", "hmac-ripemd160", "hmac-md5"),
			},
			Default: stringdefault.StaticString("hmac-sha1"),
		},
		"pp_signing_hash_alg": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the algorithm for the message digest for the generation of a digital signature. This algorithm is for only the UsernameToken postprocessing method. The default value is sha1.", "message-digest-algorithm", "").AddStringEnum("sha1", "sha256", "sha512", "ripemd160", "sha224", "sha384", "md5").AddDefaultValue("sha1").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("sha1", "sha256", "sha512", "ripemd160", "sha224", "sha384", "md5"),
			},
			Default: stringdefault.StaticString("sha1"),
		},
		"ppws_trust_header": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to return the WS-Trust token as a SOAP header. By default, the token is put in the SOAP body. When enabled, return the token as a SOAP header by wrapping the <tt>wst:RequestedSecurityToken</tt> by a <tt>wst:IssuedToken</tt> .", "ws-trust-in-header", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ppwssc_key_source": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the source of the key. For WS-Trust postprocessing, the DataPower Gateway works as an on-box WS-Trust security token service that is backed by WS-SecureConversation. A symmetric shared secret key is needed to initialize the WS-SecureConversation SecurityContext. By default, a random key is generated.", "ws-trust-key-source", "").AddStringEnum("client-entropy", "in-kerberos", "in-encryptedkey", "static", "random").AddDefaultValue("random").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("client-entropy", "in-kerberos", "in-encryptedkey", "static", "random"),
			},
			Default: stringdefault.StaticString("random"),
		},
		"pp_shared_secret_key": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the shared secret key as the WS-Trust key source.", "ws-trust-shared-key", "cryptosskey").String,
			Optional:            true,
		},
		"ppws_trust_renewal_wait": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration to allow the STS to keep an expired SecurityContext token in seconds. After a WS-Trust token expires, it can be removed from the STS and cannot be renewed. Therefore, the token must be renewed before expiry. Enter a value in the range of 0 - 2678400. The default value is 0. <p>The token is issued or renewed with a 1-hour wait time in the following situation.</p><ul><li>The WS-Trust request asks that the issued token can be renewed after expiration.</li><li>This setting has a value of 0.</li></ul>", "ws-trust-renewal-wait", "").AddIntegerRange(0, 2678400).AddDefaultValue("0").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(0, 2678400),
			},
			Default: int64default.StaticInt64(0),
		},
		"ppws_trust_new_instance": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the STS renewal request issues a new instance for WS-Trust renewal. By default, the STS renewal request renews the existing instance. When enabled, the STS renewal request creates a new instance.", "ws-trust-new-instance", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ppws_trust_new_key": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to update the context key for WS-Trust renewal.By default, the SCT renewal request uses the existing shared secret key. When enabled, the SCT renewal request does not use the existing shared secret key.", "ws-trust-new-key", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ppws_trust_never_expire": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the WS-Trust security context expires. By default, the security context expires. When enabled, the security context never expires.However, you can change the duration afterward with an explicit duration in seconds before expiry.", "ws-trust-never-expire", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ppicrx_token": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate an Extended Identity Context Reference (ICRX) for z/OS identity propagation from the authenticated credentials. When generated, the WS-Security binary token with an ICRX token is inserted into the WS-Security header. You can use this token interoperability with the CICS Transaction Server for z/OS identity propagation support.", "generate-icrx", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ppicrx_user_realm": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the realm of a user for ICRX identity propagation. The ICRX realm is defined in the SAF configuration. Generally, this value is the equivalent of the prefix for a DN in a user registry.", "icrx-user-realm", "").String,
			Optional:            true,
		},
		"ppsaml_identity_provider": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to generate a SAML assertion. The SAML assertion can contain an authentication statement, an authorization statement, an attribute statement, or any combination of these statements. The SAML attribute value can be a user LDAP Attribute value that can be retrieved in the following ways.</p><ul><li>Directly by the LDAP authentication or authorization method with the list of LDAP attribute names that are defined by user auxiliary LDAP attributes.</li><li>Indirectly with the <tt>var://context/ldap/auxiliary-attributes</tt> variable in a stylesheet or GatewayScript file. A call with <tt>dp:ldap-search</tt> to the user registry, and put the <tt>attribute-value</tt> elements of search result to the variable.</li></ul><p>To sign the SAML assertion, configure a WS-Security sign action or SAML enveloped sign action after the AAA action in the processing rule.</p>", "generate-saml-assertion", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ppsaml_protocol": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the SAML protocol to wrap up the SAML assertion. By default, the SAML assertion can be put to WS-Security wrap-up later.", "saml-protocol", "").AddStringEnum("assertion", "response").AddDefaultValue("assertion").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("assertion", "response"),
			},
			Default: stringdefault.StaticString("assertion"),
		},
		"ppsaml_response_destination": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the destination for a SAML response. This information can prevent malicious forwarding of requests to unintended recipients, which is a required protection by some protocol bindings. If it is present, the actual recipient must check that the URI reference identifies the location at which the message was received. If it does not check that the URI reference identifies the location, the request must be discarded. Some protocol bindings might require the use of this attribute.", "saml-response-destination", "").String,
			Optional:            true,
		},
		"pp_result_wrapup": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the method to generate the result. When the DataPower Gateway is configured for SOAP or WS-Security processing, different output methods can be used. By default, generates the results to an existing WS-Security message and replaces the same token in the requesting message.", "result-wrapup", "").AddStringEnum("wssec-replace", "wssec-new", "wssec-inject", "soap-body", "none").AddDefaultValue("wssec-replace").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("wssec-replace", "wssec-new", "wssec-inject", "soap-body", "none"),
			},
			Default: stringdefault.StaticString("wssec-replace"),
		},
		"ppsaml_assertion_type": GetDmSAMLStatementTypeResourceSchema("Specify the supported SAML statement types. By default, supports both attributes and authentication statements.", "saml-assertion-type", "", false),
		"ppsaml_subject_confirm": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the method that allows the destination system to confirm the subject of the SAML assertion. By default, the subject is bearer.", "saml-subject-confirm", "").AddStringEnum("bearer", "hok", "sv").AddDefaultValue("bearer").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("bearer", "hok", "sv"),
			},
			Default: stringdefault.StaticString("bearer"),
		},
		"ppsaml_name_id": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the SAML Subject element contains the name identifier. By default, the SAML subject contains the name identifier. When disabled, the SAML subject does not contain the name identifier. Use this value if the subject confirmation method is holder-of-key because the key represent the same entity as the subject.", "saml-nid", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"ppsaml_name_id_format": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the URI reference that represents the classification of string-based identifier information. Any standard or arbitrary URI is allowed. If the value is an empty string, the DataPower Gateway attempts to determine the value from the AAA context. Some SAML protocols require a specified value, such as <tt>urn:oasis:names:tc:SAML:2.0:nameid-format:entity</tt> or <tt>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</tt> .", "saml-nid-format", "").String,
			Optional:            true,
		},
		"ppsaml_recipient": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify a URI that identifies the entity or location that an attesting entity can present the assertion to. Any standard or arbitrary URI is allowed. If the value is an empty string, the optional attribute is not generated. This setting is applicable for only SAML 2.0.", "saml-recipient", "").String,
			Optional:            true,
		},
		"ppsaml_audience": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify URI references that identify an intended audience. Enter any number of the audience URIs to process the generated SAML assertion. If the value is an empty string, the SAML audience is not restricted. If there is more than one audience URI, use a + delimiter between URIs. In this case, you must convert any URI that contains the + characters to \\+.", "saml-audience", "").String,
			Optional:            true,
		},
		"ppsaml_omit_not_before": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("whether to omit the <tt>NotBefore</tt> attribute in the SAML assertion. When omitted, the assertion is considered valid even before the time it was issued. By default, the <tt>NotBefore</tt> attribute is not omitted. When enabled, the <tt>NotBefore</tt> attribute in the SAML assertion is omitted. This behavior might be required to respond to an <tt>AuthnRequest</tt> .", "saml-omit-notbefore", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"pp_one_time_use": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the destination system or relying party should cache the generated token. The generated token might contain the property for this characteristic, which is especially practical for SAML assertions. By default, the destination system can cache the generated token. When enabled, he destination system should not cache the generated token.", "one-time-use", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ppsaml_proxy": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to allow SAML proxy restriction. The generated SAML assertion provides limitations that the asserting party imposes on relying parties that want to act as asserting parties.</p><ul><li>A relying party that acts as an asserting party can issue subsequent assertions that are based on the information in the original assertion.</li><li>The relying party cannot issue an assertion that violates these restrictions.</li></ul><p>By default, proxy restrictions are not allowd. When enabled, proxy restrictions are allows.</p>", "saml-proxy", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ppsaml_proxy_audience": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the set of audiences (proxy) to whom the asserting party permits new assertions to be issued based on this assertion. If the value is an empty string, the audience for the <tt>ProxyRestriction</tt> is not issued with this SAML assertion. If there is more than one audience URI, use a + delimiter between URIs. In this case, you must convert any URI that contains the + characters to \\+.", "saml-proxy-audience", "").String,
			Optional:            true,
		},
		"ppsaml_proxy_count": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum number of indirections that the asserting party permits between this assertion and an assertion that was issued. Enter a value in the range 0 - 65535. The default value is 0. A value of 0 indicates that a relying party must not issue an assertion to another relying party based on this assertion. If greater than zero, any assertion that is issued must itself contain a <tt>ProxyRestriction</tt> element with a <tt>Count</tt> value of at most one less than this value.", "saml-proxy-count", "").AddIntegerRange(0, 65535).AddDefaultValue("0").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(0, 65535),
			},
			Default: int64default.StaticInt64(0),
		},
		"ppsaml_authz_action": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the standard action that the subject can take on the resource. The SAML specification defines the list of action identifiers with corresponding namespace URIs. By default, all HTTP operations, where <tt>urn:oasis:names:tc:SAML:1.0:action:ghpp</tt> is the namespace URI.", "saml-authz-action", "").AddStringEnum("AllHTTP", "POST", "GET", "PUT", "HEAD", "General", "Read", "Write", "Execute", "Delete", "Control", "NegatedRead", "NegatedWrite", "NegatedExecute", "NegatedDelete", "NegatedControl").AddDefaultValue("AllHTTP").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("AllHTTP", "POST", "GET", "PUT", "HEAD", "General", "Read", "Write", "Execute", "Delete", "Control", "NegatedRead", "NegatedWrite", "NegatedExecute", "NegatedDelete", "NegatedControl"),
			},
			Default: stringdefault.StaticString("AllHTTP"),
		},
		"ppsaml_attributes": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of an existing SAML attributes. The SAML attributes define the information to put in the SAML assertion to generate the attribute statement. Each SAML attribute requires the name, format or namespace, and value. The value can be from a DataPower variable.", "saml-attributes", "samlattributes").String,
			Optional:            true,
		},
		"ppltpa_insert_cookie": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to insert a <tt>Set-Cookie</tt> header in the response that contains the LTPA token. This setting is for generating LTPA tokens that are not wrapped in the WS-Security <tt>wsse:Security</tt> header. By default, inserts a Set-Cookie header in the response. When disabled, does not insert a Set-Cookie header in the response.", "ltpa-insert-cookie", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"pptampac_propagate": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to add the Access Manager privilege attribute certificate (PAC) token to an HTTP header. The PAC token was returned from the previous authentication or authorization phase. By default, does not add the PAC token. When enabled, adds the PAC token.", "propagate-tam-pac", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"pptam_header": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the HTTP header to store the token in. The default value is iv_creds, which is HTTP header that WebSEAL uses to write headers.", "tam-header", "").AddDefaultValue("iv-creds").String,
			Computed:            true,
			Optional:            true,
			Default:             stringdefault.StaticString("iv-creds"),
		},
		"pptam_header_size": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum size in bytes of HTTP headers. A value of 0 disables this function. If the value is nonzero, the PAC token is split across multiple headers of the specified length. The default value is 0.", "tam-header-size", "").AddDefaultValue("0").String,
			Computed:            true,
			Optional:            true,
			Default:             int64default.StaticInt64(0),
		},
		"pp_kerberos_use_s4u2_proxy": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to use constrained delegation, namely S4U2Proxy, when a WS-Security Kerberos AP-REQ token or a Kerberos SPNEGO token is generated. By default, does not use constrained delegation. When enabled, uses constrained delegation.", "kerberos-use-s4u2proxy", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"pp_cookie_attributes": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the cookie attribute policy to include standard or custom attributes in the cookie. The response message that contains a <tt>Set-Cookie</tt> header is updated with the attributes defined in this policy.", "cookie-attributes", "cookieattributepolicy").String,
			Optional:            true,
		},
		"pp_kerberos_use_s4u2_self_and_s4u2_proxy": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to use protocol transition, namely S4U2Self, and then use constrained delegation, namely S4U2Proxy.</p><ul><li>Use S4U2Self to convert a non-Kerberos token to a Kerberos token to the DataPower Gateway itself.</li><li>Use S4U2Proxy to generate a WS-Security Kerberos AP-REQ token or a Kerberos SPNEGO token.</li></ul><p>By default, does not use protocol transition and constrained delegation. When enabled, uses protocol transition and constrained delegation.</p>", "kerberos-use-s4u2self", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"pp_kerberos_client_source": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify where to get the principal name of the Kerberos client. By default, uses the output of credential mapping. The client principal is based on the authenticated identity, which is followed by the corresponding realm name. For example, if the authenticated user is <tt>alice</tt> , the client principal name can be <tt>HTTP/alice.datapower.com@DATAPOWER.COM</tt> . The client principal must be present in the KDC for S4U2Self to work.", "kerberos-client-source", "").AddStringEnum("mc-output", "custom-url", "ctx-var").AddDefaultValue("mc-output").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("mc-output", "custom-url", "ctx-var"),
			},
			Default: stringdefault.StaticString("mc-output"),
		},
		"pp_kerberos_self": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the principal name of the DataPower Gateway.", "kerberos-self-principal", "").String,
			Optional:            true,
		},
		"pp_kerberos_self_keytab": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the Kerberos keytab that defines the keytab for the DataPower Gateway. This keytab is required to authenticate the DataPower Gateway to the KDC.", "kerberos-self-keytab", "cryptokerberoskeytab").String,
			Optional:            true,
		},
		"pp_kerberos_client_custom_url": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the location of the stylesheet or GatewayScript file. This file returns the client principal name within the <tt>kerberos-client-principal</tt> element. This file gets the following input.</p><ul><li>The output of all the steps that are executed in this AAA action.</li><li>The incoming request message.</li></ul>", "kerberos-client-custom-url", "").String,
			Optional:            true,
		},
		"pp_kerberos_client_ctx_var": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the context variable. The value of this context variable is used as the Kerberos client principal. This context variable must be specified in the <tt>var://context/name</tt> format. For example, <tt>var://context/AAA/krb-client-princ</tt> . You can use the set variable action to set this variable in the processing rule before the AAA action.", "kerberos-client-ctx-var", "").String,
			Optional:            true,
		},
		"pp_kerberos_server_source": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify where to get the principal name of the Kerberos server. By default, the server principal name is the value that is specified by the Kerberos server principal property. Ensure that the server principal is in the correct format. For example, <tt>HTTP/was-backend.datapower.com@DATAPOWER.COM</tt> .", "kerberos-server-source", "").AddStringEnum("as-is-string", "custom-url", "ctx-var").AddDefaultValue("as-is-string").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("as-is-string", "custom-url", "ctx-var"),
			},
			Default: stringdefault.StaticString("as-is-string"),
		},
		"pp_kerberos_server_custom_url": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the location of the stylesheet or GatewayScript file. This file returns the client principal name within the <tt>kerberos-server-principal</tt> element.</p><p>When constrained delegation is not used, this file gets the following input.</p><ul><li>The output of all phases that this AAA action processes.</li><li>The incoming request message.</li></ul><p>When constrained delegation is used, this file gets the following input.</p><ul><li>The output of only the identity extraction phase.</li><li>The incoming request message.</li></ul>", "kerberos-server-custom-url", "").String,
			Optional:            true,
		},
		"pp_kerberos_server_ctx_var": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the context variable. The value of this context variable is used as the Kerberos server principal. This context variable must be specified in the <tt>var://context/name format</tt> . For example, <tt>var:///context/AAA/krb-server-princ</tt> . You can use the set variable action to set this variable in the processing rule before the AAA action.", "kerberos-server-ctx-var", "").String,
			Optional:            true,
		},
		"ppssl_client_config_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS client type to secure connections.", "ssl-client-type", "").AddStringEnum("proxy", "client").AddDefaultValue("proxy").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("proxy", "client"),
			},
			Default: stringdefault.StaticString("proxy"),
		},
		"ppssl_client_profile": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS client profile to secure connections.", "ssl-client", "sslclientprofile").String,
			Optional:            true,
		},
		"ppltpa_key_file_password_alias": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the the alias for password of the LTPA key file.", "ltpa-key-file-password-alias", "passwordalias").String,
			Optional:            true,
		},
		"ppjwt": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate a JWT token.", "jwt", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ppjwt_generator": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the JWT generator.", "generate-jwt", "aaajwtgenerator").String,
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
	if !data.PpwsDerivedKeyUsernameToken.IsNull() {
		return false
	}
	if !data.PpwsDerivedKeyUsernameTokenIterations.IsNull() {
		return false
	}
	if !data.PpwsUsernameTokenAllowReplacement.IsNull() {
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
	if !data.PpwsDerivedKeyUsernameToken.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPWSDerivedKeyUsernameToken`, tfutils.StringFromBool(data.PpwsDerivedKeyUsernameToken, ""))
	}
	if !data.PpwsDerivedKeyUsernameTokenIterations.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPWSDerivedKeyUsernameTokenIterations`, data.PpwsDerivedKeyUsernameTokenIterations.ValueInt64())
	}
	if !data.PpwsUsernameTokenAllowReplacement.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPWSUsernameTokenAllowReplacement`, tfutils.StringFromBool(data.PpwsUsernameTokenAllowReplacement, ""))
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
