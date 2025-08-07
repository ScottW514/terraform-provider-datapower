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
	_ datasource.DataSource              = &DNSNameServiceDataSource{}
	_ datasource.DataSourceWithConfigure = &DNSNameServiceDataSource{}
)

func NewDNSNameServiceDataSource() datasource.DataSource {
	return &DNSNameServiceDataSource{}
}

type DNSNameServiceDataSource struct {
	client *client.DatapowerClient
}

func (d *DNSNameServiceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnsnameservice"
}

func (d *DNSNameServiceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DNS settings (`default` domain only)",
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Administrative state",
				Computed:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: "Comments",
				Computed:            true,
			},
			"search_domains": schema.ListNestedAttribute{
				MarkdownDescription: "Search domains",
				NestedObject:        models.DmSearchDomainDataSourceSchema,
				Computed:            true,
			},
			"name_servers": schema.ListNestedAttribute{
				MarkdownDescription: "DNS servers",
				NestedObject:        models.DmNameServerDataSourceSchema,
				Computed:            true,
			},
			"static_hosts": schema.ListNestedAttribute{
				MarkdownDescription: "Static hosts",
				NestedObject:        models.DmStaticHostDataSourceSchema,
				Computed:            true,
			},
			"ip_preference": schema.StringAttribute{
				MarkdownDescription: "IP preference",
				Computed:            true,
			},
			"force_ip_preference": schema.BoolAttribute{
				MarkdownDescription: "Force IP preference",
				Computed:            true,
			},
			"load_balance_algorithm": schema.StringAttribute{
				MarkdownDescription: "Load distribution algorithm",
				Computed:            true,
			},
			"max_retries": schema.Int64Attribute{
				MarkdownDescription: "Attempts",
				Computed:            true,
			},
			"timeout": schema.Int64Attribute{
				MarkdownDescription: "Timeout",
				Computed:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (d *DNSNameServiceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *DNSNameServiceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.DNSNameService

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := d.client.Get(data.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	data.FromBody(ctx, `DNSNameService`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
