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

type AMQPBrokerList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &AMQPBrokerDataSource{}
	_ datasource.DataSourceWithConfigure = &AMQPBrokerDataSource{}
)

func NewAMQPBrokerDataSource() datasource.DataSource {
	return &AMQPBrokerDataSource{}
}

type AMQPBrokerDataSource struct {
	pData *tfutils.ProviderData
}

func (d *AMQPBrokerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_amqp_broker"
}

func (d *AMQPBrokerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "In AMQP, distributed source and target termini are managed by a broker. The broker provides messaging services for communicating applications by periodically monitoring and polling termini. The broker ensures that sent messages are directed to the correct target terminus or are routed to another server. The AMQP broker configuration corresponds to an AMQP broker that is running on another host in the network. The configured properties enable communication between the DataPower Gateway and the remote AMQP broker.",
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
						"host_name": schema.StringAttribute{
							MarkdownDescription: "Host",
							Computed:            true,
						},
						"port": schema.Int64Attribute{
							MarkdownDescription: "Port",
							Computed:            true,
						},
						"xml_manager": schema.StringAttribute{
							MarkdownDescription: "Specify the XML manager to control access to the remote AMQP server. The XML manager obtains and manages documents.",
							Computed:            true,
						},
						"container_id": schema.StringAttribute{
							MarkdownDescription: "AMQP container ID",
							Computed:            true,
						},
						"authorization": schema.StringAttribute{
							MarkdownDescription: "Specify the SASL layer that the AMQP broker uses to authenticate with the AMQP server. The default setting in no authentication.",
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
						"maximum_frame_size": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum frame size in bytes to allow. Frames Frames that are larger are rejected. When rejected, the connection is closed. Enter a value in the range 512 - 104857600. The default value is 104857600.",
							Computed:            true,
						},
						"auto_retry": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enable the automatic retry procedure after an AMQP connection failure. By default, the automatic retry behavior is enabled. This setting does not affect attempts over an established connection.",
							Computed:            true,
						},
						"retry_interval": schema.Int64Attribute{
							MarkdownDescription: "Specify the interval in seconds to wait before attempting to reestablish a failed connection. After the number of attempts is reached, attempts to reestablish a failed connection use the interval that is defined by the long retry interval. Enter a value in the range 1 - 65535. The default value is 10. <p>This setting does not affect attempts over an established connection.</p>",
							Computed:            true,
						},
						"retry_attempts": schema.Int64Attribute{
							MarkdownDescription: "Specify the number of attempts for a failed connection to the remote AMQP server. After the number of attempts is reached, the long retry interval is used. Enter a value in the range 0 - 65535. The default value is 6. The special value of 0 disables the long interval, where the retry interval is used forever.",
							Computed:            true,
						},
						"long_retry_interval": schema.Int64Attribute{
							MarkdownDescription: "Specify the interval in seconds to use after the number of attempts is reached to attempt to reestablish a failed connection. Enter a value in the range 1 - 65535. The default value is 600. <p>This setting does not affect attempts over an established connection.</p>",
							Computed:            true,
						},
						"reporting_interval": schema.Int64Attribute{
							MarkdownDescription: "Specify the interval in seconds between the writing of identical log message. Enter a value in the range 1 - 65535. The default value is 10.",
							Computed:            true,
						},
						"ssl_client": schema.StringAttribute{
							MarkdownDescription: "TLS client profile",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *AMQPBrokerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *AMQPBrokerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data AMQPBrokerList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.AMQPBroker{
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
	l := []models.AMQPBroker{}
	if resFound {
		if value := res.Get(`AMQPBroker`); value.Exists() {
			for _, v := range value.Array() {
				item := models.AMQPBroker{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.AMQPBrokerObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.AMQPBrokerObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
