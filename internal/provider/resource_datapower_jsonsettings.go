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

package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &JSONSettingsResource{}

func NewJSONSettingsResource() resource.Resource {
	return &JSONSettingsResource{}
}

type JSONSettingsResource struct {
	client *client.DatapowerClient
}

func (r *JSONSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_jsonsettings"
}

func (r *JSONSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("JSON settings", "json-settings", "").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name of the object. Must be unique among object types in application domain.", "", "").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"app_domain": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The name of the application domain the object belongs to", "", "").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					ImmutableAfterSet(),
				},
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"json_max_nesting_depth": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Max nesting depth", "max-nesting-depth", "").AddIntegerRange(64, 256).AddDefaultValue("64").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(64, 256),
				},
				Default: int64default.StaticInt64(64),
			},
			"json_max_label_length": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Max label length", "max-label-string-length", "").AddIntegerRange(256, 8192).AddDefaultValue("256").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(256, 8192),
				},
				Default: int64default.StaticInt64(256),
			},
			"json_max_value_length": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Max value length for strings", "max-value-string-length", "").AddIntegerRange(8192, 5368709121).AddDefaultValue("8192").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(8192, 5368709121),
				},
				Default: int64default.StaticInt64(8192),
			},
			"json_max_number_length": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Max value length for numbers", "max-number-length", "").AddIntegerRange(128, 256).AddDefaultValue("128").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(128, 256),
				},
				Default: int64default.StaticInt64(128),
			},
			"json_document_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Max document size", "max-document-size", "").AddIntegerRange(4194304, 5368709121).AddDefaultValue("4194304").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(4194304, 5368709121),
				},
				Default: int64default.StaticInt64(4194304),
			},
		},
	}
}

func (r *JSONSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *JSONSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.JSONSettings

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `JSONSettings`)
	_, err := r.client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "POST", err))
		return
	}

	_, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s", "POST", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *JSONSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.JSONSettings

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && (strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}

	if data.IsNull() {
		// Import
		data.FromBody(ctx, `JSONSettings`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `JSONSettings`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *JSONSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.JSONSettings

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `JSONSettings`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}
	_, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s", "POST", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *JSONSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.JSONSettings

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}

	resp.State.RemoveResource(ctx)
}
