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
)

var _ resource.Resource = &GitOpsResource{}

func NewGitOpsResource() resource.Resource {
	return &GitOpsResource{}
}

type GitOpsResource struct {
	pData *tfutils.ProviderData
}

func (r *GitOpsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gitops"
}

func (r *GitOpsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Each DataPower domain supports a single GitOps instance that operates in either read-only mode or read/write mode. The DataPower GitOps integration helps to automate configuration management through version control. This integration supports industry-standard GitOps practices and authoring experiences.", "gitops", "").String,
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
			"connection_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the protocol to secure the connection. HTTPS is the default protocol.", "type", "").AddStringEnum("https", "ssh").AddDefaultValue("https").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("https", "ssh"),
				},
				Default: stringdefault.StaticString("https"),
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the operational mode of the Git repository. The default mode is read-only.", "mode", "").AddStringEnum("read-only", "read-write").AddDefaultValue("read-write").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("read-only", "read-write"),
				},
				Default: stringdefault.StaticString("read-write"),
			},
			"commit_identifier_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the branch, commit hash, or tag for read and write GitOps operations against the repository. Use of branch is the default setting.", "commit-id-type", "").AddStringEnum("branch", "tag", "commit").AddDefaultValue("branch").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("branch", "tag", "commit"),
				},
				Default: stringdefault.StaticString("branch"),
			},
			"commit_identifier": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Commit identifier", "commit-id", "").String,
				Required:            true,
			},
			"remote_location": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Remote location", "location", "").String,
				Required:            true,
			},
			"interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the interval in minutes to poll the repository for changes. Enter a value in the range 0 - 1440. The default value is 5. To disable polling, specify 0.", "interval", "").AddIntegerRange(0, 1440).AddDefaultValue("5").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 1440),
				},
				Default: int64default.StaticInt64(5),
			},
			"ssh_client_profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SSH client profile", "ssh-client-profile", "sshclientprofile").String,
				Optional:            true,
			},
			"username": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Username", "username", "").String,
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Password", "password", "passwordalias").String,
				Optional:            true,
			},
			"ssh_authorized_keys_file": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the file that contains the authorized SSH keys. This file must be in the <tt>cert:</tt> or <tt>sharedcert:</tt> directory.", "ssh-authorized-keyfile", "").String,
				Optional:            true,
			},
			"tls_valcred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTPS validation credentials", "https-valcred", "cryptovalcred").String,
				Optional:            true,
			},
			"git_user": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the full username. Controls <tt>user.name</tt> in <tt>git config</tt> .", "name", "").String,
				Optional:            true,
			},
			"git_email": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the user emai. Controls <tt>user.email</tt> in <tt>git config</tt> .", "email", "").String,
				Optional:            true,
			},
			"json_parse_settings": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("JSON parse settings", "json-settings", "jsonsettings").String,
				Optional:            true,
			},
			"template_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the list of template policy for GitOps processing. The policy processing is in the order of the read or write GitOps operation.", "template-policy", "").String,
				NestedObject:        models.DmGitOpsTemplatePolicyResourceSchema,
				Optional:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *GitOpsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *GitOpsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.GitOps
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `GitOps`)
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

func (r *GitOpsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.GitOps
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
		data.FromBody(ctx, `GitOps`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `GitOps`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GitOpsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.GitOps
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath(), data.ToBody(ctx, `GitOps`))
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

func (r *GitOpsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.GitOps
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

func (r *GitOpsResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.GitOps

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
