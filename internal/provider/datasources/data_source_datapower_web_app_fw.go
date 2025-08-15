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

type WebAppFWList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &WebAppFWDataSource{}
	_ datasource.DataSourceWithConfigure = &WebAppFWDataSource{}
)

func NewWebAppFWDataSource() datasource.DataSource {
	return &WebAppFWDataSource{}
}

type WebAppFWDataSource struct {
	pData *tfutils.ProviderData
}

func (d *WebAppFWDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_web_app_fw"
}

func (d *WebAppFWDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The web application firewall provides filtering, security, and input validation for HTTP transactions.",
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
						"priority": schema.StringAttribute{
							MarkdownDescription: "Control the service scheduling priority. When system resources are in high demand, \"high\" priority services will be favored over lower priority services.",
							Computed:            true,
						},
						"front_side": schema.ListNestedAttribute{
							MarkdownDescription: "Source Addresses",
							NestedObject:        models.DmFrontSideDataSourceSchema,
							Computed:            true,
						},
						"remote_address": schema.StringAttribute{
							MarkdownDescription: "Specify the host name or IP address of the backend server. If the backend requires TLS, ensure that the TLS client profile is configured for client (forward) connections.",
							Computed:            true,
						},
						"remote_port": schema.Int64Attribute{
							MarkdownDescription: "Specify the port on the remote server.",
							Computed:            true,
						},
						"style_policy": schema.StringAttribute{
							MarkdownDescription: "Select a security policy from the list of available policies. Click + to create a new Policy. This policy controls which profiles are enforced.",
							Computed:            true,
						},
						"xml_manager": schema.StringAttribute{
							MarkdownDescription: "<p>An XML Manager manages the compilation and caching of stylesheets, the caching of documents, and provides configuration constraints on the size and parsing depth of documents. You can enable streaming operation by configuring an XML Manager to use a Streaming Compile Option Policy.</p><p>Note that an XML Manager can optionally employ a User Agent. The User Agent settings, in turn, can affect the behavior of the gateway when communicating with back end servers or with clients when sending back responses.</p><p>More than one firewall may use the same XML Manager. Select an existing XML Manager from the list to assign the Manager to this firewall. Click the + button to create a new XML Manager that is assigned to this firewall. A default Manager is used if you do not create one.</p>",
							Computed:            true,
						},
						"error_policy": schema.StringAttribute{
							MarkdownDescription: "If any policy is violated, this is the default policy that will handle the response sent to the client. It might be overridden in the profile that fails.",
							Computed:            true,
						},
						"uri_normalization": schema.BoolAttribute{
							MarkdownDescription: "If this property is enabled the URI is rewritten to make sure the URI is RFC compliant by escaping certain characters. Additionally, characters that are escaped that do not need to be are unescaped. This makes checking for attack sequences such as .. more reliable.",
							Computed:            true,
						},
						"rewrite_errors": schema.BoolAttribute{
							MarkdownDescription: "Error messages after a decryption action can provide an attacker who is using the padding oracle attack method with enough information to determine the contents of the plain-text data. When enabled, the default, the client receives error messages without the internal information that could lead to a discovery. When disabled, the client receives the original message with this information.",
							Computed:            true,
						},
						"delay_errors": schema.BoolAttribute{
							MarkdownDescription: "The timing difference of the error messages returned after a decryption action can provide an attacker with enough information to determine the contents of the plain-text data. When enabled, the default, the DataPower Gateway delays error messages for the defined duration. When disabled, the DataPower Gateway does not delay error messages.",
							Computed:            true,
						},
						"delay_errors_duration": schema.Int64Attribute{
							MarkdownDescription: "When enabling the delay of error messages, specify the delay duration in milliseconds. If delaying messages for 3000ms, the DataPower Gateway will not send error messages to the client until 3 seconds have elapsed since the DataPower Gateway performed decryption on the requests. Use any value of 250 - 300000. The default value is 1000.",
							Computed:            true,
						},
						"stream_output_to_back": schema.StringAttribute{
							MarkdownDescription: "Select the desired streaming behavior.",
							Computed:            true,
						},
						"stream_output_to_front": schema.StringAttribute{
							MarkdownDescription: "Select the desired streaming behavior.",
							Computed:            true,
						},
						"front_timeout": schema.Int64Attribute{
							MarkdownDescription: "Set the intra-transaction timeout for Web Application Firewall to client connections. This value is the maximum idle time to allow in a transaction on the Web Application Firewall to client connection. This timer monitors idle time in the data transfer process. If the specified idle time is exceeded, the connection is torn down. Enter a value in the range 1 - 86400. The default value is 120.",
							Computed:            true,
						},
						"back_timeout": schema.Int64Attribute{
							MarkdownDescription: "Set the intra-transaction timeout for Web Application Firewall to server connections. This value is the maximum idle time to allow in a transaction on the Web Application Firewall to server connection. This timer monitors the idle time in the data transfer process. If the specified idle time is exceeded, the connection is torn down. Enter a value in the range 1 - 86400. The default value is 120.",
							Computed:            true,
						},
						"front_persistent_timeout": schema.Int64Attribute{
							MarkdownDescription: "Set the inter-transaction timeout for Web Application Firewall to client connections. This value is the maximum idle time to allow between the completion of a TCP transaction and the initiation of a new TCP transaction on the Web Application Firewall to client connection. If the specified idle time is exceeded, the connection is torn down. Enter a value in the range 0 - 86400. The default value is 180. A value of 0 disables persistent connections.",
							Computed:            true,
						},
						"allow_cache_control_header": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to allow the HTTP GET method to pass the Cache-Control header through to the back end. If disabled, a \"Cache-Control:no-transform\" header is passed. If enabled, the client request can specify the cache control behavior, or if the client request does not specify the Cache-Control header, a \"Cache-Control:no-transform\" header is passed.",
							Computed:            true,
						},
						"back_persistent_timeout": schema.Int64Attribute{
							MarkdownDescription: "<p>Specify the maximum inter-transaction idle time in seconds for Web Application Firewall to server connections. This value is the maximum idle time between the completion of a TCP transaction and the initiation of a new TCP transaction on this connection. If the specified idle time is exceeded, the connection is torn down. Enter a value in the range 0 - 86400. The default value is 180. A value of 0 disables persistent connections.</p><p><b>Note:</b> For HTTP GET and HEAD requests, the service attempts the connection again after the specified value, Therefore, the actual timeout is twice the specified value.</p><p>An idle TCP connection can remain in the idle state for 20 seconds after the expiration of the persistence timer.</p>",
							Computed:            true,
						},
						"front_http_version": schema.StringAttribute{
							MarkdownDescription: "Select the HTTP version to be use on client responses. Incoming version 1.0 requests will always be replied to with 1.0 compatible responses regardless of this setting. The default is HTTP 1.1.",
							Computed:            true,
						},
						"back_http_version": schema.StringAttribute{
							MarkdownDescription: "Select the HTTP version to use on the server-side connection. The default is HTTP 1.1.",
							Computed:            true,
						},
						"request_side_security": schema.BoolAttribute{
							MarkdownDescription: "If this property is disabled, no request side security policies will be enforced and all requests will be allowed through.",
							Computed:            true,
						},
						"response_side_security": schema.BoolAttribute{
							MarkdownDescription: "If this property is disabled, no response side security policies will be enforced and all responses will be allowed through.",
							Computed:            true,
						},
						"do_chunked_upload": schema.BoolAttribute{
							MarkdownDescription: "Use the radio buttons to enable (on) or disable (off) the ability to send Content-Type Chunked Encoded documents to the back end server. When the device employs the HTTP/1.1 protocol, the body of the document can be delimited by either Content-Length or chunked encodings. While all servers will understand how to interpret Content-Length, many applications will fail to understand Chunked encoding. For this reason, Content-Length is the standard method used. However doing so interferes with the ability of the device to fully stream. To stream full documents towards the back end server, this property should be turned on. However, the back end server must be RFC 2616 compatible, because this feature cannot be renegotiated at run time, unlike all other HTTP/1.1 features which can be negotiated down at runtime if necessary. This property can also be enabled by configuring a User Agent to enable it on a per-URL basis.",
							Computed:            true,
						},
						"follow_redirects": schema.BoolAttribute{
							MarkdownDescription: "Some protocols generate redirects as part of the protocol - for example HTTP response code 302. If this property is enabled the firewall will try and transparently resolve those redirects.",
							Computed:            true,
						},
						"http_client_ip_label": schema.StringAttribute{
							MarkdownDescription: "Retain the default value (X-Client-IP) or provide an other value (for example, X-Forwarded-For).",
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
							MarkdownDescription: "Set the number of records for transaction diagnostics in the probe. Enter a value in the range 10 - 250. The default value is 25.",
							Computed:            true,
						},
						"debug_trigger": schema.ListNestedAttribute{
							MarkdownDescription: "The probe captures transactions that meet one or more of the conditions defined by the triggers. These triggers examine the direction or type of the message flow and examine the message for an XPath expression match. When a message meets one of these conditions, the transaction is captured in diagnostics mode and becomes part of the list of transactions that can be viewed.",
							NestedObject:        models.DmMSDebugTriggerTypeDataSourceSchema,
							Computed:            true,
						},
						"url_rewrite_policy": schema.StringAttribute{
							MarkdownDescription: "A Rewrite Policy can be used to modify the values of certain headers before the security policies are executed. One common use for this is to rewrite the contents of an HTTP Location header to reflect external DNS names. By using a Rewrite Policy, it is possible to intercept and rewrite URLs that might send a browser directly to the back end web site, rather than through the Web Application Firewall.",
							Computed:            true,
						},
						"do_host_rewriting": schema.BoolAttribute{
							MarkdownDescription: "Some protocols have distinct name based elements, separate from the URL, to de multiplex. HTTP uses the Host header for this purposes. If this feature is enabled the backside server will receive a request reflecting the final route, otherwise it will receive a request reflecting the information as it arrived at the DataPower device. Web servers issuing redirects may want to disable this feature, as they often depend on the host header for the value of their redirect.",
							Computed:            true,
						},
						"ssl_config_type": schema.StringAttribute{
							MarkdownDescription: "TLS type",
							Computed:            true,
						},
						"ssl_server": schema.StringAttribute{
							MarkdownDescription: "TLS server profile",
							Computed:            true,
						},
						"sslsni_server": schema.StringAttribute{
							MarkdownDescription: "TLS SNI server profile",
							Computed:            true,
						},
						"ssl_client": schema.StringAttribute{
							MarkdownDescription: "TLS client profile",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *WebAppFWDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *WebAppFWDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data WebAppFWList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.WebAppFW{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.WebAppFW{}
	if value := res.Get(`WebAppFW`); value.Exists() {
		for _, v := range value.Array() {
			item := models.WebAppFW{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.WebAppFWObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.WebAppFWObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
