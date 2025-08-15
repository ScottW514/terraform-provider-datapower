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

package resources

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &WSGatewayResource{}
var _ resource.ResourceWithValidateConfig = &WSGatewayResource{}

func NewWSGatewayResource() resource.Resource {
	return &WSGatewayResource{}
}

type WSGatewayResource struct {
	pData *tfutils.ProviderData
}

func (r *WSGatewayResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ws_gateway"
}

func (r *WSGatewayResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("<p>A Web Service Proxy virtualizes web services described by WSDL files. The Proxy provides alternate endpoint URLs, message validation, service level monitoring and automatic updating when provided with only one or more WSDL files.</p>", "wsgw", "").AddActions("quiesce").String,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name of the object. Must be unique among object types in application domain.", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"app_domain": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The name of the application domain the object belongs to", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`), ""),
				},
				PlanModifiers: []planmodifier.String{
					modifiers.ImmutableAfterSet(),
				},
			},
			"back_http_version": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the HTTP version to use on the server-side connection. The default is HTTP 1.1.", "http-server-version", "").AddStringEnum("HTTP/1.0", "HTTP/1.1").AddDefaultValue("HTTP/1.1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("HTTP/1.0", "HTTP/1.1"),
				},
				Default: stringdefault.StaticString("HTTP/1.1"),
			},
			"request_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Characterizes the traffic that originates from the client. The default is SOAP.", "request-type", "").AddStringEnum("soap", "xml", "unprocessed", "preprocessed").AddDefaultValue("soap").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("soap", "xml", "unprocessed", "preprocessed"),
				},
				Default: stringdefault.StaticString("soap"),
			},
			"response_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Characterizes the traffic that originates from the server. The default is SOAP.", "response-type", "").AddStringEnum("soap", "xml", "unprocessed", "preprocessed").AddDefaultValue("soap").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("soap", "xml", "unprocessed", "preprocessed"),
				},
				Default: stringdefault.StaticString("soap"),
			},
			"follow_redirects": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Some protocols generate redirects as part of the protocol - for example HTTP response code 302. If this property is enabled the gateway will try and transparently resolve those redirects.", "follow-redirects", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"allow_compression": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable or disable the GZIP compression functionality on the backside.", "compression", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"enable_compressed_payload_passthrough": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable or disable the decompression of the response payload when the backside response is compressed. When enabled, the payload is not decompressed but passed through compressed.", "enable-compressed-payload-passthrough", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the mode of operation for the Web Service Proxy. The default is Static from WSDL.", "type", "").AddStringEnum("static-backend", "dynamic-backend", "static-from-wsdl").AddDefaultValue("static-from-wsdl").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("static-backend", "dynamic-backend", "static-from-wsdl"),
				},
				Default: stringdefault.StaticString("static-from-wsdl"),
			},
			"auto_create_sources": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The front side ports that accept connections must have configured Front Side Protocol Source Handlers. If no Front Side Protocol Handler is configured and auto-create is enabled, a default Protocol Handler will be created. Currently only auto-creates HTTP sources.", "autocreate-sources", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ssl_server_config_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS server type", "ssl-config-type", "").AddStringEnum("server", "sni").AddDefaultValue("server").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("server", "sni"),
				},
				Default: stringdefault.StaticString("server"),
			},
			"ssl_server": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS server profile", "ssl-server", "ssl_server_profile").String,
				Optional:            true,
			},
			"sslsni_server": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS SNI server profile", "ssl-sni-server", "ssl_sni_server_profile").String,
				Optional:            true,
			},
			"endpoint_rewrite_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>An Endpoint Rewrite Policy determines the URL that clients use to send requests for the services described in the WSDL file(s); the URL published to directory services; and can also determine the URL of the back end application services.</p><p>This policy allows you to define service endpoint URLs for more than one local interface. You can define the URLs published to directory services, which may be different than the locally defined interface when load balancers are used. You can override the URLS defined in the WSDL files for the Web Services endpoints, redefining these URLS to a different location. This is useful for maintenance windows.</p><p>A configured Endpoint Rewrite Policy takes precedence over any Local Endpoint Rewrite and Remote Endpoint Rewrite policies.</p>", "endpoint-rewrite-policy", "ws_endpoint_rewrite_policy").String,
				Optional:            true,
			},
			"local_endpoint_rewrite": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify a URL Rewrite Policy used to modify the wsdl:port address specified in the source WSDL when the service is loaded. If no URL Rewrite Policy is specified, the default local address is the IP address of the device and the relative URI and the original port number from the wsdl:port address used in the source WSDL.</p><p>For example, the WSDL may specify <tt>*:2100/search/find</tt> and this URL Rewrite Policy might rewrite it to <tt>*:2100/searchsvc/find</tt> . Client requests to the web service would be rewritten to <tt>http://dpdeviceaddr:2100/searchsvc/find</tt> .</p>", "frontside-port-rewrite", "url_rewrite_policy").String,
				Optional:            true,
			},
			"remote_endpoint_rewrite": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify a URL Rewrite Policy used to modify the wsdl:port address specified in the source WSDL when the service is loaded. If no URL Rewrite Policy is specified, the default remote address is exactly the same wsdl:port address used in the source WSDL.</p><p>For example, the WSDL may specify <tt>http://api.beagle.com:2100/search/find\"</tt> and this URL Rewrite Policy might rewrite it to <tt>http://192.168.2.21:2000/search/find\"</tt> .</p>", "backside-port-rewrite", "url_rewrite_policy").String,
				Optional:            true,
			},
			"aaa_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the AAA policy to provide only authentication and authorization to all messages that the service endpoints handle. Unlike a AAA policy in the processing policy that can modify messages and perform postprocessing, this AAA policy does not employ any selective matching rules. Therefore, this AAA policy processes all messages.", "aaa-policy", "aaa_policy").String,
				Optional:            true,
			},
			"style_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>The service can apply a user-defined processing policy to messages processed by the service. This policy may take a wide range of action on messages. Matching rules select messages for particular processing actions, as with a firewall or proxy service. This policy takes effect in addition to the processing applied automatically by the service.</p><p>Select an existing policy from the list or click the + button to create a new policy.</p><p><b>Note:</b> The processing policy for a Web Service Proxy works with only Web Service Proxy services, not any other service type on the DataPower Gateway.</p>", "stylepolicy", "ws_style_policy").AddDefaultValue("default").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default"),
			},
			"remote_fetch_retry": models.GetDmNetworkRetryResourceSchema("Remote Fetch Retry", "remote-retry", "", false),
			"wsdl_cache_policy": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("A WSDL that is part of the cache policy is refreshed when its TTL to is reached. Refresh retrieves the WSDL from the source location and refreshes the web service proxy state. Depending upon the changes in the WSDL file itself, the Proxy may reconfigure itself in any number of ways, including adding new endpoints or removing existing endpoints.", "wsdl-cache-policy", "").String,
				NestedObject:        models.DmWSDLCachePolicyResourceSchema,
				Optional:            true,
			},
			"base_wsdl": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("A proxy may use one or more WSDL files to define the service the proxy handles. Those WSDL files are defined on this page.", "wsdl", "").String,
				NestedObject:        models.DmWSBaseWSDLResourceSchema,
				Optional:            true,
			},
			"user_toggles": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Each WSDL Operation of the web service proxied by a Web Service Proxy can have a user policy defined for that component.</p><p>Components are specified by the combination of Target Namespace, WSDL file, Service, PortType, Binding and Operation. For example, to specify all operations in the target namespace MySvc, enter \"MySvc\" in the Target Namespace field and set all other inputs to *. To specify only one particular operation, named GetLottoPick for example, enter \"wsdl:definitions//wsdl:operation/@GetLottoPick\" in the Operations field and set all others to *.</p><p>The policy applied to the specified component consists of six options: Enable, Publish, Validate Faults, Validate Headers, Validate Requests and Validate Responses. See below for more information about policy options.</p>", "user-policy", "").String,
				NestedObject:        models.DmWSUserTogglesResourceSchema,
				Optional:            true,
			},
			"client_principal": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the full name of the client principal when the Web Services Proxy needs to decrypt automatically encrypted requests. Use this property when the encryption uses a Kerberos session key or uses a key that was derived from the session key.", "client-principal", "").String,
				Optional:            true,
			},
			"server_principal": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the full name of the server principal when the Web Services Proxy needs to decrypt automatically encrypted requests. Use this property when the encryption uses a Kerberos session key or uses a key that was derived from the session key.", "server-principal", "").String,
				Optional:            true,
			},
			"kerberos_keytab": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the Kerberos keytab file that contains the principals that the Web Services Proxy uses to decrypt automatically encrypted requests.", "kerberos-keytab", "crypto_kerberos_keytab").String,
				Optional:            true,
			},
			"decrypt_key": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The ws-proxy will automatically decrypt encrypted requests. Configure the decrypt key to explicitly select the private key used in the decryption.", "decrypt-key", "crypto_key").String,
				Optional:            true,
			},
			"encrypted_key_sha1_cache_life_time": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("This is the Cache Lifetime for the decrypted generated key. Setting the value to 0 means the decrypted generated key will not be cached.", "wssec11-enckey-cache", "").AddIntegerRange(0, 604800).String,
				Optional:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 604800),
				},
			},
			"preserve_key_chain": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select 'on' to output the chain of elements being used by the decrypted Encrypted Data, such as xenc:EncryptedKey, wsc:DerivedKeyToken. Otherwise all the xenc:EncryptedKey elements will be removed after decryption, even some of the Encrypted Data may not be decrypted successfully.", "preserve-key-chain", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"decrypt_with_key_from_ed": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("In scenarios where the key is inside an EncryptedData element (such as 'encrypted SAML Assertion'), the decrypt action cannot locate the key to decrypt the corresponding EncryptedData elements. Select 'on', to enable decrypt action to attempt decryption with the key that is inside the EncryptedData element.", "decrypt-with-key-from-ed", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"soap_action_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select how to handle the SOAPAction HTTP header.", "soap-action-policy", "").AddStringEnum("lax", "strict", "off").AddDefaultValue("lax").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("lax", "strict", "off"),
				},
				Default: stringdefault.StaticString("lax"),
			},
			"wsrr_subscriptions": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Load and proxy services obtained via these subscriptions.", "wsrr-subscription", "").String,
				NestedObject:        models.DmWSRRWSDLSourceResourceSchema,
				Optional:            true,
			},
			"wsrr_saved_search_subscriptions": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The Web Service Proxy virtualizes Web service endpoints based on the WSDL files returned by the saved search.", "wsrr-saved-search-subscription", "").String,
				NestedObject:        models.DmWSRRSavedSearchWSDLSourceResourceSchema,
				Optional:            true,
			},
			"operation_priority": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Operation Priority", "operation-priority", "").String,
				NestedObject:        models.DmWSOperationSchedulerPriorityResourceSchema,
				Optional:            true,
			},
			"operation_conformance_policy": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Operation Conformance Policy", "operation-conformance", "").String,
				NestedObject:        models.DmWSOperationConformancePolicyResourceSchema,
				Optional:            true,
			},
			"operation_policy_subject_opt_out": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Operation Policy Subject Opt Out", "operation-policy-opt-out", "").String,
				NestedObject:        models.DmWSOperationPolicySubjectOptOutResourceSchema,
				Optional:            true,
			},
			"policy_parameter": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Policy Parameters", "policy-parameters", "").String,
				NestedObject:        models.DmWSPolicyParametersResourceSchema,
				Optional:            true,
			},
			"reliable_messaging": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Reliable Messaging", "reliable-messaging", "").String,
				NestedObject:        models.DmWSOperationReliableMessagingResourceSchema,
				Optional:            true,
			},
			"wsm_agent_monitor": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Control the collection of monitoring records for this service.", "wsmagent-monitor", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"wsm_agent_monitor_pcm": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Select messages to capture for analysis.</p><p>Because not all Web Services Management protocols can accommodate full message-capture, configure this property only if the spooler can forward full messages.</p><p>Full message-capture incurs a performance penalty.</p>", "wsmagent-monitor-capture-mode", "").AddStringEnum("none", "all-messages", "on-error").AddDefaultValue("all-messages").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("none", "all-messages", "on-error"),
				},
				Default: stringdefault.StaticString("all-messages"),
			},
			"process_resp_rules_on_one_way_mep": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "process-response-rule-for-oneway-wsdls", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"priority": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Control the service scheduling priority. When system resources are in high demand, \"high\" priority services will be favored over lower priority services.", "priority", "").AddStringEnum("unknown", "high-min", "high", "high-max", "normal-min", "normal", "normal-max", "low-min", "low", "low-max").AddDefaultValue("normal").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("unknown", "high-min", "high", "high-max", "normal-min", "normal", "normal-max", "low-min", "low", "low-max"),
				},
				Default: stringdefault.StaticString("normal"),
			},
			"front_protocol": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify a protocol handler to determine the network communication protocol, address, port, and other protocol-specific settings.", "front-protocol", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"xml_manager": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>An XML Manager manages the compilation and caching of style sheets, the caching of documents, and provides configuration constraints on the size and parsing depth of documents. You can enable streaming operation by configuring an XML Manager to use a Streaming Compile Option Policy.</p><p>Optionally, an XML Manager can employ a User Agent. The User Agent settings, in turn, can affect the behavior of the gateway when communicating with remote servers or with clients when sending back responses.</p>", "xml-manager", "xml_manager").AddDefaultValue("default").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default"),
			},
			"url_rewrite_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Define a URL Rewrite Policy that defines rules for the following rewrite and replacement operations:</p><ul><li>Rewrite the entire URL or a portion of the URL</li><li>Replace a header value in a message</li><li>Rewrite the HTTP POST body in a message</li></ul><p>The rewrite rules are applied before document processing. Therefore, the evaluation criteria in the Matching Rule is against the rewritten value.</p>", "urlrewrite-policy", "url_rewrite_policy").String,
				Optional:            true,
			},
			"ssl_client_config_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The TLS profile type to secure connections between the DataPower Gateway and its targets.", "ssl-client-type", "").AddStringEnum("client").AddDefaultValue("client").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("client"),
				},
				Default: stringdefault.StaticString("client"),
			},
			"ssl_client": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The TLS client profile to secure connections between the DataPower Gateway and its targets.", "ssl-client", "ssl_client_profile").String,
				Optional:            true,
			},
			"fw_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify a credentials list for the service. The credentials specify which keys and certificates are available to support processing. In the absence of credentials, all locally-stored key and certificates are available to the service.", "fwcred", "crypto_fw_cred").String,
				Optional:            true,
			},
			"header_injection": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP Header Injection", "inject", "").String,
				NestedObject:        models.DmHeaderInjectionResourceSchema,
				Optional:            true,
			},
			"header_suppression": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP Header Suppression", "suppress", "").String,
				NestedObject:        models.DmHeaderSuppressionResourceSchema,
				Optional:            true,
			},
			"stylesheet_parameters": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Style sheets used in processing policies can take stylesheet parameters. These parameters can be passed in by this object. You can define more than one parameter.", "parameter", "").String,
				NestedObject:        models.DmStylesheetParameterResourceSchema,
				Optional:            true,
			},
			"default_param_namespace": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("If a stylesheet parameter is defined without a namespace (or without explicitly specifying the null namespace), then this is the namespace into which the parameter will be assigned.", "default-param-namespace", "").AddDefaultValue("http://www.datapower.com/param/config").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("http://www.datapower.com/param/config"),
			},
			"query_param_namespace": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The namespace in which to put all parameters that are specified in the URL query string.", "query-param-namespace", "").AddDefaultValue("http://www.datapower.com/param/query").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("http://www.datapower.com/param/query"),
			},
			"backend_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the URL of the static backend server. The URL is parsed to determine the protocol to use.</p><p>To use a load balancer, specify the name of an existing Load Balancer Group instead of the address-port pair in the URL.</p><p>To use the on demand router, specify http://ODR-LBG instead of the address-port pair in the URL.</p><ul><li>When the URL starts with http:// or https://, the DataPower service uses the HTTP protocol.</li><li>When the URL starts with amqp-broker://, the DataPower service uses the AMQP protocol.</li><li>When the URL starts with ftp://, the DataPower service uses the FTP protocol. <ul><li>For a relative path to the login directory of the user, specify the URL in the ftp://user:password@host:port/path format.</li><li>For an absolute path to the root directory, specify the URL in the ftp://user:password@host/%2Fpath format. %2F is the URL encoding of a forward slash.</li></ul><p>Include a password in the URL with caution. The use of user:/password@host results in a server connection. However, with this configuration, the connection could be unable to send multiple commands to the FTP server. For a stable connection, define a basic authentication policy in the user agent. The user agent is in the XML manager associated with the DataPower service.</p></li><li>When the URL starts with dpmq://, dpmqfte, idgmq://, or idgmqmft://, the DataPower service uses the IBM MQ protocol. The referenced queue manager in the URL must exist in the current domain.</li><li>When the URL starts with dpims:// or dpimsssl://, the DataPower service uses the IMS protocol. These protocols are available with the IMS feature.</li><li>When the URL starts with dpkafka://, the DataPower service uses the Kafka protocol.</li><li>When the URL starts with dpnfs://, the DataPower service uses the NFS protocol. <ul><li>For static mounts, specify the URL in the dpnfs://mount format, where mount is the name of an existing NFS mount.</li><li>For dynamic mounts, specify the URL in the dpnfs://host or dpnfs://host/path format.</li></ul></li><li>When the URL starts with sftp://, the DataPower service uses the SSH FTP protocol. <ul><li>For an absolute path to the root directory, specify the URL in the sftp://host:port/path format.</li><li>For a relative path to the login directory of the user, specify the URL in the sftp://host:port/~/path format.</li></ul><p>If password authentication is desired, specify the password in the SFTP Client Policy of the User Agent. If Public key authentication is desired, specify the user's private key in the SFTP Client Policy of the User Agent. When no SFTP Client Policy is specified, the Basic-Auth Policy and Pubkey-Auth Policy take effect.</p></li><li>When the URL starts with dptibems://, the DataPower service uses the TIBCO EMS protocol. The referenced server in the URL must exist. These protocols are available with the TIBCO EMS feature.</li><li>When the URL starts with dpwasjms://, the DataPower service uses the WebSphere JMS protocol. The referenced server in the URL must exist.</li></ul><p>For assistance in building a URL, click the appropriate Helper button.</p>", "backend-url", "").String,
				Optional:            true,
			},
			"propagate_uri": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>URI propagation is meaningful only when:</p><ul><li>The service uses a static backend.</li><li>The service uses a dynamic backend and routing is set with a route with stylesheet (route-action) action. In this case, use dp:set-target() to define the backend server. <p>For the other routing options that are available with route-action and route-set, the URI is absolute.</p></li></ul><p>When set to <b>on</b>, the service rewrites the URI of the backend URL to the URI in the client request. If the client submits http://host/service and the backend URL is http://server/listener, the URL is rewritten to http://server/service.</p><p>If the backend URL employs AMQP, IBM MQ, Kafka, TIBCO EMS, or WebSphere JMS, set to <b>off</b>.</p>", "propagate-uri", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"service_monitors": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies a list of Web Services Monitors for the service. Web Services Monitors target Web Services endpoints. Use the Web Services wizard on the Control Panel to supply a WSDL and configure a Monitor for an endpoint. The Monitor collects statistics, establishes count and duration monitors and can take action when thresholds are met or exceeded. A Monitor must be included here to be activated. Note that this is not the same Monitor as a Service Level Monitor (SLM) that can be included in the Processing Policy of this service.", "monitor-service", "web_service_monitor").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"count_monitors": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the count monitors to define messaging events and increment counters each time the event occurs. When a certain threshold is reached, the monitor can either post a notification to a log or block service for a configured amount of time.", "monitor-count", "count_monitor").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"duration_monitors": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration monitors to watch for events that meet or exceed a configured duration. When a duration is met or exceeded, the monitor can either post a notification to a log or block service for a configured amount of time.", "monitor-duration", "duration_monitor").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"monitor_processing_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the way that the system behaves when more than one monitor is attached to a service.", "monitor-processing-policy", "").AddStringEnum("terminate-at-first-throttle", "terminate-at-first-match").AddDefaultValue("terminate-at-first-throttle").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("terminate-at-first-throttle", "terminate-at-first-match"),
				},
				Default: stringdefault.StaticString("terminate-at-first-throttle"),
			},
			"request_attachments": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select how to treat client requests with attachments. The default is strip.", "request-attachments", "").AddStringEnum("strip", "reject", "allow", "streaming", "unprocessed").AddDefaultValue("strip").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("strip", "reject", "allow", "streaming", "unprocessed"),
				},
				Default: stringdefault.StaticString("strip"),
			},
			"response_attachments": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select how to treat server responses with attachments. The default is strip.", "response-attachments", "").AddStringEnum("strip", "reject", "allow", "streaming", "unprocessed").AddDefaultValue("strip").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("strip", "reject", "allow", "streaming", "unprocessed"),
				},
				Default: stringdefault.StaticString("strip"),
			},
			"request_attachments_flow_control": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable or disable flow control for attachments in requests to servers. The default is off.", "request-attachments-flow-control", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"response_attachments_flow_control": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable or disable flow control for attachments in responses to clients. The default is off.", "response-attachments-flow-control", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"root_part_not_first_action": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("When streaming MIME messages, the action to take when the root part is not the first part of the message. If the root part must be first (for example to do BSP conformance checking) and the action is set to \"process in order\" the DP device will buffer attachments up to the root.", "root-part-not-first-action", "").AddStringEnum("process-in-order", "buffer", "abort").AddDefaultValue("process-in-order").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("process-in-order", "buffer", "abort"),
				},
				Default: stringdefault.StaticString("process-in-order"),
			},
			"front_attachment_format": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select how to interpret client requests with attachments.", "front-attachment-format", "").AddStringEnum("dynamic", "mime", "dime", "detect").AddDefaultValue("dynamic").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("dynamic", "mime", "dime", "detect"),
				},
				Default: stringdefault.StaticString("dynamic"),
			},
			"back_attachment_format": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the attachment output format to backend servers.", "back-attachment-format", "").AddStringEnum("dynamic", "mime", "dime", "detect").AddDefaultValue("dynamic").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("dynamic", "mime", "dime", "detect"),
				},
				Default: stringdefault.StaticString("dynamic"),
			},
			"mime_front_headers": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>The body of a message (that is, the payload, independent of any protocol headers) can sometimes contain MIME headers before any preamble and before the first MIME boundary contained in the body of the message. These MIME headers can contain important information not available in the protocol headers, such as the string identifying the MIME boundary. If this property is set to on, then these MIME headers will be processed by the service.</p><p>Note that if this is on and there are no MIME headers contained in the message, the device will continue to try and parse the message, using the protocol header information, if available.</p><p>When this is off and MIME headers are present in the body of the message, these MIME headers will be considered part of the preamble and not used to parse out the message. If the protocol headers (such as HTTP) indicate the MIME boundaries, the device can parse and process individual attachments. If such information is not available, no attachments can be parsed from the body of the message.</p>", "mime-front-headers", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"mime_back_headers": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>The body of a message (the payload independent of protocol headers) can sometimes contain MIME headers before any preamble and before the first MIME boundary in the body of the message. These MIME headers can contain important information that is not available in the protocol headers, such as a string that identifies the MIME boundary. If this property is set to on, these MIME headers are processed by the service.</p><p>When set to on and there are no MIME headers in the message, the device continues to try and parse the message, using the protocol header information, if available.</p><p>When set to off and MIME headers are present in the body of the message, these MIME headers are considered part of the preamble and not used to parse the message. If protocol headers, such as HTTP, indicate MIME boundaries, the device can parse and process individual attachments. If such information is not available, no attachments can be parsed from the body of the message.</p>", "mime-back-headers", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"stream_output_to_back": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the desired streaming behavior.", "stream-output-to-back", "").AddStringEnum("buffer-until-verification", "stream-until-infraction").AddDefaultValue("buffer-until-verification").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("buffer-until-verification", "stream-until-infraction"),
				},
				Default: stringdefault.StaticString("buffer-until-verification"),
			},
			"stream_output_to_front": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the desired streaming behavior.", "stream-output-to-front", "").AddStringEnum("buffer-until-verification", "stream-until-infraction").AddDefaultValue("buffer-until-verification").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("buffer-until-verification", "stream-until-infraction"),
				},
				Default: stringdefault.StaticString("buffer-until-verification"),
			},
			"max_message_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specifies the maximum allowable size in kilobytes of a message. Enter a value in the range 0 - 2097151. The default value is 0. A value of 0 specifies that no limit is enforced, and the message can be of unlimited size.</p><p>The maximum message size limit applies to JSON, SOAP, XML, and non-XML messages. If the message type is pass-through, no limit is enforced.</p>", "max-message-size", "").AddIntegerRange(0, 2097151).String,
				Optional:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 2097151),
				},
			},
			"gateway_parser_limits": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use the server-defined parser limits instead of the parser limits in the XML manager for this DataPower service. Service-defined limits override XML Manager limits.", "gateway-parser-limits", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"parser_limits_element_depth": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Defines the maximum depth of element nesting in XML parser. If any of the parser limits are set in the DataPower service, they will override those on the XML Manager.", "element-depth", "").AddDefaultValue("512").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(512),
			},
			"parser_limits_attribute_count": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Defines the maximum number of attributes of a given element. If any of the parser limits are set in the service, they will override those on the XML Manager.", "attribute-count", "").AddDefaultValue("128").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(128),
			},
			"parser_limits_max_node_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Defines the maximum size any one node may consume. The default is 32 MB. Sizes which are powers of two result in the best performance. If any of the parser limits are set in the DataPower service, they will override those on the XML Manager. Although you set an explicit value, the DataPower Gateway uses a value that is the rounded-down, largest power of two that is smaller than the defined value.", "max-node-size", "").AddIntegerRange(1024, 4294967295).AddDefaultValue("33554432").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1024, 4294967295),
				},
				Default: int64default.StaticInt64(33554432),
			},
			"parser_limits_external_references": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the special handling for input documents that contain external references, such as an external entity or external DTD definition.", "external-references", "").AddStringEnum("forbid", "ignore", "allow").AddDefaultValue("forbid").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("forbid", "ignore", "allow"),
				},
				Default: stringdefault.StaticString("forbid"),
			},
			"parser_limits_max_prefixes": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter an integer that defines the maximum number of distinct XML namespace prefixes in a document. This limit counts multiple prefixes defined for the same namespace, but does not count multiple namespaces defined in different parts of the input document under a single prefix. Enter a value in the range 0 - 262143. The default value is 1024. A value of 0 indicates that the limit is 1024.", "max-prefixes", "").AddDefaultValue("1024").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(1024),
			},
			"parser_limits_max_namespaces": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter an integer that defines the maximum number of distinct XML namespace URIs in a document. This limit counts all XML namespaces, regardless of how many prefixes are used to declare them. Enter a value in the range 0 - 65535. The default value is 1024. A value of 0 indicates that the limit is 1024.", "max-namespaces", "").AddDefaultValue("1024").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(1024),
			},
			"parser_limits_max_local_names": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter an integer that defines the maximum number of distinct XML local names in a document. This limit counts all local names, independent of the namespaces they are associated with. Enter a value in the range 0 - 1048575. The default value is 60000. A value of 0 indicates that the limit is 60000.", "max-local-names", "").AddDefaultValue("60000").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(60000),
			},
			"parser_limits_attachment_byte_count": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Defines the maximum number of bytes allowed in any single attachment. Attachments that exceed this size will result in a failure of the whole transaction. If this value is 0, no limit is enforced.", "attachment-byte-count", "").AddDefaultValue("2000000000").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(2000000000),
			},
			"parser_limits_attachment_package_byte_count": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Defines the maximum number of bytes allowed for all parts of an attachment package, including the root part. Attachment packages that exceed this size will result in a failure of the whole transaction. If this value is 0, no limit is enforced.", "attachment-package-byte-count", "").String,
				Optional:            true,
			},
			"debug_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Select the diagnostic mode for processing policies. When enabled, you can view details about the state of variables and contexts for a captured transaction in the probe. The default value is <tt>off</tt> .</p><p>Transaction diagnostic mode is not intended for use in a production environment. Transaction diagnostic mode consumes significant resources that can slow down transaction processing.</p>", "debug-mode", "").AddStringEnum("on", "off", "unbounded").AddDefaultValue("off").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("on", "off", "unbounded"),
				},
				Default: stringdefault.StaticString("off"),
			},
			"debug_history": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Set the number of records for transaction diagnostic mode in the probe. Enter a value in the range 10 - 250. The default value is 25.", "debug-history", "").AddIntegerRange(10, 250).AddDefaultValue("25").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(10, 250),
				},
				Default: int64default.StaticInt64(25),
			},
			"debug_trigger": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The probe captures transactions that meet one or more of the conditions defined by the triggers. These triggers examine the direction or type of the message flow and examine the message for an XPath expression match. When a message meets one of these conditions, the transaction is captured in diagnostics mode and becomes part of the list of transactions that can be viewed.", "debug-trigger", "").String,
				NestedObject:        models.DmMSDebugTriggerTypeResourceSchema,
				Optional:            true,
			},
			"flow_control": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Use flow control to manage the transmission of large files when the front side and back end have different latencies and throughput while in streaming mode. This option is available when the request type and the response type are Pass-Thru or Non-XML. By default, flow control is disabled.</p><p>Note that the streaming behavior for output to back and output to front must be set to Stream Messages. Also, if either the front side or the back end will use the HTTP(S) protocol, allow chunked uploads must be set to on.</p>", "flowcontrol", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"soap_schema_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("When a firewall is in SOAP mode, either on the request or response side, it will validate the incoming messages against a W3C Schema that defines what a SOAP message is supposed to look like. It is possible to customize which schema is used on a per-firewall basis by changing this property; this can be used to accommodate non-standard configurations or other special cases.", "soap-schema-url", "").AddDefaultValue("store:///schemas/soap-envelope.xsd").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("store:///schemas/soap-envelope.xsd"),
			},
			"front_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Set the intra-transaction timeout for service-to-client connections. This value is the maximum idle time to allow in a transaction on the service-to-client connection. This timer monitors idle time in the data transfer process. If the specified idle time is exceeded, the connection is torn down. Enter a value in the range 1 - 86400. The default value is 120.", "front-timeout", "").AddIntegerRange(1, 86400).AddDefaultValue("120").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(120),
			},
			"back_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Set the intra-transaction timeout for service-to-server connections. This value is the maximum idle time to allow in a transaction on the service-to-server connection. This timer monitors the idle time in the data transfer process. If the specified idle time is exceeded, the connection is torn down. Enter a value in the range 1 - 86400. The default value is 120.", "back-timeout", "").AddIntegerRange(1, 86400).AddDefaultValue("120").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(120),
			},
			"front_persistent_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Set the inter-transaction timeout for service-to-client connections. This value is the maximum idle time to allow between the completion of a TCP transaction and the initiation of a new TCP transaction on the service-to-client connection. If the specified idle time is exceeded, the connection is torn down. Enter a value in the range 0 - 86400. The default value is 180. A value of 0 disables persistent connections.", "front-persistent-timeout", "").AddIntegerRange(0, 86400).AddDefaultValue("180").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 86400),
				},
				Default: int64default.StaticInt64(180),
			},
			"back_persistent_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the maximum inter-transaction idle time in seconds for service-to-server connections. This value is the maximum idle time between the completion of a TCP transaction and the initiation of a new TCP transaction on this connection. If the specified idle time is exceeded, the connection is torn down. Enter a value in the range 0 - 86400. The default value is 180. A value of 0 disables persistent connections.</p><p><b>Note:</b> For HTTP GET and HEAD requests, the service attempts the connection again after the specified value, Therefore, the actual timeout is twice the specified value.</p><p>An idle TCP connection can remain in the idle state for 20 seconds after the expiration of the persistence timer.</p>", "back-persistent-timeout", "").AddIntegerRange(0, 86400).AddDefaultValue("180").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 86400),
				},
				Default: int64default.StaticInt64(180),
			},
			"include_response_type_encoding": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable or disable including the character set encoding in any content-type header or description produced. For example, when sending a UTF-8 encoded XML document: If this property is disabled, 'text/xml' will be sent. If this property is enabled, 'text/xml; charset=UTF-8' will be sent.", "include-content-type-encoding", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"persistent_connections": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable or disable persistent connections where appropriate on back side.", "persistent-connections", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"loop_detection": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Some protocols provide a loop detection mechanism that can be used to detect gateway loops. This is a good policy, but it can cause each gateway to be publicly recorded in the transmitted message, which might be considered an information leak.", "loop-detection", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"do_host_rewriting": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Some protocols have distinct name-based elements, separate from the URL, to demultiplex. HTTP uses the <tt>Host</tt> header for this purposes. If this feature is enabled, the remote server will receive a request that reflects the final route; otherwise, it will receive a request that reflects the information as it arrived at the DataPower service. Web servers that issue redirects might want to disable this feature, as they often depend on the <tt>Host</tt> header for the value of their redirect.", "host-rewriting", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"do_chunked_upload": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable (on) or disable (off) the ability to send Content-Type Chunked Encoded documents to the back end server. When the device employs the HTTP/1.1 protocol, the body of the document can be delimited by either Content-Length or chunked encodings. While all servers will understand how to interpret Content-Length, many applications will fail to understand Chunked encoding. For this reason, Content-Length is the standard method used. However doing so interferes with the ability of the device to fully stream. To stream full documents towards the back end server, this property should be turned on. However, the back end server must be RFC 2616 compatible, because this feature cannot be renegotiated at run time, unlike all other HTTP/1.1 features which can be negotiated down at runtime if necessary. This property can also be enabled by configuring a User Agent to enable it on a per-URL basis.", "chunked-uploads", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"process_http_errors": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>In some cases, the backend server returns a response code that indicates an error.</p><p>When enabled, the default, the service ignores the error condition and processes the response rule. If successful, returns the successful response to the client.</p><p>When disabled, the DataPower service notices the error condition and processes the error rule. If successful, propagates the response code from the backend server to the client.</p>", "process-http-errors", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"http_client_ip_label": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of an HTTP Header to read to determine the IP address of the requesting client (for example X-Forwarded-For). This value defaults to X-Client-IP.", "http-client-ip-label", "").AddDefaultValue("X-Client-IP").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("X-Client-IP"),
			},
			"http_log_cor_id_label": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of an HTTP Header to read to determine the global transaction ID for chained services. This value defaults to X-Global-Transaction-ID.", "http-global-tranID-label", "").AddDefaultValue("X-Global-Transaction-ID").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("X-Global-Transaction-ID"),
			},
			"load_balancer_hash_header": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of an HTTP Header to use rather than a Client IP address when load balancing traffic to the back end servers using a hash algorithm. The value of the header identified here will be used to calculate a hash rather than the IP address. If blank, the client IP address will be used for hash calculation.", "load-balancer-hash-header", "").String,
				Optional:            true,
			},
			"in_order_mode": models.GetDmGatewayInOrderModeResourceSchema("Enables in-order (serial) processing of queue-based messages during different parts of transactions through the service. When enabled the service respects the sequential order of the messages when writing them to queues.", "inorder-mode", "", false),
			"wsa_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the mode of Web Services Addressing (WS-Addressing). The mode specifies the WS-Addressing support that is provided by the DataPower service. The level of support is determined by the associated client and server devices. Support for a particular mode of WS-Addressing does not preclude simultaneous support for traditional addressing formats. WS-Addressing requires SOAP messages. The default is No WS-Addressing.", "wsa-mode", "").AddStringEnum("wsa2sync", "sync2wsa", "wsa2wsa", "sync2sync").AddDefaultValue("sync2sync").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("wsa2sync", "sync2wsa", "wsa2wsa", "sync2sync"),
				},
				Default: stringdefault.StaticString("sync2sync"),
			},
			"wsa_require_aaa": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Require successful execution of an AAA policy before trusting the WS-Addressing <tt>To</tt> SOAP header from the client to select the server URL to connect to.</p><p>This property is relevant when WS-Addressing is enabled and the type is a dynamic-backend (dynamic routing enabled). In this case, the <tt>To</tt> SOAP header determines the default URL for connecting to the remote server.</p><p>When this property is off and the type is dynamic-backend, clients can route SOAP messages to <em>any</em> remote server for this DataPower service, which can be a significant security risk. For this reason, the default is on.</p><p>When this property is on and the type is dynamic-backend, the follow behavior occurs based on the AAA policy. <ul><li>If the AAA policy runs successfully, the <tt>To</tt> address will be accepted.</li><li>If the AAA Policy fails (but continues processing) the back-end address determined from configuration will be used.</li><li>If there is no AAA Policy configured (or executed) in the processing rules, WS-Addressing processing will proceed as if the AAA Policy had executed and failed.</li></ul></p>", "wsa-require-aaa", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"wsa_rewrite_reply_to": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Identifies the URL Rewrite Policy used to rewrite the contents of the WS-Addressing <tt>ReplyTo</tt> header of the SOAP request, if present.</p><p>Use this property to modify the contents of an incoming WS-Addressing <tt>ReplyTo</tt> SOAP header that identifies the recipient endpoint of response messages. This typically will be used to change the <tt>ReplyTo</tt> address to one that will send the response to this service.</p><p>This property is relevant when the WS-Addressing Mode is WS-Addressing to Traditional or WS-Addressing to WS-Addressing.</p><p>Use the values-list to select the URL Rewrite Policy used to manipulate the contents of the original <tt>ReplyTo</tt> header.</p>", "wsa-replyto-rewrite", "url_rewrite_policy").String,
				Optional:            true,
			},
			"wsa_rewrite_fault_to": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Identifies the URL Rewrite Policy used to rewrite the contents of the WS-Addressing <tt>FaultTo</tt> header of the SOAP request, if present.</p><p>You can use this property to modify the contents of an incoming WS-Addressing <tt>FaultTo</tt> SOAP header that identifies the recipient endpoint of fault messages. This typically will be used to change the <tt>FaultTo</tt> address to one that will send the response to this service.</p><p>This property is relevant when the WS-Addressing Mode is WS-Addressing to Traditional, or WS-Addressing to WS-Addressing.</p><p>Use the values-list to specify the URL Rewrite Policy used to manipulate the contents of the original <tt>FaultTo</tt> header.</p>", "wsa-faultto-rewrite", "url_rewrite_policy").String,
				Optional:            true,
			},
			"wsa_rewrite_to": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Identifies the URL Rewrite Policy used to rewrite the contents of the WS-Addressing <tt>To</tt> header of the SOAP request, if present.</p><p>You can use this property to modify the contents of an incoming WS-Addressing <tt>To</tt> SOAP header that identifies the message destination. If dynamic routing is not in use, this will typically be used to change the <tt>To</tt> header from an address that matches the front side of this service to one that addresses the back-side server.</p><p>If the <tt>To</tt> header does not match any of the rewrite rules and the type is not dynamic-backend, the <tt>To</tt> will be based on the backend URL.</p><p>This property is relevant when the WS-Addressing Mode is WS-Addressing to Traditional, or WS-Addressing to WS-Addressing.</p><p>Use the values-list to specify the URL Rewrite Policy used to manipulate the contents of the original <tt>To</tt> header.</p>", "wsa-to-rewrite", "url_rewrite_policy").String,
				Optional:            true,
			},
			"wsa_strip": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Control whether to remove WS-Addressing headers when sending a request or response without using WS-Addressing.</p><ul><li>When on,the default setting, the WS-Addressing headers ( <tt>To</tt> , <tt>From</tt> , <tt>ReplyTo</tt> , <tt>FaultTo</tt> , <tt>Action</tt> , <tt>MessageID</tt> , and <tt>RelatesTo</tt> ) are stripped when sending synchronously.</li><li>When off, the WS-Addressing headers will not be removed.</li></ul><p>This property is relevant when the WS-Addressing Mode is WS-Addressing to Traditional, or Traditional to WS-Addressing.</p><p>If the WS-Addressing mode is WS-Addressing to Traditional and this is on, WS-Addressing headers are stripped from incoming client SOAP requests before they are sent to the server.</p><p>If the WS-Addressing mode is Traditional to WS-Addressing and this is on, WS-Addressing headers are stripped from incoming server SOAP responses before they are sent to the client.</p>", "wsa-strip-headers", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"wsa_default_reply_to": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Provide a default value of the WS-Addressing <tt>ReplyTo</tt> SOAP header in request messages when Force Incoming WS-Addressing is enabled.</p><p>Because the WS-Addressing specifications do not require the inclusion of the <tt>ReplyTo</tt> header, the service might receive WS-Addressing messages that do not contain a <tt>ReplyTo</tt> header, or that contain the header with no value.</p><p>If the SOAP request has no WS-Addressing <tt>ReplyTo</tt> or <tt>From</tt> SOAP header, this will provide a default value of the <tt>ReplyTo</tt> header. That header identifies the recipient endpoint of the response message.</p><p>When this happens, the service modifies the Web Service Addressing message to include a <tt>ReplyTo</tt> header that contains the URL value specified by this property.</p><p>If a default recipient endpoint of response messages is not explicitly identified by this command, the default value is <tt>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</tt> .</p><p>This property is relevant when the WS-Addressing Mode is WS-Addressing to Traditional or Traditional to WS-Addressing and Force Incoming WS-Addressing is enabled.</p>", "wsa-default-replyto", "").AddDefaultValue("http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous"),
			},
			"wsa_default_fault_to": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Provide a default value of the WS-Addressing <tt>FaultTo</tt> SOAP header in request messages.</p><p>Because the WS-Addressing specifications do not require the inclusion of the <tt>FaultTo</tt> header, the service might receive WS-Addressing requests that do not contain a <tt>FaultTo</tt> SOAP header, or that contain the header with no value.</p><p>If the SOAP request has no WS-Addressing <tt>FaultTo</tt> , <tt>ReplyTo</tt> , or <tt>From</tt> SOAP header, this will provide a default value of the <tt>FaultTo</tt> header. That header identifies the recipient endpoint of an error message.</p><p>If a default recipient endpoint of fault messages is not explicitly identified by this command, the default value is <tt>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</tt></p><p>This property is relevant when the WS-AddressingMode is WS-Addressing to Traditional, or Traditional to WS-Addressing.</p>", "wsa-default-faultto", "").AddDefaultValue("http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous"),
			},
			"wsa_force": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Force the insertion of WS-Addressing headers into traditionally addressed SOAP requests.</p><p>The service generally handles a mix of Web services addressed and traditionally addressed messages. You can use this property to ensure that all messages use WS-Addressing.</p><p>When enabled (on), the DataPower service converts traditionally addressed messages to the WS-Addressing format through the addition of <tt>To</tt> , <tt>ReplyTo</tt> , <tt>FaultTo</tt> , and <tt>MessageID</tt> WS-Addressing headers to the traditionally addressed message.</p><p>The <tt>ReplyTo</tt> header will be that configured in the Default WS-Addressing Reply-To property. The <tt>FaultTo</tt> header will be that configured in the Default WS-Addressing Fault-To property.</p><p>When disabled (off, the default), there is no insertion of WS-Addressing headers into traditionally addressed messages.</p><p>This property is relevant when the WS-Addressing Mode is WS-Addressing to Traditional, or Traditional to WS-Addressing.</p>", "wsa-force", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wsa_gen_style": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Select the message exchange model between the DataPower service and a target server that uses WS-Addressing. The default is Synchronous.</p><p>This property is relevant when the WS-Addressing Mode is Traditional to WS-Addressing or WS-Addressing to WS-Addressing.</p>", "wsa-genstyle", "").AddStringEnum("sync", "async", "oob").AddDefaultValue("sync").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("sync", "async", "oob"),
				},
				Default: stringdefault.StaticString("sync"),
			},
			"wsahttp_async_response_code": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the HTTP Response Code sent to a client device prior to transmitting the actual asynchronous server response.</p><p>If the server response to an HTTP client request is asynchronous, the DataPower service must close the original HTTP client channel with a valid response code. After the channel is closed, the DataPower service forwards the server-generated response or fault message to the client over a new channel.</p><p>Enter a value in the range 200 - 599. The default value is 204, which specifies the HTTP response code used to close the original client channel. While 204 is the default, 202 is more commonly required by current standards.</p><p>This property is relevant when the WS-Addressing mode is either Traditional to WS-Addressing or WS-Addressing to WS-Addressing.</p>", "wsa-http-async-response-code", "").AddIntegerRange(200, 599).AddDefaultValue("204").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(200, 599),
				},
				Default: int64default.StaticInt64(204),
			},
			"wsa_back_protocol": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the protocol handler that receives asynchronous server responses and forward the response to the original client. This handler must be a protocol handler on which this DataPower service already listens.</p><p>The protocol handler specified by this property can be programmatically overridden by the <tt>var://context/__WSA_REQUEST/replyto</tt> variable, which will change the WS-Addressing <tt>ReplyTo</tt> header on the request to the server.</p><p>This property is relevant when the WS-Addressing mode is either Traditional to WS-Addressing or WS-Addressing to WS-Addressing and the message generation pattern is asynchronous.</p>", "wsa-back-protocol", "").String,
				Optional:            true,
			},
			"wsa_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the maximum period of time in seconds to wait for an asynchronous response from the server before abandoning the transaction with an error.</p><p>Use any value of 1 - 4000000. The default value is 120.</p><p>This value can be programmatically overridden by the <tt>var://service/wsa/timeout</tt> variable.</p><p>The establishment of this timer can be programmatically overridden by the <tt>var://service/soap-oneway-mep</tt> variable, which eliminates waiting for a response. This can be applied to messages known to use the SOAP 1.2 one-way message exchange pattern.</p><p>This property is relevant when the WS-Addressing mode is either Traditional to WS-Addressing WS-Addressing to WS-Addressing and the message generation pattern is asynchronous.</p>", "wsa-timeout", "").AddIntegerRange(1, 4000000).AddDefaultValue("120").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 4000000),
				},
				Default: int64default.StaticInt64(120),
			},
			"wsrm_enabled": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable the Web Services Reliable Messaging feature (WS-ReliableMessaging 1.1). WS-ReliableMessaging requires SOAP messages. The default is No WS-ReliableMessaging.", "wsrm", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wsrm_sequence_expiration": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Set the target expiration interval for all WS-ReliableMessaging Sequences. If an incoming CreateSequence SOAP message has an Expires interval longer than this configured value, it will be reduced to this value in the SequenceResponse SOAP message. The same applies to the Expires interval in any accepted Offer in an incoming CreateSequence. This is also used for the requested Expires value in any CreateSequence SOAP call made to the client or server from a RM Source.</p><p>This implementation will never request or accept a non-expiring sequence (the value \"PT0S\", representing zero seconds).</p><p>This property is relevant only when WS-ReliableMessaging is enabled.</p>", "wsrm-sequence-expiration", "").AddIntegerRange(10, 86400).AddDefaultValue("3600").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(10, 86400),
				},
				Default: int64default.StaticInt64(3600),
			},
			"wsrmaaa_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the AAA Policy to perform authentication of incoming WS-ReliableMessaging (WS-RM) messages. While this is focused on protecting the WS-RM control SOAP messages, such as CreateSequence and TerminateSequence, it is also run on incoming WS-RM data messages with a Sequence header. This prevents unauthorized clients from using system resources by issuing CreateSequence requests, or from disrupting existing WS-RM sequences using CloseSequence or TerminateSequence messages, or from falsely acknowledging messages using SequenceAcknowledgement messages. This would normally be the same AAA Policy used in later processing by the request or response rules, the results are cached so it is not evaluated again.</p><p>Note that this also applies to a CreateSequence request received from the server. If the server is not providing this authentication information, configure to have the RM Source make offers.</p><p>This property is relevant only when WS-RM is enabled.</p>", "wsrm-aaa-policy", "aaa_policy").String,
				Optional:            true,
			},
			"wsrm_destination_accept_create_sequence": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Accept incoming WS-ReliableMessaging (WS-RM) <tt>CreateSequence</tt> SOAP request, and create an RM Destination when one is received. If this is enabled, both the client and the server can use WS-RM to send messages to this gateway. If disabled, the client cannot use WS-RM to communicate with this DataPower service. If disabled, the only way that an RM Destination can be created on this service is if the RM Source was configured to make offers, in which case an <tt>Offer</tt> and <tt>Accept</tt> can create an RM Destination for the server to send WS-RM messages to.</p><p>This property is relevant only when WS-RM is enabled.</p>", "wsrm-destination-accept-create-sequence", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"wsrm_destination_maximum_sequences": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Set a limit on the maximum number of WS-ReliableMessaging (WS-RM) destinations that can be created on this DataPower service by <tt>CreateSequence</tt> SOAP requests. Attempts by clients to create sequences in excess of this limit will result in a SOAP Fault. This controls memory resource utilization.</p><p>This limit does not apply to WS-RM destinations that are created by peer acceptance of offers made in <tt>CreateSequence</tt> SOAP requests sent by this DataPower service.</p><p>This property is relevant only when WS-RM and Destination Accept Incoming CreateSequence are enabled.</p>", "wsrm-destination-maximum-sequences", "").AddIntegerRange(1, 2048).AddDefaultValue("400").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 2048),
				},
				Default: int64default.StaticInt64(400),
			},
			"wsrm_destination_in_order": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Enable the <tt>InOrder</tt> delivery assurance for WS-ReliableMessaging (WS-RM) destinations, in addition to the standard <tt>ExactlyOnce</tt> delivery assurance. No messages will be passed from the receive queue for further processing unless their sequence number assigned by the client is one greater than the last one processed.</p><p><tt>InOrder</tt> delivery assurance increases memory and resource use by the WS-RM destination.</p><p>This property is relevant only when WS-RM is enabled and Destination Accept Incoming CreateSequence or Source Make Two-Way Offer are enabled.</p>", "wsrm-destination-inorder", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wsrm_destination_maximum_in_order_queue_length": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Sets the maximum length of a RM Destination queue when Destination InOrder Delivery Service is enabled. This is the maximum number of messages held in the queue beyond a gap in the received sequence numbers. This controls memory use.", "wsrm-destination-maximum-inorder-queue-length", "").AddIntegerRange(1, 256).AddDefaultValue("10").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 256),
				},
				Default: int64default.StaticInt64(10),
			},
			"wsrm_destination_accept_offers": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Accept offers for two-way WS-ReliableMessaging (WS-RM) contained in CreateSequence SOAP request. If an offer is included, the creation of an WS-RM Destination will also result in the creation of an RM Source to reliably send responses to the client.</p><p>This property is relevant only when WS-RM and Destination Accept Incoming CreateSequence are enabled.</p>", "wsrm-destination-accept-offers", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wsrm_front_force": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Require the use of WS-ReliableMessaging (WS-RM) for all SOAP messages being processed by the request rules. The client must establish a sequence using CreateSequence SOAP call and must include a Sequence in each SOAP header. Any SOAP request messages without a Sequence header will result in a SOAP fault.</p><p>This property is relevant only when WS-RM and Destination Accept Incoming CreateSequence are enabled.</p>", "wsrm-request-force", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wsrm_back_force": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Require the use of WS-ReliableMessaging (WS-RM) for all SOAP messages being processed by the response rules. Any SOAP response messages without a Sequence header will result in a SOAP fault.</p><p>When WS-Addressing is in use, SOAP messages without a WS-Addressing RelatesTo SOAP Header will be processed by the request rule, not the response rule, even if they come from the back-side server.</p><p>This property is relevant only when WS-RM is enabled and Destination Accept Incoming CreateSequence or Source Make Two-Way Offer are enabled.</p>", "wsrm-response-force", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wsrm_back_create_sequence": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Create an RM Source from the back side to the server. The RM Source is creating by sending a CreateSequence SOAP request to the server address, when there is SOAP data to send and there is no RM Source that was created by a MakeOffer from the server.</p><p>This property is relevant only when WS-RM is enabled.</p>", "wsrm-source-request-create-sequence", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wsrm_front_create_sequence": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Create an RM Source from the front side to the client. The RM Source is created by sending a CreateSequence SOAP request to the WS-Addressing <tt>ReplyTo</tt> address, when there is SOAP data to sent to the client and there is no RM Source that was created by a MakeOffer from the client.</p><p>To create this two-way connection, WS-Addressing must be in use on the front side.</p><p>This property is relevant only when WS-ReliableMessaging is enabled and the WS-Addressing mode is WS-Addressing to Traditional or WS-Addressing to WS-Addressing.</p>", "wsrm-source-response-create-sequence", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wsrm_source_make_offer": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Include an Offer for two-way WS-ReliableMessaging (WS-RM) in CreateSequence SOAP requests made to the server as a result of request processing. This may result in creating an RM Destination for the server to send responses on when the gateway creates an RM Source to send requests to the server.</p><p>If the server does not Accept the Offer, no RM Destination will be created.</p><p>This property is relevant only when WS-ReliableMessaging and Source Create Sequence on Response are enabled.</p>", "wsrm-source-make-offer", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wsrm_uses_sequence_ssl": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Use the WS-ReliableMessaging (WS-RM) SSL/TLS security model for protecting the integrity and security of messages sent from RM Sources for this gateway. This provides confidentiality (through SSL/TLS) and protects the WS-RM state from alteration by unauthorized parties.</p><p>Outgoing CreateSequence control messages will have the UsesSequenceTLS SOAP header block. This will tell the peer RM Destination that all control and data messages for this conversation must use the same SSL/TLS session. Any message to that RM Destination Identifier that is not associated with the same SSL/TLS session will be rejected.</p><p>The lifetime of a SSL/TLS protected Sequence will be limited by the lifetime of the SSL/TLS session that is used to protect that Sequence.</p><p>Acceptance and enforcement of UsesSequenceTLS on incoming WS-RM connections to RM Destinations is automatic if TLS is enabled.</p><p>This property is relevant only when WS-ReliableMessaging and when Source Create Sequence on Request, Source Create Sequence on Response, or Destination Accept Two-Way Offers are enabled.</p>", "wsrm-source-sequence-ssl", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wsrm_front_acks_to": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Selects a particular front-side protocol handler to request to receive asynchronous acknowledgments and control messages responses on the front side of the gateway. If not configured, they will be requested to be received at the URL of the original request message that causes the creation of the RM Source.</p><p>The handler must be one that is already configured as one of the active ones for this service. It can be the one that incoming requests are arriving on or can be any other one that the client WS-ReliableMessaging (WS-RM) implementation can reach.</p><p>An URL will be generated from the base URL of the protocol handler, which may also have a path added due to the WSDL describing the service.</p><p>This generated URL is used as the value of <tt>CreateSequence/AcksTo</tt> when sending a Create Sequence SOAP request, it is used as the <tt>CreateSequenceResponse/Accept/AcksTo</tt> when accepting an offer from the client. This generated URL is also used as the WS-Addressing <tt>ReplyTo</tt> address for WS-RM control SOAP messages ( <tt>CreateSequence</tt> , <tt>CloseSequence</tt> , and <tt>TerminateSequence</tt> ) sent by the RM Source to the client.</p><p>This property is relevant only when WS-ReliableMessaging and when Source Create Sequence on Response or Destination Accept Two-Way Offers are enabled.</p>", "wsrm-source-front-acks-to", "").String,
				Optional:            true,
			},
			"wsrm_back_acks_to": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Selects a particular handler to receive asynchronous WS-ReliableMessaging (WS-RM) <tt>SequenceAcknowledgement</tt> messages. If this is not configured, but there is a WS-Addressing Reply Point configured, it will be used to receive <tt>SequenceAcknowledgements</tt> . If neither is set, the back side RM Source will expect synchronous <tt>SequenceAcknowledgments</tt> in response messages.</p><p>The handler must be one that is already configured as active for this DataPower service. It can be the one that incoming front side requests are arriving on or can be any other one that the server WS-RM implementation can reach.</p><p>When configured, the <tt>AcksTo</tt> field in WS-RM <tt>CreateSequence</tt> SOAP messages sent to the server will be the base URL of this handler. If not configured but WS-Addressing Reply Point is configured, the <tt>AcksTo</tt> field will come from that handler.</p><p>If neither is configured, the <tt>AcksTo</tt> field will be <tt>http://www.w3.org/2005/08/addressing/anonymous</tt> .</p><p>This property is relevant only when WS-ReliableMessaging and when Source Create Sequence on Request or Destination Accept Two-Way Offers are enabled.</p>", "wsrm-source-back-acks-to", "").String,
				Optional:            true,
			},
			"wsrm_source_maximum_sequences": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Set a limit on the maximum number of simultaneously active Sequences from RM Sources of this Gateway. Each remote RM Destination Endpoint Reference (URL) requires one sequence. This controls memory resource utilization.</p><p>Transactions requiring to creating sequences in excess of this limit will result in a SOAP fault.</p><p>Incoming CreateSequence SOAP requests with an Offer will have the offer ignored if accepting the offer would exceed this limit.</p><p>This property is relevant only when WS-ReliableMessaging and when Source Create Sequence on Request, Source Create Sequence on Response, or Destination Accept Two-Way Offers are enabled.</p>", "wsrm-source-maximum-sequences", "").AddIntegerRange(1, 2048).AddDefaultValue("400").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 2048),
				},
				Default: int64default.StaticInt64(400),
			},
			"wsrm_source_retransmission_interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>How long a RM Source will wait for an Ack before retransmitting the message. This also applies to the retransmission of a CreateSequence, CloseSequence, or TerminateSequence SOAP requests. If Front Exponential Backoff is enabled, this is the initial timeout. Use any value of 2 - 600. The default is 10.</p><p>This property is relevant only when WS-ReliableMessaging and when Source Create Sequence on Request, Source Create Sequence on Response, or Destination Accept Two-Way Offers are enabled.</p>", "wsrm-source-retransmission-interval", "").AddIntegerRange(2, 600).AddDefaultValue("10").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(2, 600),
				},
				Default: int64default.StaticInt64(10),
			},
			"wsrm_source_exponential_backoff": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Controls whether exponential backoff is used to increase the interval between subsequent retransmissions on unacknowledged messages by a RM Source.</p><p>This property is relevant only when WS-ReliableMessaging and when Source Create Sequence on Request, Source Create Sequence on Response, or Destination Accept Two-Way Offers are enabled.</p>", "wsrm-source-exponential-backoff", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"wsrm_source_maximum_retransmissions": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>How many times a RM Source will retransmit a message before declaring a failure. This also controls retransmission of CreateSequence, CloseSequence, and TerminateSequence SOAP requests.</p><p>This property is relevant only when WS-ReliableMessaging and when Source Create Sequence on Request, Source Create Sequence on Response, or Destination Accept Two-Way Offers are enabled.</p>", "wsrm-source-retransmit-count", "").AddIntegerRange(1, 256).AddDefaultValue("4").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 256),
				},
				Default: int64default.StaticInt64(4),
			},
			"wsrm_source_maximum_queue_length": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Maximum length of a RM Source queue for responses. This is the maximum number of messages held in the queue awaiting Ack messages. This controls memory use.</p><p>This property is relevant only when WS-ReliableMessaging and when Source Create Sequence on Request, Source Create Sequence on Response, or Destination Accept Two-Way Offers are enabled.</p>", "wsrm-source-maximum-queue-length", "").AddIntegerRange(1, 256).AddDefaultValue("30").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 256),
				},
				Default: int64default.StaticInt64(30),
			},
			"wsrm_source_request_ack_count": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>How many messages the a RM Source will send before including the AckRequested SOAP header to request an acknowledgment. The AckRequested SOAP header is always included when a message is retransmitted.</p><p>This property is relevant only when WS-ReliableMessaging and when Source Create Sequence on Request, Source Create Sequence on Response, or Destination Accept Two-Way Offers are enabled.</p>", "wsrm-source-request-ack-count", "").AddIntegerRange(1, 256).AddDefaultValue("1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 256),
				},
				Default: int64default.StaticInt64(1),
			},
			"wsrm_source_inactivity_close": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>How long the a RM Source will wait for an another message to be sent before closing the sequence by sending a CloseSequence or TerminateSequence SOAP message. Use any value of 10 - 3600. The default is 10.</p><p>CloseSequence is sent when the inactivity timer expires if there are outstanding unacknowledged messages, otherwise Terminate sequence is sent.</p><p>This property is relevant only when WS-ReliableMessaging and when Source Create Sequence on Request, Source Create Sequence on Response, or Destination Accept Two-Way Offers are enabled.</p>", "wsrm-source-inactivity-close-interval", "").AddIntegerRange(10, 3600).AddDefaultValue("360").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(10, 3600),
				},
				Default: int64default.StaticInt64(360),
			},
			"force_policy_exec": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Some message patterns could include bodyless request and response messages. This approach is common with RESTful Web services where messages might not include a body but still requires the processing policy to run. To enable this capability for services whose request and response type is XML, JSON, or marked as non-XML, set this option to 'on'. By doing so, the processing policy rules will always be run.</p>", "force-policy-exec", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"rewrite_errors": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Error messages after a decryption action can provide an attacker who is using the padding oracle attack method with enough information to determine the contents of the plain-text data. When enabled, the default, the client receives error messages without the internal information that could lead to a discovery. When disabled, the client receives the original message with this information.", "rewrite-errors", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"delay_errors": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The timing difference of the error messages returned after a decryption action can provide an attacker with enough information to determine the contents of the plain-text data. When enabled, the default, the DataPower Gateway delays error messages for the defined duration. When disabled, the DataPower Gateway does not delay error messages.", "delay-errors", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"delay_errors_duration": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("When enabling the delay of error messages, specify the delay duration in milliseconds. If delaying messages for 3000ms, the DataPower Gateway will not send error messages to the client until 3 seconds have elapsed since the DataPower Gateway performed decryption on the requests. Enter a value in the range 250 - 300000. The default value is 1000.", "delay-errors-duration", "").AddIntegerRange(250, 300000).AddDefaultValue("1000").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(250, 300000),
				},
				Default: int64default.StaticInt64(1000),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *WSGatewayResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *WSGatewayResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WSGateway
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `WSGateway`)
	_, err := r.pData.Client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "POST", err))
		return
	}
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WSGatewayResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WSGateway
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && (strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}

	if data.IsNull() {
		// Import
		data.FromBody(ctx, `WSGateway`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `WSGateway`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WSGatewayResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WSGateway
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `WSGateway`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Update)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WSGatewayResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WSGateway
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *WSGatewayResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.WSGateway

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
