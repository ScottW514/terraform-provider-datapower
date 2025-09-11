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

type POPPollerSourceProtocolHandlerList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &POPPollerSourceProtocolHandlerDataSource{}
	_ datasource.DataSourceWithConfigure = &POPPollerSourceProtocolHandlerDataSource{}
)

func NewPOPPollerSourceProtocolHandlerDataSource() datasource.DataSource {
	return &POPPollerSourceProtocolHandlerDataSource{}
}

type POPPollerSourceProtocolHandlerDataSource struct {
	pData *tfutils.ProviderData
}

func (d *POPPollerSourceProtocolHandlerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_pop_poller_source_protocol_handler"
}

func (d *POPPollerSourceProtocolHandlerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The Post Office Protocol (POP) handler manages the polling of a mailbox on a mail server. The mailbox receives mail messages from external partners. The handler retrieves and deletes mail messages on each polling cycle. Each mail message that The handler retrieves results in one transaction.",
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
						"mail_server": schema.StringAttribute{
							MarkdownDescription: "The host name or IP address of the mail server.",
							Computed:            true,
						},
						"port": schema.Int64Attribute{
							MarkdownDescription: "The listening port on the mail server. STARTTLS negotiation and an unsecured connection generally use port 110. An implicit, secured connection generally uses port 995.",
							Computed:            true,
						},
						"conn_security": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"auth_method": schema.StringAttribute{
							MarkdownDescription: "The type of authentication to use. If authentication fails, no connection is made.",
							Computed:            true,
						},
						"account": schema.StringAttribute{
							MarkdownDescription: "The name to access the mailbox on the server; for example, user@example.com.",
							Computed:            true,
						},
						"password_alias": schema.StringAttribute{
							MarkdownDescription: "The password alias of the password for the account that accesses the mailbox on the server.",
							Computed:            true,
						},
						"delay_between_polls": schema.Int64Attribute{
							MarkdownDescription: "Specify the interval in seconds between polling sequences. A <em>polling sequence</em> is the time to retrieve the messages plus the time to complete their processing. Enter a value in the range 1 - 65535. The default value is 300. <p><b>Note:</b> Some mail servers restrict the number of times an account can establish a connection during a specific time period. Ensure that the configured interval complies with any restriction.</p>",
							Computed:            true,
						},
						"max_messages_per_poll": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum number of messages to retrieve in each polling cycle. Enter a value in the range 1 - 100. The default value is 5.",
							Computed:            true,
						},
						"ssl_client_config_type": schema.StringAttribute{
							MarkdownDescription: "The TLS profile type to secure connections between the DataPower Gateway and its targets.",
							Computed:            true,
						},
						"ssl_client": schema.StringAttribute{
							MarkdownDescription: "The TLS client profile to secure connections between the DataPower Gateway and its targets.",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *POPPollerSourceProtocolHandlerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *POPPollerSourceProtocolHandlerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data POPPollerSourceProtocolHandlerList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.POPPollerSourceProtocolHandler{
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
	l := []models.POPPollerSourceProtocolHandler{}
	if resFound {
		if value := res.Get(`POPPollerSourceProtocolHandler`); value.Exists() {
			for _, v := range value.Array() {
				item := models.POPPollerSourceProtocolHandler{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.POPPollerSourceProtocolHandlerObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.POPPollerSourceProtocolHandlerObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
