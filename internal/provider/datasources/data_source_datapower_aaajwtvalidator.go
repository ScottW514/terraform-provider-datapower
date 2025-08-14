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

type AAAJWTValidatorList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &AAAJWTValidatorDataSource{}
	_ datasource.DataSourceWithConfigure = &AAAJWTValidatorDataSource{}
)

func NewAAAJWTValidatorDataSource() datasource.DataSource {
	return &AAAJWTValidatorDataSource{}
}

type AAAJWTValidatorDataSource struct {
	pData *tfutils.ProviderData
}

func (d *AAAJWTValidatorDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaajwtvalidator"
}

func (d *AAAJWTValidatorDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "JWT Validator",
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
						"aud": schema.StringAttribute{
							MarkdownDescription: "Audience",
							Computed:            true,
						},
						"val_method": models.GetDmJWTValMethodDataSourceSchema("Validation method", "validate-method", ""),
						"customized_script": schema.StringAttribute{
							MarkdownDescription: "Custom validation method processing",
							Computed:            true,
						},
						"decrypt_credential_type": schema.StringAttribute{
							MarkdownDescription: "Decrypt method",
							Computed:            true,
						},
						"decrypt_key": schema.StringAttribute{
							MarkdownDescription: "Decrypt key",
							Computed:            true,
						},
						"decrypt_s_secret": schema.StringAttribute{
							MarkdownDescription: "Decrypt shared secret",
							Computed:            true,
						},
						"decrypt_jwk": schema.StringAttribute{
							MarkdownDescription: "Decrypt JWK",
							Computed:            true,
						},
						"decrypt_fetch_cred_url": schema.StringAttribute{
							MarkdownDescription: "Decrypt credential URL",
							Computed:            true,
						},
						"decrypt_fetch_cred_ssl_profile": schema.StringAttribute{
							MarkdownDescription: "Decrypt credential TLS client profile",
							Computed:            true,
						},
						"validate_custom": schema.StringAttribute{
							MarkdownDescription: "Custom decrypt/verify processing",
							Computed:            true,
						},
						"verify_credential_type": schema.StringAttribute{
							MarkdownDescription: "Verify method",
							Computed:            true,
						},
						"verify_certificate": schema.StringAttribute{
							MarkdownDescription: "Verify certificate",
							Computed:            true,
						},
						"verify_certificate_against_val_cred": schema.BoolAttribute{
							MarkdownDescription: "Signature validation credentials",
							Computed:            true,
						},
						"verify_val_cred": schema.StringAttribute{
							MarkdownDescription: "Validation credentials",
							Computed:            true,
						},
						"verify_s_secret": schema.StringAttribute{
							MarkdownDescription: "Verify shared secret",
							Computed:            true,
						},
						"verify_jwk": schema.StringAttribute{
							MarkdownDescription: "Verify JWK",
							Computed:            true,
						},
						"verify_fetch_cred_url": schema.StringAttribute{
							MarkdownDescription: "Verify credential URL",
							Computed:            true,
						},
						"verify_fetch_cred_ssl_profile": schema.StringAttribute{
							MarkdownDescription: "Verify credential TLS client profile",
							Computed:            true,
						},
						"claims": schema.ListNestedAttribute{
							MarkdownDescription: "Validate claims",
							NestedObject:        models.DmClaimDataSourceSchema,
							Computed:            true,
						},
						"username_claim": schema.StringAttribute{
							MarkdownDescription: "Claim used as username",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *AAAJWTValidatorDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *AAAJWTValidatorDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data AAAJWTValidatorList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.AAAJWTValidator{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.AAAJWTValidator{}
	if value := res.Get(`AAAJWTValidator`); value.Exists() {
		for _, v := range value.Array() {
			item := models.AAAJWTValidator{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.AAAJWTValidatorObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.AAAJWTValidatorObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
