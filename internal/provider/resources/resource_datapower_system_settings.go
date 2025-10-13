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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &SystemSettingsResource{}
var _ resource.ResourceWithImportState = &SystemSettingsResource{}

func NewSystemSettingsResource() resource.Resource {
	return &SystemSettingsResource{}
}

type SystemSettingsResource struct {
	pData *tfutils.ProviderData
}

func (r *SystemSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_settings"
}

func (r *SystemSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("<p>System settings provide the following purposes.</p><ul><li>Define system-specific information, such as contact information, location, and name.</li><li>Update serial number after a replacement.</li><li>Enable interface for custom GUI messages and custom CLI prompt.</li><li>Reserve disk space for the audit log.</li><li>Define information about the hardware for use by the SNMP system table, such as serial number, and model type</li></ul>", "system", "").String,
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>The administrative state of the configuration.</p><ul><li>To make active, set to enabled.</li><li>To make inactive, set to disabled.</li></ul>", "admin-state", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter a descriptive summary for the configuration.", "summary", "").String,
				Optional:            true,
			},
			"product_oid": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The read-only string that identifies the installed DataPower agent software.", "", "").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The read-only string that identifies the product.", "", "").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"serial_number": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The read-only string that identifies the product serial number.", "", "").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"entitlement_number": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("After an appliance replacement, the serial number of the original appliance. Without the original serial number, IBM cannot entitle the replacement appliance for maintenance or warranty services.", "entitlement", "").String,
				Optional:            true,
			},
			"product_id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The read-only string that identifies the product type.", "", "").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"capacity_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The read-only installation setting that indicates the licensed capacity mode.", "", "").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"contact": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify any information that identifies the individual or functional area that is responsible maintenance and management.", "contact", "").String,
				Optional:            true,
			},
			"system_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of the DataPower Gateway to use internally as a custom prompt and to use externally to integrate with remote systems. The name must be a 7-bit US-ASCII string of 127 characters or less consisting of letters, numbers, underscore, or embedded dashes, dots, or spaces. However, it is recommended to also be unique with a length of 64 characters or less to be compatible with most remote systems.", "name", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 127),
					stringvalidator.RegexMatches(regexp.MustCompile("^$|^[a-zA-Z0-9_][-_a-zA-Z0-9.\\ ]{0,126}$"), "Must match :"+"^$|^[a-zA-Z0-9_][-_a-zA-Z0-9.\\ ]{0,126}$"),
				},
			},
			"location": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the location of the DataPower Gateway.", "location", "").String,
				Optional:            true,
			},
			"services": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The read-only hex value that indicates support for application, presentation, session, and data-link layer services.", "", "").String,
				Computed:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"backup_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The read-only installation setting that indicates whether a secure-backup is allowed.", "", "").AddStringEnum("normal", "secure").String,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("normal", "secure"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"product_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The read-only installation setting that indicates the operational mode of the product.", "", "").AddStringEnum("normal", "cc").String,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("normal", "cc"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"custom_ui_file": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specifies the URL of the custom user interface file. This file contains custom messages for CLI and GUI sessions as well as the custom CLI prompt. The file must reside in the <tt>store:</tt> or <tt>local:</tt> directory, not on a mounted file system.</p>", "custom-ui-file", "").String,
				Optional:            true,
			},
			"audit_reserve": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the amount of disk space to reserve for audit records. When the disk is full, all services enter the down operational state and stop processing traffic. To restore disk space and resume traffic processing, manual intervention is required. Enter a value in the range 0 - 10000. The default value is 40.", "audit-reserve", "").AddIntegerRange(0, 10000).AddDefaultValue("40").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 10000),
				},
				Default: int64default.StaticInt64(40),
			},
			"detect_intrusion": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Indicates whether to check for intrusion detection.", "detect-intrusion", "").AddStringEnum("enable", "disable").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
			},
			"locale": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the locale for the operating language of the DataPower Gateway. The locale setting manages locale-specific conventions, such as date and time formats, and controls the language of log messages. The language must be enabled before you can select it.", "locale", "").AddStringEnum("de", "en", "es", "fr", "it", "ja", "ko", "pt_BR", "zh_CN", "zh_TW").AddDefaultValue("en").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("de", "en", "es", "fr", "it", "ja", "ko", "pt_BR", "zh_CN", "zh_TW"),
				},
				Default: stringdefault.StaticString("en"),
			},
			"system_log_fixed_format": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Indicates whether to enable fixed format in system logs. When enabled, the system logs are in the format that was used in version 6.0.1 and contain no serviceability improvements after this version that can help with monitoring or troubleshooting.", "system-log-fixed-format", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"uuid": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("UUID", "", "").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *SystemSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *SystemSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SystemSettings
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, "default", data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `SystemSettings`)
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
	getRes, getErr := r.pData.Client.Get(data.GetPath())
	if getErr != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object after creation (GET), got error: %s", getErr))
		return
	}
	data.UpdateUnknownFromBody(ctx, `SystemSettings`, getRes)
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SystemSettings
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

	data.UpdateFromBody(ctx, `SystemSettings`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SystemSettings
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, "default", data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath(), data.ToBody(ctx, `SystemSettings`))
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

func (r *SystemSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SystemSettings
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, "default", data.DependencyActions, actions.Delete, false)
	if resp.Diagnostics.HasError() {
		return
	}
	data.ToDefault()
	_, err := r.pData.Client.Put(data.GetPath(), data.ToBody(ctx, `SystemSettings`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to restore singleton to default, got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *SystemSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

	var data models.SystemSettings
	res, err := r.pData.Client.Get(data.GetPath())
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("Resource was not found, got error: %s", err))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		}
		return
	}

	data.FromBody(ctx, `SystemSettings`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemSettingsResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.SystemSettings

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
