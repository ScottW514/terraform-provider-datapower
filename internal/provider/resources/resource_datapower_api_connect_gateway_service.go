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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

var _ resource.Resource = &APIConnectGatewayServiceResource{}
var _ resource.ResourceWithImportState = &APIConnectGatewayServiceResource{}

func NewAPIConnectGatewayServiceResource() resource.Resource {
	return &APIConnectGatewayServiceResource{}
}

type APIConnectGatewayServiceResource struct {
	pData *tfutils.ProviderData
}

func (r *APIConnectGatewayServiceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_connect_gateway_service"
}

func (r *APIConnectGatewayServiceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("The API Connect gateway service defines the type of gateway service and manages connections with API Connect. When configured, the DataPower Gateway creates a gateway service to retrieve data from API Connect to define the configuration to process API requests.", "apic-gw-service", "").String,
		Attributes: map[string]schema.Attribute{
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
			"enabled": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>The administrative state of the configuration.</p><ul><li>To make active, set to enabled.</li><li>To make inactive, set to disabled.</li></ul>", "admin-state", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"local_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the IP address or interface through that API Connect uses to manage the gateway service. The default value is 0.0.0.0.", "local-address", "").AddDefaultValue("0.0.0.0").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("0.0.0.0"),
			},
			"local_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the listening port for the gateway service. The default value is 3000. <p><b>Note:</b> The gateway service uses four additional consecutive ports after the local port. Therefore, all five consecutive ports must be clear of conflicts.</p>", "local-port", "").AddDefaultValue("3000").String,
				Computed:            true,
				Optional:            true,
				Default:             int64default.StaticInt64(3000),
			},
			"ssl_server": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS server profile to secure connections between API Connect to the gateway service. The following restrictions apply. <ul><li>Keys and certificates are restricted to PEM and PKCS #12 formats.</li><li>The validation credentials must use PEM formatted material.</li></ul>", "ssl-server", "ssl_server_profile").String,
				Optional:            true,
			},
			"api_gateway_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the IP address or host alias to accept API requests. The default value is 0.0.0.0. This address is used with its port to create an HTTPS handler.", "api-gw-address", "").AddDefaultValue("0.0.0.0").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("0.0.0.0"),
			},
			"api_gateway_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the listening port for API requests. The default value is 9443. This port is used with its address to create an HTTPS handler.", "api-gw-port", "").AddDefaultValue("9443").String,
				Computed:            true,
				Optional:            true,
				Default:             int64default.StaticInt64(9443),
			},
			"gateway_peering": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the gateway-peering instance that manages data across the gateway peers. The following restrictions apply. <ul><li>When TLS and peer group mode are enabled, all peers must use the same crypto material.</li><li>Keys and certificates are restricted to PEM and PKCS #12 formats.</li></ul>", "gateway-peering", "gateway_peering").AddRequiredWhen(models.APIConnectGatewayServiceGatewayPeeringCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.APIConnectGatewayServiceGatewayPeeringCondVal, validators.Evaluation{}, false),
				},
			},
			"gateway_peering_manager": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the gateway-peering manager that manages gateway-peering instances for the gateway service. This property is meaningful when the gateway type is an API gateway.", "gateway-peering-manager", "gateway_peering_manager").AddDefaultValue("default").AddRequiredWhen(models.APIConnectGatewayServiceGatewayPeeringManagerCondVal.String()).AddNotValidWhen(models.APIConnectGatewayServiceGatewayPeeringManagerIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.APIConnectGatewayServiceGatewayPeeringManagerCondVal, models.APIConnectGatewayServiceGatewayPeeringManagerIgnoreVal, true),
				},
				Default: stringdefault.StaticString("default"),
			},
			"v5compatibility_mode": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the gateway service is a Multi-Protocol Gateway or an API gateway. <ui><li>When enabled, the gateway service is a Multi-Protocol Gateway that is compatible with API Connect version 5.</li><li>When disabled, that gateway service is an API gateway this is not compatible with API Connect v5.</li></ui>", "v5-compatibility-mode", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"user_defined_policies": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify user-defined policies to advertise to API Connect for use in the API Connect Assembly Editor. This property is meaningful when the gateway type is an API gateway. <p>For an assembly function that is a user-defined policy, configure the assembly function with a mechanism other than a watched file that is processed by a configuration sequence. Objects that are created through the processing of configuration sequences are not persisted to the startup configuration. The preferred method for user-defined policies is to define them explicitly so that they persist to the startup configuration.</p>", "user-defined-policies", "assembly_function").AddNotValidWhen(models.APIConnectGatewayServiceUserDefinedPoliciesIgnoreVal.String()).String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"v5c_slm_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the peer group type for the SLM policy. This property is meaningful when the gateway type is a Multi-Protocol Gateway.", "slm-mode", "").AddStringEnum("autounicast", "unicast", "multicast").AddDefaultValue("autounicast").AddRequiredWhen(models.APIConnectGatewayServiceV5CSlmModeCondVal.String()).AddNotValidWhen(models.APIConnectGatewayServiceV5CSlmModeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("autounicast", "unicast", "multicast"),
					validators.ConditionalRequiredString(models.APIConnectGatewayServiceV5CSlmModeCondVal, models.APIConnectGatewayServiceV5CSlmModeIgnoreVal, true),
				},
				Default: stringdefault.StaticString("autounicast"),
			},
			"ip_multicast": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the IP multicast configuration for the SLM policy. This property is meaningful when the gateway type is a Multi-Protocol Gateway and the peer mode is multicast.", "ip-multicast", "ip_multicast").AddRequiredWhen(models.APIConnectGatewayServiceIPMulticastCondVal.String()).AddNotValidWhen(models.APIConnectGatewayServiceIPMulticastIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.APIConnectGatewayServiceIPMulticastCondVal, models.APIConnectGatewayServiceIPMulticastIgnoreVal, false),
				},
			},
			"ip_unicast": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the address of the unicast peer group for the SLM policy. This property is meaningful when the gateway type is a Multi-Protocol Gateway and the peer mode is unicast.", "ip-unicast", "").AddRequiredWhen(models.APIConnectGatewayServiceIPUnicastCondVal.String()).AddNotValidWhen(models.APIConnectGatewayServiceIPUnicastIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.APIConnectGatewayServiceIPUnicastCondVal, models.APIConnectGatewayServiceIPUnicastIgnoreVal, false),
				},
			},
			"jwt_validation_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the JWT validation mode. This property does not control whether a token is validated. This property controls whether transactions fail when validation fails.", "jwt-validate-mode", "").AddStringEnum("request", "require").AddDefaultValue("request").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("request", "require"),
				},
				Default: stringdefault.StaticString("request"),
			},
			"jwt_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("JWT URL", "jwt-url", "").String,
				Optional:            true,
			},
			"proxy_policy":       models.GetDmAPICGSProxyPolicyResourceSchema("API Manager proxy", "proxy", "", false),
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *APIConnectGatewayServiceResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *APIConnectGatewayServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APIConnectGatewayService
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `APIConnectGatewayService`)
	_, err := r.pData.Client.Put(data.GetPath(), body)
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			_, err := r.pData.Client.Put(data.GetPath(), body)
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

func (r *APIConnectGatewayServiceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APIConnectGatewayService
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.pData.Client.Get(data.GetPath())
	if err != nil && (strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}

	data.UpdateFromBody(ctx, `APIConnectGatewayService`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIConnectGatewayServiceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APIConnectGatewayService
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath(), data.ToBody(ctx, `APIConnectGatewayService`))
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

func (r *APIConnectGatewayServiceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APIConnectGatewayService
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false)
	if resp.Diagnostics.HasError() {
		return
	}
	data.ToDefault()
	_, err := r.pData.Client.Put(data.GetPath(), data.ToBody(ctx, `APIConnectGatewayService`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to restore singleton to default, got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *APIConnectGatewayServiceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()
	appDomain := req.ID
	if !regexp.MustCompile("^[a-zA-Z0-9_-]+$").MatchString(appDomain) || len(appDomain) < 1 || len(appDomain) > 128 {
		resp.Diagnostics.AddError("Invalid Application Domain", "Application domain must be 1-128 characters and match pattern ^[a-zA-Z0-9_-]+$")
		return
	}

	var data models.APIConnectGatewayService
	data.AppDomain = types.StringValue(appDomain)
	res, err := r.pData.Client.Get(data.GetPath())
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("Resource was not found, got error: %s", err))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		}
		return
	}

	data.FromBody(ctx, `APIConnectGatewayService`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIConnectGatewayServiceResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.APIConnectGatewayService

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
