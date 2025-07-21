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

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
)

type JOSESignatureIdentifierList struct {
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
	client *client.DatapowerClient
}

func (d *JOSESignatureIdentifierDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_josesignatureidentifier"
}

func (d *JOSESignatureIdentifierDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Signature Identifier",
		Attributes: map[string]schema.Attribute{
			"app_domain": schema.StringAttribute{
				MarkdownDescription: "The name of the application domain the object belongs to",
				Required:            true,
			},
			"result": schema.ListNestedAttribute{
				MarkdownDescription: "List of objects",
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
						"type": schema.StringAttribute{
							MarkdownDescription: "Key Material Type",
							Computed:            true,
						},
						"ss_key": schema.StringAttribute{
							MarkdownDescription: "Key Material",
							Computed:            true,
						},
						"certificate": schema.StringAttribute{
							MarkdownDescription: "Key Material",
							Computed:            true,
						},
						"valid_algorithms": schema.ListAttribute{
							MarkdownDescription: "Valid algorithms",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"header_param": schema.ListNestedAttribute{
							MarkdownDescription: "Header Parameters",
							NestedObject:        models.DmJOSEHeaderDataSourceSchema,
							Computed:            true,
						},
						"verify": schema.BoolAttribute{
							MarkdownDescription: "Verify (deprecated)",
							Computed:            true,
						},
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

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *JOSESignatureIdentifierDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data JOSESignatureIdentifierList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.JOSESignatureIdentifier{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.JOSESignatureIdentifier{}
	if value := res.Get(`JOSESignatureIdentifier`); value.Exists() {
		for _, v := range value.Array() {
			item := models.JOSESignatureIdentifier{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
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
