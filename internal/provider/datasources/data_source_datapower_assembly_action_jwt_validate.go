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

type AssemblyActionJWTValidateList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &AssemblyActionJWTValidateDataSource{}
	_ datasource.DataSourceWithConfigure = &AssemblyActionJWTValidateDataSource{}
)

func NewAssemblyActionJWTValidateDataSource() datasource.DataSource {
	return &AssemblyActionJWTValidateDataSource{}
}

type AssemblyActionJWTValidateDataSource struct {
	pData *tfutils.ProviderData
}

func (d *AssemblyActionJWTValidateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_assembly_action_jwt_validate"
}

func (d *AssemblyActionJWTValidateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "<p>The JWT validate assembly action specifies credentials and methods to validate a JWT in the request. The following guidelines apply. <ul><li>All claims that are specified in the JWT validate assembly action are validated. If any specified claim fails, the JWT validation fails.</li><li>You can use a crypto object or a JWK to decrypt or verify the JWT. When both are specified, the crypto object is used.</li><li>If the original message is signed with a shared secret key, the crypto object that is specified must also be a shared secret key.</li><li>If the original message is signed with a private key, the crypto object that is specified must be a crypto certificate (public certificate).</li><li>If a JWK header parameter is included in the header of the JWT, the parameter must match the crypto object or JWK that is specified in the action. Otherwise, validation fails.</li></ul></p>",
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
							MarkdownDescription: "Specify the context variable that contains the JSON Web Token (JWT) to validate. The default value is <tt>request.headers.authorization</tt> . The format of the authorization header must be <tt>Authorization: Bearer jwt-token</tt> , where <tt>jwt-token</tt> indicates the encoded JWT.",
							Computed:            true,
						},
						"output_claims": schema.StringAttribute{
							MarkdownDescription: "Specify a context variable to store all claims that the JWT contains when the validation of the JWT succeeds. The default value is <tt>decoded.claims</tt> .",
							Computed:            true,
						},
						"issuer_claim": schema.StringAttribute{
							MarkdownDescription: "Specify the PCRE to validate the issuer claim. When specified, the \"iss\" claim in the JWT is validated. If this claim fails, validation fails. The maximum value length is 256 characters.",
							Computed:            true,
						},
						"audience_claim": schema.StringAttribute{
							MarkdownDescription: "Specify the PCRE to validate the audience claim. When specified, the \"aud\" claim in the JWT is validated. If this claim fails, the validation fails. The maximum value length is 256 characters.",
							Computed:            true,
						},
						"decrypt_crypto": schema.StringAttribute{
							MarkdownDescription: "Crypto object for JWT decryption",
							Computed:            true,
						},
						"decrypt_jwk": schema.StringAttribute{
							MarkdownDescription: "JWK for JWT decryption",
							Computed:            true,
						},
						"verify_crypto": schema.StringAttribute{
							MarkdownDescription: "Crypto object for JWT verification",
							Computed:            true,
						},
						"verify_jwk": schema.StringAttribute{
							MarkdownDescription: "JWK for JWT verification",
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
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *AssemblyActionJWTValidateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *AssemblyActionJWTValidateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data AssemblyActionJWTValidateList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.AssemblyActionJWTValidate{
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
	l := []models.AssemblyActionJWTValidate{}
	if resFound {
		if value := res.Get(`AssemblyActionJWTValidate`); value.Exists() {
			for _, v := range value.Array() {
				item := models.AssemblyActionJWTValidate{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.AssemblyActionJWTValidateObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.AssemblyActionJWTValidateObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
