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

type SocialLoginPolicyList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &SocialLoginPolicyDataSource{}
	_ datasource.DataSourceWithConfigure = &SocialLoginPolicyDataSource{}
)

func NewSocialLoginPolicyDataSource() datasource.DataSource {
	return &SocialLoginPolicyDataSource{}
}

type SocialLoginPolicyDataSource struct {
	client *client.DatapowerClient
}

func (d *SocialLoginPolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_socialloginpolicy"
}

func (d *SocialLoginPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Social Login Policy",
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
						"client_id": schema.StringAttribute{
							MarkdownDescription: "Client ID",
							Computed:            true,
						},
						"client_secret": schema.StringAttribute{
							MarkdownDescription: "Client secret",
							Computed:            true,
						},
						"client_grant": schema.StringAttribute{
							MarkdownDescription: "Client grant type",
							Computed:            true,
						},
						"client_scope": schema.StringAttribute{
							MarkdownDescription: "Scope",
							Computed:            true,
						},
						"client_redirect_uri": schema.StringAttribute{
							MarkdownDescription: "Client redirection URI",
							Computed:            true,
						},
						"client_optional_query_params": schema.StringAttribute{
							MarkdownDescription: "Client Optional Query Parameters",
							Computed:            true,
						},
						"client_ssl_profile": schema.StringAttribute{
							MarkdownDescription: "TLS client profile",
							Computed:            true,
						},
						"social_provider": schema.StringAttribute{
							MarkdownDescription: "Social login provider",
							Computed:            true,
						},
						"provider_az_endpoint": schema.StringAttribute{
							MarkdownDescription: "Authorization endpoint URL",
							Computed:            true,
						},
						"provider_token_endpoint": schema.StringAttribute{
							MarkdownDescription: "Token endpoint URL",
							Computed:            true,
						},
						"validate_jwt_token": schema.BoolAttribute{
							MarkdownDescription: "Enable JWT token validation",
							Computed:            true,
						},
						"jwt_validator": schema.StringAttribute{
							MarkdownDescription: "JWT Validator",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *SocialLoginPolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *SocialLoginPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SocialLoginPolicyList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.SocialLoginPolicy{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.SocialLoginPolicy{}
	if value := res.Get(`SocialLoginPolicy`); value.Exists() {
		for _, v := range value.Array() {
			item := models.SocialLoginPolicy{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.SocialLoginPolicyObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.SocialLoginPolicyObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
