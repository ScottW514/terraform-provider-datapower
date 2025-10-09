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
	_ datasource.DataSource              = &CertMonitorDataSource{}
	_ datasource.DataSourceWithConfigure = &CertMonitorDataSource{}
)

func NewCertMonitorDataSource() datasource.DataSource {
	return &CertMonitorDataSource{}
}

type CertMonitorDataSource struct {
	pData *tfutils.ProviderData
}

func (d *CertMonitorDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cert_monitor"
}

func (d *CertMonitorDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The certificate monitor is a task that checks the expiration date of all certificates. The defined values establish both the polling frequency and a notification window during which the monitor generates log messages that record when certificates are nearing their expiration date. The certificate monitor scans all certificates when first enabled.",
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "<p>The administrative state of the configuration.</p><ul><li>To make active, set to enabled.</li><li>To make inactive, set to disabled.</li></ul>",
				Computed:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: "Comments",
				Computed:            true,
			},
			"polling_interval": schema.Int64Attribute{
				MarkdownDescription: "Specify the frequency that the certificate monitor scans certificates for their expiration. The certificate monitor scans for expiry at each restart. If today is Monday and you set the frequency to three days and you restart on Wednesday, the next scan is three days later. Enter a value in the range 1 - 65535. The default value is 1.",
				Computed:            true,
			},
			"reminder_time": schema.Int64Attribute{
				MarkdownDescription: "Specify the notification window before certificate expiration to start the logging of certificate expiration messages. For example, the value 21 specifies that all certificates to expire in three weeks or less generate a log message at the defined priority. Enter a value in the range 1 - 65535. The default value is 30.",
				Computed:            true,
			},
			"log_level": schema.StringAttribute{
				MarkdownDescription: "Specify the priority to log messages about the impending expiration date of certificates. By default, messages are logged as warnings.",
				Computed:            true,
			},
			"disable_expired_certs": schema.BoolAttribute{
				MarkdownDescription: "Specify the behavior for expired certificates. By default, expired certificate objects are not disabled, which allows the use of expired certificates. When enabled, prevents the use of expired certificates either directly or through inheritance, which disables the use of any objects the reference expired certificates.",
				Computed:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (d *CertMonitorDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *CertMonitorDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data models.CertMonitor
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
		data.FromBody(ctx, `CertMonitor`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
