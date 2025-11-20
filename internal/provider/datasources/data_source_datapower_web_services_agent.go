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
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var (
	_ datasource.DataSource              = &WebServicesAgentDataSource{}
	_ datasource.DataSourceWithConfigure = &WebServicesAgentDataSource{}
)

func NewWebServicesAgentDataSource() datasource.DataSource {
	return &WebServicesAgentDataSource{}
}

type WebServicesAgentDataSource struct {
	pData *tfutils.ProviderData
}

func (d *WebServicesAgentDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_web_services_agent"
}

func (d *WebServicesAgentDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The Web Services Management agent provides manageability for Web Services by providing status, metrics, and transaction history to external management stations.",
		Attributes: map[string]schema.Attribute{
			"app_domain": schema.StringAttribute{
				MarkdownDescription: "The name of the application domain the object belongs to",
				Required:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "<p>The administrative state of the configuration.</p><ul><li>To make active, set to enabled.</li><li>To make inactive, set to disabled.</li></ul>",
				Computed:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: "Comments",
				Computed:            true,
			},
			"max_records": schema.Int64Attribute{
				MarkdownDescription: "Specify the maximum number of transaction records to buffer. Buffering controls the behavior of the agent when there are no registered consumers of transaction events. Records are accumulated until they reach the configured threshold. After this threshold is reached, new records are dropped. The loss of records are visible to web service managers that understand the concept.",
				Computed:            true,
			},
			"max_memory_kb": schema.Int64Attribute{
				MarkdownDescription: "Specify the maximum memory in KB to buffer transaction records. Buffering controls the behavior of the agent when there are no registered consumers of transaction events. Records are accumulated until they reach the configured threshold. After this threshold is reached, new records are dropped. The loss of records are visible to web service managers that understand the concept.",
				Computed:            true,
			},
			"capture_mode": schema.StringAttribute{
				MarkdownDescription: "Specify the mode to capture messages for further analysis. Because not all Web Services Management protocols can accommodate full message-capture, configure this property only when the spooler can forward full messages. <p>Full message-capture incurs a performance penalty.</p>",
				Computed:            true,
			},
			"mediation_metrics": schema.BoolAttribute{
				MarkdownDescription: "Specify whether to collect metrics about mediation enforcement. The default behavior is to not collect metrics.",
				Computed:            true,
			},
			"max_payload_size_kb": schema.Int64Attribute{
				MarkdownDescription: "Specify the maximum total payload size in KB of a buffered transaction record. The total payload size is the sum of the payloads that are collected at the following points. <ul><li>When the service accepts the request.</li><li>When the service sends the processed request to its target.</li><li>When the service accepts the response.</li><li>When the service sends the processed response to the client.</li></ul><p>A record is dropped when its total payload size exceeds the maximum value. The default value is 0, which indicates no limit.</p>",
				Computed:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (d *WebServicesAgentDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *WebServicesAgentDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data models.WebServicesAgent
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	path := data.GetPath()

	res, err := d.pData.Client.Get(path)
	resFound := true
	if err != nil {
		if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(d.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString())
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
	if resFound {
		data.FromBody(ctx, `WebServicesAgent`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
