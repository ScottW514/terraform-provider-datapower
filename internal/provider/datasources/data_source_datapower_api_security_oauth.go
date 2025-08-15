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

type APISecurityOAuthList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &APISecurityOAuthDataSource{}
	_ datasource.DataSourceWithConfigure = &APISecurityOAuthDataSource{}
)

func NewAPISecurityOAuthDataSource() datasource.DataSource {
	return &APISecurityOAuthDataSource{}
}

type APISecurityOAuthDataSource struct {
	pData *tfutils.ProviderData
}

func (d *APISecurityOAuthDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_security_oauth"
}

func (d *APISecurityOAuthDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "An API OAuth security definition defines the applicable settings for an OAuth provider.",
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
						"o_auth_provider": schema.StringAttribute{
							MarkdownDescription: "OAuth provider",
							Computed:            true,
						},
						"o_auth_flow": schema.StringAttribute{
							MarkdownDescription: "Specify the OAuth flow to enforce to protect the target. The value must be among the supported grant types that are defined by the referenced OAuth provider.",
							Computed:            true,
						},
						"o_auth_scope": schema.StringAttribute{
							MarkdownDescription: "Specify the scopes that the access token is valid to access. To specify multiple scopes, use a space between each scope. The order of scopes does not matter. <ul><li>The allowed scopes for this configuration must be a subset of the allowed scopes set for the API OAuth security definition.</li><li>The allowed scopes for the API OAuth security definition must be a subset of the allowed scopes set for the OAuth provider settings.</li></ul><p>Scopes ensure that the granted access token is valid to access only specific protected resources.</p>",
							Computed:            true,
						},
						"o_auth_adv_scope_url": schema.StringAttribute{
							MarkdownDescription: "Specify the URL that overrides the advanced scope URL in the OAuth provider settings. The value must be a properly formatted URL.",
							Computed:            true,
						},
						"o_auth_adv_scope_tls_profile": schema.StringAttribute{
							MarkdownDescription: "OAuth TLS profile",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *APISecurityOAuthDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *APISecurityOAuthDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data APISecurityOAuthList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.APISecurityOAuth{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.APISecurityOAuth{}
	if value := res.Get(`APISecurityOAuth`); value.Exists() {
		for _, v := range value.Array() {
			item := models.APISecurityOAuth{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.APISecurityOAuthObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.APISecurityOAuthObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
