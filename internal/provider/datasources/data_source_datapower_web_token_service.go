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

type WebTokenServiceList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &WebTokenServiceDataSource{}
	_ datasource.DataSourceWithConfigure = &WebTokenServiceDataSource{}
)

func NewWebTokenServiceDataSource() datasource.DataSource {
	return &WebTokenServiceDataSource{}
}

type WebTokenServiceDataSource struct {
	pData *tfutils.ProviderData
}

func (d *WebTokenServiceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_web_token_service"
}

func (d *WebTokenServiceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The web token service provides on-box HTTP services.",
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
							MarkdownDescription: "Specify a brief, descriptive comment.",
							Computed:            true,
						},
						"priority": schema.StringAttribute{
							MarkdownDescription: "Set the service-level priority. When resources are in high demand, the appliance handles transactions for higher priority services over lower priority services.",
							Computed:            true,
						},
						"xml_manager": schema.StringAttribute{
							MarkdownDescription: "Assign the XML manager that obtains and manages documents on behalf of the service. The user agent, which is referenced by this XML manager, controls access to remote destinations. Unless you have service-specific needs, use the <tt>default</tt> XML manager.",
							Computed:            true,
						},
						"request_type": schema.StringAttribute{
							MarkdownDescription: "Characterizes the traffic that originates from the client. The default is Non-XML.",
							Computed:            true,
						},
						"front_side": schema.ListNestedAttribute{
							MarkdownDescription: "Define the addresses and ports that the service listens on. You can define listening sources for HTTP and HTTPS requests. For HTTPS request, set the listening source to use TLS parameter and assign the TLS profile.",
							NestedObject:        models.DmSSLFrontSideDataSourceSchema,
							Computed:            true,
						},
						"style_policy": schema.StringAttribute{
							MarkdownDescription: "Assign the processing policy to the service. The processing policy defines the actions to perform on security token requests.",
							Computed:            true,
						},
						"rewrite_errors": schema.BoolAttribute{
							MarkdownDescription: "Error messages after a decryption action can provide an attacker who is using the padding oracle attack method with enough information to determine the contents of the plain-text data. When enabled, the default, the client receives error messages without the internal information that could lead to a discovery. When disabled, the client receives the original message with this information.",
							Computed:            true,
						},
						"delay_errors": schema.BoolAttribute{
							MarkdownDescription: "The timing difference of the error messages returned after a decryption action can provide an attacker with enough information to determine the contents of the plain-text data. When enabled, the default, the appliance delays error messages for the defined duration. When disabled, the appliance does not delay error messages.",
							Computed:            true,
						},
						"delay_errors_duration": schema.Int64Attribute{
							MarkdownDescription: "When enabling the delay of error messages, specify the delay duration in milliseconds. If delaying messages for 3000 ms, the appliance will not send error messages to the client until 3 seconds have elapsed since the appliance performed decryption on the requests. Enter any value in the range 250 - 300000. The default value is 1000.",
							Computed:            true,
						},
						"front_timeout": schema.Int64Attribute{
							MarkdownDescription: "Set the intra-transaction timeout for Web Token Service to client connections. This value is the maximum idle time to allow in a transaction on the Web Token Service to client connection. This timer monitors idle time in the data transfer process. If the specified idle time is exceeded, the connection is torn down. Enter a value in the range 1 - 86400. The default value is 120.",
							Computed:            true,
						},
						"front_persistent_timeout": schema.Int64Attribute{
							MarkdownDescription: "Set the inter-transaction timeout for Web Token Service to client connections. This value is the maximum idle time to allow between the completion of a TCP transaction and the initiation of a new TCP transaction on the Web Token Service to client connection. If the specified idle time is exceeded, the connection is torn down. Enter a value in the range 0 - 86400. The default value is 180. A value of 0 disables persistent connections.",
							Computed:            true,
						},
						"front_http_version": schema.StringAttribute{
							MarkdownDescription: "Set the HTTP version for client-to-service connections. If the client submits an HTTP/1.0 request, the service on the DataPower appliance always replies with HTTP/1.0 compatible responses regardless of this setting. The default version is HTTP 1.1.",
							Computed:            true,
						},
						"http_client_ip_label": schema.StringAttribute{
							MarkdownDescription: "Set the HTTP header that contains the client IP address. When defined, the client IP address is read from this HTTP header. This IP address will then be used for monitoring and logging. Retain the default value <tt>X-Client-IP</tt> or specify another value; for example, <tt>X-Forwarded-For</tt> .",
							Computed:            true,
						},
						"http_log_cor_id_label": schema.StringAttribute{
							MarkdownDescription: "Enter the name of an HTTP Header to read to determine the global transaction ID for chained services. This value defaults to X-Global-Transaction-ID.",
							Computed:            true,
						},
						"debug_mode": schema.StringAttribute{
							MarkdownDescription: "<p>Select the diagnostic mode for processing policies. When enabled, you can view details about the state of variables and contexts for a captured transaction in the probe. The default value is <tt>off</tt> .</p><p>Transaction diagnostic mode is not intended for use in a production environment. Transaction diagnostic mode consumes significant resources that can slow down transaction processing.</p>",
							Computed:            true,
						},
						"debug_history": schema.Int64Attribute{
							MarkdownDescription: "Set the number of records for transaction diagnostic mode in the probe. Enter a value in the range 10 - 250. The default value is 25.",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *WebTokenServiceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *WebTokenServiceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data WebTokenServiceList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.WebTokenService{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.WebTokenService{}
	if value := res.Get(`WebTokenService`); value.Exists() {
		for _, v := range value.Array() {
			item := models.WebTokenService{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.WebTokenServiceObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.WebTokenServiceObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
