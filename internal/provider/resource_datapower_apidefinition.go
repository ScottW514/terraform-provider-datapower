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
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &APIDefinitionResource{}

func NewAPIDefinitionResource() resource.Resource {
	return &APIDefinitionResource{}
}

type APIDefinitionResource struct {
	client *client.DatapowerClient
}

func (r *APIDefinitionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apidefinition"
}

func (r *APIDefinitionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("API definition", "api-definition", "").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name of the object. Must be unique among object types in application domain.", "", "").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"app_domain": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The name of the application domain the object belongs to", "", "").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					ImmutableAfterSet(),
				},
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"api_id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("API ID", "id", "").String,
				Optional:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("API name", "name", "").String,
				Optional:            true,
			},
			"version": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("API version", "version", "").AddDefaultValue("1.0.0").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("1.0.0"),
			},
			"base_path": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Base path", "base-path", "").AddDefaultValue("/").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/"),
			},
			"html_page": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTML page", "html-page", "").String,
				Optional:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Type", "type", "").AddStringEnum("standard", "wsdl", "graphql").AddDefaultValue("standard").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("standard", "wsdl", "graphql"),
				},
				Default: stringdefault.StaticString("standard"),
			},
			"assembly": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Assembly", "assembly", "assembly").String,
				Optional:            true,
			},
			"path": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Paths", "path", "apipath").String,
				ElementType:         types.StringType,
				Required:            true,
			},
			"consume": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Consumes", "consume", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"produce": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Produces", "produce", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"swagger_location": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("OpenAPI document", "swagger-location", "").String,
				Optional:            true,
			},
			"graph_ql_schema": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("GraphQL schema location", "graphql-schema", "apischema").String,
				Optional:            true,
			},
			"wsdl_advertised_schema_location": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WSDL advertised schema location", "wsdl-advertised-schema-location", "").String,
				Optional:            true,
			},
			"wsdl_validation_schema": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WSDL validation schema location", "wsdl-validation-schema", "apischema").String,
				Optional:            true,
			},
			"security_requirement": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Security requirements", "security-req", "apisecurityrequirement").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"require_api_mutual_tls": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("API protection", "require-api-mutual-tls", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"api_mutual_tls_source": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("API protection sources", "api-mutual-tls-source", "").AddStringEnum("header", "tls_cert").String,
				ElementType:         types.StringType,
				Optional:            true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(stringvalidator.OneOf("header", "tls_cert")),
				},
			},
			"api_mutual_tls_header_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("API protection HTTP header", "api-mutual-tls-header-name", "").AddDefaultValue("X-Client-Certificate").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("X-Client-Certificate"),
			},
			"properties": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Custom properties", "property", "").String,
				NestedObject:        models.DmAPIPropertyResourceSchema,
				Optional:            true,
			},
			"schemas": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Schemas", "schema", "").String,
				NestedObject:        models.DmAPIDataTypeDefinitionResourceSchema,
				Optional:            true,
			},
			"cors_toggle": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable CORS", "cors", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"cors_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("CORS policy", "cors-policy", "corspolicy").String,
				Optional:            true,
			},
			"activity_log_toggle": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable logging", "activity-log", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"content": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Content on success", "success-content", "").AddStringEnum("none", "activity", "header", "payload").AddDefaultValue("activity").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("none", "activity", "header", "payload"),
				},
				Default: stringdefault.StaticString("activity"),
			},
			"error_content": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Content on error", "error-content", "").AddStringEnum("none", "activity", "header", "payload").AddDefaultValue("payload").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("none", "activity", "header", "payload"),
				},
				Default: stringdefault.StaticString("payload"),
			},
			"preserved_request_header": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Request headers to preserve", "preserved-request-header", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"preserved_response_header": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Response header to preserve", "preserved-response-header", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"message_buffering": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable message buffering", "message-buffering", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"deployment_state": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Deployment state", "deployment-state", "").AddStringEnum("running", "suspended").AddDefaultValue("running").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("running", "suspended"),
				},
				Default: stringdefault.StaticString("running"),
			},
			"share_rate_limit_count": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Share rate limit count", "share-count", "").AddStringEnum("unset", "yes", "no").AddDefaultValue("unset").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("unset", "yes", "no"),
				},
				Default: stringdefault.StaticString("unset"),
			},
			"return_v5_responses": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Return v5 responses", "return-v5-responses", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"copy_id_headers_to_message": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Copy ID headers to message", "copy-id-headers-to-message", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"enforce_required_params": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enforce required parameters", "enforce-required-params", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"allow_chunked_uploads": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow chunked uploads", "allow-chunked-uploads", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"set_v5_request_headers": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Set v5 request headers", "set-v5-request-headers", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"get_raw_body_value": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Get raw body value", "get-raw-body-value", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"allowed_api_protocols": models.GetDmAPIProtocolsResourceSchema("Allowed API protocols", "allowed-api-protocols", "", false),
			"allow_trailing_slash": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow trailing slash", "allow-trailing-slash", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"enforce_all_headers_case_insensitive": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enforce all headers as case-insensitive", "enforce-all-headers-case-insensitive", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"enforce_form_data_parameter": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enforce form data parameter", "enforce-form-data-parameter", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"force_http500_for_soap11": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Force HTTP 500 for SOAP 1.1", "force-http-500-for-soap11", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
}

func (r *APIDefinitionResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *APIDefinitionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.APIDefinition

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `APIDefinition`)
	_, err := r.client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "POST", err))
		return
	}

	_, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s", "POST", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIDefinitionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.APIDefinition

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && (strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}

	if data.IsNull() {
		// Import
		data.FromBody(ctx, `APIDefinition`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `APIDefinition`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIDefinitionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.APIDefinition

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `APIDefinition`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}
	_, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s", "POST", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIDefinitionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.APIDefinition

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}

	resp.State.RemoveResource(ctx)
}
