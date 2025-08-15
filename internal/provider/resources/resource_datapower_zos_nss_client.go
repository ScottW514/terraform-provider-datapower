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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &ZosNSSClientResource{}

func NewZosNSSClientResource() resource.Resource {
	return &ZosNSSClientResource{}
}

type ZosNSSClientResource struct {
	pData *tfutils.ProviderData
}

func (r *ZosNSSClientResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zos_nss_client"
}

func (r *ZosNSSClientResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("The z/OS NSS client object provides the parameters for authentication with SAF on a z/OS Communications Server.", "zos-nss", "").String,
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
					modifiers.ImmutableAfterSet(),
				},
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("A descriptive summary for the configuration.", "summary", "").String,
				Optional:            true,
			},
			"remote_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies IP address or host name of the NSS server. In conjunction with the Remote Port, identifies the host and listening port of the NSS server. The NSS server must have the XMLAppliance discipline support enabled.", "host", "").String,
				Required:            true,
			},
			"remote_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies NSS server port. In conjunction with the Remote Address, identifies the host and listening port of the NSS server.", "port", "").String,
				Required:            true,
			},
			"client_id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specifies the client ID to be used for registration with the NSS server. Minimum length is 1. Maximum length is 24.</p><p>Valid characters are:</p><p><ul><li>a through z</li><li>A through Z</li><li>0 through 9</li><li>_ (underscore)</li><li>- (dash)</li></ul></p><p>Embedded spaces are invalid.</p><p>The Client ID identifies the client ID to register the appliance with the NSS server. The NSS client ID is a unique string used by the NSS Server to track clients.</p><p>The Client ID does not have to correspond to any preexisting object. It is provided to the server at the time of registration. If another client attempts to register with the same client ID to the same NSS Server, the NSS server will send a heartbeat to the first client. If the first client responds to the heartbeat, the second client's registration will be rejected. If the first client does not respond, the connect to the first client will be severed and the second client will be registered.</p>", "client-id", "").String,
				Required:            true,
			},
			"system_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specifies a name for the NSS client. Minimum length is 1. Maximum length is 8.</p><p>Valid characters are:</p><p><ul><li>a through z</li><li>A through Z</li><li>0 through 9</li><li>_ (underscore)</li><li>- (dash)</li></ul></p><p>Embedded spaces are invalid.</p><p>The System Name identifies the NSS client to the NSS server. NSS server commands identify NSS clients by system name in the output when displaying information for connected NSS clients.</p>", "system-name", "").String,
				Required:            true,
			},
			"user_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specifies a user name to use to authenticate to the NSS server. Minimum length is 1. Maximum length is 8.</p><p>Valid characters are:</p><p><ul><li>a through z</li><li>A through Z</li><li>0 through 9</li><li>_ (underscore)</li><li>- (dash)</li></ul></p><p>Embedded spaces are invalid.</p><p>The user name must match an existing user ID on the NSS Server.</p>", "user-name", "").String,
				Required:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specifies the password to use to authenticate to the NSS server. Minimum length is 1. Maximum length is 8.</p><p>Valid characters are:</p><p><ul><li>a through z</li><li>A through Z</li><li>0 through 9</li><li>_ (underscore)</li><li>- (dash)</li></ul></p><p>Embedded spaces are invalid.</p><p>The Password specifies the password to use in conjunction with the value provided by the User Name.</p>", "password", "").String,
				Optional:            true,
				Sensitive:           true,
			},
			"password_alias": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specifies the password alias of the password to use to authenticate to the NSS server.</p><p>The associated password is used in conjunction with the value provided by the User Name.</p>", "password-alias", "password_alias").String,
				Optional:            true,
			},
			"ssl_client_config_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The TLS profile type to secure connections between the DataPower Gateway and the NSS server.", "ssl-client-type", "").AddStringEnum("client").AddDefaultValue("client").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("client"),
				},
				Default: stringdefault.StaticString("client"),
			},
			"ssl_client": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The TLS client profile to secure connections between the DataPower Gateway and the NSS server.", "ssl-client", "ssl_client_profile").String,
				Optional:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *ZosNSSClientResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *ZosNSSClientResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.ZosNSSClient
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `ZosNSSClient`)
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

func (r *ZosNSSClientResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.ZosNSSClient
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
		data.FromBody(ctx, `ZosNSSClient`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `ZosNSSClient`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZosNSSClientResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.ZosNSSClient
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `ZosNSSClient`))
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

func (r *ZosNSSClientResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.ZosNSSClient
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

func (r *ZosNSSClientResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.ZosNSSClient

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
