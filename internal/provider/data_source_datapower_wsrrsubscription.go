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

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
)

type WSRRSubscriptionList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &WSRRSubscriptionDataSource{}
	_ datasource.DataSourceWithConfigure = &WSRRSubscriptionDataSource{}
)

func NewWSRRSubscriptionDataSource() datasource.DataSource {
	return &WSRRSubscriptionDataSource{}
}

type WSRRSubscriptionDataSource struct {
	client *client.DatapowerClient
}

func (d *WSRRSubscriptionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wsrrsubscription"
}

func (d *WSRRSubscriptionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "WSRR subscription",
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
						"server": schema.StringAttribute{
							MarkdownDescription: "WSRR server",
							Computed:            true,
						},
						"namespace": schema.StringAttribute{
							MarkdownDescription: "Namespace",
							Computed:            true,
						},
						"object_type": schema.StringAttribute{
							MarkdownDescription: "Subscription object",
							Computed:            true,
						},
						"object_name": schema.StringAttribute{
							MarkdownDescription: "Object name",
							Computed:            true,
						},
						"method": schema.StringAttribute{
							MarkdownDescription: "Synchronization method",
							Computed:            true,
						},
						"refresh_interval": schema.Int64Attribute{
							MarkdownDescription: "Refresh interval",
							Computed:            true,
						},
						"use_version": schema.BoolAttribute{
							MarkdownDescription: "Use object version",
							Computed:            true,
						},
						"object_version": schema.StringAttribute{
							MarkdownDescription: "Object version",
							Computed:            true,
						},
						"fetch_policy_attachments": schema.BoolAttribute{
							MarkdownDescription: "Fetch policy attachments",
							Computed:            true,
						},
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *WSRRSubscriptionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *WSRRSubscriptionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data WSRRSubscriptionList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.WSRRSubscription{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.WSRRSubscription{}
	if value := res.Get(`WSRRSubscription`); value.Exists() {
		for _, v := range value.Array() {
			item := models.WSRRSubscription{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.WSRRSubscriptionObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.WSRRSubscriptionObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
