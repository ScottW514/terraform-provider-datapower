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

type APILDAPRegistryList struct {
	Id        types.String `tfsdk:"id"`
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
	resp.TypeName = req.ProviderTypeName + "_api_ldap_registry"
}

func (d *APILDAPRegistryDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Configure and manage the API LDAP registry.",
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
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"ldap_host": schema.StringAttribute{
							MarkdownDescription: "Host",
							Computed:            true,
						},
						"ldap_port": schema.Int64Attribute{
							MarkdownDescription: "Specify the listening port on the LDAP server. The default value is 636.",
							Computed:            true,
						},
						"ssl_client_profile": schema.StringAttribute{
							MarkdownDescription: "TLS client profile",
							Computed:            true,
						},
						"ldap_version": schema.StringAttribute{
							MarkdownDescription: "Specify the LDAP protocol version for bind operation. The default value is v3.",
							Computed:            true,
						},
						"ldap_auth_method": schema.StringAttribute{
							MarkdownDescription: "Specify the method to create the user for authentication. <ul><li>When compose DN, the DN can be composed from the username. For example, <tt>uid=john,ou=People,dc=company,dc=com</tt> is a DN format that can be composed from the username.</li><li>When compose UPN, the UPN can be composed from the username. For example, <tt>john@example.com</tt> is a UPN format that can be composed from the username.</li><li>When search DN, the DN cannot be composed from the username. You must use an LDAP search to retrieve information that matches the username.</li></ul><p>By default, queries the LDAP server to retrieve user information. Before deciding on the method, contact your LDAP administrator.</p>",
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
							MarkdownDescription: "Specify the time to wait for a response from the LDAP server before the connection is closed. Enter a value in the range 0 - 86400. The default value is 60. A value of 0 indicates that the connection never times out.",
							Computed:            true,
						},
						"ldap_group_auth_enabled": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enable LDAP group authentication to use to check group membership for a user. The default value is off.",
							Computed:            true,
						},
						"ldap_group_auth_type": schema.StringAttribute{
							MarkdownDescription: "Specify the type of group authentication configuration to use. The default value is static.",
							Computed:            true,
						},
						"ldap_group_scope": schema.StringAttribute{
							MarkdownDescription: "Specify the depth of the LDAP group search. The default value is subtree.",
							Computed:            true,
						},
						"ldap_group_base_dn": schema.StringAttribute{
							MarkdownDescription: "Specify the base DN name to begin the group authentication search. This value identifies the entry level of the tree used by the LDAP group scope.",
							Computed:            true,
						},
						"ldap_group_filter_prefix": schema.StringAttribute{
							MarkdownDescription: "Specify the prefix of the LDAP group filter expression. An LDAP group filter expression is composed by <tt>prefix + user DN + suffix</tt> . <p>When the prefix is <tt>(&amp;(objectclass=group)(member=</tt> and the user DN is <tt>CN=bob,DN=ibm,DN=com</tt> , the LDAP search filter is <tt>(&amp;(objectclass=group)(member=CN=bob,DN=ibm,DN=com))</tt> .</p>",
							Computed:            true,
						},
						"ldap_group_filter_suffix": schema.StringAttribute{
							MarkdownDescription: "Specify the suffix of the LDAP group filter expression. <p>When the prefix is <tt>&amp;(objectclass=group)(member=</tt> , the user DN is <tt>CN=bob,DN=ibm,DN=com</tt> , and the suffix is <tt>)(CN=ibm-group))</tt> , the LDAP search filter is <tt>(&amp;(objectclass=group)(member=CN=bob,DN=ibm,DN=com)(CN=ibm-group))</tt> .</p>",
							Computed:            true,
						},
						"ldap_group_dynamic_filter": schema.StringAttribute{
							MarkdownDescription: "Specify the filter expression of the LDAP dynamic group configuration. Only for dynamic. <p>When the filter is <tt>(memberOf=CN=ibm-group,DC=ibm,DC=com)</tt> , the value is used verbatim for LDAP group dynamic search.</p>",
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
	l := []models.APILDAPRegistry{}
	if resFound {
		if value := res.Get(`APILDAPRegistry`); value.Exists() {
			for _, v := range value.Array() {
				item := models.APILDAPRegistry{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
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
