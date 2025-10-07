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

package resources

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

var _ resource.Resource = &RBMSettingsResource{}
var _ resource.ResourceWithValidateConfig = &RBMSettingsResource{}
var _ resource.ResourceWithImportState = &RBMSettingsResource{}

func NewRBMSettingsResource() resource.Resource {
	return &RBMSettingsResource{}
}

type RBMSettingsResource struct {
	pData *tfutils.ProviderData
}

func (r *RBMSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rbm_settings"
}

func (r *RBMSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("<p>Manage role-based management (RBM) settings: RBM, password policy, and account policy</p><ul><li>RBM consists of the following capabilities: Authenticating users, evaluating the access profile, enforcing access to resources</li><li>The password policy sets the password requirements for local user accounts.</li><li>The account policy sets the lockout behavior and the timeout for CLI sessions.</li></ul>", "rbm", "").AddActions("flush_cache").String,
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>The administrative state of the configuration.</p><ul><li>To make active, set to enabled.</li><li>To make inactive, set to disabled.</li></ul>", "admin-state", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"au_method": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Authentication method", "au-method", "").AddStringEnum("local", "xmlfile", "ldap", "radius", "spnego", "zosnss", "custom", "client-ssl", "oidc").AddDefaultValue("local").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("local", "xmlfile", "ldap", "radius", "spnego", "zosnss", "custom", "client-ssl", "oidc"),
				},
				Default: stringdefault.StaticString("local"),
			},
			"ssh_au_method": models.GetDmRBMSSHAuthenticateTypeResourceSchema("Specify the method to authenticate SSH users. <ul><li>When no method, the user is prompted for both username and password.</li><li>When password, the user is prompted for the password. For this method, the username must be part of the invocation. With the ssh command, the invocation is in the ssh username@host format.</li><li>When user certificate, the user is not prompted for input. The connection is successful when the invocation provides a signed SSH user certificate that is verified by the CA public key file in the <tt>cert:</tt> directory. With the ssh command, the invocation must include the -i file parameter.</li><li>When both certificate and password, processing attempts to first authenticate with the provided signed SSH user certificate. If unsuccessful, prompts for the password.</li><li>Supported RBM authentication methods with SSH authentication method are local and LDAP. <ul><li>Local authentication method extracts the certificate identity and attempts login with local User of that name.</li><li>LDAP authentication method constructs the DN through LDAP search or by applying the configured prefix and suffix. Since SSH authentication completes prior to the Authentication step no LDAP bind or authenticate will take place.</li></ul></li></ul>", "ssh-au-method", "", false),
			"ca_pub_key_file": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the certificate authority (CA) public key file in the <tt>cert:</tt> directory for SSH authentication with SSH user certificates. This public key file contains the public key for one or more certificate authorities.", "ssh-ca-pubkey-file", "").AddRequiredWhen(models.RBMSettingsCAPubKeyFileCondVal.String()).AddNotValidWhen(models.RBMSettingsCAPubKeyFileIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.RBMSettingsCAPubKeyFileCondVal, models.RBMSettingsCAPubKeyFileIgnoreVal, false),
				},
			},
			"revoked_keys": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the OpenSSH public keys to revoke for SSH authentication. Each entry is the public key file in the <tt>cert:</tt> or <tt>sharedcert:</tt> directory and must be in the OpenSSH public key format. These keys are signed by the CA user public key file.", "ssh-revoked-keys", "").AddNotValidWhen(models.RBMSettingsRevokedKeysIgnoreVal.String()).String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"au_zos_nss_config": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("z/OS NSS client", "zos-nss-au", "zos_nss_client").AddRequiredWhen(models.RBMSettingsAUZOSNSSConfigCondVal.String()).AddNotValidWhen(models.RBMSettingsAUZOSNSSConfigIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.RBMSettingsAUZOSNSSConfigCondVal, models.RBMSettingsAUZOSNSSConfigIgnoreVal, false),
				},
			},
			"au_oidc_scope": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Scope", "au-oidc-scope", "").AddDefaultValue("openid").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("openid"),
			},
			"au_oidc_client_id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Client ID", "au-oidc-client-id", "").String,
				Optional:            true,
			},
			"au_oidc_client_secret": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Client secret", "au-oidc-client-secret", "password_alias").String,
				Optional:            true,
			},
			"au_oidc_identity_service_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Identity service URL", "au-oidc-identity-service-url", "").String,
				Optional:            true,
			},
			"au_oidc_key_fetch_interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Retrieval interval for public keys.", "au-oidc-key-fetch-interval", "").AddIntegerRange(1, 86400).AddDefaultValue("30").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(30),
			},
			"au_oidc_identity_service_url_ssl_client": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "au-oidc-identity-service-ssl-client", "ssl_client_profile").String,
				Optional:            true,
			},
			"au_kerberos_keytab": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Kerberos keytab", "au-kerberos-keytab", "crypto_kerberos_keytab").AddRequiredWhen(models.RBMSettingsAUKerberosKeytabCondVal.String()).AddNotValidWhen(models.RBMSettingsAUKerberosKeytabIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.RBMSettingsAUKerberosKeytabCondVal, models.RBMSettingsAUKerberosKeytabIgnoreVal, false),
				},
			},
			"au_custom_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Custom URL", "au-custom-url", "").AddRequiredWhen(models.RBMSettingsAUCustomURLCondVal.String()).AddNotValidWhen(models.RBMSettingsAUCustomURLIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.RBMSettingsAUCustomURLCondVal, models.RBMSettingsAUCustomURLIgnoreVal, false),
				},
			},
			"au_info_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL of the XML file for authentication. The XML file can be on the DataPower Gateway or on a remote server. You can use the same XML file to map credentials.", "au-info-url", "").AddRequiredWhen(models.RBMSettingsAUInfoURLCondVal.String()).AddNotValidWhen(models.RBMSettingsAUInfoURLIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.RBMSettingsAUInfoURLCondVal, models.RBMSettingsAUInfoURLIgnoreVal, false),
				},
			},
			"au_ssl_valcred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validation credentials", "au-valcred", "crypto_val_cred").AddRequiredWhen(models.RBMSettingsAUSSLValcredCondVal.String()).AddNotValidWhen(models.RBMSettingsAUSSLValcredIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.RBMSettingsAUSSLValcredCondVal, models.RBMSettingsAUSSLValcredIgnoreVal, false),
				},
			},
			"au_host": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Server host", "au-server-host", "").AddRequiredWhen(models.RBMSettingsAUHostCondVal.String()).AddNotValidWhen(models.RBMSettingsAUHostIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.RBMSettingsAUHostCondVal, models.RBMSettingsAUHostIgnoreVal, false),
				},
			},
			"au_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Server port", "au-server-port", "").AddRequiredWhen(models.RBMSettingsAUPortCondVal.String()).AddNotValidWhen(models.RBMSettingsAUPortIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.Int64{
					validators.ConditionalRequiredInt64(models.RBMSettingsAUPortCondVal, models.RBMSettingsAUPortIgnoreVal, false),
				},
			},
			"au_ldap_search_for_dn": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to retrieve the user DN with an LDAP search. <ul><li>When enabled, the login name presented by the user is used with the LDAP search parameters for an LDAP search to retrieve the user DN.</li><li>When disabled, the login name presented by the user is used with the LDAP prefix and suffix to construct the user DN.</li></ul>", "au-ldap-search", "").AddDefaultValue("false").AddNotValidWhen(models.RBMSettingsAULDAPSearchForDNIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"au_ldap_bind_dn": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP bind DN", "au-ldap-bind-dn", "").AddNotValidWhen(models.RBMSettingsAULDAPBindDNIgnoreVal.String()).String,
				Optional:            true,
			},
			"au_ldap_bind_password_alias": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP bind password alias", "au-ldap-bind-password-alias", "password_alias").AddNotValidWhen(models.RBMSettingsAULDAPBindPasswordAliasIgnoreVal.String()).String,
				Optional:            true,
			},
			"au_ldap_search_parameters": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP search parameters", "au-ldap-parameters", "ldap_search_parameters").AddRequiredWhen(models.RBMSettingsAULDAPSearchParametersCondVal.String()).AddNotValidWhen(models.RBMSettingsAULDAPSearchParametersIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.RBMSettingsAULDAPSearchParametersCondVal, models.RBMSettingsAULDAPSearchParametersIgnoreVal, false),
				},
			},
			"au_ldap_prefix": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the string to add before the username to form the DN. If this value is <tt>CN=</tt> and the username is <tt>Bob</tt> , the complete DN is <tt>CN=Bob,O=example.com</tt> when the LDAP suffix is <tt>O=example.com</tt> .", "ldap-prefix", "").AddDefaultValue("cn=").AddNotValidWhen(models.RBMSettingsAULDAPPrefixIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("cn="),
			},
			"au_force_dn_ldap_order": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to convert the extracted DN to LDAP format. This property is essential when the extracted DN from a TLS certificate is in X.500 format. This format arranges the RDNs of the DNs from left to right with forward slashes as separators; for example, <tt>C=US/O=My Organization/CN=Fred</tt> . <p>When you retrieve the group name with an LDAP search, the authenticated DN must be in LDAP format. This format arranges the RDNs of the DNs from right to left with commas as separators; for example, <tt>CN=Fred, O=My Organization, C=US</tt> .</p>", "au-force-dn-ldap-order", "").AddDefaultValue("false").AddNotValidWhen(models.RBMSettingsAUForceDNLDAPOrderIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ldap_suffix": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the string to add after the username to form the DN. If this value is <tt>O=example.com</tt> and the username is <tt>Bob</tt> , the complete DN is <tt>CN=Bob,O=example.com</tt> when the LDAP prefix is <tt>CN=</tt> .", "ldap-suffix", "").AddNotValidWhen(models.RBMSettingsLDAPsuffixIgnoreVal.String()).String,
				Optional:            true,
			},
			"au_ldap_load_balance_group": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the load balancer group of LDAP servers. This setting overrides the settings for the server host and port.", "loadbalancer-group", "load_balancer_group").AddNotValidWhen(models.RBMSettingsAULDAPLoadBalanceGroupIgnoreVal.String()).String,
				Optional:            true,
			},
			"au_cache_allow": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Authentication cache mode", "au-cache-mode", "").AddStringEnum("absolute", "disabled", "maximum", "minimum").AddDefaultValue("absolute").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("absolute", "disabled", "maximum", "minimum"),
				},
				Default: stringdefault.StaticString("absolute"),
			},
			"au_cache_ttl": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the expiry for cached authentication decisions. Enter a value in the range 1 - 86400. The default value is 600.", "au-cache-ttl", "").AddIntegerRange(1, 86400).AddDefaultValue("600").AddNotValidWhen(models.RBMSettingsAUCacheTTLIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 86400),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, models.RBMSettingsAUCacheTTLIgnoreVal, true),
				},
				Default: int64default.StaticInt64(600),
			},
			"au_ldap_read_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the time to wait for a response from the LDAP server before the DataPower Gateway closes the LDAP connection. Enter a value in the range 0 - 86400. The default value is 60. A value of 0 indicates that the connection never times out.", "au-ldap-readtimeout", "").AddIntegerRange(0, 86400).AddDefaultValue("60").AddNotValidWhen(models.RBMSettingsAULDAPReadTimeoutIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 86400),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, models.RBMSettingsAULDAPReadTimeoutIgnoreVal, true),
				},
				Default: int64default.StaticInt64(60),
			},
			"mc_method": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Credential-mapping method", "mc-method", "").AddStringEnum("local", "xmlfile", "custom").AddDefaultValue("local").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("local", "xmlfile", "custom"),
				},
				Default: stringdefault.StaticString("local"),
			},
			"mc_custom_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Custom URL", "mc-custom-url", "").AddRequiredWhen(models.RBMSettingsMCCustomURLCondVal.String()).AddNotValidWhen(models.RBMSettingsMCCustomURLIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.RBMSettingsMCCustomURLCondVal, models.RBMSettingsMCCustomURLIgnoreVal, false),
				},
			},
			"mc_ldap_search_for_group": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to search LDAP to retrieve all user groups that match the query. <ul><li>When enabled, the authenticated DN of the user and the LDAP search parameters are used as part of the LDAP search to retrieve all user groups that match the query. When a user belongs to multiple groups, the resultant access policy for this user is additive not most restrictive.</li><li>When disabled, the authenticated identity of the user (DN or user group of local user) is used directly as the input credential.</li></ul>", "mc-ldap-search", "").AddDefaultValue("false").AddNotValidWhen(models.RBMSettingsMCLDAPSearchForGroupIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"mc_host": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Server host", "mc-server-host", "").AddRequiredWhen(models.RBMSettingsMCHostCondVal.String()).AddNotValidWhen(models.RBMSettingsMCHostIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.RBMSettingsMCHostCondVal, models.RBMSettingsMCHostIgnoreVal, false),
				},
			},
			"mc_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Server port", "mc-server-port", "").AddRequiredWhen(models.RBMSettingsMCPortCondVal.String()).AddNotValidWhen(models.RBMSettingsMCPortIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.Int64{
					validators.ConditionalRequiredInt64(models.RBMSettingsMCPortCondVal, models.RBMSettingsMCPortIgnoreVal, false),
				},
			},
			"mc_ldap_load_balance_group": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Load balancer group", "mc-loadbalancer-group", "load_balancer_group").AddNotValidWhen(models.RBMSettingsMCLDAPLoadBalanceGroupIgnoreVal.String()).String,
				Optional:            true,
			},
			"mc_ldap_bind_dn": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP bind DN", "mc-ldap-bind-dn", "").AddNotValidWhen(models.RBMSettingsMCLDAPBindDNIgnoreVal.String()).String,
				Optional:            true,
			},
			"mc_ldap_bind_password_alias": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP bind password alias", "mc-ldap-bind-password-alias", "password_alias").AddNotValidWhen(models.RBMSettingsMCLDAPBindPasswordAliasIgnoreVal.String()).String,
				Optional:            true,
			},
			"mc_ldap_search_parameters": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP search parameters", "mc-ldap-parameters", "ldap_search_parameters").AddRequiredWhen(models.RBMSettingsMCLDAPSearchParametersCondVal.String()).AddNotValidWhen(models.RBMSettingsMCLDAPSearchParametersIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.RBMSettingsMCLDAPSearchParametersCondVal, models.RBMSettingsMCLDAPSearchParametersIgnoreVal, false),
				},
			},
			"mc_info_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL of the XML file to map credentials. The XML file can be on the DataPower Gateway or on a remote server. You can use the same XML file for authentication.", "mc-info-url", "").AddRequiredWhen(models.RBMSettingsMCInfoURLCondVal.String()).AddNotValidWhen(models.RBMSettingsMCInfoURLIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.RBMSettingsMCInfoURLCondVal, models.RBMSettingsMCInfoURLIgnoreVal, false),
				},
			},
			"mc_ldap_read_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the time to wait for a response from the LDAP server before the DataPower Gateway closes the LDAP connection. Enter a value in the range 0 - 86400. The default value is 60. A value of 0 indicates that the connection never times out.", "mc-ldap-readtimeout", "").AddIntegerRange(0, 86400).AddDefaultValue("60").AddNotValidWhen(models.RBMSettingsMCLDAPReadTimeoutIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 86400),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, models.RBMSettingsMCLDAPReadTimeoutIgnoreVal, true),
				},
				Default: int64default.StaticInt64(60),
			},
			"ldap_version": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP version", "ldap-version", "").AddStringEnum("v2", "v3").AddNotValidWhen(models.RBMSettingsLDAPVersionIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("v2", "v3"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.RBMSettingsLDAPVersionIgnoreVal, false),
				},
			},
			"fallback_login": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to use local user accounts as fallback users if remote authentication fails. With fallback users, local user accounts can log on to the DataPower Gateway if authentication fails or during a network outage that affects primary authentication. The recommendation is to restrict fallback users to a subset of local user accounts. <p><b>Note: </b>When authentication uses a TLS certificate from a connection peer, you cannot enforce RBM on CLI sessions unless fallback users are supported.</p>", "fallback-login", "").AddStringEnum("disabled", "local", "restricted").AddDefaultValue("disabled").AddNotValidWhen(models.RBMSettingsFallbackLoginIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disabled", "local", "restricted"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.RBMSettingsFallbackLoginIgnoreVal, true),
				},
				Default: stringdefault.StaticString("disabled"),
			},
			"fallback_user": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Fallback users", "fallback-user", "user").AddRequiredWhen(models.RBMSettingsFallbackUserCondVal.String()).AddNotValidWhen(models.RBMSettingsFallbackUserIgnoreVal.String()).String,
				ElementType:         types.StringType,
				Optional:            true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(),
					validators.ConditionalRequiredList(models.RBMSettingsFallbackUserCondVal, models.RBMSettingsFallbackUserIgnoreVal, false),
				},
			},
			"apply_to_cli": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enforce the RBM policy on CLI sessions. When authentication uses a TLS certificate from a connection peer, you cannot enforce RBM on CLI sessions unless fallback users are supported.", "apply-cli", "").AddDefaultValue("false").AddNotValidWhen(models.RBMSettingsApplyToCLIIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"restrict_admin_to_serial_port": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Restrict admin to serial", "restrict-admin", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"min_password_length": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the minimum length of a valid password. Enter a value in the range 1 - 128. The default value depends on common criteria mode. <ul><li>When common criteria is enabled, the default value is 14.</li><li>When common criteria is not enabled, the default value is 6.</li></ul>", "pwd-minimum-length", "").AddIntegerRange(1, 128).AddDefaultValue("6").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 128),
				},
				Default: int64default.StaticInt64(6),
			},
			"require_mixed_case": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Require mixed case", "pwd-mixed-case", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"require_digit": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Require number", "pwd-digit", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"require_non_alpha_numeric": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Require non-alphanumeric", "pwd-nonalphanumeric", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"disallow_username_substring": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Disallow username substring", "pwd-username", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"do_password_aging": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable aging", "pwd-aging", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"max_password_age": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the expiry for passwords. The default value depends on common criteria mode. <ul><li>When common criteria is enabled, the default value is 90.</li><li>When common criteria is not enabled, the default value is 30.</li></ul>", "pwd-max-age", "").AddIntegerRange(1, 65535).AddDefaultValue("30").AddRequiredWhen(models.RBMSettingsMaxPasswordAgeCondVal.String()).AddNotValidWhen(models.RBMSettingsMaxPasswordAgeIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
					validators.ConditionalRequiredInt64(models.RBMSettingsMaxPasswordAgeCondVal, models.RBMSettingsMaxPasswordAgeIgnoreVal, true),
				},
				Default: int64default.StaticInt64(30),
			},
			"do_password_history": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Control reuse", "pwd-history", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"num_old_passwords": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the number of recent passwords to track to prevent reuse. The default value depends on common criteria mode. <ul><li>When common criteria is enabled, the default value is 3.</li><li>When common criteria is not enabled, the default value is 5.</li></ul>", "pwd-max-history", "").AddIntegerRange(1, 65535).AddDefaultValue("5").AddRequiredWhen(models.RBMSettingsNumOldPasswordsCondVal.String()).AddNotValidWhen(models.RBMSettingsNumOldPasswordsIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
					validators.ConditionalRequiredInt64(models.RBMSettingsNumOldPasswordsCondVal, models.RBMSettingsNumOldPasswordsIgnoreVal, true),
				},
				Default: int64default.StaticInt64(5),
			},
			"cli_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the time after which to invalidate idle CLI sessions. When invalidated, requires re-authentication. Enter a value in the range 0 - 65535. A value of 0 disables the timer. The default value depends on common criteria mode. <ul><li>When common criteria is enabled, the default value is 900.</li><li>When common criteria is not enabled, the default value is 0.</li></ul>", "cli-timeout", "").AddIntegerRange(0, 65535).AddDefaultValue("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
				Default: int64default.StaticInt64(0),
			},
			"max_failed_login": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the number of failed login attempts to allow before account lockout. Enter a value in the range of 0 - 64, where 0 disables account lockout. The default value depends on common criteria mode. <ul><li>When common criteria is enabled, the default value is 3.</li><li>When common criteria is not enabled, the default value is 0.</li></ul>", "max-login-failure", "").AddIntegerRange(0, 64).AddDefaultValue("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 64),
				},
				Default: int64default.StaticInt64(0),
			},
			"lockout_period": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration to lock out local user accounts after the maximum number of failed login attempts is exceeded. Instead of locking out accounts for a specific duration, the account can be locked out until re-enabled by a privileged user. Enter a value in the range 0 - 1000, where 0 locks out accounts until reset. The default value depends on common criteria mode. <ul><li>When common criteria is enabled, the default value is 0.</li><li>When common criteria is not enabled, the default value is 1.</li></ul><p><b>Note:</b> The duration applies to all local accounts, including the <tt>admin</tt> user account. The only difference is that the <tt>admin</tt> user account cannot be locked out until reset. When the duration is 0, the <tt>admin</tt> user account is locked out for 120 minutes or until re-enabled by a privileged user.</p>", "lockout-duration", "").AddIntegerRange(0, 1000).AddDefaultValue("1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 1000),
				},
				Default: int64default.StaticInt64(1),
			},
			"mc_force_dn_ldap_order": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to convert the extracted DN to LDAP format. This property is essential when the extracted DN from a TLS certificate is in X.500 format. This format arranges the RDNs of the DNs from left to right with forward slashes as separators; for example, <tt>C=US/O=My Organization/CN=Fred</tt> . <p>When you retrieve the group name with an LDAP search, the authenticated DN must be in LDAP format. This format arranges the RDNs of the DNs from right to left with commas as separators; for example, <tt>CN=Fred, O=My Organization, C=US</tt> .</p>", "mc-force-dn-ldap-order", "").AddDefaultValue("false").AddNotValidWhen(models.RBMSettingsMCForceDNLDAPOrderIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"password_hash_algorithm": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the algorithm to apply to passwords before they are stored. The hash algorithm affects firmware downgrade and cryptographic modes in the following ways. <ul><li>For firmware downgrade, sha256crypt is not supported in releases earlier than 6.0.1.0.</li><li>For cryptographic modes, like FIPS 140-2 Level 1, md5crypt is not supported in FIPS mode.</li></ul>", "password-hash-algorithm", "").AddStringEnum("md5crypt", "sha256crypt").AddDefaultValue("md5crypt").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("md5crypt", "sha256crypt"),
				},
				Default: stringdefault.StaticString("md5crypt"),
			},
			"ldap_ssl_client_config_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client type", "ssl-client-type", "").AddStringEnum("client").AddDefaultValue("client").AddNotValidWhen(models.RBMSettingsLDAPSSLClientConfigTypeIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("client"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.RBMSettingsLDAPSSLClientConfigTypeIgnoreVal, true),
				},
				Default: stringdefault.StaticString("client"),
			},
			"ldap_ssl_client_profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "ssl_client_profile").AddNotValidWhen(models.RBMSettingsLDAPSSLClientProfileIgnoreVal.String()).String,
				Optional:            true,
			},
			"mc_ldap_ssl_client_config_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client type", "mc-ssl-client-type", "").AddStringEnum("client").AddDefaultValue("client").AddNotValidWhen(models.RBMSettingsMCLDAPSSLClientConfigTypeIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("client"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.RBMSettingsMCLDAPSSLClientConfigTypeIgnoreVal, true),
				},
				Default: stringdefault.StaticString("client"),
			},
			"mc_ldap_ssl_client_profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "mc-ssl-client", "ssl_client_profile").AddNotValidWhen(models.RBMSettingsMCLDAPSSLClientProfileIgnoreVal.String()).String,
				Optional:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *RBMSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *RBMSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.RBMSettings
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, "default", data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `RBMSettings`)
	_, err := r.pData.Client.Put(data.GetPath(), body)
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			_, err := r.pData.Client.Put(data.GetPath(), body)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Resource already exists. Failed to update resource, got error: %s", err))
				return
			}
			resp.Diagnostics.AddWarning("Warning", "Resource already exists. Existing resource was updated.")
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create resource, got error: %s", err))
			return
		}
	}
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RBMSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.RBMSettings
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.pData.Client.Get(data.GetPath())
	if err != nil && (strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}

	data.UpdateFromBody(ctx, `RBMSettings`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RBMSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.RBMSettings
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, "default", data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath(), data.ToBody(ctx, `RBMSettings`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Update)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RBMSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.RBMSettings
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, "default", data.DependencyActions, actions.Delete, false)
	if resp.Diagnostics.HasError() {
		return
	}
	data.ToDefault()
	_, err := r.pData.Client.Put(data.GetPath(), data.ToBody(ctx, `RBMSettings`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to restore singleton to default, got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *RBMSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()
	appDomain := req.ID
	if appDomain != "default" {
		resp.Diagnostics.AddError("Invalid Application Domain", "This resourece supported on the 'default' domain only.")
		return
	}
	if !regexp.MustCompile("^[a-zA-Z0-9_-]+$").MatchString(appDomain) || len(appDomain) < 1 || len(appDomain) > 128 {
		resp.Diagnostics.AddError("Invalid Application Domain", "Application domain must be 1-128 characters and match pattern ^[a-zA-Z0-9_-]+$")
		return
	}

	var data models.RBMSettings
	res, err := r.pData.Client.Get(data.GetPath())
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("Resource was not found, got error: %s", err))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		}
		return
	}

	data.FromBody(ctx, `RBMSettings`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RBMSettingsResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.RBMSettings

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
