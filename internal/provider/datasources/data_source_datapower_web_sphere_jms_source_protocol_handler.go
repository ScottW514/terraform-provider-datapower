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

type WebSphereJMSSourceProtocolHandlerList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &WebSphereJMSSourceProtocolHandlerDataSource{}
	_ datasource.DataSourceWithConfigure = &WebSphereJMSSourceProtocolHandlerDataSource{}
)

func NewWebSphereJMSSourceProtocolHandlerDataSource() datasource.DataSource {
	return &WebSphereJMSSourceProtocolHandlerDataSource{}
}

type WebSphereJMSSourceProtocolHandlerDataSource struct {
	pData *tfutils.ProviderData
}

func (d *WebSphereJMSSourceProtocolHandlerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_web_sphere_jms_source_protocol_handler"
}

func (d *WebSphereJMSSourceProtocolHandlerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Configure the WebSphere JMS handler to manage WebSphere JMS protocol communications.",
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
						"server": schema.StringAttribute{
							MarkdownDescription: "Select the WebSphere JMS server object supported by this handler.",
							Computed:            true,
						},
						"request_topic_space": schema.StringAttribute{
							MarkdownDescription: "<p>Use this property to disambiguate a topic if the request destination is a topic whose name appears in multiple topic spaces.</p><p>A topic space is a hierarchy of topics used for publish/subscribe messaging. Topics with the same name can exist in multiple topic spaces, but there can be only one topic space with a given name in a service integration bus.</p><p>For example, consider a topic hierarchy split into the following topic spaces:</p><ul><li>library - topics for document management</li><li>sales - topics for marketing and sales tracking</li><li>engineering - topics for engineering and technology</li></ul><p>The topic <em>volumes</em> can appear in all three topic spaces, and have a different meaning in each.</p><p>Enter the name of the target topic space if necessary.</p>",
							Computed:            true,
						},
						"reply_topic_space": schema.StringAttribute{
							MarkdownDescription: "<p>Use this property to disambiguate a topic if the response destination is a topic whose name appears in multiple topic spaces.</p><p>A topic space is a hierarchy of topics used for publish/subscribe messaging. Topics with the same name can exist in multiple topic spaces, but there can be only one topic space with a given name in a service integration bus.</p><p>For example, consider a topic hierarchy split into the following topic spaces:</p><ul><li>library - topics for document management</li><li>sales - topics for marketing and sales tracking</li><li>engineering - topics for engineering and technology</li></ul><p>The topic <em>volumes</em> can appear in all three topic spaces, and have a different meaning in each.</p><p>Enter the name of the target topic space if necessary.</p>",
							Computed:            true,
						},
						"strict_message_order": schema.BoolAttribute{
							MarkdownDescription: "<p>Enable to work smoothly with WebSphere server when the \"Strict Message Ordering\" option in the corresponding destination is checked.</p>",
							Computed:            true,
						},
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"get_queue": schema.StringAttribute{
							MarkdownDescription: "<p>Enter the name of the get queue associated with the WebSphere JMS or TIBCO EMS object supported by this handler.</p><p>The handler monitors the get queue for incoming client requests. Upon message receipt, The handler forwards the extracted message to the DataPower object that will gateway the message to a remote message provider.</p>",
							Computed:            true,
						},
						"put_queue": schema.StringAttribute{
							MarkdownDescription: "<p>Enter the name of the put queue associated with the WebSphere JMS or TIBCO EMS object supported by this handler.</p><p>The put queue contains server-originated WAS JMS or TIBCO EMS reply messages.</p><p>Such messages are originated by a remote WAS JMS or TIBCO EMS message provider and put into this queue by a local WebSphere JMS or TIBCO EMS object.</p><p>Configuration of a put queue is optional.</p><p>A put queue should be configured if server replies are expected; if reply messages are not expected, a put queue need not be configured.</p><p>In the absence of a put queue, any received replies are dropped.</p>",
							Computed:            true,
						},
						"selector": schema.StringAttribute{
							MarkdownDescription: "<p>Provide an SQL-like expression to filter messages from the GET queue.</p><p>For example, <tt>DeliveryMode LIKE PERSISTENT</tt></p><p>This expression specifies that only client requests that have a DeliveryMode of PERSISTENT are forwarded to the WebSphere JMS or TIBCO EMS object for processing; all other messages are dropped from the get queue.</p><p>The message selector is a conditional expression based on a subset of SQL92 conditional expression syntax. The conditional expression enables the handler to identify <em>messages of interest</em> .</p><p>The conditional expression does not operate on the body of the message, rather it examines message headers and properties (proprietary user-created headers that might appear between the required headers and the message body).</p><p>The required headers are as follows:</p><ul><li><tt>Destination</tt> - contains the destination (queue) to which the message is being sent</li><li><tt>DeliveryMode</tt> - contains the delivery mode (PERSISTENT or NON_PERSISTENT)</li><li><tt>Expiration</tt> - contains a message TTL or a value of 0 indicating an unlimited TTL</li><li><tt>Priority</tt> - contains the message priority expressed as a digit from 0 (lowest priority) to 9 (highest priority)</li><li><tt>MessageID</tt> - contains a unique message identifier starting with the prefix ID:, or a null value, effectively disabling message ID</li><li><tt>Timestamp</tt> - contains the time the message was handed off for transmission, not the time it was actually sent</li><li><tt>CorrelationID</tt> - contains a means of associating one message (for example, a response) with another message (for example, the original request)</li><li><tt>ReplyTo</tt> - contains the destination (queue) to which a reply to this message should be sent</li><li><tt>Type</tt> - contains a message identifier provided by the application</li><li><tt>Redelivered</tt> - contains a boolean indicating that the message has been delivered in the past, but not yet acknowledged</li></ul><p>Configuration of a message selector is optional. If a message selector is not specified, all incoming client request messages are transferred by The handler to the DataPower object for processing.</p>",
							Computed:            true,
						},
						"async_message_processing": schema.BoolAttribute{
							MarkdownDescription: "<p>If enabled messages taken from the get queue will be processed not necessarily in the same order as they were queued.</p><p>This property may be set to improve performance only if associated Multi-Protocol Gateway or WS-Proxy isn't configured to process messages in order.</p>",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *WebSphereJMSSourceProtocolHandlerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *WebSphereJMSSourceProtocolHandlerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data WebSphereJMSSourceProtocolHandlerList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.WebSphereJMSSourceProtocolHandler{
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
	l := []models.WebSphereJMSSourceProtocolHandler{}
	if resFound {
		if value := res.Get(`WebSphereJMSSourceProtocolHandler`); value.Exists() {
			for _, v := range value.Array() {
				item := models.WebSphereJMSSourceProtocolHandler{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.WebSphereJMSSourceProtocolHandlerObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.WebSphereJMSSourceProtocolHandlerObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
