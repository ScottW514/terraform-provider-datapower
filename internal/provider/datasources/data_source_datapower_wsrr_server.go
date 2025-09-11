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

type WSRRServerList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &WSRRServerDataSource{}
	_ datasource.DataSourceWithConfigure = &WSRRServerDataSource{}
)

func NewWSRRServerDataSource() datasource.DataSource {
	return &WSRRServerDataSource{}
}

type WSRRServerDataSource struct {
	pData *tfutils.ProviderData
}

func (d *WSRRServerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wsrr_server"
}

func (d *WSRRServerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Configuration required to locate and access WSRR (WebSphere Service Registry and Repository).",
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
						"soap_url": schema.StringAttribute{
							MarkdownDescription: "Specify the URL of SOAP API endpoint on the WSRR server. The port in the URL depends on whether the server has enabled global security. A typical default URL is similar to <tt>https://192.18.1.120:9443/WSRRCoreSDO/services/WSRRCoreSDOPort</tt> or <tt>http://192.18.1.120:9080/WSRRCoreSDO/services/WSRRCoreSDOPort</tt> .",
							Computed:            true,
						},
						"server_prefix": schema.StringAttribute{
							MarkdownDescription: "WSRR instance prefix",
							Computed:            true,
						},
						"username": schema.StringAttribute{
							MarkdownDescription: "Specify the username for authentication with the server. Leave blank when authentication is not required.",
							Computed:            true,
						},
						"password_alias": schema.StringAttribute{
							MarkdownDescription: "Specify the password alias for authentication with the server. The password alias references the password. Leave blank when authentication is not required.",
							Computed:            true,
						},
						"ssl_client_config_type": schema.StringAttribute{
							MarkdownDescription: "TLS client type",
							Computed:            true,
						},
						"ssl_client": schema.StringAttribute{
							MarkdownDescription: "TLS client profile",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *WSRRServerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *WSRRServerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data WSRRServerList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.WSRRServer{
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
	l := []models.WSRRServer{}
	if resFound {
		if value := res.Get(`WSRRServer`); value.Exists() {
			for _, v := range value.Array() {
				item := models.WSRRServer{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.WSRRServerObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.WSRRServerObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
