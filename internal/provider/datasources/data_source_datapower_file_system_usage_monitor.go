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
	_ datasource.DataSource              = &FileSystemUsageMonitorDataSource{}
	_ datasource.DataSourceWithConfigure = &FileSystemUsageMonitorDataSource{}
)

func NewFileSystemUsageMonitorDataSource() datasource.DataSource {
	return &FileSystemUsageMonitorDataSource{}
}

type FileSystemUsageMonitorDataSource struct {
	pData *tfutils.ProviderData
}

func (d *FileSystemUsageMonitorDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_file_system_usage_monitor"
}

func (d *FileSystemUsageMonitorDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The file system usage monitor is a utility that checks file systems to determine how much space is available and generates log events to report the usage of each file system. When initially enabled, the monitor immediately checks the file systems.",
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "<p>The administrative state of the configuration.</p><ul><li>To make active, set to enabled.</li><li>To make inactive, set to disabled.</li></ul>",
				Computed:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: "Comments",
				Computed:            true,
			},
			"polling_interval": schema.Int64Attribute{
				MarkdownDescription: "Specify the interval in minutes between file system checks for their usage. Enter a value in the range 15 - 65535. The default value is 60.",
				Computed:            true,
			},
			"all_system": schema.BoolAttribute{
				MarkdownDescription: "Specify whether the utility checks all or only a subset of system file systems. By default, all file systems are scanned. <ul><li>When enabled, you can define specific file systems that override their default thresholds.</li><li>When not enabled, define the file systems to check with their thresholds.</li></ul>",
				Computed:            true,
			},
			"all_system_warning_threshold": schema.Int64Attribute{
				MarkdownDescription: "Specify the usage threshold to generate a warning event when the check is against all system file systems. The threshold is the percentage of the file system that is full. The value for the warning threshold must be less than the critical threshold. Enter a value in the range 0 - 100. The default value is 75.",
				Computed:            true,
			},
			"all_system_critical_threshold": schema.Int64Attribute{
				MarkdownDescription: "Specify the usage threshold to generate a critical event when the check is against all system file systems. The threshold is the percentage of the file system that is full. The value for the critical threshold must be greater than the warning threshold. Enter a value in the range 0 - 100. The default value is 75.",
				Computed:            true,
			},
			"system": schema.ListNestedAttribute{
				MarkdownDescription: "Specify the system file systems to check with their usage thresholds. These thresholds override the thresholds that are defined for all system file systems.",
				NestedObject:        models.GetDmFileSystemUsageDataSourceSchema(),
				Computed:            true,
			},
		},
	}
}

func (d *FileSystemUsageMonitorDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *FileSystemUsageMonitorDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data models.FileSystemUsageMonitor
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
		data.FromBody(ctx, `FileSystemUsageMonitor`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
