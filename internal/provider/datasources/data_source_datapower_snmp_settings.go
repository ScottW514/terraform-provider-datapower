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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var (
	_ datasource.DataSource              = &SNMPSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &SNMPSettingsDataSource{}
)

func NewSNMPSettingsDataSource() datasource.DataSource {
	return &SNMPSettingsDataSource{}
}

type SNMPSettingsDataSource struct {
	pData *tfutils.ProviderData
}

func (d *SNMPSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmp_settings"
}

func (d *SNMPSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Use this page to establish SNMP connectivity to the device, and to set values used by SNMP.",
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
			"local_address": schema.StringAttribute{
				MarkdownDescription: "A specific IP address monitored by the SNMP agent or engine for incoming SNMP requests. The default value of 0.0.0.0 allows the agent or engine to listen on all interfaces. Selecting the address of one interface restricts SNMP to that interface.",
				Computed:            true,
			},
			"local_port": schema.Int64Attribute{
				MarkdownDescription: "A specific UDP port monitored by the SNMP agent or engine for incoming SNMP requests. By default, the agent or engine monitors port 161.",
				Computed:            true,
			},
			"policies": schema.ListNestedAttribute{
				MarkdownDescription: "SNMPv1/v2c Communities",
				NestedObject:        models.GetDmSnmpPolicyDataSourceSchema(),
				Computed:            true,
			},
			"policies_mq": schema.ListNestedAttribute{
				MarkdownDescription: "SNMPv1/v2c Communities",
				NestedObject:        models.GetDmSnmpPolicyMQDataSourceSchema(),
				Computed:            true,
			},
			"targets": schema.ListNestedAttribute{
				MarkdownDescription: "Trap and Notification Targets",
				NestedObject:        models.GetDmSnmpTargetDataSourceSchema(),
				Computed:            true,
			},
			"users": schema.ListAttribute{
				MarkdownDescription: "The name of a user (which must have SNMP credential parameters) which is authorized to use SNMPv3 to access the MIBs on this system.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"contexts": schema.ListNestedAttribute{
				MarkdownDescription: "SNMPv3 context definitions, which provide SNMPv3 access to non-default application domains.",
				NestedObject:        models.GetDmSnmpContextDataSourceSchema(),
				Computed:            true,
			},
			"security_level": schema.StringAttribute{
				MarkdownDescription: "The minimum security level required for incoming SNMPv3 Get and Set requests. The default is Authentication, Privacy.",
				Computed:            true,
			},
			"access_level": schema.StringAttribute{
				MarkdownDescription: "The type of access allowed to MIB objects for incoming SNMPv3 Get and Set requests.",
				Computed:            true,
			},
			"enable_default_trap_subscriptions": schema.BoolAttribute{
				MarkdownDescription: "Enable or Disable the default list of event codes that generate traps. The default is Enable Trap Subscriptions.",
				Computed:            true,
			},
			"trap_priority": schema.StringAttribute{
				MarkdownDescription: "Select a minimum trap event priority. The priorities are hierarchical. The lowest is listed last. Set to the minimum that is required for your trap events.",
				Computed:            true,
			},
			"trap_event_code": schema.ListAttribute{
				MarkdownDescription: "The list of event codes generating traps. You can add event codes which will be triggering traps send to the configured trap targets.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"config_mib": schema.StringAttribute{
				MarkdownDescription: "Configuration",
				Computed:            true,
			},
			"config_mib_mq": schema.StringAttribute{
				MarkdownDescription: "Configuration",
				Computed:            true,
			},
			"status_mib": schema.StringAttribute{
				MarkdownDescription: "Status",
				Computed:            true,
			},
			"status_mib_mq": schema.StringAttribute{
				MarkdownDescription: "Status",
				Computed:            true,
			},
			"notif_mib": schema.StringAttribute{
				MarkdownDescription: "Notifications",
				Computed:            true,
			},
			"notif_mib_mq": schema.StringAttribute{
				MarkdownDescription: "Notifications",
				Computed:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (d *SNMPSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *SNMPSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data models.SNMPSettings
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
		data.FromBody(ctx, `SNMPSettings`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
