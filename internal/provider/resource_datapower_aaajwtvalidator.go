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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &AAAJWTValidatorResource{}

func NewAAAJWTValidatorResource() resource.Resource {
	return &AAAJWTValidatorResource{}
}

type AAAJWTValidatorResource struct {
	client *client.DatapowerClient
}

func (r *AAAJWTValidatorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaajwtvalidator"
}

func (r *AAAJWTValidatorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("JWT Validator", "jwt-validator", "").String,

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
			"issuer": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Issuer", "iss", "").String,
				Optional:            true,
			},
			"aud": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Audience", "aud", "").String,
				Optional:            true,
			},
			"val_method": models.GetDmJWTValMethodResourceSchema("Validation method", "validate-method", "", false),
			"customized_script": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Custom validation method processing", "customized-script", "").String,
				Optional:            true,
			},
			"decrypt_credential_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Decrypt method", "decrypt-type", "").AddStringEnum("pkix", "ssecret", "jwk", "jwk-remote", "custom").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("pkix", "ssecret", "jwk", "jwk-remote", "custom"),
				},
			},
			"decrypt_key": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Decrypt key", "decrypt-key", "cryptokey").String,
				Optional:            true,
			},
			"decrypt_s_secret": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Decrypt shared secret", "decrypt-ssecret", "cryptosskey").String,
				Optional:            true,
			},
			"decrypt_jwk": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Decrypt JWK", "decrypt-jwk", "").String,
				Optional:            true,
			},
			"decrypt_fetch_cred_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Decrypt credential URL", "decrypt-fetch-cred-url", "").AddDefaultValue("http://example.com/v3/key").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("http://example.com/v3/key"),
			},
			"decrypt_fetch_cred_ssl_profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Decrypt credential TLS client profile", "decrypt-fetch-cred-sslprofile", "sslclientprofile").String,
				Optional:            true,
			},
			"validate_custom": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Custom decrypt/verify processing", "validate-custom", "").String,
				Optional:            true,
			},
			"verify_credential_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Verify method", "verify-type", "").AddStringEnum("pkix", "ssecret", "jwk", "jwk-remote", "custom").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("pkix", "ssecret", "jwk", "jwk-remote", "custom"),
				},
			},
			"verify_certificate": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Verify certificate", "verify-certificate", "cryptocertificate").String,
				Optional:            true,
			},
			"verify_certificate_against_val_cred": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Signature validation credentials", "verify-certificate-against-valcred", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"verify_val_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validation credentials", "valcred", "cryptovalcred").String,
				Optional:            true,
			},
			"verify_s_secret": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Verify shared secret", "verify-ssecret", "cryptosskey").String,
				Optional:            true,
			},
			"verify_jwk": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Verify JWK", "verify-jwk", "").String,
				Optional:            true,
			},
			"verify_fetch_cred_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Verify credential URL", "verify-fetch-cred-url", "").AddDefaultValue("http://example.com/v3/certs").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("http://example.com/v3/certs"),
			},
			"verify_fetch_cred_ssl_profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Verify credential TLS client profile", "verify-fetch-cred-sslprofile", "sslclientprofile").String,
				Optional:            true,
			},
			"claims": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validate claims", "claims", "").String,
				NestedObject:        models.DmClaimResourceSchema,
				Optional:            true,
			},
			"username_claim": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Claim used as username", "username-claim", "").AddDefaultValue("sub").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("sub"),
			},
		},
	}
}

func (r *AAAJWTValidatorResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *AAAJWTValidatorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.AAAJWTValidator

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `AAAJWTValidator`)
	res, err := r.client.Post(data.GetPath(), body)

	if err != nil && res.RawResponse.StatusCode != 409 {
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

func (r *AAAJWTValidatorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.AAAJWTValidator

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 406") || strings.Contains(err.Error(), "StatusCode 500") || strings.Contains(err.Error(), "StatusCode 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	if data.IsNull() {
		// Import
		data.FromBody(ctx, `AAAJWTValidator`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `AAAJWTValidator`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AAAJWTValidatorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.AAAJWTValidator

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `AAAJWTValidator`))
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

func (r *AAAJWTValidatorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.AAAJWTValidator

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && res.RawResponse.StatusCode != 404 && res.RawResponse.StatusCode != 409 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	resp.State.RemoveResource(ctx)
}
