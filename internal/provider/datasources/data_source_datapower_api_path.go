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

type APIPathList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &APIPathDataSource{}
	_ datasource.DataSourceWithConfigure = &APIPathDataSource{}
)

func NewAPIPathDataSource() datasource.DataSource {
	return &APIPathDataSource{}
}

type APIPathDataSource struct {
	pData *tfutils.ProviderData
}

func (d *APIPathDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_path"
}

func (d *APIPathDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "An API Path configuration describes the operations that are available on a single path.",
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
						"path": schema.StringAttribute{
							MarkdownDescription: "Specify the relative path to access the REST APIs. The path is appended to the base path to construct the full URI. The path must start with a / character. When the path contains a parameter, ensure that you define the path parameter at either or both the path and operation levels. <ul><li>A parameter at the end of the path can contain a + qualifier to match one or more levels as in the following example. <p><tt>/petstore/{type}/{+category}</tt></p><p>The <tt>{type}</tt> parameter matches one path level. The <tt>{+category}</tt> parameter matches multiple levels. The following paths match this path template.</p><ul><li><tt>/petstore/cats/supplies</tt></li><li><tt>/petstore/cats/supplies/health</tt></li><li><tt>/petstore/cats/supplies/health/medicines</tt></li></ul></li><li>A parameter at the end of the path can contain a * qualifier to match zero or more levels as in the following example. <p><tt>/petstore/{type}/{*category}</tt></p><p>The <tt>{type}</tt> parameter matches one path level. The <tt>{*category}</tt> parameter matches multiple levels. The following paths match this path template.</p><ul><li><tt>/petstore/cats/</tt></li><li><tt>/petstore/cats/supplies</tt></li><li><tt>/petstore/cats/supplies/health</tt></li><li><tt>/petstore/cats/supplies/health/medicines</tt></li></ul></li></ul>",
							Computed:            true,
						},
						"operation": schema.ListAttribute{
							MarkdownDescription: "Specify the available operations for the path. Without a defined operation, all operations are accepted.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"request_schema": schema.StringAttribute{
							MarkdownDescription: "Request schema",
							Computed:            true,
						},
						"parameter": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the applicable parameters for all operations. The setting can be overridden for the same parameter by the setting in the API operation configuration.",
							NestedObject:        models.DmAPIParameterDataSourceSchema,
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *APIPathDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *APIPathDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data APIPathList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.APIPath{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.APIPath{}
	if value := res.Get(`APIPath`); value.Exists() {
		for _, v := range value.Array() {
			item := models.APIPath{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.APIPathObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.APIPathObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
