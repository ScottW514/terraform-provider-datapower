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
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type APIDefinitionList struct {
	ProviderTarget types.String `tfsdk:"provider_target"`
	Id             types.String `tfsdk:"id"`
	AppDomain      types.String `tfsdk:"app_domain"`
	Result         types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &APIDefinitionDataSource{}
	_ datasource.DataSourceWithConfigure = &APIDefinitionDataSource{}
)

func NewAPIDefinitionDataSource() datasource.DataSource {
	return &APIDefinitionDataSource{}
}

type APIDefinitionDataSource struct {
	pData *tfutils.ProviderData
}

func (d *APIDefinitionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_definition"
}

func (d *APIDefinitionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "An API definition uses HTTP requests to GET, PUT, POST, and DELETE data.",
		Attributes: map[string]schema.Attribute{
			"provider_target": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Target host to retrieve this data from. If not set, provider will use the top level settings.", "", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
			},
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
							MarkdownDescription: "Specify the base path on which the API is served, which is relative to the host. When the base path is not specified, the APIs are served directly under the host. The base path does not include the hostname or any additional segments for paths or operations. The base path must start but not end with slash (/). All resources in a REST API are defined relative to its base path.",
							Computed:            true,
						},
						"html_page": schema.StringAttribute{
							MarkdownDescription: "Specify the name and location of a static HTML page that the API can return. Import the file to the <tt>local:</tt> , <tt>store:</tt> , or <tt>temporary:</tt> DataPower directory.",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Type",
							Computed:            true,
						},
						"assembly": schema.StringAttribute{
							MarkdownDescription: "Specify the assembly to apply to API calls. An assembly is a rule that defines the actions to run against API requests and how to handle the processing errors.",
							Computed:            true,
						},
						"path": schema.ListAttribute{
							MarkdownDescription: "Specify the paths through which users access the API operations. A path consists of one or more HTTP operations.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"consume": schema.ListAttribute{
							MarkdownDescription: "Specify the MIME types that the API can consume. These MIME types apply to all API operations. You can override the setting for specific operations in the API operation.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"produce": schema.ListAttribute{
							MarkdownDescription: "Specify the MIME types that the API can produce. These MIME types apply to all API operations. You can override the setting for specific operations in the API operation.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"swagger_location": schema.StringAttribute{
							MarkdownDescription: "Specify the name and location of the OpenAPI document when you create the API definition from an OpenAPI document. Prepare the document as follows before you specify the location. <ol><li>When the OpenAPI document is a YAML file, convert it to JSON.</li><li>Import the JSON file to the <tt>local:</tt> or <tt>temporary:</tt> DataPower directory.</li></ol><p>When you create the API definition with API properties, this property is not applicable.</p>",
							Computed:            true,
						},
						"graphql_schema": schema.StringAttribute{
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
							MarkdownDescription: "Specify the alternative security requirements to enforce for the API as a whole. In other words, processing applies a logical <tt>OR</tt> between the security requirements. By default, the security requirement is applied to all operations in the API. However, for each API operation, you can override the API-level security by separately specifying security schemes to enforce at the operation level.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"require_api_mutual_tls": schema.BoolAttribute{
							MarkdownDescription: "API protection",
							Computed:            true,
						},
						"api_mutual_tls_source": schema.ListAttribute{
							MarkdownDescription: "Specify the sources to obtain the client certificate for mutual TLS. Because you can define multiple ways to obtain the source, ensure that you sequence the methods appropriately.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"api_mutual_tls_header_name": schema.StringAttribute{
							MarkdownDescription: "Specify the HTTP header that contains the client certificate for mutual TLS. The default value is <tt>X-Client-Certificate</tt> .",
							Computed:            true,
						},
						"properties": schema.ListNestedAttribute{
							MarkdownDescription: "Specify custom entries for API properties. An API property is a type of context variable whose value is dependent on the collection that the API is provisioned in. Collection-specific API properties allow you to use the same API definition in different collections when a property in a collection requires a unique or different value. A custom property entry defines a property and its value for one collection. For each custom property or property that needs a different value for another collection, add another entry.",
							NestedObject:        models.GetDmAPIPropertyDataSourceSchema(),
							Computed:            true,
						},
						"schemas": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the API schemas that define data types for request or message validation. An API data type consists of a name and its API schema.",
							NestedObject:        models.GetDmAPIDataTypeDefinitionDataSourceSchema(),
							Computed:            true,
						},
						"cors_toggle": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enable the API gateway to handle cross-origin resource sharing (CORS) requests. <ul><li>When enabled, the API gateway runs the API CORS action to handle all CORS requests for the API.</li><li>The routing API action matches the API to process, determines that the request is a preflight CORS request, and sets the <tt>request.attributes.isCORSPreflight</tt> flag to <tt>true</tt> .</li><li>When CORS is enabled and a preflight request is received, all assembly actions and many API actions are skipped. Only the following API actions are processed.</li><ul><li>The CORS API action configures the appropriate response headers.</li><li>The result API action sets the response headers.</li></ul><li>For all preflight requests, the security and client identification actions are always skipped.</li></ul>",
							Computed:            true,
						},
						"cors_policy": schema.StringAttribute{
							MarkdownDescription: "CORS policy",
							Computed:            true,
						},
						"activity_log_toggle": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enable API activity logging. The API gateway runs the activity log API action to generate logs. To generate log data for calls, you must enable this property and ensure that the following conditions are met. <ul><li>The logging type is not set to none.</li><li>The activity log action is added in the API rule for the API definition.</li></ul><p>When disabled, the API gateway does not generate log data.</p>",
							Computed:            true,
						},
						"content": schema.StringAttribute{
							MarkdownDescription: "Specify the content to log on success. When set to payload data, you must enable message buffering to capture all request and response data.",
							Computed:            true,
						},
						"error_content": schema.StringAttribute{
							MarkdownDescription: "Specify the content to log on error. When set to payload data, you must enable message buffering to capture all request and response data.",
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
							MarkdownDescription: "Specify whether to buffer requests and responses before processing. <ul><li>When enabled, requests and responses are buffered before processing. The message payload and the output of the invoke assembly action are read as a binary large object (BLOB).</li><li>When disabled, requests and responses are streamed. Only an asynchronous API call can read the streamed data. If the message processing requires data to be parsed at the payload level, buffering is used to capture the data.</li></ul><p>If you enable activity logging to capture payload data, you must enable message buffering to capture all request and response data.</p>",
							Computed:            true,
						},
						"deployment_state": schema.StringAttribute{
							MarkdownDescription: "Specify the deployment state of the API. By default, the deployment state is running instead of suspended.",
							Computed:            true,
						},
						"share_rate_limit_count": schema.StringAttribute{
							MarkdownDescription: "Share rate limit count",
							Computed:            true,
						},
						"return_v5responses": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to return v5-compatible responses, such as OAuth and client security error responses.",
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
							MarkdownDescription: "Specify whether to transfer documents to the server in chunks, which is based on the <tt>Transfer-Encoding: chunked</tt> header. This setting applies only to the <tt>invoke</tt> 1.5.0 policy that is deployed from API Connect from using the migration utility.",
							Computed:            true,
						},
						"set_v5request_headers": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to populate v5-compatible headers such as <tt>X-Client-IP</tt> and <tt>X-Global-Transaction-ID</tt> in the <tt>request.headers</tt> context variable.",
							Computed:            true,
						},
						"get_raw_body_value": schema.BoolAttribute{
							MarkdownDescription: "Specify whether the GatewayScript <tt>apim.getvariable()</tt> APIreturns the raw body instead of parsing. This setting applies only when the context is other than <tt>message</tt> .",
							Computed:            true,
						},
						"allowed_api_protocols": models.GetDmAPIProtocolsDataSourceSchema("Allowed API protocols", "allowed-api-protocols", ""),
						"allow_trailing_slash": schema.BoolAttribute{
							MarkdownDescription: "Allow trailing slash",
							Computed:            true,
						},
						"enforce_all_headers_case_insensitive": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enforce all headers in case-insensitive. By default, the following behavior applies to headers. <ul><li>Headers in the <tt>message.headers</tt> and <tt>request.headers</tt> contexts are case-insensitive.</li><li>Headers in a custom context, such as <tt>foo.headers</tt> , are case-sensitive. For example, <tt>foo.headers.bar</tt> and <tt>foo.headers.Bar</tt> are different headers.</li></ul>",
							Computed:            true,
						},
						"enforce_form_data_parameter": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to resolve and populate the form data parameter. <ul><li>When enabled and the operation of the API path configures any form data, resolve the payload as BLOB data and populate the corresponding fields as context variables during routing. Applicable when the <tt>Content-Type</tt> header is <tt>application/x-www-form-urlencoded</tt> or <tt>multipart/form-data</tt> .</li><li>When disabled and no required form-data parameter is specified, processing does not resolve as a BLOB and populate the parameter. Applicable when the <tt>Content-Type</tt> header is <tt>application/x-www-form-urlencoded</tt> or <tt>multipart/form-data</tt> .</li></ul>",
							Computed:            true,
						},
						"force_http500_for_soap11": schema.BoolAttribute{
							MarkdownDescription: "Force HTTP 500 for SOAP 1.1",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
						"provider_target": schema.StringAttribute{
							MarkdownDescription: tfutils.NewAttributeDescription("Target host to retrieve this data from. If not set, provider will use the top level settings.", "", "").String,
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

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *APIDefinitionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data APIDefinitionList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !d.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), d.pData.Client.GetTargetNames()))
		return
	}
	o := models.APIDefinition{
		AppDomain: data.AppDomain,
	}

	path := o.GetPath()
	if !data.Id.IsNull() {
		path = path + "/" + data.Id.ValueString()
	}

	res, err := d.pData.Client.Get(path, data.ProviderTarget)
	resFound := true
	if err != nil {
		if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(d.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		} else {
			resFound = false
		}
	}
	l := []models.APIDefinition{}
	if resFound {
		if value := res.Get(`APIDefinition`); value.Exists() {
			for _, v := range value.Array() {
				item := models.APIDefinition{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
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
