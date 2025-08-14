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

type SQLDataSourceList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &SQLDataSourceDataSource{}
	_ datasource.DataSourceWithConfigure = &SQLDataSourceDataSource{}
)

func NewSQLDataSourceDataSource() datasource.DataSource {
	return &SQLDataSourceDataSource{}
}

type SQLDataSourceDataSource struct {
	pData *tfutils.ProviderData
}

func (d *SQLDataSourceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sqldatasource"
}

func (d *SQLDataSourceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SQL data source",
		Attributes: map[string]schema.Attribute{
			"app_domain": schema.StringAttribute{
				MarkdownDescription: "The name of the application domain the object belongs to",
				Required:            true,
			},
			"result": schema.ListNestedAttribute{
				MarkdownDescription: "List of objects",
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
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"database": schema.StringAttribute{
							MarkdownDescription: "Database type",
							Computed:            true,
						},
						"username": schema.StringAttribute{
							MarkdownDescription: "Connection username",
							Computed:            true,
						},
						"password_alias": schema.StringAttribute{
							MarkdownDescription: "Connection password alias",
							Computed:            true,
						},
						"data_source_id": schema.StringAttribute{
							MarkdownDescription: "Data source ID",
							Computed:            true,
						},
						"data_source_host": schema.StringAttribute{
							MarkdownDescription: "Data source host",
							Computed:            true,
						},
						"data_source_port": schema.Int64Attribute{
							MarkdownDescription: "Data source port",
							Computed:            true,
						},
						"limit_returned_data": schema.BoolAttribute{
							MarkdownDescription: "Limit returned data",
							Computed:            true,
						},
						"limit_returned_data_size": schema.Int64Attribute{
							MarkdownDescription: "Returned data size limit",
							Computed:            true,
						},
						"sql_data_source_config_nv_pairs": schema.ListNestedAttribute{
							MarkdownDescription: "Advanced configuration parameters",
							NestedObject:        models.DmSQLDataSourceConfigNVPairDataSourceSchema,
							Computed:            true,
						},
						"max_connection": schema.Int64Attribute{
							MarkdownDescription: "Max connections",
							Computed:            true,
						},
						"oracle_data_source_type": schema.StringAttribute{
							MarkdownDescription: "Data source type - Oracle",
							Computed:            true,
						},
						"connect_timeout": schema.Int64Attribute{
							MarkdownDescription: "Connection timeout",
							Computed:            true,
						},
						"query_timeout": schema.Int64Attribute{
							MarkdownDescription: "Query timeout",
							Computed:            true,
						},
						"idle_timeout": schema.Int64Attribute{
							MarkdownDescription: "Idle connection timeout",
							Computed:            true,
						},
						"load_balancing": schema.BoolAttribute{
							MarkdownDescription: "Load distribution - Db2 for z/OS",
							Computed:            true,
						},
						"encryption_method_mssql": schema.StringAttribute{
							MarkdownDescription: "Encryption method - SQL Server",
							Computed:            true,
						},
						"encryption_method_oracle": schema.StringAttribute{
							MarkdownDescription: "Encryption method - Oracle",
							Computed:            true,
						},
						"encryption_method_db2": schema.StringAttribute{
							MarkdownDescription: "Encryption method - Db2",
							Computed:            true,
						},
						"truststore_ref": schema.StringAttribute{
							MarkdownDescription: "Truststore",
							Computed:            true,
						},
						"validate_server_certificate": schema.StringAttribute{
							MarkdownDescription: "Validate server certificate",
							Computed:            true,
						},
						"host_name_in_certificate": schema.StringAttribute{
							MarkdownDescription: "Hostname in certificate",
							Computed:            true,
						},
						"validate_host_name": schema.BoolAttribute{
							MarkdownDescription: "Validate hostname",
							Computed:            true,
						},
						"keystore_ref": schema.StringAttribute{
							MarkdownDescription: "Keystore",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *SQLDataSourceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *SQLDataSourceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data SQLDataSourceList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.SQLDataSource{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.SQLDataSource{}
	if value := res.Get(`SQLDataSource`); value.Exists() {
		for _, v := range value.Array() {
			item := models.SQLDataSource{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.SQLDataSourceObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.SQLDataSourceObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
