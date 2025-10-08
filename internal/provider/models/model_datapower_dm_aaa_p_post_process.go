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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmAAAPPostProcess struct {
	PpEnabled                             types.Bool           `tfsdk:"pp_enabled"`
	PpCustomUrl                           types.String         `tfsdk:"pp_custom_url"`
	PpSamlAuthAssertion                   types.Bool           `tfsdk:"pp_saml_auth_assertion"`
	PpSamlServerName                      types.String         `tfsdk:"pp_saml_server_name"`
	PpSamlNameQualifier                   types.String         `tfsdk:"pp_saml_name_qualifier"`
	PpKerberosTicket                      types.Bool           `tfsdk:"pp_kerberos_ticket"`
	PpKerberosClient                      types.String         `tfsdk:"pp_kerberos_client"`
	PpKerberosClientPassword              types.String         `tfsdk:"pp_kerberos_client_password"`
	PpKerberosServer                      types.String         `tfsdk:"pp_kerberos_server"`
	PpWsTrust                             types.Bool           `tfsdk:"pp_ws_trust"`
	PpTimestamp                           types.Bool           `tfsdk:"pp_timestamp"`
	PpTimestampExpiry                     types.Int64          `tfsdk:"pp_timestamp_expiry"`
	PpAllowRenewal                        types.Bool           `tfsdk:"pp_allow_renewal"`
	PpSamlVersion                         types.String         `tfsdk:"pp_saml_version"`
	PpSamlSendSlo                         types.Bool           `tfsdk:"pp_saml_send_slo"`
	PpSamlSloEndpoint                     types.String         `tfsdk:"pp_saml_slo_endpoint"`
	PpWsUsernameToken                     types.Bool           `tfsdk:"pp_ws_username_token"`
	PpWsUsernameTokenPasswordType         types.String         `tfsdk:"pp_ws_username_token_password_type"`
	PpSamlValidity                        types.Int64          `tfsdk:"pp_saml_validity"`
	PpSamlSkew                            types.Int64          `tfsdk:"pp_saml_skew"`
	PpWsUsernameTokenIncludePwd           types.Bool           `tfsdk:"pp_ws_username_token_include_pwd"`
	PpLtpa                                types.Bool           `tfsdk:"pp_ltpa"`
	PpLtpaVersion                         types.String         `tfsdk:"pp_ltpa_version"`
	PpLtpaExpiry                          types.Int64          `tfsdk:"pp_ltpa_expiry"`
	PpLtpaKeyFile                         types.String         `tfsdk:"pp_ltpa_key_file"`
	PpLtpaKeyFilePassword                 types.String         `tfsdk:"pp_ltpa_key_file_password"`
	PpLtpaStashFile                       types.String         `tfsdk:"pp_ltpa_stash_file"`
	PpKerberosSpnegoToken                 types.Bool           `tfsdk:"pp_kerberos_spnego_token"`
	PpKerberosBstValueType                types.String         `tfsdk:"pp_kerberos_bst_value_type"`
	PpSamlUseWsSec                        types.Bool           `tfsdk:"pp_saml_use_ws_sec"`
	PpKerberosClientKeytab                types.String         `tfsdk:"pp_kerberos_client_keytab"`
	PpUseWsSec                            types.Bool           `tfsdk:"pp_use_ws_sec"`
	PpActorRoleId                         types.String         `tfsdk:"pp_actor_role_id"`
	PpWsDerivedKeyUsernameToken           types.Bool           `tfsdk:"pp_ws_derived_key_username_token"`
	PpWsDerivedKeyUsernameTokenIterations types.Int64          `tfsdk:"pp_ws_derived_key_username_token_iterations"`
	PpWsUsernameTokenAllowReplacement     types.Bool           `tfsdk:"pp_ws_username_token_allow_replacement"`
	PpHmacSigningAlg                      types.String         `tfsdk:"pp_hmac_signing_alg"`
	PpSigningHashAlg                      types.String         `tfsdk:"pp_signing_hash_alg"`
	PpWsTrustHeader                       types.Bool           `tfsdk:"pp_ws_trust_header"`
	PpWsScKeySource                       types.String         `tfsdk:"pp_ws_sc_key_source"`
	PpSharedSecretKey                     types.String         `tfsdk:"pp_shared_secret_key"`
	PpWsTrustRenewalWait                  types.Int64          `tfsdk:"pp_ws_trust_renewal_wait"`
	PpWsTrustNewInstance                  types.Bool           `tfsdk:"pp_ws_trust_new_instance"`
	PpWsTrustNewKey                       types.Bool           `tfsdk:"pp_ws_trust_new_key"`
	PpWsTrustNeverExpire                  types.Bool           `tfsdk:"pp_ws_trust_never_expire"`
	PpicrxToken                           types.Bool           `tfsdk:"ppicrx_token"`
	PpicrxUserRealm                       types.String         `tfsdk:"ppicrx_user_realm"`
	PpSamlIdentityProvider                types.Bool           `tfsdk:"pp_saml_identity_provider"`
	PpSamlProtocol                        types.String         `tfsdk:"pp_saml_protocol"`
	PpSamlResponseDestination             types.String         `tfsdk:"pp_saml_response_destination"`
	PpResultWrapup                        types.String         `tfsdk:"pp_result_wrapup"`
	PpSamlAssertionType                   *DmSAMLStatementType `tfsdk:"pp_saml_assertion_type"`
	PpSamlSubjectConfirm                  types.String         `tfsdk:"pp_saml_subject_confirm"`
	PpSamlNameId                          types.Bool           `tfsdk:"pp_saml_name_id"`
	PpSamlNameIdFormat                    types.String         `tfsdk:"pp_saml_name_id_format"`
	PpSamlRecipient                       types.String         `tfsdk:"pp_saml_recipient"`
	PpSamlAudience                        types.String         `tfsdk:"pp_saml_audience"`
	PpSamlOmitNotBefore                   types.Bool           `tfsdk:"pp_saml_omit_not_before"`
	PpOneTimeUse                          types.Bool           `tfsdk:"pp_one_time_use"`
	PpSamlProxy                           types.Bool           `tfsdk:"pp_saml_proxy"`
	PpSamlProxyAudience                   types.String         `tfsdk:"pp_saml_proxy_audience"`
	PpSamlProxyCount                      types.Int64          `tfsdk:"pp_saml_proxy_count"`
	PpSamlAuthzAction                     types.String         `tfsdk:"pp_saml_authz_action"`
	PpSamlAttributes                      types.String         `tfsdk:"pp_saml_attributes"`
	PpLtpaInsertCookie                    types.Bool           `tfsdk:"pp_ltpa_insert_cookie"`
	PpTamPacPropagate                     types.Bool           `tfsdk:"pp_tam_pac_propagate"`
	PpTamHeader                           types.String         `tfsdk:"pp_tam_header"`
	PpTamHeaderSize                       types.Int64          `tfsdk:"pp_tam_header_size"`
	PpKerberosUseS4u2proxy                types.Bool           `tfsdk:"pp_kerberos_use_s4u2proxy"`
	PpCookieAttributes                    types.String         `tfsdk:"pp_cookie_attributes"`
	PpKerberosUseS4u2selfAndS4u2proxy     types.Bool           `tfsdk:"pp_kerberos_use_s4u2self_and_s4u2proxy"`
	PpKerberosClientSource                types.String         `tfsdk:"pp_kerberos_client_source"`
	PpKerberosSelf                        types.String         `tfsdk:"pp_kerberos_self"`
	PpKerberosSelfKeytab                  types.String         `tfsdk:"pp_kerberos_self_keytab"`
	PpKerberosClientCustomUrl             types.String         `tfsdk:"pp_kerberos_client_custom_url"`
	PpKerberosClientCtxVar                types.String         `tfsdk:"pp_kerberos_client_ctx_var"`
	PpKerberosServerSource                types.String         `tfsdk:"pp_kerberos_server_source"`
	PpKerberosServerCustomUrl             types.String         `tfsdk:"pp_kerberos_server_custom_url"`
	PpKerberosServerCtxVar                types.String         `tfsdk:"pp_kerberos_server_ctx_var"`
	PpSslClientConfigType                 types.String         `tfsdk:"pp_ssl_client_config_type"`
	PpSslClientProfile                    types.String         `tfsdk:"pp_ssl_client_profile"`
	PpLtpaKeyFilePasswordAlias            types.String         `tfsdk:"pp_ltpa_key_file_password_alias"`
	PpJwt                                 types.Bool           `tfsdk:"pp_jwt"`
	PpJwtGenerator                        types.String         `tfsdk:"pp_jwt_generator"`
}

var DmAAAPPostProcessPPCustomURLCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "pp_enabled",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var DmAAAPPostProcessPPCustomURLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPPostProcessPPSAMLServerNameIgnoreVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_saml_identity_provider",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_saml_auth_assertion",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_saml_send_slo",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
	},
}

var DmAAAPPostProcessPPSAMLNameQualifierIgnoreVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_saml_auth_assertion",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_saml_identity_provider",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"false"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_saml_name_id",
					AttrType:    "Bool",
					AttrDefault: "true",
					Value:       []string{"false"},
				},
			},
		},
	},
}

var DmAAAPPostProcessPPKerberosClientCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_kerberos_ticket",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_kerberos_spnego_token",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},
			},
		},
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_kerberos_use_s4u2proxy",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"false"},
				},

				{
					Evaluation: "logical-and",
					Conditions: []validators.Evaluation{

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "pp_kerberos_use_s4u2proxy",
							AttrType:    "Bool",
							AttrDefault: "false",
							Value:       []string{"true"},
						},

						{
							Evaluation:  "property-value-not-in-list",
							Attribute:   "au_method",
							AttrType:    "String",
							AttrDefault: "ldap",
							AttrPath:    "../authenticate",
							Value:       []string{"kerberos"},
						},
					},
				},
			},
		},
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_kerberos_use_s4u2self_and_s4u2proxy",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"false"},
				},

				{
					Evaluation: "logical-and",
					Conditions: []validators.Evaluation{

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "pp_kerberos_use_s4u2self_and_s4u2proxy",
							AttrType:    "Bool",
							AttrDefault: "false",
							Value:       []string{"true"},
						},

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "au_method",
							AttrType:    "String",
							AttrDefault: "ldap",
							AttrPath:    "../authenticate",
							Value:       []string{"kerberos"},
						},
					},
				},
			},
		},
	},
}

var DmAAAPPostProcessPPKerberosClientIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPPostProcessPPKerberosClientPasswordIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPPostProcessPPKerberosServerCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_kerberos_ticket",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_kerberos_spnego_token",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},
			},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "pp_kerberos_server_source",
			AttrType:    "String",
			AttrDefault: "as-is-string",
			Value:       []string{"custom-url"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "pp_kerberos_server_source",
			AttrType:    "String",
			AttrDefault: "as-is-string",
			Value:       []string{"ctx-var"},
		},
	},
}

var DmAAAPPostProcessPPKerberosServerIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPPostProcessPPTimestampIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "pp_ws_trust",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var DmAAAPPostProcessPPTimestampExpiryIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_ws_trust",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_ws_trust_never_expire",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
	},
}

var DmAAAPPostProcessPPAllowRenewalIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_ws_trust",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_ws_trust_never_expire",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
	},
}

var DmAAAPPostProcessPPSAMLVersionIgnoreVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_saml_identity_provider",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_saml_auth_assertion",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
	},
}

var DmAAAPPostProcessPPSAMLSLOEndpointIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "pp_saml_send_slo",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var DmAAAPPostProcessPPWSUsernameTokenPasswordTypeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_ws_username_token",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_ws_username_token_include_pwd",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"false"},
		},
	},
}

var DmAAAPPostProcessPPSAMLValidityIgnoreVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_saml_identity_provider",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_saml_auth_assertion",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
	},
}

var DmAAAPPostProcessPPSAMLSkewIgnoreVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_saml_identity_provider",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_saml_auth_assertion",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
	},
}

var DmAAAPPostProcessPPWSUsernameTokenIncludePwdIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "pp_ws_username_token",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var DmAAAPPostProcessPPLTPAVersionIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "pp_ltpa",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var DmAAAPPostProcessPPLTPAExpiryIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "pp_ltpa",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var DmAAAPPostProcessPPLTPAKeyFileCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "pp_ltpa",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var DmAAAPPostProcessPPLTPAKeyFileIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "pp_ltpa",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var DmAAAPPostProcessPPLTPAKeyFilePasswordIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_ltpa",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"false"},
				},

				{
					Evaluation: "logical-and",
					Conditions: []validators.Evaluation{

						{
							Evaluation:  "property-value-not-in-list",
							Attribute:   "pp_ltpa_version",
							AttrType:    "String",
							AttrDefault: "LTPA2",
							Value:       []string{"LTPA", "LTPA1FIPS", "LTPA2", "LTPA2WAS7"},
						},

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "pp_ltpa_version",
							AttrType:    "String",
							AttrDefault: "LTPA2",
							Value:       []string{"LTPADomino"},
						},
					},
				},
			},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_ltpa_key_file_password",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{""},
		},
	},
}

var DmAAAPPostProcessPPLTPAStashFileIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPPostProcessPPKerberosBstValueTypeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "pp_kerberos_ticket",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var DmAAAPPostProcessPPKerberosBstValueTypeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPPostProcessPPSAMLUseWSSecIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "pp_saml_auth_assertion",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var DmAAAPPostProcessPPKerberosClientKeytabCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_kerberos_ticket",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_kerberos_spnego_token",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},
			},
		},
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_kerberos_use_s4u2proxy",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"false"},
				},

				{
					Evaluation: "logical-and",
					Conditions: []validators.Evaluation{

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "pp_kerberos_use_s4u2proxy",
							AttrType:    "Bool",
							AttrDefault: "false",
							Value:       []string{"true"},
						},

						{
							Evaluation:  "property-value-not-in-list",
							Attribute:   "au_method",
							AttrType:    "String",
							AttrDefault: "ldap",
							AttrPath:    "../authenticate",
							Value:       []string{"kerberos"},
						},
					},
				},
			},
		},
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_kerberos_use_s4u2self_and_s4u2proxy",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"false"},
				},

				{
					Evaluation: "logical-and",
					Conditions: []validators.Evaluation{

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "pp_kerberos_use_s4u2self_and_s4u2proxy",
							AttrType:    "Bool",
							AttrDefault: "false",
							Value:       []string{"true"},
						},

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "au_method",
							AttrType:    "String",
							AttrDefault: "ldap",
							AttrPath:    "../authenticate",
							Value:       []string{"kerberos"},
						},
					},
				},
			},
		},
	},
}

var DmAAAPPostProcessPPKerberosClientKeytabIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPPostProcessPPUseWSSecIgnoreVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_ltpa",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"false"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_ltpa_version",
					AttrType:    "String",
					AttrDefault: "LTPA2",
					Value:       []string{"LTPADomino"},
				},
			},
		},
	},
}

var DmAAAPPostProcessPPActorRoleIDIgnoreVal = validators.Evaluation{
	Evaluation: "logical-not",
	Conditions: []validators.Evaluation{
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_ws_username_token",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "ppicrx_token",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_saml_identity_provider",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},

				{
					Evaluation: "logical-and",
					Conditions: []validators.Evaluation{

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "pp_ltpa",
							AttrType:    "Bool",
							AttrDefault: "false",
							Value:       []string{"true"},
						},

						{
							Evaluation:  "property-value-not-in-list",
							Attribute:   "pp_ltpa_version",
							AttrType:    "String",
							AttrDefault: "LTPA2",
							Value:       []string{"LTPADomino"},
						},

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "pp_use_ws_sec",
							AttrType:    "Bool",
							AttrDefault: "false",
							Value:       []string{"true"},
						},
					},
				},
			},
		},
	},
}

var DmAAAPPostProcessPPWSDerivedKeyUsernameTokenCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_ws_username_token",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_ws_username_token_include_pwd",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"false"},
		},
	},
}

var DmAAAPPostProcessPPWSDerivedKeyUsernameTokenIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPPostProcessPPWSDerivedKeyUsernameTokenIterationsCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_ws_username_token",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_ws_username_token_include_pwd",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_ws_derived_key_username_token",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
	},
}

var DmAAAPPostProcessPPWSDerivedKeyUsernameTokenIterationsIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPPostProcessPPWSUsernameTokenAllowReplacementIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "pp_ws_username_token",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var DmAAAPPostProcessPPHMACSigningAlgCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_ws_username_token",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_ws_username_token_include_pwd",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_ws_derived_key_username_token",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
	},
}

var DmAAAPPostProcessPPHMACSigningAlgIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPPostProcessPPSigningHashAlgCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_ws_username_token",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_ws_username_token_include_pwd",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_ws_derived_key_username_token",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
	},
}

var DmAAAPPostProcessPPSigningHashAlgIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPPostProcessPPWSTrustHeaderIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "pp_ws_trust",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var DmAAAPPostProcessPPWSSCKeySourceIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "pp_ws_trust",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var DmAAAPPostProcessPPSharedSecretKeyIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "pp_ws_sc_key_source",
	AttrType:    "String",
	AttrDefault: "random",
	Value:       []string{"static"},
}

var DmAAAPPostProcessPPWSTrustRenewalWaitIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_ws_trust",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_ws_trust_never_expire",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_allow_renewal",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
	},
}

var DmAAAPPostProcessPPWSTrustNewInstanceIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_ws_trust",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_ws_trust_never_expire",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_allow_renewal",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
	},
}

var DmAAAPPostProcessPPWSTrustNewKeyIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_ws_trust",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_ws_trust_never_expire",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_allow_renewal",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_ws_trust_new_instance",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
	},
}

var DmAAAPPostProcessPPWSTrustNeverExpireIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "pp_ws_trust",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var DmAAAPPostProcessPPICRXUserRealmIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "ppicrx_token",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var DmAAAPPostProcessPPSAMLProtocolIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "pp_saml_identity_provider",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var DmAAAPPostProcessPPSAMLResponseDestinationIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_saml_identity_provider",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "pp_saml_version",
			AttrType:    "String",
			AttrDefault: "2",
			Value:       []string{"2.0"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_saml_protocol",
			AttrType:    "String",
			AttrDefault: "assertion",
			Value:       []string{"assertion"},
		},
	},
}

var DmAAAPPostProcessPPResultWrapupIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "pp_saml_identity_provider",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var DmAAAPPostProcessPPSAMLAssertionTypeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "pp_saml_identity_provider",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var DmAAAPPostProcessPPSAMLAssertionTypeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "pp_saml_identity_provider",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var DmAAAPPostProcessPPSAMLSubjectConfirmIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "pp_saml_identity_provider",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var DmAAAPPostProcessPPSAMLNameIDIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "pp_saml_identity_provider",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var DmAAAPPostProcessPPSAMLNameIDFormatIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_saml_identity_provider",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_saml_name_id",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"false"},
		},
	},
}

var DmAAAPPostProcessPPSAMLRecipientIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "pp_saml_identity_provider",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var DmAAAPPostProcessPPSAMLAudienceIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "pp_saml_identity_provider",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var DmAAAPPostProcessPPSAMLOmitNotBeforeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "pp_saml_identity_provider",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var DmAAAPPostProcessPPOneTimeUseIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_saml_identity_provider",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_saml_version",
			AttrType:    "String",
			AttrDefault: "2",
			Value:       []string{"1.0"},
		},
	},
}

var DmAAAPPostProcessPPSAMLProxyIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_saml_identity_provider",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "pp_saml_version",
			AttrType:    "String",
			AttrDefault: "2",
			Value:       []string{"2.0"},
		},
	},
}

var DmAAAPPostProcessPPSAMLProxyAudienceIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_saml_identity_provider",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_saml_proxy",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "pp_saml_version",
			AttrType:    "String",
			AttrDefault: "2",
			Value:       []string{"2.0"},
		},
	},
}

var DmAAAPPostProcessPPSAMLProxyCountIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_saml_identity_provider",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_saml_proxy",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "pp_saml_version",
			AttrType:    "String",
			AttrDefault: "2",
			Value:       []string{"2.0"},
		},
	},
}

var DmAAAPPostProcessPPSAMLAuthzActionIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_saml_identity_provider",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "pp_saml_assertion_type",
			AttrType:    "DmSAMLStatementType",
			AttrDefault: "authentication+attribute",
			Value:       []string{"authorization"},
		},
	},
}

var DmAAAPPostProcessPPSAMLAttributesCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_saml_identity_provider",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_saml_assertion_type",
			AttrType:    "DmSAMLStatementType",
			AttrDefault: "authentication+attribute",
			Value:       []string{"attribute"},
		},
	},
}

var DmAAAPPostProcessPPSAMLAttributesIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPPostProcessPPLTPAInsertCookieIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_ltpa",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_use_ws_sec",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
	},
}

var DmAAAPPostProcessPPTAMHeaderCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "pp_tam_pac_propagate",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var DmAAAPPostProcessPPTAMHeaderIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "pp_tam_pac_propagate",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var DmAAAPPostProcessPPTAMHeaderSizeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "pp_tam_pac_propagate",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var DmAAAPPostProcessPPTAMHeaderSizeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "pp_tam_pac_propagate",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var DmAAAPPostProcessPPKerberosUseS4U2ProxyIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "ldap",
			AttrPath:    "../authenticate",
			Value:       []string{"kerberos"},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_kerberos_ticket",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"false"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_kerberos_spnego_token",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"false"},
				},
			},
		},
	},
}

var DmAAAPPostProcessPPCookieAttributesIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "pp_ltpa",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "pp_ltpa_insert_cookie",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
	},
}

var DmAAAPPostProcessPPKerberosUseS4U2SelfAndS4U2ProxyIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "ldap",
			AttrPath:    "../authenticate",
			Value:       []string{"kerberos"},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_kerberos_ticket",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"false"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_kerberos_spnego_token",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"false"},
				},
			},
		},
	},
}

var DmAAAPPostProcessPPKerberosClientSourceCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_kerberos_ticket",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_kerberos_spnego_token",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},
			},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "ldap",
			AttrPath:    "../authenticate",
			Value:       []string{"kerberos"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_kerberos_use_s4u2self_and_s4u2proxy",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
	},
}

var DmAAAPPostProcessPPKerberosClientSourceIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPPostProcessPPKerberosSelfCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_kerberos_ticket",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_kerberos_spnego_token",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},
			},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "ldap",
			AttrPath:    "../authenticate",
			Value:       []string{"kerberos"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_kerberos_use_s4u2self_and_s4u2proxy",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
	},
}

var DmAAAPPostProcessPPKerberosSelfIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPPostProcessPPKerberosSelfKeytabCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_kerberos_ticket",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_kerberos_spnego_token",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},
			},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "ldap",
			AttrPath:    "../authenticate",
			Value:       []string{"kerberos"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_kerberos_use_s4u2self_and_s4u2proxy",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
	},
}

var DmAAAPPostProcessPPKerberosSelfKeytabIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPPostProcessPPKerberosClientCustomURLCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_kerberos_ticket",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_kerberos_spnego_token",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},
			},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "ldap",
			AttrPath:    "../authenticate",
			Value:       []string{"kerberos"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_kerberos_use_s4u2self_and_s4u2proxy",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_kerberos_client_source",
			AttrType:    "String",
			AttrDefault: "mc-output",
			Value:       []string{"custom-url"},
		},
	},
}

var DmAAAPPostProcessPPKerberosClientCustomURLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPPostProcessPPKerberosClientCtxVarCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_kerberos_ticket",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_kerberos_spnego_token",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},
			},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "ldap",
			AttrPath:    "../authenticate",
			Value:       []string{"kerberos"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_kerberos_use_s4u2self_and_s4u2proxy",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_kerberos_client_source",
			AttrType:    "String",
			AttrDefault: "mc-output",
			Value:       []string{"ctx-var"},
		},
	},
}

var DmAAAPPostProcessPPKerberosClientCtxVarIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPPostProcessPPKerberosServerSourceCondVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_kerberos_ticket",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_kerberos_spnego_token",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
	},
}

var DmAAAPPostProcessPPKerberosServerSourceIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPPostProcessPPKerberosServerCustomURLCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_kerberos_ticket",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_kerberos_spnego_token",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},
			},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_kerberos_server_source",
			AttrType:    "String",
			AttrDefault: "as-is-string",
			Value:       []string{"custom-url"},
		},
	},
}

var DmAAAPPostProcessPPKerberosServerCustomURLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPPostProcessPPKerberosServerCtxVarCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_kerberos_ticket",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_kerberos_spnego_token",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},
			},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_kerberos_server_source",
			AttrType:    "String",
			AttrDefault: "as-is-string",
			Value:       []string{"ctx-var"},
		},
	},
}

var DmAAAPPostProcessPPKerberosServerCtxVarIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPPostProcessPPSSLClientConfigTypeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "pp_saml_send_slo",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var DmAAAPPostProcessPPSSLClientProfileIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_saml_send_slo",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "pp_ssl_client_config_type",
			AttrType:    "String",
			AttrDefault: "proxy",
			Value:       []string{"client"},
		},
	},
}

var DmAAAPPostProcessPPLTPAKeyFilePasswordAliasCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_ltpa_key_file_password",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{""},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_ltpa",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_ltpa_version",
					AttrType:    "String",
					AttrDefault: "LTPA2",
					Value:       []string{"LTPA", "LTPA1FIPS", "LTPA2", "LTPA2WAS7"},
				},
			},
		},
	},
}

var DmAAAPPostProcessPPLTPAKeyFilePasswordAliasIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "pp_ltpa",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-not-in-list",
					Attribute:   "pp_ltpa_version",
					AttrType:    "String",
					AttrDefault: "LTPA2",
					Value:       []string{"LTPA", "LTPA1FIPS", "LTPA2", "LTPA2WAS7"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "pp_ltpa_version",
					AttrType:    "String",
					AttrDefault: "LTPA2",
					Value:       []string{"LTPADomino"},
				},
			},
		},
	},
}

var DmAAAPPostProcessPPJWTGeneratorCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "pp_jwt",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var DmAAAPPostProcessPPJWTGeneratorIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPPostProcessObjectType = map[string]attr.Type{
	"pp_enabled":                                  types.BoolType,
	"pp_custom_url":                               types.StringType,
	"pp_saml_auth_assertion":                      types.BoolType,
	"pp_saml_server_name":                         types.StringType,
	"pp_saml_name_qualifier":                      types.StringType,
	"pp_kerberos_ticket":                          types.BoolType,
	"pp_kerberos_client":                          types.StringType,
	"pp_kerberos_client_password":                 types.StringType,
	"pp_kerberos_server":                          types.StringType,
	"pp_ws_trust":                                 types.BoolType,
	"pp_timestamp":                                types.BoolType,
	"pp_timestamp_expiry":                         types.Int64Type,
	"pp_allow_renewal":                            types.BoolType,
	"pp_saml_version":                             types.StringType,
	"pp_saml_send_slo":                            types.BoolType,
	"pp_saml_slo_endpoint":                        types.StringType,
	"pp_ws_username_token":                        types.BoolType,
	"pp_ws_username_token_password_type":          types.StringType,
	"pp_saml_validity":                            types.Int64Type,
	"pp_saml_skew":                                types.Int64Type,
	"pp_ws_username_token_include_pwd":            types.BoolType,
	"pp_ltpa":                                     types.BoolType,
	"pp_ltpa_version":                             types.StringType,
	"pp_ltpa_expiry":                              types.Int64Type,
	"pp_ltpa_key_file":                            types.StringType,
	"pp_ltpa_key_file_password":                   types.StringType,
	"pp_ltpa_stash_file":                          types.StringType,
	"pp_kerberos_spnego_token":                    types.BoolType,
	"pp_kerberos_bst_value_type":                  types.StringType,
	"pp_saml_use_ws_sec":                          types.BoolType,
	"pp_kerberos_client_keytab":                   types.StringType,
	"pp_use_ws_sec":                               types.BoolType,
	"pp_actor_role_id":                            types.StringType,
	"pp_ws_derived_key_username_token":            types.BoolType,
	"pp_ws_derived_key_username_token_iterations": types.Int64Type,
	"pp_ws_username_token_allow_replacement":      types.BoolType,
	"pp_hmac_signing_alg":                         types.StringType,
	"pp_signing_hash_alg":                         types.StringType,
	"pp_ws_trust_header":                          types.BoolType,
	"pp_ws_sc_key_source":                         types.StringType,
	"pp_shared_secret_key":                        types.StringType,
	"pp_ws_trust_renewal_wait":                    types.Int64Type,
	"pp_ws_trust_new_instance":                    types.BoolType,
	"pp_ws_trust_new_key":                         types.BoolType,
	"pp_ws_trust_never_expire":                    types.BoolType,
	"ppicrx_token":                                types.BoolType,
	"ppicrx_user_realm":                           types.StringType,
	"pp_saml_identity_provider":                   types.BoolType,
	"pp_saml_protocol":                            types.StringType,
	"pp_saml_response_destination":                types.StringType,
	"pp_result_wrapup":                            types.StringType,
	"pp_saml_assertion_type":                      types.ObjectType{AttrTypes: DmSAMLStatementTypeObjectType},
	"pp_saml_subject_confirm":                     types.StringType,
	"pp_saml_name_id":                             types.BoolType,
	"pp_saml_name_id_format":                      types.StringType,
	"pp_saml_recipient":                           types.StringType,
	"pp_saml_audience":                            types.StringType,
	"pp_saml_omit_not_before":                     types.BoolType,
	"pp_one_time_use":                             types.BoolType,
	"pp_saml_proxy":                               types.BoolType,
	"pp_saml_proxy_audience":                      types.StringType,
	"pp_saml_proxy_count":                         types.Int64Type,
	"pp_saml_authz_action":                        types.StringType,
	"pp_saml_attributes":                          types.StringType,
	"pp_ltpa_insert_cookie":                       types.BoolType,
	"pp_tam_pac_propagate":                        types.BoolType,
	"pp_tam_header":                               types.StringType,
	"pp_tam_header_size":                          types.Int64Type,
	"pp_kerberos_use_s4u2proxy":                   types.BoolType,
	"pp_cookie_attributes":                        types.StringType,
	"pp_kerberos_use_s4u2self_and_s4u2proxy":      types.BoolType,
	"pp_kerberos_client_source":                   types.StringType,
	"pp_kerberos_self":                            types.StringType,
	"pp_kerberos_self_keytab":                     types.StringType,
	"pp_kerberos_client_custom_url":               types.StringType,
	"pp_kerberos_client_ctx_var":                  types.StringType,
	"pp_kerberos_server_source":                   types.StringType,
	"pp_kerberos_server_custom_url":               types.StringType,
	"pp_kerberos_server_ctx_var":                  types.StringType,
	"pp_ssl_client_config_type":                   types.StringType,
	"pp_ssl_client_profile":                       types.StringType,
	"pp_ltpa_key_file_password_alias":             types.StringType,
	"pp_jwt":                                      types.BoolType,
	"pp_jwt_generator":                            types.StringType,
}
var DmAAAPPostProcessObjectDefault = map[string]attr.Value{
	"pp_enabled":                                  types.BoolValue(false),
	"pp_custom_url":                               types.StringNull(),
	"pp_saml_auth_assertion":                      types.BoolValue(false),
	"pp_saml_server_name":                         types.StringValue("XS"),
	"pp_saml_name_qualifier":                      types.StringNull(),
	"pp_kerberos_ticket":                          types.BoolValue(false),
	"pp_kerberos_client":                          types.StringNull(),
	"pp_kerberos_client_password":                 types.StringNull(),
	"pp_kerberos_server":                          types.StringNull(),
	"pp_ws_trust":                                 types.BoolValue(false),
	"pp_timestamp":                                types.BoolValue(true),
	"pp_timestamp_expiry":                         types.Int64Value(0),
	"pp_allow_renewal":                            types.BoolValue(false),
	"pp_saml_version":                             types.StringValue("2"),
	"pp_saml_send_slo":                            types.BoolValue(false),
	"pp_saml_slo_endpoint":                        types.StringNull(),
	"pp_ws_username_token":                        types.BoolValue(false),
	"pp_ws_username_token_password_type":          types.StringValue("Digest"),
	"pp_saml_validity":                            types.Int64Value(0),
	"pp_saml_skew":                                types.Int64Value(0),
	"pp_ws_username_token_include_pwd":            types.BoolValue(true),
	"pp_ltpa":                                     types.BoolValue(false),
	"pp_ltpa_version":                             types.StringValue("LTPA2"),
	"pp_ltpa_expiry":                              types.Int64Value(600),
	"pp_ltpa_key_file":                            types.StringNull(),
	"pp_ltpa_key_file_password":                   types.StringNull(),
	"pp_ltpa_stash_file":                          types.StringNull(),
	"pp_kerberos_spnego_token":                    types.BoolValue(false),
	"pp_kerberos_bst_value_type":                  types.StringValue("http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ"),
	"pp_saml_use_ws_sec":                          types.BoolValue(false),
	"pp_kerberos_client_keytab":                   types.StringNull(),
	"pp_use_ws_sec":                               types.BoolValue(false),
	"pp_actor_role_id":                            types.StringNull(),
	"pp_ws_derived_key_username_token":            types.BoolValue(false),
	"pp_ws_derived_key_username_token_iterations": types.Int64Value(1000),
	"pp_ws_username_token_allow_replacement":      types.BoolValue(false),
	"pp_hmac_signing_alg":                         types.StringValue("hmac-sha1"),
	"pp_signing_hash_alg":                         types.StringValue("sha1"),
	"pp_ws_trust_header":                          types.BoolValue(false),
	"pp_ws_sc_key_source":                         types.StringValue("random"),
	"pp_shared_secret_key":                        types.StringNull(),
	"pp_ws_trust_renewal_wait":                    types.Int64Value(0),
	"pp_ws_trust_new_instance":                    types.BoolValue(false),
	"pp_ws_trust_new_key":                         types.BoolValue(false),
	"pp_ws_trust_never_expire":                    types.BoolValue(false),
	"ppicrx_token":                                types.BoolValue(false),
	"ppicrx_user_realm":                           types.StringNull(),
	"pp_saml_identity_provider":                   types.BoolValue(false),
	"pp_saml_protocol":                            types.StringValue("assertion"),
	"pp_saml_response_destination":                types.StringNull(),
	"pp_result_wrapup":                            types.StringValue("wssec-replace"),
	"pp_saml_assertion_type":                      types.ObjectValueMust(DmSAMLStatementTypeObjectType, DmSAMLStatementTypeObjectDefault),
	"pp_saml_subject_confirm":                     types.StringValue("bearer"),
	"pp_saml_name_id":                             types.BoolValue(true),
	"pp_saml_name_id_format":                      types.StringNull(),
	"pp_saml_recipient":                           types.StringNull(),
	"pp_saml_audience":                            types.StringNull(),
	"pp_saml_omit_not_before":                     types.BoolValue(false),
	"pp_one_time_use":                             types.BoolValue(false),
	"pp_saml_proxy":                               types.BoolValue(false),
	"pp_saml_proxy_audience":                      types.StringNull(),
	"pp_saml_proxy_count":                         types.Int64Value(0),
	"pp_saml_authz_action":                        types.StringValue("AllHTTP"),
	"pp_saml_attributes":                          types.StringNull(),
	"pp_ltpa_insert_cookie":                       types.BoolValue(true),
	"pp_tam_pac_propagate":                        types.BoolValue(false),
	"pp_tam_header":                               types.StringValue("iv-creds"),
	"pp_tam_header_size":                          types.Int64Value(0),
	"pp_kerberos_use_s4u2proxy":                   types.BoolValue(false),
	"pp_cookie_attributes":                        types.StringNull(),
	"pp_kerberos_use_s4u2self_and_s4u2proxy":      types.BoolValue(false),
	"pp_kerberos_client_source":                   types.StringValue("mc-output"),
	"pp_kerberos_self":                            types.StringNull(),
	"pp_kerberos_self_keytab":                     types.StringNull(),
	"pp_kerberos_client_custom_url":               types.StringNull(),
	"pp_kerberos_client_ctx_var":                  types.StringNull(),
	"pp_kerberos_server_source":                   types.StringValue("as-is-string"),
	"pp_kerberos_server_custom_url":               types.StringNull(),
	"pp_kerberos_server_ctx_var":                  types.StringNull(),
	"pp_ssl_client_config_type":                   types.StringValue("proxy"),
	"pp_ssl_client_profile":                       types.StringNull(),
	"pp_ltpa_key_file_password_alias":             types.StringNull(),
	"pp_jwt":                                      types.BoolValue(false),
	"pp_jwt_generator":                            types.StringNull(),
}

func GetDmAAAPPostProcessDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmAAAPPostProcessDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"pp_enabled": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to run a custom stylesheet or GatewayScript file.", "custom-processing", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"pp_custom_url": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the custom file for the postprocessing activity.", "custom-url", "").AddRequiredWhen(DmAAAPPostProcessPPCustomURLCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPCustomURLIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_saml_auth_assertion": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate a SAML assertion that contains a SAML authentication statement for the authenticated user identity.", "saml-generate-assertion", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"pp_saml_server_name": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the value of the <tt>saml:Issuer</tt> of the generated SAML assertion or SAML SLO request. The default value is XS.</p><ul><li>If generating an SAML assertion, identifies the server that makes the assertion.</li><li>If sending an SLO request, identifies the issuer that sends the request.</li></ul>", "saml-server-name", "").AddDefaultValue("XS").AddNotValidWhen(DmAAAPPostProcessPPSAMLServerNameIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_saml_name_qualifier": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the value of the NameQualifier attribute of the NameIdentifier in the generated SAML assertion. Although the attribute is an optional attribute, some SAML implementations require that this attribute must be present.", "saml-name-qualifier", "").AddNotValidWhen(DmAAAPPostProcessPPSAMLNameQualifierIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_kerberos_ticket": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to include a WS-Security Kerberos AP-REQ BinarySecurityToken for the specified client and server principals in the WS-Security header. By default, token are not included.", "kerberos-include-token", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"pp_kerberos_client": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the client identity (cname of the Kerberos ticket) for the Kerberos client principal.", "kerberos-client-principal", "").AddRequiredWhen(DmAAAPPostProcessPPKerberosClientCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPKerberosClientIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_kerberos_client_password": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "", "").AddNotValidWhen(DmAAAPPostProcessPPKerberosClientPasswordIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_kerberos_server": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the server identity (sname of the Kerberos ticket) for the Kerberos server principal.", "kerberos-server", "").AddRequiredWhen(DmAAAPPostProcessPPKerberosServerCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPKerberosServerIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_ws_trust": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate the appropriate security token response for a valid WS-Trust SecurityContextToken request.", "ws-trust-generate-resp", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"pp_timestamp": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate a WS-Trust token time stamp for the security token response.", "ws-trust-add-timestamp", "").AddDefaultValue("true").AddNotValidWhen(DmAAAPPostProcessPPTimestampIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_timestamp_expiry": DataSourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the validity duration for the WS-Trust SCT in seconds to issue a new security context or to renew a context instance with new instance. Enter a value in the range 0 - 31622400. The default value is 0, which uses the value of the <tt>var://system/AAA/defaultexpiry</tt> variable if defined. If you did not define this variable, the value is 14400. If this setting is to renew a security context or instance, the value 0 means to use the old duration for the renewed cycle.", "ws-trust-timestamp-expiry", "").AddIntegerRange(0, 31622400).AddDefaultValue("0").AddNotValidWhen(DmAAAPPostProcessPPTimestampExpiryIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_allow_renewal": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether WS-Trust tokens can have their lifetime period reset without a new bootstrapping authentication event. If the WS-Trust request asks to renew the issued token, this setting is ignored.", "ws-trust-allow-renewal", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPPostProcessPPAllowRenewalIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_saml_version": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the protocol level of SAML messages. The version affects the identity extraction from the original message and the format of messages. The default value is 1.1.", "saml-version", "").AddStringEnum("2", "1.1", "1").AddDefaultValue("2").AddNotValidWhen(DmAAAPPostProcessPPSAMLVersionIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_saml_send_slo": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to send a SAML Logout (SLO) request to revoke the SAML Assertion token that is used for single-sign-on (SSO). The SLO is a request-response that the DataPower&#174; Gateway handles differently when it is working as a service provider (SP) or identity provider (IdP).</p><ul><li>When an SP, the DataPower Gateway sends an SLO request to the SAML SLO endpoint (IdP). On response, the DataPower Gateway processes the SLO response for its status.</li><li>When an IdP, the request to the DataPower Gateway contains the SLO request. Postprocessing validates against the SAML metadata file and sends the corresponding endpoint the SLO response.</li></ul>", "saml-send-slo", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"pp_saml_slo_endpoint": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the endpoint URL for SAML 2.0 Single Logout (SLO) messages. This endpoint is the authority that authenticated the assertion subject.", "saml-slo-endpoint", "").AddNotValidWhen(DmAAAPPostProcessPPSAMLSLOEndpointIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_ws_username_token": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to add a WS-Security UsernameToken. The username and password are taken from the output of the map credentials phase.", "wssec-add-user-name-token", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"pp_ws_username_token_password_type": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the type of password that the UsernameToken provides. By default, use the digest of the password as defined in the \"Web Services Security UsernameToken Profile 1.0\" specification.", "wssec-user-name-token-type", "").AddStringEnum("Text", "Digest").AddDefaultValue("Digest").AddNotValidWhen(DmAAAPPostProcessPPWSUsernameTokenPasswordTypeIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_saml_validity": DataSourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the validity duration of the SAML assertion in seconds. This value and the skew time are for fine control of the validity duration. The default value is 0.", "saml-validity", "").AddDefaultValue("0").AddNotValidWhen(DmAAAPPostProcessPPSAMLValidityIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_saml_skew": DataSourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the acceptable skew interval in seconds. The IdP and SP system clocks can have a skew time. When the SAML assertion is generated, the expiration takes the skew time setting into account. <ul><li>When <tt>NotBefore</tt> has the value of <tt>(CurrentTime - SkewTime)</tt> .</li><li>When <tt>NotOnOrAfter</tt> has the value of <tt>(CurrentTime + Validity + SkewTime)</tt> .</li></ul>", "saml-skew", "").AddDefaultValue("0").AddNotValidWhen(DmAAAPPostProcessPPSAMLSkewIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_ws_username_token_include_pwd": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the WS-Security UsernameToken must include the password. By default, the token must contain the password.", "wssec-user-name-token-contains-pwd", "").AddDefaultValue("true").AddNotValidWhen(DmAAAPPostProcessPPWSUsernameTokenIncludePwdIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_ltpa": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate an LTPA token.", "lpta-generate-token", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"pp_ltpa_version": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the LTPA token version to generate. By default, generates a WebSphere version 2 token.", "lpta-version", "").AddStringEnum("LTPA", "LTPA1FIPS", "LTPA2", "LTPA2WAS7", "LTPADomino").AddDefaultValue("LTPA2").AddNotValidWhen(DmAAAPPostProcessPPLTPAVersionIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_ltpa_expiry": DataSourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the lifetime of LTPA token in seconds. Enter a value in the range 1 - 628992000. The default value is 600.", "lpta-expiry", "").AddIntegerRange(1, 628992000).AddDefaultValue("600").AddNotValidWhen(DmAAAPPostProcessPPLTPAExpiryIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_ltpa_key_file": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the LTPA key file that secures the LTPA token. The LTPA key file contains the crypto material to create an LTPA token that can be consumed by WebSphere or Domino. <ul><li>For WebSphere tokens, you must export the LTPA key file from WebSphere. This file has portions encrypted by a password.</li><li>For Domino tokens, the key file should contain only the base 64-encoded Domino shared secret.</li></ul>", "lpta-key-file", "").AddRequiredWhen(DmAAAPPostProcessPPLTPAKeyFileCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPLTPAKeyFileIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_ltpa_key_file_password": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use the LTPA key file password alias.", "lpta-key-file-password", "").AddNotValidWhen(DmAAAPPostProcessPPLTPAKeyFilePasswordIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_ltpa_stash_file": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the file that contains the LTPA key file password.", "lpta-stash-file", "").AddNotValidWhen(DmAAAPPostProcessPPLTPAStashFileIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_kerberos_spnego_token": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate an SPNEGO token to insert into the HTTP WWW-Authenticate header.", "kerberos-generate-spnego", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"pp_kerberos_bst_value_type": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the value for the <tt>ValueType</tt> attribute of the WS-Security BinarySecurityToken. The Kerberos AP-REQ message contains the <tt>ValueType</tt> attribute. The default value is for WSS Kerberos Token Profile 1.1 (GSS).", "kerberos-value-type", "").AddStringEnum("http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ1510", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ1510", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ4120", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ4120", "http://docs.oasis-open.org/wss/2005/xx/oasis-2005xx-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ", "http://docs.oasis-open.org/wss/2005/xx/oasis-2005xx-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ", "http://www.docs.oasis-open.org/wss/2004/07/oasis-000000-wss-kerberos-token-profile-1.0#Kerberosv5_AP_REQ").AddDefaultValue("http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ").AddRequiredWhen(DmAAAPPostProcessPPKerberosBstValueTypeCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPKerberosBstValueTypeIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_saml_use_ws_sec": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify where to insert the SAML assertion. By default, the assertion is inserted as a child element of the SOAP header. When enabled, the assertion is inserted in a WS-Security-compliant header as defined by the WS-Security SAML token profile.", "saml-in-wssec", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPPostProcessPPSAMLUseWSSecIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_kerberos_client_keytab": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the Kerberos keytab that defines the keytab for the client. This keytab is required to authenticate the client to the KDC.", "kerberos-client-keytab", "crypto_kerberos_keytab").AddRequiredWhen(DmAAAPPostProcessPPKerberosClientKeytabCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPKerberosClientKeytabIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_use_ws_sec": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the token can be wrapped by the WS-Security <tt>wsse:Security</tt> header. This setting for the LTPA token. By default, the token cannot be wrapped by this header. When enabled, generate a WS-Security header that contains the token.", "wssec-header-wrap-token", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPPostProcessPPUseWSSecIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_actor_role_id": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the identifier for the SOAP 1.1 actor or SOAP 1.2 role for processing a WS-Security Security header. The DataPower Gateway works as that actor or role in consuming the input and generating the output for the next SOAP endpoint. This setting is meaningful when a SOAP message is being used for WS-Security 1.0 or 1.1. <table border=\"1\"><tr><td valign=\"left\">http://schemas.xmlsoap.org/soap/actor/next</td><td>Each receiver, including the intermediary and ultimate receiver, can process the Security header.</td></tr><tr><td valign=\"left\">http://www.w3.org/2003/05/soap-envelope/role/none</td><td>No one can process the Security header.</td></tr><tr><td valign=\"left\">http://www.w3.org/2003/05/soap-envelope/role/next</td><td>Each receiver, including the intermediary and ultimate receiver, can process the Security header.</td></tr><tr><td valign=\"left\">http://www.w3.org/2003/05/soap-envelope/role/ultimateReceiver</td><td>The ultimate receiver of the message can process the Security header. This value is the default value if such setting is not configured.</td></tr><tr><td valign=\"left\">&lt;blank or empty string></td><td>The empty string \"\" (without quotation marks) indicates that no actor or role identifier is configured. If no actor or role setting is configured, the ultimate receiver is assumed during message processing, and no actor or role attribute is added during the generation of the Security header. <p>This value does not generate an attribute with an empty value, which is the behavior as defined by the USE_MESSAGE_BASE_URI constant string. There cannot be more than one Security header that omits the actor or role identifier.</p></td></tr><tr><td valign=\"left\">USE_MESSAGE_BASE_URI</td><td>The constant value indicates that the actor or role identifier is the base URL of the message. If the SOAP message is transported over HTTP, the base URI is the Request-URI of the HTTP request.</td></tr><tr><td valign=\"left\">any other customized string</td><td>You can input any string to identify the actor or role of the Security header.</td></tr></table>", "wssec-actor-role-id", "").AddNotValidWhen(DmAAAPPostProcessPPActorRoleIDIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_ws_derived_key_username_token": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate a derived key from a password. By default, a derived key is not generated. When enabled, the process adds a WS-Security derived-key UsernameToken to the message and adds an HMAC signature with the derived-key. The username and password are taken from the output of the map credentials phase.", "wssec-use-derived-key", "").AddDefaultValue("false").AddRequiredWhen(DmAAAPPostProcessPPWSDerivedKeyUsernameTokenCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPWSDerivedKeyUsernameTokenIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_ws_derived_key_username_token_iterations": DataSourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the number of hashing cycles during the generation of a derived key from a password. The minimum value is 2. The default value is 1000.", "wssec-derived-key-hash-iter", "").AddIntegerRange(2, 65535).AddDefaultValue("1000").AddRequiredWhen(DmAAAPPostProcessPPWSDerivedKeyUsernameTokenIterationsCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPWSDerivedKeyUsernameTokenIterationsIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_ws_username_token_allow_replacement": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to retain the original token, not generate a new one, if the message already contains a UsernameToken. By default, the original otken is retained. When enabled, the generated token replaces any existing ones.", "wssec-replace-existing", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPPostProcessPPWSUsernameTokenAllowReplacementIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_hmac_signing_alg": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the HMAC algorithm to sign the token. This property is available to request a WS-Security UsernameToken in postprocessing and WS-Security Derived-Key UsernameToken is added to the message with an HMAC signature. The default value is hmac-sha1.", "hmac-signing-algorithm", "").AddStringEnum("hmac-sha1", "hmac-sha224", "hmac-sha256", "hmac-sha384", "hmac-sha512", "hmac-ripemd160", "hmac-md5").AddDefaultValue("hmac-sha1").AddRequiredWhen(DmAAAPPostProcessPPHMACSigningAlgCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPHMACSigningAlgIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_signing_hash_alg": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the algorithm for the message digest for the generation of a digital signature. This algorithm is for only the UsernameToken postprocessing method. The default value is sha1.", "message-digest-algorithm", "").AddStringEnum("sha1", "sha256", "sha512", "ripemd160", "sha224", "sha384", "md5").AddDefaultValue("sha1").AddRequiredWhen(DmAAAPPostProcessPPSigningHashAlgCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPSigningHashAlgIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_ws_trust_header": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to return the WS-Trust token as a SOAP header. By default, the token is put in the SOAP body. When enabled, return the token as a SOAP header by wrapping the <tt>wst:RequestedSecurityToken</tt> by a <tt>wst:IssuedToken</tt> .", "ws-trust-in-header", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPPostProcessPPWSTrustHeaderIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_ws_sc_key_source": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the source of the key. For WS-Trust postprocessing, the DataPower Gateway works as an on-box WS-Trust security token service that is backed by WS-SecureConversation. A symmetric shared secret key is needed to initialize the WS-SecureConversation SecurityContext. By default, a random key is generated.", "ws-trust-key-source", "").AddStringEnum("client-entropy", "in-kerberos", "in-encryptedkey", "static", "random").AddDefaultValue("random").AddNotValidWhen(DmAAAPPostProcessPPWSSCKeySourceIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_shared_secret_key": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the shared secret key as the WS-Trust key source.", "ws-trust-shared-key", "crypto_sskey").AddNotValidWhen(DmAAAPPostProcessPPSharedSecretKeyIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_ws_trust_renewal_wait": DataSourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration to allow the STS to keep an expired SecurityContext token in seconds. After a WS-Trust token expires, it can be removed from the STS and cannot be renewed. Therefore, the token must be renewed before expiry. Enter a value in the range of 0 - 2678400. The default value is 0. <p>The token is issued or renewed with a 1-hour wait time in the following situation.</p><ul><li>The WS-Trust request asks that the issued token can be renewed after expiration.</li><li>This setting has a value of 0.</li></ul>", "ws-trust-renewal-wait", "").AddIntegerRange(0, 2678400).AddDefaultValue("0").AddNotValidWhen(DmAAAPPostProcessPPWSTrustRenewalWaitIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_ws_trust_new_instance": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the STS renewal request issues a new instance for WS-Trust renewal. By default, the STS renewal request renews the existing instance. When enabled, the STS renewal request creates a new instance.", "ws-trust-new-instance", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPPostProcessPPWSTrustNewInstanceIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_ws_trust_new_key": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to update the context key for WS-Trust renewal.By default, the SCT renewal request uses the existing shared secret key. When enabled, the SCT renewal request does not use the existing shared secret key.", "ws-trust-new-key", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPPostProcessPPWSTrustNewKeyIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_ws_trust_never_expire": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the WS-Trust security context expires. By default, the security context expires. When enabled, the security context never expires.However, you can change the duration afterward with an explicit duration in seconds before expiry.", "ws-trust-never-expire", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPPostProcessPPWSTrustNeverExpireIgnoreVal.String()).String,
				Computed:            true,
			},
			"ppicrx_token": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate an Extended Identity Context Reference (ICRX) for z/OS identity propagation from the authenticated credentials. When generated, the WS-Security binary token with an ICRX token is inserted into the WS-Security header. You can use this token interoperability with the CICS Transaction Server for z/OS identity propagation support.", "generate-icrx", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"ppicrx_user_realm": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the realm of a user for ICRX identity propagation. The ICRX realm is defined in the SAF configuration. Generally, this value is the equivalent of the prefix for a DN in a user registry.", "icrx-user-realm", "").AddNotValidWhen(DmAAAPPostProcessPPICRXUserRealmIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_saml_identity_provider": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to generate a SAML assertion. The SAML assertion can contain an authentication statement, an authorization statement, an attribute statement, or any combination of these statements. The SAML attribute value can be a user LDAP Attribute value that can be retrieved in the following ways.</p><ul><li>Directly by the LDAP authentication or authorization method with the list of LDAP attribute names that are defined by user auxiliary LDAP attributes.</li><li>Indirectly with the <tt>var://context/ldap/auxiliary-attributes</tt> variable in a stylesheet or GatewayScript file. A call with <tt>dp:ldap-search</tt> to the user registry, and put the <tt>attribute-value</tt> elements of search result to the variable.</li></ul><p>To sign the SAML assertion, configure a WS-Security sign action or SAML enveloped sign action after the AAA action in the processing rule.</p>", "generate-saml-assertion", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"pp_saml_protocol": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the SAML protocol to wrap up the SAML assertion. By default, the SAML assertion can be put to WS-Security wrap-up later.", "saml-protocol", "").AddStringEnum("assertion", "response").AddDefaultValue("assertion").AddNotValidWhen(DmAAAPPostProcessPPSAMLProtocolIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_saml_response_destination": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the destination for a SAML response. This information can prevent malicious forwarding of requests to unintended recipients, which is a required protection by some protocol bindings. If it is present, the actual recipient must check that the URI reference identifies the location at which the message was received. If it does not check that the URI reference identifies the location, the request must be discarded. Some protocol bindings might require the use of this attribute.", "saml-response-destination", "").AddNotValidWhen(DmAAAPPostProcessPPSAMLResponseDestinationIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_result_wrapup": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the method to generate the result. When the DataPower Gateway is configured for SOAP or WS-Security processing, different output methods can be used. By default, generates the results to an existing WS-Security message and replaces the same token in the requesting message.", "result-wrapup", "").AddStringEnum("wssec-replace", "wssec-new", "wssec-inject", "soap-body", "none").AddDefaultValue("wssec-replace").AddNotValidWhen(DmAAAPPostProcessPPResultWrapupIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_saml_assertion_type": GetDmSAMLStatementTypeDataSourceSchema("Specify the supported SAML statement types. By default, supports both attributes and authentication statements.", "saml-assertion-type", ""),
			"pp_saml_subject_confirm": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the method that allows the destination system to confirm the subject of the SAML assertion. By default, the subject is bearer.", "saml-subject-confirm", "").AddStringEnum("bearer", "hok", "sv").AddDefaultValue("bearer").AddNotValidWhen(DmAAAPPostProcessPPSAMLSubjectConfirmIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_saml_name_id": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the SAML Subject element contains the name identifier. By default, the SAML subject contains the name identifier. When disabled, the SAML subject does not contain the name identifier. Use this value if the subject confirmation method is holder-of-key because the key represent the same entity as the subject.", "saml-nid", "").AddDefaultValue("true").AddNotValidWhen(DmAAAPPostProcessPPSAMLNameIDIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_saml_name_id_format": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URI reference that represents the classification of string-based identifier information. Any standard or arbitrary URI is allowed. If the value is an empty string, the DataPower Gateway attempts to determine the value from the AAA context. Some SAML protocols require a specified value, such as <tt>urn:oasis:names:tc:SAML:2.0:nameid-format:entity</tt> or <tt>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</tt> .", "saml-nid-format", "").AddNotValidWhen(DmAAAPPostProcessPPSAMLNameIDFormatIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_saml_recipient": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify a URI that identifies the entity or location that an attesting entity can present the assertion to. Any standard or arbitrary URI is allowed. If the value is an empty string, the optional attribute is not generated. This setting is applicable for only SAML 2.0.", "saml-recipient", "").AddNotValidWhen(DmAAAPPostProcessPPSAMLRecipientIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_saml_audience": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify URI references that identify an intended audience. Enter any number of the audience URIs to process the generated SAML assertion. If the value is an empty string, the SAML audience is not restricted. If there is more than one audience URI, use a + delimiter between URIs. In this case, you must convert any URI that contains the + characters to \\+.", "saml-audience", "").AddNotValidWhen(DmAAAPPostProcessPPSAMLAudienceIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_saml_omit_not_before": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("whether to omit the <tt>NotBefore</tt> attribute in the SAML assertion. When omitted, the assertion is considered valid even before the time it was issued. By default, the <tt>NotBefore</tt> attribute is not omitted. When enabled, the <tt>NotBefore</tt> attribute in the SAML assertion is omitted. This behavior might be required to respond to an <tt>AuthnRequest</tt> .", "saml-omit-notbefore", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPPostProcessPPSAMLOmitNotBeforeIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_one_time_use": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the destination system or relying party should cache the generated token. The generated token might contain the property for this characteristic, which is especially practical for SAML assertions. By default, the destination system can cache the generated token. When enabled, he destination system should not cache the generated token.", "one-time-use", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPPostProcessPPOneTimeUseIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_saml_proxy": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to allow SAML proxy restriction. The generated SAML assertion provides limitations that the asserting party imposes on relying parties that want to act as asserting parties.</p><ul><li>A relying party that acts as an asserting party can issue subsequent assertions that are based on the information in the original assertion.</li><li>The relying party cannot issue an assertion that violates these restrictions.</li></ul><p>By default, proxy restrictions are not allowd. When enabled, proxy restrictions are allows.</p>", "saml-proxy", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPPostProcessPPSAMLProxyIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_saml_proxy_audience": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the set of audiences (proxy) to whom the asserting party permits new assertions to be issued based on this assertion. If the value is an empty string, the audience for the <tt>ProxyRestriction</tt> is not issued with this SAML assertion. If there is more than one audience URI, use a + delimiter between URIs. In this case, you must convert any URI that contains the + characters to \\+.", "saml-proxy-audience", "").AddNotValidWhen(DmAAAPPostProcessPPSAMLProxyAudienceIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_saml_proxy_count": DataSourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum number of indirections that the asserting party permits between this assertion and an assertion that was issued. Enter a value in the range 0 - 65535. The default value is 0. A value of 0 indicates that a relying party must not issue an assertion to another relying party based on this assertion. If greater than zero, any assertion that is issued must itself contain a <tt>ProxyRestriction</tt> element with a <tt>Count</tt> value of at most one less than this value.", "saml-proxy-count", "").AddIntegerRange(0, 65535).AddDefaultValue("0").AddNotValidWhen(DmAAAPPostProcessPPSAMLProxyCountIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_saml_authz_action": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the standard action that the subject can take on the resource. The SAML specification defines the list of action identifiers with corresponding namespace URIs. By default, all HTTP operations, where <tt>urn:oasis:names:tc:SAML:1.0:action:ghpp</tt> is the namespace URI.", "saml-authz-action", "").AddStringEnum("AllHTTP", "POST", "GET", "PUT", "HEAD", "General", "Read", "Write", "Execute", "Delete", "Control", "NegatedRead", "NegatedWrite", "NegatedExecute", "NegatedDelete", "NegatedControl").AddDefaultValue("AllHTTP").AddNotValidWhen(DmAAAPPostProcessPPSAMLAuthzActionIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_saml_attributes": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of an existing SAML attributes. The SAML attributes define the information to put in the SAML assertion to generate the attribute statement. Each SAML attribute requires the name, format or namespace, and value. The value can be from a DataPower variable.", "saml-attributes", "saml_attributes").AddRequiredWhen(DmAAAPPostProcessPPSAMLAttributesCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPSAMLAttributesIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_ltpa_insert_cookie": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to insert a <tt>Set-Cookie</tt> header in the response that contains the LTPA token. This setting is for generating LTPA tokens that are not wrapped in the WS-Security <tt>wsse:Security</tt> header. By default, inserts a Set-Cookie header in the response. When disabled, does not insert a Set-Cookie header in the response.", "ltpa-insert-cookie", "").AddDefaultValue("true").AddNotValidWhen(DmAAAPPostProcessPPLTPAInsertCookieIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_tam_pac_propagate": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to add the Access Manager privilege attribute certificate (PAC) token to an HTTP header. The PAC token was returned from the previous authentication or authorization phase. By default, does not add the PAC token. When enabled, adds the PAC token.", "propagate-tam-pac", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"pp_tam_header": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the HTTP header to store the token in. The default value is iv_creds, which is HTTP header that WebSEAL uses to write headers.", "tam-header", "").AddDefaultValue("iv-creds").AddRequiredWhen(DmAAAPPostProcessPPTAMHeaderCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPTAMHeaderIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_tam_header_size": DataSourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum size in bytes of HTTP headers. A value of 0 disables this function. If the value is nonzero, the PAC token is split across multiple headers of the specified length. The default value is 0.", "tam-header-size", "").AddDefaultValue("0").AddRequiredWhen(DmAAAPPostProcessPPTAMHeaderSizeCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPTAMHeaderSizeIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_kerberos_use_s4u2proxy": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to use constrained delegation, namely S4U2Proxy, when a WS-Security Kerberos AP-REQ token or a Kerberos SPNEGO token is generated. By default, does not use constrained delegation. When enabled, uses constrained delegation.", "kerberos-use-s4u2proxy", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPPostProcessPPKerberosUseS4U2ProxyIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_cookie_attributes": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the cookie attribute policy to include standard or custom attributes in the cookie. The response message that contains a <tt>Set-Cookie</tt> header is updated with the attributes defined in this policy.", "cookie-attributes", "cookie_attribute_policy").AddNotValidWhen(DmAAAPPostProcessPPCookieAttributesIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_kerberos_use_s4u2self_and_s4u2proxy": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to use protocol transition, namely S4U2Self, and then use constrained delegation, namely S4U2Proxy.</p><ul><li>Use S4U2Self to convert a non-Kerberos token to a Kerberos token to the DataPower Gateway itself.</li><li>Use S4U2Proxy to generate a WS-Security Kerberos AP-REQ token or a Kerberos SPNEGO token.</li></ul><p>By default, does not use protocol transition and constrained delegation. When enabled, uses protocol transition and constrained delegation.</p>", "kerberos-use-s4u2self", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPPostProcessPPKerberosUseS4U2SelfAndS4U2ProxyIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_kerberos_client_source": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify where to get the principal name of the Kerberos client. By default, uses the output of credential mapping. The client principal is based on the authenticated identity, which is followed by the corresponding realm name. For example, if the authenticated user is <tt>alice</tt> , the client principal name can be <tt>HTTP/alice.datapower.com@DATAPOWER.COM</tt> . The client principal must be present in the KDC for S4U2Self to work.", "kerberos-client-source", "").AddStringEnum("mc-output", "custom-url", "ctx-var").AddDefaultValue("mc-output").AddRequiredWhen(DmAAAPPostProcessPPKerberosClientSourceCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPKerberosClientSourceIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_kerberos_self": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the principal name of the DataPower Gateway.", "kerberos-self-principal", "").AddRequiredWhen(DmAAAPPostProcessPPKerberosSelfCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPKerberosSelfIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_kerberos_self_keytab": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the Kerberos keytab that defines the keytab for the DataPower Gateway. This keytab is required to authenticate the DataPower Gateway to the KDC.", "kerberos-self-keytab", "crypto_kerberos_keytab").AddRequiredWhen(DmAAAPPostProcessPPKerberosSelfKeytabCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPKerberosSelfKeytabIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_kerberos_client_custom_url": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the location of the stylesheet or GatewayScript file. This file returns the client principal name within the <tt>kerberos-client-principal</tt> element. This file gets the following input.</p><ul><li>The output of all the steps that are executed in this AAA action.</li><li>The incoming request message.</li></ul>", "kerberos-client-custom-url", "").AddRequiredWhen(DmAAAPPostProcessPPKerberosClientCustomURLCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPKerberosClientCustomURLIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_kerberos_client_ctx_var": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the context variable. The value of this context variable is used as the Kerberos client principal. This context variable must be specified in the <tt>var://context/name</tt> format. For example, <tt>var://context/AAA/krb-client-princ</tt> . You can use the set variable action to set this variable in the processing rule before the AAA action.", "kerberos-client-ctx-var", "").AddRequiredWhen(DmAAAPPostProcessPPKerberosClientCtxVarCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPKerberosClientCtxVarIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_kerberos_server_source": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify where to get the principal name of the Kerberos server. By default, the server principal name is the value that is specified by the Kerberos server principal property. Ensure that the server principal is in the correct format. For example, <tt>HTTP/was-backend.datapower.com@DATAPOWER.COM</tt> .", "kerberos-server-source", "").AddStringEnum("as-is-string", "custom-url", "ctx-var").AddDefaultValue("as-is-string").AddRequiredWhen(DmAAAPPostProcessPPKerberosServerSourceCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPKerberosServerSourceIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_kerberos_server_custom_url": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the location of the stylesheet or GatewayScript file. This file returns the client principal name within the <tt>kerberos-server-principal</tt> element.</p><p>When constrained delegation is not used, this file gets the following input.</p><ul><li>The output of all phases that this AAA action processes.</li><li>The incoming request message.</li></ul><p>When constrained delegation is used, this file gets the following input.</p><ul><li>The output of only the identity extraction phase.</li><li>The incoming request message.</li></ul>", "kerberos-server-custom-url", "").AddRequiredWhen(DmAAAPPostProcessPPKerberosServerCustomURLCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPKerberosServerCustomURLIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_kerberos_server_ctx_var": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the context variable. The value of this context variable is used as the Kerberos server principal. This context variable must be specified in the <tt>var://context/name format</tt> . For example, <tt>var:///context/AAA/krb-server-princ</tt> . You can use the set variable action to set this variable in the processing rule before the AAA action.", "kerberos-server-ctx-var", "").AddRequiredWhen(DmAAAPPostProcessPPKerberosServerCtxVarCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPKerberosServerCtxVarIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_ssl_client_config_type": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS client type to secure connections.", "ssl-client-type", "").AddStringEnum("proxy", "client").AddDefaultValue("proxy").AddNotValidWhen(DmAAAPPostProcessPPSSLClientConfigTypeIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_ssl_client_profile": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS client profile to secure connections.", "ssl-client", "ssl_client_profile").AddNotValidWhen(DmAAAPPostProcessPPSSLClientProfileIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_ltpa_key_file_password_alias": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the the alias for password of the LTPA key file.", "ltpa-key-file-password-alias", "password_alias").AddRequiredWhen(DmAAAPPostProcessPPLTPAKeyFilePasswordAliasCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPLTPAKeyFilePasswordAliasIgnoreVal.String()).String,
				Computed:            true,
			},
			"pp_jwt": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate a JWT token.", "jwt", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"pp_jwt_generator": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the JWT generator.", "generate-jwt", "aaa_jwt_generator").AddRequiredWhen(DmAAAPPostProcessPPJWTGeneratorCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPJWTGeneratorIgnoreVal.String()).String,
				Computed:            true,
			},
		},
	}
	DmAAAPPostProcessDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmAAAPPostProcessDataSourceSchema
}
func GetDmAAAPPostProcessResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the custom file for the postprocessing activity.", "custom-url", "").AddRequiredWhen(DmAAAPPostProcessPPCustomURLCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPCustomURLIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPPostProcessPPCustomURLCondVal, DmAAAPPostProcessPPCustomURLIgnoreVal, false),
				},
			},
			"pp_saml_auth_assertion": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate a SAML assertion that contains a SAML authentication statement for the authenticated user identity.", "saml-generate-assertion", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"pp_saml_server_name": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the value of the <tt>saml:Issuer</tt> of the generated SAML assertion or SAML SLO request. The default value is XS.</p><ul><li>If generating an SAML assertion, identifies the server that makes the assertion.</li><li>If sending an SLO request, identifies the issuer that sends the request.</li></ul>", "saml-server-name", "").AddDefaultValue("XS").AddNotValidWhen(DmAAAPPostProcessPPSAMLServerNameIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("XS"),
			},
			"pp_saml_name_qualifier": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the value of the NameQualifier attribute of the NameIdentifier in the generated SAML assertion. Although the attribute is an optional attribute, some SAML implementations require that this attribute must be present.", "saml-name-qualifier", "").AddNotValidWhen(DmAAAPPostProcessPPSAMLNameQualifierIgnoreVal.String()).String,
				Optional:            true,
			},
			"pp_kerberos_ticket": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to include a WS-Security Kerberos AP-REQ BinarySecurityToken for the specified client and server principals in the WS-Security header. By default, token are not included.", "kerberos-include-token", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"pp_kerberos_client": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the client identity (cname of the Kerberos ticket) for the Kerberos client principal.", "kerberos-client-principal", "").AddRequiredWhen(DmAAAPPostProcessPPKerberosClientCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPKerberosClientIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPPostProcessPPKerberosClientCondVal, DmAAAPPostProcessPPKerberosClientIgnoreVal, false),
				},
			},
			"pp_kerberos_client_password": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "", "").AddNotValidWhen(DmAAAPPostProcessPPKerberosClientPasswordIgnoreVal.String()).String,
				Optional:            true,
			},
			"pp_kerberos_server": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the server identity (sname of the Kerberos ticket) for the Kerberos server principal.", "kerberos-server", "").AddRequiredWhen(DmAAAPPostProcessPPKerberosServerCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPKerberosServerIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPPostProcessPPKerberosServerCondVal, DmAAAPPostProcessPPKerberosServerIgnoreVal, false),
				},
			},
			"pp_ws_trust": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate the appropriate security token response for a valid WS-Trust SecurityContextToken request.", "ws-trust-generate-resp", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"pp_timestamp": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate a WS-Trust token time stamp for the security token response.", "ws-trust-add-timestamp", "").AddDefaultValue("true").AddNotValidWhen(DmAAAPPostProcessPPTimestampIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"pp_timestamp_expiry": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the validity duration for the WS-Trust SCT in seconds to issue a new security context or to renew a context instance with new instance. Enter a value in the range 0 - 31622400. The default value is 0, which uses the value of the <tt>var://system/AAA/defaultexpiry</tt> variable if defined. If you did not define this variable, the value is 14400. If this setting is to renew a security context or instance, the value 0 means to use the old duration for the renewed cycle.", "ws-trust-timestamp-expiry", "").AddIntegerRange(0, 31622400).AddDefaultValue("0").AddNotValidWhen(DmAAAPPostProcessPPTimestampExpiryIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 31622400),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, DmAAAPPostProcessPPTimestampExpiryIgnoreVal, true),
				},
				Default: int64default.StaticInt64(0),
			},
			"pp_allow_renewal": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether WS-Trust tokens can have their lifetime period reset without a new bootstrapping authentication event. If the WS-Trust request asks to renew the issued token, this setting is ignored.", "ws-trust-allow-renewal", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPPostProcessPPAllowRenewalIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"pp_saml_version": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the protocol level of SAML messages. The version affects the identity extraction from the original message and the format of messages. The default value is 1.1.", "saml-version", "").AddStringEnum("2", "1.1", "1").AddDefaultValue("2").AddNotValidWhen(DmAAAPPostProcessPPSAMLVersionIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("2", "1.1", "1"),
					validators.ConditionalRequiredString(validators.Evaluation{}, DmAAAPPostProcessPPSAMLVersionIgnoreVal, true),
				},
				Default: stringdefault.StaticString("2"),
			},
			"pp_saml_send_slo": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to send a SAML Logout (SLO) request to revoke the SAML Assertion token that is used for single-sign-on (SSO). The SLO is a request-response that the DataPower&#174; Gateway handles differently when it is working as a service provider (SP) or identity provider (IdP).</p><ul><li>When an SP, the DataPower Gateway sends an SLO request to the SAML SLO endpoint (IdP). On response, the DataPower Gateway processes the SLO response for its status.</li><li>When an IdP, the request to the DataPower Gateway contains the SLO request. Postprocessing validates against the SAML metadata file and sends the corresponding endpoint the SLO response.</li></ul>", "saml-send-slo", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"pp_saml_slo_endpoint": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the endpoint URL for SAML 2.0 Single Logout (SLO) messages. This endpoint is the authority that authenticated the assertion subject.", "saml-slo-endpoint", "").AddNotValidWhen(DmAAAPPostProcessPPSAMLSLOEndpointIgnoreVal.String()).String,
				Optional:            true,
			},
			"pp_ws_username_token": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to add a WS-Security UsernameToken. The username and password are taken from the output of the map credentials phase.", "wssec-add-user-name-token", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"pp_ws_username_token_password_type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the type of password that the UsernameToken provides. By default, use the digest of the password as defined in the \"Web Services Security UsernameToken Profile 1.0\" specification.", "wssec-user-name-token-type", "").AddStringEnum("Text", "Digest").AddDefaultValue("Digest").AddNotValidWhen(DmAAAPPostProcessPPWSUsernameTokenPasswordTypeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Text", "Digest"),
					validators.ConditionalRequiredString(validators.Evaluation{}, DmAAAPPostProcessPPWSUsernameTokenPasswordTypeIgnoreVal, true),
				},
				Default: stringdefault.StaticString("Digest"),
			},
			"pp_saml_validity": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the validity duration of the SAML assertion in seconds. This value and the skew time are for fine control of the validity duration. The default value is 0.", "saml-validity", "").AddDefaultValue("0").AddNotValidWhen(DmAAAPPostProcessPPSAMLValidityIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             int64default.StaticInt64(0),
			},
			"pp_saml_skew": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the acceptable skew interval in seconds. The IdP and SP system clocks can have a skew time. When the SAML assertion is generated, the expiration takes the skew time setting into account. <ul><li>When <tt>NotBefore</tt> has the value of <tt>(CurrentTime - SkewTime)</tt> .</li><li>When <tt>NotOnOrAfter</tt> has the value of <tt>(CurrentTime + Validity + SkewTime)</tt> .</li></ul>", "saml-skew", "").AddDefaultValue("0").AddNotValidWhen(DmAAAPPostProcessPPSAMLSkewIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             int64default.StaticInt64(0),
			},
			"pp_ws_username_token_include_pwd": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the WS-Security UsernameToken must include the password. By default, the token must contain the password.", "wssec-user-name-token-contains-pwd", "").AddDefaultValue("true").AddNotValidWhen(DmAAAPPostProcessPPWSUsernameTokenIncludePwdIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"pp_ltpa": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate an LTPA token.", "lpta-generate-token", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"pp_ltpa_version": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the LTPA token version to generate. By default, generates a WebSphere version 2 token.", "lpta-version", "").AddStringEnum("LTPA", "LTPA1FIPS", "LTPA2", "LTPA2WAS7", "LTPADomino").AddDefaultValue("LTPA2").AddNotValidWhen(DmAAAPPostProcessPPLTPAVersionIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("LTPA", "LTPA1FIPS", "LTPA2", "LTPA2WAS7", "LTPADomino"),
					validators.ConditionalRequiredString(validators.Evaluation{}, DmAAAPPostProcessPPLTPAVersionIgnoreVal, true),
				},
				Default: stringdefault.StaticString("LTPA2"),
			},
			"pp_ltpa_expiry": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the lifetime of LTPA token in seconds. Enter a value in the range 1 - 628992000. The default value is 600.", "lpta-expiry", "").AddIntegerRange(1, 628992000).AddDefaultValue("600").AddNotValidWhen(DmAAAPPostProcessPPLTPAExpiryIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 628992000),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, DmAAAPPostProcessPPLTPAExpiryIgnoreVal, true),
				},
				Default: int64default.StaticInt64(600),
			},
			"pp_ltpa_key_file": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the LTPA key file that secures the LTPA token. The LTPA key file contains the crypto material to create an LTPA token that can be consumed by WebSphere or Domino. <ul><li>For WebSphere tokens, you must export the LTPA key file from WebSphere. This file has portions encrypted by a password.</li><li>For Domino tokens, the key file should contain only the base 64-encoded Domino shared secret.</li></ul>", "lpta-key-file", "").AddRequiredWhen(DmAAAPPostProcessPPLTPAKeyFileCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPLTPAKeyFileIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPPostProcessPPLTPAKeyFileCondVal, DmAAAPPostProcessPPLTPAKeyFileIgnoreVal, false),
				},
			},
			"pp_ltpa_key_file_password": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use the LTPA key file password alias.", "lpta-key-file-password", "").AddNotValidWhen(DmAAAPPostProcessPPLTPAKeyFilePasswordIgnoreVal.String()).String,
				Optional:            true,
			},
			"pp_ltpa_stash_file": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the file that contains the LTPA key file password.", "lpta-stash-file", "").AddNotValidWhen(DmAAAPPostProcessPPLTPAStashFileIgnoreVal.String()).String,
				Optional:            true,
			},
			"pp_kerberos_spnego_token": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate an SPNEGO token to insert into the HTTP WWW-Authenticate header.", "kerberos-generate-spnego", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"pp_kerberos_bst_value_type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the value for the <tt>ValueType</tt> attribute of the WS-Security BinarySecurityToken. The Kerberos AP-REQ message contains the <tt>ValueType</tt> attribute. The default value is for WSS Kerberos Token Profile 1.1 (GSS).", "kerberos-value-type", "").AddStringEnum("http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ1510", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ1510", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ4120", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ4120", "http://docs.oasis-open.org/wss/2005/xx/oasis-2005xx-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ", "http://docs.oasis-open.org/wss/2005/xx/oasis-2005xx-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ", "http://www.docs.oasis-open.org/wss/2004/07/oasis-000000-wss-kerberos-token-profile-1.0#Kerberosv5_AP_REQ").AddDefaultValue("http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ").AddRequiredWhen(DmAAAPPostProcessPPKerberosBstValueTypeCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPKerberosBstValueTypeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ1510", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ1510", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ4120", "http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ4120", "http://docs.oasis-open.org/wss/2005/xx/oasis-2005xx-wss-kerberos-token-profile-1.1#Kerberosv5_AP_REQ", "http://docs.oasis-open.org/wss/2005/xx/oasis-2005xx-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ", "http://www.docs.oasis-open.org/wss/2004/07/oasis-000000-wss-kerberos-token-profile-1.0#Kerberosv5_AP_REQ"),
					validators.ConditionalRequiredString(DmAAAPPostProcessPPKerberosBstValueTypeCondVal, DmAAAPPostProcessPPKerberosBstValueTypeIgnoreVal, true),
				},
				Default: stringdefault.StaticString("http://docs.oasis-open.org/wss/oasis-wss-kerberos-token-profile-1.1#GSS_Kerberosv5_AP_REQ"),
			},
			"pp_saml_use_ws_sec": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify where to insert the SAML assertion. By default, the assertion is inserted as a child element of the SOAP header. When enabled, the assertion is inserted in a WS-Security-compliant header as defined by the WS-Security SAML token profile.", "saml-in-wssec", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPPostProcessPPSAMLUseWSSecIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"pp_kerberos_client_keytab": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the Kerberos keytab that defines the keytab for the client. This keytab is required to authenticate the client to the KDC.", "kerberos-client-keytab", "crypto_kerberos_keytab").AddRequiredWhen(DmAAAPPostProcessPPKerberosClientKeytabCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPKerberosClientKeytabIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPPostProcessPPKerberosClientKeytabCondVal, DmAAAPPostProcessPPKerberosClientKeytabIgnoreVal, false),
				},
			},
			"pp_use_ws_sec": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the token can be wrapped by the WS-Security <tt>wsse:Security</tt> header. This setting for the LTPA token. By default, the token cannot be wrapped by this header. When enabled, generate a WS-Security header that contains the token.", "wssec-header-wrap-token", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPPostProcessPPUseWSSecIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"pp_actor_role_id": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the identifier for the SOAP 1.1 actor or SOAP 1.2 role for processing a WS-Security Security header. The DataPower Gateway works as that actor or role in consuming the input and generating the output for the next SOAP endpoint. This setting is meaningful when a SOAP message is being used for WS-Security 1.0 or 1.1. <table border=\"1\"><tr><td valign=\"left\">http://schemas.xmlsoap.org/soap/actor/next</td><td>Each receiver, including the intermediary and ultimate receiver, can process the Security header.</td></tr><tr><td valign=\"left\">http://www.w3.org/2003/05/soap-envelope/role/none</td><td>No one can process the Security header.</td></tr><tr><td valign=\"left\">http://www.w3.org/2003/05/soap-envelope/role/next</td><td>Each receiver, including the intermediary and ultimate receiver, can process the Security header.</td></tr><tr><td valign=\"left\">http://www.w3.org/2003/05/soap-envelope/role/ultimateReceiver</td><td>The ultimate receiver of the message can process the Security header. This value is the default value if such setting is not configured.</td></tr><tr><td valign=\"left\">&lt;blank or empty string></td><td>The empty string \"\" (without quotation marks) indicates that no actor or role identifier is configured. If no actor or role setting is configured, the ultimate receiver is assumed during message processing, and no actor or role attribute is added during the generation of the Security header. <p>This value does not generate an attribute with an empty value, which is the behavior as defined by the USE_MESSAGE_BASE_URI constant string. There cannot be more than one Security header that omits the actor or role identifier.</p></td></tr><tr><td valign=\"left\">USE_MESSAGE_BASE_URI</td><td>The constant value indicates that the actor or role identifier is the base URL of the message. If the SOAP message is transported over HTTP, the base URI is the Request-URI of the HTTP request.</td></tr><tr><td valign=\"left\">any other customized string</td><td>You can input any string to identify the actor or role of the Security header.</td></tr></table>", "wssec-actor-role-id", "").AddNotValidWhen(DmAAAPPostProcessPPActorRoleIDIgnoreVal.String()).String,
				Optional:            true,
			},
			"pp_ws_derived_key_username_token": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate a derived key from a password. By default, a derived key is not generated. When enabled, the process adds a WS-Security derived-key UsernameToken to the message and adds an HMAC signature with the derived-key. The username and password are taken from the output of the map credentials phase.", "wssec-use-derived-key", "").AddDefaultValue("false").AddRequiredWhen(DmAAAPPostProcessPPWSDerivedKeyUsernameTokenCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPWSDerivedKeyUsernameTokenIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"pp_ws_derived_key_username_token_iterations": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the number of hashing cycles during the generation of a derived key from a password. The minimum value is 2. The default value is 1000.", "wssec-derived-key-hash-iter", "").AddIntegerRange(2, 65535).AddDefaultValue("1000").AddRequiredWhen(DmAAAPPostProcessPPWSDerivedKeyUsernameTokenIterationsCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPWSDerivedKeyUsernameTokenIterationsIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(2, 65535),
					validators.ConditionalRequiredInt64(DmAAAPPostProcessPPWSDerivedKeyUsernameTokenIterationsCondVal, DmAAAPPostProcessPPWSDerivedKeyUsernameTokenIterationsIgnoreVal, true),
				},
				Default: int64default.StaticInt64(1000),
			},
			"pp_ws_username_token_allow_replacement": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to retain the original token, not generate a new one, if the message already contains a UsernameToken. By default, the original otken is retained. When enabled, the generated token replaces any existing ones.", "wssec-replace-existing", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPPostProcessPPWSUsernameTokenAllowReplacementIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"pp_hmac_signing_alg": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the HMAC algorithm to sign the token. This property is available to request a WS-Security UsernameToken in postprocessing and WS-Security Derived-Key UsernameToken is added to the message with an HMAC signature. The default value is hmac-sha1.", "hmac-signing-algorithm", "").AddStringEnum("hmac-sha1", "hmac-sha224", "hmac-sha256", "hmac-sha384", "hmac-sha512", "hmac-ripemd160", "hmac-md5").AddDefaultValue("hmac-sha1").AddRequiredWhen(DmAAAPPostProcessPPHMACSigningAlgCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPHMACSigningAlgIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("hmac-sha1", "hmac-sha224", "hmac-sha256", "hmac-sha384", "hmac-sha512", "hmac-ripemd160", "hmac-md5"),
					validators.ConditionalRequiredString(DmAAAPPostProcessPPHMACSigningAlgCondVal, DmAAAPPostProcessPPHMACSigningAlgIgnoreVal, true),
				},
				Default: stringdefault.StaticString("hmac-sha1"),
			},
			"pp_signing_hash_alg": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the algorithm for the message digest for the generation of a digital signature. This algorithm is for only the UsernameToken postprocessing method. The default value is sha1.", "message-digest-algorithm", "").AddStringEnum("sha1", "sha256", "sha512", "ripemd160", "sha224", "sha384", "md5").AddDefaultValue("sha1").AddRequiredWhen(DmAAAPPostProcessPPSigningHashAlgCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPSigningHashAlgIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("sha1", "sha256", "sha512", "ripemd160", "sha224", "sha384", "md5"),
					validators.ConditionalRequiredString(DmAAAPPostProcessPPSigningHashAlgCondVal, DmAAAPPostProcessPPSigningHashAlgIgnoreVal, true),
				},
				Default: stringdefault.StaticString("sha1"),
			},
			"pp_ws_trust_header": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to return the WS-Trust token as a SOAP header. By default, the token is put in the SOAP body. When enabled, return the token as a SOAP header by wrapping the <tt>wst:RequestedSecurityToken</tt> by a <tt>wst:IssuedToken</tt> .", "ws-trust-in-header", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPPostProcessPPWSTrustHeaderIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"pp_ws_sc_key_source": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the source of the key. For WS-Trust postprocessing, the DataPower Gateway works as an on-box WS-Trust security token service that is backed by WS-SecureConversation. A symmetric shared secret key is needed to initialize the WS-SecureConversation SecurityContext. By default, a random key is generated.", "ws-trust-key-source", "").AddStringEnum("client-entropy", "in-kerberos", "in-encryptedkey", "static", "random").AddDefaultValue("random").AddNotValidWhen(DmAAAPPostProcessPPWSSCKeySourceIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("client-entropy", "in-kerberos", "in-encryptedkey", "static", "random"),
					validators.ConditionalRequiredString(validators.Evaluation{}, DmAAAPPostProcessPPWSSCKeySourceIgnoreVal, true),
				},
				Default: stringdefault.StaticString("random"),
			},
			"pp_shared_secret_key": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the shared secret key as the WS-Trust key source.", "ws-trust-shared-key", "crypto_sskey").AddNotValidWhen(DmAAAPPostProcessPPSharedSecretKeyIgnoreVal.String()).String,
				Optional:            true,
			},
			"pp_ws_trust_renewal_wait": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration to allow the STS to keep an expired SecurityContext token in seconds. After a WS-Trust token expires, it can be removed from the STS and cannot be renewed. Therefore, the token must be renewed before expiry. Enter a value in the range of 0 - 2678400. The default value is 0. <p>The token is issued or renewed with a 1-hour wait time in the following situation.</p><ul><li>The WS-Trust request asks that the issued token can be renewed after expiration.</li><li>This setting has a value of 0.</li></ul>", "ws-trust-renewal-wait", "").AddIntegerRange(0, 2678400).AddDefaultValue("0").AddNotValidWhen(DmAAAPPostProcessPPWSTrustRenewalWaitIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 2678400),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, DmAAAPPostProcessPPWSTrustRenewalWaitIgnoreVal, true),
				},
				Default: int64default.StaticInt64(0),
			},
			"pp_ws_trust_new_instance": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the STS renewal request issues a new instance for WS-Trust renewal. By default, the STS renewal request renews the existing instance. When enabled, the STS renewal request creates a new instance.", "ws-trust-new-instance", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPPostProcessPPWSTrustNewInstanceIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"pp_ws_trust_new_key": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to update the context key for WS-Trust renewal.By default, the SCT renewal request uses the existing shared secret key. When enabled, the SCT renewal request does not use the existing shared secret key.", "ws-trust-new-key", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPPostProcessPPWSTrustNewKeyIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"pp_ws_trust_never_expire": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the WS-Trust security context expires. By default, the security context expires. When enabled, the security context never expires.However, you can change the duration afterward with an explicit duration in seconds before expiry.", "ws-trust-never-expire", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPPostProcessPPWSTrustNeverExpireIgnoreVal.String()).String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the realm of a user for ICRX identity propagation. The ICRX realm is defined in the SAF configuration. Generally, this value is the equivalent of the prefix for a DN in a user registry.", "icrx-user-realm", "").AddNotValidWhen(DmAAAPPostProcessPPICRXUserRealmIgnoreVal.String()).String,
				Optional:            true,
			},
			"pp_saml_identity_provider": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to generate a SAML assertion. The SAML assertion can contain an authentication statement, an authorization statement, an attribute statement, or any combination of these statements. The SAML attribute value can be a user LDAP Attribute value that can be retrieved in the following ways.</p><ul><li>Directly by the LDAP authentication or authorization method with the list of LDAP attribute names that are defined by user auxiliary LDAP attributes.</li><li>Indirectly with the <tt>var://context/ldap/auxiliary-attributes</tt> variable in a stylesheet or GatewayScript file. A call with <tt>dp:ldap-search</tt> to the user registry, and put the <tt>attribute-value</tt> elements of search result to the variable.</li></ul><p>To sign the SAML assertion, configure a WS-Security sign action or SAML enveloped sign action after the AAA action in the processing rule.</p>", "generate-saml-assertion", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"pp_saml_protocol": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the SAML protocol to wrap up the SAML assertion. By default, the SAML assertion can be put to WS-Security wrap-up later.", "saml-protocol", "").AddStringEnum("assertion", "response").AddDefaultValue("assertion").AddNotValidWhen(DmAAAPPostProcessPPSAMLProtocolIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("assertion", "response"),
					validators.ConditionalRequiredString(validators.Evaluation{}, DmAAAPPostProcessPPSAMLProtocolIgnoreVal, true),
				},
				Default: stringdefault.StaticString("assertion"),
			},
			"pp_saml_response_destination": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the destination for a SAML response. This information can prevent malicious forwarding of requests to unintended recipients, which is a required protection by some protocol bindings. If it is present, the actual recipient must check that the URI reference identifies the location at which the message was received. If it does not check that the URI reference identifies the location, the request must be discarded. Some protocol bindings might require the use of this attribute.", "saml-response-destination", "").AddNotValidWhen(DmAAAPPostProcessPPSAMLResponseDestinationIgnoreVal.String()).String,
				Optional:            true,
			},
			"pp_result_wrapup": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the method to generate the result. When the DataPower Gateway is configured for SOAP or WS-Security processing, different output methods can be used. By default, generates the results to an existing WS-Security message and replaces the same token in the requesting message.", "result-wrapup", "").AddStringEnum("wssec-replace", "wssec-new", "wssec-inject", "soap-body", "none").AddDefaultValue("wssec-replace").AddNotValidWhen(DmAAAPPostProcessPPResultWrapupIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("wssec-replace", "wssec-new", "wssec-inject", "soap-body", "none"),
					validators.ConditionalRequiredString(validators.Evaluation{}, DmAAAPPostProcessPPResultWrapupIgnoreVal, true),
				},
				Default: stringdefault.StaticString("wssec-replace"),
			},
			"pp_saml_assertion_type": GetDmSAMLStatementTypeResourceSchema("Specify the supported SAML statement types. By default, supports both attributes and authentication statements.", "saml-assertion-type", "", false),
			"pp_saml_subject_confirm": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the method that allows the destination system to confirm the subject of the SAML assertion. By default, the subject is bearer.", "saml-subject-confirm", "").AddStringEnum("bearer", "hok", "sv").AddDefaultValue("bearer").AddNotValidWhen(DmAAAPPostProcessPPSAMLSubjectConfirmIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("bearer", "hok", "sv"),
					validators.ConditionalRequiredString(validators.Evaluation{}, DmAAAPPostProcessPPSAMLSubjectConfirmIgnoreVal, true),
				},
				Default: stringdefault.StaticString("bearer"),
			},
			"pp_saml_name_id": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the SAML Subject element contains the name identifier. By default, the SAML subject contains the name identifier. When disabled, the SAML subject does not contain the name identifier. Use this value if the subject confirmation method is holder-of-key because the key represent the same entity as the subject.", "saml-nid", "").AddDefaultValue("true").AddNotValidWhen(DmAAAPPostProcessPPSAMLNameIDIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"pp_saml_name_id_format": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URI reference that represents the classification of string-based identifier information. Any standard or arbitrary URI is allowed. If the value is an empty string, the DataPower Gateway attempts to determine the value from the AAA context. Some SAML protocols require a specified value, such as <tt>urn:oasis:names:tc:SAML:2.0:nameid-format:entity</tt> or <tt>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</tt> .", "saml-nid-format", "").AddNotValidWhen(DmAAAPPostProcessPPSAMLNameIDFormatIgnoreVal.String()).String,
				Optional:            true,
			},
			"pp_saml_recipient": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify a URI that identifies the entity or location that an attesting entity can present the assertion to. Any standard or arbitrary URI is allowed. If the value is an empty string, the optional attribute is not generated. This setting is applicable for only SAML 2.0.", "saml-recipient", "").AddNotValidWhen(DmAAAPPostProcessPPSAMLRecipientIgnoreVal.String()).String,
				Optional:            true,
			},
			"pp_saml_audience": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify URI references that identify an intended audience. Enter any number of the audience URIs to process the generated SAML assertion. If the value is an empty string, the SAML audience is not restricted. If there is more than one audience URI, use a + delimiter between URIs. In this case, you must convert any URI that contains the + characters to \\+.", "saml-audience", "").AddNotValidWhen(DmAAAPPostProcessPPSAMLAudienceIgnoreVal.String()).String,
				Optional:            true,
			},
			"pp_saml_omit_not_before": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("whether to omit the <tt>NotBefore</tt> attribute in the SAML assertion. When omitted, the assertion is considered valid even before the time it was issued. By default, the <tt>NotBefore</tt> attribute is not omitted. When enabled, the <tt>NotBefore</tt> attribute in the SAML assertion is omitted. This behavior might be required to respond to an <tt>AuthnRequest</tt> .", "saml-omit-notbefore", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPPostProcessPPSAMLOmitNotBeforeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"pp_one_time_use": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the destination system or relying party should cache the generated token. The generated token might contain the property for this characteristic, which is especially practical for SAML assertions. By default, the destination system can cache the generated token. When enabled, he destination system should not cache the generated token.", "one-time-use", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPPostProcessPPOneTimeUseIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"pp_saml_proxy": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to allow SAML proxy restriction. The generated SAML assertion provides limitations that the asserting party imposes on relying parties that want to act as asserting parties.</p><ul><li>A relying party that acts as an asserting party can issue subsequent assertions that are based on the information in the original assertion.</li><li>The relying party cannot issue an assertion that violates these restrictions.</li></ul><p>By default, proxy restrictions are not allowd. When enabled, proxy restrictions are allows.</p>", "saml-proxy", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPPostProcessPPSAMLProxyIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"pp_saml_proxy_audience": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the set of audiences (proxy) to whom the asserting party permits new assertions to be issued based on this assertion. If the value is an empty string, the audience for the <tt>ProxyRestriction</tt> is not issued with this SAML assertion. If there is more than one audience URI, use a + delimiter between URIs. In this case, you must convert any URI that contains the + characters to \\+.", "saml-proxy-audience", "").AddNotValidWhen(DmAAAPPostProcessPPSAMLProxyAudienceIgnoreVal.String()).String,
				Optional:            true,
			},
			"pp_saml_proxy_count": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum number of indirections that the asserting party permits between this assertion and an assertion that was issued. Enter a value in the range 0 - 65535. The default value is 0. A value of 0 indicates that a relying party must not issue an assertion to another relying party based on this assertion. If greater than zero, any assertion that is issued must itself contain a <tt>ProxyRestriction</tt> element with a <tt>Count</tt> value of at most one less than this value.", "saml-proxy-count", "").AddIntegerRange(0, 65535).AddDefaultValue("0").AddNotValidWhen(DmAAAPPostProcessPPSAMLProxyCountIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, DmAAAPPostProcessPPSAMLProxyCountIgnoreVal, true),
				},
				Default: int64default.StaticInt64(0),
			},
			"pp_saml_authz_action": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the standard action that the subject can take on the resource. The SAML specification defines the list of action identifiers with corresponding namespace URIs. By default, all HTTP operations, where <tt>urn:oasis:names:tc:SAML:1.0:action:ghpp</tt> is the namespace URI.", "saml-authz-action", "").AddStringEnum("AllHTTP", "POST", "GET", "PUT", "HEAD", "General", "Read", "Write", "Execute", "Delete", "Control", "NegatedRead", "NegatedWrite", "NegatedExecute", "NegatedDelete", "NegatedControl").AddDefaultValue("AllHTTP").AddNotValidWhen(DmAAAPPostProcessPPSAMLAuthzActionIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("AllHTTP", "POST", "GET", "PUT", "HEAD", "General", "Read", "Write", "Execute", "Delete", "Control", "NegatedRead", "NegatedWrite", "NegatedExecute", "NegatedDelete", "NegatedControl"),
					validators.ConditionalRequiredString(validators.Evaluation{}, DmAAAPPostProcessPPSAMLAuthzActionIgnoreVal, true),
				},
				Default: stringdefault.StaticString("AllHTTP"),
			},
			"pp_saml_attributes": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of an existing SAML attributes. The SAML attributes define the information to put in the SAML assertion to generate the attribute statement. Each SAML attribute requires the name, format or namespace, and value. The value can be from a DataPower variable.", "saml-attributes", "saml_attributes").AddRequiredWhen(DmAAAPPostProcessPPSAMLAttributesCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPSAMLAttributesIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPPostProcessPPSAMLAttributesCondVal, DmAAAPPostProcessPPSAMLAttributesIgnoreVal, false),
				},
			},
			"pp_ltpa_insert_cookie": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to insert a <tt>Set-Cookie</tt> header in the response that contains the LTPA token. This setting is for generating LTPA tokens that are not wrapped in the WS-Security <tt>wsse:Security</tt> header. By default, inserts a Set-Cookie header in the response. When disabled, does not insert a Set-Cookie header in the response.", "ltpa-insert-cookie", "").AddDefaultValue("true").AddNotValidWhen(DmAAAPPostProcessPPLTPAInsertCookieIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"pp_tam_pac_propagate": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to add the Access Manager privilege attribute certificate (PAC) token to an HTTP header. The PAC token was returned from the previous authentication or authorization phase. By default, does not add the PAC token. When enabled, adds the PAC token.", "propagate-tam-pac", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"pp_tam_header": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the HTTP header to store the token in. The default value is iv_creds, which is HTTP header that WebSEAL uses to write headers.", "tam-header", "").AddDefaultValue("iv-creds").AddRequiredWhen(DmAAAPPostProcessPPTAMHeaderCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPTAMHeaderIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPPostProcessPPTAMHeaderCondVal, DmAAAPPostProcessPPTAMHeaderIgnoreVal, true),
				},
				Default: stringdefault.StaticString("iv-creds"),
			},
			"pp_tam_header_size": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum size in bytes of HTTP headers. A value of 0 disables this function. If the value is nonzero, the PAC token is split across multiple headers of the specified length. The default value is 0.", "tam-header-size", "").AddDefaultValue("0").AddRequiredWhen(DmAAAPPostProcessPPTAMHeaderSizeCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPTAMHeaderSizeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					validators.ConditionalRequiredInt64(DmAAAPPostProcessPPTAMHeaderSizeCondVal, DmAAAPPostProcessPPTAMHeaderSizeIgnoreVal, true),
				},
				Default: int64default.StaticInt64(0),
			},
			"pp_kerberos_use_s4u2proxy": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to use constrained delegation, namely S4U2Proxy, when a WS-Security Kerberos AP-REQ token or a Kerberos SPNEGO token is generated. By default, does not use constrained delegation. When enabled, uses constrained delegation.", "kerberos-use-s4u2proxy", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPPostProcessPPKerberosUseS4U2ProxyIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"pp_cookie_attributes": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the cookie attribute policy to include standard or custom attributes in the cookie. The response message that contains a <tt>Set-Cookie</tt> header is updated with the attributes defined in this policy.", "cookie-attributes", "cookie_attribute_policy").AddNotValidWhen(DmAAAPPostProcessPPCookieAttributesIgnoreVal.String()).String,
				Optional:            true,
			},
			"pp_kerberos_use_s4u2self_and_s4u2proxy": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to use protocol transition, namely S4U2Self, and then use constrained delegation, namely S4U2Proxy.</p><ul><li>Use S4U2Self to convert a non-Kerberos token to a Kerberos token to the DataPower Gateway itself.</li><li>Use S4U2Proxy to generate a WS-Security Kerberos AP-REQ token or a Kerberos SPNEGO token.</li></ul><p>By default, does not use protocol transition and constrained delegation. When enabled, uses protocol transition and constrained delegation.</p>", "kerberos-use-s4u2self", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPPostProcessPPKerberosUseS4U2SelfAndS4U2ProxyIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"pp_kerberos_client_source": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify where to get the principal name of the Kerberos client. By default, uses the output of credential mapping. The client principal is based on the authenticated identity, which is followed by the corresponding realm name. For example, if the authenticated user is <tt>alice</tt> , the client principal name can be <tt>HTTP/alice.datapower.com@DATAPOWER.COM</tt> . The client principal must be present in the KDC for S4U2Self to work.", "kerberos-client-source", "").AddStringEnum("mc-output", "custom-url", "ctx-var").AddDefaultValue("mc-output").AddRequiredWhen(DmAAAPPostProcessPPKerberosClientSourceCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPKerberosClientSourceIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("mc-output", "custom-url", "ctx-var"),
					validators.ConditionalRequiredString(DmAAAPPostProcessPPKerberosClientSourceCondVal, DmAAAPPostProcessPPKerberosClientSourceIgnoreVal, true),
				},
				Default: stringdefault.StaticString("mc-output"),
			},
			"pp_kerberos_self": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the principal name of the DataPower Gateway.", "kerberos-self-principal", "").AddRequiredWhen(DmAAAPPostProcessPPKerberosSelfCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPKerberosSelfIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPPostProcessPPKerberosSelfCondVal, DmAAAPPostProcessPPKerberosSelfIgnoreVal, false),
				},
			},
			"pp_kerberos_self_keytab": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the Kerberos keytab that defines the keytab for the DataPower Gateway. This keytab is required to authenticate the DataPower Gateway to the KDC.", "kerberos-self-keytab", "crypto_kerberos_keytab").AddRequiredWhen(DmAAAPPostProcessPPKerberosSelfKeytabCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPKerberosSelfKeytabIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPPostProcessPPKerberosSelfKeytabCondVal, DmAAAPPostProcessPPKerberosSelfKeytabIgnoreVal, false),
				},
			},
			"pp_kerberos_client_custom_url": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the location of the stylesheet or GatewayScript file. This file returns the client principal name within the <tt>kerberos-client-principal</tt> element. This file gets the following input.</p><ul><li>The output of all the steps that are executed in this AAA action.</li><li>The incoming request message.</li></ul>", "kerberos-client-custom-url", "").AddRequiredWhen(DmAAAPPostProcessPPKerberosClientCustomURLCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPKerberosClientCustomURLIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPPostProcessPPKerberosClientCustomURLCondVal, DmAAAPPostProcessPPKerberosClientCustomURLIgnoreVal, false),
				},
			},
			"pp_kerberos_client_ctx_var": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the context variable. The value of this context variable is used as the Kerberos client principal. This context variable must be specified in the <tt>var://context/name</tt> format. For example, <tt>var://context/AAA/krb-client-princ</tt> . You can use the set variable action to set this variable in the processing rule before the AAA action.", "kerberos-client-ctx-var", "").AddRequiredWhen(DmAAAPPostProcessPPKerberosClientCtxVarCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPKerberosClientCtxVarIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPPostProcessPPKerberosClientCtxVarCondVal, DmAAAPPostProcessPPKerberosClientCtxVarIgnoreVal, false),
				},
			},
			"pp_kerberos_server_source": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify where to get the principal name of the Kerberos server. By default, the server principal name is the value that is specified by the Kerberos server principal property. Ensure that the server principal is in the correct format. For example, <tt>HTTP/was-backend.datapower.com@DATAPOWER.COM</tt> .", "kerberos-server-source", "").AddStringEnum("as-is-string", "custom-url", "ctx-var").AddDefaultValue("as-is-string").AddRequiredWhen(DmAAAPPostProcessPPKerberosServerSourceCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPKerberosServerSourceIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("as-is-string", "custom-url", "ctx-var"),
					validators.ConditionalRequiredString(DmAAAPPostProcessPPKerberosServerSourceCondVal, DmAAAPPostProcessPPKerberosServerSourceIgnoreVal, true),
				},
				Default: stringdefault.StaticString("as-is-string"),
			},
			"pp_kerberos_server_custom_url": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the location of the stylesheet or GatewayScript file. This file returns the client principal name within the <tt>kerberos-server-principal</tt> element.</p><p>When constrained delegation is not used, this file gets the following input.</p><ul><li>The output of all phases that this AAA action processes.</li><li>The incoming request message.</li></ul><p>When constrained delegation is used, this file gets the following input.</p><ul><li>The output of only the identity extraction phase.</li><li>The incoming request message.</li></ul>", "kerberos-server-custom-url", "").AddRequiredWhen(DmAAAPPostProcessPPKerberosServerCustomURLCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPKerberosServerCustomURLIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPPostProcessPPKerberosServerCustomURLCondVal, DmAAAPPostProcessPPKerberosServerCustomURLIgnoreVal, false),
				},
			},
			"pp_kerberos_server_ctx_var": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the context variable. The value of this context variable is used as the Kerberos server principal. This context variable must be specified in the <tt>var://context/name format</tt> . For example, <tt>var:///context/AAA/krb-server-princ</tt> . You can use the set variable action to set this variable in the processing rule before the AAA action.", "kerberos-server-ctx-var", "").AddRequiredWhen(DmAAAPPostProcessPPKerberosServerCtxVarCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPKerberosServerCtxVarIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPPostProcessPPKerberosServerCtxVarCondVal, DmAAAPPostProcessPPKerberosServerCtxVarIgnoreVal, false),
				},
			},
			"pp_ssl_client_config_type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS client type to secure connections.", "ssl-client-type", "").AddStringEnum("proxy", "client").AddDefaultValue("proxy").AddNotValidWhen(DmAAAPPostProcessPPSSLClientConfigTypeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("proxy", "client"),
					validators.ConditionalRequiredString(validators.Evaluation{}, DmAAAPPostProcessPPSSLClientConfigTypeIgnoreVal, true),
				},
				Default: stringdefault.StaticString("proxy"),
			},
			"pp_ssl_client_profile": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS client profile to secure connections.", "ssl-client", "ssl_client_profile").AddNotValidWhen(DmAAAPPostProcessPPSSLClientProfileIgnoreVal.String()).String,
				Optional:            true,
			},
			"pp_ltpa_key_file_password_alias": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the the alias for password of the LTPA key file.", "ltpa-key-file-password-alias", "password_alias").AddRequiredWhen(DmAAAPPostProcessPPLTPAKeyFilePasswordAliasCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPLTPAKeyFilePasswordAliasIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPPostProcessPPLTPAKeyFilePasswordAliasCondVal, DmAAAPPostProcessPPLTPAKeyFilePasswordAliasIgnoreVal, false),
				},
			},
			"pp_jwt": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate a JWT token.", "jwt", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"pp_jwt_generator": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the JWT generator.", "generate-jwt", "aaa_jwt_generator").AddRequiredWhen(DmAAAPPostProcessPPJWTGeneratorCondVal.String()).AddNotValidWhen(DmAAAPPostProcessPPJWTGeneratorIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPPostProcessPPJWTGeneratorCondVal, DmAAAPPostProcessPPJWTGeneratorIgnoreVal, false),
				},
			},
		},
	}
	DmAAAPPostProcessResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmAAAPPostProcessResourceSchema.Required = true
	} else {
		DmAAAPPostProcessResourceSchema.Optional = true
		DmAAAPPostProcessResourceSchema.Computed = true
	}
	return DmAAAPPostProcessResourceSchema
}

func (data DmAAAPPostProcess) IsNull() bool {
	if !data.PpEnabled.IsNull() {
		return false
	}
	if !data.PpCustomUrl.IsNull() {
		return false
	}
	if !data.PpSamlAuthAssertion.IsNull() {
		return false
	}
	if !data.PpSamlServerName.IsNull() {
		return false
	}
	if !data.PpSamlNameQualifier.IsNull() {
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
	if !data.PpWsTrust.IsNull() {
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
	if !data.PpSamlVersion.IsNull() {
		return false
	}
	if !data.PpSamlSendSlo.IsNull() {
		return false
	}
	if !data.PpSamlSloEndpoint.IsNull() {
		return false
	}
	if !data.PpWsUsernameToken.IsNull() {
		return false
	}
	if !data.PpWsUsernameTokenPasswordType.IsNull() {
		return false
	}
	if !data.PpSamlValidity.IsNull() {
		return false
	}
	if !data.PpSamlSkew.IsNull() {
		return false
	}
	if !data.PpWsUsernameTokenIncludePwd.IsNull() {
		return false
	}
	if !data.PpLtpa.IsNull() {
		return false
	}
	if !data.PpLtpaVersion.IsNull() {
		return false
	}
	if !data.PpLtpaExpiry.IsNull() {
		return false
	}
	if !data.PpLtpaKeyFile.IsNull() {
		return false
	}
	if !data.PpLtpaKeyFilePassword.IsNull() {
		return false
	}
	if !data.PpLtpaStashFile.IsNull() {
		return false
	}
	if !data.PpKerberosSpnegoToken.IsNull() {
		return false
	}
	if !data.PpKerberosBstValueType.IsNull() {
		return false
	}
	if !data.PpSamlUseWsSec.IsNull() {
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
	if !data.PpWsDerivedKeyUsernameToken.IsNull() {
		return false
	}
	if !data.PpWsDerivedKeyUsernameTokenIterations.IsNull() {
		return false
	}
	if !data.PpWsUsernameTokenAllowReplacement.IsNull() {
		return false
	}
	if !data.PpHmacSigningAlg.IsNull() {
		return false
	}
	if !data.PpSigningHashAlg.IsNull() {
		return false
	}
	if !data.PpWsTrustHeader.IsNull() {
		return false
	}
	if !data.PpWsScKeySource.IsNull() {
		return false
	}
	if !data.PpSharedSecretKey.IsNull() {
		return false
	}
	if !data.PpWsTrustRenewalWait.IsNull() {
		return false
	}
	if !data.PpWsTrustNewInstance.IsNull() {
		return false
	}
	if !data.PpWsTrustNewKey.IsNull() {
		return false
	}
	if !data.PpWsTrustNeverExpire.IsNull() {
		return false
	}
	if !data.PpicrxToken.IsNull() {
		return false
	}
	if !data.PpicrxUserRealm.IsNull() {
		return false
	}
	if !data.PpSamlIdentityProvider.IsNull() {
		return false
	}
	if !data.PpSamlProtocol.IsNull() {
		return false
	}
	if !data.PpSamlResponseDestination.IsNull() {
		return false
	}
	if !data.PpResultWrapup.IsNull() {
		return false
	}
	if data.PpSamlAssertionType != nil {
		if !data.PpSamlAssertionType.IsNull() {
			return false
		}
	}
	if !data.PpSamlSubjectConfirm.IsNull() {
		return false
	}
	if !data.PpSamlNameId.IsNull() {
		return false
	}
	if !data.PpSamlNameIdFormat.IsNull() {
		return false
	}
	if !data.PpSamlRecipient.IsNull() {
		return false
	}
	if !data.PpSamlAudience.IsNull() {
		return false
	}
	if !data.PpSamlOmitNotBefore.IsNull() {
		return false
	}
	if !data.PpOneTimeUse.IsNull() {
		return false
	}
	if !data.PpSamlProxy.IsNull() {
		return false
	}
	if !data.PpSamlProxyAudience.IsNull() {
		return false
	}
	if !data.PpSamlProxyCount.IsNull() {
		return false
	}
	if !data.PpSamlAuthzAction.IsNull() {
		return false
	}
	if !data.PpSamlAttributes.IsNull() {
		return false
	}
	if !data.PpLtpaInsertCookie.IsNull() {
		return false
	}
	if !data.PpTamPacPropagate.IsNull() {
		return false
	}
	if !data.PpTamHeader.IsNull() {
		return false
	}
	if !data.PpTamHeaderSize.IsNull() {
		return false
	}
	if !data.PpKerberosUseS4u2proxy.IsNull() {
		return false
	}
	if !data.PpCookieAttributes.IsNull() {
		return false
	}
	if !data.PpKerberosUseS4u2selfAndS4u2proxy.IsNull() {
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
	if !data.PpSslClientConfigType.IsNull() {
		return false
	}
	if !data.PpSslClientProfile.IsNull() {
		return false
	}
	if !data.PpLtpaKeyFilePasswordAlias.IsNull() {
		return false
	}
	if !data.PpJwt.IsNull() {
		return false
	}
	if !data.PpJwtGenerator.IsNull() {
		return false
	}
	return true
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
	if !data.PpSamlAuthAssertion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLAuthAssertion`, tfutils.StringFromBool(data.PpSamlAuthAssertion, ""))
	}
	if !data.PpSamlServerName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLServerName`, data.PpSamlServerName.ValueString())
	}
	if !data.PpSamlNameQualifier.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLNameQualifier`, data.PpSamlNameQualifier.ValueString())
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
	if !data.PpWsTrust.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPWSTrust`, tfutils.StringFromBool(data.PpWsTrust, ""))
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
	if !data.PpSamlVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLVersion`, data.PpSamlVersion.ValueString())
	}
	if !data.PpSamlSendSlo.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLSendSLO`, tfutils.StringFromBool(data.PpSamlSendSlo, ""))
	}
	if !data.PpSamlSloEndpoint.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLSLOEndpoint`, data.PpSamlSloEndpoint.ValueString())
	}
	if !data.PpWsUsernameToken.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPWSUsernameToken`, tfutils.StringFromBool(data.PpWsUsernameToken, ""))
	}
	if !data.PpWsUsernameTokenPasswordType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPWSUsernameTokenPasswordType`, data.PpWsUsernameTokenPasswordType.ValueString())
	}
	if !data.PpSamlValidity.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLValidity`, data.PpSamlValidity.ValueInt64())
	}
	if !data.PpSamlSkew.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLSkew`, data.PpSamlSkew.ValueInt64())
	}
	if !data.PpWsUsernameTokenIncludePwd.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPWSUsernameTokenIncludePwd`, tfutils.StringFromBool(data.PpWsUsernameTokenIncludePwd, ""))
	}
	if !data.PpLtpa.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPLTPA`, tfutils.StringFromBool(data.PpLtpa, ""))
	}
	if !data.PpLtpaVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPLTPAVersion`, data.PpLtpaVersion.ValueString())
	}
	if !data.PpLtpaExpiry.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPLTPAExpiry`, data.PpLtpaExpiry.ValueInt64())
	}
	if !data.PpLtpaKeyFile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPLTPAKeyFile`, data.PpLtpaKeyFile.ValueString())
	}
	if !data.PpLtpaKeyFilePassword.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPLTPAKeyFilePassword`, data.PpLtpaKeyFilePassword.ValueString())
	}
	if !data.PpLtpaStashFile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPLTPAStashFile`, data.PpLtpaStashFile.ValueString())
	}
	if !data.PpKerberosSpnegoToken.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPKerberosSPNEGOToken`, tfutils.StringFromBool(data.PpKerberosSpnegoToken, ""))
	}
	if !data.PpKerberosBstValueType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPKerberosBstValueType`, data.PpKerberosBstValueType.ValueString())
	}
	if !data.PpSamlUseWsSec.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLUseWSSec`, tfutils.StringFromBool(data.PpSamlUseWsSec, ""))
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
	if !data.PpWsDerivedKeyUsernameToken.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPWSDerivedKeyUsernameToken`, tfutils.StringFromBool(data.PpWsDerivedKeyUsernameToken, ""))
	}
	if !data.PpWsDerivedKeyUsernameTokenIterations.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPWSDerivedKeyUsernameTokenIterations`, data.PpWsDerivedKeyUsernameTokenIterations.ValueInt64())
	}
	if !data.PpWsUsernameTokenAllowReplacement.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPWSUsernameTokenAllowReplacement`, tfutils.StringFromBool(data.PpWsUsernameTokenAllowReplacement, ""))
	}
	if !data.PpHmacSigningAlg.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPHMACSigningAlg`, data.PpHmacSigningAlg.ValueString())
	}
	if !data.PpSigningHashAlg.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSigningHashAlg`, data.PpSigningHashAlg.ValueString())
	}
	if !data.PpWsTrustHeader.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPWSTrustHeader`, tfutils.StringFromBool(data.PpWsTrustHeader, ""))
	}
	if !data.PpWsScKeySource.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPWSSCKeySource`, data.PpWsScKeySource.ValueString())
	}
	if !data.PpSharedSecretKey.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSharedSecretKey`, data.PpSharedSecretKey.ValueString())
	}
	if !data.PpWsTrustRenewalWait.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPWSTrustRenewalWait`, data.PpWsTrustRenewalWait.ValueInt64())
	}
	if !data.PpWsTrustNewInstance.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPWSTrustNewInstance`, tfutils.StringFromBool(data.PpWsTrustNewInstance, ""))
	}
	if !data.PpWsTrustNewKey.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPWSTrustNewKey`, tfutils.StringFromBool(data.PpWsTrustNewKey, ""))
	}
	if !data.PpWsTrustNeverExpire.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPWSTrustNeverExpire`, tfutils.StringFromBool(data.PpWsTrustNeverExpire, ""))
	}
	if !data.PpicrxToken.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPICRXToken`, tfutils.StringFromBool(data.PpicrxToken, ""))
	}
	if !data.PpicrxUserRealm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPICRXUserRealm`, data.PpicrxUserRealm.ValueString())
	}
	if !data.PpSamlIdentityProvider.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLIdentityProvider`, tfutils.StringFromBool(data.PpSamlIdentityProvider, ""))
	}
	if !data.PpSamlProtocol.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLProtocol`, data.PpSamlProtocol.ValueString())
	}
	if !data.PpSamlResponseDestination.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLResponseDestination`, data.PpSamlResponseDestination.ValueString())
	}
	if !data.PpResultWrapup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPResultWrapup`, data.PpResultWrapup.ValueString())
	}
	if data.PpSamlAssertionType != nil {
		if !data.PpSamlAssertionType.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`PPSAMLAssertionType`, data.PpSamlAssertionType.ToBody(ctx, ""))
		}
	}
	if !data.PpSamlSubjectConfirm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLSubjectConfirm`, data.PpSamlSubjectConfirm.ValueString())
	}
	if !data.PpSamlNameId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLNameID`, tfutils.StringFromBool(data.PpSamlNameId, ""))
	}
	if !data.PpSamlNameIdFormat.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLNameIDFormat`, data.PpSamlNameIdFormat.ValueString())
	}
	if !data.PpSamlRecipient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLRecipient`, data.PpSamlRecipient.ValueString())
	}
	if !data.PpSamlAudience.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLAudience`, data.PpSamlAudience.ValueString())
	}
	if !data.PpSamlOmitNotBefore.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLOmitNotBefore`, tfutils.StringFromBool(data.PpSamlOmitNotBefore, ""))
	}
	if !data.PpOneTimeUse.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPOneTimeUse`, tfutils.StringFromBool(data.PpOneTimeUse, ""))
	}
	if !data.PpSamlProxy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLProxy`, tfutils.StringFromBool(data.PpSamlProxy, ""))
	}
	if !data.PpSamlProxyAudience.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLProxyAudience`, data.PpSamlProxyAudience.ValueString())
	}
	if !data.PpSamlProxyCount.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLProxyCount`, data.PpSamlProxyCount.ValueInt64())
	}
	if !data.PpSamlAuthzAction.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLAuthzAction`, data.PpSamlAuthzAction.ValueString())
	}
	if !data.PpSamlAttributes.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSAMLAttributes`, data.PpSamlAttributes.ValueString())
	}
	if !data.PpLtpaInsertCookie.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPLTPAInsertCookie`, tfutils.StringFromBool(data.PpLtpaInsertCookie, ""))
	}
	if !data.PpTamPacPropagate.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPTAMPACPropagate`, tfutils.StringFromBool(data.PpTamPacPropagate, ""))
	}
	if !data.PpTamHeader.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPTAMHeader`, data.PpTamHeader.ValueString())
	}
	if !data.PpTamHeaderSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPTAMHeaderSize`, data.PpTamHeaderSize.ValueInt64())
	}
	if !data.PpKerberosUseS4u2proxy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPKerberosUseS4U2Proxy`, tfutils.StringFromBool(data.PpKerberosUseS4u2proxy, ""))
	}
	if !data.PpCookieAttributes.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPCookieAttributes`, data.PpCookieAttributes.ValueString())
	}
	if !data.PpKerberosUseS4u2selfAndS4u2proxy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPKerberosUseS4U2SelfAndS4U2Proxy`, tfutils.StringFromBool(data.PpKerberosUseS4u2selfAndS4u2proxy, ""))
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
	if !data.PpSslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSSLClientConfigType`, data.PpSslClientConfigType.ValueString())
	}
	if !data.PpSslClientProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPSSLClientProfile`, data.PpSslClientProfile.ValueString())
	}
	if !data.PpLtpaKeyFilePasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPLTPAKeyFilePasswordAlias`, data.PpLtpaKeyFilePasswordAlias.ValueString())
	}
	if !data.PpJwt.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPJWT`, tfutils.StringFromBool(data.PpJwt, ""))
	}
	if !data.PpJwtGenerator.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PPJWTGenerator`, data.PpJwtGenerator.ValueString())
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
		data.PpSamlAuthAssertion = tfutils.BoolFromString(value.String())
	} else {
		data.PpSamlAuthAssertion = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPSAMLServerName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpSamlServerName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpSamlServerName = types.StringValue("XS")
	}
	if value := res.Get(pathRoot + `PPSAMLNameQualifier`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpSamlNameQualifier = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpSamlNameQualifier = types.StringNull()
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
		data.PpWsTrust = tfutils.BoolFromString(value.String())
	} else {
		data.PpWsTrust = types.BoolNull()
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
		data.PpSamlVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpSamlVersion = types.StringValue("2")
	}
	if value := res.Get(pathRoot + `PPSAMLSendSLO`); value.Exists() {
		data.PpSamlSendSlo = tfutils.BoolFromString(value.String())
	} else {
		data.PpSamlSendSlo = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPSAMLSLOEndpoint`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpSamlSloEndpoint = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpSamlSloEndpoint = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPWSUsernameToken`); value.Exists() {
		data.PpWsUsernameToken = tfutils.BoolFromString(value.String())
	} else {
		data.PpWsUsernameToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPWSUsernameTokenPasswordType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpWsUsernameTokenPasswordType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpWsUsernameTokenPasswordType = types.StringValue("Digest")
	}
	if value := res.Get(pathRoot + `PPSAMLValidity`); value.Exists() {
		data.PpSamlValidity = types.Int64Value(value.Int())
	} else {
		data.PpSamlValidity = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `PPSAMLSkew`); value.Exists() {
		data.PpSamlSkew = types.Int64Value(value.Int())
	} else {
		data.PpSamlSkew = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `PPWSUsernameTokenIncludePwd`); value.Exists() {
		data.PpWsUsernameTokenIncludePwd = tfutils.BoolFromString(value.String())
	} else {
		data.PpWsUsernameTokenIncludePwd = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPLTPA`); value.Exists() {
		data.PpLtpa = tfutils.BoolFromString(value.String())
	} else {
		data.PpLtpa = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPLTPAVersion`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpLtpaVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpLtpaVersion = types.StringValue("LTPA2")
	}
	if value := res.Get(pathRoot + `PPLTPAExpiry`); value.Exists() {
		data.PpLtpaExpiry = types.Int64Value(value.Int())
	} else {
		data.PpLtpaExpiry = types.Int64Value(600)
	}
	if value := res.Get(pathRoot + `PPLTPAKeyFile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpLtpaKeyFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpLtpaKeyFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPLTPAKeyFilePassword`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpLtpaKeyFilePassword = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpLtpaKeyFilePassword = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPLTPAStashFile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpLtpaStashFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpLtpaStashFile = types.StringNull()
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
		data.PpSamlUseWsSec = tfutils.BoolFromString(value.String())
	} else {
		data.PpSamlUseWsSec = types.BoolNull()
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
		data.PpWsDerivedKeyUsernameToken = tfutils.BoolFromString(value.String())
	} else {
		data.PpWsDerivedKeyUsernameToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPWSDerivedKeyUsernameTokenIterations`); value.Exists() {
		data.PpWsDerivedKeyUsernameTokenIterations = types.Int64Value(value.Int())
	} else {
		data.PpWsDerivedKeyUsernameTokenIterations = types.Int64Value(1000)
	}
	if value := res.Get(pathRoot + `PPWSUsernameTokenAllowReplacement`); value.Exists() {
		data.PpWsUsernameTokenAllowReplacement = tfutils.BoolFromString(value.String())
	} else {
		data.PpWsUsernameTokenAllowReplacement = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPHMACSigningAlg`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpHmacSigningAlg = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpHmacSigningAlg = types.StringValue("hmac-sha1")
	}
	if value := res.Get(pathRoot + `PPSigningHashAlg`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpSigningHashAlg = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpSigningHashAlg = types.StringValue("sha1")
	}
	if value := res.Get(pathRoot + `PPWSTrustHeader`); value.Exists() {
		data.PpWsTrustHeader = tfutils.BoolFromString(value.String())
	} else {
		data.PpWsTrustHeader = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPWSSCKeySource`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpWsScKeySource = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpWsScKeySource = types.StringValue("random")
	}
	if value := res.Get(pathRoot + `PPSharedSecretKey`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpSharedSecretKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpSharedSecretKey = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPWSTrustRenewalWait`); value.Exists() {
		data.PpWsTrustRenewalWait = types.Int64Value(value.Int())
	} else {
		data.PpWsTrustRenewalWait = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `PPWSTrustNewInstance`); value.Exists() {
		data.PpWsTrustNewInstance = tfutils.BoolFromString(value.String())
	} else {
		data.PpWsTrustNewInstance = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPWSTrustNewKey`); value.Exists() {
		data.PpWsTrustNewKey = tfutils.BoolFromString(value.String())
	} else {
		data.PpWsTrustNewKey = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPWSTrustNeverExpire`); value.Exists() {
		data.PpWsTrustNeverExpire = tfutils.BoolFromString(value.String())
	} else {
		data.PpWsTrustNeverExpire = types.BoolNull()
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
		data.PpSamlIdentityProvider = tfutils.BoolFromString(value.String())
	} else {
		data.PpSamlIdentityProvider = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPSAMLProtocol`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpSamlProtocol = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpSamlProtocol = types.StringValue("assertion")
	}
	if value := res.Get(pathRoot + `PPSAMLResponseDestination`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpSamlResponseDestination = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpSamlResponseDestination = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPResultWrapup`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpResultWrapup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpResultWrapup = types.StringValue("wssec-replace")
	}
	if value := res.Get(pathRoot + `PPSAMLAssertionType`); value.Exists() {
		data.PpSamlAssertionType = &DmSAMLStatementType{}
		data.PpSamlAssertionType.FromBody(ctx, "", value)
	} else {
		data.PpSamlAssertionType = nil
	}
	if value := res.Get(pathRoot + `PPSAMLSubjectConfirm`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpSamlSubjectConfirm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpSamlSubjectConfirm = types.StringValue("bearer")
	}
	if value := res.Get(pathRoot + `PPSAMLNameID`); value.Exists() {
		data.PpSamlNameId = tfutils.BoolFromString(value.String())
	} else {
		data.PpSamlNameId = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPSAMLNameIDFormat`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpSamlNameIdFormat = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpSamlNameIdFormat = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLRecipient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpSamlRecipient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpSamlRecipient = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLAudience`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpSamlAudience = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpSamlAudience = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLOmitNotBefore`); value.Exists() {
		data.PpSamlOmitNotBefore = tfutils.BoolFromString(value.String())
	} else {
		data.PpSamlOmitNotBefore = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPOneTimeUse`); value.Exists() {
		data.PpOneTimeUse = tfutils.BoolFromString(value.String())
	} else {
		data.PpOneTimeUse = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPSAMLProxy`); value.Exists() {
		data.PpSamlProxy = tfutils.BoolFromString(value.String())
	} else {
		data.PpSamlProxy = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPSAMLProxyAudience`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpSamlProxyAudience = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpSamlProxyAudience = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLProxyCount`); value.Exists() {
		data.PpSamlProxyCount = types.Int64Value(value.Int())
	} else {
		data.PpSamlProxyCount = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `PPSAMLAuthzAction`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpSamlAuthzAction = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpSamlAuthzAction = types.StringValue("AllHTTP")
	}
	if value := res.Get(pathRoot + `PPSAMLAttributes`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpSamlAttributes = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpSamlAttributes = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPLTPAInsertCookie`); value.Exists() {
		data.PpLtpaInsertCookie = tfutils.BoolFromString(value.String())
	} else {
		data.PpLtpaInsertCookie = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPTAMPACPropagate`); value.Exists() {
		data.PpTamPacPropagate = tfutils.BoolFromString(value.String())
	} else {
		data.PpTamPacPropagate = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPTAMHeader`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpTamHeader = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpTamHeader = types.StringValue("iv-creds")
	}
	if value := res.Get(pathRoot + `PPTAMHeaderSize`); value.Exists() {
		data.PpTamHeaderSize = types.Int64Value(value.Int())
	} else {
		data.PpTamHeaderSize = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `PPKerberosUseS4U2Proxy`); value.Exists() {
		data.PpKerberosUseS4u2proxy = tfutils.BoolFromString(value.String())
	} else {
		data.PpKerberosUseS4u2proxy = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPCookieAttributes`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpCookieAttributes = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpCookieAttributes = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPKerberosUseS4U2SelfAndS4U2Proxy`); value.Exists() {
		data.PpKerberosUseS4u2selfAndS4u2proxy = tfutils.BoolFromString(value.String())
	} else {
		data.PpKerberosUseS4u2selfAndS4u2proxy = types.BoolNull()
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
		data.PpSslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpSslClientConfigType = types.StringValue("proxy")
	}
	if value := res.Get(pathRoot + `PPSSLClientProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpSslClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpSslClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPLTPAKeyFilePasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpLtpaKeyFilePasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpLtpaKeyFilePasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPJWT`); value.Exists() {
		data.PpJwt = tfutils.BoolFromString(value.String())
	} else {
		data.PpJwt = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPJWTGenerator`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PpJwtGenerator = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpJwtGenerator = types.StringNull()
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
	if value := res.Get(pathRoot + `PPSAMLAuthAssertion`); value.Exists() && !data.PpSamlAuthAssertion.IsNull() {
		data.PpSamlAuthAssertion = tfutils.BoolFromString(value.String())
	} else if data.PpSamlAuthAssertion.ValueBool() {
		data.PpSamlAuthAssertion = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPSAMLServerName`); value.Exists() && !data.PpSamlServerName.IsNull() {
		data.PpSamlServerName = tfutils.ParseStringFromGJSON(value)
	} else if data.PpSamlServerName.ValueString() != "XS" {
		data.PpSamlServerName = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLNameQualifier`); value.Exists() && !data.PpSamlNameQualifier.IsNull() {
		data.PpSamlNameQualifier = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpSamlNameQualifier = types.StringNull()
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
	if value := res.Get(pathRoot + `PPWSTrust`); value.Exists() && !data.PpWsTrust.IsNull() {
		data.PpWsTrust = tfutils.BoolFromString(value.String())
	} else if data.PpWsTrust.ValueBool() {
		data.PpWsTrust = types.BoolNull()
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
	if value := res.Get(pathRoot + `PPSAMLVersion`); value.Exists() && !data.PpSamlVersion.IsNull() {
		data.PpSamlVersion = tfutils.ParseStringFromGJSON(value)
	} else if data.PpSamlVersion.ValueString() != "2" {
		data.PpSamlVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLSendSLO`); value.Exists() && !data.PpSamlSendSlo.IsNull() {
		data.PpSamlSendSlo = tfutils.BoolFromString(value.String())
	} else if data.PpSamlSendSlo.ValueBool() {
		data.PpSamlSendSlo = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPSAMLSLOEndpoint`); value.Exists() && !data.PpSamlSloEndpoint.IsNull() {
		data.PpSamlSloEndpoint = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpSamlSloEndpoint = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPWSUsernameToken`); value.Exists() && !data.PpWsUsernameToken.IsNull() {
		data.PpWsUsernameToken = tfutils.BoolFromString(value.String())
	} else if data.PpWsUsernameToken.ValueBool() {
		data.PpWsUsernameToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPWSUsernameTokenPasswordType`); value.Exists() && !data.PpWsUsernameTokenPasswordType.IsNull() {
		data.PpWsUsernameTokenPasswordType = tfutils.ParseStringFromGJSON(value)
	} else if data.PpWsUsernameTokenPasswordType.ValueString() != "Digest" {
		data.PpWsUsernameTokenPasswordType = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLValidity`); value.Exists() && !data.PpSamlValidity.IsNull() {
		data.PpSamlValidity = types.Int64Value(value.Int())
	} else if data.PpSamlValidity.ValueInt64() != 0 {
		data.PpSamlValidity = types.Int64Null()
	}
	if value := res.Get(pathRoot + `PPSAMLSkew`); value.Exists() && !data.PpSamlSkew.IsNull() {
		data.PpSamlSkew = types.Int64Value(value.Int())
	} else if data.PpSamlSkew.ValueInt64() != 0 {
		data.PpSamlSkew = types.Int64Null()
	}
	if value := res.Get(pathRoot + `PPWSUsernameTokenIncludePwd`); value.Exists() && !data.PpWsUsernameTokenIncludePwd.IsNull() {
		data.PpWsUsernameTokenIncludePwd = tfutils.BoolFromString(value.String())
	} else if !data.PpWsUsernameTokenIncludePwd.ValueBool() {
		data.PpWsUsernameTokenIncludePwd = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPLTPA`); value.Exists() && !data.PpLtpa.IsNull() {
		data.PpLtpa = tfutils.BoolFromString(value.String())
	} else if data.PpLtpa.ValueBool() {
		data.PpLtpa = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPLTPAVersion`); value.Exists() && !data.PpLtpaVersion.IsNull() {
		data.PpLtpaVersion = tfutils.ParseStringFromGJSON(value)
	} else if data.PpLtpaVersion.ValueString() != "LTPA2" {
		data.PpLtpaVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPLTPAExpiry`); value.Exists() && !data.PpLtpaExpiry.IsNull() {
		data.PpLtpaExpiry = types.Int64Value(value.Int())
	} else if data.PpLtpaExpiry.ValueInt64() != 600 {
		data.PpLtpaExpiry = types.Int64Null()
	}
	if value := res.Get(pathRoot + `PPLTPAKeyFile`); value.Exists() && !data.PpLtpaKeyFile.IsNull() {
		data.PpLtpaKeyFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpLtpaKeyFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPLTPAKeyFilePassword`); value.Exists() && !data.PpLtpaKeyFilePassword.IsNull() {
		data.PpLtpaKeyFilePassword = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpLtpaKeyFilePassword = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPLTPAStashFile`); value.Exists() && !data.PpLtpaStashFile.IsNull() {
		data.PpLtpaStashFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpLtpaStashFile = types.StringNull()
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
	if value := res.Get(pathRoot + `PPSAMLUseWSSec`); value.Exists() && !data.PpSamlUseWsSec.IsNull() {
		data.PpSamlUseWsSec = tfutils.BoolFromString(value.String())
	} else if data.PpSamlUseWsSec.ValueBool() {
		data.PpSamlUseWsSec = types.BoolNull()
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
	if value := res.Get(pathRoot + `PPWSDerivedKeyUsernameToken`); value.Exists() && !data.PpWsDerivedKeyUsernameToken.IsNull() {
		data.PpWsDerivedKeyUsernameToken = tfutils.BoolFromString(value.String())
	} else if data.PpWsDerivedKeyUsernameToken.ValueBool() {
		data.PpWsDerivedKeyUsernameToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPWSDerivedKeyUsernameTokenIterations`); value.Exists() && !data.PpWsDerivedKeyUsernameTokenIterations.IsNull() {
		data.PpWsDerivedKeyUsernameTokenIterations = types.Int64Value(value.Int())
	} else if data.PpWsDerivedKeyUsernameTokenIterations.ValueInt64() != 1000 {
		data.PpWsDerivedKeyUsernameTokenIterations = types.Int64Null()
	}
	if value := res.Get(pathRoot + `PPWSUsernameTokenAllowReplacement`); value.Exists() && !data.PpWsUsernameTokenAllowReplacement.IsNull() {
		data.PpWsUsernameTokenAllowReplacement = tfutils.BoolFromString(value.String())
	} else if data.PpWsUsernameTokenAllowReplacement.ValueBool() {
		data.PpWsUsernameTokenAllowReplacement = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPHMACSigningAlg`); value.Exists() && !data.PpHmacSigningAlg.IsNull() {
		data.PpHmacSigningAlg = tfutils.ParseStringFromGJSON(value)
	} else if data.PpHmacSigningAlg.ValueString() != "hmac-sha1" {
		data.PpHmacSigningAlg = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSigningHashAlg`); value.Exists() && !data.PpSigningHashAlg.IsNull() {
		data.PpSigningHashAlg = tfutils.ParseStringFromGJSON(value)
	} else if data.PpSigningHashAlg.ValueString() != "sha1" {
		data.PpSigningHashAlg = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPWSTrustHeader`); value.Exists() && !data.PpWsTrustHeader.IsNull() {
		data.PpWsTrustHeader = tfutils.BoolFromString(value.String())
	} else if data.PpWsTrustHeader.ValueBool() {
		data.PpWsTrustHeader = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPWSSCKeySource`); value.Exists() && !data.PpWsScKeySource.IsNull() {
		data.PpWsScKeySource = tfutils.ParseStringFromGJSON(value)
	} else if data.PpWsScKeySource.ValueString() != "random" {
		data.PpWsScKeySource = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSharedSecretKey`); value.Exists() && !data.PpSharedSecretKey.IsNull() {
		data.PpSharedSecretKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpSharedSecretKey = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPWSTrustRenewalWait`); value.Exists() && !data.PpWsTrustRenewalWait.IsNull() {
		data.PpWsTrustRenewalWait = types.Int64Value(value.Int())
	} else if data.PpWsTrustRenewalWait.ValueInt64() != 0 {
		data.PpWsTrustRenewalWait = types.Int64Null()
	}
	if value := res.Get(pathRoot + `PPWSTrustNewInstance`); value.Exists() && !data.PpWsTrustNewInstance.IsNull() {
		data.PpWsTrustNewInstance = tfutils.BoolFromString(value.String())
	} else if data.PpWsTrustNewInstance.ValueBool() {
		data.PpWsTrustNewInstance = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPWSTrustNewKey`); value.Exists() && !data.PpWsTrustNewKey.IsNull() {
		data.PpWsTrustNewKey = tfutils.BoolFromString(value.String())
	} else if data.PpWsTrustNewKey.ValueBool() {
		data.PpWsTrustNewKey = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPWSTrustNeverExpire`); value.Exists() && !data.PpWsTrustNeverExpire.IsNull() {
		data.PpWsTrustNeverExpire = tfutils.BoolFromString(value.String())
	} else if data.PpWsTrustNeverExpire.ValueBool() {
		data.PpWsTrustNeverExpire = types.BoolNull()
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
	if value := res.Get(pathRoot + `PPSAMLIdentityProvider`); value.Exists() && !data.PpSamlIdentityProvider.IsNull() {
		data.PpSamlIdentityProvider = tfutils.BoolFromString(value.String())
	} else if data.PpSamlIdentityProvider.ValueBool() {
		data.PpSamlIdentityProvider = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPSAMLProtocol`); value.Exists() && !data.PpSamlProtocol.IsNull() {
		data.PpSamlProtocol = tfutils.ParseStringFromGJSON(value)
	} else if data.PpSamlProtocol.ValueString() != "assertion" {
		data.PpSamlProtocol = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLResponseDestination`); value.Exists() && !data.PpSamlResponseDestination.IsNull() {
		data.PpSamlResponseDestination = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpSamlResponseDestination = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPResultWrapup`); value.Exists() && !data.PpResultWrapup.IsNull() {
		data.PpResultWrapup = tfutils.ParseStringFromGJSON(value)
	} else if data.PpResultWrapup.ValueString() != "wssec-replace" {
		data.PpResultWrapup = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLAssertionType`); value.Exists() {
		data.PpSamlAssertionType.UpdateFromBody(ctx, "", value)
	} else {
		data.PpSamlAssertionType = nil
	}
	if value := res.Get(pathRoot + `PPSAMLSubjectConfirm`); value.Exists() && !data.PpSamlSubjectConfirm.IsNull() {
		data.PpSamlSubjectConfirm = tfutils.ParseStringFromGJSON(value)
	} else if data.PpSamlSubjectConfirm.ValueString() != "bearer" {
		data.PpSamlSubjectConfirm = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLNameID`); value.Exists() && !data.PpSamlNameId.IsNull() {
		data.PpSamlNameId = tfutils.BoolFromString(value.String())
	} else if !data.PpSamlNameId.ValueBool() {
		data.PpSamlNameId = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPSAMLNameIDFormat`); value.Exists() && !data.PpSamlNameIdFormat.IsNull() {
		data.PpSamlNameIdFormat = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpSamlNameIdFormat = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLRecipient`); value.Exists() && !data.PpSamlRecipient.IsNull() {
		data.PpSamlRecipient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpSamlRecipient = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLAudience`); value.Exists() && !data.PpSamlAudience.IsNull() {
		data.PpSamlAudience = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpSamlAudience = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLOmitNotBefore`); value.Exists() && !data.PpSamlOmitNotBefore.IsNull() {
		data.PpSamlOmitNotBefore = tfutils.BoolFromString(value.String())
	} else if data.PpSamlOmitNotBefore.ValueBool() {
		data.PpSamlOmitNotBefore = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPOneTimeUse`); value.Exists() && !data.PpOneTimeUse.IsNull() {
		data.PpOneTimeUse = tfutils.BoolFromString(value.String())
	} else if data.PpOneTimeUse.ValueBool() {
		data.PpOneTimeUse = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPSAMLProxy`); value.Exists() && !data.PpSamlProxy.IsNull() {
		data.PpSamlProxy = tfutils.BoolFromString(value.String())
	} else if data.PpSamlProxy.ValueBool() {
		data.PpSamlProxy = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPSAMLProxyAudience`); value.Exists() && !data.PpSamlProxyAudience.IsNull() {
		data.PpSamlProxyAudience = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpSamlProxyAudience = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLProxyCount`); value.Exists() && !data.PpSamlProxyCount.IsNull() {
		data.PpSamlProxyCount = types.Int64Value(value.Int())
	} else if data.PpSamlProxyCount.ValueInt64() != 0 {
		data.PpSamlProxyCount = types.Int64Null()
	}
	if value := res.Get(pathRoot + `PPSAMLAuthzAction`); value.Exists() && !data.PpSamlAuthzAction.IsNull() {
		data.PpSamlAuthzAction = tfutils.ParseStringFromGJSON(value)
	} else if data.PpSamlAuthzAction.ValueString() != "AllHTTP" {
		data.PpSamlAuthzAction = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSAMLAttributes`); value.Exists() && !data.PpSamlAttributes.IsNull() {
		data.PpSamlAttributes = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpSamlAttributes = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPLTPAInsertCookie`); value.Exists() && !data.PpLtpaInsertCookie.IsNull() {
		data.PpLtpaInsertCookie = tfutils.BoolFromString(value.String())
	} else if !data.PpLtpaInsertCookie.ValueBool() {
		data.PpLtpaInsertCookie = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPTAMPACPropagate`); value.Exists() && !data.PpTamPacPropagate.IsNull() {
		data.PpTamPacPropagate = tfutils.BoolFromString(value.String())
	} else if data.PpTamPacPropagate.ValueBool() {
		data.PpTamPacPropagate = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPTAMHeader`); value.Exists() && !data.PpTamHeader.IsNull() {
		data.PpTamHeader = tfutils.ParseStringFromGJSON(value)
	} else if data.PpTamHeader.ValueString() != "iv-creds" {
		data.PpTamHeader = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPTAMHeaderSize`); value.Exists() && !data.PpTamHeaderSize.IsNull() {
		data.PpTamHeaderSize = types.Int64Value(value.Int())
	} else if data.PpTamHeaderSize.ValueInt64() != 0 {
		data.PpTamHeaderSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `PPKerberosUseS4U2Proxy`); value.Exists() && !data.PpKerberosUseS4u2proxy.IsNull() {
		data.PpKerberosUseS4u2proxy = tfutils.BoolFromString(value.String())
	} else if data.PpKerberosUseS4u2proxy.ValueBool() {
		data.PpKerberosUseS4u2proxy = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPCookieAttributes`); value.Exists() && !data.PpCookieAttributes.IsNull() {
		data.PpCookieAttributes = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpCookieAttributes = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPKerberosUseS4U2SelfAndS4U2Proxy`); value.Exists() && !data.PpKerberosUseS4u2selfAndS4u2proxy.IsNull() {
		data.PpKerberosUseS4u2selfAndS4u2proxy = tfutils.BoolFromString(value.String())
	} else if data.PpKerberosUseS4u2selfAndS4u2proxy.ValueBool() {
		data.PpKerberosUseS4u2selfAndS4u2proxy = types.BoolNull()
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
	if value := res.Get(pathRoot + `PPSSLClientConfigType`); value.Exists() && !data.PpSslClientConfigType.IsNull() {
		data.PpSslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else if data.PpSslClientConfigType.ValueString() != "proxy" {
		data.PpSslClientConfigType = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPSSLClientProfile`); value.Exists() && !data.PpSslClientProfile.IsNull() {
		data.PpSslClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpSslClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPLTPAKeyFilePasswordAlias`); value.Exists() && !data.PpLtpaKeyFilePasswordAlias.IsNull() {
		data.PpLtpaKeyFilePasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpLtpaKeyFilePasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `PPJWT`); value.Exists() && !data.PpJwt.IsNull() {
		data.PpJwt = tfutils.BoolFromString(value.String())
	} else if data.PpJwt.ValueBool() {
		data.PpJwt = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PPJWTGenerator`); value.Exists() && !data.PpJwtGenerator.IsNull() {
		data.PpJwtGenerator = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PpJwtGenerator = types.StringNull()
	}
}
