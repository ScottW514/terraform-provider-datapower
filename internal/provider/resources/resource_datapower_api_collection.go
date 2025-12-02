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

var (
	_ resource.Resource                   = &APICollectionResource{}
	_ resource.ResourceWithValidateConfig = &APICollectionResource{}
	_ resource.ResourceWithImportState    = &APICollectionResource{}
)

func NewAPICollectionResource() resource.Resource {
	return &APICollectionResource{}
}

type APICollectionResource struct {
	pData *tfutils.ProviderData
}

func (r *APICollectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_collection"
}

func (r *APICollectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("An API collection is a logical partition of an API gateway that packages the plans and subscribers to make APIs available to a specific group of clients. An API collection corresponds to a catalog in the API manager.", "api-collection", "").AddActions("flush_cache").String,
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
			"sandbox": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the APIs in this catalog are for test purposes. By default, a catalog is not for test purposes.", "sandbox", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"org_id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Organization ID", "organization-id", "").String,
				Required:            true,
			},
			"org_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Organization name", "organization-name", "").String,
				Required:            true,
			},
			"catalog_id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Catalog ID", "catalog-id", "").AddDefaultValue("default-catalog-id").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("default-catalog-id"),
			},
			"catalog_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Catalog name", "catalog-name", "").AddDefaultValue("default").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("default"),
			},
			"dev_portal_endpoint": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL of the Developer Portal endpoint. This endpoint can be used to provide security credentials for access to an API.", "dev-portal-endpoint", "").String,
				Optional:            true,
			},
			"cache_capacity": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum number of subscriber entries to cache. Enter a value in the range 8 - 51200. The default value is 128. When the limit is exceeded, the least recently used (LRU) entry is removed.", "cache-capacity", "").AddIntegerRange(8, 51200).AddDefaultValue("128").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(8, 51200),
				},
				Default: int64default.StaticInt64(128),
			},
			"routing_prefix": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the routing prefix to determine which API collection to route the request. You can use routing prefixes to organize your APIs and plans into collections and subcollections. For example, if you have a collection of APIs serving for a certain purpose, and the APIs are to be used by two segments of your organization, you might create two API collections with the organization name, purpose name, and segment name in the routing prefix. If the organization name is <tt>myorg</tt> , the APIs serve for purpose <tt>purpose1</tt> , and the two segments under the organization is <tt>section1</tt> and <tt>section2</tt> , the resulting URL routing prefixes are <tt>/myorg/purpose1/section1</tt> and <tt>/myorg/purpose1/section2</tt> . The resulting hostname routing prefixes are <tt>section1.purpose1.myorg</tt> and <tt>section2.purpose1.myorg</tt> . <p>The API gateway uses the routing prefix to form the complete URI <tt>routing_prefix/base_path/operation_path</tt> and accepts only the incoming requests with this URI. In the complete URI, <tt>base_path</tt> is the base path on which the API is served, and <tt>operation_path</tt> is the relative path to the base path where the operations are available.</p><p>The default routing prefix is slash (/) when the type is URI and blank when the type is hostname. An API collection becomes the default API collection in the API Gateway when the API collection has a default routing prefix. The API gateway routes a request to the default API collection when other API collections do not match. An API gateway can have only one default API collection. Therefore, regardless of the prefix type, only one API collection can be configured with the default routing prefix.</p>", "routing-prefix", "").String,
				NestedObject:        models.GetDmRoutingPrefixResourceSchema(),
				Required:            true,
			},
			"use_rate_limit_group": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use rate limit group", "use-rate-limit-group", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"default_rate_limit": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the default rate limit scheme for API requests without API keys for client identification. When not defined, requests without API keys are rejected.", "default-rate-limit", "").AddNotValidWhen(models.APICollectionDefaultRateLimitIgnoreVal.String()).String,
				NestedObject:        models.GetDmAPIRateLimitResourceSchema(),
				Optional:            true,
			},
			"rate_limit_group": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the default rate limit group for API requests without API keys for client identification. When not defined, requests without API keys are rejected.", "rate-limit-group", "rate_limit_definition_group").AddNotValidWhen(models.APICollectionRateLimitGroupIgnoreVal.String()).String,
				Optional:            true,
			},
			"assembly_burst_limit": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Assembly burst limit", "assembly-burst-limit", "").String,
				NestedObject:        models.GetDmAPIBurstLimitResourceSchema(),
				Optional:            true,
			},
			"assembly_rate_limit": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Assembly rate limit", "assembly-rate-limit", "").String,
				NestedObject:        models.GetDmAPIRateLimitResourceSchema(),
				Optional:            true,
			},
			"assembly_count_limit": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Assembly count limit", "assembly-count-limit", "").String,
				NestedObject:        models.GetDmAPICountLimitResourceSchema(),
				Optional:            true,
			},
			"enforce_pre_assembly_rate_limits": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enforce the API rate and burst limits from the plan. When disabled, only the limits specified in a rate limit assembly action are applied to this API.", "enforce-pre-assembly-rate-limits", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"api_processing_rule": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the processing rule to process API requests. When your collection requires custom processing, use API Connect global policies to define the custom rules.", "processing-rule", "api_rule").AddDefaultValue("default-api-rule").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("default-api-rule"),
			},
			"api_error_rule": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the processing rule to handle errors during API processing. When your collection requires custom processing, use API Connect global policies to define the custom rules.", "error-rule", "api_rule").AddDefaultValue("default-api-error-rule").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("default-api-error-rule"),
			},
			"assembly_preflow": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the processing rule to run before the assembly rule. When your collection requires custom processing, use API Connect global policies to configure the assembly.", "assembly-preprocessing", "assembly").String,
				Optional:            true,
			},
			"assembly_postflow": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the processing rule to run after the assembly rule. When your collection requires custom processing, use API Connect global policies to configure the assembly.", "assembly-postprocessing", "assembly").String,
				Optional:            true,
			},
			"plan": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the API plans for the collection. Each plan contains a list of APIs and defines the rate limit for the API operations.", "plan", "api_plan").String,
				ElementType:         types.StringType,
				Required:            true,
			},
			"analytics_endpoint": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Analytic endpoint", "analytics-endpoint", "analytics_endpoint").String,
				Optional:            true,
			},
			"application_type": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Application types", "application-type", "api_application_type").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"parse_settings_reference": models.GetDmDynamicParseSettingsReferenceResourceSchema("Specify the constraints to parse documents. Precedence rules apply when the constraint for the same aspect of an input document is configured with more than one method. <ul><li>You can specify a URL reference from which to retrieve the constraints definition.</li><li>You can specify a literal configuration string in XML management interface or REST management interface format that contains the constraints definition.</li><li>You can specify a parse settings configuration object to retrieve the constraints definition.</li></ul>", "parse-settings-reference", "", false),
			"dependency_actions":       actions.ActionsSchema,
		},
	}
}

func (r *APICollectionResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *APICollectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APICollection
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

	body := data.ToBody(ctx, `APICollection`)
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

func (r *APICollectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APICollection
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

	data.UpdateFromBody(ctx, `APICollection`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APICollectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APICollection
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
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `APICollection`), data.ProviderTarget)
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

func (r *APICollectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APICollection
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

func (r *APICollectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

	var data models.APICollection
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

	data.FromBody(ctx, `APICollection`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
func (r *APICollectionResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.APICollection

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
