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

type SSLSNIServerProfileList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &SSLSNIServerProfileDataSource{}
	_ datasource.DataSourceWithConfigure = &SSLSNIServerProfileDataSource{}
)

func NewSSLSNIServerProfileDataSource() datasource.DataSource {
	return &SSLSNIServerProfileDataSource{}
}

type SSLSNIServerProfileDataSource struct {
	pData *tfutils.ProviderData
}

func (d *SSLSNIServerProfileDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssl_sni_server_profile"
}

func (d *SSLSNIServerProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The TLS Server Name Indication (SNI) server profile secures connections with clients.",
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
						"sni_server_mapping": schema.StringAttribute{
							MarkdownDescription: "TLS hostname map",
							Computed:            true,
						},
						"sni_server_default": schema.StringAttribute{
							MarkdownDescription: "Specify the TLS server profile to process requests when the client does not send a <tt>ClientHello</tt> SNI extension. <p>The request is rejected when either of the following conditions apply. <ul><li>The client sends a <tt>ClientHello</tt> SNI extension that does not match a hostname in the map.</li><li>The client does not send a <tt>ClientHello</tt> SNI extension</li></ul></p>",
							Computed:            true,
						},
						"ssl_options": models.GetDmSSLOptionsDataSourceSchema("Specify the options to apply to the TLS connection that override settings in the TLS server profiles. These options have negative impact on the performance.", "ssl-options", ""),
						"max_ssl_duration": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum duration in seconds for an established TLS session. After the duration is reached, the TLS connection is closed. Enter a value in the range 1 - 691200. The default value is 3600.",
							Computed:            true,
						},
						"number_of_renegotiation_allowed": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum number of client initiated renegotiations. Enter a value in the range 0 - 512. The default value is 0, which indicates client initiated renegotiation is not allowed.",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *SSLSNIServerProfileDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *SSLSNIServerProfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data SSLSNIServerProfileList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.SSLSNIServerProfile{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.SSLSNIServerProfile{}
	if value := res.Get(`SSLSNIServerProfile`); value.Exists() {
		for _, v := range value.Array() {
			item := models.SSLSNIServerProfile{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.SSLSNIServerProfileObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.SSLSNIServerProfileObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
