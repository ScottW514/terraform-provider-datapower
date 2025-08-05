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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
)

type ParseSettingsList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &ParseSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &ParseSettingsDataSource{}
)

func NewParseSettingsDataSource() datasource.DataSource {
	return &ParseSettingsDataSource{}
}

type ParseSettingsDataSource struct {
	client *client.DatapowerClient
}

func (d *ParseSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_parsesettings"
}

func (d *ParseSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Parse settings",
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
						"document_type": schema.StringAttribute{
							MarkdownDescription: "Document type",
							Computed:            true,
						},
						"document_size": schema.Int64Attribute{
							MarkdownDescription: "Max document size",
							Computed:            true,
						},
						"nesting_depth": schema.Int64Attribute{
							MarkdownDescription: "Max nesting depth",
							Computed:            true,
						},
						"width": schema.Int64Attribute{
							MarkdownDescription: "Max width",
							Computed:            true,
						},
						"name_length": schema.Int64Attribute{
							MarkdownDescription: "Max name length",
							Computed:            true,
						},
						"value_length": schema.Int64Attribute{
							MarkdownDescription: "Max value length",
							Computed:            true,
						},
						"unique_prefixes": schema.Int64Attribute{
							MarkdownDescription: "Max unique prefixes",
							Computed:            true,
						},
						"unique_namespaces": schema.Int64Attribute{
							MarkdownDescription: "Max unique namespaces",
							Computed:            true,
						},
						"unique_names": schema.Int64Attribute{
							MarkdownDescription: "Max unique names",
							Computed:            true,
						},
						"number_length": schema.Int64Attribute{
							MarkdownDescription: "Max number length",
							Computed:            true,
						},
						"strict_utf8_encoding": schema.BoolAttribute{
							MarkdownDescription: "Strict UTF-8 encoding",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *ParseSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *ParseSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ParseSettingsList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.ParseSettings{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.ParseSettings{}
	if value := res.Get(`ParseSettings`); value.Exists() {
		for _, v := range value.Array() {
			item := models.ParseSettings{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.ParseSettingsObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.ParseSettingsObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
