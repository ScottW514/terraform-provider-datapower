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

type APIOperationList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &APIOperationDataSource{}
	_ datasource.DataSourceWithConfigure = &APIOperationDataSource{}
)

func NewAPIOperationDataSource() datasource.DataSource {
	return &APIOperationDataSource{}
}

type APIOperationDataSource struct {
	client *client.DatapowerClient
}

func (d *APIOperationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apioperation"
}

func (d *APIOperationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "API operation",
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
						"method": schema.StringAttribute{
							MarkdownDescription: "Method",
							Computed:            true,
						},
						"operation_id": schema.StringAttribute{
							MarkdownDescription: "Operation ID",
							Computed:            true,
						},
						"remove_consume": schema.BoolAttribute{
							MarkdownDescription: "Remove consume",
							Computed:            true,
						},
						"consume": schema.ListAttribute{
							MarkdownDescription: "Consumes",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"produce": schema.ListAttribute{
							MarkdownDescription: "Produces",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"request_schema": schema.StringAttribute{
							MarkdownDescription: "Request schema",
							Computed:            true,
						},
						"response_schema": schema.ListNestedAttribute{
							MarkdownDescription: "Response schema",
							NestedObject:        models.DmAPIResponseSchemaDataSourceSchema,
							Computed:            true,
						},
						"parameter": schema.ListNestedAttribute{
							MarkdownDescription: "Parameters",
							NestedObject:        models.DmAPIParameterDataSourceSchema,
							Computed:            true,
						},
						"remove_security": schema.BoolAttribute{
							MarkdownDescription: "Remove security",
							Computed:            true,
						},
						"security": schema.ListAttribute{
							MarkdownDescription: "Security requirements",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"soap_action": schema.StringAttribute{
							MarkdownDescription: "SOAP action",
							Computed:            true,
						},
						"soap_element_name": schema.StringAttribute{
							MarkdownDescription: "SOAP element name",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *APIOperationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *APIOperationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data APIOperationList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.APIOperation{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.APIOperation{}
	if value := res.Get(`APIOperation`); value.Exists() {
		for _, v := range value.Array() {
			item := models.APIOperation{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.APIOperationObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.APIOperationObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
