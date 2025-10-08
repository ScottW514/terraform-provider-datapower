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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type StylePolicyActionList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &StylePolicyActionDataSource{}
	_ datasource.DataSourceWithConfigure = &StylePolicyActionDataSource{}
)

func NewStylePolicyActionDataSource() datasource.DataSource {
	return &StylePolicyActionDataSource{}
}

type StylePolicyActionDataSource struct {
	pData *tfutils.ProviderData
}

func (d *StylePolicyActionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_style_policy_action"
}

func (d *StylePolicyActionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Define processing actions for the processing rules in a processing policy.",
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
						"type": schema.StringAttribute{
							MarkdownDescription: "Action type",
							Computed:            true,
						},
						"input": schema.StringAttribute{
							MarkdownDescription: "Specify the input context for the action, which identifies the context that contains the document to act. Enter the context name, the string <tt>PIPE</tt> for streaming mode, or the string <tt>INPUT</tt> to identify the original input of the policy rule.",
							Computed:            true,
						},
						"transform": schema.StringAttribute{
							MarkdownDescription: "Specify the location of the XSL stylesheet or transform file. The location uses one of the following formats. <ul><li>Use a URL, for example, <tt>store:///myTest.xsl</tt></li><li>Use a context variable that expands to a URL, for example, <tt>var://context/contextName/varName</tt></li><li>Use a context, for example, <tt>var://context/Name</tt> or <tt>var://context/Name/</tt> . The context runs as a stylesheet.</li></ul>",
							Computed:            true,
						},
						"parse_settings_reference": models.GetDmDynamicParseSettingsReferenceDataSourceSchema("Specify the configuration that defines the constraints against the documents to parse. Use any or all of the following ways to define the constraints. <ul><li>Specify a URL reference from which to retrieve the constraints definition.</li><li>Specify a literal configuration string in XML or JSON format that contains the constraints definition.</li><li>Specify an instance of the parse settings configuration object to retrieve constraints definition.</li></ul><p>Precedence rules apply when the constraint for the same aspect of an input document is configured with more than one method.</p>", "parse-settings-reference", ""),
						"parse_metrics_result_type": schema.StringAttribute{
							MarkdownDescription: "Parse metrics result type",
							Computed:            true,
						},
						"parse_metrics_result_location": schema.StringAttribute{
							MarkdownDescription: "Parse metrics result location",
							Computed:            true,
						},
						"input_language": schema.StringAttribute{
							MarkdownDescription: "Input language",
							Computed:            true,
						},
						"dfdl_input_root_name": schema.StringAttribute{
							MarkdownDescription: "Specify the root element in the DFDL model from which to start a parse. This property is only meaningful in the context of a DFDL parse. <p>For the input root name, specify the global xsd:element in the XSD file to use to begin the parsing of binary input. The input root name can be selected from the specified XSD file or specified as a variable. For a variable, it must resolve to a valid namespace URL with the name between braces ({}) as a prefix to the local part. For instance, if in the DFDL Schema, the target namespace is \"http://example.com/messages\" and the local element is &lt;xsd:element name=\"Message\">...&lt;/xsd:element>, the variable must resolve to {http://example.com/messages}Message.</p><p>The schema author might specify the root parse element by using the ibmSchExtn:docRoot=\"true\" element within the schema. For instance, &lt;xsd:element ibmSchExtn:docRoot=\"true\" name=\"Message\">...&lt;/xsd:element>. In this case, the input root name shows in the selection as the element name followed by (@ibmSchExtn:docRoot=\"true\"). Use of another value for the DFDL input root name overrides the value that is specified in the schema.</p>",
							Computed:            true,
						},
						"input_descriptor": schema.StringAttribute{
							MarkdownDescription: "Specify the input descriptor as understood according to the input language. <ul><li>When DFDL, the input descriptor must be a URL to a schema file that defines the input. The input descriptor can be a URL to a directory from which you can select a schema file a variable that resolves to a schema at run time.</li><li>When XML, do not specify an input descriptor.</li><li>When XSD, the input descriptor must be a URL to an XML schema.</li><li>When JSON, do not specify an input descriptor.</li></ul>",
							Computed:            true,
						},
						"output_descriptor": schema.StringAttribute{
							MarkdownDescription: "Specify the output descriptor as understood by the output language. When DFDL, the output descriptor must be a URL to a DFDL schema to serialize the output.",
							Computed:            true,
						},
						"transform_language": schema.StringAttribute{
							MarkdownDescription: "Transform language",
							Computed:            true,
						},
						"output_language": schema.StringAttribute{
							MarkdownDescription: "Output language",
							Computed:            true,
						},
						"tx_map": schema.StringAttribute{
							MarkdownDescription: "Specify the location of the Transformation Extender map file. The generated map file is either on the DataPower Gateway or on a remote HTTP or HTTPS server. Use one of the following formats. <ul><li>When the file is local, use &lt;directory>:///&lt;file></li><li>When the file is remote, use HTTP://&lt;path_qualified_file> or HTTPS://&lt;path_qualified_file></li></ul>",
							Computed:            true,
						},
						"gateway_script_location": schema.StringAttribute{
							MarkdownDescription: "Specify the location of the GatewayScript file. The file location can be specified in one of the following formats. <ul><li>As a URL, where the file is in the <tt>local:</tt> , <tt>store:</tt> , or <tt>temporary:</tt> directory.</li><li>As a context variable that expands to a URL, such as <tt>var://context/contextName/varName</tt> .</li><li>As a context, for example, <tt>var://context/Name</tt> or <tt>var://context/Name/</tt> . The context content runs as GatewayScript.</li></ul>",
							Computed:            true,
						},
						"action_debug": schema.BoolAttribute{
							MarkdownDescription: "Enable GatewayScript debug",
							Computed:            true,
						},
						"tx_top_level_map": schema.StringAttribute{
							MarkdownDescription: "Specify the name of the Transformation Extender map in the map file. A map file can contain more than one map. When not specified, the transform uses the first map in the file.",
							Computed:            true,
						},
						"tx_mode": schema.StringAttribute{
							MarkdownDescription: "Specify the Transformation Extender mode to run the map. DPA is the only mode.",
							Computed:            true,
						},
						"tx_audit_log": schema.StringAttribute{
							MarkdownDescription: "ITX audit log",
							Computed:            true,
						},
						"output": schema.StringAttribute{
							MarkdownDescription: "Specify the output context for the action, which identifies the context that receives the document when the action completes. Enter the context name, the string <tt>PIPE</tt> for streaming mode, or the string <tt>OUTPUT</tt> to identify the final output of the policy rule.",
							Computed:            true,
						},
						"no_transcode_utf8": schema.BoolAttribute{
							MarkdownDescription: "Specify whether the output from the convert action retains the input encoding or uses ISO 8859-1. An encoding is also known as a character set. For illustrative purposes, assume UTF-8 is the input encoding. <ul><li>When enabled and the input encoding is UTF-8, the output is UTF-8.</li><li>When not enabled and the input encoding is UTF-8, the output is ISO 8859-1. This behavior is the default behavior.</li></ul>",
							Computed:            true,
						},
						"named_in_out_location_type": schema.StringAttribute{
							MarkdownDescription: "Specify how to locate named inputs and outputs. Use values that are appropriate for your Transformation Extender configuration.",
							Computed:            true,
						},
						"named_inputs": schema.ListNestedAttribute{
							MarkdownDescription: "Named inputs",
							NestedObject:        models.GetDmNamedInOutDataSourceSchema(),
							Computed:            true,
						},
						"named_outputs": schema.ListNestedAttribute{
							MarkdownDescription: "Named outputs",
							NestedObject:        models.GetDmNamedInOutDataSourceSchema(),
							Computed:            true,
						},
						"destination": schema.StringAttribute{
							MarkdownDescription: "Specify the location of the resource, which might be either the source or the destination. Specify the location as either a URL or as a variable that expands to a URL.",
							Computed:            true,
						},
						"schema_url": schema.StringAttribute{
							MarkdownDescription: "Specify XML schema for document validation regardless of any <tt>xsi:schemaLocation</tt> attributes contained with the document. Identify the schema with one of the following formats. <ul><li>Use a URL, for example, <tt>store:///valHigh.xsd</tt></li><li>Use a context variable that expands to a URL, for example, <tt>var://context/contextName/varName</tt></li><li>Use a context, for example, <tt>var://context/Name</tt> or <tt>var://context/Name/</tt> . The context runs as a schema validation.</li></ul>",
							Computed:            true,
						},
						"json_schema_url": schema.StringAttribute{
							MarkdownDescription: "Specify the JSON schema for JSON document validation. Identify the schema with one of the following formats. <ul><li>Use a URL, for example, <tt>local:///valHigh.jsv</tt></li><li>Use a context variable that expands to a URL, for example, <tt>var://context/contextName/varName</tt></li><li>Use a context, for example, <tt>var://context/Name</tt> or <tt>var://context/Name/</tt> . The context runs as a JSON schema validation.</li></ul>",
							Computed:            true,
						},
						"wsdl_url": schema.StringAttribute{
							MarkdownDescription: "Specify the URL of the WSDL file that defines the operations to use during the validate action. The WSDL file can reside on the local system or on the network. By default, the WSDL validation always applies to the entire input message, which can be modified by compile options on the XML manager. <p>Identify the WSDL with one of the following formats.</p><ul><li>Use a URL, for example, <tt>local:///myTest.wsdl</tt></li><li>Use a context variable that expands to a URL, for example, <tt>var://context/contextName/varName</tt></li><li>Use a context, for example, <tt>var://context/Name</tt> or <tt>var://context/Name/</tt> . The context runs as a WSDL validation.</li></ul>",
							Computed:            true,
						},
						"policy": schema.StringAttribute{
							MarkdownDescription: "URL rewrite policy",
							Computed:            true,
						},
						"aaa": schema.StringAttribute{
							MarkdownDescription: "AAA policy",
							Computed:            true,
						},
						"dynamic_schema": schema.StringAttribute{
							MarkdownDescription: "Dynamic schema",
							Computed:            true,
						},
						"dynamic_stylesheet": schema.StringAttribute{
							MarkdownDescription: "Dynamic stylesheet",
							Computed:            true,
						},
						"input_conversion": schema.StringAttribute{
							MarkdownDescription: "Input conversion",
							Computed:            true,
						},
						"xpath": schema.StringAttribute{
							MarkdownDescription: "Specify the XPath expression to apply to the input context. Enter the XPath expression or a variable in the <tt>var://context/name</tt> format that expands to an XPath expression.",
							Computed:            true,
						},
						"variable": schema.StringAttribute{
							MarkdownDescription: "Specify the variable URL in one of the following formats. <ul><li>var://context/CONTEXT-NAME/VAR-NAME <p>This format is the primary way to reference variables. var://context/CONTEXT-NAME/_roottree is a special variable that holds the value of the context when used as input to an action. var://context/CONTEXT-NAME (or var://context/CONTEXT-NAME/) is treated as shorthand for var://context/CONTEXT-NAME/_roottree.</p></li><li>var://local/VAR-NAME <p>This format can be used to reference variables in the input or output context. Because this reference is context-sensitive, The use of var://context/CONTEXT-NAME/VAR-NAME is recommended.</p></li><li>var://system/CONTEXT-NAME/VAR-NAME <p>This format is used to reference global variables. These variables are rarely used.</p></li><li>var://service/SERVICE-NAME <p>This format is used to reference certain internal state variables. These variables are defined by the firmware.</p></li></ul>",
							Computed:            true,
						},
						"value": schema.StringAttribute{
							MarkdownDescription: "Specify the variable value. The value can be a number, a string, or another variable URL.",
							Computed:            true,
						},
						"ssl_client_config_type": schema.StringAttribute{
							MarkdownDescription: "TLS client type",
							Computed:            true,
						},
						"ssl_client_cred": schema.StringAttribute{
							MarkdownDescription: "TLS client profile",
							Computed:            true,
						},
						"attachment_uri": schema.StringAttribute{
							MarkdownDescription: "Specify the attachment to strip. When omitted, all attachments are stripped from the context.",
							Computed:            true,
						},
						"stylesheet_parameters": schema.ListNestedAttribute{
							MarkdownDescription: "Stylesheet parameters",
							NestedObject:        models.GetDmStylesheetParameterDataSourceSchema(),
							Computed:            true,
						},
						"error_mode": schema.StringAttribute{
							MarkdownDescription: "Error mode",
							Computed:            true,
						},
						"error_input": schema.StringAttribute{
							MarkdownDescription: "Error input",
							Computed:            true,
						},
						"error_output": schema.StringAttribute{
							MarkdownDescription: "Error output",
							Computed:            true,
						},
						"rule": schema.StringAttribute{
							MarkdownDescription: "Processing rule",
							Computed:            true,
						},
						"output_type": schema.StringAttribute{
							MarkdownDescription: "Output type",
							Computed:            true,
						},
						"log_level": schema.StringAttribute{
							MarkdownDescription: "Log level",
							Computed:            true,
						},
						"log_type": schema.StringAttribute{
							MarkdownDescription: "Log type",
							Computed:            true,
						},
						"transactional": schema.BoolAttribute{
							MarkdownDescription: "Transactional",
							Computed:            true,
						},
						"checkpoint_event": schema.StringAttribute{
							MarkdownDescription: "Event",
							Computed:            true,
						},
						"slm_policy": schema.StringAttribute{
							MarkdownDescription: "SLM policy",
							Computed:            true,
						},
						"sql_data_source": schema.StringAttribute{
							MarkdownDescription: "SQL Data Source",
							Computed:            true,
						},
						"sql_text": schema.StringAttribute{
							MarkdownDescription: "SQL text",
							Computed:            true,
						},
						"soap_validation": schema.StringAttribute{
							MarkdownDescription: "Specify which parts of the SOAP message to validate. This setting does not affect the validation of the input context to ensure that it is a valid document.",
							Computed:            true,
						},
						"sql_source_type": schema.StringAttribute{
							MarkdownDescription: "SQL input method",
							Computed:            true,
						},
						"jose_serialization_type": schema.StringAttribute{
							MarkdownDescription: "Serialization",
							Computed:            true,
						},
						"jwe_enc_algorithm": schema.StringAttribute{
							MarkdownDescription: "Algorithm",
							Computed:            true,
						},
						"jws_signature_object": schema.StringAttribute{
							MarkdownDescription: "Signature",
							Computed:            true,
						},
						"jwe_header_object": schema.StringAttribute{
							MarkdownDescription: "JWE Header",
							Computed:            true,
						},
						"jose_verify_type": schema.StringAttribute{
							MarkdownDescription: "Identifier type",
							Computed:            true,
						},
						"jose_decrypt_type": schema.StringAttribute{
							MarkdownDescription: "Identifier Type",
							Computed:            true,
						},
						"signature_identifier": schema.ListAttribute{
							MarkdownDescription: "Signature Identifiers",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"recipient_identifier": schema.ListAttribute{
							MarkdownDescription: "Recipient Identifiers",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"single_certificate": schema.StringAttribute{
							MarkdownDescription: "Certificate",
							Computed:            true,
						},
						"single_key": schema.StringAttribute{
							MarkdownDescription: "Private Key",
							Computed:            true,
						},
						"single_sskey": schema.StringAttribute{
							MarkdownDescription: "Shared Secret Key",
							Computed:            true,
						},
						"jwe_direct_key_object": schema.StringAttribute{
							MarkdownDescription: "Direct Key",
							Computed:            true,
						},
						"jws_verify_strip_signature": schema.BoolAttribute{
							MarkdownDescription: "Strip Signature",
							Computed:            true,
						},
						"asynchronous": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to mark the action as asynchronous. When asynchronous, the action does not need to complete for the next action to start.",
							Computed:            true,
						},
						"condition": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the conditions to check and action to run when a match is found A single condition maps an XPath expression to search for in the input context to an action to run when the condition is found. When no match is found, other conditions can be checked.",
							NestedObject:        models.GetDmConditionDataSourceSchema(),
							Computed:            true,
						},
						"results_mode": schema.StringAttribute{
							MarkdownDescription: "Specify the processing mode for multiple targets. The default behavior, is first available.",
							Computed:            true,
						},
						"retry_count": schema.Int64Attribute{
							MarkdownDescription: "Specify the number of connection attempts. The default value is 0, which indicates that the operation fails immediately if the connection fails.",
							Computed:            true,
						},
						"retry_interval": schema.Int64Attribute{
							MarkdownDescription: "Specify the duration between connection attempts in milliseconds. The default value is 1000.",
							Computed:            true,
						},
						"multiple_outputs": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to place parallel outputs into separate contexts. A results action can target several targets simultaneously by specifying a variable that contains an XML nodeset as the results target. <ul><li>When enabled, the context for the output context is ignored. The status of the individual attempts to reach the targets are recorded in separate contexts by appending a number. For example, the resultant context name is <tt>ctx_1</tt> for the <tt>ctx</tt> context name.</li><li>When not enabled, only one result is kept. <ul><li>In require-all mode, keep the first target that failed and has no remaining attempts or the last target to succeed.</li><li>In attempt-all mode, the last target attempted.</li><li>In first-available mode, the first target to succeed or the last to fail.</li></ul></li></ul>",
							Computed:            true,
						},
						"iterator_type": schema.StringAttribute{
							MarkdownDescription: "Iterator type",
							Computed:            true,
						},
						"iterator_expression": schema.StringAttribute{
							MarkdownDescription: "XPath expression",
							Computed:            true,
						},
						"iterator_count": schema.Int64Attribute{
							MarkdownDescription: "Specify the number of times to run the loop action. Enter a value in the range 1 - 32768.",
							Computed:            true,
						},
						"loop_action": schema.StringAttribute{
							MarkdownDescription: "Loop action",
							Computed:            true,
						},
						"async_action": schema.ListAttribute{
							MarkdownDescription: "Asynchronous actions",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"timeout": schema.Int64Attribute{
							MarkdownDescription: "Specify the duration to wait for asynchronous actions to complete in milliseconds. After this period, an error is logged. A value of 0 indicates to wait forever.",
							Computed:            true,
						},
						"wsdl_port_q_name": schema.StringAttribute{
							MarkdownDescription: "Specify the QName of the WSDL port that defines the service traffic to validate. The value must be a QName in the <tt>{namespace-uri}local-part</tt> format or <tt>*</tt> for all ports. When specified and not <tt>*</tt> , only messages that are defined for the named port are considered valid.",
							Computed:            true,
						},
						"wsdl_operation_name": schema.StringAttribute{
							MarkdownDescription: "Specify the name of the WSDL operation that defines the service traffic to validate. The value must be the unqualified name of the operation or <tt>*</tt> for all operations. When specified and not <tt>*</tt> , only messages that are defined for operations that match the specified name are considered valid.",
							Computed:            true,
						},
						"wsdl_message_direction_or_name": schema.StringAttribute{
							MarkdownDescription: "Specify the name or direction of the WSDL input, output, or fault that defines the service traffic to validate. Use one of the following values. <ul><li>The name of one or more WSDL input, output, or fault components.</li><ul><li><tt><i>operation</i> Request</tt> for a specific request.</li><li><tt><i>operation</i> Response\"</tt> for a specific response.</li><li>The value of the <tt>name</tt> attribute for the <tt>&lt;fault></tt> element.</li></ul><li><tt>#input</tt> for the request direction in the WSDL file.</li><li><tt>#output</tt> for the response direction in the WSDL file.</li><li><tt>*</tt> for all inputs, outputs, and faults in the WSDL file.</li></ul><p>When specified and not <tt>*</tt> , only messages that are defined for inputs, outputs, and faults that match the specified name or direction are considered valid. Faults are considered valid for the response direction.</p>",
							Computed:            true,
						},
						"wsdl_attachment_part": schema.StringAttribute{
							MarkdownDescription: "Specify the name of the WSDL message part that defines the content of a MIME attachment (mime:content/@part) to validate. The value must be the unqualified name of the message part. The name is the same as the part attribute on the corresponding <tt>mime:content</tt> component in the WSDL file. When not defined or <tt>*</tt> , the root MIME part is validated. The root MIME part is bound to a <tt>soap:body</tt> .",
							Computed:            true,
						},
						"method_rewrite_type": schema.StringAttribute{
							MarkdownDescription: "Method",
							Computed:            true,
						},
						"method_type": schema.StringAttribute{
							MarkdownDescription: "Method",
							Computed:            true,
						},
						"method_type2": schema.StringAttribute{
							MarkdownDescription: "Method",
							Computed:            true,
						},
						"policy_key": schema.StringAttribute{
							MarkdownDescription: "<b>Do not modify this value.</b> The DataPower Gateway uses this identifier to store the output from the action. The output can be used for external monitoring.",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *StylePolicyActionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *StylePolicyActionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data StylePolicyActionList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.StylePolicyAction{
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
	l := []models.StylePolicyAction{}
	if resFound {
		if value := res.Get(`StylePolicyAction`); value.Exists() {
			for _, v := range value.Array() {
				item := models.StylePolicyAction{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.StylePolicyActionObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.StylePolicyActionObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
