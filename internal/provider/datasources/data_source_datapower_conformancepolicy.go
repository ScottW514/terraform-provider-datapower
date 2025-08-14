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

type ConformancePolicyList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &ConformancePolicyDataSource{}
	_ datasource.DataSourceWithConfigure = &ConformancePolicyDataSource{}
)

func NewConformancePolicyDataSource() datasource.DataSource {
	return &ConformancePolicyDataSource{}
}

type ConformancePolicyDataSource struct {
	pData *tfutils.ProviderData
}

func (d *ConformancePolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_conformancepolicy"
}

func (d *ConformancePolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Conformance Policy",
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
						"profiles": models.GetDmConformanceProfilesDataSourceSchema("Profiles", "profiles", ""),
						"ignored_requirements": schema.ListAttribute{
							MarkdownDescription: "Ignored Requirements",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"fixup_stylesheets": schema.ListAttribute{
							MarkdownDescription: "Corrective Stylesheets",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"assert_bp10_conformance": schema.BoolAttribute{
							MarkdownDescription: "BP1.0 Conformance Claim Assertion",
							Computed:            true,
						},
						"report_level": schema.StringAttribute{
							MarkdownDescription: "Record Report",
							Computed:            true,
						},
						"log_target": schema.StringAttribute{
							MarkdownDescription: "Destination",
							Computed:            true,
						},
						"reject_level": schema.StringAttribute{
							MarkdownDescription: "Reject non-conforming messages",
							Computed:            true,
						},
						"reject_include_summary": schema.BoolAttribute{
							MarkdownDescription: "Include error summary",
							Computed:            true,
						},
						"result_is_conformance_report": schema.BoolAttribute{
							MarkdownDescription: "Use analysis as result",
							Computed:            true,
						},
						"response_properties_enabled": schema.BoolAttribute{
							MarkdownDescription: "Distinct response behavior",
							Computed:            true,
						},
						"response_report_level": schema.StringAttribute{
							MarkdownDescription: "Record Report (response direction)",
							Computed:            true,
						},
						"response_log_target": schema.StringAttribute{
							MarkdownDescription: "Destination",
							Computed:            true,
						},
						"response_reject_level": schema.StringAttribute{
							MarkdownDescription: "Reject non-conforming response messages",
							Computed:            true,
						},
						"response_reject_include_summary": schema.BoolAttribute{
							MarkdownDescription: "Include response error summary",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *ConformancePolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *ConformancePolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data ConformancePolicyList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.ConformancePolicy{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.ConformancePolicy{}
	if value := res.Get(`ConformancePolicy`); value.Exists() {
		for _, v := range value.Array() {
			item := models.ConformancePolicy{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.ConformancePolicyObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.ConformancePolicyObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
