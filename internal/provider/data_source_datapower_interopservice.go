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
	_ datasource.DataSource              = &InteropServiceDataSource{}
	_ datasource.DataSourceWithConfigure = &InteropServiceDataSource{}
)

func NewInteropServiceDataSource() datasource.DataSource {
	return &InteropServiceDataSource{}
}

type InteropServiceDataSource struct {
	client *client.DatapowerClient
}

func (d *InteropServiceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interopservice"
}

func (d *InteropServiceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Interoperability test service (`default` domain only)",
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Administrative state",
				Computed:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: "Comments",
				Computed:            true,
			},
			"xml_manager": schema.StringAttribute{
				MarkdownDescription: "XML manager",
				Computed:            true,
			},
			"aaa_policy": schema.StringAttribute{
				MarkdownDescription: "AAA policy",
				Computed:            true,
			},
			"http_service": schema.BoolAttribute{
				MarkdownDescription: "Enable over HTTP",
				Computed:            true,
			},
			"local_address": schema.StringAttribute{
				MarkdownDescription: "Local IP address",
				Computed:            true,
			},
			"local_port": schema.Int64Attribute{
				MarkdownDescription: "Local port",
				Computed:            true,
			},
			"acl": schema.StringAttribute{
				MarkdownDescription: "Access control list",
				Computed:            true,
			},
			"https_service": schema.BoolAttribute{
				MarkdownDescription: "Enable over HTTPS",
				Computed:            true,
			},
			"https_local_address": schema.StringAttribute{
				MarkdownDescription: "Local IP address",
				Computed:            true,
			},
			"https_local_port": schema.Int64Attribute{
				MarkdownDescription: "Local port",
				Computed:            true,
			},
			"https_acl": schema.StringAttribute{
				MarkdownDescription: "Access control list",
				Computed:            true,
			},
			"ssl_server_config_type": schema.StringAttribute{
				MarkdownDescription: "TLS server type",
				Computed:            true,
			},
			"ssl_server": schema.StringAttribute{
				MarkdownDescription: "TLS server profile",
				Computed:            true,
			},
			"sslsni_server": schema.StringAttribute{
				MarkdownDescription: "TLS SNI server profile",
				Computed:            true,
			},
		},
	}
}

func (d *InteropServiceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *InteropServiceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.InteropService

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := d.client.Get(data.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	data.FromBody(ctx, `InteropService`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
