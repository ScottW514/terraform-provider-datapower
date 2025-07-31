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
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &APICollectionResource{}

func NewAPICollectionResource() resource.Resource {
	return &APICollectionResource{}
}

type APICollectionResource struct {
	client *client.DatapowerClient
}

func (r *APICollectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apicollection"
}

func (r *APICollectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("API collection", "api-collection", "").String,

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
					ImmutableAfterSet(),
				},
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"sandbox": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Sandbox", "sandbox", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
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
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default-catalog-id"),
			},
			"catalog_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Catalog name", "catalog-name", "").AddDefaultValue("default").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default"),
			},
			"dev_portal_endpoint": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Developer Portal endpoint", "dev-portal-endpoint", "").String,
				Optional:            true,
			},
			"cache_capacity": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Subscriber cache capacity", "cache-capacity", "").AddIntegerRange(8, 51200).AddDefaultValue("128").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(8, 51200),
				},
				Default: int64default.StaticInt64(128),
			},
			"routing_prefix": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Routing prefixes", "routing-prefix", "").String,
				NestedObject:        models.DmRoutingPrefixResourceSchema,
				Required:            true,
			},
			"use_rate_limit_group": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use rate limit group", "use-rate-limit-group", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"default_rate_limit": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Default rate limit", "default-rate-limit", "").String,
				NestedObject:        models.DmAPIRateLimitResourceSchema,
				Optional:            true,
			},
			"rate_limit_group": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Rate limit group", "rate-limit-group", "ratelimitdefinitiongroup").String,
				Optional:            true,
			},
			"assembly_burst_limit": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Assembly burst limit", "assembly-burst-limit", "").String,
				NestedObject:        models.DmAPIBurstLimitResourceSchema,
				Optional:            true,
			},
			"assembly_rate_limit": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Assembly rate limit", "assembly-rate-limit", "").String,
				NestedObject:        models.DmAPIRateLimitResourceSchema,
				Optional:            true,
			},
			"assembly_count_limit": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Assembly count limit", "assembly-count-limit", "").String,
				NestedObject:        models.DmAPICountLimitResourceSchema,
				Optional:            true,
			},
			"enforce_pre_assembly_rate_limits": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enforce preassembly rate limits", "enforce-pre-assembly-rate-limits", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"api_processing_rule": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("API processing rule", "processing-rule", "apirule").AddDefaultValue("default-api-rule").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default-api-rule"),
			},
			"api_error_rule": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("API error rule", "error-rule", "apirule").AddDefaultValue("default-api-error-rule").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default-api-error-rule"),
			},
			"assembly_preflow": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Assembly preprocessing", "assembly-preprocessing", "assembly").String,
				Optional:            true,
			},
			"assembly_postflow": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Assembly postprocessing", "assembly-postprocessing", "assembly").String,
				Optional:            true,
			},
			"plan": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Plans", "plan", "apiplan").String,
				ElementType:         types.StringType,
				Required:            true,
			},
			"analytics_endpoint": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Analytic endpoint", "analytics-endpoint", "analyticsendpoint").String,
				Optional:            true,
			},
			"application_type": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Application types", "application-type", "apiapplicationtype").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"parse_settings_reference": models.GetDmDynamicParseSettingsReferenceResourceSchema("Parse settings", "parse-settings-reference", "", false),
		},
	}
}

func (r *APICollectionResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *APICollectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.APICollection

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `APICollection`)
	_, err := r.client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "POST", err))
		return
	}

	_, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s", "POST", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APICollectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.APICollection

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && (strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}

	if data.IsNull() {
		// Import
		data.FromBody(ctx, `APICollection`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `APICollection`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APICollectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.APICollection

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `APICollection`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}
	_, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s", "POST", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APICollectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.APICollection

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}

	resp.State.RemoveResource(ctx)
}
