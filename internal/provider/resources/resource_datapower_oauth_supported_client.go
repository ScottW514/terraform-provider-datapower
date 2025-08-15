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
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
)

var _ resource.ResourceWithModifyPlan = &OAuthSupportedClientResource{}

func NewOAuthSupportedClientResource() resource.Resource {
	return &OAuthSupportedClientResource{}
}

type OAuthSupportedClientResource struct {
	pData *tfutils.ProviderData
}

func (r *OAuthSupportedClientResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_oauth_supported_client"
}

func (r *OAuthSupportedClientResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("<p>An OAuth client profile is the basic building block for an OAuth client group. When you create an OAuth client profile, you define its role. As you select the role, the WebGUI refreshes to display the appropriate properties.</p><p>You can create the following types of OAuth client profiles. <ul><li>A client profile for authorization server endpoints: authorization endpoint and token endpoint.</li><li>A client profile for the enforcement point for the resource server.</li><li>A client profile for both authorization server endpoints and the enforcement point.</li></ul></p><p>When creating an OAuth client profile, you can use stylesheets or GatewayScript files for customization.</p><p>You can create a customized OAuth client profile that defines any combination of roles. Customization uses stylesheets or GatewayScript files that must be in the local: or store: directory. For information about the operations that these stylesheets or GatewayScript files must define, see the topic in IBM Knowledge Center.</p>", "oauth-supported-client", "").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies a brief comment that describes the configuration.", "summary", "").String,
				Optional:            true,
			},
			"customized": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Indicates whether the configuration is for a customized OAuth client. The configuration of the customized OAuth client is defined in a stylesheet or GatewayScript file in the local: or store: directory.", "customized", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"customized_process_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specifies the location of the stylesheet or GatewayScript file that defines the customized OAuth client. The stylesheet or GatewayScript file must be in the local: or store: directory</p><p>When creating a customized OAuth client, the stylesheets or GatewayScript files must define all implementation details based on the role of the client. For information about these stylesheets or GatewayScript files, see the topic in IBM Knowledge Center.</p><p>You should provide support for the following operations: <ul><li>verify-az-request: determine whether the initial OAuth request is supported or not.</li><li>issue-az-code: issue a temporary authorization code (for the authorization code grant).</li><li>verify-az-code: verify a temporary authorization.</li><li>issue-access-token: issue an access_token.</li><li>verify-access-token: verify an access_token.</li><li>verify-refresh-token: verify a refresh_token.</li><li>client-revoke-request: handle client revocation request.</li><li>owner-revoke-request: handle owner revocation request.</li></ul></p>", "customized-process-url", "").String,
				Optional:            true,
			},
			"o_auth_role": models.GetDmOAuthRoleResourceSchema("Identifies the role of the client when interacting with a request to access a protected resource.", "oauth-role", "", false),
			"az_grant":    models.GetDmOAuthAZGrantTypeResourceSchema("Identifies the method to obtain the access token for authorization based on the grant type.", "az-grant", "", false),
			"client_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Sets the type of client based on its ability to authenticate securely with authorization server endpoints. The client type is based on the definitions that the authorization server endpoints use for secure authentication and acceptable exposure of client credentials. If the client can securely authenticate, its classification is <tt>confidential</tt> .", "client-type", "").AddStringEnum("confidential", "public").AddDefaultValue("confidential").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("confidential", "public"),
				},
				Default: stringdefault.StaticString("confidential"),
			},
			"check_client_credential": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Identifies whether to verify the client credentials when the DataPower Gateway protects the resource server by using access tokens.", "check-client-credential", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"use_validation_url": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Uses a remote URL to validate the access token.", "use-validation-url", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"client_authen_method": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Identifies the method to authenticate this client.", "client-authen-method", "").AddStringEnum("secret", "ssl", "jwt").AddDefaultValue("secret").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("secret", "ssl", "jwt"),
				},
				Default: stringdefault.StaticString("secret"),
			},
			"client_val_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("An TLS credential used to authenticate the OAuth client sent by remote TLS peer during the TLS handshake.", "client-valcred", "crypto_val_cred").String,
				Optional:            true,
			},
			"generate_client_secret": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Indicates whether to generate the client secret for the OAuth client. The specification refers to the client secret as <tt>client_secret</tt> .</p><ul><li>If you generate the passphrase, the passphrase becomes the client secret.</li><li>If you do not generate the passphrase, you must explicitly define the client secret.</li></ul>", "generate-client-secret", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"client_secret": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the client secret for the OAuth client. The specification references the client secret as <tt>client_secret</tt> .", "client-secret", "").String,
				Optional:            true,
				WriteOnly:           true,
				Sensitive:           true,
			},
			"client_secret_update": schema.BoolAttribute{
				MarkdownDescription: "Set to true by provider if the WRITE ONLY value needs to be updated, otherwise provider will force this to false.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				DeprecationMessage:  "This attribute is for INTERNAL PROVIDER USE. Set values are ignored.",
			},
			"caching": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the caching mechanism to be used.", "caching", "").AddStringEnum("replay", "system", "custom", "diststore").AddDefaultValue("replay").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("replay", "system", "custom", "diststore"),
				},
				Default: stringdefault.StaticString("replay"),
			},
			"validation_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the validation url.", "validation-url", "").String,
				Optional:            true,
			},
			"validation_features": models.GetDmValidationFeaturesResourceSchema("Customize how to handle the validation grant type.", "validation-features", "", false),
			"redirect_uri": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Defines redirection URIs that the OAuth client supports to exchange tokens. Specify each redirection URI as a PCRE.</p><p>Redirection URIs help to detect malicious clients and prevent phishing attacks. The authorization endpoint must have the registered redirection URIs before the authorization endpoint can validate the authorization request from the client. For mobile applications, the redirection URI can be an application URL; for example, <tt>mobiletrafficapp://</tt> that is defined with the <tt>^mobiletrafficapp:\\/\\/?</tt> PCRE.</p>", "redirect-uri", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"custom_scope_check": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Indicates how to check the scope for authorization grants and access tokens.</p><ul><li>When checking the scope with custom processing, specify the location of the stylesheet or GatewayScript file with the <b>Scope Customized Process</b> property. The stylesheet or GatewayScript file must be in the local: or store: directory.</li><li>When checking the scope with a PCRE, specify the expression with the <b>Scope</b> property.</li></ul><p>You should use a custom scope check in the following situations. <ul><li>An authorization request where the OAuth client requests an authorization code.</li><li>An access request where the OAuth client requests an access token.</li><li>A resource request where the OAuth client requests a resource.</li></ul></p>", "custom-scope-check", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"scope": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the PCRE to check the scope. The minimum length of the expression is 1 character. The maximum length of the expression is 1023 characters.", "scope", "").String,
				Optional:            true,
			},
			"scope_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the location of the stylesheet or GatewayScript file for a custom scope check. The stylesheet or GatewayScript file must be in the local: or store: directory. The stylesheet or GatewayScript file validates and sets the scope to check.", "scope-url", "").String,
				Optional:            true,
			},
			"default_scope": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the default value of the scope if the client does not define any scope value in the request.", "default-scope", "").String,
				Optional:            true,
			},
			"token_secret": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Assigns the shared secret key to protect tokens that use the OAuth protocol. The shared secret must be at least 32 bytes in length.", "token-secret", "crypto_sskey_").String,
				Optional:            true,
			},
			"local_az_page_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specifies the location of the stylesheet or GatewayScript file that generates the authorization form for the resource owner and handles errors. The file must be in the local: or store: directory. You can use the <tt>OAuth-Generate-HTML.xsl</tt> stylesheet in the store: directory or copy this file to the local: directory and modify as needed.</p><p>The stylesheet or GatewayScript file must be on the DataPower Gateway in the local: or store: directory. The HTML authorization form remains operational for the duration defined with the <b>DataPower State Lifetime</b> property. If the user does not submit the request before the duration elapses, the authorization from the OAuth client is rejected.</p>", "local-az-page-url", "").AddDefaultValue("store:///OAuth-Generate-HTML.xsl").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("store:///OAuth-Generate-HTML.xsl"),
			},
			"dp_state_life_time": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Sets the operational duration in seconds for the local authorization page. Enter a value in the range 1 - 600. The default value is 300.</p><p>If the user does not submit the request before the duration elapses, the authorization request from the OAuth client is rejected. The location of the stylesheet or GatewayScript file that defines the local authorization page and the error handling is set with the <b>Authorization Form</b> property.</p>", "dp-state-lifetime", "").AddIntegerRange(1, 600).AddDefaultValue("300").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 600),
				},
				Default: int64default.StaticInt64(300),
			},
			"au_code_life_time": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Sets the lifetime for an authorization code in seconds. Enter a value in the range 1 - 600. The default value is 300.</p><p>An authorization code is the intermediary result of a successful authorization. The client uses authorization codes to obtain the access token. Instead of sending tokens to a client, clients receives authorization codes on their redirection URI. Each supported redirection URI for the client is defined with the <b>Redirect URI</b> property.</p>", "au-code-lifetime", "").AddIntegerRange(1, 600).AddDefaultValue("300").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 600),
				},
				Default: int64default.StaticInt64(300),
			},
			"access_token_life_time": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Sets the lifetime for the access token in seconds. Enter a value in the range 1 - 63244800. The default value is 3600.", "access-token-lifetime", "").AddIntegerRange(1, 63244800).AddDefaultValue("3600").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 63244800),
				},
				Default: int64default.StaticInt64(3600),
			},
			"refresh_token_allowed": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Sets the maximum number of refresh tokens that can be generated for a specific permission set. A permission set is defined as a combination of the resource owner, application, and protected resources. For example, <tt>jack,mobileapp1,scope1</tt> and <tt>john,mobileapp1,scope1</tt> are different permission sets. When an application asks the resource owner for access to protected resources again, the application receives a new permission set with its own counter for refresh tokens.</p><p>Enter a value in the range 0 - 4096. The default value is 0. Remember that refresh tokens and access tokens are distributed in pairs.</p>", "refresh-token-allowed", "").AddIntegerRange(0, 4096).String,
				Optional:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 4096),
				},
			},
			"refresh_token_life_time": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Sets the lifetime for the refresh token in seconds. Enter a value in the range 2 - 252979200. The default value is 5400. The lifetime for a refresh token must be longer than that for the corresponding access token.", "refresh-token-lifetime", "").AddIntegerRange(2, 252979200).AddDefaultValue("5400").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(2, 252979200),
				},
				Default: int64default.StaticInt64(5400),
			},
			"max_consent_life_time": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Maximum lifetime that the permission is valid before the application must gather consent again. Enter a value in the range 0 - 2529792000. The default value is 0, which disables this feature.</p>", "max-consent-lifetime", "").AddIntegerRange(0, 2529792000).String,
				Optional:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 2529792000),
				},
			},
			"custom_resource_owner": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Indicates whether to use a stylesheet or GatewayScript file to extract information about the resource owner. When extracting using custom processing, use the <b>Resource Owner Process</b> property to specify the location of the file. The stylesheet or GatewayScript file must be in the local: or store: directory.</p><p>By default, the resource owner is the user name from the extracted identity. For custom handling, you must provide a stylesheet or GatewayScript file that overrides information about the resource owner.</p><ul><li>For AAA identity extraction, the extraction method can be basic authentication or forms-based login.</li><li>For custom handling, the stylesheet or GatewayScript file overrides data about the resource owner with information from authentication.</li></ul><p>You should use custom handling in the following situations: <ul><li>When presenting the authorization form to the resource owner</li><li>When issuing a code for an authorization code grant type</li><li>When issuing an access token</li></ul></p>", "custom-resource-owner", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"resource_owner_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the location of the stylesheet or GatewayScript file to extract information about the resource owner. The file must be in the local: or store: directory.", "resource-owner-url", "").String,
				Optional:            true,
			},
			"additional_o_auth_process_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specifies the location of the stylesheet or GatewayScript file to process after generating a code, after generating an access token, or after generating an access token but before sending it to the resource server. The stylesheet or GatewayScript file must be in the local: or store: directory.</p><p>You can use custom additional OAuth processing in the following situations.</p><ul><li>An authorization form request allows custom processing to handle the consent form with the <tt>authorization_form</tt> operation. This operations allows custom handling of the consent form.</li><li>An authorization request after successfully generating a code for an authorization code grant with the <tt>authorization_request</tt> operation. Processing returns a node set. This information becomes part of the query string and is returned to the OAuth client during authorization code grant type.</li><li>An access request after successfully generating an access token with the <tt>access_request</tt> operation. Processing returns a node set. This information becomes part of the JSON object that contains the access token.</li><li>A resource request after successfully verifying an access token but before sending the request to the resource server with the <tt>resource_request</tt> operation.</li><li>A revoke request allows custom handling of a revocation request with the <tt>revoke_request</tt> operation. For example, this operation provides a way to persist the revocation information in a persistent store off the DataPower Gateway.</li><li>A check revocation request verifies whether an access request was revoked previously with the <tt>check_revocation_request</tt> operation. For example, this operation can be used to check against the persistent store off the DataPower Gateway to determine whether an access permission was revoked previously.</li><li>A pre-approval request allows the consent form to be by-passed in either an authorization code or implicit grant type with the <tt>preapproved_check</tt> operation. Depending on the result of this operation, the client's request is approved, denied, or the consent form to be presented.</li><li>A validation request allows custom handling of a validation request grant type with the <tt>validate_request</tt> operation. The response must be in a node set that can be converted into a JSON response, in responding to a validation request.</li><li>A miscinfo request allows the OAuth client to add miscellaneous information to the token with the <tt>miscinfo_request</tt> operation. The authorization server adds the response to the token and returns it to the OAuth client. The maximum number of characters in this information is 512.</li></ul>", "additional-oauth-process-url", "").String,
				Optional:            true,
			},
			"rs_set_header": models.GetDmOAuthRSSetHeaderResourceSchema("Identifies which HTTP headers to create and send to the remote resource server.", "rs-set-header", "", false),
			"validation_urlssl_client_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The TLS profile type to secure connections between the DataPower Gateway and its targets.", "validation-url-ssl-client-type", "").AddStringEnum("client").AddDefaultValue("client").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("client"),
				},
				Default: stringdefault.StaticString("client"),
			},
			"validation_urlssl_client": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the TLS Client Profile for the validation URL.", "validation-url-ssl-client", "ssl_client_profile").String,
				Optional:            true,
			},
			"jwt_grant_validator": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the JWT validator configuration to verify a JWT for JWT authorization grant. The JWT validator configuration must meet the following requirements.</p><ul><li>Must treat the \"sub\" claim as the resource owner.</li><li>Must check the \"iss\" claim.</li><li>Must check the \"aud\" claim. The \"aud\" claim can be a client ID or the redirect URI.</li><li>Must be configured to verify a signed JWT.</li></ul>", "jwt-grant-validator", "aaa_jwt_validator").String,
				Optional:            true,
			},
			"client_jwt_validator": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the JWT validator configuration to verify the client credentials. The JWT validator configuration must meet the following requirements.</p><ul><li>The \"sub\" claim must be the same as client ID.</li><li>Must check the \"iss\" claim.</li><li>Must check the \"aud\" claim.</li><li>Must be configured to verify a signed JWT.</li></ul>", "client-jwt-validator", "aaa_jwt_validator").String,
				Optional:            true,
			},
			"oidcid_token_generator": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the JWT generator configuration that generates an ID token. The JWT generator configuration must meet the following requirements.</p><ul><li>Must configure \"Issuer\" for the \"iss\" claim.</li><li>Must support \"Issued at\" for the \"iat\" claim.</li><li>Must support signing of the JWT.</li></ul><p>The following items are added to the JWT.</p><ul><li>Authenticated resource owner is added as the value of the \"sub\" claim.</li><li>Client ID is added as part of the \"aud\" claim.</li><li>\"Validity period\" is used to generate the value of the \"exp\" claim.</li><li>Requested \"nonce\" is used for the \"nonce\" claim.</li></ul>", "idtoken-generator", "aaa_jwt_generator").String,
				Optional:            true,
			},
			"o_auth_features":    models.GetDmOAuthFeaturesResourceSchema("Specify which features to enable.", "oauth-features", "", false),
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *OAuthSupportedClientResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *OAuthSupportedClientResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.OAuthSupportedClient
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	var config models.OAuthSupportedClient

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	data.ClientSecret = config.ClientSecret
	resp.Diagnostics.Append(resp.Private.SetKey(ctx, "OAuthSupportedClient.ClientSecret", []byte("{\"value\": \""+tfutils.GenerateHash(config.ClientSecret.ValueString())+"\"}"))...)

	body := data.ToBody(ctx, `OAuthSupportedClient`)
	_, err := r.pData.Client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "POST", err))
		return
	}
	getRes, getErr := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if getErr != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object after creation (GET), got error: %s", getErr))
		return
	}
	data.UpdateUnknownFromBody(ctx, `OAuthSupportedClient`, getRes)
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OAuthSupportedClientResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.OAuthSupportedClient
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
		data.FromBody(ctx, `OAuthSupportedClient`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `OAuthSupportedClient`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OAuthSupportedClientResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data, config models.OAuthSupportedClient
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ClientSecretUpdate.IsUnknown() {
		data.ClientSecret = config.ClientSecret
		data.ClientSecretUpdate = types.BoolValue(false)
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "OAuthSupportedClient.ClientSecret", []byte("{\"value\": \""+tfutils.GenerateHash(config.ClientSecret.ValueString())+"\"}"))...)
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `OAuthSupportedClient`))
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

func (r *OAuthSupportedClientResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.OAuthSupportedClient
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
func (r *OAuthSupportedClientResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}
	var data, config models.OAuthSupportedClient

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var val []byte
	var diags diag.Diagnostics
	val, diags = req.Private.GetKey(ctx, "OAuthSupportedClient.ClientSecret")
	resp.Diagnostics.Append(diags...)
	if val != nil {
		if hash := gjson.GetBytes(val, "value"); hash.Exists() && !tfutils.VerifyHash(config.ClientSecret.ValueString(), hash.String()) {
			data.ClientSecretUpdate = types.BoolUnknown()
		}
	}

	resp.Diagnostics.Append(resp.Plan.Set(ctx, &data)...)
}

func (r *OAuthSupportedClientResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.OAuthSupportedClient

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
