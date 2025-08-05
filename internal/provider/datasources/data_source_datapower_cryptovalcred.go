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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
)

type CryptoValCredList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &CryptoValCredDataSource{}
	_ datasource.DataSourceWithConfigure = &CryptoValCredDataSource{}
)

func NewCryptoValCredDataSource() datasource.DataSource {
	return &CryptoValCredDataSource{}
}

type CryptoValCredDataSource struct {
	client *client.DatapowerClient
}

func (d *CryptoValCredDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cryptovalcred"
}

func (d *CryptoValCredDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Validation credentials",
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
						"certificate": schema.ListAttribute{
							MarkdownDescription: "Certificates",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"cert_validation_mode": schema.StringAttribute{
							MarkdownDescription: "Certificate validation mode",
							Computed:            true,
						},
						"use_crl": schema.BoolAttribute{
							MarkdownDescription: "Use CRL",
							Computed:            true,
						},
						"require_crl": schema.BoolAttribute{
							MarkdownDescription: "Require CRL",
							Computed:            true,
						},
						"crldp_handling": schema.StringAttribute{
							MarkdownDescription: "CRL distribution points handling",
							Computed:            true,
						},
						"initial_policy_set": schema.ListAttribute{
							MarkdownDescription: "Initial certificate policy set",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"explicit_policy": schema.BoolAttribute{
							MarkdownDescription: "Require explicit certificate policy",
							Computed:            true,
						},
						"check_dates": schema.BoolAttribute{
							MarkdownDescription: "Check dates",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *CryptoValCredDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *CryptoValCredDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data CryptoValCredList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.CryptoValCred{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.CryptoValCred{}
	if value := res.Get(`CryptoValCred`); value.Exists() {
		for _, v := range value.Array() {
			item := models.CryptoValCred{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.CryptoValCredObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.CryptoValCredObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
