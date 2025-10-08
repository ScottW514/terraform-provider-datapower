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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

var _ resource.Resource = &GitOpsResource{}
var _ resource.ResourceWithImportState = &GitOpsResource{}

func NewGitOpsResource() resource.Resource {
	return &GitOpsResource{}
}

type GitOpsResource struct {
	pData *tfutils.ProviderData
}

func (r *GitOpsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_git_ops"
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
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
				PlanModifiers: []planmodifier.String{
					modifiers.ImmutableAfterSet(),
				},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>The administrative state of the configuration.</p><ul><li>To make active, set to enabled.</li><li>To make inactive, set to disabled.</li></ul>", "admin-state", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"connection_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the protocol to secure the connection. HTTPS is the default protocol.", "type", "").AddStringEnum("https", "ssh").AddDefaultValue("https").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("https", "ssh"),
				},
				Default: stringdefault.StaticString("https"),
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the operational mode of the Git repository. The default mode is read-only.", "mode", "").AddStringEnum("read-only", "read-write").AddDefaultValue("read-write").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("read-only", "read-write"),
				},
				Default: stringdefault.StaticString("read-write"),
			},
			"commit_identifier_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the branch, commit hash, or tag for read and write GitOps operations against the repository. Use of branch is the default setting.", "commit-id-type", "").AddStringEnum("branch", "tag", "commit").AddDefaultValue("branch").AddNotValidWhen(models.GitOpsCommitIdentifierTypeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("branch", "tag", "commit"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.GitOpsCommitIdentifierTypeIgnoreVal, true),
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the interval in minutes to poll the repository for changes. Enter a value in the range 0 - 1440. The default value is 5. To disable polling, specify 0.", "interval", "").AddIntegerRange(0, 1440).AddDefaultValue("5").AddRequiredWhen(models.GitOpsIntervalCondVal.String()).AddNotValidWhen(models.GitOpsIntervalIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 1440),
					validators.ConditionalRequiredInt64(models.GitOpsIntervalCondVal, models.GitOpsIntervalIgnoreVal, true),
				},
				Default: int64default.StaticInt64(5),
			},
			"ssh_client_profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SSH client profile", "ssh-client-profile", "ssh_client_profile").AddRequiredWhen(models.GitOpsSSHClientProfileCondVal.String()).AddNotValidWhen(models.GitOpsSSHClientProfileIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.GitOpsSSHClientProfileCondVal, models.GitOpsSSHClientProfileIgnoreVal, false),
				},
			},
			"username": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Username", "username", "").AddRequiredWhen(models.GitOpsUsernameCondVal.String()).AddNotValidWhen(models.GitOpsUsernameIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.GitOpsUsernameCondVal, models.GitOpsUsernameIgnoreVal, false),
				},
			},
			"password": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Password", "password", "password_alias").AddRequiredWhen(models.GitOpsPasswordCondVal.String()).AddNotValidWhen(models.GitOpsPasswordIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.GitOpsPasswordCondVal, models.GitOpsPasswordIgnoreVal, false),
				},
			},
			"ssh_authorized_keys_file": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the file that contains the authorized SSH keys. This file must be in the <tt>cert:</tt> or <tt>sharedcert:</tt> directory.", "ssh-authorized-keyfile", "").AddRequiredWhen(models.GitOpsSSHAuthorizedKeysFileCondVal.String()).AddNotValidWhen(models.GitOpsSSHAuthorizedKeysFileIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.GitOpsSSHAuthorizedKeysFileCondVal, models.GitOpsSSHAuthorizedKeysFileIgnoreVal, false),
				},
			},
			"tls_valcred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTPS validation credentials", "https-valcred", "crypto_val_cred").AddRequiredWhen(models.GitOpsTLSValcredCondVal.String()).AddNotValidWhen(models.GitOpsTLSValcredIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.GitOpsTLSValcredCondVal, models.GitOpsTLSValcredIgnoreVal, false),
				},
			},
			"git_user": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the full username. Controls <tt>user.name</tt> in <tt>git config</tt> .", "name", "").AddRequiredWhen(models.GitOpsGitUserCondVal.String()).AddNotValidWhen(models.GitOpsGitUserIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.GitOpsGitUserCondVal, models.GitOpsGitUserIgnoreVal, false),
				},
			},
			"git_email": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the user emai. Controls <tt>user.email</tt> in <tt>git config</tt> .", "email", "").AddRequiredWhen(models.GitOpsGitEmailCondVal.String()).AddNotValidWhen(models.GitOpsGitEmailIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.GitOpsGitEmailCondVal, models.GitOpsGitEmailIgnoreVal, false),
				},
			},
			"json_parse_settings": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("JSON parse settings", "json-settings", "json_settings").String,
				Optional:            true,
			},
			"template_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the list of template policy for GitOps processing. The policy processing is in the order of the read or write GitOps operation.", "template-policy", "").String,
				NestedObject:        models.GetDmGitOpsTemplatePolicyResourceSchema(),
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
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			_, err := r.pData.Client.Put(data.GetPath(), body)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Resource already exists. Failed to update resource, got error: %s", err))
				return
			}
			resp.Diagnostics.AddWarning("Warning", "Resource already exists. Existing resource was updated.")
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create resource, got error: %s", err))
			return
		}
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

	data.UpdateFromBody(ctx, `GitOps`, res)

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
	data.ToDefault()
	_, err := r.pData.Client.Put(data.GetPath(), data.ToBody(ctx, `GitOps`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to restore singleton to default, got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *GitOpsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()
	appDomain := req.ID
	if !regexp.MustCompile("^[a-zA-Z0-9_-]+$").MatchString(appDomain) || len(appDomain) < 1 || len(appDomain) > 128 {
		resp.Diagnostics.AddError("Invalid Application Domain", "Application domain must be 1-128 characters and match pattern ^[a-zA-Z0-9_-]+$")
		return
	}

	var data models.GitOps
	data.AppDomain = types.StringValue(appDomain)
	res, err := r.pData.Client.Get(data.GetPath())
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("Resource was not found, got error: %s", err))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		}
		return
	}

	data.FromBody(ctx, `GitOps`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GitOpsResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.GitOps

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
