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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &QuotaEnforcementServerResource{}

func NewQuotaEnforcementServerResource() resource.Resource {
	return &QuotaEnforcementServerResource{}
}

type QuotaEnforcementServerResource struct {
	client *client.DatapowerClient
}

func (r *QuotaEnforcementServerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_quotaenforcementserver"
}

func (r *QuotaEnforcementServerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Quota Enforcement Server (`default` domain only)", "quota-enforcement-server", "").String,

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
			"password_alias": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Password alias", "password-alias", "passwordalias").String,
				Optional:            true,
			},
			"raid_volume": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Data storage location", "raid-volume", "raidvolume").String,
				Optional:            true,
			},
			"server_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Server port", "server-port", "").AddDefaultValue("16379").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(16379),
			},
			"monitor_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Monitor port", "monitor-port", "").AddDefaultValue("26379").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(26379),
			},
			"enable_peer_group": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Peer group mode", "enable-peer-group", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"enable_ssl": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable TLS", "enable-ssl", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ssl_crypto_key": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Crypto key", "ssl-key", "cryptokey").String,
				Optional:            true,
			},
			"ssl_crypto_certificate": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Certificate", "ssl-cert", "cryptocertificate").String,
				Optional:            true,
			},
			"ip_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("IP address", "ip-address", "").String,
				Optional:            true,
			},
			"peers": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Peers", "peer", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Priority", "priority", "").AddIntegerRange(0, 255).AddDefaultValue("100").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 255),
				},
				Default: int64default.StaticInt64(100),
			},
			"strict_mode": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Strict mode", "strict-mode", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
		},
	}
}

func (r *QuotaEnforcementServerResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *QuotaEnforcementServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.QuotaEnforcementServer

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `QuotaEnforcementServer`)
	_, err := r.client.Put(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "PUT", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *QuotaEnforcementServerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.QuotaEnforcementServer

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
		data.FromBody(ctx, `QuotaEnforcementServer`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `QuotaEnforcementServer`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *QuotaEnforcementServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.QuotaEnforcementServer

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Put(data.GetPath(), data.ToBody(ctx, `QuotaEnforcementServer`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *QuotaEnforcementServerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	resp.State.RemoveResource(ctx)
}
