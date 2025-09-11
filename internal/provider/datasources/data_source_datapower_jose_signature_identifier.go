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

type JOSESignatureIdentifierList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &JOSESignatureIdentifierDataSource{}
	_ datasource.DataSourceWithConfigure = &JOSESignatureIdentifierDataSource{}
)

func NewJOSESignatureIdentifierDataSource() datasource.DataSource {
	return &JOSESignatureIdentifierDataSource{}
}

type JOSESignatureIdentifierDataSource struct {
	pData *tfutils.ProviderData
}

func (d *JOSESignatureIdentifierDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_jose_signature_identifier"
}

func (d *JOSESignatureIdentifierDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "JOSE Signature Identifier object for the JSON Web Verify.",
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
						"type": schema.StringAttribute{
							MarkdownDescription: "Key material type used to verify the signature.",
							Computed:            true,
						},
						"sskey": schema.StringAttribute{
							MarkdownDescription: "Use the shared secret key to verify the signature.",
							Computed:            true,
						},
						"certificate": schema.StringAttribute{
							MarkdownDescription: "Use the certificate to verify the signature.",
							Computed:            true,
						},
						"valid_algorithms": schema.ListAttribute{
							MarkdownDescription: "Specifies an array of algorithm values that are valid for signature verification. When specified, the JWS <tt>alg</tt> header parameter value must match a value in this set. By default, all allowed JWS <tt>alg</tt> header parameters values are valid.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"header_param": schema.ListNestedAttribute{
							MarkdownDescription: "The JOSE header parameters used to identify the signature.",
							NestedObject:        models.GetDmJOSEHeaderDataSourceSchema(),
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *JOSESignatureIdentifierDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *JOSESignatureIdentifierDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data JOSESignatureIdentifierList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.JOSESignatureIdentifier{
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
	l := []models.JOSESignatureIdentifier{}
	if resFound {
		if value := res.Get(`JOSESignatureIdentifier`); value.Exists() {
			for _, v := range value.Array() {
				item := models.JOSESignatureIdentifier{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.JOSESignatureIdentifierObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.JOSESignatureIdentifierObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
