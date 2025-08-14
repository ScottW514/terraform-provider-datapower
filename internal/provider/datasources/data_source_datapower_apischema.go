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

type APISchemaList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &APISchemaDataSource{}
	_ datasource.DataSourceWithConfigure = &APISchemaDataSource{}
)

func NewAPISchemaDataSource() datasource.DataSource {
	return &APISchemaDataSource{}
}

type APISchemaDataSource struct {
	pData *tfutils.ProviderData
}

func (d *APISchemaDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apischema"
}

func (d *APISchemaDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "An API schema defines the schemas to validate JSON, XML, WSDL, or SOAP messages.",
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
						"json_schema": schema.StringAttribute{
							MarkdownDescription: "Specify the schema URL for JSON message validation. For example, <tt>local:///petstore.json#/definitions/Pet</tt> . To accept all input, use the string <tt>accept</tt> instead of a URL. To reject all input, use the string <tt>reject</tt> instead of a URL.",
							Computed:            true,
						},
						"graph_ql_schema": schema.StringAttribute{
							MarkdownDescription: "Specify the schema URL for GraphQL message validation. For example, <tt>local:///petstore-Pet.graphql</tt> . To accept all input, use the string <tt>accept</tt> instead of a URL. To reject all input, use the string <tt>reject</tt> instead of a URL.",
							Computed:            true,
						},
						"xml_type": schema.StringAttribute{
							MarkdownDescription: "XML type",
							Computed:            true,
						},
						"xml_validation_mode": schema.StringAttribute{
							MarkdownDescription: "XML validation mode",
							Computed:            true,
						},
						"xml_schema_url": schema.StringAttribute{
							MarkdownDescription: "Specify the schema URL for XML message validation. For example, <tt>local:///petstore-Pet.xsd</tt> . To accept all input, use the string <tt>accept</tt> instead of a URL. To reject all input, use the string <tt>reject</tt> instead of a URL.",
							Computed:            true,
						},
						"wsdl_schema_url": schema.StringAttribute{
							MarkdownDescription: "Specify the schema URL for WSDL message validation. For example, <tt>local:///petstore-Pet.wsdl</tt> . To accept all input, use the string <tt>accept</tt> instead of a URL. To reject all input, use the string <tt>reject</tt> instead of a URL.",
							Computed:            true,
						},
						"wsdl_port_q_name": schema.StringAttribute{
							MarkdownDescription: "Specify the <tt>wsdl:port</tt> for the traffic to validate. The value should be a QName in the form <tt>{namespace-uri}local-part</tt> or <tt>*</tt> for all ports in the WSDL file. When specified and not <tt>*</tt> , only messages for the named port are valid.",
							Computed:            true,
						},
						"wsdl_operation_name": schema.StringAttribute{
							MarkdownDescription: "Specify the <tt>wsdl:operation</tt> for the traffic to validate. The value should be the unqualified name of the operation or <tt>*</tt> for all operations in the WSDL file. When specified and not <tt>*</tt> , only messages for the named operation are valid.",
							Computed:            true,
						},
						"wsdl_message_direction_or_name": schema.StringAttribute{
							MarkdownDescription: "Specify the <tt>wsdl:input</tt> , <tt>wsdl:output</tt> , or <tt>wsdl:fault</tt> for the traffic to validate. The value must be the name of one or more WSDL input, output, or fault components, or <tt>#input</tt> or <tt>#output</tt> for the request and response directions respectively, or <tt>*</tt> for all inputs, outputs, and faults in the WSDL file. When specified and not <tt>*</tt> , only messages that match the specified direction or name are valid. Faults are valid for the response direction.",
							Computed:            true,
						},
						"wsdl_attachment_part": schema.StringAttribute{
							MarkdownDescription: "Specify the mime:content to validate in the format <tt>mime:content/@part</tt> . The value must be the unqualified name of the message part. The name is the same as the part attribute on the corresponding <tt>mime:content</tt> component in the WSDL file. When not specified or <tt>*</tt> , the root MIME part is validated. The root MIME part is bound to a <tt>soap:Body</tt> .",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *APISchemaDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *APISchemaDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data APISchemaList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.APISchema{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.APISchema{}
	if value := res.Get(`APISchema`); value.Exists() {
		for _, v := range value.Array() {
			item := models.APISchema{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.APISchemaObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.APISchemaObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
