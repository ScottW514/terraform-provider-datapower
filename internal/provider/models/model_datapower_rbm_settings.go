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
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type RBMSettings struct {
	Enabled                           types.Bool                  `tfsdk:"enabled"`
	UserSummary                       types.String                `tfsdk:"user_summary"`
	AuMethod                          types.String                `tfsdk:"au_method"`
	SshauMethod                       *DmRBMSSHAuthenticateType   `tfsdk:"sshau_method"`
	CaPubKeyFile                      types.String                `tfsdk:"ca_pub_key_file"`
	RevokedKeys                       types.List                  `tfsdk:"revoked_keys"`
	AuzosnssConfig                    types.String                `tfsdk:"auzosnss_config"`
	AuoidcScope                       types.String                `tfsdk:"auoidc_scope"`
	AuoidcClientId                    types.String                `tfsdk:"auoidc_client_id"`
	AuoidcClientSecret                types.String                `tfsdk:"auoidc_client_secret"`
	AuoidcIdentityServiceUrl          types.String                `tfsdk:"auoidc_identity_service_url"`
	AuoidcKeyFetchInterval            types.Int64                 `tfsdk:"auoidc_key_fetch_interval"`
	AuoidcIdentityServiceUrlsslClient types.String                `tfsdk:"auoidc_identity_service_urlssl_client"`
	AuKerberosKeytab                  types.String                `tfsdk:"au_kerberos_keytab"`
	AuCustomUrl                       types.String                `tfsdk:"au_custom_url"`
	AuInfoUrl                         types.String                `tfsdk:"au_info_url"`
	AusslValcred                      types.String                `tfsdk:"aussl_valcred"`
	AuHost                            types.String                `tfsdk:"au_host"`
	AuPort                            types.Int64                 `tfsdk:"au_port"`
	AuldapSearchForDn                 types.Bool                  `tfsdk:"auldap_search_for_dn"`
	AuldapBindDn                      types.String                `tfsdk:"auldap_bind_dn"`
	AuldapBindPasswordAlias           types.String                `tfsdk:"auldap_bind_password_alias"`
	AuldapSearchParameters            types.String                `tfsdk:"auldap_search_parameters"`
	AuldapPrefix                      types.String                `tfsdk:"auldap_prefix"`
	AuForceDnldapOrder                types.Bool                  `tfsdk:"au_force_dnldap_order"`
	LdaPsuffix                        types.String                `tfsdk:"lda_psuffix"`
	AuldapLoadBalanceGroup            types.String                `tfsdk:"auldap_load_balance_group"`
	AuCacheAllow                      types.String                `tfsdk:"au_cache_allow"`
	AuCacheTtl                        types.Int64                 `tfsdk:"au_cache_ttl"`
	AuldapReadTimeout                 types.Int64                 `tfsdk:"auldap_read_timeout"`
	McMethod                          types.String                `tfsdk:"mc_method"`
	McCustomUrl                       types.String                `tfsdk:"mc_custom_url"`
	McldapSearchForGroup              types.Bool                  `tfsdk:"mcldap_search_for_group"`
	McHost                            types.String                `tfsdk:"mc_host"`
	McPort                            types.Int64                 `tfsdk:"mc_port"`
	McldapLoadBalanceGroup            types.String                `tfsdk:"mcldap_load_balance_group"`
	McldapBindDn                      types.String                `tfsdk:"mcldap_bind_dn"`
	McldapBindPasswordAlias           types.String                `tfsdk:"mcldap_bind_password_alias"`
	McldapSearchParameters            types.String                `tfsdk:"mcldap_search_parameters"`
	McInfoUrl                         types.String                `tfsdk:"mc_info_url"`
	McldapReadTimeout                 types.Int64                 `tfsdk:"mcldap_read_timeout"`
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
	McForceDnldapOrder                types.Bool                  `tfsdk:"mc_force_dnldap_order"`
	PasswordHashAlgorithm             types.String                `tfsdk:"password_hash_algorithm"`
	LdapsslClientConfigType           types.String                `tfsdk:"ldapssl_client_config_type"`
	LdapsslClientProfile              types.String                `tfsdk:"ldapssl_client_profile"`
	McldapsslClientConfigType         types.String                `tfsdk:"mcldapssl_client_config_type"`
	McldapsslClientProfile            types.String                `tfsdk:"mcldapssl_client_profile"`
	DependencyActions                 []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var RBMSettingsObjectType = map[string]attr.Type{
	"enabled":                               types.BoolType,
	"user_summary":                          types.StringType,
	"au_method":                             types.StringType,
	"sshau_method":                          types.ObjectType{AttrTypes: DmRBMSSHAuthenticateTypeObjectType},
	"ca_pub_key_file":                       types.StringType,
	"revoked_keys":                          types.ListType{ElemType: types.StringType},
	"auzosnss_config":                       types.StringType,
	"auoidc_scope":                          types.StringType,
	"auoidc_client_id":                      types.StringType,
	"auoidc_client_secret":                  types.StringType,
	"auoidc_identity_service_url":           types.StringType,
	"auoidc_key_fetch_interval":             types.Int64Type,
	"auoidc_identity_service_urlssl_client": types.StringType,
	"au_kerberos_keytab":                    types.StringType,
	"au_custom_url":                         types.StringType,
	"au_info_url":                           types.StringType,
	"aussl_valcred":                         types.StringType,
	"au_host":                               types.StringType,
	"au_port":                               types.Int64Type,
	"auldap_search_for_dn":                  types.BoolType,
	"auldap_bind_dn":                        types.StringType,
	"auldap_bind_password_alias":            types.StringType,
	"auldap_search_parameters":              types.StringType,
	"auldap_prefix":                         types.StringType,
	"au_force_dnldap_order":                 types.BoolType,
	"lda_psuffix":                           types.StringType,
	"auldap_load_balance_group":             types.StringType,
	"au_cache_allow":                        types.StringType,
	"au_cache_ttl":                          types.Int64Type,
	"auldap_read_timeout":                   types.Int64Type,
	"mc_method":                             types.StringType,
	"mc_custom_url":                         types.StringType,
	"mcldap_search_for_group":               types.BoolType,
	"mc_host":                               types.StringType,
	"mc_port":                               types.Int64Type,
	"mcldap_load_balance_group":             types.StringType,
	"mcldap_bind_dn":                        types.StringType,
	"mcldap_bind_password_alias":            types.StringType,
	"mcldap_search_parameters":              types.StringType,
	"mc_info_url":                           types.StringType,
	"mcldap_read_timeout":                   types.Int64Type,
	"ldap_version":                          types.StringType,
	"fallback_login":                        types.StringType,
	"fallback_user":                         types.ListType{ElemType: types.StringType},
	"apply_to_cli":                          types.BoolType,
	"restrict_admin_to_serial_port":         types.BoolType,
	"min_password_length":                   types.Int64Type,
	"require_mixed_case":                    types.BoolType,
	"require_digit":                         types.BoolType,
	"require_non_alpha_numeric":             types.BoolType,
	"disallow_username_substring":           types.BoolType,
	"do_password_aging":                     types.BoolType,
	"max_password_age":                      types.Int64Type,
	"do_password_history":                   types.BoolType,
	"num_old_passwords":                     types.Int64Type,
	"cli_timeout":                           types.Int64Type,
	"max_failed_login":                      types.Int64Type,
	"lockout_period":                        types.Int64Type,
	"mc_force_dnldap_order":                 types.BoolType,
	"password_hash_algorithm":               types.StringType,
	"ldapssl_client_config_type":            types.StringType,
	"ldapssl_client_profile":                types.StringType,
	"mcldapssl_client_config_type":          types.StringType,
	"mcldapssl_client_profile":              types.StringType,
	"dependency_actions":                    actions.ActionsListType,
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
	if data.SshauMethod != nil {
		if !data.SshauMethod.IsNull() {
			return false
		}
	}
	if !data.CaPubKeyFile.IsNull() {
		return false
	}
	if !data.RevokedKeys.IsNull() {
		return false
	}
	if !data.AuzosnssConfig.IsNull() {
		return false
	}
	if !data.AuoidcScope.IsNull() {
		return false
	}
	if !data.AuoidcClientId.IsNull() {
		return false
	}
	if !data.AuoidcClientSecret.IsNull() {
		return false
	}
	if !data.AuoidcIdentityServiceUrl.IsNull() {
		return false
	}
	if !data.AuoidcKeyFetchInterval.IsNull() {
		return false
	}
	if !data.AuoidcIdentityServiceUrlsslClient.IsNull() {
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
	if !data.AusslValcred.IsNull() {
		return false
	}
	if !data.AuHost.IsNull() {
		return false
	}
	if !data.AuPort.IsNull() {
		return false
	}
	if !data.AuldapSearchForDn.IsNull() {
		return false
	}
	if !data.AuldapBindDn.IsNull() {
		return false
	}
	if !data.AuldapBindPasswordAlias.IsNull() {
		return false
	}
	if !data.AuldapSearchParameters.IsNull() {
		return false
	}
	if !data.AuldapPrefix.IsNull() {
		return false
	}
	if !data.AuForceDnldapOrder.IsNull() {
		return false
	}
	if !data.LdaPsuffix.IsNull() {
		return false
	}
	if !data.AuldapLoadBalanceGroup.IsNull() {
		return false
	}
	if !data.AuCacheAllow.IsNull() {
		return false
	}
	if !data.AuCacheTtl.IsNull() {
		return false
	}
	if !data.AuldapReadTimeout.IsNull() {
		return false
	}
	if !data.McMethod.IsNull() {
		return false
	}
	if !data.McCustomUrl.IsNull() {
		return false
	}
	if !data.McldapSearchForGroup.IsNull() {
		return false
	}
	if !data.McHost.IsNull() {
		return false
	}
	if !data.McPort.IsNull() {
		return false
	}
	if !data.McldapLoadBalanceGroup.IsNull() {
		return false
	}
	if !data.McldapBindDn.IsNull() {
		return false
	}
	if !data.McldapBindPasswordAlias.IsNull() {
		return false
	}
	if !data.McldapSearchParameters.IsNull() {
		return false
	}
	if !data.McInfoUrl.IsNull() {
		return false
	}
	if !data.McldapReadTimeout.IsNull() {
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
	if !data.McForceDnldapOrder.IsNull() {
		return false
	}
	if !data.PasswordHashAlgorithm.IsNull() {
		return false
	}
	if !data.LdapsslClientConfigType.IsNull() {
		return false
	}
	if !data.LdapsslClientProfile.IsNull() {
		return false
	}
	if !data.McldapsslClientConfigType.IsNull() {
		return false
	}
	if !data.McldapsslClientProfile.IsNull() {
		return false
	}
	return true
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
	if data.SshauMethod != nil {
		if !data.SshauMethod.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`SSHAUMethod`, data.SshauMethod.ToBody(ctx, ""))
		}
	}
	if !data.CaPubKeyFile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CAPubKeyFile`, data.CaPubKeyFile.ValueString())
	}
	if !data.RevokedKeys.IsNull() {
		var values []string
		data.RevokedKeys.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`RevokedKeys`+".-1", map[string]string{"value": val})
		}
	}
	if !data.AuzosnssConfig.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUZOSNSSConfig`, data.AuzosnssConfig.ValueString())
	}
	if !data.AuoidcScope.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUOIDCScope`, data.AuoidcScope.ValueString())
	}
	if !data.AuoidcClientId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUOIDCClientID`, data.AuoidcClientId.ValueString())
	}
	if !data.AuoidcClientSecret.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUOIDCClientSecret`, data.AuoidcClientSecret.ValueString())
	}
	if !data.AuoidcIdentityServiceUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUOIDCIdentityServiceURL`, data.AuoidcIdentityServiceUrl.ValueString())
	}
	if !data.AuoidcKeyFetchInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUOIDCKeyFetchInterval`, data.AuoidcKeyFetchInterval.ValueInt64())
	}
	if !data.AuoidcIdentityServiceUrlsslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUOIDCIdentityServiceURLSSLClient`, data.AuoidcIdentityServiceUrlsslClient.ValueString())
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
	if !data.AusslValcred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUSSLValcred`, data.AusslValcred.ValueString())
	}
	if !data.AuHost.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUHost`, data.AuHost.ValueString())
	}
	if !data.AuPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUPort`, data.AuPort.ValueInt64())
	}
	if !data.AuldapSearchForDn.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPSearchForDN`, tfutils.StringFromBool(data.AuldapSearchForDn, ""))
	}
	if !data.AuldapBindDn.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPBindDN`, data.AuldapBindDn.ValueString())
	}
	if !data.AuldapBindPasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPBindPasswordAlias`, data.AuldapBindPasswordAlias.ValueString())
	}
	if !data.AuldapSearchParameters.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPSearchParameters`, data.AuldapSearchParameters.ValueString())
	}
	if !data.AuldapPrefix.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPPrefix`, data.AuldapPrefix.ValueString())
	}
	if !data.AuForceDnldapOrder.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUForceDNLDAPOrder`, tfutils.StringFromBool(data.AuForceDnldapOrder, ""))
	}
	if !data.LdaPsuffix.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPsuffix`, data.LdaPsuffix.ValueString())
	}
	if !data.AuldapLoadBalanceGroup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPLoadBalanceGroup`, data.AuldapLoadBalanceGroup.ValueString())
	}
	if !data.AuCacheAllow.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUCacheAllow`, data.AuCacheAllow.ValueString())
	}
	if !data.AuCacheTtl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUCacheTTL`, data.AuCacheTtl.ValueInt64())
	}
	if !data.AuldapReadTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AULDAPReadTimeout`, data.AuldapReadTimeout.ValueInt64())
	}
	if !data.McMethod.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCMethod`, data.McMethod.ValueString())
	}
	if !data.McCustomUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCCustomURL`, data.McCustomUrl.ValueString())
	}
	if !data.McldapSearchForGroup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCLDAPSearchForGroup`, tfutils.StringFromBool(data.McldapSearchForGroup, ""))
	}
	if !data.McHost.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCHost`, data.McHost.ValueString())
	}
	if !data.McPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCPort`, data.McPort.ValueInt64())
	}
	if !data.McldapLoadBalanceGroup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCLDAPLoadBalanceGroup`, data.McldapLoadBalanceGroup.ValueString())
	}
	if !data.McldapBindDn.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCLDAPBindDN`, data.McldapBindDn.ValueString())
	}
	if !data.McldapBindPasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCLDAPBindPasswordAlias`, data.McldapBindPasswordAlias.ValueString())
	}
	if !data.McldapSearchParameters.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCLDAPSearchParameters`, data.McldapSearchParameters.ValueString())
	}
	if !data.McInfoUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCInfoURL`, data.McInfoUrl.ValueString())
	}
	if !data.McldapReadTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCLDAPReadTimeout`, data.McldapReadTimeout.ValueInt64())
	}
	if !data.LdapVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPVersion`, data.LdapVersion.ValueString())
	}
	if !data.FallbackLogin.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FallbackLogin`, data.FallbackLogin.ValueString())
	}
	if !data.FallbackUser.IsNull() {
		var values []string
		data.FallbackUser.ElementsAs(ctx, &values, false)
		for _, val := range values {
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
	if !data.McForceDnldapOrder.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCForceDNLDAPOrder`, tfutils.StringFromBool(data.McForceDnldapOrder, ""))
	}
	if !data.PasswordHashAlgorithm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PasswordHashAlgorithm`, data.PasswordHashAlgorithm.ValueString())
	}
	if !data.LdapsslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPSSLClientConfigType`, data.LdapsslClientConfigType.ValueString())
	}
	if !data.LdapsslClientProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPSSLClientProfile`, data.LdapsslClientProfile.ValueString())
	}
	if !data.McldapsslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCLDAPSSLClientConfigType`, data.McldapsslClientConfigType.ValueString())
	}
	if !data.McldapsslClientProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MCLDAPSSLClientProfile`, data.McldapsslClientProfile.ValueString())
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
		data.SshauMethod = &DmRBMSSHAuthenticateType{}
		data.SshauMethod.FromBody(ctx, "", value)
	} else {
		data.SshauMethod = nil
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
		data.AuzosnssConfig = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuzosnssConfig = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUOIDCScope`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuoidcScope = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuoidcScope = types.StringValue("openid")
	}
	if value := res.Get(pathRoot + `AUOIDCClientID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuoidcClientId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuoidcClientId = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUOIDCClientSecret`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuoidcClientSecret = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuoidcClientSecret = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUOIDCIdentityServiceURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuoidcIdentityServiceUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuoidcIdentityServiceUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUOIDCKeyFetchInterval`); value.Exists() {
		data.AuoidcKeyFetchInterval = types.Int64Value(value.Int())
	} else {
		data.AuoidcKeyFetchInterval = types.Int64Value(30)
	}
	if value := res.Get(pathRoot + `AUOIDCIdentityServiceURLSSLClient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuoidcIdentityServiceUrlsslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuoidcIdentityServiceUrlsslClient = types.StringNull()
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
		data.AusslValcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AusslValcred = types.StringNull()
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
		data.AuldapSearchForDn = tfutils.BoolFromString(value.String())
	} else {
		data.AuldapSearchForDn = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AULDAPBindDN`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuldapBindDn = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuldapBindDn = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPBindPasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuldapBindPasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuldapBindPasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPSearchParameters`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuldapSearchParameters = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuldapSearchParameters = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPPrefix`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuldapPrefix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuldapPrefix = types.StringValue("cn=")
	}
	if value := res.Get(pathRoot + `AUForceDNLDAPOrder`); value.Exists() {
		data.AuForceDnldapOrder = tfutils.BoolFromString(value.String())
	} else {
		data.AuForceDnldapOrder = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LDAPsuffix`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdaPsuffix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdaPsuffix = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPLoadBalanceGroup`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuldapLoadBalanceGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuldapLoadBalanceGroup = types.StringNull()
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
		data.AuldapReadTimeout = types.Int64Value(value.Int())
	} else {
		data.AuldapReadTimeout = types.Int64Value(60)
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
		data.McldapSearchForGroup = tfutils.BoolFromString(value.String())
	} else {
		data.McldapSearchForGroup = types.BoolNull()
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
		data.McldapLoadBalanceGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McldapLoadBalanceGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCLDAPBindDN`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.McldapBindDn = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McldapBindDn = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCLDAPBindPasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.McldapBindPasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McldapBindPasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCLDAPSearchParameters`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.McldapSearchParameters = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McldapSearchParameters = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCInfoURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.McInfoUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McInfoUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCLDAPReadTimeout`); value.Exists() {
		data.McldapReadTimeout = types.Int64Value(value.Int())
	} else {
		data.McldapReadTimeout = types.Int64Value(60)
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
		data.McForceDnldapOrder = tfutils.BoolFromString(value.String())
	} else {
		data.McForceDnldapOrder = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PasswordHashAlgorithm`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PasswordHashAlgorithm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordHashAlgorithm = types.StringValue("md5crypt")
	}
	if value := res.Get(pathRoot + `LDAPSSLClientConfigType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapsslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapsslClientConfigType = types.StringValue("client")
	}
	if value := res.Get(pathRoot + `LDAPSSLClientProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapsslClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapsslClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCLDAPSSLClientConfigType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.McldapsslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McldapsslClientConfigType = types.StringValue("client")
	}
	if value := res.Get(pathRoot + `MCLDAPSSLClientProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.McldapsslClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McldapsslClientProfile = types.StringNull()
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
		data.SshauMethod.UpdateFromBody(ctx, "", value)
	} else {
		data.SshauMethod = nil
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
	if value := res.Get(pathRoot + `AUZOSNSSConfig`); value.Exists() && !data.AuzosnssConfig.IsNull() {
		data.AuzosnssConfig = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuzosnssConfig = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUOIDCScope`); value.Exists() && !data.AuoidcScope.IsNull() {
		data.AuoidcScope = tfutils.ParseStringFromGJSON(value)
	} else if data.AuoidcScope.ValueString() != "openid" {
		data.AuoidcScope = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUOIDCClientID`); value.Exists() && !data.AuoidcClientId.IsNull() {
		data.AuoidcClientId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuoidcClientId = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUOIDCClientSecret`); value.Exists() && !data.AuoidcClientSecret.IsNull() {
		data.AuoidcClientSecret = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuoidcClientSecret = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUOIDCIdentityServiceURL`); value.Exists() && !data.AuoidcIdentityServiceUrl.IsNull() {
		data.AuoidcIdentityServiceUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuoidcIdentityServiceUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUOIDCKeyFetchInterval`); value.Exists() && !data.AuoidcKeyFetchInterval.IsNull() {
		data.AuoidcKeyFetchInterval = types.Int64Value(value.Int())
	} else if data.AuoidcKeyFetchInterval.ValueInt64() != 30 {
		data.AuoidcKeyFetchInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AUOIDCIdentityServiceURLSSLClient`); value.Exists() && !data.AuoidcIdentityServiceUrlsslClient.IsNull() {
		data.AuoidcIdentityServiceUrlsslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuoidcIdentityServiceUrlsslClient = types.StringNull()
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
	if value := res.Get(pathRoot + `AUSSLValcred`); value.Exists() && !data.AusslValcred.IsNull() {
		data.AusslValcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AusslValcred = types.StringNull()
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
	if value := res.Get(pathRoot + `AULDAPSearchForDN`); value.Exists() && !data.AuldapSearchForDn.IsNull() {
		data.AuldapSearchForDn = tfutils.BoolFromString(value.String())
	} else if data.AuldapSearchForDn.ValueBool() {
		data.AuldapSearchForDn = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AULDAPBindDN`); value.Exists() && !data.AuldapBindDn.IsNull() {
		data.AuldapBindDn = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuldapBindDn = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPBindPasswordAlias`); value.Exists() && !data.AuldapBindPasswordAlias.IsNull() {
		data.AuldapBindPasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuldapBindPasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPSearchParameters`); value.Exists() && !data.AuldapSearchParameters.IsNull() {
		data.AuldapSearchParameters = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuldapSearchParameters = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPPrefix`); value.Exists() && !data.AuldapPrefix.IsNull() {
		data.AuldapPrefix = tfutils.ParseStringFromGJSON(value)
	} else if data.AuldapPrefix.ValueString() != "cn=" {
		data.AuldapPrefix = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUForceDNLDAPOrder`); value.Exists() && !data.AuForceDnldapOrder.IsNull() {
		data.AuForceDnldapOrder = tfutils.BoolFromString(value.String())
	} else if data.AuForceDnldapOrder.ValueBool() {
		data.AuForceDnldapOrder = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LDAPsuffix`); value.Exists() && !data.LdaPsuffix.IsNull() {
		data.LdaPsuffix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdaPsuffix = types.StringNull()
	}
	if value := res.Get(pathRoot + `AULDAPLoadBalanceGroup`); value.Exists() && !data.AuldapLoadBalanceGroup.IsNull() {
		data.AuldapLoadBalanceGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuldapLoadBalanceGroup = types.StringNull()
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
	if value := res.Get(pathRoot + `AULDAPReadTimeout`); value.Exists() && !data.AuldapReadTimeout.IsNull() {
		data.AuldapReadTimeout = types.Int64Value(value.Int())
	} else if data.AuldapReadTimeout.ValueInt64() != 60 {
		data.AuldapReadTimeout = types.Int64Null()
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
	if value := res.Get(pathRoot + `MCLDAPSearchForGroup`); value.Exists() && !data.McldapSearchForGroup.IsNull() {
		data.McldapSearchForGroup = tfutils.BoolFromString(value.String())
	} else if data.McldapSearchForGroup.ValueBool() {
		data.McldapSearchForGroup = types.BoolNull()
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
	if value := res.Get(pathRoot + `MCLDAPLoadBalanceGroup`); value.Exists() && !data.McldapLoadBalanceGroup.IsNull() {
		data.McldapLoadBalanceGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McldapLoadBalanceGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCLDAPBindDN`); value.Exists() && !data.McldapBindDn.IsNull() {
		data.McldapBindDn = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McldapBindDn = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCLDAPBindPasswordAlias`); value.Exists() && !data.McldapBindPasswordAlias.IsNull() {
		data.McldapBindPasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McldapBindPasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCLDAPSearchParameters`); value.Exists() && !data.McldapSearchParameters.IsNull() {
		data.McldapSearchParameters = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McldapSearchParameters = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCInfoURL`); value.Exists() && !data.McInfoUrl.IsNull() {
		data.McInfoUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McInfoUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCLDAPReadTimeout`); value.Exists() && !data.McldapReadTimeout.IsNull() {
		data.McldapReadTimeout = types.Int64Value(value.Int())
	} else if data.McldapReadTimeout.ValueInt64() != 60 {
		data.McldapReadTimeout = types.Int64Null()
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
	if value := res.Get(pathRoot + `MCForceDNLDAPOrder`); value.Exists() && !data.McForceDnldapOrder.IsNull() {
		data.McForceDnldapOrder = tfutils.BoolFromString(value.String())
	} else if data.McForceDnldapOrder.ValueBool() {
		data.McForceDnldapOrder = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PasswordHashAlgorithm`); value.Exists() && !data.PasswordHashAlgorithm.IsNull() {
		data.PasswordHashAlgorithm = tfutils.ParseStringFromGJSON(value)
	} else if data.PasswordHashAlgorithm.ValueString() != "md5crypt" {
		data.PasswordHashAlgorithm = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPSSLClientConfigType`); value.Exists() && !data.LdapsslClientConfigType.IsNull() {
		data.LdapsslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else if data.LdapsslClientConfigType.ValueString() != "client" {
		data.LdapsslClientConfigType = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPSSLClientProfile`); value.Exists() && !data.LdapsslClientProfile.IsNull() {
		data.LdapsslClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapsslClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCLDAPSSLClientConfigType`); value.Exists() && !data.McldapsslClientConfigType.IsNull() {
		data.McldapsslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else if data.McldapsslClientConfigType.ValueString() != "client" {
		data.McldapsslClientConfigType = types.StringNull()
	}
	if value := res.Get(pathRoot + `MCLDAPSSLClientProfile`); value.Exists() && !data.McldapsslClientProfile.IsNull() {
		data.McldapsslClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.McldapsslClientProfile = types.StringNull()
	}
}
