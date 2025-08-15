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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &GatewayPeeringResource{}

func NewGatewayPeeringResource() resource.Resource {
	return &GatewayPeeringResource{}
}

type GatewayPeeringResource struct {
	pData *tfutils.ProviderData
}

func (r *GatewayPeeringResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_peering"
}

func (r *GatewayPeeringResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("A gateway-peering instance defines how members synchronize data across members.", "gateway-peering", "").String,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name of the object. Must be unique among object types in application domain.", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`), ""),
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
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`), ""),
				},
				PlanModifiers: []planmodifier.String{
					modifiers.ImmutableAfterSet(),
				},
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"password_alias": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the password alias to secure the data store. If not specified, a system default is used. The use of the system default is classified as a security vulnerability (CVE-2022-31776).", "password-alias", "password_alias").String,
				Optional:            true,
			},
			"local_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the IP address or host alias that the gateway service listens on. The IP address can be any DataPower network interface that can be accessed by other peers in the peer group. The IP address cannot be 127.0.0.1, 0.0.0.0 or ::.", "local-address", "").String,
				Optional:            true,
			},
			"local_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the port that the gateway service listens on. The default value is 16380. Ensure that all peers use the same port.", "local-port", "").AddDefaultValue("16380").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(16380),
			},
			"peer_group": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Gateway-peering group", "peer-group", "gateway_peering_group").String,
				Optional:            true,
			},
			"monitor_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the port to monitor for state synchronization. The default value is 26380. Ensure that all peers use the same port.", "monitor-port", "").AddDefaultValue("26380").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(26380),
			},
			"enable_peer_group": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the gateway-peering instance uses a peer group. <ul><li>When enabled, the instance works in the mode that is set for the peer group. This setting is the default value.</li><li>When not enabled, the instance works in stand-alone mode.</li></ul>", "enable-peer-group", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"enable_ssl": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to use TLS to secure the connection among the members. By default, TLS is enabled. In peer-based mode, ensure that all peers use the same TLS configuration.", "", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"persistence_location": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify where to store data. Ensure that all peers in the group store data in the same location.", "persistence", "").AddStringEnum("memory", "local", "raid").AddDefaultValue("memory").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("memory", "local", "raid"),
				},
				Default: stringdefault.StaticString("memory"),
			},
			"local_directory": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the directory to store data. For example, <tt>local:///group1</tt> .", "local-directory", "").AddDefaultValue("local:///").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("local:///"),
			},
			"max_memory": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum memory for the data store. When memory reaches this limit, data is removed by using the least recently used (LRU) algorithm. The default value is 0, which means no limits. Do not over allocate memory.", "max-memory", "").AddIntegerRange(0, 1048576).String,
				Optional:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 1048576),
				},
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *GatewayPeeringResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *GatewayPeeringResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.GatewayPeering
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `GatewayPeering`)
	_, err := r.pData.Client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "POST", err))
		return
	}
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPeeringResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.GatewayPeering
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

	if data.IsNull() {
		// Import
		data.FromBody(ctx, `GatewayPeering`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `GatewayPeering`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPeeringResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.GatewayPeering
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `GatewayPeering`))
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

func (r *GatewayPeeringResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.GatewayPeering
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *GatewayPeeringResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.GatewayPeering

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
