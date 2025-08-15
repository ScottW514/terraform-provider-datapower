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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var (
	_ datasource.DataSource              = &SSHServerProfileDataSource{}
	_ datasource.DataSourceWithConfigure = &SSHServerProfileDataSource{}
)

func NewSSHServerProfileDataSource() datasource.DataSource {
	return &SSHServerProfileDataSource{}
}

type SSHServerProfileDataSource struct {
	pData *tfutils.ProviderData
}

func (d *SSHServerProfileDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssh_server_profile"
}

func (d *SSHServerProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The SSH domain client profile defines the cipher suites.",
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
			"ciphers": schema.ListAttribute{
				MarkdownDescription: "Specify the SSH cipher suites to support.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"kex_alg": schema.ListAttribute{
				MarkdownDescription: "Specify the key exchange (KEX) algorithms to support.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"mac_alg": schema.ListAttribute{
				MarkdownDescription: "Specify the message authentication codes (MAC) to support.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"send_pre_auth_msg": schema.BoolAttribute{
				MarkdownDescription: "Specify whether to include a banner message during SSH preauthentication. The banner message contains the text to display to users before the login prompt.",
				Computed:            true,
			},
			"pre_auth_msg": schema.StringAttribute{
				MarkdownDescription: "Specify the banner message to display to users before the login prompt. White space characters ( <tt>\\n</tt> , <tt>\\t</tt> ) are recognized and escaped. The banner message must be at least one character. The banner message is truncated at 4096 characters.",
				Computed:            true,
			},
			"host_key_alg":       models.GetDmHostKeyAlgorithmsDataSourceSchema("Host key algorithms", "host-key-alg", ""),
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (d *SSHServerProfileDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *SSHServerProfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data models.SSHServerProfile
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := d.pData.Client.Get(data.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	data.FromBody(ctx, `SSHServerProfile`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
