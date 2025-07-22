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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &TFIMEndpointResource{}

func NewTFIMEndpointResource() resource.Resource {
	return &TFIMEndpointResource{}
}

type TFIMEndpointResource struct {
	client *client.DatapowerClient
}

func (r *TFIMEndpointResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tfimendpoint"
}

func (r *TFIMEndpointResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Tivoli Federated Identity Manager (deprecated)", "tfim", "").String,

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
			"m_endpoint_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Endpoint type", "tfim-endpoint-type", "").AddStringEnum("tokenmapping", "oauth").AddDefaultValue("tokenmapping").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("tokenmapping", "oauth"),
				},
				Default: stringdefault.StaticString("tokenmapping"),
			},
			"m_server_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Server", "tfim-addr", "").String,
				Required:            true,
			},
			"m_server_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Port", "tfim-port", "").AddDefaultValue("9080").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(9080),
			},
			"m_compatible_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Compatibility mode", "tfim-compatible", "").AddStringEnum("v6.0", "v6.1", "v6.2").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("v6.0", "v6.1", "v6.2"),
				},
			},
			"m_req_token60_format": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Request token format", "tfim-60-req-tokenformat", "").AddStringEnum("WSUserNameToken", "SAML1.0", "SAML1.1", "Custom").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("WSUserNameToken", "SAML1.0", "SAML1.1", "Custom"),
				},
			},
			"m_req_token61_format": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Request token format", "tfim-61-req-tokenformat", "").AddStringEnum("WSUserNameToken", "SAML1.0", "SAML1.1", "SAML2.0", "WSKerberosToken", "WSX509Token", "BinarySecurityToken", "CustomToken", "Custom").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("WSUserNameToken", "SAML1.0", "SAML1.1", "SAML2.0", "WSKerberosToken", "WSX509Token", "BinarySecurityToken", "CustomToken", "Custom"),
				},
			},
			"m_req_token62_format": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Request Token Format", "tfim-62-req-tokenformat", "").AddStringEnum("WSUserNameToken", "SAML1.0", "SAML1.1", "SAML2.0", "WSKerberosToken", "WSX509Token", "BinarySecurityToken", "CustomToken", "Custom").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("WSUserNameToken", "SAML1.0", "SAML1.1", "SAML2.0", "WSKerberosToken", "WSX509Token", "BinarySecurityToken", "CustomToken", "Custom"),
				},
			},
			"m_req_custom_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Custom request", "tfim-custom-req-url", "").String,
				Optional:            true,
			},
			"m_applies_to_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("AppliesTo address", "tfim-pathaddr", "").String,
				Optional:            true,
			},
			"m_issuer": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Issuer", "tfim-issuer", "").String,
				Optional:            true,
			},
			"m_port_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Port Type", "tfim-porttype", "").String,
				Optional:            true,
			},
			"m_operation": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Operation", "tfim-operation", "").String,
				Optional:            true,
			},
			"m_schema_validate": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Schema-validate response", "tfim-schema-validate", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"m_ltpa_value_type_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LTPA BinarySecurityToken ValueType", "tfim-lpta-valuetype-mode", "").AddStringEnum("static", "dynamic").AddDefaultValue("static").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("static", "dynamic"),
				},
				Default: stringdefault.StaticString("static"),
			},
			"m_sts_username": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Username", "tfim-username", "").String,
				Optional:            true,
			},
			"m_sts_password": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Password (deprecated)", "tfim-password", "").String,
				Optional:            true,
			},
			"m_sts_password_alias": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Password alias", "tfim-password-alias", "passwordalias").String,
				Optional:            true,
			},
			"m_ssl_client_config_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client type", "ssl-client-type", "").AddStringEnum("proxy", "client").AddDefaultValue("client").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("proxy", "client"),
				},
				Default: stringdefault.StaticString("client"),
			},
			"m_ssl_client": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "sslclientprofile").String,
				Optional:            true,
			},
		},
	}
}

func (r *TFIMEndpointResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *TFIMEndpointResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.TFIMEndpoint

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `TFIMEndpoint`)
	res, err := r.client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}

	res, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TFIMEndpointResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.TFIMEndpoint

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && (strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	if data.IsNull() {
		// Import
		data.FromBody(ctx, `TFIMEndpoint`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `TFIMEndpoint`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TFIMEndpointResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.TFIMEndpoint

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `TFIMEndpoint`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s, %s", err, res.String()))
		return
	}
	res, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TFIMEndpointResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.TFIMEndpoint

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	resp.State.RemoveResource(ctx)
}
