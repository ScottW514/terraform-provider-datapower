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

type DmAAAPAuthorize struct {
	AzMethod                      types.String `tfsdk:"az_method"`
	AzCustomUrl                   types.String `tfsdk:"az_custom_url"`
	AzMapUrl                      types.String `tfsdk:"az_map_url"`
	AzHost                        types.String `tfsdk:"az_host"`
	AzPort                        types.Int64  `tfsdk:"az_port"`
	AzldapGroup                   types.String `tfsdk:"azldap_group"`
	AzValcred                     types.String `tfsdk:"az_valcred"`
	Azsamlurl                     types.String `tfsdk:"azsamlurl"`
	AzsamlType                    types.String `tfsdk:"azsaml_type"`
	AzsamlxPath                   types.String `tfsdk:"azsamlx_path"`
	AzsamlNameQualifier           types.String `tfsdk:"azsaml_name_qualifier"`
	AzCacheAllow                  types.String `tfsdk:"az_cache_allow"`
	AzCacheTtl                    types.Int64  `tfsdk:"az_cache_ttl"`
	AzNetegrityBaseUri            types.String `tfsdk:"az_netegrity_base_uri"`
	AzNetegrityOpNameExtension    types.String `tfsdk:"az_netegrity_op_name_extension"`
	AzClearTrustServerUrl         types.String `tfsdk:"az_clear_trust_server_url"`
	AzsamlVersion                 types.String `tfsdk:"azsaml_version"`
	AzldapLoadBalanceGroup        types.String `tfsdk:"azldap_load_balance_group"`
	AzldapBindDn                  types.String `tfsdk:"azldap_bind_dn"`
	AzldapBindPassword            types.String `tfsdk:"azldap_bind_password"`
	AzldapGroupAttribute          types.String `tfsdk:"azldap_group_attribute"`
	AzldapSearchScope             types.String `tfsdk:"azldap_search_scope"`
	AzldapSearchFilter            types.String `tfsdk:"azldap_search_filter"`
	AzxacmlVersion                types.String `tfsdk:"azxacml_version"`
	AzxacmlpepType                types.String `tfsdk:"azxacmlpep_type"`
	AzxacmlUseOnBoxPdp            types.Bool   `tfsdk:"azxacml_use_on_box_pdp"`
	Azxacmlpdp                    types.String `tfsdk:"azxacmlpdp"`
	AzxacmlExternalPdpUrl         types.String `tfsdk:"azxacml_external_pdp_url"`
	AzxacmlBindingMethod          types.String `tfsdk:"azxacml_binding_method"`
	AzxacmlBindingObject          types.String `tfsdk:"azxacml_binding_object"`
	AzxacmlBindingXsl             types.String `tfsdk:"azxacml_binding_xsl"`
	AzxacmlCustomObligation       types.String `tfsdk:"azxacml_custom_obligation"`
	AzxacmlUseSaml2               types.Bool   `tfsdk:"azxacml_use_saml2"`
	AztamServer                   types.String `tfsdk:"aztam_server"`
	AztamDefaultAction            types.String `tfsdk:"aztam_default_action"`
	AztamActionResourceMap        types.String `tfsdk:"aztam_action_resource_map"`
	AzxacmlUseSoap                types.Bool   `tfsdk:"azxacml_use_soap"`
	AzzosnssConfig                types.String `tfsdk:"azzosnss_config"`
	AzsafDefaultAction            types.String `tfsdk:"azsaf_default_action"`
	AzldapAttributes              types.String `tfsdk:"azldap_attributes"`
	AzSkewTime                    types.Int64  `tfsdk:"az_skew_time"`
	AzoAuthValidationEndpointType types.String `tfsdk:"azo_auth_validation_endpoint_type"`
	AztfimEndpoint                types.String `tfsdk:"aztfim_endpoint"`
	AzoAuthEnforceScope           types.Bool   `tfsdk:"azo_auth_enforce_scope"`
	AzoAuthExportHeaders          types.Bool   `tfsdk:"azo_auth_export_headers"`
	AztampacReturn                types.Bool   `tfsdk:"aztampac_return"`
	AztampacUse                   types.Bool   `tfsdk:"aztampac_use"`
	AzldapReadTimeout             types.Int64  `tfsdk:"azldap_read_timeout"`
	AzsslClientConfigType         types.String `tfsdk:"azssl_client_config_type"`
	AzsslClientProfile            types.String `tfsdk:"azssl_client_profile"`
	AzldapBindPasswordAlias       types.String `tfsdk:"azldap_bind_password_alias"`
	AzsmRequestType               types.String `tfsdk:"azsm_request_type"`
	AzsmCookieFlow                *DmSMFlow    `tfsdk:"azsm_cookie_flow"`
	AzsmHeaderFlow                *DmSMFlow    `tfsdk:"azsm_header_flow"`
	AzsmCookieAttributes          types.String `tfsdk:"azsm_cookie_attributes"`
	AzCacheControl                types.String `tfsdk:"az_cache_control"`
}

var DmAAAPAuthorizeObjectType = map[string]attr.Type{
	"az_method":                         types.StringType,
	"az_custom_url":                     types.StringType,
	"az_map_url":                        types.StringType,
	"az_host":                           types.StringType,
	"az_port":                           types.Int64Type,
	"azldap_group":                      types.StringType,
	"az_valcred":                        types.StringType,
	"azsamlurl":                         types.StringType,
	"azsaml_type":                       types.StringType,
	"azsamlx_path":                      types.StringType,
	"azsaml_name_qualifier":             types.StringType,
	"az_cache_allow":                    types.StringType,
	"az_cache_ttl":                      types.Int64Type,
	"az_netegrity_base_uri":             types.StringType,
	"az_netegrity_op_name_extension":    types.StringType,
	"az_clear_trust_server_url":         types.StringType,
	"azsaml_version":                    types.StringType,
	"azldap_load_balance_group":         types.StringType,
	"azldap_bind_dn":                    types.StringType,
	"azldap_bind_password":              types.StringType,
	"azldap_group_attribute":            types.StringType,
	"azldap_search_scope":               types.StringType,
	"azldap_search_filter":              types.StringType,
	"azxacml_version":                   types.StringType,
	"azxacmlpep_type":                   types.StringType,
	"azxacml_use_on_box_pdp":            types.BoolType,
	"azxacmlpdp":                        types.StringType,
	"azxacml_external_pdp_url":          types.StringType,
	"azxacml_binding_method":            types.StringType,
	"azxacml_binding_object":            types.StringType,
	"azxacml_binding_xsl":               types.StringType,
	"azxacml_custom_obligation":         types.StringType,
	"azxacml_use_saml2":                 types.BoolType,
	"aztam_server":                      types.StringType,
	"aztam_default_action":              types.StringType,
	"aztam_action_resource_map":         types.StringType,
	"azxacml_use_soap":                  types.BoolType,
	"azzosnss_config":                   types.StringType,
	"azsaf_default_action":              types.StringType,
	"azldap_attributes":                 types.StringType,
	"az_skew_time":                      types.Int64Type,
	"azo_auth_validation_endpoint_type": types.StringType,
	"aztfim_endpoint":                   types.StringType,
	"azo_auth_enforce_scope":            types.BoolType,
	"azo_auth_export_headers":           types.BoolType,
	"aztampac_return":                   types.BoolType,
	"aztampac_use":                      types.BoolType,
	"azldap_read_timeout":               types.Int64Type,
	"azssl_client_config_type":          types.StringType,
	"azssl_client_profile":              types.StringType,
	"azldap_bind_password_alias":        types.StringType,
	"azsm_request_type":                 types.StringType,
	"azsm_cookie_flow":                  types.ObjectType{AttrTypes: DmSMFlowObjectType},
	"azsm_header_flow":                  types.ObjectType{AttrTypes: DmSMFlowObjectType},
	"azsm_cookie_attributes":            types.StringType,
	"az_cache_control":                  types.StringType,
}
var DmAAAPAuthorizeObjectDefault = map[string]attr.Value{
	"az_method":                         types.StringValue("anyauthenticated"),
	"az_custom_url":                     types.StringNull(),
	"az_map_url":                        types.StringNull(),
	"az_host":                           types.StringNull(),
	"az_port":                           types.Int64Value(0),
	"azldap_group":                      types.StringNull(),
	"az_valcred":                        types.StringNull(),
	"azsamlurl":                         types.StringNull(),
	"azsaml_type":                       types.StringValue("any"),
	"azsamlx_path":                      types.StringNull(),
	"azsaml_name_qualifier":             types.StringNull(),
	"az_cache_allow":                    types.StringValue("absolute"),
	"az_cache_ttl":                      types.Int64Value(3),
	"az_netegrity_base_uri":             types.StringNull(),
	"az_netegrity_op_name_extension":    types.StringNull(),
	"az_clear_trust_server_url":         types.StringNull(),
	"azsaml_version":                    types.StringValue("1.1"),
	"azldap_load_balance_group":         types.StringNull(),
	"azldap_bind_dn":                    types.StringNull(),
	"azldap_bind_password":              types.StringNull(),
	"azldap_group_attribute":            types.StringValue("member"),
	"azldap_search_scope":               types.StringValue("subtree"),
	"azldap_search_filter":              types.StringValue("(objectClass=*)"),
	"azxacml_version":                   types.StringValue("2"),
	"azxacmlpep_type":                   types.StringValue("deny-biased"),
	"azxacml_use_on_box_pdp":            types.BoolValue(true),
	"azxacmlpdp":                        types.StringNull(),
	"azxacml_external_pdp_url":          types.StringNull(),
	"azxacml_binding_method":            types.StringValue("custom"),
	"azxacml_binding_object":            types.StringNull(),
	"azxacml_binding_xsl":               types.StringNull(),
	"azxacml_custom_obligation":         types.StringNull(),
	"azxacml_use_saml2":                 types.BoolValue(false),
	"aztam_server":                      types.StringNull(),
	"aztam_default_action":              types.StringValue("T"),
	"aztam_action_resource_map":         types.StringNull(),
	"azxacml_use_soap":                  types.BoolValue(false),
	"azzosnss_config":                   types.StringNull(),
	"azsaf_default_action":              types.StringValue("r"),
	"azldap_attributes":                 types.StringNull(),
	"az_skew_time":                      types.Int64Value(0),
	"azo_auth_validation_endpoint_type": types.StringValue("tfim"),
	"aztfim_endpoint":                   types.StringNull(),
	"azo_auth_enforce_scope":            types.BoolValue(false),
	"azo_auth_export_headers":           types.BoolValue(true),
	"aztampac_return":                   types.BoolValue(false),
	"aztampac_use":                      types.BoolValue(false),
	"azldap_read_timeout":               types.Int64Value(60),
	"azssl_client_config_type":          types.StringValue("client"),
	"azssl_client_profile":              types.StringNull(),
	"azldap_bind_password_alias":        types.StringNull(),
	"azsm_request_type":                 types.StringValue("webagent"),
	"azsm_cookie_flow":                  types.ObjectValueMust(DmSMFlowObjectType, DmSMFlowObjectDefault),
	"azsm_header_flow":                  types.ObjectValueMust(DmSMFlowObjectType, DmSMFlowObjectDefault),
	"azsm_cookie_attributes":            types.StringNull(),
	"az_cache_control":                  types.StringValue("default"),
}
var DmAAAPAuthorizeDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"az_method": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Method", "method", "").AddStringEnum("anyauthenticated", "passthrough", "tivoli", "netegrity", "oblix", "cleartrust", "custom", "ldap", "saml-authz", "saml-attr", "use-authen-attr", "xacml", "xmlfile", "zosnss", "oauth").AddDefaultValue("anyauthenticated").String,
			Computed:            true,
		},
		"az_custom_url": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Custom URL", "custom-url", "").String,
			Computed:            true,
		},
		"az_map_url": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("AAA information file URL", "xmlfile-url", "").String,
			Computed:            true,
		},
		"az_host": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Host", "remote-host", "").String,
			Computed:            true,
		},
		"az_port": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Port", "remote-port", "").AddDefaultValue("0").String,
			Computed:            true,
		},
		"azldap_group": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP Group DN", "ldap-group-dn", "").String,
			Computed:            true,
		},
		"az_valcred": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "cryptovalcred").String,
			Computed:            true,
		},
		"azsamlurl": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML server URL", "saml-server-url", "").String,
			Computed:            true,
		},
		"azsaml_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML match", "saml-type", "").AddStringEnum("xpath", "any", "all", "any-value", "all-values").AddDefaultValue("any").String,
			Computed:            true,
		},
		"azsamlx_path": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML XPath", "saml-xpath", "").String,
			Computed:            true,
		},
		"azsaml_name_qualifier": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML NameQualifier", "saml-name-qualifier", "").String,
			Computed:            true,
		},
		"az_cache_allow": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Cache authorization results", "cache-type", "").AddStringEnum("absolute", "disabled", "maximum", "minimum").AddDefaultValue("absolute").String,
			Computed:            true,
		},
		"az_cache_ttl": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Cache lifetime", "cache-ttl", "").AddIntegerRange(1, 86400).AddDefaultValue("3").String,
			Computed:            true,
		},
		"az_netegrity_base_uri": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("CA Single Sign-On Base URI", "netegrity-base-uri", "").String,
			Computed:            true,
		},
		"az_netegrity_op_name_extension": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Operation name extension", "netegrity-opname-ext", "").String,
			Computed:            true,
		},
		"az_clear_trust_server_url": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("ClearTrust server URL", "cleartrust-server-url", "").String,
			Computed:            true,
		},
		"azsaml_version": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML version", "saml-version", "").AddStringEnum("2.0", "1.1", "1.0").AddDefaultValue("1.1").String,
			Computed:            true,
		},
		"azldap_load_balance_group": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP load balancer group", "ldap-lbgroup", "loadbalancergroup").String,
			Computed:            true,
		},
		"azldap_bind_dn": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP bind DN", "ldap-bind-dn", "").String,
			Computed:            true,
		},
		"azldap_bind_password": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP bind password (deprecated)", "ldap-bind-password", "").String,
			Computed:            true,
		},
		"azldap_group_attribute": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP group attribute", "ldap-group-attr", "").AddDefaultValue("member").String,
			Computed:            true,
		},
		"azldap_search_scope": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP search scope", "ldap-search-scope", "").AddStringEnum("subtree", "one-level", "base").AddDefaultValue("subtree").String,
			Computed:            true,
		},
		"azldap_search_filter": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP search filter", "ldap-search-filter", "").AddDefaultValue("(objectClass=*)").String,
			Computed:            true,
		},
		"azxacml_version": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("XACML version", "xacml-version", "").AddStringEnum("2", "1").AddDefaultValue("2").String,
			Computed:            true,
		},
		"azxacmlpep_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("PEP type", "xacml-pep-type", "").AddStringEnum("base", "deny-biased", "permit-biased").AddDefaultValue("deny-biased").String,
			Computed:            true,
		},
		"azxacml_use_on_box_pdp": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Use on-box PDP", "xacml-use-builtin", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"azxacmlpdp": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Policy decision point", "xacml-pdp", "xacmlpdp").String,
			Computed:            true,
		},
		"azxacml_external_pdp_url": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL for external policy decision point", "xacml-url", "").String,
			Computed:            true,
		},
		"azxacml_binding_method": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("XACML binding method", "xacml-binding-method", "").AddStringEnum("dp-pdp", "custom").AddDefaultValue("custom").String,
			Computed:            true,
		},
		"azxacml_binding_object": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("XACML binding", "xacml-binding-object", "").String,
			Computed:            true,
		},
		"azxacml_binding_xsl": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Custom processing to bind AAA and XACML", "xacml-binding-custom-url", "").String,
			Computed:            true,
		},
		"azxacml_custom_obligation": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Custom obligation fulfillment processing", "xacml-obligation-custom-url", "").String,
			Computed:            true,
		},
		"azxacml_use_saml2": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("PDP requires SAML 2.0", "xacml-use-saml2", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"aztam_server": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("IBM Security Access Manager client", "tam", "tam").String,
			Computed:            true,
		},
		"aztam_default_action": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Default action", "tam-action-default", "").AddStringEnum("T", "c", "g", "m", "d", "b", "s", "v", "a", "BypassPOP", "tt", "r", "x", "l", "N", "W", "Add", "BypassAuthzRule", "_WebService_i", "_PDMQ_D", "_PDMQ_E").AddDefaultValue("T").String,
			Computed:            true,
		},
		"aztam_action_resource_map": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Resource-action map", "tam-action-map", "").String,
			Computed:            true,
		},
		"azxacml_use_soap": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SOAP enveloping", "xacml-use-soap", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"azzosnss_config": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("z/OS NSS client configuration", "zos-nss-az", "zosnssclient").String,
			Computed:            true,
		},
		"azsaf_default_action": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Default action", "zos-nss-default-action", "").AddStringEnum("r", "u", "a", "c").AddDefaultValue("r").String,
			Computed:            true,
		},
		"azldap_attributes": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("User auxiliary LDAP attributes", "az-ldap-attributes", "").String,
			Computed:            true,
		},
		"az_skew_time": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Skew time", "az-skew-time", "").AddDefaultValue("0").String,
			Computed:            true,
		},
		"azo_auth_validation_endpoint_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("OAuth endpoint type (deprecated)", "az-oauth-endpoint-type", "").AddStringEnum("tfim").AddDefaultValue("tfim").String,
			Computed:            true,
		},
		"aztfim_endpoint": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Federated Identity Manager endpoint (deprecated)", "az-tfim-endpoint", "tfimendpoint").String,
			Computed:            true,
		},
		"azo_auth_enforce_scope": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enforce scope", "az-oauth-enforce-scope", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"azo_auth_export_headers": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Export response attributes", "az-oauth-export-headers", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"aztampac_return": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Return Privilege Attribute Certificate", "tam-pac-return", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"aztampac_use": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Use Privilege Attribute Certificate", "use-tam-pac", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"azldap_read_timeout": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP Read Timeout", "ldap-readtimeout", "").AddIntegerRange(0, 86400).AddDefaultValue("60").String,
			Computed:            true,
		},
		"azssl_client_config_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("TLS client type", "ssl-client-type", "").AddStringEnum("proxy", "client").AddDefaultValue("client").String,
			Computed:            true,
		},
		"azssl_client_profile": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "sslclientprofile").String,
			Computed:            true,
		},
		"azldap_bind_password_alias": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP bind password alias", "ldap-bind-password-alias", "passwordalias").String,
			Computed:            true,
		},
		"azsm_request_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Request type", "sm-request-type", "").AddStringEnum("webagent", "webservice").AddDefaultValue("webagent").String,
			Computed:            true,
		},
		"azsm_cookie_flow": GetDmSMFlowDataSourceSchema("Session cookie flow", "sm-cookie-flow", ""),
		"azsm_header_flow": GetDmSMFlowDataSourceSchema("CA Single Sign-On header flow", "sm-header-flow", ""),
		"azsm_cookie_attributes": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Cookie attribute policy", "cookie-attributes", "cookieattributepolicy").String,
			Computed:            true,
		},
		"az_cache_control": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Authorization caching", "cache-control", "").AddStringEnum("default", "disable-all", "disable-ldap-failures").AddDefaultValue("default").String,
			Computed:            true,
		},
	},
}
var DmAAAPAuthorizeResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmAAAPAuthorizeObjectType,
			DmAAAPAuthorizeObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"az_method": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Method", "method", "").AddStringEnum("anyauthenticated", "passthrough", "tivoli", "netegrity", "oblix", "cleartrust", "custom", "ldap", "saml-authz", "saml-attr", "use-authen-attr", "xacml", "xmlfile", "zosnss", "oauth").AddDefaultValue("anyauthenticated").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("anyauthenticated", "passthrough", "tivoli", "netegrity", "oblix", "cleartrust", "custom", "ldap", "saml-authz", "saml-attr", "use-authen-attr", "xacml", "xmlfile", "zosnss", "oauth"),
			},
			Default: stringdefault.StaticString("anyauthenticated"),
		},
		"az_custom_url": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Custom URL", "custom-url", "").String,
			Optional:            true,
		},
		"az_map_url": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("AAA information file URL", "xmlfile-url", "").String,
			Optional:            true,
		},
		"az_host": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Host", "remote-host", "").String,
			Optional:            true,
		},
		"az_port": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Port", "remote-port", "").AddDefaultValue("0").String,
			Computed:            true,
			Optional:            true,
			Default:             int64default.StaticInt64(0),
		},
		"azldap_group": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP Group DN", "ldap-group-dn", "").String,
			Optional:            true,
		},
		"az_valcred": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "cryptovalcred").String,
			Optional:            true,
		},
		"azsamlurl": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML server URL", "saml-server-url", "").String,
			Optional:            true,
		},
		"azsaml_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML match", "saml-type", "").AddStringEnum("xpath", "any", "all", "any-value", "all-values").AddDefaultValue("any").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("xpath", "any", "all", "any-value", "all-values"),
			},
			Default: stringdefault.StaticString("any"),
		},
		"azsamlx_path": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML XPath", "saml-xpath", "").String,
			Optional:            true,
		},
		"azsaml_name_qualifier": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML NameQualifier", "saml-name-qualifier", "").String,
			Optional:            true,
		},
		"az_cache_allow": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Cache authorization results", "cache-type", "").AddStringEnum("absolute", "disabled", "maximum", "minimum").AddDefaultValue("absolute").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("absolute", "disabled", "maximum", "minimum"),
			},
			Default: stringdefault.StaticString("absolute"),
		},
		"az_cache_ttl": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Cache lifetime", "cache-ttl", "").AddIntegerRange(1, 86400).AddDefaultValue("3").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(1, 86400),
			},
			Default: int64default.StaticInt64(3),
		},
		"az_netegrity_base_uri": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("CA Single Sign-On Base URI", "netegrity-base-uri", "").String,
			Optional:            true,
		},
		"az_netegrity_op_name_extension": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Operation name extension", "netegrity-opname-ext", "").String,
			Optional:            true,
		},
		"az_clear_trust_server_url": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("ClearTrust server URL", "cleartrust-server-url", "").String,
			Optional:            true,
		},
		"azsaml_version": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SAML version", "saml-version", "").AddStringEnum("2.0", "1.1", "1.0").AddDefaultValue("1.1").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("2.0", "1.1", "1.0"),
			},
			Default: stringdefault.StaticString("1.1"),
		},
		"azldap_load_balance_group": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP load balancer group", "ldap-lbgroup", "loadbalancergroup").String,
			Optional:            true,
		},
		"azldap_bind_dn": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP bind DN", "ldap-bind-dn", "").String,
			Optional:            true,
		},
		"azldap_bind_password": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP bind password (deprecated)", "ldap-bind-password", "").String,
			Optional:            true,
		},
		"azldap_group_attribute": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP group attribute", "ldap-group-attr", "").AddDefaultValue("member").String,
			Computed:            true,
			Optional:            true,
			Default:             stringdefault.StaticString("member"),
		},
		"azldap_search_scope": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP search scope", "ldap-search-scope", "").AddStringEnum("subtree", "one-level", "base").AddDefaultValue("subtree").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("subtree", "one-level", "base"),
			},
			Default: stringdefault.StaticString("subtree"),
		},
		"azldap_search_filter": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP search filter", "ldap-search-filter", "").AddDefaultValue("(objectClass=*)").String,
			Computed:            true,
			Optional:            true,
			Default:             stringdefault.StaticString("(objectClass=*)"),
		},
		"azxacml_version": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("XACML version", "xacml-version", "").AddStringEnum("2", "1").AddDefaultValue("2").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("2", "1"),
			},
			Default: stringdefault.StaticString("2"),
		},
		"azxacmlpep_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("PEP type", "xacml-pep-type", "").AddStringEnum("base", "deny-biased", "permit-biased").AddDefaultValue("deny-biased").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("base", "deny-biased", "permit-biased"),
			},
			Default: stringdefault.StaticString("deny-biased"),
		},
		"azxacml_use_on_box_pdp": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Use on-box PDP", "xacml-use-builtin", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"azxacmlpdp": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Policy decision point", "xacml-pdp", "xacmlpdp").String,
			Optional:            true,
		},
		"azxacml_external_pdp_url": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL for external policy decision point", "xacml-url", "").String,
			Optional:            true,
		},
		"azxacml_binding_method": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("XACML binding method", "xacml-binding-method", "").AddStringEnum("dp-pdp", "custom").AddDefaultValue("custom").String,
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
			MarkdownDescription: tfutils.NewAttributeDescription("Custom processing to bind AAA and XACML", "xacml-binding-custom-url", "").String,
			Optional:            true,
		},
		"azxacml_custom_obligation": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Custom obligation fulfillment processing", "xacml-obligation-custom-url", "").String,
			Optional:            true,
		},
		"azxacml_use_saml2": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("PDP requires SAML 2.0", "xacml-use-saml2", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"aztam_server": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("IBM Security Access Manager client", "tam", "tam").String,
			Optional:            true,
		},
		"aztam_default_action": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Default action", "tam-action-default", "").AddStringEnum("T", "c", "g", "m", "d", "b", "s", "v", "a", "BypassPOP", "tt", "r", "x", "l", "N", "W", "Add", "BypassAuthzRule", "_WebService_i", "_PDMQ_D", "_PDMQ_E").AddDefaultValue("T").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("T", "c", "g", "m", "d", "b", "s", "v", "a", "BypassPOP", "tt", "r", "x", "l", "N", "W", "Add", "BypassAuthzRule", "_WebService_i", "_PDMQ_D", "_PDMQ_E"),
			},
			Default: stringdefault.StaticString("T"),
		},
		"aztam_action_resource_map": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Resource-action map", "tam-action-map", "").String,
			Optional:            true,
		},
		"azxacml_use_soap": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SOAP enveloping", "xacml-use-soap", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"azzosnss_config": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("z/OS NSS client configuration", "zos-nss-az", "zosnssclient").String,
			Optional:            true,
		},
		"azsaf_default_action": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Default action", "zos-nss-default-action", "").AddStringEnum("r", "u", "a", "c").AddDefaultValue("r").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("r", "u", "a", "c"),
			},
			Default: stringdefault.StaticString("r"),
		},
		"azldap_attributes": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("User auxiliary LDAP attributes", "az-ldap-attributes", "").String,
			Optional:            true,
		},
		"az_skew_time": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Skew time", "az-skew-time", "").AddDefaultValue("0").String,
			Computed:            true,
			Optional:            true,
			Default:             int64default.StaticInt64(0),
		},
		"azo_auth_validation_endpoint_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("OAuth endpoint type (deprecated)", "az-oauth-endpoint-type", "").AddStringEnum("tfim").AddDefaultValue("tfim").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("tfim"),
			},
			Default: stringdefault.StaticString("tfim"),
		},
		"aztfim_endpoint": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Federated Identity Manager endpoint (deprecated)", "az-tfim-endpoint", "tfimendpoint").String,
			Optional:            true,
		},
		"azo_auth_enforce_scope": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enforce scope", "az-oauth-enforce-scope", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"azo_auth_export_headers": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Export response attributes", "az-oauth-export-headers", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"aztampac_return": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Return Privilege Attribute Certificate", "tam-pac-return", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"aztampac_use": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Use Privilege Attribute Certificate", "use-tam-pac", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"azldap_read_timeout": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP Read Timeout", "ldap-readtimeout", "").AddIntegerRange(0, 86400).AddDefaultValue("60").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(0, 86400),
			},
			Default: int64default.StaticInt64(60),
		},
		"azssl_client_config_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("TLS client type", "ssl-client-type", "").AddStringEnum("proxy", "client").AddDefaultValue("client").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("proxy", "client"),
			},
			Default: stringdefault.StaticString("client"),
		},
		"azssl_client_profile": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "sslclientprofile").String,
			Optional:            true,
		},
		"azldap_bind_password_alias": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP bind password alias", "ldap-bind-password-alias", "passwordalias").String,
			Optional:            true,
		},
		"azsm_request_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Request type", "sm-request-type", "").AddStringEnum("webagent", "webservice").AddDefaultValue("webagent").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("webagent", "webservice"),
			},
			Default: stringdefault.StaticString("webagent"),
		},
		"azsm_cookie_flow": GetDmSMFlowResourceSchema("Session cookie flow", "sm-cookie-flow", "", false),
		"azsm_header_flow": GetDmSMFlowResourceSchema("CA Single Sign-On header flow", "sm-header-flow", "", false),
		"azsm_cookie_attributes": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Cookie attribute policy", "cookie-attributes", "cookieattributepolicy").String,
			Optional:            true,
		},
		"az_cache_control": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Authorization caching", "cache-control", "").AddStringEnum("default", "disable-all", "disable-ldap-failures").AddDefaultValue("default").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("default", "disable-all", "disable-ldap-failures"),
			},
			Default: stringdefault.StaticString("default"),
		},
	},
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
	if !data.AzldapBindPassword.IsNull() {
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
	if !data.AzoAuthValidationEndpointType.IsNull() {
		return false
	}
	if !data.AztfimEndpoint.IsNull() {
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
func GetDmAAAPAuthorizeDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmAAAPAuthorizeDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmAAAPAuthorizeDataSourceSchema
}

func GetDmAAAPAuthorizeResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmAAAPAuthorizeResourceSchema.Required = true
	} else {
		DmAAAPAuthorizeResourceSchema.Optional = true
		DmAAAPAuthorizeResourceSchema.Computed = true
	}
	DmAAAPAuthorizeResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmAAAPAuthorizeResourceSchema
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
	if !data.AzldapBindPassword.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZLDAPBindPassword`, data.AzldapBindPassword.ValueString())
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
		body, _ = sjson.Set(body, pathRoot+`AZXACMLUseOnBoxPDP`, tfutils.StringFromBool(data.AzxacmlUseOnBoxPdp))
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
		body, _ = sjson.Set(body, pathRoot+`AZXACMLUseSAML2`, tfutils.StringFromBool(data.AzxacmlUseSaml2))
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
		body, _ = sjson.Set(body, pathRoot+`AZXACMLUseSOAP`, tfutils.StringFromBool(data.AzxacmlUseSoap))
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
	if !data.AzoAuthValidationEndpointType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZOAuthValidationEndpointType`, data.AzoAuthValidationEndpointType.ValueString())
	}
	if !data.AztfimEndpoint.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZTFIMEndpoint`, data.AztfimEndpoint.ValueString())
	}
	if !data.AzoAuthEnforceScope.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZOAuthEnforceScope`, tfutils.StringFromBool(data.AzoAuthEnforceScope))
	}
	if !data.AzoAuthExportHeaders.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZOAuthExportHeaders`, tfutils.StringFromBool(data.AzoAuthExportHeaders))
	}
	if !data.AztampacReturn.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZTAMPACReturn`, tfutils.StringFromBool(data.AztampacReturn))
	}
	if !data.AztampacUse.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZTAMPACUse`, tfutils.StringFromBool(data.AztampacUse))
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
	if value := res.Get(pathRoot + `AZLDAPBindPassword`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzldapBindPassword = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzldapBindPassword = types.StringNull()
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
	if value := res.Get(pathRoot + `AZOAuthValidationEndpointType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzoAuthValidationEndpointType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzoAuthValidationEndpointType = types.StringValue("tfim")
	}
	if value := res.Get(pathRoot + `AZTFIMEndpoint`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AztfimEndpoint = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AztfimEndpoint = types.StringNull()
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
	if value := res.Get(pathRoot + `AZLDAPBindPassword`); value.Exists() && !data.AzldapBindPassword.IsNull() {
		data.AzldapBindPassword = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzldapBindPassword = types.StringNull()
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
	if value := res.Get(pathRoot + `AZOAuthValidationEndpointType`); value.Exists() && !data.AzoAuthValidationEndpointType.IsNull() {
		data.AzoAuthValidationEndpointType = tfutils.ParseStringFromGJSON(value)
	} else if data.AzoAuthValidationEndpointType.ValueString() != "tfim" {
		data.AzoAuthValidationEndpointType = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZTFIMEndpoint`); value.Exists() && !data.AztfimEndpoint.IsNull() {
		data.AztfimEndpoint = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AztfimEndpoint = types.StringNull()
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
