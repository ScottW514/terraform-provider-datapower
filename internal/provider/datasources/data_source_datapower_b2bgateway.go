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
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
)

type B2BGatewayList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &B2BGatewayDataSource{}
	_ datasource.DataSourceWithConfigure = &B2BGatewayDataSource{}
)

func NewB2BGatewayDataSource() datasource.DataSource {
	return &B2BGatewayDataSource{}
}

type B2BGatewayDataSource struct {
	client *client.DatapowerClient
}

func (d *B2BGatewayDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_b2bgateway"
}

func (d *B2BGatewayDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "B2B gateway",
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
						"priority": schema.StringAttribute{
							MarkdownDescription: "Service priority",
							Computed:            true,
						},
						"doc_store_location": schema.StringAttribute{
							MarkdownDescription: "Document storage location",
							Computed:            true,
						},
						"as_front_protocol": schema.ListNestedAttribute{
							MarkdownDescription: "Protocol handlers",
							NestedObject:        models.DmASFrontProtocolDataSourceSchema,
							Computed:            true,
						},
						"as1mdn_email": schema.StringAttribute{
							MarkdownDescription: "Default AS1 MDN return email",
							Computed:            true,
						},
						"as1mdnsmtp_server_connection": schema.StringAttribute{
							MarkdownDescription: "AS1 MDN SMTP server connection",
							Computed:            true,
						},
						"as2mdnurl": schema.StringAttribute{
							MarkdownDescription: "Default AS2 MDN Return URL",
							Computed:            true,
						},
						"as3mdnurl": schema.StringAttribute{
							MarkdownDescription: "Default AS3 MDN return URL",
							Computed:            true,
						},
						"b2b_profiles": schema.ListNestedAttribute{
							MarkdownDescription: "Active partner profiles",
							NestedObject:        models.DmB2BActiveProfileDataSourceSchema,
							Computed:            true,
						},
						"b2b_groups": schema.ListNestedAttribute{
							MarkdownDescription: "Active profile groups",
							NestedObject:        models.DmB2BActiveGroupDataSourceSchema,
							Computed:            true,
						},
						"document_routing_preprocessor_type": schema.StringAttribute{
							MarkdownDescription: "Processor type",
							Computed:            true,
						},
						"document_routing_preprocessor": schema.StringAttribute{
							MarkdownDescription: "File location",
							Computed:            true,
						},
						"document_routing_preprocessor_debug": schema.BoolAttribute{
							MarkdownDescription: "Enable GatewayScript debugger",
							Computed:            true,
						},
						"archive_mode": schema.StringAttribute{
							MarkdownDescription: "Purge mode",
							Computed:            true,
						},
						"archive_location": schema.StringAttribute{
							MarkdownDescription: "Archive location",
							Computed:            true,
						},
						"archive_file_name": schema.StringAttribute{
							MarkdownDescription: "Archive file base name",
							Computed:            true,
						},
						"archive_minimum_size": schema.Int64Attribute{
							MarkdownDescription: "Min size",
							Computed:            true,
						},
						"archive_document_age": schema.Int64Attribute{
							MarkdownDescription: "Document age",
							Computed:            true,
						},
						"archive_minimum_documents": schema.Int64Attribute{
							MarkdownDescription: "Min documents",
							Computed:            true,
						},
						"disk_use_check_interval": schema.Int64Attribute{
							MarkdownDescription: "Check interval",
							Computed:            true,
						},
						"max_document_disk_use": schema.Int64Attribute{
							MarkdownDescription: "Max document storage",
							Computed:            true,
						},
						"archive_monitor": schema.BoolAttribute{
							MarkdownDescription: "Monitor during archival",
							Computed:            true,
						},
						"shaping_threshold": schema.Int64Attribute{
							MarkdownDescription: "Monitor threshold",
							Computed:            true,
						},
						"archive_backup_documents": models.GetDmB2BBackupMsgTypeDataSourceSchema("Document types to archive", "arch-backup-documents", ""),
						"x_path_routing_policies": schema.ListAttribute{
							MarkdownDescription: "XPath routing policies",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"xml_manager": schema.StringAttribute{
							MarkdownDescription: "XML manager",
							Computed:            true,
						},
						"debug_mode": schema.StringAttribute{
							MarkdownDescription: "Probe setting",
							Computed:            true,
						},
						"debug_history": schema.Int64Attribute{
							MarkdownDescription: "Transaction history",
							Computed:            true,
						},
						"cpa_entries": schema.ListNestedAttribute{
							MarkdownDescription: "CPA",
							NestedObject:        models.DmB2BCPAEntryDataSourceSchema,
							Computed:            true,
						},
						"sql_data_source": schema.StringAttribute{
							MarkdownDescription: "SQL data source",
							Computed:            true,
						},
						"front_side_timeout": schema.Int64Attribute{
							MarkdownDescription: "Front timeout",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *B2BGatewayDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *B2BGatewayDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data B2BGatewayList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.B2BGateway{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.B2BGateway{}
	if value := res.Get(`B2BGateway`); value.Exists() {
		for _, v := range value.Array() {
			item := models.B2BGateway{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.B2BGatewayObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.B2BGatewayObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
