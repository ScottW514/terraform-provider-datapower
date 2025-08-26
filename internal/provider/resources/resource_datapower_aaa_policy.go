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

var _ resource.Resource = &AAAPolicyResource{}
var _ resource.ResourceWithValidateConfig = &AAAPolicyResource{}

func NewAAAPolicyResource() resource.Resource {
	return &AAAPolicyResource{}
}

type AAAPolicyResource struct {
	pData *tfutils.ProviderData
}

func (r *AAAPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaa_policy"
}

func (r *AAAPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("<p>Define the support to authenticate users and authorize their access to resources. An AAA policy consists of the following phases.</p><ul><li>During AAA processing, the identity extraction phase defines which methods the AAA policy uses to extract the claimed identity of the service requester.</li><li>After the claimed identity of the service requester is extracted, an AAA policy authenticates the claimed identity. The authentication process can use internal or external resources.</li><li>After authentication credentials are received, an AAA policy can map these credentials.</li><li>After client authentication, an AAA policy identifies the specific resource that is being requested by that client.</li><li>After requested resources are identified, you might need to map extracted resource to a form that is compatible with the authorization method.</li><li>After the service requester is authenticated from the extracted identity, an AAA policy authorizes the client to the requested resource.</li><li>After client-authorization, an AAA policy can run postprocessing activities.</li></ul>", "aaapolicy", "").AddActions("flush_cache").String,
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
			"authorized_counter": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the monitor for authorized messages. Ensure that this count monitor is configured with an XPath as the measure.", "authorized-counter", "count_monitor").String,
				Optional:            true,
			},
			"rejected_counter": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the monitor for rejected messages. Ensure that this count monitor is configured with an XPath as the measure.", "rejected-counter", "count_monitor").String,
				Optional:            true,
			},
			"namespace_mapping": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Define XML namespace maps. Each map is a prefix with its URI.", "namespace-mapping", "").String,
				NestedObject:        models.GetDmNamespaceMappingResourceSchema(),
				Optional:            true,
			},
			"extract_identity": models.GetDmAAAPExtractIdentityResourceSchema("Specify the methods to extract the identity of the service requester. For some methods, you must define more properties.", "extract-identity", "", false),
			"authenticate":     models.GetDmAAAPAuthenticateResourceSchema("Specify the method to authenticate the extracted identity. For some methods, you must define more properties.", "authenticate", "", false),
			"map_credentials":  models.GetDmAAAPMapCredentialsResourceSchema("Specify the method to map credentials for authorization. For some methods, you must define more properties.", "map-credentials", "", false),
			"extract_resource": models.GetDmAAAPExtractResourceResourceSchema("Specify the methods to extract the identity of a requested resource. For some methods, you must define more properties.", "extract-resource", "", false),
			"map_resource":     models.GetDmAAAPMapResourceResourceSchema("Specify the method to map resources for authorization. For some methods, you must define more properties.", "map-resource", "", false),
			"authorize":        models.GetDmAAAPAuthorizeResourceSchema("Specify the method to authorize the identity to resources. For some methods, you must define more properties.", "authorize", "", false),
			"post_process":     models.GetDmAAAPPostProcessResourceSchema("Specify postprocessing activities. For some methods, you must define more properties.", "post-process", "", false),
			"saml_attribute": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify SAML attributes. Each attribute consists of its namespace URI, local name, and expected value.", "saml-attribute", "").String,
				NestedObject:        models.GetDmSAMLAttributeNameAndValueResourceSchema(),
				Optional:            true,
			},
			"ltpa_attributes": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify user attributes to include in the LTPA token. Attributes are relevant for only WebSphere tokens.</p><p>For each attribute, its value can be static or resolved at run time. <ul><li>When static, its value is a fixed value.</li><li>When resolved at run time, its value is resolved at run time with an XPath expression.</li></ul></p>", "ltpa-attribute", "").String,
				NestedObject:        models.GetDmLTPAUserAttributeNameAndValueResourceSchema(),
				Optional:            true,
			},
			"transaction_priority": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Define the transactional priority for users. For each user, you must specify the name of the output credentials, their priority for scheduling or resource allocation, and whether authorization is required.", "transaction-priority", "").String,
				NestedObject:        models.GetDmAAATransactionPriorityResourceSchema(),
				Optional:            true,
			},
			"saml_valcred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SAML signature validation credentials", "saml-valcred", "crypto_val_cred").String,
				Optional:            true,
			},
			"saml_signing_key": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the key to sign SAML messages. To sign messages, you must specify a key and a certificate.", "saml-sign-key", "crypto_key").String,
				Optional:            true,
			},
			"saml_signing_cert": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the certificate to sign SAML messages. To sign messages, you must specify a key and a certificate.", "saml-sign-cert", "crypto_certificate").String,
				Optional:            true,
			},
			"saml_signing_hash_alg": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the algorithm to calculate the message digest for signing. The default value is sha1.", "saml-sign-hash", "").AddStringEnum("sha1", "sha256", "sha512", "ripemd160", "sha224", "sha384", "md5").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("sha1", "sha256", "sha512", "ripemd160", "sha224", "sha384", "md5"),
				},
			},
			"saml_signing_alg": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the algorithm to sign SAML messages. The default value is rsa.", "saml-sign-alg", "").AddStringEnum("rsa-sha1", "dsa-sha1", "rsa-sha256", "rsa-sha384", "rsa-sha512", "rsa-ripemd160", "rsa-ripemd160-2010", "sha256-rsa-MGF1", "rsa-md5", "rsa", "dsa", "ecdsa-sha1", "ecdsa-sha224", "ecdsa-sha256", "ecdsa-sha384", "ecdsa-sha512").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("rsa-sha1", "dsa-sha1", "rsa-sha256", "rsa-sha384", "rsa-sha512", "rsa-ripemd160", "rsa-ripemd160-2010", "sha256-rsa-MGF1", "rsa-md5", "rsa", "dsa", "ecdsa-sha1", "ecdsa-sha224", "ecdsa-sha256", "ecdsa-sha384", "ecdsa-sha512"),
				},
			},
			"ldap_suffix": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the LDAP suffix to add to the username to form the base distinguished name (DN) for authentication. The suffix and the username are separated with a comma. If the suffix is <tt>O=example.com</tt> and the username is <tt>Bob</tt> , the DN is <tt>CN=Bob,O=example.com</tt> .", "ldap-suffix", "").String,
				Optional:            true,
			},
			"log_allowed": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Whether to log successful AAA operations. When enabled and if needed, modify the default logging level from informational.", "log-allowed", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"log_allowed_level": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Set the logging level for successful AAA operations. The default level is informational.", "log-allowed-level", "").AddStringEnum("emerg", "alert", "critic", "error", "warn", "notice", "info", "debug").AddDefaultValue("info").AddRequiredWhen(models.AAAPolicyLogAllowedLevelCondVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("emerg", "alert", "critic", "error", "warn", "notice", "info", "debug"),
					validators.ConditionalRequiredString(models.AAAPolicyLogAllowedLevelCondVal, validators.Evaluation{}, true),
				},
				Default: stringdefault.StaticString("info"),
			},
			"log_rejected": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Whether to log unsuccessful AAA operations. When enabled and if needed, modify the default logging level from warning.", "log-rejected", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"log_rejected_level": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Set the logging level for unsuccessful AAA operations. The default level is warning.", "log-rejected-level", "").AddStringEnum("emerg", "alert", "critic", "error", "warn", "notice", "info", "debug").AddDefaultValue("warn").AddRequiredWhen(models.AAAPolicyLogRejectedLevelCondVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("emerg", "alert", "critic", "error", "warn", "notice", "info", "debug"),
					validators.ConditionalRequiredString(models.AAAPolicyLogRejectedLevelCondVal, validators.Evaluation{}, true),
				},
				Default: stringdefault.StaticString("warn"),
			},
			"ws_secure_conversation_crypto_key": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WS-Trust encryption recipient certificate", "wstrust-encrypt-key", "crypto_certificate").String,
				Optional:            true,
			},
			"saml_source_id_mapping_file": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the SAML artifact-mapping file that provides a mapping of SAML artifact source IDs to artifact retrieval endpoints. Use this setting when artifacts are retrieved from multiple endpoints and the source ID for these endpoints are encoded in the artifact itself. If only one artifact retrieval URL exists, it can be specified by the SAML artifact responder URL in the authentication phase.", "saml-artifact-mapping", "").String,
				Optional:            true,
			},
			"ping_identity_compatibility": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Control compatibility with a PingFederate identity server. By default, compatibility is disabled. Enable compatibility for SAML authentication or authorization.", "ping-identity-compatibility", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"saml2_metadata_file": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the SAML 2.0 metadata file for SAML 2.0 protocol message exchanges. The metadata in this file identifies identity provider endpoints and certificates to secure message exchanges. The file must have a root-level &lt;md:EntitiesDescriptor> element with an &lt;EntityDescriptor> child element for each identity provider.", "saml2-metadata", "").String,
				Optional:            true,
			},
			"do_s_valve": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the number of times to process the same request to protect against a denial of service (DoS) attack. Enter a value in the range 1 - 1000. The default value is 3.</p><p>With the default value, AAA processes only the first 3 signature and each signature can contain up to 3 reference URIs. Additional signatures or reference URIs are ignored.</p><p>XML processing includes encryption, decryption, message signing, and signature validation. The AAA policy supports only identity extraction with subject DN from certificate in message signature and authorization with signer certificate for digitally signed messages.</p>", "dos-valve", "").AddIntegerRange(1, 1000).AddDefaultValue("3").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 1000),
				},
				Default: int64default.StaticInt64(3),
			},
			"ldap_version": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the LDAP version to access the LDAP server. The default value is v2.", "ldap-version", "").AddStringEnum("v2", "v3").AddDefaultValue("v2").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("v2", "v3"),
				},
				Default: stringdefault.StaticString("v2"),
			},
			"enforce_soap_actor": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Whether to enforce the <tt>S11:actor</tt> or <tt>S12:role</tt> on WS-Security messages. In general, a WS-Security message has a <tt>S11:actor</tt> or <tt>S12:role</tt> attribute for its <tt>Security</tt> header. Processing can enforce these attributes when the AAA policy tries to use the <tt>Security</tt> header. For example, the <tt>Security</tt> element has only one actor or role. In this case, the AAA policy processes only the <tt>Security</tt> header for this actor or role identifier. This setting applies to all AAA phases, except postprocessing. For postprocessing, the activity generally generates a new message for next SOAP node.", "enforce-actor-role", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ws_sec_actor_role_id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Set the assumed <tt>S11:actor</tt> or <tt>S12:role</tt> identifier. The AAA policy acts as the assumed actor or role when it consumes <tt>Security</tt> headers. This setting takes effect only when the AAA policy attempts to process the incoming message before it makes an authorization decision. Postprocessing does not use this setting. Postprocessing uses its own setting in generating the message for the next SOAP node. The default value is an empty string. <table border=\"1\"><tr><td valign=\"left\"><tt>http://schemas.xmlsoap.org/soap/actor/next</tt></td><td>Every one, including the intermediary and ultimate receiver, that receives the message can process the <tt>Security</tt> header.</td></tr><tr><td valign=\"left\"><tt>http://www.w3.org/2003/05/soap-envelope/role/none</tt></td><td>No one can process the <tt>Security</tt> header.</td></tr><tr><td valign=\"left\"><tt>http://www.w3.org/2003/05/soap-envelope/role/next</tt></td><td>Every one, including the intermediary and ultimate receiver, that receives the message can process the <tt>Security</tt> header.</td></tr><tr><td valign=\"left\"><tt>http://www.w3.org/2003/05/soap-envelope/role/ultimateReceiver</tt></td><td>The ultimate receiver can process the <tt>Security</tt> header.</td></tr><tr><td valign=\"left\">No value, which is an empty string</td><td>The empty string (without quotation marks) indicates that no \"actor/role\" identifier is configured. With a configured actor/role, the ultimate receiver is assumed for the message. No actor/role attribute is added during the generation of the <tt>Security</tt> header. More than one <tt>Security</tt> header cannot omit the actor/role identifier.</td></tr><tr><td valign=\"left\"><tt>USE_MESSAGE_BASE_URI</tt></td><td>The identifier is the base URL of the message. When the SOAP message is transported, the base URI is the request-URI of the HTTP request.</td></tr><tr><td valign=\"left\">A string value</td><td>Any string to identify the actor or role of the <tt>Security</tt> header.</td></tr></table>", "actor-role-id", "").String,
				Optional:            true,
			},
			"au_sm_http_header": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify HTTP headers from CA Single Sign-On authentication responses. These headers are included as request or response headers based on the CA Single Sign-on header flow.", "au-sm-http-header", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"az_sm_http_header": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify HTTP headers from CA Single Sign-On authorization responses. These headers are included as request or response headers based on the CA Single Sign-On header flow.", "az-sm-http-header", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"dyn_config": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how to obtain the AAA policy configuration dynamically. With dynamic configuration, you can configure the AAA policy at run time. When enabled, the configuration of AAA is determined dynamically based on the template AAA policy and the parameters that the dynamic configuration custom URL returns. By default, uses no template.", "dyn-config", "").AddStringEnum("none", "current-aaa", "external-aaa").AddDefaultValue("none").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("none", "current-aaa", "external-aaa"),
				},
				Default: stringdefault.StaticString("none"),
			},
			"external_aaa_template": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify another AAA policy to use as the template. When specified, this AAA policy overwrites the current AAA policy.", "external-aaa-template", "aaa_policy").AddRequiredWhen(models.AAAPolicyExternalAAATemplateCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AAAPolicyExternalAAATemplateCondVal, validators.Evaluation{}, false),
				},
			},
			"dyn_config_custom_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the location of the custom stylesheet or GatewayScript file. The configuration of the AAA policy is obtained dynamically from this file. The obtained configuration overwrites the configuration in the template AAA policy.</p><p>In the custom file, modify only the properties to dynamically overwrite. See the <tt>ModifyAAAPolicy</tt> element in the <tt>store:///xml-mgmt.xsd</tt> schema to construct a schema-compliant AAA configuration.</p>", "dyn-config-custom-url", "").AddRequiredWhen(models.AAAPolicyDynConfigCustomURLCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AAAPolicyDynConfigCustomURLCondVal, validators.Evaluation{}, false),
				},
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *AAAPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *AAAPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AAAPolicy
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `AAAPolicy`)
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

func (r *AAAPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AAAPolicy
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
		data.FromBody(ctx, `AAAPolicy`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `AAAPolicy`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AAAPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AAAPolicy
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `AAAPolicy`))
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

func (r *AAAPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AAAPolicy
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

func (r *AAAPolicyResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.AAAPolicy

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
