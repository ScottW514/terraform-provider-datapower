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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
)

var (
	_ datasource.DataSource              = &SNMPSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &SNMPSettingsDataSource{}
)

func NewSNMPSettingsDataSource() datasource.DataSource {
	return &SNMPSettingsDataSource{}
}

type SNMPSettingsDataSource struct {
	client *client.DatapowerClient
}

func (d *SNMPSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmpsettings"
}

func (d *SNMPSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SNMP Settings (`default` domain only)",
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Administrative state",
				Computed:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: "Comments",
				Computed:            true,
			},
			"local_address": schema.StringAttribute{
				MarkdownDescription: "Local IP Address",
				Computed:            true,
			},
			"local_port": schema.Int64Attribute{
				MarkdownDescription: "Local Port",
				Computed:            true,
			},
			"policies": schema.ListNestedAttribute{
				MarkdownDescription: "SNMPv1/v2c Communities",
				NestedObject:        models.DmSnmpPolicyDataSourceSchema,
				Computed:            true,
			},
			"policies_mq": schema.ListNestedAttribute{
				MarkdownDescription: "SNMPv1/v2c Communities",
				NestedObject:        models.DmSnmpPolicyMQDataSourceSchema,
				Computed:            true,
			},
			"targets": schema.ListNestedAttribute{
				MarkdownDescription: "Trap and Notification Targets",
				NestedObject:        models.DmSnmpTargetDataSourceSchema,
				Computed:            true,
			},
			"users": schema.ListAttribute{
				MarkdownDescription: "SNMPv3 Users",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"contexts": schema.ListNestedAttribute{
				MarkdownDescription: "SNMPv3 Contexts",
				NestedObject:        models.DmSnmpContextDataSourceSchema,
				Computed:            true,
			},
			"security_level": schema.StringAttribute{
				MarkdownDescription: "SNMPv3 Security Level",
				Computed:            true,
			},
			"access_level": schema.StringAttribute{
				MarkdownDescription: "SNMPv3 Access Level",
				Computed:            true,
			},
			"enable_default_trap_subscriptions": schema.BoolAttribute{
				MarkdownDescription: "Enable Default Event Subscriptions",
				Computed:            true,
			},
			"trap_priority": schema.StringAttribute{
				MarkdownDescription: "Minimum Priority",
				Computed:            true,
			},
			"trap_event_code": schema.ListAttribute{
				MarkdownDescription: "Event Subscriptions",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"config_mib": schema.StringAttribute{
				MarkdownDescription: "Configuration",
				Computed:            true,
			},
			"config_mib_mq": schema.StringAttribute{
				MarkdownDescription: "Configuration",
				Computed:            true,
			},
			"status_mib": schema.StringAttribute{
				MarkdownDescription: "Status",
				Computed:            true,
			},
			"status_mib_mq": schema.StringAttribute{
				MarkdownDescription: "Status",
				Computed:            true,
			},
			"notif_mib": schema.StringAttribute{
				MarkdownDescription: "Notifications",
				Computed:            true,
			},
			"notif_mib_mq": schema.StringAttribute{
				MarkdownDescription: "Notifications",
				Computed:            true,
			},
		},
	}
}

func (d *SNMPSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *SNMPSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.SNMPSettings

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := d.client.Get(data.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	data.FromBody(ctx, `SNMPSettings`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
