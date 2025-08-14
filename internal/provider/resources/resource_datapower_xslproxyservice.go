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

var _ resource.Resource = &XSLProxyServiceResource{}
var _ resource.ResourceWithValidateConfig = &XSLProxyServiceResource{}

func NewXSLProxyServiceResource() resource.Resource {
	return &XSLProxyServiceResource{}
}

type XSLProxyServiceResource struct {
	pData *tfutils.ProviderData
}

func (r *XSLProxyServiceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_xslproxyservice"
}

func (r *XSLProxyServiceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("<p>The XSL Proxy is obsolete.</p>", "xslproxy", "").AddActions("quiesce").String,
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
			"type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "type", "").AddStringEnum("static-backend", "loopback-proxy", "strict-proxy", "dynamic-backend").AddDefaultValue("static-backend").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("static-backend", "loopback-proxy", "strict-proxy", "dynamic-backend"),
				},
				Default: stringdefault.StaticString("static-backend"),
			},
			"xml_manager": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "xml-manager", "xmlmanager").AddDefaultValue("default").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default"),
			},
			"url_rewrite_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "urlrewrite-policy", "urlrewritepolicy").String,
				Optional:            true,
			},
			"style_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "stylesheet-policy", "stylepolicy").AddDefaultValue("default").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default"),
			},
			"credential_charset": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "credential-charset", "").AddStringEnum("protocol", "ascii", "utf8", "big5", "cp1250", "cp1251", "cp1252", "cp1253", "cp1254", "cp1255", "cp1256", "cp1257", "cp1258", "euc_jp", "euc_kr", "gb18030", "gb2312", "iso2022_jp", "iso2022_kr", "iso8859_1", "iso8859_2", "iso8859_4", "iso8859_5", "iso8859_6", "iso8859_7", "iso8859_8", "iso8859_9", "iso8859_15", "sjis", "tis620", "unicode_le").AddDefaultValue("protocol").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("protocol", "ascii", "utf8", "big5", "cp1250", "cp1251", "cp1252", "cp1253", "cp1254", "cp1255", "cp1256", "cp1257", "cp1258", "euc_jp", "euc_kr", "gb18030", "gb2312", "iso2022_jp", "iso2022_kr", "iso8859_1", "iso8859_2", "iso8859_4", "iso8859_5", "iso8859_6", "iso8859_7", "iso8859_8", "iso8859_9", "iso8859_15", "sjis", "tis620", "unicode_le"),
				},
				Default: stringdefault.StaticString("protocol"),
			},
			"ssl_config_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "ssl-config-type", "").AddStringEnum("server", "sni").AddDefaultValue("server").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("server", "sni"),
				},
				Default: stringdefault.StaticString("server"),
			},
			"ssl_server": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "ssl-server", "sslserverprofile").String,
				Optional:            true,
			},
			"sslsni_server": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "ssl-sni-server", "sslsniserverprofile").String,
				Optional:            true,
			},
			"ssl_client": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "ssl-client", "sslclientprofile").String,
				Optional:            true,
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
			"local_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the local port to monitor for incoming client requests.", "port", "").AddIntegerRange(1, 65535).String,
				Required:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 65535),
				},
			},
			"remote_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the host name or IP address of the specific server supported by this DataPower service. If using load balancers, specify the name of the Load Balancer Group. If using the On Demand Router, specify the keyword ODR-LBG. Load balancer groups and the On Demand Router can be used only when Type is static-backend.", "remote-ip-address", "").String,
				Required:            true,
			},
			"remote_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the port number to monitor. Used only when Type is static-backend.", "remote-port", "").AddIntegerRange(1, 65535).String,
				Required:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 65535),
				},
			},
			"acl": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("This Access Control List will be used to allow or deny access to this service based on the IP address of the client. When attached to a service, an Access Control List (ACL) denies all access by default. To deny access to only selected addresses, first grant access to all addresses (allow 0.0.0.0) and then create deny entries for the desired hosts.", "acl", "accesscontrollist").String,
				Optional:            true,
			},
			"http_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the maximum number of seconds (within the range 1 through 86400, with a default of 120) that a firewall or proxy maintains an idle HTTP connection.", "http-timeout", "").AddIntegerRange(1, 86400).AddDefaultValue("120").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(120),
			},
			"http_persist_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the maximum number of seconds (within the range 0 through 7200, with a default of 180) that a firewall or proxy maintains an idle TCP connection.", "persistent-timeout", "").AddIntegerRange(0, 7200).AddDefaultValue("180").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 7200),
				},
				Default: int64default.StaticInt64(180),
			},
			"do_host_rewrite": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("When enabled, the device will rewrite the Host: header to be the address of the back-end server. This is not what a strict proxy would do, and may not be appropriate for all topologies.", "host-rewriting", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"suppress_http_warnings": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Toggle to enable or disable the generation of Transformation Applied (Warning Code: 214) messages by the HTTP service.", "silence-warning", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_compression": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Toggle to enable or disable the GZIP compression function.", "compression", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_include_response_type_encoding": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Toggle to enable or disable including the character set encoding in the HTTP content-type header produced. For example, when sending a UTF-8 encoded XML document: If this property is disabled, 'content-type=text/xml' will be sent. If this property is enabled, 'content-type=text/xml; charset=UTF-8' will be sent.", "include-response-type-encoding", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"always_show_errors": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("If set, HTTP responses may be generated with errors appended to partially generated documents. If not set error responses will only be sent to the browser if no other output has been created.", "always-show-errors", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"disallow_get": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("If set, only methods with incoming data (such as POST) are allowed.", "disallow-get", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"disallow_empty_response": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("If set, only responses with message bodies are allowed (that is, not 304 and so forth).", "disallow-empty-reply", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_persistent_connections": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Toggle to enable or disable HTTP persistent connections.", "persistent-connections", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"http_client_ip_label": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Retain the default value (X-Client-IP) or provide an other value (for example, X-Forwarded-For).", "client-address", "").AddDefaultValue("X-Client-IP").String,
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
			"http_proxy_host": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the host name or IP address of the HTTP proxy to use when the remote server can be accessed only through another HTTP proxy.", "httpproxy-address", "").String,
				Optional:            true,
			},
			"http_proxy_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the port number on the HTTP proxy server.", "httpproxy-port", "").AddIntegerRange(1, 65535).AddDefaultValue("800").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(800),
			},
			"http_version": models.GetDmHTTPClientServerVersionResourceSchema("HTTP Version", "version", "", false),
			"do_chunked_upload": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use the radio buttons to enable (on) or disable (off) the ability to send Content-Type Chunked Encoded documents to the back end server. When the device employs the HTTP/1.1 protocol, the body of the document can be delimited by either Content-Length or chunked encodings. While all servers will understand how to interpret Content-Length, many applications will fail to understand Chunked encoding. For this reason, Content-Length is the standard method used. However doing so interferes with the ability of the device to fully stream. To stream full documents towards the back end server, this property should be turned on. However, the back end server must be RFC 2616 compatible, because this feature cannot be renegotiated at run time, unlike all other HTTP/1.1 features which can be negotiated down at runtime if necessary. This property can also be enabled by configuring a User Agent to enable it on a per-URL basis.", "chunked-uploads", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"header_injection": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP Header Injection", "inject", "").String,
				NestedObject:        models.DmHeaderInjectionResourceSchema,
				Optional:            true,
			},
			"header_suppression": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP Headers can be suppressed (removed) from the message flow using this property. For example, the Via: header, which contains the name of the DataPower service handling the message, may be suppressed from messages sent by the DataPower device back to the client.", "suppress", "").String,
				NestedObject:        models.DmHeaderSuppressionResourceSchema,
				Optional:            true,
			},
			"stylesheet_parameters": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Stylesheets used in Processing Policies can take stylesheet parameters. These parameters can be passed in by this object. More than one parameter can be defined.", "parameter", "").String,
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
			"force_policy_exec": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Some message patterns may include bodyless request and response messages. This is common with RESTful web services where messages may or may not include a body but still requires the processing policy to run. To enable this capability for services whose request and response type is XML (or marked as non-XML i.e. JSON), set this option to 'on'. By doing so, the processing policy rules will always be executed.</p>", "force-policy-exec", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"count_monitors": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Count Monitors watch for defined messaging events and increment counters each time the event occurs. When a certain threshold is reached, the monitor can either post a notification to a log or block service for a configured amount of time. Select a Count Monitor from the list to activate that monitor for this firewall. Click the + button to create a new Count Monitor.", "monitor-count", "countmonitor").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"duration_monitors": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Duration Monitors watch for events meeting or exceeding a configured duration. When a duration is met or exceeded, the monitor can either post a notification to a log or block service for a configured amount of time. Select a Duration Monitor from the list to activate that monitor for this firewall. Click the + button to create a new Duration Monitor.", "monitor-duration", "durationmonitor").String,
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
			"local_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Enter a host alias or the IP address that the service listens on. Host aliases can ease migration tasks among appliances.</p><ul><li>0 or 0.0.0.0 indicates all configured IPv4 addresses.</li><li>:: indicates all configured IPv4 and IPv6 addresses.</li></ul><p><b>Attention:</b> For management services, the value of 0.0.0.0 or :: is a security risk. Use an explicit IP address to isolate management traffic from application data traffic.</p>", "ip-address", "").AddDefaultValue("0.0.0.0").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("0.0.0.0"),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *XSLProxyServiceResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *XSLProxyServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.XSLProxyService
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `XSLProxyService`)
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

func (r *XSLProxyServiceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.XSLProxyService
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
		data.FromBody(ctx, `XSLProxyService`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `XSLProxyService`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *XSLProxyServiceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.XSLProxyService
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `XSLProxyService`))
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

func (r *XSLProxyServiceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.XSLProxyService
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

func (r *XSLProxyServiceResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.XSLProxyService

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
