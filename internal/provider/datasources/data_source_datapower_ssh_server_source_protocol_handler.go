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

package datasources

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type SSHServerSourceProtocolHandlerList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &SSHServerSourceProtocolHandlerDataSource{}
	_ datasource.DataSourceWithConfigure = &SSHServerSourceProtocolHandlerDataSource{}
)

func NewSSHServerSourceProtocolHandlerDataSource() datasource.DataSource {
	return &SSHServerSourceProtocolHandlerDataSource{}
}

type SSHServerSourceProtocolHandlerDataSource struct {
	pData *tfutils.ProviderData
}

func (d *SSHServerSourceProtocolHandlerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssh_server_source_protocol_handler"
}

func (d *SSHServerSourceProtocolHandlerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "<p>The SFTP server handler provides an SSH SFTP server that can be used to submit files for processing by the system. Each file that is written results in one transaction.</p><p>There can be multiple SFTP servers, but only one server can be configured to listen on the default SSH port 22 on a given interface. There can be multiple simultaneous connections from SFTP clients to the same SFTP server.</p><p><b>Note:</b> Changes in the configuration affect only new connections to this SFTP server. Existing connections continue to use their current configuration until they disconnect.</p>",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The name of the object to retrieve.",
				Optional:            true,
			},
			"app_domain": schema.StringAttribute{
				MarkdownDescription: "The name of the application domain the object belongs to.",
				Required:            true,
			},
			"result": schema.ListNestedAttribute{
				MarkdownDescription: "List of objects. If `id` was provided and it exists, it will be the only item in the list.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Name of the object. Must be unique among object types in application domain.",
							Computed:            true,
						},
						"app_domain": schema.StringAttribute{
							MarkdownDescription: "The name of the application domain the object belongs to",
							Computed:            true,
						},
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"local_address": schema.StringAttribute{
							MarkdownDescription: "Specifies the address on which the SFTP server service listens. The default of 0.0.0.0 indicates that the service is active on all addresses. An alias name can be used to specify the address. Local host aliases can help ease migration tasks among machines.",
							Computed:            true,
						},
						"local_port": schema.Int64Attribute{
							MarkdownDescription: "Specifies the port that is monitored by the SFTP server service. This port is the port on which SFTP connections can be established. This port does not control the TCP port that is used for the data connections. Enter a value in the range 1 - 65535. The default value is 22.",
							Computed:            true,
						},
						"acl": schema.StringAttribute{
							MarkdownDescription: "<p>Specifies an Access Control List to apply. The ACL allows or denies access to the SFTP server based on the IP address of the SFTP client. When attached to a server, the default for an ACL is to deny all access. To deny access to only select IP addresses, first grant access to all addresses (allow 0.0.0.0). Then, create deny entries for the desired hosts.</p><p>If an ACL with the same name as this handler exists, the system DataPower Gateway might inadvertently use that ACL rather than the one specified here.</p>",
							Computed:            true,
						},
						"host_private_keys": schema.ListAttribute{
							MarkdownDescription: "Specifies the private keys to use for Host authentication. Keys used as host private keys cannot be password protected.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"ssh_user_authentication": models.GetDmSSHUserAuthenticationMethodsDataSourceSchema("Specifies the type(s) of SSH user authentication available for use by the client.", "user-auth", ""),
						"allow_backend_listings": schema.BoolAttribute{
							MarkdownDescription: "In transparent mode, determines whether or not SFTP directory listing (SSH_FXP_READDIR) requests to remote servers are allowed. Requires a remote FTP or SFTP server.",
							Computed:            true,
						},
						"allow_backend_delete": schema.BoolAttribute{
							MarkdownDescription: "In transparent mode, determines whether or not requests to delete files (SSH_FXP_REMOVE) to remote servers are allowed. Requires a remote FTP or SFTP server.",
							Computed:            true,
						},
						"allow_backend_stat": schema.BoolAttribute{
							MarkdownDescription: "In transparent mode, determines whether or not SFTP directory listings requests to remote servers would query the remote server to obtain file attributes (SSH_FXP_STAT/SSH_FXP_LSTAT/SSH_FXP_FSTAT), or use default values. Querying the remote server may reduce performance, but is necessary for SFTP clients that do not follow the DataPower SFTP URL naming conventions. Requires a remote FTP or SFTP server.",
							Computed:            true,
						},
						"allow_backend_mkdir": schema.BoolAttribute{
							MarkdownDescription: "In transparent mode, determines whether or not requests to create directories (SSH_FXP_MKDIR) on remote servers are allowed. Requires a remote FTP or SFTP server.",
							Computed:            true,
						},
						"allow_backend_rmdir": schema.BoolAttribute{
							MarkdownDescription: "In transparent mode, determines whether or not requests to delete directories (SSH_FXP_RMDIR) from remote servers are allowed. Requires a remote FTP or SFTP server.",
							Computed:            true,
						},
						"allow_backend_rename": schema.BoolAttribute{
							MarkdownDescription: "In transparent mode, determines whether or not requests to rename files or directories (SSH_FXP_RENAME) on remote servers are allowed. Requires a remote FTP or SFTP server.",
							Computed:            true,
						},
						"aaa_policy": schema.StringAttribute{
							MarkdownDescription: "AAA policy",
							Computed:            true,
						},
						"filesystem_type": schema.StringAttribute{
							MarkdownDescription: "File system type",
							Computed:            true,
						},
						"default_directory": schema.StringAttribute{
							MarkdownDescription: "Default directory",
							Computed:            true,
						},
						"idle_timeout": schema.Int64Attribute{
							MarkdownDescription: "Specify the duration in seconds that the SFTP connection can be idle. After the specified duration, the SFTP server closes the control connection. Enter a value in the range 0 - 65535. The default value is 0, which disables the timeout.",
							Computed:            true,
						},
						"persistent_filesystem_timeout": schema.Int64Attribute{
							MarkdownDescription: "Specify the duration in seconds that a connection to a virtual persistent file system is retained after all SFTP control connections from user identities are disconnected. When the timeout expires, the virtual file system object is destroyed. Enter a value in the range 1- 43200. The default value is 600.",
							Computed:            true,
						},
						"virtual_directories": schema.ListNestedAttribute{
							MarkdownDescription: "In virtual mode, create a directory in the virtual file system that is presented by this SFTP server. The SFTP client can use all of these directories to write file to be processed. The root directory (/) is always present and cannot be created.",
							NestedObject:        models.GetDmSFTPServerVirtualDirectoryDataSourceSchema(),
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *SSHServerSourceProtocolHandlerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *SSHServerSourceProtocolHandlerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data SSHServerSourceProtocolHandlerList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.SSHServerSourceProtocolHandler{
		AppDomain: data.AppDomain,
	}

	path := o.GetPath()
	if !data.Id.IsNull() {
		path = path + "/" + data.Id.ValueString()
	}

	res, err := d.pData.Client.Get(path)
	resFound := true
	if err != nil {
		if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(d.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString())
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		} else {
			resFound = false
		}
	}
	l := []models.SSHServerSourceProtocolHandler{}
	if resFound {
		if value := res.Get(`SSHServerSourceProtocolHandler`); value.Exists() {
			for _, v := range value.Array() {
				item := models.SSHServerSourceProtocolHandler{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.SSHServerSourceProtocolHandlerObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.SSHServerSourceProtocolHandlerObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
