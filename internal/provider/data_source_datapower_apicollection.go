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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
)

type APICollectionList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &APICollectionDataSource{}
	_ datasource.DataSourceWithConfigure = &APICollectionDataSource{}
)

func NewAPICollectionDataSource() datasource.DataSource {
	return &APICollectionDataSource{}
}

type APICollectionDataSource struct {
	client *client.DatapowerClient
}

func (d *APICollectionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apicollection"
}

func (d *APICollectionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "API collection",
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
						"sandbox": schema.BoolAttribute{
							MarkdownDescription: "Sandbox",
							Computed:            true,
						},
						"org_id": schema.StringAttribute{
							MarkdownDescription: "Organization ID",
							Computed:            true,
						},
						"org_name": schema.StringAttribute{
							MarkdownDescription: "Organization name",
							Computed:            true,
						},
						"catalog_id": schema.StringAttribute{
							MarkdownDescription: "Catalog ID",
							Computed:            true,
						},
						"catalog_name": schema.StringAttribute{
							MarkdownDescription: "Catalog name",
							Computed:            true,
						},
						"dev_portal_endpoint": schema.StringAttribute{
							MarkdownDescription: "Developer Portal endpoint",
							Computed:            true,
						},
						"cache_capacity": schema.Int64Attribute{
							MarkdownDescription: "Subscriber cache capacity",
							Computed:            true,
						},
						"routing_prefix": schema.ListNestedAttribute{
							MarkdownDescription: "Routing prefixes",
							NestedObject:        models.DmRoutingPrefixDataSourceSchema,
							Computed:            true,
						},
						"use_rate_limit_group": schema.BoolAttribute{
							MarkdownDescription: "Use rate limit group",
							Computed:            true,
						},
						"default_rate_limit": schema.ListNestedAttribute{
							MarkdownDescription: "Default rate limit",
							NestedObject:        models.DmAPIRateLimitDataSourceSchema,
							Computed:            true,
						},
						"rate_limit_group": schema.StringAttribute{
							MarkdownDescription: "Rate limit group",
							Computed:            true,
						},
						"assembly_burst_limit": schema.ListNestedAttribute{
							MarkdownDescription: "Assembly burst limit",
							NestedObject:        models.DmAPIBurstLimitDataSourceSchema,
							Computed:            true,
						},
						"assembly_rate_limit": schema.ListNestedAttribute{
							MarkdownDescription: "Assembly rate limit",
							NestedObject:        models.DmAPIRateLimitDataSourceSchema,
							Computed:            true,
						},
						"assembly_count_limit": schema.ListNestedAttribute{
							MarkdownDescription: "Assembly count limit",
							NestedObject:        models.DmAPICountLimitDataSourceSchema,
							Computed:            true,
						},
						"enforce_pre_assembly_rate_limits": schema.BoolAttribute{
							MarkdownDescription: "Enforce preassembly rate limits",
							Computed:            true,
						},
						"api_processing_rule": schema.StringAttribute{
							MarkdownDescription: "API processing rule",
							Computed:            true,
						},
						"api_error_rule": schema.StringAttribute{
							MarkdownDescription: "API error rule",
							Computed:            true,
						},
						"assembly_preflow": schema.StringAttribute{
							MarkdownDescription: "Assembly preprocessing",
							Computed:            true,
						},
						"assembly_postflow": schema.StringAttribute{
							MarkdownDescription: "Assembly postprocessing",
							Computed:            true,
						},
						"plan": schema.ListAttribute{
							MarkdownDescription: "Plans",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"analytics_endpoint": schema.StringAttribute{
							MarkdownDescription: "Analytic endpoint",
							Computed:            true,
						},
						"application_type": schema.ListAttribute{
							MarkdownDescription: "Application types",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"parse_settings_reference": models.GetDmDynamicParseSettingsReferenceDataSourceSchema("Parse settings", "parse-settings-reference", ""),
					},
				},
			},
		},
	}
}

func (d *APICollectionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *APICollectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data APICollectionList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.APICollection{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.APICollection{}
	if value := res.Get(`APICollection`); value.Exists() {
		for _, v := range value.Array() {
			item := models.APICollection{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.APICollectionObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.APICollectionObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
