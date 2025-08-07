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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
)

type SSLServerProfileList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &SSLServerProfileDataSource{}
	_ datasource.DataSourceWithConfigure = &SSLServerProfileDataSource{}
)

func NewSSLServerProfileDataSource() datasource.DataSource {
	return &SSLServerProfileDataSource{}
}

type SSLServerProfileDataSource struct {
	client *client.DatapowerClient
}

func (d *SSLServerProfileDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslserverprofile"
}

func (d *SSLServerProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "TLS server profile",
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
						"protocols": models.GetDmSSLProtoVersionsBitmapDataSourceSchema("Protocols", "protocols", ""),
						"ciphers": schema.ListAttribute{
							MarkdownDescription: "Ciphers",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"idcred": schema.StringAttribute{
							MarkdownDescription: "Identification credentials",
							Computed:            true,
						},
						"request_client_auth": schema.BoolAttribute{
							MarkdownDescription: "Request client authentication",
							Computed:            true,
						},
						"require_client_auth": schema.BoolAttribute{
							MarkdownDescription: "Require client authentication",
							Computed:            true,
						},
						"validate_client_cert": schema.BoolAttribute{
							MarkdownDescription: "Validate client certificate",
							Computed:            true,
						},
						"send_client_auth_ca_list": schema.BoolAttribute{
							MarkdownDescription: "Send client authentication CA list",
							Computed:            true,
						},
						"valcred": schema.StringAttribute{
							MarkdownDescription: "Validation credentials",
							Computed:            true,
						},
						"caching": schema.BoolAttribute{
							MarkdownDescription: "Enable session caching",
							Computed:            true,
						},
						"cache_timeout": schema.Int64Attribute{
							MarkdownDescription: "Session cache timeout",
							Computed:            true,
						},
						"cache_size": schema.Int64Attribute{
							MarkdownDescription: "Session cache size",
							Computed:            true,
						},
						"ssl_options": models.GetDmSSLOptionsDataSourceSchema("Advanced TLS options", "ssl-options", ""),
						"max_ssl_duration": schema.Int64Attribute{
							MarkdownDescription: "Maximum TLS session duration",
							Computed:            true,
						},
						"disable_renegotiation": schema.BoolAttribute{
							MarkdownDescription: "Disable renegotiation",
							Computed:            true,
						},
						"number_of_renegotiation_allowed": schema.Int64Attribute{
							MarkdownDescription: "Maximum client initiated renegotiations",
							Computed:            true,
						},
						"prohibit_resume_on_reneg": schema.BoolAttribute{
							MarkdownDescription: "Prohibit session resumption on renegotiation",
							Computed:            true,
						},
						"compression": schema.BoolAttribute{
							MarkdownDescription: "Enable compression",
							Computed:            true,
						},
						"allow_legacy_renegotiation": schema.BoolAttribute{
							MarkdownDescription: "Allow legacy renegotiation",
							Computed:            true,
						},
						"prefer_server_ciphers": schema.BoolAttribute{
							MarkdownDescription: "Use server cipher suite preferences",
							Computed:            true,
						},
						"elliptic_curves": schema.ListAttribute{
							MarkdownDescription: "Elliptic curves",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"prioritize_cha_cha": schema.BoolAttribute{
							MarkdownDescription: "Prioritize ChaCha20-Poly1305 cipher",
							Computed:            true,
						},
						"sig_algs": schema.ListAttribute{
							MarkdownDescription: "Signature algorithms",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"require_closure_notification": schema.BoolAttribute{
							MarkdownDescription: "Require closure notification",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *SSLServerProfileDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *SSLServerProfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SSLServerProfileList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.SSLServerProfile{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.SSLServerProfile{}
	if value := res.Get(`SSLServerProfile`); value.Exists() {
		for _, v := range value.Array() {
			item := models.SSLServerProfile{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.SSLServerProfileObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.SSLServerProfileObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
