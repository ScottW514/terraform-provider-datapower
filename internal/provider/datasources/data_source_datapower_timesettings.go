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
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
)

var (
	_ datasource.DataSource              = &TimeSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &TimeSettingsDataSource{}
)

func NewTimeSettingsDataSource() datasource.DataSource {
	return &TimeSettingsDataSource{}
}

type TimeSettingsDataSource struct {
	client *client.DatapowerClient
}

func (d *TimeSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_timesettings"
}

func (d *TimeSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Time Settings (`default` domain only)",
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Administrative state",
				Computed:            true,
			},
			"local_time_zone": schema.StringAttribute{
				MarkdownDescription: "Local time zone",
				Computed:            true,
			},
			"custom_tz_name": schema.StringAttribute{
				MarkdownDescription: "Name",
				Computed:            true,
			},
			"utc_direction": schema.StringAttribute{
				MarkdownDescription: "Direction from UTC",
				Computed:            true,
			},
			"offset_hours": schema.Int64Attribute{
				MarkdownDescription: "Hours from UTC",
				Computed:            true,
			},
			"offset_minutes": schema.Int64Attribute{
				MarkdownDescription: "Minutes from UTC",
				Computed:            true,
			},
			"daylight_offset_hours": schema.Int64Attribute{
				MarkdownDescription: "Daylight savings time (DST) offset",
				Computed:            true,
			},
			"tz_name_dst": schema.StringAttribute{
				MarkdownDescription: "DST name",
				Computed:            true,
			},
			"daylight_start_month": schema.StringAttribute{
				MarkdownDescription: "DST start month",
				Computed:            true,
			},
			"daylight_start_week": schema.Int64Attribute{
				MarkdownDescription: "DST start day instance",
				Computed:            true,
			},
			"daylight_start_day": schema.StringAttribute{
				MarkdownDescription: "DST start day",
				Computed:            true,
			},
			"daylight_start_time_hours": schema.Int64Attribute{
				MarkdownDescription: "DST start hour",
				Computed:            true,
			},
			"daylight_start_time_minutes": schema.Int64Attribute{
				MarkdownDescription: "DST start minute",
				Computed:            true,
			},
			"daylight_stop_month": schema.StringAttribute{
				MarkdownDescription: "DST stop month",
				Computed:            true,
			},
			"daylight_stop_week": schema.Int64Attribute{
				MarkdownDescription: "DST stop day instance",
				Computed:            true,
			},
			"daylight_stop_day": schema.StringAttribute{
				MarkdownDescription: "DST stop day",
				Computed:            true,
			},
			"daylight_stop_time_hours": schema.Int64Attribute{
				MarkdownDescription: "DST stop hour",
				Computed:            true,
			},
			"daylight_stop_time_minutes": schema.Int64Attribute{
				MarkdownDescription: "DST stop minutes",
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

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *TimeSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.TimeSettings

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := d.client.Get(data.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	data.FromBody(ctx, `TimeSettings`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
