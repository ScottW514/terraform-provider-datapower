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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var (
	_ datasource.DataSource              = &ErrorReportSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &ErrorReportSettingsDataSource{}
)

func NewErrorReportSettingsDataSource() datasource.DataSource {
	return &ErrorReportSettingsDataSource{}
}

type ErrorReportSettingsDataSource struct {
	pData *tfutils.ProviderData
}

func (d *ErrorReportSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_errorreportsettings"
}

func (d *ErrorReportSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "<p>Failure notification is a serviceability tool. By default, failure notification is disabled. To use failure notification, you must enable the configuration and allow the error report to be uploaded.</p><p>The uploading of an error report allows the capture of more diagnostic information, which safely improves serviceability. Although there is a tradeoff between performance and serviceability, you should choose serviceability by enabling the following properties:</p><ul><li>Include Internal State</li><li>Background Packet Capture</li><li>Background Log Capture</li><li>Background Memory Trace</li></ul><p>When you allow the error report to be uploaded, this setting enables the Failure Notification status provider. This status provide in combination with the report history tracks the error reports that the appliance generates, the reason why the appliance generated the error report, and its upload status to the specific destination.</p><p>You can specify an NFS, RAID, SMTP, or FTP destination. You can also specify the local temporary directory as the destination. The appliance generates error reports, the naming convention includes the serial number of the appliance and the timestamp of the report. This naming convention prevents one report from overwriting another.</p><p>You can use Event Triggers to generate error reports automatically when specific events occur.</p>",
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "<p>The administrative state of the configuration.</p><ul><li>To make active, set to enabled.</li><li>To make inactive, set to disabled.</li></ul>",
				Computed:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: "Comments",
				Computed:            true,
			},
			"upload_report": schema.BoolAttribute{
				MarkdownDescription: "<p>Specify whether to upload the error report to an NFS, RAID, SMTP, or FTP destination or write the error report to the local temporary directory. If you enable this feature:</p><ul><li>It enables the Failure Notification status provider, which tracks previous error reports</li><li>It changes the naming convention to include the serial number of the appliance and the timestamp, which prevents one report overwriting another</li></ul>",
				Computed:            true,
			},
			"use_smtp": schema.BoolAttribute{
				MarkdownDescription: "Specify whether to send an e-mail at start-up only that contains the error report. If you want to receive e-mail notification, use the upload error report property instead of this property.",
				Computed:            true,
			},
			"internal_state": schema.BoolAttribute{
				MarkdownDescription: "Specify whether to include the internal state of the appliance in the error report. The internal state can be useful in diagnosing the cause of the error.",
				Computed:            true,
			},
			"ffdc_packet_capture": schema.BoolAttribute{
				MarkdownDescription: "<p>Specify whether to use a background packet capture. This feature enables network packet capture for all interfaces including the internal loopback interface. When enabled, this feature runs continuously.</p><p>If the appliance encounters a problem or a user triggers the generation of an error report, the error report includes the data from this packet capture data. This data helps to determine the messages that the appliance was processing when it encountered the problem.</p>",
				Computed:            true,
			},
			"ffdc_event_log_capture": schema.BoolAttribute{
				MarkdownDescription: "<p>Specify whether to use a background log capture. This feature enables the capture of all log and trace points with minimal overhead. When enabled, this feature runs continuously.</p><p>If the appliance encounters a problem or a user triggers the generation of an error report, the error report includes data from this log capture. This data can help IBM Support identify the problem.</p><p>These messages are independent of messages written to log and trace targets.</p>",
				Computed:            true,
			},
			"ffdc_memory_leak_capture": schema.BoolAttribute{
				MarkdownDescription: "<p>Specify whether to enable automatic leak detection. This feature finds gradual memory leaks that occur steadily over time. This feature does not help in situations where messages are larger than the appliance can parse.</p><p>When enabled and if memory falls below an internal threshold, the appliance tracks all memory allocations. When the appliance reaches a critical condition that will lead to a crash, it generates an error report that contains information about memory allocation.</p><p>The configuration of the Throttle Settings affects this feature. The throttle settings can prevent the appliance from reaching the internal threshold.</p>",
				Computed:            true,
			},
			"always_on_startup": schema.BoolAttribute{
				MarkdownDescription: "<p>Specify whether the appliance generates an error report when it reboots or reloads. If this feature is enabled logs will be collected before they are overwritten when the system reloads. When the appliance reboots or reloads and when using the upload error report feature, the status provider lists <tt>on-start</tt> as the reason code. If the appliance reloads due to a crash, the status provider lists <tt>crash</tt> as the reason code.</p><p><b>Best Practices:</b></p><ul><li>In general enable this feature.</li><li>In particular when in production, enable this feature to track reloads and reboots.</li></ul>",
				Computed:            true,
			},
			"always_on_shutdown": schema.BoolAttribute{
				MarkdownDescription: "<p>Specify whether the appliance generates an error report when it shuts down. When the appliance shuts down and when using the upload error report feature, the status provider lists <tt>on-shutdown</tt> as the reason code. If the appliance shuts down due to a crash, the status provider lists <tt>crash</tt> as the reason code.</p><p><b>Best Practices:</b></p><ul><li>In general enable this feature.</li><li>In particular when in production, enable this feature to make sure debugging information is captured when the appliance shuts down.</li></ul>",
				Computed:            true,
			},
			"protocol": schema.StringAttribute{
				MarkdownDescription: "Specify the protocol to use to upload the error report. Note that the selection you will see depends on your device features and licenses.",
				Computed:            true,
			},
			"location_identifier": schema.StringAttribute{
				MarkdownDescription: "Specify text to include in the subject of an e-mail notification. In general, this value should be how you identify this appliance in your environment. When using the upload error report feature, this property is not necessary. In this case, failure notification uses the serial number of the appliance and the timestamp of the error report.",
				Computed:            true,
			},
			"smtp_server": schema.StringAttribute{
				MarkdownDescription: "Specify the host name or IP address of the remote SMTP server to which to send the error report.",
				Computed:            true,
			},
			"email_address": schema.StringAttribute{
				MarkdownDescription: "Specify the e-mail address to which to send the error report.",
				Computed:            true,
			},
			"email_sender_address": schema.StringAttribute{
				MarkdownDescription: "Specify the e-mail address of the sender ( <tt>MAIL FROM</tt> ). If not specified, the configuration uses the e-mail address of the recipient.",
				Computed:            true,
			},
			"ftp_server": schema.StringAttribute{
				MarkdownDescription: "Specify the host name or IP address of the remote FTP server to which to upload the error report.",
				Computed:            true,
			},
			"ftp_path": schema.StringAttribute{
				MarkdownDescription: "Specify the directory on the FTP server to which to upload the error report. Use <tt>%2F</tt> to specify an absolute path.",
				Computed:            true,
			},
			"ftp_user_agent": schema.StringAttribute{
				MarkdownDescription: "Specify the User Agent that describes how to connect to remote FTP servers. In addition to the FTP Policy to define the connection, ensure that this User Agent defines the basic authentication policy (user name and password) to connect to the FTP server.",
				Computed:            true,
			},
			"nfs_mount": schema.StringAttribute{
				MarkdownDescription: "Specify the NFS mount point to which to upload the error report.",
				Computed:            true,
			},
			"nfs_path": schema.StringAttribute{
				MarkdownDescription: "This describes the NFS path location for the Error Report",
				Computed:            true,
			},
			"raid_volume": schema.StringAttribute{
				MarkdownDescription: "Specify the volume on the RAID array to which to upload the error report.",
				Computed:            true,
			},
			"raid_path": schema.StringAttribute{
				MarkdownDescription: "Specify the directory on the RAID volume to which to upload the error report.",
				Computed:            true,
			},
			"report_history_kept": schema.Int64Attribute{
				MarkdownDescription: "<p>Specify the maximum number of local error reports to maintain when using the upload error report feature. After reaching this limit, the next local error report overwrites the oldest local error report. Use any value of 2 - 10. The default value is 5.</p><p>This feature only applies to locally stored error reports, including temporary and Raid.</p><p>To view the history, see the Failure Notification status provider.</p>",
				Computed:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (d *ErrorReportSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *ErrorReportSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data models.ErrorReportSettings
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := d.pData.Client.Get(data.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	data.FromBody(ctx, `ErrorReportSettings`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
