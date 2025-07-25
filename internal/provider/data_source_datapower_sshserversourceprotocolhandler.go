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

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
)

type SSHServerSourceProtocolHandlerList struct {
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
	client *client.DatapowerClient
}

func (d *SSHServerSourceProtocolHandlerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sshserversourceprotocolhandler"
}

func (d *SSHServerSourceProtocolHandlerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SFTP server handler",
		Attributes: map[string]schema.Attribute{
			"app_domain": schema.StringAttribute{
				MarkdownDescription: "The name of the application domain the object belongs to",
				Required:            true,
			},
			"result": schema.ListNestedAttribute{
				MarkdownDescription: "List of objects",
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
							MarkdownDescription: "Local IP address",
							Computed:            true,
						},
						"local_port": schema.Int64Attribute{
							MarkdownDescription: "Port",
							Computed:            true,
						},
						"acl": schema.StringAttribute{
							MarkdownDescription: "Access control list",
							Computed:            true,
						},
						"host_private_keys": schema.ListAttribute{
							MarkdownDescription: "Host private keys",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"ssh_user_authentication": models.GetDmSSHUserAuthenticationMethodsDataSourceSchema("User authentication", "user-auth", ""),
						"allow_backend_listings": schema.BoolAttribute{
							MarkdownDescription: "Allow back-end listings",
							Computed:            true,
						},
						"allow_backend_delete": schema.BoolAttribute{
							MarkdownDescription: "Allow file deletes",
							Computed:            true,
						},
						"allow_backend_stat": schema.BoolAttribute{
							MarkdownDescription: "Allow stat operations",
							Computed:            true,
						},
						"allow_backend_mkdir": schema.BoolAttribute{
							MarkdownDescription: "Allow directory creates",
							Computed:            true,
						},
						"allow_backend_rmdir": schema.BoolAttribute{
							MarkdownDescription: "Allow directory deletes",
							Computed:            true,
						},
						"allow_backend_rename": schema.BoolAttribute{
							MarkdownDescription: "Allow file/directory renames",
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
							MarkdownDescription: "Idle timeout",
							Computed:            true,
						},
						"persistent_filesystem_timeout": schema.Int64Attribute{
							MarkdownDescription: "Persistent file system timeout",
							Computed:            true,
						},
						"virtual_directories": schema.ListNestedAttribute{
							MarkdownDescription: "Virtual directories",
							NestedObject:        models.DmSFTPServerVirtualDirectoryDataSourceSchema,
							Computed:            true,
						},
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

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *SSHServerSourceProtocolHandlerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SSHServerSourceProtocolHandlerList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.SSHServerSourceProtocolHandler{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.SSHServerSourceProtocolHandler{}
	if value := res.Get(`SSHServerSourceProtocolHandler`); value.Exists() {
		for _, v := range value.Array() {
			item := models.SSHServerSourceProtocolHandler{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
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
