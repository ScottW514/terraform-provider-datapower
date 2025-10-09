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
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type AnalyticsEndpointList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &AnalyticsEndpointDataSource{}
	_ datasource.DataSourceWithConfigure = &AnalyticsEndpointDataSource{}
)

func NewAnalyticsEndpointDataSource() datasource.DataSource {
	return &AnalyticsEndpointDataSource{}
}

type AnalyticsEndpointDataSource struct {
	pData *tfutils.ProviderData
}

func (d *AnalyticsEndpointDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_analytics_endpoint"
}

func (d *AnalyticsEndpointDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "An analytics endpoint buffers API event data and offloads the collected data as a bulk activity log to a remote server. When offloaded, you can use this data for display and analysis.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The name of the object to retrieve.",
				Optional:            true,
			},
			"app_domain": schema.StringAttribute{
				MarkdownDescription: "The name of the application domain the object belongs to.",
				Required:            true,
			},
			"result": schema.ListNestedAttribute{
				MarkdownDescription: "List of objects. If `id` was provided and it exists, it will be the only item in the list.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Name of the object. Must be unique among object types in application domain.",
							Computed:            true,
						},
						"app_domain": schema.StringAttribute{
							MarkdownDescription: "The name of the application domain the object belongs to",
							Computed:            true,
						},
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"analytics_server_url": schema.StringAttribute{
							MarkdownDescription: "Specify the URL to offload the collected API event data. The URL can start with <tt>http</tt> or <tt>https</tt> for an Elasticsearch server or start with <tt>dpkafka</tt> for a Kafka server. <ul><li>For an Elasticsearch server, specify the full URL to the endpoint starting with the <tt>http</tt> or <tt>https</tt> protocol identifier. With HTTPS, you must assign a TLS client profile.</li><li>For a Kafka server, specify only the name of the existing Kafka cluster configuration after the <tt>dpkafka</tt> protocol identifier. To complete the URL, you must specify which request topic to offload analytics data.</li></ul>",
							Computed:            true,
						},
						"ssl_client": schema.StringAttribute{
							MarkdownDescription: "TLS client profile",
							Computed:            true,
						},
						"request_topic": schema.StringAttribute{
							MarkdownDescription: "Request topic",
							Computed:            true,
						},
						"max_records": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum number of records that can be buffered for each API gateway. The collected analytics data for an API gateway is offloaded when 80% of this value or the defined interval is reached. The value must be a power of 2. Enter a value in the range 256 - 65536. The default value is 1024.",
							Computed:            true,
						},
						"max_records_memory_kb": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum size for each record in KB. Enter a value in the range 4 - 1024. The default value is 512.",
							Computed:            true,
						},
						"max_delivery_memory_mb": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum size for each delivery in MB. Enter a value in the range 1 - 1024. The default value is 512.",
							Computed:            true,
						},
						"interval": schema.Int64Attribute{
							MarkdownDescription: "Specify the interval in seconds between offloads. Data is offloaded at this interval or when an API gateway reaches 80% of the value set for maximum records. Enter a value in the range 1 - 3600. The default value is 600",
							Computed:            true,
						},
						"delivery_connections": schema.Int64Attribute{
							MarkdownDescription: "Specify the number of connections to establish per delivery to the remote server to offload analytics data. Each connection can carry a bulk activity log. Enter a value in the range 1 - 100. The default value is 1.",
							Computed:            true,
						},
						"enable_jwt": schema.BoolAttribute{
							MarkdownDescription: "Enable JWT feature sending logs to analytics server.",
							Computed:            true,
						},
						"management_url": schema.StringAttribute{
							MarkdownDescription: "Specify the URL of management platform endpoint to retrieve a JWT. The URL must use the <tt>http</tt> or <tt>https</tt> protocol.",
							Computed:            true,
						},
						"management_url_ssl_client": schema.StringAttribute{
							MarkdownDescription: "Management platform TLS client profile",
							Computed:            true,
						},
						"client_id": schema.StringAttribute{
							MarkdownDescription: "Client ID",
							Computed:            true,
						},
						"client_secret_alias": schema.StringAttribute{
							MarkdownDescription: "Client secret",
							Computed:            true,
						},
						"grant_type": schema.StringAttribute{
							MarkdownDescription: "Specify the grant type for requesting JWT tokens. Only the client credentials grant type is supported.",
							Computed:            true,
						},
						"scope": schema.StringAttribute{
							MarkdownDescription: "Specify the scope for requesting JWT tokens. The value is in the <tt>openid analytics_subsystem_ID/name</tt> format.",
							Computed:            true,
						},
						"persistent_connection": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to negotiate persistent connections. By default, persistent connections are enabled. The HTTP/2 protocol controls persistent connections and reuse. Therefore, these settings are ignored.",
							Computed:            true,
						},
						"timeout": schema.Int64Attribute{
							MarkdownDescription: "Specify the intra-transaction timeout for connections, which is the maximum idle time to allow in a transaction. This timer monitors idle time in the data transfer process. When the idle time is exceeded, the connection is torn down. Enter a value in the range 1 - 86400. The default value is 90.",
							Computed:            true,
						},
						"persistent_timeout": schema.Int64Attribute{
							MarkdownDescription: "Specify the inter-transaction timeout for connections, which is the maximum idle time to allow between the completion of a TCP transaction and the initiation of a new TCP transaction. When the idle time is exceeded, the connection is torn down. Enter a value in the range 1 - 86400. The default value is 60.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *AnalyticsEndpointDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *AnalyticsEndpointDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data AnalyticsEndpointList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.AnalyticsEndpoint{
		AppDomain: data.AppDomain,
	}

	path := o.GetPath()
	if !data.Id.IsNull() {
		path = path + "/" + data.Id.ValueString()
	}

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
	l := []models.AnalyticsEndpoint{}
	if resFound {
		if value := res.Get(`AnalyticsEndpoint`); value.Exists() {
			for _, v := range value.Array() {
				item := models.AnalyticsEndpoint{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.AnalyticsEndpointObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.AnalyticsEndpointObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
