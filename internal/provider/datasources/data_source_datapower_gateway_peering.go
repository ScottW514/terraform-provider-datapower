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

type GatewayPeeringList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &GatewayPeeringDataSource{}
	_ datasource.DataSourceWithConfigure = &GatewayPeeringDataSource{}
)

func NewGatewayPeeringDataSource() datasource.DataSource {
	return &GatewayPeeringDataSource{}
}

type GatewayPeeringDataSource struct {
	pData *tfutils.ProviderData
}

func (d *GatewayPeeringDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_peering"
}

func (d *GatewayPeeringDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "A gateway-peering instance defines how members synchronize data across members.",
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
						"password_alias": schema.StringAttribute{
							MarkdownDescription: "Specify the password alias to secure the data store. If not specified, a system default is used. The use of the system default is classified as a security vulnerability (CVE-2022-31776).",
							Computed:            true,
						},
						"local_address": schema.StringAttribute{
							MarkdownDescription: "Specify the IP address or host alias that the gateway service listens on. The IP address can be any DataPower network interface that can be accessed by other peers in the peer group. The IP address cannot be 127.0.0.1, 0.0.0.0 or ::.",
							Computed:            true,
						},
						"local_port": schema.Int64Attribute{
							MarkdownDescription: "Specify the port that the gateway service listens on. The default value is 16380. Ensure that all peers use the same port.",
							Computed:            true,
						},
						"peer_group": schema.StringAttribute{
							MarkdownDescription: "Gateway-peering group",
							Computed:            true,
						},
						"monitor_port": schema.Int64Attribute{
							MarkdownDescription: "Specify the port to monitor for state synchronization. The default value is 26380. Ensure that all peers use the same port.",
							Computed:            true,
						},
						"enable_peer_group": schema.BoolAttribute{
							MarkdownDescription: "Specify whether the gateway-peering instance uses a peer group. <ul><li>When enabled, the instance works in the mode that is set for the peer group. This setting is the default value.</li><li>When not enabled, the instance works in stand-alone mode.</li></ul>",
							Computed:            true,
						},
						"enable_ssl": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to use TLS to secure the connection among the members. By default, TLS is enabled. In peer-based mode, ensure that all peers use the same TLS configuration.",
							Computed:            true,
						},
						"persistence_location": schema.StringAttribute{
							MarkdownDescription: "Specify where to store data. Ensure that all peers in the group store data in the same location.",
							Computed:            true,
						},
						"local_directory": schema.StringAttribute{
							MarkdownDescription: "Specify the directory to store data. For example, <tt>local:///group1</tt> .",
							Computed:            true,
						},
						"max_memory": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum memory for the data store. When memory reaches this limit, data is removed by using the least recently used (LRU) algorithm. The default value is 0, which means no limits. Do not over allocate memory.",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *GatewayPeeringDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *GatewayPeeringDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data GatewayPeeringList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.GatewayPeering{
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
	l := []models.GatewayPeering{}
	if resFound {
		if value := res.Get(`GatewayPeering`); value.Exists() {
			for _, v := range value.Array() {
				item := models.GatewayPeering{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.GatewayPeeringObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.GatewayPeeringObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
