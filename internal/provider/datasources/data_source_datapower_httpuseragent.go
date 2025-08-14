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

type HTTPUserAgentList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &HTTPUserAgentDataSource{}
	_ datasource.DataSourceWithConfigure = &HTTPUserAgentDataSource{}
)

func NewHTTPUserAgentDataSource() datasource.DataSource {
	return &HTTPUserAgentDataSource{}
}

type HTTPUserAgentDataSource struct {
	pData *tfutils.ProviderData
}

func (d *HTTPUserAgentDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_httpuseragent"
}

func (d *HTTPUserAgentDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "User agent",
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
						"identifier": schema.StringAttribute{
							MarkdownDescription: "HTTP request-header",
							Computed:            true,
						},
						"max_redirects": schema.Int64Attribute{
							MarkdownDescription: "Max redirects",
							Computed:            true,
						},
						"timeout": schema.Int64Attribute{
							MarkdownDescription: "Timeout",
							Computed:            true,
						},
						"proxy_policies": schema.ListNestedAttribute{
							MarkdownDescription: "Proxy policy",
							NestedObject:        models.DmProxyPolicyDataSourceSchema,
							Computed:            true,
						},
						"ssl_policies": schema.ListNestedAttribute{
							MarkdownDescription: "TLS profile policy",
							NestedObject:        models.DmSSLPolicyDataSourceSchema,
							Computed:            true,
						},
						"basic_auth_policies": schema.ListNestedAttribute{
							MarkdownDescription: "Basic authentication policy",
							NestedObject:        models.DmBasicAuthPolicyDataSourceSchema,
							Computed:            true,
						},
						"soap_action_policies": schema.ListNestedAttribute{
							MarkdownDescription: "SOAPAction policy",
							NestedObject:        models.DmSoapActionPolicyDataSourceSchema,
							Computed:            true,
						},
						"pubkey_auth_policies": schema.ListNestedAttribute{
							MarkdownDescription: "Public key authentication policy",
							NestedObject:        models.DmPubkeyAuthPolicyDataSourceSchema,
							Computed:            true,
						},
						"allow_compression_policies": schema.ListNestedAttribute{
							MarkdownDescription: "Allow compression policy",
							NestedObject:        models.DmAllowCompressionPolicyDataSourceSchema,
							Computed:            true,
						},
						"header_retention_policies": schema.ListNestedAttribute{
							MarkdownDescription: "Header retention policy",
							NestedObject:        models.DmHeaderRetentionPolicyDataSourceSchema,
							Computed:            true,
						},
						"http_version_policies": schema.ListNestedAttribute{
							MarkdownDescription: "HTTP version policy",
							NestedObject:        models.DmHTTPVersionPolicyDataSourceSchema,
							Computed:            true,
						},
						"add_header_policies": schema.ListNestedAttribute{
							MarkdownDescription: "Header injection policy",
							NestedObject:        models.DmAddHeaderPolicyDataSourceSchema,
							Computed:            true,
						},
						"upload_chunked_policies": schema.ListNestedAttribute{
							MarkdownDescription: "Chunked upload policy",
							NestedObject:        models.DmUploadChunkedPolicyDataSourceSchema,
							Computed:            true,
						},
						"ftp_policies": schema.ListNestedAttribute{
							MarkdownDescription: "FTP client policy",
							NestedObject:        models.DmFTPPolicyDataSourceSchema,
							Computed:            true,
						},
						"smtp_policies": schema.ListNestedAttribute{
							MarkdownDescription: "SMTP client policy",
							NestedObject:        models.DmSMTPPolicyDataSourceSchema,
							Computed:            true,
						},
						"sftp_policies": schema.ListNestedAttribute{
							MarkdownDescription: "SFTP client policy",
							NestedObject:        models.DmSFTPPolicyDataSourceSchema,
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *HTTPUserAgentDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *HTTPUserAgentDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data HTTPUserAgentList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.HTTPUserAgent{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.HTTPUserAgent{}
	if value := res.Get(`HTTPUserAgent`); value.Exists() {
		for _, v := range value.Array() {
			item := models.HTTPUserAgent{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.HTTPUserAgentObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.HTTPUserAgentObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
