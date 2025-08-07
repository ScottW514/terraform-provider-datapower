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

type TAMList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &TAMDataSource{}
	_ datasource.DataSourceWithConfigure = &TAMDataSource{}
)

func NewTAMDataSource() datasource.DataSource {
	return &TAMDataSource{}
}

type TAMDataSource struct {
	client *client.DatapowerClient
}

func (d *TAMDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tam"
}

func (d *TAMDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Access Manager Client",
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
						"ad_use_ad": schema.BoolAttribute{
							MarkdownDescription: "Use Active Directory",
							Computed:            true,
						},
						"tam_version": schema.StringAttribute{
							MarkdownDescription: "Access Manager Client Version",
							Computed:            true,
						},
						"configuration_file": schema.StringAttribute{
							MarkdownDescription: "Configuration File for Access Manager",
							Computed:            true,
						},
						"ad_configuration_file": schema.StringAttribute{
							MarkdownDescription: "Configuration File for Directories",
							Computed:            true,
						},
						"ssl_key_file": schema.StringAttribute{
							MarkdownDescription: "TLS Key File",
							Computed:            true,
						},
						"ssl_key_stash_file": schema.StringAttribute{
							MarkdownDescription: "TLS Key Stash File",
							Computed:            true,
						},
						"use_local_mode": schema.BoolAttribute{
							MarkdownDescription: "Use Local Policy Database",
							Computed:            true,
						},
						"poll_interval": schema.StringAttribute{
							MarkdownDescription: "Local Database Refresh Interval",
							Computed:            true,
						},
						"listen_mode": schema.BoolAttribute{
							MarkdownDescription: "Accept Update Notifications",
							Computed:            true,
						},
						"listen_port": schema.Int64Attribute{
							MarkdownDescription: "Update Notification Port",
							Computed:            true,
						},
						"returning_user_attributes": schema.BoolAttribute{
							MarkdownDescription: "Return User Attributes",
							Computed:            true,
						},
						"ldap_use_ssl": schema.BoolAttribute{
							MarkdownDescription: "Use TLS with Registry Server",
							Computed:            true,
						},
						"ldapssl_port": schema.Int64Attribute{
							MarkdownDescription: "LDAP TLS Port",
							Computed:            true,
						},
						"ldapssl_key_file": schema.StringAttribute{
							MarkdownDescription: "Registry Server TLS Key File",
							Computed:            true,
						},
						"ldapssl_key_file_password_alias": schema.StringAttribute{
							MarkdownDescription: "Registry Server TLS Key File Password Alias",
							Computed:            true,
						},
						"ldapssl_key_file_label": schema.StringAttribute{
							MarkdownDescription: "Registry Server TLS Key File Label",
							Computed:            true,
						},
						"tam_use_fips": schema.BoolAttribute{
							MarkdownDescription: "Run in FIPS Mode",
							Computed:            true,
						},
						"tam_choose_nist": schema.StringAttribute{
							MarkdownDescription: "Select a NIST Compliance Mode",
							Computed:            true,
						},
						"tam_use_basic_user": schema.BoolAttribute{
							MarkdownDescription: "Enable basic user mode",
							Computed:            true,
						},
						"user_principal_attribute": schema.StringAttribute{
							MarkdownDescription: "User principal attribute",
							Computed:            true,
						},
						"user_no_duplicates": schema.BoolAttribute{
							MarkdownDescription: "Disallow duplicate principals",
							Computed:            true,
						},
						"user_search_suffixes": schema.ListAttribute{
							MarkdownDescription: "Principal search suffixes",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"user_suffix_optimiser": schema.BoolAttribute{
							MarkdownDescription: "Enable suffix optimization",
							Computed:            true,
						},
						"tam_fed_dirs": schema.ListNestedAttribute{
							MarkdownDescription: "Federated directories",
							NestedObject:        models.DmTAMFedDirDataSourceSchema,
							Computed:            true,
						},
						"tamaz_replicas": schema.ListNestedAttribute{
							MarkdownDescription: "Authorization Server Replicas",
							NestedObject:        models.DmTAMAZReplicaDataSourceSchema,
							Computed:            true,
						},
						"tamras_trace": models.GetDmTAMRASTraceDataSourceSchema("Trace Logging", "tam-ras-trace", ""),
						"auto_retry": schema.BoolAttribute{
							MarkdownDescription: "Enable Automatic Restart Attempts",
							Computed:            true,
						},
						"retry_interval": schema.Int64Attribute{
							MarkdownDescription: "Attempt Interval",
							Computed:            true,
						},
						"retry_attempts": schema.Int64Attribute{
							MarkdownDescription: "Number of Attempts",
							Computed:            true,
						},
						"long_retry_interval": schema.Int64Attribute{
							MarkdownDescription: "Long Attempt Interval",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *TAMDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *TAMDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data TAMList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.TAM{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.TAM{}
	if value := res.Get(`TAM`); value.Exists() {
		for _, v := range value.Array() {
			item := models.TAM{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.TAMObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.TAMObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
