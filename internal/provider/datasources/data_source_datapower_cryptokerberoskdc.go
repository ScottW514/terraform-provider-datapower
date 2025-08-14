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

type CryptoKerberosKDCList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &CryptoKerberosKDCDataSource{}
	_ datasource.DataSourceWithConfigure = &CryptoKerberosKDCDataSource{}
)

func NewCryptoKerberosKDCDataSource() datasource.DataSource {
	return &CryptoKerberosKDCDataSource{}
}

type CryptoKerberosKDCDataSource struct {
	pData *tfutils.ProviderData
}

func (d *CryptoKerberosKDCDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cryptokerberoskdc"
}

func (d *CryptoKerberosKDCDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Kerberos KDC Server",
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
						"realm": schema.StringAttribute{
							MarkdownDescription: "Kerberos realm name",
							Computed:            true,
						},
						"server": schema.StringAttribute{
							MarkdownDescription: "Kerberos KDC Server",
							Computed:            true,
						},
						"use_tcp": schema.BoolAttribute{
							MarkdownDescription: "Use TCP",
							Computed:            true,
						},
						"server_port": schema.Int64Attribute{
							MarkdownDescription: "Server Port Number",
							Computed:            true,
						},
						"udp_timeout": schema.Int64Attribute{
							MarkdownDescription: "UDP Timeout",
							Computed:            true,
						},
						"cache_tickets": schema.BoolAttribute{
							MarkdownDescription: "Cache Service Tickets",
							Computed:            true,
						},
						"max_cached_tickets": schema.Int64Attribute{
							MarkdownDescription: "Maximum Number of Cached Service Tickets",
							Computed:            true,
						},
						"min_cached_ticket_validity": schema.Int64Attribute{
							MarkdownDescription: "Minimum Cached Ticket Validity",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *CryptoKerberosKDCDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *CryptoKerberosKDCDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data CryptoKerberosKDCList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.CryptoKerberosKDC{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.CryptoKerberosKDC{}
	if value := res.Get(`CryptoKerberosKDC`); value.Exists() {
		for _, v := range value.Array() {
			item := models.CryptoKerberosKDC{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.CryptoKerberosKDCObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.CryptoKerberosKDCObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
