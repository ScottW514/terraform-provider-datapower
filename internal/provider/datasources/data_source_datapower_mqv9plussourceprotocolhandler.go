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
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type MQv9PlusSourceProtocolHandlerList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &MQv9PlusSourceProtocolHandlerDataSource{}
	_ datasource.DataSourceWithConfigure = &MQv9PlusSourceProtocolHandlerDataSource{}
)

func NewMQv9PlusSourceProtocolHandlerDataSource() datasource.DataSource {
	return &MQv9PlusSourceProtocolHandlerDataSource{}
}

type MQv9PlusSourceProtocolHandlerDataSource struct {
	pData *tfutils.ProviderData
}

func (d *MQv9PlusSourceProtocolHandlerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mqv9plussourceprotocolhandler"
}

func (d *MQv9PlusSourceProtocolHandlerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Configure the IBM MQ v9+ handler to manage IBM MQ protocol communications.",
		Attributes: map[string]schema.Attribute{
			"app_domain": schema.StringAttribute{
				MarkdownDescription: "The name of the application domain the object belongs to",
				Required:            true,
			},
			"result": schema.ListNestedAttribute{
				MarkdownDescription: "List of objects",
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
						"queue_manager": schema.StringAttribute{
							MarkdownDescription: "Specify the name of the queue manager that provides messaging services for communicating applications by periodically monitoring or polling queues and by ensuring that messages are directed to the correct receive queue or routed to another queue manager. The local queue manager corresponds to a queue manager running on another host on the network.",
							Computed:            true,
						},
						"get_queue": schema.StringAttribute{
							MarkdownDescription: "Specify the name of the get queue associated with the queue manager. The handler gets messages from this queue.",
							Computed:            true,
						},
						"subscribe_topic_string": schema.StringAttribute{
							MarkdownDescription: "Specify the the topic string associated with the queue manager. The handler subscribes to this topic string and gets messages from it.",
							Computed:            true,
						},
						"subscription_name": schema.StringAttribute{
							MarkdownDescription: "Specify the subscription name of a durable subscription associated with the queue manager. This name identifies the subscription after reestablishing a lost connection.",
							Computed:            true,
						},
						"put_queue": schema.StringAttribute{
							MarkdownDescription: "Specify the the name of the put queue associated with the queue manager. The handler puts server-originated reply messages to this queue.",
							Computed:            true,
						},
						"publish_topic_string": schema.StringAttribute{
							MarkdownDescription: "Specify the publish topic string associated with the queue manager. The topic string describes the subject of the information that is published in a publish or subscribe message. The handler publishes messages to this topic string. If the put queue is specified, this property is ignored.",
							Computed:            true,
						},
						"code_page": schema.Int64Attribute{
							MarkdownDescription: "<p>Specify the coded character set identifier to which the remote IBM MQ queue manager converts output data. This property is meaningful only when the queue manager has the convert property set to on. The CCSI will be put in the CodeCharSetId field of MQMD.</p><p>The default CCSI is for ISO-8859-1 (latin-1).</p><p>For MQCCSI_EMBEDDED enter 4294967295; for MQCCSI_INHERIT enter 4294967294. For the other CCSIDs, refer to the IBM Code Pages.</p>",
							Computed:            true,
						},
						"get_message_options": schema.Int64Attribute{
							MarkdownDescription: "Specify the cumulative value of the MQGET options that are applicable to an IBM MQ message in decimal or hex format. The value is passed directly to the IBM MQ API. The default value is 1, which is the decimal value for the MQGMO_WAIT option.",
							Computed:            true,
						},
						"message_selector": schema.StringAttribute{
							MarkdownDescription: "Specify the selector that filters the messages from a queue or a subscription by their properties. The selector uses the SQL92 conditional expression syntax. The selector enables the handler to filter messages from a queue or a subscription.",
							Computed:            true,
						},
						"parse_properties": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to parse the properties of the incoming messages from a queue or a subscription. By default, properties are not parsed.",
							Computed:            true,
						},
						"async_put": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to asynchronously put a message to a queue without waiting for a response from the queue manager.",
							Computed:            true,
						},
						"exclude_headers": models.GetDmMQHeadersDataSourceSchema("Specify the headers after MQMD to strip from the message. By default only the MQMD header is parsed.", "exclude-headers", ""),
						"concurrent_connections": schema.Int64Attribute{
							MarkdownDescription: "Specify the number of concurrent IBM MQ conversations to allocate. The default value is 1 but can be increased to improve performance.",
							Computed:            true,
						},
						"polling_interval": schema.Int64Attribute{
							MarkdownDescription: "Specify the duration in seconds to wait after processing all messages before attempting to retrieve messages from the get queue.",
							Computed:            true,
						},
						"batch_size": schema.Int64Attribute{
							MarkdownDescription: "Specify the number of messages to process as a batch. The handler gathers the specified number of messages and processes them as a batch.",
							Computed:            true,
						},
						"content_type_header": schema.StringAttribute{
							MarkdownDescription: "Header to extract Content-Type",
							Computed:            true,
						},
						"content_type_x_path": schema.StringAttribute{
							MarkdownDescription: "XPath expression to extract Content-Type from IBM MQ header",
							Computed:            true,
						},
						"retrieve_backout_settings": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to retrieve backout setting from the IBM MQ server. <p>When enabled, retrieves the <b>Backout threshold</b> and <b>Backout requeue queue name</b> settings from the IBM MQ server and checks these values. On a reattempt, the handler uses the higher priority backout settings from the server. If the server does not contain backout settings, The handler uses any existing backout values, either empty or populated, from the local IBM MQ queue manager. If there are no backout settings, the backout function is disabled.</p><p>When an alias queue is used, its attributes are retrieved, not those of the base queue.</p>",
							Computed:            true,
						},
						"use_qm_name_in_url": schema.BoolAttribute{
							MarkdownDescription: "Specify whether the var://service/URL-in variable returns the name of the local queue manager or queue manager group when this configuration defines a queue manager group as the queue manager. <ul><li>When enabled, the variable returns the name of the queue manager.</li><li>When not enabled, the variable returns the name of the queue manager group. This setting is the default value.</li></ul>",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *MQv9PlusSourceProtocolHandlerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *MQv9PlusSourceProtocolHandlerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data MQv9PlusSourceProtocolHandlerList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.MQv9PlusSourceProtocolHandler{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.MQv9PlusSourceProtocolHandler{}
	if value := res.Get(`MQv9PlusSourceProtocolHandler`); value.Exists() {
		for _, v := range value.Array() {
			item := models.MQv9PlusSourceProtocolHandler{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.MQv9PlusSourceProtocolHandlerObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.MQv9PlusSourceProtocolHandlerObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
