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

type MQv9PlusMFTSourceProtocolHandlerList struct {
	ProviderTarget types.String `tfsdk:"provider_target"`
	Id             types.String `tfsdk:"id"`
	AppDomain      types.String `tfsdk:"app_domain"`
	Result         types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &MQv9PlusMFTSourceProtocolHandlerDataSource{}
	_ datasource.DataSourceWithConfigure = &MQv9PlusMFTSourceProtocolHandlerDataSource{}
)

func NewMQv9PlusMFTSourceProtocolHandlerDataSource() datasource.DataSource {
	return &MQv9PlusMFTSourceProtocolHandlerDataSource{}
}

type MQv9PlusMFTSourceProtocolHandlerDataSource struct {
	pData *tfutils.ProviderData
}

func (d *MQv9PlusMFTSourceProtocolHandlerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mqv9_plus_mft_source_protocol_handler"
}

func (d *MQv9PlusMFTSourceProtocolHandlerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Configure the IBM MQ v9+ MFT handle to manage IBM MQ MFT protocol communications.",
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
						"queue_manager": schema.StringAttribute{
							MarkdownDescription: "Specify the name of the queue manager that provides messaging services for communicating applications by periodically monitoring or polling queues and by ensuring that messages are directed to the correct receive queue or routed to another queue manager. The local queue manager corresponds to a queue manager running on another host on the network.",
							Computed:            true,
						},
						"get_queue": schema.StringAttribute{
							MarkdownDescription: "Specify the name of the get queue associated with the queue manager. The handler gets messages from this queue.",
							Computed:            true,
						},
						"get_message_options": schema.Int64Attribute{
							MarkdownDescription: "Specify the cumulative value of the MQGET options that are applicable to an IBM MQ message in decimal or hex format. The value is passed directly to the IBM MQ API. The default value is 32769, which is the decimal value for the <tt>MQGMO_WAIT</tt> and <tt>MQGMO_LOGICAL_ORDER</tt> options.",
							Computed:            true,
						},
						"concurrent_connections": schema.Int64Attribute{
							MarkdownDescription: "Specify the number of concurrent IBM MQ conversations to allocate. The default value is 1 but can be increased to improve performance.",
							Computed:            true,
						},
						"polling_interval": schema.Int64Attribute{
							MarkdownDescription: "Specify the duration in seconds to wait after processing all messages before attempting to retrieve messages from the get queue.",
							Computed:            true,
						},
						"retrieve_backout_settings": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to retrieve backout setting from the IBM MQ server. <p>When enabled, retrieves the <b>Backout threshold</b> and <b>Backout requeue queue name</b> settings from the IBM MQ server and checks these values. On a reattempt, the handler uses the higher priority backout settings from the server. If the server does not contain backout settings, The handler uses any existing backout values, either empty or populated, from the local IBM MQ queue manager. If there are no backout settings, the backout function is disabled.</p><p>When an alias queue is used, its attributes are retrieved, not those of the base queue.</p>",
							Computed:            true,
						},
						"ignore_backout_errors": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to ignore backout errors. <ul><li>>When enabled, ignore the error in sending the transfer to the backout queue and commit the transfer from the get queue.</li><li>When not enabled roll back and retry the transfer. This setting is the default value.</li></ul>",
							Computed:            true,
						},
						"use_qm_name_in_url": schema.BoolAttribute{
							MarkdownDescription: "Specify whether the var://service/URL-in variable returns the name of the local queue manager or queue manager group when this configuration defines a queue manager group as the queue manager. <ul><li>When enabled, the variable returns the name of the queue manager.</li><li>When not enabled, the variable returns the name of the queue manager group. This setting is the default value.</li></ul>",
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

func (d *MQv9PlusMFTSourceProtocolHandlerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *MQv9PlusMFTSourceProtocolHandlerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data MQv9PlusMFTSourceProtocolHandlerList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !d.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), d.pData.Client.GetTargetNames()))
		return
	}
	o := models.MQv9PlusMFTSourceProtocolHandler{
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
	l := []models.MQv9PlusMFTSourceProtocolHandler{}
	if resFound {
		if value := res.Get(`MQv9PlusMFTSourceProtocolHandler`); value.Exists() {
			for _, v := range value.Array() {
				item := models.MQv9PlusMFTSourceProtocolHandler{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.MQv9PlusMFTSourceProtocolHandlerObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.MQv9PlusMFTSourceProtocolHandlerObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
