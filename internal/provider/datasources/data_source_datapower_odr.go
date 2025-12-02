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
	_ datasource.DataSource              = &ODRDataSource{}
	_ datasource.DataSourceWithConfigure = &ODRDataSource{}
)

func NewODRDataSource() datasource.DataSource {
	return &ODRDataSource{}
}

type ODRDataSource struct {
	pData *tfutils.ProviderData
}

func (d *ODRDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_odr"
}

func (d *ODRDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "<p>Specifies an on demand router (ODR). The ODR feature acts as a collection of load balancer groups that distribute traffic to various clusters within a WebSphere cell or Liberty Collective. If multiple ODR connector groups are defined, the ODR distributes traffic to any of the clusters.</p><p>The ODR feature on the DataPower Gateway supports a subset of On Demand Router in Intelligent Management.</p>",
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
				MarkdownDescription: "Enter a descriptive summary for the configuration.",
				Computed:            true,
			},
			"odr_server_name": schema.StringAttribute{
				MarkdownDescription: "<p>Enter the name of the ODR routing rules server that is used to determine whether routing rules are delivered to the DataPower Gateway. If the DataPower Gateway is not using routing rules, you can use any server name.</p><p>You must define the <tt>RoutingRulesConnectorClusterName</tt> custom property to set the name of the ODR connector group from which routing rules are accepted.</p><p>The web server name is configured in on one of the following ways. <ul><li>In a Liberty Collective, defined as an attribute of the <tt>routingRules</tt> element.</li><li>In a WebSphere cell as the value of the <tt>serverName</tt> parameter of <tt>WebServerRoutingRule</tt> command.</li></ul></p>",
				Computed:            true,
			},
			"odr_connector_groups": schema.ListAttribute{
				MarkdownDescription: "Define one ODR connector group for each WebSphere cell or Liberty Collective. Each connector group establishes the communications with an Intelligent management service. The DataPower Gateway retrieves topology information, application information, routing rules, and other information over the connectors in the connector group.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"odr_custom_properties": schema.ListNestedAttribute{
				MarkdownDescription: "<p>Defines custom property name-value strings to connect to Liberty Collective or WebSphere cell to define which ODR connection group accepts routing rules.</p><ul><li>The scheme that the DataPower ODR feature uses to connect to WebSphere Application Server does not apply to Liberty Collective. For any HTTP or HTTPS request to connect to Liberty Collective, you must define the <tt>profileType</tt> property. <ol><li>Set the <b>Name</b> to <tt>profileType</tt> .</li><li>Set the <b>Value</b> to <tt>Liberty</tt> .</li></ol></li><li>To set the name of the ODR connector group from which routing rules are accepted, you must define the <tt>RoutingRulesConnectorClusterName</tt> property. The following example shows setting the <tt>liberty-collective-node03</tt> connector group as the group that accepts routing rules. <ol><li>Set the <b>Name</b> to <tt>RoutingRulesConnectorClusterName</tt> .</li><li>Set the <b>Value</b> to <tt>liberty-collective-node03</tt> .</li></ol></li></ul><p>Beyond these situations, use custom properties only when directed by IBM Support.</p>",
				NestedObject:        models.GetDmODRPropertyDataSourceSchema(),
				Computed:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (d *ODRDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *ODRDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data models.ODR
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
		data.FromBody(ctx, `ODR`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
