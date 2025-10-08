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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

var _ resource.Resource = &SocialLoginPolicyResource{}
var _ resource.ResourceWithImportState = &SocialLoginPolicyResource{}

func NewSocialLoginPolicyResource() resource.Resource {
	return &SocialLoginPolicyResource{}
}

type SocialLoginPolicyResource struct {
	pData *tfutils.ProviderData
}

func (r *SocialLoginPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_social_login_policy"
}

func (r *SocialLoginPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("DataPower can act as an OpenID Connect client. In this case, a social login policy enables DataPower to redirect the user to a social login provider like Google for user authentication and consent for authorization.", "social-login-policy", "").String,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name of the object. Must be unique among object types in application domain.", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
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
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
				PlanModifiers: []planmodifier.String{
					modifiers.ImmutableAfterSet(),
				},
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter a descriptive summary for the configuration.", "summary", "").String,
				Optional:            true,
			},
			"client_id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the ID of DataPower that is registered with the social login provider.", "client-id", "").String,
				Required:            true,
			},
			"client_secret": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the secret of DataPower that is registered with the social login provider.", "client-secret", "").String,
				Required:            true,
			},
			"client_grant": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Controls how DataPower generates the client request to the social login provider.", "client-grant", "").AddStringEnum("code", "implicit", "password", "client", "code-id_token").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("code", "implicit", "password", "client", "code-id_token"),
				},
			},
			"client_scope": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the scope value that defines what access privileges are requested for access tokens, ID tokens, or both. Use space separated strings. For example, <tt>openid email</tt> .", "client-scope", "").String,
				Optional:            true,
			},
			"client_redirect_uri": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specifies the URI that the social login policy redirects the client to after the client obtains a code or an access token. The URI must match with what is registered at the social login provider for DataPower as the OAuth/OIDC client. The URI is included in the OAuth/OIDC client request that DataPower generates.</p><p>Note that the social login provider Google mandates that the redirect URI must be a fully qualified host name instead of an IP address.</p><p>Note that the redirect URI should end with the suffix '/social-login-callback' in the pathname in order to differentiate between the callback requests and other types of requests coming into the service.</p><p>You can specify the value of this redirect URI in the following forms.</p><ul><li>Static string. Enter a static string as the redirect URI. Must end with the suffix '/social-login-callback'.</li><li>URL-in/suffix. In this case, it takes the value from the inbound URL service variable var://service/URL-in and then suffixes the value with whatever is specified after <tt>URL-in</tt> as the redirect URI. For example, the value of this property is 'URL-in/social-login-callback' and the incoming URL is 'https://datapower.ibm.com:10087/getresources', then the redirect URI is constructed as 'https://datapower.ibm.com:10087/getresources/social-login-callback'.</li><li>Context variable. You can set a context variable before you invoke this AAA action and specify the context variable name for this value. For example, var://context/AAA/social-login-redirect-uri</li></ul>", "client-redirect-uri", "").AddDefaultValue("URL-in/social-login-callback").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 255),
					stringvalidator.RegexMatches(regexp.MustCompile("(^var://)|(/social-login-callback$)"), "Must match :"+"(^var://)|(/social-login-callback$)"),
				},
				Default: stringdefault.StaticString("URL-in/social-login-callback"),
			},
			"client_optional_query_params": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the optional query parameters to include in the initial OAuth/OIDC request that DataPower sends to the social login provider. Enter the optional query parameters as name=value pairs and separate each pair with an ampersand. For example, prompt=consent&amp;login_hint=jsmith@example.com&amp;openid.realm=example.comi&amp;hd=example.com.", "client-opt-query-params", "").String,
				Optional:            true,
			},
			"client_ssl_profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the TLS client profile to secure connections when DataPower obtains an access token from the social login provider.", "client-ssl", "ssl_client_profile").String,
				Required:            true,
			},
			"social_provider": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Controls which social login provider to use.", "provider", "").AddStringEnum("google", "oidc", "facebook", "custom").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("google", "oidc", "facebook", "custom"),
				},
			},
			"provider_az_endpoint": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the provider's endpoint URL that accepts an authorization request from a client to perform social login with the provider. When the provider is Google, you can retrieve the authorization endpoint URL from the Discovery document for Google's OpenID Connect service.", "provider-az-endpoint", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(7, 255),
					stringvalidator.RegexMatches(regexp.MustCompile("^https?://\\[?[-_a-z0-9A-Z.:]+\\]?(:[0-9]+)?/[ -~]*$"), "Must match :"+"^https?://\\[?[-_a-z0-9A-Z.:]+\\]?(:[0-9]+)?/[ -~]*$"),
				},
			},
			"provider_token_endpoint": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the provider's endpoint URL that accepts an authorization grant, or code, from a client in exchange for an access token from the social login provider. When the provider is Google, you can retrieve the token endpoint URL from the Discovery document for Google's OpenID Connect service.", "provider-token-endpoint", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(7, 255),
					stringvalidator.RegexMatches(regexp.MustCompile("^https?://\\[?[-_a-z0-9A-Z.:]+\\]?(:[0-9]+)?/[ -~]*$"), "Must match :"+"^https?://\\[?[-_a-z0-9A-Z.:]+\\]?(:[0-9]+)?/[ -~]*$"),
				},
			},
			"validate_jwt_token": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Controls whether to validate the JWT token (ID token)from the provider. If yes, it is recommended that you validate the ID token that is obtained from Google by defining the following settings in the JWT Validator configuration.</p><p><ol><li>Verify the signature by fetching the certs from https://www.googleapis.com/oauth2/v3/certs</li><li>Verify that the <tt>aud</tt> claim matches the client ID of DataPower.</li><li>Verify that the <tt>iss</tt> claim matches accounts.google.com or https://accounts.google.com</li></ol></p><p>For other recommendations on validating the ID token from Google, see https://developers.google.com/identity/protocols/OpenIDConnect.</p>", "validate-jwt-token", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"jwt_validator": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the JWT Validator configuration that defines how to validate and verify the ID token.", "jwt-validator", "aaa_jwt_validator").AddRequiredWhen(models.SocialLoginPolicyJWTValidatorCondVal.String()).AddNotValidWhen(models.SocialLoginPolicyJWTValidatorIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.SocialLoginPolicyJWTValidatorCondVal, models.SocialLoginPolicyJWTValidatorIgnoreVal, false),
				},
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *SocialLoginPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *SocialLoginPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SocialLoginPolicy
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `SocialLoginPolicy`)
	_, err := r.pData.Client.Post(data.GetPath(), body)
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), body)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Resource already exists. Failed to update resource, got error: %s", err))
				return
			}
			resp.Diagnostics.AddWarning("Warning", "Resource already exists. Existing resource was updated.")
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create resource, got error: %s", err))
			return
		}
	}
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SocialLoginPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SocialLoginPolicy
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

	data.UpdateFromBody(ctx, `SocialLoginPolicy`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SocialLoginPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SocialLoginPolicy
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `SocialLoginPolicy`))
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

func (r *SocialLoginPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SocialLoginPolicy
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			resp.Diagnostics.AddWarning("Resource Conflict", fmt.Sprintf("Resource is no longer tracked by Terraform, but may need to be manually deleted on DataPower host. Got error: %s", err))
		} else if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete resource, got error: %s", err))
			return
		}
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *SocialLoginPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError("Invalid Import ID", "Expected format: <app_domain>/<id>. Got: "+req.ID)
		return
	}

	appDomain := parts[0]
	id := parts[1]

	if !regexp.MustCompile("^[a-zA-Z0-9_-]+$").MatchString(id) || len(id) < 1 || len(id) > 128 {
		resp.Diagnostics.AddError("Invalid ID", "ID must be 1-128 characters and match pattern ^[a-zA-Z0-9_-]+$")
		return
	}
	if !regexp.MustCompile("^[a-zA-Z0-9_-]+$").MatchString(appDomain) || len(appDomain) < 1 || len(appDomain) > 128 {
		resp.Diagnostics.AddError("Invalid Application Domain", "Application domain must be 1-128 characters and match pattern ^[a-zA-Z0-9_-]+$")
		return
	}

	var data models.SocialLoginPolicy
	data.AppDomain = types.StringValue(appDomain)
	data.Id = types.StringValue(id)
	res, err := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("Resource was not found, got error: %s", err))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		}
		return
	}

	data.FromBody(ctx, `SocialLoginPolicy`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SocialLoginPolicyResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.SocialLoginPolicy

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
