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

type SocialLoginPolicyList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &SocialLoginPolicyDataSource{}
	_ datasource.DataSourceWithConfigure = &SocialLoginPolicyDataSource{}
)

func NewSocialLoginPolicyDataSource() datasource.DataSource {
	return &SocialLoginPolicyDataSource{}
}

type SocialLoginPolicyDataSource struct {
	pData *tfutils.ProviderData
}

func (d *SocialLoginPolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_social_login_policy"
}

func (d *SocialLoginPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DataPower can act as an OpenID Connect client. In this case, a social login policy enables DataPower to redirect the user to a social login provider like Google for user authentication and consent for authorization.",
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
							MarkdownDescription: "Enter a descriptive summary for the configuration.",
							Computed:            true,
						},
						"client_id": schema.StringAttribute{
							MarkdownDescription: "Specify the ID of DataPower that is registered with the social login provider.",
							Computed:            true,
						},
						"client_secret": schema.StringAttribute{
							MarkdownDescription: "Specify the secret of DataPower that is registered with the social login provider.",
							Computed:            true,
						},
						"client_grant": schema.StringAttribute{
							MarkdownDescription: "Controls how DataPower generates the client request to the social login provider.",
							Computed:            true,
						},
						"client_scope": schema.StringAttribute{
							MarkdownDescription: "Specifies the scope value that defines what access privileges are requested for access tokens, ID tokens, or both. Use space separated strings. For example, <tt>openid email</tt> .",
							Computed:            true,
						},
						"client_redirect_uri": schema.StringAttribute{
							MarkdownDescription: "<p>Specifies the URI that the social login policy redirects the client to after the client obtains a code or an access token. The URI must match with what is registered at the social login provider for DataPower as the OAuth/OIDC client. The URI is included in the OAuth/OIDC client request that DataPower generates.</p><p>Note that the social login provider Google mandates that the redirect URI must be a fully qualified host name instead of an IP address.</p><p>Note that the redirect URI should end with the suffix '/social-login-callback' in the pathname in order to differentiate between the callback requests and other types of requests coming into the service.</p><p>You can specify the value of this redirect URI in the following forms.</p><ul><li>Static string. Enter a static string as the redirect URI. Must end with the suffix '/social-login-callback'.</li><li>URL-in/suffix. In this case, it takes the value from the inbound URL service variable var://service/URL-in and then suffixes the value with whatever is specified after <tt>URL-in</tt> as the redirect URI. For example, the value of this property is 'URL-in/social-login-callback' and the incoming URL is 'https://datapower.ibm.com:10087/getresources', then the redirect URI is constructed as 'https://datapower.ibm.com:10087/getresources/social-login-callback'.</li><li>Context variable. You can set a context variable before you invoke this AAA action and specify the context variable name for this value. For example, var://context/AAA/social-login-redirect-uri</li></ul>",
							Computed:            true,
						},
						"client_optional_query_params": schema.StringAttribute{
							MarkdownDescription: "Specifies the optional query parameters to include in the initial OAuth/OIDC request that DataPower sends to the social login provider. Enter the optional query parameters as name=value pairs and separate each pair with an ampersand. For example, prompt=consent&amp;login_hint=jsmith@example.com&amp;openid.realm=example.comi&amp;hd=example.com.",
							Computed:            true,
						},
						"client_ssl_profile": schema.StringAttribute{
							MarkdownDescription: "Specifies the TLS client profile to secure connections when DataPower obtains an access token from the social login provider.",
							Computed:            true,
						},
						"social_provider": schema.StringAttribute{
							MarkdownDescription: "Controls which social login provider to use.",
							Computed:            true,
						},
						"provider_az_endpoint": schema.StringAttribute{
							MarkdownDescription: "Specify the provider's endpoint URL that accepts an authorization request from a client to perform social login with the provider. When the provider is Google, you can retrieve the authorization endpoint URL from the Discovery document for Google's OpenID Connect service.",
							Computed:            true,
						},
						"provider_token_endpoint": schema.StringAttribute{
							MarkdownDescription: "Specify the provider's endpoint URL that accepts an authorization grant, or code, from a client in exchange for an access token from the social login provider. When the provider is Google, you can retrieve the token endpoint URL from the Discovery document for Google's OpenID Connect service.",
							Computed:            true,
						},
						"validate_jwt_token": schema.BoolAttribute{
							MarkdownDescription: "<p>Controls whether to validate the JWT token (ID token)from the provider. If yes, it is recommended that you validate the ID token that is obtained from Google by defining the following settings in the JWT Validator configuration.</p><p><ol><li>Verify the signature by fetching the certs from https://www.googleapis.com/oauth2/v3/certs</li><li>Verify that the <tt>aud</tt> claim matches the client ID of DataPower.</li><li>Verify that the <tt>iss</tt> claim matches accounts.google.com or https://accounts.google.com</li></ol></p><p>For other recommendations on validating the ID token from Google, see https://developers.google.com/identity/protocols/OpenIDConnect.</p>",
							Computed:            true,
						},
						"jwt_validator": schema.StringAttribute{
							MarkdownDescription: "Specify the JWT Validator configuration that defines how to validate and verify the ID token.",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *SocialLoginPolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *SocialLoginPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data SocialLoginPolicyList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.SocialLoginPolicy{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.SocialLoginPolicy{}
	if value := res.Get(`SocialLoginPolicy`); value.Exists() {
		for _, v := range value.Array() {
			item := models.SocialLoginPolicy{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.SocialLoginPolicyObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.SocialLoginPolicyObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
