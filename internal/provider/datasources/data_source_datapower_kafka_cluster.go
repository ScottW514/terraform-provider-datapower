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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type KafkaClusterList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &KafkaClusterDataSource{}
	_ datasource.DataSourceWithConfigure = &KafkaClusterDataSource{}
)

func NewKafkaClusterDataSource() datasource.DataSource {
	return &KafkaClusterDataSource{}
}

type KafkaClusterDataSource struct {
	pData *tfutils.ProviderData
}

func (d *KafkaClusterDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_kafka_cluster"
}

func (d *KafkaClusterDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Define the Kafka cluster that is responsible for the messaging services. The Kafka cluster periodically monitors and polls topics. The Kafka cluster ensures that sent messages are directed to the correct response topic or are routed to another server.",
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
						"protocol": schema.StringAttribute{
							MarkdownDescription: "Specify the transport protocol for the Kafka bootstrap connection. The selected protocol is used for the exchange of information between the Kafka server and the bootstrap server. By default, uses a non-encrypted transport.",
							Computed:            true,
						},
						"endpoint": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the endpoints for the bootstrap process. A bootstrap server uses a host name or IP address and a port to define an endpoint address. You can add multiple nondefault bootstrap servers. For failover capability, the endpoints must be members of the same cluster.",
							NestedObject:        models.GetDmKafkaEndpointDataSourceSchema(),
							Computed:            true,
						},
						"sasl_mechanism": schema.StringAttribute{
							MarkdownDescription: "Specify the Simple Authentication and Security Layer (SASL) mechanism to communicate with the Kafka cluster. By default, uses a clear text password.",
							Computed:            true,
						},
						"user_name": schema.StringAttribute{
							MarkdownDescription: "Username",
							Computed:            true,
						},
						"password_alias": schema.StringAttribute{
							MarkdownDescription: "Password alias",
							Computed:            true,
						},
						"autocommit": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to commit offsets at the defined interval or at process-completion. <ul><li>When enabled, commits offsets at the defined interval. The default interval is 5 seconds. To change the interval, set the <tt>auto.commit.interval.ms</tt> property.</li><li>When disabled, commits offsets at process-completion. You can use the batch size setting for the Kafka handle to define the number of messages to attempt to receive from the consumer.</li></ul>",
							Computed:            true,
						},
						"ssl_client": schema.StringAttribute{
							MarkdownDescription: "TLS client profile",
							Computed:            true,
						},
						"memory_threshold": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum memory to allocate in bytes. Enter a value in the range 10485760 - 1073741824. The default value is 268435456.",
							Computed:            true,
						},
						"maximum_message_size": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum message size in bytes. Enter a value in the range 0 - 1073741824. The default value is 1048576. A value of 0 disables the enforcement of a maximum message size.",
							Computed:            true,
						},
						"auto_retry": schema.BoolAttribute{
							MarkdownDescription: "Automatic retry",
							Computed:            true,
						},
						"retry_interval": schema.Int64Attribute{
							MarkdownDescription: "Specify the interval between attempts to reestablish a connection in seconds. Enter a value in the range 1 - 65535. The default value is 10.",
							Computed:            true,
						},
						"property": schema.ListNestedAttribute{
							MarkdownDescription: "Specify extra property to configure the connection to the Kafka server. Use this property for each extra property that is required. Some properties are unsupported and will cause a configuration failure.",
							NestedObject:        models.GetDmKafkaPropertyDataSourceSchema(),
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *KafkaClusterDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *KafkaClusterDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data KafkaClusterList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.KafkaCluster{
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
	l := []models.KafkaCluster{}
	if resFound {
		if value := res.Get(`KafkaCluster`); value.Exists() {
			for _, v := range value.Array() {
				item := models.KafkaCluster{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.KafkaClusterObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.KafkaClusterObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
