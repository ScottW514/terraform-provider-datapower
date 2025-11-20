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

package resources

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

var _ resource.Resource = &LoadBalancerGroupResource{}
var _ resource.ResourceWithImportState = &LoadBalancerGroupResource{}

func NewLoadBalancerGroupResource() resource.Resource {
	return &LoadBalancerGroupResource{}
}

type LoadBalancerGroupResource struct {
	pData *tfutils.ProviderData
}

func (r *LoadBalancerGroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_load_balancer_group"
}

func (r *LoadBalancerGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("<p>The DataPower device distributes traffic to members of a Load Balancer Group. These are back end servers and not additional DataPower devices. A Load Balancer Group lists members of a virtual server group and sets the algorithm for balancing them. Periodic health checks can be performed. Load Balancers may also be used to provide redundant LDAP server access.</p><p>When created, a DataPower service can use a Load Balancer Group by associating it with an XML manager that is associated with this service.</p><p>The back end destination URL is set to the name of the Load Balancer Group (example: \"BackEndServers\").</p>", "loadbalancer-group", "").String,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name of the object. Must be unique among object types in application domain.", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"app_domain": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The name of the application domain the object belongs to", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
				PlanModifiers: []planmodifier.String{
					modifiers.ImmutableAfterSet(),
				},
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"algorithm": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the algorithm to use to balance the real servers.", "algorithm", "").AddStringEnum("round-robin", "weighted-round-robin", "hash", "least-connections", "first-alive", "weighted-least-connections").AddDefaultValue("round-robin").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("round-robin", "weighted-round-robin", "hash", "least-connections", "first-alive", "weighted-least-connections"),
				},
				Default: stringdefault.StaticString("round-robin"),
			},
			"retrieve_info": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use this setting to control whether this Load Balancer Group has membership and weight information automatically retrieved from the work load management repository WebSphere Cell. When disabled, the static configuration is used.", "retrieve-wlm-info", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wlm_retrieval": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Contains the back end work load management repository selection type. Select 'WebSphere Cell' if your back-end is a WebSphere Application Server (WAS) Network Deployment (ND) or WAS Virtual Enterprise (VE).", "wlm-type", "").AddStringEnum("use-websphere").AddDefaultValue("use-websphere").AddNotValidWhen(models.LoadBalancerGroupWLMRetrievalIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("use-websphere"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.LoadBalancerGroupWLMRetrievalIgnoreVal, true),
				},
				Default: stringdefault.StaticString("use-websphere"),
			},
			"web_sphere_cell_config": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("If you selected 'WebSphere Cell' for Workload Management Retrieval, you need to select a WebSphere Cell object that retrieves this information. If no objects are available in the pull down, you must create one.", "websphere-cell", "wcc_service").AddNotValidWhen(models.LoadBalancerGroupWebSphereCellConfigIgnoreVal.String()).String,
				Optional:            true,
			},
			"wlm_group": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The Workload Management Group Name is used to define a group. In a WebSphere Application Server environment, the back end group is a cluster name. Once specified, the Load Balancer Group will be populated with the members and weights retrieved from the back end.", "wlm-group", "").AddNotValidWhen(models.LoadBalancerGroupWLMGroupIgnoreVal.String()).String,
				Optional:            true,
			},
			"wlm_transport": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify either HTTP or HTTPS for the Load Balancer Group protocol. This protocol is used to forward traffic between the DataPower Gateway and the members of the Load Balancer Group.", "wlm-transport", "").AddStringEnum("http", "https").AddDefaultValue("http").AddNotValidWhen(models.LoadBalancerGroupWLMTransportIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("http", "https"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.LoadBalancerGroupWLMTransportIgnoreVal, true),
				},
				Default: stringdefault.StaticString("http"),
			},
			"damp": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("When a real server is observed to be non-functioning, it is temporarily disabled. When the damp time has elapsed, it is re-enabled. Allowable values are in the range 1 - 86400.", "damp", "").AddIntegerRange(1, 86400).AddDefaultValue("120").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(120),
			},
			"never_return_sick_member": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("During normal operation, when all members of the load-balancing group are down and a new request for that group is made, the first member of the group is automatically selected. If this property is turned on, no attempt will be made to connect under these circumstances.", "giveup-when-all-members-down", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"lb_group_members": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Members", "server", "").String,
				NestedObject:        models.GetDmLBGroupMemberResourceSchema(),
				Optional:            true,
			},
			"try_every_server_before_failing": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("This property applies only when none of the group members are in the \"up\" state. If this value is set, every server in the group is tried before failing the connection attempt. It is a \"last best-effort\" attempt.", "try-every-server", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"lb_group_checks": models.GetDmLBGroupCheckResourceSchema("The members of a Load Balancer Group can be periodically polled to verify the health of the server. If the server is found to be unresponsive, it is removed from the list of actively available servers until the unresponsive server passes a health check.", "health-check", "", false),
			"masquerade_member": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("If you set this value, the host name presented to the server will be the name of the group instead of the name of the member being used for that specific transaction.", "masquerade", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"application_routing": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>If set to on, the load balancer group will route to the back end cluster depending on the following conditions.</p><ul><li>the application for which this request is targeted</li><li>the application status on the back end servers</li></ul><p>Application Routing is required for Application Edition (group or atomic) rollout. If you need Application Edition support, set the Update Type to Subscribe in the WebSphere Cell object.</p>", "appl-routing", "").AddDefaultValue("false").AddNotValidWhen(models.LoadBalancerGroupApplicationRoutingIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"lb_group_affinity_conf": models.GetDmLBGroupAffinityResourceSchema("Session affinity allows applications to maintain sessions with clients.", "session-affinity", "", false),
			"monitored_cookies": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The DataPower Gateway enforces session affinity when the application server attempts to establish session affinity using one of these cookie names.", "monitored-cookie", "").AddNotValidWhen(models.LoadBalancerGroupMonitoredCookiesIgnoreVal.String()).String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *LoadBalancerGroupResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *LoadBalancerGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.LoadBalancerGroup
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `LoadBalancerGroup`)
	_, err := r.pData.Client.Post(data.GetPath(), body)
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), body)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Resource already exists. Failed to update resource, got error: %s", err))
				return
			}
			resp.Diagnostics.AddWarning("Warning", "Resource already exists. Existing resource was updated.")
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString())
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create resource, got error: %s", err))
			return
		}
	}
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LoadBalancerGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.LoadBalancerGroup
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil {
		if strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400") {
			resp.State.RemoveResource(ctx)
			return
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString())
			if !resp.Diagnostics.HasError() {
				resp.State.RemoveResource(ctx)
			}
			return
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
			return
		}
	}

	data.UpdateFromBody(ctx, `LoadBalancerGroup`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LoadBalancerGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.LoadBalancerGroup
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `LoadBalancerGroup`))
	if err != nil {
		if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString())
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
			return
		}
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Update)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LoadBalancerGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.LoadBalancerGroup
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			resp.Diagnostics.AddWarning("Resource Conflict", fmt.Sprintf("Resource is no longer tracked by Terraform, but may need to be manually deleted on DataPower host. Got error: %s", err))
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString())
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete resource, got error: %s", err))
			return
		}
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *LoadBalancerGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError("Invalid Import ID", "Expected format: <app_domain>/<id>. Got: "+req.ID)
		return
	}

	appDomain := parts[0]
	id := parts[1]

	if !regexp.MustCompile("^[a-zA-Z0-9_-]+$").MatchString(id) || len(id) < 1 || len(id) > 128 {
		resp.Diagnostics.AddError("Invalid ID", "ID must be 1-128 characters and match pattern ^[a-zA-Z0-9_-]+$")
		return
	}
	if !regexp.MustCompile("^[a-zA-Z0-9_-]+$").MatchString(appDomain) || len(appDomain) < 1 || len(appDomain) > 128 {
		resp.Diagnostics.AddError("Invalid Application Domain", "Application domain must be 1-128 characters and match pattern ^[a-zA-Z0-9_-]+$")
		return
	}

	var data models.LoadBalancerGroup
	data.AppDomain = types.StringValue(appDomain)
	data.Id = types.StringValue(id)
	res, err := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("Resource was not found, got error: %s", err))
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString())
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		}
		return
	}

	data.FromBody(ctx, `LoadBalancerGroup`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LoadBalancerGroupResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.LoadBalancerGroup

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
