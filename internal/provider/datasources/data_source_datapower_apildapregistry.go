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

type APILDAPRegistryList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &APILDAPRegistryDataSource{}
	_ datasource.DataSourceWithConfigure = &APILDAPRegistryDataSource{}
)

func NewAPILDAPRegistryDataSource() datasource.DataSource {
	return &APILDAPRegistryDataSource{}
}

type APILDAPRegistryDataSource struct {
	pData *tfutils.ProviderData
}

func (d *APILDAPRegistryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apildapregistry"
}

func (d *APILDAPRegistryDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "API LDAP registry",
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
						"ldap_host": schema.StringAttribute{
							MarkdownDescription: "Host",
							Computed:            true,
						},
						"ldap_port": schema.Int64Attribute{
							MarkdownDescription: "Port",
							Computed:            true,
						},
						"ssl_client_profile": schema.StringAttribute{
							MarkdownDescription: "TLS client profile",
							Computed:            true,
						},
						"ldap_version": schema.StringAttribute{
							MarkdownDescription: "LDAP version",
							Computed:            true,
						},
						"ldap_auth_method": schema.StringAttribute{
							MarkdownDescription: "LDAP authentication method",
							Computed:            true,
						},
						"ldap_bind_dn": schema.StringAttribute{
							MarkdownDescription: "LDAP bind DN",
							Computed:            true,
						},
						"ldap_bind_password_alias": schema.StringAttribute{
							MarkdownDescription: "LDAP bind password alias",
							Computed:            true,
						},
						"ldap_search_parameters": schema.StringAttribute{
							MarkdownDescription: "LDAP search parameters",
							Computed:            true,
						},
						"ldap_read_timeout": schema.Int64Attribute{
							MarkdownDescription: "LDAP read timeout",
							Computed:            true,
						},
						"ldap_group_auth_enabled": schema.BoolAttribute{
							MarkdownDescription: "Enable LDAP group authentication",
							Computed:            true,
						},
						"ldap_group_auth_type": schema.StringAttribute{
							MarkdownDescription: "LDAP group authentication type",
							Computed:            true,
						},
						"ldap_group_scope": schema.StringAttribute{
							MarkdownDescription: "LDAP group scope",
							Computed:            true,
						},
						"ldap_group_base_dn": schema.StringAttribute{
							MarkdownDescription: "LDAP static group base DN",
							Computed:            true,
						},
						"ldap_group_filter_prefix": schema.StringAttribute{
							MarkdownDescription: "LDAP static group filter prefix",
							Computed:            true,
						},
						"ldap_group_filter_suffix": schema.StringAttribute{
							MarkdownDescription: "LDAP static group filter suffix",
							Computed:            true,
						},
						"ldap_group_dynamic_filter": schema.StringAttribute{
							MarkdownDescription: "LDAP dynamic filter",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *APILDAPRegistryDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *APILDAPRegistryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data APILDAPRegistryList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.APILDAPRegistry{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.APILDAPRegistry{}
	if value := res.Get(`APILDAPRegistry`); value.Exists() {
		for _, v := range value.Array() {
			item := models.APILDAPRegistry{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.APILDAPRegistryObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.APILDAPRegistryObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
