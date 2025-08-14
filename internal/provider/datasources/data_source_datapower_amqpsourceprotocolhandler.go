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

type AMQPSourceProtocolHandlerList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &AMQPSourceProtocolHandlerDataSource{}
	_ datasource.DataSourceWithConfigure = &AMQPSourceProtocolHandlerDataSource{}
)

func NewAMQPSourceProtocolHandlerDataSource() datasource.DataSource {
	return &AMQPSourceProtocolHandlerDataSource{}
}

type AMQPSourceProtocolHandlerDataSource struct {
	pData *tfutils.ProviderData
}

func (d *AMQPSourceProtocolHandlerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_amqpsourceprotocolhandler"
}

func (d *AMQPSourceProtocolHandlerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "In AMQP, distributed source and target termini are managed by a broker. The broker periodically monitors and polls termini. The broker ensures that sent messages are directed to the correct target terminus or are routed to another server. The AMQP broker configuration corresponds to an AMQP broker that is running on another host in the network.",
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
						"broker": schema.StringAttribute{
							MarkdownDescription: "AMQP broker",
							Computed:            true,
						},
						"from": schema.StringAttribute{
							MarkdownDescription: "Source terminus",
							Computed:            true,
						},
						"to": schema.StringAttribute{
							MarkdownDescription: "Target terminus",
							Computed:            true,
						},
						"credit": schema.Int64Attribute{
							MarkdownDescription: "Specify the number of concurrent messages that a receiver can handle. The minimum value is 1. The default value is 100.",
							Computed:            true,
						},
						"ignore_reply_to": schema.BoolAttribute{
							MarkdownDescription: "<p>Specify whether to ignore the AMQP <tt>reply-to</tt> property. The default behavior is to ignore the property. <ul><li>When enabled, ignore the <tt>reply-to</tt> address when sending an AMQP response message.</li><li>When disabled, use the <tt>reply-to</tt> address instead of the address of the target terminus.</li></ul></p>",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *AMQPSourceProtocolHandlerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *AMQPSourceProtocolHandlerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data AMQPSourceProtocolHandlerList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.AMQPSourceProtocolHandler{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.AMQPSourceProtocolHandler{}
	if value := res.Get(`AMQPSourceProtocolHandler`); value.Exists() {
		for _, v := range value.Array() {
			item := models.AMQPSourceProtocolHandler{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.AMQPSourceProtocolHandlerObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.AMQPSourceProtocolHandlerObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
