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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

var _ resource.Resource = &B2BCPACollaborationResource{}

func NewB2BCPACollaborationResource() resource.Resource {
	return &B2BCPACollaborationResource{}
}

type B2BCPACollaborationResource struct {
	pData *tfutils.ProviderData
}

func (r *B2BCPACollaborationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_b2b_cpa_collaboration"
}

func (r *B2BCPACollaborationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("B2B CPA collaboration", "b2b-cpa-collaboration", "").String,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name of the object. Must be unique among object types in application domain.", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
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
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
				PlanModifiers: []planmodifier.String{
					modifiers.ImmutableAfterSet(),
				},
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"internal_role": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the authorized role of the internal partner in a business collaboration service. Each role is authorized for specific actions. For example, a <tt>Buyer</tt> role has the authority for purchasing actions.", "internal-role", "").AddRequiredWhen(models.B2BCPACollaborationInternalRoleCondVal.String()).AddNotValidWhen(models.B2BCPACollaborationInternalRoleIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.B2BCPACollaborationInternalRoleCondVal, models.B2BCPACollaborationInternalRoleIgnoreVal, false),
				},
			},
			"external_role": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the authorized role of the external partner in a business collaboration service. Each role is authorized for specific actions. For example, a <tt>Supplier</tt> role has the authority for selling actions.", "external-role", "").AddRequiredWhen(models.B2BCPACollaborationExternalRoleCondVal.String()).AddNotValidWhen(models.B2BCPACollaborationExternalRoleIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.B2BCPACollaborationExternalRoleCondVal, models.B2BCPACollaborationExternalRoleIgnoreVal, false),
				},
			},
			"process_specification": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the process specification document that defines the interactions between the internal and external partners. For example, <tt>http://www.rosettanet.org/processes/3A4</tt> .", "process-spec", "").String,
				Optional:            true,
			},
			"service": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the value of the service that acts on the message. The value is used to specify and identify the value of the Service element in the outbound and inbound ebMS message header. The service is one of the following types. <ul><li>A business collaboration service for exchanging business messages.</li><li>An MSH signal service for exchanging MSH signals.</li></ul><p>The value of <tt>urn:oasis:names:tc:ebxml-msg:service;</tt> is an MSH signal service. Any other value represents a business collaboration service.</p>", "service", "").String,
				Required:            true,
			},
			"service_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the value of the service type. If you specify the type, the value is present in the type attribute of the <tt>Service</tt> element within the message to be sent. If the type is empty, the value of the <tt>Service</tt> element must be a URI.", "service-type", "").AddNotValidWhen(models.B2BCPACollaborationServiceTypeIgnoreVal.String()).String,
				Optional:            true,
			},
			"sender_msh_setting": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the default MSH signal sender to send ebMS MSH signals to send MSH level signals that include <tt>Acknowledgment</tt> , <tt>Error</tt> , <tt>StatusRequest</tt> , <tt>StatusResponse</tt> , <tt>Ping</tt> , and <tt>Pong</tt> .", "sender-msh-setting", "b2b_cpa_sender_setting").AddRequiredWhen(models.B2BCPACollaborationSenderMshSettingCondVal.String()).AddNotValidWhen(models.B2BCPACollaborationSenderMshSettingIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.B2BCPACollaborationSenderMshSettingCondVal, models.B2BCPACollaborationSenderMshSettingIgnoreVal, false),
				},
			},
			"receiver_msh_setting": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the default MSH signal receiver to receive ebMS MSH signals to receive MSH level signals that include <tt>Acknowledgment</tt> , <tt>Error</tt> , <tt>StatusRequest</tt> , <tt>StatusResponse</tt> , <tt>Ping</tt> , and <tt>Pong</tt> .", "receiver-msh-setting", "b2b_cpa_receiver_setting").AddRequiredWhen(models.B2BCPACollaborationReceiverMshSettingCondVal.String()).AddNotValidWhen(models.B2BCPACollaborationReceiverMshSettingIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.B2BCPACollaborationReceiverMshSettingCondVal, models.B2BCPACollaborationReceiverMshSettingIgnoreVal, false),
				},
			},
			"actions": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify CPA actions to bind. For a business collaboration, each action entry identifies a business message that a party can send or receive. For a collaboration of MSH level signal, the action overrides the sending or receiving behaviors of the default sender setting or default receiver setting.", "action", "").AddRequiredWhen(models.B2BCPACollaborationActionsCondVal.String()).String,
				NestedObject:        models.GetDmCPACollaborationActionResourceSchema(),
				Required:            true,
				Validators: []validator.List{
					validators.ConditionalRequiredList(models.B2BCPACollaborationActionsCondVal, validators.Evaluation{}, false),
				},
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *B2BCPACollaborationResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *B2BCPACollaborationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.B2BCPACollaboration
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `B2BCPACollaboration`)
	_, err := r.pData.Client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "POST", err))
		return
	}
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *B2BCPACollaborationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.B2BCPACollaboration
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && (strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}

	if data.IsNull() {
		// Import
		data.FromBody(ctx, `B2BCPACollaboration`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `B2BCPACollaboration`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *B2BCPACollaborationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.B2BCPACollaboration
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `B2BCPACollaboration`))
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

func (r *B2BCPACollaborationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.B2BCPACollaboration
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *B2BCPACollaborationResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.B2BCPACollaboration

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
