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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var (
	_ datasource.DataSource              = &APISecurityTokenManagerDataSource{}
	_ datasource.DataSourceWithConfigure = &APISecurityTokenManagerDataSource{}
)

func NewAPISecurityTokenManagerDataSource() datasource.DataSource {
	return &APISecurityTokenManagerDataSource{}
}

type APISecurityTokenManagerDataSource struct {
	pData *tfutils.ProviderData
}

func (d *APISecurityTokenManagerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_security_token_manager"
}

func (d *APISecurityTokenManagerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "<p>The API security token manager provides the storage configuration for security objects, which include OAuth providers and user security. Each domain has the <tt>default</tt> API security token manager. This instance is used by the domain to store and manage API details.</p><p>The API security token service uses gateway-peering instances for the internal and external token stores.</p>",
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
			"gateway_peering": schema.StringAttribute{
				MarkdownDescription: "Specify the gateway-peering instance to store and manage internal OAuth token data in this domain. Native OAuth tokens that are managed by an external token management service are not stored in this gateway-peering instance. This gateway-peering instance must be configured to persist data across a restart.",
				Computed:            true,
			},
			"gateway_peering_external": schema.StringAttribute{
				MarkdownDescription: "Specify the gateway-peering instance to store and manage responses from external OAuth token management services in this domain. This gateway-peering instance does not require that data persist across a restart.",
				Computed:            true,
			},
			"expired_token_cleanup_interval": schema.Int64Attribute{
				MarkdownDescription: "Cleanup interval",
				Computed:            true,
			},
		},
	}
}

func (d *APISecurityTokenManagerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *APISecurityTokenManagerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data models.APISecurityTokenManager
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	path := data.GetPath()

	res, err := d.pData.Client.Get(path)
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
		data.FromBody(ctx, `APISecurityTokenManager`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
