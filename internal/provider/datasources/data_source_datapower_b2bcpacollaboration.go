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

type B2BCPACollaborationList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &B2BCPACollaborationDataSource{}
	_ datasource.DataSourceWithConfigure = &B2BCPACollaborationDataSource{}
)

func NewB2BCPACollaborationDataSource() datasource.DataSource {
	return &B2BCPACollaborationDataSource{}
}

type B2BCPACollaborationDataSource struct {
	client *client.DatapowerClient
}

func (d *B2BCPACollaborationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_b2bcpacollaboration"
}

func (d *B2BCPACollaborationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "B2B CPA collaboration",
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
						"internal_role": schema.StringAttribute{
							MarkdownDescription: "Internal role",
							Computed:            true,
						},
						"external_role": schema.StringAttribute{
							MarkdownDescription: "External role",
							Computed:            true,
						},
						"process_specification": schema.StringAttribute{
							MarkdownDescription: "Process Specification",
							Computed:            true,
						},
						"service": schema.StringAttribute{
							MarkdownDescription: "Service",
							Computed:            true,
						},
						"service_type": schema.StringAttribute{
							MarkdownDescription: "Service Type",
							Computed:            true,
						},
						"sender_msh_setting": schema.StringAttribute{
							MarkdownDescription: "Default MSH signal sender",
							Computed:            true,
						},
						"receiver_msh_setting": schema.StringAttribute{
							MarkdownDescription: "Default MSH signal receiver",
							Computed:            true,
						},
						"actions": schema.ListNestedAttribute{
							MarkdownDescription: "Actions",
							NestedObject:        models.DmCPACollaborationActionDataSourceSchema,
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *B2BCPACollaborationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *B2BCPACollaborationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data B2BCPACollaborationList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.B2BCPACollaboration{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.B2BCPACollaboration{}
	if value := res.Get(`B2BCPACollaboration`); value.Exists() {
		for _, v := range value.Array() {
			item := models.B2BCPACollaboration{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.B2BCPACollaborationObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.B2BCPACollaborationObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
