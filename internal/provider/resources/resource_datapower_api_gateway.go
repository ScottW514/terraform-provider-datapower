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

var _ resource.Resource = &APIGatewayResource{}
var _ resource.ResourceWithValidateConfig = &APIGatewayResource{}

func NewAPIGatewayResource() resource.Resource {
	return &APIGatewayResource{}
}

type APIGatewayResource struct {
	pData *tfutils.ProviderData
}

func (r *APIGatewayResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_gateway"
}

func (r *APIGatewayResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("An API gateway matches the API to process API requests and to route each request to the matched API.", "apigw", "").AddActions("flush_stylesheet_cache", "flush_document_cache").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the stylesheet refresh policy. Stylesheets cached by this gateway are refreshed in accordance with policy rules.", "xslrefresh", "url_refresh_policy").String,
				Optional:            true,
			},
			"cache_memory_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum size of the stylesheet cache. The default value is 2147483647. A value of 0 disables caching. Stylesheets are purged when either the cache size or the cache count is reached.", "xsl-cache-memorysize", "").AddDefaultValue("2147483647").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(2147483647),
			},
			"cache_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum number of stylesheets to cache. Enter a value in the range 5 - 250000. The default value is 256. Stylesheets are purged when either the cache size or the cache count is reached.", "xsl-cache-size", "").AddIntegerRange(5, 250000).AddDefaultValue("256").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(5, 250000),
				},
				Default: int64default.StaticInt64(256),
			},
			"sha1caching": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how to manage SHA1-assisted stylesheet caching. With SHA1 caching enabled, stylesheets are cached by both URL and SHA1 message digest value. With SHA1 caching disabled, stylesheets are cached only by URL.", "xsl-checksummed-cache", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"static_document_calls": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how to manage static document calls. The latest XSLT specifications require that multiple document calls in the same transformation return the same result. Disable this setting to allow all document calls to operate independently.", "static-document-calls", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"virtual_servers": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Load balancer groups", "loadbalancer-group", "load_balancer_group").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"doc_cache_max_docs": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum number of documents to cache. Enter a value in the range 1 - 250000. The default value is 5000.", "maxdocs", "").AddIntegerRange(1, 250000).AddDefaultValue("5000").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 250000),
				},
				Default: int64default.StaticInt64(5000),
			},
			"doc_cache_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum size of the document cache. Regardless of the specified size, no document that is greater than 1073741824 bytes is cached. This restriction applies even if the cache has available space.", "size", "").String,
				Optional:            true,
			},
			"doc_max_writes": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum number of concurrent write requests to create documents or refresh expired documents in the document cache. Enter a value in the range 1 - 32768. The default value is 32768. After the maximum number is reached, requests are forwarded to the target server and the response is not written to the cache.", "max-writes", "").AddIntegerRange(1, 32768).AddDefaultValue("32768").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 32768),
				},
				Default: int64default.StaticInt64(32768),
			},
			"doc_cache_policy": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the document cache policies to associate a set of URLs with a specific cache policy. A document cache policy allows the administrator to determine how documents are cached. The policy offers time-to-live, priority, and type. The document cache is distinct from the stylesheet cache.", "policy", "").String,
				NestedObject:        models.GetDmDocCachePolicyResourceSchema(),
				Optional:            true,
			},
			"scheduled_rule": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the processing rules to run at defined intervals. Certain applications require the running of a processing rule. For example, the integration with a CA Unicenter Manager is facilitated by a regularly scheduled processing rule that obtains relationship data from the Unicenter Manager.", "schedule-rule", "").String,
				NestedObject:        models.GetDmScheduledRuleResourceSchema(),
				Optional:            true,
			},
			"api_collection": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the API collections to serve a group of clients. Each collection packages the plans and subscribers to serve a specific group of clients.", "collection", "api_collection").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("", "assembly-burst-limit", "").String,
				NestedObject:        models.GetDmAPIBurstLimitResourceSchema(),
				Optional:            true,
			},
			"assembly_rate_limit": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Assembly rate limits", "assembly-rate-limit", "").String,
				NestedObject:        models.GetDmAPIRateLimitResourceSchema(),
				Optional:            true,
			},
			"assembly_count_limit": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Assembly count limits", "assembly-count-limit", "").String,
				NestedObject:        models.GetDmAPICountLimitResourceSchema(),
				Optional:            true,
			},
			"ldap_conn_pool": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP connection pool", "ldap-pool", "ldap_connection_pool").String,
				Optional:            true,
			},
			"proxy_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the proxy policies to associate a set of URLs with a specific HTTP proxy. When multiple proxy policies are defined, URLs are evaluated against each policy in order.", "proxy", "").String,
				NestedObject:        models.GetDmAPIProxyPolicyResourceSchema(),
				Optional:            true,
			},
			"front_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the intra-transaction timeout for client connections. This value is the maximum idle time to allow in a transaction for a client connection. This timer monitors idle time in the data transfer process. If the specified idle time is exceeded, the connection is torn down. Enter a value in the range 1 - 86400. The default value is 120.", "front-timeout", "").AddIntegerRange(1, 86400).AddDefaultValue("120").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(120),
			},
			"front_persistent_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the inter-transaction timeout for client connections. This value is the maximum idle time to allow between the completion of a transaction and the initiation of a new transaction for a client connection. If the specified idle time is exceeded, the connection is torn down. Enter a value in the range 0 - 86400. The default value is 180. A value of 0 disables persistent connections.", "front-persistent-timeout", "").AddIntegerRange(0, 86400).AddDefaultValue("180").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 86400),
				},
				Default: int64default.StaticInt64(180),
			},
			"open_telemetry": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("OpenTelemetry instance", "otel", "open_telemetry").String,
				Optional:            true,
			},
			"open_telemetry_resource_attribute": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("OpenTelemetry resource attributes", "otel-resource-attribute", "").String,
				NestedObject:        models.GetDmOpenTelemetryResourceAttributeResourceSchema(),
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

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *APIGatewayResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APIGateway
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `APIGateway`)
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

func (r *APIGatewayResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APIGateway
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
		data.FromBody(ctx, `APIGateway`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `APIGateway`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIGatewayResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APIGateway
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `APIGateway`))
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

func (r *APIGatewayResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APIGateway
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

func (r *APIGatewayResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.APIGateway

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
