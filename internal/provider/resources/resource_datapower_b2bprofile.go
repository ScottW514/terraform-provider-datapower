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

var _ resource.Resource = &B2BProfileResource{}

func NewB2BProfileResource() resource.Resource {
	return &B2BProfileResource{}
}

type B2BProfileResource struct {
	client *client.DatapowerClient
}

func (r *B2BProfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_b2bprofile"
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
			"profile_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Profile type", "profile-type", "").AddStringEnum("internal", "external").AddDefaultValue("internal").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("internal", "external"),
				},
				Default: stringdefault.StaticString("internal"),
			},
			"business_i_ds": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Partner business IDs", "business-id", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"business_i_ds_duns": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Partner business IDs (DUNS)", "business-id-duns", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"business_i_ds_duns_plus4": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Partner business IDs (DUNS+4)", "business-id-duns4", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"custom_style_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Processing policy", "stylepolicy", "stylepolicy").String,
				Optional:            true,
			},
			"response_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Response traffic type", "response-type", "").AddStringEnum("soap", "xml", "unprocessed", "preprocessed").AddDefaultValue("preprocessed").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("soap", "xml", "unprocessed", "preprocessed"),
				},
				Default: stringdefault.StaticString("preprocessed"),
			},
			"email_addresses": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Partner email addresses", "email-address", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"destinations": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Destinations", "destination", "").String,
				NestedObject:        models.DmB2BDestinationResourceSchema,
				Required:            true,
			},
			"inbound_verify_val_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Inbound signature validation credentials", "verify-valcred", "cryptovalcred").String,
				Optional:            true,
			},
			"inbound_require_signed": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Require signature", "require-signed", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"inbound_require_encrypted": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Require encryption", "require-encrypted", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"inbound_decrypt_id_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Inbound decryption identification credentials", "decrypt-idcred", "cryptoidentcred").String,
				Optional:            true,
			},
			"outbound_sign": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Sign outbound messages", "sign", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"outbound_sign_id_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Signing identification credentials", "sign-idcred", "cryptoidentcred").String,
				Optional:            true,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Signing S/MIME version", "sign-micalg-version", "").AddStringEnum("SMIME3.1", "SMIME3.2").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("SMIME3.1", "SMIME3.2"),
				},
			},
			"contacts": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Contacts", "contact", "").String,
				NestedObject:        models.DmB2BContactResourceSchema,
				Optional:            true,
			},
			"override_asid": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Override AS identifier", "override-as-identifier", "").String,
				Optional:            true,
			},
			"as_allow_duplicate_message": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow duplicate AS inbound message", "as-allow-dup-msg", "").AddStringEnum("never", "always", "on-error").AddDefaultValue("never").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("never", "always", "on-error"),
				},
				Default: stringdefault.StaticString("never"),
			},
			"preserve_filename": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Preserve file name", "preserve-filename", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ebms_role": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Role", "ebms-role", "").String,
				Optional:            true,
			},
			"ebms_persist_duration": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Persist duration", "ebms-persist-duration", "").AddIntegerRange(0, 6000000).String,
				Optional:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 6000000),
				},
			},
			"ebms_ack_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Acknowledgment URL", "ebms-ack-url", "").String,
				Optional:            true,
			},
			"ebms_error_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Error URL", "ebms-error-url", "").String,
				Optional:            true,
			},
			"ebms_inbound_send_receipt": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Reply with receipt", "ebms-inbound-send-receipt", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ebms_inbound_send_signed_receipt": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Reply with signed receipt", "ebms-inbound-send-signed-receipt", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ebms_inbound_receipt_reply_pattern": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Receipt reply pattern", "ebms-inbound-receipt-reply-pattern", "").AddStringEnum("Response", "Callback").AddDefaultValue("Response").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Response", "Callback"),
				},
				Default: stringdefault.StaticString("Response"),
			},
			"ebms_receipt_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Receipt URL", "ebms-receipt-url", "").String,
				Optional:            true,
			},
			"ebms_inbound_error_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Error URL", "ebms-inbound-error-url", "").String,
				Optional:            true,
			},
			"ebms_inbound_verify_val_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Inbound signature validation credentials", "ebms-verify-valcred", "cryptovalcred").String,
				Optional:            true,
			},
			"ebms_default_signer_cert": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Default inbound signature validation certificate", "ebms-default-signer-cert", "cryptocertificate").String,
				Optional:            true,
			},
			"ebms_inbound_require_signed": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Require signature", "ebms-require-signed", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ebms_inbound_require_encrypted": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Require encryption", "ebms-require-encrypted", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ebms_inbound_decrypt_id_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Inbound decryption identification credentials", "ebms-decrypt-idcred", "cryptoidentcred").String,
				Optional:            true,
			},
			"ebms_outbound_sign": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Sign outbound messages", "ebms-sign", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ebms_outbound_sign_id_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Signing identification credentials", "ebms-sign-idcred", "cryptoidentcred").String,
				Optional:            true,
			},
			"ebms_outbound_signature_alg": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Signature algorithm", "ebms-signature-alg", "").AddStringEnum("dsa-sha1", "rsa-sha1", "rsa-sha256", "rsa-sha384", "rsa-sha512", "rsa-ripemd160", "rsa-ripemd160-2010", "sha256-rsa-MGF1", "rsa-md5", "ecdsa-sha1", "ecdsa-sha224", "ecdsa-sha256", "ecdsa-sha384", "ecdsa-sha512").AddDefaultValue("dsa-sha1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("dsa-sha1", "rsa-sha1", "rsa-sha256", "rsa-sha384", "rsa-sha512", "rsa-ripemd160", "rsa-ripemd160-2010", "sha256-rsa-MGF1", "rsa-md5", "ecdsa-sha1", "ecdsa-sha224", "ecdsa-sha256", "ecdsa-sha384", "ecdsa-sha512"),
				},
				Default: stringdefault.StaticString("dsa-sha1"),
			},
			"ebms_outbound_signature_c14n_alg": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Signature canonicalization method", "ebms-signature-c14n-alg", "").AddStringEnum("c14n", "exc-c14n", "c14n-comments", "exc-c14n-comments", "c14n11", "c14n11-comments").AddDefaultValue("c14n").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("c14n", "exc-c14n", "c14n-comments", "exc-c14n-comments", "c14n11", "c14n11-comments"),
				},
				Default: stringdefault.StaticString("c14n"),
			},
			"ebms_outbound_sign_digest_alg": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Signing digest algorithm", "ebms-sign-digest-alg", "").AddStringEnum("sha1", "sha256", "sha512", "ripemd160", "sha224", "sha384", "md5").AddDefaultValue("sha1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("sha1", "sha256", "sha512", "ripemd160", "sha224", "sha384", "md5"),
				},
				Default: stringdefault.StaticString("sha1"),
			},
			"ebms_enable_cpa_binding": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable CPA bindings", "ebms-enable-cpa-bindings", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ebms_profile_cpa_bindings": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("CPA bindings", "profile-cpa-binding", "").String,
				NestedObject:        models.DmProfileCPABindingResourceSchema,
				Optional:            true,
			},
			"ebms_cpa_id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Default CPA ID", "ebms-cpa-id", "").String,
				Optional:            true,
			},
			"ebms_service": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Default service", "ebms-service", "").String,
				Optional:            true,
			},
			"ebms_action": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Default action", "ebms-action", "").String,
				Optional:            true,
			},
			"ebms_start_parameter": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Generate start parameter", "ebms-start-parameter", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ebms_allow_duplicate_message": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow duplicate ebMS2 inbound message", "ebms-allow-dup-msg", "").AddStringEnum("never", "always", "on-error").AddDefaultValue("never").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("never", "always", "on-error"),
				},
				Default: stringdefault.StaticString("never"),
			},
			"mdnssl_client_config_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("MDN TLS client type", "mdn-ssl-client-type", "").AddStringEnum("client").AddDefaultValue("client").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("client"),
				},
				Default: stringdefault.StaticString("client"),
			},
			"mdnssl_client": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("MDN TLS client profile", "mdn-ssl-client", "sslclientprofile").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Acknowledgment/Error TLS client profile", "ebmsack-ssl-client", "sslclientprofile").String,
				Optional:            true,
			},
			"ebms3_outbound_sign": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Sign outbound messages", "ebms3-sign", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ebms3_outbound_sign_id_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Signing identification credentials", "ebms3-sign-idcred", "cryptoidentcred").String,
				Optional:            true,
			},
			"ebms3_outbound_sign_digest_alg": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Signing digest algorithm", "ebms3-sign-digest-alg", "").AddStringEnum("sha1", "sha256", "sha512", "ripemd160", "sha224", "sha384", "md5").AddDefaultValue("sha1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("sha1", "sha256", "sha512", "ripemd160", "sha224", "sha384", "md5"),
				},
				Default: stringdefault.StaticString("sha1"),
			},
			"ebms3_outbound_signature_alg": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Signature algorithm", "ebms3-signature-alg", "").AddStringEnum("dsa-sha1", "rsa-sha1", "rsa-sha256", "rsa-sha384", "rsa-sha512", "rsa-ripemd160", "rsa-ripemd160-2010", "sha256-rsa-MGF1", "rsa-md5", "ecdsa-sha1", "ecdsa-sha224", "ecdsa-sha256", "ecdsa-sha384", "ecdsa-sha512").AddDefaultValue("rsa-sha1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("dsa-sha1", "rsa-sha1", "rsa-sha256", "rsa-sha384", "rsa-sha512", "rsa-ripemd160", "rsa-ripemd160-2010", "sha256-rsa-MGF1", "rsa-md5", "ecdsa-sha1", "ecdsa-sha224", "ecdsa-sha256", "ecdsa-sha384", "ecdsa-sha512"),
				},
				Default: stringdefault.StaticString("rsa-sha1"),
			},
			"ebms3_outbound_signature_c14n_alg": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Signature canonicalization method", "ebms3-signature-c14n-alg", "").AddStringEnum("c14n", "exc-c14n", "c14n-comments", "exc-c14n-comments", "c14n11", "c14n11-comments").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("c14n", "exc-c14n", "c14n-comments", "exc-c14n-comments", "c14n11", "c14n11-comments"),
				},
			},
			"ebms3_inbound_verify_val_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Inbound signature validation credentials", "ebms3-verify-valcred", "cryptovalcred").String,
				Optional:            true,
			},
			"ebms3_default_signer_cert": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Default inbound signature validation certificate", "ebms3-default-signer-cert", "cryptocertificate").String,
				Optional:            true,
			},
			"ebms3_inbound_require_signed": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Require signature", "ebms3-require-signed", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ebms3_inbound_require_encrypted": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Require encryption", "ebms3-require-encrypted", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ebms3_inbound_decrypt_id_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Inbound decryption identification credentials", "ebms3-decrypt-idcred", "cryptoidentcred").String,
				Optional:            true,
			},
			"ebms3_inbound_require_compressed": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Require compression", "ebms3-require-compressed", "").AddDefaultValue("false").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Receipt/Error TLS client profile", "ebms3-receipt-ssl-client", "sslclientprofile").String,
				Optional:            true,
			},
			"ebms_notification": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable notification", "ebms-notification", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ebms_notification_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Notification URL", "ebms-notification-url", "").String,
				Optional:            true,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Notification TLS client profile", "ebms-notification-ssl-client", "sslclientprofile").String,
				Optional:            true,
			},
			"ebms3_allow_duplicate_message": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow duplicate ebMS3 inbound message", "ebms3-allow-dup-msg", "").AddStringEnum("never", "always", "on-error").AddDefaultValue("never").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("never", "always", "on-error"),
				},
				Default: stringdefault.StaticString("never"),
			},
			"ebms3_duplicate_detection_notification": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Duplicate detection notification", "ebms3-duplicate-detection-notification", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ebms_message_properties": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("ebMS3 message properties", "ebms-messageproperties", "").String,
				NestedObject:        models.DmB2BMessagePropertiesResourceSchema,
				Optional:            true,
			},
		},
	}
}

func (r *B2BProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *B2BProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.B2BProfile

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `B2BProfile`)
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

func (r *B2BProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.B2BProfile

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
		data.FromBody(ctx, `B2BProfile`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `B2BProfile`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *B2BProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.B2BProfile

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `B2BProfile`))
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

func (r *B2BProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.B2BProfile

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
