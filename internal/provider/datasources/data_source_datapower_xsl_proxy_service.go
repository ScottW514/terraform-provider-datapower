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

type XSLProxyServiceList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &XSLProxyServiceDataSource{}
	_ datasource.DataSourceWithConfigure = &XSLProxyServiceDataSource{}
)

func NewXSLProxyServiceDataSource() datasource.DataSource {
	return &XSLProxyServiceDataSource{}
}

type XSLProxyServiceDataSource struct {
	pData *tfutils.ProviderData
}

func (d *XSLProxyServiceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_xsl_proxy_service"
}

func (d *XSLProxyServiceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "<p>The XSL Proxy is obsolete.</p>",
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
						"type": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"xml_manager": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"url_rewrite_policy": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"style_policy": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"credential_charset": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"ssl_config_type": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"ssl_server": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"sslsni_server": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"ssl_client": schema.StringAttribute{
							MarkdownDescription: "",
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
						"local_port": schema.Int64Attribute{
							MarkdownDescription: "Specify the local port to monitor for incoming client requests.",
							Computed:            true,
						},
						"remote_address": schema.StringAttribute{
							MarkdownDescription: "Specify the host name or IP address of the specific server supported by this DataPower service. If using load balancers, specify the name of the Load Balancer Group. If using the On Demand Router, specify the keyword ODR-LBG. Load balancer groups and the On Demand Router can be used only when Type is static-backend.",
							Computed:            true,
						},
						"remote_port": schema.Int64Attribute{
							MarkdownDescription: "Specify the port number to monitor. Used only when Type is static-backend.",
							Computed:            true,
						},
						"acl": schema.StringAttribute{
							MarkdownDescription: "This Access Control List will be used to allow or deny access to this service based on the IP address of the client. When attached to a service, an Access Control List (ACL) denies all access by default. To deny access to only selected addresses, first grant access to all addresses (allow 0.0.0.0) and then create deny entries for the desired hosts.",
							Computed:            true,
						},
						"http_timeout": schema.Int64Attribute{
							MarkdownDescription: "Specifies the maximum number of seconds (within the range 1 through 86400, with a default of 120) that a firewall or proxy maintains an idle HTTP connection.",
							Computed:            true,
						},
						"http_persist_timeout": schema.Int64Attribute{
							MarkdownDescription: "Specifies the maximum number of seconds (within the range 0 through 7200, with a default of 180) that a firewall or proxy maintains an idle TCP connection.",
							Computed:            true,
						},
						"do_host_rewrite": schema.BoolAttribute{
							MarkdownDescription: "When enabled, the device will rewrite the Host: header to be the address of the back-end server. This is not what a strict proxy would do, and may not be appropriate for all topologies.",
							Computed:            true,
						},
						"suppress_http_warnings": schema.BoolAttribute{
							MarkdownDescription: "Toggle to enable or disable the generation of Transformation Applied (Warning Code: 214) messages by the HTTP service.",
							Computed:            true,
						},
						"http_compression": schema.BoolAttribute{
							MarkdownDescription: "Toggle to enable or disable the GZIP compression function.",
							Computed:            true,
						},
						"http_include_response_type_encoding": schema.BoolAttribute{
							MarkdownDescription: "Toggle to enable or disable including the character set encoding in the HTTP content-type header produced. For example, when sending a UTF-8 encoded XML document: If this property is disabled, 'content-type=text/xml' will be sent. If this property is enabled, 'content-type=text/xml; charset=UTF-8' will be sent.",
							Computed:            true,
						},
						"always_show_errors": schema.BoolAttribute{
							MarkdownDescription: "If set, HTTP responses may be generated with errors appended to partially generated documents. If not set error responses will only be sent to the browser if no other output has been created.",
							Computed:            true,
						},
						"disallow_get": schema.BoolAttribute{
							MarkdownDescription: "If set, only methods with incoming data (such as POST) are allowed.",
							Computed:            true,
						},
						"disallow_empty_response": schema.BoolAttribute{
							MarkdownDescription: "If set, only responses with message bodies are allowed (that is, not 304 and so forth).",
							Computed:            true,
						},
						"http_persistent_connections": schema.BoolAttribute{
							MarkdownDescription: "Toggle to enable or disable HTTP persistent connections.",
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
						"http_proxy_host": schema.StringAttribute{
							MarkdownDescription: "Specify the host name or IP address of the HTTP proxy to use when the remote server can be accessed only through another HTTP proxy.",
							Computed:            true,
						},
						"http_proxy_port": schema.Int64Attribute{
							MarkdownDescription: "Specify the port number on the HTTP proxy server.",
							Computed:            true,
						},
						"http_version": models.GetDmHTTPClientServerVersionDataSourceSchema("HTTP Version", "version", ""),
						"do_chunked_upload": schema.BoolAttribute{
							MarkdownDescription: "Use the radio buttons to enable (on) or disable (off) the ability to send Content-Type Chunked Encoded documents to the back end server. When the device employs the HTTP/1.1 protocol, the body of the document can be delimited by either Content-Length or chunked encodings. While all servers will understand how to interpret Content-Length, many applications will fail to understand Chunked encoding. For this reason, Content-Length is the standard method used. However doing so interferes with the ability of the device to fully stream. To stream full documents towards the back end server, this property should be turned on. However, the back end server must be RFC 2616 compatible, because this feature cannot be renegotiated at run time, unlike all other HTTP/1.1 features which can be negotiated down at runtime if necessary. This property can also be enabled by configuring a User Agent to enable it on a per-URL basis.",
							Computed:            true,
						},
						"header_injection": schema.ListNestedAttribute{
							MarkdownDescription: "HTTP Header Injection",
							NestedObject:        models.DmHeaderInjectionDataSourceSchema,
							Computed:            true,
						},
						"header_suppression": schema.ListNestedAttribute{
							MarkdownDescription: "HTTP Headers can be suppressed (removed) from the message flow using this property. For example, the Via: header, which contains the name of the DataPower service handling the message, may be suppressed from messages sent by the DataPower device back to the client.",
							NestedObject:        models.DmHeaderSuppressionDataSourceSchema,
							Computed:            true,
						},
						"stylesheet_parameters": schema.ListNestedAttribute{
							MarkdownDescription: "Stylesheets used in Processing Policies can take stylesheet parameters. These parameters can be passed in by this object. More than one parameter can be defined.",
							NestedObject:        models.DmStylesheetParameterDataSourceSchema,
							Computed:            true,
						},
						"default_param_namespace": schema.StringAttribute{
							MarkdownDescription: "If a stylesheet parameter is defined without a namespace (or without explicitly specifying the null namespace), then this is the namespace into which the parameter will be assigned.",
							Computed:            true,
						},
						"query_param_namespace": schema.StringAttribute{
							MarkdownDescription: "The namespace in which to put all parameters that are specified in the URL query string.",
							Computed:            true,
						},
						"force_policy_exec": schema.BoolAttribute{
							MarkdownDescription: "<p>Some message patterns may include bodyless request and response messages. This is common with RESTful web services where messages may or may not include a body but still requires the processing policy to run. To enable this capability for services whose request and response type is XML (or marked as non-XML i.e. JSON), set this option to 'on'. By doing so, the processing policy rules will always be executed.</p>",
							Computed:            true,
						},
						"count_monitors": schema.ListAttribute{
							MarkdownDescription: "Count Monitors watch for defined messaging events and increment counters each time the event occurs. When a certain threshold is reached, the monitor can either post a notification to a log or block service for a configured amount of time. Select a Count Monitor from the list to activate that monitor for this firewall. Click the + button to create a new Count Monitor.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"duration_monitors": schema.ListAttribute{
							MarkdownDescription: "Duration Monitors watch for events meeting or exceeding a configured duration. When a duration is met or exceeded, the monitor can either post a notification to a log or block service for a configured amount of time. Select a Duration Monitor from the list to activate that monitor for this firewall. Click the + button to create a new Duration Monitor.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"monitor_processing_policy": schema.StringAttribute{
							MarkdownDescription: "Select the way that the system behaves when more than one monitor is attached to a service.",
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
						"debug_trigger": schema.ListNestedAttribute{
							MarkdownDescription: "The probe captures transactions that meet one or more of the conditions defined by the triggers. These triggers examine the direction or type of the message flow and examine the message for an XPath expression match. When a message meets one of these conditions, the transaction is captured in diagnostics mode and becomes part of the list of transactions that can be viewed.",
							NestedObject:        models.DmMSDebugTriggerTypeDataSourceSchema,
							Computed:            true,
						},
						"local_address": schema.StringAttribute{
							MarkdownDescription: "<p>Enter a host alias or the IP address that the service listens on. Host aliases can ease migration tasks among appliances.</p><ul><li>0 or 0.0.0.0 indicates all configured IPv4 addresses.</li><li>:: indicates all configured IPv4 and IPv6 addresses.</li></ul><p><b>Attention:</b> For management services, the value of 0.0.0.0 or :: is a security risk. Use an explicit IP address to isolate management traffic from application data traffic.</p>",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *XSLProxyServiceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *XSLProxyServiceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data XSLProxyServiceList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.XSLProxyService{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.XSLProxyService{}
	if value := res.Get(`XSLProxyService`); value.Exists() {
		for _, v := range value.Array() {
			item := models.XSLProxyService{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.XSLProxyServiceObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.XSLProxyServiceObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
