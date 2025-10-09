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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var (
	_ datasource.DataSource              = &APIConnectGatewayServiceDataSource{}
	_ datasource.DataSourceWithConfigure = &APIConnectGatewayServiceDataSource{}
)

func NewAPIConnectGatewayServiceDataSource() datasource.DataSource {
	return &APIConnectGatewayServiceDataSource{}
}

type APIConnectGatewayServiceDataSource struct {
	pData *tfutils.ProviderData
}

func (d *APIConnectGatewayServiceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_connect_gateway_service"
}

func (d *APIConnectGatewayServiceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The API Connect gateway service defines the type of gateway service and manages connections with API Connect. When configured, the DataPower Gateway creates a gateway service to retrieve data from API Connect to define the configuration to process API requests.",
		Attributes: map[string]schema.Attribute{
			"app_domain": schema.StringAttribute{
				MarkdownDescription: "The name of the application domain the object belongs to",
				Required:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "<p>The administrative state of the configuration.</p><ul><li>To make active, set to enabled.</li><li>To make inactive, set to disabled.</li></ul>",
				Computed:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: "Comments",
				Computed:            true,
			},
			"local_address": schema.StringAttribute{
				MarkdownDescription: "Specify the IP address or interface through that API Connect uses to manage the gateway service. The default value is 0.0.0.0.",
				Computed:            true,
			},
			"local_port": schema.Int64Attribute{
				MarkdownDescription: "Specify the listening port for the gateway service. The default value is 3000. <p><b>Note:</b> The gateway service uses four additional consecutive ports after the local port. Therefore, all five consecutive ports must be clear of conflicts.</p>",
				Computed:            true,
			},
			"ssl_server": schema.StringAttribute{
				MarkdownDescription: "Specify the TLS server profile to secure connections between API Connect to the gateway service. The following restrictions apply. <ul><li>Keys and certificates are restricted to PEM and PKCS #12 formats.</li><li>The validation credentials must use PEM formatted material.</li></ul>",
				Computed:            true,
			},
			"api_gateway_address": schema.StringAttribute{
				MarkdownDescription: "Specify the IP address or host alias to accept API requests. The default value is 0.0.0.0. This address is used with its port to create an HTTPS handler.",
				Computed:            true,
			},
			"api_gateway_port": schema.Int64Attribute{
				MarkdownDescription: "Specify the listening port for API requests. The default value is 9443. This port is used with its address to create an HTTPS handler.",
				Computed:            true,
			},
			"gateway_peering": schema.StringAttribute{
				MarkdownDescription: "Specify the gateway-peering instance that manages data across the gateway peers. The following restrictions apply. <ul><li>When TLS and peer group mode are enabled, all peers must use the same crypto material.</li><li>Keys and certificates are restricted to PEM and PKCS #12 formats.</li></ul>",
				Computed:            true,
			},
			"gateway_peering_manager": schema.StringAttribute{
				MarkdownDescription: "Specify the gateway-peering manager that manages gateway-peering instances for the gateway service. This property is meaningful when the gateway type is an API gateway.",
				Computed:            true,
			},
			"v5compatibility_mode": schema.BoolAttribute{
				MarkdownDescription: "Specify whether the gateway service is a Multi-Protocol Gateway or an API gateway. <ui><li>When enabled, the gateway service is a Multi-Protocol Gateway that is compatible with API Connect version 5.</li><li>When disabled, that gateway service is an API gateway this is not compatible with API Connect v5.</li></ui>",
				Computed:            true,
			},
			"user_defined_policies": schema.ListAttribute{
				MarkdownDescription: "Specify user-defined policies to advertise to API Connect for use in the API Connect Assembly Editor. This property is meaningful when the gateway type is an API gateway. <p>For an assembly function that is a user-defined policy, configure the assembly function with a mechanism other than a watched file that is processed by a configuration sequence. Objects that are created through the processing of configuration sequences are not persisted to the startup configuration. The preferred method for user-defined policies is to define them explicitly so that they persist to the startup configuration.</p>",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"v5c_slm_mode": schema.StringAttribute{
				MarkdownDescription: "Specify the peer group type for the SLM policy. This property is meaningful when the gateway type is a Multi-Protocol Gateway.",
				Computed:            true,
			},
			"ip_multicast": schema.StringAttribute{
				MarkdownDescription: "Specify the IP multicast configuration for the SLM policy. This property is meaningful when the gateway type is a Multi-Protocol Gateway and the peer mode is multicast.",
				Computed:            true,
			},
			"ip_unicast": schema.StringAttribute{
				MarkdownDescription: "Specify the address of the unicast peer group for the SLM policy. This property is meaningful when the gateway type is a Multi-Protocol Gateway and the peer mode is unicast.",
				Computed:            true,
			},
			"jwt_validation_mode": schema.StringAttribute{
				MarkdownDescription: "Specify the JWT validation mode. This property does not control whether a token is validated. This property controls whether transactions fail when validation fails.",
				Computed:            true,
			},
			"jwt_url": schema.StringAttribute{
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

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *APIConnectGatewayServiceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data models.APIConnectGatewayService
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	path := data.GetPath()

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
	if resFound {
		data.FromBody(ctx, `APIConnectGatewayService`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
