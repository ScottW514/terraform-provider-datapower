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

type WebAppRequestList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &WebAppRequestDataSource{}
	_ datasource.DataSourceWithConfigure = &WebAppRequestDataSource{}
)

func NewWebAppRequestDataSource() datasource.DataSource {
	return &WebAppRequestDataSource{}
}

type WebAppRequestDataSource struct {
	client *client.DatapowerClient
}

func (d *WebAppRequestDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_webapprequest"
}

func (d *WebAppRequestDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Web Request Profile",
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
						"policy_type": schema.StringAttribute{
							MarkdownDescription: "Style",
							Computed:            true,
						},
						"ssl_policy": schema.StringAttribute{
							MarkdownDescription: "Allow SSL",
							Computed:            true,
						},
						"aaa": schema.StringAttribute{
							MarkdownDescription: "AAA Policy",
							Computed:            true,
						},
						"ss_key": schema.StringAttribute{
							MarkdownDescription: "Shared Secret Key",
							Computed:            true,
						},
						"rate_limiter": schema.StringAttribute{
							MarkdownDescription: "Rate Limiting",
							Computed:            true,
						},
						"acl": schema.StringAttribute{
							MarkdownDescription: "Access Control List",
							Computed:            true,
						},
						"ok_methods":  models.GetDmHTTPRequestMethodsDataSourceSchema("Methods", "request-methods", ""),
						"ok_versions": models.GetDmHTTPVersionMaskDataSourceSchema("Versions", "request-versions", ""),
						"min_body_size": schema.Int64Attribute{
							MarkdownDescription: "Minimum Size",
							Computed:            true,
						},
						"max_body_size": schema.Int64Attribute{
							MarkdownDescription: "Maximum Size",
							Computed:            true,
						},
						"xml_policy": schema.StringAttribute{
							MarkdownDescription: "XML Processing",
							Computed:            true,
						},
						"xml_rule": schema.StringAttribute{
							MarkdownDescription: "XML Transformation Rule",
							Computed:            true,
						},
						"non_xml_policy": schema.StringAttribute{
							MarkdownDescription: "Non-XML Processing",
							Computed:            true,
						},
						"non_xml_rule": schema.StringAttribute{
							MarkdownDescription: "Non-XML Processing Rule",
							Computed:            true,
						},
						"error_policy": schema.StringAttribute{
							MarkdownDescription: "Error Policy",
							Computed:            true,
						},
						"session_management_profile": schema.StringAttribute{
							MarkdownDescription: "Session Policy",
							Computed:            true,
						},
						"header_gnvc": schema.StringAttribute{
							MarkdownDescription: "Header Name-Value Profile",
							Computed:            true,
						},
						"url_encoded_gnvc": schema.StringAttribute{
							MarkdownDescription: "URL-Encoded Body Name-Value Profile",
							Computed:            true,
						},
						"query_string_policy": schema.StringAttribute{
							MarkdownDescription: "Allow Query String",
							Computed:            true,
						},
						"query_string_gnvc": schema.StringAttribute{
							MarkdownDescription: "QueryString Name-Value Profile",
							Computed:            true,
						},
						"sql_injection": schema.BoolAttribute{
							MarkdownDescription: "SQL Injection Filter",
							Computed:            true,
						},
						"max_uri_size": schema.Int64Attribute{
							MarkdownDescription: "Maximum URI Length",
							Computed:            true,
						},
						"uri_filter_unicode": schema.BoolAttribute{
							MarkdownDescription: "Filter Unicode",
							Computed:            true,
						},
						"uri_filter_dot_dot": schema.BoolAttribute{
							MarkdownDescription: "Filter Dot Dot",
							Computed:            true,
						},
						"uri_filter_exe": schema.BoolAttribute{
							MarkdownDescription: "Filter .exe",
							Computed:            true,
						},
						"uri_filter_fragment": schema.StringAttribute{
							MarkdownDescription: "Fragmented URI Policy",
							Computed:            true,
						},
						"content_types": schema.ListAttribute{
							MarkdownDescription: "Content-Type List",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"multipart_form_data": models.GetDmMultipartFormDataProfileDataSourceSchema("Multipart/Form-Data Profile", "multipart-form-data", ""),
						"cookie_profile":      models.GetDmCookieProfileDataSourceSchema("Cookie Profile", "cookie-policy", ""),
						"process_all_cookie": schema.BoolAttribute{
							MarkdownDescription: "Sign or Encrypt All Cookies",
							Computed:            true,
						},
						"cookie_name_vector": schema.ListAttribute{
							MarkdownDescription: "Cookie Names",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"sql_injection_patterns_file": schema.StringAttribute{
							MarkdownDescription: "SQL Injection Patterns File",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *WebAppRequestDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *WebAppRequestDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data WebAppRequestList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.WebAppRequest{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.WebAppRequest{}
	if value := res.Get(`WebAppRequest`); value.Exists() {
		for _, v := range value.Array() {
			item := models.WebAppRequest{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.WebAppRequestObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.WebAppRequestObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
