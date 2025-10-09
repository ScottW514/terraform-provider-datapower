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
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type APIOperationList struct {
	Id        types.String `tfsdk:"id"`
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
	pData *tfutils.ProviderData
}

func (d *APIOperationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_operation"
}

func (d *APIOperationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "An API operation describes the actions to perform against the resource.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The name of the object to retrieve.",
				Optional:            true,
			},
			"app_domain": schema.StringAttribute{
				MarkdownDescription: "The name of the application domain the object belongs to.",
				Required:            true,
			},
			"result": schema.ListNestedAttribute{
				MarkdownDescription: "List of objects. If `id` was provided and it exists, it will be the only item in the list.",
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
							MarkdownDescription: "Specify whether to remove the API-level consume declaration. By default, the API-level consume declaration is applied to the operation. When removed, the operation can always be performed regardless of the content type.",
							Computed:            true,
						},
						"consume": schema.ListAttribute{
							MarkdownDescription: "Specify MIME types that the operation can consume. This setting overrides the API-level consume declaration that is defined in the API definition.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"produce": schema.ListAttribute{
							MarkdownDescription: "Specify MIME types that the operation can produce. This setting overrides the API-level produce declaration that is defined in the API definition.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"request_schema": schema.StringAttribute{
							MarkdownDescription: "Request schema",
							Computed:            true,
						},
						"response_schema": schema.ListNestedAttribute{
							MarkdownDescription: "Response schema",
							NestedObject:        models.GetDmAPIResponseSchemaDataSourceSchema(),
							Computed:            true,
						},
						"parameter": schema.ListNestedAttribute{
							MarkdownDescription: "Specify applicable parameters for the API operation. This setting overrides the setting in the API path configuration for the same parameter.",
							NestedObject:        models.GetDmAPIParameterDataSourceSchema(),
							Computed:            true,
						},
						"remove_security": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to remove the API-level security declaration that is defined for the API. By default, the API-level security declaration is applied to the operation. When removed, the operation can be performed without security check.",
							Computed:            true,
						},
						"security": schema.ListAttribute{
							MarkdownDescription: "Specify the alternative security requirements to enforce for the operation (that is, there is a logical OR between the security requirements). This setting overrides any declared API-level security.",
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

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *APIOperationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data APIOperationList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.APIOperation{
		AppDomain: data.AppDomain,
	}

	path := o.GetPath()
	if !data.Id.IsNull() {
		path = path + "/" + data.Id.ValueString()
	}

	res, err := d.pData.Client.Get(path)
	resFound := true
	if err != nil {
		if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		} else {
			resFound = false
		}
	}
	l := []models.APIOperation{}
	if resFound {
		if value := res.Get(`APIOperation`); value.Exists() {
			for _, v := range value.Array() {
				item := models.APIOperation{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
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
