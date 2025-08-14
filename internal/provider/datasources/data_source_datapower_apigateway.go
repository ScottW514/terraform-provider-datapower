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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type APIGatewayList struct {
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
	resp.TypeName = req.ProviderTypeName + "_apigateway"
}

func (d *APIGatewayDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "API gateway",
		Attributes: map[string]schema.Attribute{
			"app_domain": schema.StringAttribute{
				MarkdownDescription: "The name of the application domain the object belongs to",
				Required:            true,
			},
			"result": schema.ListNestedAttribute{
				MarkdownDescription: "List of objects",
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
							MarkdownDescription: "URL refresh policy",
							Computed:            true,
						},
						"cache_memory_size": schema.Int64Attribute{
							MarkdownDescription: "Stylesheet cache size",
							Computed:            true,
						},
						"cache_size": schema.Int64Attribute{
							MarkdownDescription: "Stylesheet cache count",
							Computed:            true,
						},
						"sha1_caching": schema.BoolAttribute{
							MarkdownDescription: "SHA1 caching",
							Computed:            true,
						},
						"static_document_calls": schema.BoolAttribute{
							MarkdownDescription: "Static document calls",
							Computed:            true,
						},
						"virtual_servers": schema.ListAttribute{
							MarkdownDescription: "Load balancer groups",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"doc_cache_max_docs": schema.Int64Attribute{
							MarkdownDescription: "Document cache count",
							Computed:            true,
						},
						"doc_cache_size": schema.Int64Attribute{
							MarkdownDescription: "Document cache size",
							Computed:            true,
						},
						"doc_max_writes": schema.Int64Attribute{
							MarkdownDescription: "Maximum concurrent writes",
							Computed:            true,
						},
						"doc_cache_policy": schema.ListNestedAttribute{
							MarkdownDescription: "Document cache policy",
							NestedObject:        models.DmDocCachePolicyDataSourceSchema,
							Computed:            true,
						},
						"scheduled_rule": schema.ListNestedAttribute{
							MarkdownDescription: "Scheduled processing rule",
							NestedObject:        models.DmScheduledRuleDataSourceSchema,
							Computed:            true,
						},
						"api_collection": schema.ListAttribute{
							MarkdownDescription: "API collection",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"share_rate_limit_count": schema.StringAttribute{
							MarkdownDescription: "Share rate limit count",
							Computed:            true,
						},
						"assembly_burst_limit": schema.ListNestedAttribute{
							MarkdownDescription: "Assembly burst limits",
							NestedObject:        models.DmAPIBurstLimitDataSourceSchema,
							Computed:            true,
						},
						"assembly_rate_limit": schema.ListNestedAttribute{
							MarkdownDescription: "Assembly rate limits",
							NestedObject:        models.DmAPIRateLimitDataSourceSchema,
							Computed:            true,
						},
						"assembly_count_limit": schema.ListNestedAttribute{
							MarkdownDescription: "Assembly count limits",
							NestedObject:        models.DmAPICountLimitDataSourceSchema,
							Computed:            true,
						},
						"ldap_conn_pool": schema.StringAttribute{
							MarkdownDescription: "LDAP connection pool",
							Computed:            true,
						},
						"proxy_policies": schema.ListNestedAttribute{
							MarkdownDescription: "Proxy policy",
							NestedObject:        models.DmAPIProxyPolicyDataSourceSchema,
							Computed:            true,
						},
						"front_timeout": schema.Int64Attribute{
							MarkdownDescription: "Front side timeout",
							Computed:            true,
						},
						"front_persistent_timeout": schema.Int64Attribute{
							MarkdownDescription: "Front persistent timeout",
							Computed:            true,
						},
						"open_telemetry": schema.StringAttribute{
							MarkdownDescription: "OpenTelemetry instance",
							Computed:            true,
						},
						"open_telemetry_resource_attribute": schema.ListNestedAttribute{
							MarkdownDescription: "OpenTelemetry resource attributes",
							NestedObject:        models.DmOpenTelemetryResourceAttributeDataSourceSchema,
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

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.APIGateway{}
	if value := res.Get(`APIGateway`); value.Exists() {
		for _, v := range value.Array() {
			item := models.APIGateway{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
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
