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

type AssemblyActionWebSocketUpgradeList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &AssemblyActionWebSocketUpgradeDataSource{}
	_ datasource.DataSourceWithConfigure = &AssemblyActionWebSocketUpgradeDataSource{}
)

func NewAssemblyActionWebSocketUpgradeDataSource() datasource.DataSource {
	return &AssemblyActionWebSocketUpgradeDataSource{}
}

type AssemblyActionWebSocketUpgradeDataSource struct {
	pData *tfutils.ProviderData
}

func (d *AssemblyActionWebSocketUpgradeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_assembly_action_web_socket_upgrade"
}

func (d *AssemblyActionWebSocketUpgradeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The WebSocket upgrade assembly action processes API requests and responses through a WebSocket connection.",
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
						"url": schema.StringAttribute{
							MarkdownDescription: "Specify the URL to invoke. You can reference a custom API property that resolves as the value. To reference an API property, use the <tt>$(api.properties. <i>property_name</i> )</tt> format, where <tt><i>property_name</i></tt> is the name of the property to reference. You can use the short <tt>$( <i>property_name</i> )</tt> format when the assembly action does not have a property with the same name.",
							Computed:            true,
						},
						"ssl_client": schema.StringAttribute{
							MarkdownDescription: "TLS client profile",
							Computed:            true,
						},
						"timeout": schema.Int64Attribute{
							MarkdownDescription: "Specify the duration in seconds to wait for a reply from the target. The default value is 60.",
							Computed:            true,
						},
						"user_name": schema.StringAttribute{
							MarkdownDescription: "User name",
							Computed:            true,
						},
						"password": schema.StringAttribute{
							MarkdownDescription: "Password alias",
							Computed:            true,
						},
						"follow_redirects": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to follow the redirects. When enabled, attempts to resolve any redirect transparently.",
							Computed:            true,
						},
						"decode_request_params": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to decode request parameters in the target URL. When enabled, request parameters are decoded. The default behavior is to not decode request parameters.",
							Computed:            true,
						},
						"encode_plus_char": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to encode + characters in query strings. When enabled, + characters are encoded to <tt>%2F</tt> . The default behavior is to not encode + characters.",
							Computed:            true,
						},
						"inject_user_agent_header": schema.BoolAttribute{
							MarkdownDescription: "Inject User-Agent header",
							Computed:            true,
						},
						"inject_proxy_headers": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to inject proxy-related headers when not found in the request. The proxy-related headers are <tt>X-Forwarded-For</tt> , <tt>X-Forwarded-Host</tt> , and <tt>X-Forwarded-Port</tt> .",
							Computed:            true,
						},
						"header_control_list": schema.StringAttribute{
							MarkdownDescription: "Specify the control list that manages whether to accept or reject headers. The default behavior is to accept all headers.",
							Computed:            true,
						},
						"parameter_control_list": schema.StringAttribute{
							MarkdownDescription: "Specify the control list that manages whether to accept or reject URL parameters. The default behavior is to reject all URL parameters.",
							Computed:            true,
						},
						"api_request_processing_assembly": schema.StringAttribute{
							MarkdownDescription: "API request processing assembly",
							Computed:            true,
						},
						"api_response_processing_assembly": schema.StringAttribute{
							MarkdownDescription: "API response processing assembly",
							Computed:            true,
						},
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"title": schema.StringAttribute{
							MarkdownDescription: "Title",
							Computed:            true,
						},
						"correlation_path": schema.StringAttribute{
							MarkdownDescription: "Specify the path that correlates the API action to a specific part of the API specification. The correlation path specifies the part of the API definition that correlates with the API action. This path is exposed in the debug data by the API gateway for use by debugging tools. For example, for an API configuration that is retrieved from API Connect and specified in an OpenAPI document with IBM extensions, this path is the JSON path to the assembly policy in the IBM extensions section of the document. The path can be expressed in any form that the debugging tool can correlate to the API definition.",
							Computed:            true,
						},
						"action_debug": schema.BoolAttribute{
							MarkdownDescription: "<p>Specify whether to enable the GatewayScript debugger to troubleshoot the following GatewayScript files or script.</p><ul><li>Troubleshoot a GatewayScript file that is called from the GatewayScript assembly action.</li><li>Troubleshoot a GatewayScript file that is called from an XSLT assembly action that uses the <tt>gatewayscript()</tt> extension function.</li><li>Troubleshoot a GatewayScript script that is called through the <tt>value</tt> or <tt>default</tt> property in the JSON file from the map assembly action.</li></ul><p>To debug a file or script, the following conditions must be met.</p><ul><li>The file contains one or more <tt>debugger;</tt> statements at the points in your script where you want to start debugging.</li><li>The GatewayScript debugger is enabled.</li></ul><p>You run the <tt>debug-action</tt> command.</p>",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *AssemblyActionWebSocketUpgradeDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *AssemblyActionWebSocketUpgradeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data AssemblyActionWebSocketUpgradeList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.AssemblyActionWebSocketUpgrade{
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
	l := []models.AssemblyActionWebSocketUpgrade{}
	if resFound {
		if value := res.Get(`AssemblyActionWebSocketUpgrade`); value.Exists() {
			for _, v := range value.Array() {
				item := models.AssemblyActionWebSocketUpgrade{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.AssemblyActionWebSocketUpgradeObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.AssemblyActionWebSocketUpgradeObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
