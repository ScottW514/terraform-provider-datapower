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
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &TimeSettingsResource{}

func NewTimeSettingsResource() resource.Resource {
	return &TimeSettingsResource{}
}

type TimeSettingsResource struct {
	client *client.DatapowerClient
}

func (r *TimeSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_timesettings"
}

func (r *TimeSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Time Settings (`default` domain only)", "timezone", "").String,
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Administrative state", "admin-state", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"local_time_zone": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Local time zone", "name", "").AddStringEnum("HST10", "AKST9AKDT", "PST8PDT", "MST7MDT", "CST6CDT", "EST5EDT", "AST4ADT", "UTC", "GMT0BST", "CET-1CEST", "EET-2EEST", "MKS-3MSD", "MSK-3MSD", "AST-3", "KRT-5", "IST-5:30", "NOVST-6NOVDT", "CST-8", "WST-8", "WST-8WDT", "JST-9", "CST-9:30CDT", "EST-10EDT", "EST-10", "Custom").AddDefaultValue("EST5EDT").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("HST10", "AKST9AKDT", "PST8PDT", "MST7MDT", "CST6CDT", "EST5EDT", "AST4ADT", "UTC", "GMT0BST", "CET-1CEST", "EET-2EEST", "MKS-3MSD", "MSK-3MSD", "AST-3", "KRT-5", "IST-5:30", "NOVST-6NOVDT", "CST-8", "WST-8", "WST-8WDT", "JST-9", "CST-9:30CDT", "EST-10EDT", "EST-10", "Custom"),
				},
				Default: stringdefault.StaticString("EST5EDT"),
			},
			"custom_tz_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name", "custom", "").AddDefaultValue("STD").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("STD"),
			},
			"utc_direction": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Direction from UTC", "direction", "").AddStringEnum("East", "West").AddDefaultValue("West").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("East", "West"),
				},
				Default: stringdefault.StaticString("West"),
			},
			"offset_hours": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Hours from UTC", "offset-hours", "").AddIntegerRange(0, 12).String,
				Optional:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 12),
				},
			},
			"offset_minutes": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Minutes from UTC", "offset-minutes", "").AddIntegerRange(0, 59).String,
				Optional:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 59),
				},
			},
			"daylight_offset_hours": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Daylight savings time (DST) offset", "daylight-offset", "").AddIntegerRange(0, 12).AddDefaultValue("1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 12),
				},
				Default: int64default.StaticInt64(1),
			},
			"tz_name_dst": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("DST name", "daylight-name", "").AddDefaultValue("DST").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("DST"),
			},
			"daylight_start_month": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("DST start month", "daylight-start-month", "").AddStringEnum("January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December").AddDefaultValue("March").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"),
				},
				Default: stringdefault.StaticString("March"),
			},
			"daylight_start_week": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("DST start day instance", "daylight-start-week", "").AddIntegerRange(1, 5).AddDefaultValue("2").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 5),
				},
				Default: int64default.StaticInt64(2),
			},
			"daylight_start_day": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("DST start day", "daylight-start-day", "").AddStringEnum("Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday").AddDefaultValue("Sunday").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"),
				},
				Default: stringdefault.StaticString("Sunday"),
			},
			"daylight_start_time_hours": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("DST start hour", "daylight-start-hours", "").AddIntegerRange(0, 23).AddDefaultValue("2").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 23),
				},
				Default: int64default.StaticInt64(2),
			},
			"daylight_start_time_minutes": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("DST start minute", "daylight-start-minutes", "").AddIntegerRange(0, 59).String,
				Optional:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 59),
				},
			},
			"daylight_stop_month": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("DST stop month", "daylight-stop-month", "").AddStringEnum("January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December").AddDefaultValue("November").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"),
				},
				Default: stringdefault.StaticString("November"),
			},
			"daylight_stop_week": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("DST stop day instance", "daylight-stop-week", "").AddIntegerRange(1, 5).AddDefaultValue("1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 5),
				},
				Default: int64default.StaticInt64(1),
			},
			"daylight_stop_day": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("DST stop day", "daylight-stop-day", "").AddStringEnum("Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday").AddDefaultValue("Sunday").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"),
				},
				Default: stringdefault.StaticString("Sunday"),
			},
			"daylight_stop_time_hours": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("DST stop hour", "daylight-stop-hours", "").AddIntegerRange(0, 23).AddDefaultValue("2").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 23),
				},
				Default: int64default.StaticInt64(2),
			},
			"daylight_stop_time_minutes": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("DST stop minutes", "daylight-stop-minutes", "").AddIntegerRange(0, 59).String,
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

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *TimeSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.TimeSettings

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, "default", data.DependencyActions, actions.Create, false)

	body := data.ToBody(ctx, `TimeSettings`)
	_, err := r.client.Put(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "PUT", err))
		return
	}
	actions.PostProcess(ctx, &resp.Diagnostics, r.client, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TimeSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.TimeSettings

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
		data.FromBody(ctx, `TimeSettings`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `TimeSettings`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TimeSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.TimeSettings

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, "default", data.DependencyActions, actions.Update, false)
	_, err := r.client.Put(data.GetPath(), data.ToBody(ctx, `TimeSettings`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, r.client, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TimeSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.TimeSettings

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, "default", data.DependencyActions, actions.Delete, false)
	actions.PostProcess(ctx, &resp.Diagnostics, r.client, data.DependencyActions, actions.Create)
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
