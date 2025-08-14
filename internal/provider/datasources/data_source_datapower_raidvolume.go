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
	_ datasource.DataSource              = &RaidVolumeDataSource{}
	_ datasource.DataSourceWithConfigure = &RaidVolumeDataSource{}
)

func NewRaidVolumeDataSource() datasource.DataSource {
	return &RaidVolumeDataSource{}
}

type RaidVolumeDataSource struct {
	pData *tfutils.ProviderData
}

func (d *RaidVolumeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_raidvolume"
}

func (d *RaidVolumeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "<p>Configure a RAID volume for data storage.</p><ol><li>The access permission to files on the storage volume. With the B2B feature, access permission must be read/write. B2B storage requires write access. Setting to read-only is ignored but generates a warning.</li><li>The subdirectory where files on the storage volume are available in the <tt>local:</tt> and <tt>logstore:</tt> directories. Each domain contains these subdirectories, and these subdirectories are not shared across domains.</li></ol>",
		Attributes: map[string]schema.Attribute{
			"user_summary": schema.StringAttribute{
				MarkdownDescription: "A descriptive summary for the configuration.",
				Computed:            true,
			},
			"read_only": schema.BoolAttribute{
				MarkdownDescription: "<p>The access permission to files on the storage volume.</p><ul><li>When enabled, access permission is read-only. With the B2B feature, access permission must be read/write. B2B storage requires write access to the RAID volume. Setting to read-only is ignored but generates a warning.</li><li>When disabled, the default value, access permission is read/write.</li></ul>",
				Computed:            true,
			},
			"directory": schema.StringAttribute{
				MarkdownDescription: "The subdirectory where the files on the storage volume are available. The name can be up to 64 characters long. The name cannot start with a period. The name can use all alphanumeric characters and the following special characters: . - _.",
				Computed:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (d *RaidVolumeDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *RaidVolumeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data models.RaidVolume
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := d.pData.Client.Get(data.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	data.FromBody(ctx, `RaidVolume`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
