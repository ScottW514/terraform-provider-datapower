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

type LDAPSearchParametersList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &LDAPSearchParametersDataSource{}
	_ datasource.DataSourceWithConfigure = &LDAPSearchParametersDataSource{}
)

func NewLDAPSearchParametersDataSource() datasource.DataSource {
	return &LDAPSearchParametersDataSource{}
}

type LDAPSearchParametersDataSource struct {
	pData *tfutils.ProviderData
}

func (d *LDAPSearchParametersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ldap_search_parameters"
}

func (d *LDAPSearchParametersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "LDAP search parameters are a container for the parameters to use to perform an LDAP search operation. When used with authentication, the search retrieves the distinguished name (DN) for the user. When used with credential authorization mapping, the search retrieves the group name (DN or attribute value) based on the DN of the authenticated user.",
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
						"ldap_base_dn": schema.StringAttribute{
							MarkdownDescription: "Specify the base DN to begin the search. This value identifies the entry level of the tree.",
							Computed:            true,
						},
						"ldap_returned_attribute": schema.StringAttribute{
							MarkdownDescription: "Specify the LDAP attribute to return for each entry that matches the search filter. The default value is dn.",
							Computed:            true,
						},
						"ldap_filter_prefix": schema.StringAttribute{
							MarkdownDescription: "Specify the prefix of the LDAP filter expression. An LDAP filter expression is composed by <tt>prefix + username + suffix</tt> . If the prefix is <tt>(&amp;(uid=</tt> and the username is <tt>bob</tt> , the LDAP search filter is <tt>(&amp;(uid=bob</tt> .",
							Computed:            true,
						},
						"ldap_filter_suffix": schema.StringAttribute{
							MarkdownDescription: "Specify the suffix of the LDAP filter expression. An LDAP filter expression is composed by <tt>prefix + username + suffix</tt> . If the prefix is <tt>(&amp;(uid=</tt> , the username is <tt>bob</tt> , and the suffix is <tt>)(objectClass=person))</tt> , the LDAP search filter is <tt>(&amp;(uid=bob)(objectClass=person))</tt> .",
							Computed:            true,
						},
						"ldap_scope": schema.StringAttribute{
							MarkdownDescription: "Specify the depth of the LDAP search. The default is subtree.",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *LDAPSearchParametersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *LDAPSearchParametersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data LDAPSearchParametersList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.LDAPSearchParameters{
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
	l := []models.LDAPSearchParameters{}
	if resFound {
		if value := res.Get(`LDAPSearchParameters`); value.Exists() {
			for _, v := range value.Array() {
				item := models.LDAPSearchParameters{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.LDAPSearchParametersObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.LDAPSearchParametersObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
