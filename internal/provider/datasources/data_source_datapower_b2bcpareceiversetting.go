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

type B2BCPAReceiverSettingList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &B2BCPAReceiverSettingDataSource{}
	_ datasource.DataSourceWithConfigure = &B2BCPAReceiverSettingDataSource{}
)

func NewB2BCPAReceiverSettingDataSource() datasource.DataSource {
	return &B2BCPAReceiverSettingDataSource{}
}

type B2BCPAReceiverSettingDataSource struct {
	pData *tfutils.ProviderData
}

func (d *B2BCPAReceiverSettingDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_b2bcpareceiversetting"
}

func (d *B2BCPAReceiverSettingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "CPA receiver setting",
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
						"local_endpoint_uri": schema.StringAttribute{
							MarkdownDescription: "Local URI",
							Computed:            true,
						},
						"sync_reply_mode": schema.StringAttribute{
							MarkdownDescription: "Sync reply mode",
							Computed:            true,
						},
						"ack_requested": schema.StringAttribute{
							MarkdownDescription: "Expect acknowledgment requests",
							Computed:            true,
						},
						"ack_signature_requested": schema.StringAttribute{
							MarkdownDescription: "Expect signed acknowledgment requests",
							Computed:            true,
						},
						"allow_duplicate_message": schema.StringAttribute{
							MarkdownDescription: "Allow duplicate messages",
							Computed:            true,
						},
						"persist_duration": schema.Int64Attribute{
							MarkdownDescription: "Persistence duration",
							Computed:            true,
						},
						"encryption_required": schema.BoolAttribute{
							MarkdownDescription: "Require encryption",
							Computed:            true,
						},
						"decrypt_id_cred": schema.StringAttribute{
							MarkdownDescription: "Decryption identification credentials",
							Computed:            true,
						},
						"signature_required": schema.BoolAttribute{
							MarkdownDescription: "Require signature",
							Computed:            true,
						},
						"verify_val_cred": schema.StringAttribute{
							MarkdownDescription: "Signature validation credentials",
							Computed:            true,
						},
						"default_signer_cert": schema.StringAttribute{
							MarkdownDescription: "Default signature certificate",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *B2BCPAReceiverSettingDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *B2BCPAReceiverSettingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data B2BCPAReceiverSettingList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.B2BCPAReceiverSetting{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.B2BCPAReceiverSetting{}
	if value := res.Get(`B2BCPAReceiverSetting`); value.Exists() {
		for _, v := range value.Array() {
			item := models.B2BCPAReceiverSetting{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.B2BCPAReceiverSettingObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.B2BCPAReceiverSettingObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
