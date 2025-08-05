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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
)

var (
	_ datasource.DataSource              = &NFSDynamicMountsDataSource{}
	_ datasource.DataSourceWithConfigure = &NFSDynamicMountsDataSource{}
)

func NewNFSDynamicMountsDataSource() datasource.DataSource {
	return &NFSDynamicMountsDataSource{}
}

type NFSDynamicMountsDataSource struct {
	client *client.DatapowerClient
}

func (d *NFSDynamicMountsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nfsdynamicmounts"
}

func (d *NFSDynamicMountsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "NFS dynamic mounts",
		Attributes: map[string]schema.Attribute{
			"app_domain": schema.StringAttribute{
				MarkdownDescription: "The name of the application domain the object belongs to",
				Required:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Administrative state",
				Computed:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: "Comments",
				Computed:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "NFS version",
				Computed:            true,
			},
			"transport": schema.StringAttribute{
				MarkdownDescription: "Transport protocol",
				Computed:            true,
			},
			"mount_type": schema.StringAttribute{
				MarkdownDescription: "Mount type",
				Computed:            true,
			},
			"read_only": schema.BoolAttribute{
				MarkdownDescription: "Read-Only",
				Computed:            true,
			},
			"read_size": schema.Int64Attribute{
				MarkdownDescription: "Read size",
				Computed:            true,
			},
			"write_size": schema.Int64Attribute{
				MarkdownDescription: "Write size",
				Computed:            true,
			},
			"timeout": schema.Int64Attribute{
				MarkdownDescription: "Retransmission timeout",
				Computed:            true,
			},
			"retransmissions": schema.Int64Attribute{
				MarkdownDescription: "Max retransmissions",
				Computed:            true,
			},
			"idle_unmount_seconds": schema.Int64Attribute{
				MarkdownDescription: "Inactivity timeout",
				Computed:            true,
			},
			"mount_timeout_seconds": schema.Int64Attribute{
				MarkdownDescription: "Mount timeout",
				Computed:            true,
			},
		},
	}
}

func (d *NFSDynamicMountsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *NFSDynamicMountsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.NFSDynamicMounts

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := d.client.Get(data.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	data.FromBody(ctx, `NFSDynamicMounts`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
