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
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
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
	client *client.DatapowerClient
}

func (d *APIPlanDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apiplan"
}

func (d *APIPlanDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "API plan",
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
							MarkdownDescription: "Product ID",
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
							MarkdownDescription: "Rate limit",
							NestedObject:        models.DmAPIRateLimitDataSourceSchema,
							Computed:            true,
						},
						"burst_limit": schema.ListNestedAttribute{
							MarkdownDescription: "Burst limit",
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
							MarkdownDescription: "Assembly burst limit",
							NestedObject:        models.DmAPIBurstLimitDataSourceSchema,
							Computed:            true,
						},
						"assembly_burst_limit_definition": schema.ListNestedAttribute{
							MarkdownDescription: "Assembly burst limit definition",
							NestedObject:        models.DmDefinitionLinkDataSourceSchema,
							Computed:            true,
						},
						"assembly_rate_limit": schema.ListNestedAttribute{
							MarkdownDescription: "Assembly rate limit",
							NestedObject:        models.DmAPIRateLimitDataSourceSchema,
							Computed:            true,
						},
						"assembly_rate_limit_definition": schema.ListNestedAttribute{
							MarkdownDescription: "Assembly rate limit definition",
							NestedObject:        models.DmDefinitionLinkDataSourceSchema,
							Computed:            true,
						},
						"assembly_count_limit": schema.ListNestedAttribute{
							MarkdownDescription: "Assembly count limit",
							NestedObject:        models.DmAPICountLimitDataSourceSchema,
							Computed:            true,
						},
						"assembly_count_limit_definition": schema.ListNestedAttribute{
							MarkdownDescription: "Assembly count limit definition",
							NestedObject:        models.DmDefinitionLinkDataSourceSchema,
							Computed:            true,
						},
						"space_id": schema.StringAttribute{
							MarkdownDescription: "Space ID",
							Computed:            true,
						},
						"space_name": schema.StringAttribute{
							MarkdownDescription: "Space name",
							Computed:            true,
						},
						"api": schema.ListAttribute{
							MarkdownDescription: "API",
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
							MarkdownDescription: "Rate limit scope",
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

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *APIPlanDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data APIPlanList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.APIPlan{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
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
