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

type NFSStaticMountList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &NFSStaticMountDataSource{}
	_ datasource.DataSourceWithConfigure = &NFSStaticMountDataSource{}
)

func NewNFSStaticMountDataSource() datasource.DataSource {
	return &NFSStaticMountDataSource{}
}

type NFSStaticMountDataSource struct {
	client *client.DatapowerClient
}

func (d *NFSStaticMountDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nfsstaticmount"
}

func (d *NFSStaticMountDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "NFS static mounts",
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
						"remote": schema.StringAttribute{
							MarkdownDescription: "Remote NFS export",
							Computed:            true,
						},
						"local_filesystem_access": schema.BoolAttribute{
							MarkdownDescription: "Local file system access",
							Computed:            true,
						},
						"version": schema.Int64Attribute{
							MarkdownDescription: "NFS version",
							Computed:            true,
						},
						"transport": schema.StringAttribute{
							MarkdownDescription: "Transport protocol",
							Computed:            true,
						},
						"mount_type": schema.StringAttribute{
							MarkdownDescription: "Mount type",
							Computed:            true,
						},
						"read_only": schema.BoolAttribute{
							MarkdownDescription: "Read-Only",
							Computed:            true,
						},
						"read_size": schema.Int64Attribute{
							MarkdownDescription: "Read size",
							Computed:            true,
						},
						"write_size": schema.Int64Attribute{
							MarkdownDescription: "Write size",
							Computed:            true,
						},
						"timeout": schema.Int64Attribute{
							MarkdownDescription: "Retransmission timeout",
							Computed:            true,
						},
						"retransmissions": schema.Int64Attribute{
							MarkdownDescription: "Max retransmissions",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *NFSStaticMountDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *NFSStaticMountDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data NFSStaticMountList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.NFSStaticMount{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.NFSStaticMount{}
	if value := res.Get(`NFSStaticMount`); value.Exists() {
		for _, v := range value.Array() {
			item := models.NFSStaticMount{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.NFSStaticMountObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.NFSStaticMountObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
