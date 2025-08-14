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

type RateLimitDefinitionList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &RateLimitDefinitionDataSource{}
	_ datasource.DataSourceWithConfigure = &RateLimitDefinitionDataSource{}
)

func NewRateLimitDefinitionDataSource() datasource.DataSource {
	return &RateLimitDefinitionDataSource{}
}

type RateLimitDefinitionDataSource struct {
	pData *tfutils.ProviderData
}

func (d *RateLimitDefinitionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ratelimitdefinition"
}

func (d *RateLimitDefinitionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Rate limit definition",
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
						"short_name": schema.StringAttribute{
							MarkdownDescription: "Short name",
							Computed:            true,
						},
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Limit type",
							Computed:            true,
						},
						"rate": schema.Int64Attribute{
							MarkdownDescription: "Rate",
							Computed:            true,
						},
						"interval": schema.Int64Attribute{
							MarkdownDescription: "Interval",
							Computed:            true,
						},
						"unit": schema.StringAttribute{
							MarkdownDescription: "Unit",
							Computed:            true,
						},
						"hard_limit": schema.BoolAttribute{
							MarkdownDescription: "Enable hard limit",
							Computed:            true,
						},
						"is_client": schema.BoolAttribute{
							MarkdownDescription: "Is client",
							Computed:            true,
						},
						"use_api_name": schema.BoolAttribute{
							MarkdownDescription: "Use API name",
							Computed:            true,
						},
						"use_app_id": schema.BoolAttribute{
							MarkdownDescription: "Use application ID",
							Computed:            true,
						},
						"use_client_id": schema.BoolAttribute{
							MarkdownDescription: "Use client ID",
							Computed:            true,
						},
						"auto_replenish": schema.BoolAttribute{
							MarkdownDescription: "Autoreplenish",
							Computed:            true,
						},
						"dynamic_value": schema.StringAttribute{
							MarkdownDescription: "Dynamic value",
							Computed:            true,
						},
						"weight": schema.StringAttribute{
							MarkdownDescription: "Weight expression",
							Computed:            true,
						},
						"response_headers": schema.BoolAttribute{
							MarkdownDescription: "Response headers",
							Computed:            true,
						},
						"emulate_burst_headers": schema.BoolAttribute{
							MarkdownDescription: "Emulate burst headers",
							Computed:            true,
						},
						"use_interval_offset": schema.BoolAttribute{
							MarkdownDescription: "Use interval offset",
							Computed:            true,
						},
						"allow_cache_fallback": schema.BoolAttribute{
							MarkdownDescription: "Allow cache fallback",
							Computed:            true,
						},
						"use_cache": schema.BoolAttribute{
							MarkdownDescription: "Use cache",
							Computed:            true,
						},
						"parameters": schema.ListNestedAttribute{
							MarkdownDescription: "Parameters",
							NestedObject:        models.DmRateLimitDefinitionNameValuePairDataSourceSchema,
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *RateLimitDefinitionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *RateLimitDefinitionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data RateLimitDefinitionList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.RateLimitDefinition{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.RateLimitDefinition{}
	if value := res.Get(`RateLimitDefinition`); value.Exists() {
		for _, v := range value.Array() {
			item := models.RateLimitDefinition{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.RateLimitDefinitionObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.RateLimitDefinitionObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
