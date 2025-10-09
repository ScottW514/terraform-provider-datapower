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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var (
	_ datasource.DataSource              = &GitOpsDataSource{}
	_ datasource.DataSourceWithConfigure = &GitOpsDataSource{}
)

func NewGitOpsDataSource() datasource.DataSource {
	return &GitOpsDataSource{}
}

type GitOpsDataSource struct {
	pData *tfutils.ProviderData
}

func (d *GitOpsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_git_ops"
}

func (d *GitOpsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Each DataPower domain supports a single GitOps instance that operates in either read-only mode or read/write mode. The DataPower GitOps integration helps to automate configuration management through version control. This integration supports industry-standard GitOps practices and authoring experiences.",
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
			"connection_type": schema.StringAttribute{
				MarkdownDescription: "Specify the protocol to secure the connection. HTTPS is the default protocol.",
				Computed:            true,
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "Specify the operational mode of the Git repository. The default mode is read-only.",
				Computed:            true,
			},
			"commit_identifier_type": schema.StringAttribute{
				MarkdownDescription: "Specify the branch, commit hash, or tag for read and write GitOps operations against the repository. Use of branch is the default setting.",
				Computed:            true,
			},
			"commit_identifier": schema.StringAttribute{
				MarkdownDescription: "Commit identifier",
				Computed:            true,
			},
			"remote_location": schema.StringAttribute{
				MarkdownDescription: "Remote location",
				Computed:            true,
			},
			"interval": schema.Int64Attribute{
				MarkdownDescription: "Specify the interval in minutes to poll the repository for changes. Enter a value in the range 0 - 1440. The default value is 5. To disable polling, specify 0.",
				Computed:            true,
			},
			"ssh_client_profile": schema.StringAttribute{
				MarkdownDescription: "SSH client profile",
				Computed:            true,
			},
			"username": schema.StringAttribute{
				MarkdownDescription: "Username",
				Computed:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Password",
				Computed:            true,
			},
			"ssh_authorized_keys_file": schema.StringAttribute{
				MarkdownDescription: "Specify the file that contains the authorized SSH keys. This file must be in the <tt>cert:</tt> or <tt>sharedcert:</tt> directory.",
				Computed:            true,
			},
			"tls_valcred": schema.StringAttribute{
				MarkdownDescription: "HTTPS validation credentials",
				Computed:            true,
			},
			"git_user": schema.StringAttribute{
				MarkdownDescription: "Specify the full username. Controls <tt>user.name</tt> in <tt>git config</tt> .",
				Computed:            true,
			},
			"git_email": schema.StringAttribute{
				MarkdownDescription: "Specify the user emai. Controls <tt>user.email</tt> in <tt>git config</tt> .",
				Computed:            true,
			},
			"json_parse_settings": schema.StringAttribute{
				MarkdownDescription: "JSON parse settings",
				Computed:            true,
			},
			"template_policies": schema.ListNestedAttribute{
				MarkdownDescription: "Specify the list of template policy for GitOps processing. The policy processing is in the order of the read or write GitOps operation.",
				NestedObject:        models.GetDmGitOpsTemplatePolicyDataSourceSchema(),
				Computed:            true,
			},
		},
	}
}

func (d *GitOpsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *GitOpsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data models.GitOps
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	path := data.GetPath()

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
	if resFound {
		data.FromBody(ctx, `GitOps`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
