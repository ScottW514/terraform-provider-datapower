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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

var _ resource.Resource = &XMLFirewallServiceResource{}
var _ resource.ResourceWithValidateConfig = &XMLFirewallServiceResource{}
var _ resource.ResourceWithImportState = &XMLFirewallServiceResource{}

func NewXMLFirewallServiceResource() resource.Resource {
	return &XMLFirewallServiceResource{}
}

type XMLFirewallServiceResource struct {
	pData *tfutils.ProviderData
}

func (r *XMLFirewallServiceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_xml_firewall_service"
}

func (r *XMLFirewallServiceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Create or edit an XML Firewall on local IP/port. This XML Firewall can communicate with a dynamically identified servers, a static back end server or as a loopback. The XML Firewall applies the selected processing policy to messages. The XML Firewall can rewrite client request URLs byusing a URL rewrite policy. The XML Firewall can use TLS communications to client, server or both directions if applicable.", "xmlfirewall", "").AddActions("quiesce").String,
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
			"type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the mode of the XML Firewall. The default is Dynamic Backend.", "type", "").AddStringEnum("static-backend", "loopback-proxy", "dynamic-backend").AddDefaultValue("dynamic-backend").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("static-backend", "loopback-proxy", "dynamic-backend"),
				},
				Default: stringdefault.StaticString("dynamic-backend"),
			},
			"xml_manager": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("An XML Manager manages the compilation and caching of stylesheets, the caching of documents, and provides configuration constraints on the size and parsing depth of documents. You can enable streaming operation by configuring an XML Manager to use a Streaming Compile Option Policy. More than one firewall may use the same XML Manager. Select an existing XML Manager from the list to assign the Manager to this firewall. Click the + button to create a new XML Manager that is assigned to this firewall. A default Manager is used if you do not create one.", "xml-manager", "xml_manager").AddDefaultValue("default").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("default"),
			},
			"url_rewrite_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("A URL Rewrite Policy applies the rules established in the named policy to rewrite the URL used by the client to connect to the service this firewall represents. This helps to mask the URL and location of the actual service. Select an existing URL rewrite policy from the list of available policies. Click the + button to create a new URL rewrite policy that is then assigned to this firewall.", "urlrewrite-policy", "url_rewrite_policy").String,
				Optional:            true,
			},
			"style_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Assign the processing policy to the service. The processing policy defines the actions to take against messages that flow through the service.", "stylesheet-policy", "style_policy").AddDefaultValue("default").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("default"),
			},
			"max_message_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specifies the maximum allowable size in kilobytes of a message. Enter a value in the range 0 - 2097151. The default value is 0. A value of 0 specifies that no limit is enforced, and the message can be of unlimited size.</p><p>The maximum message size limit applies to JSON, SOAP, XML, and non-XML messages. If the message type is pass-through, no limit is enforced.</p>", "max-message-size", "").AddIntegerRange(0, 2097151).AddDefaultValue("0").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 2097151),
				},
				Default: int64default.StaticInt64(0),
			},
			"request_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Characterizes the type of traffic that originates from the client. The default is SOAP.", "request-type", "").AddStringEnum("soap", "xml", "unprocessed", "preprocessed", "json").AddDefaultValue("soap").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("soap", "xml", "unprocessed", "preprocessed", "json"),
				},
				Default: stringdefault.StaticString("soap"),
			},
			"response_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Characterizes the type of traffic that originates from the server. The default is SOAP.", "response-type", "").AddStringEnum("soap", "xml", "unprocessed", "preprocessed", "json").AddDefaultValue("soap").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("soap", "xml", "unprocessed", "preprocessed", "json"),
				},
				Default: stringdefault.StaticString("soap"),
			},
			"fw_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies an optional Firewall Credentials list. A Firewall Credentials list specifies which keys and certificates are available to support firewall processing. In the absence of an identified Firewall Credentials list, all locally-stored key and certificates are available to the firewall.", "fwcred", "crypto_fw_cred").String,
				Optional:            true,
			},
			"service_monitors": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies a list of Web Service Monitors. Web Service Monitors target Web Service endpoints. Use the Web Services wizard on the Control Panel to supply a WSDL and configure a Monitor for an endpoint. The Monitor collects statistics, establishes count and duration monitors and can take action when thresholds are met or exceeded. A Monitor must be selected here to be activated. Click + to create a new Web Services Monitor. Note that this Monitor is not the Service Level Monitor (SLM) used by a Multi-Protocol Gateway or Web Service Proxy.", "monitor-service", "web_service_monitor").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"request_attachments": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select how to treat client requests with attachments. The default is Strip.", "request-attachments", "").AddStringEnum("strip", "reject", "allow", "streaming", "unprocessed").AddDefaultValue("strip").AddNotValidWhen(models.XMLFirewallServiceRequestAttachmentsIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("strip", "reject", "allow", "streaming", "unprocessed"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.XMLFirewallServiceRequestAttachmentsIgnoreVal, true),
				},
				Default: stringdefault.StaticString("strip"),
			},
			"response_attachments": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select how to treat server responses with attachments. The default is Strip.", "response-attachments", "").AddStringEnum("strip", "reject", "allow", "streaming", "unprocessed").AddDefaultValue("strip").AddNotValidWhen(models.XMLFirewallServiceResponseAttachmentsIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("strip", "reject", "allow", "streaming", "unprocessed"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.XMLFirewallServiceResponseAttachmentsIgnoreVal, true),
				},
				Default: stringdefault.StaticString("strip"),
			},
			"root_part_not_first_action": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("When streaming MIME messages, the action to take when the root part is not the first part of the message. If the root part must be first (for example to do BSP conformance checking) and the action is set to \"process in order\" the DP device will buffer attachments up to the root.", "root-part-not-first-action", "").AddStringEnum("process-in-order", "buffer", "abort").AddDefaultValue("process-in-order").AddNotValidWhen(models.XMLFirewallServiceRootPartNotFirstActionIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("process-in-order", "buffer", "abort"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.XMLFirewallServiceRootPartNotFirstActionIgnoreVal, true),
				},
				Default: stringdefault.StaticString("process-in-order"),
			},
			"front_attachment_format": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select how to interpret client requests with attachments. If an attachment arrives that is not of the type chosen, the attachment will be rejected.", "front-attachment-format", "").AddStringEnum("dynamic", "mime", "dime").AddDefaultValue("dynamic").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("dynamic", "mime", "dime"),
				},
				Default: stringdefault.StaticString("dynamic"),
			},
			"back_attachment_format": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the attachment output format to backend servers. If an attachment arrives that is not of the type chosen, the attachment will be rejected.", "back-attachment-format", "").AddStringEnum("dynamic", "mime", "dime").AddDefaultValue("dynamic").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("dynamic", "mime", "dime"),
				},
				Default: stringdefault.StaticString("dynamic"),
			},
			"mime_headers": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>The body of a message (that is, the payload, independent of any protocol headers) can sometimes contain MIME headers before any preamble and before the first MIME boundary contained in the body of the message. These MIME headers can contain important information not available in the protocol headers, such as the string identifying the MIME boundary. If this property is set to on, then these MIME headers will be processed by the Firewall.</p><p>Note that if this is on and there are no MIME headers contained in the message, the device will continue to try and parse the message, using the protocol header information, if available.</p><p>When this is off and MIME headers are present in the body of the message, these MIME headers will be considered part of the preamble and not used to parse out the message. If the protocol headers (such as HTTP) indicate the MIME boundaries, the device can parse and process individual attachments. If such information is not available, no attachments can be parsed from the body of the message.</p>", "mime-headers", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"rewrite_errors": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Error messages after a decryption action can provide an attacker who is using the padding oracle attack method with enough information to determine the contents of the plaintext data. When enabled, the default, the client receives error messages without the internal information that could lead to a discovery. When disabled, the client receives the original message with this information.", "rewrite-errors", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"delay_errors": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The timing difference of the error messages returned after a decryption action can provide an attacker with enough information to determine the contents of the plain-text data. When enabled, the default, the DataPower Gateway delays error messages for the defined duration. When disabled, the DataPower Gateway does not delay error messages.", "delay-errors", "").AddDefaultValue("true").AddNotValidWhen(models.XMLFirewallServiceDelayErrorsIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"delay_errors_duration": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("When enabling the delay of error messages, specify the delay duration in milliseconds. If delaying messages for 3000 ms, the DataPower Gateway will not send error messages to the client until 3 seconds have elapsed since the DataPower Gateway performed decryption on the requests. Use any value of 250 - 300000. The default value is 1000.", "delay-errors-duration", "").AddIntegerRange(250, 300000).AddDefaultValue("1000").AddRequiredWhen(models.XMLFirewallServiceDelayErrorsDurationCondVal.String()).AddNotValidWhen(models.XMLFirewallServiceDelayErrorsDurationIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(250, 300000),
					validators.ConditionalRequiredInt64(models.XMLFirewallServiceDelayErrorsDurationCondVal, models.XMLFirewallServiceDelayErrorsDurationIgnoreVal, true),
				},
				Default: int64default.StaticInt64(1000),
			},
			"soap_schema_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("When a firewall is in SOAP mode, either on the request or response side, it will validate the incoming messages against a W3C Schema that defines what a SOAP message is supposed to look like. It is possible to customize which schema is used on a per-firewall basis by changing this property; this can be used to accommodate non-standard configurations or other special cases.", "soap-schema-url", "").AddDefaultValue("store:///schemas/soap-envelope.xsd").AddNotValidWhen(models.XMLFirewallServiceSOAPSchemaURLIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("store:///schemas/soap-envelope.xsd"),
			},
			"wsdl_response_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select how the firewall handles requests for .NET WSDL requests via the http://domain.com/service?wsdl convention. The default is Off.", "wsdl-response-policy", "").AddStringEnum("off", "intercept", "serve").AddDefaultValue("off").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("off", "intercept", "serve"),
				},
				Default: stringdefault.StaticString("off"),
			},
			"wsdl_file_location": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("URL of the file to return when WSDL response policy is set to serve local.", "wsdl-file-location", "").AddNotValidWhen(models.XMLFirewallServiceWSDLFileLocationIgnoreVal.String()).String,
				Optional:            true,
			},
			"firewall_parser_limits": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use the firewall parser limits instead of the parser limits in the XML Manager for this firewall. Firewall limits override XML Manager limits.", "firewall-parser-limits", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"parser_limits_bytes_scanned": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Defines the maximum number of bytes scanned by the XML parser. This applies to any XML document that is parsed. If any of the parser limits are set in the XML Firewall, they will override those on the XML Manager. If this value is 0, no limit is enforced.", "bytes-scanned", "").AddDefaultValue("4194304").AddRequiredWhen(models.XMLFirewallServiceParserLimitsBytesScannedCondVal.String()).AddNotValidWhen(models.XMLFirewallServiceParserLimitsBytesScannedIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					validators.ConditionalRequiredInt64(models.XMLFirewallServiceParserLimitsBytesScannedCondVal, models.XMLFirewallServiceParserLimitsBytesScannedIgnoreVal, true),
				},
				Default: int64default.StaticInt64(4194304),
			},
			"parser_limits_element_depth": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Defines the maximum depth of element nesting in XML parser. If any of the parser limits are set in the XML Firewall, they will override those on the XML Manager.", "element-depth", "").AddIntegerRange(0, 65535).AddDefaultValue("512").AddRequiredWhen(models.XMLFirewallServiceParserLimitsElementDepthCondVal.String()).AddNotValidWhen(models.XMLFirewallServiceParserLimitsElementDepthIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
					validators.ConditionalRequiredInt64(models.XMLFirewallServiceParserLimitsElementDepthCondVal, models.XMLFirewallServiceParserLimitsElementDepthIgnoreVal, true),
				},
				Default: int64default.StaticInt64(512),
			},
			"parser_limits_attribute_count": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Defines the maximum number of attributes of a given element. If any of the parser limits are set in the XML Firewall, they will override those on the XML Manager.", "attribute-count", "").AddIntegerRange(0, 65535).AddDefaultValue("128").AddRequiredWhen(models.XMLFirewallServiceParserLimitsAttributeCountCondVal.String()).AddNotValidWhen(models.XMLFirewallServiceParserLimitsAttributeCountIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
					validators.ConditionalRequiredInt64(models.XMLFirewallServiceParserLimitsAttributeCountCondVal, models.XMLFirewallServiceParserLimitsAttributeCountIgnoreVal, true),
				},
				Default: int64default.StaticInt64(128),
			},
			"parser_limits_max_node_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Defines the maximum size any one node may consume. The default is 32 MB. Sizes which are powers of two result in the best performance. If any of the parser limits are set in the XML Firewall, they will override those on the XML Manager. Although you set an explicit value, the DataPower Gateway uses a value that is the rounded-down, largest power of two that is smaller than the defined value.", "max-node-size", "").AddIntegerRange(1024, 4294967295).AddDefaultValue("33554432").AddRequiredWhen(models.XMLFirewallServiceParserLimitsMaxNodeSizeCondVal.String()).AddNotValidWhen(models.XMLFirewallServiceParserLimitsMaxNodeSizeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1024, 4294967295),
					validators.ConditionalRequiredInt64(models.XMLFirewallServiceParserLimitsMaxNodeSizeCondVal, models.XMLFirewallServiceParserLimitsMaxNodeSizeIgnoreVal, true),
				},
				Default: int64default.StaticInt64(33554432),
			},
			"parser_limits_max_prefixes": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter an integer that defines the maximum number of distinct XML namespace prefixes in a document. This limit counts multiple prefixes defined for the same namespace, but does not count multiple namespaces defined in different parts of the input document under a single prefix. Enter a value in the rage 0 - 262143. The default value is 1024. A value of 0 indicates that the limit is 1024.", "max-prefixes", "").AddIntegerRange(0, 4294967295).AddDefaultValue("1024").AddRequiredWhen(models.XMLFirewallServiceParserLimitsMaxPrefixesCondVal.String()).AddNotValidWhen(models.XMLFirewallServiceParserLimitsMaxPrefixesIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
					validators.ConditionalRequiredInt64(models.XMLFirewallServiceParserLimitsMaxPrefixesCondVal, models.XMLFirewallServiceParserLimitsMaxPrefixesIgnoreVal, true),
				},
				Default: int64default.StaticInt64(1024),
			},
			"parser_limits_max_namespaces": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter an integer that defines the maximum number of distinct XML namespace URIs in a document. This limit counts all XML namespaces, regardless of how many prefixes are used to declare them. Enter a value in the range 0 - 65535. The default value is 1024. A value of 0 indicates that the limit is 1024.", "max-namespaces", "").AddDefaultValue("1024").AddRequiredWhen(models.XMLFirewallServiceParserLimitsMaxNamespacesCondVal.String()).AddNotValidWhen(models.XMLFirewallServiceParserLimitsMaxNamespacesIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					validators.ConditionalRequiredInt64(models.XMLFirewallServiceParserLimitsMaxNamespacesCondVal, models.XMLFirewallServiceParserLimitsMaxNamespacesIgnoreVal, true),
				},
				Default: int64default.StaticInt64(1024),
			},
			"parser_limits_max_local_names": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter an integer that defines the maximum number of distinct XML local names in a document. This limit counts all local names, independent of the namespaces they are associated with. Enter a value in the range 0 - 1048575. The default value is 60000. A value of 0 indicates that the limit is 60000.", "max-local-names", "").AddIntegerRange(0, 4294967295).AddDefaultValue("60000").AddRequiredWhen(models.XMLFirewallServiceParserLimitsMaxLocalNamesCondVal.String()).AddNotValidWhen(models.XMLFirewallServiceParserLimitsMaxLocalNamesIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
					validators.ConditionalRequiredInt64(models.XMLFirewallServiceParserLimitsMaxLocalNamesCondVal, models.XMLFirewallServiceParserLimitsMaxLocalNamesIgnoreVal, true),
				},
				Default: int64default.StaticInt64(60000),
			},
			"parser_limits_attachment_byte_count": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Defines the maximum number of bytes allowed in any single attachment. Attachments that exceed this size will result in a failure of the whole transaction. If this value is 0, no limit is enforced.", "attachment-byte-count", "").AddIntegerRange(0, 4294967295).AddDefaultValue("2000000000").AddRequiredWhen(models.XMLFirewallServiceParserLimitsAttachmentByteCountCondVal.String()).AddNotValidWhen(models.XMLFirewallServiceParserLimitsAttachmentByteCountIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
					validators.ConditionalRequiredInt64(models.XMLFirewallServiceParserLimitsAttachmentByteCountCondVal, models.XMLFirewallServiceParserLimitsAttachmentByteCountIgnoreVal, true),
				},
				Default: int64default.StaticInt64(2000000000),
			},
			"parser_limits_attachment_package_byte_count": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Defines the maximum number of bytes allowed for all parts of an attachment package, including the root part. Attachment packages that exceed this size will result in a failure of the transaction. If this value is 0, no limit is enforced.", "attachment-package-byte-count", "").AddRequiredWhen(models.XMLFirewallServiceParserLimitsAttachmentPackageByteCountCondVal.String()).AddNotValidWhen(models.XMLFirewallServiceParserLimitsAttachmentPackageByteCountIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.Int64{
					validators.ConditionalRequiredInt64(models.XMLFirewallServiceParserLimitsAttachmentPackageByteCountCondVal, models.XMLFirewallServiceParserLimitsAttachmentPackageByteCountIgnoreVal, false),
				},
			},
			"parser_limits_external_references": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the special handling for input documents that contain external references, such as an external entity or external DTD definition.", "external-references", "").AddStringEnum("forbid", "ignore", "allow").AddDefaultValue("forbid").AddRequiredWhen(models.XMLFirewallServiceParserLimitsExternalReferencesCondVal.String()).AddNotValidWhen(models.XMLFirewallServiceParserLimitsExternalReferencesIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("forbid", "ignore", "allow"),
					validators.ConditionalRequiredString(models.XMLFirewallServiceParserLimitsExternalReferencesCondVal, models.XMLFirewallServiceParserLimitsExternalReferencesIgnoreVal, true),
				},
				Default: stringdefault.StaticString("forbid"),
			},
			"credential_charset": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the character encoding of the original basic authentication values. Basic authentication credentials are combined and base64 encoded in the authorization header of the request. The DataPower Gateway transcodes the contents of the authorization header to UTF-8. Defaults to Protocol which is ISO-8859-1, Latin 1.", "credential-charset", "").AddStringEnum("protocol", "ascii", "utf8", "big5", "cp1250", "cp1251", "cp1252", "cp1253", "cp1254", "cp1255", "cp1256", "cp1257", "cp1258", "euc_jp", "euc_kr", "gb18030", "gb2312", "iso2022_jp", "iso2022_kr", "iso8859_1", "iso8859_2", "iso8859_4", "iso8859_5", "iso8859_6", "iso8859_7", "iso8859_8", "iso8859_9", "iso8859_15", "sjis", "tis620", "unicode_le").AddDefaultValue("protocol").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("protocol", "ascii", "utf8", "big5", "cp1250", "cp1251", "cp1252", "cp1253", "cp1254", "cp1255", "cp1256", "cp1257", "cp1258", "euc_jp", "euc_kr", "gb18030", "gb2312", "iso2022_jp", "iso2022_kr", "iso8859_1", "iso8859_2", "iso8859_4", "iso8859_5", "iso8859_6", "iso8859_7", "iso8859_8", "iso8859_9", "iso8859_15", "sjis", "tis620", "unicode_le"),
				},
				Default: stringdefault.StaticString("protocol"),
			},
			"ssl_config_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS type", "ssl-config-type", "").AddStringEnum("server", "sni").AddDefaultValue("server").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("server", "sni"),
				},
				Default: stringdefault.StaticString("server"),
			},
			"ssl_server": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS server profile", "ssl-server", "ssl_server_profile").AddNotValidWhen(models.XMLFirewallServiceSSLServerIgnoreVal.String()).String,
				Optional:            true,
			},
			"ssl_sni_server": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS SNI server profile", "ssl-sni-server", "ssl_sni_server_profile").AddNotValidWhen(models.XMLFirewallServiceSSLSNIServerIgnoreVal.String()).String,
				Optional:            true,
			},
			"ssl_client": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "ssl_client_profile").AddNotValidWhen(models.XMLFirewallServiceSSLClientIgnoreVal.String()).String,
				Optional:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"priority": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Control the service scheduling priority. When system resources are in high demand, \"high\" priority services will be favored over lower priority services.", "priority", "").AddStringEnum("unknown", "high-min", "high", "high-max", "normal-min", "normal", "normal-max", "low-min", "low", "low-max").AddDefaultValue("normal").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("unknown", "high-min", "high", "high-max", "normal-min", "normal", "normal-max", "low-min", "low", "low-max"),
				},
				Default: stringdefault.StaticString("normal"),
			},
			"local_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the local port to monitor for incoming client requests.", "port", "").AddIntegerRange(1, 65535).String,
				Required:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"remote_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the host name or IP address of the specific server supported by this DataPower service. If using load balancers, specify the name of the Load Balancer Group. If using the On Demand Router, specify the keyword ODR-LBG. Load balancer groups and the On Demand Router can be used only when Type is static-backend.", "remote-ip-address", "").AddRequiredWhen(models.XMLFirewallServiceRemoteAddressCondVal.String()).AddNotValidWhen(models.XMLFirewallServiceRemoteAddressIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.XMLFirewallServiceRemoteAddressCondVal, models.XMLFirewallServiceRemoteAddressIgnoreVal, false),
				},
			},
			"remote_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the port number to monitor. Used only when Type is static-backend.", "remote-port", "").AddIntegerRange(1, 65535).AddRequiredWhen(models.XMLFirewallServiceRemotePortCondVal.String()).AddNotValidWhen(models.XMLFirewallServiceRemotePortIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
					validators.ConditionalRequiredInt64(models.XMLFirewallServiceRemotePortCondVal, models.XMLFirewallServiceRemotePortIgnoreVal, false),
				},
			},
			"acl": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("This Access Control List will be used to allow or deny access to this service based on the IP address of the client. When attached to a service, an Access Control List (ACL) denies all access by default. To deny access to only selected addresses, first grant access to all addresses (allow 0.0.0.0) and then create deny entries for the desired hosts.", "acl", "access_control_list").String,
				Optional:            true,
			},
			"http_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the maximum number of seconds (within the range 1 through 86400, with a default of 120) that a firewall or proxy maintains an idle HTTP connection.", "http-timeout", "").AddIntegerRange(1, 86400).AddDefaultValue("120").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(120),
			},
			"http_persist_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the maximum number of seconds (within the range 0 through 7200, with a default of 180) that a firewall or proxy maintains an idle TCP connection.", "persistent-timeout", "").AddIntegerRange(0, 7200).AddDefaultValue("180").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 7200),
				},
				Default: int64default.StaticInt64(180),
			},
			"do_host_rewrite": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("When enabled, the device will rewrite the Host: header to be the address of the back-end server. This is not what a strict proxy would do, and may not be appropriate for all topologies.", "host-rewriting", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"suppress_http_warnings": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Toggle to enable or disable the generation of Transformation Applied (Warning Code: 214) messages by the HTTP service.", "silence-warning", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_compression": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Toggle to enable or disable the GZIP compression function.", "compression", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_include_response_type_encoding": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Toggle to enable or disable including the character set encoding in the HTTP content-type header produced. For example, when sending a UTF-8 encoded XML document: If this property is disabled, 'content-type=text/xml' will be sent. If this property is enabled, 'content-type=text/xml; charset=UTF-8' will be sent.", "include-response-type-encoding", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"always_show_errors": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("If set, HTTP responses may be generated with errors appended to partially generated documents. If not set error responses will only be sent to the browser if no other output has been created.", "always-show-errors", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"disallow_get": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("If set, only methods with incoming data (such as POST) are allowed.", "disallow-get", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"disallow_empty_response": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("If set, only responses with message bodies are allowed (that is, not 304 and so forth).", "disallow-empty-reply", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_persistent_connections": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Toggle to enable or disable HTTP persistent connections.", "persistent-connections", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"http_client_ip_label": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Retain the default value (X-Client-IP) or provide an other value (for example, X-Forwarded-For).", "client-address", "").AddDefaultValue("X-Client-IP").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("X-Client-IP"),
			},
			"http_log_cor_id_label": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of an HTTP Header to read to determine the global transaction ID for chained services. This value defaults to X-Global-Transaction-ID.", "http-global-tranID-label", "").AddDefaultValue("X-Global-Transaction-ID").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("X-Global-Transaction-ID"),
			},
			"http_proxy_host": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the host name or IP address of the HTTP proxy to use when the remote server can be accessed only through another HTTP proxy.", "httpproxy-address", "").String,
				Optional:            true,
			},
			"http_proxy_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the port number on the HTTP proxy server.", "httpproxy-port", "").AddIntegerRange(1, 65535).AddDefaultValue("800").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(800),
			},
			"http_version": models.GetDmHTTPClientServerVersionResourceSchema("HTTP Version", "version", "", false),
			"do_chunked_upload": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use the radio buttons to enable (on) or disable (off) the ability to send Content-Type Chunked Encoded documents to the back end server. When the device employs the HTTP/1.1 protocol, the body of the document can be delimited by either Content-Length or chunked encodings. While all servers will understand how to interpret Content-Length, many applications will fail to understand Chunked encoding. For this reason, Content-Length is the standard method used. However doing so interferes with the ability of the device to fully stream. To stream full documents towards the back end server, this property should be turned on. However, the back end server must be RFC 2616 compatible, because this feature cannot be renegotiated at run time, unlike all other HTTP/1.1 features which can be negotiated down at runtime if necessary. This property can also be enabled by configuring a User Agent to enable it on a per-URL basis.", "chunked-uploads", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"header_injection": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP Header Injection", "inject", "").String,
				NestedObject:        models.GetDmHeaderInjectionResourceSchema(),
				Optional:            true,
			},
			"header_suppression": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP Headers can be suppressed (removed) from the message flow using this property. For example, the Via: header, which contains the name of the DataPower service handling the message, may be suppressed from messages sent by the DataPower device back to the client.", "suppress", "").String,
				NestedObject:        models.GetDmHeaderSuppressionResourceSchema(),
				Optional:            true,
			},
			"stylesheet_parameters": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Stylesheets used in Processing Policies can take stylesheet parameters. These parameters can be passed in by this object. More than one parameter can be defined.", "parameter", "").String,
				NestedObject:        models.GetDmStylesheetParameterResourceSchema(),
				Optional:            true,
			},
			"default_param_namespace": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("If a stylesheet parameter is defined without a namespace (or without explicitly specifying the null namespace), then this is the namespace into which the parameter will be assigned.", "default-param-namespace", "").AddDefaultValue("http://www.datapower.com/param/config").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("http://www.datapower.com/param/config"),
			},
			"query_param_namespace": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The namespace in which to put all parameters that are specified in the URL query string.", "query-param-namespace", "").AddDefaultValue("http://www.datapower.com/param/query").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("http://www.datapower.com/param/query"),
			},
			"force_policy_exec": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Some message patterns may include bodyless request and response messages. This is common with RESTful web services where messages may or may not include a body but still requires the processing policy to run. To enable this capability for services whose request and response type is XML (or marked as non-XML i.e. JSON), set this option to 'on'. By doing so, the processing policy rules will always be executed.</p>", "force-policy-exec", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"count_monitors": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Count Monitors watch for defined messaging events and increment counters each time the event occurs. When a certain threshold is reached, the monitor can either post a notification to a log or block service for a configured amount of time. Select a Count Monitor from the list to activate that monitor for this firewall. Click the + button to create a new Count Monitor.", "monitor-count", "count_monitor").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"duration_monitors": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Duration Monitors watch for events meeting or exceeding a configured duration. When a duration is met or exceeded, the monitor can either post a notification to a log or block service for a configured amount of time. Select a Duration Monitor from the list to activate that monitor for this firewall. Click the + button to create a new Duration Monitor.", "monitor-duration", "duration_monitor").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"monitor_processing_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the way that the system behaves when more than one monitor is attached to a service.", "monitor-processing-policy", "").AddStringEnum("terminate-at-first-throttle", "terminate-at-first-match").AddDefaultValue("terminate-at-first-throttle").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("terminate-at-first-throttle", "terminate-at-first-match"),
				},
				Default: stringdefault.StaticString("terminate-at-first-throttle"),
			},
			"debug_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Select the diagnostic mode for processing policies. When enabled, you can view details about the state of variables and contexts for a captured transaction in the probe. The default value is <tt>off</tt> .</p><p>Transaction diagnostic mode is not intended for use in a production environment. Transaction diagnostic mode consumes significant resources that can slow down transaction processing.</p>", "debug-mode", "").AddStringEnum("on", "off", "unbounded").AddDefaultValue("off").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("on", "off", "unbounded"),
				},
				Default: stringdefault.StaticString("off"),
			},
			"debug_history": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Set the number of records for transaction diagnostic mode in the probe. Enter a value in the range 10 - 250. The default value is 25.", "debug-history", "").AddIntegerRange(10, 250).AddDefaultValue("25").AddRequiredWhen(models.XMLFirewallServiceDebugHistoryCondVal.String()).AddNotValidWhen(models.XMLFirewallServiceDebugHistoryIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(10, 250),
					validators.ConditionalRequiredInt64(models.XMLFirewallServiceDebugHistoryCondVal, models.XMLFirewallServiceDebugHistoryIgnoreVal, true),
				},
				Default: int64default.StaticInt64(25),
			},
			"debug_trigger": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The probe captures transactions that meet one or more of the conditions defined by the triggers. These triggers examine the direction or type of the message flow and examine the message for an XPath expression match. When a message meets one of these conditions, the transaction is captured in diagnostics mode and becomes part of the list of transactions that can be viewed.", "debug-trigger", "").AddNotValidWhen(models.XMLFirewallServiceDebugTriggerIgnoreVal.String()).String,
				NestedObject:        models.GetDmMSDebugTriggerTypeResourceSchema(),
				Optional:            true,
			},
			"local_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Enter a host alias or the IP address that the service listens on. Host aliases can ease migration tasks among appliances.</p><ul><li>0 or 0.0.0.0 indicates all configured IPv4 addresses.</li><li>:: indicates all configured IPv4 and IPv6 addresses.</li></ul><p><b>Attention:</b> For management services, the value of 0.0.0.0 or :: is a security risk. Use an explicit IP address to isolate management traffic from application data traffic.</p>", "ip-address", "").AddDefaultValue("0.0.0.0").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("0.0.0.0"),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *XMLFirewallServiceResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *XMLFirewallServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.XMLFirewallService
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `XMLFirewallService`)
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

func (r *XMLFirewallServiceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.XMLFirewallService
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

	data.UpdateFromBody(ctx, `XMLFirewallService`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *XMLFirewallServiceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.XMLFirewallService
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `XMLFirewallService`))
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

func (r *XMLFirewallServiceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.XMLFirewallService
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

func (r *XMLFirewallServiceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

	var data models.XMLFirewallService
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

	data.FromBody(ctx, `XMLFirewallService`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *XMLFirewallServiceResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.XMLFirewallService

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
