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
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type OpenTelemetryExporterList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &OpenTelemetryExporterDataSource{}
	_ datasource.DataSourceWithConfigure = &OpenTelemetryExporterDataSource{}
)

func NewOpenTelemetryExporterDataSource() datasource.DataSource {
	return &OpenTelemetryExporterDataSource{}
}

type OpenTelemetryExporterDataSource struct {
	pData *tfutils.ProviderData
}

func (d *OpenTelemetryExporterDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_open_telemetry_exporter"
}

func (d *OpenTelemetryExporterDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "OpenTelemetry exporter",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The name of the object to retrieve.",
				Optional:            true,
			},
			"app_domain": schema.StringAttribute{
				MarkdownDescription: "The name of the application domain the object belongs to.",
				Required:            true,
			},
			"result": schema.ListNestedAttribute{
				MarkdownDescription: "List of objects. If `id` was provided and it exists, it will be the only item in the list.",
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
						"type": schema.StringAttribute{
							MarkdownDescription: "Exporter type",
							Computed:            true,
						},
						"host_name": schema.StringAttribute{
							MarkdownDescription: "Host",
							Computed:            true,
						},
						"traces_path": schema.StringAttribute{
							MarkdownDescription: "Traces",
							Computed:            true,
						},
						"port": schema.Int64Attribute{
							MarkdownDescription: "Port",
							Computed:            true,
						},
						"http_content_type": schema.StringAttribute{
							MarkdownDescription: "Content-Type",
							Computed:            true,
						},
						"timeout": schema.Int64Attribute{
							MarkdownDescription: "Timeout",
							Computed:            true,
						},
						"header": schema.ListNestedAttribute{
							MarkdownDescription: "Headers",
							NestedObject:        models.GetDmOpenTelemetryExporterHeaderDataSourceSchema(),
							Computed:            true,
						},
						"processor": schema.StringAttribute{
							MarkdownDescription: "Processor",
							Computed:            true,
						},
						"max_queue_size": schema.Int64Attribute{
							MarkdownDescription: "Max queue size",
							Computed:            true,
						},
						"max_export_size": schema.Int64Attribute{
							MarkdownDescription: "Max export size",
							Computed:            true,
						},
						"export_delay_interval": schema.Int64Attribute{
							MarkdownDescription: "Export interval",
							Computed:            true,
						},
						"proxy_policies": schema.ListNestedAttribute{
							MarkdownDescription: "Proxy policies",
							NestedObject:        models.GetDmAPIProxyPolicyDataSourceSchema(),
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

func (d *OpenTelemetryExporterDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *OpenTelemetryExporterDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data OpenTelemetryExporterList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.OpenTelemetryExporter{
		AppDomain: data.AppDomain,
	}

	path := o.GetPath()
	if !data.Id.IsNull() {
		path = path + "/" + data.Id.ValueString()
	}

	res, err := d.pData.Client.Get(path)
	resFound := true
	if err != nil {
		if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(d.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString())
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		} else {
			resFound = false
		}
	}
	l := []models.OpenTelemetryExporter{}
	if resFound {
		if value := res.Get(`OpenTelemetryExporter`); value.Exists() {
			for _, v := range value.Array() {
				item := models.OpenTelemetryExporter{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.OpenTelemetryExporterObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.OpenTelemetryExporterObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
