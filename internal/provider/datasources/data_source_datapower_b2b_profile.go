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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type B2BProfileList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &B2BProfileDataSource{}
	_ datasource.DataSourceWithConfigure = &B2BProfileDataSource{}
)

func NewB2BProfileDataSource() datasource.DataSource {
	return &B2BProfileDataSource{}
}

type B2BProfileDataSource struct {
	pData *tfutils.ProviderData
}

func (d *B2BProfileDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_b2b_profile"
}

func (d *B2BProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "B2B partner profile",
		Attributes: map[string]schema.Attribute{
			"app_domain": schema.StringAttribute{
				MarkdownDescription: "The name of the application domain the object belongs to",
				Required:            true,
			},
			"result": schema.ListNestedAttribute{
				MarkdownDescription: "List of objects",
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
						"profile_type": schema.StringAttribute{
							MarkdownDescription: "Specify whether the profile is for an internal or external partner. The default value is internal.",
							Computed:            true,
						},
						"business_i_ds": schema.ListAttribute{
							MarkdownDescription: "Specify the identifier for the partners. When configuring a trading partner, its identifier (ID) must be unique within a specific B2B gateway. The identifiers are equivalent to one of the following values. <ul><li>In AS messages, a value for an <tt>AS*-From</tt> or <tt>AS*-To</tt> header, where * is 1, 2, or 3.</li><li>In ebMS messages, a value for an <tt>PartyID</tt> element, where the PartyID element can be under either From or To element.</li><li>In other messages, a value that is extracted from the body of the EDI message: <tt>SenderID</tt> or <tt>ReceiverID</tt> .</li></ul>",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"business_i_ds_duns": schema.ListAttribute{
							MarkdownDescription: "Specify the 9-digit DUNS (Data Universal Numbering System) identification number for the partner. When configuring a trading partner, the identifier (ID) must be unique not only within the 3 types of ID System (Freeform, DUNS, and DUNS+4) but also within a specific B2B gateway.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"business_i_ds_duns_plus4": schema.ListAttribute{
							MarkdownDescription: "Specifies the 13-digit D-U-N-S (Data Universal Numbering System + 4) identification number for the partner. When configuring a trading partner, the identifier (ID) must be unique not only within the 3 types of ID System (Freeform, DUNS, and DUNS+4) but also within a specific B2B gateway.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"custom_style_policy": schema.StringAttribute{
							MarkdownDescription: "Processing policy",
							Computed:            true,
						},
						"response_type": schema.StringAttribute{
							MarkdownDescription: "Characterize the traffic that originates from <tt>To</tt> partners. For inbound transaction, this is the traffic type that originates from internal partner. For outbound transaction, this is the traffic type that originates from external partner. The default value is Non-XML.",
							Computed:            true,
						},
						"email_addresses": schema.ListAttribute{
							MarkdownDescription: "Specify the identifier (e-mail address) for the partner. When configuring a partner, their email address must be unique within a specific B2B gateway configuration. The email addresses for the partner must be equivalent to one of the following values. <ul><li>In AS1 messages, a value for a <tt>From</tt> or <tt>To</tt> header</li><li>In non-AS1 messages, a value that is extracted from the body of the EDI message - <tt>SenderID</tt> or <tt>ReceiverID</tt> .</li></ul>",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"destinations": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the routing information for the partner. The first destination is the default destination. The gateway uses the default destination when no specific destination is assigned or when no matching destination is found.",
							NestedObject:        models.GetDmB2BDestinationDataSourceSchema(),
							Computed:            true,
						},
						"inbound_verify_val_cred": schema.StringAttribute{
							MarkdownDescription: "Inbound signature validation credentials",
							Computed:            true,
						},
						"inbound_require_signed": schema.BoolAttribute{
							MarkdownDescription: "Specify whether inbound AS messages must be signed. The default behavior is off.",
							Computed:            true,
						},
						"inbound_require_encrypted": schema.BoolAttribute{
							MarkdownDescription: "Specify whether inbound AS messages must be encrypted. The default behavior is off.",
							Computed:            true,
						},
						"inbound_decrypt_id_cred": schema.StringAttribute{
							MarkdownDescription: "Inbound decryption identification credentials",
							Computed:            true,
						},
						"outbound_sign": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to sign outbound messages. The default behavior is off. <ul><li>If enabled, sign outbound messages with the configured identification credentials and algorithm. If the destination indicates to send messages unsigned, messages from this partner to that destination are not signed.</li><li>If disabled, does not sign outbound messages.</li></ul><p>This setting has no effect on outbound MDN messages. Regardless of this setting and if a partner requests a signed MDN, the outbound MDN is signed if this partner has a configured identification credentials.</p>",
							Computed:            true,
						},
						"outbound_sign_id_cred": schema.StringAttribute{
							MarkdownDescription: "Specify the identification credentials to sign an outbound message or outbound MDN. The outbound MDN is signed if requested by a partner and this setting is configured.",
							Computed:            true,
						},
						"outbound_sign_digest_alg": schema.StringAttribute{
							MarkdownDescription: "Signing digest algorithm",
							Computed:            true,
						},
						"outbound_sign_mic_alg_version": schema.StringAttribute{
							MarkdownDescription: "Specify the S/MIME specification version to generate the micalg parameter value in a multipart/signed message. The following value is a sample <tt>Content-Type</tt> header of the multipart/signed message. <p><tt>Content-Type: multipart/signed;protocol=\"application/pkcs7-signature\";micalg=sha1; boundary=...</tt></p><p>The micalg parameter indicates which message digest algorithm (such as, MD5, SHA-1, SHA-256, SHA-384, and SHA-512) to use in the calculation of the Message Integrity Check (MIC). The formats of the micalg value differ between S/MIME Version 3.1 (RFC 3851) and Version 3.2 (RFC 5751).</p><ul><li>In Version 3.1, the micalg parameter value has micalg=[md5|sha1|sha256|sha384|sha512].</li><li>In Version 3.2, the micalg parameter value has micalg=[md5|sha-1|sha-256|sha-384|sha-512].</li></ul>",
							Computed:            true,
						},
						"contacts": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the contact information for partners. To define the contact information, you must provide information for at least one property.",
							NestedObject:        models.GetDmB2BContactDataSourceSchema(),
							Computed:            true,
						},
						"override_asid": schema.StringAttribute{
							MarkdownDescription: "Specify the override the identifiers in AS message headers. <ul><li>For outbound messages, the value to use for the <tt>AS*-From</tt> or <tt>AS*-To</tt> header, where * is 1, 2, or 3. If blank, the transaction retains the identifiers that were extracted from the message payload.</li><li>For inbound messages, the value becomes another valid identifier for the partner in addition to those already defined for the partner.</li></ul>",
							Computed:            true,
						},
						"as_allow_duplicate_message": schema.StringAttribute{
							MarkdownDescription: "Specify when to allow and reprocess duplicate AS inbound messages. The default behavior is never. This option does not apply to MDN.",
							Computed:            true,
						},
						"preserve_filename": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to expose the <tt>Content-Disposition</tt> header of inbound AS messages for file name preservation. According to RFC 2183, the file name information is optional in <tt>Content-Disposition</tt> header. When enabled, the <tt>Content-Disposition</tt> header of the inbound AS message is exposed if the inbound AS message is in S/MIME format. Then, the receiving partner can retrieve the original file name that is contained in the header and transfer the received file to its backend system with the received file name.",
							Computed:            true,
						},
						"ebms_role": schema.StringAttribute{
							MarkdownDescription: "Specify the name of authorized role of the party. When sending outbound message, the role associated with internal partner presents the <tt>From</tt> party and the role associated with external partner presents the <tt>To</tt> party. The value is referencing the <tt>Role</tt> specified in CPA. A <tt>Role</tt> is better defined as a URI, for example, http://rosettanet.org/roles/buyer.",
							Computed:            true,
						},
						"ebms_persist_duration": schema.Int64Attribute{
							MarkdownDescription: "Specify the duration in seconds to retain messages in persistent storage. This value is used to compute the <tt>TimeToLive</tt> value. Until the value of the <tt>TimeToLive</tt> element elapses, the message cannot be archived.",
							Computed:            true,
						},
						"ebms_ack_url": schema.StringAttribute{
							MarkdownDescription: "Specify the redirection URL to send the acknowledgment message when received ebMS2 message requests an asynchronous response. When an asynchronous reply is requested by the inbound ebMS2 document, this field is required to determine where to send the acknowledgment.",
							Computed:            true,
						},
						"ebms_error_url": schema.StringAttribute{
							MarkdownDescription: "Specify the redirection URL to send the error message when received ebMS2 message requests an asynchronous response. When an asynchronous reply is requested by the inbound ebMS2 document, this field is used as the error reporting location to send the error message that contains the error code and the description of the problem. Error URL cannot be empty when the acknowledgment URL is specified.",
							Computed:            true,
						},
						"ebms_inbound_send_receipt": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to reply to the MSH with a receipt signal message for the received ebMS messages. The default behavior is off.",
							Computed:            true,
						},
						"ebms_inbound_send_signed_receipt": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to reply to the MSH with a signed receipt signal message for the received ebMS3 message. The default behavior is off.",
							Computed:            true,
						},
						"ebms_inbound_receipt_reply_pattern": schema.StringAttribute{
							MarkdownDescription: "Specify the pattern to send the receipt signal. The default behavior is response.",
							Computed:            true,
						},
						"ebms_receipt_url": schema.StringAttribute{
							MarkdownDescription: "When the receipt reply pattern is callback, specify the URL to send the receipt signal.",
							Computed:            true,
						},
						"ebms_inbound_error_url": schema.StringAttribute{
							MarkdownDescription: "When the error is sent asynchronously, specify the URL to send the error signal.",
							Computed:            true,
						},
						"ebms_inbound_verify_val_cred": schema.StringAttribute{
							MarkdownDescription: "Specify the validation credentials to verify the signature on an acknowledgment or inbound ebMS2 message. For ebMS2 messages, only <tt>X509Data</tt> and <tt>KeyName</tt> signature methods are supported.",
							Computed:            true,
						},
						"ebms_default_signer_cert": schema.StringAttribute{
							MarkdownDescription: "Specify the default validation certificate to verify the signature of an inbound ebMS2 message or an acknowledgment. The default signature validation certificate is used if either the <tt>keyInfo</tt> element is missing or signature method is unsupported.",
							Computed:            true,
						},
						"ebms_inbound_require_signed": schema.BoolAttribute{
							MarkdownDescription: "Specify whether inbound ebMS2 messages must be signed. The default behavior is off.",
							Computed:            true,
						},
						"ebms_inbound_require_encrypted": schema.BoolAttribute{
							MarkdownDescription: "Specify whether inbound ebMS2 messages must be encrypted. The default behavior is off.",
							Computed:            true,
						},
						"ebms_inbound_decrypt_id_cred": schema.StringAttribute{
							MarkdownDescription: "Specify the identification credentials to decrypt inbound ebMS2 messages.",
							Computed:            true,
						},
						"ebms_outbound_sign": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to sign outbound messages. The default behavior is disabled. <ul><li>If enabled, signs outbound messages with the configured identification credentials and algorithm. If the configuration of a destination indicates to send messages unsigned, messages from this partner to that destination are not signed.</li><li>If disabled, does not sign outbound messages.</li></ul><p>This setting has no effect on outbound acknowledgment messages. Regardless of this setting and if a partner requests a signed acknowledgment, the outbound acknowledgment is signed if this partner has a configured identification credentials.</p>",
							Computed:            true,
						},
						"ebms_outbound_sign_id_cred": schema.StringAttribute{
							MarkdownDescription: "Specify the identification credentials to sign an outbound message or acknowledgment. The outbound acknowledgment is signed if requested by a partner and this setting is configured. For ebMS2 messages, only <tt>X509Data</tt> and <tt>KeyName</tt> signature methods are supported.",
							Computed:            true,
						},
						"ebms_outbound_signature_alg": schema.StringAttribute{
							MarkdownDescription: "Specify the algorithm to sign the outbound ebMS2 message. The default value is dsa-sha1, which is recommended by the ebMS specification.",
							Computed:            true,
						},
						"ebms_outbound_signature_c14n_alg": schema.StringAttribute{
							MarkdownDescription: "Specify the algorithm to canonicalize the SOAP Envelope XML and exclude comments before signing outbound ebMS2 message. The default value is c14n, which is recommended by the ebMS specification.",
							Computed:            true,
						},
						"ebms_outbound_sign_digest_alg": schema.StringAttribute{
							MarkdownDescription: "Specify the algorithm to hash the payload during signing. The default value is sha1. See http://www.w3.org/TR/xmldsig-core/",
							Computed:            true,
						},
						"ebms_enable_cpa_binding": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enable CPA bindings. When enabled, the CPA, service, and action that are specified by the matched CPA binding are used for the outbound ebMS2 messages instead of the default CPA, service, and action of the external profile. The CPA binding is matched by the internal partner profile.",
							Computed:            true,
						},
						"ebms_profile_cpa_bindings": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the CPA binding for the external partner profile. A CPA binding binds a CPA entry (CPA, service, and action) that you prefer to use. When a CPA entry is matched through the internal partner profile, outbound messages from the internal partner use the CPA, service, and action that are specified by the matched CPA entry. You must have the CPA entry in the ebXML settings of the associated B2B gateway before you add the entry to the CPA bindings list.",
							NestedObject:        models.GetDmProfileCPABindingDataSourceSchema(),
							Computed:            true,
						},
						"ebms_cpa_id": schema.StringAttribute{
							MarkdownDescription: "Specify the CPA ID to package in the outbound ebMS2 message header. This value is used when the B2B gateway is CPA-enforced.",
							Computed:            true,
						},
						"ebms_service": schema.StringAttribute{
							MarkdownDescription: "Specify the service to package in the outbound ebMS2 message header. This value is used when when the B2B gateway is CPA-enforced.",
							Computed:            true,
						},
						"ebms_action": schema.StringAttribute{
							MarkdownDescription: "Specify the action to package in the outbound ebMS2 message header. This value is used when the B2B gateway is CPA-enforced.",
							Computed:            true,
						},
						"ebms_start_parameter": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to generate a start parameter for the ebMS2 message. The start parameter identifies the root part of the ebMS message. This setting is disabled by default. If enabled, the start parameter is generated with this value enclosed in angle brackets.",
							Computed:            true,
						},
						"ebms_allow_duplicate_message": schema.StringAttribute{
							MarkdownDescription: "Specify when to allow and reprocess duplicate ebMS2 inbound messages. The default behavior is never. This option does not apply to acknowledgments.",
							Computed:            true,
						},
						"mdnssl_client_config_type": schema.StringAttribute{
							MarkdownDescription: "MDN TLS client type",
							Computed:            true,
						},
						"mdnssl_client": schema.StringAttribute{
							MarkdownDescription: "MDN TLS client profile",
							Computed:            true,
						},
						"ebms_ack_ssl_client_config_type": schema.StringAttribute{
							MarkdownDescription: "Acknowledgment/Error TLS client type",
							Computed:            true,
						},
						"ebms_ack_ssl_client": schema.StringAttribute{
							MarkdownDescription: "Acknowledgment/Error TLS client profile",
							Computed:            true,
						},
						"ebms3_outbound_sign": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to sign outbound messages. The default behavior is off. <ul><li>When enabled, signs outbound messages with the configured identification credentials and algorithm. If the configuration of a destination indicates to send messages unsigned, messages from this partner to that destination are not signed.</li><li>When disabled, does not sign outbound messages.</li></ul><p>This setting has no effect on outbound receipt signal messages. Regardless of this setting and whether an external partner requests a signed receipt signal, the outbound receipt signal is signed when this internal partner has a configured identification credentials.</p>",
							Computed:            true,
						},
						"ebms3_outbound_sign_id_cred": schema.StringAttribute{
							MarkdownDescription: "Specify the identification credentials to sign an outbound message or outbound receipt signal. The outbound receipt signal is signed if requested by a partner and this setting is configured.",
							Computed:            true,
						},
						"ebms3_outbound_sign_digest_alg": schema.StringAttribute{
							MarkdownDescription: "Specify the algorithm to hash the payload during signing. The default value is sha1. See http://www.w3.org/TR/xmldsig-core/",
							Computed:            true,
						},
						"ebms3_outbound_signature_alg": schema.StringAttribute{
							MarkdownDescription: "Specify the algorithm to sign outbound ebMS3 messages. The default value is dsa-sha1, which is recommended by the ebMS specification.",
							Computed:            true,
						},
						"ebms3_outbound_signature_c14n_alg": schema.StringAttribute{
							MarkdownDescription: "Specify the algorithm to canonicalize the SOAP Envelope XML and exclude comments before signing outbound ebMS3 messages. The default value is exc-c14n, which is recommended by the ebMS specification.",
							Computed:            true,
						},
						"ebms3_inbound_verify_val_cred": schema.StringAttribute{
							MarkdownDescription: "Inbound signature validation credentials",
							Computed:            true,
						},
						"ebms3_default_signer_cert": schema.StringAttribute{
							MarkdownDescription: "Specify the default validation certificate to verify the signature of an inbound ebMS3 message or a receipt. This certificate is used when the <tt>keyInfo</tt> element is missing or the signature method is unsupported.",
							Computed:            true,
						},
						"ebms3_inbound_require_signed": schema.BoolAttribute{
							MarkdownDescription: "Specify whether inbound ebMS3 messages must be signed. The default behavior is off.",
							Computed:            true,
						},
						"ebms3_inbound_require_encrypted": schema.BoolAttribute{
							MarkdownDescription: "Specify whether inbound ebMS3 messages must be encrypted. The default behavior is off.",
							Computed:            true,
						},
						"ebms3_inbound_decrypt_id_cred": schema.StringAttribute{
							MarkdownDescription: "Inbound decryption identification credentials",
							Computed:            true,
						},
						"ebms3_inbound_require_compressed": schema.BoolAttribute{
							MarkdownDescription: "Specify whether internal partners require compressed inbound ebMS3 messages. The default behavior is not to require compression.",
							Computed:            true,
						},
						"ebms3_receipt_ssl_client_config_type": schema.StringAttribute{
							MarkdownDescription: "Receipt/Error TLS client type",
							Computed:            true,
						},
						"ebms3_receipt_ssl_client": schema.StringAttribute{
							MarkdownDescription: "Receipt/Error TLS client profile",
							Computed:            true,
						},
						"ebms_notification": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enable notification. When enabled, send notifications to the message producer or consumer when there are specific errors. The default behavior is off.",
							Computed:            true,
						},
						"ebms_notification_url": schema.StringAttribute{
							MarkdownDescription: "Notification URL",
							Computed:            true,
						},
						"ebms_notification_ssl_client_config_type": schema.StringAttribute{
							MarkdownDescription: "Notification TLS client type",
							Computed:            true,
						},
						"ebms_notification_ssl_client": schema.StringAttribute{
							MarkdownDescription: "Notification TLS client profile",
							Computed:            true,
						},
						"ebms3_allow_duplicate_message": schema.StringAttribute{
							MarkdownDescription: "Specify when to allow and reprocess duplicate ebMS3 inbound messages. The default behavior is never. This option does not apply to the receipt signal messages.",
							Computed:            true,
						},
						"ebms3_duplicate_detection_notification": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enable duplicate detection notification. When enabled, send notification to the message consumer when a duplicate message is detected. The default behavior is disabled.",
							Computed:            true,
						},
						"ebms_message_properties": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the <tt>eb:Property</tt> elements to add to the <tt>eb:MessageProperties</tt> node. Define message properties to meet your business needs and the agreement between partners.",
							NestedObject:        models.GetDmB2BMessagePropertiesDataSourceSchema(),
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *B2BProfileDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *B2BProfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data B2BProfileList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.B2BProfile{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.B2BProfile{}
	if value := res.Get(`B2BProfile`); value.Exists() {
		for _, v := range value.Array() {
			item := models.B2BProfile{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.B2BProfileObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.B2BProfileObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
