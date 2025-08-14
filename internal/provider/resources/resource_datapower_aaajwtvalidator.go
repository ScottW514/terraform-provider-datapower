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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &AAAJWTValidatorResource{}

func NewAAAJWTValidatorResource() resource.Resource {
	return &AAAJWTValidatorResource{}
}

type AAAJWTValidatorResource struct {
	pData *tfutils.ProviderData
}

func (r *AAAJWTValidatorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaajwtvalidator"
}

func (r *AAAJWTValidatorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("<p>The JSON Web Token (JWT) Validator specifies credentials and different methods to validate a JWT.</p>", "jwt-validator", "").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("A descriptive summary for the JWT Validator configuration.", "summary", "").String,
				Optional:            true,
			},
			"issuer": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The optional issuer claim. The \"iss\" PCRE can be used to verify the JWT. The maximum length of the value is 256 characters.", "iss", "").String,
				Optional:            true,
			},
			"aud": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The optional audience claim. The \"aud\" PCRE can be used to verify the JWT. The maximum length of the value is 256 characters.", "aud", "").String,
				Optional:            true,
			},
			"val_method": models.GetDmJWTValMethodResourceSchema("Various methods can be used to validate the JWT. You can decrypt the JWT, verify the JWT signature, and process a custom GatewayScript or XSLT file for further processing.", "validate-method", "", false),
			"customized_script": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("A custom GatewayScript or XSLT file is processed to validate the JWT. The GatewayScript or XSLT file must be stored in the <tt>local:</tt> (the default) or <tt>store:</tt> directory. This field is meaningful when you select <tt>Custom processing</tt> in the Validation method field.", "customized-script", "").String,
				Optional:            true,
			},
			"decrypt_credential_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Various decryption methods (such as PKIX, shared secret key, JSON Web Key (JWK), custom processing, remotely retrieve JWK) can be used to decrypt the JWT. The default method is PKIX. This field is meaningful when you select <tt>Decrypt</tt> in the Validation method field.", "decrypt-type", "").AddStringEnum("pkix", "ssecret", "jwk", "jwk-remote", "custom").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("pkix", "ssecret", "jwk", "jwk-remote", "custom"),
				},
			},
			"decrypt_key": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The private key can be used to decrypt the JWT. You can get the key alias by configuring the Crypto Key. This field is meaningful when you select <tt>Decrypt</tt> in the Validation Method field and choose <tt>PKIX</tt> from the Decrypt method list.", "decrypt-key", "cryptokey").String,
				Optional:            true,
			},
			"decrypt_s_secret": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The shared secret key can be used to decrypt the JWT. You can get the shared secret key alias by configuring the Crypto Shared Secret Key. This field is meaningful when you select <tt>Decrypt</tt> in the Validation method field and choose <tt>Shared secret</tt> from the Decrypt method list.", "decrypt-ssecret", "cryptosskey").String,
				Optional:            true,
			},
			"decrypt_jwk": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The file containing the JWK or key set is fetched to decrypt the JWT. The file must be stored in the local: or store: directory. This field is meaningful when you select <tt>Decrypt</tt> in the Validation method field and choose <tt>JWK</tt> from the Decrypt method list.", "decrypt-jwk", "").String,
				Optional:            true,
			},
			"decrypt_fetch_cred_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The URL indicates the source location where the decryption credentials can be fetched for decrypting the JWT. The URL must be in the format of http or https. By default, the URL is http://example.com/v3/key. This field is meaningful when you choose <tt>Decrypt</tt> in the Validation method field and choose <tt>Remotely retrieve JWK</tt> from the Decrypt method list.", "decrypt-fetch-cred-url", "").AddDefaultValue("http://example.com/v3/key").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("http://example.com/v3/key"),
			},
			"decrypt_fetch_cred_ssl_profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The TLS client profile is specified for fetching the decryption credentials. This field is meaningful when you select <tt>Decrypt</tt> in the Validation method field and choose <tt>Remotely retrieve JWK</tt> from the Decrypt method list.", "decrypt-fetch-cred-sslprofile", "sslclientprofile").String,
				Optional:            true,
			},
			"validate_custom": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("A custom GatewayScript or XSLT file provides the key material information to decrypt or verify the JWT. This field is meaningful when you select <tt>Custom</tt> for the Decrypt method or Verify method list.", "validate-custom", "").String,
				Optional:            true,
			},
			"verify_credential_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Various methods (such as PKIX, shared secret key, JWK, custom processing, remotely retrieve JWK) can be used to verify the JWT signature. The default method is PKIX. This field is meaningful when you select <tt>Verify</tt> in the Validation method field.", "verify-type", "").AddStringEnum("pkix", "ssecret", "jwk", "jwk-remote", "custom").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("pkix", "ssecret", "jwk", "jwk-remote", "custom"),
				},
			},
			"verify_certificate": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The certificate can be used to verify the JWT signature. You can get the certificate by configuring the Crypto Certificate. This field is meaningful when you select <tt>Verify</tt> in the Validation method field and choose <tt>PKIX</tt> from the Verify method field.", "verify-certificate", "cryptocertificate").String,
				Optional:            true,
			},
			"verify_certificate_against_val_cred": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("You decide whether to use validation credentials to verify the JWT signature. This field is meaningful when you select <tt>Verify</tt> in the Validation method field and choose <tt>PKIX</tt> from the Verify method list.", "verify-certificate-against-valcred", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"verify_val_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The validation credentials can be used to verify the signers' certificate for the JWT. You can get credentials by configuring the Crypto Validation Credentials. This field is meaningful when you select <tt>on</tt> in the Signature validation credentials field.", "valcred", "cryptovalcred").String,
				Optional:            true,
			},
			"verify_s_secret": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The shared secret key can be used to verify the JWT signature. This field is meaningful when you select <tt>Verify</tt> in the Validation method field and choose <tt>Shared secret</tt> from the Verify method list.", "verify-ssecret", "cryptosskey").String,
				Optional:            true,
			},
			"verify_jwk": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The file containing the JWK or key set is fetched to verify the JWT signature. The file must be stored in the local: or store: directory. This field is meaningful when you select <tt>Verify</tt> in the Validation method field and choose <tt>JWK</tt> from the Verify method list.", "verify-jwk", "").String,
				Optional:            true,
			},
			"verify_fetch_cred_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The URL indicates the source location where the verification credentials can be fetched for verifying the JWT signature. The URL must be in the format of http or https. By default, the URL is http://example.com/v3/certs. This field is meaningful when you select <tt>Verify</tt> in the Validation method field and choose <tt>Remotely retrieve JWK</tt> from the Verify method list.", "verify-fetch-cred-url", "").AddDefaultValue("http://example.com/v3/certs").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("http://example.com/v3/certs"),
			},
			"verify_fetch_cred_ssl_profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The TLS client profile is provided for fetching the verification credentials. This field is meaningful when you select <tt>Verify</tt> in the Validation method field and choose <tt>Remotely retrieve JWK</tt> from the Verify method list.", "verify-fetch-cred-sslprofile", "sslclientprofile").String,
				Optional:            true,
			},
			"claims": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("JWT claims need to be validated. You must enter the claim name, value, and data type. If the data type is string, specify the PCRE regular expression to verify the claim value.", "claims", "").String,
				NestedObject:        models.DmClaimResourceSchema,
				Optional:            true,
			},
			"username_claim": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("This field is applicable only when the JWT Validator is used in the AAA identity extraction phase. By default, the value of the \"sub\" claim is populated as the username element of the identity extraction output. Ensure that the claim specified in this field is present in the incoming JWT. If no match is found, no username is populated in the AAA processing.", "username-claim", "").AddDefaultValue("sub").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("sub"),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *AAAJWTValidatorResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *AAAJWTValidatorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AAAJWTValidator
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `AAAJWTValidator`)
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

func (r *AAAJWTValidatorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AAAJWTValidator
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
		data.FromBody(ctx, `AAAJWTValidator`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `AAAJWTValidator`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AAAJWTValidatorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AAAJWTValidator
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `AAAJWTValidator`))
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

func (r *AAAJWTValidatorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AAAJWTValidator
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

func (r *AAAJWTValidatorResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.AAAJWTValidator

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
