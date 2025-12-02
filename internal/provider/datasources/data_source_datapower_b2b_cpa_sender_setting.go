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
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type B2BCPASenderSettingList struct {
	ProviderTarget types.String `tfsdk:"provider_target"`
	Id             types.String `tfsdk:"id"`
	AppDomain      types.String `tfsdk:"app_domain"`
	Result         types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &B2BCPASenderSettingDataSource{}
	_ datasource.DataSourceWithConfigure = &B2BCPASenderSettingDataSource{}
)

func NewB2BCPASenderSettingDataSource() datasource.DataSource {
	return &B2BCPASenderSettingDataSource{}
}

type B2BCPASenderSettingDataSource struct {
	pData *tfutils.ProviderData
}

func (d *B2BCPASenderSettingDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_b2b_cpa_sender_setting"
}

func (d *B2BCPASenderSettingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "CPA sender setting",
		Attributes: map[string]schema.Attribute{
			"provider_target": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Target host to retrieve this data from. If not set, provider will use the top level settings.", "", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
			},
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
						"enabled_doc_type": models.GetDmB2BEnabledDocTypeDataSourceSchema("Enabled document types", "enabled-doc-type", ""),
						"dest_endpoint_url": schema.StringAttribute{
							MarkdownDescription: "Specify the destination URL for sending the message to the external party endpoint. For load distribution, use the name of the load-balancing group instead of the address-port pair in the URL.",
							Computed:            true,
						},
						"user_name": schema.StringAttribute{
							MarkdownDescription: "Username",
							Computed:            true,
						},
						"password_alias": schema.StringAttribute{
							MarkdownDescription: "Password alias",
							Computed:            true,
						},
						"connection_timeout": schema.Int64Attribute{
							MarkdownDescription: "Specify the duration in seconds to maintain an idle connection. Enter a value in the range 3 - 7200. The default value is 300.",
							Computed:            true,
						},
						"sync_reply_mode": schema.StringAttribute{
							MarkdownDescription: "Sync reply mode",
							Computed:            true,
						},
						"duplicate_elimination": schema.StringAttribute{
							MarkdownDescription: "For an outbound ebMS message, specify whether the internal sending party requests the external receiving party to check duplicate elimination. The request is made by presenting the <tt>DuplicateElimination</tt> element in the <tt>MessageHeader</tt> element in the ebMS SOAP header. <p>When imported from CPA, the <tt>duplicateElimination</tt> attribute on the internal party <tt>DeliveryChannel</tt> element in the <tt>MessagingCharacteristics</tt> element.</p>",
							Computed:            true,
						},
						"ack_requested": schema.StringAttribute{
							MarkdownDescription: "Request acknowledgment",
							Computed:            true,
						},
						"ack_signature_requested": schema.StringAttribute{
							MarkdownDescription: "Request signed acknowledgment",
							Computed:            true,
						},
						"retry": schema.BoolAttribute{
							MarkdownDescription: "Retransmit unacknowledged messages",
							Computed:            true,
						},
						"max_retries": schema.Int64Attribute{
							MarkdownDescription: "Specify the number of attempts to retransmit an unacknowledged message. Enter a value in the range 1 - 30. The default value is 3.",
							Computed:            true,
						},
						"retry_interval": schema.Int64Attribute{
							MarkdownDescription: "Specify the interval in seconds between retransmit attempts. Enter a value in the range 1 - 86400. The default value in 1800.",
							Computed:            true,
						},
						"persist_duration": schema.Int64Attribute{
							MarkdownDescription: "Specify the duration in seconds to retain messages in persistent storage. This value is used to compute the <tt>TimeToLive</tt> value. Until the value of the <tt>TimeToLive</tt> element elapses, the message cannot be archived.",
							Computed:            true,
						},
						"include_time_to_live": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to include the <tt>TimeToLive</tt> element in the outbound messages. This element indicates when the message expires. The receiving partner can accept the message only when it has not expired.",
							Computed:            true,
						},
						"encryption_required": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to encrypt outbound messages. Encryption does not apply to MSH level signals.",
							Computed:            true,
						},
						"encrypt_cert": schema.StringAttribute{
							MarkdownDescription: "Encryption certificate",
							Computed:            true,
						},
						"encrypt_algorithm": schema.StringAttribute{
							MarkdownDescription: "Encryption algorithm",
							Computed:            true,
						},
						"signature_required": schema.BoolAttribute{
							MarkdownDescription: "Require signature",
							Computed:            true,
						},
						"sign_id_cred": schema.StringAttribute{
							MarkdownDescription: "Signature identification credentials",
							Computed:            true,
						},
						"signature_algorithm": schema.StringAttribute{
							MarkdownDescription: "Signature algorithm",
							Computed:            true,
						},
						"sign_digest_algorithm": schema.StringAttribute{
							MarkdownDescription: "Signing digest algorithm",
							Computed:            true,
						},
						"signature_c14n_algorithm": schema.StringAttribute{
							MarkdownDescription: "Signature canonicalization method",
							Computed:            true,
						},
						"ssl_client_config_type": schema.StringAttribute{
							MarkdownDescription: "TLS client type",
							Computed:            true,
						},
						"ssl_client": schema.StringAttribute{
							MarkdownDescription: "TLS client profile",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
						"provider_target": schema.StringAttribute{
							MarkdownDescription: tfutils.NewAttributeDescription("Target host to retrieve this data from. If not set, provider will use the top level settings.", "", "").String,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *B2BCPASenderSettingDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *B2BCPASenderSettingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data B2BCPASenderSettingList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !d.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), d.pData.Client.GetTargetNames()))
		return
	}
	o := models.B2BCPASenderSetting{
		AppDomain: data.AppDomain,
	}

	path := o.GetPath()
	if !data.Id.IsNull() {
		path = path + "/" + data.Id.ValueString()
	}

	res, err := d.pData.Client.Get(path, data.ProviderTarget)
	resFound := true
	if err != nil {
		if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(d.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
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
	l := []models.B2BCPASenderSetting{}
	if resFound {
		if value := res.Get(`B2BCPASenderSetting`); value.Exists() {
			for _, v := range value.Array() {
				item := models.B2BCPASenderSetting{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.B2BCPASenderSettingObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.B2BCPASenderSettingObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
