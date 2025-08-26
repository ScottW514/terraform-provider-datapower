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

type MQManagerList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &MQManagerDataSource{}
	_ datasource.DataSourceWithConfigure = &MQManagerDataSource{}
)

func NewMQManagerDataSource() datasource.DataSource {
	return &MQManagerDataSource{}
}

type MQManagerDataSource struct {
	pData *tfutils.ProviderData
}

func (d *MQManagerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mq_manager"
}

func (d *MQManagerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "In IBM MQ, distributed send and receive queues are managed by a component called a queue manager. The queue manager provides messaging services for communicating applications by periodically monitoring/polling queues, by ensuring that sent messages are directed to the correct receive queue, or that messages are routed to another queue manager. This queue manager object corresponds to a queue manager that is running on an IBM MQ server in your network. These properties enable communication between the DataPower Gateway and the queue manager on the remote IBM MQ server.",
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
						"host_name": schema.StringAttribute{
							MarkdownDescription: "<p>Specify the hostname or IP address of the IBM MQ server with its listening port. This value indicates where the queue manager is running. If you do not specify the port, the default is 1414.</p><p>The value depends on IP family.</p><ul><li>For IPv4 <ul><li>hostname:port -- For example, server1:1414</li><li>hostname(port) -- For example, server1(1414)</li><li>hostname -- For example, server1</li><li>address:port -- For example, 10.10.1.2:1414</li><li>address(port) -- For example, 10.10.1.2(1414)</li><li>address -- For example, 10.10.1.2</li></ul></li><li>For IPv6 <ul><li>hostname:port -- For example, server1:1414</li><li>hostname(port) -- For example, server1(1414)</li><li>hostname -- For example, server1</li><li>[address]:port -- For example, [2202::148:248]:1414</li><li>address(port) -- For example, 2202::148:248(1414)</li><li>address -- For example, 2202::148:248</li></ul></li></ul>",
							Computed:            true,
						},
						"q_mname": schema.StringAttribute{
							MarkdownDescription: "Specify the name of the queue manager when the queue manager is not the default on the identified host.",
							Computed:            true,
						},
						"ccsid": schema.Int64Attribute{
							MarkdownDescription: "Specify the coded character set identifier (CCSID) to present to the remote queue manager during the connection. This setting has the same effect as setting the MQCCSID environment variable for an IBM MQ client. For more information, see the IBM MQ documentation. Unless necessary, retain the default value.",
							Computed:            true,
						},
						"channel_name": schema.StringAttribute{
							MarkdownDescription: "Specify the name of the channel to use as an alternative to the default SYSTEM.DEF.SVRCONN.",
							Computed:            true,
						},
						"csp_user_id": schema.StringAttribute{
							MarkdownDescription: "<p>Specify the user ID value of the MQCSP connection security parameter when the MQCSP structure is used for authorization service.</p><p>MQCSP support enables the authorization service to authenticate a user ID and password. You can specify the MQCSP connection security parameters structure on an MQCONNX call. Before you use MQCSP support, you must define a security exit in the queue manager on the IBM MQ server. Ensure that your MQCSP user ID and password in the security exit are consistent with what you define for the local queue manager. Either an inconsistent MQCSP user ID or an inconsistent password causes a connection failure with the IBM MQ server.</p><p><b>Notes:</b><ul><li>If neither user ID or password is defined, connects to the IBM MQ server without MQCSP settings.</li><li>If only one is defined, a warning occurs and the local queue manager is not up.</li><li>If both are defined but one is inconsistent with the value on the IBM MQ server, the connection fails and the local queue manager is not up.</li></ul></p>",
							Computed:            true,
						},
						"csp_password_alias": schema.StringAttribute{
							MarkdownDescription: "<p>Specify the password alias of the password value of the MQCSP connection security parameter when the MQCSP structure is used for authorization service.</p><p>MQCSP support enables the authorization service to authenticate a user ID and password. You can specify the MQCSP connection security parameters structure on an MQCONNX call. Before you use MQCSP support, you must define a security exit in the queue manager on the IBM MQ server. Ensure that your MQCSP user ID and password in the security exit are consistent with what you define for the local queue manager. Either an inconsistent MQCSP user ID or an inconsistent password causes a connection failure with the IBM MQ server.</p><p><b>Notes:</b><ul><li>If neither user ID or password is defined, connects to the IBM MQ server without MQCSP settings.</li><li>If only one is defined, a warning occurs and the local queue manager is not up.</li><li>If both are defined but one is inconsistent with the value on the IBM MQ server, the connection fails and the local queue manager is not up.</li></ul></p>",
							Computed:            true,
						},
						"heartbeat": schema.Int64Attribute{
							MarkdownDescription: "Specify the approximate time in seconds between heartbeat flows on a channel when waiting for a message on a queue. Enter a value in the range 0 - 999999. If 0, there in no heartbeat flows exchanged when waiting for a message on the channel. This property does not set the heartbeat on the channel. Instead, it is used to negotiate the heartbeat value with the channel. The greater of the two values is used.",
							Computed:            true,
						},
						"maximum_message_size": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum size in bytes of allowed messages. Enter a value that is equal to or greater than the <tt>MaxMsgLength</tt> attribute of the channel and of the queue on the IBM MQ server. Messages that are greater than this size are rejected. Enter a value in the range 1024 - 104857600.",
							Computed:            true,
						},
						"cache_timeout": schema.Int64Attribute{
							MarkdownDescription: "<p>Specify the duration in seconds to retain (keep alive) a dynamic connection in the connection cache. Enter a value in the range 0 - 65535. The default value is 60. Enter 0 to disable the timer.</p><p>Use a value that is greater than the negotiated heartbeat interval but less than the keep alive interval.</p><ul><li>The negotiated heartbeat interval is between the client and the IBM MQ server. The channel heartbeat defines the starting value for the negotiation.</li><li>The keep alive (timeout) interval is on the IBM MQ server. The <tt>KAINT</tt> attribute on the IBM MQ server defines the timeout value for a channel. Not all channels have a defined, explicit keep alive interval on the IBM MQ server. Some queue managers use an automatic timeout setting when the <tt>KAINT</tt> attribute set to <tt>AUTO</tt> . In these cases, the keep alive interval is the negotiated heartbeat interval plus 60 seconds.</li></ul><p>When an inactive connection reaches this threshold, the dynamic connection is removed from the local cache. When the cache no longer contains dynamic connections, the client deletes the dynamic queue manager. Without a dynamic queue manager, there is no connection with the IBM MQ server.</p><p>The cache timeout value is the only way to configure a timeout value with the IBM MQ server. No other local configuration setting can time out an IBM MQ connection.</p>",
							Computed:            true,
						},
						"ffst_size": schema.Int64Attribute{
							MarkdownDescription: "FFST file size",
							Computed:            true,
						},
						"ffst_rotate": schema.Int64Attribute{
							MarkdownDescription: "Number of FFST file rotations",
							Computed:            true,
						},
						"units_of_work": schema.Int64Attribute{
							MarkdownDescription: "Specify whether to use units-of-work. <ul><li>When 0, the default value, get and put messages without provisions for rollback. Either the operation succeeds or not. Undeliverable messages are silently discarded, which leaves higher level protocols with the responsibility to detect and retransmit lost packets</li><li>When 1, use syncpoints. A syncpoint commits and rolls back each message, not the entire transaction. When specified, the local queue manager does not remove the message that it gets from a server queue until it completes its transaction by using that message (such as placing the message on a server queue for processing). If the transaction fails and the message is left available on the server queue, the local queue manager can attempt to get the message from the server queue and process it again.</li></ul>",
							Computed:            true,
						},
						"automatic_backout": schema.BoolAttribute{
							MarkdownDescription: "<p>Specify whether to enable automatic routing of undeliverable messages, which is the automatic backout of poison messages. A poison message is any message that the receiving application does not know how to process.</p><p>Usually an application rolls back the get of this message, which leaves the message on the input queue. However, the backout count ( <tt>MQMD.Backoutcount</tt> ) is incremented. As the queue manager continues to re-get the message, the backout count continues to increase. When the backout count exceeds the backout threshold, the queue manager moves the message to the backout queue.</p><p>When disabled, the poison message remains on the get queue and continues to be reprocessed by the client until the server queue manager that manages the get queue removes it or the client reroutes the offending message. The message can be rerouted by a custom stylesheet in the request rule.</p>",
							Computed:            true,
						},
						"backout_threshold": schema.Int64Attribute{
							MarkdownDescription: "Specify the total number of processing attempts. After all attempts fail, the following actions occur. <ul><li>The poison message is moved to the backout queue.</li><li>The unit of work that contains this message is committed.</li></ul><p>Enter a value that is equal to or greater than 1.</p>",
							Computed:            true,
						},
						"backout_queue_name": schema.StringAttribute{
							MarkdownDescription: "Specify the name of the queue for undeliverable messages. This queue contains messages that exceeded the backout threshold. This queue must be managed by the same queue manager on the IBM MQ server as the defined GET queue. The backout queue, typically <tt>SYSTEM.DEAD.LETTER.QUEUE</tt> contains messages that cannot be processed or delivered.",
							Computed:            true,
						},
						"total_connection_limit": schema.Int64Attribute{
							MarkdownDescription: "Specify the total number of open TCP connections to allow. Enter a value in the range 1 - 10000. The default value is 250.",
							Computed:            true,
						},
						"initial_connections": schema.Int64Attribute{
							MarkdownDescription: "Specify the number of TCP connections to open immediately with the IBM MQ server. Enter a value in the range 0 - 10000. The default value is 1.",
							Computed:            true,
						},
						"sharing_conversations": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum number of conversations to share a single TCP connection. Enter a value in the range 1 - 5000. The default value is 1.",
							Computed:            true,
						},
						"ssl_key": schema.StringAttribute{
							MarkdownDescription: "<p>Specify the location of the key database file where keys and certificates are stored.</p><p>Use this setting with the TLS cipher setting to enable TLS communication when the TLS artifacts were created with IBM Global Security Kit (GSKit).</p><p><b>Note:</b> To integrate with IBM MQ for z/OS, do not use these properties. Use a TLS client profile.</p><p>Each key database file has an associated stash file. The stash file holds encrypted passwords that allow programmatic access to the key database. The stash file must reside in the same directory as the key database file, have the same file name as the key database file, and have the <tt>.sth</tt> file extension. For example, if the key database file is <tt>MQkeys.pem</tt> or <tt>MQkeys.kidb</tt> , the stash file must be <tt>MQkeys.sth</tt> .</p><p>If these file are not in the <tt>cert:</tt> directory, upload or fetch them.</p>",
							Computed:            true,
						},
						"permit_insecure_servers": schema.BoolAttribute{
							MarkdownDescription: "When the configuration uses the TLS key repository, specify whether to permit connections to IBM MQ servers that do not support RFC 5746. Such servers are vulnerable to MITM attacks as documented in CVE-2009-3555. By default, insecure connections are rejected during the handshake.",
							Computed:            true,
						},
						"ssl_cipher": schema.StringAttribute{
							MarkdownDescription: "<p>When the configuration uses a TLS key repository, specify the TLS cipher suite. The cipher suite must match the ciphers that the remote queue manager uses.</p><p>Use this setting with the TLS key repository setting to enable TLS communication when the TLS artifacts were created with IBM Global Security Kit (GSKit).</p><p><b>Note:</b> To integrate with IBM MQ for z/OS, do not use these settings. Use a TLS client profile.</p>",
							Computed:            true,
						},
						"ssl_cert_label": schema.StringAttribute{
							MarkdownDescription: "TLS certificate label",
							Computed:            true,
						},
						"convert_input": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to convert input messages to a different CCSI (coded character set identifier) than the one in the incoming message. This conversion is done by the queue manager on the IBM MQ server.",
							Computed:            true,
						},
						"auto_retry": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to reconnect to the server after a critical connection error. By default, automatically attempts to reconnect to the IBM MQ server. This setting does not affect established connections.",
							Computed:            true,
						},
						"retry_interval": schema.Int64Attribute{
							MarkdownDescription: "Specify the interval in seconds between failed connection attempts to the IBM MQ server. The default value is 10. This setting does not affect established connections.",
							Computed:            true,
						},
						"retry_attempts": schema.Int64Attribute{
							MarkdownDescription: "Specify the number of failed connection attempts. After the number is reached, the long interval is used. The default value is 6. When 0, the long retry interval is not used. The retry interval is used forever.",
							Computed:            true,
						},
						"long_retry_interval": schema.Int64Attribute{
							MarkdownDescription: "Specify the long retry interval in seconds for failed connections. The long retry interval is used after the number of retry attempts is reached. The default value is 600. The long retry interval must be greater than the retry interval.",
							Computed:            true,
						},
						"reporting_interval": schema.Int64Attribute{
							MarkdownDescription: "Specify the interval in seconds between error messages for failed connection attempts. This setting filters the generation of identical error messages to IBM MQ logging targets. The default value is 10.",
							Computed:            true,
						},
						"alternate_user": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to use the ALTERNATE_USER_AUTHORITY flag. The valdue determines whether to use <tt>MQOD.AlternateUserId</tt> as the value of the username setting. <ul><li>When enabled, uses <tt>MQOD.AlternateUserId</tt> .</li><li>When disabled, uses <tt>MQMD.UserIdentifier</tt> .</li></ul>",
							Computed:            true,
						},
						"local_address": schema.StringAttribute{
							MarkdownDescription: "Specify the local address for the channel, which is the local address for outbound connections to a specific local interface and port. Supported formats are 1.1.1.1 or 1.1.1.1(1) or just (1) to tell TCP to bind to port 1 and for a range of ports use (1,10) or 1.1.1.1(1,10). If the port is set, the range must be greater than the total number of allowed connections.",
							Computed:            true,
						},
						"xml_manager": schema.StringAttribute{
							MarkdownDescription: "XML manager",
							Computed:            true,
						},
						"ssl_client": schema.StringAttribute{
							MarkdownDescription: "TLS client profile",
							Computed:            true,
						},
						"outbound_sni": schema.StringAttribute{
							MarkdownDescription: "Specify the Outbound SNI settings.",
							Computed:            true,
						},
						"ocsp_check_extensions": schema.BoolAttribute{
							MarkdownDescription: "Check OCSP extensions",
							Computed:            true,
						},
						"ocsp_authentication": schema.StringAttribute{
							MarkdownDescription: "OCSP authentication",
							Computed:            true,
						},
						"cdp_check_extensions": schema.BoolAttribute{
							MarkdownDescription: "Check CDP extensions",
							Computed:            true,
						},
						"client_revocation_checks": schema.StringAttribute{
							MarkdownDescription: "Client revocation checking",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *MQManagerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *MQManagerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data MQManagerList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.MQManager{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.MQManager{}
	if value := res.Get(`MQManager`); value.Exists() {
		for _, v := range value.Array() {
			item := models.MQManager{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.MQManagerObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.MQManagerObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
