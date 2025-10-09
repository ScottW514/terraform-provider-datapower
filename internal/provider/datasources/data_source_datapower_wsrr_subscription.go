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

type WSRRSubscriptionList struct {
	Id        types.String `tfsdk:"id"`
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
	pData *tfutils.ProviderData
}

func (d *WSRRSubscriptionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wsrr_subscription"
}

func (d *WSRRSubscriptionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The WSRR (WebSphere Service Registry and Repository) subscription is useful when you want to deploy web services with a direct reference to a service document. The configuration references a WSRR server, the name of the WSRR resource, and its namespace. If more than one version of the service document exists, you must specify the version to reference. <p>The management of service documents is controlled on the WSRR server. The service configuration is updated based on the synchronization method.</p>",
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
						"server": schema.StringAttribute{
							MarkdownDescription: "WSRR server",
							Computed:            true,
						},
						"namespace": schema.StringAttribute{
							MarkdownDescription: "Specify the namespace to unambiguously identify the WSRR resource. This property is used with the object name.",
							Computed:            true,
						},
						"object_type": schema.StringAttribute{
							MarkdownDescription: "Subscription object",
							Computed:            true,
						},
						"object_name": schema.StringAttribute{
							MarkdownDescription: "Specify the object name to unambiguously identify the WSRR resource. This property is used with the namespace.",
							Computed:            true,
						},
						"method": schema.StringAttribute{
							MarkdownDescription: "Synchronization method",
							Computed:            true,
						},
						"refresh_interval": schema.Int64Attribute{
							MarkdownDescription: "Specify the refresh interval in seconds between polls to synchronize the local copy with the registry version.",
							Computed:            true,
						},
						"use_version": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to query the registry for a specific object version. Set this property when the registry contains more than one version of an object.",
							Computed:            true,
						},
						"object_version": schema.StringAttribute{
							MarkdownDescription: "Object version",
							Computed:            true,
						},
						"fetch_policy_attachments": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to fetch external policy attachments. When enabled, the registry is queried for external policy attachments for retrieved resources. These policies are processed when the service allow external policy attachments.",
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

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *WSRRSubscriptionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data WSRRSubscriptionList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.WSRRSubscription{
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
	l := []models.WSRRSubscription{}
	if resFound {
		if value := res.Get(`WSRRSubscription`); value.Exists() {
			for _, v := range value.Array() {
				item := models.WSRRSubscription{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
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
