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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type AS3SourceProtocolHandlerList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &AS3SourceProtocolHandlerDataSource{}
	_ datasource.DataSourceWithConfigure = &AS3SourceProtocolHandlerDataSource{}
)

func NewAS3SourceProtocolHandlerDataSource() datasource.DataSource {
	return &AS3SourceProtocolHandlerDataSource{}
}

type AS3SourceProtocolHandlerDataSource struct {
	pData *tfutils.ProviderData
}

func (d *AS3SourceProtocolHandlerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_as3sourceprotocolhandler"
}

func (d *AS3SourceProtocolHandlerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "AS3 handler",
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
							MarkdownDescription: "Local port",
							Computed:            true,
						},
						"filesystem_type": schema.StringAttribute{
							MarkdownDescription: "File system type",
							Computed:            true,
						},
						"persistent_filesystem_timeout": schema.Int64Attribute{
							MarkdownDescription: "Persistent timeout",
							Computed:            true,
						},
						"virtual_directories": schema.ListNestedAttribute{
							MarkdownDescription: "Virtual directories",
							NestedObject:        models.DmFTPServerVirtualDirectoryDataSourceSchema,
							Computed:            true,
						},
						"default_directory": schema.StringAttribute{
							MarkdownDescription: "Default directory",
							Computed:            true,
						},
						"max_filename_length": schema.Int64Attribute{
							MarkdownDescription: "Max file name length",
							Computed:            true,
						},
						"acl": schema.StringAttribute{
							MarkdownDescription: "Access control list",
							Computed:            true,
						},
						"require_tls": schema.StringAttribute{
							MarkdownDescription: "Require TLS",
							Computed:            true,
						},
						"password_aaa_policy": schema.StringAttribute{
							MarkdownDescription: "Username-password AAA policy",
							Computed:            true,
						},
						"certificate_aaa_policy": schema.StringAttribute{
							MarkdownDescription: "Password-required AAA policy",
							Computed:            true,
						},
						"allow_ccc": schema.BoolAttribute{
							MarkdownDescription: "Allow CCC command",
							Computed:            true,
						},
						"passive": schema.StringAttribute{
							MarkdownDescription: "Passive (PASV) command",
							Computed:            true,
						},
						"use_pasv_port_range": schema.BoolAttribute{
							MarkdownDescription: "Limit port range for passive connections",
							Computed:            true,
						},
						"pasv_min_port": schema.Int64Attribute{
							MarkdownDescription: "Min passive port",
							Computed:            true,
						},
						"pasv_max_port": schema.Int64Attribute{
							MarkdownDescription: "Max passive port",
							Computed:            true,
						},
						"pasv_idle_time_out": schema.Int64Attribute{
							MarkdownDescription: "Passive data connection idle timeout",
							Computed:            true,
						},
						"disable_pasvip_check": schema.BoolAttribute{
							MarkdownDescription: "Disable passive data connection (PASV) IP security check",
							Computed:            true,
						},
						"disable_portip_check": schema.BoolAttribute{
							MarkdownDescription: "Disable active data connection (PORT) IP security check",
							Computed:            true,
						},
						"use_alternate_pasv_addr": schema.BoolAttribute{
							MarkdownDescription: "Use alternate PASV IP address",
							Computed:            true,
						},
						"alternate_pasv_addr": schema.StringAttribute{
							MarkdownDescription: "Alternate PASV IP address",
							Computed:            true,
						},
						"allow_list_cmd": schema.BoolAttribute{
							MarkdownDescription: "Enable LIST command support",
							Computed:            true,
						},
						"allow_dele_cmd": schema.BoolAttribute{
							MarkdownDescription: "Enable DELE command support",
							Computed:            true,
						},
						"data_encryption": schema.StringAttribute{
							MarkdownDescription: "File transfer data encryption",
							Computed:            true,
						},
						"allow_compression": schema.BoolAttribute{
							MarkdownDescription: "Allow compression",
							Computed:            true,
						},
						"allow_stou": schema.BoolAttribute{
							MarkdownDescription: "Allow unique file name (STOU)",
							Computed:            true,
						},
						"unique_filename_prefix": schema.StringAttribute{
							MarkdownDescription: "Unique file name prefix",
							Computed:            true,
						},
						"allow_rest": schema.BoolAttribute{
							MarkdownDescription: "Allow restart (REST)",
							Computed:            true,
						},
						"restart_timeout": schema.Int64Attribute{
							MarkdownDescription: "Restart timeout",
							Computed:            true,
						},
						"idle_timeout": schema.Int64Attribute{
							MarkdownDescription: "Idle timeout",
							Computed:            true,
						},
						"response_type": schema.StringAttribute{
							MarkdownDescription: "Response type",
							Computed:            true,
						},
						"response_storage": schema.StringAttribute{
							MarkdownDescription: "Response storage",
							Computed:            true,
						},
						"temporary_storage_size": schema.Int64Attribute{
							MarkdownDescription: "Temporary storage size",
							Computed:            true,
						},
						"response_nfs_mount": schema.StringAttribute{
							MarkdownDescription: "Response NFS mount",
							Computed:            true,
						},
						"response_suffix": schema.StringAttribute{
							MarkdownDescription: "Response suffix",
							Computed:            true,
						},
						"response_url": schema.StringAttribute{
							MarkdownDescription: "Response URL",
							Computed:            true,
						},
						"ssl_server_config_type": schema.StringAttribute{
							MarkdownDescription: "TLS server type",
							Computed:            true,
						},
						"ssl_server": schema.StringAttribute{
							MarkdownDescription: "TLS server profile",
							Computed:            true,
						},
						"sslsni_server": schema.StringAttribute{
							MarkdownDescription: "TLS SNI server profile",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *AS3SourceProtocolHandlerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *AS3SourceProtocolHandlerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data AS3SourceProtocolHandlerList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.AS3SourceProtocolHandler{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.AS3SourceProtocolHandler{}
	if value := res.Get(`AS3SourceProtocolHandler`); value.Exists() {
		for _, v := range value.Array() {
			item := models.AS3SourceProtocolHandler{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.AS3SourceProtocolHandlerObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.AS3SourceProtocolHandlerObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
