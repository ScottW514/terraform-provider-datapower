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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
)

type APIDefinitionList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &APIDefinitionDataSource{}
	_ datasource.DataSourceWithConfigure = &APIDefinitionDataSource{}
)

func NewAPIDefinitionDataSource() datasource.DataSource {
	return &APIDefinitionDataSource{}
}

type APIDefinitionDataSource struct {
	client *client.DatapowerClient
}

func (d *APIDefinitionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apidefinition"
}

func (d *APIDefinitionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "API definition",
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
						"api_id": schema.StringAttribute{
							MarkdownDescription: "API ID",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "API name",
							Computed:            true,
						},
						"version": schema.StringAttribute{
							MarkdownDescription: "API version",
							Computed:            true,
						},
						"base_path": schema.StringAttribute{
							MarkdownDescription: "Base path",
							Computed:            true,
						},
						"html_page": schema.StringAttribute{
							MarkdownDescription: "HTML page",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Type",
							Computed:            true,
						},
						"assembly": schema.StringAttribute{
							MarkdownDescription: "Assembly",
							Computed:            true,
						},
						"path": schema.ListAttribute{
							MarkdownDescription: "Paths",
							ElementType:         types.StringType,
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
						"swagger_location": schema.StringAttribute{
							MarkdownDescription: "OpenAPI document",
							Computed:            true,
						},
						"graph_ql_schema": schema.StringAttribute{
							MarkdownDescription: "GraphQL schema location",
							Computed:            true,
						},
						"wsdl_advertised_schema_location": schema.StringAttribute{
							MarkdownDescription: "WSDL advertised schema location",
							Computed:            true,
						},
						"wsdl_validation_schema": schema.StringAttribute{
							MarkdownDescription: "WSDL validation schema location",
							Computed:            true,
						},
						"security_requirement": schema.ListAttribute{
							MarkdownDescription: "Security requirements",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"require_api_mutual_tls": schema.BoolAttribute{
							MarkdownDescription: "API protection",
							Computed:            true,
						},
						"api_mutual_tls_source": schema.ListAttribute{
							MarkdownDescription: "API protection sources",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"api_mutual_tls_header_name": schema.StringAttribute{
							MarkdownDescription: "API protection HTTP header",
							Computed:            true,
						},
						"properties": schema.ListNestedAttribute{
							MarkdownDescription: "Custom properties",
							NestedObject:        models.DmAPIPropertyDataSourceSchema,
							Computed:            true,
						},
						"schemas": schema.ListNestedAttribute{
							MarkdownDescription: "Schemas",
							NestedObject:        models.DmAPIDataTypeDefinitionDataSourceSchema,
							Computed:            true,
						},
						"cors_toggle": schema.BoolAttribute{
							MarkdownDescription: "Enable CORS",
							Computed:            true,
						},
						"cors_policy": schema.StringAttribute{
							MarkdownDescription: "CORS policy",
							Computed:            true,
						},
						"activity_log_toggle": schema.BoolAttribute{
							MarkdownDescription: "Enable logging",
							Computed:            true,
						},
						"content": schema.StringAttribute{
							MarkdownDescription: "Content on success",
							Computed:            true,
						},
						"error_content": schema.StringAttribute{
							MarkdownDescription: "Content on error",
							Computed:            true,
						},
						"preserved_request_header": schema.ListAttribute{
							MarkdownDescription: "Request headers to preserve",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"preserved_response_header": schema.ListAttribute{
							MarkdownDescription: "Response header to preserve",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"message_buffering": schema.BoolAttribute{
							MarkdownDescription: "Enable message buffering",
							Computed:            true,
						},
						"deployment_state": schema.StringAttribute{
							MarkdownDescription: "Deployment state",
							Computed:            true,
						},
						"share_rate_limit_count": schema.StringAttribute{
							MarkdownDescription: "Share rate limit count",
							Computed:            true,
						},
						"return_v5_responses": schema.BoolAttribute{
							MarkdownDescription: "Return v5 responses",
							Computed:            true,
						},
						"copy_id_headers_to_message": schema.BoolAttribute{
							MarkdownDescription: "Copy ID headers to message",
							Computed:            true,
						},
						"enforce_required_params": schema.BoolAttribute{
							MarkdownDescription: "Enforce required parameters",
							Computed:            true,
						},
						"allow_chunked_uploads": schema.BoolAttribute{
							MarkdownDescription: "Allow chunked uploads",
							Computed:            true,
						},
						"set_v5_request_headers": schema.BoolAttribute{
							MarkdownDescription: "Set v5 request headers",
							Computed:            true,
						},
						"get_raw_body_value": schema.BoolAttribute{
							MarkdownDescription: "Get raw body value",
							Computed:            true,
						},
						"allowed_api_protocols": models.GetDmAPIProtocolsDataSourceSchema("Allowed API protocols", "allowed-api-protocols", ""),
						"allow_trailing_slash": schema.BoolAttribute{
							MarkdownDescription: "Allow trailing slash",
							Computed:            true,
						},
						"enforce_all_headers_case_insensitive": schema.BoolAttribute{
							MarkdownDescription: "Enforce all headers as case-insensitive",
							Computed:            true,
						},
						"enforce_form_data_parameter": schema.BoolAttribute{
							MarkdownDescription: "Enforce form data parameter",
							Computed:            true,
						},
						"force_http500_for_soap11": schema.BoolAttribute{
							MarkdownDescription: "Force HTTP 500 for SOAP 1.1",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *APIDefinitionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *APIDefinitionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data APIDefinitionList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.APIDefinition{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.APIDefinition{}
	if value := res.Get(`APIDefinition`); value.Exists() {
		for _, v := range value.Array() {
			item := models.APIDefinition{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.APIDefinitionObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.APIDefinitionObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
