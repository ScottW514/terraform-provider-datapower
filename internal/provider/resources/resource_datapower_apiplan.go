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
	"github.com/scottw514/terraform-provider-datapower/client"
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
	client *client.DatapowerClient
}

func (r *APIPlanResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apiplan"
}

func (r *APIPlanResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("API plan", "api-plan", "").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Product ID", "product-id", "").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Rate limit", "rate-limit", "").String,
				NestedObject:        models.DmAPIRateLimitResourceSchema,
				Optional:            true,
			},
			"burst_limit": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Burst limit", "burst-limit", "").String,
				NestedObject:        models.DmAPIBurstLimitResourceSchema,
				Optional:            true,
			},
			"rate_limit_group": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Rate limit group", "rate-limit-group", "ratelimitdefinitiongroup").String,
				Optional:            true,
			},
			"use_limit_definitions": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use limit definitions", "use-definitions", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"assembly_burst_limit": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Assembly burst limit", "assembly-burst-limit", "").String,
				NestedObject:        models.DmAPIBurstLimitResourceSchema,
				Optional:            true,
			},
			"assembly_burst_limit_definition": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Assembly burst limit definition", "assembly-burst-limit-def", "").String,
				NestedObject:        models.DmDefinitionLinkResourceSchema,
				Optional:            true,
			},
			"assembly_rate_limit": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Assembly rate limit", "assembly-rate-limit", "").String,
				NestedObject:        models.DmAPIRateLimitResourceSchema,
				Optional:            true,
			},
			"assembly_rate_limit_definition": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Assembly rate limit definition", "assembly-rate-limit-def", "").String,
				NestedObject:        models.DmDefinitionLinkResourceSchema,
				Optional:            true,
			},
			"assembly_count_limit": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Assembly count limit", "assembly-count-limit", "").String,
				NestedObject:        models.DmAPICountLimitResourceSchema,
				Optional:            true,
			},
			"assembly_count_limit_definition": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Assembly count limit definition", "assembly-count-limit-def", "").String,
				NestedObject:        models.DmDefinitionLinkResourceSchema,
				Optional:            true,
			},
			"space_id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Space ID", "space-id", "").String,
				Optional:            true,
			},
			"space_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Space name", "space-name", "").String,
				Optional:            true,
			},
			"api": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("API", "api", "apidefinition").String,
				ElementType:         types.StringType,
				Required:            true,
			},
			"exclude_operation": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Exclude operation", "exclude", "apioperation").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"override": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Operation rate limit", "override", "operationratelimit").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"rate_limit_scope": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Rate limit scope", "rate-limit-scope", "").AddStringEnum("per-application", "per-client-id").AddDefaultValue("per-application").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("per-application", "per-client-id"),
				},
				Default: stringdefault.StaticString("per-application"),
			},
			"graph_ql_schema_options": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("GraphQL schema options", "graphql-schema-options", "graphqlschemaoptions").String,
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

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *APIPlanResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.APIPlan

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, data.AppDomain.ValueString(), data.DependencyActions, actions.Create)

	body := data.ToBody(ctx, `APIPlan`)
	_, err := r.client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "POST", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIPlanResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.APIPlan

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
		data.FromBody(ctx, `APIPlan`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `APIPlan`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIPlanResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.APIPlan

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, data.AppDomain.ValueString(), data.DependencyActions, actions.Update)
	_, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `APIPlan`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIPlanResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.APIPlan

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete)
	_, err := r.client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}

	resp.State.RemoveResource(ctx)
}
