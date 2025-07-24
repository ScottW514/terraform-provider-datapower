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

type AssemblyActionJWTGenerateList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &AssemblyActionJWTGenerateDataSource{}
	_ datasource.DataSourceWithConfigure = &AssemblyActionJWTGenerateDataSource{}
)

func NewAssemblyActionJWTGenerateDataSource() datasource.DataSource {
	return &AssemblyActionJWTGenerateDataSource{}
}

type AssemblyActionJWTGenerateDataSource struct {
	client *client.DatapowerClient
}

func (d *AssemblyActionJWTGenerateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_assemblyactionjwtgenerate"
}

func (d *AssemblyActionJWTGenerateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "JWT generate assembly action",
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
						"jwt": schema.StringAttribute{
							MarkdownDescription: "JWT output location",
							Computed:            true,
						},
						"jwtid_claims": schema.BoolAttribute{
							MarkdownDescription: "JWT ID claims",
							Computed:            true,
						},
						"issuer_claim": schema.StringAttribute{
							MarkdownDescription: "Issuer claim",
							Computed:            true,
						},
						"subject_claim": schema.StringAttribute{
							MarkdownDescription: "Subject claim",
							Computed:            true,
						},
						"audience_claim": schema.StringAttribute{
							MarkdownDescription: "Audience claim",
							Computed:            true,
						},
						"validity_period": schema.Int64Attribute{
							MarkdownDescription: "Validity period",
							Computed:            true,
						},
						"private_claim": schema.StringAttribute{
							MarkdownDescription: "Private claim",
							Computed:            true,
						},
						"sign_jwk": schema.StringAttribute{
							MarkdownDescription: "JWK for JWT signature",
							Computed:            true,
						},
						"crypto_algorithm": schema.StringAttribute{
							MarkdownDescription: "Crypto algorithm for JWT signature",
							Computed:            true,
						},
						"sign_crypto": schema.StringAttribute{
							MarkdownDescription: "Crypto object for JWT signature",
							Computed:            true,
						},
						"custom_kid_value_jws": schema.StringAttribute{
							MarkdownDescription: "Key ID value for JWS",
							Computed:            true,
						},
						"encrypt_algorithm": schema.StringAttribute{
							MarkdownDescription: "Encrypt algorithm for JWT encryption",
							Computed:            true,
						},
						"encrypt_jwk": schema.StringAttribute{
							MarkdownDescription: "JWK for JWT encryption",
							Computed:            true,
						},
						"key_encrypt_algorithm": schema.StringAttribute{
							MarkdownDescription: "Key encrypt algorithm for JWT encryption",
							Computed:            true,
						},
						"encrypt_crypto": schema.StringAttribute{
							MarkdownDescription: "Crypto object for JWT encryption",
							Computed:            true,
						},
						"custom_kid_value_jwe": schema.StringAttribute{
							MarkdownDescription: "Key ID value for JWE",
							Computed:            true,
						},
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"title": schema.StringAttribute{
							MarkdownDescription: "Title",
							Computed:            true,
						},
						"correlation_path": schema.StringAttribute{
							MarkdownDescription: "Correlation path",
							Computed:            true,
						},
						"action_debug": schema.BoolAttribute{
							MarkdownDescription: "Enable debugging",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *AssemblyActionJWTGenerateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *AssemblyActionJWTGenerateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AssemblyActionJWTGenerateList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.AssemblyActionJWTGenerate{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.AssemblyActionJWTGenerate{}
	if value := res.Get(`AssemblyActionJWTGenerate`); value.Exists() {
		for _, v := range value.Array() {
			item := models.AssemblyActionJWTGenerate{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.AssemblyActionJWTGenerateObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.AssemblyActionJWTGenerateObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
