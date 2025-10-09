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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var (
	_ datasource.DataSource              = &SystemSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &SystemSettingsDataSource{}
)

func NewSystemSettingsDataSource() datasource.DataSource {
	return &SystemSettingsDataSource{}
}

type SystemSettingsDataSource struct {
	pData *tfutils.ProviderData
}

func (d *SystemSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_settings"
}

func (d *SystemSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "<p>System settings provide the following purposes.</p><ul><li>Define system-specific information, such as contact information, location, and name.</li><li>Update serial number after a replacement.</li><li>Enable interface for custom GUI messages and custom CLI prompt.</li><li>Reserve disk space for the audit log.</li><li>Define information about the hardware for use by the SNMP system table, such as serial number, and model type</li></ul>",
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "<p>The administrative state of the configuration.</p><ul><li>To make active, set to enabled.</li><li>To make inactive, set to disabled.</li></ul>",
				Computed:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: "Enter a descriptive summary for the configuration.",
				Computed:            true,
			},
			"product_oid": schema.StringAttribute{
				MarkdownDescription: "The read-only string that identifies the installed DataPower agent software.",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The read-only string that identifies the product.",
				Computed:            true,
			},
			"serial_number": schema.StringAttribute{
				MarkdownDescription: "The read-only string that identifies the product serial number.",
				Computed:            true,
			},
			"entitlement_number": schema.StringAttribute{
				MarkdownDescription: "After an appliance replacement, the serial number of the original appliance. Without the original serial number, IBM cannot entitle the replacement appliance for maintenance or warranty services.",
				Computed:            true,
			},
			"product_id": schema.StringAttribute{
				MarkdownDescription: "The read-only string that identifies the product type.",
				Computed:            true,
			},
			"capacity_mode": schema.StringAttribute{
				MarkdownDescription: "The read-only installation setting that indicates the licensed capacity mode.",
				Computed:            true,
			},
			"contact": schema.StringAttribute{
				MarkdownDescription: "Specify any information that identifies the individual or functional area that is responsible maintenance and management.",
				Computed:            true,
			},
			"system_name": schema.StringAttribute{
				MarkdownDescription: "Enter the name of the DataPower Gateway to use internally as a custom prompt and to use externally to integrate with remote systems. The name must be a 7-bit US-ASCII string of 127 characters or less consisting of letters, numbers, underscore, or embedded dashes, dots, or spaces. However, it is recommended to also be unique with a length of 64 characters or less to be compatible with most remote systems.",
				Computed:            true,
			},
			"location": schema.StringAttribute{
				MarkdownDescription: "Enter the location of the DataPower Gateway.",
				Computed:            true,
			},
			"services": schema.Int64Attribute{
				MarkdownDescription: "The read-only hex value that indicates support for application, presentation, session, and data-link layer services.",
				Computed:            true,
			},
			"backup_mode": schema.StringAttribute{
				MarkdownDescription: "The read-only installation setting that indicates whether a secure-backup is allowed.",
				Computed:            true,
			},
			"product_mode": schema.StringAttribute{
				MarkdownDescription: "The read-only installation setting that indicates the operational mode of the product.",
				Computed:            true,
			},
			"custom_ui_file": schema.StringAttribute{
				MarkdownDescription: "<p>Specifies the URL of the custom user interface file. This file contains custom messages for CLI and GUI sessions as well as the custom CLI prompt. The file must reside in the <tt>store:</tt> or <tt>local:</tt> directory, not on a mounted file system.</p>",
				Computed:            true,
			},
			"audit_reserve": schema.Int64Attribute{
				MarkdownDescription: "Specifies the amount of disk space to reserve for audit records. When the disk is full, all services enter the down operational state and stop processing traffic. To restore disk space and resume traffic processing, manual intervention is required. Enter a value in the range 0 - 10000. The default value is 40.",
				Computed:            true,
			},
			"detect_intrusion": schema.StringAttribute{
				MarkdownDescription: "Indicates whether to check for intrusion detection.",
				Computed:            true,
			},
			"hardware_xml_acceleration": schema.BoolAttribute{
				MarkdownDescription: "<p>Indicates whether to use the Hardware XML accelerator.</p><p><b>Attention: </b>Disable the XML hardware accelerator only when directed by IBM Support. When disabled, XML hardware acceleration rules in the compile options policy have no effect.</p><p>After you change the state, restart the appliance to remove items from the cache. After the appliance restarts, the specified state is in effect. You can view the status of the XML hardware accelerator in the appliance version information. The XML accelerator shows the type appended with <tt>(disabled)</tt> when the accelerator is disabled.</p>",
				Computed:            true,
			},
			"locale": schema.StringAttribute{
				MarkdownDescription: "Specifies the locale for the operating language of the DataPower Gateway. The locale setting manages locale-specific conventions, such as date and time formats, and controls the language of log messages. The language must be enabled before you can select it.",
				Computed:            true,
			},
			"system_log_fixed_format": schema.BoolAttribute{
				MarkdownDescription: "Indicates whether to enable fixed format in system logs. When enabled, the system logs are in the format that was used in version 6.0.1 and contain no serviceability improvements after this version that can help with monitoring or troubleshooting.",
				Computed:            true,
			},
			"uuid": schema.StringAttribute{
				MarkdownDescription: "UUID",
				Computed:            true,
			},
		},
	}
}

func (d *SystemSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *SystemSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data models.SystemSettings
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	path := data.GetPath()

	res, err := d.pData.Client.Get(path)
	resFound := true
	if err != nil {
		if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		} else {
			resFound = false
		}
	}
	if resFound {
		data.FromBody(ctx, `SystemSettings`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
