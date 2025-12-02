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
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type WebSphereJMSServerList struct {
	ProviderTarget types.String `tfsdk:"provider_target"`
	Id             types.String `tfsdk:"id"`
	AppDomain      types.String `tfsdk:"app_domain"`
	Result         types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &WebSphereJMSServerDataSource{}
	_ datasource.DataSourceWithConfigure = &WebSphereJMSServerDataSource{}
)

func NewWebSphereJMSServerDataSource() datasource.DataSource {
	return &WebSphereJMSServerDataSource{}
}

type WebSphereJMSServerDataSource struct {
	pData *tfutils.ProviderData
}

func (d *WebSphereJMSServerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_web_sphere_jms_server"
}

func (d *WebSphereJMSServerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Define the connection details to the remote WebSphere JMS server environment. This configuration is responsible for JMS messaging by periodically monitoring and polling queues to ensure that sent messages are directed to the correct receive queue or are routed to another queue manager.",
		Attributes: map[string]schema.Attribute{
			"provider_target": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Target host to retrieve this data from. If not set, provider will use the top level settings.", "", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
			},
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
						"endpoint": schema.ListNestedAttribute{
							MarkdownDescription: "<p>Specify JMS nondefault bootstrap server endpoints by host, port, and protocol to use for the bootstrap process.</p><p>A service integration bus (SIB) supports applications that use message-based and service-oriented architectures. A bus is a group of interconnected servers and clusters that were added as members. Applications connect to a bus at one of the messaging engines that is associated with its bus members.</p><p>A messaging engine is a component that runs inside a server. The messaging engine manages messaging resources for a bus member. Applications connect to a messaging engine when a SIB is accessed.</p><p>Applications, such as the WebSphere JMS message provider, that run outside the WebSphere Application Server environment cannot directly locate a suitable messaging engine to connect to the target bus. In such cases, the remote clients or servers must access the bus through a bootstrap server that is a member of the target bus. A bootstrap server is an application server that runs the SIB process. The bootstrap server does not need to run any message engines. Rather the bootstrap server selects a messaging engine that is running in an application server that supports the bootstrap protocol that was requested.</p><p>The connection sequence to a messaging engine is as follows. <ol><li>The remote application first connects to a bootstrap server.</li><li>The bootstrap server selects a messaging engine.</li><li>The bootstrap server tells the client application to connect to that message engine to gain bus access.</li></ol></p><p>A bootstrap server uses a host name or IP address, with a port number and a bootstrap transport chain. The transport chain identifies the protocol stack that is offered by the bootstrap server to define an endpoint address.</p><p>The protocol stack that is used to access the bootstrap server does not need to be the same protocol stack that is used for actual message transfer over the bus.</p><p>You can add multiple nondefault bootstrap servers. For failover capability, the endpoints must be members of the same WebSphere Application Server cluster.</p>",
							NestedObject:        models.GetDmWebSphereJMSEndpointDataSourceSchema(),
							Computed:            true,
						},
						"target_transport_chain": schema.StringAttribute{
							MarkdownDescription: "<p>Specify the predefined transport chain to use for WebSphere JMS message exchanges. The transport chain used for message exchange need not match the chain used for bootstrap access.</p><p>This property takes one of the following values. The default value is <tt>InboundBasicMessaging</tt> . <dl><dt><tt>InboundBasicMessaging</tt></dt><dd>Specifies the predefined <tt>InboundBasicMessaging</tt> transport chain (JFAP-TCP/IP).</dd><dt><tt>InboundHTTPMessaging</tt></dt><dd>Specifies the predefined <tt>InboundHTTPMessaging</tt> transport chain (tunnels JFAP using HTTP wrappers).</dd><dt><tt>InboundSecureMessaging</tt></dt><dd>Specifies the predefined <tt>InboundSecureMessaging</tt> transport chain (JFAP-SSL-TCP/IP).</dd></dl></p><p>If you have access to the WebSphere Administrative Console, you can view transport chain information through the <b>Application Servers/serverName/Transport Chain</b> menu.</p>",
							Computed:            true,
						},
						"messaging_bus": schema.StringAttribute{
							MarkdownDescription: "<p>Specify the SIB to access the remote WebSphere JMS service provider.</p><p>A SIB supports applications that use message-based and service-oriented architectures. A bus is a group of interconnected servers and clusters that were added as members of the bus. Applications connect to a bus at one of the messaging engines that is associated with its bus members.</p><p>If you have access to the WebSphere Application Server administrative console, you can view bus information through the <b>Service integration/Buses</b> menu. This information includes bus members and messaging engines, queues, topics, and the bus-specific default topic space.</p>",
							Computed:            true,
						},
						"ssl_cipher": schema.StringAttribute{
							MarkdownDescription: "Specify the IBM cipher specification when the TLS profile establishes a secure connection with the WebSphere JMS message provider. When you specify a TLS profile, the cipher suite that is associated with the profile is replaced by an IBM default cipher specification. The default value is <tt>SSL_RSA_WITH_NULL_MD5</tt> .",
							Computed:            true,
						},
						"fips": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to require the assignment of a FIPS-compliant cipher for the WebSphere JMS message provider. By default, a FIPS-compliant cipher is not required. When enabled, which prevents the use of non-FIPs compliance ciphers, requires the use of one of the following ciphers. <ul><li><tt>TLS_RSA_WITH_AES_128_CBC_SHA</tt></li><li><tt>TLS_RSA_WITH_AES_256_CBC_SHA</tt></li><li><tt>TLS_RSA_WITH_AES_128_CBC_SHA256</tt></li><li><tt>TLS_RSA_WITH_AES_256_CBC_SHA256</tt></li></ul>",
							Computed:            true,
						},
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"user_name": schema.StringAttribute{
							MarkdownDescription: "Username",
							Computed:            true,
						},
						"password_alias": schema.StringAttribute{
							MarkdownDescription: "Password alias",
							Computed:            true,
						},
						"transactional": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to acknowledge messages only after the transaction succeeds. When enabled, messages are acknowledged only after the transaction succeeds. By default, transaction-based processing is disabled.",
							Computed:            true,
						},
						"memory_threshold": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum memory allocation for pending messages in bytes. Enter a value in the range 10485760 - 1073741824. The default value is 268435456.",
							Computed:            true,
						},
						"maximum_message_size": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum message size in bytes. Enter a value in the range 0 - 1073741824. The default value is 1048576. A value of 0 disables the enforcement of a maximum message size.",
							Computed:            true,
						},
						"default_message_type": schema.StringAttribute{
							MarkdownDescription: "<p>Specify the default JMS message type. This message type is provided only when the message type cannot be determined from the JMS message headers. By default, the message payload is accessed as a Java byte array.</p><p>On the transaction level, the <tt>DP_JMSMessageType</tt> header can change this setting.</p>",
							Computed:            true,
						},
						"total_connection_limit": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum number of simultaneous open connections. Enter a value in the range 1 - 65535. The default value is 64.",
							Computed:            true,
						},
						"sessions_per_connection": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum number of concurrent multiplexed sessions that a single connection can support. Enter a value in the range 1 - 65535. The default is 100. <p>Session requests in excess of the defined value trigger the establishment of a new connection to the server. A new connection cannot be established unless the number of current connections is less than the value set for total connections.</p><p>For example, with values of 20 sessions per connection, 5 total connections, and 3 active fully subscribed connections, a new session request generates the establishment of a fourth connection.</p>",
							Computed:            true,
						},
						"auto_retry": schema.BoolAttribute{
							MarkdownDescription: "Specify how to control the automatic error-recovery procedure for the connection. By default, automatic error-recovery is enabled to reestablish a connection that was broken in response to an error condition.",
							Computed:            true,
						},
						"retry_interval": schema.Int64Attribute{
							MarkdownDescription: "Specify the interval between attempts to reestablish a connection in seconds. Enter a value in the range 1 - 65535. The default value is 1.",
							Computed:            true,
						},
						"enable_logging": schema.BoolAttribute{
							MarkdownDescription: "Specify how to control JMS-specific expanded logging. By default, expanded logging is disabled.",
							Computed:            true,
						},
						"ssl_client_config_type": schema.StringAttribute{
							MarkdownDescription: "TLS client type",
							Computed:            true,
						},
						"ssl_client": schema.StringAttribute{
							MarkdownDescription: "TLS client profile",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
						"provider_target": schema.StringAttribute{
							MarkdownDescription: tfutils.NewAttributeDescription("Target host to retrieve this data from. If not set, provider will use the top level settings.", "", "").String,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *WebSphereJMSServerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *WebSphereJMSServerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data WebSphereJMSServerList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !d.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), d.pData.Client.GetTargetNames()))
		return
	}
	o := models.WebSphereJMSServer{
		AppDomain: data.AppDomain,
	}

	path := o.GetPath()
	if !data.Id.IsNull() {
		path = path + "/" + data.Id.ValueString()
	}

	res, err := d.pData.Client.Get(path, data.ProviderTarget)
	resFound := true
	if err != nil {
		if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(d.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
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
	l := []models.WebSphereJMSServer{}
	if resFound {
		if value := res.Get(`WebSphereJMSServer`); value.Exists() {
			for _, v := range value.Array() {
				item := models.WebSphereJMSServer{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.WebSphereJMSServerObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.WebSphereJMSServerObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
