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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &EBMS2SourceProtocolHandlerResource{}
var _ resource.ResourceWithValidateConfig = &EBMS2SourceProtocolHandlerResource{}

func NewEBMS2SourceProtocolHandlerResource() resource.Resource {
	return &EBMS2SourceProtocolHandlerResource{}
}

type EBMS2SourceProtocolHandlerResource struct {
	pData *tfutils.ProviderData
}

func (r *EBMS2SourceProtocolHandlerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ebms2_source_protocol_handler"
}

func (r *EBMS2SourceProtocolHandlerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("The configuration for an ebMS2 handler.", "source-ebms2", "").AddActions("quiesce").String,
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
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"local_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the local IP address that the service listens. The default value is 0.0.0.0, which indicates that the service is active on all addresses. You can use a local host alias to help ease migration.", "local-address", "").AddDefaultValue("0.0.0.0").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("0.0.0.0"),
			},
			"local_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the port that this service monitors. Enter a value in the range 1 - 65535. The default value is 80.", "port", "").AddIntegerRange(1, 65535).AddDefaultValue("80").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(80),
			},
			"http_version": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the HTTP version for client connection. The default value is HTTP/1.1.", "http-client-version", "").AddStringEnum("HTTP/1.0", "HTTP/1.1").AddDefaultValue("HTTP/1.1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("HTTP/1.0", "HTTP/1.1"),
				},
				Default: stringdefault.StaticString("HTTP/1.1"),
			},
			"persistent_connections": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable persistent connections with clients. By default, persistent connections are enabled.", "persistent-connections", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"allow_compression": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable GZIP compression with clients. By default, compression is not enabled.", "compression", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"max_url_len": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum URL length in bytes to accept from clients. Enter a value in the range 1 - 128000. The default value is 16384. The length includes any query string or fragment identifier.", "max-url-len", "").AddIntegerRange(1, 128000).AddDefaultValue("16384").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 128000),
				},
				Default: int64default.StaticInt64(16384),
			},
			"max_total_hdr_len": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum aggregate size of headers in bytes in client requests. Enter a value in the range 5 - 128000. The default value is 128000.", "max-total-header-len", "").AddIntegerRange(5, 128000).AddDefaultValue("128000").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(5, 128000),
				},
				Default: int64default.StaticInt64(128000),
			},
			"max_hdr_count": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum number of headers to allow in client requests. The default value is 0, which means unlimited.", "max-header-count", "").AddIntegerRange(0, 65535).AddDefaultValue("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 65535),
				},
				Default: int64default.StaticInt64(0),
			},
			"max_name_hdr_len": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum length to accept for any header name. A header is expressed as a name-value pair. The default value is 0, which means unlimited.", "max-header-name-len", "").AddIntegerRange(0, 65535).AddDefaultValue("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 65535),
				},
				Default: int64default.StaticInt64(0),
			},
			"max_value_hdr_len": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum length to accept for any header value. A header is expressed as a name-value pair. The default value is 0, which means unlimited.", "max-header-value-len", "").AddIntegerRange(0, 65535).AddDefaultValue("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 65535),
				},
				Default: int64default.StaticInt64(0),
			},
			"max_query_string_len": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum length of the query string. The query string is the portion of the URL after the <tt>?</tt> character. The default value is 0, which indicates no limit.", "max-querystring-len", "").AddIntegerRange(0, 65535).AddDefaultValue("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 65535),
				},
				Default: int64default.StaticInt64(0),
			},
			"acl": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Access control list", "acl", "access_control_list").String,
				Optional:            true,
			},
			"aaa_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("AAA policy", "aaa-policy", "aaa_policy").String,
				Optional:            true,
			},
			"credential_charset": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the character encoding of the original basic authentication values. The default value is protocol, which is ISO-8859-1 (Latin 1). <p>Basic authentication credentials are combined and base64 encoded in the HTTP <tt>Authorization</tt> request header. The contents of the <tt>Authorization</tt> header is transcoded to UTF-8.</p>", "credential-charset", "").AddStringEnum("protocol", "ascii", "utf8", "big5", "cp1250", "cp1251", "cp1252", "cp1253", "cp1254", "cp1255", "cp1256", "cp1257", "cp1258", "euc_jp", "euc_kr", "gb18030", "gb2312", "iso2022_jp", "iso2022_kr", "iso8859_1", "iso8859_2", "iso8859_4", "iso8859_5", "iso8859_6", "iso8859_7", "iso8859_8", "iso8859_9", "iso8859_15", "sjis", "tis620", "unicode_le").AddDefaultValue("protocol").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("protocol", "ascii", "utf8", "big5", "cp1250", "cp1251", "cp1252", "cp1253", "cp1254", "cp1255", "cp1256", "cp1257", "cp1258", "euc_jp", "euc_kr", "gb18030", "gb2312", "iso2022_jp", "iso2022_kr", "iso8859_1", "iso8859_2", "iso8859_4", "iso8859_5", "iso8859_6", "iso8859_7", "iso8859_8", "iso8859_9", "iso8859_15", "sjis", "tis620", "unicode_le"),
				},
				Default: stringdefault.StaticString("protocol"),
			},
			"ssl_server_config_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the type of TLS profile to secure client connections.", "ssl-config-type", "").AddStringEnum("server", "sni").AddDefaultValue("server").String,
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
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *EBMS2SourceProtocolHandlerResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *EBMS2SourceProtocolHandlerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.EBMS2SourceProtocolHandler
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `EBMS2SourceProtocolHandler`)
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

func (r *EBMS2SourceProtocolHandlerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.EBMS2SourceProtocolHandler
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
		data.FromBody(ctx, `EBMS2SourceProtocolHandler`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `EBMS2SourceProtocolHandler`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EBMS2SourceProtocolHandlerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.EBMS2SourceProtocolHandler
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `EBMS2SourceProtocolHandler`))
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

func (r *EBMS2SourceProtocolHandlerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.EBMS2SourceProtocolHandler
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

func (r *EBMS2SourceProtocolHandlerResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.EBMS2SourceProtocolHandler

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
