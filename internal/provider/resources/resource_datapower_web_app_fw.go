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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

var (
	_ resource.Resource                   = &WebAppFWResource{}
	_ resource.ResourceWithValidateConfig = &WebAppFWResource{}
	_ resource.ResourceWithImportState    = &WebAppFWResource{}
)

func NewWebAppFWResource() resource.Resource {
	return &WebAppFWResource{}
}

type WebAppFWResource struct {
	pData *tfutils.ProviderData
}

func (r *WebAppFWResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_web_app_fw"
}

func (r *WebAppFWResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("The web application firewall provides filtering, security, and input validation for HTTP transactions.", "web-application-firewall", "").AddActions("quiesce").String,
		Attributes: map[string]schema.Attribute{
			"provider_target": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Target host for this resource. If not set, provider will use the top level settings.", "", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name of the object. Must be unique among object types in application domain.", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
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
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
				PlanModifiers: []planmodifier.String{
					modifiers.ImmutableAfterSet(),
				},
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"priority": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Control the service scheduling priority. When system resources are in high demand, \"high\" priority services will be favored over lower priority services.", "priority", "").AddStringEnum("unknown", "high-min", "high", "high-max", "normal-min", "normal", "normal-max", "low-min", "low", "low-max").AddDefaultValue("normal").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("unknown", "high-min", "high", "high-max", "normal-min", "normal", "normal-max", "low-min", "low", "low-max"),
				},
				Default: stringdefault.StaticString("normal"),
			},
			"front_side": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Source Addresses", "listen-on", "").String,
				NestedObject:        models.GetDmFrontSideResourceSchema(),
				Optional:            true,
			},
			"remote_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the host name or IP address of the backend server. If the backend requires TLS, ensure that the TLS client profile is configured for client (forward) connections.", "remote-address", "").String,
				Required:            true,
			},
			"remote_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the port on the remote server.", "remote-port", "").AddIntegerRange(1, 65535).AddDefaultValue("80").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(80),
			},
			"style_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select a security policy from the list of available policies. Click + to create a new Policy. This policy controls which profiles are enforced.", "security-policy", "app_security_policy").String,
				Required:            true,
			},
			"xml_manager": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>An XML Manager manages the compilation and caching of stylesheets, the caching of documents, and provides configuration constraints on the size and parsing depth of documents. You can enable streaming operation by configuring an XML Manager to use a Streaming Compile Option Policy.</p><p>Note that an XML Manager can optionally employ a User Agent. The User Agent settings, in turn, can affect the behavior of the gateway when communicating with back end servers or with clients when sending back responses.</p><p>More than one firewall may use the same XML Manager. Select an existing XML Manager from the list to assign the Manager to this firewall. Click the + button to create a new XML Manager that is assigned to this firewall. A default Manager is used if you do not create one.</p>", "xml-manager", "xml_manager").AddDefaultValue("default").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("default"),
			},
			"error_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("If any policy is violated, this is the default policy that will handle the response sent to the client. It might be overridden in the profile that fails.", "error-policy", "web_app_error_handling_policy").String,
				Optional:            true,
			},
			"uri_normalization": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("If this property is enabled the URI is rewritten to make sure the URI is RFC compliant by escaping certain characters. Additionally, characters that are escaped that do not need to be are unescaped. This makes checking for attack sequences such as .. more reliable.", "uri-normalization", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"rewrite_errors": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Error messages after a decryption action can provide an attacker who is using the padding oracle attack method with enough information to determine the contents of the plain-text data. When enabled, the default, the client receives error messages without the internal information that could lead to a discovery. When disabled, the client receives the original message with this information.", "rewrite-errors", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"delay_errors": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The timing difference of the error messages returned after a decryption action can provide an attacker with enough information to determine the contents of the plain-text data. When enabled, the default, the DataPower Gateway delays error messages for the defined duration. When disabled, the DataPower Gateway does not delay error messages.", "delay-errors", "").AddDefaultValue("true").AddNotValidWhen(models.WebAppFWDelayErrorsIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"delay_errors_duration": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("When enabling the delay of error messages, specify the delay duration in milliseconds. If delaying messages for 3000ms, the DataPower Gateway will not send error messages to the client until 3 seconds have elapsed since the DataPower Gateway performed decryption on the requests. Use any value of 250 - 300000. The default value is 1000.", "delay-errors-duration", "").AddIntegerRange(250, 300000).AddDefaultValue("1000").AddRequiredWhen(models.WebAppFWDelayErrorsDurationCondVal.String()).AddNotValidWhen(models.WebAppFWDelayErrorsDurationIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(250, 300000),
					validators.ConditionalRequiredInt64(models.WebAppFWDelayErrorsDurationCondVal, models.WebAppFWDelayErrorsDurationIgnoreVal, true),
				},
				Default: int64default.StaticInt64(1000),
			},
			"stream_output_to_back": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the desired streaming behavior.", "stream-output-to-back", "").AddStringEnum("buffer-until-verification", "stream-until-infraction").AddDefaultValue("buffer-until-verification").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("buffer-until-verification", "stream-until-infraction"),
				},
				Default: stringdefault.StaticString("buffer-until-verification"),
			},
			"stream_output_to_front": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the desired streaming behavior.", "stream-output-to-front", "").AddStringEnum("buffer-until-verification", "stream-until-infraction").AddDefaultValue("buffer-until-verification").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("buffer-until-verification", "stream-until-infraction"),
				},
				Default: stringdefault.StaticString("buffer-until-verification"),
			},
			"front_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Set the intra-transaction timeout for Web Application Firewall to client connections. This value is the maximum idle time to allow in a transaction on the Web Application Firewall to client connection. This timer monitors idle time in the data transfer process. If the specified idle time is exceeded, the connection is torn down. Enter a value in the range 1 - 86400. The default value is 120.", "front-timeout", "").AddIntegerRange(1, 86400).AddDefaultValue("120").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(120),
			},
			"back_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Set the intra-transaction timeout for Web Application Firewall to server connections. This value is the maximum idle time to allow in a transaction on the Web Application Firewall to server connection. This timer monitors the idle time in the data transfer process. If the specified idle time is exceeded, the connection is torn down. Enter a value in the range 1 - 86400. The default value is 120.", "back-timeout", "").AddIntegerRange(1, 86400).AddDefaultValue("120").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(120),
			},
			"front_persistent_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Set the inter-transaction timeout for Web Application Firewall to client connections. This value is the maximum idle time to allow between the completion of a TCP transaction and the initiation of a new TCP transaction on the Web Application Firewall to client connection. If the specified idle time is exceeded, the connection is torn down. Enter a value in the range 0 - 86400. The default value is 180. A value of 0 disables persistent connections.", "front-persistent-timeout", "").AddIntegerRange(0, 86400).AddDefaultValue("180").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 86400),
				},
				Default: int64default.StaticInt64(180),
			},
			"allow_cache_control_header": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to allow the HTTP GET method to pass the Cache-Control header through to the back end. If disabled, a \"Cache-Control:no-transform\" header is passed. If enabled, the client request can specify the cache control behavior, or if the client request does not specify the Cache-Control header, a \"Cache-Control:no-transform\" header is passed.", "allow-cache-control", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"back_persistent_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the maximum inter-transaction idle time in seconds for Web Application Firewall to server connections. This value is the maximum idle time between the completion of a TCP transaction and the initiation of a new TCP transaction on this connection. If the specified idle time is exceeded, the connection is torn down. Enter a value in the range 0 - 86400. The default value is 180. A value of 0 disables persistent connections.</p><p><b>Note:</b> For HTTP GET and HEAD requests, the service attempts the connection again after the specified value, Therefore, the actual timeout is twice the specified value.</p><p>An idle TCP connection can remain in the idle state for 20 seconds after the expiration of the persistence timer.</p>", "back-persistent-timeout", "").AddIntegerRange(0, 86400).AddDefaultValue("180").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 86400),
				},
				Default: int64default.StaticInt64(180),
			},
			"front_http_version": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the HTTP version to be use on client responses. Incoming version 1.0 requests will always be replied to with 1.0 compatible responses regardless of this setting. The default is HTTP 1.1.", "http-front-version", "").AddStringEnum("HTTP/1.0", "HTTP/1.1").AddDefaultValue("HTTP/1.1").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("HTTP/1.0", "HTTP/1.1"),
				},
				Default: stringdefault.StaticString("HTTP/1.1"),
			},
			"back_http_version": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the HTTP version to use on the server-side connection. The default is HTTP 1.1.", "http-back-version", "").AddStringEnum("HTTP/1.0", "HTTP/1.1").AddDefaultValue("HTTP/1.1").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("HTTP/1.0", "HTTP/1.1"),
				},
				Default: stringdefault.StaticString("HTTP/1.1"),
			},
			"request_side_security": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("If this property is disabled, no request side security policies will be enforced and all requests will be allowed through.", "request-security", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"response_side_security": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("If this property is disabled, no response side security policies will be enforced and all responses will be allowed through.", "response-security", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"do_chunked_upload": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use the radio buttons to enable (on) or disable (off) the ability to send Content-Type Chunked Encoded documents to the back end server. When the device employs the HTTP/1.1 protocol, the body of the document can be delimited by either Content-Length or chunked encodings. While all servers will understand how to interpret Content-Length, many applications will fail to understand Chunked encoding. For this reason, Content-Length is the standard method used. However doing so interferes with the ability of the device to fully stream. To stream full documents towards the back end server, this property should be turned on. However, the back end server must be RFC 2616 compatible, because this feature cannot be renegotiated at run time, unlike all other HTTP/1.1 features which can be negotiated down at runtime if necessary. This property can also be enabled by configuring a User Agent to enable it on a per-URL basis.", "chunked-uploads", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"follow_redirects": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Some protocols generate redirects as part of the protocol - for example HTTP response code 302. If this property is enabled the firewall will try and transparently resolve those redirects.", "follow-redirects", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"http_client_ip_label": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Retain the default value (X-Client-IP) or provide an other value (for example, X-Forwarded-For).", "http-client-ip-label", "").AddDefaultValue("X-Client-IP").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("X-Client-IP"),
			},
			"http_log_cor_id_label": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of an HTTP Header to read to determine the global transaction ID for chained services. This value defaults to X-Global-Transaction-ID.", "http-global-tranID-label", "").AddDefaultValue("X-Global-Transaction-ID").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("X-Global-Transaction-ID"),
			},
			"debug_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Select the diagnostic mode for processing policies. When enabled, you can view details about the state of variables and contexts for a captured transaction in the probe. The default value is <tt>off</tt> .</p><p>Transaction diagnostic mode is not intended for use in a production environment. Transaction diagnostic mode consumes significant resources that can slow down transaction processing.</p>", "debug-mode", "").AddStringEnum("on", "off", "unbounded").AddDefaultValue("off").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("on", "off", "unbounded"),
				},
				Default: stringdefault.StaticString("off"),
			},
			"debug_history": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Set the number of records for transaction diagnostics in the probe. Enter a value in the range 10 - 250. The default value is 25.", "debug-history", "").AddIntegerRange(10, 250).AddDefaultValue("25").AddRequiredWhen(models.WebAppFWDebugHistoryCondVal.String()).AddNotValidWhen(models.WebAppFWDebugHistoryIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(10, 250),
					validators.ConditionalRequiredInt64(models.WebAppFWDebugHistoryCondVal, models.WebAppFWDebugHistoryIgnoreVal, true),
				},
				Default: int64default.StaticInt64(25),
			},
			"debug_trigger": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The probe captures transactions that meet one or more of the conditions defined by the triggers. These triggers examine the direction or type of the message flow and examine the message for an XPath expression match. When a message meets one of these conditions, the transaction is captured in diagnostics mode and becomes part of the list of transactions that can be viewed.", "debug-trigger", "").AddNotValidWhen(models.WebAppFWDebugTriggerIgnoreVal.String()).String,
				NestedObject:        models.GetDmMSDebugTriggerTypeResourceSchema(),
				Optional:            true,
			},
			"url_rewrite_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("A Rewrite Policy can be used to modify the values of certain headers before the security policies are executed. One common use for this is to rewrite the contents of an HTTP Location header to reflect external DNS names. By using a Rewrite Policy, it is possible to intercept and rewrite URLs that might send a browser directly to the back end web site, rather than through the Web Application Firewall.", "urlrewrite-policy", "url_rewrite_policy").String,
				Optional:            true,
			},
			"do_host_rewriting": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Some protocols have distinct name based elements, separate from the URL, to de multiplex. HTTP uses the Host header for this purposes. If this feature is enabled the backside server will receive a request reflecting the final route, otherwise it will receive a request reflecting the information as it arrived at the DataPower device. Web servers issuing redirects may want to disable this feature, as they often depend on the host header for the value of their redirect.", "host-rewriting", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ssl_config_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS type", "ssl-config-type", "").AddStringEnum("server", "sni").AddDefaultValue("server").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("server", "sni"),
				},
				Default: stringdefault.StaticString("server"),
			},
			"ssl_server": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS server profile", "ssl-server", "ssl_server_profile").AddNotValidWhen(models.WebAppFWSSLServerIgnoreVal.String()).String,
				Optional:            true,
			},
			"ssl_sni_server": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS SNI server profile", "ssl-sni-server", "ssl_sni_server_profile").AddNotValidWhen(models.WebAppFWSSLSNIServerIgnoreVal.String()).String,
				Optional:            true,
			},
			"ssl_client": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "ssl_client_profile").AddNotValidWhen(models.WebAppFWSSLClientIgnoreVal.String()).String,
				Optional:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *WebAppFWResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *WebAppFWResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WebAppFW
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !r.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), r.pData.Client.GetTargetNames()))
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `WebAppFW`)
	_, err := r.pData.Client.Post(data.GetPath(), body, data.ProviderTarget)
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), body, data.ProviderTarget)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Resource already exists. Failed to update resource, got error: %s", err))
				return
			}
			resp.Diagnostics.AddWarning("Warning", "Resource already exists. Existing resource was updated.")
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create resource, got error: %s", err))
			return
		}
	}
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}

	tfutils.SaveDomain(ctx, &resp.Diagnostics, r.pData.Client, data.AppDomain, data.ProviderTarget)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WebAppFWResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WebAppFW
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !r.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), r.pData.Client.GetTargetNames()))
		return
	}
	res, err := r.pData.Client.Get(data.GetPath()+"/"+data.Id.ValueString(), data.ProviderTarget)
	if err != nil {
		if strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400") {
			resp.State.RemoveResource(ctx)
			return
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.State.RemoveResource(ctx)
			}
			return
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
			return
		}
	}

	data.UpdateFromBody(ctx, `WebAppFW`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WebAppFWResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WebAppFW
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !r.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), r.pData.Client.GetTargetNames()))
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `WebAppFW`), data.ProviderTarget)
	if err != nil {
		if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
			return
		}
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Update, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}

	tfutils.SaveDomain(ctx, &resp.Diagnostics, r.pData.Client, data.AppDomain, data.ProviderTarget)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WebAppFWResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WebAppFW
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !r.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), r.pData.Client.GetTargetNames()))
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Delete(data.GetPath()+"/"+data.Id.ValueString(), data.ProviderTarget)
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			resp.Diagnostics.AddWarning("Resource Conflict", fmt.Sprintf("Resource is no longer tracked by Terraform, but may need to be manually deleted on DataPower host. Got error: %s", err))
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete resource, got error: %s", err))
			return
		}
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}

	tfutils.SaveDomain(ctx, &resp.Diagnostics, r.pData.Client, data.AppDomain, data.ProviderTarget)

	resp.State.RemoveResource(ctx)
}

func (r *WebAppFWResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError("Invalid Import ID", "Expected format: <app_domain>/<id>. Got: "+req.ID)
		return
	}

	appDomain := parts[0]
	id := parts[1]

	if !regexp.MustCompile("^[a-zA-Z0-9_-]+$").MatchString(id) || len(id) < 1 || len(id) > 128 {
		resp.Diagnostics.AddError("Invalid ID", "ID must be 1-128 characters and match pattern ^[a-zA-Z0-9_-]+$")
		return
	}
	if !regexp.MustCompile("^[a-zA-Z0-9_-]+$").MatchString(appDomain) || len(appDomain) < 1 || len(appDomain) > 128 {
		resp.Diagnostics.AddError("Invalid Application Domain", "Application domain must be 1-128 characters and match pattern ^[a-zA-Z0-9_-]+$")
		return
	}

	var data models.WebAppFW
	data.AppDomain = types.StringValue(appDomain)
	data.Id = types.StringValue(id)
	res, err := r.pData.Client.Get(data.GetPath()+"/"+data.Id.ValueString(), data.ProviderTarget)
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("Resource was not found, got error: %s", err))
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		}
		return
	}

	data.FromBody(ctx, `WebAppFW`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
func (r *WebAppFWResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.WebAppFW

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
