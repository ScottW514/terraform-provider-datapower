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

type ZosNSSClientList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &ZosNSSClientDataSource{}
	_ datasource.DataSourceWithConfigure = &ZosNSSClientDataSource{}
)

func NewZosNSSClientDataSource() datasource.DataSource {
	return &ZosNSSClientDataSource{}
}

type ZosNSSClientDataSource struct {
	pData *tfutils.ProviderData
}

func (d *ZosNSSClientDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zos_nss_client"
}

func (d *ZosNSSClientDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The z/OS NSS client object provides the parameters for authentication with SAF on a z/OS Communications Server.",
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
							MarkdownDescription: "A descriptive summary for the configuration.",
							Computed:            true,
						},
						"remote_address": schema.StringAttribute{
							MarkdownDescription: "Specifies IP address or host name of the NSS server. In conjunction with the Remote Port, identifies the host and listening port of the NSS server. The NSS server must have the XMLAppliance discipline support enabled.",
							Computed:            true,
						},
						"remote_port": schema.Int64Attribute{
							MarkdownDescription: "Specifies NSS server port. In conjunction with the Remote Address, identifies the host and listening port of the NSS server.",
							Computed:            true,
						},
						"client_id": schema.StringAttribute{
							MarkdownDescription: "<p>Specifies the client ID to be used for registration with the NSS server. Minimum length is 1. Maximum length is 24.</p><p>Valid characters are:</p><p><ul><li>a through z</li><li>A through Z</li><li>0 through 9</li><li>_ (underscore)</li><li>- (dash)</li></ul></p><p>Embedded spaces are invalid.</p><p>The Client ID identifies the client ID to register the appliance with the NSS server. The NSS client ID is a unique string used by the NSS Server to track clients.</p><p>The Client ID does not have to correspond to any preexisting object. It is provided to the server at the time of registration. If another client attempts to register with the same client ID to the same NSS Server, the NSS server will send a heartbeat to the first client. If the first client responds to the heartbeat, the second client's registration will be rejected. If the first client does not respond, the connect to the first client will be severed and the second client will be registered.</p>",
							Computed:            true,
						},
						"system_name": schema.StringAttribute{
							MarkdownDescription: "<p>Specifies a name for the NSS client. Minimum length is 1. Maximum length is 8.</p><p>Valid characters are:</p><p><ul><li>a through z</li><li>A through Z</li><li>0 through 9</li><li>_ (underscore)</li><li>- (dash)</li></ul></p><p>Embedded spaces are invalid.</p><p>The System Name identifies the NSS client to the NSS server. NSS server commands identify NSS clients by system name in the output when displaying information for connected NSS clients.</p>",
							Computed:            true,
						},
						"user_name": schema.StringAttribute{
							MarkdownDescription: "<p>Specifies a user name to use to authenticate to the NSS server. Minimum length is 1. Maximum length is 8.</p><p>Valid characters are:</p><p><ul><li>a through z</li><li>A through Z</li><li>0 through 9</li><li>_ (underscore)</li><li>- (dash)</li></ul></p><p>Embedded spaces are invalid.</p><p>The user name must match an existing user ID on the NSS Server.</p>",
							Computed:            true,
						},
						"password_alias": schema.StringAttribute{
							MarkdownDescription: "<p>Specifies the password alias of the password to use to authenticate to the NSS server.</p><p>The associated password is used in conjunction with the value provided by the User Name.</p>",
							Computed:            true,
						},
						"ssl_client_config_type": schema.StringAttribute{
							MarkdownDescription: "The TLS profile type to secure connections between the DataPower Gateway and the NSS server.",
							Computed:            true,
						},
						"ssl_client": schema.StringAttribute{
							MarkdownDescription: "The TLS client profile to secure connections between the DataPower Gateway and the NSS server.",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *ZosNSSClientDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *ZosNSSClientDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data ZosNSSClientList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.ZosNSSClient{
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
	l := []models.ZosNSSClient{}
	if resFound {
		if value := res.Get(`ZosNSSClient`); value.Exists() {
			for _, v := range value.Array() {
				item := models.ZosNSSClient{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.ZosNSSClientObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.ZosNSSClientObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
