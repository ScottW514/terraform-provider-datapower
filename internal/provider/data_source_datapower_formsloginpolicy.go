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

type FormsLoginPolicyList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &FormsLoginPolicyDataSource{}
	_ datasource.DataSourceWithConfigure = &FormsLoginPolicyDataSource{}
)

func NewFormsLoginPolicyDataSource() datasource.DataSource {
	return &FormsLoginPolicyDataSource{}
}

type FormsLoginPolicyDataSource struct {
	client *client.DatapowerClient
}

func (d *FormsLoginPolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_formsloginpolicy"
}

func (d *FormsLoginPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "HTML forms login policy",
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
						"login_form": schema.StringAttribute{
							MarkdownDescription: "Login",
							Computed:            true,
						},
						"use_cookie_attributes": schema.BoolAttribute{
							MarkdownDescription: "Attach cookie attribute policy",
							Computed:            true,
						},
						"cookie_attributes": schema.StringAttribute{
							MarkdownDescription: "Cookie attributes",
							Computed:            true,
						},
						"use_ssl_for_login": schema.BoolAttribute{
							MarkdownDescription: "Use TLS for Login",
							Computed:            true,
						},
						"enable_migration": schema.BoolAttribute{
							MarkdownDescription: "Enable session migration",
							Computed:            true,
						},
						"ssl_port": schema.Int64Attribute{
							MarkdownDescription: "TLS port",
							Computed:            true,
						},
						"shared_secret": schema.StringAttribute{
							MarkdownDescription: "Shared secret",
							Computed:            true,
						},
						"error_page": schema.StringAttribute{
							MarkdownDescription: "Error",
							Computed:            true,
						},
						"logout_page": schema.StringAttribute{
							MarkdownDescription: "Logout",
							Computed:            true,
						},
						"default_url": schema.StringAttribute{
							MarkdownDescription: "Default",
							Computed:            true,
						},
						"forms_location": schema.StringAttribute{
							MarkdownDescription: "Forms storage location",
							Computed:            true,
						},
						"local_login_form": schema.StringAttribute{
							MarkdownDescription: "Local login form",
							Computed:            true,
						},
						"local_error_page": schema.StringAttribute{
							MarkdownDescription: "Local error page",
							Computed:            true,
						},
						"local_logout_page": schema.StringAttribute{
							MarkdownDescription: "Local logout page",
							Computed:            true,
						},
						"username_field": schema.StringAttribute{
							MarkdownDescription: "Username field name",
							Computed:            true,
						},
						"password_field": schema.StringAttribute{
							MarkdownDescription: "Password field name",
							Computed:            true,
						},
						"redirect_field": schema.StringAttribute{
							MarkdownDescription: "Target URL field name",
							Computed:            true,
						},
						"form_processing_url": schema.StringAttribute{
							MarkdownDescription: "Form POST action URL",
							Computed:            true,
						},
						"inactivity_timeout": schema.Int64Attribute{
							MarkdownDescription: "Inactivity timeout",
							Computed:            true,
						},
						"session_lifetime": schema.Int64Attribute{
							MarkdownDescription: "Session lifetime",
							Computed:            true,
						},
						"redirect_url_type": schema.StringAttribute{
							MarkdownDescription: "Redirect URL Type",
							Computed:            true,
						},
						"form_support_type": schema.StringAttribute{
							MarkdownDescription: "Source for form processing",
							Computed:            true,
						},
						"form_support_script": schema.StringAttribute{
							MarkdownDescription: "Custom processing for form",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *FormsLoginPolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *FormsLoginPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data FormsLoginPolicyList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.FormsLoginPolicy{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.FormsLoginPolicy{}
	if value := res.Get(`FormsLoginPolicy`); value.Exists() {
		for _, v := range value.Array() {
			item := models.FormsLoginPolicy{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.FormsLoginPolicyObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.FormsLoginPolicyObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
