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
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

var _ resource.Resource = &AAAJWTGeneratorResource{}

func NewAAAJWTGeneratorResource() resource.Resource {
	return &AAAJWTGeneratorResource{}
}

type AAAJWTGeneratorResource struct {
	pData *tfutils.ProviderData
}

func (r *AAAJWTGeneratorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaa_jwt_generator"
}

func (r *AAAJWTGeneratorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("<p>The JSON Web Token (JWT) Generator specifies the JWT content and the cryptographic methods, such as signing and encryption methods, used for generating a JWT during the AAA postprocessing phase.</p>", "jwt-generator", "").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("A descriptive summary for the JWT Generator configuration.", "summary", "").String,
				Optional:            true,
			},
			"issuer": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The issuer claim, \"iss\", identifies the principal that issues the JWT. The maximum length is 256 characters. The default value is <tt>idg</tt> .", "iss", "").AddDefaultValue("idg").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 256),
				},
				Default: stringdefault.StaticString("idg"),
			},
			"duration": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The validity period identifies the expiration time, \"exp\" claim. Enter a value in the range 1 - 31622400. The default value is 3600.", "duration", "").AddIntegerRange(1, 31622400).AddDefaultValue("3600").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 31622400),
				},
				Default: int64default.StaticInt64(3600),
			},
			"additional_claims": models.GetDmJWTClaimsResourceSchema("<p>Additional JWT claims, such as audience \"aud\" claim, not before \"nbf\" claim, issued at \"iat\" claim, JWT ID \"jit\" claim, \"nonce\" claim, and custom claim, can be added for JWT.</p><p>The subject, \"sub\" claim is added by default. You can override the subject claim value by specifying the \"sub\" claim in the Custom claims field.</p>", "add-claims", "", false),
			"audience": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The audience, \"aud\", claim identifies the recipients that the JWT is intended for. The maximum length of the Audience claim is 256 characters.", "aud", "").AddRequiredWhen(models.AAAJWTGeneratorAudienceCondVal.String()).String,
				ElementType:         types.StringType,
				Optional:            true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(),
					validators.ConditionalRequiredList(models.AAAJWTGeneratorAudienceCondVal, validators.Evaluation{}, false),
				},
			},
			"not_before": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The not before claim, \"nbf\", indicates the time before which the JWT must not be accepted for processing. Enter a value in the range 0 - 480. The default value is 0.", "nbf", "").AddIntegerRange(0, 480).AddRequiredWhen(models.AAAJWTGeneratorNotBeforeCondVal.String()).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 480),
					validators.ConditionalRequiredInt64(models.AAAJWTGeneratorNotBeforeCondVal, validators.Evaluation{}, false),
				},
			},
			"custom_claims": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The GatewayScript or XSLT file is processed to specify the custom claim. The GatewayScript or XSLT file must be stored in the <tt>local:</tt> or <tt>store:</tt> directory.", "custom-claims", "").AddRequiredWhen(models.AAAJWTGeneratorCustomClaimsCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AAAJWTGeneratorCustomClaimsCondVal, validators.Evaluation{}, false),
				},
			},
			"gen_method": models.GetDmJWTGenMethodResourceSchema("The signing and encryption methods can be used to secure and generate a JWT.", "generate-method", "", false),
			"sign_algorithm": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Various signing algorithms can be used to generate the JWT signature, such as HS256, HS384, HS512, RS256, RS384, and RS512. The default value is RS256.", "sign-alg", "").AddStringEnum("HS256", "HS384", "HS512", "RS256", "RS384", "RS512", "ES256", "ES384", "ES512", "PS256", "PS384", "PS512").AddDefaultValue("RS256").AddRequiredWhen(models.AAAJWTGeneratorSignAlgorithmCondVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("HS256", "HS384", "HS512", "RS256", "RS384", "RS512", "ES256", "ES384", "ES512", "PS256", "PS384", "PS512"),
					validators.ConditionalRequiredString(models.AAAJWTGeneratorSignAlgorithmCondVal, validators.Evaluation{}, true),
				},
				Default: stringdefault.StaticString("RS256"),
			},
			"sign_key": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The key alias can be used to sign the JWT. You can get a key alias by configuring the Crypto Key.", "sign-key", "crypto_key").AddRequiredWhen(models.AAAJWTGeneratorSignKeyCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AAAJWTGeneratorSignKeyCondVal, validators.Evaluation{}, false),
				},
			},
			"sign_sskey": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The shared secret key alias can be used to sign the JWT. You can get the shared secret key alias by configuring the Crypto Shared Secret Key.", "sign-sskey", "crypto_sskey").AddRequiredWhen(models.AAAJWTGeneratorSignSSKeyCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AAAJWTGeneratorSignSSKeyCondVal, validators.Evaluation{}, false),
				},
			},
			"enc_algorithm": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Various encryption algorithms can be used to encrypt the JWT, such as A128CBC-HS256, A192CBC-HS384, A256CBC-HS512, A128GCM, A192GCM, and A256GCM. The default value is A128CBC-HS256.", "enc", "").AddStringEnum("A128CBC-HS256", "A192CBC-HS384", "A256CBC-HS512", "A128GCM", "A192GCM", "A256GCM").AddDefaultValue("A128CBC-HS256").AddRequiredWhen(models.AAAJWTGeneratorEncAlgorithmCondVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("A128CBC-HS256", "A192CBC-HS384", "A256CBC-HS512", "A128GCM", "A192GCM", "A256GCM"),
					validators.ConditionalRequiredString(models.AAAJWTGeneratorEncAlgorithmCondVal, validators.Evaluation{}, true),
				},
				Default: stringdefault.StaticString("A128CBC-HS256"),
			},
			"encrypt_algorithm": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Various key management algorithms can be used to encrypt the JWT, such as RSA1_5, RSA-OAEP, RSA-OAEP-256, A128KW, A192KW, A256KW, and dir. The default value is RSA1_5.", "enc-alg", "").AddStringEnum("RSA1_5", "RSA-OAEP", "RSA-OAEP-256", "A128KW", "A192KW", "A256KW", "dir").AddDefaultValue("RSA1_5").AddRequiredWhen(models.AAAJWTGeneratorEncryptAlgorithmCondVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("RSA1_5", "RSA-OAEP", "RSA-OAEP-256", "A128KW", "A192KW", "A256KW", "dir"),
					validators.ConditionalRequiredString(models.AAAJWTGeneratorEncryptAlgorithmCondVal, validators.Evaluation{}, true),
				},
				Default: stringdefault.StaticString("RSA1_5"),
			},
			"encrypt_certificate": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The certificate alias can be used to encrypt the JWT. You can get the certificate alias by configuring the Crypto Certificate.", "enc-cert", "crypto_certificate").AddRequiredWhen(models.AAAJWTGeneratorEncryptCertificateCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AAAJWTGeneratorEncryptCertificateCondVal, validators.Evaluation{}, false),
				},
			},
			"encrypt_sskey": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The shared secret key alias can be used to encrypt the JWT. You can get the shared secret key alias by configuring the Crypto Shared Secret Key.", "enc-sskey", "crypto_sskey").AddRequiredWhen(models.AAAJWTGeneratorEncryptSSKeyCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AAAJWTGeneratorEncryptSSKeyCondVal, validators.Evaluation{}, false),
				},
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *AAAJWTGeneratorResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *AAAJWTGeneratorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AAAJWTGenerator
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `AAAJWTGenerator`)
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

func (r *AAAJWTGeneratorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AAAJWTGenerator
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
		data.FromBody(ctx, `AAAJWTGenerator`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `AAAJWTGenerator`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AAAJWTGeneratorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AAAJWTGenerator
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `AAAJWTGenerator`))
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

func (r *AAAJWTGeneratorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AAAJWTGenerator
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

func (r *AAAJWTGeneratorResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.AAAJWTGenerator

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
