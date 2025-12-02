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
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type EBMS2SourceProtocolHandlerList struct {
	ProviderTarget types.String `tfsdk:"provider_target"`
	Id             types.String `tfsdk:"id"`
	AppDomain      types.String `tfsdk:"app_domain"`
	Result         types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &EBMS2SourceProtocolHandlerDataSource{}
	_ datasource.DataSourceWithConfigure = &EBMS2SourceProtocolHandlerDataSource{}
)

func NewEBMS2SourceProtocolHandlerDataSource() datasource.DataSource {
	return &EBMS2SourceProtocolHandlerDataSource{}
}

type EBMS2SourceProtocolHandlerDataSource struct {
	pData *tfutils.ProviderData
}

func (d *EBMS2SourceProtocolHandlerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ebms2_source_protocol_handler"
}

func (d *EBMS2SourceProtocolHandlerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The configuration for an ebMS2 handler.",
		Attributes: map[string]schema.Attribute{
			"provider_target": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Target host to retrieve this data from. If not set, provider will use the top level settings.", "", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
			},
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
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"local_address": schema.StringAttribute{
							MarkdownDescription: "Specify the local IP address that the service listens. The default value is 0.0.0.0, which indicates that the service is active on all addresses. You can use a local host alias to help ease migration.",
							Computed:            true,
						},
						"local_port": schema.Int64Attribute{
							MarkdownDescription: "Specifies the port that this service monitors. Enter a value in the range 1 - 65535. The default value is 80.",
							Computed:            true,
						},
						"http_version": schema.StringAttribute{
							MarkdownDescription: "Specify the HTTP version for client connection. The default value is HTTP/1.1.",
							Computed:            true,
						},
						"persistent_connections": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enable persistent connections with clients. By default, persistent connections are enabled.",
							Computed:            true,
						},
						"allow_compression": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enable GZIP compression with clients. By default, compression is not enabled.",
							Computed:            true,
						},
						"max_url_len": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum URL length in bytes to accept from clients. Enter a value in the range 1 - 128000. The default value is 16384. The length includes any query string or fragment identifier.",
							Computed:            true,
						},
						"max_total_hdr_len": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum aggregate size of headers in bytes in client requests. Enter a value in the range 5 - 128000. The default value is 128000.",
							Computed:            true,
						},
						"max_hdr_count": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum number of headers to allow in client requests. The default value is 0, which means unlimited.",
							Computed:            true,
						},
						"max_name_hdr_len": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum length to accept for any header name. A header is expressed as a name-value pair. The default value is 0, which means unlimited.",
							Computed:            true,
						},
						"max_value_hdr_len": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum length to accept for any header value. A header is expressed as a name-value pair. The default value is 0, which means unlimited.",
							Computed:            true,
						},
						"max_query_string_len": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum length of the query string. The query string is the portion of the URL after the <tt>?</tt> character. The default value is 0, which indicates no limit.",
							Computed:            true,
						},
						"acl": schema.StringAttribute{
							MarkdownDescription: "Access control list",
							Computed:            true,
						},
						"aaa_policy": schema.StringAttribute{
							MarkdownDescription: "AAA policy",
							Computed:            true,
						},
						"credential_charset": schema.StringAttribute{
							MarkdownDescription: "Specify the character encoding of the original basic authentication values. The default value is protocol, which is ISO-8859-1 (Latin 1). <p>Basic authentication credentials are combined and base64 encoded in the HTTP <tt>Authorization</tt> request header. The contents of the <tt>Authorization</tt> header is transcoded to UTF-8.</p>",
							Computed:            true,
						},
						"ssl_server_config_type": schema.StringAttribute{
							MarkdownDescription: "Specify the type of TLS profile to secure client connections.",
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
						"dependency_actions": actions.ActionsSchema,
						"provider_target": schema.StringAttribute{
							MarkdownDescription: tfutils.NewAttributeDescription("Target host to retrieve this data from. If not set, provider will use the top level settings.", "", "").String,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *EBMS2SourceProtocolHandlerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *EBMS2SourceProtocolHandlerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data EBMS2SourceProtocolHandlerList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !d.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), d.pData.Client.GetTargetNames()))
		return
	}
	o := models.EBMS2SourceProtocolHandler{
		AppDomain: data.AppDomain,
	}

	path := o.GetPath()
	if !data.Id.IsNull() {
		path = path + "/" + data.Id.ValueString()
	}

	res, err := d.pData.Client.Get(path, data.ProviderTarget)
	resFound := true
	if err != nil {
		if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(d.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		} else {
			resFound = false
		}
	}
	l := []models.EBMS2SourceProtocolHandler{}
	if resFound {
		if value := res.Get(`EBMS2SourceProtocolHandler`); value.Exists() {
			for _, v := range value.Array() {
				item := models.EBMS2SourceProtocolHandler{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.EBMS2SourceProtocolHandlerObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.EBMS2SourceProtocolHandlerObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
