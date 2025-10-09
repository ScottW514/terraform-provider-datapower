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
	_ datasource.DataSource              = &TimeSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &TimeSettingsDataSource{}
)

func NewTimeSettingsDataSource() datasource.DataSource {
	return &TimeSettingsDataSource{}
}

type TimeSettingsDataSource struct {
	pData *tfutils.ProviderData
}

func (d *TimeSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_time_settings"
}

func (d *TimeSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "<p>The time zone for the system. The time zone affects the time that the system displays. All timestamps use this time zone.</p><p><b>Note: </b>The system clock runs on Coordinated Universal Time (UTC).</p><p>When daylight saving time (DST), or summer time, applies to a time zone, the system adjusts the time when a DST boundary is crossed.</p>",
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "<p>The administrative state of the configuration.</p><ul><li>To make active, set to enabled.</li><li>To make inactive, set to disabled.</li></ul>",
				Computed:            true,
			},
			"local_time_zone": schema.StringAttribute{
				MarkdownDescription: "Specify the time zone to use in management interfaces. The default value is EST-5EDT.",
				Computed:            true,
			},
			"custom_tz_name": schema.StringAttribute{
				MarkdownDescription: "Specify the symbolic name for the custom time zone. This name is appended to local times. The name must be three or more alphabetic characters. If you use any other characters, the time zone becomes UTC.",
				Computed:            true,
			},
			"utc_direction": schema.StringAttribute{
				MarkdownDescription: "Specify the direction relative to UTC for the custom time zone. Asia is east. North America is west. The default value is East.",
				Computed:            true,
			},
			"offset_hours": schema.Int64Attribute{
				MarkdownDescription: "Specify the number of hours the custom time zone is from UTC. If 2 hours and 30 minutes from UTC, enter 2.",
				Computed:            true,
			},
			"offset_minutes": schema.Int64Attribute{
				MarkdownDescription: "Specify the number of minutes the time zone is from UTC. If 2 hours and 30 minutes from UTC, enter 30.",
				Computed:            true,
			},
			"daylight_offset_hours": schema.Int64Attribute{
				MarkdownDescription: "Specify the offset in hours when the custom time zone observes DST. Generally, the offset is 1 hour. The default value is 1.",
				Computed:            true,
			},
			"tz_name_dst": schema.StringAttribute{
				MarkdownDescription: "Specify the symbolic name for the custom time zone during DST. This name is appended to local times. The name must be three or more alphabetic characters. If you use any other characters, the time zone becomes UTC.",
				Computed:            true,
			},
			"daylight_start_month": schema.StringAttribute{
				MarkdownDescription: "Specify the month when DST starts for the custom time zone. The default value is March.",
				Computed:            true,
			},
			"daylight_start_week": schema.Int64Attribute{
				MarkdownDescription: "Specify the instance of the day in the month when DST starts for the custom time zone. If DST starts on the second Sunday in the month, enter 2.",
				Computed:            true,
			},
			"daylight_start_day": schema.StringAttribute{
				MarkdownDescription: "Specify the day of the week when DST starts for the custom time zone. The default value is Sunday.",
				Computed:            true,
			},
			"daylight_start_time_hours": schema.Int64Attribute{
				MarkdownDescription: "Specify the hour when DST starts for the custom time zone. If the start boundary is 2:30 AM, enter 2.",
				Computed:            true,
			},
			"daylight_start_time_minutes": schema.Int64Attribute{
				MarkdownDescription: "Specify the minute when DST starts for the custom time zone. If the start boundary is 2:30 AM, enter 30.",
				Computed:            true,
			},
			"daylight_stop_month": schema.StringAttribute{
				MarkdownDescription: "Specify the month when DST ends for the custom time zone. The default value is November.",
				Computed:            true,
			},
			"daylight_stop_week": schema.Int64Attribute{
				MarkdownDescription: "Specify the instance of the day in the month when DST ends for the custom time zone. If DST ends on the second Sunday in the month, enter 2.",
				Computed:            true,
			},
			"daylight_stop_day": schema.StringAttribute{
				MarkdownDescription: "Specify the day of the week when DST ends for the custom time zone. The default value is Sunday.",
				Computed:            true,
			},
			"daylight_stop_time_hours": schema.Int64Attribute{
				MarkdownDescription: "Specify the hour when DST ends for the custom time zone. If the end boundary is 2:30 AM, enter 2.",
				Computed:            true,
			},
			"daylight_stop_time_minutes": schema.Int64Attribute{
				MarkdownDescription: "Specify the minute when DST ends for the custom time zone. If the end boundary is 2:30 AM, enter 30.",
				Computed:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (d *TimeSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *TimeSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data models.TimeSettings
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
		data.FromBody(ctx, `TimeSettings`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
