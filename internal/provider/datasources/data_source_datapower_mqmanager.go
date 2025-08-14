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

type MQManagerList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &MQManagerDataSource{}
	_ datasource.DataSourceWithConfigure = &MQManagerDataSource{}
)

func NewMQManagerDataSource() datasource.DataSource {
	return &MQManagerDataSource{}
}

type MQManagerDataSource struct {
	pData *tfutils.ProviderData
}

func (d *MQManagerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mqmanager"
}

func (d *MQManagerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "IBM MQ v9+ queue manager",
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
						"host_name": schema.StringAttribute{
							MarkdownDescription: "Host",
							Computed:            true,
						},
						"q_mname": schema.StringAttribute{
							MarkdownDescription: "Queue manager name",
							Computed:            true,
						},
						"ccsid": schema.Int64Attribute{
							MarkdownDescription: "Coded character set ID",
							Computed:            true,
						},
						"channel_name": schema.StringAttribute{
							MarkdownDescription: "Channel name",
							Computed:            true,
						},
						"csp_user_id": schema.StringAttribute{
							MarkdownDescription: "MQCSP user ID",
							Computed:            true,
						},
						"csp_password_alias": schema.StringAttribute{
							MarkdownDescription: "MQCSP password alias",
							Computed:            true,
						},
						"heartbeat": schema.Int64Attribute{
							MarkdownDescription: "Channel heartbeat",
							Computed:            true,
						},
						"maximum_message_size": schema.Int64Attribute{
							MarkdownDescription: "Max message size",
							Computed:            true,
						},
						"cache_timeout": schema.Int64Attribute{
							MarkdownDescription: "Cache timeout",
							Computed:            true,
						},
						"ffst_size": schema.Int64Attribute{
							MarkdownDescription: "FFST file size",
							Computed:            true,
						},
						"ffst_rotate": schema.Int64Attribute{
							MarkdownDescription: "Number of FFST file rotations",
							Computed:            true,
						},
						"units_of_work": schema.Int64Attribute{
							MarkdownDescription: "Units-of-work",
							Computed:            true,
						},
						"automatic_backout": schema.BoolAttribute{
							MarkdownDescription: "Automatic backout",
							Computed:            true,
						},
						"backout_threshold": schema.Int64Attribute{
							MarkdownDescription: "Backout threshold",
							Computed:            true,
						},
						"backout_queue_name": schema.StringAttribute{
							MarkdownDescription: "Backout queue",
							Computed:            true,
						},
						"total_connection_limit": schema.Int64Attribute{
							MarkdownDescription: "Total connection limit",
							Computed:            true,
						},
						"initial_connections": schema.Int64Attribute{
							MarkdownDescription: "Initial connections",
							Computed:            true,
						},
						"sharing_conversations": schema.Int64Attribute{
							MarkdownDescription: "Sharing conversations",
							Computed:            true,
						},
						"ss_lkey": schema.StringAttribute{
							MarkdownDescription: "TLS key repository",
							Computed:            true,
						},
						"permit_insecure_servers": schema.BoolAttribute{
							MarkdownDescription: "Permit insecure connections",
							Computed:            true,
						},
						"ss_lcipher": schema.StringAttribute{
							MarkdownDescription: "TLS cipher specification",
							Computed:            true,
						},
						"ssl_cert_label": schema.StringAttribute{
							MarkdownDescription: "TLS certificate label",
							Computed:            true,
						},
						"convert_input": schema.BoolAttribute{
							MarkdownDescription: "Convert input",
							Computed:            true,
						},
						"auto_retry": schema.BoolAttribute{
							MarkdownDescription: "Automatic retry",
							Computed:            true,
						},
						"retry_interval": schema.Int64Attribute{
							MarkdownDescription: "Retry interval",
							Computed:            true,
						},
						"retry_attempts": schema.Int64Attribute{
							MarkdownDescription: "Retry attempts",
							Computed:            true,
						},
						"long_retry_interval": schema.Int64Attribute{
							MarkdownDescription: "Long retry interval",
							Computed:            true,
						},
						"reporting_interval": schema.Int64Attribute{
							MarkdownDescription: "Reporting interval",
							Computed:            true,
						},
						"alternate_user": schema.BoolAttribute{
							MarkdownDescription: "Alternate user",
							Computed:            true,
						},
						"local_address": schema.StringAttribute{
							MarkdownDescription: "Local address",
							Computed:            true,
						},
						"xml_manager": schema.StringAttribute{
							MarkdownDescription: "XML manager",
							Computed:            true,
						},
						"ssl_client": schema.StringAttribute{
							MarkdownDescription: "TLS client profile",
							Computed:            true,
						},
						"outbound_sni": schema.StringAttribute{
							MarkdownDescription: "Outbound SNI",
							Computed:            true,
						},
						"ocsp_check_extensions": schema.BoolAttribute{
							MarkdownDescription: "Check OCSP extensions",
							Computed:            true,
						},
						"ocsp_authentication": schema.StringAttribute{
							MarkdownDescription: "OCSP authentication",
							Computed:            true,
						},
						"cdp_check_extensions": schema.BoolAttribute{
							MarkdownDescription: "Check CDP extensions",
							Computed:            true,
						},
						"client_revocation_checks": schema.StringAttribute{
							MarkdownDescription: "Client revocation checking",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *MQManagerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *MQManagerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data MQManagerList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.MQManager{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.MQManager{}
	if value := res.Get(`MQManager`); value.Exists() {
		for _, v := range value.Array() {
			item := models.MQManager{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.MQManagerObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.MQManagerObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
