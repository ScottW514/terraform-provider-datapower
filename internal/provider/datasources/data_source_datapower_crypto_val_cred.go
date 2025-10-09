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

type CryptoValCredList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &CryptoValCredDataSource{}
	_ datasource.DataSourceWithConfigure = &CryptoValCredDataSource{}
)

func NewCryptoValCredDataSource() datasource.DataSource {
	return &CryptoValCredDataSource{}
}

type CryptoValCredDataSource struct {
	pData *tfutils.ProviderData
}

func (d *CryptoValCredDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_crypto_val_cred"
}

func (d *CryptoValCredDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "<p>Validation credentials authenticate certificates that are received from TLS peers. Validation credentials can be used to validate certificates that are used in digital signature and encryption operations.</p><p>a TLS client requires validation credentials only when it authenticates the certificate that is presented by the remote TLS server. The TLS standard does not require authentication of the server certificate.</p><p>a TLS server requires validation credentials only when it authenticates remote TLS clients. The TLS standard does not require authentication of the client certificate.</p>",
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
						"certificate": schema.ListAttribute{
							MarkdownDescription: "Specify the list of certificates for the validation credentials. Each certificate in the validation credentials is the certificate that a TLS peer might send or is the certificate of the certification authority (CA) that signed the certificate that is sent by a peer or is the root certificate.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"cert_validation_mode": schema.StringAttribute{
							MarkdownDescription: "Certificate validation mode",
							Computed:            true,
						},
						"use_crl": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to check certificate revocation lists (CRLs) during certificate validation. When enabled, CRLs are checked. Otherwise, CRLs are not checked.",
							Computed:            true,
						},
						"require_crl": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to mandate CRLs during certificate validation. When enabled, certificate validation fails if no CRL is available. Otherwise, validation succeeds independent of the availability of a CRL.",
							Computed:            true,
						},
						"crl_dp_handling": schema.StringAttribute{
							MarkdownDescription: "Specify the support of certificate extensions for X.509 certificate distribution points. This certificate extension specifies how to obtain CRL information. For more information, see RFC 2527 and RFC 3280.",
							Computed:            true,
						},
						"initial_policy_set": schema.ListAttribute{
							MarkdownDescription: "Specify the unique object identifiers for the certificate policy. <p>RFC 3280 refers to the input variable for certificate chain validation as <tt>user-initial-policy-set</tt> . These OIDs specify the allow values of certificate policies. To use this functionality, you need to require an explicit certificate policy. Otherwise, this set is used only if there are policy constraint extensions in the certificate chain.</p><p>By default, the initial certificate policy set consists of the single OID 2.5.29.32.0, which identifies <tt>anyPolicy</tt> .</p>",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"explicit_policy": schema.BoolAttribute{
							MarkdownDescription: "Specify support for the initial explicit policy variable as defined by RFC 3280. When enabled, the chain validation algorithm must end with a non-empty policy tree. Otherwise, the algorithm can end with an empty policy tree unless policy constraint extensions in the chain require an explicit policy.",
							Computed:            true,
						},
						"check_dates": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to check dates during certificate validation. This validation checks the current date and time against the <tt>notBefore</tt> and <tt>notAfter</tt> values in certificates and CRLs. When enabled, the date values are checked and expired certificates cause validation to fail. Otherwise, the date values are ignored and and do not cause validation to fail when a certificate is expired.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *CryptoValCredDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *CryptoValCredDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data CryptoValCredList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.CryptoValCred{
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
	l := []models.CryptoValCred{}
	if resFound {
		if value := res.Get(`CryptoValCred`); value.Exists() {
			for _, v := range value.Array() {
				item := models.CryptoValCred{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.CryptoValCredObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.CryptoValCredObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
