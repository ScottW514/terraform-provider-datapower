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

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

var _ resource.Resource = &NFSDynamicMountsResource{}

func NewNFSDynamicMountsResource() resource.Resource {
	return &NFSDynamicMountsResource{}
}

type NFSDynamicMountsResource struct {
	pData *tfutils.ProviderData
}

func (r *NFSDynamicMountsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nfs_dynamic_mounts"
}

func (r *NFSDynamicMountsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Configure parameters of dynamic NFS mounts for dpnfs URL calls. These mounts support URL access in the form <tt>dpnfs://&lt;host>/&lt;path>/&lt;file></tt> . The system automatically mounts any dynamic mounts. Dynamic mounts remain mounted until the inactivity timer elapses.", "nfs-dynamic-mounts", "").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("<p>The administrative state of the configuration.</p><ul><li>To make active, set to enabled.</li><li>To make inactive, set to disabled.</li></ul>", "admin-state", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the preferred NFS protocol version. Enter a value in the range 2 - 4. The default value is 3. <ul><li>If version 3 and the server only implements version 2, the client falls back to version 2.</li><li>If version 4, the remote export paths are different and prevents fallback.</li></ul>", "version", "").AddIntegerRange(2, 4).AddDefaultValue("3").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(2, 4),
				},
				Default: int64default.StaticInt64(3),
			},
			"transport": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the transport protocol. The default transport protocol is TCP.", "transport", "").AddStringEnum("tcp", "udp").AddDefaultValue("tcp").AddNotValidWhen(models.NFSDynamicMountsTransportIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("tcp", "udp"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.NFSDynamicMountsTransportIgnoreVal, true),
				},
				Default: stringdefault.StaticString("tcp"),
			},
			"mount_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the type of NFS mount. The default mount type is a hard mount.", "mount-type", "").AddStringEnum("hard", "soft").AddDefaultValue("hard").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("hard", "soft"),
				},
				Default: stringdefault.StaticString("hard"),
			},
			"read_only": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the mount is read-only. By default, the mount is not read-only.", "read-only", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"read_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the size in bytes for NFS read operations. Enter a value in the range 1024 - 32768. The default value is 4096.", "rsize", "").AddIntegerRange(1024, 32768).AddDefaultValue("4096").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1024, 32768),
				},
				Default: int64default.StaticInt64(4096),
			},
			"write_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the size in bytes for NFS write operations. Enter a value in the range 1024 - 32768. The default value is 4096.", "wsize", "").AddIntegerRange(1024, 32768).AddDefaultValue("4096").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1024, 32768),
				},
				Default: int64default.StaticInt64(4096),
			},
			"timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the time in tenths of seconds until the first retransmission on RPC times out. Enter a value in the range 1 - 600. The default value is 7.", "timeo", "").AddIntegerRange(1, 600).AddDefaultValue("7").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 600),
				},
				Default: int64default.StaticInt64(7),
			},
			"retransmissions": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the number of minor RPC timeouts and retransmissions until a major timeout. Enter a value in the range 1 - 60. The default value is 3.", "retrans", "").AddIntegerRange(1, 60).AddDefaultValue("3").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 60),
				},
				Default: int64default.StaticInt64(3),
			},
			"idle_unmount_seconds": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the inactivity duration in seconds to wait before the mount is unmounted. The default value is 900. The value of 0 disables the timer.", "inactivity-timeout", "").AddDefaultValue("900").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(900),
			},
			"mount_timeout_seconds": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration in seconds to attempt to mount a dynamic mount. When the timer elapses, related file open operations fail.", "mount-timeout", "").AddIntegerRange(10, 240).AddDefaultValue("30").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(10, 240),
				},
				Default: int64default.StaticInt64(30),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *NFSDynamicMountsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *NFSDynamicMountsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.NFSDynamicMounts
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `NFSDynamicMounts`)
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

func (r *NFSDynamicMountsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.NFSDynamicMounts
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
		data.FromBody(ctx, `NFSDynamicMounts`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `NFSDynamicMounts`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NFSDynamicMountsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.NFSDynamicMounts
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath(), data.ToBody(ctx, `NFSDynamicMounts`))
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

func (r *NFSDynamicMountsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.NFSDynamicMounts
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

func (r *NFSDynamicMountsResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.NFSDynamicMounts

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
