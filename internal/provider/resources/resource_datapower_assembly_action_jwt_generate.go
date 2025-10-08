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
)

var _ resource.Resource = &AssemblyActionJWTGenerateResource{}
var _ resource.ResourceWithImportState = &AssemblyActionJWTGenerateResource{}

func NewAssemblyActionJWTGenerateResource() resource.Resource {
	return &AssemblyActionJWTGenerateResource{}
}

type AssemblyActionJWTGenerateResource struct {
	pData *tfutils.ProviderData
}

func (r *AssemblyActionJWTGenerateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_assembly_action_jwt_generate"
}

func (r *AssemblyActionJWTGenerateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("The JWT generate assembly action generates JWT claims and specifies the crypto material to generate a JWT during processing.", "assembly-jwt-generate", "").String,
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
			"jwt": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the variable to store the generated JWT. The default value is <tt>generated.jwt</tt> . When the variable is not set, the generated JWT is written to the Authorization Header as a Bearer token.", "jwt", "").AddDefaultValue("generated.jwt").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("generated.jwt"),
			},
			"jwt_id_claims": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to add a JWT ID (jti) claim to the JWT. When enabled, a UUID is generated and set as the value of the JWT ID claim.", "jti-claim", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"issuer_claim": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the variable from which to retrieve the issuer (iss) claim value. The default value is <tt>iss.claim</tt> . The maximum value length is 256 characters.", "iss-claim", "").AddDefaultValue("iss.claim").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 256),
				},
				Default: stringdefault.StaticString("iss.claim"),
			},
			"subject_claim": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the variable from which to retrieve the subject (sub) claim value. The maximum value length is 256 characters.", "sub-claim", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 256),
				},
			},
			"audience_claim": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the variable from which to retrieve the audience (aud) claim value. The maximum value length is 256 characters. This value can be a single string, a comma-separated string of values, or an array of one or more values when you set the variable through GatewayScript processing.", "aud-claim", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 256),
				},
			},
			"validity_period": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the validity period in seconds to calculate the expiration (exp) claim. This value is added to the current date and time to be the value for the \"exp\" claim. The JWT is considered valid until expiry. Enter a value in the range 1 - 31622400. The default value is 3600.", "exp-claim", "").AddIntegerRange(1, 31622400).AddDefaultValue("3600").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 31622400),
				},
				Default: int64default.StaticInt64(3600),
			},
			"private_claim": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the variable from which to retrieve a valid set of JSON claims. These claims are added to any set of claims that are specified previously.", "private-claims", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 256),
				},
			},
			"sign_jwk": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("JWK for JWT signature", "jws-jwk", "").String,
				Optional:            true,
			},
			"crypto_algorithm": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the crypto algorithm to use. Use one of the following values. <ul><li><tt>HS256</tt> - HMAC using SHA-256</li><li><tt>HS384</tt> - HMAC using SHA-384</li><li><tt>HS512</tt> - HMAC using SHA-512</li><li><tt>RS256</tt> - RSASSA-PKCS-v1_5 using SHA-256</li><li><tt>RS384</tt> - RSASSA-PKCS-v1_5 using SHA 384</li><li><tt>RS512</tt> - RSASSA-PKCS-v1_5 using SHA-512</li><li><tt>ES256</tt> - ECDSA using P-256 and SHA-256</li><li><tt>ES384</tt> - ECDSA using P-384 and SHA-384</li><li><tt>ES512</tt> - ECDSA using P-521 and SHA-512</li><li><tt>none</tt> - Do not sign the JWT, which is unsecured and provides no integrity protection but can be used for a nest JWT</li><li>An inline parameter to read at runtime</li></ul>", "jws-alg", "").String,
				Optional:            true,
			},
			"sign_crypto": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the crypto object (a shared secret key or certificate) to use to sign the JWT.", "jws-crypto", "").String,
				Optional:            true,
			},
			"custom_kid_value_jws": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the value of the <tt>kid</tt> claim of the JWT for JWS. The maximum length is 256 characters. This value can be a single string, a comma-separated string of values, or an array of values when you set the variable through GatewayScript processing.", "custom-kid-value-jws", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 256),
				},
			},
			"encrypt_algorithm": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the encryption algorithm to use. Use one of the following values. <ul><li><tt>A128CBC-HS256</tt> - AES_128_CBC_HMAC_SHA_256 authenticated encryption algorithm</li><li><tt>A192CBC-HS384</tt> - AES_192_CBC_HMAC_SHA_384 authenticated encryption algorithm</li><li><tt>A256CBC-HS512</tt> - AES_256_CBC_HMAC_SHA_512 authenticated encryption algorithm</li><li><tt>A128GCM</tt> - AES GCM using 128-bit key</li><li><tt>A192GCM</tt> - AES GCM using 192-bit key</li><li><tt>A256GCM</tt> - AES GCM using 256-bit key</li><li>An inline parameter to read at runtime</li></ul>", "jwe-enc", "").String,
				Optional:            true,
			},
			"encrypt_jwk": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("JWK for JWT encryption", "jwe-jwk", "").String,
				Optional:            true,
			},
			"key_encrypt_algorithm": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the key encryption algorithm to use. Use one of the following values. <ul><li><tt>RSA1_5</tt> - RSAES-PKCS1-V1_5</li><li><tt>RSA-OAEP</tt> - RSAES OAEP using default parameters</li><li><tt>RSA-OAEP-256</tt> - RSAES OAEP using SHA-256 and MGF1 with SHA-256</li><li><tt>A128KW</tt> - AES Key Wrap with default initial value using 128 bit key</li><li><tt>A192KW</tt> - AES Key Wrap with default initial value using 192 bit key</li><li><tt>A256KW</tt> - AES Key Wrap with default initial value using 256 bit key</li><li><tt>dir</tt> - Direct use of a shared symmetric key as the CEK</li><li>An inline parameter to read at runtime</li></ul>", "jwe-alg", "").String,
				Optional:            true,
			},
			"encrypt_crypto": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Crypto object for JWT encryption", "jwe-crypto", "").String,
				Optional:            true,
			},
			"custom_kid_value_jwe": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the value of the <tt>kid</tt> claim of the JWT for JWE. The maximum length is 256 characters. This value can be a single string, a comma-separated string of values, or an array of values when you set the variable through GatewayScript processing.", "custom-kid-value-jwe", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 256),
				},
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"title": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Title", "title", "").String,
				Optional:            true,
			},
			"correlation_path": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the path that correlates the API action to a specific part of the API specification. The correlation path specifies the part of the API definition that correlates with the API action. This path is exposed in the debug data by the API gateway for use by debugging tools. For example, for an API configuration that is retrieved from API Connect and specified in an OpenAPI document with IBM extensions, this path is the JSON path to the assembly policy in the IBM extensions section of the document. The path can be expressed in any form that the debugging tool can correlate to the API definition.", "correlation-path", "").String,
				Optional:            true,
			},
			"action_debug": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to enable the GatewayScript debugger to troubleshoot the following GatewayScript files or script.</p><ul><li>Troubleshoot a GatewayScript file that is called from the GatewayScript assembly action.</li><li>Troubleshoot a GatewayScript file that is called from an XSLT assembly action that uses the <tt>gatewayscript()</tt> extension function.</li><li>Troubleshoot a GatewayScript script that is called through the <tt>value</tt> or <tt>default</tt> property in the JSON file from the map assembly action.</li></ul><p>To debug a file or script, the following conditions must be met.</p><ul><li>The file contains one or more <tt>debugger;</tt> statements at the points in your script where you want to start debugging.</li><li>The GatewayScript debugger is enabled.</li></ul><p>You run the <tt>debug-action</tt> command.</p>", "debug", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *AssemblyActionJWTGenerateResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *AssemblyActionJWTGenerateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AssemblyActionJWTGenerate
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `AssemblyActionJWTGenerate`)
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

func (r *AssemblyActionJWTGenerateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AssemblyActionJWTGenerate
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

	data.UpdateFromBody(ctx, `AssemblyActionJWTGenerate`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AssemblyActionJWTGenerateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AssemblyActionJWTGenerate
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `AssemblyActionJWTGenerate`))
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

func (r *AssemblyActionJWTGenerateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AssemblyActionJWTGenerate
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

func (r *AssemblyActionJWTGenerateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

	var data models.AssemblyActionJWTGenerate
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

	data.FromBody(ctx, `AssemblyActionJWTGenerate`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AssemblyActionJWTGenerateResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.AssemblyActionJWTGenerate

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
