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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type WebSphereJMSServerList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &WebSphereJMSServerDataSource{}
	_ datasource.DataSourceWithConfigure = &WebSphereJMSServerDataSource{}
)

func NewWebSphereJMSServerDataSource() datasource.DataSource {
	return &WebSphereJMSServerDataSource{}
}

type WebSphereJMSServerDataSource struct {
	pData *tfutils.ProviderData
}

func (d *WebSphereJMSServerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_webspherejmsserver"
}

func (d *WebSphereJMSServerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "WebSphere JMS Server",
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
						"endpoint": schema.ListNestedAttribute{
							MarkdownDescription: "Endpoints",
							NestedObject:        models.DmWebSphereJMSEndpointDataSourceSchema,
							Computed:            true,
						},
						"target_transport_chain": schema.StringAttribute{
							MarkdownDescription: "Target transport chain",
							Computed:            true,
						},
						"messaging_bus": schema.StringAttribute{
							MarkdownDescription: "Messaging bus",
							Computed:            true,
						},
						"ssl_cipher": schema.StringAttribute{
							MarkdownDescription: "TLS ciphers",
							Computed:            true,
						},
						"fips": schema.BoolAttribute{
							MarkdownDescription: "FIPS compliant ciphers suite",
							Computed:            true,
						},
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"user_name": schema.StringAttribute{
							MarkdownDescription: "Username",
							Computed:            true,
						},
						"password_alias": schema.StringAttribute{
							MarkdownDescription: "Password alias",
							Computed:            true,
						},
						"transactional": schema.BoolAttribute{
							MarkdownDescription: "Transactional",
							Computed:            true,
						},
						"memory_threshold": schema.Int64Attribute{
							MarkdownDescription: "Memory threshold",
							Computed:            true,
						},
						"maximum_message_size": schema.Int64Attribute{
							MarkdownDescription: "Max message size",
							Computed:            true,
						},
						"default_message_type": schema.StringAttribute{
							MarkdownDescription: "Default message type",
							Computed:            true,
						},
						"total_connection_limit": schema.Int64Attribute{
							MarkdownDescription: "Total connection limit",
							Computed:            true,
						},
						"sessions_per_connection": schema.Int64Attribute{
							MarkdownDescription: "Max number of sessions per connection",
							Computed:            true,
						},
						"auto_retry": schema.BoolAttribute{
							MarkdownDescription: "Automatic retry",
							Computed:            true,
						},
						"retry_interval": schema.Int64Attribute{
							MarkdownDescription: "Retry interval",
							Computed:            true,
						},
						"enable_logging": schema.BoolAttribute{
							MarkdownDescription: "Enable JMS-specific logging",
							Computed:            true,
						},
						"ssl_client_config_type": schema.StringAttribute{
							MarkdownDescription: "TLS client type",
							Computed:            true,
						},
						"ssl_client": schema.StringAttribute{
							MarkdownDescription: "TLS client profile",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *WebSphereJMSServerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *WebSphereJMSServerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data WebSphereJMSServerList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.WebSphereJMSServer{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.WebSphereJMSServer{}
	if value := res.Get(`WebSphereJMSServer`); value.Exists() {
		for _, v := range value.Array() {
			item := models.WebSphereJMSServer{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.WebSphereJMSServerObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.WebSphereJMSServerObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
