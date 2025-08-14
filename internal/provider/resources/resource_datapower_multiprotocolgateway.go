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

var _ resource.Resource = &MultiProtocolGatewayResource{}
var _ resource.ResourceWithValidateConfig = &MultiProtocolGatewayResource{}

func NewMultiProtocolGatewayResource() resource.Resource {
	return &MultiProtocolGatewayResource{}
}

type MultiProtocolGatewayResource struct {
	pData *tfutils.ProviderData
}

func (r *MultiProtocolGatewayResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_multiprotocolgateway"
}

func (r *MultiProtocolGatewayResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Multi-Protocol Gateway", "mpgw", "").AddActions("quiesce").String,
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
			"back_http_version": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP version to server", "http-server-version", "").AddStringEnum("HTTP/1.0", "HTTP/1.1", "HTTP/2").AddDefaultValue("HTTP/1.1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("HTTP/1.0", "HTTP/1.1", "HTTP/2"),
				},
				Default: stringdefault.StaticString("HTTP/1.1"),
			},
			"http2_required": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP/2 required", "http2-required", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"request_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Request type", "request-type", "").AddStringEnum("soap", "xml", "unprocessed", "preprocessed", "json").AddDefaultValue("soap").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("soap", "xml", "unprocessed", "preprocessed", "json"),
				},
				Default: stringdefault.StaticString("soap"),
			},
			"response_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Response type", "response-type", "").AddStringEnum("soap", "xml", "unprocessed", "preprocessed", "json").AddDefaultValue("soap").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("soap", "xml", "unprocessed", "preprocessed", "json"),
				},
				Default: stringdefault.StaticString("soap"),
			},
			"follow_redirects": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Follow redirects", "follow-redirects", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"rewrite_location_header": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Rewrite Location URL", "rewrite-location-header", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"style_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Processing policy", "stylepolicy", "stylepolicy").AddDefaultValue("default").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default"),
			},
			"type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Type", "type", "").AddStringEnum("static-backend", "dynamic-backend", "static-from-wsdl").AddDefaultValue("static-backend").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("static-backend", "dynamic-backend", "static-from-wsdl"),
				},
				Default: stringdefault.StaticString("static-backend"),
			},
			"allow_compression": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Compression", "compression", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"enable_compressed_payload_passthrough": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable pass through of compressed payload", "enable-compressed-payload-passthrough", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"allow_cache_control_header": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow Cache-Control header", "allow-cache-control", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wsrr_saved_search_subscriptions": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WSRR saved search subscriptions", "wsrr-saved-search-subscription", "").String,
				NestedObject:        models.DmWSRRSavedSearchWSDLSourceResourceSchema,
				Optional:            true,
			},
			"wsrr_subscriptions": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WSRR subscriptions", "wsrr-subscription", "").String,
				NestedObject:        models.DmWSRRWSDLSourceResourceSchema,
				Optional:            true,
			},
			"policy_attachments": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Policy attachments", "policy-attachments", "policyattachments").String,
				Optional:            true,
			},
			"policy_parameter": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Policy parameters", "policy-parameters", "").String,
				NestedObject:        models.DmWSPolicyParametersResourceSchema,
				Optional:            true,
			},
			"wsm_agent_monitor": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Monitor with Web Services Management agent", "wsmagent-monitor", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wsm_agent_monitor_pcm": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Message capture with Web Services Management agent", "wsmagent-monitor-capture-mode", "").AddStringEnum("none", "all-messages", "on-error").AddDefaultValue("all-messages").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("none", "all-messages", "on-error"),
				},
				Default: stringdefault.StaticString("all-messages"),
			},
			"proxy_http_response": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Proxy HTTP response", "proxy-http-response", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"error_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Error policy", "error-policy", "mpgwerrorhandlingpolicy").String,
				Optional:            true,
			},
			"transaction_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Transaction timeout", "transaction-timeout", "").String,
				Optional:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"priority": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Service Priority", "priority", "").AddStringEnum("unknown", "high-min", "high", "high-max", "normal-min", "normal", "normal-max", "low-min", "low", "low-max").AddDefaultValue("normal").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("unknown", "high-min", "high", "high-max", "normal-min", "normal", "normal-max", "low-min", "low", "low-max"),
				},
				Default: stringdefault.StaticString("normal"),
			},
			"front_protocol": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Front Side Protocol", "front-protocol", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"xml_manager": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML Manager", "xml-manager", "xmlmanager").AddDefaultValue("default").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default"),
			},
			"url_rewrite_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("URL Rewrite Policy", "urlrewrite-policy", "urlrewritepolicy").String,
				Optional:            true,
			},
			"ssl_client_config_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client type", "ssl-client-type", "").AddStringEnum("proxy", "client").AddDefaultValue("client").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("proxy", "client"),
				},
				Default: stringdefault.StaticString("client"),
			},
			"ssl_client": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "sslclientprofile").String,
				Optional:            true,
			},
			"fw_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Crypto Credentials", "fwcred", "cryptofwcred").String,
				Optional:            true,
			},
			"header_injection": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP Header Injection", "inject", "").String,
				NestedObject:        models.DmHeaderInjectionResourceSchema,
				Optional:            true,
			},
			"header_suppression": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP Header Suppression", "suppress", "").String,
				NestedObject:        models.DmHeaderSuppressionResourceSchema,
				Optional:            true,
			},
			"stylesheet_parameters": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Stylesheet Parameter", "parameter", "").String,
				NestedObject:        models.DmStylesheetParameterResourceSchema,
				Optional:            true,
			},
			"default_param_namespace": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Default parameter namespace", "default-param-namespace", "").AddDefaultValue("http://www.datapower.com/param/config").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("http://www.datapower.com/param/config"),
			},
			"query_param_namespace": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Query parameter namespace", "query-param-namespace", "").AddDefaultValue("http://www.datapower.com/param/query").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("http://www.datapower.com/param/query"),
			},
			"backend_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Backend URL", "backend-url", "").String,
				Optional:            true,
			},
			"propagate_uri": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Propagate URI", "propagate-uri", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"service_monitors": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Service Monitors", "monitor-service", "webservicemonitor").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"count_monitors": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Count Monitors", "monitor-count", "countmonitor").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"duration_monitors": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Duration Monitors", "monitor-duration", "durationmonitor").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"monitor_processing_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Monitors evaluation method", "monitor-processing-policy", "").AddStringEnum("terminate-at-first-throttle", "terminate-at-first-match").AddDefaultValue("terminate-at-first-throttle").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("terminate-at-first-throttle", "terminate-at-first-match"),
				},
				Default: stringdefault.StaticString("terminate-at-first-throttle"),
			},
			"request_attachments": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Request attachment processing mode", "request-attachments", "").AddStringEnum("strip", "reject", "allow", "streaming", "unprocessed").AddDefaultValue("strip").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("strip", "reject", "allow", "streaming", "unprocessed"),
				},
				Default: stringdefault.StaticString("strip"),
			},
			"response_attachments": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Response attachment processing mode", "response-attachments", "").AddStringEnum("strip", "reject", "allow", "streaming", "unprocessed").AddDefaultValue("strip").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("strip", "reject", "allow", "streaming", "unprocessed"),
				},
				Default: stringdefault.StaticString("strip"),
			},
			"request_attachments_flow_control": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Request attachment flow control mode", "request-attachments-flow-control", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"response_attachments_flow_control": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Response attachment flow control mode", "response-attachments-flow-control", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"root_part_not_first_action": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Action when required root part is not first", "root-part-not-first-action", "").AddStringEnum("process-in-order", "buffer", "abort").AddDefaultValue("process-in-order").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("process-in-order", "buffer", "abort"),
				},
				Default: stringdefault.StaticString("process-in-order"),
			},
			"front_attachment_format": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Front attachment processing format", "front-attachment-format", "").AddStringEnum("dynamic", "mime", "dime", "detect").AddDefaultValue("dynamic").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("dynamic", "mime", "dime", "detect"),
				},
				Default: stringdefault.StaticString("dynamic"),
			},
			"back_attachment_format": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Back attachment processing format", "back-attachment-format", "").AddStringEnum("dynamic", "mime", "dime", "detect").AddDefaultValue("dynamic").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("dynamic", "mime", "dime", "detect"),
				},
				Default: stringdefault.StaticString("dynamic"),
			},
			"mime_front_headers": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Front Side MIME Header Processing", "mime-front-headers", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"mime_back_headers": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Back Side MIME Header Processing", "mime-back-headers", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"stream_output_to_back": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Stream Output to Back", "stream-output-to-back", "").AddStringEnum("buffer-until-verification", "stream-until-infraction").AddDefaultValue("buffer-until-verification").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("buffer-until-verification", "stream-until-infraction"),
				},
				Default: stringdefault.StaticString("buffer-until-verification"),
			},
			"stream_output_to_front": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Stream Output to Front", "stream-output-to-front", "").AddStringEnum("buffer-until-verification", "stream-until-infraction").AddDefaultValue("buffer-until-verification").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("buffer-until-verification", "stream-until-infraction"),
				},
				Default: stringdefault.StaticString("buffer-until-verification"),
			},
			"max_message_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Maximum Message Size", "max-message-size", "").AddIntegerRange(0, 2097151).String,
				Optional:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 2097151),
				},
			},
			"gateway_parser_limits": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Parser limits", "gateway-parser-limits", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"parser_limits_element_depth": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML Element Depth", "element-depth", "").AddDefaultValue("512").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(512),
			},
			"parser_limits_attribute_count": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML Attribute Count", "attribute-count", "").AddDefaultValue("128").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(128),
			},
			"parser_limits_max_node_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML Maximum Node Size", "max-node-size", "").AddIntegerRange(1024, 4294967295).AddDefaultValue("33554432").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1024, 4294967295),
				},
				Default: int64default.StaticInt64(33554432),
			},
			"parser_limits_external_references": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML External Reference Handling", "external-references", "").AddStringEnum("forbid", "ignore", "allow").AddDefaultValue("forbid").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("forbid", "ignore", "allow"),
				},
				Default: stringdefault.StaticString("forbid"),
			},
			"parser_limits_max_prefixes": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML Maximum Distinct Prefixes", "max-prefixes", "").AddDefaultValue("1024").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(1024),
			},
			"parser_limits_max_namespaces": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML Maximum Distinct Namespaces", "max-namespaces", "").AddDefaultValue("1024").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(1024),
			},
			"parser_limits_max_local_names": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML Maximum Distinct Local Names", "max-local-names", "").AddDefaultValue("60000").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(60000),
			},
			"parser_limits_attachment_byte_count": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Attachment Byte Count Limit", "attachment-byte-count", "").AddDefaultValue("2000000000").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(2000000000),
			},
			"parser_limits_attachment_package_byte_count": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Attachment Package Byte Count Limit", "attachment-package-byte-count", "").String,
				Optional:            true,
			},
			"debug_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Probe setting", "debug-mode", "").AddStringEnum("on", "off", "unbounded").AddDefaultValue("off").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("on", "off", "unbounded"),
				},
				Default: stringdefault.StaticString("off"),
			},
			"debug_history": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Transaction History", "debug-history", "").AddIntegerRange(10, 250).AddDefaultValue("25").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(10, 250),
				},
				Default: int64default.StaticInt64(25),
			},
			"debug_trigger": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Probe Triggers", "debug-trigger", "").String,
				NestedObject:        models.DmMSDebugTriggerTypeResourceSchema,
				Optional:            true,
			},
			"flow_control": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Flow Control", "flowcontrol", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"soap_schema_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SOAP Schema URL", "soap-schema-url", "").AddDefaultValue("store:///schemas/soap-envelope.xsd").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("store:///schemas/soap-envelope.xsd"),
			},
			"front_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Front Side Timeout", "front-timeout", "").AddIntegerRange(1, 86400).AddDefaultValue("120").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(120),
			},
			"back_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Back Side Timeout", "back-timeout", "").AddIntegerRange(1, 86400).AddDefaultValue("120").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(120),
			},
			"front_persistent_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Front Persistent Timeout", "front-persistent-timeout", "").AddIntegerRange(0, 86400).AddDefaultValue("180").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 86400),
				},
				Default: int64default.StaticInt64(180),
			},
			"back_persistent_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Back Persistent Timeout", "back-persistent-timeout", "").AddIntegerRange(0, 86400).AddDefaultValue("180").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 86400),
				},
				Default: int64default.StaticInt64(180),
			},
			"include_response_type_encoding": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Include charset in response-type", "include-content-type-encoding", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"persistent_connections": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Persistent Connections", "persistent-connections", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"loop_detection": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Loop Detection", "loop-detection", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"do_host_rewriting": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Rewrite Host Names when Redirecting", "host-rewriting", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"do_chunked_upload": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow Chunked Uploads", "chunked-uploads", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"process_http_errors": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Process Backend Errors", "process-http-errors", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"http_client_ip_label": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP Client IP Label", "http-client-ip-label", "").AddDefaultValue("X-Client-IP").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("X-Client-IP"),
			},
			"http_log_cor_id_label": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP Global Transaction ID Label", "http-global-tranID-label", "").AddDefaultValue("X-Global-Transaction-ID").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("X-Global-Transaction-ID"),
			},
			"load_balancer_hash_header": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Load Balancer Hash Header", "load-balancer-hash-header", "").String,
				Optional:            true,
			},
			"in_order_mode": models.GetDmGatewayInOrderModeResourceSchema("Message Processing Modes", "inorder-mode", "", false),
			"wsa_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WS-Addressing Mode", "wsa-mode", "").AddStringEnum("wsa2sync", "sync2wsa", "wsa2wsa", "sync2sync").AddDefaultValue("sync2sync").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("wsa2sync", "sync2wsa", "wsa2wsa", "sync2sync"),
				},
				Default: stringdefault.StaticString("sync2sync"),
			},
			"wsa_require_aaa": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Require AAA for WS-Addressing trust", "wsa-require-aaa", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"wsa_rewrite_reply_to": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Rewrite WS Addressing Reply To Header", "wsa-replyto-rewrite", "urlrewritepolicy").String,
				Optional:            true,
			},
			"wsa_rewrite_fault_to": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Rewrite WS Addressing Fault To Header", "wsa-faultto-rewrite", "urlrewritepolicy").String,
				Optional:            true,
			},
			"wsa_rewrite_to": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Rewrite WS Addressing To Header", "wsa-to-rewrite", "urlrewritepolicy").String,
				Optional:            true,
			},
			"wsa_strip": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Strip WS-Addressing Headers", "wsa-strip-headers", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"wsa_default_reply_to": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Default WS-Addressing Reply-To", "wsa-default-replyto", "").AddDefaultValue("http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous"),
			},
			"wsa_default_fault_to": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Default WS-Addressing Fault-To", "wsa-default-faultto", "").AddDefaultValue("http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous"),
			},
			"wsa_force": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Force Insertion of WS-Addressing Headers", "wsa-force", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wsa_gen_style": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WS-Addressing Message Generation Pattern", "wsa-genstyle", "").AddStringEnum("sync", "async", "oob").AddDefaultValue("sync").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("sync", "async", "oob"),
				},
				Default: stringdefault.StaticString("sync"),
			},
			"wsahttp_async_response_code": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Asynchronous HTTP ResponseCode", "wsa-http-async-response-code", "").AddIntegerRange(200, 599).AddDefaultValue("204").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(200, 599),
				},
				Default: int64default.StaticInt64(204),
			},
			"wsa_back_protocol": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WS-Addressing Reply Point", "wsa-back-protocol", "").String,
				Optional:            true,
			},
			"wsa_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WS-Addressing Asynchronous Reply Timeout", "wsa-timeout", "").AddIntegerRange(1, 4000000).AddDefaultValue("120").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 4000000),
				},
				Default: int64default.StaticInt64(120),
			},
			"wsrm_enabled": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WS-ReliableMessaging", "wsrm", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wsrm_sequence_expiration": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Target Sequence Expiration Interval", "wsrm-sequence-expiration", "").AddIntegerRange(10, 86400).AddDefaultValue("3600").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(10, 86400),
				},
				Default: int64default.StaticInt64(3600),
			},
			"wsrmaaa_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WS-RM AAA Policy", "wsrm-aaa-policy", "aaapolicy").String,
				Optional:            true,
			},
			"wsrm_destination_accept_create_sequence": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Destination Accept Incoming CreateSequence", "wsrm-destination-accept-create-sequence", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"wsrm_destination_maximum_sequences": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Destination Maximum Simultaneous Sequences", "wsrm-destination-maximum-sequences", "").AddIntegerRange(1, 2048).AddDefaultValue("400").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 2048),
				},
				Default: int64default.StaticInt64(400),
			},
			"wsrm_destination_in_order": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Destination InOrder Delivery Assurance", "wsrm-destination-inorder", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wsrm_destination_maximum_in_order_queue_length": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Destination Maximum InOrder Queue Length", "wsrm-destination-maximum-inorder-queue-length", "").AddIntegerRange(1, 256).AddDefaultValue("10").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 256),
				},
				Default: int64default.StaticInt64(10),
			},
			"wsrm_destination_accept_offers": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Destination Accept Two-Way Offers", "wsrm-destination-accept-offers", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wsrm_front_force": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WS-RM Required on Request", "wsrm-request-force", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wsrm_back_force": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WS-RM Required on Response", "wsrm-response-force", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wsrm_back_create_sequence": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Source Create Sequence on Request", "wsrm-source-request-create-sequence", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wsrm_front_create_sequence": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Source Create Sequence on Response", "wsrm-source-response-create-sequence", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wsrm_source_make_offer": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Source Make Two-Way Offer", "wsrm-source-make-offer", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wsrm_uses_sequence_ssl": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Source use SSL/TLS Session Binding", "wsrm-source-sequence-ssl", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wsrm_front_acks_to": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Source Front Reply Point", "wsrm-source-front-acks-to", "").String,
				Optional:            true,
			},
			"wsrm_back_acks_to": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Source Back AcksTo Reply Point", "wsrm-source-back-acks-to", "").String,
				Optional:            true,
			},
			"wsrm_source_maximum_sequences": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Source Maximum Simultaneous Sequences", "wsrm-source-maximum-sequences", "").AddIntegerRange(1, 2048).AddDefaultValue("400").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 2048),
				},
				Default: int64default.StaticInt64(400),
			},
			"wsrm_source_retransmission_interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Source Retransmission Interval", "wsrm-source-retransmission-interval", "").AddIntegerRange(2, 600).AddDefaultValue("10").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(2, 600),
				},
				Default: int64default.StaticInt64(10),
			},
			"wsrm_source_exponential_backoff": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Source Exponential Backoff", "wsrm-source-exponential-backoff", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"wsrm_source_maximum_retransmissions": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Source Maximum Retransmissions", "wsrm-source-retransmit-count", "").AddIntegerRange(1, 256).AddDefaultValue("4").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 256),
				},
				Default: int64default.StaticInt64(4),
			},
			"wsrm_source_maximum_queue_length": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Source Maximum Queue Length", "wsrm-source-maximum-queue-length", "").AddIntegerRange(1, 256).AddDefaultValue("30").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 256),
				},
				Default: int64default.StaticInt64(30),
			},
			"wsrm_source_request_ack_count": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Source Request Ack Count", "wsrm-source-request-ack-count", "").AddIntegerRange(1, 256).AddDefaultValue("1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 256),
				},
				Default: int64default.StaticInt64(1),
			},
			"wsrm_source_inactivity_close": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Source Inactivity Close Interval", "wsrm-source-inactivity-close-interval", "").AddIntegerRange(10, 3600).AddDefaultValue("360").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(10, 3600),
				},
				Default: int64default.StaticInt64(360),
			},
			"force_policy_exec": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Process Messages Whose Body Is Empty", "force-policy-exec", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"rewrite_errors": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Rewrite Error Messages", "rewrite-errors", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"delay_errors": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Delay Error Messages", "delay-errors", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"delay_errors_duration": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Duration to Delay Error Messages", "delay-errors-duration", "").AddIntegerRange(250, 300000).AddDefaultValue("1000").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(250, 300000),
				},
				Default: int64default.StaticInt64(1000),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *MultiProtocolGatewayResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *MultiProtocolGatewayResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.MultiProtocolGateway
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `MultiProtocolGateway`)
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

func (r *MultiProtocolGatewayResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.MultiProtocolGateway
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
		data.FromBody(ctx, `MultiProtocolGateway`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `MultiProtocolGateway`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MultiProtocolGatewayResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.MultiProtocolGateway
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `MultiProtocolGateway`))
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

func (r *MultiProtocolGatewayResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.MultiProtocolGateway
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

func (r *MultiProtocolGatewayResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.MultiProtocolGateway

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
