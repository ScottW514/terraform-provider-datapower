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
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &ErrorReportSettingsResource{}

func NewErrorReportSettingsResource() resource.Resource {
	return &ErrorReportSettingsResource{}
}

type ErrorReportSettingsResource struct {
	client *client.DatapowerClient
}

func (r *ErrorReportSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_errorreportsettings"
}

func (r *ErrorReportSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Failure Notification (`default` domain only)", "failure-notification", "").String,
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Administrative state", "admin-state", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"upload_report": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Upload Error Report", "upload-report", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"use_smtp": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("E-mail Notification", "use-smtp", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"internal_state": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Include Internal State", "internal-state", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ffdc_packet_capture": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Background Packet Capture", "ffdc packet-capture", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ffdc_event_log_capture": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Background Log Capture", "ffdc event-log", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ffdc_memory_leak_capture": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Background Memory Trace", "ffdc memory-trace", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"always_on_startup": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Always On Startup", "always-on-startup", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"always_on_shutdown": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Always On Shutdown", "always-on-shutdown", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"protocol": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Report Destination Protocol", "protocol", "").AddStringEnum("ftp", "nfs", "raid", "smtp", "temporary", "mqdiag").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("ftp", "nfs", "raid", "smtp", "temporary", "mqdiag"),
				},
			},
			"location_identifier": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Location", "location-id", "").String,
				Optional:            true,
			},
			"smtp_server": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SMTP Server", "remote-address", "").String,
				Optional:            true,
			},
			"email_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("E-mail Address", "email-address", "").String,
				Optional:            true,
			},
			"email_sender_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("E-mail Sender Address", "email-sender-address", "").String,
				Optional:            true,
			},
			"ftp_server": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("FTP Server", "ftp-server", "").String,
				Optional:            true,
			},
			"ftp_path": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("FTP Path", "ftp-path", "").String,
				Optional:            true,
			},
			"ftp_user_agent": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("FTP User Agent", "ftp-user-agent", "httpuseragent").String,
				Optional:            true,
			},
			"nfs_mount": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("NFS Mount", "nfs-mount", "nfsstaticmount").String,
				Optional:            true,
			},
			"nfs_path": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("NFS Path", "nfs-path", "").String,
				Optional:            true,
			},
			"raid_volume": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("RAID Volume", "raid-volume", "raidvolume").String,
				Optional:            true,
			},
			"raid_path": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("RAID Volume Path", "raid-path", "").String,
				Optional:            true,
			},
			"report_history_kept": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Report History", "report-history", "").AddIntegerRange(2, 10).AddDefaultValue("5").String,
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

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *ErrorReportSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.ErrorReportSettings

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, "default", data.DependencyActions, actions.Create)

	body := data.ToBody(ctx, `ErrorReportSettings`)
	_, err := r.client.Put(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "PUT", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ErrorReportSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.ErrorReportSettings

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Get(data.GetPath())
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
	var data models.ErrorReportSettings

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, "default", data.DependencyActions, actions.Update)
	_, err := r.client.Put(data.GetPath(), data.ToBody(ctx, `ErrorReportSettings`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ErrorReportSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	resp.State.RemoveResource(ctx)
}
