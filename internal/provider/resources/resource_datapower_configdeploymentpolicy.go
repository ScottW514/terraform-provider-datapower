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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &ConfigDeploymentPolicyResource{}

func NewConfigDeploymentPolicyResource() resource.Resource {
	return &ConfigDeploymentPolicyResource{}
}

type ConfigDeploymentPolicyResource struct {
	pData *tfutils.ProviderData
}

func (r *ConfigDeploymentPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_configdeploymentpolicy"
}

func (r *ConfigDeploymentPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("A deployment policy contains a set of rules that are applied during a configuration import. A deployment policy can be used to accept, filter, or modify configuration during import.", "deployment-policy", "").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"accepted_config": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Matching configuration is accepted during import. To create a match statement, type a correctly formatted resource match in the horizontal text box or select Build. Selecting Build displays the Configuration Match Builder in a popup window. <p>A match statement takes the following general form: <br/><i>Addr</i> / <i>Domain</i> / <i>Resource</i> [? <i>Name=resource-name</i> &amp; <i>Property=property-name</i> &amp; <i>Value=property-value</i> ]</p><table><tr><td valign=\"top\">Addr</td><td>Device Address. Specifies IP address or host alias. The value (*) matches all IP addresses.</td></tr><tr><td valign=\"top\">Domain</td><td>Application Domain. The name of the application domain. The value (*) matches all domains.</td></tr><tr><td valign=\"top\">Resource</td><td>Resource Type. The name of the resource type. The value (*) matches all resource types.</td></tr><tr><td valign=\"top\">Name=resource-name</td><td>An additional specification field, such as \"Name\". Limits the match statement to resources of the specified name. Use a PCRE to select groups of resource instances. For example, \"Name=foo.*\" would match all resources with names that start with \"foo\".</td></tr><tr><td valign=\"top\">Property=property-name</td><td>Property Name. The name of the configuration property. Limits the match statement to resources of the specified property. If change specified, set property-name to null string.</td></tr><tr><td valign=\"top\">Value=property-value</td><td>Property Value. Specifies the value for the configuration property. This property limits the match statement to resources with the specified property value.</td></tr></table>", "accept", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"filtered_config": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Matching configuration is filtered during import. Match statements do not control whether to import files that are part of the imported configuration package. Files are imported regardless of whether the filter disallows the resources with which the files are associated. <p>To create a match statement, type a correctly formatted resource match in the horizontal text box or select Build. Selecting Build displays the Configuration Match Builder in a popup window.</p><p>A match statement takes the following general form: <br/><i>Addr</i> / <i>Domain</i> / <i>Resource</i> [? <i>Name=resource-name</i> &amp; <i>Property=property-name</i> &amp; <i>Value=property-value</i> ]</p><table><tr><td valign=\"top\">Addr</td><td>Device Address. Specifies IP address or host alias. The value (*) matches all IP addresses.</td></tr><tr><td valign=\"top\">Domain</td><td>Application Domain. The name of the application domain. The value (*) matches all domains.</td></tr><tr><td valign=\"top\">Resource</td><td>Resource Type. The name of the resource type. The value (*) matches all resource types.</td></tr><tr><td valign=\"top\">Name=resource-name</td><td>An additional specification field, such as \"Name\". Limits the match statement to resources of the specified name. Use a PCRE to select groups of resource instances. For example, \"Name=foo.*\" would match all resources with names that start with \"foo\".</td></tr><tr><td valign=\"top\">Property=property-name</td><td>Property Name. The name of the configuration property. Limits the match statement to resources of the specified property. If change specified, set property-name to null string.</td></tr><tr><td valign=\"top\">Value=property-value</td><td>Property Value. Specifies the value for the configuration property. This property limits the match statement to resources with the specified property value.</td></tr></table>", "filter", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"modified_config": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Matching configuration is modified during import. The matching configuration may be changed, added, or deleted. To create a match statement, type a correctly formatted resource match in the horizontal text box or select Build. Selecting Build displays the Configuration Match Builder in a popup window. <p>A match statement takes the following general form: <br/><i>Addr</i> / <i>Domain</i> / <i>Resource</i> [? <i>Name=resource-name</i> &amp; <i>Property=property-name</i> &amp; <i>Value=property-value</i> ]</p><table><tr><td valign=\"top\">Addr</td><td>Device Address. Specifies IP address or host alias. The value (*) matches all IP addresses.</td></tr><tr><td valign=\"top\">Domain</td><td>Application Domain. The name of the application domain. The value (*) matches all domains.</td></tr><tr><td valign=\"top\">Resource</td><td>Resource Type. The name of the resource type. The value (*) matches all resource types.</td></tr><tr><td valign=\"top\">Name=resource-name</td><td>An additional specification field, such as \"Name\". Limits the match statement to resources of the specified name. Use a PCRE to select groups of resource instances. For example, \"Name=foo.*\" would match all resources with names that start with \"foo\".</td></tr><tr><td valign=\"top\">Property=property-name</td><td>Property Name. The name of the configuration property. Limits the match statement to resources of the specified property. If change specified, set property-name to null string.</td></tr><tr><td valign=\"top\">Value=property-value</td><td>Property Value. Specifies the value for the configuration property. This property limits the match statement to resources with the specified property value.</td></tr></table>", "modify", "").String,
				NestedObject:        models.DmConfigModifyTypeResourceSchema,
				Optional:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *ConfigDeploymentPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *ConfigDeploymentPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.ConfigDeploymentPolicy
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `ConfigDeploymentPolicy`)
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

func (r *ConfigDeploymentPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.ConfigDeploymentPolicy
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
		data.FromBody(ctx, `ConfigDeploymentPolicy`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `ConfigDeploymentPolicy`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ConfigDeploymentPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.ConfigDeploymentPolicy
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `ConfigDeploymentPolicy`))
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

func (r *ConfigDeploymentPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.ConfigDeploymentPolicy
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

func (r *ConfigDeploymentPolicyResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.ConfigDeploymentPolicy

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
