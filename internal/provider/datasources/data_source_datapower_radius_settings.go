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
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var (
	_ datasource.DataSource              = &RADIUSSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &RADIUSSettingsDataSource{}
)

func NewRADIUSSettingsDataSource() datasource.DataSource {
	return &RADIUSSettingsDataSource{}
}

type RADIUSSettingsDataSource struct {
	pData *tfutils.ProviderData
}

func (d *RADIUSSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_radius_settings"
}

func (d *RADIUSSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "RADIUS settings define RADIUS servers. You can use RADIUS servers to authenticate access with RBM or in a AAA policy. The DataPower Gateway is a client to RADIUS servers.",
		Attributes: map[string]schema.Attribute{
			"provider_target": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Target host to retrieve this data from. If not set, provider will use the top level settings.", "", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "<p>The administrative state of the configuration.</p><ul><li>To make active, set to enabled.</li><li>To make inactive, set to disabled.</li></ul>",
				Computed:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: "Comments",
				Computed:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"timeout": schema.Int64Attribute{
				MarkdownDescription: "Specify the RADIUS retransmit interval in milliseconds. This timeout is the duration that the RADIUS client waits before an unacknowledged request is retransmitted. Enter a value in the range of 1 - 30000. The default value is 1000.",
				Computed:            true,
			},
			"retries": schema.Int64Attribute{
				MarkdownDescription: "Specify the maximum number of times that the RADIUS client can retransmit an unacknowledged request to a server. Enter a value in the range 1 - 10. The default value is 3.",
				Computed:            true,
			},
			"aaa_servers": schema.ListNestedAttribute{
				MarkdownDescription: "AAA and RBM",
				NestedObject:        models.GetDmRadiusServerDataSourceSchema(),
				Computed:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (d *RADIUSSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *RADIUSSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data models.RADIUSSettingsWO
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !d.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), d.pData.Client.GetTargetNames()))
		return
	}

	path := data.GetPath()

	res, err := d.pData.Client.Get(path, data.ProviderTarget)
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
		data.FromBody(ctx, `RADIUSSettings`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
