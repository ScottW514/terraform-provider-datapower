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
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &OAuthProviderSettingsResource{}

func NewOAuthProviderSettingsResource() resource.Resource {
	return &OAuthProviderSettingsResource{}
}

type OAuthProviderSettingsResource struct {
	client *client.DatapowerClient
}

func (r *OAuthProviderSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_oauthprovidersettings"
}

func (r *OAuthProviderSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("OAuth provider settings", "oauth-provider-settings", "").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name of the object. Must be unique among object types in application domain.", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"app_domain": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The name of the application domain the object belongs to", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`), ""),
				},
				PlanModifiers: []planmodifier.String{
					modifiers.ImmutableAfterSet(),
				},
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"enable_debug_mode": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable debug headers", "enable-debug-mode", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"provider_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Provider type", "provider-type", "").AddStringEnum("native", "third_party").AddDefaultValue("native").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("native", "third_party"),
				},
				Default: stringdefault.StaticString("native"),
			},
			"scopes_allowed": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allowed scopes", "scopes-allowed", "").String,
				Optional:            true,
			},
			"default_scopes": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Default scopes", "default-scopes", "").String,
				Optional:            true,
			},
			"supported_grant_types":  models.GetDmOAuthProviderGrantTypeResourceSchema("Supported grant types", "supported-grant-types", "", false),
			"supported_client_types": models.GetDmAllowedClientTypeResourceSchema("Supported client types", "supported-client-types", "", false),
			"apic_provider_base_path": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Provider base path", "apic-provider-base-path", "").AddDefaultValue("/").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/"),
			},
			"apic_authorize_endpoint": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Authorization endpoint", "apic-authorize-endpoint", "").AddDefaultValue("/oauth2/authorize").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/oauth2/authorize"),
			},
			"apic_token_endpoint": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Token endpoint", "apic-token-endpoint", "").AddDefaultValue("/oauth2/token").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/oauth2/token"),
			},
			"apic_enable_introspection": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable token introspection", "apic-enable-introspection", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"apic_introspect_endpoint": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Introspection endpoint", "apic-introspect-endpoint", "").AddDefaultValue("/oauth2/introspect").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/oauth2/introspect"),
			},
			"apic_token_secret": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Token secret", "apic-token-secret", "cryptosskey").String,
				Optional:            true,
			},
			"apic_one_time_use_accesstoken": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("One-time use access token", "apic-enable-one-time-use-access-token", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"apic_access_token_ttl": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Access token TTL", "apic-access-token-ttl", "").AddIntegerRange(1, 63244800).AddDefaultValue("3600").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 63244800),
				},
				Default: int64default.StaticInt64(3600),
			},
			"apic_auth_code_ttl": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Authorization code TTL", "apic-auth-code-ttl", "").AddIntegerRange(1, 600).AddDefaultValue("300").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 600),
				},
				Default: int64default.StaticInt64(300),
			},
			"apic_enable_refresh_token": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable refresh tokens", "apic-enable-refresh-token", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"apic_one_time_use_refreshtoken": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("One-time use refresh token", "apic-enable-one-time-use-refresh-token", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"apic_refresh_token_limit": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("APIC refresh token count", "apic-refresh-token-limit", "").AddIntegerRange(1, 4096).AddDefaultValue("10").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 4096),
				},
				Default: int64default.StaticInt64(10),
			},
			"apic_refresh_token_ttl": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Refresh token TTL", "apic-refresh-token-ttl", "").AddIntegerRange(2, 252979200).AddDefaultValue("5400").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(2, 252979200),
				},
				Default: int64default.StaticInt64(5400),
			},
			"apic_maximum_consent_ttl": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Consent TTL", "apic-maximum-consent-ttl", "").AddIntegerRange(0, 2529792000).AddDefaultValue("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 2529792000),
				},
				Default: int64default.StaticInt64(0),
			},
			"adv_scope_validation_enabled": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable advanced scope validation", "advanced-scope-validation-enabled", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"advanced_scope_url_override": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use URL from API Security Definition", "advanced-scope-url-from-security", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"advanced_scope_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Advanced scope URL", "advanced-scope-url", "").String,
				Optional:            true,
			},
			"adv_scope_tls_profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS profile to access advanced scope URL", "advanced-scope-tls-profile", "sslclientprofile").String,
				Optional:            true,
			},
			"adv_scope_url_security_enabled": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable advanced scope endpoint security", "advanced-scope-url-security-enabled", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"adv_scope_url_security": models.GetDmSecurityTypeResourceSchema("Advanced scope endpoint security", "advanced-scope-url-security", "", false),
			"adv_scope_basic_auth_user_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Basic authentication username", "advanced-scope-basic-auth-username", "").String,
				Optional:            true,
			},
			"adv_scope_basic_auth_password": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Basic authentication password", "advanced-scope-basic-auth-password-alias", "passwordalias").String,
				Optional:            true,
			},
			"adv_scope_basic_auth_header_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Basic authentication header name", "advanced-scope-basic-auth-headername", "").String,
				Optional:            true,
			},
			"advanced_scope_custom_headers": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Request headers", "advanced-scope-request-headers", "").String,
				Optional:            true,
			},
			"advanced_scope_custom_contexts": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Response context variables", "advanced-scope-response-contexts", "").String,
				Optional:            true,
			},
			"apic_enable_oidc": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable OIDC", "apic-enable-oidc", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"apicoidc_hybrid_response_types": models.GetDmOIDCHybridResponseTypeResourceSchema("OIDC Hybrid Flow Response Types", "apic-oidc-hybrid-response-types", "", false),
			"apic_support_pkce": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Support PKCE", "apic-support-pkce", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"apic_require_pkce": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Require PKCE", "apic-require-pkce", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"apic_support_pkce_plain": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Support PKCE 'plain' challenge method", "apic-support-pkce-plain", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"apic_token_type_to_generate": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Type of token to generate", "apic-token-type-to-generate", "").AddStringEnum("Bearer", "jwt").AddDefaultValue("Bearer").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Bearer", "jwt"),
				},
				Default: stringdefault.StaticString("Bearer"),
			},
			"metadata_from": models.GetDmMetadataFromTypeResourceSchema("Obtain metadata from", "metadata-from", "", false),
			"metadata_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("External metadata URL", "metadata-url", "").String,
				Optional:            true,
			},
			"metadata_ssl_profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS profile to access metadata URL", "metadata-ssl-profile", "sslclientprofile").String,
				Optional:            true,
			},
			"metadata_header_for_access_token": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Response header for access token", "mdheader-for-accesstoken", "").AddDefaultValue("X-API-OAuth-Metadata-For-AccessToken").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("X-API-OAuth-Metadata-For-AccessToken"),
			},
			"metadata_header_for_payload": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Response header for payload", "mdheader-for-payload", "").AddDefaultValue("X-API-OAuth-Metadata-For-Payload").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("X-API-OAuth-Metadata-For-Payload"),
			},
			"external_revocation_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("External management URL", "external-revocation-url", "").String,
				Optional:            true,
			},
			"external_revocation_ssl_profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("External management TLS client profile", "external-revocation-ssl-profile", "sslclientprofile").String,
				Optional:            true,
			},
			"external_revocation_url_security": models.GetDmSecurityTypeResourceSchema("External management security", "external-revocation-url-security", "", false),
			"external_revocation_basic_auth_user_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Basic authentication username", "external-revocation-basic-auth-username", "").String,
				Optional:            true,
			},
			"external_revocation_basic_auth_password": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Basic authentication password", "external-revocation-basic-auth-password-alias", "passwordalias").String,
				Optional:            true,
			},
			"external_revocation_basic_auth_header_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Basic authentication header name", "external-revocation-basic-auth-headername", "").String,
				Optional:            true,
			},
			"external_revocation_custom_header_format": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Custom header pattern", "external-revocation-custom-headername-format", "").String,
				Optional:            true,
			},
			"external_revocation_cache_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Cache type", "external-revocation-cache-type", "").AddStringEnum("Protocol", "NoCache", "TimeToLive").AddDefaultValue("NoCache").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Protocol", "NoCache", "TimeToLive"),
				},
				Default: stringdefault.StaticString("NoCache"),
			},
			"external_revocation_cache_time_to_live": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Time to live", "external-revocation-cache-ttl", "").AddIntegerRange(0, 4294967295).AddDefaultValue("900").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 4294967295),
				},
				Default: int64default.StaticInt64(900),
			},
			"external_revocation_fail_on_error": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Fail on error", "external-revocation-fail-on-error", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"enable_token_management": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable token management", "enable-token-management", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"token_manager_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Token manager type", "token-manager-type", "").AddStringEnum("native", "external").AddDefaultValue("native").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("native", "external"),
				},
				Default: stringdefault.StaticString("native"),
			},
			"api_security_token_manager": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("API security token manager", "api-security-token-manager", "apisecuritytokenmanager").AddDefaultValue("default").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default"),
			},
			"enable_application_revocation": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Application revocation", "apic-app-revoke-enable", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"application_revocation_endpoint": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Application revocation endpoint", "apic-app-revoke-endpoint", "").AddDefaultValue("/oauth2/revoke").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/oauth2/revoke"),
			},
			"enable_owner_revocation": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Resource owner revocation", "apic-owner-revoke-enable", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"owner_revocation_endpoint": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Resource owner revocation endpoint", "apic-owner-revoke-endpoint", "").AddDefaultValue("/oauth2/issued").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/oauth2/issued"),
			},
			"token_validation_req": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Token validation requirement", "token-validation-requirement", "").AddStringEnum("connected", "active", "custom").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("connected", "active", "custom"),
				},
			},
			"third_party_azurl": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Authorization endpoint", "third-party-az-url", "").String,
				Optional:            true,
			},
			"third_party_token_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Token endpoint", "third-party-token-url", "").String,
				Optional:            true,
			},
			"third_party_introspect_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Introspection endpoint", "third-party-introspect-url", "").String,
				Optional:            true,
			},
			"third_party_introspect_cache_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Cache type", "third-party-introspect-cache-type", "").AddStringEnum("Protocol", "NoCache", "TimeToLive").AddDefaultValue("NoCache").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Protocol", "NoCache", "TimeToLive"),
				},
				Default: stringdefault.StaticString("NoCache"),
			},
			"third_party_introspect_cache_time_to_live": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Time to live", "third-party-introspect-cache-ttl", "").AddIntegerRange(0, 4294967295).AddDefaultValue("900").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 4294967295),
				},
				Default: int64default.StaticInt64(900),
			},
			"third_party_authorization_header_pass_thru": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Retain Authorization header", "third-party-authorization-header-pass-thru", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"third_party_introspect_url_security": models.GetDmSecurityTypeResourceSchema("Introspection endpoint security", "third-party-introspect-url-security", "", false),
			"third_party_introspect_basic_auth_user_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Basic authentication username", "third-party-introspect-basic-auth-username", "").String,
				Optional:            true,
			},
			"third_party_introspect_basic_auth_password": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Basic authentication password", "third-party-introspect-basic-auth-password-alias", "passwordalias").String,
				Optional:            true,
			},
			"third_party_basic_auth_header_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Basic authentication header name", "third-party-introspect-basic-auth-headername", "").String,
				Optional:            true,
			},
			"third_party_custom_header_name_format": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Custom header pattern", "third-party-introspect-custom-headername-format", "").String,
				Optional:            true,
			},
			"third_party_introspect_ssl_profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "third-party-introspect-ssl-profile", "sslclientprofile").String,
				Optional:            true,
			},
		},
	}
}

func (r *OAuthProviderSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *OAuthProviderSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.OAuthProviderSettings

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `OAuthProviderSettings`)
	_, err := r.client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "POST", err))
		return
	}

	_, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s", "POST", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OAuthProviderSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.OAuthProviderSettings

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && (strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}

	if data.IsNull() {
		// Import
		data.FromBody(ctx, `OAuthProviderSettings`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `OAuthProviderSettings`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OAuthProviderSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.OAuthProviderSettings

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `OAuthProviderSettings`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}
	_, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s", "POST", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OAuthProviderSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.OAuthProviderSettings

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}

	resp.State.RemoveResource(ctx)
}
