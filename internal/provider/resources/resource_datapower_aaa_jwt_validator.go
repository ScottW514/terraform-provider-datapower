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

var _ resource.Resource = &AAAJWTValidatorResource{}
var _ resource.ResourceWithImportState = &AAAJWTValidatorResource{}

func NewAAAJWTValidatorResource() resource.Resource {
	return &AAAJWTValidatorResource{}
}

type AAAJWTValidatorResource struct {
	pData *tfutils.ProviderData
}

func (r *AAAJWTValidatorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaa_jwt_validator"
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
				MarkdownDescription: tfutils.NewAttributeDescription("A descriptive summary for the JWT Validator configuration.", "summary", "").String,
				Optional:            true,
			},
			"issuer": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The optional issuer claim. The \"iss\" PCRE can be used to verify the JWT. The maximum length of the value is 256 characters.", "iss", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 256),
				},
			},
			"aud": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The optional audience claim. The \"aud\" PCRE can be used to verify the JWT. The maximum length of the value is 256 characters.", "aud", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 256),
				},
			},
			"val_method": models.GetDmJWTValMethodResourceSchema("Various methods can be used to validate the JWT. You can decrypt the JWT, verify the JWT signature, and process a custom GatewayScript or XSLT file for further processing.", "validate-method", "", false),
			"customized_script": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("A custom GatewayScript or XSLT file is processed to validate the JWT. The GatewayScript or XSLT file must be stored in the <tt>local:</tt> (the default) or <tt>store:</tt> directory. This field is meaningful when you select <tt>Custom processing</tt> in the Validation method field.", "customized-script", "").AddRequiredWhen(models.AAAJWTValidatorCustomizedScriptCondVal.String()).AddNotValidWhen(models.AAAJWTValidatorCustomizedScriptIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AAAJWTValidatorCustomizedScriptCondVal, models.AAAJWTValidatorCustomizedScriptIgnoreVal, false),
				},
			},
			"decrypt_credential_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Various decryption methods (such as PKIX, shared secret key, JSON Web Key (JWK), custom processing, remotely retrieve JWK) can be used to decrypt the JWT. The default method is PKIX. This field is meaningful when you select <tt>Decrypt</tt> in the Validation method field.", "decrypt-type", "").AddStringEnum("pkix", "ssecret", "jwk", "jwk-remote", "custom").AddRequiredWhen(models.AAAJWTValidatorDecryptCredentialTypeCondVal.String()).AddNotValidWhen(models.AAAJWTValidatorDecryptCredentialTypeIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("pkix", "ssecret", "jwk", "jwk-remote", "custom"),
					validators.ConditionalRequiredString(models.AAAJWTValidatorDecryptCredentialTypeCondVal, models.AAAJWTValidatorDecryptCredentialTypeIgnoreVal, false),
				},
			},
			"decrypt_key": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The private key can be used to decrypt the JWT. You can get the key alias by configuring the Crypto Key. This field is meaningful when you select <tt>Decrypt</tt> in the Validation Method field and choose <tt>PKIX</tt> from the Decrypt method list.", "decrypt-key", "crypto_key").AddRequiredWhen(models.AAAJWTValidatorDecryptKeyCondVal.String()).AddNotValidWhen(models.AAAJWTValidatorDecryptKeyIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AAAJWTValidatorDecryptKeyCondVal, models.AAAJWTValidatorDecryptKeyIgnoreVal, false),
				},
			},
			"decrypt_s_secret": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The shared secret key can be used to decrypt the JWT. You can get the shared secret key alias by configuring the Crypto Shared Secret Key. This field is meaningful when you select <tt>Decrypt</tt> in the Validation method field and choose <tt>Shared secret</tt> from the Decrypt method list.", "decrypt-ssecret", "crypto_sskey").AddRequiredWhen(models.AAAJWTValidatorDecryptSSecretCondVal.String()).AddNotValidWhen(models.AAAJWTValidatorDecryptSSecretIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AAAJWTValidatorDecryptSSecretCondVal, models.AAAJWTValidatorDecryptSSecretIgnoreVal, false),
				},
			},
			"decrypt_jwk": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The file containing the JWK or key set is fetched to decrypt the JWT. The file must be stored in the local: or store: directory. This field is meaningful when you select <tt>Decrypt</tt> in the Validation method field and choose <tt>JWK</tt> from the Decrypt method list.", "decrypt-jwk", "").AddRequiredWhen(models.AAAJWTValidatorDecryptJWKCondVal.String()).AddNotValidWhen(models.AAAJWTValidatorDecryptJWKIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AAAJWTValidatorDecryptJWKCondVal, models.AAAJWTValidatorDecryptJWKIgnoreVal, false),
				},
			},
			"decrypt_fetch_cred_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The URL indicates the source location where the decryption credentials can be fetched for decrypting the JWT. The URL must be in the format of http or https. By default, the URL is http://example.com/v3/key. This field is meaningful when you choose <tt>Decrypt</tt> in the Validation method field and choose <tt>Remotely retrieve JWK</tt> from the Decrypt method list.", "decrypt-fetch-cred-url", "").AddDefaultValue("http://example.com/v3/key").AddRequiredWhen(models.AAAJWTValidatorDecryptFetchCredURLCondVal.String()).AddNotValidWhen(models.AAAJWTValidatorDecryptFetchCredURLIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(7, 255),
					validators.ConditionalRequiredString(models.AAAJWTValidatorDecryptFetchCredURLCondVal, models.AAAJWTValidatorDecryptFetchCredURLIgnoreVal, true),
				},
				Default: stringdefault.StaticString("http://example.com/v3/key"),
			},
			"decrypt_fetch_cred_ssl_profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The TLS client profile is specified for fetching the decryption credentials. This field is meaningful when you select <tt>Decrypt</tt> in the Validation method field and choose <tt>Remotely retrieve JWK</tt> from the Decrypt method list.", "decrypt-fetch-cred-sslprofile", "ssl_client_profile").AddNotValidWhen(models.AAAJWTValidatorDecryptFetchCredSSLProfileIgnoreVal.String()).String,
				Optional:            true,
			},
			"validate_custom": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("A custom GatewayScript or XSLT file provides the key material information to decrypt or verify the JWT. This field is meaningful when you select <tt>Custom</tt> for the Decrypt method or Verify method list.", "validate-custom", "").AddRequiredWhen(models.AAAJWTValidatorValidateCustomCondVal.String()).AddNotValidWhen(models.AAAJWTValidatorValidateCustomIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AAAJWTValidatorValidateCustomCondVal, models.AAAJWTValidatorValidateCustomIgnoreVal, false),
				},
			},
			"verify_credential_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Various methods (such as PKIX, shared secret key, JWK, custom processing, remotely retrieve JWK) can be used to verify the JWT signature. The default method is PKIX. This field is meaningful when you select <tt>Verify</tt> in the Validation method field.", "verify-type", "").AddStringEnum("pkix", "ssecret", "jwk", "jwk-remote", "custom").AddRequiredWhen(models.AAAJWTValidatorVerifyCredentialTypeCondVal.String()).AddNotValidWhen(models.AAAJWTValidatorVerifyCredentialTypeIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("pkix", "ssecret", "jwk", "jwk-remote", "custom"),
					validators.ConditionalRequiredString(models.AAAJWTValidatorVerifyCredentialTypeCondVal, models.AAAJWTValidatorVerifyCredentialTypeIgnoreVal, false),
				},
			},
			"verify_certificate": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The certificate can be used to verify the JWT signature. You can get the certificate by configuring the Crypto Certificate. This field is meaningful when you select <tt>Verify</tt> in the Validation method field and choose <tt>PKIX</tt> from the Verify method field.", "verify-certificate", "crypto_certificate").AddRequiredWhen(models.AAAJWTValidatorVerifyCertificateCondVal.String()).AddNotValidWhen(models.AAAJWTValidatorVerifyCertificateIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AAAJWTValidatorVerifyCertificateCondVal, models.AAAJWTValidatorVerifyCertificateIgnoreVal, false),
				},
			},
			"verify_certificate_against_val_cred": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("You decide whether to use validation credentials to verify the JWT signature. This field is meaningful when you select <tt>Verify</tt> in the Validation method field and choose <tt>PKIX</tt> from the Verify method list.", "verify-certificate-against-valcred", "").AddDefaultValue("false").AddRequiredWhen(models.AAAJWTValidatorVerifyCertificateAgainstValCredCondVal.String()).AddNotValidWhen(models.AAAJWTValidatorVerifyCertificateAgainstValCredIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"verify_val_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The validation credentials can be used to verify the signers' certificate for the JWT. You can get credentials by configuring the Crypto Validation Credentials. This field is meaningful when you select <tt>on</tt> in the Signature validation credentials field.", "valcred", "crypto_val_cred").AddRequiredWhen(models.AAAJWTValidatorVerifyValCredCondVal.String()).AddNotValidWhen(models.AAAJWTValidatorVerifyValCredIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AAAJWTValidatorVerifyValCredCondVal, models.AAAJWTValidatorVerifyValCredIgnoreVal, false),
				},
			},
			"verify_s_secret": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The shared secret key can be used to verify the JWT signature. This field is meaningful when you select <tt>Verify</tt> in the Validation method field and choose <tt>Shared secret</tt> from the Verify method list.", "verify-ssecret", "crypto_sskey").AddRequiredWhen(models.AAAJWTValidatorVerifySSecretCondVal.String()).AddNotValidWhen(models.AAAJWTValidatorVerifySSecretIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AAAJWTValidatorVerifySSecretCondVal, models.AAAJWTValidatorVerifySSecretIgnoreVal, false),
				},
			},
			"verify_jwk": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The file containing the JWK or key set is fetched to verify the JWT signature. The file must be stored in the local: or store: directory. This field is meaningful when you select <tt>Verify</tt> in the Validation method field and choose <tt>JWK</tt> from the Verify method list.", "verify-jwk", "").AddRequiredWhen(models.AAAJWTValidatorVerifyJWKCondVal.String()).AddNotValidWhen(models.AAAJWTValidatorVerifyJWKIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AAAJWTValidatorVerifyJWKCondVal, models.AAAJWTValidatorVerifyJWKIgnoreVal, false),
				},
			},
			"verify_fetch_cred_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The URL indicates the source location where the verification credentials can be fetched for verifying the JWT signature. The URL must be in the format of http or https. By default, the URL is http://example.com/v3/certs. This field is meaningful when you select <tt>Verify</tt> in the Validation method field and choose <tt>Remotely retrieve JWK</tt> from the Verify method list.", "verify-fetch-cred-url", "").AddDefaultValue("http://example.com/v3/certs").AddRequiredWhen(models.AAAJWTValidatorVerifyFetchCredURLCondVal.String()).AddNotValidWhen(models.AAAJWTValidatorVerifyFetchCredURLIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(7, 255),
					validators.ConditionalRequiredString(models.AAAJWTValidatorVerifyFetchCredURLCondVal, models.AAAJWTValidatorVerifyFetchCredURLIgnoreVal, true),
				},
				Default: stringdefault.StaticString("http://example.com/v3/certs"),
			},
			"verify_fetch_cred_ssl_profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The TLS client profile is provided for fetching the verification credentials. This field is meaningful when you select <tt>Verify</tt> in the Validation method field and choose <tt>Remotely retrieve JWK</tt> from the Verify method list.", "verify-fetch-cred-sslprofile", "ssl_client_profile").AddNotValidWhen(models.AAAJWTValidatorVerifyFetchCredSSLProfileIgnoreVal.String()).String,
				Optional:            true,
			},
			"claims": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("JWT claims need to be validated. You must enter the claim name, value, and data type. If the data type is string, specify the PCRE regular expression to verify the claim value.", "claims", "").String,
				NestedObject:        models.GetDmClaimResourceSchema(),
				Optional:            true,
			},
			"username_claim": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("This field is applicable only when the JWT Validator is used in the AAA identity extraction phase. By default, the value of the \"sub\" claim is populated as the username element of the identity extraction output. Ensure that the claim specified in this field is present in the incoming JWT. If no match is found, no username is populated in the AAA processing.", "username-claim", "").AddDefaultValue("sub").String,
				Computed:            true,
				Optional:            true,
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
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), body)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Resource already exists. Failed to update resource, got error: %s", err))
				return
			}
			resp.Diagnostics.AddWarning("Warning", "Resource already exists. Existing resource was updated.")
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString())
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
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

func (r *AAAJWTValidatorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AAAJWTValidator
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil {
		if strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400") {
			resp.State.RemoveResource(ctx)
			return
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString())
			if !resp.Diagnostics.HasError() {
				resp.State.RemoveResource(ctx)
			}
			return
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
			return
		}
	}

	data.UpdateFromBody(ctx, `AAAJWTValidator`, res)

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
		if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString())
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
			return
		}
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
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			resp.Diagnostics.AddWarning("Resource Conflict", fmt.Sprintf("Resource is no longer tracked by Terraform, but may need to be manually deleted on DataPower host. Got error: %s", err))
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString())
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
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

func (r *AAAJWTValidatorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

	var data models.AAAJWTValidator
	data.AppDomain = types.StringValue(appDomain)
	data.Id = types.StringValue(id)
	res, err := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("Resource was not found, got error: %s", err))
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString())
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		}
		return
	}

	data.FromBody(ctx, `AAAJWTValidator`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AAAJWTValidatorResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.AAAJWTValidator

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
