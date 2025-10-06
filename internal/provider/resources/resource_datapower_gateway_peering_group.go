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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
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

var _ resource.Resource = &GatewayPeeringGroupResource{}
var _ resource.ResourceWithImportState = &GatewayPeeringGroupResource{}

func NewGatewayPeeringGroupResource() resource.Resource {
	return &GatewayPeeringGroupResource{}
}

type GatewayPeeringGroupResource struct {
	pData *tfutils.ProviderData
}

func (r *GatewayPeeringGroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_peering_group"
}

func (r *GatewayPeeringGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("A gateway-peering group defines members as a group to synchronize data across members. When a group can work in stand-alone mode, peer-based mode, or cluster-based mode. <ul><li>For stand-alone, no peers defined. This mode is for only development or testing purposes.</li><li>For a peer group, add peers and configure the connection among the peers.</li><li>For a cluster, add cluster nodes and the other nodes that are in the same data center.</li></ul>", "gateway-peering-group", "").String,
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
			"mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Mode", "mode", "").AddStringEnum("peer", "cluster").AddDefaultValue("peer").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("peer", "cluster"),
				},
				Default: stringdefault.StaticString("peer"),
			},
			"peer_nodes": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify peers for the the group. To add a peer, enter its local IP address or host alias and its priority.", "peer-node", "").AddRequiredWhen(models.GatewayPeeringGroupPeerNodesCondVal.String()).AddNotValidWhen(models.GatewayPeeringGroupPeerNodesIgnoreVal.String()).String,
				NestedObject:        models.GetDmGatewayPeeringGroupPeerNodeResourceSchema(),
				Optional:            true,
				Validators: []validator.List{
					validators.ConditionalRequiredList(models.GatewayPeeringGroupPeerNodesCondVal, models.GatewayPeeringGroupPeerNodesIgnoreVal, false),
				},
			},
			"cluster_primary_count": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Primary count", "cluster-primary-count", "").AddStringEnum("3").AddDefaultValue("3").AddRequiredWhen(models.GatewayPeeringGroupClusterPrimaryCountCondVal.String()).AddNotValidWhen(models.GatewayPeeringGroupClusterPrimaryCountIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("3"),
					validators.ConditionalRequiredString(models.GatewayPeeringGroupClusterPrimaryCountCondVal, models.GatewayPeeringGroupClusterPrimaryCountIgnoreVal, true),
				},
				Default: stringdefault.StaticString("3"),
			},
			"cluster_nodes": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify nodes for the cluster group. To add a node, enter its local IP address or host alias and the comma-separated list of local IP addresses or host aliases of the other nodes that are in the same data center. <p>Because the primary count is 3, the configuration requires a minimum of 6 nodes that are generally in 2 data centers. Each node is defined on a different DataPower Gateway. The minimal configuration is 3 primary-secondary pairs. Each pair is a shard that manages a subset of slots.</p><p>Each primary node can have more than one secondary node, but each primary node requires the same number of secondary nodes. In other words, you can define an environment of 9 nodes, which is a configuration of 3 primary nodes and 6 secondary nodes. In this configuration, each primary node has 2 secondary nodes.</p>", "cluster-node", "").AddRequiredWhen(models.GatewayPeeringGroupClusterNodesCondVal.String()).AddNotValidWhen(models.GatewayPeeringGroupClusterNodesIgnoreVal.String()).String,
				NestedObject:        models.GetDmGatewayPeeringGroupClusterNodeResourceSchema(),
				Optional:            true,
				Validators: []validator.List{
					validators.ConditionalRequiredList(models.GatewayPeeringGroupClusterNodesCondVal, models.GatewayPeeringGroupClusterNodesIgnoreVal, false),
				},
			},
			"cluster_auto_config": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the cluster configuration is managed automatically. By default, cluster configuration is managed automatically. Unless directed by IBM Support, do not change this setting.", "cluster-auto-config", "").AddDefaultValue("true").AddRequiredWhen(models.GatewayPeeringGroupClusterAutoConfigCondVal.String()).AddNotValidWhen(models.GatewayPeeringGroupClusterAutoConfigIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"enable_ssl": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to use TLS to secure the connection among the members. By default, TLS is enabled. When enabled, ensure that all members use the same TLS configuration.", "enable-ssl", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"idcred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the identification credentials that contains the credentials that the current member uses to identify itself to other peers. Client authentication uses mutual TLS.", "idcred", "crypto_ident_cred").AddRequiredWhen(models.GatewayPeeringGroupIdcredCondVal.String()).AddNotValidWhen(models.GatewayPeeringGroupIdcredIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.GatewayPeeringGroupIdcredCondVal, models.GatewayPeeringGroupIdcredIgnoreVal, false),
				},
			},
			"valcred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validation credentials", "valcred", "crypto_val_cred").AddNotValidWhen(models.GatewayPeeringGroupValcredIgnoreVal.String()).String,
				Optional:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *GatewayPeeringGroupResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *GatewayPeeringGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.GatewayPeeringGroup
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `GatewayPeeringGroup`)
	_, err := r.pData.Client.Post(data.GetPath(), body)
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), body)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Resource already exists. Failed to update resource, got error: %s", err))
				return
			}
			resp.Diagnostics.AddWarning("Warning", "Resource already exists. Existing resource was updated.")
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

func (r *GatewayPeeringGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.GatewayPeeringGroup
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && (strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}

	data.UpdateFromBody(ctx, `GatewayPeeringGroup`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPeeringGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.GatewayPeeringGroup
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `GatewayPeeringGroup`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Update)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPeeringGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.GatewayPeeringGroup
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

func (r *GatewayPeeringGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

	var data models.GatewayPeeringGroup
	data.AppDomain = types.StringValue(appDomain)
	data.Id = types.StringValue(id)
	res, err := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("Resource was not found, got error: %s", err))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		}
		return
	}

	data.FromBody(ctx, `GatewayPeeringGroup`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPeeringGroupResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.GatewayPeeringGroup

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
