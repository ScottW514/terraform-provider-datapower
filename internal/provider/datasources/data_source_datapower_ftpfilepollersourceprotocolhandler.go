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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
)

type FTPFilePollerSourceProtocolHandlerList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &FTPFilePollerSourceProtocolHandlerDataSource{}
	_ datasource.DataSourceWithConfigure = &FTPFilePollerSourceProtocolHandlerDataSource{}
)

func NewFTPFilePollerSourceProtocolHandlerDataSource() datasource.DataSource {
	return &FTPFilePollerSourceProtocolHandlerDataSource{}
}

type FTPFilePollerSourceProtocolHandlerDataSource struct {
	client *client.DatapowerClient
}

func (d *FTPFilePollerSourceProtocolHandlerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ftpfilepollersourceprotocolhandler"
}

func (d *FTPFilePollerSourceProtocolHandlerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "FTP poller handler",
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
						"target_directory": schema.StringAttribute{
							MarkdownDescription: "Target directory",
							Computed:            true,
						},
						"delay_between_polls": schema.Int64Attribute{
							MarkdownDescription: "Delay between polls",
							Computed:            true,
						},
						"input_file_match_pattern": schema.StringAttribute{
							MarkdownDescription: "Input file match pattern",
							Computed:            true,
						},
						"processing_rename_pattern": schema.StringAttribute{
							MarkdownDescription: "Processing file renaming pattern",
							Computed:            true,
						},
						"delete_on_success": schema.BoolAttribute{
							MarkdownDescription: "Delete input file on success",
							Computed:            true,
						},
						"success_rename_pattern": schema.StringAttribute{
							MarkdownDescription: "Success file renaming pattern",
							Computed:            true,
						},
						"delete_on_error": schema.BoolAttribute{
							MarkdownDescription: "Delete file on processing error",
							Computed:            true,
						},
						"error_rename_pattern": schema.StringAttribute{
							MarkdownDescription: "Error file renaming pattern",
							Computed:            true,
						},
						"generate_result_file": schema.BoolAttribute{
							MarkdownDescription: "Generate result file",
							Computed:            true,
						},
						"result_name_pattern": schema.StringAttribute{
							MarkdownDescription: "Result file name pattern",
							Computed:            true,
						},
						"processing_seize_timeout": schema.Int64Attribute{
							MarkdownDescription: "Processing seize timeout",
							Computed:            true,
						},
						"processing_seize_pattern": schema.StringAttribute{
							MarkdownDescription: "Processing seize pattern",
							Computed:            true,
						},
						"xml_manager": schema.StringAttribute{
							MarkdownDescription: "XML manager",
							Computed:            true,
						},
						"max_transfers_per_poll": schema.Int64Attribute{
							MarkdownDescription: "Maximum file transfers per poll cycle",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *FTPFilePollerSourceProtocolHandlerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *FTPFilePollerSourceProtocolHandlerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data FTPFilePollerSourceProtocolHandlerList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.FTPFilePollerSourceProtocolHandler{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.FTPFilePollerSourceProtocolHandler{}
	if value := res.Get(`FTPFilePollerSourceProtocolHandler`); value.Exists() {
		for _, v := range value.Array() {
			item := models.FTPFilePollerSourceProtocolHandler{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.FTPFilePollerSourceProtocolHandlerObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.FTPFilePollerSourceProtocolHandlerObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
