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

type SMTPServerConnectionList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &SMTPServerConnectionDataSource{}
	_ datasource.DataSourceWithConfigure = &SMTPServerConnectionDataSource{}
)

func NewSMTPServerConnectionDataSource() datasource.DataSource {
	return &SMTPServerConnectionDataSource{}
}

type SMTPServerConnectionDataSource struct {
	pData *tfutils.ProviderData
}

func (d *SMTPServerConnectionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_smtp_server_connection"
}

func (d *SMTPServerConnectionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "<p>The SMTP server connection defines the connection details for a Simple Mail Transport Protocol (SMTP) server. The DataPower Gateway uses the SMTP server connection for the following purposes</p><ul><li>B2B partners use this configuration to send an e-mail message to an AS1 or ESMTP destination.</li><li>B2B gateways use this configuration to request an AS1 MDN.</li></ul><p>For ease of configuration, the DataPower Gateway provides the <tt>default</tt> SMTP server connection configuration in each domain. By default, this configuration is empty and disabled.</p>",
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
						"mail_server_host": schema.StringAttribute{
							MarkdownDescription: "The IP address or host name of the SMTP server to send outgoing e-mail messages.",
							Computed:            true,
						},
						"mail_server_port": schema.Int64Attribute{
							MarkdownDescription: "The listening port on the SMTP server to send outgoing e-mail messages.",
							Computed:            true,
						},
						"options": models.GetDmSMTPOptionsDataSourceSchema("The SMTP options to enable for the SMTP client. If blank, the configuration uses the setting from the SMTP client policy in the associated user agent.", "options", ""),
						"auth": schema.StringAttribute{
							MarkdownDescription: "With the client authentication option, the method to authenticate the SMTP client. If blank, the configuration uses the setting from the SMTP client policy in the associated user agent.",
							Computed:            true,
						},
						"account_name": schema.StringAttribute{
							MarkdownDescription: "The account or user name of the SMTP client to authenticate on the SMTP server. The account generally takes the <tt>name@domain.com</tt> form. If blank, the configuration uses the setting from the basic authentication policy in the associated user agent.",
							Computed:            true,
						},
						"account_password_alias": schema.StringAttribute{
							MarkdownDescription: "The password alias of the password for the SMTP client account or the user name that is authenticated to the SMTP server. If password or alias are blank, the configuration uses the setting from the basic authentication policy in the associated user agent.",
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

func (d *SMTPServerConnectionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *SMTPServerConnectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data SMTPServerConnectionList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.SMTPServerConnection{
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
	l := []models.SMTPServerConnection{}
	if resFound {
		if value := res.Get(`SMTPServerConnection`); value.Exists() {
			for _, v := range value.Array() {
				item := models.SMTPServerConnection{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.SMTPServerConnectionObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.SMTPServerConnectionObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
