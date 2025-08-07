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
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &APIGatewayResource{}
var _ resource.ResourceWithValidateConfig = &APIGatewayResource{}

func NewAPIGatewayResource() resource.Resource {
	return &APIGatewayResource{}
}

type APIGatewayResource struct {
	client *client.DatapowerClient
}

func (r *APIGatewayResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apigateway"
}

func (r *APIGatewayResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("API gateway", "apigw", "").AddActions("flush_stylesheet_cache", "flush_document_cache").String,
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
			"gateway_service_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Gateway service name", "gateway-service-name", "").String,
				Optional:            true,
			},
			"front_protocol": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Protocol handler (reference to HTTP or HTTPS Source Protocol Hander)", "front-protocol", "").String,
				ElementType:         types.StringType,
				Required:            true,
			},
			"url_refresh_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("URL refresh policy", "xslrefresh", "urlrefreshpolicy").String,
				Optional:            true,
			},
			"cache_memory_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Stylesheet cache size", "xsl-cache-memorysize", "").AddDefaultValue("2147483647").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(2147483647),
			},
			"cache_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Stylesheet cache count", "xsl-cache-size", "").AddIntegerRange(5, 250000).AddDefaultValue("256").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(5, 250000),
				},
				Default: int64default.StaticInt64(256),
			},
			"sha1_caching": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SHA1 caching", "xsl-checksummed-cache", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"static_document_calls": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Static document calls", "static-document-calls", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"virtual_servers": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Load balancer groups", "loadbalancer-group", "loadbalancergroup").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"doc_cache_max_docs": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Document cache count", "maxdocs", "").AddIntegerRange(1, 250000).AddDefaultValue("5000").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 250000),
				},
				Default: int64default.StaticInt64(5000),
			},
			"doc_cache_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Document cache size", "size", "").String,
				Optional:            true,
			},
			"doc_max_writes": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Maximum concurrent writes", "max-writes", "").AddIntegerRange(1, 32768).AddDefaultValue("32768").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 32768),
				},
				Default: int64default.StaticInt64(32768),
			},
			"doc_cache_policy": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Document cache policy", "policy", "").String,
				NestedObject:        models.DmDocCachePolicyResourceSchema,
				Optional:            true,
			},
			"scheduled_rule": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Scheduled processing rule", "schedule-rule", "").String,
				NestedObject:        models.DmScheduledRuleResourceSchema,
				Optional:            true,
			},
			"api_collection": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("API collection", "collection", "apicollection").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"share_rate_limit_count": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Share rate limit count", "share-count", "").AddStringEnum("yes", "no").AddDefaultValue("yes").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("yes", "no"),
				},
				Default: stringdefault.StaticString("yes"),
			},
			"assembly_burst_limit": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Assembly burst limits", "assembly-burst-limit", "").String,
				NestedObject:        models.DmAPIBurstLimitResourceSchema,
				Optional:            true,
			},
			"assembly_rate_limit": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Assembly rate limits", "assembly-rate-limit", "").String,
				NestedObject:        models.DmAPIRateLimitResourceSchema,
				Optional:            true,
			},
			"assembly_count_limit": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Assembly count limits", "assembly-count-limit", "").String,
				NestedObject:        models.DmAPICountLimitResourceSchema,
				Optional:            true,
			},
			"ldap_conn_pool": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP connection pool", "ldap-pool", "ldapconnectionpool").String,
				Optional:            true,
			},
			"proxy_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Proxy policy", "proxy", "").String,
				NestedObject:        models.DmAPIProxyPolicyResourceSchema,
				Optional:            true,
			},
			"front_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Front side timeout", "front-timeout", "").AddIntegerRange(1, 86400).AddDefaultValue("120").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(120),
			},
			"front_persistent_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Front persistent timeout", "front-persistent-timeout", "").AddIntegerRange(0, 86400).AddDefaultValue("180").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 86400),
				},
				Default: int64default.StaticInt64(180),
			},
			"open_telemetry": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("OpenTelemetry instance", "otel", "opentelemetry").String,
				Optional:            true,
			},
			"open_telemetry_resource_attribute": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("OpenTelemetry resource attributes", "otel-resource-attribute", "").String,
				NestedObject:        models.DmOpenTelemetryResourceAttributeResourceSchema,
				Optional:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *APIGatewayResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *APIGatewayResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.APIGateway

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)

	body := data.ToBody(ctx, `APIGateway`)
	_, err := r.client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "POST", err))
		return
	}
	actions.PostProcess(ctx, &resp.Diagnostics, r.client, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIGatewayResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.APIGateway

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
		data.FromBody(ctx, `APIGateway`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `APIGateway`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIGatewayResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.APIGateway

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	_, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `APIGateway`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, r.client, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIGatewayResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.APIGateway

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false)
	_, err := r.client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}
	actions.PostProcess(ctx, &resp.Diagnostics, r.client, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *APIGatewayResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.APIGateway

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
