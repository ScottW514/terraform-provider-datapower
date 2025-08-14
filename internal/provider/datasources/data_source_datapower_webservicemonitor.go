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

type WebServiceMonitorList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &WebServiceMonitorDataSource{}
	_ datasource.DataSourceWithConfigure = &WebServiceMonitorDataSource{}
)

func NewWebServiceMonitorDataSource() datasource.DataSource {
	return &WebServiceMonitorDataSource{}
}

type WebServiceMonitorDataSource struct {
	pData *tfutils.ProviderData
}

func (d *WebServiceMonitorDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_webservicemonitor"
}

func (d *WebServiceMonitorDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "A service level monitor (SLM) for a web service watches web services traffic to a specific endpoint.",
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
						"wsdlurl": schema.StringAttribute{
							MarkdownDescription: "Specify the URL of the WSDL file that defines the endpoints, transport type and operations to monitor. The WSDL file can be in the file system or remote.",
							Computed:            true,
						},
						"operations": schema.ListNestedAttribute{
							MarkdownDescription: "Operations to monitor",
							NestedObject:        models.DmWSSLMOpsDataSourceSchema,
							Computed:            true,
						},
						"endpoint_name": schema.StringAttribute{
							MarkdownDescription: "Endpoint",
							Computed:            true,
						},
						"endpoint_url": schema.StringAttribute{
							MarkdownDescription: "Endpoint URL",
							Computed:            true,
						},
						"frontend_url": schema.StringAttribute{
							MarkdownDescription: "Specify the URL that clients use to access the web service. This value cannot be the same value as the endpoint URL.. Wildcards are supported.",
							Computed:            true,
						},
						"transport": schema.StringAttribute{
							MarkdownDescription: "Specify the transport type for the endpoint. The transport type must agree with the transport type in the WSDL file.",
							Computed:            true,
						},
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *WebServiceMonitorDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *WebServiceMonitorDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data WebServiceMonitorList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.WebServiceMonitor{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.WebServiceMonitor{}
	if value := res.Get(`WebServiceMonitor`); value.Exists() {
		for _, v := range value.Array() {
			item := models.WebServiceMonitor{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.WebServiceMonitorObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.WebServiceMonitorObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
