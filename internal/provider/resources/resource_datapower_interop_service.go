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
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

var _ resource.Resource = &InteropServiceResource{}

func NewInteropServiceResource() resource.Resource {
	return &InteropServiceResource{}
}

type InteropServiceResource struct {
	pData *tfutils.ProviderData
}

func (r *InteropServiceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interop_service"
}

func (r *InteropServiceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Configure and use the interoperability test service to test stylesheets or maps. The service is available over HTTP or HTTPS and is disabled by default.", "iop-mgmt", "").String,
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>The administrative state of the configuration.</p><ul><li>To make active, set to enabled.</li><li>To make inactive, set to disabled.</li></ul>", "admin-state", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"xml_manager": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML manager", "xml-manager", "xml_manager").String,
				Optional:            true,
			},
			"aaa_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("AAA policy", "aaa-policy", "aaa_policy").String,
				Optional:            true,
			},
			"http_service": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable over HTTP", "http-service", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"local_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the IP address or host alias that the service listens. The default value is 0.0.0.0, which indicates that the service is active on all addresses.", "http-ip-address", "").AddDefaultValue("0.0.0.0").AddRequiredWhen(models.InteropServiceLocalAddressCondVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.InteropServiceLocalAddressCondVal, validators.Evaluation{}, true),
				},
				Default: stringdefault.StaticString("0.0.0.0"),
			},
			"local_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Local port", "http-port", "").AddIntegerRange(1000, 61000).AddDefaultValue("9990").AddRequiredWhen(models.InteropServiceLocalPortCondVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1000, 61000),
					validators.ConditionalRequiredInt64(models.InteropServiceLocalPortCondVal, validators.Evaluation{}, true),
				},
				Default: int64default.StaticInt64(9990),
			},
			"acl": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Access control list", "http-acl", "access_control_list").String,
				Optional:            true,
			},
			"https_service": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable over HTTPS", "https-service", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"https_local_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the IP address or host alias that the service listens. The default value is 0.0.0.0, which indicates that the service is active on all addresses.", "https-ip-address", "").AddDefaultValue("0.0.0.0").AddRequiredWhen(models.InteropServiceHttpsLocalAddressCondVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.InteropServiceHttpsLocalAddressCondVal, validators.Evaluation{}, true),
				},
				Default: stringdefault.StaticString("0.0.0.0"),
			},
			"https_local_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Local port", "https-port", "").AddIntegerRange(1000, 61000).AddDefaultValue("9991").AddRequiredWhen(models.InteropServiceHttpsLocalPortCondVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1000, 61000),
					validators.ConditionalRequiredInt64(models.InteropServiceHttpsLocalPortCondVal, validators.Evaluation{}, true),
				},
				Default: int64default.StaticInt64(9991),
			},
			"https_acl": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Access control list", "https-acl", "access_control_list").String,
				Optional:            true,
			},
			"ssl_server_config_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS server type", "ssl-config-type", "").AddStringEnum("server", "sni").AddDefaultValue("server").AddRequiredWhen(models.InteropServiceSSLServerConfigTypeCondVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("server", "sni"),
					validators.ConditionalRequiredString(models.InteropServiceSSLServerConfigTypeCondVal, validators.Evaluation{}, true),
				},
				Default: stringdefault.StaticString("server"),
			},
			"ssl_server": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS server profile", "ssl-server", "ssl_server_profile").AddRequiredWhen(models.InteropServiceSSLServerCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.InteropServiceSSLServerCondVal, validators.Evaluation{}, false),
				},
			},
			"sslsni_server": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS SNI server profile", "ssl-sni-server", "ssl_sni_server_profile").AddRequiredWhen(models.InteropServiceSSLSNIServerCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.InteropServiceSSLSNIServerCondVal, validators.Evaluation{}, false),
				},
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *InteropServiceResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *InteropServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.InteropService
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, "default", data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `InteropService`)
	_, err := r.pData.Client.Put(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "PUT", err))
		return
	}
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *InteropServiceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.InteropService
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

	if data.IsNull() {
		// Import
		data.FromBody(ctx, `InteropService`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `InteropService`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *InteropServiceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.InteropService
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, "default", data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath(), data.ToBody(ctx, `InteropService`))
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

func (r *InteropServiceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.InteropService
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, "default", data.DependencyActions, actions.Delete, false)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *InteropServiceResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.InteropService

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
