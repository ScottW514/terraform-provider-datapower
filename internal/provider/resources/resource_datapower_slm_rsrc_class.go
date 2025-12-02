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

var (
	_ resource.Resource                = &SLMRsrcClassResource{}
	_ resource.ResourceWithImportState = &SLMRsrcClassResource{}
)

func NewSLMRsrcClassResource() resource.Resource {
	return &SLMRsrcClassResource{}
}

type SLMRsrcClassResource struct {
	pData *tfutils.ProviderData
}

func (r *SLMRsrcClassResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_slm_rsrc_class"
}

func (r *SLMRsrcClassResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("<p>An SLM resource class identifies a set of resources to be subject to an SLM policy. A resource class consists of the following types and properties.</p><ul><li>A resource type that specifies the manner to identify resources.</li><li>A match type that determines the resources to apply the policy.</li><li>Depending on the resource and match type, properties that identify specific instances of resources.</li></ul>", "slm-rsrc", "").String,
		Attributes: map[string]schema.Attribute{
			"provider_target": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Target host for this resource. If not set, provider will use the top level settings.", "", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
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
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"rsrc_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Resource type", "type", "").AddStringEnum("aaa-mapped-resource", "front-url", "destination-url", "xpath-filter", "request-message", "response-message", "soap-fault", "errorcode", "custom-stylesheet", "concurrent-connections", "concurrent-transactions", "wsdl", "wsdl-service", "wsdl-port", "wsdl-operation", "request-mq-qname", "reply-mq-qname", "uddi-subscription", "wsrr-subscription", "wsrr-saved-search-subscription").AddDefaultValue("aaa-mapped-resource").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("aaa-mapped-resource", "front-url", "destination-url", "xpath-filter", "request-message", "response-message", "soap-fault", "errorcode", "custom-stylesheet", "concurrent-connections", "concurrent-transactions", "wsdl", "wsdl-service", "wsdl-port", "wsdl-operation", "request-mq-qname", "reply-mq-qname", "uddi-subscription", "wsrr-subscription", "wsrr-saved-search-subscription"),
				},
				Default: stringdefault.StaticString("aaa-mapped-resource"),
			},
			"rsrc_match_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Match type", "match-type", "").AddStringEnum("per-extracted-value", "exact-match", "regexp-match").AddDefaultValue("per-extracted-value").AddRequiredWhen(models.SLMRsrcClassRsrcMatchTypeCondVal.String()).AddNotValidWhen(models.SLMRsrcClassRsrcMatchTypeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("per-extracted-value", "exact-match", "regexp-match"),
					validators.ConditionalRequiredString(models.SLMRsrcClassRsrcMatchTypeCondVal, models.SLMRsrcClassRsrcMatchTypeIgnoreVal, true),
				},
				Default: stringdefault.StaticString("per-extracted-value"),
			},
			"rsrc_value": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Resource value", "value", "").AddNotValidWhen(models.SLMRsrcClassRsrcValueIgnoreVal.String()).String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"stylesheet": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Custom stylesheet", "stylesheet", "").AddRequiredWhen(models.SLMRsrcClassStylesheetCondVal.String()).AddNotValidWhen(models.SLMRsrcClassStylesheetIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.SLMRsrcClassStylesheetCondVal, models.SLMRsrcClassStylesheetIgnoreVal, false),
				},
			},
			"xpath_filter": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XPath filter", "xpath-filter", "").AddRequiredWhen(models.SLMRsrcClassXPathFilterCondVal.String()).AddNotValidWhen(models.SLMRsrcClassXPathFilterIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.SLMRsrcClassXPathFilterCondVal, models.SLMRsrcClassXPathFilterIgnoreVal, false),
				},
			},
			"subscription": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("UDDI subscription (deprecated)", "subscription", "").AddRequiredWhen(models.SLMRsrcClassSubscriptionCondVal.String()).AddNotValidWhen(models.SLMRsrcClassSubscriptionIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.SLMRsrcClassSubscriptionCondVal, models.SLMRsrcClassSubscriptionIgnoreVal, false),
				},
			},
			"wsrr_subscription": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WSRR subscription", "wsrr-subscription", "wsrr_subscription").AddRequiredWhen(models.SLMRsrcClassWSRRSubscriptionCondVal.String()).AddNotValidWhen(models.SLMRsrcClassWSRRSubscriptionIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.SLMRsrcClassWSRRSubscriptionCondVal, models.SLMRsrcClassWSRRSubscriptionIgnoreVal, false),
				},
			},
			"wsrr_saved_search_subscription": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WSRR saved search subscription", "wsrr-saved-search-subscription", "wsrr_saved_search_subscription").AddRequiredWhen(models.SLMRsrcClassWSRRSavedSearchSubscriptionCondVal.String()).AddNotValidWhen(models.SLMRsrcClassWSRRSavedSearchSubscriptionIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.SLMRsrcClassWSRRSavedSearchSubscriptionCondVal, models.SLMRsrcClassWSRRSavedSearchSubscriptionIgnoreVal, false),
				},
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *SLMRsrcClassResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *SLMRsrcClassResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SLMRsrcClass
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !r.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), r.pData.Client.GetTargetNames()))
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `SLMRsrcClass`)
	_, err := r.pData.Client.Post(data.GetPath(), body, data.ProviderTarget)
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), body, data.ProviderTarget)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Resource already exists. Failed to update resource, got error: %s", err))
				return
			}
			resp.Diagnostics.AddWarning("Warning", "Resource already exists. Existing resource was updated.")
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create resource, got error: %s", err))
			return
		}
	}
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}

	tfutils.SaveDomain(ctx, &resp.Diagnostics, r.pData.Client, data.AppDomain, data.ProviderTarget)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SLMRsrcClassResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SLMRsrcClass
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !r.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), r.pData.Client.GetTargetNames()))
		return
	}
	res, err := r.pData.Client.Get(data.GetPath()+"/"+data.Id.ValueString(), data.ProviderTarget)
	if err != nil {
		if strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400") {
			resp.State.RemoveResource(ctx)
			return
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.State.RemoveResource(ctx)
			}
			return
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
			return
		}
	}

	data.UpdateFromBody(ctx, `SLMRsrcClass`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SLMRsrcClassResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SLMRsrcClass
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !r.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), r.pData.Client.GetTargetNames()))
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `SLMRsrcClass`), data.ProviderTarget)
	if err != nil {
		if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
			return
		}
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Update, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}

	tfutils.SaveDomain(ctx, &resp.Diagnostics, r.pData.Client, data.AppDomain, data.ProviderTarget)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SLMRsrcClassResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SLMRsrcClass
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !r.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), r.pData.Client.GetTargetNames()))
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Delete(data.GetPath()+"/"+data.Id.ValueString(), data.ProviderTarget)
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			resp.Diagnostics.AddWarning("Resource Conflict", fmt.Sprintf("Resource is no longer tracked by Terraform, but may need to be manually deleted on DataPower host. Got error: %s", err))
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete resource, got error: %s", err))
			return
		}
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}

	tfutils.SaveDomain(ctx, &resp.Diagnostics, r.pData.Client, data.AppDomain, data.ProviderTarget)

	resp.State.RemoveResource(ctx)
}

func (r *SLMRsrcClassResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

	var data models.SLMRsrcClass
	data.AppDomain = types.StringValue(appDomain)
	data.Id = types.StringValue(id)
	res, err := r.pData.Client.Get(data.GetPath()+"/"+data.Id.ValueString(), data.ProviderTarget)
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("Resource was not found, got error: %s", err))
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		}
		return
	}

	data.FromBody(ctx, `SLMRsrcClass`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
