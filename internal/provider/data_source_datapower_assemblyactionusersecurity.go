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

type AssemblyActionUserSecurityList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &AssemblyActionUserSecurityDataSource{}
	_ datasource.DataSourceWithConfigure = &AssemblyActionUserSecurityDataSource{}
)

func NewAssemblyActionUserSecurityDataSource() datasource.DataSource {
	return &AssemblyActionUserSecurityDataSource{}
}

type AssemblyActionUserSecurityDataSource struct {
	client *client.DatapowerClient
}

func (d *AssemblyActionUserSecurityDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_assemblyactionusersecurity"
}

func (d *AssemblyActionUserSecurityDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "User security assembly action",
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
						"factor_id": schema.StringAttribute{
							MarkdownDescription: "Factor identifier",
							Computed:            true,
						},
						"extract_identity_method": schema.StringAttribute{
							MarkdownDescription: "Identity extraction method",
							Computed:            true,
						},
						"ei_stop_on_error": schema.BoolAttribute{
							MarkdownDescription: "Stop on error",
							Computed:            true,
						},
						"user_context_variable": schema.StringAttribute{
							MarkdownDescription: "Username context variable",
							Computed:            true,
						},
						"pass_context_variable": schema.StringAttribute{
							MarkdownDescription: "Password context variable",
							Computed:            true,
						},
						"redirect_url": schema.StringAttribute{
							MarkdownDescription: "Redirect URL",
							Computed:            true,
						},
						"redirect_time_limit": schema.Int64Attribute{
							MarkdownDescription: "Redirect time limit",
							Computed:            true,
						},
						"query_parameters": schema.StringAttribute{
							MarkdownDescription: "Query parameters",
							Computed:            true,
						},
						"ei_default_form": schema.BoolAttribute{
							MarkdownDescription: "Use default form?",
							Computed:            true,
						},
						"ei_custom_form": schema.StringAttribute{
							MarkdownDescription: "Custom form location",
							Computed:            true,
						},
						"ei_custom_form_client_profile": schema.StringAttribute{
							MarkdownDescription: "Custom form TLS client profile",
							Computed:            true,
						},
						"ei_custom_form_content_security_policy": schema.StringAttribute{
							MarkdownDescription: "Custom form CSP header value",
							Computed:            true,
						},
						"ei_form_time_limit": schema.Int64Attribute{
							MarkdownDescription: "HTML form time limit",
							Computed:            true,
						},
						"user_auth_method": schema.StringAttribute{
							MarkdownDescription: "Authentication method",
							Computed:            true,
						},
						"au_stop_on_error": schema.BoolAttribute{
							MarkdownDescription: "Stop on error",
							Computed:            true,
						},
						"user_registry": schema.StringAttribute{
							MarkdownDescription: "User registry",
							Computed:            true,
						},
						"auth_response_headers_pattern": schema.StringAttribute{
							MarkdownDescription: "Authentication response headers pattern",
							Computed:            true,
						},
						"auth_response_credential_header": schema.StringAttribute{
							MarkdownDescription: "Authenticated credential header",
							Computed:            true,
						},
						"user_az_method": schema.StringAttribute{
							MarkdownDescription: "Authorization method",
							Computed:            true,
						},
						"az_stop_on_error": schema.BoolAttribute{
							MarkdownDescription: "Stop on error",
							Computed:            true,
						},
						"az_default_form": schema.BoolAttribute{
							MarkdownDescription: "Use default form?",
							Computed:            true,
						},
						"az_custom_form": schema.StringAttribute{
							MarkdownDescription: "Custom form location",
							Computed:            true,
						},
						"az_custom_form_client_profile": schema.StringAttribute{
							MarkdownDescription: "Custom form TLS client profile",
							Computed:            true,
						},
						"az_custom_form_content_security_policy": schema.StringAttribute{
							MarkdownDescription: "Custom form CSP header value",
							Computed:            true,
						},
						"az_form_time_limit": schema.Int64Attribute{
							MarkdownDescription: "HTML form time limit",
							Computed:            true,
						},
						"az_table_display_checkboxes": schema.BoolAttribute{
							MarkdownDescription: "Display table check boxes",
							Computed:            true,
						},
						"az_table_dynamic_entries": schema.StringAttribute{
							MarkdownDescription: "Dynamic table entries",
							Computed:            true,
						},
						"az_table_default_entry": schema.ListNestedAttribute{
							MarkdownDescription: "Default table entry",
							NestedObject:        models.DmTableEntryDataSourceSchema,
							Computed:            true,
						},
						"hostname": schema.StringAttribute{
							MarkdownDescription: "Hostname",
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

func (d *AssemblyActionUserSecurityDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *AssemblyActionUserSecurityDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AssemblyActionUserSecurityList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.AssemblyActionUserSecurity{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.AssemblyActionUserSecurity{}
	if value := res.Get(`AssemblyActionUserSecurity`); value.Exists() {
		for _, v := range value.Array() {
			item := models.AssemblyActionUserSecurity{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.AssemblyActionUserSecurityObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.AssemblyActionUserSecurityObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
