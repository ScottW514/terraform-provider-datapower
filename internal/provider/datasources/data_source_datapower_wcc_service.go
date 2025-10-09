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
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type WCCServiceList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &WCCServiceDataSource{}
	_ datasource.DataSourceWithConfigure = &WCCServiceDataSource{}
)

func NewWCCServiceDataSource() datasource.DataSource {
	return &WCCServiceDataSource{}
}

type WCCServiceDataSource struct {
	pData *tfutils.ProviderData
}

func (d *WCCServiceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wcc_service"
}

func (d *WCCServiceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Use the WebSphere Cell configuration to retrieve configuration information and runtime information from a Network Deployment or Virtual Enterprise product. Load balancer groups use this information to automate configuration and updates.",
		Attributes: map[string]schema.Attribute{
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
						"odc_info_hostname": schema.StringAttribute{
							MarkdownDescription: "Host name or IP address of the Deployment Manager. It is used to communicate with the ODCInfo servlet. The ODCInfo servlet must be installed in a WebSphere Application Server environment.",
							Computed:            true,
						},
						"odc_info_port": schema.Int64Attribute{
							MarkdownDescription: "Specifies the listening port of the ODCInfo servlet on the deployment manager. To determine the listening port number, log in to the WebSphere Application Server administrative console and click System Administration > Deployment Manager > ports. From the list, use the port associated with WC_adminhost for HTTP or the port associated with WC_adminhost_secure for HTTPS.",
							Computed:            true,
						},
						"update_type": schema.StringAttribute{
							MarkdownDescription: "Specifies the method for retrieving WebSphere information. The method defines the responsiveness of the updates received by the DataPower Gateway. The method also affects the amount of system resource used for providing those updates. The default is poll.",
							Computed:            true,
						},
						"time_interval": schema.Int64Attribute{
							MarkdownDescription: "<p>If the update method is poll, the time interval specifies the amount of time in seconds between poll requests.</p><p>If the update method is subscribe, the time interval specifies the maximum duration of the request in seconds.</p><p>Use any value of 1 - 86400. The default is 10.</p>",
							Computed:            true,
						},
						"ssl_client_config_type": schema.StringAttribute{
							MarkdownDescription: "The TLS profile type to secure connections between the DataPower Gateway and its targets.",
							Computed:            true,
						},
						"ssl_client": schema.StringAttribute{
							MarkdownDescription: "The TLS client profile to secure connections between the DataPower Gateway and its targets.",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *WCCServiceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *WCCServiceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data WCCServiceList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.WCCService{
		AppDomain: data.AppDomain,
	}

	path := o.GetPath()
	if !data.Id.IsNull() {
		path = path + "/" + data.Id.ValueString()
	}

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
	l := []models.WCCService{}
	if resFound {
		if value := res.Get(`WCCService`); value.Exists() {
			for _, v := range value.Array() {
				item := models.WCCService{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.WCCServiceObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.WCCServiceObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
