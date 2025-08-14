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

type AssemblyActionJWTValidateList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &AssemblyActionJWTValidateDataSource{}
	_ datasource.DataSourceWithConfigure = &AssemblyActionJWTValidateDataSource{}
)

func NewAssemblyActionJWTValidateDataSource() datasource.DataSource {
	return &AssemblyActionJWTValidateDataSource{}
}

type AssemblyActionJWTValidateDataSource struct {
	pData *tfutils.ProviderData
}

func (d *AssemblyActionJWTValidateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_assemblyactionjwtvalidate"
}

func (d *AssemblyActionJWTValidateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "JWT validate assembly action",
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
							MarkdownDescription: "JWT location",
							Computed:            true,
						},
						"output_claims": schema.StringAttribute{
							MarkdownDescription: "Output claims",
							Computed:            true,
						},
						"issuer_claim": schema.StringAttribute{
							MarkdownDescription: "Issuer claim",
							Computed:            true,
						},
						"audience_claim": schema.StringAttribute{
							MarkdownDescription: "Audience claim",
							Computed:            true,
						},
						"decrypt_crypto": schema.StringAttribute{
							MarkdownDescription: "Crypto object for JWT decryption",
							Computed:            true,
						},
						"decrypt_jwk": schema.StringAttribute{
							MarkdownDescription: "JWK for JWT decryption",
							Computed:            true,
						},
						"verify_crypto": schema.StringAttribute{
							MarkdownDescription: "Crypto object for JWT verification",
							Computed:            true,
						},
						"verify_jwk": schema.StringAttribute{
							MarkdownDescription: "JWK for JWT verification",
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
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *AssemblyActionJWTValidateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *AssemblyActionJWTValidateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data AssemblyActionJWTValidateList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.AssemblyActionJWTValidate{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.AssemblyActionJWTValidate{}
	if value := res.Get(`AssemblyActionJWTValidate`); value.Exists() {
		for _, v := range value.Array() {
			item := models.AssemblyActionJWTValidate{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.AssemblyActionJWTValidateObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.AssemblyActionJWTValidateObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
