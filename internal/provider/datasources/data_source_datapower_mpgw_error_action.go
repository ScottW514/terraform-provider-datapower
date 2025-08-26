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

type MPGWErrorActionList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &MPGWErrorActionDataSource{}
	_ datasource.DataSourceWithConfigure = &MPGWErrorActionDataSource{}
)

func NewMPGWErrorActionDataSource() datasource.DataSource {
	return &MPGWErrorActionDataSource{}
}

type MPGWErrorActionDataSource struct {
	pData *tfutils.ProviderData
}

func (d *MPGWErrorActionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mpgw_error_action"
}

func (d *MPGWErrorActionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Define how the Multi-Protocol gateway handles errors and generates error responses to the client.",
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
							MarkdownDescription: "Select which mode to handle the errors and generate the responses. The default mode is \"Static (Local)\".",
							Computed:            true,
						},
						"remote_url": schema.StringAttribute{
							MarkdownDescription: "Specify the URL of the remote error page.",
							Computed:            true,
						},
						"local_url": schema.StringAttribute{
							MarkdownDescription: "Specify the URL of the local error page.",
							Computed:            true,
						},
						"error_rule": schema.StringAttribute{
							MarkdownDescription: "Specify the custom error rule that the appliance runs to handle errors.",
							Computed:            true,
						},
						"status_code": schema.Int64Attribute{
							MarkdownDescription: "Specify the HTTP status code that the appliance returns to the client. Enter a value in the range 0 - 999.",
							Computed:            true,
						},
						"reason_phrase": schema.StringAttribute{
							MarkdownDescription: "Specify the HTTP reason phrase that the appliance returns to the client. For a proxy mode, the specified reason phrase overrides the fetched value.",
							Computed:            true,
						},
						"header_injection": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the name and the value for the HTTP header that the appliance injects.",
							NestedObject:        models.GetDmWebGWErrorRespHeaderInjectionDataSourceSchema(),
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *MPGWErrorActionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *MPGWErrorActionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data MPGWErrorActionList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.MPGWErrorAction{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.MPGWErrorAction{}
	if value := res.Get(`MPGWErrorAction`); value.Exists() {
		for _, v := range value.Array() {
			item := models.MPGWErrorAction{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.MPGWErrorActionObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.MPGWErrorActionObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
