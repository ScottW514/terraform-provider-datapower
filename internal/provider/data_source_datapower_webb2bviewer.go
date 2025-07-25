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
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
)

var (
	_ datasource.DataSource              = &WebB2BViewerDataSource{}
	_ datasource.DataSourceWithConfigure = &WebB2BViewerDataSource{}
)

func NewWebB2BViewerDataSource() datasource.DataSource {
	return &WebB2BViewerDataSource{}
}

type WebB2BViewerDataSource struct {
	client *client.DatapowerClient
}

func (d *WebB2BViewerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_webb2bviewer"
}

func (d *WebB2BViewerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "B2B viewer management service (`default` domain only)",
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Administrative state",
				Computed:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: "Comments",
				Computed:            true,
			},
			"local_port": schema.Int64Attribute{
				MarkdownDescription: "Port Number",
				Computed:            true,
			},
			"user_agent": schema.StringAttribute{
				MarkdownDescription: "Custom user agent",
				Computed:            true,
			},
			"idle_timeout": schema.Int64Attribute{
				MarkdownDescription: "Idle timeout",
				Computed:            true,
			},
			"acl": schema.StringAttribute{
				MarkdownDescription: "Access control list",
				Computed:            true,
			},
			"ssl_server_config_type": schema.StringAttribute{
				MarkdownDescription: "Custom TLS server type",
				Computed:            true,
			},
			"ssl_server": schema.StringAttribute{
				MarkdownDescription: "Custom TLS server profile",
				Computed:            true,
			},
			"sslsni_server": schema.StringAttribute{
				MarkdownDescription: "Custom TLS SNI server profile",
				Computed:            true,
			},
			"local_address": schema.StringAttribute{
				MarkdownDescription: "Local address",
				Computed:            true,
			},
		},
	}
}

func (d *WebB2BViewerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *WebB2BViewerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.WebB2BViewer

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := d.client.Get(data.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	data.FromBody(ctx, `WebB2BViewer`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
