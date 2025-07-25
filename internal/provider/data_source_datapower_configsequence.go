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

type ConfigSequenceList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &ConfigSequenceDataSource{}
	_ datasource.DataSourceWithConfigure = &ConfigSequenceDataSource{}
)

func NewConfigSequenceDataSource() datasource.DataSource {
	return &ConfigSequenceDataSource{}
}

type ConfigSequenceDataSource struct {
	client *client.DatapowerClient
}

func (d *ConfigSequenceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_configsequence"
}

func (d *ConfigSequenceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Configuration sequence",
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
						"locations": schema.ListNestedAttribute{
							MarkdownDescription: "Location profiles",
							NestedObject:        models.DmConfigSequenceLocationDataSourceSchema,
							Computed:            true,
						},
						"match_pattern": schema.StringAttribute{
							MarkdownDescription: "Matching pattern",
							Computed:            true,
						},
						"result_name_pattern": schema.StringAttribute{
							MarkdownDescription: "Result file-naming pattern",
							Computed:            true,
						},
						"status_name_pattern": schema.StringAttribute{
							MarkdownDescription: "Status file-naming pattern",
							Computed:            true,
						},
						"watch": schema.BoolAttribute{
							MarkdownDescription: "Watch",
							Computed:            true,
						},
						"use_output_location": schema.BoolAttribute{
							MarkdownDescription: "Use output location",
							Computed:            true,
						},
						"output_location": schema.StringAttribute{
							MarkdownDescription: "Output file location",
							Computed:            true,
						},
						"delete_unused": schema.BoolAttribute{
							MarkdownDescription: "Delete unused",
							Computed:            true,
						},
						"run_sequence_interval": schema.Int64Attribute{
							MarkdownDescription: "Run interval",
							Computed:            true,
						},
						"capabilities": models.GetDmConfigSequenceCapabilitiesDataSourceSchema("Capabilities", "", ""),
					},
				},
			},
		},
	}
}

func (d *ConfigSequenceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *ConfigSequenceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ConfigSequenceList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.ConfigSequence{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.ConfigSequence{}
	if value := res.Get(`ConfigSequence`); value.Exists() {
		for _, v := range value.Array() {
			item := models.ConfigSequence{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.ConfigSequenceObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.ConfigSequenceObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
