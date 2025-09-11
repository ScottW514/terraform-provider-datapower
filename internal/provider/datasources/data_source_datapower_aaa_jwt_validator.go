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

type AAAJWTValidatorList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &AAAJWTValidatorDataSource{}
	_ datasource.DataSourceWithConfigure = &AAAJWTValidatorDataSource{}
)

func NewAAAJWTValidatorDataSource() datasource.DataSource {
	return &AAAJWTValidatorDataSource{}
}

type AAAJWTValidatorDataSource struct {
	pData *tfutils.ProviderData
}

func (d *AAAJWTValidatorDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaa_jwt_validator"
}

func (d *AAAJWTValidatorDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "<p>The JSON Web Token (JWT) Validator specifies credentials and different methods to validate a JWT.</p>",
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
							MarkdownDescription: "A descriptive summary for the JWT Validator configuration.",
							Computed:            true,
						},
						"issuer": schema.StringAttribute{
							MarkdownDescription: "The optional issuer claim. The \"iss\" PCRE can be used to verify the JWT. The maximum length of the value is 256 characters.",
							Computed:            true,
						},
						"aud": schema.StringAttribute{
							MarkdownDescription: "The optional audience claim. The \"aud\" PCRE can be used to verify the JWT. The maximum length of the value is 256 characters.",
							Computed:            true,
						},
						"val_method": models.GetDmJWTValMethodDataSourceSchema("Various methods can be used to validate the JWT. You can decrypt the JWT, verify the JWT signature, and process a custom GatewayScript or XSLT file for further processing.", "validate-method", ""),
						"customized_script": schema.StringAttribute{
							MarkdownDescription: "A custom GatewayScript or XSLT file is processed to validate the JWT. The GatewayScript or XSLT file must be stored in the <tt>local:</tt> (the default) or <tt>store:</tt> directory. This field is meaningful when you select <tt>Custom processing</tt> in the Validation method field.",
							Computed:            true,
						},
						"decrypt_credential_type": schema.StringAttribute{
							MarkdownDescription: "Various decryption methods (such as PKIX, shared secret key, JSON Web Key (JWK), custom processing, remotely retrieve JWK) can be used to decrypt the JWT. The default method is PKIX. This field is meaningful when you select <tt>Decrypt</tt> in the Validation method field.",
							Computed:            true,
						},
						"decrypt_key": schema.StringAttribute{
							MarkdownDescription: "The private key can be used to decrypt the JWT. You can get the key alias by configuring the Crypto Key. This field is meaningful when you select <tt>Decrypt</tt> in the Validation Method field and choose <tt>PKIX</tt> from the Decrypt method list.",
							Computed:            true,
						},
						"decrypt_s_secret": schema.StringAttribute{
							MarkdownDescription: "The shared secret key can be used to decrypt the JWT. You can get the shared secret key alias by configuring the Crypto Shared Secret Key. This field is meaningful when you select <tt>Decrypt</tt> in the Validation method field and choose <tt>Shared secret</tt> from the Decrypt method list.",
							Computed:            true,
						},
						"decrypt_jwk": schema.StringAttribute{
							MarkdownDescription: "The file containing the JWK or key set is fetched to decrypt the JWT. The file must be stored in the local: or store: directory. This field is meaningful when you select <tt>Decrypt</tt> in the Validation method field and choose <tt>JWK</tt> from the Decrypt method list.",
							Computed:            true,
						},
						"decrypt_fetch_cred_url": schema.StringAttribute{
							MarkdownDescription: "The URL indicates the source location where the decryption credentials can be fetched for decrypting the JWT. The URL must be in the format of http or https. By default, the URL is http://example.com/v3/key. This field is meaningful when you choose <tt>Decrypt</tt> in the Validation method field and choose <tt>Remotely retrieve JWK</tt> from the Decrypt method list.",
							Computed:            true,
						},
						"decrypt_fetch_cred_ssl_profile": schema.StringAttribute{
							MarkdownDescription: "The TLS client profile is specified for fetching the decryption credentials. This field is meaningful when you select <tt>Decrypt</tt> in the Validation method field and choose <tt>Remotely retrieve JWK</tt> from the Decrypt method list.",
							Computed:            true,
						},
						"validate_custom": schema.StringAttribute{
							MarkdownDescription: "A custom GatewayScript or XSLT file provides the key material information to decrypt or verify the JWT. This field is meaningful when you select <tt>Custom</tt> for the Decrypt method or Verify method list.",
							Computed:            true,
						},
						"verify_credential_type": schema.StringAttribute{
							MarkdownDescription: "Various methods (such as PKIX, shared secret key, JWK, custom processing, remotely retrieve JWK) can be used to verify the JWT signature. The default method is PKIX. This field is meaningful when you select <tt>Verify</tt> in the Validation method field.",
							Computed:            true,
						},
						"verify_certificate": schema.StringAttribute{
							MarkdownDescription: "The certificate can be used to verify the JWT signature. You can get the certificate by configuring the Crypto Certificate. This field is meaningful when you select <tt>Verify</tt> in the Validation method field and choose <tt>PKIX</tt> from the Verify method field.",
							Computed:            true,
						},
						"verify_certificate_against_val_cred": schema.BoolAttribute{
							MarkdownDescription: "You decide whether to use validation credentials to verify the JWT signature. This field is meaningful when you select <tt>Verify</tt> in the Validation method field and choose <tt>PKIX</tt> from the Verify method list.",
							Computed:            true,
						},
						"verify_val_cred": schema.StringAttribute{
							MarkdownDescription: "The validation credentials can be used to verify the signers' certificate for the JWT. You can get credentials by configuring the Crypto Validation Credentials. This field is meaningful when you select <tt>on</tt> in the Signature validation credentials field.",
							Computed:            true,
						},
						"verify_s_secret": schema.StringAttribute{
							MarkdownDescription: "The shared secret key can be used to verify the JWT signature. This field is meaningful when you select <tt>Verify</tt> in the Validation method field and choose <tt>Shared secret</tt> from the Verify method list.",
							Computed:            true,
						},
						"verify_jwk": schema.StringAttribute{
							MarkdownDescription: "The file containing the JWK or key set is fetched to verify the JWT signature. The file must be stored in the local: or store: directory. This field is meaningful when you select <tt>Verify</tt> in the Validation method field and choose <tt>JWK</tt> from the Verify method list.",
							Computed:            true,
						},
						"verify_fetch_cred_url": schema.StringAttribute{
							MarkdownDescription: "The URL indicates the source location where the verification credentials can be fetched for verifying the JWT signature. The URL must be in the format of http or https. By default, the URL is http://example.com/v3/certs. This field is meaningful when you select <tt>Verify</tt> in the Validation method field and choose <tt>Remotely retrieve JWK</tt> from the Verify method list.",
							Computed:            true,
						},
						"verify_fetch_cred_ssl_profile": schema.StringAttribute{
							MarkdownDescription: "The TLS client profile is provided for fetching the verification credentials. This field is meaningful when you select <tt>Verify</tt> in the Validation method field and choose <tt>Remotely retrieve JWK</tt> from the Verify method list.",
							Computed:            true,
						},
						"claims": schema.ListNestedAttribute{
							MarkdownDescription: "JWT claims need to be validated. You must enter the claim name, value, and data type. If the data type is string, specify the PCRE regular expression to verify the claim value.",
							NestedObject:        models.GetDmClaimDataSourceSchema(),
							Computed:            true,
						},
						"username_claim": schema.StringAttribute{
							MarkdownDescription: "This field is applicable only when the JWT Validator is used in the AAA identity extraction phase. By default, the value of the \"sub\" claim is populated as the username element of the identity extraction output. Ensure that the claim specified in this field is present in the incoming JWT. If no match is found, no username is populated in the AAA processing.",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *AAAJWTValidatorDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *AAAJWTValidatorDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data AAAJWTValidatorList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.AAAJWTValidator{
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
	l := []models.AAAJWTValidator{}
	if resFound {
		if value := res.Get(`AAAJWTValidator`); value.Exists() {
			for _, v := range value.Array() {
				item := models.AAAJWTValidator{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.AAAJWTValidatorObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.AAAJWTValidatorObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
