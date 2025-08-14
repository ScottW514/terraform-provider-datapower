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

type HTTPSourceProtocolHandlerList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &HTTPSourceProtocolHandlerDataSource{}
	_ datasource.DataSourceWithConfigure = &HTTPSourceProtocolHandlerDataSource{}
)

func NewHTTPSourceProtocolHandlerDataSource() datasource.DataSource {
	return &HTTPSourceProtocolHandlerDataSource{}
}

type HTTPSourceProtocolHandlerDataSource struct {
	pData *tfutils.ProviderData
}

func (d *HTTPSourceProtocolHandlerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_httpsourceprotocolhandler"
}

func (d *HTTPSourceProtocolHandlerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "HTTP handler",
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
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"local_address": schema.StringAttribute{
							MarkdownDescription: "Local IP address",
							Computed:            true,
						},
						"local_port": schema.Int64Attribute{
							MarkdownDescription: "Port",
							Computed:            true,
						},
						"http_version": schema.StringAttribute{
							MarkdownDescription: "HTTP version to client",
							Computed:            true,
						},
						"allowed_features": models.GetDmSourceHTTPFeatureTypeDataSourceSchema("Allowed methods and versions", "allowed-features", ""),
						"persistent_connections": schema.BoolAttribute{
							MarkdownDescription: "Negotiate persistent connections",
							Computed:            true,
						},
						"max_persistent_connections_reuse": schema.Int64Attribute{
							MarkdownDescription: "Maximum persistent reuse",
							Computed:            true,
						},
						"allow_compression": schema.BoolAttribute{
							MarkdownDescription: "Enable compression",
							Computed:            true,
						},
						"allow_web_socket_upgrade": schema.BoolAttribute{
							MarkdownDescription: "Allow WebSocket upgrade",
							Computed:            true,
						},
						"web_socket_idle_timeout": schema.Int64Attribute{
							MarkdownDescription: "WebSocket idle timeout",
							Computed:            true,
						},
						"max_url_len": schema.Int64Attribute{
							MarkdownDescription: "Maximum URL length",
							Computed:            true,
						},
						"max_total_hdr_len": schema.Int64Attribute{
							MarkdownDescription: "Maximum total header length",
							Computed:            true,
						},
						"max_hdr_count": schema.Int64Attribute{
							MarkdownDescription: "Maximum request headers",
							Computed:            true,
						},
						"max_name_hdr_len": schema.Int64Attribute{
							MarkdownDescription: "Maximum header name length",
							Computed:            true,
						},
						"max_value_hdr_len": schema.Int64Attribute{
							MarkdownDescription: "Maximum header value length",
							Computed:            true,
						},
						"max_query_string_len": schema.Int64Attribute{
							MarkdownDescription: "Maximum query string length",
							Computed:            true,
						},
						"acl": schema.StringAttribute{
							MarkdownDescription: "Access control list",
							Computed:            true,
						},
						"credential_charset": schema.StringAttribute{
							MarkdownDescription: "Credential character set",
							Computed:            true,
						},
						"http2_max_streams": schema.Int64Attribute{
							MarkdownDescription: "HTTP/2 maximum streams",
							Computed:            true,
						},
						"http2_max_frame_size": schema.Int64Attribute{
							MarkdownDescription: "HTTP/2 max frame size",
							Computed:            true,
						},
						"http2_stream_header": schema.BoolAttribute{
							MarkdownDescription: "Enable HTTP/2 stream header",
							Computed:            true,
						},
						"chunked_encoding": schema.BoolAttribute{
							MarkdownDescription: "Enable chunked encoding responses",
							Computed:            true,
						},
						"header_timeout": schema.Int64Attribute{
							MarkdownDescription: "Request headers processing timeout",
							Computed:            true,
						},
						"http2_idle_timeout": schema.Int64Attribute{
							MarkdownDescription: "HTTP/2 idle timeout",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *HTTPSourceProtocolHandlerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *HTTPSourceProtocolHandlerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data HTTPSourceProtocolHandlerList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.HTTPSourceProtocolHandler{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.HTTPSourceProtocolHandler{}
	if value := res.Get(`HTTPSourceProtocolHandler`); value.Exists() {
		for _, v := range value.Array() {
			item := models.HTTPSourceProtocolHandler{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.HTTPSourceProtocolHandlerObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.HTTPSourceProtocolHandlerObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
