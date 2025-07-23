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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &APIConnectGatewayServiceResource{}

func NewAPIConnectGatewayServiceResource() resource.Resource {
	return &APIConnectGatewayServiceResource{}
}

type APIConnectGatewayServiceResource struct {
	client *client.DatapowerClient
}

func (r *APIConnectGatewayServiceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apiconnectgatewayservice"
}

func (r *APIConnectGatewayServiceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("API Connect gateway service", "apic-gw-service", "").String,

		Attributes: map[string]schema.Attribute{
			"app_domain": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The name of the application domain the object belongs to", "", "").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					ImmutableAfterSet(),
				},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Administrative state", "admin-state", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"local_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Local address", "local-address", "").AddDefaultValue("0.0.0.0").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("0.0.0.0"),
			},
			"local_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Local port", "local-port", "").AddDefaultValue("3000").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(3000),
			},
			"ssl_server": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS server profile", "ssl-server", "sslserverprofile").String,
				Optional:            true,
			},
			"api_gateway_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("API gateway address", "api-gw-address", "").AddDefaultValue("0.0.0.0").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("0.0.0.0"),
			},
			"api_gateway_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("API gateway port", "api-gw-port", "").AddDefaultValue("9443").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(9443),
			},
			"gateway_peering": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Gateway peering", "gateway-peering", "gatewaypeering").String,
				Optional:            true,
			},
			"gateway_peering_manager": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Gateway-peering manager", "gateway-peering-manager", "gatewaypeeringmanager").AddDefaultValue("default").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default"),
			},
			"v5_compatibility_mode": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("V5 compatibility mode", "v5-compatibility-mode", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"user_defined_policies": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("User-defined policies", "user-defined-policies", "assemblyfunction").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"v5c_slm_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SLM peer mode", "slm-mode", "").AddStringEnum("autounicast", "unicast", "multicast").AddDefaultValue("autounicast").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("autounicast", "unicast", "multicast"),
				},
				Default: stringdefault.StaticString("autounicast"),
			},
			"ip_multicast": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("IP multicast", "ip-multicast", "ipmulticast").String,
				Optional:            true,
			},
			"ip_unicast": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("IP unicast", "ip-unicast", "").String,
				Optional:            true,
			},
			"jwt_validation_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("JWT validation mode", "jwt-validate-mode", "").AddStringEnum("request", "require").AddDefaultValue("request").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("request", "require"),
				},
				Default: stringdefault.StaticString("request"),
			},
			"jwturl": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("JWT URL", "jwt-url", "").String,
				Optional:            true,
			},
			"proxy_policy": models.GetDmAPICGSProxyPolicyResourceSchema("API Manager proxy", "proxy", "", false),
		},
	}
}

func (r *APIConnectGatewayServiceResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *APIConnectGatewayServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.APIConnectGatewayService

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `APIConnectGatewayService`)
	_, err := r.client.Put(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "PUT", err))
		return
	}

	_, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s", "POST", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIConnectGatewayServiceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.APIConnectGatewayService

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Get(data.GetPath())
	if err != nil && (strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}

	if data.IsNull() {
		// Import
		data.FromBody(ctx, `APIConnectGatewayService`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `APIConnectGatewayService`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIConnectGatewayServiceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.APIConnectGatewayService

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Put(data.GetPath(), data.ToBody(ctx, `APIConnectGatewayService`))
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

func (r *APIConnectGatewayServiceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	resp.State.RemoveResource(ctx)
}
