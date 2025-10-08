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
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type CountMonitorList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &CountMonitorDataSource{}
	_ datasource.DataSourceWithConfigure = &CountMonitorDataSource{}
)

func NewCountMonitorDataSource() datasource.DataSource {
	return &CountMonitorDataSource{}
}

type CountMonitorDataSource struct {
	pData *tfutils.ProviderData
}

func (d *CountMonitorDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_count_monitor"
}

func (d *CountMonitorDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "On-Line Help for Message-Count Monitors",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The name of the object to retrieve.",
				Optional:            true,
			},
			"app_domain": schema.StringAttribute{
				MarkdownDescription: "The name of the application domain the object belongs to.",
				Required:            true,
			},
			"result": schema.ListNestedAttribute{
				MarkdownDescription: "List of objects. If `id` was provided and it exists, it will be the only item in the list.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Name of the object. Must be unique among object types in application domain.",
							Computed:            true,
						},
						"app_domain": schema.StringAttribute{
							MarkdownDescription: "The name of the application domain the object belongs to",
							Computed:            true,
						},
						"measure": schema.StringAttribute{
							MarkdownDescription: "Select the action that advances the counter. The default is Requests.",
							Computed:            true,
						},
						"source": schema.StringAttribute{
							MarkdownDescription: "Select how monitoring is applied to IP addresses. This setting is meaningful only if an associated traffic definition contains and inclusive or exclusive IP address criterion. The default is All.",
							Computed:            true,
						},
						"header": schema.StringAttribute{
							MarkdownDescription: "The name of the HTTP header to read to determine the value of the source IP address.",
							Computed:            true,
						},
						"filter": schema.ListNestedAttribute{
							MarkdownDescription: "Click this tab to define a message-count monitor threshold and assign an action (Message Filter Action) that is taken when the threshold is reached.",
							NestedObject:        models.GetDmCountMonitorFilterDataSourceSchema(),
							Computed:            true,
						},
						"max_source_s": schema.Int64Attribute{
							MarkdownDescription: "When utilizing the each-ip aggregate addressing policy the system organizes the counts per address by the addresses most recently used. When too many distinct counts have been observed, the Addresses not seen in the longest time are discarded. This parameter specifies how many distinct addresses are tracked.",
							Computed:            true,
						},
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "A Message Monitor observes traffic that is incuded by the Message Type definition (which in turn is a collection of Message Matching objects). The Monitor measures only that traffic selected by the Measure field. On the Filters page, traffic which meets the filter criteria causes the Monitor to take the corresponding action (which is defined by a Message Filter Type object).",
							Computed:            true,
						},
						"message_type": schema.StringAttribute{
							MarkdownDescription: "Select the message type monitored by this message-count monitor.",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *CountMonitorDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *CountMonitorDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data CountMonitorList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.CountMonitor{
		AppDomain: data.AppDomain,
	}

	path := o.GetPath()
	if !data.Id.IsNull() {
		path = path + "/" + data.Id.ValueString()
	}

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
	l := []models.CountMonitor{}
	if resFound {
		if value := res.Get(`CountMonitor`); value.Exists() {
			for _, v := range value.Array() {
				item := models.CountMonitor{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.CountMonitorObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.CountMonitorObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
