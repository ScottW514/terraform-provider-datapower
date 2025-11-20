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

type AAAJWTGeneratorList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &AAAJWTGeneratorDataSource{}
	_ datasource.DataSourceWithConfigure = &AAAJWTGeneratorDataSource{}
)

func NewAAAJWTGeneratorDataSource() datasource.DataSource {
	return &AAAJWTGeneratorDataSource{}
}

type AAAJWTGeneratorDataSource struct {
	pData *tfutils.ProviderData
}

func (d *AAAJWTGeneratorDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaa_jwt_generator"
}

func (d *AAAJWTGeneratorDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "<p>The JSON Web Token (JWT) Generator specifies the JWT content and the cryptographic methods, such as signing and encryption methods, used for generating a JWT during the AAA postprocessing phase.</p>",
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
							MarkdownDescription: "A descriptive summary for the JWT Generator configuration.",
							Computed:            true,
						},
						"issuer": schema.StringAttribute{
							MarkdownDescription: "The issuer claim, \"iss\", identifies the principal that issues the JWT. The maximum length is 256 characters. The default value is <tt>idg</tt> .",
							Computed:            true,
						},
						"duration": schema.Int64Attribute{
							MarkdownDescription: "The validity period identifies the expiration time, \"exp\" claim. Enter a value in the range 1 - 31622400. The default value is 3600.",
							Computed:            true,
						},
						"additional_claims": models.GetDmJWTClaimsDataSourceSchema("<p>Additional JWT claims, such as audience \"aud\" claim, not before \"nbf\" claim, issued at \"iat\" claim, JWT ID \"jit\" claim, \"nonce\" claim, and custom claim, can be added for JWT.</p><p>The subject, \"sub\" claim is added by default. You can override the subject claim value by specifying the \"sub\" claim in the Custom claims field.</p>", "add-claims", ""),
						"audience": schema.ListAttribute{
							MarkdownDescription: "The audience, \"aud\", claim identifies the recipients that the JWT is intended for. The maximum length of the Audience claim is 256 characters.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"not_before": schema.Int64Attribute{
							MarkdownDescription: "The not before claim, \"nbf\", indicates the time before which the JWT must not be accepted for processing. Enter a value in the range 0 - 480. The default value is 0.",
							Computed:            true,
						},
						"custom_claims": schema.StringAttribute{
							MarkdownDescription: "The GatewayScript or XSLT file is processed to specify the custom claim. The GatewayScript or XSLT file must be stored in the <tt>local:</tt> or <tt>store:</tt> directory.",
							Computed:            true,
						},
						"gen_method": models.GetDmJWTGenMethodDataSourceSchema("The signing and encryption methods can be used to secure and generate a JWT.", "generate-method", ""),
						"sign_algorithm": schema.StringAttribute{
							MarkdownDescription: "Various signing algorithms can be used to generate the JWT signature, such as HS256, HS384, HS512, RS256, RS384, and RS512. The default value is RS256.",
							Computed:            true,
						},
						"sign_key": schema.StringAttribute{
							MarkdownDescription: "The key alias can be used to sign the JWT. You can get a key alias by configuring the Crypto Key.",
							Computed:            true,
						},
						"sign_sskey": schema.StringAttribute{
							MarkdownDescription: "The shared secret key alias can be used to sign the JWT. You can get the shared secret key alias by configuring the Crypto Shared Secret Key.",
							Computed:            true,
						},
						"enc_algorithm": schema.StringAttribute{
							MarkdownDescription: "Various encryption algorithms can be used to encrypt the JWT, such as A128CBC-HS256, A192CBC-HS384, A256CBC-HS512, A128GCM, A192GCM, and A256GCM. The default value is A128CBC-HS256.",
							Computed:            true,
						},
						"encrypt_algorithm": schema.StringAttribute{
							MarkdownDescription: "Various key management algorithms can be used to encrypt the JWT, such as RSA1_5, RSA-OAEP, RSA-OAEP-256, A128KW, A192KW, A256KW, and dir. The default value is RSA1_5.",
							Computed:            true,
						},
						"encrypt_certificate": schema.StringAttribute{
							MarkdownDescription: "The certificate alias can be used to encrypt the JWT. You can get the certificate alias by configuring the Crypto Certificate.",
							Computed:            true,
						},
						"encrypt_sskey": schema.StringAttribute{
							MarkdownDescription: "The shared secret key alias can be used to encrypt the JWT. You can get the shared secret key alias by configuring the Crypto Shared Secret Key.",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *AAAJWTGeneratorDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *AAAJWTGeneratorDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data AAAJWTGeneratorList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.AAAJWTGenerator{
		AppDomain: data.AppDomain,
	}

	path := o.GetPath()
	if !data.Id.IsNull() {
		path = path + "/" + data.Id.ValueString()
	}

	res, err := d.pData.Client.Get(path)
	resFound := true
	if err != nil {
		if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(d.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString())
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		} else {
			resFound = false
		}
	}
	l := []models.AAAJWTGenerator{}
	if resFound {
		if value := res.Get(`AAAJWTGenerator`); value.Exists() {
			for _, v := range value.Array() {
				item := models.AAAJWTGenerator{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.AAAJWTGeneratorObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.AAAJWTGeneratorObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
