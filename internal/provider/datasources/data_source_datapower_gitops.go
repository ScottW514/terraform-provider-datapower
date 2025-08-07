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
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
)

var (
	_ datasource.DataSource              = &GitOpsDataSource{}
	_ datasource.DataSourceWithConfigure = &GitOpsDataSource{}
)

func NewGitOpsDataSource() datasource.DataSource {
	return &GitOpsDataSource{}
}

type GitOpsDataSource struct {
	client *client.DatapowerClient
}

func (d *GitOpsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gitops"
}

func (d *GitOpsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GitOps",
		Attributes: map[string]schema.Attribute{
			"app_domain": schema.StringAttribute{
				MarkdownDescription: "The name of the application domain the object belongs to",
				Required:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Administrative state",
				Computed:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: "Comments",
				Computed:            true,
			},
			"connection_type": schema.StringAttribute{
				MarkdownDescription: "Connection type",
				Computed:            true,
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "Operation mode",
				Computed:            true,
			},
			"commit_identifier_type": schema.StringAttribute{
				MarkdownDescription: "Commit identifier type",
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
				MarkdownDescription: "Sync Interval",
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
				MarkdownDescription: "SSH authorized key file",
				Computed:            true,
			},
			"tls_valcred": schema.StringAttribute{
				MarkdownDescription: "HTTPS validation credentials",
				Computed:            true,
			},
			"git_user": schema.StringAttribute{
				MarkdownDescription: "Git user",
				Computed:            true,
			},
			"git_email": schema.StringAttribute{
				MarkdownDescription: "Git email",
				Computed:            true,
			},
			"json_parse_settings": schema.StringAttribute{
				MarkdownDescription: "JSON parse settings",
				Computed:            true,
			},
			"template_policies": schema.ListNestedAttribute{
				MarkdownDescription: "Template policies",
				NestedObject:        models.DmGitOpsTemplatePolicyDataSourceSchema,
				Computed:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (d *GitOpsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *GitOpsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.GitOps

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := d.client.Get(data.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	data.FromBody(ctx, `GitOps`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
