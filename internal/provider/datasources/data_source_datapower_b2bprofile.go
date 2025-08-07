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
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
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
	client *client.DatapowerClient
}

func (d *B2BProfileDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_b2bprofile"
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
							MarkdownDescription: "Profile type",
							Computed:            true,
						},
						"business_i_ds": schema.ListAttribute{
							MarkdownDescription: "Partner business IDs",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"business_i_ds_duns": schema.ListAttribute{
							MarkdownDescription: "Partner business IDs (DUNS)",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"business_i_ds_duns_plus4": schema.ListAttribute{
							MarkdownDescription: "Partner business IDs (DUNS+4)",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"custom_style_policy": schema.StringAttribute{
							MarkdownDescription: "Processing policy",
							Computed:            true,
						},
						"response_type": schema.StringAttribute{
							MarkdownDescription: "Response traffic type",
							Computed:            true,
						},
						"email_addresses": schema.ListAttribute{
							MarkdownDescription: "Partner email addresses",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"destinations": schema.ListNestedAttribute{
							MarkdownDescription: "Destinations",
							NestedObject:        models.DmB2BDestinationDataSourceSchema,
							Computed:            true,
						},
						"inbound_verify_val_cred": schema.StringAttribute{
							MarkdownDescription: "Inbound signature validation credentials",
							Computed:            true,
						},
						"inbound_require_signed": schema.BoolAttribute{
							MarkdownDescription: "Require signature",
							Computed:            true,
						},
						"inbound_require_encrypted": schema.BoolAttribute{
							MarkdownDescription: "Require encryption",
							Computed:            true,
						},
						"inbound_decrypt_id_cred": schema.StringAttribute{
							MarkdownDescription: "Inbound decryption identification credentials",
							Computed:            true,
						},
						"outbound_sign": schema.BoolAttribute{
							MarkdownDescription: "Sign outbound messages",
							Computed:            true,
						},
						"outbound_sign_id_cred": schema.StringAttribute{
							MarkdownDescription: "Signing identification credentials",
							Computed:            true,
						},
						"outbound_sign_digest_alg": schema.StringAttribute{
							MarkdownDescription: "Signing digest algorithm",
							Computed:            true,
						},
						"outbound_sign_mic_alg_version": schema.StringAttribute{
							MarkdownDescription: "Signing S/MIME version",
							Computed:            true,
						},
						"contacts": schema.ListNestedAttribute{
							MarkdownDescription: "Contacts",
							NestedObject:        models.DmB2BContactDataSourceSchema,
							Computed:            true,
						},
						"override_asid": schema.StringAttribute{
							MarkdownDescription: "Override AS identifier",
							Computed:            true,
						},
						"as_allow_duplicate_message": schema.StringAttribute{
							MarkdownDescription: "Allow duplicate AS inbound message",
							Computed:            true,
						},
						"preserve_filename": schema.BoolAttribute{
							MarkdownDescription: "Preserve file name",
							Computed:            true,
						},
						"ebms_role": schema.StringAttribute{
							MarkdownDescription: "Role",
							Computed:            true,
						},
						"ebms_persist_duration": schema.Int64Attribute{
							MarkdownDescription: "Persist duration",
							Computed:            true,
						},
						"ebms_ack_url": schema.StringAttribute{
							MarkdownDescription: "Acknowledgment URL",
							Computed:            true,
						},
						"ebms_error_url": schema.StringAttribute{
							MarkdownDescription: "Error URL",
							Computed:            true,
						},
						"ebms_inbound_send_receipt": schema.BoolAttribute{
							MarkdownDescription: "Reply with receipt",
							Computed:            true,
						},
						"ebms_inbound_send_signed_receipt": schema.BoolAttribute{
							MarkdownDescription: "Reply with signed receipt",
							Computed:            true,
						},
						"ebms_inbound_receipt_reply_pattern": schema.StringAttribute{
							MarkdownDescription: "Receipt reply pattern",
							Computed:            true,
						},
						"ebms_receipt_url": schema.StringAttribute{
							MarkdownDescription: "Receipt URL",
							Computed:            true,
						},
						"ebms_inbound_error_url": schema.StringAttribute{
							MarkdownDescription: "Error URL",
							Computed:            true,
						},
						"ebms_inbound_verify_val_cred": schema.StringAttribute{
							MarkdownDescription: "Inbound signature validation credentials",
							Computed:            true,
						},
						"ebms_default_signer_cert": schema.StringAttribute{
							MarkdownDescription: "Default inbound signature validation certificate",
							Computed:            true,
						},
						"ebms_inbound_require_signed": schema.BoolAttribute{
							MarkdownDescription: "Require signature",
							Computed:            true,
						},
						"ebms_inbound_require_encrypted": schema.BoolAttribute{
							MarkdownDescription: "Require encryption",
							Computed:            true,
						},
						"ebms_inbound_decrypt_id_cred": schema.StringAttribute{
							MarkdownDescription: "Inbound decryption identification credentials",
							Computed:            true,
						},
						"ebms_outbound_sign": schema.BoolAttribute{
							MarkdownDescription: "Sign outbound messages",
							Computed:            true,
						},
						"ebms_outbound_sign_id_cred": schema.StringAttribute{
							MarkdownDescription: "Signing identification credentials",
							Computed:            true,
						},
						"ebms_outbound_signature_alg": schema.StringAttribute{
							MarkdownDescription: "Signature algorithm",
							Computed:            true,
						},
						"ebms_outbound_signature_c14n_alg": schema.StringAttribute{
							MarkdownDescription: "Signature canonicalization method",
							Computed:            true,
						},
						"ebms_outbound_sign_digest_alg": schema.StringAttribute{
							MarkdownDescription: "Signing digest algorithm",
							Computed:            true,
						},
						"ebms_enable_cpa_binding": schema.BoolAttribute{
							MarkdownDescription: "Enable CPA bindings",
							Computed:            true,
						},
						"ebms_profile_cpa_bindings": schema.ListNestedAttribute{
							MarkdownDescription: "CPA bindings",
							NestedObject:        models.DmProfileCPABindingDataSourceSchema,
							Computed:            true,
						},
						"ebms_cpa_id": schema.StringAttribute{
							MarkdownDescription: "Default CPA ID",
							Computed:            true,
						},
						"ebms_service": schema.StringAttribute{
							MarkdownDescription: "Default service",
							Computed:            true,
						},
						"ebms_action": schema.StringAttribute{
							MarkdownDescription: "Default action",
							Computed:            true,
						},
						"ebms_start_parameter": schema.BoolAttribute{
							MarkdownDescription: "Generate start parameter",
							Computed:            true,
						},
						"ebms_allow_duplicate_message": schema.StringAttribute{
							MarkdownDescription: "Allow duplicate ebMS2 inbound message",
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
							MarkdownDescription: "Sign outbound messages",
							Computed:            true,
						},
						"ebms3_outbound_sign_id_cred": schema.StringAttribute{
							MarkdownDescription: "Signing identification credentials",
							Computed:            true,
						},
						"ebms3_outbound_sign_digest_alg": schema.StringAttribute{
							MarkdownDescription: "Signing digest algorithm",
							Computed:            true,
						},
						"ebms3_outbound_signature_alg": schema.StringAttribute{
							MarkdownDescription: "Signature algorithm",
							Computed:            true,
						},
						"ebms3_outbound_signature_c14n_alg": schema.StringAttribute{
							MarkdownDescription: "Signature canonicalization method",
							Computed:            true,
						},
						"ebms3_inbound_verify_val_cred": schema.StringAttribute{
							MarkdownDescription: "Inbound signature validation credentials",
							Computed:            true,
						},
						"ebms3_default_signer_cert": schema.StringAttribute{
							MarkdownDescription: "Default inbound signature validation certificate",
							Computed:            true,
						},
						"ebms3_inbound_require_signed": schema.BoolAttribute{
							MarkdownDescription: "Require signature",
							Computed:            true,
						},
						"ebms3_inbound_require_encrypted": schema.BoolAttribute{
							MarkdownDescription: "Require encryption",
							Computed:            true,
						},
						"ebms3_inbound_decrypt_id_cred": schema.StringAttribute{
							MarkdownDescription: "Inbound decryption identification credentials",
							Computed:            true,
						},
						"ebms3_inbound_require_compressed": schema.BoolAttribute{
							MarkdownDescription: "Require compression",
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
							MarkdownDescription: "Enable notification",
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
							MarkdownDescription: "Allow duplicate ebMS3 inbound message",
							Computed:            true,
						},
						"ebms3_duplicate_detection_notification": schema.BoolAttribute{
							MarkdownDescription: "Duplicate detection notification",
							Computed:            true,
						},
						"ebms_message_properties": schema.ListNestedAttribute{
							MarkdownDescription: "ebMS3 message properties",
							NestedObject:        models.DmB2BMessagePropertiesDataSourceSchema,
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

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *B2BProfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data B2BProfileList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.B2BProfile{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
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
