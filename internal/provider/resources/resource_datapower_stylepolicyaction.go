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
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &StylePolicyActionResource{}

func NewStylePolicyActionResource() resource.Resource {
	return &StylePolicyActionResource{}
}

type StylePolicyActionResource struct {
	client *client.DatapowerClient
}

func (r *StylePolicyActionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_stylepolicyaction"
}

func (r *StylePolicyActionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Processing action", "action", "").String,

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
				MarkdownDescription: tfutils.NewAttributeDescription("Input", "input", "").String,
				Optional:            true,
			},
			"transform": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Transform file", "transform", "").String,
				Optional:            true,
			},
			"parse_settings_reference": models.GetDmDynamicParseSettingsReferenceResourceSchema("Parse settings", "parse-settings-reference", "", false),
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
				MarkdownDescription: tfutils.NewAttributeDescription("DFDL input root name", "input-root-name", "").String,
				Optional:            true,
			},
			"input_descriptor": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Input descriptor", "input-descriptor", "").String,
				Optional:            true,
			},
			"output_descriptor": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Output descriptor", "output-descriptor", "").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("ITX map file", "tx-map", "").String,
				Optional:            true,
			},
			"gateway_script_location": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("GatewayScript file", "gatewayscript-location", "").String,
				Optional:            true,
			},
			"action_debug": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable GatewayScript debug", "debug", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"tx_top_level_map": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("ITX top-level map name", "tx-tlm", "").String,
				Optional:            true,
			},
			"tx_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("ITX map mode", "tx-mode", "").AddStringEnum("default", "no-map", "dpa").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Output", "output", "").String,
				Optional:            true,
			},
			"no_transcode_utf8": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Retain input encoding", "charset-transparency", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"named_in_out_location_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Locate named inputs and outputs", "named-inouts", "").AddStringEnum("default", "explicit", "dynamic", "interop").AddDefaultValue("default").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("External URL", "destination", "").String,
				Optional:            true,
			},
			"schema_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Schema URL", "schema-url", "").String,
				Optional:            true,
			},
			"json_schema_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("JSON schema URL", "json-schema-url", "").String,
				Optional:            true,
			},
			"wsdl_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WSDL URL", "wsdl-url", "").String,
				Optional:            true,
			},
			"policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("URL rewrite policy", "urlrewrite-policy", "urlrewritepolicy").String,
				Optional:            true,
			},
			"aaa": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("AAA policy", "aaa-policy", "aaapolicy").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Input conversion", "input-conversion", "httpinputconversionmap").String,
				Optional:            true,
			},
			"x_path": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XPath", "xpath", "").String,
				Optional:            true,
			},
			"variable": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Variable name", "variable", "").String,
				Optional:            true,
			},
			"value": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Variable value", "value", "").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "sslclientprofile").String,
				Optional:            true,
			},
			"attachment_uri": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Attachment URI", "attachment-uri", "").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Log type", "log-type", "loglabel").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("SQL Data Source", "sql-source", "sqldatasource").String,
				Optional:            true,
			},
			"sql_text": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SQL text", "sql-text", "").String,
				Optional:            true,
			},
			"soap_validation": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SOAP validation", "soap-validation", "").AddStringEnum("body", "body-or-detail", "ignore-faults", "envelope").AddDefaultValue("body").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Signature", "jws-signature", "jwssignature").String,
				Optional:            true,
			},
			"jwe_header_object": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("JWE Header", "jwe-header", "jweheader").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Signature Identifiers", "signature-identifier", "josesignatureidentifier").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"recipient_identifier": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Recipient Identifiers", "recipient-identifier", "joserecipientidentifier").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"single_certificate": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Certificate", "single-cert", "cryptocertificate").String,
				Optional:            true,
			},
			"single_key": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Private Key", "single-key", "cryptokey").String,
				Optional:            true,
			},
			"single_ss_key": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Shared Secret Key", "single-sskey", "cryptosskey").String,
				Optional:            true,
			},
			"jwe_direct_key_object": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Direct Key", "direct-key", "cryptosskey").String,
				Optional:            true,
			},
			"jws_verify_strip_signature": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Strip Signature", "strip-signature", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"asynchronous": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Asynchronous", "asynchronous", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"condition": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Condition", "condition", "").String,
				NestedObject:        models.DmConditionResourceSchema,
				Optional:            true,
			},
			"results_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Multi-way results mode", "results-mode", "").AddStringEnum("first-available", "require-all", "attempt-all").AddDefaultValue("first-available").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("first-available", "require-all", "attempt-all"),
				},
				Default: stringdefault.StaticString("first-available"),
			},
			"retry_count": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Number of retries", "retry-count", "").String,
				Optional:            true,
			},
			"retry_interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Retry interval", "retry-interval", "").AddDefaultValue("1000").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(1000),
			},
			"multiple_outputs": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use multiple outputs", "multiple-outputs", "").AddDefaultValue("false").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Iterator count", "iterator-count", "").AddIntegerRange(1, 32768).String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Timeout", "timeout", "").String,
				Optional:            true,
			},
			"wsdl_port_q_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WSDL Port QName", "wsdl-port", "").String,
				Optional:            true,
			},
			"wsdl_operation_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WSDL Operation Name", "wsdl-operation", "").String,
				Optional:            true,
			},
			"wsdl_message_direction_or_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WSDL Message Direction or Name", "wsdl-message-direction-or-name", "").String,
				Optional:            true,
			},
			"wsdl_attachment_part": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WSDL Attachment Part Name", "wsdl-attachment-part", "").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Policy identifier (for internal use)", "policy-key", "").String,
				Optional:            true,
			},
		},
	}
}

func (r *StylePolicyActionResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *StylePolicyActionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.StylePolicyAction

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `StylePolicyAction`)
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

func (r *StylePolicyActionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.StylePolicyAction

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
		data.FromBody(ctx, `StylePolicyAction`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `StylePolicyAction`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StylePolicyActionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.StylePolicyAction

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `StylePolicyAction`))
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

func (r *StylePolicyActionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.StylePolicyAction

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
