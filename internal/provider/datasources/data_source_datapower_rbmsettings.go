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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
)

var (
	_ datasource.DataSource              = &RBMSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &RBMSettingsDataSource{}
)

func NewRBMSettingsDataSource() datasource.DataSource {
	return &RBMSettingsDataSource{}
}

type RBMSettingsDataSource struct {
	client *client.DatapowerClient
}

func (d *RBMSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rbmsettings"
}

func (d *RBMSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "RBM settings (`default` domain only)",
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Administrative state",
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
			"sshau_method": models.GetDmRBMSSHAuthenticateTypeDataSourceSchema("SSH authentication method", "ssh-au-method", ""),
			"ca_pub_key_file": schema.StringAttribute{
				MarkdownDescription: "CA user public key file",
				Computed:            true,
			},
			"revoked_keys": schema.ListAttribute{
				MarkdownDescription: "Revoked keys",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"auzosnss_config": schema.StringAttribute{
				MarkdownDescription: "z/OS NSS client",
				Computed:            true,
			},
			"auoidc_scope": schema.StringAttribute{
				MarkdownDescription: "Scope",
				Computed:            true,
			},
			"auoidc_client_id": schema.StringAttribute{
				MarkdownDescription: "Client ID",
				Computed:            true,
			},
			"auoidc_client_secret": schema.StringAttribute{
				MarkdownDescription: "Client secret",
				Computed:            true,
			},
			"auoidc_identity_service_url": schema.StringAttribute{
				MarkdownDescription: "Identity service URL",
				Computed:            true,
			},
			"auoidc_key_fetch_interval": schema.Int64Attribute{
				MarkdownDescription: "Retrieval interval for public keys.",
				Computed:            true,
			},
			"auoidc_identity_service_urlssl_client": schema.StringAttribute{
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
				MarkdownDescription: "XML file URL",
				Computed:            true,
			},
			"aussl_valcred": schema.StringAttribute{
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
			"auldap_search_for_dn": schema.BoolAttribute{
				MarkdownDescription: "Search LDAP for DN",
				Computed:            true,
			},
			"auldap_bind_dn": schema.StringAttribute{
				MarkdownDescription: "LDAP bind DN",
				Computed:            true,
			},
			"auldap_bind_password_alias": schema.StringAttribute{
				MarkdownDescription: "LDAP bind password alias",
				Computed:            true,
			},
			"auldap_search_parameters": schema.StringAttribute{
				MarkdownDescription: "LDAP search parameters",
				Computed:            true,
			},
			"auldap_prefix": schema.StringAttribute{
				MarkdownDescription: "LDAP prefix",
				Computed:            true,
			},
			"au_force_dnldap_order": schema.BoolAttribute{
				MarkdownDescription: "Convert DN to LDAP format",
				Computed:            true,
			},
			"lda_psuffix": schema.StringAttribute{
				MarkdownDescription: "LDAP suffix",
				Computed:            true,
			},
			"auldap_load_balance_group": schema.StringAttribute{
				MarkdownDescription: "Load balancer group",
				Computed:            true,
			},
			"au_cache_allow": schema.StringAttribute{
				MarkdownDescription: "Authentication cache mode",
				Computed:            true,
			},
			"au_cache_ttl": schema.Int64Attribute{
				MarkdownDescription: "Authentication cache lifetime",
				Computed:            true,
			},
			"auldap_read_timeout": schema.Int64Attribute{
				MarkdownDescription: "LDAP read timeout",
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
			"mcldap_search_for_group": schema.BoolAttribute{
				MarkdownDescription: "Search LDAP for group name",
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
			"mcldap_load_balance_group": schema.StringAttribute{
				MarkdownDescription: "Load balancer group",
				Computed:            true,
			},
			"mcldap_bind_dn": schema.StringAttribute{
				MarkdownDescription: "LDAP bind DN",
				Computed:            true,
			},
			"mcldap_bind_password_alias": schema.StringAttribute{
				MarkdownDescription: "LDAP bind password alias",
				Computed:            true,
			},
			"mcldap_search_parameters": schema.StringAttribute{
				MarkdownDescription: "LDAP search parameters",
				Computed:            true,
			},
			"mc_info_url": schema.StringAttribute{
				MarkdownDescription: "XML file URL",
				Computed:            true,
			},
			"mcldap_read_timeout": schema.Int64Attribute{
				MarkdownDescription: "LDAP read timeout",
				Computed:            true,
			},
			"ldap_version": schema.StringAttribute{
				MarkdownDescription: "LDAP version",
				Computed:            true,
			},
			"fallback_login": schema.StringAttribute{
				MarkdownDescription: "Local accounts for fallback",
				Computed:            true,
			},
			"fallback_user": schema.ListAttribute{
				MarkdownDescription: "Fallback users",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"apply_to_cli": schema.BoolAttribute{
				MarkdownDescription: "Enforce RBM on CLI",
				Computed:            true,
			},
			"restrict_admin_to_serial_port": schema.BoolAttribute{
				MarkdownDescription: "Restrict admin to serial",
				Computed:            true,
			},
			"min_password_length": schema.Int64Attribute{
				MarkdownDescription: "Minimum length",
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
				MarkdownDescription: "Max age",
				Computed:            true,
			},
			"do_password_history": schema.BoolAttribute{
				MarkdownDescription: "Control reuse",
				Computed:            true,
			},
			"num_old_passwords": schema.Int64Attribute{
				MarkdownDescription: "Reuse history",
				Computed:            true,
			},
			"cli_timeout": schema.Int64Attribute{
				MarkdownDescription: "CLI idle timeout",
				Computed:            true,
			},
			"max_failed_login": schema.Int64Attribute{
				MarkdownDescription: "Max failed logins",
				Computed:            true,
			},
			"lockout_period": schema.Int64Attribute{
				MarkdownDescription: "Lockout duration",
				Computed:            true,
			},
			"mc_force_dnldap_order": schema.BoolAttribute{
				MarkdownDescription: "Convert DN to LDAP format",
				Computed:            true,
			},
			"password_hash_algorithm": schema.StringAttribute{
				MarkdownDescription: "Password hash algorithm",
				Computed:            true,
			},
			"ldapssl_client_config_type": schema.StringAttribute{
				MarkdownDescription: "TLS client type",
				Computed:            true,
			},
			"ldapssl_client_profile": schema.StringAttribute{
				MarkdownDescription: "TLS client profile",
				Computed:            true,
			},
			"mcldapssl_client_config_type": schema.StringAttribute{
				MarkdownDescription: "TLS client type",
				Computed:            true,
			},
			"mcldapssl_client_profile": schema.StringAttribute{
				MarkdownDescription: "TLS client profile",
				Computed:            true,
			},
		},
	}
}

func (d *RBMSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *RBMSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.RBMSettings

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := d.client.Get(data.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	data.FromBody(ctx, `RBMSettings`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
