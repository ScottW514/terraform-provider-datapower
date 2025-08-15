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

type WSRRSavedSearchSubscriptionList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &WSRRSavedSearchSubscriptionDataSource{}
	_ datasource.DataSourceWithConfigure = &WSRRSavedSearchSubscriptionDataSource{}
)

func NewWSRRSavedSearchSubscriptionDataSource() datasource.DataSource {
	return &WSRRSavedSearchSubscriptionDataSource{}
}

type WSRRSavedSearchSubscriptionDataSource struct {
	pData *tfutils.ProviderData
}

func (d *WSRRSavedSearchSubscriptionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wsrr_saved_search_subscription"
}

func (d *WSRRSavedSearchSubscriptionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The WSRR (WebSphere Service Registry and Repository) saved subscription is useful when you want to deploy services with an indirect reference to WSRR resources. The configuration references a WSRR server, a saved search and a saved search or named query on the server. <p>The management of service documents is controlled on the WSRR server. The service configuration is updated based on the synchronization method.</p>",
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
						"saved_search_name": schema.StringAttribute{
							MarkdownDescription: "Saved search or named query",
							Computed:            true,
						},
						"saved_search_parameters": schema.ListAttribute{
							MarkdownDescription: "Specify the parameters to include in the query. The query to the registry uses these parameters. A parameter can be up to 255 characters in length. You can define a maximum of 9 parameters. <p>If you define parameters and they are not require parameters, an error is logged.</p><p><b>Note:</b> In WSRR, a named query and a saved search can have the same name. WSRR matches named queries before saved searches. Therefore, WSRR never finds a saved search with the same name as a named query.</p>",
							ElementType:         types.StringType,
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
						"fetch_policy_attachments": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to fetch external policy attachments. When enabled, the registry is queried for external policy attachments for retrieved resources. These policies are processed when the service allow external policy attachments.",
							Computed:            true,
						},
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *WSRRSavedSearchSubscriptionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *WSRRSavedSearchSubscriptionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data WSRRSavedSearchSubscriptionList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.WSRRSavedSearchSubscription{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.WSRRSavedSearchSubscription{}
	if value := res.Get(`WSRRSavedSearchSubscription`); value.Exists() {
		for _, v := range value.Array() {
			item := models.WSRRSavedSearchSubscription{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.WSRRSavedSearchSubscriptionObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.WSRRSavedSearchSubscriptionObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
