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

var _ resource.Resource = &CryptoValCredResource{}
var _ resource.ResourceWithImportState = &CryptoValCredResource{}

func NewCryptoValCredResource() resource.Resource {
	return &CryptoValCredResource{}
}

type CryptoValCredResource struct {
	pData *tfutils.ProviderData
}

func (r *CryptoValCredResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_crypto_val_cred"
}

func (r *CryptoValCredResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("<p>Validation credentials authenticate certificates that are received from TLS peers. Validation credentials can be used to validate certificates that are used in digital signature and encryption operations.</p><p>a TLS client requires validation credentials only when it authenticates the certificate that is presented by the remote TLS server. The TLS standard does not require authentication of the server certificate.</p><p>a TLS server requires validation credentials only when it authenticates remote TLS clients. The TLS standard does not require authentication of the client certificate.</p>", "valcred", "").String,
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
			"certificate": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the list of certificates for the validation credentials. Each certificate in the validation credentials is the certificate that a TLS peer might send or is the certificate of the certification authority (CA) that signed the certificate that is sent by a peer or is the root certificate.", "certificate", "crypto_certificate").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"cert_validation_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Certificate validation mode", "cert-validation-mode", "").AddStringEnum("legacy", "pkix", "exact-match").AddDefaultValue("legacy").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("legacy", "pkix", "exact-match"),
				},
				Default: stringdefault.StaticString("legacy"),
			},
			"use_crl": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to check certificate revocation lists (CRLs) during certificate validation. When enabled, CRLs are checked. Otherwise, CRLs are not checked.", "use-crl", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"require_crl": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to mandate CRLs during certificate validation. When enabled, certificate validation fails if no CRL is available. Otherwise, validation succeeds independent of the availability of a CRL.", "require-crl", "").AddDefaultValue("false").AddNotValidWhen(models.CryptoValCredRequireCRLIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"crl_dp_handling": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the support of certificate extensions for X.509 certificate distribution points. This certificate extension specifies how to obtain CRL information. For more information, see RFC 2527 and RFC 3280.", "crldp", "").AddStringEnum("ignore", "require").AddDefaultValue("ignore").AddNotValidWhen(models.CryptoValCredCRLDPHandlingIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("ignore", "require"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.CryptoValCredCRLDPHandlingIgnoreVal, true),
				},
				Default: stringdefault.StaticString("ignore"),
			},
			"initial_policy_set": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the unique object identifiers for the certificate policy. <p>RFC 3280 refers to the input variable for certificate chain validation as <tt>user-initial-policy-set</tt> . These OIDs specify the allow values of certificate policies. To use this functionality, you need to require an explicit certificate policy. Otherwise, this set is used only if there are policy constraint extensions in the certificate chain.</p><p>By default, the initial certificate policy set consists of the single OID 2.5.29.32.0, which identifies <tt>anyPolicy</tt> .</p>", "initial-policy-set", "").AddNotValidWhen(models.CryptoValCredInitialPolicySetIgnoreVal.String()).String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"explicit_policy": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify support for the initial explicit policy variable as defined by RFC 3280. When enabled, the chain validation algorithm must end with a non-empty policy tree. Otherwise, the algorithm can end with an empty policy tree unless policy constraint extensions in the chain require an explicit policy.", "explicit-policy", "").AddDefaultValue("false").AddNotValidWhen(models.CryptoValCredExplicitPolicyIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"check_dates": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to check dates during certificate validation. This validation checks the current date and time against the <tt>notBefore</tt> and <tt>notAfter</tt> values in certificates and CRLs. When enabled, the date values are checked and expired certificates cause validation to fail. Otherwise, the date values are ignored and and do not cause validation to fail when a certificate is expired.", "check-dates", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *CryptoValCredResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *CryptoValCredResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.CryptoValCred
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `CryptoValCred`)
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

func (r *CryptoValCredResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.CryptoValCred
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

	data.UpdateFromBody(ctx, `CryptoValCred`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CryptoValCredResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.CryptoValCred
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `CryptoValCred`))
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

func (r *CryptoValCredResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.CryptoValCred
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

func (r *CryptoValCredResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

	var data models.CryptoValCred
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

	data.FromBody(ctx, `CryptoValCred`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CryptoValCredResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.CryptoValCred

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
