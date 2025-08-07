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

type B2BCPASenderSettingList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &B2BCPASenderSettingDataSource{}
	_ datasource.DataSourceWithConfigure = &B2BCPASenderSettingDataSource{}
)

func NewB2BCPASenderSettingDataSource() datasource.DataSource {
	return &B2BCPASenderSettingDataSource{}
}

type B2BCPASenderSettingDataSource struct {
	client *client.DatapowerClient
}

func (d *B2BCPASenderSettingDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_b2bcpasendersetting"
}

func (d *B2BCPASenderSettingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "CPA sender setting",
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
						"enabled_doc_type": models.GetDmB2BEnabledDocTypeDataSourceSchema("Enabled document types", "enabled-doc-type", ""),
						"dest_endpoint_url": schema.StringAttribute{
							MarkdownDescription: "Destination URL",
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
							MarkdownDescription: "Connection timeout",
							Computed:            true,
						},
						"sync_reply_mode": schema.StringAttribute{
							MarkdownDescription: "Sync reply mode",
							Computed:            true,
						},
						"duplicate_elimination": schema.StringAttribute{
							MarkdownDescription: "Duplicate elimination",
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
							MarkdownDescription: "Retransmit attempts",
							Computed:            true,
						},
						"retry_interval": schema.Int64Attribute{
							MarkdownDescription: "Retransmit interval",
							Computed:            true,
						},
						"persist_duration": schema.Int64Attribute{
							MarkdownDescription: "Persistence duration",
							Computed:            true,
						},
						"include_time_to_live": schema.BoolAttribute{
							MarkdownDescription: "Include TimeToLive element",
							Computed:            true,
						},
						"encryption_required": schema.BoolAttribute{
							MarkdownDescription: "Require encryption",
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

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *B2BCPASenderSettingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data B2BCPASenderSettingList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.B2BCPASenderSetting{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.B2BCPASenderSetting{}
	if value := res.Get(`B2BCPASenderSetting`); value.Exists() {
		for _, v := range value.Array() {
			item := models.B2BCPASenderSetting{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
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
