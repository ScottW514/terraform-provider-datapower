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
)

var _ resource.Resource = &APIPlanResource{}

func NewAPIPlanResource() resource.Resource {
	return &APIPlanResource{}
}

type APIPlanResource struct {
	pData *tfutils.ProviderData
}

func (r *APIPlanResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_plan"
}

func (r *APIPlanResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("An API plan packages a list of APIs to expose. An API is not exposed unless you add the API to a plan. When you configure an API plan, define the rate limit schemes to enforce against APIs. By default, the rate limit scheme in the plan applies to all operations. You can override plan-level rate limit schemes with operation-specific rate limit schemes.", "api-plan", "").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Plan name", "name", "").String,
				Optional:            true,
			},
			"product_id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the product ID for the plan. A product makes a set of APIs and plans into one offering to make available to API developers.", "product-id", "").String,
				Optional:            true,
			},
			"product_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Product name", "product-name", "").String,
				Optional:            true,
			},
			"product_version": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Product version", "product-version", "").AddDefaultValue("1.0.0").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("1.0.0"),
			},
			"product_title": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Product title", "product-title", "").String,
				Optional:            true,
			},
			"use_rate_limit_group": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use rate limit group", "use-rate-limit-group", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"rate_limit": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the rate limit scheme to enforce. This scheme defines the maximum rate to allow during a specified interval and the actions to take when the limit is exceeded.", "rate-limit", "").String,
				NestedObject:        models.DmAPIRateLimitResourceSchema,
				Optional:            true,
			},
			"burst_limit": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the burst limit scheme to enforce. This scheme defines the maximum burst rate to allow during a specified interval. The burst limit helps to prevent spikes that might damage the infrastructure. When a message arrives within an interval, the burst limit takes priority over the rate limit. In other words, a message is first checked against the burst limit scheme and then against the rate limit scheme.", "burst-limit", "").String,
				NestedObject:        models.DmAPIBurstLimitResourceSchema,
				Optional:            true,
			},
			"rate_limit_group": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Rate limit group", "rate-limit-group", "rate_limit_definition_group").String,
				Optional:            true,
			},
			"use_limit_definitions": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use limit definitions", "use-definitions", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"assembly_burst_limit": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the burst limit scheme that the rate limit assembly action enforces. This scheme defines the maximum burst rate to allow during a specified interval. This scheme helps to prevent spikes that might damage the infrastructure. When a message arrives within an interval, the burst limit takes priority over the rate limit. In other words, a message is first checked against the burst limit scheme and then against the rate limit scheme.", "assembly-burst-limit", "").String,
				NestedObject:        models.DmAPIBurstLimitResourceSchema,
				Optional:            true,
			},
			"assembly_burst_limit_definition": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify a burst limit definition that the rate limit assembly action enforces. A burst limit definition defines the maximum burst rate to allow during a specified interval. This scheme helps to prevent spikes that might damage infrastructure. When a message arrives within an interval, the burst limit takes priority over the rate limit. A message is first checked against the burst limit scheme and then against the rate limit scheme.", "assembly-burst-limit-def", "").String,
				NestedObject:        models.DmDefinitionLinkResourceSchema,
				Optional:            true,
			},
			"assembly_rate_limit": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the rate limit scheme that the rate limit assembly action enforces. This scheme defines the maximum rate to allow during a specified interval and the actions to take when the limit is exceeded.", "assembly-rate-limit", "").String,
				NestedObject:        models.DmAPIRateLimitResourceSchema,
				Optional:            true,
			},
			"assembly_rate_limit_definition": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify a rate limit definition that the rate limit assembly action enforces. A rate limit definition defines the maximum rate that is allowed in a specified interval and the actions to take when the limit is exceeded.", "assembly-rate-limit-def", "").String,
				NestedObject:        models.DmDefinitionLinkResourceSchema,
				Optional:            true,
			},
			"assembly_count_limit": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the count limit scheme that the rate limit assembly action enforces. This scheme defines the maximum count to allow and the actions to take when the limit is exceeded.", "assembly-count-limit", "").String,
				NestedObject:        models.DmAPICountLimitResourceSchema,
				Optional:            true,
			},
			"assembly_count_limit_definition": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify a count limit definition that the rate limit assembly action enforces. A count limit definition defines the maximum count that is allowed and the actions to take when the limit is exceeded.", "assembly-count-limit-def", "").String,
				NestedObject:        models.DmDefinitionLinkResourceSchema,
				Optional:            true,
			},
			"space_id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the space ID for the product in the catalog. When space is enabled for a catalog, the catalog can be partitioned to spaces. Spaces enable each team to manage their APIs independently.", "space-id", "").String,
				Optional:            true,
			},
			"space_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the space name for the product in the catalog. When space is enabled for a catalog, the catalog can be partitioned to spaces. Spaces enable each team to manage their APIs independently.", "space-name", "").String,
				Optional:            true,
			},
			"api": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the APIs to package for the plan. An API is exposed through a plan by associating the API to the plan.", "api", "api_definition").String,
				ElementType:         types.StringType,
				Required:            true,
			},
			"exclude_operation": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Exclude operation", "exclude", "api_operation").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"override": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Operation rate limit", "override", "operation_rate_limit").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"rate_limit_scope": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the scope to apply the rate limit schemes to. You can apply schemes against the application or client ID. For example, <tt>application1</tt> has <tt>client1</tt> and <tt>client2</tt> , and the rate limit is 10 calls per hour. <ul><li>When against the application, <tt>application1</tt> limits 10 calls per hour from either <tt>client1</tt> or <tt>client2.</tt></li><li>When against the client ID, <tt>application1</tt> limits 10 calls per hour from each <tt>client1</tt> and <tt>client2</tt> .</li></ul>", "rate-limit-scope", "").AddStringEnum("per-application", "per-client-id").AddDefaultValue("per-application").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("per-application", "per-client-id"),
				},
				Default: stringdefault.StaticString("per-application"),
			},
			"graph_ql_schema_options": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("GraphQL schema options", "graphql-schema-options", "graphql_schema_options").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *APIPlanResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *APIPlanResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APIPlan
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `APIPlan`)
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

func (r *APIPlanResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APIPlan
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
		data.FromBody(ctx, `APIPlan`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `APIPlan`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIPlanResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APIPlan
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `APIPlan`))
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

func (r *APIPlanResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APIPlan
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

func (r *APIPlanResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.APIPlan

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
