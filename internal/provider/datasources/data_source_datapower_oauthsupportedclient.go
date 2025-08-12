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

type OAuthSupportedClientWOList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &OAuthSupportedClientDataSource{}
	_ datasource.DataSourceWithConfigure = &OAuthSupportedClientDataSource{}
)

func NewOAuthSupportedClientDataSource() datasource.DataSource {
	return &OAuthSupportedClientDataSource{}
}

type OAuthSupportedClientDataSource struct {
	client *client.DatapowerClient
}

func (d *OAuthSupportedClientDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_oauthsupportedclient"
}

func (d *OAuthSupportedClientDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "OAuth Client Profile",
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
						"customized": schema.BoolAttribute{
							MarkdownDescription: "Customized OAuth",
							Computed:            true,
						},
						"customized_process_url": schema.StringAttribute{
							MarkdownDescription: "Customized OAuth Process",
							Computed:            true,
						},
						"o_auth_role": models.GetDmOAuthRoleDataSourceSchema("OAuth Role", "oauth-role", ""),
						"az_grant":    models.GetDmOAuthAZGrantTypeDataSourceSchema("Supported Type", "az-grant", ""),
						"client_type": schema.StringAttribute{
							MarkdownDescription: "Client Type",
							Computed:            true,
						},
						"check_client_credential": schema.BoolAttribute{
							MarkdownDescription: "Verify Client Credential",
							Computed:            true,
						},
						"use_validation_url": schema.BoolAttribute{
							MarkdownDescription: "Use Validation URL",
							Computed:            true,
						},
						"client_authen_method": schema.StringAttribute{
							MarkdownDescription: "Authentication Method",
							Computed:            true,
						},
						"client_val_cred": schema.StringAttribute{
							MarkdownDescription: "Client TLS Credentials",
							Computed:            true,
						},
						"generate_client_secret": schema.BoolAttribute{
							MarkdownDescription: "Generate Client Secret",
							Computed:            true,
						},
						"caching": schema.StringAttribute{
							MarkdownDescription: "Caching",
							Computed:            true,
						},
						"validation_url": schema.StringAttribute{
							MarkdownDescription: "Validation URL",
							Computed:            true,
						},
						"validation_features": models.GetDmValidationFeaturesDataSourceSchema("Validation Grant Features", "validation-features", ""),
						"redirect_uri": schema.ListAttribute{
							MarkdownDescription: "Redirect URI",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"custom_scope_check": schema.BoolAttribute{
							MarkdownDescription: "Customized Scope Check",
							Computed:            true,
						},
						"scope": schema.StringAttribute{
							MarkdownDescription: "Scope",
							Computed:            true,
						},
						"scope_url": schema.StringAttribute{
							MarkdownDescription: "Scope Customized Process",
							Computed:            true,
						},
						"default_scope": schema.StringAttribute{
							MarkdownDescription: "Default Scope",
							Computed:            true,
						},
						"token_secret": schema.StringAttribute{
							MarkdownDescription: "Shared Secret",
							Computed:            true,
						},
						"local_az_page_url": schema.StringAttribute{
							MarkdownDescription: "Authorization Form",
							Computed:            true,
						},
						"dp_state_life_time": schema.Int64Attribute{
							MarkdownDescription: "DataPower State Lifetime",
							Computed:            true,
						},
						"au_code_life_time": schema.Int64Attribute{
							MarkdownDescription: "Authorization Grant Lifetime",
							Computed:            true,
						},
						"access_token_life_time": schema.Int64Attribute{
							MarkdownDescription: "Access Token Lifetime",
							Computed:            true,
						},
						"refresh_token_allowed": schema.Int64Attribute{
							MarkdownDescription: "Number of Refresh Token Allowed",
							Computed:            true,
						},
						"refresh_token_life_time": schema.Int64Attribute{
							MarkdownDescription: "Refresh Token Lifetime",
							Computed:            true,
						},
						"max_consent_life_time": schema.Int64Attribute{
							MarkdownDescription: "Maximum Consent Lifetime",
							Computed:            true,
						},
						"custom_resource_owner": schema.BoolAttribute{
							MarkdownDescription: "Custom Resource Owner Handling",
							Computed:            true,
						},
						"resource_owner_url": schema.StringAttribute{
							MarkdownDescription: "Resource Owner Process",
							Computed:            true,
						},
						"additional_o_auth_process_url": schema.StringAttribute{
							MarkdownDescription: "Additional OAuth Process",
							Computed:            true,
						},
						"rs_set_header": models.GetDmOAuthRSSetHeaderDataSourceSchema("Create HTTP Headers for", "rs-set-header", ""),
						"validation_urlssl_client_type": schema.StringAttribute{
							MarkdownDescription: "Validation TLS client type",
							Computed:            true,
						},
						"validation_urlssl_client": schema.StringAttribute{
							MarkdownDescription: "Validation TLS Client Profile",
							Computed:            true,
						},
						"jwt_grant_validator": schema.StringAttribute{
							MarkdownDescription: "Authorization grant JWT validator",
							Computed:            true,
						},
						"client_jwt_validator": schema.StringAttribute{
							MarkdownDescription: "Client authentication JWT validator",
							Computed:            true,
						},
						"oidcid_token_generator": schema.StringAttribute{
							MarkdownDescription: "ID token JWT generator",
							Computed:            true,
						},
						"o_auth_features":    models.GetDmOAuthFeaturesDataSourceSchema("Features", "oauth-features", ""),
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *OAuthSupportedClientDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *OAuthSupportedClientDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data OAuthSupportedClientWOList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.OAuthSupportedClientWO{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.OAuthSupportedClientWO{}
	if value := res.Get(`OAuthSupportedClient`); value.Exists() {
		for _, v := range value.Array() {
			item := models.OAuthSupportedClientWO{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.OAuthSupportedClientObjectTypeWO}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.OAuthSupportedClientObjectTypeWO})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
