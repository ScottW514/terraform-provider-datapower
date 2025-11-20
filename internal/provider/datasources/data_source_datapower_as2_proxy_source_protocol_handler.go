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

type AS2ProxySourceProtocolHandlerList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &AS2ProxySourceProtocolHandlerDataSource{}
	_ datasource.DataSourceWithConfigure = &AS2ProxySourceProtocolHandlerDataSource{}
)

func NewAS2ProxySourceProtocolHandlerDataSource() datasource.DataSource {
	return &AS2ProxySourceProtocolHandlerDataSource{}
}

type AS2ProxySourceProtocolHandlerDataSource struct {
	pData *tfutils.ProviderData
}

func (d *AS2ProxySourceProtocolHandlerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_as2_proxy_source_protocol_handler"
}

func (d *AS2ProxySourceProtocolHandlerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "A MEIG AS2 proxy handler receives AS2 requests over HTTP or HTTPS and forwards them to the back end which is assumed to be an IBM Multi-Enterprise Integration Gateway (MEIG) server. AS2 proxy handlers conform to gateway specifications of RFC 2616 and AS2 specification of RFC 4130.",
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
						"local_address": schema.StringAttribute{
							MarkdownDescription: "Specify the IP address or host alias that the service listens. The default value indicates that the service listens on all IP addresses.",
							Computed:            true,
						},
						"local_port": schema.Int64Attribute{
							MarkdownDescription: "Specify the port that the service listens. Enter a value in the range 1 - 65535. The default value is 80.",
							Computed:            true,
						},
						"http_version": schema.StringAttribute{
							MarkdownDescription: "Specify the HTTP version for client connections. The default value is HTTP/1.1.",
							Computed:            true,
						},
						"allowed_features": models.GetDmSourceAS2FeatureTypeDataSourceSchema("Allowed methods and versions", "allowed-features", ""),
						"persistent_connections": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to negotiate persistent connections with clients. The default value is enabled.",
							Computed:            true,
						},
						"max_persistent_connections_reuse": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum number of times a client can reuse a persistent connection. When the maximum reuse count is reached, an explicit <tt>HTTP Connection: close</tt> header is sent in the response and the TCP connection is closed. The default value is 0, which means unlimited reuse.",
							Computed:            true,
						},
						"allow_compression": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to negotiate GZIP compression for client connections. The default value to not negoitate compression. When enabled and the <tt>Accept-Encoding</tt> HTTP header indicates that compressed documents can be processed, the service uses GZIP to compress HTTP transmissions. The <tt>Transfer-Encoding</tt> HTTP header indicates compression.",
							Computed:            true,
						},
						"max_url_len": schema.Int64Attribute{
							MarkdownDescription: "Specify the length in bytes of the longest incoming URL to accept. The length includes any query string or fragment identifier. Enter a value in the range 1 - 128000. The default value is 16384.",
							Computed:            true,
						},
						"max_total_hdr_len": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum aggregate length of incoming HTTP headers in bytes to allow. Enter a value in the range 5 - 128000. The default value is 128000.",
							Computed:            true,
						},
						"max_hdr_count": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum number of headers to allow in requests. The default value is 0, which indicates no limit.",
							Computed:            true,
						},
						"max_name_hdr_len": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum length of header names in bytes to allow. Each HTTP header is expressed as a name-value pair. This setting specifies the maximum length of the name portion for HTTP headers. The default value is 0, which indicates no limit.",
							Computed:            true,
						},
						"max_value_hdr_len": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum length of HTTP header values in bytes to allow. Each HTTP header is expressed as a name-value pair. This setting specifies the maximum length of the value portion of that header. The default value is 0, which indicates no limit.",
							Computed:            true,
						},
						"acl": schema.StringAttribute{
							MarkdownDescription: "Specifies the access control list to allow or deny access to this service based on the IP address of the client. When attached to a service, an access control list denies all access by default. To deny access to only selected addresses, first grant access to all addresses (allow 0.0.0.0) and then create deny entries for the desired hosts.",
							Computed:            true,
						},
						"credential_charset": schema.StringAttribute{
							MarkdownDescription: "Specifies the character encoding of the original basic authentication values. Basic authentication credentials are combined and base64 encoded in the authorization header of the request. The DataPower Gateway transcodes the contents of the authorization header to UTF-8. The default value is Protocol, which represents ISO-8859-1, Latin 1.",
							Computed:            true,
						},
						"remote_address": schema.StringAttribute{
							MarkdownDescription: "Specifies the IP address, host name, or name of a load balancer group of the Multi-Enterprise Integration Gateway server.",
							Computed:            true,
						},
						"remote_port": schema.Int64Attribute{
							MarkdownDescription: "Specifies the destination port on the Multi-Enterprise Integration Gateway server.",
							Computed:            true,
						},
						"remote_connection_timeout": schema.Int64Attribute{
							MarkdownDescription: "Specify the duration in seconds to wait to establish a connection with the server. Enter a value in the range 1 - 86400. The default value is 60.",
							Computed:            true,
						},
						"xml_manager": schema.StringAttribute{
							MarkdownDescription: "Specifies an existing XML manager. An XML manager obtains and manages XML documents, stylesheets, and other document resources on behalf of one or more services.",
							Computed:            true,
						},
						"enable_passthrough": schema.BoolAttribute{
							MarkdownDescription: "Controls whether to pass the original AS2 requests to the processing policy of DataPower service. <ul><li>When enabled, the AS2 proxy handler passes the original AS2 requests to DataPower service processing policy.</li><li>When disabled, the AS2 proxy handler first uses the cryptographic information in the partner exchange profile to decrypt the incoming AS2 requests and verify the signature. The AS2 proxy handler then passes the decrypted request body with signature removed to DataPower service for processing.</li></ul>",
							Computed:            true,
						},
						"enable_visibility_event": schema.BoolAttribute{
							MarkdownDescription: "Controls whether to send the visibility events generated by the AS2 proxy handler to the MEIG visibility event endpoint. These visibility events are correlated to those generated by the Multi-Enterprise Integration Gateway server in one transaction thread.",
							Computed:            true,
						},
						"visibility_event_endpoint": schema.StringAttribute{
							MarkdownDescription: "Specifies the URL of the MEIG visibility event endpoint. Enter the URL in the format of static IBM MQ URL that provides the information about the IBM MQ server name, queue manager name, and name of the channel configured in the Multi-Enterprise Integration Gateway server. For example, dpmq://NAME_OF_MQ_OBJECT/?RequestQueue=QUEUE_NAME_FOR_VISIBILITY_EVENT",
							Computed:            true,
						},
						"enable_hmac_authentication": schema.BoolAttribute{
							MarkdownDescription: "Controls whether to use Hash-based Message Authentication Code (HMAC) to secure all visibility events sent to the visibility event endpoint. If HMAC is enabled in the Multi-Enterprise Integration Gateway server, you must enable HMAC authentication in the AS2 proxy handler to avoid message rejection.",
							Computed:            true,
						},
						"hmac_passphrase_alias": schema.StringAttribute{
							MarkdownDescription: "Specifies the password alias of the passphrase used to calculate the HMAC token for message authentication and integrity checking in the Multi-Enterprise Integration Gateway server.",
							Computed:            true,
						},
						"ssl_server_config_type": schema.StringAttribute{
							MarkdownDescription: "The TLS profile type to secure connections between clients and the DataPower Gateway.",
							Computed:            true,
						},
						"ssl_server": schema.StringAttribute{
							MarkdownDescription: "The TLS server profile to secure connections between clients and the DataPower Gateway.",
							Computed:            true,
						},
						"ssl_sni_server": schema.StringAttribute{
							MarkdownDescription: "The TLS SNI server profile to secure connections between clients and the DataPower Gateway.",
							Computed:            true,
						},
						"ssl_client_config_type": schema.StringAttribute{
							MarkdownDescription: "<p>The TLS profile type to secure connections between the DataPower Gateway and the remote Multi-Enterprise Integration Gateway server. This communication must be TLS protected. You can define the TLS proxy profile for this communication in one of the following places:</p><ul><li>Define the TLS profile in the user agent that is assigned to the XML manager for the DataPower service service.</li><li>Define the TLS profile here.</li></ul><p>Ensure that the TLS profile is defined in one of these places. Without the remote TLS profile, processing is stopped and an error is logged.</p>",
							Computed:            true,
						},
						"ssl_client": schema.StringAttribute{
							MarkdownDescription: "TLS client profile",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *AS2ProxySourceProtocolHandlerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *AS2ProxySourceProtocolHandlerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data AS2ProxySourceProtocolHandlerList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.AS2ProxySourceProtocolHandler{
		AppDomain: data.AppDomain,
	}

	path := o.GetPath()
	if !data.Id.IsNull() {
		path = path + "/" + data.Id.ValueString()
	}

	res, err := d.pData.Client.Get(path)
	resFound := true
	if err != nil {
		if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(d.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString())
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
	l := []models.AS2ProxySourceProtocolHandler{}
	if resFound {
		if value := res.Get(`AS2ProxySourceProtocolHandler`); value.Exists() {
			for _, v := range value.Array() {
				item := models.AS2ProxySourceProtocolHandler{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.AS2ProxySourceProtocolHandlerObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.AS2ProxySourceProtocolHandlerObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
