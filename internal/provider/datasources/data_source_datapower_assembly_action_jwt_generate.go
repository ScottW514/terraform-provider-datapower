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

type AssemblyActionJWTGenerateList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &AssemblyActionJWTGenerateDataSource{}
	_ datasource.DataSourceWithConfigure = &AssemblyActionJWTGenerateDataSource{}
)

func NewAssemblyActionJWTGenerateDataSource() datasource.DataSource {
	return &AssemblyActionJWTGenerateDataSource{}
}

type AssemblyActionJWTGenerateDataSource struct {
	pData *tfutils.ProviderData
}

func (d *AssemblyActionJWTGenerateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_assembly_action_jwt_generate"
}

func (d *AssemblyActionJWTGenerateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The JWT generate assembly action generates JWT claims and specifies the crypto material to generate a JWT during processing.",
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
						"jwt": schema.StringAttribute{
							MarkdownDescription: "Specify the variable to store the generated JWT. The default value is <tt>generated.jwt</tt> . When the variable is not set, the generated JWT is written to the Authorization Header as a Bearer token.",
							Computed:            true,
						},
						"jwt_id_claims": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to add a JWT ID (jti) claim to the JWT. When enabled, a UUID is generated and set as the value of the JWT ID claim.",
							Computed:            true,
						},
						"issuer_claim": schema.StringAttribute{
							MarkdownDescription: "Specify the variable from which to retrieve the issuer (iss) claim value. The default value is <tt>iss.claim</tt> . The maximum value length is 256 characters.",
							Computed:            true,
						},
						"subject_claim": schema.StringAttribute{
							MarkdownDescription: "Specify the variable from which to retrieve the subject (sub) claim value. The maximum value length is 256 characters.",
							Computed:            true,
						},
						"audience_claim": schema.StringAttribute{
							MarkdownDescription: "Specify the variable from which to retrieve the audience (aud) claim value. The maximum value length is 256 characters. This value can be a single string, a comma-separated string of values, or an array of one or more values when you set the variable through GatewayScript processing.",
							Computed:            true,
						},
						"validity_period": schema.Int64Attribute{
							MarkdownDescription: "Specify the validity period in seconds to calculate the expiration (exp) claim. This value is added to the current date and time to be the value for the \"exp\" claim. The JWT is considered valid until expiry. Enter a value in the range 1 - 31622400. The default value is 3600.",
							Computed:            true,
						},
						"private_claim": schema.StringAttribute{
							MarkdownDescription: "Specify the variable from which to retrieve a valid set of JSON claims. These claims are added to any set of claims that are specified previously.",
							Computed:            true,
						},
						"sign_jwk": schema.StringAttribute{
							MarkdownDescription: "JWK for JWT signature",
							Computed:            true,
						},
						"crypto_algorithm": schema.StringAttribute{
							MarkdownDescription: "Specify the crypto algorithm to use. Use one of the following values. <ul><li><tt>HS256</tt> - HMAC using SHA-256</li><li><tt>HS384</tt> - HMAC using SHA-384</li><li><tt>HS512</tt> - HMAC using SHA-512</li><li><tt>RS256</tt> - RSASSA-PKCS-v1_5 using SHA-256</li><li><tt>RS384</tt> - RSASSA-PKCS-v1_5 using SHA 384</li><li><tt>RS512</tt> - RSASSA-PKCS-v1_5 using SHA-512</li><li><tt>ES256</tt> - ECDSA using P-256 and SHA-256</li><li><tt>ES384</tt> - ECDSA using P-384 and SHA-384</li><li><tt>ES512</tt> - ECDSA using P-521 and SHA-512</li><li><tt>none</tt> - Do not sign the JWT, which is unsecured and provides no integrity protection but can be used for a nest JWT</li><li>An inline parameter to read at runtime</li></ul>",
							Computed:            true,
						},
						"sign_crypto": schema.StringAttribute{
							MarkdownDescription: "Specify the crypto object (a shared secret key or certificate) to use to sign the JWT.",
							Computed:            true,
						},
						"custom_kid_value_jws": schema.StringAttribute{
							MarkdownDescription: "Specify the value of the <tt>kid</tt> claim of the JWT for JWS. The maximum length is 256 characters. This value can be a single string, a comma-separated string of values, or an array of values when you set the variable through GatewayScript processing.",
							Computed:            true,
						},
						"encrypt_algorithm": schema.StringAttribute{
							MarkdownDescription: "Specify the encryption algorithm to use. Use one of the following values. <ul><li><tt>A128CBC-HS256</tt> - AES_128_CBC_HMAC_SHA_256 authenticated encryption algorithm</li><li><tt>A192CBC-HS384</tt> - AES_192_CBC_HMAC_SHA_384 authenticated encryption algorithm</li><li><tt>A256CBC-HS512</tt> - AES_256_CBC_HMAC_SHA_512 authenticated encryption algorithm</li><li><tt>A128GCM</tt> - AES GCM using 128-bit key</li><li><tt>A192GCM</tt> - AES GCM using 192-bit key</li><li><tt>A256GCM</tt> - AES GCM using 256-bit key</li><li>An inline parameter to read at runtime</li></ul>",
							Computed:            true,
						},
						"encrypt_jwk": schema.StringAttribute{
							MarkdownDescription: "JWK for JWT encryption",
							Computed:            true,
						},
						"key_encrypt_algorithm": schema.StringAttribute{
							MarkdownDescription: "Specify the key encryption algorithm to use. Use one of the following values. <ul><li><tt>RSA1_5</tt> - RSAES-PKCS1-V1_5</li><li><tt>RSA-OAEP</tt> - RSAES OAEP using default parameters</li><li><tt>RSA-OAEP-256</tt> - RSAES OAEP using SHA-256 and MGF1 with SHA-256</li><li><tt>A128KW</tt> - AES Key Wrap with default initial value using 128 bit key</li><li><tt>A192KW</tt> - AES Key Wrap with default initial value using 192 bit key</li><li><tt>A256KW</tt> - AES Key Wrap with default initial value using 256 bit key</li><li><tt>dir</tt> - Direct use of a shared symmetric key as the CEK</li><li>An inline parameter to read at runtime</li></ul>",
							Computed:            true,
						},
						"encrypt_crypto": schema.StringAttribute{
							MarkdownDescription: "Crypto object for JWT encryption",
							Computed:            true,
						},
						"custom_kid_value_jwe": schema.StringAttribute{
							MarkdownDescription: "Specify the value of the <tt>kid</tt> claim of the JWT for JWE. The maximum length is 256 characters. This value can be a single string, a comma-separated string of values, or an array of values when you set the variable through GatewayScript processing.",
							Computed:            true,
						},
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"title": schema.StringAttribute{
							MarkdownDescription: "Title",
							Computed:            true,
						},
						"correlation_path": schema.StringAttribute{
							MarkdownDescription: "Specify the path that correlates the API action to a specific part of the API specification. The correlation path specifies the part of the API definition that correlates with the API action. This path is exposed in the debug data by the API gateway for use by debugging tools. For example, for an API configuration that is retrieved from API Connect and specified in an OpenAPI document with IBM extensions, this path is the JSON path to the assembly policy in the IBM extensions section of the document. The path can be expressed in any form that the debugging tool can correlate to the API definition.",
							Computed:            true,
						},
						"action_debug": schema.BoolAttribute{
							MarkdownDescription: "<p>Specify whether to enable the GatewayScript debugger to troubleshoot the following GatewayScript files or script.</p><ul><li>Troubleshoot a GatewayScript file that is called from the GatewayScript assembly action.</li><li>Troubleshoot a GatewayScript file that is called from an XSLT assembly action that uses the <tt>gatewayscript()</tt> extension function.</li><li>Troubleshoot a GatewayScript script that is called through the <tt>value</tt> or <tt>default</tt> property in the JSON file from the map assembly action.</li></ul><p>To debug a file or script, the following conditions must be met.</p><ul><li>The file contains one or more <tt>debugger;</tt> statements at the points in your script where you want to start debugging.</li><li>The GatewayScript debugger is enabled.</li></ul><p>You run the <tt>debug-action</tt> command.</p>",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *AssemblyActionJWTGenerateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *AssemblyActionJWTGenerateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data AssemblyActionJWTGenerateList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.AssemblyActionJWTGenerate{
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
	l := []models.AssemblyActionJWTGenerate{}
	if resFound {
		if value := res.Get(`AssemblyActionJWTGenerate`); value.Exists() {
			for _, v := range value.Array() {
				item := models.AssemblyActionJWTGenerate{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.AssemblyActionJWTGenerateObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.AssemblyActionJWTGenerateObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
