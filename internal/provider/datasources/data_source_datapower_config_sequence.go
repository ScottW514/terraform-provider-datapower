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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type ConfigSequenceList struct {
	Id        types.String `tfsdk:"id"`
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
	pData *tfutils.ProviderData
}

func (d *ConfigSequenceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_config_sequence"
}

func (d *ConfigSequenceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "A configuration sequence defines a series of configuration files to load after the startup configuration. By default, changes in configuration files are detected and reloaded.",
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
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"locations": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the locations to watch for changes and the permissions for each location. Each entry specifies a directory where the configuration files to match are stored. The DataPower Gateway watches the location and reloads the configuration when a change is detected that match the PCRE match pattern. The entries are processed in the listed order. The assess profile indicates the permissions for processing.",
							NestedObject:        models.GetDmConfigSequenceLocationDataSourceSchema(),
							Computed:            true,
						},
						"match_pattern": schema.StringAttribute{
							MarkdownDescription: "Specify the PCRE pattern to determine whether a file is considered part of the location match. For example, when the configuration files to match are <tt>NNNNNN.input</tt> , the PCRE pattern is <tt>\"([0-9]{6})\\.input$\"</tt> .",
							Computed:            true,
						},
						"result_name_pattern": schema.StringAttribute{
							MarkdownDescription: "Specify the PCRE pattern to name the result file. This pattern normally has a back-reference to the base input file name. For example, when input files are <tt>NNNNNN.input</tt> and the wanted result file name is <tt>NNNNNN.result</tt> , the pattern is <tt>\"$1.result\"</tt> .",
							Computed:            true,
						},
						"status_name_pattern": schema.StringAttribute{
							MarkdownDescription: "Specify the PCRE pattern to name the status file. This pattern normally has a back-reference to the base input file name. For example, when the input files are <tt>NNNNNN.input</tt> and the wanted status file name is <tt>NNNNNN.json</tt> , the pattern is <tt>\"$1.json\"</tt> .",
							Computed:            true,
						},
						"watch": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to watch the specified directory for configuration file changes and automatically reload the configuration when a change is detected. By default, the specified directory is watched.",
							Computed:            true,
						},
						"use_output_location": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to place output log and status files in the configured output location after processing. By default, output files are stored in the input file location.",
							Computed:            true,
						},
						"output_location": schema.StringAttribute{
							MarkdownDescription: "Specify the directory to store the output files that processing generates. When not specified, the input file location is used.",
							Computed:            true,
						},
						"delete_unused": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to clean up objects that are no longer needed. When enabled, the configuration sequence detects and attempts to delete objects that are no longer modified by any configuration file. By default, the configuration sequence does not delete unneeded objects.",
							Computed:            true,
						},
						"run_sequence_interval": schema.Int64Attribute{
							MarkdownDescription: "Specify the interval in milliseconds between the processing of changes. This delay enables multiple file events to be aggregated and handled within the same sequence run. Enter a value in the range 100 - 60000. The default value is 100.",
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

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *ConfigSequenceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data ConfigSequenceList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.ConfigSequence{
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
	l := []models.ConfigSequence{}
	if resFound {
		if value := res.Get(`ConfigSequence`); value.Exists() {
			for _, v := range value.Array() {
				item := models.ConfigSequence{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
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
