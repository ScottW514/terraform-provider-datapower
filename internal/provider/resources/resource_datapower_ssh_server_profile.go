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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var (
	_ resource.Resource                = &SSHServerProfileResource{}
	_ resource.ResourceWithImportState = &SSHServerProfileResource{}
)

func NewSSHServerProfileResource() resource.Resource {
	return &SSHServerProfileResource{}
}

type SSHServerProfileResource struct {
	pData *tfutils.ProviderData
}

func (r *SSHServerProfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssh_server_profile"
}

func (r *SSHServerProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("The SSH domain client profile defines the cipher suites.", "sshserverprofile", "").String,
		Attributes: map[string]schema.Attribute{
			"provider_target": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Target host for this resource. If not set, provider will use the top level settings.", "", "").String,
				Optional:            true,
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
			"enabled": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>The administrative state of the configuration.</p><ul><li>To make active, set to enabled.</li><li>To make inactive, set to disabled.</li></ul>", "admin-state", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"ciphers": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the SSH cipher suites to support.", "ciphers", "").AddStringEnum("CHACHA20-POLY1305_AT_OPENSSH.COM", "AES128-CTR", "AES192-CTR", "AES256-CTR", "AES128-GCM_AT_OPENSSH.COM", "AES256-GCM_AT_OPENSSH.COM").String,
				ElementType:         types.StringType,
				Computed:            true,
				Optional:            true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOf("CHACHA20-POLY1305_AT_OPENSSH.COM", "AES128-CTR", "AES192-CTR", "AES256-CTR", "AES128-GCM_AT_OPENSSH.COM", "AES256-GCM_AT_OPENSSH.COM"),
					),
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
				Computed:            true,
				Optional:            true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOf("DIFFIE-HELLMAN-GROUP-EXCHANGE-SHA256", "ECDH-SHA2-NISTP256", "ECDH-SHA2-NISTP384", "ECDH-SHA2-NISTP521", "CURVE25519-SHA256_AT_LIBSSH.ORG"),
					),
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
				Computed:            true,
				Optional:            true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOf("HMAC-SHA1", "HMAC-SHA2-256", "HMAC-SHA2-512", "UMAC-64_AT_OPENSSH.COM", "UMAC-128_AT_OPENSSH.COM", "HMAC-SHA1-ETM_AT_OPENSSH.COM", "HMAC-SHA2-256-ETM_AT_OPENSSH.COM", "HMAC-SHA2-512-ETM_AT_OPENSSH.COM", "UMAC-64-ETM_AT_OPENSSH.COM", "UMAC-128-ETM_AT_OPENSSH.COM"),
					),
				},
				Default: listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue("UMAC-64-ETM_AT_OPENSSH.COM"),
					types.StringValue("UMAC-128-ETM_AT_OPENSSH.COM"),
					types.StringValue("HMAC-SHA2-256"),
					types.StringValue("HMAC-SHA2-512"),
					types.StringValue("HMAC-SHA1"),
				})),
			},
			"send_pre_auth_msg": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to include a banner message during SSH preauthentication. The banner message contains the text to display to users before the login prompt.", "send-preauth-msg", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"pre_auth_msg": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the banner message to display to users before the login prompt. White space characters ( <tt>\\n</tt> , <tt>\\t</tt> ) are recognized and escaped. The banner message must be at least one character. The banner message is truncated at 4096 characters.", "preauth-msg", "").AddNotValidWhen(models.SSHServerProfilePreAuthMsgIgnoreVal.String()).String,
				Optional:            true,
			},
			"host_key_alg":       models.GetDmHostKeyAlgorithmsResourceSchema("Host key algorithms", "host-key-alg", "", false),
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *SSHServerProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *SSHServerProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SSHServerProfile
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !r.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), r.pData.Client.GetTargetNames()))
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `SSHServerProfile`)
	_, err := r.pData.Client.Put(data.GetPath(), body, data.ProviderTarget)
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			_, err := r.pData.Client.Put(data.GetPath(), body, data.ProviderTarget)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Resource already exists. Failed to update resource, got error: %s", err))
				return
			}
			resp.Diagnostics.AddWarning("Warning", "Resource already exists. Existing resource was updated.")
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create resource, got error: %s", err))
			return
		}
	}
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}

	tfutils.SaveDomain(ctx, &resp.Diagnostics, r.pData.Client, data.AppDomain, data.ProviderTarget)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SSHServerProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SSHServerProfile
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !r.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), r.pData.Client.GetTargetNames()))
		return
	}
	res, err := r.pData.Client.Get(data.GetPath(), data.ProviderTarget)
	if err != nil {
		if strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400") {
			resp.State.RemoveResource(ctx)
			return
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.State.RemoveResource(ctx)
			}
			return
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
			return
		}
	}

	data.UpdateFromBody(ctx, `SSHServerProfile`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SSHServerProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SSHServerProfile
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !r.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), r.pData.Client.GetTargetNames()))
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath(), data.ToBody(ctx, `SSHServerProfile`), data.ProviderTarget)
	if err != nil {
		if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
			return
		}
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Update, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}

	tfutils.SaveDomain(ctx, &resp.Diagnostics, r.pData.Client, data.AppDomain, data.ProviderTarget)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SSHServerProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SSHServerProfile
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !r.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), r.pData.Client.GetTargetNames()))
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}
	data.ToDefault()
	_, err := r.pData.Client.Put(data.GetPath(), data.ToBody(ctx, `SSHServerProfile`), data.ProviderTarget)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to restore singleton to default, got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}

	tfutils.SaveDomain(ctx, &resp.Diagnostics, r.pData.Client, data.AppDomain, data.ProviderTarget)

	resp.State.RemoveResource(ctx)
}

func (r *SSHServerProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()
	appDomain := req.ID
	if !regexp.MustCompile("^[a-zA-Z0-9_-]+$").MatchString(appDomain) || len(appDomain) < 1 || len(appDomain) > 128 {
		resp.Diagnostics.AddError("Invalid Application Domain", "Application domain must be 1-128 characters and match pattern ^[a-zA-Z0-9_-]+$")
		return
	}

	var data models.SSHServerProfile
	data.AppDomain = types.StringValue(appDomain)
	res, err := r.pData.Client.Get(data.GetPath(), data.ProviderTarget)
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("Resource was not found, got error: %s", err))
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		}
		return
	}

	data.FromBody(ctx, `SSHServerProfile`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
