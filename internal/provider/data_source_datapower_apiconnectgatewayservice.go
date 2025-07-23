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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
)

var (
	_ datasource.DataSource              = &APIConnectGatewayServiceDataSource{}
	_ datasource.DataSourceWithConfigure = &APIConnectGatewayServiceDataSource{}
)

func NewAPIConnectGatewayServiceDataSource() datasource.DataSource {
	return &APIConnectGatewayServiceDataSource{}
}

type APIConnectGatewayServiceDataSource struct {
	client *client.DatapowerClient
}

func (d *APIConnectGatewayServiceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apiconnectgatewayservice"
}

func (d *APIConnectGatewayServiceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "API Connect gateway service",
		Attributes: map[string]schema.Attribute{
			"app_domain": schema.StringAttribute{
				MarkdownDescription: "The name of the application domain the object belongs to",
				Required:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Administrative state",
				Computed:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: "Comments",
				Computed:            true,
			},
			"local_address": schema.StringAttribute{
				MarkdownDescription: "Local address",
				Computed:            true,
			},
			"local_port": schema.Int64Attribute{
				MarkdownDescription: "Local port",
				Computed:            true,
			},
			"ssl_server": schema.StringAttribute{
				MarkdownDescription: "TLS server profile",
				Computed:            true,
			},
			"api_gateway_address": schema.StringAttribute{
				MarkdownDescription: "API gateway address",
				Computed:            true,
			},
			"api_gateway_port": schema.Int64Attribute{
				MarkdownDescription: "API gateway port",
				Computed:            true,
			},
			"gateway_peering": schema.StringAttribute{
				MarkdownDescription: "Gateway peering",
				Computed:            true,
			},
			"gateway_peering_manager": schema.StringAttribute{
				MarkdownDescription: "Gateway-peering manager",
				Computed:            true,
			},
			"v5_compatibility_mode": schema.BoolAttribute{
				MarkdownDescription: "V5 compatibility mode",
				Computed:            true,
			},
			"user_defined_policies": schema.ListAttribute{
				MarkdownDescription: "User-defined policies",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"v5c_slm_mode": schema.StringAttribute{
				MarkdownDescription: "SLM peer mode",
				Computed:            true,
			},
			"ip_multicast": schema.StringAttribute{
				MarkdownDescription: "IP multicast",
				Computed:            true,
			},
			"ip_unicast": schema.StringAttribute{
				MarkdownDescription: "IP unicast",
				Computed:            true,
			},
			"jwt_validation_mode": schema.StringAttribute{
				MarkdownDescription: "JWT validation mode",
				Computed:            true,
			},
			"jwturl": schema.StringAttribute{
				MarkdownDescription: "JWT URL",
				Computed:            true,
			},
			"proxy_policy": models.GetDmAPICGSProxyPolicyDataSourceSchema("API Manager proxy", "proxy", ""),
		},
	}
}

func (d *APIConnectGatewayServiceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *APIConnectGatewayServiceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.APIConnectGatewayService

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := d.client.Get(data.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	data.FromBody(ctx, `APIConnectGatewayService`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
