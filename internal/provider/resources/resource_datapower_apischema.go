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

package resources

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &APISchemaResource{}

func NewAPISchemaResource() resource.Resource {
	return &APISchemaResource{}
}

type APISchemaResource struct {
	pData *tfutils.ProviderData
}

func (r *APISchemaResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apischema"
}

func (r *APISchemaResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("An API schema defines the schemas to validate JSON, XML, WSDL, or SOAP messages.", "api-schema", "").String,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name of the object. Must be unique among object types in application domain.", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"app_domain": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The name of the application domain the object belongs to", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`), ""),
				},
				PlanModifiers: []planmodifier.String{
					modifiers.ImmutableAfterSet(),
				},
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"json_schema": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the schema URL for JSON message validation. For example, <tt>local:///petstore.json#/definitions/Pet</tt> . To accept all input, use the string <tt>accept</tt> instead of a URL. To reject all input, use the string <tt>reject</tt> instead of a URL.", "json-schema", "").String,
				Optional:            true,
			},
			"graph_ql_schema": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the schema URL for GraphQL message validation. For example, <tt>local:///petstore-Pet.graphql</tt> . To accept all input, use the string <tt>accept</tt> instead of a URL. To reject all input, use the string <tt>reject</tt> instead of a URL.", "graphql-schema", "").String,
				Optional:            true,
			},
			"xml_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML type", "xml-type", "").AddStringEnum("xml", "wsdl").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("xml", "wsdl"),
				},
			},
			"xml_validation_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML validation mode", "xml-validation-mode", "").AddStringEnum("xsd", "soap-body").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("xsd", "soap-body"),
				},
			},
			"xml_schema_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the schema URL for XML message validation. For example, <tt>local:///petstore-Pet.xsd</tt> . To accept all input, use the string <tt>accept</tt> instead of a URL. To reject all input, use the string <tt>reject</tt> instead of a URL.", "xml-schema-url", "").String,
				Optional:            true,
			},
			"wsdl_schema_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the schema URL for WSDL message validation. For example, <tt>local:///petstore-Pet.wsdl</tt> . To accept all input, use the string <tt>accept</tt> instead of a URL. To reject all input, use the string <tt>reject</tt> instead of a URL.", "wsdl-schema-url", "").String,
				Optional:            true,
			},
			"wsdl_port_q_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the <tt>wsdl:port</tt> for the traffic to validate. The value should be a QName in the form <tt>{namespace-uri}local-part</tt> or <tt>*</tt> for all ports in the WSDL file. When specified and not <tt>*</tt> , only messages for the named port are valid.", "wsdl-port", "").String,
				Optional:            true,
			},
			"wsdl_operation_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the <tt>wsdl:operation</tt> for the traffic to validate. The value should be the unqualified name of the operation or <tt>*</tt> for all operations in the WSDL file. When specified and not <tt>*</tt> , only messages for the named operation are valid.", "wsdl-operation", "").String,
				Optional:            true,
			},
			"wsdl_message_direction_or_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the <tt>wsdl:input</tt> , <tt>wsdl:output</tt> , or <tt>wsdl:fault</tt> for the traffic to validate. The value must be the name of one or more WSDL input, output, or fault components, or <tt>#input</tt> or <tt>#output</tt> for the request and response directions respectively, or <tt>*</tt> for all inputs, outputs, and faults in the WSDL file. When specified and not <tt>*</tt> , only messages that match the specified direction or name are valid. Faults are valid for the response direction.", "wsdl-message-direction-or-name", "").String,
				Optional:            true,
			},
			"wsdl_attachment_part": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the mime:content to validate in the format <tt>mime:content/@part</tt> . The value must be the unqualified name of the message part. The name is the same as the part attribute on the corresponding <tt>mime:content</tt> component in the WSDL file. When not specified or <tt>*</tt> , the root MIME part is validated. The root MIME part is bound to a <tt>soap:Body</tt> .", "wsdl-attachment-part", "").String,
				Optional:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *APISchemaResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *APISchemaResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APISchema
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `APISchema`)
	_, err := r.pData.Client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "POST", err))
		return
	}
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APISchemaResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APISchema
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && (strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}

	if data.IsNull() {
		// Import
		data.FromBody(ctx, `APISchema`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `APISchema`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APISchemaResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APISchema
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `APISchema`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Update)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APISchemaResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APISchema
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *APISchemaResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.APISchema

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
