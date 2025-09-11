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
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type SSLClientProfileList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &SSLClientProfileDataSource{}
	_ datasource.DataSourceWithConfigure = &SSLClientProfileDataSource{}
)

func NewSSLClientProfileDataSource() datasource.DataSource {
	return &SSLClientProfileDataSource{}
}

type SSLClientProfileDataSource struct {
	pData *tfutils.ProviderData
}

func (d *SSLClientProfileDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssl_client_profile"
}

func (d *SSLClientProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "A TLS client profile define how to secure a connection remote endpoint.",
		Attributes: map[string]schema.Attribute{

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
						"protocols": models.GetDmSSLProtoVersionsBitmapDataSourceSchema("Protocols", "protocols", ""),
						"ciphers": schema.ListAttribute{
							MarkdownDescription: "Specify the list of cipher suites to support in preference order. Ensure that the displayed order of cipher suites are in preference order. Otherwise, change the sequence to meet your preference.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"idcred": schema.StringAttribute{
							MarkdownDescription: "Identification credentials",
							Computed:            true,
						},
						"validate_server_cert": schema.BoolAttribute{
							MarkdownDescription: "Validate server certificate",
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
							MarkdownDescription: "Specify the duration in seconds that TLS sessions remain in the session cache before they are removed. Enter a value in the range 1 - 86400. The default value is 300.",
							Computed:            true,
						},
						"cache_size": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum number of entries of the session cache. Enter a value in the range 1 - 500000. The default value is 100.",
							Computed:            true,
						},
						"ssl_client_features": models.GetDmSSLClientFeaturesDataSourceSchema("Features", "ssl-client-features", ""),
						"elliptic_curves": schema.ListAttribute{
							MarkdownDescription: "Elliptic curves",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"use_custom_sni_hostname": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to use a custom server name in the SNI extension in the TLS client <tt>hello</tt> message. By default, the hostname of the target is used in the SNI extension.",
							Computed:            true,
						},
						"custom_sni_hostname": schema.StringAttribute{
							MarkdownDescription: "Custom SNI hostname",
							Computed:            true,
						},
						"validate_hostname": schema.BoolAttribute{
							MarkdownDescription: "Validate server hostname",
							Computed:            true,
						},
						"hostname_validation_flags": models.GetDmSSLHostnameValidationFlagsDataSourceSchema("Specify the flags that fine tune the validation methods and settings during the handshake. The default behavior uses the subject DN only when the <tt>Subject Alternative Name</tt> (SAN) extension contains no DNS name.", "hostname-validation-flags", ""),
						"hostname_validation_fail_on_error": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to terminate the handshake when hostname validation fails or to ignore the failure, log an event, and continue with server certificate validation. The default behavior is to ignore the failure, log an event, and continue with any configured server certificate validation.",
							Computed:            true,
						},
						"enable_tls13_compat": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to send extra TLS messages to look more like TLSv1.2. Some network middleboxes might not recognize TLSv1.3 and drop the connection. Enable this option to send dummy Change Cipher Spec messages, which makes TLSv1.3 look more like TLSv1.2.",
							Computed:            true,
						},
						"disable_renegotiation": schema.BoolAttribute{
							MarkdownDescription: "Disable renegotiation",
							Computed:            true,
						},
						"sig_algs": schema.ListAttribute{
							MarkdownDescription: "Specify the list of signature algorithms to advertise and support. An empty list implies the use of all of the default algorithms.",
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

func (d *SSLClientProfileDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *SSLClientProfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data SSLClientProfileList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.SSLClientProfile{
		AppDomain: data.AppDomain,
	}

	path := o.GetPath()
	if !data.Id.IsNull() {
		path = path + "/" + data.Id.ValueString()
	}

	res, err := d.pData.Client.Get(path)
	resFound := true
	if err != nil {
		if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		} else {
			resFound = false
		}
	}
	l := []models.SSLClientProfile{}
	if resFound {
		if value := res.Get(`SSLClientProfile`); value.Exists() {
			for _, v := range value.Array() {
				item := models.SSLClientProfile{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.SSLClientProfileObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.SSLClientProfileObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
