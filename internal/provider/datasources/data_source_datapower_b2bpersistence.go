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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var (
	_ datasource.DataSource              = &B2BPersistenceDataSource{}
	_ datasource.DataSourceWithConfigure = &B2BPersistenceDataSource{}
)

func NewB2BPersistenceDataSource() datasource.DataSource {
	return &B2BPersistenceDataSource{}
}

type B2BPersistenceDataSource struct {
	pData *tfutils.ProviderData
}

func (d *B2BPersistenceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_b2bpersistence"
}

func (d *B2BPersistenceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "B2B persistence (`default` domain only)",
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Administrative state",
				Computed:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: "Comments",
				Computed:            true,
			},
			"raid_volume": schema.StringAttribute{
				MarkdownDescription: "RAID volume",
				Computed:            true,
			},
			"storage_size": schema.Int64Attribute{
				MarkdownDescription: "Storage size",
				Computed:            true,
			},
			"ha_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable high availability",
				Computed:            true,
			},
			"ha_other_hosts": models.GetDmB2BHAHostDataSourceSchema("Alternate host", "ha-other-hosts", ""),
			"ha_local_ip": schema.StringAttribute{
				MarkdownDescription: "Replication address",
				Computed:            true,
			},
			"ha_local_port": schema.Int64Attribute{
				MarkdownDescription: "Replication port",
				Computed:            true,
			},
			"ha_virtual_ip": schema.StringAttribute{
				MarkdownDescription: "Virtual IP address",
				Computed:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (d *B2BPersistenceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *B2BPersistenceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data models.B2BPersistence
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := d.pData.Client.Get(data.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	data.FromBody(ctx, `B2BPersistence`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
