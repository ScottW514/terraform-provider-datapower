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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var (
	_ datasource.DataSource              = &B2BPersistenceDataSource{}
	_ datasource.DataSourceWithConfigure = &B2BPersistenceDataSource{}
)

func NewB2BPersistenceDataSource() datasource.DataSource {
	return &B2BPersistenceDataSource{}
}

type B2BPersistenceDataSource struct {
	pData *tfutils.ProviderData
}

func (d *B2BPersistenceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_b2bpersistence"
}

func (d *B2BPersistenceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Define the data storage for B2B transaction data.",
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "<p>The administrative state of the configuration.</p><ul><li>To make active, set to enabled.</li><li>To make inactive, set to disabled.</li></ul>",
				Computed:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: "Comments",
				Computed:            true,
			},
			"raid_volume": schema.StringAttribute{
				MarkdownDescription: "Specify the location of the data store. The location must be the <tt>raid0</tt> RAID volume. During firmware initialization, the <tt>raid0</tt> volume is associated with persistent storage.",
				Computed:            true,
			},
			"storage_size": schema.Int64Attribute{
				MarkdownDescription: "Specify the maximum size for the data store in MB. Enter a value in the range 1024 - 65536. The default is 1024. <p>This data store is for all B2B gateway services. These services store transaction metadata on the unencrypted partition of the RAID volume. These services store copies of the messages on the encrypted portion of the RAID volume.</p><p>The storage location for messages is defined on a service-by-service basis during the configuration of the B2B gateway.</p><p><b>Attention:</b> The maximum size for the persistent data store cannot be changed to a smaller value. Changing to a larger value might interrupt transactions that are in flight.</p>",
				Computed:            true,
			},
			"ha_enabled": schema.BoolAttribute{
				MarkdownDescription: "When on, the appliance is in active-passive high availability mode with the configured 'Alternate Host'.",
				Computed:            true,
			},
			"ha_other_hosts": models.GetDmB2BHAHostDataSourceSchema("Alternate host", "ha-other-hosts", ""),
			"ha_local_ip": schema.StringAttribute{
				MarkdownDescription: "Replication address",
				Computed:            true,
			},
			"ha_local_port": schema.Int64Attribute{
				MarkdownDescription: "Replication port",
				Computed:            true,
			},
			"ha_virtual_ip": schema.StringAttribute{
				MarkdownDescription: "Virtual IP address",
				Computed:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (d *B2BPersistenceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *B2BPersistenceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data models.B2BPersistence
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := d.pData.Client.Get(data.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	data.FromBody(ctx, `B2BPersistence`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
