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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type UserGroupList struct {
	Id     types.String `tfsdk:"id"`
	Result types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &UserGroupDataSource{}
	_ datasource.DataSourceWithConfigure = &UserGroupDataSource{}
)

func NewUserGroupDataSource() datasource.DataSource {
	return &UserGroupDataSource{}
}

type UserGroupDataSource struct {
	pData *tfutils.ProviderData
}

func (d *UserGroupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user_group"
}

func (d *UserGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Create or edit user groups and their access privileges.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The name of the object to retrieve.",
				Optional:            true,
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
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"access_policies": schema.ListAttribute{
							MarkdownDescription: "Specify the access policies that define privileges for the access profile. When more than one policy applies to a resource, the most specific policy is used. To create access policies, enter a policy statement in the following format. <p><tt><i>address</i> / <i>domain</i> / <i>resource</i> ?Access= <i>privileges</i> &amp; <i>field</i> = <i>value</i></tt></p><dl><dt><i>address</i></dt><dd>The complete local IP address or host alias. Use the * character to designate all addresses.</dd><dt><i>domain</i></dt><dd>The complete domain name. Use the * character to designate all domains.</dd><dt>resource</dt><dd>The complete value for the resource type. Use the * character to designate all resource types.</dd><dt><i>privileges</i></dt><dd>The privileges to apply. Separate permissions with the + character. For example, <tt>a+d+x+r+w</tt> defines add, delete, execute, read, and write privileges.</dd><dt><i>field</i></dt><dd>The complete name of a specific property in the configuration; for example, <tt>Name</tt> .</dd><dt><i>value</i></dt><dd>The PCRE match for the property value; For example, <tt>foo(.*)bar</tt> .</dd></dl>",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"command_group": schema.ListAttribute{
							MarkdownDescription: "Specify the command groups to which the user group has CLI access. This property is superseded by the access profile when role-based management is applied to the CLI.",
							ElementType:         types.StringType,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *UserGroupDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *UserGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data UserGroupList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.UserGroup{}

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
	l := []models.UserGroup{}
	if resFound {
		if value := res.Get(`UserGroup`); value.Exists() {
			for _, v := range value.Array() {
				item := models.UserGroup{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.UserGroupObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.UserGroupObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
