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

type XMLFirewallServiceList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &XMLFirewallServiceDataSource{}
	_ datasource.DataSourceWithConfigure = &XMLFirewallServiceDataSource{}
)

func NewXMLFirewallServiceDataSource() datasource.DataSource {
	return &XMLFirewallServiceDataSource{}
}

type XMLFirewallServiceDataSource struct {
	pData *tfutils.ProviderData
}

func (d *XMLFirewallServiceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_xml_firewall_service"
}

func (d *XMLFirewallServiceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Create or edit an XML Firewall on local IP/port. This XML Firewall can communicate with a dynamically identified servers, a static back end server or as a loopback. The XML Firewall applies the selected processing policy to messages. The XML Firewall can rewrite client request URLs byusing a URL rewrite policy. The XML Firewall can use TLS communications to client, server or both directions if applicable.",
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
						"type": schema.StringAttribute{
							MarkdownDescription: "Specify the mode of the XML Firewall. The default is Dynamic Backend.",
							Computed:            true,
						},
						"xml_manager": schema.StringAttribute{
							MarkdownDescription: "An XML Manager manages the compilation and caching of stylesheets, the caching of documents, and provides configuration constraints on the size and parsing depth of documents. You can enable streaming operation by configuring an XML Manager to use a Streaming Compile Option Policy. More than one firewall may use the same XML Manager. Select an existing XML Manager from the list to assign the Manager to this firewall. Click the + button to create a new XML Manager that is assigned to this firewall. A default Manager is used if you do not create one.",
							Computed:            true,
						},
						"url_rewrite_policy": schema.StringAttribute{
							MarkdownDescription: "A URL Rewrite Policy applies the rules established in the named policy to rewrite the URL used by the client to connect to the service this firewall represents. This helps to mask the URL and location of the actual service. Select an existing URL rewrite policy from the list of available policies. Click the + button to create a new URL rewrite policy that is then assigned to this firewall.",
							Computed:            true,
						},
						"style_policy": schema.StringAttribute{
							MarkdownDescription: "Assign the processing policy to the service. The processing policy defines the actions to take against messages that flow through the service.",
							Computed:            true,
						},
						"max_message_size": schema.Int64Attribute{
							MarkdownDescription: "<p>Specifies the maximum allowable size in kilobytes of a message. Enter a value in the range 0 - 2097151. The default value is 0. A value of 0 specifies that no limit is enforced, and the message can be of unlimited size.</p><p>The maximum message size limit applies to JSON, SOAP, XML, and non-XML messages. If the message type is pass-through, no limit is enforced.</p>",
							Computed:            true,
						},
						"request_type": schema.StringAttribute{
							MarkdownDescription: "Characterizes the type of traffic that originates from the client. The default is SOAP.",
							Computed:            true,
						},
						"response_type": schema.StringAttribute{
							MarkdownDescription: "Characterizes the type of traffic that originates from the server. The default is SOAP.",
							Computed:            true,
						},
						"fw_cred": schema.StringAttribute{
							MarkdownDescription: "Specifies an optional Firewall Credentials list. A Firewall Credentials list specifies which keys and certificates are available to support firewall processing. In the absence of an identified Firewall Credentials list, all locally-stored key and certificates are available to the firewall.",
							Computed:            true,
						},
						"service_monitors": schema.ListAttribute{
							MarkdownDescription: "Specifies a list of Web Service Monitors. Web Service Monitors target Web Service endpoints. Use the Web Services wizard on the Control Panel to supply a WSDL and configure a Monitor for an endpoint. The Monitor collects statistics, establishes count and duration monitors and can take action when thresholds are met or exceeded. A Monitor must be selected here to be activated. Click + to create a new Web Services Monitor. Note that this Monitor is not the Service Level Monitor (SLM) used by a Multi-Protocol Gateway or Web Service Proxy.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"request_attachments": schema.StringAttribute{
							MarkdownDescription: "Select how to treat client requests with attachments. The default is Strip.",
							Computed:            true,
						},
						"response_attachments": schema.StringAttribute{
							MarkdownDescription: "Select how to treat server responses with attachments. The default is Strip.",
							Computed:            true,
						},
						"root_part_not_first_action": schema.StringAttribute{
							MarkdownDescription: "When streaming MIME messages, the action to take when the root part is not the first part of the message. If the root part must be first (for example to do BSP conformance checking) and the action is set to \"process in order\" the DP device will buffer attachments up to the root.",
							Computed:            true,
						},
						"front_attachment_format": schema.StringAttribute{
							MarkdownDescription: "Select how to interpret client requests with attachments. If an attachment arrives that is not of the type chosen, the attachment will be rejected.",
							Computed:            true,
						},
						"back_attachment_format": schema.StringAttribute{
							MarkdownDescription: "Select the attachment output format to backend servers. If an attachment arrives that is not of the type chosen, the attachment will be rejected.",
							Computed:            true,
						},
						"mime_headers": schema.BoolAttribute{
							MarkdownDescription: "<p>The body of a message (that is, the payload, independent of any protocol headers) can sometimes contain MIME headers before any preamble and before the first MIME boundary contained in the body of the message. These MIME headers can contain important information not available in the protocol headers, such as the string identifying the MIME boundary. If this property is set to on, then these MIME headers will be processed by the Firewall.</p><p>Note that if this is on and there are no MIME headers contained in the message, the device will continue to try and parse the message, using the protocol header information, if available.</p><p>When this is off and MIME headers are present in the body of the message, these MIME headers will be considered part of the preamble and not used to parse out the message. If the protocol headers (such as HTTP) indicate the MIME boundaries, the device can parse and process individual attachments. If such information is not available, no attachments can be parsed from the body of the message.</p>",
							Computed:            true,
						},
						"rewrite_errors": schema.BoolAttribute{
							MarkdownDescription: "Error messages after a decryption action can provide an attacker who is using the padding oracle attack method with enough information to determine the contents of the plaintext data. When enabled, the default, the client receives error messages without the internal information that could lead to a discovery. When disabled, the client receives the original message with this information.",
							Computed:            true,
						},
						"delay_errors": schema.BoolAttribute{
							MarkdownDescription: "The timing difference of the error messages returned after a decryption action can provide an attacker with enough information to determine the contents of the plain-text data. When enabled, the default, the DataPower Gateway delays error messages for the defined duration. When disabled, the DataPower Gateway does not delay error messages.",
							Computed:            true,
						},
						"delay_errors_duration": schema.Int64Attribute{
							MarkdownDescription: "When enabling the delay of error messages, specify the delay duration in milliseconds. If delaying messages for 3000 ms, the DataPower Gateway will not send error messages to the client until 3 seconds have elapsed since the DataPower Gateway performed decryption on the requests. Use any value of 250 - 300000. The default value is 1000.",
							Computed:            true,
						},
						"soap_schema_url": schema.StringAttribute{
							MarkdownDescription: "When a firewall is in SOAP mode, either on the request or response side, it will validate the incoming messages against a W3C Schema that defines what a SOAP message is supposed to look like. It is possible to customize which schema is used on a per-firewall basis by changing this property; this can be used to accommodate non-standard configurations or other special cases.",
							Computed:            true,
						},
						"wsdl_response_policy": schema.StringAttribute{
							MarkdownDescription: "Select how the firewall handles requests for .NET WSDL requests via the http://domain.com/service?wsdl convention. The default is Off.",
							Computed:            true,
						},
						"wsdl_file_location": schema.StringAttribute{
							MarkdownDescription: "URL of the file to return when WSDL response policy is set to serve local.",
							Computed:            true,
						},
						"firewall_parser_limits": schema.BoolAttribute{
							MarkdownDescription: "Use the firewall parser limits instead of the parser limits in the XML Manager for this firewall. Firewall limits override XML Manager limits.",
							Computed:            true,
						},
						"parser_limits_bytes_scanned": schema.Int64Attribute{
							MarkdownDescription: "Defines the maximum number of bytes scanned by the XML parser. This applies to any XML document that is parsed. If any of the parser limits are set in the XML Firewall, they will override those on the XML Manager. If this value is 0, no limit is enforced.",
							Computed:            true,
						},
						"parser_limits_element_depth": schema.Int64Attribute{
							MarkdownDescription: "Defines the maximum depth of element nesting in XML parser. If any of the parser limits are set in the XML Firewall, they will override those on the XML Manager.",
							Computed:            true,
						},
						"parser_limits_attribute_count": schema.Int64Attribute{
							MarkdownDescription: "Defines the maximum number of attributes of a given element. If any of the parser limits are set in the XML Firewall, they will override those on the XML Manager.",
							Computed:            true,
						},
						"parser_limits_max_node_size": schema.Int64Attribute{
							MarkdownDescription: "Defines the maximum size any one node may consume. The default is 32 MB. Sizes which are powers of two result in the best performance. If any of the parser limits are set in the XML Firewall, they will override those on the XML Manager. Although you set an explicit value, the DataPower Gateway uses a value that is the rounded-down, largest power of two that is smaller than the defined value.",
							Computed:            true,
						},
						"parser_limits_max_prefixes": schema.Int64Attribute{
							MarkdownDescription: "Enter an integer that defines the maximum number of distinct XML namespace prefixes in a document. This limit counts multiple prefixes defined for the same namespace, but does not count multiple namespaces defined in different parts of the input document under a single prefix. Enter a value in the rage 0 - 262143. The default value is 1024. A value of 0 indicates that the limit is 1024.",
							Computed:            true,
						},
						"parser_limits_max_namespaces": schema.Int64Attribute{
							MarkdownDescription: "Enter an integer that defines the maximum number of distinct XML namespace URIs in a document. This limit counts all XML namespaces, regardless of how many prefixes are used to declare them. Enter a value in the range 0 - 65535. The default value is 1024. A value of 0 indicates that the limit is 1024.",
							Computed:            true,
						},
						"parser_limits_max_local_names": schema.Int64Attribute{
							MarkdownDescription: "Enter an integer that defines the maximum number of distinct XML local names in a document. This limit counts all local names, independent of the namespaces they are associated with. Enter a value in the range 0 - 1048575. The default value is 60000. A value of 0 indicates that the limit is 60000.",
							Computed:            true,
						},
						"parser_limits_attachment_byte_count": schema.Int64Attribute{
							MarkdownDescription: "Defines the maximum number of bytes allowed in any single attachment. Attachments that exceed this size will result in a failure of the whole transaction. If this value is 0, no limit is enforced.",
							Computed:            true,
						},
						"parser_limits_attachment_package_byte_count": schema.Int64Attribute{
							MarkdownDescription: "Defines the maximum number of bytes allowed for all parts of an attachment package, including the root part. Attachment packages that exceed this size will result in a failure of the transaction. If this value is 0, no limit is enforced.",
							Computed:            true,
						},
						"parser_limits_external_references": schema.StringAttribute{
							MarkdownDescription: "Select the special handling for input documents that contain external references, such as an external entity or external DTD definition.",
							Computed:            true,
						},
						"credential_charset": schema.StringAttribute{
							MarkdownDescription: "Specifies the character encoding of the original basic authentication values. Basic authentication credentials are combined and base64 encoded in the authorization header of the request. The DataPower Gateway transcodes the contents of the authorization header to UTF-8. Defaults to Protocol which is ISO-8859-1, Latin 1.",
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
						"ssl_sni_server": schema.StringAttribute{
							MarkdownDescription: "TLS SNI server profile",
							Computed:            true,
						},
						"ssl_client": schema.StringAttribute{
							MarkdownDescription: "TLS client profile",
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
							NestedObject:        models.GetDmHeaderInjectionDataSourceSchema(),
							Computed:            true,
						},
						"header_suppression": schema.ListNestedAttribute{
							MarkdownDescription: "HTTP Headers can be suppressed (removed) from the message flow using this property. For example, the Via: header, which contains the name of the DataPower service handling the message, may be suppressed from messages sent by the DataPower device back to the client.",
							NestedObject:        models.GetDmHeaderSuppressionDataSourceSchema(),
							Computed:            true,
						},
						"stylesheet_parameters": schema.ListNestedAttribute{
							MarkdownDescription: "Stylesheets used in Processing Policies can take stylesheet parameters. These parameters can be passed in by this object. More than one parameter can be defined.",
							NestedObject:        models.GetDmStylesheetParameterDataSourceSchema(),
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
							NestedObject:        models.GetDmMSDebugTriggerTypeDataSourceSchema(),
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

func (d *XMLFirewallServiceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *XMLFirewallServiceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data XMLFirewallServiceList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.XMLFirewallService{
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
	l := []models.XMLFirewallService{}
	if resFound {
		if value := res.Get(`XMLFirewallService`); value.Exists() {
			for _, v := range value.Array() {
				item := models.XMLFirewallService{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.XMLFirewallServiceObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.XMLFirewallServiceObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
