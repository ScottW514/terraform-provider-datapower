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

package provider

import (
	"context"
	"fmt"
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
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &XSLProxyServiceResource{}

func NewXSLProxyServiceResource() resource.Resource {
	return &XSLProxyServiceResource{}
}

type XSLProxyServiceResource struct {
	client *client.DatapowerClient
}

func (r *XSLProxyServiceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_xslproxyservice"
}

func (r *XSLProxyServiceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("XSL Proxy", "xslproxy", "").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name of the object. Must be unique among object types in application domain.", "", "").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"app_domain": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The name of the application domain the object belongs to", "", "").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					ImmutableAfterSet(),
				},
			},
			"type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Type", "type", "").AddStringEnum("static-backend", "loopback-proxy", "strict-proxy", "dynamic-backend").AddDefaultValue("static-backend").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("static-backend", "loopback-proxy", "strict-proxy", "dynamic-backend"),
				},
				Default: stringdefault.StaticString("static-backend"),
			},
			"xml_manager": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML Manager", "xml-manager", "xmlmanager").AddDefaultValue("default").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default"),
			},
			"url_rewrite_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("URL Rewrite Policy", "urlrewrite-policy", "urlrewritepolicy").String,
				Optional:            true,
			},
			"style_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Processing Policy", "stylesheet-policy", "stylepolicy").AddDefaultValue("default").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default"),
			},
			"credential_charset": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Credential Character Set", "credential-charset", "").AddStringEnum("protocol", "ascii", "utf8", "big5", "cp1250", "cp1251", "cp1252", "cp1253", "cp1254", "cp1255", "cp1256", "cp1257", "cp1258", "euc_jp", "euc_kr", "gb18030", "gb2312", "iso2022_jp", "iso2022_kr", "iso8859_1", "iso8859_2", "iso8859_4", "iso8859_5", "iso8859_6", "iso8859_7", "iso8859_8", "iso8859_9", "iso8859_15", "sjis", "tis620", "unicode_le").AddDefaultValue("protocol").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("protocol", "ascii", "utf8", "big5", "cp1250", "cp1251", "cp1252", "cp1253", "cp1254", "cp1255", "cp1256", "cp1257", "cp1258", "euc_jp", "euc_kr", "gb18030", "gb2312", "iso2022_jp", "iso2022_kr", "iso8859_1", "iso8859_2", "iso8859_4", "iso8859_5", "iso8859_6", "iso8859_7", "iso8859_8", "iso8859_9", "iso8859_15", "sjis", "tis620", "unicode_le"),
				},
				Default: stringdefault.StaticString("protocol"),
			},
			"ssl_config_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS type", "ssl-config-type", "").AddStringEnum("server", "sni").AddDefaultValue("server").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("server", "sni"),
				},
				Default: stringdefault.StaticString("server"),
			},
			"ssl_server": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS server profile", "ssl-server", "sslserverprofile").String,
				Optional:            true,
			},
			"sslsni_server": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS SNI server profile", "ssl-sni-server", "sslsniserverprofile").String,
				Optional:            true,
			},
			"ssl_client": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "sslclientprofile").String,
				Optional:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"priority": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Service Priority", "priority", "").AddStringEnum("unknown", "high-min", "high", "high-max", "normal-min", "normal", "normal-max", "low-min", "low", "low-max").AddDefaultValue("normal").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("unknown", "high-min", "high", "high-max", "normal-min", "normal", "normal-max", "low-min", "low", "low-max"),
				},
				Default: stringdefault.StaticString("normal"),
			},
			"local_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Port Number", "port", "").AddIntegerRange(1, 65535).String,
				Required:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 65535),
				},
			},
			"remote_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Remote Host", "remote-ip-address", "").String,
				Required:            true,
			},
			"remote_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Remote Port", "remote-port", "").AddIntegerRange(1, 65535).String,
				Required:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 65535),
				},
			},
			"acl": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Access Control List", "acl", "accesscontrollist").String,
				Optional:            true,
			},
			"http_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP Timeout", "http-timeout", "").AddIntegerRange(1, 86400).AddDefaultValue("120").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(120),
			},
			"http_persist_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP Persistent Timeout", "persistent-timeout", "").AddIntegerRange(0, 7200).AddDefaultValue("180").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 7200),
				},
				Default: int64default.StaticInt64(180),
			},
			"do_host_rewrite": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Host Rewrite", "host-rewriting", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"suppress_http_warnings": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP Warning Suppression", "silence-warning", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_compression": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP Compression", "compression", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_include_response_type_encoding": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP Include charset in response-type", "include-response-type-encoding", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"always_show_errors": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Always provide full errors", "always-show-errors", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"disallow_get": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Disallow GET (and HEAD)", "disallow-get", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"disallow_empty_response": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Don't allow empty response bodies", "disallow-empty-reply", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_persistent_connections": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP Persistent Connections", "persistent-connections", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"http_client_ip_label": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP Client IP Label", "client-address", "").AddDefaultValue("X-Client-IP").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("X-Client-IP"),
			},
			"http_log_cor_id_label": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP Global Transaction ID Label", "http-global-tranID-label", "").AddDefaultValue("X-Global-Transaction-ID").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("X-Global-Transaction-ID"),
			},
			"http_proxy_host": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Proxy Host", "httpproxy-address", "").String,
				Optional:            true,
			},
			"http_proxy_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Proxy Port", "httpproxy-port", "").AddIntegerRange(1, 65535).AddDefaultValue("800").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(800),
			},
			"http_version": models.GetDmHTTPClientServerVersionResourceSchema("HTTP Version", "version", "", false),
			"do_chunked_upload": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow Chunked Uploads", "chunked-uploads", "").AddDefaultValue("false").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP Header Suppression", "suppress", "").String,
				NestedObject:        models.DmHeaderSuppressionResourceSchema,
				Optional:            true,
			},
			"stylesheet_parameters": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Stylesheet Parameter", "parameter", "").String,
				NestedObject:        models.DmStylesheetParameterResourceSchema,
				Optional:            true,
			},
			"default_param_namespace": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Default parameter namespace", "default-param-namespace", "").AddDefaultValue("http://www.datapower.com/param/config").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("http://www.datapower.com/param/config"),
			},
			"query_param_namespace": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Query parameter namespace", "query-param-namespace", "").AddDefaultValue("http://www.datapower.com/param/query").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("http://www.datapower.com/param/query"),
			},
			"force_policy_exec": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Process Messages Whose Body Is Empty", "force-policy-exec", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"count_monitors": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Count Monitors", "monitor-count", "countmonitor").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"duration_monitors": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Duration Monitors", "monitor-duration", "durationmonitor").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"monitor_processing_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Monitors evaluation method", "monitor-processing-policy", "").AddStringEnum("terminate-at-first-throttle", "terminate-at-first-match").AddDefaultValue("terminate-at-first-throttle").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("terminate-at-first-throttle", "terminate-at-first-match"),
				},
				Default: stringdefault.StaticString("terminate-at-first-throttle"),
			},
			"debug_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Probe setting", "debug-mode", "").AddStringEnum("on", "off", "unbounded").AddDefaultValue("off").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("on", "off", "unbounded"),
				},
				Default: stringdefault.StaticString("off"),
			},
			"debug_history": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Transaction History", "debug-history", "").AddIntegerRange(10, 250).AddDefaultValue("25").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(10, 250),
				},
				Default: int64default.StaticInt64(25),
			},
			"debug_trigger": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Probe Triggers", "debug-trigger", "").String,
				NestedObject:        models.DmMSDebugTriggerTypeResourceSchema,
				Optional:            true,
			},
			"local_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Local address", "ip-address", "").AddDefaultValue("0.0.0.0").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("0.0.0.0"),
			},
		},
	}
}

func (r *XSLProxyServiceResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *XSLProxyServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.XSLProxyService

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `XSLProxyService`)
	_, err := r.client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "POST", err))
		return
	}

	_, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s", "POST", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *XSLProxyServiceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.XSLProxyService

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Get(data.GetPath() + "/" + data.Id.ValueString())
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
	var data models.XSLProxyService

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `XSLProxyService`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}
	_, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s", "POST", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *XSLProxyServiceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.XSLProxyService

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}

	resp.State.RemoveResource(ctx)
}
