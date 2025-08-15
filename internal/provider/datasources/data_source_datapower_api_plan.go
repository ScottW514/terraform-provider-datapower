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

type APIPlanList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &APIPlanDataSource{}
	_ datasource.DataSourceWithConfigure = &APIPlanDataSource{}
)

func NewAPIPlanDataSource() datasource.DataSource {
	return &APIPlanDataSource{}
}

type APIPlanDataSource struct {
	pData *tfutils.ProviderData
}

func (d *APIPlanDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_plan"
}

func (d *APIPlanDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "An API plan packages a list of APIs to expose. An API is not exposed unless you add the API to a plan. When you configure an API plan, define the rate limit schemes to enforce against APIs. By default, the rate limit scheme in the plan applies to all operations. You can override plan-level rate limit schemes with operation-specific rate limit schemes.",
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
						"name": schema.StringAttribute{
							MarkdownDescription: "Plan name",
							Computed:            true,
						},
						"product_id": schema.StringAttribute{
							MarkdownDescription: "Specify the product ID for the plan. A product makes a set of APIs and plans into one offering to make available to API developers.",
							Computed:            true,
						},
						"product_name": schema.StringAttribute{
							MarkdownDescription: "Product name",
							Computed:            true,
						},
						"product_version": schema.StringAttribute{
							MarkdownDescription: "Product version",
							Computed:            true,
						},
						"product_title": schema.StringAttribute{
							MarkdownDescription: "Product title",
							Computed:            true,
						},
						"use_rate_limit_group": schema.BoolAttribute{
							MarkdownDescription: "Use rate limit group",
							Computed:            true,
						},
						"rate_limit": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the rate limit scheme to enforce. This scheme defines the maximum rate to allow during a specified interval and the actions to take when the limit is exceeded.",
							NestedObject:        models.DmAPIRateLimitDataSourceSchema,
							Computed:            true,
						},
						"burst_limit": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the burst limit scheme to enforce. This scheme defines the maximum burst rate to allow during a specified interval. The burst limit helps to prevent spikes that might damage the infrastructure. When a message arrives within an interval, the burst limit takes priority over the rate limit. In other words, a message is first checked against the burst limit scheme and then against the rate limit scheme.",
							NestedObject:        models.DmAPIBurstLimitDataSourceSchema,
							Computed:            true,
						},
						"rate_limit_group": schema.StringAttribute{
							MarkdownDescription: "Rate limit group",
							Computed:            true,
						},
						"use_limit_definitions": schema.BoolAttribute{
							MarkdownDescription: "Use limit definitions",
							Computed:            true,
						},
						"assembly_burst_limit": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the burst limit scheme that the rate limit assembly action enforces. This scheme defines the maximum burst rate to allow during a specified interval. This scheme helps to prevent spikes that might damage the infrastructure. When a message arrives within an interval, the burst limit takes priority over the rate limit. In other words, a message is first checked against the burst limit scheme and then against the rate limit scheme.",
							NestedObject:        models.DmAPIBurstLimitDataSourceSchema,
							Computed:            true,
						},
						"assembly_burst_limit_definition": schema.ListNestedAttribute{
							MarkdownDescription: "Specify a burst limit definition that the rate limit assembly action enforces. A burst limit definition defines the maximum burst rate to allow during a specified interval. This scheme helps to prevent spikes that might damage infrastructure. When a message arrives within an interval, the burst limit takes priority over the rate limit. A message is first checked against the burst limit scheme and then against the rate limit scheme.",
							NestedObject:        models.DmDefinitionLinkDataSourceSchema,
							Computed:            true,
						},
						"assembly_rate_limit": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the rate limit scheme that the rate limit assembly action enforces. This scheme defines the maximum rate to allow during a specified interval and the actions to take when the limit is exceeded.",
							NestedObject:        models.DmAPIRateLimitDataSourceSchema,
							Computed:            true,
						},
						"assembly_rate_limit_definition": schema.ListNestedAttribute{
							MarkdownDescription: "Specify a rate limit definition that the rate limit assembly action enforces. A rate limit definition defines the maximum rate that is allowed in a specified interval and the actions to take when the limit is exceeded.",
							NestedObject:        models.DmDefinitionLinkDataSourceSchema,
							Computed:            true,
						},
						"assembly_count_limit": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the count limit scheme that the rate limit assembly action enforces. This scheme defines the maximum count to allow and the actions to take when the limit is exceeded.",
							NestedObject:        models.DmAPICountLimitDataSourceSchema,
							Computed:            true,
						},
						"assembly_count_limit_definition": schema.ListNestedAttribute{
							MarkdownDescription: "Specify a count limit definition that the rate limit assembly action enforces. A count limit definition defines the maximum count that is allowed and the actions to take when the limit is exceeded.",
							NestedObject:        models.DmDefinitionLinkDataSourceSchema,
							Computed:            true,
						},
						"space_id": schema.StringAttribute{
							MarkdownDescription: "Specify the space ID for the product in the catalog. When space is enabled for a catalog, the catalog can be partitioned to spaces. Spaces enable each team to manage their APIs independently.",
							Computed:            true,
						},
						"space_name": schema.StringAttribute{
							MarkdownDescription: "Specify the space name for the product in the catalog. When space is enabled for a catalog, the catalog can be partitioned to spaces. Spaces enable each team to manage their APIs independently.",
							Computed:            true,
						},
						"api": schema.ListAttribute{
							MarkdownDescription: "Specify the APIs to package for the plan. An API is exposed through a plan by associating the API to the plan.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"exclude_operation": schema.ListAttribute{
							MarkdownDescription: "Exclude operation",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"override": schema.ListAttribute{
							MarkdownDescription: "Operation rate limit",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"rate_limit_scope": schema.StringAttribute{
							MarkdownDescription: "Specify the scope to apply the rate limit schemes to. You can apply schemes against the application or client ID. For example, <tt>application1</tt> has <tt>client1</tt> and <tt>client2</tt> , and the rate limit is 10 calls per hour. <ul><li>When against the application, <tt>application1</tt> limits 10 calls per hour from either <tt>client1</tt> or <tt>client2.</tt></li><li>When against the client ID, <tt>application1</tt> limits 10 calls per hour from each <tt>client1</tt> and <tt>client2</tt> .</li></ul>",
							Computed:            true,
						},
						"graph_ql_schema_options": schema.ListAttribute{
							MarkdownDescription: "GraphQL schema options",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *APIPlanDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *APIPlanDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data APIPlanList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.APIPlan{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.APIPlan{}
	if value := res.Get(`APIPlan`); value.Exists() {
		for _, v := range value.Array() {
			item := models.APIPlan{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.APIPlanObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.APIPlanObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
