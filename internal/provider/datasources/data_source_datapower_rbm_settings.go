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

package datasources

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var (
	_ datasource.DataSource              = &RBMSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &RBMSettingsDataSource{}
)

func NewRBMSettingsDataSource() datasource.DataSource {
	return &RBMSettingsDataSource{}
}

type RBMSettingsDataSource struct {
	pData *tfutils.ProviderData
}

func (d *RBMSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rbm_settings"
}

func (d *RBMSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "<p>Manage role-based management (RBM) settings: RBM, password policy, and account policy</p><ul><li>RBM consists of the following capabilities: Authenticating users, evaluating the access profile, enforcing access to resources</li><li>The password policy sets the password requirements for local user accounts.</li><li>The account policy sets the lockout behavior and the timeout for CLI sessions.</li></ul>",
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "<p>The administrative state of the configuration.</p><ul><li>To make active, set to enabled.</li><li>To make inactive, set to disabled.</li></ul>",
				Computed:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: "Comments",
				Computed:            true,
			},
			"au_method": schema.StringAttribute{
				MarkdownDescription: "Authentication method",
				Computed:            true,
			},
			"ssh_au_method": models.GetDmRBMSSHAuthenticateTypeDataSourceSchema("Specify the method to authenticate SSH users. <ul><li>When no method, the user is prompted for both username and password.</li><li>When password, the user is prompted for the password. For this method, the username must be part of the invocation. With the ssh command, the invocation is in the ssh username@host format.</li><li>When user certificate, the user is not prompted for input. The connection is successful when the invocation provides a signed SSH user certificate that is verified by the CA public key file in the <tt>cert:</tt> directory. With the ssh command, the invocation must include the -i file parameter.</li><li>When both certificate and password, processing attempts to first authenticate with the provided signed SSH user certificate. If unsuccessful, prompts for the password.</li><li>Supported RBM authentication methods with SSH authentication method are local and LDAP. <ul><li>Local authentication method extracts the certificate identity and attempts login with local User of that name.</li><li>LDAP authentication method constructs the DN through LDAP search or by applying the configured prefix and suffix. Since SSH authentication completes prior to the Authentication step no LDAP bind or authenticate will take place.</li></ul></li></ul>", "ssh-au-method", ""),
			"ca_pub_key_file": schema.StringAttribute{
				MarkdownDescription: "Specify the certificate authority (CA) public key file in the <tt>cert:</tt> directory for SSH authentication with SSH user certificates. This public key file contains the public key for one or more certificate authorities.",
				Computed:            true,
			},
			"revoked_keys": schema.ListAttribute{
				MarkdownDescription: "Specify the OpenSSH public keys to revoke for SSH authentication. Each entry is the public key file in the <tt>cert:</tt> or <tt>sharedcert:</tt> directory and must be in the OpenSSH public key format. These keys are signed by the CA user public key file.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"au_zos_nss_config": schema.StringAttribute{
				MarkdownDescription: "z/OS NSS client",
				Computed:            true,
			},
			"au_oidc_scope": schema.StringAttribute{
				MarkdownDescription: "Scope",
				Computed:            true,
			},
			"au_oidc_client_id": schema.StringAttribute{
				MarkdownDescription: "Client ID",
				Computed:            true,
			},
			"au_oidc_client_secret": schema.StringAttribute{
				MarkdownDescription: "Client secret",
				Computed:            true,
			},
			"au_oidc_identity_service_url": schema.StringAttribute{
				MarkdownDescription: "Identity service URL",
				Computed:            true,
			},
			"au_oidc_key_fetch_interval": schema.Int64Attribute{
				MarkdownDescription: "Retrieval interval for public keys.",
				Computed:            true,
			},
			"au_oidc_identity_service_url_ssl_client": schema.StringAttribute{
				MarkdownDescription: "TLS client profile",
				Computed:            true,
			},
			"au_kerberos_keytab": schema.StringAttribute{
				MarkdownDescription: "Kerberos keytab",
				Computed:            true,
			},
			"au_custom_url": schema.StringAttribute{
				MarkdownDescription: "Custom URL",
				Computed:            true,
			},
			"au_info_url": schema.StringAttribute{
				MarkdownDescription: "Specify the URL of the XML file for authentication. The XML file can be on the DataPower Gateway or on a remote server. You can use the same XML file to map credentials.",
				Computed:            true,
			},
			"au_ssl_valcred": schema.StringAttribute{
				MarkdownDescription: "Validation credentials",
				Computed:            true,
			},
			"au_host": schema.StringAttribute{
				MarkdownDescription: "Server host",
				Computed:            true,
			},
			"au_port": schema.Int64Attribute{
				MarkdownDescription: "Server port",
				Computed:            true,
			},
			"au_ldap_search_for_dn": schema.BoolAttribute{
				MarkdownDescription: "Specify whether to retrieve the user DN with an LDAP search. <ul><li>When enabled, the login name presented by the user is used with the LDAP search parameters for an LDAP search to retrieve the user DN.</li><li>When disabled, the login name presented by the user is used with the LDAP prefix and suffix to construct the user DN.</li></ul>",
				Computed:            true,
			},
			"au_ldap_bind_dn": schema.StringAttribute{
				MarkdownDescription: "LDAP bind DN",
				Computed:            true,
			},
			"au_ldap_bind_password_alias": schema.StringAttribute{
				MarkdownDescription: "LDAP bind password alias",
				Computed:            true,
			},
			"au_ldap_search_parameters": schema.StringAttribute{
				MarkdownDescription: "LDAP search parameters",
				Computed:            true,
			},
			"au_ldap_prefix": schema.StringAttribute{
				MarkdownDescription: "Specify the string to add before the username to form the DN. If this value is <tt>CN=</tt> and the username is <tt>Bob</tt> , the complete DN is <tt>CN=Bob,O=example.com</tt> when the LDAP suffix is <tt>O=example.com</tt> .",
				Computed:            true,
			},
			"au_force_dn_ldap_order": schema.BoolAttribute{
				MarkdownDescription: "Specify whether to convert the extracted DN to LDAP format. This property is essential when the extracted DN from a TLS certificate is in X.500 format. This format arranges the RDNs of the DNs from left to right with forward slashes as separators; for example, <tt>C=US/O=My Organization/CN=Fred</tt> . <p>When you retrieve the group name with an LDAP search, the authenticated DN must be in LDAP format. This format arranges the RDNs of the DNs from right to left with commas as separators; for example, <tt>CN=Fred, O=My Organization, C=US</tt> .</p>",
				Computed:            true,
			},
			"ldap_suffix": schema.StringAttribute{
				MarkdownDescription: "Specify the string to add after the username to form the DN. If this value is <tt>O=example.com</tt> and the username is <tt>Bob</tt> , the complete DN is <tt>CN=Bob,O=example.com</tt> when the LDAP prefix is <tt>CN=</tt> .",
				Computed:            true,
			},
			"au_ldap_load_balance_group": schema.StringAttribute{
				MarkdownDescription: "Specify the load balancer group of LDAP servers. This setting overrides the settings for the server host and port.",
				Computed:            true,
			},
			"au_cache_allow": schema.StringAttribute{
				MarkdownDescription: "Authentication cache mode",
				Computed:            true,
			},
			"au_cache_ttl": schema.Int64Attribute{
				MarkdownDescription: "Specify the expiry for cached authentication decisions. Enter a value in the range 1 - 86400. The default value is 600.",
				Computed:            true,
			},
			"au_ldap_read_timeout": schema.Int64Attribute{
				MarkdownDescription: "Specify the time to wait for a response from the LDAP server before the DataPower Gateway closes the LDAP connection. Enter a value in the range 0 - 86400. The default value is 60. A value of 0 indicates that the connection never times out.",
				Computed:            true,
			},
			"mc_method": schema.StringAttribute{
				MarkdownDescription: "Credential-mapping method",
				Computed:            true,
			},
			"mc_custom_url": schema.StringAttribute{
				MarkdownDescription: "Custom URL",
				Computed:            true,
			},
			"mc_ldap_search_for_group": schema.BoolAttribute{
				MarkdownDescription: "Specify whether to search LDAP to retrieve all user groups that match the query. <ul><li>When enabled, the authenticated DN of the user and the LDAP search parameters are used as part of the LDAP search to retrieve all user groups that match the query. When a user belongs to multiple groups, the resultant access policy for this user is additive not most restrictive.</li><li>When disabled, the authenticated identity of the user (DN or user group of local user) is used directly as the input credential.</li></ul>",
				Computed:            true,
			},
			"mc_host": schema.StringAttribute{
				MarkdownDescription: "Server host",
				Computed:            true,
			},
			"mc_port": schema.Int64Attribute{
				MarkdownDescription: "Server port",
				Computed:            true,
			},
			"mc_ldap_load_balance_group": schema.StringAttribute{
				MarkdownDescription: "Load balancer group",
				Computed:            true,
			},
			"mc_ldap_bind_dn": schema.StringAttribute{
				MarkdownDescription: "LDAP bind DN",
				Computed:            true,
			},
			"mc_ldap_bind_password_alias": schema.StringAttribute{
				MarkdownDescription: "LDAP bind password alias",
				Computed:            true,
			},
			"mc_ldap_search_parameters": schema.StringAttribute{
				MarkdownDescription: "LDAP search parameters",
				Computed:            true,
			},
			"mc_info_url": schema.StringAttribute{
				MarkdownDescription: "Specify the URL of the XML file to map credentials. The XML file can be on the DataPower Gateway or on a remote server. You can use the same XML file for authentication.",
				Computed:            true,
			},
			"mc_ldap_read_timeout": schema.Int64Attribute{
				MarkdownDescription: "Specify the time to wait for a response from the LDAP server before the DataPower Gateway closes the LDAP connection. Enter a value in the range 0 - 86400. The default value is 60. A value of 0 indicates that the connection never times out.",
				Computed:            true,
			},
			"ldap_version": schema.StringAttribute{
				MarkdownDescription: "LDAP version",
				Computed:            true,
			},
			"fallback_login": schema.StringAttribute{
				MarkdownDescription: "Specify whether to use local user accounts as fallback users if remote authentication fails. With fallback users, local user accounts can log on to the DataPower Gateway if authentication fails or during a network outage that affects primary authentication. The recommendation is to restrict fallback users to a subset of local user accounts. <p><b>Note: </b>When authentication uses a TLS certificate from a connection peer, you cannot enforce RBM on CLI sessions unless fallback users are supported.</p>",
				Computed:            true,
			},
			"fallback_user": schema.ListAttribute{
				MarkdownDescription: "Fallback users",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"apply_to_cli": schema.BoolAttribute{
				MarkdownDescription: "Specify whether to enforce the RBM policy on CLI sessions. When authentication uses a TLS certificate from a connection peer, you cannot enforce RBM on CLI sessions unless fallback users are supported.",
				Computed:            true,
			},
			"restrict_admin_to_serial_port": schema.BoolAttribute{
				MarkdownDescription: "Restrict admin to serial",
				Computed:            true,
			},
			"min_password_length": schema.Int64Attribute{
				MarkdownDescription: "Specify the minimum length of a valid password. Enter a value in the range 1 - 128. The default value depends on common criteria mode. <ul><li>When common criteria is enabled, the default value is 14.</li><li>When common criteria is not enabled, the default value is 6.</li></ul>",
				Computed:            true,
			},
			"require_mixed_case": schema.BoolAttribute{
				MarkdownDescription: "Require mixed case",
				Computed:            true,
			},
			"require_digit": schema.BoolAttribute{
				MarkdownDescription: "Require number",
				Computed:            true,
			},
			"require_non_alpha_numeric": schema.BoolAttribute{
				MarkdownDescription: "Require non-alphanumeric",
				Computed:            true,
			},
			"disallow_username_substring": schema.BoolAttribute{
				MarkdownDescription: "Disallow username substring",
				Computed:            true,
			},
			"do_password_aging": schema.BoolAttribute{
				MarkdownDescription: "Enable aging",
				Computed:            true,
			},
			"max_password_age": schema.Int64Attribute{
				MarkdownDescription: "Specify the expiry for passwords. The default value depends on common criteria mode. <ul><li>When common criteria is enabled, the default value is 90.</li><li>When common criteria is not enabled, the default value is 30.</li></ul>",
				Computed:            true,
			},
			"do_password_history": schema.BoolAttribute{
				MarkdownDescription: "Control reuse",
				Computed:            true,
			},
			"num_old_passwords": schema.Int64Attribute{
				MarkdownDescription: "Specify the number of recent passwords to track to prevent reuse. The default value depends on common criteria mode. <ul><li>When common criteria is enabled, the default value is 3.</li><li>When common criteria is not enabled, the default value is 5.</li></ul>",
				Computed:            true,
			},
			"cli_timeout": schema.Int64Attribute{
				MarkdownDescription: "Specify the time after which to invalidate idle CLI sessions. When invalidated, requires re-authentication. Enter a value in the range 0 - 65535. A value of 0 disables the timer. The default value depends on common criteria mode. <ul><li>When common criteria is enabled, the default value is 900.</li><li>When common criteria is not enabled, the default value is 0.</li></ul>",
				Computed:            true,
			},
			"max_failed_login": schema.Int64Attribute{
				MarkdownDescription: "Specify the number of failed login attempts to allow before account lockout. Enter a value in the range of 0 - 64, where 0 disables account lockout. The default value depends on common criteria mode. <ul><li>When common criteria is enabled, the default value is 3.</li><li>When common criteria is not enabled, the default value is 0.</li></ul>",
				Computed:            true,
			},
			"lockout_period": schema.Int64Attribute{
				MarkdownDescription: "Specify the duration to lock out local user accounts after the maximum number of failed login attempts is exceeded. Instead of locking out accounts for a specific duration, the account can be locked out until re-enabled by a privileged user. Enter a value in the range 0 - 1000, where 0 locks out accounts until reset. The default value depends on common criteria mode. <ul><li>When common criteria is enabled, the default value is 0.</li><li>When common criteria is not enabled, the default value is 1.</li></ul><p><b>Note:</b> The duration applies to all local accounts, including the <tt>admin</tt> user account. The only difference is that the <tt>admin</tt> user account cannot be locked out until reset. When the duration is 0, the <tt>admin</tt> user account is locked out for 120 minutes or until re-enabled by a privileged user.</p>",
				Computed:            true,
			},
			"mc_force_dn_ldap_order": schema.BoolAttribute{
				MarkdownDescription: "Specify whether to convert the extracted DN to LDAP format. This property is essential when the extracted DN from a TLS certificate is in X.500 format. This format arranges the RDNs of the DNs from left to right with forward slashes as separators; for example, <tt>C=US/O=My Organization/CN=Fred</tt> . <p>When you retrieve the group name with an LDAP search, the authenticated DN must be in LDAP format. This format arranges the RDNs of the DNs from right to left with commas as separators; for example, <tt>CN=Fred, O=My Organization, C=US</tt> .</p>",
				Computed:            true,
			},
			"password_hash_algorithm": schema.StringAttribute{
				MarkdownDescription: "Specify the algorithm to apply to passwords before they are stored. The hash algorithm affects firmware downgrade and cryptographic modes in the following ways. <ul><li>For firmware downgrade, sha256crypt is not supported in releases earlier than 6.0.1.0.</li><li>For cryptographic modes, like FIPS 140-2 Level 1, md5crypt is not supported in FIPS mode.</li></ul>",
				Computed:            true,
			},
			"ldap_ssl_client_config_type": schema.StringAttribute{
				MarkdownDescription: "TLS client type",
				Computed:            true,
			},
			"ldap_ssl_client_profile": schema.StringAttribute{
				MarkdownDescription: "TLS client profile",
				Computed:            true,
			},
			"mc_ldap_ssl_client_config_type": schema.StringAttribute{
				MarkdownDescription: "TLS client type",
				Computed:            true,
			},
			"mc_ldap_ssl_client_profile": schema.StringAttribute{
				MarkdownDescription: "TLS client profile",
				Computed:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (d *RBMSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *RBMSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data models.RBMSettings
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	path := data.GetPath()

	res, err := d.pData.Client.Get(path)
	resFound := true
	if err != nil {
		if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		} else {
			resFound = false
		}
	}
	if resFound {
		data.FromBody(ctx, `RBMSettings`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
