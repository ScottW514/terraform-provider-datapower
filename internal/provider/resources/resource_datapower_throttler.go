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
	"regexp"
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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

var _ resource.Resource = &ThrottlerResource{}
var _ resource.ResourceWithImportState = &ThrottlerResource{}

func NewThrottlerResource() resource.Resource {
	return &ThrottlerResource{}
}

type ThrottlerResource struct {
	pData *tfutils.ProviderData
}

func (r *ThrottlerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_throttler"
}

func (r *ThrottlerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Define throttle settings that monitor memory, temporary file space, XML names, and JSON keys. The system responds to low conditions by refusing to accept new connections. If the refusal does not free sufficient resources after the defined duration, the system restarts.", "throttle", "").String,
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>The administrative state of the configuration.</p><ul><li>To make active, set to enabled.</li><li>To make inactive, set to disabled.</li></ul>", "admin-state", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"throttle_at": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the throttle threshold as a percentage of available memory. At this threshold, the system rejects new connections for the timeout period to allow memory usage to recover. Enter a value in the range 0 - 100. The default value is 20.", "memory-throttle", "").AddIntegerRange(0, 100).AddDefaultValue("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 100),
				},
				Default: int64default.StaticInt64(0),
			},
			"terminate_at": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the terminate threshold as a percentage of available memory. At this threshold, the system reboots. Enter a value in the range 0 - 100. The default value is 5.", "memory-terminate", "").AddIntegerRange(0, 100).AddDefaultValue("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 100),
				},
				Default: int64default.StaticInt64(0),
			},
			"temp_fs_throttle_at": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the throttle threshold as a percentage of available temporary file space. At this threshold, the system rejects new connections for the timeout period to allow temporary file space usage to recover. Enter a value in the range 0 - 100. The default value is 0.", "temp-fs-throttle", "").AddIntegerRange(0, 100).AddDefaultValue("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 100),
				},
				Default: int64default.StaticInt64(0),
			},
			"temp_fs_terminate_at": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the terminate threshold as a percentage of available temporary file space. At this threshold, the system reboots. Enter a value in the range 0 - 100. The default value is 0.", "temp-fs-terminate", "").AddIntegerRange(0, 100).AddDefaultValue("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 100),
				},
				Default: int64default.StaticInt64(0),
			},
			"qname_warn_at": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the threshold as a percentage of available XML names and JSON keys before the system writes an alert to the logs. This threshold is when the number of available XML names or JSON keys in any pool is less than the threshold. Enter a value in the range 5 - 100. The default value is 10.</p><p>As this threshold is approached, the system attempts to free unused resources to prevent this threshold from being reached. If you receive this alert, schedule a reload as soon as possible to prevent an unscheduled restart. If the percentage for any resource pool is less than 5%, the system reboots.</p>", "qcode-warn", "").AddIntegerRange(5, 100).AddDefaultValue("10").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(5, 100),
				},
				Default: int64default.StaticInt64(10),
			},
			"timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration in seconds to reject new connections after a throttle threshold is reached. The default value is 30.", "timeout", "").AddDefaultValue("30").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(30),
			},
			"statistics": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to collect throttle log messages. By default, logging is disabled.", "status-log", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"log_level": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the criticality level for throttle messages. By default, logging is at debug level.", "status-loglevel", "").AddStringEnum("emerg", "alert", "critic", "error", "warn", "notice", "info", "debug").AddDefaultValue("debug").AddNotValidWhen(models.ThrottlerLogLevelIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("emerg", "alert", "critic", "error", "warn", "notice", "info", "debug"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.ThrottlerLogLevelIgnoreVal, true),
				},
				Default: stringdefault.StaticString("debug"),
			},
			"environmental_log": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to collect messages about fans and power supplies and generate messages if a failure event occurs. By default, monitoring is enabled.", "sensors-log", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"backlog_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the size of the backlog queue where incoming requests are routed if a throttling threshold is reached. Enter a value in the range 0 - 500. The default value is 0, which indicates that no requests are routed to the backlog queue.", "backlog-size", "").AddIntegerRange(0, 500).AddDefaultValue("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 500),
				},
				Default: int64default.StaticInt64(0),
			},
			"backlog_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration in seconds that a request remains in the backlog queue before it is rejected if a throttling threshold is reached. Specify a value that is less than the timeout value of your browser. The default value is 30.", "backlog-timeout", "").AddDefaultValue("30").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(30),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *ThrottlerResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *ThrottlerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.Throttler
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, "default", data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `Throttler`)
	_, err := r.pData.Client.Put(data.GetPath(), body)
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			_, err := r.pData.Client.Put(data.GetPath(), body)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Resource already exists. Failed to update resource, got error: %s", err))
				return
			}
			resp.Diagnostics.AddWarning("Warning", "Resource already exists. Existing resource was updated.")
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create resource, got error: %s", err))
			return
		}
	}
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ThrottlerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.Throttler
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

	data.UpdateFromBody(ctx, `Throttler`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ThrottlerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.Throttler
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, "default", data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath(), data.ToBody(ctx, `Throttler`))
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

func (r *ThrottlerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.Throttler
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

func (r *ThrottlerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()
	appDomain := req.ID
	if appDomain != "default" {
		resp.Diagnostics.AddError("Invalid Application Domain", "This resourece supported on the 'default' domain only.")
		return
	}
	if !regexp.MustCompile("^[a-zA-Z0-9_-]+$").MatchString(appDomain) || len(appDomain) < 1 || len(appDomain) > 128 {
		resp.Diagnostics.AddError("Invalid Application Domain", "Application domain must be 1-128 characters and match pattern ^[a-zA-Z0-9_-]+$")
		return
	}

	var data models.Throttler
	res, err := r.pData.Client.Get(data.GetPath())
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("Resource was not found, got error: %s", err))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		}
		return
	}

	data.FromBody(ctx, `Throttler`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ThrottlerResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.Throttler

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
