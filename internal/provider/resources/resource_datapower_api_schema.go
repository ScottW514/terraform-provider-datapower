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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

var _ resource.Resource = &APISchemaResource{}
var _ resource.ResourceWithImportState = &APISchemaResource{}

func NewAPISchemaResource() resource.Resource {
	return &APISchemaResource{}
}

type APISchemaResource struct {
	pData *tfutils.ProviderData
}

func (r *APISchemaResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_schema"
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
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
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
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
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
			"graphql_schema": schema.StringAttribute{
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
				MarkdownDescription: tfutils.NewAttributeDescription("XML validation mode", "xml-validation-mode", "").AddStringEnum("xsd", "soap-body").AddNotValidWhen(models.APISchemaXMLValidationModeIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("xsd", "soap-body"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.APISchemaXMLValidationModeIgnoreVal, false),
				},
			},
			"xml_schema_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the schema URL for XML message validation. For example, <tt>local:///petstore-Pet.xsd</tt> . To accept all input, use the string <tt>accept</tt> instead of a URL. To reject all input, use the string <tt>reject</tt> instead of a URL.", "xml-schema-url", "").AddNotValidWhen(models.APISchemaXMLSchemaURLIgnoreVal.String()).String,
				Optional:            true,
			},
			"wsdl_schema_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the schema URL for WSDL message validation. For example, <tt>local:///petstore-Pet.wsdl</tt> . To accept all input, use the string <tt>accept</tt> instead of a URL. To reject all input, use the string <tt>reject</tt> instead of a URL.", "wsdl-schema-url", "").AddNotValidWhen(models.APISchemaWSDLSchemaURLIgnoreVal.String()).String,
				Optional:            true,
			},
			"wsdl_port_q_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the <tt>wsdl:port</tt> for the traffic to validate. The value should be a QName in the form <tt>{namespace-uri}local-part</tt> or <tt>*</tt> for all ports in the WSDL file. When specified and not <tt>*</tt> , only messages for the named port are valid.", "wsdl-port", "").AddNotValidWhen(models.APISchemaWSDLPortQNameIgnoreVal.String()).String,
				Optional:            true,
			},
			"wsdl_operation_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the <tt>wsdl:operation</tt> for the traffic to validate. The value should be the unqualified name of the operation or <tt>*</tt> for all operations in the WSDL file. When specified and not <tt>*</tt> , only messages for the named operation are valid.", "wsdl-operation", "").AddNotValidWhen(models.APISchemaWSDLOperationNameIgnoreVal.String()).String,
				Optional:            true,
			},
			"wsdl_message_direction_or_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the <tt>wsdl:input</tt> , <tt>wsdl:output</tt> , or <tt>wsdl:fault</tt> for the traffic to validate. The value must be the name of one or more WSDL input, output, or fault components, or <tt>#input</tt> or <tt>#output</tt> for the request and response directions respectively, or <tt>*</tt> for all inputs, outputs, and faults in the WSDL file. When specified and not <tt>*</tt> , only messages that match the specified direction or name are valid. Faults are valid for the response direction.", "wsdl-message-direction-or-name", "").AddNotValidWhen(models.APISchemaWSDLMessageDirectionOrNameIgnoreVal.String()).String,
				Optional:            true,
			},
			"wsdl_attachment_part": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the mime:content to validate in the format <tt>mime:content/@part</tt> . The value must be the unqualified name of the message part. The name is the same as the part attribute on the corresponding <tt>mime:content</tt> component in the WSDL file. When not specified or <tt>*</tt> , the root MIME part is validated. The root MIME part is bound to a <tt>soap:Body</tt> .", "wsdl-attachment-part", "").AddNotValidWhen(models.APISchemaWSDLAttachmentPartIgnoreVal.String()).String,
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
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), body)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Resource already exists. Failed to update resource, got error: %s", err))
				return
			}
			resp.Diagnostics.AddWarning("Warning", "Resource already exists. Existing resource was updated.")
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create resource, got error: %s", err))
			return
		}
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

	data.UpdateFromBody(ctx, `APISchema`, res)

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
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			resp.Diagnostics.AddWarning("Resource Conflict", fmt.Sprintf("Resource is no longer tracked by Terraform, but may need to be manually deleted on DataPower host. Got error: %s", err))
		} else if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete resource, got error: %s", err))
			return
		}
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *APISchemaResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError("Invalid Import ID", "Expected format: <app_domain>/<id>. Got: "+req.ID)
		return
	}

	appDomain := parts[0]
	id := parts[1]

	if !regexp.MustCompile("^[a-zA-Z0-9_-]+$").MatchString(id) || len(id) < 1 || len(id) > 128 {
		resp.Diagnostics.AddError("Invalid ID", "ID must be 1-128 characters and match pattern ^[a-zA-Z0-9_-]+$")
		return
	}
	if !regexp.MustCompile("^[a-zA-Z0-9_-]+$").MatchString(appDomain) || len(appDomain) < 1 || len(appDomain) > 128 {
		resp.Diagnostics.AddError("Invalid Application Domain", "Application domain must be 1-128 characters and match pattern ^[a-zA-Z0-9_-]+$")
		return
	}

	var data models.APISchema
	data.AppDomain = types.StringValue(appDomain)
	data.Id = types.StringValue(id)
	res, err := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("Resource was not found, got error: %s", err))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		}
		return
	}

	data.FromBody(ctx, `APISchema`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APISchemaResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.APISchema

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
