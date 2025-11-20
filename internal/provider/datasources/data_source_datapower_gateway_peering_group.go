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

type GatewayPeeringGroupList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &GatewayPeeringGroupDataSource{}
	_ datasource.DataSourceWithConfigure = &GatewayPeeringGroupDataSource{}
)

func NewGatewayPeeringGroupDataSource() datasource.DataSource {
	return &GatewayPeeringGroupDataSource{}
}

type GatewayPeeringGroupDataSource struct {
	pData *tfutils.ProviderData
}

func (d *GatewayPeeringGroupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_peering_group"
}

func (d *GatewayPeeringGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "A gateway-peering group defines members as a group to synchronize data across members. When a group can work in stand-alone mode, peer-based mode, or cluster-based mode. <ul><li>For stand-alone, no peers defined. This mode is for only development or testing purposes.</li><li>For a peer group, add peers and configure the connection among the peers.</li><li>For a cluster, add cluster nodes and the other nodes that are in the same data center.</li></ul>",
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
						"mode": schema.StringAttribute{
							MarkdownDescription: "Mode",
							Computed:            true,
						},
						"peer_nodes": schema.ListNestedAttribute{
							MarkdownDescription: "Specify peers for the the group. To add a peer, enter its local IP address or host alias and its priority.",
							NestedObject:        models.GetDmGatewayPeeringGroupPeerNodeDataSourceSchema(),
							Computed:            true,
						},
						"cluster_primary_count": schema.StringAttribute{
							MarkdownDescription: "Primary count",
							Computed:            true,
						},
						"cluster_nodes": schema.ListNestedAttribute{
							MarkdownDescription: "Specify nodes for the cluster group. To add a node, enter its local IP address or host alias and the comma-separated list of local IP addresses or host aliases of the other nodes that are in the same data center. <p>Because the primary count is 3, the configuration requires a minimum of 6 nodes that are generally in 2 data centers. Each node is defined on a different DataPower Gateway. The minimal configuration is 3 primary-secondary pairs. Each pair is a shard that manages a subset of slots.</p><p>Each primary node can have more than one secondary node, but each primary node requires the same number of secondary nodes. In other words, you can define an environment of 9 nodes, which is a configuration of 3 primary nodes and 6 secondary nodes. In this configuration, each primary node has 2 secondary nodes.</p>",
							NestedObject:        models.GetDmGatewayPeeringGroupClusterNodeDataSourceSchema(),
							Computed:            true,
						},
						"cluster_auto_config": schema.BoolAttribute{
							MarkdownDescription: "Specify whether the cluster configuration is managed automatically. By default, cluster configuration is managed automatically. Unless directed by IBM Support, do not change this setting.",
							Computed:            true,
						},
						"enable_ssl": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to use TLS to secure the connection among the members. By default, TLS is enabled. When enabled, ensure that all members use the same TLS configuration.",
							Computed:            true,
						},
						"idcred": schema.StringAttribute{
							MarkdownDescription: "Specify the identification credentials that contains the credentials that the current member uses to identify itself to other peers. Client authentication uses mutual TLS.",
							Computed:            true,
						},
						"valcred": schema.StringAttribute{
							MarkdownDescription: "Validation credentials",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *GatewayPeeringGroupDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *GatewayPeeringGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data GatewayPeeringGroupList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.GatewayPeeringGroup{
		AppDomain: data.AppDomain,
	}

	path := o.GetPath()
	if !data.Id.IsNull() {
		path = path + "/" + data.Id.ValueString()
	}

	res, err := d.pData.Client.Get(path)
	resFound := true
	if err != nil {
		if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(d.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString())
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		} else {
			resFound = false
		}
	}
	l := []models.GatewayPeeringGroup{}
	if resFound {
		if value := res.Get(`GatewayPeeringGroup`); value.Exists() {
			for _, v := range value.Array() {
				item := models.GatewayPeeringGroup{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.GatewayPeeringGroupObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.GatewayPeeringGroupObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
