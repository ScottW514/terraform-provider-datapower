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

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
)

var (
	_ datasource.DataSource              = &ErrorReportSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &ErrorReportSettingsDataSource{}
)

func NewErrorReportSettingsDataSource() datasource.DataSource {
	return &ErrorReportSettingsDataSource{}
}

type ErrorReportSettingsDataSource struct {
	client *client.DatapowerClient
}

func (d *ErrorReportSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_errorreportsettings"
}

func (d *ErrorReportSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Failure Notification (`default` domain only)",
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Administrative state",
				Computed:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: "Comments",
				Computed:            true,
			},
			"upload_report": schema.BoolAttribute{
				MarkdownDescription: "Upload Error Report",
				Computed:            true,
			},
			"use_smtp": schema.BoolAttribute{
				MarkdownDescription: "E-mail Notification",
				Computed:            true,
			},
			"internal_state": schema.BoolAttribute{
				MarkdownDescription: "Include Internal State",
				Computed:            true,
			},
			"ffdc_packet_capture": schema.BoolAttribute{
				MarkdownDescription: "Background Packet Capture",
				Computed:            true,
			},
			"ffdc_event_log_capture": schema.BoolAttribute{
				MarkdownDescription: "Background Log Capture",
				Computed:            true,
			},
			"ffdc_memory_leak_capture": schema.BoolAttribute{
				MarkdownDescription: "Background Memory Trace",
				Computed:            true,
			},
			"always_on_startup": schema.BoolAttribute{
				MarkdownDescription: "Always On Startup",
				Computed:            true,
			},
			"always_on_shutdown": schema.BoolAttribute{
				MarkdownDescription: "Always On Shutdown",
				Computed:            true,
			},
			"protocol": schema.StringAttribute{
				MarkdownDescription: "Report Destination Protocol",
				Computed:            true,
			},
			"location_identifier": schema.StringAttribute{
				MarkdownDescription: "Location",
				Computed:            true,
			},
			"smtp_server": schema.StringAttribute{
				MarkdownDescription: "SMTP Server",
				Computed:            true,
			},
			"email_address": schema.StringAttribute{
				MarkdownDescription: "E-mail Address",
				Computed:            true,
			},
			"email_sender_address": schema.StringAttribute{
				MarkdownDescription: "E-mail Sender Address",
				Computed:            true,
			},
			"ftp_server": schema.StringAttribute{
				MarkdownDescription: "FTP Server",
				Computed:            true,
			},
			"ftp_path": schema.StringAttribute{
				MarkdownDescription: "FTP Path",
				Computed:            true,
			},
			"ftp_user_agent": schema.StringAttribute{
				MarkdownDescription: "FTP User Agent",
				Computed:            true,
			},
			"nfs_mount": schema.StringAttribute{
				MarkdownDescription: "NFS Mount",
				Computed:            true,
			},
			"nfs_path": schema.StringAttribute{
				MarkdownDescription: "NFS Path",
				Computed:            true,
			},
			"raid_volume": schema.StringAttribute{
				MarkdownDescription: "RAID Volume",
				Computed:            true,
			},
			"raid_path": schema.StringAttribute{
				MarkdownDescription: "RAID Volume Path",
				Computed:            true,
			},
			"report_history_kept": schema.Int64Attribute{
				MarkdownDescription: "Report History",
				Computed:            true,
			},
		},
	}
}

func (d *ErrorReportSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *ErrorReportSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.ErrorReportSettings

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := d.client.Get(data.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	data.FromBody(ctx, `ErrorReportSettings`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
