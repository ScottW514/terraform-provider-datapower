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

var _ resource.Resource = &AssemblyActionInvokeResource{}
var _ resource.ResourceWithImportState = &AssemblyActionInvokeResource{}

func NewAssemblyActionInvokeResource() resource.Resource {
	return &AssemblyActionInvokeResource{}
}

type AssemblyActionInvokeResource struct {
	pData *tfutils.ProviderData
}

func (r *AssemblyActionInvokeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_assembly_action_invoke"
}

func (r *AssemblyActionInvokeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("The invoke assembly action call a service from your assembly.", "assembly-invoke", "").String,
		Attributes: map[string]schema.Attribute{
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
			"url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL of the target. You can reference a custom API property that resolves as the value. To reference an API property, use the <tt>$(api.properties. <i>property_name</i> )</tt> format, where <tt><i>property_name</i></tt> is the name of the property to reference. You can use the short <tt>$( <i>property_name</i> )</tt> format when the assembly action does not have a property with the same name.", "url", "").String,
				Required:            true,
			},
			"ssl_client": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "").String,
				Optional:            true,
			},
			"timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration in seconds to wait for a reply from the target. The default value is 60.", "timeout", "").AddDefaultValue("60").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(60),
			},
			"user_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Username", "username", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile("^[^ ]+$"), "Must match :"+"^[^ ]+$"),
				},
			},
			"password": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Password", "password", "password_alias").String,
				Optional:            true,
			},
			"method": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP method", "method", "").AddStringEnum("Keep", "GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS", "TRACE").AddDefaultValue("Keep").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Keep", "GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS", "TRACE"),
				},
				Default: stringdefault.StaticString("Keep"),
			},
			"backend_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Backend type", "backend-type", "").AddStringEnum("detect", "xml", "json", "binary", "graphql").AddDefaultValue("detect").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("detect", "xml", "json", "binary", "graphql"),
				},
				Default: stringdefault.StaticString("detect"),
			},
			"graphql_send_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the type of payload to send for GraphQL POST requests. When GraphQL or JSON, this setting overrides the message type of the payload.", "graphql-send-type", "").AddStringEnum("detect", "graphql", "json").AddDefaultValue("detect").AddRequiredWhen(models.AssemblyActionInvokeGraphQLSendTypeCondVal.String()).AddNotValidWhen(models.AssemblyActionInvokeGraphQLSendTypeIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("detect", "graphql", "json"),
					validators.ConditionalRequiredString(models.AssemblyActionInvokeGraphQLSendTypeCondVal, models.AssemblyActionInvokeGraphQLSendTypeIgnoreVal, true),
				},
				Default: stringdefault.StaticString("detect"),
			},
			"compression": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable compression", "compression", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"cache_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how to cache documents.", "cache-type", "").AddStringEnum("Protocol", "NoCache", "TimeToLive").AddDefaultValue("Protocol").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Protocol", "NoCache", "TimeToLive"),
				},
				Default: stringdefault.StaticString("Protocol"),
			},
			"time_to_live": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the validity period in seconds for documents in the cache. The default value is 900.", "ttl", "").AddDefaultValue("900").AddRequiredWhen(models.AssemblyActionInvokeTimeToLiveCondVal.String()).AddNotValidWhen(models.AssemblyActionInvokeTimeToLiveIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					validators.ConditionalRequiredInt64(models.AssemblyActionInvokeTimeToLiveCondVal, models.AssemblyActionInvokeTimeToLiveIgnoreVal, true),
				},
				Default: int64default.StaticInt64(900),
			},
			"cache_unsafe_response": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to cache responses to POST and PUT requests when the cache policy type is set to time to live. The response to these requests is the result of an action on the server that might change its resource state. You might want to cache responses to these requests when you know that the action (for example: HTTP POST) will not change the server state.", "cache-unsafe-response", "").AddDefaultValue("false").AddNotValidWhen(models.AssemblyActionInvokeCacheUnsafeResponseIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"cache_key": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the string for the cache key. If omitted, the entire URL is used as the key.", "cache-key", "").String,
				Optional:            true,
			},
			"follow_redirects": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to follow redirects. Some protocols generate redirects. When enabled, the action attempts to resolve redirects transparently.", "follow-redirects", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_version": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the HTTP version for server-side connections. The default value is HTTP/1.1.", "http-version", "").AddStringEnum("HTTP/1.0", "HTTP/1.1", "HTTP/2").AddDefaultValue("HTTP/1.1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("HTTP/1.0", "HTTP/1.1", "HTTP/2"),
				},
				Default: stringdefault.StaticString("HTTP/1.1"),
			},
			"http2_required": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether an HTTP/2 connection is required when connecting to the server. Only applicable when the HTTP version to the server is set to HTTP/2 and the connection uses TLS. The default value is off.", "http2-required", "").AddDefaultValue("false").AddNotValidWhen(models.AssemblyActionInvokeHTTP2RequiredIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"do_chunked_upload": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable uploading of HTTP/1.1 chunked-encoded documents. For HTTP/1.1, the document body can be delimited by either <tt>Content-Length</tt> or chunked encoding. While all servers understand <tt>Content-Length</tt> , many servers fail to understand chunked encoding. For this reason, <tt>Content-Length</tt> is the standard method. However, the use of <tt>Content-Length</tt> can interfere with streaming. To stream full documents to an RFC 2616 compatible server, enable this property. Unlike other HTTP/1.1 features, you must know that the target server is RFC 2616 compatible.", "chunked-uploads", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"persistent_connection": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Persistent connection", "persistent-connection", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"stop_on_error": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Stop on error", "stop-on-error", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"error_types": models.GetDmInvokeErrorTypeResourceSchema("Error types", "error-types", "", false),
			"output": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the variable to store results. By default, results are stored in the <tt>message.body</tt> , <tt>message.headers</tt> , <tt>message.statuscode</tt> variables.", "output", "").AddDefaultValue("message").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("message"),
			},
			"decode_request_params": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to decode the request parameters in the target URL. When enabled, request parameters are decoded. By default, request parameters are not decoded.", "decode-request-params", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"encode_plus_char": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to encode + characters in query strings. When enabled, + characters are encoded to <tt>%2F</tt> . By default, + characters are not encoded.", "encode-plus-char", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"keep_payload": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to include the payload for DELETE requests. When enabled, DELETE requests include the payload. By default, DELETE requests do not include the payload.", "keep-payload", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"inject_user_agent_header": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to inject the default <tt>User-Agent</tt> header. When the <tt>User-Agent</tt> header is not in the request, inject this header to the request.", "inject-user-agent", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"inject_proxy_headers": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to inject proxy-related headers. When the <tt>X-Forwarded-For</tt> , <tt>X-Forwarded-Host</tt> , and <tt>X-Forwarded-Port</tt> headers are not found in the request, inject theses headers to the request.", "inject-proxy-headers", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"header_control_list": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the control list that uses headers to accept or reject requests. By default, accepts all requests with headers.", "header-control-list", "control_list").AddDefaultValue("default-accept-all").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default-accept-all"),
			},
			"parameter_control_list": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the control list that uses URL parameters to accept or reject requests. By default, rejects all requests with URL parameters.", "parameter-control-list", "control_list").AddDefaultValue("default-reject-all").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default-reject-all"),
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"title": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Title", "title", "").String,
				Optional:            true,
			},
			"correlation_path": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the path that correlates the API action to a specific part of the API specification. The correlation path specifies the part of the API definition that correlates with the API action. This path is exposed in the debug data by the API gateway for use by debugging tools. For example, for an API configuration that is retrieved from API Connect and specified in an OpenAPI document with IBM extensions, this path is the JSON path to the assembly policy in the IBM extensions section of the document. The path can be expressed in any form that the debugging tool can correlate to the API definition.", "correlation-path", "").String,
				Optional:            true,
			},
			"action_debug": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to enable the GatewayScript debugger to troubleshoot the following GatewayScript files or script.</p><ul><li>Troubleshoot a GatewayScript file that is called from the GatewayScript assembly action.</li><li>Troubleshoot a GatewayScript file that is called from an XSLT assembly action that uses the <tt>gatewayscript()</tt> extension function.</li><li>Troubleshoot a GatewayScript script that is called through the <tt>value</tt> or <tt>default</tt> property in the JSON file from the map assembly action.</li></ul><p>To debug a file or script, the following conditions must be met.</p><ul><li>The file contains one or more <tt>debugger;</tt> statements at the points in your script where you want to start debugging.</li><li>The GatewayScript debugger is enabled.</li></ul><p>You run the <tt>debug-action</tt> command.</p>", "debug", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *AssemblyActionInvokeResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *AssemblyActionInvokeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AssemblyActionInvoke
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `AssemblyActionInvoke`)
	_, err := r.pData.Client.Post(data.GetPath(), body)
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), body)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Resource already exists. Failed to update resource, got error: %s", err))
				return
			}
			resp.Diagnostics.AddWarning("Warning", "Resource already exists. Existing resource was updated.")
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create resource, got error: %s", err))
			return
		}
	}
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AssemblyActionInvokeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AssemblyActionInvoke
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

	data.UpdateFromBody(ctx, `AssemblyActionInvoke`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AssemblyActionInvokeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AssemblyActionInvoke
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `AssemblyActionInvoke`))
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

func (r *AssemblyActionInvokeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AssemblyActionInvoke
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			resp.Diagnostics.AddWarning("Resource Conflict", fmt.Sprintf("Resource is no longer tracked by Terraform, but may need to be manually deleted on DataPower host. Got error: %s", err))
		} else if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete resource, got error: %s", err))
			return
		}
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *AssemblyActionInvokeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

	var data models.AssemblyActionInvoke
	data.AppDomain = types.StringValue(appDomain)
	data.Id = types.StringValue(id)
	res, err := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("Resource was not found, got error: %s", err))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		}
		return
	}

	data.FromBody(ctx, `AssemblyActionInvoke`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AssemblyActionInvokeResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.AssemblyActionInvoke

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
