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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
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
	client *client.DatapowerClient
}

func (d *OAuthProviderSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_oauthprovidersettings"
}

func (d *OAuthProviderSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "OAuth provider settings",
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
							MarkdownDescription: "Enable debug headers",
							Computed:            true,
						},
						"provider_type": schema.StringAttribute{
							MarkdownDescription: "Provider type",
							Computed:            true,
						},
						"scopes_allowed": schema.StringAttribute{
							MarkdownDescription: "Allowed scopes",
							Computed:            true,
						},
						"default_scopes": schema.StringAttribute{
							MarkdownDescription: "Default scopes",
							Computed:            true,
						},
						"supported_grant_types":  models.GetDmOAuthProviderGrantTypeDataSourceSchema("Supported grant types", "supported-grant-types", ""),
						"supported_client_types": models.GetDmAllowedClientTypeDataSourceSchema("Supported client types", "supported-client-types", ""),
						"apic_provider_base_path": schema.StringAttribute{
							MarkdownDescription: "Provider base path",
							Computed:            true,
						},
						"apic_authorize_endpoint": schema.StringAttribute{
							MarkdownDescription: "Authorization endpoint",
							Computed:            true,
						},
						"apic_token_endpoint": schema.StringAttribute{
							MarkdownDescription: "Token endpoint",
							Computed:            true,
						},
						"apic_enable_introspection": schema.BoolAttribute{
							MarkdownDescription: "Enable token introspection",
							Computed:            true,
						},
						"apic_introspect_endpoint": schema.StringAttribute{
							MarkdownDescription: "Introspection endpoint",
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
							MarkdownDescription: "Access token TTL",
							Computed:            true,
						},
						"apic_auth_code_ttl": schema.Int64Attribute{
							MarkdownDescription: "Authorization code TTL",
							Computed:            true,
						},
						"apic_enable_refresh_token": schema.BoolAttribute{
							MarkdownDescription: "Enable refresh tokens",
							Computed:            true,
						},
						"apic_one_time_use_refreshtoken": schema.BoolAttribute{
							MarkdownDescription: "One-time use refresh token",
							Computed:            true,
						},
						"apic_refresh_token_limit": schema.Int64Attribute{
							MarkdownDescription: "APIC refresh token count",
							Computed:            true,
						},
						"apic_refresh_token_ttl": schema.Int64Attribute{
							MarkdownDescription: "Refresh token TTL",
							Computed:            true,
						},
						"apic_maximum_consent_ttl": schema.Int64Attribute{
							MarkdownDescription: "Consent TTL",
							Computed:            true,
						},
						"adv_scope_validation_enabled": schema.BoolAttribute{
							MarkdownDescription: "Enable advanced scope validation",
							Computed:            true,
						},
						"advanced_scope_url_override": schema.BoolAttribute{
							MarkdownDescription: "Use URL from API Security Definition",
							Computed:            true,
						},
						"advanced_scope_url": schema.StringAttribute{
							MarkdownDescription: "Advanced scope URL",
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
							MarkdownDescription: "Basic authentication header name",
							Computed:            true,
						},
						"advanced_scope_custom_headers": schema.StringAttribute{
							MarkdownDescription: "Request headers",
							Computed:            true,
						},
						"advanced_scope_custom_contexts": schema.StringAttribute{
							MarkdownDescription: "Response context variables",
							Computed:            true,
						},
						"apic_enable_oidc": schema.BoolAttribute{
							MarkdownDescription: "Enable OIDC",
							Computed:            true,
						},
						"apicoidc_hybrid_response_types": models.GetDmOIDCHybridResponseTypeDataSourceSchema("OIDC Hybrid Flow Response Types", "apic-oidc-hybrid-response-types", ""),
						"apic_support_pkce": schema.BoolAttribute{
							MarkdownDescription: "Support PKCE",
							Computed:            true,
						},
						"apic_require_pkce": schema.BoolAttribute{
							MarkdownDescription: "Require PKCE",
							Computed:            true,
						},
						"apic_support_pkce_plain": schema.BoolAttribute{
							MarkdownDescription: "Support PKCE 'plain' challenge method",
							Computed:            true,
						},
						"apic_token_type_to_generate": schema.StringAttribute{
							MarkdownDescription: "Type of token to generate",
							Computed:            true,
						},
						"metadata_from": models.GetDmMetadataFromTypeDataSourceSchema("Obtain metadata from", "metadata-from", ""),
						"metadata_url": schema.StringAttribute{
							MarkdownDescription: "External metadata URL",
							Computed:            true,
						},
						"metadata_ssl_profile": schema.StringAttribute{
							MarkdownDescription: "TLS profile to access metadata URL",
							Computed:            true,
						},
						"metadata_header_for_access_token": schema.StringAttribute{
							MarkdownDescription: "Response header for access token",
							Computed:            true,
						},
						"metadata_header_for_payload": schema.StringAttribute{
							MarkdownDescription: "Response header for payload",
							Computed:            true,
						},
						"external_revocation_url": schema.StringAttribute{
							MarkdownDescription: "External management URL",
							Computed:            true,
						},
						"external_revocation_ssl_profile": schema.StringAttribute{
							MarkdownDescription: "External management TLS client profile",
							Computed:            true,
						},
						"external_revocation_url_security": models.GetDmSecurityTypeDataSourceSchema("External management security", "external-revocation-url-security", ""),
						"external_revocation_basic_auth_user_name": schema.StringAttribute{
							MarkdownDescription: "Basic authentication username",
							Computed:            true,
						},
						"external_revocation_basic_auth_password": schema.StringAttribute{
							MarkdownDescription: "Basic authentication password",
							Computed:            true,
						},
						"external_revocation_basic_auth_header_name": schema.StringAttribute{
							MarkdownDescription: "Basic authentication header name",
							Computed:            true,
						},
						"external_revocation_custom_header_format": schema.StringAttribute{
							MarkdownDescription: "Custom header pattern",
							Computed:            true,
						},
						"external_revocation_cache_type": schema.StringAttribute{
							MarkdownDescription: "Cache type",
							Computed:            true,
						},
						"external_revocation_cache_time_to_live": schema.Int64Attribute{
							MarkdownDescription: "Time to live",
							Computed:            true,
						},
						"external_revocation_fail_on_error": schema.BoolAttribute{
							MarkdownDescription: "Fail on error",
							Computed:            true,
						},
						"enable_token_management": schema.BoolAttribute{
							MarkdownDescription: "Enable token management",
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
							MarkdownDescription: "Application revocation",
							Computed:            true,
						},
						"application_revocation_endpoint": schema.StringAttribute{
							MarkdownDescription: "Application revocation endpoint",
							Computed:            true,
						},
						"enable_owner_revocation": schema.BoolAttribute{
							MarkdownDescription: "Resource owner revocation",
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
							MarkdownDescription: "Introspection endpoint",
							Computed:            true,
						},
						"third_party_introspect_cache_type": schema.StringAttribute{
							MarkdownDescription: "Cache type",
							Computed:            true,
						},
						"third_party_introspect_cache_time_to_live": schema.Int64Attribute{
							MarkdownDescription: "Time to live",
							Computed:            true,
						},
						"third_party_authorization_header_pass_thru": schema.BoolAttribute{
							MarkdownDescription: "Retain Authorization header",
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
							MarkdownDescription: "Basic authentication header name",
							Computed:            true,
						},
						"third_party_custom_header_name_format": schema.StringAttribute{
							MarkdownDescription: "Custom header pattern",
							Computed:            true,
						},
						"third_party_introspect_ssl_profile": schema.StringAttribute{
							MarkdownDescription: "TLS client profile",
							Computed:            true,
						},
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

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *OAuthProviderSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OAuthProviderSettingsList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.OAuthProviderSettings{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
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
