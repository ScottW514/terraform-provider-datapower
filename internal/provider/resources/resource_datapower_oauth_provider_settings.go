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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &OAuthProviderSettingsResource{}

func NewOAuthProviderSettingsResource() resource.Resource {
	return &OAuthProviderSettingsResource{}
}

type OAuthProviderSettingsResource struct {
	pData *tfutils.ProviderData
}

func (r *OAuthProviderSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_oauth_provider_settings"
}

func (r *OAuthProviderSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("An OAuth provider settings configuration defines how a client application is authorized to access resources on behalf of the resource owner.", "oauth-provider-settings", "").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable debug mode to add security error details in response headers. In debug mode when you use a validation endpoint, security error details are sent in the <tt>x-apic-debug-oauth-error</tt> and <tt>x-apic-debug-oauth-error-desc</tt> response headers.", "enable-debug-mode", "").AddDefaultValue("false").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the scopes that the access token is valid to access. To specify multiple scopes, use a space between each scope. The order of scopes does not matter. Scopes ensure that the granted access token is valid to access only specific protected resources.", "scopes-allowed", "").String,
				Optional:            true,
			},
			"default_scopes": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the default scopes to apply when the request does not contain a scope. To specify multiple scopes, use a space between each scope. The order of scopes does not matter. <p>The default scopes must be a subset of the allowed scopes in the API security OAuth requirement. Without defined scopes and the request does not contain a scope, an invalid scope error is returned.</p><p>Scopes ensure that the granted access token is valid to access only specific protected resources.</p>", "default-scopes", "").String,
				Optional:            true,
			},
			"supported_grant_types":  models.GetDmOAuthProviderGrantTypeResourceSchema("Specify the supported grant types. Each grant type defines a method to grant authorization to client applications.", "supported-grant-types", "", false),
			"supported_client_types": models.GetDmAllowedClientTypeResourceSchema("Supported client types", "supported-client-types", "", false),
			"apic_provider_base_path": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the base path on which the OAuth provider API is served. The default value is <tt>/</tt> .", "apic-provider-base-path", "").AddDefaultValue("/").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/"),
			},
			"apic_authorize_endpoint": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the endpoint where the client application obtains authorization grant. The default value is <tt>/oauth2/authorize</tt> .", "apic-authorize-endpoint", "").AddDefaultValue("/oauth2/authorize").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/oauth2/authorize"),
			},
			"apic_token_endpoint": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the endpoint where the client application exchanges an authorization grant for an access token. The default value is <tt>/oauth2/token</tt> .", "apic-token-endpoint", "").AddDefaultValue("/oauth2/token").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/oauth2/token"),
			},
			"apic_enable_introspection": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable the introspection of access tokens. When enabled, authorized protected resources can introspect the access token to determine the metadata for making appropriate authorization decisions. By default, token introspection is disabled.", "apic-enable-introspection", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"apic_introspect_endpoint": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the endpoint for token introspection. The default value is <tt>/oauth2/introspect</tt> .", "apic-introspect-endpoint", "").AddDefaultValue("/oauth2/introspect").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/oauth2/introspect"),
			},
			"apic_token_secret": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Token secret", "apic-token-secret", "crypto_sskey_").String,
				Optional:            true,
			},
			"apic_one_time_use_accesstoken": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("One-time use access token", "apic-enable-one-time-use-access-token", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"apic_access_token_ttl": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the time in seconds that an access token remains valid. The default value is 3600.", "apic-access-token-ttl", "").AddIntegerRange(1, 63244800).AddDefaultValue("3600").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 63244800),
				},
				Default: int64default.StaticInt64(3600),
			},
			"apic_auth_code_ttl": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the time in seconds that an authorization code remains valid. The default value is 300.", "apic-auth-code-ttl", "").AddIntegerRange(1, 600).AddDefaultValue("300").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 600),
				},
				Default: int64default.StaticInt64(300),
			},
			"apic_enable_refresh_token": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable issuing refresh tokens. Refresh tokens are issued to the client. Refresh tokens are used to obtain a new access token when the current access token becomes invalid, expires, or are used to obtain additional access tokens with identical or narrower scope. By default, this setting is disabled.", "apic-enable-refresh-token", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"apic_one_time_use_refreshtoken": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether a refresh tokens is one-time use. <ul><li>When enabled, the refresh token is one-time use. This setting is the default value.</li><li>When disabled, the refresh token can be reused until it expires or is revoked.</li></ul>", "apic-enable-one-time-use-refresh-token", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"apic_refresh_token_limit": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the number of refresh tokens to allow to be generated. The default value is 10.", "apic-refresh-token-limit", "").AddIntegerRange(1, 4096).AddDefaultValue("10").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 4096),
				},
				Default: int64default.StaticInt64(10),
			},
			"apic_refresh_token_ttl": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the time in seconds that a refresh token remains valid. The default value is 5400.", "apic-refresh-token-ttl", "").AddIntegerRange(2, 252979200).AddDefaultValue("5400").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(2, 252979200),
				},
				Default: int64default.StaticInt64(5400),
			},
			"apic_maximum_consent_ttl": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the time in seconds that a consent remains valid. The default value is 0, which disables maximum consent.", "apic-maximum-consent-ttl", "").AddIntegerRange(0, 2529792000).AddDefaultValue("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 2529792000),
				},
				Default: int64default.StaticInt64(0),
			},
			"adv_scope_validation_enabled": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable advanced scope validation that you can use to provide additional scope checking.", "advanced-scope-validation-enabled", "").AddDefaultValue("false").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL to send scope information for validation. This external endpoint is where the specified scope is verified. You must define this property for advanced scope validation.", "advanced-scope-url", "").String,
				Optional:            true,
			},
			"adv_scope_tls_profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS profile to access advanced scope URL", "advanced-scope-tls-profile", "ssl_client_profile").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Basic authentication password", "advanced-scope-basic-auth-password-alias", "password_alias").String,
				Optional:            true,
			},
			"adv_scope_basic_auth_header_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the header name to use for sending encoded or non-encoded authentication string in the header. For example, <tt>x-basic-authorization-header</tt> .", "advanced-scope-basic-auth-headername", "").String,
				Optional:            true,
			},
			"advanced_scope_custom_headers": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the custom headers to send with the advanced scope validation request. Use a regular expression match to include headers from the initial request.", "advanced-scope-request-headers", "").String,
				Optional:            true,
			},
			"advanced_scope_custom_contexts": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify custom context variables to save headers from the advanced scope validation request. Use a regular expression to include headers from the advanced scope endpoint response.", "advanced-scope-response-contexts", "").String,
				Optional:            true,
			},
			"apic_enable_oidc": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable OIDC to verify the identity of the user. When enabled, the client application verifies the identity of the user based on the requirement of an OIDC provider before requesting access to client resources. By default, OIDC token generation is enabled. OIDC is only available for implicit and AZ code grant types.", "apic-enable-oidc", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"apicoidc_hybrid_response_types": models.GetDmOIDCHybridResponseTypeResourceSchema("OIDC Hybrid Flow Response Types", "apic-oidc-hybrid-response-types", "", false),
			"apic_support_pkce": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the application should enforce PKCE if provided by the client. For more information, see RFC 7636.", "apic-support-pkce", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"apic_require_pkce": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the application must enforce PKCE. For more information, see RFC 7636.", "apic-require-pkce", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"apic_support_pkce_plain": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to support the PKCE <tt>plain</tt> code challenge transform method. For more information, see RFC 7636.", "apic-support-pkce-plain", "").AddDefaultValue("false").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL to a remote server where the custom metadata is stored. The value can include one or more runtime context variables in the <tt>$(variable)</tt> format.", "metadata-url", "").String,
				Optional:            true,
			},
			"metadata_ssl_profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS profile to access metadata URL", "metadata-ssl-profile", "ssl_client_profile").String,
				Optional:            true,
			},
			"metadata_header_for_access_token": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the response header to place in the access token. These headers are in the response from the authentication or metadata endpoint.", "mdheader-for-accesstoken", "").AddDefaultValue("X-API-OAuth-Metadata-For-AccessToken").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("X-API-OAuth-Metadata-For-AccessToken"),
			},
			"metadata_header_for_payload": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the response header to place in the response payload. These headers are in the response from the authentication or metadata endpoint.", "mdheader-for-payload", "").AddDefaultValue("X-API-OAuth-Metadata-For-Payload").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("X-API-OAuth-Metadata-For-Payload"),
			},
			"external_revocation_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify an external endpoint through which the token management is accomplished. The value can include one or more runtime context variables in the <tt>$(variable)</tt> format.", "external-revocation-url", "").String,
				Optional:            true,
			},
			"external_revocation_ssl_profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("External management TLS client profile", "external-revocation-ssl-profile", "ssl_client_profile").String,
				Optional:            true,
			},
			"external_revocation_url_security": models.GetDmSecurityTypeResourceSchema("External management security", "external-revocation-url-security", "", false),
			"external_revocation_basic_auth_user_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "external-revocation-basic-auth-username", "").String,
				Optional:            true,
			},
			"external_revocation_basic_auth_password": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Basic authentication password", "external-revocation-basic-auth-password-alias", "password_alias").String,
				Optional:            true,
			},
			"external_revocation_basic_auth_header_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the header name to use for sending encoded or non-encoded authentication string in the header. For example, <tt>x-external-basic-authorization-header</tt> .", "external-revocation-basic-auth-headername", "").String,
				Optional:            true,
			},
			"external_revocation_custom_header_format": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the pattern of header names to include from the original message. For example, <tt>x-external-management-*</tt> .", "external-revocation-custom-headername-format", "").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the validity period in seconds for external management service responses in the cache. The default value is 900.", "external-revocation-cache-ttl", "").AddIntegerRange(0, 4294967295).AddDefaultValue("900").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 4294967295),
				},
				Default: int64default.StaticInt64(900),
			},
			"external_revocation_fail_on_error": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to stop processing if connection to external management service fails. If failed, stops token-generation or use, and returns an error.", "external-revocation-fail-on-error", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"enable_token_management": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify if security token details should be managed and stored. Enabling token management for security token details provides the ability to create one-time use tokens, prevent AZ code reuse, and support allow-listing through the use of the token manager.", "enable-token-management", "").AddDefaultValue("true").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("API security token manager", "api-security-token-manager", "api_security_token_manager").AddDefaultValue("default").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default"),
			},
			"enable_application_revocation": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable revocation by application. Enabling application revocation allows the application to revoke consent before the token expires.", "apic-app-revoke-enable", "").AddDefaultValue("false").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable revocation by resource owner. Enabling resource owner revocation allows the resource owner to revoke consent before the token expires.", "apic-owner-revoke-enable", "").AddDefaultValue("false").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the endpoint for token-introspection operation. The value can include one or more runtime context variables in the <tt>$(variable)</tt> format.", "third-party-introspect-url", "").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the validity period in seconds for third-party provider responses in the cache. The default value is 900.", "third-party-introspect-cache-ttl", "").AddIntegerRange(0, 4294967295).AddDefaultValue("900").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 4294967295),
				},
				Default: int64default.StaticInt64(900),
			},
			"third_party_authorization_header_pass_thru": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to retain or remove the <tt>Authorization</tt> header for a bearer token. The default behavior is to remove this header.", "third-party-authorization-header-pass-thru", "").AddDefaultValue("false").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Basic authentication password", "third-party-introspect-basic-auth-password-alias", "password_alias").String,
				Optional:            true,
			},
			"third_party_basic_auth_header_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the header name to send the encoded or non-encoded authentication string. For example, <tt>x-introspect-basic-authorization-header</tt> .", "third-party-introspect-basic-auth-headername", "").String,
				Optional:            true,
			},
			"third_party_custom_header_name_format": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the pattern of header name to send additional information. For example, <tt>x-introspect-*</tt> .", "third-party-introspect-custom-headername-format", "").String,
				Optional:            true,
			},
			"third_party_introspect_ssl_profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "third-party-introspect-ssl-profile", "ssl_client_profile").String,
				Optional:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *OAuthProviderSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *OAuthProviderSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.OAuthProviderSettings
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `OAuthProviderSettings`)
	_, err := r.pData.Client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "POST", err))
		return
	}
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OAuthProviderSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.OAuthProviderSettings
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
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
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.OAuthProviderSettings
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `OAuthProviderSettings`))
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

func (r *OAuthProviderSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.OAuthProviderSettings
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *OAuthProviderSettingsResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.OAuthProviderSettings

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
