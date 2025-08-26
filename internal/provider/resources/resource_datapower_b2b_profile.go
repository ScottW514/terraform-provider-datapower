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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

var _ resource.Resource = &B2BProfileResource{}

func NewB2BProfileResource() resource.Resource {
	return &B2BProfileResource{}
}

type B2BProfileResource struct {
	pData *tfutils.ProviderData
}

func (r *B2BProfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_b2b_profile"
}

func (r *B2BProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("B2B partner profile", "b2b-profile", "").String,
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
			"profile_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the profile is for an internal or external partner. The default value is internal.", "profile-type", "").AddStringEnum("internal", "external").AddDefaultValue("internal").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("internal", "external"),
				},
				Default: stringdefault.StaticString("internal"),
			},
			"business_ids": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the identifier for the partners. When configuring a trading partner, its identifier (ID) must be unique within a specific B2B gateway. The identifiers are equivalent to one of the following values. <ul><li>In AS messages, a value for an <tt>AS*-From</tt> or <tt>AS*-To</tt> header, where * is 1, 2, or 3.</li><li>In ebMS messages, a value for an <tt>PartyID</tt> element, where the PartyID element can be under either From or To element.</li><li>In other messages, a value that is extracted from the body of the EDI message: <tt>SenderID</tt> or <tt>ReceiverID</tt> .</li></ul>", "business-id", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"business_id_duns": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the 9-digit DUNS (Data Universal Numbering System) identification number for the partner. When configuring a trading partner, the identifier (ID) must be unique not only within the 3 types of ID System (Freeform, DUNS, and DUNS+4) but also within a specific B2B gateway.", "business-id-duns", "").String,
				ElementType:         types.StringType,
				Optional:            true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.RegexMatches(regexp.MustCompile("^[0-9]{9}$"), "Must match :"+"^[0-9]{9}$"),
					),
				},
			},
			"business_id_duns4": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the 13-digit D-U-N-S (Data Universal Numbering System + 4) identification number for the partner. When configuring a trading partner, the identifier (ID) must be unique not only within the 3 types of ID System (Freeform, DUNS, and DUNS+4) but also within a specific B2B gateway.", "business-id-duns4", "").String,
				ElementType:         types.StringType,
				Optional:            true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.RegexMatches(regexp.MustCompile("^[0-9]{13}$"), "Must match :"+"^[0-9]{13}$"),
					),
				},
			},
			"custom_style_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Processing policy", "stylepolicy", "style_policy").String,
				Optional:            true,
			},
			"response_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Characterize the traffic that originates from <tt>To</tt> partners. For inbound transaction, this is the traffic type that originates from internal partner. For outbound transaction, this is the traffic type that originates from external partner. The default value is Non-XML.", "response-type", "").AddStringEnum("soap", "xml", "unprocessed", "preprocessed").AddDefaultValue("preprocessed").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("soap", "xml", "unprocessed", "preprocessed"),
				},
				Default: stringdefault.StaticString("preprocessed"),
			},
			"email_addresses": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the identifier (e-mail address) for the partner. When configuring a partner, their email address must be unique within a specific B2B gateway configuration. The email addresses for the partner must be equivalent to one of the following values. <ul><li>In AS1 messages, a value for a <tt>From</tt> or <tt>To</tt> header</li><li>In non-AS1 messages, a value that is extracted from the body of the EDI message - <tt>SenderID</tt> or <tt>ReceiverID</tt> .</li></ul>", "email-address", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"destinations": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the routing information for the partner. The first destination is the default destination. The gateway uses the default destination when no specific destination is assigned or when no matching destination is found.", "destination", "").String,
				NestedObject:        models.GetDmB2BDestinationResourceSchema(),
				Required:            true,
			},
			"inbound_verify_val_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Inbound signature validation credentials", "verify-valcred", "crypto_val_cred").String,
				Optional:            true,
			},
			"inbound_require_signed": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether inbound AS messages must be signed. The default behavior is off.", "require-signed", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"inbound_require_encrypted": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether inbound AS messages must be encrypted. The default behavior is off.", "require-encrypted", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"inbound_decrypt_id_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Inbound decryption identification credentials", "decrypt-idcred", "crypto_ident_cred").AddRequiredWhen(models.B2BProfileInboundDecryptIdCredCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.B2BProfileInboundDecryptIdCredCondVal, validators.Evaluation{}, false),
				},
			},
			"outbound_sign": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to sign outbound messages. The default behavior is off. <ul><li>If enabled, sign outbound messages with the configured identification credentials and algorithm. If the destination indicates to send messages unsigned, messages from this partner to that destination are not signed.</li><li>If disabled, does not sign outbound messages.</li></ul><p>This setting has no effect on outbound MDN messages. Regardless of this setting and if a partner requests a signed MDN, the outbound MDN is signed if this partner has a configured identification credentials.</p>", "sign", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"outbound_sign_id_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the identification credentials to sign an outbound message or outbound MDN. The outbound MDN is signed if requested by a partner and this setting is configured.", "sign-idcred", "crypto_ident_cred").AddRequiredWhen(models.B2BProfileOutboundSignIdCredCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.B2BProfileOutboundSignIdCredCondVal, validators.Evaluation{}, false),
				},
			},
			"outbound_sign_digest_alg": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Signing digest algorithm", "sign-digest-alg", "").AddStringEnum("sha1", "md5", "sha256", "sha384", "sha512").AddDefaultValue("sha1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("sha1", "md5", "sha256", "sha384", "sha512"),
				},
				Default: stringdefault.StaticString("sha1"),
			},
			"outbound_sign_mic_alg_version": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the S/MIME specification version to generate the micalg parameter value in a multipart/signed message. The following value is a sample <tt>Content-Type</tt> header of the multipart/signed message. <p><tt>Content-Type: multipart/signed;protocol=\"application/pkcs7-signature\";micalg=sha1; boundary=...</tt></p><p>The micalg parameter indicates which message digest algorithm (such as, MD5, SHA-1, SHA-256, SHA-384, and SHA-512) to use in the calculation of the Message Integrity Check (MIC). The formats of the micalg value differ between S/MIME Version 3.1 (RFC 3851) and Version 3.2 (RFC 5751).</p><ul><li>In Version 3.1, the micalg parameter value has micalg=[md5|sha1|sha256|sha384|sha512].</li><li>In Version 3.2, the micalg parameter value has micalg=[md5|sha-1|sha-256|sha-384|sha-512].</li></ul>", "sign-micalg-version", "").AddStringEnum("SMIME3.1", "SMIME3.2").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("SMIME3.1", "SMIME3.2"),
				},
			},
			"contacts": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the contact information for partners. To define the contact information, you must provide information for at least one property.", "contact", "").String,
				NestedObject:        models.GetDmB2BContactResourceSchema(),
				Optional:            true,
			},
			"override_asid": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the override the identifiers in AS message headers. <ul><li>For outbound messages, the value to use for the <tt>AS*-From</tt> or <tt>AS*-To</tt> header, where * is 1, 2, or 3. If blank, the transaction retains the identifiers that were extracted from the message payload.</li><li>For inbound messages, the value becomes another valid identifier for the partner in addition to those already defined for the partner.</li></ul>", "override-as-identifier", "").String,
				Optional:            true,
			},
			"as_allow_duplicate_message": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify when to allow and reprocess duplicate AS inbound messages. The default behavior is never. This option does not apply to MDN.", "as-allow-dup-msg", "").AddStringEnum("never", "always", "on-error").AddDefaultValue("never").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("never", "always", "on-error"),
				},
				Default: stringdefault.StaticString("never"),
			},
			"preserve_filename": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to expose the <tt>Content-Disposition</tt> header of inbound AS messages for file name preservation. According to RFC 2183, the file name information is optional in <tt>Content-Disposition</tt> header. When enabled, the <tt>Content-Disposition</tt> header of the inbound AS message is exposed if the inbound AS message is in S/MIME format. Then, the receiving partner can retrieve the original file name that is contained in the header and transfer the received file to its backend system with the received file name.", "preserve-filename", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ebms_role": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of authorized role of the party. When sending outbound message, the role associated with internal partner presents the <tt>From</tt> party and the role associated with external partner presents the <tt>To</tt> party. The value is referencing the <tt>Role</tt> specified in CPA. A <tt>Role</tt> is better defined as a URI, for example, http://rosettanet.org/roles/buyer.", "ebms-role", "").String,
				Optional:            true,
			},
			"ebms_persist_duration": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration in seconds to retain messages in persistent storage. This value is used to compute the <tt>TimeToLive</tt> value. Until the value of the <tt>TimeToLive</tt> element elapses, the message cannot be archived.", "ebms-persist-duration", "").AddIntegerRange(0, 6000000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 6000000),
				},
			},
			"ebms_ack_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the redirection URL to send the acknowledgment message when received ebMS2 message requests an asynchronous response. When an asynchronous reply is requested by the inbound ebMS2 document, this field is required to determine where to send the acknowledgment.", "ebms-ack-url", "").String,
				Optional:            true,
			},
			"ebms_error_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the redirection URL to send the error message when received ebMS2 message requests an asynchronous response. When an asynchronous reply is requested by the inbound ebMS2 document, this field is used as the error reporting location to send the error message that contains the error code and the description of the problem. Error URL cannot be empty when the acknowledgment URL is specified.", "ebms-error-url", "").AddRequiredWhen(models.B2BProfileEBMSErrorURLCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.B2BProfileEBMSErrorURLCondVal, validators.Evaluation{}, false),
				},
			},
			"ebms_inbound_send_receipt": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to reply to the MSH with a receipt signal message for the received ebMS messages. The default behavior is off.", "ebms-inbound-send-receipt", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ebms_inbound_send_signed_receipt": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to reply to the MSH with a signed receipt signal message for the received ebMS3 message. The default behavior is off.", "ebms-inbound-send-signed-receipt", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ebms_inbound_receipt_reply_pattern": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the pattern to send the receipt signal. The default behavior is response.", "ebms-inbound-receipt-reply-pattern", "").AddStringEnum("Response", "Callback").AddDefaultValue("Response").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Response", "Callback"),
				},
				Default: stringdefault.StaticString("Response"),
			},
			"ebms_receipt_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("When the receipt reply pattern is callback, specify the URL to send the receipt signal.", "ebms-receipt-url", "").AddRequiredWhen(models.B2BProfileEBMSReceiptURLCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.B2BProfileEBMSReceiptURLCondVal, validators.Evaluation{}, false),
				},
			},
			"ebms_inbound_error_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("When the error is sent asynchronously, specify the URL to send the error signal.", "ebms-inbound-error-url", "").AddRequiredWhen(models.B2BProfileEBMSInboundErrorURLCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.B2BProfileEBMSInboundErrorURLCondVal, validators.Evaluation{}, false),
				},
			},
			"ebms_inbound_verify_val_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the validation credentials to verify the signature on an acknowledgment or inbound ebMS2 message. For ebMS2 messages, only <tt>X509Data</tt> and <tt>KeyName</tt> signature methods are supported.", "ebms-verify-valcred", "crypto_val_cred").String,
				Optional:            true,
			},
			"ebms_default_signer_cert": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the default validation certificate to verify the signature of an inbound ebMS2 message or an acknowledgment. The default signature validation certificate is used if either the <tt>keyInfo</tt> element is missing or signature method is unsupported.", "ebms-default-signer-cert", "crypto_certificate").String,
				Optional:            true,
			},
			"ebms_inbound_require_signed": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether inbound ebMS2 messages must be signed. The default behavior is off.", "ebms-require-signed", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ebms_inbound_require_encrypted": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether inbound ebMS2 messages must be encrypted. The default behavior is off.", "ebms-require-encrypted", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ebms_inbound_decrypt_id_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the identification credentials to decrypt inbound ebMS2 messages.", "ebms-decrypt-idcred", "crypto_ident_cred").AddRequiredWhen(models.B2BProfileEBMSInboundDecryptIdCredCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.B2BProfileEBMSInboundDecryptIdCredCondVal, validators.Evaluation{}, false),
				},
			},
			"ebms_outbound_sign": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to sign outbound messages. The default behavior is disabled. <ul><li>If enabled, signs outbound messages with the configured identification credentials and algorithm. If the configuration of a destination indicates to send messages unsigned, messages from this partner to that destination are not signed.</li><li>If disabled, does not sign outbound messages.</li></ul><p>This setting has no effect on outbound acknowledgment messages. Regardless of this setting and if a partner requests a signed acknowledgment, the outbound acknowledgment is signed if this partner has a configured identification credentials.</p>", "ebms-sign", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ebms_outbound_sign_id_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the identification credentials to sign an outbound message or acknowledgment. The outbound acknowledgment is signed if requested by a partner and this setting is configured. For ebMS2 messages, only <tt>X509Data</tt> and <tt>KeyName</tt> signature methods are supported.", "ebms-sign-idcred", "crypto_ident_cred").AddRequiredWhen(models.B2BProfileEBMSOutboundSignIdCredCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.B2BProfileEBMSOutboundSignIdCredCondVal, validators.Evaluation{}, false),
				},
			},
			"ebms_outbound_signature_alg": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the algorithm to sign the outbound ebMS2 message. The default value is dsa-sha1, which is recommended by the ebMS specification.", "ebms-signature-alg", "").AddStringEnum("dsa-sha1", "rsa-sha1", "rsa-sha256", "rsa-sha384", "rsa-sha512", "rsa-ripemd160", "rsa-ripemd160-2010", "sha256-rsa-MGF1", "rsa-md5", "ecdsa-sha1", "ecdsa-sha224", "ecdsa-sha256", "ecdsa-sha384", "ecdsa-sha512").AddDefaultValue("dsa-sha1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("dsa-sha1", "rsa-sha1", "rsa-sha256", "rsa-sha384", "rsa-sha512", "rsa-ripemd160", "rsa-ripemd160-2010", "sha256-rsa-MGF1", "rsa-md5", "ecdsa-sha1", "ecdsa-sha224", "ecdsa-sha256", "ecdsa-sha384", "ecdsa-sha512"),
				},
				Default: stringdefault.StaticString("dsa-sha1"),
			},
			"ebms_outbound_signature_c14n_alg": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the algorithm to canonicalize the SOAP Envelope XML and exclude comments before signing outbound ebMS2 message. The default value is c14n, which is recommended by the ebMS specification.", "ebms-signature-c14n-alg", "").AddStringEnum("c14n", "exc-c14n", "c14n-comments", "exc-c14n-comments", "c14n11", "c14n11-comments").AddDefaultValue("c14n").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("c14n", "exc-c14n", "c14n-comments", "exc-c14n-comments", "c14n11", "c14n11-comments"),
				},
				Default: stringdefault.StaticString("c14n"),
			},
			"ebms_outbound_sign_digest_alg": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the algorithm to hash the payload during signing. The default value is sha1. See http://www.w3.org/TR/xmldsig-core/", "ebms-sign-digest-alg", "").AddStringEnum("sha1", "sha256", "sha512", "ripemd160", "sha224", "sha384", "md5").AddDefaultValue("sha1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("sha1", "sha256", "sha512", "ripemd160", "sha224", "sha384", "md5"),
				},
				Default: stringdefault.StaticString("sha1"),
			},
			"ebms_enable_cpa_binding": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable CPA bindings. When enabled, the CPA, service, and action that are specified by the matched CPA binding are used for the outbound ebMS2 messages instead of the default CPA, service, and action of the external profile. The CPA binding is matched by the internal partner profile.", "ebms-enable-cpa-bindings", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ebms_profile_cpa_bindings": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the CPA binding for the external partner profile. A CPA binding binds a CPA entry (CPA, service, and action) that you prefer to use. When a CPA entry is matched through the internal partner profile, outbound messages from the internal partner use the CPA, service, and action that are specified by the matched CPA entry. You must have the CPA entry in the ebXML settings of the associated B2B gateway before you add the entry to the CPA bindings list.", "profile-cpa-binding", "").String,
				NestedObject:        models.GetDmProfileCPABindingResourceSchema(),
				Optional:            true,
			},
			"ebms_cpa_id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the CPA ID to package in the outbound ebMS2 message header. This value is used when the B2B gateway is CPA-enforced.", "ebms-cpa-id", "").String,
				Optional:            true,
			},
			"ebms_service": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the service to package in the outbound ebMS2 message header. This value is used when when the B2B gateway is CPA-enforced.", "ebms-service", "").String,
				Optional:            true,
			},
			"ebms_action": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the action to package in the outbound ebMS2 message header. This value is used when the B2B gateway is CPA-enforced.", "ebms-action", "").String,
				Optional:            true,
			},
			"ebms_start_parameter": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate a start parameter for the ebMS2 message. The start parameter identifies the root part of the ebMS message. This setting is disabled by default. If enabled, the start parameter is generated with this value enclosed in angle brackets.", "ebms-start-parameter", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ebms_allow_duplicate_message": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify when to allow and reprocess duplicate ebMS2 inbound messages. The default behavior is never. This option does not apply to acknowledgments.", "ebms-allow-dup-msg", "").AddStringEnum("never", "always", "on-error").AddDefaultValue("never").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("never", "always", "on-error"),
				},
				Default: stringdefault.StaticString("never"),
			},
			"mdn_ssl_client_config_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("MDN TLS client type", "mdn-ssl-client-type", "").AddStringEnum("client").AddDefaultValue("client").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("client"),
				},
				Default: stringdefault.StaticString("client"),
			},
			"mdn_ssl_client": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("MDN TLS client profile", "mdn-ssl-client", "ssl_client_profile").String,
				Optional:            true,
			},
			"ebms_ack_ssl_client_config_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Acknowledgment/Error TLS client type", "ebmsack-ssl-client-type", "").AddStringEnum("client").AddDefaultValue("client").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("client"),
				},
				Default: stringdefault.StaticString("client"),
			},
			"ebms_ack_ssl_client": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Acknowledgment/Error TLS client profile", "ebmsack-ssl-client", "ssl_client_profile").String,
				Optional:            true,
			},
			"ebms3_outbound_sign": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to sign outbound messages. The default behavior is off. <ul><li>When enabled, signs outbound messages with the configured identification credentials and algorithm. If the configuration of a destination indicates to send messages unsigned, messages from this partner to that destination are not signed.</li><li>When disabled, does not sign outbound messages.</li></ul><p>This setting has no effect on outbound receipt signal messages. Regardless of this setting and whether an external partner requests a signed receipt signal, the outbound receipt signal is signed when this internal partner has a configured identification credentials.</p>", "ebms3-sign", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ebms3_outbound_sign_id_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the identification credentials to sign an outbound message or outbound receipt signal. The outbound receipt signal is signed if requested by a partner and this setting is configured.", "ebms3-sign-idcred", "crypto_ident_cred").AddRequiredWhen(models.B2BProfileEBMS3OutboundSignIdCredCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.B2BProfileEBMS3OutboundSignIdCredCondVal, validators.Evaluation{}, false),
				},
			},
			"ebms3_outbound_sign_digest_alg": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the algorithm to hash the payload during signing. The default value is sha1. See http://www.w3.org/TR/xmldsig-core/", "ebms3-sign-digest-alg", "").AddStringEnum("sha1", "sha256", "sha512", "ripemd160", "sha224", "sha384", "md5").AddDefaultValue("sha1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("sha1", "sha256", "sha512", "ripemd160", "sha224", "sha384", "md5"),
				},
				Default: stringdefault.StaticString("sha1"),
			},
			"ebms3_outbound_signature_alg": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the algorithm to sign outbound ebMS3 messages. The default value is dsa-sha1, which is recommended by the ebMS specification.", "ebms3-signature-alg", "").AddStringEnum("dsa-sha1", "rsa-sha1", "rsa-sha256", "rsa-sha384", "rsa-sha512", "rsa-ripemd160", "rsa-ripemd160-2010", "sha256-rsa-MGF1", "rsa-md5", "ecdsa-sha1", "ecdsa-sha224", "ecdsa-sha256", "ecdsa-sha384", "ecdsa-sha512").AddDefaultValue("rsa-sha1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("dsa-sha1", "rsa-sha1", "rsa-sha256", "rsa-sha384", "rsa-sha512", "rsa-ripemd160", "rsa-ripemd160-2010", "sha256-rsa-MGF1", "rsa-md5", "ecdsa-sha1", "ecdsa-sha224", "ecdsa-sha256", "ecdsa-sha384", "ecdsa-sha512"),
				},
				Default: stringdefault.StaticString("rsa-sha1"),
			},
			"ebms3_outbound_signature_c14n_alg": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the algorithm to canonicalize the SOAP Envelope XML and exclude comments before signing outbound ebMS3 messages. The default value is exc-c14n, which is recommended by the ebMS specification.", "ebms3-signature-c14n-alg", "").AddStringEnum("c14n", "exc-c14n", "c14n-comments", "exc-c14n-comments", "c14n11", "c14n11-comments").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("c14n", "exc-c14n", "c14n-comments", "exc-c14n-comments", "c14n11", "c14n11-comments"),
				},
			},
			"ebms3_inbound_verify_val_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Inbound signature validation credentials", "ebms3-verify-valcred", "crypto_val_cred").String,
				Optional:            true,
			},
			"ebms3_default_signer_cert": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the default validation certificate to verify the signature of an inbound ebMS3 message or a receipt. This certificate is used when the <tt>keyInfo</tt> element is missing or the signature method is unsupported.", "ebms3-default-signer-cert", "crypto_certificate").String,
				Optional:            true,
			},
			"ebms3_inbound_require_signed": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether inbound ebMS3 messages must be signed. The default behavior is off.", "ebms3-require-signed", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ebms3_inbound_require_encrypted": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether inbound ebMS3 messages must be encrypted. The default behavior is off.", "ebms3-require-encrypted", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ebms3_inbound_decrypt_id_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Inbound decryption identification credentials", "ebms3-decrypt-idcred", "crypto_ident_cred").AddRequiredWhen(models.B2BProfileEBMS3InboundDecryptIdCredCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.B2BProfileEBMS3InboundDecryptIdCredCondVal, validators.Evaluation{}, false),
				},
			},
			"ebms3_inbound_require_compressed": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether internal partners require compressed inbound ebMS3 messages. The default behavior is not to require compression.", "ebms3-require-compressed", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ebms3_receipt_ssl_client_config_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Receipt/Error TLS client type", "ebms3-receipt-ssl-client-type", "").AddStringEnum("client").AddDefaultValue("client").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("client"),
				},
				Default: stringdefault.StaticString("client"),
			},
			"ebms3_receipt_ssl_client": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Receipt/Error TLS client profile", "ebms3-receipt-ssl-client", "ssl_client_profile").AddRequiredWhen(models.B2BProfileEBMS3ReceiptSSLClientCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.B2BProfileEBMS3ReceiptSSLClientCondVal, validators.Evaluation{}, false),
				},
			},
			"ebms_notification": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable notification. When enabled, send notifications to the message producer or consumer when there are specific errors. The default behavior is off.", "ebms-notification", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ebms_notification_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Notification URL", "ebms-notification-url", "").AddRequiredWhen(models.B2BProfileEBMSNotificationURLCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.B2BProfileEBMSNotificationURLCondVal, validators.Evaluation{}, false),
				},
			},
			"ebms_notification_ssl_client_config_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Notification TLS client type", "ebms-notification-ssl-client-type", "").AddStringEnum("client").AddDefaultValue("client").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("client"),
				},
				Default: stringdefault.StaticString("client"),
			},
			"ebms_notification_ssl_client": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Notification TLS client profile", "ebms-notification-ssl-client", "ssl_client_profile").AddRequiredWhen(models.B2BProfileEBMSNotificationSSLClientCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.B2BProfileEBMSNotificationSSLClientCondVal, validators.Evaluation{}, false),
				},
			},
			"ebms3_allow_duplicate_message": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify when to allow and reprocess duplicate ebMS3 inbound messages. The default behavior is never. This option does not apply to the receipt signal messages.", "ebms3-allow-dup-msg", "").AddStringEnum("never", "always", "on-error").AddDefaultValue("never").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("never", "always", "on-error"),
				},
				Default: stringdefault.StaticString("never"),
			},
			"ebms3_duplicate_detection_notification": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable duplicate detection notification. When enabled, send notification to the message consumer when a duplicate message is detected. The default behavior is disabled.", "ebms3-duplicate-detection-notification", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ebms_message_properties": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the <tt>eb:Property</tt> elements to add to the <tt>eb:MessageProperties</tt> node. Define message properties to meet your business needs and the agreement between partners.", "ebms-messageproperties", "").String,
				NestedObject:        models.GetDmB2BMessagePropertiesResourceSchema(),
				Optional:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *B2BProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *B2BProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.B2BProfile
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `B2BProfile`)
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

func (r *B2BProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.B2BProfile
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
		data.FromBody(ctx, `B2BProfile`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `B2BProfile`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *B2BProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.B2BProfile
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `B2BProfile`))
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

func (r *B2BProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.B2BProfile
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

func (r *B2BProfileResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.B2BProfile

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
