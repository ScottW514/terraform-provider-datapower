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
	resp.TypeName = req.ProviderTypeName + "_sql_data_source"
}

func (d *SQLDataSourceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "<p>An SQL data source the configuration to establish a direct connection to a database instance on a remote data server. When configured, a DataPower service can dynamically run database operations, such as <tt>SELECT</tt> and <tt>INSERT</tt> , on the remote database.</p><p>An SQL data source is used by SQL actions in processing policies. The SQL action retrieves the data for further processing. Conversely, the processing policy can store the processed data in the configured database instance.</p><p>When you configure an SQL data source, you can define valid configuration parameters for your data server connection. Configuration parameters modify the behavior of the services that run with a data server. Some configuration parameters in the configuration file are informational and define characteristics about the environment. These configuration parameters cannot be modified.</p>",
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
							MarkdownDescription: "Specify the user to establish the connection to the SQL database. The server maintains this information.",
							Computed:            true,
						},
						"password_alias": schema.StringAttribute{
							MarkdownDescription: "Specify the password alias of the user password to establish connection with the SQL database. The password alias looks up the password for the user. The server maintains the password.",
							Computed:            true,
						},
						"data_source_id": schema.StringAttribute{
							MarkdownDescription: "Specify the identifier of the data source. The terminology differs by vendor. <ul><li><b>Db2</b> - The IBM Db2 database alias.</li><li><b>IMS</b> - The name of the IBM IMS data store.</li><li><b>Microsoft SQL Server</b> - The name of the Microsoft SQL Server data source.</li><li><b>Oracle</b> - The Oracle system identifier (SID) or service name.</li><li><b>Sybase</b> - The name of the Sybase database.</li></ul>",
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
							MarkdownDescription: "Specify whether to limit the data from a <b>SELECT</b> statement. By default, the response size is not limited.",
							Computed:            true,
						},
						"limit_returned_data_size": schema.Int64Attribute{
							MarkdownDescription: "Specify the limit in KB on returned data from a <b>SELECT</b> statement. The default value is 128.",
							Computed:            true,
						},
						"sql_data_source_config_nv_pairs": schema.ListNestedAttribute{
							MarkdownDescription: "Specify configuration parameters for the data server connection. Configuration parameters modify the behavior of the services that run with a data server. Some parameters in the configuration file are informational and define characteristics about the environment. These parameters cannot be modified.",
							NestedObject:        models.GetDmSQLDataSourceConfigNVPairDataSourceSchema(),
							Computed:            true,
						},
						"max_connection": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum number of concurrent SQL connections. Enter a value in the range 1 - 65535. The default value is 10.",
							Computed:            true,
						},
						"oracle_data_source_type": schema.StringAttribute{
							MarkdownDescription: "Data source type - Oracle",
							Computed:            true,
						},
						"connect_timeout": schema.Int64Attribute{
							MarkdownDescription: "<p>Specify the duration in seconds to wait to establish a connection to the data server. Enter a value in the range 0 - 4294967295. The value of 0 disables the timeout. The default value is 15.</p><p>A new connection is the initial connection and each new connection from the connection pool. Reuse of a connection from the connection pool is not considered establishing a new connection.</p><p>The connection timeout must be less than the query timeout. With this configuration, the initial query has time to establish the connection to the data server.</p>",
							Computed:            true,
						},
						"query_timeout": schema.Int64Attribute{
							MarkdownDescription: "Specify the duration in seconds to wait for an SQL request to complete. Enter a value in the range 0 - 4294967295. The default value is 0, which uses the standard timeout in the user agent. <p>The duration is from when the service sends the request to when the service receives the results.</p><p>The query timeout must be greater than the connection timeout. With this configuration, the initial query has time to establish the connection to the remote data server.</p>",
							Computed:            true,
						},
						"idle_timeout": schema.Int64Attribute{
							MarkdownDescription: "Specify the duration that a connection from the connection pool can remain idle before the connection is released. Enter a value in the range 0 - 4294967295. The default value is 180. The value of 0 disables the timer.",
							Computed:            true,
						},
						"load_balancing": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enable Db2 workload balancing and automatic client reroute for Db2 for z/OS. <p>When enabled, this feature set uses the z/OS Sysplex Distributor for real-time load distribution of SQL calls to the sysplex-aware Db2 instance.</p><p>When enabled, you must specify the sysplex DVIPA as the data source host.</p>",
							Computed:            true,
						},
						"encryption_method_mssql": schema.StringAttribute{
							MarkdownDescription: "Specify the TLS encryption method for a Microsoft SQL Server database. When the server does not support the specified encryption method, the connection fails.",
							Computed:            true,
						},
						"encryption_method_oracle": schema.StringAttribute{
							MarkdownDescription: "Specify the TLS encryption method for an Oracle database. When the server does not support the specified encryption method, the connection fails. The default behavior is to not encrypt or decrypt data.",
							Computed:            true,
						},
						"encryption_method_db2": schema.StringAttribute{
							MarkdownDescription: "Specify the TLS encryption method for an IBM Db2 database. When the server does not support the specified encryption method, the connection fails. The default behavior is to not encrypt or decrypt data.",
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
							MarkdownDescription: "Specify the hostname that the certificate must contain for hostname validation. Hostname validation provides extra security against man-in-the-middle (MITM) attacks by ensuring that the connection is to the requested server.",
							Computed:            true,
						},
						"validate_host_name": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to validate the hostname against the hostname in the server certificate. Hostname validate uses the value of the data source host. Hostname validation provides extra security against man-in-the-middle (MITM) attacks by ensuring that the connection is to the requested server.",
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
