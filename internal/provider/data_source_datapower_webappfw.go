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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
)

type WebAppFWList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &WebAppFWDataSource{}
	_ datasource.DataSourceWithConfigure = &WebAppFWDataSource{}
)

func NewWebAppFWDataSource() datasource.DataSource {
	return &WebAppFWDataSource{}
}

type WebAppFWDataSource struct {
	client *client.DatapowerClient
}

func (d *WebAppFWDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_webappfw"
}

func (d *WebAppFWDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Web Application Firewall",
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
						"priority": schema.StringAttribute{
							MarkdownDescription: "Service Priority",
							Computed:            true,
						},
						"front_side": schema.ListNestedAttribute{
							MarkdownDescription: "Source Addresses",
							NestedObject:        models.DmFrontSideDataSourceSchema,
							Computed:            true,
						},
						"remote_address": schema.StringAttribute{
							MarkdownDescription: "Remote Host",
							Computed:            true,
						},
						"remote_port": schema.Int64Attribute{
							MarkdownDescription: "Remote Port",
							Computed:            true,
						},
						"style_policy": schema.StringAttribute{
							MarkdownDescription: "Security Policy",
							Computed:            true,
						},
						"xml_manager": schema.StringAttribute{
							MarkdownDescription: "XML Manager",
							Computed:            true,
						},
						"error_policy": schema.StringAttribute{
							MarkdownDescription: "Error Policy",
							Computed:            true,
						},
						"uri_normalization": schema.BoolAttribute{
							MarkdownDescription: "Normalize URI",
							Computed:            true,
						},
						"rewrite_errors": schema.BoolAttribute{
							MarkdownDescription: "Rewrite Error Messages",
							Computed:            true,
						},
						"delay_errors": schema.BoolAttribute{
							MarkdownDescription: "Delay Error Messages",
							Computed:            true,
						},
						"delay_errors_duration": schema.Int64Attribute{
							MarkdownDescription: "Duration to Delay Error Messages",
							Computed:            true,
						},
						"stream_output_to_back": schema.StringAttribute{
							MarkdownDescription: "Stream Output to Back",
							Computed:            true,
						},
						"stream_output_to_front": schema.StringAttribute{
							MarkdownDescription: "Stream Output to Front",
							Computed:            true,
						},
						"front_timeout": schema.Int64Attribute{
							MarkdownDescription: "Front Side Timeout",
							Computed:            true,
						},
						"back_timeout": schema.Int64Attribute{
							MarkdownDescription: "Back Side Timeout",
							Computed:            true,
						},
						"front_persistent_timeout": schema.Int64Attribute{
							MarkdownDescription: "Front Persistent Timeout",
							Computed:            true,
						},
						"allow_cache_control_header": schema.BoolAttribute{
							MarkdownDescription: "Allow Cache-Control Header",
							Computed:            true,
						},
						"back_persistent_timeout": schema.Int64Attribute{
							MarkdownDescription: "Back Persistent Timeout",
							Computed:            true,
						},
						"front_http_version": schema.StringAttribute{
							MarkdownDescription: "HTTP Response Version",
							Computed:            true,
						},
						"back_http_version": schema.StringAttribute{
							MarkdownDescription: "HTTP Version to Server",
							Computed:            true,
						},
						"request_side_security": schema.BoolAttribute{
							MarkdownDescription: "Request Security",
							Computed:            true,
						},
						"response_side_security": schema.BoolAttribute{
							MarkdownDescription: "Response Security",
							Computed:            true,
						},
						"do_chunked_upload": schema.BoolAttribute{
							MarkdownDescription: "Allow Chunked Uploads",
							Computed:            true,
						},
						"follow_redirects": schema.BoolAttribute{
							MarkdownDescription: "Follow Redirects",
							Computed:            true,
						},
						"http_client_ip_label": schema.StringAttribute{
							MarkdownDescription: "HTTP Client IP Label",
							Computed:            true,
						},
						"http_log_cor_id_label": schema.StringAttribute{
							MarkdownDescription: "HTTP Global Transaction ID Label",
							Computed:            true,
						},
						"debug_mode": schema.StringAttribute{
							MarkdownDescription: "Probe setting",
							Computed:            true,
						},
						"debug_history": schema.Int64Attribute{
							MarkdownDescription: "Transaction History",
							Computed:            true,
						},
						"debug_trigger": schema.ListNestedAttribute{
							MarkdownDescription: "Probe Triggers",
							NestedObject:        models.DmMSDebugTriggerTypeDataSourceSchema,
							Computed:            true,
						},
						"url_rewrite_policy": schema.StringAttribute{
							MarkdownDescription: "Header and URL Rewrite Policy",
							Computed:            true,
						},
						"do_host_rewriting": schema.BoolAttribute{
							MarkdownDescription: "Rewrite Host Names When Gatewaying",
							Computed:            true,
						},
						"ssl_config_type": schema.StringAttribute{
							MarkdownDescription: "TLS type",
							Computed:            true,
						},
						"ssl_server": schema.StringAttribute{
							MarkdownDescription: "TLS server profile",
							Computed:            true,
						},
						"sslsni_server": schema.StringAttribute{
							MarkdownDescription: "TLS SNI server profile",
							Computed:            true,
						},
						"ssl_client": schema.StringAttribute{
							MarkdownDescription: "TLS client profile",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *WebAppFWDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *WebAppFWDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data WebAppFWList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.WebAppFW{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.WebAppFW{}
	if value := res.Get(`WebAppFW`); value.Exists() {
		for _, v := range value.Array() {
			item := models.WebAppFW{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.WebAppFWObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.WebAppFWObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
