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
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &GatewayPeeringResource{}

func NewGatewayPeeringResource() resource.Resource {
	return &GatewayPeeringResource{}
}

type GatewayPeeringResource struct {
	client *client.DatapowerClient
}

func (r *GatewayPeeringResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gatewaypeering"
}

func (r *GatewayPeeringResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Gateway peering", "gateway-peering", "").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name of the object. Must be unique among object types in application domain.", "", "").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"app_domain": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The name of the application domain the object belongs to", "", "").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					ImmutableAfterSet(),
				},
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"password_alias": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Password alias", "password-alias", "passwordalias").String,
				Optional:            true,
			},
			"local_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Local address", "local-address", "").String,
				Optional:            true,
			},
			"local_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Local port", "local-port", "").AddDefaultValue("16380").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(16380),
			},
			"peer_group": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Gateway-peering group", "peer-group", "gatewaypeeringgroup").String,
				Optional:            true,
			},
			"primary_count": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Primary count (deprecated)", "primary-count", "").AddStringEnum("1", "3").AddDefaultValue("1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1", "3"),
				},
				Default: stringdefault.StaticString("1"),
			},
			"monitor_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Monitor port", "monitor-port", "").AddDefaultValue("26380").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(26380),
			},
			"cluster_auto_config": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Auto manage cluster configuration (deprecated)", "cluster-auto-config", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"enable_peer_group": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use peer group", "enable-peer-group", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"peers": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Peers (deprecated)", "peer", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"cluster_nodes": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Cluster nodes (deprecated)", "cluster-node", "").String,
				NestedObject:        models.DmGatewayPeeringClusterNodeResourceSchema,
				Optional:            true,
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Priority (deprecated)", "priority", "").AddIntegerRange(0, 255).AddDefaultValue("100").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 255),
				},
				Default: int64default.StaticInt64(100),
			},
			"enable_ssl": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable TLS (deprecated)", "enable-ssl", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"idcred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Identification credentials (deprecated)", "idcred", "cryptoidentcred").String,
				Optional:            true,
			},
			"valcred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validation credentials (deprecated)", "valcred", "cryptovalcred").String,
				Optional:            true,
			},
			"ssl_crypto_key": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Key (deprecated)", "ssl-key", "cryptokey").String,
				Optional:            true,
			},
			"ssl_crypto_certificate": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Certificate (deprecated)", "ssl-cert", "cryptocertificate").String,
				Optional:            true,
			},
			"persistence_location": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Persistence location", "persistence", "").AddStringEnum("memory", "local", "raid").AddDefaultValue("memory").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("memory", "local", "raid"),
				},
				Default: stringdefault.StaticString("memory"),
			},
			"local_directory": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Local directory", "local-directory", "").AddDefaultValue("local:///").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("local:///"),
			},
			"max_memory": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Max memory", "max-memory", "").AddIntegerRange(0, 1048576).String,
				Optional:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 1048576),
				},
			},
		},
	}
}

func (r *GatewayPeeringResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *GatewayPeeringResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.GatewayPeering

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `GatewayPeering`)
	_, err := r.client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "POST", err))
		return
	}

	_, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s", "POST", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPeeringResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.GatewayPeering

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Get(data.GetPath() + "/" + data.Id.ValueString())
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
	var data models.GatewayPeering

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `GatewayPeering`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}
	_, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s", "POST", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPeeringResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.GatewayPeering

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}

	resp.State.RemoveResource(ctx)
}
