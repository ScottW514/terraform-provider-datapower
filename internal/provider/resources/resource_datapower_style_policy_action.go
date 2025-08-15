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

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
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

var _ resource.Resource = &StylePolicyActionResource{}

func NewStylePolicyActionResource() resource.Resource {
	return &StylePolicyActionResource{}
}

type StylePolicyActionResource struct {
	pData *tfutils.ProviderData
}

func (r *StylePolicyActionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_style_policy_action"
}

func (r *StylePolicyActionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Define processing actions for the processing rules in a processing policy.", "action", "").String,
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
			"type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Action type", "type", "").AddStringEnum("validate", "filter", "route-action", "aaa", "xformpi", "xformbin", "cryptobin", "xform", "convert-http", "log", "results", "results-async", "setvar", "fetch", "extract", "rewrite", "route-set", "strip-attachments", "call", "on-error", "checkpoint", "slm", "quota-enforcement", "sql", "conditional", "for-each", "event-sink", "method-rewrite", "xformng", "gatewayscript", "jose-sign", "jose-verify", "jose-encrypt", "jose-decrypt", "parse").AddDefaultValue("xform").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("validate", "filter", "route-action", "aaa", "xformpi", "xformbin", "cryptobin", "xform", "convert-http", "log", "results", "results-async", "setvar", "fetch", "extract", "rewrite", "route-set", "strip-attachments", "call", "on-error", "checkpoint", "slm", "quota-enforcement", "sql", "conditional", "for-each", "event-sink", "method-rewrite", "xformng", "gatewayscript", "jose-sign", "jose-verify", "jose-encrypt", "jose-decrypt", "parse"),
				},
				Default: stringdefault.StaticString("xform"),
			},
			"input": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the input context for the action, which identifies the context that contains the document to act. Enter the context name, the string <tt>PIPE</tt> for streaming mode, or the string <tt>INPUT</tt> to identify the original input of the policy rule.", "input", "").String,
				Optional:            true,
			},
			"transform": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the XSL stylesheet or transform file. The location uses one of the following formats. <ul><li>Use a URL, for example, <tt>store:///myTest.xsl</tt></li><li>Use a context variable that expands to a URL, for example, <tt>var://context/contextName/varName</tt></li><li>Use a context, for example, <tt>var://context/Name</tt> or <tt>var://context/Name/</tt> . The context runs as a stylesheet.</li></ul>", "transform", "").String,
				Optional:            true,
			},
			"parse_settings_reference": models.GetDmDynamicParseSettingsReferenceResourceSchema("Specify the configuration that defines the constraints against the documents to parse. Use any or all of the following ways to define the constraints. <ul><li>Specify a URL reference from which to retrieve the constraints definition.</li><li>Specify a literal configuration string in XML or JSON format that contains the constraints definition.</li><li>Specify an instance of the parse settings configuration object to retrieve constraints definition.</li></ul><p>Precedence rules apply when the constraint for the same aspect of an input document is configured with more than one method.</p>", "parse-settings-reference", "", false),
			"parse_metrics_result_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Parse metrics result type", "parse-settings-result-type", "").AddStringEnum("none", "xml", "json", "graphql").AddDefaultValue("none").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("none", "xml", "json", "graphql"),
				},
				Default: stringdefault.StaticString("none"),
			},
			"parse_metrics_result_location": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Parse metrics result location", "parse-settings-result-location", "").String,
				Optional:            true,
			},
			"input_language": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Input language", "input-language", "").AddStringEnum("xml", "dfdl", "xsd", "json").AddDefaultValue("xml").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("xml", "dfdl", "xsd", "json"),
				},
				Default: stringdefault.StaticString("xml"),
			},
			"dfdl_input_root_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the root element in the DFDL model from which to start a parse. This property is only meaningful in the context of a DFDL parse. <p>For the input root name, specify the global xsd:element in the XSD file to use to begin the parsing of binary input. The input root name can be selected from the specified XSD file or specified as a variable. For a variable, it must resolve to a valid namespace URL with the name between braces ({}) as a prefix to the local part. For instance, if in the DFDL Schema, the target namespace is \"http://example.com/messages\" and the local element is &lt;xsd:element name=\"Message\">...&lt;/xsd:element>, the variable must resolve to {http://example.com/messages}Message.</p><p>The schema author might specify the root parse element by using the ibmSchExtn:docRoot=\"true\" element within the schema. For instance, &lt;xsd:element ibmSchExtn:docRoot=\"true\" name=\"Message\">...&lt;/xsd:element>. In this case, the input root name shows in the selection as the element name followed by (@ibmSchExtn:docRoot=\"true\"). Use of another value for the DFDL input root name overrides the value that is specified in the schema.</p>", "input-root-name", "").String,
				Optional:            true,
			},
			"input_descriptor": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the input descriptor as understood according to the input language. <ul><li>When DFDL, the input descriptor must be a URL to a schema file that defines the input. The input descriptor can be a URL to a directory from which you can select a schema file a variable that resolves to a schema at run time.</li><li>When XML, do not specify an input descriptor.</li><li>When XSD, the input descriptor must be a URL to an XML schema.</li><li>When JSON, do not specify an input descriptor.</li></ul>", "input-descriptor", "").String,
				Optional:            true,
			},
			"output_descriptor": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the output descriptor as understood by the output language. When DFDL, the output descriptor must be a URL to a DFDL schema to serialize the output.", "output-descriptor", "").String,
				Optional:            true,
			},
			"transform_language": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Transform language", "transform-language", "").AddStringEnum("none", "xquery").AddDefaultValue("none").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("none", "xquery"),
				},
				Default: stringdefault.StaticString("none"),
			},
			"output_language": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Output language", "output-language", "").AddStringEnum("default", "dfdl").AddDefaultValue("default").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default", "dfdl"),
				},
				Default: stringdefault.StaticString("default"),
			},
			"tx_map": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the Transformation Extender map file. The generated map file is either on the DataPower Gateway or on a remote HTTP or HTTPS server. Use one of the following formats. <ul><li>When the file is local, use &lt;directory>:///&lt;file></li><li>When the file is remote, use HTTP://&lt;path_qualified_file> or HTTPS://&lt;path_qualified_file></li></ul>", "tx-map", "").String,
				Optional:            true,
			},
			"gateway_script_location": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the GatewayScript file. The file location can be specified in one of the following formats. <ul><li>As a URL, where the file is in the <tt>local:</tt> , <tt>store:</tt> , or <tt>temporary:</tt> directory.</li><li>As a context variable that expands to a URL, such as <tt>var://context/contextName/varName</tt> .</li><li>As a context, for example, <tt>var://context/Name</tt> or <tt>var://context/Name/</tt> . The context content runs as GatewayScript.</li></ul>", "gatewayscript-location", "").String,
				Optional:            true,
			},
			"action_debug": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable GatewayScript debug", "debug", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"tx_top_level_map": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the Transformation Extender map in the map file. A map file can contain more than one map. When not specified, the transform uses the first map in the file.", "tx-tlm", "").String,
				Optional:            true,
			},
			"tx_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the Transformation Extender mode to run the map. DPA is the only mode.", "tx-mode", "").AddStringEnum("default", "no-map", "dpa").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default", "no-map", "dpa"),
				},
			},
			"tx_audit_log": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("ITX audit log", "tx-audit-log", "").String,
				Optional:            true,
			},
			"output": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the output context for the action, which identifies the context that receives the document when the action completes. Enter the context name, the string <tt>PIPE</tt> for streaming mode, or the string <tt>OUTPUT</tt> to identify the final output of the policy rule.", "output", "").String,
				Optional:            true,
			},
			"no_transcode_utf8": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the output from the convert action retains the input encoding or uses ISO 8859-1. An encoding is also known as a character set. For illustrative purposes, assume UTF-8 is the input encoding. <ul><li>When enabled and the input encoding is UTF-8, the output is UTF-8.</li><li>When not enabled and the input encoding is UTF-8, the output is ISO 8859-1. This behavior is the default behavior.</li></ul>", "charset-transparency", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"named_in_out_location_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how to locate named inputs and outputs. Use values that are appropriate for your Transformation Extender configuration.", "named-inouts", "").AddStringEnum("default", "explicit", "dynamic", "interop").AddDefaultValue("default").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default", "explicit", "dynamic", "interop"),
				},
				Default: stringdefault.StaticString("default"),
			},
			"named_inputs": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Named inputs", "named-input", "").String,
				NestedObject:        models.DmNamedInOutResourceSchema,
				Optional:            true,
			},
			"named_outputs": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Named outputs", "named-output", "").String,
				NestedObject:        models.DmNamedInOutResourceSchema,
				Optional:            true,
			},
			"destination": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the resource, which might be either the source or the destination. Specify the location as either a URL or as a variable that expands to a URL.", "destination", "").String,
				Optional:            true,
			},
			"schema_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify XML schema for document validation regardless of any <tt>xsi:schemaLocation</tt> attributes contained with the document. Identify the schema with one of the following formats. <ul><li>Use a URL, for example, <tt>store:///valHigh.xsd</tt></li><li>Use a context variable that expands to a URL, for example, <tt>var://context/contextName/varName</tt></li><li>Use a context, for example, <tt>var://context/Name</tt> or <tt>var://context/Name/</tt> . The context runs as a schema validation.</li></ul>", "schema-url", "").String,
				Optional:            true,
			},
			"json_schema_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the JSON schema for JSON document validation. Identify the schema with one of the following formats. <ul><li>Use a URL, for example, <tt>local:///valHigh.jsv</tt></li><li>Use a context variable that expands to a URL, for example, <tt>var://context/contextName/varName</tt></li><li>Use a context, for example, <tt>var://context/Name</tt> or <tt>var://context/Name/</tt> . The context runs as a JSON schema validation.</li></ul>", "json-schema-url", "").String,
				Optional:            true,
			},
			"wsdl_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL of the WSDL file that defines the operations to use during the validate action. The WSDL file can reside on the local system or on the network. By default, the WSDL validation always applies to the entire input message, which can be modified by compile options on the XML manager. <p>Identify the WSDL with one of the following formats.</p><ul><li>Use a URL, for example, <tt>local:///myTest.wsdl</tt></li><li>Use a context variable that expands to a URL, for example, <tt>var://context/contextName/varName</tt></li><li>Use a context, for example, <tt>var://context/Name</tt> or <tt>var://context/Name/</tt> . The context runs as a WSDL validation.</li></ul>", "wsdl-url", "").String,
				Optional:            true,
			},
			"policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("URL rewrite policy", "urlrewrite-policy", "url_rewrite_policy").String,
				Optional:            true,
			},
			"aaa": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("AAA policy", "aaa-policy", "aaa_policy").String,
				Optional:            true,
			},
			"dynamic_schema": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Dynamic schema", "dynamic-schema", "").String,
				Optional:            true,
			},
			"dynamic_stylesheet": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Dynamic stylesheet", "dynamic-stylesheet", "").String,
				Optional:            true,
			},
			"input_conversion": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Input conversion", "input-conversion", "http_input_conversion_map").String,
				Optional:            true,
			},
			"x_path": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the XPath expression to apply to the input context. Enter the XPath expression or a variable in the <tt>var://context/name</tt> format that expands to an XPath expression.", "xpath", "").String,
				Optional:            true,
			},
			"variable": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the variable URL in one of the following formats. <ul><li>var://context/CONTEXT-NAME/VAR-NAME <p>This format is the primary way to reference variables. var://context/CONTEXT-NAME/_roottree is a special variable that holds the value of the context when used as input to an action. var://context/CONTEXT-NAME (or var://context/CONTEXT-NAME/) is treated as shorthand for var://context/CONTEXT-NAME/_roottree.</p></li><li>var://local/VAR-NAME <p>This format can be used to reference variables in the input or output context. Because this reference is context-sensitive, The use of var://context/CONTEXT-NAME/VAR-NAME is recommended.</p></li><li>var://system/CONTEXT-NAME/VAR-NAME <p>This format is used to reference global variables. These variables are rarely used.</p></li><li>var://service/SERVICE-NAME <p>This format is used to reference certain internal state variables. These variables are defined by the firmware.</p></li></ul>", "variable", "").String,
				Optional:            true,
			},
			"value": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the variable value. The value can be a number, a string, or another variable URL.", "value", "").String,
				Optional:            true,
			},
			"ssl_client_config_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client type", "ssl-client-type", "").AddStringEnum("client").AddDefaultValue("client").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("client"),
				},
				Default: stringdefault.StaticString("client"),
			},
			"ssl_client_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "ssl_client_profile").String,
				Optional:            true,
			},
			"attachment_uri": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the attachment to strip. When omitted, all attachments are stripped from the context.", "attachment-uri", "").String,
				Optional:            true,
			},
			"stylesheet_parameters": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Stylesheet parameters", "parameter", "").String,
				NestedObject:        models.DmStylesheetParameterResourceSchema,
				Optional:            true,
			},
			"error_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Error mode", "error-mode", "").AddStringEnum("abort", "continue", "alternative").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("abort", "continue", "alternative"),
				},
			},
			"error_input": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Error input", "error-input", "").String,
				Optional:            true,
			},
			"error_output": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Error output", "error-output", "").String,
				Optional:            true,
			},
			"rule": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Processing rule", "rule", "").String,
				Optional:            true,
			},
			"output_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Output type", "output-type", "").AddStringEnum("default", "binary", "xml").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default", "binary", "xml"),
				},
			},
			"log_level": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Log level", "log-level", "").AddStringEnum("emerg", "alert", "critic", "error", "warn", "notice", "info", "debug").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("emerg", "alert", "critic", "error", "warn", "notice", "info", "debug"),
				},
			},
			"log_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Log type", "log-type", "log_label").String,
				Optional:            true,
			},
			"transactional": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Transactional", "transactional", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"checkpoint_event": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Event", "event", "").AddStringEnum("Request", "Response", "Fault", "AuthComplete").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Request", "Response", "Fault", "AuthComplete"),
				},
			},
			"slm_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SLM policy", "slm", "").String,
				Optional:            true,
			},
			"sql_data_source": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SQL Data Source", "sql-source", "sql_data_source").String,
				Optional:            true,
			},
			"sql_text": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SQL text", "sql-text", "").String,
				Optional:            true,
			},
			"soap_validation": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify which parts of the SOAP message to validate. This setting does not affect the validation of the input context to ensure that it is a valid document.", "soap-validation", "").AddStringEnum("body", "body-or-detail", "ignore-faults", "envelope").AddDefaultValue("body").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("body", "body-or-detail", "ignore-faults", "envelope"),
				},
				Default: stringdefault.StaticString("body"),
			},
			"sql_source_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SQL input method", "sql-source-type", "").AddStringEnum("static", "variable", "stylesheet", "static_internal").AddDefaultValue("static").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("static", "variable", "stylesheet", "static_internal"),
				},
				Default: stringdefault.StaticString("static"),
			},
			"jose_serialization_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Serialization", "serialization", "").AddStringEnum("compact", "json", "json_flat").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("compact", "json", "json_flat"),
				},
			},
			"jwe_enc_algorithm": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Algorithm", "jwe-enc", "").AddStringEnum("A128CBC-HS256", "A192CBC-HS384", "A256CBC-HS512", "A128GCM", "A192GCM", "A256GCM").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("A128CBC-HS256", "A192CBC-HS384", "A256CBC-HS512", "A128GCM", "A192GCM", "A256GCM"),
				},
			},
			"jws_signature_object": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Signature", "jws-signature", "jws_signature").String,
				Optional:            true,
			},
			"jwe_header_object": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("JWE Header", "jwe-header", "jwe_header").String,
				Optional:            true,
			},
			"jose_verify_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Identifier type", "jose-verify-type", "").AddStringEnum("identifiers", "single-cert", "single-sskey").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("identifiers", "single-cert", "single-sskey"),
				},
			},
			"jose_decrypt_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Identifier Type", "jose-decrypt-type", "").AddStringEnum("identifiers", "single-key", "single-sskey", "direct-key").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("identifiers", "single-key", "single-sskey", "direct-key"),
				},
			},
			"signature_identifier": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Signature Identifiers", "signature-identifier", "jose_signature_identifier").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"recipient_identifier": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Recipient Identifiers", "recipient-identifier", "jose_recipient_identifier").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"single_certificate": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Certificate", "single-cert", "crypto_certificate").String,
				Optional:            true,
			},
			"single_key": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Private Key", "single-key", "crypto_key").String,
				Optional:            true,
			},
			"single_ss_key": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Shared Secret Key", "single-sskey", "crypto_sskey").String,
				Optional:            true,
			},
			"jwe_direct_key_object": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Direct Key", "direct-key", "crypto_sskey").String,
				Optional:            true,
			},
			"jws_verify_strip_signature": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Strip Signature", "strip-signature", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"asynchronous": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to mark the action as asynchronous. When asynchronous, the action does not need to complete for the next action to start.", "asynchronous", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"condition": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the conditions to check and action to run when a match is found A single condition maps an XPath expression to search for in the input context to an action to run when the condition is found. When no match is found, other conditions can be checked.", "condition", "").String,
				NestedObject:        models.DmConditionResourceSchema,
				Optional:            true,
			},
			"results_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the processing mode for multiple targets. The default behavior, is first available.", "results-mode", "").AddStringEnum("first-available", "require-all", "attempt-all").AddDefaultValue("first-available").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("first-available", "require-all", "attempt-all"),
				},
				Default: stringdefault.StaticString("first-available"),
			},
			"retry_count": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the number of connection attempts. The default value is 0, which indicates that the operation fails immediately if the connection fails.", "retry-count", "").String,
				Optional:            true,
			},
			"retry_interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration between connection attempts in milliseconds. The default value is 1000.", "retry-interval", "").AddDefaultValue("1000").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(1000),
			},
			"multiple_outputs": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to place parallel outputs into separate contexts. A results action can target several targets simultaneously by specifying a variable that contains an XML nodeset as the results target. <ul><li>When enabled, the context for the output context is ignored. The status of the individual attempts to reach the targets are recorded in separate contexts by appending a number. For example, the resultant context name is <tt>ctx_1</tt> for the <tt>ctx</tt> context name.</li><li>When not enabled, only one result is kept. <ul><li>In require-all mode, keep the first target that failed and has no remaining attempts or the last target to succeed.</li><li>In attempt-all mode, the last target attempted.</li><li>In first-available mode, the first target to succeed or the last to fail.</li></ul></li></ul>", "multiple-outputs", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"iterator_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Iterator type", "iterator-type", "").AddStringEnum("XPATH", "COUNT").AddDefaultValue("XPATH").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("XPATH", "COUNT"),
				},
				Default: stringdefault.StaticString("XPATH"),
			},
			"iterator_expression": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XPath expression", "iterator-expression", "").String,
				Optional:            true,
			},
			"iterator_count": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the number of times to run the loop action. Enter a value in the range 1 - 32768.", "iterator-count", "").AddIntegerRange(1, 32768).String,
				Optional:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 32768),
				},
			},
			"loop_action": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Loop action", "loop-action", "").String,
				Optional:            true,
			},
			"async_action": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Asynchronous actions", "async-action", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration to wait for asynchronous actions to complete in milliseconds. After this period, an error is logged. A value of 0 indicates to wait forever.", "timeout", "").String,
				Optional:            true,
			},
			"wsdl_port_q_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the QName of the WSDL port that defines the service traffic to validate. The value must be a QName in the <tt>{namespace-uri}local-part</tt> format or <tt>*</tt> for all ports. When specified and not <tt>*</tt> , only messages that are defined for the named port are considered valid.", "wsdl-port", "").String,
				Optional:            true,
			},
			"wsdl_operation_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the WSDL operation that defines the service traffic to validate. The value must be the unqualified name of the operation or <tt>*</tt> for all operations. When specified and not <tt>*</tt> , only messages that are defined for operations that match the specified name are considered valid.", "wsdl-operation", "").String,
				Optional:            true,
			},
			"wsdl_message_direction_or_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name or direction of the WSDL input, output, or fault that defines the service traffic to validate. Use one of the following values. <ul><li>The name of one or more WSDL input, output, or fault components.</li><ul><li><tt><i>operation</i> Request</tt> for a specific request.</li><li><tt><i>operation</i> Response\"</tt> for a specific response.</li><li>The value of the <tt>name</tt> attribute for the <tt>&lt;fault></tt> element.</li></ul><li><tt>#input</tt> for the request direction in the WSDL file.</li><li><tt>#output</tt> for the response direction in the WSDL file.</li><li><tt>*</tt> for all inputs, outputs, and faults in the WSDL file.</li></ul><p>When specified and not <tt>*</tt> , only messages that are defined for inputs, outputs, and faults that match the specified name or direction are considered valid. Faults are considered valid for the response direction.</p>", "wsdl-message-direction-or-name", "").String,
				Optional:            true,
			},
			"wsdl_attachment_part": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the WSDL message part that defines the content of a MIME attachment (mime:content/@part) to validate. The value must be the unqualified name of the message part. The name is the same as the part attribute on the corresponding <tt>mime:content</tt> component in the WSDL file. When not defined or <tt>*</tt> , the root MIME part is validated. The root MIME part is bound to a <tt>soap:body</tt> .", "wsdl-attachment-part", "").String,
				Optional:            true,
			},
			"method_rewrite_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Method", "http-method", "").AddStringEnum("POST", "GET", "PUT", "PATCH", "DELETE", "HEAD").AddDefaultValue("GET").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("POST", "GET", "PUT", "PATCH", "DELETE", "HEAD"),
				},
				Default: stringdefault.StaticString("GET"),
			},
			"method_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Method", "http-method-limited", "").AddStringEnum("POST", "GET", "PUT", "PATCH", "DELETE", "HEAD").AddDefaultValue("POST").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("POST", "GET", "PUT", "PATCH", "DELETE", "HEAD"),
				},
				Default: stringdefault.StaticString("POST"),
			},
			"method_type2": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Method", "http-method-limited2", "").AddStringEnum("POST", "GET", "PUT", "PATCH", "DELETE", "HEAD").AddDefaultValue("POST").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("POST", "GET", "PUT", "PATCH", "DELETE", "HEAD"),
				},
				Default: stringdefault.StaticString("POST"),
			},
			"policy_key": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<b>Do not modify this value.</b> The DataPower Gateway uses this identifier to store the output from the action. The output can be used for external monitoring.", "policy-key", "").String,
				Optional:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *StylePolicyActionResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *StylePolicyActionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.StylePolicyAction
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `StylePolicyAction`)
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

func (r *StylePolicyActionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.StylePolicyAction
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
		data.FromBody(ctx, `StylePolicyAction`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `StylePolicyAction`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StylePolicyActionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.StylePolicyAction
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `StylePolicyAction`))
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

func (r *StylePolicyActionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.StylePolicyAction
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

func (r *StylePolicyActionResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.StylePolicyAction

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
