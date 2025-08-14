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

type MultiProtocolGatewayList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &MultiProtocolGatewayDataSource{}
	_ datasource.DataSourceWithConfigure = &MultiProtocolGatewayDataSource{}
)

func NewMultiProtocolGatewayDataSource() datasource.DataSource {
	return &MultiProtocolGatewayDataSource{}
}

type MultiProtocolGatewayDataSource struct {
	pData *tfutils.ProviderData
}

func (d *MultiProtocolGatewayDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_multiprotocolgateway"
}

func (d *MultiProtocolGatewayDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "<p>A Multi-Protocol Gateway service can accept client-originated messages in a variety of protocols. The Gateway can subsequently pass messages to a back end server using a variety of protocols. The protocol used on the client side need not be the same as that on the back side.</p><p>A Gateway can support more than one client, or Front Side, protocol. Similarly, the Gateway can support more than one back side, or server protocol. The messages may be processed and routed using any of the standard Document Processing Actions available to a firewall or a proxy.</p>",
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
						"back_http_version": schema.StringAttribute{
							MarkdownDescription: "Select the HTTP version to use on the server-side connection. The default value is HTTP/1.1.",
							Computed:            true,
						},
						"http2_required": schema.BoolAttribute{
							MarkdownDescription: "Select whether an HTTP/2 connection is required when connecting to the server. Only applicable when the HTTP version to the server is set to HTTP/2 and the connection uses TLS. The default value is off.",
							Computed:            true,
						},
						"request_type": schema.StringAttribute{
							MarkdownDescription: "Characterizes the traffic that originates from the client. The default value is SOAP.",
							Computed:            true,
						},
						"response_type": schema.StringAttribute{
							MarkdownDescription: "Characterizes the traffic that originates from the server. The default value is SOAP.",
							Computed:            true,
						},
						"follow_redirects": schema.BoolAttribute{
							MarkdownDescription: "Some protocols generate redirects as part of the protocol - for example HTTP response code 302. If this property is enabled the gateway tries to transparently resolve those redirects.",
							Computed:            true,
						},
						"rewrite_location_header": schema.BoolAttribute{
							MarkdownDescription: "<p>When you enable this property and the following conditions are true, the gateway rewrites the host name and port in the <tt>Location</tt> response header with the gateway's host name and port.</p><ul><li>The host name and port in the <tt>Location</tt> header matches that of the target server.</li><li>The request contains a <tt>Host</tt> header.</li></ul><p>After the rewrite of the <tt>Location</tt> response header, subsequent requests are always sent to the gateway.</p>",
							Computed:            true,
						},
						"style_policy": schema.StringAttribute{
							MarkdownDescription: "Assign the processing policy to the service. The processing policy defines the actions to take against messages that flow through the service.",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Specify the mode of the gateway. The default value is static backend.",
							Computed:            true,
						},
						"allow_compression": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enable or disable the GZIP compression functionality on backside.",
							Computed:            true,
						},
						"enable_compressed_payload_passthrough": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enable or disable the decompression of the response payload when the backside response is compressed. When enabled, the payload is not decompressed but passed through compressed.",
							Computed:            true,
						},
						"allow_cache_control_header": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to allow the HTTP GET method to pass the <tt>Cache-Control</tt> header through to the back end. If disabled, a <tt>Cache-Control: no-transform</tt> header is passed. If enabled, the client request can specify the cache control behavior, or if the client request does not specify the <tt>Cache-Control</tt> header, a <tt>Cache-Control: no-transform</tt> header is passed.",
							Computed:            true,
						},
						"wsrr_saved_search_subscriptions": schema.ListNestedAttribute{
							MarkdownDescription: "The Gateway services endpoints based on the metadata returned by the saved search or named query.",
							NestedObject:        models.DmWSRRSavedSearchWSDLSourceDataSourceSchema,
							Computed:            true,
						},
						"wsrr_subscriptions": schema.ListNestedAttribute{
							MarkdownDescription: "The Gateway services endpoints based on the metadata obtained using these subscriptions.",
							NestedObject:        models.DmWSRRWSDLSourceDataSourceSchema,
							Computed:            true,
						},
						"policy_attachments": schema.StringAttribute{
							MarkdownDescription: "Select an existing policy attachment object to configure the handling of XML element attached policies and to create external policy attachments to this gateway.",
							Computed:            true,
						},
						"policy_parameter": schema.ListNestedAttribute{
							MarkdownDescription: "Policy parameters influence the enforcement of policy in the gateway",
							NestedObject:        models.DmWSPolicyParametersDataSourceSchema,
							Computed:            true,
						},
						"wsm_agent_monitor": schema.BoolAttribute{
							MarkdownDescription: "Control the collection of monitoring records for this service. The automatic collection of Web Services Management records applies to only transactions over HTTP protocols.",
							Computed:            true,
						},
						"wsm_agent_monitor_pcm": schema.StringAttribute{
							MarkdownDescription: "<p>Select messages to capture for analysis.</p><p>Because not all Web Services Management protocols can accommodate full message-capture, configure this property only if the spooler can forward full messages.</p><p>Full message-capture incurs a performance penalty.</p>",
							Computed:            true,
						},
						"proxy_http_response": schema.BoolAttribute{
							MarkdownDescription: "When you enable this setting, the service passes the response code from the server to the client without modification. If an error occurs when the service runs the response rule, the service does not pass the response code but returns the default error code, which is an HTTP 500. If you require the service to pass the response code in this error condition, set the <tt>var://service/mpgw/proxy-http-response</tt> variable to 1 in the request or response rule. However, when you set a custom response code in the error rule with the <tt>var://service/error-protocol-response</tt> variable, the service passes this custom response code instead.",
							Computed:            true,
						},
						"error_policy": schema.StringAttribute{
							MarkdownDescription: "Assign an error policy to the service. The error policy defines the actions to take against errors in an HTTP or HTTPS flow that no precedent error handler handles.",
							Computed:            true,
						},
						"transaction_timeout": schema.Int64Attribute{
							MarkdownDescription: "Sets the duration for a transaction in milliseconds before it is canceled. The transaction starts when the handler receives the request. You can override this timeout dynamically during processing with the <tt>var://service/mpgw/transaction-timeout-ms-threshold</tt> variable or adjust the current timer value with the <tt>var://service/mpgw/transaction-timeout-ms-count variable</tt> .",
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
						"front_protocol": schema.ListAttribute{
							MarkdownDescription: "Specify a protocol handler to determine the network communication protocol, address, port, and other protocol-specific settings.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"xml_manager": schema.StringAttribute{
							MarkdownDescription: "<p>An XML Manager manages the compilation and caching of style sheets, the caching of documents, and provides configuration constraints on the size and parsing depth of documents. You can enable streaming operation by configuring an XML Manager to use a Streaming Compile Option Policy.</p><p>Optionally, an XML Manager can employ a User Agent. The User Agent settings, in turn, can affect the behavior of the gateway when communicating with remote servers or with clients when sending back responses.</p>",
							Computed:            true,
						},
						"url_rewrite_policy": schema.StringAttribute{
							MarkdownDescription: "<p>Define a URL Rewrite Policy that defines rules for the following rewrite and replacement operations:</p><ul><li>Rewrite the entire URL or a portion of the URL</li><li>Replace a header value in a message</li><li>Rewrite the HTTP POST body in a message</li></ul><p>The rewrite rules are applied before document processing. Therefore, the evaluation criteria in the Matching Rule is against the rewritten value.</p>",
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
						"fw_cred": schema.StringAttribute{
							MarkdownDescription: "Specify a credentials list for the service. The credentials specify which keys and certificates are available to support processing. In the absence of credentials, all locally-stored key and certificates are available to the service.",
							Computed:            true,
						},
						"header_injection": schema.ListNestedAttribute{
							MarkdownDescription: "HTTP Header Injection",
							NestedObject:        models.DmHeaderInjectionDataSourceSchema,
							Computed:            true,
						},
						"header_suppression": schema.ListNestedAttribute{
							MarkdownDescription: "HTTP Header Suppression",
							NestedObject:        models.DmHeaderSuppressionDataSourceSchema,
							Computed:            true,
						},
						"stylesheet_parameters": schema.ListNestedAttribute{
							MarkdownDescription: "Style sheets used in processing policies can take stylesheet parameters. These parameters can be passed in by this object. You can define more than one parameter.",
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
						"backend_url": schema.StringAttribute{
							MarkdownDescription: "<p>Specify the URL of the static backend server. The URL is parsed to determine the protocol to use.</p><p>To use a load balancer, specify the name of an existing Load Balancer Group instead of the address-port pair in the URL.</p><p>To use the on demand router, specify http://ODR-LBG instead of the address-port pair in the URL.</p><ul><li>When the URL starts with http:// or https://, the DataPower service uses the HTTP protocol.</li><li>When the URL starts with amqp-broker://, the DataPower service uses the AMQP protocol.</li><li>When the URL starts with ftp://, the DataPower service uses the FTP protocol. <ul><li>For a relative path to the login directory of the user, specify the URL in the ftp://user:password@host:port/path format.</li><li>For an absolute path to the root directory, specify the URL in the ftp://user:password@host/%2Fpath format. %2F is the URL encoding of a forward slash.</li></ul><p>Include a password in the URL with caution. The use of user:/password@host results in a server connection. However, with this configuration, the connection could be unable to send multiple commands to the FTP server. For a stable connection, define a basic authentication policy in the user agent. The user agent is in the XML manager associated with the DataPower service.</p></li><li>When the URL starts with dpmq://, dpmqfte, idgmq://, or idgmqmft://, the DataPower service uses the IBM MQ protocol. The referenced queue manager in the URL must exist in the current domain.</li><li>When the URL starts with dpims:// or dpimsssl://, the DataPower service uses the IMS protocol. These protocols are available with the IMS feature.</li><li>When the URL starts with dpkafka://, the DataPower service uses the Kafka protocol.</li><li>When the URL starts with dpnfs://, the DataPower service uses the NFS protocol. <ul><li>For static mounts, specify the URL in the dpnfs://mount format, where mount is the name of an existing NFS mount.</li><li>For dynamic mounts, specify the URL in the dpnfs://host or dpnfs://host/path format.</li></ul></li><li>When the URL starts with sftp://, the DataPower service uses the SSH FTP protocol. <ul><li>For an absolute path to the root directory, specify the URL in the sftp://host:port/path format.</li><li>For a relative path to the login directory of the user, specify the URL in the sftp://host:port/~/path format.</li></ul><p>If password authentication is desired, specify the password in the SFTP Client Policy of the User Agent. If Public key authentication is desired, specify the user's private key in the SFTP Client Policy of the User Agent. When no SFTP Client Policy is specified, the Basic-Auth Policy and Pubkey-Auth Policy take effect.</p></li><li>When the URL starts with dptibems://, the DataPower service uses the TIBCO EMS protocol. The referenced server in the URL must exist. These protocols are available with the TIBCO EMS feature.</li><li>When the URL starts with dpwasjms://, the DataPower service uses the WebSphere JMS protocol. The referenced server in the URL must exist.</li></ul><p>For assistance in building a URL, click the appropriate Helper button.</p>",
							Computed:            true,
						},
						"propagate_uri": schema.BoolAttribute{
							MarkdownDescription: "<p>URI propagation is meaningful only when:</p><ul><li>The service uses a static backend.</li><li>The service uses a dynamic backend and routing is set with a route with stylesheet (route-action) action. In this case, use dp:set-target() to define the backend server. <p>For the other routing options that are available with route-action and route-set, the URI is absolute.</p></li></ul><p>When set to <b>on</b>, the service rewrites the URI of the backend URL to the URI in the client request. If the client submits http://host/service and the backend URL is http://server/listener, the URL is rewritten to http://server/service.</p><p>If the backend URL employs AMQP, IBM MQ, Kafka, TIBCO EMS, or WebSphere JMS, set to <b>off</b>.</p>",
							Computed:            true,
						},
						"service_monitors": schema.ListAttribute{
							MarkdownDescription: "Specifies a list of Web Services Monitors for the service. Web Services Monitors target Web Services endpoints. Use the Web Services wizard on the Control Panel to supply a WSDL and configure a Monitor for an endpoint. The Monitor collects statistics, establishes count and duration monitors and can take action when thresholds are met or exceeded. A Monitor must be included here to be activated. Note that this is not the same Monitor as a Service Level Monitor (SLM) that can be included in the Processing Policy of this service.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"count_monitors": schema.ListAttribute{
							MarkdownDescription: "Specify the count monitors to define messaging events and increment counters each time the event occurs. When a certain threshold is reached, the monitor can either post a notification to a log or block service for a configured amount of time.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"duration_monitors": schema.ListAttribute{
							MarkdownDescription: "Specify the duration monitors to watch for events that meet or exceed a configured duration. When a duration is met or exceeded, the monitor can either post a notification to a log or block service for a configured amount of time.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"monitor_processing_policy": schema.StringAttribute{
							MarkdownDescription: "Select the way that the system behaves when more than one monitor is attached to a service.",
							Computed:            true,
						},
						"request_attachments": schema.StringAttribute{
							MarkdownDescription: "Select how to treat client requests with attachments. The default is strip.",
							Computed:            true,
						},
						"response_attachments": schema.StringAttribute{
							MarkdownDescription: "Select how to treat server responses with attachments. The default is strip.",
							Computed:            true,
						},
						"request_attachments_flow_control": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enable or disable flow control for attachments in requests to servers. The default is off.",
							Computed:            true,
						},
						"response_attachments_flow_control": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enable or disable flow control for attachments in responses to clients. The default is off.",
							Computed:            true,
						},
						"root_part_not_first_action": schema.StringAttribute{
							MarkdownDescription: "When streaming MIME messages, the action to take when the root part is not the first part of the message. If the root part must be first (for example to do BSP conformance checking) and the action is set to \"process in order\" the DP device will buffer attachments up to the root.",
							Computed:            true,
						},
						"front_attachment_format": schema.StringAttribute{
							MarkdownDescription: "Select how to interpret client requests with attachments.",
							Computed:            true,
						},
						"back_attachment_format": schema.StringAttribute{
							MarkdownDescription: "Select the attachment output format to backend servers.",
							Computed:            true,
						},
						"mime_front_headers": schema.BoolAttribute{
							MarkdownDescription: "<p>The body of a message (that is, the payload, independent of any protocol headers) can sometimes contain MIME headers before any preamble and before the first MIME boundary contained in the body of the message. These MIME headers can contain important information not available in the protocol headers, such as the string identifying the MIME boundary. If this property is set to on, then these MIME headers will be processed by the service.</p><p>Note that if this is on and there are no MIME headers contained in the message, the device will continue to try and parse the message, using the protocol header information, if available.</p><p>When this is off and MIME headers are present in the body of the message, these MIME headers will be considered part of the preamble and not used to parse out the message. If the protocol headers (such as HTTP) indicate the MIME boundaries, the device can parse and process individual attachments. If such information is not available, no attachments can be parsed from the body of the message.</p>",
							Computed:            true,
						},
						"mime_back_headers": schema.BoolAttribute{
							MarkdownDescription: "<p>The body of a message (the payload independent of protocol headers) can sometimes contain MIME headers before any preamble and before the first MIME boundary in the body of the message. These MIME headers can contain important information that is not available in the protocol headers, such as a string that identifies the MIME boundary. If this property is set to on, these MIME headers are processed by the service.</p><p>When set to on and there are no MIME headers in the message, the device continues to try and parse the message, using the protocol header information, if available.</p><p>When set to off and MIME headers are present in the body of the message, these MIME headers are considered part of the preamble and not used to parse the message. If protocol headers, such as HTTP, indicate MIME boundaries, the device can parse and process individual attachments. If such information is not available, no attachments can be parsed from the body of the message.</p>",
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
						"max_message_size": schema.Int64Attribute{
							MarkdownDescription: "<p>Specifies the maximum allowable size in kilobytes of a message. Enter a value in the range 0 - 2097151. The default value is 0. A value of 0 specifies that no limit is enforced, and the message can be of unlimited size.</p><p>The maximum message size limit applies to JSON, SOAP, XML, and non-XML messages. If the message type is pass-through, no limit is enforced.</p>",
							Computed:            true,
						},
						"gateway_parser_limits": schema.BoolAttribute{
							MarkdownDescription: "Use the server-defined parser limits instead of the parser limits in the XML manager for this DataPower service. Service-defined limits override XML Manager limits.",
							Computed:            true,
						},
						"parser_limits_element_depth": schema.Int64Attribute{
							MarkdownDescription: "Defines the maximum depth of element nesting in XML parser. If any of the parser limits are set in the DataPower service, they will override those on the XML Manager.",
							Computed:            true,
						},
						"parser_limits_attribute_count": schema.Int64Attribute{
							MarkdownDescription: "Defines the maximum number of attributes of a given element. If any of the parser limits are set in the service, they will override those on the XML Manager.",
							Computed:            true,
						},
						"parser_limits_max_node_size": schema.Int64Attribute{
							MarkdownDescription: "Defines the maximum size any one node may consume. The default is 32 MB. Sizes which are powers of two result in the best performance. If any of the parser limits are set in the DataPower service, they will override those on the XML Manager. Although you set an explicit value, the DataPower Gateway uses a value that is the rounded-down, largest power of two that is smaller than the defined value.",
							Computed:            true,
						},
						"parser_limits_external_references": schema.StringAttribute{
							MarkdownDescription: "Select the special handling for input documents that contain external references, such as an external entity or external DTD definition.",
							Computed:            true,
						},
						"parser_limits_max_prefixes": schema.Int64Attribute{
							MarkdownDescription: "Enter an integer that defines the maximum number of distinct XML namespace prefixes in a document. This limit counts multiple prefixes defined for the same namespace, but does not count multiple namespaces defined in different parts of the input document under a single prefix. Enter a value in the range 0 - 262143. The default value is 1024. A value of 0 indicates that the limit is 1024.",
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
							MarkdownDescription: "Defines the maximum number of bytes allowed for all parts of an attachment package, including the root part. Attachment packages that exceed this size will result in a failure of the whole transaction. If this value is 0, no limit is enforced.",
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
						"flow_control": schema.BoolAttribute{
							MarkdownDescription: "<p>Use flow control to manage the transmission of large files when the front side and back end have different latencies and throughput while in streaming mode. This option is available when the request type and the response type are Pass-Thru or Non-XML. By default, flow control is disabled.</p><p>Note that the streaming behavior for output to back and output to front must be set to Stream Messages. Also, if either the front side or the back end will use the HTTP(S) protocol, allow chunked uploads must be set to on.</p>",
							Computed:            true,
						},
						"soap_schema_url": schema.StringAttribute{
							MarkdownDescription: "When a firewall is in SOAP mode, either on the request or response side, it will validate the incoming messages against a W3C Schema that defines what a SOAP message is supposed to look like. It is possible to customize which schema is used on a per-firewall basis by changing this property; this can be used to accommodate non-standard configurations or other special cases.",
							Computed:            true,
						},
						"front_timeout": schema.Int64Attribute{
							MarkdownDescription: "Set the intra-transaction timeout for service-to-client connections. This value is the maximum idle time to allow in a transaction on the service-to-client connection. This timer monitors idle time in the data transfer process. If the specified idle time is exceeded, the connection is torn down. Enter a value in the range 1 - 86400. The default value is 120.",
							Computed:            true,
						},
						"back_timeout": schema.Int64Attribute{
							MarkdownDescription: "Set the intra-transaction timeout for service-to-server connections. This value is the maximum idle time to allow in a transaction on the service-to-server connection. This timer monitors the idle time in the data transfer process. If the specified idle time is exceeded, the connection is torn down. Enter a value in the range 1 - 86400. The default value is 120.",
							Computed:            true,
						},
						"front_persistent_timeout": schema.Int64Attribute{
							MarkdownDescription: "Set the inter-transaction timeout for service-to-client connections. This value is the maximum idle time to allow between the completion of a TCP transaction and the initiation of a new TCP transaction on the service-to-client connection. If the specified idle time is exceeded, the connection is torn down. Enter a value in the range 0 - 86400. The default value is 180. A value of 0 disables persistent connections.",
							Computed:            true,
						},
						"back_persistent_timeout": schema.Int64Attribute{
							MarkdownDescription: "<p>Specify the maximum inter-transaction idle time in seconds for service-to-server connections. This value is the maximum idle time between the completion of a TCP transaction and the initiation of a new TCP transaction on this connection. If the specified idle time is exceeded, the connection is torn down. Enter a value in the range 0 - 86400. The default value is 180. A value of 0 disables persistent connections.</p><p><b>Note:</b> For HTTP GET and HEAD requests, the service attempts the connection again after the specified value, Therefore, the actual timeout is twice the specified value.</p><p>An idle TCP connection can remain in the idle state for 20 seconds after the expiration of the persistence timer.</p>",
							Computed:            true,
						},
						"include_response_type_encoding": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enable or disable including the character set encoding in any content-type header or description produced. For example, when sending a UTF-8 encoded XML document: If this property is disabled, 'text/xml' will be sent. If this property is enabled, 'text/xml; charset=UTF-8' will be sent.",
							Computed:            true,
						},
						"persistent_connections": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enable or disable persistent connections where appropriate on back side.",
							Computed:            true,
						},
						"loop_detection": schema.BoolAttribute{
							MarkdownDescription: "Some protocols provide a loop detection mechanism that can be used to detect gateway loops. This is a good policy, but it can cause each gateway to be publicly recorded in the transmitted message, which might be considered an information leak.",
							Computed:            true,
						},
						"do_host_rewriting": schema.BoolAttribute{
							MarkdownDescription: "Some protocols have distinct name-based elements, separate from the URL, to demultiplex. HTTP uses the <tt>Host</tt> header for this purposes. If this feature is enabled, the remote server will receive a request that reflects the final route; otherwise, it will receive a request that reflects the information as it arrived at the DataPower service. Web servers that issue redirects might want to disable this feature, as they often depend on the <tt>Host</tt> header for the value of their redirect.",
							Computed:            true,
						},
						"do_chunked_upload": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enable (on) or disable (off) the ability to send Content-Type Chunked Encoded documents to the back end server. When the device employs the HTTP/1.1 protocol, the body of the document can be delimited by either Content-Length or chunked encodings. While all servers will understand how to interpret Content-Length, many applications will fail to understand Chunked encoding. For this reason, Content-Length is the standard method used. However doing so interferes with the ability of the device to fully stream. To stream full documents towards the back end server, this property should be turned on. However, the back end server must be RFC 2616 compatible, because this feature cannot be renegotiated at run time, unlike all other HTTP/1.1 features which can be negotiated down at runtime if necessary. This property can also be enabled by configuring a User Agent to enable it on a per-URL basis.",
							Computed:            true,
						},
						"process_http_errors": schema.BoolAttribute{
							MarkdownDescription: "<p>In some cases, the backend server returns a response code that indicates an error.</p><p>When enabled, the default, the service ignores the error condition and processes the response rule. If successful, returns the successful response to the client.</p><p>When disabled, the DataPower service notices the error condition and processes the error rule. If successful, propagates the response code from the backend server to the client.</p>",
							Computed:            true,
						},
						"http_client_ip_label": schema.StringAttribute{
							MarkdownDescription: "Enter the name of an HTTP Header to read to determine the IP address of the requesting client (for example X-Forwarded-For). This value defaults to X-Client-IP.",
							Computed:            true,
						},
						"http_log_cor_id_label": schema.StringAttribute{
							MarkdownDescription: "Enter the name of an HTTP Header to read to determine the global transaction ID for chained services. This value defaults to X-Global-Transaction-ID.",
							Computed:            true,
						},
						"load_balancer_hash_header": schema.StringAttribute{
							MarkdownDescription: "Enter the name of an HTTP Header to use rather than a Client IP address when load balancing traffic to the back end servers using a hash algorithm. The value of the header identified here will be used to calculate a hash rather than the IP address. If blank, the client IP address will be used for hash calculation.",
							Computed:            true,
						},
						"in_order_mode": models.GetDmGatewayInOrderModeDataSourceSchema("Enables in-order (serial) processing of queue-based messages during different parts of transactions through the service. When enabled the service respects the sequential order of the messages when writing them to queues.", "inorder-mode", ""),
						"wsa_mode": schema.StringAttribute{
							MarkdownDescription: "Select the mode of Web Services Addressing (WS-Addressing). The mode specifies the WS-Addressing support that is provided by the DataPower service. The level of support is determined by the associated client and server devices. Support for a particular mode of WS-Addressing does not preclude simultaneous support for traditional addressing formats. WS-Addressing requires SOAP messages. The default is No WS-Addressing.",
							Computed:            true,
						},
						"wsa_require_aaa": schema.BoolAttribute{
							MarkdownDescription: "<p>Require successful execution of an AAA policy before trusting the WS-Addressing <tt>To</tt> SOAP header from the client to select the server URL to connect to.</p><p>This property is relevant when WS-Addressing is enabled and the type is a dynamic-backend (dynamic routing enabled). In this case, the <tt>To</tt> SOAP header determines the default URL for connecting to the remote server.</p><p>When this property is off and the type is dynamic-backend, clients can route SOAP messages to <em>any</em> remote server for this DataPower service, which can be a significant security risk. For this reason, the default is on.</p><p>When this property is on and the type is dynamic-backend, the follow behavior occurs based on the AAA policy. <ul><li>If the AAA policy runs successfully, the <tt>To</tt> address will be accepted.</li><li>If the AAA Policy fails (but continues processing) the back-end address determined from configuration will be used.</li><li>If there is no AAA Policy configured (or executed) in the processing rules, WS-Addressing processing will proceed as if the AAA Policy had executed and failed.</li></ul></p>",
							Computed:            true,
						},
						"wsa_rewrite_reply_to": schema.StringAttribute{
							MarkdownDescription: "<p>Identifies the URL Rewrite Policy used to rewrite the contents of the WS-Addressing <tt>ReplyTo</tt> header of the SOAP request, if present.</p><p>Use this property to modify the contents of an incoming WS-Addressing <tt>ReplyTo</tt> SOAP header that identifies the recipient endpoint of response messages. This typically will be used to change the <tt>ReplyTo</tt> address to one that will send the response to this service.</p><p>This property is relevant when the WS-Addressing Mode is WS-Addressing to Traditional or WS-Addressing to WS-Addressing.</p><p>Use the values-list to select the URL Rewrite Policy used to manipulate the contents of the original <tt>ReplyTo</tt> header.</p>",
							Computed:            true,
						},
						"wsa_rewrite_fault_to": schema.StringAttribute{
							MarkdownDescription: "<p>Identifies the URL Rewrite Policy used to rewrite the contents of the WS-Addressing <tt>FaultTo</tt> header of the SOAP request, if present.</p><p>You can use this property to modify the contents of an incoming WS-Addressing <tt>FaultTo</tt> SOAP header that identifies the recipient endpoint of fault messages. This typically will be used to change the <tt>FaultTo</tt> address to one that will send the response to this service.</p><p>This property is relevant when the WS-Addressing Mode is WS-Addressing to Traditional, or WS-Addressing to WS-Addressing.</p><p>Use the values-list to specify the URL Rewrite Policy used to manipulate the contents of the original <tt>FaultTo</tt> header.</p>",
							Computed:            true,
						},
						"wsa_rewrite_to": schema.StringAttribute{
							MarkdownDescription: "<p>Identifies the URL Rewrite Policy used to rewrite the contents of the WS-Addressing <tt>To</tt> header of the SOAP request, if present.</p><p>You can use this property to modify the contents of an incoming WS-Addressing <tt>To</tt> SOAP header that identifies the message destination. If dynamic routing is not in use, this will typically be used to change the <tt>To</tt> header from an address that matches the front side of this service to one that addresses the back-side server.</p><p>If the <tt>To</tt> header does not match any of the rewrite rules and the type is not dynamic-backend, the <tt>To</tt> will be based on the backend URL.</p><p>This property is relevant when the WS-Addressing Mode is WS-Addressing to Traditional, or WS-Addressing to WS-Addressing.</p><p>Use the values-list to specify the URL Rewrite Policy used to manipulate the contents of the original <tt>To</tt> header.</p>",
							Computed:            true,
						},
						"wsa_strip": schema.BoolAttribute{
							MarkdownDescription: "<p>Control whether to remove WS-Addressing headers when sending a request or response without using WS-Addressing.</p><ul><li>When on,the default setting, the WS-Addressing headers ( <tt>To</tt> , <tt>From</tt> , <tt>ReplyTo</tt> , <tt>FaultTo</tt> , <tt>Action</tt> , <tt>MessageID</tt> , and <tt>RelatesTo</tt> ) are stripped when sending synchronously.</li><li>When off, the WS-Addressing headers will not be removed.</li></ul><p>This property is relevant when the WS-Addressing Mode is WS-Addressing to Traditional, or Traditional to WS-Addressing.</p><p>If the WS-Addressing mode is WS-Addressing to Traditional and this is on, WS-Addressing headers are stripped from incoming client SOAP requests before they are sent to the server.</p><p>If the WS-Addressing mode is Traditional to WS-Addressing and this is on, WS-Addressing headers are stripped from incoming server SOAP responses before they are sent to the client.</p>",
							Computed:            true,
						},
						"wsa_default_reply_to": schema.StringAttribute{
							MarkdownDescription: "<p>Provide a default value of the WS-Addressing <tt>ReplyTo</tt> SOAP header in request messages when Force Incoming WS-Addressing is enabled.</p><p>Because the WS-Addressing specifications do not require the inclusion of the <tt>ReplyTo</tt> header, the service might receive WS-Addressing messages that do not contain a <tt>ReplyTo</tt> header, or that contain the header with no value.</p><p>If the SOAP request has no WS-Addressing <tt>ReplyTo</tt> or <tt>From</tt> SOAP header, this will provide a default value of the <tt>ReplyTo</tt> header. That header identifies the recipient endpoint of the response message.</p><p>When this happens, the service modifies the Web Service Addressing message to include a <tt>ReplyTo</tt> header that contains the URL value specified by this property.</p><p>If a default recipient endpoint of response messages is not explicitly identified by this command, the default value is <tt>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</tt> .</p><p>This property is relevant when the WS-Addressing Mode is WS-Addressing to Traditional or Traditional to WS-Addressing and Force Incoming WS-Addressing is enabled.</p>",
							Computed:            true,
						},
						"wsa_default_fault_to": schema.StringAttribute{
							MarkdownDescription: "<p>Provide a default value of the WS-Addressing <tt>FaultTo</tt> SOAP header in request messages.</p><p>Because the WS-Addressing specifications do not require the inclusion of the <tt>FaultTo</tt> header, the service might receive WS-Addressing requests that do not contain a <tt>FaultTo</tt> SOAP header, or that contain the header with no value.</p><p>If the SOAP request has no WS-Addressing <tt>FaultTo</tt> , <tt>ReplyTo</tt> , or <tt>From</tt> SOAP header, this will provide a default value of the <tt>FaultTo</tt> header. That header identifies the recipient endpoint of an error message.</p><p>If a default recipient endpoint of fault messages is not explicitly identified by this command, the default value is <tt>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</tt></p><p>This property is relevant when the WS-AddressingMode is WS-Addressing to Traditional, or Traditional to WS-Addressing.</p>",
							Computed:            true,
						},
						"wsa_force": schema.BoolAttribute{
							MarkdownDescription: "<p>Force the insertion of WS-Addressing headers into traditionally addressed SOAP requests.</p><p>The service generally handles a mix of Web services addressed and traditionally addressed messages. You can use this property to ensure that all messages use WS-Addressing.</p><p>When enabled (on), the DataPower service converts traditionally addressed messages to the WS-Addressing format through the addition of <tt>To</tt> , <tt>ReplyTo</tt> , <tt>FaultTo</tt> , and <tt>MessageID</tt> WS-Addressing headers to the traditionally addressed message.</p><p>The <tt>ReplyTo</tt> header will be that configured in the Default WS-Addressing Reply-To property. The <tt>FaultTo</tt> header will be that configured in the Default WS-Addressing Fault-To property.</p><p>When disabled (off, the default), there is no insertion of WS-Addressing headers into traditionally addressed messages.</p><p>This property is relevant when the WS-Addressing Mode is WS-Addressing to Traditional, or Traditional to WS-Addressing.</p>",
							Computed:            true,
						},
						"wsa_gen_style": schema.StringAttribute{
							MarkdownDescription: "<p>Select the message exchange model between the DataPower service and a target server that uses WS-Addressing. The default is Synchronous.</p><p>This property is relevant when the WS-Addressing Mode is Traditional to WS-Addressing or WS-Addressing to WS-Addressing.</p>",
							Computed:            true,
						},
						"wsahttp_async_response_code": schema.Int64Attribute{
							MarkdownDescription: "<p>Specify the HTTP Response Code sent to a client device prior to transmitting the actual asynchronous server response.</p><p>If the server response to an HTTP client request is asynchronous, the DataPower service must close the original HTTP client channel with a valid response code. After the channel is closed, the DataPower service forwards the server-generated response or fault message to the client over a new channel.</p><p>Enter a value in the range 200 - 599. The default value is 204, which specifies the HTTP response code used to close the original client channel. While 204 is the default, 202 is more commonly required by current standards.</p><p>This property is relevant when the WS-Addressing mode is either Traditional to WS-Addressing or WS-Addressing to WS-Addressing.</p>",
							Computed:            true,
						},
						"wsa_back_protocol": schema.StringAttribute{
							MarkdownDescription: "<p>Specify the protocol handler that receives asynchronous server responses and forward the response to the original client. This handler must be a protocol handler on which this DataPower service already listens.</p><p>The protocol handler specified by this property can be programmatically overridden by the <tt>var://context/__WSA_REQUEST/replyto</tt> variable, which will change the WS-Addressing <tt>ReplyTo</tt> header on the request to the server.</p><p>This property is relevant when the WS-Addressing mode is either Traditional to WS-Addressing or WS-Addressing to WS-Addressing and the message generation pattern is asynchronous.</p>",
							Computed:            true,
						},
						"wsa_timeout": schema.Int64Attribute{
							MarkdownDescription: "<p>Specify the maximum period of time in seconds to wait for an asynchronous response from the server before abandoning the transaction with an error.</p><p>Use any value of 1 - 4000000. The default value is 120.</p><p>This value can be programmatically overridden by the <tt>var://service/wsa/timeout</tt> variable.</p><p>The establishment of this timer can be programmatically overridden by the <tt>var://service/soap-oneway-mep</tt> variable, which eliminates waiting for a response. This can be applied to messages known to use the SOAP 1.2 one-way message exchange pattern.</p><p>This property is relevant when the WS-Addressing mode is either Traditional to WS-Addressing WS-Addressing to WS-Addressing and the message generation pattern is asynchronous.</p>",
							Computed:            true,
						},
						"wsrm_enabled": schema.BoolAttribute{
							MarkdownDescription: "Enable the Web Services Reliable Messaging feature (WS-ReliableMessaging 1.1). WS-ReliableMessaging requires SOAP messages. The default is No WS-ReliableMessaging.",
							Computed:            true,
						},
						"wsrm_sequence_expiration": schema.Int64Attribute{
							MarkdownDescription: "<p>Set the target expiration interval for all WS-ReliableMessaging Sequences. If an incoming CreateSequence SOAP message has an Expires interval longer than this configured value, it will be reduced to this value in the SequenceResponse SOAP message. The same applies to the Expires interval in any accepted Offer in an incoming CreateSequence. This is also used for the requested Expires value in any CreateSequence SOAP call made to the client or server from a RM Source.</p><p>This implementation will never request or accept a non-expiring sequence (the value \"PT0S\", representing zero seconds).</p><p>This property is relevant only when WS-ReliableMessaging is enabled.</p>",
							Computed:            true,
						},
						"wsrmaaa_policy": schema.StringAttribute{
							MarkdownDescription: "<p>Specify the AAA Policy to perform authentication of incoming WS-ReliableMessaging (WS-RM) messages. While this is focused on protecting the WS-RM control SOAP messages, such as CreateSequence and TerminateSequence, it is also run on incoming WS-RM data messages with a Sequence header. This prevents unauthorized clients from using system resources by issuing CreateSequence requests, or from disrupting existing WS-RM sequences using CloseSequence or TerminateSequence messages, or from falsely acknowledging messages using SequenceAcknowledgement messages. This would normally be the same AAA Policy used in later processing by the request or response rules, the results are cached so it is not evaluated again.</p><p>Note that this also applies to a CreateSequence request received from the server. If the server is not providing this authentication information, configure to have the RM Source make offers.</p><p>This property is relevant only when WS-RM is enabled.</p>",
							Computed:            true,
						},
						"wsrm_destination_accept_create_sequence": schema.BoolAttribute{
							MarkdownDescription: "<p>Accept incoming WS-ReliableMessaging (WS-RM) <tt>CreateSequence</tt> SOAP request, and create an RM Destination when one is received. If this is enabled, both the client and the server can use WS-RM to send messages to this gateway. If disabled, the client cannot use WS-RM to communicate with this DataPower service. If disabled, the only way that an RM Destination can be created on this service is if the RM Source was configured to make offers, in which case an <tt>Offer</tt> and <tt>Accept</tt> can create an RM Destination for the server to send WS-RM messages to.</p><p>This property is relevant only when WS-RM is enabled.</p>",
							Computed:            true,
						},
						"wsrm_destination_maximum_sequences": schema.Int64Attribute{
							MarkdownDescription: "<p>Set a limit on the maximum number of WS-ReliableMessaging (WS-RM) destinations that can be created on this DataPower service by <tt>CreateSequence</tt> SOAP requests. Attempts by clients to create sequences in excess of this limit will result in a SOAP Fault. This controls memory resource utilization.</p><p>This limit does not apply to WS-RM destinations that are created by peer acceptance of offers made in <tt>CreateSequence</tt> SOAP requests sent by this DataPower service.</p><p>This property is relevant only when WS-RM and Destination Accept Incoming CreateSequence are enabled.</p>",
							Computed:            true,
						},
						"wsrm_destination_in_order": schema.BoolAttribute{
							MarkdownDescription: "<p>Enable the <tt>InOrder</tt> delivery assurance for WS-ReliableMessaging (WS-RM) destinations, in addition to the standard <tt>ExactlyOnce</tt> delivery assurance. No messages will be passed from the receive queue for further processing unless their sequence number assigned by the client is one greater than the last one processed.</p><p><tt>InOrder</tt> delivery assurance increases memory and resource use by the WS-RM destination.</p><p>This property is relevant only when WS-RM is enabled and Destination Accept Incoming CreateSequence or Source Make Two-Way Offer are enabled.</p>",
							Computed:            true,
						},
						"wsrm_destination_maximum_in_order_queue_length": schema.Int64Attribute{
							MarkdownDescription: "Sets the maximum length of a RM Destination queue when Destination InOrder Delivery Service is enabled. This is the maximum number of messages held in the queue beyond a gap in the received sequence numbers. This controls memory use.",
							Computed:            true,
						},
						"wsrm_destination_accept_offers": schema.BoolAttribute{
							MarkdownDescription: "<p>Accept offers for two-way WS-ReliableMessaging (WS-RM) contained in CreateSequence SOAP request. If an offer is included, the creation of an WS-RM Destination will also result in the creation of an RM Source to reliably send responses to the client.</p><p>This property is relevant only when WS-RM and Destination Accept Incoming CreateSequence are enabled.</p>",
							Computed:            true,
						},
						"wsrm_front_force": schema.BoolAttribute{
							MarkdownDescription: "<p>Require the use of WS-ReliableMessaging (WS-RM) for all SOAP messages being processed by the request rules. The client must establish a sequence using CreateSequence SOAP call and must include a Sequence in each SOAP header. Any SOAP request messages without a Sequence header will result in a SOAP fault.</p><p>This property is relevant only when WS-RM and Destination Accept Incoming CreateSequence are enabled.</p>",
							Computed:            true,
						},
						"wsrm_back_force": schema.BoolAttribute{
							MarkdownDescription: "<p>Require the use of WS-ReliableMessaging (WS-RM) for all SOAP messages being processed by the response rules. Any SOAP response messages without a Sequence header will result in a SOAP fault.</p><p>When WS-Addressing is in use, SOAP messages without a WS-Addressing RelatesTo SOAP Header will be processed by the request rule, not the response rule, even if they come from the back-side server.</p><p>This property is relevant only when WS-RM is enabled and Destination Accept Incoming CreateSequence or Source Make Two-Way Offer are enabled.</p>",
							Computed:            true,
						},
						"wsrm_back_create_sequence": schema.BoolAttribute{
							MarkdownDescription: "<p>Create an RM Source from the back side to the server. The RM Source is creating by sending a CreateSequence SOAP request to the server address, when there is SOAP data to send and there is no RM Source that was created by a MakeOffer from the server.</p><p>This property is relevant only when WS-RM is enabled.</p>",
							Computed:            true,
						},
						"wsrm_front_create_sequence": schema.BoolAttribute{
							MarkdownDescription: "<p>Create an RM Source from the front side to the client. The RM Source is created by sending a CreateSequence SOAP request to the WS-Addressing <tt>ReplyTo</tt> address, when there is SOAP data to sent to the client and there is no RM Source that was created by a MakeOffer from the client.</p><p>To create this two-way connection, WS-Addressing must be in use on the front side.</p><p>This property is relevant only when WS-ReliableMessaging is enabled and the WS-Addressing mode is WS-Addressing to Traditional or WS-Addressing to WS-Addressing.</p>",
							Computed:            true,
						},
						"wsrm_source_make_offer": schema.BoolAttribute{
							MarkdownDescription: "<p>Include an Offer for two-way WS-ReliableMessaging (WS-RM) in CreateSequence SOAP requests made to the server as a result of request processing. This may result in creating an RM Destination for the server to send responses on when the gateway creates an RM Source to send requests to the server.</p><p>If the server does not Accept the Offer, no RM Destination will be created.</p><p>This property is relevant only when WS-ReliableMessaging and Source Create Sequence on Response are enabled.</p>",
							Computed:            true,
						},
						"wsrm_uses_sequence_ssl": schema.BoolAttribute{
							MarkdownDescription: "<p>Use the WS-ReliableMessaging (WS-RM) SSL/TLS security model for protecting the integrity and security of messages sent from RM Sources for this gateway. This provides confidentiality (through SSL/TLS) and protects the WS-RM state from alteration by unauthorized parties.</p><p>Outgoing CreateSequence control messages will have the UsesSequenceTLS SOAP header block. This will tell the peer RM Destination that all control and data messages for this conversation must use the same SSL/TLS session. Any message to that RM Destination Identifier that is not associated with the same SSL/TLS session will be rejected.</p><p>The lifetime of a SSL/TLS protected Sequence will be limited by the lifetime of the SSL/TLS session that is used to protect that Sequence.</p><p>Acceptance and enforcement of UsesSequenceTLS on incoming WS-RM connections to RM Destinations is automatic if TLS is enabled.</p><p>This property is relevant only when WS-ReliableMessaging and when Source Create Sequence on Request, Source Create Sequence on Response, or Destination Accept Two-Way Offers are enabled.</p>",
							Computed:            true,
						},
						"wsrm_front_acks_to": schema.StringAttribute{
							MarkdownDescription: "<p>Selects a particular front-side protocol handler to request to receive asynchronous acknowledgments and control messages responses on the front side of the gateway. If not configured, they will be requested to be received at the URL of the original request message that causes the creation of the RM Source.</p><p>The handler must be one that is already configured as one of the active ones for this service. It can be the one that incoming requests are arriving on or can be any other one that the client WS-ReliableMessaging (WS-RM) implementation can reach.</p><p>An URL will be generated from the base URL of the protocol handler, which may also have a path added due to the WSDL describing the service.</p><p>This generated URL is used as the value of <tt>CreateSequence/AcksTo</tt> when sending a Create Sequence SOAP request, it is used as the <tt>CreateSequenceResponse/Accept/AcksTo</tt> when accepting an offer from the client. This generated URL is also used as the WS-Addressing <tt>ReplyTo</tt> address for WS-RM control SOAP messages ( <tt>CreateSequence</tt> , <tt>CloseSequence</tt> , and <tt>TerminateSequence</tt> ) sent by the RM Source to the client.</p><p>This property is relevant only when WS-ReliableMessaging and when Source Create Sequence on Response or Destination Accept Two-Way Offers are enabled.</p>",
							Computed:            true,
						},
						"wsrm_back_acks_to": schema.StringAttribute{
							MarkdownDescription: "<p>Selects a particular handler to receive asynchronous WS-ReliableMessaging (WS-RM) <tt>SequenceAcknowledgement</tt> messages. If this is not configured, but there is a WS-Addressing Reply Point configured, it will be used to receive <tt>SequenceAcknowledgements</tt> . If neither is set, the back side RM Source will expect synchronous <tt>SequenceAcknowledgments</tt> in response messages.</p><p>The handler must be one that is already configured as active for this DataPower service. It can be the one that incoming front side requests are arriving on or can be any other one that the server WS-RM implementation can reach.</p><p>When configured, the <tt>AcksTo</tt> field in WS-RM <tt>CreateSequence</tt> SOAP messages sent to the server will be the base URL of this handler. If not configured but WS-Addressing Reply Point is configured, the <tt>AcksTo</tt> field will come from that handler.</p><p>If neither is configured, the <tt>AcksTo</tt> field will be <tt>http://www.w3.org/2005/08/addressing/anonymous</tt> .</p><p>This property is relevant only when WS-ReliableMessaging and when Source Create Sequence on Request or Destination Accept Two-Way Offers are enabled.</p>",
							Computed:            true,
						},
						"wsrm_source_maximum_sequences": schema.Int64Attribute{
							MarkdownDescription: "<p>Set a limit on the maximum number of simultaneously active Sequences from RM Sources of this Gateway. Each remote RM Destination Endpoint Reference (URL) requires one sequence. This controls memory resource utilization.</p><p>Transactions requiring to creating sequences in excess of this limit will result in a SOAP fault.</p><p>Incoming CreateSequence SOAP requests with an Offer will have the offer ignored if accepting the offer would exceed this limit.</p><p>This property is relevant only when WS-ReliableMessaging and when Source Create Sequence on Request, Source Create Sequence on Response, or Destination Accept Two-Way Offers are enabled.</p>",
							Computed:            true,
						},
						"wsrm_source_retransmission_interval": schema.Int64Attribute{
							MarkdownDescription: "<p>How long a RM Source will wait for an Ack before retransmitting the message. This also applies to the retransmission of a CreateSequence, CloseSequence, or TerminateSequence SOAP requests. If Front Exponential Backoff is enabled, this is the initial timeout. Use any value of 2 - 600. The default is 10.</p><p>This property is relevant only when WS-ReliableMessaging and when Source Create Sequence on Request, Source Create Sequence on Response, or Destination Accept Two-Way Offers are enabled.</p>",
							Computed:            true,
						},
						"wsrm_source_exponential_backoff": schema.BoolAttribute{
							MarkdownDescription: "<p>Controls whether exponential backoff is used to increase the interval between subsequent retransmissions on unacknowledged messages by a RM Source.</p><p>This property is relevant only when WS-ReliableMessaging and when Source Create Sequence on Request, Source Create Sequence on Response, or Destination Accept Two-Way Offers are enabled.</p>",
							Computed:            true,
						},
						"wsrm_source_maximum_retransmissions": schema.Int64Attribute{
							MarkdownDescription: "<p>How many times a RM Source will retransmit a message before declaring a failure. This also controls retransmission of CreateSequence, CloseSequence, and TerminateSequence SOAP requests.</p><p>This property is relevant only when WS-ReliableMessaging and when Source Create Sequence on Request, Source Create Sequence on Response, or Destination Accept Two-Way Offers are enabled.</p>",
							Computed:            true,
						},
						"wsrm_source_maximum_queue_length": schema.Int64Attribute{
							MarkdownDescription: "<p>Maximum length of a RM Source queue for responses. This is the maximum number of messages held in the queue awaiting Ack messages. This controls memory use.</p><p>This property is relevant only when WS-ReliableMessaging and when Source Create Sequence on Request, Source Create Sequence on Response, or Destination Accept Two-Way Offers are enabled.</p>",
							Computed:            true,
						},
						"wsrm_source_request_ack_count": schema.Int64Attribute{
							MarkdownDescription: "<p>How many messages the a RM Source will send before including the AckRequested SOAP header to request an acknowledgment. The AckRequested SOAP header is always included when a message is retransmitted.</p><p>This property is relevant only when WS-ReliableMessaging and when Source Create Sequence on Request, Source Create Sequence on Response, or Destination Accept Two-Way Offers are enabled.</p>",
							Computed:            true,
						},
						"wsrm_source_inactivity_close": schema.Int64Attribute{
							MarkdownDescription: "<p>How long the a RM Source will wait for an another message to be sent before closing the sequence by sending a CloseSequence or TerminateSequence SOAP message. Use any value of 10 - 3600. The default is 10.</p><p>CloseSequence is sent when the inactivity timer expires if there are outstanding unacknowledged messages, otherwise Terminate sequence is sent.</p><p>This property is relevant only when WS-ReliableMessaging and when Source Create Sequence on Request, Source Create Sequence on Response, or Destination Accept Two-Way Offers are enabled.</p>",
							Computed:            true,
						},
						"force_policy_exec": schema.BoolAttribute{
							MarkdownDescription: "<p>Some message patterns could include bodyless request and response messages. This approach is common with RESTful Web services where messages might not include a body but still requires the processing policy to run. To enable this capability for services whose request and response type is XML, JSON, or marked as non-XML, set this option to 'on'. By doing so, the processing policy rules will always be run.</p>",
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
							MarkdownDescription: "When enabling the delay of error messages, specify the delay duration in milliseconds. If delaying messages for 3000ms, the DataPower Gateway will not send error messages to the client until 3 seconds have elapsed since the DataPower Gateway performed decryption on the requests. Enter a value in the range 250 - 300000. The default value is 1000.",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *MultiProtocolGatewayDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *MultiProtocolGatewayDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data MultiProtocolGatewayList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.MultiProtocolGateway{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.MultiProtocolGateway{}
	if value := res.Get(`MultiProtocolGateway`); value.Exists() {
		for _, v := range value.Array() {
			item := models.MultiProtocolGateway{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.MultiProtocolGatewayObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.MultiProtocolGatewayObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
