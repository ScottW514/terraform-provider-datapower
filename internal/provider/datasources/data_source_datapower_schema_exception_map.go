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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type SchemaExceptionMapList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &SchemaExceptionMapDataSource{}
	_ datasource.DataSourceWithConfigure = &SchemaExceptionMapDataSource{}
)

func NewSchemaExceptionMapDataSource() datasource.DataSource {
	return &SchemaExceptionMapDataSource{}
}

type SchemaExceptionMapDataSource struct {
	pData *tfutils.ProviderData
}

func (d *SchemaExceptionMapDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_schema_exception_map"
}

func (d *SchemaExceptionMapDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "This object describes exceptions to an XML Schema by identifying elements of the schema that may be encrypted in an instance document. When Schema Validation is performed on the document using this Exception Map, encrypted elements are considered valid when they appear in place of the specified elements.",
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
						"original_schema_url": schema.StringAttribute{
							MarkdownDescription: "This is the location of the full original schema document. This is required. This may reside on or off the device.",
							Computed:            true,
						},
						"schema_exception_rules": schema.ListNestedAttribute{
							MarkdownDescription: "Rules describe what elements may be excepted from the schema validation when the Original Schema is used for Schema Validation.",
							NestedObject:        models.DmSchemaExceptionRuleDataSourceSchema,
							Computed:            true,
						},
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *SchemaExceptionMapDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *SchemaExceptionMapDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data SchemaExceptionMapList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.SchemaExceptionMap{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.SchemaExceptionMap{}
	if value := res.Get(`SchemaExceptionMap`); value.Exists() {
		for _, v := range value.Array() {
			item := models.SchemaExceptionMap{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.SchemaExceptionMapObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.SchemaExceptionMapObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
