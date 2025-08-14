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

type AnalyticsEndpointList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &AnalyticsEndpointDataSource{}
	_ datasource.DataSourceWithConfigure = &AnalyticsEndpointDataSource{}
)

func NewAnalyticsEndpointDataSource() datasource.DataSource {
	return &AnalyticsEndpointDataSource{}
}

type AnalyticsEndpointDataSource struct {
	pData *tfutils.ProviderData
}

func (d *AnalyticsEndpointDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_analyticsendpoint"
}

func (d *AnalyticsEndpointDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Analytics endpoint",
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
						"analytics_server_url": schema.StringAttribute{
							MarkdownDescription: "Analytics server URL",
							Computed:            true,
						},
						"ssl_client": schema.StringAttribute{
							MarkdownDescription: "TLS client profile",
							Computed:            true,
						},
						"request_topic": schema.StringAttribute{
							MarkdownDescription: "Request topic",
							Computed:            true,
						},
						"max_records": schema.Int64Attribute{
							MarkdownDescription: "Maximum records",
							Computed:            true,
						},
						"max_records_memory_kb": schema.Int64Attribute{
							MarkdownDescription: "Maximum size for each record",
							Computed:            true,
						},
						"max_delivery_memory_mb": schema.Int64Attribute{
							MarkdownDescription: "Maximum size for each delivery",
							Computed:            true,
						},
						"interval": schema.Int64Attribute{
							MarkdownDescription: "Interval",
							Computed:            true,
						},
						"delivery_connections": schema.Int64Attribute{
							MarkdownDescription: "Connections per delivery",
							Computed:            true,
						},
						"enable_jwt": schema.BoolAttribute{
							MarkdownDescription: "Enable JWT",
							Computed:            true,
						},
						"management_url": schema.StringAttribute{
							MarkdownDescription: "Management platform endpoint",
							Computed:            true,
						},
						"management_url_ssl_client": schema.StringAttribute{
							MarkdownDescription: "Management platform TLS client profile",
							Computed:            true,
						},
						"client_id": schema.StringAttribute{
							MarkdownDescription: "Client ID",
							Computed:            true,
						},
						"client_secret_alias": schema.StringAttribute{
							MarkdownDescription: "Client secret",
							Computed:            true,
						},
						"grant_type": schema.StringAttribute{
							MarkdownDescription: "Grant type",
							Computed:            true,
						},
						"scope": schema.StringAttribute{
							MarkdownDescription: "Scope",
							Computed:            true,
						},
						"persistent_connection": schema.BoolAttribute{
							MarkdownDescription: "Negotiate persistent connections",
							Computed:            true,
						},
						"timeout": schema.Int64Attribute{
							MarkdownDescription: "Timeout",
							Computed:            true,
						},
						"persistent_timeout": schema.Int64Attribute{
							MarkdownDescription: "Persistent timeout",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *AnalyticsEndpointDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *AnalyticsEndpointDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data AnalyticsEndpointList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.AnalyticsEndpoint{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.AnalyticsEndpoint{}
	if value := res.Get(`AnalyticsEndpoint`); value.Exists() {
		for _, v := range value.Array() {
			item := models.AnalyticsEndpoint{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.AnalyticsEndpointObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.AnalyticsEndpointObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
