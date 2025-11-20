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

type LoadBalancerGroupList struct {
	Id        types.String `tfsdk:"id"`
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
	pData *tfutils.ProviderData
}

func (d *LoadBalancerGroupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_load_balancer_group"
}

func (d *LoadBalancerGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "<p>The DataPower device distributes traffic to members of a Load Balancer Group. These are back end servers and not additional DataPower devices. A Load Balancer Group lists members of a virtual server group and sets the algorithm for balancing them. Periodic health checks can be performed. Load Balancers may also be used to provide redundant LDAP server access.</p><p>When created, a DataPower service can use a Load Balancer Group by associating it with an XML manager that is associated with this service.</p><p>The back end destination URL is set to the name of the Load Balancer Group (example: \"BackEndServers\").</p>",
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
						"algorithm": schema.StringAttribute{
							MarkdownDescription: "Select the algorithm to use to balance the real servers.",
							Computed:            true,
						},
						"retrieve_info": schema.BoolAttribute{
							MarkdownDescription: "Use this setting to control whether this Load Balancer Group has membership and weight information automatically retrieved from the work load management repository WebSphere Cell. When disabled, the static configuration is used.",
							Computed:            true,
						},
						"wlm_retrieval": schema.StringAttribute{
							MarkdownDescription: "Contains the back end work load management repository selection type. Select 'WebSphere Cell' if your back-end is a WebSphere Application Server (WAS) Network Deployment (ND) or WAS Virtual Enterprise (VE).",
							Computed:            true,
						},
						"web_sphere_cell_config": schema.StringAttribute{
							MarkdownDescription: "If you selected 'WebSphere Cell' for Workload Management Retrieval, you need to select a WebSphere Cell object that retrieves this information. If no objects are available in the pull down, you must create one.",
							Computed:            true,
						},
						"wlm_group": schema.StringAttribute{
							MarkdownDescription: "The Workload Management Group Name is used to define a group. In a WebSphere Application Server environment, the back end group is a cluster name. Once specified, the Load Balancer Group will be populated with the members and weights retrieved from the back end.",
							Computed:            true,
						},
						"wlm_transport": schema.StringAttribute{
							MarkdownDescription: "Specify either HTTP or HTTPS for the Load Balancer Group protocol. This protocol is used to forward traffic between the DataPower Gateway and the members of the Load Balancer Group.",
							Computed:            true,
						},
						"damp": schema.Int64Attribute{
							MarkdownDescription: "When a real server is observed to be non-functioning, it is temporarily disabled. When the damp time has elapsed, it is re-enabled. Allowable values are in the range 1 - 86400.",
							Computed:            true,
						},
						"never_return_sick_member": schema.BoolAttribute{
							MarkdownDescription: "During normal operation, when all members of the load-balancing group are down and a new request for that group is made, the first member of the group is automatically selected. If this property is turned on, no attempt will be made to connect under these circumstances.",
							Computed:            true,
						},
						"lb_group_members": schema.ListNestedAttribute{
							MarkdownDescription: "Members",
							NestedObject:        models.GetDmLBGroupMemberDataSourceSchema(),
							Computed:            true,
						},
						"try_every_server_before_failing": schema.BoolAttribute{
							MarkdownDescription: "This property applies only when none of the group members are in the \"up\" state. If this value is set, every server in the group is tried before failing the connection attempt. It is a \"last best-effort\" attempt.",
							Computed:            true,
						},
						"lb_group_checks": models.GetDmLBGroupCheckDataSourceSchema("The members of a Load Balancer Group can be periodically polled to verify the health of the server. If the server is found to be unresponsive, it is removed from the list of actively available servers until the unresponsive server passes a health check.", "health-check", ""),
						"masquerade_member": schema.BoolAttribute{
							MarkdownDescription: "If you set this value, the host name presented to the server will be the name of the group instead of the name of the member being used for that specific transaction.",
							Computed:            true,
						},
						"application_routing": schema.BoolAttribute{
							MarkdownDescription: "<p>If set to on, the load balancer group will route to the back end cluster depending on the following conditions.</p><ul><li>the application for which this request is targeted</li><li>the application status on the back end servers</li></ul><p>Application Routing is required for Application Edition (group or atomic) rollout. If you need Application Edition support, set the Update Type to Subscribe in the WebSphere Cell object.</p>",
							Computed:            true,
						},
						"lb_group_affinity_conf": models.GetDmLBGroupAffinityDataSourceSchema("Session affinity allows applications to maintain sessions with clients.", "session-affinity", ""),
						"monitored_cookies": schema.ListAttribute{
							MarkdownDescription: "The DataPower Gateway enforces session affinity when the application server attempts to establish session affinity using one of these cookie names.",
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

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *LoadBalancerGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data LoadBalancerGroupList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.LoadBalancerGroup{
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
	l := []models.LoadBalancerGroup{}
	if resFound {
		if value := res.Get(`LoadBalancerGroup`); value.Exists() {
			for _, v := range value.Array() {
				item := models.LoadBalancerGroup{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
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
