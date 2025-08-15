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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &TimeSettingsResource{}

func NewTimeSettingsResource() resource.Resource {
	return &TimeSettingsResource{}
}

type TimeSettingsResource struct {
	pData *tfutils.ProviderData
}

func (r *TimeSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_time_settings"
}

func (r *TimeSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("<p>The time zone for the system. The time zone affects the time that the system displays. All timestamps use this time zone.</p><p><b>Note: </b>The system clock runs on Coordinated Universal Time (UTC).</p><p>When daylight saving time (DST), or summer time, applies to a time zone, the system adjusts the time when a DST boundary is crossed.</p>", "timezone", "").String,
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>The administrative state of the configuration.</p><ul><li>To make active, set to enabled.</li><li>To make inactive, set to disabled.</li></ul>", "admin-state", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"local_time_zone": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the time zone to use in management interfaces. The default value is EST-5EDT.", "name", "").AddStringEnum("HST10", "AKST9AKDT", "PST8PDT", "MST7MDT", "CST6CDT", "EST5EDT", "AST4ADT", "UTC", "GMT0BST", "CET-1CEST", "EET-2EEST", "MKS-3MSD", "MSK-3MSD", "AST-3", "KRT-5", "IST-5:30", "NOVST-6NOVDT", "CST-8", "WST-8", "WST-8WDT", "JST-9", "CST-9:30CDT", "EST-10EDT", "EST-10", "Custom").AddDefaultValue("EST5EDT").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("HST10", "AKST9AKDT", "PST8PDT", "MST7MDT", "CST6CDT", "EST5EDT", "AST4ADT", "UTC", "GMT0BST", "CET-1CEST", "EET-2EEST", "MKS-3MSD", "MSK-3MSD", "AST-3", "KRT-5", "IST-5:30", "NOVST-6NOVDT", "CST-8", "WST-8", "WST-8WDT", "JST-9", "CST-9:30CDT", "EST-10EDT", "EST-10", "Custom"),
				},
				Default: stringdefault.StaticString("EST5EDT"),
			},
			"custom_tz_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the symbolic name for the custom time zone. This name is appended to local times. The name must be three or more alphabetic characters. If you use any other characters, the time zone becomes UTC.", "custom", "").AddDefaultValue("STD").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("STD"),
			},
			"utc_direction": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the direction relative to UTC for the custom time zone. Asia is east. North America is west. The default value is East.", "direction", "").AddStringEnum("East", "West").AddDefaultValue("West").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("East", "West"),
				},
				Default: stringdefault.StaticString("West"),
			},
			"offset_hours": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the number of hours the custom time zone is from UTC. If 2 hours and 30 minutes from UTC, enter 2.", "offset-hours", "").AddIntegerRange(0, 12).String,
				Optional:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 12),
				},
			},
			"offset_minutes": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the number of minutes the time zone is from UTC. If 2 hours and 30 minutes from UTC, enter 30.", "offset-minutes", "").AddIntegerRange(0, 59).String,
				Optional:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 59),
				},
			},
			"daylight_offset_hours": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the offset in hours when the custom time zone observes DST. Generally, the offset is 1 hour. The default value is 1.", "daylight-offset", "").AddIntegerRange(0, 12).AddDefaultValue("1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 12),
				},
				Default: int64default.StaticInt64(1),
			},
			"tz_name_dst": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the symbolic name for the custom time zone during DST. This name is appended to local times. The name must be three or more alphabetic characters. If you use any other characters, the time zone becomes UTC.", "daylight-name", "").AddDefaultValue("DST").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("DST"),
			},
			"daylight_start_month": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the month when DST starts for the custom time zone. The default value is March.", "daylight-start-month", "").AddStringEnum("January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December").AddDefaultValue("March").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"),
				},
				Default: stringdefault.StaticString("March"),
			},
			"daylight_start_week": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the instance of the day in the month when DST starts for the custom time zone. If DST starts on the second Sunday in the month, enter 2.", "daylight-start-week", "").AddIntegerRange(1, 5).AddDefaultValue("2").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 5),
				},
				Default: int64default.StaticInt64(2),
			},
			"daylight_start_day": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the day of the week when DST starts for the custom time zone. The default value is Sunday.", "daylight-start-day", "").AddStringEnum("Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday").AddDefaultValue("Sunday").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"),
				},
				Default: stringdefault.StaticString("Sunday"),
			},
			"daylight_start_time_hours": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the hour when DST starts for the custom time zone. If the start boundary is 2:30 AM, enter 2.", "daylight-start-hours", "").AddIntegerRange(0, 23).AddDefaultValue("2").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 23),
				},
				Default: int64default.StaticInt64(2),
			},
			"daylight_start_time_minutes": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the minute when DST starts for the custom time zone. If the start boundary is 2:30 AM, enter 30.", "daylight-start-minutes", "").AddIntegerRange(0, 59).String,
				Optional:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 59),
				},
			},
			"daylight_stop_month": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the month when DST ends for the custom time zone. The default value is November.", "daylight-stop-month", "").AddStringEnum("January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December").AddDefaultValue("November").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"),
				},
				Default: stringdefault.StaticString("November"),
			},
			"daylight_stop_week": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the instance of the day in the month when DST ends for the custom time zone. If DST ends on the second Sunday in the month, enter 2.", "daylight-stop-week", "").AddIntegerRange(1, 5).AddDefaultValue("1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 5),
				},
				Default: int64default.StaticInt64(1),
			},
			"daylight_stop_day": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the day of the week when DST ends for the custom time zone. The default value is Sunday.", "daylight-stop-day", "").AddStringEnum("Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday").AddDefaultValue("Sunday").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"),
				},
				Default: stringdefault.StaticString("Sunday"),
			},
			"daylight_stop_time_hours": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the hour when DST ends for the custom time zone. If the end boundary is 2:30 AM, enter 2.", "daylight-stop-hours", "").AddIntegerRange(0, 23).AddDefaultValue("2").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 23),
				},
				Default: int64default.StaticInt64(2),
			},
			"daylight_stop_time_minutes": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the minute when DST ends for the custom time zone. If the end boundary is 2:30 AM, enter 30.", "daylight-stop-minutes", "").AddIntegerRange(0, 59).String,
				Optional:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 59),
				},
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *TimeSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *TimeSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.TimeSettings
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, "default", data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `TimeSettings`)
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

func (r *TimeSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.TimeSettings
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
		data.FromBody(ctx, `TimeSettings`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `TimeSettings`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TimeSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.TimeSettings
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, "default", data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath(), data.ToBody(ctx, `TimeSettings`))
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

func (r *TimeSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.TimeSettings
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

func (r *TimeSettingsResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.TimeSettings

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
