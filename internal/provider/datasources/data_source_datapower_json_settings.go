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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type JSONSettingsList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &JSONSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &JSONSettingsDataSource{}
)

func NewJSONSettingsDataSource() datasource.DataSource {
	return &JSONSettingsDataSource{}
}

type JSONSettingsDataSource struct {
	pData *tfutils.ProviderData
}

func (d *JSONSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_json_settings"
}

func (d *JSONSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The JSON settings specify constraints for parsing JSON messages. JSON settings work in conjunction with the parser limits in the XML manager. The more restrictive limits apply.",
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
						"json_max_nesting_depth": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum nesting depth in levels of label-value pairs. Enter a value in the range 64 - 256. The default value is 64.",
							Computed:            true,
						},
						"json_max_label_length": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum length in bytes of the label portion of label-value pair. The length includes any white space between quotation marks. Enter a value in the range 256 - 8192. The default value is 256.",
							Computed:            true,
						},
						"json_max_value_length": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum number in bytes for string values of label-value pairs. The length includes any white space between quotation marks. Enter a value in the range 8192 - 5368709121. The default value is 8192.",
							Computed:            true,
						},
						"json_max_number_length": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum number in bytes for number values of label-value pairs. The number must be a contiguous string of bytes that contain no white space. The number can include a minus sign and a positive or negative exponent. Enter a value in the range 128 - 256. The default value is 128.",
							Computed:            true,
						},
						"json_document_size": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum document size in bytes for the body of the JSON message. If the message is converted to JSONx, the maximum document size specifies the size before the conversion to JSONx. <b>Note:</b> The document size of the JSON message and the size of the JSONx equivalent might differ. Enter a value in the range 4194304 - 5368709121. The default value is 4194304.",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *JSONSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *JSONSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data JSONSettingsList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.JSONSettings{
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
	l := []models.JSONSettings{}
	if resFound {
		if value := res.Get(`JSONSettings`); value.Exists() {
			for _, v := range value.Array() {
				item := models.JSONSettings{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.JSONSettingsObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.JSONSettingsObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
