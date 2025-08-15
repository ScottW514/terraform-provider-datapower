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
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &DomainResource{}
var _ resource.ResourceWithValidateConfig = &DomainResource{}

func NewDomainResource() resource.Resource {
	return &DomainResource{}
}

type DomainResource struct {
	pData *tfutils.ProviderData
}

func (r *DomainResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_domain"
}

func (r *DomainResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("An application domain contains the resources that support DataPower services. The DataPower Gateway supports multiple domains. Domains can share read access to files in their <tt>local:</tt> directory. All domains share the contents of the <tt>store:</tt> directory. <p>After a user logs in to a domain, everything the user does applies to only this domain.</p><p>Except for the <tt>default</tt> domain, all domains can be restarted independently. For the <tt>default</tt> domain, you must restart the DataPower Gateway. When an domain or the DataPower Gateway is restarted, the persisted configuration is used. The persisted configuration can differ from the running configuration.</p><p>The configuration of a domain can be locally stored or can be retrieved from a remote server. The use of a remote configuration file enables centralized management of domains.</p>", "domain", "").AddActions("quiesce", "restart").String,
		Attributes: map[string]schema.Attribute{
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
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"config_dir": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify where the configuration file for this domain is stored. This property is read-only because it is configured in domain settings.", "", "").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"neighbor_domain": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify which domains have their <tt>local:</tt> directory visible to this domain. <p>References to visible domains are explicit, not bidirectional. If domain <tt>domainB</tt> is made visible to domain <tt>domainA</tt> , the following conditions apply.</p><ul><li>Domain <tt>domainA</tt> has read-only access to the <tt>local:</tt> directory of domain <tt>domainB</tt> .</li><li>Domain <tt>domainB</tt> cannot see domain <tt>domainA</tt> .</li></ul><p>In this case, you cannot make domain <tt>domainA</tt> visible to domain <tt>domainB</tt> . References to visible domains cannot be circular.</p>", "visible-domain", "domain").String,
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default: listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue("default"),
				})),
			},
			"domain_user": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the set of CLI users who can access this domain. These users can log into this domain through the CLI. This setting can be superseded by an existing access policy of the user.", "domain-user", "user").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"file_map":       models.GetDmDomainFileMapResourceSchema("Specify the file permissions to apply to the <tt>local:</tt> directory. When access permissions are defined and with role-based management, users are granted the lesser privilege.", "file-permissions", "", false),
			"monitoring_map": models.GetDmDomainMonitoringMapResourceSchema("File-monitoring of the local: directory", "file-monitoring", "", false),
			"config_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Configuration mode", "config-mode", "").AddStringEnum("local", "import").AddDefaultValue("local").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("local", "import"),
				},
				Default: stringdefault.StaticString("local"),
			},
			"import_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Import URL", "import-url", "").String,
				Optional:            true,
			},
			"import_format": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Import format", "import-format", "").AddStringEnum("ZIP", "XML").AddDefaultValue("ZIP").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("ZIP", "XML"),
				},
				Default: stringdefault.StaticString("ZIP"),
			},
			"deployment_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Deployment policy", "deployment-policy", "config_deployment_policy").String,
				Optional:            true,
			},
			"deployment_policy_parameters": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Deployment policy variables", "deployment-policy-variables", "deployment_policy_parameters_binding").String,
				Optional:            true,
			},
			"local_ip_rewrite": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to rewrite local IP addresses during import. When enabled, local IP addresses in the import package are rewritten to match the local IP address on the DataPower Gateway. In other words, a service that binds to <tt>eth10</tt> in the import package is rewritten to bind to the local IP address of <tt>eth10</tt> on the DataPower Gateway.", "local-ip-rewrite", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"max_chkpoints": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Checkpoint limit", "maxchkpoints", "").AddIntegerRange(1, 5).AddDefaultValue("3").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 5),
				},
				Default: int64default.StaticInt64(3),
			},
			"config_permissions_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Configuration permissions mode", "config-permissions-mode", "").AddStringEnum("scope-domain", "global-profile", "specific-profile").AddDefaultValue("scope-domain").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("scope-domain", "global-profile", "specific-profile"),
				},
				Default: stringdefault.StaticString("scope-domain"),
			},
			"config_permissions_profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Global permissions profile", "config-permissions-profile", "access_profile").String,
				Optional:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *DomainResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *DomainResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.Domain
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, true)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `Domain`)
	_, err := r.pData.Client.Put(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "PUT", err))
		return
	}
	getRes, getErr := r.pData.Client.Get(data.GetPath())
	if getErr != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object after creation (GET), got error: %s", getErr))
		return
	}
	data.UpdateUnknownFromBody(ctx, `Domain`, getRes)
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DomainResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.Domain
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
		data.FromBody(ctx, `Domain`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `Domain`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DomainResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.Domain
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, true)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath(), data.ToBody(ctx, `Domain`))
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

func (r *DomainResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.Domain
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, true)
	if resp.Diagnostics.HasError() {
		return
	}

	// Special case for Application Domains - we need make sure there are no active user sessions before deleting
	auRes, auErr := r.pData.Client.Get("/mgmt/status/default/ActiveUsers")
	if auErr != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to get Active Users, got error: %s", auErr))
		return
	}
	if userList := auRes.Get("ActiveUsers"); userList.Exists() {
		for _, session := range userList.Array() {
			if d := session.Get("domain"); d.Exists() && d.Str == data.AppDomain.ValueString() {
				if id := session.Get("session"); id.Exists() {
					r.pData.Client.Post("/mgmt/actionqueue/default", fmt.Sprintf("{\"Disconnect\": {\"id\": %d}}", id.Int()))
				}
			}
		}
	}
	_, err := r.pData.Client.Delete(data.GetPath())
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

func (r *DomainResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.Domain

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
