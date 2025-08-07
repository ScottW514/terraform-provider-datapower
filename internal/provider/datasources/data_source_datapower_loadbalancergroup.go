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

type LoadBalancerGroupList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &LoadBalancerGroupDataSource{}
	_ datasource.DataSourceWithConfigure = &LoadBalancerGroupDataSource{}
)

func NewLoadBalancerGroupDataSource() datasource.DataSource {
	return &LoadBalancerGroupDataSource{}
}

type LoadBalancerGroupDataSource struct {
	client *client.DatapowerClient
}

func (d *LoadBalancerGroupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_loadbalancergroup"
}

func (d *LoadBalancerGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Load Balancer Group",
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
						"algorithm": schema.StringAttribute{
							MarkdownDescription: "Algorithm",
							Computed:            true,
						},
						"retrieve_info": schema.BoolAttribute{
							MarkdownDescription: "Retrieve Workload Management Information",
							Computed:            true,
						},
						"wlm_retrieval": schema.StringAttribute{
							MarkdownDescription: "Workload Management Retrieval",
							Computed:            true,
						},
						"web_sphere_cell_config": schema.StringAttribute{
							MarkdownDescription: "WebSphere Cell",
							Computed:            true,
						},
						"wlm_group": schema.StringAttribute{
							MarkdownDescription: "Workload Management Group Name",
							Computed:            true,
						},
						"wlm_transport": schema.StringAttribute{
							MarkdownDescription: "Protocol",
							Computed:            true,
						},
						"damp": schema.Int64Attribute{
							MarkdownDescription: "Damp Time",
							Computed:            true,
						},
						"never_return_sick_member": schema.BoolAttribute{
							MarkdownDescription: "Do not Bypass Down State",
							Computed:            true,
						},
						"lb_group_members": schema.ListNestedAttribute{
							MarkdownDescription: "Members",
							NestedObject:        models.DmLBGroupMemberDataSourceSchema,
							Computed:            true,
						},
						"try_every_server_before_failing": schema.BoolAttribute{
							MarkdownDescription: "Try Every Server Before Failing",
							Computed:            true,
						},
						"lb_group_checks": models.GetDmLBGroupCheckDataSourceSchema("Health Checks", "health-check", ""),
						"masquerade_member": schema.BoolAttribute{
							MarkdownDescription: "Masquerade As Group Name",
							Computed:            true,
						},
						"application_routing": schema.BoolAttribute{
							MarkdownDescription: "Enable Application Routing",
							Computed:            true,
						},
						"lb_group_affinity_conf": models.GetDmLBGroupAffinityDataSourceSchema("Session Affinity", "session-affinity", ""),
						"monitored_cookies": schema.ListAttribute{
							MarkdownDescription: "Monitored Cookies",
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

func (d *LoadBalancerGroupDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *LoadBalancerGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data LoadBalancerGroupList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.LoadBalancerGroup{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.LoadBalancerGroup{}
	if value := res.Get(`LoadBalancerGroup`); value.Exists() {
		for _, v := range value.Array() {
			item := models.LoadBalancerGroup{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.LoadBalancerGroupObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.LoadBalancerGroupObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
