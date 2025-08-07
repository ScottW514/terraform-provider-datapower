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

type AAAJWTGeneratorList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &AAAJWTGeneratorDataSource{}
	_ datasource.DataSourceWithConfigure = &AAAJWTGeneratorDataSource{}
)

func NewAAAJWTGeneratorDataSource() datasource.DataSource {
	return &AAAJWTGeneratorDataSource{}
}

type AAAJWTGeneratorDataSource struct {
	client *client.DatapowerClient
}

func (d *AAAJWTGeneratorDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaajwtgenerator"
}

func (d *AAAJWTGeneratorDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "JWT Generator",
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
						"issuer": schema.StringAttribute{
							MarkdownDescription: "Issuer",
							Computed:            true,
						},
						"duration": schema.Int64Attribute{
							MarkdownDescription: "Validity period",
							Computed:            true,
						},
						"additional_claims": models.GetDmJWTClaimsDataSourceSchema("Additional claims", "add-claims", ""),
						"audience": schema.ListAttribute{
							MarkdownDescription: "Audience claim",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"not_before": schema.Int64Attribute{
							MarkdownDescription: "Delta for not before claim",
							Computed:            true,
						},
						"custom_claims": schema.StringAttribute{
							MarkdownDescription: "Custom claims",
							Computed:            true,
						},
						"gen_method": models.GetDmJWTGenMethodDataSourceSchema("JWT generation method", "generate-method", ""),
						"sign_algorithm": schema.StringAttribute{
							MarkdownDescription: "Signing algorithm",
							Computed:            true,
						},
						"sign_key": schema.StringAttribute{
							MarkdownDescription: "Signing key",
							Computed:            true,
						},
						"sign_ss_key": schema.StringAttribute{
							MarkdownDescription: "Signing shared secret",
							Computed:            true,
						},
						"enc_algorithm": schema.StringAttribute{
							MarkdownDescription: "Encryption algorithm",
							Computed:            true,
						},
						"encrypt_algorithm": schema.StringAttribute{
							MarkdownDescription: "Key management algorithm",
							Computed:            true,
						},
						"encrypt_certificate": schema.StringAttribute{
							MarkdownDescription: "Encryption certificate",
							Computed:            true,
						},
						"encrypt_ss_key": schema.StringAttribute{
							MarkdownDescription: "Encryption key",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *AAAJWTGeneratorDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *AAAJWTGeneratorDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AAAJWTGeneratorList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.AAAJWTGenerator{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.AAAJWTGenerator{}
	if value := res.Get(`AAAJWTGenerator`); value.Exists() {
		for _, v := range value.Array() {
			item := models.AAAJWTGenerator{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.AAAJWTGeneratorObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.AAAJWTGeneratorObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
