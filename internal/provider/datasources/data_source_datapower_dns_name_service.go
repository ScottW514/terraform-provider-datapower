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
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var (
	_ datasource.DataSource              = &DNSNameServiceDataSource{}
	_ datasource.DataSourceWithConfigure = &DNSNameServiceDataSource{}
)

func NewDNSNameServiceDataSource() datasource.DataSource {
	return &DNSNameServiceDataSource{}
}

type DNSNameServiceDataSource struct {
	pData *tfutils.ProviderData
}

func (d *DNSNameServiceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dns_name_service"
}

func (d *DNSNameServiceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Configure the DNS client with the DNS servers to contact to resolve hostnames to IP addresses.",
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "<p>The administrative state of the configuration.</p><ul><li>To make active, set to enabled.</li><li>To make inactive, set to disabled.</li></ul>",
				Computed:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: "Comments",
				Computed:            true,
			},
			"search_domains": schema.ListNestedAttribute{
				MarkdownDescription: "Specify the list of search domains to resolve partial hostnames.",
				NestedObject:        models.GetDmSearchDomainDataSourceSchema(),
				Computed:            true,
			},
			"name_servers": schema.ListNestedAttribute{
				MarkdownDescription: "Specify the list of DNS servers to contact to resolve hostnames. If you define multiple servers, ensure that the sequence to contact the servers is your preferred order.",
				NestedObject:        models.GetDmNameServerDataSourceSchema(),
				Computed:            true,
			},
			"static_hosts": schema.ListNestedAttribute{
				MarkdownDescription: "Specify the static map of hostnames to IP addresses that do not use DNS resolution. Because the local resolver uses a cache, static hosts do not improve performance.",
				NestedObject:        models.GetDmStaticHostDataSourceSchema(),
				Computed:            true,
			},
			"ip_preference": schema.StringAttribute{
				MarkdownDescription: "Specify the preferred IP version to resolve hostnames. When a hostname resolves to both IPv4 and IPv6 addresses, this setting determines which version to use.",
				Computed:            true,
			},
			"force_ip_preference": schema.BoolAttribute{
				MarkdownDescription: "Specify whether to restrict DNS queries to the preferred IP version to resolve hostnames. You want to force the IP preference except when both IPv4 and IPv6 addresses are in use. When not forced, the device resolves each hostname by querying A and AAAA records and waiting for both responses or a timeout. Waiting for the response or timeout for both records can introduce unnecessary latency in DNS resolution.",
				Computed:            true,
			},
			"load_balance_algorithm": schema.StringAttribute{
				MarkdownDescription: "Specify the load distribution algorithm to resolve hostnames. The default algorithm is first-alive.",
				Computed:            true,
			},
			"max_retries": schema.Int64Attribute{
				MarkdownDescription: "For the first alive algorithm, specify the maximum number of resolution attempts to send a query to the list of name servers before an error is returned. By default, an unacknowledged resolution request is attempted 3 times.",
				Computed:            true,
			},
			"timeout": schema.Int64Attribute{
				MarkdownDescription: "For the first alive algorithm, specify the duration in seconds that the resolver waits for a response from a DNS server. After expiry, the resolver attempts the query to a different DNS server. The default value is 5.",
				Computed:            true,
			},
		},
	}
}

func (d *DNSNameServiceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *DNSNameServiceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data models.DNSNameService
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	path := data.GetPath()

	res, err := d.pData.Client.Get(path)
	resFound := true
	if err != nil {
		if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		} else {
			resFound = false
		}
	}
	if resFound {
		data.FromBody(ctx, `DNSNameService`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
