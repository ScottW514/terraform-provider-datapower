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

package datasources

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type APIGatewayList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &APIGatewayDataSource{}
	_ datasource.DataSourceWithConfigure = &APIGatewayDataSource{}
)

func NewAPIGatewayDataSource() datasource.DataSource {
	return &APIGatewayDataSource{}
}

type APIGatewayDataSource struct {
	pData *tfutils.ProviderData
}

func (d *APIGatewayDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_gateway"
}

func (d *APIGatewayDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "An API gateway matches the API to process API requests and to route each request to the matched API.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The name of the object to retrieve.",
				Optional:            true,
			},
			"app_domain": schema.StringAttribute{
				MarkdownDescription: "The name of the application domain the object belongs to.",
				Required:            true,
			},
			"result": schema.ListNestedAttribute{
				MarkdownDescription: "List of objects. If `id` was provided and it exists, it will be the only item in the list.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Name of the object. Must be unique among object types in application domain.",
							Computed:            true,
						},
						"app_domain": schema.StringAttribute{
							MarkdownDescription: "The name of the application domain the object belongs to",
							Computed:            true,
						},
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"gateway_service_name": schema.StringAttribute{
							MarkdownDescription: "Gateway service name",
							Computed:            true,
						},
						"front_protocol": schema.ListAttribute{
							MarkdownDescription: "Protocol handler (reference to HTTP or HTTPS Source Protocol Hander)",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"url_refresh_policy": schema.StringAttribute{
							MarkdownDescription: "Specify the stylesheet refresh policy. Stylesheets cached by this gateway are refreshed in accordance with policy rules.",
							Computed:            true,
						},
						"cache_memory_size": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum size of the stylesheet cache. The default value is 2147483647. A value of 0 disables caching. Stylesheets are purged when either the cache size or the cache count is reached.",
							Computed:            true,
						},
						"cache_size": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum number of stylesheets to cache. Enter a value in the range 5 - 250000. The default value is 256. Stylesheets are purged when either the cache size or the cache count is reached.",
							Computed:            true,
						},
						"sha1caching": schema.BoolAttribute{
							MarkdownDescription: "Specify how to manage SHA1-assisted stylesheet caching. With SHA1 caching enabled, stylesheets are cached by both URL and SHA1 message digest value. With SHA1 caching disabled, stylesheets are cached only by URL.",
							Computed:            true,
						},
						"static_document_calls": schema.BoolAttribute{
							MarkdownDescription: "Specify how to manage static document calls. The latest XSLT specifications require that multiple document calls in the same transformation return the same result. Disable this setting to allow all document calls to operate independently.",
							Computed:            true,
						},
						"virtual_servers": schema.ListAttribute{
							MarkdownDescription: "Load balancer groups",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"doc_cache_max_docs": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum number of documents to cache. Enter a value in the range 1 - 250000. The default value is 5000.",
							Computed:            true,
						},
						"doc_cache_size": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum size of the document cache. Regardless of the specified size, no document that is greater than 1073741824 bytes is cached. This restriction applies even if the cache has available space.",
							Computed:            true,
						},
						"doc_max_writes": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum number of concurrent write requests to create documents or refresh expired documents in the document cache. Enter a value in the range 1 - 32768. The default value is 32768. After the maximum number is reached, requests are forwarded to the target server and the response is not written to the cache.",
							Computed:            true,
						},
						"doc_cache_policy": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the document cache policies to associate a set of URLs with a specific cache policy. A document cache policy allows the administrator to determine how documents are cached. The policy offers time-to-live, priority, and type. The document cache is distinct from the stylesheet cache.",
							NestedObject:        models.GetDmDocCachePolicyDataSourceSchema(),
							Computed:            true,
						},
						"scheduled_rule": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the processing rules to run at defined intervals. Certain applications require the running of a processing rule. For example, the integration with a CA Unicenter Manager is facilitated by a regularly scheduled processing rule that obtains relationship data from the Unicenter Manager.",
							NestedObject:        models.GetDmScheduledRuleDataSourceSchema(),
							Computed:            true,
						},
						"api_collection": schema.ListAttribute{
							MarkdownDescription: "Specify the API collections to serve a group of clients. Each collection packages the plans and subscribers to serve a specific group of clients.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"share_rate_limit_count": schema.StringAttribute{
							MarkdownDescription: "Share rate limit count",
							Computed:            true,
						},
						"assembly_burst_limit": schema.ListNestedAttribute{
							MarkdownDescription: "",
							NestedObject:        models.GetDmAPIBurstLimitDataSourceSchema(),
							Computed:            true,
						},
						"assembly_rate_limit": schema.ListNestedAttribute{
							MarkdownDescription: "Assembly rate limits",
							NestedObject:        models.GetDmAPIRateLimitDataSourceSchema(),
							Computed:            true,
						},
						"assembly_count_limit": schema.ListNestedAttribute{
							MarkdownDescription: "Assembly count limits",
							NestedObject:        models.GetDmAPICountLimitDataSourceSchema(),
							Computed:            true,
						},
						"ldap_conn_pool": schema.StringAttribute{
							MarkdownDescription: "LDAP connection pool",
							Computed:            true,
						},
						"proxy_policies": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the proxy policies to associate a set of URLs with a specific HTTP proxy. When multiple proxy policies are defined, URLs are evaluated against each policy in order.",
							NestedObject:        models.GetDmAPIProxyPolicyDataSourceSchema(),
							Computed:            true,
						},
						"front_timeout": schema.Int64Attribute{
							MarkdownDescription: "Specify the intra-transaction timeout for client connections. This value is the maximum idle time to allow in a transaction for a client connection. This timer monitors idle time in the data transfer process. If the specified idle time is exceeded, the connection is torn down. Enter a value in the range 1 - 86400. The default value is 120.",
							Computed:            true,
						},
						"front_persistent_timeout": schema.Int64Attribute{
							MarkdownDescription: "Specify the inter-transaction timeout for client connections. This value is the maximum idle time to allow between the completion of a transaction and the initiation of a new transaction for a client connection. If the specified idle time is exceeded, the connection is torn down. Enter a value in the range 0 - 86400. The default value is 180. A value of 0 disables persistent connections.",
							Computed:            true,
						},
						"open_telemetry": schema.StringAttribute{
							MarkdownDescription: "OpenTelemetry instance",
							Computed:            true,
						},
						"open_telemetry_resource_attribute": schema.ListNestedAttribute{
							MarkdownDescription: "OpenTelemetry resource attributes",
							NestedObject:        models.GetDmOpenTelemetryResourceAttributeDataSourceSchema(),
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *APIGatewayDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *APIGatewayDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data APIGatewayList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.APIGateway{
		AppDomain: data.AppDomain,
	}

	path := o.GetPath()
	if !data.Id.IsNull() {
		path = path + "/" + data.Id.ValueString()
	}

	res, err := d.pData.Client.Get(path)
	resFound := true
	if err != nil {
		if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		} else {
			resFound = false
		}
	}
	l := []models.APIGateway{}
	if resFound {
		if value := res.Get(`APIGateway`); value.Exists() {
			for _, v := range value.Array() {
				item := models.APIGateway{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.APIGatewayObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.APIGatewayObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
