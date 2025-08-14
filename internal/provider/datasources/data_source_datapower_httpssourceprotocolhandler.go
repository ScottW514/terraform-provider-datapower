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

type HTTPSSourceProtocolHandlerList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &HTTPSSourceProtocolHandlerDataSource{}
	_ datasource.DataSourceWithConfigure = &HTTPSSourceProtocolHandlerDataSource{}
)

func NewHTTPSSourceProtocolHandlerDataSource() datasource.DataSource {
	return &HTTPSSourceProtocolHandlerDataSource{}
}

type HTTPSSourceProtocolHandlerDataSource struct {
	pData *tfutils.ProviderData
}

func (d *HTTPSSourceProtocolHandlerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_httpssourceprotocolhandler"
}

func (d *HTTPSSourceProtocolHandlerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "An HTTPS handler receives HTTP requests over TLS and forwards them to the appropriate DataPower service. HTTPS handlers conform to RFC 2616.",
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
						"local_address": schema.StringAttribute{
							MarkdownDescription: "Specify the IP address or host alias that the handler listens. The default value indicates that The handler listens on all IPv4 addresses.",
							Computed:            true,
						},
						"local_port": schema.Int64Attribute{
							MarkdownDescription: "Port",
							Computed:            true,
						},
						"http_version": schema.StringAttribute{
							MarkdownDescription: "Specify the HTTP version for client connections. The default value is HTTP/1.1. For the HTTP/2 protocol, requests and responses are always HTTP/2. When HTTP/2, this setting is ignored.",
							Computed:            true,
						},
						"allowed_features": models.GetDmSourceHTTPFeatureTypeDataSourceSchema("Allowed methods and versions", "allowed-features", ""),
						"persistent_connections": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to negotiate persistent connections with clients. The HTTP/2 protocol controls persistent connections and reuse. Therefore, this setting is ignored for the HTTP/2 protocol.",
							Computed:            true,
						},
						"max_persistent_connections_reuse": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum number of times that a client can reuse a persistent connection. When this count is reached, an explicit <tt>HTTP Connection: close</tt> header is sent in the response, and the TCP connection is closed. The default value is 0, which means unlimited reuse.",
							Computed:            true,
						},
						"allow_compression": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to negotiate GZIP compression for client connections. When enabled and the <tt>Accept-Encoding</tt> HTTP header indicates that compressed documents can be processed, the service uses GZIP to compress HTTP transmissions. The <tt>Transfer-Encoding</tt> HTTP header indicates compression.",
							Computed:            true,
						},
						"allow_web_socket_upgrade": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to allow WebSocket upgrade requests from clients. The default value is disabled. This request is to switch the existing connection to use the WebSocket protocol. WebSocket upgrade requests require that The handler allows GET methods.",
							Computed:            true,
						},
						"web_socket_idle_timeout": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum idle time in seconds for client connections. This timer monitors the idle time in the data transfer process. When the specified idle time is exceeded, the connection is torn down. Enter a value in the range 0 - 86400. The default value is 0, which indicates that the timer is disabled.",
							Computed:            true,
						},
						"max_url_len": schema.Int64Attribute{
							MarkdownDescription: "Specify the length in bytes of the longest incoming URL to accept. The length includes any query string or fragment identifier. Enter a value in the range 1 - 128000. The default value is 16384.",
							Computed:            true,
						},
						"max_total_hdr_len": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum aggregate length in bytes of HTTP headers to allow. Enter a value in the range 5 - 128000. The default value is 128000.",
							Computed:            true,
						},
						"max_hdr_count": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum number of headers to allow in client requests. The default value is 0, which indicates no limit.",
							Computed:            true,
						},
						"max_name_hdr_len": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum length in bytes of a header name to allow in client requests. Each HTTP header is expressed as a name-value pair. This setting sets the maximum length of the name portion of a header. The default value is 0, which indicates no limit.",
							Computed:            true,
						},
						"max_value_hdr_len": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum length of a header value in bytes to allow in client requests. Each HTTP header is expressed as a name-value pair. This setting sets the maximum length of the value portion of a header. The default value is 0, which indicates no limit.",
							Computed:            true,
						},
						"max_query_string_len": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum length of the query string in bytes to allow in client requests. The query string is the portion of the URL after the ? character. The default value is 0, which indicates no limit.",
							Computed:            true,
						},
						"acl": schema.StringAttribute{
							MarkdownDescription: "Access control list",
							Computed:            true,
						},
						"credential_charset": schema.StringAttribute{
							MarkdownDescription: "Specify the character encoding of the original basic authentication values. Basic authentication credentials are combined and base64 encoded in the authorization header of the request. The contents of the <tt>Authorization</tt> header are transcoded to UTF-8. The default value represents ISO-8859-1 Latin 1.",
							Computed:            true,
						},
						"ssl_server_config_type": schema.StringAttribute{
							MarkdownDescription: "TLS server type",
							Computed:            true,
						},
						"ssl_server": schema.StringAttribute{
							MarkdownDescription: "Specify the TLS server profile to secure connections from clients. <b>Note:</b> The TLS server profile that secures the HTTP/2 connection must use the TLS 1.2 or later protocol with a cipher that is secure according to RFC 7540.",
							Computed:            true,
						},
						"sslsni_server": schema.StringAttribute{
							MarkdownDescription: "Specify the TLS SNI server profile to secure connections from clients. <b>Note:</b> The TLS SNI server profile that secures the HTTP/2 connection must use the TLS 1.2 or later protocol with a cipher that is secure according to RFC 7540.",
							Computed:            true,
						},
						"http2_max_streams": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum number of concurrent streams that the client can have outstanding at the same time. Enter a value in the range 1 - 500. The default value is 100. <p>The limit applies to the number of streams that the client allows the target to create. The greater the number of streams in use, the more resources the client uses. Resources include memory and the network connections to the destination.</p>",
							Computed:            true,
						},
						"http2_max_frame_size": schema.Int64Attribute{
							MarkdownDescription: "Specify the largest payload frame size that the client can send. Enter a value in the range 16384 - 16777215. The default value is 16384.",
							Computed:            true,
						},
						"http2_stream_header": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enable the HTTP/2 stream identifier header in the request or response. When enabled, the HTTP/2 stream identifier is included in the <tt>X-DP-http2-stream</tt> header. With this header, you can correlate the HTTP/2 stream. The default behavior is disabled.",
							Computed:            true,
						},
						"chunked_encoding": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enable responses to use chunked transfer-encoding. By default, HTTP responses use <tt>Transfer-Encoding: chunked</tt> .",
							Computed:            true,
						},
						"header_timeout": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum duration in milliseconds to allow for request headers processing. When the value is greater than 0, request header processing must complete before the duration elapses. Enter a value in the range 0 - 3600000, where a value of 0 disables the timer. The default value is 30000.",
							Computed:            true,
						},
						"http2_idle_timeout": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum idle duration in milliseconds to allow before closing the HTTP/2 connection. Enter a value in the range 0 - 3600000, where a value of 0 disables the timer. The default value is 0.",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *HTTPSSourceProtocolHandlerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *HTTPSSourceProtocolHandlerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data HTTPSSourceProtocolHandlerList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.HTTPSSourceProtocolHandler{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.HTTPSSourceProtocolHandler{}
	if value := res.Get(`HTTPSSourceProtocolHandler`); value.Exists() {
		for _, v := range value.Array() {
			item := models.HTTPSSourceProtocolHandler{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.HTTPSSourceProtocolHandlerObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.HTTPSSourceProtocolHandlerObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
