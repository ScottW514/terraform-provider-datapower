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

type AAAPolicyList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &AAAPolicyDataSource{}
	_ datasource.DataSourceWithConfigure = &AAAPolicyDataSource{}
)

func NewAAAPolicyDataSource() datasource.DataSource {
	return &AAAPolicyDataSource{}
}

type AAAPolicyDataSource struct {
	pData *tfutils.ProviderData
}

func (d *AAAPolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaa_policy"
}

func (d *AAAPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "<p>Define the support to authenticate users and authorize their access to resources. An AAA policy consists of the following phases.</p><ul><li>During AAA processing, the identity extraction phase defines which methods the AAA policy uses to extract the claimed identity of the service requester.</li><li>After the claimed identity of the service requester is extracted, an AAA policy authenticates the claimed identity. The authentication process can use internal or external resources.</li><li>After authentication credentials are received, an AAA policy can map these credentials.</li><li>After client authentication, an AAA policy identifies the specific resource that is being requested by that client.</li><li>After requested resources are identified, you might need to map extracted resource to a form that is compatible with the authorization method.</li><li>After the service requester is authenticated from the extracted identity, an AAA policy authorizes the client to the requested resource.</li><li>After client-authorization, an AAA policy can run postprocessing activities.</li></ul>",
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
						"authorized_counter": schema.StringAttribute{
							MarkdownDescription: "Specify the monitor for authorized messages. Ensure that this count monitor is configured with an XPath as the measure.",
							Computed:            true,
						},
						"rejected_counter": schema.StringAttribute{
							MarkdownDescription: "Specify the monitor for rejected messages. Ensure that this count monitor is configured with an XPath as the measure.",
							Computed:            true,
						},
						"namespace_mapping": schema.ListNestedAttribute{
							MarkdownDescription: "Define XML namespace maps. Each map is a prefix with its URI.",
							NestedObject:        models.GetDmNamespaceMappingDataSourceSchema(),
							Computed:            true,
						},
						"extract_identity": models.GetDmAAAPExtractIdentityDataSourceSchema("Specify the methods to extract the identity of the service requester. For some methods, you must define more properties.", "extract-identity", ""),
						"authenticate":     models.GetDmAAAPAuthenticateDataSourceSchema("Specify the method to authenticate the extracted identity. For some methods, you must define more properties.", "authenticate", ""),
						"map_credentials":  models.GetDmAAAPMapCredentialsDataSourceSchema("Specify the method to map credentials for authorization. For some methods, you must define more properties.", "map-credentials", ""),
						"extract_resource": models.GetDmAAAPExtractResourceDataSourceSchema("Specify the methods to extract the identity of a requested resource. For some methods, you must define more properties.", "extract-resource", ""),
						"map_resource":     models.GetDmAAAPMapResourceDataSourceSchema("Specify the method to map resources for authorization. For some methods, you must define more properties.", "map-resource", ""),
						"authorize":        models.GetDmAAAPAuthorizeDataSourceSchema("Specify the method to authorize the identity to resources. For some methods, you must define more properties.", "authorize", ""),
						"post_process":     models.GetDmAAAPPostProcessDataSourceSchema("Specify postprocessing activities. For some methods, you must define more properties.", "post-process", ""),
						"saml_attribute": schema.ListNestedAttribute{
							MarkdownDescription: "Specify SAML attributes. Each attribute consists of its namespace URI, local name, and expected value.",
							NestedObject:        models.GetDmSAMLAttributeNameAndValueDataSourceSchema(),
							Computed:            true,
						},
						"ltpa_attributes": schema.ListNestedAttribute{
							MarkdownDescription: "<p>Specify user attributes to include in the LTPA token. Attributes are relevant for only WebSphere tokens.</p><p>For each attribute, its value can be static or resolved at run time. <ul><li>When static, its value is a fixed value.</li><li>When resolved at run time, its value is resolved at run time with an XPath expression.</li></ul></p>",
							NestedObject:        models.GetDmLTPAUserAttributeNameAndValueDataSourceSchema(),
							Computed:            true,
						},
						"transaction_priority": schema.ListNestedAttribute{
							MarkdownDescription: "Define the transactional priority for users. For each user, you must specify the name of the output credentials, their priority for scheduling or resource allocation, and whether authorization is required.",
							NestedObject:        models.GetDmAAATransactionPriorityDataSourceSchema(),
							Computed:            true,
						},
						"saml_valcred": schema.StringAttribute{
							MarkdownDescription: "SAML signature validation credentials",
							Computed:            true,
						},
						"saml_signing_key": schema.StringAttribute{
							MarkdownDescription: "Specify the key to sign SAML messages. To sign messages, you must specify a key and a certificate.",
							Computed:            true,
						},
						"saml_signing_cert": schema.StringAttribute{
							MarkdownDescription: "Specify the certificate to sign SAML messages. To sign messages, you must specify a key and a certificate.",
							Computed:            true,
						},
						"saml_signing_hash_alg": schema.StringAttribute{
							MarkdownDescription: "Specify the algorithm to calculate the message digest for signing. The default value is sha1.",
							Computed:            true,
						},
						"saml_signing_alg": schema.StringAttribute{
							MarkdownDescription: "Specify the algorithm to sign SAML messages. The default value is rsa.",
							Computed:            true,
						},
						"ldap_suffix": schema.StringAttribute{
							MarkdownDescription: "Specify the LDAP suffix to add to the username to form the base distinguished name (DN) for authentication. The suffix and the username are separated with a comma. If the suffix is <tt>O=example.com</tt> and the username is <tt>Bob</tt> , the DN is <tt>CN=Bob,O=example.com</tt> .",
							Computed:            true,
						},
						"log_allowed": schema.BoolAttribute{
							MarkdownDescription: "Whether to log successful AAA operations. When enabled and if needed, modify the default logging level from informational.",
							Computed:            true,
						},
						"log_allowed_level": schema.StringAttribute{
							MarkdownDescription: "Set the logging level for successful AAA operations. The default level is informational.",
							Computed:            true,
						},
						"log_rejected": schema.BoolAttribute{
							MarkdownDescription: "Whether to log unsuccessful AAA operations. When enabled and if needed, modify the default logging level from warning.",
							Computed:            true,
						},
						"log_rejected_level": schema.StringAttribute{
							MarkdownDescription: "Set the logging level for unsuccessful AAA operations. The default level is warning.",
							Computed:            true,
						},
						"ws_secure_conversation_crypto_key": schema.StringAttribute{
							MarkdownDescription: "WS-Trust encryption recipient certificate",
							Computed:            true,
						},
						"saml_source_id_mapping_file": schema.StringAttribute{
							MarkdownDescription: "Specify the location of the SAML artifact-mapping file that provides a mapping of SAML artifact source IDs to artifact retrieval endpoints. Use this setting when artifacts are retrieved from multiple endpoints and the source ID for these endpoints are encoded in the artifact itself. If only one artifact retrieval URL exists, it can be specified by the SAML artifact responder URL in the authentication phase.",
							Computed:            true,
						},
						"ping_identity_compatibility": schema.BoolAttribute{
							MarkdownDescription: "Control compatibility with a PingFederate identity server. By default, compatibility is disabled. Enable compatibility for SAML authentication or authorization.",
							Computed:            true,
						},
						"saml2_metadata_file": schema.StringAttribute{
							MarkdownDescription: "Specify the location of the SAML 2.0 metadata file for SAML 2.0 protocol message exchanges. The metadata in this file identifies identity provider endpoints and certificates to secure message exchanges. The file must have a root-level &lt;md:EntitiesDescriptor> element with an &lt;EntityDescriptor> child element for each identity provider.",
							Computed:            true,
						},
						"dos_valve": schema.Int64Attribute{
							MarkdownDescription: "<p>Specify the number of times to process the same request to protect against a denial of service (DoS) attack. Enter a value in the range 1 - 1000. The default value is 3.</p><p>With the default value, AAA processes only the first 3 signature and each signature can contain up to 3 reference URIs. Additional signatures or reference URIs are ignored.</p><p>XML processing includes encryption, decryption, message signing, and signature validation. The AAA policy supports only identity extraction with subject DN from certificate in message signature and authorization with signer certificate for digitally signed messages.</p>",
							Computed:            true,
						},
						"ldap_version": schema.StringAttribute{
							MarkdownDescription: "Specify the LDAP version to access the LDAP server. The default value is v2.",
							Computed:            true,
						},
						"enforce_soap_actor": schema.BoolAttribute{
							MarkdownDescription: "Whether to enforce the <tt>S11:actor</tt> or <tt>S12:role</tt> on WS-Security messages. In general, a WS-Security message has a <tt>S11:actor</tt> or <tt>S12:role</tt> attribute for its <tt>Security</tt> header. Processing can enforce these attributes when the AAA policy tries to use the <tt>Security</tt> header. For example, the <tt>Security</tt> element has only one actor or role. In this case, the AAA policy processes only the <tt>Security</tt> header for this actor or role identifier. This setting applies to all AAA phases, except postprocessing. For postprocessing, the activity generally generates a new message for next SOAP node.",
							Computed:            true,
						},
						"ws_sec_actor_role_id": schema.StringAttribute{
							MarkdownDescription: "Set the assumed <tt>S11:actor</tt> or <tt>S12:role</tt> identifier. The AAA policy acts as the assumed actor or role when it consumes <tt>Security</tt> headers. This setting takes effect only when the AAA policy attempts to process the incoming message before it makes an authorization decision. Postprocessing does not use this setting. Postprocessing uses its own setting in generating the message for the next SOAP node. The default value is an empty string. <table border=\"1\"><tr><td valign=\"left\"><tt>http://schemas.xmlsoap.org/soap/actor/next</tt></td><td>Every one, including the intermediary and ultimate receiver, that receives the message can process the <tt>Security</tt> header.</td></tr><tr><td valign=\"left\"><tt>http://www.w3.org/2003/05/soap-envelope/role/none</tt></td><td>No one can process the <tt>Security</tt> header.</td></tr><tr><td valign=\"left\"><tt>http://www.w3.org/2003/05/soap-envelope/role/next</tt></td><td>Every one, including the intermediary and ultimate receiver, that receives the message can process the <tt>Security</tt> header.</td></tr><tr><td valign=\"left\"><tt>http://www.w3.org/2003/05/soap-envelope/role/ultimateReceiver</tt></td><td>The ultimate receiver can process the <tt>Security</tt> header.</td></tr><tr><td valign=\"left\">No value, which is an empty string</td><td>The empty string (without quotation marks) indicates that no \"actor/role\" identifier is configured. With a configured actor/role, the ultimate receiver is assumed for the message. No actor/role attribute is added during the generation of the <tt>Security</tt> header. More than one <tt>Security</tt> header cannot omit the actor/role identifier.</td></tr><tr><td valign=\"left\"><tt>USE_MESSAGE_BASE_URI</tt></td><td>The identifier is the base URL of the message. When the SOAP message is transported, the base URI is the request-URI of the HTTP request.</td></tr><tr><td valign=\"left\">A string value</td><td>Any string to identify the actor or role of the <tt>Security</tt> header.</td></tr></table>",
							Computed:            true,
						},
						"au_sm_http_header": schema.ListAttribute{
							MarkdownDescription: "Specify HTTP headers from CA Single Sign-On authentication responses. These headers are included as request or response headers based on the CA Single Sign-on header flow.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"az_sm_http_header": schema.ListAttribute{
							MarkdownDescription: "Specify HTTP headers from CA Single Sign-On authorization responses. These headers are included as request or response headers based on the CA Single Sign-On header flow.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"dyn_config": schema.StringAttribute{
							MarkdownDescription: "Specify how to obtain the AAA policy configuration dynamically. With dynamic configuration, you can configure the AAA policy at run time. When enabled, the configuration of AAA is determined dynamically based on the template AAA policy and the parameters that the dynamic configuration custom URL returns. By default, uses no template.",
							Computed:            true,
						},
						"external_aaa_template": schema.StringAttribute{
							MarkdownDescription: "Specify another AAA policy to use as the template. When specified, this AAA policy overwrites the current AAA policy.",
							Computed:            true,
						},
						"dyn_config_custom_url": schema.StringAttribute{
							MarkdownDescription: "<p>Specify the location of the custom stylesheet or GatewayScript file. The configuration of the AAA policy is obtained dynamically from this file. The obtained configuration overwrites the configuration in the template AAA policy.</p><p>In the custom file, modify only the properties to dynamically overwrite. See the <tt>ModifyAAAPolicy</tt> element in the <tt>store:///xml-mgmt.xsd</tt> schema to construct a schema-compliant AAA configuration.</p>",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *AAAPolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *AAAPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data AAAPolicyList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.AAAPolicy{
		AppDomain: data.AppDomain,
	}

	path := o.GetPath()
	if !data.Id.IsNull() {
		path = path + "/" + data.Id.ValueString()
	}

	res, err := d.pData.Client.Get(path)
	resFound := true
	if err != nil {
		if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(d.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString())
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
	l := []models.AAAPolicy{}
	if resFound {
		if value := res.Get(`AAAPolicy`); value.Exists() {
			for _, v := range value.Array() {
				item := models.AAAPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.AAAPolicyObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.AAAPolicyObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
