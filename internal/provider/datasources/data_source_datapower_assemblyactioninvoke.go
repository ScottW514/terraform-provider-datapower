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
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
)

type AssemblyActionInvokeList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &AssemblyActionInvokeDataSource{}
	_ datasource.DataSourceWithConfigure = &AssemblyActionInvokeDataSource{}
)

func NewAssemblyActionInvokeDataSource() datasource.DataSource {
	return &AssemblyActionInvokeDataSource{}
}

type AssemblyActionInvokeDataSource struct {
	client *client.DatapowerClient
}

func (d *AssemblyActionInvokeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_assemblyactioninvoke"
}

func (d *AssemblyActionInvokeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Invoke assembly action",
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
						"url": schema.StringAttribute{
							MarkdownDescription: "URL",
							Computed:            true,
						},
						"ssl_client": schema.StringAttribute{
							MarkdownDescription: "TLS client profile",
							Computed:            true,
						},
						"timeout": schema.Int64Attribute{
							MarkdownDescription: "Timeout",
							Computed:            true,
						},
						"user_name": schema.StringAttribute{
							MarkdownDescription: "Username",
							Computed:            true,
						},
						"password": schema.StringAttribute{
							MarkdownDescription: "Password",
							Computed:            true,
						},
						"method": schema.StringAttribute{
							MarkdownDescription: "HTTP method",
							Computed:            true,
						},
						"backend_type": schema.StringAttribute{
							MarkdownDescription: "Backend type",
							Computed:            true,
						},
						"graph_ql_send_type": schema.StringAttribute{
							MarkdownDescription: "GraphQL send type",
							Computed:            true,
						},
						"compression": schema.BoolAttribute{
							MarkdownDescription: "Enable compression",
							Computed:            true,
						},
						"cache_type": schema.StringAttribute{
							MarkdownDescription: "Cache type",
							Computed:            true,
						},
						"time_to_live": schema.Int64Attribute{
							MarkdownDescription: "Time to live",
							Computed:            true,
						},
						"cache_unsafe_response": schema.BoolAttribute{
							MarkdownDescription: "Cache response to POST and PUT requests",
							Computed:            true,
						},
						"cache_key": schema.StringAttribute{
							MarkdownDescription: "Cache key",
							Computed:            true,
						},
						"follow_redirects": schema.BoolAttribute{
							MarkdownDescription: "Follow redirects",
							Computed:            true,
						},
						"http_version": schema.StringAttribute{
							MarkdownDescription: "HTTP version to server",
							Computed:            true,
						},
						"http2_required": schema.BoolAttribute{
							MarkdownDescription: "HTTP/2 required",
							Computed:            true,
						},
						"do_chunked_upload": schema.BoolAttribute{
							MarkdownDescription: "Allow chunked uploads",
							Computed:            true,
						},
						"persistent_connection": schema.BoolAttribute{
							MarkdownDescription: "Persistent connection",
							Computed:            true,
						},
						"stop_on_error": schema.BoolAttribute{
							MarkdownDescription: "Stop on error",
							Computed:            true,
						},
						"error_types": models.GetDmInvokeErrorTypeDataSourceSchema("Error types", "error-types", ""),
						"output": schema.StringAttribute{
							MarkdownDescription: "Output",
							Computed:            true,
						},
						"decode_request_params": schema.BoolAttribute{
							MarkdownDescription: "Decode request parameters",
							Computed:            true,
						},
						"encode_plus_char": schema.BoolAttribute{
							MarkdownDescription: "Encode + characters in query",
							Computed:            true,
						},
						"keep_payload": schema.BoolAttribute{
							MarkdownDescription: "Keep payload",
							Computed:            true,
						},
						"inject_user_agent_header": schema.BoolAttribute{
							MarkdownDescription: "Inject User-Agent header",
							Computed:            true,
						},
						"inject_proxy_headers": schema.BoolAttribute{
							MarkdownDescription: "Inject proxy headers",
							Computed:            true,
						},
						"header_control_list": schema.StringAttribute{
							MarkdownDescription: "Header control list",
							Computed:            true,
						},
						"parameter_control_list": schema.StringAttribute{
							MarkdownDescription: "Parameter control list",
							Computed:            true,
						},
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"title": schema.StringAttribute{
							MarkdownDescription: "Title",
							Computed:            true,
						},
						"correlation_path": schema.StringAttribute{
							MarkdownDescription: "Correlation path",
							Computed:            true,
						},
						"action_debug": schema.BoolAttribute{
							MarkdownDescription: "Enable debugging",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *AssemblyActionInvokeDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *AssemblyActionInvokeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AssemblyActionInvokeList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.AssemblyActionInvoke{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.AssemblyActionInvoke{}
	if value := res.Get(`AssemblyActionInvoke`); value.Exists() {
		for _, v := range value.Array() {
			item := models.AssemblyActionInvoke{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.AssemblyActionInvokeObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.AssemblyActionInvokeObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
