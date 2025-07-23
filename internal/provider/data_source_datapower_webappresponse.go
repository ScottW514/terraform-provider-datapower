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

type WebAppResponseList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &WebAppResponseDataSource{}
	_ datasource.DataSourceWithConfigure = &WebAppResponseDataSource{}
)

func NewWebAppResponseDataSource() datasource.DataSource {
	return &WebAppResponseDataSource{}
}

type WebAppResponseDataSource struct {
	client *client.DatapowerClient
}

func (d *WebAppResponseDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_webappresponse"
}

func (d *WebAppResponseDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Web Response Profile",
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
						"ok_codes":    models.GetDmHTTPResponseCodesDataSourceSchema("Response Codes", "response-codes", ""),
						"ok_versions": models.GetDmHTTPVersionMaskDataSourceSchema("Response Versions", "response-versions", ""),
						"min_body_size": schema.Int64Attribute{
							MarkdownDescription: "Minimum Size",
							Computed:            true,
						},
						"max_body_size": schema.Int64Attribute{
							MarkdownDescription: "Maximum Size",
							Computed:            true,
						},
						"header_gnvc": schema.StringAttribute{
							MarkdownDescription: "Header Name-Value Profile",
							Computed:            true,
						},
						"content_types": schema.ListAttribute{
							MarkdownDescription: "Content-Type List",
							ElementType:         types.StringType,
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
					},
				},
			},
		},
	}
}

func (d *WebAppResponseDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *WebAppResponseDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data WebAppResponseList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.WebAppResponse{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.WebAppResponse{}
	if value := res.Get(`WebAppResponse`); value.Exists() {
		for _, v := range value.Array() {
			item := models.WebAppResponse{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.WebAppResponseObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.WebAppResponseObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
