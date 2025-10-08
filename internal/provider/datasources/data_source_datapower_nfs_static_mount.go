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

type NFSStaticMountList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &NFSStaticMountDataSource{}
	_ datasource.DataSourceWithConfigure = &NFSStaticMountDataSource{}
)

func NewNFSStaticMountDataSource() datasource.DataSource {
	return &NFSStaticMountDataSource{}
}

type NFSStaticMountDataSource struct {
	pData *tfutils.ProviderData
}

func (d *NFSStaticMountDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nfs_static_mount"
}

func (d *NFSStaticMountDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Create static NFS mounts for URL or file system access. These mounts remain mounted as long as their application domain is up.",
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
						"remote": schema.StringAttribute{
							MarkdownDescription: "Specify the remote NFS file system to mount. Use the form <tt>host:/path</tt> , where <tt>host</tt> is the DNS name or IP address of the NFS server, and <tt>path</tt> is the path exported by the host to mount.",
							Computed:            true,
						},
						"local_filesystem_access": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to allow local file system access through the <tt>nfs-&lt;name></tt> directory. By default, local access is not enabled. When enabled, the NFS mount is available for file system access through the CLI in the <tt>nfs-&lt;name></tt> directory, where <tt>&lt;name></tt> is the name of the mount.",
							Computed:            true,
						},
						"version": schema.Int64Attribute{
							MarkdownDescription: "Specify the preferred NFS protocol version. Enter a value in the range 2 - 4. The default value is 3. <ul><li>If version 3 and the server only implements version 2, the client falls back to version 2.</li><li>If version 4, the remote export paths are different and prevents fallback.</li></ul>",
							Computed:            true,
						},
						"transport": schema.StringAttribute{
							MarkdownDescription: "Specify the transport protocol. The default transport protocol is TCP.",
							Computed:            true,
						},
						"mount_type": schema.StringAttribute{
							MarkdownDescription: "Specify the type of NFS mount. The default mount type is a hard mount.",
							Computed:            true,
						},
						"read_only": schema.BoolAttribute{
							MarkdownDescription: "Specify whether the mount is read-only. By default, the mount is not read-only.",
							Computed:            true,
						},
						"read_size": schema.Int64Attribute{
							MarkdownDescription: "Specify the size in bytes for NFS read operations. Enter a value in the range 1024 - 32768. The default value is 4096.",
							Computed:            true,
						},
						"write_size": schema.Int64Attribute{
							MarkdownDescription: "Specify the size in bytes for NFS write operations. Enter a value in the range 1024 - 32768. The default value is 4096.",
							Computed:            true,
						},
						"timeout": schema.Int64Attribute{
							MarkdownDescription: "Specify the time in tenths of seconds until the first retransmission on RPC times out. Enter a value in the range 1 - 600. The default value is 7.",
							Computed:            true,
						},
						"retransmissions": schema.Int64Attribute{
							MarkdownDescription: "Specify the number of minor RPC timeouts and retransmissions until a major timeout. Enter a value in the range 1 - 60. The default value is 3.",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *NFSStaticMountDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *NFSStaticMountDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data NFSStaticMountList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.NFSStaticMount{
		AppDomain: data.AppDomain,
	}

	path := o.GetPath()
	if !data.Id.IsNull() {
		path = path + "/" + data.Id.ValueString()
	}

	res, err := d.pData.Client.Get(path)
	resFound := true
	if err != nil {
		if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		} else {
			resFound = false
		}
	}
	l := []models.NFSStaticMount{}
	if resFound {
		if value := res.Get(`NFSStaticMount`); value.Exists() {
			for _, v := range value.Array() {
				item := models.NFSStaticMount{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.NFSStaticMountObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.NFSStaticMountObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
