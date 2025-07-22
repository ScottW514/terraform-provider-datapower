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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &JOSERecipientIdentifierResource{}

func NewJOSERecipientIdentifierResource() resource.Resource {
	return &JOSERecipientIdentifierResource{}
}

type JOSERecipientIdentifierResource struct {
	client *client.DatapowerClient
}

func (r *JOSERecipientIdentifierResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_joserecipientidentifier"
}

func (r *JOSERecipientIdentifierResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Recipient Identifier", "jose-recipient-identifier", "").String,

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
			"type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Key Material Type", "type", "").AddStringEnum("key", "sskey").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("key", "sskey"),
				},
			},
			"key": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Key Material", "key", "cryptokey").String,
				Optional:            true,
			},
			"ss_key": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Key Material", "sskey", "cryptosskey").String,
				Optional:            true,
			},
			"header_param": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Header Parameters", "header-param", "").String,
				NestedObject:        models.DmJOSEHeaderResourceSchema,
				Required:            true,
			},
		},
	}
}

func (r *JOSERecipientIdentifierResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *JOSERecipientIdentifierResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.JOSERecipientIdentifier

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `JOSERecipientIdentifier`)
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

func (r *JOSERecipientIdentifierResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.JOSERecipientIdentifier

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
		data.FromBody(ctx, `JOSERecipientIdentifier`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `JOSERecipientIdentifier`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *JOSERecipientIdentifierResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.JOSERecipientIdentifier

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `JOSERecipientIdentifier`))
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

func (r *JOSERecipientIdentifierResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.JOSERecipientIdentifier

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
