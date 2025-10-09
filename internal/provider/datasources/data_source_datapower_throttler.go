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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var (
	_ datasource.DataSource              = &ThrottlerDataSource{}
	_ datasource.DataSourceWithConfigure = &ThrottlerDataSource{}
)

func NewThrottlerDataSource() datasource.DataSource {
	return &ThrottlerDataSource{}
}

type ThrottlerDataSource struct {
	pData *tfutils.ProviderData
}

func (d *ThrottlerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_throttler"
}

func (d *ThrottlerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Define throttle settings that monitor memory, temporary file space, XML names, and JSON keys. The system responds to low conditions by refusing to accept new connections. If the refusal does not free sufficient resources after the defined duration, the system restarts.",
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "<p>The administrative state of the configuration.</p><ul><li>To make active, set to enabled.</li><li>To make inactive, set to disabled.</li></ul>",
				Computed:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: "Comments",
				Computed:            true,
			},
			"throttle_at": schema.Int64Attribute{
				MarkdownDescription: "Specify the throttle threshold as a percentage of available memory. At this threshold, the system rejects new connections for the timeout period to allow memory usage to recover. Enter a value in the range 0 - 100. The default value is 20.",
				Computed:            true,
			},
			"terminate_at": schema.Int64Attribute{
				MarkdownDescription: "Specify the terminate threshold as a percentage of available memory. At this threshold, the system reboots. Enter a value in the range 0 - 100. The default value is 5.",
				Computed:            true,
			},
			"temp_fs_throttle_at": schema.Int64Attribute{
				MarkdownDescription: "Specify the throttle threshold as a percentage of available temporary file space. At this threshold, the system rejects new connections for the timeout period to allow temporary file space usage to recover. Enter a value in the range 0 - 100. The default value is 0.",
				Computed:            true,
			},
			"temp_fs_terminate_at": schema.Int64Attribute{
				MarkdownDescription: "Specify the terminate threshold as a percentage of available temporary file space. At this threshold, the system reboots. Enter a value in the range 0 - 100. The default value is 0.",
				Computed:            true,
			},
			"qname_warn_at": schema.Int64Attribute{
				MarkdownDescription: "<p>Specify the threshold as a percentage of available XML names and JSON keys before the system writes an alert to the logs. This threshold is when the number of available XML names or JSON keys in any pool is less than the threshold. Enter a value in the range 5 - 100. The default value is 10.</p><p>As this threshold is approached, the system attempts to free unused resources to prevent this threshold from being reached. If you receive this alert, schedule a reload as soon as possible to prevent an unscheduled restart. If the percentage for any resource pool is less than 5%, the system reboots.</p>",
				Computed:            true,
			},
			"timeout": schema.Int64Attribute{
				MarkdownDescription: "Specify the duration in seconds to reject new connections after a throttle threshold is reached. The default value is 30.",
				Computed:            true,
			},
			"statistics": schema.BoolAttribute{
				MarkdownDescription: "Specify whether to collect throttle log messages. By default, logging is disabled.",
				Computed:            true,
			},
			"log_level": schema.StringAttribute{
				MarkdownDescription: "Specify the criticality level for throttle messages. By default, logging is at debug level.",
				Computed:            true,
			},
			"environmental_log": schema.BoolAttribute{
				MarkdownDescription: "Specify whether to collect messages about fans and power supplies and generate messages if a failure event occurs. By default, monitoring is enabled.",
				Computed:            true,
			},
			"backlog_size": schema.Int64Attribute{
				MarkdownDescription: "Specify the size of the backlog queue where incoming requests are routed if a throttling threshold is reached. Enter a value in the range 0 - 500. The default value is 0, which indicates that no requests are routed to the backlog queue.",
				Computed:            true,
			},
			"backlog_timeout": schema.Int64Attribute{
				MarkdownDescription: "Specify the duration in seconds that a request remains in the backlog queue before it is rejected if a throttling threshold is reached. Specify a value that is less than the timeout value of your browser. The default value is 30.",
				Computed:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (d *ThrottlerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *ThrottlerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data models.Throttler
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
		data.FromBody(ctx, `Throttler`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
