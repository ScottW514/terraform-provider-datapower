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

type AssemblyLogicOperationSwitchList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &AssemblyLogicOperationSwitchDataSource{}
	_ datasource.DataSourceWithConfigure = &AssemblyLogicOperationSwitchDataSource{}
)

func NewAssemblyLogicOperationSwitchDataSource() datasource.DataSource {
	return &AssemblyLogicOperationSwitchDataSource{}
}

type AssemblyLogicOperationSwitchDataSource struct {
	pData *tfutils.ProviderData
}

func (d *AssemblyLogicOperationSwitchDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_assembly_logic_operation_switch"
}

func (d *AssemblyLogicOperationSwitchDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The operation switch assembly action evaluates a list of conditions and runs the API rule that is associated with the first condition that matches the API context.",
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
						"case": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the condition to evaluate and the API rule to run when the condition matches. Conditions are evaluated in order. The API rule of the first condition that evaluates to true is run. If an error occurs during the processing, the transaction fails.",
							NestedObject:        models.DmAssemblyLogicOperationSwitchCaseDataSourceSchema,
							Computed:            true,
						},
						"otherwise": schema.StringAttribute{
							MarkdownDescription: "Otherwise",
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

func (d *AssemblyLogicOperationSwitchDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *AssemblyLogicOperationSwitchDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data AssemblyLogicOperationSwitchList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.AssemblyLogicOperationSwitch{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.AssemblyLogicOperationSwitch{}
	if value := res.Get(`AssemblyLogicOperationSwitch`); value.Exists() {
		for _, v := range value.Array() {
			item := models.AssemblyLogicOperationSwitch{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.AssemblyLogicOperationSwitchObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.AssemblyLogicOperationSwitchObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
