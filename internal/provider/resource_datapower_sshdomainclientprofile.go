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
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &SSHDomainClientProfileResource{}

func NewSSHDomainClientProfileResource() resource.Resource {
	return &SSHDomainClientProfileResource{}
}

type SSHDomainClientProfileResource struct {
	client *client.DatapowerClient
}

func (r *SSHDomainClientProfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sshdomainclientprofile"
}

func (r *SSHDomainClientProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("SSH domain client profile", "sshdomainclientprofile", "").String,

		Attributes: map[string]schema.Attribute{
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
			"ciphers": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Ciphers", "ciphers", "").AddStringEnum("CHACHA20-POLY1305_AT_OPENSSH.COM", "AES128-CTR", "AES192-CTR", "AES256-CTR", "AES128-GCM_AT_OPENSSH.COM", "AES256-GCM_AT_OPENSSH.COM").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Key exchange algorithms", "kex-alg", "").AddStringEnum("DIFFIE-HELLMAN-GROUP-EXCHANGE-SHA256", "ECDH-SHA2-NISTP256", "ECDH-SHA2-NISTP384", "ECDH-SHA2-NISTP521", "CURVE25519-SHA256_AT_LIBSSH.ORG").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Message authentication codes", "mac-alg", "").AddStringEnum("HMAC-SHA1", "HMAC-SHA2-256", "HMAC-SHA2-512", "UMAC-64_AT_OPENSSH.COM", "UMAC-128_AT_OPENSSH.COM", "HMAC-SHA1-ETM_AT_OPENSSH.COM", "HMAC-SHA2-256-ETM_AT_OPENSSH.COM", "HMAC-SHA2-512-ETM_AT_OPENSSH.COM", "UMAC-64-ETM_AT_OPENSSH.COM", "UMAC-128-ETM_AT_OPENSSH.COM").String,
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
		},
	}
}

func (r *SSHDomainClientProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *SSHDomainClientProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.SSHDomainClientProfile

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `SSHDomainClientProfile`)
	_, err := r.client.Put(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "PUT", err))
		return
	}

	_, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s", "POST", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SSHDomainClientProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.SSHDomainClientProfile

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
		data.FromBody(ctx, `SSHDomainClientProfile`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `SSHDomainClientProfile`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SSHDomainClientProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.SSHDomainClientProfile

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Put(data.GetPath(), data.ToBody(ctx, `SSHDomainClientProfile`))
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

func (r *SSHDomainClientProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	resp.State.RemoveResource(ctx)
}
