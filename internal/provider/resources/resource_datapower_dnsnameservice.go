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
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &DNSNameServiceResource{}

func NewDNSNameServiceResource() resource.Resource {
	return &DNSNameServiceResource{}
}

type DNSNameServiceResource struct {
	client *client.DatapowerClient
}

func (r *DNSNameServiceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnsnameservice"
}

func (r *DNSNameServiceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("DNS settings (`default` domain only)", "dns", "").String,
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
			"search_domains": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Search domains", "search-domain", "").String,
				NestedObject:        models.DmSearchDomainResourceSchema,
				Optional:            true,
			},
			"name_servers": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("DNS servers", "name-server", "").String,
				NestedObject:        models.DmNameServerResourceSchema,
				Optional:            true,
			},
			"static_hosts": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Static hosts", "static-host", "").String,
				NestedObject:        models.DmStaticHostResourceSchema,
				Optional:            true,
			},
			"ip_preference": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("IP preference", "ip-preference", "").AddStringEnum("4", "6").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("4", "6"),
				},
			},
			"force_ip_preference": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Force IP preference", "force-ip-preference", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"load_balance_algorithm": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Load distribution algorithm", "load-balance", "").AddStringEnum("round-robin", "first-alive").AddDefaultValue("first-alive").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("round-robin", "first-alive"),
				},
				Default: stringdefault.StaticString("first-alive"),
			},
			"max_retries": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Attempts", "retries", "").AddDefaultValue("2").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(2),
			},
			"timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Timeout", "timeout", "").AddDefaultValue("5").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(5),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *DNSNameServiceResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *DNSNameServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.DNSNameService

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, "default", data.DependencyActions, actions.Create, false)

	body := data.ToBody(ctx, `DNSNameService`)
	_, err := r.client.Put(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "PUT", err))
		return
	}
	actions.PostProcess(ctx, &resp.Diagnostics, r.client, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DNSNameServiceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.DNSNameService

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
		data.FromBody(ctx, `DNSNameService`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `DNSNameService`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DNSNameServiceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.DNSNameService

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, "default", data.DependencyActions, actions.Update, false)
	_, err := r.client.Put(data.GetPath(), data.ToBody(ctx, `DNSNameService`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, r.client, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DNSNameServiceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.DNSNameService

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, "default", data.DependencyActions, actions.Delete, false)
	actions.PostProcess(ctx, &resp.Diagnostics, r.client, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *DNSNameServiceResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.DNSNameService

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
