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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &SystemSettingsResource{}

func NewSystemSettingsResource() resource.Resource {
	return &SystemSettingsResource{}
}

type SystemSettingsResource struct {
	client *client.DatapowerClient
}

func (r *SystemSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemsettings"
}

func (r *SystemSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("System settings (`default` domain only)", "system", "").String,

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
			"product_oid": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Product OID", "", "").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Description", "", "").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"serial_number": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Serial number", "", "").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"entitlement_number": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Entitlement serial number", "entitlement", "").String,
				Optional:            true,
			},
			"product_id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Product ID", "", "").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"capacity_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Licensed capacity mode", "", "").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"contact": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Contact", "contact", "").String,
				Optional:            true,
			},
			"system_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("System name", "name", "").String,
				Optional:            true,
			},
			"location": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Location", "location", "").String,
				Optional:            true,
			},
			"services": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Services", "", "").String,
				Computed:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"backup_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Backup mode", "", "").AddStringEnum("normal", "secure").String,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("normal", "secure"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"product_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Product Mode", "", "").AddStringEnum("normal", "cc").String,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("normal", "cc"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"custom_ui_file": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Custom user interface file", "custom-ui-file", "").String,
				Optional:            true,
			},
			"audit_reserve": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Audit reserve space", "audit-reserve", "").AddIntegerRange(0, 10000).AddDefaultValue("40").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 10000),
				},
				Default: int64default.StaticInt64(40),
			},
			"detect_intrusion": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Detect intrusion (physical appliances only)", "detect-intrusion", "").AddStringEnum("enable", "disable").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
			},
			"hardware_xml_acceleration": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable hardware XML acceleration (physical appliances only)", "xml-accelerator", "").String,
				Optional:            true,
			},
			"locale": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("System locale", "locale", "").AddStringEnum("de", "en", "es", "fr", "it", "ja", "ko", "pt_BR", "zh_CN", "zh_TW").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("de", "en", "es", "fr", "it", "ja", "ko", "pt_BR", "zh_CN", "zh_TW"),
				},
			},
			"system_log_fixed_format": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable fixed format", "system-log-fixed-format", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"uuid": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("UUID", "", "").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *SystemSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *SystemSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.SystemSettings

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `SystemSettings`)
	_, err := r.client.Put(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "PUT", err))
		return
	}
	getRes, getErr := r.client.Get(data.GetPath())
	if getErr != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object after creation (GET), got error: %s", getErr))
		return
	}
	data.UpdateUnknownFromBody(ctx, `SystemSettings`, getRes)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.SystemSettings

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
		data.FromBody(ctx, `SystemSettings`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `SystemSettings`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.SystemSettings

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Put(data.GetPath(), data.ToBody(ctx, `SystemSettings`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	resp.State.RemoveResource(ctx)
}
