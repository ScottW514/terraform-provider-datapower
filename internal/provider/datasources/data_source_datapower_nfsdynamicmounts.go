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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var (
	_ datasource.DataSource              = &NFSDynamicMountsDataSource{}
	_ datasource.DataSourceWithConfigure = &NFSDynamicMountsDataSource{}
)

func NewNFSDynamicMountsDataSource() datasource.DataSource {
	return &NFSDynamicMountsDataSource{}
}

type NFSDynamicMountsDataSource struct {
	pData *tfutils.ProviderData
}

func (d *NFSDynamicMountsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nfsdynamicmounts"
}

func (d *NFSDynamicMountsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Configure parameters of dynamic NFS mounts for dpnfs URL calls. These mounts support URL access in the form <tt>dpnfs://&lt;host>/&lt;path>/&lt;file></tt> . The system automatically mounts any dynamic mounts. Dynamic mounts remain mounted until the inactivity timer elapses.",
		Attributes: map[string]schema.Attribute{
			"app_domain": schema.StringAttribute{
				MarkdownDescription: "The name of the application domain the object belongs to",
				Required:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "<p>The administrative state of the configuration.</p><ul><li>To make active, set to enabled.</li><li>To make inactive, set to disabled.</li></ul>",
				Computed:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: "Comments",
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
			"idle_unmount_seconds": schema.Int64Attribute{
				MarkdownDescription: "Specify the inactivity duration in seconds to wait before the mount is unmounted. The default value is 900. The value of 0 disables the timer.",
				Computed:            true,
			},
			"mount_timeout_seconds": schema.Int64Attribute{
				MarkdownDescription: "Specify the duration in seconds to attempt to mount a dynamic mount. When the timer elapses, related file open operations fail.",
				Computed:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (d *NFSDynamicMountsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *NFSDynamicMountsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data models.NFSDynamicMounts
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := d.pData.Client.Get(data.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	data.FromBody(ctx, `NFSDynamicMounts`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
