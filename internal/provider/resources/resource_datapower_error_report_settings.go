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
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &ErrorReportSettingsResource{}

func NewErrorReportSettingsResource() resource.Resource {
	return &ErrorReportSettingsResource{}
}

type ErrorReportSettingsResource struct {
	pData *tfutils.ProviderData
}

func (r *ErrorReportSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_error_report_settings"
}

func (r *ErrorReportSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("<p>Failure notification is a serviceability tool. By default, failure notification is disabled. To use failure notification, you must enable the configuration and allow the error report to be uploaded.</p><p>The uploading of an error report allows the capture of more diagnostic information, which safely improves serviceability. Although there is a tradeoff between performance and serviceability, you should choose serviceability by enabling the following properties:</p><ul><li>Include Internal State</li><li>Background Packet Capture</li><li>Background Log Capture</li><li>Background Memory Trace</li></ul><p>When you allow the error report to be uploaded, this setting enables the Failure Notification status provider. This status provide in combination with the report history tracks the error reports that the appliance generates, the reason why the appliance generated the error report, and its upload status to the specific destination.</p><p>You can specify an NFS, RAID, SMTP, or FTP destination. You can also specify the local temporary directory as the destination. The appliance generates error reports, the naming convention includes the serial number of the appliance and the timestamp of the report. This naming convention prevents one report from overwriting another.</p><p>You can use Event Triggers to generate error reports automatically when specific events occur.</p>", "failure-notification", "").String,
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>The administrative state of the configuration.</p><ul><li>To make active, set to enabled.</li><li>To make inactive, set to disabled.</li></ul>", "admin-state", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"upload_report": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to upload the error report to an NFS, RAID, SMTP, or FTP destination or write the error report to the local temporary directory. If you enable this feature:</p><ul><li>It enables the Failure Notification status provider, which tracks previous error reports</li><li>It changes the naming convention to include the serial number of the appliance and the timestamp, which prevents one report overwriting another</li></ul>", "upload-report", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"use_smtp": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to send an e-mail at start-up only that contains the error report. If you want to receive e-mail notification, use the upload error report property instead of this property.", "use-smtp", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"internal_state": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to include the internal state of the appliance in the error report. The internal state can be useful in diagnosing the cause of the error.", "internal-state", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ffdc_packet_capture": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to use a background packet capture. This feature enables network packet capture for all interfaces including the internal loopback interface. When enabled, this feature runs continuously.</p><p>If the appliance encounters a problem or a user triggers the generation of an error report, the error report includes the data from this packet capture data. This data helps to determine the messages that the appliance was processing when it encountered the problem.</p>", "ffdc packet-capture", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ffdc_event_log_capture": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to use a background log capture. This feature enables the capture of all log and trace points with minimal overhead. When enabled, this feature runs continuously.</p><p>If the appliance encounters a problem or a user triggers the generation of an error report, the error report includes data from this log capture. This data can help IBM Support identify the problem.</p><p>These messages are independent of messages written to log and trace targets.</p>", "ffdc event-log", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ffdc_memory_leak_capture": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to enable automatic leak detection. This feature finds gradual memory leaks that occur steadily over time. This feature does not help in situations where messages are larger than the appliance can parse.</p><p>When enabled and if memory falls below an internal threshold, the appliance tracks all memory allocations. When the appliance reaches a critical condition that will lead to a crash, it generates an error report that contains information about memory allocation.</p><p>The configuration of the Throttle Settings affects this feature. The throttle settings can prevent the appliance from reaching the internal threshold.</p>", "ffdc memory-trace", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"always_on_startup": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether the appliance generates an error report when it reboots or reloads. If this feature is enabled logs will be collected before they are overwritten when the system reloads. When the appliance reboots or reloads and when using the upload error report feature, the status provider lists <tt>on-start</tt> as the reason code. If the appliance reloads due to a crash, the status provider lists <tt>crash</tt> as the reason code.</p><p><b>Best Practices:</b></p><ul><li>In general enable this feature.</li><li>In particular when in production, enable this feature to track reloads and reboots.</li></ul>", "always-on-startup", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"always_on_shutdown": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether the appliance generates an error report when it shuts down. When the appliance shuts down and when using the upload error report feature, the status provider lists <tt>on-shutdown</tt> as the reason code. If the appliance shuts down due to a crash, the status provider lists <tt>crash</tt> as the reason code.</p><p><b>Best Practices:</b></p><ul><li>In general enable this feature.</li><li>In particular when in production, enable this feature to make sure debugging information is captured when the appliance shuts down.</li></ul>", "always-on-shutdown", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"protocol": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the protocol to use to upload the error report. Note that the selection you will see depends on your device features and licenses.", "protocol", "").AddStringEnum("ftp", "nfs", "raid", "smtp", "temporary", "mqdiag").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("ftp", "nfs", "raid", "smtp", "temporary", "mqdiag"),
				},
			},
			"location_identifier": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify text to include in the subject of an e-mail notification. In general, this value should be how you identify this appliance in your environment. When using the upload error report feature, this property is not necessary. In this case, failure notification uses the serial number of the appliance and the timestamp of the error report.", "location-id", "").String,
				Optional:            true,
			},
			"smtp_server": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the host name or IP address of the remote SMTP server to which to send the error report.", "remote-address", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.NoneOf([]string{"127.0.0.1", "localhost"}...),
				},
			},
			"email_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the e-mail address to which to send the error report.", "email-address", "").String,
				Optional:            true,
			},
			"email_sender_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the e-mail address of the sender ( <tt>MAIL FROM</tt> ). If not specified, the configuration uses the e-mail address of the recipient.", "email-sender-address", "").String,
				Optional:            true,
			},
			"ftp_server": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the host name or IP address of the remote FTP server to which to upload the error report.", "ftp-server", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 128),
					stringvalidator.NoneOf([]string{"127.0.0.1", "localhost"}...),
				},
			},
			"ftp_path": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the directory on the FTP server to which to upload the error report. Use <tt>%2F</tt> to specify an absolute path.", "ftp-path", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 128),
				},
			},
			"ftp_user_agent": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the User Agent that describes how to connect to remote FTP servers. In addition to the FTP Policy to define the connection, ensure that this User Agent defines the basic authentication policy (user name and password) to connect to the FTP server.", "ftp-user-agent", "http_user_agent").String,
				Optional:            true,
			},
			"nfs_mount": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the NFS mount point to which to upload the error report.", "nfs-mount", "nfs_static_mount").String,
				Optional:            true,
			},
			"nfs_path": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("This describes the NFS path location for the Error Report", "nfs-path", "").String,
				Optional:            true,
			},
			"raid_volume": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the volume on the RAID array to which to upload the error report.", "raid-volume", "raid_volume").String,
				Optional:            true,
			},
			"raid_path": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the directory on the RAID volume to which to upload the error report.", "raid-path", "").String,
				Optional:            true,
			},
			"report_history_kept": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the maximum number of local error reports to maintain when using the upload error report feature. After reaching this limit, the next local error report overwrites the oldest local error report. Use any value of 2 - 10. The default value is 5.</p><p>This feature only applies to locally stored error reports, including temporary and Raid.</p><p>To view the history, see the Failure Notification status provider.</p>", "report-history", "").AddIntegerRange(2, 10).AddDefaultValue("5").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(2, 10),
				},
				Default: int64default.StaticInt64(5),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *ErrorReportSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *ErrorReportSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.ErrorReportSettings
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, "default", data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `ErrorReportSettings`)
	_, err := r.pData.Client.Put(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "PUT", err))
		return
	}
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ErrorReportSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.ErrorReportSettings
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.pData.Client.Get(data.GetPath())
	if err != nil && (strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}

	if data.IsNull() {
		// Import
		data.FromBody(ctx, `ErrorReportSettings`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `ErrorReportSettings`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ErrorReportSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.ErrorReportSettings
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, "default", data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath(), data.ToBody(ctx, `ErrorReportSettings`))
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

func (r *ErrorReportSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.ErrorReportSettings
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, "default", data.DependencyActions, actions.Delete, false)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *ErrorReportSettingsResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.ErrorReportSettings

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
