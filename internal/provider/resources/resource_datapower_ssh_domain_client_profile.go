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

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &SSHDomainClientProfileResource{}

func NewSSHDomainClientProfileResource() resource.Resource {
	return &SSHDomainClientProfileResource{}
}

type SSHDomainClientProfileResource struct {
	pData *tfutils.ProviderData
}

func (r *SSHDomainClientProfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssh_domain_client_profile"
}

func (r *SSHDomainClientProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("The SSH domain client profile defines the cipher suites.", "sshdomainclientprofile", "").String,
		Attributes: map[string]schema.Attribute{
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
			"ciphers": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the SSH cipher suites to support.", "ciphers", "").AddStringEnum("CHACHA20-POLY1305_AT_OPENSSH.COM", "AES128-CTR", "AES192-CTR", "AES256-CTR", "AES128-GCM_AT_OPENSSH.COM", "AES256-GCM_AT_OPENSSH.COM").String,
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(stringvalidator.OneOf("CHACHA20-POLY1305_AT_OPENSSH.COM", "AES128-CTR", "AES192-CTR", "AES256-CTR", "AES128-GCM_AT_OPENSSH.COM", "AES256-GCM_AT_OPENSSH.COM")),
				},
				Default: listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue("AES128-CTR"),
					types.StringValue("AES192-CTR"),
					types.StringValue("AES256-CTR"),
					types.StringValue("AES128-GCM_AT_OPENSSH.COM"),
					types.StringValue("AES256-GCM_AT_OPENSSH.COM"),
				})),
			},
			"kex_alg": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the key exchange (KEX) algorithms to support.", "kex-alg", "").AddStringEnum("DIFFIE-HELLMAN-GROUP-EXCHANGE-SHA256", "ECDH-SHA2-NISTP256", "ECDH-SHA2-NISTP384", "ECDH-SHA2-NISTP521", "CURVE25519-SHA256_AT_LIBSSH.ORG").String,
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(stringvalidator.OneOf("DIFFIE-HELLMAN-GROUP-EXCHANGE-SHA256", "ECDH-SHA2-NISTP256", "ECDH-SHA2-NISTP384", "ECDH-SHA2-NISTP521", "CURVE25519-SHA256_AT_LIBSSH.ORG")),
				},
				Default: listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue("CURVE25519-SHA256_AT_LIBSSH.ORG"),
					types.StringValue("ECDH-SHA2-NISTP256"),
					types.StringValue("ECDH-SHA2-NISTP384"),
					types.StringValue("ECDH-SHA2-NISTP521"),
					types.StringValue("DIFFIE-HELLMAN-GROUP-EXCHANGE-SHA256"),
				})),
			},
			"mac_alg": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the message authentication codes (MAC) to support.", "mac-alg", "").AddStringEnum("HMAC-SHA1", "HMAC-SHA2-256", "HMAC-SHA2-512", "UMAC-64_AT_OPENSSH.COM", "UMAC-128_AT_OPENSSH.COM", "HMAC-SHA1-ETM_AT_OPENSSH.COM", "HMAC-SHA2-256-ETM_AT_OPENSSH.COM", "HMAC-SHA2-512-ETM_AT_OPENSSH.COM", "UMAC-64-ETM_AT_OPENSSH.COM", "UMAC-128-ETM_AT_OPENSSH.COM").String,
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(stringvalidator.OneOf("HMAC-SHA1", "HMAC-SHA2-256", "HMAC-SHA2-512", "UMAC-64_AT_OPENSSH.COM", "UMAC-128_AT_OPENSSH.COM", "HMAC-SHA1-ETM_AT_OPENSSH.COM", "HMAC-SHA2-256-ETM_AT_OPENSSH.COM", "HMAC-SHA2-512-ETM_AT_OPENSSH.COM", "UMAC-64-ETM_AT_OPENSSH.COM", "UMAC-128-ETM_AT_OPENSSH.COM")),
				},
				Default: listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue("UMAC-64_AT_OPENSSH.COM"),
					types.StringValue("UMAC-128_AT_OPENSSH.COM"),
					types.StringValue("HMAC-SHA2-256"),
					types.StringValue("HMAC-SHA2-512"),
					types.StringValue("HMAC-SHA1"),
				})),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *SSHDomainClientProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *SSHDomainClientProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SSHDomainClientProfile
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `SSHDomainClientProfile`)
	_, err := r.pData.Client.Put(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "PUT", err))
		return
	}
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SSHDomainClientProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SSHDomainClientProfile
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

	if data.IsNull() {
		// Import
		data.FromBody(ctx, `SSHDomainClientProfile`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `SSHDomainClientProfile`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SSHDomainClientProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SSHDomainClientProfile
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath(), data.ToBody(ctx, `SSHDomainClientProfile`))
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

func (r *SSHDomainClientProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SSHDomainClientProfile
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *SSHDomainClientProfileResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.SSHDomainClientProfile

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
