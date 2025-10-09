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
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type OAuthSupportedClientWOList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &OAuthSupportedClientDataSource{}
	_ datasource.DataSourceWithConfigure = &OAuthSupportedClientDataSource{}
)

func NewOAuthSupportedClientDataSource() datasource.DataSource {
	return &OAuthSupportedClientDataSource{}
}

type OAuthSupportedClientDataSource struct {
	pData *tfutils.ProviderData
}

func (d *OAuthSupportedClientDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_oauth_supported_client"
}

func (d *OAuthSupportedClientDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "<p>An OAuth client profile is the basic building block for an OAuth client group. When you create an OAuth client profile, you define its role. As you select the role, the WebGUI refreshes to display the appropriate properties.</p><p>You can create the following types of OAuth client profiles. <ul><li>A client profile for authorization server endpoints: authorization endpoint and token endpoint.</li><li>A client profile for the enforcement point for the resource server.</li><li>A client profile for both authorization server endpoints and the enforcement point.</li></ul></p><p>When creating an OAuth client profile, you can use stylesheets or GatewayScript files for customization.</p><p>You can create a customized OAuth client profile that defines any combination of roles. Customization uses stylesheets or GatewayScript files that must be in the local: or store: directory. For information about the operations that these stylesheets or GatewayScript files must define, see the topic in IBM Knowledge Center.</p>",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The name of the object to retrieve.",
				Optional:            true,
			},
			"app_domain": schema.StringAttribute{
				MarkdownDescription: "The name of the application domain the object belongs to.",
				Required:            true,
			},
			"result": schema.ListNestedAttribute{
				MarkdownDescription: "List of objects. If `id` was provided and it exists, it will be the only item in the list.",
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
							MarkdownDescription: "Specifies a brief comment that describes the configuration.",
							Computed:            true,
						},
						"customized": schema.BoolAttribute{
							MarkdownDescription: "Indicates whether the configuration is for a customized OAuth client. The configuration of the customized OAuth client is defined in a stylesheet or GatewayScript file in the local: or store: directory.",
							Computed:            true,
						},
						"customized_process_url": schema.StringAttribute{
							MarkdownDescription: "<p>Specifies the location of the stylesheet or GatewayScript file that defines the customized OAuth client. The stylesheet or GatewayScript file must be in the local: or store: directory</p><p>When creating a customized OAuth client, the stylesheets or GatewayScript files must define all implementation details based on the role of the client. For information about these stylesheets or GatewayScript files, see the topic in IBM Knowledge Center.</p><p>You should provide support for the following operations: <ul><li>verify-az-request: determine whether the initial OAuth request is supported or not.</li><li>issue-az-code: issue a temporary authorization code (for the authorization code grant).</li><li>verify-az-code: verify a temporary authorization.</li><li>issue-access-token: issue an access_token.</li><li>verify-access-token: verify an access_token.</li><li>verify-refresh-token: verify a refresh_token.</li><li>client-revoke-request: handle client revocation request.</li><li>owner-revoke-request: handle owner revocation request.</li></ul></p>",
							Computed:            true,
						},
						"oauth_role": models.GetDmOAuthRoleDataSourceSchema("Identifies the role of the client when interacting with a request to access a protected resource.", "oauth-role", ""),
						"az_grant":   models.GetDmOAuthAZGrantTypeDataSourceSchema("Identifies the method to obtain the access token for authorization based on the grant type.", "az-grant", ""),
						"client_type": schema.StringAttribute{
							MarkdownDescription: "Sets the type of client based on its ability to authenticate securely with authorization server endpoints. The client type is based on the definitions that the authorization server endpoints use for secure authentication and acceptable exposure of client credentials. If the client can securely authenticate, its classification is <tt>confidential</tt> .",
							Computed:            true,
						},
						"check_client_credential": schema.BoolAttribute{
							MarkdownDescription: "Identifies whether to verify the client credentials when the DataPower Gateway protects the resource server by using access tokens.",
							Computed:            true,
						},
						"use_validation_url": schema.BoolAttribute{
							MarkdownDescription: "Uses a remote URL to validate the access token.",
							Computed:            true,
						},
						"client_authen_method": schema.StringAttribute{
							MarkdownDescription: "Identifies the method to authenticate this client.",
							Computed:            true,
						},
						"client_val_cred": schema.StringAttribute{
							MarkdownDescription: "An TLS credential used to authenticate the OAuth client sent by remote TLS peer during the TLS handshake.",
							Computed:            true,
						},
						"generate_client_secret": schema.BoolAttribute{
							MarkdownDescription: "<p>Indicates whether to generate the client secret for the OAuth client. The specification refers to the client secret as <tt>client_secret</tt> .</p><ul><li>If you generate the passphrase, the passphrase becomes the client secret.</li><li>If you do not generate the passphrase, you must explicitly define the client secret.</li></ul>",
							Computed:            true,
						},
						"caching": schema.StringAttribute{
							MarkdownDescription: "Specifies the caching mechanism to be used.",
							Computed:            true,
						},
						"validation_url": schema.StringAttribute{
							MarkdownDescription: "Specifies the validation url.",
							Computed:            true,
						},
						"validation_features": models.GetDmValidationFeaturesDataSourceSchema("Customize how to handle the validation grant type.", "validation-features", ""),
						"redirect_uri": schema.ListAttribute{
							MarkdownDescription: "<p>Defines redirection URIs that the OAuth client supports to exchange tokens. Specify each redirection URI as a PCRE.</p><p>Redirection URIs help to detect malicious clients and prevent phishing attacks. The authorization endpoint must have the registered redirection URIs before the authorization endpoint can validate the authorization request from the client. For mobile applications, the redirection URI can be an application URL; for example, <tt>mobiletrafficapp://</tt> that is defined with the <tt>^mobiletrafficapp:\\/\\/?</tt> PCRE.</p>",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"custom_scope_check": schema.BoolAttribute{
							MarkdownDescription: "<p>Indicates how to check the scope for authorization grants and access tokens.</p><ul><li>When checking the scope with custom processing, specify the location of the stylesheet or GatewayScript file with the <b>Scope Customized Process</b> property. The stylesheet or GatewayScript file must be in the local: or store: directory.</li><li>When checking the scope with a PCRE, specify the expression with the <b>Scope</b> property.</li></ul><p>You should use a custom scope check in the following situations. <ul><li>An authorization request where the OAuth client requests an authorization code.</li><li>An access request where the OAuth client requests an access token.</li><li>A resource request where the OAuth client requests a resource.</li></ul></p>",
							Computed:            true,
						},
						"scope": schema.StringAttribute{
							MarkdownDescription: "Specifies the PCRE to check the scope. The minimum length of the expression is 1 character. The maximum length of the expression is 1023 characters.",
							Computed:            true,
						},
						"scope_url": schema.StringAttribute{
							MarkdownDescription: "Specifies the location of the stylesheet or GatewayScript file for a custom scope check. The stylesheet or GatewayScript file must be in the local: or store: directory. The stylesheet or GatewayScript file validates and sets the scope to check.",
							Computed:            true,
						},
						"default_scope": schema.StringAttribute{
							MarkdownDescription: "Specifies the default value of the scope if the client does not define any scope value in the request.",
							Computed:            true,
						},
						"token_secret": schema.StringAttribute{
							MarkdownDescription: "Assigns the shared secret key to protect tokens that use the OAuth protocol. The shared secret must be at least 32 bytes in length.",
							Computed:            true,
						},
						"local_az_page_url": schema.StringAttribute{
							MarkdownDescription: "<p>Specifies the location of the stylesheet or GatewayScript file that generates the authorization form for the resource owner and handles errors. The file must be in the local: or store: directory. You can use the <tt>OAuth-Generate-HTML.xsl</tt> stylesheet in the store: directory or copy this file to the local: directory and modify as needed.</p><p>The stylesheet or GatewayScript file must be on the DataPower Gateway in the local: or store: directory. The HTML authorization form remains operational for the duration defined with the <b>DataPower State Lifetime</b> property. If the user does not submit the request before the duration elapses, the authorization from the OAuth client is rejected.</p>",
							Computed:            true,
						},
						"dp_state_life_time": schema.Int64Attribute{
							MarkdownDescription: "<p>Sets the operational duration in seconds for the local authorization page. Enter a value in the range 1 - 600. The default value is 300.</p><p>If the user does not submit the request before the duration elapses, the authorization request from the OAuth client is rejected. The location of the stylesheet or GatewayScript file that defines the local authorization page and the error handling is set with the <b>Authorization Form</b> property.</p>",
							Computed:            true,
						},
						"au_code_life_time": schema.Int64Attribute{
							MarkdownDescription: "<p>Sets the lifetime for an authorization code in seconds. Enter a value in the range 1 - 600. The default value is 300.</p><p>An authorization code is the intermediary result of a successful authorization. The client uses authorization codes to obtain the access token. Instead of sending tokens to a client, clients receives authorization codes on their redirection URI. Each supported redirection URI for the client is defined with the <b>Redirect URI</b> property.</p>",
							Computed:            true,
						},
						"access_token_life_time": schema.Int64Attribute{
							MarkdownDescription: "Sets the lifetime for the access token in seconds. Enter a value in the range 1 - 63244800. The default value is 3600.",
							Computed:            true,
						},
						"refresh_token_allowed": schema.Int64Attribute{
							MarkdownDescription: "<p>Sets the maximum number of refresh tokens that can be generated for a specific permission set. A permission set is defined as a combination of the resource owner, application, and protected resources. For example, <tt>jack,mobileapp1,scope1</tt> and <tt>john,mobileapp1,scope1</tt> are different permission sets. When an application asks the resource owner for access to protected resources again, the application receives a new permission set with its own counter for refresh tokens.</p><p>Enter a value in the range 0 - 4096. The default value is 0. Remember that refresh tokens and access tokens are distributed in pairs.</p>",
							Computed:            true,
						},
						"refresh_token_life_time": schema.Int64Attribute{
							MarkdownDescription: "Sets the lifetime for the refresh token in seconds. Enter a value in the range 2 - 252979200. The default value is 5400. The lifetime for a refresh token must be longer than that for the corresponding access token.",
							Computed:            true,
						},
						"max_consent_life_time": schema.Int64Attribute{
							MarkdownDescription: "<p>Maximum lifetime that the permission is valid before the application must gather consent again. Enter a value in the range 0 - 2529792000. The default value is 0, which disables this feature.</p>",
							Computed:            true,
						},
						"custom_resource_owner": schema.BoolAttribute{
							MarkdownDescription: "<p>Indicates whether to use a stylesheet or GatewayScript file to extract information about the resource owner. When extracting using custom processing, use the <b>Resource Owner Process</b> property to specify the location of the file. The stylesheet or GatewayScript file must be in the local: or store: directory.</p><p>By default, the resource owner is the user name from the extracted identity. For custom handling, you must provide a stylesheet or GatewayScript file that overrides information about the resource owner.</p><ul><li>For AAA identity extraction, the extraction method can be basic authentication or forms-based login.</li><li>For custom handling, the stylesheet or GatewayScript file overrides data about the resource owner with information from authentication.</li></ul><p>You should use custom handling in the following situations: <ul><li>When presenting the authorization form to the resource owner</li><li>When issuing a code for an authorization code grant type</li><li>When issuing an access token</li></ul></p>",
							Computed:            true,
						},
						"resource_owner_url": schema.StringAttribute{
							MarkdownDescription: "Specifies the location of the stylesheet or GatewayScript file to extract information about the resource owner. The file must be in the local: or store: directory.",
							Computed:            true,
						},
						"additional_oauth_process_url": schema.StringAttribute{
							MarkdownDescription: "<p>Specifies the location of the stylesheet or GatewayScript file to process after generating a code, after generating an access token, or after generating an access token but before sending it to the resource server. The stylesheet or GatewayScript file must be in the local: or store: directory.</p><p>You can use custom additional OAuth processing in the following situations.</p><ul><li>An authorization form request allows custom processing to handle the consent form with the <tt>authorization_form</tt> operation. This operations allows custom handling of the consent form.</li><li>An authorization request after successfully generating a code for an authorization code grant with the <tt>authorization_request</tt> operation. Processing returns a node set. This information becomes part of the query string and is returned to the OAuth client during authorization code grant type.</li><li>An access request after successfully generating an access token with the <tt>access_request</tt> operation. Processing returns a node set. This information becomes part of the JSON object that contains the access token.</li><li>A resource request after successfully verifying an access token but before sending the request to the resource server with the <tt>resource_request</tt> operation.</li><li>A revoke request allows custom handling of a revocation request with the <tt>revoke_request</tt> operation. For example, this operation provides a way to persist the revocation information in a persistent store off the DataPower Gateway.</li><li>A check revocation request verifies whether an access request was revoked previously with the <tt>check_revocation_request</tt> operation. For example, this operation can be used to check against the persistent store off the DataPower Gateway to determine whether an access permission was revoked previously.</li><li>A pre-approval request allows the consent form to be by-passed in either an authorization code or implicit grant type with the <tt>preapproved_check</tt> operation. Depending on the result of this operation, the client's request is approved, denied, or the consent form to be presented.</li><li>A validation request allows custom handling of a validation request grant type with the <tt>validate_request</tt> operation. The response must be in a node set that can be converted into a JSON response, in responding to a validation request.</li><li>A miscinfo request allows the OAuth client to add miscellaneous information to the token with the <tt>miscinfo_request</tt> operation. The authorization server adds the response to the token and returns it to the OAuth client. The maximum number of characters in this information is 512.</li></ul>",
							Computed:            true,
						},
						"rs_set_header": models.GetDmOAuthRSSetHeaderDataSourceSchema("Identifies which HTTP headers to create and send to the remote resource server.", "rs-set-header", ""),
						"validation_url_ssl_client_type": schema.StringAttribute{
							MarkdownDescription: "The TLS profile type to secure connections between the DataPower Gateway and its targets.",
							Computed:            true,
						},
						"validation_url_ssl_client": schema.StringAttribute{
							MarkdownDescription: "Specifies the TLS Client Profile for the validation URL.",
							Computed:            true,
						},
						"jwt_grant_validator": schema.StringAttribute{
							MarkdownDescription: "<p>Specify the JWT validator configuration to verify a JWT for JWT authorization grant. The JWT validator configuration must meet the following requirements.</p><ul><li>Must treat the \"sub\" claim as the resource owner.</li><li>Must check the \"iss\" claim.</li><li>Must check the \"aud\" claim. The \"aud\" claim can be a client ID or the redirect URI.</li><li>Must be configured to verify a signed JWT.</li></ul>",
							Computed:            true,
						},
						"client_jwt_validator": schema.StringAttribute{
							MarkdownDescription: "<p>Specify the JWT validator configuration to verify the client credentials. The JWT validator configuration must meet the following requirements.</p><ul><li>The \"sub\" claim must be the same as client ID.</li><li>Must check the \"iss\" claim.</li><li>Must check the \"aud\" claim.</li><li>Must be configured to verify a signed JWT.</li></ul>",
							Computed:            true,
						},
						"oidc_id_token_generator": schema.StringAttribute{
							MarkdownDescription: "<p>Specify the JWT generator configuration that generates an ID token. The JWT generator configuration must meet the following requirements.</p><ul><li>Must configure \"Issuer\" for the \"iss\" claim.</li><li>Must support \"Issued at\" for the \"iat\" claim.</li><li>Must support signing of the JWT.</li></ul><p>The following items are added to the JWT.</p><ul><li>Authenticated resource owner is added as the value of the \"sub\" claim.</li><li>Client ID is added as part of the \"aud\" claim.</li><li>\"Validity period\" is used to generate the value of the \"exp\" claim.</li><li>Requested \"nonce\" is used for the \"nonce\" claim.</li></ul>",
							Computed:            true,
						},
						"oauth_features":     models.GetDmOAuthFeaturesDataSourceSchema("Specify which features to enable.", "oauth-features", ""),
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *OAuthSupportedClientDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *OAuthSupportedClientDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data OAuthSupportedClientWOList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.OAuthSupportedClientWO{
		AppDomain: data.AppDomain,
	}

	path := o.GetPath()
	if !data.Id.IsNull() {
		path = path + "/" + data.Id.ValueString()
	}

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
	l := []models.OAuthSupportedClientWO{}
	if resFound {
		if value := res.Get(`OAuthSupportedClient`); value.Exists() {
			for _, v := range value.Array() {
				item := models.OAuthSupportedClientWO{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.OAuthSupportedClientObjectTypeWO}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.OAuthSupportedClientObjectTypeWO})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
