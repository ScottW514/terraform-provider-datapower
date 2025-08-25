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
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type ODRConnectorGroupList struct {
	Result types.List `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &ODRConnectorGroupDataSource{}
	_ datasource.DataSourceWithConfigure = &ODRConnectorGroupDataSource{}
)

func NewODRConnectorGroupDataSource() datasource.DataSource {
	return &ODRConnectorGroupDataSource{}
}

type ODRConnectorGroupDataSource struct {
	pData *tfutils.ProviderData
}

func (d *ODRConnectorGroupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_odr_connector_group"
}

func (d *ODRConnectorGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "A collection of on demand router (ODR) connectors used to communicate with the Intelligent Management service.",
		Attributes: map[string]schema.Attribute{
			"result": schema.ListNestedAttribute{
				MarkdownDescription: "List of objects",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Name of the object. Must be unique among object types in application domain.",
							Computed:            true,
						},
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"odr_group_connectors": schema.ListNestedAttribute{
							MarkdownDescription: "The ODR connectors that are used to retrieve ODR information. An ODR connector defines a connection to the Intelligent Management service. The DataPower Gateway retrieves topology, weights, session affinity, and other information from the WebSphere cell over the connection. If multiple connectors are configured, the top most connection is tried first, followed by the second, and so on. After a connection is established, other endpoints might be retrieved and used if the configured connector endpoint is down. You must define at least one connector in an ODR connector group but cannot define more than 16 connectors.",
							NestedObject:        models.GetDmODRConnectorDataSourceSchema(),
							Computed:            true,
						},
						"max_retry_interval": schema.Int64Attribute{
							MarkdownDescription: "The time to wait before attempting to reestablish a connection to the Intelligent Management Service. Enter a value in the range 1 - 120. The default value is 60.",
							Computed:            true,
						},
						"xml_manager": schema.StringAttribute{
							MarkdownDescription: "The XML manager to use when processing transactions with the Intelligent management service.",
							Computed:            true,
						},
						"odr_conn_group_properties": schema.ListNestedAttribute{
							MarkdownDescription: "The custom properties that are associated with the ODR connector group.",
							NestedObject:        models.GetDmODRConnPropertyDataSourceSchema(),
							Computed:            true,
						},
						"ssl_client_config_type": schema.StringAttribute{
							MarkdownDescription: "The TLS profile type to secure connections between the DataPower Gateway and its targets.",
							Computed:            true,
						},
						"ssl_client": schema.StringAttribute{
							MarkdownDescription: "The TLS client profile to secure connections between the DataPower Gateway and its targets.",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *ODRConnectorGroupDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *ODRConnectorGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data ODRConnectorGroupList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.ODRConnectorGroup{}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.ODRConnectorGroup{}
	if value := res.Get(`ODRConnectorGroup`); value.Exists() {
		for _, v := range value.Array() {
			item := models.ODRConnectorGroup{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.ODRConnectorGroupObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.ODRConnectorGroupObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
