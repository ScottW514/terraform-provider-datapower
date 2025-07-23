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
	_ datasource.DataSource              = &QuotaEnforcementServerDataSource{}
	_ datasource.DataSourceWithConfigure = &QuotaEnforcementServerDataSource{}
)

func NewQuotaEnforcementServerDataSource() datasource.DataSource {
	return &QuotaEnforcementServerDataSource{}
}

type QuotaEnforcementServerDataSource struct {
	client *client.DatapowerClient
}

func (d *QuotaEnforcementServerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_quotaenforcementserver"
}

func (d *QuotaEnforcementServerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Quota Enforcement Server (`default` domain only)",
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Administrative state",
				Computed:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: "Comments",
				Computed:            true,
			},
			"password_alias": schema.StringAttribute{
				MarkdownDescription: "Password alias",
				Computed:            true,
			},
			"raid_volume": schema.StringAttribute{
				MarkdownDescription: "Data storage location",
				Computed:            true,
			},
			"server_port": schema.Int64Attribute{
				MarkdownDescription: "Server port",
				Computed:            true,
			},
			"monitor_port": schema.Int64Attribute{
				MarkdownDescription: "Monitor port",
				Computed:            true,
			},
			"enable_peer_group": schema.BoolAttribute{
				MarkdownDescription: "Peer group mode",
				Computed:            true,
			},
			"enable_ssl": schema.BoolAttribute{
				MarkdownDescription: "Enable TLS",
				Computed:            true,
			},
			"ssl_crypto_key": schema.StringAttribute{
				MarkdownDescription: "Crypto key",
				Computed:            true,
			},
			"ssl_crypto_certificate": schema.StringAttribute{
				MarkdownDescription: "Certificate",
				Computed:            true,
			},
			"ip_address": schema.StringAttribute{
				MarkdownDescription: "IP address",
				Computed:            true,
			},
			"peers": schema.ListAttribute{
				MarkdownDescription: "Peers",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: "Priority",
				Computed:            true,
			},
			"strict_mode": schema.BoolAttribute{
				MarkdownDescription: "Strict mode",
				Computed:            true,
			},
		},
	}
}

func (d *QuotaEnforcementServerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *QuotaEnforcementServerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.QuotaEnforcementServer

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := d.client.Get(data.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	data.FromBody(ctx, `QuotaEnforcementServer`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
