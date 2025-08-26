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

type LogTargetWOList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &LogTargetDataSource{}
	_ datasource.DataSourceWithConfigure = &LogTargetDataSource{}
)

func NewLogTargetDataSource() datasource.DataSource {
	return &LogTargetDataSource{}
}

type LogTargetDataSource struct {
	pData *tfutils.ProviderData
}

func (d *LogTargetDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_log_target"
}

func (d *LogTargetDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "<p>Log targets capture messages that are posted by the various objects and services. Target types enable additional capabilities.</p><p>Messages in log targets can be filtered by priority, event code, event category, object type, or IP address. By default, a log target cannot accept messages until it is subscribed to one or more events.</p>",
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
						"backup": schema.StringAttribute{
							MarkdownDescription: "Sets another Log Target object as a backup to receive redirected events in case of an error on the current file-based log target. This setting has no effect on network-based log targets. For network-based log targets, set a load balancer group as the remote host.",
							Computed:            true,
						},
						"log_events": schema.ListNestedAttribute{
							MarkdownDescription: "<p>Subscribes the log target to particular event categories. Some example categories include:</p><dl><dt>auth</dt><dd>Authorization events</dd><dt>mgmt</dt><dd>Configuration management events</dd><dt>xslt</dt><dd>XSLT processing events</dd></dl><p>For each event category chosen (including the <tt>all</tt> category), you can establish a priority level that must be met before the log message will be captured by the log target. Without event subscriptions, no events are included by default. To allow the log target to capture messages, the configuration must include at least one event subscription. The category can be the <tt>all</tt> category.</p>",
							NestedObject:        models.GetDmLogEventDataSourceSchema(),
							Computed:            true,
						},
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Specify the type of the log target. The default value is file.",
							Computed:            true,
						},
						"priority": schema.StringAttribute{
							MarkdownDescription: "Specify the priority to control the scheduling of logs. When system resources are in high demand, high priority operations are favored over lower priority operations.",
							Computed:            true,
						},
						"soap_version": schema.StringAttribute{
							MarkdownDescription: "SOAP version",
							Computed:            true,
						},
						"format": schema.StringAttribute{
							MarkdownDescription: "Log format",
							Computed:            true,
						},
						"timestamp_format": schema.StringAttribute{
							MarkdownDescription: "Specify the format of the timestamp for log entries. The default format is ISO UTC format.",
							Computed:            true,
						},
						"fixed_format": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to make the format of logs unchanging. The log format fixed at version 6.0.1. New fields added to log formats are ignored.",
							Computed:            true,
						},
						"local_identifier": schema.StringAttribute{
							MarkdownDescription: "Specify a descriptive string that identifies the log target to remote recipients. For syslog destinations, do not include spaces.",
							Computed:            true,
						},
						"email_address": schema.StringAttribute{
							MarkdownDescription: "Recipient email address",
							Computed:            true,
						},
						"sender_address": schema.StringAttribute{
							MarkdownDescription: "Specify the email address of the sender. The value must match the email address of the crypto key when email messages are signed.",
							Computed:            true,
						},
						"smtp_domain": schema.StringAttribute{
							MarkdownDescription: "Specify the fully-qualified domain name of the SMTP client. This information is part of the SMTP session initiation (HELO command).",
							Computed:            true,
						},
						"size": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum size of file-based log targets. Enter a value in the range 100 - 50000. The default value is 500.",
							Computed:            true,
						},
						"url": schema.StringAttribute{
							MarkdownDescription: "Specify the HTTP URL to send log entries. Entries are sent with the POST method and uses the default user agent.",
							Computed:            true,
						},
						"nf_sm_ount": schema.StringAttribute{
							MarkdownDescription: "NFS static mount",
							Computed:            true,
						},
						"local_file": schema.StringAttribute{
							MarkdownDescription: "Specify the name of the log file. For example, <tt>logtemp:///filename.log</tt> or <tt>logstore:///filename.log</tt> .",
							Computed:            true,
						},
						"nfs_file": schema.StringAttribute{
							MarkdownDescription: "Specify the path to the log file. The path is relative to the NFS mount. Use a regular expression in the <tt>^[_a-z0-9A-Z/][-_a-z0-9A-Z/.]*$</tt> format. Do not end the path with a forward slash (/).",
							Computed:            true,
						},
						"archive_mode": schema.StringAttribute{
							MarkdownDescription: "Archive mode",
							Computed:            true,
						},
						"upload_method": schema.StringAttribute{
							MarkdownDescription: "Upload protocol",
							Computed:            true,
						},
						"rotate": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum number of rotations. Enter a value in the range 1 - 100. The default value is 3.",
							Computed:            true,
						},
						"use_ansi_color": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enable the use of ANSI color scheme. When enabled, ANSI X3.64 escape sequences color-code messages by log level.",
							Computed:            true,
						},
						"remote_address": schema.StringAttribute{
							MarkdownDescription: "Specify the host name or IP address of the remote server. To establish a secure TLS connection to the server, set this value to the value of the remote host of a TLS proxy service. The local TLS proxy service then securely forwards the log entries to its configured remote server.",
							Computed:            true,
						},
						"remote_port": schema.Int64Attribute{
							MarkdownDescription: "Specify the listening port on the remote server. If using a local TLS proxy service to establish a secure TLS connection, set this value to the value of the remote port of the TLS proxy service.",
							Computed:            true,
						},
						"remote_login": schema.StringAttribute{
							MarkdownDescription: "Remote login",
							Computed:            true,
						},
						"remote_directory": schema.StringAttribute{
							MarkdownDescription: "Specify an existing writable directory on the remote server to upload files. <ul><li>To denote an absolute directory from the root directory, specify a single forward slash character (/) or equivalent encoded character (%2F) before the fully qualified path. <ul><li>For SCP or SFTP, enter / to resolve to //.</li><li>For FTP, enter %2F to resolve to /%2F.</li></ul></li><li>To denote a directory that is relative to the home directory of a user, do not specify a forward slash character or encoded character before the fully qualified file name.</li></ul>",
							Computed:            true,
						},
						"local_address": schema.StringAttribute{
							MarkdownDescription: "Local address",
							Computed:            true,
						},
						"syslog_facility": schema.StringAttribute{
							MarkdownDescription: "Specify the syslog log facility (per RFC 3164) to include in messages sent to the syslog log target.",
							Computed:            true,
						},
						"rate_limit": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum number of events to log per second. Enter a value in the range 1 - 1000. The default value is 100. <ul><li>Remote log targets might receive more than this number of events within a second, depending on network latency and buffering. syslog over TCP log targets are exclusive, because only a single TCP connection is made to the server.</li><li>In the case of syslog over TCP log targets, the rate limit is the maximum number of events transmitted over the connection within one second. A value of 0 disables rate-limiting by the logging target.</li></ul>",
							Computed:            true,
						},
						"max_connections": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum number of concurrent connections that can be opened to the syslog-tcp server. Enter a value in the range 1 - 100. The default value is 1.",
							Computed:            true,
						},
						"connect_timeout": schema.Int64Attribute{
							MarkdownDescription: "Specify the time to wait in seconds for a connection to the server to be established before generating an error. At this time, a log message is generated in the default log and connection retry attempts are made. Enter a value in the range 1 - 90. The default value is 60.",
							Computed:            true,
						},
						"idle_timeout": schema.Int64Attribute{
							MarkdownDescription: "Specify the time in seconds to wait before closing an established but inactive connection to the server. Enter a value in the range 1 - 600. The default value is 15. <p><b>Attention:</b> If multiple log targets have the following configuration, they might share connections. <ul><li>The same local address and port</li><li>The same remote address and port</li></ul> Because of potential connection-sharing, set the same idle timeout value for these log targets.</p>",
							Computed:            true,
						},
						"active_timeout": schema.Int64Attribute{
							MarkdownDescription: "Specify the time to wait in seconds before closing an established and active connection to the server. Enter a value in the range 0 - 60. A value of 0 allows the log target to most efficiently send messages to the server by maintaining a healthy connection indefinitely. The default value is 0. <p><b>Attention:</b> If multiple log targets have the following configuration, they might share connections. <ul><li>The same local address and port</li><li>The same remote address and port</li></ul> Because of potential connection-sharing, set the same active timeout value for these log targets.</p>",
							Computed:            true,
						},
						"feedback_detection": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to suppress events from the logging subsystem itself. A log target always suppresses its own events, but will record events from other log targets. Under certain circumstances with multiple log targets, these events could create a positive feedback loop that could cause resource contention. Enable to suppress all log events from the logging subsystem and prevent resource contention.",
							Computed:            true,
						},
						"log_event_code": schema.ListAttribute{
							MarkdownDescription: "Specify specific events to allow in the log. Subscription filters allow only those log messages that contain the configured event codes. With this filter, it is possible to create a log target that collects only log messages for a specific set of event codes.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"log_event_filter": schema.ListAttribute{
							MarkdownDescription: "Specify specific events to suppress in the log. Suppression filters suppress those log messages that contain the configured event codes. With this filter, it is possible to create a log target that collects a wide range of log messages except for a specific set of event codes.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"log_objects": schema.ListNestedAttribute{
							MarkdownDescription: "Specify specific objects to log events for. Object filters allow only those log messages for specific objects to be written to this log target. Object filters are based on object classes. With this filter, you can create a log target that collects only log messages generated by particular instances of the specified object classes.",
							NestedObject:        models.GetDmLogObjectDataSourceSchema(),
							Computed:            true,
						},
						"log_ip_filter": schema.ListNestedAttribute{
							MarkdownDescription: "Specify specific IP addresses to log events for. IP address filters allow only those log messages from specific IP addresses to be written to this log target.",
							NestedObject:        models.GetDmLogIPFilterDataSourceSchema(),
							Computed:            true,
						},
						"log_triggers": schema.ListNestedAttribute{
							MarkdownDescription: "Specify event trigger points. Event triggers start actions only when triggered by a specified message ID or event code. With this filter, it is possible to create a log target that collects only the results of the specified trigger action. For example, to trigger the generation of an error report when a certain event occurs use the <b>save error-report</b> command.",
							NestedObject:        models.GetDmLogTriggerDataSourceSchema(),
							Computed:            true,
						},
						"ssl_client_profile": schema.StringAttribute{
							MarkdownDescription: "TLS client profile",
							Computed:            true,
						},
						"ssl_client_config_type": schema.StringAttribute{
							MarkdownDescription: "TLS client type",
							Computed:            true,
						},
						"retry_interval": schema.Int64Attribute{
							MarkdownDescription: "Specify the time to wait in seconds before attempting to reestablish a failed connection to the syslog server. Enter a value in the range 1 - 600. The default value is 1.",
							Computed:            true,
						},
						"retry_attempts": schema.Int64Attribute{
							MarkdownDescription: "Specify the number of attempts for a failed connection to the syslog server. After the number of attempts is reached, connection attempts use the value set for the long retry interval. When the long interval is disabled, the log target repeatedly attempts to reconnect to the syslog server with the value set for the retry interval. <p><b>Note:</b> 0 means that the long retry interval is never used and retries forever by using the retry interval.</p>",
							Computed:            true,
						},
						"long_retry_interval": schema.Int64Attribute{
							MarkdownDescription: "Specify the time to wait in seconds before attempting to reestablish a failed connection to the syslog server after the number of attempts is reached. Enter a value in the range 0 - 600. The default value is 20. <p><b>Note:</b> The long retry interval must be greater than the retry interval or it will take no effect.</p>",
							Computed:            true,
						},
						"log_precision": schema.StringAttribute{
							MarkdownDescription: "Specify the precision for the timestamp of log messages. The default value is seconds.",
							Computed:            true,
						},
						"event_buffer_size": schema.StringAttribute{
							MarkdownDescription: "Specify the buffer size in number of event entries. The buffer stores log events before they are written to the target. A buffer of this size is allocated for each connection.",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *LogTargetDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *LogTargetDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data LogTargetWOList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.LogTargetWO{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.LogTargetWO{}
	if value := res.Get(`LogTarget`); value.Exists() {
		for _, v := range value.Array() {
			item := models.LogTargetWO{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.LogTargetObjectTypeWO}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.LogTargetObjectTypeWO})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
