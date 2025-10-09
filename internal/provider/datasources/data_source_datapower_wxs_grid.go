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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type WXSGridList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &WXSGridDataSource{}
	_ datasource.DataSourceWithConfigure = &WXSGridDataSource{}
)

func NewWXSGridDataSource() datasource.DataSource {
	return &WXSGridDataSource{}
}

type WXSGridDataSource struct {
	pData *tfutils.ProviderData
}

func (d *WXSGridDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wxs_grid"
}

func (d *WXSGridDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The configuration of an eXtreme Scale Grid defines the connection details to an eXtreme Scale grid in an eXtreme Scale collective. To define this configuration, you must define the eXtreme Scale collective, the grid name, and the user and password for the user account who connects to the eXtreme Scale. If you need to secure connections to eXtreme Scale, you must assign a TLS Proxy Profile.",
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
							MarkdownDescription: "Specify a brief, but descriptive, summary of the configuration.",
							Computed:            true,
						},
						"collective": schema.StringAttribute{
							MarkdownDescription: "Specify the Load Balancer Group that contains members in the collective of eXtreme Scale. You must define at least one member in the collective.",
							Computed:            true,
						},
						"grid": schema.StringAttribute{
							MarkdownDescription: "Specify the name of the eXtreme Scale grid. The value cannot contain whitespace or the following characters: <tt>^ . \\ / , # $ @ : ; * ? &lt; > | = + &amp; % [ ] \" '</tt> .",
							Computed:            true,
						},
						"user_name": schema.StringAttribute{
							MarkdownDescription: "<p>Specify the user account of the eXtreme Scale user who connects to the eXtreme Scale collective. The value can be up to 64 characters in length and cannot be blank. You can use all alphanumeric characters and most special characters. You cannot use spaces or the following special characters: <tt># &lt;</tt> .</p><p>The user must have sufficient eXtreme Scale permissions to access the grid.</p>",
							Computed:            true,
						},
						"password_alias": schema.StringAttribute{
							MarkdownDescription: "Specify the password alias to use to look up the password of the eXtreme Scale user who connects to the eXtreme Scale collective.",
							Computed:            true,
						},
						"timeout": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum time to wait to establish a connection to an eXtreme Scale. If unable to establish a connection, the operation fails. Enter a value in the range 10 - 86400000. The default value is 1000.",
							Computed:            true,
						},
						"ssl_client": schema.StringAttribute{
							MarkdownDescription: "The TLS client profile to secure connections between the DataPower Gateway and its targets.",
							Computed:            true,
						},
						"encrypt": schema.BoolAttribute{
							MarkdownDescription: "Indicates whether the data in the eXtreme Scale data grid is encrypted. If encrypted, the data is encrypted, when writing to, and decrypted, when reading from, the eXtreme Scale data grid.",
							Computed:            true,
						},
						"encrypt_sskey": schema.StringAttribute{
							MarkdownDescription: "Specify the shared secret for PKCS #7 encryption and decryption. When writing data to the data grid, encrypts the data. When reading data from the eXtreme Scale data grid, decrypts the data.",
							Computed:            true,
						},
						"encrypt_alg": schema.StringAttribute{
							MarkdownDescription: "Specify the PKCS #7 algorithm for encryption and decryption. When writing data to the data grid, encrypts the data. When reading data from the eXtreme Scale data grid, decrypts the data.",
							Computed:            true,
						},
						"key_obfuscation": schema.BoolAttribute{
							MarkdownDescription: "Indicate whether to apply a hash algorithm to obfuscate keys before reading data from or writing data to the eXtreme Scale data grid.",
							Computed:            true,
						},
						"key_obfuscation_alg": schema.StringAttribute{
							MarkdownDescription: "Specify the hash algorithm to obfuscate keys before reading data from or writing data to the eXtreme Scale data grid.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *WXSGridDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *WXSGridDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data WXSGridList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.WXSGrid{
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
	l := []models.WXSGrid{}
	if resFound {
		if value := res.Get(`WXSGrid`); value.Exists() {
			for _, v := range value.Array() {
				item := models.WXSGrid{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.WXSGridObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.WXSGridObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
