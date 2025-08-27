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
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type OAuthProviderSettings struct {
	Id                                    types.String                `tfsdk:"id"`
	AppDomain                             types.String                `tfsdk:"app_domain"`
	UserSummary                           types.String                `tfsdk:"user_summary"`
	EnableDebugMode                       types.Bool                  `tfsdk:"enable_debug_mode"`
	ProviderType                          types.String                `tfsdk:"provider_type"`
	ScopesAllowed                         types.String                `tfsdk:"scopes_allowed"`
	DefaultScopes                         types.String                `tfsdk:"default_scopes"`
	SupportedGrantTypes                   *DmOAuthProviderGrantType   `tfsdk:"supported_grant_types"`
	SupportedClientTypes                  *DmAllowedClientType        `tfsdk:"supported_client_types"`
	ApiCProviderBasePath                  types.String                `tfsdk:"api_c_provider_base_path"`
	ApiCAuthorizeEndpoint                 types.String                `tfsdk:"api_c_authorize_endpoint"`
	ApiCTokenEndpoint                     types.String                `tfsdk:"api_c_token_endpoint"`
	ApiCEnableIntrospection               types.Bool                  `tfsdk:"api_c_enable_introspection"`
	ApiCIntrospectEndpoint                types.String                `tfsdk:"api_c_introspect_endpoint"`
	ApiCTokenSecret                       types.String                `tfsdk:"api_c_token_secret"`
	ApiCOneTimeUseAccesstoken             types.Bool                  `tfsdk:"api_c_one_time_use_accesstoken"`
	ApiCAccessTokenTtl                    types.Int64                 `tfsdk:"api_c_access_token_ttl"`
	ApiCAuthCodeTtl                       types.Int64                 `tfsdk:"api_c_auth_code_ttl"`
	ApiCEnableRefreshToken                types.Bool                  `tfsdk:"api_c_enable_refresh_token"`
	ApiCOneTimeUseRefreshtoken            types.Bool                  `tfsdk:"api_c_one_time_use_refreshtoken"`
	ApiCRefreshTokenLimit                 types.Int64                 `tfsdk:"api_c_refresh_token_limit"`
	ApiCRefreshTokenTtl                   types.Int64                 `tfsdk:"api_c_refresh_token_ttl"`
	ApiCMaximumConsentTtl                 types.Int64                 `tfsdk:"api_c_maximum_consent_ttl"`
	AdvScopeValidationEnabled             types.Bool                  `tfsdk:"adv_scope_validation_enabled"`
	AdvancedScopeUrlOverride              types.Bool                  `tfsdk:"advanced_scope_url_override"`
	AdvancedScopeUrl                      types.String                `tfsdk:"advanced_scope_url"`
	AdvScopeTlsProfile                    types.String                `tfsdk:"adv_scope_tls_profile"`
	AdvScopeUrlSecurityEnabled            types.Bool                  `tfsdk:"adv_scope_url_security_enabled"`
	AdvScopeUrlSecurity                   *DmSecurityType             `tfsdk:"adv_scope_url_security"`
	AdvScopeBasicAuthUserName             types.String                `tfsdk:"adv_scope_basic_auth_user_name"`
	AdvScopeBasicAuthPassword             types.String                `tfsdk:"adv_scope_basic_auth_password"`
	AdvScopeBasicAuthHeaderName           types.String                `tfsdk:"adv_scope_basic_auth_header_name"`
	AdvancedScopeCustomHeaders            types.String                `tfsdk:"advanced_scope_custom_headers"`
	AdvancedScopeCustomContexts           types.String                `tfsdk:"advanced_scope_custom_contexts"`
	ApiCEnableOidc                        types.Bool                  `tfsdk:"api_c_enable_oidc"`
	ApiCOidcHybridResponseTypes           *DmOIDCHybridResponseType   `tfsdk:"api_c_oidc_hybrid_response_types"`
	ApiCSupportPkce                       types.Bool                  `tfsdk:"api_c_support_pkce"`
	ApiCRequirePkce                       types.Bool                  `tfsdk:"api_c_require_pkce"`
	ApiCSupportPkcePlain                  types.Bool                  `tfsdk:"api_c_support_pkce_plain"`
	ApiCTokenTypeToGenerate               types.String                `tfsdk:"api_c_token_type_to_generate"`
	MetadataFrom                          *DmMetadataFromType         `tfsdk:"metadata_from"`
	MetadataUrl                           types.String                `tfsdk:"metadata_url"`
	MetadataSslProfile                    types.String                `tfsdk:"metadata_ssl_profile"`
	MetadataHeaderForAccessToken          types.String                `tfsdk:"metadata_header_for_access_token"`
	MetadataHeaderForPayload              types.String                `tfsdk:"metadata_header_for_payload"`
	ExternalRevocationUrl                 types.String                `tfsdk:"external_revocation_url"`
	ExternalRevocationSslProfile          types.String                `tfsdk:"external_revocation_ssl_profile"`
	ExternalRevocationUrlSecurity         *DmSecurityType             `tfsdk:"external_revocation_url_security"`
	ExternalRevocationBasicAuthUserName   types.String                `tfsdk:"external_revocation_basic_auth_user_name"`
	ExternalRevocationBasicAuthPassword   types.String                `tfsdk:"external_revocation_basic_auth_password"`
	ExternalRevocationBasicAuthHeaderName types.String                `tfsdk:"external_revocation_basic_auth_header_name"`
	ExternalRevocationCustomHeaderFormat  types.String                `tfsdk:"external_revocation_custom_header_format"`
	ExternalRevocationCacheType           types.String                `tfsdk:"external_revocation_cache_type"`
	ExternalRevocationCacheTimeToLive     types.Int64                 `tfsdk:"external_revocation_cache_time_to_live"`
	ExternalRevocationFailOnError         types.Bool                  `tfsdk:"external_revocation_fail_on_error"`
	EnableTokenManagement                 types.Bool                  `tfsdk:"enable_token_management"`
	TokenManagerType                      types.String                `tfsdk:"token_manager_type"`
	ApiSecurityTokenManager               types.String                `tfsdk:"api_security_token_manager"`
	EnableApplicationRevocation           types.Bool                  `tfsdk:"enable_application_revocation"`
	ApplicationRevocationEndpoint         types.String                `tfsdk:"application_revocation_endpoint"`
	EnableOwnerRevocation                 types.Bool                  `tfsdk:"enable_owner_revocation"`
	OwnerRevocationEndpoint               types.String                `tfsdk:"owner_revocation_endpoint"`
	TokenValidationReq                    types.String                `tfsdk:"token_validation_req"`
	ThirdPartyAzUrl                       types.String                `tfsdk:"third_party_az_url"`
	ThirdPartyTokenUrl                    types.String                `tfsdk:"third_party_token_url"`
	ThirdPartyIntrospectUrl               types.String                `tfsdk:"third_party_introspect_url"`
	ThirdPartyIntrospectCacheType         types.String                `tfsdk:"third_party_introspect_cache_type"`
	ThirdPartyIntrospectCacheTimeToLive   types.Int64                 `tfsdk:"third_party_introspect_cache_time_to_live"`
	ThirdPartyAuthorizationHeaderPassThru types.Bool                  `tfsdk:"third_party_authorization_header_pass_thru"`
	ThirdPartyIntrospectUrlSecurity       *DmSecurityType             `tfsdk:"third_party_introspect_url_security"`
	ThirdPartyIntrospectBasicAuthUserName types.String                `tfsdk:"third_party_introspect_basic_auth_user_name"`
	ThirdPartyIntrospectBasicAuthPassword types.String                `tfsdk:"third_party_introspect_basic_auth_password"`
	ThirdPartyBasicAuthHeaderName         types.String                `tfsdk:"third_party_basic_auth_header_name"`
	ThirdPartyCustomHeaderNameFormat      types.String                `tfsdk:"third_party_custom_header_name_format"`
	ThirdPartyIntrospectSslProfile        types.String                `tfsdk:"third_party_introspect_ssl_profile"`
	DependencyActions                     []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var OAuthProviderSettingsSupportedClientTypesCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "provider_type",
	AttrType:    "String",
	AttrDefault: "native",
	Value:       []string{"native"},
}
var OAuthProviderSettingsAPICProviderBasePathCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "provider_type",
	AttrType:    "String",
	AttrDefault: "native",
	Value:       []string{"native"},
}
var OAuthProviderSettingsAPICAuthorizeEndpointCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "provider_type",
	AttrType:    "String",
	AttrDefault: "native",
	Value:       []string{"native"},
}
var OAuthProviderSettingsAPICTokenEndpointCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "provider_type",
	AttrType:    "String",
	AttrDefault: "native",
	Value:       []string{"native"},
}
var OAuthProviderSettingsAPICIntrospectEndpointCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "api_c_enable_introspection",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
	},
}
var OAuthProviderSettingsAPICTokenSecretCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "provider_type",
	AttrType:    "String",
	AttrDefault: "native",
	Value:       []string{"native"},
}
var OAuthProviderSettingsAPICOneTimeUseAccesstokenCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "provider_type",
	AttrType:    "String",
	AttrDefault: "native",
	Value:       []string{"native"},
}
var OAuthProviderSettingsAPICAccessTokenTTLCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "provider_type",
	AttrType:    "String",
	AttrDefault: "native",
	Value:       []string{"native"},
}
var OAuthProviderSettingsAPICAuthCodeTTLCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "supported_grant_types",
			AttrType:    "DmOAuthProviderGrantType",
			AttrDefault: "",
			Value:       []string{"access_code"},
		},
	},
}
var OAuthProviderSettingsAPICRefreshTokenLimitCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "api_c_enable_refresh_token",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "api_c_one_time_use_refreshtoken",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
	},
}
var OAuthProviderSettingsAPICRefreshTokenTTLCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "api_c_enable_refresh_token",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
	},
}
var OAuthProviderSettingsAdvancedScopeURLCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "advanced_scope_url_override",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "adv_scope_validation_enabled",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
	},
}
var OAuthProviderSettingsMetadataURLCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "metadata_from",
			AttrType:    "DmMetadataFromType",
			AttrDefault: "",
			Value:       []string{"external_url"},
		},
	},
}
var OAuthProviderSettingsExternalRevocationURLCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "enable_token_management",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "token_manager_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"external"},
		},
	},
}
var OAuthProviderSettingsAPISecurityTokenManagerCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "enable_token_management",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "token_manager_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
	},
}
var OAuthProviderSettingsEnableApplicationRevocationCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "enable_token_management",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "token_manager_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
	},
}
var OAuthProviderSettingsApplicationRevocationEndpointCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "enable_application_revocation",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "enable_token_management",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "token_manager_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
	},
}
var OAuthProviderSettingsEnableOwnerRevocationCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "enable_token_management",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "token_manager_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
	},
}
var OAuthProviderSettingsOwnerRevocationEndpointCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "enable_owner_revocation",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "enable_token_management",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "token_manager_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
	},
}
var OAuthProviderSettingsTokenValidationReqCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "provider_type",
	AttrType:    "String",
	AttrDefault: "native",
	Value:       []string{"third_party"},
}
var OAuthProviderSettingsThirdPartyIntrospectURLCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "provider_type",
	AttrType:    "String",
	AttrDefault: "native",
	Value:       []string{"third_party"},
}
var OAuthProviderSettingsThirdPartyIntrospectCacheTypeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "provider_type",
	AttrType:    "String",
	AttrDefault: "native",
	Value:       []string{"third_party"},
}
var OAuthProviderSettingsThirdPartyIntrospectCacheTimeToLiveCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"third_party"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "third_party_introspect_cache_type",
			AttrType:    "String",
			AttrDefault: "NoCache",
			Value:       []string{"TimeToLive"},
		},
	},
}
var OAuthProviderSettingsDefaultScopesIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "provider_type",
	AttrType:    "String",
	AttrDefault: "native",
	Value:       []string{"third_party"},
}
var OAuthProviderSettingsSupportedClientTypesIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var OAuthProviderSettingsAPICProviderBasePathIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var OAuthProviderSettingsAPICAuthorizeEndpointIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var OAuthProviderSettingsAPICTokenEndpointIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var OAuthProviderSettingsAPICEnableIntrospectionIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "provider_type",
	AttrType:    "String",
	AttrDefault: "native",
	Value:       []string{"native"},
}
var OAuthProviderSettingsAPICIntrospectEndpointIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var OAuthProviderSettingsAPICTokenSecretIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var OAuthProviderSettingsAPICOneTimeUseAccesstokenIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var OAuthProviderSettingsAPICAccessTokenTTLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var OAuthProviderSettingsAPICAuthCodeTTLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var OAuthProviderSettingsAPICEnableRefreshTokenIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "provider_type",
	AttrType:    "String",
	AttrDefault: "native",
	Value:       []string{"native"},
}
var OAuthProviderSettingsAPICOneTimeUseRefreshtokenIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "api_c_enable_refresh_token",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
	},
}
var OAuthProviderSettingsAPICRefreshTokenLimitIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var OAuthProviderSettingsAPICRefreshTokenTTLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var OAuthProviderSettingsAPICMaximumConsentTTLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "api_c_enable_refresh_token",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
	},
}
var OAuthProviderSettingsAdvancedScopeURLOverrideIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "adv_scope_validation_enabled",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var OAuthProviderSettingsAdvancedScopeURLIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "adv_scope_validation_enabled",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var OAuthProviderSettingsAdvScopeTLSProfileIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "adv_scope_validation_enabled",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var OAuthProviderSettingsAdvScopeURLSecurityEnabledIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "adv_scope_validation_enabled",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var OAuthProviderSettingsAdvScopeURLSecurityIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "adv_scope_url_security_enabled",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "adv_scope_validation_enabled",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
	},
}
var OAuthProviderSettingsAdvScopeBasicAuthUserNameIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "adv_scope_validation_enabled",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "adv_scope_url_security_enabled",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "adv_scope_url_security",
			AttrType:    "DmSecurityType",
			AttrDefault: "basic-auth",
			Value:       []string{"basic-auth"},
		},
	},
}
var OAuthProviderSettingsAdvScopeBasicAuthPasswordIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "adv_scope_validation_enabled",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "adv_scope_url_security_enabled",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "adv_scope_url_security",
			AttrType:    "DmSecurityType",
			AttrDefault: "basic-auth",
			Value:       []string{"basic-auth"},
		},
	},
}
var OAuthProviderSettingsAdvScopeBasicAuthHeaderNameIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "adv_scope_validation_enabled",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "adv_scope_url_security_enabled",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "adv_scope_url_security",
			AttrType:    "DmSecurityType",
			AttrDefault: "basic-auth",
			Value:       []string{"basic-auth"},
		},
	},
}
var OAuthProviderSettingsAdvancedScopeCustomHeadersIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "adv_scope_validation_enabled",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var OAuthProviderSettingsAdvancedScopeCustomContextsIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "adv_scope_validation_enabled",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var OAuthProviderSettingsAPICEnableOIDCIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{
				{
					Evaluation:  "property-value-not-in-list",
					Attribute:   "supported_grant_types",
					AttrType:    "DmOAuthProviderGrantType",
					AttrDefault: "",
					Value:       []string{"access_code"},
				},
				{
					Evaluation:  "property-value-not-in-list",
					Attribute:   "supported_grant_types",
					AttrType:    "DmOAuthProviderGrantType",
					AttrDefault: "",
					Value:       []string{"implicit"},
				},
			},
		},
	},
}
var OAuthProviderSettingsAPICOIDCHybridResponseTypesIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{
				{
					Evaluation:  "property-value-not-in-list",
					Attribute:   "supported_grant_types",
					AttrType:    "DmOAuthProviderGrantType",
					AttrDefault: "",
					Value:       []string{"access_code"},
				},
				{
					Evaluation:  "property-value-not-in-list",
					Attribute:   "supported_grant_types",
					AttrType:    "DmOAuthProviderGrantType",
					AttrDefault: "",
					Value:       []string{"implicit"},
				},
			},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "api_c_enable_oidc",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
	},
}
var OAuthProviderSettingsAPICSupportPKCEIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "supported_grant_types",
			AttrType:    "DmOAuthProviderGrantType",
			AttrDefault: "",
			Value:       []string{"access_code"},
		},
	},
}
var OAuthProviderSettingsAPICRequirePKCEIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "supported_grant_types",
			AttrType:    "DmOAuthProviderGrantType",
			AttrDefault: "",
			Value:       []string{"access_code"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "api_c_support_pkce",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
	},
}
var OAuthProviderSettingsAPICSupportPKCEPlainIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "supported_grant_types",
			AttrType:    "DmOAuthProviderGrantType",
			AttrDefault: "",
			Value:       []string{"access_code"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "api_c_support_pkce",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
	},
}
var OAuthProviderSettingsAPICTokenTypeToGenerateIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "provider_type",
	AttrType:    "String",
	AttrDefault: "native",
	Value:       []string{"native"},
}
var OAuthProviderSettingsMetadataFromIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "provider_type",
	AttrType:    "String",
	AttrDefault: "native",
	Value:       []string{"native"},
}
var OAuthProviderSettingsMetadataURLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var OAuthProviderSettingsMetadataSSLProfileIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "metadata_from",
			AttrType:    "DmMetadataFromType",
			AttrDefault: "",
			Value:       []string{"external_url"},
		},
	},
}
var OAuthProviderSettingsMetadataHeaderForAccessTokenIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "metadata_from",
			AttrType:    "DmMetadataFromType",
			AttrDefault: "",
			Value:       []string{"authentication_url", "external_url"},
		},
	},
}
var OAuthProviderSettingsMetadataHeaderForPayloadIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "metadata_from",
			AttrType:    "DmMetadataFromType",
			AttrDefault: "",
			Value:       []string{"authentication_url", "external_url"},
		},
	},
}
var OAuthProviderSettingsExternalRevocationURLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var OAuthProviderSettingsExternalRevocationSSLProfileIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "enable_token_management",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "token_manager_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"external"},
		},
	},
}
var OAuthProviderSettingsExternalRevocationURLSecurityIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "enable_token_management",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "token_manager_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"external"},
		},
	},
}
var OAuthProviderSettingsExternalRevocationBasicAuthUserNameIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "enable_token_management",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "token_manager_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"external"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "external_revocation_url_security",
			AttrType:    "DmSecurityType",
			AttrDefault: "basic-auth",
			Value:       []string{"basic-auth"},
		},
	},
}
var OAuthProviderSettingsExternalRevocationBasicAuthPasswordIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "enable_token_management",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "token_manager_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"external"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "external_revocation_url_security",
			AttrType:    "DmSecurityType",
			AttrDefault: "basic-auth",
			Value:       []string{"basic-auth"},
		},
	},
}
var OAuthProviderSettingsExternalRevocationBasicAuthHeaderNameIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "enable_token_management",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "token_manager_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"external"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "external_revocation_url_security",
			AttrType:    "DmSecurityType",
			AttrDefault: "basic-auth",
			Value:       []string{"basic-auth"},
		},
	},
}
var OAuthProviderSettingsExternalRevocationCustomHeaderFormatIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "enable_token_management",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "token_manager_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"external"},
		},
	},
}
var OAuthProviderSettingsExternalRevocationCacheTypeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "enable_token_management",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "token_manager_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"external"},
		},
	},
}
var OAuthProviderSettingsExternalRevocationCacheTimeToLiveIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "enable_token_management",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "token_manager_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"external"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "external_revocation_cache_type",
			AttrType:    "String",
			AttrDefault: "NoCache",
			Value:       []string{"TimeToLive"},
		},
	},
}
var OAuthProviderSettingsExternalRevocationFailOnErrorIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"native"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "enable_token_management",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "token_manager_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"external"},
		},
	},
}
var OAuthProviderSettingsEnableTokenManagementIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "provider_type",
	AttrType:    "String",
	AttrDefault: "native",
	Value:       []string{"third_party"},
}
var OAuthProviderSettingsTokenManagerTypeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "enable_token_management",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"third_party"},
		},
	},
}
var OAuthProviderSettingsAPISecurityTokenManagerIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var OAuthProviderSettingsEnableApplicationRevocationIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var OAuthProviderSettingsApplicationRevocationEndpointIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var OAuthProviderSettingsEnableOwnerRevocationIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var OAuthProviderSettingsOwnerRevocationEndpointIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var OAuthProviderSettingsTokenValidationReqIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var OAuthProviderSettingsThirdPartyAZURLIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "provider_type",
	AttrType:    "String",
	AttrDefault: "native",
	Value:       []string{"third_party"},
}
var OAuthProviderSettingsThirdPartyTokenURLIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "provider_type",
	AttrType:    "String",
	AttrDefault: "native",
	Value:       []string{"third_party"},
}
var OAuthProviderSettingsThirdPartyIntrospectURLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var OAuthProviderSettingsThirdPartyIntrospectCacheTypeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var OAuthProviderSettingsThirdPartyIntrospectCacheTimeToLiveIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var OAuthProviderSettingsThirdPartyAuthorizationHeaderPassThruIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "provider_type",
	AttrType:    "String",
	AttrDefault: "native",
	Value:       []string{"native"},
}
var OAuthProviderSettingsThirdPartyIntrospectURLSecurityIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "provider_type",
	AttrType:    "String",
	AttrDefault: "native",
	Value:       []string{"native"},
}
var OAuthProviderSettingsThirdPartyIntrospectBasicAuthUserNameIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"third_party"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "third_party_introspect_url_security",
			AttrType:    "DmSecurityType",
			AttrDefault: "basic-auth",
			Value:       []string{"basic-auth"},
		},
	},
}
var OAuthProviderSettingsThirdPartyIntrospectBasicAuthPasswordIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"third_party"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "third_party_introspect_url_security",
			AttrType:    "DmSecurityType",
			AttrDefault: "basic-auth",
			Value:       []string{"basic-auth"},
		},
	},
}
var OAuthProviderSettingsThirdPartyBasicAuthHeaderNameIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "provider_type",
			AttrType:    "String",
			AttrDefault: "native",
			Value:       []string{"third_party"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "third_party_introspect_url_security",
			AttrType:    "DmSecurityType",
			AttrDefault: "basic-auth",
			Value:       []string{"basic-auth"},
		},
	},
}
var OAuthProviderSettingsThirdPartyCustomHeaderNameFormatIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "provider_type",
	AttrType:    "String",
	AttrDefault: "native",
	Value:       []string{"third_party"},
}
var OAuthProviderSettingsThirdPartyIntrospectSSLProfileIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "provider_type",
	AttrType:    "String",
	AttrDefault: "native",
	Value:       []string{"native"},
}

var OAuthProviderSettingsObjectType = map[string]attr.Type{
	"id":                                          types.StringType,
	"app_domain":                                  types.StringType,
	"user_summary":                                types.StringType,
	"enable_debug_mode":                           types.BoolType,
	"provider_type":                               types.StringType,
	"scopes_allowed":                              types.StringType,
	"default_scopes":                              types.StringType,
	"supported_grant_types":                       types.ObjectType{AttrTypes: DmOAuthProviderGrantTypeObjectType},
	"supported_client_types":                      types.ObjectType{AttrTypes: DmAllowedClientTypeObjectType},
	"api_c_provider_base_path":                    types.StringType,
	"api_c_authorize_endpoint":                    types.StringType,
	"api_c_token_endpoint":                        types.StringType,
	"api_c_enable_introspection":                  types.BoolType,
	"api_c_introspect_endpoint":                   types.StringType,
	"api_c_token_secret":                          types.StringType,
	"api_c_one_time_use_accesstoken":              types.BoolType,
	"api_c_access_token_ttl":                      types.Int64Type,
	"api_c_auth_code_ttl":                         types.Int64Type,
	"api_c_enable_refresh_token":                  types.BoolType,
	"api_c_one_time_use_refreshtoken":             types.BoolType,
	"api_c_refresh_token_limit":                   types.Int64Type,
	"api_c_refresh_token_ttl":                     types.Int64Type,
	"api_c_maximum_consent_ttl":                   types.Int64Type,
	"adv_scope_validation_enabled":                types.BoolType,
	"advanced_scope_url_override":                 types.BoolType,
	"advanced_scope_url":                          types.StringType,
	"adv_scope_tls_profile":                       types.StringType,
	"adv_scope_url_security_enabled":              types.BoolType,
	"adv_scope_url_security":                      types.ObjectType{AttrTypes: DmSecurityTypeObjectType},
	"adv_scope_basic_auth_user_name":              types.StringType,
	"adv_scope_basic_auth_password":               types.StringType,
	"adv_scope_basic_auth_header_name":            types.StringType,
	"advanced_scope_custom_headers":               types.StringType,
	"advanced_scope_custom_contexts":              types.StringType,
	"api_c_enable_oidc":                           types.BoolType,
	"api_c_oidc_hybrid_response_types":            types.ObjectType{AttrTypes: DmOIDCHybridResponseTypeObjectType},
	"api_c_support_pkce":                          types.BoolType,
	"api_c_require_pkce":                          types.BoolType,
	"api_c_support_pkce_plain":                    types.BoolType,
	"api_c_token_type_to_generate":                types.StringType,
	"metadata_from":                               types.ObjectType{AttrTypes: DmMetadataFromTypeObjectType},
	"metadata_url":                                types.StringType,
	"metadata_ssl_profile":                        types.StringType,
	"metadata_header_for_access_token":            types.StringType,
	"metadata_header_for_payload":                 types.StringType,
	"external_revocation_url":                     types.StringType,
	"external_revocation_ssl_profile":             types.StringType,
	"external_revocation_url_security":            types.ObjectType{AttrTypes: DmSecurityTypeObjectType},
	"external_revocation_basic_auth_user_name":    types.StringType,
	"external_revocation_basic_auth_password":     types.StringType,
	"external_revocation_basic_auth_header_name":  types.StringType,
	"external_revocation_custom_header_format":    types.StringType,
	"external_revocation_cache_type":              types.StringType,
	"external_revocation_cache_time_to_live":      types.Int64Type,
	"external_revocation_fail_on_error":           types.BoolType,
	"enable_token_management":                     types.BoolType,
	"token_manager_type":                          types.StringType,
	"api_security_token_manager":                  types.StringType,
	"enable_application_revocation":               types.BoolType,
	"application_revocation_endpoint":             types.StringType,
	"enable_owner_revocation":                     types.BoolType,
	"owner_revocation_endpoint":                   types.StringType,
	"token_validation_req":                        types.StringType,
	"third_party_az_url":                          types.StringType,
	"third_party_token_url":                       types.StringType,
	"third_party_introspect_url":                  types.StringType,
	"third_party_introspect_cache_type":           types.StringType,
	"third_party_introspect_cache_time_to_live":   types.Int64Type,
	"third_party_authorization_header_pass_thru":  types.BoolType,
	"third_party_introspect_url_security":         types.ObjectType{AttrTypes: DmSecurityTypeObjectType},
	"third_party_introspect_basic_auth_user_name": types.StringType,
	"third_party_introspect_basic_auth_password":  types.StringType,
	"third_party_basic_auth_header_name":          types.StringType,
	"third_party_custom_header_name_format":       types.StringType,
	"third_party_introspect_ssl_profile":          types.StringType,
	"dependency_actions":                          actions.ActionsListType,
}

func (data OAuthProviderSettings) GetPath() string {
	rest_path := "/mgmt/config/{domain}/OAuthProviderSettings"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data OAuthProviderSettings) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.EnableDebugMode.IsNull() {
		return false
	}
	if !data.ProviderType.IsNull() {
		return false
	}
	if !data.ScopesAllowed.IsNull() {
		return false
	}
	if !data.DefaultScopes.IsNull() {
		return false
	}
	if data.SupportedGrantTypes != nil {
		if !data.SupportedGrantTypes.IsNull() {
			return false
		}
	}
	if data.SupportedClientTypes != nil {
		if !data.SupportedClientTypes.IsNull() {
			return false
		}
	}
	if !data.ApiCProviderBasePath.IsNull() {
		return false
	}
	if !data.ApiCAuthorizeEndpoint.IsNull() {
		return false
	}
	if !data.ApiCTokenEndpoint.IsNull() {
		return false
	}
	if !data.ApiCEnableIntrospection.IsNull() {
		return false
	}
	if !data.ApiCIntrospectEndpoint.IsNull() {
		return false
	}
	if !data.ApiCTokenSecret.IsNull() {
		return false
	}
	if !data.ApiCOneTimeUseAccesstoken.IsNull() {
		return false
	}
	if !data.ApiCAccessTokenTtl.IsNull() {
		return false
	}
	if !data.ApiCAuthCodeTtl.IsNull() {
		return false
	}
	if !data.ApiCEnableRefreshToken.IsNull() {
		return false
	}
	if !data.ApiCOneTimeUseRefreshtoken.IsNull() {
		return false
	}
	if !data.ApiCRefreshTokenLimit.IsNull() {
		return false
	}
	if !data.ApiCRefreshTokenTtl.IsNull() {
		return false
	}
	if !data.ApiCMaximumConsentTtl.IsNull() {
		return false
	}
	if !data.AdvScopeValidationEnabled.IsNull() {
		return false
	}
	if !data.AdvancedScopeUrlOverride.IsNull() {
		return false
	}
	if !data.AdvancedScopeUrl.IsNull() {
		return false
	}
	if !data.AdvScopeTlsProfile.IsNull() {
		return false
	}
	if !data.AdvScopeUrlSecurityEnabled.IsNull() {
		return false
	}
	if data.AdvScopeUrlSecurity != nil {
		if !data.AdvScopeUrlSecurity.IsNull() {
			return false
		}
	}
	if !data.AdvScopeBasicAuthUserName.IsNull() {
		return false
	}
	if !data.AdvScopeBasicAuthPassword.IsNull() {
		return false
	}
	if !data.AdvScopeBasicAuthHeaderName.IsNull() {
		return false
	}
	if !data.AdvancedScopeCustomHeaders.IsNull() {
		return false
	}
	if !data.AdvancedScopeCustomContexts.IsNull() {
		return false
	}
	if !data.ApiCEnableOidc.IsNull() {
		return false
	}
	if data.ApiCOidcHybridResponseTypes != nil {
		if !data.ApiCOidcHybridResponseTypes.IsNull() {
			return false
		}
	}
	if !data.ApiCSupportPkce.IsNull() {
		return false
	}
	if !data.ApiCRequirePkce.IsNull() {
		return false
	}
	if !data.ApiCSupportPkcePlain.IsNull() {
		return false
	}
	if !data.ApiCTokenTypeToGenerate.IsNull() {
		return false
	}
	if data.MetadataFrom != nil {
		if !data.MetadataFrom.IsNull() {
			return false
		}
	}
	if !data.MetadataUrl.IsNull() {
		return false
	}
	if !data.MetadataSslProfile.IsNull() {
		return false
	}
	if !data.MetadataHeaderForAccessToken.IsNull() {
		return false
	}
	if !data.MetadataHeaderForPayload.IsNull() {
		return false
	}
	if !data.ExternalRevocationUrl.IsNull() {
		return false
	}
	if !data.ExternalRevocationSslProfile.IsNull() {
		return false
	}
	if data.ExternalRevocationUrlSecurity != nil {
		if !data.ExternalRevocationUrlSecurity.IsNull() {
			return false
		}
	}
	if !data.ExternalRevocationBasicAuthUserName.IsNull() {
		return false
	}
	if !data.ExternalRevocationBasicAuthPassword.IsNull() {
		return false
	}
	if !data.ExternalRevocationBasicAuthHeaderName.IsNull() {
		return false
	}
	if !data.ExternalRevocationCustomHeaderFormat.IsNull() {
		return false
	}
	if !data.ExternalRevocationCacheType.IsNull() {
		return false
	}
	if !data.ExternalRevocationCacheTimeToLive.IsNull() {
		return false
	}
	if !data.ExternalRevocationFailOnError.IsNull() {
		return false
	}
	if !data.EnableTokenManagement.IsNull() {
		return false
	}
	if !data.TokenManagerType.IsNull() {
		return false
	}
	if !data.ApiSecurityTokenManager.IsNull() {
		return false
	}
	if !data.EnableApplicationRevocation.IsNull() {
		return false
	}
	if !data.ApplicationRevocationEndpoint.IsNull() {
		return false
	}
	if !data.EnableOwnerRevocation.IsNull() {
		return false
	}
	if !data.OwnerRevocationEndpoint.IsNull() {
		return false
	}
	if !data.TokenValidationReq.IsNull() {
		return false
	}
	if !data.ThirdPartyAzUrl.IsNull() {
		return false
	}
	if !data.ThirdPartyTokenUrl.IsNull() {
		return false
	}
	if !data.ThirdPartyIntrospectUrl.IsNull() {
		return false
	}
	if !data.ThirdPartyIntrospectCacheType.IsNull() {
		return false
	}
	if !data.ThirdPartyIntrospectCacheTimeToLive.IsNull() {
		return false
	}
	if !data.ThirdPartyAuthorizationHeaderPassThru.IsNull() {
		return false
	}
	if data.ThirdPartyIntrospectUrlSecurity != nil {
		if !data.ThirdPartyIntrospectUrlSecurity.IsNull() {
			return false
		}
	}
	if !data.ThirdPartyIntrospectBasicAuthUserName.IsNull() {
		return false
	}
	if !data.ThirdPartyIntrospectBasicAuthPassword.IsNull() {
		return false
	}
	if !data.ThirdPartyBasicAuthHeaderName.IsNull() {
		return false
	}
	if !data.ThirdPartyCustomHeaderNameFormat.IsNull() {
		return false
	}
	if !data.ThirdPartyIntrospectSslProfile.IsNull() {
		return false
	}
	return true
}

func (data OAuthProviderSettings) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.EnableDebugMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnableDebugMode`, tfutils.StringFromBool(data.EnableDebugMode, ""))
	}
	if !data.ProviderType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ProviderType`, data.ProviderType.ValueString())
	}
	if !data.ScopesAllowed.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ScopesAllowed`, data.ScopesAllowed.ValueString())
	}
	if !data.DefaultScopes.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DefaultScopes`, data.DefaultScopes.ValueString())
	}
	if data.SupportedGrantTypes != nil {
		if !data.SupportedGrantTypes.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`SupportedGrantTypes`, data.SupportedGrantTypes.ToBody(ctx, ""))
		}
	}
	if data.SupportedClientTypes != nil {
		if !data.SupportedClientTypes.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`SupportedClientTypes`, data.SupportedClientTypes.ToBody(ctx, ""))
		}
	}
	if !data.ApiCProviderBasePath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APICProviderBasePath`, data.ApiCProviderBasePath.ValueString())
	}
	if !data.ApiCAuthorizeEndpoint.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APICAuthorizeEndpoint`, data.ApiCAuthorizeEndpoint.ValueString())
	}
	if !data.ApiCTokenEndpoint.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APICTokenEndpoint`, data.ApiCTokenEndpoint.ValueString())
	}
	if !data.ApiCEnableIntrospection.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APICEnableIntrospection`, tfutils.StringFromBool(data.ApiCEnableIntrospection, ""))
	}
	if !data.ApiCIntrospectEndpoint.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APICIntrospectEndpoint`, data.ApiCIntrospectEndpoint.ValueString())
	}
	if !data.ApiCTokenSecret.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APICTokenSecret`, data.ApiCTokenSecret.ValueString())
	}
	if !data.ApiCOneTimeUseAccesstoken.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APICOneTimeUseAccesstoken`, tfutils.StringFromBool(data.ApiCOneTimeUseAccesstoken, ""))
	}
	if !data.ApiCAccessTokenTtl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APICAccessTokenTTL`, data.ApiCAccessTokenTtl.ValueInt64())
	}
	if !data.ApiCAuthCodeTtl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APICAuthCodeTTL`, data.ApiCAuthCodeTtl.ValueInt64())
	}
	if !data.ApiCEnableRefreshToken.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APICEnableRefreshToken`, tfutils.StringFromBool(data.ApiCEnableRefreshToken, ""))
	}
	if !data.ApiCOneTimeUseRefreshtoken.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APICOneTimeUseRefreshtoken`, tfutils.StringFromBool(data.ApiCOneTimeUseRefreshtoken, ""))
	}
	if !data.ApiCRefreshTokenLimit.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APICRefreshTokenLimit`, data.ApiCRefreshTokenLimit.ValueInt64())
	}
	if !data.ApiCRefreshTokenTtl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APICRefreshTokenTTL`, data.ApiCRefreshTokenTtl.ValueInt64())
	}
	if !data.ApiCMaximumConsentTtl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APICMaximumConsentTTL`, data.ApiCMaximumConsentTtl.ValueInt64())
	}
	if !data.AdvScopeValidationEnabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AdvScopeValidationEnabled`, tfutils.StringFromBool(data.AdvScopeValidationEnabled, ""))
	}
	if !data.AdvancedScopeUrlOverride.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AdvancedScopeURLOverride`, tfutils.StringFromBool(data.AdvancedScopeUrlOverride, ""))
	}
	if !data.AdvancedScopeUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AdvancedScopeURL`, data.AdvancedScopeUrl.ValueString())
	}
	if !data.AdvScopeTlsProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AdvScopeTLSProfile`, data.AdvScopeTlsProfile.ValueString())
	}
	if !data.AdvScopeUrlSecurityEnabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AdvScopeURLSecurityEnabled`, tfutils.StringFromBool(data.AdvScopeUrlSecurityEnabled, ""))
	}
	if data.AdvScopeUrlSecurity != nil {
		if !data.AdvScopeUrlSecurity.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`AdvScopeURLSecurity`, data.AdvScopeUrlSecurity.ToBody(ctx, ""))
		}
	}
	if !data.AdvScopeBasicAuthUserName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AdvScopeBasicAuthUserName`, data.AdvScopeBasicAuthUserName.ValueString())
	}
	if !data.AdvScopeBasicAuthPassword.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AdvScopeBasicAuthPassword`, data.AdvScopeBasicAuthPassword.ValueString())
	}
	if !data.AdvScopeBasicAuthHeaderName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AdvScopeBasicAuthHeaderName`, data.AdvScopeBasicAuthHeaderName.ValueString())
	}
	if !data.AdvancedScopeCustomHeaders.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AdvancedScopeCustomHeaders`, data.AdvancedScopeCustomHeaders.ValueString())
	}
	if !data.AdvancedScopeCustomContexts.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AdvancedScopeCustomContexts`, data.AdvancedScopeCustomContexts.ValueString())
	}
	if !data.ApiCEnableOidc.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APICEnableOIDC`, tfutils.StringFromBool(data.ApiCEnableOidc, ""))
	}
	if data.ApiCOidcHybridResponseTypes != nil {
		if !data.ApiCOidcHybridResponseTypes.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`APICOIDCHybridResponseTypes`, data.ApiCOidcHybridResponseTypes.ToBody(ctx, ""))
		}
	}
	if !data.ApiCSupportPkce.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APICSupportPKCE`, tfutils.StringFromBool(data.ApiCSupportPkce, ""))
	}
	if !data.ApiCRequirePkce.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APICRequirePKCE`, tfutils.StringFromBool(data.ApiCRequirePkce, ""))
	}
	if !data.ApiCSupportPkcePlain.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APICSupportPKCEPlain`, tfutils.StringFromBool(data.ApiCSupportPkcePlain, ""))
	}
	if !data.ApiCTokenTypeToGenerate.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APICTokenTypeToGenerate`, data.ApiCTokenTypeToGenerate.ValueString())
	}
	if data.MetadataFrom != nil {
		if !data.MetadataFrom.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`MetadataFrom`, data.MetadataFrom.ToBody(ctx, ""))
		}
	}
	if !data.MetadataUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MetadataURL`, data.MetadataUrl.ValueString())
	}
	if !data.MetadataSslProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MetadataSSLProfile`, data.MetadataSslProfile.ValueString())
	}
	if !data.MetadataHeaderForAccessToken.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MetadataHeaderForAccessToken`, data.MetadataHeaderForAccessToken.ValueString())
	}
	if !data.MetadataHeaderForPayload.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MetadataHeaderForPayload`, data.MetadataHeaderForPayload.ValueString())
	}
	if !data.ExternalRevocationUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ExternalRevocationURL`, data.ExternalRevocationUrl.ValueString())
	}
	if !data.ExternalRevocationSslProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ExternalRevocationSSLProfile`, data.ExternalRevocationSslProfile.ValueString())
	}
	if data.ExternalRevocationUrlSecurity != nil {
		if !data.ExternalRevocationUrlSecurity.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`ExternalRevocationURLSecurity`, data.ExternalRevocationUrlSecurity.ToBody(ctx, ""))
		}
	}
	if !data.ExternalRevocationBasicAuthUserName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ExternalRevocationBasicAuthUserName`, data.ExternalRevocationBasicAuthUserName.ValueString())
	}
	if !data.ExternalRevocationBasicAuthPassword.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ExternalRevocationBasicAuthPassword`, data.ExternalRevocationBasicAuthPassword.ValueString())
	}
	if !data.ExternalRevocationBasicAuthHeaderName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ExternalRevocationBasicAuthHeaderName`, data.ExternalRevocationBasicAuthHeaderName.ValueString())
	}
	if !data.ExternalRevocationCustomHeaderFormat.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ExternalRevocationCustomHeaderFormat`, data.ExternalRevocationCustomHeaderFormat.ValueString())
	}
	if !data.ExternalRevocationCacheType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ExternalRevocationCacheType`, data.ExternalRevocationCacheType.ValueString())
	}
	if !data.ExternalRevocationCacheTimeToLive.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ExternalRevocationCacheTimeToLive`, data.ExternalRevocationCacheTimeToLive.ValueInt64())
	}
	if !data.ExternalRevocationFailOnError.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ExternalRevocationFailOnError`, tfutils.StringFromBool(data.ExternalRevocationFailOnError, ""))
	}
	if !data.EnableTokenManagement.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnableTokenManagement`, tfutils.StringFromBool(data.EnableTokenManagement, ""))
	}
	if !data.TokenManagerType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TokenManagerType`, data.TokenManagerType.ValueString())
	}
	if !data.ApiSecurityTokenManager.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APISecurityTokenManager`, data.ApiSecurityTokenManager.ValueString())
	}
	if !data.EnableApplicationRevocation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnableApplicationRevocation`, tfutils.StringFromBool(data.EnableApplicationRevocation, ""))
	}
	if !data.ApplicationRevocationEndpoint.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ApplicationRevocationEndpoint`, data.ApplicationRevocationEndpoint.ValueString())
	}
	if !data.EnableOwnerRevocation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnableOwnerRevocation`, tfutils.StringFromBool(data.EnableOwnerRevocation, ""))
	}
	if !data.OwnerRevocationEndpoint.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OwnerRevocationEndpoint`, data.OwnerRevocationEndpoint.ValueString())
	}
	if !data.TokenValidationReq.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TokenValidationReq`, data.TokenValidationReq.ValueString())
	}
	if !data.ThirdPartyAzUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ThirdPartyAZURL`, data.ThirdPartyAzUrl.ValueString())
	}
	if !data.ThirdPartyTokenUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ThirdPartyTokenURL`, data.ThirdPartyTokenUrl.ValueString())
	}
	if !data.ThirdPartyIntrospectUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ThirdPartyIntrospectURL`, data.ThirdPartyIntrospectUrl.ValueString())
	}
	if !data.ThirdPartyIntrospectCacheType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ThirdPartyIntrospectCacheType`, data.ThirdPartyIntrospectCacheType.ValueString())
	}
	if !data.ThirdPartyIntrospectCacheTimeToLive.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ThirdPartyIntrospectCacheTimeToLive`, data.ThirdPartyIntrospectCacheTimeToLive.ValueInt64())
	}
	if !data.ThirdPartyAuthorizationHeaderPassThru.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ThirdPartyAuthorizationHeaderPassThru`, tfutils.StringFromBool(data.ThirdPartyAuthorizationHeaderPassThru, ""))
	}
	if data.ThirdPartyIntrospectUrlSecurity != nil {
		if !data.ThirdPartyIntrospectUrlSecurity.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`ThirdPartyIntrospectURLSecurity`, data.ThirdPartyIntrospectUrlSecurity.ToBody(ctx, ""))
		}
	}
	if !data.ThirdPartyIntrospectBasicAuthUserName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ThirdPartyIntrospectBasicAuthUserName`, data.ThirdPartyIntrospectBasicAuthUserName.ValueString())
	}
	if !data.ThirdPartyIntrospectBasicAuthPassword.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ThirdPartyIntrospectBasicAuthPassword`, data.ThirdPartyIntrospectBasicAuthPassword.ValueString())
	}
	if !data.ThirdPartyBasicAuthHeaderName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ThirdPartyBasicAuthHeaderName`, data.ThirdPartyBasicAuthHeaderName.ValueString())
	}
	if !data.ThirdPartyCustomHeaderNameFormat.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ThirdPartyCustomHeaderNameFormat`, data.ThirdPartyCustomHeaderNameFormat.ValueString())
	}
	if !data.ThirdPartyIntrospectSslProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ThirdPartyIntrospectSSLProfile`, data.ThirdPartyIntrospectSslProfile.ValueString())
	}
	return body
}

func (data *OAuthProviderSettings) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `EnableDebugMode`); value.Exists() {
		data.EnableDebugMode = tfutils.BoolFromString(value.String())
	} else {
		data.EnableDebugMode = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ProviderType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ProviderType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProviderType = types.StringValue("native")
	}
	if value := res.Get(pathRoot + `ScopesAllowed`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ScopesAllowed = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ScopesAllowed = types.StringNull()
	}
	if value := res.Get(pathRoot + `DefaultScopes`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DefaultScopes = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DefaultScopes = types.StringNull()
	}
	if value := res.Get(pathRoot + `SupportedGrantTypes`); value.Exists() {
		data.SupportedGrantTypes = &DmOAuthProviderGrantType{}
		data.SupportedGrantTypes.FromBody(ctx, "", value)
	} else {
		data.SupportedGrantTypes = nil
	}
	if value := res.Get(pathRoot + `SupportedClientTypes`); value.Exists() {
		data.SupportedClientTypes = &DmAllowedClientType{}
		data.SupportedClientTypes.FromBody(ctx, "", value)
	} else {
		data.SupportedClientTypes = nil
	}
	if value := res.Get(pathRoot + `APICProviderBasePath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ApiCProviderBasePath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ApiCProviderBasePath = types.StringValue("/")
	}
	if value := res.Get(pathRoot + `APICAuthorizeEndpoint`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ApiCAuthorizeEndpoint = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ApiCAuthorizeEndpoint = types.StringValue("/oauth2/authorize")
	}
	if value := res.Get(pathRoot + `APICTokenEndpoint`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ApiCTokenEndpoint = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ApiCTokenEndpoint = types.StringValue("/oauth2/token")
	}
	if value := res.Get(pathRoot + `APICEnableIntrospection`); value.Exists() {
		data.ApiCEnableIntrospection = tfutils.BoolFromString(value.String())
	} else {
		data.ApiCEnableIntrospection = types.BoolNull()
	}
	if value := res.Get(pathRoot + `APICIntrospectEndpoint`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ApiCIntrospectEndpoint = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ApiCIntrospectEndpoint = types.StringValue("/oauth2/introspect")
	}
	if value := res.Get(pathRoot + `APICTokenSecret`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ApiCTokenSecret = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ApiCTokenSecret = types.StringNull()
	}
	if value := res.Get(pathRoot + `APICOneTimeUseAccesstoken`); value.Exists() {
		data.ApiCOneTimeUseAccesstoken = tfutils.BoolFromString(value.String())
	} else {
		data.ApiCOneTimeUseAccesstoken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `APICAccessTokenTTL`); value.Exists() {
		data.ApiCAccessTokenTtl = types.Int64Value(value.Int())
	} else {
		data.ApiCAccessTokenTtl = types.Int64Value(3600)
	}
	if value := res.Get(pathRoot + `APICAuthCodeTTL`); value.Exists() {
		data.ApiCAuthCodeTtl = types.Int64Value(value.Int())
	} else {
		data.ApiCAuthCodeTtl = types.Int64Value(300)
	}
	if value := res.Get(pathRoot + `APICEnableRefreshToken`); value.Exists() {
		data.ApiCEnableRefreshToken = tfutils.BoolFromString(value.String())
	} else {
		data.ApiCEnableRefreshToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `APICOneTimeUseRefreshtoken`); value.Exists() {
		data.ApiCOneTimeUseRefreshtoken = tfutils.BoolFromString(value.String())
	} else {
		data.ApiCOneTimeUseRefreshtoken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `APICRefreshTokenLimit`); value.Exists() {
		data.ApiCRefreshTokenLimit = types.Int64Value(value.Int())
	} else {
		data.ApiCRefreshTokenLimit = types.Int64Value(10)
	}
	if value := res.Get(pathRoot + `APICRefreshTokenTTL`); value.Exists() {
		data.ApiCRefreshTokenTtl = types.Int64Value(value.Int())
	} else {
		data.ApiCRefreshTokenTtl = types.Int64Value(5400)
	}
	if value := res.Get(pathRoot + `APICMaximumConsentTTL`); value.Exists() {
		data.ApiCMaximumConsentTtl = types.Int64Value(value.Int())
	} else {
		data.ApiCMaximumConsentTtl = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `AdvScopeValidationEnabled`); value.Exists() {
		data.AdvScopeValidationEnabled = tfutils.BoolFromString(value.String())
	} else {
		data.AdvScopeValidationEnabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AdvancedScopeURLOverride`); value.Exists() {
		data.AdvancedScopeUrlOverride = tfutils.BoolFromString(value.String())
	} else {
		data.AdvancedScopeUrlOverride = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AdvancedScopeURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AdvancedScopeUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AdvancedScopeUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AdvScopeTLSProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AdvScopeTlsProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AdvScopeTlsProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `AdvScopeURLSecurityEnabled`); value.Exists() {
		data.AdvScopeUrlSecurityEnabled = tfutils.BoolFromString(value.String())
	} else {
		data.AdvScopeUrlSecurityEnabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AdvScopeURLSecurity`); value.Exists() {
		data.AdvScopeUrlSecurity = &DmSecurityType{}
		data.AdvScopeUrlSecurity.FromBody(ctx, "", value)
	} else {
		data.AdvScopeUrlSecurity = nil
	}
	if value := res.Get(pathRoot + `AdvScopeBasicAuthUserName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AdvScopeBasicAuthUserName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AdvScopeBasicAuthUserName = types.StringNull()
	}
	if value := res.Get(pathRoot + `AdvScopeBasicAuthPassword`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AdvScopeBasicAuthPassword = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AdvScopeBasicAuthPassword = types.StringNull()
	}
	if value := res.Get(pathRoot + `AdvScopeBasicAuthHeaderName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AdvScopeBasicAuthHeaderName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AdvScopeBasicAuthHeaderName = types.StringNull()
	}
	if value := res.Get(pathRoot + `AdvancedScopeCustomHeaders`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AdvancedScopeCustomHeaders = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AdvancedScopeCustomHeaders = types.StringNull()
	}
	if value := res.Get(pathRoot + `AdvancedScopeCustomContexts`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AdvancedScopeCustomContexts = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AdvancedScopeCustomContexts = types.StringNull()
	}
	if value := res.Get(pathRoot + `APICEnableOIDC`); value.Exists() {
		data.ApiCEnableOidc = tfutils.BoolFromString(value.String())
	} else {
		data.ApiCEnableOidc = types.BoolNull()
	}
	if value := res.Get(pathRoot + `APICOIDCHybridResponseTypes`); value.Exists() {
		data.ApiCOidcHybridResponseTypes = &DmOIDCHybridResponseType{}
		data.ApiCOidcHybridResponseTypes.FromBody(ctx, "", value)
	} else {
		data.ApiCOidcHybridResponseTypes = nil
	}
	if value := res.Get(pathRoot + `APICSupportPKCE`); value.Exists() {
		data.ApiCSupportPkce = tfutils.BoolFromString(value.String())
	} else {
		data.ApiCSupportPkce = types.BoolNull()
	}
	if value := res.Get(pathRoot + `APICRequirePKCE`); value.Exists() {
		data.ApiCRequirePkce = tfutils.BoolFromString(value.String())
	} else {
		data.ApiCRequirePkce = types.BoolNull()
	}
	if value := res.Get(pathRoot + `APICSupportPKCEPlain`); value.Exists() {
		data.ApiCSupportPkcePlain = tfutils.BoolFromString(value.String())
	} else {
		data.ApiCSupportPkcePlain = types.BoolNull()
	}
	if value := res.Get(pathRoot + `APICTokenTypeToGenerate`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ApiCTokenTypeToGenerate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ApiCTokenTypeToGenerate = types.StringValue("Bearer")
	}
	if value := res.Get(pathRoot + `MetadataFrom`); value.Exists() {
		data.MetadataFrom = &DmMetadataFromType{}
		data.MetadataFrom.FromBody(ctx, "", value)
	} else {
		data.MetadataFrom = nil
	}
	if value := res.Get(pathRoot + `MetadataURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MetadataUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MetadataUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `MetadataSSLProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MetadataSslProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MetadataSslProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `MetadataHeaderForAccessToken`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MetadataHeaderForAccessToken = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MetadataHeaderForAccessToken = types.StringValue("X-API-OAuth-Metadata-For-AccessToken")
	}
	if value := res.Get(pathRoot + `MetadataHeaderForPayload`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MetadataHeaderForPayload = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MetadataHeaderForPayload = types.StringValue("X-API-OAuth-Metadata-For-Payload")
	}
	if value := res.Get(pathRoot + `ExternalRevocationURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ExternalRevocationUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExternalRevocationUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExternalRevocationSSLProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ExternalRevocationSslProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExternalRevocationSslProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExternalRevocationURLSecurity`); value.Exists() {
		data.ExternalRevocationUrlSecurity = &DmSecurityType{}
		data.ExternalRevocationUrlSecurity.FromBody(ctx, "", value)
	} else {
		data.ExternalRevocationUrlSecurity = nil
	}
	if value := res.Get(pathRoot + `ExternalRevocationBasicAuthUserName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ExternalRevocationBasicAuthUserName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExternalRevocationBasicAuthUserName = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExternalRevocationBasicAuthPassword`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ExternalRevocationBasicAuthPassword = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExternalRevocationBasicAuthPassword = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExternalRevocationBasicAuthHeaderName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ExternalRevocationBasicAuthHeaderName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExternalRevocationBasicAuthHeaderName = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExternalRevocationCustomHeaderFormat`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ExternalRevocationCustomHeaderFormat = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExternalRevocationCustomHeaderFormat = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExternalRevocationCacheType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ExternalRevocationCacheType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExternalRevocationCacheType = types.StringValue("NoCache")
	}
	if value := res.Get(pathRoot + `ExternalRevocationCacheTimeToLive`); value.Exists() {
		data.ExternalRevocationCacheTimeToLive = types.Int64Value(value.Int())
	} else {
		data.ExternalRevocationCacheTimeToLive = types.Int64Value(900)
	}
	if value := res.Get(pathRoot + `ExternalRevocationFailOnError`); value.Exists() {
		data.ExternalRevocationFailOnError = tfutils.BoolFromString(value.String())
	} else {
		data.ExternalRevocationFailOnError = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EnableTokenManagement`); value.Exists() {
		data.EnableTokenManagement = tfutils.BoolFromString(value.String())
	} else {
		data.EnableTokenManagement = types.BoolNull()
	}
	if value := res.Get(pathRoot + `TokenManagerType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TokenManagerType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TokenManagerType = types.StringValue("native")
	}
	if value := res.Get(pathRoot + `APISecurityTokenManager`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ApiSecurityTokenManager = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ApiSecurityTokenManager = types.StringValue("default")
	}
	if value := res.Get(pathRoot + `EnableApplicationRevocation`); value.Exists() {
		data.EnableApplicationRevocation = tfutils.BoolFromString(value.String())
	} else {
		data.EnableApplicationRevocation = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ApplicationRevocationEndpoint`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ApplicationRevocationEndpoint = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ApplicationRevocationEndpoint = types.StringValue("/oauth2/revoke")
	}
	if value := res.Get(pathRoot + `EnableOwnerRevocation`); value.Exists() {
		data.EnableOwnerRevocation = tfutils.BoolFromString(value.String())
	} else {
		data.EnableOwnerRevocation = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OwnerRevocationEndpoint`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OwnerRevocationEndpoint = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OwnerRevocationEndpoint = types.StringValue("/oauth2/issued")
	}
	if value := res.Get(pathRoot + `TokenValidationReq`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TokenValidationReq = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TokenValidationReq = types.StringNull()
	}
	if value := res.Get(pathRoot + `ThirdPartyAZURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ThirdPartyAzUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ThirdPartyAzUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `ThirdPartyTokenURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ThirdPartyTokenUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ThirdPartyTokenUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `ThirdPartyIntrospectURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ThirdPartyIntrospectUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ThirdPartyIntrospectUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `ThirdPartyIntrospectCacheType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ThirdPartyIntrospectCacheType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ThirdPartyIntrospectCacheType = types.StringValue("NoCache")
	}
	if value := res.Get(pathRoot + `ThirdPartyIntrospectCacheTimeToLive`); value.Exists() {
		data.ThirdPartyIntrospectCacheTimeToLive = types.Int64Value(value.Int())
	} else {
		data.ThirdPartyIntrospectCacheTimeToLive = types.Int64Value(900)
	}
	if value := res.Get(pathRoot + `ThirdPartyAuthorizationHeaderPassThru`); value.Exists() {
		data.ThirdPartyAuthorizationHeaderPassThru = tfutils.BoolFromString(value.String())
	} else {
		data.ThirdPartyAuthorizationHeaderPassThru = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ThirdPartyIntrospectURLSecurity`); value.Exists() {
		data.ThirdPartyIntrospectUrlSecurity = &DmSecurityType{}
		data.ThirdPartyIntrospectUrlSecurity.FromBody(ctx, "", value)
	} else {
		data.ThirdPartyIntrospectUrlSecurity = nil
	}
	if value := res.Get(pathRoot + `ThirdPartyIntrospectBasicAuthUserName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ThirdPartyIntrospectBasicAuthUserName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ThirdPartyIntrospectBasicAuthUserName = types.StringNull()
	}
	if value := res.Get(pathRoot + `ThirdPartyIntrospectBasicAuthPassword`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ThirdPartyIntrospectBasicAuthPassword = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ThirdPartyIntrospectBasicAuthPassword = types.StringNull()
	}
	if value := res.Get(pathRoot + `ThirdPartyBasicAuthHeaderName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ThirdPartyBasicAuthHeaderName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ThirdPartyBasicAuthHeaderName = types.StringNull()
	}
	if value := res.Get(pathRoot + `ThirdPartyCustomHeaderNameFormat`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ThirdPartyCustomHeaderNameFormat = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ThirdPartyCustomHeaderNameFormat = types.StringNull()
	}
	if value := res.Get(pathRoot + `ThirdPartyIntrospectSSLProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ThirdPartyIntrospectSslProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ThirdPartyIntrospectSslProfile = types.StringNull()
	}
}

func (data *OAuthProviderSettings) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `EnableDebugMode`); value.Exists() && !data.EnableDebugMode.IsNull() {
		data.EnableDebugMode = tfutils.BoolFromString(value.String())
	} else if data.EnableDebugMode.ValueBool() {
		data.EnableDebugMode = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ProviderType`); value.Exists() && !data.ProviderType.IsNull() {
		data.ProviderType = tfutils.ParseStringFromGJSON(value)
	} else if data.ProviderType.ValueString() != "native" {
		data.ProviderType = types.StringNull()
	}
	if value := res.Get(pathRoot + `ScopesAllowed`); value.Exists() && !data.ScopesAllowed.IsNull() {
		data.ScopesAllowed = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ScopesAllowed = types.StringNull()
	}
	if value := res.Get(pathRoot + `DefaultScopes`); value.Exists() && !data.DefaultScopes.IsNull() {
		data.DefaultScopes = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DefaultScopes = types.StringNull()
	}
	if value := res.Get(pathRoot + `SupportedGrantTypes`); value.Exists() {
		data.SupportedGrantTypes.UpdateFromBody(ctx, "", value)
	} else {
		data.SupportedGrantTypes = nil
	}
	if value := res.Get(pathRoot + `SupportedClientTypes`); value.Exists() {
		data.SupportedClientTypes.UpdateFromBody(ctx, "", value)
	} else {
		data.SupportedClientTypes = nil
	}
	if value := res.Get(pathRoot + `APICProviderBasePath`); value.Exists() && !data.ApiCProviderBasePath.IsNull() {
		data.ApiCProviderBasePath = tfutils.ParseStringFromGJSON(value)
	} else if data.ApiCProviderBasePath.ValueString() != "/" {
		data.ApiCProviderBasePath = types.StringNull()
	}
	if value := res.Get(pathRoot + `APICAuthorizeEndpoint`); value.Exists() && !data.ApiCAuthorizeEndpoint.IsNull() {
		data.ApiCAuthorizeEndpoint = tfutils.ParseStringFromGJSON(value)
	} else if data.ApiCAuthorizeEndpoint.ValueString() != "/oauth2/authorize" {
		data.ApiCAuthorizeEndpoint = types.StringNull()
	}
	if value := res.Get(pathRoot + `APICTokenEndpoint`); value.Exists() && !data.ApiCTokenEndpoint.IsNull() {
		data.ApiCTokenEndpoint = tfutils.ParseStringFromGJSON(value)
	} else if data.ApiCTokenEndpoint.ValueString() != "/oauth2/token" {
		data.ApiCTokenEndpoint = types.StringNull()
	}
	if value := res.Get(pathRoot + `APICEnableIntrospection`); value.Exists() && !data.ApiCEnableIntrospection.IsNull() {
		data.ApiCEnableIntrospection = tfutils.BoolFromString(value.String())
	} else if data.ApiCEnableIntrospection.ValueBool() {
		data.ApiCEnableIntrospection = types.BoolNull()
	}
	if value := res.Get(pathRoot + `APICIntrospectEndpoint`); value.Exists() && !data.ApiCIntrospectEndpoint.IsNull() {
		data.ApiCIntrospectEndpoint = tfutils.ParseStringFromGJSON(value)
	} else if data.ApiCIntrospectEndpoint.ValueString() != "/oauth2/introspect" {
		data.ApiCIntrospectEndpoint = types.StringNull()
	}
	if value := res.Get(pathRoot + `APICTokenSecret`); value.Exists() && !data.ApiCTokenSecret.IsNull() {
		data.ApiCTokenSecret = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ApiCTokenSecret = types.StringNull()
	}
	if value := res.Get(pathRoot + `APICOneTimeUseAccesstoken`); value.Exists() && !data.ApiCOneTimeUseAccesstoken.IsNull() {
		data.ApiCOneTimeUseAccesstoken = tfutils.BoolFromString(value.String())
	} else if data.ApiCOneTimeUseAccesstoken.ValueBool() {
		data.ApiCOneTimeUseAccesstoken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `APICAccessTokenTTL`); value.Exists() && !data.ApiCAccessTokenTtl.IsNull() {
		data.ApiCAccessTokenTtl = types.Int64Value(value.Int())
	} else if data.ApiCAccessTokenTtl.ValueInt64() != 3600 {
		data.ApiCAccessTokenTtl = types.Int64Null()
	}
	if value := res.Get(pathRoot + `APICAuthCodeTTL`); value.Exists() && !data.ApiCAuthCodeTtl.IsNull() {
		data.ApiCAuthCodeTtl = types.Int64Value(value.Int())
	} else if data.ApiCAuthCodeTtl.ValueInt64() != 300 {
		data.ApiCAuthCodeTtl = types.Int64Null()
	}
	if value := res.Get(pathRoot + `APICEnableRefreshToken`); value.Exists() && !data.ApiCEnableRefreshToken.IsNull() {
		data.ApiCEnableRefreshToken = tfutils.BoolFromString(value.String())
	} else if data.ApiCEnableRefreshToken.ValueBool() {
		data.ApiCEnableRefreshToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `APICOneTimeUseRefreshtoken`); value.Exists() && !data.ApiCOneTimeUseRefreshtoken.IsNull() {
		data.ApiCOneTimeUseRefreshtoken = tfutils.BoolFromString(value.String())
	} else if !data.ApiCOneTimeUseRefreshtoken.ValueBool() {
		data.ApiCOneTimeUseRefreshtoken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `APICRefreshTokenLimit`); value.Exists() && !data.ApiCRefreshTokenLimit.IsNull() {
		data.ApiCRefreshTokenLimit = types.Int64Value(value.Int())
	} else if data.ApiCRefreshTokenLimit.ValueInt64() != 10 {
		data.ApiCRefreshTokenLimit = types.Int64Null()
	}
	if value := res.Get(pathRoot + `APICRefreshTokenTTL`); value.Exists() && !data.ApiCRefreshTokenTtl.IsNull() {
		data.ApiCRefreshTokenTtl = types.Int64Value(value.Int())
	} else if data.ApiCRefreshTokenTtl.ValueInt64() != 5400 {
		data.ApiCRefreshTokenTtl = types.Int64Null()
	}
	if value := res.Get(pathRoot + `APICMaximumConsentTTL`); value.Exists() && !data.ApiCMaximumConsentTtl.IsNull() {
		data.ApiCMaximumConsentTtl = types.Int64Value(value.Int())
	} else if data.ApiCMaximumConsentTtl.ValueInt64() != 0 {
		data.ApiCMaximumConsentTtl = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AdvScopeValidationEnabled`); value.Exists() && !data.AdvScopeValidationEnabled.IsNull() {
		data.AdvScopeValidationEnabled = tfutils.BoolFromString(value.String())
	} else if data.AdvScopeValidationEnabled.ValueBool() {
		data.AdvScopeValidationEnabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AdvancedScopeURLOverride`); value.Exists() && !data.AdvancedScopeUrlOverride.IsNull() {
		data.AdvancedScopeUrlOverride = tfutils.BoolFromString(value.String())
	} else if data.AdvancedScopeUrlOverride.ValueBool() {
		data.AdvancedScopeUrlOverride = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AdvancedScopeURL`); value.Exists() && !data.AdvancedScopeUrl.IsNull() {
		data.AdvancedScopeUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AdvancedScopeUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AdvScopeTLSProfile`); value.Exists() && !data.AdvScopeTlsProfile.IsNull() {
		data.AdvScopeTlsProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AdvScopeTlsProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `AdvScopeURLSecurityEnabled`); value.Exists() && !data.AdvScopeUrlSecurityEnabled.IsNull() {
		data.AdvScopeUrlSecurityEnabled = tfutils.BoolFromString(value.String())
	} else if data.AdvScopeUrlSecurityEnabled.ValueBool() {
		data.AdvScopeUrlSecurityEnabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AdvScopeURLSecurity`); value.Exists() {
		data.AdvScopeUrlSecurity.UpdateFromBody(ctx, "", value)
	} else {
		data.AdvScopeUrlSecurity = nil
	}
	if value := res.Get(pathRoot + `AdvScopeBasicAuthUserName`); value.Exists() && !data.AdvScopeBasicAuthUserName.IsNull() {
		data.AdvScopeBasicAuthUserName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AdvScopeBasicAuthUserName = types.StringNull()
	}
	if value := res.Get(pathRoot + `AdvScopeBasicAuthPassword`); value.Exists() && !data.AdvScopeBasicAuthPassword.IsNull() {
		data.AdvScopeBasicAuthPassword = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AdvScopeBasicAuthPassword = types.StringNull()
	}
	if value := res.Get(pathRoot + `AdvScopeBasicAuthHeaderName`); value.Exists() && !data.AdvScopeBasicAuthHeaderName.IsNull() {
		data.AdvScopeBasicAuthHeaderName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AdvScopeBasicAuthHeaderName = types.StringNull()
	}
	if value := res.Get(pathRoot + `AdvancedScopeCustomHeaders`); value.Exists() && !data.AdvancedScopeCustomHeaders.IsNull() {
		data.AdvancedScopeCustomHeaders = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AdvancedScopeCustomHeaders = types.StringNull()
	}
	if value := res.Get(pathRoot + `AdvancedScopeCustomContexts`); value.Exists() && !data.AdvancedScopeCustomContexts.IsNull() {
		data.AdvancedScopeCustomContexts = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AdvancedScopeCustomContexts = types.StringNull()
	}
	if value := res.Get(pathRoot + `APICEnableOIDC`); value.Exists() && !data.ApiCEnableOidc.IsNull() {
		data.ApiCEnableOidc = tfutils.BoolFromString(value.String())
	} else if !data.ApiCEnableOidc.ValueBool() {
		data.ApiCEnableOidc = types.BoolNull()
	}
	if value := res.Get(pathRoot + `APICOIDCHybridResponseTypes`); value.Exists() {
		data.ApiCOidcHybridResponseTypes.UpdateFromBody(ctx, "", value)
	} else {
		data.ApiCOidcHybridResponseTypes = nil
	}
	if value := res.Get(pathRoot + `APICSupportPKCE`); value.Exists() && !data.ApiCSupportPkce.IsNull() {
		data.ApiCSupportPkce = tfutils.BoolFromString(value.String())
	} else if !data.ApiCSupportPkce.ValueBool() {
		data.ApiCSupportPkce = types.BoolNull()
	}
	if value := res.Get(pathRoot + `APICRequirePKCE`); value.Exists() && !data.ApiCRequirePkce.IsNull() {
		data.ApiCRequirePkce = tfutils.BoolFromString(value.String())
	} else if data.ApiCRequirePkce.ValueBool() {
		data.ApiCRequirePkce = types.BoolNull()
	}
	if value := res.Get(pathRoot + `APICSupportPKCEPlain`); value.Exists() && !data.ApiCSupportPkcePlain.IsNull() {
		data.ApiCSupportPkcePlain = tfutils.BoolFromString(value.String())
	} else if data.ApiCSupportPkcePlain.ValueBool() {
		data.ApiCSupportPkcePlain = types.BoolNull()
	}
	if value := res.Get(pathRoot + `APICTokenTypeToGenerate`); value.Exists() && !data.ApiCTokenTypeToGenerate.IsNull() {
		data.ApiCTokenTypeToGenerate = tfutils.ParseStringFromGJSON(value)
	} else if data.ApiCTokenTypeToGenerate.ValueString() != "Bearer" {
		data.ApiCTokenTypeToGenerate = types.StringNull()
	}
	if value := res.Get(pathRoot + `MetadataFrom`); value.Exists() {
		data.MetadataFrom.UpdateFromBody(ctx, "", value)
	} else {
		data.MetadataFrom = nil
	}
	if value := res.Get(pathRoot + `MetadataURL`); value.Exists() && !data.MetadataUrl.IsNull() {
		data.MetadataUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MetadataUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `MetadataSSLProfile`); value.Exists() && !data.MetadataSslProfile.IsNull() {
		data.MetadataSslProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MetadataSslProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `MetadataHeaderForAccessToken`); value.Exists() && !data.MetadataHeaderForAccessToken.IsNull() {
		data.MetadataHeaderForAccessToken = tfutils.ParseStringFromGJSON(value)
	} else if data.MetadataHeaderForAccessToken.ValueString() != "X-API-OAuth-Metadata-For-AccessToken" {
		data.MetadataHeaderForAccessToken = types.StringNull()
	}
	if value := res.Get(pathRoot + `MetadataHeaderForPayload`); value.Exists() && !data.MetadataHeaderForPayload.IsNull() {
		data.MetadataHeaderForPayload = tfutils.ParseStringFromGJSON(value)
	} else if data.MetadataHeaderForPayload.ValueString() != "X-API-OAuth-Metadata-For-Payload" {
		data.MetadataHeaderForPayload = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExternalRevocationURL`); value.Exists() && !data.ExternalRevocationUrl.IsNull() {
		data.ExternalRevocationUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExternalRevocationUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExternalRevocationSSLProfile`); value.Exists() && !data.ExternalRevocationSslProfile.IsNull() {
		data.ExternalRevocationSslProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExternalRevocationSslProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExternalRevocationURLSecurity`); value.Exists() {
		data.ExternalRevocationUrlSecurity.UpdateFromBody(ctx, "", value)
	} else {
		data.ExternalRevocationUrlSecurity = nil
	}
	if value := res.Get(pathRoot + `ExternalRevocationBasicAuthUserName`); value.Exists() && !data.ExternalRevocationBasicAuthUserName.IsNull() {
		data.ExternalRevocationBasicAuthUserName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExternalRevocationBasicAuthUserName = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExternalRevocationBasicAuthPassword`); value.Exists() && !data.ExternalRevocationBasicAuthPassword.IsNull() {
		data.ExternalRevocationBasicAuthPassword = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExternalRevocationBasicAuthPassword = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExternalRevocationBasicAuthHeaderName`); value.Exists() && !data.ExternalRevocationBasicAuthHeaderName.IsNull() {
		data.ExternalRevocationBasicAuthHeaderName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExternalRevocationBasicAuthHeaderName = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExternalRevocationCustomHeaderFormat`); value.Exists() && !data.ExternalRevocationCustomHeaderFormat.IsNull() {
		data.ExternalRevocationCustomHeaderFormat = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExternalRevocationCustomHeaderFormat = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExternalRevocationCacheType`); value.Exists() && !data.ExternalRevocationCacheType.IsNull() {
		data.ExternalRevocationCacheType = tfutils.ParseStringFromGJSON(value)
	} else if data.ExternalRevocationCacheType.ValueString() != "NoCache" {
		data.ExternalRevocationCacheType = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExternalRevocationCacheTimeToLive`); value.Exists() && !data.ExternalRevocationCacheTimeToLive.IsNull() {
		data.ExternalRevocationCacheTimeToLive = types.Int64Value(value.Int())
	} else if data.ExternalRevocationCacheTimeToLive.ValueInt64() != 900 {
		data.ExternalRevocationCacheTimeToLive = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ExternalRevocationFailOnError`); value.Exists() && !data.ExternalRevocationFailOnError.IsNull() {
		data.ExternalRevocationFailOnError = tfutils.BoolFromString(value.String())
	} else if !data.ExternalRevocationFailOnError.ValueBool() {
		data.ExternalRevocationFailOnError = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EnableTokenManagement`); value.Exists() && !data.EnableTokenManagement.IsNull() {
		data.EnableTokenManagement = tfutils.BoolFromString(value.String())
	} else if !data.EnableTokenManagement.ValueBool() {
		data.EnableTokenManagement = types.BoolNull()
	}
	if value := res.Get(pathRoot + `TokenManagerType`); value.Exists() && !data.TokenManagerType.IsNull() {
		data.TokenManagerType = tfutils.ParseStringFromGJSON(value)
	} else if data.TokenManagerType.ValueString() != "native" {
		data.TokenManagerType = types.StringNull()
	}
	if value := res.Get(pathRoot + `APISecurityTokenManager`); value.Exists() && !data.ApiSecurityTokenManager.IsNull() {
		data.ApiSecurityTokenManager = tfutils.ParseStringFromGJSON(value)
	} else if data.ApiSecurityTokenManager.ValueString() != "default" {
		data.ApiSecurityTokenManager = types.StringNull()
	}
	if value := res.Get(pathRoot + `EnableApplicationRevocation`); value.Exists() && !data.EnableApplicationRevocation.IsNull() {
		data.EnableApplicationRevocation = tfutils.BoolFromString(value.String())
	} else if data.EnableApplicationRevocation.ValueBool() {
		data.EnableApplicationRevocation = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ApplicationRevocationEndpoint`); value.Exists() && !data.ApplicationRevocationEndpoint.IsNull() {
		data.ApplicationRevocationEndpoint = tfutils.ParseStringFromGJSON(value)
	} else if data.ApplicationRevocationEndpoint.ValueString() != "/oauth2/revoke" {
		data.ApplicationRevocationEndpoint = types.StringNull()
	}
	if value := res.Get(pathRoot + `EnableOwnerRevocation`); value.Exists() && !data.EnableOwnerRevocation.IsNull() {
		data.EnableOwnerRevocation = tfutils.BoolFromString(value.String())
	} else if data.EnableOwnerRevocation.ValueBool() {
		data.EnableOwnerRevocation = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OwnerRevocationEndpoint`); value.Exists() && !data.OwnerRevocationEndpoint.IsNull() {
		data.OwnerRevocationEndpoint = tfutils.ParseStringFromGJSON(value)
	} else if data.OwnerRevocationEndpoint.ValueString() != "/oauth2/issued" {
		data.OwnerRevocationEndpoint = types.StringNull()
	}
	if value := res.Get(pathRoot + `TokenValidationReq`); value.Exists() && !data.TokenValidationReq.IsNull() {
		data.TokenValidationReq = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TokenValidationReq = types.StringNull()
	}
	if value := res.Get(pathRoot + `ThirdPartyAZURL`); value.Exists() && !data.ThirdPartyAzUrl.IsNull() {
		data.ThirdPartyAzUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ThirdPartyAzUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `ThirdPartyTokenURL`); value.Exists() && !data.ThirdPartyTokenUrl.IsNull() {
		data.ThirdPartyTokenUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ThirdPartyTokenUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `ThirdPartyIntrospectURL`); value.Exists() && !data.ThirdPartyIntrospectUrl.IsNull() {
		data.ThirdPartyIntrospectUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ThirdPartyIntrospectUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `ThirdPartyIntrospectCacheType`); value.Exists() && !data.ThirdPartyIntrospectCacheType.IsNull() {
		data.ThirdPartyIntrospectCacheType = tfutils.ParseStringFromGJSON(value)
	} else if data.ThirdPartyIntrospectCacheType.ValueString() != "NoCache" {
		data.ThirdPartyIntrospectCacheType = types.StringNull()
	}
	if value := res.Get(pathRoot + `ThirdPartyIntrospectCacheTimeToLive`); value.Exists() && !data.ThirdPartyIntrospectCacheTimeToLive.IsNull() {
		data.ThirdPartyIntrospectCacheTimeToLive = types.Int64Value(value.Int())
	} else if data.ThirdPartyIntrospectCacheTimeToLive.ValueInt64() != 900 {
		data.ThirdPartyIntrospectCacheTimeToLive = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ThirdPartyAuthorizationHeaderPassThru`); value.Exists() && !data.ThirdPartyAuthorizationHeaderPassThru.IsNull() {
		data.ThirdPartyAuthorizationHeaderPassThru = tfutils.BoolFromString(value.String())
	} else if data.ThirdPartyAuthorizationHeaderPassThru.ValueBool() {
		data.ThirdPartyAuthorizationHeaderPassThru = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ThirdPartyIntrospectURLSecurity`); value.Exists() {
		data.ThirdPartyIntrospectUrlSecurity.UpdateFromBody(ctx, "", value)
	} else {
		data.ThirdPartyIntrospectUrlSecurity = nil
	}
	if value := res.Get(pathRoot + `ThirdPartyIntrospectBasicAuthUserName`); value.Exists() && !data.ThirdPartyIntrospectBasicAuthUserName.IsNull() {
		data.ThirdPartyIntrospectBasicAuthUserName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ThirdPartyIntrospectBasicAuthUserName = types.StringNull()
	}
	if value := res.Get(pathRoot + `ThirdPartyIntrospectBasicAuthPassword`); value.Exists() && !data.ThirdPartyIntrospectBasicAuthPassword.IsNull() {
		data.ThirdPartyIntrospectBasicAuthPassword = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ThirdPartyIntrospectBasicAuthPassword = types.StringNull()
	}
	if value := res.Get(pathRoot + `ThirdPartyBasicAuthHeaderName`); value.Exists() && !data.ThirdPartyBasicAuthHeaderName.IsNull() {
		data.ThirdPartyBasicAuthHeaderName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ThirdPartyBasicAuthHeaderName = types.StringNull()
	}
	if value := res.Get(pathRoot + `ThirdPartyCustomHeaderNameFormat`); value.Exists() && !data.ThirdPartyCustomHeaderNameFormat.IsNull() {
		data.ThirdPartyCustomHeaderNameFormat = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ThirdPartyCustomHeaderNameFormat = types.StringNull()
	}
	if value := res.Get(pathRoot + `ThirdPartyIntrospectSSLProfile`); value.Exists() && !data.ThirdPartyIntrospectSslProfile.IsNull() {
		data.ThirdPartyIntrospectSslProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ThirdPartyIntrospectSslProfile = types.StringNull()
	}
}
