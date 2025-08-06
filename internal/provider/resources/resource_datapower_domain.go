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
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &DomainResource{}

func NewDomainResource() resource.Resource {
	return &DomainResource{}
}

type DomainResource struct {
	client *client.DatapowerClient
}

func (r *DomainResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_domain"
}

func (r *DomainResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Application domain", "domain", "").String,

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
				MarkdownDescription: tfutils.NewAttributeDescription("Configuration directory", "", "").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"neighbor_domain": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Visible domains", "visible-domain", "domain").String,
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default: listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue("default"),
				})),
			},
			"domain_user": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("CLI user access", "domain-user", "user").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"file_map":       models.GetDmDomainFileMapResourceSchema("File permission to the local: directory", "file-permissions", "", false),
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
				MarkdownDescription: tfutils.NewAttributeDescription("Deployment policy", "deployment-policy", "configdeploymentpolicy").String,
				Optional:            true,
			},
			"deployment_policy_parameters": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Deployment policy variables", "deployment-policy-variables", "deploymentpolicyparametersbinding").String,
				Optional:            true,
			},
			"local_ip_rewrite": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Rewrite local IP addresses", "local-ip-rewrite", "").AddDefaultValue("true").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Global permissions profile", "config-permissions-profile", "accessprofile").String,
				Optional:            true,
			},
			"object_actions": actions.ActionsSchema,
		},
	}
}

func (r *DomainResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *DomainResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.Domain

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, data.ObjectActions, actions.Create)

	body := data.ToBody(ctx, `Domain`)
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

	_, err = r.client.Post("/mgmt/actionqueue/default", "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s", "POST", err))
		return
	}
	getRes, getErr := r.client.Get(data.GetPath())
	if getErr != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object after creation (GET), got error: %s", getErr))
		return
	}
	data.UpdateUnknownFromBody(ctx, `Domain`, getRes)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DomainResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.Domain

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
		data.FromBody(ctx, `Domain`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `Domain`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DomainResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.Domain

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, data.ObjectActions, actions.Update)
	_, err := r.client.Put(data.GetPath(), data.ToBody(ctx, `Domain`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}
	_, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s", "POST", err))
		return
	}

	_, err = r.client.Post("/mgmt/actionqueue/default", "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s", "POST", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DomainResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.Domain

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, data.ObjectActions, actions.Delete)
	// Special case for Application Domains - we need make sure there are no active user sessions before deleting
	auRes, auErr := r.client.Get("/mgmt/status/default/ActiveUsers")
	if auErr != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to get Active Users, got error: %s", auErr))
		return
	}
	if userList := auRes.Get("ActiveUsers"); userList.Exists() {
		for _, session := range userList.Array() {
			if d := session.Get("domain"); d.Exists() && d.Str == data.AppDomain.ValueString() {
				if id := session.Get("session"); id.Exists() {
					r.client.Post("/mgmt/actionqueue/default", fmt.Sprintf("{\"Disconnect\": {\"id\": %d}}", id.Int()))
				}
			}
		}
	}
	_, err := r.client.Delete(data.GetPath())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}

	_, err = r.client.Post("/mgmt/actionqueue/default", "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s", "POST", err))
		return
	}

	resp.State.RemoveResource(ctx)
}
