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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
)

type TFIMEndpointList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &TFIMEndpointDataSource{}
	_ datasource.DataSourceWithConfigure = &TFIMEndpointDataSource{}
)

func NewTFIMEndpointDataSource() datasource.DataSource {
	return &TFIMEndpointDataSource{}
}

type TFIMEndpointDataSource struct {
	client *client.DatapowerClient
}

func (d *TFIMEndpointDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tfimendpoint"
}

func (d *TFIMEndpointDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Tivoli Federated Identity Manager (deprecated)",
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
						"m_endpoint_type": schema.StringAttribute{
							MarkdownDescription: "Endpoint type",
							Computed:            true,
						},
						"m_server_url": schema.StringAttribute{
							MarkdownDescription: "Server",
							Computed:            true,
						},
						"m_server_port": schema.Int64Attribute{
							MarkdownDescription: "Port",
							Computed:            true,
						},
						"m_compatible_mode": schema.StringAttribute{
							MarkdownDescription: "Compatibility mode",
							Computed:            true,
						},
						"m_req_token60_format": schema.StringAttribute{
							MarkdownDescription: "Request token format",
							Computed:            true,
						},
						"m_req_token61_format": schema.StringAttribute{
							MarkdownDescription: "Request token format",
							Computed:            true,
						},
						"m_req_token62_format": schema.StringAttribute{
							MarkdownDescription: "Request Token Format",
							Computed:            true,
						},
						"m_req_custom_url": schema.StringAttribute{
							MarkdownDescription: "Custom request",
							Computed:            true,
						},
						"m_applies_to_address": schema.StringAttribute{
							MarkdownDescription: "AppliesTo address",
							Computed:            true,
						},
						"m_issuer": schema.StringAttribute{
							MarkdownDescription: "Issuer",
							Computed:            true,
						},
						"m_port_type": schema.StringAttribute{
							MarkdownDescription: "Port Type",
							Computed:            true,
						},
						"m_operation": schema.StringAttribute{
							MarkdownDescription: "Operation",
							Computed:            true,
						},
						"m_schema_validate": schema.BoolAttribute{
							MarkdownDescription: "Schema-validate response",
							Computed:            true,
						},
						"m_ltpa_value_type_mode": schema.StringAttribute{
							MarkdownDescription: "LTPA BinarySecurityToken ValueType",
							Computed:            true,
						},
						"m_sts_username": schema.StringAttribute{
							MarkdownDescription: "Username",
							Computed:            true,
						},
						"m_sts_password_alias": schema.StringAttribute{
							MarkdownDescription: "Password alias",
							Computed:            true,
						},
						"m_ssl_client_config_type": schema.StringAttribute{
							MarkdownDescription: "TLS client type",
							Computed:            true,
						},
						"m_ssl_client": schema.StringAttribute{
							MarkdownDescription: "TLS client profile",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *TFIMEndpointDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *TFIMEndpointDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data TFIMEndpointList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.TFIMEndpoint{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.TFIMEndpoint{}
	if value := res.Get(`TFIMEndpoint`); value.Exists() {
		for _, v := range value.Array() {
			item := models.TFIMEndpoint{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.TFIMEndpointObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.TFIMEndpointObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
