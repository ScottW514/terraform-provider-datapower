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

type AssemblyActionInvokeList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &AssemblyActionInvokeDataSource{}
	_ datasource.DataSourceWithConfigure = &AssemblyActionInvokeDataSource{}
)

func NewAssemblyActionInvokeDataSource() datasource.DataSource {
	return &AssemblyActionInvokeDataSource{}
}

type AssemblyActionInvokeDataSource struct {
	pData *tfutils.ProviderData
}

func (d *AssemblyActionInvokeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_assembly_action_invoke"
}

func (d *AssemblyActionInvokeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The invoke assembly action call a service from your assembly.",
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
							MarkdownDescription: "Specify the URL of the target. You can reference a custom API property that resolves as the value. To reference an API property, use the <tt>$(api.properties. <i>property_name</i> )</tt> format, where <tt><i>property_name</i></tt> is the name of the property to reference. You can use the short <tt>$( <i>property_name</i> )</tt> format when the assembly action does not have a property with the same name.",
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
							MarkdownDescription: "Username",
							Computed:            true,
						},
						"password": schema.StringAttribute{
							MarkdownDescription: "Password",
							Computed:            true,
						},
						"method": schema.StringAttribute{
							MarkdownDescription: "HTTP method",
							Computed:            true,
						},
						"backend_type": schema.StringAttribute{
							MarkdownDescription: "Backend type",
							Computed:            true,
						},
						"graphql_send_type": schema.StringAttribute{
							MarkdownDescription: "Specify the type of payload to send for GraphQL POST requests. When GraphQL or JSON, this setting overrides the message type of the payload.",
							Computed:            true,
						},
						"compression": schema.BoolAttribute{
							MarkdownDescription: "Enable compression",
							Computed:            true,
						},
						"cache_type": schema.StringAttribute{
							MarkdownDescription: "Specify how to cache documents.",
							Computed:            true,
						},
						"time_to_live": schema.Int64Attribute{
							MarkdownDescription: "Specify the validity period in seconds for documents in the cache. The default value is 900.",
							Computed:            true,
						},
						"cache_unsafe_response": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to cache responses to POST and PUT requests when the cache policy type is set to time to live. The response to these requests is the result of an action on the server that might change its resource state. You might want to cache responses to these requests when you know that the action (for example: HTTP POST) will not change the server state.",
							Computed:            true,
						},
						"cache_key": schema.StringAttribute{
							MarkdownDescription: "Specify the string for the cache key. If omitted, the entire URL is used as the key.",
							Computed:            true,
						},
						"follow_redirects": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to follow redirects. Some protocols generate redirects. When enabled, the action attempts to resolve redirects transparently.",
							Computed:            true,
						},
						"http_version": schema.StringAttribute{
							MarkdownDescription: "Specify the HTTP version for server-side connections. The default value is HTTP/1.1.",
							Computed:            true,
						},
						"http2_required": schema.BoolAttribute{
							MarkdownDescription: "Specify whether an HTTP/2 connection is required when connecting to the server. Only applicable when the HTTP version to the server is set to HTTP/2 and the connection uses TLS. The default value is off.",
							Computed:            true,
						},
						"do_chunked_upload": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enable uploading of HTTP/1.1 chunked-encoded documents. For HTTP/1.1, the document body can be delimited by either <tt>Content-Length</tt> or chunked encoding. While all servers understand <tt>Content-Length</tt> , many servers fail to understand chunked encoding. For this reason, <tt>Content-Length</tt> is the standard method. However, the use of <tt>Content-Length</tt> can interfere with streaming. To stream full documents to an RFC 2616 compatible server, enable this property. Unlike other HTTP/1.1 features, you must know that the target server is RFC 2616 compatible.",
							Computed:            true,
						},
						"persistent_connection": schema.BoolAttribute{
							MarkdownDescription: "Persistent connection",
							Computed:            true,
						},
						"stop_on_error": schema.BoolAttribute{
							MarkdownDescription: "Stop on error",
							Computed:            true,
						},
						"error_types": models.GetDmInvokeErrorTypeDataSourceSchema("Error types", "error-types", ""),
						"output": schema.StringAttribute{
							MarkdownDescription: "Specify the variable to store results. By default, results are stored in the <tt>message.body</tt> , <tt>message.headers</tt> , <tt>message.statuscode</tt> variables.",
							Computed:            true,
						},
						"decode_request_params": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to decode the request parameters in the target URL. When enabled, request parameters are decoded. By default, request parameters are not decoded.",
							Computed:            true,
						},
						"encode_plus_char": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to encode + characters in query strings. When enabled, + characters are encoded to <tt>%2F</tt> . By default, + characters are not encoded.",
							Computed:            true,
						},
						"keep_payload": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to include the payload for DELETE requests. When enabled, DELETE requests include the payload. By default, DELETE requests do not include the payload.",
							Computed:            true,
						},
						"inject_user_agent_header": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to inject the default <tt>User-Agent</tt> header. When the <tt>User-Agent</tt> header is not in the request, inject this header to the request.",
							Computed:            true,
						},
						"inject_proxy_headers": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to inject proxy-related headers. When the <tt>X-Forwarded-For</tt> , <tt>X-Forwarded-Host</tt> , and <tt>X-Forwarded-Port</tt> headers are not found in the request, inject theses headers to the request.",
							Computed:            true,
						},
						"header_control_list": schema.StringAttribute{
							MarkdownDescription: "Specify the control list that uses headers to accept or reject requests. By default, accepts all requests with headers.",
							Computed:            true,
						},
						"parameter_control_list": schema.StringAttribute{
							MarkdownDescription: "Specify the control list that uses URL parameters to accept or reject requests. By default, rejects all requests with URL parameters.",
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
					},
				},
			},
		},
	}
}

func (d *AssemblyActionInvokeDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *AssemblyActionInvokeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data AssemblyActionInvokeList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.AssemblyActionInvoke{
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
	l := []models.AssemblyActionInvoke{}
	if resFound {
		if value := res.Get(`AssemblyActionInvoke`); value.Exists() {
			for _, v := range value.Array() {
				item := models.AssemblyActionInvoke{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.AssemblyActionInvokeObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.AssemblyActionInvokeObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
