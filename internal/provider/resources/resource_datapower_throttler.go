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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &ThrottlerResource{}

func NewThrottlerResource() resource.Resource {
	return &ThrottlerResource{}
}

type ThrottlerResource struct {
	client *client.DatapowerClient
}

func (r *ThrottlerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_throttler"
}

func (r *ThrottlerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Throttle settings (`default` domain only)", "throttle", "").String,

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
			"throttle_at": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Memory throttle threshold", "memory-throttle", "").AddIntegerRange(0, 100).AddDefaultValue("20").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 100),
				},
				Default: int64default.StaticInt64(20),
			},
			"terminate_at": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Memory terminate threshold", "memory-terminate", "").AddIntegerRange(0, 100).AddDefaultValue("5").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 100),
				},
				Default: int64default.StaticInt64(5),
			},
			"temp_fs_throttle_at": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Temporary file space throttle threshold", "temp-fs-throttle", "").AddIntegerRange(0, 100).AddDefaultValue("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 100),
				},
				Default: int64default.StaticInt64(0),
			},
			"temp_fs_terminate_at": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Temporary file space terminate threshold", "temp-fs-terminate", "").AddIntegerRange(0, 100).AddDefaultValue("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 100),
				},
				Default: int64default.StaticInt64(0),
			},
			"qname_warn_at": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML names and JSON keys threshold", "qcode-warn", "").AddIntegerRange(5, 100).AddDefaultValue("10").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(5, 100),
				},
				Default: int64default.StaticInt64(10),
			},
			"timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Timeout", "timeout", "").AddDefaultValue("30").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(30),
			},
			"statistics": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Status log", "status-log", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"log_level": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Log level", "status-loglevel", "").AddStringEnum("emerg", "alert", "critic", "error", "warn", "notice", "info", "debug").AddDefaultValue("debug").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("emerg", "alert", "critic", "error", "warn", "notice", "info", "debug"),
				},
				Default: stringdefault.StaticString("debug"),
			},
			"environmental_log": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Monitor sensors", "sensors-log", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"backlog_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Backlog size", "backlog-size", "").AddIntegerRange(0, 500).AddDefaultValue("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 500),
				},
				Default: int64default.StaticInt64(0),
			},
			"backlog_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Backlog timeout", "backlog-timeout", "").AddDefaultValue("30").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(30),
			},
		},
	}
}

func (r *ThrottlerResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *ThrottlerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.Throttler

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `Throttler`)
	_, err := r.client.Put(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "PUT", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ThrottlerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.Throttler

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
		data.FromBody(ctx, `Throttler`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `Throttler`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ThrottlerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.Throttler

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Put(data.GetPath(), data.ToBody(ctx, `Throttler`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ThrottlerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	resp.State.RemoveResource(ctx)
}
