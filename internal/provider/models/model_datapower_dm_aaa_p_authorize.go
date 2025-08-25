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
	AzldapGroup                types.String `tfsdk:"azldap_group"`
	AzValcred                  types.String `tfsdk:"az_valcred"`
	Azsamlurl                  types.String `tfsdk:"azsamlurl"`
	AzsamlType                 types.String `tfsdk:"azsaml_type"`
	AzsamlxPath                types.String `tfsdk:"azsamlx_path"`
	AzsamlNameQualifier        types.String `tfsdk:"azsaml_name_qualifier"`
	AzCacheAllow               types.String `tfsdk:"az_cache_allow"`
	AzCacheTtl                 types.Int64  `tfsdk:"az_cache_ttl"`
	AzNetegrityBaseUri         types.String `tfsdk:"az_netegrity_base_uri"`
	AzNetegrityOpNameExtension types.String `tfsdk:"az_netegrity_op_name_extension"`
	AzClearTrustServerUrl      types.String `tfsdk:"az_clear_trust_server_url"`
	AzsamlVersion              types.String `tfsdk:"azsaml_version"`
	AzldapLoadBalanceGroup     types.String `tfsdk:"azldap_load_balance_group"`
	AzldapBindDn               types.String `tfsdk:"azldap_bind_dn"`
	AzldapGroupAttribute       types.String `tfsdk:"azldap_group_attribute"`
	AzldapSearchScope          types.String `tfsdk:"azldap_search_scope"`
	AzldapSearchFilter         types.String `tfsdk:"azldap_search_filter"`
	AzxacmlVersion             types.String `tfsdk:"azxacml_version"`
	AzxacmlpepType             types.String `tfsdk:"azxacmlpep_type"`
	AzxacmlUseOnBoxPdp         types.Bool   `tfsdk:"azxacml_use_on_box_pdp"`
	Azxacmlpdp                 types.String `tfsdk:"azxacmlpdp"`
	AzxacmlExternalPdpUrl      types.String `tfsdk:"azxacml_external_pdp_url"`
	AzxacmlBindingMethod       types.String `tfsdk:"azxacml_binding_method"`
	AzxacmlBindingObject       types.String `tfsdk:"azxacml_binding_object"`
	AzxacmlBindingXsl          types.String `tfsdk:"azxacml_binding_xsl"`
	AzxacmlCustomObligation    types.String `tfsdk:"azxacml_custom_obligation"`
	AzxacmlUseSaml2            types.Bool   `tfsdk:"azxacml_use_saml2"`
	AztamServer                types.String `tfsdk:"aztam_server"`
	AztamDefaultAction         types.String `tfsdk:"aztam_default_action"`
	AztamActionResourceMap     types.String `tfsdk:"aztam_action_resource_map"`
	AzxacmlUseSoap             types.Bool   `tfsdk:"azxacml_use_soap"`
	AzzosnssConfig             types.String `tfsdk:"azzosnss_config"`
	AzsafDefaultAction         types.String `tfsdk:"azsaf_default_action"`
	AzldapAttributes           types.String `tfsdk:"azldap_attributes"`
	AzSkewTime                 types.Int64  `tfsdk:"az_skew_time"`
	AzoAuthEnforceScope        types.Bool   `tfsdk:"azo_auth_enforce_scope"`
	AzoAuthExportHeaders       types.Bool   `tfsdk:"azo_auth_export_headers"`
	AztampacReturn             types.Bool   `tfsdk:"aztampac_return"`
	AztampacUse                types.Bool   `tfsdk:"aztampac_use"`
	AzldapReadTimeout          types.Int64  `tfsdk:"azldap_read_timeout"`
	AzsslClientConfigType      types.String `tfsdk:"azssl_client_config_type"`
	AzsslClientProfile         types.String `tfsdk:"azssl_client_profile"`
	AzldapBindPasswordAlias    types.String `tfsdk:"azldap_bind_password_alias"`
	AzsmRequestType            types.String `tfsdk:"azsm_request_type"`
	AzsmCookieFlow             *DmSMFlow    `tfsdk:"azsm_cookie_flow"`
	AzsmHeaderFlow             *DmSMFlow    `tfsdk:"azsm_header_flow"`
	AzsmCookieAttributes       types.String `tfsdk:"azsm_cookie_attributes"`
	AzCacheControl             types.String `tfsdk:"az_cache_control"`
}

var DmAAAPAuthorizeAZCustomURLCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"custom"},
}
var DmAAAPAuthorizeAZMapURLCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"xmlfile"},
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
			Attribute:   "azldap_load_balance_group",
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
			Attribute:   "azldap_load_balance_group",
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
var DmAAAPAuthorizeAZSAMLURLCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"saml-attr", "saml-authz"},
}
var DmAAAPAuthorizeAZSAMLTypeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"saml-attr", "saml-authz", "use-authen-attr"},
}
var DmAAAPAuthorizeAZSAMLXPathCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "azsaml_type",
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
var DmAAAPAuthorizeAZClearTrustServerURLCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"cleartrust"},
}
var DmAAAPAuthorizeAZXACMLVersionCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"xacml"},
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
			Attribute:   "azxacml_version",
			AttrType:    "String",
			AttrDefault: "2",
			Value:       []string{"1.0"},
		},
	},
}
var DmAAAPAuthorizeAZXACMLUseOnBoxPDPCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"xacml"},
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
			Attribute:   "azxacml_use_on_box_pdp",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
	},
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
			Attribute:   "azxacml_use_on_box_pdp",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"false"},
		},
	},
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
			Attribute:   "azxacml_binding_method",
			AttrType:    "String",
			AttrDefault: "custom",
			Value:       []string{"dp-pdp"},
		},
	},
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
			Attribute:   "azxacml_version",
			AttrType:    "String",
			AttrDefault: "2",
			Value:       []string{"2.0"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "azxacml_use_on_box_pdp",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"false"},
		},
	},
}
var DmAAAPAuthorizeAZTAMServerCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
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
			Attribute:   "azxacml_use_on_box_pdp",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"false"},
		},
	},
}
var DmAAAPAuthorizeAZZOSNSSConfigCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"zosnss"},
}
var DmAAAPAuthorizeAZOAuthEnforceScopeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"oauth"},
}
var DmAAAPAuthorizeAZOAuthExportHeadersCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"oauth"},
}
var DmAAAPAuthorizeAZTAMPACReturnCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"tivoli"},
}
var DmAAAPAuthorizeAZTAMPACUseCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"tivoli"},
}
var DmAAAPAuthorizeAZSMRequestTypeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "az_method",
	AttrType:    "String",
	AttrDefault: "anyauthenticated",
	Value:       []string{"netegrity"},
}

var DmAAAPAuthorizeObjectType = map[string]attr.Type{
	"az_method":                      types.StringType,
	"az_custom_url":                  types.StringType,
	"az_map_url":                     types.StringType,
	"az_host":                        types.StringType,
	"az_port":                        types.Int64Type,
	"azldap_group":                   types.StringType,
	"az_valcred":                     types.StringType,
	"azsamlurl":                      types.StringType,
	"azsaml_type":                    types.StringType,
	"azsamlx_path":                   types.StringType,
	"azsaml_name_qualifier":          types.StringType,
	"az_cache_allow":                 types.StringType,
	"az_cache_ttl":                   types.Int64Type,
	"az_netegrity_base_uri":          types.StringType,
	"az_netegrity_op_name_extension": types.StringType,
	"az_clear_trust_server_url":      types.StringType,
	"azsaml_version":                 types.StringType,
	"azldap_load_balance_group":      types.StringType,
	"azldap_bind_dn":                 types.StringType,
	"azldap_group_attribute":         types.StringType,
	"azldap_search_scope":            types.StringType,
	"azldap_search_filter":           types.StringType,
	"azxacml_version":                types.StringType,
	"azxacmlpep_type":                types.StringType,
	"azxacml_use_on_box_pdp":         types.BoolType,
	"azxacmlpdp":                     types.StringType,
	"azxacml_external_pdp_url":       types.StringType,
	"azxacml_binding_method":         types.StringType,
	"azxacml_binding_object":         types.StringType,
	"azxacml_binding_xsl":            types.StringType,
	"azxacml_custom_obligation":      types.StringType,
	"azxacml_use_saml2":              types.BoolType,
	"aztam_server":                   types.StringType,
	"aztam_default_action":           types.StringType,
	"aztam_action_resource_map":      types.StringType,
	"azxacml_use_soap":               types.BoolType,
	"azzosnss_config":                types.StringType,
	"azsaf_default_action":           types.StringType,
	"azldap_attributes":              types.StringType,
	"az_skew_time":                   types.Int64Type,
	"azo_auth_enforce_scope":         types.BoolType,
	"azo_auth_export_headers":        types.BoolType,
	"aztampac_return":                types.BoolType,
	"aztampac_use":                   types.BoolType,
	"azldap_read_timeout":            types.Int64Type,
	"azssl_client_config_type":       types.StringType,
	"azssl_client_profile":           types.StringType,
	"azldap_bind_password_alias":     types.StringType,
	"azsm_request_type":              types.StringType,
	"azsm_cookie_flow":               types.ObjectType{AttrTypes: DmSMFlowObjectType},
	"azsm_header_flow":               types.ObjectType{AttrTypes: DmSMFlowObjectType},
	"azsm_cookie_attributes":         types.StringType,
	"az_cache_control":               types.StringType,
}
var DmAAAPAuthorizeObjectDefault = map[string]attr.Value{
	"az_method":                      types.StringValue("anyauthenticated"),
	"az_custom_url":                  types.StringNull(),
	"az_map_url":                     types.StringNull(),
	"az_host":                        types.StringNull(),
	"az_port":                        types.Int64Value(0),
	"azldap_group":                   types.StringNull(),
	"az_valcred":                     types.StringNull(),
	"azsamlurl":                      types.StringNull(),
	"azsaml_type":                    types.StringValue("any"),
	"azsamlx_path":                   types.StringNull(),
	"azsaml_name_qualifier":          types.StringNull(),
	"az_cache_allow":                 types.StringValue("absolute"),
	"az_cache_ttl":                   types.Int64Value(3),
	"az_netegrity_base_uri":          types.StringNull(),
	"az_netegrity_op_name_extension": types.StringNull(),
	"az_clear_trust_server_url":      types.StringNull(),
	"azsaml_version":                 types.StringValue("1.1"),
	"azldap_load_balance_group":      types.StringNull(),
	"azldap_bind_dn":                 types.StringNull(),
	"azldap_group_attribute":         types.StringValue("member"),
	"azldap_search_scope":            types.StringValue("subtree"),
	"azldap_search_filter":           types.StringValue("(objectClass=*)"),
	"azxacml_version":                types.StringValue("2"),
	"azxacmlpep_type":                types.StringValue("deny-biased"),
	"azxacml_use_on_box_pdp":         types.BoolValue(true),
	"azxacmlpdp":                     types.StringNull(),
	"azxacml_external_pdp_url":       types.StringNull(),
	"azxacml_binding_method":         types.StringValue("custom"),
	"azxacml_binding_object":         types.StringNull(),
	"azxacml_binding_xsl":            types.StringNull(),
	"azxacml_custom_obligation":      types.StringNull(),
	"azxacml_use_saml2":              types.BoolValue(false),
	"aztam_server":                   types.StringNull(),
	"aztam_default_action":           types.StringValue("T"),
	"aztam_action_resource_map":      types.StringNull(),
	"azxacml_use_soap":               types.BoolValue(false),
	"azzosnss_config":                types.StringNull(),
	"azsaf_default_action":           types.StringValue("r"),
	"azldap_attributes":              types.StringNull(),
	"az_skew_time":                   types.Int64Value(0),
	"azo_auth_enforce_scope":         types.BoolValue(false),
	"azo_auth_export_headers":        types.BoolValue(true),
	"aztampac_return":                types.BoolValue(false),
	"aztampac_use":                   types.BoolValue(false),
	"azldap_read_timeout":            types.Int64Value(60),
	"azssl_client_config_type":       types.StringValue("client"),
	"azssl_client_profile":           types.StringNull(),
	"azldap_bind_password_alias":     types.StringNull(),
	"azsm_request_type":              types.StringValue("webagent"),
	"azsm_cookie_flow":               types.ObjectValueMust(DmSMFlowObjectType, DmSMFlowObjectDefault),
	"azsm_header_flow":               types.ObjectValueMust(DmSMFlowObjectType, DmSMFlowObjectDefault),
	"azsm_cookie_attributes":         types.StringNull(),
	"az_cache_control":               types.StringValue("default"),
}

func GetDmAAAPAuthorizeDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmAAAPAuthorizeDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"az_method": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the authorization method.", "method", "").AddStringEnum("anyauthenticated", "passthrough", "tivoli", "netegrity", "oblix", "cleartrust", "custom", "ldap", "saml-authz", "saml-attr", "use-authen-attr", "xacml", "xmlfile", "zosnss", "oauth").AddDefaultValue("anyauthenticated").String,
				Computed:            true,
			},
			"az_custom_url": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the stylesheet or GatewayScript file for custom authorization.", "custom-url", "").String,
				Computed:            true,
			},
			"az_map_url": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the AAA information file.", "xmlfile-url", "").String,
				Computed:            true,
			},
			"az_host": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the host name or IP address of the authorization server.", "remote-host", "").String,
				Computed:            true,
			},
			"az_port": DataSourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the listening port on the authorization server.", "remote-port", "").AddDefaultValue("0").String,
				Computed:            true,
			},
			"azldap_group": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the DN of the required LDAP group.", "ldap-group-dn", "").String,
				Computed:            true,
			},
			"az_valcred": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "", "crypto_val_cred").String,
				Computed:            true,
			},
			"azsamlurl": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL of the SAML server.", "saml-server-url", "").String,
				Computed:            true,
			},
			"azsaml_type": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how to match SAML attributes and values. The default value is any.", "saml-type", "").AddStringEnum("xpath", "any", "all", "any-value", "all-values").AddDefaultValue("any").String,
				Computed:            true,
			},
			"azsamlx_path": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the XPath expression to run against the SAML statement.", "saml-xpath", "").String,
				Computed:            true,
			},
			"azsaml_name_qualifier": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the value of the <tt>NameQualifier</tt> attribute of the <tt>NameIdentifier</tt> in the generated SAML query. Although the <tt>NameQualifier</tt> attribute is an optional attribute, some SAML implementations require it to be present.", "saml-name-qualifier", "").String,
				Computed:            true,
			},
			"az_cache_allow": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how to control the caching of AAA authorization results. The default value is absolute. A protocol TTL is available only with SAML or OAuth with a Federated Identity Manager endpoint. Federated Identity Manager integration is deprecated.", "cache-type", "").AddStringEnum("absolute", "disabled", "maximum", "minimum").AddDefaultValue("absolute").String,
				Computed:            true,
			},
			"az_cache_ttl": DataSourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration in seconds to cache authorization decisions. Enter a value in the range 1 - 86400. The default value is 3.", "cache-ttl", "").AddIntegerRange(1, 86400).AddDefaultValue("3").String,
				Computed:            true,
			},
			"az_netegrity_base_uri": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URI sent to CA Single Sign-On (formerly Netegrity SiteMinder) server. The CA Single Sign-On base URI is combined with the host, port, and CA Single Sign-On operation name extension to form the URL for attempting CA Single Sign-On authentication. The URL is of the http://host:port/NetegrityBaseURI/operationNetegrityOpNameExtension form, where NetegrityOpNameExtension is concatenated directly with the operation name.", "netegrity-base-uri", "").String,
				Computed:            true,
			},
			"az_netegrity_op_name_extension": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the extension for URI sent to CA Single Sign-On (formerly Netegrity SiteMinder) server. The CA Single Sign-On base URI is combined with the host, port, and CA Single Sign-On operation name extension to form the URL for attempting CA Single Sign-On authentication. The URL is of the http://host:port/NetegrityBaseURI/operationNetegrityOpNameExtension form, where NetegrityOpNameExtension is concatenated directly with the operation name.", "netegrity-opname-ext", "").String,
				Computed:            true,
			},
			"az_clear_trust_server_url": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL to access the ClearTrust server for authorization.", "cleartrust-server-url", "").String,
				Computed:            true,
			},
			"azsaml_version": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the version to use for SAML messages. The default value is 1.1.", "saml-version", "").AddStringEnum("2.0", "1.1", "1.0").AddDefaultValue("1.1").String,
				Computed:            true,
			},
			"azldap_load_balance_group": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the load balancer group that contains the LDAP servers.", "ldap-lbgroup", "load_balancer_group").String,
				Computed:            true,
			},
			"azldap_bind_dn": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the DN to bind to the LDAP server.", "ldap-bind-dn", "").String,
				Computed:            true,
			},
			"azldap_group_attribute": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the attribute name of the LDAP group to check for membership. The authorizing identity must exist as an attribute value in the group.", "ldap-group-attr", "").AddDefaultValue("member").String,
				Computed:            true,
			},
			"azldap_search_scope": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the scope of the search relative to the input. The default value is subtree.", "ldap-search-scope", "").AddStringEnum("subtree", "one-level", "base").AddDefaultValue("subtree").String,
				Computed:            true,
			},
			"azldap_search_filter": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the LDAP search filter for the search. The default value is <tt>(objectClass=*)</tt> .", "ldap-search-filter", "").AddDefaultValue("(objectClass=*)").String,
				Computed:            true,
			},
			"azxacml_version": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the XACML version for communication between the PDP and the AAA policy. The AAA policy acts as an XACML policy enforcement point (PEP). The default value is 2.0.", "xacml-version", "").AddStringEnum("2", "1").AddDefaultValue("2").String,
				Computed:            true,
			},
			"azxacmlpep_type": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how the AAA policy processes the PDP authorization response. The AAA policy acts as an XACML PEP. The default value is deny-based PEP.", "xacml-pep-type", "").AddStringEnum("base", "deny-biased", "permit-biased").AddDefaultValue("deny-biased").String,
				Computed:            true,
			},
			"azxacml_use_on_box_pdp": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to use the on-box XACML policy decision point (PDP). By default, the AAA policy uses the XACML PDP configuration on the DataPower Gateway.", "xacml-use-builtin", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"azxacmlpdp": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the XACML policy decision point (PDP) configuration.", "xacml-pdp", "xacml_pdp").String,
				Computed:            true,
			},
			"azxacml_external_pdp_url": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL for the external XACML PDP service. The AAA policy sends the authorization request to and receives the authorization response from this service.", "xacml-url", "").String,
				Computed:            true,
			},
			"azxacml_binding_method": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the method to generate the XACML context request. The default value is custom processing.", "xacml-binding-method", "").AddStringEnum("dp-pdp", "custom").AddDefaultValue("custom").String,
				Computed:            true,
			},
			"azxacml_binding_object": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XACML binding", "xacml-binding-object", "").String,
				Computed:            true,
			},
			"azxacml_binding_xsl": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the stylesheet or GatewayScript file that generates the XACML context request. This file maps the AAA result, input message, or both AAA result and input message to the XACML context request.", "xacml-binding-custom-url", "").String,
				Computed:            true,
			},
			"azxacml_custom_obligation": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the stylesheet or GatewayScript file that can fulfill XACML obligations. he file must understand the obligations from the PDP and take the appropriate action to fulfill the obligations that are based on the request context. <ul><li>For fulfilled obligations, the output is &lt; <tt>xsl:value-of select=\"true()\"</tt> />.</li><li>For unfulfilled obligations, the output is &lt; <tt>xsl:value-of select=\"false()\"</tt> />.</li></ul>", "xacml-obligation-custom-url", "").String,
				Computed:            true,
			},
			"azxacml_use_saml2": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to use SAML2.0 Profile to communicate with the external PDP service. By default, the PEP does not use SAML2.0 Profile. <ul><li>When enabled, the PEP communicates with the external PDP service by using &lt; <tt>xacml-samlp:XACMLAuthzDecisionQuery</tt> > as defined by SAML2.0 Profile. You can combine this setting with SOAP enveloping if <tt>xacml-samlp:XACMLAuthzDecisionQuery</tt> must be wrapped by a SOAP <tt>Body</tt> element.</li><li>When disabled, the PEP does not use SAML2.0 Profile to communicate with the external PDP service.</li></ul>", "xacml-use-saml2", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"aztam_server": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the IBM Security Access Manager client.", "tam", "tam").String,
				Computed:            true,
			},
			"aztam_default_action": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the default Access Manager action. The default value is T (traverse).", "tam-action-default", "").AddStringEnum("T", "c", "g", "m", "d", "b", "s", "v", "a", "BypassPOP", "tt", "r", "x", "l", "N", "W", "Add", "BypassAuthzRule", "_WebService_i", "_PDMQ_D", "_PDMQ_E").AddDefaultValue("T").String,
				Computed:            true,
			},
			"aztam_action_resource_map": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the XML file that contains the resource-action map. Each entry in the resource-action map defines a PCRE pattern to match the resource, the action to run, and whether to map the action to WebSEAL. This file is in the <tt>local:</tt> or <tt>store:</tt> directory.", "tam-action-map", "").String,
				Computed:            true,
			},
			"azxacml_use_soap": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the external PDP requires a SOAP envelope. By default, a SOAP envelope is not required. If the stylesheet or GatewayScript file for custom binding generates the SOAP envelope, retain the default value.", "xacml-use-soap", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"azzosnss_config": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the z/OS NSS client for SAF communication.", "zos-nss-az", "zos_nss_client").String,
				Computed:            true,
			},
			"azsaf_default_action": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the default action. The default value is r (Read).", "zos-nss-default-action", "").AddStringEnum("r", "u", "a", "c").AddDefaultValue("r").String,
				Computed:            true,
			},
			"azldap_attributes": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the list of the extra user attributes to retrieve from the LDAP registry. The attributes that are retrieved from the registry and stored in the var://context/ldap/auxiliary-attributes context variable for future use, such as in the AAA postprocessing phase. To specify multiple attributes, use a comma as the delimiter. For example, enter <tt>email, cn, userPassword</tt> to retrieve these attributes from the registry.", "az-ldap-attributes", "").String,
				Computed:            true,
			},
			"az_skew_time": DataSourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the skew time in seconds. The skew time is the difference between the clock time on the DataPower Gateway and the time on other systems. The default value is 0. When defined, the expiration of the SAML assertion takes the time difference into account.</p><ul><li>For <tt>NotBefore</tt> , validates with <tt>CurrentTime</tt> minus <tt>SkewTime</tt> .</li><li>For <tt>NotOnOrAfter</tt> , validates with <tt>CurrentTime</tt> plus <tt>SkewTime</tt> .</li></ul>", "az-skew-time", "").AddDefaultValue("0").String,
				Computed:            true,
			},
			"azo_auth_enforce_scope": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how to enforce the scope of the access token. The scope is returned by the server as part of the validation process. By default, the scope is enforced by the resource server. <ul><li>When enabled, the mapped resource is enforced by the DataPower Gateway against the scope.</li><li>When disabled, the remote resource server enforces the scope.</li></ul>", "az-oauth-enforce-scope", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"azo_auth_export_headers": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to export response attributes that Federated Identity Manager might return a set of response headers. AAA processing places the response attributes as input to the postprocessing phase for use in a custom stylesheet or GatewayScript file. To access the node in the postprocessing input, specify <tt>/container/ResponseAttributes</tt> as the XPath expression. By default, all response attributes are exported to HTTP headers.", "az-oauth-export-headers", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"aztampac_return": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to return the Access Manager privilege attribute certificate (PAC) token from a successful authorization. You can use the PAC in the postprocessing phase. By default, does not return a PAC token.</p><p>This property is mutually exclusive to the same property in the authentication phase. If you select this property for both authentication and authorization, the setting is automatically cleared for authorization when applied.</p>", "tam-pac-return", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"aztampac_use": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to use the existing identity or the PAC token for authorization. By default, uses the exiting identity.</p><p>When enabled, use the PAC token that was returned in the authentication or map credentials phase. You can use the PAC token in the postprocessing phase.</p>", "use-tam-pac", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"azldap_read_timeout": DataSourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the duration in seconds to wait for a response from LDAP server before the connection is closed. Enter a value in the range 0 - 86400. The default value is 60. A value of 0 indicates that the connection never times out.</p><p>If you configure an LDAP connection pool and assign it to the AAA policy's XML manager, the AAA policy can use the connection pool. The LDAP read timer of the AAA policy can work with the idle timer of the LDAP connection pool to remove idle connections from the connection pool.</p>", "ldap-readtimeout", "").AddIntegerRange(0, 86400).AddDefaultValue("60").String,
				Computed:            true,
			},
			"azssl_client_config_type": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS profile type to secure connections.", "ssl-client-type", "").AddStringEnum("client").AddDefaultValue("client").String,
				Computed:            true,
			},
			"azssl_client_profile": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the client profile type to secure connections.", "ssl-client", "ssl_client_profile").String,
				Computed:            true,
			},
			"azldap_bind_password_alias": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the password alias to bind to the LDAP server.", "ldap-bind-password-alias", "password_alias").String,
				Computed:            true,
			},
			"azsm_request_type": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the type of request to make. By default, the request against the CA Single Sign-On web agent.", "sm-request-type", "").AddStringEnum("webagent", "webservice").AddDefaultValue("webagent").String,
				Computed:            true,
			},
			"azsm_cookie_flow": GetDmSMFlowDataSourceSchema("Specify which flows to include the authorization session cookie.", "sm-cookie-flow", ""),
			"azsm_header_flow": GetDmSMFlowDataSourceSchema("Specify which flows to include the CA Single Sign-On HTTP headers that are generated during authorization. The CA Single Sign-On HTTP headers has a prefix of <tt>SM_</tt> .", "sm-header-flow", ""),
			"azsm_cookie_attributes": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the cookie attribute policy that allows predefined or custom attributes to be included in CA Single Sign-On cookies.", "cookie-attributes", "cookie_attribute_policy").String,
				Computed:            true,
			},
			"az_cache_control": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how to manage the caching of authorization failures. By default, all failures are cached.", "cache-control", "").AddStringEnum("default", "disable-all", "disable-ldap-failures").AddDefaultValue("default").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the stylesheet or GatewayScript file for custom authorization.", "custom-url", "").String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZCustomURLCondVal, validators.Evaluation{}, false),
				},
			},
			"az_map_url": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the AAA information file.", "xmlfile-url", "").String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZMapURLCondVal, validators.Evaluation{}, false),
				},
			},
			"az_host": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the host name or IP address of the authorization server.", "remote-host", "").String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZHostCondVal, validators.Evaluation{}, false),
				},
			},
			"az_port": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the listening port on the authorization server.", "remote-port", "").AddDefaultValue("0").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					validators.ConditionalRequiredInt64(DmAAAPAuthorizeAZPortCondVal, validators.Evaluation{}, true),
				},
				Default: int64default.StaticInt64(0),
			},
			"azldap_group": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the DN of the required LDAP group.", "ldap-group-dn", "").String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZLDAPGroupCondVal, validators.Evaluation{}, false),
				},
			},
			"az_valcred": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "", "crypto_val_cred").String,
				Optional:            true,
			},
			"azsamlurl": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL of the SAML server.", "saml-server-url", "").String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZSAMLURLCondVal, validators.Evaluation{}, false),
				},
			},
			"azsaml_type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how to match SAML attributes and values. The default value is any.", "saml-type", "").AddStringEnum("xpath", "any", "all", "any-value", "all-values").AddDefaultValue("any").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("xpath", "any", "all", "any-value", "all-values"),
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZSAMLTypeCondVal, validators.Evaluation{}, true),
				},
				Default: stringdefault.StaticString("any"),
			},
			"azsamlx_path": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the XPath expression to run against the SAML statement.", "saml-xpath", "").String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZSAMLXPathCondVal, validators.Evaluation{}, false),
				},
			},
			"azsaml_name_qualifier": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the value of the <tt>NameQualifier</tt> attribute of the <tt>NameIdentifier</tt> in the generated SAML query. Although the <tt>NameQualifier</tt> attribute is an optional attribute, some SAML implementations require it to be present.", "saml-name-qualifier", "").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration in seconds to cache authorization decisions. Enter a value in the range 1 - 86400. The default value is 3.", "cache-ttl", "").AddIntegerRange(1, 86400).AddDefaultValue("3").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(3),
			},
			"az_netegrity_base_uri": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URI sent to CA Single Sign-On (formerly Netegrity SiteMinder) server. The CA Single Sign-On base URI is combined with the host, port, and CA Single Sign-On operation name extension to form the URL for attempting CA Single Sign-On authentication. The URL is of the http://host:port/NetegrityBaseURI/operationNetegrityOpNameExtension form, where NetegrityOpNameExtension is concatenated directly with the operation name.", "netegrity-base-uri", "").String,
				Optional:            true,
			},
			"az_netegrity_op_name_extension": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the extension for URI sent to CA Single Sign-On (formerly Netegrity SiteMinder) server. The CA Single Sign-On base URI is combined with the host, port, and CA Single Sign-On operation name extension to form the URL for attempting CA Single Sign-On authentication. The URL is of the http://host:port/NetegrityBaseURI/operationNetegrityOpNameExtension form, where NetegrityOpNameExtension is concatenated directly with the operation name.", "netegrity-opname-ext", "").String,
				Optional:            true,
			},
			"az_clear_trust_server_url": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL to access the ClearTrust server for authorization.", "cleartrust-server-url", "").String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZClearTrustServerURLCondVal, validators.Evaluation{}, false),
				},
			},
			"azsaml_version": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the version to use for SAML messages. The default value is 1.1.", "saml-version", "").AddStringEnum("2.0", "1.1", "1.0").AddDefaultValue("1.1").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("2.0", "1.1", "1.0"),
				},
				Default: stringdefault.StaticString("1.1"),
			},
			"azldap_load_balance_group": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the load balancer group that contains the LDAP servers.", "ldap-lbgroup", "load_balancer_group").String,
				Optional:            true,
			},
			"azldap_bind_dn": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the DN to bind to the LDAP server.", "ldap-bind-dn", "").String,
				Optional:            true,
			},
			"azldap_group_attribute": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the attribute name of the LDAP group to check for membership. The authorizing identity must exist as an attribute value in the group.", "ldap-group-attr", "").AddDefaultValue("member").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("member"),
			},
			"azldap_search_scope": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the scope of the search relative to the input. The default value is subtree.", "ldap-search-scope", "").AddStringEnum("subtree", "one-level", "base").AddDefaultValue("subtree").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("subtree", "one-level", "base"),
				},
				Default: stringdefault.StaticString("subtree"),
			},
			"azldap_search_filter": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the LDAP search filter for the search. The default value is <tt>(objectClass=*)</tt> .", "ldap-search-filter", "").AddDefaultValue("(objectClass=*)").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("(objectClass=*)"),
			},
			"azxacml_version": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the XACML version for communication between the PDP and the AAA policy. The AAA policy acts as an XACML policy enforcement point (PEP). The default value is 2.0.", "xacml-version", "").AddStringEnum("2", "1").AddDefaultValue("2").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("2", "1"),
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZXACMLVersionCondVal, validators.Evaluation{}, true),
				},
				Default: stringdefault.StaticString("2"),
			},
			"azxacmlpep_type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how the AAA policy processes the PDP authorization response. The AAA policy acts as an XACML PEP. The default value is deny-based PEP.", "xacml-pep-type", "").AddStringEnum("base", "deny-biased", "permit-biased").AddDefaultValue("deny-biased").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("base", "deny-biased", "permit-biased"),
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZXACMLPEPTypeCondVal, validators.Evaluation{}, true),
				},
				Default: stringdefault.StaticString("deny-biased"),
			},
			"azxacml_use_on_box_pdp": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to use the on-box XACML policy decision point (PDP). By default, the AAA policy uses the XACML PDP configuration on the DataPower Gateway.", "xacml-use-builtin", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"azxacmlpdp": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the XACML policy decision point (PDP) configuration.", "xacml-pdp", "xacml_pdp").String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZXACMLPDPCondVal, validators.Evaluation{}, false),
				},
			},
			"azxacml_external_pdp_url": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL for the external XACML PDP service. The AAA policy sends the authorization request to and receives the authorization response from this service.", "xacml-url", "").String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZXACMLExternalPDPUrlCondVal, validators.Evaluation{}, false),
				},
			},
			"azxacml_binding_method": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the method to generate the XACML context request. The default value is custom processing.", "xacml-binding-method", "").AddStringEnum("dp-pdp", "custom").AddDefaultValue("custom").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("dp-pdp", "custom"),
				},
				Default: stringdefault.StaticString("custom"),
			},
			"azxacml_binding_object": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XACML binding", "xacml-binding-object", "").String,
				Optional:            true,
			},
			"azxacml_binding_xsl": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the stylesheet or GatewayScript file that generates the XACML context request. This file maps the AAA result, input message, or both AAA result and input message to the XACML context request.", "xacml-binding-custom-url", "").String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZXACMLBindingXSLCondVal, validators.Evaluation{}, false),
				},
			},
			"azxacml_custom_obligation": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the stylesheet or GatewayScript file that can fulfill XACML obligations. he file must understand the obligations from the PDP and take the appropriate action to fulfill the obligations that are based on the request context. <ul><li>For fulfilled obligations, the output is &lt; <tt>xsl:value-of select=\"true()\"</tt> />.</li><li>For unfulfilled obligations, the output is &lt; <tt>xsl:value-of select=\"false()\"</tt> />.</li></ul>", "xacml-obligation-custom-url", "").String,
				Optional:            true,
			},
			"azxacml_use_saml2": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to use SAML2.0 Profile to communicate with the external PDP service. By default, the PEP does not use SAML2.0 Profile. <ul><li>When enabled, the PEP communicates with the external PDP service by using &lt; <tt>xacml-samlp:XACMLAuthzDecisionQuery</tt> > as defined by SAML2.0 Profile. You can combine this setting with SOAP enveloping if <tt>xacml-samlp:XACMLAuthzDecisionQuery</tt> must be wrapped by a SOAP <tt>Body</tt> element.</li><li>When disabled, the PEP does not use SAML2.0 Profile to communicate with the external PDP service.</li></ul>", "xacml-use-saml2", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"aztam_server": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the IBM Security Access Manager client.", "tam", "tam").String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZTAMServerCondVal, validators.Evaluation{}, false),
				},
			},
			"aztam_default_action": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the default Access Manager action. The default value is T (traverse).", "tam-action-default", "").AddStringEnum("T", "c", "g", "m", "d", "b", "s", "v", "a", "BypassPOP", "tt", "r", "x", "l", "N", "W", "Add", "BypassAuthzRule", "_WebService_i", "_PDMQ_D", "_PDMQ_E").AddDefaultValue("T").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("T", "c", "g", "m", "d", "b", "s", "v", "a", "BypassPOP", "tt", "r", "x", "l", "N", "W", "Add", "BypassAuthzRule", "_WebService_i", "_PDMQ_D", "_PDMQ_E"),
				},
				Default: stringdefault.StaticString("T"),
			},
			"aztam_action_resource_map": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the XML file that contains the resource-action map. Each entry in the resource-action map defines a PCRE pattern to match the resource, the action to run, and whether to map the action to WebSEAL. This file is in the <tt>local:</tt> or <tt>store:</tt> directory.", "tam-action-map", "").String,
				Optional:            true,
			},
			"azxacml_use_soap": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the external PDP requires a SOAP envelope. By default, a SOAP envelope is not required. If the stylesheet or GatewayScript file for custom binding generates the SOAP envelope, retain the default value.", "xacml-use-soap", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"azzosnss_config": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the z/OS NSS client for SAF communication.", "zos-nss-az", "zos_nss_client").String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZZOSNSSConfigCondVal, validators.Evaluation{}, false),
				},
			},
			"azsaf_default_action": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the default action. The default value is r (Read).", "zos-nss-default-action", "").AddStringEnum("r", "u", "a", "c").AddDefaultValue("r").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("r", "u", "a", "c"),
				},
				Default: stringdefault.StaticString("r"),
			},
			"azldap_attributes": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the list of the extra user attributes to retrieve from the LDAP registry. The attributes that are retrieved from the registry and stored in the var://context/ldap/auxiliary-attributes context variable for future use, such as in the AAA postprocessing phase. To specify multiple attributes, use a comma as the delimiter. For example, enter <tt>email, cn, userPassword</tt> to retrieve these attributes from the registry.", "az-ldap-attributes", "").String,
				Optional:            true,
			},
			"az_skew_time": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the skew time in seconds. The skew time is the difference between the clock time on the DataPower Gateway and the time on other systems. The default value is 0. When defined, the expiration of the SAML assertion takes the time difference into account.</p><ul><li>For <tt>NotBefore</tt> , validates with <tt>CurrentTime</tt> minus <tt>SkewTime</tt> .</li><li>For <tt>NotOnOrAfter</tt> , validates with <tt>CurrentTime</tt> plus <tt>SkewTime</tt> .</li></ul>", "az-skew-time", "").AddDefaultValue("0").String,
				Computed:            true,
				Optional:            true,
				Default:             int64default.StaticInt64(0),
			},
			"azo_auth_enforce_scope": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how to enforce the scope of the access token. The scope is returned by the server as part of the validation process. By default, the scope is enforced by the resource server. <ul><li>When enabled, the mapped resource is enforced by the DataPower Gateway against the scope.</li><li>When disabled, the remote resource server enforces the scope.</li></ul>", "az-oauth-enforce-scope", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"azo_auth_export_headers": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to export response attributes that Federated Identity Manager might return a set of response headers. AAA processing places the response attributes as input to the postprocessing phase for use in a custom stylesheet or GatewayScript file. To access the node in the postprocessing input, specify <tt>/container/ResponseAttributes</tt> as the XPath expression. By default, all response attributes are exported to HTTP headers.", "az-oauth-export-headers", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"aztampac_return": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to return the Access Manager privilege attribute certificate (PAC) token from a successful authorization. You can use the PAC in the postprocessing phase. By default, does not return a PAC token.</p><p>This property is mutually exclusive to the same property in the authentication phase. If you select this property for both authentication and authorization, the setting is automatically cleared for authorization when applied.</p>", "tam-pac-return", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"aztampac_use": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to use the existing identity or the PAC token for authorization. By default, uses the exiting identity.</p><p>When enabled, use the PAC token that was returned in the authentication or map credentials phase. You can use the PAC token in the postprocessing phase.</p>", "use-tam-pac", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"azldap_read_timeout": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the duration in seconds to wait for a response from LDAP server before the connection is closed. Enter a value in the range 0 - 86400. The default value is 60. A value of 0 indicates that the connection never times out.</p><p>If you configure an LDAP connection pool and assign it to the AAA policy's XML manager, the AAA policy can use the connection pool. The LDAP read timer of the AAA policy can work with the idle timer of the LDAP connection pool to remove idle connections from the connection pool.</p>", "ldap-readtimeout", "").AddIntegerRange(0, 86400).AddDefaultValue("60").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 86400),
				},
				Default: int64default.StaticInt64(60),
			},
			"azssl_client_config_type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS profile type to secure connections.", "ssl-client-type", "").AddStringEnum("client").AddDefaultValue("client").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("client"),
				},
				Default: stringdefault.StaticString("client"),
			},
			"azssl_client_profile": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the client profile type to secure connections.", "ssl-client", "ssl_client_profile").String,
				Optional:            true,
			},
			"azldap_bind_password_alias": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the password alias to bind to the LDAP server.", "ldap-bind-password-alias", "password_alias").String,
				Optional:            true,
			},
			"azsm_request_type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the type of request to make. By default, the request against the CA Single Sign-On web agent.", "sm-request-type", "").AddStringEnum("webagent", "webservice").AddDefaultValue("webagent").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("webagent", "webservice"),
					validators.ConditionalRequiredString(DmAAAPAuthorizeAZSMRequestTypeCondVal, validators.Evaluation{}, true),
				},
				Default: stringdefault.StaticString("webagent"),
			},
			"azsm_cookie_flow": GetDmSMFlowResourceSchema("Specify which flows to include the authorization session cookie.", "sm-cookie-flow", "", false),
			"azsm_header_flow": GetDmSMFlowResourceSchema("Specify which flows to include the CA Single Sign-On HTTP headers that are generated during authorization. The CA Single Sign-On HTTP headers has a prefix of <tt>SM_</tt> .", "sm-header-flow", "", false),
			"azsm_cookie_attributes": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the cookie attribute policy that allows predefined or custom attributes to be included in CA Single Sign-On cookies.", "cookie-attributes", "cookie_attribute_policy").String,
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
	if !data.AzldapGroup.IsNull() {
		return false
	}
	if !data.AzValcred.IsNull() {
		return false
	}
	if !data.Azsamlurl.IsNull() {
		return false
	}
	if !data.AzsamlType.IsNull() {
		return false
	}
	if !data.AzsamlxPath.IsNull() {
		return false
	}
	if !data.AzsamlNameQualifier.IsNull() {
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
	if !data.AzsamlVersion.IsNull() {
		return false
	}
	if !data.AzldapLoadBalanceGroup.IsNull() {
		return false
	}
	if !data.AzldapBindDn.IsNull() {
		return false
	}
	if !data.AzldapGroupAttribute.IsNull() {
		return false
	}
	if !data.AzldapSearchScope.IsNull() {
		return false
	}
	if !data.AzldapSearchFilter.IsNull() {
		return false
	}
	if !data.AzxacmlVersion.IsNull() {
		return false
	}
	if !data.AzxacmlpepType.IsNull() {
		return false
	}
	if !data.AzxacmlUseOnBoxPdp.IsNull() {
		return false
	}
	if !data.Azxacmlpdp.IsNull() {
		return false
	}
	if !data.AzxacmlExternalPdpUrl.IsNull() {
		return false
	}
	if !data.AzxacmlBindingMethod.IsNull() {
		return false
	}
	if !data.AzxacmlBindingObject.IsNull() {
		return false
	}
	if !data.AzxacmlBindingXsl.IsNull() {
		return false
	}
	if !data.AzxacmlCustomObligation.IsNull() {
		return false
	}
	if !data.AzxacmlUseSaml2.IsNull() {
		return false
	}
	if !data.AztamServer.IsNull() {
		return false
	}
	if !data.AztamDefaultAction.IsNull() {
		return false
	}
	if !data.AztamActionResourceMap.IsNull() {
		return false
	}
	if !data.AzxacmlUseSoap.IsNull() {
		return false
	}
	if !data.AzzosnssConfig.IsNull() {
		return false
	}
	if !data.AzsafDefaultAction.IsNull() {
		return false
	}
	if !data.AzldapAttributes.IsNull() {
		return false
	}
	if !data.AzSkewTime.IsNull() {
		return false
	}
	if !data.AzoAuthEnforceScope.IsNull() {
		return false
	}
	if !data.AzoAuthExportHeaders.IsNull() {
		return false
	}
	if !data.AztampacReturn.IsNull() {
		return false
	}
	if !data.AztampacUse.IsNull() {
		return false
	}
	if !data.AzldapReadTimeout.IsNull() {
		return false
	}
	if !data.AzsslClientConfigType.IsNull() {
		return false
	}
	if !data.AzsslClientProfile.IsNull() {
		return false
	}
	if !data.AzldapBindPasswordAlias.IsNull() {
		return false
	}
	if !data.AzsmRequestType.IsNull() {
		return false
	}
	if data.AzsmCookieFlow != nil {
		if !data.AzsmCookieFlow.IsNull() {
			return false
		}
	}
	if data.AzsmHeaderFlow != nil {
		if !data.AzsmHeaderFlow.IsNull() {
			return false
		}
	}
	if !data.AzsmCookieAttributes.IsNull() {
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
	if !data.AzldapGroup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZLDAPGroup`, data.AzldapGroup.ValueString())
	}
	if !data.AzValcred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZValcred`, data.AzValcred.ValueString())
	}
	if !data.Azsamlurl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZSAMLURL`, data.Azsamlurl.ValueString())
	}
	if !data.AzsamlType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZSAMLType`, data.AzsamlType.ValueString())
	}
	if !data.AzsamlxPath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZSAMLXPath`, data.AzsamlxPath.ValueString())
	}
	if !data.AzsamlNameQualifier.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZSAMLNameQualifier`, data.AzsamlNameQualifier.ValueString())
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
	if !data.AzsamlVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZSAMLVersion`, data.AzsamlVersion.ValueString())
	}
	if !data.AzldapLoadBalanceGroup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZLDAPLoadBalanceGroup`, data.AzldapLoadBalanceGroup.ValueString())
	}
	if !data.AzldapBindDn.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZLDAPBindDN`, data.AzldapBindDn.ValueString())
	}
	if !data.AzldapGroupAttribute.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZLDAPGroupAttribute`, data.AzldapGroupAttribute.ValueString())
	}
	if !data.AzldapSearchScope.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZLDAPSearchScope`, data.AzldapSearchScope.ValueString())
	}
	if !data.AzldapSearchFilter.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZLDAPSearchFilter`, data.AzldapSearchFilter.ValueString())
	}
	if !data.AzxacmlVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZXACMLVersion`, data.AzxacmlVersion.ValueString())
	}
	if !data.AzxacmlpepType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZXACMLPEPType`, data.AzxacmlpepType.ValueString())
	}
	if !data.AzxacmlUseOnBoxPdp.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZXACMLUseOnBoxPDP`, tfutils.StringFromBool(data.AzxacmlUseOnBoxPdp, ""))
	}
	if !data.Azxacmlpdp.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZXACMLPDP`, data.Azxacmlpdp.ValueString())
	}
	if !data.AzxacmlExternalPdpUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZXACMLExternalPDPUrl`, data.AzxacmlExternalPdpUrl.ValueString())
	}
	if !data.AzxacmlBindingMethod.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZXACMLBindingMethod`, data.AzxacmlBindingMethod.ValueString())
	}
	if !data.AzxacmlBindingObject.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZXACMLBindingObject`, data.AzxacmlBindingObject.ValueString())
	}
	if !data.AzxacmlBindingXsl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZXACMLBindingXSL`, data.AzxacmlBindingXsl.ValueString())
	}
	if !data.AzxacmlCustomObligation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZXACMLCustomObligation`, data.AzxacmlCustomObligation.ValueString())
	}
	if !data.AzxacmlUseSaml2.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZXACMLUseSAML2`, tfutils.StringFromBool(data.AzxacmlUseSaml2, ""))
	}
	if !data.AztamServer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZTAMServer`, data.AztamServer.ValueString())
	}
	if !data.AztamDefaultAction.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZTAMDefaultAction`, data.AztamDefaultAction.ValueString())
	}
	if !data.AztamActionResourceMap.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZTAMActionResourceMap`, data.AztamActionResourceMap.ValueString())
	}
	if !data.AzxacmlUseSoap.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZXACMLUseSOAP`, tfutils.StringFromBool(data.AzxacmlUseSoap, ""))
	}
	if !data.AzzosnssConfig.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZZOSNSSConfig`, data.AzzosnssConfig.ValueString())
	}
	if !data.AzsafDefaultAction.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZSAFDefaultAction`, data.AzsafDefaultAction.ValueString())
	}
	if !data.AzldapAttributes.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZLDAPAttributes`, data.AzldapAttributes.ValueString())
	}
	if !data.AzSkewTime.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZSkewTime`, data.AzSkewTime.ValueInt64())
	}
	if !data.AzoAuthEnforceScope.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZOAuthEnforceScope`, tfutils.StringFromBool(data.AzoAuthEnforceScope, ""))
	}
	if !data.AzoAuthExportHeaders.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZOAuthExportHeaders`, tfutils.StringFromBool(data.AzoAuthExportHeaders, ""))
	}
	if !data.AztampacReturn.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZTAMPACReturn`, tfutils.StringFromBool(data.AztampacReturn, ""))
	}
	if !data.AztampacUse.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZTAMPACUse`, tfutils.StringFromBool(data.AztampacUse, ""))
	}
	if !data.AzldapReadTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZLDAPReadTimeout`, data.AzldapReadTimeout.ValueInt64())
	}
	if !data.AzsslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZSSLClientConfigType`, data.AzsslClientConfigType.ValueString())
	}
	if !data.AzsslClientProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZSSLClientProfile`, data.AzsslClientProfile.ValueString())
	}
	if !data.AzldapBindPasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZLDAPBindPasswordAlias`, data.AzldapBindPasswordAlias.ValueString())
	}
	if !data.AzsmRequestType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZSMRequestType`, data.AzsmRequestType.ValueString())
	}
	if data.AzsmCookieFlow != nil {
		if !data.AzsmCookieFlow.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`AZSMCookieFlow`, data.AzsmCookieFlow.ToBody(ctx, ""))
		}
	}
	if data.AzsmHeaderFlow != nil {
		if !data.AzsmHeaderFlow.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`AZSMHeaderFlow`, data.AzsmHeaderFlow.ToBody(ctx, ""))
		}
	}
	if !data.AzsmCookieAttributes.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZSMCookieAttributes`, data.AzsmCookieAttributes.ValueString())
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
		data.AzldapGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzldapGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZValcred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzValcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzValcred = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSAMLURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Azsamlurl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Azsamlurl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSAMLType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzsamlType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzsamlType = types.StringValue("any")
	}
	if value := res.Get(pathRoot + `AZSAMLXPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzsamlxPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzsamlxPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSAMLNameQualifier`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzsamlNameQualifier = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzsamlNameQualifier = types.StringNull()
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
		data.AzsamlVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzsamlVersion = types.StringValue("1.1")
	}
	if value := res.Get(pathRoot + `AZLDAPLoadBalanceGroup`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzldapLoadBalanceGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzldapLoadBalanceGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZLDAPBindDN`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzldapBindDn = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzldapBindDn = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZLDAPGroupAttribute`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzldapGroupAttribute = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzldapGroupAttribute = types.StringValue("member")
	}
	if value := res.Get(pathRoot + `AZLDAPSearchScope`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzldapSearchScope = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzldapSearchScope = types.StringValue("subtree")
	}
	if value := res.Get(pathRoot + `AZLDAPSearchFilter`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzldapSearchFilter = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzldapSearchFilter = types.StringValue("(objectClass=*)")
	}
	if value := res.Get(pathRoot + `AZXACMLVersion`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzxacmlVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzxacmlVersion = types.StringValue("2")
	}
	if value := res.Get(pathRoot + `AZXACMLPEPType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzxacmlpepType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzxacmlpepType = types.StringValue("deny-biased")
	}
	if value := res.Get(pathRoot + `AZXACMLUseOnBoxPDP`); value.Exists() {
		data.AzxacmlUseOnBoxPdp = tfutils.BoolFromString(value.String())
	} else {
		data.AzxacmlUseOnBoxPdp = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZXACMLPDP`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Azxacmlpdp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Azxacmlpdp = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZXACMLExternalPDPUrl`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzxacmlExternalPdpUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzxacmlExternalPdpUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZXACMLBindingMethod`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzxacmlBindingMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzxacmlBindingMethod = types.StringValue("custom")
	}
	if value := res.Get(pathRoot + `AZXACMLBindingObject`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzxacmlBindingObject = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzxacmlBindingObject = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZXACMLBindingXSL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzxacmlBindingXsl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzxacmlBindingXsl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZXACMLCustomObligation`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzxacmlCustomObligation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzxacmlCustomObligation = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZXACMLUseSAML2`); value.Exists() {
		data.AzxacmlUseSaml2 = tfutils.BoolFromString(value.String())
	} else {
		data.AzxacmlUseSaml2 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZTAMServer`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AztamServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AztamServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZTAMDefaultAction`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AztamDefaultAction = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AztamDefaultAction = types.StringValue("T")
	}
	if value := res.Get(pathRoot + `AZTAMActionResourceMap`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AztamActionResourceMap = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AztamActionResourceMap = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZXACMLUseSOAP`); value.Exists() {
		data.AzxacmlUseSoap = tfutils.BoolFromString(value.String())
	} else {
		data.AzxacmlUseSoap = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZZOSNSSConfig`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzzosnssConfig = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzzosnssConfig = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSAFDefaultAction`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzsafDefaultAction = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzsafDefaultAction = types.StringValue("r")
	}
	if value := res.Get(pathRoot + `AZLDAPAttributes`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzldapAttributes = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzldapAttributes = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSkewTime`); value.Exists() {
		data.AzSkewTime = types.Int64Value(value.Int())
	} else {
		data.AzSkewTime = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `AZOAuthEnforceScope`); value.Exists() {
		data.AzoAuthEnforceScope = tfutils.BoolFromString(value.String())
	} else {
		data.AzoAuthEnforceScope = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZOAuthExportHeaders`); value.Exists() {
		data.AzoAuthExportHeaders = tfutils.BoolFromString(value.String())
	} else {
		data.AzoAuthExportHeaders = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZTAMPACReturn`); value.Exists() {
		data.AztampacReturn = tfutils.BoolFromString(value.String())
	} else {
		data.AztampacReturn = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZTAMPACUse`); value.Exists() {
		data.AztampacUse = tfutils.BoolFromString(value.String())
	} else {
		data.AztampacUse = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZLDAPReadTimeout`); value.Exists() {
		data.AzldapReadTimeout = types.Int64Value(value.Int())
	} else {
		data.AzldapReadTimeout = types.Int64Value(60)
	}
	if value := res.Get(pathRoot + `AZSSLClientConfigType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzsslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzsslClientConfigType = types.StringValue("client")
	}
	if value := res.Get(pathRoot + `AZSSLClientProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzsslClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzsslClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZLDAPBindPasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzldapBindPasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzldapBindPasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSMRequestType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzsmRequestType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzsmRequestType = types.StringValue("webagent")
	}
	if value := res.Get(pathRoot + `AZSMCookieFlow`); value.Exists() {
		data.AzsmCookieFlow = &DmSMFlow{}
		data.AzsmCookieFlow.FromBody(ctx, "", value)
	} else {
		data.AzsmCookieFlow = nil
	}
	if value := res.Get(pathRoot + `AZSMHeaderFlow`); value.Exists() {
		data.AzsmHeaderFlow = &DmSMFlow{}
		data.AzsmHeaderFlow.FromBody(ctx, "", value)
	} else {
		data.AzsmHeaderFlow = nil
	}
	if value := res.Get(pathRoot + `AZSMCookieAttributes`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzsmCookieAttributes = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzsmCookieAttributes = types.StringNull()
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
	if value := res.Get(pathRoot + `AZLDAPGroup`); value.Exists() && !data.AzldapGroup.IsNull() {
		data.AzldapGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzldapGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZValcred`); value.Exists() && !data.AzValcred.IsNull() {
		data.AzValcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzValcred = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSAMLURL`); value.Exists() && !data.Azsamlurl.IsNull() {
		data.Azsamlurl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Azsamlurl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSAMLType`); value.Exists() && !data.AzsamlType.IsNull() {
		data.AzsamlType = tfutils.ParseStringFromGJSON(value)
	} else if data.AzsamlType.ValueString() != "any" {
		data.AzsamlType = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSAMLXPath`); value.Exists() && !data.AzsamlxPath.IsNull() {
		data.AzsamlxPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzsamlxPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSAMLNameQualifier`); value.Exists() && !data.AzsamlNameQualifier.IsNull() {
		data.AzsamlNameQualifier = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzsamlNameQualifier = types.StringNull()
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
	if value := res.Get(pathRoot + `AZSAMLVersion`); value.Exists() && !data.AzsamlVersion.IsNull() {
		data.AzsamlVersion = tfutils.ParseStringFromGJSON(value)
	} else if data.AzsamlVersion.ValueString() != "1.1" {
		data.AzsamlVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZLDAPLoadBalanceGroup`); value.Exists() && !data.AzldapLoadBalanceGroup.IsNull() {
		data.AzldapLoadBalanceGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzldapLoadBalanceGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZLDAPBindDN`); value.Exists() && !data.AzldapBindDn.IsNull() {
		data.AzldapBindDn = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzldapBindDn = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZLDAPGroupAttribute`); value.Exists() && !data.AzldapGroupAttribute.IsNull() {
		data.AzldapGroupAttribute = tfutils.ParseStringFromGJSON(value)
	} else if data.AzldapGroupAttribute.ValueString() != "member" {
		data.AzldapGroupAttribute = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZLDAPSearchScope`); value.Exists() && !data.AzldapSearchScope.IsNull() {
		data.AzldapSearchScope = tfutils.ParseStringFromGJSON(value)
	} else if data.AzldapSearchScope.ValueString() != "subtree" {
		data.AzldapSearchScope = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZLDAPSearchFilter`); value.Exists() && !data.AzldapSearchFilter.IsNull() {
		data.AzldapSearchFilter = tfutils.ParseStringFromGJSON(value)
	} else if data.AzldapSearchFilter.ValueString() != "(objectClass=*)" {
		data.AzldapSearchFilter = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZXACMLVersion`); value.Exists() && !data.AzxacmlVersion.IsNull() {
		data.AzxacmlVersion = tfutils.ParseStringFromGJSON(value)
	} else if data.AzxacmlVersion.ValueString() != "2" {
		data.AzxacmlVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZXACMLPEPType`); value.Exists() && !data.AzxacmlpepType.IsNull() {
		data.AzxacmlpepType = tfutils.ParseStringFromGJSON(value)
	} else if data.AzxacmlpepType.ValueString() != "deny-biased" {
		data.AzxacmlpepType = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZXACMLUseOnBoxPDP`); value.Exists() && !data.AzxacmlUseOnBoxPdp.IsNull() {
		data.AzxacmlUseOnBoxPdp = tfutils.BoolFromString(value.String())
	} else if !data.AzxacmlUseOnBoxPdp.ValueBool() {
		data.AzxacmlUseOnBoxPdp = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZXACMLPDP`); value.Exists() && !data.Azxacmlpdp.IsNull() {
		data.Azxacmlpdp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Azxacmlpdp = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZXACMLExternalPDPUrl`); value.Exists() && !data.AzxacmlExternalPdpUrl.IsNull() {
		data.AzxacmlExternalPdpUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzxacmlExternalPdpUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZXACMLBindingMethod`); value.Exists() && !data.AzxacmlBindingMethod.IsNull() {
		data.AzxacmlBindingMethod = tfutils.ParseStringFromGJSON(value)
	} else if data.AzxacmlBindingMethod.ValueString() != "custom" {
		data.AzxacmlBindingMethod = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZXACMLBindingObject`); value.Exists() && !data.AzxacmlBindingObject.IsNull() {
		data.AzxacmlBindingObject = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzxacmlBindingObject = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZXACMLBindingXSL`); value.Exists() && !data.AzxacmlBindingXsl.IsNull() {
		data.AzxacmlBindingXsl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzxacmlBindingXsl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZXACMLCustomObligation`); value.Exists() && !data.AzxacmlCustomObligation.IsNull() {
		data.AzxacmlCustomObligation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzxacmlCustomObligation = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZXACMLUseSAML2`); value.Exists() && !data.AzxacmlUseSaml2.IsNull() {
		data.AzxacmlUseSaml2 = tfutils.BoolFromString(value.String())
	} else if data.AzxacmlUseSaml2.ValueBool() {
		data.AzxacmlUseSaml2 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZTAMServer`); value.Exists() && !data.AztamServer.IsNull() {
		data.AztamServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AztamServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZTAMDefaultAction`); value.Exists() && !data.AztamDefaultAction.IsNull() {
		data.AztamDefaultAction = tfutils.ParseStringFromGJSON(value)
	} else if data.AztamDefaultAction.ValueString() != "T" {
		data.AztamDefaultAction = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZTAMActionResourceMap`); value.Exists() && !data.AztamActionResourceMap.IsNull() {
		data.AztamActionResourceMap = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AztamActionResourceMap = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZXACMLUseSOAP`); value.Exists() && !data.AzxacmlUseSoap.IsNull() {
		data.AzxacmlUseSoap = tfutils.BoolFromString(value.String())
	} else if data.AzxacmlUseSoap.ValueBool() {
		data.AzxacmlUseSoap = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZZOSNSSConfig`); value.Exists() && !data.AzzosnssConfig.IsNull() {
		data.AzzosnssConfig = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzzosnssConfig = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSAFDefaultAction`); value.Exists() && !data.AzsafDefaultAction.IsNull() {
		data.AzsafDefaultAction = tfutils.ParseStringFromGJSON(value)
	} else if data.AzsafDefaultAction.ValueString() != "r" {
		data.AzsafDefaultAction = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZLDAPAttributes`); value.Exists() && !data.AzldapAttributes.IsNull() {
		data.AzldapAttributes = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzldapAttributes = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSkewTime`); value.Exists() && !data.AzSkewTime.IsNull() {
		data.AzSkewTime = types.Int64Value(value.Int())
	} else if data.AzSkewTime.ValueInt64() != 0 {
		data.AzSkewTime = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AZOAuthEnforceScope`); value.Exists() && !data.AzoAuthEnforceScope.IsNull() {
		data.AzoAuthEnforceScope = tfutils.BoolFromString(value.String())
	} else if data.AzoAuthEnforceScope.ValueBool() {
		data.AzoAuthEnforceScope = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZOAuthExportHeaders`); value.Exists() && !data.AzoAuthExportHeaders.IsNull() {
		data.AzoAuthExportHeaders = tfutils.BoolFromString(value.String())
	} else if !data.AzoAuthExportHeaders.ValueBool() {
		data.AzoAuthExportHeaders = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZTAMPACReturn`); value.Exists() && !data.AztampacReturn.IsNull() {
		data.AztampacReturn = tfutils.BoolFromString(value.String())
	} else if data.AztampacReturn.ValueBool() {
		data.AztampacReturn = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZTAMPACUse`); value.Exists() && !data.AztampacUse.IsNull() {
		data.AztampacUse = tfutils.BoolFromString(value.String())
	} else if data.AztampacUse.ValueBool() {
		data.AztampacUse = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZLDAPReadTimeout`); value.Exists() && !data.AzldapReadTimeout.IsNull() {
		data.AzldapReadTimeout = types.Int64Value(value.Int())
	} else if data.AzldapReadTimeout.ValueInt64() != 60 {
		data.AzldapReadTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AZSSLClientConfigType`); value.Exists() && !data.AzsslClientConfigType.IsNull() {
		data.AzsslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else if data.AzsslClientConfigType.ValueString() != "client" {
		data.AzsslClientConfigType = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSSLClientProfile`); value.Exists() && !data.AzsslClientProfile.IsNull() {
		data.AzsslClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzsslClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZLDAPBindPasswordAlias`); value.Exists() && !data.AzldapBindPasswordAlias.IsNull() {
		data.AzldapBindPasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzldapBindPasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSMRequestType`); value.Exists() && !data.AzsmRequestType.IsNull() {
		data.AzsmRequestType = tfutils.ParseStringFromGJSON(value)
	} else if data.AzsmRequestType.ValueString() != "webagent" {
		data.AzsmRequestType = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZSMCookieFlow`); value.Exists() {
		data.AzsmCookieFlow.UpdateFromBody(ctx, "", value)
	} else {
		data.AzsmCookieFlow = nil
	}
	if value := res.Get(pathRoot + `AZSMHeaderFlow`); value.Exists() {
		data.AzsmHeaderFlow.UpdateFromBody(ctx, "", value)
	} else {
		data.AzsmHeaderFlow = nil
	}
	if value := res.Get(pathRoot + `AZSMCookieAttributes`); value.Exists() && !data.AzsmCookieAttributes.IsNull() {
		data.AzsmCookieAttributes = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzsmCookieAttributes = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZCacheControl`); value.Exists() && !data.AzCacheControl.IsNull() {
		data.AzCacheControl = tfutils.ParseStringFromGJSON(value)
	} else if data.AzCacheControl.ValueString() != "default" {
		data.AzCacheControl = types.StringNull()
	}
}
