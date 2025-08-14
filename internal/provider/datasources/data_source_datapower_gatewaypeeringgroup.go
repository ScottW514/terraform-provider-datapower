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

type GatewayPeeringGroupList struct {
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
	resp.TypeName = req.ProviderTypeName + "_gatewaypeeringgroup"
}

func (d *GatewayPeeringGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Gateway-peering group",
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
						"mode": schema.StringAttribute{
							MarkdownDescription: "Mode",
							Computed:            true,
						},
						"peer_nodes": schema.ListNestedAttribute{
							MarkdownDescription: "Peers",
							NestedObject:        models.DmGatewayPeeringGroupPeerNodeDataSourceSchema,
							Computed:            true,
						},
						"cluster_primary_count": schema.StringAttribute{
							MarkdownDescription: "Primary count",
							Computed:            true,
						},
						"cluster_nodes": schema.ListNestedAttribute{
							MarkdownDescription: "Cluster nodes",
							NestedObject:        models.DmGatewayPeeringGroupClusterNodeDataSourceSchema,
							Computed:            true,
						},
						"cluster_auto_config": schema.BoolAttribute{
							MarkdownDescription: "Auto manage cluster configuration",
							Computed:            true,
						},
						"enable_ssl": schema.BoolAttribute{
							MarkdownDescription: "Enable TLS",
							Computed:            true,
						},
						"idcred": schema.StringAttribute{
							MarkdownDescription: "Identification credentials",
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

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.GatewayPeeringGroup{}
	if value := res.Get(`GatewayPeeringGroup`); value.Exists() {
		for _, v := range value.Array() {
			item := models.GatewayPeeringGroup{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
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
