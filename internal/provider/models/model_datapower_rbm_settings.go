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
	"path"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type RBMSettings struct {
	Enabled                           types.Bool                  `tfsdk:"enabled"`
	UserSummary                       types.String                `tfsdk:"user_summary"`
	AuMethod                          types.String                `tfsdk:"au_method"`
	SshAuMethod                       *DmRBMSSHAuthenticateType   `tfsdk:"ssh_au_method"`
	CaPubKeyFile                      types.String                `tfsdk:"ca_pub_key_file"`
	RevokedKeys                       types.List                  `tfsdk:"revoked_keys"`
	AuZosNssConfig                    types.String                `tfsdk:"au_zos_nss_config"`
	AuOidcScope                       types.String                `tfsdk:"au_oidc_scope"`
	AuOidcClientId                    types.String                `tfsdk:"au_oidc_client_id"`
	AuOidcClientSecret                types.String                `tfsdk:"au_oidc_client_secret"`
	AuOidcIdentityServiceUrl          types.String                `tfsdk:"au_oidc_identity_service_url"`
	AuOidcKeyFetchInterval            types.Int64                 `tfsdk:"au_oidc_key_fetch_interval"`
	AuOidcIdentityServiceUrlSslClient types.String                `tfsdk:"au_oidc_identity_service_url_ssl_client"`
	AuKerberosKeytab                  types.String                `tfsdk:"au_kerberos_keytab"`
	AuCustomUrl                       types.String                `tfsdk:"au_custom_url"`
	AuInfoUrl                         types.String                `tfsdk:"au_info_url"`
	AuSslValcred                      types.String                `tfsdk:"au_ssl_valcred"`
	AuHost                            types.String                `tfsdk:"au_host"`
	AuPort                            types.Int64                 `tfsdk:"au_port"`
	AuLdapSearchForDn                 types.Bool                  `tfsdk:"au_ldap_search_for_dn"`
	AuLdapBindDn                      types.String                `tfsdk:"au_ldap_bind_dn"`
	AuLdapBindPasswordAlias           types.String                `tfsdk:"au_ldap_bind_password_alias"`
	AuLdapSearchParameters            types.String                `tfsdk:"au_ldap_search_parameters"`
	AuLdapPrefix                      types.String                `tfsdk:"au_ldap_prefix"`
	AuForceDnLdapOrder                types.Bool                  `tfsdk:"au_force_dn_ldap_order"`
	LdapSuffix                        types.String                `tfsdk:"ldap_suffix"`
	AuLdapLoadBalanceGroup            types.String                `tfsdk:"au_ldap_load_balance_group"`
	AuCacheAllow                      types.String                `tfsdk:"au_cache_allow"`
	AuCacheTtl                        types.Int64                 `tfsdk:"au_cache_ttl"`
	AuLdapReadTimeout                 types.Int64                 `tfsdk:"au_ldap_read_timeout"`
	McMethod                          types.String                `tfsdk:"mc_method"`
	McCustomUrl                       types.String                `tfsdk:"mc_custom_url"`
	McLdapSearchForGroup              types.Bool                  `tfsdk:"mc_ldap_search_for_group"`
	McHost                            types.String                `tfsdk:"mc_host"`
	McPort                            types.Int64                 `tfsdk:"mc_port"`
	McLdapLoadBalanceGroup            types.String                `tfsdk:"mc_ldap_load_balance_group"`
	McLdapBindDn                      types.String                `tfsdk:"mc_ldap_bind_dn"`
	McLdapBindPasswordAlias           types.String                `tfsdk:"mc_ldap_bind_password_alias"`
	McLdapSearchParameters            types.String                `tfsdk:"mc_ldap_search_parameters"`
	McInfoUrl                         types.String                `tfsdk:"mc_info_url"`
	McLdapReadTimeout                 types.Int64                 `tfsdk:"mc_ldap_read_timeout"`
	LdapVersion                       types.String                `tfsdk:"ldap_version"`
	FallbackLogin                     types.String                `tfsdk:"fallback_login"`
	FallbackUser                      types.List                  `tfsdk:"fallback_user"`
	ApplyToCli                        types.Bool                  `tfsdk:"apply_to_cli"`
	RestrictAdminToSerialPort         types.Bool                  `tfsdk:"restrict_admin_to_serial_port"`
	MinPasswordLength                 types.Int64                 `tfsdk:"min_password_length"`
	RequireMixedCase                  types.Bool                  `tfsdk:"require_mixed_case"`
	RequireDigit                      types.Bool                  `tfsdk:"require_digit"`
	RequireNonAlphaNumeric            types.Bool                  `tfsdk:"require_non_alpha_numeric"`
	DisallowUsernameSubstring         types.Bool                  `tfsdk:"disallow_username_substring"`
	DoPasswordAging                   types.Bool                  `tfsdk:"do_password_aging"`
	MaxPasswordAge                    types.Int64                 `tfsdk:"max_password_age"`
	DoPasswordHistory                 types.Bool                  `tfsdk:"do_password_history"`
	NumOldPasswords                   types.Int64                 `tfsdk:"num_old_passwords"`
	CliTimeout                        types.Int64                 `tfsdk:"cli_timeout"`
	MaxFailedLogin                    types.Int64                 `tfsdk:"max_failed_login"`
	LockoutPeriod                     types.Int64                 `tfsdk:"lockout_period"`
	McForceDnLdapOrder                types.Bool                  `tfsdk:"mc_force_dn_ldap_order"`
	PasswordHashAlgorithm             types.String                `tfsdk:"password_hash_algorithm"`
	LdapSslClientConfigType           types.String                `tfsdk:"ldap_ssl_client_config_type"`
	LdapSslClientProfile              types.String                `tfsdk:"ldap_ssl_client_profile"`
	McLdapSslClientConfigType         types.String                `tfsdk:"mc_ldap_ssl_client_config_type"`
	McLdapSslClientProfile            types.String                `tfsdk:"mc_ldap_ssl_client_profile"`
	DependencyActions                 []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var RBMSettingsCAPubKeyFileCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "ssh_au_method",
	AttrType:    "DmRBMSSHAuthenticateType",
	AttrDefault: "",
	Value:       []string{"certificate"},
}
var RBMSettingsAUZOSNSSConfigCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "local",
	Value:       []string{"zosnss"},
}
var RBMSettingsAUKerberosKeytabCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "local",
	Value:       []string{"spnego"},
}
var RBMSettingsAUCustomURLCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "local",
	Value:       []string{"custom"},
}
var RBMSettingsAUInfoURLCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "local",
	Value:       []string{"xmlfile"},
}
var RBMSettingsAUSSLValcredCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "local",
	Value:       []string{"client-ssl"},
}
var RBMSettingsAUHostCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "local",
			Value:       []string{"ldap"},
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
var RBMSettingsAUPortCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "local",
			Value:       []string{"ldap"},
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
var RBMSettingsAULDAPSearchParametersCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "local",
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
var RBMSettingsMCCustomURLCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "mc_method",
	AttrType:    "String",
	AttrDefault: "local",
	Value:       []string{"custom"},
}
var RBMSettingsMCHostCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "mc_ldap_search_for_group",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "mc_method",
			AttrType:    "String",
			AttrDefault: "local",
			Value:       []string{"custom"},
		},
	},
}
var RBMSettingsMCPortCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "mc_ldap_search_for_group",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "mc_method",
			AttrType:    "String",
			AttrDefault: "local",
			Value:       []string{"custom"},
		},
	},
}
var RBMSettingsMCLDAPSearchParametersCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "mc_ldap_search_for_group",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "mc_method",
			AttrType:    "String",
			AttrDefault: "local",
			Value:       []string{"custom"},
		},
	},
}
var RBMSettingsMCInfoURLCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "mc_method",
	AttrType:    "String",
	AttrDefault: "local",
	Value:       []string{"xmlfile"},
}
var RBMSettingsFallbackUserCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation: "logical-not",
			Conditions: []validators.Evaluation{
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "au_method",
					AttrType:    "String",
					AttrDefault: "local",
					Value:       []string{"local"},
				},
			},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "fallback_login",
			AttrType:    "String",
			AttrDefault: "disabled",
			Value:       []string{"restricted"},
		},
	},
}
var RBMSettingsMaxPasswordAgeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "do_password_aging",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var RBMSettingsNumOldPasswordsCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "do_password_history",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var RBMSettingsCAPubKeyFileIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "ssh_au_method",
	AttrType:    "DmRBMSSHAuthenticateType",
	AttrDefault: "",
	Value:       []string{"certificate"},
}
var RBMSettingsRevokedKeysIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "ssh_au_method",
	AttrType:    "DmRBMSSHAuthenticateType",
	AttrDefault: "",
	Value:       []string{"certificate"},
}
var RBMSettingsAUZOSNSSConfigIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var RBMSettingsAUKerberosKeytabIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var RBMSettingsAUCustomURLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var RBMSettingsAUInfoURLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var RBMSettingsAUSSLValcredIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var RBMSettingsAUHostIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var RBMSettingsAUPortIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var RBMSettingsAULDAPSearchForDNIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "local",
	Value:       []string{"ldap"},
}
var RBMSettingsAULDAPBindDNIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "local",
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
var RBMSettingsAULDAPBindPasswordAliasIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "local",
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
var RBMSettingsAULDAPSearchParametersIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "local",
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
var RBMSettingsAULDAPPrefixIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "local",
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
var RBMSettingsAUForceDNLDAPOrderIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "local",
	Value:       []string{"client-ssl"},
}
var RBMSettingsLDAPsuffixIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "local",
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
var RBMSettingsAULDAPLoadBalanceGroupIgnoreVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "local",
			Value:       []string{"ldap"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "local",
			Value:       []string{"custom"},
		},
	},
}
var RBMSettingsAUCacheTTLIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "au_cache_allow",
	AttrType:    "String",
	AttrDefault: "absolute",
	Value:       []string{"disabled"},
}
var RBMSettingsAULDAPReadTimeoutIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "local",
	Value:       []string{"ldap"},
}
var RBMSettingsMCCustomURLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var RBMSettingsMCLDAPSearchForGroupIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "mc_method",
	AttrType:    "String",
	AttrDefault: "local",
	Value:       []string{"custom"},
}
var RBMSettingsMCHostIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var RBMSettingsMCPortIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var RBMSettingsMCLDAPLoadBalanceGroupIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "mc_ldap_search_for_group",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "mc_method",
			AttrType:    "String",
			AttrDefault: "local",
			Value:       []string{"custom"},
		},
	},
}
var RBMSettingsMCLDAPBindDNIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "mc_ldap_search_for_group",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "mc_method",
			AttrType:    "String",
			AttrDefault: "local",
			Value:       []string{"custom"},
		},
	},
}
var RBMSettingsMCLDAPBindPasswordAliasIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "mc_ldap_search_for_group",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "mc_method",
			AttrType:    "String",
			AttrDefault: "local",
			Value:       []string{"custom"},
		},
	},
}
var RBMSettingsMCLDAPSearchParametersIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var RBMSettingsMCInfoURLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var RBMSettingsMCLDAPReadTimeoutIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "mc_ldap_search_for_group",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "mc_method",
			AttrType:    "String",
			AttrDefault: "local",
			Value:       []string{"custom"},
		},
	},
}
var RBMSettingsLDAPVersionIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "local",
	Value:       []string{"ldap"},
}
var RBMSettingsFallbackLoginIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "local",
	Value:       []string{"local"},
}
var RBMSettingsFallbackUserIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var RBMSettingsApplyToCLIIgnoreVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "local",
			Value:       []string{"client-ssl"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "fallback_login",
			AttrType:    "String",
			AttrDefault: "disabled",
			Value:       []string{"disabled"},
		},
	},
}
var RBMSettingsMaxPasswordAgeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var RBMSettingsNumOldPasswordsIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var RBMSettingsMCForceDNLDAPOrderIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "local",
			Value:       []string{"client-ssl"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "mc_ldap_search_for_group",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "mc_method",
			AttrType:    "String",
			AttrDefault: "local",
			Value:       []string{"custom"},
		},
	},
}
var RBMSettingsLDAPSSLClientConfigTypeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "au_method",
	AttrType:    "String",
	AttrDefault: "local",
	Value:       []string{"ldap"},
}
var RBMSettingsLDAPSSLClientProfileIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "au_method",
			AttrType:    "String",
			AttrDefault: "local",
			Value:       []string{"ldap"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "ldap_ssl_client_config_type",
			AttrType:    "String",
			AttrDefault: "client",
			Value:       []string{"client"},
		},
	},
}
var RBMSettingsMCLDAPSSLClientConfigTypeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "mc_ldap_search_for_group",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "mc_method",
			AttrType:    "String",
			AttrDefault: "local",
			Value:       []string{"custom"},
		},
	},
}
var RBMSettingsMCLDAPSSLClientProfileIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "mc_ldap_search_for_group",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "mc_method",
			AttrType:    "String",
			AttrDefault: "local",
			Value:       []string{"custom"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "mc_ldap_ssl_client_config_type",
			AttrType:    "String",
			AttrDefault: "client",
			Value:       []string{"client"},
		},
	},
}

var RBMSettingsObjectType = map[string]attr.Type{
	"enabled":                                 types.BoolType,
	"user_summary":                            types.StringType,
	"au_method":                               types.StringType,
	"ssh_au_method":                           types.ObjectType{AttrTypes: DmRBMSSHAuthenticateTypeObjectType},
	"ca_pub_key_file":                         types.StringType,
	"revoked_keys":                            types.ListType{ElemType: types.StringType},
	"au_zos_nss_config":                       types.StringType,
	"au_oidc_scope":                           types.StringType,
	"au_oidc_client_id":                       types.StringType,
	"au_oidc_client_secret":                   types.StringType,
	"au_oidc_identity_service_url":            types.StringType,
	"au_oidc_key_fetch_interval":              types.Int64Type,
	"au_oidc_identity_service_url_ssl_client": types.StringType,
	"au_kerberos_keytab":                      types.StringType,
	"au_custom_url":                           types.StringType,
	"au_info_url":                             types.StringType,
	"au_ssl_valcred":                          types.StringType,
	"au_host":                                 types.StringType,
	"au_port":                                 types.Int64Type,
	"au_ldap_search_for_dn":                   types.BoolType,
	"au_ldap_bind_dn":                         types.StringType,
	"au_ldap_bind_password_alias":             types.StringType,
	"au_ldap_search_parameters":               types.StringType,
	"au_ldap_prefix":                          types.StringType,
	"au_force_dn_ldap_order":                  types.BoolType,
	"ldap_suffix":                             types.StringType,
	"au_ldap_load_balance_group":              types.StringType,
	"au_cache_allow":                          types.StringType,
	"au_cache_ttl":                            types.Int64Type,
	"au_ldap_read_timeout":                    types.Int64Type,
	"mc_method":                               types.StringType,
	"mc_custom_url":                           types.StringType,
	"mc_ldap_search_for_group":                types.BoolType,
	"mc_host":                                 types.StringType,
	"mc_port":                                 types.Int64Type,
	"mc_ldap_load_balance_group":              types.StringType,
	"mc_ldap_bind_dn":                         types.StringType,
	"mc_ldap_bind_password_alias":             types.StringType,
	"mc_ldap_search_parameters":               types.StringType,
	"mc_info_url":                             types.StringType,
	"mc_ldap_read_timeout":                    types.Int64Type,
	"ldap_version":                            types.StringType,
	"fallback_login":                          types.StringType,
	"fallback_user":                           types.ListType{ElemType: types.StringType},
	"apply_to_cli":                            types.BoolType,
	"restrict_admin_to_serial_port":           types.BoolType,
	"min_password_length":                     types.Int64Type,
	"require_mixed_case":                      types.BoolType,
	"require_digit":                           types.BoolType,
	"require_non_alpha_numeric":               types.BoolType,
	"disallow_username_substring":             types.BoolType,
	"do_password_aging":                       types.BoolType,
	"max_password_age":                        types.Int64Type,
	"do_password_history":                     types.BoolType,
	"num_old_passwords":                       types.Int64Type,
	"cli_timeout":                             types.Int64Type,
	"max_failed_login":                        types.Int64Type,
	"lockout_period":                          types.Int64Type,
	"mc_force_dn_ldap_order":                  types.BoolType,
	"password_hash_algorithm":                 types.StringType,
	"ldap_ssl_client_config_type":             types.StringType,
	"ldap_ssl_client_profile":                 types.StringType,
	"mc_ldap_ssl_client_config_type":          types.StringType,
	"mc_ldap_ssl_client_profile":              types.StringType,
	"dependency_actions":                      actions.ActionsListType,
}

func (data RBMSettings) GetPath() string {
	rest_path := "/mgmt/config/default/RBMSettings/RBM-Settings"
	return rest_path
}

func (data RBMSettings) IsNull() bool {
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.AuMethod.IsNull() {
		return false
	}
	if data.SshAuMethod != nil {
		if !data.SshAuMethod.IsNull() {
			return false
		}
	}
	if !data.CaPubKeyFile.IsNull() {
		return false
	}
	if !data.RevokedKeys.IsNull() {
		return false
	}
	if !data.AuZosNssConfig.IsNull() {
		return false
	}
	if !data.AuOidcScope.IsNull() {
		return false
	}
	if !data.AuOidcClientId.IsNull() {
		return false
	}
	if !data.AuOidcClientSecret.IsNull() {
		return false
	}
	if !data.AuOidcIdentityServiceUrl.IsNull() {
		return false
	}
	if !data.AuOidcKeyFetchInterval.IsNull() {
		return false
	}
	if !data.AuOidcIdentityServiceUrlSslClient.IsNull() {
		return false
	}
	if !data.AuKerberosKeytab.IsNull() {
		return false
	}
	if !data.AuCustomUrl.IsNull() {
		return false
	}
	if !data.AuInfoUrl.IsNull() {
		return false
	}
	if !data.AuSslValcred.IsNull() {
		return false
	}
	if !data.AuHost.IsNull() {
		return false
	}
	if !data.AuPort.IsNull() {
		return false
	}
	if !data.AuLdapSearchForDn.IsNull() {
		return false
	}
	if !data.AuLdapBindDn.IsNull() {
		return false
	}
	if !data.AuLdapBindPasswordAlias.IsNull() {
		return false
	}
	if !data.AuLdapSearchParameters.IsNull() {
		return false
	}
	if !data.AuLdapPrefix.IsNull() {
		return false
	}
	if !data.AuForceDnLdapOrder.IsNull() {
		return false
	}
	if !data.LdapSuffix.IsNull() {
		return false
	}
	if !data.AuLdapLoadBalanceGroup.IsNull() {
		return false
	}
	if !data.AuCacheAllow.IsNull() {
		return false
	}
	if !data.AuCacheTtl.IsNull() {
		return false
	}
	if !data.AuLdapReadTimeout.IsNull() {
		return false
	}
	if !data.McMethod.IsNull() {
		return false
	}
	if !data.McCustomUrl.IsNull() {
		return false
	}
	if !data.McLdapSearchForGroup.IsNull() {
		return false
	}
	if !data.McHost.IsNull() {
		return false
	}
	if !data.McPort.IsNull() {
		return false
	}
	if !data.McLdapLoadBalanceGroup.IsNull() {
		return false
	}
	if !data.McLdapBindDn.IsNull() {
		return false
	}
	if !data.McLdapBindPasswordAlias.IsNull() {
		return false
	}
	if !data.McLdapSearchParameters.IsNull() {
		return false
	}
	if !data.McInfoUrl.IsNull() {
		return false
	}
	if !data.McLdapReadTimeout.IsNull() {
		return false
	}
	if !data.LdapVersion.IsNull() {
		return false
	}
	if !data.FallbackLogin.IsNull() {
		return false
	}
	if !data.FallbackUser.IsNull() {
		return false
	}
	if !data.ApplyToCli.IsNull() {
		return false
	}
	if !data.RestrictAdminToSerialPort.IsNull() {
		return false
	}
	if !data.MinPasswordLength.IsNull() {
		return false
	}
	if !data.RequireMixedCase.IsNull() {
		return false
	}
	if !data.RequireDigit.IsNull() {
		return false
	}
	if !data.RequireNonAlphaNumeric.IsNull() {
		return false
	}
	if !data.DisallowUsernameSubstring.IsNull() {
		return false
	}
	if !data.DoPasswordAging.IsNull() {
		return false
	}
	if !data.MaxPasswordAge.IsNull() {
		return false
	}
	if !data.DoPasswordHistory.IsNull() {
		return false
	}
	if !data.NumOldPasswords.IsNull() {
		return false
	}
	if !data.CliTimeout.IsNull() {
		return false
	}
	if !data.MaxFailedLogin.IsNull() {
		return false
	}
	if !data.LockoutPeriod.IsNull() {
		return false
	}
	if !data.McForceDnLdapOrder.IsNull() {
		return false
	}
	if !data.PasswordHashAlgorithm.IsNull() {
		return false
	}
	if !data.LdapSslClientConfigType.IsNull() {
		return false
	}
	if !data.LdapSslClientProfile.IsNull() {
		return false
	}
	if !data.McLdapSslClientConfigType.IsNull() {
		return false
	}
	if !data.McLdapSslClientProfile.IsNull() {
		return false
	}
	return true
}
func (data *RBMSettings) ToDefault() {
	data.Enabled = types.BoolValue(true)
	data.UserSummary = types.StringNull()
	data.AuMethod = types.StringValue("local")
	data.SshAuMethod = &DmRBMSSHAuthenticateType{}
	data.SshAuMethod.ToDefault()
	data.CaPubKeyFile = types.StringNull()
	data.RevokedKeys = types.ListNull(types.StringType)
	data.AuZosNssConfig = types.StringNull()
	data.AuOidcScope = types.StringValue("openid")
	data.AuOidcClientId = types.StringNull()
	data.AuOidcClientSecret = types.StringNull()
	data.AuOidcIdentityServiceUrl = types.StringNull()
	data.AuOidcKeyFetchInterval = types.Int64Value(30)
	data.AuOidcIdentityServiceUrlSslClient = types.StringNull()
	data.AuKerberosKeytab = types.StringNull()
	data.AuCustomUrl = types.StringNull()
	data.AuInfoUrl = types.StringNull()
	data.AuSslValcred = types.StringNull()
	data.AuHost = types.StringNull()
	data.AuPort = types.Int64Null()
	data.AuLdapSearchForDn = types.BoolValue(false)
	data.AuLdapBindDn = types.StringNull()
	data.AuLdapBindPasswordAlias = types.StringNull()
	data.AuLdapSearchParameters = types.StringNull()
	data.AuLdapPrefix = types.StringValue("cn=")
	data.AuForceDnLdapOrder = types.BoolValue(false)
	data.LdapSuffix = types.StringNull()
	data.AuLdapLoadBalanceGroup = types.StringNull()
	data.AuCacheAllow = types.StringValue("absolute")
	data.AuCacheTtl = types.Int64Value(600)
	data.AuLdapReadTimeout = types.Int64Value(60)
	data.McMethod = types.StringValue("local")
	data.McCustomUrl = types.StringNull()
	data.McLdapSearchForGroup = types.BoolValue(false)
	data.McHost = types.StringNull()
	data.McPort = types.Int64Null()
	data.McLdapLoadBalanceGroup = types.StringNull()
	data.McLdapBindDn = types.StringNull()
	data.McLdapBindPasswordAlias = types.StringNull()
	data.McLdapSearchParameters = types.StringNull()
	data.McInfoUrl = types.StringNull()
	data.McLdapReadTimeout = types.Int64Value(60)
	data.LdapVersion = types.StringNull()
	data.FallbackLogin = types.StringValue("disabled")
	data.FallbackUser = types.ListNull(types.StringType)
	data.ApplyToCli = types.BoolValue(false)
	data.RestrictAdminToSerialPort = types.BoolValue(false)
	data.MinPasswordLength = types.Int64Value(6)
	data.RequireMixedCase = types.BoolValue(false)
	data.RequireDigit = types.BoolValue(false)
	data.RequireNonAlphaNumeric = types.BoolValue(false)
	data.DisallowUsernameSubstring = types.BoolValue(false)
	data.DoPasswordAging = types.BoolValue(false)
	data.MaxPasswordAge = types.Int64Value(30)
	data.DoPasswordHistory = types.BoolValue(false)
	data.NumOldPasswords = types.Int64Value(5)
	data.CliTimeout = types.Int64Value(0)
	data.MaxFailedLogin = types.Int64Value(0)
	data.LockoutPeriod = types.Int64Value(1)
	data.McForceDnLdapOrder = types.BoolValue(false)
	data.PasswordHashAlgorithm = types.StringValue("md5crypt")
	data.LdapSslClientConfigType = types.StringValue("client")
	data.LdapSslClientProfile = types.StringNull()
	data.McLdapSslClientConfigType = types.StringValue("client")
	data.McLdapSslClientProfile = types.StringNull()
}

func (data RBMSettings) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "RBMSettings.name", path.Base("/mgmt/config/default/RBMSettings/RBM-Settings"))

	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.AuMethod.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUMethod`, data.AuMethod.ValueString())
	}
	if data.SshAuMethod != nil {
		if !data.SshAuMethod.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`SSHAUMethod`, data.SshAuMethod.ToBody(ctx, ""))
		}
	}
	if !data.CaPubKeyFile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CAPubKeyFile`, data.CaPubKeyFile.ValueString())
	}
	if !data.RevokedKeys.IsNull() {
		var dataValues []string
		data.RevokedKeys.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`RevokedKeys`+".-1", map[string]string{"value": val})
		}
	}
	if !data.AuZosNssConfig.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUZOSNSSConfig`, data.AuZosNssConfig.ValueString())
	}
	if !data.AuOidcScope.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUOIDCScope`, data.AuOidcScope.ValueString())
	}
	if !data.AuOidcClientId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUOIDCClientID`, data.AuOidcClientId.ValueString())
	}
	if !data.AuOidcClientSecret.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUOIDCClientSecret`, data.AuOidcClientSecret.ValueString())
	}
	if !data.AuOidcIdentityServiceUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUOIDCIdentityServiceURL`, data.AuOidcIdentityServiceUrl.ValueString())
	}
	if !data.AuOidcKeyFetchInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUOIDCKeyFetchInterval`, data.AuOidcKeyFetchInterval.ValueInt64())
	}
	if !data.AuOidcIdentityServiceUrlSslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUOIDCIdentityServiceURLSSLClient`, data.AuOidcIdentityServiceUrlSslClient.ValueString())
	}
	if !data.AuKerberosKeytab.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUKerberosKeytab`, data.AuKerberosKeytab.ValueString())
	}
	if !data.AuCustomUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUCustomURL`, data.AuCustomUrl.ValueString())
	}
	if !data.AuInfoUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUInfoURL`, data.AuInfoUrl.ValueString())
	}
	if !data.AuSslValcred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUSSLValcred`, data.AuSslValcred.ValueString())
	}
	if !data.AuHost.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUHost`, data.AuHost.ValueString())
	}
	if !data.AuPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUPort`, data.AuPort.ValueInt64())
	}
	if !data.AuLdapSearchForDn.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPSearchForDN`, tfutils.StringFromBool(data.AuLdapSearchForDn, ""))
	}
	if !data.AuLdapBindDn.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPBindDN`, data.AuLdapBindDn.ValueString())
	}
	if !data.AuLdapBindPasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPBindPasswordAlias`, data.AuLdapBindPasswordAlias.ValueString())
	}
	if !data.AuLdapSearchParameters.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPSearchParameters`, data.AuLdapSearchParameters.ValueString())
	}
	if !data.AuLdapPrefix.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPPrefix`, data.AuLdapPrefix.ValueString())
	}
	if !data.AuForceDnLdapOrder.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUForceDNLDAPOrder`, tfutils.StringFromBool(data.AuForceDnLdapOrder, ""))
	}
	if !data.LdapSuffix.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPsuffix`, data.LdapSuffix.ValueString())
	}
	if !data.AuLdapLoadBalanceGroup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPLoadBalanceGroup`, data.AuLdapLoadBalanceGroup.ValueString())
	}
	if !data.AuCacheAllow.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUCacheAllow`, data.AuCacheAllow.ValueString())
	}
	if !data.AuCacheTtl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUCacheTTL`, data.AuCacheTtl.ValueInt64())
	}
	if !data.AuLdapReadTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPReadTimeout`, data.AuLdapReadTimeout.ValueInt64())
	}
	if !data.McMethod.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCMethod`, data.McMethod.ValueString())
	}
	if !data.McCustomUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCCustomURL`, data.McCustomUrl.ValueString())
	}
	if !data.McLdapSearchForGroup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCLDAPSearchForGroup`, tfutils.StringFromBool(data.McLdapSearchForGroup, ""))
	}
	if !data.McHost.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCHost`, data.McHost.ValueString())
	}
	if !data.McPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCPort`, data.McPort.ValueInt64())
	}
	if !data.McLdapLoadBalanceGroup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCLDAPLoadBalanceGroup`, data.McLdapLoadBalanceGroup.ValueString())
	}
	if !data.McLdapBindDn.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCLDAPBindDN`, data.McLdapBindDn.ValueString())
	}
	if !data.McLdapBindPasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCLDAPBindPasswordAlias`, data.McLdapBindPasswordAlias.ValueString())
	}
	if !data.McLdapSearchParameters.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCLDAPSearchParameters`, data.McLdapSearchParameters.ValueString())
	}
	if !data.McInfoUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCInfoURL`, data.McInfoUrl.ValueString())
	}
	if !data.McLdapReadTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCLDAPReadTimeout`, data.McLdapReadTimeout.ValueInt64())
	}
	if !data.LdapVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPVersion`, data.LdapVersion.ValueString())
	}
	if !data.FallbackLogin.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FallbackLogin`, data.FallbackLogin.ValueString())
	}
	if !data.FallbackUser.IsNull() {
		var dataValues []string
		data.FallbackUser.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`FallbackUser`+".-1", map[string]string{"value": val})
		}
	}
	if !data.ApplyToCli.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ApplyToCLI`, tfutils.StringFromBool(data.ApplyToCli, ""))
	}
	if !data.RestrictAdminToSerialPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RestrictAdminToSerialPort`, tfutils.StringFromBool(data.RestrictAdminToSerialPort, ""))
	}
	if !data.MinPasswordLength.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MinPasswordLength`, data.MinPasswordLength.ValueInt64())
	}
	if !data.RequireMixedCase.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RequireMixedCase`, tfutils.StringFromBool(data.RequireMixedCase, ""))
	}
	if !data.RequireDigit.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RequireDigit`, tfutils.StringFromBool(data.RequireDigit, ""))
	}
	if !data.RequireNonAlphaNumeric.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RequireNonAlphaNumeric`, tfutils.StringFromBool(data.RequireNonAlphaNumeric, ""))
	}
	if !data.DisallowUsernameSubstring.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DisallowUsernameSubstring`, tfutils.StringFromBool(data.DisallowUsernameSubstring, ""))
	}
	if !data.DoPasswordAging.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DoPasswordAging`, tfutils.StringFromBool(data.DoPasswordAging, ""))
	}
	if !data.MaxPasswordAge.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxPasswordAge`, data.MaxPasswordAge.ValueInt64())
	}
	if !data.DoPasswordHistory.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DoPasswordHistory`, tfutils.StringFromBool(data.DoPasswordHistory, ""))
	}
	if !data.NumOldPasswords.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`NumOldPasswords`, data.NumOldPasswords.ValueInt64())
	}
	if !data.CliTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CliTimeout`, data.CliTimeout.ValueInt64())
	}
	if !data.MaxFailedLogin.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxFailedLogin`, data.MaxFailedLogin.ValueInt64())
	}
	if !data.LockoutPeriod.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LockoutPeriod`, data.LockoutPeriod.ValueInt64())
	}
	if !data.McForceDnLdapOrder.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCForceDNLDAPOrder`, tfutils.StringFromBool(data.McForceDnLdapOrder, ""))
	}
	if !data.PasswordHashAlgorithm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PasswordHashAlgorithm`, data.PasswordHashAlgorithm.ValueString())
	}
	if !data.LdapSslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPSSLClientConfigType`, data.LdapSslClientConfigType.ValueString())
	}
	if !data.LdapSslClientProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPSSLClientProfile`, data.LdapSslClientProfile.ValueString())
	}
	if !data.McLdapSslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCLDAPSSLClientConfigType`, data.McLdapSslClientConfigType.ValueString())
	}
	if !data.McLdapSslClientProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCLDAPSSLClientProfile`, data.McLdapSslClientProfile.ValueString())
	}
	return body
}

func (data *RBMSettings) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUMethod`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuMethod = types.StringValue("local")
	}
	if value := res.Get(pathRoot + `SSHAUMethod`); value.Exists() {
		data.SshAuMethod = &DmRBMSSHAuthenticateType{}
		data.SshAuMethod.FromBody(ctx, "", value)
	} else {
		data.SshAuMethod = nil
	}
	if value := res.Get(pathRoot + `CAPubKeyFile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CaPubKeyFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CaPubKeyFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `RevokedKeys`); value.Exists() {
		data.RevokedKeys = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.RevokedKeys = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `AUZOSNSSConfig`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuZosNssConfig = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuZosNssConfig = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUOIDCScope`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuOidcScope = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuOidcScope = types.StringValue("openid")
	}
	if value := res.Get(pathRoot + `AUOIDCClientID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuOidcClientId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuOidcClientId = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUOIDCClientSecret`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuOidcClientSecret = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuOidcClientSecret = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUOIDCIdentityServiceURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuOidcIdentityServiceUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuOidcIdentityServiceUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUOIDCKeyFetchInterval`); value.Exists() {
		data.AuOidcKeyFetchInterval = types.Int64Value(value.Int())
	} else {
		data.AuOidcKeyFetchInterval = types.Int64Value(30)
	}
	if value := res.Get(pathRoot + `AUOIDCIdentityServiceURLSSLClient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuOidcIdentityServiceUrlSslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuOidcIdentityServiceUrlSslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUKerberosKeytab`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuKerberosKeytab = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuKerberosKeytab = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUCustomURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuCustomUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuCustomUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUInfoURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuInfoUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuInfoUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSSLValcred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuSslValcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuSslValcred = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUHost`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuHost = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuHost = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUPort`); value.Exists() {
		data.AuPort = types.Int64Value(value.Int())
	} else {
		data.AuPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AULDAPSearchForDN`); value.Exists() {
		data.AuLdapSearchForDn = tfutils.BoolFromString(value.String())
	} else {
		data.AuLdapSearchForDn = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AULDAPBindDN`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuLdapBindDn = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuLdapBindDn = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPBindPasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuLdapBindPasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuLdapBindPasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPSearchParameters`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuLdapSearchParameters = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuLdapSearchParameters = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPPrefix`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuLdapPrefix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuLdapPrefix = types.StringValue("cn=")
	}
	if value := res.Get(pathRoot + `AUForceDNLDAPOrder`); value.Exists() {
		data.AuForceDnLdapOrder = tfutils.BoolFromString(value.String())
	} else {
		data.AuForceDnLdapOrder = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LDAPsuffix`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapSuffix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapSuffix = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPLoadBalanceGroup`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuLdapLoadBalanceGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuLdapLoadBalanceGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUCacheAllow`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuCacheAllow = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuCacheAllow = types.StringValue("absolute")
	}
	if value := res.Get(pathRoot + `AUCacheTTL`); value.Exists() {
		data.AuCacheTtl = types.Int64Value(value.Int())
	} else {
		data.AuCacheTtl = types.Int64Value(600)
	}
	if value := res.Get(pathRoot + `AULDAPReadTimeout`); value.Exists() {
		data.AuLdapReadTimeout = types.Int64Value(value.Int())
	} else {
		data.AuLdapReadTimeout = types.Int64Value(60)
	}
	if value := res.Get(pathRoot + `MCMethod`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.McMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McMethod = types.StringValue("local")
	}
	if value := res.Get(pathRoot + `MCCustomURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.McCustomUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McCustomUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCLDAPSearchForGroup`); value.Exists() {
		data.McLdapSearchForGroup = tfutils.BoolFromString(value.String())
	} else {
		data.McLdapSearchForGroup = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MCHost`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.McHost = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McHost = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCPort`); value.Exists() {
		data.McPort = types.Int64Value(value.Int())
	} else {
		data.McPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MCLDAPLoadBalanceGroup`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.McLdapLoadBalanceGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McLdapLoadBalanceGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCLDAPBindDN`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.McLdapBindDn = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McLdapBindDn = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCLDAPBindPasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.McLdapBindPasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McLdapBindPasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCLDAPSearchParameters`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.McLdapSearchParameters = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McLdapSearchParameters = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCInfoURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.McInfoUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McInfoUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCLDAPReadTimeout`); value.Exists() {
		data.McLdapReadTimeout = types.Int64Value(value.Int())
	} else {
		data.McLdapReadTimeout = types.Int64Value(60)
	}
	if value := res.Get(pathRoot + `LDAPVersion`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `FallbackLogin`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.FallbackLogin = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FallbackLogin = types.StringValue("disabled")
	}
	if value := res.Get(pathRoot + `FallbackUser`); value.Exists() {
		data.FallbackUser = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.FallbackUser = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `ApplyToCLI`); value.Exists() {
		data.ApplyToCli = tfutils.BoolFromString(value.String())
	} else {
		data.ApplyToCli = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RestrictAdminToSerialPort`); value.Exists() {
		data.RestrictAdminToSerialPort = tfutils.BoolFromString(value.String())
	} else {
		data.RestrictAdminToSerialPort = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MinPasswordLength`); value.Exists() {
		data.MinPasswordLength = types.Int64Value(value.Int())
	} else {
		data.MinPasswordLength = types.Int64Value(6)
	}
	if value := res.Get(pathRoot + `RequireMixedCase`); value.Exists() {
		data.RequireMixedCase = tfutils.BoolFromString(value.String())
	} else {
		data.RequireMixedCase = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RequireDigit`); value.Exists() {
		data.RequireDigit = tfutils.BoolFromString(value.String())
	} else {
		data.RequireDigit = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RequireNonAlphaNumeric`); value.Exists() {
		data.RequireNonAlphaNumeric = tfutils.BoolFromString(value.String())
	} else {
		data.RequireNonAlphaNumeric = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DisallowUsernameSubstring`); value.Exists() {
		data.DisallowUsernameSubstring = tfutils.BoolFromString(value.String())
	} else {
		data.DisallowUsernameSubstring = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DoPasswordAging`); value.Exists() {
		data.DoPasswordAging = tfutils.BoolFromString(value.String())
	} else {
		data.DoPasswordAging = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MaxPasswordAge`); value.Exists() {
		data.MaxPasswordAge = types.Int64Value(value.Int())
	} else {
		data.MaxPasswordAge = types.Int64Value(30)
	}
	if value := res.Get(pathRoot + `DoPasswordHistory`); value.Exists() {
		data.DoPasswordHistory = tfutils.BoolFromString(value.String())
	} else {
		data.DoPasswordHistory = types.BoolNull()
	}
	if value := res.Get(pathRoot + `NumOldPasswords`); value.Exists() {
		data.NumOldPasswords = types.Int64Value(value.Int())
	} else {
		data.NumOldPasswords = types.Int64Value(5)
	}
	if value := res.Get(pathRoot + `CliTimeout`); value.Exists() {
		data.CliTimeout = types.Int64Value(value.Int())
	} else {
		data.CliTimeout = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `MaxFailedLogin`); value.Exists() {
		data.MaxFailedLogin = types.Int64Value(value.Int())
	} else {
		data.MaxFailedLogin = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `LockoutPeriod`); value.Exists() {
		data.LockoutPeriod = types.Int64Value(value.Int())
	} else {
		data.LockoutPeriod = types.Int64Value(1)
	}
	if value := res.Get(pathRoot + `MCForceDNLDAPOrder`); value.Exists() {
		data.McForceDnLdapOrder = tfutils.BoolFromString(value.String())
	} else {
		data.McForceDnLdapOrder = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PasswordHashAlgorithm`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PasswordHashAlgorithm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordHashAlgorithm = types.StringValue("md5crypt")
	}
	if value := res.Get(pathRoot + `LDAPSSLClientConfigType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapSslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapSslClientConfigType = types.StringValue("client")
	}
	if value := res.Get(pathRoot + `LDAPSSLClientProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapSslClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapSslClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCLDAPSSLClientConfigType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.McLdapSslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McLdapSslClientConfigType = types.StringValue("client")
	}
	if value := res.Get(pathRoot + `MCLDAPSSLClientProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.McLdapSslClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McLdapSslClientProfile = types.StringNull()
	}
}

func (data *RBMSettings) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else if !data.Enabled.ValueBool() {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUMethod`); value.Exists() && !data.AuMethod.IsNull() {
		data.AuMethod = tfutils.ParseStringFromGJSON(value)
	} else if data.AuMethod.ValueString() != "local" {
		data.AuMethod = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSHAUMethod`); value.Exists() {
		data.SshAuMethod.UpdateFromBody(ctx, "", value)
	} else {
		data.SshAuMethod = nil
	}
	if value := res.Get(pathRoot + `CAPubKeyFile`); value.Exists() && !data.CaPubKeyFile.IsNull() {
		data.CaPubKeyFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CaPubKeyFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `RevokedKeys`); value.Exists() && !data.RevokedKeys.IsNull() {
		data.RevokedKeys = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.RevokedKeys = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `AUZOSNSSConfig`); value.Exists() && !data.AuZosNssConfig.IsNull() {
		data.AuZosNssConfig = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuZosNssConfig = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUOIDCScope`); value.Exists() && !data.AuOidcScope.IsNull() {
		data.AuOidcScope = tfutils.ParseStringFromGJSON(value)
	} else if data.AuOidcScope.ValueString() != "openid" {
		data.AuOidcScope = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUOIDCClientID`); value.Exists() && !data.AuOidcClientId.IsNull() {
		data.AuOidcClientId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuOidcClientId = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUOIDCClientSecret`); value.Exists() && !data.AuOidcClientSecret.IsNull() {
		data.AuOidcClientSecret = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuOidcClientSecret = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUOIDCIdentityServiceURL`); value.Exists() && !data.AuOidcIdentityServiceUrl.IsNull() {
		data.AuOidcIdentityServiceUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuOidcIdentityServiceUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUOIDCKeyFetchInterval`); value.Exists() && !data.AuOidcKeyFetchInterval.IsNull() {
		data.AuOidcKeyFetchInterval = types.Int64Value(value.Int())
	} else if data.AuOidcKeyFetchInterval.ValueInt64() != 30 {
		data.AuOidcKeyFetchInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AUOIDCIdentityServiceURLSSLClient`); value.Exists() && !data.AuOidcIdentityServiceUrlSslClient.IsNull() {
		data.AuOidcIdentityServiceUrlSslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuOidcIdentityServiceUrlSslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUKerberosKeytab`); value.Exists() && !data.AuKerberosKeytab.IsNull() {
		data.AuKerberosKeytab = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuKerberosKeytab = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUCustomURL`); value.Exists() && !data.AuCustomUrl.IsNull() {
		data.AuCustomUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuCustomUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUInfoURL`); value.Exists() && !data.AuInfoUrl.IsNull() {
		data.AuInfoUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuInfoUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSSLValcred`); value.Exists() && !data.AuSslValcred.IsNull() {
		data.AuSslValcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuSslValcred = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUHost`); value.Exists() && !data.AuHost.IsNull() {
		data.AuHost = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuHost = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUPort`); value.Exists() && !data.AuPort.IsNull() {
		data.AuPort = types.Int64Value(value.Int())
	} else {
		data.AuPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AULDAPSearchForDN`); value.Exists() && !data.AuLdapSearchForDn.IsNull() {
		data.AuLdapSearchForDn = tfutils.BoolFromString(value.String())
	} else if data.AuLdapSearchForDn.ValueBool() {
		data.AuLdapSearchForDn = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AULDAPBindDN`); value.Exists() && !data.AuLdapBindDn.IsNull() {
		data.AuLdapBindDn = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuLdapBindDn = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPBindPasswordAlias`); value.Exists() && !data.AuLdapBindPasswordAlias.IsNull() {
		data.AuLdapBindPasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuLdapBindPasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPSearchParameters`); value.Exists() && !data.AuLdapSearchParameters.IsNull() {
		data.AuLdapSearchParameters = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuLdapSearchParameters = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPPrefix`); value.Exists() && !data.AuLdapPrefix.IsNull() {
		data.AuLdapPrefix = tfutils.ParseStringFromGJSON(value)
	} else if data.AuLdapPrefix.ValueString() != "cn=" {
		data.AuLdapPrefix = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUForceDNLDAPOrder`); value.Exists() && !data.AuForceDnLdapOrder.IsNull() {
		data.AuForceDnLdapOrder = tfutils.BoolFromString(value.String())
	} else if data.AuForceDnLdapOrder.ValueBool() {
		data.AuForceDnLdapOrder = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LDAPsuffix`); value.Exists() && !data.LdapSuffix.IsNull() {
		data.LdapSuffix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapSuffix = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPLoadBalanceGroup`); value.Exists() && !data.AuLdapLoadBalanceGroup.IsNull() {
		data.AuLdapLoadBalanceGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuLdapLoadBalanceGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUCacheAllow`); value.Exists() && !data.AuCacheAllow.IsNull() {
		data.AuCacheAllow = tfutils.ParseStringFromGJSON(value)
	} else if data.AuCacheAllow.ValueString() != "absolute" {
		data.AuCacheAllow = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUCacheTTL`); value.Exists() && !data.AuCacheTtl.IsNull() {
		data.AuCacheTtl = types.Int64Value(value.Int())
	} else if data.AuCacheTtl.ValueInt64() != 600 {
		data.AuCacheTtl = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AULDAPReadTimeout`); value.Exists() && !data.AuLdapReadTimeout.IsNull() {
		data.AuLdapReadTimeout = types.Int64Value(value.Int())
	} else if data.AuLdapReadTimeout.ValueInt64() != 60 {
		data.AuLdapReadTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MCMethod`); value.Exists() && !data.McMethod.IsNull() {
		data.McMethod = tfutils.ParseStringFromGJSON(value)
	} else if data.McMethod.ValueString() != "local" {
		data.McMethod = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCCustomURL`); value.Exists() && !data.McCustomUrl.IsNull() {
		data.McCustomUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McCustomUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCLDAPSearchForGroup`); value.Exists() && !data.McLdapSearchForGroup.IsNull() {
		data.McLdapSearchForGroup = tfutils.BoolFromString(value.String())
	} else if data.McLdapSearchForGroup.ValueBool() {
		data.McLdapSearchForGroup = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MCHost`); value.Exists() && !data.McHost.IsNull() {
		data.McHost = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McHost = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCPort`); value.Exists() && !data.McPort.IsNull() {
		data.McPort = types.Int64Value(value.Int())
	} else {
		data.McPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MCLDAPLoadBalanceGroup`); value.Exists() && !data.McLdapLoadBalanceGroup.IsNull() {
		data.McLdapLoadBalanceGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McLdapLoadBalanceGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCLDAPBindDN`); value.Exists() && !data.McLdapBindDn.IsNull() {
		data.McLdapBindDn = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McLdapBindDn = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCLDAPBindPasswordAlias`); value.Exists() && !data.McLdapBindPasswordAlias.IsNull() {
		data.McLdapBindPasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McLdapBindPasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCLDAPSearchParameters`); value.Exists() && !data.McLdapSearchParameters.IsNull() {
		data.McLdapSearchParameters = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McLdapSearchParameters = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCInfoURL`); value.Exists() && !data.McInfoUrl.IsNull() {
		data.McInfoUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McInfoUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCLDAPReadTimeout`); value.Exists() && !data.McLdapReadTimeout.IsNull() {
		data.McLdapReadTimeout = types.Int64Value(value.Int())
	} else if data.McLdapReadTimeout.ValueInt64() != 60 {
		data.McLdapReadTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `LDAPVersion`); value.Exists() && !data.LdapVersion.IsNull() {
		data.LdapVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `FallbackLogin`); value.Exists() && !data.FallbackLogin.IsNull() {
		data.FallbackLogin = tfutils.ParseStringFromGJSON(value)
	} else if data.FallbackLogin.ValueString() != "disabled" {
		data.FallbackLogin = types.StringNull()
	}
	if value := res.Get(pathRoot + `FallbackUser`); value.Exists() && !data.FallbackUser.IsNull() {
		data.FallbackUser = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.FallbackUser = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `ApplyToCLI`); value.Exists() && !data.ApplyToCli.IsNull() {
		data.ApplyToCli = tfutils.BoolFromString(value.String())
	} else if data.ApplyToCli.ValueBool() {
		data.ApplyToCli = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RestrictAdminToSerialPort`); value.Exists() && !data.RestrictAdminToSerialPort.IsNull() {
		data.RestrictAdminToSerialPort = tfutils.BoolFromString(value.String())
	} else if data.RestrictAdminToSerialPort.ValueBool() {
		data.RestrictAdminToSerialPort = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MinPasswordLength`); value.Exists() && !data.MinPasswordLength.IsNull() {
		data.MinPasswordLength = types.Int64Value(value.Int())
	} else if data.MinPasswordLength.ValueInt64() != 6 {
		data.MinPasswordLength = types.Int64Null()
	}
	if value := res.Get(pathRoot + `RequireMixedCase`); value.Exists() && !data.RequireMixedCase.IsNull() {
		data.RequireMixedCase = tfutils.BoolFromString(value.String())
	} else if data.RequireMixedCase.ValueBool() {
		data.RequireMixedCase = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RequireDigit`); value.Exists() && !data.RequireDigit.IsNull() {
		data.RequireDigit = tfutils.BoolFromString(value.String())
	} else if data.RequireDigit.ValueBool() {
		data.RequireDigit = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RequireNonAlphaNumeric`); value.Exists() && !data.RequireNonAlphaNumeric.IsNull() {
		data.RequireNonAlphaNumeric = tfutils.BoolFromString(value.String())
	} else if data.RequireNonAlphaNumeric.ValueBool() {
		data.RequireNonAlphaNumeric = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DisallowUsernameSubstring`); value.Exists() && !data.DisallowUsernameSubstring.IsNull() {
		data.DisallowUsernameSubstring = tfutils.BoolFromString(value.String())
	} else if data.DisallowUsernameSubstring.ValueBool() {
		data.DisallowUsernameSubstring = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DoPasswordAging`); value.Exists() && !data.DoPasswordAging.IsNull() {
		data.DoPasswordAging = tfutils.BoolFromString(value.String())
	} else if data.DoPasswordAging.ValueBool() {
		data.DoPasswordAging = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MaxPasswordAge`); value.Exists() && !data.MaxPasswordAge.IsNull() {
		data.MaxPasswordAge = types.Int64Value(value.Int())
	} else if data.MaxPasswordAge.ValueInt64() != 30 {
		data.MaxPasswordAge = types.Int64Null()
	}
	if value := res.Get(pathRoot + `DoPasswordHistory`); value.Exists() && !data.DoPasswordHistory.IsNull() {
		data.DoPasswordHistory = tfutils.BoolFromString(value.String())
	} else if data.DoPasswordHistory.ValueBool() {
		data.DoPasswordHistory = types.BoolNull()
	}
	if value := res.Get(pathRoot + `NumOldPasswords`); value.Exists() && !data.NumOldPasswords.IsNull() {
		data.NumOldPasswords = types.Int64Value(value.Int())
	} else if data.NumOldPasswords.ValueInt64() != 5 {
		data.NumOldPasswords = types.Int64Null()
	}
	if value := res.Get(pathRoot + `CliTimeout`); value.Exists() && !data.CliTimeout.IsNull() {
		data.CliTimeout = types.Int64Value(value.Int())
	} else if data.CliTimeout.ValueInt64() != 0 {
		data.CliTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxFailedLogin`); value.Exists() && !data.MaxFailedLogin.IsNull() {
		data.MaxFailedLogin = types.Int64Value(value.Int())
	} else if data.MaxFailedLogin.ValueInt64() != 0 {
		data.MaxFailedLogin = types.Int64Null()
	}
	if value := res.Get(pathRoot + `LockoutPeriod`); value.Exists() && !data.LockoutPeriod.IsNull() {
		data.LockoutPeriod = types.Int64Value(value.Int())
	} else if data.LockoutPeriod.ValueInt64() != 1 {
		data.LockoutPeriod = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MCForceDNLDAPOrder`); value.Exists() && !data.McForceDnLdapOrder.IsNull() {
		data.McForceDnLdapOrder = tfutils.BoolFromString(value.String())
	} else if data.McForceDnLdapOrder.ValueBool() {
		data.McForceDnLdapOrder = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PasswordHashAlgorithm`); value.Exists() && !data.PasswordHashAlgorithm.IsNull() {
		data.PasswordHashAlgorithm = tfutils.ParseStringFromGJSON(value)
	} else if data.PasswordHashAlgorithm.ValueString() != "md5crypt" {
		data.PasswordHashAlgorithm = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPSSLClientConfigType`); value.Exists() && !data.LdapSslClientConfigType.IsNull() {
		data.LdapSslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else if data.LdapSslClientConfigType.ValueString() != "client" {
		data.LdapSslClientConfigType = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPSSLClientProfile`); value.Exists() && !data.LdapSslClientProfile.IsNull() {
		data.LdapSslClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapSslClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCLDAPSSLClientConfigType`); value.Exists() && !data.McLdapSslClientConfigType.IsNull() {
		data.McLdapSslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else if data.McLdapSslClientConfigType.ValueString() != "client" {
		data.McLdapSslClientConfigType = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCLDAPSSLClientProfile`); value.Exists() && !data.McLdapSslClientProfile.IsNull() {
		data.McLdapSslClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McLdapSslClientProfile = types.StringNull()
	}
}
