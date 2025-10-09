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
	_ datasource.DataSource              = &WebGUIDataSource{}
	_ datasource.DataSourceWithConfigure = &WebGUIDataSource{}
)

func NewWebGUIDataSource() datasource.DataSource {
	return &WebGUIDataSource{}
}

type WebGUIDataSource struct {
	pData *tfutils.ProviderData
}

func (d *WebGUIDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_web_gui"
}

func (d *WebGUIDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Modify the configuration of the web management service. <p>If you do not assign a TLS profile, the service uses a profile with a self-signed certificate.</p>",
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "<p>The administrative state of the configuration.</p><ul><li>To make active, set to enabled.</li><li>To make inactive, set to disabled.</li></ul>",
				Computed:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: "Comments",
				Computed:            true,
			},
			"local_port": schema.Int64Attribute{
				MarkdownDescription: "Specify the TCP port that the service monitors. The default value is 9090.",
				Computed:            true,
			},
			"user_agent": schema.StringAttribute{
				MarkdownDescription: "Custom user agent",
				Computed:            true,
			},
			"save_config_overwrites": schema.BoolAttribute{
				MarkdownDescription: "Specify whether saving the configuration overwrites the startup configuration. <ul><li>When enable, saving the configuration overwrites the startup configuration with the running configuration.</li><li>When not enabled, a manual step is required that prevents overwriting the manually edited startup configuration.</li></ul>",
				Computed:            true,
			},
			"idle_timeout": schema.Int64Attribute{
				MarkdownDescription: "Specify the time after which to invalidate idle sessions. When invalidated, requires reauthentication. Enter a value in the range 0 - 65535, in seconds. A value of 0 disables the timer.",
				Computed:            true,
			},
			"acl": schema.StringAttribute{
				MarkdownDescription: "Edit the <tt>web-mgmt</tt> access control list to define the client IP addresses to allow or deny.",
				Computed:            true,
			},
			"ssl_server_config_type": schema.StringAttribute{
				MarkdownDescription: "Custom TLS server type",
				Computed:            true,
			},
			"ssl_server": schema.StringAttribute{
				MarkdownDescription: "Custom TLS server profile",
				Computed:            true,
			},
			"ssl_sni_server": schema.StringAttribute{
				MarkdownDescription: "Custom TLS SNI server profile",
				Computed:            true,
			},
			"enable_sts": schema.BoolAttribute{
				MarkdownDescription: "Specify whether to enable HTTP Strict Transport Security headers. When enabled, responses inject HTTP Strict Transport Security headers.",
				Computed:            true,
			},
			"local_address": schema.StringAttribute{
				MarkdownDescription: "<p>Enter a host alias or the IP address that the service listens on. Host aliases can ease migration tasks among appliances.</p><ul><li>0 or 0.0.0.0 indicates all configured IPv4 addresses.</li><li>:: indicates all configured IPv4 and IPv6 addresses.</li></ul><p><b>Attention:</b> For management services, the value of 0.0.0.0 or :: is a security risk. Use an explicit IP address to isolate management traffic from application data traffic.</p>",
				Computed:            true,
			},
		},
	}
}

func (d *WebGUIDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *WebGUIDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data models.WebGUI
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
		data.FromBody(ctx, `WebGUI`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
