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
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type OAuthProviderSettingsList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &OAuthProviderSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &OAuthProviderSettingsDataSource{}
)

func NewOAuthProviderSettingsDataSource() datasource.DataSource {
	return &OAuthProviderSettingsDataSource{}
}

type OAuthProviderSettingsDataSource struct {
	pData *tfutils.ProviderData
}

func (d *OAuthProviderSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_oauthprovidersettings"
}

func (d *OAuthProviderSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "An OAuth provider settings configuration defines how a client application is authorized to access resources on behalf of the resource owner.",
		Attributes: map[string]schema.Attribute{
			"app_domain": schema.StringAttribute{
				MarkdownDescription: "The name of the application domain the object belongs to",
				Required:            true,
			},
			"result": schema.ListNestedAttribute{
				MarkdownDescription: "List of objects",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Name of the object. Must be unique among object types in application domain.",
							Computed:            true,
						},
						"app_domain": schema.StringAttribute{
							MarkdownDescription: "The name of the application domain the object belongs to",
							Computed:            true,
						},
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"enable_debug_mode": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enable debug mode to add security error details in response headers. In debug mode when you use a validation endpoint, security error details are sent in the <tt>x-apic-debug-oauth-error</tt> and <tt>x-apic-debug-oauth-error-desc</tt> response headers.",
							Computed:            true,
						},
						"provider_type": schema.StringAttribute{
							MarkdownDescription: "Provider type",
							Computed:            true,
						},
						"scopes_allowed": schema.StringAttribute{
							MarkdownDescription: "Specify the scopes that the access token is valid to access. To specify multiple scopes, use a space between each scope. The order of scopes does not matter. Scopes ensure that the granted access token is valid to access only specific protected resources.",
							Computed:            true,
						},
						"default_scopes": schema.StringAttribute{
							MarkdownDescription: "Specify the default scopes to apply when the request does not contain a scope. To specify multiple scopes, use a space between each scope. The order of scopes does not matter. <p>The default scopes must be a subset of the allowed scopes in the API security OAuth requirement. Without defined scopes and the request does not contain a scope, an invalid scope error is returned.</p><p>Scopes ensure that the granted access token is valid to access only specific protected resources.</p>",
							Computed:            true,
						},
						"supported_grant_types":  models.GetDmOAuthProviderGrantTypeDataSourceSchema("Specify the supported grant types. Each grant type defines a method to grant authorization to client applications.", "supported-grant-types", ""),
						"supported_client_types": models.GetDmAllowedClientTypeDataSourceSchema("Supported client types", "supported-client-types", ""),
						"apic_provider_base_path": schema.StringAttribute{
							MarkdownDescription: "Specify the base path on which the OAuth provider API is served. The default value is <tt>/</tt> .",
							Computed:            true,
						},
						"apic_authorize_endpoint": schema.StringAttribute{
							MarkdownDescription: "Specify the endpoint where the client application obtains authorization grant. The default value is <tt>/oauth2/authorize</tt> .",
							Computed:            true,
						},
						"apic_token_endpoint": schema.StringAttribute{
							MarkdownDescription: "Specify the endpoint where the client application exchanges an authorization grant for an access token. The default value is <tt>/oauth2/token</tt> .",
							Computed:            true,
						},
						"apic_enable_introspection": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enable the introspection of access tokens. When enabled, authorized protected resources can introspect the access token to determine the metadata for making appropriate authorization decisions. By default, token introspection is disabled.",
							Computed:            true,
						},
						"apic_introspect_endpoint": schema.StringAttribute{
							MarkdownDescription: "Specify the endpoint for token introspection. The default value is <tt>/oauth2/introspect</tt> .",
							Computed:            true,
						},
						"apic_token_secret": schema.StringAttribute{
							MarkdownDescription: "Token secret",
							Computed:            true,
						},
						"apic_one_time_use_accesstoken": schema.BoolAttribute{
							MarkdownDescription: "One-time use access token",
							Computed:            true,
						},
						"apic_access_token_ttl": schema.Int64Attribute{
							MarkdownDescription: "Specify the time in seconds that an access token remains valid. The default value is 3600.",
							Computed:            true,
						},
						"apic_auth_code_ttl": schema.Int64Attribute{
							MarkdownDescription: "Specify the time in seconds that an authorization code remains valid. The default value is 300.",
							Computed:            true,
						},
						"apic_enable_refresh_token": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enable issuing refresh tokens. Refresh tokens are issued to the client. Refresh tokens are used to obtain a new access token when the current access token becomes invalid, expires, or are used to obtain additional access tokens with identical or narrower scope. By default, this setting is disabled.",
							Computed:            true,
						},
						"apic_one_time_use_refreshtoken": schema.BoolAttribute{
							MarkdownDescription: "Specify whether a refresh tokens is one-time use. <ul><li>When enabled, the refresh token is one-time use. This setting is the default value.</li><li>When disabled, the refresh token can be reused until it expires or is revoked.</li></ul>",
							Computed:            true,
						},
						"apic_refresh_token_limit": schema.Int64Attribute{
							MarkdownDescription: "Specify the number of refresh tokens to allow to be generated. The default value is 10.",
							Computed:            true,
						},
						"apic_refresh_token_ttl": schema.Int64Attribute{
							MarkdownDescription: "Specify the time in seconds that a refresh token remains valid. The default value is 5400.",
							Computed:            true,
						},
						"apic_maximum_consent_ttl": schema.Int64Attribute{
							MarkdownDescription: "Specify the time in seconds that a consent remains valid. The default value is 0, which disables maximum consent.",
							Computed:            true,
						},
						"adv_scope_validation_enabled": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enable advanced scope validation that you can use to provide additional scope checking.",
							Computed:            true,
						},
						"advanced_scope_url_override": schema.BoolAttribute{
							MarkdownDescription: "Use URL from API Security Definition",
							Computed:            true,
						},
						"advanced_scope_url": schema.StringAttribute{
							MarkdownDescription: "Specify the URL to send scope information for validation. This external endpoint is where the specified scope is verified. You must define this property for advanced scope validation.",
							Computed:            true,
						},
						"adv_scope_tls_profile": schema.StringAttribute{
							MarkdownDescription: "TLS profile to access advanced scope URL",
							Computed:            true,
						},
						"adv_scope_url_security_enabled": schema.BoolAttribute{
							MarkdownDescription: "Enable advanced scope endpoint security",
							Computed:            true,
						},
						"adv_scope_url_security": models.GetDmSecurityTypeDataSourceSchema("Advanced scope endpoint security", "advanced-scope-url-security", ""),
						"adv_scope_basic_auth_user_name": schema.StringAttribute{
							MarkdownDescription: "Basic authentication username",
							Computed:            true,
						},
						"adv_scope_basic_auth_password": schema.StringAttribute{
							MarkdownDescription: "Basic authentication password",
							Computed:            true,
						},
						"adv_scope_basic_auth_header_name": schema.StringAttribute{
							MarkdownDescription: "Specify the header name to use for sending encoded or non-encoded authentication string in the header. For example, <tt>x-basic-authorization-header</tt> .",
							Computed:            true,
						},
						"advanced_scope_custom_headers": schema.StringAttribute{
							MarkdownDescription: "Specify the custom headers to send with the advanced scope validation request. Use a regular expression match to include headers from the initial request.",
							Computed:            true,
						},
						"advanced_scope_custom_contexts": schema.StringAttribute{
							MarkdownDescription: "Specify custom context variables to save headers from the advanced scope validation request. Use a regular expression to include headers from the advanced scope endpoint response.",
							Computed:            true,
						},
						"apic_enable_oidc": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enable OIDC to verify the identity of the user. When enabled, the client application verifies the identity of the user based on the requirement of an OIDC provider before requesting access to client resources. By default, OIDC token generation is enabled. OIDC is only available for implicit and AZ code grant types.",
							Computed:            true,
						},
						"apicoidc_hybrid_response_types": models.GetDmOIDCHybridResponseTypeDataSourceSchema("OIDC Hybrid Flow Response Types", "apic-oidc-hybrid-response-types", ""),
						"apic_support_pkce": schema.BoolAttribute{
							MarkdownDescription: "Specify whether the application should enforce PKCE if provided by the client. For more information, see RFC 7636.",
							Computed:            true,
						},
						"apic_require_pkce": schema.BoolAttribute{
							MarkdownDescription: "Specify whether the application must enforce PKCE. For more information, see RFC 7636.",
							Computed:            true,
						},
						"apic_support_pkce_plain": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to support the PKCE <tt>plain</tt> code challenge transform method. For more information, see RFC 7636.",
							Computed:            true,
						},
						"apic_token_type_to_generate": schema.StringAttribute{
							MarkdownDescription: "Type of token to generate",
							Computed:            true,
						},
						"metadata_from": models.GetDmMetadataFromTypeDataSourceSchema("Obtain metadata from", "metadata-from", ""),
						"metadata_url": schema.StringAttribute{
							MarkdownDescription: "Specify the URL to a remote server where the custom metadata is stored. The value can include one or more runtime context variables in the <tt>$(variable)</tt> format.",
							Computed:            true,
						},
						"metadata_ssl_profile": schema.StringAttribute{
							MarkdownDescription: "TLS profile to access metadata URL",
							Computed:            true,
						},
						"metadata_header_for_access_token": schema.StringAttribute{
							MarkdownDescription: "Specify the response header to place in the access token. These headers are in the response from the authentication or metadata endpoint.",
							Computed:            true,
						},
						"metadata_header_for_payload": schema.StringAttribute{
							MarkdownDescription: "Specify the response header to place in the response payload. These headers are in the response from the authentication or metadata endpoint.",
							Computed:            true,
						},
						"external_revocation_url": schema.StringAttribute{
							MarkdownDescription: "Specify an external endpoint through which the token management is accomplished. The value can include one or more runtime context variables in the <tt>$(variable)</tt> format.",
							Computed:            true,
						},
						"external_revocation_ssl_profile": schema.StringAttribute{
							MarkdownDescription: "External management TLS client profile",
							Computed:            true,
						},
						"external_revocation_url_security": models.GetDmSecurityTypeDataSourceSchema("External management security", "external-revocation-url-security", ""),
						"external_revocation_basic_auth_user_name": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"external_revocation_basic_auth_password": schema.StringAttribute{
							MarkdownDescription: "Basic authentication password",
							Computed:            true,
						},
						"external_revocation_basic_auth_header_name": schema.StringAttribute{
							MarkdownDescription: "Specify the header name to use for sending encoded or non-encoded authentication string in the header. For example, <tt>x-external-basic-authorization-header</tt> .",
							Computed:            true,
						},
						"external_revocation_custom_header_format": schema.StringAttribute{
							MarkdownDescription: "Specify the pattern of header names to include from the original message. For example, <tt>x-external-management-*</tt> .",
							Computed:            true,
						},
						"external_revocation_cache_type": schema.StringAttribute{
							MarkdownDescription: "Cache type",
							Computed:            true,
						},
						"external_revocation_cache_time_to_live": schema.Int64Attribute{
							MarkdownDescription: "Specify the validity period in seconds for external management service responses in the cache. The default value is 900.",
							Computed:            true,
						},
						"external_revocation_fail_on_error": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to stop processing if connection to external management service fails. If failed, stops token-generation or use, and returns an error.",
							Computed:            true,
						},
						"enable_token_management": schema.BoolAttribute{
							MarkdownDescription: "Specify if security token details should be managed and stored. Enabling token management for security token details provides the ability to create one-time use tokens, prevent AZ code reuse, and support allow-listing through the use of the token manager.",
							Computed:            true,
						},
						"token_manager_type": schema.StringAttribute{
							MarkdownDescription: "Token manager type",
							Computed:            true,
						},
						"api_security_token_manager": schema.StringAttribute{
							MarkdownDescription: "API security token manager",
							Computed:            true,
						},
						"enable_application_revocation": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enable revocation by application. Enabling application revocation allows the application to revoke consent before the token expires.",
							Computed:            true,
						},
						"application_revocation_endpoint": schema.StringAttribute{
							MarkdownDescription: "Application revocation endpoint",
							Computed:            true,
						},
						"enable_owner_revocation": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enable revocation by resource owner. Enabling resource owner revocation allows the resource owner to revoke consent before the token expires.",
							Computed:            true,
						},
						"owner_revocation_endpoint": schema.StringAttribute{
							MarkdownDescription: "Resource owner revocation endpoint",
							Computed:            true,
						},
						"token_validation_req": schema.StringAttribute{
							MarkdownDescription: "Token validation requirement",
							Computed:            true,
						},
						"third_party_azurl": schema.StringAttribute{
							MarkdownDescription: "Authorization endpoint",
							Computed:            true,
						},
						"third_party_token_url": schema.StringAttribute{
							MarkdownDescription: "Token endpoint",
							Computed:            true,
						},
						"third_party_introspect_url": schema.StringAttribute{
							MarkdownDescription: "Specify the endpoint for token-introspection operation. The value can include one or more runtime context variables in the <tt>$(variable)</tt> format.",
							Computed:            true,
						},
						"third_party_introspect_cache_type": schema.StringAttribute{
							MarkdownDescription: "Cache type",
							Computed:            true,
						},
						"third_party_introspect_cache_time_to_live": schema.Int64Attribute{
							MarkdownDescription: "Specify the validity period in seconds for third-party provider responses in the cache. The default value is 900.",
							Computed:            true,
						},
						"third_party_authorization_header_pass_thru": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to retain or remove the <tt>Authorization</tt> header for a bearer token. The default behavior is to remove this header.",
							Computed:            true,
						},
						"third_party_introspect_url_security": models.GetDmSecurityTypeDataSourceSchema("Introspection endpoint security", "third-party-introspect-url-security", ""),
						"third_party_introspect_basic_auth_user_name": schema.StringAttribute{
							MarkdownDescription: "Basic authentication username",
							Computed:            true,
						},
						"third_party_introspect_basic_auth_password": schema.StringAttribute{
							MarkdownDescription: "Basic authentication password",
							Computed:            true,
						},
						"third_party_basic_auth_header_name": schema.StringAttribute{
							MarkdownDescription: "Specify the header name to send the encoded or non-encoded authentication string. For example, <tt>x-introspect-basic-authorization-header</tt> .",
							Computed:            true,
						},
						"third_party_custom_header_name_format": schema.StringAttribute{
							MarkdownDescription: "Specify the pattern of header name to send additional information. For example, <tt>x-introspect-*</tt> .",
							Computed:            true,
						},
						"third_party_introspect_ssl_profile": schema.StringAttribute{
							MarkdownDescription: "TLS client profile",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *OAuthProviderSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *OAuthProviderSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data OAuthProviderSettingsList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.OAuthProviderSettings{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.OAuthProviderSettings{}
	if value := res.Get(`OAuthProviderSettings`); value.Exists() {
		for _, v := range value.Array() {
			item := models.OAuthProviderSettings{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.OAuthProviderSettingsObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.OAuthProviderSettingsObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
