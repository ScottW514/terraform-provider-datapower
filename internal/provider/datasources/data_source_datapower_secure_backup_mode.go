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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var (
	_ datasource.DataSource              = &SecureBackupModeDataSource{}
	_ datasource.DataSourceWithConfigure = &SecureBackupModeDataSource{}
)

func NewSecureBackupModeDataSource() datasource.DataSource {
	return &SecureBackupModeDataSource{}
}

type SecureBackupModeDataSource struct {
	pData *tfutils.ProviderData
}

func (d *SecureBackupModeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_secure_backup_mode"
}

func (d *SecureBackupModeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Sets the backup operational mode. This mode controls whether a secure-backup is allowed. After this mode is set, it cannot be changed. Operational modes are set the first time the DataPower Gateway is started.",
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "<p>The administrative state of the configuration.</p><ul><li>To make active, set to enabled.</li><li>To make inactive, set to disabled.</li></ul>",
				Computed:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: "A descriptive summary for the configuration.",
				Computed:            true,
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "<p>Sets the backup mode to support, when this operational mode is not previously set.</p><p>The creation of a secure backup is available only when secure backup mode is enabled. Unlike a standard backup, a secure backup contains private data (certificates, keys, and user data), which the DataPower Gateway encrypts with a customer-provided and a DataPower certificate.</p>",
				Computed:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (d *SecureBackupModeDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *SecureBackupModeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data models.SecureBackupMode
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
		data.FromBody(ctx, `SecureBackupMode`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
