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

var _ resource.Resource = &HTTPSSourceProtocolHandlerResource{}
var _ resource.ResourceWithValidateConfig = &HTTPSSourceProtocolHandlerResource{}
var _ resource.ResourceWithImportState = &HTTPSSourceProtocolHandlerResource{}

func NewHTTPSSourceProtocolHandlerResource() resource.Resource {
	return &HTTPSSourceProtocolHandlerResource{}
}

type HTTPSSourceProtocolHandlerResource struct {
	pData *tfutils.ProviderData
}

func (r *HTTPSSourceProtocolHandlerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_https_source_protocol_handler"
}

func (r *HTTPSSourceProtocolHandlerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("An HTTPS handler receives HTTP requests over TLS and forwards them to the appropriate DataPower service. HTTPS handlers conform to RFC 2616.", "source-https", "").AddActions("quiesce").String,
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
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"local_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the IP address or host alias that the handler listens. The default value indicates that The handler listens on all IPv4 addresses.", "local-address", "").AddDefaultValue("0.0.0.0").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("0.0.0.0"),
			},
			"local_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Port", "port", "").AddIntegerRange(1, 65535).AddDefaultValue("443").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(443),
			},
			"http_version": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the HTTP version for client connections. The default value is HTTP/1.1. For the HTTP/2 protocol, requests and responses are always HTTP/2. When HTTP/2, this setting is ignored.", "http-client-version", "").AddStringEnum("HTTP/1.0", "HTTP/1.1").AddDefaultValue("HTTP/1.1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("HTTP/1.0", "HTTP/1.1"),
				},
				Default: stringdefault.StaticString("HTTP/1.1"),
			},
			"allowed_features": models.GetDmSourceHTTPFeatureTypeResourceSchema("Allowed methods and versions", "allowed-features", "", false),
			"persistent_connections": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to negotiate persistent connections with clients. The HTTP/2 protocol controls persistent connections and reuse. Therefore, this setting is ignored for the HTTP/2 protocol.", "persistent-connections", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"max_persistent_connections_reuse": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum number of times that a client can reuse a persistent connection. When this count is reached, an explicit <tt>HTTP Connection: close</tt> header is sent in the response, and the TCP connection is closed. The default value is 0, which means unlimited reuse.", "max-persistent-reuse", "").AddNotValidWhen(models.HTTPSSourceProtocolHandlerMaxPersistentConnectionsReuseIgnoreVal.String()).String,
				Optional:            true,
			},
			"allow_compression": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to negotiate GZIP compression for client connections. When enabled and the <tt>Accept-Encoding</tt> HTTP header indicates that compressed documents can be processed, the service uses GZIP to compress HTTP transmissions. The <tt>Transfer-Encoding</tt> HTTP header indicates compression.", "compression", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"allow_web_socket_upgrade": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to allow WebSocket upgrade requests from clients. The default value is disabled. This request is to switch the existing connection to use the WebSocket protocol. WebSocket upgrade requests require that The handler allows GET methods.", "websocket-upgrade", "").AddDefaultValue("false").AddNotValidWhen(models.HTTPSSourceProtocolHandlerAllowWebSocketUpgradeIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"web_socket_idle_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum idle time in seconds for client connections. This timer monitors the idle time in the data transfer process. When the specified idle time is exceeded, the connection is torn down. Enter a value in the range 0 - 86400. The default value is 0, which indicates that the timer is disabled.", "websocket-idle-timeout", "").AddIntegerRange(0, 86400).AddRequiredWhen(models.HTTPSSourceProtocolHandlerWebSocketIdleTimeoutCondVal.String()).AddNotValidWhen(models.HTTPSSourceProtocolHandlerWebSocketIdleTimeoutIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 86400),
					validators.ConditionalRequiredInt64(models.HTTPSSourceProtocolHandlerWebSocketIdleTimeoutCondVal, models.HTTPSSourceProtocolHandlerWebSocketIdleTimeoutIgnoreVal, false),
				},
			},
			"max_url_len": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the length in bytes of the longest incoming URL to accept. The length includes any query string or fragment identifier. Enter a value in the range 1 - 128000. The default value is 16384.", "max-url-len", "").AddIntegerRange(1, 128000).AddDefaultValue("16384").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 128000),
				},
				Default: int64default.StaticInt64(16384),
			},
			"max_total_hdr_len": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum aggregate length in bytes of HTTP headers to allow. Enter a value in the range 5 - 128000. The default value is 128000.", "max-total-header-len", "").AddIntegerRange(5, 128000).AddDefaultValue("128000").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(5, 128000),
				},
				Default: int64default.StaticInt64(128000),
			},
			"max_hdr_count": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum number of headers to allow in client requests. The default value is 0, which indicates no limit.", "max-header-count", "").String,
				Optional:            true,
			},
			"max_name_hdr_len": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum length in bytes of a header name to allow in client requests. Each HTTP header is expressed as a name-value pair. This setting sets the maximum length of the name portion of a header. The default value is 0, which indicates no limit.", "max-header-name-len", "").String,
				Optional:            true,
			},
			"max_value_hdr_len": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum length of a header value in bytes to allow in client requests. Each HTTP header is expressed as a name-value pair. This setting sets the maximum length of the value portion of a header. The default value is 0, which indicates no limit.", "max-header-value-len", "").String,
				Optional:            true,
			},
			"max_query_string_len": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum length of the query string in bytes to allow in client requests. The query string is the portion of the URL after the ? character. The default value is 0, which indicates no limit.", "max-querystring-len", "").String,
				Optional:            true,
			},
			"acl": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Access control list", "acl", "access_control_list").String,
				Optional:            true,
			},
			"credential_charset": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the character encoding of the original basic authentication values. Basic authentication credentials are combined and base64 encoded in the authorization header of the request. The contents of the <tt>Authorization</tt> header are transcoded to UTF-8. The default value represents ISO-8859-1 Latin 1.", "credential-charset", "").AddStringEnum("protocol", "ascii", "utf8", "big5", "cp1250", "cp1251", "cp1252", "cp1253", "cp1254", "cp1255", "cp1256", "cp1257", "cp1258", "euc_jp", "euc_kr", "gb18030", "gb2312", "iso2022_jp", "iso2022_kr", "iso8859_1", "iso8859_2", "iso8859_4", "iso8859_5", "iso8859_6", "iso8859_7", "iso8859_8", "iso8859_9", "iso8859_15", "sjis", "tis620", "unicode_le").AddDefaultValue("protocol").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("protocol", "ascii", "utf8", "big5", "cp1250", "cp1251", "cp1252", "cp1253", "cp1254", "cp1255", "cp1256", "cp1257", "cp1258", "euc_jp", "euc_kr", "gb18030", "gb2312", "iso2022_jp", "iso2022_kr", "iso8859_1", "iso8859_2", "iso8859_4", "iso8859_5", "iso8859_6", "iso8859_7", "iso8859_8", "iso8859_9", "iso8859_15", "sjis", "tis620", "unicode_le"),
				},
				Default: stringdefault.StaticString("protocol"),
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS server profile to secure connections from clients. <b>Note:</b> The TLS server profile that secures the HTTP/2 connection must use the TLS 1.2 or later protocol with a cipher that is secure according to RFC 7540.", "ssl-server", "ssl_server_profile").AddRequiredWhen(models.HTTPSSourceProtocolHandlerSSLServerCondVal.String()).AddNotValidWhen(models.HTTPSSourceProtocolHandlerSSLServerIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.HTTPSSourceProtocolHandlerSSLServerCondVal, models.HTTPSSourceProtocolHandlerSSLServerIgnoreVal, false),
				},
			},
			"ssl_sni_server": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS SNI server profile to secure connections from clients. <b>Note:</b> The TLS SNI server profile that secures the HTTP/2 connection must use the TLS 1.2 or later protocol with a cipher that is secure according to RFC 7540.", "ssl-sni-server", "ssl_sni_server_profile").AddRequiredWhen(models.HTTPSSourceProtocolHandlerSSLSNIServerCondVal.String()).AddNotValidWhen(models.HTTPSSourceProtocolHandlerSSLSNIServerIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.HTTPSSourceProtocolHandlerSSLSNIServerCondVal, models.HTTPSSourceProtocolHandlerSSLSNIServerIgnoreVal, false),
				},
			},
			"http2_max_streams": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum number of concurrent streams that the client can have outstanding at the same time. Enter a value in the range 1 - 500. The default value is 100. <p>The limit applies to the number of streams that the client allows the target to create. The greater the number of streams in use, the more resources the client uses. Resources include memory and the network connections to the destination.</p>", "http2-max-streams", "").AddIntegerRange(1, 500).AddDefaultValue("100").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 500),
				},
				Default: int64default.StaticInt64(100),
			},
			"http2_max_frame_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the largest payload frame size that the client can send. Enter a value in the range 16384 - 16777215. The default value is 16384.", "http2-max-frame", "").AddIntegerRange(16384, 16777215).AddDefaultValue("16384").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(16384, 16777215),
				},
				Default: int64default.StaticInt64(16384),
			},
			"http2_stream_header": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable the HTTP/2 stream identifier header in the request or response. When enabled, the HTTP/2 stream identifier is included in the <tt>X-DP-http2-stream</tt> header. With this header, you can correlate the HTTP/2 stream. The default behavior is disabled.", "http2-stream-header", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"chunked_encoding": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable responses to use chunked transfer-encoding. By default, HTTP responses use <tt>Transfer-Encoding: chunked</tt> .", "chunked-encoding", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"header_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum duration in milliseconds to allow for request headers processing. When the value is greater than 0, request header processing must complete before the duration elapses. Enter a value in the range 0 - 3600000, where a value of 0 disables the timer. The default value is 30000.", "header-timeout", "").AddIntegerRange(0, 3600000).AddDefaultValue("30000").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 3600000),
				},
				Default: int64default.StaticInt64(30000),
			},
			"http2_idle_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum idle duration in milliseconds to allow before closing the HTTP/2 connection. Enter a value in the range 0 - 3600000, where a value of 0 disables the timer. The default value is 0.", "http2-idle-timeout", "").AddIntegerRange(0, 3600000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 3600000),
				},
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *HTTPSSourceProtocolHandlerResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *HTTPSSourceProtocolHandlerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.HTTPSSourceProtocolHandler
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `HTTPSSourceProtocolHandler`)
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

func (r *HTTPSSourceProtocolHandlerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.HTTPSSourceProtocolHandler
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

	data.UpdateFromBody(ctx, `HTTPSSourceProtocolHandler`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HTTPSSourceProtocolHandlerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.HTTPSSourceProtocolHandler
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `HTTPSSourceProtocolHandler`))
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

func (r *HTTPSSourceProtocolHandlerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.HTTPSSourceProtocolHandler
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

func (r *HTTPSSourceProtocolHandlerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

	var data models.HTTPSSourceProtocolHandler
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

	data.FromBody(ctx, `HTTPSSourceProtocolHandler`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HTTPSSourceProtocolHandlerResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.HTTPSSourceProtocolHandler

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
