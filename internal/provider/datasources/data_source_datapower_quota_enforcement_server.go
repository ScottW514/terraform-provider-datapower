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
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var (
	_ datasource.DataSource              = &QuotaEnforcementServerDataSource{}
	_ datasource.DataSourceWithConfigure = &QuotaEnforcementServerDataSource{}
)

func NewQuotaEnforcementServerDataSource() datasource.DataSource {
	return &QuotaEnforcementServerDataSource{}
}

type QuotaEnforcementServerDataSource struct {
	pData *tfutils.ProviderData
}

func (d *QuotaEnforcementServerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_quota_enforcement_server"
}

func (d *QuotaEnforcementServerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "<p>On each DataPower Gateway, you configure the quota enforcement server to store the thresholds and associated metadata in memory or persist them on the RAID volume.</p><p>The quota enforcement server can work in standalone mode or peer group mode.</p><p>A peer group is collection of at least three nodes across which quota enforcement is implemented. In each peer group, one DataPower Gateway is the primary node and others are replicas. Failover might occur when the primary node becomes unavailable.</p><p>When you enable the peer group mode, the appropriate configuration properties are displayed. You must configure the connection among peers.</p><p>Based on your requirements for quota enforcement, you can enable or disable strict mode. Strict mode affects data-consistency across the peer group.</p><p>When strict mode is enabled, the following effects are caused: <ul><li>When the primary node is operational and when strict mode of all nodes in a peer group is enabled, threshold synchronization is more frequent to ensure data-consistency across the peer group. However, more network bandwidth is used. Therefore, strict mode is suitable for peers in the same data center.</li><li>When the primary node becomes unavailable, before failover occurs, the replica with enabled strict mode cannot process the request. <p>If service performance and availability are more important than data-consistency, you can disable strict mode for the replica so that this replica can process the request. The replica with disabled strict mode writes the threshold and associated metadata to the local data storage. After failover occurs, the connection is resumed between the replica and the new primary node. The threshold and associated metadata stored by the replica might be overwritten by the new primary node when the new primay node synchronizes the data to all replicas. Data-consistency might be affected across the peer group.</p></li></ul></p><p>By default, data is stored in memory and quota enforcement server works in standalone mode.</p>",
		Attributes: map[string]schema.Attribute{
			"provider_target": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Target host to retrieve this data from. If not set, provider will use the top level settings.", "", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "<p>The administrative state of the configuration.</p><ul><li>To make active, set to enabled.</li><li>To make inactive, set to disabled.</li></ul>",
				Computed:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: "Specifies a brief descriptive summary for the configuration.",
				Computed:            true,
			},
			"password_alias": schema.StringAttribute{
				MarkdownDescription: "Specifies the password alias to secure the data store. If not specified, a system default is used.",
				Computed:            true,
			},
			"raid_volume": schema.StringAttribute{
				MarkdownDescription: "<p>Specifies whether data storage is persistent or in-memory. <ul><li>For persistent storage, select the RAID volume that must be <tt>raid0</tt> .</li><li>For in-memory storage, do not select the RAID volume.</li></ul></p>",
				Computed:            true,
			},
			"server_port": schema.Int64Attribute{
				MarkdownDescription: "Indicates the listening port used by the GatewayScript module to communicate with the quota enforcement server. The default port value is 16379.",
				Computed:            true,
			},
			"monitor_port": schema.Int64Attribute{
				MarkdownDescription: "Indicates the listening port for operational state monitoring for the quota enforcement server. The default port is 26379.",
				Computed:            true,
			},
			"enable_peer_group": schema.BoolAttribute{
				MarkdownDescription: "Indicates whether the DataPower Gateway is enabled to work in peer group mode. By default, peer group mode is disabled, which indicates that the DataPower Gateway works in standalone mode.",
				Computed:            true,
			},
			"enable_ssl": schema.BoolAttribute{
				MarkdownDescription: "Indicates whether TLS is used to secure connection among the peers of the peer group. By default, the TLS is enabled.",
				Computed:            true,
			},
			"ssl_crypto_key": schema.StringAttribute{
				MarkdownDescription: "Indicates the key alias for the DataPower Gateway to authenticate a peer of the DataPower Gateway during the TLS handshake.",
				Computed:            true,
			},
			"ssl_crypto_certificate": schema.StringAttribute{
				MarkdownDescription: "Indicates the certificate alias that is sent to a peer when the DataPower Gateway negotiates a TLS connection with the peer.",
				Computed:            true,
			},
			"ip_address": schema.StringAttribute{
				MarkdownDescription: "<p>Indicates the IP address of the DataPower Gateway for other peers to connect to. The IP address can be the IP address on any interface of the DataPower Gateway and must be accessible by other peers in the peer group. The IP address cannot be 127.0.0.1, 0.0.0.0 or ::. This IP address uniquely identifies the DataPower Gateway.</p><p>You can use a local host alias instead of a static IP address. A host alias resolves a locally configured alias to a static IP address.</p>",
				Computed:            true,
			},
			"peers": schema.ListAttribute{
				MarkdownDescription: "<p>Specifies peers of the DataPower Gateway in the peer group. The DataPower Gateway connects to each peer in the order in which peers are added in the list. It is not necessary to specify all peers in the Peers list.</p><ul><li>When the DataPower Gateway connects to no peer in the list, this DataPower Gateway is the first active server and joins the peer group as the primary node.</li><li>When the DataPower Gateway connects to any peer in the list, it joins the peer group as a replica.</li></ul><p>You can use a local host alias instead of a static IP address. A host alias resolves a locally configured alias to a static IP address. Aliasing can help when you move configurations among appliances.</p><p>Note: Do not specify the IP address or hostname of this DataPower Gateway.</p>",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: "<p>Indicates the priority that is used to decide which replica is promoted to the primary node when failover occurs.</p><p>Enter a value in range 0 - 255. The default value is 100. The replica with the lowest priority number is promoted. A replica with the value of 0 can never be promoted.</p>",
				Computed:            true,
			},
			"strict_mode": schema.BoolAttribute{
				MarkdownDescription: "Based on your requirements for quota enforcement, enable or disable strict mode. By default, the strict mode is enabled.",
				Computed:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (d *QuotaEnforcementServerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *QuotaEnforcementServerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data models.QuotaEnforcementServer
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !d.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), d.pData.Client.GetTargetNames()))
		return
	}

	path := data.GetPath()

	res, err := d.pData.Client.Get(path, data.ProviderTarget)
	resFound := true
	if err != nil {
		if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		} else {
			resFound = false
		}
	}
	if resFound {
		data.FromBody(ctx, `QuotaEnforcementServer`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
