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
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &B2BCPAReceiverSettingResource{}

func NewB2BCPAReceiverSettingResource() resource.Resource {
	return &B2BCPAReceiverSettingResource{}
}

type B2BCPAReceiverSettingResource struct {
	client *client.DatapowerClient
}

func (r *B2BCPAReceiverSettingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_b2bcpareceiversetting"
}

func (r *B2BCPAReceiverSettingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("CPA receiver setting", "cpa-receiver-setting", "").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name of the object. Must be unique among object types in application domain.", "", "").String,
				Required:            true,
				Validators: []validator.String{

					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"app_domain": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The name of the application domain the object belongs to", "", "").String,
				Required:            true,
				Validators: []validator.String{

					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`), ""),
				},
				PlanModifiers: []planmodifier.String{
					ImmutableAfterSet(),
				},
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"local_endpoint_uri": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Local URI", "local-uri", "").String,
				Optional:            true,
			},
			"sync_reply_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Sync reply mode", "syncreply-mode", "").AddStringEnum("mshSignalsOnly", "none").AddDefaultValue("none").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("mshSignalsOnly", "none"),
				},
				Default: stringdefault.StaticString("none"),
			},
			"ack_requested": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Expect acknowledgment requests", "ack-requested", "").AddStringEnum("never", "always", "perMessage").AddDefaultValue("perMessage").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("never", "always", "perMessage"),
				},
				Default: stringdefault.StaticString("perMessage"),
			},
			"ack_signature_requested": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Expect signed acknowledgment requests", "ack-signature-requested", "").AddStringEnum("never", "always", "perMessage").AddDefaultValue("perMessage").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("never", "always", "perMessage"),
				},
				Default: stringdefault.StaticString("perMessage"),
			},
			"allow_duplicate_message": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow duplicate messages", "allow-dup-msg", "").AddStringEnum("never", "always", "on-error").AddDefaultValue("never").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("never", "always", "on-error"),
				},
				Default: stringdefault.StaticString("never"),
			},
			"persist_duration": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Persistence duration", "persist-duration", "").AddIntegerRange(0, 6000000).String,
				Optional:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 6000000),
				},
			},
			"encryption_required": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Require encryption", "encrypt-required", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"decrypt_id_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Decryption identification credentials", "decrypt-idcred", "cryptoidentcred").String,
				Optional:            true,
			},
			"signature_required": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Require signature", "sign-required", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"verify_val_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Signature validation credentials", "verify-valcred", "cryptovalcred").String,
				Optional:            true,
			},
			"default_signer_cert": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Default signature certificate", "default-signer-cert", "cryptocertificate").String,
				Optional:            true,
			},
		},
	}
}

func (r *B2BCPAReceiverSettingResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *B2BCPAReceiverSettingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.B2BCPAReceiverSetting

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `B2BCPAReceiverSetting`)
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

func (r *B2BCPAReceiverSettingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.B2BCPAReceiverSetting

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
		data.FromBody(ctx, `B2BCPAReceiverSetting`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `B2BCPAReceiverSetting`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *B2BCPAReceiverSettingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.B2BCPAReceiverSetting

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `B2BCPAReceiverSetting`))
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

func (r *B2BCPAReceiverSettingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.B2BCPAReceiverSetting

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
