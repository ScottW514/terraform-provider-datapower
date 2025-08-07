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
	_ datasource.DataSource              = &ThrottlerDataSource{}
	_ datasource.DataSourceWithConfigure = &ThrottlerDataSource{}
)

func NewThrottlerDataSource() datasource.DataSource {
	return &ThrottlerDataSource{}
}

type ThrottlerDataSource struct {
	client *client.DatapowerClient
}

func (d *ThrottlerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_throttler"
}

func (d *ThrottlerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Throttle settings (`default` domain only)",
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Administrative state",
				Computed:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: "Comments",
				Computed:            true,
			},
			"throttle_at": schema.Int64Attribute{
				MarkdownDescription: "Memory throttle threshold",
				Computed:            true,
			},
			"terminate_at": schema.Int64Attribute{
				MarkdownDescription: "Memory terminate threshold",
				Computed:            true,
			},
			"temp_fs_throttle_at": schema.Int64Attribute{
				MarkdownDescription: "Temporary file space throttle threshold",
				Computed:            true,
			},
			"temp_fs_terminate_at": schema.Int64Attribute{
				MarkdownDescription: "Temporary file space terminate threshold",
				Computed:            true,
			},
			"qname_warn_at": schema.Int64Attribute{
				MarkdownDescription: "XML names and JSON keys threshold",
				Computed:            true,
			},
			"timeout": schema.Int64Attribute{
				MarkdownDescription: "Timeout",
				Computed:            true,
			},
			"statistics": schema.BoolAttribute{
				MarkdownDescription: "Status log",
				Computed:            true,
			},
			"log_level": schema.StringAttribute{
				MarkdownDescription: "Log level",
				Computed:            true,
			},
			"environmental_log": schema.BoolAttribute{
				MarkdownDescription: "Monitor sensors",
				Computed:            true,
			},
			"backlog_size": schema.Int64Attribute{
				MarkdownDescription: "Backlog size",
				Computed:            true,
			},
			"backlog_timeout": schema.Int64Attribute{
				MarkdownDescription: "Backlog timeout",
				Computed:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (d *ThrottlerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *ThrottlerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.Throttler

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := d.client.Get(data.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	data.FromBody(ctx, `Throttler`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
