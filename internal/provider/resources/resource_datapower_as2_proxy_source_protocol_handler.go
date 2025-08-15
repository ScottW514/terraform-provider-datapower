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

package resources

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &AS2ProxySourceProtocolHandlerResource{}
var _ resource.ResourceWithValidateConfig = &AS2ProxySourceProtocolHandlerResource{}

func NewAS2ProxySourceProtocolHandlerResource() resource.Resource {
	return &AS2ProxySourceProtocolHandlerResource{}
}

type AS2ProxySourceProtocolHandlerResource struct {
	pData *tfutils.ProviderData
}

func (r *AS2ProxySourceProtocolHandlerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_as2_proxy_source_protocol_handler"
}

func (r *AS2ProxySourceProtocolHandlerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("A MEIG AS2 proxy handler receives AS2 requests over HTTP or HTTPS and forwards them to the back end which is assumed to be an IBM Multi-Enterprise Integration Gateway (MEIG) server. AS2 proxy handlers conform to gateway specifications of RFC 2616 and AS2 specification of RFC 4130.", "source-as2-proxy", "").AddActions("quiesce").String,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name of the object. Must be unique among object types in application domain.", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"app_domain": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The name of the application domain the object belongs to", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`), ""),
				},
				PlanModifiers: []planmodifier.String{
					modifiers.ImmutableAfterSet(),
				},
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"local_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the IP address or host alias that the service listens. The default value indicates that the service listens on all IP addresses.", "local-address", "").AddDefaultValue("0.0.0.0").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("0.0.0.0"),
			},
			"local_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the port that the service listens. Enter a value in the range 1 - 65535. The default value is 80.", "port", "").AddIntegerRange(1, 65535).AddDefaultValue("80").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(80),
			},
			"http_version": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the HTTP version for client connections. The default value is HTTP/1.1.", "http-client-version", "").AddStringEnum("HTTP/1.0", "HTTP/1.1").AddDefaultValue("HTTP/1.1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("HTTP/1.0", "HTTP/1.1"),
				},
				Default: stringdefault.StaticString("HTTP/1.1"),
			},
			"allowed_features": models.GetDmSourceAS2FeatureTypeResourceSchema("Allowed methods and versions", "allowed-features", "", false),
			"persistent_connections": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to negotiate persistent connections with clients. The default value is enabled.", "persistent-connections", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"max_persistent_connections_reuse": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum number of times a client can reuse a persistent connection. When the maximum reuse count is reached, an explicit <tt>HTTP Connection: close</tt> header is sent in the response and the TCP connection is closed. The default value is 0, which means unlimited reuse.", "max-persistent-reuse", "").String,
				Optional:            true,
			},
			"allow_compression": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to negotiate GZIP compression for client connections. The default value to not negoitate compression. When enabled and the <tt>Accept-Encoding</tt> HTTP header indicates that compressed documents can be processed, the service uses GZIP to compress HTTP transmissions. The <tt>Transfer-Encoding</tt> HTTP header indicates compression.", "compression", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"max_url_len": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the length in bytes of the longest incoming URL to accept. The length includes any query string or fragment identifier. Enter a value in the range 1 - 128000. The default value is 16384.", "max-url-len", "").AddIntegerRange(1, 128000).AddDefaultValue("16384").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 128000),
				},
				Default: int64default.StaticInt64(16384),
			},
			"max_total_hdr_len": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum aggregate length of incoming HTTP headers in bytes to allow. Enter a value in the range 5 - 128000. The default value is 128000.", "max-total-header-len", "").AddIntegerRange(5, 128000).AddDefaultValue("128000").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(5, 128000),
				},
				Default: int64default.StaticInt64(128000),
			},
			"max_hdr_count": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum number of headers to allow in requests. The default value is 0, which indicates no limit.", "max-header-count", "").AddIntegerRange(0, 65535).AddDefaultValue("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 65535),
				},
				Default: int64default.StaticInt64(0),
			},
			"max_name_hdr_len": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum length of header names in bytes to allow. Each HTTP header is expressed as a name-value pair. This setting specifies the maximum length of the name portion for HTTP headers. The default value is 0, which indicates no limit.", "max-header-name-len", "").AddIntegerRange(0, 65535).AddDefaultValue("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 65535),
				},
				Default: int64default.StaticInt64(0),
			},
			"max_value_hdr_len": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum length of HTTP header values in bytes to allow. Each HTTP header is expressed as a name-value pair. This setting specifies the maximum length of the value portion of that header. The default value is 0, which indicates no limit.", "max-header-value-len", "").AddIntegerRange(0, 65535).AddDefaultValue("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 65535),
				},
				Default: int64default.StaticInt64(0),
			},
			"acl": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the access control list to allow or deny access to this service based on the IP address of the client. When attached to a service, an access control list denies all access by default. To deny access to only selected addresses, first grant access to all addresses (allow 0.0.0.0) and then create deny entries for the desired hosts.", "acl", "access_control_list").String,
				Optional:            true,
			},
			"credential_charset": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the character encoding of the original basic authentication values. Basic authentication credentials are combined and base64 encoded in the authorization header of the request. The DataPower Gateway transcodes the contents of the authorization header to UTF-8. The default value is Protocol, which represents ISO-8859-1, Latin 1.", "credential-charset", "").AddStringEnum("protocol", "ascii", "utf8", "big5", "cp1250", "cp1251", "cp1252", "cp1253", "cp1254", "cp1255", "cp1256", "cp1257", "cp1258", "euc_jp", "euc_kr", "gb18030", "gb2312", "iso2022_jp", "iso2022_kr", "iso8859_1", "iso8859_2", "iso8859_4", "iso8859_5", "iso8859_6", "iso8859_7", "iso8859_8", "iso8859_9", "iso8859_15", "sjis", "tis620", "unicode_le").AddDefaultValue("protocol").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("protocol", "ascii", "utf8", "big5", "cp1250", "cp1251", "cp1252", "cp1253", "cp1254", "cp1255", "cp1256", "cp1257", "cp1258", "euc_jp", "euc_kr", "gb18030", "gb2312", "iso2022_jp", "iso2022_kr", "iso8859_1", "iso8859_2", "iso8859_4", "iso8859_5", "iso8859_6", "iso8859_7", "iso8859_8", "iso8859_9", "iso8859_15", "sjis", "tis620", "unicode_le"),
				},
				Default: stringdefault.StaticString("protocol"),
			},
			"remote_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the IP address, host name, or name of a load balancer group of the Multi-Enterprise Integration Gateway server.", "remote-hostname", "").String,
				Required:            true,
			},
			"remote_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the destination port on the Multi-Enterprise Integration Gateway server.", "remote-port", "").AddIntegerRange(1, 65535).String,
				Required:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 65535),
				},
			},
			"remote_connection_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration in seconds to wait to establish a connection with the server. Enter a value in the range 1 - 86400. The default value is 60.", "remote-connect-timeout", "").AddIntegerRange(1, 86400).AddDefaultValue("60").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(60),
			},
			"xml_manager": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies an existing XML manager. An XML manager obtains and manages XML documents, stylesheets, and other document resources on behalf of one or more services.", "xml-manager", "xml_manager").AddDefaultValue("default").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default"),
			},
			"enable_passthrough": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Controls whether to pass the original AS2 requests to the processing policy of DataPower service. <ul><li>When enabled, the AS2 proxy handler passes the original AS2 requests to DataPower service processing policy.</li><li>When disabled, the AS2 proxy handler first uses the cryptographic information in the partner exchange profile to decrypt the incoming AS2 requests and verify the signature. The AS2 proxy handler then passes the decrypted request body with signature removed to DataPower service for processing.</li></ul>", "enable-passthrough", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"enable_visibility_event": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Controls whether to send the visibility events generated by the AS2 proxy handler to the MEIG visibility event endpoint. These visibility events are correlated to those generated by the Multi-Enterprise Integration Gateway server in one transaction thread.", "enable-visibility-event", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"visibility_event_endpoint": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the URL of the MEIG visibility event endpoint. Enter the URL in the format of static IBM MQ URL that provides the information about the IBM MQ server name, queue manager name, and name of the channel configured in the Multi-Enterprise Integration Gateway server. For example, dpmq://NAME_OF_MQ_OBJECT/?RequestQueue=QUEUE_NAME_FOR_VISIBILITY_EVENT", "visibility-event-endpoint", "").String,
				Optional:            true,
			},
			"enable_hmac_authentication": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Controls whether to use Hash-based Message Authentication Code (HMAC) to secure all visibility events sent to the visibility event endpoint. If HMAC is enabled in the Multi-Enterprise Integration Gateway server, you must enable HMAC authentication in the AS2 proxy handler to avoid message rejection.", "enable-hmac-authentication", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"hmac_passphrase_alias": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the password alias of the passphrase used to calculate the HMAC token for message authentication and integrity checking in the Multi-Enterprise Integration Gateway server.", "hmac-passphrase-alias", "password_alias").String,
				Optional:            true,
			},
			"ssl_server_config_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The TLS profile type to secure connections between clients and the DataPower Gateway.", "ssl-config-type", "").AddStringEnum("server", "sni").AddDefaultValue("server").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("server", "sni"),
				},
				Default: stringdefault.StaticString("server"),
			},
			"ssl_server": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The TLS server profile to secure connections between clients and the DataPower Gateway.", "ssl-server", "ssl_server_profile").String,
				Optional:            true,
			},
			"sslsni_server": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The TLS SNI server profile to secure connections between clients and the DataPower Gateway.", "ssl-sni-server", "ssl_sni_server_profile").String,
				Optional:            true,
			},
			"ssl_client_config_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>The TLS profile type to secure connections between the DataPower Gateway and the remote Multi-Enterprise Integration Gateway server. This communication must be TLS protected. You can define the TLS proxy profile for this communication in one of the following places:</p><ul><li>Define the TLS profile in the user agent that is assigned to the XML manager for the DataPower service service.</li><li>Define the TLS profile here.</li></ul><p>Ensure that the TLS profile is defined in one of these places. Without the remote TLS profile, processing is stopped and an error is logged.</p>", "ssl-client-type", "").AddStringEnum("client").AddDefaultValue("client").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("client"),
				},
				Default: stringdefault.StaticString("client"),
			},
			"ssl_client": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "ssl_client_profile").String,
				Optional:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *AS2ProxySourceProtocolHandlerResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *AS2ProxySourceProtocolHandlerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AS2ProxySourceProtocolHandler
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `AS2ProxySourceProtocolHandler`)
	_, err := r.pData.Client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "POST", err))
		return
	}
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AS2ProxySourceProtocolHandlerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AS2ProxySourceProtocolHandler
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && (strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}

	if data.IsNull() {
		// Import
		data.FromBody(ctx, `AS2ProxySourceProtocolHandler`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `AS2ProxySourceProtocolHandler`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AS2ProxySourceProtocolHandlerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AS2ProxySourceProtocolHandler
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `AS2ProxySourceProtocolHandler`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Update)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AS2ProxySourceProtocolHandlerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AS2ProxySourceProtocolHandler
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *AS2ProxySourceProtocolHandlerResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.AS2ProxySourceProtocolHandler

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
