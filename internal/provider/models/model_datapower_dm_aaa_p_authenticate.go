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

type DmAAAPAuthenticate struct {
	AuMethod                        types.String        `tfsdk:"au_method"`
	AuCustomUrl                     types.String        `tfsdk:"au_custom_url"`
	AuMapUrl                        types.String        `tfsdk:"au_map_url"`
	AuHost                          types.String        `tfsdk:"au_host"`
	AuPort                          types.Int64         `tfsdk:"au_port"`
	AuSslValcred                    types.String        `tfsdk:"au_ssl_valcred"`
	AuCacheAllow                    types.String        `tfsdk:"au_cache_allow"`
	AuCacheTtl                      types.Int64         `tfsdk:"au_cache_ttl"`
	AuKerberosPrincipal             types.String        `tfsdk:"au_kerberos_principal"`
	AuKerberosPassword              types.String        `tfsdk:"au_kerberos_password"`
	AuClearTrustServerUrl           types.String        `tfsdk:"au_clear_trust_server_url"`
	AuClearTrustApplication         types.String        `tfsdk:"au_clear_trust_application"`
	AuSamlArtifactResponder         types.String        `tfsdk:"au_saml_artifact_responder"`
	AuKerberosVerifySignature       types.Bool          `tfsdk:"au_kerberos_verify_signature"`
	AuNetegrityBaseUri              types.String        `tfsdk:"au_netegrity_base_uri"`
	AuSamlAuthQueryServer           types.String        `tfsdk:"au_saml_auth_query_server"`
	AuSamlVersion                   types.String        `tfsdk:"au_saml_version"`
	AuLdapPrefix                    types.String        `tfsdk:"au_ldap_prefix"`
	AuLdapSuffix                    types.String        `tfsdk:"au_ldap_suffix"`
	AuLdapLoadBalanceGroup          types.String        `tfsdk:"au_ldap_load_balance_group"`
	AuKerberosKeytab                types.String        `tfsdk:"au_kerberos_keytab"`
	AuWsTrustUrl                    types.String        `tfsdk:"au_ws_trust_url"`
	AuSaml2Issuer                   types.String        `tfsdk:"au_saml2_issuer"`
	AuSignerValcred                 types.String        `tfsdk:"au_signer_valcred"`
	AuSignedXpath                   types.String        `tfsdk:"au_signed_xpath"`
	AuLdapBindDn                    types.String        `tfsdk:"au_ldap_bind_dn"`
	AuLdapSearchAttribute           types.String        `tfsdk:"au_ldap_search_attribute"`
	AuLtpaTokenVersionsBitmap       *DmLTPATokenVersion `tfsdk:"au_ltpa_token_versions_bitmap"`
	AuLtpaKeyFile                   types.String        `tfsdk:"au_ltpa_key_file"`
	AuBinaryTokenX509Valcred        types.String        `tfsdk:"au_binary_token_x509_valcred"`
	AuTamServer                     types.String        `tfsdk:"au_tam_server"`
	AuAllowRemoteTokenReference     types.Bool          `tfsdk:"au_allow_remote_token_reference"`
	AuRemoteTokenProcessService     types.String        `tfsdk:"au_remote_token_process_service"`
	AuWsTrustVersion                types.String        `tfsdk:"au_ws_trust_version"`
	AuLdapSearchForDn               types.Bool          `tfsdk:"au_ldap_search_for_dn"`
	AuLdapSearchParameters          types.String        `tfsdk:"au_ldap_search_parameters"`
	AuWsTrustRequireClientEntropy   types.Bool          `tfsdk:"au_ws_trust_require_client_entropy"`
	AuWsTrustClientEntropySize      types.Int64         `tfsdk:"au_ws_trust_client_entropy_size"`
	AuWsTrustRequireServerEntropy   types.Bool          `tfsdk:"au_ws_trust_require_server_entropy"`
	AuWsTrustServerEntropySize      types.Int64         `tfsdk:"au_ws_trust_server_entropy_size"`
	AuWsTrustRequireRstc            types.Bool          `tfsdk:"au_ws_trust_require_rstc"`
	AuWsTrustRequireAppliesToHeader types.Bool          `tfsdk:"au_ws_trust_require_applies_to_header"`
	AuWsTrustAppliesToHeader        types.String        `tfsdk:"au_ws_trust_applies_to_header"`
	AuWsTrustEncryptionCertificate  types.String        `tfsdk:"au_ws_trust_encryption_certificate"`
	AuZosNssConfig                  types.String        `tfsdk:"au_zos_nss_config"`
	AuLdapAttributes                types.String        `tfsdk:"au_ldap_attributes"`
	AuSkewTime                      types.Int64         `tfsdk:"au_skew_time"`
	AuTamPacReturn                  types.Bool          `tfsdk:"au_tam_pac_return"`
	AuLdapReadTimeout               types.Int64         `tfsdk:"au_ldap_read_timeout"`
	AuSslClientConfigType           types.String        `tfsdk:"au_ssl_client_config_type"`
	AuSslClientProfile              types.String        `tfsdk:"au_ssl_client_profile"`
	AuLdapBindPasswordAlias         types.String        `tfsdk:"au_ldap_bind_password_alias"`
	AuLtpaKeyFilePasswordAlias      types.String        `tfsdk:"au_ltpa_key_file_password_alias"`
	AuSmRequestType                 types.String        `tfsdk:"au_sm_request_type"`
	AuSmCookieFlow                  *DmSMFlow           `tfsdk:"au_sm_cookie_flow"`
	AuSmHeaderFlow                  *DmSMFlow           `tfsdk:"au_sm_header_flow"`
	AuSmCookieAttributes            types.String        `tfsdk:"au_sm_cookie_attributes"`
	AuCacheControl                  types.String        `tfsdk:"au_cache_control"`
}

var DmAAAPAuthenticateAUCustomURLCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"custom"},
}

var DmAAAPAuthenticateAUCustomURLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthenticateAUMapURLCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"xmlfile"},
}

var DmAAAPAuthenticateAUMapURLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthenticateAUHostCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "ldap",
			Value:       []string{"ldap", "oblix", "netegrity"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "au_ldap_load_balance_group",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{""},
		},
	},
}

var DmAAAPAuthenticateAUHostIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "ldap",
			Value:       []string{"ldap", "netegrity", "oblix"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "au_ldap_load_balance_group",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{""},
		},
	},
}

var DmAAAPAuthenticateAUPortCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "ldap",
			Value:       []string{"ldap", "oblix", "netegrity"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "au_ldap_load_balance_group",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{""},
		},
	},
}

var DmAAAPAuthenticateAUPortIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "ldap",
			Value:       []string{"ldap", "netegrity", "oblix"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "au_ldap_load_balance_group",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{""},
		},
	},
}

var DmAAAPAuthenticateAUSSLValcredIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"client-ssl"},
}

var DmAAAPAuthenticateAUCacheTTLIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "au_cache_allow",
	AttrType:    "String",
	AttrDefault: "absolute",
	Value:       []string{"disabled"},
}

var DmAAAPAuthenticateAUKerberosPrincipalIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"kerberos"},
}

var DmAAAPAuthenticateAUKerberosPasswordIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthenticateAUClearTrustServerURLCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"cleartrust"},
}

var DmAAAPAuthenticateAUClearTrustServerURLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthenticateAUClearTrustApplicationIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthenticateAUSAMLArtifactResponderIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"saml-artifact"},
}

var DmAAAPAuthenticateAUKerberosVerifySignatureIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthenticateAUNetegrityBaseURIIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"netegrity"},
}

var DmAAAPAuthenticateAUSAMLAuthQueryServerIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"saml-authen-query"},
}

var DmAAAPAuthenticateAUSAMLVersionCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"saml-artifact", "saml-authen-query"},
}

var DmAAAPAuthenticateAUSAMLVersionIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthenticateAULDAPPrefixIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "ldap",
			Value:       []string{"ldap"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "au_ldap_search_for_dn",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
	},
}

var DmAAAPAuthenticateAULDAPSuffixIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "ldap",
			Value:       []string{"ldap"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "au_ldap_search_for_dn",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
	},
}

var DmAAAPAuthenticateAULDAPLoadBalanceGroupIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"ldap"},
}

var DmAAAPAuthenticateAUKerberosKeytabCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"kerberos"},
}

var DmAAAPAuthenticateAUKerberosKeytabIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthenticateAUWSTrustURLCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"ws-trust"},
}

var DmAAAPAuthenticateAUWSTrustURLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthenticateAUSAML2IssuerIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthenticateAUSignerValcredIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"validate-signer"},
}

var DmAAAPAuthenticateAUSignedXPathIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"validate-signer"},
}

var DmAAAPAuthenticateAULDAPBindDNIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"ldap"},
}

var DmAAAPAuthenticateAULDAPSearchAttributeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"ldap"},
}

var DmAAAPAuthenticateAULTPATokenVersionsBitmapCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"ltpa"},
}

var DmAAAPAuthenticateAULTPATokenVersionsBitmapIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthenticateAULTPAKeyFileCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"ltpa"},
}

var DmAAAPAuthenticateAULTPAKeyFileIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthenticateAUBinaryTokenX509ValcredCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"binarytokenx509"},
}

var DmAAAPAuthenticateAUBinaryTokenX509ValcredIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthenticateAUTAMServerCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"tivoli"},
}

var DmAAAPAuthenticateAUTAMServerIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthenticateAUAllowRemoteTokenReferenceIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"saml-signature", "validate-signer"},
}

var DmAAAPAuthenticateAURemoteTokenProcessServiceIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "ldap",
			Value:       []string{"saml-signature", "validate-signer"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "au_allow_remote_token_reference",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
	},
}

var DmAAAPAuthenticateAUWSTrustVersionIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"ws-trust"},
}

var DmAAAPAuthenticateAULDAPSearchForDNIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"ldap"},
}

var DmAAAPAuthenticateAULDAPSearchParametersCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "ldap",
			Value:       []string{"ldap"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "au_ldap_search_for_dn",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
	},
}

var DmAAAPAuthenticateAULDAPSearchParametersIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "ldap",
			Value:       []string{"ldap"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "au_ldap_search_for_dn",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
	},
}

var DmAAAPAuthenticateAUWSTrustRequireClientEntropyIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"ws-trust"},
}

var DmAAAPAuthenticateAUWSTrustClientEntropySizeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "ldap",
			Value:       []string{"ws-trust"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "au_ws_trust_require_client_entropy",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
	},
}

var DmAAAPAuthenticateAUWSTrustRequireServerEntropyIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"ws-trust"},
}

var DmAAAPAuthenticateAUWSTrustServerEntropySizeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "ldap",
			Value:       []string{"ws-trust"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "au_ws_trust_require_server_entropy",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
	},
}

var DmAAAPAuthenticateAUWSTrustRequireRSTCIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"ws-trust"},
}

var DmAAAPAuthenticateAUWSTrustRequireAppliesToHeaderIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"ws-trust"},
}

var DmAAAPAuthenticateAUWSTrustAppliesToHeaderIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "ldap",
			Value:       []string{"ws-trust"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "au_ws_trust_require_applies_to_header",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
	},
}

var DmAAAPAuthenticateAUWSTrustEncryptionCertificateIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"ws-trust"},
}

var DmAAAPAuthenticateAUZOSNSSConfigCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"zosnss"},
}

var DmAAAPAuthenticateAUZOSNSSConfigIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthenticateAULDAPAttributesIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"ldap"},
}

var DmAAAPAuthenticateAUSkewTimeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"saml-artifact", "saml-authen-query", "saml-signature"},
}

var DmAAAPAuthenticateAUTAMPACReturnCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"tivoli"},
}

var DmAAAPAuthenticateAUTAMPACReturnIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthenticateAULDAPReadTimeoutIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"ldap"},
}

var DmAAAPAuthenticateAUSSLClientConfigTypeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "ldap",
			Value:       []string{"cleartrust", "ldap", "netegrity", "saml-artifact", "saml-authen-query", "ws-trust"},
		},
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-not-in-list",
					Attribute:   "au_method",
					AttrType:    "String",
					AttrDefault: "ldap",
					Value:       []string{"saml-signature", "validate-signer"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "au_allow_remote_token_reference",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"false"},
				},
			},
		},
	},
}

var DmAAAPAuthenticateAUSSLClientProfileIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-not-in-list",
					Attribute:   "au_method",
					AttrType:    "String",
					AttrDefault: "ldap",
					Value:       []string{"cleartrust", "ldap", "netegrity", "saml-artifact", "saml-authen-query", "ws-trust"},
				},

				{
					Evaluation: "logical-or",
					Conditions: []validators.Evaluation{

						{
							Evaluation:  "property-value-not-in-list",
							Attribute:   "au_method",
							AttrType:    "String",
							AttrDefault: "ldap",
							Value:       []string{"saml-signature", "validate-signer"},
						},

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "au_allow_remote_token_reference",
							AttrType:    "Bool",
							AttrDefault: "false",
							Value:       []string{"false"},
						},
					},
				},
			},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "au_ssl_client_config_type",
			AttrType:    "String",
			AttrDefault: "client",
			Value:       []string{"client"},
		},
	},
}

var DmAAAPAuthenticateAULDAPBindPasswordAliasIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"ldap"},
}

var DmAAAPAuthenticateAULTPAKeyFilePasswordAliasCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "ldap",
			Value:       []string{"ltpa"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "au_ltpa_token_versions_bitmap",
			AttrType:    "DmLTPATokenVersion",
			AttrDefault: "LTPA2",
			Value:       []string{"LTPA", "LTPA2"},
		},
	},
}

var DmAAAPAuthenticateAULTPAKeyFilePasswordAliasIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "ldap",
			Value:       []string{"ltpa"},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-not-in-list",
					Attribute:   "au_ltpa_token_versions_bitmap",
					AttrType:    "DmLTPATokenVersion",
					AttrDefault: "LTPA2",
					Value:       []string{"LTPA", "LTPA2"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "au_ltpa_token_versions_bitmap",
					AttrType:    "DmLTPATokenVersion",
					AttrDefault: "LTPA2",
					Value:       []string{"LTPADomino"},
				},
			},
		},
	},
}

var DmAAAPAuthenticateAUSMRequestTypeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"netegrity"},
}

var DmAAAPAuthenticateAUSMRequestTypeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"netegrity"},
}

var DmAAAPAuthenticateAUSMCookieFlowIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"netegrity"},
}

var DmAAAPAuthenticateAUSMHeaderFlowIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "ldap",
	Value:       []string{"netegrity"},
}

var DmAAAPAuthenticateAUSMCookieAttributesIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_sm_cookie_flow",
	AttrType:    "DmSMFlow",
	AttrDefault: "",
	Value:       []string{"frontend", "frontend+backend", "backend+frontend"},
}

var DmAAAPAuthenticateObjectType = map[string]attr.Type{
	"au_method":                             types.StringType,
	"au_custom_url":                         types.StringType,
	"au_map_url":                            types.StringType,
	"au_host":                               types.StringType,
	"au_port":                               types.Int64Type,
	"au_ssl_valcred":                        types.StringType,
	"au_cache_allow":                        types.StringType,
	"au_cache_ttl":                          types.Int64Type,
	"au_kerberos_principal":                 types.StringType,
	"au_kerberos_password":                  types.StringType,
	"au_clear_trust_server_url":             types.StringType,
	"au_clear_trust_application":            types.StringType,
	"au_saml_artifact_responder":            types.StringType,
	"au_kerberos_verify_signature":          types.BoolType,
	"au_netegrity_base_uri":                 types.StringType,
	"au_saml_auth_query_server":             types.StringType,
	"au_saml_version":                       types.StringType,
	"au_ldap_prefix":                        types.StringType,
	"au_ldap_suffix":                        types.StringType,
	"au_ldap_load_balance_group":            types.StringType,
	"au_kerberos_keytab":                    types.StringType,
	"au_ws_trust_url":                       types.StringType,
	"au_saml2_issuer":                       types.StringType,
	"au_signer_valcred":                     types.StringType,
	"au_signed_xpath":                       types.StringType,
	"au_ldap_bind_dn":                       types.StringType,
	"au_ldap_search_attribute":              types.StringType,
	"au_ltpa_token_versions_bitmap":         types.ObjectType{AttrTypes: DmLTPATokenVersionObjectType},
	"au_ltpa_key_file":                      types.StringType,
	"au_binary_token_x509_valcred":          types.StringType,
	"au_tam_server":                         types.StringType,
	"au_allow_remote_token_reference":       types.BoolType,
	"au_remote_token_process_service":       types.StringType,
	"au_ws_trust_version":                   types.StringType,
	"au_ldap_search_for_dn":                 types.BoolType,
	"au_ldap_search_parameters":             types.StringType,
	"au_ws_trust_require_client_entropy":    types.BoolType,
	"au_ws_trust_client_entropy_size":       types.Int64Type,
	"au_ws_trust_require_server_entropy":    types.BoolType,
	"au_ws_trust_server_entropy_size":       types.Int64Type,
	"au_ws_trust_require_rstc":              types.BoolType,
	"au_ws_trust_require_applies_to_header": types.BoolType,
	"au_ws_trust_applies_to_header":         types.StringType,
	"au_ws_trust_encryption_certificate":    types.StringType,
	"au_zos_nss_config":                     types.StringType,
	"au_ldap_attributes":                    types.StringType,
	"au_skew_time":                          types.Int64Type,
	"au_tam_pac_return":                     types.BoolType,
	"au_ldap_read_timeout":                  types.Int64Type,
	"au_ssl_client_config_type":             types.StringType,
	"au_ssl_client_profile":                 types.StringType,
	"au_ldap_bind_password_alias":           types.StringType,
	"au_ltpa_key_file_password_alias":       types.StringType,
	"au_sm_request_type":                    types.StringType,
	"au_sm_cookie_flow":                     types.ObjectType{AttrTypes: DmSMFlowObjectType},
	"au_sm_header_flow":                     types.ObjectType{AttrTypes: DmSMFlowObjectType},
	"au_sm_cookie_attributes":               types.StringType,
	"au_cache_control":                      types.StringType,
}
var DmAAAPAuthenticateObjectDefault = map[string]attr.Value{
	"au_method":                             types.StringValue("ldap"),
	"au_custom_url":                         types.StringNull(),
	"au_map_url":                            types.StringNull(),
	"au_host":                               types.StringNull(),
	"au_port":                               types.Int64Value(389),
	"au_ssl_valcred":                        types.StringNull(),
	"au_cache_allow":                        types.StringValue("absolute"),
	"au_cache_ttl":                          types.Int64Value(3),
	"au_kerberos_principal":                 types.StringNull(),
	"au_kerberos_password":                  types.StringNull(),
	"au_clear_trust_server_url":             types.StringNull(),
	"au_clear_trust_application":            types.StringNull(),
	"au_saml_artifact_responder":            types.StringNull(),
	"au_kerberos_verify_signature":          types.BoolValue(true),
	"au_netegrity_base_uri":                 types.StringNull(),
	"au_saml_auth_query_server":             types.StringNull(),
	"au_saml_version":                       types.StringValue("1.1"),
	"au_ldap_prefix":                        types.StringValue("cn="),
	"au_ldap_suffix":                        types.StringNull(),
	"au_ldap_load_balance_group":            types.StringNull(),
	"au_kerberos_keytab":                    types.StringNull(),
	"au_ws_trust_url":                       types.StringNull(),
	"au_saml2_issuer":                       types.StringNull(),
	"au_signer_valcred":                     types.StringNull(),
	"au_signed_xpath":                       types.StringNull(),
	"au_ldap_bind_dn":                       types.StringNull(),
	"au_ldap_search_attribute":              types.StringValue("userPassword"),
	"au_ltpa_token_versions_bitmap":         types.ObjectValueMust(DmLTPATokenVersionObjectType, DmLTPATokenVersionObjectDefault),
	"au_ltpa_key_file":                      types.StringNull(),
	"au_binary_token_x509_valcred":          types.StringNull(),
	"au_tam_server":                         types.StringNull(),
	"au_allow_remote_token_reference":       types.BoolValue(false),
	"au_remote_token_process_service":       types.StringNull(),
	"au_ws_trust_version":                   types.StringValue("1.2"),
	"au_ldap_search_for_dn":                 types.BoolValue(false),
	"au_ldap_search_parameters":             types.StringNull(),
	"au_ws_trust_require_client_entropy":    types.BoolValue(false),
	"au_ws_trust_client_entropy_size":       types.Int64Value(32),
	"au_ws_trust_require_server_entropy":    types.BoolValue(false),
	"au_ws_trust_server_entropy_size":       types.Int64Value(32),
	"au_ws_trust_require_rstc":              types.BoolValue(false),
	"au_ws_trust_require_applies_to_header": types.BoolValue(false),
	"au_ws_trust_applies_to_header":         types.StringNull(),
	"au_ws_trust_encryption_certificate":    types.StringNull(),
	"au_zos_nss_config":                     types.StringNull(),
	"au_ldap_attributes":                    types.StringNull(),
	"au_skew_time":                          types.Int64Value(0),
	"au_tam_pac_return":                     types.BoolValue(false),
	"au_ldap_read_timeout":                  types.Int64Value(60),
	"au_ssl_client_config_type":             types.StringValue("client"),
	"au_ssl_client_profile":                 types.StringNull(),
	"au_ldap_bind_password_alias":           types.StringNull(),
	"au_ltpa_key_file_password_alias":       types.StringNull(),
	"au_sm_request_type":                    types.StringValue("webagent"),
	"au_sm_cookie_flow":                     types.ObjectValueMust(DmSMFlowObjectType, DmSMFlowObjectDefault),
	"au_sm_header_flow":                     types.ObjectValueMust(DmSMFlowObjectType, DmSMFlowObjectDefault),
	"au_sm_cookie_attributes":               types.StringNull(),
	"au_cache_control":                      types.StringValue("default"),
}

func GetDmAAAPAuthenticateDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmAAAPAuthenticateDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"au_method": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the method to authenticate the extracted identity.",
				Computed:            true,
			},
			"au_custom_url": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the location of the processing file. This file is the stylesheet or GatewayScript that authenticates the extracted identity.",
				Computed:            true,
			},
			"au_map_url": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the location of the AAA information file. This file contains a list of authenticated identities and the various values needed to authenticate successfully.",
				Computed:            true,
			},
			"au_host": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the host name or IP address of the authentication server.",
				Computed:            true,
			},
			"au_port": DataSourceSchema.Int64Attribute{
				MarkdownDescription: "Specify this listening port on the authentication server.",
				Computed:            true,
			},
			"au_ssl_valcred": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the validation credentials that contain the certificate to validate the remote TLS peer.",
				Computed:            true,
			},
			"au_cache_allow": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify how to control the caching of AAA authentication results. A protocol TTL is available only with SAML. The default value is absolute.",
				Computed:            true,
			},
			"au_cache_ttl": DataSourceSchema.Int64Attribute{
				MarkdownDescription: "Specify the duration in seconds to cache authentication decisions. Enter a value in the range 1 - 86400. The default value is 3.",
				Computed:            true,
			},
			"au_kerberos_principal": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the principal name that must appear as the server name in the Kerberos ticket.This value must be a full principal name, including the Kerberos realm. For example, <tt>foo/bar@REALM</tt> .",
				Computed:            true,
			},
			"au_kerberos_password": DataSourceSchema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"au_clear_trust_server_url": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the URL to access the ClearTrust server for authentication.",
				Computed:            true,
			},
			"au_clear_trust_application": DataSourceSchema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"au_saml_artifact_responder": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the URL of the SAML artifact responder.",
				Computed:            true,
			},
			"au_kerberos_verify_signature": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"au_netegrity_base_uri": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the base URI sent to CA Single Sign-On server. The base URI is combined with the host and port to form the URL for attempting CA Single Sign-On authentication. This base URI must equal the concatenation of the <tt>servlet-name</tt> and its <tt>url-pattern</tt> set in its <tt>web.xml</tt> configuration file. If the <tt>servlet-name</tt> is <tt>datapoweragent</tt> and the <tt>url-pattern</tt> is <tt>/</tt> , then the base URI must be <tt>datapoweragent/</tt> .",
				Computed:            true,
			},
			"au_saml_auth_query_server": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the URL to access the SAML authentication query server and to post a SAML authentication query.",
				Computed:            true,
			},
			"au_saml_version": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the version of SAML messages. The default value is 1.1.",
				Computed:            true,
			},
			"au_ldap_prefix": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the prefix to construct the LDAP lookup DN. The default value is <tt>cn=</tt> .",
				Computed:            true,
			},
			"au_ldap_suffix": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the suffix used to construct the LDAP lookup DN.",
				Computed:            true,
			},
			"au_ldap_load_balance_group": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the name of the load balancer group that contains the LDAP servers.",
				Computed:            true,
			},
			"au_kerberos_keytab": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the name of the keytab for the Kerberos server principal. This keytab is required to decrypt the client Kerberos ticket.",
				Computed:            true,
			},
			"au_ws_trust_url": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the URL to access the WS-Trust server.",
				Computed:            true,
			},
			"au_saml2_issuer": DataSourceSchema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"au_signer_valcred": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the name of the validation credentials to verify the signature validity for the incoming message. When validation credentials are set, the signer certificate must be contained in the validation credentials or the signature is rejected as untrusted.",
				Computed:            true,
			},
			"au_signed_xpath": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the XPath expression for the XML entity that is protected by signature. After the signature validity is verified, this property verifies if the specific XPath expression is part of the signed message.",
				Computed:            true,
			},
			"au_ldap_bind_dn": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the DN to bind to the LDAP server for an LDAP search. This value is used when the password from the extract identity phase is a WS-Security UsernameToken PasswordDigest. The LDAP server is searched for the corresponding password to verify the PasswordDigest.",
				Computed:            true,
			},
			"au_ldap_search_attribute": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the attribute to use in the LDAP search. The default value is userPassword.",
				Computed:            true,
			},
			"au_ltpa_token_versions_bitmap": GetDmLTPATokenVersionDataSourceSchema("Specify which versions of LTPA tokens are acceptable.", "lpta-version", ""),
			"au_ltpa_key_file": DataSourceSchema.StringAttribute{
				MarkdownDescription: "<p>Specify the LTPA key file that contains the crypto material to create an LTPA token that can be consumed by WebSphere (both version 1 and version 2) or Domino.</p><ul><li>For WebSphere token creation, you must export the LTPA key file from WebSphere. This file has portions encrypted by a password.</li><li>For Domino token creation, the key file contains only the base 64-encoded Domino shared secret.</li></ul>",
				Computed:            true,
			},
			"au_binary_token_x509_valcred": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the name of the validation credentials to validate the X.509 certificate in the BinarySecurityToken.",
				Computed:            true,
			},
			"au_tam_server": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the name of the Access Manager client.",
				Computed:            true,
			},
			"au_allow_remote_token_reference": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "<p>Specify whether to allow the retrieval of a remote security token. By default, retrieval is prohibited.</p><p>Examples of remote security tokens are as follows.</p><ul><li>The SAML assertion holds the signer public certificate.</li><li>The SAML assertion at which the signed Security Token Reference (STR dereference transform) points.</li></ul>",
				Computed:            true,
			},
			"au_remote_token_process_service": DataSourceSchema.StringAttribute{
				MarkdownDescription: "<p>Specify the URL for a service that can process the remote security token. This service accepts the WS-Security token as the request of the SOAP call and, if successful, provides the final security token as the response.</p><p>The remote WS-Security token can be signed, encrypted, or encoded DataPower services with different processing actions can process the remote token. Processing can be by decrypting parts of a remote SAML assertion, doing an XSLT transform, or with an AAA policy to assert the token.</p>",
				Computed:            true,
			},
			"au_ws_trust_version": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the version of the WS-Trust and the WS-SecureConversation specifications to use when WS-Trust authentication sends a request to the remote STS. Usually these specifications are updated together. The default value is 1.2.",
				Computed:            true,
			},
			"au_ldap_search_for_dn": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Specify whether to retrieve the user DN with an LDAP search. By default, the login name that the user presents is used with the LDAP prefix and LDAP suffix to construct the user DN.",
				Computed:            true,
			},
			"au_ldap_search_parameters": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the name of an LDAP search parameters that retrieves the user DN.",
				Computed:            true,
			},
			"au_ws_trust_require_client_entropy": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "<p>Specify whether to require client entropy in the WS-Trust request. When required, a WS-Trust entropy element is sent by the client as part of the security token request exchange. By default, entropy is not required.</p><ul><li>If a WS-Trust encryption certificate is used, the client entropy material is encrypted.</li><li>If a WS-Trust encryption certificate is not used, a WS-Trust <tt>BinarySecret</tt> element contains the entropy material. In this case, use a TLS profile to secure the exchange with the WS-Trust server.</li></ul>",
				Computed:            true,
			},
			"au_ws_trust_client_entropy_size": DataSourceSchema.Int64Attribute{
				MarkdownDescription: "Specify the size in bytes of the WS-Trust client entropy material. The size refers to the length of the entropy before base 64-encoding. Enter a value in the range 8 - 128. The default value is 32.",
				Computed:            true,
			},
			"au_ws_trust_require_server_entropy": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Specify whether to require server entropy in the WS-Trust response. When required, a WS-Trust <tt>entropy</tt> element must be returned to the client as part of the security token request exchange. By default, entropy is not required.",
				Computed:            true,
			},
			"au_ws_trust_server_entropy_size": DataSourceSchema.Int64Attribute{
				MarkdownDescription: "Specify the minimum size in bytes for the WS-Trust server entropy. The size refers to the length of the entropy before base 64-encoding. Enter any value in the range 8 - 128. The default value is 32.",
				Computed:            true,
			},
			"au_ws_trust_require_rstc": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Specify whether the message exchange with the client requires a WS-Trust <tt>RequestSecurityToken</tt> or WS-Trust <tt>RequestSecurityTokenCollection</tt> element to be sent by the client. By default, the element is not required.",
				Computed:            true,
			},
			"au_ws_trust_require_applies_to_header": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Specify whether to require a WS-Addressing <tt>AppliesTo</tt> header in the message exchange. By default, the header is not required.",
				Computed:            true,
			},
			"au_ws_trust_applies_to_header": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the value for the WS-Addressing <tt>AppliesTo</tt> header. The <tt>header</tt> element is included in the WS-Trust request security token message sent to the WS-Trust server.",
				Computed:            true,
			},
			"au_ws_trust_encryption_certificate": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the certificate to encrypt WS-Trust elements for recipient. If client entropy is configured, the certificate public key encrypts the material for the recipient. If client entropy is configured and this certificate is not specified, use a TLS profile to secure the message exchange.",
				Computed:            true,
			},
			"au_zos_nss_config": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the name of the z/OS NSS client for SAF communication.",
				Computed:            true,
			},
			"au_ldap_attributes": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the list of the extra user attributes to retrieve from the LDAP user store and kept in a <tt>var://context/ldap/auxiliary-attributes</tt> context variable for future use, such as AAA postprocessing. To define the list of LDAP attributes as the auxiliary information for AAA, use the comma (,) as the delimiter. For example, <tt>email, cn, userPassword</tt> .",
				Computed:            true,
			},
			"au_skew_time": DataSourceSchema.Int64Attribute{
				MarkdownDescription: "Specify the skew time in seconds. The skew time is the difference between the clock time on the DataPower Gateway and the time on other systems. The default value is 0. <p>When defined, the expiration of the SAML assertion takes the time difference into account.</p><ul><li>For <tt>NotBefore</tt> , validates with <tt>CurrentTime</tt> minus <tt>SkewTime</tt> .</li><li>For <tt>NotOnOrAfter</tt> , validates with <tt>CurrentTime</tt> plus <tt>SkewTime</tt> .</li></ul>",
				Computed:            true,
			},
			"au_tam_pac_return": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Specify whether to return the Access Manager privilege attribute certificate (PAC) token from a successful authentication for further use. You can use the PAC in the postprocessing phase. By default, The default the PAC token is not returned.",
				Computed:            true,
			},
			"au_ldap_read_timeout": DataSourceSchema.Int64Attribute{
				MarkdownDescription: "Specify the duration in seconds to wait for a response from the LDAP server before the LDAP connection is closed. Enter a value in the range 0 - 86400. The default value is 60. A value of 0 indicates that the connection never times out. <p>If you configure an LDAP connection pool and assign it to the XML manager, the service can use this LDAP connection pool. The LDAP read timer can work with the idle timer of the LDAP connection pool to remove idle connections from the LDAP connection pool.</p>",
				Computed:            true,
			},
			"au_ssl_client_config_type": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the TLS profile type to secure connections.",
				Computed:            true,
			},
			"au_ssl_client_profile": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the TLS client profile to secure connections.",
				Computed:            true,
			},
			"au_ldap_bind_password_alias": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the alias for the password to bind to the LDAP server for the LDAP search. This value is used when the password from the extract identity phase is a WS-Security UsernameToken PasswordDigest. The LDAP server is searched for the corresponding password to verify the PasswordDigest.",
				Computed:            true,
			},
			"au_ltpa_key_file_password_alias": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the password alias of the password that decrypts the LTPA key file. This password decrypts certain entries in a WebSphere LTPA key file. This password is not applicable to Domino key files.",
				Computed:            true,
			},
			"au_sm_request_type": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the type of request to make. By default, the request is against the CA Single Sign-On web agent.",
				Computed:            true,
			},
			"au_sm_cookie_flow": GetDmSMFlowDataSourceSchema("Specify which flow to include the authentication session cookie.", "sm-cookie-flow", ""),
			"au_sm_header_flow": GetDmSMFlowDataSourceSchema("Identifies the flow to include the CA Single Sign-On headers that are generated during authentication. The CA Single Sign-On HTTP headers start with <tt>SM_</tt> .", "sm-header-flow", ""),
			"au_sm_cookie_attributes": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the cookie attribute policy that allows predefined or custom attributes to be included in CA Single Sign-On cookies.",
				Computed:            true,
			},
			"au_cache_control": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify how to manage the caching of failures. By default, all failures are cached.",
				Computed:            true,
			},
		},
	}
	DmAAAPAuthenticateDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmAAAPAuthenticateDataSourceSchema
}
func GetDmAAAPAuthenticateResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmAAAPAuthenticateResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmAAAPAuthenticateObjectType,
				DmAAAPAuthenticateObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"au_method": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the method to authenticate the extracted identity.", "method", "").AddStringEnum("xmlfile", "ldap", "tivoli", "netegrity", "oblix", "cleartrust", "radius", "client-ssl", "validate-signer", "saml-signature", "saml-artifact", "saml-authen-query", "ws-trust", "ws-secureconversation", "token", "kerberos", "ltpa", "binarytokenx509", "zosnss", "verified-oauth", "custom").AddDefaultValue("ldap").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("xmlfile", "ldap", "tivoli", "netegrity", "oblix", "cleartrust", "radius", "client-ssl", "validate-signer", "saml-signature", "saml-artifact", "saml-authen-query", "ws-trust", "ws-secureconversation", "token", "kerberos", "ltpa", "binarytokenx509", "zosnss", "verified-oauth", "custom"),
				},
				Default: stringdefault.StaticString("ldap"),
			},
			"au_custom_url": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the processing file. This file is the stylesheet or GatewayScript that authenticates the extracted identity.", "custom-url", "").AddRequiredWhen(DmAAAPAuthenticateAUCustomURLCondVal.String()).AddNotValidWhen(DmAAAPAuthenticateAUCustomURLIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthenticateAUCustomURLCondVal, DmAAAPAuthenticateAUCustomURLIgnoreVal, false),
				},
			},
			"au_map_url": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the AAA information file. This file contains a list of authenticated identities and the various values needed to authenticate successfully.", "xmlfile-url", "").AddRequiredWhen(DmAAAPAuthenticateAUMapURLCondVal.String()).AddNotValidWhen(DmAAAPAuthenticateAUMapURLIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthenticateAUMapURLCondVal, DmAAAPAuthenticateAUMapURLIgnoreVal, false),
				},
			},
			"au_host": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the host name or IP address of the authentication server.", "remote-host", "").AddRequiredWhen(DmAAAPAuthenticateAUHostCondVal.String()).AddNotValidWhen(DmAAAPAuthenticateAUHostIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthenticateAUHostCondVal, DmAAAPAuthenticateAUHostIgnoreVal, false),
				},
			},
			"au_port": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify this listening port on the authentication server.", "remote-port", "").AddDefaultValue("389").AddRequiredWhen(DmAAAPAuthenticateAUPortCondVal.String()).AddNotValidWhen(DmAAAPAuthenticateAUPortIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					validators.ConditionalRequiredInt64(DmAAAPAuthenticateAUPortCondVal, DmAAAPAuthenticateAUPortIgnoreVal, true),
				},
				Default: int64default.StaticInt64(389),
			},
			"au_ssl_valcred": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the validation credentials that contain the certificate to validate the remote TLS peer.", "valcred", "crypto_val_cred").AddNotValidWhen(DmAAAPAuthenticateAUSSLValcredIgnoreVal.String()).String,
				Optional:            true,
			},
			"au_cache_allow": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how to control the caching of AAA authentication results. A protocol TTL is available only with SAML. The default value is absolute.", "cache-type", "").AddStringEnum("absolute", "disabled", "maximum", "minimum").AddDefaultValue("absolute").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("absolute", "disabled", "maximum", "minimum"),
				},
				Default: stringdefault.StaticString("absolute"),
			},
			"au_cache_ttl": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration in seconds to cache authentication decisions. Enter a value in the range 1 - 86400. The default value is 3.", "cache-ttl", "").AddIntegerRange(1, 86400).AddDefaultValue("3").AddNotValidWhen(DmAAAPAuthenticateAUCacheTTLIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 86400),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, DmAAAPAuthenticateAUCacheTTLIgnoreVal, true),
				},
				Default: int64default.StaticInt64(3),
			},
			"au_kerberos_principal": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the principal name that must appear as the server name in the Kerberos ticket.This value must be a full principal name, including the Kerberos realm. For example, <tt>foo/bar@REALM</tt> .", "kerberos-principal", "").AddNotValidWhen(DmAAAPAuthenticateAUKerberosPrincipalIgnoreVal.String()).String,
				Optional:            true,
			},
			"au_kerberos_password": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "", "").AddNotValidWhen(DmAAAPAuthenticateAUKerberosPasswordIgnoreVal.String()).String,
				Optional:            true,
			},
			"au_clear_trust_server_url": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL to access the ClearTrust server for authentication.", "cleartrust-url", "").AddRequiredWhen(DmAAAPAuthenticateAUClearTrustServerURLCondVal.String()).AddNotValidWhen(DmAAAPAuthenticateAUClearTrustServerURLIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthenticateAUClearTrustServerURLCondVal, DmAAAPAuthenticateAUClearTrustServerURLIgnoreVal, false),
				},
			},
			"au_clear_trust_application": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "", "").AddNotValidWhen(DmAAAPAuthenticateAUClearTrustApplicationIgnoreVal.String()).String,
				Optional:            true,
			},
			"au_saml_artifact_responder": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL of the SAML artifact responder.", "saml-artifact-responder", "").AddNotValidWhen(DmAAAPAuthenticateAUSAMLArtifactResponderIgnoreVal.String()).String,
				Optional:            true,
			},
			"au_kerberos_verify_signature": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "", "").AddDefaultValue("true").AddNotValidWhen(DmAAAPAuthenticateAUKerberosVerifySignatureIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"au_netegrity_base_uri": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the base URI sent to CA Single Sign-On server. The base URI is combined with the host and port to form the URL for attempting CA Single Sign-On authentication. This base URI must equal the concatenation of the <tt>servlet-name</tt> and its <tt>url-pattern</tt> set in its <tt>web.xml</tt> configuration file. If the <tt>servlet-name</tt> is <tt>datapoweragent</tt> and the <tt>url-pattern</tt> is <tt>/</tt> , then the base URI must be <tt>datapoweragent/</tt> .", "netegrity-base-uri", "").AddNotValidWhen(DmAAAPAuthenticateAUNetegrityBaseURIIgnoreVal.String()).String,
				Optional:            true,
			},
			"au_saml_auth_query_server": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL to access the SAML authentication query server and to post a SAML authentication query.", "saml-authen-query-url", "").AddNotValidWhen(DmAAAPAuthenticateAUSAMLAuthQueryServerIgnoreVal.String()).String,
				Optional:            true,
			},
			"au_saml_version": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the version of SAML messages. The default value is 1.1.", "saml-version", "").AddStringEnum("2.0", "1.1", "1.0").AddDefaultValue("1.1").AddRequiredWhen(DmAAAPAuthenticateAUSAMLVersionCondVal.String()).AddNotValidWhen(DmAAAPAuthenticateAUSAMLVersionIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("2.0", "1.1", "1.0"),
					validators.ConditionalRequiredString(DmAAAPAuthenticateAUSAMLVersionCondVal, DmAAAPAuthenticateAUSAMLVersionIgnoreVal, true),
				},
				Default: stringdefault.StaticString("1.1"),
			},
			"au_ldap_prefix": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the prefix to construct the LDAP lookup DN. The default value is <tt>cn=</tt> .", "ldap-prefix", "").AddDefaultValue("cn=").AddNotValidWhen(DmAAAPAuthenticateAULDAPPrefixIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("cn="),
			},
			"au_ldap_suffix": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the suffix used to construct the LDAP lookup DN.", "ldap-suffix", "").AddNotValidWhen(DmAAAPAuthenticateAULDAPSuffixIgnoreVal.String()).String,
				Optional:            true,
			},
			"au_ldap_load_balance_group": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the load balancer group that contains the LDAP servers.", "ldap-lbgroup", "load_balancer_group").AddNotValidWhen(DmAAAPAuthenticateAULDAPLoadBalanceGroupIgnoreVal.String()).String,
				Optional:            true,
			},
			"au_kerberos_keytab": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the keytab for the Kerberos server principal. This keytab is required to decrypt the client Kerberos ticket.", "kerberos-keytab", "crypto_kerberos_keytab").AddRequiredWhen(DmAAAPAuthenticateAUKerberosKeytabCondVal.String()).AddNotValidWhen(DmAAAPAuthenticateAUKerberosKeytabIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthenticateAUKerberosKeytabCondVal, DmAAAPAuthenticateAUKerberosKeytabIgnoreVal, false),
				},
			},
			"au_ws_trust_url": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL to access the WS-Trust server.", "ws-trust-url", "").AddRequiredWhen(DmAAAPAuthenticateAUWSTrustURLCondVal.String()).AddNotValidWhen(DmAAAPAuthenticateAUWSTrustURLIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthenticateAUWSTrustURLCondVal, DmAAAPAuthenticateAUWSTrustURLIgnoreVal, false),
				},
			},
			"au_saml2_issuer": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "", "").AddNotValidWhen(DmAAAPAuthenticateAUSAML2IssuerIgnoreVal.String()).String,
				Optional:            true,
			},
			"au_signer_valcred": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the validation credentials to verify the signature validity for the incoming message. When validation credentials are set, the signer certificate must be contained in the validation credentials or the signature is rejected as untrusted.", "valcred", "crypto_val_cred").AddNotValidWhen(DmAAAPAuthenticateAUSignerValcredIgnoreVal.String()).String,
				Optional:            true,
			},
			"au_signed_xpath": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the XPath expression for the XML entity that is protected by signature. After the signature validity is verified, this property verifies if the specific XPath expression is part of the signed message.", "signed-xpath", "").AddNotValidWhen(DmAAAPAuthenticateAUSignedXPathIgnoreVal.String()).String,
				Optional:            true,
			},
			"au_ldap_bind_dn": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the DN to bind to the LDAP server for an LDAP search. This value is used when the password from the extract identity phase is a WS-Security UsernameToken PasswordDigest. The LDAP server is searched for the corresponding password to verify the PasswordDigest.", "ldap-bind-dn", "").AddNotValidWhen(DmAAAPAuthenticateAULDAPBindDNIgnoreVal.String()).String,
				Optional:            true,
			},
			"au_ldap_search_attribute": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the attribute to use in the LDAP search. The default value is userPassword.", "ldap-search-attr", "").AddDefaultValue("userPassword").AddNotValidWhen(DmAAAPAuthenticateAULDAPSearchAttributeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("userPassword"),
			},
			"au_ltpa_token_versions_bitmap": GetDmLTPATokenVersionResourceSchema("Specify which versions of LTPA tokens are acceptable.", "lpta-version", "", false),
			"au_ltpa_key_file": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the LTPA key file that contains the crypto material to create an LTPA token that can be consumed by WebSphere (both version 1 and version 2) or Domino.</p><ul><li>For WebSphere token creation, you must export the LTPA key file from WebSphere. This file has portions encrypted by a password.</li><li>For Domino token creation, the key file contains only the base 64-encoded Domino shared secret.</li></ul>", "lpta-key-file", "").AddRequiredWhen(DmAAAPAuthenticateAULTPAKeyFileCondVal.String()).AddNotValidWhen(DmAAAPAuthenticateAULTPAKeyFileIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthenticateAULTPAKeyFileCondVal, DmAAAPAuthenticateAULTPAKeyFileIgnoreVal, false),
				},
			},
			"au_binary_token_x509_valcred": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the validation credentials to validate the X.509 certificate in the BinarySecurityToken.", "x509-bin-token-valcred", "crypto_val_cred").AddRequiredWhen(DmAAAPAuthenticateAUBinaryTokenX509ValcredCondVal.String()).AddNotValidWhen(DmAAAPAuthenticateAUBinaryTokenX509ValcredIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthenticateAUBinaryTokenX509ValcredCondVal, DmAAAPAuthenticateAUBinaryTokenX509ValcredIgnoreVal, false),
				},
			},
			"au_tam_server": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the Access Manager client.", "tam", "tam").AddRequiredWhen(DmAAAPAuthenticateAUTAMServerCondVal.String()).AddNotValidWhen(DmAAAPAuthenticateAUTAMServerIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthenticateAUTAMServerCondVal, DmAAAPAuthenticateAUTAMServerIgnoreVal, false),
				},
			},
			"au_allow_remote_token_reference": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to allow the retrieval of a remote security token. By default, retrieval is prohibited.</p><p>Examples of remote security tokens are as follows.</p><ul><li>The SAML assertion holds the signer public certificate.</li><li>The SAML assertion at which the signed Security Token Reference (STR dereference transform) points.</li></ul>", "remote-token-allowed", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPAuthenticateAUAllowRemoteTokenReferenceIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"au_remote_token_process_service": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the URL for a service that can process the remote security token. This service accepts the WS-Security token as the request of the SOAP call and, if successful, provides the final security token as the response.</p><p>The remote WS-Security token can be signed, encrypted, or encoded DataPower services with different processing actions can process the remote token. Processing can be by decrypting parts of a remote SAML assertion, doing an XSLT transform, or with an AAA policy to assert the token.</p>", "remote-token-url", "").AddNotValidWhen(DmAAAPAuthenticateAURemoteTokenProcessServiceIgnoreVal.String()).String,
				Optional:            true,
			},
			"au_ws_trust_version": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the version of the WS-Trust and the WS-SecureConversation specifications to use when WS-Trust authentication sends a request to the remote STS. Usually these specifications are updated together. The default value is 1.2.", "ws-trust-version", "").AddStringEnum("1.3", "1.2", "1.1").AddDefaultValue("1.2").AddNotValidWhen(DmAAAPAuthenticateAUWSTrustVersionIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1.3", "1.2", "1.1"),
					validators.ConditionalRequiredString(validators.Evaluation{}, DmAAAPAuthenticateAUWSTrustVersionIgnoreVal, true),
				},
				Default: stringdefault.StaticString("1.2"),
			},
			"au_ldap_search_for_dn": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to retrieve the user DN with an LDAP search. By default, the login name that the user presents is used with the LDAP prefix and LDAP suffix to construct the user DN.", "ldap-search-for-dn", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPAuthenticateAULDAPSearchForDNIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"au_ldap_search_parameters": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of an LDAP search parameters that retrieves the user DN.", "ldap-search-param", "ldap_search_parameters").AddRequiredWhen(DmAAAPAuthenticateAULDAPSearchParametersCondVal.String()).AddNotValidWhen(DmAAAPAuthenticateAULDAPSearchParametersIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthenticateAULDAPSearchParametersCondVal, DmAAAPAuthenticateAULDAPSearchParametersIgnoreVal, false),
				},
			},
			"au_ws_trust_require_client_entropy": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to require client entropy in the WS-Trust request. When required, a WS-Trust entropy element is sent by the client as part of the security token request exchange. By default, entropy is not required.</p><ul><li>If a WS-Trust encryption certificate is used, the client entropy material is encrypted.</li><li>If a WS-Trust encryption certificate is not used, a WS-Trust <tt>BinarySecret</tt> element contains the entropy material. In this case, use a TLS profile to secure the exchange with the WS-Trust server.</li></ul>", "trust-require-client-entropy", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPAuthenticateAUWSTrustRequireClientEntropyIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"au_ws_trust_client_entropy_size": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the size in bytes of the WS-Trust client entropy material. The size refers to the length of the entropy before base 64-encoding. Enter a value in the range 8 - 128. The default value is 32.", "trust-client-entropy-size", "").AddIntegerRange(8, 128).AddDefaultValue("32").AddNotValidWhen(DmAAAPAuthenticateAUWSTrustClientEntropySizeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(8, 128),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, DmAAAPAuthenticateAUWSTrustClientEntropySizeIgnoreVal, true),
				},
				Default: int64default.StaticInt64(32),
			},
			"au_ws_trust_require_server_entropy": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to require server entropy in the WS-Trust response. When required, a WS-Trust <tt>entropy</tt> element must be returned to the client as part of the security token request exchange. By default, entropy is not required.", "trust-require-server-entropy", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPAuthenticateAUWSTrustRequireServerEntropyIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"au_ws_trust_server_entropy_size": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the minimum size in bytes for the WS-Trust server entropy. The size refers to the length of the entropy before base 64-encoding. Enter any value in the range 8 - 128. The default value is 32.", "trust-server-entropy-size", "").AddIntegerRange(8, 128).AddDefaultValue("32").AddNotValidWhen(DmAAAPAuthenticateAUWSTrustServerEntropySizeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(8, 128),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, DmAAAPAuthenticateAUWSTrustServerEntropySizeIgnoreVal, true),
				},
				Default: int64default.StaticInt64(32),
			},
			"au_ws_trust_require_rstc": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the message exchange with the client requires a WS-Trust <tt>RequestSecurityToken</tt> or WS-Trust <tt>RequestSecurityTokenCollection</tt> element to be sent by the client. By default, the element is not required.", "trust-require-rstc", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPAuthenticateAUWSTrustRequireRSTCIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"au_ws_trust_require_applies_to_header": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to require a WS-Addressing <tt>AppliesTo</tt> header in the message exchange. By default, the header is not required.", "trust-require-applies-to-header", "").AddDefaultValue("false").AddNotValidWhen(DmAAAPAuthenticateAUWSTrustRequireAppliesToHeaderIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"au_ws_trust_applies_to_header": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the value for the WS-Addressing <tt>AppliesTo</tt> header. The <tt>header</tt> element is included in the WS-Trust request security token message sent to the WS-Trust server.", "trust-applies-to-header", "").AddNotValidWhen(DmAAAPAuthenticateAUWSTrustAppliesToHeaderIgnoreVal.String()).String,
				Optional:            true,
			},
			"au_ws_trust_encryption_certificate": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the certificate to encrypt WS-Trust elements for recipient. If client entropy is configured, the certificate public key encrypts the material for the recipient. If client entropy is configured and this certificate is not specified, use a TLS profile to secure the message exchange.", "trust-encryption-certificate", "crypto_certificate").AddNotValidWhen(DmAAAPAuthenticateAUWSTrustEncryptionCertificateIgnoreVal.String()).String,
				Optional:            true,
			},
			"au_zos_nss_config": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the z/OS NSS client for SAF communication.", "zos-nss-au", "zos_nss_client").AddRequiredWhen(DmAAAPAuthenticateAUZOSNSSConfigCondVal.String()).AddNotValidWhen(DmAAAPAuthenticateAUZOSNSSConfigIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthenticateAUZOSNSSConfigCondVal, DmAAAPAuthenticateAUZOSNSSConfigIgnoreVal, false),
				},
			},
			"au_ldap_attributes": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the list of the extra user attributes to retrieve from the LDAP user store and kept in a <tt>var://context/ldap/auxiliary-attributes</tt> context variable for future use, such as AAA postprocessing. To define the list of LDAP attributes as the auxiliary information for AAA, use the comma (,) as the delimiter. For example, <tt>email, cn, userPassword</tt> .", "au-ldap-attributes", "").AddNotValidWhen(DmAAAPAuthenticateAULDAPAttributesIgnoreVal.String()).String,
				Optional:            true,
			},
			"au_skew_time": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the skew time in seconds. The skew time is the difference between the clock time on the DataPower Gateway and the time on other systems. The default value is 0. <p>When defined, the expiration of the SAML assertion takes the time difference into account.</p><ul><li>For <tt>NotBefore</tt> , validates with <tt>CurrentTime</tt> minus <tt>SkewTime</tt> .</li><li>For <tt>NotOnOrAfter</tt> , validates with <tt>CurrentTime</tt> plus <tt>SkewTime</tt> .</li></ul>", "au-skew-time", "").AddDefaultValue("0").AddNotValidWhen(DmAAAPAuthenticateAUSkewTimeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             int64default.StaticInt64(0),
			},
			"au_tam_pac_return": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to return the Access Manager privilege attribute certificate (PAC) token from a successful authentication for further use. You can use the PAC in the postprocessing phase. By default, The default the PAC token is not returned.", "tam-pac-return", "").AddDefaultValue("false").AddRequiredWhen(DmAAAPAuthenticateAUTAMPACReturnCondVal.String()).AddNotValidWhen(DmAAAPAuthenticateAUTAMPACReturnIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"au_ldap_read_timeout": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration in seconds to wait for a response from the LDAP server before the LDAP connection is closed. Enter a value in the range 0 - 86400. The default value is 60. A value of 0 indicates that the connection never times out. <p>If you configure an LDAP connection pool and assign it to the XML manager, the service can use this LDAP connection pool. The LDAP read timer can work with the idle timer of the LDAP connection pool to remove idle connections from the LDAP connection pool.</p>", "ldap-readtimeout", "").AddIntegerRange(0, 86400).AddDefaultValue("60").AddNotValidWhen(DmAAAPAuthenticateAULDAPReadTimeoutIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 86400),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, DmAAAPAuthenticateAULDAPReadTimeoutIgnoreVal, true),
				},
				Default: int64default.StaticInt64(60),
			},
			"au_ssl_client_config_type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS profile type to secure connections.", "ssl-client-type", "").AddStringEnum("client").AddDefaultValue("client").AddNotValidWhen(DmAAAPAuthenticateAUSSLClientConfigTypeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("client"),
					validators.ConditionalRequiredString(validators.Evaluation{}, DmAAAPAuthenticateAUSSLClientConfigTypeIgnoreVal, true),
				},
				Default: stringdefault.StaticString("client"),
			},
			"au_ssl_client_profile": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS client profile to secure connections.", "ssl-client", "ssl_client_profile").AddNotValidWhen(DmAAAPAuthenticateAUSSLClientProfileIgnoreVal.String()).String,
				Optional:            true,
			},
			"au_ldap_bind_password_alias": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the alias for the password to bind to the LDAP server for the LDAP search. This value is used when the password from the extract identity phase is a WS-Security UsernameToken PasswordDigest. The LDAP server is searched for the corresponding password to verify the PasswordDigest.", "ldap-bind-password-alias", "password_alias").AddNotValidWhen(DmAAAPAuthenticateAULDAPBindPasswordAliasIgnoreVal.String()).String,
				Optional:            true,
			},
			"au_ltpa_key_file_password_alias": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the password alias of the password that decrypts the LTPA key file. This password decrypts certain entries in a WebSphere LTPA key file. This password is not applicable to Domino key files.", "ltpa-key-password-alias", "password_alias").AddRequiredWhen(DmAAAPAuthenticateAULTPAKeyFilePasswordAliasCondVal.String()).AddNotValidWhen(DmAAAPAuthenticateAULTPAKeyFilePasswordAliasIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthenticateAULTPAKeyFilePasswordAliasCondVal, DmAAAPAuthenticateAULTPAKeyFilePasswordAliasIgnoreVal, false),
				},
			},
			"au_sm_request_type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the type of request to make. By default, the request is against the CA Single Sign-On web agent.", "sm-request-type", "").AddStringEnum("webagent", "webservice").AddDefaultValue("webagent").AddRequiredWhen(DmAAAPAuthenticateAUSMRequestTypeCondVal.String()).AddNotValidWhen(DmAAAPAuthenticateAUSMRequestTypeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("webagent", "webservice"),
					validators.ConditionalRequiredString(DmAAAPAuthenticateAUSMRequestTypeCondVal, DmAAAPAuthenticateAUSMRequestTypeIgnoreVal, true),
				},
				Default: stringdefault.StaticString("webagent"),
			},
			"au_sm_cookie_flow": GetDmSMFlowResourceSchema("Specify which flow to include the authentication session cookie.", "sm-cookie-flow", "", false),
			"au_sm_header_flow": GetDmSMFlowResourceSchema("Identifies the flow to include the CA Single Sign-On headers that are generated during authentication. The CA Single Sign-On HTTP headers start with <tt>SM_</tt> .", "sm-header-flow", "", false),
			"au_sm_cookie_attributes": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the cookie attribute policy that allows predefined or custom attributes to be included in CA Single Sign-On cookies.", "cookie-attributes", "cookie_attribute_policy").AddNotValidWhen(DmAAAPAuthenticateAUSMCookieAttributesIgnoreVal.String()).String,
				Optional:            true,
			},
			"au_cache_control": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how to manage the caching of failures. By default, all failures are cached.", "cache-control", "").AddStringEnum("default", "disable-all", "disable-ldap-failures").AddDefaultValue("default").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default", "disable-all", "disable-ldap-failures"),
				},
				Default: stringdefault.StaticString("default"),
			},
		},
	}
	DmAAAPAuthenticateResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmAAAPAuthenticateResourceSchema.Required = true
	} else {
		DmAAAPAuthenticateResourceSchema.Optional = true
		DmAAAPAuthenticateResourceSchema.Computed = true
	}
	return DmAAAPAuthenticateResourceSchema
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
	if !data.AuSslValcred.IsNull() {
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
	if !data.AuSamlArtifactResponder.IsNull() {
		return false
	}
	if !data.AuKerberosVerifySignature.IsNull() {
		return false
	}
	if !data.AuNetegrityBaseUri.IsNull() {
		return false
	}
	if !data.AuSamlAuthQueryServer.IsNull() {
		return false
	}
	if !data.AuSamlVersion.IsNull() {
		return false
	}
	if !data.AuLdapPrefix.IsNull() {
		return false
	}
	if !data.AuLdapSuffix.IsNull() {
		return false
	}
	if !data.AuLdapLoadBalanceGroup.IsNull() {
		return false
	}
	if !data.AuKerberosKeytab.IsNull() {
		return false
	}
	if !data.AuWsTrustUrl.IsNull() {
		return false
	}
	if !data.AuSaml2Issuer.IsNull() {
		return false
	}
	if !data.AuSignerValcred.IsNull() {
		return false
	}
	if !data.AuSignedXpath.IsNull() {
		return false
	}
	if !data.AuLdapBindDn.IsNull() {
		return false
	}
	if !data.AuLdapSearchAttribute.IsNull() {
		return false
	}
	if data.AuLtpaTokenVersionsBitmap != nil {
		if !data.AuLtpaTokenVersionsBitmap.IsNull() {
			return false
		}
	}
	if !data.AuLtpaKeyFile.IsNull() {
		return false
	}
	if !data.AuBinaryTokenX509Valcred.IsNull() {
		return false
	}
	if !data.AuTamServer.IsNull() {
		return false
	}
	if !data.AuAllowRemoteTokenReference.IsNull() {
		return false
	}
	if !data.AuRemoteTokenProcessService.IsNull() {
		return false
	}
	if !data.AuWsTrustVersion.IsNull() {
		return false
	}
	if !data.AuLdapSearchForDn.IsNull() {
		return false
	}
	if !data.AuLdapSearchParameters.IsNull() {
		return false
	}
	if !data.AuWsTrustRequireClientEntropy.IsNull() {
		return false
	}
	if !data.AuWsTrustClientEntropySize.IsNull() {
		return false
	}
	if !data.AuWsTrustRequireServerEntropy.IsNull() {
		return false
	}
	if !data.AuWsTrustServerEntropySize.IsNull() {
		return false
	}
	if !data.AuWsTrustRequireRstc.IsNull() {
		return false
	}
	if !data.AuWsTrustRequireAppliesToHeader.IsNull() {
		return false
	}
	if !data.AuWsTrustAppliesToHeader.IsNull() {
		return false
	}
	if !data.AuWsTrustEncryptionCertificate.IsNull() {
		return false
	}
	if !data.AuZosNssConfig.IsNull() {
		return false
	}
	if !data.AuLdapAttributes.IsNull() {
		return false
	}
	if !data.AuSkewTime.IsNull() {
		return false
	}
	if !data.AuTamPacReturn.IsNull() {
		return false
	}
	if !data.AuLdapReadTimeout.IsNull() {
		return false
	}
	if !data.AuSslClientConfigType.IsNull() {
		return false
	}
	if !data.AuSslClientProfile.IsNull() {
		return false
	}
	if !data.AuLdapBindPasswordAlias.IsNull() {
		return false
	}
	if !data.AuLtpaKeyFilePasswordAlias.IsNull() {
		return false
	}
	if !data.AuSmRequestType.IsNull() {
		return false
	}
	if data.AuSmCookieFlow != nil {
		if !data.AuSmCookieFlow.IsNull() {
			return false
		}
	}
	if data.AuSmHeaderFlow != nil {
		if !data.AuSmHeaderFlow.IsNull() {
			return false
		}
	}
	if !data.AuSmCookieAttributes.IsNull() {
		return false
	}
	if !data.AuCacheControl.IsNull() {
		return false
	}
	return true
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
	if !data.AuSslValcred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUSSLValcred`, data.AuSslValcred.ValueString())
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
	if !data.AuSamlArtifactResponder.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUSAMLArtifactResponder`, data.AuSamlArtifactResponder.ValueString())
	}
	if !data.AuKerberosVerifySignature.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUKerberosVerifySignature`, tfutils.StringFromBool(data.AuKerberosVerifySignature, ""))
	}
	if !data.AuNetegrityBaseUri.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUNetegrityBaseURI`, data.AuNetegrityBaseUri.ValueString())
	}
	if !data.AuSamlAuthQueryServer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUSAMLAuthQueryServer`, data.AuSamlAuthQueryServer.ValueString())
	}
	if !data.AuSamlVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUSAMLVersion`, data.AuSamlVersion.ValueString())
	}
	if !data.AuLdapPrefix.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPPrefix`, data.AuLdapPrefix.ValueString())
	}
	if !data.AuLdapSuffix.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPSuffix`, data.AuLdapSuffix.ValueString())
	}
	if !data.AuLdapLoadBalanceGroup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPLoadBalanceGroup`, data.AuLdapLoadBalanceGroup.ValueString())
	}
	if !data.AuKerberosKeytab.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUKerberosKeytab`, data.AuKerberosKeytab.ValueString())
	}
	if !data.AuWsTrustUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUWSTrustURL`, data.AuWsTrustUrl.ValueString())
	}
	if !data.AuSaml2Issuer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUSAML2Issuer`, data.AuSaml2Issuer.ValueString())
	}
	if !data.AuSignerValcred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUSignerValcred`, data.AuSignerValcred.ValueString())
	}
	if !data.AuSignedXpath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUSignedXPath`, data.AuSignedXpath.ValueString())
	}
	if !data.AuLdapBindDn.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPBindDN`, data.AuLdapBindDn.ValueString())
	}
	if !data.AuLdapSearchAttribute.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPSearchAttribute`, data.AuLdapSearchAttribute.ValueString())
	}
	if data.AuLtpaTokenVersionsBitmap != nil {
		if !data.AuLtpaTokenVersionsBitmap.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`AULTPATokenVersionsBitmap`, data.AuLtpaTokenVersionsBitmap.ToBody(ctx, ""))
		}
	}
	if !data.AuLtpaKeyFile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULTPAKeyFile`, data.AuLtpaKeyFile.ValueString())
	}
	if !data.AuBinaryTokenX509Valcred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUBinaryTokenX509Valcred`, data.AuBinaryTokenX509Valcred.ValueString())
	}
	if !data.AuTamServer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUTAMServer`, data.AuTamServer.ValueString())
	}
	if !data.AuAllowRemoteTokenReference.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUAllowRemoteTokenReference`, tfutils.StringFromBool(data.AuAllowRemoteTokenReference, ""))
	}
	if !data.AuRemoteTokenProcessService.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AURemoteTokenProcessService`, data.AuRemoteTokenProcessService.ValueString())
	}
	if !data.AuWsTrustVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUWSTrustVersion`, data.AuWsTrustVersion.ValueString())
	}
	if !data.AuLdapSearchForDn.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPSearchForDN`, tfutils.StringFromBool(data.AuLdapSearchForDn, ""))
	}
	if !data.AuLdapSearchParameters.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPSearchParameters`, data.AuLdapSearchParameters.ValueString())
	}
	if !data.AuWsTrustRequireClientEntropy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUWSTrustRequireClientEntropy`, tfutils.StringFromBool(data.AuWsTrustRequireClientEntropy, ""))
	}
	if !data.AuWsTrustClientEntropySize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUWSTrustClientEntropySize`, data.AuWsTrustClientEntropySize.ValueInt64())
	}
	if !data.AuWsTrustRequireServerEntropy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUWSTrustRequireServerEntropy`, tfutils.StringFromBool(data.AuWsTrustRequireServerEntropy, ""))
	}
	if !data.AuWsTrustServerEntropySize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUWSTrustServerEntropySize`, data.AuWsTrustServerEntropySize.ValueInt64())
	}
	if !data.AuWsTrustRequireRstc.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUWSTrustRequireRSTC`, tfutils.StringFromBool(data.AuWsTrustRequireRstc, ""))
	}
	if !data.AuWsTrustRequireAppliesToHeader.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUWSTrustRequireAppliesToHeader`, tfutils.StringFromBool(data.AuWsTrustRequireAppliesToHeader, ""))
	}
	if !data.AuWsTrustAppliesToHeader.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUWSTrustAppliesToHeader`, data.AuWsTrustAppliesToHeader.ValueString())
	}
	if !data.AuWsTrustEncryptionCertificate.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUWSTrustEncryptionCertificate`, data.AuWsTrustEncryptionCertificate.ValueString())
	}
	if !data.AuZosNssConfig.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUZOSNSSConfig`, data.AuZosNssConfig.ValueString())
	}
	if !data.AuLdapAttributes.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPAttributes`, data.AuLdapAttributes.ValueString())
	}
	if !data.AuSkewTime.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUSkewTime`, data.AuSkewTime.ValueInt64())
	}
	if !data.AuTamPacReturn.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUTAMPACReturn`, tfutils.StringFromBool(data.AuTamPacReturn, ""))
	}
	if !data.AuLdapReadTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPReadTimeout`, data.AuLdapReadTimeout.ValueInt64())
	}
	if !data.AuSslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUSSLClientConfigType`, data.AuSslClientConfigType.ValueString())
	}
	if !data.AuSslClientProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUSSLClientProfile`, data.AuSslClientProfile.ValueString())
	}
	if !data.AuLdapBindPasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPBindPasswordAlias`, data.AuLdapBindPasswordAlias.ValueString())
	}
	if !data.AuLtpaKeyFilePasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULTPAKeyFilePasswordAlias`, data.AuLtpaKeyFilePasswordAlias.ValueString())
	}
	if !data.AuSmRequestType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUSMRequestType`, data.AuSmRequestType.ValueString())
	}
	if data.AuSmCookieFlow != nil {
		if !data.AuSmCookieFlow.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`AUSMCookieFlow`, data.AuSmCookieFlow.ToBody(ctx, ""))
		}
	}
	if data.AuSmHeaderFlow != nil {
		if !data.AuSmHeaderFlow.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`AUSMHeaderFlow`, data.AuSmHeaderFlow.ToBody(ctx, ""))
		}
	}
	if !data.AuSmCookieAttributes.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUSMCookieAttributes`, data.AuSmCookieAttributes.ValueString())
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
		data.AuSslValcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuSslValcred = types.StringNull()
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
		data.AuSamlArtifactResponder = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuSamlArtifactResponder = types.StringNull()
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
		data.AuSamlAuthQueryServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuSamlAuthQueryServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSAMLVersion`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuSamlVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuSamlVersion = types.StringValue("1.1")
	}
	if value := res.Get(pathRoot + `AULDAPPrefix`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuLdapPrefix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuLdapPrefix = types.StringValue("cn=")
	}
	if value := res.Get(pathRoot + `AULDAPSuffix`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuLdapSuffix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuLdapSuffix = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPLoadBalanceGroup`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuLdapLoadBalanceGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuLdapLoadBalanceGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUKerberosKeytab`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuKerberosKeytab = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuKerberosKeytab = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUWSTrustURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuWsTrustUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuWsTrustUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSAML2Issuer`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuSaml2Issuer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuSaml2Issuer = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSignerValcred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuSignerValcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuSignerValcred = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSignedXPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuSignedXpath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuSignedXpath = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPBindDN`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuLdapBindDn = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuLdapBindDn = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPSearchAttribute`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuLdapSearchAttribute = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuLdapSearchAttribute = types.StringValue("userPassword")
	}
	if value := res.Get(pathRoot + `AULTPATokenVersionsBitmap`); value.Exists() {
		data.AuLtpaTokenVersionsBitmap = &DmLTPATokenVersion{}
		data.AuLtpaTokenVersionsBitmap.FromBody(ctx, "", value)
	} else {
		data.AuLtpaTokenVersionsBitmap = nil
	}
	if value := res.Get(pathRoot + `AULTPAKeyFile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuLtpaKeyFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuLtpaKeyFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUBinaryTokenX509Valcred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuBinaryTokenX509Valcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuBinaryTokenX509Valcred = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUTAMServer`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuTamServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuTamServer = types.StringNull()
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
		data.AuWsTrustVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuWsTrustVersion = types.StringValue("1.2")
	}
	if value := res.Get(pathRoot + `AULDAPSearchForDN`); value.Exists() {
		data.AuLdapSearchForDn = tfutils.BoolFromString(value.String())
	} else {
		data.AuLdapSearchForDn = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AULDAPSearchParameters`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuLdapSearchParameters = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuLdapSearchParameters = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUWSTrustRequireClientEntropy`); value.Exists() {
		data.AuWsTrustRequireClientEntropy = tfutils.BoolFromString(value.String())
	} else {
		data.AuWsTrustRequireClientEntropy = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AUWSTrustClientEntropySize`); value.Exists() {
		data.AuWsTrustClientEntropySize = types.Int64Value(value.Int())
	} else {
		data.AuWsTrustClientEntropySize = types.Int64Value(32)
	}
	if value := res.Get(pathRoot + `AUWSTrustRequireServerEntropy`); value.Exists() {
		data.AuWsTrustRequireServerEntropy = tfutils.BoolFromString(value.String())
	} else {
		data.AuWsTrustRequireServerEntropy = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AUWSTrustServerEntropySize`); value.Exists() {
		data.AuWsTrustServerEntropySize = types.Int64Value(value.Int())
	} else {
		data.AuWsTrustServerEntropySize = types.Int64Value(32)
	}
	if value := res.Get(pathRoot + `AUWSTrustRequireRSTC`); value.Exists() {
		data.AuWsTrustRequireRstc = tfutils.BoolFromString(value.String())
	} else {
		data.AuWsTrustRequireRstc = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AUWSTrustRequireAppliesToHeader`); value.Exists() {
		data.AuWsTrustRequireAppliesToHeader = tfutils.BoolFromString(value.String())
	} else {
		data.AuWsTrustRequireAppliesToHeader = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AUWSTrustAppliesToHeader`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuWsTrustAppliesToHeader = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuWsTrustAppliesToHeader = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUWSTrustEncryptionCertificate`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuWsTrustEncryptionCertificate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuWsTrustEncryptionCertificate = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUZOSNSSConfig`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuZosNssConfig = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuZosNssConfig = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPAttributes`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuLdapAttributes = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuLdapAttributes = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSkewTime`); value.Exists() {
		data.AuSkewTime = types.Int64Value(value.Int())
	} else {
		data.AuSkewTime = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `AUTAMPACReturn`); value.Exists() {
		data.AuTamPacReturn = tfutils.BoolFromString(value.String())
	} else {
		data.AuTamPacReturn = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AULDAPReadTimeout`); value.Exists() {
		data.AuLdapReadTimeout = types.Int64Value(value.Int())
	} else {
		data.AuLdapReadTimeout = types.Int64Value(60)
	}
	if value := res.Get(pathRoot + `AUSSLClientConfigType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuSslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuSslClientConfigType = types.StringValue("client")
	}
	if value := res.Get(pathRoot + `AUSSLClientProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuSslClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuSslClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPBindPasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuLdapBindPasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuLdapBindPasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULTPAKeyFilePasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuLtpaKeyFilePasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuLtpaKeyFilePasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSMRequestType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuSmRequestType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuSmRequestType = types.StringValue("webagent")
	}
	if value := res.Get(pathRoot + `AUSMCookieFlow`); value.Exists() {
		data.AuSmCookieFlow = &DmSMFlow{}
		data.AuSmCookieFlow.FromBody(ctx, "", value)
	} else {
		data.AuSmCookieFlow = nil
	}
	if value := res.Get(pathRoot + `AUSMHeaderFlow`); value.Exists() {
		data.AuSmHeaderFlow = &DmSMFlow{}
		data.AuSmHeaderFlow.FromBody(ctx, "", value)
	} else {
		data.AuSmHeaderFlow = nil
	}
	if value := res.Get(pathRoot + `AUSMCookieAttributes`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuSmCookieAttributes = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuSmCookieAttributes = types.StringNull()
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
	if value := res.Get(pathRoot + `AUSSLValcred`); value.Exists() && !data.AuSslValcred.IsNull() {
		data.AuSslValcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuSslValcred = types.StringNull()
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
	if value := res.Get(pathRoot + `AUSAMLArtifactResponder`); value.Exists() && !data.AuSamlArtifactResponder.IsNull() {
		data.AuSamlArtifactResponder = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuSamlArtifactResponder = types.StringNull()
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
	if value := res.Get(pathRoot + `AUSAMLAuthQueryServer`); value.Exists() && !data.AuSamlAuthQueryServer.IsNull() {
		data.AuSamlAuthQueryServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuSamlAuthQueryServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSAMLVersion`); value.Exists() && !data.AuSamlVersion.IsNull() {
		data.AuSamlVersion = tfutils.ParseStringFromGJSON(value)
	} else if data.AuSamlVersion.ValueString() != "1.1" {
		data.AuSamlVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPPrefix`); value.Exists() && !data.AuLdapPrefix.IsNull() {
		data.AuLdapPrefix = tfutils.ParseStringFromGJSON(value)
	} else if data.AuLdapPrefix.ValueString() != "cn=" {
		data.AuLdapPrefix = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPSuffix`); value.Exists() && !data.AuLdapSuffix.IsNull() {
		data.AuLdapSuffix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuLdapSuffix = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPLoadBalanceGroup`); value.Exists() && !data.AuLdapLoadBalanceGroup.IsNull() {
		data.AuLdapLoadBalanceGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuLdapLoadBalanceGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUKerberosKeytab`); value.Exists() && !data.AuKerberosKeytab.IsNull() {
		data.AuKerberosKeytab = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuKerberosKeytab = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUWSTrustURL`); value.Exists() && !data.AuWsTrustUrl.IsNull() {
		data.AuWsTrustUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuWsTrustUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSAML2Issuer`); value.Exists() && !data.AuSaml2Issuer.IsNull() {
		data.AuSaml2Issuer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuSaml2Issuer = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSignerValcred`); value.Exists() && !data.AuSignerValcred.IsNull() {
		data.AuSignerValcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuSignerValcred = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSignedXPath`); value.Exists() && !data.AuSignedXpath.IsNull() {
		data.AuSignedXpath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuSignedXpath = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPBindDN`); value.Exists() && !data.AuLdapBindDn.IsNull() {
		data.AuLdapBindDn = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuLdapBindDn = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPSearchAttribute`); value.Exists() && !data.AuLdapSearchAttribute.IsNull() {
		data.AuLdapSearchAttribute = tfutils.ParseStringFromGJSON(value)
	} else if data.AuLdapSearchAttribute.ValueString() != "userPassword" {
		data.AuLdapSearchAttribute = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULTPATokenVersionsBitmap`); value.Exists() {
		data.AuLtpaTokenVersionsBitmap.UpdateFromBody(ctx, "", value)
	} else {
		data.AuLtpaTokenVersionsBitmap = nil
	}
	if value := res.Get(pathRoot + `AULTPAKeyFile`); value.Exists() && !data.AuLtpaKeyFile.IsNull() {
		data.AuLtpaKeyFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuLtpaKeyFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUBinaryTokenX509Valcred`); value.Exists() && !data.AuBinaryTokenX509Valcred.IsNull() {
		data.AuBinaryTokenX509Valcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuBinaryTokenX509Valcred = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUTAMServer`); value.Exists() && !data.AuTamServer.IsNull() {
		data.AuTamServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuTamServer = types.StringNull()
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
	if value := res.Get(pathRoot + `AUWSTrustVersion`); value.Exists() && !data.AuWsTrustVersion.IsNull() {
		data.AuWsTrustVersion = tfutils.ParseStringFromGJSON(value)
	} else if data.AuWsTrustVersion.ValueString() != "1.2" {
		data.AuWsTrustVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPSearchForDN`); value.Exists() && !data.AuLdapSearchForDn.IsNull() {
		data.AuLdapSearchForDn = tfutils.BoolFromString(value.String())
	} else if data.AuLdapSearchForDn.ValueBool() {
		data.AuLdapSearchForDn = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AULDAPSearchParameters`); value.Exists() && !data.AuLdapSearchParameters.IsNull() {
		data.AuLdapSearchParameters = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuLdapSearchParameters = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUWSTrustRequireClientEntropy`); value.Exists() && !data.AuWsTrustRequireClientEntropy.IsNull() {
		data.AuWsTrustRequireClientEntropy = tfutils.BoolFromString(value.String())
	} else if data.AuWsTrustRequireClientEntropy.ValueBool() {
		data.AuWsTrustRequireClientEntropy = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AUWSTrustClientEntropySize`); value.Exists() && !data.AuWsTrustClientEntropySize.IsNull() {
		data.AuWsTrustClientEntropySize = types.Int64Value(value.Int())
	} else if data.AuWsTrustClientEntropySize.ValueInt64() != 32 {
		data.AuWsTrustClientEntropySize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AUWSTrustRequireServerEntropy`); value.Exists() && !data.AuWsTrustRequireServerEntropy.IsNull() {
		data.AuWsTrustRequireServerEntropy = tfutils.BoolFromString(value.String())
	} else if data.AuWsTrustRequireServerEntropy.ValueBool() {
		data.AuWsTrustRequireServerEntropy = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AUWSTrustServerEntropySize`); value.Exists() && !data.AuWsTrustServerEntropySize.IsNull() {
		data.AuWsTrustServerEntropySize = types.Int64Value(value.Int())
	} else if data.AuWsTrustServerEntropySize.ValueInt64() != 32 {
		data.AuWsTrustServerEntropySize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AUWSTrustRequireRSTC`); value.Exists() && !data.AuWsTrustRequireRstc.IsNull() {
		data.AuWsTrustRequireRstc = tfutils.BoolFromString(value.String())
	} else if data.AuWsTrustRequireRstc.ValueBool() {
		data.AuWsTrustRequireRstc = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AUWSTrustRequireAppliesToHeader`); value.Exists() && !data.AuWsTrustRequireAppliesToHeader.IsNull() {
		data.AuWsTrustRequireAppliesToHeader = tfutils.BoolFromString(value.String())
	} else if data.AuWsTrustRequireAppliesToHeader.ValueBool() {
		data.AuWsTrustRequireAppliesToHeader = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AUWSTrustAppliesToHeader`); value.Exists() && !data.AuWsTrustAppliesToHeader.IsNull() {
		data.AuWsTrustAppliesToHeader = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuWsTrustAppliesToHeader = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUWSTrustEncryptionCertificate`); value.Exists() && !data.AuWsTrustEncryptionCertificate.IsNull() {
		data.AuWsTrustEncryptionCertificate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuWsTrustEncryptionCertificate = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUZOSNSSConfig`); value.Exists() && !data.AuZosNssConfig.IsNull() {
		data.AuZosNssConfig = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuZosNssConfig = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPAttributes`); value.Exists() && !data.AuLdapAttributes.IsNull() {
		data.AuLdapAttributes = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuLdapAttributes = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSkewTime`); value.Exists() && !data.AuSkewTime.IsNull() {
		data.AuSkewTime = types.Int64Value(value.Int())
	} else if data.AuSkewTime.ValueInt64() != 0 {
		data.AuSkewTime = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AUTAMPACReturn`); value.Exists() && !data.AuTamPacReturn.IsNull() {
		data.AuTamPacReturn = tfutils.BoolFromString(value.String())
	} else if data.AuTamPacReturn.ValueBool() {
		data.AuTamPacReturn = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AULDAPReadTimeout`); value.Exists() && !data.AuLdapReadTimeout.IsNull() {
		data.AuLdapReadTimeout = types.Int64Value(value.Int())
	} else if data.AuLdapReadTimeout.ValueInt64() != 60 {
		data.AuLdapReadTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AUSSLClientConfigType`); value.Exists() && !data.AuSslClientConfigType.IsNull() {
		data.AuSslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else if data.AuSslClientConfigType.ValueString() != "client" {
		data.AuSslClientConfigType = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSSLClientProfile`); value.Exists() && !data.AuSslClientProfile.IsNull() {
		data.AuSslClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuSslClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPBindPasswordAlias`); value.Exists() && !data.AuLdapBindPasswordAlias.IsNull() {
		data.AuLdapBindPasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuLdapBindPasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULTPAKeyFilePasswordAlias`); value.Exists() && !data.AuLtpaKeyFilePasswordAlias.IsNull() {
		data.AuLtpaKeyFilePasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuLtpaKeyFilePasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSMRequestType`); value.Exists() && !data.AuSmRequestType.IsNull() {
		data.AuSmRequestType = tfutils.ParseStringFromGJSON(value)
	} else if data.AuSmRequestType.ValueString() != "webagent" {
		data.AuSmRequestType = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSMCookieFlow`); value.Exists() {
		data.AuSmCookieFlow.UpdateFromBody(ctx, "", value)
	} else {
		data.AuSmCookieFlow = nil
	}
	if value := res.Get(pathRoot + `AUSMHeaderFlow`); value.Exists() {
		data.AuSmHeaderFlow.UpdateFromBody(ctx, "", value)
	} else {
		data.AuSmHeaderFlow = nil
	}
	if value := res.Get(pathRoot + `AUSMCookieAttributes`); value.Exists() && !data.AuSmCookieAttributes.IsNull() {
		data.AuSmCookieAttributes = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuSmCookieAttributes = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUCacheControl`); value.Exists() && !data.AuCacheControl.IsNull() {
		data.AuCacheControl = tfutils.ParseStringFromGJSON(value)
	} else if data.AuCacheControl.ValueString() != "default" {
		data.AuCacheControl = types.StringNull()
	}
}
