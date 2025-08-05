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
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
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
	client *client.DatapowerClient
}

func (d *MQv9PlusSourceProtocolHandlerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mqv9plussourceprotocolhandler"
}

func (d *MQv9PlusSourceProtocolHandlerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "IBM MQ v9+ handler",
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
							MarkdownDescription: "Queue manager (reference to MQManger or MQManagerGroup)",
							Computed:            true,
						},
						"get_queue": schema.StringAttribute{
							MarkdownDescription: "Get queue",
							Computed:            true,
						},
						"subscribe_topic_string": schema.StringAttribute{
							MarkdownDescription: "Subscribe topic string",
							Computed:            true,
						},
						"subscription_name": schema.StringAttribute{
							MarkdownDescription: "Subscription name",
							Computed:            true,
						},
						"put_queue": schema.StringAttribute{
							MarkdownDescription: "Put queue",
							Computed:            true,
						},
						"publish_topic_string": schema.StringAttribute{
							MarkdownDescription: "Publish topic string",
							Computed:            true,
						},
						"code_page": schema.Int64Attribute{
							MarkdownDescription: "CCSI",
							Computed:            true,
						},
						"get_message_options": schema.Int64Attribute{
							MarkdownDescription: "Get message options",
							Computed:            true,
						},
						"message_selector": schema.StringAttribute{
							MarkdownDescription: "Selector",
							Computed:            true,
						},
						"parse_properties": schema.BoolAttribute{
							MarkdownDescription: "Parse properties",
							Computed:            true,
						},
						"async_put": schema.BoolAttribute{
							MarkdownDescription: "Async put",
							Computed:            true,
						},
						"exclude_headers": models.GetDmMQHeadersDataSourceSchema("Exclude message headers", "exclude-headers", ""),
						"concurrent_connections": schema.Int64Attribute{
							MarkdownDescription: "Concurrent conversations",
							Computed:            true,
						},
						"polling_interval": schema.Int64Attribute{
							MarkdownDescription: "Polling interval",
							Computed:            true,
						},
						"batch_size": schema.Int64Attribute{
							MarkdownDescription: "Batch size",
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
							MarkdownDescription: "Retrieve backout settings",
							Computed:            true,
						},
						"use_qm_name_in_url": schema.BoolAttribute{
							MarkdownDescription: "Use queue manager in URL",
							Computed:            true,
						},
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

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *MQv9PlusSourceProtocolHandlerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data MQv9PlusSourceProtocolHandlerList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.MQv9PlusSourceProtocolHandler{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
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
