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
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &WebAppFWResource{}

func NewWebAppFWResource() resource.Resource {
	return &WebAppFWResource{}
}

type WebAppFWResource struct {
	client *client.DatapowerClient
}

func (r *WebAppFWResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_webappfw"
}

func (r *WebAppFWResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Web Application Firewall", "web-application-firewall", "").String,

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
			"front_side": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Source Addresses", "listen-on", "").String,
				NestedObject:        models.DmFrontSideResourceSchema,
				Optional:            true,
			},
			"remote_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Remote Host", "remote-address", "").String,
				Required:            true,
			},
			"remote_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Remote Port", "remote-port", "").AddIntegerRange(1, 65535).AddDefaultValue("80").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(80),
			},
			"style_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Security Policy", "security-policy", "appsecuritypolicy").String,
				Required:            true,
			},
			"xml_manager": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML Manager", "xml-manager", "xmlmanager").AddDefaultValue("default").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default"),
			},
			"error_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Error Policy", "error-policy", "webapperrorhandlingpolicy").String,
				Optional:            true,
			},
			"uri_normalization": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Normalize URI", "uri-normalization", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"rewrite_errors": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Rewrite Error Messages", "rewrite-errors", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"delay_errors": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Delay Error Messages", "delay-errors", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"delay_errors_duration": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Duration to Delay Error Messages", "delay-errors-duration", "").AddIntegerRange(250, 300000).AddDefaultValue("1000").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(250, 300000),
				},
				Default: int64default.StaticInt64(1000),
			},
			"stream_output_to_back": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Stream Output to Back", "stream-output-to-back", "").AddStringEnum("buffer-until-verification", "stream-until-infraction").AddDefaultValue("buffer-until-verification").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("buffer-until-verification", "stream-until-infraction"),
				},
				Default: stringdefault.StaticString("buffer-until-verification"),
			},
			"stream_output_to_front": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Stream Output to Front", "stream-output-to-front", "").AddStringEnum("buffer-until-verification", "stream-until-infraction").AddDefaultValue("buffer-until-verification").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("buffer-until-verification", "stream-until-infraction"),
				},
				Default: stringdefault.StaticString("buffer-until-verification"),
			},
			"front_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Front Side Timeout", "front-timeout", "").AddIntegerRange(1, 86400).AddDefaultValue("120").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(120),
			},
			"back_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Back Side Timeout", "back-timeout", "").AddIntegerRange(1, 86400).AddDefaultValue("120").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(120),
			},
			"front_persistent_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Front Persistent Timeout", "front-persistent-timeout", "").AddIntegerRange(0, 86400).AddDefaultValue("180").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 86400),
				},
				Default: int64default.StaticInt64(180),
			},
			"allow_cache_control_header": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow Cache-Control Header", "allow-cache-control", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"back_persistent_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Back Persistent Timeout", "back-persistent-timeout", "").AddIntegerRange(0, 86400).AddDefaultValue("180").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 86400),
				},
				Default: int64default.StaticInt64(180),
			},
			"front_http_version": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP Response Version", "http-front-version", "").AddStringEnum("HTTP/1.0", "HTTP/1.1").AddDefaultValue("HTTP/1.1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("HTTP/1.0", "HTTP/1.1"),
				},
				Default: stringdefault.StaticString("HTTP/1.1"),
			},
			"back_http_version": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP Version to Server", "http-back-version", "").AddStringEnum("HTTP/1.0", "HTTP/1.1").AddDefaultValue("HTTP/1.1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("HTTP/1.0", "HTTP/1.1"),
				},
				Default: stringdefault.StaticString("HTTP/1.1"),
			},
			"request_side_security": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Request Security", "request-security", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"response_side_security": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Response Security", "response-security", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"do_chunked_upload": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow Chunked Uploads", "chunked-uploads", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"follow_redirects": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Follow Redirects", "follow-redirects", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"http_client_ip_label": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP Client IP Label", "http-client-ip-label", "").AddDefaultValue("X-Client-IP").String,
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
			"url_rewrite_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Header and URL Rewrite Policy", "urlrewrite-policy", "urlrewritepolicy").String,
				Optional:            true,
			},
			"do_host_rewriting": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Rewrite Host Names When Gatewaying", "host-rewriting", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
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
		},
	}
}

func (r *WebAppFWResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *WebAppFWResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.WebAppFW

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `WebAppFW`)
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

func (r *WebAppFWResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.WebAppFW

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
		data.FromBody(ctx, `WebAppFW`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `WebAppFW`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WebAppFWResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.WebAppFW

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `WebAppFW`))
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

func (r *WebAppFWResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.WebAppFW

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
