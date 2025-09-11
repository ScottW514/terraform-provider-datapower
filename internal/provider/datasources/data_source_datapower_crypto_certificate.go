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

type CryptoCertificateList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &CryptoCertificateDataSource{}
	_ datasource.DataSourceWithConfigure = &CryptoCertificateDataSource{}
)

func NewCryptoCertificateDataSource() datasource.DataSource {
	return &CryptoCertificateDataSource{}
}

type CryptoCertificateDataSource struct {
	pData *tfutils.ProviderData
}

func (d *CryptoCertificateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_crypto_certificate"
}

func (d *CryptoCertificateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "A certificate provides an added layer of security by supplying an indirect reference to a file that contains a public key.",
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
						"filename": schema.StringAttribute{
							MarkdownDescription: "Specify the file that contains the public key. The name might be the same as the file that contains the private key. <ul><li>When in the public crypto area, the file is in the <tt>pubcert:</tt> directory.</li><li>When in the private crypto area, the file is in the <tt>cert:</tt> directory.</li><li>When in the public area, the file is in the <tt>local:</tt> directory. <p><b>Attention:</b> Any file in the <tt>local:</tt> directory can be downloaded or included in an export. Therefore, consider carefully before you store crypto files in this directory.</p></li><li>When retrieved from z/OS, the file is in the <tt>saf-cert:</tt> directory.</li><li>When on the SafeNet Luna HSM, the file is in a partition of the <tt>luna-cert:</tt> directory.</li></ul>",
							Computed:            true,
						},
						"alias": schema.StringAttribute{
							MarkdownDescription: "Password alias",
							Computed:            true,
						},
						"ignore_expiration": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to ignore the expiration date of the certificate. When enabled, the certificate remains in the <tt>up</tt> operational state outside of its expiration values. <p>Although the certificate is in the <tt>up</tt> state, other configurations that reference the certificate use the internal expiration values. In other words, the certificate itself is in the <tt>up</tt> state, but any configuration that references the certificate adheres to the internal expiration values.</p><ul><li>If the certificate is for certificate chain validation in validation credentials and the certificate is invalid, validation fails. Expired certificates cause the validation to fail unless you disable check dates in validation credentials.</li><li>If the certificate is in identification credentials, the DataPower Gateway sends the certificate to the peer. The peer can reject the certificate as invalid.</li></ul>",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *CryptoCertificateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *CryptoCertificateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data CryptoCertificateList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.CryptoCertificate{
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
	l := []models.CryptoCertificate{}
	if resFound {
		if value := res.Get(`CryptoCertificate`); value.Exists() {
			for _, v := range value.Array() {
				item := models.CryptoCertificate{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.CryptoCertificateObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.CryptoCertificateObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
