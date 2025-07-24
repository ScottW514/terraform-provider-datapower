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

package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &RBMSettingsResource{}

func NewRBMSettingsResource() resource.Resource {
	return &RBMSettingsResource{}
}

type RBMSettingsResource struct {
	client *client.DatapowerClient
}

func (r *RBMSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rbmsettings"
}

func (r *RBMSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("RBM settings (`default` domain only)", "rbm", "").String,

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Administrative state", "admin-state", "").AddDefaultValue("true").String,
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
			"sshau_method": models.GetDmRBMSSHAuthenticateTypeResourceSchema("SSH authentication method", "ssh-au-method", "", false),
			"ca_pub_key_file": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("CA user public key file", "ssh-ca-pubkey-file", "").String,
				Optional:            true,
			},
			"revoked_keys": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Revoked keys", "ssh-revoked-keys", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"auzosnss_config": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("z/OS NSS client", "zos-nss-au", "zosnssclient").String,
				Optional:            true,
			},
			"auoidc_scope": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Scope", "au-oidc-scope", "").AddDefaultValue("openid").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("openid"),
			},
			"auoidc_client_id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Client ID", "au-oidc-client-id", "").String,
				Optional:            true,
			},
			"auoidc_client_secret": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Client secret", "au-oidc-client-secret", "passwordalias").String,
				Optional:            true,
			},
			"auoidc_identity_service_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Identity service URL", "au-oidc-identity-service-url", "").String,
				Optional:            true,
			},
			"auoidc_key_fetch_interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Retrieval interval for public keys.", "au-oidc-key-fetch-interval", "").AddIntegerRange(1, 86400).AddDefaultValue("30").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(30),
			},
			"auoidc_identity_service_urlssl_client": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "au-oidc-identity-service-ssl-client", "sslclientprofile").String,
				Optional:            true,
			},
			"au_kerberos_keytab": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Kerberos keytab", "au-kerberos-keytab", "cryptokerberoskeytab").String,
				Optional:            true,
			},
			"au_custom_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Custom URL", "au-custom-url", "").String,
				Optional:            true,
			},
			"au_info_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML file URL", "au-info-url", "").String,
				Optional:            true,
			},
			"aussl_valcred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validation credentials", "au-valcred", "cryptovalcred").String,
				Optional:            true,
			},
			"au_host": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Server host", "au-server-host", "").String,
				Optional:            true,
			},
			"au_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Server port", "au-server-port", "").String,
				Optional:            true,
			},
			"auldap_search_for_dn": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Search LDAP for DN", "au-ldap-search", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"auldap_bind_dn": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP bind DN", "au-ldap-bind-dn", "").String,
				Optional:            true,
			},
			"auldap_bind_password_alias": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP bind password alias", "au-ldap-bind-password-alias", "passwordalias").String,
				Optional:            true,
			},
			"auldap_search_parameters": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP search parameters", "au-ldap-parameters", "ldapsearchparameters").String,
				Optional:            true,
			},
			"auldap_prefix": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP prefix", "ldap-prefix", "").AddDefaultValue("cn=").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("cn="),
			},
			"au_force_dnldap_order": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Convert DN to LDAP format", "au-force-dn-ldap-order", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"lda_psuffix": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP suffix", "ldap-suffix", "").String,
				Optional:            true,
			},
			"auldap_load_balance_group": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Load balancer group", "loadbalancer-group", "loadbalancergroup").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Authentication cache lifetime", "au-cache-ttl", "").AddIntegerRange(1, 86400).AddDefaultValue("600").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(600),
			},
			"auldap_read_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP read timeout", "au-ldap-readtimeout", "").AddIntegerRange(0, 86400).AddDefaultValue("60").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 86400),
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
				MarkdownDescription: tfutils.NewAttributeDescription("Custom URL", "mc-custom-url", "").String,
				Optional:            true,
			},
			"mcldap_search_for_group": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Search LDAP for group name", "mc-ldap-search", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"mc_host": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Server host", "mc-server-host", "").String,
				Optional:            true,
			},
			"mc_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Server port", "mc-server-port", "").String,
				Optional:            true,
			},
			"mcldap_load_balance_group": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Load balancer group", "mc-loadbalancer-group", "loadbalancergroup").String,
				Optional:            true,
			},
			"mcldap_bind_dn": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP bind DN", "mc-ldap-bind-dn", "").String,
				Optional:            true,
			},
			"mcldap_bind_password_alias": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP bind password alias", "mc-ldap-bind-password-alias", "passwordalias").String,
				Optional:            true,
			},
			"mcldap_search_parameters": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP search parameters", "mc-ldap-parameters", "ldapsearchparameters").String,
				Optional:            true,
			},
			"mc_info_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML file URL", "mc-info-url", "").String,
				Optional:            true,
			},
			"mcldap_read_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP read timeout", "mc-ldap-readtimeout", "").AddIntegerRange(0, 86400).AddDefaultValue("60").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 86400),
				},
				Default: int64default.StaticInt64(60),
			},
			"ldap_version": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP version", "ldap-version", "").AddStringEnum("v2", "v3").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("v2", "v3"),
				},
			},
			"fallback_login": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Local accounts for fallback", "fallback-login", "").AddStringEnum("disabled", "local", "restricted").AddDefaultValue("disabled").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disabled", "local", "restricted"),
				},
				Default: stringdefault.StaticString("disabled"),
			},
			"fallback_user": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Fallback users", "fallback-user", "user").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"apply_to_cli": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enforce RBM on CLI", "apply-cli", "").AddDefaultValue("false").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Minimum length", "pwd-minimum-length", "").AddIntegerRange(1, 128).AddDefaultValue("6").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Max age", "pwd-max-age", "").AddIntegerRange(1, 65535).AddDefaultValue("30").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 65535),
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
				MarkdownDescription: tfutils.NewAttributeDescription("Reuse history", "pwd-max-history", "").AddIntegerRange(1, 65535).AddDefaultValue("5").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(5),
			},
			"cli_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("CLI idle timeout", "cli-timeout", "").AddIntegerRange(0, 65535).AddDefaultValue("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 65535),
				},
				Default: int64default.StaticInt64(0),
			},
			"max_failed_login": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Max failed logins", "max-login-failure", "").AddIntegerRange(0, 64).AddDefaultValue("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 64),
				},
				Default: int64default.StaticInt64(0),
			},
			"lockout_period": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Lockout duration", "lockout-duration", "").AddIntegerRange(0, 1000).AddDefaultValue("1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 1000),
				},
				Default: int64default.StaticInt64(1),
			},
			"mc_force_dnldap_order": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Convert DN to LDAP format", "mc-force-dn-ldap-order", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"password_hash_algorithm": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Password hash algorithm", "password-hash-algorithm", "").AddStringEnum("md5crypt", "sha256crypt").AddDefaultValue("md5crypt").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("md5crypt", "sha256crypt"),
				},
				Default: stringdefault.StaticString("md5crypt"),
			},
			"ldapssl_client_config_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client type", "ssl-client-type", "").AddStringEnum("client").AddDefaultValue("client").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("client"),
				},
				Default: stringdefault.StaticString("client"),
			},
			"ldapssl_client_profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "sslclientprofile").String,
				Optional:            true,
			},
			"mcldapssl_client_config_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client type", "mc-ssl-client-type", "").AddStringEnum("client").AddDefaultValue("client").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("client"),
				},
				Default: stringdefault.StaticString("client"),
			},
			"mcldapssl_client_profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "mc-ssl-client", "sslclientprofile").String,
				Optional:            true,
			},
		},
	}
}

func (r *RBMSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *RBMSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.RBMSettings

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `RBMSettings`)
	_, err := r.client.Put(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "PUT", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RBMSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.RBMSettings

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Get(data.GetPath())
	if err != nil && (strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}

	if data.IsNull() {
		// Import
		data.FromBody(ctx, `RBMSettings`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `RBMSettings`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RBMSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.RBMSettings

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Put(data.GetPath(), data.ToBody(ctx, `RBMSettings`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RBMSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	resp.State.RemoveResource(ctx)
}
