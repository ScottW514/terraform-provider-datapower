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
	_ datasource.DataSource              = &GWScriptSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &GWScriptSettingsDataSource{}
)

func NewGWScriptSettingsDataSource() datasource.DataSource {
	return &GWScriptSettingsDataSource{}
}

type GWScriptSettingsDataSource struct {
	pData *tfutils.ProviderData
}

func (d *GWScriptSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gw_script_settings"
}

func (d *GWScriptSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "You can configure the following GatewayScript settings. <ul><li>Whether to freeze the GatewayScript built-in objects.</li><li>How to manage untrusted code mitigation to protect against SSCA.</li><li>The maximum allowed duration that a GatewayScript action can continuously use the CPU.</li></ul>",
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "<p>The administrative state of the configuration.</p><ul><li>To make active, set to enabled.</li><li>To make inactive, set to disabled.</li></ul>",
				Computed:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: "Comments",
				Computed:            true,
			},
			"frozen_enabled": schema.BoolAttribute{
				MarkdownDescription: "Specify whether to freeze the GatewayScript built-in object prototypes. <ul><li>When frozen, you cannot modify, add, or remove prototypes. This setting is the default setting.</li><li>When not frozen, you can manipulate prototypes.</li></ul><p>If you change and persist this property, the change is pending and requires a firmware reload.</p>",
				Computed:            true,
			},
			"untrusted_code_mitigated": schema.BoolAttribute{
				MarkdownDescription: "Specify whether to enable untrusted code mitigation to protect against Speculative Side-Channel Attacks (SSCA). <ul><li>When enabled, protects untrusted code against SSCA. This setting is the default setting.</li><li>When disabled, does not protect untrusted code against SSCA.</li></ul><p>If you change and persist this property, the change is pending and requires a firmware reload.</p>",
				Computed:            true,
			},
			"reload_needed": schema.BoolAttribute{
				MarkdownDescription: "System reload needed",
				Computed:            true,
			},
			"terminate_time": schema.Int64Attribute{
				MarkdownDescription: "Specify the maximum duration in seconds that a GatewayScript action can continuously use CPU without yielding back to the system event loop. When the processing of a GatewayScript action exceeds the duration, processing is stopped and an error is logged. Enter a value in the range 1 - 300. The default value is 0, which indicates unlimited.",
				Computed:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (d *GWScriptSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *GWScriptSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data models.GWScriptSettings
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := d.pData.Client.Get(data.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	data.FromBody(ctx, `GWScriptSettings`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
