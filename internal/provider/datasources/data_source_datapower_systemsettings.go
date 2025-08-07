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
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
)

var (
	_ datasource.DataSource              = &SystemSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &SystemSettingsDataSource{}
)

func NewSystemSettingsDataSource() datasource.DataSource {
	return &SystemSettingsDataSource{}
}

type SystemSettingsDataSource struct {
	client *client.DatapowerClient
}

func (d *SystemSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemsettings"
}

func (d *SystemSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "System settings (`default` domain only)",
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Administrative state",
				Computed:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: "Comments",
				Computed:            true,
			},
			"product_oid": schema.StringAttribute{
				MarkdownDescription: "Product OID",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description",
				Computed:            true,
			},
			"serial_number": schema.StringAttribute{
				MarkdownDescription: "Serial number",
				Computed:            true,
			},
			"entitlement_number": schema.StringAttribute{
				MarkdownDescription: "Entitlement serial number",
				Computed:            true,
			},
			"product_id": schema.StringAttribute{
				MarkdownDescription: "Product ID",
				Computed:            true,
			},
			"capacity_mode": schema.StringAttribute{
				MarkdownDescription: "Licensed capacity mode",
				Computed:            true,
			},
			"contact": schema.StringAttribute{
				MarkdownDescription: "Contact",
				Computed:            true,
			},
			"system_name": schema.StringAttribute{
				MarkdownDescription: "System name",
				Computed:            true,
			},
			"location": schema.StringAttribute{
				MarkdownDescription: "Location",
				Computed:            true,
			},
			"services": schema.Int64Attribute{
				MarkdownDescription: "Services",
				Computed:            true,
			},
			"backup_mode": schema.StringAttribute{
				MarkdownDescription: "Backup mode",
				Computed:            true,
			},
			"product_mode": schema.StringAttribute{
				MarkdownDescription: "Product Mode",
				Computed:            true,
			},
			"custom_ui_file": schema.StringAttribute{
				MarkdownDescription: "Custom user interface file",
				Computed:            true,
			},
			"audit_reserve": schema.Int64Attribute{
				MarkdownDescription: "Audit reserve space",
				Computed:            true,
			},
			"detect_intrusion": schema.StringAttribute{
				MarkdownDescription: "Detect intrusion (physical appliances only)",
				Computed:            true,
			},
			"hardware_xml_acceleration": schema.BoolAttribute{
				MarkdownDescription: "Enable hardware XML acceleration (physical appliances only)",
				Computed:            true,
			},
			"locale": schema.StringAttribute{
				MarkdownDescription: "System locale",
				Computed:            true,
			},
			"system_log_fixed_format": schema.BoolAttribute{
				MarkdownDescription: "Enable fixed format",
				Computed:            true,
			},
			"uuid": schema.StringAttribute{
				MarkdownDescription: "UUID",
				Computed:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (d *SystemSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *SystemSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.SystemSettings

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := d.client.Get(data.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	data.FromBody(ctx, `SystemSettings`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
