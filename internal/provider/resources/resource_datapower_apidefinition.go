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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &APIDefinitionResource{}

func NewAPIDefinitionResource() resource.Resource {
	return &APIDefinitionResource{}
}

type APIDefinitionResource struct {
	pData *tfutils.ProviderData
}

func (r *APIDefinitionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apidefinition"
}

func (r *APIDefinitionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("An API definition uses HTTP requests to GET, PUT, POST, and DELETE data.", "api-definition", "").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the base path on which the API is served, which is relative to the host. When the base path is not specified, the APIs are served directly under the host. The base path does not include the hostname or any additional segments for paths or operations. The base path must start but not end with slash (/). All resources in a REST API are defined relative to its base path.", "base-path", "").AddDefaultValue("/").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/"),
			},
			"html_page": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name and location of a static HTML page that the API can return. Import the file to the <tt>local:</tt> , <tt>store:</tt> , or <tt>temporary:</tt> DataPower directory.", "html-page", "").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the assembly to apply to API calls. An assembly is a rule that defines the actions to run against API requests and how to handle the processing errors.", "assembly", "assembly").String,
				Optional:            true,
			},
			"path": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the paths through which users access the API operations. A path consists of one or more HTTP operations.", "path", "apipath").String,
				ElementType:         types.StringType,
				Required:            true,
			},
			"consume": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the MIME types that the API can consume. These MIME types apply to all API operations. You can override the setting for specific operations in the API operation.", "consume", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"produce": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the MIME types that the API can produce. These MIME types apply to all API operations. You can override the setting for specific operations in the API operation.", "produce", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"swagger_location": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name and location of the OpenAPI document when you create the API definition from an OpenAPI document. Prepare the document as follows before you specify the location. <ol><li>When the OpenAPI document is a YAML file, convert it to JSON.</li><li>Import the JSON file to the <tt>local:</tt> or <tt>temporary:</tt> DataPower directory.</li></ol><p>When you create the API definition with API properties, this property is not applicable.</p>", "swagger-location", "").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the alternative security requirements to enforce for the API as a whole. In other words, processing applies a logical <tt>OR</tt> between the security requirements. By default, the security requirement is applied to all operations in the API. However, for each API operation, you can override the API-level security by separately specifying security schemes to enforce at the operation level.", "security-req", "apisecurityrequirement").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the sources to obtain the client certificate for mutual TLS. Because you can define multiple ways to obtain the source, ensure that you sequence the methods appropriately.", "api-mutual-tls-source", "").AddStringEnum("header", "tls_cert").String,
				ElementType:         types.StringType,
				Optional:            true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(stringvalidator.OneOf("header", "tls_cert")),
				},
			},
			"api_mutual_tls_header_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the HTTP header that contains the client certificate for mutual TLS. The default value is <tt>X-Client-Certificate</tt> .", "api-mutual-tls-header-name", "").AddDefaultValue("X-Client-Certificate").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("X-Client-Certificate"),
			},
			"properties": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify custom entries for API properties. An API property is a type of context variable whose value is dependent on the collection that the API is provisioned in. Collection-specific API properties allow you to use the same API definition in different collections when a property in a collection requires a unique or different value. A custom property entry defines a property and its value for one collection. For each custom property or property that needs a different value for another collection, add another entry.", "property", "").String,
				NestedObject:        models.DmAPIPropertyResourceSchema,
				Optional:            true,
			},
			"schemas": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the API schemas that define data types for request or message validation. An API data type consists of a name and its API schema.", "schema", "").String,
				NestedObject:        models.DmAPIDataTypeDefinitionResourceSchema,
				Optional:            true,
			},
			"cors_toggle": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable the API gateway to handle cross-origin resource sharing (CORS) requests. <ul><li>When enabled, the API gateway runs the API CORS action to handle all CORS requests for the API.</li><li>The routing API action matches the API to process, determines that the request is a preflight CORS request, and sets the <tt>request.attributes.isCORSPreflight</tt> flag to <tt>true</tt> .</li><li>When CORS is enabled and a preflight request is received, all assembly actions and many API actions are skipped. Only the following API actions are processed.</li><ul><li>The CORS API action configures the appropriate response headers.</li><li>The result API action sets the response headers.</li></ul><li>For all preflight requests, the security and client identification actions are always skipped.</li></ul>", "cors", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"cors_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("CORS policy", "cors-policy", "corspolicy").String,
				Optional:            true,
			},
			"activity_log_toggle": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable API activity logging. The API gateway runs the activity log API action to generate logs. To generate log data for calls, you must enable this property and ensure that the following conditions are met. <ul><li>The logging type is not set to none.</li><li>The activity log action is added in the API rule for the API definition.</li></ul><p>When disabled, the API gateway does not generate log data.</p>", "activity-log", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"content": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the content to log on success. When set to payload data, you must enable message buffering to capture all request and response data.", "success-content", "").AddStringEnum("none", "activity", "header", "payload").AddDefaultValue("activity").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("none", "activity", "header", "payload"),
				},
				Default: stringdefault.StaticString("activity"),
			},
			"error_content": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the content to log on error. When set to payload data, you must enable message buffering to capture all request and response data.", "error-content", "").AddStringEnum("none", "activity", "header", "payload").AddDefaultValue("payload").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to buffer requests and responses before processing. <ul><li>When enabled, requests and responses are buffered before processing. The message payload and the output of the invoke assembly action are read as a binary large object (BLOB).</li><li>When disabled, requests and responses are streamed. Only an asynchronous API call can read the streamed data. If the message processing requires data to be parsed at the payload level, buffering is used to capture the data.</li></ul><p>If you enable activity logging to capture payload data, you must enable message buffering to capture all request and response data.</p>", "message-buffering", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"deployment_state": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the deployment state of the API. By default, the deployment state is running instead of suspended.", "deployment-state", "").AddStringEnum("running", "suspended").AddDefaultValue("running").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to return v5-compatible responses, such as OAuth and client security error responses.", "return-v5-responses", "").AddDefaultValue("false").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to transfer documents to the server in chunks, which is based on the <tt>Transfer-Encoding: chunked</tt> header. This setting applies only to the <tt>invoke</tt> 1.5.0 policy that is deployed from API Connect from using the migration utility.", "allow-chunked-uploads", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"set_v5_request_headers": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to populate v5-compatible headers such as <tt>X-Client-IP</tt> and <tt>X-Global-Transaction-ID</tt> in the <tt>request.headers</tt> context variable.", "set-v5-request-headers", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"get_raw_body_value": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the GatewayScript <tt>apim.getvariable()</tt> APIreturns the raw body instead of parsing. This setting applies only when the context is other than <tt>message</tt> .", "get-raw-body-value", "").AddDefaultValue("false").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enforce all headers in case-insensitive. By default, the following behavior applies to headers. <ul><li>Headers in the <tt>message.headers</tt> and <tt>request.headers</tt> contexts are case-insensitive.</li><li>Headers in a custom context, such as <tt>foo.headers</tt> , are case-sensitive. For example, <tt>foo.headers.bar</tt> and <tt>foo.headers.Bar</tt> are different headers.</li></ul>", "enforce-all-headers-case-insensitive", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"enforce_form_data_parameter": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to resolve and populate the form data parameter. <ul><li>When enabled and the operation of the API path configures any form data, resolve the payload as BLOB data and populate the corresponding fields as context variables during routing. Applicable when the <tt>Content-Type</tt> header is <tt>application/x-www-form-urlencoded</tt> or <tt>multipart/form-data</tt> .</li><li>When disabled and no required form-data parameter is specified, processing does not resolve as a BLOB and populate the parameter. Applicable when the <tt>Content-Type</tt> header is <tt>application/x-www-form-urlencoded</tt> or <tt>multipart/form-data</tt> .</li></ul>", "enforce-form-data-parameter", "").AddDefaultValue("true").String,
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
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *APIDefinitionResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *APIDefinitionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APIDefinition
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `APIDefinition`)
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

func (r *APIDefinitionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APIDefinition
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
		data.FromBody(ctx, `APIDefinition`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `APIDefinition`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIDefinitionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APIDefinition
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `APIDefinition`))
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

func (r *APIDefinitionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APIDefinition
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

func (r *APIDefinitionResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.APIDefinition

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
