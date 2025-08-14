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

package resources

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &SQLDataSourceResource{}

func NewSQLDataSourceResource() resource.Resource {
	return &SQLDataSourceResource{}
}

type SQLDataSourceResource struct {
	pData *tfutils.ProviderData
}

func (r *SQLDataSourceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sqldatasource"
}

func (r *SQLDataSourceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("<p>An SQL data source the configuration to establish a direct connection to a database instance on a remote data server. When configured, a DataPower service can dynamically run database operations, such as <tt>SELECT</tt> and <tt>INSERT</tt> , on the remote database.</p><p>An SQL data source is used by SQL actions in processing policies. The SQL action retrieves the data for further processing. Conversely, the processing policy can store the processed data in the configured database instance.</p><p>When you configure an SQL data source, you can define valid configuration parameters for your data server connection. Configuration parameters modify the behavior of the services that run with a data server. Some configuration parameters in the configuration file are informational and define characteristics about the environment. These configuration parameters cannot be modified.</p>", "sql-source", "").String,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name of the object. Must be unique among object types in application domain.", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"app_domain": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The name of the application domain the object belongs to", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`), ""),
				},
				PlanModifiers: []planmodifier.String{
					modifiers.ImmutableAfterSet(),
				},
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"database": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Database type", "db", "").AddStringEnum("DB2", "Oracle", "Sybase", "MSSQLServer", "DB2v9", "IMS").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("DB2", "Oracle", "Sybase", "MSSQLServer", "DB2v9", "IMS"),
				},
			},
			"username": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the user to establish the connection to the SQL database. The server maintains this information.", "username", "").String,
				Required:            true,
			},
			"password_alias": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the password alias of the user password to establish connection with the SQL database. The password alias looks up the password for the user. The server maintains the password.", "password-alias", "passwordalias").String,
				Required:            true,
			},
			"data_source_id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the identifier of the data source. The terminology differs by vendor. <ul><li><b>Db2</b> - The IBM Db2 database alias.</li><li><b>IMS</b> - The name of the IBM IMS data store.</li><li><b>Microsoft SQL Server</b> - The name of the Microsoft SQL Server data source.</li><li><b>Oracle</b> - The Oracle system identifier (SID) or service name.</li><li><b>Sybase</b> - The name of the Sybase database.</li></ul>", "id", "").String,
				Required:            true,
			},
			"data_source_host": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Data source host", "host", "").String,
				Required:            true,
			},
			"data_source_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Data source port", "port", "").String,
				Required:            true,
			},
			"limit_returned_data": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to limit the data from a <b>SELECT</b> statement. By default, the response size is not limited.", "limit", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"limit_returned_data_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the limit in KB on returned data from a <b>SELECT</b> statement. The default value is 128.", "limit-size", "").AddIntegerRange(1, 65535).AddDefaultValue("128").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(128),
			},
			"sql_data_source_config_nv_pairs": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify configuration parameters for the data server connection. Configuration parameters modify the behavior of the services that run with a data server. Some parameters in the configuration file are informational and define characteristics about the environment. These parameters cannot be modified.", "sql-config-param", "").String,
				NestedObject:        models.DmSQLDataSourceConfigNVPairResourceSchema,
				Optional:            true,
			},
			"max_connection": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum number of concurrent SQL connections. Enter a value in the range 1 - 65535. The default value is 10.", "maximum-connections", "").AddIntegerRange(1, 65535).AddDefaultValue("10").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(10),
			},
			"oracle_data_source_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Data source type - Oracle", "oracle-datasource-type", "").AddStringEnum("SID", "ServiceName").AddDefaultValue("SID").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("SID", "ServiceName"),
				},
				Default: stringdefault.StaticString("SID"),
			},
			"connect_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the duration in seconds to wait to establish a connection to the data server. Enter a value in the range 0 - 4294967295. The value of 0 disables the timeout. The default value is 15.</p><p>A new connection is the initial connection and each new connection from the connection pool. Reuse of a connection from the connection pool is not considered establishing a new connection.</p><p>The connection timeout must be less than the query timeout. With this configuration, the initial query has time to establish the connection to the data server.</p>", "connect-timeout", "").AddDefaultValue("15").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(15),
			},
			"query_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration in seconds to wait for an SQL request to complete. Enter a value in the range 0 - 4294967295. The default value is 0, which uses the standard timeout in the user agent. <p>The duration is from when the service sends the request to when the service receives the results.</p><p>The query timeout must be greater than the connection timeout. With this configuration, the initial query has time to establish the connection to the remote data server.</p>", "query-timeout", "").String,
				Required:            true,
			},
			"idle_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration that a connection from the connection pool can remain idle before the connection is released. Enter a value in the range 0 - 4294967295. The default value is 180. The value of 0 disables the timer.", "idle-timeout", "").AddDefaultValue("180").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(180),
			},
			"load_balancing": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable Db2 workload balancing and automatic client reroute for Db2 for z/OS. <p>When enabled, this feature set uses the z/OS Sysplex Distributor for real-time load distribution of SQL calls to the sysplex-aware Db2 instance.</p><p>When enabled, you must specify the sysplex DVIPA as the data source host.</p>", "load-balancing", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"encryption_method_mssql": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS encryption method for a Microsoft SQL Server database. When the server does not support the specified encryption method, the connection fails.", "mssql-encryption-method", "").AddStringEnum("NoEncryption", "SSL", "RequestSSL", "LoginSSL").AddDefaultValue("NoEncryption").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("NoEncryption", "SSL", "RequestSSL", "LoginSSL"),
				},
				Default: stringdefault.StaticString("NoEncryption"),
			},
			"encryption_method_oracle": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS encryption method for an Oracle database. When the server does not support the specified encryption method, the connection fails. The default behavior is to not encrypt or decrypt data.", "oracle-encryption-method", "").AddStringEnum("NoEncryption", "SSL").AddDefaultValue("NoEncryption").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("NoEncryption", "SSL"),
				},
				Default: stringdefault.StaticString("NoEncryption"),
			},
			"encryption_method_db2": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS encryption method for an IBM Db2 database. When the server does not support the specified encryption method, the connection fails. The default behavior is to not encrypt or decrypt data.", "db2-encryption-method", "").AddStringEnum("NoEncryption", "SSL").AddDefaultValue("NoEncryption").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("NoEncryption", "SSL"),
				},
				Default: stringdefault.StaticString("NoEncryption"),
			},
			"truststore_ref": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Truststore", "truststore", "cryptovalcred").String,
				Optional:            true,
			},
			"validate_server_certificate": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validate server certificate", "validate-server-certificate", "").AddStringEnum("Disabled", "Enabled").AddDefaultValue("Enabled").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Disabled", "Enabled"),
				},
				Default: stringdefault.StaticString("Enabled"),
			},
			"host_name_in_certificate": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the hostname that the certificate must contain for hostname validation. Hostname validation provides extra security against man-in-the-middle (MITM) attacks by ensuring that the connection is to the requested server.", "hostname-in-certificate", "").String,
				Optional:            true,
			},
			"validate_host_name": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to validate the hostname against the hostname in the server certificate. Hostname validate uses the value of the data source host. Hostname validation provides extra security against man-in-the-middle (MITM) attacks by ensuring that the connection is to the requested server.", "validate-host-name", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"keystore_ref": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Keystore", "keystore", "cryptoidentcred").String,
				Optional:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *SQLDataSourceResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *SQLDataSourceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SQLDataSource
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `SQLDataSource`)
	_, err := r.pData.Client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "POST", err))
		return
	}
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SQLDataSourceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SQLDataSource
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && (strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}

	if data.IsNull() {
		// Import
		data.FromBody(ctx, `SQLDataSource`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `SQLDataSource`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SQLDataSourceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SQLDataSource
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `SQLDataSource`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Update)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SQLDataSourceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SQLDataSource
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *SQLDataSourceResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.SQLDataSource

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
