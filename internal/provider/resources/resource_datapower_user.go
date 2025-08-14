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
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
)

var _ resource.ResourceWithModifyPlan = &UserResource{}

func NewUserResource() resource.Resource {
	return &UserResource{}
}

type UserResource struct {
	pData *tfutils.ProviderData
}

func (r *UserResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

func (r *UserResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Create or edit a local user account. When you modify a local user account, the changes do not affect remote user accounts in a role-based management XML information file. <p>An administrator can change the password for the account. An administrator can force the account owner to change the password during the next log in attempt.</p>", "user", "").String,
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
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the password for the account. The password must comply to the password policy in RBM settings.", "password", "").String,
				Required:            true,
				WriteOnly:           true,
				Sensitive:           true,
			},
			"password_update": schema.BoolAttribute{
				MarkdownDescription: "Set to true by provider if the WRITE ONLY value needs to be updated, otherwise provider will force this to false.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				DeprecationMessage:  "This attribute is for INTERNAL PROVIDER USE. Set values are ignored.",
			},
			"access_level": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Access level", "access-level", "").AddStringEnum("none", "privileged", "group-defined", "technician", "expired", "config-sequence").AddDefaultValue("group-defined").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("none", "privileged", "group-defined", "technician", "expired", "config-sequence"),
				},
				Default: stringdefault.StaticString("group-defined"),
			},
			"group_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the user group for the user account. Each user group has an access profile to manage access rights. A user account inherits access rights from its user group.", "group", "usergroup").String,
				Optional:            true,
			},
			"snmp_creds": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify SNMP V3 credentials for the user account. SNMP V3 users are a type of user account with SNMP V3 credentials. This type of user account creates an account and adds SNMP V3 credentials. Each account can have multiple SNMP V3 credentials, one for each SNMP V3 engine that is identified by an engine ID value. The current implementation supports an SNMP V3 credential for the local engine ID only. Therefore, define only one SNMP V3 credential for each account. <p>SNMP V3 requests can use authentication with the <tt>AuthKey</tt> and can use privacy (encryption and decryption) with the <tt>PrivKey</tt> . The use of these keys is between the user and the local SNMP engine. <ul><li>The <tt>AuthKey</tt> provides data integrity and data origin authentication.</li><li>The <tt>PrivKey</tt> provides data encryption and decryption.</li></ul></p><p>The <tt>AuthKey</tt> and the <tt>PrivKey</tt> can be explicit values or be generated by the DataPower Gateway. Specifying explicit values is useful when the key is created on another system. <ul><li>When the authentication protocol is MD5, enter the hex representation of the 16-byte key.</li><li>When the authentication protocol is SHA, enter the hex representation of the 20-byte key.</li><li>When the DataPower Gateway generates and stores the appropriate length key, enter a plaintext string that is at least 8 characters long as the hash.</li></ul></p>", "snmp-cred", "").String,
				NestedObject:        models.DmSnmpCredResourceSchema,
				Optional:            true,
			},
			"hashed_snmp_creds": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "snmp-cred-hashed", "").String,
				NestedObject:        models.DmSnmpCredMaskedResourceSchema,
				Optional:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *UserResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *UserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.User
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, "default", data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	var config models.User

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	data.Password = config.Password
	resp.Diagnostics.Append(resp.Private.SetKey(ctx, "User.Password", []byte("{\"value\": \""+tfutils.GenerateHash(config.Password.ValueString())+"\"}"))...)

	body := data.ToBody(ctx, `User`)
	_, err := r.pData.Client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "POST", err))
		return
	}
	getRes, getErr := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if getErr != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object after creation (GET), got error: %s", getErr))
		return
	}
	data.UpdateUnknownFromBody(ctx, `User`, getRes)
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *UserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.User
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
		data.FromBody(ctx, `User`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `User`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *UserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data, config models.User
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, "default", data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.PasswordUpdate.IsUnknown() {
		data.Password = config.Password
		data.PasswordUpdate = types.BoolValue(false)
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, "User.Password", []byte("{\"value\": \""+tfutils.GenerateHash(config.Password.ValueString())+"\"}"))...)
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `User`))
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

func (r *UserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.User
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, "default", data.DependencyActions, actions.Delete, false)
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
func (r *UserResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}
	var data, config models.User

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var val []byte
	var diags diag.Diagnostics
	val, diags = req.Private.GetKey(ctx, "User.Password")
	resp.Diagnostics.Append(diags...)
	if val != nil {
		if hash := gjson.GetBytes(val, "value"); hash.Exists() && !tfutils.VerifyHash(config.Password.ValueString(), hash.String()) {
			data.PasswordUpdate = types.BoolUnknown()
		}
	}

	resp.Diagnostics.Append(resp.Plan.Set(ctx, &data)...)
}

func (r *UserResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.User

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
