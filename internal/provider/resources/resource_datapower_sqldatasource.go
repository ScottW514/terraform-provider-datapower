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
		MarkdownDescription: tfutils.NewAttributeDescription("SQL data source", "sql-source", "").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Connection username", "username", "").String,
				Required:            true,
			},
			"password_alias": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Connection password alias", "password-alias", "passwordalias").String,
				Required:            true,
			},
			"data_source_id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Data source ID", "id", "").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Limit returned data", "limit", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"limit_returned_data_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Returned data size limit", "limit-size", "").AddIntegerRange(1, 65535).AddDefaultValue("128").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(128),
			},
			"sql_data_source_config_nv_pairs": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Advanced configuration parameters", "sql-config-param", "").String,
				NestedObject:        models.DmSQLDataSourceConfigNVPairResourceSchema,
				Optional:            true,
			},
			"max_connection": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Max connections", "maximum-connections", "").AddIntegerRange(1, 65535).AddDefaultValue("10").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Connection timeout", "connect-timeout", "").AddDefaultValue("15").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(15),
			},
			"query_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Query timeout", "query-timeout", "").String,
				Required:            true,
			},
			"idle_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Idle connection timeout", "idle-timeout", "").AddDefaultValue("180").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(180),
			},
			"load_balancing": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Load distribution - Db2 for z/OS", "load-balancing", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"encryption_method_mssql": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Encryption method - SQL Server", "mssql-encryption-method", "").AddStringEnum("NoEncryption", "SSL", "RequestSSL", "LoginSSL").AddDefaultValue("NoEncryption").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("NoEncryption", "SSL", "RequestSSL", "LoginSSL"),
				},
				Default: stringdefault.StaticString("NoEncryption"),
			},
			"encryption_method_oracle": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Encryption method - Oracle", "oracle-encryption-method", "").AddStringEnum("NoEncryption", "SSL").AddDefaultValue("NoEncryption").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("NoEncryption", "SSL"),
				},
				Default: stringdefault.StaticString("NoEncryption"),
			},
			"encryption_method_db2": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Encryption method - Db2", "db2-encryption-method", "").AddStringEnum("NoEncryption", "SSL").AddDefaultValue("NoEncryption").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Hostname in certificate", "hostname-in-certificate", "").String,
				Optional:            true,
			},
			"validate_host_name": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validate hostname", "validate-host-name", "").AddDefaultValue("true").String,
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
