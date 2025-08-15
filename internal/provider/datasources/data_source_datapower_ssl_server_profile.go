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
	pData *tfutils.ProviderData
}

func (d *SSLServerProfileDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssl_server_profile"
}

func (d *SSLServerProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The TLS server profile secures connections with clients.",
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
							MarkdownDescription: "Specify the cipher suites to support in preference order. Ensure that the order of cipher suites are in preference order. Otherwise, change the sequence order to meet your preference.",
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
							MarkdownDescription: "Specify whether to require client authentication during the TLS handshake. When enabled, the handshake is aborted if the client certificate is not provided. Otherwise, the request does not fail when there is no client certificate.",
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
							MarkdownDescription: "Specify the number of seconds that TLS sessions are allowed to remain in the TLS session cache before they are removed. Enter a value in the range 1 - 86400. The default value is 300.",
							Computed:            true,
						},
						"cache_size": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum number of entries (multiplied by 1024) in the session cache. Enter a value in the range 1 - 500. The default value is 20.",
							Computed:            true,
						},
						"ssl_options": models.GetDmSSLOptionsDataSourceSchema("Specify the options to apply to the TLS connection. These options have negative impact on the performance.", "ssl-options", ""),
						"max_ssl_duration": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum duration of an established TLS session. The TLS connection terminates when the duration is reached. Enter a value in the range 1 - 11520. The default value is 60.",
							Computed:            true,
						},
						"disable_renegotiation": schema.BoolAttribute{
							MarkdownDescription: "Disable renegotiation",
							Computed:            true,
						},
						"number_of_renegotiation_allowed": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum number of client initiated renegotiations to allow. Enter a value in the range 0 - 512. The default value is 0, which indicates TLS client initiated renegotiation is not allowed.",
							Computed:            true,
						},
						"prohibit_resume_on_reneg": schema.BoolAttribute{
							MarkdownDescription: "Prohibit session resumption on renegotiation",
							Computed:            true,
						},
						"compression": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enable TLS compression. TLS compression in HTTPS is dangerous because the connection becomes vulnerable to the CRIME attack.",
							Computed:            true,
						},
						"allow_legacy_renegotiation": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to allow TLS renegotiation with TLS clients that do not support RFC 5746. By default, this support is disabled because renegotiation is vulnerable to man-in-the-middle attacks as documented in CVE-2009-3555. Renegotiation with TLS clients that support RFC 5746 is permitted regardless of the setting.",
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
							MarkdownDescription: "Specify whether to move ChaCha20-Poly1305 cipher to the top of preference list sent to the client when this cipher is at the top of client cipher list When server cipher suite preferences is in effect, enabling this property temporarily moves the ChaCha20-Poly1305 cipher to the top of preference list when clients that present ChaCha20-Poly1305 cipher have this cipher at the top of their preference list. This setting allows the client to negotiate ChaCha20-Poly1305 cipher while other clients can use other ciphers.",
							Computed:            true,
						},
						"sig_algs": schema.ListAttribute{
							MarkdownDescription: "Specify the signature algorithms to advertise and support. An empty list uses the default algorithms.",
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

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *SSLServerProfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data SSLServerProfileList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.SSLServerProfile{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
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
