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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

var _ resource.Resource = &LogTargetResource{}
var _ resource.ResourceWithImportState = &LogTargetResource{}

func NewLogTargetResource() resource.Resource {
	return &LogTargetResource{}
}

type LogTargetResource struct {
	pData *tfutils.ProviderData
}

func (r *LogTargetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_log_target"
}

func (r *LogTargetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("<p>Log targets capture messages that are posted by the various objects and services. Target types enable additional capabilities.</p><p>Messages in log targets can be filtered by priority, event code, event category, object type, or IP address. By default, a log target cannot accept messages until it is subscribed to one or more events.</p>", "logging target", "").String,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name of the object. Must be unique among object types in application domain.", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
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
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
				PlanModifiers: []planmodifier.String{
					modifiers.ImmutableAfterSet(),
				},
			},
			"backup": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Sets another Log Target object as a backup to receive redirected events in case of an error on the current file-based log target. This setting has no effect on network-based log targets. For network-based log targets, set a load balancer group as the remote host.", "backup", "log_target").AddNotValidWhen(models.LogTargetBackupIgnoreVal.String()).String,
				Optional:            true,
			},
			"log_events": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Subscribes the log target to particular event categories. Some example categories include:</p><dl><dt>auth</dt><dd>Authorization events</dd><dt>mgmt</dt><dd>Configuration management events</dd><dt>xslt</dt><dd>XSLT processing events</dd></dl><p>For each event category chosen (including the <tt>all</tt> category), you can establish a priority level that must be met before the log message will be captured by the log target. Without event subscriptions, no events are included by default. To allow the log target to capture messages, the configuration must include at least one event subscription. The category can be the <tt>all</tt> category.</p>", "event", "").String,
				NestedObject:        models.GetDmLogEventResourceSchema(),
				Optional:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the type of the log target. The default value is file.", "type", "").AddStringEnum("console", "cache", "syslog", "syslog-tcp", "smtp", "file", "soap", "snmp", "nfs").AddDefaultValue("file").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("console", "cache", "syslog", "syslog-tcp", "smtp", "file", "soap", "snmp", "nfs"),
				},
				Default: stringdefault.StaticString("file"),
			},
			"priority": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the priority to control the scheduling of logs. When system resources are in high demand, high priority operations are favored over lower priority operations.", "priority", "").AddStringEnum("unknown", "high-min", "high", "high-max", "normal-min", "normal", "normal-max", "low-min", "low", "low-max").AddDefaultValue("normal").AddNotValidWhen(models.LogTargetPriorityIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("unknown", "high-min", "high", "high-max", "normal-min", "normal", "normal-max", "low-min", "low", "low-max"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.LogTargetPriorityIgnoreVal, true),
				},
				Default: stringdefault.StaticString("normal"),
			},
			"soap_version": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SOAP version", "soap-version", "").AddStringEnum("soap11", "soap12").AddDefaultValue("soap11").AddNotValidWhen(models.LogTargetSoapVersionIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("soap11", "soap12"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.LogTargetSoapVersionIgnoreVal, true),
				},
				Default: stringdefault.StaticString("soap11"),
			},
			"format": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Log format", "format", "").AddStringEnum("text", "raw", "xml", "json-icp", "cbe", "csv", "audit", "diag").AddDefaultValue("xml").AddNotValidWhen(models.LogTargetFormatIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("text", "raw", "xml", "json-icp", "cbe", "csv", "audit", "diag"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.LogTargetFormatIgnoreVal, true),
				},
				Default: stringdefault.StaticString("xml"),
			},
			"timestamp_format": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the format of the timestamp for log entries. The default format is ISO UTC format.", "timestamp", "").AddStringEnum("syslog", "numeric", "zulu").AddDefaultValue("zulu").AddNotValidWhen(models.LogTargetTimestampFormatIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("syslog", "numeric", "zulu"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.LogTargetTimestampFormatIgnoreVal, true),
				},
				Default: stringdefault.StaticString("zulu"),
			},
			"fixed_format": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to make the format of logs unchanging. The log format fixed at version 6.0.1. New fields added to log formats are ignored.", "fixed-format", "").AddDefaultValue("false").AddNotValidWhen(models.LogTargetFixedFormatIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"local_identifier": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify a descriptive string that identifies the log target to remote recipients. For syslog destinations, do not include spaces.", "local-ident", "").AddRequiredWhen(models.LogTargetLocalIdentifierCondVal.String()).AddNotValidWhen(models.LogTargetLocalIdentifierIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.LogTargetLocalIdentifierCondVal, models.LogTargetLocalIdentifierIgnoreVal, false),
				},
			},
			"email_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Recipient email address", "email-address", "").AddRequiredWhen(models.LogTargetEmailAddressCondVal.String()).AddNotValidWhen(models.LogTargetEmailAddressIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.LogTargetEmailAddressCondVal, models.LogTargetEmailAddressIgnoreVal, false),
				},
			},
			"sender_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the email address of the sender. The value must match the email address of the crypto key when email messages are signed.", "sender-address", "").AddRequiredWhen(models.LogTargetSenderAddressCondVal.String()).AddNotValidWhen(models.LogTargetSenderAddressIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.LogTargetSenderAddressCondVal, models.LogTargetSenderAddressIgnoreVal, false),
				},
			},
			"smtp_domain": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the fully-qualified domain name of the SMTP client. This information is part of the SMTP session initiation (HELO command).", "smtp-domain", "").AddRequiredWhen(models.LogTargetSMTPDomainCondVal.String()).AddNotValidWhen(models.LogTargetSMTPDomainIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.LogTargetSMTPDomainCondVal, models.LogTargetSMTPDomainIgnoreVal, false),
				},
			},
			"size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum size of file-based log targets. Enter a value in the range 100 - 50000. The default value is 500.", "size", "").AddIntegerRange(100, 50000).AddDefaultValue("500").AddRequiredWhen(models.LogTargetSizeCondVal.String()).AddNotValidWhen(models.LogTargetSizeIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(100, 50000),
					validators.ConditionalRequiredInt64(models.LogTargetSizeCondVal, models.LogTargetSizeIgnoreVal, true),
				},
				Default: int64default.StaticInt64(500),
			},
			"url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the HTTP URL to send log entries. Entries are sent with the POST method and uses the default user agent.", "url", "").AddRequiredWhen(models.LogTargetURLCondVal.String()).AddNotValidWhen(models.LogTargetURLIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.LogTargetURLCondVal, models.LogTargetURLIgnoreVal, false),
				},
			},
			"nf_sm_ount": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("NFS static mount", "nfs-static-mount", "nfs_static_mount").AddRequiredWhen(models.LogTargetNFSMountCondVal.String()).AddNotValidWhen(models.LogTargetNFSMountIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.LogTargetNFSMountCondVal, models.LogTargetNFSMountIgnoreVal, false),
				},
			},
			"local_file": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the log file. For example, <tt>logtemp:///filename.log</tt> or <tt>logstore:///filename.log</tt> .", "local-file", "").AddRequiredWhen(models.LogTargetLocalFileCondVal.String()).AddNotValidWhen(models.LogTargetLocalFileIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile("^(logtemp|logstore):[_a-z0-9A-Z/][-_a-z0-9A-Z/.]*$"), "Must match :"+"^(logtemp|logstore):[_a-z0-9A-Z/][-_a-z0-9A-Z/.]*$"),
					validators.ConditionalRequiredString(models.LogTargetLocalFileCondVal, models.LogTargetLocalFileIgnoreVal, false),
				},
			},
			"nfs_file": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the path to the log file. The path is relative to the NFS mount. Use a regular expression in the <tt>^[_a-z0-9A-Z/][-_a-z0-9A-Z/.]*$</tt> format. Do not end the path with a forward slash (/).", "nfs-file", "").AddRequiredWhen(models.LogTargetNFSFileCondVal.String()).AddNotValidWhen(models.LogTargetNFSFileIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile("^[_a-z0-9A-Z/][-_a-z0-9A-Z/.]*$"), "Must match :"+"^[_a-z0-9A-Z/][-_a-z0-9A-Z/.]*$"),
					validators.ConditionalRequiredString(models.LogTargetNFSFileCondVal, models.LogTargetNFSFileIgnoreVal, false),
				},
			},
			"archive_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Archive mode", "archive-mode", "").AddStringEnum("rotate", "upload").AddDefaultValue("rotate").AddRequiredWhen(models.LogTargetArchiveModeCondVal.String()).AddNotValidWhen(models.LogTargetArchiveModeIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("rotate", "upload"),
					validators.ConditionalRequiredString(models.LogTargetArchiveModeCondVal, models.LogTargetArchiveModeIgnoreVal, true),
				},
				Default: stringdefault.StaticString("rotate"),
			},
			"upload_method": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Upload protocol", "upload-method", "").AddStringEnum("ftp", "scp", "sftp", "smtp").AddDefaultValue("ftp").AddRequiredWhen(models.LogTargetUploadMethodCondVal.String()).AddNotValidWhen(models.LogTargetUploadMethodIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("ftp", "scp", "sftp", "smtp"),
					validators.ConditionalRequiredString(models.LogTargetUploadMethodCondVal, models.LogTargetUploadMethodIgnoreVal, true),
				},
				Default: stringdefault.StaticString("ftp"),
			},
			"rotate": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum number of rotations. Enter a value in the range 1 - 100. The default value is 3.", "rotate", "").AddIntegerRange(1, 100).AddDefaultValue("3").AddRequiredWhen(models.LogTargetRotateCondVal.String()).AddNotValidWhen(models.LogTargetRotateIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 100),
					validators.ConditionalRequiredInt64(models.LogTargetRotateCondVal, models.LogTargetRotateIgnoreVal, true),
				},
				Default: int64default.StaticInt64(3),
			},
			"use_ansi_color": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable the use of ANSI color scheme. When enabled, ANSI X3.64 escape sequences color-code messages by log level.", "ansi-color", "").AddDefaultValue("false").AddNotValidWhen(models.LogTargetUseANSIColorIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"remote_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the host name or IP address of the remote server. To establish a secure TLS connection to the server, set this value to the value of the remote host of a TLS proxy service. The local TLS proxy service then securely forwards the log entries to its configured remote server.", "remote-address", "").AddRequiredWhen(models.LogTargetRemoteAddressCondVal.String()).AddNotValidWhen(models.LogTargetRemoteAddressIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.LogTargetRemoteAddressCondVal, models.LogTargetRemoteAddressIgnoreVal, false),
				},
			},
			"remote_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the listening port on the remote server. If using a local TLS proxy service to establish a secure TLS connection, set this value to the value of the remote port of the TLS proxy service.", "remote-port", "").AddIntegerRange(1, 65535).AddNotValidWhen(models.LogTargetRemotePortIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, models.LogTargetRemotePortIgnoreVal, false),
				},
			},
			"remote_login": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Remote login", "remote-login", "").AddRequiredWhen(models.LogTargetRemoteLoginCondVal.String()).AddNotValidWhen(models.LogTargetRemoteLoginIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.LogTargetRemoteLoginCondVal, models.LogTargetRemoteLoginIgnoreVal, false),
				},
			},
			"remote_password_wo": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the password for the account or username for non-public key authentication. Public key authentication can be configured through the default user agent.", "", "").AddRequiredWhen(models.LogTargetRemotePasswordCondVal.String()).AddNotValidWhen(models.LogTargetRemotePasswordIgnoreVal.String()).String,
				Optional:            true,
				WriteOnly:           true,
				Sensitive:           true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.LogTargetRemotePasswordCondVal, models.LogTargetRemotePasswordIgnoreVal, false),
				},
			},
			"remote_password_wo_version": schema.Int64Attribute{
				MarkdownDescription: "Changes to this value trigger an update to `write_only` value.",
				Optional:            true,
				Validators: []validator.Int64{
					validators.ConditionalRequiredInt64(
						validators.Evaluation{
							Evaluation:  "property-value-not-in-list",
							Attribute:   "remote_password_wo",
							AttrType:    "String",
							AttrDefault: "",
							Value:       []string{""},
						}, validators.Evaluation{}, false),
				},
			},
			"remote_directory": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify an existing writable directory on the remote server to upload files. <ul><li>To denote an absolute directory from the root directory, specify a single forward slash character (/) or equivalent encoded character (%2F) before the fully qualified path. <ul><li>For SCP or SFTP, enter / to resolve to //.</li><li>For FTP, enter %2F to resolve to /%2F.</li></ul></li><li>To denote a directory that is relative to the home directory of a user, do not specify a forward slash character or encoded character before the fully qualified file name.</li></ul>", "remote-directory", "").AddNotValidWhen(models.LogTargetRemoteDirectoryIgnoreVal.String()).String,
				Optional:            true,
			},
			"local_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Local address", "local-address", "").AddRequiredWhen(models.LogTargetLocalAddressCondVal.String()).AddNotValidWhen(models.LogTargetLocalAddressIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.LogTargetLocalAddressCondVal, models.LogTargetLocalAddressIgnoreVal, false),
				},
			},
			"syslog_facility": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the syslog log facility (per RFC 3164) to include in messages sent to the syslog log target.", "facility", "").AddStringEnum("user", "security", "authpriv", "local0", "local1", "local2", "local3", "local4", "local5", "local6", "local7").AddDefaultValue("user").AddRequiredWhen(models.LogTargetSyslogFacilityCondVal.String()).AddNotValidWhen(models.LogTargetSyslogFacilityIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("user", "security", "authpriv", "local0", "local1", "local2", "local3", "local4", "local5", "local6", "local7"),
					validators.ConditionalRequiredString(models.LogTargetSyslogFacilityCondVal, models.LogTargetSyslogFacilityIgnoreVal, true),
				},
				Default: stringdefault.StaticString("user"),
			},
			"rate_limit": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum number of events to log per second. Enter a value in the range 1 - 1000. The default value is 100. <ul><li>Remote log targets might receive more than this number of events within a second, depending on network latency and buffering. syslog over TCP log targets are exclusive, because only a single TCP connection is made to the server.</li><li>In the case of syslog over TCP log targets, the rate limit is the maximum number of events transmitted over the connection within one second. A value of 0 disables rate-limiting by the logging target.</li></ul>", "rate-limit", "").AddIntegerRange(0, 1000).AddDefaultValue("100").AddNotValidWhen(models.LogTargetRateLimitIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 1000),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, models.LogTargetRateLimitIgnoreVal, true),
				},
				Default: int64default.StaticInt64(100),
			},
			"max_connections": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum number of concurrent connections that can be opened to the syslog-tcp server. Enter a value in the range 1 - 100. The default value is 1.", "maximum-connections", "").AddIntegerRange(1, 100).AddDefaultValue("1").AddNotValidWhen(models.LogTargetMaxConnectionsIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 100),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, models.LogTargetMaxConnectionsIgnoreVal, true),
				},
				Default: int64default.StaticInt64(1),
			},
			"connect_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the time to wait in seconds for a connection to the server to be established before generating an error. At this time, a log message is generated in the default log and connection retry attempts are made. Enter a value in the range 1 - 90. The default value is 60.", "connect-timeout", "").AddIntegerRange(1, 90).AddDefaultValue("60").AddNotValidWhen(models.LogTargetConnectTimeoutIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 90),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, models.LogTargetConnectTimeoutIgnoreVal, true),
				},
				Default: int64default.StaticInt64(60),
			},
			"idle_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the time in seconds to wait before closing an established but inactive connection to the server. Enter a value in the range 1 - 600. The default value is 15. <p><b>Attention:</b> If multiple log targets have the following configuration, they might share connections. <ul><li>The same local address and port</li><li>The same remote address and port</li></ul> Because of potential connection-sharing, set the same idle timeout value for these log targets.</p>", "idle-timeout", "").AddIntegerRange(1, 600).AddDefaultValue("15").AddNotValidWhen(models.LogTargetIdleTimeoutIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 600),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, models.LogTargetIdleTimeoutIgnoreVal, true),
				},
				Default: int64default.StaticInt64(15),
			},
			"active_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the time to wait in seconds before closing an established and active connection to the server. Enter a value in the range 0 - 60. A value of 0 allows the log target to most efficiently send messages to the server by maintaining a healthy connection indefinitely. The default value is 0. <p><b>Attention:</b> If multiple log targets have the following configuration, they might share connections. <ul><li>The same local address and port</li><li>The same remote address and port</li></ul> Because of potential connection-sharing, set the same active timeout value for these log targets.</p>", "active-timeout", "").AddIntegerRange(0, 60).AddDefaultValue("0").AddNotValidWhen(models.LogTargetActiveTimeoutIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 60),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, models.LogTargetActiveTimeoutIgnoreVal, true),
				},
				Default: int64default.StaticInt64(0),
			},
			"feedback_detection": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to suppress events from the logging subsystem itself. A log target always suppresses its own events, but will record events from other log targets. Under certain circumstances with multiple log targets, these events could create a positive feedback loop that could cause resource contention. Enable to suppress all log events from the logging subsystem and prevent resource contention.", "feedback-detection", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"log_event_code": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify specific events to allow in the log. Subscription filters allow only those log messages that contain the configured event codes. With this filter, it is possible to create a log target that collects only log messages for a specific set of event codes.", "event-code", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"log_event_filter": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify specific events to suppress in the log. Suppression filters suppress those log messages that contain the configured event codes. With this filter, it is possible to create a log target that collects a wide range of log messages except for a specific set of event codes.", "event-filter", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"log_objects": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify specific objects to log events for. Object filters allow only those log messages for specific objects to be written to this log target. Object filters are based on object classes. With this filter, you can create a log target that collects only log messages generated by particular instances of the specified object classes.", "object", "").String,
				NestedObject:        models.GetDmLogObjectResourceSchema(),
				Optional:            true,
			},
			"log_ip_filter": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify specific IP addresses to log events for. IP address filters allow only those log messages from specific IP addresses to be written to this log target.", "ip-filter", "").String,
				NestedObject:        models.GetDmLogIPFilterResourceSchema(),
				Optional:            true,
			},
			"log_triggers": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify event trigger points. Event triggers start actions only when triggered by a specified message ID or event code. With this filter, it is possible to create a log target that collects only the results of the specified trigger action. For example, to trigger the generation of an error report when a certain event occurs use the <b>save error-report</b> command.", "trigger", "").String,
				NestedObject:        models.GetDmLogTriggerResourceSchema(),
				Optional:            true,
			},
			"ssl_client_profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "ssl_client_profile").AddNotValidWhen(models.LogTargetSSLClientProfileIgnoreVal.String()).String,
				Optional:            true,
			},
			"ssl_client_config_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client type", "ssl-client-type", "").AddStringEnum("client").AddDefaultValue("client").AddNotValidWhen(models.LogTargetSSLClientConfigTypeIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("client"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.LogTargetSSLClientConfigTypeIgnoreVal, true),
				},
				Default: stringdefault.StaticString("client"),
			},
			"retry_interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the time to wait in seconds before attempting to reestablish a failed connection to the syslog server. Enter a value in the range 1 - 600. The default value is 1.", "retry-interval", "").AddIntegerRange(1, 600).AddDefaultValue("1").AddNotValidWhen(models.LogTargetRetryIntervalIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 600),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, models.LogTargetRetryIntervalIgnoreVal, true),
				},
				Default: int64default.StaticInt64(1),
			},
			"retry_attempts": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the number of attempts for a failed connection to the syslog server. After the number of attempts is reached, connection attempts use the value set for the long retry interval. When the long interval is disabled, the log target repeatedly attempts to reconnect to the syslog server with the value set for the retry interval. <p><b>Note:</b> 0 means that the long retry interval is never used and retries forever by using the retry interval.</p>", "retry-attempts", "").AddIntegerRange(1, 100).AddDefaultValue("1").AddNotValidWhen(models.LogTargetRetryAttemptsIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 100),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, models.LogTargetRetryAttemptsIgnoreVal, true),
				},
				Default: int64default.StaticInt64(1),
			},
			"long_retry_interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the time to wait in seconds before attempting to reestablish a failed connection to the syslog server after the number of attempts is reached. Enter a value in the range 0 - 600. The default value is 20. <p><b>Note:</b> The long retry interval must be greater than the retry interval or it will take no effect.</p>", "long-retry-interval", "").AddIntegerRange(1, 600).AddDefaultValue("20").AddNotValidWhen(models.LogTargetLongRetryIntervalIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 600),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, models.LogTargetLongRetryIntervalIgnoreVal, true),
				},
				Default: int64default.StaticInt64(20),
			},
			"log_precision": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the precision for the timestamp of log messages. The default value is seconds.", "precision", "").AddStringEnum("second", "microsecond").AddDefaultValue("second").AddRequiredWhen(models.LogTargetLogPrecisionCondVal.String()).AddNotValidWhen(models.LogTargetLogPrecisionIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("second", "microsecond"),
					validators.ConditionalRequiredString(models.LogTargetLogPrecisionCondVal, models.LogTargetLogPrecisionIgnoreVal, true),
				},
				Default: stringdefault.StaticString("second"),
			},
			"event_buffer_size": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the buffer size in number of event entries. The buffer stores log events before they are written to the target. A buffer of this size is allocated for each connection.", "buffer-size", "").AddStringEnum("2048", "16384", "65536", "131072", "262144", "524288", "1048576", "2097152", "4194304", "8388608").AddDefaultValue("2048").AddNotValidWhen(models.LogTargetEventBufferSizeIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("2048", "16384", "65536", "131072", "262144", "524288", "1048576", "2097152", "4194304", "8388608"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.LogTargetEventBufferSizeIgnoreVal, true),
				},
				Default: stringdefault.StaticString("2048"),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *LogTargetResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *LogTargetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data, config models.LogTarget
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `LogTarget`, &config)
	_, err := r.pData.Client.Post(data.GetPath(), body)
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), body)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Resource already exists. Failed to update resource, got error: %s", err))
				return
			}
			resp.Diagnostics.AddWarning("Warning", "Resource already exists. Existing resource was updated.")
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create resource, got error: %s", err))
			return
		}
	}
	getRes, getErr := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if getErr != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object after creation (GET), got error: %s", getErr))
		return
	}
	data.UpdateUnknownFromBody(ctx, `LogTarget`, getRes)
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LogTargetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.LogTarget
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

	data.UpdateFromBody(ctx, `LogTarget`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LogTargetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data, config models.LogTarget
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `LogTarget`, &config))
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

func (r *LogTargetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.LogTarget
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			resp.Diagnostics.AddWarning("Resource Conflict", fmt.Sprintf("Resource is no longer tracked by Terraform, but may need to be manually deleted on DataPower host. Got error: %s", err))
		} else if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete resource, got error: %s", err))
			return
		}
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *LogTargetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError("Invalid Import ID", "Expected format: <app_domain>/<id>. Got: "+req.ID)
		return
	}

	appDomain := parts[0]
	id := parts[1]

	if !regexp.MustCompile("^[a-zA-Z0-9_-]+$").MatchString(id) || len(id) < 1 || len(id) > 128 {
		resp.Diagnostics.AddError("Invalid ID", "ID must be 1-128 characters and match pattern ^[a-zA-Z0-9_-]+$")
		return
	}
	if !regexp.MustCompile("^[a-zA-Z0-9_-]+$").MatchString(appDomain) || len(appDomain) < 1 || len(appDomain) > 128 {
		resp.Diagnostics.AddError("Invalid Application Domain", "Application domain must be 1-128 characters and match pattern ^[a-zA-Z0-9_-]+$")
		return
	}

	var data models.LogTarget
	data.AppDomain = types.StringValue(appDomain)
	data.Id = types.StringValue(id)
	res, err := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("Resource was not found, got error: %s", err))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		}
		return
	}

	data.FromBody(ctx, `LogTarget`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LogTargetResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.LogTarget

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
