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

type HTTPUserAgentList struct {
	Id        types.String `tfsdk:"id"`
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
	resp.TypeName = req.ProviderTypeName + "_http_user_agent"
}

func (d *HTTPUserAgentDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "A user agent define how to retrieve resources from remote servers.",
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
						"identifier": schema.StringAttribute{
							MarkdownDescription: "Specify the string that the user agent includes as the <tt>request-header</tt> field. This field contains information about the user agent that initiates the request. By default, the system does not include a <tt>request-header</tt> field.",
							Computed:            true,
						},
						"max_redirects": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum number of HTTP redirect messages received before the target URL is declared unreachable. Enter a value in the range 0 - 128. The default value is 8.",
							Computed:            true,
						},
						"timeout": schema.Int64Attribute{
							MarkdownDescription: "the maximum idle time in seconds before an established connection to a remote server is torn down. Enter a avlue in the range 1 - 86400. The default value is 300.",
							Computed:            true,
						},
						"proxy_policies": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the proxy policy that associates a set of URLs with a specific HTTP proxy.",
							NestedObject:        models.GetDmProxyPolicyDataSourceSchema(),
							Computed:            true,
						},
						"ssl_policies": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the policy that associates a set of URLs with a specific TLS profile. When a URL matches the expression, the agent uses the corresponding TLS profile to secure connections with the resource.",
							NestedObject:        models.GetDmSSLPolicyDataSourceSchema(),
							Computed:            true,
						},
						"basic_auth_policies": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the policy that associates a set of URLs with a specific username and password for basic authentication.",
							NestedObject:        models.GetDmBasicAuthPolicyDataSourceSchema(),
							Computed:            true,
						},
						"soap_action_policies": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the policy that associates a set of URLs with a specific HTTP SOAPAction header.",
							NestedObject:        models.GetDmSoapActionPolicyDataSourceSchema(),
							Computed:            true,
						},
						"pubkey_auth_policies": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the policy that associates a set of URLs with a specific private key for public key authentication. The remote host must possess and reference the corresponding public key (certificate) to connect successfully.",
							NestedObject:        models.GetDmPubkeyAuthPolicyDataSourceSchema(),
							Computed:            true,
						},
						"allow_compression_policies": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the policy that associates a set of URLS that allow compression.",
							NestedObject:        models.GetDmAllowCompressionPolicyDataSourceSchema(),
							Computed:            true,
						},
						"header_retention_policies": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the policy that associates a set of URLS to retain specific heads in messages.",
							NestedObject:        models.GetDmHeaderRetentionPolicyDataSourceSchema(),
							Computed:            true,
						},
						"http_version_policies": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the policy that associates a set of URLs to specific HTTP versions. This policy is cumulative. If any transaction, URL match, or gateway have an HTTP version policy, that transaction is processed at the requested HTTP version.",
							NestedObject:        models.GetDmHTTPVersionPolicyDataSourceSchema(),
							Computed:            true,
						},
						"add_header_policies": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the policy that associates a set of URLS to inject HTTP headers into the message.",
							NestedObject:        models.GetDmAddHeaderPolicyDataSourceSchema(),
							Computed:            true,
						},
						"upload_chunked_policies": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the policy that associates a set of URL to control whether to send chunked-encoded documents. With HTTP/1.1, the body of the document can be delimited by <tt>Content-Length</tt> or chunked encoding. All servers understand <tt>Content-Length</tt> but many applications fail to understand chunked encoding. Therefore, <tt>Content-Length</tt> is used. However, the use of <tt>Content-Length</tt> interferes with the ability of the service to fully stream. <p>Unlike all other HTTP/1.1 features that can be negotiated down at run time, you must know beforehand that the target server is RFC 2616 compatible.</p>",
							NestedObject:        models.GetDmUploadChunkedPolicyDataSourceSchema(),
							Computed:            true,
						},
						"ftp_policies": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the policy that associate a set of URLs to control FTP client options for outgoing connections. These settings override the compiled-in defaults and can be further overridden by query parameters that initiates the file transfer.",
							NestedObject:        models.GetDmFTPPolicyDataSourceSchema(),
							Computed:            true,
						},
						"smtp_policies": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the policy that associates a set of URLS to control SMTP client options for outgoing connections. These settings override the compiled-in defaults and can be further overridden by query parameters that sends the e-mail message.",
							NestedObject:        models.GetDmSMTPPolicyDataSourceSchema(),
							Computed:            true,
						},
						"sftp_policies": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the policy that associate a set of URLs to control SSH client options for outgoing connections. These settings override the compiled-in defaults and can be further overridden by query parameters that initiates the file transfer.",
							NestedObject:        models.GetDmSFTPPolicyDataSourceSchema(),
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
	l := []models.HTTPUserAgent{}
	if resFound {
		if value := res.Get(`HTTPUserAgent`); value.Exists() {
			for _, v := range value.Array() {
				item := models.HTTPUserAgent{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
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
