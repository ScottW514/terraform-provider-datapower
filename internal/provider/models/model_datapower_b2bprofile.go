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

package models

import (
	"context"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type B2BProfile struct {
	Id                                  types.String `tfsdk:"id"`
	AppDomain                           types.String `tfsdk:"app_domain"`
	UserSummary                         types.String `tfsdk:"user_summary"`
	ProfileType                         types.String `tfsdk:"profile_type"`
	BusinessIDs                         types.List   `tfsdk:"business_i_ds"`
	BusinessIDsDuns                     types.List   `tfsdk:"business_i_ds_duns"`
	BusinessIDsDunsPlus4                types.List   `tfsdk:"business_i_ds_duns_plus4"`
	CustomStylePolicy                   types.String `tfsdk:"custom_style_policy"`
	ResponseType                        types.String `tfsdk:"response_type"`
	EmailAddresses                      types.List   `tfsdk:"email_addresses"`
	Destinations                        types.List   `tfsdk:"destinations"`
	InboundVerifyValCred                types.String `tfsdk:"inbound_verify_val_cred"`
	InboundRequireSigned                types.Bool   `tfsdk:"inbound_require_signed"`
	InboundRequireEncrypted             types.Bool   `tfsdk:"inbound_require_encrypted"`
	InboundDecryptIdCred                types.String `tfsdk:"inbound_decrypt_id_cred"`
	OutboundSign                        types.Bool   `tfsdk:"outbound_sign"`
	OutboundSignIdCred                  types.String `tfsdk:"outbound_sign_id_cred"`
	OutboundSignDigestAlg               types.String `tfsdk:"outbound_sign_digest_alg"`
	OutboundSignMicAlgVersion           types.String `tfsdk:"outbound_sign_mic_alg_version"`
	Contacts                            types.List   `tfsdk:"contacts"`
	OverrideAsid                        types.String `tfsdk:"override_asid"`
	AsAllowDuplicateMessage             types.String `tfsdk:"as_allow_duplicate_message"`
	PreserveFilename                    types.Bool   `tfsdk:"preserve_filename"`
	EbmsRole                            types.String `tfsdk:"ebms_role"`
	EbmsPersistDuration                 types.Int64  `tfsdk:"ebms_persist_duration"`
	EbmsAckUrl                          types.String `tfsdk:"ebms_ack_url"`
	EbmsErrorUrl                        types.String `tfsdk:"ebms_error_url"`
	EbmsInboundSendReceipt              types.Bool   `tfsdk:"ebms_inbound_send_receipt"`
	EbmsInboundSendSignedReceipt        types.Bool   `tfsdk:"ebms_inbound_send_signed_receipt"`
	EbmsInboundReceiptReplyPattern      types.String `tfsdk:"ebms_inbound_receipt_reply_pattern"`
	EbmsReceiptUrl                      types.String `tfsdk:"ebms_receipt_url"`
	EbmsInboundErrorUrl                 types.String `tfsdk:"ebms_inbound_error_url"`
	EbmsInboundVerifyValCred            types.String `tfsdk:"ebms_inbound_verify_val_cred"`
	EbmsDefaultSignerCert               types.String `tfsdk:"ebms_default_signer_cert"`
	EbmsInboundRequireSigned            types.Bool   `tfsdk:"ebms_inbound_require_signed"`
	EbmsInboundRequireEncrypted         types.Bool   `tfsdk:"ebms_inbound_require_encrypted"`
	EbmsInboundDecryptIdCred            types.String `tfsdk:"ebms_inbound_decrypt_id_cred"`
	EbmsOutboundSign                    types.Bool   `tfsdk:"ebms_outbound_sign"`
	EbmsOutboundSignIdCred              types.String `tfsdk:"ebms_outbound_sign_id_cred"`
	EbmsOutboundSignatureAlg            types.String `tfsdk:"ebms_outbound_signature_alg"`
	EbmsOutboundSignatureC14nAlg        types.String `tfsdk:"ebms_outbound_signature_c14n_alg"`
	EbmsOutboundSignDigestAlg           types.String `tfsdk:"ebms_outbound_sign_digest_alg"`
	EbmsEnableCpaBinding                types.Bool   `tfsdk:"ebms_enable_cpa_binding"`
	EbmsProfileCpaBindings              types.List   `tfsdk:"ebms_profile_cpa_bindings"`
	EbmsCpaId                           types.String `tfsdk:"ebms_cpa_id"`
	EbmsService                         types.String `tfsdk:"ebms_service"`
	EbmsAction                          types.String `tfsdk:"ebms_action"`
	EbmsStartParameter                  types.Bool   `tfsdk:"ebms_start_parameter"`
	EbmsAllowDuplicateMessage           types.String `tfsdk:"ebms_allow_duplicate_message"`
	MdnsslClientConfigType              types.String `tfsdk:"mdnssl_client_config_type"`
	MdnsslClient                        types.String `tfsdk:"mdnssl_client"`
	EbmsAckSslClientConfigType          types.String `tfsdk:"ebms_ack_ssl_client_config_type"`
	EbmsAckSslClient                    types.String `tfsdk:"ebms_ack_ssl_client"`
	Ebms3OutboundSign                   types.Bool   `tfsdk:"ebms3_outbound_sign"`
	Ebms3OutboundSignIdCred             types.String `tfsdk:"ebms3_outbound_sign_id_cred"`
	Ebms3OutboundSignDigestAlg          types.String `tfsdk:"ebms3_outbound_sign_digest_alg"`
	Ebms3OutboundSignatureAlg           types.String `tfsdk:"ebms3_outbound_signature_alg"`
	Ebms3OutboundSignatureC14nAlg       types.String `tfsdk:"ebms3_outbound_signature_c14n_alg"`
	Ebms3InboundVerifyValCred           types.String `tfsdk:"ebms3_inbound_verify_val_cred"`
	Ebms3DefaultSignerCert              types.String `tfsdk:"ebms3_default_signer_cert"`
	Ebms3InboundRequireSigned           types.Bool   `tfsdk:"ebms3_inbound_require_signed"`
	Ebms3InboundRequireEncrypted        types.Bool   `tfsdk:"ebms3_inbound_require_encrypted"`
	Ebms3InboundDecryptIdCred           types.String `tfsdk:"ebms3_inbound_decrypt_id_cred"`
	Ebms3InboundRequireCompressed       types.Bool   `tfsdk:"ebms3_inbound_require_compressed"`
	Ebms3ReceiptSslClientConfigType     types.String `tfsdk:"ebms3_receipt_ssl_client_config_type"`
	Ebms3ReceiptSslClient               types.String `tfsdk:"ebms3_receipt_ssl_client"`
	EbmsNotification                    types.Bool   `tfsdk:"ebms_notification"`
	EbmsNotificationUrl                 types.String `tfsdk:"ebms_notification_url"`
	EbmsNotificationSslClientConfigType types.String `tfsdk:"ebms_notification_ssl_client_config_type"`
	EbmsNotificationSslClient           types.String `tfsdk:"ebms_notification_ssl_client"`
	Ebms3AllowDuplicateMessage          types.String `tfsdk:"ebms3_allow_duplicate_message"`
	Ebms3DuplicateDetectionNotification types.Bool   `tfsdk:"ebms3_duplicate_detection_notification"`
	EbmsMessageProperties               types.List   `tfsdk:"ebms_message_properties"`
}

var B2BProfileObjectType = map[string]attr.Type{
	"id":                                       types.StringType,
	"app_domain":                               types.StringType,
	"user_summary":                             types.StringType,
	"profile_type":                             types.StringType,
	"business_i_ds":                            types.ListType{ElemType: types.StringType},
	"business_i_ds_duns":                       types.ListType{ElemType: types.StringType},
	"business_i_ds_duns_plus4":                 types.ListType{ElemType: types.StringType},
	"custom_style_policy":                      types.StringType,
	"response_type":                            types.StringType,
	"email_addresses":                          types.ListType{ElemType: types.StringType},
	"destinations":                             types.ListType{ElemType: types.ObjectType{AttrTypes: DmB2BDestinationObjectType}},
	"inbound_verify_val_cred":                  types.StringType,
	"inbound_require_signed":                   types.BoolType,
	"inbound_require_encrypted":                types.BoolType,
	"inbound_decrypt_id_cred":                  types.StringType,
	"outbound_sign":                            types.BoolType,
	"outbound_sign_id_cred":                    types.StringType,
	"outbound_sign_digest_alg":                 types.StringType,
	"outbound_sign_mic_alg_version":            types.StringType,
	"contacts":                                 types.ListType{ElemType: types.ObjectType{AttrTypes: DmB2BContactObjectType}},
	"override_asid":                            types.StringType,
	"as_allow_duplicate_message":               types.StringType,
	"preserve_filename":                        types.BoolType,
	"ebms_role":                                types.StringType,
	"ebms_persist_duration":                    types.Int64Type,
	"ebms_ack_url":                             types.StringType,
	"ebms_error_url":                           types.StringType,
	"ebms_inbound_send_receipt":                types.BoolType,
	"ebms_inbound_send_signed_receipt":         types.BoolType,
	"ebms_inbound_receipt_reply_pattern":       types.StringType,
	"ebms_receipt_url":                         types.StringType,
	"ebms_inbound_error_url":                   types.StringType,
	"ebms_inbound_verify_val_cred":             types.StringType,
	"ebms_default_signer_cert":                 types.StringType,
	"ebms_inbound_require_signed":              types.BoolType,
	"ebms_inbound_require_encrypted":           types.BoolType,
	"ebms_inbound_decrypt_id_cred":             types.StringType,
	"ebms_outbound_sign":                       types.BoolType,
	"ebms_outbound_sign_id_cred":               types.StringType,
	"ebms_outbound_signature_alg":              types.StringType,
	"ebms_outbound_signature_c14n_alg":         types.StringType,
	"ebms_outbound_sign_digest_alg":            types.StringType,
	"ebms_enable_cpa_binding":                  types.BoolType,
	"ebms_profile_cpa_bindings":                types.ListType{ElemType: types.ObjectType{AttrTypes: DmProfileCPABindingObjectType}},
	"ebms_cpa_id":                              types.StringType,
	"ebms_service":                             types.StringType,
	"ebms_action":                              types.StringType,
	"ebms_start_parameter":                     types.BoolType,
	"ebms_allow_duplicate_message":             types.StringType,
	"mdnssl_client_config_type":                types.StringType,
	"mdnssl_client":                            types.StringType,
	"ebms_ack_ssl_client_config_type":          types.StringType,
	"ebms_ack_ssl_client":                      types.StringType,
	"ebms3_outbound_sign":                      types.BoolType,
	"ebms3_outbound_sign_id_cred":              types.StringType,
	"ebms3_outbound_sign_digest_alg":           types.StringType,
	"ebms3_outbound_signature_alg":             types.StringType,
	"ebms3_outbound_signature_c14n_alg":        types.StringType,
	"ebms3_inbound_verify_val_cred":            types.StringType,
	"ebms3_default_signer_cert":                types.StringType,
	"ebms3_inbound_require_signed":             types.BoolType,
	"ebms3_inbound_require_encrypted":          types.BoolType,
	"ebms3_inbound_decrypt_id_cred":            types.StringType,
	"ebms3_inbound_require_compressed":         types.BoolType,
	"ebms3_receipt_ssl_client_config_type":     types.StringType,
	"ebms3_receipt_ssl_client":                 types.StringType,
	"ebms_notification":                        types.BoolType,
	"ebms_notification_url":                    types.StringType,
	"ebms_notification_ssl_client_config_type": types.StringType,
	"ebms_notification_ssl_client":             types.StringType,
	"ebms3_allow_duplicate_message":            types.StringType,
	"ebms3_duplicate_detection_notification":   types.BoolType,
	"ebms_message_properties":                  types.ListType{ElemType: types.ObjectType{AttrTypes: DmB2BMessagePropertiesObjectType}},
}

func (data B2BProfile) GetPath() string {
	rest_path := "/mgmt/config/{domain}/B2BProfile"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data B2BProfile) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.ProfileType.IsNull() {
		return false
	}
	if !data.BusinessIDs.IsNull() {
		return false
	}
	if !data.BusinessIDsDuns.IsNull() {
		return false
	}
	if !data.BusinessIDsDunsPlus4.IsNull() {
		return false
	}
	if !data.CustomStylePolicy.IsNull() {
		return false
	}
	if !data.ResponseType.IsNull() {
		return false
	}
	if !data.EmailAddresses.IsNull() {
		return false
	}
	if !data.Destinations.IsNull() {
		return false
	}
	if !data.InboundVerifyValCred.IsNull() {
		return false
	}
	if !data.InboundRequireSigned.IsNull() {
		return false
	}
	if !data.InboundRequireEncrypted.IsNull() {
		return false
	}
	if !data.InboundDecryptIdCred.IsNull() {
		return false
	}
	if !data.OutboundSign.IsNull() {
		return false
	}
	if !data.OutboundSignIdCred.IsNull() {
		return false
	}
	if !data.OutboundSignDigestAlg.IsNull() {
		return false
	}
	if !data.OutboundSignMicAlgVersion.IsNull() {
		return false
	}
	if !data.Contacts.IsNull() {
		return false
	}
	if !data.OverrideAsid.IsNull() {
		return false
	}
	if !data.AsAllowDuplicateMessage.IsNull() {
		return false
	}
	if !data.PreserveFilename.IsNull() {
		return false
	}
	if !data.EbmsRole.IsNull() {
		return false
	}
	if !data.EbmsPersistDuration.IsNull() {
		return false
	}
	if !data.EbmsAckUrl.IsNull() {
		return false
	}
	if !data.EbmsErrorUrl.IsNull() {
		return false
	}
	if !data.EbmsInboundSendReceipt.IsNull() {
		return false
	}
	if !data.EbmsInboundSendSignedReceipt.IsNull() {
		return false
	}
	if !data.EbmsInboundReceiptReplyPattern.IsNull() {
		return false
	}
	if !data.EbmsReceiptUrl.IsNull() {
		return false
	}
	if !data.EbmsInboundErrorUrl.IsNull() {
		return false
	}
	if !data.EbmsInboundVerifyValCred.IsNull() {
		return false
	}
	if !data.EbmsDefaultSignerCert.IsNull() {
		return false
	}
	if !data.EbmsInboundRequireSigned.IsNull() {
		return false
	}
	if !data.EbmsInboundRequireEncrypted.IsNull() {
		return false
	}
	if !data.EbmsInboundDecryptIdCred.IsNull() {
		return false
	}
	if !data.EbmsOutboundSign.IsNull() {
		return false
	}
	if !data.EbmsOutboundSignIdCred.IsNull() {
		return false
	}
	if !data.EbmsOutboundSignatureAlg.IsNull() {
		return false
	}
	if !data.EbmsOutboundSignatureC14nAlg.IsNull() {
		return false
	}
	if !data.EbmsOutboundSignDigestAlg.IsNull() {
		return false
	}
	if !data.EbmsEnableCpaBinding.IsNull() {
		return false
	}
	if !data.EbmsProfileCpaBindings.IsNull() {
		return false
	}
	if !data.EbmsCpaId.IsNull() {
		return false
	}
	if !data.EbmsService.IsNull() {
		return false
	}
	if !data.EbmsAction.IsNull() {
		return false
	}
	if !data.EbmsStartParameter.IsNull() {
		return false
	}
	if !data.EbmsAllowDuplicateMessage.IsNull() {
		return false
	}
	if !data.MdnsslClientConfigType.IsNull() {
		return false
	}
	if !data.MdnsslClient.IsNull() {
		return false
	}
	if !data.EbmsAckSslClientConfigType.IsNull() {
		return false
	}
	if !data.EbmsAckSslClient.IsNull() {
		return false
	}
	if !data.Ebms3OutboundSign.IsNull() {
		return false
	}
	if !data.Ebms3OutboundSignIdCred.IsNull() {
		return false
	}
	if !data.Ebms3OutboundSignDigestAlg.IsNull() {
		return false
	}
	if !data.Ebms3OutboundSignatureAlg.IsNull() {
		return false
	}
	if !data.Ebms3OutboundSignatureC14nAlg.IsNull() {
		return false
	}
	if !data.Ebms3InboundVerifyValCred.IsNull() {
		return false
	}
	if !data.Ebms3DefaultSignerCert.IsNull() {
		return false
	}
	if !data.Ebms3InboundRequireSigned.IsNull() {
		return false
	}
	if !data.Ebms3InboundRequireEncrypted.IsNull() {
		return false
	}
	if !data.Ebms3InboundDecryptIdCred.IsNull() {
		return false
	}
	if !data.Ebms3InboundRequireCompressed.IsNull() {
		return false
	}
	if !data.Ebms3ReceiptSslClientConfigType.IsNull() {
		return false
	}
	if !data.Ebms3ReceiptSslClient.IsNull() {
		return false
	}
	if !data.EbmsNotification.IsNull() {
		return false
	}
	if !data.EbmsNotificationUrl.IsNull() {
		return false
	}
	if !data.EbmsNotificationSslClientConfigType.IsNull() {
		return false
	}
	if !data.EbmsNotificationSslClient.IsNull() {
		return false
	}
	if !data.Ebms3AllowDuplicateMessage.IsNull() {
		return false
	}
	if !data.Ebms3DuplicateDetectionNotification.IsNull() {
		return false
	}
	if !data.EbmsMessageProperties.IsNull() {
		return false
	}
	return true
}

func (data B2BProfile) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.ProfileType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ProfileType`, data.ProfileType.ValueString())
	}
	if !data.BusinessIDs.IsNull() {
		var values []string
		data.BusinessIDs.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`BusinessIDs`+".-1", map[string]string{"value": val})
		}
	}
	if !data.BusinessIDsDuns.IsNull() {
		var values []string
		data.BusinessIDsDuns.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`BusinessIDsDUNS`+".-1", map[string]string{"value": val})
		}
	}
	if !data.BusinessIDsDunsPlus4.IsNull() {
		var values []string
		data.BusinessIDsDunsPlus4.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`BusinessIDsDUNSPlus4`+".-1", map[string]string{"value": val})
		}
	}
	if !data.CustomStylePolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CustomStylePolicy`, data.CustomStylePolicy.ValueString())
	}
	if !data.ResponseType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ResponseType`, data.ResponseType.ValueString())
	}
	if !data.EmailAddresses.IsNull() {
		var values []string
		data.EmailAddresses.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`EmailAddresses`+".-1", map[string]string{"value": val})
		}
	}
	if !data.Destinations.IsNull() {
		var values []DmB2BDestination
		data.Destinations.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`Destinations`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.InboundVerifyValCred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`InboundVerifyValCred`, data.InboundVerifyValCred.ValueString())
	}
	if !data.InboundRequireSigned.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`InboundRequireSigned`, tfutils.StringFromBool(data.InboundRequireSigned, ""))
	}
	if !data.InboundRequireEncrypted.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`InboundRequireEncrypted`, tfutils.StringFromBool(data.InboundRequireEncrypted, ""))
	}
	if !data.InboundDecryptIdCred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`InboundDecryptIdCred`, data.InboundDecryptIdCred.ValueString())
	}
	if !data.OutboundSign.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OutboundSign`, tfutils.StringFromBool(data.OutboundSign, ""))
	}
	if !data.OutboundSignIdCred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OutboundSignIdCred`, data.OutboundSignIdCred.ValueString())
	}
	if !data.OutboundSignDigestAlg.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OutboundSignDigestAlg`, data.OutboundSignDigestAlg.ValueString())
	}
	if !data.OutboundSignMicAlgVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OutboundSignMICAlgVersion`, data.OutboundSignMicAlgVersion.ValueString())
	}
	if !data.Contacts.IsNull() {
		var values []DmB2BContact
		data.Contacts.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`Contacts`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.OverrideAsid.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OverrideASID`, data.OverrideAsid.ValueString())
	}
	if !data.AsAllowDuplicateMessage.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ASAllowDuplicateMessage`, data.AsAllowDuplicateMessage.ValueString())
	}
	if !data.PreserveFilename.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PreserveFilename`, tfutils.StringFromBool(data.PreserveFilename, ""))
	}
	if !data.EbmsRole.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSRole`, data.EbmsRole.ValueString())
	}
	if !data.EbmsPersistDuration.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSPersistDuration`, data.EbmsPersistDuration.ValueInt64())
	}
	if !data.EbmsAckUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSAckURL`, data.EbmsAckUrl.ValueString())
	}
	if !data.EbmsErrorUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSErrorURL`, data.EbmsErrorUrl.ValueString())
	}
	if !data.EbmsInboundSendReceipt.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSInboundSendReceipt`, tfutils.StringFromBool(data.EbmsInboundSendReceipt, ""))
	}
	if !data.EbmsInboundSendSignedReceipt.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSInboundSendSignedReceipt`, tfutils.StringFromBool(data.EbmsInboundSendSignedReceipt, ""))
	}
	if !data.EbmsInboundReceiptReplyPattern.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSInboundReceiptReplyPattern`, data.EbmsInboundReceiptReplyPattern.ValueString())
	}
	if !data.EbmsReceiptUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSReceiptURL`, data.EbmsReceiptUrl.ValueString())
	}
	if !data.EbmsInboundErrorUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSInboundErrorURL`, data.EbmsInboundErrorUrl.ValueString())
	}
	if !data.EbmsInboundVerifyValCred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSInboundVerifyValCred`, data.EbmsInboundVerifyValCred.ValueString())
	}
	if !data.EbmsDefaultSignerCert.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSDefaultSignerCert`, data.EbmsDefaultSignerCert.ValueString())
	}
	if !data.EbmsInboundRequireSigned.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSInboundRequireSigned`, tfutils.StringFromBool(data.EbmsInboundRequireSigned, ""))
	}
	if !data.EbmsInboundRequireEncrypted.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSInboundRequireEncrypted`, tfutils.StringFromBool(data.EbmsInboundRequireEncrypted, ""))
	}
	if !data.EbmsInboundDecryptIdCred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSInboundDecryptIdCred`, data.EbmsInboundDecryptIdCred.ValueString())
	}
	if !data.EbmsOutboundSign.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSOutboundSign`, tfutils.StringFromBool(data.EbmsOutboundSign, ""))
	}
	if !data.EbmsOutboundSignIdCred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSOutboundSignIdCred`, data.EbmsOutboundSignIdCred.ValueString())
	}
	if !data.EbmsOutboundSignatureAlg.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSOutboundSignatureAlg`, data.EbmsOutboundSignatureAlg.ValueString())
	}
	if !data.EbmsOutboundSignatureC14nAlg.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSOutboundSignatureC14NAlg`, data.EbmsOutboundSignatureC14nAlg.ValueString())
	}
	if !data.EbmsOutboundSignDigestAlg.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSOutboundSignDigestAlg`, data.EbmsOutboundSignDigestAlg.ValueString())
	}
	if !data.EbmsEnableCpaBinding.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSEnableCPABinding`, tfutils.StringFromBool(data.EbmsEnableCpaBinding, ""))
	}
	if !data.EbmsProfileCpaBindings.IsNull() {
		var values []DmProfileCPABinding
		data.EbmsProfileCpaBindings.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`EBMSProfileCPABindings`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.EbmsCpaId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSCpaId`, data.EbmsCpaId.ValueString())
	}
	if !data.EbmsService.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSService`, data.EbmsService.ValueString())
	}
	if !data.EbmsAction.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSAction`, data.EbmsAction.ValueString())
	}
	if !data.EbmsStartParameter.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSStartParameter`, tfutils.StringFromBool(data.EbmsStartParameter, ""))
	}
	if !data.EbmsAllowDuplicateMessage.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSAllowDuplicateMessage`, data.EbmsAllowDuplicateMessage.ValueString())
	}
	if !data.MdnsslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MDNSSLClientConfigType`, data.MdnsslClientConfigType.ValueString())
	}
	if !data.MdnsslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MDNSSLClient`, data.MdnsslClient.ValueString())
	}
	if !data.EbmsAckSslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSAckSSLClientConfigType`, data.EbmsAckSslClientConfigType.ValueString())
	}
	if !data.EbmsAckSslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSAckSSLClient`, data.EbmsAckSslClient.ValueString())
	}
	if !data.Ebms3OutboundSign.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMS3OutboundSign`, tfutils.StringFromBool(data.Ebms3OutboundSign, ""))
	}
	if !data.Ebms3OutboundSignIdCred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMS3OutboundSignIdCred`, data.Ebms3OutboundSignIdCred.ValueString())
	}
	if !data.Ebms3OutboundSignDigestAlg.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMS3OutboundSignDigestAlg`, data.Ebms3OutboundSignDigestAlg.ValueString())
	}
	if !data.Ebms3OutboundSignatureAlg.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMS3OutboundSignatureAlg`, data.Ebms3OutboundSignatureAlg.ValueString())
	}
	if !data.Ebms3OutboundSignatureC14nAlg.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMS3OutboundSignatureC14NAlg`, data.Ebms3OutboundSignatureC14nAlg.ValueString())
	}
	if !data.Ebms3InboundVerifyValCred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMS3InboundVerifyValCred`, data.Ebms3InboundVerifyValCred.ValueString())
	}
	if !data.Ebms3DefaultSignerCert.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMS3DefaultSignerCert`, data.Ebms3DefaultSignerCert.ValueString())
	}
	if !data.Ebms3InboundRequireSigned.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMS3InboundRequireSigned`, tfutils.StringFromBool(data.Ebms3InboundRequireSigned, ""))
	}
	if !data.Ebms3InboundRequireEncrypted.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMS3InboundRequireEncrypted`, tfutils.StringFromBool(data.Ebms3InboundRequireEncrypted, ""))
	}
	if !data.Ebms3InboundDecryptIdCred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMS3InboundDecryptIdCred`, data.Ebms3InboundDecryptIdCred.ValueString())
	}
	if !data.Ebms3InboundRequireCompressed.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMS3InboundRequireCompressed`, tfutils.StringFromBool(data.Ebms3InboundRequireCompressed, ""))
	}
	if !data.Ebms3ReceiptSslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMS3ReceiptSSLClientConfigType`, data.Ebms3ReceiptSslClientConfigType.ValueString())
	}
	if !data.Ebms3ReceiptSslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMS3ReceiptSSLClient`, data.Ebms3ReceiptSslClient.ValueString())
	}
	if !data.EbmsNotification.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSNotification`, tfutils.StringFromBool(data.EbmsNotification, ""))
	}
	if !data.EbmsNotificationUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSNotificationURL`, data.EbmsNotificationUrl.ValueString())
	}
	if !data.EbmsNotificationSslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSNotificationSSLClientConfigType`, data.EbmsNotificationSslClientConfigType.ValueString())
	}
	if !data.EbmsNotificationSslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSNotificationSSLClient`, data.EbmsNotificationSslClient.ValueString())
	}
	if !data.Ebms3AllowDuplicateMessage.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMS3AllowDuplicateMessage`, data.Ebms3AllowDuplicateMessage.ValueString())
	}
	if !data.Ebms3DuplicateDetectionNotification.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMS3DuplicateDetectionNotification`, tfutils.StringFromBool(data.Ebms3DuplicateDetectionNotification, ""))
	}
	if !data.EbmsMessageProperties.IsNull() {
		var values []DmB2BMessageProperties
		data.EbmsMessageProperties.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`EBMSMessageProperties`+".-1", val.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *B2BProfile) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProfileType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ProfileType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProfileType = types.StringValue("internal")
	}
	if value := res.Get(pathRoot + `BusinessIDs`); value.Exists() {
		data.BusinessIDs = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.BusinessIDs = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `BusinessIDsDUNS`); value.Exists() {
		data.BusinessIDsDuns = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.BusinessIDsDuns = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `BusinessIDsDUNSPlus4`); value.Exists() {
		data.BusinessIDsDunsPlus4 = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.BusinessIDsDunsPlus4 = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `CustomStylePolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CustomStylePolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CustomStylePolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `ResponseType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ResponseType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ResponseType = types.StringValue("preprocessed")
	}
	if value := res.Get(pathRoot + `EmailAddresses`); value.Exists() {
		data.EmailAddresses = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.EmailAddresses = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `Destinations`); value.Exists() {
		l := []DmB2BDestination{}
		if value := res.Get(`Destinations`); value.Exists() {
			for _, v := range value.Array() {
				item := DmB2BDestination{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Destinations, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmB2BDestinationObjectType}, l)
		} else {
			data.Destinations = types.ListNull(types.ObjectType{AttrTypes: DmB2BDestinationObjectType})
		}
	} else {
		data.Destinations = types.ListNull(types.ObjectType{AttrTypes: DmB2BDestinationObjectType})
	}
	if value := res.Get(pathRoot + `InboundVerifyValCred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.InboundVerifyValCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.InboundVerifyValCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `InboundRequireSigned`); value.Exists() {
		data.InboundRequireSigned = tfutils.BoolFromString(value.String())
	} else {
		data.InboundRequireSigned = types.BoolNull()
	}
	if value := res.Get(pathRoot + `InboundRequireEncrypted`); value.Exists() {
		data.InboundRequireEncrypted = tfutils.BoolFromString(value.String())
	} else {
		data.InboundRequireEncrypted = types.BoolNull()
	}
	if value := res.Get(pathRoot + `InboundDecryptIdCred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.InboundDecryptIdCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.InboundDecryptIdCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `OutboundSign`); value.Exists() {
		data.OutboundSign = tfutils.BoolFromString(value.String())
	} else {
		data.OutboundSign = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OutboundSignIdCred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OutboundSignIdCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OutboundSignIdCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `OutboundSignDigestAlg`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OutboundSignDigestAlg = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OutboundSignDigestAlg = types.StringValue("sha1")
	}
	if value := res.Get(pathRoot + `OutboundSignMICAlgVersion`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OutboundSignMicAlgVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OutboundSignMicAlgVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `Contacts`); value.Exists() {
		l := []DmB2BContact{}
		if value := res.Get(`Contacts`); value.Exists() {
			for _, v := range value.Array() {
				item := DmB2BContact{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Contacts, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmB2BContactObjectType}, l)
		} else {
			data.Contacts = types.ListNull(types.ObjectType{AttrTypes: DmB2BContactObjectType})
		}
	} else {
		data.Contacts = types.ListNull(types.ObjectType{AttrTypes: DmB2BContactObjectType})
	}
	if value := res.Get(pathRoot + `OverrideASID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OverrideAsid = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OverrideAsid = types.StringNull()
	}
	if value := res.Get(pathRoot + `ASAllowDuplicateMessage`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AsAllowDuplicateMessage = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AsAllowDuplicateMessage = types.StringValue("never")
	}
	if value := res.Get(pathRoot + `PreserveFilename`); value.Exists() {
		data.PreserveFilename = tfutils.BoolFromString(value.String())
	} else {
		data.PreserveFilename = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSRole`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsRole = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsRole = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSPersistDuration`); value.Exists() {
		data.EbmsPersistDuration = types.Int64Value(value.Int())
	} else {
		data.EbmsPersistDuration = types.Int64Null()
	}
	if value := res.Get(pathRoot + `EBMSAckURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsAckUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsAckUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSErrorURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsErrorUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsErrorUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSInboundSendReceipt`); value.Exists() {
		data.EbmsInboundSendReceipt = tfutils.BoolFromString(value.String())
	} else {
		data.EbmsInboundSendReceipt = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSInboundSendSignedReceipt`); value.Exists() {
		data.EbmsInboundSendSignedReceipt = tfutils.BoolFromString(value.String())
	} else {
		data.EbmsInboundSendSignedReceipt = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSInboundReceiptReplyPattern`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsInboundReceiptReplyPattern = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsInboundReceiptReplyPattern = types.StringValue("Response")
	}
	if value := res.Get(pathRoot + `EBMSReceiptURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsReceiptUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsReceiptUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSInboundErrorURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsInboundErrorUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsInboundErrorUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSInboundVerifyValCred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsInboundVerifyValCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsInboundVerifyValCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSDefaultSignerCert`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsDefaultSignerCert = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsDefaultSignerCert = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSInboundRequireSigned`); value.Exists() {
		data.EbmsInboundRequireSigned = tfutils.BoolFromString(value.String())
	} else {
		data.EbmsInboundRequireSigned = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSInboundRequireEncrypted`); value.Exists() {
		data.EbmsInboundRequireEncrypted = tfutils.BoolFromString(value.String())
	} else {
		data.EbmsInboundRequireEncrypted = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSInboundDecryptIdCred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsInboundDecryptIdCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsInboundDecryptIdCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSOutboundSign`); value.Exists() {
		data.EbmsOutboundSign = tfutils.BoolFromString(value.String())
	} else {
		data.EbmsOutboundSign = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSOutboundSignIdCred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsOutboundSignIdCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsOutboundSignIdCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSOutboundSignatureAlg`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsOutboundSignatureAlg = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsOutboundSignatureAlg = types.StringValue("dsa-sha1")
	}
	if value := res.Get(pathRoot + `EBMSOutboundSignatureC14NAlg`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsOutboundSignatureC14nAlg = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsOutboundSignatureC14nAlg = types.StringValue("c14n")
	}
	if value := res.Get(pathRoot + `EBMSOutboundSignDigestAlg`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsOutboundSignDigestAlg = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsOutboundSignDigestAlg = types.StringValue("sha1")
	}
	if value := res.Get(pathRoot + `EBMSEnableCPABinding`); value.Exists() {
		data.EbmsEnableCpaBinding = tfutils.BoolFromString(value.String())
	} else {
		data.EbmsEnableCpaBinding = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSProfileCPABindings`); value.Exists() {
		l := []DmProfileCPABinding{}
		if value := res.Get(`EBMSProfileCPABindings`); value.Exists() {
			for _, v := range value.Array() {
				item := DmProfileCPABinding{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.EbmsProfileCpaBindings, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmProfileCPABindingObjectType}, l)
		} else {
			data.EbmsProfileCpaBindings = types.ListNull(types.ObjectType{AttrTypes: DmProfileCPABindingObjectType})
		}
	} else {
		data.EbmsProfileCpaBindings = types.ListNull(types.ObjectType{AttrTypes: DmProfileCPABindingObjectType})
	}
	if value := res.Get(pathRoot + `EBMSCpaId`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsCpaId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsCpaId = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSService`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsService = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsService = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSAction`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsAction = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsAction = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSStartParameter`); value.Exists() {
		data.EbmsStartParameter = tfutils.BoolFromString(value.String())
	} else {
		data.EbmsStartParameter = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSAllowDuplicateMessage`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsAllowDuplicateMessage = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsAllowDuplicateMessage = types.StringValue("never")
	}
	if value := res.Get(pathRoot + `MDNSSLClientConfigType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MdnsslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MdnsslClientConfigType = types.StringValue("client")
	}
	if value := res.Get(pathRoot + `MDNSSLClient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MdnsslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MdnsslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSAckSSLClientConfigType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsAckSslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsAckSslClientConfigType = types.StringValue("client")
	}
	if value := res.Get(pathRoot + `EBMSAckSSLClient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsAckSslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsAckSslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMS3OutboundSign`); value.Exists() {
		data.Ebms3OutboundSign = tfutils.BoolFromString(value.String())
	} else {
		data.Ebms3OutboundSign = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMS3OutboundSignIdCred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Ebms3OutboundSignIdCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Ebms3OutboundSignIdCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMS3OutboundSignDigestAlg`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Ebms3OutboundSignDigestAlg = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Ebms3OutboundSignDigestAlg = types.StringValue("sha1")
	}
	if value := res.Get(pathRoot + `EBMS3OutboundSignatureAlg`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Ebms3OutboundSignatureAlg = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Ebms3OutboundSignatureAlg = types.StringValue("rsa-sha1")
	}
	if value := res.Get(pathRoot + `EBMS3OutboundSignatureC14NAlg`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Ebms3OutboundSignatureC14nAlg = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Ebms3OutboundSignatureC14nAlg = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMS3InboundVerifyValCred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Ebms3InboundVerifyValCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Ebms3InboundVerifyValCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMS3DefaultSignerCert`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Ebms3DefaultSignerCert = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Ebms3DefaultSignerCert = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMS3InboundRequireSigned`); value.Exists() {
		data.Ebms3InboundRequireSigned = tfutils.BoolFromString(value.String())
	} else {
		data.Ebms3InboundRequireSigned = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMS3InboundRequireEncrypted`); value.Exists() {
		data.Ebms3InboundRequireEncrypted = tfutils.BoolFromString(value.String())
	} else {
		data.Ebms3InboundRequireEncrypted = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMS3InboundDecryptIdCred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Ebms3InboundDecryptIdCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Ebms3InboundDecryptIdCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMS3InboundRequireCompressed`); value.Exists() {
		data.Ebms3InboundRequireCompressed = tfutils.BoolFromString(value.String())
	} else {
		data.Ebms3InboundRequireCompressed = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMS3ReceiptSSLClientConfigType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Ebms3ReceiptSslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Ebms3ReceiptSslClientConfigType = types.StringValue("client")
	}
	if value := res.Get(pathRoot + `EBMS3ReceiptSSLClient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Ebms3ReceiptSslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Ebms3ReceiptSslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSNotification`); value.Exists() {
		data.EbmsNotification = tfutils.BoolFromString(value.String())
	} else {
		data.EbmsNotification = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSNotificationURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsNotificationUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsNotificationUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSNotificationSSLClientConfigType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsNotificationSslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsNotificationSslClientConfigType = types.StringValue("client")
	}
	if value := res.Get(pathRoot + `EBMSNotificationSSLClient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsNotificationSslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsNotificationSslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMS3AllowDuplicateMessage`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Ebms3AllowDuplicateMessage = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Ebms3AllowDuplicateMessage = types.StringValue("never")
	}
	if value := res.Get(pathRoot + `EBMS3DuplicateDetectionNotification`); value.Exists() {
		data.Ebms3DuplicateDetectionNotification = tfutils.BoolFromString(value.String())
	} else {
		data.Ebms3DuplicateDetectionNotification = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSMessageProperties`); value.Exists() {
		l := []DmB2BMessageProperties{}
		if value := res.Get(`EBMSMessageProperties`); value.Exists() {
			for _, v := range value.Array() {
				item := DmB2BMessageProperties{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.EbmsMessageProperties, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmB2BMessagePropertiesObjectType}, l)
		} else {
			data.EbmsMessageProperties = types.ListNull(types.ObjectType{AttrTypes: DmB2BMessagePropertiesObjectType})
		}
	} else {
		data.EbmsMessageProperties = types.ListNull(types.ObjectType{AttrTypes: DmB2BMessagePropertiesObjectType})
	}
}

func (data *B2BProfile) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProfileType`); value.Exists() && !data.ProfileType.IsNull() {
		data.ProfileType = tfutils.ParseStringFromGJSON(value)
	} else if data.ProfileType.ValueString() != "internal" {
		data.ProfileType = types.StringNull()
	}
	if value := res.Get(pathRoot + `BusinessIDs`); value.Exists() && !data.BusinessIDs.IsNull() {
		data.BusinessIDs = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.BusinessIDs = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `BusinessIDsDUNS`); value.Exists() && !data.BusinessIDsDuns.IsNull() {
		data.BusinessIDsDuns = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.BusinessIDsDuns = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `BusinessIDsDUNSPlus4`); value.Exists() && !data.BusinessIDsDunsPlus4.IsNull() {
		data.BusinessIDsDunsPlus4 = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.BusinessIDsDunsPlus4 = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `CustomStylePolicy`); value.Exists() && !data.CustomStylePolicy.IsNull() {
		data.CustomStylePolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CustomStylePolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `ResponseType`); value.Exists() && !data.ResponseType.IsNull() {
		data.ResponseType = tfutils.ParseStringFromGJSON(value)
	} else if data.ResponseType.ValueString() != "preprocessed" {
		data.ResponseType = types.StringNull()
	}
	if value := res.Get(pathRoot + `EmailAddresses`); value.Exists() && !data.EmailAddresses.IsNull() {
		data.EmailAddresses = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.EmailAddresses = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `Destinations`); value.Exists() && !data.Destinations.IsNull() {
		l := []DmB2BDestination{}
		for _, v := range value.Array() {
			item := DmB2BDestination{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.Destinations, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmB2BDestinationObjectType}, l)
		} else {
			data.Destinations = types.ListNull(types.ObjectType{AttrTypes: DmB2BDestinationObjectType})
		}
	} else {
		data.Destinations = types.ListNull(types.ObjectType{AttrTypes: DmB2BDestinationObjectType})
	}
	if value := res.Get(pathRoot + `InboundVerifyValCred`); value.Exists() && !data.InboundVerifyValCred.IsNull() {
		data.InboundVerifyValCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.InboundVerifyValCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `InboundRequireSigned`); value.Exists() && !data.InboundRequireSigned.IsNull() {
		data.InboundRequireSigned = tfutils.BoolFromString(value.String())
	} else if data.InboundRequireSigned.ValueBool() {
		data.InboundRequireSigned = types.BoolNull()
	}
	if value := res.Get(pathRoot + `InboundRequireEncrypted`); value.Exists() && !data.InboundRequireEncrypted.IsNull() {
		data.InboundRequireEncrypted = tfutils.BoolFromString(value.String())
	} else if data.InboundRequireEncrypted.ValueBool() {
		data.InboundRequireEncrypted = types.BoolNull()
	}
	if value := res.Get(pathRoot + `InboundDecryptIdCred`); value.Exists() && !data.InboundDecryptIdCred.IsNull() {
		data.InboundDecryptIdCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.InboundDecryptIdCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `OutboundSign`); value.Exists() && !data.OutboundSign.IsNull() {
		data.OutboundSign = tfutils.BoolFromString(value.String())
	} else if data.OutboundSign.ValueBool() {
		data.OutboundSign = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OutboundSignIdCred`); value.Exists() && !data.OutboundSignIdCred.IsNull() {
		data.OutboundSignIdCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OutboundSignIdCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `OutboundSignDigestAlg`); value.Exists() && !data.OutboundSignDigestAlg.IsNull() {
		data.OutboundSignDigestAlg = tfutils.ParseStringFromGJSON(value)
	} else if data.OutboundSignDigestAlg.ValueString() != "sha1" {
		data.OutboundSignDigestAlg = types.StringNull()
	}
	if value := res.Get(pathRoot + `OutboundSignMICAlgVersion`); value.Exists() && !data.OutboundSignMicAlgVersion.IsNull() {
		data.OutboundSignMicAlgVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OutboundSignMicAlgVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `Contacts`); value.Exists() && !data.Contacts.IsNull() {
		l := []DmB2BContact{}
		for _, v := range value.Array() {
			item := DmB2BContact{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.Contacts, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmB2BContactObjectType}, l)
		} else {
			data.Contacts = types.ListNull(types.ObjectType{AttrTypes: DmB2BContactObjectType})
		}
	} else {
		data.Contacts = types.ListNull(types.ObjectType{AttrTypes: DmB2BContactObjectType})
	}
	if value := res.Get(pathRoot + `OverrideASID`); value.Exists() && !data.OverrideAsid.IsNull() {
		data.OverrideAsid = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OverrideAsid = types.StringNull()
	}
	if value := res.Get(pathRoot + `ASAllowDuplicateMessage`); value.Exists() && !data.AsAllowDuplicateMessage.IsNull() {
		data.AsAllowDuplicateMessage = tfutils.ParseStringFromGJSON(value)
	} else if data.AsAllowDuplicateMessage.ValueString() != "never" {
		data.AsAllowDuplicateMessage = types.StringNull()
	}
	if value := res.Get(pathRoot + `PreserveFilename`); value.Exists() && !data.PreserveFilename.IsNull() {
		data.PreserveFilename = tfutils.BoolFromString(value.String())
	} else if data.PreserveFilename.ValueBool() {
		data.PreserveFilename = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSRole`); value.Exists() && !data.EbmsRole.IsNull() {
		data.EbmsRole = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsRole = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSPersistDuration`); value.Exists() && !data.EbmsPersistDuration.IsNull() {
		data.EbmsPersistDuration = types.Int64Value(value.Int())
	} else {
		data.EbmsPersistDuration = types.Int64Null()
	}
	if value := res.Get(pathRoot + `EBMSAckURL`); value.Exists() && !data.EbmsAckUrl.IsNull() {
		data.EbmsAckUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsAckUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSErrorURL`); value.Exists() && !data.EbmsErrorUrl.IsNull() {
		data.EbmsErrorUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsErrorUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSInboundSendReceipt`); value.Exists() && !data.EbmsInboundSendReceipt.IsNull() {
		data.EbmsInboundSendReceipt = tfutils.BoolFromString(value.String())
	} else if data.EbmsInboundSendReceipt.ValueBool() {
		data.EbmsInboundSendReceipt = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSInboundSendSignedReceipt`); value.Exists() && !data.EbmsInboundSendSignedReceipt.IsNull() {
		data.EbmsInboundSendSignedReceipt = tfutils.BoolFromString(value.String())
	} else if data.EbmsInboundSendSignedReceipt.ValueBool() {
		data.EbmsInboundSendSignedReceipt = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSInboundReceiptReplyPattern`); value.Exists() && !data.EbmsInboundReceiptReplyPattern.IsNull() {
		data.EbmsInboundReceiptReplyPattern = tfutils.ParseStringFromGJSON(value)
	} else if data.EbmsInboundReceiptReplyPattern.ValueString() != "Response" {
		data.EbmsInboundReceiptReplyPattern = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSReceiptURL`); value.Exists() && !data.EbmsReceiptUrl.IsNull() {
		data.EbmsReceiptUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsReceiptUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSInboundErrorURL`); value.Exists() && !data.EbmsInboundErrorUrl.IsNull() {
		data.EbmsInboundErrorUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsInboundErrorUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSInboundVerifyValCred`); value.Exists() && !data.EbmsInboundVerifyValCred.IsNull() {
		data.EbmsInboundVerifyValCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsInboundVerifyValCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSDefaultSignerCert`); value.Exists() && !data.EbmsDefaultSignerCert.IsNull() {
		data.EbmsDefaultSignerCert = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsDefaultSignerCert = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSInboundRequireSigned`); value.Exists() && !data.EbmsInboundRequireSigned.IsNull() {
		data.EbmsInboundRequireSigned = tfutils.BoolFromString(value.String())
	} else if data.EbmsInboundRequireSigned.ValueBool() {
		data.EbmsInboundRequireSigned = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSInboundRequireEncrypted`); value.Exists() && !data.EbmsInboundRequireEncrypted.IsNull() {
		data.EbmsInboundRequireEncrypted = tfutils.BoolFromString(value.String())
	} else if data.EbmsInboundRequireEncrypted.ValueBool() {
		data.EbmsInboundRequireEncrypted = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSInboundDecryptIdCred`); value.Exists() && !data.EbmsInboundDecryptIdCred.IsNull() {
		data.EbmsInboundDecryptIdCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsInboundDecryptIdCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSOutboundSign`); value.Exists() && !data.EbmsOutboundSign.IsNull() {
		data.EbmsOutboundSign = tfutils.BoolFromString(value.String())
	} else if data.EbmsOutboundSign.ValueBool() {
		data.EbmsOutboundSign = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSOutboundSignIdCred`); value.Exists() && !data.EbmsOutboundSignIdCred.IsNull() {
		data.EbmsOutboundSignIdCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsOutboundSignIdCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSOutboundSignatureAlg`); value.Exists() && !data.EbmsOutboundSignatureAlg.IsNull() {
		data.EbmsOutboundSignatureAlg = tfutils.ParseStringFromGJSON(value)
	} else if data.EbmsOutboundSignatureAlg.ValueString() != "dsa-sha1" {
		data.EbmsOutboundSignatureAlg = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSOutboundSignatureC14NAlg`); value.Exists() && !data.EbmsOutboundSignatureC14nAlg.IsNull() {
		data.EbmsOutboundSignatureC14nAlg = tfutils.ParseStringFromGJSON(value)
	} else if data.EbmsOutboundSignatureC14nAlg.ValueString() != "c14n" {
		data.EbmsOutboundSignatureC14nAlg = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSOutboundSignDigestAlg`); value.Exists() && !data.EbmsOutboundSignDigestAlg.IsNull() {
		data.EbmsOutboundSignDigestAlg = tfutils.ParseStringFromGJSON(value)
	} else if data.EbmsOutboundSignDigestAlg.ValueString() != "sha1" {
		data.EbmsOutboundSignDigestAlg = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSEnableCPABinding`); value.Exists() && !data.EbmsEnableCpaBinding.IsNull() {
		data.EbmsEnableCpaBinding = tfutils.BoolFromString(value.String())
	} else if data.EbmsEnableCpaBinding.ValueBool() {
		data.EbmsEnableCpaBinding = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSProfileCPABindings`); value.Exists() && !data.EbmsProfileCpaBindings.IsNull() {
		l := []DmProfileCPABinding{}
		for _, v := range value.Array() {
			item := DmProfileCPABinding{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.EbmsProfileCpaBindings, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmProfileCPABindingObjectType}, l)
		} else {
			data.EbmsProfileCpaBindings = types.ListNull(types.ObjectType{AttrTypes: DmProfileCPABindingObjectType})
		}
	} else {
		data.EbmsProfileCpaBindings = types.ListNull(types.ObjectType{AttrTypes: DmProfileCPABindingObjectType})
	}
	if value := res.Get(pathRoot + `EBMSCpaId`); value.Exists() && !data.EbmsCpaId.IsNull() {
		data.EbmsCpaId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsCpaId = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSService`); value.Exists() && !data.EbmsService.IsNull() {
		data.EbmsService = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsService = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSAction`); value.Exists() && !data.EbmsAction.IsNull() {
		data.EbmsAction = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsAction = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSStartParameter`); value.Exists() && !data.EbmsStartParameter.IsNull() {
		data.EbmsStartParameter = tfutils.BoolFromString(value.String())
	} else if data.EbmsStartParameter.ValueBool() {
		data.EbmsStartParameter = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSAllowDuplicateMessage`); value.Exists() && !data.EbmsAllowDuplicateMessage.IsNull() {
		data.EbmsAllowDuplicateMessage = tfutils.ParseStringFromGJSON(value)
	} else if data.EbmsAllowDuplicateMessage.ValueString() != "never" {
		data.EbmsAllowDuplicateMessage = types.StringNull()
	}
	if value := res.Get(pathRoot + `MDNSSLClientConfigType`); value.Exists() && !data.MdnsslClientConfigType.IsNull() {
		data.MdnsslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else if data.MdnsslClientConfigType.ValueString() != "client" {
		data.MdnsslClientConfigType = types.StringNull()
	}
	if value := res.Get(pathRoot + `MDNSSLClient`); value.Exists() && !data.MdnsslClient.IsNull() {
		data.MdnsslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MdnsslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSAckSSLClientConfigType`); value.Exists() && !data.EbmsAckSslClientConfigType.IsNull() {
		data.EbmsAckSslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else if data.EbmsAckSslClientConfigType.ValueString() != "client" {
		data.EbmsAckSslClientConfigType = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSAckSSLClient`); value.Exists() && !data.EbmsAckSslClient.IsNull() {
		data.EbmsAckSslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsAckSslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMS3OutboundSign`); value.Exists() && !data.Ebms3OutboundSign.IsNull() {
		data.Ebms3OutboundSign = tfutils.BoolFromString(value.String())
	} else if data.Ebms3OutboundSign.ValueBool() {
		data.Ebms3OutboundSign = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMS3OutboundSignIdCred`); value.Exists() && !data.Ebms3OutboundSignIdCred.IsNull() {
		data.Ebms3OutboundSignIdCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Ebms3OutboundSignIdCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMS3OutboundSignDigestAlg`); value.Exists() && !data.Ebms3OutboundSignDigestAlg.IsNull() {
		data.Ebms3OutboundSignDigestAlg = tfutils.ParseStringFromGJSON(value)
	} else if data.Ebms3OutboundSignDigestAlg.ValueString() != "sha1" {
		data.Ebms3OutboundSignDigestAlg = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMS3OutboundSignatureAlg`); value.Exists() && !data.Ebms3OutboundSignatureAlg.IsNull() {
		data.Ebms3OutboundSignatureAlg = tfutils.ParseStringFromGJSON(value)
	} else if data.Ebms3OutboundSignatureAlg.ValueString() != "rsa-sha1" {
		data.Ebms3OutboundSignatureAlg = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMS3OutboundSignatureC14NAlg`); value.Exists() && !data.Ebms3OutboundSignatureC14nAlg.IsNull() {
		data.Ebms3OutboundSignatureC14nAlg = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Ebms3OutboundSignatureC14nAlg = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMS3InboundVerifyValCred`); value.Exists() && !data.Ebms3InboundVerifyValCred.IsNull() {
		data.Ebms3InboundVerifyValCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Ebms3InboundVerifyValCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMS3DefaultSignerCert`); value.Exists() && !data.Ebms3DefaultSignerCert.IsNull() {
		data.Ebms3DefaultSignerCert = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Ebms3DefaultSignerCert = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMS3InboundRequireSigned`); value.Exists() && !data.Ebms3InboundRequireSigned.IsNull() {
		data.Ebms3InboundRequireSigned = tfutils.BoolFromString(value.String())
	} else if data.Ebms3InboundRequireSigned.ValueBool() {
		data.Ebms3InboundRequireSigned = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMS3InboundRequireEncrypted`); value.Exists() && !data.Ebms3InboundRequireEncrypted.IsNull() {
		data.Ebms3InboundRequireEncrypted = tfutils.BoolFromString(value.String())
	} else if data.Ebms3InboundRequireEncrypted.ValueBool() {
		data.Ebms3InboundRequireEncrypted = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMS3InboundDecryptIdCred`); value.Exists() && !data.Ebms3InboundDecryptIdCred.IsNull() {
		data.Ebms3InboundDecryptIdCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Ebms3InboundDecryptIdCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMS3InboundRequireCompressed`); value.Exists() && !data.Ebms3InboundRequireCompressed.IsNull() {
		data.Ebms3InboundRequireCompressed = tfutils.BoolFromString(value.String())
	} else if data.Ebms3InboundRequireCompressed.ValueBool() {
		data.Ebms3InboundRequireCompressed = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMS3ReceiptSSLClientConfigType`); value.Exists() && !data.Ebms3ReceiptSslClientConfigType.IsNull() {
		data.Ebms3ReceiptSslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else if data.Ebms3ReceiptSslClientConfigType.ValueString() != "client" {
		data.Ebms3ReceiptSslClientConfigType = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMS3ReceiptSSLClient`); value.Exists() && !data.Ebms3ReceiptSslClient.IsNull() {
		data.Ebms3ReceiptSslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Ebms3ReceiptSslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSNotification`); value.Exists() && !data.EbmsNotification.IsNull() {
		data.EbmsNotification = tfutils.BoolFromString(value.String())
	} else if data.EbmsNotification.ValueBool() {
		data.EbmsNotification = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSNotificationURL`); value.Exists() && !data.EbmsNotificationUrl.IsNull() {
		data.EbmsNotificationUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsNotificationUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSNotificationSSLClientConfigType`); value.Exists() && !data.EbmsNotificationSslClientConfigType.IsNull() {
		data.EbmsNotificationSslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else if data.EbmsNotificationSslClientConfigType.ValueString() != "client" {
		data.EbmsNotificationSslClientConfigType = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSNotificationSSLClient`); value.Exists() && !data.EbmsNotificationSslClient.IsNull() {
		data.EbmsNotificationSslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsNotificationSslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMS3AllowDuplicateMessage`); value.Exists() && !data.Ebms3AllowDuplicateMessage.IsNull() {
		data.Ebms3AllowDuplicateMessage = tfutils.ParseStringFromGJSON(value)
	} else if data.Ebms3AllowDuplicateMessage.ValueString() != "never" {
		data.Ebms3AllowDuplicateMessage = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMS3DuplicateDetectionNotification`); value.Exists() && !data.Ebms3DuplicateDetectionNotification.IsNull() {
		data.Ebms3DuplicateDetectionNotification = tfutils.BoolFromString(value.String())
	} else if data.Ebms3DuplicateDetectionNotification.ValueBool() {
		data.Ebms3DuplicateDetectionNotification = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSMessageProperties`); value.Exists() && !data.EbmsMessageProperties.IsNull() {
		l := []DmB2BMessageProperties{}
		for _, v := range value.Array() {
			item := DmB2BMessageProperties{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.EbmsMessageProperties, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmB2BMessagePropertiesObjectType}, l)
		} else {
			data.EbmsMessageProperties = types.ListNull(types.ObjectType{AttrTypes: DmB2BMessagePropertiesObjectType})
		}
	} else {
		data.EbmsMessageProperties = types.ListNull(types.ObjectType{AttrTypes: DmB2BMessagePropertiesObjectType})
	}
}
