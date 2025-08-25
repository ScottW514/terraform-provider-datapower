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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

var _ resource.Resource = &DNSNameServiceResource{}

func NewDNSNameServiceResource() resource.Resource {
	return &DNSNameServiceResource{}
}

type DNSNameServiceResource struct {
	pData *tfutils.ProviderData
}

func (r *DNSNameServiceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dns_name_service"
}

func (r *DNSNameServiceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Configure the DNS client with the DNS servers to contact to resolve hostnames to IP addresses.", "dns", "").String,
		Attributes: map[string]schema.Attribute{
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
			"search_domains": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the list of search domains to resolve partial hostnames.", "search-domain", "").String,
				NestedObject:        models.GetDmSearchDomainResourceSchema(),
				Optional:            true,
			},
			"name_servers": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the list of DNS servers to contact to resolve hostnames. If you define multiple servers, ensure that the sequence to contact the servers is your preferred order.", "name-server", "").String,
				NestedObject:        models.GetDmNameServerResourceSchema(),
				Optional:            true,
			},
			"static_hosts": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the static map of hostnames to IP addresses that do not use DNS resolution. Because the local resolver uses a cache, static hosts do not improve performance.", "static-host", "").String,
				NestedObject:        models.GetDmStaticHostResourceSchema(),
				Optional:            true,
			},
			"ip_preference": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the preferred IP version to resolve hostnames. When a hostname resolves to both IPv4 and IPv6 addresses, this setting determines which version to use.", "ip-preference", "").AddStringEnum("4", "6").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("4", "6"),
				},
			},
			"force_ip_preference": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to restrict DNS queries to the preferred IP version to resolve hostnames. You want to force the IP preference except when both IPv4 and IPv6 addresses are in use. When not forced, the device resolves each hostname by querying A and AAAA records and waiting for both responses or a timeout. Waiting for the response or timeout for both records can introduce unnecessary latency in DNS resolution.", "force-ip-preference", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"load_balance_algorithm": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the load distribution algorithm to resolve hostnames. The default algorithm is first-alive.", "load-balance", "").AddStringEnum("round-robin", "first-alive").AddDefaultValue("first-alive").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("round-robin", "first-alive"),
				},
				Default: stringdefault.StaticString("first-alive"),
			},
			"max_retries": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("For the first alive algorithm, specify the maximum number of resolution attempts to send a query to the list of name servers before an error is returned. By default, an unacknowledged resolution request is attempted 3 times.", "retries", "").AddDefaultValue("2").AddRequiredWhen(models.DNSNameServiceMaxRetriesCondVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					validators.ConditionalRequiredInt64(models.DNSNameServiceMaxRetriesCondVal, validators.Evaluation{}, true),
				},
				Default: int64default.StaticInt64(2),
			},
			"timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("For the first alive algorithm, specify the duration in seconds that the resolver waits for a response from a DNS server. After expiry, the resolver attempts the query to a different DNS server. The default value is 5.", "timeout", "").AddDefaultValue("5").AddRequiredWhen(models.DNSNameServiceTimeoutCondVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					validators.ConditionalRequiredInt64(models.DNSNameServiceTimeoutCondVal, validators.Evaluation{}, true),
				},
				Default: int64default.StaticInt64(5),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *DNSNameServiceResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *DNSNameServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.DNSNameService
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, "default", data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `DNSNameService`)
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

func (r *DNSNameServiceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.DNSNameService
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
		data.FromBody(ctx, `DNSNameService`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `DNSNameService`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DNSNameServiceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.DNSNameService
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, "default", data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath(), data.ToBody(ctx, `DNSNameService`))
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

func (r *DNSNameServiceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.DNSNameService
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, "default", data.DependencyActions, actions.Delete, false)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
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
