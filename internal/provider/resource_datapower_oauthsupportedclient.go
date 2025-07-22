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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &OAuthSupportedClientResource{}

func NewOAuthSupportedClientResource() resource.Resource {
	return &OAuthSupportedClientResource{}
}

type OAuthSupportedClientResource struct {
	client *client.DatapowerClient
}

func (r *OAuthSupportedClientResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_oauthsupportedclient"
}

func (r *OAuthSupportedClientResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("OAuth Client Profile", "oauth-supported-client", "").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name of the object. Must be unique among object types in application domain.", "", "").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"app_domain": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The name of the application domain the object belongs to", "", "").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					ImmutableAfterSet(),
				},
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"customized": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Customized OAuth", "customized", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"customized_process_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Customized OAuth Process", "customized-process-url", "").String,
				Optional:            true,
			},
			"o_auth_role": models.GetDmOAuthRoleResourceSchema("OAuth Role", "oauth-role", "", false),
			"az_grant":    models.GetDmOAuthAZGrantTypeResourceSchema("Supported Type", "az-grant", "", false),
			"client_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Client Type", "client-type", "").AddStringEnum("confidential", "public").AddDefaultValue("confidential").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("confidential", "public"),
				},
				Default: stringdefault.StaticString("confidential"),
			},
			"check_client_credential": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Verify Client Credential", "check-client-credential", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"use_validation_url": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use Validation URL", "use-validation-url", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"client_authen_method": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Authentication Method", "client-authen-method", "").AddStringEnum("secret", "ssl", "jwt").AddDefaultValue("secret").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("secret", "ssl", "jwt"),
				},
				Default: stringdefault.StaticString("secret"),
			},
			"client_val_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Client TLS Credentials", "client-valcred", "cryptovalcred").String,
				Optional:            true,
			},
			"generate_client_secret": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Generate Client Secret", "generate-client-secret", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"client_secret": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Client Secret", "client-secret", "").String,
				Optional:            true,
			},
			"caching": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Caching", "caching", "").AddStringEnum("replay", "system", "custom", "diststore").AddDefaultValue("replay").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("replay", "system", "custom", "diststore"),
				},
				Default: stringdefault.StaticString("replay"),
			},
			"validation_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validation URL", "validation-url", "").String,
				Optional:            true,
			},
			"validation_features": models.GetDmValidationFeaturesResourceSchema("Validation Grant Features", "validation-features", "", false),
			"redirect_uri": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Redirect URI", "redirect-uri", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"custom_scope_check": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Customized Scope Check", "custom-scope-check", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"scope": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Scope", "scope", "").String,
				Optional:            true,
			},
			"scope_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Scope Customized Process", "scope-url", "").String,
				Optional:            true,
			},
			"default_scope": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Default Scope", "default-scope", "").String,
				Optional:            true,
			},
			"token_secret": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Shared Secret", "token-secret", "cryptosskey").String,
				Optional:            true,
			},
			"local_az_page_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Authorization Form", "local-az-page-url", "").AddDefaultValue("store:///OAuth-Generate-HTML.xsl").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("store:///OAuth-Generate-HTML.xsl"),
			},
			"dp_state_life_time": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("DataPower State Lifetime", "dp-state-lifetime", "").AddIntegerRange(1, 600).AddDefaultValue("300").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 600),
				},
				Default: int64default.StaticInt64(300),
			},
			"au_code_life_time": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Authorization Grant Lifetime", "au-code-lifetime", "").AddIntegerRange(1, 600).AddDefaultValue("300").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 600),
				},
				Default: int64default.StaticInt64(300),
			},
			"access_token_life_time": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Access Token Lifetime", "access-token-lifetime", "").AddIntegerRange(1, 63244800).AddDefaultValue("3600").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 63244800),
				},
				Default: int64default.StaticInt64(3600),
			},
			"refresh_token_allowed": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Number of Refresh Token Allowed", "refresh-token-allowed", "").AddIntegerRange(0, 4096).String,
				Optional:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 4096),
				},
			},
			"refresh_token_life_time": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Refresh Token Lifetime", "refresh-token-lifetime", "").AddIntegerRange(2, 252979200).AddDefaultValue("5400").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(2, 252979200),
				},
				Default: int64default.StaticInt64(5400),
			},
			"max_consent_life_time": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Maximum Consent Lifetime", "max-consent-lifetime", "").AddIntegerRange(0, 2529792000).String,
				Optional:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 2529792000),
				},
			},
			"custom_resource_owner": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Custom Resource Owner Handling", "custom-resource-owner", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"resource_owner_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Resource Owner Process", "resource-owner-url", "").String,
				Optional:            true,
			},
			"additional_o_auth_process_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Additional OAuth Process", "additional-oauth-process-url", "").String,
				Optional:            true,
			},
			"rs_set_header": models.GetDmOAuthRSSetHeaderResourceSchema("Create HTTP Headers for", "rs-set-header", "", false),
			"validation_urlssl_client_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validation TLS client type", "validation-url-ssl-client-type", "").AddStringEnum("proxy", "client").AddDefaultValue("client").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("proxy", "client"),
				},
				Default: stringdefault.StaticString("client"),
			},
			"validation_urlssl_client": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validation TLS Client Profile", "validation-url-ssl-client", "sslclientprofile").String,
				Optional:            true,
			},
			"jwt_grant_validator": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Authorization grant JWT validator", "jwt-grant-validator", "aaajwtvalidator").String,
				Optional:            true,
			},
			"client_jwt_validator": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Client authentication JWT validator", "client-jwt-validator", "aaajwtvalidator").String,
				Optional:            true,
			},
			"oidcid_token_generator": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("ID token JWT generator", "idtoken-generator", "aaajwtgenerator").String,
				Optional:            true,
			},
			"o_auth_features": models.GetDmOAuthFeaturesResourceSchema("Features", "oauth-features", "", false),
		},
	}
}

func (r *OAuthSupportedClientResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *OAuthSupportedClientResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.OAuthSupportedClient

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `OAuthSupportedClient`)
	res, err := r.client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}

	res, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OAuthSupportedClientResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.OAuthSupportedClient

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && (strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
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
	var data models.OAuthSupportedClient

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `OAuthSupportedClient`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s, %s", err, res.String()))
		return
	}
	res, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OAuthSupportedClientResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.OAuthSupportedClient

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	resp.State.RemoveResource(ctx)
}
