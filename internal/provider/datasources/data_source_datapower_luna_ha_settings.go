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
	_ datasource.DataSource              = &LunaHASettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &LunaHASettingsDataSource{}
)

func NewLunaHASettingsDataSource() datasource.DataSource {
	return &LunaHASettingsDataSource{}
}

type LunaHASettingsDataSource struct {
	pData *tfutils.ProviderData
}

func (d *LunaHASettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_luna_ha_settings"
}

func (d *LunaHASettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Defines the high availability (HA) settings for the SafeNet Luna Network HSM HA group.",
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "<p>The administrative state of the configuration.</p><ul><li>To make active, set to enabled.</li><li>To make inactive, set to disabled.</li></ul>",
				Computed:            true,
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "Specify the HA recovery mode. The default value is activeBasic.",
				Computed:            true,
			},
			"recovery_count": schema.Int64Attribute{
				MarkdownDescription: "Specify the number of recovery attempts for failed members in the HA group. Enter a value in the range 0 - 500. The default value is 0, which disables automatic recovery.",
				Computed:            true,
			},
			"interval": schema.Int64Attribute{
				MarkdownDescription: "Specify the interval in seconds between recovery attempts for failed members in the HA group. Enter a value in the range 60 - 1200. The default value is 60.",
				Computed:            true,
			},
		},
	}
}

func (d *LunaHASettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *LunaHASettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data models.LunaHASettings
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
		data.FromBody(ctx, `LunaHASettings`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
