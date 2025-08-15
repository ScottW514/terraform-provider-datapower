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
	resp.TypeName = req.ProviderTypeName + "_rate_limit_definition"
}

func (d *RateLimitDefinitionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "A rate limit definition contains the parameters to enforce a rate limit.",
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
							MarkdownDescription: "Specify the name shown in rate limit response headers. If not specified, the object name is used in response headers.",
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
							MarkdownDescription: "Specify the maximum number of requests that the API gateway can handle in an interval. The value of 0 indicates no limit.",
							Computed:            true,
						},
						"interval": schema.Int64Attribute{
							MarkdownDescription: "Specify the interval for the rate limit. Enter a value that is greater than or equal to 1. The default value is 1.",
							Computed:            true,
						},
						"unit": schema.StringAttribute{
							MarkdownDescription: "Specify the time unit for the rate limit. The default value is minute. When type is burst, the unit can be second or minute.",
							Computed:            true,
						},
						"hard_limit": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to reject requests when the specified rate limit is exceeded. By default, requests are rejected when the limit is exceeded. When disabled, requests are accepted but a warning is logged.",
							Computed:            true,
						},
						"is_client": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to apply the rate limit to the client or to an internal component. By default, the rate limit is applied to the client. Client rate limits return a 429 error when exceeded. When disabled, rate limit information is not applied to the client. Non-client rate limits return a 503 error when exceeded.",
							Computed:            true,
						},
						"use_api_name": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to include the API name in the rate limit key. By default, the API name is not included. When enabled, the API name is included.",
							Computed:            true,
						},
						"use_app_id": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to include the application ID in the rate limit key. By default, the application ID is not included. When enabled, the application ID is included.",
							Computed:            true,
						},
						"use_client_id": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to include the client ID in the rate limit key. By default, the client ID is not included. When enabled, the client ID is included.",
							Computed:            true,
						},
						"auto_replenish": schema.BoolAttribute{
							MarkdownDescription: "Specify whether the count limit is automatically replenished at the end of the transaction. By default, the count limit is automatically replenished. When disabled, the count limit is replenished only by applying a rate limit assembly action that contains the count limit with a replenish operation.",
							Computed:            true,
						},
						"dynamic_value": schema.StringAttribute{
							MarkdownDescription: "Specify the dynamic value string for the rate limit, which should contain one or more context variables. The default value is an empty string. <p>The dynamic value makes it possible to use a context variable to enforce the rate limit based on parameters other than those defined in the rate limit scheme, such as a username, incoming IP address, or server name. The context variable can be set in a GatewayScript action and then included in the dynamic value.</p><p>The following example uses the context object in a GatewayScript action to add the <tt>my.server</tt> variable to the API context. The dynamic value can then include the variable <tt>my.server</tt> , which resolves to the server name <tt>server34</tt> .</p><p><tt>context.set(\"my.server\", \"server34\")</tt></p>",
							Computed:            true,
						},
						"weight": schema.StringAttribute{
							MarkdownDescription: "Specify a JSONata expression to assign a weight value to the rate limit. For each API call, the value computed by the weight expression is applied to the rate limit. The default value is 1. If the weight expression evaluates to a value that is less than or equal to 0, it is set to 1. An empty string results in an error.",
							Computed:            true,
						},
						"response_headers": schema.BoolAttribute{
							MarkdownDescription: "Specify whether response headers include rate limit information. By default, headers include rate limit information. When disabled, headers exclude rate limit information.",
							Computed:            true,
						},
						"emulate_burst_headers": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to return information about the rate limit in burst limit response headers instead of in rate limit response headers. By default, the information is in rate limit headers. When enabled, information is in burst limit headers.",
							Computed:            true,
						},
						"use_interval_offset": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to allow limit intervals to start at different offsets. By default, intervals can start at different offsets. When disabled, intervals cannot start at different offsets.",
							Computed:            true,
						},
						"allow_cache_fallback": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to use the cache as a fallback when gateway-peering instances cannot be contacted. By default, the cache can enforce rate limits when the cache is disabled. When disabled, the cache cannot enforce rate limits.",
							Computed:            true,
						},
						"use_cache": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to use a cache to store rate limit information. A cache might be faster when the number of API calls is low. A cache can cause degraded performance when the number of API calls is exceptionally high. By default, a cache cannot store information. When enabled, the cache can store information.",
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
