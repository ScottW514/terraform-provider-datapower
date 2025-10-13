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

type DmAAAPAuthorize struct {
	AzMethod                   types.String `tfsdk:"az_method"`
	AzCustomUrl                types.String `tfsdk:"az_custom_url"`
	AzMapUrl                   types.String `tfsdk:"az_map_url"`
	AzHost                     types.String `tfsdk:"az_host"`
	AzPort                     types.Int64  `tfsdk:"az_port"`
	AzLdapGroup                types.String `tfsdk:"az_ldap_group"`
	AzValcred                  types.String `tfsdk:"az_valcred"`
	AzSamlUrl                  types.String `tfsdk:"az_saml_url"`
	AzSamlType                 types.String `tfsdk:"az_saml_type"`
	AzSamlXpath                types.String `tfsdk:"az_saml_xpath"`
	AzSamlNameQualifier        types.String `tfsdk:"az_saml_name_qualifier"`
	AzCacheAllow               types.String `tfsdk:"az_cache_allow"`
	AzCacheTtl                 types.Int64  `tfsdk:"az_cache_ttl"`
	AzNetegrityBaseUri         types.String `tfsdk:"az_netegrity_base_uri"`
	AzNetegrityOpNameExtension types.String `tfsdk:"az_netegrity_op_name_extension"`
	AzClearTrustServerUrl      types.String `tfsdk:"az_clear_trust_server_url"`
	AzSamlVersion              types.String `tfsdk:"az_saml_version"`
	AzLdapLoadBalanceGroup     types.String `tfsdk:"az_ldap_load_balance_group"`
	AzLdapBindDn               types.String `tfsdk:"az_ldap_bind_dn"`
	AzLdapGroupAttribute       types.String `tfsdk:"az_ldap_group_attribute"`
	AzLdapSearchScope          types.String `tfsdk:"az_ldap_search_scope"`
	AzLdapSearchFilter         types.String `tfsdk:"az_ldap_search_filter"`
	AzXacmlVersion             types.String `tfsdk:"az_xacml_version"`
	AzXacmlPepType             types.String `tfsdk:"az_xacml_pep_type"`
	AzXacmlUseOnBoxPdp         types.Bool   `tfsdk:"az_xacml_use_on_box_pdp"`
	AzXacmlPdp                 types.String `tfsdk:"az_xacml_pdp"`
	AzXacmlExternalPdpUrl      types.String `tfsdk:"az_xacml_external_pdp_url"`
	AzXacmlBindingMethod       types.String `tfsdk:"az_xacml_binding_method"`
	AzXacmlBindingXsl          types.String `tfsdk:"az_xacml_binding_xsl"`
	AzXacmlCustomObligation    types.String `tfsdk:"az_xacml_custom_obligation"`
	AzXacmlUseSaml2            types.Bool   `tfsdk:"az_xacml_use_saml2"`
	AzTamServer                types.String `tfsdk:"az_tam_server"`
	AzTamDefaultAction         types.String `tfsdk:"az_tam_default_action"`
	AzTamActionResourceMap     types.String `tfsdk:"az_tam_action_resource_map"`
	AzXacmlUseSoap             types.Bool   `tfsdk:"az_xacml_use_soap"`
	AzZosNssConfig             types.String `tfsdk:"az_zos_nss_config"`
	AzSafDefaultAction         types.String `tfsdk:"az_saf_default_action"`
	AzLdapAttributes           types.String `tfsdk:"az_ldap_attributes"`
	AzSkewTime                 types.Int64  `tfsdk:"az_skew_time"`
	AzOauthEnforceScope        types.Bool   `tfsdk:"az_oauth_enforce_scope"`
	AzOauthExportHeaders       types.Bool   `tfsdk:"az_oauth_export_headers"`
	AzTamPacReturn             types.Bool   `tfsdk:"az_tam_pac_return"`
	AzTamPacUse                types.Bool   `tfsdk:"az_tam_pac_use"`
	AzLdapReadTimeout          types.Int64  `tfsdk:"az_ldap_read_timeout"`
	AzSslClientConfigType      types.String `tfsdk:"az_ssl_client_config_type"`
	AzSslClientProfile         types.String `tfsdk:"az_ssl_client_profile"`
	AzLdapBindPasswordAlias    types.String `tfsdk:"az_ldap_bind_password_alias"`
	AzSmRequestType            types.String `tfsdk:"az_sm_request_type"`
	AzSmCookieFlow             *DmSMFlow    `tfsdk:"az_sm_cookie_flow"`
	AzSmHeaderFlow             *DmSMFlow    `tfsdk:"az_sm_header_flow"`
	AzSmCookieAttributes       types.String `tfsdk:"az_sm_cookie_attributes"`
	AzCacheControl             types.String `tfsdk:"az_cache_control"`
}

var DmAAAPAuthorizeAZCustomURLCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"custom"},
}

var DmAAAPAuthorizeAZCustomURLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthorizeAZMapURLCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"xmlfile"},
}

var DmAAAPAuthorizeAZMapURLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthorizeAZHostCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_method",
			AttrType:    "String",
			AttrDefault: "anyauthenticated",
			Value:       []string{"ldap", "oblix", "netegrity"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_ldap_load_balance_group",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{""},
		},
	},
}

var DmAAAPAuthorizeAZHostIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "az_method",
			AttrType:    "String",
			AttrDefault: "anyauthenticated",
			Value:       []string{"ldap", "netegrity", "oblix"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "az_ldap_load_balance_group",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{""},
		},
	},
}

var DmAAAPAuthorizeAZPortCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_method",
			AttrType:    "String",
			AttrDefault: "anyauthenticated",
			Value:       []string{"ldap", "oblix", "netegrity"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_ldap_load_balance_group",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{""},
		},
	},
}

var DmAAAPAuthorizeAZPortIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "az_method",
			AttrType:    "String",
			AttrDefault: "anyauthenticated",
			Value:       []string{"ldap", "netegrity", "oblix"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "az_ldap_load_balance_group",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{""},
		},
	},
}

var DmAAAPAuthorizeAZLDAPGroupCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"ldap"},
}

var DmAAAPAuthorizeAZLDAPGroupIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthorizeAZValcredIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthorizeAZSAMLURLCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"saml-attr", "saml-authz"},
}

var DmAAAPAuthorizeAZSAMLURLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthorizeAZSAMLTypeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"saml-attr", "saml-authz", "use-authen-attr"},
}

var DmAAAPAuthorizeAZSAMLTypeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthorizeAZSAMLXPathCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_saml_type",
			AttrType:    "String",
			AttrDefault: "any",
			Value:       []string{"xpath"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_method",
			AttrType:    "String",
			AttrDefault: "anyauthenticated",
			Value:       []string{"saml-attr", "saml-authz", "use-authen-attr"},
		},
	},
}

var DmAAAPAuthorizeAZSAMLXPathIgnoreVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "az_saml_type",
			AttrType:    "String",
			AttrDefault: "any",
			Value:       []string{"xpath"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "az_method",
			AttrType:    "String",
			AttrDefault: "anyauthenticated",
			Value:       []string{"saml-attr", "saml-authz", "use-authen-attr"},
		},
	},
}

var DmAAAPAuthorizeAZSAMLNameQualifierIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"saml-attr", "saml-authz"},
}

var DmAAAPAuthorizeAZCacheTTLIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "az_cache_allow",
	AttrType:    "String",
	AttrDefault: "absolute",
	Value:       []string{"disabled"},
}

var DmAAAPAuthorizeAZNetegrityBaseURIIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"netegrity"},
}

var DmAAAPAuthorizeAZNetegrityOpNameExtensionIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "az_method",
			AttrType:    "String",
			AttrDefault: "anyauthenticated",
			Value:       []string{"netegrity"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_sm_request_type",
			AttrType:    "String",
			AttrDefault: "webagent",
			Value:       []string{"webservice"},
		},
	},
}

var DmAAAPAuthorizeAZClearTrustServerURLCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"cleartrust"},
}

var DmAAAPAuthorizeAZClearTrustServerURLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthorizeAZSAMLVersionIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"saml-attr", "saml-authz"},
}

var DmAAAPAuthorizeAZLDAPLoadBalanceGroupIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"ldap"},
}

var DmAAAPAuthorizeAZLDAPBindDNIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"ldap"},
}

var DmAAAPAuthorizeAZLDAPGroupAttributeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"ldap"},
}

var DmAAAPAuthorizeAZLDAPSearchScopeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"ldap"},
}

var DmAAAPAuthorizeAZLDAPSearchFilterIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"ldap"},
}

var DmAAAPAuthorizeAZXACMLVersionCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"xacml"},
}

var DmAAAPAuthorizeAZXACMLVersionIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthorizeAZXACMLPEPTypeCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_method",
			AttrType:    "String",
			AttrDefault: "anyauthenticated",
			Value:       []string{"xacml"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "az_xacml_version",
			AttrType:    "String",
			AttrDefault: "2",
			Value:       []string{"1.0"},
		},
	},
}

var DmAAAPAuthorizeAZXACMLPEPTypeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthorizeAZXACMLUseOnBoxPDPCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"xacml"},
}

var DmAAAPAuthorizeAZXACMLUseOnBoxPDPIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthorizeAZXACMLPDPCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_method",
			AttrType:    "String",
			AttrDefault: "anyauthenticated",
			Value:       []string{"xacml"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_xacml_use_on_box_pdp",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
	},
}

var DmAAAPAuthorizeAZXACMLPDPIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthorizeAZXACMLExternalPDPUrlCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_method",
			AttrType:    "String",
			AttrDefault: "anyauthenticated",
			Value:       []string{"xacml"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_xacml_use_on_box_pdp",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"false"},
		},
	},
}

var DmAAAPAuthorizeAZXACMLExternalPDPUrlIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthorizeAZXACMLBindingMethodIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"xacml"},
}

var DmAAAPAuthorizeAZXACMLBindingXSLCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_method",
			AttrType:    "String",
			AttrDefault: "anyauthenticated",
			Value:       []string{"xacml"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "az_xacml_binding_method",
			AttrType:    "String",
			AttrDefault: "custom",
			Value:       []string{"dp-pdp"},
		},
	},
}

var DmAAAPAuthorizeAZXACMLBindingXSLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthorizeAZXACMLCustomObligationIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"xacml"},
}

var DmAAAPAuthorizeAZXACMLUseSAML2CondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_method",
			AttrType:    "String",
			AttrDefault: "anyauthenticated",
			Value:       []string{"xacml"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_xacml_version",
			AttrType:    "String",
			AttrDefault: "2",
			Value:       []string{"2.0"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_xacml_use_on_box_pdp",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"false"},
		},
	},
}

var DmAAAPAuthorizeAZXACMLUseSAML2IgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthorizeAZTAMServerCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"tivoli"},
}

var DmAAAPAuthorizeAZTAMServerIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthorizeAZTAMDefaultActionIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"tivoli"},
}

var DmAAAPAuthorizeAZTAMActionResourceMapIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"tivoli"},
}

var DmAAAPAuthorizeAZXACMLUseSOAPCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_method",
			AttrType:    "String",
			AttrDefault: "anyauthenticated",
			Value:       []string{"xacml"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_xacml_use_on_box_pdp",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"false"},
		},
	},
}

var DmAAAPAuthorizeAZXACMLUseSOAPIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthorizeAZZOSNSSConfigCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"zosnss"},
}

var DmAAAPAuthorizeAZZOSNSSConfigIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthorizeAZSAFDefaultActionIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"zosnss"},
}

var DmAAAPAuthorizeAZLDAPAttributesIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"ldap"},
}

var DmAAAPAuthorizeAZSkewTimeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"saml-authz", "saml-attr"},
}

var DmAAAPAuthorizeAZOAuthEnforceScopeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"oauth"},
}

var DmAAAPAuthorizeAZOAuthEnforceScopeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthorizeAZOAuthExportHeadersCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"oauth"},
}

var DmAAAPAuthorizeAZOAuthExportHeadersIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthorizeAZTAMPACReturnCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"tivoli"},
}

var DmAAAPAuthorizeAZTAMPACReturnIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthorizeAZTAMPACUseCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"tivoli"},
}

var DmAAAPAuthorizeAZTAMPACUseIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPAuthorizeAZLDAPReadTimeoutIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"ldap"},
}

var DmAAAPAuthorizeAZSSLClientConfigTypeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"cleartrust", "ldap", "netegrity", "saml-attr", "saml-authz"},
}

var DmAAAPAuthorizeAZSSLClientProfileIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "az_method",
			AttrType:    "String",
			AttrDefault: "anyauthenticated",
			Value:       []string{"cleartrust", "ldap", "netegrity", "saml-attr", "saml-authz"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "az_ssl_client_config_type",
			AttrType:    "String",
			AttrDefault: "client",
			Value:       []string{"client"},
		},
	},
}

var DmAAAPAuthorizeAZLDAPBindPasswordAliasIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"ldap"},
}

var DmAAAPAuthorizeAZSMRequestTypeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"netegrity"},
}

var DmAAAPAuthorizeAZSMRequestTypeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"netegrity"},
}

var DmAAAPAuthorizeAZSMCookieFlowIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"netegrity"},
}

var DmAAAPAuthorizeAZSMHeaderFlowIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"netegrity"},
}

var DmAAAPAuthorizeAZSMCookieAttributesIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "az_sm_cookie_flow",
	AttrType:    "DmSMFlow",
	AttrDefault: "",
	Value:       []string{"frontend", "frontend+backend", "backend+frontend"},
}

var DmAAAPAuthorizeObjectType = map[string]attr.Type{
	"az_method":                      types.StringType,
	"az_custom_url":                  types.StringType,
	"az_map_url":                     types.StringType,
	"az_host":                        types.StringType,
	"az_port":                        types.Int64Type,
	"az_ldap_group":                  types.StringType,
	"az_valcred":                     types.StringType,
	"az_saml_url":                    types.StringType,
	"az_saml_type":                   types.StringType,
	"az_saml_xpath":                  types.StringType,
	"az_saml_name_qualifier":         types.StringType,
	"az_cache_allow":                 types.StringType,
	"az_cache_ttl":                   types.Int64Type,
	"az_netegrity_base_uri":          types.StringType,
	"az_netegrity_op_name_extension": types.StringType,
	"az_clear_trust_server_url":      types.StringType,
	"az_saml_version":                types.StringType,
	"az_ldap_load_balance_group":     types.StringType,
	"az_ldap_bind_dn":                types.StringType,
	"az_ldap_group_attribute":        types.StringType,
	"az_ldap_search_scope":           types.StringType,
	"az_ldap_search_filter":          types.StringType,
	"az_xacml_version":               types.StringType,
	"az_xacml_pep_type":              types.StringType,
	"az_xacml_use_on_box_pdp":        types.BoolType,
	"az_xacml_pdp":                   types.StringType,
	"az_xacml_external_pdp_url":      types.StringType,
	"az_xacml_binding_method":        types.StringType,
	"az_xacml_binding_xsl":           types.StringType,
	"az_xacml_custom_obligation":     types.StringType,
	"az_xacml_use_saml2":             types.BoolType,
	"az_tam_server":                  types.StringType,
	"az_tam_default_action":          types.StringType,
	"az_tam_action_resource_map":     types.StringType,
	"az_xacml_use_soap":              types.BoolType,
	"az_zos_nss_config":              types.StringType,
	"az_saf_default_action":          types.StringType,
	"az_ldap_attributes":             types.StringType,
	"az_skew_time":                   types.Int64Type,
	"az_oauth_enforce_scope":         types.BoolType,
	"az_oauth_export_headers":        types.BoolType,
	"az_tam_pac_return":              types.BoolType,
	"az_tam_pac_use":                 types.BoolType,
	"az_ldap_read_timeout":           types.Int64Type,
	"az_ssl_client_config_type":      types.StringType,
	"az_ssl_client_profile":          types.StringType,
	"az_ldap_bind_password_alias":    types.StringType,
	"az_sm_request_type":             types.StringType,
	"az_sm_cookie_flow":              types.ObjectType{AttrTypes: DmSMFlowObjectType},
	"az_sm_header_flow":              types.ObjectType{AttrTypes: DmSMFlowObjectType},
	"az_sm_cookie_attributes":        types.StringType,
	"az_cache_control":               types.StringType,
}
var DmAAAPAuthorizeObjectDefault = map[string]attr.Value{
	"az_method":                      types.StringValue("anyauthenticated"),
	"az_custom_url":                  types.StringNull(),
	"az_map_url":                     types.StringNull(),
	"az_host":                        types.StringNull(),
	"az_port":                        types.Int64Value(0),
	"az_ldap_group":                  types.StringNull(),
	"az_valcred":                     types.StringNull(),
	"az_saml_url":                    types.StringNull(),
	"az_saml_type":                   types.StringValue("any"),
	"az_saml_xpath":                  types.StringNull(),
	"az_saml_name_qualifier":         types.StringNull(),
	"az_cache_allow":                 types.StringValue("absolute"),
	"az_cache_ttl":                   types.Int64Value(3),
	"az_netegrity_base_uri":          types.StringNull(),
	"az_netegrity_op_name_extension": types.StringNull(),
	"az_clear_trust_server_url":      types.StringNull(),
	"az_saml_version":                types.StringValue("1.1"),
	"az_ldap_load_balance_group":     types.StringNull(),
	"az_ldap_bind_dn":                types.StringNull(),
	"az_ldap_group_attribute":        types.StringValue("member"),
	"az_ldap_search_scope":           types.StringValue("subtree"),
	"az_ldap_search_filter":          types.StringValue("(objectClass=*)"),
	"az_xacml_version":               types.StringValue("2"),
	"az_xacml_pep_type":              types.StringValue("deny-biased"),
	"az_xacml_use_on_box_pdp":        types.BoolValue(true),
	"az_xacml_pdp":                   types.StringNull(),
	"az_xacml_external_pdp_url":      types.StringNull(),
	"az_xacml_binding_method":        types.StringValue("custom"),
	"az_xacml_binding_xsl":           types.StringNull(),
	"az_xacml_custom_obligation":     types.StringNull(),
	"az_xacml_use_saml2":             types.BoolValue(false),
	"az_tam_server":                  types.StringNull(),
	"az_tam_default_action":          types.StringValue("T"),
	"az_tam_action_resource_map":     types.StringNull(),
	"az_xacml_use_soap":              types.BoolValue(false),
	"az_zos_nss_config":              types.StringNull(),
	"az_saf_default_action":          types.StringValue("r"),
	"az_ldap_attributes":             types.StringNull(),
	"az_skew_time":                   types.Int64Value(0),
	"az_oauth_enforce_scope":         types.BoolValue(false),
	"az_oauth_export_headers":        types.BoolValue(true),
	"az_tam_pac_return":              types.BoolValue(false),
	"az_tam_pac_use":                 types.BoolValue(false),
	"az_ldap_read_timeout":           types.Int64Value(60),
	"az_ssl_client_config_type":      types.StringValue("client"),
	"az_ssl_client_profile":          types.StringNull(),
	"az_ldap_bind_password_alias":    types.StringNull(),
	"az_sm_request_type":             types.StringValue("webagent"),
	"az_sm_cookie_flow":              types.ObjectValueMust(DmSMFlowObjectType, DmSMFlowObjectDefault),
	"az_sm_header_flow":              types.ObjectValueMust(DmSMFlowObjectType, DmSMFlowObjectDefault),
	"az_sm_cookie_attributes":        types.StringNull(),
	"az_cache_control":               types.StringValue("default"),
}

func GetDmAAAPAuthorizeDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmAAAPAuthorizeDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"az_method": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the authorization method.",
				Computed:            true,
			},
			"az_custom_url": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the location of the stylesheet or GatewayScript file for custom authorization.",
				Computed:            true,
			},
			"az_map_url": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the location of the AAA information file.",
				Computed:            true,
			},
			"az_host": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the host name or IP address of the authorization server.",
				Computed:            true,
			},
			"az_port": DataSourceSchema.Int64Attribute{
				MarkdownDescription: "Specify the listening port on the authorization server.",
				Computed:            true,
			},
			"az_ldap_group": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the DN of the required LDAP group.",
				Computed:            true,
			},
			"az_valcred": DataSourceSchema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"az_saml_url": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the URL of the SAML server.",
				Computed:            true,
			},
			"az_saml_type": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify how to match SAML attributes and values. The default value is any.",
				Computed:            true,
			},
			"az_saml_xpath": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the XPath expression to run against the SAML statement.",
				Computed:            true,
			},
			"az_saml_name_qualifier": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specifies the value of the <tt>NameQualifier</tt> attribute of the <tt>NameIdentifier</tt> in the generated SAML query. Although the <tt>NameQualifier</tt> attribute is an optional attribute, some SAML implementations require it to be present.",
				Computed:            true,
			},
			"az_cache_allow": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify how to control the caching of AAA authorization results. The default value is absolute. A protocol TTL is available only with SAML or OAuth with a Federated Identity Manager endpoint. Federated Identity Manager integration is deprecated.",
				Computed:            true,
			},
			"az_cache_ttl": DataSourceSchema.Int64Attribute{
				MarkdownDescription: "Specify the duration in seconds to cache authorization decisions. Enter a value in the range 1 - 86400. The default value is 3.",
				Computed:            true,
			},
			"az_netegrity_base_uri": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the URI sent to CA Single Sign-On (formerly Netegrity SiteMinder) server. The CA Single Sign-On base URI is combined with the host, port, and CA Single Sign-On operation name extension to form the URL for attempting CA Single Sign-On authentication. The URL is of the http://host:port/NetegrityBaseURI/operationNetegrityOpNameExtension form, where NetegrityOpNameExtension is concatenated directly with the operation name.",
				Computed:            true,
			},
			"az_netegrity_op_name_extension": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the extension for URI sent to CA Single Sign-On (formerly Netegrity SiteMinder) server. The CA Single Sign-On base URI is combined with the host, port, and CA Single Sign-On operation name extension to form the URL for attempting CA Single Sign-On authentication. The URL is of the http://host:port/NetegrityBaseURI/operationNetegrityOpNameExtension form, where NetegrityOpNameExtension is concatenated directly with the operation name.",
				Computed:            true,
			},
			"az_clear_trust_server_url": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the URL to access the ClearTrust server for authorization.",
				Computed:            true,
			},
			"az_saml_version": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the version to use for SAML messages. The default value is 1.1.",
				Computed:            true,
			},
			"az_ldap_load_balance_group": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the load balancer group that contains the LDAP servers.",
				Computed:            true,
			},
			"az_ldap_bind_dn": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the DN to bind to the LDAP server.",
				Computed:            true,
			},
			"az_ldap_group_attribute": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the attribute name of the LDAP group to check for membership. The authorizing identity must exist as an attribute value in the group.",
				Computed:            true,
			},
			"az_ldap_search_scope": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the scope of the search relative to the input. The default value is subtree.",
				Computed:            true,
			},
			"az_ldap_search_filter": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the LDAP search filter for the search. The default value is <tt>(objectClass=*)</tt> .",
				Computed:            true,
			},
			"az_xacml_version": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the XACML version for communication between the PDP and the AAA policy. The AAA policy acts as an XACML policy enforcement point (PEP). The default value is 2.0.",
				Computed:            true,
			},
			"az_xacml_pep_type": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify how the AAA policy processes the PDP authorization response. The AAA policy acts as an XACML PEP. The default value is deny-based PEP.",
				Computed:            true,
			},
			"az_xacml_use_on_box_pdp": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Specify whether to use the on-box XACML policy decision point (PDP). By default, the AAA policy uses the XACML PDP configuration on the DataPower Gateway.",
				Computed:            true,
			},
			"az_xacml_pdp": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the XACML policy decision point (PDP) configuration.",
				Computed:            true,
			},
			"az_xacml_external_pdp_url": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the URL for the external XACML PDP service. The AAA policy sends the authorization request to and receives the authorization response from this service.",
				Computed:            true,
			},
			"az_xacml_binding_method": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the method to generate the XACML context request. The default value is custom processing.",
				Computed:            true,
			},
			"az_xacml_binding_xsl": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the location of the stylesheet or GatewayScript file that generates the XACML context request. This file maps the AAA result, input message, or both AAA result and input message to the XACML context request.",
				Computed:            true,
			},
			"az_xacml_custom_obligation": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the location of the stylesheet or GatewayScript file that can fulfill XACML obligations. he file must understand the obligations from the PDP and take the appropriate action to fulfill the obligations that are based on the request context. <ul><li>For fulfilled obligations, the output is &lt; <tt>xsl:value-of select=\"true()\"</tt> />.</li><li>For unfulfilled obligations, the output is &lt; <tt>xsl:value-of select=\"false()\"</tt> />.</li></ul>",
				Computed:            true,
			},
			"az_xacml_use_saml2": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Specify whether to use SAML2.0 Profile to communicate with the external PDP service. By default, the PEP does not use SAML2.0 Profile. <ul><li>When enabled, the PEP communicates with the external PDP service by using &lt; <tt>xacml-samlp:XACMLAuthzDecisionQuery</tt> > as defined by SAML2.0 Profile. You can combine this setting with SOAP enveloping if <tt>xacml-samlp:XACMLAuthzDecisionQuery</tt> must be wrapped by a SOAP <tt>Body</tt> element.</li><li>When disabled, the PEP does not use SAML2.0 Profile to communicate with the external PDP service.</li></ul>",
				Computed:            true,
			},
			"az_tam_server": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the IBM Security Access Manager client.",
				Computed:            true,
			},
			"az_tam_default_action": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the default Access Manager action. The default value is T (traverse).",
				Computed:            true,
			},
			"az_tam_action_resource_map": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the location of the XML file that contains the resource-action map. Each entry in the resource-action map defines a PCRE pattern to match the resource, the action to run, and whether to map the action to WebSEAL. This file is in the <tt>local:</tt> or <tt>store:</tt> directory.",
				Computed:            true,
			},
			"az_xacml_use_soap": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Specify whether the external PDP requires a SOAP envelope. By default, a SOAP envelope is not required. If the stylesheet or GatewayScript file for custom binding generates the SOAP envelope, retain the default value.",
				Computed:            true,
			},
			"az_zos_nss_config": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the z/OS NSS client for SAF communication.",
				Computed:            true,
			},
			"az_saf_default_action": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the default action. The default value is r (Read).",
				Computed:            true,
			},
			"az_ldap_attributes": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the list of the extra user attributes to retrieve from the LDAP registry. The attributes that are retrieved from the registry and stored in the var://context/ldap/auxiliary-attributes context variable for future use, such as in the AAA postprocessing phase. To specify multiple attributes, use a comma as the delimiter. For example, enter <tt>email, cn, userPassword</tt> to retrieve these attributes from the registry.",
				Computed:            true,
			},
			"az_skew_time": DataSourceSchema.Int64Attribute{
				MarkdownDescription: "<p>Specify the skew time in seconds. The skew time is the difference between the clock time on the DataPower Gateway and the time on other systems. The default value is 0. When defined, the expiration of the SAML assertion takes the time difference into account.</p><ul><li>For <tt>NotBefore</tt> , validates with <tt>CurrentTime</tt> minus <tt>SkewTime</tt> .</li><li>For <tt>NotOnOrAfter</tt> , validates with <tt>CurrentTime</tt> plus <tt>SkewTime</tt> .</li></ul>",
				Computed:            true,
			},
			"az_oauth_enforce_scope": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Specify how to enforce the scope of the access token. The scope is returned by the server as part of the validation process. By default, the scope is enforced by the resource server. <ul><li>When enabled, the mapped resource is enforced by the DataPower Gateway against the scope.</li><li>When disabled, the remote resource server enforces the scope.</li></ul>",
				Computed:            true,
			},
			"az_oauth_export_headers": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Specify whether to export response attributes that Federated Identity Manager might return a set of response headers. AAA processing places the response attributes as input to the postprocessing phase for use in a custom stylesheet or GatewayScript file. To access the node in the postprocessing input, specify <tt>/container/ResponseAttributes</tt> as the XPath expression. By default, all response attributes are exported to HTTP headers.",
				Computed:            true,
			},
			"az_tam_pac_return": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "<p>Specify whether to return the Access Manager privilege attribute certificate (PAC) token from a successful authorization. You can use the PAC in the postprocessing phase. By default, does not return a PAC token.</p><p>This property is mutually exclusive to the same property in the authentication phase. If you select this property for both authentication and authorization, the setting is automatically cleared for authorization when applied.</p>",
				Computed:            true,
			},
			"az_tam_pac_use": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "<p>Specify whether to use the existing identity or the PAC token for authorization. By default, uses the exiting identity.</p><p>When enabled, use the PAC token that was returned in the authentication or map credentials phase. You can use the PAC token in the postprocessing phase.</p>",
				Computed:            true,
			},
			"az_ldap_read_timeout": DataSourceSchema.Int64Attribute{
				MarkdownDescription: "<p>Specify the duration in seconds to wait for a response from LDAP server before the connection is closed. Enter a value in the range 0 - 86400. The default value is 60. A value of 0 indicates that the connection never times out.</p><p>If you configure an LDAP connection pool and assign it to the AAA policy's XML manager, the AAA policy can use the connection pool. The LDAP read timer of the AAA policy can work with the idle timer of the LDAP connection pool to remove idle connections from the connection pool.</p>",
				Computed:            true,
			},
			"az_ssl_client_config_type": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the TLS profile type to secure connections.",
				Computed:            true,
			},
			"az_ssl_client_profile": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the client profile type to secure connections.",
				Computed:            true,
			},
			"az_ldap_bind_password_alias": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the password alias to bind to the LDAP server.",
				Computed:            true,
			},
			"az_sm_request_type": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the type of request to make. By default, the request against the CA Single Sign-On web agent.",
				Computed:            true,
			},
			"az_sm_cookie_flow": GetDmSMFlowDataSourceSchema("Specify which flows to include the authorization session cookie.", "sm-cookie-flow", ""),
			"az_sm_header_flow": GetDmSMFlowDataSourceSchema("Specify which flows to include the CA Single Sign-On HTTP headers that are generated during authorization. The CA Single Sign-On HTTP headers has a prefix of <tt>SM_</tt> .", "sm-header-flow", ""),
			"az_sm_cookie_attributes": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the cookie attribute policy that allows predefined or custom attributes to be included in CA Single Sign-On cookies.",
				Computed:            true,
			},
			"az_cache_control": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify how to manage the caching of authorization failures. By default, all failures are cached.",
				Computed:            true,
			},
		},
	}
	DmAAAPAuthorizeDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmAAAPAuthorizeDataSourceSchema
}
func GetDmAAAPAuthorizeResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmAAAPAuthorizeResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmAAAPAuthorizeObjectType,
				DmAAAPAuthorizeObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"az_method": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the authorization method.", "method", "").AddStringEnum("anyauthenticated", "passthrough", "tivoli", "netegrity", "oblix", "cleartrust", "custom", "ldap", "saml-authz", "saml-attr", "use-authen-attr", "xacml", "xmlfile", "zosnss", "oauth").AddDefaultValue("anyauthenticated").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("anyauthenticated", "passthrough", "tivoli", "netegrity", "oblix", "cleartrust", "custom", "ldap", "saml-authz", "saml-attr", "use-authen-attr", "xacml", "xmlfile", "zosnss", "oauth"),
				},
				Default: stringdefault.StaticString("anyauthenticated"),
			},
			"az_custom_url": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the stylesheet or GatewayScript file for custom authorization.", "custom-url", "").AddRequiredWhen(DmAAAPAuthorizeAZCustomURLCondVal.String()).AddNotValidWhen(DmAAAPAuthorizeAZCustomURLIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZCustomURLCondVal, DmAAAPAuthorizeAZCustomURLIgnoreVal, false),
				},
			},
			"az_map_url": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the AAA information file.", "xmlfile-url", "").AddRequiredWhen(DmAAAPAuthorizeAZMapURLCondVal.String()).AddNotValidWhen(DmAAAPAuthorizeAZMapURLIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZMapURLCondVal, DmAAAPAuthorizeAZMapURLIgnoreVal, false),
				},
			},
			"az_host": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the host name or IP address of the authorization server.", "remote-host", "").AddRequiredWhen(DmAAAPAuthorizeAZHostCondVal.String()).AddNotValidWhen(DmAAAPAuthorizeAZHostIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZHostCondVal, DmAAAPAuthorizeAZHostIgnoreVal, false),
				},
			},
			"az_port": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the listening port on the authorization server.", "remote-port", "").AddDefaultValue("0").AddRequiredWhen(DmAAAPAuthorizeAZPortCondVal.String()).AddNotValidWhen(DmAAAPAuthorizeAZPortIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					validators.ConditionalRequiredInt64(DmAAAPAuthorizeAZPortCondVal, DmAAAPAuthorizeAZPortIgnoreVal, true),
				},
				Default: int64default.StaticInt64(0),
			},
			"az_ldap_group": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the DN of the required LDAP group.", "ldap-group-dn", "").AddRequiredWhen(DmAAAPAuthorizeAZLDAPGroupCondVal.String()).AddNotValidWhen(DmAAAPAuthorizeAZLDAPGroupIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZLDAPGroupCondVal, DmAAAPAuthorizeAZLDAPGroupIgnoreVal, false),
				},
			},
			"az_valcred": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "", "crypto_val_cred").AddNotValidWhen(DmAAAPAuthorizeAZValcredIgnoreVal.String()).String,
				Optional:            true,
			},
			"az_saml_url": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL of the SAML server.", "saml-server-url", "").AddRequiredWhen(DmAAAPAuthorizeAZSAMLURLCondVal.String()).AddNotValidWhen(DmAAAPAuthorizeAZSAMLURLIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZSAMLURLCondVal, DmAAAPAuthorizeAZSAMLURLIgnoreVal, false),
				},
			},
			"az_saml_type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how to match SAML attributes and values. The default value is any.", "saml-type", "").AddStringEnum("xpath", "any", "all", "any-value", "all-values").AddDefaultValue("any").AddRequiredWhen(DmAAAPAuthorizeAZSAMLTypeCondVal.String()).AddNotValidWhen(DmAAAPAuthorizeAZSAMLTypeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("xpath", "any", "all", "any-value", "all-values"),
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZSAMLTypeCondVal, DmAAAPAuthorizeAZSAMLTypeIgnoreVal, true),
				},
				Default: stringdefault.StaticString("any"),
			},
			"az_saml_xpath": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the XPath expression to run against the SAML statement.", "saml-xpath", "").AddRequiredWhen(DmAAAPAuthorizeAZSAMLXPathCondVal.String()).AddNotValidWhen(DmAAAPAuthorizeAZSAMLXPathIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZSAMLXPathCondVal, DmAAAPAuthorizeAZSAMLXPathIgnoreVal, false),
				},
			},
			"az_saml_name_qualifier": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the value of the <tt>NameQualifier</tt> attribute of the <tt>NameIdentifier</tt> in the generated SAML query. Although the <tt>NameQualifier</tt> attribute is an optional attribute, some SAML implementations require it to be present.", "saml-name-qualifier", "").AddNotValidWhen(DmAAAPAuthorizeAZSAMLNameQualifierIgnoreVal.String()).String,
				Optional:            true,
			},
			"az_cache_allow": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how to control the caching of AAA authorization results. The default value is absolute. A protocol TTL is available only with SAML or OAuth with a Federated Identity Manager endpoint. Federated Identity Manager integration is deprecated.", "cache-type", "").AddStringEnum("absolute", "disabled", "maximum", "minimum").AddDefaultValue("absolute").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("absolute", "disabled", "maximum", "minimum"),
				},
				Default: stringdefault.StaticString("absolute"),
			},
			"az_cache_ttl": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration in seconds to cache authorization decisions. Enter a value in the range 1 - 86400. The default value is 3.", "cache-ttl", "").AddIntegerRange(1, 86400).AddDefaultValue("3").AddNotValidWhen(DmAAAPAuthorizeAZCacheTTLIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 86400),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, DmAAAPAuthorizeAZCacheTTLIgnoreVal, true),
				},
				Default: int64default.StaticInt64(3),
			},
			"az_netegrity_base_uri": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URI sent to CA Single Sign-On (formerly Netegrity SiteMinder) server. The CA Single Sign-On base URI is combined with the host, port, and CA Single Sign-On operation name extension to form the URL for attempting CA Single Sign-On authentication. The URL is of the http://host:port/NetegrityBaseURI/operationNetegrityOpNameExtension form, where NetegrityOpNameExtension is concatenated directly with the operation name.", "netegrity-base-uri", "").AddNotValidWhen(DmAAAPAuthorizeAZNetegrityBaseURIIgnoreVal.String()).String,
				Optional:            true,
			},
			"az_netegrity_op_name_extension": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the extension for URI sent to CA Single Sign-On (formerly Netegrity SiteMinder) server. The CA Single Sign-On base URI is combined with the host, port, and CA Single Sign-On operation name extension to form the URL for attempting CA Single Sign-On authentication. The URL is of the http://host:port/NetegrityBaseURI/operationNetegrityOpNameExtension form, where NetegrityOpNameExtension is concatenated directly with the operation name.", "netegrity-opname-ext", "").AddNotValidWhen(DmAAAPAuthorizeAZNetegrityOpNameExtensionIgnoreVal.String()).String,
				Optional:            true,
			},
			"az_clear_trust_server_url": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL to access the ClearTrust server for authorization.", "cleartrust-server-url", "").AddRequiredWhen(DmAAAPAuthorizeAZClearTrustServerURLCondVal.String()).AddNotValidWhen(DmAAAPAuthorizeAZClearTrustServerURLIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZClearTrustServerURLCondVal, DmAAAPAuthorizeAZClearTrustServerURLIgnoreVal, false),
				},
			},
			"az_saml_version": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the version to use for SAML messages. The default value is 1.1.", "saml-version", "").AddStringEnum("2.0", "1.1", "1.0").AddDefaultValue("1.1").AddNotValidWhen(DmAAAPAuthorizeAZSAMLVersionIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("2.0", "1.1", "1.0"),
					validators.ConditionalRequiredString(validators.Evaluation{}, DmAAAPAuthorizeAZSAMLVersionIgnoreVal, true),
				},
				Default: stringdefault.StaticString("1.1"),
			},
			"az_ldap_load_balance_group": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the load balancer group that contains the LDAP servers.", "ldap-lbgroup", "load_balancer_group").AddNotValidWhen(DmAAAPAuthorizeAZLDAPLoadBalanceGroupIgnoreVal.String()).String,
				Optional:            true,
			},
			"az_ldap_bind_dn": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the DN to bind to the LDAP server.", "ldap-bind-dn", "").AddNotValidWhen(DmAAAPAuthorizeAZLDAPBindDNIgnoreVal.String()).String,
				Optional:            true,
			},
			"az_ldap_group_attribute": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the attribute name of the LDAP group to check for membership. The authorizing identity must exist as an attribute value in the group.", "ldap-group-attr", "").AddDefaultValue("member").AddNotValidWhen(DmAAAPAuthorizeAZLDAPGroupAttributeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("member"),
			},
			"az_ldap_search_scope": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the scope of the search relative to the input. The default value is subtree.", "ldap-search-scope", "").AddStringEnum("subtree", "one-level", "base").AddDefaultValue("subtree").AddNotValidWhen(DmAAAPAuthorizeAZLDAPSearchScopeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("subtree", "one-level", "base"),
					validators.ConditionalRequiredString(validators.Evaluation{}, DmAAAPAuthorizeAZLDAPSearchScopeIgnoreVal, true),
				},
				Default: stringdefault.StaticString("subtree"),
			},
			"az_ldap_search_filter": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the LDAP search filter for the search. The default value is <tt>(objectClass=*)</tt> .", "ldap-search-filter", "").AddDefaultValue("(objectClass=*)").AddNotValidWhen(DmAAAPAuthorizeAZLDAPSearchFilterIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("(objectClass=*)"),
			},
			"az_xacml_version": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the XACML version for communication between the PDP and the AAA policy. The AAA policy acts as an XACML policy enforcement point (PEP). The default value is 2.0.", "xacml-version", "").AddStringEnum("2", "1").AddDefaultValue("2").AddRequiredWhen(DmAAAPAuthorizeAZXACMLVersionCondVal.String()).AddNotValidWhen(DmAAAPAuthorizeAZXACMLVersionIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("2", "1"),
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZXACMLVersionCondVal, DmAAAPAuthorizeAZXACMLVersionIgnoreVal, true),
				},
				Default: stringdefault.StaticString("2"),
			},
			"az_xacml_pep_type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how the AAA policy processes the PDP authorization response. The AAA policy acts as an XACML PEP. The default value is deny-based PEP.", "xacml-pep-type", "").AddStringEnum("base", "deny-biased", "permit-biased").AddDefaultValue("deny-biased").AddRequiredWhen(DmAAAPAuthorizeAZXACMLPEPTypeCondVal.String()).AddNotValidWhen(DmAAAPAuthorizeAZXACMLPEPTypeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("base", "deny-biased", "permit-biased"),
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZXACMLPEPTypeCondVal, DmAAAPAuthorizeAZXACMLPEPTypeIgnoreVal, true),
				},
				Default: stringdefault.StaticString("deny-biased"),
			},
			"az_xacml_use_on_box_pdp": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to use the on-box XACML policy decision point (PDP). By default, the AAA policy uses the XACML PDP configuration on the DataPower Gateway.", "xacml-use-builtin", "").AddDefaultValue("true").AddRequiredWhen(DmAAAPAuthorizeAZXACMLUseOnBoxPDPCondVal.String()).AddNotValidWhen(DmAAAPAuthorizeAZXACMLUseOnBoxPDPIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"az_xacml_pdp": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the XACML policy decision point (PDP) configuration.", "xacml-pdp", "xacml_pdp").AddRequiredWhen(DmAAAPAuthorizeAZXACMLPDPCondVal.String()).AddNotValidWhen(DmAAAPAuthorizeAZXACMLPDPIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZXACMLPDPCondVal, DmAAAPAuthorizeAZXACMLPDPIgnoreVal, false),
				},
			},
			"az_xacml_external_pdp_url": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL for the external XACML PDP service. The AAA policy sends the authorization request to and receives the authorization response from this service.", "xacml-url", "").AddRequiredWhen(DmAAAPAuthorizeAZXACMLExternalPDPUrlCondVal.String()).AddNotValidWhen(DmAAAPAuthorizeAZXACMLExternalPDPUrlIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZXACMLExternalPDPUrlCondVal, DmAAAPAuthorizeAZXACMLExternalPDPUrlIgnoreVal, false),
				},
			},
			"az_xacml_binding_method": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the method to generate the XACML context request. The default value is custom processing.", "xacml-binding-method", "").AddStringEnum("dp-pdp", "custom").AddDefaultValue("custom").AddNotValidWhen(DmAAAPAuthorizeAZXACMLBindingMethodIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("dp-pdp", "custom"),
					validators.ConditionalRequiredString(validators.Evaluation{}, DmAAAPAuthorizeAZXACMLBindingMethodIgnoreVal, true),
				},
				Default: stringdefault.StaticString("custom"),
			},
			"az_xacml_binding_xsl": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the stylesheet or GatewayScript file that generates the XACML context request. This file maps the AAA result, input message, or both AAA result and input message to the XACML context request.", "xacml-binding-custom-url", "").AddRequiredWhen(DmAAAPAuthorizeAZXACMLBindingXSLCondVal.String()).AddNotValidWhen(DmAAAPAuthorizeAZXACMLBindingXSLIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZXACMLBindingXSLCondVal, DmAAAPAuthorizeAZXACMLBindingXSLIgnoreVal, false),
				},
			},
			"az_xacml_custom_obligation": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the stylesheet or GatewayScript file that can fulfill XACML obligations. he file must understand the obligations from the PDP and take the appropriate action to fulfill the obligations that are based on the request context. <ul><li>For fulfilled obligations, the output is &lt; <tt>xsl:value-of select=\"true()\"</tt> />.</li><li>For unfulfilled obligations, the output is &lt; <tt>xsl:value-of select=\"false()\"</tt> />.</li></ul>", "xacml-obligation-custom-url", "").AddNotValidWhen(DmAAAPAuthorizeAZXACMLCustomObligationIgnoreVal.String()).String,
				Optional:            true,
			},
			"az_xacml_use_saml2": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to use SAML2.0 Profile to communicate with the external PDP service. By default, the PEP does not use SAML2.0 Profile. <ul><li>When enabled, the PEP communicates with the external PDP service by using &lt; <tt>xacml-samlp:XACMLAuthzDecisionQuery</tt> > as defined by SAML2.0 Profile. You can combine this setting with SOAP enveloping if <tt>xacml-samlp:XACMLAuthzDecisionQuery</tt> must be wrapped by a SOAP <tt>Body</tt> element.</li><li>When disabled, the PEP does not use SAML2.0 Profile to communicate with the external PDP service.</li></ul>", "xacml-use-saml2", "").AddDefaultValue("false").AddRequiredWhen(DmAAAPAuthorizeAZXACMLUseSAML2CondVal.String()).AddNotValidWhen(DmAAAPAuthorizeAZXACMLUseSAML2IgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"az_tam_server": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the IBM Security Access Manager client.", "tam", "tam").AddRequiredWhen(DmAAAPAuthorizeAZTAMServerCondVal.String()).AddNotValidWhen(DmAAAPAuthorizeAZTAMServerIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZTAMServerCondVal, DmAAAPAuthorizeAZTAMServerIgnoreVal, false),
				},
			},
			"az_tam_default_action": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the default Access Manager action. The default value is T (traverse).", "tam-action-default", "").AddStringEnum("T", "c", "g", "m", "d", "b", "s", "v", "a", "BypassPOP", "tt", "r", "x", "l", "N", "W", "Add", "BypassAuthzRule", "_WebService_i", "_PDMQ_D", "_PDMQ_E").AddDefaultValue("T").AddNotValidWhen(DmAAAPAuthorizeAZTAMDefaultActionIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("T", "c", "g", "m", "d", "b", "s", "v", "a", "BypassPOP", "tt", "r", "x", "l", "N", "W", "Add", "BypassAuthzRule", "_WebService_i", "_PDMQ_D", "_PDMQ_E"),
					validators.ConditionalRequiredString(validators.Evaluation{}, DmAAAPAuthorizeAZTAMDefaultActionIgnoreVal, true),
				},
				Default: stringdefault.StaticString("T"),
			},
			"az_tam_action_resource_map": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the XML file that contains the resource-action map. Each entry in the resource-action map defines a PCRE pattern to match the resource, the action to run, and whether to map the action to WebSEAL. This file is in the <tt>local:</tt> or <tt>store:</tt> directory.", "tam-action-map", "").AddNotValidWhen(DmAAAPAuthorizeAZTAMActionResourceMapIgnoreVal.String()).String,
				Optional:            true,
			},
			"az_xacml_use_soap": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the external PDP requires a SOAP envelope. By default, a SOAP envelope is not required. If the stylesheet or GatewayScript file for custom binding generates the SOAP envelope, retain the default value.", "xacml-use-soap", "").AddDefaultValue("false").AddRequiredWhen(DmAAAPAuthorizeAZXACMLUseSOAPCondVal.String()).AddNotValidWhen(DmAAAPAuthorizeAZXACMLUseSOAPIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"az_zos_nss_config": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the z/OS NSS client for SAF communication.", "zos-nss-az", "zos_nss_client").AddRequiredWhen(DmAAAPAuthorizeAZZOSNSSConfigCondVal.String()).AddNotValidWhen(DmAAAPAuthorizeAZZOSNSSConfigIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZZOSNSSConfigCondVal, DmAAAPAuthorizeAZZOSNSSConfigIgnoreVal, false),
				},
			},
			"az_saf_default_action": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the default action. The default value is r (Read).", "zos-nss-default-action", "").AddStringEnum("r", "u", "a", "c").AddDefaultValue("r").AddNotValidWhen(DmAAAPAuthorizeAZSAFDefaultActionIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("r", "u", "a", "c"),
					validators.ConditionalRequiredString(validators.Evaluation{}, DmAAAPAuthorizeAZSAFDefaultActionIgnoreVal, true),
				},
				Default: stringdefault.StaticString("r"),
			},
			"az_ldap_attributes": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the list of the extra user attributes to retrieve from the LDAP registry. The attributes that are retrieved from the registry and stored in the var://context/ldap/auxiliary-attributes context variable for future use, such as in the AAA postprocessing phase. To specify multiple attributes, use a comma as the delimiter. For example, enter <tt>email, cn, userPassword</tt> to retrieve these attributes from the registry.", "az-ldap-attributes", "").AddNotValidWhen(DmAAAPAuthorizeAZLDAPAttributesIgnoreVal.String()).String,
				Optional:            true,
			},
			"az_skew_time": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the skew time in seconds. The skew time is the difference between the clock time on the DataPower Gateway and the time on other systems. The default value is 0. When defined, the expiration of the SAML assertion takes the time difference into account.</p><ul><li>For <tt>NotBefore</tt> , validates with <tt>CurrentTime</tt> minus <tt>SkewTime</tt> .</li><li>For <tt>NotOnOrAfter</tt> , validates with <tt>CurrentTime</tt> plus <tt>SkewTime</tt> .</li></ul>", "az-skew-time", "").AddDefaultValue("0").AddNotValidWhen(DmAAAPAuthorizeAZSkewTimeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             int64default.StaticInt64(0),
			},
			"az_oauth_enforce_scope": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how to enforce the scope of the access token. The scope is returned by the server as part of the validation process. By default, the scope is enforced by the resource server. <ul><li>When enabled, the mapped resource is enforced by the DataPower Gateway against the scope.</li><li>When disabled, the remote resource server enforces the scope.</li></ul>", "az-oauth-enforce-scope", "").AddDefaultValue("false").AddRequiredWhen(DmAAAPAuthorizeAZOAuthEnforceScopeCondVal.String()).AddNotValidWhen(DmAAAPAuthorizeAZOAuthEnforceScopeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"az_oauth_export_headers": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to export response attributes that Federated Identity Manager might return a set of response headers. AAA processing places the response attributes as input to the postprocessing phase for use in a custom stylesheet or GatewayScript file. To access the node in the postprocessing input, specify <tt>/container/ResponseAttributes</tt> as the XPath expression. By default, all response attributes are exported to HTTP headers.", "az-oauth-export-headers", "").AddDefaultValue("true").AddRequiredWhen(DmAAAPAuthorizeAZOAuthExportHeadersCondVal.String()).AddNotValidWhen(DmAAAPAuthorizeAZOAuthExportHeadersIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"az_tam_pac_return": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to return the Access Manager privilege attribute certificate (PAC) token from a successful authorization. You can use the PAC in the postprocessing phase. By default, does not return a PAC token.</p><p>This property is mutually exclusive to the same property in the authentication phase. If you select this property for both authentication and authorization, the setting is automatically cleared for authorization when applied.</p>", "tam-pac-return", "").AddDefaultValue("false").AddRequiredWhen(DmAAAPAuthorizeAZTAMPACReturnCondVal.String()).AddNotValidWhen(DmAAAPAuthorizeAZTAMPACReturnIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"az_tam_pac_use": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to use the existing identity or the PAC token for authorization. By default, uses the exiting identity.</p><p>When enabled, use the PAC token that was returned in the authentication or map credentials phase. You can use the PAC token in the postprocessing phase.</p>", "use-tam-pac", "").AddDefaultValue("false").AddRequiredWhen(DmAAAPAuthorizeAZTAMPACUseCondVal.String()).AddNotValidWhen(DmAAAPAuthorizeAZTAMPACUseIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"az_ldap_read_timeout": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the duration in seconds to wait for a response from LDAP server before the connection is closed. Enter a value in the range 0 - 86400. The default value is 60. A value of 0 indicates that the connection never times out.</p><p>If you configure an LDAP connection pool and assign it to the AAA policy's XML manager, the AAA policy can use the connection pool. The LDAP read timer of the AAA policy can work with the idle timer of the LDAP connection pool to remove idle connections from the connection pool.</p>", "ldap-readtimeout", "").AddIntegerRange(0, 86400).AddDefaultValue("60").AddNotValidWhen(DmAAAPAuthorizeAZLDAPReadTimeoutIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 86400),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, DmAAAPAuthorizeAZLDAPReadTimeoutIgnoreVal, true),
				},
				Default: int64default.StaticInt64(60),
			},
			"az_ssl_client_config_type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS profile type to secure connections.", "ssl-client-type", "").AddStringEnum("client").AddDefaultValue("client").AddNotValidWhen(DmAAAPAuthorizeAZSSLClientConfigTypeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("client"),
					validators.ConditionalRequiredString(validators.Evaluation{}, DmAAAPAuthorizeAZSSLClientConfigTypeIgnoreVal, true),
				},
				Default: stringdefault.StaticString("client"),
			},
			"az_ssl_client_profile": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the client profile type to secure connections.", "ssl-client", "ssl_client_profile").AddNotValidWhen(DmAAAPAuthorizeAZSSLClientProfileIgnoreVal.String()).String,
				Optional:            true,
			},
			"az_ldap_bind_password_alias": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the password alias to bind to the LDAP server.", "ldap-bind-password-alias", "password_alias").AddNotValidWhen(DmAAAPAuthorizeAZLDAPBindPasswordAliasIgnoreVal.String()).String,
				Optional:            true,
			},
			"az_sm_request_type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the type of request to make. By default, the request against the CA Single Sign-On web agent.", "sm-request-type", "").AddStringEnum("webagent", "webservice").AddDefaultValue("webagent").AddRequiredWhen(DmAAAPAuthorizeAZSMRequestTypeCondVal.String()).AddNotValidWhen(DmAAAPAuthorizeAZSMRequestTypeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("webagent", "webservice"),
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZSMRequestTypeCondVal, DmAAAPAuthorizeAZSMRequestTypeIgnoreVal, true),
				},
				Default: stringdefault.StaticString("webagent"),
			},
			"az_sm_cookie_flow": GetDmSMFlowResourceSchema("Specify which flows to include the authorization session cookie.", "sm-cookie-flow", "", false),
			"az_sm_header_flow": GetDmSMFlowResourceSchema("Specify which flows to include the CA Single Sign-On HTTP headers that are generated during authorization. The CA Single Sign-On HTTP headers has a prefix of <tt>SM_</tt> .", "sm-header-flow", "", false),
			"az_sm_cookie_attributes": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the cookie attribute policy that allows predefined or custom attributes to be included in CA Single Sign-On cookies.", "cookie-attributes", "cookie_attribute_policy").AddNotValidWhen(DmAAAPAuthorizeAZSMCookieAttributesIgnoreVal.String()).String,
				Optional:            true,
			},
			"az_cache_control": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how to manage the caching of authorization failures. By default, all failures are cached.", "cache-control", "").AddStringEnum("default", "disable-all", "disable-ldap-failures").AddDefaultValue("default").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default", "disable-all", "disable-ldap-failures"),
				},
				Default: stringdefault.StaticString("default"),
			},
		},
	}
	DmAAAPAuthorizeResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmAAAPAuthorizeResourceSchema.Required = true
	} else {
		DmAAAPAuthorizeResourceSchema.Optional = true
		DmAAAPAuthorizeResourceSchema.Computed = true
	}
	return DmAAAPAuthorizeResourceSchema
}

func (data DmAAAPAuthorize) IsNull() bool {
	if !data.AzMethod.IsNull() {
		return false
	}
	if !data.AzCustomUrl.IsNull() {
		return false
	}
	if !data.AzMapUrl.IsNull() {
		return false
	}
	if !data.AzHost.IsNull() {
		return false
	}
	if !data.AzPort.IsNull() {
		return false
	}
	if !data.AzLdapGroup.IsNull() {
		return false
	}
	if !data.AzValcred.IsNull() {
		return false
	}
	if !data.AzSamlUrl.IsNull() {
		return false
	}
	if !data.AzSamlType.IsNull() {
		return false
	}
	if !data.AzSamlXpath.IsNull() {
		return false
	}
	if !data.AzSamlNameQualifier.IsNull() {
		return false
	}
	if !data.AzCacheAllow.IsNull() {
		return false
	}
	if !data.AzCacheTtl.IsNull() {
		return false
	}
	if !data.AzNetegrityBaseUri.IsNull() {
		return false
	}
	if !data.AzNetegrityOpNameExtension.IsNull() {
		return false
	}
	if !data.AzClearTrustServerUrl.IsNull() {
		return false
	}
	if !data.AzSamlVersion.IsNull() {
		return false
	}
	if !data.AzLdapLoadBalanceGroup.IsNull() {
		return false
	}
	if !data.AzLdapBindDn.IsNull() {
		return false
	}
	if !data.AzLdapGroupAttribute.IsNull() {
		return false
	}
	if !data.AzLdapSearchScope.IsNull() {
		return false
	}
	if !data.AzLdapSearchFilter.IsNull() {
		return false
	}
	if !data.AzXacmlVersion.IsNull() {
		return false
	}
	if !data.AzXacmlPepType.IsNull() {
		return false
	}
	if !data.AzXacmlUseOnBoxPdp.IsNull() {
		return false
	}
	if !data.AzXacmlPdp.IsNull() {
		return false
	}
	if !data.AzXacmlExternalPdpUrl.IsNull() {
		return false
	}
	if !data.AzXacmlBindingMethod.IsNull() {
		return false
	}
	if !data.AzXacmlBindingXsl.IsNull() {
		return false
	}
	if !data.AzXacmlCustomObligation.IsNull() {
		return false
	}
	if !data.AzXacmlUseSaml2.IsNull() {
		return false
	}
	if !data.AzTamServer.IsNull() {
		return false
	}
	if !data.AzTamDefaultAction.IsNull() {
		return false
	}
	if !data.AzTamActionResourceMap.IsNull() {
		return false
	}
	if !data.AzXacmlUseSoap.IsNull() {
		return false
	}
	if !data.AzZosNssConfig.IsNull() {
		return false
	}
	if !data.AzSafDefaultAction.IsNull() {
		return false
	}
	if !data.AzLdapAttributes.IsNull() {
		return false
	}
	if !data.AzSkewTime.IsNull() {
		return false
	}
	if !data.AzOauthEnforceScope.IsNull() {
		return false
	}
	if !data.AzOauthExportHeaders.IsNull() {
		return false
	}
	if !data.AzTamPacReturn.IsNull() {
		return false
	}
	if !data.AzTamPacUse.IsNull() {
		return false
	}
	if !data.AzLdapReadTimeout.IsNull() {
		return false
	}
	if !data.AzSslClientConfigType.IsNull() {
		return false
	}
	if !data.AzSslClientProfile.IsNull() {
		return false
	}
	if !data.AzLdapBindPasswordAlias.IsNull() {
		return false
	}
	if !data.AzSmRequestType.IsNull() {
		return false
	}
	if data.AzSmCookieFlow != nil {
		if !data.AzSmCookieFlow.IsNull() {
			return false
		}
	}
	if data.AzSmHeaderFlow != nil {
		if !data.AzSmHeaderFlow.IsNull() {
			return false
		}
	}
	if !data.AzSmCookieAttributes.IsNull() {
		return false
	}
	if !data.AzCacheControl.IsNull() {
		return false
	}
	return true
}

func (data DmAAAPAuthorize) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.AzMethod.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZMethod`, data.AzMethod.ValueString())
	}
	if !data.AzCustomUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZCustomURL`, data.AzCustomUrl.ValueString())
	}
	if !data.AzMapUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZMapURL`, data.AzMapUrl.ValueString())
	}
	if !data.AzHost.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZHost`, data.AzHost.ValueString())
	}
	if !data.AzPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZPort`, data.AzPort.ValueInt64())
	}
	if !data.AzLdapGroup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZLDAPGroup`, data.AzLdapGroup.ValueString())
	}
	if !data.AzValcred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZValcred`, data.AzValcred.ValueString())
	}
	if !data.AzSamlUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZSAMLURL`, data.AzSamlUrl.ValueString())
	}
	if !data.AzSamlType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZSAMLType`, data.AzSamlType.ValueString())
	}
	if !data.AzSamlXpath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZSAMLXPath`, data.AzSamlXpath.ValueString())
	}
	if !data.AzSamlNameQualifier.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZSAMLNameQualifier`, data.AzSamlNameQualifier.ValueString())
	}
	if !data.AzCacheAllow.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZCacheAllow`, data.AzCacheAllow.ValueString())
	}
	if !data.AzCacheTtl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZCacheTTL`, data.AzCacheTtl.ValueInt64())
	}
	if !data.AzNetegrityBaseUri.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZNetegrityBaseURI`, data.AzNetegrityBaseUri.ValueString())
	}
	if !data.AzNetegrityOpNameExtension.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZNetegrityOpNameExtension`, data.AzNetegrityOpNameExtension.ValueString())
	}
	if !data.AzClearTrustServerUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZClearTrustServerURL`, data.AzClearTrustServerUrl.ValueString())
	}
	if !data.AzSamlVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZSAMLVersion`, data.AzSamlVersion.ValueString())
	}
	if !data.AzLdapLoadBalanceGroup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZLDAPLoadBalanceGroup`, data.AzLdapLoadBalanceGroup.ValueString())
	}
	if !data.AzLdapBindDn.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZLDAPBindDN`, data.AzLdapBindDn.ValueString())
	}
	if !data.AzLdapGroupAttribute.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZLDAPGroupAttribute`, data.AzLdapGroupAttribute.ValueString())
	}
	if !data.AzLdapSearchScope.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZLDAPSearchScope`, data.AzLdapSearchScope.ValueString())
	}
	if !data.AzLdapSearchFilter.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZLDAPSearchFilter`, data.AzLdapSearchFilter.ValueString())
	}
	if !data.AzXacmlVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZXACMLVersion`, data.AzXacmlVersion.ValueString())
	}
	if !data.AzXacmlPepType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZXACMLPEPType`, data.AzXacmlPepType.ValueString())
	}
	if !data.AzXacmlUseOnBoxPdp.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZXACMLUseOnBoxPDP`, tfutils.StringFromBool(data.AzXacmlUseOnBoxPdp, ""))
	}
	if !data.AzXacmlPdp.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZXACMLPDP`, data.AzXacmlPdp.ValueString())
	}
	if !data.AzXacmlExternalPdpUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZXACMLExternalPDPUrl`, data.AzXacmlExternalPdpUrl.ValueString())
	}
	if !data.AzXacmlBindingMethod.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZXACMLBindingMethod`, data.AzXacmlBindingMethod.ValueString())
	}
	if !data.AzXacmlBindingXsl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZXACMLBindingXSL`, data.AzXacmlBindingXsl.ValueString())
	}
	if !data.AzXacmlCustomObligation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZXACMLCustomObligation`, data.AzXacmlCustomObligation.ValueString())
	}
	if !data.AzXacmlUseSaml2.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZXACMLUseSAML2`, tfutils.StringFromBool(data.AzXacmlUseSaml2, ""))
	}
	if !data.AzTamServer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZTAMServer`, data.AzTamServer.ValueString())
	}
	if !data.AzTamDefaultAction.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZTAMDefaultAction`, data.AzTamDefaultAction.ValueString())
	}
	if !data.AzTamActionResourceMap.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZTAMActionResourceMap`, data.AzTamActionResourceMap.ValueString())
	}
	if !data.AzXacmlUseSoap.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZXACMLUseSOAP`, tfutils.StringFromBool(data.AzXacmlUseSoap, ""))
	}
	if !data.AzZosNssConfig.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZZOSNSSConfig`, data.AzZosNssConfig.ValueString())
	}
	if !data.AzSafDefaultAction.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZSAFDefaultAction`, data.AzSafDefaultAction.ValueString())
	}
	if !data.AzLdapAttributes.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZLDAPAttributes`, data.AzLdapAttributes.ValueString())
	}
	if !data.AzSkewTime.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZSkewTime`, data.AzSkewTime.ValueInt64())
	}
	if !data.AzOauthEnforceScope.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZOAuthEnforceScope`, tfutils.StringFromBool(data.AzOauthEnforceScope, ""))
	}
	if !data.AzOauthExportHeaders.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZOAuthExportHeaders`, tfutils.StringFromBool(data.AzOauthExportHeaders, ""))
	}
	if !data.AzTamPacReturn.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZTAMPACReturn`, tfutils.StringFromBool(data.AzTamPacReturn, ""))
	}
	if !data.AzTamPacUse.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZTAMPACUse`, tfutils.StringFromBool(data.AzTamPacUse, ""))
	}
	if !data.AzLdapReadTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZLDAPReadTimeout`, data.AzLdapReadTimeout.ValueInt64())
	}
	if !data.AzSslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZSSLClientConfigType`, data.AzSslClientConfigType.ValueString())
	}
	if !data.AzSslClientProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZSSLClientProfile`, data.AzSslClientProfile.ValueString())
	}
	if !data.AzLdapBindPasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZLDAPBindPasswordAlias`, data.AzLdapBindPasswordAlias.ValueString())
	}
	if !data.AzSmRequestType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZSMRequestType`, data.AzSmRequestType.ValueString())
	}
	if data.AzSmCookieFlow != nil {
		if !data.AzSmCookieFlow.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`AZSMCookieFlow`, data.AzSmCookieFlow.ToBody(ctx, ""))
		}
	}
	if data.AzSmHeaderFlow != nil {
		if !data.AzSmHeaderFlow.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`AZSMHeaderFlow`, data.AzSmHeaderFlow.ToBody(ctx, ""))
		}
	}
	if !data.AzSmCookieAttributes.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZSMCookieAttributes`, data.AzSmCookieAttributes.ValueString())
	}
	if !data.AzCacheControl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZCacheControl`, data.AzCacheControl.ValueString())
	}
	return body
}

func (data *DmAAAPAuthorize) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `AZMethod`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzMethod = types.StringValue("anyauthenticated")
	}
	if value := res.Get(pathRoot + `AZCustomURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzCustomUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzCustomUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZMapURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzMapUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzMapUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZHost`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzHost = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzHost = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZPort`); value.Exists() {
		data.AzPort = types.Int64Value(value.Int())
	} else {
		data.AzPort = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `AZLDAPGroup`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzLdapGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzLdapGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZValcred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzValcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzValcred = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSAMLURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzSamlUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzSamlUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSAMLType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzSamlType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzSamlType = types.StringValue("any")
	}
	if value := res.Get(pathRoot + `AZSAMLXPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzSamlXpath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzSamlXpath = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSAMLNameQualifier`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzSamlNameQualifier = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzSamlNameQualifier = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZCacheAllow`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzCacheAllow = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzCacheAllow = types.StringValue("absolute")
	}
	if value := res.Get(pathRoot + `AZCacheTTL`); value.Exists() {
		data.AzCacheTtl = types.Int64Value(value.Int())
	} else {
		data.AzCacheTtl = types.Int64Value(3)
	}
	if value := res.Get(pathRoot + `AZNetegrityBaseURI`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzNetegrityBaseUri = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzNetegrityBaseUri = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZNetegrityOpNameExtension`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzNetegrityOpNameExtension = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzNetegrityOpNameExtension = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZClearTrustServerURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzClearTrustServerUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzClearTrustServerUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSAMLVersion`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzSamlVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzSamlVersion = types.StringValue("1.1")
	}
	if value := res.Get(pathRoot + `AZLDAPLoadBalanceGroup`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzLdapLoadBalanceGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzLdapLoadBalanceGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZLDAPBindDN`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzLdapBindDn = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzLdapBindDn = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZLDAPGroupAttribute`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzLdapGroupAttribute = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzLdapGroupAttribute = types.StringValue("member")
	}
	if value := res.Get(pathRoot + `AZLDAPSearchScope`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzLdapSearchScope = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzLdapSearchScope = types.StringValue("subtree")
	}
	if value := res.Get(pathRoot + `AZLDAPSearchFilter`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzLdapSearchFilter = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzLdapSearchFilter = types.StringValue("(objectClass=*)")
	}
	if value := res.Get(pathRoot + `AZXACMLVersion`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzXacmlVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzXacmlVersion = types.StringValue("2")
	}
	if value := res.Get(pathRoot + `AZXACMLPEPType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzXacmlPepType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzXacmlPepType = types.StringValue("deny-biased")
	}
	if value := res.Get(pathRoot + `AZXACMLUseOnBoxPDP`); value.Exists() {
		data.AzXacmlUseOnBoxPdp = tfutils.BoolFromString(value.String())
	} else {
		data.AzXacmlUseOnBoxPdp = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZXACMLPDP`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzXacmlPdp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzXacmlPdp = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZXACMLExternalPDPUrl`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzXacmlExternalPdpUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzXacmlExternalPdpUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZXACMLBindingMethod`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzXacmlBindingMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzXacmlBindingMethod = types.StringValue("custom")
	}
	if value := res.Get(pathRoot + `AZXACMLBindingXSL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzXacmlBindingXsl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzXacmlBindingXsl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZXACMLCustomObligation`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzXacmlCustomObligation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzXacmlCustomObligation = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZXACMLUseSAML2`); value.Exists() {
		data.AzXacmlUseSaml2 = tfutils.BoolFromString(value.String())
	} else {
		data.AzXacmlUseSaml2 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZTAMServer`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzTamServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzTamServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZTAMDefaultAction`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzTamDefaultAction = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzTamDefaultAction = types.StringValue("T")
	}
	if value := res.Get(pathRoot + `AZTAMActionResourceMap`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzTamActionResourceMap = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzTamActionResourceMap = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZXACMLUseSOAP`); value.Exists() {
		data.AzXacmlUseSoap = tfutils.BoolFromString(value.String())
	} else {
		data.AzXacmlUseSoap = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZZOSNSSConfig`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzZosNssConfig = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzZosNssConfig = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSAFDefaultAction`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzSafDefaultAction = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzSafDefaultAction = types.StringValue("r")
	}
	if value := res.Get(pathRoot + `AZLDAPAttributes`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzLdapAttributes = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzLdapAttributes = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSkewTime`); value.Exists() {
		data.AzSkewTime = types.Int64Value(value.Int())
	} else {
		data.AzSkewTime = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `AZOAuthEnforceScope`); value.Exists() {
		data.AzOauthEnforceScope = tfutils.BoolFromString(value.String())
	} else {
		data.AzOauthEnforceScope = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZOAuthExportHeaders`); value.Exists() {
		data.AzOauthExportHeaders = tfutils.BoolFromString(value.String())
	} else {
		data.AzOauthExportHeaders = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZTAMPACReturn`); value.Exists() {
		data.AzTamPacReturn = tfutils.BoolFromString(value.String())
	} else {
		data.AzTamPacReturn = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZTAMPACUse`); value.Exists() {
		data.AzTamPacUse = tfutils.BoolFromString(value.String())
	} else {
		data.AzTamPacUse = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZLDAPReadTimeout`); value.Exists() {
		data.AzLdapReadTimeout = types.Int64Value(value.Int())
	} else {
		data.AzLdapReadTimeout = types.Int64Value(60)
	}
	if value := res.Get(pathRoot + `AZSSLClientConfigType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzSslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzSslClientConfigType = types.StringValue("client")
	}
	if value := res.Get(pathRoot + `AZSSLClientProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzSslClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzSslClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZLDAPBindPasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzLdapBindPasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzLdapBindPasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSMRequestType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzSmRequestType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzSmRequestType = types.StringValue("webagent")
	}
	if value := res.Get(pathRoot + `AZSMCookieFlow`); value.Exists() {
		data.AzSmCookieFlow = &DmSMFlow{}
		data.AzSmCookieFlow.FromBody(ctx, "", value)
	} else {
		data.AzSmCookieFlow = nil
	}
	if value := res.Get(pathRoot + `AZSMHeaderFlow`); value.Exists() {
		data.AzSmHeaderFlow = &DmSMFlow{}
		data.AzSmHeaderFlow.FromBody(ctx, "", value)
	} else {
		data.AzSmHeaderFlow = nil
	}
	if value := res.Get(pathRoot + `AZSMCookieAttributes`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzSmCookieAttributes = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzSmCookieAttributes = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZCacheControl`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzCacheControl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzCacheControl = types.StringValue("default")
	}
}

func (data *DmAAAPAuthorize) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `AZMethod`); value.Exists() && !data.AzMethod.IsNull() {
		data.AzMethod = tfutils.ParseStringFromGJSON(value)
	} else if data.AzMethod.ValueString() != "anyauthenticated" {
		data.AzMethod = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZCustomURL`); value.Exists() && !data.AzCustomUrl.IsNull() {
		data.AzCustomUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzCustomUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZMapURL`); value.Exists() && !data.AzMapUrl.IsNull() {
		data.AzMapUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzMapUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZHost`); value.Exists() && !data.AzHost.IsNull() {
		data.AzHost = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzHost = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZPort`); value.Exists() && !data.AzPort.IsNull() {
		data.AzPort = types.Int64Value(value.Int())
	} else if data.AzPort.ValueInt64() != 0 {
		data.AzPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AZLDAPGroup`); value.Exists() && !data.AzLdapGroup.IsNull() {
		data.AzLdapGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzLdapGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZValcred`); value.Exists() && !data.AzValcred.IsNull() {
		data.AzValcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzValcred = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSAMLURL`); value.Exists() && !data.AzSamlUrl.IsNull() {
		data.AzSamlUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzSamlUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSAMLType`); value.Exists() && !data.AzSamlType.IsNull() {
		data.AzSamlType = tfutils.ParseStringFromGJSON(value)
	} else if data.AzSamlType.ValueString() != "any" {
		data.AzSamlType = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSAMLXPath`); value.Exists() && !data.AzSamlXpath.IsNull() {
		data.AzSamlXpath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzSamlXpath = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSAMLNameQualifier`); value.Exists() && !data.AzSamlNameQualifier.IsNull() {
		data.AzSamlNameQualifier = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzSamlNameQualifier = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZCacheAllow`); value.Exists() && !data.AzCacheAllow.IsNull() {
		data.AzCacheAllow = tfutils.ParseStringFromGJSON(value)
	} else if data.AzCacheAllow.ValueString() != "absolute" {
		data.AzCacheAllow = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZCacheTTL`); value.Exists() && !data.AzCacheTtl.IsNull() {
		data.AzCacheTtl = types.Int64Value(value.Int())
	} else if data.AzCacheTtl.ValueInt64() != 3 {
		data.AzCacheTtl = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AZNetegrityBaseURI`); value.Exists() && !data.AzNetegrityBaseUri.IsNull() {
		data.AzNetegrityBaseUri = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzNetegrityBaseUri = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZNetegrityOpNameExtension`); value.Exists() && !data.AzNetegrityOpNameExtension.IsNull() {
		data.AzNetegrityOpNameExtension = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzNetegrityOpNameExtension = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZClearTrustServerURL`); value.Exists() && !data.AzClearTrustServerUrl.IsNull() {
		data.AzClearTrustServerUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzClearTrustServerUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSAMLVersion`); value.Exists() && !data.AzSamlVersion.IsNull() {
		data.AzSamlVersion = tfutils.ParseStringFromGJSON(value)
	} else if data.AzSamlVersion.ValueString() != "1.1" {
		data.AzSamlVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZLDAPLoadBalanceGroup`); value.Exists() && !data.AzLdapLoadBalanceGroup.IsNull() {
		data.AzLdapLoadBalanceGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzLdapLoadBalanceGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZLDAPBindDN`); value.Exists() && !data.AzLdapBindDn.IsNull() {
		data.AzLdapBindDn = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzLdapBindDn = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZLDAPGroupAttribute`); value.Exists() && !data.AzLdapGroupAttribute.IsNull() {
		data.AzLdapGroupAttribute = tfutils.ParseStringFromGJSON(value)
	} else if data.AzLdapGroupAttribute.ValueString() != "member" {
		data.AzLdapGroupAttribute = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZLDAPSearchScope`); value.Exists() && !data.AzLdapSearchScope.IsNull() {
		data.AzLdapSearchScope = tfutils.ParseStringFromGJSON(value)
	} else if data.AzLdapSearchScope.ValueString() != "subtree" {
		data.AzLdapSearchScope = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZLDAPSearchFilter`); value.Exists() && !data.AzLdapSearchFilter.IsNull() {
		data.AzLdapSearchFilter = tfutils.ParseStringFromGJSON(value)
	} else if data.AzLdapSearchFilter.ValueString() != "(objectClass=*)" {
		data.AzLdapSearchFilter = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZXACMLVersion`); value.Exists() && !data.AzXacmlVersion.IsNull() {
		data.AzXacmlVersion = tfutils.ParseStringFromGJSON(value)
	} else if data.AzXacmlVersion.ValueString() != "2" {
		data.AzXacmlVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZXACMLPEPType`); value.Exists() && !data.AzXacmlPepType.IsNull() {
		data.AzXacmlPepType = tfutils.ParseStringFromGJSON(value)
	} else if data.AzXacmlPepType.ValueString() != "deny-biased" {
		data.AzXacmlPepType = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZXACMLUseOnBoxPDP`); value.Exists() && !data.AzXacmlUseOnBoxPdp.IsNull() {
		data.AzXacmlUseOnBoxPdp = tfutils.BoolFromString(value.String())
	} else if !data.AzXacmlUseOnBoxPdp.ValueBool() {
		data.AzXacmlUseOnBoxPdp = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZXACMLPDP`); value.Exists() && !data.AzXacmlPdp.IsNull() {
		data.AzXacmlPdp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzXacmlPdp = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZXACMLExternalPDPUrl`); value.Exists() && !data.AzXacmlExternalPdpUrl.IsNull() {
		data.AzXacmlExternalPdpUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzXacmlExternalPdpUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZXACMLBindingMethod`); value.Exists() && !data.AzXacmlBindingMethod.IsNull() {
		data.AzXacmlBindingMethod = tfutils.ParseStringFromGJSON(value)
	} else if data.AzXacmlBindingMethod.ValueString() != "custom" {
		data.AzXacmlBindingMethod = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZXACMLBindingXSL`); value.Exists() && !data.AzXacmlBindingXsl.IsNull() {
		data.AzXacmlBindingXsl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzXacmlBindingXsl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZXACMLCustomObligation`); value.Exists() && !data.AzXacmlCustomObligation.IsNull() {
		data.AzXacmlCustomObligation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzXacmlCustomObligation = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZXACMLUseSAML2`); value.Exists() && !data.AzXacmlUseSaml2.IsNull() {
		data.AzXacmlUseSaml2 = tfutils.BoolFromString(value.String())
	} else if data.AzXacmlUseSaml2.ValueBool() {
		data.AzXacmlUseSaml2 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZTAMServer`); value.Exists() && !data.AzTamServer.IsNull() {
		data.AzTamServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzTamServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZTAMDefaultAction`); value.Exists() && !data.AzTamDefaultAction.IsNull() {
		data.AzTamDefaultAction = tfutils.ParseStringFromGJSON(value)
	} else if data.AzTamDefaultAction.ValueString() != "T" {
		data.AzTamDefaultAction = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZTAMActionResourceMap`); value.Exists() && !data.AzTamActionResourceMap.IsNull() {
		data.AzTamActionResourceMap = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzTamActionResourceMap = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZXACMLUseSOAP`); value.Exists() && !data.AzXacmlUseSoap.IsNull() {
		data.AzXacmlUseSoap = tfutils.BoolFromString(value.String())
	} else if data.AzXacmlUseSoap.ValueBool() {
		data.AzXacmlUseSoap = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZZOSNSSConfig`); value.Exists() && !data.AzZosNssConfig.IsNull() {
		data.AzZosNssConfig = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzZosNssConfig = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSAFDefaultAction`); value.Exists() && !data.AzSafDefaultAction.IsNull() {
		data.AzSafDefaultAction = tfutils.ParseStringFromGJSON(value)
	} else if data.AzSafDefaultAction.ValueString() != "r" {
		data.AzSafDefaultAction = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZLDAPAttributes`); value.Exists() && !data.AzLdapAttributes.IsNull() {
		data.AzLdapAttributes = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzLdapAttributes = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSkewTime`); value.Exists() && !data.AzSkewTime.IsNull() {
		data.AzSkewTime = types.Int64Value(value.Int())
	} else if data.AzSkewTime.ValueInt64() != 0 {
		data.AzSkewTime = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AZOAuthEnforceScope`); value.Exists() && !data.AzOauthEnforceScope.IsNull() {
		data.AzOauthEnforceScope = tfutils.BoolFromString(value.String())
	} else if data.AzOauthEnforceScope.ValueBool() {
		data.AzOauthEnforceScope = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZOAuthExportHeaders`); value.Exists() && !data.AzOauthExportHeaders.IsNull() {
		data.AzOauthExportHeaders = tfutils.BoolFromString(value.String())
	} else if !data.AzOauthExportHeaders.ValueBool() {
		data.AzOauthExportHeaders = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZTAMPACReturn`); value.Exists() && !data.AzTamPacReturn.IsNull() {
		data.AzTamPacReturn = tfutils.BoolFromString(value.String())
	} else if data.AzTamPacReturn.ValueBool() {
		data.AzTamPacReturn = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZTAMPACUse`); value.Exists() && !data.AzTamPacUse.IsNull() {
		data.AzTamPacUse = tfutils.BoolFromString(value.String())
	} else if data.AzTamPacUse.ValueBool() {
		data.AzTamPacUse = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZLDAPReadTimeout`); value.Exists() && !data.AzLdapReadTimeout.IsNull() {
		data.AzLdapReadTimeout = types.Int64Value(value.Int())
	} else if data.AzLdapReadTimeout.ValueInt64() != 60 {
		data.AzLdapReadTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AZSSLClientConfigType`); value.Exists() && !data.AzSslClientConfigType.IsNull() {
		data.AzSslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else if data.AzSslClientConfigType.ValueString() != "client" {
		data.AzSslClientConfigType = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSSLClientProfile`); value.Exists() && !data.AzSslClientProfile.IsNull() {
		data.AzSslClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzSslClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZLDAPBindPasswordAlias`); value.Exists() && !data.AzLdapBindPasswordAlias.IsNull() {
		data.AzLdapBindPasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzLdapBindPasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSMRequestType`); value.Exists() && !data.AzSmRequestType.IsNull() {
		data.AzSmRequestType = tfutils.ParseStringFromGJSON(value)
	} else if data.AzSmRequestType.ValueString() != "webagent" {
		data.AzSmRequestType = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSMCookieFlow`); value.Exists() {
		data.AzSmCookieFlow.UpdateFromBody(ctx, "", value)
	} else {
		data.AzSmCookieFlow = nil
	}
	if value := res.Get(pathRoot + `AZSMHeaderFlow`); value.Exists() {
		data.AzSmHeaderFlow.UpdateFromBody(ctx, "", value)
	} else {
		data.AzSmHeaderFlow = nil
	}
	if value := res.Get(pathRoot + `AZSMCookieAttributes`); value.Exists() && !data.AzSmCookieAttributes.IsNull() {
		data.AzSmCookieAttributes = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzSmCookieAttributes = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZCacheControl`); value.Exists() && !data.AzCacheControl.IsNull() {
		data.AzCacheControl = tfutils.ParseStringFromGJSON(value)
	} else if data.AzCacheControl.ValueString() != "default" {
		data.AzCacheControl = types.StringNull()
	}
}
