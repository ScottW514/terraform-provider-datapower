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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

var _ resource.Resource = &SSHServerSourceProtocolHandlerResource{}
var _ resource.ResourceWithValidateConfig = &SSHServerSourceProtocolHandlerResource{}
var _ resource.ResourceWithImportState = &SSHServerSourceProtocolHandlerResource{}

func NewSSHServerSourceProtocolHandlerResource() resource.Resource {
	return &SSHServerSourceProtocolHandlerResource{}
}

type SSHServerSourceProtocolHandlerResource struct {
	pData *tfutils.ProviderData
}

func (r *SSHServerSourceProtocolHandlerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssh_server_source_protocol_handler"
}

func (r *SSHServerSourceProtocolHandlerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("<p>The SFTP server handler provides an SSH SFTP server that can be used to submit files for processing by the system. Each file that is written results in one transaction.</p><p>There can be multiple SFTP servers, but only one server can be configured to listen on the default SSH port 22 on a given interface. There can be multiple simultaneous connections from SFTP clients to the same SFTP server.</p><p><b>Note:</b> Changes in the configuration affect only new connections to this SFTP server. Existing connections continue to use their current configuration until they disconnect.</p>", "source-ssh-server", "").AddActions("quiesce").String,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name of the object. Must be unique among object types in application domain.", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
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
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
				PlanModifiers: []planmodifier.String{
					modifiers.ImmutableAfterSet(),
				},
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"local_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the address on which the SFTP server service listens. The default of 0.0.0.0 indicates that the service is active on all addresses. An alias name can be used to specify the address. Local host aliases can help ease migration tasks among machines.", "address", "").AddDefaultValue("0.0.0.0").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("0.0.0.0"),
			},
			"local_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the port that is monitored by the SFTP server service. This port is the port on which SFTP connections can be established. This port does not control the TCP port that is used for the data connections. Enter a value in the range 1 - 65535. The default value is 22.", "port", "").AddIntegerRange(1, 65534).AddDefaultValue("22").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65534),
				},
				Default: int64default.StaticInt64(22),
			},
			"acl": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specifies an Access Control List to apply. The ACL allows or denies access to the SFTP server based on the IP address of the SFTP client. When attached to a server, the default for an ACL is to deny all access. To deny access to only select IP addresses, first grant access to all addresses (allow 0.0.0.0). Then, create deny entries for the desired hosts.</p><p>If an ACL with the same name as this handler exists, the system DataPower Gateway might inadvertently use that ACL rather than the one specified here.</p>", "acl", "access_control_list").String,
				Optional:            true,
			},
			"host_private_keys": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the private keys to use for Host authentication. Keys used as host private keys cannot be password protected.", "host-private-key", "crypto_key").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"ssh_user_authentication": models.GetDmSSHUserAuthenticationMethodsResourceSchema("Specifies the type(s) of SSH user authentication available for use by the client.", "user-auth", "", false),
			"allow_backend_listings": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("In transparent mode, determines whether or not SFTP directory listing (SSH_FXP_READDIR) requests to remote servers are allowed. Requires a remote FTP or SFTP server.", "allow-backend-listings", "").AddDefaultValue("true").AddNotValidWhen(models.SSHServerSourceProtocolHandlerAllowBackendListingsIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"allow_backend_delete": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("In transparent mode, determines whether or not requests to delete files (SSH_FXP_REMOVE) to remote servers are allowed. Requires a remote FTP or SFTP server.", "allow-backend-delete", "").AddDefaultValue("false").AddNotValidWhen(models.SSHServerSourceProtocolHandlerAllowBackendDeleteIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"allow_backend_stat": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("In transparent mode, determines whether or not SFTP directory listings requests to remote servers would query the remote server to obtain file attributes (SSH_FXP_STAT/SSH_FXP_LSTAT/SSH_FXP_FSTAT), or use default values. Querying the remote server may reduce performance, but is necessary for SFTP clients that do not follow the DataPower SFTP URL naming conventions. Requires a remote FTP or SFTP server.", "allow-backend-stat", "").AddDefaultValue("false").AddNotValidWhen(models.SSHServerSourceProtocolHandlerAllowBackendStatIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"allow_backend_mkdir": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("In transparent mode, determines whether or not requests to create directories (SSH_FXP_MKDIR) on remote servers are allowed. Requires a remote FTP or SFTP server.", "allow-backend-mkdir", "").AddDefaultValue("false").AddNotValidWhen(models.SSHServerSourceProtocolHandlerAllowBackendMkdirIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"allow_backend_rmdir": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("In transparent mode, determines whether or not requests to delete directories (SSH_FXP_RMDIR) from remote servers are allowed. Requires a remote FTP or SFTP server.", "allow-backend-rmdir", "").AddDefaultValue("false").AddNotValidWhen(models.SSHServerSourceProtocolHandlerAllowBackendRmdirIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"allow_backend_rename": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("In transparent mode, determines whether or not requests to rename files or directories (SSH_FXP_RENAME) on remote servers are allowed. Requires a remote FTP or SFTP server.", "allow-backend-rename", "").AddDefaultValue("false").AddNotValidWhen(models.SSHServerSourceProtocolHandlerAllowBackendRenameIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"aaa_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("AAA policy", "aaa-policy", "aaa_policy").String,
				Optional:            true,
			},
			"filesystem_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("File system type", "filesystem", "").AddStringEnum("transparent", "virtual-ephemeral", "virtual-persistent").AddDefaultValue("transparent").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("transparent", "virtual-ephemeral", "virtual-persistent"),
				},
				Default: stringdefault.StaticString("transparent"),
			},
			"default_directory": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Default directory", "default-directory", "").AddDefaultValue("/").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile("^/([^/]+(/[^/]+)*)?$"), "Must match :"+"^/([^/]+(/[^/]+)*)?$"),
				},
				Default: stringdefault.StaticString("/"),
			},
			"idle_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration in seconds that the SFTP connection can be idle. After the specified duration, the SFTP server closes the control connection. Enter a value in the range 0 - 65535. The default value is 0, which disables the timeout.", "idle-timeout", "").AddIntegerRange(0, 65535).AddDefaultValue("0").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
				Default: int64default.StaticInt64(0),
			},
			"persistent_filesystem_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration in seconds that a connection to a virtual persistent file system is retained after all SFTP control connections from user identities are disconnected. When the timeout expires, the virtual file system object is destroyed. Enter a value in the range 1- 43200. The default value is 600.", "persistent-filesystem-timeout", "").AddIntegerRange(1, 43200).AddDefaultValue("600").AddNotValidWhen(models.SSHServerSourceProtocolHandlerPersistentFilesystemTimeoutIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 43200),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, models.SSHServerSourceProtocolHandlerPersistentFilesystemTimeoutIgnoreVal, true),
				},
				Default: int64default.StaticInt64(600),
			},
			"virtual_directories": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("In virtual mode, create a directory in the virtual file system that is presented by this SFTP server. The SFTP client can use all of these directories to write file to be processed. The root directory (/) is always present and cannot be created.", "virtual-directory", "").AddNotValidWhen(models.SSHServerSourceProtocolHandlerVirtualDirectoriesIgnoreVal.String()).String,
				NestedObject:        models.GetDmSFTPServerVirtualDirectoryResourceSchema(),
				Optional:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *SSHServerSourceProtocolHandlerResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *SSHServerSourceProtocolHandlerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SSHServerSourceProtocolHandler
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `SSHServerSourceProtocolHandler`)
	_, err := r.pData.Client.Post(data.GetPath(), body)
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), body)
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

func (r *SSHServerSourceProtocolHandlerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SSHServerSourceProtocolHandler
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

	data.UpdateFromBody(ctx, `SSHServerSourceProtocolHandler`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SSHServerSourceProtocolHandlerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SSHServerSourceProtocolHandler
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `SSHServerSourceProtocolHandler`))
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

func (r *SSHServerSourceProtocolHandlerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SSHServerSourceProtocolHandler
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			resp.Diagnostics.AddWarning("Resource Conflict", fmt.Sprintf("Resource is no longer tracked by Terraform, but may need to be manually deleted on DataPower host. Got error: %s", err))
		} else if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete resource, got error: %s", err))
			return
		}
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *SSHServerSourceProtocolHandlerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError("Invalid Import ID", "Expected format: <app_domain>/<id>. Got: "+req.ID)
		return
	}

	appDomain := parts[0]
	id := parts[1]

	if !regexp.MustCompile("^[a-zA-Z0-9_-]+$").MatchString(id) || len(id) < 1 || len(id) > 128 {
		resp.Diagnostics.AddError("Invalid ID", "ID must be 1-128 characters and match pattern ^[a-zA-Z0-9_-]+$")
		return
	}
	if !regexp.MustCompile("^[a-zA-Z0-9_-]+$").MatchString(appDomain) || len(appDomain) < 1 || len(appDomain) > 128 {
		resp.Diagnostics.AddError("Invalid Application Domain", "Application domain must be 1-128 characters and match pattern ^[a-zA-Z0-9_-]+$")
		return
	}

	var data models.SSHServerSourceProtocolHandler
	data.AppDomain = types.StringValue(appDomain)
	data.Id = types.StringValue(id)
	res, err := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("Resource was not found, got error: %s", err))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		}
		return
	}

	data.FromBody(ctx, `SSHServerSourceProtocolHandler`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SSHServerSourceProtocolHandlerResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.SSHServerSourceProtocolHandler

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
