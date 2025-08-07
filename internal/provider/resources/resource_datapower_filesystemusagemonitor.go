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

var _ resource.Resource = &FileSystemUsageMonitorResource{}

func NewFileSystemUsageMonitorResource() resource.Resource {
	return &FileSystemUsageMonitorResource{}
}

type FileSystemUsageMonitorResource struct {
	client *client.DatapowerClient
}

func (r *FileSystemUsageMonitorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_filesystemusagemonitor"
}

func (r *FileSystemUsageMonitorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("File system usage monitor (`default` domain only)", "fs-usage-monitor", "").String,
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
			"polling_interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Polling interval", "poll", "").AddIntegerRange(15, 65535).AddDefaultValue("60").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(15, 65535),
				},
				Default: int64default.StaticInt64(60),
			},
			"all_system": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Monitor all system file systems", "all-system", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"all_system_warning_threshold": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Warning threshold for all system file systems", "all-system-warning", "").AddIntegerRange(0, 100).AddDefaultValue("75").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 100),
				},
				Default: int64default.StaticInt64(75),
			},
			"all_system_critical_threshold": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Critical threshold for all system file systems", "all-system-critical", "").AddIntegerRange(0, 100).AddDefaultValue("90").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 100),
				},
				Default: int64default.StaticInt64(90),
			},
			"system": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("System file systems", "system", "").String,
				NestedObject:        models.DmFileSystemUsageResourceSchema,
				Optional:            true,
			},
			"all_qm_warning_threshold": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Warning threshold for queue manager file systems", "all-qm-warning", "").AddIntegerRange(0, 100).AddDefaultValue("75").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 100),
				},
				Default: int64default.StaticInt64(75),
			},
			"all_qm_critical_threshold": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Critical threshold for queue manager file systems", "all-qm-critical", "").AddIntegerRange(0, 100).AddDefaultValue("90").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 100),
				},
				Default: int64default.StaticInt64(90),
			},
			"queue_manager": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Queue manager file systems", "qm", "").String,
				NestedObject:        models.DmQMFileSystemUsageResourceSchema,
				Optional:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *FileSystemUsageMonitorResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *FileSystemUsageMonitorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.FileSystemUsageMonitor

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, "default", data.DependencyActions, actions.Create)

	body := data.ToBody(ctx, `FileSystemUsageMonitor`)
	_, err := r.client.Put(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "PUT", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FileSystemUsageMonitorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.FileSystemUsageMonitor

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
		data.FromBody(ctx, `FileSystemUsageMonitor`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `FileSystemUsageMonitor`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FileSystemUsageMonitorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.FileSystemUsageMonitor

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, "default", data.DependencyActions, actions.Update)
	_, err := r.client.Put(data.GetPath(), data.ToBody(ctx, `FileSystemUsageMonitor`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FileSystemUsageMonitorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	resp.State.RemoveResource(ctx)
}
