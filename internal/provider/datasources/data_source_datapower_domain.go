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
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
)

type DomainList struct {
	Result types.List `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &DomainDataSource{}
	_ datasource.DataSourceWithConfigure = &DomainDataSource{}
)

func NewDomainDataSource() datasource.DataSource {
	return &DomainDataSource{}
}

type DomainDataSource struct {
	client *client.DatapowerClient
}

func (d *DomainDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_domain"
}

func (d *DomainDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Application domain",
		Attributes: map[string]schema.Attribute{
			"result": schema.ListNestedAttribute{
				MarkdownDescription: "List of objects",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"app_domain": schema.StringAttribute{
							MarkdownDescription: "The name of the application domain the object belongs to",
							Computed:            true,
						},
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"config_dir": schema.StringAttribute{
							MarkdownDescription: "Configuration directory",
							Computed:            true,
						},
						"neighbor_domain": schema.ListAttribute{
							MarkdownDescription: "Visible domains",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"domain_user": schema.ListAttribute{
							MarkdownDescription: "CLI user access",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"file_map":       models.GetDmDomainFileMapDataSourceSchema("File permission to the local: directory", "file-permissions", ""),
						"monitoring_map": models.GetDmDomainMonitoringMapDataSourceSchema("File-monitoring of the local: directory", "file-monitoring", ""),
						"config_mode": schema.StringAttribute{
							MarkdownDescription: "Configuration mode",
							Computed:            true,
						},
						"import_url": schema.StringAttribute{
							MarkdownDescription: "Import URL",
							Computed:            true,
						},
						"import_format": schema.StringAttribute{
							MarkdownDescription: "Import format",
							Computed:            true,
						},
						"deployment_policy": schema.StringAttribute{
							MarkdownDescription: "Deployment policy",
							Computed:            true,
						},
						"deployment_policy_parameters": schema.StringAttribute{
							MarkdownDescription: "Deployment policy variables",
							Computed:            true,
						},
						"local_ip_rewrite": schema.BoolAttribute{
							MarkdownDescription: "Rewrite local IP addresses",
							Computed:            true,
						},
						"max_chkpoints": schema.Int64Attribute{
							MarkdownDescription: "Checkpoint limit",
							Computed:            true,
						},
						"config_permissions_mode": schema.StringAttribute{
							MarkdownDescription: "Configuration permissions mode",
							Computed:            true,
						},
						"config_permissions_profile": schema.StringAttribute{
							MarkdownDescription: "Global permissions profile",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *DomainDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *DomainDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DomainList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.Domain{}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.Domain{}
	if value := res.Get(`Domain`); value.Exists() {
		for _, v := range value.Array() {
			item := models.Domain{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.DomainObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.DomainObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
