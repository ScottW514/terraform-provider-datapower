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

type AccessProfileList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &AccessProfileDataSource{}
	_ datasource.DataSourceWithConfigure = &AccessProfileDataSource{}
)

func NewAccessProfileDataSource() datasource.DataSource {
	return &AccessProfileDataSource{}
}

type AccessProfileDataSource struct {
	pData *tfutils.ProviderData
}

func (d *AccessProfileDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_accessprofile"
}

func (d *AccessProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Create or edit an access profile.",
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
						"access_policies": schema.ListAttribute{
							MarkdownDescription: "Specify the access policies that define privileges for the access profile. When more than one policy applies to a resource, the most specific policy is used. To create access policies, enter a policy statement in the following format. <p><tt><i>address</i> / <i>domain</i> / <i>resource</i> ?Access= <i>privileges</i> &amp; <i>field</i> = <i>value</i></tt></p><dl><dt><i>address</i></dt><dd>The complete local IP address or host alias. Use the * character to designate all addresses.</dd><dt><i>domain</i></dt><dd>The complete domain name. Use the * character to designate all domains.</dd><dt>resource</dt><dd>The complete value for the resource type. Use the * character to designate all resource types.</dd><dt><i>privileges</i></dt><dd>The privileges to apply. Separate permissions with the + character. For example, <tt>a+d+x+r+w</tt> defines add, delete, execute, read, and write privileges.</dd><dt><i>field</i></dt><dd>The complete name of a specific property in the configuration; for example, <tt>Name</tt> .</dd><dt><i>value</i></dt><dd>The PCRE match for the property value; For example, <tt>foo(.*)bar</tt> .</dd></dl>",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *AccessProfileDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *AccessProfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data AccessProfileList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.AccessProfile{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.AccessProfile{}
	if value := res.Get(`AccessProfile`); value.Exists() {
		for _, v := range value.Array() {
			item := models.AccessProfile{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.AccessProfileObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.AccessProfileObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
